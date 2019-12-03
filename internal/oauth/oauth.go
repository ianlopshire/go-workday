// Package oauth implements a sub set of OAuth2 specific to Workday.
package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/golang/groupcache/singleflight"
)

// Token represents a response from the token endpoint.
type Token struct {
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
}

// TokenSource handles fetching and caching a token.
//
// All method on TokenSource are thread safe. Token source should be instantiated
// with the NewTokenSource function.
type TokenSource struct {
	mu           sync.RWMutex
	requestGroup singleflight.Group

	token        Token
	clientID     string
	clientSecret string
	tokenURL     string
	httpClient   *http.Client
	expireAt     time.Time
}

// NewTokenSource creates a new token source.
//
// If the provided http.client is nil, http.Default client will be used.
func NewTokenSource(clientID, clientSecret, refreshToken, tokenURL string, client *http.Client) *TokenSource {
	if client == nil {
		client = http.DefaultClient
	}

	return &TokenSource{
		token: Token{
			RefreshToken: refreshToken,
		},
		clientID:     clientID,
		clientSecret: clientSecret,
		tokenURL:     tokenURL,
		httpClient:   client,
	}
}

// Token returns a token or an error.
//
// The token returned should always be valid. Caching and refreshing the token is
// handled automatically.
func (s *TokenSource) Token() (Token, error) {
	token, ok := s.readToken()
	if ok {
		return token, nil
	}

	return s.refreshToken()
}

// refreshToken refreshes the token.
//
// A sync.Singleflight group is used to ensure only a single token is being
// fetched at any given time.
func (s *TokenSource) refreshToken() (Token, error) {
	token, err := s.requestGroup.Do("token", func() (i interface{}, e error) {

		fmt.Println("fetching new token")
		token, _ := s.readToken()
		params := url.Values{
			"grant_type":    []string{"refresh_token"},
			"refresh_token": []string{token.RefreshToken},
		}

		req, err := http.NewRequest("GET", s.tokenURL, strings.NewReader(params.Encode()))
		if err != nil {
			return Token{}, err
		}
		req.SetBasicAuth(s.clientID, s.clientSecret)

		res, err := s.httpClient.Do(req)
		if err != nil {
			return Token{}, err
		}
		defer res.Body.Close()

		err = json.NewDecoder(res.Body).Decode(&token)
		if err != nil {
			return Token{}, err
		}

		s.writeToken(token, time.Now().Add(60*time.Second))
		return token, nil
	})
	return token.(Token), err
}

// writeToken writes the token to cache in a thread safe manner.
func (s *TokenSource) writeToken(token Token, expireAt time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.expireAt = expireAt
	s.token = token
}

// readToken reads the token from cache in a thread safe manner.
//
// If the token in not valid, false will be returned in the
// second return parameter.
func (s *TokenSource) readToken() (Token, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.token.AccessToken == "" || s.expireAt.Before(time.Now()) {
		return s.token, false
	}

	return s.token, true
}
