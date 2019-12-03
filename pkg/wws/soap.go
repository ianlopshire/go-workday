package wws

import (
	"encoding/xml"
)

// Envelope is a SOAP envelope.
type envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  header
	Body    body
}

// newEnvelope creates a new envelope with the necessary headers.
func newEnvelope(seed int64, username, password string, bdy interface{}) envelope {
	return envelope{
		Header: header{
			SecurityHeader: newSecurityHeader(seed, username, password),
			CommonHeader:   commonHeader{IncludeReferenceDescriptorsInResponse: true},
		},
		Body: body{
			Content: bdy,
		},
	}
}

// Header contains WWS related SOAP headers.
type header struct {
	XMLName        xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	CommonHeader   commonHeader
	SecurityHeader securityHeader
}

// CommonHeader is the WWS common header.
type commonHeader struct {
	XMLName                               xml.Name `xml:"urn:com.workday/bsvc Workday_Common_Header"`
	IncludeReferenceDescriptorsInResponse bool     `xml:"urn:com.workday/bsvc Include_Reference_Descriptors_In_Response"`
}

// Body is a SOAP body.
type body struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Fault   *Fault      `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

func (b *body) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &Fault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

// Fault is a SOAP fault.
type Fault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	Code    string   `xml:"faultcode,omitempty"`
	String  string   `xml:"faultstring,omitempty"`
	Actor   string   `xml:"faultactor,omitempty"`
	Detail  string   `xml:"detail,omitempty"`
}

func (f *Fault) Error() string {
	return f.String
}
