package wws

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"
)

func TestEnvelope_MarshalXML(t *testing.T) {
	for _, tt := range []struct {
		name        string
		seed        int64
		username    string
		password    string
		body        interface{}
		expectedXML []byte
	}{
		{
			name:        "empty body",
			seed:        time.Time{}.UnixNano(),
			username:    "user@tenant",
			password:    "pass",
			body:        nil,
			expectedXML: []byte(`<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Header xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Workday_Common_Header xmlns="urn:com.workday/bsvc"><Include_Reference_Descriptors_In_Response xmlns="urn:com.workday/bsvc">true</Include_Reference_Descriptors_In_Response></Workday_Common_Header><wsse:Security xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" mustUnderstand="1"><wsse:UsernameToken xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd" xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" wsu:Id="UsernameToken-8wAd9q1uI"><wsse:Username xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd">user@tenant</wsse:Username><wsse:Password xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">pass</wsse:Password></wsse:UsernameToken></wsse:Security></Header><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"></Body></Envelope>`),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			envelope := newEnvelope(tt.seed, tt.username, tt.password, tt.body)
			x, err := xml.Marshal(envelope)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !bytes.Equal(x, tt.expectedXML) {
				t.Errorf("unexpected xml have %s, want %s", x, tt.expectedXML)
			}

		})
	}
}
