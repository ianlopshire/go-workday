//go:generate -command wwsgen go run ../../../../cmd/wwsgen/main.go
//go:generate wwsgen --service=Human_Resources --version=v31.2 --pkg=hr

package hr

// Manually fixed issue after wwsgen:
// ./wwsgen_types.go:16925:45: cannot convert &layout.T.ServerTimestampData (type *[]time.Time) to type *xsdDateTime
// ./wwsgen_types.go:16935:46: cannot convert &overlay.T.ServerTimestampData (type *[]time.Time) to type *xsdDateTime

//func (t *ServerTimestampType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
//	type T ServerTimestampType
//var layout struct {
//	*T
//	ServerTimestampData []xsdDateTime `xml:"urn:com.workday/bsvc Server_Timestamp_Data,omitempty"`
//}
//layout.T = (*T)(t)
//
//for _, d := range t.ServerTimestampData {
//layout.ServerTimestampData = append(layout.ServerTimestampData, xsdDateTime(d))
//}
//
//return e.EncodeElement(layout, start)
//}
//func (t *ServerTimestampType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
//	type T ServerTimestampType
//	var overlay struct {
//		*T
//		ServerTimestampData []xsdDateTime `xml:"urn:com.workday/bsvc Server_Timestamp_Data,omitempty"`
//	}
//	overlay.T = (*T)(t)
//
//	err := d.DecodeElement(&overlay, &start)
//	if err != nil {
//		return err
//	}
//
//	for _, d := range overlay.ServerTimestampData {
//		t.ServerTimestampData = append(t.ServerTimestampData, time.Time(d))
//	}
//
//	return nil
//}
