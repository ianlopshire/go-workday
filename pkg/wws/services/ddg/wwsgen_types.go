package ddg

import (
	"bytes"
	"encoding/xml"
	"time"
)

// Contains a unique identifier for an instance of an object.
type ExternalFieldObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type ExternalFieldObjectType struct {
	ID         []ExternalFieldObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Request element for Get Text Block Categories.
type GetTextBlockCategoriesRequestType struct {
	RequestReferences *TextBlockCategoryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *TextBlockCategoryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *TextBlockCategoryResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Get Text Block Categories.
type GetTextBlockCategoriesResponseType struct {
	RequestReferences *TextBlockCategoryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *TextBlockCategoryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *TextBlockCategoryResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *TextBlockCategoryResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request Criteria for Get Text Blocks.
type GetTextBlocksRequestType struct {
	RequestReferences *TextBlockRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *TextBlockRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Get Text Blocks.
type GetTextBlocksResponseType struct {
	RequestReferences *TextBlockRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *TextBlockRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType            `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *TextBlockResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

type ProcessingFaultType struct {
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
}

// Request element for Put Text Block Category
type PutTextBlockCategoryRequestType struct {
	TextBlockCategoryReference *TextBlockCategoryObjectType `xml:"urn:com.workday/bsvc Text_Block_Category_Reference,omitempty"`
	TextBlockCategoryData      *TextBlockCategoryDataType   `xml:"urn:com.workday/bsvc Text_Block_Category_Data"`
	AddOnly                    bool                         `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Put Text Bock Category.
type PutTextBlockCategoryResponseType struct {
	TextBlockCategoryReference *TextBlockCategoryObjectType `xml:"urn:com.workday/bsvc Text_Block_Category_Reference,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for Put Text Block
type PutTextBlockRequestType struct {
	TextBlockReference *TextBlockObjectType `xml:"urn:com.workday/bsvc Text_Block_Reference,omitempty"`
	TextBlockData      *TextBlockDataType   `xml:"urn:com.workday/bsvc Text_Block_Data"`
	AddOnly            bool                 `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version            string               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Documentation Response element for Put Text Bock
type PutTextBlockResponseType struct {
	TextBlockReference *TextBlockObjectType `xml:"urn:com.workday/bsvc Text_Block_Reference,omitempty"`
	Version            string               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type TemplateTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TemplateTypeObjectType struct {
	ID         []TemplateTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The text block category data.
type TextBlockCategoryDataType struct {
	TextBlockCategoryID   string `xml:"urn:com.workday/bsvc Text_Block_Category_ID,omitempty"`
	TextBlockCategoryName string `xml:"urn:com.workday/bsvc Text_Block_Category_Name"`
}

// Contains a unique identifier for an instance of an object.
type TextBlockCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TextBlockCategoryObjectType struct {
	ID         []TextBlockCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Request Criteria for Get Text Block Categories.
type TextBlockCategoryRequestCriteriaType struct {
}

// The Text Block Category reference(s) to be retrieved.
type TextBlockCategoryRequestReferencesType struct {
	TextBlockCategoryReference []TextBlockCategoryObjectType `xml:"urn:com.workday/bsvc Text_Block_Category_Reference"`
}

// Text Block Category Response Data containing the retrieved Text Block Category information.
type TextBlockCategoryResponseDataType struct {
	TextBlockCategory []TextBlockCategoryType `xml:"urn:com.workday/bsvc Text_Block_Category,omitempty"`
}

// Text Block Category Response Data containing the retrieved Text Block Category information.
type TextBlockCategoryResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// The text block category.
type TextBlockCategoryType struct {
	TextBlockCategoryReference *TextBlockCategoryObjectType `xml:"urn:com.workday/bsvc Text_Block_Category_Reference,omitempty"`
	TextBlockCategoryData      []TextBlockCategoryDataType  `xml:"urn:com.workday/bsvc Text_Block_Category_Data,omitempty"`
}

// The text block data.
type TextBlockDataType struct {
	ID                         string                       `xml:"urn:com.workday/bsvc ID,omitempty"`
	TextBlockName              string                       `xml:"urn:com.workday/bsvc Text_Block_Name"`
	SourceReference            *TemplateTypeObjectType      `xml:"urn:com.workday/bsvc Source_Reference"`
	TextBlockCategoryReference *TextBlockCategoryObjectType `xml:"urn:com.workday/bsvc Text_Block_Category_Reference,omitempty"`
	Inactive                   *bool                        `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	DoNotAllowEditing          *bool                        `xml:"urn:com.workday/bsvc Do_Not_Allow_Editing,omitempty"`
	TextBlockContent           string                       `xml:"urn:com.workday/bsvc Text_Block_Content"`
	ExternalFieldReference     []ExternalFieldObjectType    `xml:"urn:com.workday/bsvc External_Field_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TextBlockObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TextBlockObjectType struct {
	ID         []TextBlockObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Request Criteria for Get Text Blocks.
type TextBlockRequestCriteriaType struct {
	TextBlockCategoryReference *TextBlockCategoryObjectType `xml:"urn:com.workday/bsvc Text_Block_Category_Reference,omitempty"`
	TextBlockName              string                       `xml:"urn:com.workday/bsvc Text_Block_Name,omitempty"`
}

// The Text Block reference(s) to be retrieved.
type TextBlockRequestReferencesType struct {
	TextBlockReference []TextBlockObjectType `xml:"urn:com.workday/bsvc Text_Block_Reference"`
}

// Text Block Response Data containing the retrieved Text Block information.
type TextBlockResponseDataType struct {
	TextBlock []TextBlockType `xml:"urn:com.workday/bsvc Text_Block,omitempty"`
}

// The text block information retrieved as a result of the request.
type TextBlockType struct {
	TextBlockReference *TextBlockObjectType `xml:"urn:com.workday/bsvc Text_Block_Reference,omitempty"`
	TextBlockData      []TextBlockDataType  `xml:"urn:com.workday/bsvc Text_Block_Data,omitempty"`
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
