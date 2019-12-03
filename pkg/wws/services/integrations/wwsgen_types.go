package integrations

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// AS2 Settings Data
type AS2SettingsDataType struct {
	AS2FromID                         string                        `xml:"urn:com.workday/bsvc AS2_From_ID"`
	AS2ID                             string                        `xml:"urn:com.workday/bsvc AS2_ID"`
	PublicKeyforEncryptionReference   *X509PublicKeyObjectType      `xml:"urn:com.workday/bsvc Public_Key_for_Encryption_Reference"`
	PrivateKeyPairforSigningReference *X509PrivateKeyPairObjectType `xml:"urn:com.workday/bsvc Private_Key_Pair_for_Signing_Reference"`
	PublicKeyforSSLReference          *X509PublicKeyObjectType      `xml:"urn:com.workday/bsvc Public_Key_for_SSL_Reference,omitempty"`
	PublicKeyforSSLAcceptReference    *X509PublicKeyObjectType      `xml:"urn:com.workday/bsvc Public_Key_for_SSL_Accept_Reference,omitempty"`
}

// AS2 Transport Protocol Data element
type AS2TransportProtocolDataType struct {
	AS2Endpoint     string               `xml:"urn:com.workday/bsvc AS2_Endpoint"`
	AS2SettingsData *AS2SettingsDataType `xml:"urn:com.workday/bsvc AS2_Settings_Data"`
}

