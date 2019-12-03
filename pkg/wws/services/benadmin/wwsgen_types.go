package benadmin

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// ACA 1095-C Recipient Form
type ACA1095CFormDataType struct {
	EmployeeInformationData             *EmployeeInformationDataType              `xml:"urn:com.workday/bsvc Employee_Information_Data,omitempty"`
	PlanStartMonth                      string                                    `xml:"urn:com.workday/bsvc Plan_Start_Month,omitempty"`
	Line14MonthData                     *Line14MonthDataType                      `xml:"urn:com.workday/bsvc Line_14_Month_Data,omitempty"`
	Line15MonthData                     *Line15MonthDataType                      `xml:"urn:com.workday/bsvc Line_15_Month_Data,omitempty"`
	Line16MonthData                     *Line16MonthDataType                      `xml:"urn:com.workday/bsvc Line_16_Month_Data,omitempty"`
	Line1734CoveredIndividualsMonthData []Line1734CoveredIndividualsMonthDataType `xml:"urn:com.workday/bsvc Line_17-34_Covered_Individuals_Month_Data,omitempty"`
}

// Request Criteria
type ACA1095CFormRequestCriteriaType struct {
	CalendarYearReference           *CalendarYearObjectType            `xml:"urn:com.workday/bsvc Calendar_Year_Reference"`
	CompanyReference                *CompanyObjectType                 `xml:"urn:com.workday/bsvc Company_Reference"`
	EmployeeReference               []WorkerObjectType                 `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	External1095CRecipientReference []External1095CRecipientObjectType `xml:"urn:com.workday/bsvc External_1095-C_Recipient_Reference,omitempty"`
}

// GET ACA 1095-C Forms Response Data
type ACA1095CFormResponseDataType struct {
	ACA1095CForms []ACA1095CFormsType `xml:"urn:com.workday/bsvc ACA_1095-C_Forms,omitempty"`
}

// ACA 1095-C Recipient Form
type ACA1095CFormType struct {
	EmployeeReference               *WorkerObjectType                 `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	External1095CRecipientReference *External1095CRecipientObjectType `xml:"urn:com.workday/bsvc External_1095-C_Recipient_Reference,omitempty"`
	GenerationSource                string                            `xml:"urn:com.workday/bsvc Generation_Source,omitempty"`
	IsTransmitted                   *bool                             `xml:"urn:com.workday/bsvc Is_Transmitted,omitempty"`
	ACA1095CFormData                *ACA1095CFormDataType             `xml:"urn:com.workday/bsvc ACA_1095-C_Form_Data"`
}

// ACA 1095-C Forms Data
type ACA1095CFormsType struct {
	CalendarYearReference *CalendarYearObjectType `xml:"urn:com.workday/bsvc Calendar_Year_Reference"`
	CompanyReference      *CompanyObjectType      `xml:"urn:com.workday/bsvc Company_Reference"`
	ACA1095CForm          []ACA1095CFormType      `xml:"urn:com.workday/bsvc ACA_1095-C_Form,omitempty"`
}

// Wrapper Element for the Add Dependent business process web service and its sub business processes.
type AddDependentDataType struct {
	DependentID                                               string                                               `xml:"urn:com.workday/bsvc Dependent_ID,omitempty"`
	EffectiveDate                                             *time.Time                                           `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	ReasonReference                                           *BenefitsEventSubcategoryObjectType                  `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
	EmployeeReference                                         *WorkerObjectType                                    `xml:"urn:com.workday/bsvc Employee_Reference"`
	ExistingRelatedPersonReference                            *RelatedPersonforWorkerObjectType                    `xml:"urn:com.workday/bsvc Existing_Related_Person_Reference,omitempty"`
	RelatedPersonRelationshipReference                        *RelatedPersonRelationshipObjectType                 `xml:"urn:com.workday/bsvc Related_Person_Relationship_Reference,omitempty"`
	UseEmployeeAddress                                        *bool                                                `xml:"urn:com.workday/bsvc Use_Employee_Address,omitempty"`
	UseEmployeePhone                                          *bool                                                `xml:"urn:com.workday/bsvc Use_Employee_Phone,omitempty"`
	DependentPersonalInformationData                          *RelatedPersonPersonalInformationDataType            `xml:"urn:com.workday/bsvc Dependent_Personal_Information_Data,omitempty"`
	QualifiedDomesticRelationsOrderData                       []QualifiedDomesticRelationsOrderReplacementDataType `xml:"urn:com.workday/bsvc Qualified_Domestic_Relations_Order_Data,omitempty"`
	CouldBeCoveredForHealthCareCoverageElsewhere              *bool                                                `xml:"urn:com.workday/bsvc Could_Be_Covered_For_Health_Care_Coverage_Elsewhere,omitempty"`
	CouldBeCoveredForHealthCareCoverageElsewhereEffectiveDate *time.Time                                           `xml:"urn:com.workday/bsvc Could_Be_Covered_For_Health_Care_Coverage_Elsewhere_Effective_Date,omitempty"`
}

