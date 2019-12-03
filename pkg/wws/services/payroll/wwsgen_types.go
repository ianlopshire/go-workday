package payroll

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// Contains a unique identifier for an instance of an object.
type ARRCOAGIRCRubricValueObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ARRCOAGIRCRubricValueObjectType struct {
	ID         []ARRCOAGIRCRubricValueObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This is the wrapper element for Box 13 W-2 Additional Data.
type AdditionalDataType struct {
	StatutoryEmployee *bool `xml:"urn:com.workday/bsvc Statutory_Employee,omitempty"`
	RetirementPlan    *bool `xml:"urn:com.workday/bsvc Retirement_Plan,omitempty"`
	ThirdPartySickPay *bool `xml:"urn:com.workday/bsvc Third_Party_Sick_Pay,omitempty"`
}

// Additional Input Details
type AdditionalInputDetailsType struct {
	RelatedCalculationReference *RelatedCalculationAllObjectType `xml:"urn:com.workday/bsvc Related_Calculation_Reference,omitempty"`
	InputValue                  float64                          `xml:"urn:com.workday/bsvc Input_Value,omitempty"`
}

// Address information
type AddressInformationDataType struct {
	CountryReference             *CountryObjectType                            `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	LastModified                 *time.Time                                    `xml:"urn:com.workday/bsvc Last_Modified,omitempty"`
	AddressLineData              []AddressLineInformationDataType              `xml:"urn:com.workday/bsvc Address_Line_Data,omitempty"`
	Municipality                 string                                        `xml:"urn:com.workday/bsvc Municipality,omitempty"`
	CountryCityReference         *CountryCityObjectType                        `xml:"urn:com.workday/bsvc Country_City_Reference,omitempty"`
	SubmunicipalityData          []SubmunicipalityInformationDataType          `xml:"urn:com.workday/bsvc Submunicipality_Data,omitempty"`
	CountryRegionReference       *CountryRegionObjectType                      `xml:"urn:com.workday/bsvc Country_Region_Reference,omitempty"`
	CountryRegionDescriptor      string                                        `xml:"urn:com.workday/bsvc Country_Region_Descriptor,omitempty"`
	SubregionData                []SubregionInformationDataType                `xml:"urn:com.workday/bsvc Subregion_Data,omitempty"`
	PostalCode                   string                                        `xml:"urn:com.workday/bsvc Postal_Code,omitempty"`
	UsageData                    []CommunicationMethodUsageInformationDataType `xml:"urn:com.workday/bsvc Usage_Data,omitempty"`
	NumberofDays                 float64                                       `xml:"urn:com.workday/bsvc Number_of_Days,omitempty"`
	MunicipalityLocal            string                                        `xml:"urn:com.workday/bsvc Municipality_Local,omitempty"`
	AddressReference             *AddressReferenceObjectType                   `xml:"urn:com.workday/bsvc Address_Reference,omitempty"`
	AddressID                    string                                        `xml:"urn:com.workday/bsvc Address_ID,omitempty"`
	FormattedAddress             string                                        `xml:"urn:com.workday/bsvc Formatted_Address,attr,omitempty"`
	AddressFormatType            string                                        `xml:"urn:com.workday/bsvc Address_Format_Type,attr,omitempty"`
	DefaultedBusinessSiteAddress bool                                          `xml:"urn:com.workday/bsvc Defaulted_Business_Site_Address,attr,omitempty"`
	Delete                       bool                                          `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	DoNotReplaceAll              bool                                          `xml:"urn:com.workday/bsvc Do_Not_Replace_All,attr,omitempty"`
	EffectiveDate                time.Time                                     `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
}

func (t *AddressInformationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AddressInformationDataType
	var layout struct {
		*T
		LastModified  *xsdDateTime `xml:"urn:com.workday/bsvc Last_Modified,omitempty"`
		EffectiveDate *xsdDate     `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LastModified = (*xsdDateTime)(layout.T.LastModified)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *AddressInformationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AddressInformationDataType
	var overlay struct {
		*T
		LastModified  *xsdDateTime `xml:"urn:com.workday/bsvc Last_Modified,omitempty"`
		EffectiveDate *xsdDate     `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LastModified = (*xsdDateTime)(overlay.T.LastModified)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// The address line for the address. This typically contains Street name, street number, apartment, suite number.
type AddressLineInformationDataType struct {
	Value      string `xml:",chardata"`
	Descriptor string `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
	Type       string `xml:"urn:com.workday/bsvc Type,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AddressReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AddressReferenceObjectType struct {
	ID         []AddressReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AlternateNameUsageObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AlternateNameUsageObjectType struct {
	ID         []AlternateNameUsageObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ApplicationBatchObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ApplicationBatchObjectType struct {
	ID         []ApplicationBatchObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing application related exceptions data
type ApplicationInstanceExceptionsDataType struct {
	ExceptionData []ExceptionDataType `xml:"urn:com.workday/bsvc Exception_Data,omitempty"`
}

// Element containing Exceptions Data
type ApplicationInstanceRelatedExceptionsDataType struct {
	ExceptionsData []ApplicationInstanceExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Data,omitempty"`
}

// Specifies the options for the business process, as well as the data for the Costing Allocation
type AssignCostingAllocationRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	CostingAllocationData     *CostingAllocationDataType     `xml:"urn:com.workday/bsvc Costing_Allocation_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the Event and Costing data that was created by this request
type AssignCostingAllocationResponseType struct {
	AssignCostingAllocationEventReference *UniqueIdentifierObjectType     `xml:"urn:com.workday/bsvc Assign_Costing_Allocation_Event_Reference,omitempty"`
	WorkerReference                       *WorkerObjectType               `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	PositionRestrictionReference          *PositionRestrictionsObjectType `xml:"urn:com.workday/bsvc Position_Restriction_Reference,omitempty"`
	Version                               string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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
type BackgroundProcessObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type BackgroundProcessObjectType struct {
	ID         []BackgroundProcessObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Background Process Runtime Data
type BackgroundProcessRuntimeDataType struct {
	ID                           string                               `xml:"urn:com.workday/bsvc ID,omitempty"`
	BackgroundProcessMessageData []BackgroundProcessMessageDataWSType `xml:"urn:com.workday/bsvc Background_Process_Message_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BackgroundProcessRuntimeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BackgroundProcessRuntimeObjectType struct {
	ID         []BackgroundProcessRuntimeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type BackgroundProcessTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BackgroundProcessTypeObjectType struct {
	ID         []BackgroundProcessTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BalancePeriodAllObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BalancePeriodAllObjectType struct {
	ID         []BalancePeriodAllObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Balance Period Data for Payroll Balances
type BalancePeriodDataforGetPayrollBalancesType struct {
	BalancePeriodName       string                        `xml:"urn:com.workday/bsvc Balance_Period_Name,omitempty"`
	BalancePeriodStartDate  *time.Time                    `xml:"urn:com.workday/bsvc Balance_Period_Start_Date,omitempty"`
	BalancePeriodEndDate    *time.Time                    `xml:"urn:com.workday/bsvc Balance_Period_End_Date,omitempty"`
	BalancePeriodDetailData []BalancePeriodDetailDataType `xml:"urn:com.workday/bsvc Balance_Period_Detail_Data,omitempty"`
}

func (t *BalancePeriodDataforGetPayrollBalancesType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BalancePeriodDataforGetPayrollBalancesType
	var layout struct {
		*T
		BalancePeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Balance_Period_Start_Date,omitempty"`
		BalancePeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Balance_Period_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.BalancePeriodStartDate = (*xsdDate)(layout.T.BalancePeriodStartDate)
	layout.BalancePeriodEndDate = (*xsdDate)(layout.T.BalancePeriodEndDate)
	return e.EncodeElement(layout, start)
}
func (t *BalancePeriodDataforGetPayrollBalancesType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BalancePeriodDataforGetPayrollBalancesType
	var overlay struct {
		*T
		BalancePeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Balance_Period_Start_Date,omitempty"`
		BalancePeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Balance_Period_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.BalancePeriodStartDate = (*xsdDate)(overlay.T.BalancePeriodStartDate)
	overlay.BalancePeriodEndDate = (*xsdDate)(overlay.T.BalancePeriodEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Balance Period Detail Data
type BalancePeriodDetailDataType struct {
	WorktagReference     []PayrollWorktagDataType   `xml:"urn:com.workday/bsvc Worktag_Reference,omitempty"`
	BalanceAmount        float64                    `xml:"urn:com.workday/bsvc Balance_Amount,omitempty"`
	CurrencyReference    *CurrencyReferenceDataType `xml:"urn:com.workday/bsvc Currency_Reference"`
	PayrollIDBalanceData []PayrollIDBalanceDataType `xml:"urn:com.workday/bsvc Payroll_ID_Balance_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankAccountObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankAccountObjectType struct {
	ID         []BankAccountObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankAccountTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankAccountTypeObjectType struct {
	ID         []BankAccountTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Bankruptcy Order Specific Data
type BankruptcyOrderSpecificDataType struct {
	Chapter13 bool `xml:"urn:com.workday/bsvc Chapter_13"`
	Chapter7  bool `xml:"urn:com.workday/bsvc Chapter_7"`
}

// Box 16 A-A Data
type Box16AADataType struct {
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Code   string  `xml:"urn:com.workday/bsvc Code,omitempty"`
}

// Box 16 A-B Data
type Box16ABDataType struct {
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Code   string  `xml:"urn:com.workday/bsvc Code,omitempty"`
}

// Box 16 A-C Data
type Box16ACDataType struct {
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Code   string  `xml:"urn:com.workday/bsvc Code,omitempty"`
}

// Box 16 A-D Data
type Box16ADDataType struct {
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Code   string  `xml:"urn:com.workday/bsvc Code,omitempty"`
}

// Box 16 A-E Data
type Box16AEDataType struct {
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Code   string  `xml:"urn:com.workday/bsvc Code,omitempty"`
}

// Box 16 A-F Data
type Box16AFDataType struct {
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Code   string  `xml:"urn:com.workday/bsvc Code,omitempty"`
}

// Box 16 A-G Data
type Box16AGDataType struct {
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Code   string  `xml:"urn:com.workday/bsvc Code,omitempty"`
}

// Box 16 A-H Data
type Box16AHDataType struct {
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Code   string  `xml:"urn:com.workday/bsvc Code,omitempty"`
}

// Box 16 A-I Data
type Box16AIDataType struct {
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Code   string  `xml:"urn:com.workday/bsvc Code,omitempty"`
}

// Box 16 Data Wrapper
type Box16DataWrapperType struct {
	OvertimeDuringEmergency              []Box16AADataType `xml:"urn:com.workday/bsvc Overtime_During_Emergency,omitempty"`
	OvertimeByPoliceMember               []Box16ABDataType `xml:"urn:com.workday/bsvc Overtime_By_Police_Member,omitempty"`
	Stipends                             []Box16ACDataType `xml:"urn:com.workday/bsvc Stipends,omitempty"`
	CompensationforResearcherorScientist []Box16ADDataType `xml:"urn:com.workday/bsvc Compensation_for_Researcher_or_Scientist,omitempty"`
	SalaryNotOver40K                     []Box16AEDataType `xml:"urn:com.workday/bsvc Salary_Not_Over_40K,omitempty"`
	VacationandSickPayPublicEmployees    []Box16AFDataType `xml:"urn:com.workday/bsvc Vacation_and_Sick_Pay_Public_Employees,omitempty"`
	DisasterAssistance                   []Box16AGDataType `xml:"urn:com.workday/bsvc Disaster_Assistance,omitempty"`
	PublicEmployeesVoluntaryTransition   []Box16AHDataType `xml:"urn:com.workday/bsvc Public_Employees_Voluntary_Transition,omitempty"`
	CompensationforDismissal             []Box16AIDataType `xml:"urn:com.workday/bsvc Compensation_for_Dismissal,omitempty"`
}

// Business Entity Alternate Name Data
type BusinessEntityAlternateNameDataType struct {
	AlternateName               string                         `xml:"urn:com.workday/bsvc Alternate_Name"`
	AlternateNameUsageReference []AlternateNameUsageObjectType `xml:"urn:com.workday/bsvc Alternate_Name_Usage_Reference"`
}

// Business Entity Logo Image Data is an element containing the logo data.
type BusinessEntityLogoImageDataType struct {
	Filename string  `xml:"urn:com.workday/bsvc Filename"`
	Image    *[]byte `xml:"urn:com.workday/bsvc Image,omitempty"`
}

func (t *BusinessEntityLogoImageDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BusinessEntityLogoImageDataType
	var layout struct {
		*T
		Image *xsdBase64Binary `xml:"urn:com.workday/bsvc Image,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Image = (*xsdBase64Binary)(layout.T.Image)
	return e.EncodeElement(layout, start)
}
func (t *BusinessEntityLogoImageDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BusinessEntityLogoImageDataType
	var overlay struct {
		*T
		Image *xsdBase64Binary `xml:"urn:com.workday/bsvc Image,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Image = (*xsdBase64Binary)(overlay.T.Image)
	return d.DecodeElement(&overlay, &start)
}

// Business Entity Data Element which is a wrapper for all business entity data of name and contract information
type BusinessEntityWWSDataType struct {
	BusinessEntityName          string                           `xml:"urn:com.workday/bsvc Business_Entity_Name"`
	BusinessEntityPhoneticName  string                           `xml:"urn:com.workday/bsvc Business_Entity_Phonetic_Name,omitempty"`
	BusinessEntityTaxID         string                           `xml:"urn:com.workday/bsvc Business_Entity_Tax_ID,omitempty"`
	ExternalEntityID            string                           `xml:"urn:com.workday/bsvc External_Entity_ID,omitempty"`
	ContactData                 *ContactInformationDataType      `xml:"urn:com.workday/bsvc Contact_Data,omitempty"`
	BusinessEntityLogoImageData *BusinessEntityLogoImageDataType `xml:"urn:com.workday/bsvc Business_Entity_Logo_Image_Data,omitempty"`
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
type BusinessUnitObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BusinessUnitObjectType struct {
	ID         []BusinessUnitObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains sets of status counts of Payroll Results
type CalcStatusType struct {
	PayGroupPeriodPayCalculationStatusAsOfNow []PayGroupPeriodPayCalculationStatusAsOfNowType `xml:"urn:com.workday/bsvc Pay_Group_Period_Pay_Calculation_Status_As_Of_Now,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CalendarQuarterObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CalendarQuarterObjectType struct {
	ID         []CalendarQuarterObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CalendarYearObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CalendarYearObjectType struct {
	ID         []CalendarYearObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This is the Canadian Record of Employment Data element containing attributes to be updated on the Canadian Record of Employment.
type CanadianRecordofEmploymentDataDataType struct {
	SerialNumber             string                      `xml:"urn:com.workday/bsvc Serial_Number,omitempty"`
	DateIssued               *time.Time                  `xml:"urn:com.workday/bsvc Date_Issued,omitempty"`
	PayrollRefNumber         string                      `xml:"urn:com.workday/bsvc Payroll_Ref_Number,omitempty"`
	BusinessNumber           string                      `xml:"urn:com.workday/bsvc Business_Number"`
	SIN                      string                      `xml:"urn:com.workday/bsvc SIN"`
	FirstDayWorked           time.Time                   `xml:"urn:com.workday/bsvc First_Day_Worked"`
	LastDayForWhichPaid      time.Time                   `xml:"urn:com.workday/bsvc Last_Day_For_Which_Paid"`
	FinalPayPeriodEndingDate time.Time                   `xml:"urn:com.workday/bsvc Final_Pay_Period_Ending_Date"`
	TotalInsurableEarnings   float64                     `xml:"urn:com.workday/bsvc Total_Insurable_Earnings,omitempty"`
	ROEStatusReference       *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc ROE_Status_Reference"`
}

func (t *CanadianRecordofEmploymentDataDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CanadianRecordofEmploymentDataDataType
	var layout struct {
		*T
		DateIssued               *xsdDate `xml:"urn:com.workday/bsvc Date_Issued,omitempty"`
		FirstDayWorked           *xsdDate `xml:"urn:com.workday/bsvc First_Day_Worked"`
		LastDayForWhichPaid      *xsdDate `xml:"urn:com.workday/bsvc Last_Day_For_Which_Paid"`
		FinalPayPeriodEndingDate *xsdDate `xml:"urn:com.workday/bsvc Final_Pay_Period_Ending_Date"`
	}
	layout.T = (*T)(t)
	layout.DateIssued = (*xsdDate)(layout.T.DateIssued)
	layout.FirstDayWorked = (*xsdDate)(&layout.T.FirstDayWorked)
	layout.LastDayForWhichPaid = (*xsdDate)(&layout.T.LastDayForWhichPaid)
	layout.FinalPayPeriodEndingDate = (*xsdDate)(&layout.T.FinalPayPeriodEndingDate)
	return e.EncodeElement(layout, start)
}
func (t *CanadianRecordofEmploymentDataDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CanadianRecordofEmploymentDataDataType
	var overlay struct {
		*T
		DateIssued               *xsdDate `xml:"urn:com.workday/bsvc Date_Issued,omitempty"`
		FirstDayWorked           *xsdDate `xml:"urn:com.workday/bsvc First_Day_Worked"`
		LastDayForWhichPaid      *xsdDate `xml:"urn:com.workday/bsvc Last_Day_For_Which_Paid"`
		FinalPayPeriodEndingDate *xsdDate `xml:"urn:com.workday/bsvc Final_Pay_Period_Ending_Date"`
	}
	overlay.T = (*T)(t)
	overlay.DateIssued = (*xsdDate)(overlay.T.DateIssued)
	overlay.FirstDayWorked = (*xsdDate)(&overlay.T.FirstDayWorked)
	overlay.LastDayForWhichPaid = (*xsdDate)(&overlay.T.LastDayForWhichPaid)
	overlay.FinalPayPeriodEndingDate = (*xsdDate)(&overlay.T.FinalPayPeriodEndingDate)
	return d.DecodeElement(&overlay, &start)
}

// Get Canadian Record of Employment Data Request Criteria
type CanadianRecordofEmploymentDataRequestCriteriaType struct {
}

// Canadian Record of Employment Data Request References
type CanadianRecordofEmploymentDataRequestReferencesType struct {
	CanadianRecordofEmploymentDataReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Canadian_Record_of_Employment_Data_Reference,omitempty"`
}

// Get Canadian Record of Employment Data Response
type CanadianRecordofEmploymentDataResponseDataType struct {
	CanadianRecordofEmploymentData []CanadianRecordofEmploymentDataType `xml:"urn:com.workday/bsvc Canadian_Record_of_Employment_Data,omitempty"`
}

// Canadian Record of Employment Data Response Group
type CanadianRecordofEmploymentDataResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Canadian Record of Employment Data
type CanadianRecordofEmploymentDataType struct {
	PrintingLanguage                 string                                                                      `xml:"urn:com.workday/bsvc Printing_Language,omitempty"`
	IssueType                        string                                                                      `xml:"urn:com.workday/bsvc Issue_Type,omitempty"`
	ROEDataReference                 *UniqueIdentifierObjectType                                                 `xml:"urn:com.workday/bsvc ROE_Data_Reference,omitempty"`
	ROESerialNumber                  string                                                                      `xml:"urn:com.workday/bsvc ROE_Serial_Number,omitempty"`
	EmployeeID                       string                                                                      `xml:"urn:com.workday/bsvc Employee_ID,omitempty"`
	EmployerPayrollReferenceNumber   string                                                                      `xml:"urn:com.workday/bsvc Employer_Payroll_Reference_Number"`
	OrganizationNumber               string                                                                      `xml:"urn:com.workday/bsvc Organization_Number,omitempty"`
	FolderCD                         string                                                                      `xml:"urn:com.workday/bsvc Folder_CD,omitempty"`
	PayPeriodType                    string                                                                      `xml:"urn:com.workday/bsvc Pay_Period_Type"`
	SocialInsuranceNumber            string                                                                      `xml:"urn:com.workday/bsvc Social_Insurance_Number,omitempty"`
	ROEAddressLine1                  string                                                                      `xml:"urn:com.workday/bsvc ROE_Address_Line_1"`
	ROEAddressLine2                  string                                                                      `xml:"urn:com.workday/bsvc ROE_Address_Line_2,omitempty"`
	ROEAddressLine3                  string                                                                      `xml:"urn:com.workday/bsvc ROE_Address_Line_3,omitempty"`
	PostalCode                       string                                                                      `xml:"urn:com.workday/bsvc Postal_Code,omitempty"`
	LegalFirstName                   string                                                                      `xml:"urn:com.workday/bsvc Legal_First_Name"`
	LegalLastName                    string                                                                      `xml:"urn:com.workday/bsvc Legal_Last_Name"`
	LegalMiddleName                  string                                                                      `xml:"urn:com.workday/bsvc Legal_Middle_Name,omitempty"`
	FirstDayWorked                   time.Time                                                                   `xml:"urn:com.workday/bsvc First_Day_Worked"`
	LastDayForWhichPaid              time.Time                                                                   `xml:"urn:com.workday/bsvc Last_Day_For_Which_Paid"`
	FinalPayPeriodEndDate            time.Time                                                                   `xml:"urn:com.workday/bsvc Final_Pay_Period_End_Date"`
	EmployeeOccupation               string                                                                      `xml:"urn:com.workday/bsvc Employee_Occupation,omitempty"`
	ROEExpectedRecallCode            string                                                                      `xml:"urn:com.workday/bsvc ROE_Expected_Recall_Code,omitempty"`
	ExpectedDateofReturn             *time.Time                                                                  `xml:"urn:com.workday/bsvc Expected_Date_of_Return,omitempty"`
	TotalInsurableHours              float64                                                                     `xml:"urn:com.workday/bsvc Total_Insurable_Hours"`
	TotalInsurableEarnings           float64                                                                     `xml:"urn:com.workday/bsvc Total_Insurable_Earnings"`
	ROEInsurableEarningsByPeriodData []CanadianRecordofEmploymentInsurableEarningsByPeriodDataType               `xml:"urn:com.workday/bsvc ROE_Insurable_Earnings_By_Period_Data"`
	ROEReasonContactAreaCodeNumber   string                                                                      `xml:"urn:com.workday/bsvc ROE_Reason_Contact_Area_Code_Number"`
	ROEReasonCodeforROEReason        string                                                                      `xml:"urn:com.workday/bsvc ROE_Reason_Code_for_ROE_Reason"`
	ROEReasonPhoneExtensionNumber    string                                                                      `xml:"urn:com.workday/bsvc ROE_Reason_Phone_Extension_Number,omitempty"`
	ROEReasonContactFirstName        string                                                                      `xml:"urn:com.workday/bsvc ROE_Reason_Contact_First_Name"`
	ROEReasonContactLastName         string                                                                      `xml:"urn:com.workday/bsvc ROE_Reason_Contact_Last_Name"`
	ROEReasonContactPhoneNumber      string                                                                      `xml:"urn:com.workday/bsvc ROE_Reason_Contact_Phone_Number"`
	VacationPayAmount                float64                                                                     `xml:"urn:com.workday/bsvc Vacation_Pay_Amount,omitempty"`
	VacationPayData                  []ROEVacationPayDataType                                                    `xml:"urn:com.workday/bsvc Vacation_Pay_Data,omitempty"`
	StatutoryHolidayData             []CanadianRecordofEmploymentStatutoryHolidayDetailDataType                  `xml:"urn:com.workday/bsvc Statutory_Holiday_Data,omitempty"`
	OtherMoniesData                  []OtherMoniesDataType                                                       `xml:"urn:com.workday/bsvc Other_Monies_Data,omitempty"`
	Comments                         string                                                                      `xml:"urn:com.workday/bsvc Comments,omitempty"`
	SpecialPaymentData               []CanadianRecordofEmploymentPaidSickLeaveMaternityandWageLossDetailDataType `xml:"urn:com.workday/bsvc Special_Payment_Data,omitempty"`
	PreferredUserLanguage            string                                                                      `xml:"urn:com.workday/bsvc Preferred_User_Language,omitempty"`
}

func (t *CanadianRecordofEmploymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CanadianRecordofEmploymentDataType
	var layout struct {
		*T
		FirstDayWorked        *xsdDate `xml:"urn:com.workday/bsvc First_Day_Worked"`
		LastDayForWhichPaid   *xsdDate `xml:"urn:com.workday/bsvc Last_Day_For_Which_Paid"`
		FinalPayPeriodEndDate *xsdDate `xml:"urn:com.workday/bsvc Final_Pay_Period_End_Date"`
		ExpectedDateofReturn  *xsdDate `xml:"urn:com.workday/bsvc Expected_Date_of_Return,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FirstDayWorked = (*xsdDate)(&layout.T.FirstDayWorked)
	layout.LastDayForWhichPaid = (*xsdDate)(&layout.T.LastDayForWhichPaid)
	layout.FinalPayPeriodEndDate = (*xsdDate)(&layout.T.FinalPayPeriodEndDate)
	layout.ExpectedDateofReturn = (*xsdDate)(layout.T.ExpectedDateofReturn)
	return e.EncodeElement(layout, start)
}
func (t *CanadianRecordofEmploymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CanadianRecordofEmploymentDataType
	var overlay struct {
		*T
		FirstDayWorked        *xsdDate `xml:"urn:com.workday/bsvc First_Day_Worked"`
		LastDayForWhichPaid   *xsdDate `xml:"urn:com.workday/bsvc Last_Day_For_Which_Paid"`
		FinalPayPeriodEndDate *xsdDate `xml:"urn:com.workday/bsvc Final_Pay_Period_End_Date"`
		ExpectedDateofReturn  *xsdDate `xml:"urn:com.workday/bsvc Expected_Date_of_Return,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FirstDayWorked = (*xsdDate)(&overlay.T.FirstDayWorked)
	overlay.LastDayForWhichPaid = (*xsdDate)(&overlay.T.LastDayForWhichPaid)
	overlay.FinalPayPeriodEndDate = (*xsdDate)(&overlay.T.FinalPayPeriodEndDate)
	overlay.ExpectedDateofReturn = (*xsdDate)(overlay.T.ExpectedDateofReturn)
	return d.DecodeElement(&overlay, &start)
}

// Insurable Earnings by Period
type CanadianRecordofEmploymentInsurableEarningsByPeriodDataType struct {
	PayrollROEPeriodNumber float64 `xml:"urn:com.workday/bsvc Payroll_ROE_Period_Number"`
	EIEarningValue         float64 `xml:"urn:com.workday/bsvc EI_Earning_Value"`
}

// Special Payment information
type CanadianRecordofEmploymentPaidSickLeaveMaternityandWageLossDetailDataType struct {
	ROESpecialPaymentCode   string     `xml:"urn:com.workday/bsvc ROE_Special_Payment_Code,omitempty"`
	SpecialPaymentDate      *time.Time `xml:"urn:com.workday/bsvc Special_Payment_Date,omitempty"`
	SpecialPaymentStartDate *time.Time `xml:"urn:com.workday/bsvc Special_Payment_Start_Date,omitempty"`
	SpecialPaymentEndDate   *time.Time `xml:"urn:com.workday/bsvc Special_Payment_End_Date,omitempty"`
	SpecialPaymentAmount    float64    `xml:"urn:com.workday/bsvc Special_Payment_Amount,omitempty"`
	PerDay                  *bool      `xml:"urn:com.workday/bsvc Per_Day,omitempty"`
	PerWeek                 *bool      `xml:"urn:com.workday/bsvc Per_Week,omitempty"`
}

func (t *CanadianRecordofEmploymentPaidSickLeaveMaternityandWageLossDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CanadianRecordofEmploymentPaidSickLeaveMaternityandWageLossDetailDataType
	var layout struct {
		*T
		SpecialPaymentDate      *xsdDate `xml:"urn:com.workday/bsvc Special_Payment_Date,omitempty"`
		SpecialPaymentStartDate *xsdDate `xml:"urn:com.workday/bsvc Special_Payment_Start_Date,omitempty"`
		SpecialPaymentEndDate   *xsdDate `xml:"urn:com.workday/bsvc Special_Payment_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.SpecialPaymentDate = (*xsdDate)(layout.T.SpecialPaymentDate)
	layout.SpecialPaymentStartDate = (*xsdDate)(layout.T.SpecialPaymentStartDate)
	layout.SpecialPaymentEndDate = (*xsdDate)(layout.T.SpecialPaymentEndDate)
	return e.EncodeElement(layout, start)
}
func (t *CanadianRecordofEmploymentPaidSickLeaveMaternityandWageLossDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CanadianRecordofEmploymentPaidSickLeaveMaternityandWageLossDetailDataType
	var overlay struct {
		*T
		SpecialPaymentDate      *xsdDate `xml:"urn:com.workday/bsvc Special_Payment_Date,omitempty"`
		SpecialPaymentStartDate *xsdDate `xml:"urn:com.workday/bsvc Special_Payment_Start_Date,omitempty"`
		SpecialPaymentEndDate   *xsdDate `xml:"urn:com.workday/bsvc Special_Payment_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.SpecialPaymentDate = (*xsdDate)(overlay.T.SpecialPaymentDate)
	overlay.SpecialPaymentStartDate = (*xsdDate)(overlay.T.SpecialPaymentStartDate)
	overlay.SpecialPaymentEndDate = (*xsdDate)(overlay.T.SpecialPaymentEndDate)
	return d.DecodeElement(&overlay, &start)
}

// This is the parent tag for the list of Statutory Holiday information
type CanadianRecordofEmploymentStatutoryHolidayDetailDataType struct {
	ROEDetailDate   *time.Time `xml:"urn:com.workday/bsvc ROE_Detail_Date,omitempty"`
	ROEDetailAmount float64    `xml:"urn:com.workday/bsvc ROE_Detail_Amount,omitempty"`
}

func (t *CanadianRecordofEmploymentStatutoryHolidayDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CanadianRecordofEmploymentStatutoryHolidayDetailDataType
	var layout struct {
		*T
		ROEDetailDate *xsdDate `xml:"urn:com.workday/bsvc ROE_Detail_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ROEDetailDate = (*xsdDate)(layout.T.ROEDetailDate)
	return e.EncodeElement(layout, start)
}
func (t *CanadianRecordofEmploymentStatutoryHolidayDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CanadianRecordofEmploymentStatutoryHolidayDetailDataType
	var overlay struct {
		*T
		ROEDetailDate *xsdDate `xml:"urn:com.workday/bsvc ROE_Detail_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ROEDetailDate = (*xsdDate)(overlay.T.ROEDetailDate)
	return d.DecodeElement(&overlay, &start)
}

// Change No Retro Processing Prior to Data
type ChangeNoRetroProcessingPriorToDataType struct {
	EmployeeReference                *EmployeeObjectType `xml:"urn:com.workday/bsvc Employee_Reference"`
	ProposedNoRetroProcessingPriorTo *time.Time          `xml:"urn:com.workday/bsvc Proposed_No_Retro_Processing_Prior_To,omitempty"`
}

func (t *ChangeNoRetroProcessingPriorToDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ChangeNoRetroProcessingPriorToDataType
	var layout struct {
		*T
		ProposedNoRetroProcessingPriorTo *xsdDate `xml:"urn:com.workday/bsvc Proposed_No_Retro_Processing_Prior_To,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ProposedNoRetroProcessingPriorTo = (*xsdDate)(layout.T.ProposedNoRetroProcessingPriorTo)
	return e.EncodeElement(layout, start)
}
func (t *ChangeNoRetroProcessingPriorToDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ChangeNoRetroProcessingPriorToDataType
	var overlay struct {
		*T
		ProposedNoRetroProcessingPriorTo *xsdDate `xml:"urn:com.workday/bsvc Proposed_No_Retro_Processing_Prior_To,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ProposedNoRetroProcessingPriorTo = (*xsdDate)(overlay.T.ProposedNoRetroProcessingPriorTo)
	return d.DecodeElement(&overlay, &start)
}

// Change No Retro Processing Prior To Request.
type ChangeNoRetroProcessingPriorToRequestType struct {
	BusinessProcessParameters          *BusinessProcessParametersType          `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	ChangeNoRetroProcessingPriorToData *ChangeNoRetroProcessingPriorToDataType `xml:"urn:com.workday/bsvc Change_No_Retro_Processing_Prior_To_Data"`
	Version                            string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put No Retro Processing Prior to date for worker Response.
type ChangeNoRetroProcessingPriorToResponseType struct {
	PayrollRetroWorkerDataEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Payroll_Retro_Worker_Data_Event_Reference,omitempty"`
	Version                              string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Encapsulating element for all Communication Method Usage data.
type CommunicationMethodUsageDataType struct {
	TypeReference   []CommunicationUsageTypeReferenceType `xml:"urn:com.workday/bsvc Type_Reference"`
	UseForReference []string                              `xml:"urn:com.workday/bsvc Use_For_Reference,omitempty"`
	Comments        string                                `xml:"urn:com.workday/bsvc Comments,omitempty"`
	Public          bool                                  `xml:"urn:com.workday/bsvc Public,attr,omitempty"`
}

// Encapsulating element for all Communication Method Usage data.
type CommunicationMethodUsageInformationDataType struct {
	TypeData                []CommunicationUsageTypeDataType               `xml:"urn:com.workday/bsvc Type_Data"`
	UseForReference         []CommunicationUsageBehaviorObjectType         `xml:"urn:com.workday/bsvc Use_For_Reference,omitempty"`
	UseForTenantedReference []CommunicationUsageBehaviorTenantedObjectType `xml:"urn:com.workday/bsvc Use_For_Tenanted_Reference,omitempty"`
	Comments                string                                         `xml:"urn:com.workday/bsvc Comments,omitempty"`
	Public                  bool                                           `xml:"urn:com.workday/bsvc Public,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CommunicationUsageBehaviorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CommunicationUsageBehaviorObjectType struct {
	ID         []CommunicationUsageBehaviorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CommunicationUsageBehaviorTenantedObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CommunicationUsageBehaviorTenantedObjectType struct {
	ID         []CommunicationUsageBehaviorTenantedObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Reference ID for the communication usage type.
type CommunicationUsageTypeDataType struct {
	TypeReference *CommunicationUsageTypeObjectType `xml:"urn:com.workday/bsvc Type_Reference"`
	Primary       bool                              `xml:"urn:com.workday/bsvc Primary,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CommunicationUsageTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CommunicationUsageTypeObjectType struct {
	ID         []CommunicationUsageTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Reference element representing a unique instance of Communication Usage Type.
type CommunicationUsageTypeReferenceType struct {
	Value   string `xml:",chardata"`
	Primary bool   `xml:"urn:com.workday/bsvc Primary,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompanyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompanyObjectType struct {
	ID         []CompanyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payment Dates for Company and Date Range Data element
type CompanyPaymentDatesDataType struct {
	CompanyReference *CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	PaymentDate      *time.Time         `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
}

func (t *CompanyPaymentDatesDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CompanyPaymentDatesDataType
	var layout struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PaymentDate = (*xsdDate)(layout.T.PaymentDate)
	return e.EncodeElement(layout, start)
}
func (t *CompanyPaymentDatesDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompanyPaymentDatesDataType
	var overlay struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDate = (*xsdDate)(overlay.T.PaymentDate)
	return d.DecodeElement(&overlay, &start)
}

// Payment Dates for Comapny and Date Range Request Criteria element
type CompanyPaymentDatesRequestCriteriaType struct {
	CompanyReference                      *CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference"`
	IncludeRelatedCompaniesforLegalEntity *bool              `xml:"urn:com.workday/bsvc Include_Related_Companies_for_Legal_Entity,omitempty"`
	ResultCompletedFrom                   time.Time          `xml:"urn:com.workday/bsvc Result_Completed_From"`
	ResultCompletedTo                     time.Time          `xml:"urn:com.workday/bsvc Result_Completed_To"`
}

func (t *CompanyPaymentDatesRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CompanyPaymentDatesRequestCriteriaType
	var layout struct {
		*T
		ResultCompletedFrom *xsdDateTime `xml:"urn:com.workday/bsvc Result_Completed_From"`
		ResultCompletedTo   *xsdDateTime `xml:"urn:com.workday/bsvc Result_Completed_To"`
	}
	layout.T = (*T)(t)
	layout.ResultCompletedFrom = (*xsdDateTime)(&layout.T.ResultCompletedFrom)
	layout.ResultCompletedTo = (*xsdDateTime)(&layout.T.ResultCompletedTo)
	return e.EncodeElement(layout, start)
}
func (t *CompanyPaymentDatesRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompanyPaymentDatesRequestCriteriaType
	var overlay struct {
		*T
		ResultCompletedFrom *xsdDateTime `xml:"urn:com.workday/bsvc Result_Completed_From"`
		ResultCompletedTo   *xsdDateTime `xml:"urn:com.workday/bsvc Result_Completed_To"`
	}
	overlay.T = (*T)(t)
	overlay.ResultCompletedFrom = (*xsdDateTime)(&overlay.T.ResultCompletedFrom)
	overlay.ResultCompletedTo = (*xsdDateTime)(&overlay.T.ResultCompletedTo)
	return d.DecodeElement(&overlay, &start)
}

// Payment Dates for Company and Date Range Response Data element
type CompanyPaymentDatesResponseDataType struct {
	CompanyPaymentDatesData []CompanyPaymentDatesDataType `xml:"urn:com.workday/bsvc Company_Payment_Dates_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompanyRelationshipTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompanyRelationshipTypeObjectType struct {
	ID         []CompanyRelationshipTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// All of the person's contact data (address, phone, email, instant messenger, web address).
type ContactInformationDataType struct {
	AddressData          []AddressInformationDataType          `xml:"urn:com.workday/bsvc Address_Data,omitempty"`
	PhoneData            []PhoneInformationDataType            `xml:"urn:com.workday/bsvc Phone_Data,omitempty"`
	EmailAddressData     []EmailAddressInformationDataType     `xml:"urn:com.workday/bsvc Email_Address_Data,omitempty"`
	InstantMessengerData []InstantMessengerInformationDataType `xml:"urn:com.workday/bsvc Instant_Messenger_Data,omitempty"`
	WebAddressData       []WebAddressInformationDataType       `xml:"urn:com.workday/bsvc Web_Address_Data,omitempty"`
}

// Reference element representing a unique instance of Contingent Worker.
type ContingentWorkerReferenceDataType struct {
	IntegrationIDReference *ExternalIntegrationIDReferenceDataType `xml:"urn:com.workday/bsvc Integration_ID_Reference"`
}

// Contains a unique identifier for an instance of an object.
type CostCenterObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CostCenterObjectType struct {
	ID         []CostCenterObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Specifies what the allocation will be applied to
type CostingAllocationDataType struct {
	WorkerReference                           *WorkerObjectType                   `xml:"urn:com.workday/bsvc Worker_Reference"`
	PositionRestrictionReference              *PositionElementObjectType          `xml:"urn:com.workday/bsvc Position_Restriction_Reference"`
	PositionReference                         *PositionElementObjectType          `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	EarningReference                          *EarningAllObjectType               `xml:"urn:com.workday/bsvc Earning_Reference,omitempty"`
	PositionElementEffectiveAsofDate          *time.Time                          `xml:"urn:com.workday/bsvc Position_Element_Effective_As-of_Date,omitempty"`
	ReplaceExistingCostingAllocationIntervals *bool                               `xml:"urn:com.workday/bsvc Replace_Existing_Costing_Allocation_Intervals,omitempty"`
	CostingAllocationIntervalData             []CostingAllocationIntervalDataType `xml:"urn:com.workday/bsvc Costing_Allocation_Interval_Data,omitempty"`
}

func (t *CostingAllocationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CostingAllocationDataType
	var layout struct {
		*T
		PositionElementEffectiveAsofDate *xsdDate `xml:"urn:com.workday/bsvc Position_Element_Effective_As-of_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PositionElementEffectiveAsofDate = (*xsdDate)(layout.T.PositionElementEffectiveAsofDate)
	return e.EncodeElement(layout, start)
}
func (t *CostingAllocationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CostingAllocationDataType
	var overlay struct {
		*T
		PositionElementEffectiveAsofDate *xsdDate `xml:"urn:com.workday/bsvc Position_Element_Effective_As-of_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PositionElementEffectiveAsofDate = (*xsdDate)(overlay.T.PositionElementEffectiveAsofDate)
	return d.DecodeElement(&overlay, &start)
}

// Details of the allocation, (e.g., a set of allocation dimensions and percentages)
type CostingAllocationDetailReplacementDataType struct {
	Order                                       string                                            `xml:"urn:com.workday/bsvc Order"`
	DefaultfromOrganizationAssignment           bool                                              `xml:"urn:com.workday/bsvc Default_from_Organization_Assignment"`
	CostingOverrideWorktagReference             []TenantedPayrollWorktagObjectType                `xml:"urn:com.workday/bsvc Costing_Override_Worktag_Reference"`
	DistributionPercent                         float64                                           `xml:"urn:com.workday/bsvc Distribution_Percent"`
	SalaryOvertheCapCostingAllocationDetailData []SalaryOvertheCapCostingAllocationDetailDataType `xml:"urn:com.workday/bsvc Salary_Over_the_Cap_Costing_Allocation_Detail_Data,omitempty"`
}

// Specifies an allocation for a given date range
type CostingAllocationIntervalDataType struct {
	CostingIntervalUpdateKey    []CostingIntervalUpdateKeyType               `xml:"urn:com.workday/bsvc Costing_Interval_Update_Key,omitempty"`
	CostingOverrideID           string                                       `xml:"urn:com.workday/bsvc Costing_Override_ID,omitempty"`
	StartDate                   *time.Time                                   `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                     *time.Time                                   `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	CostingAllocationDetailData []CostingAllocationDetailReplacementDataType `xml:"urn:com.workday/bsvc Costing_Allocation_Detail_Data,omitempty"`
}

func (t *CostingAllocationIntervalDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CostingAllocationIntervalDataType
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
func (t *CostingAllocationIntervalDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CostingAllocationIntervalDataType
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

// Submit only when updating or deleting an existing Costing Allocation Interval.  One key (Costing Override ID or Start Date) is required.
type CostingIntervalUpdateKeyType struct {
	CostingOverrideIDUpdateKey string    `xml:"urn:com.workday/bsvc Costing_Override_ID_Update_Key"`
	StartDateUpdateKey         time.Time `xml:"urn:com.workday/bsvc Start_Date_Update_Key"`
	Delete                     bool      `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

func (t *CostingIntervalUpdateKeyType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CostingIntervalUpdateKeyType
	var layout struct {
		*T
		StartDateUpdateKey *xsdDate `xml:"urn:com.workday/bsvc Start_Date_Update_Key"`
	}
	layout.T = (*T)(t)
	layout.StartDateUpdateKey = (*xsdDate)(&layout.T.StartDateUpdateKey)
	return e.EncodeElement(layout, start)
}
func (t *CostingIntervalUpdateKeyType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CostingIntervalUpdateKeyType
	var overlay struct {
		*T
		StartDateUpdateKey *xsdDate `xml:"urn:com.workday/bsvc Start_Date_Update_Key"`
	}
	overlay.T = (*T)(t)
	overlay.StartDateUpdateKey = (*xsdDate)(&overlay.T.StartDateUpdateKey)
	return d.DecodeElement(&overlay, &start)
}

// High-level criteria identifying individual Costing Allocation Overrides to be returned by the Get Worker Costing Allocations service.
type CostingOverrideCriteriaType struct {
	WorkerReference              []WorkerObjectType               `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	PositionRestrictionReference []PositionRestrictionsObjectType `xml:"urn:com.workday/bsvc Position_Restriction_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CountryCityObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type CountryCityObjectType struct {
	ID         []CountryCityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Country ISO code. If the Country ISO code is specified, then this ISO code will be used to determine the Country Phone Code for (eliminate space between f and o) the phone. Pass this ISO code to distinguish between multiple countries sharing the same Country Phone Code. (For example, 1 is the Country Phone Code that is shared by USA, Canada, Dominican Republic, Bermuda, Jamaica, and Puerto Rico.)
type CountryReferenceType struct {
	CountryISOCode string `xml:"urn:com.workday/bsvc Country_ISO_Code"`
}

// Contains a unique identifier for an instance of an object.
type CountryRegionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CountryRegionObjectType struct {
	ID         []CountryRegionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Creditor Garnishment Specific Data
type CreditorGarnishmentSpecificDataType struct {
	CreditorGarnishmentTypeReference    *UniqueIdentifierObjectType    `xml:"urn:com.workday/bsvc Creditor_Garnishment_Type_Reference"`
	NonResidentStateReference           *PayrollTaxAuthorityObjectType `xml:"urn:com.workday/bsvc Non_Resident_State_Reference,omitempty"`
	HeadofHousehold                     *bool                          `xml:"urn:com.workday/bsvc Head_of_Household,omitempty"`
	NumberofDependents                  float64                        `xml:"urn:com.workday/bsvc Number_of_Dependents,omitempty"`
	PayrollLocalCountyAuthorityFIPSCode string                         `xml:"urn:com.workday/bsvc Payroll_Local_County_Authority_FIPS_Code,omitempty"`
	WorkerisLaborerorMechanic           *bool                          `xml:"urn:com.workday/bsvc Worker_is_Laborer_or_Mechanic,omitempty"`
	WorkerIncomeisPovertyLevel          *bool                          `xml:"urn:com.workday/bsvc Worker_Income_is_Poverty_Level,omitempty"`
	GoodCauseLimitPercent               float64                        `xml:"urn:com.workday/bsvc Good_Cause_Limit_Percent,omitempty"`
	WeeklyGrossWages                    float64                        `xml:"urn:com.workday/bsvc Weekly_Gross_Wages,omitempty"`
	ExpectedAnnualEarnings              float64                        `xml:"urn:com.workday/bsvc Expected_Annual_Earnings,omitempty"`
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

// This element references a unique type of Currency.
type CurrencyReferenceDataType struct {
	CurrencyCode string `xml:"urn:com.workday/bsvc Currency_Code"`
}

// Contains a unique identifier for an instance of an object.
type CustomOrganizationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomOrganizationObjectType struct {
	ID         []CustomOrganizationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Custom Organizations Worktag Data
type CustomOrganizationWorktagDataType struct {
	OrganizationTypeReference          *OrganizationTypeObjectType   `xml:"urn:com.workday/bsvc Organization_Type_Reference,omitempty"`
	CustomOrganizationWorktagReference *CustomOrganizationObjectType `xml:"urn:com.workday/bsvc Custom_Organization_Worktag_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag01ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag01ObjectType struct {
	ID         []CustomWorktag01ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag02ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag02ObjectType struct {
	ID         []CustomWorktag02ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag03ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag03ObjectType struct {
	ID         []CustomWorktag03ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag04ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag04ObjectType struct {
	ID         []CustomWorktag04ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag05ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag05ObjectType struct {
	ID         []CustomWorktag05ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag06ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag06ObjectType struct {
	ID         []CustomWorktag06ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag07ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag07ObjectType struct {
	ID         []CustomWorktag07ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag08ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag08ObjectType struct {
	ID         []CustomWorktag08ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag09ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag09ObjectType struct {
	ID         []CustomWorktag09ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag10ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag10ObjectType struct {
	ID         []CustomWorktag10ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag11ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag11ObjectType struct {
	ID         []CustomWorktag11ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag12ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag12ObjectType struct {
	ID         []CustomWorktag12ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag13ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag13ObjectType struct {
	ID         []CustomWorktag13ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag14ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag14ObjectType struct {
	ID         []CustomWorktag14ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomWorktag15ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomWorktag15ObjectType struct {
	ID         []CustomWorktag15ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DeductionAllObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DeductionAllObjectType struct {
	ID         []DeductionAllObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DeductionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DeductionObjectType struct {
	ID         []DeductionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DeductionRecipientObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DeductionRecipientObjectType struct {
	ID         []DeductionRecipientObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This is a name uniquely identifying a deduction already established in the Workday Payroll system.
type DeductionReferenceType struct {
	Code string `xml:"urn:com.workday/bsvc Code"`
}

// Contains a unique identifier for an instance of an object.
type DeductionWorkdayOwnedObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DeductionWorkdayOwnedObjectType struct {
	ID         []DeductionWorkdayOwnedObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This is wrapper element for W-2 Box 12 Data.
type DeferredandOtherCompensationDataType struct {
	Code   string  `xml:"urn:com.workday/bsvc Code,omitempty"`
	Year   string  `xml:"urn:com.workday/bsvc Year,omitempty"`
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
}

// Result of the evaluation of an External Field based on a contextual instance.
type DocumentFieldResultDataType struct {
	FieldReference *IntegrationDocumentFieldObjectType `xml:"urn:com.workday/bsvc Field_Reference,omitempty"`
	Value          string                              `xml:"urn:com.workday/bsvc Value,omitempty"`
}

// This is the wrapper element for Box 13 W-2 Additional Data.
type EFW2AdditionalDataType struct {
	StatutoryEmployee *bool `xml:"urn:com.workday/bsvc Statutory_Employee,omitempty"`
	RetirementPlan    *bool `xml:"urn:com.workday/bsvc Retirement_Plan,omitempty"`
	ThirdPartySickPay *bool `xml:"urn:com.workday/bsvc Third_Party_Sick_Pay,omitempty"`
}

// This element contains the worker W-2 amounts for box 12.
type EFW2DeferredandOtherCompensationDataType struct {
	Code   string  `xml:"urn:com.workday/bsvc Code,omitempty"`
	Year   string  `xml:"urn:com.workday/bsvc Year,omitempty"`
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
}

// This element contains the company W-2 total amounts for box 12.
type EFW2EmployerDeferredandOtherCompensationDataType struct {
	Code   string  `xml:"urn:com.workday/bsvc Code,omitempty"`
	Year   string  `xml:"urn:com.workday/bsvc Year,omitempty"`
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
}

// This element contains the worker W-2 amounts for box 14.
type EFW2OtherDataType struct {
	Label      string  `xml:"urn:com.workday/bsvc Label,omitempty"`
	Amount     float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	PlanNumber string  `xml:"urn:com.workday/bsvc Plan_Number,omitempty"`
}

// This element contains the W-2 amounts for boxes 1 through  11.
//
// Box 1 search string:  Wages, tips, and other compensation
// Box 2 search string:  Federal income tax withheld
// Box 3 search string:  Social security wages
// Box 4 search string:  Social security tax withheld
// Box 5 search string:  Medicare wages and tips
// Box 6 search string:  Medicare tax withheld
// Box 7 search string:  Social security tips
// Box 8 search string:  Allocated tips
// Box 10 search string:  Dependent care benefits
// Box 11 NQPS search string:  Non-Qualified Pension-Section 457 Dist YTD
// Box 11 NQPNS search string:  Non-Qualified Pension-Non-section 457 Dist YTD
// Box 11 NQPSC search string:  Non-Qualified Pension-Section 457 Contrib YTD
// Box 11 NQPNSC search string:  Non-Qualified Pension-Non-section 457 Contrib YTD
type EFW2YearEndEmployerBoxDataType struct {
	BoxDescription string  `xml:"urn:com.workday/bsvc Box_Description,omitempty"`
	Amount         float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
}

// Request Criteria
type EFW2YearEndEmployerFilingDataRequestCriteriaType struct {
	CompanyReference      *CompanyObjectType      `xml:"urn:com.workday/bsvc Company_Reference"`
	CalendarYearReference *CalendarYearObjectType `xml:"urn:com.workday/bsvc Calendar_Year_Reference"`
	IncludeAllWorkers     *bool                   `xml:"urn:com.workday/bsvc Include_All_Workers,omitempty"`
}

// This element contains the employer details for the W-2 data.
type EFW2YearEndEmployerFilingDataType struct {
	CalendarYearReference                *CalendarYearObjectType                            `xml:"urn:com.workday/bsvc Calendar_Year_Reference,omitempty"`
	EIN                                  string                                             `xml:"urn:com.workday/bsvc EIN,omitempty"`
	CompanyName                          string                                             `xml:"urn:com.workday/bsvc Company_Name,omitempty"`
	KindofEmployerReference              *KindofEmployerforPayrollTaxFilingObjectType       `xml:"urn:com.workday/bsvc Kind_of_Employer_Reference,omitempty"`
	ThirdPartySickPayIndicator           *bool                                              `xml:"urn:com.workday/bsvc Third_Party_Sick_Pay_Indicator,omitempty"`
	EFW2YearEndSubmitterData             []EFW2YearEndSubmitterDataType                     `xml:"urn:com.workday/bsvc EFW2_Year_End_Submitter_Data,omitempty"`
	EFW2YearEndEmployerBoxData           []EFW2YearEndEmployerBoxDataType                   `xml:"urn:com.workday/bsvc EFW2_Year_End_Employer_Box_Data,omitempty"`
	EFW2DeferredandOtherCompensationData []EFW2EmployerDeferredandOtherCompensationDataType `xml:"urn:com.workday/bsvc EFW2_Deferred_and_Other_Compensation_Data,omitempty"`
}

// Response Data
type EFW2YearEndEmployerFilingResponseDataType struct {
	EFW2YearEndEmployerFiling []EFW2YearEndEmployerFilingType `xml:"urn:com.workday/bsvc EFW2_Year_End_Employer_Filing,omitempty"`
}

// Response data for Employer W-2 totals.
type EFW2YearEndEmployerFilingType struct {
	CompanyReference              *CompanyObjectType                 `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	EFW2YearEndEmployerFilingData *EFW2YearEndEmployerFilingDataType `xml:"urn:com.workday/bsvc EFW2_Year_End_Employer_Filing_Data,omitempty"`
}

// This element contains information about the submitter of the data and contact information if there are processing problems.
type EFW2YearEndSubmitterDataType struct {
	EIN                 string `xml:"urn:com.workday/bsvc EIN,omitempty"`
	CompanyName         string `xml:"urn:com.workday/bsvc Company_Name,omitempty"`
	AddressLine1        string `xml:"urn:com.workday/bsvc Address_Line_1,omitempty"`
	AddressLine2        string `xml:"urn:com.workday/bsvc Address_Line_2,omitempty"`
	City                string `xml:"urn:com.workday/bsvc City,omitempty"`
	State               string `xml:"urn:com.workday/bsvc State,omitempty"`
	PostalCode          string `xml:"urn:com.workday/bsvc Postal_Code,omitempty"`
	ContactName         string `xml:"urn:com.workday/bsvc Contact_Name,omitempty"`
	ContactPhoneNumber  string `xml:"urn:com.workday/bsvc Contact_Phone_Number,omitempty"`
	ContactEmailAddress string `xml:"urn:com.workday/bsvc Contact_Email_Address,omitempty"`
}

// This element contains the W-2 amounts for boxes 1 through  11.
//
// Box 1 search string:  Wages, tips, and other compensation
// Box 2 search string:  Federal income tax withheld
// Box 3 search string:  Social security wages
// Box 4 search string:  Social security tax withheld
// Box 5 search string:  Medicare wages and tips
// Box 6 search string:  Medicare tax withheld
// Box 7 search string:  Social security tips
// Box 8 search string:  Allocated tips
// Box 10 search string:  Dependent care benefits
// Box 11 NQPS search string:  Non-Qualified Pension-Section 457 Dist YTD
// Box 11 NQPNS search string:  Non-Qualified Pension-Non-section 457 Dist YTD
// Box 11 NQPSC search string:  Non-Qualified Pension-Section 457 Contrib YTD
// Box 11 NQPNSC search string:  Non-Qualified Pension-Non-section 457 Contrib YTD
type EFW2YearEndWorkerBoxDataType struct {
	BoxDescription string  `xml:"urn:com.workday/bsvc Box_Description,omitempty"`
	Amount         float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
}

// Request Criteria
type EFW2YearEndWorkerFilingDataRequestCriteriaType struct {
	CompanyReference      *CompanyObjectType      `xml:"urn:com.workday/bsvc Company_Reference"`
	CalendarYearReference *CalendarYearObjectType `xml:"urn:com.workday/bsvc Calendar_Year_Reference"`
	IncludeAllWorkers     *bool                   `xml:"urn:com.workday/bsvc Include_All_Workers,omitempty"`
}

// This element contains the details of the W-2 for a worker.
type EFW2YearEndWorkerFilingDataType struct {
	WorkerReference                      *WorkerObjectType                          `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	SSN                                  string                                     `xml:"urn:com.workday/bsvc SSN,omitempty"`
	EmployeeFirstName                    string                                     `xml:"urn:com.workday/bsvc Employee_First_Name,omitempty"`
	EmployeeMiddleInitial                string                                     `xml:"urn:com.workday/bsvc Employee_Middle_Initial,omitempty"`
	EmployeeLastName                     string                                     `xml:"urn:com.workday/bsvc Employee_Last_Name,omitempty"`
	Suffix                               string                                     `xml:"urn:com.workday/bsvc Suffix,omitempty"`
	AddressLine1                         string                                     `xml:"urn:com.workday/bsvc Address_Line_1,omitempty"`
	AddressLine2                         string                                     `xml:"urn:com.workday/bsvc Address_Line_2,omitempty"`
	City                                 string                                     `xml:"urn:com.workday/bsvc City,omitempty"`
	State                                string                                     `xml:"urn:com.workday/bsvc State,omitempty"`
	PostalCode                           string                                     `xml:"urn:com.workday/bsvc Postal_Code,omitempty"`
	CountryName                          string                                     `xml:"urn:com.workday/bsvc Country_Name,omitempty"`
	EFW2YearEndWorkerBoxData             []EFW2YearEndWorkerBoxDataType             `xml:"urn:com.workday/bsvc EFW2_Year_End_Worker_Box_Data,omitempty"`
	EFW2DeferredandOtherCompensationData []EFW2DeferredandOtherCompensationDataType `xml:"urn:com.workday/bsvc EFW2_Deferred_and_Other_Compensation_Data,omitempty"`
	EFW2AdditionalData                   []EFW2AdditionalDataType                   `xml:"urn:com.workday/bsvc EFW2_Additional_Data,omitempty"`
	EFW2OtherData                        []EFW2OtherDataType                        `xml:"urn:com.workday/bsvc EFW2_Other_Data,omitempty"`
	EFW2YearEndWorkerStateData           []EFW2YearEndWorkerStateDataType           `xml:"urn:com.workday/bsvc EFW2_Year_End_Worker_State_Data,omitempty"`
}

// Response Data
type EFW2YearEndWorkerFilingResponseDataType struct {
	EFW2YearEndWorkerFiling []EFW2YearEndWorkerFilingType `xml:"urn:com.workday/bsvc EFW2_Year_End_Worker_Filing,omitempty"`
}

// Response data for W-2
type EFW2YearEndWorkerFilingType struct {
	W2InstanceID                string                           `xml:"urn:com.workday/bsvc W-2_Instance_ID,omitempty"`
	CompletedMoment             *time.Time                       `xml:"urn:com.workday/bsvc Completed_Moment,omitempty"`
	EFW2YearEndWorkerFilingData *EFW2YearEndWorkerFilingDataType `xml:"urn:com.workday/bsvc EFW2_Year_End_Worker_Filing_Data,omitempty"`
}

func (t *EFW2YearEndWorkerFilingType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EFW2YearEndWorkerFilingType
	var layout struct {
		*T
		CompletedMoment *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Moment,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CompletedMoment = (*xsdDateTime)(layout.T.CompletedMoment)
	return e.EncodeElement(layout, start)
}
func (t *EFW2YearEndWorkerFilingType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EFW2YearEndWorkerFilingType
	var overlay struct {
		*T
		CompletedMoment *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Moment,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CompletedMoment = (*xsdDateTime)(overlay.T.CompletedMoment)
	return d.DecodeElement(&overlay, &start)
}

// This element contains the local data for the W-2.
type EFW2YearEndWorkerLocalDataType struct {
	Locality          string  `xml:"urn:com.workday/bsvc Locality,omitempty"`
	TaxTypeCode       string  `xml:"urn:com.workday/bsvc Tax_Type_Code,omitempty"`
	LocalTaxableWages float64 `xml:"urn:com.workday/bsvc Local_Taxable_Wages,omitempty"`
	LocalTaxWithheld  float64 `xml:"urn:com.workday/bsvc Local_Tax_Withheld,omitempty"`
	OtherInformation  string  `xml:"urn:com.workday/bsvc Other_Information,omitempty"`
}

// This element contains the state and local data for the W-2.
type EFW2YearEndWorkerStateDataType struct {
	StateEIN                   string                           `xml:"urn:com.workday/bsvc State_EIN,omitempty"`
	State                      string                           `xml:"urn:com.workday/bsvc State,omitempty"`
	StateFIPSCode              string                           `xml:"urn:com.workday/bsvc State_FIPS_Code,omitempty"`
	StateTaxableWages          float64                          `xml:"urn:com.workday/bsvc State_Taxable_Wages,omitempty"`
	StateTaxWithheld           float64                          `xml:"urn:com.workday/bsvc State_Tax_Withheld,omitempty"`
	EFW2YearEndWorkerLocalData []EFW2YearEndWorkerLocalDataType `xml:"urn:com.workday/bsvc EFW2_Year_End_Worker_Local_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EarningAllObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EarningAllObjectType struct {
	ID         []EarningAllObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This is a name uniquely identifying an earning already established in the Workday Payroll system.
type EarningReferenceType struct {
	Code string `xml:"urn:com.workday/bsvc Code"`
}

// Email Address Data
type EmailAddressDataType struct {
	EmailAddress string                             `xml:"urn:com.workday/bsvc Email_Address,omitempty"`
	EmailComment string                             `xml:"urn:com.workday/bsvc Email_Comment,omitempty"`
	UsageData    []CommunicationMethodUsageDataType `xml:"urn:com.workday/bsvc Usage_Data,omitempty"`
}

// Email Address Information
type EmailAddressInformationDataType struct {
	EmailAddress    string                                        `xml:"urn:com.workday/bsvc Email_Address,omitempty"`
	EmailComment    string                                        `xml:"urn:com.workday/bsvc Email_Comment,omitempty"`
	UsageData       []CommunicationMethodUsageInformationDataType `xml:"urn:com.workday/bsvc Usage_Data,omitempty"`
	EmailReference  *EmailReferenceObjectType                     `xml:"urn:com.workday/bsvc Email_Reference,omitempty"`
	ID              string                                        `xml:"urn:com.workday/bsvc ID,omitempty"`
	Delete          bool                                          `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	DoNotReplaceAll bool                                          `xml:"urn:com.workday/bsvc Do_Not_Replace_All,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EmailReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmailReferenceObjectType struct {
	ID         []EmailReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EmployeeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmployeeObjectType struct {
	ID         []EmployeeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Reference element representing a unique instance of Employee.
type EmployeeReferenceType struct {
	IntegrationIDReference *ExternalIntegrationIDReferenceDataType `xml:"urn:com.workday/bsvc Integration_ID_Reference"`
}

// Contains a unique identifier for an instance of an object.
type EstablishmentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EstablishmentObjectType struct {
	ID         []EstablishmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Exception (Errors and Warning) associated with the transaction.
type ExceptionDataType struct {
	Classification string `xml:"urn:com.workday/bsvc Classification,omitempty"`
	Message        string `xml:"urn:com.workday/bsvc Message,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ExternalFieldandParameterInitializationProviderObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ExternalFieldandParameterInitializationProviderObjectType struct {
	ID         []ExternalFieldandParameterInitializationProviderObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Integration ID reference is used as a unique identifier for integratable objects in the Workday system.
type ExternalIntegrationIDReferenceDataType struct {
	ID         *IDType `xml:"urn:com.workday/bsvc ID"`
	Descriptor string  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Federal Student Loan Specific Data
type FederalStudentLoanSpecificDataType struct {
	DepartmentofEducationStudentLoan *bool `xml:"urn:com.workday/bsvc Department_of_Education_Student_Loan,omitempty"`
}

// Federal Tax Levy Specific Data
type FederalTaxLevyDependentReferenceType struct {
	DependentName                 string `xml:"urn:com.workday/bsvc Dependent_Name,omitempty"`
	DependentIdentificationNumber string `xml:"urn:com.workday/bsvc Dependent_Identification_Number,omitempty"`
}

// Federal Tax Levy Specific Data
type FederalTaxLevySpecificDataType struct {
	TaxLevySignedPart3Date           *time.Time                             `xml:"urn:com.workday/bsvc Tax_Levy_Signed_Part_3_Date,omitempty"`
	PayrollMaritalStatusReference    []PayrollMaritalStatusReferenceType    `xml:"urn:com.workday/bsvc Payroll_Marital_Status_Reference,omitempty"`
	PersonalExemptions               float64                                `xml:"urn:com.workday/bsvc Personal_Exemptions,omitempty"`
	AdditionalExemptions65orBlind    float64                                `xml:"urn:com.workday/bsvc Additional_Exemptions_65_or_Blind,omitempty"`
	ExemptionAmountOverride          float64                                `xml:"urn:com.workday/bsvc Exemption_Amount_Override,omitempty"`
	ExemptionFrequencyReference      *FrequencyObjectType                   `xml:"urn:com.workday/bsvc Exemption_Frequency_Reference,omitempty"`
	TaxLevyTerminationDate           *time.Time                             `xml:"urn:com.workday/bsvc Tax_Levy_Termination_Date,omitempty"`
	FederalTaxLevyDependentReference []FederalTaxLevyDependentReferenceType `xml:"urn:com.workday/bsvc Federal_Tax_Levy_Dependent_Reference,omitempty"`
	LockTaxElections                 *bool                                  `xml:"urn:com.workday/bsvc Lock_Tax_Elections,omitempty"`
	LoadBaselineRestrictions         *bool                                  `xml:"urn:com.workday/bsvc Load_Baseline_Restrictions,omitempty"`
}

func (t *FederalTaxLevySpecificDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T FederalTaxLevySpecificDataType
	var layout struct {
		*T
		TaxLevySignedPart3Date *xsdDate `xml:"urn:com.workday/bsvc Tax_Levy_Signed_Part_3_Date,omitempty"`
		TaxLevyTerminationDate *xsdDate `xml:"urn:com.workday/bsvc Tax_Levy_Termination_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TaxLevySignedPart3Date = (*xsdDate)(layout.T.TaxLevySignedPart3Date)
	layout.TaxLevyTerminationDate = (*xsdDate)(layout.T.TaxLevyTerminationDate)
	return e.EncodeElement(layout, start)
}
func (t *FederalTaxLevySpecificDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T FederalTaxLevySpecificDataType
	var overlay struct {
		*T
		TaxLevySignedPart3Date *xsdDate `xml:"urn:com.workday/bsvc Tax_Levy_Signed_Part_3_Date,omitempty"`
		TaxLevyTerminationDate *xsdDate `xml:"urn:com.workday/bsvc Tax_Levy_Termination_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TaxLevySignedPart3Date = (*xsdDate)(overlay.T.TaxLevySignedPart3Date)
	overlay.TaxLevyTerminationDate = (*xsdDate)(overlay.T.TaxLevyTerminationDate)
	return d.DecodeElement(&overlay, &start)
}

// Field Result Request Criteria
type FieldAndParameterCriteriaDataType struct {
	ProviderReference []ExternalFieldandParameterInitializationProviderObjectType `xml:"urn:com.workday/bsvc Provider_Reference"`
}

// Contains a unique identifier for an instance of an object.
type FinancialInstitutionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type FinancialInstitutionObjectType struct {
	ID         []FinancialInstitutionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type FrequencyBehaviorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type FrequencyBehaviorObjectType struct {
	ID         []FrequencyBehaviorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type FrequencyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type FrequencyObjectType struct {
	ID         []FrequencyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type FundObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type FundObjectType struct {
	ID         []FundObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Get Canadian Record of Employment Data Request
type GetCanadianRecordofEmploymentDataRequestType struct {
	RequestReferences *CanadianRecordofEmploymentDataRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CanadianRecordofEmploymentDataRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CanadianRecordofEmploymentDataResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Canadian Record of Employment Data Response
type GetCanadianRecordofEmploymentDataResponseType struct {
	RequestReferences *CanadianRecordofEmploymentDataRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CanadianRecordofEmploymentDataRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CanadianRecordofEmploymentDataResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *CanadianRecordofEmploymentDataResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payment Dates for Company and Date Range Request element
type GetCompanyPaymentDatesRequestType struct {
	RequestCriteria *CompanyPaymentDatesRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version         string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payment Dates for Company and Date Range Response element
type GetCompanyPaymentDatesResponseType struct {
	RequestCriteria *CompanyPaymentDatesRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *CompanyPaymentDatesResponseDataType    `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// EFW2 Year End Employer Filing Data Request
type GetEFW2YearEndEmployerFilingDataRequestType struct {
	RequestCriteria *EFW2YearEndEmployerFilingDataRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	Version         string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response Data
type GetEFW2YearEndEmployerFilingDataResponseType struct {
	RequestCriteria *EFW2YearEndEmployerFilingDataRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseData    *EFW2YearEndEmployerFilingResponseDataType        `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// EFW2 Year End Worker Filing Data Request
type GetEFW2YearEndWorkerFilingDataRequestType struct {
	RequestCriteria *EFW2YearEndWorkerFilingDataRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version         string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response Data
type GetEFW2YearEndWorkerFilingDataResponseType struct {
	RequestCriteria *EFW2YearEndWorkerFilingDataRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults *ResponseResultsType                            `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *EFW2YearEndWorkerFilingResponseDataType        `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Monthly Worker Tax Filing Data Request
type GetMonthlyWorkerTaxFilingDataRequestType struct {
	RequestCriteria *MonthlyWorkerTaxFilingDataCriteriaType      `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *MonthlyWorkerTaxFilingDataResponseGroupType `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version         string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Monthly Worker Tax Filing Data Response
type GetMonthlyWorkerTaxFilingDataResponseType struct {
	RequestCriteria *MonthlyWorkerTaxFilingDataCriteriaType      `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *MonthlyWorkerTaxFilingDataResponseGroupType `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults []ResponseResultsType                        `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    []MonthlyWorkerTaxFilingDataResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request for a multiple worksite report.
type GetMultipleWorksiteReportRequestType struct {
	RequestCriteria *MultipleWorksiteReportRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version         string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response for a Get Multiple Worksites Report request.
type GetMultipleWorksiteReportResponseType struct {
	RequestCriteria *MultipleWorksiteReportRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults *ResponseResultsType                       `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *MultipleWorksiteReportResponseDataType    `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The information required to obtain a full or subset of records containing workers Paycheck Delivery and Payslip Printing options.
type GetPaycheckDeliveriesRequestType struct {
	RequestReferences *PaycheckDeliveryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PaycheckDeliveryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PaycheckDeliveryResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Paycheck Deliveries Response Element contains the request and response information for the request.
type GetPaycheckDeliveriesResponseType struct {
	RequestReferences *PaycheckDeliveryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PaycheckDeliveryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PaycheckDeliveryResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PaycheckDeliveryResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Balances Request
type GetPayrollBalancesRequestType struct {
	RequestCriteria *PayrollBalanceRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *PayrollBalanceResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version         string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Balances Response
type GetPayrollBalancesResponseType struct {
	RequestCriteria *PayrollBalanceRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *PayrollBalanceResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults *ResponseResultsType               `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *PayrollBalanceResponseDataType    `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Deduction Recipient wrapper element
type GetPayrollDeductionRecipientsRequestType struct {
	RequestReferences *PayrollDeductionRecipientRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollDeductionRecipientResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element Containing Payroll Deduction Recipient Response
type GetPayrollDeductionRecipientsResponseType struct {
	RequestReferences *PayrollDeductionRecipientRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollDeductionRecipientResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                            `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PayrollDeductionRecipientResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold all Payroll Payee W-4 Request Information
type GetPayrollFederalW4TaxElectionsRequestType struct {
	RequestReferences *PayrollFederalW4TaxElectionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollFederalW4TaxElectionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollFederalW4TaxElectionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Payee W-4s Response Element
type GetPayrollFederalW4TaxElectionsResponseType struct {
	RequestReferences *PayrollFederalW4TaxElectionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollFederalW4TaxElectionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollFederalW4TaxElectionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                              `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PayrollFederalW4TaxElectionResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This public web service method retrieves historical payroll payments that have been loaded into Workday Payroll from External Payroll systems.
type GetPayrollHistoryPaymentsRequestType struct {
	RequestReferences *PaymentHistoryRequestReferencesType     `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollHistoryResultRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response for Get Payroll History Payment Web Service method
type GetPayrollHistoryPaymentsResponseType struct {
	RequestReferences *PaymentHistoryRequestReferencesType   `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PayrollHistoryPaymentResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Involuntary Withholding Orders Request
type GetPayrollInvoluntaryWithholdingOrdersRequestType struct {
	RequestReferences *PayrollInvoluntaryWithholdingOrderRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollInvoluntaryWithholdingOrderRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollInvoluntaryWithholdingOrderResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Involuntary Withholding Orders Response
type GetPayrollInvoluntaryWithholdingOrdersResponseType struct {
	RequestReferences *PayrollInvoluntaryWithholdingOrderRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollInvoluntaryWithholdingOrderRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollInvoluntaryWithholdingOrderResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                                     `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PayrollInvoluntaryWithholdingOrderResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold all Payroll Limit Override Request Information
type GetPayrollLimitOverridesRequestType struct {
	RequestReferences *PayrollLimitOverrideRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollLimitOverrideRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollLimitOverrideResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold the Payroll Limit Override Response.
type GetPayrollLimitOverridesResponseType struct {
	RequestReferences *PayrollLimitOverrideRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollLimitOverrideRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollLimitOverrideResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                       `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PayrollLimitOverrideResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Off-cycle Payments Request element
type GetPayrollOffcyclePaymentsRequestType struct {
	RequestReferences *PayrollOffcycleResultOverrideRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollOffcycleResultOverrideRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Off-cycle Payments Response element
type GetPayrollOffcyclePaymentsResponseType struct {
	RequestReferences *PayrollOffcycleResultOverrideRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollOffcycleResultOverrideRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                                `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PayrollOffcycleResultOverrideResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Payee FICAs Request element
type GetPayrollPayeeFICAsRequestType struct {
	RequestReferences *PayrollPayeeTaxDataRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollPayeeTaxDataRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollPayeeTaxDataResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Payee FICAs Response element
type GetPayrollPayeeFICAsResponseType struct {
	RequestReferences *PayrollPayeeTaxDataRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollPayeeTaxDataRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollPayeeTaxDataResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                      `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PayrollPayeeTaxDataResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The main FUTA Request
type GetPayrollPayeeFUTAsRequestType struct {
	RequestReferences *PayrollPayeeFUTARequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollPayeeFUTARequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollPayeeFUTAResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Payee FUTAs Response Element
type GetPayrollPayeeFUTAsResponseType struct {
	RequestReferences *PayrollPayeeFUTARequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollPayeeFUTARequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollPayeeFUTAResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PayrollPayeeFUTAResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold all Payroll Payee Ongoing Work Jurisdiction Tax Election Request Information
type GetPayrollPayeeOngoingWorkJurisdictionTaxElectionRequestType struct {
	RequestReferences *PayrollPayeeOngoingWorkJurisdictionTaxElectionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollPayeeOngoingWorkJurisdictionTaxElectionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold Get Payroll Payee Ongoing Work Jurisdiction Tax Election Response Information
type GetPayrollPayeeOngoingWorkJurisdictionTaxElectionResponseType struct {
	RequestReferences *PayrollPayeeOngoingWorkJurisdictionTaxElectionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollPayeeOngoingWorkJurisdictionTaxElectionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                                                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PayrollPayeeOngoingWorkJurisdictionTaxElectionResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Payee PT1s Request
type GetPayrollPayeePT1sRequestType struct {
	RequestReferences *PayrollPayeePT1RequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollPayeePT1RequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollPayeePT1ResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Payee PT1s Response
type GetPayrollPayeePT1sResponseType struct {
	RequestReferences *PayrollPayeePT1RequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollPayeePT1RequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollPayeePT1ResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                  `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PayrollPayeePT1ResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Payee RPP or DPSP Registration Numbers Request
type GetPayrollPayeeRPPorDPSPRegistrationNumbersRequestType struct {
	RequestReferences *PayrollPayeeRPPorDPSPRegistrationNumberRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                           `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollPayeeRPPorDPSPRegistrationNumberResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Payee RPP or DPSP Registration Numbers Response Element
type GetPayrollPayeeRPPorDPSPRegistrationNumbersResponseType struct {
	RequestReferences *PayrollPayeeRPPorDPSPRegistrationNumberRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                           `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollPayeeRPPorDPSPRegistrationNumberResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                                          `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PayrollPayeeRPPorDPSPRegistrationNumberResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Payee TD1s Request
type GetPayrollPayeeTD1sRequestType struct {
	RequestReferences *PayrollPayeeTD1RequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollPayeeTD1RequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollPayeeTD1ResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Payee TD1s Response
type GetPayrollPayeeTD1sResponseType struct {
	RequestReferences *PayrollPayeeTD1RequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollPayeeTD1RequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollPayeeTD1ResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                  `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PayrollPayeeTD1ResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold all Payroll Payee Tax Location Mappings Request Information
type GetPayrollPayeeTaxLocationMappingsRequestType struct {
	RequestReferences *PayrollPayeeTaxLocationMappingRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollPayeeTaxLocationMappingRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold Get Payroll Payee Tax Location Mapping Response Information
type GetPayrollPayeeTaxLocationMappingsResponseType struct {
	RequestReferences *PayrollPayeeTaxLocationMappingRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollPayeeTaxLocationMappingRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PayrollPayeeTaxLocationMappingResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Payslips Request
type GetPayrollPayslipsRequestType struct {
	RequestCriteria *PayrollPayslipRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *PayrollPayslipResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version         string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Payslips Response
type GetPayrollPayslipsResponseType struct {
	RequestCriteria *PayrollPayslipRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *PayrollPayslipResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults *ResponseResultsType               `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *PayrollPayslipResponseDataType    `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Pre-Printed Payslips Request
type GetPayrollPrePrintedPayslipsRequestType struct {
	RequestReference *PayrollPrePrintedPayslipsRequestReferenceType `xml:"urn:com.workday/bsvc Request_Reference"`
	RequestCriteria  *PayrollPrePrintedPayslipsRequestCriteriaType  `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter   *ResponseFilterType                            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup    *PayrollPrePrintedPayslipsResponseGroupType    `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version          string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element that includes the Repository Document instance and data
type GetPayrollPrePrintedPayslipsResponseDataType struct {
	EmployeeNameReference                     *WorkerObjectType                `xml:"urn:com.workday/bsvc Employee_Name_Reference,omitempty"`
	PayGroupDetailReference                   *PayGroupDetailObjectType        `xml:"urn:com.workday/bsvc Pay_Group_Detail_Reference,omitempty"`
	CompanyReference                          *CompanyObjectType               `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	PayrollResultReference                    *UniqueIdentifierObjectType      `xml:"urn:com.workday/bsvc Payroll_Result_Reference,omitempty"`
	PeriodStartDate                           *time.Time                       `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
	PeriodEndDate                             *time.Time                       `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	PaymentDate                               *time.Time                       `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	GrossAmount                               float64                          `xml:"urn:com.workday/bsvc Gross_Amount,omitempty"`
	NetAmount                                 float64                          `xml:"urn:com.workday/bsvc Net_Amount,omitempty"`
	PayrollPrePrintedPayslipDocumentID        string                           `xml:"urn:com.workday/bsvc Payroll_Pre-Printed_Payslip_Document_ID,omitempty"`
	PayrollPrePrintedPayslipDocumentReference []RepositoryDocumentMetadataType `xml:"urn:com.workday/bsvc Payroll_Pre-Printed_Payslip_Document_Reference,omitempty"`
	FileContent                               *[]byte                          `xml:"urn:com.workday/bsvc File_Content,omitempty"`
}

func (t *GetPayrollPrePrintedPayslipsResponseDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T GetPayrollPrePrintedPayslipsResponseDataType
	var layout struct {
		*T
		PeriodStartDate *xsdDate         `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
		PeriodEndDate   *xsdDate         `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
		PaymentDate     *xsdDate         `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
		FileContent     *xsdBase64Binary `xml:"urn:com.workday/bsvc File_Content,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PeriodStartDate = (*xsdDate)(layout.T.PeriodStartDate)
	layout.PeriodEndDate = (*xsdDate)(layout.T.PeriodEndDate)
	layout.PaymentDate = (*xsdDate)(layout.T.PaymentDate)
	layout.FileContent = (*xsdBase64Binary)(layout.T.FileContent)
	return e.EncodeElement(layout, start)
}
func (t *GetPayrollPrePrintedPayslipsResponseDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T GetPayrollPrePrintedPayslipsResponseDataType
	var overlay struct {
		*T
		PeriodStartDate *xsdDate         `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
		PeriodEndDate   *xsdDate         `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
		PaymentDate     *xsdDate         `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
		FileContent     *xsdBase64Binary `xml:"urn:com.workday/bsvc File_Content,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PeriodStartDate = (*xsdDate)(overlay.T.PeriodStartDate)
	overlay.PeriodEndDate = (*xsdDate)(overlay.T.PeriodEndDate)
	overlay.PaymentDate = (*xsdDate)(overlay.T.PaymentDate)
	overlay.FileContent = (*xsdBase64Binary)(overlay.T.FileContent)
	return d.DecodeElement(&overlay, &start)
}

// Get Payroll Pre-Printed Payslips Response
type GetPayrollPrePrintedPayslipsResponseType struct {
	RequestReferences []PayrollPrePrintedPayslipsRequestReferenceType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   []PayrollPrePrintedPayslipsRequestCriteriaType  `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseGroup     []PayrollPrePrintedPayslipsResponseGroupType    `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	RequestFilter     []ResponseFilterType                            `xml:"urn:com.workday/bsvc Request_Filter,omitempty"`
	ResponseResults   []ResponseResultsType                           `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []GetPayrollPrePrintedPayslipsResponseDataType  `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request Element to return workers and their associated payroll reporting codes as of an effective date.
type GetPayrollReportingCodesforWorkersRequestType struct {
	RequestReferences *PayrollReportingCodeForWorkersRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing request response, including paging details and request references.
type GetPayrollReportingCodesforWorkersResponseType struct {
	RequestReferences []PayrollReportingCodeForWorkersRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    []ResponseFilterType                                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   []ResponseResultsType                                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []PayrollReportingCodesforWorkersResponseDataType     `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Results Request
type GetPayrollResultsRequestType struct {
	RequestReferences *PayrollResultRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollResultRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollResultResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payroll Results Response
type GetPayrollResultsResponseType struct {
	RequestReferences *PayrollResultRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollResultRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollResultResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PayrollResultResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold all Payroll Tax Mapping on Location Request information
type GetPayrollTaxMappingsonLocationRequestType struct {
	RequestReferences *PayrollTaxMappingsonLocationRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollTaxMappingsonLocationRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Documentation Element to hold Get Payroll Tax Mappings on Location Response information
type GetPayrollTaxMappingsonLocationResponseType struct {
	RequestReference []PayrollTaxMappingsonLocationRequestReferencesType `xml:"urn:com.workday/bsvc Request_Reference,omitempty"`
	RequestCriteria  []PayrollTaxMappingsonLocationRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter   []ResponseFilterType                                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults  []ResponseResultsType                               `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData     []PayrollTaxMappingsonLocationResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version          string                                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request for Get Payroll USA State and Local Tax Elections
type GetPayrollUSAStateandLocalTaxElectionsRequestType struct {
	RequestCriteria *PayrollUSAStateandLocalTaxElectionRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *PayrollUSAStateandLocalTaxElectionResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version         string                                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Root element for the Response on the Get operation. Contains the instances returned by the Get operation and their accompanying data.
type GetPayrollUSAStateandLocalTaxElectionsResponseType struct {
	RequestCriteria *PayrollUSAStateandLocalTaxElectionRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *PayrollUSAStateandLocalTaxElectionResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults *ResponseResultsType                                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *PayrollStateandLocalTaxElectionResponseDataType       `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Worker Tax Treaties Request
type GetPayrollWorkerTaxTreatiesRequestType struct {
	RequestReferences *WorkerTaxTreatiesRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *WorkerTaxTreatiesRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Period Schedules Request
type GetPeriodSchedulesRequestType struct {
	RequestReferences *PeriodScheduleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PeriodScheduleRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PeriodScheduleResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Period Schedules Response
type GetPeriodSchedulesResponseType struct {
	RequestReferences *PeriodScheduleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PeriodScheduleRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PeriodScheduleResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PeriodScheduleResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Periodic Company CAN Tax Remittance Data Request element
type GetPeriodicCompanyCANTaxRemittanceDataRequestType struct {
	RequestCriteria *PeriodicCANTaxRemittanceDataRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *PeriodicCANTaxRemittanceDataResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version         string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Periodic Company CAN Tax Remittance Data Response element
type GetPeriodicCompanyCANTaxRemittanceDataResponseType struct {
	RequestCriteria *PeriodicCANTaxRemittanceDataRequestCriteriaType     `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *PeriodicCANTaxRemittanceDataResponseGroupType       `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults *ResponseResultsType                                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *PeriodicCompanyCANTaxRemittanceDataResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Periodic Company Tax Filing Data Request element
type GetPeriodicCompanyTaxFilingDataRequestType struct {
	RequestCriteria *PeriodicTaxFilingDataRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *PeriodicTaxFilingDataResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version         string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Periodic Company Tax Filing Data Response element
type GetPeriodicCompanyTaxFilingDataResponseType struct {
	RequestCriteria *PeriodicTaxFilingDataRequestCriteriaType     `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                           `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *PeriodicTaxFilingDataResponseGroupType       `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults *ResponseResultsType                          `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *PeriodicCompanyTaxFilingDataResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Periodic Worker Tax Filing Data
type GetPeriodicWorkerTaxFilingDataRequestType struct {
	RequestCriteria *PeriodicTaxFilingDataRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *PeriodicTaxFilingDataResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version         string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Periodic Worker Tax Filing Data Response
type GetPeriodicWorkerTaxFilingDataResponseType struct {
	RequestCriteria *PeriodicTaxFilingDataRequestCriteriaType    `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *PeriodicTaxFilingDataResponseGroupType      `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults *ResponseResultsType                         `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *PeriodicWorkerTaxFilingDataResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Quarterly Worker Request
type GetQuarterlyWorkerTaxDataRequestType struct {
	RequestCriteria *QuarterlyWorkerTaxDataRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *QuarterlyWorkerTaxDataResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version         string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response Data
type GetQuarterlyWorkerTaxDataResponseType struct {
	RequestCriteria *QuarterlyWorkerTaxDataRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria"`
	ResponseFilter  *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *QuarterlyWorkerTaxDataResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults *ResponseResultsType                       `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *QuarterlyWorkerTaxDataResponseDataType    `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Root element for the Request on the Get operation. Contains references, criteria, filter and response group for specifying which instances to return in the Response.
type GetROEHistoryDataRequestType struct {
	RequestReferences *ROEHistoryDataRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *ROEHistoryDataRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *ROEHistoryDataResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get ROE Data History Response
type GetROEHistoryDataResponseType struct {
	RequestReferences *ROEHistoryDataRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *ROEHistoryDataRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *ROEHistoryDataResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *ROEHistoryDataResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Run Pay Calculation Request
type GetRunPayCalculationRequestType struct {
	RequestReferences *RunPayCalculationRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *RunPayCalculationRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *RunPayCalculationResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Run Pay Calculation Response
type GetRunPayCalculationResponseType struct {
	RequestReferences *RunPayCalculationRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *RunPayCalculationRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *RunPayCalculationResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *RunPayCalculationResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Single Legal Entities Request
type GetSingleLegalEntitiesRequestType struct {
	RequestReferences *SingleLegalEntityRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *SingleLegalEntityResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Single Legal Entities Response
type GetSingleLegalEntitiesResponseType struct {
	RequestReferences *SingleLegalEntityRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *SingleLegalEntityResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *SingleLegalEntityResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Submit Payroll Inputs Request
type GetSubmitPayrollInputsRequestType struct {
	RequestReferences *PayrollInputRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollInputRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollInputResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Submit Payroll Inputs Response
type GetSubmitPayrollInputsResponseType struct {
	RequestReferences *PayrollInputRequestReferencesType  `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PayrollInputRequestCriteriaType    `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PayrollInputResponseGroupType      `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *SubmitPayrollInputResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Successor Employers Request
type GetSuccessorEmployersRequestType struct {
	RequestReferences *SuccessorEmployerRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *SuccessorEmployerResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Successor Employers Response
type GetSuccessorEmployersResponseType struct {
	RequestReferences *SuccessorEmployerRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *SuccessorEmployerResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *SuccessorEmployerResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This web service is designed to get the payroll year end form printing elections for Worker(s). These elections determine whether the worker will receive a paper copy of their payroll year end form(s).
type GetTaxDocumentDeliveriesRequestType struct {
	RequestReferences *TaxDocumentDeliveryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *TaxDocumentDeliveryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *TaxDocumentDeliveryResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Tax Document Deliveries Response
type GetTaxDocumentDeliveriesResponseType struct {
	RequestReferences *TaxDocumentDeliveryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *TaxDocumentDeliveryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *TaxDocumentDeliveryResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                      `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *TaxDocumentDeliveryResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Tax Levy Deduction Restrictions Request
type GetTaxLevyDeductionRestrictionsRequestType struct {
	RequestReferences *TaxLevyDeductionRestrictionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *TaxLevyDeductionRestrictionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *TaxLevyDeductionRestrictionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Tax Levy Deduction Restrictions Response
type GetTaxLevyDeductionRestrictionsResponseType struct {
	RequestReferences *TaxLevyDeductionRestrictionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *TaxLevyDeductionRestrictionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *TaxLevyDeductionRestrictionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                              `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *TaxLevyDeductionRestrictionsResponseDataType     `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold all W-2/W-2C Printing Election Request Information
type GetW2W2CPrintingElectionRequestType struct {
	RequestReferences *W2W2CPrintingElectionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *W2W2CPrintingElectionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *W2W2CPrintingElectionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get W-2/W-2C Printing Election Response Element
type GetW2W2CPrintingElectionResponseType struct {
	RequestReferences *W2W2CPrintingElectionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *W2W2CPrintingElectionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *W2W2CPrintingElectionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                        `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *W2W2CPrintingElectionResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Withholding Order Local Minimum Wage Response Data
type GetWithholdingOrderLocalMinimumWageDataType struct {
	LocalMinimumWageReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Local_Minimum_Wage_Reference,omitempty"`
	WorkerReference           *WorkerObjectType           `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	CompanyReference          *CompanyObjectType          `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	WithholdingOrderReference *WithholdingOrderObjectType `xml:"urn:com.workday/bsvc Withholding_Order_Reference,omitempty"`
	EffectiveDate             *time.Time                  `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	LocalMinimumWageRate      float64                     `xml:"urn:com.workday/bsvc Local_Minimum_Wage_Rate,omitempty"`
	Comment                   string                      `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

func (t *GetWithholdingOrderLocalMinimumWageDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T GetWithholdingOrderLocalMinimumWageDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *GetWithholdingOrderLocalMinimumWageDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T GetWithholdingOrderLocalMinimumWageDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Get Withholding Order Local Minimum Wage Rate Request
type GetWithholdingOrderLocalMinimumWageRateRequestType struct {
	RequestCriteria *WithholdingOrderLocalMinimumWageRateRequestType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version         string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Withholding Order Local Minimum Wage Rate Response
type GetWithholdingOrderLocalMinimumWageRateResponseType struct {
	RequestCriteria *WithholdingOrderLocalMinimumWageRateRequestType     `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	RequestFilter   *ResponseFilterType                                  `xml:"urn:com.workday/bsvc Request_Filter,omitempty"`
	ResponseResults *ResponseResultsType                                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *GetWithholdingOrderLocalMinimumWageResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Withholding Order Local Minimum Wage Rate Response
type GetWithholdingOrderLocalMinimumWageResponseDataType struct {
	WithholdingOrderLocalMinimumWageData []GetWithholdingOrderLocalMinimumWageDataType `xml:"urn:com.workday/bsvc Withholding_Order_Local_Minimum_Wage_Data,omitempty"`
}

// Request criteria (specific and universal) for the Get Worker Costing Allocations service.
type GetWorkerCostingAllocationsRequestType struct {
	RequestCriteria *WorkerCostingAllocationsRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *WorkerCostingAllocationsResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version         string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper for the data returned by the Get Worker Costing Allocations service.
type GetWorkerCostingAllocationsResponseType struct {
	RequestCriteria *WorkerCostingAllocationsRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *WorkerCostingAllocationsResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults *ResponseResultsType                         `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *WorkerCostingAllocationsResponseDataType    `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Worker Tax Treaties Response
type GetWorkerTaxTreatiesResponseType struct {
	RequestReferences *WorkerTaxTreatiesRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *WorkerTaxTreatiesRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *WorkerTaxTreatiesResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type GiftObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type GiftObjectType struct {
	ID         []GiftObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type GrantObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type GrantObjectType struct {
	ID         []GrantObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The History Payment Data element contains attributes specific to this payment, such as the Gross and Net amounts, an optional Sub-period definition (if the payment is not for the full payroll Period) and optional Third Party Sick Pay indicator.
type HistoryPaymentDataType struct {
	ThirdPartySickPay  *bool      `xml:"urn:com.workday/bsvc Third_Party_Sick_Pay,omitempty"`
	PaymentDate        time.Time  `xml:"urn:com.workday/bsvc Payment_Date"`
	GrossAmount        float64    `xml:"urn:com.workday/bsvc Gross_Amount,omitempty"`
	NetAmount          float64    `xml:"urn:com.workday/bsvc Net_Amount,omitempty"`
	SubPeriodStartDate *time.Time `xml:"urn:com.workday/bsvc Sub-Period_Start_Date,omitempty"`
	SubPeriodEndDate   *time.Time `xml:"urn:com.workday/bsvc Sub-Period_End_Date,omitempty"`
}

func (t *HistoryPaymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T HistoryPaymentDataType
	var layout struct {
		*T
		PaymentDate        *xsdDate `xml:"urn:com.workday/bsvc Payment_Date"`
		SubPeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Sub-Period_Start_Date,omitempty"`
		SubPeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Sub-Period_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PaymentDate = (*xsdDate)(&layout.T.PaymentDate)
	layout.SubPeriodStartDate = (*xsdDate)(layout.T.SubPeriodStartDate)
	layout.SubPeriodEndDate = (*xsdDate)(layout.T.SubPeriodEndDate)
	return e.EncodeElement(layout, start)
}
func (t *HistoryPaymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T HistoryPaymentDataType
	var overlay struct {
		*T
		PaymentDate        *xsdDate `xml:"urn:com.workday/bsvc Payment_Date"`
		SubPeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Sub-Period_Start_Date,omitempty"`
		SubPeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Sub-Period_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDate = (*xsdDate)(&overlay.T.PaymentDate)
	overlay.SubPeriodStartDate = (*xsdDate)(overlay.T.SubPeriodStartDate)
	overlay.SubPeriodEndDate = (*xsdDate)(overlay.T.SubPeriodEndDate)
	return d.DecodeElement(&overlay, &start)
}

// This element contains the earnings and deductions detail that contributes to the history payment's Gross and Net.
type HistoryPaymentInputDataType struct {
	EarningReference            *EarningAllObjectType             `xml:"urn:com.workday/bsvc Earning_Reference"`
	DeductionReference          *DeductionAllObjectType           `xml:"urn:com.workday/bsvc Deduction_Reference"`
	PositionReference           *PositionElementObjectType        `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	Amount                      float64                           `xml:"urn:com.workday/bsvc Amount,omitempty"`
	TaxableWages                float64                           `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages                float64                           `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	PayrollInputWorktagsData    *PayrollInputWorktagsDataType     `xml:"urn:com.workday/bsvc Payroll_Input_Worktags_Data,omitempty"`
	Comment                     string                            `xml:"urn:com.workday/bsvc Comment,omitempty"`
	HistoryPaymentInputLineData []HistoryPaymentInputLineDataType `xml:"urn:com.workday/bsvc History_Payment_Input_Line_Data,omitempty"`
}

// This element contains the value details (amount, hours, rate, etc.) that contribute to an earning or deduction.
type HistoryPaymentInputLineDataType struct {
	RelatedCalculationReference *RelatedCalculationAllObjectType `xml:"urn:com.workday/bsvc Related_Calculation_Reference"`
	InputValue                  float64                          `xml:"urn:com.workday/bsvc Input_Value,omitempty"`
}

// External ID that uniquely identifies the integratable object within the context of the integration system identified by the System ID attribute.
type IDType struct {
	Value    string `xml:",chardata"`
	SystemID string `xml:"urn:com.workday/bsvc System_ID,attr,omitempty"`
}

// Header element of the Import Maintain Payroll Reporting Codes for Worker Events web service task.
type ImportMaintainPayrollReportingCodesforWorkerRequestType struct {
	WorkerPayrollReportingCodeData []MaintainWorkerPayrollReportingCodesDataHVType `xml:"urn:com.workday/bsvc Worker_Payroll_Reporting_Code_Data,omitempty"`
	Version                        string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for Import Payroll History Payment WSBP.
type ImportPayrollHistoryPaymentRequestType struct {
	PayrollHistoryPaymentData []PayrollHistoryPaymentDataType `xml:"urn:com.workday/bsvc Payroll_History_Payment_Data,omitempty"`
	Version                   string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Import Payroll Input Request
type ImportPayrollInputRequestType struct {
	DefaultBatchID   string                       `xml:"urn:com.workday/bsvc Default_Batch_ID"`
	PayrollInputData []SubmitPayrollInputDataType `xml:"urn:com.workday/bsvc Payroll_Input_Data,omitempty"`
	Version          string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This is the request element for the import payroll off-cycle payment web service background process.
type ImportPayrollOffcyclePaymentRequestType struct {
	PayrollOffcyclePaymentData []PayrollOffcyclePaymentDataType `xml:"urn:com.workday/bsvc Payroll_Off-cycle_Payment_Data,omitempty"`
	Version                    string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Import ROE Prior Period History Results Request
type ImportROEPriorPeriodHistoryResultsRequestType struct {
	PayrollROEPriorPeriodHistoryResultsData []PayrollROEPriorHistoryResultsDataType `xml:"urn:com.workday/bsvc Payroll_ROE_Prior_Period_History_Results_Data,omitempty"`
	Version                                 string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Instant Messenger Address.
type InstantMessengerInformationDataType struct {
	InstantMessengerAddress       string                                        `xml:"urn:com.workday/bsvc Instant_Messenger_Address,omitempty"`
	InstantMessengerTypeReference *InstantMessengerTypeObjectType               `xml:"urn:com.workday/bsvc Instant_Messenger_Type_Reference,omitempty"`
	InstantMessengerComment       string                                        `xml:"urn:com.workday/bsvc Instant_Messenger_Comment,omitempty"`
	UsageData                     []CommunicationMethodUsageInformationDataType `xml:"urn:com.workday/bsvc Usage_Data,omitempty"`
	InstantMessengerReference     *InstantMessengerReferenceObjectType          `xml:"urn:com.workday/bsvc Instant_Messenger_Reference,omitempty"`
	ID                            string                                        `xml:"urn:com.workday/bsvc ID,omitempty"`
	Delete                        bool                                          `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	DoNotReplaceAll               bool                                          `xml:"urn:com.workday/bsvc Do_Not_Replace_All,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type InstantMessengerReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InstantMessengerReferenceObjectType struct {
	ID         []InstantMessengerReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type InstantMessengerTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InstantMessengerTypeObjectType struct {
	ID         []InstantMessengerTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type IntegrationSystemAuditedObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntegrationSystemAuditedObjectType struct {
	ID         []IntegrationSystemAuditedObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Intermediary Bank Account information including Order, Bank Account Type, Currency, Routing Transit Number, etc.
type IntermediaryBankAccountWWSDataType struct {
	IntermediaryBankAccountID   string                     `xml:"urn:com.workday/bsvc Intermediary_Bank_Account_ID,omitempty"`
	IntermediaryBankOrder       string                     `xml:"urn:com.workday/bsvc Intermediary_Bank_Order,omitempty"`
	CountryReference            *CountryObjectType         `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CurrencyReference           *CurrencyObjectType        `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	BankAccountNickname         string                     `xml:"urn:com.workday/bsvc Bank_Account_Nickname,omitempty"`
	BankAccountTypeReference    *BankAccountTypeObjectType `xml:"urn:com.workday/bsvc Bank_Account_Type_Reference,omitempty"`
	BankName                    string                     `xml:"urn:com.workday/bsvc Bank_Name,omitempty"`
	RoutingTransitNumber        string                     `xml:"urn:com.workday/bsvc Routing_Transit_Number,omitempty"`
	BranchID                    string                     `xml:"urn:com.workday/bsvc Branch_ID,omitempty"`
	BranchName                  string                     `xml:"urn:com.workday/bsvc Branch_Name,omitempty"`
	BankAccountNumber           string                     `xml:"urn:com.workday/bsvc Bank_Account_Number,omitempty"`
	CheckDigit                  string                     `xml:"urn:com.workday/bsvc Check_Digit,omitempty"`
	BankAccountName             string                     `xml:"urn:com.workday/bsvc Bank_Account_Name,omitempty"`
	RollNumber                  string                     `xml:"urn:com.workday/bsvc Roll_Number,omitempty"`
	IBAN                        string                     `xml:"urn:com.workday/bsvc IBAN,omitempty"`
	SWIFTBankIdentificationCode string                     `xml:"urn:com.workday/bsvc SWIFT_Bank_Identification_Code,omitempty"`
	Inactive                    *bool                      `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	BankInstructions            string                     `xml:"urn:com.workday/bsvc Bank_Instructions,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JobProfileObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobProfileObjectType struct {
	ID         []JobProfileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type KindofEmployerforPayrollTaxFilingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type KindofEmployerforPayrollTaxFilingObjectType struct {
	ID         []KindofEmployerforPayrollTaxFilingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element for all Last Name data.
type LastNameDataType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc Type,attr,omitempty"`
}

// Limit Interface for the Payroll Limit Override
type LimitInterfaceDataType struct {
	PayComponentReference          *PayComponentObjectType          `xml:"urn:com.workday/bsvc Pay_Component_Reference"`
	RelatedCalculationAllReference *RelatedCalculationAllObjectType `xml:"urn:com.workday/bsvc Related_Calculation__All__Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LimitInterfaceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LimitInterfaceObjectType struct {
	ID         []LimitInterfaceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element for all Local Last Name data, for countries supporting Last Name or Secondary Last Name in local script.
type LocalLastNameDataType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc Type,attr,omitempty"`
}

// Contains the components of a name in local script, such as the First Name and Last Name, for supporting countries.
type LocalNameDataType struct {
	FirstName   string                  `xml:"urn:com.workday/bsvc First_Name,omitempty"`
	MiddleName  string                  `xml:"urn:com.workday/bsvc Middle_Name,omitempty"`
	LastName    []LocalLastNameDataType `xml:"urn:com.workday/bsvc Last_Name,omitempty"`
	LocalName   string                  `xml:"urn:com.workday/bsvc Local_Name,attr,omitempty"`
	LocalScript string                  `xml:"urn:com.workday/bsvc Local_Script,attr,omitempty"`
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

// Monthly Tax Filing Deduction Data
type MTDTaxFilingDataType struct {
	TaxWithheld  float64 `xml:"urn:com.workday/bsvc Tax_Withheld,omitempty"`
	TaxableWages float64 `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages float64 `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	GrossWages   float64 `xml:"urn:com.workday/bsvc Gross_Wages,omitempty"`
}

// MTD Tax Remittance Data for Periodic Company
type MTDTaxRemittanceDataforPeriodicCompanyType struct {
	TaxWithheld       float64 `xml:"urn:com.workday/bsvc Tax_Withheld,omitempty"`
	TaxableWages      float64 `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages      float64 `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	ExemptWages       float64 `xml:"urn:com.workday/bsvc Exempt_Wages,omitempty"`
	WCBInsurableWages float64 `xml:"urn:com.workday/bsvc WCB_Insurable_Wages,omitempty"`
	WCBGrossWages     float64 `xml:"urn:com.workday/bsvc WCB_Gross_Wages,omitempty"`
	WCBOtherWages     float64 `xml:"urn:com.workday/bsvc WCB_Other_Wages,omitempty"`
	WCBExcessWages    float64 `xml:"urn:com.workday/bsvc WCB_Excess_Wages,omitempty"`
}

// Element containing singular Worker/Position, Effective Date, and Payroll Reporting Codes.
type MaintainWorkerPayrollReportingCodesDataHVType struct {
	EmployeeReference        *EmployeeObjectType                                  `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	PositionReference        *PositionElementObjectType                           `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	EffectiveDate            time.Time                                            `xml:"urn:com.workday/bsvc Effective_Date"`
	PayrollReportingCodeData []MaintainWorkerPayrollReportingCodesEventDataHVType `xml:"urn:com.workday/bsvc Payroll_Reporting_Code_Data,omitempty"`
}

func (t *MaintainWorkerPayrollReportingCodesDataHVType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MaintainWorkerPayrollReportingCodesDataHVType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *MaintainWorkerPayrollReportingCodesDataHVType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MaintainWorkerPayrollReportingCodesDataHVType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Element containing a single Payroll Reporting Code Entry.
type MaintainWorkerPayrollReportingCodesEventDataHVType struct {
	PayrollReportingCodeReference *PayrollReportingCodeAllObjectType `xml:"urn:com.workday/bsvc Payroll_Reporting_Code_Reference,omitempty"`
}

// Manual Payment Data
type ManualPaymentDataType struct {
	ThirdPartySickPay    *bool                  `xml:"urn:com.workday/bsvc Third_Party_Sick_Pay,omitempty"`
	NetAmount            float64                `xml:"urn:com.workday/bsvc Net_Amount,omitempty"`
	CheckNumber          string                 `xml:"urn:com.workday/bsvc Check_Number,omitempty"`
	BankAccountReference *BankAccountObjectType `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
}

// Medicare Exempt Data
type MedicareExemptDataType struct {
	ExemptfromMedicare                  *bool                                    `xml:"urn:com.workday/bsvc Exempt_from_Medicare,omitempty"`
	MedicareReasonforExemptionReference []PayrollPayeeFICAExemptReasonObjectType `xml:"urn:com.workday/bsvc Medicare_Reason_for_Exemption_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MetadataPayrollWorktagObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MetadataPayrollWorktagObjectType struct {
	ID         []MetadataPayrollWorktagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MonthObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MonthObjectType struct {
	ID         []MonthObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Monthly Worker Tax Filing Data
type MonthlyTaxFilingDeductionDataType struct {
	DeductionReference           *PayrollCalculationObjectType      `xml:"urn:com.workday/bsvc Deduction_Reference"`
	PayrollTaxAuthorityReference []MetadataPayrollWorktagObjectType `xml:"urn:com.workday/bsvc Payroll_Tax_Authority_Reference"`
	TaxFilingCode                string                             `xml:"urn:com.workday/bsvc Tax_Filing_Code,omitempty"`
	MTDData                      *MTDTaxFilingDataType              `xml:"urn:com.workday/bsvc MTD_Data,omitempty"`
	YTDData                      *YTDTaxFilingDataType              `xml:"urn:com.workday/bsvc YTD_Data,omitempty"`
}

// Monthly Worker Tax Filing Data
type MonthlyWorkerTaxFilingDataCriteriaType struct {
	CompanyReference      *CompanyObjectType      `xml:"urn:com.workday/bsvc Company_Reference"`
	CalendarYearReference *CalendarYearObjectType `xml:"urn:com.workday/bsvc Calendar_Year_Reference"`
	MonthReference        *MonthObjectType        `xml:"urn:com.workday/bsvc Month_Reference"`
	AsOfDate              *time.Time              `xml:"urn:com.workday/bsvc As_Of_Date,omitempty"`
}

func (t *MonthlyWorkerTaxFilingDataCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MonthlyWorkerTaxFilingDataCriteriaType
	var layout struct {
		*T
		AsOfDate *xsdDate `xml:"urn:com.workday/bsvc As_Of_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AsOfDate = (*xsdDate)(layout.T.AsOfDate)
	return e.EncodeElement(layout, start)
}
func (t *MonthlyWorkerTaxFilingDataCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MonthlyWorkerTaxFilingDataCriteriaType
	var overlay struct {
		*T
		AsOfDate *xsdDate `xml:"urn:com.workday/bsvc As_Of_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AsOfDate = (*xsdDate)(overlay.T.AsOfDate)
	return d.DecodeElement(&overlay, &start)
}

// Get Monthly Worker Tax Filing Data Response
type MonthlyWorkerTaxFilingDataResponseDataType struct {
	MonthlyWorkerTaxFilingData []MonthlyWorkerTaxFilingDataType `xml:"urn:com.workday/bsvc Monthly_Worker_Tax_Filing_Data,omitempty"`
}

// Monthly Worker Tax Filing Data Response Group
type MonthlyWorkerTaxFilingDataResponseGroupType struct {
	IncludeYTDData *bool `xml:"urn:com.workday/bsvc Include_YTD_Data,omitempty"`
}

// Monthly Worker Tax Filing Data
type MonthlyWorkerTaxFilingDataType struct {
	WorkerReference          *WorkerObjectType                    `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	CompanyReference         *CompanyObjectType                   `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	CalendarYearReference    *CalendarYearObjectType              `xml:"urn:com.workday/bsvc Calendar_Year_Reference,omitempty"`
	MonthReference           *MonthObjectType                     `xml:"urn:com.workday/bsvc Month_Reference,omitempty"`
	StartDate                *time.Time                           `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                  *time.Time                           `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	CurrencyReference        *CurrencyObjectType                  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	WeeksWorked              float64                              `xml:"urn:com.workday/bsvc Weeks_Worked,omitempty"`
	HoursWorked              float64                              `xml:"urn:com.workday/bsvc Hours_Worked,omitempty"`
	SeasonalWorker           *bool                                `xml:"urn:com.workday/bsvc Seasonal_Worker,omitempty"`
	CorporateOfficer         *bool                                `xml:"urn:com.workday/bsvc Corporate_Officer,omitempty"`
	PayrollReportingCodeData []WorkerPayrollReportingCodeDataType `xml:"urn:com.workday/bsvc Payroll_Reporting_Code_Data,omitempty"`
	MonthlyDeductionTaxData  []MonthlyTaxFilingDeductionDataType  `xml:"urn:com.workday/bsvc Monthly_Deduction_Tax_Data,omitempty"`
}

func (t *MonthlyWorkerTaxFilingDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MonthlyWorkerTaxFilingDataType
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
func (t *MonthlyWorkerTaxFilingDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MonthlyWorkerTaxFilingDataType
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

// Data for the multiple worksite report.
type MultipleWorksiteReportDataType struct {
	CompanyReference                *CompanyObjectType                    `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	CompanyLegalName                string                                `xml:"urn:com.workday/bsvc Company_Legal_Name,omitempty"`
	CompanyFEIN                     string                                `xml:"urn:com.workday/bsvc Company_FEIN,omitempty"`
	CalendarYearReference           *CalendarYearObjectType               `xml:"urn:com.workday/bsvc Calendar_Year_Reference,omitempty"`
	CalendarQuarterReference        *CalendarQuarterObjectType            `xml:"urn:com.workday/bsvc Calendar_Quarter_Reference,omitempty"`
	QuarterNumber                   float64                               `xml:"urn:com.workday/bsvc Quarter_Number,omitempty"`
	MultipleWorksiteReportStateData []MultipleWorksiteReportStateDataType `xml:"urn:com.workday/bsvc Multiple_Worksite_Report_State_Data,omitempty"`
}

// Contains request criteria for the multiple worksite report.
type MultipleWorksiteReportRequestCriteriaType struct {
	CompanyReference             *CompanyObjectType                `xml:"urn:com.workday/bsvc Company_Reference"`
	CalendarQuarterReference     *CalendarQuarterObjectType        `xml:"urn:com.workday/bsvc Calendar_Quarter_Reference"`
	PayrollTaxAuthorityReference []PayrollStateAuthorityObjectType `xml:"urn:com.workday/bsvc Payroll_Tax_Authority_Reference,omitempty"`
}

// Contains the data for the multiple worksite report response.
type MultipleWorksiteReportResponseDataType struct {
	MultipleWorksiteReportData *MultipleWorksiteReportDataType `xml:"urn:com.workday/bsvc Multiple_Worksite_Report_Data,omitempty"`
}

// Data for the state tax authority which has Locations (Business Sites) with workers who receive wages subject to unemployment insurance.
type MultipleWorksiteReportStateDataType struct {
	StateTaxAuthorityReference *PayrollTaxAuthorityObjectType `xml:"urn:com.workday/bsvc State_Tax_Authority_Reference,omitempty"`
	StateUIAccountNumber       string                         `xml:"urn:com.workday/bsvc State_UI_Account_Number,omitempty"`
	WorksiteData               []WorksiteDataType             `xml:"urn:com.workday/bsvc Worksite_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type NICategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type NICategoryObjectType struct {
	ID         []NICategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element for all of the name data such as first and last name
type NameDataType struct {
	CountryReference   *CountryReferenceType `xml:"urn:com.workday/bsvc Country_Reference"`
	AdditionalNameType string                `xml:"urn:com.workday/bsvc Additional_Name_Type,omitempty"`
	Prefix             []PrefixNameDataType  `xml:"urn:com.workday/bsvc Prefix,omitempty"`
	FirstName          string                `xml:"urn:com.workday/bsvc First_Name,omitempty"`
	MiddleName         string                `xml:"urn:com.workday/bsvc Middle_Name,omitempty"`
	LastName           []LastNameDataType    `xml:"urn:com.workday/bsvc Last_Name,omitempty"`
	LocalNameData      *LocalNameDataType    `xml:"urn:com.workday/bsvc Local_Name_Data,omitempty"`
	Suffix             []SuffixNameDataType  `xml:"urn:com.workday/bsvc Suffix,omitempty"`
	IsLegal            bool                  `xml:"urn:com.workday/bsvc Is_Legal,attr,omitempty"`
	IsPreferred        bool                  `xml:"urn:com.workday/bsvc Is_Preferred,attr,omitempty"`
	EffectiveDate      time.Time             `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	LastModified       time.Time             `xml:"urn:com.workday/bsvc Last_Modified,attr,omitempty"`
}

func (t *NameDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T NameDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate     `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
		LastModified  *xsdDateTime `xml:"urn:com.workday/bsvc Last_Modified,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	layout.LastModified = (*xsdDateTime)(&layout.T.LastModified)
	return e.EncodeElement(layout, start)
}
func (t *NameDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T NameDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate     `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
		LastModified  *xsdDateTime `xml:"urn:com.workday/bsvc Last_Modified,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	overlay.LastModified = (*xsdDateTime)(&overlay.T.LastModified)
	return d.DecodeElement(&overlay, &start)
}

// Name Data for Get Payroll Results
type NameDataforGetPayrollResultsType struct {
	NameData []NameDataType `xml:"urn:com.workday/bsvc Name_Data,omitempty"`
}

// Wrapper for National Identifier Data.
type NationalIDDataType struct {
	ID                  string                    `xml:"urn:com.workday/bsvc ID,omitempty"`
	IDTypeReference     *NationalIDTypeObjectType `xml:"urn:com.workday/bsvc ID_Type_Reference,omitempty"`
	CountryReference    *CountryObjectType        `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	IssuedDate          *time.Time                `xml:"urn:com.workday/bsvc Issued_Date,omitempty"`
	ExpirationDate      *time.Time                `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	VerificationDate    *time.Time                `xml:"urn:com.workday/bsvc Verification_Date,omitempty"`
	Series              string                    `xml:"urn:com.workday/bsvc Series,omitempty"`
	IssuingAgency       string                    `xml:"urn:com.workday/bsvc Issuing_Agency,omitempty"`
	VerifiedByReference *WorkerObjectType         `xml:"urn:com.workday/bsvc Verified_By_Reference,omitempty"`
}

func (t *NationalIDDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T NationalIDDataType
	var layout struct {
		*T
		IssuedDate       *xsdDate `xml:"urn:com.workday/bsvc Issued_Date,omitempty"`
		ExpirationDate   *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
		VerificationDate *xsdDate `xml:"urn:com.workday/bsvc Verification_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.IssuedDate = (*xsdDate)(layout.T.IssuedDate)
	layout.ExpirationDate = (*xsdDate)(layout.T.ExpirationDate)
	layout.VerificationDate = (*xsdDate)(layout.T.VerificationDate)
	return e.EncodeElement(layout, start)
}
func (t *NationalIDDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T NationalIDDataType
	var overlay struct {
		*T
		IssuedDate       *xsdDate `xml:"urn:com.workday/bsvc Issued_Date,omitempty"`
		ExpirationDate   *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
		VerificationDate *xsdDate `xml:"urn:com.workday/bsvc Verification_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.IssuedDate = (*xsdDate)(overlay.T.IssuedDate)
	overlay.ExpirationDate = (*xsdDate)(overlay.T.ExpirationDate)
	overlay.VerificationDate = (*xsdDate)(overlay.T.VerificationDate)
	return d.DecodeElement(&overlay, &start)
}

// National ID Data for Get Payroll Results
type NationalIDDataforGetPayrollResultsType struct {
	NationalIDData []NationalIDDataType `xml:"urn:com.workday/bsvc National_ID_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type NationalIDTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type NationalIDTypeObjectType struct {
	ID         []NationalIDTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This is wrapper element for Annual Non Qualified plans Data.
type NonQualifiedPensionDataType struct {
	Section457DistYTD       float64 `xml:"urn:com.workday/bsvc Section_457_Dist_YTD,omitempty"`
	Nonsection457DistYTD    float64 `xml:"urn:com.workday/bsvc Non-section_457_Dist_YTD,omitempty"`
	Section457ContribYTD    float64 `xml:"urn:com.workday/bsvc Section_457_Contrib_YTD,omitempty"`
	Nonsection457ContribYTD float64 `xml:"urn:com.workday/bsvc Non-section_457_Contrib_YTD,omitempty"`
}

// OASDI Exempt Data
type OASDIExemptDataType struct {
	ExemptfromOASDI                  *bool                                    `xml:"urn:com.workday/bsvc Exempt_from_OASDI,omitempty"`
	OASDIReasonforExemptionReference []PayrollPayeeFICAExemptReasonObjectType `xml:"urn:com.workday/bsvc OASDI_Reason_for_Exemption_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ObjectClassObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ObjectClassObjectType struct {
	ID         []ObjectClassObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The Data for Off-Cycle Payments, which contains the Result Type, the Reason, and whether or not it is an Additional Payment
type OffCycleDataType struct {
	ResultTypeReference *PayrollOffcycleTypeObjectType `xml:"urn:com.workday/bsvc Result_Type_Reference,omitempty"`
	AdditionalPay       *bool                          `xml:"urn:com.workday/bsvc Additional_Pay,omitempty"`
	ReasonReference     *PayrollActionReasonObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
}

// Off-cycle Input Data
type OffcycleInputDataType struct {
	EarningReference    *EarningAllObjectType         `xml:"urn:com.workday/bsvc Earning_Reference"`
	DeductionReference  *DeductionAllObjectType       `xml:"urn:com.workday/bsvc Deduction_Reference"`
	PositionReference   *PositionElementObjectType    `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	Amount              float64                       `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Hours               float64                       `xml:"urn:com.workday/bsvc Hours,omitempty"`
	Rate                float64                       `xml:"urn:com.workday/bsvc Rate,omitempty"`
	Adjustment          *bool                         `xml:"urn:com.workday/bsvc Adjustment,omitempty"`
	ReferenceDate       *time.Time                    `xml:"urn:com.workday/bsvc Reference_Date,omitempty"`
	CurrencyReference   *CurrencyObjectType           `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	CoverageStartDate   *time.Time                    `xml:"urn:com.workday/bsvc Coverage_Start_Date,omitempty"`
	CoverageEndDate     *time.Time                    `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	PayrollWorktagsData *PayrollInputWorktagsDataType `xml:"urn:com.workday/bsvc Payroll_Worktags_Data,omitempty"`
	InputLineData       []OffcycleInputLineDataType   `xml:"urn:com.workday/bsvc Input_Line_Data,omitempty"`
	CompanyReference    *CompanyObjectType            `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	Comment             string                        `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

func (t *OffcycleInputDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T OffcycleInputDataType
	var layout struct {
		*T
		ReferenceDate     *xsdDate `xml:"urn:com.workday/bsvc Reference_Date,omitempty"`
		CoverageStartDate *xsdDate `xml:"urn:com.workday/bsvc Coverage_Start_Date,omitempty"`
		CoverageEndDate   *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ReferenceDate = (*xsdDate)(layout.T.ReferenceDate)
	layout.CoverageStartDate = (*xsdDate)(layout.T.CoverageStartDate)
	layout.CoverageEndDate = (*xsdDate)(layout.T.CoverageEndDate)
	return e.EncodeElement(layout, start)
}
func (t *OffcycleInputDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T OffcycleInputDataType
	var overlay struct {
		*T
		ReferenceDate     *xsdDate `xml:"urn:com.workday/bsvc Reference_Date,omitempty"`
		CoverageStartDate *xsdDate `xml:"urn:com.workday/bsvc Coverage_Start_Date,omitempty"`
		CoverageEndDate   *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ReferenceDate = (*xsdDate)(overlay.T.ReferenceDate)
	overlay.CoverageStartDate = (*xsdDate)(overlay.T.CoverageStartDate)
	overlay.CoverageEndDate = (*xsdDate)(overlay.T.CoverageEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Off-cycle Input Line Data
type OffcycleInputLineDataType struct {
	RelatedCalculationReference *RelatedCalculationAllObjectType `xml:"urn:com.workday/bsvc Related_Calculation_Reference,omitempty"`
	Value                       float64                          `xml:"urn:com.workday/bsvc Value,omitempty"`
}

// On Demand Payment Data
type OnDemandPaymentDataType struct {
	UseSupplementalTaxRate           *bool                          `xml:"urn:com.workday/bsvc Use_Supplemental_Tax_Rate,omitempty"`
	OverridePaymentMethod            *bool                          `xml:"urn:com.workday/bsvc Override_Payment_Method,omitempty"`
	PaymentTypeIDReference           *PaymentTypeObjectType         `xml:"urn:com.workday/bsvc Payment_Type_ID_Reference,omitempty"`
	PaytoBalanceAccount              *bool                          `xml:"urn:com.workday/bsvc Pay_to_Balance_Account,omitempty"`
	TakeAdditionalWithholding        *bool                          `xml:"urn:com.workday/bsvc Take_Additional_Withholding,omitempty"`
	IncludeRetroDifferencesinPayment *bool                          `xml:"urn:com.workday/bsvc Include_Retro_Differences_in_Payment,omitempty"`
	LoadorRefreshInput               *bool                          `xml:"urn:com.workday/bsvc Load_or_Refresh_Input,omitempty"`
	TaxFrequencyOverride             []TaxFrequencyOverrideDataType `xml:"urn:com.workday/bsvc Tax_Frequency_Override,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type OrganizationTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type OrganizationTypeObjectType struct {
	ID         []OrganizationTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Originating Party Bank Data. Can be conditionally returned by setting the "Include Originating Bank data" grouping flag.
type OriginatingPartyBankDataType struct {
	BankAccountReference          []BankAccountObjectType         `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	AccountNickname               string                          `xml:"urn:com.workday/bsvc Account_Nickname,omitempty"`
	AccountNumber                 string                          `xml:"urn:com.workday/bsvc Account_Number,omitempty"`
	AccountTypeCode               string                          `xml:"urn:com.workday/bsvc Account_Type_Code,omitempty"`
	FinancialInstitutionReference *FinancialInstitutionObjectType `xml:"urn:com.workday/bsvc Financial_Institution_Reference,omitempty"`
	IBAN                          string                          `xml:"urn:com.workday/bsvc IBAN,omitempty"`
	RoutingTransitNumber          string                          `xml:"urn:com.workday/bsvc Routing_Transit_Number,omitempty"`
	BankIdentificationCode        string                          `xml:"urn:com.workday/bsvc Bank_Identification_Code,omitempty"`
	BranchName                    string                          `xml:"urn:com.workday/bsvc Branch_Name,omitempty"`
	BranchIDNumber                string                          `xml:"urn:com.workday/bsvc Branch_ID_Number,omitempty"`
	CountryReference              *CountryObjectType              `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CurrencyReference             *CurrencyObjectType             `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	BankAccountName               string                          `xml:"urn:com.workday/bsvc Bank_Account_Name,omitempty"`
	CheckDigit                    string                          `xml:"urn:com.workday/bsvc Check_Digit,omitempty"`
	RollNumber                    string                          `xml:"urn:com.workday/bsvc Roll_Number,omitempty"`
	BankAddressData               []AddressInformationDataType    `xml:"urn:com.workday/bsvc Bank_Address_Data,omitempty"`
}

// Originating Party Data
type OriginatingPartyWWSDataType struct {
	OrganizationName               string                      `xml:"urn:com.workday/bsvc Organization_Name,omitempty"`
	TransactionTaxID               string                      `xml:"urn:com.workday/bsvc Transaction_Tax_ID,omitempty"`
	EmailAddressData               []EmailAddressDataType      `xml:"urn:com.workday/bsvc Email_Address_Data,omitempty"`
	AddressInformationData         *AddressInformationDataType `xml:"urn:com.workday/bsvc Address_Information_Data,omitempty"`
	OrganizationPrimaryPhoneNumber string                      `xml:"urn:com.workday/bsvc Organization_Primary_Phone_Number,omitempty"`
}

// This element is wrapper for W-2 Box 14 Other Data.
type OtherDataType struct {
	Label         string  `xml:"urn:com.workday/bsvc Label,omitempty"`
	Amount        float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	CustomerOwned *bool   `xml:"urn:com.workday/bsvc Customer_Owned,omitempty"`
	PlanNumber    string  `xml:"urn:com.workday/bsvc Plan_Number,omitempty"`
}

// This element is wrapper for W-2 Box 14 Other Data.
type OtherDataWrapperType struct {
	OtherData []OtherDataType `xml:"urn:com.workday/bsvc Other_Data,omitempty"`
}

// Other Monies Information
type OtherMoniesDataType struct {
	ROEOtherMoniesCode      string     `xml:"urn:com.workday/bsvc ROE_Other_Monies_Code,omitempty"`
	ROEOtherMoniesStartDate *time.Time `xml:"urn:com.workday/bsvc ROE_Other_Monies_Start_Date,omitempty"`
	ROEOtherMoniesEndDate   *time.Time `xml:"urn:com.workday/bsvc ROE_Other_Monies_End_Date,omitempty"`
	ROEOtherMoniesAmount    float64    `xml:"urn:com.workday/bsvc ROE_Other_Monies_Amount,omitempty"`
}

func (t *OtherMoniesDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T OtherMoniesDataType
	var layout struct {
		*T
		ROEOtherMoniesStartDate *xsdDate `xml:"urn:com.workday/bsvc ROE_Other_Monies_Start_Date,omitempty"`
		ROEOtherMoniesEndDate   *xsdDate `xml:"urn:com.workday/bsvc ROE_Other_Monies_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ROEOtherMoniesStartDate = (*xsdDate)(layout.T.ROEOtherMoniesStartDate)
	layout.ROEOtherMoniesEndDate = (*xsdDate)(layout.T.ROEOtherMoniesEndDate)
	return e.EncodeElement(layout, start)
}
func (t *OtherMoniesDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T OtherMoniesDataType
	var overlay struct {
		*T
		ROEOtherMoniesStartDate *xsdDate `xml:"urn:com.workday/bsvc ROE_Other_Monies_Start_Date,omitempty"`
		ROEOtherMoniesEndDate   *xsdDate `xml:"urn:com.workday/bsvc ROE_Other_Monies_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ROEOtherMoniesStartDate = (*xsdDate)(overlay.T.ROEOtherMoniesStartDate)
	overlay.ROEOtherMoniesEndDate = (*xsdDate)(overlay.T.ROEOtherMoniesEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type OutsourcedPaymentGroupObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type OutsourcedPaymentGroupObjectType struct {
	ID         []OutsourcedPaymentGroupObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayAccumulationAllObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayAccumulationAllObjectType struct {
	ID         []PayAccumulationAllObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payroll Result element
type PayAccumulationDataforGetPayrollResultsType struct {
	PayAccumulationReference []PayrollCalculationObjectType `xml:"urn:com.workday/bsvc Pay_Accumulation_Reference,omitempty"`
	Amount                   float64                        `xml:"urn:com.workday/bsvc Amount,omitempty"`
	QTDAmount                float64                        `xml:"urn:com.workday/bsvc QTD_Amount,omitempty"`
	YTDAmount                float64                        `xml:"urn:com.workday/bsvc YTD_Amount,omitempty"`
}

// Pay Calculation Balance Data
type PayCalculationBalanceDataType struct {
	PayCalculationReference []PayCalculationReferenceType                `xml:"urn:com.workday/bsvc Pay_Calculation_Reference,omitempty"`
	BalancePeriodData       []BalancePeriodDataforGetPayrollBalancesType `xml:"urn:com.workday/bsvc Balance_Period_Data,omitempty"`
}

// Pay Calculation Reference
type PayCalculationReferenceType struct {
	EarningReference         *EarningAllObjectType         `xml:"urn:com.workday/bsvc Earning_Reference,omitempty"`
	DeductionReference       *DeductionAllObjectType       `xml:"urn:com.workday/bsvc Deduction_Reference,omitempty"`
	PayAccumulationReference *PayAccumulationAllObjectType `xml:"urn:com.workday/bsvc Pay_Accumulation_Reference,omitempty"`
}

// Pay Calculation Selection for Get Payroll Balances Request
type PayCalculationSelectionType struct {
	EarningReference         []EarningAllObjectType         `xml:"urn:com.workday/bsvc Earning_Reference,omitempty"`
	DeductionReference       []DeductionAllObjectType       `xml:"urn:com.workday/bsvc Deduction_Reference,omitempty"`
	PayAccumulationReference []PayAccumulationAllObjectType `xml:"urn:com.workday/bsvc Pay_Accumulation_Reference,omitempty"`
}

// Pay Component Selection
type PayCalculationsSelectedType struct {
	EarningReference         []EarningAllObjectType         `xml:"urn:com.workday/bsvc Earning_Reference,omitempty"`
	DeductionReference       []DeductionAllObjectType       `xml:"urn:com.workday/bsvc Deduction_Reference,omitempty"`
	PayAccumulationReference []PayAccumulationAllObjectType `xml:"urn:com.workday/bsvc Pay_Accumulation_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayComponentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayComponentObjectType struct {
	ID         []PayComponentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Pay Component Reference
type PayComponentReferenceType struct {
	Earning   *EarningReferenceType   `xml:"urn:com.workday/bsvc Earning"`
	Deduction *DeductionReferenceType `xml:"urn:com.workday/bsvc Deduction"`
}

// Contains a unique identifier for an instance of an object.
type PayGroupDetailObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayGroupDetailObjectType struct {
	ID         []PayGroupDetailObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayGroupObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayGroupObjectType struct {
	ID         []PayGroupObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayGroupPayRunGroupSelectionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayGroupPayRunGroupSelectionObjectType struct {
	ID         []PayGroupPayRunGroupSelectionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains counts of Pay Results in different states along with the Pay Cycle Type and Action Reason
type PayGroupPeriodPayCalculationStatusAsOfNowType struct {
	PayCycleTypeReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Pay_Cycle_Type_Reference,omitempty"`
	ActionReason          string                      `xml:"urn:com.workday/bsvc Action_Reason,omitempty"`
	InProgress            float64                     `xml:"urn:com.workday/bsvc In_Progress,omitempty"`
	RequiresRecalculation float64                     `xml:"urn:com.workday/bsvc Requires_Recalculation,omitempty"`
	PendingCompletion     float64                     `xml:"urn:com.workday/bsvc Pending_Completion,omitempty"`
	InError               float64                     `xml:"urn:com.workday/bsvc In_Error,omitempty"`
	OnHold                float64                     `xml:"urn:com.workday/bsvc On_Hold,omitempty"`
	Completed             float64                     `xml:"urn:com.workday/bsvc Completed,omitempty"`
	Total                 float64                     `xml:"urn:com.workday/bsvc Total,omitempty"`
}

// Pay Groups being Calculated
type PayGroupsBeingCalculatedType struct {
	PayrollRunDetailPayCalcStatus []PayrollRunDetailPayCalcStatusType `xml:"urn:com.workday/bsvc Payroll_Run_Detail_Pay_Calc_Status,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayRateTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayRateTypeObjectType struct {
	ID         []PayRateTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayRunGroupSelectionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayRunGroupSelectionObjectType struct {
	ID         []PayRunGroupSelectionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// An individual workers Paycheck Delivery Method
type PaycheckDeliveryDataType struct {
	PaycheckDeliveryID               string                             `xml:"urn:com.workday/bsvc Paycheck_Delivery_ID,omitempty"`
	WorkerReference                  *WorkerObjectType                  `xml:"urn:com.workday/bsvc Worker_Reference"`
	DateLastUpdated                  *time.Time                         `xml:"urn:com.workday/bsvc Date_Last_Updated,omitempty"`
	PaycheckDeliveryMethodReference  *PaycheckDeliveryMethodObjectType  `xml:"urn:com.workday/bsvc Paycheck_Delivery_Method_Reference,omitempty"`
	PayslipPrintingOverrideReference *PayslipPrintingOverrideObjectType `xml:"urn:com.workday/bsvc Payslip_Printing_Override_Reference,omitempty"`
	CompanyReference                 *CompanyObjectType                 `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
}

func (t *PaycheckDeliveryDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PaycheckDeliveryDataType
	var layout struct {
		*T
		DateLastUpdated *xsdDate `xml:"urn:com.workday/bsvc Date_Last_Updated,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateLastUpdated = (*xsdDate)(layout.T.DateLastUpdated)
	return e.EncodeElement(layout, start)
}
func (t *PaycheckDeliveryDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PaycheckDeliveryDataType
	var overlay struct {
		*T
		DateLastUpdated *xsdDate `xml:"urn:com.workday/bsvc Date_Last_Updated,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateLastUpdated = (*xsdDate)(overlay.T.DateLastUpdated)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PaycheckDeliveryMethodObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaycheckDeliveryMethodObjectType struct {
	ID         []PaycheckDeliveryMethodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PaycheckDeliveryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaycheckDeliveryObjectType struct {
	ID         []PaycheckDeliveryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Used to return a selected set of Workers Pay Check Delivery records
type PaycheckDeliveryRequestCriteriaType struct {
}

// Set of referenceis to workers paycheck delivery methods.
type PaycheckDeliveryRequestReferencesType struct {
	PaycheckDeliveryReference []PaycheckDeliveryObjectType `xml:"urn:com.workday/bsvc Paycheck_Delivery_Reference"`
}

// Complete set of all Response Data returned.
type PaycheckDeliveryResponseDataType struct {
	PaycheckDelivery []PaycheckDeliveryType `xml:"urn:com.workday/bsvc Paycheck_Delivery,omitempty"`
}

// Response Group Data
type PaycheckDeliveryResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Delivery Response Data
type PaycheckDeliveryType struct {
	PaycheckDeliveryReference *PaycheckDeliveryObjectType `xml:"urn:com.workday/bsvc Paycheck_Delivery_Reference,omitempty"`
	PaycheckDeliveryData      []PaycheckDeliveryDataType  `xml:"urn:com.workday/bsvc Paycheck_Delivery_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayeeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayeeObjectType struct {
	ID         []PayeeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PaymentCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentCategoryObjectType struct {
	ID         []PaymentCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// One or more references to existing history payments
type PaymentHistoryRequestReferencesType struct {
	PaymentHistoryRequestReference []PayrollOffcycleResultOverrideObjectType `xml:"urn:com.workday/bsvc Payment_History_Request_Reference"`
}

// Contains a unique identifier for an instance of an object.
type PaymentMethodObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentMethodObjectType struct {
	ID         []PaymentMethodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PaymentTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentTypeObjectType struct {
	ID         []PaymentTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollActionReasonObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollActionReasonObjectType struct {
	ID         []PayrollActionReasonObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Get Payroll Balances Request
type PayrollBalanceRequestCriteriaType struct {
	EmployeeReference                  []EmployeeObjectType         `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	CompanyReference                   []CompanyObjectType          `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	PayCalculationCriteria             *PayCalculationSelectionType `xml:"urn:com.workday/bsvc Pay_Calculation_Criteria"`
	BalancePeriodDefinitionReference   []BalancePeriodAllObjectType `xml:"urn:com.workday/bsvc Balance_Period_Definition_Reference"`
	BalancePeriodAsofDate              time.Time                    `xml:"urn:com.workday/bsvc Balance_Period_As_of_Date"`
	BalancePeriodCutoffDate            *time.Time                   `xml:"urn:com.workday/bsvc Balance_Period_Cutoff_Date,omitempty"`
	PeriodforBalanceReference          *PeriodObjectType            `xml:"urn:com.workday/bsvc Period_for_Balance_Reference,omitempty"`
	IncludeResultsProcessedbyReference *time.Time                   `xml:"urn:com.workday/bsvc Include_Results_Processed_by_Reference,omitempty"`
}

func (t *PayrollBalanceRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollBalanceRequestCriteriaType
	var layout struct {
		*T
		BalancePeriodAsofDate              *xsdDate     `xml:"urn:com.workday/bsvc Balance_Period_As_of_Date"`
		BalancePeriodCutoffDate            *xsdDate     `xml:"urn:com.workday/bsvc Balance_Period_Cutoff_Date,omitempty"`
		IncludeResultsProcessedbyReference *xsdDateTime `xml:"urn:com.workday/bsvc Include_Results_Processed_by_Reference,omitempty"`
	}
	layout.T = (*T)(t)
	layout.BalancePeriodAsofDate = (*xsdDate)(&layout.T.BalancePeriodAsofDate)
	layout.BalancePeriodCutoffDate = (*xsdDate)(layout.T.BalancePeriodCutoffDate)
	layout.IncludeResultsProcessedbyReference = (*xsdDateTime)(layout.T.IncludeResultsProcessedbyReference)
	return e.EncodeElement(layout, start)
}
func (t *PayrollBalanceRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollBalanceRequestCriteriaType
	var overlay struct {
		*T
		BalancePeriodAsofDate              *xsdDate     `xml:"urn:com.workday/bsvc Balance_Period_As_of_Date"`
		BalancePeriodCutoffDate            *xsdDate     `xml:"urn:com.workday/bsvc Balance_Period_Cutoff_Date,omitempty"`
		IncludeResultsProcessedbyReference *xsdDateTime `xml:"urn:com.workday/bsvc Include_Results_Processed_by_Reference,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.BalancePeriodAsofDate = (*xsdDate)(&overlay.T.BalancePeriodAsofDate)
	overlay.BalancePeriodCutoffDate = (*xsdDate)(overlay.T.BalancePeriodCutoffDate)
	overlay.IncludeResultsProcessedbyReference = (*xsdDateTime)(overlay.T.IncludeResultsProcessedbyReference)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Balance Response Data
type PayrollBalanceResponseDataType struct {
	PayrollBalance []PayrollBalanceType `xml:"urn:com.workday/bsvc Payroll_Balance,omitempty"`
}

// Payroll Balance Response Group
type PayrollBalanceResponseGroupType struct {
	IncludeNameData       *bool `xml:"urn:com.workday/bsvc Include_Name_Data,omitempty"`
	IncludeNationalIDData *bool `xml:"urn:com.workday/bsvc Include_National_ID_Data,omitempty"`
}

// Payroll Balance
type PayrollBalanceType struct {
	WorkerReference                 *WorkerObjectType               `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	WorkerExternalIntegrationIDData *WorkerReferenceType            `xml:"urn:com.workday/bsvc Worker_External_Integration_ID_Data,omitempty"`
	NameData                        []NameDataType                  `xml:"urn:com.workday/bsvc Name_Data,omitempty"`
	NationalIDData                  []NationalIDDataType            `xml:"urn:com.workday/bsvc National_ID_Data,omitempty"`
	PayCalculationBalanceData       []PayCalculationBalanceDataType `xml:"urn:com.workday/bsvc Pay_Calculation_Balance_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollCalculationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollCalculationObjectType struct {
	ID         []PayrollCalculationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollConstantNumberObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollConstantNumberObjectType struct {
	ID         []PayrollConstantNumberObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollConstantPercentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollConstantPercentObjectType struct {
	ID         []PayrollConstantPercentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollConstantTextObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollConstantTextObjectType struct {
	ID         []PayrollConstantTextObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element Containing Payroll Deduction Recipient Response Data.
// When a Deduction Recipient already exists and  is being updated through the Web Service will we be interpreting the absence of business entity and settlement account elements as persisting the existing data. In the event that Business Entity or settlement account Data is being updated through the web service, we will replace the existing data with the new data being added in the Web Service.
type PayrollDeductionRecipientDataType struct {
	PayrollDeductionRecipientName          string                         `xml:"urn:com.workday/bsvc Payroll_Deduction_Recipient_Name"`
	PayrollAlternateDeductionRecipientName string                         `xml:"urn:com.workday/bsvc Payroll_Alternate_Deduction_Recipient_Name,omitempty"`
	PayrollDeductionRecipientID            string                         `xml:"urn:com.workday/bsvc Payroll_Deduction_Recipient_ID,omitempty"`
	PaymentTypeReference                   *PaymentTypeObjectType         `xml:"urn:com.workday/bsvc Payment_Type_Reference"`
	BusinessEntityData                     *BusinessEntityWWSDataType     `xml:"urn:com.workday/bsvc Business_Entity_Data,omitempty"`
	SettlementAccountData                  []SettlementAccountWWSDataType `xml:"urn:com.workday/bsvc Settlement_Account_Data,omitempty"`
}

// Payroll Deduction Recipient Reference element contains the specific instance set containing the requested Payroll Deduction Recipient.
type PayrollDeductionRecipientRequestReferencesType struct {
	PayrollDeductionRecipientReference []DeductionRecipientObjectType `xml:"urn:com.workday/bsvc Payroll_Deduction_Recipient_Reference"`
}

// Element Containing Payroll Deduction Recipient Response Data.
// When a Deduction Recipient already exists and  is being updated through the Web Service will we be interpreting the absence of business entity and settlement account elements as persisting the existing data. In the event that Business Entity or settlement account Data is being updated through the web service, we will replace the existing data with the new data being added in the Web Service.
type PayrollDeductionRecipientResponseDataType struct {
	PayrollDeductionRecipient []PayrollDeductionRecipientType `xml:"urn:com.workday/bsvc Payroll_Deduction_Recipient,omitempty"`
}

// Element Containing Payroll Deduction Recipient Response
type PayrollDeductionRecipientResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Deduction Recipient Name
type PayrollDeductionRecipientType struct {
	PayrollDeductionRecipientReference *DeductionRecipientObjectType       `xml:"urn:com.workday/bsvc Payroll_Deduction_Recipient_Reference,omitempty"`
	PayrollDeductionRecipientData      []PayrollDeductionRecipientDataType `xml:"urn:com.workday/bsvc Payroll_Deduction_Recipient_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollEventCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollEventCategoryObjectType struct {
	ID         []PayrollEventCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element to hold all of the Data to Put or Get for a W-4
type PayrollFederalW4TaxElectionDataType struct {
	WorkerReference               *WorkerObjectType                   `xml:"urn:com.workday/bsvc Worker_Reference"`
	Effectiveasof                 time.Time                           `xml:"urn:com.workday/bsvc Effective_as_of"`
	CompanyReference              *CompanyObjectType                  `xml:"urn:com.workday/bsvc Company_Reference"`
	MaritalStatusReference        *PayrollWithholdingStatusObjectType `xml:"urn:com.workday/bsvc Marital_Status_Reference,omitempty"`
	NumberofAllowances            float64                             `xml:"urn:com.workday/bsvc Number_of_Allowances,omitempty"`
	AdditionalAmount              float64                             `xml:"urn:com.workday/bsvc Additional_Amount,omitempty"`
	Exempt                        *bool                               `xml:"urn:com.workday/bsvc Exempt,omitempty"`
	NonresidentAlien              *bool                               `xml:"urn:com.workday/bsvc Nonresident_Alien,omitempty"`
	ExemptfromNRAAdditionalAmount *bool                               `xml:"urn:com.workday/bsvc Exempt_from_NRA_Additional_Amount,omitempty"`
	LockInLetter                  *bool                               `xml:"urn:com.workday/bsvc Lock_In_Letter,omitempty"`
	NoWageNoTax                   *bool                               `xml:"urn:com.workday/bsvc No_Wage_No_Tax,omitempty"`
}

func (t *PayrollFederalW4TaxElectionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollFederalW4TaxElectionDataType
	var layout struct {
		*T
		Effectiveasof *xsdDate `xml:"urn:com.workday/bsvc Effective_as_of"`
	}
	layout.T = (*T)(t)
	layout.Effectiveasof = (*xsdDate)(&layout.T.Effectiveasof)
	return e.EncodeElement(layout, start)
}
func (t *PayrollFederalW4TaxElectionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollFederalW4TaxElectionDataType
	var overlay struct {
		*T
		Effectiveasof *xsdDate `xml:"urn:com.workday/bsvc Effective_as_of"`
	}
	overlay.T = (*T)(t)
	overlay.Effectiveasof = (*xsdDate)(&overlay.T.Effectiveasof)
	return d.DecodeElement(&overlay, &start)
}

// Element to hold all criteria used to select which W-4s to return
type PayrollFederalW4TaxElectionRequestCriteriaType struct {
	WorkerReference  []WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	CompanyReference *CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	EffectiveAsOf    *time.Time         `xml:"urn:com.workday/bsvc Effective_As_Of,omitempty"`
}

func (t *PayrollFederalW4TaxElectionRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollFederalW4TaxElectionRequestCriteriaType
	var layout struct {
		*T
		EffectiveAsOf *xsdDate `xml:"urn:com.workday/bsvc Effective_As_Of,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveAsOf = (*xsdDate)(layout.T.EffectiveAsOf)
	return e.EncodeElement(layout, start)
}
func (t *PayrollFederalW4TaxElectionRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollFederalW4TaxElectionRequestCriteriaType
	var overlay struct {
		*T
		EffectiveAsOf *xsdDate `xml:"urn:com.workday/bsvc Effective_As_Of,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveAsOf = (*xsdDate)(overlay.T.EffectiveAsOf)
	return d.DecodeElement(&overlay, &start)
}

// Element to hold all Payroll Payee Request References
type PayrollFederalW4TaxElectionRequestReferencesType struct {
	PayrollPayeeW4Reference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_W-4_Reference"`
}

// Element to hold all data to Respond with about the Payroll Payee W-4s
type PayrollFederalW4TaxElectionResponseDataType struct {
	PayrollFederalW4TaxElection []PayrollFederalW4TaxElectionType `xml:"urn:com.workday/bsvc Payroll_Federal_W-4_Tax_Election,omitempty"`
}

// Element to hold all information about the Payroll Payee W-4 Response Group
type PayrollFederalW4TaxElectionResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element to hold all information about a Payroll Payee W-4
type PayrollFederalW4TaxElectionType struct {
	PayrollPayeeW4Reference         *UniqueIdentifierObjectType          `xml:"urn:com.workday/bsvc Payroll_Payee_W-4_Reference,omitempty"`
	PayrollFederalW4TaxElectionData *PayrollFederalW4TaxElectionDataType `xml:"urn:com.workday/bsvc Payroll_Federal_W-4_Tax_Election_Data,omitempty"`
}

// The Payroll History Payment Data element contains all of the information necessary to define a history payment, including the functional keys (Worker, Period, Pay Group), the technical key (Payment ID), the payment information (Payment Date, Gross and Net Amounts, Sub-Period dates), Payroll Worktag overrides, and the line-item earnings and deductions detail.
type PayrollHistoryPaymentDataType struct {
	BatchID                 string                          `xml:"urn:com.workday/bsvc Batch_ID,omitempty"`
	PaymentID               string                          `xml:"urn:com.workday/bsvc Payment_ID"`
	EmployeeReference       *EmployeeObjectType             `xml:"urn:com.workday/bsvc Employee_Reference"`
	PeriodDate              time.Time                       `xml:"urn:com.workday/bsvc Period_Date"`
	RunCategoryReference    *RunCategoryObjectType          `xml:"urn:com.workday/bsvc Run_Category_Reference,omitempty"`
	PayGroupReference       *PayGroupObjectType             `xml:"urn:com.workday/bsvc Pay_Group_Reference,omitempty"`
	PayrollIDReference      *PayrollIDObjectType            `xml:"urn:com.workday/bsvc Payroll_ID_Reference,omitempty"`
	HistoryPaymentData      *HistoryPaymentDataType         `xml:"urn:com.workday/bsvc History_Payment_Data"`
	ResultWorktagOverrides  *ResultWorktagOverridesDataType `xml:"urn:com.workday/bsvc Result_Worktag_Overrides,omitempty"`
	AddtoExistingInputData  *bool                           `xml:"urn:com.workday/bsvc Add_to_Existing_Input_Data,omitempty"`
	HistoryPaymentInputData []HistoryPaymentInputDataType   `xml:"urn:com.workday/bsvc History_Payment_Input_Data"`
}

func (t *PayrollHistoryPaymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollHistoryPaymentDataType
	var layout struct {
		*T
		PeriodDate *xsdDate `xml:"urn:com.workday/bsvc Period_Date"`
	}
	layout.T = (*T)(t)
	layout.PeriodDate = (*xsdDate)(&layout.T.PeriodDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollHistoryPaymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollHistoryPaymentDataType
	var overlay struct {
		*T
		PeriodDate *xsdDate `xml:"urn:com.workday/bsvc Period_Date"`
	}
	overlay.T = (*T)(t)
	overlay.PeriodDate = (*xsdDate)(&overlay.T.PeriodDate)
	return d.DecodeElement(&overlay, &start)
}

// The Payroll History Payment Data element contains all of the information necessary to define a history payment, including the functional keys (Worker, Period, Pay Group), the technical key (Payment ID), the payment information (Payment Date, Gross and Net Amounts, Sub-Period dates), Payroll Worktag overrides, and the line-item earnings and deductions detail.
type PayrollHistoryPaymentResponseDataType struct {
	PayrollHistoryPayment []PayrollHistoryPaymentType `xml:"urn:com.workday/bsvc Payroll_History_Payment,omitempty"`
}

// Response for Get Payroll History Payment Web Service method
type PayrollHistoryPaymentType struct {
	PayrollHistoryPaymentReference *PayrollOffcycleResultOverrideObjectType `xml:"urn:com.workday/bsvc Payroll_History_Payment_Reference,omitempty"`
	PayrollHistoryPaymentData      *PayrollHistoryPaymentDataType           `xml:"urn:com.workday/bsvc Payroll_History_Payment_Data,omitempty"`
}

// Payroll History Result Request Criteria element.
type PayrollHistoryResultRequestCriteriaType struct {
	BatchIDReference []ApplicationBatchObjectType `xml:"urn:com.workday/bsvc Batch_ID_Reference"`
}

// Payroll ID Balance Data
type PayrollIDBalanceDataType struct {
	PayrollIDDetailData []PayrollIDDataDisplayType `xml:"urn:com.workday/bsvc Payroll_ID_Detail_Data,omitempty"`
	BalanceAmount       float64                    `xml:"urn:com.workday/bsvc Balance_Amount,omitempty"`
}

// Payroll ID Data
type PayrollIDDataDisplayType struct {
	ID                  string            `xml:"urn:com.workday/bsvc ID,omitempty"`
	WorkerReference     *WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	EffectiveDate       *time.Time        `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	AssignmentReason    string            `xml:"urn:com.workday/bsvc Assignment_Reason,omitempty"`
	AutomaticAssignment *bool             `xml:"urn:com.workday/bsvc Automatic_Assignment,omitempty"`
}

func (t *PayrollIDDataDisplayType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollIDDataDisplayType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollIDDataDisplayType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollIDDataDisplayType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PayrollIDObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollIDObjectType struct {
	ID         []PayrollIDObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollIRSCountryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollIRSCountryObjectType struct {
	ID         []PayrollIRSCountryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollIncomeCodeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollIncomeCodeObjectType struct {
	ID         []PayrollIncomeCodeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollIncomeCodeSubtypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollIncomeCodeSubtypeObjectType struct {
	ID         []PayrollIncomeCodeSubtypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollInputObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollInputObjectType struct {
	ID         []PayrollInputObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payroll Input Request Criteria
type PayrollInputRequestCriteriaType struct {
	WorkerReference        *WorkerObjectType        `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	EarningReference       []EarningAllObjectType   `xml:"urn:com.workday/bsvc Earning_Reference,omitempty"`
	DeductionReference     []DeductionAllObjectType `xml:"urn:com.workday/bsvc Deduction_Reference,omitempty"`
	StartDate              *time.Time               `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                *time.Time               `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	OpenEndedOngoingInputs *bool                    `xml:"urn:com.workday/bsvc Open-Ended_Ongoing_Inputs,omitempty"`
}

func (t *PayrollInputRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollInputRequestCriteriaType
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
func (t *PayrollInputRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollInputRequestCriteriaType
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

// Payroll Input Request References
type PayrollInputRequestReferencesType struct {
	PayrollInputReference []PayrollInputObjectType `xml:"urn:com.workday/bsvc Payroll_Input_Reference"`
}

// Payroll Input Response Group
type PayrollInputResponseGroupType struct {
	IncludeReference        *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludePayrollInputData *bool `xml:"urn:com.workday/bsvc Include_Payroll_Input_Data,omitempty"`
}

// Payroll Worktags
type PayrollInputWorktagsDataType struct {
	LocationReference                *LocationObjectType                            `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	RegionReference                  *RegionObjectType                              `xml:"urn:com.workday/bsvc Region_Reference,omitempty"`
	JobProfileReference              *JobProfileObjectType                          `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	CostCenterReference              *CostCenterObjectType                          `xml:"urn:com.workday/bsvc Cost_Center_Reference,omitempty"`
	ProjectReference                 *ProjectObjectType                             `xml:"urn:com.workday/bsvc Project_Reference,omitempty"`
	ProjectPhaseReference            *ProjectPlanPhaseObjectType                    `xml:"urn:com.workday/bsvc Project_Phase_Reference,omitempty"`
	ProjectTaskReference             *ProjectPlanTaskObjectType                     `xml:"urn:com.workday/bsvc Project_Task_Reference,omitempty"`
	WithholdingOrderCaseReference    *WithholdingOrderCaseObjectType                `xml:"urn:com.workday/bsvc Withholding_Order_Case_Reference,omitempty"`
	StateAuthorityReference          *PayrollStateAuthorityObjectType               `xml:"urn:com.workday/bsvc State_Authority_Reference,omitempty"`
	WorkersCompensationCodeReference *WorkersCompensationCodeObjectType             `xml:"urn:com.workday/bsvc Workers_Compensation_Code_Reference,omitempty"`
	CountyAuthorityReference         *PayrollLocalCountyAuthorityObjectType         `xml:"urn:com.workday/bsvc County_Authority_Reference,omitempty"`
	CityAuthorityReference           *PayrollLocalCityAuthorityObjectType           `xml:"urn:com.workday/bsvc City_Authority_Reference,omitempty"`
	SchoolDistrictAuthorityReference *PayrollLocalSchoolDistrictAuthorityObjectType `xml:"urn:com.workday/bsvc School_District_Authority_Reference,omitempty"`
	CustomWorktag01Reference         *CustomWorktag01ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_01_Reference,omitempty"`
	CustomWorktag02Reference         *CustomWorktag02ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_02_Reference,omitempty"`
	CustomWorktag03Reference         *CustomWorktag03ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_03_Reference,omitempty"`
	CustomWorktag04Reference         *CustomWorktag04ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_04_Reference,omitempty"`
	CustomWorktag05Reference         *CustomWorktag05ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_05_Reference,omitempty"`
	FundReference                    *FundObjectType                                `xml:"urn:com.workday/bsvc Fund_Reference,omitempty"`
	GrantReference                   *GrantObjectType                               `xml:"urn:com.workday/bsvc Grant_Reference,omitempty"`
	GiftReference                    *GiftObjectType                                `xml:"urn:com.workday/bsvc Gift_Reference,omitempty"`
	ProgramReference                 *ProgramObjectType                             `xml:"urn:com.workday/bsvc Program_Reference,omitempty"`
	BusinessUnitReference            *BusinessUnitObjectType                        `xml:"urn:com.workday/bsvc Business_Unit_Reference,omitempty"`
	ObjectClassReference             *ObjectClassObjectType                         `xml:"urn:com.workday/bsvc Object_Class_Reference,omitempty"`
	CustomOrganizationReference      []CustomOrganizationObjectType                 `xml:"urn:com.workday/bsvc Custom_Organization_Reference,omitempty"`
	CustomWorktag06Reference         *CustomWorktag06ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_06_Reference,omitempty"`
	CustomWorktag07Reference         *CustomWorktag07ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_07_Reference,omitempty"`
	CustomWorktag08Reference         *CustomWorktag08ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_08_Reference,omitempty"`
	CustomWorktag09Reference         *CustomWorktag09ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_09_Reference,omitempty"`
	CustomWorktag10Reference         *CustomWorktag10ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_10_Reference,omitempty"`
	CustomWorktag11Reference         *CustomWorktag11ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_11_Reference,omitempty"`
	CustomWorktag12Reference         *CustomWorktag12ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_12_Reference,omitempty"`
	CustomWorktag13Reference         *CustomWorktag13ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_13_Reference,omitempty"`
	CustomWorktag14Reference         *CustomWorktag14ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_14_Reference,omitempty"`
	CustomWorktag15Reference         *CustomWorktag15ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_15_Reference,omitempty"`
	LocalOtherTaxAuthorityReference  *PayrollOtherAuthorityObjectType               `xml:"urn:com.workday/bsvc Local_Other_Tax_Authority_Reference,omitempty"`
	NICategoryReference              *NICategoryObjectType                          `xml:"urn:com.workday/bsvc NI_Category_Reference,omitempty"`
	ARRCOAGIRCCategoryReference      *ARRCOAGIRCRubricValueObjectType               `xml:"urn:com.workday/bsvc ARRCO-AGIRC_Category_Reference,omitempty"`
}

// Payroll Involuntary Withholding Order
type PayrollInvoluntaryWithholdingOrderDataType struct {
	EmployeeReference                        *WorkerObjectType                              `xml:"urn:com.workday/bsvc Employee_Reference"`
	WithholdingOrderTypeReference            *WithholdingOrderTypeObjectType                `xml:"urn:com.workday/bsvc Withholding_Order_Type_Reference"`
	WithholdingOrderCaseReference            *WithholdingOrderCaseObjectType                `xml:"urn:com.workday/bsvc Withholding_Order_Case_Reference"`
	CaseNumber                               string                                         `xml:"urn:com.workday/bsvc Case_Number"`
	WithholdingOrderAdditionalOrderNumber    string                                         `xml:"urn:com.workday/bsvc Withholding_Order_Additional_Order_Number,omitempty"`
	OrderDate                                time.Time                                      `xml:"urn:com.workday/bsvc Order_Date"`
	ReceivedDate                             time.Time                                      `xml:"urn:com.workday/bsvc Received_Date"`
	BeginDate                                time.Time                                      `xml:"urn:com.workday/bsvc Begin_Date"`
	EndDate                                  *time.Time                                     `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	CompanyReference                         *CompanyObjectType                             `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	InactiveOrder                            *bool                                          `xml:"urn:com.workday/bsvc Inactive_Order,omitempty"`
	WithholdingOrderAmountTypeReference      *WithholdingOrderAmountTypeObjectType          `xml:"urn:com.workday/bsvc Withholding_Order_Amount_Type_Reference,omitempty"`
	WithholdingOrderAmount                   float64                                        `xml:"urn:com.workday/bsvc Withholding_Order_Amount,omitempty"`
	WithholdingOrderAmountasPercent          float64                                        `xml:"urn:com.workday/bsvc Withholding_Order_Amount_as_Percent,omitempty"`
	PayPeriodFrequencyReference              *FrequencyObjectType                           `xml:"urn:com.workday/bsvc Pay_Period_Frequency_Reference"`
	TotalDebtAmount                          float64                                        `xml:"urn:com.workday/bsvc Total_Debt_Amount,omitempty"`
	MonthlyLimit                             float64                                        `xml:"urn:com.workday/bsvc Monthly_Limit,omitempty"`
	IssuedinReference                        *PayrollTaxAuthorityObjectType                 `xml:"urn:com.workday/bsvc Issued_in_Reference"`
	DeductionRecipientReference              *DeductionRecipientObjectType                  `xml:"urn:com.workday/bsvc Deduction_Recipient_Reference"`
	OriginatingAuthority                     string                                         `xml:"urn:com.workday/bsvc Originating_Authority,omitempty"`
	Memo                                     string                                         `xml:"urn:com.workday/bsvc Memo,omitempty"`
	BankruptcyOrderData                      *BankruptcyOrderSpecificDataType               `xml:"urn:com.workday/bsvc Bankruptcy_Order_Data,omitempty"`
	CreditorGarnishmentData                  *CreditorGarnishmentSpecificDataType           `xml:"urn:com.workday/bsvc Creditor_Garnishment_Data,omitempty"`
	FederalTaxLevyData                       *FederalTaxLevySpecificDataType                `xml:"urn:com.workday/bsvc Federal_Tax_Levy_Data,omitempty"`
	StateTaxLevyOrderData                    *StateTaxLevyOrderDataType                     `xml:"urn:com.workday/bsvc State_Tax_Levy_Order_Data,omitempty"`
	SupportOrderData                         *PayrollSupportOrderDataType                   `xml:"urn:com.workday/bsvc Support_Order_Data,omitempty"`
	PayrollSupportOrderLumpSumData           *PayrollSupportOrderLumpSumDataType            `xml:"urn:com.workday/bsvc Payroll_Support_Order_Lump_Sum_Data,omitempty"`
	WageAssignmentData                       *WageAssignmentSpecificDataType                `xml:"urn:com.workday/bsvc Wage_Assignment_Data,omitempty"`
	FederalStudentLoanData                   []FederalStudentLoanSpecificDataType           `xml:"urn:com.workday/bsvc Federal_Student_Loan_Data,omitempty"`
	WithholdingOrderFeeData                  []WithholdingOrderFeeDataType                  `xml:"urn:com.workday/bsvc Withholding_Order_Fee_Data,omitempty"`
	WithholdingOrderWithholdingFrequencyData []WithholdingOrderWithholdingFrequencyDataType `xml:"urn:com.workday/bsvc Withholding_Order_Withholding_Frequency_Data,omitempty"`
	CurrencyReference                        *CurrencyObjectType                            `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	DeductionRecipientAddressData            []AddressInformationDataType                   `xml:"urn:com.workday/bsvc Deduction_Recipient_Address_Data,omitempty"`
	WithholdingOrderID                       string                                         `xml:"urn:com.workday/bsvc Withholding_Order_ID,omitempty"`
	CopiedMessage                            string                                         `xml:"urn:com.workday/bsvc Copied_Message,attr,omitempty"`
	CopiedTo                                 string                                         `xml:"urn:com.workday/bsvc Copied_To,attr,omitempty"`
	CopiedFrom                               string                                         `xml:"urn:com.workday/bsvc Copied_From,attr,omitempty"`
	ProcessingMessage                        string                                         `xml:"urn:com.workday/bsvc Processing_Message,attr,omitempty"`
	CopiedFromDate                           time.Time                                      `xml:"urn:com.workday/bsvc Copied_From_Date,attr,omitempty"`
	CopiedToDate                             time.Time                                      `xml:"urn:com.workday/bsvc Copied_To_Date,attr,omitempty"`
}

func (t *PayrollInvoluntaryWithholdingOrderDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollInvoluntaryWithholdingOrderDataType
	var layout struct {
		*T
		OrderDate      *xsdDate     `xml:"urn:com.workday/bsvc Order_Date"`
		ReceivedDate   *xsdDateTime `xml:"urn:com.workday/bsvc Received_Date"`
		BeginDate      *xsdDate     `xml:"urn:com.workday/bsvc Begin_Date"`
		EndDate        *xsdDate     `xml:"urn:com.workday/bsvc End_Date,omitempty"`
		CopiedFromDate *xsdDate     `xml:"urn:com.workday/bsvc Copied_From_Date,attr,omitempty"`
		CopiedToDate   *xsdDate     `xml:"urn:com.workday/bsvc Copied_To_Date,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.OrderDate = (*xsdDate)(&layout.T.OrderDate)
	layout.ReceivedDate = (*xsdDateTime)(&layout.T.ReceivedDate)
	layout.BeginDate = (*xsdDate)(&layout.T.BeginDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	layout.CopiedFromDate = (*xsdDate)(&layout.T.CopiedFromDate)
	layout.CopiedToDate = (*xsdDate)(&layout.T.CopiedToDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollInvoluntaryWithholdingOrderDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollInvoluntaryWithholdingOrderDataType
	var overlay struct {
		*T
		OrderDate      *xsdDate     `xml:"urn:com.workday/bsvc Order_Date"`
		ReceivedDate   *xsdDateTime `xml:"urn:com.workday/bsvc Received_Date"`
		BeginDate      *xsdDate     `xml:"urn:com.workday/bsvc Begin_Date"`
		EndDate        *xsdDate     `xml:"urn:com.workday/bsvc End_Date,omitempty"`
		CopiedFromDate *xsdDate     `xml:"urn:com.workday/bsvc Copied_From_Date,attr,omitempty"`
		CopiedToDate   *xsdDate     `xml:"urn:com.workday/bsvc Copied_To_Date,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.OrderDate = (*xsdDate)(&overlay.T.OrderDate)
	overlay.ReceivedDate = (*xsdDateTime)(&overlay.T.ReceivedDate)
	overlay.BeginDate = (*xsdDate)(&overlay.T.BeginDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	overlay.CopiedFromDate = (*xsdDate)(&overlay.T.CopiedFromDate)
	overlay.CopiedToDate = (*xsdDate)(&overlay.T.CopiedToDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Involuntary Withholding Order Request Criteria
type PayrollInvoluntaryWithholdingOrderRequestCriteriaType struct {
	WorkerReference []WorkerObjectType               `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	TypeReference   []WithholdingOrderTypeObjectType `xml:"urn:com.workday/bsvc Type_Reference,omitempty"`
	EffectiveAsOf   *time.Time                       `xml:"urn:com.workday/bsvc Effective_As_Of,omitempty"`
	CaseReference   []WithholdingOrderCaseObjectType `xml:"urn:com.workday/bsvc Case_Reference,omitempty"`
	CaseNumber      string                           `xml:"urn:com.workday/bsvc Case_Number,omitempty"`
}

func (t *PayrollInvoluntaryWithholdingOrderRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollInvoluntaryWithholdingOrderRequestCriteriaType
	var layout struct {
		*T
		EffectiveAsOf *xsdDate `xml:"urn:com.workday/bsvc Effective_As_Of,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveAsOf = (*xsdDate)(layout.T.EffectiveAsOf)
	return e.EncodeElement(layout, start)
}
func (t *PayrollInvoluntaryWithholdingOrderRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollInvoluntaryWithholdingOrderRequestCriteriaType
	var overlay struct {
		*T
		EffectiveAsOf *xsdDate `xml:"urn:com.workday/bsvc Effective_As_Of,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveAsOf = (*xsdDate)(overlay.T.EffectiveAsOf)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Involuntary Withholding Order Request References
type PayrollInvoluntaryWithholdingOrderRequestReferencesType struct {
	PayrollInvoluntaryWithholdingOrderReference []WithholdingOrderObjectType `xml:"urn:com.workday/bsvc Payroll_Involuntary_Withholding_Order_Reference"`
}

// Get Payroll Involuntary Withholding Orders Response
type PayrollInvoluntaryWithholdingOrderResponseDataType struct {
	PayrollInvoluntaryWithholdingOrder []PayrollInvoluntaryWithholdingOrderType `xml:"urn:com.workday/bsvc Payroll_Involuntary_Withholding_Order,omitempty"`
}

// Payroll Involuntary Withholding Order Response Group
type PayrollInvoluntaryWithholdingOrderResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Payroll Involuntary Withholding Order Response Data
type PayrollInvoluntaryWithholdingOrderType struct {
	PayrollInvoluntaryWithholdingOrderReference *WithholdingOrderObjectType                  `xml:"urn:com.workday/bsvc Payroll_Involuntary_Withholding_Order_Reference,omitempty"`
	PayrollInvoluntaryWithholdingOrderData      []PayrollInvoluntaryWithholdingOrderDataType `xml:"urn:com.workday/bsvc Payroll_Involuntary_Withholding_Order_Data,omitempty"`
}

// Element to hold all information about the Payroll Limit Override Data for Put or Get.
type PayrollLimitOverrideDataType struct {
	ID                    string                     `xml:"urn:com.workday/bsvc ID,omitempty"`
	WorkerReference       *WorkerObjectType          `xml:"urn:com.workday/bsvc Worker_Reference"`
	PositionReference     *PositionElementObjectType `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	LimitInterfaceData    []LimitInterfaceDataType   `xml:"urn:com.workday/bsvc Limit_Interface_Data"`
	PeriodStartDate       time.Time                  `xml:"urn:com.workday/bsvc Period_Start_Date"`
	PeriodEndDate         *time.Time                 `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	UseLimitOverrideValue *bool                      `xml:"urn:com.workday/bsvc Use_Limit_Override_Value,omitempty"`
	LimitOverrideValue    float64                    `xml:"urn:com.workday/bsvc Limit_Override_Value,omitempty"`
}

func (t *PayrollLimitOverrideDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollLimitOverrideDataType
	var layout struct {
		*T
		PeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Period_Start_Date"`
		PeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PeriodStartDate = (*xsdDate)(&layout.T.PeriodStartDate)
	layout.PeriodEndDate = (*xsdDate)(layout.T.PeriodEndDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollLimitOverrideDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollLimitOverrideDataType
	var overlay struct {
		*T
		PeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Period_Start_Date"`
		PeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PeriodStartDate = (*xsdDate)(&overlay.T.PeriodStartDate)
	overlay.PeriodEndDate = (*xsdDate)(overlay.T.PeriodEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PayrollLimitOverrideObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollLimitOverrideObjectType struct {
	ID         []PayrollLimitOverrideObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element to hold all criteria used to select Payroll Limit Override.
type PayrollLimitOverrideRequestCriteriaType struct {
	WorkerReference       []WorkerObjectType         `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	PayComponentReference []LimitInterfaceObjectType `xml:"urn:com.workday/bsvc Pay_Component_Reference,omitempty"`
	PeriodStartDate       *time.Time                 `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
	PeriodEndDate         *time.Time                 `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
}

func (t *PayrollLimitOverrideRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollLimitOverrideRequestCriteriaType
	var layout struct {
		*T
		PeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
		PeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PeriodStartDate = (*xsdDate)(layout.T.PeriodStartDate)
	layout.PeriodEndDate = (*xsdDate)(layout.T.PeriodEndDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollLimitOverrideRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollLimitOverrideRequestCriteriaType
	var overlay struct {
		*T
		PeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
		PeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PeriodStartDate = (*xsdDate)(overlay.T.PeriodStartDate)
	overlay.PeriodEndDate = (*xsdDate)(overlay.T.PeriodEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Element to hold all Payroll Limit Override Reference Response.
type PayrollLimitOverrideRequestReferencesType struct {
	PayrollLimitOverrideReference []PayrollLimitOverrideObjectType `xml:"urn:com.workday/bsvc Payroll_Limit_Override_Reference"`
}

// Element to hold all data to Respond about Payroll Limit Override.
type PayrollLimitOverrideResponseDataType struct {
	PayrollLimitOverride []PayrollLimitOverrideType `xml:"urn:com.workday/bsvc Payroll_Limit_Override,omitempty"`
}

// Element to hold all information about the Payroll Limit Override Response Group
type PayrollLimitOverrideResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element to hold all information about Payroll Limit Override.
type PayrollLimitOverrideType struct {
	PayrollLimitOverrideReference *PayrollLimitOverrideObjectType `xml:"urn:com.workday/bsvc Payroll_Limit_Override_Reference,omitempty"`
	PayrollLimitOverrideData      *PayrollLimitOverrideDataType   `xml:"urn:com.workday/bsvc Payroll_Limit_Override_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollLocalAuthorityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollLocalAuthorityObjectType struct {
	ID         []PayrollLocalAuthorityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollLocalCityAuthorityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollLocalCityAuthorityObjectType struct {
	ID         []PayrollLocalCityAuthorityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollLocalCountyAuthorityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollLocalCountyAuthorityObjectType struct {
	ID         []PayrollLocalCountyAuthorityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollLocalSchoolDistrictAuthorityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollLocalSchoolDistrictAuthorityObjectType struct {
	ID         []PayrollLocalSchoolDistrictAuthorityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payroll Marital Status Reference
type PayrollMaritalStatusReferenceType struct {
	PayrollMaritalStatusReference string `xml:"urn:com.workday/bsvc Payroll_Marital_Status_Reference,omitempty"`
}

// Element containing off-cycle payment request to add/update off cycle payment from external systems into Workday.
type PayrollOffcyclePaymentDataType struct {
	BatchID                    string                          `xml:"urn:com.workday/bsvc Batch_ID,omitempty"`
	PaymentID                  string                          `xml:"urn:com.workday/bsvc Payment_ID"`
	EmployeeReference          *EmployeeObjectType             `xml:"urn:com.workday/bsvc Employee_Reference"`
	PaymentDate                time.Time                       `xml:"urn:com.workday/bsvc Payment_Date"`
	PeriodDate                 time.Time                       `xml:"urn:com.workday/bsvc Period_Date"`
	PaymentPriority            float64                         `xml:"urn:com.workday/bsvc Payment_Priority"`
	RunCategoryReference       *RunCategoryObjectType          `xml:"urn:com.workday/bsvc Run_Category_Reference,omitempty"`
	PayGroupReference          *PayGroupObjectType             `xml:"urn:com.workday/bsvc Pay_Group_Reference,omitempty"`
	ResultTypeReference        *PayrollOffcycleTypeObjectType  `xml:"urn:com.workday/bsvc Result_Type_Reference"`
	Replacement                *bool                           `xml:"urn:com.workday/bsvc Replacement,omitempty"`
	ReasonReference            *PayrollActionReasonObjectType  `xml:"urn:com.workday/bsvc Reason_Reference"`
	OnDemandPaymentData        *OnDemandPaymentDataType        `xml:"urn:com.workday/bsvc On_Demand_Payment_Data,omitempty"`
	ManualPaymentData          *ManualPaymentDataType          `xml:"urn:com.workday/bsvc Manual_Payment_Data,omitempty"`
	ResultWorktagOverridesData *ResultWorktagOverridesDataType `xml:"urn:com.workday/bsvc Result_Worktag_Overrides_Data,omitempty"`
	OffcycleInputData          []OffcycleInputDataType         `xml:"urn:com.workday/bsvc Off-cycle_Input_Data,omitempty"`
}

func (t *PayrollOffcyclePaymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollOffcyclePaymentDataType
	var layout struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date"`
		PeriodDate  *xsdDate `xml:"urn:com.workday/bsvc Period_Date"`
	}
	layout.T = (*T)(t)
	layout.PaymentDate = (*xsdDate)(&layout.T.PaymentDate)
	layout.PeriodDate = (*xsdDate)(&layout.T.PeriodDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollOffcyclePaymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollOffcyclePaymentDataType
	var overlay struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date"`
		PeriodDate  *xsdDate `xml:"urn:com.workday/bsvc Period_Date"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDate = (*xsdDate)(&overlay.T.PaymentDate)
	overlay.PeriodDate = (*xsdDate)(&overlay.T.PeriodDate)
	return d.DecodeElement(&overlay, &start)
}

// Get Payroll Off-cycle payment response.
type PayrollOffcyclePaymentResponseDataType struct {
	BatchID                    string                           `xml:"urn:com.workday/bsvc Batch_ID,omitempty"`
	PaymentID                  string                           `xml:"urn:com.workday/bsvc Payment_ID,omitempty"`
	EmployeeReference          *EmployeeObjectType              `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	PaymentDate                *time.Time                       `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	PeriodDate                 *time.Time                       `xml:"urn:com.workday/bsvc Period_Date,omitempty"`
	PaymentPriority            float64                          `xml:"urn:com.workday/bsvc Payment_Priority,omitempty"`
	RunCategoryReference       *RunCategoryObjectType           `xml:"urn:com.workday/bsvc Run_Category_Reference,omitempty"`
	PayGroupReference          *PayGroupObjectType              `xml:"urn:com.workday/bsvc Pay_Group_Reference,omitempty"`
	ResultTypeReference        *PayrollOffcycleTypeObjectType   `xml:"urn:com.workday/bsvc Result_Type_Reference,omitempty"`
	Replacement                *bool                            `xml:"urn:com.workday/bsvc Replacement,omitempty"`
	ReasonReference            *PayrollActionReasonObjectType   `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
	OnDemandPaymentData        *OnDemandPaymentDataType         `xml:"urn:com.workday/bsvc On_Demand_Payment_Data,omitempty"`
	ManualPaymentData          *ManualPaymentDataType           `xml:"urn:com.workday/bsvc Manual_Payment_Data,omitempty"`
	ResultWorktagOverridesData []ResultWorktagOverridesDataType `xml:"urn:com.workday/bsvc Result_Worktag_Overrides_Data,omitempty"`
	OffcycleInputData          []OffcycleInputDataType          `xml:"urn:com.workday/bsvc Off-cycle_Input_Data,omitempty"`
}

func (t *PayrollOffcyclePaymentResponseDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollOffcyclePaymentResponseDataType
	var layout struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
		PeriodDate  *xsdDate `xml:"urn:com.workday/bsvc Period_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PaymentDate = (*xsdDate)(layout.T.PaymentDate)
	layout.PeriodDate = (*xsdDate)(layout.T.PeriodDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollOffcyclePaymentResponseDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollOffcyclePaymentResponseDataType
	var overlay struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
		PeriodDate  *xsdDate `xml:"urn:com.workday/bsvc Period_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDate = (*xsdDate)(overlay.T.PaymentDate)
	overlay.PeriodDate = (*xsdDate)(overlay.T.PeriodDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Off-cycle Payment element
type PayrollOffcyclePaymentType struct {
	PayrollOffcycleResultOverrideReference *PayrollOffcycleResultOverrideObjectType `xml:"urn:com.workday/bsvc Payroll_Off-cycle_Result_Override_Reference,omitempty"`
	PayrollOffcyclePaymentData             []PayrollOffcyclePaymentResponseDataType `xml:"urn:com.workday/bsvc Payroll_Off-cycle_Payment_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollOffcycleResultOverrideObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollOffcycleResultOverrideObjectType struct {
	ID         []PayrollOffcycleResultOverrideObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payroll Off-cycle Result Override Request Criteria element.
type PayrollOffcycleResultOverrideRequestCriteriaType struct {
	PaymentID            string                          `xml:"urn:com.workday/bsvc Payment_ID,omitempty"`
	EmployeeReference    []EmployeeObjectType            `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	PeriodReference      []PeriodObjectType              `xml:"urn:com.workday/bsvc Period_Reference,omitempty"`
	RunCategoryReference []RunCategoryObjectType         `xml:"urn:com.workday/bsvc Run_Category_Reference,omitempty"`
	PayGroupReference    []PayGroupObjectType            `xml:"urn:com.workday/bsvc Pay_Group_Reference,omitempty"`
	BatchID              string                          `xml:"urn:com.workday/bsvc Batch_ID,omitempty"`
	PaymentDate          *time.Time                      `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	ReasonReference      []PayrollActionReasonObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
}

func (t *PayrollOffcycleResultOverrideRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollOffcycleResultOverrideRequestCriteriaType
	var layout struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PaymentDate = (*xsdDate)(layout.T.PaymentDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollOffcycleResultOverrideRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollOffcycleResultOverrideRequestCriteriaType
	var overlay struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDate = (*xsdDate)(overlay.T.PaymentDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Off-cycle Result Override Request References element
type PayrollOffcycleResultOverrideRequestReferencesType struct {
	PayrollOffcyclePaymentReference []PayrollOffcycleResultOverrideObjectType `xml:"urn:com.workday/bsvc Payroll_Off-cycle_Payment_Reference"`
}

// Payroll Off-cycle Result Override Response Data element
type PayrollOffcycleResultOverrideResponseDataType struct {
	PayrollOffcycleResultOverride []PayrollOffcyclePaymentType `xml:"urn:com.workday/bsvc Payroll_Off-cycle_Result_Override,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollOffcycleTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollOffcycleTypeObjectType struct {
	ID         []PayrollOffcycleTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollOtherAuthorityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollOtherAuthorityObjectType struct {
	ID         []PayrollOtherAuthorityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Local City Tax Election
type PayrollPayeeCityElectionforStateandLocalType struct {
	EffectiveDate                      *time.Time                        `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	TaxAddressTypeReference            *TaxAddressTypeObjectType         `xml:"urn:com.workday/bsvc Tax_Address_Type_Reference"`
	PayrollLocalCityAuthorityReference *PayrollLocalAuthorityObjectType  `xml:"urn:com.workday/bsvc Payroll_Local_City_Authority_Reference,omitempty"`
	Exemptions                         float64                           `xml:"urn:com.workday/bsvc Exemptions,omitempty"`
	WithholdingPercentReference        *PayrollConstantPercentObjectType `xml:"urn:com.workday/bsvc Withholding_Percent_Reference,omitempty"`
	PreviousEmployerDeductedAmount     float64                           `xml:"urn:com.workday/bsvc Previous_Employer_Deducted_Amount,omitempty"`
	AdditionalAmount                   float64                           `xml:"urn:com.workday/bsvc Additional_Amount,omitempty"`
	NonresidentAlien                   *bool                             `xml:"urn:com.workday/bsvc Nonresident_Alien,omitempty"`
	CertificateofNonResidence          *bool                             `xml:"urn:com.workday/bsvc Certificate_of_Non-Residence,omitempty"`
	AllocationPercent                  float64                           `xml:"urn:com.workday/bsvc Allocation_Percent,omitempty"`
	LowIncomeThreshold                 float64                           `xml:"urn:com.workday/bsvc Low_Income_Threshold,omitempty"`
	Exempt                             *bool                             `xml:"urn:com.workday/bsvc Exempt,omitempty"`
	ExemptReasonReference              *PayrollConstantTextObjectType    `xml:"urn:com.workday/bsvc Exempt_Reason_Reference,omitempty"`
	InactivateStateTax                 *bool                             `xml:"urn:com.workday/bsvc Inactivate_State_Tax,omitempty"`
	PrimaryEITPennsylvania             *bool                             `xml:"urn:com.workday/bsvc Primary_EIT__Pennsylvania_,omitempty"`
	NotSubjecttoEITPennsylvania        *bool                             `xml:"urn:com.workday/bsvc Not_Subject_to_EIT__Pennsylvania_,omitempty"`
}

func (t *PayrollPayeeCityElectionforStateandLocalType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeCityElectionforStateandLocalType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeCityElectionforStateandLocalType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeCityElectionforStateandLocalType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// State County Tax Election Data
type PayrollPayeeCountyElectionforStateandLocalType struct {
	EffectiveDate                        *time.Time                       `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	TaxAddressTypeReference              *TaxAddressTypeObjectType        `xml:"urn:com.workday/bsvc Tax_Address_Type_Reference"`
	PayrollLocalCountyAuthorityReference *PayrollLocalAuthorityObjectType `xml:"urn:com.workday/bsvc Payroll_Local_County_Authority_Reference"`
	CountyAdditionalAmount               float64                          `xml:"urn:com.workday/bsvc County_Additional_Amount,omitempty"`
	InactivateStateTax                   *bool                            `xml:"urn:com.workday/bsvc Inactivate_State_Tax,omitempty"`
}

func (t *PayrollPayeeCountyElectionforStateandLocalType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeCountyElectionforStateandLocalType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeCountyElectionforStateandLocalType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeCountyElectionforStateandLocalType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Earned Income Credit Election
type PayrollPayeeEICElectionforStateandLocalType struct {
	EffectiveDate      *time.Time `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	Exempt             *bool      `xml:"urn:com.workday/bsvc Exempt,omitempty"`
	NumberofAllowances float64    `xml:"urn:com.workday/bsvc Number_of_Allowances,omitempty"`
}

func (t *PayrollPayeeEICElectionforStateandLocalType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeEICElectionforStateandLocalType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeEICElectionforStateandLocalType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeEICElectionforStateandLocalType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Payee FICA Medicare and OASDI Data
type PayrollPayeeFICADataType struct {
	WorkerReference    *WorkerObjectType          `xml:"urn:com.workday/bsvc Worker_Reference"`
	PositionReference  *PositionElementObjectType `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	AllPositions       *bool                      `xml:"urn:com.workday/bsvc All_Positions,omitempty"`
	EffectiveAsOf      time.Time                  `xml:"urn:com.workday/bsvc Effective_As_Of"`
	ApplyToWorker      *bool                      `xml:"urn:com.workday/bsvc Apply_To_Worker,omitempty"`
	MedicareExemptData []MedicareExemptDataType   `xml:"urn:com.workday/bsvc Medicare_Exempt_Data,omitempty"`
	OASDIExemptData    []OASDIExemptDataType      `xml:"urn:com.workday/bsvc OASDI_Exempt_Data,omitempty"`
}

func (t *PayrollPayeeFICADataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeFICADataType
	var layout struct {
		*T
		EffectiveAsOf *xsdDate `xml:"urn:com.workday/bsvc Effective_As_Of"`
	}
	layout.T = (*T)(t)
	layout.EffectiveAsOf = (*xsdDate)(&layout.T.EffectiveAsOf)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeFICADataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeFICADataType
	var overlay struct {
		*T
		EffectiveAsOf *xsdDate `xml:"urn:com.workday/bsvc Effective_As_Of"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveAsOf = (*xsdDate)(&overlay.T.EffectiveAsOf)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PayrollPayeeFICAExemptReasonObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollPayeeFICAExemptReasonObjectType struct {
	ID         []PayrollPayeeFICAExemptReasonObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payroll Payee FUTA Data
type PayrollPayeeFUTADataType struct {
	WorkerReference  *WorkerObjectType  `xml:"urn:com.workday/bsvc Worker_Reference"`
	CompanyReference *CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference"`
	EffectiveAsOf    time.Time          `xml:"urn:com.workday/bsvc Effective_As_Of"`
	ExemptIndicator  *bool              `xml:"urn:com.workday/bsvc Exempt_Indicator,omitempty"`
}

func (t *PayrollPayeeFUTADataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeFUTADataType
	var layout struct {
		*T
		EffectiveAsOf *xsdDate `xml:"urn:com.workday/bsvc Effective_As_Of"`
	}
	layout.T = (*T)(t)
	layout.EffectiveAsOf = (*xsdDate)(&layout.T.EffectiveAsOf)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeFUTADataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeFUTADataType
	var overlay struct {
		*T
		EffectiveAsOf *xsdDate `xml:"urn:com.workday/bsvc Effective_As_Of"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveAsOf = (*xsdDate)(&overlay.T.EffectiveAsOf)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PayrollPayeeFUTAObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollPayeeFUTAObjectType struct {
	ID         []PayrollPayeeFUTAObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element to hold all criteria used to select which FUTAs to return
type PayrollPayeeFUTARequestCriteriaType struct {
	WorkerReference  []WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	CompanyReference *CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	EffectiveDate    *time.Time         `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
}

func (t *PayrollPayeeFUTARequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeFUTARequestCriteriaType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeFUTARequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeFUTARequestCriteriaType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Element to hold all Payroll Payee Request References
type PayrollPayeeFUTARequestReferencesType struct {
	PayrollPayeeFUTAReference []PayrollPayeeFUTAObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_FUTA_Reference"`
}

// Element to hold all data to Respond with about the Payroll Payee FUTAs
type PayrollPayeeFUTAResponseDataType struct {
	PayrollPayeeFUTA []PayrollPayeeFUTAType `xml:"urn:com.workday/bsvc Payroll_Payee_FUTA,omitempty"`
}

// Element to hold all information about the Payroll Payee FUTA Response Group
type PayrollPayeeFUTAResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element to hold all information about a Payroll Payee FUTA
type PayrollPayeeFUTAType struct {
	PayrollPayeeFUTAReference *PayrollPayeeFUTAObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_FUTA_Reference,omitempty"`
	PayrollPayeeFUTAData      *PayrollPayeeFUTADataType   `xml:"urn:com.workday/bsvc Payroll_Payee_FUTA_Data,omitempty"`
}

// Element to hold Payroll Payee Ongoing Jurisdiction Splits Tax Election Data Information
type PayrollPayeeOngoingJurisdictionSplitsTaxElectionDataType struct {
	StateReference    *PayrollStateAuthorityObjectType       `xml:"urn:com.workday/bsvc State_Reference"`
	CountyReference   *PayrollLocalCountyAuthorityObjectType `xml:"urn:com.workday/bsvc County_Reference,omitempty"`
	CityReference     *PayrollLocalCityAuthorityObjectType   `xml:"urn:com.workday/bsvc City_Reference,omitempty"`
	OtherReference    *PayrollOtherAuthorityObjectType       `xml:"urn:com.workday/bsvc Other_Reference,omitempty"`
	AllocationPercent float64                                `xml:"urn:com.workday/bsvc Allocation_Percent"`
}

// Element to hold Payroll Payee Ongoing Jurisdiction Tax Election Data Information
type PayrollPayeeOngoingWorkJurisdictionTaxElectionDataType struct {
	WorkerforTaxElectionReference        *WorkerObjectType                                          `xml:"urn:com.workday/bsvc Worker_for_Tax_Election_Reference"`
	EffectiveDate                        time.Time                                                  `xml:"urn:com.workday/bsvc Effective_Date"`
	CompanyforTaxElectionReference       *CompanyObjectType                                         `xml:"urn:com.workday/bsvc Company_for_Tax_Election_Reference,omitempty"`
	Inactive                             *bool                                                      `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	PayrollPayeeOngoingJurisdictionSplit []PayrollPayeeOngoingJurisdictionSplitsTaxElectionDataType `xml:"urn:com.workday/bsvc Payroll_Payee_Ongoing_Jurisdiction_Split,omitempty"`
}

func (t *PayrollPayeeOngoingWorkJurisdictionTaxElectionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeOngoingWorkJurisdictionTaxElectionDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeOngoingWorkJurisdictionTaxElectionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeOngoingWorkJurisdictionTaxElectionDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Element to hold Payroll Payee Ongoing Jurisdiction Request Criteria Information
type PayrollPayeeOngoingWorkJurisdictionTaxElectionRequestCriteriaType struct {
	WorkerReference  []WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	CompanyReference *CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	EffectiveasOf    *time.Time         `xml:"urn:com.workday/bsvc Effective_as_Of,omitempty"`
}

func (t *PayrollPayeeOngoingWorkJurisdictionTaxElectionRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeOngoingWorkJurisdictionTaxElectionRequestCriteriaType
	var layout struct {
		*T
		EffectiveasOf *xsdDate `xml:"urn:com.workday/bsvc Effective_as_Of,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveasOf = (*xsdDate)(layout.T.EffectiveasOf)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeOngoingWorkJurisdictionTaxElectionRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeOngoingWorkJurisdictionTaxElectionRequestCriteriaType
	var overlay struct {
		*T
		EffectiveasOf *xsdDate `xml:"urn:com.workday/bsvc Effective_as_Of,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveasOf = (*xsdDate)(overlay.T.EffectiveasOf)
	return d.DecodeElement(&overlay, &start)
}

// Element to hold Payroll Payee Ongoing Jurisdiction Request References information
type PayrollPayeeOngoingWorkJurisdictionTaxElectionRequestReferencesType struct {
	PayrollPayeeOngoingJurisdictionReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_Ongoing_Jurisdiction_Reference"`
}

// Element to hold Payroll Payee Ongoing Jurisdiction Response Data Information
type PayrollPayeeOngoingWorkJurisdictionTaxElectionResponseDataType struct {
	PayrollPayeeOngoingWorkJurisdictionTaxElection []PayrollPayeeOngoingWorkJurisdictionTaxElectionType `xml:"urn:com.workday/bsvc Payroll_Payee_Ongoing_Work_Jurisdiction_Tax_Election,omitempty"`
}

// Element to hold Payroll Payee Ongoing Jurisdiction Response Data Information
type PayrollPayeeOngoingWorkJurisdictionTaxElectionType struct {
	PayrollPayeeOngoingWorkJurisdictionTaxElectionReference *UniqueIdentifierObjectType                             `xml:"urn:com.workday/bsvc Payroll_Payee_Ongoing_Work_Jurisdiction_Tax_Election_Reference,omitempty"`
	PayrollPayeeOngoingWorkJurisdictionTaxElectionData      *PayrollPayeeOngoingWorkJurisdictionTaxElectionDataType `xml:"urn:com.workday/bsvc Payroll_Payee_Ongoing_Work_Jurisdiction_Tax_Election_Data,omitempty"`
}

// Tax Election for Other State Taxes
type PayrollPayeeOtherElectionforStateandLocalType struct {
	EffectiveDate                  *time.Time                       `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	TaxAddressTypeReference        *TaxAddressTypeObjectType        `xml:"urn:com.workday/bsvc Tax_Address_Type_Reference"`
	PayrollOtherAuthorityReference *PayrollLocalAuthorityObjectType `xml:"urn:com.workday/bsvc Payroll_Other_Authority_Reference,omitempty"`
	Exempt                         *bool                            `xml:"urn:com.workday/bsvc Exempt,omitempty"`
	Inactive                       *bool                            `xml:"urn:com.workday/bsvc Inactive,omitempty"`
}

func (t *PayrollPayeeOtherElectionforStateandLocalType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeOtherElectionforStateandLocalType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeOtherElectionforStateandLocalType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeOtherElectionforStateandLocalType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Payee PT1 Data
type PayrollPayeePT1DataType struct {
	PayrollPayeePT1Reference                           *PayrollPayeePT1ObjectType       `xml:"urn:com.workday/bsvc Payroll_Payee_PT1_Reference,omitempty"`
	CompanyReference                                   *CompanyObjectType               `xml:"urn:com.workday/bsvc Company_Reference"`
	WorkerReference                                    *WorkerObjectType                `xml:"urn:com.workday/bsvc Worker_Reference"`
	EffectiveAsOf                                      time.Time                        `xml:"urn:com.workday/bsvc Effective_As_Of"`
	PayrollStateAuthorityReference                     *PayrollStateAuthorityObjectType `xml:"urn:com.workday/bsvc Payroll_State_Authority_Reference"`
	BasicPersonalAmount                                float64                          `xml:"urn:com.workday/bsvc Basic_Personal_Amount,omitempty"`
	ChildAmount                                        float64                          `xml:"urn:com.workday/bsvc Child_Amount,omitempty"`
	AgeAmount                                          float64                          `xml:"urn:com.workday/bsvc Age_Amount,omitempty"`
	AgeAmountSupplement                                float64                          `xml:"urn:com.workday/bsvc Age_Amount_Supplement,omitempty"`
	SeniorSupplementaryAmountforWithholding            float64                          `xml:"urn:com.workday/bsvc Senior_Supplementary_Amount_for_Withholding,omitempty"`
	AmountforWorkers65orOlder                          float64                          `xml:"urn:com.workday/bsvc Amount_for_Workers_65_or_Older,omitempty"`
	PensionIncomeAmount                                float64                          `xml:"urn:com.workday/bsvc Pension_Income_Amount,omitempty"`
	TuitionEducationandTextbookAmounts                 float64                          `xml:"urn:com.workday/bsvc Tuition__Education_and_Textbook_Amounts,omitempty"`
	DisabilityAmount                                   float64                          `xml:"urn:com.workday/bsvc Disability_Amount,omitempty"`
	SpouseorCommonlawPartnerAmount                     float64                          `xml:"urn:com.workday/bsvc Spouse_or_Common-law_Partner_Amount,omitempty"`
	SpouseorCommonlawPartnerAmountSupplement           float64                          `xml:"urn:com.workday/bsvc Spouse_or_Common-law_Partner_Amount_Supplement,omitempty"`
	AdditionalAmount                                   float64                          `xml:"urn:com.workday/bsvc Additional_Amount,omitempty"`
	DeductionforLivinginaPrescribedZone                float64                          `xml:"urn:com.workday/bsvc Deduction_for_Living_in_a_Prescribed_Zone,omitempty"`
	AmountforanEligibleDependant                       float64                          `xml:"urn:com.workday/bsvc Amount_for_an_Eligible_Dependant,omitempty"`
	AmountforanEligibleDependantSupplement             float64                          `xml:"urn:com.workday/bsvc Amount_for_an_Eligible_Dependant_Supplement,omitempty"`
	CaregiverAmount                                    float64                          `xml:"urn:com.workday/bsvc Caregiver_Amount,omitempty"`
	AmountforInfirmDependantsAge18orOlder              float64                          `xml:"urn:com.workday/bsvc Amount_for_Infirm_Dependants_Age_18_or_Older,omitempty"`
	AmountsTransferredfromyourSpouseorCommonlawPartner float64                          `xml:"urn:com.workday/bsvc Amounts_Transferred_from_your_Spouse_or_Common-law_Partner,omitempty"`
	AmountsTransferredfromaDependant                   float64                          `xml:"urn:com.workday/bsvc Amounts_Transferred_from_a_Dependant,omitempty"`
	ManitobaFamilyTaxBenefitAmount                     float64                          `xml:"urn:com.workday/bsvc Manitoba_Family_Tax_Benefit_Amount,omitempty"`
	TotalIncomelessthanTotalClaimAmount                *bool                            `xml:"urn:com.workday/bsvc Total_Income_less_than_Total_Claim_Amount,omitempty"`
	MorethanOneEmployerReference                       *bool                            `xml:"urn:com.workday/bsvc More_than_One_Employer_Reference,omitempty"`
	DependentsUnderAge19                               float64                          `xml:"urn:com.workday/bsvc Dependents_Under_Age_19,omitempty"`
	DisabledDependents                                 float64                          `xml:"urn:com.workday/bsvc Disabled_Dependents,omitempty"`
	Exempt                                             *bool                            `xml:"urn:com.workday/bsvc Exempt,omitempty"`
	ExemptfromHealthContribution                       *bool                            `xml:"urn:com.workday/bsvc Exempt_from_Health_Contribution,omitempty"`
	LastUpdated                                        *time.Time                       `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
}

func (t *PayrollPayeePT1DataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeePT1DataType
	var layout struct {
		*T
		EffectiveAsOf *xsdDate     `xml:"urn:com.workday/bsvc Effective_As_Of"`
		LastUpdated   *xsdDateTime `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveAsOf = (*xsdDate)(&layout.T.EffectiveAsOf)
	layout.LastUpdated = (*xsdDateTime)(layout.T.LastUpdated)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeePT1DataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeePT1DataType
	var overlay struct {
		*T
		EffectiveAsOf *xsdDate     `xml:"urn:com.workday/bsvc Effective_As_Of"`
		LastUpdated   *xsdDateTime `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveAsOf = (*xsdDate)(&overlay.T.EffectiveAsOf)
	overlay.LastUpdated = (*xsdDateTime)(overlay.T.LastUpdated)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PayrollPayeePT1ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollPayeePT1ObjectType struct {
	ID         []PayrollPayeePT1ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payroll Payee PT1 Request Criteria [EL]
type PayrollPayeePT1RequestCriteriaType struct {
}

// Payroll Payee PT1 Request References
type PayrollPayeePT1RequestReferencesType struct {
	PayrollPayeePT1Reference []PayrollPayeePT1ObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_PT1_Reference"`
}

// Payroll Payee PT1 Response Data
type PayrollPayeePT1ResponseDataType struct {
	PayrollPayeePT1 []PayrollPayeePT1Type `xml:"urn:com.workday/bsvc Payroll_Payee_PT1,omitempty"`
}

// Payroll Payee PT1 Response Group
type PayrollPayeePT1ResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Payroll Payee PT1 Record Data
type PayrollPayeePT1Type struct {
	PayrollPayeePT1Reference *PayrollPayeePT1ObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_PT1_Reference,omitempty"`
	PayrollPayeePT1Data      []PayrollPayeePT1DataType  `xml:"urn:com.workday/bsvc Payroll_Payee_PT1_Data,omitempty"`
}

// Payroll Payee RPP or DPSP Registration Number Data
type PayrollPayeeRPPorDPSPRegistrationNumberDataType struct {
	PayrollPayeeRPPorDPSPRegistrationNumberReference *PayrollPayeeRPPorDPSPRegistrationNumbersObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_RPP_or_DPSP_Registration_Number_Reference,omitempty"`
	CompanyReference                                 *CompanyObjectType                                  `xml:"urn:com.workday/bsvc Company_Reference"`
	WorkerReference                                  *WorkerObjectType                                   `xml:"urn:com.workday/bsvc Worker_Reference"`
	EffectiveDate                                    time.Time                                           `xml:"urn:com.workday/bsvc Effective_Date"`
	RPPorDPSPRegistrationNumber                      string                                              `xml:"urn:com.workday/bsvc RPP_or_DPSP_Registration_Number"`
	LastUpdated                                      *time.Time                                          `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
}

func (t *PayrollPayeeRPPorDPSPRegistrationNumberDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeRPPorDPSPRegistrationNumberDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate     `xml:"urn:com.workday/bsvc Effective_Date"`
		LastUpdated   *xsdDateTime `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	layout.LastUpdated = (*xsdDateTime)(layout.T.LastUpdated)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeRPPorDPSPRegistrationNumberDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeRPPorDPSPRegistrationNumberDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate     `xml:"urn:com.workday/bsvc Effective_Date"`
		LastUpdated   *xsdDateTime `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	overlay.LastUpdated = (*xsdDateTime)(overlay.T.LastUpdated)
	return d.DecodeElement(&overlay, &start)
}

// Element to hold all Payroll Payee Request References
type PayrollPayeeRPPorDPSPRegistrationNumberRequestReferencesType struct {
	PayrollPayeeRPPorDPSPRegistrationNumberReference []PayrollPayeeRPPorDPSPRegistrationNumbersObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_RPP_or_DPSP_Registration_Number_Reference"`
}

// Element to hold all data to Respond with about the Payroll Payee RPP or DPSP Registration Numbers
type PayrollPayeeRPPorDPSPRegistrationNumberResponseDataType struct {
	PayrollPayeeRPPorDPSPRegistrationNumber []PayrollPayeeRPPorDPSPRegistrationNumberType `xml:"urn:com.workday/bsvc Payroll_Payee_RPP_or_DPSP_Registration_Number,omitempty"`
}

// Element to hold all information about the Payroll Payee RPP or DPSP Registration Number Response Group
type PayrollPayeeRPPorDPSPRegistrationNumberResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element to hold all information about the Payroll Payee RPP or DPSP Registration Number Response Group
type PayrollPayeeRPPorDPSPRegistrationNumberType struct {
	PayrollPayeeRPPorDPSPRegistrationNumberReference *PayrollPayeeRPPorDPSPRegistrationNumbersObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_RPP_or_DPSP_Registration_Number_Reference,omitempty"`
	PayrollPayeeRPPorDPSPRegistrationNumberData      []PayrollPayeeRPPorDPSPRegistrationNumberDataType   `xml:"urn:com.workday/bsvc Payroll_Payee_RPP_or_DPSP_Registration_Number_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollPayeeRPPorDPSPRegistrationNumbersObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollPayeeRPPorDPSPRegistrationNumbersObjectType struct {
	ID         []PayrollPayeeRPPorDPSPRegistrationNumbersObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// State Unemployment Exemption Data
type PayrollPayeeSUTAElectonforStateandLocalType struct {
	EffectiveDate *time.Time `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	Exempt        *bool      `xml:"urn:com.workday/bsvc Exempt,omitempty"`
}

func (t *PayrollPayeeSUTAElectonforStateandLocalType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeSUTAElectonforStateandLocalType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeSUTAElectonforStateandLocalType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeSUTAElectonforStateandLocalType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Local School District Tax Election Data
type PayrollPayeeSchoolDistrictElectionforStateandLocalType struct {
	EffectiveDate                                *time.Time                       `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	PayrollLocalSchoolDistrictAuthorityReference *PayrollLocalAuthorityObjectType `xml:"urn:com.workday/bsvc Payroll_Local_School_District_Authority_Reference,omitempty"`
	Exempt                                       *bool                            `xml:"urn:com.workday/bsvc Exempt,omitempty"`
	ExemptReasonReference                        *PayrollConstantTextObjectType   `xml:"urn:com.workday/bsvc Exempt_Reason_Reference,omitempty"`
	PreviousEmployerDeductedAmount               float64                          `xml:"urn:com.workday/bsvc Previous_Employer_Deducted_Amount,omitempty"`
	InactivateStateTax                           *bool                            `xml:"urn:com.workday/bsvc Inactivate_State_Tax,omitempty"`
}

func (t *PayrollPayeeSchoolDistrictElectionforStateandLocalType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeSchoolDistrictElectionforStateandLocalType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeSchoolDistrictElectionforStateandLocalType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeSchoolDistrictElectionforStateandLocalType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// State Income Tax Election Data
type PayrollPayeeStateElectionforStateandLocalType struct {
	EffectiveDate                                    *time.Time                          `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	MarriedFilingJointlyOptionalCalculation          *bool                               `xml:"urn:com.workday/bsvc Married_Filing_Jointly_Optional_Calculation,omitempty"`
	PayrollWithholdingStatusReference                *PayrollWithholdingStatusObjectType `xml:"urn:com.workday/bsvc Payroll_Withholding_Status_Reference,omitempty"`
	VeteranExemption                                 *bool                               `xml:"urn:com.workday/bsvc Veteran_Exemption,omitempty"`
	ExemptionforDependentsComplete                   float64                             `xml:"urn:com.workday/bsvc Exemption_for_Dependents_Complete,omitempty"`
	ExemptionforDependentsJointCustody               float64                             `xml:"urn:com.workday/bsvc Exemption_for_Dependents_Joint_Custody,omitempty"`
	AllowanceonSpecialDeduction                      *bool                               `xml:"urn:com.workday/bsvc Allowance_on_Special_Deduction,omitempty"`
	NewJerseyRateTableReference                      *PayrollWithholdingStatusObjectType `xml:"urn:com.workday/bsvc New_Jersey_Rate_Table_Reference,omitempty"`
	DependentAllowance                               float64                             `xml:"urn:com.workday/bsvc Dependent_Allowance,omitempty"`
	NumberofAllowances                               float64                             `xml:"urn:com.workday/bsvc Number_of_Allowances,omitempty"`
	NumberofAllowancesReference                      *PayrollConstantNumberObjectType    `xml:"urn:com.workday/bsvc Number_of_Allowances_Reference,omitempty"`
	EstimatedDeductions                              float64                             `xml:"urn:com.workday/bsvc Estimated_Deductions,omitempty"`
	Exemptions                                       float64                             `xml:"urn:com.workday/bsvc Exemptions,omitempty"`
	WithholdingExemption                             float64                             `xml:"urn:com.workday/bsvc Withholding_Exemption,omitempty"`
	EmployeeBlind                                    *bool                               `xml:"urn:com.workday/bsvc Employee_Blind,omitempty"`
	AdditionalAmount                                 float64                             `xml:"urn:com.workday/bsvc Additional_Amount,omitempty"`
	AdditionalPercent                                float64                             `xml:"urn:com.workday/bsvc Additional_Percent,omitempty"`
	ServicesLocalizedinIllinois                      *bool                               `xml:"urn:com.workday/bsvc Services_Localized_in_Illinois,omitempty"`
	ReciprocityTaxCreditAmount                       float64                             `xml:"urn:com.workday/bsvc Reciprocity_Tax_Credit_Amount,omitempty"`
	XMLNAME6A                                        *bool                               `xml:"urn:com.workday/bsvc XMLNAME_6A,omitempty"`
	XMLNAME6B                                        *bool                               `xml:"urn:com.workday/bsvc XMLNAME_6B,omitempty"`
	XMLNAME6C                                        *bool                               `xml:"urn:com.workday/bsvc XMLNAME_6C,omitempty"`
	XMLNAME6D                                        *bool                               `xml:"urn:com.workday/bsvc XMLNAME_6D,omitempty"`
	CertificateofNonResidence                        *bool                               `xml:"urn:com.workday/bsvc Certificate_of_Non-Residence,omitempty"`
	CertificateofResidence                           *bool                               `xml:"urn:com.workday/bsvc Certificate_of_Residence,omitempty"`
	CertificateofWithholdingExemptionandCountyStatus *bool                               `xml:"urn:com.workday/bsvc Certificate_of_Withholding_Exemption_and_County_Status,omitempty"`
	LockInLetter                                     *bool                               `xml:"urn:com.workday/bsvc Lock_In_Letter,omitempty"`
	AllocationPercent                                float64                             `xml:"urn:com.workday/bsvc Allocation_Percent,omitempty"`
	WithholdingPercentReference                      *PayrollConstantPercentObjectType   `xml:"urn:com.workday/bsvc Withholding_Percent_Reference,omitempty"`
	PayPeriodAmount                                  float64                             `xml:"urn:com.workday/bsvc Pay_Period_Amount,omitempty"`
	SpouseIndicator                                  *bool                               `xml:"urn:com.workday/bsvc Spouse_Indicator,omitempty"`
	FulltimeStudentIndicator                         *bool                               `xml:"urn:com.workday/bsvc Full-time_Student_Indicator,omitempty"`
	HeadofHousehold                                  *bool                               `xml:"urn:com.workday/bsvc Head_of_Household,omitempty"`
	AnnualTaxCredits                                 float64                             `xml:"urn:com.workday/bsvc Annual_Tax_Credits,omitempty"`
	AdditionalAllowance                              float64                             `xml:"urn:com.workday/bsvc Additional_Allowance,omitempty"`
	ReducedWithholdingAmount                         float64                             `xml:"urn:com.workday/bsvc Reduced_Withholding_Amount,omitempty"`
	NoWageNoTax                                      *bool                               `xml:"urn:com.workday/bsvc No_Wage_No_Tax,omitempty"`
	Exempt                                           *bool                               `xml:"urn:com.workday/bsvc Exempt,omitempty"`
	ExemptReasonReference                            *UniqueIdentifierObjectType         `xml:"urn:com.workday/bsvc Exempt_Reason_Reference,omitempty"`
	WithholdingSubstantiated                         *bool                               `xml:"urn:com.workday/bsvc Withholding_Substantiated,omitempty"`
	InactivateStateTax                               *bool                               `xml:"urn:com.workday/bsvc Inactivate_State_Tax,omitempty"`
	MSRRExempt                                       *bool                               `xml:"urn:com.workday/bsvc MSRR_Exempt,omitempty"`
	EntrepreneurExemption                            *bool                               `xml:"urn:com.workday/bsvc Entrepreneur_Exemption,omitempty"`
	DomicileStateReference                           *PayrollStateAuthorityObjectType    `xml:"urn:com.workday/bsvc Domicile_State_Reference,omitempty"`
	IncreaseorDecreaseWithholdingAmount              float64                             `xml:"urn:com.workday/bsvc Increase_or_Decrease_Withholding_Amount,omitempty"`
	LowerTaxRateorLowIncome                          *bool                               `xml:"urn:com.workday/bsvc Lower_Tax_Rate_or_Low_Income,omitempty"`
	ActiveDutyOklahoma                               *bool                               `xml:"urn:com.workday/bsvc Active_Duty__Oklahoma_,omitempty"`
	FortCampbellExemptKentucky                       *bool                               `xml:"urn:com.workday/bsvc Fort_Campbell_Exempt__Kentucky_,omitempty"`
}

func (t *PayrollPayeeStateElectionforStateandLocalType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeStateElectionforStateandLocalType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeStateElectionforStateandLocalType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeStateElectionforStateandLocalType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll State and Local Tax Election Response Data
type PayrollPayeeStateandLocalTaxElectionsType struct {
	WorkerReference           *WorkerObjectType          `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	CompanyReference          *CompanyObjectType         `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	EffectiveDate             *time.Time                 `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	PayrollStateAuthorityData []StateTaxElectionDataType `xml:"urn:com.workday/bsvc Payroll_State_Authority_Data"`
}

func (t *PayrollPayeeStateandLocalTaxElectionsType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeStateandLocalTaxElectionsType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeStateandLocalTaxElectionsType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeStateandLocalTaxElectionsType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PayrollPayeeT1ObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollPayeeT1ObjectType struct {
	ID         []PayrollPayeeT1ObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payroll Payee TD1 Data
type PayrollPayeeTD1DataType struct {
	PayrollPayeeTD1Reference                           *PayrollPayeeT1ObjectType   `xml:"urn:com.workday/bsvc Payroll_Payee_TD1_Reference,omitempty"`
	CompanyReference                                   *CompanyObjectType          `xml:"urn:com.workday/bsvc Company_Reference"`
	WorkerReference                                    *WorkerObjectType           `xml:"urn:com.workday/bsvc Worker_Reference"`
	EffectiveAsOf                                      time.Time                   `xml:"urn:com.workday/bsvc Effective_As_Of"`
	BasicPersonalAmount                                float64                     `xml:"urn:com.workday/bsvc Basic_Personal_Amount,omitempty"`
	ChildAmount                                        float64                     `xml:"urn:com.workday/bsvc Child_Amount,omitempty"`
	AgeAmount                                          float64                     `xml:"urn:com.workday/bsvc Age_Amount,omitempty"`
	PensionIncomeAmount                                float64                     `xml:"urn:com.workday/bsvc Pension_Income_Amount,omitempty"`
	TuitionEducationandTextbookAmounts                 float64                     `xml:"urn:com.workday/bsvc Tuition__Education_and_Textbook_Amounts,omitempty"`
	DisabilityAmount                                   float64                     `xml:"urn:com.workday/bsvc Disability_Amount,omitempty"`
	SpouseorCommonlawPartnerAmount                     float64                     `xml:"urn:com.workday/bsvc Spouse_or_Common-law_Partner_Amount,omitempty"`
	AmountforanEligibleDependant                       float64                     `xml:"urn:com.workday/bsvc Amount_for_an_Eligible_Dependant,omitempty"`
	CaregiverAmount                                    float64                     `xml:"urn:com.workday/bsvc Caregiver_Amount,omitempty"`
	AmountforInfirmDependantsAge18orOlder              float64                     `xml:"urn:com.workday/bsvc Amount_for_Infirm_Dependants_Age_18_or_Older,omitempty"`
	AmountsTransferredfromyourSpouseorCommonlawPartner float64                     `xml:"urn:com.workday/bsvc Amounts_Transferred_from_your_Spouse_or_Common-law_Partner,omitempty"`
	AmountsTransferredfromaDependant                   float64                     `xml:"urn:com.workday/bsvc Amounts_Transferred_from_a_Dependant,omitempty"`
	TotalIncomelessthanTotalClaimAmount                *bool                       `xml:"urn:com.workday/bsvc Total_Income_less_than_Total_Claim_Amount,omitempty"`
	MorethanOneEmployerorPayerattheSameTime            *bool                       `xml:"urn:com.workday/bsvc More_than_One_Employer_or_Payer_at_the_Same_Time,omitempty"`
	DeductionforLivinginaPrescribedZone                float64                     `xml:"urn:com.workday/bsvc Deduction_for_Living_in_a_Prescribed_Zone,omitempty"`
	AdditionalAmount                                   float64                     `xml:"urn:com.workday/bsvc Additional_Amount,omitempty"`
	NonResident                                        *bool                       `xml:"urn:com.workday/bsvc Non_Resident,omitempty"`
	ESSElectronicSignatureConfirmation                 *time.Time                  `xml:"urn:com.workday/bsvc ESS_Electronic_Signature_Confirmation,omitempty"`
	PayrollPayeeTaxDataEventReference                  *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_Tax_Data_Event_Reference,omitempty"`
	LastUpdated                                        *time.Time                  `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	LastUpdatedByReference                             *WorkerObjectType           `xml:"urn:com.workday/bsvc Last_Updated_By_Reference,omitempty"`
	CurrencyReference                                  *CurrencyObjectType         `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
}

func (t *PayrollPayeeTD1DataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeTD1DataType
	var layout struct {
		*T
		EffectiveAsOf                      *xsdDate     `xml:"urn:com.workday/bsvc Effective_As_Of"`
		ESSElectronicSignatureConfirmation *xsdDateTime `xml:"urn:com.workday/bsvc ESS_Electronic_Signature_Confirmation,omitempty"`
		LastUpdated                        *xsdDateTime `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveAsOf = (*xsdDate)(&layout.T.EffectiveAsOf)
	layout.ESSElectronicSignatureConfirmation = (*xsdDateTime)(layout.T.ESSElectronicSignatureConfirmation)
	layout.LastUpdated = (*xsdDateTime)(layout.T.LastUpdated)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeTD1DataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeTD1DataType
	var overlay struct {
		*T
		EffectiveAsOf                      *xsdDate     `xml:"urn:com.workday/bsvc Effective_As_Of"`
		ESSElectronicSignatureConfirmation *xsdDateTime `xml:"urn:com.workday/bsvc ESS_Electronic_Signature_Confirmation,omitempty"`
		LastUpdated                        *xsdDateTime `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveAsOf = (*xsdDate)(&overlay.T.EffectiveAsOf)
	overlay.ESSElectronicSignatureConfirmation = (*xsdDateTime)(overlay.T.ESSElectronicSignatureConfirmation)
	overlay.LastUpdated = (*xsdDateTime)(overlay.T.LastUpdated)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Payee TD1 Request Criteria
type PayrollPayeeTD1RequestCriteriaType struct {
}

// Payroll Payee TD1 Request References
type PayrollPayeeTD1RequestReferencesType struct {
	TD1ElectionReference []PayrollPayeeT1ObjectType `xml:"urn:com.workday/bsvc TD1_Election_Reference"`
}

// Payroll Payee TD1 Response Data
type PayrollPayeeTD1ResponseDataType struct {
	PayrollPayeeTD1 []PayrollPayeeTD1Type `xml:"urn:com.workday/bsvc Payroll_Payee_TD1,omitempty"`
}

// Payroll Payee TD1 Response Group
type PayrollPayeeTD1ResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Payroll Payee TD1
type PayrollPayeeTD1Type struct {
	PayrollPayeeTD1Data []PayrollPayeeTD1DataType `xml:"urn:com.workday/bsvc Payroll_Payee_TD1_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollPayeeTaxDataObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollPayeeTaxDataObjectType struct {
	ID         []PayrollPayeeTaxDataObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payroll Payee Tax Data Request Criteria element
type PayrollPayeeTaxDataRequestCriteriaType struct {
	WorkerReference   []WorkerObjectType          `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	PositionReference []PositionElementObjectType `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	EffectiveAsOf     *time.Time                  `xml:"urn:com.workday/bsvc Effective_As_Of,omitempty"`
}

func (t *PayrollPayeeTaxDataRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeTaxDataRequestCriteriaType
	var layout struct {
		*T
		EffectiveAsOf *xsdDate `xml:"urn:com.workday/bsvc Effective_As_Of,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveAsOf = (*xsdDate)(layout.T.EffectiveAsOf)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeTaxDataRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeTaxDataRequestCriteriaType
	var overlay struct {
		*T
		EffectiveAsOf *xsdDate `xml:"urn:com.workday/bsvc Effective_As_Of,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveAsOf = (*xsdDate)(overlay.T.EffectiveAsOf)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Payee Tax Data Request References element
type PayrollPayeeTaxDataRequestReferencesType struct {
	PayrollPayeeTaxDataReference []PayrollPayeeTaxDataObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_Tax_Data_Reference"`
}

// Payroll Payee Tax Data Response Data element.
type PayrollPayeeTaxDataResponseDataType struct {
	PayrollPayeeTaxData []PayrollPayeeTaxDataType `xml:"urn:com.workday/bsvc Payroll_Payee_Tax_Data,omitempty"`
}

// Payroll Payee Tax Data Response Group element
type PayrollPayeeTaxDataResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Payroll Payee Tax Data element.
type PayrollPayeeTaxDataType struct {
	PayrollPayeeTaxDataReference *PayrollPayeeTaxDataObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_Tax_Data_Reference,omitempty"`
	PayrollPayeeFICAData         []PayrollPayeeFICADataType     `xml:"urn:com.workday/bsvc Payroll_Payee_FICA_Data,omitempty"`
}

// Element to hold Payroll Payee Tax Location Mapping Data Information
type PayrollPayeeTaxLocationMappingDataType struct {
	PayrollPayeeTaxLocationMappingID string                            `xml:"urn:com.workday/bsvc Payroll_Payee_Tax_Location_Mapping_ID,omitempty"`
	WorkerReference                  *WorkerObjectType                 `xml:"urn:com.workday/bsvc Worker_Reference"`
	EffectiveDate                    time.Time                         `xml:"urn:com.workday/bsvc Effective_Date"`
	CompanyReference                 *CompanyObjectType                `xml:"urn:com.workday/bsvc Company_Reference"`
	InactiveTaxLocationMapping       *bool                             `xml:"urn:com.workday/bsvc Inactive_Tax_Location_Mapping,omitempty"`
	InactivateStateReference         []PayrollStateAuthorityObjectType `xml:"urn:com.workday/bsvc Inactivate_State_Reference,omitempty"`
}

func (t *PayrollPayeeTaxLocationMappingDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeTaxLocationMappingDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeTaxLocationMappingDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeTaxLocationMappingDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Element to hold Payroll Payee Tax Location Mapping Response Data Information
type PayrollPayeeTaxLocationMappingGetDataType struct {
	PayrollPayeeTaxLocationMappingReference *PayrollPayeeTaxLocationMappingObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_Tax_Location_Mapping_Reference,omitempty"`
	PayrollPayeeTaxLocationMappingData      *PayrollPayeeTaxLocationMappingDataType   `xml:"urn:com.workday/bsvc Payroll_Payee_Tax_Location_Mapping_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollPayeeTaxLocationMappingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollPayeeTaxLocationMappingObjectType struct {
	ID         []PayrollPayeeTaxLocationMappingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element to hold Payroll Payee Tax Location Mapping Request Criteria Information.  Utilize this element to find existing Tax Location Mappings in the System.
type PayrollPayeeTaxLocationMappingRequestCriteriaType struct {
	WorkerReference  *WorkerObjectType  `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	CompanyReference *CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	EffectiveasOf    *time.Time         `xml:"urn:com.workday/bsvc Effective_as_Of,omitempty"`
}

func (t *PayrollPayeeTaxLocationMappingRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayeeTaxLocationMappingRequestCriteriaType
	var layout struct {
		*T
		EffectiveasOf *xsdDate `xml:"urn:com.workday/bsvc Effective_as_Of,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveasOf = (*xsdDate)(layout.T.EffectiveasOf)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayeeTaxLocationMappingRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayeeTaxLocationMappingRequestCriteriaType
	var overlay struct {
		*T
		EffectiveasOf *xsdDate `xml:"urn:com.workday/bsvc Effective_as_Of,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveasOf = (*xsdDate)(overlay.T.EffectiveasOf)
	return d.DecodeElement(&overlay, &start)
}

// Element to hold Payroll Payee Tax Location Mapping Request References information
type PayrollPayeeTaxLocationMappingRequestReferencesType struct {
	PayrollPayeeTaxLocationMappingReference []PayrollPayeeTaxLocationMappingObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_Tax_Location_Mapping_Reference"`
}

// Element to hold Payroll Payee Tax Location Mapping Response Data Information
type PayrollPayeeTaxLocationMappingResponseDataType struct {
	PayrollPayeeTaxLocationMapping []PayrollPayeeTaxLocationMappingGetDataType `xml:"urn:com.workday/bsvc Payroll_Payee_Tax_Location_Mapping,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollPayeeTaxReportTypeDeliveryGroupObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollPayeeTaxReportTypeDeliveryGroupObjectType struct {
	ID         []PayrollPayeeTaxReportTypeDeliveryGroupObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollPayeeTaxTreatyUSAObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollPayeeTaxTreatyUSAObjectType struct {
	ID         []PayrollPayeeTaxTreatyUSAObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payroll Payslip Data
type PayrollPayslipDataType struct {
	RunCategoryData              *PayslipRunCategoryDataType   `xml:"urn:com.workday/bsvc Run_Category_Data,omitempty"`
	PaymentDate                  *time.Time                    `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	PayrollRemittanceData        *PayrollRemittanceDataType    `xml:"urn:com.workday/bsvc Payroll_Remittance_Data,omitempty"`
	PaymentData                  []PayslipPaymentDataType      `xml:"urn:com.workday/bsvc Payment_Data,omitempty"`
	IntegrationFieldOverrideData []DocumentFieldResultDataType `xml:"urn:com.workday/bsvc Integration_Field_Override_Data,omitempty"`
	OffCycleData                 []OffCycleDataType            `xml:"urn:com.workday/bsvc Off_Cycle_Data,omitempty"`
}

func (t *PayrollPayslipDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayslipDataType
	var layout struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PaymentDate = (*xsdDate)(layout.T.PaymentDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayslipDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayslipDataType
	var overlay struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDate = (*xsdDate)(overlay.T.PaymentDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Payslip Request Criteria
type PayrollPayslipRequestCriteriaType struct {
	OutsourcedPaymentGroupReference   *OutsourcedPaymentGroupObjectType  `xml:"urn:com.workday/bsvc Outsourced_Payment_Group_Reference"`
	EmployeeReference                 []WorkerObjectType                 `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	PaymentIntervalStartDateReference *time.Time                         `xml:"urn:com.workday/bsvc Payment_Interval_Start_Date_Reference,omitempty"`
	PaymentIntervalEndDateReference   *time.Time                         `xml:"urn:com.workday/bsvc Payment_Interval_End_Date_Reference,omitempty"`
	PayslipPrintingOptionReference    *PayslipPrintingOptionObjectType   `xml:"urn:com.workday/bsvc Payslip_Printing_Option_Reference,omitempty"`
	FieldAndParameterCriteriaData     *FieldAndParameterCriteriaDataType `xml:"urn:com.workday/bsvc Field_And_Parameter_Criteria_Data,omitempty"`
}

func (t *PayrollPayslipRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPayslipRequestCriteriaType
	var layout struct {
		*T
		PaymentIntervalStartDateReference *xsdDate `xml:"urn:com.workday/bsvc Payment_Interval_Start_Date_Reference,omitempty"`
		PaymentIntervalEndDateReference   *xsdDate `xml:"urn:com.workday/bsvc Payment_Interval_End_Date_Reference,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PaymentIntervalStartDateReference = (*xsdDate)(layout.T.PaymentIntervalStartDateReference)
	layout.PaymentIntervalEndDateReference = (*xsdDate)(layout.T.PaymentIntervalEndDateReference)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPayslipRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPayslipRequestCriteriaType
	var overlay struct {
		*T
		PaymentIntervalStartDateReference *xsdDate `xml:"urn:com.workday/bsvc Payment_Interval_Start_Date_Reference,omitempty"`
		PaymentIntervalEndDateReference   *xsdDate `xml:"urn:com.workday/bsvc Payment_Interval_End_Date_Reference,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentIntervalStartDateReference = (*xsdDate)(overlay.T.PaymentIntervalStartDateReference)
	overlay.PaymentIntervalEndDateReference = (*xsdDate)(overlay.T.PaymentIntervalEndDateReference)
	return d.DecodeElement(&overlay, &start)
}

// Get Payroll Payslips Response
type PayrollPayslipResponseDataType struct {
	Payslip []PayrollPayslipType `xml:"urn:com.workday/bsvc Payslip,omitempty"`
}

// Payroll Payslip Response Group
type PayrollPayslipResponseGroupType struct {
	IncludeOriginatingContactData    *bool `xml:"urn:com.workday/bsvc Include_Originating_Contact_Data,omitempty"`
	IncludeOriginatingBankData       *bool `xml:"urn:com.workday/bsvc Include_Originating_Bank_Data,omitempty"`
	IncludeReceivingPartyContactData *bool `xml:"urn:com.workday/bsvc Include_Receiving_Party_Contact_Data,omitempty"`
	IncludeReceivingPartyBankData    *bool `xml:"urn:com.workday/bsvc Include_Receiving_Party_Bank_Data,omitempty"`
}

// Payroll Payslip
type PayrollPayslipType struct {
	PayrollPayslipReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Payroll_Payslip_Reference,omitempty"`
	PayslipData             *PayrollPayslipDataType     `xml:"urn:com.workday/bsvc Payslip_Data,omitempty"`
}

// Utilize the following element to find Pre-Printed Payslips in the Workday system.
type PayrollPrePrintedPayslipsRequestCriteriaType struct {
	PrePrintedPayslipCreationFromMoment time.Time                        `xml:"urn:com.workday/bsvc Pre-Printed_Payslip_Creation_From_Moment"`
	PrePrintedPayslipCreationToMoment   time.Time                        `xml:"urn:com.workday/bsvc Pre-Printed_Payslip_Creation_To_Moment"`
	PayRunGroupSelectionReference       []PayRunGroupSelectionObjectType `xml:"urn:com.workday/bsvc Pay_Run_Group_Selection_Reference,omitempty"`
}

func (t *PayrollPrePrintedPayslipsRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollPrePrintedPayslipsRequestCriteriaType
	var layout struct {
		*T
		PrePrintedPayslipCreationFromMoment *xsdDateTime `xml:"urn:com.workday/bsvc Pre-Printed_Payslip_Creation_From_Moment"`
		PrePrintedPayslipCreationToMoment   *xsdDateTime `xml:"urn:com.workday/bsvc Pre-Printed_Payslip_Creation_To_Moment"`
	}
	layout.T = (*T)(t)
	layout.PrePrintedPayslipCreationFromMoment = (*xsdDateTime)(&layout.T.PrePrintedPayslipCreationFromMoment)
	layout.PrePrintedPayslipCreationToMoment = (*xsdDateTime)(&layout.T.PrePrintedPayslipCreationToMoment)
	return e.EncodeElement(layout, start)
}
func (t *PayrollPrePrintedPayslipsRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollPrePrintedPayslipsRequestCriteriaType
	var overlay struct {
		*T
		PrePrintedPayslipCreationFromMoment *xsdDateTime `xml:"urn:com.workday/bsvc Pre-Printed_Payslip_Creation_From_Moment"`
		PrePrintedPayslipCreationToMoment   *xsdDateTime `xml:"urn:com.workday/bsvc Pre-Printed_Payslip_Creation_To_Moment"`
	}
	overlay.T = (*T)(t)
	overlay.PrePrintedPayslipCreationFromMoment = (*xsdDateTime)(&overlay.T.PrePrintedPayslipCreationFromMoment)
	overlay.PrePrintedPayslipCreationToMoment = (*xsdDateTime)(&overlay.T.PrePrintedPayslipCreationToMoment)
	return d.DecodeElement(&overlay, &start)
}

// Get Payroll Pre-Printed Payslips Request Reference
type PayrollPrePrintedPayslipsRequestReferenceType struct {
	PayrollPrePrintedPayslipsRequestReference []RepositoryDocumentObjectType `xml:"urn:com.workday/bsvc Payroll_Pre-Printed_Payslips_Request_Reference"`
}

// Response Group
type PayrollPrePrintedPayslipsResponseGroupType struct {
}

// This is the Canadian Payroll ROE Prior History Results Data Element that contains the attributes and instances to be loaded into Workday system for ROE processing.
type PayrollROEPriorHistoryResultsDataType struct {
	EmployeeReference               *EmployeeObjectType               `xml:"urn:com.workday/bsvc Employee_Reference"`
	CompanyReference                *CompanyObjectType                `xml:"urn:com.workday/bsvc Company_Reference"`
	PayrollReferenceNumberReference *PayrollReferenceNumberObjectType `xml:"urn:com.workday/bsvc Payroll_Reference_Number_Reference"`
	ReportingPeriodReference        *PeriodObjectType                 `xml:"urn:com.workday/bsvc Reporting_Period_Reference"`
	EIWages                         float64                           `xml:"urn:com.workday/bsvc EI_Wages,omitempty"`
	EIHours                         float64                           `xml:"urn:com.workday/bsvc EI_Hours,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollReferenceNumberObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollReferenceNumberObjectType struct {
	ID         []PayrollReferenceNumberObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payroll Remittance Absence Plan
type PayrollRemittanceAbsencePlanType struct {
	GBName  string  `xml:"urn:com.workday/bsvc GB_Name,omitempty"`
	Name    string  `xml:"urn:com.workday/bsvc Name,omitempty"`
	Accrued float64 `xml:"urn:com.workday/bsvc Accrued,omitempty"`
	Reduced float64 `xml:"urn:com.workday/bsvc Reduced,omitempty"`
	Balance float64 `xml:"urn:com.workday/bsvc Balance,omitempty"`
}

// Payroll Remittance Data. Will always be included in the Get Payroll Payslips web service.  Can be conditionally returned by setting the "Include Payroll Remittance Data" grouping flag in the Get Payments web service.
type PayrollRemittanceDataType struct {
	PeriodStartDate                      *time.Time                                      `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
	PeriodEndDate                        *time.Time                                      `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	FederalMaritalStatusReference        *PayrollWithholdingStatusObjectType             `xml:"urn:com.workday/bsvc Federal_Marital_Status_Reference,omitempty"`
	FederalAllowance                     float64                                         `xml:"urn:com.workday/bsvc Federal_Allowance,omitempty"`
	FederalAdditionalAllowance           float64                                         `xml:"urn:com.workday/bsvc Federal_Additional_Allowance,omitempty"`
	WorkStateMaritalStatusReference      *PayrollWithholdingStatusObjectType             `xml:"urn:com.workday/bsvc Work_State_Marital_Status_Reference,omitempty"`
	WorkStateAllowance                   float64                                         `xml:"urn:com.workday/bsvc Work_State_Allowance,omitempty"`
	WorkStateAdditionalAmount            float64                                         `xml:"urn:com.workday/bsvc Work_State_Additional_Amount,omitempty"`
	CalculatedFederalTotalClaimAmount    float64                                         `xml:"urn:com.workday/bsvc Calculated_Federal_Total_Claim_Amount,omitempty"`
	CalculatedProvincialTotalClaimAmount float64                                         `xml:"urn:com.workday/bsvc Calculated_Provincial_Total_Claim_Amount,omitempty"`
	FederalTotalClaimAmount              float64                                         `xml:"urn:com.workday/bsvc Federal_Total_Claim_Amount,omitempty"`
	ProvinceTotalClaimAmount             float64                                         `xml:"urn:com.workday/bsvc Province_Total_Claim_Amount,omitempty"`
	FederalLivingPrescribedZone          float64                                         `xml:"urn:com.workday/bsvc Federal_Living_Prescribed_Zone,omitempty"`
	ProvinceLivingPrescribedZone         float64                                         `xml:"urn:com.workday/bsvc Province_Living_Prescribed_Zone,omitempty"`
	FederalAdditionalTaxAmounts          float64                                         `xml:"urn:com.workday/bsvc Federal_Additional_Tax_Amounts,omitempty"`
	ProvinceAdditionalTaxAmounts         float64                                         `xml:"urn:com.workday/bsvc Province_Additional_Tax_Amounts,omitempty"`
	FederalAnnualDeductionsCredits       float64                                         `xml:"urn:com.workday/bsvc Federal_Annual_Deductions_Credits,omitempty"`
	ProvinceAnnualDeductionsCredits      float64                                         `xml:"urn:com.workday/bsvc Province_Annual_Deductions_Credits,omitempty"`
	FederalLabourSponsoredFund           float64                                         `xml:"urn:com.workday/bsvc Federal_Labour_Sponsored_Fund,omitempty"`
	ProvinceLabourSponsoredFund          float64                                         `xml:"urn:com.workday/bsvc Province_Labour_Sponsored_Fund,omitempty"`
	PayGroupDetailReference              *PayGroupDetailObjectType                       `xml:"urn:com.workday/bsvc Pay_Group_Detail_Reference,omitempty"`
	FrequencyName                        string                                          `xml:"urn:com.workday/bsvc Frequency_Name,omitempty"`
	GrossAmount                          float64                                         `xml:"urn:com.workday/bsvc Gross_Amount,omitempty"`
	TotalUnits                           float64                                         `xml:"urn:com.workday/bsvc Total_Units,omitempty"`
	TotalHoursWorked                     float64                                         `xml:"urn:com.workday/bsvc Total_Hours_Worked,omitempty"`
	YTDTotalHoursWorked                  float64                                         `xml:"urn:com.workday/bsvc YTD_Total_Hours_Worked,omitempty"`
	NetPay                               float64                                         `xml:"urn:com.workday/bsvc Net_Pay,omitempty"`
	TaxesDeductions                      float64                                         `xml:"urn:com.workday/bsvc Taxes_Deductions,omitempty"`
	PrintPayslip                         *bool                                           `xml:"urn:com.workday/bsvc Print_Payslip,omitempty"`
	PrintCheckonPayslip                  *bool                                           `xml:"urn:com.workday/bsvc Print_Check_on_Payslip,omitempty"`
	PayslipDistributionSortingCriteria   string                                          `xml:"urn:com.workday/bsvc Payslip_Distribution_Sorting_Criteria,omitempty"`
	OregonBusinessIdentificationNumber   string                                          `xml:"urn:com.workday/bsvc Oregon_Business_Identification_Number,omitempty"`
	PayRateTypeReference                 *PayRateTypeObjectType                          `xml:"urn:com.workday/bsvc Pay_Rate_Type_Reference,omitempty"`
	PayRate                              float64                                         `xml:"urn:com.workday/bsvc Pay_Rate,omitempty"`
	WorkerData                           []PayrollRemittanceWorkerDataType               `xml:"urn:com.workday/bsvc Worker_Data,omitempty"`
	WorkAddressData                      []AddressInformationDataType                    `xml:"urn:com.workday/bsvc Work_Address_Data,omitempty"`
	GrossAndNetData                      []PayrollRemittanceGrossAndNetDataType          `xml:"urn:com.workday/bsvc Gross_And_Net_Data,omitempty"`
	TotalData                            []PayrollRemittanceTotalDataType                `xml:"urn:com.workday/bsvc Total_Data,omitempty"`
	EarningsData                         []PayrollRemittanceEarningsDataType             `xml:"urn:com.workday/bsvc Earnings_Data,omitempty"`
	PreTaxDeductionsData                 []PayrollRemittancePreTaxDeductionDataType      `xml:"urn:com.workday/bsvc Pre_Tax_Deductions_Data,omitempty"`
	PostTaxDeductionsData                []PayrollRemittancePostTaxDataType              `xml:"urn:com.workday/bsvc Post_Tax_Deductions_Data,omitempty"`
	EmployeeTaxesData                    []PayrollRemittanceEmployeeTaxesDataType        `xml:"urn:com.workday/bsvc Employee_Taxes_Data,omitempty"`
	EmployerPaidBenefitsData             []PayrollRemittanceEmployerPaidBenefitsDataType `xml:"urn:com.workday/bsvc Employer_Paid_Benefits_Data,omitempty"`
	TaxableWagesData                     []PayrollRemittanceTaxableWagesDataType         `xml:"urn:com.workday/bsvc Taxable_Wages_Data,omitempty"`
	AbsencePlansData                     []PayrollRemittanceAbsencePlanType              `xml:"urn:com.workday/bsvc Absence_Plans_Data,omitempty"`
	PayslipMessageData                   []PayrollRemittancePayslipMessageDataType       `xml:"urn:com.workday/bsvc Payslip_Message_Data,omitempty"`
}

func (t *PayrollRemittanceDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollRemittanceDataType
	var layout struct {
		*T
		PeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
		PeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PeriodStartDate = (*xsdDate)(layout.T.PeriodStartDate)
	layout.PeriodEndDate = (*xsdDate)(layout.T.PeriodEndDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollRemittanceDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollRemittanceDataType
	var overlay struct {
		*T
		PeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
		PeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PeriodStartDate = (*xsdDate)(overlay.T.PeriodStartDate)
	overlay.PeriodEndDate = (*xsdDate)(overlay.T.PeriodEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll remitance Earnings
type PayrollRemittanceEarningsDataType struct {
	PayComponentReference    []PayComponentReferenceType    `xml:"urn:com.workday/bsvc Pay_Component_Reference,omitempty"`
	GBName                   string                         `xml:"urn:com.workday/bsvc GB_Name,omitempty"`
	PayComponent             string                         `xml:"urn:com.workday/bsvc Pay_Component,omitempty"`
	StartDate                *time.Time                     `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                  *time.Time                     `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	PayrollWorktagsReference []PayrollWorktagObjectType     `xml:"urn:com.workday/bsvc Payroll_Worktags_Reference,omitempty"`
	Amount                   float64                        `xml:"urn:com.workday/bsvc Amount,omitempty"`
	YTDAmount                float64                        `xml:"urn:com.workday/bsvc YTD_Amount,omitempty"`
	UnitsTypeReference       []PayrollCalculationObjectType `xml:"urn:com.workday/bsvc Units_Type_Reference,omitempty"`
	UnitsValue               float64                        `xml:"urn:com.workday/bsvc Units_Value,omitempty"`
	YTDUnitsValue            float64                        `xml:"urn:com.workday/bsvc YTD_Units_Value,omitempty"`
	RateTypeReference        []PayrollCalculationObjectType `xml:"urn:com.workday/bsvc Rate_Type_Reference,omitempty"`
	Rate                     float64                        `xml:"urn:com.workday/bsvc Rate,omitempty"`
}

func (t *PayrollRemittanceEarningsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollRemittanceEarningsDataType
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
func (t *PayrollRemittanceEarningsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollRemittanceEarningsDataType
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

// Payroll Remittance Employee Taxes Data
type PayrollRemittanceEmployeeTaxesDataType struct {
	GBName                   string                      `xml:"urn:com.workday/bsvc GB_Name,omitempty"`
	PayComponentReference    []PayComponentReferenceType `xml:"urn:com.workday/bsvc Pay_Component_Reference,omitempty"`
	PayComponent             string                      `xml:"urn:com.workday/bsvc Pay_Component,omitempty"`
	StartDate                *time.Time                  `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                  *time.Time                  `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	PayrollWorktagsReference []PayrollWorktagObjectType  `xml:"urn:com.workday/bsvc Payroll_Worktags_Reference,omitempty"`
	Amount                   float64                     `xml:"urn:com.workday/bsvc Amount,omitempty"`
	YTD                      float64                     `xml:"urn:com.workday/bsvc YTD,omitempty"`
}

func (t *PayrollRemittanceEmployeeTaxesDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollRemittanceEmployeeTaxesDataType
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
func (t *PayrollRemittanceEmployeeTaxesDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollRemittanceEmployeeTaxesDataType
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

// Employer Paid Benefits
type PayrollRemittanceEmployerPaidBenefitsDataType struct {
	PayComponentReference    []PayComponentReferenceType `xml:"urn:com.workday/bsvc Pay_Component_Reference,omitempty"`
	GBName                   string                      `xml:"urn:com.workday/bsvc GB_Name,omitempty"`
	PayComponent             string                      `xml:"urn:com.workday/bsvc Pay_Component,omitempty"`
	StartDate                *time.Time                  `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                  *time.Time                  `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	PayrollWorktagsReference []PayrollWorktagObjectType  `xml:"urn:com.workday/bsvc Payroll_Worktags_Reference,omitempty"`
	Amount                   float64                     `xml:"urn:com.workday/bsvc Amount,omitempty"`
	YTD                      float64                     `xml:"urn:com.workday/bsvc YTD,omitempty"`
}

func (t *PayrollRemittanceEmployerPaidBenefitsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollRemittanceEmployerPaidBenefitsDataType
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
func (t *PayrollRemittanceEmployerPaidBenefitsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollRemittanceEmployerPaidBenefitsDataType
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

// Payroll Remittance Gross And Net Data
type PayrollRemittanceGrossAndNetDataType struct {
	TypeReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Type_Reference,omitempty"`
	Amount        float64                     `xml:"urn:com.workday/bsvc Amount,omitempty"`
	YTD           float64                     `xml:"urn:com.workday/bsvc YTD,omitempty"`
}

// Payroll Remittance Payslip Message
type PayrollRemittancePayslipMessageDataType struct {
	PayslipMessage string `xml:"urn:com.workday/bsvc Payslip_Message,omitempty"`
}

// Payroll Remittance Post Tax Data
type PayrollRemittancePostTaxDataType struct {
	PayComponentReference    []PayComponentReferenceType `xml:"urn:com.workday/bsvc Pay_Component_Reference,omitempty"`
	GBName                   string                      `xml:"urn:com.workday/bsvc GB_Name,omitempty"`
	PayComponent             string                      `xml:"urn:com.workday/bsvc Pay_Component,omitempty"`
	StartDate                *time.Time                  `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                  *time.Time                  `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	PayrollWorktagsReference []PayrollWorktagObjectType  `xml:"urn:com.workday/bsvc Payroll_Worktags_Reference,omitempty"`
	Amount                   float64                     `xml:"urn:com.workday/bsvc Amount,omitempty"`
	YTD                      float64                     `xml:"urn:com.workday/bsvc YTD,omitempty"`
}

func (t *PayrollRemittancePostTaxDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollRemittancePostTaxDataType
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
func (t *PayrollRemittancePostTaxDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollRemittancePostTaxDataType
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

// Payroll Remittance Pre Tax Deduction Data
type PayrollRemittancePreTaxDeductionDataType struct {
	PayComponentReference    []PayComponentReferenceType `xml:"urn:com.workday/bsvc Pay_Component_Reference,omitempty"`
	GBName                   string                      `xml:"urn:com.workday/bsvc GB_Name,omitempty"`
	PayComponent             string                      `xml:"urn:com.workday/bsvc Pay_Component,omitempty"`
	StartDate                *time.Time                  `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                  *time.Time                  `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	PayrollWorktagsReference []PayrollWorktagObjectType  `xml:"urn:com.workday/bsvc Payroll_Worktags_Reference,omitempty"`
	Amount                   float64                     `xml:"urn:com.workday/bsvc Amount,omitempty"`
	YTD                      float64                     `xml:"urn:com.workday/bsvc YTD,omitempty"`
}

func (t *PayrollRemittancePreTaxDeductionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollRemittancePreTaxDeductionDataType
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
func (t *PayrollRemittancePreTaxDeductionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollRemittancePreTaxDeductionDataType
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

// Payroll Remittance Taxable Wages
type PayrollRemittanceTaxableWagesDataType struct {
	GBName string  `xml:"urn:com.workday/bsvc GB_Name,omitempty"`
	Name   string  `xml:"urn:com.workday/bsvc Name,omitempty"`
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	YTD    float64 `xml:"urn:com.workday/bsvc YTD,omitempty"`
}

// Payroll Remittance Total Data
type PayrollRemittanceTotalDataType struct {
	Label         string  `xml:"urn:com.workday/bsvc Label,omitempty"`
	CurrentPeriod float64 `xml:"urn:com.workday/bsvc Current_Period,omitempty"`
	YTD           float64 `xml:"urn:com.workday/bsvc YTD,omitempty"`
}

// Payroll Remittance Data
type PayrollRemittanceWorkerDataType struct {
	EmployeeID                string                             `xml:"urn:com.workday/bsvc Employee_ID,omitempty"`
	EmployeeName              string                             `xml:"urn:com.workday/bsvc Employee_Name,omitempty"`
	PreferredName             string                             `xml:"urn:com.workday/bsvc Preferred_Name,omitempty"`
	LegalName                 string                             `xml:"urn:com.workday/bsvc Legal_Name,omitempty"`
	LegalFirstName            string                             `xml:"urn:com.workday/bsvc Legal_First_Name,omitempty"`
	LegalLastName             string                             `xml:"urn:com.workday/bsvc Legal_Last_Name,omitempty"`
	LegalMiddleName           string                             `xml:"urn:com.workday/bsvc Legal_Middle_Name,omitempty"`
	DateofBirth               *time.Time                         `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	SSN                       string                             `xml:"urn:com.workday/bsvc SSN,omitempty"`
	Company                   string                             `xml:"urn:com.workday/bsvc Company,omitempty"`
	Occupation                string                             `xml:"urn:com.workday/bsvc Occupation,omitempty"`
	Phone                     string                             `xml:"urn:com.workday/bsvc Phone,omitempty"`
	LocationReference         *LocationObjectType                `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	DepartmentNameReference   *SupervisoryOrganizationObjectType `xml:"urn:com.workday/bsvc Department_Name_Reference,omitempty"`
	SupervisoryOrganizationID string                             `xml:"urn:com.workday/bsvc Supervisory_Organization_ID,omitempty"`
	RegionReference           []RegionObjectType                 `xml:"urn:com.workday/bsvc Region_Reference,omitempty"`
	RegionID                  string                             `xml:"urn:com.workday/bsvc Region_ID,omitempty"`
	CostCenterReference       []CostCenterObjectType             `xml:"urn:com.workday/bsvc Cost_Center_Reference,omitempty"`
	CostCenterID              string                             `xml:"urn:com.workday/bsvc Cost_Center_ID,omitempty"`
	ManagerReference          *WorkerObjectType                  `xml:"urn:com.workday/bsvc Manager_Reference,omitempty"`
	HomeAddressData           []AddressInformationDataType       `xml:"urn:com.workday/bsvc Home_Address_Data,omitempty"`
}

func (t *PayrollRemittanceWorkerDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollRemittanceWorkerDataType
	var layout struct {
		*T
		DateofBirth *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateofBirth = (*xsdDate)(layout.T.DateofBirth)
	return e.EncodeElement(layout, start)
}
func (t *PayrollRemittanceWorkerDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollRemittanceWorkerDataType
	var overlay struct {
		*T
		DateofBirth *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateofBirth = (*xsdDate)(overlay.T.DateofBirth)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PayrollReportingCodeAllObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type PayrollReportingCodeAllObjectType struct {
	ID         []PayrollReportingCodeAllObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing the Worker(s), Company, and Effective Date for the web service request. Either Worker or Company fields need to be populated. If both are supplied, the response will default to returning the workers supplied and Company will be ignored.
type PayrollReportingCodeForWorkersRequestReferencesType struct {
	WorkersReference []WorkerObjectType `xml:"urn:com.workday/bsvc Workers_Reference,omitempty"`
	CompanyReference *CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	EffectiveDate    *time.Time         `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
}

func (t *PayrollReportingCodeForWorkersRequestReferencesType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollReportingCodeForWorkersRequestReferencesType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollReportingCodeForWorkersRequestReferencesType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollReportingCodeForWorkersRequestReferencesType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains Position and associated Payroll Reporting Codes for that Position as of Effective Date. If Company was passed into the web service request, then only Positions within that Company are returned for the Worker.
type PayrollReportingCodesforPositionType struct {
	PositionReference              *PositionObjectType                 `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	PayrollReportingCodesReference []PayrollReportingCodeAllObjectType `xml:"urn:com.workday/bsvc Payroll_Reporting_Codes_Reference,omitempty"`
}

// Element containing one worker and the worker's associated payroll reporting codes as of the effective date.
type PayrollReportingCodesforWorkerType struct {
	WorkerReference                      *WorkerObjectType                      `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	PositionandPayrollReportingCodesData []PayrollReportingCodesforPositionType `xml:"urn:com.workday/bsvc Position_and_Payroll_Reporting_Codes_Data,omitempty"`
}

// Wrapper element for containing multiple Payroll Reporting Codes for Worker elements.
type PayrollReportingCodesforWorkersResponseDataType struct {
	WorkerPayrollCodesData []PayrollReportingCodesforWorkerType `xml:"urn:com.workday/bsvc Worker_Payroll_Codes_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollReportingTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollReportingTypeObjectType struct {
	ID         []PayrollReportingTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payroll Result Data element
type PayrollResultDataType struct {
	PaymentDate             *time.Time                    `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	ResultStartDate         *time.Time                    `xml:"urn:com.workday/bsvc Result_Start_Date,omitempty"`
	ResultEndDate           *time.Time                    `xml:"urn:com.workday/bsvc Result_End_Date,omitempty"`
	ResultCalculationMoment *time.Time                    `xml:"urn:com.workday/bsvc Result_Calculation_Moment,omitempty"`
	CurrencyReference       *CurrencyObjectType           `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	PayGroupReference       *PayGroupObjectType           `xml:"urn:com.workday/bsvc Pay_Group_Reference,omitempty"`
	PeriodReference         *PeriodObjectType             `xml:"urn:com.workday/bsvc Period_Reference"`
	PeriodStartDate         *time.Time                    `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
	PeriodEndDate           *time.Time                    `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	ThirdPartySickPay       *bool                         `xml:"urn:com.workday/bsvc Third_Party_Sick_Pay,omitempty"`
	PayrollResultDetailData []PayrollResultDetailDataType `xml:"urn:com.workday/bsvc Payroll_Result_Detail_Data,omitempty"`
}

func (t *PayrollResultDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollResultDataType
	var layout struct {
		*T
		PaymentDate             *xsdDate     `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
		ResultStartDate         *xsdDate     `xml:"urn:com.workday/bsvc Result_Start_Date,omitempty"`
		ResultEndDate           *xsdDate     `xml:"urn:com.workday/bsvc Result_End_Date,omitempty"`
		ResultCalculationMoment *xsdDateTime `xml:"urn:com.workday/bsvc Result_Calculation_Moment,omitempty"`
		PeriodStartDate         *xsdDate     `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
		PeriodEndDate           *xsdDate     `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PaymentDate = (*xsdDate)(layout.T.PaymentDate)
	layout.ResultStartDate = (*xsdDate)(layout.T.ResultStartDate)
	layout.ResultEndDate = (*xsdDate)(layout.T.ResultEndDate)
	layout.ResultCalculationMoment = (*xsdDateTime)(layout.T.ResultCalculationMoment)
	layout.PeriodStartDate = (*xsdDate)(layout.T.PeriodStartDate)
	layout.PeriodEndDate = (*xsdDate)(layout.T.PeriodEndDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollResultDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollResultDataType
	var overlay struct {
		*T
		PaymentDate             *xsdDate     `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
		ResultStartDate         *xsdDate     `xml:"urn:com.workday/bsvc Result_Start_Date,omitempty"`
		ResultEndDate           *xsdDate     `xml:"urn:com.workday/bsvc Result_End_Date,omitempty"`
		ResultCalculationMoment *xsdDateTime `xml:"urn:com.workday/bsvc Result_Calculation_Moment,omitempty"`
		PeriodStartDate         *xsdDate     `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
		PeriodEndDate           *xsdDate     `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDate = (*xsdDate)(overlay.T.PaymentDate)
	overlay.ResultStartDate = (*xsdDate)(overlay.T.ResultStartDate)
	overlay.ResultEndDate = (*xsdDate)(overlay.T.ResultEndDate)
	overlay.ResultCalculationMoment = (*xsdDateTime)(overlay.T.ResultCalculationMoment)
	overlay.PeriodStartDate = (*xsdDate)(overlay.T.PeriodStartDate)
	overlay.PeriodEndDate = (*xsdDate)(overlay.T.PeriodEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Result Detail Data element
type PayrollResultDetailDataType struct {
	EmployerPaid                 *bool                                  `xml:"urn:com.workday/bsvc Employer_Paid,omitempty"`
	DeductionReference           *PayrollCalculationObjectType          `xml:"urn:com.workday/bsvc Deduction_Reference"`
	PayrollTaxAuthorityReference *PayrollTaxAuthorityObjectType         `xml:"urn:com.workday/bsvc Payroll_Tax_Authority_Reference"`
	WorkPSDCode                  string                                 `xml:"urn:com.workday/bsvc Work_PSD_Code,omitempty"`
	ResidentPSDCode              string                                 `xml:"urn:com.workday/bsvc Resident_PSD_Code,omitempty"`
	TaxFilingCode                string                                 `xml:"urn:com.workday/bsvc Tax_Filing_Code,omitempty"`
	SUIRate                      float64                                `xml:"urn:com.workday/bsvc SUI_Rate,omitempty"`
	TaxWithheld                  float64                                `xml:"urn:com.workday/bsvc Tax_Withheld,omitempty"`
	TaxableWages                 float64                                `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages                 float64                                `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	GrossWages                   float64                                `xml:"urn:com.workday/bsvc Gross_Wages,omitempty"`
	TipWages                     float64                                `xml:"urn:com.workday/bsvc Tip_Wages,omitempty"`
	QTDData                      *QTDTaxFilingDataforPeriodicWorkerType `xml:"urn:com.workday/bsvc QTD_Data,omitempty"`
	YTDData                      *YTDTaxFilingDataforPeriodicWorkerType `xml:"urn:com.workday/bsvc YTD_Data,omitempty"`
}

// Payroll Result Request Criteria
type PayrollResultRequestCriteriaType struct {
	StartDate                    time.Time                      `xml:"urn:com.workday/bsvc Start_Date"`
	EndDate                      time.Time                      `xml:"urn:com.workday/bsvc End_Date"`
	PeriodSelectionDateReference *PeriodDateIndicatorObjectType `xml:"urn:com.workday/bsvc Period_Selection_Date_Reference"`
	EmployeeReference            []EmployeeObjectType           `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	CompanyReference             []CompanyObjectType            `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	PayCalculationCriteria       *PayCalculationsSelectedType   `xml:"urn:com.workday/bsvc Pay_Calculation_Criteria,omitempty"`
	RunCategoryReference         []RunCategoryObjectType        `xml:"urn:com.workday/bsvc Run_Category_Reference,omitempty"`
	PayGroupReference            []PayGroupObjectType           `xml:"urn:com.workday/bsvc Pay_Group_Reference,omitempty"`
	PeriodReference              []PeriodObjectType             `xml:"urn:com.workday/bsvc Period_Reference,omitempty"`
	SettlementIDReference        []UniqueIdentifierObjectType   `xml:"urn:com.workday/bsvc Settlement_ID_Reference,omitempty"`
}

func (t *PayrollResultRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollResultRequestCriteriaType
	var layout struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(&layout.T.StartDate)
	layout.EndDate = (*xsdDate)(&layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollResultRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollResultRequestCriteriaType
	var overlay struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(&overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(&overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Result Request References
type PayrollResultRequestReferencesType struct {
	PayrollResultReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Payroll_Result_Reference"`
}

// Payroll Result Response Data
type PayrollResultResponseDataType struct {
	PayrollResult []PayrollResultType `xml:"urn:com.workday/bsvc Payroll_Result,omitempty"`
}

// Payroll Result Response Group
type PayrollResultResponseGroupType struct {
	IncludeNameData                     *bool `xml:"urn:com.workday/bsvc Include_Name_Data,omitempty"`
	IncludeNationalIDData               *bool `xml:"urn:com.workday/bsvc Include_National_ID_Data,omitempty"`
	IncludeRelatedCalculationResultData *bool `xml:"urn:com.workday/bsvc Include_Related_Calculation_Result_Data,omitempty"`
	IncludeWithholdingOrderData         *bool `xml:"urn:com.workday/bsvc Include_Withholding_Order_Data,omitempty"`
	IncludePayrollWorktagData           *bool `xml:"urn:com.workday/bsvc Include_Payroll_Worktag_Data,omitempty"`
	IncludeQTDData                      *bool `xml:"urn:com.workday/bsvc Include_QTD_Data,omitempty"`
	IncludeYTDData                      *bool `xml:"urn:com.workday/bsvc Include_YTD_Data,omitempty"`
}

// Payroll Result
type PayrollResultType struct {
	EmployeeReference                  *EmployeeObjectType                           `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	EmployeeNameData                   *NameDataforGetPayrollResultsType             `xml:"urn:com.workday/bsvc Employee_Name_Data,omitempty"`
	NationalIDData                     *NationalIDDataforGetPayrollResultsType       `xml:"urn:com.workday/bsvc National_ID_Data,omitempty"`
	CompanyReference                   *CompanyObjectType                            `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	PayrollIDReference                 *PayrollIDObjectType                          `xml:"urn:com.workday/bsvc Payroll_ID_Reference,omitempty"`
	PayrollAccountNumber               string                                        `xml:"urn:com.workday/bsvc Payroll_Account_Number,omitempty"`
	EstablishmentReference             *EstablishmentObjectType                      `xml:"urn:com.workday/bsvc Establishment_Reference,omitempty"`
	PayGroupReference                  *PayGroupDetailObjectType                     `xml:"urn:com.workday/bsvc Pay_Group_Reference,omitempty"`
	RunCategoryReference               *RunCategoryObjectType                        `xml:"urn:com.workday/bsvc Run_Category_Reference,omitempty"`
	PayPeriodFrequency                 float64                                       `xml:"urn:com.workday/bsvc Pay_Period_Frequency,omitempty"`
	OffcycleTypeReference              *PayrollOffcycleTypeObjectType                `xml:"urn:com.workday/bsvc Off-cycle_Type_Reference,omitempty"`
	ForAdditionalPay                   *bool                                         `xml:"urn:com.workday/bsvc For_Additional_Pay,omitempty"`
	ResultCompletedDateTime            *time.Time                                    `xml:"urn:com.workday/bsvc Result_Completed_Date_Time,omitempty"`
	PeriodReference                    *PeriodObjectType                             `xml:"urn:com.workday/bsvc Period_Reference,omitempty"`
	PeriodStartDate                    *time.Time                                    `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
	PeriodEndDate                      *time.Time                                    `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	SubperiodStartDate                 *time.Time                                    `xml:"urn:com.workday/bsvc Subperiod_Start_Date,omitempty"`
	SubperiodEndDate                   *time.Time                                    `xml:"urn:com.workday/bsvc Subperiod_End_Date,omitempty"`
	PaymentDate                        *time.Time                                    `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	GrossAmount                        float64                                       `xml:"urn:com.workday/bsvc Gross_Amount,omitempty"`
	NetAmount                          float64                                       `xml:"urn:com.workday/bsvc Net_Amount,omitempty"`
	CurrencyReference                  *CurrencyReferenceDataType                    `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	PayrollResultData                  []ResultLineDetailsforGetPayrollResultsType   `xml:"urn:com.workday/bsvc Payroll_Result_Data,omitempty"`
	PayAccumulationData                []PayAccumulationDataforGetPayrollResultsType `xml:"urn:com.workday/bsvc Pay_Accumulation_Data,omitempty"`
	PaymentDateofOriginalPayrollResult *time.Time                                    `xml:"urn:com.workday/bsvc Payment_Date_of_Original_Payroll_Result,omitempty"`
}

func (t *PayrollResultType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollResultType
	var layout struct {
		*T
		ResultCompletedDateTime            *xsdDateTime `xml:"urn:com.workday/bsvc Result_Completed_Date_Time,omitempty"`
		PeriodStartDate                    *xsdDate     `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
		PeriodEndDate                      *xsdDate     `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
		SubperiodStartDate                 *xsdDate     `xml:"urn:com.workday/bsvc Subperiod_Start_Date,omitempty"`
		SubperiodEndDate                   *xsdDate     `xml:"urn:com.workday/bsvc Subperiod_End_Date,omitempty"`
		PaymentDate                        *xsdDate     `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
		PaymentDateofOriginalPayrollResult *xsdDate     `xml:"urn:com.workday/bsvc Payment_Date_of_Original_Payroll_Result,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ResultCompletedDateTime = (*xsdDateTime)(layout.T.ResultCompletedDateTime)
	layout.PeriodStartDate = (*xsdDate)(layout.T.PeriodStartDate)
	layout.PeriodEndDate = (*xsdDate)(layout.T.PeriodEndDate)
	layout.SubperiodStartDate = (*xsdDate)(layout.T.SubperiodStartDate)
	layout.SubperiodEndDate = (*xsdDate)(layout.T.SubperiodEndDate)
	layout.PaymentDate = (*xsdDate)(layout.T.PaymentDate)
	layout.PaymentDateofOriginalPayrollResult = (*xsdDate)(layout.T.PaymentDateofOriginalPayrollResult)
	return e.EncodeElement(layout, start)
}
func (t *PayrollResultType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollResultType
	var overlay struct {
		*T
		ResultCompletedDateTime            *xsdDateTime `xml:"urn:com.workday/bsvc Result_Completed_Date_Time,omitempty"`
		PeriodStartDate                    *xsdDate     `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
		PeriodEndDate                      *xsdDate     `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
		SubperiodStartDate                 *xsdDate     `xml:"urn:com.workday/bsvc Subperiod_Start_Date,omitempty"`
		SubperiodEndDate                   *xsdDate     `xml:"urn:com.workday/bsvc Subperiod_End_Date,omitempty"`
		PaymentDate                        *xsdDate     `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
		PaymentDateofOriginalPayrollResult *xsdDate     `xml:"urn:com.workday/bsvc Payment_Date_of_Original_Payroll_Result,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ResultCompletedDateTime = (*xsdDateTime)(overlay.T.ResultCompletedDateTime)
	overlay.PeriodStartDate = (*xsdDate)(overlay.T.PeriodStartDate)
	overlay.PeriodEndDate = (*xsdDate)(overlay.T.PeriodEndDate)
	overlay.SubperiodStartDate = (*xsdDate)(overlay.T.SubperiodStartDate)
	overlay.SubperiodEndDate = (*xsdDate)(overlay.T.SubperiodEndDate)
	overlay.PaymentDate = (*xsdDate)(overlay.T.PaymentDate)
	overlay.PaymentDateofOriginalPayrollResult = (*xsdDate)(overlay.T.PaymentDateofOriginalPayrollResult)
	return d.DecodeElement(&overlay, &start)
}

// Contains Pay Groups being calculated as part of running a payroll calculation
type PayrollRunDetailPayCalcStatusType struct {
	PayGroupReference    *PayGroupObjectType    `xml:"urn:com.workday/bsvc Pay_Group_Reference,omitempty"`
	RunCategoryReference *RunCategoryObjectType `xml:"urn:com.workday/bsvc Run_Category_Reference,omitempty"`
	CalcStatus           []CalcStatusType       `xml:"urn:com.workday/bsvc Calc_Status,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollStateAuthorityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollStateAuthorityObjectType struct {
	ID         []PayrollStateAuthorityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Response for Get Payroll USA State and Local Tax Elections
type PayrollStateandLocalTaxElectionResponseDataType struct {
	PayrollStateandLocalTaxElection []PayrollPayeeStateandLocalTaxElectionsType `xml:"urn:com.workday/bsvc Payroll_State_and_Local_Tax_Election,omitempty"`
}

// Include this element to Amend or Terminate an existing Support Order.
type PayrollSupportOrderAmendorTerminateDataType struct {
	SupportOrderReference *SupportOrderObjectType `xml:"urn:com.workday/bsvc Support_Order_Reference,omitempty"`
}

// Payroll Support Order Data
type PayrollSupportOrderDataType struct {
	OriginalOrder                       *bool                               `xml:"urn:com.workday/bsvc Original_Order,omitempty"`
	AmendedOrder                        *bool                               `xml:"urn:com.workday/bsvc Amended_Order,omitempty"`
	TerminationOrder                    *bool                               `xml:"urn:com.workday/bsvc Termination_Order,omitempty"`
	CustodialPartyName                  string                              `xml:"urn:com.workday/bsvc Custodial_Party_Name,omitempty"`
	SupportsSecondFamily                *bool                               `xml:"urn:com.workday/bsvc Supports_Second_Family,omitempty"`
	RemittanceIDOverride                string                              `xml:"urn:com.workday/bsvc Remittance_ID_Override,omitempty"`
	SupportOrderDependantReference      []SupportOrderDependentDataType     `xml:"urn:com.workday/bsvc Support_Order_Dependant_Reference,omitempty"`
	SupportOrderDetailsReference        []PayrollSupportOrderDetailDataType `xml:"urn:com.workday/bsvc Support_Order_Details_Reference,omitempty"`
	PayrollLocalCountyAuthorityFIPSCode string                              `xml:"urn:com.workday/bsvc Payroll_Local_County_Authority_FIPS_Code,omitempty"`
}

// Payroll Support Order Detail Data
type PayrollSupportOrderDetailDataType struct {
	OrderFormAmount         float64                `xml:"urn:com.workday/bsvc Order_Form_Amount,omitempty"`
	PayPeriodAmountOverride float64                `xml:"urn:com.workday/bsvc Pay_Period_Amount_Override,omitempty"`
	AmountasPercent         float64                `xml:"urn:com.workday/bsvc Amount_as_Percent,omitempty"`
	SupportTypeReference    *SupportTypeObjectType `xml:"urn:com.workday/bsvc Support_Type_Reference"`
	ArrearsOver12Weeks      *bool                  `xml:"urn:com.workday/bsvc Arrears_Over_12_Weeks,omitempty"`
}

// Payroll Support Order Lump Sum Data
type PayrollSupportOrderLumpSumDataType struct {
	ProcessinRegularPeriod               *bool                                  `xml:"urn:com.workday/bsvc Process_in_Regular_Period,omitempty"`
	CustodialPartyName                   string                                 `xml:"urn:com.workday/bsvc Custodial_Party_Name,omitempty"`
	RemittanceIDOverride                 string                                 `xml:"urn:com.workday/bsvc Remittance_ID_Override,omitempty"`
	SupportsSecondFamily                 *bool                                  `xml:"urn:com.workday/bsvc Supports_Second_Family,omitempty"`
	PayrollLocalCountyAuthorityReference *PayrollLocalCountyAuthorityObjectType `xml:"urn:com.workday/bsvc Payroll_Local_County_Authority_Reference,omitempty"`
	DependantsReference                  []SupportOrderDependentDataType        `xml:"urn:com.workday/bsvc Dependants_Reference,omitempty"`
	SupportOrderDetailReference          *PayrollSupportOrderDetailDataType     `xml:"urn:com.workday/bsvc Support_Order_Detail_Reference,omitempty"`
	PayrollLocalCountyAuthorityFIPSCode  string                                 `xml:"urn:com.workday/bsvc Payroll_Local_County_Authority_FIPS_Code,omitempty"`
}

// Quarterly Worker Tax Filing Data
type PayrollTaxAuthorityDataType struct {
	PayrollTaxAuthorityReference *PayrollTaxAuthorityObjectType `xml:"urn:com.workday/bsvc Payroll_Tax_Authority_Reference,omitempty"`
	Attribute                    string                         `xml:"urn:com.workday/bsvc Attribute,omitempty"`
	Value                        string                         `xml:"urn:com.workday/bsvc Value,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollTaxAuthorityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollTaxAuthorityObjectType struct {
	ID         []PayrollTaxAuthorityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element to hold Payroll Tax Location Mapping Data information
type PayrollTaxLocationMappingDataType struct {
	LocationReference                    *LocationObjectType                    `xml:"urn:com.workday/bsvc Location_Reference"`
	EffectiveDate                        time.Time                              `xml:"urn:com.workday/bsvc Effective_Date"`
	Inactive                             *bool                                  `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	PayrollStateAuthorityReference       *PayrollStateAuthorityObjectType       `xml:"urn:com.workday/bsvc Payroll_State_Authority_Reference,omitempty"`
	PayrollLocalCountyAuthorityReference *PayrollLocalCountyAuthorityObjectType `xml:"urn:com.workday/bsvc Payroll_Local_County_Authority_Reference,omitempty"`
	PayrollLocalCityAuthorityReference   *PayrollLocalCityAuthorityObjectType   `xml:"urn:com.workday/bsvc Payroll_Local_City_Authority_Reference,omitempty"`
	PayrollOtherAuthorityReference       *PayrollOtherAuthorityObjectType       `xml:"urn:com.workday/bsvc Payroll_Other_Authority_Reference,omitempty"`
}

func (t *PayrollTaxLocationMappingDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollTaxLocationMappingDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollTaxLocationMappingDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollTaxLocationMappingDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PayrollTaxLocationMappingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollTaxLocationMappingObjectType struct {
	ID         []PayrollTaxLocationMappingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element to hold Payroll Tax Mappings on Location Response Data Information
type PayrollTaxLocationMappingType struct {
	PayrollTaxLocationMappingReference *PayrollTaxLocationMappingObjectType `xml:"urn:com.workday/bsvc Payroll_Tax_Location_Mapping_Reference,omitempty"`
	PayrollTaxLocationMappingData      *PayrollTaxLocationMappingDataType   `xml:"urn:com.workday/bsvc Payroll_Tax_Location_Mapping_Data,omitempty"`
}

// Documentation Element to hold Payroll Tax Mapping on Location Request Criteria Information
type PayrollTaxMappingsonLocationRequestCriteriaType struct {
	LocationReference      []LocationObjectType `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	EffectiveDateReference *time.Time           `xml:"urn:com.workday/bsvc Effective_Date_Reference,omitempty"`
}

func (t *PayrollTaxMappingsonLocationRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollTaxMappingsonLocationRequestCriteriaType
	var layout struct {
		*T
		EffectiveDateReference *xsdDate `xml:"urn:com.workday/bsvc Effective_Date_Reference,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDateReference = (*xsdDate)(layout.T.EffectiveDateReference)
	return e.EncodeElement(layout, start)
}
func (t *PayrollTaxMappingsonLocationRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollTaxMappingsonLocationRequestCriteriaType
	var overlay struct {
		*T
		EffectiveDateReference *xsdDate `xml:"urn:com.workday/bsvc Effective_Date_Reference,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDateReference = (*xsdDate)(overlay.T.EffectiveDateReference)
	return d.DecodeElement(&overlay, &start)
}

// Documentation Element to hold Payroll Tax Mappings on Location Request References information
type PayrollTaxMappingsonLocationRequestReferencesType struct {
	PayrollTaxLocationMappingReference []PayrollTaxLocationMappingObjectType `xml:"urn:com.workday/bsvc Payroll_Tax_Location_Mapping_Reference,omitempty"`
}

// Element to hold Payroll Tax Mappings on Location Response Data information
type PayrollTaxMappingsonLocationResponseDataType struct {
	PayrollTaxLocationMapping []PayrollTaxLocationMappingType `xml:"urn:com.workday/bsvc Payroll_Tax_Location_Mapping,omitempty"`
}

// Request Criteria used to filter State and Local Tax Elections.
type PayrollUSAStateandLocalTaxElectionRequestCriteriaType struct {
	CompanyReference *CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	WorkerReference  []WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	EffectiveDate    *time.Time         `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
}

func (t *PayrollUSAStateandLocalTaxElectionRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollUSAStateandLocalTaxElectionRequestCriteriaType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollUSAStateandLocalTaxElectionRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollUSAStateandLocalTaxElectionRequestCriteriaType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Response Group for Payroll USA State and Local Tax Elections
type PayrollUSAStateandLocalTaxElectionResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollWithholdingStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollWithholdingStatusObjectType struct {
	ID         []PayrollWithholdingStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payroll Worktag Data
type PayrollWorktagDataType struct {
	CompanyWorktagReference                     []CompanyObjectType                             `xml:"urn:com.workday/bsvc Company_Worktag_Reference,omitempty"`
	LocationWorktagReference                    []LocationObjectType                            `xml:"urn:com.workday/bsvc Location_Worktag_Reference,omitempty"`
	RegionWorktagReference                      []RegionObjectType                              `xml:"urn:com.workday/bsvc Region_Worktag_Reference,omitempty"`
	JobProfileWorktagReference                  []JobProfileObjectType                          `xml:"urn:com.workday/bsvc Job_Profile_Worktag_Reference,omitempty"`
	CostCenterWorktagReference                  []CostCenterObjectType                          `xml:"urn:com.workday/bsvc Cost_Center_Worktag_Reference,omitempty"`
	ProjectWorktagReference                     []ProjectObjectType                             `xml:"urn:com.workday/bsvc Project_Worktag_Reference,omitempty"`
	WithholdingOrderWorktagReference            []WithholdingOrderCaseObjectType                `xml:"urn:com.workday/bsvc Withholding_Order_Worktag_Reference,omitempty"`
	PayrollStateAuthorityWorktagReference       []PayrollStateAuthorityObjectType               `xml:"urn:com.workday/bsvc Payroll_State_Authority_Worktag_Reference,omitempty"`
	WorkersCompensationCodeReference            *WorkersCompensationCodeObjectType              `xml:"urn:com.workday/bsvc Workers_Compensation_Code_Reference,omitempty"`
	PayrollLocalCountyAuthorityWorktagReference []PayrollLocalCountyAuthorityObjectType         `xml:"urn:com.workday/bsvc Payroll_Local_County_Authority_Worktag_Reference,omitempty"`
	PayrollLocalCityAuthorityWorktagReference   []PayrollLocalCityAuthorityObjectType           `xml:"urn:com.workday/bsvc Payroll_Local_City_Authority_Worktag_Reference,omitempty"`
	PayrollLocalSchoolDistrictWorktagReference  []PayrollLocalSchoolDistrictAuthorityObjectType `xml:"urn:com.workday/bsvc Payroll_Local_School_District_Worktag_Reference,omitempty"`
	PayrollLocalOtherWorktagReference           []PayrollOtherAuthorityObjectType               `xml:"urn:com.workday/bsvc Payroll_Local_Other_Worktag_Reference,omitempty"`
	CustomWorktag1Reference                     []CustomWorktag01ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_1_Reference,omitempty"`
	CustomWorktag2Reference                     []CustomWorktag02ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_2_Reference,omitempty"`
	CustomWorktag3Reference                     []CustomWorktag03ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_3_Reference,omitempty"`
	CustomWorktag4Reference                     []CustomWorktag04ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_4_Reference,omitempty"`
	CustomWorktag5Reference                     []CustomWorktag05ObjectType                     `xml:"urn:com.workday/bsvc Custom_Worktag_5_Reference,omitempty"`
	FundWorktagReference                        *FundObjectType                                 `xml:"urn:com.workday/bsvc Fund_Worktag_Reference,omitempty"`
	GrantWorktagReference                       *GrantObjectType                                `xml:"urn:com.workday/bsvc Grant_Worktag_Reference,omitempty"`
	GiftWorktagReference                        *GiftObjectType                                 `xml:"urn:com.workday/bsvc Gift_Worktag_Reference,omitempty"`
	ProgramWorktagReference                     *ProgramObjectType                              `xml:"urn:com.workday/bsvc Program_Worktag_Reference,omitempty"`
	BusinessUnitWorktagReference                *BusinessUnitObjectType                         `xml:"urn:com.workday/bsvc Business_Unit_Worktag_Reference,omitempty"`
	ObjectClassWorktagReference                 *ObjectClassObjectType                          `xml:"urn:com.workday/bsvc Object_Class_Worktag_Reference,omitempty"`
	ProjectPhaseWorktagReference                *ProjectPlanPhaseObjectType                     `xml:"urn:com.workday/bsvc Project_Phase_Worktag_Reference,omitempty"`
	ProjectTaskWorktagReference                 *ProjectPlanTaskObjectType                      `xml:"urn:com.workday/bsvc Project_Task_Worktag_Reference,omitempty"`
	CustomOrganizationWorktagData               []CustomOrganizationWorktagDataType             `xml:"urn:com.workday/bsvc Custom_Organization_Worktag_Data,omitempty"`
	CustomWorktag06Reference                    *CustomWorktag06ObjectType                      `xml:"urn:com.workday/bsvc Custom_Worktag_06_Reference,omitempty"`
	CustomWorktag07Reference                    *CustomWorktag07ObjectType                      `xml:"urn:com.workday/bsvc Custom_Worktag_07_Reference,omitempty"`
	CustomWorktag08Reference                    *CustomWorktag08ObjectType                      `xml:"urn:com.workday/bsvc Custom_Worktag_08_Reference,omitempty"`
	CustomWorktag09Reference                    *CustomWorktag09ObjectType                      `xml:"urn:com.workday/bsvc Custom_Worktag_09_Reference,omitempty"`
	CustomWorktag10Reference                    *CustomWorktag10ObjectType                      `xml:"urn:com.workday/bsvc Custom_Worktag_10_Reference,omitempty"`
	CustomWorktag11Reference                    *CustomWorktag11ObjectType                      `xml:"urn:com.workday/bsvc Custom_Worktag_11_Reference,omitempty"`
	CustomWorktag12Reference                    *CustomWorktag12ObjectType                      `xml:"urn:com.workday/bsvc Custom_Worktag_12_Reference,omitempty"`
	CustomWorktag13Reference                    *CustomWorktag13ObjectType                      `xml:"urn:com.workday/bsvc Custom_Worktag_13_Reference,omitempty"`
	CustomWorktag14Reference                    *CustomWorktag14ObjectType                      `xml:"urn:com.workday/bsvc Custom_Worktag_14_Reference,omitempty"`
	CustomWorktag15Reference                    *CustomWorktag15ObjectType                      `xml:"urn:com.workday/bsvc Custom_Worktag_15_Reference,omitempty"`
	NICategoryReference                         *NICategoryObjectType                           `xml:"urn:com.workday/bsvc NI_Category_Reference,omitempty"`
	ARRCOAGIRCCategoryReference                 *ARRCOAGIRCRubricValueObjectType                `xml:"urn:com.workday/bsvc ARRCO-AGIRC_Category_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollWorktagObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type PayrollWorktagObjectType struct {
	ID         []PayrollWorktagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payslip Payment Data
type PayslipPaymentDataType struct {
	PayrollPaymentDisplayOrder float64                          `xml:"urn:com.workday/bsvc Payroll_Payment_Display_Order,omitempty"`
	PaymentCategoryReference   *PaymentCategoryObjectType       `xml:"urn:com.workday/bsvc Payment_Category_Reference,omitempty"`
	PaymentTypeReference       *PaymentTypeObjectType           `xml:"urn:com.workday/bsvc Payment_Type_Reference,omitempty"`
	PaymentMethodReference     *PaymentMethodObjectType         `xml:"urn:com.workday/bsvc Payment_Method_Reference,omitempty"`
	BankName                   string                           `xml:"urn:com.workday/bsvc Bank_Name,omitempty"`
	BankAccountReference       *SettlementInstructionObjectType `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	AccountNumber              string                           `xml:"urn:com.workday/bsvc Account_Number,omitempty"`
	PaymentAmount              float64                          `xml:"urn:com.workday/bsvc Payment_Amount,omitempty"`
	PaymentCurrencyReference   *CurrencyObjectType              `xml:"urn:com.workday/bsvc Payment_Currency_Reference,omitempty"`
	AmountasText               string                           `xml:"urn:com.workday/bsvc Amount_as_Text,omitempty"`
	AmountinPayGroupCurrency   float64                          `xml:"urn:com.workday/bsvc Amount_in_Pay_Group_Currency,omitempty"`
	PayGroupCurrencyReference  *CurrencyObjectType              `xml:"urn:com.workday/bsvc Pay_Group_Currency_Reference,omitempty"`
	CheckNumber                string                           `xml:"urn:com.workday/bsvc Check_Number,omitempty"`
	PaymentReferenceNumber     string                           `xml:"urn:com.workday/bsvc Payment_Reference_Number,omitempty"`
	PaymentMemo                string                           `xml:"urn:com.workday/bsvc Payment_Memo,omitempty"`
	CompanyReference           *CompanyObjectType               `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	OriginatingContactData     []OriginatingPartyWWSDataType    `xml:"urn:com.workday/bsvc Originating_Contact_Data,omitempty"`
	OriginatingBankData        []OriginatingPartyBankDataType   `xml:"urn:com.workday/bsvc Originating_Bank_Data,omitempty"`
	ReceivingPartyReference    *PayeeObjectType                 `xml:"urn:com.workday/bsvc Receiving_Party_Reference,omitempty"`
	ReceivingPartyContactData  []ReceivingPartyWWSDataType      `xml:"urn:com.workday/bsvc Receiving_Party_Contact_Data,omitempty"`
	ReceivingPartyBankData     []ReceivingPartyBankDataType     `xml:"urn:com.workday/bsvc Receiving_Party_Bank_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayslipPrintingOptionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayslipPrintingOptionObjectType struct {
	ID         []PayslipPrintingOptionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayslipPrintingOverrideObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayslipPrintingOverrideObjectType struct {
	ID         []PayslipPrintingOverrideObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Run Category Data
type PayslipRunCategoryDataType struct {
	RunCategoryReference *RunCategoryObjectType `xml:"urn:com.workday/bsvc Run_Category_Reference,omitempty"`
	RegularRunCategory   *bool                  `xml:"urn:com.workday/bsvc Regular_Run_Category,omitempty"`
}

// Peridoic Worker Tax Filing Data
type PeridoicWorkerTaxFilingDataType struct {
	WorkerReference             *WorkerObjectType       `xml:"urn:com.workday/bsvc Worker_Reference"`
	CompanyReference            *CompanyObjectType      `xml:"urn:com.workday/bsvc Company_Reference"`
	ThirdPartyProvidesSickPayW2 *bool                   `xml:"urn:com.workday/bsvc Third_Party_Provides_Sick_Pay_W-2,omitempty"`
	PayrollResultData           []PayrollResultDataType `xml:"urn:com.workday/bsvc Payroll_Result_Data,omitempty"`
}

// Period Content Data
type PeriodDataType struct {
	PeriodStartDate           time.Time                       `xml:"urn:com.workday/bsvc Period_Start_Date"`
	PeriodEndDate             time.Time                       `xml:"urn:com.workday/bsvc Period_End_Date"`
	PayrollPaymentDate        *time.Time                      `xml:"urn:com.workday/bsvc Payroll_Payment_Date,omitempty"`
	PayrollGLAccruals         *bool                           `xml:"urn:com.workday/bsvc Payroll_GL_Accruals,omitempty"`
	DaystoAccrue              float64                         `xml:"urn:com.workday/bsvc Days_to_Accrue,omitempty"`
	DaysinBasis               float64                         `xml:"urn:com.workday/bsvc Days_in_Basis,omitempty"`
	AccrualDate               *time.Time                      `xml:"urn:com.workday/bsvc Accrual_Date,omitempty"`
	AccrualReversalDate       *time.Time                      `xml:"urn:com.workday/bsvc Accrual_Reversal_Date,omitempty"`
	OpenTimeEntry             *time.Time                      `xml:"urn:com.workday/bsvc Open_Time_Entry,omitempty"`
	LockTimeEntry             *time.Time                      `xml:"urn:com.workday/bsvc Lock_Time_Entry,omitempty"`
	UnlockforAdjustments      *time.Time                      `xml:"urn:com.workday/bsvc Unlock_for_Adjustments,omitempty"`
	CloseTimeEntry            *time.Time                      `xml:"urn:com.workday/bsvc Close_Time_Entry,omitempty"`
	PayrollPeriodReference    *PeriodObjectType               `xml:"urn:com.workday/bsvc Payroll_Period_Reference,omitempty"`
	ValuationDateOverrideData []ValuationDateOverrideDataType `xml:"urn:com.workday/bsvc Valuation_Date_Override_Data,omitempty"`
	PeriodID                  string                          `xml:"urn:com.workday/bsvc Period_ID,omitempty"`
}

func (t *PeriodDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PeriodDataType
	var layout struct {
		*T
		PeriodStartDate      *xsdDate     `xml:"urn:com.workday/bsvc Period_Start_Date"`
		PeriodEndDate        *xsdDate     `xml:"urn:com.workday/bsvc Period_End_Date"`
		PayrollPaymentDate   *xsdDate     `xml:"urn:com.workday/bsvc Payroll_Payment_Date,omitempty"`
		AccrualDate          *xsdDate     `xml:"urn:com.workday/bsvc Accrual_Date,omitempty"`
		AccrualReversalDate  *xsdDate     `xml:"urn:com.workday/bsvc Accrual_Reversal_Date,omitempty"`
		OpenTimeEntry        *xsdDateTime `xml:"urn:com.workday/bsvc Open_Time_Entry,omitempty"`
		LockTimeEntry        *xsdDateTime `xml:"urn:com.workday/bsvc Lock_Time_Entry,omitempty"`
		UnlockforAdjustments *xsdDateTime `xml:"urn:com.workday/bsvc Unlock_for_Adjustments,omitempty"`
		CloseTimeEntry       *xsdDateTime `xml:"urn:com.workday/bsvc Close_Time_Entry,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PeriodStartDate = (*xsdDate)(&layout.T.PeriodStartDate)
	layout.PeriodEndDate = (*xsdDate)(&layout.T.PeriodEndDate)
	layout.PayrollPaymentDate = (*xsdDate)(layout.T.PayrollPaymentDate)
	layout.AccrualDate = (*xsdDate)(layout.T.AccrualDate)
	layout.AccrualReversalDate = (*xsdDate)(layout.T.AccrualReversalDate)
	layout.OpenTimeEntry = (*xsdDateTime)(layout.T.OpenTimeEntry)
	layout.LockTimeEntry = (*xsdDateTime)(layout.T.LockTimeEntry)
	layout.UnlockforAdjustments = (*xsdDateTime)(layout.T.UnlockforAdjustments)
	layout.CloseTimeEntry = (*xsdDateTime)(layout.T.CloseTimeEntry)
	return e.EncodeElement(layout, start)
}
func (t *PeriodDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PeriodDataType
	var overlay struct {
		*T
		PeriodStartDate      *xsdDate     `xml:"urn:com.workday/bsvc Period_Start_Date"`
		PeriodEndDate        *xsdDate     `xml:"urn:com.workday/bsvc Period_End_Date"`
		PayrollPaymentDate   *xsdDate     `xml:"urn:com.workday/bsvc Payroll_Payment_Date,omitempty"`
		AccrualDate          *xsdDate     `xml:"urn:com.workday/bsvc Accrual_Date,omitempty"`
		AccrualReversalDate  *xsdDate     `xml:"urn:com.workday/bsvc Accrual_Reversal_Date,omitempty"`
		OpenTimeEntry        *xsdDateTime `xml:"urn:com.workday/bsvc Open_Time_Entry,omitempty"`
		LockTimeEntry        *xsdDateTime `xml:"urn:com.workday/bsvc Lock_Time_Entry,omitempty"`
		UnlockforAdjustments *xsdDateTime `xml:"urn:com.workday/bsvc Unlock_for_Adjustments,omitempty"`
		CloseTimeEntry       *xsdDateTime `xml:"urn:com.workday/bsvc Close_Time_Entry,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PeriodStartDate = (*xsdDate)(&overlay.T.PeriodStartDate)
	overlay.PeriodEndDate = (*xsdDate)(&overlay.T.PeriodEndDate)
	overlay.PayrollPaymentDate = (*xsdDate)(overlay.T.PayrollPaymentDate)
	overlay.AccrualDate = (*xsdDate)(overlay.T.AccrualDate)
	overlay.AccrualReversalDate = (*xsdDate)(overlay.T.AccrualReversalDate)
	overlay.OpenTimeEntry = (*xsdDateTime)(overlay.T.OpenTimeEntry)
	overlay.LockTimeEntry = (*xsdDateTime)(overlay.T.LockTimeEntry)
	overlay.UnlockforAdjustments = (*xsdDateTime)(overlay.T.UnlockforAdjustments)
	overlay.CloseTimeEntry = (*xsdDateTime)(overlay.T.CloseTimeEntry)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PeriodDateIndicatorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PeriodDateIndicatorObjectType struct {
	ID         []PeriodDateIndicatorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PeriodObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PeriodObjectType struct {
	ID         []PeriodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Name of Period Schedule.
type PeriodScheduleDataType struct {
	PeriodScheduleID                     string                                  `xml:"urn:com.workday/bsvc Period_Schedule_ID,omitempty"`
	PeriodScheduleName                   string                                  `xml:"urn:com.workday/bsvc Period_Schedule_Name"`
	FrequencyReference                   *FrequencyObjectType                    `xml:"urn:com.workday/bsvc Frequency_Reference"`
	UsageReference                       []PeriodScheduleUsageObjectType         `xml:"urn:com.workday/bsvc Usage_Reference,omitempty"`
	PaymentDateAutoAdjustSaturday        float64                                 `xml:"urn:com.workday/bsvc Payment_Date_Auto-Adjust_Saturday,omitempty"`
	PaymentDateAutoAdjustSunday          float64                                 `xml:"urn:com.workday/bsvc Payment_Date_Auto-Adjust_Sunday,omitempty"`
	AllowTimesheetChanges                *bool                                   `xml:"urn:com.workday/bsvc Allow_Timesheet_Changes,omitempty"`
	PayrollPeriodScheduleReference       *PeriodScheduleObjectType               `xml:"urn:com.workday/bsvc Payroll_Period_Schedule_Reference,omitempty"`
	TimeTrackingEligibilityRuleReference []TimeTrackingEligibilityRuleObjectType `xml:"urn:com.workday/bsvc Time_Tracking_Eligibility_Rule_Reference,omitempty"`
	PeriodData                           []PeriodDataType                        `xml:"urn:com.workday/bsvc Period_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PeriodScheduleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PeriodScheduleObjectType struct {
	ID         []PeriodScheduleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Period Schedule Request Criteria
type PeriodScheduleRequestCriteriaType struct {
}

// Period Schedule Request References
type PeriodScheduleRequestReferencesType struct {
	PeriodScheduleReference []PeriodScheduleObjectType `xml:"urn:com.workday/bsvc Period_Schedule_Reference"`
}

// Period Schedule Response Data
type PeriodScheduleResponseDataType struct {
	PeriodSchedule []PeriodScheduleType `xml:"urn:com.workday/bsvc Period_Schedule,omitempty"`
}

// Period Schedule Response Group
type PeriodScheduleResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Period Schedule
type PeriodScheduleType struct {
	PeriodScheduleReference *PeriodScheduleObjectType `xml:"urn:com.workday/bsvc Period_Schedule_Reference,omitempty"`
	PeriodScheduleData      *PeriodScheduleDataType   `xml:"urn:com.workday/bsvc Period_Schedule_Data"`
}

// Contains a unique identifier for an instance of an object.
type PeriodScheduleUsageObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PeriodScheduleUsageObjectType struct {
	ID         []PeriodScheduleUsageObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Request Criteria for Periodic Company CAN Tax Remittance Data
type PeriodicCANTaxRemittanceDataRequestCriteriaType struct {
	CompanyReference                      *CompanyObjectType   `xml:"urn:com.workday/bsvc Company_Reference"`
	IncludeRelatedCompaniesforLegalEntity *bool                `xml:"urn:com.workday/bsvc Include_Related_Companies_for_Legal_Entity,omitempty"`
	ResultCompletedFrom                   *time.Time           `xml:"urn:com.workday/bsvc Result_Completed_From,omitempty"`
	ResultCompletedTo                     *time.Time           `xml:"urn:com.workday/bsvc Result_Completed_To,omitempty"`
	PayGroupReference                     []PayGroupObjectType `xml:"urn:com.workday/bsvc Pay_Group_Reference,omitempty"`
	PeriodReference                       []PeriodObjectType   `xml:"urn:com.workday/bsvc Period_Reference,omitempty"`
	IncludeWCB                            *bool                `xml:"urn:com.workday/bsvc Include_WCB,omitempty"`
	PaymentDate                           *time.Time           `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
}

func (t *PeriodicCANTaxRemittanceDataRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PeriodicCANTaxRemittanceDataRequestCriteriaType
	var layout struct {
		*T
		ResultCompletedFrom *xsdDateTime `xml:"urn:com.workday/bsvc Result_Completed_From,omitempty"`
		ResultCompletedTo   *xsdDateTime `xml:"urn:com.workday/bsvc Result_Completed_To,omitempty"`
		PaymentDate         *xsdDate     `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ResultCompletedFrom = (*xsdDateTime)(layout.T.ResultCompletedFrom)
	layout.ResultCompletedTo = (*xsdDateTime)(layout.T.ResultCompletedTo)
	layout.PaymentDate = (*xsdDate)(layout.T.PaymentDate)
	return e.EncodeElement(layout, start)
}
func (t *PeriodicCANTaxRemittanceDataRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PeriodicCANTaxRemittanceDataRequestCriteriaType
	var overlay struct {
		*T
		ResultCompletedFrom *xsdDateTime `xml:"urn:com.workday/bsvc Result_Completed_From,omitempty"`
		ResultCompletedTo   *xsdDateTime `xml:"urn:com.workday/bsvc Result_Completed_To,omitempty"`
		PaymentDate         *xsdDate     `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ResultCompletedFrom = (*xsdDateTime)(overlay.T.ResultCompletedFrom)
	overlay.ResultCompletedTo = (*xsdDateTime)(overlay.T.ResultCompletedTo)
	overlay.PaymentDate = (*xsdDate)(overlay.T.PaymentDate)
	return d.DecodeElement(&overlay, &start)
}

// Response Group for Periodic Company CAN Tax Remittance
type PeriodicCANTaxRemittanceDataResponseGroupType struct {
	IncludeMTDData *bool `xml:"urn:com.workday/bsvc Include_MTD_Data,omitempty"`
	IncludeQTDData *bool `xml:"urn:com.workday/bsvc Include_QTD_Data,omitempty"`
	IncludeYTDData *bool `xml:"urn:com.workday/bsvc Include_YTD_Data,omitempty"`
}

// Periodic Company CAN Tax Remittance Data Response Data element
type PeriodicCompanyCANTaxRemittanceDataResponseDataType struct {
	PeriodicCompanyCANTaxRemittanceData []PeriodicCompanyCANTaxRemittanceDataType `xml:"urn:com.workday/bsvc Periodic_Company_CAN_Tax_Remittance_Data,omitempty"`
}

// Periodic Company CAN Tax Remittance Data element
type PeriodicCompanyCANTaxRemittanceDataType struct {
	CompanyReference               *CompanyObjectType                          `xml:"urn:com.workday/bsvc Company_Reference"`
	PeriodEndDate                  *time.Time                                  `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	PaymentDate                    *time.Time                                  `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	CurrencyReference              *CurrencyObjectType                         `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	PayrollTaxAuthorityReference   *MetadataPayrollWorktagObjectType           `xml:"urn:com.workday/bsvc Payroll_Tax_Authority_Reference,omitempty"`
	PayrollTaxAccountNumber        string                                      `xml:"urn:com.workday/bsvc Payroll_Tax_Account_Number,omitempty"`
	WCBClassificationCodeReference *WorkersCompensationCodeObjectType          `xml:"urn:com.workday/bsvc WCB_Classification_Code_Reference,omitempty"`
	WCBPayrollReportId             string                                      `xml:"urn:com.workday/bsvc WCB_Payroll_Report_Id,omitempty"`
	WCBEBankingNumber              string                                      `xml:"urn:com.workday/bsvc WCB_E_Banking_Number,omitempty"`
	NumberEmployeesPaidLastPeriod  float64                                     `xml:"urn:com.workday/bsvc Number_Employees_Paid_Last_Period,omitempty"`
	PayGroupData                   []PeriodicCompanyTaxFilingPayGroupDataType  `xml:"urn:com.workday/bsvc Pay_Group_Data,omitempty"`
	PeriodReference                []PeriodObjectType                          `xml:"urn:com.workday/bsvc Period_Reference,omitempty"`
	EmployerPaid                   *bool                                       `xml:"urn:com.workday/bsvc Employer_Paid,omitempty"`
	DeductionReference             *PayrollCalculationObjectType               `xml:"urn:com.workday/bsvc Deduction_Reference,omitempty"`
	TaxFilingCode                  string                                      `xml:"urn:com.workday/bsvc Tax_Filing_Code,omitempty"`
	TaxRate                        float64                                     `xml:"urn:com.workday/bsvc Tax_Rate,omitempty"`
	WCBTaxRate                     float64                                     `xml:"urn:com.workday/bsvc WCB_Tax_Rate,omitempty"`
	TaxWithheld                    float64                                     `xml:"urn:com.workday/bsvc Tax_Withheld,omitempty"`
	TaxableWages                   float64                                     `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages                   float64                                     `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	ExemptWages                    float64                                     `xml:"urn:com.workday/bsvc Exempt_Wages,omitempty"`
	WCBInsurableWages              float64                                     `xml:"urn:com.workday/bsvc WCB_Insurable_Wages,omitempty"`
	WCBGrossWages                  float64                                     `xml:"urn:com.workday/bsvc WCB_Gross_Wages,omitempty"`
	WCBOtherWages                  float64                                     `xml:"urn:com.workday/bsvc WCB_Other_Wages,omitempty"`
	WCBExcessWages                 float64                                     `xml:"urn:com.workday/bsvc WCB_Excess_Wages,omitempty"`
	MTDData                        *MTDTaxRemittanceDataforPeriodicCompanyType `xml:"urn:com.workday/bsvc MTD_Data,omitempty"`
	QTDData                        *QTDTaxRemittanceDataforPeriodicCompanyType `xml:"urn:com.workday/bsvc QTD_Data,omitempty"`
	YTDData                        *YTDTaxRemittanceDataforPeriodicCompanyType `xml:"urn:com.workday/bsvc YTD_Data,omitempty"`
}

func (t *PeriodicCompanyCANTaxRemittanceDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PeriodicCompanyCANTaxRemittanceDataType
	var layout struct {
		*T
		PeriodEndDate *xsdDate `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
		PaymentDate   *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PeriodEndDate = (*xsdDate)(layout.T.PeriodEndDate)
	layout.PaymentDate = (*xsdDate)(layout.T.PaymentDate)
	return e.EncodeElement(layout, start)
}
func (t *PeriodicCompanyCANTaxRemittanceDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PeriodicCompanyCANTaxRemittanceDataType
	var overlay struct {
		*T
		PeriodEndDate *xsdDate `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
		PaymentDate   *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PeriodEndDate = (*xsdDate)(overlay.T.PeriodEndDate)
	overlay.PaymentDate = (*xsdDate)(overlay.T.PaymentDate)
	return d.DecodeElement(&overlay, &start)
}

// Periodic Company Tax Filing Data Response Data element
type PeriodicCompanyTaxFilingDataResponseDataType struct {
	PeriodicCompanyTaxFilingData []PeriodicCompanyTaxFilingDataType `xml:"urn:com.workday/bsvc Periodic_Company_Tax_Filing_Data,omitempty"`
}

// Periodic Company Tax Filing Data
type PeriodicCompanyTaxFilingDataType struct {
	CompanyReference             *CompanyObjectType                         `xml:"urn:com.workday/bsvc Company_Reference"`
	ThirdPartyProvidesSickPayW2  *bool                                      `xml:"urn:com.workday/bsvc Third_Party_Provides_Sick_Pay_W-2,omitempty"`
	PeriodEndDate                *time.Time                                 `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	PaymentDate                  *time.Time                                 `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	CurrencyReference            *CurrencyObjectType                        `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	PayGroupData                 []PeriodicCompanyTaxFilingPayGroupDataType `xml:"urn:com.workday/bsvc Pay_Group_Data,omitempty"`
	PeriodReference              []PeriodObjectType                         `xml:"urn:com.workday/bsvc Period_Reference,omitempty"`
	FrequencyReference           []FrequencyBehaviorObjectType              `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	ThirdPartySickPay            *bool                                      `xml:"urn:com.workday/bsvc Third_Party_Sick_Pay,omitempty"`
	EmployerPaid                 *bool                                      `xml:"urn:com.workday/bsvc Employer_Paid,omitempty"`
	DeductionReference           *PayrollCalculationObjectType              `xml:"urn:com.workday/bsvc Deduction_Reference,omitempty"`
	PayrollTaxAuthorityReference *MetadataPayrollWorktagObjectType          `xml:"urn:com.workday/bsvc Payroll_Tax_Authority_Reference,omitempty"`
	TaxFilingCode                string                                     `xml:"urn:com.workday/bsvc Tax_Filing_Code,omitempty"`
	MaleCount                    float64                                    `xml:"urn:com.workday/bsvc Male_Count,omitempty"`
	FemaleCount                  float64                                    `xml:"urn:com.workday/bsvc Female_Count,omitempty"`
	UnknownGenderCount           float64                                    `xml:"urn:com.workday/bsvc Unknown_Gender_Count,omitempty"`
	SUIRate                      float64                                    `xml:"urn:com.workday/bsvc SUI_Rate,omitempty"`
	TaxWithheld                  float64                                    `xml:"urn:com.workday/bsvc Tax_Withheld,omitempty"`
	TaxableWages                 float64                                    `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages                 float64                                    `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	GrossWages                   float64                                    `xml:"urn:com.workday/bsvc Gross_Wages,omitempty"`
	TipWages                     float64                                    `xml:"urn:com.workday/bsvc Tip_Wages,omitempty"`
	QTDData                      *QTDTaxFilingDataforPeriodicCompanyType    `xml:"urn:com.workday/bsvc QTD_Data,omitempty"`
	YTDData                      *YTDTaxFilingDataforPeriodicCompanyType    `xml:"urn:com.workday/bsvc YTD_Data,omitempty"`
}

func (t *PeriodicCompanyTaxFilingDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PeriodicCompanyTaxFilingDataType
	var layout struct {
		*T
		PeriodEndDate *xsdDate `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
		PaymentDate   *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PeriodEndDate = (*xsdDate)(layout.T.PeriodEndDate)
	layout.PaymentDate = (*xsdDate)(layout.T.PaymentDate)
	return e.EncodeElement(layout, start)
}
func (t *PeriodicCompanyTaxFilingDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PeriodicCompanyTaxFilingDataType
	var overlay struct {
		*T
		PeriodEndDate *xsdDate `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
		PaymentDate   *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PeriodEndDate = (*xsdDate)(overlay.T.PeriodEndDate)
	overlay.PaymentDate = (*xsdDate)(overlay.T.PaymentDate)
	return d.DecodeElement(&overlay, &start)
}

// Pay Group Data element for Pay Group Rerences
type PeriodicCompanyTaxFilingPayGroupDataType struct {
	PayGroupReference *PayGroupObjectType `xml:"urn:com.workday/bsvc Pay_Group_Reference,omitempty"`
}

// Request Criteria for Periodic Worker Tax Filing Data
type PeriodicTaxFilingDataRequestCriteriaType struct {
	CompanyReference                      *CompanyObjectType   `xml:"urn:com.workday/bsvc Company_Reference"`
	IncludeRelatedCompaniesforLegalEntity *bool                `xml:"urn:com.workday/bsvc Include_Related_Companies_for_Legal_Entity,omitempty"`
	ResultCompletedFrom                   *time.Time           `xml:"urn:com.workday/bsvc Result_Completed_From,omitempty"`
	ResultCompletedTo                     *time.Time           `xml:"urn:com.workday/bsvc Result_Completed_To,omitempty"`
	PayGroupReference                     []PayGroupObjectType `xml:"urn:com.workday/bsvc Pay_Group_Reference,omitempty"`
	PeriodReference                       []PeriodObjectType   `xml:"urn:com.workday/bsvc Period_Reference,omitempty"`
	PaymentDate                           *time.Time           `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	YTDTaxFilter                          []YTDTaxFilterType   `xml:"urn:com.workday/bsvc YTD_Tax_Filter,omitempty"`
}

func (t *PeriodicTaxFilingDataRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PeriodicTaxFilingDataRequestCriteriaType
	var layout struct {
		*T
		ResultCompletedFrom *xsdDateTime `xml:"urn:com.workday/bsvc Result_Completed_From,omitempty"`
		ResultCompletedTo   *xsdDateTime `xml:"urn:com.workday/bsvc Result_Completed_To,omitempty"`
		PaymentDate         *xsdDate     `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ResultCompletedFrom = (*xsdDateTime)(layout.T.ResultCompletedFrom)
	layout.ResultCompletedTo = (*xsdDateTime)(layout.T.ResultCompletedTo)
	layout.PaymentDate = (*xsdDate)(layout.T.PaymentDate)
	return e.EncodeElement(layout, start)
}
func (t *PeriodicTaxFilingDataRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PeriodicTaxFilingDataRequestCriteriaType
	var overlay struct {
		*T
		ResultCompletedFrom *xsdDateTime `xml:"urn:com.workday/bsvc Result_Completed_From,omitempty"`
		ResultCompletedTo   *xsdDateTime `xml:"urn:com.workday/bsvc Result_Completed_To,omitempty"`
		PaymentDate         *xsdDate     `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ResultCompletedFrom = (*xsdDateTime)(overlay.T.ResultCompletedFrom)
	overlay.ResultCompletedTo = (*xsdDateTime)(overlay.T.ResultCompletedTo)
	overlay.PaymentDate = (*xsdDate)(overlay.T.PaymentDate)
	return d.DecodeElement(&overlay, &start)
}

// Response Group for Periodic Worker Tax Filing Response Group
type PeriodicTaxFilingDataResponseGroupType struct {
	IncludeQTDData *bool `xml:"urn:com.workday/bsvc Include_QTD_Data,omitempty"`
	IncludeYTDData *bool `xml:"urn:com.workday/bsvc Include_YTD_Data,omitempty"`
}

// Get Periodic Worker Tax Filing Data Response
type PeriodicWorkerTaxFilingDataResponseDataType struct {
	PeriodicWorkerTaxFilingData []PeridoicWorkerTaxFilingDataType `xml:"urn:com.workday/bsvc Periodic_Worker_Tax_Filing_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PhoneDeviceTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PhoneDeviceTypeObjectType struct {
	ID         []PhoneDeviceTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Phone Information
type PhoneInformationDataType struct {
	CountryISOCode              string                                        `xml:"urn:com.workday/bsvc Country_ISO_Code,omitempty"`
	InternationalPhoneCode      string                                        `xml:"urn:com.workday/bsvc International_Phone_Code,omitempty"`
	PhoneNumber                 string                                        `xml:"urn:com.workday/bsvc Phone_Number,omitempty"`
	PhoneExtension              string                                        `xml:"urn:com.workday/bsvc Phone_Extension,omitempty"`
	PhoneDeviceTypeReference    *PhoneDeviceTypeObjectType                    `xml:"urn:com.workday/bsvc Phone_Device_Type_Reference,omitempty"`
	UsageData                   []CommunicationMethodUsageInformationDataType `xml:"urn:com.workday/bsvc Usage_Data,omitempty"`
	PhoneReference              *PhoneReferenceObjectType                     `xml:"urn:com.workday/bsvc Phone_Reference,omitempty"`
	ID                          string                                        `xml:"urn:com.workday/bsvc ID,omitempty"`
	AreaCode                    string                                        `xml:"urn:com.workday/bsvc Area_Code,attr,omitempty"`
	TenantFormattedPhone        string                                        `xml:"urn:com.workday/bsvc Tenant_Formatted_Phone,attr,omitempty"`
	InternationalFormattedPhone string                                        `xml:"urn:com.workday/bsvc International_Formatted_Phone,attr,omitempty"`
	PhoneNumberWithoutAreaCode  string                                        `xml:"urn:com.workday/bsvc Phone_Number_Without_Area_Code,attr,omitempty"`
	NationalFormattedPhone      string                                        `xml:"urn:com.workday/bsvc National_Formatted_Phone,attr,omitempty"`
	E164FormattedPhone          string                                        `xml:"urn:com.workday/bsvc E164_Formatted_Phone,attr,omitempty"`
	Delete                      bool                                          `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	DoNotReplaceAll             bool                                          `xml:"urn:com.workday/bsvc Do_Not_Replace_All,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PhoneReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PhoneReferenceObjectType struct {
	ID         []PhoneReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PositionElementObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PositionElementObjectType struct {
	ID         []PositionElementObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PositionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PositionObjectType struct {
	ID         []PositionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PositionRestrictionsObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PositionRestrictionsObjectType struct {
	ID         []PositionRestrictionsObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element for all Prefix Name data.
type PrefixNameDataType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc Type,attr,omitempty"`
}

type ProcessingFaultType struct {
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProgramObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProgramObjectType struct {
	ID         []ProgramObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProjectObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProjectObjectType struct {
	ID         []ProjectObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProjectPlanPhaseObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProjectPlanPhaseObjectType struct {
	ID         []ProjectPlanPhaseObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProjectPlanTaskObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProjectPlanTaskObjectType struct {
	ID         []ProjectPlanTaskObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This is Canadian Record of Employment Request element that contains the Canadian Record of Employment attributes to be updated.
type PutCanadianRecordofEmploymentDataRequestType struct {
	CanadianRecordofEmploymentDataData []CanadianRecordofEmploymentDataDataType `xml:"urn:com.workday/bsvc Canadian_Record_of_Employment_Data_Data,omitempty"`
	Version                            string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This is the Canadian Record of Employment Data Response element
type PutCanadianRecordofEmploymentDataResponseType struct {
	CanadianRecordofEmploymentDataReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Canadian_Record_of_Employment_Data_Reference,omitempty"`
	Version                                 string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Import Process Response element
type PutImportProcessResponseType struct {
	ImportProcessReference  *WebServiceBackgroundProcessRuntimeObjectType `xml:"urn:com.workday/bsvc Import_Process_Reference,omitempty"`
	HeaderInstanceReference *InstanceObjectType                           `xml:"urn:com.workday/bsvc Header_Instance_Reference,omitempty"`
	Version                 string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top Level element containing PUT request information.
type PutPaycheckDeliveryPublicRequestType struct {
	PaycheckDeliveryData *PaycheckDeliveryDataType `xml:"urn:com.workday/bsvc Paycheck_Delivery_Data"`
	Version              string                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Paycheck Printing Preference Record for Worker
type PutPaycheckDeliveryPublicResponseType struct {
	PaycheckDeliveryReference        *PaycheckDeliveryObjectType        `xml:"urn:com.workday/bsvc Paycheck_Delivery_Reference,omitempty"`
	WorkerReference                  *WorkerObjectType                  `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	CompanyReference                 *CompanyObjectType                 `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	PaycheckDeliveryMethodReference  *PaycheckDeliveryMethodObjectType  `xml:"urn:com.workday/bsvc Paycheck_Delivery_Method_Reference,omitempty"`
	PayslipPrintingOverrideReference *PayslipPrintingOverrideObjectType `xml:"urn:com.workday/bsvc Payslip_Printing_Override_Reference,omitempty"`
	LastUpdated                      *time.Time                         `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	Version                          string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

func (t *PutPaycheckDeliveryPublicResponseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PutPaycheckDeliveryPublicResponseType
	var layout struct {
		*T
		LastUpdated *xsdDate `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LastUpdated = (*xsdDate)(layout.T.LastUpdated)
	return e.EncodeElement(layout, start)
}
func (t *PutPaycheckDeliveryPublicResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PutPaycheckDeliveryPublicResponseType
	var overlay struct {
		*T
		LastUpdated *xsdDate `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LastUpdated = (*xsdDate)(overlay.T.LastUpdated)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for all Payroll Deduction Recipient data
type PutPayrollDeductionRecipientRequestType struct {
	PayrollDeductionRecipientReference *DeductionRecipientObjectType      `xml:"urn:com.workday/bsvc Payroll_Deduction_Recipient_Reference,omitempty"`
	PayrollDeductionRecipientData      *PayrollDeductionRecipientDataType `xml:"urn:com.workday/bsvc Payroll_Deduction_Recipient_Data,omitempty"`
	AddOnly                            bool                               `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                            string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Payroll Deduction Recipient Wrapper element
type PutPayrollDeductionRecipientResponseType struct {
	PayrollDeductionRecipientReference *DeductionRecipientObjectType `xml:"urn:com.workday/bsvc Payroll_Deduction_Recipient_Reference,omitempty"`
	Version                            string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Payroll Payee W-4 Request Element
type PutPayrollFederalW4TaxElectionRequestType struct {
	PayrollFederalW4TaxElectionData *PayrollFederalW4TaxElectionDataType `xml:"urn:com.workday/bsvc Payroll_Federal_W-4_Tax_Election_Data"`
	Version                         string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The data that should be displayed for the W-4 that was made as of the Put
type PutPayrollFederalW4TaxElectionResponseDataType struct {
	WorkerReference                        *WorkerObjectType  `xml:"urn:com.workday/bsvc Worker_Reference"`
	Effectiveasof                          time.Time          `xml:"urn:com.workday/bsvc Effective_as_of"`
	CompanyforPayrollPayeeTaxDataReference *CompanyObjectType `xml:"urn:com.workday/bsvc Company_for_Payroll_Payee_Tax_Data_Reference"`
}

func (t *PutPayrollFederalW4TaxElectionResponseDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PutPayrollFederalW4TaxElectionResponseDataType
	var layout struct {
		*T
		Effectiveasof *xsdDate `xml:"urn:com.workday/bsvc Effective_as_of"`
	}
	layout.T = (*T)(t)
	layout.Effectiveasof = (*xsdDate)(&layout.T.Effectiveasof)
	return e.EncodeElement(layout, start)
}
func (t *PutPayrollFederalW4TaxElectionResponseDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PutPayrollFederalW4TaxElectionResponseDataType
	var overlay struct {
		*T
		Effectiveasof *xsdDate `xml:"urn:com.workday/bsvc Effective_as_of"`
	}
	overlay.T = (*T)(t)
	overlay.Effectiveasof = (*xsdDate)(&overlay.T.Effectiveasof)
	return d.DecodeElement(&overlay, &start)
}

// Put Payroll Payee W-4 Response Element
type PutPayrollFederalW4TaxElectionResponseType struct {
	PayrollPayeeW4Reference                           *UniqueIdentifierObjectType                     `xml:"urn:com.workday/bsvc Payroll_Payee_W-4_Reference,omitempty"`
	PutPayrollFederalW4TaxElectionResponseDataElement *PutPayrollFederalW4TaxElectionResponseDataType `xml:"urn:com.workday/bsvc Put_Payroll_Federal_W-4_Tax_Election_Response_Data_Element,omitempty"`
	Version                                           string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This is the top-level Request element carrying all information necessary to make a Put Payroll History Payment service request.
type PutPayrollHistoryPaymentRequestType struct {
	PayrollHistoryPaymentReference *PayrollOffcycleResultOverrideObjectType `xml:"urn:com.workday/bsvc Payroll_History_Payment_Reference,omitempty"`
	PayrollHistoryPaymentData      *PayrollHistoryPaymentDataType           `xml:"urn:com.workday/bsvc Payroll_History_Payment_Data"`
	Version                        string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This element contains the reference to the Payroll History Payment added or updated by the web service operation.
type PutPayrollHistoryPaymentResponseType struct {
	PayrollHistoryPaymentReference *PayrollOffcycleResultOverrideObjectType `xml:"urn:com.workday/bsvc Payroll_History_Payment_Reference,omitempty"`
	Version                        string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Payroll Involuntary Withholding Order Request
type PutPayrollInvoluntaryWithholdingOrderRequestType struct {
	PayrollInvoluntaryWithholdingOrderReference *WithholdingOrderObjectType                  `xml:"urn:com.workday/bsvc Payroll_Involuntary_Withholding_Order_Reference,omitempty"`
	PayrollSupportOrderAmendorTerminateData     *PayrollSupportOrderAmendorTerminateDataType `xml:"urn:com.workday/bsvc Payroll_Support_Order_Amend_or_Terminate_Data,omitempty"`
	PayrollInvoluntaryWithholdingOrderData      *PayrollInvoluntaryWithholdingOrderDataType  `xml:"urn:com.workday/bsvc Payroll_Involuntary_Withholding_Order_Data,omitempty"`
	Version                                     string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Payroll Involuntary Withholding Order Response Data
type PutPayrollInvoluntaryWithholdingOrderResponseDataType struct {
	WithholdingOrderReference     *WithholdingOrderObjectType     `xml:"urn:com.workday/bsvc Withholding_Order_Reference,omitempty"`
	WorkerReference               *WorkerObjectType               `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	WithholdingOrderType          string                          `xml:"urn:com.workday/bsvc Withholding_Order_Type,omitempty"`
	WithholdingOrderCaseReference *WithholdingOrderCaseObjectType `xml:"urn:com.workday/bsvc Withholding_Order_Case_Reference,omitempty"`
}

// Put Payroll Involuntary Withholding Order Response
type PutPayrollInvoluntaryWithholdingOrderResponseType struct {
	PutPayrollInvoluntaryWithholdingOrderResponseData []PutPayrollInvoluntaryWithholdingOrderResponseDataType `xml:"urn:com.workday/bsvc Put_Payroll_Involuntary_Withholding_Order_Response_Data,omitempty"`
	Version                                           string                                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold all Payroll Limit Override Request Information for Put.
type PutPayrollLimitOverrideRequestType struct {
	PayrollLimitOverrideReference *PayrollLimitOverrideObjectType `xml:"urn:com.workday/bsvc Payroll_Limit_Override_Reference,omitempty"`
	PayrollLimitOverrideData      *PayrollLimitOverrideDataType   `xml:"urn:com.workday/bsvc Payroll_Limit_Override_Data,omitempty"`
	Version                       string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Payroll Limit Override Reponse Element.
type PutPayrollLimitOverrideResponseType struct {
	PayrollLimitOverrideReference *PayrollLimitOverrideObjectType `xml:"urn:com.workday/bsvc Payroll_Limit_Override_Reference,omitempty"`
	Version                       string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Payroll Off-cycle Payment Request
type PutPayrollOffcyclePaymentRequestType struct {
	PayrollOffcyclePaymentReference *PayrollOffcycleResultOverrideObjectType `xml:"urn:com.workday/bsvc Payroll_Off-cycle_Payment_Reference,omitempty"`
	PayrollOffcyclePaymentData      *PayrollOffcyclePaymentDataType          `xml:"urn:com.workday/bsvc Payroll_Off-cycle_Payment_Data"`
	Version                         string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Payroll Off-cycle Payment Response
type PutPayrollOffcyclePaymentResponseType struct {
	PayrollOffcyclePaymentReference *PayrollOffcycleResultOverrideObjectType `xml:"urn:com.workday/bsvc Payroll_Off-cycle_Payment_Reference,omitempty"`
	BatchID                         string                                   `xml:"urn:com.workday/bsvc Batch_ID,omitempty"`
	Version                         string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for Payroll Payee FICA Tax Election data
type PutPayrollPayeeFICARequestType struct {
	PayrollPayeeFICAData *PayrollPayeeFICADataType `xml:"urn:com.workday/bsvc Payroll_Payee_FICA_Data,omitempty"`
	Version              string                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Payroll Payee FICA Tax Election data
type PutPayrollPayeeFICAResponseType struct {
	PayrollPayeeFICAReference []PayrollPayeeTaxDataObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_FICA_Reference,omitempty"`
	Version                   string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// FUTA Record Data
type PutPayrollPayeeFUTARequestType struct {
	PayrollPayeeFUTAData *PayrollPayeeFUTADataType `xml:"urn:com.workday/bsvc Payroll_Payee_FUTA_Data,omitempty"`
	Version              string                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The data that should be displayed for the FUTA that was made as of the Put
type PutPayrollPayeeFUTAResponseDataType struct {
	WorkerReference  *WorkerObjectType  `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	CompanyReference *CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	EffectiveDate    *time.Time         `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
}

func (t *PutPayrollPayeeFUTAResponseDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PutPayrollPayeeFUTAResponseDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PutPayrollPayeeFUTAResponseDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PutPayrollPayeeFUTAResponseDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Shows what record was entered as a response.
type PutPayrollPayeeFUTAResponseType struct {
	PayrollPayeeFUTAReference       []PayrollPayeeFUTAObjectType          `xml:"urn:com.workday/bsvc Payroll_Payee_FUTA_Reference,omitempty"`
	PutPayrollPayeeFUTAResponseData []PutPayrollPayeeFUTAResponseDataType `xml:"urn:com.workday/bsvc Put_Payroll_Payee_FUTA_Response_Data,omitempty"`
	Version                         string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold Payroll Payee Ongoing Work Jurisdiction Tax Election Request Information
type PutPayrollPayeeOngoingWorkJurisdictionTaxElectionRequestType struct {
	PayrollPayeeOngoingWorkJurisdictionTaxElectionData *PayrollPayeeOngoingWorkJurisdictionTaxElectionDataType `xml:"urn:com.workday/bsvc Payroll_Payee_Ongoing_Work_Jurisdiction_Tax_Election_Data"`
	Version                                            string                                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold Payroll Payee Ongoing Jurisdiction Tax Election Response Data Information
type PutPayrollPayeeOngoingWorkJurisdictionTaxElectionResponseDataType struct {
	WorkerforTaxElectionReference  *WorkerObjectType  `xml:"urn:com.workday/bsvc Worker_for_Tax_Election_Reference"`
	EffectiveDate                  time.Time          `xml:"urn:com.workday/bsvc Effective_Date"`
	CompanyforTaxElectionReference *CompanyObjectType `xml:"urn:com.workday/bsvc Company_for_Tax_Election_Reference"`
}

func (t *PutPayrollPayeeOngoingWorkJurisdictionTaxElectionResponseDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PutPayrollPayeeOngoingWorkJurisdictionTaxElectionResponseDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PutPayrollPayeeOngoingWorkJurisdictionTaxElectionResponseDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PutPayrollPayeeOngoingWorkJurisdictionTaxElectionResponseDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Element to hold Payroll Payee Ongoing Work Jurisdiction Tax Election Response Information
type PutPayrollPayeeOngoingWorkJurisdictionTaxElectionResponseType struct {
	PutPayrollPayeeOngoingJurisdictionReference *UniqueIdentifierObjectType                                        `xml:"urn:com.workday/bsvc Put_Payroll_Payee_Ongoing_Jurisdiction_Reference,omitempty"`
	OngoingJurisdictionResponseTaxElectionData  *PutPayrollPayeeOngoingWorkJurisdictionTaxElectionResponseDataType `xml:"urn:com.workday/bsvc Ongoing_Jurisdiction_Response_Tax_Election_Data,omitempty"`
	Version                                     string                                                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Payroll Payee PT1 Request
type PutPayrollPayeePT1RequestType struct {
	PayrollPayeePT1Data *PayrollPayeePT1DataType `xml:"urn:com.workday/bsvc Payroll_Payee_PT1_Data"`
	Version             string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Shows what record was entered as a response.
type PutPayrollPayeePT1ResponseType struct {
	PayrollPayeePT1Reference *PayrollPayeePT1ObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_PT1_Reference,omitempty"`
	Version                  string                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// RPP or DPSP Registration Number Record Data
type PutPayrollPayeeRPPorDPSPRegistrationNumberRequestType struct {
	PayrollPayeeRPPorDPSPRegistrationNumberData *PayrollPayeeRPPorDPSPRegistrationNumberDataType `xml:"urn:com.workday/bsvc Payroll_Payee_RPP_or_DPSP_Registration_Number_Data,omitempty"`
	Version                                     string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Shows what record was entered as a response.
type PutPayrollPayeeRPPorDPSPRegistrationNumberResponseType struct {
	PayrollPayeeRPPorDPSPRegistrationNumberReference *PayrollPayeeRPPorDPSPRegistrationNumbersObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_RPP_or_DPSP_Registration_Number_Reference,omitempty"`
	Version                                          string                                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Payroll Payee T1 Response
type PutPayrollPayeeT1ResponseType struct {
	PayrollPayeeT1Reference *PayrollPayeeT1ObjectType `xml:"urn:com.workday/bsvc Payroll_Payee_T1_Reference,omitempty"`
	Version                 string                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Payroll Payee TD1 Request
type PutPayrollPayeeTD1RequestType struct {
	PayrollPayeeTD1Data *PayrollPayeeTD1DataType `xml:"urn:com.workday/bsvc Payroll_Payee_TD1_Data"`
	Version             string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold Payroll Payee Tax Location Mapping Request Information
type PutPayrollPayeeTaxLocationMappingRequestType struct {
	PayrollPayeeTaxLocationMappingData *PayrollPayeeTaxLocationMappingDataType `xml:"urn:com.workday/bsvc Payroll_Payee_Tax_Location_Mapping_Data"`
	Version                            string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold Payroll Payee Tax Location Mapping Response Information
type PutPayrollPayeeTaxLocationMappingResponseType struct {
	PayrollPayeeTaxLocationMappingReference *PayrollPayeeTaxLocationMappingObjectType     `xml:"urn:com.workday/bsvc Payroll_Payee_Tax_Location_Mapping_Reference,omitempty"`
	PayrollPayeeTaxLocationMappingData      *ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Payroll_Payee_Tax_Location_Mapping_Data,omitempty"`
	Version                                 string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold Payroll Tax Location Mapping Response Data Information
type PutPayrollTaxLocationMappingResponseDataType struct {
	EffectiveDate                  *time.Time                       `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	Inactive                       *bool                            `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	PayrollStateAuthorityReference *PayrollStateAuthorityObjectType `xml:"urn:com.workday/bsvc Payroll_State_Authority_Reference,omitempty"`
}

func (t *PutPayrollTaxLocationMappingResponseDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PutPayrollTaxLocationMappingResponseDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PutPayrollTaxLocationMappingResponseDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PutPayrollTaxLocationMappingResponseDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Element to hold Payroll Tax Location Mapping Request Information
type PutPayrollTaxMappingonLocationRequestType struct {
	PayrollTaxLocationMappingData *PayrollTaxLocationMappingDataType `xml:"urn:com.workday/bsvc Payroll_Tax_Location_Mapping_Data"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold Payroll Tax Location Mapping Response Information
type PutPayrollTaxMappingonLocationResponseType struct {
	PutPayrollTaxLocationMappingReference    *PayrollTaxLocationMappingObjectType          `xml:"urn:com.workday/bsvc Put_Payroll_Tax_Location_Mapping_Reference,omitempty"`
	PutPayrollTaxLocationMappingResponseData *PutPayrollTaxLocationMappingResponseDataType `xml:"urn:com.workday/bsvc Put_Payroll_Tax_Location_Mapping_Response_Data,omitempty"`
	Version                                  string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// State and Local Tax Elections for USA Payroll State Authorities
type PutPayrollUSAStateandLocalTaxElectionRequestType struct {
	CompanyReference                       *CompanyObjectType         `xml:"urn:com.workday/bsvc Company_Reference"`
	WorkerReference                        *WorkerObjectType          `xml:"urn:com.workday/bsvc Worker_Reference"`
	EffectiveDate                          time.Time                  `xml:"urn:com.workday/bsvc Effective_Date"`
	PayrollUSAStateandLocalTaxElectionData []StateTaxElectionDataType `xml:"urn:com.workday/bsvc Payroll_USA_State_and_Local_Tax_Election_Data"`
	Version                                string                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

func (t *PutPayrollUSAStateandLocalTaxElectionRequestType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PutPayrollUSAStateandLocalTaxElectionRequestType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PutPayrollUSAStateandLocalTaxElectionRequestType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PutPayrollUSAStateandLocalTaxElectionRequestType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// The results of the last Put Payroll USA State and Local Tax Election Put Request
type PutPayrollUSAStateandLocalTaxElectionResponseType struct {
	CompanyReference *CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	WorkerReference  *WorkerObjectType  `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	EffectiveDate    *time.Time         `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	Version          string             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

func (t *PutPayrollUSAStateandLocalTaxElectionResponseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PutPayrollUSAStateandLocalTaxElectionResponseType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PutPayrollUSAStateandLocalTaxElectionResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PutPayrollUSAStateandLocalTaxElectionResponseType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Create or Update a Tax Treaty for a Worker
type PutPayrollWorkerTaxTreatyRequestType struct {
	WorkerTaxTreatyData *WorkerTaxTreatyDataType `xml:"urn:com.workday/bsvc Worker_Tax_Treaty_Data,omitempty"`
	Version             string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Add/Update Period Schedule Request
type PutPeriodScheduleRequestType struct {
	PeriodScheduleReference *PeriodScheduleObjectType `xml:"urn:com.workday/bsvc Period_Schedule_Reference,omitempty"`
	PeriodScheduleData      *PeriodScheduleDataType   `xml:"urn:com.workday/bsvc Period_Schedule_Data,omitempty"`
	AddOnly                 bool                      `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                 string                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Add/Update Period Schedule Response
type PutPeriodScheduleResponseType struct {
	PeriodScheduleReference *PeriodScheduleObjectType `xml:"urn:com.workday/bsvc Period_Schedule_Reference,omitempty"`
	Version                 string                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Canadian Record of Employment History data request element that contains Record of Employment History data to be updated.
type PutROEHistoryDataRequestType struct {
	ROEHistoryData *ROEHistoryDataType `xml:"urn:com.workday/bsvc ROE_History_Data"`
	Version        string              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This is the Canadian Record of Employment History Data Response element
type PutROEHistoryDataResponseType struct {
	ROEDataReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc ROE_Data_Reference,omitempty"`
	Version          string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Initial element for the Put Run Pay Calculation
type PutRunPayCalculationRequestType struct {
	RunPayCalculationData *RunPayCalculationDataType `xml:"urn:com.workday/bsvc Run_Pay_Calculation_Data"`
	Version               string                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response Element for Put Pay Calculation
type PutRunPayCalculationResponseType struct {
	ApplicationInstanceRelatedExceptionsData []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Application_Instance_Related_Exceptions_Data,omitempty"`
	BackgroundProcessRuntimeData             []BackgroundProcessRuntimeDataType             `xml:"urn:com.workday/bsvc Background_Process_Runtime_Data,omitempty"`
	Version                                  string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Single Legal Entity Request
type PutSingleLegalEntityRequestType struct {
	SingleLegalEntityReference *SingleLegalEntityObjectType `xml:"urn:com.workday/bsvc Single_Legal_Entity_Reference,omitempty"`
	SingleLegalEntityData      *SingleLegalEntityDataType   `xml:"urn:com.workday/bsvc Single_Legal_Entity_Data,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Single Legal Entity Response
type PutSingleLegalEntityResponseType struct {
	SingleLegalEntityReference *SingleLegalEntityObjectType `xml:"urn:com.workday/bsvc Single_Legal_Entity_Reference,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Successor Employer Request
type PutSuccessorEmployerRequestType struct {
	SuccessorEmployerReference *SuccessorEmployerObjectType `xml:"urn:com.workday/bsvc Successor_Employer_Reference,omitempty"`
	SuccessorEmployerData      *SuccessorEmployerDataType   `xml:"urn:com.workday/bsvc Successor_Employer_Data,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Successor Employer Response
type PutSuccessorEmployerResponseType struct {
	SuccessorEmployerReference *SuccessorEmployerObjectType `xml:"urn:com.workday/bsvc Successor_Employer_Reference,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request structure for adding or updating a payroll tax document printing election for a worker.
type PutTaxDocumentDeliveryRequestType struct {
	TaxDocumentDeliveryReference *UniqueIdentifierObjectType  `xml:"urn:com.workday/bsvc Tax_Document_Delivery_Reference,omitempty"`
	TaxDocumentDeliveryData      *TaxDocumentDeliveryDataType `xml:"urn:com.workday/bsvc Tax_Document_Delivery_Data,omitempty"`
	Version                      string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Tax document delivery response
type PutTaxDocumentDeliveryResponseType struct {
	TaxDocumentDeliveryReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Tax_Document_Delivery_Reference,omitempty"`
	Version                      string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Tax Levy Deduction Restriction Request
type PutTaxLevyDeductionRestrictionRequestType struct {
	TaxLevyReference                *TaxLevyObjectType                   `xml:"urn:com.workday/bsvc Tax_Levy_Reference"`
	TaxLevyDeductionRestrictionData *TaxLevyDeductionRestrictionDataType `xml:"urn:com.workday/bsvc Tax_Levy_Deduction_Restriction_Data,omitempty"`
	Version                         string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Tax Levy Deduction Restriction Response
type PutTaxLevyDeductionRestrictionResponseType struct {
	TaxLevyReference                *TaxLevyObjectType                    `xml:"urn:com.workday/bsvc Tax_Levy_Reference,omitempty"`
	TaxLevyDeductionRestrictionData []TaxLevyDeductionRestrictionDataType `xml:"urn:com.workday/bsvc Tax_Levy_Deduction_Restriction_Data,omitempty"`
	Version                         string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put W-2/W-2C Printing Election Request Element
type PutW2W2CPrintingElectionRequestType struct {
	W2W2CPrintingElectionRequestData []W2W2CPrintingElectionRequestDataType `xml:"urn:com.workday/bsvc W-2_W-2C_Printing_Election_Request_Data,omitempty"`
	Version                          string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element to hold W-2/W-2C Printing Election Response Data
type PutW2W2CPrintingElectionforWorkerResponseType struct {
	WorkerwithUpdatedW2W2CPrintElectionReference []WorkerObjectType `xml:"urn:com.workday/bsvc Worker_with_Updated_W-2_W-2C_Print_Election_Reference,omitempty"`
	Version                                      string             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Withholding Order Local Minimum Wage Data
type PutWithholdingOrderLocalMinimumWageDataType struct {
	PayrollInvoluntaryWithholdingOrderReference *WithholdingOrderObjectType `xml:"urn:com.workday/bsvc Payroll_Involuntary_Withholding_Order_Reference,omitempty"`
	LocalMinimumWageReference                   *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Local_Minimum_Wage_Reference,omitempty"`
	EffectiveDate                               time.Time                   `xml:"urn:com.workday/bsvc Effective_Date"`
	LocalMinimumWage                            float64                     `xml:"urn:com.workday/bsvc Local_Minimum_Wage,omitempty"`
	Comment                                     string                      `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

func (t *PutWithholdingOrderLocalMinimumWageDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PutWithholdingOrderLocalMinimumWageDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PutWithholdingOrderLocalMinimumWageDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PutWithholdingOrderLocalMinimumWageDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Put Withholding Order Local Minimum Wage Rate Request
type PutWithholdingOrderLocalMinimumWageRateRequestType struct {
	WithholdingOrderLocalMinimumWageData []PutWithholdingOrderLocalMinimumWageDataType `xml:"urn:com.workday/bsvc Withholding_Order_Local_Minimum_Wage_Data"`
	Version                              string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Withholding Order Local Minimum Wage Rate Response
type PutWithholdingOrderLocalMinimumWageRateResponseType struct {
	LocalMinimumWageReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Local_Minimum_Wage_Reference,omitempty"`
	Version                   string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Worker Tax Treaty Response Data
type PutWorkerTaxTreatyResponseDataType struct {
	WorkerReference     *WorkerObjectType            `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	TaxYearReference    *CalendarYearObjectType      `xml:"urn:com.workday/bsvc Tax_Year_Reference,omitempty"`
	IncomeCodeReference *PayrollIncomeCodeObjectType `xml:"urn:com.workday/bsvc Income_Code_Reference,omitempty"`
}

// Put Worker Tax Treaty Response
type PutWorkerTaxTreatyResponseType struct {
	WorkerTaxTreatyReference    *PayrollPayeeTaxTreatyUSAObjectType  `xml:"urn:com.workday/bsvc Worker_Tax_Treaty_Reference,omitempty"`
	WorkerTaxTreatyResponseData []PutWorkerTaxTreatyResponseDataType `xml:"urn:com.workday/bsvc Worker_Tax_Treaty_Response_Data,omitempty"`
	Version                     string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// QTD Tax Filing Data for Periodic Company
type QTDTaxFilingDataforPeriodicCompanyType struct {
	TaxWithheld  float64 `xml:"urn:com.workday/bsvc Tax_Withheld,omitempty"`
	TaxableWages float64 `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages float64 `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	GrossWages   float64 `xml:"urn:com.workday/bsvc Gross_Wages,omitempty"`
}

// QTD Tax Filing Data for Periodic Worker
type QTDTaxFilingDataforPeriodicWorkerType struct {
	TaxWithheld  float64 `xml:"urn:com.workday/bsvc Tax_Withheld,omitempty"`
	TaxableWages float64 `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages float64 `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	GrossWages   float64 `xml:"urn:com.workday/bsvc Gross_Wages,omitempty"`
	TipWages     float64 `xml:"urn:com.workday/bsvc Tip_Wages,omitempty"`
}

// QTD Tax Filing Data for Quarterly Worker
type QTDTaxFilingDataforQuarterlyWorkerType struct {
	TaxWithheld       float64 `xml:"urn:com.workday/bsvc Tax_Withheld,omitempty"`
	TaxableWages      float64 `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages      float64 `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	GrossWages        float64 `xml:"urn:com.workday/bsvc Gross_Wages,omitempty"`
	TipWages          float64 `xml:"urn:com.workday/bsvc Tip_Wages,omitempty"`
	UncollectedTipTax float64 `xml:"urn:com.workday/bsvc Uncollected_Tip_Tax,omitempty"`
	UncollectedGTL    float64 `xml:"urn:com.workday/bsvc Uncollected_GTL,omitempty"`
}

// QTD Tax Remittance Data for Periodic Company
type QTDTaxRemittanceDataforPeriodicCompanyType struct {
	TaxWithheld       float64 `xml:"urn:com.workday/bsvc Tax_Withheld,omitempty"`
	TaxableWages      float64 `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages      float64 `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	ExemptWages       float64 `xml:"urn:com.workday/bsvc Exempt_Wages,omitempty"`
	WCBInsurableWages float64 `xml:"urn:com.workday/bsvc WCB_Insurable_Wages,omitempty"`
	WCBGrossWages     float64 `xml:"urn:com.workday/bsvc WCB_Gross_Wages,omitempty"`
	WCBOtherWages     float64 `xml:"urn:com.workday/bsvc WCB_Other_Wages,omitempty"`
	WCBExcessWages    float64 `xml:"urn:com.workday/bsvc WCB_Excess_Wages,omitempty"`
}

// Wrapper element for Quarterly Deduction Data.
type QuarterlyDeductionTaxFilingDataType struct {
	ThirdPartySickPay            *bool                                   `xml:"urn:com.workday/bsvc Third_Party_Sick_Pay,omitempty"`
	EmployerPaid                 *bool                                   `xml:"urn:com.workday/bsvc Employer_Paid,omitempty"`
	DeductionReference           *PayrollCalculationObjectType           `xml:"urn:com.workday/bsvc Deduction_Reference"`
	PayrollTaxAuthorityReference []MetadataPayrollWorktagObjectType      `xml:"urn:com.workday/bsvc Payroll_Tax_Authority_Reference"`
	WorkPSDCode                  string                                  `xml:"urn:com.workday/bsvc Work_PSD_Code,omitempty"`
	ResidentPSDCode              string                                  `xml:"urn:com.workday/bsvc Resident_PSD_Code,omitempty"`
	TaxFilingCode                string                                  `xml:"urn:com.workday/bsvc Tax_Filing_Code,omitempty"`
	SUIRate                      float64                                 `xml:"urn:com.workday/bsvc SUI_Rate,omitempty"`
	ActiveforMonth1              *bool                                   `xml:"urn:com.workday/bsvc Active_for_Month_1,omitempty"`
	Activefor12thDayMonth1       *bool                                   `xml:"urn:com.workday/bsvc Active_for_12th_Day_Month_1,omitempty"`
	ActiveforMonth2              *bool                                   `xml:"urn:com.workday/bsvc Active_for_Month_2,omitempty"`
	Activefor12thDayMonth2       *bool                                   `xml:"urn:com.workday/bsvc Active_for_12th_Day_Month_2,omitempty"`
	ActiveforMonth3              *bool                                   `xml:"urn:com.workday/bsvc Active_for_Month_3,omitempty"`
	Activefor12thDayMonth3       *bool                                   `xml:"urn:com.workday/bsvc Active_for_12th_Day_Month_3,omitempty"`
	EmployeeTaxElectionData      []WorkerTaxElectionsType                `xml:"urn:com.workday/bsvc Employee_Tax_Election_Data,omitempty"`
	QTDData                      *QTDTaxFilingDataforQuarterlyWorkerType `xml:"urn:com.workday/bsvc QTD_Data,omitempty"`
	YTDData                      *YTDTaxFilingDataforQuarterlyWorkerType `xml:"urn:com.workday/bsvc YTD_Data,omitempty"`
}

// Tax Category Data
type QuarterlyTaxFilingCategoryDataType struct {
	ThirdPartySickPay               *bool                       `xml:"urn:com.workday/bsvc Third_Party_Sick_Pay,omitempty"`
	TaxFilingCategoryReference      *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Tax_Filing_Category_Reference,omitempty"`
	TotalTaxFilingCategoryAmount    float64                     `xml:"urn:com.workday/bsvc Total_Tax_Filing_Category_Amount,omitempty"`
	YTDTotalTaxFilingCategoryAmount float64                     `xml:"urn:com.workday/bsvc YTD_Total_Tax_Filing_Category_Amount,omitempty"`
}

// Wrapper element for Tax Filing Category Data.
type QuarterlyTaxFilingConfigurationDataType struct {
	QuarterlyTaxCategoryData []QuarterlyTaxFilingCategoryDataType `xml:"urn:com.workday/bsvc Quarterly_Tax_Category_Data,omitempty"`
}

// Wrapper element for Tax Filing Category Data.
type QuarterlyTaxFilingConfigurationDataWrapperType struct {
	QuarterlyTaxConfigurationData []QuarterlyTaxFilingConfigurationDataType `xml:"urn:com.workday/bsvc Quarterly_Tax_Configuration_Data,omitempty"`
}

// Request Criteria
type QuarterlyWorkerTaxDataRequestCriteriaType struct {
	CompanyReference              *CompanyObjectType                 `xml:"urn:com.workday/bsvc Company_Reference"`
	CalendarQuarterReference      *CalendarQuarterObjectType         `xml:"urn:com.workday/bsvc Calendar_Quarter_Reference"`
	WorkerReference               []WorkerObjectType                 `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	FieldAndParameterCriteriaData *FieldAndParameterCriteriaDataType `xml:"urn:com.workday/bsvc Field_And_Parameter_Criteria_Data,omitempty"`
}

// Response Data
type QuarterlyWorkerTaxDataResponseDataType struct {
	QuarterlyWorkerTaxFilingData []QuarterlyWorkerTaxFilingDataType `xml:"urn:com.workday/bsvc Quarterly_Worker_Tax_Filing_Data,omitempty"`
}

// Response Group
type QuarterlyWorkerTaxDataResponseGroupType struct {
	IncludeYTDData                           *bool `xml:"urn:com.workday/bsvc Include_YTD_Data,omitempty"`
	ExcludeActiveFlagsforallTaxesExceptSUI   *bool `xml:"urn:com.workday/bsvc Exclude_Active_Flags_for_all_Taxes_Except_SUI,omitempty"`
	ExcludeTaxFilingConfigurationDataElement *bool `xml:"urn:com.workday/bsvc Exclude_Tax_Filing_Configuration_Data_Element,omitempty"`
	ExcludeWAIIFDataElement                  *bool `xml:"urn:com.workday/bsvc Exclude_WA_IIF_Data_Element,omitempty"`
	ExcludeAnnualDataElementsinQ4            *bool `xml:"urn:com.workday/bsvc Exclude_Annual_Data_Elements_in_Q4,omitempty"`
	ExcludePayrollTaxAuthorityDataElement    *bool `xml:"urn:com.workday/bsvc Exclude_Payroll_Tax_Authority_Data_Element,omitempty"`
}

// Quarterly Worker Tax Filing Data
type QuarterlyWorkerTaxFilingDataType struct {
	WorkerReference                         *WorkerObjectType                                `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	CompanyReference                        *CompanyObjectType                               `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	KindofEmployerReference                 *KindofEmployerforPayrollTaxFilingObjectType     `xml:"urn:com.workday/bsvc Kind_of_Employer_Reference,omitempty"`
	ThirdPartyProvidesSickPayW2             *bool                                            `xml:"urn:com.workday/bsvc Third_Party_Provides_Sick_Pay_W-2,omitempty"`
	CaliforniaSDIEnabledforCompany          *bool                                            `xml:"urn:com.workday/bsvc California_SDI_Enabled_for_Company,omitempty"`
	CaliforniaVDIEnabledforCompany          *bool                                            `xml:"urn:com.workday/bsvc California_VDI_Enabled_for_Company,omitempty"`
	CalendarQuarterReference                *CalendarQuarterObjectType                       `xml:"urn:com.workday/bsvc Calendar_Quarter_Reference,omitempty"`
	QuarterName                             string                                           `xml:"urn:com.workday/bsvc Quarter_Name,omitempty"`
	StartDate                               *time.Time                                       `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                                 *time.Time                                       `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	QuarterNumber                           float64                                          `xml:"urn:com.workday/bsvc Quarter_Number,omitempty"`
	CurrencyReference                       *CurrencyObjectType                              `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	WeeksWorked                             float64                                          `xml:"urn:com.workday/bsvc Weeks_Worked,omitempty"`
	HoursWorked                             float64                                          `xml:"urn:com.workday/bsvc Hours_Worked,omitempty"`
	HoursNotWorked                          float64                                          `xml:"urn:com.workday/bsvc Hours_Not_Worked,omitempty"`
	SeasonalWorker                          *bool                                            `xml:"urn:com.workday/bsvc Seasonal_Worker,omitempty"`
	MedicalCoverage                         *bool                                            `xml:"urn:com.workday/bsvc Medical_Coverage,omitempty"`
	OutsideCoverage                         *bool                                            `xml:"urn:com.workday/bsvc Outside_Coverage,omitempty"`
	CorporateOfficer                        *bool                                            `xml:"urn:com.workday/bsvc Corporate_Officer,omitempty"`
	WCCodeReference                         *WorkersCompensationCodeObjectType               `xml:"urn:com.workday/bsvc WC_Code_Reference,omitempty"`
	WCRate                                  float64                                          `xml:"urn:com.workday/bsvc WC_Rate,omitempty"`
	WCExempt                                *bool                                            `xml:"urn:com.workday/bsvc WC_Exempt,omitempty"`
	EEClassCodeIndicatorReference           *PayrollReportingCodeAllObjectType               `xml:"urn:com.workday/bsvc EE_Class_Code_Indicator_Reference,omitempty"`
	CorporateOfficerTypeReference           *PayrollReportingCodeAllObjectType               `xml:"urn:com.workday/bsvc Corporate_Officer_Type_Reference,omitempty"`
	CoverageTypeUnemploymentWorkersCompBoth string                                           `xml:"urn:com.workday/bsvc Coverage_Type__Unemployment_Workers_Comp_Both_,omitempty"`
	PayrollReportingCodeData                []WorkerPayrollReportingCodeDataType             `xml:"urn:com.workday/bsvc Payroll_Reporting_Code_Data,omitempty"`
	PayrollTaxAuthorityData                 []PayrollTaxAuthorityDataType                    `xml:"urn:com.workday/bsvc Payroll_Tax_Authority_Data,omitempty"`
	DocumentFieldResultData                 []DocumentFieldResultDataType                    `xml:"urn:com.workday/bsvc Document_Field_Result_Data,omitempty"`
	QuarterlyDeductionTaxData               []QuarterlyDeductionTaxFilingDataType            `xml:"urn:com.workday/bsvc Quarterly_Deduction_Tax_Data,omitempty"`
	WAIIFDeductionData                      []WAIIFDeductionDataType                         `xml:"urn:com.workday/bsvc WA_IIF_Deduction_Data,omitempty"`
	QuarterlyTaxConfigurationDataWrapper    []QuarterlyTaxFilingConfigurationDataWrapperType `xml:"urn:com.workday/bsvc Quarterly_Tax_Configuration_Data_Wrapper,omitempty"`
	WorkerAnnualTaxDataWrapper              []WorkerAnnualTaxDataWrapperType                 `xml:"urn:com.workday/bsvc Worker_Annual_Tax_Data_Wrapper,omitempty"`
	WorkerPuertoRicoAnnualTaxDataWrapper    []WorkerPuertoRicoAnnualTaxDataWrapperType       `xml:"urn:com.workday/bsvc Worker_Puerto_Rico_Annual_Tax_Data_Wrapper,omitempty"`
	WorkerGuamAnnualTaxDataWrapper          []WorkerGuamAnnualTaxDataWrapperType             `xml:"urn:com.workday/bsvc Worker_Guam_Annual_Tax_Data_Wrapper,omitempty"`
	WorkerVirginIslandsAnnualTaxDataWrapper []WorkerVirginIslandsAnnualTaxDataWrapperType    `xml:"urn:com.workday/bsvc Worker_Virgin_Islands_Annual_Tax_Data_Wrapper,omitempty"`
	CompanyProvidesDependentBenefits        *bool                                            `xml:"urn:com.workday/bsvc Company_Provides_Dependent_Benefits,omitempty"`
}

func (t *QuarterlyWorkerTaxFilingDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T QuarterlyWorkerTaxFilingDataType
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
func (t *QuarterlyWorkerTaxFilingDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T QuarterlyWorkerTaxFilingDataType
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

// ROE History Data Request Criteria Element
type ROEHistoryDataRequestCriteriaType struct {
}

// The Request References element contains the specific instance set that should be returned in the Get operation. Either a Reference ID or the Workday ID (GUID) should be specified for each instance to be returned.
type ROEHistoryDataRequestReferencesType struct {
	ROEHistoryDataReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc ROE_History_Data_Reference"`
}

// ROE History Data Response Data element
type ROEHistoryDataResponseDataType struct {
	ROEHistory []ROEHistoryType `xml:"urn:com.workday/bsvc ROE_History,omitempty"`
}

// The Response Group allows the request to specify which data attributes should be returned in the Response, such as whether to include reference elements, attachments, etc.
type ROEHistoryDataResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// ROE History Data element containing attributes needed
type ROEHistoryDataType struct {
	PayrollROEHistoryDataReference   *UniqueIdentifierObjectType       `xml:"urn:com.workday/bsvc Payroll_ROE_History_Data_Reference,omitempty"`
	CompanyReference                 *CompanyObjectType                `xml:"urn:com.workday/bsvc Company_Reference"`
	WorkerReference                  *WorkerObjectType                 `xml:"urn:com.workday/bsvc Worker_Reference"`
	PayrollReferenceNumberReference  *PayrollReferenceNumberObjectType `xml:"urn:com.workday/bsvc Payroll_Reference_Number_Reference"`
	DateIssued                       time.Time                         `xml:"urn:com.workday/bsvc Date_Issued"`
	SerialNumber                     string                            `xml:"urn:com.workday/bsvc Serial_Number,omitempty"`
	ReturnfromLeaveDatePriortoGoLive *time.Time                        `xml:"urn:com.workday/bsvc Return_from_Leave_Date__Prior_to_Go_Live_,omitempty"`
}

func (t *ROEHistoryDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ROEHistoryDataType
	var layout struct {
		*T
		DateIssued                       *xsdDate `xml:"urn:com.workday/bsvc Date_Issued"`
		ReturnfromLeaveDatePriortoGoLive *xsdDate `xml:"urn:com.workday/bsvc Return_from_Leave_Date__Prior_to_Go_Live_,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateIssued = (*xsdDate)(&layout.T.DateIssued)
	layout.ReturnfromLeaveDatePriortoGoLive = (*xsdDate)(layout.T.ReturnfromLeaveDatePriortoGoLive)
	return e.EncodeElement(layout, start)
}
func (t *ROEHistoryDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ROEHistoryDataType
	var overlay struct {
		*T
		DateIssued                       *xsdDate `xml:"urn:com.workday/bsvc Date_Issued"`
		ReturnfromLeaveDatePriortoGoLive *xsdDate `xml:"urn:com.workday/bsvc Return_from_Leave_Date__Prior_to_Go_Live_,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateIssued = (*xsdDate)(&overlay.T.DateIssued)
	overlay.ReturnfromLeaveDatePriortoGoLive = (*xsdDate)(overlay.T.ReturnfromLeaveDatePriortoGoLive)
	return d.DecodeElement(&overlay, &start)
}

// ROE History element containing Payroll ROE instance and necessary data for ROE History
type ROEHistoryType struct {
	PayrollROEDataReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Payroll_ROE_Data_Reference,omitempty"`
	ROEHistoryData          []ROEHistoryDataType        `xml:"urn:com.workday/bsvc ROE_History_Data,omitempty"`
}

// Vacation Pay Data - b17a
type ROEVacationPayDataType struct {
	VacationPayCode      string     `xml:"urn:com.workday/bsvc Vacation_Pay_Code,omitempty"`
	VacationPayStartDate *time.Time `xml:"urn:com.workday/bsvc Vacation_Pay_Start_Date,omitempty"`
	VacationPayEndDate   *time.Time `xml:"urn:com.workday/bsvc Vacation_Pay_End_Date,omitempty"`
	VacationPayAmount    float64    `xml:"urn:com.workday/bsvc Vacation_Pay_Amount,omitempty"`
}

func (t *ROEVacationPayDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ROEVacationPayDataType
	var layout struct {
		*T
		VacationPayStartDate *xsdDate `xml:"urn:com.workday/bsvc Vacation_Pay_Start_Date,omitempty"`
		VacationPayEndDate   *xsdDate `xml:"urn:com.workday/bsvc Vacation_Pay_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.VacationPayStartDate = (*xsdDate)(layout.T.VacationPayStartDate)
	layout.VacationPayEndDate = (*xsdDate)(layout.T.VacationPayEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ROEVacationPayDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ROEVacationPayDataType
	var overlay struct {
		*T
		VacationPayStartDate *xsdDate `xml:"urn:com.workday/bsvc Vacation_Pay_Start_Date,omitempty"`
		VacationPayEndDate   *xsdDate `xml:"urn:com.workday/bsvc Vacation_Pay_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.VacationPayStartDate = (*xsdDate)(overlay.T.VacationPayStartDate)
	overlay.VacationPayEndDate = (*xsdDate)(overlay.T.VacationPayEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Receiving Party Bank Data
type ReceivingPartyBankDataType struct {
	BankAccountNickname  string                               `xml:"urn:com.workday/bsvc Bank_Account_Nickname,omitempty"`
	AccountNumber        string                               `xml:"urn:com.workday/bsvc Account_Number,omitempty"`
	MaskedAccountNumber  string                               `xml:"urn:com.workday/bsvc Masked_Account_Number,omitempty"`
	AccountType          string                               `xml:"urn:com.workday/bsvc Account_Type,omitempty"`
	BankName             string                               `xml:"urn:com.workday/bsvc Bank_Name,omitempty"`
	IBAN                 string                               `xml:"urn:com.workday/bsvc IBAN,omitempty"`
	BankIDNumber         string                               `xml:"urn:com.workday/bsvc Bank_ID_Number,omitempty"`
	BIC                  string                               `xml:"urn:com.workday/bsvc BIC,omitempty"`
	BranchName           string                               `xml:"urn:com.workday/bsvc Branch_Name,omitempty"`
	BranchIDNumber       string                               `xml:"urn:com.workday/bsvc Branch_ID_Number,omitempty"`
	CountryReference     *CountryObjectType                   `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CurrencyReference    *InstanceObjectType                  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	BankAccountName      string                               `xml:"urn:com.workday/bsvc Bank_Account_Name,omitempty"`
	CheckDigit           string                               `xml:"urn:com.workday/bsvc Check_Digit,omitempty"`
	RollNumber           string                               `xml:"urn:com.workday/bsvc Roll_Number,omitempty"`
	BankInstructions     string                               `xml:"urn:com.workday/bsvc Bank_Instructions,omitempty"`
	IntermediaryBankData []IntermediaryBankAccountWWSDataType `xml:"urn:com.workday/bsvc Intermediary_Bank_Data,omitempty"`
}

// Receiving Party Data
type ReceivingPartyWWSDataType struct {
	PayeeName                        string                                `xml:"urn:com.workday/bsvc Payee_Name"`
	PayeeLegalName                   string                                `xml:"urn:com.workday/bsvc Payee_Legal_Name"`
	PayeePreferredName               string                                `xml:"urn:com.workday/bsvc Payee_Preferred_Name"`
	PayeeID                          string                                `xml:"urn:com.workday/bsvc Payee_ID,omitempty"`
	ExpensePayeeEmployeeNumber       string                                `xml:"urn:com.workday/bsvc Expense_Payee_Employee_Number,omitempty"`
	TaxID                            string                                `xml:"urn:com.workday/bsvc Tax_ID,omitempty"`
	TaxIDType                        string                                `xml:"urn:com.workday/bsvc Tax_ID_Type,omitempty"`
	CorporateCreditCardAccountNumber string                                `xml:"urn:com.workday/bsvc Corporate_Credit_Card_Account_Number,omitempty"`
	MarketCode                       string                                `xml:"urn:com.workday/bsvc Market_Code,omitempty"`
	CountryReference                 *CountryObjectType                    `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	PayeeAlternateName               []BusinessEntityAlternateNameDataType `xml:"urn:com.workday/bsvc Payee_Alternate_Name,omitempty"`
	EmailAddressData                 []EmailAddressDataType                `xml:"urn:com.workday/bsvc Email_Address_Data,omitempty"`
	AddressInformationData           []AddressInformationDataType          `xml:"urn:com.workday/bsvc Address_Information_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RegionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RegionObjectType struct {
	ID         []RegionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RelatedCalculationAllObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RelatedCalculationAllObjectType struct {
	ID         []RelatedCalculationAllObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Related Calculation Data for Get Payroll Results
type RelatedCalculationDataforGetPayrollResultsType struct {
	RelatedCalculationReference *RelatedCalculationAllObjectType `xml:"urn:com.workday/bsvc Related_Calculation_Reference,omitempty"`
	RelatedAmountValue          float64                          `xml:"urn:com.workday/bsvc Related_Amount_Value,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ReportOutputTagObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ReportOutputTagObjectType struct {
	ID         []ReportOutputTagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Data Element for Repository Document Metadata
type RepositoryDocumentMetadataType struct {
	FileName           string                      `xml:"urn:com.workday/bsvc File_Name,omitempty"`
	Createdby          string                      `xml:"urn:com.workday/bsvc Created_by,omitempty"`
	DateandTimeCreated *time.Time                  `xml:"urn:com.workday/bsvc Date_and_Time_Created,omitempty"`
	ReportTagReference []ReportOutputTagObjectType `xml:"urn:com.workday/bsvc Report_Tag_Reference,omitempty"`
}

func (t *RepositoryDocumentMetadataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RepositoryDocumentMetadataType
	var layout struct {
		*T
		DateandTimeCreated *xsdDateTime `xml:"urn:com.workday/bsvc Date_and_Time_Created,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateandTimeCreated = (*xsdDateTime)(layout.T.DateandTimeCreated)
	return e.EncodeElement(layout, start)
}
func (t *RepositoryDocumentMetadataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RepositoryDocumentMetadataType
	var overlay struct {
		*T
		DateandTimeCreated *xsdDateTime `xml:"urn:com.workday/bsvc Date_and_Time_Created,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateandTimeCreated = (*xsdDateTime)(overlay.T.DateandTimeCreated)
	return d.DecodeElement(&overlay, &start)
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

// Result Line Details for Get Payroll Results
type ResultLineDetailsforGetPayrollResultsType struct {
	PayComponentReference          []PayComponentReferenceType                      `xml:"urn:com.workday/bsvc Pay_Component_Reference,omitempty"`
	PayComponentFrequency          float64                                          `xml:"urn:com.workday/bsvc Pay_Component_Frequency,omitempty"`
	PeriodsRemainingInCalendarYear float64                                          `xml:"urn:com.workday/bsvc Periods_Remaining_In_Calendar_Year,omitempty"`
	SubperiodStartDate             *time.Time                                       `xml:"urn:com.workday/bsvc Subperiod_Start_Date,omitempty"`
	SubperiodEndDate               *time.Time                                       `xml:"urn:com.workday/bsvc Subperiod_End_Date,omitempty"`
	PayrollWorktagData             *PayrollWorktagDataType                          `xml:"urn:com.workday/bsvc Payroll_Worktag_Data,omitempty"`
	ResultLineAmount               float64                                          `xml:"urn:com.workday/bsvc Result_Line_Amount,omitempty"`
	ResultLineQTDAmount            float64                                          `xml:"urn:com.workday/bsvc Result_Line_QTD_Amount,omitempty"`
	ResultLineYTDAmount            float64                                          `xml:"urn:com.workday/bsvc Result_Line_YTD_Amount,omitempty"`
	ResultLineArrearsAmount        float64                                          `xml:"urn:com.workday/bsvc Result_Line_Arrears_Amount,omitempty"`
	RelatedCalculationResultData   []RelatedCalculationDataforGetPayrollResultsType `xml:"urn:com.workday/bsvc Related_Calculation_Result_Data,omitempty"`
	WithholdingOrderData           *WithholdingOrderDataforGetPayrollResultsType    `xml:"urn:com.workday/bsvc Withholding_Order_Data,omitempty"`
	PayrollInputComment            string                                           `xml:"urn:com.workday/bsvc Payroll_Input_Comment,omitempty"`
}

func (t *ResultLineDetailsforGetPayrollResultsType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ResultLineDetailsforGetPayrollResultsType
	var layout struct {
		*T
		SubperiodStartDate *xsdDate `xml:"urn:com.workday/bsvc Subperiod_Start_Date,omitempty"`
		SubperiodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Subperiod_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.SubperiodStartDate = (*xsdDate)(layout.T.SubperiodStartDate)
	layout.SubperiodEndDate = (*xsdDate)(layout.T.SubperiodEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ResultLineDetailsforGetPayrollResultsType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ResultLineDetailsforGetPayrollResultsType
	var overlay struct {
		*T
		SubperiodStartDate *xsdDate `xml:"urn:com.workday/bsvc Subperiod_Start_Date,omitempty"`
		SubperiodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Subperiod_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.SubperiodStartDate = (*xsdDate)(overlay.T.SubperiodStartDate)
	overlay.SubperiodEndDate = (*xsdDate)(overlay.T.SubperiodEndDate)
	return d.DecodeElement(&overlay, &start)
}

// This can be used to enter worktag overrides (if applicable) that will be respected at the 'result' level.  If no override is entered for a specific worktag, an applicable default worktag value (as of Period end date) will be used.
type ResultWorktagOverridesDataType struct {
	CompanyReference                *CompanyObjectType                             `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	RegionReference                 *RegionObjectType                              `xml:"urn:com.workday/bsvc Region_Reference,omitempty"`
	LocationReference               *LocationObjectType                            `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	CostCenterReference             *CostCenterObjectType                          `xml:"urn:com.workday/bsvc Cost_Center_Reference,omitempty"`
	JobProfileReference             *JobProfileObjectType                          `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	StateWorkReference              *PayrollStateAuthorityObjectType               `xml:"urn:com.workday/bsvc State__Work__Reference,omitempty"`
	StateResidentReference          *PayrollStateAuthorityObjectType               `xml:"urn:com.workday/bsvc State__Resident__Reference,omitempty"`
	CountyWorkReference             *PayrollLocalCountyAuthorityObjectType         `xml:"urn:com.workday/bsvc County__Work__Reference,omitempty"`
	CountyResidentReference         *PayrollLocalCountyAuthorityObjectType         `xml:"urn:com.workday/bsvc County__Resident__Reference,omitempty"`
	CityWorkReference               *PayrollLocalCityAuthorityObjectType           `xml:"urn:com.workday/bsvc City__Work__Reference,omitempty"`
	CityResidentReference           *PayrollLocalCityAuthorityObjectType           `xml:"urn:com.workday/bsvc City__Resident__Reference,omitempty"`
	SchoolDistrictResidentReference *PayrollLocalSchoolDistrictAuthorityObjectType `xml:"urn:com.workday/bsvc School_District__Resident__Reference,omitempty"`
	PayrollReferenceNumberReference *PayrollReferenceNumberObjectType              `xml:"urn:com.workday/bsvc Payroll_Reference_Number_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RunCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RunCategoryObjectType struct {
	ID         []RunCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The element to provide all the variables for Run Pay Calculation.
type RunPayCalculationDataType struct {
	PeriodReference                         *PeriodObjectType                        `xml:"urn:com.workday/bsvc Period_Reference,omitempty"`
	PayGrouporPayRunGroupSelectionReference []PayGroupPayRunGroupSelectionObjectType `xml:"urn:com.workday/bsvc Pay_Group_or_Pay_Run_Group_Selection_Reference"`
	RunCategoryReference                    *RunCategoryObjectType                   `xml:"urn:com.workday/bsvc Run_Category_Reference,omitempty"`
	SmartCalculationReference               *bool                                    `xml:"urn:com.workday/bsvc Smart_Calculation_Reference,omitempty"`
	BasedonEventsReference                  []PayrollEventCategoryObjectType         `xml:"urn:com.workday/bsvc Based_on_Events_Reference,omitempty"`
	IncludeRequiresReCalculation            *bool                                    `xml:"urn:com.workday/bsvc Include_Requires_ReCalculation,omitempty"`
	IncludeError                            *bool                                    `xml:"urn:com.workday/bsvc Include_Error,omitempty"`
	IncludeNotYetStarted                    *bool                                    `xml:"urn:com.workday/bsvc Include_Not_Yet_Started,omitempty"`
	IncludeInProgress                       *bool                                    `xml:"urn:com.workday/bsvc Include_In_Progress,omitempty"`
	AccountingOnlyforCalculationCriteria    *bool                                    `xml:"urn:com.workday/bsvc Accounting_Only_for_Calculation_Criteria,omitempty"`
}

// Run Pay Calculation Message is created when there is information returned while executing Run Payroll Calculation
type RunPayCalculationMessageType struct {
	DateandTime                *time.Time                          `xml:"urn:com.workday/bsvc Date_and_Time,omitempty"`
	Severity                   string                              `xml:"urn:com.workday/bsvc Severity,omitempty"`
	Message                    string                              `xml:"urn:com.workday/bsvc Message,omitempty"`
	BackgroundProcessReference *BackgroundProcessRuntimeObjectType `xml:"urn:com.workday/bsvc Background_Process_Reference,omitempty"`
}

func (t *RunPayCalculationMessageType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RunPayCalculationMessageType
	var layout struct {
		*T
		DateandTime *xsdDateTime `xml:"urn:com.workday/bsvc Date_and_Time,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateandTime = (*xsdDateTime)(layout.T.DateandTime)
	return e.EncodeElement(layout, start)
}
func (t *RunPayCalculationMessageType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RunPayCalculationMessageType
	var overlay struct {
		*T
		DateandTime *xsdDateTime `xml:"urn:com.workday/bsvc Date_and_Time,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateandTime = (*xsdDateTime)(overlay.T.DateandTime)
	return d.DecodeElement(&overlay, &start)
}

// Run Pay Calculation Messages contains all the messages returned while executing the Run Payroll Calculation task
type RunPayCalculationMessagesType struct {
	RunPayCalculationMessage []RunPayCalculationMessageType `xml:"urn:com.workday/bsvc Run_Pay_Calculation_Message,omitempty"`
}

// Calculation Criteria
type RunPayCalculationProcessCalculationCriteriaType struct {
	PeriodReference                                     *PeriodObjectType                `xml:"urn:com.workday/bsvc Period_Reference,omitempty"`
	PayRunGroupSelectionReference                       []PayRunGroupSelectionObjectType `xml:"urn:com.workday/bsvc Pay_Run_Group_Selection_Reference,omitempty"`
	CalculationOptions                                  []RunPayCalculationStatusesType  `xml:"urn:com.workday/bsvc Calculation_Options,omitempty"`
	LimitAccountingPopulationBasedonCalculationCriteria *bool                            `xml:"urn:com.workday/bsvc Limit_Accounting_Population_Based_on_Calculation_Criteria,omitempty"`
}

// Process History
type RunPayCalculationProcessHistoryType struct {
	ActualStartDateandTime        *time.Time `xml:"urn:com.workday/bsvc Actual_Start_Date_and_Time,omitempty"`
	CompletedDateandTime          *time.Time `xml:"urn:com.workday/bsvc Completed_Date_and_Time,omitempty"`
	TotalProcessingTimehourminsec string     `xml:"urn:com.workday/bsvc Total_Processing_Time__hour_min_sec_,omitempty"`
}

func (t *RunPayCalculationProcessHistoryType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RunPayCalculationProcessHistoryType
	var layout struct {
		*T
		ActualStartDateandTime *xsdDateTime `xml:"urn:com.workday/bsvc Actual_Start_Date_and_Time,omitempty"`
		CompletedDateandTime   *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Date_and_Time,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ActualStartDateandTime = (*xsdDateTime)(layout.T.ActualStartDateandTime)
	layout.CompletedDateandTime = (*xsdDateTime)(layout.T.CompletedDateandTime)
	return e.EncodeElement(layout, start)
}
func (t *RunPayCalculationProcessHistoryType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RunPayCalculationProcessHistoryType
	var overlay struct {
		*T
		ActualStartDateandTime *xsdDateTime `xml:"urn:com.workday/bsvc Actual_Start_Date_and_Time,omitempty"`
		CompletedDateandTime   *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Date_and_Time,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ActualStartDateandTime = (*xsdDateTime)(overlay.T.ActualStartDateandTime)
	overlay.CompletedDateandTime = (*xsdDateTime)(overlay.T.CompletedDateandTime)
	return d.DecodeElement(&overlay, &start)
}

// Process Information
type RunPayCalculationProcessInformationType struct {
	ProcessTypeReference              *BackgroundProcessTypeObjectType        `xml:"urn:com.workday/bsvc Process_Type_Reference,omitempty"`
	ScheduledStartDateandTime         *time.Time                              `xml:"urn:com.workday/bsvc Scheduled_Start_Date_and_Time,omitempty"`
	ActualStartDateandTime            *time.Time                              `xml:"urn:com.workday/bsvc Actual_Start_Date_and_Time,omitempty"`
	CompletedDateandTime              *time.Time                              `xml:"urn:com.workday/bsvc Completed_Date_and_Time,omitempty"`
	SubmittedByReference              *InstanceObjectType                     `xml:"urn:com.workday/bsvc Submitted_By_Reference,omitempty"`
	TotalProcessingTime               string                                  `xml:"urn:com.workday/bsvc Total_Processing_Time,omitempty"`
	JobDefinitionReference            *UniqueIdentifierObjectType             `xml:"urn:com.workday/bsvc Job_Definition_Reference,omitempty"`
	JobModeReference                  *UniqueIdentifierObjectType             `xml:"urn:com.workday/bsvc Job_Mode_Reference,omitempty"`
	JobSize                           float64                                 `xml:"urn:com.workday/bsvc Job_Size,omitempty"`
	JobParameters                     float64                                 `xml:"urn:com.workday/bsvc Job_Parameters,omitempty"`
	ApplicationKeys                   float64                                 `xml:"urn:com.workday/bsvc Application_Keys,omitempty"`
	RunPayCalculationTechnicalDetails []RunPayCalculationTechnicalDetailsType `xml:"urn:com.workday/bsvc Run_Pay_Calculation_Technical_Details,omitempty"`
	RunPayCalculationStepDetails      []RunPayCalculationStepDetailsType      `xml:"urn:com.workday/bsvc Run_Pay_Calculation_Step_Details,omitempty"`
}

func (t *RunPayCalculationProcessInformationType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RunPayCalculationProcessInformationType
	var layout struct {
		*T
		ScheduledStartDateandTime *xsdDateTime `xml:"urn:com.workday/bsvc Scheduled_Start_Date_and_Time,omitempty"`
		ActualStartDateandTime    *xsdDateTime `xml:"urn:com.workday/bsvc Actual_Start_Date_and_Time,omitempty"`
		CompletedDateandTime      *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Date_and_Time,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ScheduledStartDateandTime = (*xsdDateTime)(layout.T.ScheduledStartDateandTime)
	layout.ActualStartDateandTime = (*xsdDateTime)(layout.T.ActualStartDateandTime)
	layout.CompletedDateandTime = (*xsdDateTime)(layout.T.CompletedDateandTime)
	return e.EncodeElement(layout, start)
}
func (t *RunPayCalculationProcessInformationType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RunPayCalculationProcessInformationType
	var overlay struct {
		*T
		ScheduledStartDateandTime *xsdDateTime `xml:"urn:com.workday/bsvc Scheduled_Start_Date_and_Time,omitempty"`
		ActualStartDateandTime    *xsdDateTime `xml:"urn:com.workday/bsvc Actual_Start_Date_and_Time,omitempty"`
		CompletedDateandTime      *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Date_and_Time,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ScheduledStartDateandTime = (*xsdDateTime)(overlay.T.ScheduledStartDateandTime)
	overlay.ActualStartDateandTime = (*xsdDateTime)(overlay.T.ActualStartDateandTime)
	overlay.CompletedDateandTime = (*xsdDateTime)(overlay.T.CompletedDateandTime)
	return d.DecodeElement(&overlay, &start)
}

// Criteria used to identify calculation process if Process Integration ID is not used
type RunPayCalculationRequestCriteriaType struct {
	PeriodReference                           *PeriodObjectType                        `xml:"urn:com.workday/bsvc Period_Reference,omitempty"`
	PayGroupsorPayRunGroupsSelectionReference []PayGroupPayRunGroupSelectionObjectType `xml:"urn:com.workday/bsvc Pay_Groups_or_Pay_Run_Groups_Selection_Reference"`
	RunCategoryReference                      *RunCategoryObjectType                   `xml:"urn:com.workday/bsvc Run_Category_Reference,omitempty"`
	IncludeRequiresRecalculation              *bool                                    `xml:"urn:com.workday/bsvc Include_Requires_Recalculation,omitempty"`
	IncludeError                              *bool                                    `xml:"urn:com.workday/bsvc Include_Error,omitempty"`
	IncludeNotYetStarted                      *bool                                    `xml:"urn:com.workday/bsvc Include_Not_Yet_Started,omitempty"`
	IncludeInProgress                         *bool                                    `xml:"urn:com.workday/bsvc Include_In_Progress,omitempty"`
	AccountingOnlyforCalculationCriteria      *bool                                    `xml:"urn:com.workday/bsvc Accounting_Only_for_Calculation_Criteria,omitempty"`
	IncludeQueuedProesses                     *bool                                    `xml:"urn:com.workday/bsvc Include_Queued_Proesses,omitempty"`
}

// Pay Calculation Request References
type RunPayCalculationRequestReferencesType struct {
	RunPayCalculationReference []BackgroundProcessRuntimeObjectType `xml:"urn:com.workday/bsvc Run_Pay_Calculation_Reference"`
}

// Pay Calculation Response
type RunPayCalculationResponseDataType struct {
	RunPayCalculation []RunPayCalculationType `xml:"urn:com.workday/bsvc Run_Pay_Calculation,omitempty"`
}

// Used to specify what groups of information are provided in the response
type RunPayCalculationResponseGroupType struct {
	IncludeProcessCalculationCriteria *bool `xml:"urn:com.workday/bsvc Include_Process_Calculation_Criteria,omitempty"`
	IncludePayGroupsBeingCalculated   *bool `xml:"urn:com.workday/bsvc Include_Pay_Groups_Being_Calculated,omitempty"`
	IncludeProcessInformation         *bool `xml:"urn:com.workday/bsvc Include_Process_Information,omitempty"`
	IncludeProcessHistory             *bool `xml:"urn:com.workday/bsvc Include_Process_History,omitempty"`
	IncludeMessages                   *bool `xml:"urn:com.workday/bsvc Include_Messages,omitempty"`
}

// Statuses for Payroll Run
type RunPayCalculationStatusesType struct {
	PayCalculationStatusReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Pay_Calculation_Status_Reference,omitempty"`
}

// Process Information
type RunPayCalculationStepDetailType struct {
	JobStepRuntimeReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Job_Step_Runtime_Reference,omitempty"`
	StatusReference         *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Status_Reference,omitempty"`
	StartDateTime           *time.Time                  `xml:"urn:com.workday/bsvc Start_Date_Time,omitempty"`
	EndDateTime             *time.Time                  `xml:"urn:com.workday/bsvc End_Date_Time,omitempty"`
	CumulativeStepTimemin   float64                     `xml:"urn:com.workday/bsvc Cumulative_Step_Time__min_,omitempty"`
	ResultCount             float64                     `xml:"urn:com.workday/bsvc Result_Count,omitempty"`
	PartitionsCompleted     float64                     `xml:"urn:com.workday/bsvc Partitions_Completed,omitempty"`
	TotalPartitions         float64                     `xml:"urn:com.workday/bsvc Total_Partitions,omitempty"`
	RuntimeParameters       float64                     `xml:"urn:com.workday/bsvc Runtime_Parameters,omitempty"`
}

func (t *RunPayCalculationStepDetailType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RunPayCalculationStepDetailType
	var layout struct {
		*T
		StartDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Start_Date_Time,omitempty"`
		EndDateTime   *xsdDateTime `xml:"urn:com.workday/bsvc End_Date_Time,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDateTime = (*xsdDateTime)(layout.T.StartDateTime)
	layout.EndDateTime = (*xsdDateTime)(layout.T.EndDateTime)
	return e.EncodeElement(layout, start)
}
func (t *RunPayCalculationStepDetailType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RunPayCalculationStepDetailType
	var overlay struct {
		*T
		StartDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Start_Date_Time,omitempty"`
		EndDateTime   *xsdDateTime `xml:"urn:com.workday/bsvc End_Date_Time,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDateTime = (*xsdDateTime)(overlay.T.StartDateTime)
	overlay.EndDateTime = (*xsdDateTime)(overlay.T.EndDateTime)
	return d.DecodeElement(&overlay, &start)
}

// Step Details
type RunPayCalculationStepDetailsType struct {
	RunPayCalculationStepDetail []RunPayCalculationStepDetailType `xml:"urn:com.workday/bsvc Run_Pay_Calculation_Step_Detail,omitempty"`
}

// Technical Details
type RunPayCalculationTechnicalDetailsType struct {
	WorkdayIntegrationCloudPlatformESBProcessID string `xml:"urn:com.workday/bsvc Workday_Integration_Cloud_Platform__ESB__Process_ID,omitempty"`
}

// Pay Calculation Response Data
type RunPayCalculationType struct {
	RunPayCalculationReference          *BackgroundProcessRuntimeObjectType              `xml:"urn:com.workday/bsvc Run_Pay_Calculation_Reference,omitempty"`
	ProcessReference                    *BackgroundProcessObjectType                     `xml:"urn:com.workday/bsvc Process_Reference,omitempty"`
	StatusReference                     *BackgroundProcessRuntimeStatusObjectType        `xml:"urn:com.workday/bsvc Status_Reference,omitempty"`
	CurrentProcessingTime               string                                           `xml:"urn:com.workday/bsvc Current_Processing_Time,omitempty"`
	AverageProcessingTime               string                                           `xml:"urn:com.workday/bsvc Average_Processing_Time,omitempty"`
	RunPayCalculationProcessCriteria    *RunPayCalculationProcessCalculationCriteriaType `xml:"urn:com.workday/bsvc Run_Pay_Calculation_Process_Criteria,omitempty"`
	PayGroupsBeingCalculated            *PayGroupsBeingCalculatedType                    `xml:"urn:com.workday/bsvc Pay_Groups_Being_Calculated,omitempty"`
	RunPayCalculationProcessInformation *RunPayCalculationProcessInformationType         `xml:"urn:com.workday/bsvc Run_Pay_Calculation_Process_Information,omitempty"`
	RunPayCalculationProcessHistory     *RunPayCalculationProcessHistoryType             `xml:"urn:com.workday/bsvc Run_Pay_Calculation_Process_History,omitempty"`
	RunPayCalculationMessages           *RunPayCalculationMessagesType                   `xml:"urn:com.workday/bsvc Run_Pay_Calculation_Messages,omitempty"`
}

// Details of the Salary over the Cap allocation, (e.g., a set of allocation dimensions and percentages)
type SalaryOvertheCapCostingAllocationDetailDataType struct {
	SOCOrder                             string                             `xml:"urn:com.workday/bsvc SOC_Order"`
	SOCDefaultfromOrganizationAssignment bool                               `xml:"urn:com.workday/bsvc SOC_Default_from_Organization_Assignment"`
	SOCOverrideWorktagReference          []TenantedPayrollWorktagObjectType `xml:"urn:com.workday/bsvc SOC_Override_Worktag_Reference"`
	SOCDistributionPercent               float64                            `xml:"urn:com.workday/bsvc SOC_Distribution_Percent"`
}

// Settlement Account information including Bank Account Type, Currency, Routing Transit Number, etc.
//
// Bank Account Type becomes required when using v30.1 of the web service.
type SettlementAccountWWSDataType struct {
	SettlementBankAccountID      string                               `xml:"urn:com.workday/bsvc Settlement_Bank_Account_ID,omitempty"`
	CountryReference             *CountryObjectType                   `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CurrencyReference            *CurrencyObjectType                  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	BankAccountNickname          string                               `xml:"urn:com.workday/bsvc Bank_Account_Nickname,omitempty"`
	BankAccountTypeReference     *BankAccountTypeObjectType           `xml:"urn:com.workday/bsvc Bank_Account_Type_Reference,omitempty"`
	BankName                     string                               `xml:"urn:com.workday/bsvc Bank_Name,omitempty"`
	RoutingTransitNumber         string                               `xml:"urn:com.workday/bsvc Routing_Transit_Number,omitempty"`
	BranchID                     string                               `xml:"urn:com.workday/bsvc Branch_ID,omitempty"`
	BranchName                   string                               `xml:"urn:com.workday/bsvc Branch_Name,omitempty"`
	BankAccountNumber            string                               `xml:"urn:com.workday/bsvc Bank_Account_Number,omitempty"`
	CheckDigit                   string                               `xml:"urn:com.workday/bsvc Check_Digit,omitempty"`
	BankAccountName              string                               `xml:"urn:com.workday/bsvc Bank_Account_Name,omitempty"`
	RollNumber                   string                               `xml:"urn:com.workday/bsvc Roll_Number,omitempty"`
	IBAN                         string                               `xml:"urn:com.workday/bsvc IBAN,omitempty"`
	SWIFTBankIdentificationCode  string                               `xml:"urn:com.workday/bsvc SWIFT_Bank_Identification_Code,omitempty"`
	AcceptsPaymentTypesReference []PaymentTypeObjectType              `xml:"urn:com.workday/bsvc Accepts_Payment_Types_Reference,omitempty"`
	PaymentTypesReference        []PaymentTypeObjectType              `xml:"urn:com.workday/bsvc Payment_Types_Reference,omitempty"`
	ForSupplierConnectionsOnly   *bool                                `xml:"urn:com.workday/bsvc For_Supplier_Connections_Only,omitempty"`
	RequiresPrenote              *bool                                `xml:"urn:com.workday/bsvc Requires_Prenote,omitempty"`
	PaymentTypePrenoteReference  *PaymentTypeObjectType               `xml:"urn:com.workday/bsvc Payment_Type_Prenote_Reference,omitempty"`
	Inactive                     *bool                                `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	BankInstructions             string                               `xml:"urn:com.workday/bsvc Bank_Instructions,omitempty"`
	IntermediaryBankData         []IntermediaryBankAccountWWSDataType `xml:"urn:com.workday/bsvc Intermediary_Bank_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SettlementBankAccountObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type SettlementBankAccountObjectType struct {
	ID         []SettlementBankAccountObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SettlementInstructionObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type SettlementInstructionObjectType struct {
	ID         []SettlementInstructionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Single Legal Entity Data
type SingleLegalEntityDataType struct {
	ReportingTaxYearReference          *CalendarYearObjectType          `xml:"urn:com.workday/bsvc Reporting_Tax_Year_Reference"`
	InactiveasofTaxYearReference       *CalendarYearObjectType          `xml:"urn:com.workday/bsvc Inactive_as-of_Tax_Year_Reference,omitempty"`
	LegalEntityReference               *CompanyObjectType               `xml:"urn:com.workday/bsvc Legal_Entity_Reference"`
	RelatedCompaniesReference          []CompanyObjectType              `xml:"urn:com.workday/bsvc Related_Companies_Reference"`
	WithholdingOrderTypeReference      []WithholdingOrderTypeObjectType `xml:"urn:com.workday/bsvc Withholding_Order_Type_Reference,omitempty"`
	PeriodicCombinedTaxFiling          *bool                            `xml:"urn:com.workday/bsvc Periodic_Combined_Tax_Filing,omitempty"`
	QuarterlySeparateTaxFiling         *bool                            `xml:"urn:com.workday/bsvc Quarterly_Separate_Tax_Filing,omitempty"`
	CombinedACAReporting               *bool                            `xml:"urn:com.workday/bsvc Combined_ACA_Reporting,omitempty"`
	FederalTaxID                       string                           `xml:"urn:com.workday/bsvc Federal_Tax_ID,attr,omitempty"`
	SingleLegalEntityType              string                           `xml:"urn:com.workday/bsvc Single_Legal_Entity_Type,attr,omitempty"`
	FinancialTaxReporting              string                           `xml:"urn:com.workday/bsvc Financial_Tax_Reporting,attr,omitempty"`
	CombinedFinancialYEforallcompanies string                           `xml:"urn:com.workday/bsvc Combined_Financial_YE_for_all_companies,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SingleLegalEntityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SingleLegalEntityObjectType struct {
	ID         []SingleLegalEntityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Single Legal Entity Request References
type SingleLegalEntityRequestReferencesType struct {
	SingleLegalEntityReference []SingleLegalEntityObjectType `xml:"urn:com.workday/bsvc Single_Legal_Entity_Reference,omitempty"`
}

// Single Legal Entity Response Data
type SingleLegalEntityResponseDataType struct {
	SingleLegalEntity []SingleLegalEntityType `xml:"urn:com.workday/bsvc Single_Legal_Entity,omitempty"`
}

// Single Legal Entity Response Group
type SingleLegalEntityResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Single Legal Entity
type SingleLegalEntityType struct {
	SingleLegalEntityReference *SingleLegalEntityObjectType `xml:"urn:com.workday/bsvc Single_Legal_Entity_Reference,omitempty"`
	SingleLegalEntityData      *SingleLegalEntityDataType   `xml:"urn:com.workday/bsvc Single_Legal_Entity_Data,omitempty"`
}

// Payroll State and Local Tax Elections
type StateTaxElectionDataType struct {
	PayrollStateAuthorityReference    *PayrollStateAuthorityObjectType                         `xml:"urn:com.workday/bsvc Payroll_State_Authority_Reference,omitempty"`
	StateIncomeTaxData                []PayrollPayeeStateElectionforStateandLocalType          `xml:"urn:com.workday/bsvc State_Income_Tax_Data,omitempty"`
	StateUnemploymentData             *PayrollPayeeSUTAElectonforStateandLocalType             `xml:"urn:com.workday/bsvc State_Unemployment_Data,omitempty"`
	EarnedIncomeCreditData            *PayrollPayeeEICElectionforStateandLocalType             `xml:"urn:com.workday/bsvc Earned_Income_Credit_Data,omitempty"`
	PayrollStateCountyTaxData         []PayrollPayeeCountyElectionforStateandLocalType         `xml:"urn:com.workday/bsvc Payroll_State_County_Tax_Data,omitempty"`
	PayrollStateCityTaxData           []PayrollPayeeCityElectionforStateandLocalType           `xml:"urn:com.workday/bsvc Payroll_State_City_Tax_Data,omitempty"`
	PayrollStateSchoolDistrictTaxData []PayrollPayeeSchoolDistrictElectionforStateandLocalType `xml:"urn:com.workday/bsvc Payroll_State_School_District_Tax_Data,omitempty"`
	PayrollStateOtherTaxData          []PayrollPayeeOtherElectionforStateandLocalType          `xml:"urn:com.workday/bsvc Payroll_State_Other_Tax_Data,omitempty"`
}

// Federal Data used in State Tax Levy Withholding Orders
type StateTaxLevyFederalDataType struct {
	Part3EffectiveDate               *time.Time                         `xml:"urn:com.workday/bsvc Part_3_Effective_Date,omitempty"`
	PayPeriodExemptionOverrideAmount float64                            `xml:"urn:com.workday/bsvc Pay_Period_Exemption_Override_Amount,omitempty"`
	PayrollMaritalStatusReference    *PayrollMaritalStatusReferenceType `xml:"urn:com.workday/bsvc Payroll_Marital_Status_Reference,omitempty"`
	PersonalExemptions               float64                            `xml:"urn:com.workday/bsvc Personal_Exemptions,omitempty"`
	Additional65orBlindExemptions    float64                            `xml:"urn:com.workday/bsvc Additional_65__or_Blind_Exemptions,omitempty"`
	TerminationDate                  *time.Time                         `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
}

func (t *StateTaxLevyFederalDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T StateTaxLevyFederalDataType
	var layout struct {
		*T
		Part3EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Part_3_Effective_Date,omitempty"`
		TerminationDate    *xsdDate `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Part3EffectiveDate = (*xsdDate)(layout.T.Part3EffectiveDate)
	layout.TerminationDate = (*xsdDate)(layout.T.TerminationDate)
	return e.EncodeElement(layout, start)
}
func (t *StateTaxLevyFederalDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T StateTaxLevyFederalDataType
	var overlay struct {
		*T
		Part3EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Part_3_Effective_Date,omitempty"`
		TerminationDate    *xsdDate `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Part3EffectiveDate = (*xsdDate)(overlay.T.Part3EffectiveDate)
	overlay.TerminationDate = (*xsdDate)(overlay.T.TerminationDate)
	return d.DecodeElement(&overlay, &start)
}

// State Tax Levy Order Data
type StateTaxLevyOrderDataType struct {
	NumberofDependents                  float64                                               `xml:"urn:com.workday/bsvc Number_of_Dependents,omitempty"`
	PayrollLocalCountyAuthorityFIPSCode string                                                `xml:"urn:com.workday/bsvc Payroll_Local_County_Authority_FIPS_Code,omitempty"`
	WorkerisLaborerorMechanic           *bool                                                 `xml:"urn:com.workday/bsvc Worker_is_Laborer_or_Mechanic,omitempty"`
	WorkerIncomeisPovertyLevel          *bool                                                 `xml:"urn:com.workday/bsvc Worker_Income_is_Poverty_Level,omitempty"`
	StateTaxLevyFederalData             *StateTaxLevyFederalDataType                          `xml:"urn:com.workday/bsvc State_Tax_Levy_Federal_Data,omitempty"`
	StateTaxLevyDependantReference      []FederalTaxLevyDependentReferenceType                `xml:"urn:com.workday/bsvc State_Tax_Levy_Dependant_Reference,omitempty"`
	GoodCauseLimitPercent               float64                                               `xml:"urn:com.workday/bsvc Good_Cause_Limit_Percent,omitempty"`
	ProcessUntilReference               *WithholdingOrderOverrideCompletionCriteriaObjectType `xml:"urn:com.workday/bsvc Process_Until_Reference,omitempty"`
	ProrateUntilDate                    *time.Time                                            `xml:"urn:com.workday/bsvc Prorate_Until_Date,omitempty"`
}

func (t *StateTaxLevyOrderDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T StateTaxLevyOrderDataType
	var layout struct {
		*T
		ProrateUntilDate *xsdDate `xml:"urn:com.workday/bsvc Prorate_Until_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ProrateUntilDate = (*xsdDate)(layout.T.ProrateUntilDate)
	return e.EncodeElement(layout, start)
}
func (t *StateTaxLevyOrderDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T StateTaxLevyOrderDataType
	var overlay struct {
		*T
		ProrateUntilDate *xsdDate `xml:"urn:com.workday/bsvc Prorate_Until_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ProrateUntilDate = (*xsdDate)(overlay.T.ProrateUntilDate)
	return d.DecodeElement(&overlay, &start)
}

// Submit Payroll Input Data
type SubmitPayrollInputDataType struct {
	PayrollInputID             string                              `xml:"urn:com.workday/bsvc Payroll_Input_ID,omitempty"`
	BatchID                    string                              `xml:"urn:com.workday/bsvc Batch_ID,omitempty"`
	SourceReference            *IntegrationSystemAuditedObjectType `xml:"urn:com.workday/bsvc Source_Reference,omitempty"`
	OngoingInput               *bool                               `xml:"urn:com.workday/bsvc Ongoing_Input,omitempty"`
	StartDate                  time.Time                           `xml:"urn:com.workday/bsvc Start_Date"`
	EndDate                    *time.Time                          `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	RunCategoryReference       []RunCategoryObjectType             `xml:"urn:com.workday/bsvc Run_Category_Reference,omitempty"`
	WorkerReference            *WorkerObjectType                   `xml:"urn:com.workday/bsvc Worker_Reference"`
	PositionReference          *PositionElementObjectType          `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	EarningReference           *EarningAllObjectType               `xml:"urn:com.workday/bsvc Earning_Reference"`
	DeductionReference         *DeductionAllObjectType             `xml:"urn:com.workday/bsvc Deduction_Reference"`
	Amount                     float64                             `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Hours                      float64                             `xml:"urn:com.workday/bsvc Hours,omitempty"`
	Rate                       float64                             `xml:"urn:com.workday/bsvc Rate,omitempty"`
	Adjustment                 *bool                               `xml:"urn:com.workday/bsvc Adjustment,omitempty"`
	WorktagData                *PayrollInputWorktagsDataType       `xml:"urn:com.workday/bsvc Worktag_Data,omitempty"`
	AdditionalInputDetailsData []AdditionalInputDetailsType        `xml:"urn:com.workday/bsvc Additional_Input_Details_Data,omitempty"`
	Comment                    string                              `xml:"urn:com.workday/bsvc Comment,omitempty"`
	CurrencyReference          *CurrencyObjectType                 `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	MatchExistingInput         *bool                               `xml:"urn:com.workday/bsvc Match_Existing_Input,omitempty"`
	CompanyReference           *CompanyObjectType                  `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	CoverageStartDate          *time.Time                          `xml:"urn:com.workday/bsvc Coverage_Start_Date,omitempty"`
	CoverageEndDate            *time.Time                          `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	LastPeriodEndDate          time.Time                           `xml:"urn:com.workday/bsvc Last_Period_End_Date,attr,omitempty"`
}

func (t *SubmitPayrollInputDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T SubmitPayrollInputDataType
	var layout struct {
		*T
		StartDate         *xsdDate `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate           *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
		CoverageStartDate *xsdDate `xml:"urn:com.workday/bsvc Coverage_Start_Date,omitempty"`
		CoverageEndDate   *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
		LastPeriodEndDate *xsdDate `xml:"urn:com.workday/bsvc Last_Period_End_Date,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(&layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	layout.CoverageStartDate = (*xsdDate)(layout.T.CoverageStartDate)
	layout.CoverageEndDate = (*xsdDate)(layout.T.CoverageEndDate)
	layout.LastPeriodEndDate = (*xsdDate)(&layout.T.LastPeriodEndDate)
	return e.EncodeElement(layout, start)
}
func (t *SubmitPayrollInputDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SubmitPayrollInputDataType
	var overlay struct {
		*T
		StartDate         *xsdDate `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate           *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
		CoverageStartDate *xsdDate `xml:"urn:com.workday/bsvc Coverage_Start_Date,omitempty"`
		CoverageEndDate   *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
		LastPeriodEndDate *xsdDate `xml:"urn:com.workday/bsvc Last_Period_End_Date,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(&overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	overlay.CoverageStartDate = (*xsdDate)(overlay.T.CoverageStartDate)
	overlay.CoverageEndDate = (*xsdDate)(overlay.T.CoverageEndDate)
	overlay.LastPeriodEndDate = (*xsdDate)(&overlay.T.LastPeriodEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Submit Payroll Input Request
type SubmitPayrollInputRequestType struct {
	PayrollInputData []SubmitPayrollInputDataType `xml:"urn:com.workday/bsvc Payroll_Input_Data"`
	Version          string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Submit Payroll Inputs Response
type SubmitPayrollInputResponseDataType struct {
	PayrollInput []SubmitPayrollInputType `xml:"urn:com.workday/bsvc Payroll_Input,omitempty"`
}

// Submit Payroll Input Response
type SubmitPayrollInputResponseType struct {
	PayrollInputReference []PayrollInputObjectType `xml:"urn:com.workday/bsvc Payroll_Input_Reference,omitempty"`
	Version               string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Submit Payroll Input
type SubmitPayrollInputType struct {
	PayrollInputReference *PayrollInputObjectType      `xml:"urn:com.workday/bsvc Payroll_Input_Reference,omitempty"`
	PayrollInputData      []SubmitPayrollInputDataType `xml:"urn:com.workday/bsvc Payroll_Input_Data,omitempty"`
}

// The submunicipality of the address.
type SubmunicipalityInformationDataType struct {
	Value                string `xml:",chardata"`
	AddressComponentName string `xml:"urn:com.workday/bsvc Address_Component_Name,attr,omitempty"`
	Type                 string `xml:"urn:com.workday/bsvc Type,attr,omitempty"`
}

// The subregion part of the address.
type SubregionInformationDataType struct {
	Value      string `xml:",chardata"`
	Descriptor string `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
	Type       string `xml:"urn:com.workday/bsvc Type,attr,omitempty"`
}

// Successor Employer Data
type SuccessorEmployerDataType struct {
	SuccessorEmployerTypeReference                           *CompanyRelationshipTypeObjectType `xml:"urn:com.workday/bsvc Successor_Employer_Type_Reference"`
	AcquisitionorMergerDate                                  time.Time                          `xml:"urn:com.workday/bsvc Acquisition_or_Merger_Date"`
	PayrollProcessingBeginDate                               time.Time                          `xml:"urn:com.workday/bsvc Payroll_Processing_Begin_Date"`
	SuccessorCompanyReference                                *CompanyObjectType                 `xml:"urn:com.workday/bsvc Successor_Company_Reference"`
	PredecessorCompanyReference                              []CompanyObjectType                `xml:"urn:com.workday/bsvc Predecessor_Company_Reference"`
	WithholdingOrderTypeReference                            []WithholdingOrderTypeObjectType   `xml:"urn:com.workday/bsvc Withholding_Order_Type_Reference,omitempty"`
	TaxAuthoritiesnotRecognizingPredecessorPaymentsReference []PayrollTaxAuthorityObjectType    `xml:"urn:com.workday/bsvc Tax_Authorities_not_Recognizing_Predecessor_Payments_Reference,omitempty"`
	ID                                                       string                             `xml:"urn:com.workday/bsvc ID,attr,omitempty"`
	FederalTaxID                                             string                             `xml:"urn:com.workday/bsvc Federal_Tax_ID,attr,omitempty"`
}

func (t *SuccessorEmployerDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T SuccessorEmployerDataType
	var layout struct {
		*T
		AcquisitionorMergerDate    *xsdDate `xml:"urn:com.workday/bsvc Acquisition_or_Merger_Date"`
		PayrollProcessingBeginDate *xsdDate `xml:"urn:com.workday/bsvc Payroll_Processing_Begin_Date"`
	}
	layout.T = (*T)(t)
	layout.AcquisitionorMergerDate = (*xsdDate)(&layout.T.AcquisitionorMergerDate)
	layout.PayrollProcessingBeginDate = (*xsdDate)(&layout.T.PayrollProcessingBeginDate)
	return e.EncodeElement(layout, start)
}
func (t *SuccessorEmployerDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SuccessorEmployerDataType
	var overlay struct {
		*T
		AcquisitionorMergerDate    *xsdDate `xml:"urn:com.workday/bsvc Acquisition_or_Merger_Date"`
		PayrollProcessingBeginDate *xsdDate `xml:"urn:com.workday/bsvc Payroll_Processing_Begin_Date"`
	}
	overlay.T = (*T)(t)
	overlay.AcquisitionorMergerDate = (*xsdDate)(&overlay.T.AcquisitionorMergerDate)
	overlay.PayrollProcessingBeginDate = (*xsdDate)(&overlay.T.PayrollProcessingBeginDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type SuccessorEmployerObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SuccessorEmployerObjectType struct {
	ID         []SuccessorEmployerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Successor Employer Request References
type SuccessorEmployerRequestReferencesType struct {
	SuccessorEmployerReference []SuccessorEmployerObjectType `xml:"urn:com.workday/bsvc Successor_Employer_Reference"`
}

// Successor Employer Response Data
type SuccessorEmployerResponseDataType struct {
	SuccessorEmployer []SuccessorEmployerType `xml:"urn:com.workday/bsvc Successor_Employer,omitempty"`
}

// Successor Employer Response Group
type SuccessorEmployerResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Successor Employer
type SuccessorEmployerType struct {
	SuccessorEmployerReference *SuccessorEmployerObjectType `xml:"urn:com.workday/bsvc Successor_Employer_Reference,omitempty"`
	SuccessorEmployerData      *SuccessorEmployerDataType   `xml:"urn:com.workday/bsvc Successor_Employer_Data,omitempty"`
}

// Encapsulating element for all Suffix Name data.
type SuffixNameDataType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc Type,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SupervisoryOrganizationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SupervisoryOrganizationObjectType struct {
	ID         []SupervisoryOrganizationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Support Order Dependent Data
type SupportOrderDependentDataType struct {
	ChildsNameLastFirstMI string    `xml:"urn:com.workday/bsvc Child_s_Name__Last__First__MI_"`
	ChildsBirthDate       time.Time `xml:"urn:com.workday/bsvc Child_s_Birth_Date"`
}

func (t *SupportOrderDependentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T SupportOrderDependentDataType
	var layout struct {
		*T
		ChildsBirthDate *xsdDate `xml:"urn:com.workday/bsvc Child_s_Birth_Date"`
	}
	layout.T = (*T)(t)
	layout.ChildsBirthDate = (*xsdDate)(&layout.T.ChildsBirthDate)
	return e.EncodeElement(layout, start)
}
func (t *SupportOrderDependentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SupportOrderDependentDataType
	var overlay struct {
		*T
		ChildsBirthDate *xsdDate `xml:"urn:com.workday/bsvc Child_s_Birth_Date"`
	}
	overlay.T = (*T)(t)
	overlay.ChildsBirthDate = (*xsdDate)(&overlay.T.ChildsBirthDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type SupportOrderObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SupportOrderObjectType struct {
	ID         []SupportOrderObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SupportTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SupportTypeObjectType struct {
	ID         []SupportTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TaxAddressTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TaxAddressTypeObjectType struct {
	ID         []TaxAddressTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Tax document delivery data
type TaxDocumentDeliveryDataType struct {
	WorkerReference         *WorkerObjectType                                 `xml:"urn:com.workday/bsvc Worker_Reference"`
	CompanyReference        *CompanyObjectType                                `xml:"urn:com.workday/bsvc Company_Reference"`
	FormGroupReference      *PayrollPayeeTaxReportTypeDeliveryGroupObjectType `xml:"urn:com.workday/bsvc Form_Group_Reference"`
	OverrideDefaultDelivery *bool                                             `xml:"urn:com.workday/bsvc Override_Default_Delivery,omitempty"`
}

// Request criteria to determine the tax document delivery elections you would like to view
type TaxDocumentDeliveryRequestCriteriaType struct {
	CompanyReference   *CompanyObjectType                                `xml:"urn:com.workday/bsvc Company_Reference"`
	FormGroupReference *PayrollPayeeTaxReportTypeDeliveryGroupObjectType `xml:"urn:com.workday/bsvc Form_Group_Reference"`
	WorkerReference    []WorkerObjectType                                `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
}

// Reference(s) to the tax document delivery election you would like to view
type TaxDocumentDeliveryRequestReferencesType struct {
	TaxDocumentDeliveryReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Tax_Document_Delivery_Reference"`
}

// Tax document deliveries response
type TaxDocumentDeliveryResponseDataType struct {
	TaxDocumentDelivery []TaxDocumentDeliveryType `xml:"urn:com.workday/bsvc Tax_Document_Delivery,omitempty"`
}

// Indicate if you want to Include the tax document delivery election references in the response
type TaxDocumentDeliveryResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Tax document delivery data
type TaxDocumentDeliveryType struct {
	TaxDocumentDeliveryReference *UniqueIdentifierObjectType   `xml:"urn:com.workday/bsvc Tax_Document_Delivery_Reference,omitempty"`
	TaxDocumentDeliveryData      []TaxDocumentDeliveryDataType `xml:"urn:com.workday/bsvc Tax_Document_Delivery_Data,omitempty"`
}

// Override the pay group default frequency for taxation with a Tax Frequency Override (USA Payroll Only).
type TaxFrequencyOverrideDataType struct {
	TaxFrequencyValue           float64                             `xml:"urn:com.workday/bsvc Tax_Frequency_Value,omitempty"`
	TaxFrequencyPeriodReference *TimeUnitforAnnualizationObjectType `xml:"urn:com.workday/bsvc Tax_Frequency_Period_Reference,omitempty"`
}

// Tax Levy Deduction Restriction Baseline Data
type TaxLevyDeductionRestrictionBaselineDataType struct {
	BaselineData []TaxLevyDeductionforBaselineDataType `xml:"urn:com.workday/bsvc Baseline_Data,omitempty"`
}

// Tax Levy Deduction Restriction Response Data
type TaxLevyDeductionRestrictionDataType struct {
	WithholdingOrderID               string                                       `xml:"urn:com.workday/bsvc Withholding_Order_ID,omitempty"`
	LockTaxElections                 *bool                                        `xml:"urn:com.workday/bsvc Lock_Tax_Elections,omitempty"`
	BaselineDeductionRestrictionData *TaxLevyDeductionRestrictionBaselineDataType `xml:"urn:com.workday/bsvc Baseline_Deduction_Restriction_Data,omitempty"`
	OverrideDeductionRestrictionData *TaxLevyDeductionRestrictionOverrideDataType `xml:"urn:com.workday/bsvc Override_Deduction_Restriction_Data,omitempty"`
}

// Tax Levy Deduction Restriction Override Data
type TaxLevyDeductionRestrictionOverrideDataType struct {
	OverrideData []TaxLevyDeductionforOverrideDataType `xml:"urn:com.workday/bsvc Override_Data,omitempty"`
}

// Tax Levy Deduction Restriction Request Criteria
type TaxLevyDeductionRestrictionRequestCriteriaType struct {
	WorkerReference  []WorkerObjectType  `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	CompanyReference []CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	EffectiveAsOf    *time.Time          `xml:"urn:com.workday/bsvc Effective_As_Of,omitempty"`
}

func (t *TaxLevyDeductionRestrictionRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T TaxLevyDeductionRestrictionRequestCriteriaType
	var layout struct {
		*T
		EffectiveAsOf *xsdDate `xml:"urn:com.workday/bsvc Effective_As_Of,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveAsOf = (*xsdDate)(layout.T.EffectiveAsOf)
	return e.EncodeElement(layout, start)
}
func (t *TaxLevyDeductionRestrictionRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T TaxLevyDeductionRestrictionRequestCriteriaType
	var overlay struct {
		*T
		EffectiveAsOf *xsdDate `xml:"urn:com.workday/bsvc Effective_As_Of,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveAsOf = (*xsdDate)(overlay.T.EffectiveAsOf)
	return d.DecodeElement(&overlay, &start)
}

// Tax Levy Deduction Restriction Request References
type TaxLevyDeductionRestrictionRequestReferencesType struct {
	TaxLevyReference []TaxLevyObjectType `xml:"urn:com.workday/bsvc Tax_Levy_Reference"`
}

// Tax Levy Deduction Restriction Response Group
type TaxLevyDeductionRestrictionResponseGroupType struct {
	ExcludeBaselineRestrictions *bool `xml:"urn:com.workday/bsvc Exclude_Baseline_Restrictions,omitempty"`
}

// Get Tax Levy Deduction Restrictions Response
type TaxLevyDeductionRestrictionsResponseDataType struct {
	TaxLevyDeductionRestriction []TaxLevyDeductionRestrictionsType `xml:"urn:com.workday/bsvc Tax_Levy_Deduction_Restriction,omitempty"`
}

// Tax Levy Deduction Restrictions
type TaxLevyDeductionRestrictionsType struct {
	TaxLevyReference                *TaxLevyObjectType                    `xml:"urn:com.workday/bsvc Tax_Levy_Reference,omitempty"`
	TaxLevyDeductionRestrictionData []TaxLevyDeductionRestrictionDataType `xml:"urn:com.workday/bsvc Tax_Levy_Deduction_Restriction_Data,omitempty"`
}

// Tax Levy Deduction for Baseline Data
type TaxLevyDeductionforBaselineDataType struct {
	EffectiveDate      time.Time            `xml:"urn:com.workday/bsvc Effective_Date"`
	DeductionReference *DeductionObjectType `xml:"urn:com.workday/bsvc Deduction_Reference"`
	Amount             float64              `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Percent            float64              `xml:"urn:com.workday/bsvc Percent,omitempty"`
}

func (t *TaxLevyDeductionforBaselineDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T TaxLevyDeductionforBaselineDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *TaxLevyDeductionforBaselineDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T TaxLevyDeductionforBaselineDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Tax Levy Deduction for Override Data
type TaxLevyDeductionforOverrideDataType struct {
	EffectiveDate      time.Time            `xml:"urn:com.workday/bsvc Effective_Date"`
	DeductionReference *DeductionObjectType `xml:"urn:com.workday/bsvc Deduction_Reference"`
	Amount             float64              `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Percent            float64              `xml:"urn:com.workday/bsvc Percent,omitempty"`
	NeverReducesDE     *bool                `xml:"urn:com.workday/bsvc Never_Reduces_DE,omitempty"`
	AlwaysReducesDE    *bool                `xml:"urn:com.workday/bsvc Always_Reduces_DE,omitempty"`
}

func (t *TaxLevyDeductionforOverrideDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T TaxLevyDeductionforOverrideDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *TaxLevyDeductionforOverrideDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T TaxLevyDeductionforOverrideDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type TaxLevyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TaxLevyObjectType struct {
	ID         []TaxLevyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TenantedPayrollWorktagObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type TenantedPayrollWorktagObjectType struct {
	ID         []TenantedPayrollWorktagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TimeTrackingEligibilityRuleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimeTrackingEligibilityRuleObjectType struct {
	ID         []TimeTrackingEligibilityRuleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TimeUnitforAnnualizationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimeUnitforAnnualizationObjectType struct {
	ID         []TimeUnitforAnnualizationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Container element for Valuation Date Override Data
type ValuationDateOverrideDataType struct {
	CountryReference               *CountryObjectType `xml:"urn:com.workday/bsvc Country_Reference"`
	PayrollSettlementValuationDate time.Time          `xml:"urn:com.workday/bsvc Payroll_Settlement_Valuation_Date"`
}

func (t *ValuationDateOverrideDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ValuationDateOverrideDataType
	var layout struct {
		*T
		PayrollSettlementValuationDate *xsdDate `xml:"urn:com.workday/bsvc Payroll_Settlement_Valuation_Date"`
	}
	layout.T = (*T)(t)
	layout.PayrollSettlementValuationDate = (*xsdDate)(&layout.T.PayrollSettlementValuationDate)
	return e.EncodeElement(layout, start)
}
func (t *ValuationDateOverrideDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ValuationDateOverrideDataType
	var overlay struct {
		*T
		PayrollSettlementValuationDate *xsdDate `xml:"urn:com.workday/bsvc Payroll_Settlement_Valuation_Date"`
	}
	overlay.T = (*T)(t)
	overlay.PayrollSettlementValuationDate = (*xsdDate)(&overlay.T.PayrollSettlementValuationDate)
	return d.DecodeElement(&overlay, &start)
}

// This is the wrapper element for Box 13 W-2GU Additional Data.
type W2GUAdditionalDataType struct {
	StatutoryEmployee *bool `xml:"urn:com.workday/bsvc Statutory_Employee,omitempty"`
	RetirementPlan    *bool `xml:"urn:com.workday/bsvc Retirement_Plan,omitempty"`
	ThirdPartySickPay *bool `xml:"urn:com.workday/bsvc Third_Party_Sick_Pay,omitempty"`
}

// This is wrapper element for W-2GU Box 12 Data.
type W2GUDeferredandOtherCompensationDataType struct {
	Code   string  `xml:"urn:com.workday/bsvc Code,omitempty"`
	Year   string  `xml:"urn:com.workday/bsvc Year,omitempty"`
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
}

// This is wrapper element for Annual Non Qualified plans Data.
type W2GUNonQualifiedPensionDataType struct {
	NonQualifiedPensionSection457DistYTD       float64 `xml:"urn:com.workday/bsvc Non-Qualified_Pension-Section_457_Dist_YTD,omitempty"`
	NonQualifiedPensionNonsection457DistYTD    float64 `xml:"urn:com.workday/bsvc Non-Qualified_Pension-Non-section_457_Dist_YTD,omitempty"`
	NonQualifiedPensionSection457ContribYTD    float64 `xml:"urn:com.workday/bsvc Non-Qualified_Pension-Section_457_Contrib_YTD,omitempty"`
	NonQualifiedPensionNonsection457ContribYTD float64 `xml:"urn:com.workday/bsvc Non-Qualified_Pension-Non-section_457_Contrib_YTD,omitempty"`
}

// This element is wrapper for W-2GU Box 14 Other Data.
type W2GUOtherDataType struct {
	Label                     string  `xml:"urn:com.workday/bsvc Label,omitempty"`
	Amount                    float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	OtherInformationFEINorEIN string  `xml:"urn:com.workday/bsvc Other_Information_-_FEIN_or_EIN,omitempty"`
}

// This element is wrapper for W-2GU Box 14 Other Data.
type W2GUOtherDataWrapperType struct {
	OtherData []W2GUOtherDataType `xml:"urn:com.workday/bsvc Other_Data,omitempty"`
}

// W-2PR Employer Boxes
type W2PREmployerBoxesType struct {
	QualifiedPhysician *bool  `xml:"urn:com.workday/bsvc Qualified_Physician,omitempty"`
	DomesticServices   *bool  `xml:"urn:com.workday/bsvc Domestic_Services,omitempty"`
	Others             *bool  `xml:"urn:com.workday/bsvc Others,omitempty"`
	OthersDescription  string `xml:"urn:com.workday/bsvc Others_Description,omitempty"`
}

// W-2PR Exempt Salary and Code Boxes
type W2PRExemptSalaryandCodeBoxesType struct {
	Box16Code    string  `xml:"urn:com.workday/bsvc Box_16_Code,omitempty"`
	Box16Amount  float64 `xml:"urn:com.workday/bsvc Box_16_Amount,omitempty"`
	Box16ACode   string  `xml:"urn:com.workday/bsvc Box_16A_Code,omitempty"`
	Box16AAmount float64 `xml:"urn:com.workday/bsvc Box_16A_Amount,omitempty"`
	Box16BCode   string  `xml:"urn:com.workday/bsvc Box_16B_Code,omitempty"`
	Box16BAmount float64 `xml:"urn:com.workday/bsvc Box_16B_Amount,omitempty"`
}

// This is the wrapper element for Box 13 W-2VI Additional Data.
type W2VIAdditionalDataType struct {
	StatutoryEmployee *bool `xml:"urn:com.workday/bsvc Statutory_Employee,omitempty"`
	RetirementPlan    *bool `xml:"urn:com.workday/bsvc Retirement_Plan,omitempty"`
	ThirdPartySickPay *bool `xml:"urn:com.workday/bsvc Third_Party_Sick_Pay,omitempty"`
}

// This is wrapper element for W-2VI Box 12 Data.
type W2VIDeferredandOtherCompensationDataType struct {
	Code   string  `xml:"urn:com.workday/bsvc Code,omitempty"`
	Year   string  `xml:"urn:com.workday/bsvc Year,omitempty"`
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
}

// This is wrapper element for Annual Non Qualified plans Data.
type W2VINonQualifiedPensionDataType struct {
	NonQualifiedPensionSection457DistYTD       float64 `xml:"urn:com.workday/bsvc Non-Qualified_Pension-Section_457_Dist_YTD,omitempty"`
	NonQualifiedPensionNonsection457DistYTD    float64 `xml:"urn:com.workday/bsvc Non-Qualified_Pension-Non-section_457_Dist_YTD,omitempty"`
	NonQualifiedPensionSection457ContribYTD    float64 `xml:"urn:com.workday/bsvc Non-Qualified_Pension-Section_457_Contrib_YTD,omitempty"`
	NonQualifiedPensionNonsection457ContribYTD float64 `xml:"urn:com.workday/bsvc Non-Qualified_Pension-Non-section_457_Contrib_YTD,omitempty"`
}

// This element is wrapper for W-2VI Box 14 Other Data.
type W2VIOtherDataType struct {
	Label                     string  `xml:"urn:com.workday/bsvc Label,omitempty"`
	Amount                    float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	OtherInformationFEINorEIN string  `xml:"urn:com.workday/bsvc Other_Information_-_FEIN_or_EIN,omitempty"`
}

// This element is wrapper for W-2VI Box 14 Other Data.
type W2VIOtherDataWrapperType struct {
	OtherData []W2VIOtherDataType `xml:"urn:com.workday/bsvc Other_Data,omitempty"`
}

// Element to hold all of the Data to Get the W-2/W-2C Print Election info.
type W2W2CPrintingElectionDataType struct {
	CompanyReference             *CompanyObjectType          `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	TaxDocumentDeliveryReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Tax_Document_Delivery_Reference,omitempty"`
	LastUpdated                  *time.Time                  `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	CurrentW2W2CPrintingElection string                      `xml:"urn:com.workday/bsvc Current_W-2_W-2C_Printing_Election,omitempty"`
}

func (t *W2W2CPrintingElectionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T W2W2CPrintingElectionDataType
	var layout struct {
		*T
		LastUpdated *xsdDate `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LastUpdated = (*xsdDate)(layout.T.LastUpdated)
	return e.EncodeElement(layout, start)
}
func (t *W2W2CPrintingElectionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T W2W2CPrintingElectionDataType
	var overlay struct {
		*T
		LastUpdated *xsdDate `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LastUpdated = (*xsdDate)(overlay.T.LastUpdated)
	return d.DecodeElement(&overlay, &start)
}

// Element to hold all criteria used to select which W-2/W-2C Printing Elections to return
type W2W2CPrintingElectionRequestCriteriaType struct {
	WorkerReference *WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
}

// Element to hold all of the Data to Put or Get for a W-2/W-2C Printing Election
type W2W2CPrintingElectionRequestDataType struct {
	WorkerReference                             *WorkerObjectType   `xml:"urn:com.workday/bsvc Worker_Reference"`
	CompanyReference                            []CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference"`
	ReceiveonlyelectroniccopiesofW2W2CReference *bool               `xml:"urn:com.workday/bsvc Receive_only_electronic_copies_of_W-2_W-2C_Reference,omitempty"`
}

// Element to hold Request References
type W2W2CPrintingElectionRequestReferencesType struct {
	CompanyReference *CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference"`
}

// Element to hold all data to Respond with about the W-2/W-2C Elections
type W2W2CPrintingElectionResponseDataType struct {
	W2W2CPrintingElection []W2W2CPrintingElectionType `xml:"urn:com.workday/bsvc W-2_W-2C_Printing_Election,omitempty"`
}

// Element to hold all information about the W-2/W-2C Printing Election Response Group
type W2W2CPrintingElectionResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element to hold all information about a W-2/W-2C Print Election
type W2W2CPrintingElectionType struct {
	XMLNAMEWorkerReference    *WorkerObjectType              `xml:"urn:com.workday/bsvc XMLNAME__Worker__Reference,omitempty"`
	W2W2CPrintingElectionData *W2W2CPrintingElectionDataType `xml:"urn:com.workday/bsvc W-2_W-2C_Printing_Election_Data,omitempty"`
}

// Washington Industrial Insurance Fund Data:  Contains position based Workers Compensation data
type WAIIFDeductionDataType struct {
	PayrollTaxAuthorityReference              []MetadataPayrollWorktagObjectType  `xml:"urn:com.workday/bsvc Payroll_Tax_Authority_Reference"`
	PositionReference                         []PositionObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	WorkersCompensationCodeReference          []WorkersCompensationCodeObjectType `xml:"urn:com.workday/bsvc Workers_Compensation_Code_Reference,omitempty"`
	WorkersCompensationRateReference          float64                             `xml:"urn:com.workday/bsvc Workers_Compensation_Rate_Reference,omitempty"`
	WashingtonIndustrialInsuranceFundQTDValue float64                             `xml:"urn:com.workday/bsvc Washington_Industrial_Insurance_Fund_QTD_Value,omitempty"`
	WashingtonIndustrialInsuranceFundYTDValue float64                             `xml:"urn:com.workday/bsvc Washington_Industrial_Insurance_Fund_YTD_Value,omitempty"`
	QTDTaxableHours                           float64                             `xml:"urn:com.workday/bsvc QTD_Taxable_Hours,omitempty"`
	YTDTaxableHours                           float64                             `xml:"urn:com.workday/bsvc YTD_Taxable_Hours,omitempty"`
}

// Wage Assignment Specific Data
type WageAssignmentSpecificDataType struct {
	HeadofHousehold *bool `xml:"urn:com.workday/bsvc Head_of_Household,omitempty"`
	Married         *bool `xml:"urn:com.workday/bsvc Married,omitempty"`
}

// Web Address Information
type WebAddressInformationDataType struct {
	WebAddress          string                                        `xml:"urn:com.workday/bsvc Web_Address,omitempty"`
	WebAddressComment   string                                        `xml:"urn:com.workday/bsvc Web_Address_Comment,omitempty"`
	UsageData           []CommunicationMethodUsageInformationDataType `xml:"urn:com.workday/bsvc Usage_Data,omitempty"`
	WebAddressReference *WebAddressReferenceObjectType                `xml:"urn:com.workday/bsvc Web_Address_Reference,omitempty"`
	ID                  string                                        `xml:"urn:com.workday/bsvc ID,omitempty"`
	Delete              bool                                          `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	DoNotReplaceAll     bool                                          `xml:"urn:com.workday/bsvc Do_Not_Replace_All,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WebAddressReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WebAddressReferenceObjectType struct {
	ID         []WebAddressReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type WithholdingOrderAmountTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WithholdingOrderAmountTypeObjectType struct {
	ID         []WithholdingOrderAmountTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WithholdingOrderCaseObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WithholdingOrderCaseObjectType struct {
	ID         []WithholdingOrderCaseObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Withholding Order Data for Get Payroll Results
type WithholdingOrderDataforGetPayrollResultsType struct {
	CaseNumber                                  string                                `xml:"urn:com.workday/bsvc Case_Number,omitempty"`
	WithholdingOrderType                        string                                `xml:"urn:com.workday/bsvc Withholding_Order_Type,omitempty"`
	OrderDate                                   *time.Time                            `xml:"urn:com.workday/bsvc Order_Date,omitempty"`
	ReceivedDate                                *time.Time                            `xml:"urn:com.workday/bsvc Received_Date,omitempty"`
	BeginDate                                   *time.Time                            `xml:"urn:com.workday/bsvc Begin_Date,omitempty"`
	EndDate                                     *time.Time                            `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	OrderStatus                                 string                                `xml:"urn:com.workday/bsvc Order_Status,omitempty"`
	OrderPriority                               float64                               `xml:"urn:com.workday/bsvc Order_Priority,omitempty"`
	TotalDebtAmount                             float64                               `xml:"urn:com.workday/bsvc Total_Debt_Amount,omitempty"`
	PayPeriodAmountTypeReference                *WithholdingOrderAmountTypeObjectType `xml:"urn:com.workday/bsvc Pay_Period_Amount_Type_Reference,omitempty"`
	PayPeriodAmount                             float64                               `xml:"urn:com.workday/bsvc Pay_Period_Amount,omitempty"`
	PayPeriodAmountasPercent                    float64                               `xml:"urn:com.workday/bsvc Pay_Period_Amount_as_Percent,omitempty"`
	PayPeriodFrequencyReference                 *FrequencyObjectType                  `xml:"urn:com.workday/bsvc Pay_Period_Frequency_Reference,omitempty"`
	PayrollTaxAuthorityReference                *PayrollTaxAuthorityObjectType        `xml:"urn:com.workday/bsvc Payroll_Tax_Authority_Reference,omitempty"`
	DeductionRecipientReference                 *DeductionRecipientObjectType         `xml:"urn:com.workday/bsvc Deduction_Recipient_Reference,omitempty"`
	DeductionRecipientInstantMessengerReference []UniqueIdentifierObjectType          `xml:"urn:com.workday/bsvc Deduction_Recipient_Instant_Messenger_Reference,omitempty"`
	DeductionRecipientBankAccountReference      *SettlementBankAccountObjectType      `xml:"urn:com.workday/bsvc Deduction_Recipient_Bank_Account_Reference,omitempty"`
	OriginatingEntity                           string                                `xml:"urn:com.workday/bsvc Originating_Entity,omitempty"`
	Memo                                        string                                `xml:"urn:com.workday/bsvc Memo,omitempty"`
	DeductionRecipientAddressData               []AddressInformationDataType          `xml:"urn:com.workday/bsvc Deduction_Recipient_Address_Data,omitempty"`
}

func (t *WithholdingOrderDataforGetPayrollResultsType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WithholdingOrderDataforGetPayrollResultsType
	var layout struct {
		*T
		OrderDate    *xsdDate     `xml:"urn:com.workday/bsvc Order_Date,omitempty"`
		ReceivedDate *xsdDateTime `xml:"urn:com.workday/bsvc Received_Date,omitempty"`
		BeginDate    *xsdDate     `xml:"urn:com.workday/bsvc Begin_Date,omitempty"`
		EndDate      *xsdDate     `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.OrderDate = (*xsdDate)(layout.T.OrderDate)
	layout.ReceivedDate = (*xsdDateTime)(layout.T.ReceivedDate)
	layout.BeginDate = (*xsdDate)(layout.T.BeginDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *WithholdingOrderDataforGetPayrollResultsType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WithholdingOrderDataforGetPayrollResultsType
	var overlay struct {
		*T
		OrderDate    *xsdDate     `xml:"urn:com.workday/bsvc Order_Date,omitempty"`
		ReceivedDate *xsdDateTime `xml:"urn:com.workday/bsvc Received_Date,omitempty"`
		BeginDate    *xsdDate     `xml:"urn:com.workday/bsvc Begin_Date,omitempty"`
		EndDate      *xsdDate     `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.OrderDate = (*xsdDate)(overlay.T.OrderDate)
	overlay.ReceivedDate = (*xsdDateTime)(overlay.T.ReceivedDate)
	overlay.BeginDate = (*xsdDate)(overlay.T.BeginDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Withholding Order Fee Data
type WithholdingOrderFeeDataType struct {
	FeeAmount                   float64                            `xml:"urn:com.workday/bsvc Fee_Amount"`
	FeeTypeReference            *WithholdingOrderFeeTypeObjectType `xml:"urn:com.workday/bsvc Fee_Type_Reference"`
	DeductionRecipientReference *DeductionRecipientObjectType      `xml:"urn:com.workday/bsvc Deduction_Recipient_Reference,omitempty"`
	OverrideFeeSchedule         *bool                              `xml:"urn:com.workday/bsvc Override_Fee_Schedule,omitempty"`
	BeginDate                   *time.Time                         `xml:"urn:com.workday/bsvc Begin_Date,omitempty"`
	EndDate                     *time.Time                         `xml:"urn:com.workday/bsvc End_Date,omitempty"`
}

func (t *WithholdingOrderFeeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WithholdingOrderFeeDataType
	var layout struct {
		*T
		BeginDate *xsdDate `xml:"urn:com.workday/bsvc Begin_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.BeginDate = (*xsdDate)(layout.T.BeginDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *WithholdingOrderFeeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WithholdingOrderFeeDataType
	var overlay struct {
		*T
		BeginDate *xsdDate `xml:"urn:com.workday/bsvc Begin_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.BeginDate = (*xsdDate)(overlay.T.BeginDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type WithholdingOrderFeeTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WithholdingOrderFeeTypeObjectType struct {
	ID         []WithholdingOrderFeeTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Withholding Order Local Minimum Wage Rate Request
type WithholdingOrderLocalMinimumWageRateRequestType struct {
	CompanyReference []CompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	WorkerReference  []WorkerObjectType  `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WithholdingOrderObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WithholdingOrderObjectType struct {
	ID         []WithholdingOrderObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WithholdingOrderOverrideCompletionCriteriaObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WithholdingOrderOverrideCompletionCriteriaObjectType struct {
	ID         []WithholdingOrderOverrideCompletionCriteriaObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WithholdingOrderTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WithholdingOrderTypeObjectType struct {
	ID         []WithholdingOrderTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Withholding Order Withholding Frequency Data
type WithholdingOrderWithholdingFrequencyDataType struct {
	WithholdingOrderWithholdingFrequencyReference []WithholdingOrderWithholdingFrequencyObjectType `xml:"urn:com.workday/bsvc Withholding_Order_Withholding_Frequency_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WithholdingOrderWithholdingFrequencyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WithholdingOrderWithholdingFrequencyObjectType struct {
	ID         []WithholdingOrderWithholdingFrequencyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type WorkdayCommonHeaderType struct {
	IncludeReferenceDescriptorsInResponse *bool `xml:"urn:com.workday/bsvc Include_Reference_Descriptors_In_Response,omitempty"`
}

// Wrapper element for Non-3PSP Annual Tax Filing Data
type WorkerAnnualTaxDataType struct {
	ThirdPartySickPay                *bool                                  `xml:"urn:com.workday/bsvc Third_Party_Sick_Pay,omitempty"`
	SocialSecurityTips               float64                                `xml:"urn:com.workday/bsvc Social_Security_Tips,omitempty"`
	AllocatedTips                    float64                                `xml:"urn:com.workday/bsvc Allocated_Tips,omitempty"`
	DependentCareBenefits            float64                                `xml:"urn:com.workday/bsvc Dependent_Care_Benefits,omitempty"`
	PrintWorkersTaxDocuments         *bool                                  `xml:"urn:com.workday/bsvc Print_Workers_Tax_Documents,omitempty"`
	NonQualifiedPensionData          *NonQualifiedPensionDataType           `xml:"urn:com.workday/bsvc Non-Qualified_Pension_Data,omitempty"`
	DeferredandOtherCompensationData []DeferredandOtherCompensationDataType `xml:"urn:com.workday/bsvc Deferred_and_Other_Compensation_Data,omitempty"`
	AdditionalData                   *AdditionalDataType                    `xml:"urn:com.workday/bsvc Additional_Data,omitempty"`
	OtherDataWrapper                 []OtherDataWrapperType                 `xml:"urn:com.workday/bsvc Other_Data_Wrapper,omitempty"`
}

// This is a wrapper element for Worker Annual Tax Filing Data.
type WorkerAnnualTaxDataWrapperType struct {
	WorkerAnnualTaxData []WorkerAnnualTaxDataType `xml:"urn:com.workday/bsvc Worker_Annual_Tax_Data,omitempty"`
}

// Details of the allocation, (e.g., a set of allocation dimensions and percentages)
type WorkerCostingAllocationDetailDataType struct {
	Order                                       string                                            `xml:"urn:com.workday/bsvc Order,omitempty"`
	DefaultfromOrganizationAssignment           *bool                                             `xml:"urn:com.workday/bsvc Default_from_Organization_Assignment,omitempty"`
	CostingOverrideWorktagReference             []TenantedPayrollWorktagObjectType                `xml:"urn:com.workday/bsvc Costing_Override_Worktag_Reference,omitempty"`
	DistributionPercent                         float64                                           `xml:"urn:com.workday/bsvc Distribution_Percent,omitempty"`
	SalaryOvertheCapCostingAllocationDetailData []SalaryOvertheCapCostingAllocationDetailDataType `xml:"urn:com.workday/bsvc Salary_Over_the_Cap_Costing_Allocation_Detail_Data,omitempty"`
}

// Costing Allocation Data effective over a specified date range.
type WorkerCostingAllocationIntervalDataType struct {
	CostingOverrideID                 string                                  `xml:"urn:com.workday/bsvc Costing_Override_ID,omitempty"`
	StartDate                         *time.Time                              `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                           *time.Time                              `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	WorkerCostingAllocationDetailData []WorkerCostingAllocationDetailDataType `xml:"urn:com.workday/bsvc Worker_Costing_Allocation_Detail_Data,omitempty"`
}

func (t *WorkerCostingAllocationIntervalDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerCostingAllocationIntervalDataType
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
func (t *WorkerCostingAllocationIntervalDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerCostingAllocationIntervalDataType
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

// Primary (Worker or Position Restrictions) and Secondary (Position and Earning) keys identifying the level of the costing allocation data.
type WorkerCostingAllocationsDataType struct {
	WorkerReference                     *WorkerObjectType                         `xml:"urn:com.workday/bsvc Worker_Reference"`
	PositionRestrictionsReference       *PositionRestrictionsObjectType           `xml:"urn:com.workday/bsvc Position_Restrictions_Reference"`
	PositionReference                   *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	EarningReference                    *EarningAllObjectType                     `xml:"urn:com.workday/bsvc Earning_Reference,omitempty"`
	WorkerCostingAllocationIntervalData []WorkerCostingAllocationIntervalDataType `xml:"urn:com.workday/bsvc Worker_Costing_Allocation_Interval_Data,omitempty"`
}

// Specific request criteria for the Get Worker Costing Allocations service.
type WorkerCostingAllocationsRequestCriteriaType struct {
	CostingOverrideCriteria                          *CostingOverrideCriteriaType `xml:"urn:com.workday/bsvc Costing_Override_Criteria,omitempty"`
	IncludeAllWorkerCostingAllocations               *bool                        `xml:"urn:com.workday/bsvc Include_All_Worker_Costing_Allocations,omitempty"`
	IncludeAllPositionRestrictionsCostingAllocations *bool                        `xml:"urn:com.workday/bsvc Include_All_Position_Restrictions_Costing_Allocations,omitempty"`
	IncludeOnlyCurrentCostingOverrideInterval        *bool                        `xml:"urn:com.workday/bsvc Include_Only_Current_Costing_Override_Interval,omitempty"`
}

// Wrapper for the data returned by the Get Worker Costing Allocations service.
type WorkerCostingAllocationsResponseDataType struct {
	WorkerCostingAllocationsData []WorkerCostingAllocationsDataType `xml:"urn:com.workday/bsvc Worker_Costing_Allocations_Data,omitempty"`
}

// Configure the content returned for the Get Worker Costing Allocations service.
type WorkerCostingAllocationsResponseGroupType struct {
	ExcludeAllocationDetailData *bool `xml:"urn:com.workday/bsvc Exclude_Allocation_Detail_Data,omitempty"`
}

// Worker Guam Annual Tax Data
type WorkerGuamAnnualTaxDataType struct {
	WorkerW2GUData                   []WorkerW2GUDataType                       `xml:"urn:com.workday/bsvc Worker_W-2GU_Data,omitempty"`
	Controlnumber                    string                                     `xml:"urn:com.workday/bsvc Control_number,omitempty"`
	NonQualifiedPensionData          *W2GUNonQualifiedPensionDataType           `xml:"urn:com.workday/bsvc Non-Qualified_Pension_Data,omitempty"`
	DeferredandOtherCompensationData []W2GUDeferredandOtherCompensationDataType `xml:"urn:com.workday/bsvc Deferred_and_Other_Compensation_Data,omitempty"`
	AdditionalData                   *W2GUAdditionalDataType                    `xml:"urn:com.workday/bsvc Additional_Data,omitempty"`
	OtherDataWrapper                 []W2GUOtherDataWrapperType                 `xml:"urn:com.workday/bsvc Other_Data_Wrapper,omitempty"`
}

// Wrapper element for Worker Guam Annual Tax Filing Data.
type WorkerGuamAnnualTaxDataWrapperType struct {
	WorkerGuamAnnualTaxData []WorkerGuamAnnualTaxDataType `xml:"urn:com.workday/bsvc Worker_Guam_Annual_Tax_Data,omitempty"`
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

// This element contains Worker Payroll Reporting Code Data.
type WorkerPayrollReportingCodeDataType struct {
	PayrollReportingTypeReference *PayrollReportingTypeObjectType `xml:"urn:com.workday/bsvc Payroll_Reporting_Type_Reference,omitempty"`
	PayrollReportingCode          string                          `xml:"urn:com.workday/bsvc Payroll_Reporting_Code,omitempty"`
	FormattedPayrollReportingCode string                          `xml:"urn:com.workday/bsvc Formatted_Payroll_Reporting_Code,omitempty"`
}

// Worker Puerto Rico Annual Tax Data
type WorkerPuertoRicoAnnualTaxDataType struct {
	WorkerW2PRData                       []WorkerW2PRDataType               `xml:"urn:com.workday/bsvc Worker_W-2PR_Data,omitempty"`
	EmployerSponsoredHealthCare          float64                            `xml:"urn:com.workday/bsvc Employer_Sponsored_Health_Care,omitempty"`
	CharitableContributions              float64                            `xml:"urn:com.workday/bsvc Charitable_Contributions,omitempty"`
	EmployerBoxes                        []W2PREmployerBoxesType            `xml:"urn:com.workday/bsvc Employer_Boxes,omitempty"`
	Commissions                          float64                            `xml:"urn:com.workday/bsvc Commissions,omitempty"`
	Allowances                           float64                            `xml:"urn:com.workday/bsvc Allowances,omitempty"`
	ReimbursedExpenses                   float64                            `xml:"urn:com.workday/bsvc Reimbursed_Expenses,omitempty"`
	GovernmentalRetirementFund           float64                            `xml:"urn:com.workday/bsvc Governmental_Retirement_Fund,omitempty"`
	ContributionstoCODAPlans             float64                            `xml:"urn:com.workday/bsvc Contributions_to_CODA_Plans,omitempty"`
	ExemptSalariesandCodesBoxInformation []W2PRExemptSalaryandCodeBoxesType `xml:"urn:com.workday/bsvc Exempt_Salaries_and_Codes_Box_Information,omitempty"`
	ExemptSalariesandCodes               []Box16DataWrapperType             `xml:"urn:com.workday/bsvc Exempt_Salaries_and_Codes,omitempty"`
	ContributionstoSADYMProgram          float64                            `xml:"urn:com.workday/bsvc Contributions_to_SADYM_Program,omitempty"`
	UncollectedSocialSecurityTaxonTips   float64                            `xml:"urn:com.workday/bsvc Uncollected_Social_Security_Tax_on_Tips,omitempty"`
	UncollectedMedicareTaxonTips         float64                            `xml:"urn:com.workday/bsvc Uncollected_Medicare_Tax_on_Tips,omitempty"`
	ControlNumber                        string                             `xml:"urn:com.workday/bsvc Control_Number,omitempty"`
	PrintWorkersTaxDocuments             *bool                              `xml:"urn:com.workday/bsvc Print_Workers_Tax_Documents,omitempty"`
}

// Wrapper element for Worker Puerto Rico Annual Tax Filing Data.
type WorkerPuertoRicoAnnualTaxDataWrapperType struct {
	WorkerPuertoRicoAnnualTaxData []WorkerPuertoRicoAnnualTaxDataType `xml:"urn:com.workday/bsvc Worker_Puerto_Rico_Annual_Tax_Data,omitempty"`
}

// Reference element representing a unique instance of Worker (e.g. Employee or Contingent Worker).
type WorkerReferenceType struct {
	EmployeeReference         *EmployeeReferenceType             `xml:"urn:com.workday/bsvc Employee_Reference"`
	ContingentWorkerReference *ContingentWorkerReferenceDataType `xml:"urn:com.workday/bsvc Contingent_Worker_Reference"`
}

// This Element contains Worker Tax Elections data.
type WorkerTaxElectionsType struct {
	NumberofAllowances float64 `xml:"urn:com.workday/bsvc Number_of_Allowances,omitempty"`
	MaritalStatus      string  `xml:"urn:com.workday/bsvc Marital_Status,omitempty"`
	Exempted           *bool   `xml:"urn:com.workday/bsvc Exempted,omitempty"`
}

// Get Worker Tax Treaty Request Criteria.  The values in this element are used to filter the list of Worker Tax Treaty instances returned by the Get Worker Tax Treaty Request.
type WorkerTaxTreatiesRequestCriteriaType struct {
	WorkerReference            []WorkerObjectType                  `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	TaxYearReference           *CalendarYearObjectType             `xml:"urn:com.workday/bsvc Tax_Year_Reference,omitempty"`
	IncomeCodeReference        *PayrollIncomeCodeObjectType        `xml:"urn:com.workday/bsvc Income_Code_Reference,omitempty"`
	IncomeCodeSubtypeReference *PayrollIncomeCodeSubtypeObjectType `xml:"urn:com.workday/bsvc Income_Code_Subtype_Reference,omitempty"`
}

// Provide one or more Worker tax Treaty References to retrieve a specific set of Worker Tax Treaties
type WorkerTaxTreatiesRequestReferencesType struct {
	WorkerTaxTreatyReference []PayrollPayeeTaxTreatyUSAObjectType `xml:"urn:com.workday/bsvc Worker_Tax_Treaty_Reference"`
}

// Get Worker Tax Treaties Response Data
type WorkerTaxTreatiesResponseDataType struct {
	WorkerTaxTreaty []WorkerTaxTreatyType `xml:"urn:com.workday/bsvc Worker_Tax_Treaty,omitempty"`
}

// Worker Tax Treaty Data
type WorkerTaxTreatyDataType struct {
	WorkerReference                  *WorkerObjectType                   `xml:"urn:com.workday/bsvc Worker_Reference"`
	TaxYearReference                 *CalendarYearObjectType             `xml:"urn:com.workday/bsvc Tax_Year_Reference"`
	TreatyBenefitStartDate           time.Time                           `xml:"urn:com.workday/bsvc Treaty_Benefit_Start_Date"`
	TreatyBenefitEndDate             time.Time                           `xml:"urn:com.workday/bsvc Treaty_Benefit_End_Date"`
	MaximumBenefitAmount             float64                             `xml:"urn:com.workday/bsvc Maximum_Benefit_Amount,omitempty"`
	NoTreatyMaximum                  *bool                               `xml:"urn:com.workday/bsvc No_Treaty_Maximum,omitempty"`
	TreatyEligibilityBegin           *time.Time                          `xml:"urn:com.workday/bsvc Treaty_Eligibility_Begin,omitempty"`
	TreatyEligibilityEnd             *time.Time                          `xml:"urn:com.workday/bsvc Treaty_Eligibility_End,omitempty"`
	IncomeCodeReference              *PayrollIncomeCodeObjectType        `xml:"urn:com.workday/bsvc Income_Code_Reference"`
	IncomeCodeSubtypeReference       *PayrollIncomeCodeSubtypeObjectType `xml:"urn:com.workday/bsvc Income_Code_Subtype_Reference,omitempty"`
	EligibleforWithholdingAllowance  *bool                               `xml:"urn:com.workday/bsvc Eligible_for_Withholding_Allowance,omitempty"`
	WithholdingRate                  float64                             `xml:"urn:com.workday/bsvc Withholding_Rate,omitempty"`
	TaxResidencyCountryCodeReference *PayrollIRSCountryObjectType        `xml:"urn:com.workday/bsvc Tax_Residency_Country_Code_Reference"`
}

func (t *WorkerTaxTreatyDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerTaxTreatyDataType
	var layout struct {
		*T
		TreatyBenefitStartDate *xsdDate `xml:"urn:com.workday/bsvc Treaty_Benefit_Start_Date"`
		TreatyBenefitEndDate   *xsdDate `xml:"urn:com.workday/bsvc Treaty_Benefit_End_Date"`
		TreatyEligibilityBegin *xsdDate `xml:"urn:com.workday/bsvc Treaty_Eligibility_Begin,omitempty"`
		TreatyEligibilityEnd   *xsdDate `xml:"urn:com.workday/bsvc Treaty_Eligibility_End,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TreatyBenefitStartDate = (*xsdDate)(&layout.T.TreatyBenefitStartDate)
	layout.TreatyBenefitEndDate = (*xsdDate)(&layout.T.TreatyBenefitEndDate)
	layout.TreatyEligibilityBegin = (*xsdDate)(layout.T.TreatyEligibilityBegin)
	layout.TreatyEligibilityEnd = (*xsdDate)(layout.T.TreatyEligibilityEnd)
	return e.EncodeElement(layout, start)
}
func (t *WorkerTaxTreatyDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerTaxTreatyDataType
	var overlay struct {
		*T
		TreatyBenefitStartDate *xsdDate `xml:"urn:com.workday/bsvc Treaty_Benefit_Start_Date"`
		TreatyBenefitEndDate   *xsdDate `xml:"urn:com.workday/bsvc Treaty_Benefit_End_Date"`
		TreatyEligibilityBegin *xsdDate `xml:"urn:com.workday/bsvc Treaty_Eligibility_Begin,omitempty"`
		TreatyEligibilityEnd   *xsdDate `xml:"urn:com.workday/bsvc Treaty_Eligibility_End,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TreatyBenefitStartDate = (*xsdDate)(&overlay.T.TreatyBenefitStartDate)
	overlay.TreatyBenefitEndDate = (*xsdDate)(&overlay.T.TreatyBenefitEndDate)
	overlay.TreatyEligibilityBegin = (*xsdDate)(overlay.T.TreatyEligibilityBegin)
	overlay.TreatyEligibilityEnd = (*xsdDate)(overlay.T.TreatyEligibilityEnd)
	return d.DecodeElement(&overlay, &start)
}

// Worker Tax Treaty
type WorkerTaxTreatyType struct {
	WorkerTaxTreatyReference *PayrollPayeeTaxTreatyUSAObjectType `xml:"urn:com.workday/bsvc Worker_Tax_Treaty_Reference,omitempty"`
	WorkerTaxTreatyData      []WorkerTaxTreatyDataType           `xml:"urn:com.workday/bsvc Worker_Tax_Treaty_Data,omitempty"`
}

// Worker Virgin Islands Annual Tax Data
type WorkerVirginIslandsAnnualTaxDataType struct {
	WorkerW2VIData                   []WorkerW2VIDataType                       `xml:"urn:com.workday/bsvc Worker_W-2VI_Data,omitempty"`
	Controlnumber                    string                                     `xml:"urn:com.workday/bsvc Control_number,omitempty"`
	NonQualifiedPensionData          *W2VINonQualifiedPensionDataType           `xml:"urn:com.workday/bsvc Non-Qualified_Pension_Data,omitempty"`
	DeferredandOtherCompensationData []W2VIDeferredandOtherCompensationDataType `xml:"urn:com.workday/bsvc Deferred_and_Other_Compensation_Data,omitempty"`
	AdditionalData                   *W2VIAdditionalDataType                    `xml:"urn:com.workday/bsvc Additional_Data,omitempty"`
	OtherDataWrapper                 []W2VIOtherDataWrapperType                 `xml:"urn:com.workday/bsvc Other_Data_Wrapper,omitempty"`
}

// Wrapper element for Worker Virgin Islands Annual Tax Filing Data.
type WorkerVirginIslandsAnnualTaxDataWrapperType struct {
	WorkerVirginIslandsAnnualTaxData []WorkerVirginIslandsAnnualTaxDataType `xml:"urn:com.workday/bsvc Worker_Virgin_Islands_Annual_Tax_Data,omitempty"`
}

// Worker W-2GU Data
type WorkerW2GUDataType struct {
	W2GUInstanceID  string     `xml:"urn:com.workday/bsvc W-2GU_Instance_ID,omitempty"`
	CompletedMoment *time.Time `xml:"urn:com.workday/bsvc Completed_Moment,omitempty"`
}

func (t *WorkerW2GUDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerW2GUDataType
	var layout struct {
		*T
		CompletedMoment *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Moment,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CompletedMoment = (*xsdDateTime)(layout.T.CompletedMoment)
	return e.EncodeElement(layout, start)
}
func (t *WorkerW2GUDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerW2GUDataType
	var overlay struct {
		*T
		CompletedMoment *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Moment,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CompletedMoment = (*xsdDateTime)(overlay.T.CompletedMoment)
	return d.DecodeElement(&overlay, &start)
}

// Worker Puerto Rico Annual Tax Data
type WorkerW2PRDataType struct {
	W2PRInstanceID  string     `xml:"urn:com.workday/bsvc W-2PR_Instance_ID,omitempty"`
	CompletedMoment *time.Time `xml:"urn:com.workday/bsvc Completed_Moment,omitempty"`
}

func (t *WorkerW2PRDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerW2PRDataType
	var layout struct {
		*T
		CompletedMoment *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Moment,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CompletedMoment = (*xsdDateTime)(layout.T.CompletedMoment)
	return e.EncodeElement(layout, start)
}
func (t *WorkerW2PRDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerW2PRDataType
	var overlay struct {
		*T
		CompletedMoment *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Moment,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CompletedMoment = (*xsdDateTime)(overlay.T.CompletedMoment)
	return d.DecodeElement(&overlay, &start)
}

// Worker W-2VI Data
type WorkerW2VIDataType struct {
	W2VIInstanceID  string     `xml:"urn:com.workday/bsvc W-2VI_Instance_ID,omitempty"`
	CompletedMoment *time.Time `xml:"urn:com.workday/bsvc Completed_Moment,omitempty"`
}

func (t *WorkerW2VIDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerW2VIDataType
	var layout struct {
		*T
		CompletedMoment *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Moment,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CompletedMoment = (*xsdDateTime)(layout.T.CompletedMoment)
	return e.EncodeElement(layout, start)
}
func (t *WorkerW2VIDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerW2VIDataType
	var overlay struct {
		*T
		CompletedMoment *xsdDateTime `xml:"urn:com.workday/bsvc Completed_Moment,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CompletedMoment = (*xsdDateTime)(overlay.T.CompletedMoment)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type WorkersCompensationCodeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkersCompensationCodeObjectType struct {
	ID         []WorkersCompensationCodeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Data for Location (Business Site).  Includes identification, address, number of employees, and wages.
type WorksiteDataType struct {
	WorksiteReference          *LocationObjectType `xml:"urn:com.workday/bsvc Worksite_Reference,omitempty"`
	TradeName                  string              `xml:"urn:com.workday/bsvc Trade_Name,omitempty"`
	WorksiteStreetAddress      string              `xml:"urn:com.workday/bsvc Worksite_Street_Address,omitempty"`
	WorksiteCity               string              `xml:"urn:com.workday/bsvc Worksite_City,omitempty"`
	WorksiteState              string              `xml:"urn:com.workday/bsvc Worksite_State,omitempty"`
	WorksiteZipCode            string              `xml:"urn:com.workday/bsvc Worksite_Zip_Code,omitempty"`
	WorksiteExpandedZipCode    string              `xml:"urn:com.workday/bsvc Worksite_Expanded_Zip_Code,omitempty"`
	WorksiteIdentificationCode string              `xml:"urn:com.workday/bsvc Worksite_Identification_Code,omitempty"`
	WorksiteDescription        string              `xml:"urn:com.workday/bsvc Worksite_Description,omitempty"`
	EmployeesActiveMonth1Day12 float64             `xml:"urn:com.workday/bsvc Employees_Active_Month_1_Day_12,omitempty"`
	EmployeesActiveMonth2Day12 float64             `xml:"urn:com.workday/bsvc Employees_Active_Month_2_Day_12,omitempty"`
	EmployeesActiveMonth3Day12 float64             `xml:"urn:com.workday/bsvc Employees_Active_Month_3_Day_12,omitempty"`
	Wages                      float64             `xml:"urn:com.workday/bsvc Wages,omitempty"`
}

// Monthly Tax Filing Deduction Data
type YTDTaxFilingDataType struct {
	TaxWithheld  float64 `xml:"urn:com.workday/bsvc Tax_Withheld,omitempty"`
	TaxableWages float64 `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages float64 `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	GrossWages   float64 `xml:"urn:com.workday/bsvc Gross_Wages,omitempty"`
}

// YTD Tax Filing Data for Periodic Company
type YTDTaxFilingDataforPeriodicCompanyType struct {
	TaxWithheld  float64 `xml:"urn:com.workday/bsvc Tax_Withheld,omitempty"`
	TaxableWages float64 `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages float64 `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	GrossWages   float64 `xml:"urn:com.workday/bsvc Gross_Wages,omitempty"`
}

// YTD Tax Filing Data for Periodic Worker
type YTDTaxFilingDataforPeriodicWorkerType struct {
	TaxWithheld  float64 `xml:"urn:com.workday/bsvc Tax_Withheld,omitempty"`
	TaxableWages float64 `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages float64 `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	GrossAmount  float64 `xml:"urn:com.workday/bsvc Gross_Amount,omitempty"`
	TipWages     float64 `xml:"urn:com.workday/bsvc Tip_Wages,omitempty"`
}

// YTD Tax Filing Data for Quarterly Worker
type YTDTaxFilingDataforQuarterlyWorkerType struct {
	TaxWithheld       float64 `xml:"urn:com.workday/bsvc Tax_Withheld,omitempty"`
	TaxableWages      float64 `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages      float64 `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	GrossWages        float64 `xml:"urn:com.workday/bsvc Gross_Wages,omitempty"`
	TipWages          float64 `xml:"urn:com.workday/bsvc Tip_Wages,omitempty"`
	UncollectedTipTax float64 `xml:"urn:com.workday/bsvc Uncollected_Tip_Tax,omitempty"`
	UncollectedGTL    float64 `xml:"urn:com.workday/bsvc Uncollected_GTL,omitempty"`
	W2Eligible        *bool   `xml:"urn:com.workday/bsvc W-2_Eligible,omitempty"`
}

// Filter to be used for YTD Results when Payroll Tax Authority and Deductions are specified.
type YTDTaxFilterType struct {
	PayrollTaxAuthorityReference []PayrollTaxAuthorityObjectType   `xml:"urn:com.workday/bsvc Payroll_Tax_Authority_Reference,omitempty"`
	DeductionReference           []DeductionWorkdayOwnedObjectType `xml:"urn:com.workday/bsvc Deduction_Reference,omitempty"`
}

// YTD Tax Remittance Data for Periodic Company
type YTDTaxRemittanceDataforPeriodicCompanyType struct {
	TaxWithheld       float64 `xml:"urn:com.workday/bsvc Tax_Withheld,omitempty"`
	TaxableWages      float64 `xml:"urn:com.workday/bsvc Taxable_Wages,omitempty"`
	SubjectWages      float64 `xml:"urn:com.workday/bsvc Subject_Wages,omitempty"`
	ExemptWages       float64 `xml:"urn:com.workday/bsvc Exempt_Wages,omitempty"`
	WCBInsurableWages float64 `xml:"urn:com.workday/bsvc WCB_Insurable_Wages,omitempty"`
	WCBGrossWages     float64 `xml:"urn:com.workday/bsvc WCB_Gross_Wages,omitempty"`
	WCBOtherWages     float64 `xml:"urn:com.workday/bsvc WCB_Other_Wages,omitempty"`
	WCBExcessWages    float64 `xml:"urn:com.workday/bsvc WCB_Excess_Wages,omitempty"`
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
