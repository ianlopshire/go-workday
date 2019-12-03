package learning

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// Contains a unique identifier for an instance of an object.
type AccountingWorktagTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AccountingWorktagTypeObjectType struct {
	ID         []AccountingWorktagTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AccountingWorktagandAggregationDimensionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AccountingWorktagandAggregationDimensionObjectType struct {
	ID         []AccountingWorktagandAggregationDimensionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains an advanced expiration rule that will apply a custom expiration date/duration to specific learning enrollment
type AdvancedExpirationRuleDataType struct {
	LearningExpirationRuleOrder string                   `xml:"urn:com.workday/bsvc Learning_Expiration_Rule_Order"`
	LearnerGroupRuleReference   *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Learner_Group_Rule_Reference"`
	ExpirationDuration          *ExpiryDurationDataType  `xml:"urn:com.workday/bsvc Expiration_Duration,omitempty"`
	ExpirationDate              *time.Time               `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
}

func (t *AdvancedExpirationRuleDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AdvancedExpirationRuleDataType
	var layout struct {
		*T
		ExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ExpirationDate = (*xsdDate)(layout.T.ExpirationDate)
	return e.EncodeElement(layout, start)
}
func (t *AdvancedExpirationRuleDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AdvancedExpirationRuleDataType
	var overlay struct {
		*T
		ExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ExpirationDate = (*xsdDate)(overlay.T.ExpirationDate)
	return d.DecodeElement(&overlay, &start)
}

// Reference to existing Allowed Related Worktags or Add new Allowed Related Worktags
type AllowedWorktagDataType struct {
	AllowedWorktagReference *AccountingWorktagandAggregationDimensionObjectType `xml:"urn:com.workday/bsvc Allowed_Worktag_Reference,omitempty"`
	DeleteAllowedValue      bool                                                `xml:"urn:com.workday/bsvc Delete_Allowed_Value,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AndOrOperatorsObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AndOrOperatorsObjectType struct {
	ID         []AndOrOperatorsObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AttendanceStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AttendanceStatusObjectType struct {
	ID         []AttendanceStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element for the attachments pertaining to a Event entered through a web service.
type BusinessProcessAttachmentDataType struct {
	FileName                         string                             `xml:"urn:com.workday/bsvc File_Name"`
	EventAttachmentDescription       string                             `xml:"urn:com.workday/bsvc Event_Attachment_Description,omitempty"`
	EventAttachmentCategoryReference *EventAttachmentCategoryObjectType `xml:"urn:com.workday/bsvc Event_Attachment_Category_Reference,omitempty"`
	File                             *[]byte                            `xml:"urn:com.workday/bsvc File,omitempty"`
	ContentType                      string                             `xml:"urn:com.workday/bsvc Content_Type,omitempty"`
}

func (t *BusinessProcessAttachmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BusinessProcessAttachmentDataType
	var layout struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	layout.T = (*T)(t)
	layout.File = (*xsdBase64Binary)(layout.T.File)
	return e.EncodeElement(layout, start)
}
func (t *BusinessProcessAttachmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BusinessProcessAttachmentDataType
	var overlay struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.File = (*xsdBase64Binary)(overlay.T.File)
	return d.DecodeElement(&overlay, &start)
}

// Captures a comment for the Business Process.
type BusinessProcessCommentDataType struct {
	Comment         string            `xml:"urn:com.workday/bsvc Comment,omitempty"`
	WorkerReference *WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
}

// Container for the processing options for a business process. If no options are submitted (or the options are submitted as 'false') then the business process is simply initiated as if it where submitted on-line with approvals, reviews, notifications and to-do's in place. If the Initiator is an Integration System User, any validations you configured on the Initiation step are ignored.
type BusinessProcessParametersType struct {
	AutoComplete                  *bool                               `xml:"urn:com.workday/bsvc Auto_Complete,omitempty"`
	RunNow                        *bool                               `xml:"urn:com.workday/bsvc Run_Now,omitempty"`
	CommentData                   *BusinessProcessCommentDataType     `xml:"urn:com.workday/bsvc Comment_Data,omitempty"`
	BusinessProcessAttachmentData []BusinessProcessAttachmentDataType `xml:"urn:com.workday/bsvc Business_Process_Attachment_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ConditionEntryOptionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ConditionEntryOptionObjectType struct {
	ID         []ConditionEntryOptionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element containing data for each Condition Item.
type ConditionItemDataWWSType struct {
	Order                         string                          `xml:"urn:com.workday/bsvc Order"`
	AndOrOperatorReference        *AndOrOperatorsObjectType       `xml:"urn:com.workday/bsvc And_Or_Operator_Reference"`
	OpenParentheses               string                          `xml:"urn:com.workday/bsvc Open_Parentheses,omitempty"`
	RelationalOperatorReference   *RelationalOperatorObjectType   `xml:"urn:com.workday/bsvc Relational_Operator_Reference"`
	ConditionEntryOptionReference *ConditionEntryOptionObjectType `xml:"urn:com.workday/bsvc Condition_Entry_Option_Reference,omitempty"`
	CloseParentheses              string                          `xml:"urn:com.workday/bsvc Close_Parentheses,omitempty"`
	SourceExternalFieldReference  *ExternalFieldObjectType        `xml:"urn:com.workday/bsvc Source_External_Field_Reference,omitempty"`
	SourceConditionRuleReference  *ConditionRuleObjectType        `xml:"urn:com.workday/bsvc Source_Condition_Rule_Reference,omitempty"`
	FilterBoolean                 *bool                           `xml:"urn:com.workday/bsvc Filter_Boolean,omitempty"`
	FilterDate                    *time.Time                      `xml:"urn:com.workday/bsvc Filter_Date,omitempty"`
	FilterDateTimeZone            *FilterDateTimeZoneDataType     `xml:"urn:com.workday/bsvc Filter_DateTimeZone,omitempty"`
	FilterTime                    *time.Time                      `xml:"urn:com.workday/bsvc Filter_Time,omitempty"`
	FilterNumber                  float64                         `xml:"urn:com.workday/bsvc Filter_Number,omitempty"`
	FilterText                    string                          `xml:"urn:com.workday/bsvc Filter_Text,omitempty"`
	TargetExternalFieldReference  *ExternalFieldObjectType        `xml:"urn:com.workday/bsvc Target_External_Field_Reference,omitempty"`
	TargetInstanceReference       []InstanceObjectType            `xml:"urn:com.workday/bsvc Target_Instance_Reference,omitempty"`
}

func (t *ConditionItemDataWWSType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ConditionItemDataWWSType
	var layout struct {
		*T
		FilterDate *xsdDate `xml:"urn:com.workday/bsvc Filter_Date,omitempty"`
		FilterTime *xsdTime `xml:"urn:com.workday/bsvc Filter_Time,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FilterDate = (*xsdDate)(layout.T.FilterDate)
	layout.FilterTime = (*xsdTime)(layout.T.FilterTime)
	return e.EncodeElement(layout, start)
}
func (t *ConditionItemDataWWSType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ConditionItemDataWWSType
	var overlay struct {
		*T
		FilterDate *xsdDate `xml:"urn:com.workday/bsvc Filter_Date,omitempty"`
		FilterTime *xsdTime `xml:"urn:com.workday/bsvc Filter_Time,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FilterDate = (*xsdDate)(overlay.T.FilterDate)
	overlay.FilterTime = (*xsdTime)(overlay.T.FilterTime)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type ConditionRuleCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ConditionRuleCategoryObjectType struct {
	ID         []ConditionRuleCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper containing data for Condition Rule.
type ConditionRuleDataWWSType struct {
	ConditionRuleID                string                            `xml:"urn:com.workday/bsvc Condition_Rule_ID,omitempty"`
	RuleDescription                string                            `xml:"urn:com.workday/bsvc Rule_Description"`
	Comment                        string                            `xml:"urn:com.workday/bsvc Comment,omitempty"`
	ConditionRuleCategoryReference []ConditionRuleCategoryObjectType `xml:"urn:com.workday/bsvc Condition_Rule_Category_Reference,omitempty"`
	ConditionItemData              []ConditionItemDataWWSType        `xml:"urn:com.workday/bsvc Condition_Item_Data"`
	CountriesReference             []CountryObjectType               `xml:"urn:com.workday/bsvc Countries_Reference,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type ContractLineTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ContractLineTypeObjectType struct {
	ID         []ContractLineTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CountryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CountryObjectType struct {
	ID         []CountryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CountryPredefinedPersonNameComponentValueObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CountryPredefinedPersonNameComponentValueObjectType struct {
	ID         []CountryPredefinedPersonNameComponentValueObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

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
type CustomValidationContextObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomValidationContextObjectType struct {
	ID         []CustomValidationContextObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Validation Condition Rule and optional Notification Message Components
type CustomValidationDataType struct {
	ConditionRuleData     *ConditionRuleDataWWSType  `xml:"urn:com.workday/bsvc Condition_Rule_Data"`
	ValidationMessageData *ValidationMessageDataType `xml:"urn:com.workday/bsvc Validation_Message_Data,omitempty"`
}

// Wrapper for data of the Custom Validation Rule.
type CustomValidationRuleDataType struct {
	CustomValidationRuleID             string                     `xml:"urn:com.workday/bsvc Custom_Validation_Rule_ID"`
	CustomValidationRuleforTransaction string                     `xml:"urn:com.workday/bsvc Custom_Validation_Rule_for_Transaction,omitempty"`
	CustomValidationRuleClassification string                     `xml:"urn:com.workday/bsvc Custom_Validation_Rule_Classification,omitempty"`
	CustomValidationData               []CustomValidationDataType `xml:"urn:com.workday/bsvc Custom_Validation_Data,omitempty"`
}

// The Request Criteria is a wrapper element around a list of elements representing the operation specific criteria needed to search for Custom Validation Rules instances.
type CustomValidationRuleRequestCriteriaType struct {
}

// Wrapper element for references to Custom Validation Rules.
type CustomValidationRuleRequestReferencesType struct {
	CustomValidationContextReference []CustomValidationContextObjectType `xml:"urn:com.workday/bsvc Custom_Validation_Context_Reference"`
}

// Wrapper containing all the Custom Validation Rule objects returned in the "Get" Response operation.
type CustomValidationRuleResponseDataType struct {
	CustomValidationRule []CustomValidationRuleType `xml:"urn:com.workday/bsvc Custom_Validation_Rule,omitempty"`
}

// The Response Group is a wrapper element around a list of elements representing the amount of data that should be included in the response.
type CustomValidationRuleResponseGroupType struct {
}

// Wrapper element for each Custom Validation Rule data in the "Get" Response operation.
type CustomValidationRuleType struct {
	CustomValidationContextReference *CustomValidationContextObjectType `xml:"urn:com.workday/bsvc Custom_Validation_Context_Reference,omitempty"`
	CustomValidationRuleData         []CustomValidationRuleDataType     `xml:"urn:com.workday/bsvc Custom_Validation_Rule_Data,omitempty"`
}

// Reference to existing Default Realted Worktag or Add new Default Related Worktag
type DefaultWorktagDataType struct {
	DefaultWorktagReference *AccountingWorktagandAggregationDimensionObjectType `xml:"urn:com.workday/bsvc Default_Worktag_Reference,omitempty"`
	DeleteDefaultValue      bool                                                `xml:"urn:com.workday/bsvc Delete_Default_Value,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DeliveryModeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DeliveryModeObjectType struct {
	ID         []DeliveryModeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DigitalCourseObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DigitalCourseObjectType struct {
	ID         []DigitalCourseObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains detailed data about the enrollment.
type EnrollInLearningCourseData2Type struct {
	LearningCourseReference *LearningRegisterableObjectType `xml:"urn:com.workday/bsvc Learning_Course_Reference"`
	LearnerReference        *RoleObjectType                 `xml:"urn:com.workday/bsvc Learner_Reference"`
}

// Contains all the request data and business process parameters required to launch the 'Enroll in Course' business process.
type EnrollInLearningCourseRequestType struct {
	BusinessProcessParameters  *BusinessProcessParametersType   `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EnrollInLearningCourseData *EnrollInLearningCourseData2Type `xml:"urn:com.workday/bsvc Enroll_In_Learning_Course_Data"`
	Version                    string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the references for the enroll in course event and the corresponding enrollment that were created.
type EnrollInLearningCourseResponseType struct {
	EventReference              *UniqueIdentifierObjectType   `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	ProposedEnrollmentReference *LearningEnrollmentObjectType `xml:"urn:com.workday/bsvc Proposed_Enrollment_Reference,omitempty"`
	Version                     string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EnrollableCourseBehaviourObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EnrollableCourseBehaviourObjectType struct {
	ID         []EnrollableCourseBehaviourObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EventAttachmentCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EventAttachmentCategoryObjectType struct {
	ID         []EventAttachmentCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a course expiry period that is relative to the date a learner successfully completes this course.
type ExpiryDurationDataType struct {
	FormatReference *LearningTimeUnitObjectType `xml:"urn:com.workday/bsvc Format_Reference"`
	Value           float64                     `xml:"urn:com.workday/bsvc Value"`
}

// Contains detailed information about an Extended Enterprise Affiliation.
type ExtendedEnterpriseAffiliationDataType struct {
	ID   string `xml:"urn:com.workday/bsvc ID,omitempty"`
	Name string `xml:"urn:com.workday/bsvc Name"`
}

// Contains a unique identifier for an instance of an object.
type ExtendedEnterpriseAffiliationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ExtendedEnterpriseAffiliationObjectType struct {
	ID         []ExtendedEnterpriseAffiliationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Retrieves specific instances of Extended Enterprise Affiliations and their associated data.
type ExtendedEnterpriseAffiliationRequestReferencesType struct {
	ExtendedEnterpriseAffiliationReference []ExtendedEnterpriseAffiliationObjectType `xml:"urn:com.workday/bsvc Extended_Enterprise_Affiliation_Reference"`
}

// Contains Extended Enterprise Affiliation data for each request reference.
type ExtendedEnterpriseAffiliationResponseDataType struct {
	ExtendedEnterpriseAffiliation []ExtendedEnterpriseAffiliationType `xml:"urn:com.workday/bsvc Extended_Enterprise_Affiliation,omitempty"`
}

// Contains Extended Enterprise Affiliation data for each request reference.
type ExtendedEnterpriseAffiliationType struct {
	ExtendedEnterpriseAffiliationReference *ExtendedEnterpriseAffiliationObjectType `xml:"urn:com.workday/bsvc Extended_Enterprise_Affiliation_Reference,omitempty"`
	ExtendedEnterpriseAffiliationData      []ExtendedEnterpriseAffiliationDataType  `xml:"urn:com.workday/bsvc Extended_Enterprise_Affiliation_Data,omitempty"`
}

// Contains detailed information about an Extended Enterprise Learner.
type ExtendedEnterpriseLearnerDataType struct {
	ID                                         string                                          `xml:"urn:com.workday/bsvc ID,omitempty"`
	LegalNameData                              *LegalNameDataType                              `xml:"urn:com.workday/bsvc Legal_Name_Data,omitempty"`
	EmailAddress                               string                                          `xml:"urn:com.workday/bsvc Email_Address,omitempty"`
	SystemUserforExtendedEnterpriseLearnerData *SystemUserforExtendedEnterpriseLearnerDataType `xml:"urn:com.workday/bsvc System_User_for_Extended_Enterprise_Learner_Data"`
	ExtendedEnterpriseAffiliationReference     *ExtendedEnterpriseAffiliationObjectType        `xml:"urn:com.workday/bsvc Extended_Enterprise_Affiliation_Reference"`
	PhotoData                                  *ExtendedEnterpriseLearnerPhotoDataType         `xml:"urn:com.workday/bsvc Photo_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ExtendedEnterpriseLearnerObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ExtendedEnterpriseLearnerObjectType struct {
	ID         []ExtendedEnterpriseLearnerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Photo Data for Extended Enterprise Learner.
type ExtendedEnterpriseLearnerPhotoDataType struct {
	FileName string  `xml:"urn:com.workday/bsvc File_Name,omitempty"`
	File     *[]byte `xml:"urn:com.workday/bsvc File,omitempty"`
}

func (t *ExtendedEnterpriseLearnerPhotoDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ExtendedEnterpriseLearnerPhotoDataType
	var layout struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	layout.T = (*T)(t)
	layout.File = (*xsdBase64Binary)(layout.T.File)
	return e.EncodeElement(layout, start)
}
func (t *ExtendedEnterpriseLearnerPhotoDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ExtendedEnterpriseLearnerPhotoDataType
	var overlay struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.File = (*xsdBase64Binary)(overlay.T.File)
	return d.DecodeElement(&overlay, &start)
}

// Retrieves specific instances of Extended Enterprise Learners and their associated data.
type ExtendedEnterpriseLearnerRequestReferencesType struct {
	ExtendedEnterpriseLearnerReference []ExtendedEnterpriseLearnerObjectType `xml:"urn:com.workday/bsvc Extended_Enterprise_Learner_Reference"`
}

// Contains Extended Enterprise Learner data for each request reference.
type ExtendedEnterpriseLearnerResponseDataType struct {
	ExtendedEnterpriseLearner []ExtendedEnterpriseLearnerType `xml:"urn:com.workday/bsvc Extended_Enterprise_Learner,omitempty"`
}

// Contains Extended Enterprise Learner data for each request reference.
type ExtendedEnterpriseLearnerType struct {
	ExtendedEnterpriseLearnerReference *ExtendedEnterpriseLearnerObjectType `xml:"urn:com.workday/bsvc Extended_Enterprise_Learner_Reference,omitempty"`
	ExtendedEnterpriseLearnerData      []ExtendedEnterpriseLearnerDataType  `xml:"urn:com.workday/bsvc Extended_Enterprise_Learner_Data,omitempty"`
}

// Contains details about external-content lessons in the course.
type ExternalContentLessonDataType struct {
	ExternalContentURL         string `xml:"urn:com.workday/bsvc External_Content_URL"`
	LearningCourseLessonTitle  string `xml:"urn:com.workday/bsvc Learning_Course_Lesson_Title"`
	ExternalContentDescription string `xml:"urn:com.workday/bsvc External_Content_Description,omitempty"`
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

// DateTimeZone value for filter.
type FilterDateTimeZoneDataType struct {
	FilterDateTime          *time.Time          `xml:"urn:com.workday/bsvc Filter_DateTime,omitempty"`
	FilterTimeZoneReference *TimeZoneObjectType `xml:"urn:com.workday/bsvc Filter_TimeZone_Reference,omitempty"`
}

func (t *FilterDateTimeZoneDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T FilterDateTimeZoneDataType
	var layout struct {
		*T
		FilterDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Filter_DateTime,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FilterDateTime = (*xsdDateTime)(layout.T.FilterDateTime)
	return e.EncodeElement(layout, start)
}
func (t *FilterDateTimeZoneDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T FilterDateTimeZoneDataType
	var overlay struct {
		*T
		FilterDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Filter_DateTime,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FilterDateTime = (*xsdDateTime)(overlay.T.FilterDateTime)
	return d.DecodeElement(&overlay, &start)
}

// Top-level element wrapper for Get Custom Validation Rules web service operation. Contains all data for the operation request.
type GetCustomValidationRulesRequestType struct {
	RequestReferences *CustomValidationRuleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CustomValidationRuleRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CustomValidationRuleResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This element is the top-level response element for all "Get" Custom Validation Rules operations.
type GetCustomValidationRulesResponseType struct {
	RequestReferences []CustomValidationRuleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   []CustomValidationRuleRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    []ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     []CustomValidationRuleResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   []ResponseResultsType                       `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []CustomValidationRuleResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to find and get Extended Enterprise Affiliations and their associated data.
type GetExtendedEnterpriseAffiliationsRequestType struct {
	RequestReferences *ExtendedEnterpriseAffiliationRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response containing data for requested Extended Enterprise Affiliation(s)
type GetExtendedEnterpriseAffiliationsResponseType struct {
	RequestReferences *ExtendedEnterpriseAffiliationRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                                `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *ExtendedEnterpriseAffiliationResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to find and get Extended Enterprise Learners and their associated data.
type GetExtendedEnterpriseLearnersRequestType struct {
	RequestReferences *ExtendedEnterpriseLearnerRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response containing data for requested Extended Enterprise Learner(s)
type GetExtendedEnterpriseLearnersResponseType struct {
	RequestReferences *ExtendedEnterpriseLearnerRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                            `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *ExtendedEnterpriseLearnerResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to find and get blended courses and their associated data.
type GetLearningBlendedCoursesRequestType struct {
	RequestReferences *LearningBlendedCourseRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningBlendedCourseRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *LearningCourseResponseGroupType            `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to blended courses and their associated data.
type GetLearningBlendedCoursesResponseType struct {
	RequestReferences *LearningBlendedCourseRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningBlendedCourseRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *LearningCourseResponseGroupType            `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                        `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *LearningCourseResponseDataType             `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to find and get course offerings and their associated data.
type GetLearningCourseOfferingsRequestType struct {
	RequestReferences *LearningCourseOfferingRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningCourseOfferingRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references of course offerings and their associated data.
type GetLearningCourseOfferingsResponseType struct {
	RequestReferences *LearningCourseOfferingRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningCourseOfferingRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                         `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *LearningCourseOfferingResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to find and get digital courses and their associated data.
type GetLearningDigitalCoursesRequestType struct {
	RequestReferences *LearningDigitalCourseRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningDigitalCourseRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *LearningCourseResponseGroupType            `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to digital courses and their associated data.
type GetLearningDigitalCoursesResponseType struct {
	RequestReferences *LearningDigitalCourseRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningDigitalCourseRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *LearningCourseResponseGroupType            `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                        `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *LearningDigitalCourseResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to find and get enrollments and their associated data.
type GetLearningEnrollmentsRequestType struct {
	RequestReferences *LearningEnrollmentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningEnrollmentRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references of enrollments and their associated data.
type GetLearningEnrollmentsResponseType struct {
	RequestReferences *LearningEnrollmentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningEnrollmentRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                     `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *LearningEnrollmentResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to find and get instructors and their associated data.
type GetLearningInstructorsRequestType struct {
	RequestReferences *LearningInstructorRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningInstructorRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to instructors and their associated data.
type GetLearningInstructorsResponseType struct {
	RequestReferences *LearningInstructorRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningInstructorRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                     `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *LearningInstructorResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to find and get lessons and their associated data.
type GetLearningLessonsRequestType struct {
	RequestReferences *LearningLessonRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningLessonRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *LearningCourseResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to lessons and their associated data.
type GetLearningLessonsResponseType struct {
	RequestReferences *LearningLessonRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningLessonRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *LearningCourseResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *LearningLessonResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to find and get other unit types and their associated data.
type GetLearningOtherUnitTypesRequestType struct {
	RequestReferences *LearningOtherUnitTypeRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningOtherUnitTypeRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to other unit types and their associated data.
type GetLearningOtherUnitTypesResponseType struct {
	RequestReferences *LearningOtherUnitTypeRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningOtherUnitTypeRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                        `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *LearningOtherUnitTypeResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to find and get learning programs and their associated data.
type GetLearningProgramsRequestType struct {
	RequestReferences *LearningProgramRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningProgramRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *LearningProgramResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to learning programs and their associated data.
type GetLearningProgramsResponseType struct {
	RequestReferences *LearningProgramRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningProgramRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *LearningProgramResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                  `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *LearningProgramResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to find and get security categories and their associated data.
type GetLearningSecurityCategoriesRequestType struct {
	RequestReferences *LearningSecurityCategoryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to security categories and their associated data.
type GetLearningSecurityCategoriesResponseType struct {
	RequestReferences *LearningSecurityCategoryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                           `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *LearningSecurityCategoryResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to find and get learning topic security segments and their associated data.
type GetLearningTopicSecuritySegmentsRequestType struct {
	RequestReferences *LearningTopicSecuritySegmentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningTopicSecuritySegmentRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to learning topic security segment and their associated data.
type GetLearningTopicSecuritySegmentsResponseType struct {
	RequestReferences *LearningTopicSecuritySegmentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningTopicSecuritySegmentRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                               `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *LearningTopicSecuritySegmentResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to find and get topics and their associated data.
type GetLearningTopicsRequestType struct {
	RequestReferences *LearningTopicRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningTopicRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references to topics and their associated data.
type GetLearningTopicsResponseType struct {
	RequestReferences *LearningTopicRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *LearningTopicRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *LearningTopicResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Sales Item data
type GetSalesItemsRequestType struct {
	RequestReferences *SalesItemRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *SalesItemRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *SalesItemResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Sales Item response elements including received criteria, references, data based on criteria and grouping, results filter and result summary
type GetSalesItemsResponseType struct {
	RequestReferences *SalesItemRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *SalesItemRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *SalesItemResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType            `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *SalesItemResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IconObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IconObjectType struct {
	ID         []IconObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ImageObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ImageObjectType struct {
	ID         []ImageObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains details of instructor-led classroom training lessons in the course.
type InstructorLedLessonDataType struct {
	Title                                           string                      `xml:"urn:com.workday/bsvc Title"`
	Description                                     string                      `xml:"urn:com.workday/bsvc Description,omitempty"`
	InstructorLedUnitTrackAttendance                *bool                       `xml:"urn:com.workday/bsvc Instructor_Led_Unit_Track_Attendance,omitempty"`
	InstructorLedUnitTrackGrades                    *bool                       `xml:"urn:com.workday/bsvc Instructor_Led_Unit_Track_Grades,omitempty"`
	LearningGradingSchemeforLearningCourseReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Learning_Grading_Scheme_for_Learning_Course_Reference,omitempty"`
}

// Contains data for an instructor led classroom training lesson in the course offering
type InstructorLedLessonOfferingDataType struct {
	Title                                             string                         `xml:"urn:com.workday/bsvc Title"`
	Description                                       string                         `xml:"urn:com.workday/bsvc Description,omitempty"`
	InstructorsReference                              []LearningInstructorObjectType `xml:"urn:com.workday/bsvc Instructors_Reference"`
	LocationReference                                 *LocationObjectType            `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	StartDate                                         time.Time                      `xml:"urn:com.workday/bsvc Start_Date"`
	EndDate                                           time.Time                      `xml:"urn:com.workday/bsvc End_Date"`
	InstructorLedUnitTrackAttendance                  *bool                          `xml:"urn:com.workday/bsvc Instructor_Led_Unit_Track_Attendance,omitempty"`
	InstructorLedUnitTrackGrades                      *bool                          `xml:"urn:com.workday/bsvc Instructor_Led_Unit_Track_Grades,omitempty"`
	LearningGradingSchemeforActivityOfferingReference *UniqueIdentifierObjectType    `xml:"urn:com.workday/bsvc Learning_Grading_Scheme_for_Activity_Offering_Reference,omitempty"`
}

func (t *InstructorLedLessonOfferingDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InstructorLedLessonOfferingDataType
	var layout struct {
		*T
		StartDate *xsdDateTime `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate   *xsdDateTime `xml:"urn:com.workday/bsvc End_Date"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDateTime)(&layout.T.StartDate)
	layout.EndDate = (*xsdDateTime)(&layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *InstructorLedLessonOfferingDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InstructorLedLessonOfferingDataType
	var overlay struct {
		*T
		StartDate *xsdDateTime `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate   *xsdDateTime `xml:"urn:com.workday/bsvc End_Date"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDateTime)(&overlay.T.StartDate)
	overlay.EndDate = (*xsdDateTime)(&overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains details of instructor-led webinar lessons in the course.
type InstructorLedWebinarLessonDataType struct {
	Title                                           string                      `xml:"urn:com.workday/bsvc Title"`
	WebinarLessonUnitTrackAttendance                *bool                       `xml:"urn:com.workday/bsvc Webinar_Lesson_Unit_Track_Attendance,omitempty"`
	WebinarLessonUnitTrackGrades                    *bool                       `xml:"urn:com.workday/bsvc Webinar_Lesson_Unit_Track_Grades,omitempty"`
	LearningGradingSchemeforLearningCourseReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Learning_Grading_Scheme_for_Learning_Course_Reference,omitempty"`
	Description                                     string                      `xml:"urn:com.workday/bsvc Description,omitempty"`
}

// Contains data for a Webinar lesson in the course offering.
type InstructorLedWebinarLessonOfferingDataType struct {
	Title                                             string                         `xml:"urn:com.workday/bsvc Title"`
	Description                                       string                         `xml:"urn:com.workday/bsvc Description,omitempty"`
	InstructorsReference                              []LearningInstructorObjectType `xml:"urn:com.workday/bsvc Instructors_Reference"`
	WebinarURL                                        string                         `xml:"urn:com.workday/bsvc Webinar_URL"`
	WebinarLoginDetails                               string                         `xml:"urn:com.workday/bsvc Webinar_Login_Details,omitempty"`
	WebinarAdditionalInformation                      string                         `xml:"urn:com.workday/bsvc Webinar_Additional_Information,omitempty"`
	StartDate                                         time.Time                      `xml:"urn:com.workday/bsvc Start_Date"`
	EndDate                                           time.Time                      `xml:"urn:com.workday/bsvc End_Date"`
	TimezoneReference                                 *TimeZoneObjectType            `xml:"urn:com.workday/bsvc Timezone_Reference"`
	WebinarLessonUnitTrackAttendance                  *bool                          `xml:"urn:com.workday/bsvc Webinar_Lesson_Unit_Track_Attendance,omitempty"`
	WebinarLessonUnitTrackGrades                      *bool                          `xml:"urn:com.workday/bsvc Webinar_Lesson_Unit_Track_Grades,omitempty"`
	LearningGradingSchemeforActivityOfferingReference *UniqueIdentifierObjectType    `xml:"urn:com.workday/bsvc Learning_Grading_Scheme_for_Activity_Offering_Reference,omitempty"`
}

func (t *InstructorLedWebinarLessonOfferingDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InstructorLedWebinarLessonOfferingDataType
	var layout struct {
		*T
		StartDate *xsdDateTime `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate   *xsdDateTime `xml:"urn:com.workday/bsvc End_Date"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDateTime)(&layout.T.StartDate)
	layout.EndDate = (*xsdDateTime)(&layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *InstructorLedWebinarLessonOfferingDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InstructorLedWebinarLessonOfferingDataType
	var overlay struct {
		*T
		StartDate *xsdDateTime `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate   *xsdDateTime `xml:"urn:com.workday/bsvc End_Date"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDateTime)(&overlay.T.StartDate)
	overlay.EndDate = (*xsdDateTime)(&overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Container for Learning Blended Course Request Criteria
type LearningBlendedCourseRequestCriteriaType struct {
}

// Retrieves specific instances of blended courses and their associated data.
type LearningBlendedCourseRequestReferencesType struct {
	LearningBlendedCourseReference []MultiCourseTemplateObjectType `xml:"urn:com.workday/bsvc Learning_Blended_Course_Reference"`
}

// Contains detailed information about a course.
type LearningCourseDataType struct {
	ID                               string                                `xml:"urn:com.workday/bsvc ID,omitempty"`
	EffectiveDate                    *time.Time                            `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	VersionLabel                     string                                `xml:"urn:com.workday/bsvc Version_Label,omitempty"`
	VersionNotes                     string                                `xml:"urn:com.workday/bsvc Version_Notes,omitempty"`
	Inactive                         *bool                                 `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	CourseTitle                      string                                `xml:"urn:com.workday/bsvc Course_Title"`
	SkillLevelReference              *LearningSkillLevelObjectType         `xml:"urn:com.workday/bsvc Skill_Level_Reference,omitempty"`
	Description                      string                                `xml:"urn:com.workday/bsvc Description"`
	CourseNumber                     string                                `xml:"urn:com.workday/bsvc Course_Number,omitempty"`
	TopicReference                   []LearningTopicObjectType             `xml:"urn:com.workday/bsvc Topic_Reference"`
	SecurityCategoryReference        []LearningSecurityCategoryObjectType  `xml:"urn:com.workday/bsvc Security_Category_Reference,omitempty"`
	RegistrableStatusReference       *LearningRegisterableStatusObjectType `xml:"urn:com.workday/bsvc Registrable_Status_Reference,omitempty"`
	OnDemand                         *bool                                 `xml:"urn:com.workday/bsvc On_Demand,omitempty"`
	MinimumEnrollmentCapacity        float64                               `xml:"urn:com.workday/bsvc Minimum_Enrollment_Capacity,omitempty"`
	MaximumEnrollmentCapacity        float64                               `xml:"urn:com.workday/bsvc Maximum_Enrollment_Capacity,omitempty"`
	WaitlistCapacity                 float64                               `xml:"urn:com.workday/bsvc Waitlist_Capacity,omitempty"`
	UnlimitedCapacity                *bool                                 `xml:"urn:com.workday/bsvc Unlimited_Capacity,omitempty"`
	ExpirationDate                   *time.Time                            `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	ExpirationDuration               *ExpiryDurationDataType               `xml:"urn:com.workday/bsvc Expiration_Duration,omitempty"`
	AdvancedExpirationRules          []AdvancedExpirationRuleDataType      `xml:"urn:com.workday/bsvc Advanced_Expiration_Rules,omitempty"`
	LearningOtherUnitTypeData        []OtherUnitTypeDataType               `xml:"urn:com.workday/bsvc Learning_Other_Unit_Type_Data,omitempty"`
	SalesItemReference               []SalesItemObjectType                 `xml:"urn:com.workday/bsvc Sales_Item_Reference,omitempty"`
	ContactPersonReference           []WorkerObjectType                    `xml:"urn:com.workday/bsvc Contact_Person_Reference,omitempty"`
	TimeValueReference               *LearningTimeUnitObjectType           `xml:"urn:com.workday/bsvc Time_Value_Reference,omitempty"`
	TotalCourseDuration              float64                               `xml:"urn:com.workday/bsvc Total_Course_Duration,omitempty"`
	EnableAutoEnrollmentfromWaitlist *bool                                 `xml:"urn:com.workday/bsvc Enable_Auto_Enrollment_from_Waitlist,omitempty"`
	CourseImageData                  *LearningImageDataType                `xml:"urn:com.workday/bsvc Course_Image_Data,omitempty"`
	LegacyCourse                     *bool                                 `xml:"urn:com.workday/bsvc Legacy_Course,omitempty"`
	AllowedInstructorReference       []LearningInstructorObjectType        `xml:"urn:com.workday/bsvc Allowed_Instructor_Reference,omitempty"`
	AllowedLocationReference         []LocationObjectType                  `xml:"urn:com.workday/bsvc Allowed_Location_Reference,omitempty"`
	ExcludefromRecommendations       *bool                                 `xml:"urn:com.workday/bsvc Exclude_from_Recommendations,omitempty"`
	CourseLessonData                 []LearningCourseLessonDataType        `xml:"urn:com.workday/bsvc Course_Lesson_Data,omitempty"`
}

func (t *LearningCourseDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LearningCourseDataType
	var layout struct {
		*T
		EffectiveDate  *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		ExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	layout.ExpirationDate = (*xsdDate)(layout.T.ExpirationDate)
	return e.EncodeElement(layout, start)
}
func (t *LearningCourseDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LearningCourseDataType
	var overlay struct {
		*T
		EffectiveDate  *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		ExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	overlay.ExpirationDate = (*xsdDate)(overlay.T.ExpirationDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains information about the lessons on the course.
type LearningCourseLessonDataType struct {
	LessonOrder                    float64                             `xml:"urn:com.workday/bsvc Lesson_Order"`
	MakeLessonMandatory            *bool                               `xml:"urn:com.workday/bsvc Make_Lesson_Mandatory,omitempty"`
	InstructorLedLessonData        *InstructorLedLessonDataType        `xml:"urn:com.workday/bsvc Instructor_Led_Lesson_Data,omitempty"`
	InstructorLedWebinarLessonData *InstructorLedWebinarLessonDataType `xml:"urn:com.workday/bsvc Instructor_Led_Webinar_Lesson_Data,omitempty"`
	ExternalContentLessonData      *ExternalContentLessonDataType      `xml:"urn:com.workday/bsvc External_Content_Lesson_Data,omitempty"`
	MediaLessonData                *MediaLessonDataType                `xml:"urn:com.workday/bsvc Media_Lesson_Data,omitempty"`
	SurveyLessonData               *SurveyLessonDataType               `xml:"urn:com.workday/bsvc Survey_Lesson_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LearningCourseObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningCourseObjectType struct {
	ID         []LearningCourseObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains data about the course offering, the blended course it is scheduled for, and the reference ID for the course offering.
type LearningCourseOfferingDataType struct {
	ID                               string                                 `xml:"urn:com.workday/bsvc ID,omitempty"`
	BlendedCourseReference           *MultiCourseTemplateObjectType         `xml:"urn:com.workday/bsvc Blended_Course_Reference"`
	LearningCourseOfferingDetailData []LearningCourseOfferingDetailDataType `xml:"urn:com.workday/bsvc Learning_Course_Offering_Detail_Data"`
}

// Contains data about the blended course for the course offering specified in the course offering reference.
type LearningCourseOfferingDatafromCourseType struct {
	Title               string                        `xml:"urn:com.workday/bsvc Title,omitempty"`
	SkillLevelReference *LearningSkillLevelObjectType `xml:"urn:com.workday/bsvc Skill_Level_Reference,omitempty"`
	Description         string                        `xml:"urn:com.workday/bsvc Description,omitempty"`
	TopicsReference     []LearningTopicObjectType     `xml:"urn:com.workday/bsvc Topics_Reference,omitempty"`
	AttachmentReference *ImageObjectType              `xml:"urn:com.workday/bsvc Attachment_Reference,omitempty"`
}

// Contains information about a course offering.
type LearningCourseOfferingDetailDataType struct {
	OfferingNumber                      string                                 `xml:"urn:com.workday/bsvc Offering_Number,omitempty"`
	RegisterableStatusReference         *LearningRegisterableStatusObjectType  `xml:"urn:com.workday/bsvc Registerable_Status_Reference,omitempty"`
	UpdateStatusfromOfferingDates       *bool                                  `xml:"urn:com.workday/bsvc Update_Status_from_Offering_Dates,omitempty"`
	AdditionalDetails                   string                                 `xml:"urn:com.workday/bsvc Additional_Details,omitempty"`
	MinEnrollmentCapacity               float64                                `xml:"urn:com.workday/bsvc Min_Enrollment_Capacity,omitempty"`
	MaxEnrollmentCapacity               float64                                `xml:"urn:com.workday/bsvc Max_Enrollment_Capacity,omitempty"`
	EnableAutoEnrollmentfromWaitlist    *bool                                  `xml:"urn:com.workday/bsvc Enable_Auto_Enrollment_from_Waitlist,omitempty"`
	WaitlistCapacity                    float64                                `xml:"urn:com.workday/bsvc Waitlist_Capacity,omitempty"`
	UnlimitedCapacity                   *bool                                  `xml:"urn:com.workday/bsvc Unlimited_Capacity,omitempty"`
	ExpirationDate                      *time.Time                             `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	ExpirationDuration                  *ExpiryDurationDataType                `xml:"urn:com.workday/bsvc Expiration_Duration,omitempty"`
	AdvancedExpirationRules             []AdvancedExpirationRuleDataType       `xml:"urn:com.workday/bsvc Advanced_Expiration_Rules,omitempty"`
	ContactPersonReference              []WorkerObjectType                     `xml:"urn:com.workday/bsvc Contact_Person_Reference,omitempty"`
	PrimaryLearningInstructorsReference []LearningInstructorObjectType         `xml:"urn:com.workday/bsvc Primary_Learning_Instructors_Reference,omitempty"`
	PrimaryLocationReference            *LocationObjectType                    `xml:"urn:com.workday/bsvc Primary_Location_Reference,omitempty"`
	VersionLabel                        string                                 `xml:"urn:com.workday/bsvc Version_Label,omitempty"`
	OtherUnitTypeValueData              []OtherUnitTypeValueDataType           `xml:"urn:com.workday/bsvc Other_Unit_Type_Value_Data,omitempty"`
	SalesItemReference                  []SalesItemObjectType                  `xml:"urn:com.workday/bsvc Sales_Item_Reference,omitempty"`
	OfferingLessonData                  []LearningCourseOfferingLessonDataType `xml:"urn:com.workday/bsvc Offering_Lesson_Data,omitempty"`
}

func (t *LearningCourseOfferingDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LearningCourseOfferingDetailDataType
	var layout struct {
		*T
		ExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ExpirationDate = (*xsdDate)(layout.T.ExpirationDate)
	return e.EncodeElement(layout, start)
}
func (t *LearningCourseOfferingDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LearningCourseOfferingDetailDataType
	var overlay struct {
		*T
		ExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ExpirationDate = (*xsdDate)(overlay.T.ExpirationDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains data for the lessons on the course offering.
type LearningCourseOfferingLessonDataType struct {
	LearningCourseLessonReference          *UniqueIdentifierObjectType                 `xml:"urn:com.workday/bsvc Learning_Course_Lesson_Reference,omitempty"`
	LessonOrder                            float64                                     `xml:"urn:com.workday/bsvc Lesson_Order"`
	RequiredLearningCourseContent          *bool                                       `xml:"urn:com.workday/bsvc Required_Learning_Course_Content,omitempty"`
	InstructorLedLessonOfferingData        *InstructorLedLessonOfferingDataType        `xml:"urn:com.workday/bsvc Instructor_Led_Lesson_Offering_Data,omitempty"`
	InstructorLedWebinarLessonOfferingData *InstructorLedWebinarLessonOfferingDataType `xml:"urn:com.workday/bsvc Instructor_Led_Webinar_Lesson_Offering_Data,omitempty"`
	ExternalContentLessonData              *ExternalContentLessonDataType              `xml:"urn:com.workday/bsvc External_Content_Lesson_Data,omitempty"`
	SurveyLessonData                       *SurveyLessonDataType                       `xml:"urn:com.workday/bsvc Survey_Lesson_Data,omitempty"`
	MediaLessonData                        *MediaLessonDataType                        `xml:"urn:com.workday/bsvc Media_Lesson_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LearningCourseOfferingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningCourseOfferingObjectType struct {
	ID         []LearningCourseOfferingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Course Offering Set Request Criteria element for Web Service.
type LearningCourseOfferingRequestCriteriaType struct {
}

// Course Offering Request References
type LearningCourseOfferingRequestReferencesType struct {
	LearningCourseOfferingReference []LearningCourseOfferingObjectType `xml:"urn:com.workday/bsvc Learning_Course_Offering_Reference"`
}

// Contains course offering data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningCourseOfferingResponseDataType struct {
	LearningCourseOffering []LearningCourseOfferingType `xml:"urn:com.workday/bsvc Learning_Course_Offering,omitempty"`
}

// Contains course offering data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningCourseOfferingType struct {
	LearningCourseOfferingReference      *LearningCourseOfferingObjectType          `xml:"urn:com.workday/bsvc Learning_Course_Offering_Reference,omitempty"`
	LearningCourseOfferingDatafromCourse []LearningCourseOfferingDatafromCourseType `xml:"urn:com.workday/bsvc Learning_Course_Offering_Data_from_Course,omitempty"`
	DeliveryModeReference                *DeliveryModeObjectType                    `xml:"urn:com.workday/bsvc Delivery_Mode_Reference,omitempty"`
	LearningCourseOfferingData           []LearningCourseOfferingDataType           `xml:"urn:com.workday/bsvc Learning_Course_Offering_Data,omitempty"`
}

// Contains blended course data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningCourseResponseDataType struct {
	LearningCourse []LearningCourseType `xml:"urn:com.workday/bsvc Learning_Course,omitempty"`
}

// Contains rules for formatting the response
type LearningCourseResponseGroupType struct {
	ExcludeCourseImageData *bool `xml:"urn:com.workday/bsvc Exclude_Course_Image_Data,omitempty"`
}

// Contains blended course data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningCourseType struct {
	LearningBlendedCourseReference *MultiCourseTemplateObjectType `xml:"urn:com.workday/bsvc Learning_Blended_Course_Reference,omitempty"`
	LearningCourseData             *LearningCourseDataType        `xml:"urn:com.workday/bsvc Learning_Course_Data,omitempty"`
}

// Container for Digital Course Request Criteria
type LearningDigitalCourseRequestCriteriaType struct {
}

// Retrieves specific instances of digital courses and their associated data.
type LearningDigitalCourseRequestReferencesType struct {
	LearningDigitalCourseReference []DigitalCourseObjectType `xml:"urn:com.workday/bsvc Learning_Digital_Course_Reference"`
}

// Contains digital course data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningDigitalCourseResponseDataType struct {
	LearningDigitalCourse []LearningDigitalCourseType `xml:"urn:com.workday/bsvc Learning_Digital_Course,omitempty"`
}

// Contains digital course data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningDigitalCourseType struct {
	LearningDigitalCourseReference *DigitalCourseObjectType `xml:"urn:com.workday/bsvc Learning_Digital_Course_Reference,omitempty"`
	LearningDigitalCourseData      *LearningCourseDataType  `xml:"urn:com.workday/bsvc Learning_Digital_Course_Data,omitempty"`
}

// Contains detailed data about the enrollment.
type LearningEnrollmentDataType struct {
	ID                               string                          `xml:"urn:com.workday/bsvc ID,omitempty"`
	LearningContentReference         *LearningRegisterableObjectType `xml:"urn:com.workday/bsvc Learning_Content_Reference"`
	LearnerReference                 *RoleObjectType                 `xml:"urn:com.workday/bsvc Learner_Reference"`
	RegisteredDate                   *time.Time                      `xml:"urn:com.workday/bsvc Registered_Date,omitempty"`
	LearningEnrollmentCompletionDate *time.Time                      `xml:"urn:com.workday/bsvc Learning_Enrollment_Completion_Date,omitempty"`
	OverallCourseScore               float64                         `xml:"urn:com.workday/bsvc Overall_Course_Score,omitempty"`
	LearningGradeReference           *LearningGradeObjectType        `xml:"urn:com.workday/bsvc Learning_Grade_Reference,omitempty"`
	AttendanceStatusReference        *AttendanceStatusObjectType     `xml:"urn:com.workday/bsvc Attendance_Status_Reference,omitempty"`
	VersionLabel                     string                          `xml:"urn:com.workday/bsvc Version_Label,omitempty"`
	ExpirationDate                   *time.Time                      `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	RescindEnrollment                *bool                           `xml:"urn:com.workday/bsvc Rescind_Enrollment,omitempty"`
}

func (t *LearningEnrollmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LearningEnrollmentDataType
	var layout struct {
		*T
		RegisteredDate                   *xsdDate     `xml:"urn:com.workday/bsvc Registered_Date,omitempty"`
		LearningEnrollmentCompletionDate *xsdDateTime `xml:"urn:com.workday/bsvc Learning_Enrollment_Completion_Date,omitempty"`
		ExpirationDate                   *xsdDate     `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.RegisteredDate = (*xsdDate)(layout.T.RegisteredDate)
	layout.LearningEnrollmentCompletionDate = (*xsdDateTime)(layout.T.LearningEnrollmentCompletionDate)
	layout.ExpirationDate = (*xsdDate)(layout.T.ExpirationDate)
	return e.EncodeElement(layout, start)
}
func (t *LearningEnrollmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LearningEnrollmentDataType
	var overlay struct {
		*T
		RegisteredDate                   *xsdDate     `xml:"urn:com.workday/bsvc Registered_Date,omitempty"`
		LearningEnrollmentCompletionDate *xsdDateTime `xml:"urn:com.workday/bsvc Learning_Enrollment_Completion_Date,omitempty"`
		ExpirationDate                   *xsdDate     `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.RegisteredDate = (*xsdDate)(overlay.T.RegisteredDate)
	overlay.LearningEnrollmentCompletionDate = (*xsdDateTime)(overlay.T.LearningEnrollmentCompletionDate)
	overlay.ExpirationDate = (*xsdDate)(overlay.T.ExpirationDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type LearningEnrollmentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningEnrollmentObjectType struct {
	ID         []LearningEnrollmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Container for Learning Enrollment Request Criteria
type LearningEnrollmentRequestCriteriaType struct {
	Imported *bool `xml:"urn:com.workday/bsvc Imported,omitempty"`
}

// Retrieves specific instances of enrollments and their associated data.
type LearningEnrollmentRequestReferencesType struct {
	LearningEnrollmentReference []LearningEnrollmentObjectType `xml:"urn:com.workday/bsvc Learning_Enrollment_Reference"`
}

// Contains enrollment data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningEnrollmentResponseDataType struct {
	LearningEnrollment []LearningEnrollmentType `xml:"urn:com.workday/bsvc Learning_Enrollment,omitempty"`
}

// Contains enrollment data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningEnrollmentType struct {
	LearningEnrollmentReference *LearningEnrollmentObjectType `xml:"urn:com.workday/bsvc Learning_Enrollment_Reference,omitempty"`
	LearningEnrollmentData      *LearningEnrollmentDataType   `xml:"urn:com.workday/bsvc Learning_Enrollment_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LearningGradeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningGradeObjectType struct {
	ID         []LearningGradeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the cover image for the course.
type LearningImageDataType struct {
	ImageID     string `xml:"urn:com.workday/bsvc Image_ID,omitempty"`
	FileName    string `xml:"urn:com.workday/bsvc File_Name"`
	ContentType string `xml:"urn:com.workday/bsvc Content_Type"`
	Image       []byte `xml:"urn:com.workday/bsvc Image"`
}

func (t *LearningImageDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LearningImageDataType
	var layout struct {
		*T
		Image *xsdBase64Binary `xml:"urn:com.workday/bsvc Image"`
	}
	layout.T = (*T)(t)
	layout.Image = (*xsdBase64Binary)(&layout.T.Image)
	return e.EncodeElement(layout, start)
}
func (t *LearningImageDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LearningImageDataType
	var overlay struct {
		*T
		Image *xsdBase64Binary `xml:"urn:com.workday/bsvc Image"`
	}
	overlay.T = (*T)(t)
	overlay.Image = (*xsdBase64Binary)(&overlay.T.Image)
	return d.DecodeElement(&overlay, &start)
}

// Contains detailed information about an instructor.
type LearningInstructorDataType struct {
	ID                         string            `xml:"urn:com.workday/bsvc ID,omitempty"`
	WorkerReference            *WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference"`
	LearningInstructorInactive *bool             `xml:"urn:com.workday/bsvc Learning_Instructor_Inactive,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LearningInstructorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningInstructorObjectType struct {
	ID         []LearningInstructorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Container for Learning Instructor Request Criteria
type LearningInstructorRequestCriteriaType struct {
}

// Retrieves specific instances of instructors and their associated data.
type LearningInstructorRequestReferencesType struct {
	LearningInstructorReference []LearningInstructorObjectType `xml:"urn:com.workday/bsvc Learning_Instructor_Reference"`
}

// Contains instructor data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningInstructorResponseDataType struct {
	LearningInstructor []LearningInstructorType `xml:"urn:com.workday/bsvc Learning_Instructor,omitempty"`
}

// Contains instructor data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningInstructorType struct {
	LearningInstructorReference *LearningInstructorObjectType `xml:"urn:com.workday/bsvc Learning_Instructor_Reference,omitempty"`
	LearningInstructorData      *LearningInstructorDataType   `xml:"urn:com.workday/bsvc Learning_Instructor_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LearningLessonBehaviourObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningLessonBehaviourObjectType struct {
	ID         []LearningLessonBehaviourObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains detailed information about the lesson.
type LearningLessonContentDataType struct {
	LearningLessonTitle        string                    `xml:"urn:com.workday/bsvc Learning_Lesson_Title"`
	LearningLessonDescription  string                    `xml:"urn:com.workday/bsvc Learning_Lesson_Description"`
	Inactive                   *bool                     `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	ExcludefromRecommendations *bool                     `xml:"urn:com.workday/bsvc Exclude_from_Recommendations,omitempty"`
	LearningTopicReference     []LearningTopicObjectType `xml:"urn:com.workday/bsvc Learning_Topic_Reference"`
	LearningMediaReference     *MediaObjectType          `xml:"urn:com.workday/bsvc Learning_Media_Reference"`
}

// Contains detailed information about a lesson.
type LearningLessonDataType struct {
	ID                        string                         `xml:"urn:com.workday/bsvc ID,omitempty"`
	LearningLessonContentData *LearningLessonContentDataType `xml:"urn:com.workday/bsvc Learning_Lesson_Content_Data"`
	LearningImageData         *LearningImageDataType         `xml:"urn:com.workday/bsvc Learning_Image_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LearningLessonObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningLessonObjectType struct {
	ID         []LearningLessonObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains request criteria to get specific lessons.
type LearningLessonRequestCriteriaType struct {
}

// Retrieves specific instances of lessons and their associated data.
type LearningLessonRequestReferencesType struct {
	LearningLessonReference []LearningLessonObjectType `xml:"urn:com.workday/bsvc Learning_Lesson_Reference"`
}

// Contains lesson data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningLessonResponseDataType struct {
	LearningLesson []LearningLessonType `xml:"urn:com.workday/bsvc Learning_Lesson,omitempty"`
}

// Contains lesson data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningLessonType struct {
	LearningLessonReference *LearningLessonObjectType `xml:"urn:com.workday/bsvc Learning_Lesson_Reference,omitempty"`
	LearningLessonData      *LearningLessonDataType   `xml:"urn:com.workday/bsvc Learning_Lesson_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LearningOtherUnitTypeAllObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningOtherUnitTypeAllObjectType struct {
	ID         []LearningOtherUnitTypeAllObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains detailed information about an other unit type.
type LearningOtherUnitTypeDataType struct {
	Name string `xml:"urn:com.workday/bsvc Name"`
}

// Contains a unique identifier for an instance of an object.
type LearningOtherUnitTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningOtherUnitTypeObjectType struct {
	ID         []LearningOtherUnitTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Documentation Container for Learning Other Unit Type Request Criteria
type LearningOtherUnitTypeRequestCriteriaType struct {
}

// Retrieves specific instances of other unit types and their associated data.
type LearningOtherUnitTypeRequestReferencesType struct {
	LearningOtherUnitTypeReference []LearningOtherUnitTypeObjectType `xml:"urn:com.workday/bsvc Learning_Other_Unit_Type_Reference"`
}

// Contains other unit type data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningOtherUnitTypeResponseDataType struct {
	LearningOtherUnitType []LearningOtherUnitTypeType `xml:"urn:com.workday/bsvc Learning_Other_Unit_Type,omitempty"`
}

// Contains other unit type data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningOtherUnitTypeType struct {
	LearningOtherUnitTypeReference *LearningOtherUnitTypeObjectType `xml:"urn:com.workday/bsvc Learning_Other_Unit_Type_Reference,omitempty"`
	LearningOtherUnitTypeData      []LearningOtherUnitTypeDataType  `xml:"urn:com.workday/bsvc Learning_Other_Unit_Type_Data,omitempty"`
}

// Contains information about the learning program content Data
type LearningProgramContentDataType struct {
	ContentOrder                 float64                              `xml:"urn:com.workday/bsvc Content_Order"`
	RequiresSuccessfulCompletion *bool                                `xml:"urn:com.workday/bsvc Requires_Successful_Completion,omitempty"`
	BlendedCourseReference       *MultiCourseTemplateObjectType       `xml:"urn:com.workday/bsvc Blended_Course_Reference"`
	DigitalCourseReference       *EnrollableCourseBehaviourObjectType `xml:"urn:com.workday/bsvc Digital_Course_Reference"`
	LessonReference              *LearningLessonBehaviourObjectType   `xml:"urn:com.workday/bsvc Lesson_Reference"`
	RecognizePreviousCompletion  *bool                                `xml:"urn:com.workday/bsvc Recognize_Previous_Completion,omitempty"`
	AllPreviousCompletions       *bool                                `xml:"urn:com.workday/bsvc All_Previous_Completions,omitempty"`
	AsOfDate                     *bool                                `xml:"urn:com.workday/bsvc As-Of_Date,omitempty"`
	WithinDateRange              *bool                                `xml:"urn:com.workday/bsvc Within_Date_Range,omitempty"`
	StartDate                    *time.Time                           `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                      *time.Time                           `xml:"urn:com.workday/bsvc End_Date,omitempty"`
}

func (t *LearningProgramContentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LearningProgramContentDataType
	var layout struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *LearningProgramContentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LearningProgramContentDataType
	var overlay struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Documentation Contains detailed information about a learning program.
type LearningProgramDataType struct {
	ID                         string                                `xml:"urn:com.workday/bsvc ID,omitempty"`
	EffectiveDate              time.Time                             `xml:"urn:com.workday/bsvc Effective_Date"`
	VersionLabel               string                                `xml:"urn:com.workday/bsvc Version_Label,omitempty"`
	VersionNotes               string                                `xml:"urn:com.workday/bsvc Version_Notes,omitempty"`
	Inactive                   *bool                                 `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	ProgramTitle               string                                `xml:"urn:com.workday/bsvc Program_Title"`
	SkillLevelReference        *LearningSkillLevelObjectType         `xml:"urn:com.workday/bsvc Skill_Level_Reference,omitempty"`
	Description                string                                `xml:"urn:com.workday/bsvc Description"`
	ProgramNumber              string                                `xml:"urn:com.workday/bsvc Program_Number,omitempty"`
	SalesItemReference         []SalesItemObjectType                 `xml:"urn:com.workday/bsvc Sales_Item_Reference,omitempty"`
	TopicReference             []LearningTopicObjectType             `xml:"urn:com.workday/bsvc Topic_Reference"`
	SecurityCategoryReference  []LearningSecurityCategoryObjectType  `xml:"urn:com.workday/bsvc Security_Category_Reference,omitempty"`
	RegistrableStatusReference *LearningRegisterableStatusObjectType `xml:"urn:com.workday/bsvc Registrable_Status_Reference"`
	ContactPersonReference     []WorkerObjectType                    `xml:"urn:com.workday/bsvc Contact_Person_Reference,omitempty"`
	ProgramImageData           *LearningImageDataType                `xml:"urn:com.workday/bsvc Program_Image_Data,omitempty"`
	ExcludefromRecommendations *bool                                 `xml:"urn:com.workday/bsvc Exclude_from_Recommendations,omitempty"`
	ProgramContentData         []LearningProgramContentDataType      `xml:"urn:com.workday/bsvc Program_Content_Data"`
}

func (t *LearningProgramDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LearningProgramDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *LearningProgramDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LearningProgramDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type LearningProgramObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningProgramObjectType struct {
	ID         []LearningProgramObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains request criteria to get specific learning programs.
type LearningProgramRequestCriteriaType struct {
}

// Retrieves specific instances of learning programs and their associated data.
type LearningProgramRequestReferencesType struct {
	LearningProgramReference []LearningProgramObjectType `xml:"urn:com.workday/bsvc Learning_Program_Reference"`
}

// Contains learning program data for each request reference. The data is current as of the dates in the response filter and doesn't include all historical information.
type LearningProgramResponseDataType struct {
	LearningProgram []LearningProgramType `xml:"urn:com.workday/bsvc Learning_Program,omitempty"`
}

// Contains rules for formatting the response.
type LearningProgramResponseGroupType struct {
	ExcludeProgramImageData *bool `xml:"urn:com.workday/bsvc Exclude_Program_Image_Data,omitempty"`
}

// Contains learning program data for each request reference. The data is current as of the dates in the response filter and doesn't include all historical information.
type LearningProgramType struct {
	LearningProgramReference *LearningProgramObjectType `xml:"urn:com.workday/bsvc Learning_Program_Reference,omitempty"`
	LearningProgramData      *LearningProgramDataType   `xml:"urn:com.workday/bsvc Learning_Program_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LearningRegisterableObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningRegisterableObjectType struct {
	ID         []LearningRegisterableObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LearningRegisterableStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningRegisterableStatusObjectType struct {
	ID         []LearningRegisterableStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains detailed information about the security category
type LearningSecurityCategoryDataType struct {
	Name string `xml:"urn:com.workday/bsvc Name"`
}

// Contains a unique identifier for an instance of an object.
type LearningSecurityCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningSecurityCategoryObjectType struct {
	ID         []LearningSecurityCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Retrieves specific instances of security categories and their associated data.
type LearningSecurityCategoryRequestReferencesType struct {
	LearningSecurityCategoryReference []LearningSecurityCategoryObjectType `xml:"urn:com.workday/bsvc Learning_Security_Category_Reference"`
}

// Contains security category data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningSecurityCategoryResponseDataType struct {
	LearningSecurityCategory []LearningSecurityCategoryType `xml:"urn:com.workday/bsvc Learning_Security_Category,omitempty"`
}

// Contains security category data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningSecurityCategoryType struct {
	LearningSecurityCategoryReference *LearningSecurityCategoryObjectType `xml:"urn:com.workday/bsvc Learning_Security_Category_Reference,omitempty"`
	LearningSecurityCategoryData      []LearningSecurityCategoryDataType  `xml:"urn:com.workday/bsvc Learning_Security_Category_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LearningSkillLevelObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningSkillLevelObjectType struct {
	ID         []LearningSkillLevelObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LearningTimeUnitObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningTimeUnitObjectType struct {
	ID         []LearningTimeUnitObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains detailed information about a topic.
type LearningTopicDataType struct {
	LearningTopicName        string                 `xml:"urn:com.workday/bsvc Learning_Topic_Name"`
	LearningTopicDescription string                 `xml:"urn:com.workday/bsvc Learning_Topic_Description,omitempty"`
	Inactive                 *bool                  `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	IconReference            *IconObjectType        `xml:"urn:com.workday/bsvc Icon_Reference,omitempty"`
	LearningImageData        *LearningImageDataType `xml:"urn:com.workday/bsvc Learning_Image_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LearningTopicObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningTopicObjectType struct {
	ID         []LearningTopicObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Request criteria for Learning Topic
type LearningTopicRequestCriteriaType struct {
}

// Retrieves specific instances of topics and their associated data.
type LearningTopicRequestReferencesType struct {
	LearningTopicReference []LearningTopicObjectType `xml:"urn:com.workday/bsvc Learning_Topic_Reference"`
}

// Contains topic data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningTopicResponseDataType struct {
	LearningTopic []LearningTopicType `xml:"urn:com.workday/bsvc Learning_Topic,omitempty"`
}

// Contains detailed information about the learning topic security segment
type LearningTopicSecuritySegmentDataType struct {
	ID                     string                    `xml:"urn:com.workday/bsvc ID"`
	LearningTopicReference []LearningTopicObjectType `xml:"urn:com.workday/bsvc Learning_Topic_Reference"`
}

// Contains a unique identifier for an instance of an object.
type LearningTopicSecuritySegmentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LearningTopicSecuritySegmentObjectType struct {
	ID         []LearningTopicSecuritySegmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Retrieves specific instances of learning topic security segments based of certain search criterion
type LearningTopicSecuritySegmentRequestCriteriaType struct {
}

// Retrieves specific instances of learning topic security segments and their associated data.
type LearningTopicSecuritySegmentRequestReferencesType struct {
	LearningTopicSecuritySegmentReference []LearningTopicSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Learning_Topic_Security_Segment_Reference"`
}

// Contains learning topic security segment data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningTopicSecuritySegmentResponseDataType struct {
	LearningTopicSecuritySegment []LearningTopicSecuritySegmentType `xml:"urn:com.workday/bsvc Learning_Topic_Security_Segment,omitempty"`
}

// Contains learning topic security segment data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningTopicSecuritySegmentType struct {
	LearningTopicSecuritySegmentReference *LearningTopicSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Learning_Topic_Security_Segment_Reference,omitempty"`
	LearningTopicSecuritySegmentData      *LearningTopicSecuritySegmentDataType   `xml:"urn:com.workday/bsvc Learning_Topic_Security_Segment_Data,omitempty"`
}

// Contains topic data for each request reference. The data is current as of the dates in the response filter, and doesn't include all historical information.
type LearningTopicType struct {
	LearningTopicReference *LearningTopicObjectType `xml:"urn:com.workday/bsvc Learning_Topic_Reference,omitempty"`
	LearningTopicData      *LearningTopicDataType   `xml:"urn:com.workday/bsvc Learning_Topic_Data,omitempty"`
}

// Contains the legal name for a person.  A person must name one and only one legal name.
type LegalNameDataType struct {
	NameDetailData *PersonNameDetailDataType `xml:"urn:com.workday/bsvc Name_Detail_Data"`
}

// Contains the components of a name in local script, such as the First Name and Last Name, for supporting countries.
type LocalPersonNameDetailDataType struct {
	FirstName          string `xml:"urn:com.workday/bsvc First_Name,omitempty"`
	MiddleName         string `xml:"urn:com.workday/bsvc Middle_Name,omitempty"`
	LastName           string `xml:"urn:com.workday/bsvc Last_Name,omitempty"`
	SecondaryLastName  string `xml:"urn:com.workday/bsvc Secondary_Last_Name,omitempty"`
	FirstName2         string `xml:"urn:com.workday/bsvc First_Name_2,omitempty"`
	MiddleName2        string `xml:"urn:com.workday/bsvc Middle_Name_2,omitempty"`
	LastName2          string `xml:"urn:com.workday/bsvc Last_Name_2,omitempty"`
	SecondaryLastName2 string `xml:"urn:com.workday/bsvc Secondary_Last_Name_2,omitempty"`
	LocalName          string `xml:"urn:com.workday/bsvc Local_Name,attr,omitempty"`
	LocalScript        string `xml:"urn:com.workday/bsvc Local_Script,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LocationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LocationObjectType struct {
	ID         []LocationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains all the request data to add or update learning courses.
type ManageLearningCourseDataType struct {
	PutLearningBlendedCourseRequest *PutLearningBlendedCoursesRequestType `xml:"urn:com.workday/bsvc Put_Learning_Blended_Course_Request"`
	PutLearningDigitalCourseRequest *PutLearningDigitalCourseRequestType  `xml:"urn:com.workday/bsvc Put_Learning_Digital_Course_Request"`
}

// Contains all the request data to add or update learning courses.
type ManageLearningCourseRequestType struct {
	ManageLearningCourseData  *ManageLearningCourseDataType  `xml:"urn:com.workday/bsvc Manage_Learning_Course_Data"`
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the references for the course event and corresponding digital course created or updated.
type ManageLearningCourseResponseType struct {
	CourseEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Course_Event_Reference,omitempty"`
	CourseReference      *LearningCourseObjectType   `xml:"urn:com.workday/bsvc Course_Reference,omitempty"`
	Version              string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update learning programs.
type ManageLearningProgramDataType struct {
	LearningProgramReference *LearningProgramObjectType `xml:"urn:com.workday/bsvc Learning_Program_Reference,omitempty"`
	LearningProgramData      *LearningProgramDataType   `xml:"urn:com.workday/bsvc Learning_Program_Data"`
}

// Contains all the request data to add or update learning programs.
type ManageLearningProgramRequestType struct {
	ManageLearningProgramData *ManageLearningProgramDataType `xml:"urn:com.workday/bsvc Manage_Learning_Program_Data"`
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the references for the learning program event and corresponding learning program created or updated.
type ManageLearningProgramResponseType struct {
	ProgramEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Program_Event_Reference,omitempty"`
	ProgramReference      *LearningProgramObjectType  `xml:"urn:com.workday/bsvc Program_Reference,omitempty"`
	Version               string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update lessons.
type ManageLessonRequestType struct {
	ManageLessonData          *PutLearningLessonsRequestType `xml:"urn:com.workday/bsvc Manage_Lesson_Data"`
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references for the lessons added or updated.
type ManageLessonResponseType struct {
	LearningLessonReference *LearningLessonObjectType   `xml:"urn:com.workday/bsvc Learning_Lesson_Reference,omitempty"`
	EventReference          *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version                 string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains details about media lessons in the course.
type MediaLessonDataType struct {
	MediaReference            *MediaObjectType `xml:"urn:com.workday/bsvc Media_Reference"`
	LearningCourseLessonTitle string           `xml:"urn:com.workday/bsvc Learning_Course_Lesson_Title"`
	MediaContentDescription   string           `xml:"urn:com.workday/bsvc Media_Content_Description,omitempty"`
	ProvideCourseGrade        *bool            `xml:"urn:com.workday/bsvc Provide_Course_Grade,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MediaObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MediaObjectType struct {
	ID         []MediaObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MultiCourseTemplateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MultiCourseTemplateObjectType struct {
	ID         []MultiCourseTemplateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element defining a component to be used in composing a message.
type NotificationMessageComponentDataType struct {
	Order                  string                   `xml:"urn:com.workday/bsvc Order"`
	Text                   string                   `xml:"urn:com.workday/bsvc Text"`
	ExternalFieldReference *ExternalFieldObjectType `xml:"urn:com.workday/bsvc External_Field_Reference"`
}

// Contains data for Other Unit Type values for the course.
type OtherUnitTypeDataType struct {
	OtherUnitTypeReference *LearningOtherUnitTypeAllObjectType `xml:"urn:com.workday/bsvc Other_Unit_Type_Reference"`
	OtherUnitValue         float64                             `xml:"urn:com.workday/bsvc Other_Unit_Value"`
}

// Contains information about Other Unit Types for the course offering.
type OtherUnitTypeValueDataType struct {
	OtherUnitTypeReference *LearningOtherUnitTypeAllObjectType `xml:"urn:com.workday/bsvc Other_Unit_Type_Reference"`
	OtherUnitValue         float64                             `xml:"urn:com.workday/bsvc Other_Unit_Value"`
}

// Contains the components of a name, such as the First Name and Last Name.
type PersonNameDetailDataType struct {
	CountryReference                *CountryObjectType             `xml:"urn:com.workday/bsvc Country_Reference"`
	PrefixData                      *PersonNamePrefixDataType      `xml:"urn:com.workday/bsvc Prefix_Data,omitempty"`
	FirstName                       string                         `xml:"urn:com.workday/bsvc First_Name,omitempty"`
	MiddleName                      string                         `xml:"urn:com.workday/bsvc Middle_Name,omitempty"`
	LastName                        string                         `xml:"urn:com.workday/bsvc Last_Name,omitempty"`
	SecondaryLastName               string                         `xml:"urn:com.workday/bsvc Secondary_Last_Name,omitempty"`
	TertiaryLastName                string                         `xml:"urn:com.workday/bsvc Tertiary_Last_Name,omitempty"`
	LocalNameDetailData             *LocalPersonNameDetailDataType `xml:"urn:com.workday/bsvc Local_Name_Detail_Data,omitempty"`
	SuffixData                      *PersonNameSuffixDataType      `xml:"urn:com.workday/bsvc Suffix_Data,omitempty"`
	FullNameforSingaporeandMalaysia string                         `xml:"urn:com.workday/bsvc Full_Name_for_Singapore_and_Malaysia,omitempty"`
	FormattedName                   string                         `xml:"urn:com.workday/bsvc Formatted_Name,attr,omitempty"`
	ReportingName                   string                         `xml:"urn:com.workday/bsvc Reporting_Name,attr,omitempty"`
}

// Contains the prefixes for a name.
type PersonNamePrefixDataType struct {
	TitleReference      *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Title_Reference,omitempty"`
	TitleDescriptor     string                                               `xml:"urn:com.workday/bsvc Title_Descriptor,omitempty"`
	SalutationReference *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Salutation_Reference,omitempty"`
}

// Contains the suffixes for a name.
type PersonNameSuffixDataType struct {
	SocialSuffixReference       *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Social_Suffix_Reference,omitempty"`
	SocialSuffixDescriptor      string                                               `xml:"urn:com.workday/bsvc Social_Suffix_Descriptor,omitempty"`
	AcademicSuffixReference     *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Academic_Suffix_Reference,omitempty"`
	HereditarySuffixReference   *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Hereditary_Suffix_Reference,omitempty"`
	HonorarySuffixReference     *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Honorary_Suffix_Reference,omitempty"`
	ProfessionalSuffixReference *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Professional_Suffix_Reference,omitempty"`
	ReligiousSuffixReference    *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Religious_Suffix_Reference,omitempty"`
	RoyalSuffixReference        *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Royal_Suffix_Reference,omitempty"`
}

type ProcessingFaultType struct {
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProjectTemplateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProjectTemplateObjectType struct {
	ID         []ProjectTemplateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PurchaseItemObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PurchaseItemObjectType struct {
	ID         []PurchaseItemObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Root element for the Request on the Put operation.
type PutCustomValidationRuleRequestType struct {
	CustomValidationRuleReference *CustomValidationContextObjectType `xml:"urn:com.workday/bsvc Custom_Validation_Rule_Reference,omitempty"`
	CustomValidationRuleData      *CustomValidationRuleDataType      `xml:"urn:com.workday/bsvc Custom_Validation_Rule_Data"`
	AddOnly                       bool                               `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Root element for the Response on the Put operation.
type PutCustomValidationRuleResponseType struct {
	CustomValidationRuleReference *CustomValidationContextObjectType `xml:"urn:com.workday/bsvc Custom_Validation_Rule_Reference,omitempty"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update an Extended Enterprise Affiliation.
type PutExtendedEnterpriseAffiliationRequestType struct {
	ExtendedEnterpriseAffiliationReference *ExtendedEnterpriseAffiliationObjectType `xml:"urn:com.workday/bsvc Extended_Enterprise_Affiliation_Reference,omitempty"`
	ExtendedEnterpriseAffiliationData      *ExtendedEnterpriseAffiliationDataType   `xml:"urn:com.workday/bsvc Extended_Enterprise_Affiliation_Data,omitempty"`
	Version                                string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains reference for the Extended Enterprise Affiliation added or updated.
type PutExtendedEnterpriseAffiliationResponseType struct {
	ExtendedEnterpriseAffiliationReference *ExtendedEnterpriseAffiliationObjectType `xml:"urn:com.workday/bsvc Extended_Enterprise_Affiliation_Reference,omitempty"`
	Version                                string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update Extended Enterprise Learners.
type PutExtendedEnterpriseLearnerRequestType struct {
	ExtendedEnterpriseLearnerReference *ExtendedEnterpriseLearnerObjectType `xml:"urn:com.workday/bsvc Extended_Enterprise_Learner_Reference,omitempty"`
	ExtendedEnterpriseLearnerData      *ExtendedEnterpriseLearnerDataType   `xml:"urn:com.workday/bsvc Extended_Enterprise_Learner_Data"`
	Version                            string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references for Extended Enterprise Learners added or updated.
type PutExtendedEnterpriseLearnerResponseType struct {
	ExtendedEnterpriseLearnerReference *ExtendedEnterpriseLearnerObjectType `xml:"urn:com.workday/bsvc Extended_Enterprise_Learner_Reference,omitempty"`
	Version                            string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update blended courses.
type PutLearningBlendedCoursesRequestType struct {
	LearningBlendedCourseReference *MultiCourseTemplateObjectType `xml:"urn:com.workday/bsvc Learning_Blended_Course_Reference,omitempty"`
	LearningCourseData             *LearningCourseDataType        `xml:"urn:com.workday/bsvc Learning_Course_Data"`
}

// Contains all the request data to add or update a course offering.
type PutLearningCourseOfferingRequestType struct {
	LearningCourseOfferingReference *LearningCourseOfferingObjectType `xml:"urn:com.workday/bsvc Learning_Course_Offering_Reference,omitempty"`
	LearningCourseOfferingData      []LearningCourseOfferingDataType  `xml:"urn:com.workday/bsvc Learning_Course_Offering_Data"`
	Version                         string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references for the course offerings created or updated.
type PutLearningCourseOfferingResponseType struct {
	LearningCourseOfferingReference *LearningCourseOfferingObjectType `xml:"urn:com.workday/bsvc Learning_Course_Offering_Reference,omitempty"`
	Version                         string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update a digital course.
type PutLearningDigitalCourseRequestType struct {
	LearningDigitalCourseReference *DigitalCourseObjectType `xml:"urn:com.workday/bsvc Learning_Digital_Course_Reference,omitempty"`
	LearningDigitalCourseData      *LearningCourseDataType  `xml:"urn:com.workday/bsvc Learning_Digital_Course_Data"`
}

// Contains all the request data to add or update legacy enrollments.
type PutLearningEnrollmentRequestType struct {
	LearningEnrollmentReference *LearningEnrollmentObjectType `xml:"urn:com.workday/bsvc Learning_Enrollment_Reference,omitempty"`
	LearningEnrollmentData      *LearningEnrollmentDataType   `xml:"urn:com.workday/bsvc Learning_Enrollment_Data"`
	Version                     string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the references for the enrollments created or updated.
type PutLearningEnrollmentResponseType struct {
	LearningEnrollmentReference *LearningEnrollmentObjectType `xml:"urn:com.workday/bsvc Learning_Enrollment_Reference,omitempty"`
	Version                     string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update instructors.
type PutLearningInstructorsRequestType struct {
	LearningInstructorReference *LearningInstructorObjectType `xml:"urn:com.workday/bsvc Learning_Instructor_Reference,omitempty"`
	LearningInstructorData      *LearningInstructorDataType   `xml:"urn:com.workday/bsvc Learning_Instructor_Data"`
	Version                     string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references for instructors added or updated.
type PutLearningInstructorsResponseType struct {
	LearningInstructorReference *LearningInstructorObjectType `xml:"urn:com.workday/bsvc Learning_Instructor_Reference,omitempty"`
	Version                     string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update lessons.
type PutLearningLessonsRequestType struct {
	LearningLessonReference *LearningLessonObjectType `xml:"urn:com.workday/bsvc Learning_Lesson_Reference,omitempty"`
	LearningLessonData      *LearningLessonDataType   `xml:"urn:com.workday/bsvc Learning_Lesson_Data,omitempty"`
}

// Contains all the request data to add or update other unit types.
type PutLearningOtherUnitTypeRequestType struct {
	LearningOtherUnitTypeReference *LearningOtherUnitTypeObjectType `xml:"urn:com.workday/bsvc Learning_Other_Unit_Type_Reference,omitempty"`
	LearningOtherUnitTypeData      *LearningOtherUnitTypeDataType   `xml:"urn:com.workday/bsvc Learning_Other_Unit_Type_Data"`
	Version                        string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the references for the other unit types added or updated.
type PutLearningOtherUnitTypeResponseType struct {
	LearningOtherUnitTypeReference *LearningOtherUnitTypeObjectType `xml:"urn:com.workday/bsvc Learning_Other_Unit_Type_Reference,omitempty"`
	Version                        string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update security categories
type PutLearningSecurityCategoryRequestType struct {
	LearningSecurityCategoryReference *LearningSecurityCategoryObjectType `xml:"urn:com.workday/bsvc Learning_Security_Category_Reference,omitempty"`
	LearningSecurityCategoryData      *LearningSecurityCategoryDataType   `xml:"urn:com.workday/bsvc Learning_Security_Category_Data"`
	Version                           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the security category to be created or updated
type PutLearningSecurityCategoryResponseType struct {
	LearningSecurityCategoryReference *LearningSecurityCategoryObjectType `xml:"urn:com.workday/bsvc Learning_Security_Category_Reference,omitempty"`
	Version                           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update topics.
type PutLearningTopicRequestType struct {
	LearningTopicReference *LearningTopicObjectType `xml:"urn:com.workday/bsvc Learning_Topic_Reference,omitempty"`
	LearningTopicData      *LearningTopicDataType   `xml:"urn:com.workday/bsvc Learning_Topic_Data"`
	Version                string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains references for the topics added or updated.
type PutLearningTopicResponseType struct {
	LearningTopicReference *LearningTopicObjectType `xml:"urn:com.workday/bsvc Learning_Topic_Reference,omitempty"`
	Version                string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update learning topic security segments
type PutLearningTopicSecuritySegmentRequestType struct {
	LearningTopicSecuritySegmentReference *LearningTopicSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Learning_Topic_Security_Segment_Reference,omitempty"`
	LearningTopicSecuritySegmentData      *LearningTopicSecuritySegmentDataType   `xml:"urn:com.workday/bsvc Learning_Topic_Security_Segment_Data"`
	Version                               string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the Learning Topic Security Segment to be created or updated
type PutLearningTopicSecuritySegmentResponseType struct {
	LearningTopicSecuritySegmentReference *LearningTopicSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Learning_Topic_Security_Segment_Reference,omitempty"`
	Version                               string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Reference to the existing Allowed or Default Related Worktags.
type RelatedWorktagWidgetDataType struct {
	RelatedWorktagsbyTypeData []RelatedWorktagsbyWorktagTypeDataType `xml:"urn:com.workday/bsvc Related_Worktags_by_Type_Data,omitempty"`
	ReplaceAll                bool                                   `xml:"urn:com.workday/bsvc Replace_All,attr,omitempty"`
}

// Allowed and Default Related Worktags Data by Worktag Type.
type RelatedWorktagsbyWorktagTypeDataType struct {
	WorktagTypeReference               *AccountingWorktagTypeObjectType `xml:"urn:com.workday/bsvc Worktag_Type_Reference,omitempty"`
	RequiredOnTransaction              *bool                            `xml:"urn:com.workday/bsvc Required_On_Transaction,omitempty"`
	RequiredOnTransactionForValidation *bool                            `xml:"urn:com.workday/bsvc Required_On_Transaction_For_Validation,omitempty"`
	DefaultWorktagData                 *DefaultWorktagDataType          `xml:"urn:com.workday/bsvc Default_Worktag_Data,omitempty"`
	ReplaceAllAllowedValues            *bool                            `xml:"urn:com.workday/bsvc Replace_All_Allowed_Values,omitempty"`
	AllowedWorktagData                 []AllowedWorktagDataType         `xml:"urn:com.workday/bsvc Allowed_Worktag_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RelationalOperatorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RelationalOperatorObjectType struct {
	ID         []RelationalOperatorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type RevenueCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RevenueCategoryObjectType struct {
	ID         []RevenueCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RevenueRecognitionScheduleTemplateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RevenueRecognitionScheduleTemplateObjectType struct {
	ID         []RevenueRecognitionScheduleTemplateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Element containing all Sales Item data
type SalesItemDataType struct {
	SalesItemID                                 string                                        `xml:"urn:com.workday/bsvc Sales_Item_ID,omitempty"`
	ItemName                                    string                                        `xml:"urn:com.workday/bsvc Item_Name"`
	AlternateName                               string                                        `xml:"urn:com.workday/bsvc Alternate_Name,omitempty"`
	Inactive                                    *bool                                         `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	WorktagOnly                                 *bool                                         `xml:"urn:com.workday/bsvc Worktag_Only,omitempty"`
	RevenueCategoryReference                    *RevenueCategoryObjectType                    `xml:"urn:com.workday/bsvc Revenue_Category_Reference,omitempty"`
	ItemDescription                             string                                        `xml:"urn:com.workday/bsvc Item_Description,omitempty"`
	ItemIdentifier                              string                                        `xml:"urn:com.workday/bsvc Item_Identifier,omitempty"`
	UnitofMeasureReference                      *UnitofMeasureObjectType                      `xml:"urn:com.workday/bsvc Unit_of_Measure_Reference,omitempty"`
	UnitofMeasure2Reference                     *UnitofMeasureObjectType                      `xml:"urn:com.workday/bsvc Unit_of_Measure_2_Reference,omitempty"`
	EnableBulkQuantityPricing                   *bool                                         `xml:"urn:com.workday/bsvc Enable_Bulk_Quantity_Pricing,omitempty"`
	ItemUnitPrice                               float64                                       `xml:"urn:com.workday/bsvc Item_Unit_Price,omitempty"`
	CurrencyReference                           *CurrencyObjectType                           `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	UOMPricedatQuantitiesof                     float64                                       `xml:"urn:com.workday/bsvc UOM_Priced_at_Quantities_of,omitempty"`
	TaxApplicabilityReference                   *TaxApplicabilityObjectType                   `xml:"urn:com.workday/bsvc Tax_Applicability_Reference,omitempty"`
	SalesItemGroupReference                     []SalesItemGroupObjectType                    `xml:"urn:com.workday/bsvc Sales_Item_Group_Reference,omitempty"`
	ProjectTemplateReference                    *ProjectTemplateObjectType                    `xml:"urn:com.workday/bsvc Project_Template_Reference,omitempty"`
	DeferredRevenue                             *bool                                         `xml:"urn:com.workday/bsvc Deferred_Revenue,omitempty"`
	RevenueRecognitionScheduleTemplateReference *RevenueRecognitionScheduleTemplateObjectType `xml:"urn:com.workday/bsvc Revenue_Recognition_Schedule_Template_Reference,omitempty"`
	RevenueRecognitionMethodReference           *ScheduleDistributionMethodObjectType         `xml:"urn:com.workday/bsvc Revenue_Recognition_Method_Reference,omitempty"`
	VSOEPrice                                   float64                                       `xml:"urn:com.workday/bsvc VSOE_Price,omitempty"`
	ThirdPartyEvidence                          float64                                       `xml:"urn:com.workday/bsvc Third_Party_Evidence,omitempty"`
	EstimatedSellingPrice                       float64                                       `xml:"urn:com.workday/bsvc Estimated_Selling_Price,omitempty"`
	PriceTypeReference                          *ContractLineTypeObjectType                   `xml:"urn:com.workday/bsvc Price_Type_Reference,omitempty"`
	RelatedWorktagsData                         *RelatedWorktagWidgetDataType                 `xml:"urn:com.workday/bsvc Related_Worktags_Data,omitempty"`
	MapstoPurchaseItemReference                 *PurchaseItemObjectType                       `xml:"urn:com.workday/bsvc Maps_to_Purchase_Item_Reference,omitempty"`
	IsaBundle                                   *bool                                         `xml:"urn:com.workday/bsvc Is_a_Bundle,omitempty"`
	SalesItemToIncludeInBundleReference         []SalesItemObjectType                         `xml:"urn:com.workday/bsvc Sales_Item_To_Include_In_Bundle_Reference,omitempty"`
	SalesItemBundleReference                    []SalesItemObjectType                         `xml:"urn:com.workday/bsvc Sales_Item_Bundle_Reference,omitempty"`
	FulfillmentRequired                         *bool                                         `xml:"urn:com.workday/bsvc Fulfillment_Required,omitempty"`
	Renewable                                   *bool                                         `xml:"urn:com.workday/bsvc Renewable,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SalesItemGroupObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SalesItemGroupObjectType struct {
	ID         []SalesItemGroupObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SalesItemObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SalesItemObjectType struct {
	ID         []SalesItemObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing Sales Item request criteria
type SalesItemRequestCriteriaType struct {
}

// Element containing reference instances for a Sales Item
type SalesItemRequestReferencesType struct {
	SalesItemReference []SalesItemObjectType `xml:"urn:com.workday/bsvc Sales_Item_Reference"`
}

// Wrapper Element that includes Sales Item Instance and Data
type SalesItemResponseDataType struct {
	SalesItem []SalesItemType `xml:"urn:com.workday/bsvc Sales_Item,omitempty"`
}

// Element containing Sales Item response grouping options
type SalesItemResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Wrapper Element that includes Sales Item Instance and Data
type SalesItemType struct {
	SalesItemReference *SalesItemObjectType `xml:"urn:com.workday/bsvc Sales_Item_Reference,omitempty"`
	SalesItemData      *SalesItemDataType   `xml:"urn:com.workday/bsvc Sales_Item_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ScheduleDistributionMethodObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ScheduleDistributionMethodObjectType struct {
	ID         []ScheduleDistributionMethodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains details about the survey lessons in the course.
type SurveyLessonDataType struct {
	SurveyReference *SurveyObjectType `xml:"urn:com.workday/bsvc Survey_Reference"`
	Description     string            `xml:"urn:com.workday/bsvc Description,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SurveyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SurveyObjectType struct {
	ID         []SurveyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Information related to system account for Extended Enterprise Learner.
type SystemUserforExtendedEnterpriseLearnerDataType struct {
	SystemUserNameforPersoninRole    string     `xml:"urn:com.workday/bsvc System_User_Name_for_Person_in_Role,omitempty"`
	AccountExpirationDate            *time.Time `xml:"urn:com.workday/bsvc Account_Expiration_Date,omitempty"`
	AccountDisabled                  *bool      `xml:"urn:com.workday/bsvc Account_Disabled,omitempty"`
	SendEmailNotificationforUsername *bool      `xml:"urn:com.workday/bsvc Send_Email_Notification_for_Username,omitempty"`
	SendEmailNotificationforPassword *bool      `xml:"urn:com.workday/bsvc Send_Email_Notification_for_Password,omitempty"`
}

func (t *SystemUserforExtendedEnterpriseLearnerDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T SystemUserforExtendedEnterpriseLearnerDataType
	var layout struct {
		*T
		AccountExpirationDate *xsdDateTime `xml:"urn:com.workday/bsvc Account_Expiration_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AccountExpirationDate = (*xsdDateTime)(layout.T.AccountExpirationDate)
	return e.EncodeElement(layout, start)
}
func (t *SystemUserforExtendedEnterpriseLearnerDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SystemUserforExtendedEnterpriseLearnerDataType
	var overlay struct {
		*T
		AccountExpirationDate *xsdDateTime `xml:"urn:com.workday/bsvc Account_Expiration_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AccountExpirationDate = (*xsdDateTime)(overlay.T.AccountExpirationDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type TaxApplicabilityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TaxApplicabilityObjectType struct {
	ID         []TaxApplicabilityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type UniqueIdentifierObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type UniqueIdentifierObjectType struct {
	ID         []UniqueIdentifierObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type UnitofMeasureObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type UnitofMeasureObjectType struct {
	ID         []UnitofMeasureObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type ValidationErrorType struct {
	Message       string `xml:"urn:com.workday/bsvc Message,omitempty"`
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
	Xpath         string `xml:"urn:com.workday/bsvc Xpath,omitempty"`
}

type ValidationFaultType struct {
	ValidationError []ValidationErrorType `xml:"urn:com.workday/bsvc Validation_Error,omitempty"`
}

// Wrapper element for Validation Message Component element
type ValidationMessageDataType struct {
	MessageComponentData []NotificationMessageComponentDataType `xml:"urn:com.workday/bsvc Message_Component_Data,omitempty"`
}

type WorkdayCommonHeaderType struct {
	IncludeReferenceDescriptorsInResponse *bool `xml:"urn:com.workday/bsvc Include_Reference_Descriptors_In_Response,omitempty"`
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