// Contains a unique identifier for an instance of an object.
type AS2TransportProtocolObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AS2TransportProtocolObjectType struct {
	ID         []AS2TransportProtocolObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AbstractExternalParameterObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AbstractExternalParameterObjectType struct {
	ID         []AbstractExternalParameterObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Sequence Generator Data
type AbstractSequenceGeneratorDataType struct {
	SequenceID                   string                  `xml:"urn:com.workday/bsvc Sequence_ID,omitempty"`
	SequenceName                 string                  `xml:"urn:com.workday/bsvc Sequence_Name"`
	StartNumber                  float64                 `xml:"urn:com.workday/bsvc Start_Number,omitempty"`
	IncrementBy                  float64                 `xml:"urn:com.workday/bsvc Increment_By"`
	RestartDateIntervalReference *DateIntervalObjectType `xml:"urn:com.workday/bsvc Restart_Date_Interval_Reference,omitempty"`
	UseTimeZoneReference         *TimeZoneObjectType     `xml:"urn:com.workday/bsvc Use_Time_Zone_Reference,omitempty"`
	LastNumberUsed               float64                 `xml:"urn:com.workday/bsvc Last_Number_Used,omitempty"`
	LastDateTimeUsed             *time.Time              `xml:"urn:com.workday/bsvc Last_DateTime_Used,omitempty"`
	PaddingWithZero              float64                 `xml:"urn:com.workday/bsvc Padding_With_Zero,omitempty"`
	Format                       string                  `xml:"urn:com.workday/bsvc Format"`
}

func (t *AbstractSequenceGeneratorDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AbstractSequenceGeneratorDataType
	var layout struct {
		*T
		LastDateTimeUsed *xsdDateTime `xml:"urn:com.workday/bsvc Last_DateTime_Used,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LastDateTimeUsed = (*xsdDateTime)(layout.T.LastDateTimeUsed)
	return e.EncodeElement(layout, start)
}
func (t *AbstractSequenceGeneratorDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AbstractSequenceGeneratorDataType
	var overlay struct {
		*T
		LastDateTimeUsed *xsdDateTime `xml:"urn:com.workday/bsvc Last_DateTime_Used,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LastDateTimeUsed = (*xsdDateTime)(overlay.T.LastDateTimeUsed)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type ActionEventObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ActionEventObjectType struct {
	ID         []ActionEventObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing application related exceptions data
type ApplicationInstanceExceptionsDataType struct {
	ExceptionData []ExceptionDataType `xml:"urn:com.workday/bsvc Exception_Data,omitempty"`
}

// Element containing Exceptions Data
type ApplicationInstanceRelatedExceptionsDataType struct {
	ExceptionsData []ApplicationInstanceExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Data,omitempty"`
}

// Top-level request element for the Approve Business Process operation
type ApproveBusinessProcessRequestType struct {
	EventReference             *EventObjectType         `xml:"urn:com.workday/bsvc Event_Reference"`
	ApproveBusinessProcessData *BusinessProcessDataType `xml:"urn:com.workday/bsvc Approve_Business_Process_Data,omitempty"`
	Version                    string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper that contains all sub response elements.
type ApproveBusinessProcessResponseType struct {
	EventReference *EventObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The details of a resume (resume, name, comments).
type AttachmentDataWWSType struct {
	FileID   string  `xml:"urn:com.workday/bsvc File_ID,omitempty"`
	File     *[]byte `xml:"urn:com.workday/bsvc File,omitempty"`
	FileName string  `xml:"urn:com.workday/bsvc FileName,omitempty"`
	Comments string  `xml:"urn:com.workday/bsvc Comments,omitempty"`
}

func (t *AttachmentDataWWSType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AttachmentDataWWSType
	var layout struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	layout.T = (*T)(t)
	layout.File = (*xsdBase64Binary)(layout.T.File)
	return e.EncodeElement(layout, start)
}
func (t *AttachmentDataWWSType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AttachmentDataWWSType
	var overlay struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.File = (*xsdBase64Binary)(overlay.T.File)
	return d.DecodeElement(&overlay, &start)
}

// This element is wrapper for Event data.
type AwaitingTaskDataType struct {
	TaskReference           *UniqueIdentifierObjectType  `xml:"urn:com.workday/bsvc Task_Reference,omitempty"`
	CommentReference        []CommentObjectType          `xml:"urn:com.workday/bsvc Comment_Reference,omitempty"`
	TaskStatusReference     *EventRecordActionObjectType `xml:"urn:com.workday/bsvc Task_Status_Reference,omitempty"`
	AssignmentDate          *time.Time                   `xml:"urn:com.workday/bsvc Assignment_Date,omitempty"`
	DueDate                 *time.Time                   `xml:"urn:com.workday/bsvc Due_Date,omitempty"`
	AwaitingPersonReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Awaiting_Person_Reference,omitempty"`
}

func (t *AwaitingTaskDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AwaitingTaskDataType
	var layout struct {
		*T
		AssignmentDate *xsdDateTime `xml:"urn:com.workday/bsvc Assignment_Date,omitempty"`
		DueDate        *xsdDate     `xml:"urn:com.workday/bsvc Due_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssignmentDate = (*xsdDateTime)(layout.T.AssignmentDate)
	layout.DueDate = (*xsdDate)(layout.T.DueDate)
	return e.EncodeElement(layout, start)
}
func (t *AwaitingTaskDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AwaitingTaskDataType
	var overlay struct {
		*T
		AssignmentDate *xsdDateTime `xml:"urn:com.workday/bsvc Assignment_Date,omitempty"`
		DueDate        *xsdDate     `xml:"urn:com.workday/bsvc Due_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssignmentDate = (*xsdDateTime)(overlay.T.AssignmentDate)
	overlay.DueDate = (*xsdDate)(overlay.T.DueDate)
	return d.DecodeElement(&overlay, &start)
}

// Message
type BackgroundProcessMessageDataWSType struct {
	Timestamp      *time.Time `xml:"urn:com.workday/bsvc Timestamp,omitempty"`
	Severity       string     `xml:"urn:com.workday/bsvc Severity,omitempty"`
	MessageSummary string     `xml:"urn:com.workday/bsvc Message_Summary,omitempty"`
	LineNumber     float64    `xml:"urn:com.workday/bsvc Line_Number,omitempty"`
	LineIdentifier string     `xml:"urn:com.workday/bsvc Line_Identifier,omitempty"`
}

func (t *BackgroundProcessMessageDataWSType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BackgroundProcessMessageDataWSType
	var layout struct {
		*T
		Timestamp *xsdDateTime `xml:"urn:com.workday/bsvc Timestamp,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Timestamp = (*xsdDateTime)(layout.T.Timestamp)
	return e.EncodeElement(layout, start)
}
func (t *BackgroundProcessMessageDataWSType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BackgroundProcessMessageDataWSType
	var overlay struct {
		*T
		Timestamp *xsdDateTime `xml:"urn:com.workday/bsvc Timestamp,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Timestamp = (*xsdDateTime)(overlay.T.Timestamp)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type BackgroundProcessMessageSeverityLevelObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BackgroundProcessMessageSeverityLevelObjectType struct {
	ID         []BackgroundProcessMessageSeverityLevelObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BackgroundProcessRuntimeStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BackgroundProcessRuntimeStatusObjectType struct {
	ID         []BackgroundProcessRuntimeStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BusinessObjectObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BusinessObjectObjectType struct {
	ID         []BusinessObjectObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing the data for the Denial (Comment)
type BusinessProcessDataType struct {
	Comment string `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BusinessProcessTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BusinessProcessTypeObjectType struct {
	ID         []BusinessProcessTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Provide additional details for the Business Process cancelation.
type CancelBusinessProcessDataType struct {
	Comment               string `xml:"urn:com.workday/bsvc Comment,omitempty"`
	SuppressNotifications *bool  `xml:"urn:com.workday/bsvc Suppress_Notifications,omitempty"`
}

// Cancel Business Process Request
type CancelBusinessProcessRequestType struct {
	EventReference            *ActionEventObjectType         `xml:"urn:com.workday/bsvc Event_Reference"`
	CancelBusinessProcessData *CancelBusinessProcessDataType `xml:"urn:com.workday/bsvc Cancel_Business_Process_Data,omitempty"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Cancel Business Process Response
type CancelBusinessProcessResponseType struct {
	EventReference *ActionEventObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ChecklistObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ChecklistObjectType struct {
	ID         []ChecklistObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type ChecklistReferenceEnumeration string

// Contains a unique identifier for an instance of an object.
type CloudCollectionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CloudCollectionObjectType struct {
	ID         []CloudCollectionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CommentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CommentObjectType struct {
	ID         []CommentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element to configure Concurrency on Integration Background Future Process
type ConcurrencyConfigurationDataType struct {
	ConcurrencyOption      *bool                        `xml:"urn:com.workday/bsvc Concurrency_Option,omitempty"`
	ParameterExceptionData []ParameterExceptionDataType `xml:"urn:com.workday/bsvc Parameter_Exception_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ConditionRuleObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type ConditionRuleObjectType struct {
	ID         []ConditionRuleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Container element for an Abstract Value.
type CopyofIntegrationAbstractValueDataType struct {
	Text              string               `xml:"urn:com.workday/bsvc Text,omitempty"`
	Date              *time.Time           `xml:"urn:com.workday/bsvc Date,omitempty"`
	DateTime          *time.Time           `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	Numeric           float64              `xml:"urn:com.workday/bsvc Numeric,omitempty"`
	Boolean           *bool                `xml:"urn:com.workday/bsvc Boolean,omitempty"`
	InstanceReference []InstanceObjectType `xml:"urn:com.workday/bsvc Instance_Reference,omitempty"`
}

func (t *CopyofIntegrationAbstractValueDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CopyofIntegrationAbstractValueDataType
	var layout struct {
		*T
		Date     *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		DateTime *xsdDateTime `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Date = (*xsdDate)(layout.T.Date)
	layout.DateTime = (*xsdDateTime)(layout.T.DateTime)
	return e.EncodeElement(layout, start)
}
func (t *CopyofIntegrationAbstractValueDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CopyofIntegrationAbstractValueDataType
	var overlay struct {
		*T
		Date     *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		DateTime *xsdDateTime `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Date = (*xsdDate)(overlay.T.Date)
	overlay.DateTime = (*xsdDateTime)(overlay.T.DateTime)
	return d.DecodeElement(&overlay, &start)
}

type Currency float64

// Contains a unique identifier for an instance of an object.
type CurrencyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CurrencyObjectType struct {
	ID         []CurrencyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomBusinessFormLayoutObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomBusinessFormLayoutObjectType struct {
	ID         []CustomBusinessFormLayoutObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomObjectObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type CustomObjectObjectType struct {
	ID         []CustomObjectObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomReportDefinitionObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type CustomReportDefinitionObjectType struct {
	ID         []CustomReportDefinitionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Custom Report Transformation Data element
type CustomReportTransformationDataType struct {
	CustomReportTransformationReference *CustomReportTransformationObjectType `xml:"urn:com.workday/bsvc Custom_Report_Transformation_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomReportTransformationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomReportTransformationObjectType struct {
	ID         []CustomReportTransformationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DateIntervalObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DateIntervalObjectType struct {
	ID         []DateIntervalObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DeliveredBusinessFormLayoutObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DeliveredBusinessFormLayoutObjectType struct {
	ID         []DeliveredBusinessFormLayoutObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Delivered Transformation Data element
type DeliveredTransformationDataType struct {
	WorkdayTransformationReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Workday_Transformation_Reference"`
	TemplateModelReference         *TemplateModelObjectType    `xml:"urn:com.workday/bsvc Template_Model_Reference"`
	CustomObjectReference          *CustomObjectObjectType     `xml:"urn:com.workday/bsvc Custom_Object_Reference"`
}

// Top-level request element for the Deny Business Process operation
type DenyBusinessProcessRequestType struct {
	EventReference          *EventObjectType         `xml:"urn:com.workday/bsvc Event_Reference"`
	DenyBusinessProcessData *BusinessProcessDataType `xml:"urn:com.workday/bsvc Deny_Business_Process_Data,omitempty"`
	Version                 string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper that contains all sub response elements.
type DenyBusinessProcessResponseType struct {
	EventReference *EventObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DisplayOptionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DisplayOptionObjectType struct {
	ID         []DisplayOptionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DocumentTagObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DocumentTagObjectType struct {
	ID         []DocumentTagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DocumentTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DocumentTypeObjectType struct {
	ID         []DocumentTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Dynamic Filename Definition Assignment Data element
type DynamicFilenameDefinitionAssignmentDataType struct {
	TagReference           *UniqueIdentifierObjectType                               `xml:"urn:com.workday/bsvc Tag_Reference"`
	ReportPromptReference  *ExternalPromptableObjectType                             `xml:"urn:com.workday/bsvc Report_Prompt_Reference"`
	ExternalFieldReference *ExternalFieldObjectType                                  `xml:"urn:com.workday/bsvc External_Field_Reference"`
	FieldPrompt            []IntegrationFieldOverrideParameterInitializationDataType `xml:"urn:com.workday/bsvc Field_Prompt,omitempty"`
}

// Dynamic Filename Definition Data element
type DynamicFilenameDefinitionDataType struct {
	FilenameDefinition                      string                                        `xml:"urn:com.workday/bsvc Filename_Definition,omitempty"`
	DynamicFilenameDefinitionAssignmentData []DynamicFilenameDefinitionAssignmentDataType `xml:"urn:com.workday/bsvc Dynamic_Filename_Definition_Assignment_Data,omitempty"`
}

// Attachment Data Source Data element
type EIBAttachmentDataSourceDataType struct {
	LoadFromAttachment           bool                           `xml:"urn:com.workday/bsvc Load_From_Attachment"`
	WebServiceOperationReference *WebServiceOperationObjectType `xml:"urn:com.workday/bsvc Web_Service_Operation_Reference,omitempty"`
	CustomObjectReference        *CustomObjectObjectType        `xml:"urn:com.workday/bsvc Custom_Object_Reference,omitempty"`
}

// Extensibility API Transport Protocol Data element
type EIBExtensibilityAPITransportProtocolDataType struct {
	CustomObjectReference *CustomObjectObjectType `xml:"urn:com.workday/bsvc Custom_Object_Reference,omitempty"`
}

// External File Data Source Data element
type EIBExternalFileDataSourceDataType struct {
	IntegrationTransportProtocolData *EIBIntegrationTransportProtocolDataWWSType `xml:"urn:com.workday/bsvc Integration_Transport_Protocol_Data"`
	WebServiceOperationReference     *WebServiceOperationObjectType              `xml:"urn:com.workday/bsvc Web_Service_Operation_Reference,omitempty"`
	CustomObjectReference            *CustomObjectObjectType                     `xml:"urn:com.workday/bsvc Custom_Object_Reference,omitempty"`
}

// Integration Transport Protocol Data element
type EIBIntegrationTransportProtocolDataWWSType struct {
	ID                             string                                               `xml:"urn:com.workday/bsvc ID,omitempty"`
	FTPSTransportProtocolReference *FTPSSLTransportProtocolObjectType                   `xml:"urn:com.workday/bsvc FTPS_Transport_Protocol_Reference,omitempty"`
	FTPSTransportProtocolData      *FTPSTransportProtocolDataType                       `xml:"urn:com.workday/bsvc FTPS_Transport_Protocol_Data,omitempty"`
	SFTPTransportProtocolReference *SFTPTransportProtocolObjectType                     `xml:"urn:com.workday/bsvc SFTP_Transport_Protocol_Reference,omitempty"`
	SFTPTransportProtocolData      *SFTPTransportProtocolDatawithDualAuthenticationType `xml:"urn:com.workday/bsvc SFTP_Transport_Protocol_Data,omitempty"`
}

// Email Transport Protocol Data element
type EmailTransportProtocolDataType struct {
	ToEmailAddress           string `xml:"urn:com.workday/bsvc To_Email_Address"`
	CcInternetEmailAddress   string `xml:"urn:com.workday/bsvc Cc_Internet_Email_Address,omitempty"`
	BccInternetEmailAddress  string `xml:"urn:com.workday/bsvc Bcc_Internet_Email_Address,omitempty"`
	FromInternetEmailAddress string `xml:"urn:com.workday/bsvc From_Internet_Email_Address"`
	EmailSubject             string `xml:"urn:com.workday/bsvc Email_Subject"`
	EmailText                string `xml:"urn:com.workday/bsvc Email_Text,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EmailTransportProtocolObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmailTransportProtocolObjectType struct {
	ID         []EmailTransportProtocolObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EventClassificationSubcategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EventClassificationSubcategoryObjectType struct {
	ID         []EventClassificationSubcategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This element is wrapper for Event data.
type EventDetailDataType struct {
	CreationDate              *time.Time                                `xml:"urn:com.workday/bsvc Creation_Date,omitempty"`
	DueDate                   *time.Time                                `xml:"urn:com.workday/bsvc Due_Date,omitempty"`
	CompletedDate             *time.Time                                `xml:"urn:com.workday/bsvc Completed_Date,omitempty"`
	EventCategoryReference    *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Event_Category_Reference,omitempty"`
	EventStateReference       *UniqueIdentifierObjectType               `xml:"urn:com.workday/bsvc Event_State_Reference,omitempty"`
	EventTargetReference      *EventTargetObjectType                    `xml:"urn:com.workday/bsvc Event_Target_Reference,omitempty"`
	InitiatingPersonReference *UniqueIdentifierObjectType               `xml:"urn:com.workday/bsvc Initiating_Person_Reference,omitempty"`
	AwaitingTaskData          []AwaitingTaskDataType                    `xml:"urn:com.workday/bsvc Awaiting_Task_Data,omitempty"`
	SubEvent                  []EventWWSType                            `xml:"urn:com.workday/bsvc Sub-Event,omitempty"`
	ProcessHistoryData        []EventRecordProcessViewWWSType           `xml:"urn:com.workday/bsvc Process_History_Data,omitempty"`
	RemainingProcessData      []EventRemainingProcessViewWWSType        `xml:"urn:com.workday/bsvc Remaining_Process_Data,omitempty"`
}

func (t *EventDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EventDetailDataType
	var layout struct {
		*T
		CreationDate  *xsdDateTime `xml:"urn:com.workday/bsvc Creation_Date,omitempty"`
		DueDate       *xsdDate     `xml:"urn:com.workday/bsvc Due_Date,omitempty"`
		CompletedDate *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreationDate = (*xsdDateTime)(layout.T.CreationDate)
	layout.DueDate = (*xsdDate)(layout.T.DueDate)
	layout.CompletedDate = (*xsdDateTime)(layout.T.CompletedDate)
	return e.EncodeElement(layout, start)
}
func (t *EventDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EventDetailDataType
	var overlay struct {
		*T
		CreationDate  *xsdDateTime `xml:"urn:com.workday/bsvc Creation_Date,omitempty"`
		DueDate       *xsdDate     `xml:"urn:com.workday/bsvc Due_Date,omitempty"`
		CompletedDate *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreationDate = (*xsdDateTime)(overlay.T.CreationDate)
	overlay.DueDate = (*xsdDate)(overlay.T.DueDate)
	overlay.CompletedDate = (*xsdDateTime)(overlay.T.CompletedDate)
	return d.DecodeElement(&overlay, &start)
}

// Event Documents Response Data element
type EventDocumentsResponseDataType struct {
	EventDocuments []EventDocumentsType `xml:"urn:com.workday/bsvc Event_Documents,omitempty"`
}

// For each Integration Event, this element contains all the available Documents
type EventDocumentsType struct {
	IntegrationEventReference *IntegrationEventAbstractObjectType `xml:"urn:com.workday/bsvc Integration_Event_Reference,omitempty"`
	RepositoryDocument        []RepositoryDocumentSummaryDataType `xml:"urn:com.workday/bsvc Repository_Document,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EventObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EventObjectType struct {
	ID         []EventObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EventRecordActionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EventRecordActionObjectType struct {
	ID         []EventRecordActionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// All completed or in progress steps for the event
type EventRecordProcessViewWWSType struct {
	ProcessReference                    *ActionEventObjectType                 `xml:"urn:com.workday/bsvc Process_Reference,omitempty"`
	WorkflowStep                        *WorkflowStepType                      `xml:"urn:com.workday/bsvc Workflow_Step,omitempty"`
	StatusReference                     *EventRecordActionObjectType           `xml:"urn:com.workday/bsvc Status_Reference,omitempty"`
	CompletedDate                       *time.Time                             `xml:"urn:com.workday/bsvc Completed_Date,omitempty"`
	DueDate                             *time.Time                             `xml:"urn:com.workday/bsvc Due_Date,omitempty"`
	WorkflowProcessParticipantReference []WorkflowProcessParticipantObjectType `xml:"urn:com.workday/bsvc Workflow_Process_Participant_Reference,omitempty"`
}

func (t *EventRecordProcessViewWWSType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EventRecordProcessViewWWSType
	var layout struct {
		*T
		CompletedDate *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Date,omitempty"`
		DueDate       *xsdDate     `xml:"urn:com.workday/bsvc Due_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CompletedDate = (*xsdDateTime)(layout.T.CompletedDate)
	layout.DueDate = (*xsdDate)(layout.T.DueDate)
	return e.EncodeElement(layout, start)
}
func (t *EventRecordProcessViewWWSType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EventRecordProcessViewWWSType
	var overlay struct {
		*T
		CompletedDate *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Date,omitempty"`
		DueDate       *xsdDate     `xml:"urn:com.workday/bsvc Due_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CompletedDate = (*xsdDateTime)(overlay.T.CompletedDate)
	overlay.DueDate = (*xsdDate)(overlay.T.DueDate)
	return d.DecodeElement(&overlay, &start)
}

// All remaining steps for the event
type EventRemainingProcessViewWWSType struct {
	EventReference *EventObjectType   `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	WorkflowStep   []WorkflowStepType `xml:"urn:com.workday/bsvc Workflow_Step,omitempty"`
}

// This element contains references to Event for data retrieval.
type EventRequestReferencesType struct {
	EventReference []ActionEventObjectType `xml:"urn:com.workday/bsvc Event_Reference"`
}

// Element containing the response data with event details.
type EventResponseDataType struct {
	Event []EventWWSType `xml:"urn:com.workday/bsvc Event,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EventServiceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EventServiceObjectType struct {
	ID         []EventServiceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type EventServiceReferenceEnumeration string

// Contains a unique identifier for an instance of an object.
type EventTargetObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type EventTargetObjectType struct {
	ID         []EventTargetObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This element is wrapper for Event data.
type EventWWSType struct {
	EventReference  *ActionEventObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	EventDetailData *EventDetailDataType   `xml:"urn:com.workday/bsvc Event_Detail_Data,omitempty"`
}

// Exception (Errors and Warning) associated with the transaction.
type ExceptionDataType struct {
	Classification string `xml:"urn:com.workday/bsvc Classification,omitempty"`
	Message        string `xml:"urn:com.workday/bsvc Message,omitempty"`
}

// Integration Document Field for Document Extension Data
type ExtendedIntegrationDocumentFieldDataType struct {
	FieldName                                                             string                                                    `xml:"urn:com.workday/bsvc Field_Name"`
	Description                                                           string                                                    `xml:"urn:com.workday/bsvc Description,omitempty"`
	FieldDataTypeReference                                                *UniqueIdentifierObjectType                               `xml:"urn:com.workday/bsvc Field_Data_Type_Reference,omitempty"`
	ReferenceIDType                                                       string                                                    `xml:"urn:com.workday/bsvc Reference_ID_Type,omitempty"`
	ExternalFieldReference                                                *ExternalFieldObjectType                                  `xml:"urn:com.workday/bsvc External_Field_Reference"`
	IntegrationFieldOverrideParameterInitializationData                   []IntegrationFieldOverrideParameterInitializationDataType `xml:"urn:com.workday/bsvc Integration_Field_Override_Parameter_Initialization_Data,omitempty"`
	IntegrationDocumentStacksforIntegrationDocumentFieldConfigurationData []IntegrationDocumentStackDataType                        `xml:"urn:com.workday/bsvc Integration_Document_Stacks_for_Integration_Document_Field_Configuration_Data,omitempty"`
	IntegrationDocumentFieldOptions                                       []IntegrationDocumentFieldOptionsType                     `xml:"urn:com.workday/bsvc Integration_Document_Field_Options,omitempty"`
}

// Endpoint Info Data element
type ExternalEndpointDataType struct {
	WebServiceAPIVersionReference *WebServiceAPIVersionObjectType `xml:"urn:com.workday/bsvc Web_Service_API_Version_Reference"`
	UseDeployedServiceEndpoint    *bool                           `xml:"urn:com.workday/bsvc Use_Deployed_Service_Endpoint,omitempty"`
	SubscriberURL                 string                          `xml:"urn:com.workday/bsvc Subscriber_URL,omitempty"`
	SubscriptionUserName          string                          `xml:"urn:com.workday/bsvc Subscription_User_Name,omitempty"`
	SubscriptionPassword          string                          `xml:"urn:com.workday/bsvc Subscription_Password,omitempty"`
	DisableEndpoint               *bool                           `xml:"urn:com.workday/bsvc Disable_Endpoint,omitempty"`
	OMSEnvironmentReference       []OMSEnvironmentObjectType      `xml:"urn:com.workday/bsvc OMS_Environment_Reference,omitempty"`
}

// Container element for External Field definition.
type ExternalFieldAddorReferenceType struct {
	ClassReportFieldReference  *ExternalFieldObjectType  `xml:"urn:com.workday/bsvc Class_Report_Field_Reference,omitempty"`
	CalculatedFieldClassName   string                    `xml:"urn:com.workday/bsvc Calculated_Field_Class_Name,omitempty"`
	CalculatedFieldReferenceID string                    `xml:"urn:com.workday/bsvc Calculated_Field_Reference_ID,omitempty"`
	CalculatedFieldName        string                    `xml:"urn:com.workday/bsvc Calculated_Field_Name,omitempty"`
	BusinessObjectReference    *BusinessObjectObjectType `xml:"urn:com.workday/bsvc Business_Object_Reference,omitempty"`
}

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

// Contains a unique identifier for an instance of an object.
type ExternalInstanceIDType struct {
	Value      string `xml:",chardata"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
	Type       string `xml:"urn:com.workday/bsvc type,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ExternalInstanceObjectType struct {
	ID         []ExternalInstanceIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ExternalPromptableObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type ExternalPromptableObjectType struct {
	ID         []ExternalPromptableObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type FTPSSLTransportProtocolObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type FTPSSLTransportProtocolObjectType struct {
	ID         []FTPSSLTransportProtocolObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// FTPS Transport Protocol Data element
type FTPSTransportProtocolDataType struct {
	FTPSAddress    string `xml:"urn:com.workday/bsvc FTPS_Address"`
	Directory      string `xml:"urn:com.workday/bsvc Directory,omitempty"`
	UsePassiveMode bool   `xml:"urn:com.workday/bsvc Use_Passive_Mode"`
	UserID         string `xml:"urn:com.workday/bsvc User_ID"`
	Password       string `xml:"urn:com.workday/bsvc Password"`
	UseTempFile    *bool  `xml:"urn:com.workday/bsvc Use_Temp_File,omitempty"`
}

// FTP Transport Protocol Data element
type FTPTransportProtocolDataType struct {
	FTPAddress     string `xml:"urn:com.workday/bsvc FTP_Address"`
	Directory      string `xml:"urn:com.workday/bsvc Directory,omitempty"`
	UsePassiveMode *bool  `xml:"urn:com.workday/bsvc Use_Passive_Mode,omitempty"`
	UserID         string `xml:"urn:com.workday/bsvc User_ID"`
	Password       string `xml:"urn:com.workday/bsvc Password"`
	UseTempFile    *bool  `xml:"urn:com.workday/bsvc Use_Temp_File,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type FTPTransportProtocolObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type FTPTransportProtocolObjectType struct {
	ID         []FTPTransportProtocolObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Get_Event_Details request element
type GetEventDetailsRequestType struct {
	RequestReferences *EventRequestReferencesType `xml:"urn:com.workday/bsvc Request_References"`
	ResponseFilter    *ResponseFilterType         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get_Event_Details response element
type GetEventDetailsResponseType struct {
	RequestReferences *EventRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType        `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *EventResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Event Documents Request element
type GetEventDocumentsRequestType struct {
	RequestReferences *IntegrationEventRequestReferencesType `xml:"urn:com.workday/bsvc Request_References"`
	ResponseFilter    *ResponseFilterNoEntryMomentType       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Event Documents Response element
type GetEventDocumentsResponseType struct {
	RequestReferences *IntegrationEventRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *EventDocumentsResponseDataType        `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Incoming request element
type GetImportProcessMessagesRequestType struct {
	RequestCriteria *ImportProcessMessagesRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version         string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response
type GetImportProcessMessagesResponseType struct {
	RequestCriteria []ImportProcessMessagesRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  []ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults []ResponseResultsType                      `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    []ImportProcessMessageResponseDataType     `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Incoming element for Get Import Processes
type GetImportProcessesRequestType struct {
	RequestReferences *ImportProcessRequestReferencesType `xml:"urn:com.workday/bsvc Request_References"`
	ResponseFilter    *ResponseFilterType                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Import Processes Response element
type GetImportProcessesResponseType struct {
	RequestReferences *ImportProcessRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *ImportProcessResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for the Get Integration Events operation
type GetIntegrationEventsRequestType struct {
	RequestReferences *IntegrationEventRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *IntegrationEventRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterNoEntryMomentType       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Integration Events Response
type GetIntegrationEventsResponseType struct {
	RequestReferences *IntegrationEventRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *IntegrationEventRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *IntegrationEventResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Integration System Users Request element
type GetIntegrationSystemUsersRequestType struct {
	RequestReferences                *IntegrationSystemRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	IntegrationSystemRequestCriteria *IntegrationSystemRequestCriteriaType   `xml:"urn:com.workday/bsvc Integration_System_Request_Criteria,omitempty"`
	ResponseFilter                   *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version                          string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Root Response Element
type GetIntegrationSystemUsersResponseType struct {
	RequestReferences *IntegrationSystemRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *IntegrationSystemRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *IntegrationSystemUserResponseDataType  `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Criteria for filtering the Integration Systems that get returned in the response.
type GetIntegrationSystemsCriteriaType struct {
	IntegrationTemplateReference *IntegrationTemplateObjectType `xml:"urn:com.workday/bsvc Integration_Template_Reference,omitempty"`
	WorkdayAccountReference      *SystemUserObjectType          `xml:"urn:com.workday/bsvc Workday_Account_Reference,omitempty"`
	CloudCollectionReference     *CloudCollectionObjectType     `xml:"urn:com.workday/bsvc Cloud_Collection_Reference,omitempty"`
}

// Request element used to find and get Integration Systems and their associated data.
type GetIntegrationSystemsRequestType struct {
	RequestReferences *IntegrationSystemRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *GetIntegrationSystemsCriteriaType      `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *IntegrationSystemResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing instances of Integration System and each of its associated data.
type GetIntegrationSystemsResponseType struct {
	RequestReferences *IntegrationSystemRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *GetIntegrationSystemsCriteriaType      `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *IntegrationSystemResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *IntegrationSystemResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The "Request Criteria" element for the web service operation contains the filtering logic to limit the data that is returned in the response.
type GetReferencesRequestCriteriaType struct {
	ReferenceIDType            string `xml:"urn:com.workday/bsvc Reference_ID_Type"`
	IncludeDefaultedValuesOnly *bool  `xml:"urn:com.workday/bsvc Include_Defaulted_Values_Only,omitempty"`
}

// Root request element for the operation
type GetReferencesRequestType struct {
	RequestReferences *ReferencesRequestReferencesType  `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *GetReferencesRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The "Response Data" element contains the core data results based on the inbound request that was processed.
type GetReferencesResponseDataType struct {
	ReferenceID []ReferenceIDType `xml:"urn:com.workday/bsvc Reference_ID,omitempty"`
}

// Root element for the Response on the Get operation. Contains the instances returned by the Get operation and their accompanying data.
type GetReferencesResponseType struct {
	RequestCriteria *GetReferencesRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults *ResponseResultsType              `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *GetReferencesResponseDataType    `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Sequence Generator request element.
type GetSequenceGeneratorsRequestType struct {
	RequestReferences *SequenceGeneratorRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing instances of Sequence Generator and each of its associated data.
type GetSequenceGeneratorsResponseType struct {
	RequestReferences *SequenceGeneratorRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *SequenceGeneratorResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Subscriptions Request
type GetSubscriptionsRequestType struct {
	RequestReferences *SubscriptionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *SubscriptionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Subscriptions Response
type GetSubscriptionsResponseType struct {
	RequestReferences *SubscriptionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType               `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *SubscriptionResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// HTTP Header Data
type HTTPHeaderDataType struct {
	HTTPHeader []HTTPHeaderType `xml:"urn:com.workday/bsvc HTTP_Header"`
}

// HTTP Header
type HTTPHeaderType struct {
	HeaderValue string `xml:"urn:com.workday/bsvc Header_Value,omitempty"`
	HeaderName  string `xml:"urn:com.workday/bsvc Header_Name,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type HTTPSSLTransportProtocolObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type HTTPSSLTransportProtocolObjectType struct {
	ID         []HTTPSSLTransportProtocolObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// HTTP Transport Protocol Data element
type HTTPTransportProtocolDataType struct {
	HTTPAddress                       string                              `xml:"urn:com.workday/bsvc HTTP_Address"`
	HTTPDeliveryMethodReference       *UniqueIdentifierObjectType         `xml:"urn:com.workday/bsvc HTTP_Delivery_Method_Reference,omitempty"`
	WebServiceInvocationTypeReference *WebServiceInvocationTypeObjectType `xml:"urn:com.workday/bsvc Web_Service_Invocation_Type_Reference"`
	HTTPHeaderData                    *HTTPHeaderDataType                 `xml:"urn:com.workday/bsvc HTTP_Header_Data,omitempty"`
	UserID                            string                              `xml:"urn:com.workday/bsvc User_ID,omitempty"`
	Password                          string                              `xml:"urn:com.workday/bsvc Password,omitempty"`
}

// Import Process Data
type ImportProcessDataType struct {
	ProcessDescription       string                                    `xml:"urn:com.workday/bsvc Process_Description,omitempty"`
	PercentComplete          float64                                   `xml:"urn:com.workday/bsvc Percent_Complete,omitempty"`
	ProcessCompletedDateTime *time.Time                                `xml:"urn:com.workday/bsvc Process_Completed_DateTime,omitempty"`
	ImportHeaderReference    *InstanceObjectType                       `xml:"urn:com.workday/bsvc Import_Header_Reference,omitempty"`
	StatusReference          *BackgroundProcessRuntimeStatusObjectType `xml:"urn:com.workday/bsvc Status_Reference,omitempty"`
	HasMessages              *bool                                     `xml:"urn:com.workday/bsvc Has_Messages,omitempty"`
	ValidateOnly             *bool                                     `xml:"urn:com.workday/bsvc Validate_Only,omitempty"`
}

func (t *ImportProcessDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ImportProcessDataType
	var layout struct {
		*T
		ProcessCompletedDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Process_Completed_DateTime,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ProcessCompletedDateTime = (*xsdDateTime)(layout.T.ProcessCompletedDateTime)
	return e.EncodeElement(layout, start)
}
func (t *ImportProcessDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ImportProcessDataType
	var overlay struct {
		*T
		ProcessCompletedDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Process_Completed_DateTime,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ProcessCompletedDateTime = (*xsdDateTime)(overlay.T.ProcessCompletedDateTime)
	return d.DecodeElement(&overlay, &start)
}

// Response Data
type ImportProcessMessageResponseDataType struct {
	ImportProcessMessage []ImportProcessMessageType `xml:"urn:com.workday/bsvc Import_Process_Message,omitempty"`
}

// Message
type ImportProcessMessageType struct {
	ImportProcessMessageReference *UniqueIdentifierObjectType         `xml:"urn:com.workday/bsvc Import_Process_Message_Reference,omitempty"`
	ImportProcessMessageData      *BackgroundProcessMessageDataWSType `xml:"urn:com.workday/bsvc Import_Process_Message_Data,omitempty"`
}

// Request Criteria
type ImportProcessMessagesRequestCriteriaType struct {
	ImportProcessReference *WebServiceBackgroundProcessRuntimeObjectType `xml:"urn:com.workday/bsvc Import_Process_Reference"`
}

// Import Process Request References
type ImportProcessRequestReferencesType struct {
	ImportProcessReference *WebServiceBackgroundProcessRuntimeObjectType `xml:"urn:com.workday/bsvc Import_Process_Reference"`
}

// Import Process Response Data
type ImportProcessResponseDataType struct {
	ImportProcess []ImportProcessType `xml:"urn:com.workday/bsvc Import_Process,omitempty"`
}

// Import Process
type ImportProcessType struct {
	ImportProcessReference *WebServiceBackgroundProcessRuntimeObjectType `xml:"urn:com.workday/bsvc Import_Process_Reference,omitempty"`
	ImportProcessData      *ImportProcessDataType                        `xml:"urn:com.workday/bsvc Import_Process_Data,omitempty"`
}

// Request element for Increment of Sequence Generator.
type IncrementSequenceGeneratorRequestType struct {
	SequenceGeneratorReference *SequenceGeneratorObjectType `xml:"urn:com.workday/bsvc Sequence_Generator_Reference"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Increment of Sequence Generator.
type IncrementSequenceGeneratorResponseType struct {
	SequenceGeneratorReference *SequenceGeneratorObjectType `xml:"urn:com.workday/bsvc Sequence_Generator_Reference,omitempty"`
	SequencedValue             string                       `xml:"urn:com.workday/bsvc Sequenced_Value,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

type InstanceIDType struct {
	Value      string `xml:",chardata"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
	Type       string `xml:"urn:com.workday/bsvc type,attr,omitempty"`
}

type InstanceObjectType struct {
	ID         []InstanceIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegratableObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntegratableObjectType struct {
	ID         []IntegratableObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Container element for an Abstract Value.
type IntegrationAbstractValueDataType struct {
	Text              string               `xml:"urn:com.workday/bsvc Text,omitempty"`
	Date              *time.Time           `xml:"urn:com.workday/bsvc Date,omitempty"`
	DateTime          *time.Time           `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	Numeric           float64              `xml:"urn:com.workday/bsvc Numeric,omitempty"`
	Boolean           *bool                `xml:"urn:com.workday/bsvc Boolean,omitempty"`
	InstanceReference []InstanceObjectType `xml:"urn:com.workday/bsvc Instance_Reference,omitempty"`
}

func (t *IntegrationAbstractValueDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T IntegrationAbstractValueDataType
	var layout struct {
		*T
		Date     *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		DateTime *xsdDateTime `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Date = (*xsdDate)(layout.T.Date)
	layout.DateTime = (*xsdDateTime)(layout.T.DateTime)
	return e.EncodeElement(layout, start)
}
func (t *IntegrationAbstractValueDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T IntegrationAbstractValueDataType
	var overlay struct {
		*T
		Date     *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		DateTime *xsdDateTime `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Date = (*xsdDate)(overlay.T.Date)
	overlay.DateTime = (*xsdDateTime)(overlay.T.DateTime)
	return d.DecodeElement(&overlay, &start)
}

// Integration Attachment Configuration Data Element
type IntegrationAttachmentConfigurationDataType struct {
	AttachmentData []AttachmentDataWWSType `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
}

// Encapsulating element containing all Integration Attachment data.
type IntegrationAttachmentDataType struct {
	FileID               string              `xml:"urn:com.workday/bsvc File_ID,omitempty"`
	FileContent          *[]byte             `xml:"urn:com.workday/bsvc File_Content,omitempty"`
	ContentTypeReference *MimeTypeObjectType `xml:"urn:com.workday/bsvc Content_Type_Reference,omitempty"`
	Comments             string              `xml:"urn:com.workday/bsvc Comments,omitempty"`
	ContentType          string              `xml:"urn:com.workday/bsvc Content_Type,attr,omitempty"`
	Filename             string              `xml:"urn:com.workday/bsvc Filename,attr,omitempty"`
	Encoding             string              `xml:"urn:com.workday/bsvc Encoding,attr,omitempty"`
	Compressed           bool                `xml:"urn:com.workday/bsvc Compressed,attr,omitempty"`
}

func (t *IntegrationAttachmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T IntegrationAttachmentDataType
	var layout struct {
		*T
		FileContent *xsdBase64Binary `xml:"urn:com.workday/bsvc File_Content,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FileContent = (*xsdBase64Binary)(layout.T.FileContent)
	return e.EncodeElement(layout, start)
}
func (t *IntegrationAttachmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T IntegrationAttachmentDataType
	var overlay struct {
		*T
		FileContent *xsdBase64Binary `xml:"urn:com.workday/bsvc File_Content,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FileContent = (*xsdBase64Binary)(overlay.T.FileContent)
	return d.DecodeElement(&overlay, &start)
}

// Container element for definitions of custom Integration Attributes associated to this System
type IntegrationAttributeDataType struct {
	Order                                  string                                      `xml:"urn:com.workday/bsvc Order"`
	Name                                   string                                      `xml:"urn:com.workday/bsvc Name"`
	Description                            string                                      `xml:"urn:com.workday/bsvc Description,omitempty"`
	DataTypeExternalFieldReference         *ExternalFieldObjectType                    `xml:"urn:com.workday/bsvc Data_Type_External_Field_Reference"`
	DataTypeEnumerationDefinitionReference *IntegrationEnumerationDefinitionObjectType `xml:"urn:com.workday/bsvc Data_Type_Enumeration_Definition_Reference"`
	MultiSelect                            *bool                                       `xml:"urn:com.workday/bsvc Multi-Select,omitempty"`
	DisplayOptionReference                 []DisplayOptionObjectType                   `xml:"urn:com.workday/bsvc Display_Option_Reference,omitempty"`
}

// Contains Integration Attributes Name and Values associated to above Attributable instance.
type IntegrationAttributeIntegrationAttributeValueDataWWSType struct {
	IntegrationAttributeReference *ExternalInstanceObjectType            `xml:"urn:com.workday/bsvc Integration_Attribute_Reference,omitempty"`
	IntegrationAttributeName      string                                 `xml:"urn:com.workday/bsvc Integration_Attribute_Name,omitempty"`
	IntegrationAttributeValueData []IntegrationAttributeValueDataWWSType `xml:"urn:com.workday/bsvc Integration_Attribute_Value_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationAttributeObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type IntegrationAttributeObjectType struct {
	ID         []IntegrationAttributeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains Integration Attributes Value and OMS Environment restriction for above Attribute and Integration System.
type IntegrationAttributeValueDataWWSType struct {
	Text                 string                     `xml:"urn:com.workday/bsvc Text,omitempty"`
	Date                 *time.Time                 `xml:"urn:com.workday/bsvc Date,omitempty"`
	DateTime             *time.Time                 `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	Numeric              float64                    `xml:"urn:com.workday/bsvc Numeric,omitempty"`
	Boolean              *bool                      `xml:"urn:com.workday/bsvc Boolean,omitempty"`
	InstanceReference    []InstanceObjectType       `xml:"urn:com.workday/bsvc Instance_Reference,omitempty"`
	EnvironmentReference []OMSEnvironmentObjectType `xml:"urn:com.workday/bsvc Environment_Reference,omitempty"`
}

func (t *IntegrationAttributeValueDataWWSType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T IntegrationAttributeValueDataWWSType
	var layout struct {
		*T
		Date     *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		DateTime *xsdDateTime `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Date = (*xsdDate)(layout.T.Date)
	layout.DateTime = (*xsdDateTime)(layout.T.DateTime)
	return e.EncodeElement(layout, start)
}
func (t *IntegrationAttributeValueDataWWSType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T IntegrationAttributeValueDataWWSType
	var overlay struct {
		*T
		Date     *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		DateTime *xsdDateTime `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Date = (*xsdDate)(overlay.T.Date)
	overlay.DateTime = (*xsdDateTime)(overlay.T.DateTime)
	return d.DecodeElement(&overlay, &start)
}

// Integration Background Process Definition Data element
type IntegrationBackgroundProcessDefinitionDataType struct {
	ProcessDescription                       string                                     `xml:"urn:com.workday/bsvc Process_Description"`
	DynamicDescriptionExternalFieldReference *ExternalFieldObjectType                   `xml:"urn:com.workday/bsvc Dynamic_Description_External_Field_Reference"`
	IntegrationBackgroundProcessData         *IntegrationScheduledFutureProcessDataType `xml:"urn:com.workday/bsvc Integration_Background_Process_Data"`
}

// Connector Configuration
type IntegrationConnectorConfigurationDataType struct {
	LocalInEndpoint string `xml:"urn:com.workday/bsvc Local-In_Endpoint,omitempty"`
}

// Integration Custom Object Field Configuration Data
type IntegrationCustomObjectAliasConfigurationDataType struct {
	CustomObjectAliasReference  *IntegrationDocumentFieldObjectType `xml:"urn:com.workday/bsvc Custom_Object_Alias_Reference"`
	CustomObjectReference       *CustomObjectObjectType             `xml:"urn:com.workday/bsvc Custom_Object_Reference,omitempty"`
	CaptureMultipleValues       *bool                               `xml:"urn:com.workday/bsvc Capture_Multiple_Values,omitempty"`
	CustomObjectRESTURL         string                              `xml:"urn:com.workday/bsvc Custom_Object_REST_URL,omitempty"`
	MultipleCustomObjectRESTURL string                              `xml:"urn:com.workday/bsvc Multiple_Custom_Object_REST_URL,omitempty"`
}

// Integration Custom Object Configuration Data
type IntegrationCustomObjectConfigurationDataType struct {
	IntegrationCustomObjectAliasConfigurationData []IntegrationCustomObjectAliasConfigurationDataType `xml:"urn:com.workday/bsvc Integration_Custom_Object_Alias_Configuration_Data"`
}

// Container element for data related to a Data Change Service and its use within an Integration System.
type IntegrationDataChangeConfigurationDataType struct {
	ExternalFieldReference []ExternalFieldObjectType `xml:"urn:com.workday/bsvc External_Field_Reference,omitempty"`
}

// Integration Data Initialization Configuration Data
type IntegrationDataInitializationConfigurationDataType struct {
	IntegrationFieldAttributesFieldConfigurationData []IntegrationFieldAttributesFieldConfigurationDataType `xml:"urn:com.workday/bsvc Integration_Field_Attributes_Field_Configuration_Data,omitempty"`
	PopulationEligibilityData                        *PopulationEligibilityDataType                         `xml:"urn:com.workday/bsvc Population_Eligibility_Data,omitempty"`
}

// Integration Data Initialization Override Configuration Data
type IntegrationDataInitializationOverrideConfigurationDataType struct {
	ParameterData                                    []IntegrationParameterAssignmentDataType     `xml:"urn:com.workday/bsvc Parameter_Data,omitempty"`
	IntegrationDocumentFieldforFieldOverrideData     []OverriddenIntegrationDocumentFieldDataType `xml:"urn:com.workday/bsvc Integration_Document_Field_for_Field_Override_Data,omitempty"`
	IntegrationDocumentFieldforDocumentExtensionData []ExtendedIntegrationDocumentFieldDataType   `xml:"urn:com.workday/bsvc Integration_Document_Field_for_Document_Extension_Data,omitempty"`
}

// Integration Delivery Configuration Data element
type IntegrationDeliveryConfigurationDataType struct {
}

// Contains a unique identifier for an instance of an object.
type IntegrationDocumentFieldObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type IntegrationDocumentFieldObjectType struct {
	ID         []IntegrationDocumentFieldObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Integration Document Field Options
type IntegrationDocumentFieldOptionsType struct {
	IntegrationDocumentOptionValueReference []IntegrationDocumentOptionValueObjectType `xml:"urn:com.workday/bsvc Integration_Document_Option_Value_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationDocumentOptionValueObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type IntegrationDocumentOptionValueObjectType struct {
	ID         []IntegrationDocumentOptionValueObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Integration Document Stacks for Integration Document Field Configuration Data
type IntegrationDocumentStackDataType struct {
	StackLevel                        float64                             `xml:"urn:com.workday/bsvc Stack_Level"`
	IntegrationDocumentFieldReference *IntegrationDocumentFieldObjectType `xml:"urn:com.workday/bsvc Integration_Document_Field_Reference"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationESBInvocationAbstractObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntegrationESBInvocationAbstractObjectType struct {
	ID         []IntegrationESBInvocationAbstractObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing records counts for the Integration Enterprise Event
type IntegrationEnterpriseEventRecordsDataType struct {
	TotalRecords         float64 `xml:"urn:com.workday/bsvc Total_Records,omitempty"`
	TotalRecordsUploaded float64 `xml:"urn:com.workday/bsvc Total_Records_Uploaded,omitempty"`
	TotalFailedRecords   float64 `xml:"urn:com.workday/bsvc Total_Failed_Records,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationEnumerationDefinitionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntegrationEnumerationDefinitionObjectType struct {
	ID         []IntegrationEnumerationDefinitionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationEnumerationObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type IntegrationEnumerationObjectType struct {
	ID         []IntegrationEnumerationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationEventAbstractObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntegrationEventAbstractObjectType struct {
	ID         []IntegrationEventAbstractObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element containing all Integration Event data.
type IntegrationEventDataWWSType struct {
	IntegrationSystemReference              *IntegrationSystemObjectType                  `xml:"urn:com.workday/bsvc Integration_System_Reference,omitempty"`
	ProcessDescription                      string                                        `xml:"urn:com.workday/bsvc Process_Description,omitempty"`
	InitiatedDateTime                       *time.Time                                    `xml:"urn:com.workday/bsvc Initiated_DateTime,omitempty"`
	IntegrationResponseMessage              string                                        `xml:"urn:com.workday/bsvc Integration_Response_Message,omitempty"`
	CompletedDateTime                       *time.Time                                    `xml:"urn:com.workday/bsvc Completed_DateTime,omitempty"`
	IntegrationEventMemberReference         []IntegratableObjectType                      `xml:"urn:com.workday/bsvc Integration_Event_Member_Reference,omitempty"`
	InitiatedBySystemAccountReference       *SystemUserObjectType                         `xml:"urn:com.workday/bsvc Initiated_By_System_Account_Reference,omitempty"`
	PercentComplete                         float64                                       `xml:"urn:com.workday/bsvc Percent_Complete,omitempty"`
	IntegrationRuntimeParameterData         []IntegrationRuntimeParameterDataType         `xml:"urn:com.workday/bsvc Integration_Runtime_Parameter_Data,omitempty"`
	IntegrationServiceGeneratedSequenceData []IntegrationServiceGeneratedSequenceDataType `xml:"urn:com.workday/bsvc Integration_Service_Generated_Sequence_Data,omitempty"`
	IntegrationEnterpriseEventRecordsData   []IntegrationEnterpriseEventRecordsDataType   `xml:"urn:com.workday/bsvc Integration_Enterprise_Event_Records_Data,omitempty"`
}

func (t *IntegrationEventDataWWSType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T IntegrationEventDataWWSType
	var layout struct {
		*T
		InitiatedDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Initiated_DateTime,omitempty"`
		CompletedDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Completed_DateTime,omitempty"`
	}
	layout.T = (*T)(t)
	layout.InitiatedDateTime = (*xsdDateTime)(layout.T.InitiatedDateTime)
	layout.CompletedDateTime = (*xsdDateTime)(layout.T.CompletedDateTime)
	return e.EncodeElement(layout, start)
}
func (t *IntegrationEventDataWWSType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T IntegrationEventDataWWSType
	var overlay struct {
		*T
		InitiatedDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Initiated_DateTime,omitempty"`
		CompletedDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Completed_DateTime,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.InitiatedDateTime = (*xsdDateTime)(overlay.T.InitiatedDateTime)
	overlay.CompletedDateTime = (*xsdDateTime)(overlay.T.CompletedDateTime)
	return d.DecodeElement(&overlay, &start)
}

// Element for derived information regarding a Background Process Instance
type IntegrationEventReportingDetailsDataType struct {
	BackgroundProcessInstanceStatusReference *BackgroundProcessRuntimeStatusObjectType `xml:"urn:com.workday/bsvc Background_Process_Instance_Status_Reference,omitempty"`
	ParentEventReference                     *EventObjectType                          `xml:"urn:com.workday/bsvc Parent_Event_Reference,omitempty"`
	BackgroundProcessMessageData             []IntegrationMessageDetailDataType        `xml:"urn:com.workday/bsvc Background_Process_Message_Data,omitempty"`
	OutputDocumentData                       []IntegrationRepositoryDocumentType       `xml:"urn:com.workday/bsvc Output_Document_Data,omitempty"`
	ConsolidatedReportReference              *RepositoryDocumentObjectType             `xml:"urn:com.workday/bsvc Consolidated_Report_Reference,omitempty"`
	LogFileReference                         []RepositoryDocumentObjectType            `xml:"urn:com.workday/bsvc Log_File_Reference,omitempty"`
}

// Criteria for filtering the Integration Events that get returned in the response.
type IntegrationEventRequestCriteriaType struct {
	IntegrationSystemReference      *IntegrationSystemObjectType               `xml:"urn:com.workday/bsvc Integration_System_Reference,omitempty"`
	IntegrationEventStatusReference []BackgroundProcessRuntimeStatusObjectType `xml:"urn:com.workday/bsvc Integration_Event_Status_Reference,omitempty"`
	SentAfter                       *time.Time                                 `xml:"urn:com.workday/bsvc Sent_After,omitempty"`
	SentBefore                      *time.Time                                 `xml:"urn:com.workday/bsvc Sent_Before,omitempty"`
}

func (t *IntegrationEventRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T IntegrationEventRequestCriteriaType
	var layout struct {
		*T
		SentAfter  *xsdDateTime `xml:"urn:com.workday/bsvc Sent_After,omitempty"`
		SentBefore *xsdDateTime `xml:"urn:com.workday/bsvc Sent_Before,omitempty"`
	}
	layout.T = (*T)(t)
	layout.SentAfter = (*xsdDateTime)(layout.T.SentAfter)
	layout.SentBefore = (*xsdDateTime)(layout.T.SentBefore)
	return e.EncodeElement(layout, start)
}
func (t *IntegrationEventRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T IntegrationEventRequestCriteriaType
	var overlay struct {
		*T
		SentAfter  *xsdDateTime `xml:"urn:com.workday/bsvc Sent_After,omitempty"`
		SentBefore *xsdDateTime `xml:"urn:com.workday/bsvc Sent_Before,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.SentAfter = (*xsdDateTime)(overlay.T.SentAfter)
	overlay.SentBefore = (*xsdDateTime)(overlay.T.SentBefore)
	return d.DecodeElement(&overlay, &start)
}

// Integration Event Request References
type IntegrationEventRequestReferencesType struct {
	IntegrationEventReference []IntegrationESBInvocationAbstractObjectType `xml:"urn:com.workday/bsvc Integration_Event_Reference"`
}

// Integration Event Response Data
type IntegrationEventResponseDataType struct {
	IntegrationEvent []IntegrationEventType `xml:"urn:com.workday/bsvc Integration_Event,omitempty"`
}

// Integration Event
type IntegrationEventType struct {
	IntegrationEventReference     *IntegrationESBInvocationAbstractObjectType `xml:"urn:com.workday/bsvc Integration_Event_Reference,omitempty"`
	IntegrationEventData          *IntegrationEventDataWWSType                `xml:"urn:com.workday/bsvc Integration_Event_Data,omitempty"`
	BackgroundProcessInstanceData *IntegrationEventReportingDetailsDataType   `xml:"urn:com.workday/bsvc Background_Process_Instance_Data,omitempty"`
}

// Integration Field Attributes Configuration Data
type IntegrationFieldAttributesConfigurationDataType struct {
	IntegrationFieldAttributesFieldConfigurationData []IntegrationFieldAttributesFieldConfigurationDataType `xml:"urn:com.workday/bsvc Integration_Field_Attributes_Field_Configuration_Data"`
}

// Integration Field Attributes Field Configuration Data
type IntegrationFieldAttributesFieldConfigurationDataType struct {
	FieldReference                                   *IntegrationDocumentFieldObjectType                    `xml:"urn:com.workday/bsvc Field_Reference"`
	WebServiceAlias                                  string                                                 `xml:"urn:com.workday/bsvc Web_Service_Alias,omitempty"`
	IncludeinOutput                                  *bool                                                  `xml:"urn:com.workday/bsvc Include_in_Output,omitempty"`
	RequiredField                                    *bool                                                  `xml:"urn:com.workday/bsvc Required_Field,omitempty"`
	MaximumLength                                    float64                                                `xml:"urn:com.workday/bsvc Maximum_Length,omitempty"`
	IntegrationDocumentFieldOptions                  []IntegrationDocumentFieldOptionsType                  `xml:"urn:com.workday/bsvc Integration_Document_Field_Options,omitempty"`
	IntegrationFieldAttributesFieldConfigurationData []IntegrationFieldAttributesFieldConfigurationDataType `xml:"urn:com.workday/bsvc Integration_Field_Attributes_Field_Configuration_Data,omitempty"`
}

// Integration Field Override Configuration Data
type IntegrationFieldOverrideConfigurationDataType struct {
	IntegrationBusinessObjectReference                *UniqueIdentifierObjectType                          `xml:"urn:com.workday/bsvc Integration_Business_Object_Reference,omitempty"`
	IntegrationDocumentFieldOverrideConfigurationData []IntegrationFieldOverrideFieldConfigurationDataType `xml:"urn:com.workday/bsvc Integration_Document_Field_Override_Configuration_Data"`
}

// Integration Document Field Override Configuration Data
type IntegrationFieldOverrideFieldConfigurationDataType struct {
	FieldReference                  *IntegrationDocumentFieldObjectType                       `xml:"urn:com.workday/bsvc Field_Reference"`
	ExternalFieldReference          *ExternalFieldObjectType                                  `xml:"urn:com.workday/bsvc External_Field_Reference,omitempty"`
	ExternalParameterAssignmentData []IntegrationFieldOverrideParameterInitializationDataType `xml:"urn:com.workday/bsvc External_Parameter_Assignment_Data,omitempty"`
	RequiredField                   *bool                                                     `xml:"urn:com.workday/bsvc Required_Field,omitempty"`
	MaximumLength                   float64                                                   `xml:"urn:com.workday/bsvc Maximum_Length,omitempty"`
}

// For each parameter, assignment details.
type IntegrationFieldOverrideParameterInitializationDataType struct {
	ExternalParameterReference  *AbstractExternalParameterObjectType `xml:"urn:com.workday/bsvc External_Parameter_Reference"`
	ParameterInitializationData *ParameterInitializationWWSDataType  `xml:"urn:com.workday/bsvc Parameter_Initialization_Data"`
}

// File Utility Data element
type IntegrationFileUtilityDataType struct {
	ID                            string                             `xml:"urn:com.workday/bsvc ID,omitempty"`
	Filename                      string                             `xml:"urn:com.workday/bsvc Filename,omitempty"`
	SequenceGeneratorData         *AbstractSequenceGeneratorDataType `xml:"urn:com.workday/bsvc Sequence_Generator_Data,omitempty"`
	DynamicFilenameDefinitionData *DynamicFilenameDefinitionDataType `xml:"urn:com.workday/bsvc Dynamic_Filename_Definition_Data,omitempty"`
	MimeTypeReference             *MimeTypeObjectType                `xml:"urn:com.workday/bsvc Mime_Type_Reference,omitempty"`
	Compressed                    *bool                              `xml:"urn:com.workday/bsvc Compressed,omitempty"`
	DocumentRetentionPolicy       float64                            `xml:"urn:com.workday/bsvc Document_Retention_Policy,omitempty"`
	PGPEncryptionSettingsData     *PGPEncryptionSettingsDataType     `xml:"urn:com.workday/bsvc PGP_Encryption_Settings_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationLaunchOptionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntegrationLaunchOptionObjectType struct {
	ID         []IntegrationLaunchOptionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Integration Launch Parameter Data
type IntegrationLaunchParameterDataType struct {
	LaunchParameterReference *LaunchParameterObjectType        `xml:"urn:com.workday/bsvc Launch_Parameter_Reference"`
	LaunchParameterValueData *IntegrationAbstractValueDataType `xml:"urn:com.workday/bsvc Launch_Parameter_Value_Data,omitempty"`
}

// Container element for definitions of custom Integration Maps associated to this System
type IntegrationMapDataType struct {
	Order                                                string                                      `xml:"urn:com.workday/bsvc Order"`
	Name                                                 string                                      `xml:"urn:com.workday/bsvc Name"`
	Description                                          string                                      `xml:"urn:com.workday/bsvc Description,omitempty"`
	InternalValueDataTypeExternalFieldReference          *ExternalFieldObjectType                    `xml:"urn:com.workday/bsvc Internal_Value_Data_Type_External_Field_Reference"`
	InternalValueDataTypeIntegrationEnumerationReference *IntegrationEnumerationDefinitionObjectType `xml:"urn:com.workday/bsvc Internal_Value_Data_Type_Integration_Enumeration_Reference"`
	InternalValueDisplayOptionReference                  []DisplayOptionObjectType                   `xml:"urn:com.workday/bsvc Internal_Value_Display_Option_Reference,omitempty"`
	ExternalValueDataTypeExternalFieldReference          *ExternalFieldObjectType                    `xml:"urn:com.workday/bsvc External_Value_Data_Type_External_Field_Reference"`
	ExternalValueDataTypeIntegrationEnumerationReference *IntegrationEnumerationDefinitionObjectType `xml:"urn:com.workday/bsvc External_Value_Data_Type_Integration_Enumeration_Reference"`
	ExternalValueDisplayOptionReference                  []DisplayOptionObjectType                   `xml:"urn:com.workday/bsvc External_Value_Display_Option_Reference,omitempty"`
}

// Container element for Integration Map value.
type IntegrationMapIntegrationMapValueDataWWSType struct {
	IntegrationMapReference *ExternalInstanceObjectType       `xml:"urn:com.workday/bsvc Integration_Map_Reference,omitempty"`
	IntegrationMapName      string                            `xml:"urn:com.workday/bsvc Integration_Map_Name,omitempty"`
	DefaultValueData        *IntegrationAbstractValueDataType `xml:"urn:com.workday/bsvc Default_Value_Data,omitempty"`
	IntegrationMapValueData []IntegrationMapValueDataWWSType  `xml:"urn:com.workday/bsvc Integration_Map_Value_Data,omitempty"`
}

// Container element for each row of data within an Integration Map value.
type IntegrationMapValueDataWWSType struct {
	InternalValueData *IntegrationAbstractValueDataType       `xml:"urn:com.workday/bsvc Internal_Value_Data,omitempty"`
	ExternalValueData *CopyofIntegrationAbstractValueDataType `xml:"urn:com.workday/bsvc External_Value_Data,omitempty"`
}

// Encapsulating element containing all Integration Message data.
type IntegrationMessageDataType struct {
	IntegrationEventReference  *IntegrationESBInvocationAbstractObjectType      `xml:"urn:com.workday/bsvc Integration_Event_Reference"`
	IntegrationSystemReference *IntegrationSystemAuditedObjectType              `xml:"urn:com.workday/bsvc Integration_System_Reference"`
	MessageSummary             string                                           `xml:"urn:com.workday/bsvc Message_Summary"`
	MessageDetail              string                                           `xml:"urn:com.workday/bsvc Message_Detail,omitempty"`
	SeverityLevelReference     *BackgroundProcessMessageSeverityLevelObjectType `xml:"urn:com.workday/bsvc Severity_Level_Reference"`
	MessageTargetReference     []InstanceObjectType                             `xml:"urn:com.workday/bsvc Message_Target_Reference,omitempty"`
	IntegrationAttachmentData  []IntegrationAttachmentDataType                  `xml:"urn:com.workday/bsvc Integration_Attachment_Data,omitempty"`
	RepositoryDocumentData     []IntegrationRepositoryDocumentDataType          `xml:"urn:com.workday/bsvc Repository_Document_Data,omitempty"`
	EnqueueNotificationMessage *bool                                            `xml:"urn:com.workday/bsvc Enqueue_Notification_Message,omitempty"`
}

// Element for Messages related to Background Process Instances
type IntegrationMessageDetailDataType struct {
	Timestamp              *time.Time                                       `xml:"urn:com.workday/bsvc Timestamp,omitempty"`
	SeverityLevelReference *BackgroundProcessMessageSeverityLevelObjectType `xml:"urn:com.workday/bsvc Severity_Level_Reference,omitempty"`
	MessageSummary         string                                           `xml:"urn:com.workday/bsvc Message_Summary,omitempty"`
	MessageDetail          string                                           `xml:"urn:com.workday/bsvc Message_Detail,omitempty"`
	MessageTargetReference []InstanceObjectType                             `xml:"urn:com.workday/bsvc Message_Target_Reference,omitempty"`
}

func (t *IntegrationMessageDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T IntegrationMessageDetailDataType
	var layout struct {
		*T
		Timestamp *xsdDateTime `xml:"urn:com.workday/bsvc Timestamp,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Timestamp = (*xsdDateTime)(layout.T.Timestamp)
	return e.EncodeElement(layout, start)
}
func (t *IntegrationMessageDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T IntegrationMessageDetailDataType
	var overlay struct {
		*T
		Timestamp *xsdDateTime `xml:"urn:com.workday/bsvc Timestamp,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Timestamp = (*xsdDateTime)(overlay.T.Timestamp)
	return d.DecodeElement(&overlay, &start)
}

// Element containing configuration of Notifications.
type IntegrationNotificationDataType struct {
	TriggeronLaunch               bool                                                  `xml:"urn:com.workday/bsvc Trigger_on_Launch"`
	TriggeronStatusReference      []BackgroundProcessRuntimeStatusObjectType            `xml:"urn:com.workday/bsvc Trigger_on_Status_Reference"`
	NotificationConditionData     []IntegrationNotificationIntegrationConditionDataType `xml:"urn:com.workday/bsvc Notification_Condition_Data,omitempty"`
	LoopsOnExternalFieldReference *ExternalFieldObjectType                              `xml:"urn:com.workday/bsvc Loops_On_External_Field_Reference,omitempty"`
	NotificationNotifiesData      *NotificationNotifiesDataType                         `xml:"urn:com.workday/bsvc Notification_Notifies_Data,omitempty"`
	SecurityGroupReference        []SecurityGroupObjectType                             `xml:"urn:com.workday/bsvc Security_Group_Reference,omitempty"`
	EmailAddressData              []InternetEmailAddressDataforNotificationsType        `xml:"urn:com.workday/bsvc Email_Address_Data,omitempty"`
	NotificationSubjectData       *NotificationSubjectDataType                          `xml:"urn:com.workday/bsvc Notification_Subject_Data,omitempty"`
	NotificationBodyData          *NotificationBodyDataType                             `xml:"urn:com.workday/bsvc Notification_Body_Data,omitempty"`
	NotificationAttachmentData    *NotificationAttachmentDataType                       `xml:"urn:com.workday/bsvc Notification_Attachment_Data,omitempty"`
}

// Element containing the conditions to be evaluated for a Notification
type IntegrationNotificationIntegrationConditionDataType struct {
	ConditionRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Condition_Rule_Reference"`
}

// Parameter Data
type IntegrationParameterAssignmentDataType struct {
	IntegrationParameterName                 string                                     `xml:"urn:com.workday/bsvc Integration_Parameter_Name"`
	DataTypeExternalFieldReference           *ExternalFieldObjectType                   `xml:"urn:com.workday/bsvc Data_Type_External_Field_Reference"`
	Boolean                                  *bool                                      `xml:"urn:com.workday/bsvc Boolean,omitempty"`
	Currency                                 float64                                    `xml:"urn:com.workday/bsvc Currency,omitempty"`
	Date                                     *time.Time                                 `xml:"urn:com.workday/bsvc Date,omitempty"`
	DateTime                                 *time.Time                                 `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	Time                                     *time.Time                                 `xml:"urn:com.workday/bsvc Time,omitempty"`
	Numeric                                  float64                                    `xml:"urn:com.workday/bsvc Numeric,omitempty"`
	Text                                     string                                     `xml:"urn:com.workday/bsvc Text,omitempty"`
	BusinessObjectInstanceReference          []InstanceObjectType                       `xml:"urn:com.workday/bsvc Business_Object_Instance_Reference,omitempty"`
	ExternalFieldReference                   *ExternalFieldObjectType                   `xml:"urn:com.workday/bsvc External_Field_Reference,omitempty"`
	IntegrationAttributeReference            *IntegrationAttributeObjectType            `xml:"urn:com.workday/bsvc Integration_Attribute_Reference,omitempty"`
	LaunchParameterReference                 *LaunchParameterObjectType                 `xml:"urn:com.workday/bsvc Launch_Parameter_Reference,omitempty"`
	DatefromDateTimeZone                     *time.Time                                 `xml:"urn:com.workday/bsvc Date_from_DateTimeZone,omitempty"`
	TimefromDateTimeZone                     *time.Time                                 `xml:"urn:com.workday/bsvc Time_from_DateTimeZone,omitempty"`
	TimeZonefromDateTimeZoneReference        *TimeZoneObjectType                        `xml:"urn:com.workday/bsvc TimeZone_from_DateTimeZone_Reference,omitempty"`
	ParameterInitializationOperatorReference *ParameterInitializationOperatorObjectType `xml:"urn:com.workday/bsvc Parameter_Initialization_Operator_Reference"`
	CurrencyReference                        *CurrencyObjectType                        `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	IntegrationDocumentStackData             []IntegrationDocumentStackDataType         `xml:"urn:com.workday/bsvc Integration_Document_Stack_Data,omitempty"`
	Global                                   bool                                       `xml:"urn:com.workday/bsvc Global,attr,omitempty"`
}

func (t *IntegrationParameterAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T IntegrationParameterAssignmentDataType
	var layout struct {
		*T
		Date                 *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		DateTime             *xsdDateTime `xml:"urn:com.workday/bsvc DateTime,omitempty"`
		Time                 *xsdTime     `xml:"urn:com.workday/bsvc Time,omitempty"`
		DatefromDateTimeZone *xsdDate     `xml:"urn:com.workday/bsvc Date_from_DateTimeZone,omitempty"`
		TimefromDateTimeZone *xsdTime     `xml:"urn:com.workday/bsvc Time_from_DateTimeZone,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Date = (*xsdDate)(layout.T.Date)
	layout.DateTime = (*xsdDateTime)(layout.T.DateTime)
	layout.Time = (*xsdTime)(layout.T.Time)
	layout.DatefromDateTimeZone = (*xsdDate)(layout.T.DatefromDateTimeZone)
	layout.TimefromDateTimeZone = (*xsdTime)(layout.T.TimefromDateTimeZone)
	return e.EncodeElement(layout, start)
}
func (t *IntegrationParameterAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T IntegrationParameterAssignmentDataType
	var overlay struct {
		*T
		Date                 *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		DateTime             *xsdDateTime `xml:"urn:com.workday/bsvc DateTime,omitempty"`
		Time                 *xsdTime     `xml:"urn:com.workday/bsvc Time,omitempty"`
		DatefromDateTimeZone *xsdDate     `xml:"urn:com.workday/bsvc Date_from_DateTimeZone,omitempty"`
		TimefromDateTimeZone *xsdTime     `xml:"urn:com.workday/bsvc Time_from_DateTimeZone,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Date = (*xsdDate)(overlay.T.Date)
	overlay.DateTime = (*xsdDateTime)(overlay.T.DateTime)
	overlay.Time = (*xsdTime)(overlay.T.Time)
	overlay.DatefromDateTimeZone = (*xsdDate)(overlay.T.DatefromDateTimeZone)
	overlay.TimefromDateTimeZone = (*xsdTime)(overlay.T.TimefromDateTimeZone)
	return d.DecodeElement(&overlay, &start)
}

// Integration Parameter Initialization Data element
type IntegrationParameterInitializationDataType struct {
	LaunchParameterReference    *LaunchParameterObjectType          `xml:"urn:com.workday/bsvc Launch_Parameter_Reference"`
	ParameterInitializationData *ParameterInitializationWWSDataType `xml:"urn:com.workday/bsvc Parameter_Initialization_Data"`
}

// Integration Parameter Reference
type IntegrationParameterReferenceType struct {
	Global                       *bool                              `xml:"urn:com.workday/bsvc Global,omitempty"`
	IntegrationParameterName     string                             `xml:"urn:com.workday/bsvc Integration_Parameter_Name"`
	IntegrationDocumentStackData []IntegrationDocumentStackDataType `xml:"urn:com.workday/bsvc Integration_Document_Stack_Data,omitempty"`
}

// Integration Payload Data element
type IntegrationPayloadDataType struct {
	PGPEncryptionSettingsData *PGPEncryptionSettingsDataType `xml:"urn:com.workday/bsvc PGP_Encryption_Settings_Data,omitempty"`
	Compressed                *bool                          `xml:"urn:com.workday/bsvc Compressed,omitempty"`
	UseImprovedCompression    *bool                          `xml:"urn:com.workday/bsvc Use_Improved_Compression,omitempty"`
}

// Integration Report Field Configuration Data
type IntegrationReportFieldConfigurationDataType struct {
	ReportAliasReference            *IntegrationDocumentFieldObjectType `xml:"urn:com.workday/bsvc Report_Alias_Reference"`
	CustomReportDefinitionReference *CustomReportDefinitionObjectType   `xml:"urn:com.workday/bsvc Custom_Report_Definition_Reference,omitempty"`
	RESTEndpoint                    string                              `xml:"urn:com.workday/bsvc REST_Endpoint,omitempty"`
}

// Integration Report Parameter Initialization Data element
type IntegrationReportParameterInitializationDataType struct {
	SimpleWorkDataReference     *UniqueIdentifierObjectType         `xml:"urn:com.workday/bsvc Simple_Work_Data_Reference"`
	XMLName                     string                              `xml:"urn:com.workday/bsvc XML_Name,omitempty"`
	ParameterInitializationData *ParameterInitializationWWSDataType `xml:"urn:com.workday/bsvc Parameter_Initialization_Data"`
}

// Integration Report Service Configuration Data
type IntegrationReportServiceConfigurationDataType struct {
	IntegrationReportFieldConfigurationData []IntegrationReportFieldConfigurationDataType `xml:"urn:com.workday/bsvc Integration_Report_Field_Configuration_Data"`
}

// Encapsulating element containing all Integration Repository Document data.
type IntegrationRepositoryDocumentDataType struct {
	FileSize                 float64                  `xml:"urn:com.workday/bsvc File_Size,omitempty"`
	ContentTypeReference     *MimeTypeObjectType      `xml:"urn:com.workday/bsvc Content_Type_Reference,omitempty"`
	ContentType              string                   `xml:"urn:com.workday/bsvc Content_Type,omitempty"`
	DocumentTypeReference    *DocumentTypeObjectType  `xml:"urn:com.workday/bsvc Document_Type_Reference,omitempty"`
	ExpirationTimestamp      time.Time                `xml:"urn:com.workday/bsvc Expiration_Timestamp"`
	OwnerReference           *SystemAccountObjectType `xml:"urn:com.workday/bsvc Owner_Reference,omitempty"`
	DocumentTagReference     []DocumentTagObjectType  `xml:"urn:com.workday/bsvc Document_Tag_Reference,omitempty"`
	CustomDocumentTag        []string                 `xml:"urn:com.workday/bsvc Custom_Document_Tag,omitempty"`
	SecuredInstanceReference []InstanceObjectType     `xml:"urn:com.workday/bsvc Secured_Instance_Reference,omitempty"`
	DocumentID               string                   `xml:"urn:com.workday/bsvc Document_ID,attr,omitempty"`
	FileName                 string                   `xml:"urn:com.workday/bsvc File_Name,attr,omitempty"`
}

func (t *IntegrationRepositoryDocumentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T IntegrationRepositoryDocumentDataType
	var layout struct {
		*T
		ExpirationTimestamp *xsdDateTime `xml:"urn:com.workday/bsvc Expiration_Timestamp"`
	}
	layout.T = (*T)(t)
	layout.ExpirationTimestamp = (*xsdDateTime)(&layout.T.ExpirationTimestamp)
	return e.EncodeElement(layout, start)
}
func (t *IntegrationRepositoryDocumentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T IntegrationRepositoryDocumentDataType
	var overlay struct {
		*T
		ExpirationTimestamp *xsdDateTime `xml:"urn:com.workday/bsvc Expiration_Timestamp"`
	}
	overlay.T = (*T)(t)
	overlay.ExpirationTimestamp = (*xsdDateTime)(&overlay.T.ExpirationTimestamp)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type IntegrationRepositoryDocumentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntegrationRepositoryDocumentObjectType struct {
	ID         []IntegrationRepositoryDocumentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Integration Repository Document
type IntegrationRepositoryDocumentType struct {
	IntegrationRepositoryDocumentReference *IntegrationRepositoryDocumentObjectType `xml:"urn:com.workday/bsvc Integration_Repository_Document_Reference,omitempty"`
	IntegrationRepositoryDocumentData      *IntegrationRepositoryDocumentDataType   `xml:"urn:com.workday/bsvc Integration_Repository_Document_Data,omitempty"`
}

// Integration Retrieval Configuration Data element
type IntegrationRetrievalConfigurationDataType struct {
}

// Integration Runtime Parameter Data
type IntegrationRuntimeParameterDataType struct {
	LaunchConfigurableName string                                   `xml:"urn:com.workday/bsvc Launch_Configurable_Name,omitempty"`
	ParameterName          []SimpleWorkDataRuntimeParameterNameType `xml:"urn:com.workday/bsvc Parameter_Name,omitempty"`
	Text                   []TextAttributeType                      `xml:"urn:com.workday/bsvc Text,omitempty"`
	Date                   *time.Time                               `xml:"urn:com.workday/bsvc Date,omitempty"`
	Boolean                *bool                                    `xml:"urn:com.workday/bsvc Boolean,omitempty"`
	InstanceSetReference   []InstanceObjectType                     `xml:"urn:com.workday/bsvc Instance_Set_Reference,omitempty"`
}

func (t *IntegrationRuntimeParameterDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T IntegrationRuntimeParameterDataType
	var layout struct {
		*T
		Date *xsdDateTime `xml:"urn:com.workday/bsvc Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Date = (*xsdDateTime)(layout.T.Date)
	return e.EncodeElement(layout, start)
}
func (t *IntegrationRuntimeParameterDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T IntegrationRuntimeParameterDataType
	var overlay struct {
		*T
		Date *xsdDateTime `xml:"urn:com.workday/bsvc Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Date = (*xsdDateTime)(overlay.T.Date)
	return d.DecodeElement(&overlay, &start)
}

// Integration Scheduled Future Process Data element
type IntegrationScheduledFutureProcessDataType struct {
	IntegrationSystemReference                   *IntegrationSystemIdentifierObjectType             `xml:"urn:com.workday/bsvc Integration_System_Reference"`
	IntegrationParameterInitializationData       []IntegrationParameterInitializationDataType       `xml:"urn:com.workday/bsvc Integration_Parameter_Initialization_Data,omitempty"`
	IntegrationReportParameterInitializationData []IntegrationReportParameterInitializationDataType `xml:"urn:com.workday/bsvc Integration_Report_Parameter_Initialization_Data,omitempty"`
	ConcurrencyConfigurationData                 *ConcurrencyConfigurationDataType                  `xml:"urn:com.workday/bsvc Concurrency_Configuration_Data,omitempty"`
}

// Container element for data related to a Sequence Generator Service and its use within an Integration System.
type IntegrationSequenceGeneratorConfigurationDataType struct {
	IntegrationSequencerReference *IntegrationSequencerObjectType `xml:"urn:com.workday/bsvc Integration_Sequencer_Reference"`
	SequenceGeneratorReference    *SequenceGeneratorObjectType    `xml:"urn:com.workday/bsvc Sequence_Generator_Reference"`
	OMSEnvironmentReference       []OMSEnvironmentObjectType      `xml:"urn:com.workday/bsvc OMS_Environment_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationSequenceGeneratorServiceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntegrationSequenceGeneratorServiceObjectType struct {
	ID         []IntegrationSequenceGeneratorServiceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Integration Sequencer Generated Sequence Data
type IntegrationSequencerGeneratedSequenceDataType struct {
	IntegrationSequencerReference *IntegrationSequencerObjectType `xml:"urn:com.workday/bsvc Integration_Sequencer_Reference,omitempty"`
	SequencedValue                string                          `xml:"urn:com.workday/bsvc Sequenced_Value,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationSequencerObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type IntegrationSequencerObjectType struct {
	ID         []IntegrationSequencerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Service Component Field Override Data element
type IntegrationServiceComponentFieldOverrideDataType struct {
	FieldName              string                            `xml:"urn:com.workday/bsvc Field_Name"`
	SpecifyValueData       *IntegrationAbstractValueDataType `xml:"urn:com.workday/bsvc Specify_Value_Data"`
	AttachmentData         *IntegrationAttachmentDataType    `xml:"urn:com.workday/bsvc Attachment_Data"`
	ExternalFieldReference *ExternalFieldObjectType          `xml:"urn:com.workday/bsvc External_Field_Reference"`
}

// Integration Service Component Override Data element
type IntegrationServiceComponentOverrideDataType struct {
	DataSource        bool                                               `xml:"urn:com.workday/bsvc Data_Source"`
	Transformation    bool                                               `xml:"urn:com.workday/bsvc Transformation"`
	FileUtility       bool                                               `xml:"urn:com.workday/bsvc File_Utility"`
	TransportProtocol bool                                               `xml:"urn:com.workday/bsvc Transport_Protocol"`
	FieldOverrideData []IntegrationServiceComponentFieldOverrideDataType `xml:"urn:com.workday/bsvc Field_Override_Data"`
}

// Container element for data related to an Integration Service and its use within an Integration System.
type IntegrationServiceDataWWSType struct {
	IntegrationServiceReference                            *IntegrationServiceObjectType                               `xml:"urn:com.workday/bsvc Integration_Service_Reference"`
	IntegrationTransactionLogServiceData                   *IntegrationTransactionLogConfigurationDataType             `xml:"urn:com.workday/bsvc Integration_Transaction_Log_Service_Data,omitempty"`
	IntegrationSequenceGeneratorServiceData                []IntegrationSequenceGeneratorConfigurationDataType         `xml:"urn:com.workday/bsvc Integration_Sequence_Generator_Service_Data,omitempty"`
	IntegrationDataChangeServiceData                       *IntegrationDataChangeConfigurationDataType                 `xml:"urn:com.workday/bsvc Integration_Data_Change_Service_Data,omitempty"`
	IntegrationFieldOverrideServiceData                    *IntegrationFieldOverrideConfigurationDataType              `xml:"urn:com.workday/bsvc Integration_Field_Override_Service_Data,omitempty"`
	IntegrationDataInitializationOverrideConfigurationData *IntegrationDataInitializationOverrideConfigurationDataType `xml:"urn:com.workday/bsvc Integration_Data_Initialization_Override_Configuration_Data,omitempty"`
	IntegrationFieldAttributesServiceData                  *IntegrationFieldAttributesConfigurationDataType            `xml:"urn:com.workday/bsvc Integration_Field_Attributes_Service_Data,omitempty"`
	IntegrationDeliveryConfigurationData                   *IntegrationDeliveryConfigurationDataType                   `xml:"urn:com.workday/bsvc Integration_Delivery_Configuration_Data,omitempty"`
	JobIntegrationDeliveryConfigurationData                *JobIntegrationDeliveryConfigurationType                    `xml:"urn:com.workday/bsvc Job_Integration_Delivery_Configuration_Data,omitempty"`
	JobIntegrationTransformationConfigurationData          *JobIntegrationTransformationConfigurationDataType          `xml:"urn:com.workday/bsvc Job_Integration_Transformation_Configuration_Data,omitempty"`
	IntegrationRetrievalConfigurationData                  *IntegrationRetrievalConfigurationDataType                  `xml:"urn:com.workday/bsvc Integration_Retrieval_Configuration_Data,omitempty"`
	JobIntegrationRetrievalConfigurationData               *JobIntegrationRetrievalConfigurationDataType               `xml:"urn:com.workday/bsvc Job_Integration_Retrieval_Configuration_Data,omitempty"`
	IntegrationAttachmentServiceData                       *IntegrationAttachmentConfigurationDataType                 `xml:"urn:com.workday/bsvc Integration_Attachment_Service_Data,omitempty"`
	IntegrationReportServiceConfigurationData              *IntegrationReportServiceConfigurationDataType              `xml:"urn:com.workday/bsvc Integration_Report_Service_Configuration_Data,omitempty"`
	IntegrationCustomObjectConfigurationData               *IntegrationCustomObjectConfigurationDataType               `xml:"urn:com.workday/bsvc Integration_Custom_Object_Configuration_Data,omitempty"`
	IntegrationConnectorConfigurationData                  *IntegrationConnectorConfigurationDataType                  `xml:"urn:com.workday/bsvc Integration_Connector_Configuration_Data,omitempty"`
	IntegrationDataInitializationServiceData               *IntegrationDataInitializationConfigurationDataType         `xml:"urn:com.workday/bsvc Integration_Data_Initialization_Service_Data,omitempty"`
	WID                                                    string                                                      `xml:"urn:com.workday/bsvc WID,attr,omitempty"`
	Enabled                                                bool                                                        `xml:"urn:com.workday/bsvc Enabled,attr,omitempty"`
}

// Integration Service Generated Sequence Data
type IntegrationServiceGeneratedSequenceDataType struct {
	IntegrationServiceReference               *IntegrationSequenceGeneratorServiceObjectType  `xml:"urn:com.workday/bsvc Integration_Service_Reference,omitempty"`
	IntegrationSequencerGeneratedSequenceData []IntegrationSequencerGeneratedSequenceDataType `xml:"urn:com.workday/bsvc Integration_Sequencer_Generated_Sequence_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationServiceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntegrationServiceObjectType struct {
	ID         []IntegrationServiceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationSystemAuditedObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntegrationSystemAuditedObjectType struct {
	ID         []IntegrationSystemAuditedObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing data for Contacts related to the Integration System
type IntegrationSystemContactDataWWSType struct {
	IntegrationSystemContactReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Integration_System_Contact_Reference,omitempty"`
	Description                       string                      `xml:"urn:com.workday/bsvc Description,omitempty"`
	ContactReference                  *WorkerObjectType           `xml:"urn:com.workday/bsvc Contact_Reference,omitempty"`
}

// Element containing data for an Integration System, including Template reference, Attributes and Maps.
type IntegrationSystemDataWWSType struct {
	IntegrationSystemID          string                                                         `xml:"urn:com.workday/bsvc Integration_System_ID,omitempty"`
	IntegrationSystemName        string                                                         `xml:"urn:com.workday/bsvc Integration_System_Name,omitempty"`
	IntegrationTemplateReference *IntegrationTemplateObjectType                                 `xml:"urn:com.workday/bsvc Integration_Template_Reference,omitempty"`
	IntegrationTagData           []IntegrationTagforIntegrationSystemWWSDataType                `xml:"urn:com.workday/bsvc Integration_Tag_Data,omitempty"`
	IntegrationSystemContact     []IntegrationSystemContactDataWWSType                          `xml:"urn:com.workday/bsvc Integration_System_Contact,omitempty"`
	IntegrationServiceData       []IntegrationServiceDataWWSType                                `xml:"urn:com.workday/bsvc Integration_Service_Data,omitempty"`
	IntegrationAttributesData    []IntegrationToolProviderIntegrationAttributeValuesDataWWSType `xml:"urn:com.workday/bsvc Integration_Attributes_Data,omitempty"`
	IntegrationMapsData          []IntegrationToolProviderIntegrationMapValuesDataWWSType       `xml:"urn:com.workday/bsvc Integration_Maps_Data,omitempty"`
	CustomAttributeData          []IntegrationAttributeDataType                                 `xml:"urn:com.workday/bsvc Custom_Attribute_Data,omitempty"`
	CustomMapData                []IntegrationMapDataType                                       `xml:"urn:com.workday/bsvc Custom_Map_Data,omitempty"`
	CustomLaunchParameterData    []LaunchParameterDataType                                      `xml:"urn:com.workday/bsvc Custom_Launch_Parameter_Data,omitempty"`
	IntegrationNotificationData  []IntegrationNotificationDataType                              `xml:"urn:com.workday/bsvc Integration_Notification_Data,omitempty"`
	SubscriptionData             []SubscriptionDataWWSType                                      `xml:"urn:com.workday/bsvc Subscription_Data,omitempty"`
	WebServiceOperationData      []WebServiceOperationDataWWSType                               `xml:"urn:com.workday/bsvc Web_Service_Operation_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationSystemIdentifierObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type IntegrationSystemIdentifierObjectType struct {
	ID         []IntegrationSystemIdentifierObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationSystemObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntegrationSystemObjectType struct {
	ID         []IntegrationSystemObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Pass in a request criteria
type IntegrationSystemRequestCriteriaType struct {
	SystemUserReference []SystemUserObjectType `xml:"urn:com.workday/bsvc System_User_Reference,omitempty"`
}

// Utilize the Request References element to retrieve a specific instance(s) of Integration System and its associated data.
type IntegrationSystemRequestReferencesType struct {
	IntegrationSystemReference []IntegrationSystemAuditedObjectType `xml:"urn:com.workday/bsvc Integration_System_Reference"`
}

// Contains each Integration System based on the Request Criteria.  The data returned is the current information as of the dates in the response filter.
type IntegrationSystemResponseDataType struct {
	IntegrationSystem []IntegrationSystemWWSType `xml:"urn:com.workday/bsvc Integration_System,omitempty"`
}

// The response groups allows for the amount of data contained within the response to be tailored to only included elements that the user is looking for.  If no response group is provided in the request, then ONLY the payee references will be returned. If no response group is provided AND the Page and Count parameters within the Response Filter element are also omitted, then ALL of the payee references will be returned in your response (no paging is performed).
type IntegrationSystemResponseGroupType struct {
	IncludeReference             *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	ShowValuesForAllEnvironments *bool `xml:"urn:com.workday/bsvc Show_Values_For_All_Environments,omitempty"`
}

// Integration System User Data element
type IntegrationSystemUserDataType struct {
	IntegrationSystemReference          *IntegrationSystemAuditedObjectType       `xml:"urn:com.workday/bsvc Integration_System_Reference,omitempty"`
	UserName                            string                                    `xml:"urn:com.workday/bsvc User_Name,omitempty"`
	Password                            string                                    `xml:"urn:com.workday/bsvc Password,omitempty"`
	GenerateRandomPassword              *bool                                     `xml:"urn:com.workday/bsvc Generate_Random_Password,omitempty"`
	RequireNewPasswordAtNextSignIn      *bool                                     `xml:"urn:com.workday/bsvc Require_New_Password_At_Next_Sign_In,omitempty"`
	SessionTimeoutMinutes               float64                                   `xml:"urn:com.workday/bsvc Session_Timeout_Minutes,omitempty"`
	DoNotAllowUISessions                *bool                                     `xml:"urn:com.workday/bsvc Do_Not_Allow_UI_Sessions,omitempty"`
	WebServiceSecurityConfigurationData []WebServiceSecurityConfigurationDataType `xml:"urn:com.workday/bsvc Web_Service_Security_Configuration_Data,omitempty"`
}

// Integration System User Response Data element
type IntegrationSystemUserResponseDataType struct {
	IntegrationSystemUser []IntegrationSystemUserType `xml:"urn:com.workday/bsvc Integration_System_User,omitempty"`
}

// Integration System User element
type IntegrationSystemUserType struct {
	IntegrationSystemReference *IntegrationSystemAuditedObjectType `xml:"urn:com.workday/bsvc Integration_System_Reference"`
	IntegrationSystemUserData  *IntegrationSystemUserDataType      `xml:"urn:com.workday/bsvc Integration_System_User_Data,omitempty"`
}

// Encapsulating element containing all Integration System data.
type IntegrationSystemWWSType struct {
	IntegrationSystemReference *IntegrationSystemAuditedObjectType            `xml:"urn:com.workday/bsvc Integration_System_Reference,omitempty"`
	IntegrationSystemData      *IntegrationSystemDataWWSType                  `xml:"urn:com.workday/bsvc Integration_System_Data,omitempty"`
	ExceptionsResponseData     []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationTagObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntegrationTagObjectType struct {
	ID         []IntegrationTagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Integration Tag Data
type IntegrationTagforIntegrationSystemWWSDataType struct {
	IntegrationTagReference []IntegrationTagObjectType `xml:"urn:com.workday/bsvc Integration_Tag_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationTemplateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntegrationTemplateObjectType struct {
	ID         []IntegrationTemplateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Container element for all the Integration Attributes associated with a specific Attributable instance (e.g. Integration Template, Integration Service).
type IntegrationToolProviderIntegrationAttributeValuesDataWWSType struct {
	IntegrationAttributeProviderReference *ExternalInstanceObjectType                                `xml:"urn:com.workday/bsvc Integration_Attribute_Provider_Reference,omitempty"`
	IntegrationAttributeData              []IntegrationAttributeIntegrationAttributeValueDataWWSType `xml:"urn:com.workday/bsvc Integration_Attribute_Data"`
}

// Container element for all the Integration Maps associated with a specific Attributable instance (e.g. Integration Template, Integration Service).
type IntegrationToolProviderIntegrationMapValuesDataWWSType struct {
	IntegrationMapProviderReference *ExternalInstanceObjectType                    `xml:"urn:com.workday/bsvc Integration_Map_Provider_Reference,omitempty"`
	IntegrationMapData              []IntegrationMapIntegrationMapValueDataWWSType `xml:"urn:com.workday/bsvc Integration_Map_Data,omitempty"`
}

// Container element for data related to a Transaction Log Service and its use within an Integration System.
type IntegrationTransactionLogConfigurationDataType struct {
	AllBusinessProcesses                 bool                            `xml:"urn:com.workday/bsvc All_Business_Processes"`
	ExcludedBusinessProcessTypeReference []BusinessProcessTypeObjectType `xml:"urn:com.workday/bsvc Excluded_Business_Process_Type_Reference"`
	AllTransactionTypes                  bool                            `xml:"urn:com.workday/bsvc All_Transaction_Types"`
	ExcludedTransactionLogTypeReference  []TransactionLogTypeObjectType  `xml:"urn:com.workday/bsvc Excluded_Transaction_Log_Type_Reference"`
	IncludedTransactionLogTypeReference  []TransactionLogTypeObjectType  `xml:"urn:com.workday/bsvc Included_Transaction_Log_Type_Reference"`
}

// Integration Transformation Data element
type IntegrationTransformationDataWWSType struct {
	Name                                  string                                  `xml:"urn:com.workday/bsvc Name"`
	DeliveredTransformationData           *DeliveredTransformationDataType        `xml:"urn:com.workday/bsvc Delivered_Transformation_Data"`
	XSLTAttachmentTransformationReference *XSLTAttachmentTransformationObjectType `xml:"urn:com.workday/bsvc XSLT_Attachment_Transformation_Reference"`
	CustomReportTransformationData        *CustomReportTransformationDataType     `xml:"urn:com.workday/bsvc Custom_Report_Transformation_Data"`
}

// Integration Data Communication Data element
type IntegrationTransportProtocolAssignmentDataType struct {
	IntegrationTransportProtocolData *IntegrationTransportProtocolDataWWSType `xml:"urn:com.workday/bsvc Integration_Transport_Protocol_Data"`
	IntegrationFileUtilityData       *IntegrationFileUtilityDataType          `xml:"urn:com.workday/bsvc Integration_File_Utility_Data,omitempty"`
	IntegrationPayloadData           *IntegrationPayloadDataType              `xml:"urn:com.workday/bsvc Integration_Payload_Data,omitempty"`
	RestrictedToEnvironmentReference []OMSEnvironmentObjectType               `xml:"urn:com.workday/bsvc Restricted_To_Environment_Reference,omitempty"`
	PreviewOnly                      *bool                                    `xml:"urn:com.workday/bsvc Preview_Only,omitempty"`
}

// Integration Transport Protocol Data element
type IntegrationTransportProtocolDataWWSType struct {
	EmailTransportProtocolReference        *EmailTransportProtocolObjectType                    `xml:"urn:com.workday/bsvc Email_Transport_Protocol_Reference,omitempty"`
	EmailTransportProtocolData             *EmailTransportProtocolDataType                      `xml:"urn:com.workday/bsvc Email_Transport_Protocol_Data,omitempty"`
	FTPSTransportProtocolReference         *FTPSSLTransportProtocolObjectType                   `xml:"urn:com.workday/bsvc FTPS_Transport_Protocol_Reference,omitempty"`
	FTPSTransportProtocolData              *FTPSTransportProtocolDataType                       `xml:"urn:com.workday/bsvc FTPS_Transport_Protocol_Data,omitempty"`
	FTPTransportProtocolReference          *FTPTransportProtocolObjectType                      `xml:"urn:com.workday/bsvc FTP_Transport_Protocol_Reference,omitempty"`
	FTPTransportProtocolData               *FTPTransportProtocolDataType                        `xml:"urn:com.workday/bsvc FTP_Transport_Protocol_Data,omitempty"`
	HTTPTransportProtocolReference         *HTTPSSLTransportProtocolObjectType                  `xml:"urn:com.workday/bsvc HTTP_Transport_Protocol_Reference,omitempty"`
	HTTPTransportProtocolData              *HTTPTransportProtocolDataType                       `xml:"urn:com.workday/bsvc HTTP_Transport_Protocol_Data,omitempty"`
	SFTPTransportProtocolReference         *SFTPTransportProtocolObjectType                     `xml:"urn:com.workday/bsvc SFTP_Transport_Protocol_Reference,omitempty"`
	SFTPTransportProtocolData              *SFTPTransportProtocolDatawithDualAuthenticationType `xml:"urn:com.workday/bsvc SFTP_Transport_Protocol_Data,omitempty"`
	WorkdayAttachmentTransportProtocolData *WorkdayAttachmentTransportProtocolDataType          `xml:"urn:com.workday/bsvc Workday_Attachment_Transport_Protocol_Data,omitempty"`
	WorkdayWebServiceTransportProtocolData *WorkdayWebServiceTransportProtocolDataType          `xml:"urn:com.workday/bsvc Workday_Web_Service_Transport_Protocol_Data,omitempty"`
	AS2TransportProtocolReference          *AS2TransportProtocolObjectType                      `xml:"urn:com.workday/bsvc AS2_Transport_Protocol_Reference,omitempty"`
	AS2TransportProtocolData               *AS2TransportProtocolDataType                        `xml:"urn:com.workday/bsvc AS2_Transport_Protocol_Data,omitempty"`
	ExtensibilityAPITransportProtocolData  *EIBExtensibilityAPITransportProtocolDataType        `xml:"urn:com.workday/bsvc Extensibility_API_Transport_Protocol_Data,omitempty"`
}

// Wraps notification email addresses.
type InternetEmailAddressDataforNotificationsType struct {
	EmailAddress string `xml:"urn:com.workday/bsvc Email_Address"`
}

// Contains Information about the configuration of a Job Integration Delivery Configuration, used for configuring an EIB delivery step
type JobIntegrationDeliveryConfigurationType struct {
	IntegrationTransportProtocolAssignmentData []IntegrationTransportProtocolAssignmentDataType `xml:"urn:com.workday/bsvc Integration_Transport_Protocol_Assignment_Data,omitempty"`
}

// Contains Information about the configuration of a Job Integration Retrieval Configuration, used for configuring an EIB retrieval step
type JobIntegrationRetrievalConfigurationDataType struct {
	IntegrationDataSourceData []JobIntegrationRetrievalProtocolDataType `xml:"urn:com.workday/bsvc Integration_Data_Source_Data,omitempty"`
}

// Describes how the Job Integration Retrieval Configuration (Audited) is configured
type JobIntegrationRetrievalProtocolDataType struct {
	InboundProtocolData                    *JobRetrievalInboundProtocolDataType `xml:"urn:com.workday/bsvc Inbound_Protocol_Data"`
	DecryptUsingPGPPrivateKeyPairReference *PGPPrivateKeyPairObjectType         `xml:"urn:com.workday/bsvc Decrypt_Using_PGP_Private_Key_Pair_Reference,omitempty"`
	Filename                               string                               `xml:"urn:com.workday/bsvc Filename,omitempty"`
	RestrictedToEnvironmentReference       []OMSEnvironmentObjectType           `xml:"urn:com.workday/bsvc Restricted_To_Environment_Reference,omitempty"`
}

// Transformation Configuration element
type JobIntegrationTransformationConfigurationDataType struct {
	IntegrationTransformationData []IntegrationTransformationDataWWSType `xml:"urn:com.workday/bsvc Integration_Transformation_Data,omitempty"`
}

// Describes how the input file is retrieved
type JobRetrievalInboundProtocolDataType struct {
	AttachmentDataSourceData        *EIBAttachmentDataSourceDataType                     `xml:"urn:com.workday/bsvc Attachment_Data_Source_Data"`
	CustomReportDataSourceData      *ReportBackgroundFutureProcessasCustomReportDataType `xml:"urn:com.workday/bsvc Custom_Report_Data_Source_Data"`
	ExternalFileDataSource          *EIBExternalFileDataSourceDataType                   `xml:"urn:com.workday/bsvc External_File_Data_Source"`
	RESTEndpointDataSourceReference *RESTEndpointDataSourceObjectType                    `xml:"urn:com.workday/bsvc REST_Endpoint_Data_Source_Reference"`
	RESTEndpointDataSourceData      *RESTEndpointDataSourceDataType                      `xml:"urn:com.workday/bsvc REST_Endpoint_Data_Source_Data"`
	WebServiceDataSourceReference   *WebServiceDataSourceObjectType                      `xml:"urn:com.workday/bsvc Web_Service_Data_Source_Reference"`
}

// Launch EIB Request
type LaunchEIBRequestType struct {
	IntegrationSystemReference *IntegrationSystemAuditedObjectType           `xml:"urn:com.workday/bsvc Integration_System_Reference"`
	ServiceComponentData       []IntegrationServiceComponentOverrideDataType `xml:"urn:com.workday/bsvc Service_Component_Data,omitempty"`
	Version                    string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Launch EIB Response
type LaunchEIBResponseType struct {
	ApplicationInstanceExceptions []ApplicationInstanceExceptionsDataType `xml:"urn:com.workday/bsvc Application_Instance_Exceptions,omitempty"`
	IntegrationEvent              *IntegrationEventType                   `xml:"urn:com.workday/bsvc Integration_Event,omitempty"`
	Version                       string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Launch Integration Event Data
type LaunchIntegrationEventDataType struct {
	IntegrationEventReference               *IntegrationESBInvocationAbstractObjectType   `xml:"urn:com.workday/bsvc Integration_Event_Reference,omitempty"`
	IntegrationSystemReference              *IntegrationSystemObjectType                  `xml:"urn:com.workday/bsvc Integration_System_Reference,omitempty"`
	DocumentID                              string                                        `xml:"urn:com.workday/bsvc Document_ID,omitempty"`
	EventInitializationDocumentID           string                                        `xml:"urn:com.workday/bsvc Event_Initialization_Document_ID,omitempty"`
	IntegrationTemplateReference            *IntegrationTemplateObjectType                `xml:"urn:com.workday/bsvc Integration_Template_Reference,omitempty"`
	SentOn                                  *time.Time                                    `xml:"urn:com.workday/bsvc Sent_On,omitempty"`
	SystemUserReference                     *SystemUserObjectType                         `xml:"urn:com.workday/bsvc System_User_Reference,omitempty"`
	ResendfromIntegrationEventReference     *IntegrationEventAbstractObjectType           `xml:"urn:com.workday/bsvc Resend_from_Integration_Event_Reference,omitempty"`
	IntegrationRuntimeParameterData         []IntegrationRuntimeParameterDataType         `xml:"urn:com.workday/bsvc Integration_Runtime_Parameter_Data,omitempty"`
	IntegrationServiceGeneratedSequenceData []IntegrationServiceGeneratedSequenceDataType `xml:"urn:com.workday/bsvc Integration_Service_Generated_Sequence_Data,omitempty"`
	ParentEventReference                    *EventObjectType                              `xml:"urn:com.workday/bsvc Parent_Event_Reference,omitempty"`
	WINIntegration                          bool                                          `xml:"urn:com.workday/bsvc WIN_Integration,attr,omitempty"`
}

func (t *LaunchIntegrationEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LaunchIntegrationEventDataType
	var layout struct {
		*T
		SentOn *xsdDateTime `xml:"urn:com.workday/bsvc Sent_On,omitempty"`
	}
	layout.T = (*T)(t)
	layout.SentOn = (*xsdDateTime)(layout.T.SentOn)
	return e.EncodeElement(layout, start)
}
func (t *LaunchIntegrationEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LaunchIntegrationEventDataType
	var overlay struct {
		*T
		SentOn *xsdDateTime `xml:"urn:com.workday/bsvc Sent_On,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.SentOn = (*xsdDateTime)(overlay.T.SentOn)
	return d.DecodeElement(&overlay, &start)
}

// Launch Integration Event Request
type LaunchIntegrationEventRequestType struct {
	IntegrationSystemReference     *IntegrationSystemAuditedObjectType     `xml:"urn:com.workday/bsvc Integration_System_Reference"`
	IntegrationLaunchParameterData []IntegrationLaunchParameterDataType    `xml:"urn:com.workday/bsvc Integration_Launch_Parameter_Data,omitempty"`
	ListenerDocumentData           []IntegrationRepositoryDocumentDataType `xml:"urn:com.workday/bsvc Listener_Document_Data,omitempty"`
	DebugMode                      bool                                    `xml:"urn:com.workday/bsvc Debug_Mode,attr,omitempty"`
	Version                        string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Launch Integration Event Response
type LaunchIntegrationEventResponseType struct {
	IntegrationEvent              *IntegrationEventType                   `xml:"urn:com.workday/bsvc Integration_Event,omitempty"`
	LaunchIntegrationEventData    *LaunchIntegrationEventDataType         `xml:"urn:com.workday/bsvc Launch_Integration_Event_Data,omitempty"`
	ApplicationInstanceExceptions []ApplicationInstanceExceptionsDataType `xml:"urn:com.workday/bsvc Application_Instance_Exceptions,omitempty"`
	DebugMode                     bool                                    `xml:"urn:com.workday/bsvc Debug_Mode,attr,omitempty"`
	Version                       string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Container element for all the Integration Launch Parameters for this Integration System.
type LaunchParameterDataType struct {
	Order                                  string                                      `xml:"urn:com.workday/bsvc Order"`
	Name                                   string                                      `xml:"urn:com.workday/bsvc Name"`
	Description                            string                                      `xml:"urn:com.workday/bsvc Description,omitempty"`
	DataTypeEnumerationDefinitionReference *IntegrationEnumerationDefinitionObjectType `xml:"urn:com.workday/bsvc Data_Type_Enumeration_Definition_Reference"`
	DataTypeExternalFieldReference         *ExternalFieldObjectType                    `xml:"urn:com.workday/bsvc Data_Type_External_Field_Reference"`
	LaunchOptionReference                  []IntegrationLaunchOptionObjectType         `xml:"urn:com.workday/bsvc Launch_Option_Reference,omitempty"`
	LaunchParameterDefaultData             *LaunchParameterDefaultDataType             `xml:"urn:com.workday/bsvc Launch_Parameter_Default_Data,omitempty"`
}

// Element to define Default processing for Launch Parameter
type LaunchParameterDefaultDataType struct {
	ValueTypeReference     *ParameterInitializationOperatorObjectType `xml:"urn:com.workday/bsvc Value_Type_Reference"`
	ExternalFieldReference *ExternalFieldObjectType                   `xml:"urn:com.workday/bsvc External_Field_Reference,omitempty"`
	Boolean                *bool                                      `xml:"urn:com.workday/bsvc Boolean,omitempty"`
	Currency               float64                                    `xml:"urn:com.workday/bsvc Currency,omitempty"`
	CurrencyReference      *CurrencyObjectType                        `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	Date                   *time.Time                                 `xml:"urn:com.workday/bsvc Date,omitempty"`
	DateTime               *time.Time                                 `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	Numeric                float64                                    `xml:"urn:com.workday/bsvc Numeric,omitempty"`
	Text                   string                                     `xml:"urn:com.workday/bsvc Text,omitempty"`
	InstanceReference      []InstanceObjectType                       `xml:"urn:com.workday/bsvc Instance_Reference,omitempty"`
	EnumerationReference   *IntegrationEnumerationObjectType          `xml:"urn:com.workday/bsvc Enumeration_Reference,omitempty"`
}

func (t *LaunchParameterDefaultDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LaunchParameterDefaultDataType
	var layout struct {
		*T
		Date     *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		DateTime *xsdDateTime `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Date = (*xsdDate)(layout.T.Date)
	layout.DateTime = (*xsdDateTime)(layout.T.DateTime)
	return e.EncodeElement(layout, start)
}
func (t *LaunchParameterDefaultDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LaunchParameterDefaultDataType
	var overlay struct {
		*T
		Date     *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		DateTime *xsdDateTime `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Date = (*xsdDate)(overlay.T.Date)
	overlay.DateTime = (*xsdDateTime)(overlay.T.DateTime)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type LaunchParameterObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type LaunchParameterObjectType struct {
	ID         []LaunchParameterObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MimeTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MimeTypeObjectType struct {
	ID         []MimeTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Notification Attachment Data
type NotificationAttachmentDataType struct {
	AttachmentExternalFieldReference []ExternalFieldObjectType `xml:"urn:com.workday/bsvc Attachment_External_Field_Reference"`
}

// Element containing the components that make up the Body of the email message.
type NotificationBodyDataType struct {
	NotificationBodyData []NotificationMessageComponentDataType `xml:"urn:com.workday/bsvc Notification_Body_Data,omitempty"`
}

// Element defining a component to be used in composing a message.
type NotificationMessageComponentDataType struct {
	Order                  string                   `xml:"urn:com.workday/bsvc Order"`
	Text                   string                   `xml:"urn:com.workday/bsvc Text"`
	ExternalFieldReference *ExternalFieldObjectType `xml:"urn:com.workday/bsvc External_Field_Reference"`
}

// Notification Notifies Data
type NotificationNotifiesDataType struct {
	NotifiesRecipientExternalFieldReference []ExternalFieldObjectType `xml:"urn:com.workday/bsvc Notifies_Recipient_External_Field_Reference"`
}

// Element containing the components that make up the Subject of the email message.
type NotificationSubjectDataType struct {
	NotificationMessageComponentData []NotificationMessageComponentDataType `xml:"urn:com.workday/bsvc Notification_Message_Component_Data,omitempty"`
}

type Numeric float64

// Contains a unique identifier for an instance of an object.
type OMSEnvironmentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type OMSEnvironmentObjectType struct {
	ID         []OMSEnvironmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Integration Document Field for Override Data
type OverriddenIntegrationDocumentFieldDataType struct {
	IntegrationDocumentFieldReference                                     *IntegrationDocumentFieldObjectType                       `xml:"urn:com.workday/bsvc Integration_Document_Field_Reference"`
	ExternalFieldReference                                                *ExternalFieldObjectType                                  `xml:"urn:com.workday/bsvc External_Field_Reference"`
	IntegrationFieldOverrideParameterInitializationData                   []IntegrationFieldOverrideParameterInitializationDataType `xml:"urn:com.workday/bsvc Integration_Field_Override_Parameter_Initialization_Data,omitempty"`
	IntegrationDocumentStacksforIntegrationDocumentFieldConfigurationData []IntegrationDocumentStackDataType                        `xml:"urn:com.workday/bsvc Integration_Document_Stacks_for_Integration_Document_Field_Configuration_Data,omitempty"`
	IntegrationDocumentFieldOptions                                       []IntegrationDocumentFieldOptionsType                     `xml:"urn:com.workday/bsvc Integration_Document_Field_Options,omitempty"`
}

// PGP Encryption Settings Data element
type PGPEncryptionSettingsDataType struct {
	CertificateReference          *PGPPublicKeyObjectType      `xml:"urn:com.workday/bsvc Certificate_Reference,omitempty"`
	TextMode                      *bool                        `xml:"urn:com.workday/bsvc Text_Mode,omitempty"`
	AsciiArmored                  *bool                        `xml:"urn:com.workday/bsvc Ascii_Armored,omitempty"`
	ContainingIntegrityCheck      *bool                        `xml:"urn:com.workday/bsvc Containing_Integrity_Check,omitempty"`
	DecryptedFilename             string                       `xml:"urn:com.workday/bsvc Decrypted_Filename,omitempty"`
	PGP26xCompatible              *bool                        `xml:"urn:com.workday/bsvc PGP_2.6x_Compatible,omitempty"`
	DigitallySignKeyPairReference *PGPPrivateKeyPairObjectType `xml:"urn:com.workday/bsvc Digitally_Sign_Key_Pair_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PGPPrivateKeyPairObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PGPPrivateKeyPairObjectType struct {
	ID         []PGPPrivateKeyPairObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PGPPublicKeyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PGPPublicKeyObjectType struct {
	ID         []PGPPublicKeyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element to define Parameters Exception to not fire Events in Concurrency
type ParameterExceptionDataType struct {
	LaunchParameterReference *LaunchParameterObjectType  `xml:"urn:com.workday/bsvc Launch_Parameter_Reference,omitempty"`
	SimpleWorkDataReference  *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Simple_Work_Data_Reference,omitempty"`
	XMLName                  string                      `xml:"urn:com.workday/bsvc XML_Name,omitempty"`
}

// Defines the value to be used for parameter assignment.
type ParameterInitializationDataType struct {
	Boolean                                  *bool                                      `xml:"urn:com.workday/bsvc Boolean,omitempty"`
	Currency                                 float64                                    `xml:"urn:com.workday/bsvc Currency,omitempty"`
	Date                                     *time.Time                                 `xml:"urn:com.workday/bsvc Date,omitempty"`
	DateTime                                 *time.Time                                 `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	Numeric                                  float64                                    `xml:"urn:com.workday/bsvc Numeric,omitempty"`
	Text                                     string                                     `xml:"urn:com.workday/bsvc Text,omitempty"`
	BusinessObjectInstanceReference          []InstanceObjectType                       `xml:"urn:com.workday/bsvc Business_Object_Instance_Reference,omitempty"`
	ExternalFieldContent                     *ExternalFieldAddorReferenceType           `xml:"urn:com.workday/bsvc External_Field_Content,omitempty"`
	ParameterInitializationOperatorReference *ParameterInitializationOperatorObjectType `xml:"urn:com.workday/bsvc Parameter_Initialization_Operator_Reference"`
	CurrencyReference                        *CurrencyObjectType                        `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
}

func (t *ParameterInitializationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ParameterInitializationDataType
	var layout struct {
		*T
		Date     *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		DateTime *xsdDateTime `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Date = (*xsdDate)(layout.T.Date)
	layout.DateTime = (*xsdDateTime)(layout.T.DateTime)
	return e.EncodeElement(layout, start)
}
func (t *ParameterInitializationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ParameterInitializationDataType
	var overlay struct {
		*T
		Date     *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		DateTime *xsdDateTime `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Date = (*xsdDate)(overlay.T.Date)
	overlay.DateTime = (*xsdDateTime)(overlay.T.DateTime)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type ParameterInitializationOperatorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ParameterInitializationOperatorObjectType struct {
	ID         []ParameterInitializationOperatorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Defines the value to be used for parameter assignment.
type ParameterInitializationWWSDataType struct {
	Boolean                                  *bool                                      `xml:"urn:com.workday/bsvc Boolean,omitempty"`
	Currency                                 float64                                    `xml:"urn:com.workday/bsvc Currency,omitempty"`
	Date                                     *time.Time                                 `xml:"urn:com.workday/bsvc Date,omitempty"`
	DateTime                                 *time.Time                                 `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	Numeric                                  float64                                    `xml:"urn:com.workday/bsvc Numeric,omitempty"`
	Text                                     string                                     `xml:"urn:com.workday/bsvc Text,omitempty"`
	BusinessObjectInstanceReference          []InstanceObjectType                       `xml:"urn:com.workday/bsvc Business_Object_Instance_Reference,omitempty"`
	ExternalFieldReference                   *ExternalFieldObjectType                   `xml:"urn:com.workday/bsvc External_Field_Reference,omitempty"`
	ReportPromptReference                    *ExternalPromptableObjectType              `xml:"urn:com.workday/bsvc Report_Prompt_Reference,omitempty"`
	IntegrationParameterReference            *IntegrationParameterReferenceType         `xml:"urn:com.workday/bsvc Integration_Parameter_Reference,omitempty"`
	ParameterInitializationOperatorReference *ParameterInitializationOperatorObjectType `xml:"urn:com.workday/bsvc Parameter_Initialization_Operator_Reference"`
	CurrencyReference                        *CurrencyObjectType                        `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
}

func (t *ParameterInitializationWWSDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ParameterInitializationWWSDataType
	var layout struct {
		*T
		Date     *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		DateTime *xsdDateTime `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Date = (*xsdDate)(layout.T.Date)
	layout.DateTime = (*xsdDateTime)(layout.T.DateTime)
	return e.EncodeElement(layout, start)
}
func (t *ParameterInitializationWWSDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ParameterInitializationWWSDataType
	var overlay struct {
		*T
		Date     *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		DateTime *xsdDateTime `xml:"urn:com.workday/bsvc DateTime,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Date = (*xsdDate)(overlay.T.Date)
	overlay.DateTime = (*xsdDateTime)(overlay.T.DateTime)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PartitionedBackgroundProcessObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PartitionedBackgroundProcessObjectType struct {
	ID         []PartitionedBackgroundProcessObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type PartitionedBackgroundProcessReferenceEnumeration string

// Post Population Predicate
type PopulationEligibilityDataType struct {
	ExternalFieldReference                              *ExternalFieldObjectType                                  `xml:"urn:com.workday/bsvc External_Field_Reference,omitempty"`
	IntegrationFieldOverrideParameterInitializationData []IntegrationFieldOverrideParameterInitializationDataType `xml:"urn:com.workday/bsvc Integration_Field_Override_Parameter_Initialization_Data,omitempty"`
}

type ProcessingFaultType struct {
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
}

// Integration event request element.
type PutIntegrationEventRequestType struct {
	IntegrationEventReference *IntegrationESBInvocationAbstractObjectType `xml:"urn:com.workday/bsvc Integration_Event_Reference,omitempty"`
	IntegrationEventData      *IntegrationEventDataWWSType                `xml:"urn:com.workday/bsvc Integration_Event_Data"`
	Version                   string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Integration Event response element.
type PutIntegrationEventResponseType struct {
	IntegrationEventReference *IntegrationESBInvocationAbstractObjectType `xml:"urn:com.workday/bsvc Integration_Event_Reference,omitempty"`
	Version                   string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Integration Message request element.
type PutIntegrationMessageRequestType struct {
	IntegrationMessageData *IntegrationMessageDataType `xml:"urn:com.workday/bsvc Integration_Message_Data"`
	Version                string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Integration Message response element.
type PutIntegrationMessageResponseType struct {
	IntegrationMessageReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Integration_Message_Reference,omitempty"`
	Version                     string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Integration System request element.
type PutIntegrationSystemRequestType struct {
	IntegrationSystemReference *IntegrationSystemAuditedObjectType `xml:"urn:com.workday/bsvc Integration_System_Reference,omitempty"`
	IntegrationSystemData      *IntegrationSystemDataWWSType       `xml:"urn:com.workday/bsvc Integration_System_Data"`
	AddOnly                    bool                                `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                    string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Integration System Response element
type PutIntegrationSystemResponseType struct {
	IntegrationSystemReference *IntegrationSystemAuditedObjectType            `xml:"urn:com.workday/bsvc Integration_System_Reference,omitempty"`
	ExceptionsResponseData     []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                    string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Integration System User Request element
type PutIntegrationSystemUserRequestType struct {
	IntegrationSystemReference *IntegrationSystemAuditedObjectType `xml:"urn:com.workday/bsvc Integration_System_Reference,omitempty"`
	IntegrationSystemUserData  *IntegrationSystemUserDataType      `xml:"urn:com.workday/bsvc Integration_System_User_Data,omitempty"`
	Version                    string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Integration System User Response element
type PutIntegrationSystemUserResponseType struct {
	IntegrationSystemReference     *IntegrationSystemAuditedObjectType `xml:"urn:com.workday/bsvc Integration_System_Reference,omitempty"`
	IntegrationSystemUserReference *SystemUserObjectType               `xml:"urn:com.workday/bsvc Integration_System_User_Reference,omitempty"`
	Version                        string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for Put Reference.
type PutReferenceRequestType struct {
	ReferenceIDReference *ReferenceIndexObjectType `xml:"urn:com.workday/bsvc Reference_ID_Reference,omitempty"`
	ReferenceIDData      *ReferenceIDDataType      `xml:"urn:com.workday/bsvc Reference_ID_Data,omitempty"`
	Version              string                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Put Reference.
type PutReferenceResponseType struct {
	ReferenceIDReference *ReferenceIndexObjectType `xml:"urn:com.workday/bsvc Reference_ID_Reference,omitempty"`
	Version              string                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Sequence Generator Request element
type PutSequenceGeneratorRequestType struct {
	SequenceGeneratorReference *SequenceGeneratorObjectType       `xml:"urn:com.workday/bsvc Sequence_Generator_Reference,omitempty"`
	SequenceGeneratorData      *AbstractSequenceGeneratorDataType `xml:"urn:com.workday/bsvc Sequence_Generator_Data"`
	AddOnly                    bool                               `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                    string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Sequence Generator Response element
type PutSequenceGeneratorResponseType struct {
	SequenceGeneratorReference *SequenceGeneratorObjectType `xml:"urn:com.workday/bsvc Sequence_Generator_Reference,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Subscription Request
type PutSubscriptionRequestType struct {
	SubscriptionReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Subscription_Reference,omitempty"`
	SubscriptionData      *SubscriptionDataType       `xml:"urn:com.workday/bsvc Subscription_Data"`
	Version               string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Subscription Response
type PutSubscriptionResponseType struct {
	SubscriptionReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Subscription_Reference,omitempty"`
	Version               string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// REST Endpoint Data Source Data element
type RESTEndpointDataSourceDataType struct {
	RESTEndpoint  string `xml:"urn:com.workday/bsvc REST_Endpoint"`
	UserID        string `xml:"urn:com.workday/bsvc User_ID,omitempty"`
	Password      string `xml:"urn:com.workday/bsvc Password,omitempty"`
	IsInternalURL bool   `xml:"urn:com.workday/bsvc isInternalURL,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RESTEndpointDataSourceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RESTEndpointDataSourceObjectType struct {
	ID         []RESTEndpointDataSourceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing the data for the reassignment.
type ReassignBusinessProcessStepDataType struct {
	ReassignToReference   *RoleObjectType `xml:"urn:com.workday/bsvc Reassign_To_Reference"`
	ReassignFromReference *RoleObjectType `xml:"urn:com.workday/bsvc Reassign_From_Reference,omitempty"`
	Reason                string          `xml:"urn:com.workday/bsvc Reason,omitempty"`
}

// Top-level request element for the Reassign Business Process Step operation
type ReassignBusinessProcessStepRequestType struct {
	EventRecordReference            *UniqueIdentifierObjectType          `xml:"urn:com.workday/bsvc Event_Record_Reference"`
	ReassignBusinessProcessStepData *ReassignBusinessProcessStepDataType `xml:"urn:com.workday/bsvc Reassign_Business_Process_Step_Data"`
	Version                         string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Reference to business process event record processed.
type ReassignBusinessProcessStepResponseType struct {
	EventRecordReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Record_Reference,omitempty"`
	Version              string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing reference ID and type data.
type ReferenceIDDataType struct {
	ID                         string `xml:"urn:com.workday/bsvc ID,omitempty"`
	NewID                      string `xml:"urn:com.workday/bsvc New_ID,omitempty"`
	ReferenceIDType            string `xml:"urn:com.workday/bsvc Reference_ID_Type"`
	ReferencedObjectDescriptor string `xml:"urn:com.workday/bsvc Referenced_Object_Descriptor,omitempty"`
}

// The "Business Object" element contains object instances along with their corresponding Reference and Workday IDs.
type ReferenceIDType struct {
	ReferenceIDReference *ReferenceIndexObjectType `xml:"urn:com.workday/bsvc Reference_ID_Reference,omitempty"`
	ReferenceIDData      *ReferenceIDDataType      `xml:"urn:com.workday/bsvc Reference_ID_Data,omitempty"`
	Descriptor           string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ReferenceIndexObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type ReferenceIndexObjectType struct {
	ID         []ReferenceIndexObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ReferencedTaskObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ReferencedTaskObjectType struct {
	ID         []ReferencedTaskObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type ReferencedTaskReferenceEnumeration string

// Wrapper element containing references to specific Reference IDs.
type ReferencesRequestReferencesType struct {
	ReferenceIDReference []ReferenceIndexObjectType `xml:"urn:com.workday/bsvc Reference_ID_Reference,omitempty"`
}

// Report Background Future Process as Custom Report Data element
type ReportBackgroundFutureProcessasCustomReportDataType struct {
	CustomReportDefinitionReference *CustomReportDefinitionObjectType `xml:"urn:com.workday/bsvc Custom_Report_Definition_Reference"`
	OutputFormatReference           *UniqueIdentifierObjectType       `xml:"urn:com.workday/bsvc Output_Format_Reference,omitempty"`
	RunAsSystemUserReference        *SystemUserObjectType             `xml:"urn:com.workday/bsvc Run_As_System_User_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ReportGroupObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ReportGroupObjectType struct {
	ID         []ReportGroupObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Indicates how a report parameter is initialized
type ReportParameterInitializationDataType struct {
	XMLName                                   string                                         `xml:"urn:com.workday/bsvc XML_Name,omitempty"`
	SimpleWorkDataParameterInitializationData *SimpleWorkDataParameterInitializationDataType `xml:"urn:com.workday/bsvc Simple_Work_Data_Parameter_Initialization_Data"`
}

// Contains a unique identifier for an instance of an object.
type RepositoryDocumentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RepositoryDocumentObjectType struct {
	ID         []RepositoryDocumentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Summary of Repository Document
type RepositoryDocumentSummaryDataType struct {
	RepositoryDocumentReference *RepositoryDocumentObjectType           `xml:"urn:com.workday/bsvc Repository_Document_Reference,omitempty"`
	RepositoryDocumentData      []IntegrationRepositoryDocumentDataType `xml:"urn:com.workday/bsvc Repository_Document_Data,omitempty"`
}

// Provide additional details for the Business Process rescind.
type RescindBusinessProcessDataType struct {
	Comment               string `xml:"urn:com.workday/bsvc Comment,omitempty"`
	SuppressNotifications *bool  `xml:"urn:com.workday/bsvc Suppress_Notifications,omitempty"`
}

// Rescind Business Process Request
type RescindBusinessProcessRequestType struct {
	EventReference             *ActionEventObjectType          `xml:"urn:com.workday/bsvc Event_Reference"`
	RescindBusinessProcessData *RescindBusinessProcessDataType `xml:"urn:com.workday/bsvc Rescind_Business_Process_Data,omitempty"`
	Version                    string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Rescind Business Process Response
type RescindBusinessProcessResponseType struct {
	EventReference *ActionEventObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Parameters that let you filter the data returned in the response. You can filter returned data by dates and page attributes.
type ResponseFilterNoEntryMomentType struct {
	AsOfEffectiveDate *time.Time `xml:"urn:com.workday/bsvc As_Of_Effective_Date,omitempty"`
	AsOfEntryDateTime *time.Time `xml:"urn:com.workday/bsvc As_Of_Entry_DateTime,omitempty"`
	Page              float64    `xml:"urn:com.workday/bsvc Page,omitempty"`
	Count             float64    `xml:"urn:com.workday/bsvc Count,omitempty"`
}

func (t *ResponseFilterNoEntryMomentType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ResponseFilterNoEntryMomentType
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
func (t *ResponseFilterNoEntryMomentType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ResponseFilterNoEntryMomentType
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
type RoleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RoleObjectType struct {
	ID         []RoleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// SFTP Transport Protocol Data element
type SFTPTransportProtocolDatawithDualAuthenticationType struct {
	SFTPAddress                string                        `xml:"urn:com.workday/bsvc SFTP_Address"`
	Directory                  string                        `xml:"urn:com.workday/bsvc Directory,omitempty"`
	DualAuthentication         *bool                         `xml:"urn:com.workday/bsvc Dual_Authentication,omitempty"`
	UserID                     string                        `xml:"urn:com.workday/bsvc User_ID"`
	Password                   string                        `xml:"urn:com.workday/bsvc Password,omitempty"`
	SSHAuthenticationReference *X509PrivateKeyPairObjectType `xml:"urn:com.workday/bsvc SSH_Authentication_Reference,omitempty"`
	UseTempFile                *bool                         `xml:"urn:com.workday/bsvc Use_Temp_File,omitempty"`
	BlockSize                  float64                       `xml:"urn:com.workday/bsvc Block_Size,omitempty"`
	BlockSizeName              string                        `xml:"urn:com.workday/bsvc Block_Size_Name,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SFTPTransportProtocolObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SFTPTransportProtocolObjectType struct {
	ID         []SFTPTransportProtocolObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SecurityGroupObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SecurityGroupObjectType struct {
	ID         []SecurityGroupObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SequenceGeneratorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SequenceGeneratorObjectType struct {
	ID         []SequenceGeneratorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Sequence Generator Request References
type SequenceGeneratorRequestReferencesType struct {
	SequenceGeneratorReference []SequenceGeneratorObjectType `xml:"urn:com.workday/bsvc Sequence_Generator_Reference"`
}

// Sequence Generator Response Data
type SequenceGeneratorResponseDataType struct {
	SequenceGenerator []SequenceGeneratorType `xml:"urn:com.workday/bsvc Sequence_Generator,omitempty"`
}

// Sequence Generator
type SequenceGeneratorType struct {
	SequenceGeneratorReference *SequenceGeneratorObjectType       `xml:"urn:com.workday/bsvc Sequence_Generator_Reference,omitempty"`
	SequenceGeneratorData      *AbstractSequenceGeneratorDataType `xml:"urn:com.workday/bsvc Sequence_Generator_Data,omitempty"`
	SequencedValue             string                             `xml:"urn:com.workday/bsvc Sequenced_Value,omitempty"`
}

// Simple Work Data Parameter Initialization Data element
type SimpleWorkDataParameterInitializationDataType struct {
	SimpleWorkDataReference     *UniqueIdentifierObjectType      `xml:"urn:com.workday/bsvc Simple_Work_Data_Reference,omitempty"`
	ParameterInitializationData *ParameterInitializationDataType `xml:"urn:com.workday/bsvc Parameter_Initialization_Data"`
}

// Parameter Name
type SimpleWorkDataRuntimeParameterNameType struct {
	Value string `xml:",chardata"`
	Label string `xml:"urn:com.workday/bsvc Label,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SubscriberObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SubscriberObjectType struct {
	ID         []SubscriberObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Subscription Data element
type SubscriptionDataType struct {
	SubscriberReference                    *SubscriberObjectType                           `xml:"urn:com.workday/bsvc Subscriber_Reference"`
	EndpointInfoData                       []ExternalEndpointDataType                      `xml:"urn:com.workday/bsvc Endpoint_Info_Data,omitempty"`
	RunAsUserReference                     *SystemUserObjectType                           `xml:"urn:com.workday/bsvc Run_As_User_Reference,omitempty"`
	IntegrationBackgroundFutureProcessData *IntegrationBackgroundProcessDefinitionDataType `xml:"urn:com.workday/bsvc Integration_Background_Future_Process_Data,omitempty"`
	SubscribetoallTransactionTypes         *bool                                           `xml:"urn:com.workday/bsvc Subscribe_to_all_Transaction_Types,omitempty"`
	ExcludedTransactionLogTypeReference    []TransactionLogTypeObjectType                  `xml:"urn:com.workday/bsvc Excluded_Transaction_Log_Type_Reference,omitempty"`
	SubscribetoallBusinessProcesses        *bool                                           `xml:"urn:com.workday/bsvc Subscribe_to_all_Business_Processes,omitempty"`
	ExcludedBusinessProcessTypeReference   []BusinessProcessTypeObjectType                 `xml:"urn:com.workday/bsvc Excluded_Business_Process_Type_Reference,omitempty"`
	IncludedTransactionLogTypeReference    []TransactionLogTypeObjectType                  `xml:"urn:com.workday/bsvc Included_Transaction_Log_Type_Reference,omitempty"`
}

// Subscription Data element
type SubscriptionDataWWSType struct {
	EndpointInfoData                                     []ExternalEndpointDataType                       `xml:"urn:com.workday/bsvc Endpoint_Info_Data,omitempty"`
	SystemUsertoFireIntegrationReference                 *SystemUserObjectType                            `xml:"urn:com.workday/bsvc System_User_to_Fire_Integration_Reference,omitempty"`
	IntegrationBackgroundFutureProcessData               []IntegrationBackgroundProcessDefinitionDataType `xml:"urn:com.workday/bsvc Integration_Background_Future_Process_Data,omitempty"`
	SubscribetoallTransactionTypes                       *bool                                            `xml:"urn:com.workday/bsvc Subscribe_to_all_Transaction_Types,omitempty"`
	TransactionLogTypeExcludedfromSubscriptionReference  []TransactionLogTypeObjectType                   `xml:"urn:com.workday/bsvc Transaction_Log_Type_Excluded_from_Subscription_Reference,omitempty"`
	SubscribetoallBusinessProcesses                      *bool                                            `xml:"urn:com.workday/bsvc Subscribe_to_all_Business_Processes,omitempty"`
	BusinessProcessTypeExcludedFromSubscriptionReference []BusinessProcessTypeObjectType                  `xml:"urn:com.workday/bsvc Business_Process_Type_Excluded_From_Subscription_Reference,omitempty"`
	SpecificTransactionTypeforSuscriptionReference       []TransactionLogTypeObjectType                   `xml:"urn:com.workday/bsvc Specific_Transaction_Type_for_Suscription_Reference,omitempty"`
}

// Criteria for filtering the Subscriptions that get returned in the response
type SubscriptionRequestCriteriaType struct {
	IntegrationSystemReference *IntegrationSystemAuditedObjectType `xml:"urn:com.workday/bsvc Integration_System_Reference,omitempty"`
}

// Subscription Request References element
type SubscriptionRequestReferencesType struct {
	SubscriptionReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Subscription_Reference"`
}

// Subscription Response Data element
type SubscriptionResponseDataType struct {
	Subscription []SubscriptionType `xml:"urn:com.workday/bsvc Subscription,omitempty"`
}

// Subscription element
type SubscriptionType struct {
	SubscriptionReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Subscription_Reference,omitempty"`
	SubscriptionData      *SubscriptionDataType       `xml:"urn:com.workday/bsvc Subscription_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SystemAccountObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SystemAccountObjectType struct {
	ID         []SystemAccountObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SystemUserObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SystemUserObjectType struct {
	ID         []SystemUserObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TemplateModelObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TemplateModelObjectType struct {
	ID         []TemplateModelObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Text
type TextAttributeType struct {
	Value     string `xml:",chardata"`
	Sensitive bool   `xml:"urn:com.workday/bsvc Sensitive,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TimeZoneObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimeZoneObjectType struct {
	ID         []TimeZoneObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ToDoObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ToDoObjectType struct {
	ID         []ToDoObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type ToDoReferenceEnumeration string

// Contains a unique identifier for an instance of an object.
type TransactionLogTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TransactionLogTypeObjectType struct {
	ID         []TransactionLogTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type UniqueIdentifierObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type UniqueIdentifierObjectType struct {
	ID         []UniqueIdentifierObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type ValidationErrorType struct {
	Message       string `xml:"urn:com.workday/bsvc Message,omitempty"`
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
	Xpath         string `xml:"urn:com.workday/bsvc Xpath,omitempty"`
}

type ValidationFaultType struct {
	ValidationError []ValidationErrorType `xml:"urn:com.workday/bsvc Validation_Error,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WebServiceAPIVersionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WebServiceAPIVersionObjectType struct {
	ID         []WebServiceAPIVersionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WebServiceBackgroundProcessRuntimeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WebServiceBackgroundProcessRuntimeObjectType struct {
	ID         []WebServiceBackgroundProcessRuntimeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WebServiceDataSourceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WebServiceDataSourceObjectType struct {
	ID         []WebServiceDataSourceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WebServiceInvocationTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WebServiceInvocationTypeObjectType struct {
	ID         []WebServiceInvocationTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Container element for definitions of Web Service Operations associated to Studio integrations
type WebServiceOperationDataWWSType struct {
	Order                           string                         `xml:"urn:com.workday/bsvc Order"`
	WebServiceOperationReference    *WebServiceOperationObjectType `xml:"urn:com.workday/bsvc Web_Service_Operation_Reference"`
	WebServiceOperationConnectorWID string                         `xml:"urn:com.workday/bsvc Web_Service_Operation_Connector_WID,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WebServiceOperationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WebServiceOperationObjectType struct {
	ID         []WebServiceOperationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Reference element representing a unique instance of Web Service Operation.
type WebServiceOperationReferenceDataType struct {
	WebServiceName          string `xml:"urn:com.workday/bsvc Web_Service_Name"`
	WebServiceOperationName string `xml:"urn:com.workday/bsvc Web_Service_Operation_Name"`
}

// A valid instance of Web Service Security Configuration
type WebServiceSecurityConfigurationDataType struct {
	Enablex509TokenAuthentication      *bool                    `xml:"urn:com.workday/bsvc Enable_x509_Token_Authentication,omitempty"`
	X509PublicKeyReference             *X509PublicKeyObjectType `xml:"urn:com.workday/bsvc X509_Public_Key_Reference,omitempty"`
	EnableSAMLAuthentication           *bool                    `xml:"urn:com.workday/bsvc Enable_SAML_Authentication,omitempty"`
	SAMLIssuer                         string                   `xml:"urn:com.workday/bsvc SAML_Issuer,omitempty"`
	IdentityProviderPublicKeyReference *X509PublicKeyObjectType `xml:"urn:com.workday/bsvc Identity_Provider_Public_Key_Reference,omitempty"`
	HolderofKeyPublicKeyReference      *X509PublicKeyObjectType `xml:"urn:com.workday/bsvc Holder-of-Key_Public_Key_Reference,omitempty"`
}

// Workday Attachment Transport Protocol Data element
type WorkdayAttachmentTransportProtocolDataType struct {
	AttachToWorkday bool `xml:"urn:com.workday/bsvc Attach_To_Workday"`
}

type WorkdayCommonHeaderType struct {
	IncludeReferenceDescriptorsInResponse *bool `xml:"urn:com.workday/bsvc Include_Reference_Descriptors_In_Response,omitempty"`
}

// Workday Web Service Transport Protocol Data element
type WorkdayWebServiceTransportProtocolDataType struct {
	WebServiceOperationReferenceData *WebServiceOperationReferenceDataType `xml:"urn:com.workday/bsvc Web_Service_Operation_Reference_Data"`
}

// Contains a unique identifier for an instance of an object.
type WorkerObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkerObjectType struct {
	ID         []WorkerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WorkflowProcessParticipantObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkflowProcessParticipantObjectType struct {
	ID         []WorkflowProcessParticipantObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type WorkflowProcessParticipantReferenceEnumeration string

// A scheduled future process for a report group definition
type WorkflowReportGroupStepBackgroundFutureProcessDefinitionDataType struct {
	ReportGroupReference              *ReportGroupObjectType                  `xml:"urn:com.workday/bsvc Report_Group_Reference"`
	ReportParameterInitializationData []ReportParameterInitializationDataType `xml:"urn:com.workday/bsvc Report_Parameter_Initialization_Data,omitempty"`
	DoNotOutputEmptyReport            *bool                                   `xml:"urn:com.workday/bsvc Do_Not_Output_Empty_Report,omitempty"`
	FileExpirationinDays              float64                                 `xml:"urn:com.workday/bsvc File_Expiration_in_Days"`
}

// Some details about a step in a Business Process (Workflow)
type WorkflowStepContentDataType struct {
	Order                                      string                                                            `xml:"urn:com.workday/bsvc Order,omitempty"`
	WorkflowStepTypeReference                  *WorkflowStepTypeObjectType                                       `xml:"urn:com.workday/bsvc Workflow_Step_Type_Reference,omitempty"`
	TaskReference                              *ReferencedTaskObjectType                                         `xml:"urn:com.workday/bsvc Task_Reference,omitempty"`
	ChecklistReference                         *ChecklistObjectType                                              `xml:"urn:com.workday/bsvc Checklist_Reference,omitempty"`
	ReportBackgroundProcessDefinitionData      *WorkflowStepReportBackgroundProcessDefinitionDataType            `xml:"urn:com.workday/bsvc Report_Background_Process_Definition_Data,omitempty"`
	ReportGroupBackgroundProcessDefinitionData *WorkflowReportGroupStepBackgroundFutureProcessDefinitionDataType `xml:"urn:com.workday/bsvc Report_Group_Background_Process_Definition_Data,omitempty"`
	PartitionedBackgroundProcessReference      *PartitionedBackgroundProcessObjectType                           `xml:"urn:com.workday/bsvc Partitioned_Background_Process_Reference,omitempty"`
	ToDoReference                              *ToDoObjectType                                                   `xml:"urn:com.workday/bsvc To_Do_Reference,omitempty"`
	EventServiceReference                      *EventServiceObjectType                                           `xml:"urn:com.workday/bsvc Event_Service_Reference,omitempty"`
	IntegrationBackgroundProcessDefinitionData *IntegrationBackgroundProcessDefinitionDataType                   `xml:"urn:com.workday/bsvc Integration_Background_Process_Definition_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WorkflowStepObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkflowStepObjectType struct {
	ID         []WorkflowStepObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// A scheduled future process for a custom report definition or a report task.
type WorkflowStepReportBackgroundProcessDefinitionDataType struct {
	DocumentTypeReference                *DocumentTypeObjectType                 `xml:"urn:com.workday/bsvc Document_Type_Reference"`
	CustomBusinessFormLayoutReference    *CustomBusinessFormLayoutObjectType     `xml:"urn:com.workday/bsvc Custom_Business_Form_Layout_Reference,omitempty"`
	DeliveredBusinessFormLayoutReference *DeliveredBusinessFormLayoutObjectType  `xml:"urn:com.workday/bsvc Delivered_Business_Form_Layout_Reference,omitempty"`
	ReferencedTaskReference              *ReferencedTaskObjectType               `xml:"urn:com.workday/bsvc Referenced_Task_Reference"`
	TenantedReportDefinitionReference    *CustomReportDefinitionObjectType       `xml:"urn:com.workday/bsvc Tenanted_Report_Definition_Reference"`
	ReportParameterInitializationData    []ReportParameterInitializationDataType `xml:"urn:com.workday/bsvc Report_Parameter_Initialization_Data,omitempty"`
	FileExpirationinDays                 float64                                 `xml:"urn:com.workday/bsvc File_Expiration_in_Days"`
}

// A step in a Business Process (Workflow)
type WorkflowStepType struct {
	WorkflowStepReference *WorkflowStepObjectType       `xml:"urn:com.workday/bsvc Workflow_Step_Reference,omitempty"`
	WorkflowStepData      []WorkflowStepContentDataType `xml:"urn:com.workday/bsvc Workflow_Step_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WorkflowStepTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkflowStepTypeObjectType struct {
	ID         []WorkflowStepTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type WorkflowStepTypeReferenceEnumeration string

// Contains a unique identifier for an instance of an object.
type X509PrivateKeyPairObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type X509PrivateKeyPairObjectType struct {
	ID         []X509PrivateKeyPairObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type X509PublicKeyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type X509PublicKeyObjectType struct {
	ID         []X509PublicKeyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type XSLTAttachmentTransformationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type XSLTAttachmentTransformationObjectType struct {
	ID         []XSLTAttachmentTransformationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
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

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("15:04:05.999999999")), nil
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