func (t *AddDependentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AddDependentDataType
	var layout struct {
		*T
		EffectiveDate                                             *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		CouldBeCoveredForHealthCareCoverageElsewhereEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Could_Be_Covered_For_Health_Care_Coverage_Elsewhere_Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	layout.CouldBeCoveredForHealthCareCoverageElsewhereEffectiveDate = (*xsdDate)(layout.T.CouldBeCoveredForHealthCareCoverageElsewhereEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *AddDependentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AddDependentDataType
	var overlay struct {
		*T
		EffectiveDate                                             *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		CouldBeCoveredForHealthCareCoverageElsewhereEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Could_Be_Covered_For_Health_Care_Coverage_Elsewhere_Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	overlay.CouldBeCoveredForHealthCareCoverageElsewhereEffectiveDate = (*xsdDate)(overlay.T.CouldBeCoveredForHealthCareCoverageElsewhereEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper Element for the Add Dependent web service operation. The earliest effective date provided for an address will apply to all addresses, emails, web addresses, instant messengers and phone numbers.
type AddDependentRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AddDependentData          *AddDependentDataType          `xml:"urn:com.workday/bsvc Add_Dependent_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the response to an Add Dependent web service operation request.
type AddDependentResponseType struct {
	EventReference         *UniqueIdentifierObjectType                    `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	EmployeeReference      *WorkerObjectType                              `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	DependentReference     *DependentObjectType                           `xml:"urn:com.workday/bsvc Dependent_Reference,omitempty"`
	ExceptionsResponseData []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AdditionalBenefitsCoverageTargetObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AdditionalBenefitsCoverageTargetObjectType struct {
	ID         []AdditionalBenefitsCoverageTargetObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Additional Benefits Election data
type AdditionalBenefitsElectionDataType struct {
	AdditionalBenefitsPlanReference               *AdditionalBenefitsPlanObjectType                           `xml:"urn:com.workday/bsvc Additional_Benefits_Plan_Reference"`
	AdditionalBenefitsCoverageTargetReference     *AdditionalBenefitsCoverageTargetObjectType                 `xml:"urn:com.workday/bsvc Additional_Benefits_Coverage_Target_Reference,omitempty"`
	AdditionalBenefitsPercentageContributionValue float64                                                     `xml:"urn:com.workday/bsvc Additional_Benefits_Percentage_Contribution_Value,omitempty"`
	AdditionalBenefitsFlatContributionAmount      float64                                                     `xml:"urn:com.workday/bsvc Additional_Benefits_Flat_Contribution_Amount,omitempty"`
	BenefitIndividualRateData                     *BenefitIndividualRateDataforChangeBenefitsforLifeEventType `xml:"urn:com.workday/bsvc Benefit_Individual_Rate_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AdditionalBenefitsPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AdditionalBenefitsPlanObjectType struct {
	ID         []AdditionalBenefitsPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the additional names for a person, other than their legal and preferred names.  Additional names are not valid for applicants.
type AdditionalNameDataType struct {
	NameDetailData    *PersonNameDetailDataType     `xml:"urn:com.workday/bsvc Name_Detail_Data"`
	NameTypeReference *AdditionalNameTypeObjectType `xml:"urn:com.workday/bsvc Name_Type_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AdditionalNameTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AdditionalNameTypeObjectType struct {
	ID         []AdditionalNameTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains the Affordable Care Act Worker Hours And Wages Data for a worker.
type AffordableCareActWorkerHoursAndWagesDataType struct {
	DeleteInstance        *bool             `xml:"urn:com.workday/bsvc Delete_Instance,omitempty"`
	WorkerReference       *WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference"`
	TaxableWagesforPeriod float64           `xml:"urn:com.workday/bsvc Taxable_Wages_for_Period,omitempty"`
	PayEndDate            time.Time         `xml:"urn:com.workday/bsvc Pay_End_Date"`
	PaycheckDate          time.Time         `xml:"urn:com.workday/bsvc Paycheck_Date"`
	PaycheckNumber        string            `xml:"urn:com.workday/bsvc Paycheck_Number"`
	PaidHoursWorked       float64           `xml:"urn:com.workday/bsvc Paid_Hours_Worked,omitempty"`
	PaidHoursNotWorked    float64           `xml:"urn:com.workday/bsvc Paid_Hours_Not_Worked,omitempty"`
}

func (t *AffordableCareActWorkerHoursAndWagesDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AffordableCareActWorkerHoursAndWagesDataType
	var layout struct {
		*T
		PayEndDate   *xsdDate `xml:"urn:com.workday/bsvc Pay_End_Date"`
		PaycheckDate *xsdDate `xml:"urn:com.workday/bsvc Paycheck_Date"`
	}
	layout.T = (*T)(t)
	layout.PayEndDate = (*xsdDate)(&layout.T.PayEndDate)
	layout.PaycheckDate = (*xsdDate)(&layout.T.PaycheckDate)
	return e.EncodeElement(layout, start)
}
func (t *AffordableCareActWorkerHoursAndWagesDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AffordableCareActWorkerHoursAndWagesDataType
	var overlay struct {
		*T
		PayEndDate   *xsdDate `xml:"urn:com.workday/bsvc Pay_End_Date"`
		PaycheckDate *xsdDate `xml:"urn:com.workday/bsvc Paycheck_Date"`
	}
	overlay.T = (*T)(t)
	overlay.PayEndDate = (*xsdDate)(&overlay.T.PayEndDate)
	overlay.PaycheckDate = (*xsdDate)(&overlay.T.PaycheckDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type AffordableCareActWorkerHoursAndWagesObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type AffordableCareActWorkerHoursAndWagesObjectType struct {
	ID         []AffordableCareActWorkerHoursAndWagesObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Age in Years
type AgeinYearsDataType struct {
	AsOfDateData *AsOfDateDataType `xml:"urn:com.workday/bsvc As_Of_Date_Data"`
}

// If the dependent is allowed for tax deduction
type AllowedforTaxDeductionDataType struct {
	EffectiveDate          *time.Time `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	AllowedforTaxDeduction *bool      `xml:"urn:com.workday/bsvc Allowed_for_Tax_Deduction,omitempty"`
}

func (t *AllowedforTaxDeductionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AllowedforTaxDeductionDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *AllowedforTaxDeductionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AllowedforTaxDeductionDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Dependent's annual income
type AnnualIncomeDataType struct {
	EffectiveDate     *time.Time          `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	CurrencyReference *CurrencyObjectType `xml:"urn:com.workday/bsvc Currency_Reference"`
	AnnualIncome      float64             `xml:"urn:com.workday/bsvc Annual_Income,omitempty"`
}

func (t *AnnualIncomeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AnnualIncomeDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *AnnualIncomeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AnnualIncomeDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Element containing application related exceptions data
type ApplicationInstanceExceptionsDataType struct {
	ExceptionData []ExceptionDataType `xml:"urn:com.workday/bsvc Exception_Data,omitempty"`
}

// Element containing Exceptions Data
type ApplicationInstanceRelatedExceptionsDataType struct {
	ExceptionsData []ApplicationInstanceExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Data,omitempty"`
}

// As of Date
type AsOfDateDataType struct {
	SpecificMonth            float64                    `xml:"urn:com.workday/bsvc Specific_Month,omitempty"`
	SpecificDay              float64                    `xml:"urn:com.workday/bsvc Specific_Day,omitempty"`
	OfPriorPlanYear          *bool                      `xml:"urn:com.workday/bsvc Of_Prior_Plan_Year,omitempty"`
	BenefitAsOfRuleReference *BenefitAsOfRuleObjectType `xml:"urn:com.workday/bsvc Benefit_As_Of_Rule_Reference"`
}

// Contains a unique identifier for an instance of an object.
type AuthorityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AuthorityObjectType struct {
	ID         []AuthorityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the Primary or Contingent percentage to be allocated to the Beneficiary.
type BeneficaryAllocationPercentageDataType struct {
	PrimaryPercentage    float64 `xml:"urn:com.workday/bsvc Primary_Percentage"`
	ContingentPercentage float64 `xml:"urn:com.workday/bsvc Contingent_Percentage"`
}

// Beneficiary Allocation data
type BeneficiaryAllocationforChangeBenefitsDataType struct {
	BeneficiaryReference *BeneficiaryObjectType `xml:"urn:com.workday/bsvc Beneficiary_Reference"`
	PrimaryPercentage    float64                `xml:"urn:com.workday/bsvc Primary_Percentage"`
	ContingentPercentage float64                `xml:"urn:com.workday/bsvc Contingent_Percentage"`
}

// Contains a unique identifier for an instance of an object.
type BeneficiaryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BeneficiaryObjectType struct {
	ID         []BeneficiaryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for Beneficiary Person data.
type BeneficiaryPersonDataType struct {
	ExistingRelatedPersonReference           *RelatedPersonforWorkerObjectType                    `xml:"urn:com.workday/bsvc Existing_Related_Person_Reference,omitempty"`
	RelatedPersonRelationshipReference       *RelatedPersonRelationshipObjectType                 `xml:"urn:com.workday/bsvc Related_Person_Relationship_Reference,omitempty"`
	Irrevocable                              *bool                                                `xml:"urn:com.workday/bsvc Irrevocable,omitempty"`
	UseEmployeeAddress                       *bool                                                `xml:"urn:com.workday/bsvc Use_Employee_Address,omitempty"`
	UseEmployeePhone                         *bool                                                `xml:"urn:com.workday/bsvc Use_Employee_Phone,omitempty"`
	BeneficiaryPersonPersonalInformationData []BeneficiaryPersonalInformationDataType             `xml:"urn:com.workday/bsvc Beneficiary_Person_Personal_Information_Data,omitempty"`
	CourtOrder                               []QualifiedDomesticRelationsOrderReplacementDataType `xml:"urn:com.workday/bsvc Court_Order,omitempty"`
}

// Wrapper element to hold the name, contact, ID, and personal information of the Beneficiary.
type BeneficiaryPersonalInformationDataType struct {
	PersonNameData           *PersonNameDataType           `xml:"urn:com.workday/bsvc Person_Name_Data,omitempty"`
	ContactInformationData   *ContactInformationDataType   `xml:"urn:com.workday/bsvc Contact_Information_Data,omitempty"`
	PersonIdentificationData *PersonIdentificationDataType `xml:"urn:com.workday/bsvc Person_Identification_Data,omitempty"`
	DateofBirth              *time.Time                    `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	DateofDeath              *time.Time                    `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
	GenderReference          *GenderObjectType             `xml:"urn:com.workday/bsvc Gender_Reference,omitempty"`
}

func (t *BeneficiaryPersonalInformationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BeneficiaryPersonalInformationDataType
	var layout struct {
		*T
		DateofBirth *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
		DateofDeath *xsdDate `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateofBirth = (*xsdDate)(layout.T.DateofBirth)
	layout.DateofDeath = (*xsdDate)(layout.T.DateofDeath)
	return e.EncodeElement(layout, start)
}
func (t *BeneficiaryPersonalInformationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BeneficiaryPersonalInformationDataType
	var overlay struct {
		*T
		DateofBirth *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
		DateofDeath *xsdDate `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateofBirth = (*xsdDate)(overlay.T.DateofBirth)
	overlay.DateofDeath = (*xsdDate)(overlay.T.DateofDeath)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for Beneficiary Trust data.
type BeneficiaryTrustInformationDataType struct {
	TrustName                      string                              `xml:"urn:com.workday/bsvc Trust_Name"`
	TaxID                          string                              `xml:"urn:com.workday/bsvc Tax_ID,omitempty"`
	TrustDate                      *time.Time                          `xml:"urn:com.workday/bsvc Trust_Date,omitempty"`
	TrusteePersonalInformationData *TrusteePersonalInformationDataType `xml:"urn:com.workday/bsvc Trustee_Personal_Information_Data,omitempty"`
}

func (t *BeneficiaryTrustInformationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BeneficiaryTrustInformationDataType
	var layout struct {
		*T
		TrustDate *xsdDate `xml:"urn:com.workday/bsvc Trust_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TrustDate = (*xsdDate)(layout.T.TrustDate)
	return e.EncodeElement(layout, start)
}
func (t *BeneficiaryTrustInformationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BeneficiaryTrustInformationDataType
	var overlay struct {
		*T
		TrustDate *xsdDate `xml:"urn:com.workday/bsvc Trust_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TrustDate = (*xsdDate)(overlay.T.TrustDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains the Benefit Annual Credit data.
type BenefitAnnualCreditDataType struct {
	EffectiveDate                        time.Time                          `xml:"urn:com.workday/bsvc Effective_Date"`
	WorkerReference                      *WorkerObjectType                  `xml:"urn:com.workday/bsvc Worker_Reference"`
	BenefitAnnualCredit                  float64                            `xml:"urn:com.workday/bsvc Benefit_Annual_Credit,omitempty"`
	BenefitAnnualCreditCurrencyReference *CurrencyObjectType                `xml:"urn:com.workday/bsvc Benefit_Annual_Credit_Currency_Reference"`
	BenefitAnnualCreditTypeReference     *BenefitAnnualCreditTypeObjectType `xml:"urn:com.workday/bsvc Benefit_Annual_Credit_Type_Reference"`
}

func (t *BenefitAnnualCreditDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BenefitAnnualCreditDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *BenefitAnnualCreditDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BenefitAnnualCreditDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element containing the list of Benefit Annual Credits for which data is requested.
type BenefitAnnualCreditRequestCriteriaType struct {
	WorkerReference []WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
}

// Wrapper element for Benefit Annual Credits Response.
type BenefitAnnualCreditResponseDataType struct {
	BenefitAnnualCredit []BenefitAnnualCreditType `xml:"urn:com.workday/bsvc Benefit_Annual_Credit,omitempty"`
}

// Wrapper element for Benefit Annual Credit.
type BenefitAnnualCreditType struct {
	BenefitAnnualCreditReference *UniqueIdentifierObjectType  `xml:"urn:com.workday/bsvc Benefit_Annual_Credit_Reference"`
	BenefitAnnualCreditData      *BenefitAnnualCreditDataType `xml:"urn:com.workday/bsvc Benefit_Annual_Credit_Data"`
}

// Contains a unique identifier for an instance of an object.
type BenefitAnnualCreditTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitAnnualCreditTypeObjectType struct {
	ID         []BenefitAnnualCreditTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the Benefit Annual Rate data.
type BenefitAnnualRateDataType struct {
	EffectiveDate                  time.Time                        `xml:"urn:com.workday/bsvc Effective_Date"`
	WorkerReference                *WorkerObjectType                `xml:"urn:com.workday/bsvc Worker_Reference"`
	BenefitAnnualRate              float64                          `xml:"urn:com.workday/bsvc Benefit_Annual_Rate,omitempty"`
	CurrencyReference              *CurrencyObjectType              `xml:"urn:com.workday/bsvc Currency_Reference"`
	BenefitAnnualRateTypeReference *BenefitAnnualRateTypeObjectType `xml:"urn:com.workday/bsvc Benefit_Annual_Rate_Type_Reference,omitempty"`
}

func (t *BenefitAnnualRateDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BenefitAnnualRateDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *BenefitAnnualRateDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BenefitAnnualRateDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element containing the list of Benefit Annual Rates for which data is requested.
type BenefitAnnualRateRequestCriteriaType struct {
	WorkerReference []WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
}

// Wrapper element for Benefit Annual Rates Response.
type BenefitAnnualRateResponseDataType struct {
	BenefitAnnualRate []BenefitAnnualRateType `xml:"urn:com.workday/bsvc Benefit_Annual_Rate,omitempty"`
}

// Wrapper element for Benefit Annual Rate.
type BenefitAnnualRateType struct {
	BenefitAnnualRateReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Benefit_Annual_Rate_Reference"`
	BenefitAnnualRateData      *BenefitAnnualRateDataType  `xml:"urn:com.workday/bsvc Benefit_Annual_Rate_Data"`
}

// Contains a unique identifier for an instance of an object.
type BenefitAnnualRateTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitAnnualRateTypeObjectType struct {
	ID         []BenefitAnnualRateTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BenefitAsOfRuleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitAsOfRuleObjectType struct {
	ID         []BenefitAsOfRuleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BenefitCoverageTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitCoverageTypeObjectType struct {
	ID         []BenefitCoverageTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Benefit coverage type level values
type BenefitCoverageTypeforRetirementSavingsElectionsDataType struct {
	BenefitCoverageTypeReference                 *BenefitCoverageTypeObjectType                     `xml:"urn:com.workday/bsvc Benefit_Coverage_Type_Reference"`
	EmployeeContributionPercentage               float64                                            `xml:"urn:com.workday/bsvc Employee_Contribution_Percentage,omitempty"`
	EmployeeContributionAmount                   float64                                            `xml:"urn:com.workday/bsvc Employee_Contribution_Amount,omitempty"`
	Donotautoenrollinunspecifiedplans            *bool                                              `xml:"urn:com.workday/bsvc Do_not_auto_enroll_in_unspecified_plans,omitempty"`
	RetirementSavingsElectionforCoverageTypeData []RetirementSavingsElectionforCoverageTypeDataType `xml:"urn:com.workday/bsvc Retirement_Savings_Election_for_Coverage_Type_Data,omitempty"`
}

// Generic benefit election data for this election
type BenefitElectionChangeBenefitsDataType struct {
	CoverageBeginDate                *time.Time                                `xml:"urn:com.workday/bsvc Coverage_Begin_Date,omitempty"`
	OriginalCoverageBeginDate        *time.Time                                `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
	DeductionBeginDate               *time.Time                                `xml:"urn:com.workday/bsvc Deduction_Begin_Date,omitempty"`
	OriginalDeductionBeginDate       *time.Time                                `xml:"urn:com.workday/bsvc Original_Deduction_Begin_Date,omitempty"`
	HealthCareElectionData           *HealthCareElectionChangeBenefitsDataType `xml:"urn:com.workday/bsvc Health_Care_Election_Data"`
	HealthSavingsAccountElectionData *HealthSavingsAccountElectionDataType     `xml:"urn:com.workday/bsvc Health_Savings_Account_Election_Data"`
	SpendingAccountElectionData      *SpendingAccountElectionDataType          `xml:"urn:com.workday/bsvc Spending_Account_Election_Data"`
	InsuranceElectionData            *InsuranceElectionDataChangeBenefitsType  `xml:"urn:com.workday/bsvc Insurance_Election_Data"`
	RetirementSavingsElectionData    *RetirementSavingsElectionDataType        `xml:"urn:com.workday/bsvc Retirement_Savings_Election_Data"`
	AdditionalBenefitsElectionData   *AdditionalBenefitsElectionDataType       `xml:"urn:com.workday/bsvc Additional_Benefits_Election_Data"`
}

func (t *BenefitElectionChangeBenefitsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BenefitElectionChangeBenefitsDataType
	var layout struct {
		*T
		CoverageBeginDate          *xsdDate `xml:"urn:com.workday/bsvc Coverage_Begin_Date,omitempty"`
		OriginalCoverageBeginDate  *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
		DeductionBeginDate         *xsdDate `xml:"urn:com.workday/bsvc Deduction_Begin_Date,omitempty"`
		OriginalDeductionBeginDate *xsdDate `xml:"urn:com.workday/bsvc Original_Deduction_Begin_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CoverageBeginDate = (*xsdDate)(layout.T.CoverageBeginDate)
	layout.OriginalCoverageBeginDate = (*xsdDate)(layout.T.OriginalCoverageBeginDate)
	layout.DeductionBeginDate = (*xsdDate)(layout.T.DeductionBeginDate)
	layout.OriginalDeductionBeginDate = (*xsdDate)(layout.T.OriginalDeductionBeginDate)
	return e.EncodeElement(layout, start)
}
func (t *BenefitElectionChangeBenefitsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BenefitElectionChangeBenefitsDataType
	var overlay struct {
		*T
		CoverageBeginDate          *xsdDate `xml:"urn:com.workday/bsvc Coverage_Begin_Date,omitempty"`
		OriginalCoverageBeginDate  *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
		DeductionBeginDate         *xsdDate `xml:"urn:com.workday/bsvc Deduction_Begin_Date,omitempty"`
		OriginalDeductionBeginDate *xsdDate `xml:"urn:com.workday/bsvc Original_Deduction_Begin_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CoverageBeginDate = (*xsdDate)(overlay.T.CoverageBeginDate)
	overlay.OriginalCoverageBeginDate = (*xsdDate)(overlay.T.OriginalCoverageBeginDate)
	overlay.DeductionBeginDate = (*xsdDate)(overlay.T.DeductionBeginDate)
	overlay.OriginalDeductionBeginDate = (*xsdDate)(overlay.T.OriginalDeductionBeginDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type BenefitEventTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitEventTypeObjectType struct {
	ID         []BenefitEventTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Benefit Individual Rate Data contains all cost values for an individual rate.
type BenefitIndividualRateDataType struct {
	BenefitIndividualRateID string  `xml:"urn:com.workday/bsvc Benefit_Individual_Rate_ID"`
	EmployeeCostPreTax      float64 `xml:"urn:com.workday/bsvc Employee_Cost_PreTax,omitempty"`
	EmployeePostTaxAmount   float64 `xml:"urn:com.workday/bsvc Employee_PostTax_Amount,omitempty"`
	EmployerCostNonTaxable  float64 `xml:"urn:com.workday/bsvc Employer_Cost_NonTaxable,omitempty"`
	EmployerCostTaxable     float64 `xml:"urn:com.workday/bsvc Employer_Cost_Taxable,omitempty"`
}

// Data for Benefit Individual Rate
type BenefitIndividualRateDataforChangeBenefitsforLifeEventType struct {
	BenefitIndividualRateID string  `xml:"urn:com.workday/bsvc Benefit_Individual_Rate_ID,omitempty"`
	EmployeeCostPreTax      float64 `xml:"urn:com.workday/bsvc Employee_Cost_PreTax,omitempty"`
	EmployeeCostPostTax     float64 `xml:"urn:com.workday/bsvc Employee_Cost_PostTax,omitempty"`
	EmployerCostNonTaxable  float64 `xml:"urn:com.workday/bsvc Employer_Cost_NonTaxable,omitempty"`
	EmployerCostTaxable     float64 `xml:"urn:com.workday/bsvc Employer_Cost_Taxable,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BenefitIndividualRateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitIndividualRateObjectType struct {
	ID         []BenefitIndividualRateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The "Request Criteria" element for the web service operation contains the filtering logic to limit the data that is returned in the response.
type BenefitIndividualRateRequestCriteriaType struct {
	WorkerReference []WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
}

// A list of references to retrieve using the get operation.
type BenefitIndividualRateRequestReferencesType struct {
	BenefitIndividualRateReference []BenefitIndividualRateObjectType `xml:"urn:com.workday/bsvc Benefit_Individual_Rate_Reference"`
}

// Response for a Get Benefits Individual Rate web service request
type BenefitIndividualRateResponseDataType struct {
	BenefitIndividualRate []BenefitIndividualRateType `xml:"urn:com.workday/bsvc Benefit_Individual_Rate,omitempty"`
}

// Benefit Individual Rate
type BenefitIndividualRateType struct {
	BenefitIndividualRateReference *BenefitIndividualRateObjectType `xml:"urn:com.workday/bsvc Benefit_Individual_Rate_Reference"`
	BenefitIndividualRateData      *BenefitIndividualRateDataType   `xml:"urn:com.workday/bsvc Benefit_Individual_Rate_Data"`
}

// Container for specifying the benefit jobs for a worker.
type BenefitJobsDataType struct {
	UsePrimaryJob     bool                        `xml:"urn:com.workday/bsvc Use_Primary_Job"`
	PositionReference []PositionElementObjectType `xml:"urn:com.workday/bsvc Position_Reference"`
	EffectiveDate     time.Time                   `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
}

func (t *BenefitJobsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BenefitJobsDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *BenefitJobsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BenefitJobsDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type BenefitPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitPlanObjectType struct {
	ID         []BenefitPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BenefitRateBandTobaccoUseObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitRateBandTobaccoUseObjectType struct {
	ID         []BenefitRateBandTobaccoUseObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Benefit Rate Flat Value Data
type BenefitRateValueFlatDataType struct {
	EmployeeCostAmount                float64                              `xml:"urn:com.workday/bsvc Employee_Cost_Amount,omitempty"`
	EmployeeCostAmountMin             float64                              `xml:"urn:com.workday/bsvc Employee_Cost_Amount_Min,omitempty"`
	EmployeeCostAmountMax             float64                              `xml:"urn:com.workday/bsvc Employee_Cost_Amount_Max,omitempty"`
	EmployeeCostAmountPostTax         float64                              `xml:"urn:com.workday/bsvc Employee_Cost_Amount_PostTax,omitempty"`
	EmployeeCostAmountMinPostTax      float64                              `xml:"urn:com.workday/bsvc Employee_Cost_Amount_Min_PostTax,omitempty"`
	EmployeeCostAmountMaxPostTax      float64                              `xml:"urn:com.workday/bsvc Employee_Cost_Amount_Max_PostTax,omitempty"`
	EmployerContributionAmount        float64                              `xml:"urn:com.workday/bsvc Employer_Contribution_Amount,omitempty"`
	EmployerContributionAmountTaxable float64                              `xml:"urn:com.workday/bsvc Employer_Contribution_Amount_Taxable,omitempty"`
	FlexCreditAmount                  float64                              `xml:"urn:com.workday/bsvc Flex_Credit_Amount,omitempty"`
	FairMarketValueperDependent       float64                              `xml:"urn:com.workday/bsvc Fair_Market_Value_per_Dependent,omitempty"`
	BenefitCreditReference            []BenefitsCalculableCreditObjectType `xml:"urn:com.workday/bsvc Benefit_Credit_Reference,omitempty"`
	BenefitSurchargeReference         []BenefitsCalculableCreditObjectType `xml:"urn:com.workday/bsvc Benefit_Surcharge_Reference,omitempty"`
}

// Benefit Rate Value Percent of Salary
type BenefitRateValuePercentofSalaryDataType struct {
	EmployeeCostPercentage        float64                              `xml:"urn:com.workday/bsvc Employee_Cost_Percentage,omitempty"`
	EmployeeCostAmountMin         float64                              `xml:"urn:com.workday/bsvc Employee_Cost_Amount_Min,omitempty"`
	EmployeeCostAmountMax         float64                              `xml:"urn:com.workday/bsvc Employee_Cost_Amount_Max,omitempty"`
	EmployeeCostPercentagePostTax float64                              `xml:"urn:com.workday/bsvc Employee_Cost_Percentage_PostTax,omitempty"`
	EmployeeCostAmountMinPostTax  float64                              `xml:"urn:com.workday/bsvc Employee_Cost_Amount_Min_PostTax,omitempty"`
	EmployeeCostAmountMaxPostTax  float64                              `xml:"urn:com.workday/bsvc Employee_Cost_Amount_Max_PostTax,omitempty"`
	TotalCostNonTaxable           float64                              `xml:"urn:com.workday/bsvc Total_Cost_NonTaxable,omitempty"`
	TotalCostTaxable              float64                              `xml:"urn:com.workday/bsvc Total_Cost_Taxable,omitempty"`
	FlexCreditAmount              float64                              `xml:"urn:com.workday/bsvc Flex_Credit_Amount,omitempty"`
	FairMarketValueperDependent   float64                              `xml:"urn:com.workday/bsvc Fair_Market_Value_per_Dependent,omitempty"`
	FlatAmountData                []FlatAmountDataType                 `xml:"urn:com.workday/bsvc Flat_Amount_Data,omitempty"`
	BenefitCreditReference        []BenefitsCalculableCreditObjectType `xml:"urn:com.workday/bsvc Benefit_Credit_Reference,omitempty"`
	BenefitSurchargeReference     []BenefitsCalculableCreditObjectType `xml:"urn:com.workday/bsvc Benefit_Surcharge_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BenefitRoundDirectionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitRoundDirectionObjectType struct {
	ID         []BenefitRoundDirectionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains Benefit Rounding Rule Final Data
type BenefitRoundingRuleFinalDataType struct {
	RoundAmountTo              float64                          `xml:"urn:com.workday/bsvc Round_Amount_To,omitempty"`
	RoundingDirectionReference *BenefitRoundDirectionObjectType `xml:"urn:com.workday/bsvc Rounding_Direction_Reference,omitempty"`
	RoundonNumber              float64                          `xml:"urn:com.workday/bsvc Round_on_Number,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BenefitSalarySourceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitSalarySourceObjectType struct {
	ID         []BenefitSalarySourceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BenefitsCalculableCreditObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitsCalculableCreditObjectType struct {
	ID         []BenefitsCalculableCreditObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BenefitsEventSubcategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitsEventSubcategoryObjectType struct {
	ID         []BenefitsEventSubcategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element for all Biographical data (e.g. DOB, Gender, etc.) associated with a person.
type BiographicDataType struct {
	CountryOfBirthReference *CountryOfBirthReferenceType `xml:"urn:com.workday/bsvc Country_Of_Birth_Reference,omitempty"`
	PlaceOfBirth            string                       `xml:"urn:com.workday/bsvc Place_Of_Birth,omitempty"`
	DateOfBirth             *time.Time                   `xml:"urn:com.workday/bsvc Date_Of_Birth,omitempty"`
	GenderReference         *GenderReferenceType         `xml:"urn:com.workday/bsvc Gender_Reference,omitempty"`
	DisabilityReference     []DisabilityReferenceType    `xml:"urn:com.workday/bsvc Disability_Reference,omitempty"`
	UsesTobacco             *bool                        `xml:"urn:com.workday/bsvc Uses_Tobacco,omitempty"`
}

func (t *BiographicDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BiographicDataType
	var layout struct {
		*T
		DateOfBirth *xsdDate `xml:"urn:com.workday/bsvc Date_Of_Birth,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateOfBirth = (*xsdDate)(layout.T.DateOfBirth)
	return e.EncodeElement(layout, start)
}
func (t *BiographicDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BiographicDataType
	var overlay struct {
		*T
		DateOfBirth *xsdDate `xml:"urn:com.workday/bsvc Date_Of_Birth,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateOfBirth = (*xsdDate)(overlay.T.DateOfBirth)
	return d.DecodeElement(&overlay, &start)
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
type COBRAEligibilityReasonObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type COBRAEligibilityReasonObjectType struct {
	ID         []COBRAEligibilityReasonObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The information (data) for a single COBRA Eligiblity Status
type COBRAParticipantEligibilityStatusDataType struct {
	COBRAParticipantReference  *COBRAParticipantObjectType       `xml:"urn:com.workday/bsvc COBRA_Participant_Reference"`
	EligibilityReasonReference *COBRAEligibilityReasonObjectType `xml:"urn:com.workday/bsvc Eligibility_Reason_Reference"`
	EligibleDate               time.Time                         `xml:"urn:com.workday/bsvc Eligible_Date"`
	LossofCoverageDate         *time.Time                        `xml:"urn:com.workday/bsvc Loss_of_Coverage_Date,omitempty"`
	CoverageEndDate            *time.Time                        `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	BenefitPlanReference       *BenefitPlanObjectType            `xml:"urn:com.workday/bsvc Benefit_Plan_Reference,omitempty"`
}

func (t *COBRAParticipantEligibilityStatusDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T COBRAParticipantEligibilityStatusDataType
	var layout struct {
		*T
		EligibleDate       *xsdDate `xml:"urn:com.workday/bsvc Eligible_Date"`
		LossofCoverageDate *xsdDate `xml:"urn:com.workday/bsvc Loss_of_Coverage_Date,omitempty"`
		CoverageEndDate    *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EligibleDate = (*xsdDate)(&layout.T.EligibleDate)
	layout.LossofCoverageDate = (*xsdDate)(layout.T.LossofCoverageDate)
	layout.CoverageEndDate = (*xsdDate)(layout.T.CoverageEndDate)
	return e.EncodeElement(layout, start)
}
func (t *COBRAParticipantEligibilityStatusDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T COBRAParticipantEligibilityStatusDataType
	var overlay struct {
		*T
		EligibleDate       *xsdDate `xml:"urn:com.workday/bsvc Eligible_Date"`
		LossofCoverageDate *xsdDate `xml:"urn:com.workday/bsvc Loss_of_Coverage_Date,omitempty"`
		CoverageEndDate    *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EligibleDate = (*xsdDate)(&overlay.T.EligibleDate)
	overlay.LossofCoverageDate = (*xsdDate)(overlay.T.LossofCoverageDate)
	overlay.CoverageEndDate = (*xsdDate)(overlay.T.CoverageEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type COBRAParticipantObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type COBRAParticipantObjectType struct {
	ID         []COBRAParticipantObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Wrapper element to hold the Beneficiary data for an Employee.
type ChangeBeneficiaryBusinessProcessDataType struct {
	EmployeeReference    *EmployeeObjectType                  `xml:"urn:com.workday/bsvc Employee_Reference"`
	BeneficiaryID        string                               `xml:"urn:com.workday/bsvc Beneficiary_ID,omitempty"`
	BeneficiaryReference *BeneficiaryObjectType               `xml:"urn:com.workday/bsvc Beneficiary_Reference,omitempty"`
	PersonData           *BeneficiaryPersonDataType           `xml:"urn:com.workday/bsvc Person_Data,omitempty"`
	TrustData            *BeneficiaryTrustInformationDataType `xml:"urn:com.workday/bsvc Trust_Data,omitempty"`
	Delete               bool                                 `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// This web service allows the adding and updating of a Beneficiary for an Employee.
type ChangeBeneficiaryRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType            `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	ChangeBeneficiaryData     *ChangeBeneficiaryBusinessProcessDataType `xml:"urn:com.workday/bsvc Change_Beneficiary_Data"`
	Version                   string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response for the change beneficiary web service.
type ChangeBeneficiaryResponseType struct {
	BeneficiaryEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Beneficiary_Event_Reference,omitempty"`
	Version                   string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the request to change a worker's benefit jobs. Changing a worker's benefit jobs changes how the CRFs for benefit eligibility behave. If a worker has benefit jobs specified, then only the specified jobs will be used in the CRF calculations used to evaluate their benefit eligibility. If benefit jobs are not specified, then the HR Primary job is used for those CRFs.
type ChangeBenefitJobsRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	WorkerBenefitJobsData     *WorkerBenefitJobsDataType     `xml:"urn:com.workday/bsvc Worker_Benefit_Jobs_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response to the Add Worker Benefit Jobs web service.
type ChangeBenefitJobsResponseType struct {
	EventReference        *UniqueIdentifierObjectType                    `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	WorkerReference       *WorkerObjectType                              `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	ExceptionResponseData []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exception_Response_Data,omitempty"`
	Version               string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Data for life change events
type ChangeBenefitsDataType struct {
	EmployeeReference                 *WorkerObjectType                                          `xml:"urn:com.workday/bsvc Employee_Reference"`
	BenefitEventTypeReference         []BenefitEventTypeObjectType                               `xml:"urn:com.workday/bsvc Benefit_Event_Type_Reference"`
	EventDate                         time.Time                                                  `xml:"urn:com.workday/bsvc Event_Date"`
	BenefitDeductionPeriodsRemaining  float64                                                    `xml:"urn:com.workday/bsvc Benefit_Deduction_Periods_Remaining,omitempty"`
	RemainingPeriodFrequencyReference *FrequencyObjectType                                       `xml:"urn:com.workday/bsvc Remaining_Period_Frequency_Reference,omitempty"`
	ExcessCreditstoHSA                float64                                                    `xml:"urn:com.workday/bsvc Excess_Credits_to_HSA,omitempty"`
	HasExistingMedicalCoverage        *bool                                                      `xml:"urn:com.workday/bsvc Has_Existing_Medical_Coverage,omitempty"`
	EnrollmentSignatureDate           *time.Time                                                 `xml:"urn:com.workday/bsvc Enrollment_Signature_Date,omitempty"`
	SigningWorkerReference            *WorkerObjectType                                          `xml:"urn:com.workday/bsvc Signing_Worker_Reference,omitempty"`
	BenefitElectionData               []BenefitElectionChangeBenefitsDataType                    `xml:"urn:com.workday/bsvc Benefit_Election_Data,omitempty"`
	RetirementSavingsCoverageTypeData []BenefitCoverageTypeforRetirementSavingsElectionsDataType `xml:"urn:com.workday/bsvc Retirement_Savings_Coverage_Type_Data,omitempty"`
	BiographicData                    []BiographicDataType                                       `xml:"urn:com.workday/bsvc Biographic_Data,omitempty"`
}

func (t *ChangeBenefitsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ChangeBenefitsDataType
	var layout struct {
		*T
		EventDate               *xsdDate `xml:"urn:com.workday/bsvc Event_Date"`
		EnrollmentSignatureDate *xsdDate `xml:"urn:com.workday/bsvc Enrollment_Signature_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EventDate = (*xsdDate)(&layout.T.EventDate)
	layout.EnrollmentSignatureDate = (*xsdDate)(layout.T.EnrollmentSignatureDate)
	return e.EncodeElement(layout, start)
}
func (t *ChangeBenefitsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ChangeBenefitsDataType
	var overlay struct {
		*T
		EventDate               *xsdDate `xml:"urn:com.workday/bsvc Event_Date"`
		EnrollmentSignatureDate *xsdDate `xml:"urn:com.workday/bsvc Enrollment_Signature_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EventDate = (*xsdDate)(&overlay.T.EventDate)
	overlay.EnrollmentSignatureDate = (*xsdDate)(overlay.T.EnrollmentSignatureDate)
	return d.DecodeElement(&overlay, &start)
}

// Change benefits request
type ChangeBenefitsRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	ChangeBenefitsData        *ChangeBenefitsDataType        `xml:"urn:com.workday/bsvc Change_Benefits_Data,omitempty"`
	AddOnly                   bool                           `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response for a Change Benefits for Life Event web service request
type ChangeBenefitsResponseType struct {
	ChangeBenefitsReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Change_Benefits_Reference,omitempty"`
	Version                 string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CitizenshipStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CitizenshipStatusObjectType struct {
	ID         []CitizenshipStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type CompanyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompanyObjectType struct {
	ID         []CompanyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompensationPayEarningObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationPayEarningObjectType struct {
	ID         []CompensationPayEarningObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// All of the person's contact data (address, phone, email, instant messenger, web address).
type ContactInformationDataForRelatedPersonType struct {
	AddressData          []AddressInformationDataType          `xml:"urn:com.workday/bsvc Address_Data,omitempty"`
	PhoneData            []PhoneInformationDataType            `xml:"urn:com.workday/bsvc Phone_Data,omitempty"`
	EmailAddressData     []EmailAddressInformationDataType     `xml:"urn:com.workday/bsvc Email_Address_Data,omitempty"`
	InstantMessengerData []InstantMessengerInformationDataType `xml:"urn:com.workday/bsvc Instant_Messenger_Data,omitempty"`
	WebAddressData       []WebAddressInformationDataType       `xml:"urn:com.workday/bsvc Web_Address_Data,omitempty"`
}

// All of the person's contact data (address, phone, email, instant messenger, web address).
type ContactInformationDataType struct {
	AddressData          []AddressInformationDataType          `xml:"urn:com.workday/bsvc Address_Data,omitempty"`
	PhoneData            []PhoneInformationDataType            `xml:"urn:com.workday/bsvc Phone_Data,omitempty"`
	EmailAddressData     []EmailAddressInformationDataType     `xml:"urn:com.workday/bsvc Email_Address_Data,omitempty"`
	InstantMessengerData []InstantMessengerInformationDataType `xml:"urn:com.workday/bsvc Instant_Messenger_Data,omitempty"`
	WebAddressData       []WebAddressInformationDataType       `xml:"urn:com.workday/bsvc Web_Address_Data,omitempty"`
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

// Reference element representing a unique instance of Country of Birth.
type CountryOfBirthReferenceType struct {
	CountryReference *CountryReferenceType `xml:"urn:com.workday/bsvc Country_Reference"`
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

// Contains a unique identifier for an instance of an object.
type CurrencyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CurrencyObjectType struct {
	ID         []CurrencyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper for Custom Identifier Data.
type CustomIDDataType struct {
	ID                      string                  `xml:"urn:com.workday/bsvc ID,omitempty"`
	IDTypeReference         *CustomIDTypeObjectType `xml:"urn:com.workday/bsvc ID_Type_Reference,omitempty"`
	IssuedDate              *time.Time              `xml:"urn:com.workday/bsvc Issued_Date,omitempty"`
	ExpirationDate          *time.Time              `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	OrganizationIDReference *OrganizationObjectType `xml:"urn:com.workday/bsvc Organization_ID_Reference,omitempty"`
	CustomDescription       string                  `xml:"urn:com.workday/bsvc Custom_Description,omitempty"`
}

func (t *CustomIDDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CustomIDDataType
	var layout struct {
		*T
		IssuedDate     *xsdDate `xml:"urn:com.workday/bsvc Issued_Date,omitempty"`
		ExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.IssuedDate = (*xsdDate)(layout.T.IssuedDate)
	layout.ExpirationDate = (*xsdDate)(layout.T.ExpirationDate)
	return e.EncodeElement(layout, start)
}
func (t *CustomIDDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CustomIDDataType
	var overlay struct {
		*T
		IssuedDate     *xsdDate `xml:"urn:com.workday/bsvc Issued_Date,omitempty"`
		ExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.IssuedDate = (*xsdDate)(overlay.T.IssuedDate)
	overlay.ExpirationDate = (*xsdDate)(overlay.T.ExpirationDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element for all Custom Identifier data.
type CustomIDType struct {
	CustomIDReference       *UniqueIdentifierObjectType          `xml:"urn:com.workday/bsvc Custom_ID_Reference,omitempty"`
	CustomIDData            *CustomIDDataType                    `xml:"urn:com.workday/bsvc Custom_ID_Data,omitempty"`
	CustomIDSharedReference *CustomIdentifierReferenceObjectType `xml:"urn:com.workday/bsvc Custom_ID_Shared_Reference,omitempty"`
	Delete                  bool                                 `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomIDTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomIDTypeObjectType struct {
	ID         []CustomIDTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomIdentifierReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomIdentifierReferenceObjectType struct {
	ID         []CustomIdentifierReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Defines the Beneficiary recipient and the allocation percentage for the Retirement Savings Election.  This is required when the retirement savings plan requires beneficiaries and the election amount or election percentage is populated (i.e. greater than 0).
type DefinedContributionBeneficiaryAllocationDataType struct {
	BeneficiaryReference                *BeneficiaryObjectType                   `xml:"urn:com.workday/bsvc Beneficiary_Reference"`
	BeneficiaryAllocationPercentageData []BeneficaryAllocationPercentageDataType `xml:"urn:com.workday/bsvc Beneficiary_Allocation_Percentage_Data,omitempty"`
}

// Contains election percent or amount and beneficiary allocation information.  If the Election Percentage and Election Amount are left blank, then an existing election is removed.  If there is no existing election for the benefit plan and the Election Percentage and Election Amount are left blank, then no election will be created.
type DefinedContributionElectionDataType struct {
	DefinedContributionPlanReference *RetirementSavingsPlanObjectType                   `xml:"urn:com.workday/bsvc Defined_Contribution_Plan_Reference"`
	ElectionPercentage               float64                                            `xml:"urn:com.workday/bsvc Election_Percentage,omitempty"`
	ElectionAmount                   float64                                            `xml:"urn:com.workday/bsvc Election_Amount,omitempty"`
	CurrencyReference                *CurrencyObjectType                                `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference               *FrequencyObjectType                               `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	BeneficiaryAllocationData        []DefinedContributionBeneficiaryAllocationDataType `xml:"urn:com.workday/bsvc Beneficiary_Allocation_Data,omitempty"`
}

// Contains the retirement savings election package reference and it's detailed information.
type DefinedContributionPackageType struct {
	EmployeeDefinedContributionElectionData *EmployeeDefinedContributionElectionDataType `xml:"urn:com.workday/bsvc Employee_Defined_Contribution_Election_Data,omitempty"`
}

// Contains the Data for the Dependent Election
type DependentElectionDataType struct {
	DependentReference        *DependentObjectType `xml:"urn:com.workday/bsvc Dependent_Reference"`
	ProviderID                string               `xml:"urn:com.workday/bsvc Provider_ID,omitempty"`
	OriginalCoverageBeginDate *time.Time           `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
}

func (t *DependentElectionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DependentElectionDataType
	var layout struct {
		*T
		OriginalCoverageBeginDate *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.OriginalCoverageBeginDate = (*xsdDate)(layout.T.OriginalCoverageBeginDate)
	return e.EncodeElement(layout, start)
}
func (t *DependentElectionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DependentElectionDataType
	var overlay struct {
		*T
		OriginalCoverageBeginDate *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.OriginalCoverageBeginDate = (*xsdDate)(overlay.T.OriginalCoverageBeginDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type DependentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DependentObjectType struct {
	ID         []DependentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for this dependent's personal information.
type DependentPersonalInformationDataType struct {
	PersonNameData           *PersonNameDataType           `xml:"urn:com.workday/bsvc Person_Name_Data,omitempty"`
	ContactInformationData   *ContactInformationDataType   `xml:"urn:com.workday/bsvc Contact_Information_Data,omitempty"`
	PersonIdentificationData *PersonIdentificationDataType `xml:"urn:com.workday/bsvc Person_Identification_Data,omitempty"`
	DateofBirth              *time.Time                    `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	DateofDeath              *time.Time                    `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
	GenderReference          *GenderObjectType             `xml:"urn:com.workday/bsvc Gender_Reference,omitempty"`
	UsesTobacco              *bool                         `xml:"urn:com.workday/bsvc Uses_Tobacco,omitempty"`
	FullTimeStudent          *bool                         `xml:"urn:com.workday/bsvc Full-Time_Student,omitempty"`
	StudentStatusStartDate   *time.Time                    `xml:"urn:com.workday/bsvc Student_Status_Start_Date,omitempty"`
	StudentStatusEndDate     *time.Time                    `xml:"urn:com.workday/bsvc Student_Status_End_Date,omitempty"`
	Disabled                 *bool                         `xml:"urn:com.workday/bsvc Disabled,omitempty"`
	InactiveDate             *time.Time                    `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
}

func (t *DependentPersonalInformationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DependentPersonalInformationDataType
	var layout struct {
		*T
		DateofBirth            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
		DateofDeath            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
		StudentStatusStartDate *xsdDate `xml:"urn:com.workday/bsvc Student_Status_Start_Date,omitempty"`
		StudentStatusEndDate   *xsdDate `xml:"urn:com.workday/bsvc Student_Status_End_Date,omitempty"`
		InactiveDate           *xsdDate `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateofBirth = (*xsdDate)(layout.T.DateofBirth)
	layout.DateofDeath = (*xsdDate)(layout.T.DateofDeath)
	layout.StudentStatusStartDate = (*xsdDate)(layout.T.StudentStatusStartDate)
	layout.StudentStatusEndDate = (*xsdDate)(layout.T.StudentStatusEndDate)
	layout.InactiveDate = (*xsdDate)(layout.T.InactiveDate)
	return e.EncodeElement(layout, start)
}
func (t *DependentPersonalInformationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DependentPersonalInformationDataType
	var overlay struct {
		*T
		DateofBirth            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
		DateofDeath            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
		StudentStatusStartDate *xsdDate `xml:"urn:com.workday/bsvc Student_Status_Start_Date,omitempty"`
		StudentStatusEndDate   *xsdDate `xml:"urn:com.workday/bsvc Student_Status_End_Date,omitempty"`
		InactiveDate           *xsdDate `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateofBirth = (*xsdDate)(overlay.T.DateofBirth)
	overlay.DateofDeath = (*xsdDate)(overlay.T.DateofDeath)
	overlay.StudentStatusStartDate = (*xsdDate)(overlay.T.StudentStatusStartDate)
	overlay.StudentStatusEndDate = (*xsdDate)(overlay.T.StudentStatusEndDate)
	overlay.InactiveDate = (*xsdDate)(overlay.T.InactiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type DisabilityGradeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DisabilityGradeObjectType struct {
	ID         []DisabilityGradeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for all Disability Status data for the Dependent.
type DisabilityInformationDataforRelatedPersonType struct {
	ReplaceAll                      *bool                                                 `xml:"urn:com.workday/bsvc Replace_All,omitempty"`
	DisabilityStatusInformationData []DisabilityStatusInformationDataforRelatedPersonType `xml:"urn:com.workday/bsvc Disability_Status_Information_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DisabilityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DisabilityObjectType struct {
	ID         []DisabilityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Reference element representing a unique instance of Disability.  Does not support the update of Disability Status.  Use Put Applicant web service to update the Disability Status for a Person.
type DisabilityReferenceType struct {
	DisabilityName                        string                      `xml:"urn:com.workday/bsvc Disability_Name"`
	RegulatoryRegionOrganizationReference []OrganizationReferenceType `xml:"urn:com.workday/bsvc Regulatory_Region_Organization_Reference,omitempty"`
}

// Wrapper element for each disability status entry.
type DisabilityStatusInformationDataforRelatedPersonType struct {
	DisabilityStatusReference *DisabilityStatusReferenceObjectType         `xml:"urn:com.workday/bsvc Disability_Status_Reference,omitempty"`
	Delete                    *bool                                        `xml:"urn:com.workday/bsvc Delete,omitempty"`
	DisabilityStatusData      *DisabilityStatusSubDataforRelatedPersonType `xml:"urn:com.workday/bsvc Disability_Status_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DisabilityStatusReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DisabilityStatusReferenceObjectType struct {
	ID         []DisabilityStatusReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Container for disability status data for the Personal Information Change business process.
type DisabilityStatusSubDataforRelatedPersonType struct {
	DisabilityReference      *DisabilityObjectType      `xml:"urn:com.workday/bsvc Disability_Reference"`
	DisabilityStatusDate     *time.Time                 `xml:"urn:com.workday/bsvc Disability_Status_Date,omitempty"`
	DateKnown                *time.Time                 `xml:"urn:com.workday/bsvc Date_Known,omitempty"`
	DisabilityEndDate        *time.Time                 `xml:"urn:com.workday/bsvc Disability_End_Date,omitempty"`
	DisabilityGradeReference *DisabilityGradeObjectType `xml:"urn:com.workday/bsvc Disability_Grade_Reference,omitempty"`
}

func (t *DisabilityStatusSubDataforRelatedPersonType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DisabilityStatusSubDataforRelatedPersonType
	var layout struct {
		*T
		DisabilityStatusDate *xsdDate `xml:"urn:com.workday/bsvc Disability_Status_Date,omitempty"`
		DateKnown            *xsdDate `xml:"urn:com.workday/bsvc Date_Known,omitempty"`
		DisabilityEndDate    *xsdDate `xml:"urn:com.workday/bsvc Disability_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DisabilityStatusDate = (*xsdDate)(layout.T.DisabilityStatusDate)
	layout.DateKnown = (*xsdDate)(layout.T.DateKnown)
	layout.DisabilityEndDate = (*xsdDate)(layout.T.DisabilityEndDate)
	return e.EncodeElement(layout, start)
}
func (t *DisabilityStatusSubDataforRelatedPersonType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DisabilityStatusSubDataforRelatedPersonType
	var overlay struct {
		*T
		DisabilityStatusDate *xsdDate `xml:"urn:com.workday/bsvc Disability_Status_Date,omitempty"`
		DateKnown            *xsdDate `xml:"urn:com.workday/bsvc Date_Known,omitempty"`
		DisabilityEndDate    *xsdDate `xml:"urn:com.workday/bsvc Disability_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DisabilityStatusDate = (*xsdDate)(overlay.T.DisabilityStatusDate)
	overlay.DateKnown = (*xsdDate)(overlay.T.DateKnown)
	overlay.DisabilityEndDate = (*xsdDate)(overlay.T.DisabilityEndDate)
	return d.DecodeElement(&overlay, &start)
}

// ESRD Information
type ESRDDataType struct {
	CoordinationPeriodStartDate time.Time  `xml:"urn:com.workday/bsvc Coordination_Period_Start_Date"`
	CoordinationPeriodEndDate   *time.Time `xml:"urn:com.workday/bsvc Coordination_Period_End_Date,omitempty"`
	ESRDSelfTrainingDate        *time.Time `xml:"urn:com.workday/bsvc ESRD_Self_Training_Date,omitempty"`
	FirstDialysisDate           *time.Time `xml:"urn:com.workday/bsvc First_Dialysis_Date,omitempty"`
	TransplantDate              *time.Time `xml:"urn:com.workday/bsvc Transplant_Date,omitempty"`
	TransplantFailureDate       *time.Time `xml:"urn:com.workday/bsvc Transplant_Failure_Date,omitempty"`
}

func (t *ESRDDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ESRDDataType
	var layout struct {
		*T
		CoordinationPeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Coordination_Period_Start_Date"`
		CoordinationPeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Coordination_Period_End_Date,omitempty"`
		ESRDSelfTrainingDate        *xsdDate `xml:"urn:com.workday/bsvc ESRD_Self_Training_Date,omitempty"`
		FirstDialysisDate           *xsdDate `xml:"urn:com.workday/bsvc First_Dialysis_Date,omitempty"`
		TransplantDate              *xsdDate `xml:"urn:com.workday/bsvc Transplant_Date,omitempty"`
		TransplantFailureDate       *xsdDate `xml:"urn:com.workday/bsvc Transplant_Failure_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CoordinationPeriodStartDate = (*xsdDate)(&layout.T.CoordinationPeriodStartDate)
	layout.CoordinationPeriodEndDate = (*xsdDate)(layout.T.CoordinationPeriodEndDate)
	layout.ESRDSelfTrainingDate = (*xsdDate)(layout.T.ESRDSelfTrainingDate)
	layout.FirstDialysisDate = (*xsdDate)(layout.T.FirstDialysisDate)
	layout.TransplantDate = (*xsdDate)(layout.T.TransplantDate)
	layout.TransplantFailureDate = (*xsdDate)(layout.T.TransplantFailureDate)
	return e.EncodeElement(layout, start)
}
func (t *ESRDDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ESRDDataType
	var overlay struct {
		*T
		CoordinationPeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Coordination_Period_Start_Date"`
		CoordinationPeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Coordination_Period_End_Date,omitempty"`
		ESRDSelfTrainingDate        *xsdDate `xml:"urn:com.workday/bsvc ESRD_Self_Training_Date,omitempty"`
		FirstDialysisDate           *xsdDate `xml:"urn:com.workday/bsvc First_Dialysis_Date,omitempty"`
		TransplantDate              *xsdDate `xml:"urn:com.workday/bsvc Transplant_Date,omitempty"`
		TransplantFailureDate       *xsdDate `xml:"urn:com.workday/bsvc Transplant_Failure_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CoordinationPeriodStartDate = (*xsdDate)(&overlay.T.CoordinationPeriodStartDate)
	overlay.CoordinationPeriodEndDate = (*xsdDate)(overlay.T.CoordinationPeriodEndDate)
	overlay.ESRDSelfTrainingDate = (*xsdDate)(overlay.T.ESRDSelfTrainingDate)
	overlay.FirstDialysisDate = (*xsdDate)(overlay.T.FirstDialysisDate)
	overlay.TransplantDate = (*xsdDate)(overlay.T.TransplantDate)
	overlay.TransplantFailureDate = (*xsdDate)(overlay.T.TransplantFailureDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for the Edit Dependent web service operation.
type EditDependentDataType struct {
	DependentReference     *DependentObjectType        `xml:"urn:com.workday/bsvc Dependent_Reference"`
	EditDependentEventData *EditDependentEventDataType `xml:"urn:com.workday/bsvc Edit_Dependent_Event_Data,omitempty"`
}

// Wrapper element for the Dependent's data.
type EditDependentEventDataType struct {
	DependentID                                               string                                               `xml:"urn:com.workday/bsvc Dependent_ID,omitempty"`
	EffectiveDate                                             *time.Time                                           `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	ReasonReference                                           *BenefitsEventSubcategoryObjectType                  `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
	RelatedPersonRelationshipReference                        *RelatedPersonRelationshipObjectType                 `xml:"urn:com.workday/bsvc Related_Person_Relationship_Reference,omitempty"`
	UseEmployeeAddress                                        *bool                                                `xml:"urn:com.workday/bsvc Use_Employee_Address,omitempty"`
	UseEmployeePhone                                          *bool                                                `xml:"urn:com.workday/bsvc Use_Employee_Phone,omitempty"`
	DependentPersonalInformationData                          *RelatedPersonPersonalInformationDataType            `xml:"urn:com.workday/bsvc Dependent_Personal_Information_Data,omitempty"`
	QualifiedDomesticRelationsOrderReplacementData            []QualifiedDomesticRelationsOrderReplacementDataType `xml:"urn:com.workday/bsvc Qualified_Domestic_Relations_Order_Replacement_Data,omitempty"`
	CouldBeCoveredForHealthCareCoverageElsewhere              *bool                                                `xml:"urn:com.workday/bsvc Could_Be_Covered_For_Health_Care_Coverage_Elsewhere,omitempty"`
	CouldBeCoveredForHealthCareCoverageElsewhereEffectiveDate *time.Time                                           `xml:"urn:com.workday/bsvc Could_Be_Covered_For_Health_Care_Coverage_Elsewhere_Effective_Date,omitempty"`
}

func (t *EditDependentEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EditDependentEventDataType
	var layout struct {
		*T
		EffectiveDate                                             *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		CouldBeCoveredForHealthCareCoverageElsewhereEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Could_Be_Covered_For_Health_Care_Coverage_Elsewhere_Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	layout.CouldBeCoveredForHealthCareCoverageElsewhereEffectiveDate = (*xsdDate)(layout.T.CouldBeCoveredForHealthCareCoverageElsewhereEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EditDependentEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EditDependentEventDataType
	var overlay struct {
		*T
		EffectiveDate                                             *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		CouldBeCoveredForHealthCareCoverageElsewhereEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Could_Be_Covered_For_Health_Care_Coverage_Elsewhere_Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	overlay.CouldBeCoveredForHealthCareCoverageElsewhereEffectiveDate = (*xsdDate)(overlay.T.CouldBeCoveredForHealthCareCoverageElsewhereEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for the Edit Dependent web service operation.
type EditDependentRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EditDependentData         *EditDependentDataType         `xml:"urn:com.workday/bsvc Edit_Dependent_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the response to an Edit Dependent web service operation request.
type EditDependentResponseType struct {
	EventReference        *UniqueIdentifierObjectType                    `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	EmployeeReference     *WorkerObjectType                              `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	DependentReference    *DependentObjectType                           `xml:"urn:com.workday/bsvc Dependent_Reference,omitempty"`
	ExceptionResponseData []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exception_Response_Data,omitempty"`
	Version               string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Retirement Savings Elections to be added, updated or deleted.
type EmployeeDefinedContributionElectionDataType struct {
	EmployeeReference               *EmployeeObjectType                   `xml:"urn:com.workday/bsvc Employee_Reference"`
	TakesEffectOnDate               time.Time                             `xml:"urn:com.workday/bsvc Takes_Effect_On_Date"`
	DefinedContributionElectionData []DefinedContributionElectionDataType `xml:"urn:com.workday/bsvc Defined_Contribution_Election_Data,omitempty"`
}

func (t *EmployeeDefinedContributionElectionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeDefinedContributionElectionDataType
	var layout struct {
		*T
		TakesEffectOnDate *xsdDate `xml:"urn:com.workday/bsvc Takes_Effect_On_Date"`
	}
	layout.T = (*T)(t)
	layout.TakesEffectOnDate = (*xsdDate)(&layout.T.TakesEffectOnDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeDefinedContributionElectionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeDefinedContributionElectionDataType
	var overlay struct {
		*T
		TakesEffectOnDate *xsdDate `xml:"urn:com.workday/bsvc Takes_Effect_On_Date"`
	}
	overlay.T = (*T)(t)
	overlay.TakesEffectOnDate = (*xsdDate)(&overlay.T.TakesEffectOnDate)
	return d.DecodeElement(&overlay, &start)
}

// Employee Request References
type EmployeeDefinedContributionElectionsRequestCriteriaType struct {
	EmployeeReference []EmployeeObjectType `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
}

// Retirement Savings Elections Response Data
type EmployeeDefinedContributionElectionsResponseDataType struct {
	EmployeeDefinedContributionElection []DefinedContributionPackageType `xml:"urn:com.workday/bsvc Employee_Defined_Contribution_Election,omitempty"`
}

// ACA 1095-C Form Part I - Employee
type EmployeeInformationDataType struct {
	EmployeeFirstName                   string `xml:"urn:com.workday/bsvc Employee_First_Name"`
	EmployeeMiddleName                  string `xml:"urn:com.workday/bsvc Employee_Middle_Name,omitempty"`
	EmployeeLastName                    string `xml:"urn:com.workday/bsvc Employee_Last_Name"`
	EmployeeSuffixName                  string `xml:"urn:com.workday/bsvc Employee_Suffix_Name,omitempty"`
	SocialSecurityNumber                string `xml:"urn:com.workday/bsvc Social_Security_Number"`
	StreetAddress                       string `xml:"urn:com.workday/bsvc Street_Address"`
	CityorTown                          string `xml:"urn:com.workday/bsvc City_or_Town"`
	TwoLetterStateCodeorForeignProvince string `xml:"urn:com.workday/bsvc Two_Letter_State_Code_or_Foreign_Province"`
	TwoLetterISOCountryCode             string `xml:"urn:com.workday/bsvc Two_Letter_ISO_Country_Code"`
	ZIPCodeorForeignPostalCode          string `xml:"urn:com.workday/bsvc ZIP_Code_or_Foreign_Postal_Code"`
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

// Enroll in Retirement Savings Plans Data
type EnrollinRetirementSavingsPlansDataType struct {
	EmployeeReference                 *WorkerObjectType                                                  `xml:"urn:com.workday/bsvc Employee_Reference"`
	EventDate                         time.Time                                                          `xml:"urn:com.workday/bsvc Event_Date"`
	RetirementSavingsElectionData     []RetirementSavingsElectionDataforChangeRetirementSavingsPlansType `xml:"urn:com.workday/bsvc Retirement_Savings_Election_Data,omitempty"`
	RetirementSavingsCoverageTypeData []BenefitCoverageTypeforRetirementSavingsElectionsDataType         `xml:"urn:com.workday/bsvc Retirement_Savings_Coverage_Type_Data,omitempty"`
}

func (t *EnrollinRetirementSavingsPlansDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EnrollinRetirementSavingsPlansDataType
	var layout struct {
		*T
		EventDate *xsdDate `xml:"urn:com.workday/bsvc Event_Date"`
	}
	layout.T = (*T)(t)
	layout.EventDate = (*xsdDate)(&layout.T.EventDate)
	return e.EncodeElement(layout, start)
}
func (t *EnrollinRetirementSavingsPlansDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EnrollinRetirementSavingsPlansDataType
	var overlay struct {
		*T
		EventDate *xsdDate `xml:"urn:com.workday/bsvc Event_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EventDate = (*xsdDate)(&overlay.T.EventDate)
	return d.DecodeElement(&overlay, &start)
}

// Enroll in Retirement Savings Plans Request
type EnrollinRetirementSavingsPlansRequestType struct {
	BusinessProcessParameters          *BusinessProcessParametersType          `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EnrollinRetirementSavingsPlansData *EnrollinRetirementSavingsPlansDataType `xml:"urn:com.workday/bsvc Enroll_in_Retirement_Savings_Plans_Data,omitempty"`
	Version                            string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Enroll in Retirement Savings Plans Response
type EnrollinRetirementSavingsPlansResponseType struct {
	EnrollinRetirementSavingsPlansReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Enroll_in_Retirement_Savings_Plans_Reference,omitempty"`
	Version                                 string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Update Evidence of Insurability
type EvidenceofInsurabilityDataType struct {
	WorkerReference      []WorkerObjectType      `xml:"urn:com.workday/bsvc Worker_Reference"`
	BenefitPlanReference []BenefitPlanObjectType `xml:"urn:com.workday/bsvc Benefit_Plan_Reference"`
	Approveforselected   bool                    `xml:"urn:com.workday/bsvc Approve_for_selected"`
	Denyforselected      bool                    `xml:"urn:com.workday/bsvc Deny_for_selected"`
	EOIApproveOrDenyDate time.Time               `xml:"urn:com.workday/bsvc EOI_Approve_Or_Deny_Date"`
}

func (t *EvidenceofInsurabilityDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EvidenceofInsurabilityDataType
	var layout struct {
		*T
		EOIApproveOrDenyDate *xsdDate `xml:"urn:com.workday/bsvc EOI_Approve_Or_Deny_Date"`
	}
	layout.T = (*T)(t)
	layout.EOIApproveOrDenyDate = (*xsdDate)(&layout.T.EOIApproveOrDenyDate)
	return e.EncodeElement(layout, start)
}
func (t *EvidenceofInsurabilityDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EvidenceofInsurabilityDataType
	var overlay struct {
		*T
		EOIApproveOrDenyDate *xsdDate `xml:"urn:com.workday/bsvc EOI_Approve_Or_Deny_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EOIApproveOrDenyDate = (*xsdDate)(&overlay.T.EOIApproveOrDenyDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type EvidenceofInsurabilityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EvidenceofInsurabilityObjectType struct {
	ID         []EvidenceofInsurabilityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Exception (Errors and Warning) associated with the transaction.
type ExceptionDataType struct {
	Classification string `xml:"urn:com.workday/bsvc Classification,omitempty"`
	Message        string `xml:"urn:com.workday/bsvc Message,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type External1095CRecipientObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type External1095CRecipientObjectType struct {
	ID         []External1095CRecipientObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Integration ID reference is used as a unique identifier for integratable objects in the Workday system.
type ExternalIntegrationIDReferenceDataType struct {
	ID         *IDType `xml:"urn:com.workday/bsvc ID"`
	Descriptor string  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Flat Amount Data
type FlatAmountDataType struct {
	EmployeeCostAmount        float64 `xml:"urn:com.workday/bsvc Employee_Cost_Amount,omitempty"`
	EmployeeCostAmountPostTax float64 `xml:"urn:com.workday/bsvc Employee_Cost_Amount_PostTax,omitempty"`
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
type GenderObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type GenderObjectType struct {
	ID         []GenderObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Reference element representing a unique instance of Gender.
type GenderReferenceType struct {
	GenderDescription string `xml:"urn:com.workday/bsvc Gender_Description"`
}

// Contains a unique identifier for an instance of an object.
type GeneralEventSubcategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type GeneralEventSubcategoryObjectType struct {
	ID         []GeneralEventSubcategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Get ACA 1095-C Forms Data Web Service Request
type GetACA1095CFormsDataRequestType struct {
	RequestCriteria *ACA1095CFormRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version         string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get ACA 1095-C Forms Data Response
type GetACA1095CFormsDataResponseType struct {
	RequestCriteria *ACA1095CFormRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults *ResponseResultsType             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *ACA1095CFormResponseDataType    `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for Get Benefit Annual Credits Request.
type GetBenefitAnnualCreditsRequestType struct {
	RequestCriteria []BenefitAnnualCreditRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  []ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version         string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for Get Benefit Annual Credits Response.
type GetBenefitAnnualCreditsResponseType struct {
	RequestReferences []BenefitAnnualCreditRequestCriteriaType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    []ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   []ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []BenefitAnnualCreditResponseDataType    `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for Get Benefit Annual Rates Request.
type GetBenefitAnnualRatesRequestType struct {
	RequestCriteria []BenefitAnnualRateRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  []ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version         string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for Get Benefit Annual Rates Response.
type GetBenefitAnnualRatesResponseType struct {
	RequestReferences []BenefitAnnualRateRequestCriteriaType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    []ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   []ResponseResultsType                  `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []BenefitAnnualRateResponseDataType    `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This service operation retrieves Benefits Individual Rates for one or more employees.
type GetBenefitIndividualRatesRequestType struct {
	RequestReferences *BenefitIndividualRateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BenefitIndividualRateRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response for a Get Benefit Individual Rate web service request
type GetBenefitIndividualRatesResponseType struct {
	RequestReferences *BenefitIndividualRateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BenefitIndividualRateRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                        `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BenefitIndividualRateResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request for Retirement Savings Elections
type GetEmployeeDefinedContributionElectionsRequestType struct {
	RequestCriteria *EmployeeDefinedContributionElectionsRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version         string                                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response Data for Retirement Savings Elections
type GetEmployeeDefinedContributionElectionsResponseType struct {
	RequestCriteria *EmployeeDefinedContributionElectionsRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults *ResponseResultsType                                     `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *EmployeeDefinedContributionElectionsResponseDataType    `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Health Care Rates Web Service Request
type GetHealthCareRatesRequestType struct {
	RequestReferences *HealthCareRateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *HealthCareRateRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *HealthCareRateResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Health Care Rates Web Service Response
type GetHealthCareRatesResponseType struct {
	RequestReferences *HealthCareRateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *HealthCareRateRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *HealthCareRateResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *HealthCareRateResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get All Medicare Part D EGWP Data
type GetMedicarePartDEGWPRequestType struct {
	ResponseFilter *ResponseFilterType `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version        string              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Container for All Medicare Part D EGWP Data
type GetMedicarePartDEGWPResponseType struct {
	ResponseFilter  *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults *ResponseResultsType               `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *MedicarePartDEGWPResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper for the Government Identifier Data.
type GovernmentIDDataType struct {
	ID                  string                         `xml:"urn:com.workday/bsvc ID,omitempty"`
	IDTypeReference     *GovernmentIDTypeAllObjectType `xml:"urn:com.workday/bsvc ID_Type_Reference,omitempty"`
	CountryReference    *CountryObjectType             `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	IssuedDate          *time.Time                     `xml:"urn:com.workday/bsvc Issued_Date,omitempty"`
	ExpirationDate      *time.Time                     `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	VerificationDate    *time.Time                     `xml:"urn:com.workday/bsvc Verification_Date,omitempty"`
	VerifiedbyReference *WorkerObjectType              `xml:"urn:com.workday/bsvc Verified_by_Reference,omitempty"`
}

func (t *GovernmentIDDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T GovernmentIDDataType
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
func (t *GovernmentIDDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T GovernmentIDDataType
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

// Encapsulating element for all Government Identifier data.
type GovernmentIDType struct {
	GovernmentIDReference       *UniqueIdentifierObjectType              `xml:"urn:com.workday/bsvc Government_ID_Reference,omitempty"`
	GovernmentIDData            *GovernmentIDDataType                    `xml:"urn:com.workday/bsvc Government_ID_Data,omitempty"`
	GovernmentIDSharedReference *GovernmentIdentifierReferenceObjectType `xml:"urn:com.workday/bsvc Government_ID_Shared_Reference,omitempty"`
	Delete                      bool                                     `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type GovernmentIDTypeAllObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type GovernmentIDTypeAllObjectType struct {
	ID         []GovernmentIDTypeAllObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type GovernmentIdentifierReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type GovernmentIdentifierReferenceObjectType struct {
	ID         []GovernmentIdentifierReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Top level request to modify COBRA Eligibility Data
type GrantCOBRAEligibilityRequestType struct {
	BusinessProcessParameters  *BusinessProcessParametersType              `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	COBRAEligibilityStatusData []COBRAParticipantEligibilityStatusDataType `xml:"urn:com.workday/bsvc COBRA_Eligibility_Status_Data,omitempty"`
	AddOnly                    bool                                        `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                    string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// A container for A reference to the newly created or modified COBRA Eligiblity
type GrantCOBRAEligibilityResponseType struct {
	COBRAParticipantEligibilityStatusReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc COBRA_Participant_Eligibility_Status_Reference,omitempty"`
	Version                                    string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// If the dependent has health insurance
type HasHealthInsuranceDataType struct {
	EffectiveDate      *time.Time `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	HasHealthInsurance *bool      `xml:"urn:com.workday/bsvc Has_Health_Insurance,omitempty"`
}

func (t *HasHealthInsuranceDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T HasHealthInsuranceDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *HasHealthInsuranceDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T HasHealthInsuranceDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type HealthCareBandedRateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type HealthCareBandedRateObjectType struct {
	ID         []HealthCareBandedRateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type HealthCareCoveragePlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type HealthCareCoveragePlanObjectType struct {
	ID         []HealthCareCoveragePlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type HealthCareCoverageTargetObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type HealthCareCoverageTargetObjectType struct {
	ID         []HealthCareCoverageTargetObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Health Care Election data for the Change Benefits web services
type HealthCareElectionChangeBenefitsDataType struct {
	HealthCareCoveragePlanReference   *HealthCareCoveragePlanObjectType                           `xml:"urn:com.workday/bsvc Health_Care_Coverage_Plan_Reference"`
	HealthCareCoverageTargetReference *HealthCareCoverageTargetObjectType                         `xml:"urn:com.workday/bsvc Health_Care_Coverage_Target_Reference,omitempty"`
	ProviderID                        string                                                      `xml:"urn:com.workday/bsvc Provider_ID,omitempty"`
	DependentElectionData             []DependentElectionDataType                                 `xml:"urn:com.workday/bsvc Dependent_Election_Data,omitempty"`
	BenefitIndividualRateData         *BenefitIndividualRateDataforChangeBenefitsforLifeEventType `xml:"urn:com.workday/bsvc Benefit_Individual_Rate_Data,omitempty"`
}

// Contains Health Care Rate Data
type HealthCareRateDataType struct {
	HealthCareRateID                  string                                `xml:"urn:com.workday/bsvc Health_Care_Rate_ID,omitempty"`
	Name                              string                                `xml:"urn:com.workday/bsvc Name"`
	Description                       string                                `xml:"urn:com.workday/bsvc Description,omitempty"`
	UseMinimumCost                    *bool                                 `xml:"urn:com.workday/bsvc Use_Minimum_Cost,omitempty"`
	UseMaximumCost                    *bool                                 `xml:"urn:com.workday/bsvc Use_Maximum_Cost,omitempty"`
	FrequencyReference                *FrequencyObjectType                  `xml:"urn:com.workday/bsvc Frequency_Reference"`
	CurrencyReference                 *CurrencyObjectType                   `xml:"urn:com.workday/bsvc Currency_Reference"`
	HealthCareCoverageTargetReference []HealthCareCoverageTargetObjectType  `xml:"urn:com.workday/bsvc Health_Care_Coverage_Target_Reference,omitempty"`
	AgeinYearsData                    *AgeinYearsDataType                   `xml:"urn:com.workday/bsvc Age_in_Years_Data,omitempty"`
	LengthofServiceData               *LengthofServiceDataType              `xml:"urn:com.workday/bsvc Length_of_Service_Data,omitempty"`
	SalaryData                        *SalaryDataType                       `xml:"urn:com.workday/bsvc Salary_Data,omitempty"`
	HealthCareRateFlatData            []HealthCareRateFlatDataType          `xml:"urn:com.workday/bsvc Health_Care_Rate_Flat_Data,omitempty"`
	HealthCareRatePercentSalaryData   []HealthCareRatePercentSalaryDataType `xml:"urn:com.workday/bsvc Health_Care_Rate_Percent_Salary_Data,omitempty"`
	ConsiderTobacco                   *bool                                 `xml:"urn:com.workday/bsvc Consider_Tobacco,omitempty"`
	EffectiveDate                     time.Time                             `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
}

func (t *HealthCareRateDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T HealthCareRateDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *HealthCareRateDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T HealthCareRateDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Flat Health Care Rate
type HealthCareRateFlatDataType struct {
	BenefitRateID                     string                               `xml:"urn:com.workday/bsvc Benefit_Rate_ID,omitempty"`
	Minimum                           float64                              `xml:"urn:com.workday/bsvc Minimum,omitempty"`
	Maximum                           float64                              `xml:"urn:com.workday/bsvc Maximum,omitempty"`
	SalaryMinimum                     float64                              `xml:"urn:com.workday/bsvc Salary_Minimum,omitempty"`
	SalaryMaximum                     float64                              `xml:"urn:com.workday/bsvc Salary_Maximum,omitempty"`
	HealthCareCoverageTargetReference *HealthCareCoverageTargetObjectType  `xml:"urn:com.workday/bsvc Health_Care_Coverage_Target_Reference"`
	RateValueFlatData                 []BenefitRateValueFlatDataType       `xml:"urn:com.workday/bsvc Rate_Value_Flat_Data,omitempty"`
	TobaccoUseReference               *BenefitRateBandTobaccoUseObjectType `xml:"urn:com.workday/bsvc Tobacco_Use_Reference,omitempty"`
}

// Contains Health Care Rate Percent Salary Data
type HealthCareRatePercentSalaryDataType struct {
	BenefitRateSalarySourceData       []HealthCareRateSalarySourceDataType    `xml:"urn:com.workday/bsvc Benefit_Rate_Salary_Source_Data"`
	BenefitRoundingRuleFinalData      []BenefitRoundingRuleFinalDataType      `xml:"urn:com.workday/bsvc Benefit_Rounding_Rule_Final_Data,omitempty"`
	HealthCareRatePercentofSalaryData []HealthCareRatePercentofSalaryDataType `xml:"urn:com.workday/bsvc Health_Care_Rate_Percent_of_Salary_Data"`
}

// Contains Health Care Rate Percent of Salary Data
type HealthCareRatePercentofSalaryDataType struct {
	BenefitRateID                     string                                    `xml:"urn:com.workday/bsvc Benefit_Rate_ID,omitempty"`
	Minimum                           float64                                   `xml:"urn:com.workday/bsvc Minimum,omitempty"`
	Maximum                           float64                                   `xml:"urn:com.workday/bsvc Maximum,omitempty"`
	SalaryMinimum                     float64                                   `xml:"urn:com.workday/bsvc Salary_Minimum,omitempty"`
	SalaryMaximum                     float64                                   `xml:"urn:com.workday/bsvc Salary_Maximum,omitempty"`
	HealthCareCoverageTargetReference *HealthCareCoverageTargetObjectType       `xml:"urn:com.workday/bsvc Health_Care_Coverage_Target_Reference,omitempty"`
	RateValuePercentofSalaryData      []BenefitRateValuePercentofSalaryDataType `xml:"urn:com.workday/bsvc Rate_Value_Percent_of_Salary_Data"`
	TobaccoUseReference               *BenefitRateBandTobaccoUseObjectType      `xml:"urn:com.workday/bsvc Tobacco_Use_Reference,omitempty"`
}

// Contains Health Care Rate Request Criteria
type HealthCareRateRequestCriteriaType struct {
}

// Contains Health Care Rate Request References
type HealthCareRateRequestReferencesType struct {
	HealthCareRateReference []HealthCareBandedRateObjectType `xml:"urn:com.workday/bsvc Health_Care_Rate_Reference"`
}

// Contains Health Care Rate Response Data
type HealthCareRateResponseDataType struct {
	HealthCareRate []HealthCareRateType `xml:"urn:com.workday/bsvc Health_Care_Rate,omitempty"`
}

// Contains Health Care Rate Response Group
type HealthCareRateResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Contains Health Care Rate Salary Source Data
type HealthCareRateSalarySourceDataType struct {
	SalaryData []SalaryDataType `xml:"urn:com.workday/bsvc Salary_Data,omitempty"`
}

// Contains Health Care Rate
type HealthCareRateType struct {
	HealthCareRateReference *HealthCareBandedRateObjectType `xml:"urn:com.workday/bsvc Health_Care_Rate_Reference,omitempty"`
	HealthCareRateData      []HealthCareRateDataType        `xml:"urn:com.workday/bsvc Health_Care_Rate_Data,omitempty"`
}

// Election data for a health savings account
type HealthSavingsAccountElectionDataType struct {
	HealthSavingsAccountPlanReference *HealthSavingsAccountPlanObjectType              `xml:"urn:com.workday/bsvc Health_Savings_Account_Plan_Reference"`
	YTDContributionAmount             float64                                          `xml:"urn:com.workday/bsvc YTD_Contribution_Amount,omitempty"`
	AnnualContribution                float64                                          `xml:"urn:com.workday/bsvc Annual_Contribution,omitempty"`
	EmployeeCost                      float64                                          `xml:"urn:com.workday/bsvc Employee_Cost,omitempty"`
	BeneficiaryAllocationData         []BeneficiaryAllocationforChangeBenefitsDataType `xml:"urn:com.workday/bsvc Beneficiary_Allocation_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type HealthSavingsAccountPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type HealthSavingsAccountPlanObjectType struct {
	ID         []HealthSavingsAccountPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// External ID that uniquely identifies the integratable object within the context of the integration system identified by the System ID attribute.
type IDType struct {
	Value    string `xml:",chardata"`
	SystemID string `xml:"urn:com.workday/bsvc System_ID,attr,omitempty"`
}

// ACA 1095-C Forms Data Request
type ImportACA1095CFormsDataRequestType struct {
	ACA1095CFormsData *ACA1095CFormsType `xml:"urn:com.workday/bsvc ACA_1095-C_Forms_Data,omitempty"`
	Version           string             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Import Medicare Part D EGWP Data
type ImportMedicarePartDEGWPRequestType struct {
	MedicarePartDEGWPData []MedicarePartDEGWPDataType `xml:"urn:com.workday/bsvc Medicare_Part_D_EGWP_Data,omitempty"`
	Version               string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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
type InsuranceCoverageMasterAmountObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InsuranceCoverageMasterAmountObjectType struct {
	ID         []InsuranceCoverageMasterAmountObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type InsuranceCoveragePlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InsuranceCoveragePlanObjectType struct {
	ID         []InsuranceCoveragePlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Insurance election data
type InsuranceElectionDataChangeBenefitsType struct {
	InsuranceCoveragePlanReference *InsuranceCoveragePlanObjectType                            `xml:"urn:com.workday/bsvc Insurance_Coverage_Plan_Reference"`
	DependentReference             []DependentObjectType                                       `xml:"urn:com.workday/bsvc Dependent_Reference,omitempty"`
	CoverageAmountReference        *InsuranceCoverageMasterAmountObjectType                    `xml:"urn:com.workday/bsvc Coverage_Amount_Reference"`
	BeneficiaryAllocationData      []BeneficiaryAllocationforChangeBenefitsDataType            `xml:"urn:com.workday/bsvc Beneficiary_Allocation_Data,omitempty"`
	BenefitIndividualRateData      *BenefitIndividualRateDataforChangeBenefitsforLifeEventType `xml:"urn:com.workday/bsvc Benefit_Individual_Rate_Data,omitempty"`
}

// Contains the legal name for a person.  A person must name one and only one legal name.
type LegalNameDataType struct {
	NameDetailData *PersonNameDetailDataType `xml:"urn:com.workday/bsvc Name_Detail_Data"`
}

// Service Length of Health Care Rate
type LengthofServiceDataType struct {
	AsOfDateData *AsOfDateDataType `xml:"urn:com.workday/bsvc As_Of_Date_Data"`
}

// Wrapper for License Identifier Data.
type LicenseIDDataType struct {
	ID                      string                      `xml:"urn:com.workday/bsvc ID,omitempty"`
	IDTypeReference         *LicenseIDTypeAllObjectType `xml:"urn:com.workday/bsvc ID_Type_Reference,omitempty"`
	CountryReference        *CountryObjectType          `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CountryRegionReference  *CountryRegionObjectType    `xml:"urn:com.workday/bsvc Country_Region_Reference,omitempty"`
	CountryRegionDescriptor string                      `xml:"urn:com.workday/bsvc Country_Region_Descriptor,omitempty"`
	AuthorityReference      *AuthorityObjectType        `xml:"urn:com.workday/bsvc Authority_Reference,omitempty"`
	LicenseClass            string                      `xml:"urn:com.workday/bsvc License_Class,omitempty"`
	IssuedDate              *time.Time                  `xml:"urn:com.workday/bsvc Issued_Date,omitempty"`
	ExpirationDate          *time.Time                  `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	VerificationDate        *time.Time                  `xml:"urn:com.workday/bsvc Verification_Date,omitempty"`
	VerifiedByReference     *WorkerObjectType           `xml:"urn:com.workday/bsvc Verified_By_Reference,omitempty"`
}

func (t *LicenseIDDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LicenseIDDataType
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
func (t *LicenseIDDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LicenseIDDataType
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

// Encapsulating element for all License Identifier data.
type LicenseIDType struct {
	LicenseIDReference       *LicenseIdentifierObjectType          `xml:"urn:com.workday/bsvc License_ID_Reference,omitempty"`
	LicenseIDData            *LicenseIDDataType                    `xml:"urn:com.workday/bsvc License_ID_Data,omitempty"`
	LicenseIDSharedReference *LicenseIdentifierReferenceObjectType `xml:"urn:com.workday/bsvc License_ID_Shared_Reference,omitempty"`
	Delete                   bool                                  `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LicenseIDTypeAllObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LicenseIDTypeAllObjectType struct {
	ID         []LicenseIDTypeAllObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LicenseIdentifierObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LicenseIdentifierObjectType struct {
	ID         []LicenseIdentifierObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LicenseIdentifierReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LicenseIdentifierReferenceObjectType struct {
	ID         []LicenseIdentifierReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// ACA 1095-C Form Line 14 - Offer Of Coverage
type Line14MonthDataType struct {
	All12Months string `xml:"urn:com.workday/bsvc All_12_Months,omitempty"`
	Jan         string `xml:"urn:com.workday/bsvc Jan,omitempty"`
	Feb         string `xml:"urn:com.workday/bsvc Feb,omitempty"`
	Mar         string `xml:"urn:com.workday/bsvc Mar,omitempty"`
	Apr         string `xml:"urn:com.workday/bsvc Apr,omitempty"`
	May         string `xml:"urn:com.workday/bsvc May,omitempty"`
	Jun         string `xml:"urn:com.workday/bsvc Jun,omitempty"`
	Jul         string `xml:"urn:com.workday/bsvc Jul,omitempty"`
	Aug         string `xml:"urn:com.workday/bsvc Aug,omitempty"`
	Sep         string `xml:"urn:com.workday/bsvc Sep,omitempty"`
	Oct         string `xml:"urn:com.workday/bsvc Oct,omitempty"`
	Nov         string `xml:"urn:com.workday/bsvc Nov,omitempty"`
	Dec         string `xml:"urn:com.workday/bsvc Dec,omitempty"`
}

// ACA 1095-C Form Line 15 - Employee Share of Lowest Cost Monthly Premium, for Self-Only Minimum Value Coverage
type Line15MonthDataType struct {
	All12Months string `xml:"urn:com.workday/bsvc All_12_Months,omitempty"`
	Jan         string `xml:"urn:com.workday/bsvc Jan,omitempty"`
	Feb         string `xml:"urn:com.workday/bsvc Feb,omitempty"`
	Mar         string `xml:"urn:com.workday/bsvc Mar,omitempty"`
	Apr         string `xml:"urn:com.workday/bsvc Apr,omitempty"`
	May         string `xml:"urn:com.workday/bsvc May,omitempty"`
	Jun         string `xml:"urn:com.workday/bsvc Jun,omitempty"`
	Jul         string `xml:"urn:com.workday/bsvc Jul,omitempty"`
	Aug         string `xml:"urn:com.workday/bsvc Aug,omitempty"`
	Sep         string `xml:"urn:com.workday/bsvc Sep,omitempty"`
	Oct         string `xml:"urn:com.workday/bsvc Oct,omitempty"`
	Nov         string `xml:"urn:com.workday/bsvc Nov,omitempty"`
	Dec         string `xml:"urn:com.workday/bsvc Dec,omitempty"`
}

// ACA 1095-C Form Line 16 - Applicable Section 4980H Safe Harbor
type Line16MonthDataType struct {
	All12Months string `xml:"urn:com.workday/bsvc All_12_Months,omitempty"`
	Jan         string `xml:"urn:com.workday/bsvc Jan,omitempty"`
	Feb         string `xml:"urn:com.workday/bsvc Feb,omitempty"`
	Mar         string `xml:"urn:com.workday/bsvc Mar,omitempty"`
	Apr         string `xml:"urn:com.workday/bsvc Apr,omitempty"`
	May         string `xml:"urn:com.workday/bsvc May,omitempty"`
	Jun         string `xml:"urn:com.workday/bsvc Jun,omitempty"`
	Jul         string `xml:"urn:com.workday/bsvc Jul,omitempty"`
	Aug         string `xml:"urn:com.workday/bsvc Aug,omitempty"`
	Sep         string `xml:"urn:com.workday/bsvc Sep,omitempty"`
	Oct         string `xml:"urn:com.workday/bsvc Oct,omitempty"`
	Nov         string `xml:"urn:com.workday/bsvc Nov,omitempty"`
	Dec         string `xml:"urn:com.workday/bsvc Dec,omitempty"`
}

// ACA 1095-C Form Part III - Covered Individuals
type Line1734CoveredIndividualsMonthDataType struct {
	CoveredIndividualReference  string     `xml:"urn:com.workday/bsvc Covered_Individual_Reference,omitempty"`
	CoveredIndividualFirstName  string     `xml:"urn:com.workday/bsvc Covered_Individual_First_Name,omitempty"`
	CoveredIndividualMiddleName string     `xml:"urn:com.workday/bsvc Covered_Individual_Middle_Name,omitempty"`
	CoveredIndividualLastName   string     `xml:"urn:com.workday/bsvc Covered_Individual_Last_Name,omitempty"`
	CoveredIndividualSuffixName string     `xml:"urn:com.workday/bsvc Covered_Individual_Suffix_Name,omitempty"`
	CoveredIndividualSSNorITIN  string     `xml:"urn:com.workday/bsvc Covered_Individual_SSN_or_ITIN,omitempty"`
	CoveredIndividualDOB        *time.Time `xml:"urn:com.workday/bsvc Covered_Individual_DOB,omitempty"`
	All12Months                 *bool      `xml:"urn:com.workday/bsvc All_12_Months,omitempty"`
	Jan                         *bool      `xml:"urn:com.workday/bsvc Jan,omitempty"`
	Feb                         *bool      `xml:"urn:com.workday/bsvc Feb,omitempty"`
	Mar                         *bool      `xml:"urn:com.workday/bsvc Mar,omitempty"`
	Apr                         *bool      `xml:"urn:com.workday/bsvc Apr,omitempty"`
	May                         *bool      `xml:"urn:com.workday/bsvc May,omitempty"`
	Jun                         *bool      `xml:"urn:com.workday/bsvc Jun,omitempty"`
	Jul                         *bool      `xml:"urn:com.workday/bsvc Jul,omitempty"`
	Aug                         *bool      `xml:"urn:com.workday/bsvc Aug,omitempty"`
	Sep                         *bool      `xml:"urn:com.workday/bsvc Sep,omitempty"`
	Oct                         *bool      `xml:"urn:com.workday/bsvc Oct,omitempty"`
	Nov                         *bool      `xml:"urn:com.workday/bsvc Nov,omitempty"`
	Dec                         *bool      `xml:"urn:com.workday/bsvc Dec,omitempty"`
}

func (t *Line1734CoveredIndividualsMonthDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Line1734CoveredIndividualsMonthDataType
	var layout struct {
		*T
		CoveredIndividualDOB *xsdDate `xml:"urn:com.workday/bsvc Covered_Individual_DOB,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CoveredIndividualDOB = (*xsdDate)(layout.T.CoveredIndividualDOB)
	return e.EncodeElement(layout, start)
}
func (t *Line1734CoveredIndividualsMonthDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Line1734CoveredIndividualsMonthDataType
	var overlay struct {
		*T
		CoveredIndividualDOB *xsdDate `xml:"urn:com.workday/bsvc Covered_Individual_DOB,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CoveredIndividualDOB = (*xsdDate)(overlay.T.CoveredIndividualDOB)
	return d.DecodeElement(&overlay, &start)
}

// If the dependent lives with worker
type LivesWithWorkerDataType struct {
	EffectiveDate   *time.Time `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	LivesWithWorker *bool      `xml:"urn:com.workday/bsvc Lives_With_Worker,omitempty"`
}

func (t *LivesWithWorkerDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LivesWithWorkerDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *LivesWithWorkerDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LivesWithWorkerDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
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

// Data for a Manage Medicare Information Event
type ManageMedicareInformationDataType struct {
	PersonReference          *RoleObjectType                    `xml:"urn:com.workday/bsvc Person_Reference"`
	EventDate                time.Time                          `xml:"urn:com.workday/bsvc Event_Date"`
	MBI                      string                             `xml:"urn:com.workday/bsvc MBI,omitempty"`
	HICN                     string                             `xml:"urn:com.workday/bsvc HICN,omitempty"`
	ReportingReasonReference *GeneralEventSubcategoryObjectType `xml:"urn:com.workday/bsvc Reporting_Reason_Reference"`
	MedicareReasonReference  *MedicareReasonObjectType          `xml:"urn:com.workday/bsvc Medicare_Reason_Reference"`
	MedicarePartAData        *MedicarePartADataType             `xml:"urn:com.workday/bsvc Medicare_Part_A_Data,omitempty"`
	MedicarePartBData        *MedicarePartBDataType             `xml:"urn:com.workday/bsvc Medicare_Part_B_Data,omitempty"`
	MedicarePartDData        *MedicarePartDDataType             `xml:"urn:com.workday/bsvc Medicare_Part_D_Data,omitempty"`
	ESRDData                 *ESRDDataType                      `xml:"urn:com.workday/bsvc ESRD_Data,omitempty"`
}

func (t *ManageMedicareInformationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ManageMedicareInformationDataType
	var layout struct {
		*T
		EventDate *xsdDate `xml:"urn:com.workday/bsvc Event_Date"`
	}
	layout.T = (*T)(t)
	layout.EventDate = (*xsdDate)(&layout.T.EventDate)
	return e.EncodeElement(layout, start)
}
func (t *ManageMedicareInformationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ManageMedicareInformationDataType
	var overlay struct {
		*T
		EventDate *xsdDate `xml:"urn:com.workday/bsvc Event_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EventDate = (*xsdDate)(&overlay.T.EventDate)
	return d.DecodeElement(&overlay, &start)
}

// Request to create a Manage Medicare Information Event
type ManageMedicareInformationRequestType struct {
	BusinessProcessParameters     *BusinessProcessParametersType     `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	ManageMedicareInformationData *ManageMedicareInformationDataType `xml:"urn:com.workday/bsvc Manage_Medicare_Information_Data"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response generated after submitting a Manage Medicare Information Event
type ManageMedicareInformationResponseType struct {
	MedicareInformationReference *MedicareInformationEventObjectType `xml:"urn:com.workday/bsvc Medicare_Information_Reference,omitempty"`
	Version                      string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MedicareEGWPDisenrollmentReasonObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MedicareEGWPDisenrollmentReasonObjectType struct {
	ID         []MedicareEGWPDisenrollmentReasonObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MedicareEGWPEnrollmentStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MedicareEGWPEnrollmentStatusObjectType struct {
	ID         []MedicareEGWPEnrollmentStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MedicareEGWPReasonRejectCodeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MedicareEGWPReasonRejectCodeObjectType struct {
	ID         []MedicareEGWPReasonRejectCodeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MedicareInformationEventObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MedicareInformationEventObjectType struct {
	ID         []MedicareInformationEventObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Medicare Part A Information
type MedicarePartADataType struct {
	EffectiveDate   time.Time  `xml:"urn:com.workday/bsvc Effective_Date"`
	TerminationDate *time.Time `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
}

func (t *MedicarePartADataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MedicarePartADataType
	var layout struct {
		*T
		EffectiveDate   *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
		TerminationDate *xsdDate `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	layout.TerminationDate = (*xsdDate)(layout.T.TerminationDate)
	return e.EncodeElement(layout, start)
}
func (t *MedicarePartADataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MedicarePartADataType
	var overlay struct {
		*T
		EffectiveDate   *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
		TerminationDate *xsdDate `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	overlay.TerminationDate = (*xsdDate)(overlay.T.TerminationDate)
	return d.DecodeElement(&overlay, &start)
}

// Medicare Part B Information
type MedicarePartBDataType struct {
	EffectiveDate   time.Time  `xml:"urn:com.workday/bsvc Effective_Date"`
	TerminationDate *time.Time `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
}

func (t *MedicarePartBDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MedicarePartBDataType
	var layout struct {
		*T
		EffectiveDate   *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
		TerminationDate *xsdDate `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	layout.TerminationDate = (*xsdDate)(layout.T.TerminationDate)
	return e.EncodeElement(layout, start)
}
func (t *MedicarePartBDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MedicarePartBDataType
	var overlay struct {
		*T
		EffectiveDate   *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
		TerminationDate *xsdDate `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	overlay.TerminationDate = (*xsdDate)(overlay.T.TerminationDate)
	return d.DecodeElement(&overlay, &start)
}

// Medicare Part D Information
type MedicarePartDDataType struct {
	EnrollmentDate       time.Time  `xml:"urn:com.workday/bsvc Enrollment_Date"`
	TerminationDate      *time.Time `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
	EligibilityStartDate time.Time  `xml:"urn:com.workday/bsvc Eligibility_Start_Date"`
	EligibilityStopDate  *time.Time `xml:"urn:com.workday/bsvc Eligibility_Stop_Date,omitempty"`
	PlanContractorNumber string     `xml:"urn:com.workday/bsvc Plan_Contractor_Number,omitempty"`
}

func (t *MedicarePartDDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MedicarePartDDataType
	var layout struct {
		*T
		EnrollmentDate       *xsdDate `xml:"urn:com.workday/bsvc Enrollment_Date"`
		TerminationDate      *xsdDate `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
		EligibilityStartDate *xsdDate `xml:"urn:com.workday/bsvc Eligibility_Start_Date"`
		EligibilityStopDate  *xsdDate `xml:"urn:com.workday/bsvc Eligibility_Stop_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EnrollmentDate = (*xsdDate)(&layout.T.EnrollmentDate)
	layout.TerminationDate = (*xsdDate)(layout.T.TerminationDate)
	layout.EligibilityStartDate = (*xsdDate)(&layout.T.EligibilityStartDate)
	layout.EligibilityStopDate = (*xsdDate)(layout.T.EligibilityStopDate)
	return e.EncodeElement(layout, start)
}
func (t *MedicarePartDDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MedicarePartDDataType
	var overlay struct {
		*T
		EnrollmentDate       *xsdDate `xml:"urn:com.workday/bsvc Enrollment_Date"`
		TerminationDate      *xsdDate `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
		EligibilityStartDate *xsdDate `xml:"urn:com.workday/bsvc Eligibility_Start_Date"`
		EligibilityStopDate  *xsdDate `xml:"urn:com.workday/bsvc Eligibility_Stop_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EnrollmentDate = (*xsdDate)(&overlay.T.EnrollmentDate)
	overlay.TerminationDate = (*xsdDate)(overlay.T.TerminationDate)
	overlay.EligibilityStartDate = (*xsdDate)(&overlay.T.EligibilityStartDate)
	overlay.EligibilityStopDate = (*xsdDate)(overlay.T.EligibilityStopDate)
	return d.DecodeElement(&overlay, &start)
}

// Individual Rows of Medicare Part D EGWP Data
type MedicarePartDEGWPDataType struct {
	EffectiveDate                *time.Time                                 `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	MBI                          string                                     `xml:"urn:com.workday/bsvc MBI,omitempty"`
	HICN                         string                                     `xml:"urn:com.workday/bsvc HICN,omitempty"`
	SubscriberNumber             string                                     `xml:"urn:com.workday/bsvc Subscriber_Number"`
	EnrollmentStatusReference    *MedicareEGWPEnrollmentStatusObjectType    `xml:"urn:com.workday/bsvc Enrollment_Status_Reference"`
	EnrollmentDate               *time.Time                                 `xml:"urn:com.workday/bsvc Enrollment_Date,omitempty"`
	DisenrollmentDate            *time.Time                                 `xml:"urn:com.workday/bsvc Disenrollment_Date,omitempty"`
	DisenrollmentReasonReference *MedicareEGWPDisenrollmentReasonObjectType `xml:"urn:com.workday/bsvc Disenrollment_Reason_Reference,omitempty"`
	ReasonRejectCodeReference    *MedicareEGWPReasonRejectCodeObjectType    `xml:"urn:com.workday/bsvc Reason_Reject_Code_Reference,omitempty"`
}

func (t *MedicarePartDEGWPDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MedicarePartDEGWPDataType
	var layout struct {
		*T
		EffectiveDate     *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		EnrollmentDate    *xsdDate `xml:"urn:com.workday/bsvc Enrollment_Date,omitempty"`
		DisenrollmentDate *xsdDate `xml:"urn:com.workday/bsvc Disenrollment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	layout.EnrollmentDate = (*xsdDate)(layout.T.EnrollmentDate)
	layout.DisenrollmentDate = (*xsdDate)(layout.T.DisenrollmentDate)
	return e.EncodeElement(layout, start)
}
func (t *MedicarePartDEGWPDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MedicarePartDEGWPDataType
	var overlay struct {
		*T
		EffectiveDate     *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		EnrollmentDate    *xsdDate `xml:"urn:com.workday/bsvc Enrollment_Date,omitempty"`
		DisenrollmentDate *xsdDate `xml:"urn:com.workday/bsvc Disenrollment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	overlay.EnrollmentDate = (*xsdDate)(overlay.T.EnrollmentDate)
	overlay.DisenrollmentDate = (*xsdDate)(overlay.T.DisenrollmentDate)
	return d.DecodeElement(&overlay, &start)
}

// Medicare Part D EGWP Response Data
type MedicarePartDEGWPResponseDataType struct {
	MedicarePartDEGWP []MedicarePartDEGWPType `xml:"urn:com.workday/bsvc Medicare_Part_D_EGWP,omitempty"`
}

// Container for Individual Rows of Medicare Part D EGWP Data
type MedicarePartDEGWPType struct {
	MedicarePartDEGWPData []MedicarePartDEGWPDataType `xml:"urn:com.workday/bsvc Medicare_Part_D_EGWP_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MedicareReasonObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MedicareReasonObjectType struct {
	ID         []MedicareReasonObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Wrapper element for all National Identifier Data.
type NationalIDType struct {
	NationalIDReference       *UniqueIdentifierObjectType            `xml:"urn:com.workday/bsvc National_ID_Reference,omitempty"`
	NationalIDData            *NationalIDDataType                    `xml:"urn:com.workday/bsvc National_ID_Data,omitempty"`
	NationalIDSharedReference *NationalIdentifierReferenceObjectType `xml:"urn:com.workday/bsvc National_ID_Shared_Reference,omitempty"`
	Delete                    bool                                   `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type NationalIdentifierReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type NationalIdentifierReferenceObjectType struct {
	ID         []NationalIdentifierReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Occupation of the Dependent
type OccupationDataType struct {
	EffectiveDate *time.Time `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	Occupation    string     `xml:"urn:com.workday/bsvc Occupation,omitempty"`
}

func (t *OccupationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T OccupationDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *OccupationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T OccupationDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type OrganizationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type OrganizationObjectType struct {
	ID         []OrganizationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Reference element representing a unique instance of Organization.
type OrganizationReferenceType struct {
	IntegrationIDReference *ExternalIntegrationIDReferenceDataType `xml:"urn:com.workday/bsvc Integration_ID_Reference"`
}

// Wrapper for Passport Identifier data.
type PassportIDDataType struct {
	ID                  string                       `xml:"urn:com.workday/bsvc ID,omitempty"`
	IDTypeReference     *PassportIDTypeAllObjectType `xml:"urn:com.workday/bsvc ID_Type_Reference,omitempty"`
	CountryReference    *CountryObjectType           `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	IssuedDate          *time.Time                   `xml:"urn:com.workday/bsvc Issued_Date,omitempty"`
	ExpirationDate      *time.Time                   `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	VerificationDate    *time.Time                   `xml:"urn:com.workday/bsvc Verification_Date,omitempty"`
	VerifiedByReference *WorkerObjectType            `xml:"urn:com.workday/bsvc Verified_By_Reference,omitempty"`
}

func (t *PassportIDDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PassportIDDataType
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
func (t *PassportIDDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PassportIDDataType
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

// Encapsulating element for all Passport Identifier data.
type PassportIDType struct {
	PassportIDReference       *UniqueIdentifierObjectType            `xml:"urn:com.workday/bsvc Passport_ID_Reference,omitempty"`
	PassportIDData            *PassportIDDataType                    `xml:"urn:com.workday/bsvc Passport_ID_Data,omitempty"`
	PassportIDSharedReference *PassportIdentifierReferenceObjectType `xml:"urn:com.workday/bsvc Passport_ID_Shared_Reference,omitempty"`
	Delete                    bool                                   `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PassportIDTypeAllObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PassportIDTypeAllObjectType struct {
	ID         []PassportIDTypeAllObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PassportIdentifierReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PassportIdentifierReferenceObjectType struct {
	ID         []PassportIdentifierReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper for Person Identification Data. Includes National Identifiers, Government Identifiers, Visa Identifiers, Passport Identifiers, License Identifiers and Custom Identifiers.
type PersonIdentificationDataType struct {
	NationalID   []NationalIDType   `xml:"urn:com.workday/bsvc National_ID,omitempty"`
	GovernmentID []GovernmentIDType `xml:"urn:com.workday/bsvc Government_ID,omitempty"`
	VisaID       []VisaIDType       `xml:"urn:com.workday/bsvc Visa_ID,omitempty"`
	PassportID   []PassportIDType   `xml:"urn:com.workday/bsvc Passport_ID,omitempty"`
	LicenseID    []LicenseIDType    `xml:"urn:com.workday/bsvc License_ID,omitempty"`
	CustomID     []CustomIDType     `xml:"urn:com.workday/bsvc Custom_ID,omitempty"`
}

// Contains the legal, preferred, and additional names for a person.
type PersonNameDataType struct {
	LegalNameData      *LegalNameDataType       `xml:"urn:com.workday/bsvc Legal_Name_Data,omitempty"`
	PreferredNameData  *PreferredNameDataType   `xml:"urn:com.workday/bsvc Preferred_Name_Data,omitempty"`
	AdditionalNameData []AdditionalNameDataType `xml:"urn:com.workday/bsvc Additional_Name_Data,omitempty"`
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

// The postal codes to be added to the set. Each code will be checked for validity before being added to the set.
type PostalCodeDataUnverifiedType struct {
	PostalCode             string                   `xml:"urn:com.workday/bsvc Postal_Code"`
	CountryRegionReference *CountryRegionObjectType `xml:"urn:com.workday/bsvc Country_Region_Reference,omitempty"`
}

// Data to use to create or update a postal code set.
type PostalCodeSetDataType struct {
	Name           string                         `xml:"urn:com.workday/bsvc Name"`
	Description    string                         `xml:"urn:com.workday/bsvc Description,omitempty"`
	PostalCodeData []PostalCodeDataUnverifiedType `xml:"urn:com.workday/bsvc Postal_Code_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PostalCodeSetObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PostalCodeSetObjectType struct {
	ID         []PostalCodeSetObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the preferred name for a person.  If no preferred name is returned then the legal name is assumed to be the preferred name.  If a preferred name is not provided in a request then the legal name is assumed to be the preferred name.
type PreferredNameDataType struct {
	NameDetailData *PersonNameDetailDataType `xml:"urn:com.workday/bsvc Name_Detail_Data"`
}

type ProcessingFaultType struct {
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
}

// Wrapper element for the Put Affordable Care Act Worker Hours And Wages Request information.
type PutAffordableCareActWorkerHoursAndWagesRequestType struct {
	AffordableCareActWorkerHoursAndWagesData *AffordableCareActWorkerHoursAndWagesDataType `xml:"urn:com.workday/bsvc Affordable_Care_Act_Worker_Hours_And_Wages_Data"`
	Version                                  string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Put Affordable Care Act Worker Hours Response information.
type PutAffordableCareActWorkerHoursAndWagesResponseType struct {
	AffordableCareActWorkerHoursAndWagesReference *AffordableCareActWorkerHoursAndWagesObjectType `xml:"urn:com.workday/bsvc Affordable_Care_Act_Worker_Hours_And_Wages_Reference,omitempty"`
	Version                                       string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Put Benefit Annual Credit Request information.
type PutBenefitAnnualCreditRequestType struct {
	BenefitAnnualCreditData *BenefitAnnualCreditDataType `xml:"urn:com.workday/bsvc Benefit_Annual_Credit_Data"`
	Version                 string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Put Benefit Annual Credit Response information.
type PutBenefitAnnualCreditResponseType struct {
	BenefitAnnualCreditReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Benefit_Annual_Credit_Reference,omitempty"`
	Version                      string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Put Benefit Annual Rate Request information.
type PutBenefitAnnualRateRequestType struct {
	BenefitAnnualRateData *BenefitAnnualRateDataType `xml:"urn:com.workday/bsvc Benefit_Annual_Rate_Data"`
	Version               string                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Put Benefit Annual Rate Response information.
type PutBenefitAnnualRateResponseType struct {
	BenefitAnnualRateReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Benefit_Annual_Rate_Reference,omitempty"`
	Version                    string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Updates existing Benefit Individual Rates.
type PutBenefitIndividualRateRequestType struct {
	BenefitIndividualRateReference *BenefitIndividualRateObjectType `xml:"urn:com.workday/bsvc Benefit_Individual_Rate_Reference,omitempty"`
	BenefitIndividualRateData      *BenefitIndividualRateDataType   `xml:"urn:com.workday/bsvc Benefit_Individual_Rate_Data,omitempty"`
	Version                        string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response for a Put Benefit Individual Rate web service request
type PutBenefitIndividualRateResponseType struct {
	BenefitIndividualRateReference *BenefitIndividualRateObjectType `xml:"urn:com.workday/bsvc Benefit_Individual_Rate_Reference,omitempty"`
	Version                        string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The Wrapper data element for the Put Dependent Web Service.
type PutDependentDataType struct {
	DependentID                        string                                               `xml:"urn:com.workday/bsvc Dependent_ID,omitempty"`
	EmployeeReference                  *EmployeeObjectType                                  `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	ExistingRelatedPersonReference     *RelatedPersonforWorkerObjectType                    `xml:"urn:com.workday/bsvc Existing_Related_Person_Reference,omitempty"`
	RelatedPersonRelationshipReference *RelatedPersonRelationshipObjectType                 `xml:"urn:com.workday/bsvc Related_Person_Relationship_Reference,omitempty"`
	UseEmployeeAddress                 *bool                                                `xml:"urn:com.workday/bsvc Use_Employee_Address,omitempty"`
	UseEmployeePhone                   *bool                                                `xml:"urn:com.workday/bsvc Use_Employee_Phone,omitempty"`
	Irrevocable                        *bool                                                `xml:"urn:com.workday/bsvc Irrevocable,omitempty"`
	DependentPersonalInformationData   *DependentPersonalInformationDataType                `xml:"urn:com.workday/bsvc Dependent_Personal_Information_Data,omitempty"`
	CourtOrderReplacementData          []QualifiedDomesticRelationsOrderReplacementDataType `xml:"urn:com.workday/bsvc Court_Order_Replacement_Data,omitempty"`
}

// Wrapper element for the Put Dependent Request Information.
type PutDependentRequestType struct {
	DependentReference *DependentObjectType  `xml:"urn:com.workday/bsvc Dependent_Reference,omitempty"`
	DependentData      *PutDependentDataType `xml:"urn:com.workday/bsvc Dependent_Data"`
	AddOnly            bool                  `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version            string                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Put Dependent Returned data.
type PutDependentResponseType struct {
	DependentReference     *DependentObjectType                           `xml:"urn:com.workday/bsvc Dependent_Reference,omitempty"`
	EmployeeReference      *EmployeeObjectType                            `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	ExceptionsResponseData []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the retirement savings elections to be added, updated or deleted for an employee.
type PutEmployeeDefinedContributionElectionsRequestType struct {
	EmployeeDefinedContributionElectionData *EmployeeDefinedContributionElectionDataType `xml:"urn:com.workday/bsvc Employee_Defined_Contribution_Election_Data"`
	AddOnly                                 bool                                         `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                                 string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a reference to the employee and takes effect on date for the retirement savings elections that were added or updated for the request.
type PutEmployeeDefinedContributionElectionsResponseType struct {
	EmployeeReference *EmployeeObjectType `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	TakesEffectOnDate *time.Time          `xml:"urn:com.workday/bsvc Takes_Effect_On_Date,omitempty"`
	Version           string              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

func (t *PutEmployeeDefinedContributionElectionsResponseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PutEmployeeDefinedContributionElectionsResponseType
	var layout struct {
		*T
		TakesEffectOnDate *xsdDate `xml:"urn:com.workday/bsvc Takes_Effect_On_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TakesEffectOnDate = (*xsdDate)(layout.T.TakesEffectOnDate)
	return e.EncodeElement(layout, start)
}
func (t *PutEmployeeDefinedContributionElectionsResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PutEmployeeDefinedContributionElectionsResponseType
	var overlay struct {
		*T
		TakesEffectOnDate *xsdDate `xml:"urn:com.workday/bsvc Takes_Effect_On_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TakesEffectOnDate = (*xsdDate)(overlay.T.TakesEffectOnDate)
	return d.DecodeElement(&overlay, &start)
}

// Update Evidence of Insurability
type PutEvidenceofInsurabilityRequestType struct {
	EvidenceofInsurabilityData *EvidenceofInsurabilityDataType `xml:"urn:com.workday/bsvc Evidence_of_Insurability_Data"`
	Version                    string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Update Evidence of Insurability
type PutEvidenceofInsurabilityResponseType struct {
	EvidenceofInsurabilityReference *EvidenceofInsurabilityObjectType `xml:"urn:com.workday/bsvc Evidence_of_Insurability_Reference,omitempty"`
	Version                         string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Health Care Rate Request
type PutHealthCareRateRequestType struct {
	HealthCareRateReference *HealthCareBandedRateObjectType `xml:"urn:com.workday/bsvc Health_Care_Rate_Reference,omitempty"`
	HealthCareRateData      *HealthCareRateDataType         `xml:"urn:com.workday/bsvc Health_Care_Rate_Data"`
	Version                 string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Health Care Rate Response EL
type PutHealthCareRateResponseType struct {
	HealthCareRateReference *HealthCareBandedRateObjectType `xml:"urn:com.workday/bsvc Health_Care_Rate_Reference,omitempty"`
	Version                 string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Import Process Response element
type PutImportProcessResponseType struct {
	ImportProcessReference  *WebServiceBackgroundProcessRuntimeObjectType `xml:"urn:com.workday/bsvc Import_Process_Reference,omitempty"`
	HeaderInstanceReference *InstanceObjectType                           `xml:"urn:com.workday/bsvc Header_Instance_Reference,omitempty"`
	Version                 string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top-Level Request
type PutPostalCodeSetRequestType struct {
	PostalCodeSetReference *PostalCodeSetObjectType `xml:"urn:com.workday/bsvc Postal_Code_Set_Reference,omitempty"`
	PostalCodeSetData      *PostalCodeSetDataType   `xml:"urn:com.workday/bsvc Postal_Code_Set_Data"`
	Version                string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Put Postal Code Set Returned data.
type PutPostalCodeSetResponseType struct {
	PostalCodeSetReference *PostalCodeSetObjectType `xml:"urn:com.workday/bsvc Postal_Code_Set_Reference,omitempty"`
	Version                string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Put Worker Wellness Data Request information.
type PutWorkerWellnessDataRequestType struct {
	WorkerWellnessDataData *WorkerWellnessDataDataType `xml:"urn:com.workday/bsvc Worker_Wellness_Data_Data"`
	Version                string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Put Worker Wellness Data Response information.
type PutWorkerWellnessDataResponseType struct {
	WorkerWellnessDataReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Worker_Wellness_Data_Reference,omitempty"`
	Version                     string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Court Order Details. NOTE: You must pass in the entire set of court orders. Any existing court orders that are not submitted will be removed.
type QualifiedDomesticRelationsOrderReplacementDataType struct {
	BenefitCoverageTypeReference *BenefitCoverageTypeObjectType `xml:"urn:com.workday/bsvc Benefit_Coverage_Type_Reference"`
	StartDate                    time.Time                      `xml:"urn:com.workday/bsvc Start_Date"`
	EndDate                      time.Time                      `xml:"urn:com.workday/bsvc End_Date"`
}

func (t *QualifiedDomesticRelationsOrderReplacementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T QualifiedDomesticRelationsOrderReplacementDataType
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
func (t *QualifiedDomesticRelationsOrderReplacementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T QualifiedDomesticRelationsOrderReplacementDataType
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

// All of the personal data. This includes name, contact information (Email, Phone, Address, Web, Instant Messenger), Visa and ID information, Biographic, Demographic, and Background Check Information.
type RelatedPersonPersonalInformationDataType struct {
	PersonNameData                *LegalNameDataType                             `xml:"urn:com.workday/bsvc Person_Name_Data,omitempty"`
	ContactInformationData        *ContactInformationDataForRelatedPersonType    `xml:"urn:com.workday/bsvc Contact_Information_Data,omitempty"`
	PersonIdentificationData      *PersonIdentificationDataType                  `xml:"urn:com.workday/bsvc Person_Identification_Data,omitempty"`
	DateofBirth                   *time.Time                                     `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	DateofDeath                   *time.Time                                     `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
	GenderReference               *GenderObjectType                              `xml:"urn:com.workday/bsvc Gender_Reference,omitempty"`
	UsesTobacco                   *bool                                          `xml:"urn:com.workday/bsvc Uses_Tobacco,omitempty"`
	FullTimeStudent               *bool                                          `xml:"urn:com.workday/bsvc Full-Time_Student,omitempty"`
	StudentStatusStartDate        *time.Time                                     `xml:"urn:com.workday/bsvc Student_Status_Start_Date,omitempty"`
	StudentStatusEndDate          *time.Time                                     `xml:"urn:com.workday/bsvc Student_Status_End_Date,omitempty"`
	Disabled                      *bool                                          `xml:"urn:com.workday/bsvc Disabled,omitempty"`
	DisabilityData                *DisabilityInformationDataforRelatedPersonType `xml:"urn:com.workday/bsvc Disability_Data,omitempty"`
	InactiveDate                  *time.Time                                     `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
	DependentforPayrollPurposes   *bool                                          `xml:"urn:com.workday/bsvc Dependent_for_Payroll_Purposes,omitempty"`
	CitizenshipStatusReference    []CitizenshipStatusObjectType                  `xml:"urn:com.workday/bsvc Citizenship_Status_Reference,omitempty"`
	CountryofNationalityReference *CountryObjectType                             `xml:"urn:com.workday/bsvc Country_of_Nationality_Reference,omitempty"`
	CountryofBirthReference       *CountryObjectType                             `xml:"urn:com.workday/bsvc Country_of_Birth_Reference,omitempty"`
	RegionofBirthReference        *CountryRegionObjectType                       `xml:"urn:com.workday/bsvc Region_of_Birth_Reference,omitempty"`
	CityofBirth                   string                                         `xml:"urn:com.workday/bsvc City_of_Birth,omitempty"`
	LivesWithWorkerData           *LivesWithWorkerDataType                       `xml:"urn:com.workday/bsvc Lives_With_Worker_Data,omitempty"`
	HasHealthInsuranceData        *HasHealthInsuranceDataType                    `xml:"urn:com.workday/bsvc Has_Health_Insurance_Data,omitempty"`
	AllowedforTaxDeductionData    *AllowedforTaxDeductionDataType                `xml:"urn:com.workday/bsvc Allowed_for_Tax_Deduction_Data,omitempty"`
	AnnualIncomeData              *AnnualIncomeDataType                          `xml:"urn:com.workday/bsvc Annual_Income_Data,omitempty"`
	OccupationData                *OccupationDataType                            `xml:"urn:com.workday/bsvc Occupation_Data,omitempty"`
}

func (t *RelatedPersonPersonalInformationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RelatedPersonPersonalInformationDataType
	var layout struct {
		*T
		DateofBirth            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
		DateofDeath            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
		StudentStatusStartDate *xsdDate `xml:"urn:com.workday/bsvc Student_Status_Start_Date,omitempty"`
		StudentStatusEndDate   *xsdDate `xml:"urn:com.workday/bsvc Student_Status_End_Date,omitempty"`
		InactiveDate           *xsdDate `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateofBirth = (*xsdDate)(layout.T.DateofBirth)
	layout.DateofDeath = (*xsdDate)(layout.T.DateofDeath)
	layout.StudentStatusStartDate = (*xsdDate)(layout.T.StudentStatusStartDate)
	layout.StudentStatusEndDate = (*xsdDate)(layout.T.StudentStatusEndDate)
	layout.InactiveDate = (*xsdDate)(layout.T.InactiveDate)
	return e.EncodeElement(layout, start)
}
func (t *RelatedPersonPersonalInformationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RelatedPersonPersonalInformationDataType
	var overlay struct {
		*T
		DateofBirth            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
		DateofDeath            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
		StudentStatusStartDate *xsdDate `xml:"urn:com.workday/bsvc Student_Status_Start_Date,omitempty"`
		StudentStatusEndDate   *xsdDate `xml:"urn:com.workday/bsvc Student_Status_End_Date,omitempty"`
		InactiveDate           *xsdDate `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateofBirth = (*xsdDate)(overlay.T.DateofBirth)
	overlay.DateofDeath = (*xsdDate)(overlay.T.DateofDeath)
	overlay.StudentStatusStartDate = (*xsdDate)(overlay.T.StudentStatusStartDate)
	overlay.StudentStatusEndDate = (*xsdDate)(overlay.T.StudentStatusEndDate)
	overlay.InactiveDate = (*xsdDate)(overlay.T.InactiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type RelatedPersonRelationshipObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RelatedPersonRelationshipObjectType struct {
	ID         []RelatedPersonRelationshipObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RelatedPersonforWorkerObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RelatedPersonforWorkerObjectType struct {
	ID         []RelatedPersonforWorkerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Retirement Savings Election data
type RetirementSavingsElectionDataType struct {
	RetirementSavingsPlanReference *RetirementSavingsPlanObjectType                 `xml:"urn:com.workday/bsvc Retirement_Savings_Plan_Reference"`
	ElectionPercentage             float64                                          `xml:"urn:com.workday/bsvc Election_Percentage,omitempty"`
	ElectionAmount                 float64                                          `xml:"urn:com.workday/bsvc Election_Amount,omitempty"`
	BeneficiaryAllocationData      []BeneficiaryAllocationforChangeBenefitsDataType `xml:"urn:com.workday/bsvc Beneficiary_Allocation_Data,omitempty"`
}

// Retirement Savings Election Data
type RetirementSavingsElectionDataforChangeRetirementSavingsPlansType struct {
	CoverageBeginDate                time.Time                                        `xml:"urn:com.workday/bsvc Coverage_Begin_Date"`
	OriginalCoverageBeginDate        *time.Time                                       `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
	DeductionBeginDate               *time.Time                                       `xml:"urn:com.workday/bsvc Deduction_Begin_Date,omitempty"`
	LoadedOriginalDeductionBeginDate *time.Time                                       `xml:"urn:com.workday/bsvc Loaded_Original_Deduction_Begin_Date,omitempty"`
	Elect                            *bool                                            `xml:"urn:com.workday/bsvc Elect,omitempty"`
	RetirementSavingsPlanReference   *RetirementSavingsPlanObjectType                 `xml:"urn:com.workday/bsvc Retirement_Savings_Plan_Reference"`
	ElectionPercentage               float64                                          `xml:"urn:com.workday/bsvc Election_Percentage,omitempty"`
	ElectionAmount                   float64                                          `xml:"urn:com.workday/bsvc Election_Amount,omitempty"`
	BeneficiaryAllocationData        []BeneficiaryAllocationforChangeBenefitsDataType `xml:"urn:com.workday/bsvc Beneficiary_Allocation_Data,omitempty"`
}

func (t *RetirementSavingsElectionDataforChangeRetirementSavingsPlansType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RetirementSavingsElectionDataforChangeRetirementSavingsPlansType
	var layout struct {
		*T
		CoverageBeginDate                *xsdDate `xml:"urn:com.workday/bsvc Coverage_Begin_Date"`
		OriginalCoverageBeginDate        *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
		DeductionBeginDate               *xsdDate `xml:"urn:com.workday/bsvc Deduction_Begin_Date,omitempty"`
		LoadedOriginalDeductionBeginDate *xsdDate `xml:"urn:com.workday/bsvc Loaded_Original_Deduction_Begin_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CoverageBeginDate = (*xsdDate)(&layout.T.CoverageBeginDate)
	layout.OriginalCoverageBeginDate = (*xsdDate)(layout.T.OriginalCoverageBeginDate)
	layout.DeductionBeginDate = (*xsdDate)(layout.T.DeductionBeginDate)
	layout.LoadedOriginalDeductionBeginDate = (*xsdDate)(layout.T.LoadedOriginalDeductionBeginDate)
	return e.EncodeElement(layout, start)
}
func (t *RetirementSavingsElectionDataforChangeRetirementSavingsPlansType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RetirementSavingsElectionDataforChangeRetirementSavingsPlansType
	var overlay struct {
		*T
		CoverageBeginDate                *xsdDate `xml:"urn:com.workday/bsvc Coverage_Begin_Date"`
		OriginalCoverageBeginDate        *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
		DeductionBeginDate               *xsdDate `xml:"urn:com.workday/bsvc Deduction_Begin_Date,omitempty"`
		LoadedOriginalDeductionBeginDate *xsdDate `xml:"urn:com.workday/bsvc Loaded_Original_Deduction_Begin_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CoverageBeginDate = (*xsdDate)(&overlay.T.CoverageBeginDate)
	overlay.OriginalCoverageBeginDate = (*xsdDate)(overlay.T.OriginalCoverageBeginDate)
	overlay.DeductionBeginDate = (*xsdDate)(overlay.T.DeductionBeginDate)
	overlay.LoadedOriginalDeductionBeginDate = (*xsdDate)(overlay.T.LoadedOriginalDeductionBeginDate)
	return d.DecodeElement(&overlay, &start)
}

// Benefit Coverage Type level values
type RetirementSavingsElectionforCoverageTypeDataType struct {
	CoverageBeginDate                     time.Time                                        `xml:"urn:com.workday/bsvc Coverage_Begin_Date"`
	OriginalCoverageBeginDate             *time.Time                                       `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
	DeductionBeginDate                    *time.Time                                       `xml:"urn:com.workday/bsvc Deduction_Begin_Date,omitempty"`
	OriginalDeductionBeginDate            *time.Time                                       `xml:"urn:com.workday/bsvc Original_Deduction_Begin_Date,omitempty"`
	RetirementSavingsPlanReference        *RetirementSavingsPlanObjectType                 `xml:"urn:com.workday/bsvc Retirement_Savings_Plan_Reference"`
	EmployeeContributionAllocationPercent float64                                          `xml:"urn:com.workday/bsvc Employee_Contribution_Allocation_Percent,omitempty"`
	EmployerContributionAllocationPercent float64                                          `xml:"urn:com.workday/bsvc Employer_Contribution_Allocation_Percent,omitempty"`
	BeneficiaryAllocationData             []BeneficiaryAllocationforChangeBenefitsDataType `xml:"urn:com.workday/bsvc Beneficiary_Allocation_Data,omitempty"`
}

func (t *RetirementSavingsElectionforCoverageTypeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RetirementSavingsElectionforCoverageTypeDataType
	var layout struct {
		*T
		CoverageBeginDate          *xsdDate `xml:"urn:com.workday/bsvc Coverage_Begin_Date"`
		OriginalCoverageBeginDate  *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
		DeductionBeginDate         *xsdDate `xml:"urn:com.workday/bsvc Deduction_Begin_Date,omitempty"`
		OriginalDeductionBeginDate *xsdDate `xml:"urn:com.workday/bsvc Original_Deduction_Begin_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CoverageBeginDate = (*xsdDate)(&layout.T.CoverageBeginDate)
	layout.OriginalCoverageBeginDate = (*xsdDate)(layout.T.OriginalCoverageBeginDate)
	layout.DeductionBeginDate = (*xsdDate)(layout.T.DeductionBeginDate)
	layout.OriginalDeductionBeginDate = (*xsdDate)(layout.T.OriginalDeductionBeginDate)
	return e.EncodeElement(layout, start)
}
func (t *RetirementSavingsElectionforCoverageTypeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RetirementSavingsElectionforCoverageTypeDataType
	var overlay struct {
		*T
		CoverageBeginDate          *xsdDate `xml:"urn:com.workday/bsvc Coverage_Begin_Date"`
		OriginalCoverageBeginDate  *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
		DeductionBeginDate         *xsdDate `xml:"urn:com.workday/bsvc Deduction_Begin_Date,omitempty"`
		OriginalDeductionBeginDate *xsdDate `xml:"urn:com.workday/bsvc Original_Deduction_Begin_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CoverageBeginDate = (*xsdDate)(&overlay.T.CoverageBeginDate)
	overlay.OriginalCoverageBeginDate = (*xsdDate)(overlay.T.OriginalCoverageBeginDate)
	overlay.DeductionBeginDate = (*xsdDate)(overlay.T.DeductionBeginDate)
	overlay.OriginalDeductionBeginDate = (*xsdDate)(overlay.T.OriginalDeductionBeginDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type RetirementSavingsPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RetirementSavingsPlanObjectType struct {
	ID         []RetirementSavingsPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Salary for Health Care Rate
type SalaryDataType struct {
	CombineAllJobsforSalarySource  *bool                              `xml:"urn:com.workday/bsvc Combine_All_Jobs_for_Salary_Source,omitempty"`
	BenefitSalarySourceReference   *BenefitSalarySourceObjectType     `xml:"urn:com.workday/bsvc Benefit_Salary_Source_Reference"`
	CompensationElementReference   []CompensationPayEarningObjectType `xml:"urn:com.workday/bsvc Compensation_Element_Reference,omitempty"`
	FrequencyReference             *FrequencyObjectType               `xml:"urn:com.workday/bsvc Frequency_Reference"`
	BenefitAnnualRateTypeReference *BenefitAnnualRateTypeObjectType   `xml:"urn:com.workday/bsvc Benefit_Annual_Rate_Type_Reference,omitempty"`
	AsOfDateData                   *AsOfDateDataType                  `xml:"urn:com.workday/bsvc As_Of_Date_Data"`
}

// Election data for spending account elections
type SpendingAccountElectionDataType struct {
	SpendingAccountPlanReference *SpendingAccountPlanObjectType                   `xml:"urn:com.workday/bsvc Spending_Account_Plan_Reference"`
	YTDContributionAmount        float64                                          `xml:"urn:com.workday/bsvc YTD_Contribution_Amount,omitempty"`
	AnnualContribution           float64                                          `xml:"urn:com.workday/bsvc Annual_Contribution,omitempty"`
	EmployeeCost                 float64                                          `xml:"urn:com.workday/bsvc Employee_Cost,omitempty"`
	BeneficiaryAllocationData    []BeneficiaryAllocationforChangeBenefitsDataType `xml:"urn:com.workday/bsvc Beneficiary_Allocation_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SpendingAccountPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SpendingAccountPlanObjectType struct {
	ID         []SpendingAccountPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Personal data for the Trustee. This includes name and contact information.
type TrusteePersonalInformationDataType struct {
	NameData               []PersonNameDetailDataType                  `xml:"urn:com.workday/bsvc Name_Data,omitempty"`
	ContactInformationData *ContactInformationDataForRelatedPersonType `xml:"urn:com.workday/bsvc Contact_Information_Data,omitempty"`
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

// Wrapper element for Visa identifier data.
type VisaIDDataType struct {
	ID                  string                `xml:"urn:com.workday/bsvc ID,omitempty"`
	IDTypeReference     *VisaIDTypeObjectType `xml:"urn:com.workday/bsvc ID_Type_Reference,omitempty"`
	CountryReference    *CountryObjectType    `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	IssuedDate          *time.Time            `xml:"urn:com.workday/bsvc Issued_Date,omitempty"`
	ExpirationDate      *time.Time            `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	VerificationDate    *time.Time            `xml:"urn:com.workday/bsvc Verification_Date,omitempty"`
	VerifiedByReference *WorkerObjectType     `xml:"urn:com.workday/bsvc Verified_By_Reference,omitempty"`
}

func (t *VisaIDDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T VisaIDDataType
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
func (t *VisaIDDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T VisaIDDataType
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

// Encapsulating element for all Visa Identifier data.
type VisaIDType struct {
	VisaIDReference       *UniqueIdentifierObjectType        `xml:"urn:com.workday/bsvc Visa_ID_Reference,omitempty"`
	VisaIDData            *VisaIDDataType                    `xml:"urn:com.workday/bsvc Visa_ID_Data,omitempty"`
	VisaIDSharedReference *VisaIdentifierReferenceObjectType `xml:"urn:com.workday/bsvc Visa_ID_Shared_Reference,omitempty"`
	Delete                bool                               `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type VisaIDTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type VisaIDTypeObjectType struct {
	ID         []VisaIDTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type VisaIdentifierReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type VisaIdentifierReferenceObjectType struct {
	ID         []VisaIdentifierReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

type WorkdayCommonHeaderType struct {
	IncludeReferenceDescriptorsInResponse *bool `xml:"urn:com.workday/bsvc Include_Reference_Descriptors_In_Response,omitempty"`
}

// Container to specify the Worker and select their benefit jobs, or to specify that the HR Primary job should be used for eligibility.
type WorkerBenefitJobsDataType struct {
	WorkerReference *WorkerObjectType    `xml:"urn:com.workday/bsvc Worker_Reference"`
	BenefitJobsData *BenefitJobsDataType `xml:"urn:com.workday/bsvc Benefit_Jobs_Data"`
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

// Contains the Worker Wellness Data data.
type WorkerWellnessDataDataType struct {
	EffectiveDate                 *time.Time           `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	WorkerReference               *WorkerObjectType    `xml:"urn:com.workday/bsvc Worker_Reference"`
	DependentReference            *DependentObjectType `xml:"urn:com.workday/bsvc Dependent_Reference"`
	ParticipatesinWellnessDate    *time.Time           `xml:"urn:com.workday/bsvc Participates_in_Wellness_Date,omitempty"`
	ParticipatesinWellnessProgram *bool                `xml:"urn:com.workday/bsvc Participates_in_Wellness_Program,omitempty"`
	WellnessScoreDate             *time.Time           `xml:"urn:com.workday/bsvc Wellness_Score_Date,omitempty"`
	WellnessScore                 float64              `xml:"urn:com.workday/bsvc Wellness_Score,omitempty"`
	UsesTobaccoDate               *time.Time           `xml:"urn:com.workday/bsvc Uses_Tobacco_Date,omitempty"`
	UsesTobacco                   *bool                `xml:"urn:com.workday/bsvc Uses_Tobacco,omitempty"`
}

func (t *WorkerWellnessDataDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerWellnessDataDataType
	var layout struct {
		*T
		EffectiveDate              *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		ParticipatesinWellnessDate *xsdDate `xml:"urn:com.workday/bsvc Participates_in_Wellness_Date,omitempty"`
		WellnessScoreDate          *xsdDate `xml:"urn:com.workday/bsvc Wellness_Score_Date,omitempty"`
		UsesTobaccoDate            *xsdDate `xml:"urn:com.workday/bsvc Uses_Tobacco_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	layout.ParticipatesinWellnessDate = (*xsdDate)(layout.T.ParticipatesinWellnessDate)
	layout.WellnessScoreDate = (*xsdDate)(layout.T.WellnessScoreDate)
	layout.UsesTobaccoDate = (*xsdDate)(layout.T.UsesTobaccoDate)
	return e.EncodeElement(layout, start)
}
func (t *WorkerWellnessDataDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerWellnessDataDataType
	var overlay struct {
		*T
		EffectiveDate              *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		ParticipatesinWellnessDate *xsdDate `xml:"urn:com.workday/bsvc Participates_in_Wellness_Date,omitempty"`
		WellnessScoreDate          *xsdDate `xml:"urn:com.workday/bsvc Wellness_Score_Date,omitempty"`
		UsesTobaccoDate            *xsdDate `xml:"urn:com.workday/bsvc Uses_Tobacco_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	overlay.ParticipatesinWellnessDate = (*xsdDate)(overlay.T.ParticipatesinWellnessDate)
	overlay.WellnessScoreDate = (*xsdDate)(overlay.T.WellnessScoreDate)
	overlay.UsesTobaccoDate = (*xsdDate)(overlay.T.UsesTobaccoDate)
	return d.DecodeElement(&overlay, &start)
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
