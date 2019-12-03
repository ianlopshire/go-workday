package extensibility

import (
	"bytes"
	"encoding/xml"
	"time"
)

// Custom List Data
type CustomListDataType struct {
	WebServiceAlias     string                `xml:"urn:com.workday/bsvc Web_Service_Alias,omitempty"`
	CustomFieldTypeName string                `xml:"urn:com.workday/bsvc Custom_Field_Type_Name,omitempty"`
	DoNotUse            *bool                 `xml:"urn:com.workday/bsvc Do_Not_Use,omitempty"`
	CustomListValueData []CustomListValueType `xml:"urn:com.workday/bsvc Custom_List_Value_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomListObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomListObjectType struct {
	ID         []CustomListObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Request Criteria for Bringing back Custom List
type CustomListRequestCriteriaType struct {
	CustomListAlias      string `xml:"urn:com.workday/bsvc Custom_List_Alias,omitempty"`
	CustomListValueAlias string `xml:"urn:com.workday/bsvc Custom_List_Value_Alias,omitempty"`
}

// Custom List Reference
type CustomListRequestReferencesType struct {
	CustomListReference []CustomListObjectType `xml:"urn:com.workday/bsvc Custom_List_Reference"`
}

// Custom List and  Data
type CustomListResponseDataType struct {
	CustomList []CustomListType `xml:"urn:com.workday/bsvc Custom_List,omitempty"`
}

// Custom List Response Group
type CustomListResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Custom List Reference and Data
type CustomListType struct {
	CustomListReference *CustomListObjectType `xml:"urn:com.workday/bsvc Custom_List_Reference,omitempty"`
	CustomListData      []CustomListDataType  `xml:"urn:com.workday/bsvc Custom_List_Data,omitempty"`
}

// Custom List Value Data
type CustomListValueType struct {
	Order           string `xml:"urn:com.workday/bsvc Order,omitempty"`
	ListValueName   string `xml:"urn:com.workday/bsvc List_Value_Name,omitempty"`
	WebServiceAlias string `xml:"urn:com.workday/bsvc Web_Service_Alias"`
	Deprecated      *bool  `xml:"urn:com.workday/bsvc Deprecated,omitempty"`
}

// Request for a Specific Custom List or All Custom Lists
type GetCustomListsRequestType struct {
	RequestReferences *CustomListRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CustomListRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CustomListResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Custom List Response
type GetCustomListsResponseType struct {
	RequestReferences *CustomListRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CustomListRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CustomListResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *CustomListResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

type ProcessingFaultType struct {
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
}

// Request Element . A custom list can be edited by specifying a reference or a webservice alias
type PutCustomListRequestType struct {
	CustomListReference *CustomListObjectType `xml:"urn:com.workday/bsvc Custom_List_Reference,omitempty"`
	CustomListData      *CustomListDataType   `xml:"urn:com.workday/bsvc Custom_List_Data,omitempty"`
	AddOnly             bool                  `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version             string                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response Element for Custom List
type PutCustomListResponseType struct {
	CustomListReference *CustomListObjectType `xml:"urn:com.workday/bsvc Custom_List_Reference,omitempty"`
	Version             string                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Parameters that let you filter the data returned in the response. You can filter returned data by dates and page attributes.
type ResponseFilterType struct {
	AsOfEffectiveDate *time.Time `xml:"urn:com.workday/bsvc As_Of_Effective_Date,omitempty"`
	AsOfEntryDateTime *time.Time `xml:"urn:com.workday/bsvc As_Of_Entry_DateTime,omitempty"`
	Page              float64    `xml:"urn:com.workday/bsvc Page,omitempty"`
	Count             float64    `xml:"urn:com.workday/bsvc Count,omitempty"`
}

func (t *ResponseFilterType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ResponseFilterType
	var layout struct {
		*T
		AsOfEffectiveDate *xsdDate     `xml:"urn:com.workday/bsvc As_Of_Effective_Date,omitempty"`
		AsOfEntryDateTime *xsdDateTime `xml:"urn:com.workday/bsvc As_Of_Entry_DateTime,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AsOfEffectiveDate = (*xsdDate)(layout.T.AsOfEffectiveDate)
	layout.AsOfEntryDateTime = (*xsdDateTime)(layout.T.AsOfEntryDateTime)
	return e.EncodeElement(layout, start)
}
func (t *ResponseFilterType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ResponseFilterType
	var overlay struct {
		*T
		AsOfEffectiveDate *xsdDate     `xml:"urn:com.workday/bsvc As_Of_Effective_Date,omitempty"`
		AsOfEntryDateTime *xsdDateTime `xml:"urn:com.workday/bsvc As_Of_Entry_DateTime,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AsOfEffectiveDate = (*xsdDate)(overlay.T.AsOfEffectiveDate)
	overlay.AsOfEntryDateTime = (*xsdDateTime)(overlay.T.AsOfEntryDateTime)
	return d.DecodeElement(&overlay, &start)
}

// The "Response_Results" element contains summary information about the data that has been returned from your request including "Total_Results", "Total_Pages", and the current "Page" returned.
type ResponseResultsType struct {
	TotalResults float64 `xml:"urn:com.workday/bsvc Total_Results,omitempty"`
	TotalPages   float64 `xml:"urn:com.workday/bsvc Total_Pages,omitempty"`
	PageResults  float64 `xml:"urn:com.workday/bsvc Page_Results,omitempty"`
	Page         float64 `xml:"urn:com.workday/bsvc Page,omitempty"`
}

type ValidationErrorType struct {
	Message       string `xml:"urn:com.workday/bsvc Message,omitempty"`
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
	Xpath         string `xml:"urn:com.workday/bsvc Xpath,omitempty"`
}

type ValidationFaultType struct {
	ValidationError []ValidationErrorType `xml:"urn:com.workday/bsvc Validation_Error,omitempty"`
}

type WorkdayCommonHeaderType struct {
	IncludeReferenceDescriptorsInResponse *bool `xml:"urn:com.workday/bsvc Include_Reference_Descriptors_In_Response,omitempty"`
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02")), nil
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
