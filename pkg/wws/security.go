package wws

import (
	"encoding/xml"
	"math/rand"
)

// securityHeader is the soap header that is used to authenticate with Workday web services.
type securityHeader struct {
	XMLName        xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ wsse:Security"`
	XmlNSWsse      string   `xml:"xmlns:wsse,attr"`
	MustUnderstand string   `xml:"mustUnderstand,attr,omitempty"`
	UsernameToken  usernameToken
}

// newSecurityHeader creates a new security header.
func newSecurityHeader(seed int64, username, password string) securityHeader {
	return securityHeader{
		XmlNSWsse:      XMLNSWSS,
		MustUnderstand: "1",
		UsernameToken: usernameToken{
			XmlNSWsse: XMLNSWSS,
			XmlNSWsu:  XMLNSWSU,
			Id:        "UsernameToken-" + randStringBytesMaskImprSrc(seed, 9),
			Username: usernameTokenUsername{
				XmlNSWsse: XMLNSWSS,
				Data:      username,
			},
			Password: usernameTokenPassword{
				XmlNSWsse: XMLNSWSS,
				XmlNSType: XMLNSPasswordText,
				Data:      password,
			},
		},
	}
}

type usernameToken struct {
	XMLName   xml.Name `xml:"wsse:UsernameToken"`
	XmlNSWsu  string   `xml:"xmlns:wsu,attr"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	Id        string   `xml:"wsu:Id,attr,omitempty"`
	Username  usernameTokenUsername
	Password  usernameTokenPassword
}

type usernameTokenUsername struct {
	XMLName   xml.Name `xml:"wsse:Username"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	Data      string   `xml:",chardata"`
}

type usernameTokenPassword struct {
	XMLName   xml.Name `xml:"wsse:Password"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	XmlNSType string   `xml:"Type,attr"`
	Data      string   `xml:",chardata"`
}

// **********
// Accepted solution from http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
// Author: Icza - http://stackoverflow.com/users/1705598/icza

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrc(seed int64, n int) string {
	src := rand.NewSource(seed)
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// **********
