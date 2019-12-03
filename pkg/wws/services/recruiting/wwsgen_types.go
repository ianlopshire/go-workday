package recruiting

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

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

// Contains a unique identifier for an instance of an object.
type AllowanceUnitPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AllowanceUnitPlanObjectType struct {
	ID         []AllowanceUnitPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AllowanceValuePlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AllowanceValuePlanObjectType struct {
	ID         []AllowanceValuePlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The main Wrapper element for the Applicant Data Web Service.
type ApplicantInformationHVType struct {
	ApplicantReference *ApplicantObjectType `xml:"urn:com.workday/bsvc Applicant_Reference,omitempty"`
	ApplicantData      []PreHireDataWWSType `xml:"urn:com.workday/bsvc Applicant_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ApplicantObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ApplicantObjectType struct {
	ID         []ApplicantObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Utilize the following criteria options to search for Applicants within the Workday system.  The Applicant references that are returned are those that satisfy ALL criteria included in the request.  Therefore, the result set will become more limited with every criterium that is populated.
type ApplicantRequestCriteriaType struct {
	WorkerReference               *WorkerObjectType                  `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	FormerWorkerReference         *FormerWorkerObjectType            `xml:"urn:com.workday/bsvc Former_Worker_Reference,omitempty"`
	EmailAddress                  string                             `xml:"urn:com.workday/bsvc Email_Address,omitempty"`
	FieldAndParameterCriteriaData *FieldAndParameterCriteriaDataType `xml:"urn:com.workday/bsvc Field_And_Parameter_Criteria_Data,omitempty"`
}

// Wrapper element  containing the list of applicants for which data is requested.
type ApplicantRequestReferencesType struct {
	ApplicantReference []ApplicantObjectType `xml:"urn:com.workday/bsvc Applicant_Reference"`
}

// Wrapper element containing the list of applicant information to be included in the response.
type ApplicantResponseGroupType struct {
	IncludeReference                 *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludePersonalInformation       *bool `xml:"urn:com.workday/bsvc Include_Personal_Information,omitempty"`
	IncludeRecruitingInformation     *bool `xml:"urn:com.workday/bsvc Include_Recruiting_Information,omitempty"`
	IncludeQualificationProfile      *bool `xml:"urn:com.workday/bsvc Include_Qualification_Profile,omitempty"`
	IncludeResume                    *bool `xml:"urn:com.workday/bsvc Include_Resume,omitempty"`
	IncludeBackgroundCheck           *bool `xml:"urn:com.workday/bsvc Include_Background_Check,omitempty"`
	IncludeExternalIntegrationIDData *bool `xml:"urn:com.workday/bsvc Include_External_Integration_ID_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ApplicantResumeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ApplicantResumeObjectType struct {
	ID         []ApplicantResumeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ApplicantSourceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ApplicantSourceObjectType struct {
	ID         []ApplicantSourceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for the Applicant Data.
type ApplicantWWSType struct {
	ApplicantReference *ApplicantObjectType `xml:"urn:com.workday/bsvc Applicant_Reference,omitempty"`
	ApplicantData      *PreHireDataWWSType  `xml:"urn:com.workday/bsvc Applicant_Data,omitempty"`
}

// Wrapper element for the Applicant Data.
type ApplicantsResponseDataType struct {
	Applicant []ApplicantWWSType `xml:"urn:com.workday/bsvc Applicant,omitempty"`
}

// Element containing application related exceptions data
type ApplicationInstanceExceptionsDataType struct {
	ExceptionData []ExceptionDataType `xml:"urn:com.workday/bsvc Exception_Data,omitempty"`
}

// Element containing Exceptions Data
type ApplicationInstanceRelatedExceptionsDataType struct {
	ExceptionsData []ApplicationInstanceExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Data,omitempty"`
}

// Contains Candidate Personal Info: Gender, Ethnicity, Hispanic/Latino, Military Service, Disabilities
type ApplicationPersonalInformationDataType struct {
	GenderReference           *GenderObjectType                               `xml:"urn:com.workday/bsvc Gender_Reference,omitempty"`
	EthnicityReference        []EthnicityObjectType                           `xml:"urn:com.workday/bsvc Ethnicity_Reference,omitempty"`
	VeteransStatusReference   []ArmedForcesStatusObjectType                   `xml:"urn:com.workday/bsvc Veterans_Status_Reference,omitempty"`
	HispanicorLatino          *bool                                           `xml:"urn:com.workday/bsvc Hispanic_or_Latino,omitempty"`
	DisabilityStatusReference *SelfIdentificationofDisabilityStatusObjectType `xml:"urn:com.workday/bsvc Disability_Status_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ArmedForcesStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ArmedForcesStatusObjectType struct {
	ID         []ArmedForcesStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// References to the Candidate Assessment to be created or modified along with the data to be used in the creation or modification.
type AssessCandidateEventDataType struct {
	EventReference          *UniqueIdentifierObjectType   `xml:"urn:com.workday/bsvc Event_Reference"`
	JobApplicationReference *JobApplicationObjectType     `xml:"urn:com.workday/bsvc Job_Application_Reference"`
	CandidateCriteriaData   *CandidateCriteriaType        `xml:"urn:com.workday/bsvc Candidate_Criteria_Data"`
	CandidateAssessmentData *RecruitingAssessmentDataType `xml:"urn:com.workday/bsvc Candidate_Assessment_Data"`
}

// Data relating to a Candidate Assessment Event.
type AssessCandidateEventType struct {
	EventReference               *UniqueIdentifierObjectType      `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	JobApplicationReference      *JobApplicationObjectType        `xml:"urn:com.workday/bsvc Job_Application_Reference,omitempty"`
	CandidateReference           *CandidateObjectType             `xml:"urn:com.workday/bsvc Candidate_Reference,omitempty"`
	JobRequisitionReference      *JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	JobApplicationEventReference *UniqueIdentifierObjectType      `xml:"urn:com.workday/bsvc Job_Application_Event_Reference,omitempty"`
	CandidateAssessmentData      *RecruitingAssessmentDataType    `xml:"urn:com.workday/bsvc Candidate_Assessment_Data,omitempty"`
}

// Criteria used to determine which Candidate Assessment to return.
type AssessCandidateRequestCriteriaType struct {
	JobApplicationReference      *JobApplicationObjectType   `xml:"urn:com.workday/bsvc Job_Application_Reference,omitempty"`
	CandidateCriteriaData        *CandidateCriteriaType      `xml:"urn:com.workday/bsvc Candidate_Criteria_Data,omitempty"`
	JobApplicationEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Job_Application_Event_Reference,omitempty"`
	IncludeInlineAssessments     *bool                       `xml:"urn:com.workday/bsvc Include_Inline_Assessments,omitempty"`
	FromMoment                   *time.Time                  `xml:"urn:com.workday/bsvc From_Moment,omitempty"`
	ToMoment                     *time.Time                  `xml:"urn:com.workday/bsvc To_Moment,omitempty"`
}

func (t *AssessCandidateRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AssessCandidateRequestCriteriaType
	var layout struct {
		*T
		FromMoment *xsdDateTime `xml:"urn:com.workday/bsvc From_Moment,omitempty"`
		ToMoment   *xsdDateTime `xml:"urn:com.workday/bsvc To_Moment,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FromMoment = (*xsdDateTime)(layout.T.FromMoment)
	layout.ToMoment = (*xsdDateTime)(layout.T.ToMoment)
	return e.EncodeElement(layout, start)
}
func (t *AssessCandidateRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AssessCandidateRequestCriteriaType
	var overlay struct {
		*T
		FromMoment *xsdDateTime `xml:"urn:com.workday/bsvc From_Moment,omitempty"`
		ToMoment   *xsdDateTime `xml:"urn:com.workday/bsvc To_Moment,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FromMoment = (*xsdDateTime)(overlay.T.FromMoment)
	overlay.ToMoment = (*xsdDateTime)(overlay.T.ToMoment)
	return d.DecodeElement(&overlay, &start)
}

// References provided by the web service caller to indicate the Candidate Assessment information to return.
type AssessCandidateRequestReferencesType struct {
	AssessCandidateReference []UniqueIdentifierObjectType     `xml:"urn:com.workday/bsvc Assess_Candidate_Reference"`
	AssessmentReference      []RecruitingAssessmentObjectType `xml:"urn:com.workday/bsvc Assessment_Reference"`
}

// A request to create or modify a Candidate Assessment.
type AssessCandidateRequestType struct {
	DynamicBusinessProcessParameters *DynamicBusinessProcessParametersType `xml:"urn:com.workday/bsvc Dynamic_Business_Process_Parameters,omitempty"`
	AssessCandidateData              *AssessCandidateEventDataType         `xml:"urn:com.workday/bsvc Assess_Candidate_Data"`
	Version                          string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The result of a Candidate Assessment request.  Will contain the result of the added or modified Candidate Assessment.
type AssessCandidateResponseDataType struct {
	AssessCandidateData  []AssessCandidateEventType `xml:"urn:com.workday/bsvc Assess_Candidate_Data,omitempty"`
	InlineAssessmentData []InlineAssessmentDataType `xml:"urn:com.workday/bsvc Inline_Assessment_Data,omitempty"`
}

// The result of an attempt to create or modify a Candidate Assessment.
type AssessCandidateResponseType struct {
	CandidateAssessmentEventReference *UniqueIdentifierObjectType     `xml:"urn:com.workday/bsvc Candidate_Assessment_Event_Reference,omitempty"`
	CandidateAssessmentReference      *RecruitingAssessmentObjectType `xml:"urn:com.workday/bsvc Candidate_Assessment_Reference,omitempty"`
	Version                           string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Specifies the costing allocation for the position.
type AssignCostingAllocationEventDataType struct {
	CostingOverrideData               []CostingAllocationIntervalDataType `xml:"urn:com.workday/bsvc Costing_Override_Data"`
	CostingAllocationLevelReference   *CostingAllocationLevelObjectType   `xml:"urn:com.workday/bsvc Costing_Allocation_Level_Reference"`
	CostingAllocationEarningReference *EarningAllObjectType               `xml:"urn:com.workday/bsvc Costing_Allocation_Earning_Reference,omitempty"`
}

// Contains optional costing allocation override values
type AssignCostingAllocationSubBusinessProcessType struct {
	BusinessSubProcessParameters            *BusinessSubProcessParametersType     `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	PositionCostingAllocationAssignmentData *AssignCostingAllocationEventDataType `xml:"urn:com.workday/bsvc Position_Costing_Allocation_Assignment_Data,omitempty"`
}

// Wrapper for Organization Role Assignment Data.  Includes Role Assignments.
type AssignOrganizationRolesEventDataType struct {
	RoleAssigneeReference *RoleeObjectType     `xml:"urn:com.workday/bsvc Role_Assignee_Reference,omitempty"`
	EffectiveDate         *time.Time           `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	RoleAssignment        []RoleAssignmentType `xml:"urn:com.workday/bsvc Role_Assignment,omitempty"`
}

func (t *AssignOrganizationRolesEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AssignOrganizationRolesEventDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *AssignOrganizationRolesEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AssignOrganizationRolesEventDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper to hold the business process configuration and organization role assignment details. For fields that accept worker, worker's position as of specified effective date will be determined when event is submitted, not when it is completed. This means if worker is being assigned a new position and worker is specified, role will be assigned to their existing position and not their new position. Workday recommends you use 'Assign_Roles_Subprocess' instead to reduce the risk of your role assignments being inadvertently reversed due to another role assignment event being processed at the same time or later-dated role assignments.
type AssignOrganizationRolesSubBusinessProcessType struct {
	BusinessSubProcessParameters     *BusinessSubProcessParametersType     `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	AssignOrganizationRolesEventData *AssignOrganizationRolesEventDataType `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Event_Data,omitempty"`
}

// Wrapper element for Assign Pay Group for Position Restriction when it is a sub business process to Create Position.
type AssignPayGroupforPositionRestrictionDataType struct {
	PayGroupReference *PayGroupObjectType `xml:"urn:com.workday/bsvc Pay_Group_Reference"`
}

// Wrapper element for the Assign Pay Group to Position Restriction sub business process.
type AssignPayGroupforPositionRestrictionsSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType             `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	AssignPayGroupData           *AssignPayGroupforPositionRestrictionDataType `xml:"urn:com.workday/bsvc Assign_Pay_Group_Data,omitempty"`
}

// Wrapper element for Assign Position Organization business process.
type AssignPositionOrganizationEventDataType struct {
	CompanyAssignmentsReference      []CompanyObjectType                    `xml:"urn:com.workday/bsvc Company_Assignments_Reference,omitempty"`
	CostCenterAssignmentsReference   []CostCenterObjectType                 `xml:"urn:com.workday/bsvc Cost_Center_Assignments_Reference,omitempty"`
	RegionAssignmentsReference       []RegionObjectType                     `xml:"urn:com.workday/bsvc Region_Assignments_Reference,omitempty"`
	CustomOrganizationAssignmentData []CustomOrganizationAssignmentDataType `xml:"urn:com.workday/bsvc Custom_Organization_Assignment_Data,omitempty"`
	FundAssignmentsReference         []FundObjectType                       `xml:"urn:com.workday/bsvc Fund_Assignments_Reference,omitempty"`
	GrantAssignmentsReference        []GrantObjectType                      `xml:"urn:com.workday/bsvc Grant_Assignments_Reference,omitempty"`
	ProgramAssignmentsReference      []ProgramObjectType                    `xml:"urn:com.workday/bsvc Program_Assignments_Reference,omitempty"`
	BusinessUnitAssignmentsReference []BusinessUnitObjectType               `xml:"urn:com.workday/bsvc Business_Unit_Assignments_Reference,omitempty"`
	GiftAssignmentsReference         []GiftObjectType                       `xml:"urn:com.workday/bsvc Gift_Assignments_Reference,omitempty"`
}

// Contains references to Job Requisitions to retrieve.
type AssignRecruitingSelfScheduleCalendarRequestReferencesType struct {
	JobRequisitionReference []JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
}

// Top level request for assigning recruiting self-schedule event calendars to job requisitions.
type AssignRecruitingSelfScheduleCalendarRequestType struct {
	JobRequisitionReference                  *JobRequisitionEnabledObjectType               `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
	AssignRecruitingSelfScheduleCalendarData *AssignRecruitingSelfScheduleCalendarsDataType `xml:"urn:com.workday/bsvc Assign_Recruiting_Self-Schedule_Calendar_Data"`
	Version                                  string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Recruiting Self-Schedule Event Calendar Assignment Data
type AssignRecruitingSelfScheduleCalendarResponseDataType struct {
	AssignRecruitingSelfScheduleCalendar []AssignRecruitingSelfScheduleCalendarType `xml:"urn:com.workday/bsvc Assign_Recruiting_Self-Schedule_Calendar,omitempty"`
}

// Response data for recruiting self-schedule event calendars assignment
type AssignRecruitingSelfScheduleCalendarResponseType struct {
	JobRequisitionReference *JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	Version                 string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Recruiting Self-Schedule Calendar Assignment
type AssignRecruitingSelfScheduleCalendarType struct {
	JobRequisitionReference                  *JobRequisitionEnabledObjectType               `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	AssignRecruitingSelfScheduleCalendarData *AssignRecruitingSelfScheduleCalendarsDataType `xml:"urn:com.workday/bsvc Assign_Recruiting_Self-Schedule_Calendar_Data,omitempty"`
}

// This holds the calendar assignment data.
type AssignRecruitingSelfScheduleCalendarsDataType struct {
	JobRequisitionID                        string                                     `xml:"urn:com.workday/bsvc Job_Requisition_ID,omitempty"`
	RecruitingSelfScheduleCalendarReference []RecruitingSelfScheduleCalendarObjectType `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AssignableRoleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AssignableRoleObjectType struct {
	ID         []AssignableRoleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Attachment WWS Data element
type AttachmentWWSDataType struct {
	ID                string              `xml:"urn:com.workday/bsvc ID,omitempty"`
	Filename          string              `xml:"urn:com.workday/bsvc Filename"`
	FileContent       []byte              `xml:"urn:com.workday/bsvc File_Content"`
	MimeTypeReference *MimeTypeObjectType `xml:"urn:com.workday/bsvc Mime_Type_Reference,omitempty"`
	Comment           string              `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

func (t *AttachmentWWSDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AttachmentWWSDataType
	var layout struct {
		*T
		FileContent *xsdBase64Binary `xml:"urn:com.workday/bsvc File_Content"`
	}
	layout.T = (*T)(t)
	layout.FileContent = (*xsdBase64Binary)(&layout.T.FileContent)
	return e.EncodeElement(layout, start)
}
func (t *AttachmentWWSDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AttachmentWWSDataType
	var overlay struct {
		*T
		FileContent *xsdBase64Binary `xml:"urn:com.workday/bsvc File_Content"`
	}
	overlay.T = (*T)(t)
	overlay.FileContent = (*xsdBase64Binary)(&overlay.T.FileContent)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type AuditedAccountingWorktagObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AuditedAccountingWorktagObjectType struct {
	ID         []AuditedAccountingWorktagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type AwardObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AwardObjectType struct {
	ID         []AwardObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for Award and Activity information.
type AwardandActivityAchievementDataType struct {
	AwardandActivityID            string                          `xml:"urn:com.workday/bsvc Award_and_Activity_ID,omitempty"`
	RemoveAwardandActivity        *bool                           `xml:"urn:com.workday/bsvc Remove_Award_and_Activity,omitempty"`
	AwardandActivityTypeReference *AwardandActivityTypeObjectType `xml:"urn:com.workday/bsvc Award_and_Activity_Type_Reference,omitempty"`
	Title                         string                          `xml:"urn:com.workday/bsvc Title,omitempty"`
	SponsorIssuer                 string                          `xml:"urn:com.workday/bsvc Sponsor_Issuer,omitempty"`
	StartDate                     *time.Time                      `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                       *time.Time                      `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	Description                   string                          `xml:"urn:com.workday/bsvc Description,omitempty"`
	URL                           string                          `xml:"urn:com.workday/bsvc URL,omitempty"`
	RelatedPositionReference      *PositionElementObjectType      `xml:"urn:com.workday/bsvc Related_Position_Reference,omitempty"`
}

func (t *AwardandActivityAchievementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AwardandActivityAchievementDataType
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
func (t *AwardandActivityAchievementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AwardandActivityAchievementDataType
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

// Wrapper element for Award and Activity information.
type AwardandActivityType struct {
	AwardandActivityReference *AwardObjectType                     `xml:"urn:com.workday/bsvc Award_and_Activity_Reference,omitempty"`
	AwardandActivityData      *AwardandActivityAchievementDataType `xml:"urn:com.workday/bsvc Award_and_Activity_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AwardandActivityTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AwardandActivityTypeObjectType struct {
	ID         []AwardandActivityTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for Background Check data.
type BackgroundCheckDataType struct {
	RecipientReference        *RoleObjectType                       `xml:"urn:com.workday/bsvc Recipient_Reference"`
	EventReference            *BackgroundCheckEventObjectType       `xml:"urn:com.workday/bsvc Event_Reference"`
	BackgroundCheckStatusData *BackgroundCheckOverallStatusDataType `xml:"urn:com.workday/bsvc Background_Check_Status_Data,omitempty"`
	PackageReferenceData      *PackageReferenceDataType             `xml:"urn:com.workday/bsvc Package_Reference_Data,omitempty"`
	TestReferenceData         []TestReferenceDataType               `xml:"urn:com.workday/bsvc Test_Reference_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BackgroundCheckEventObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BackgroundCheckEventObjectType struct {
	ID         []BackgroundCheckEventObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Details for a Background Check Event
type BackgroundCheckEventType struct {
	EventReference            *BackgroundCheckEventObjectType       `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	RecruitingSubProcess      *bool                                 `xml:"urn:com.workday/bsvc Recruiting_Sub_Process,omitempty"`
	RequesterReference        *WorkerObjectType                     `xml:"urn:com.workday/bsvc Requester_Reference,omitempty"`
	RecipientReference        *RoleObjectType                       `xml:"urn:com.workday/bsvc Recipient_Reference,omitempty"`
	RequisitionReference      *JobRequisitionEnabledObjectType      `xml:"urn:com.workday/bsvc Requisition_Reference,omitempty"`
	SubmissionDate            *time.Time                            `xml:"urn:com.workday/bsvc Submission_Date,omitempty"`
	BackgroundCheckStatusData *BackgroundCheckOverallStatusDataType `xml:"urn:com.workday/bsvc Background_Check_Status_Data,omitempty"`
	WorkflowStateReference    *UniqueIdentifierObjectType           `xml:"urn:com.workday/bsvc Workflow_State_Reference,omitempty"`
	PackageReferenceData      []PackageReferenceDataType            `xml:"urn:com.workday/bsvc Package_Reference_Data,omitempty"`
	TestReferenceData         []TestReferenceDataType               `xml:"urn:com.workday/bsvc Test_Reference_Data,omitempty"`
	DocumentFieldResultData   []DocumentFieldResultDataType         `xml:"urn:com.workday/bsvc Document_Field_Result_Data,omitempty"`
}

func (t *BackgroundCheckEventType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BackgroundCheckEventType
	var layout struct {
		*T
		SubmissionDate *xsdDateTime `xml:"urn:com.workday/bsvc Submission_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.SubmissionDate = (*xsdDateTime)(layout.T.SubmissionDate)
	return e.EncodeElement(layout, start)
}
func (t *BackgroundCheckEventType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BackgroundCheckEventType
	var overlay struct {
		*T
		SubmissionDate *xsdDateTime `xml:"urn:com.workday/bsvc Submission_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.SubmissionDate = (*xsdDateTime)(overlay.T.SubmissionDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for Background Check data.
type BackgroundCheckOverallStatusDataType struct {
	StatusDate      time.Time                        `xml:"urn:com.workday/bsvc Status_Date"`
	StatusReference *BackgroundCheckStatusObjectType `xml:"urn:com.workday/bsvc Status_Reference"`
	StatusComment   string                           `xml:"urn:com.workday/bsvc Status_Comment,omitempty"`
}

func (t *BackgroundCheckOverallStatusDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BackgroundCheckOverallStatusDataType
	var layout struct {
		*T
		StatusDate *xsdDate `xml:"urn:com.workday/bsvc Status_Date"`
	}
	layout.T = (*T)(t)
	layout.StatusDate = (*xsdDate)(&layout.T.StatusDate)
	return e.EncodeElement(layout, start)
}
func (t *BackgroundCheckOverallStatusDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BackgroundCheckOverallStatusDataType
	var overlay struct {
		*T
		StatusDate *xsdDate `xml:"urn:com.workday/bsvc Status_Date"`
	}
	overlay.T = (*T)(t)
	overlay.StatusDate = (*xsdDate)(&overlay.T.StatusDate)
	return d.DecodeElement(&overlay, &start)
}

// Element containing the Background Check Package information and contained tests.
type BackgroundCheckPackageDataType struct {
	BackgroundCheckPackage     string                           `xml:"urn:com.workday/bsvc Background_Check_Package"`
	PackageReferenceID         string                           `xml:"urn:com.workday/bsvc Package_Reference_ID,omitempty"`
	BackgroundCheckDescription string                           `xml:"urn:com.workday/bsvc Background_Check_Description,omitempty"`
	Inactive                   *bool                            `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	BackgroundCheckTest        []BackgroundCheckPackageTestType `xml:"urn:com.workday/bsvc Background_Check_Test,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BackgroundCheckPackageObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BackgroundCheckPackageObjectType struct {
	ID         []BackgroundCheckPackageObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains Background Check Package references
type BackgroundCheckPackageRequestReferencesType struct {
	BackgroundCheckPackageReference []BackgroundCheckPackageObjectType `xml:"urn:com.workday/bsvc Background_Check_Package_Reference"`
}

// Wrapper element containing Background Check Package Response Data for requested references or criteria and for requested response group.
type BackgroundCheckPackageResponseDataType struct {
	BackgroundCheckPackage []BackgroundCheckPackageType `xml:"urn:com.workday/bsvc Background_Check_Package,omitempty"`
}

// Documentation Element containing Background Check Package response grouping options
type BackgroundCheckPackageResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing the background test details.
type BackgroundCheckPackageTestDataType struct {
	BackgroundCheckTest            string `xml:"urn:com.workday/bsvc Background_Check_Test"`
	TestReference                  string `xml:"urn:com.workday/bsvc Test_Reference,omitempty"`
	BackgroundCheckTestDescription string `xml:"urn:com.workday/bsvc Background_Check_Test_Description,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BackgroundCheckPackageTestObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BackgroundCheckPackageTestObjectType struct {
	ID         []BackgroundCheckPackageTestObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing Background Check Package data and related tests.
type BackgroundCheckPackageTestType struct {
	BackgroundCheckPackageTestReferenceID string                              `xml:"urn:com.workday/bsvc Background_Check_Package_Test_Reference_ID,omitempty"`
	BackgroundCheckTestData               *BackgroundCheckPackageTestDataType `xml:"urn:com.workday/bsvc Background_Check_Test_Data"`
}

// Contains Background Check Package reference and details.
type BackgroundCheckPackageType struct {
	BackgroundCheckPackageReference *BackgroundCheckPackageObjectType `xml:"urn:com.workday/bsvc Background_Check_Package_Reference,omitempty"`
	BackgroundCheckPackageData      *BackgroundCheckPackageDataType   `xml:"urn:com.workday/bsvc Background_Check_Package_Data,omitempty"`
}

// Utilize the Request Criteria element to filter the returned Background Check Events to those Payees that meet ALL of the criteria.
type BackgroundCheckRequestCriteriaType struct {
	RequesterReference            *WorkerObjectType                  `xml:"urn:com.workday/bsvc Requester_Reference,omitempty"`
	RecipientReference            *RoleObjectType                    `xml:"urn:com.workday/bsvc Recipient_Reference,omitempty"`
	SubmissionDate                *time.Time                         `xml:"urn:com.workday/bsvc Submission_Date,omitempty"`
	WorkflowStateReference        *UniqueIdentifierObjectType        `xml:"urn:com.workday/bsvc Workflow_State_Reference,omitempty"`
	PackageSelected               *bool                              `xml:"urn:com.workday/bsvc Package_Selected,omitempty"`
	FieldAndParameterCriteriaData *FieldAndParameterCriteriaDataType `xml:"urn:com.workday/bsvc Field_And_Parameter_Criteria_Data,omitempty"`
}

func (t *BackgroundCheckRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BackgroundCheckRequestCriteriaType
	var layout struct {
		*T
		SubmissionDate *xsdDate `xml:"urn:com.workday/bsvc Submission_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.SubmissionDate = (*xsdDate)(layout.T.SubmissionDate)
	return e.EncodeElement(layout, start)
}
func (t *BackgroundCheckRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BackgroundCheckRequestCriteriaType
	var overlay struct {
		*T
		SubmissionDate *xsdDate `xml:"urn:com.workday/bsvc Submission_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.SubmissionDate = (*xsdDate)(overlay.T.SubmissionDate)
	return d.DecodeElement(&overlay, &start)
}

// Background Check Event request references.
type BackgroundCheckRequestReferencesType struct {
	EventReference []BackgroundCheckEventObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
}

// Data for Background Check Events. Used for Background Check Events that have Background Check Packages.
type BackgroundCheckResponseDataType struct {
	BackgroundCheckEvent []BackgroundCheckEventType `xml:"urn:com.workday/bsvc Background_Check_Event,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BackgroundCheckStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BackgroundCheckStatusObjectType struct {
	ID         []BackgroundCheckStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BloodTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BloodTypeObjectType struct {
	ID         []BloodTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BonusPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BonusPlanObjectType struct {
	ID         []BonusPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Container for the processing options for sub-business processes within a business process. If no options are submitted (or the options are submitted as 'false') then the sub-business process is simply initiated as if it where submitted on-line with approvals, reviews, notifications and to-do's in place. If the Initiator is an Integration System User, any validations you configured on the Initiation step are ignored.
type BusinessSubProcessParametersType struct {
	AutoComplete                  *bool                               `xml:"urn:com.workday/bsvc Auto_Complete,omitempty"`
	Skip                          *bool                               `xml:"urn:com.workday/bsvc Skip,omitempty"`
	BusinessProcessCommentData    *BusinessProcessCommentDataType     `xml:"urn:com.workday/bsvc Business_Process_Comment_Data,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type CalculatedPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CalculatedPlanObjectType struct {
	ID         []CalculatedPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains information about a Attachment to add to a Candidate.
type CandidateAttachmentDataType struct {
	CandidateAttachmentReference *CandidateAttachmentObjectType `xml:"urn:com.workday/bsvc Candidate_Attachment_Reference,omitempty"`
	AttachmentData               *AttachmentWWSDataType         `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
	DocumentCategoryReference    *DocumentCategoryAllObjectType `xml:"urn:com.workday/bsvc Document_Category_Reference"`
}

// Contains a unique identifier for an instance of an object.
type CandidateAttachmentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CandidateAttachmentObjectType struct {
	ID         []CandidateAttachmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Utilize the following criteria options to search for Candidate attachments within the Workday system.
type CandidateAttachmentRequestCriteriaType struct {
	CandidateReference []CandidateObjectType `xml:"urn:com.workday/bsvc Candidate_Reference,omitempty"`
}

// Contains reference to a candidate attachment.
type CandidateAttachmentRequestReferencesType struct {
	CandidateAttachmentReference []CandidateAttachmentObjectType `xml:"urn:com.workday/bsvc Candidate_Attachment_Reference,omitempty"`
	SkipNonExistingInstances     bool                            `xml:"urn:com.workday/bsvc Skip_Non_Existing_Instances,attr,omitempty"`
}

// Contains Candidate attachment data.
type CandidateAttachmentResponseDataType struct {
	CandidateAttachment []CandidateAttachmentWWSType `xml:"urn:com.workday/bsvc Candidate_Attachment,omitempty"`
}

// Element containing Candidate attachment response grouping options
type CandidateAttachmentResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Contains Candidate attachment data.
type CandidateAttachmentWWSType struct {
	CandidateReference      *CandidateObjectType         `xml:"urn:com.workday/bsvc Candidate_Reference,omitempty"`
	CandidateAttachmentData *CandidateAttachmentDataType `xml:"urn:com.workday/bsvc Candidate_Attachment_Data,omitempty"`
}

// Contains contact data for a candidate.
type CandidateContactDataType struct {
	CountryPhoneCodeReference *CountryPhoneCodeObjectType `xml:"urn:com.workday/bsvc Country_Phone_Code_Reference,omitempty"`
	PhoneNumber               string                      `xml:"urn:com.workday/bsvc Phone_Number,omitempty"`
	PhoneExtension            string                      `xml:"urn:com.workday/bsvc Phone_Extension,omitempty"`
	EmailAddress              string                      `xml:"urn:com.workday/bsvc Email_Address,omitempty"`
	WebsiteData               []CandidateWebsitesDataType `xml:"urn:com.workday/bsvc Website_Data,omitempty"`
	LocationData              *CandidateLocationDataType  `xml:"urn:com.workday/bsvc Location_Data,omitempty"`
}

// A Candidate and optional Job Requisition reference that should be used to create or modify a Candidate Assessment.
type CandidateCriteriaType struct {
	CandidateReference      *CandidateObjectType             `xml:"urn:com.workday/bsvc Candidate_Reference"`
	JobRequisitionReference *JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
}

// Contains Candidate data.
type CandidateDataType struct {
	CandidateID                 string                                `xml:"urn:com.workday/bsvc Candidate_ID,omitempty"`
	PreHireReference            *ApplicantObjectType                  `xml:"urn:com.workday/bsvc Pre-Hire_Reference,omitempty"`
	WorkerReference             *WorkerObjectType                     `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	NameData                    *CandidateNameDataType                `xml:"urn:com.workday/bsvc Name_Data,omitempty"`
	ContactData                 *CandidateContactDataType             `xml:"urn:com.workday/bsvc Contact_Data,omitempty"`
	SocialMediaAccountData      []CandidateSocialMediaAccountDataType `xml:"urn:com.workday/bsvc Social_Media_Account_Data,omitempty"`
	StatusData                  *CandidateStatusDataType              `xml:"urn:com.workday/bsvc Status_Data,omitempty"`
	JobApplicationData          []JobApplicationDataType              `xml:"urn:com.workday/bsvc Job_Application_Data,omitempty"`
	ProspectData                *RecruitingProspectDataType           `xml:"urn:com.workday/bsvc Prospect_Data,omitempty"`
	CandidateIdentificationData *CandidateIdentificationDataType      `xml:"urn:com.workday/bsvc Candidate_Identification_Data,omitempty"`
	LanguageReference           *UserLanguageObjectType               `xml:"urn:com.workday/bsvc Language_Reference,omitempty"`
	CandidateTagReference       []CandidateTagObjectType              `xml:"urn:com.workday/bsvc Candidate_Tag_Reference,omitempty"`
	CandidatePoolData           *CandidatePoolDataType                `xml:"urn:com.workday/bsvc Candidate_Pool_Data,omitempty"`
}

// Wrapper element for all Disability Status data for the candidate.
type CandidateDisabilityStatusDataType struct {
	DisabilityStatusReference *DisabilityStatusReferenceObjectType `xml:"urn:com.workday/bsvc Disability_Status_Reference,omitempty"`
	DisabilityReference       *DisabilityObjectType                `xml:"urn:com.workday/bsvc Disability_Reference,omitempty"`
	DisabilityStatusDate      *time.Time                           `xml:"urn:com.workday/bsvc Disability_Status_Date,omitempty"`
}

func (t *CandidateDisabilityStatusDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CandidateDisabilityStatusDataType
	var layout struct {
		*T
		DisabilityStatusDate *xsdDate `xml:"urn:com.workday/bsvc Disability_Status_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DisabilityStatusDate = (*xsdDate)(layout.T.DisabilityStatusDate)
	return e.EncodeElement(layout, start)
}
func (t *CandidateDisabilityStatusDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CandidateDisabilityStatusDataType
	var overlay struct {
		*T
		DisabilityStatusDate *xsdDate `xml:"urn:com.workday/bsvc Disability_Status_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DisabilityStatusDate = (*xsdDate)(overlay.T.DisabilityStatusDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains education details for the candidate.
type CandidateEducationDataType struct {
	SchoolReference       *SchoolObjectType       `xml:"urn:com.workday/bsvc School_Reference,omitempty"`
	SchoolName            string                  `xml:"urn:com.workday/bsvc School_Name,omitempty"`
	FirstYearAttended     float64                 `xml:"urn:com.workday/bsvc First_Year_Attended,omitempty"`
	LastYearAttended      float64                 `xml:"urn:com.workday/bsvc Last_Year_Attended,omitempty"`
	DegreeReference       *DegreeObjectType       `xml:"urn:com.workday/bsvc Degree_Reference,omitempty"`
	FieldofStudyReference *FieldofStudyObjectType `xml:"urn:com.workday/bsvc Field_of_Study_Reference,omitempty"`
	GradeAverage          string                  `xml:"urn:com.workday/bsvc Grade_Average,omitempty"`
}

// Contains the work experience for the candidate.
type CandidateExperienceDataType struct {
	CompanyReference  *JobHistoryCompanyObjectType `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	CompanyName       string                       `xml:"urn:com.workday/bsvc Company_Name,omitempty"`
	Title             string                       `xml:"urn:com.workday/bsvc Title"`
	Location          string                       `xml:"urn:com.workday/bsvc Location,omitempty"`
	StartMonth        float64                      `xml:"urn:com.workday/bsvc Start_Month,omitempty"`
	StartYear         float64                      `xml:"urn:com.workday/bsvc Start_Year"`
	EndMonth          float64                      `xml:"urn:com.workday/bsvc End_Month,omitempty"`
	EndYear           float64                      `xml:"urn:com.workday/bsvc End_Year,omitempty"`
	CurrentlyWorkHere *bool                        `xml:"urn:com.workday/bsvc Currently_Work_Here,omitempty"`
	Description       string                       `xml:"urn:com.workday/bsvc Description,omitempty"`
}

// Contains Candidate Identification Data
type CandidateIdentificationDataType struct {
	NationalID   []NationalIDType   `xml:"urn:com.workday/bsvc National_ID,omitempty"`
	GovernmentID []GovernmentIDType `xml:"urn:com.workday/bsvc Government_ID,omitempty"`
}

// Contains information about the job applications for the Candidate.
type CandidateJobApplicationDataType struct {
	JobApplicationReference *JobApplicationObjectType        `xml:"urn:com.workday/bsvc Job_Application_Reference,omitempty"`
	JobRequisitionReference *JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
}

// Contains the jobs the candidate has applied to.
type CandidateJobAppliedToDataType struct {
	JobApplicationReference       *JobApplicationObjectType                          `xml:"urn:com.workday/bsvc Job_Application_Reference,omitempty"`
	JobRequisitionReference       *JobRequisitionEnabledObjectType                   `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
	JobApplicationDate            *time.Time                                         `xml:"urn:com.workday/bsvc Job_Application_Date,omitempty"`
	StageReference                *RecruitingStageWorkdayOwnedObjectType             `xml:"urn:com.workday/bsvc Stage_Reference"`
	WorkflowStepReference         *WorkflowStepObjectType                            `xml:"urn:com.workday/bsvc Workflow_Step_Reference,omitempty"`
	DispositionReference          *RecruitingDispositionObjectType                   `xml:"urn:com.workday/bsvc Disposition_Reference,omitempty"`
	StatusTimestamp               *time.Time                                         `xml:"urn:com.workday/bsvc Status_Timestamp,omitempty"`
	SourceReference               *ApplicantSourceObjectType                         `xml:"urn:com.workday/bsvc Source_Reference,omitempty"`
	ReferredByWorkerReference     *WorkerObjectType                                  `xml:"urn:com.workday/bsvc Referred_By_Worker_Reference,omitempty"`
	AddedByWorkerReference        *ProcessmaintainedRoleObjectType                   `xml:"urn:com.workday/bsvc Added_By_Worker_Reference,omitempty"`
	PersonalInformationData       *ApplicationPersonalInformationDataType            `xml:"urn:com.workday/bsvc Personal_Information_Data,omitempty"`
	GlobalPersonalInformationData *PersonBiographicandDemographicInformationDataType `xml:"urn:com.workday/bsvc Global_Personal_Information_Data,omitempty"`
}

func (t *CandidateJobAppliedToDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CandidateJobAppliedToDataType
	var layout struct {
		*T
		JobApplicationDate *xsdDateTime `xml:"urn:com.workday/bsvc Job_Application_Date,omitempty"`
		StatusTimestamp    *xsdDateTime `xml:"urn:com.workday/bsvc Status_Timestamp,omitempty"`
	}
	layout.T = (*T)(t)
	layout.JobApplicationDate = (*xsdDateTime)(layout.T.JobApplicationDate)
	layout.StatusTimestamp = (*xsdDateTime)(layout.T.StatusTimestamp)
	return e.EncodeElement(layout, start)
}
func (t *CandidateJobAppliedToDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CandidateJobAppliedToDataType
	var overlay struct {
		*T
		JobApplicationDate *xsdDateTime `xml:"urn:com.workday/bsvc Job_Application_Date,omitempty"`
		StatusTimestamp    *xsdDateTime `xml:"urn:com.workday/bsvc Status_Timestamp,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.JobApplicationDate = (*xsdDateTime)(overlay.T.JobApplicationDate)
	overlay.StatusTimestamp = (*xsdDateTime)(overlay.T.StatusTimestamp)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for Language Achievement information
type CandidateLanguageDataType struct {
	Native          *bool                 `xml:"urn:com.workday/bsvc Native,omitempty"`
	LanguageAbility []LanguageAbilityType `xml:"urn:com.workday/bsvc Language_Ability"`
}

// Wrapper element for language profile
type CandidateLanguageSkillDataType struct {
	LanguageReference *LanguageObjectType         `xml:"urn:com.workday/bsvc Language_Reference"`
	Language          []CandidateLanguageDataType `xml:"urn:com.workday/bsvc Language"`
}

// Contains localized name data for a candidate
type CandidateLocalNameDetailDataType struct {
	FirstNameLocal  string `xml:"urn:com.workday/bsvc First_Name_-_Local,omitempty"`
	LastNameLocal   string `xml:"urn:com.workday/bsvc Last_Name_-_Local,omitempty"`
	FirstName2Local string `xml:"urn:com.workday/bsvc First_Name_2_-_Local,omitempty"`
	LastName2Local  string `xml:"urn:com.workday/bsvc Last_Name_2_-_Local,omitempty"`
}

// Contains Location data for the Candidate
type CandidateLocationDataType struct {
	CountryReference        *CountryObjectType       `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	AddressLine1            string                   `xml:"urn:com.workday/bsvc Address_Line_1,omitempty"`
	AddressLine2            string                   `xml:"urn:com.workday/bsvc Address_Line_2,omitempty"`
	AddressLine3            string                   `xml:"urn:com.workday/bsvc Address_Line_3,omitempty"`
	AddressLine4            string                   `xml:"urn:com.workday/bsvc Address_Line_4,omitempty"`
	AddressLine5            string                   `xml:"urn:com.workday/bsvc Address_Line_5,omitempty"`
	AddressLine6            string                   `xml:"urn:com.workday/bsvc Address_Line_6,omitempty"`
	AddressLine7            string                   `xml:"urn:com.workday/bsvc Address_Line_7,omitempty"`
	AddressLine8            string                   `xml:"urn:com.workday/bsvc Address_Line_8,omitempty"`
	AddressLine9            string                   `xml:"urn:com.workday/bsvc Address_Line_9,omitempty"`
	AddressLine1Local       string                   `xml:"urn:com.workday/bsvc Address_Line_1_-_Local,omitempty"`
	AddressLine2Local       string                   `xml:"urn:com.workday/bsvc Address_Line_2_-_Local,omitempty"`
	AddressLine3Local       string                   `xml:"urn:com.workday/bsvc Address_Line_3_-_Local,omitempty"`
	AddressLine4Local       string                   `xml:"urn:com.workday/bsvc Address_Line_4_-_Local,omitempty"`
	AddressLine5Local       string                   `xml:"urn:com.workday/bsvc Address_Line_5_-_Local,omitempty"`
	City                    string                   `xml:"urn:com.workday/bsvc City,omitempty"`
	CityLocal               string                   `xml:"urn:com.workday/bsvc City_-_Local,omitempty"`
	CitySubdivision1        string                   `xml:"urn:com.workday/bsvc City_Subdivision_1,omitempty"`
	CitySubdivision1Local   string                   `xml:"urn:com.workday/bsvc City_Subdivision_1_-_Local,omitempty"`
	CountryRegionReference  *CountryRegionObjectType `xml:"urn:com.workday/bsvc Country_Region_Reference,omitempty"`
	RegionSubdivision1      string                   `xml:"urn:com.workday/bsvc Region_Subdivision_1,omitempty"`
	RegionSubdivision2      string                   `xml:"urn:com.workday/bsvc Region_Subdivision_2,omitempty"`
	RegionSubdivision1Local string                   `xml:"urn:com.workday/bsvc Region_Subdivision_1_-_Local,omitempty"`
	PostalCode              string                   `xml:"urn:com.workday/bsvc Postal_Code,omitempty"`
}

// Wrapper element for each Military Service entry.
type CandidateMilitaryServiceInformationDataType struct {
	MilitaryServiceReference *MilitaryServiceReferenceObjectType `xml:"urn:com.workday/bsvc Military_Service_Reference,omitempty"`
	MilitaryServiceData      *MilitaryServiceSubDataType         `xml:"urn:com.workday/bsvc Military_Service_Data,omitempty"`
}

// Contains name data for a candidate.
type CandidateNameDataType struct {
	FirstName         string                            `xml:"urn:com.workday/bsvc First_Name,omitempty"`
	MiddleName        string                            `xml:"urn:com.workday/bsvc Middle_Name,omitempty"`
	LastName          string                            `xml:"urn:com.workday/bsvc Last_Name,omitempty"`
	SecondaryLastName string                            `xml:"urn:com.workday/bsvc Secondary_Last_Name,omitempty"`
	LocalPersonName   *CandidateLocalNameDetailDataType `xml:"urn:com.workday/bsvc Local_Person_Name,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CandidateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CandidateObjectType struct {
	ID         []CandidateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains candidate photo data.
type CandidatePhotoDataType struct {
	FileName string `xml:"urn:com.workday/bsvc File_Name,omitempty"`
	File     []byte `xml:"urn:com.workday/bsvc File"`
}

func (t *CandidatePhotoDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CandidatePhotoDataType
	var layout struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File"`
	}
	layout.T = (*T)(t)
	layout.File = (*xsdBase64Binary)(&layout.T.File)
	return e.EncodeElement(layout, start)
}
func (t *CandidatePhotoDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CandidatePhotoDataType
	var overlay struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File"`
	}
	overlay.T = (*T)(t)
	overlay.File = (*xsdBase64Binary)(&overlay.T.File)
	return d.DecodeElement(&overlay, &start)
}

// Contains Candidate Photo data.
type CandidatePhotoType struct {
	CandidateReference *CandidateObjectType    `xml:"urn:com.workday/bsvc Candidate_Reference,omitempty"`
	CandidatePhotoData *CandidatePhotoDataType `xml:"urn:com.workday/bsvc Candidate_Photo_Data,omitempty"`
}

// Contains candidate photo data.
type CandidatePhotosResponseDataType struct {
	CandidatePhoto []CandidatePhotoType `xml:"urn:com.workday/bsvc Candidate_Photo,omitempty"`
}

// Candidate pool data for the candidate.
type CandidatePoolDataType struct {
	CandidatePoolMembershipData []CandidatePoolMembershipDataType `xml:"urn:com.workday/bsvc Candidate_Pool_Membership_Data,omitempty"`
}

// Candidate pool membership and readiness status for the candidate.
type CandidatePoolMembershipDataType struct {
	StaticCandidatePoolReference *StaticCandidatePoolObjectType      `xml:"urn:com.workday/bsvc Static_Candidate_Pool_Reference"`
	ReadinessStatusReference     *CandidateReadinessStatusObjectType `xml:"urn:com.workday/bsvc Readiness_Status_Reference,omitempty"`
}

// Contains the questionnaire response for the candidate.
type CandidateQuestionnaireResponseDataType struct {
	ResponseData *QuestionnairesResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CandidateReadinessStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CandidateReadinessStatusObjectType struct {
	ID         []CandidateReadinessStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Utilize the following criteria options to search for Candidates within the Workday system. The Candidate references that are returned are those that satisfy ALL criteria included in the request. Therefore, the result set will become more limited with every criterium that is populated.
type CandidateRequestCriteriaType struct {
	FirstName                string                                  `xml:"urn:com.workday/bsvc First_Name,omitempty"`
	LastName                 string                                  `xml:"urn:com.workday/bsvc Last_Name,omitempty"`
	CandidateEmailAddress    string                                  `xml:"urn:com.workday/bsvc Candidate_Email_Address,omitempty"`
	PreHireReference         []ApplicantObjectType                   `xml:"urn:com.workday/bsvc Pre-Hire_Reference,omitempty"`
	WorkerReference          []WorkerObjectType                      `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	JobRequisitionReference  []JobRequisitionEnabledObjectType       `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	RecruitingStageReference []RecruitingStageWorkdayOwnedObjectType `xml:"urn:com.workday/bsvc Recruiting_Stage_Reference,omitempty"`
	ApplicantSourceReference []ApplicantSourceObjectType             `xml:"urn:com.workday/bsvc Applicant_Source_Reference,omitempty"`
	AppliedFrom              *time.Time                              `xml:"urn:com.workday/bsvc Applied_From,omitempty"`
	AppliedThrough           *time.Time                              `xml:"urn:com.workday/bsvc Applied_Through,omitempty"`
	InternalWorkersOnly      *bool                                   `xml:"urn:com.workday/bsvc Internal_Workers_Only,omitempty"`
	CreatedFrom              *time.Time                              `xml:"urn:com.workday/bsvc Created_From,omitempty"`
	CreatedThrough           *time.Time                              `xml:"urn:com.workday/bsvc Created_Through,omitempty"`
}

func (t *CandidateRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CandidateRequestCriteriaType
	var layout struct {
		*T
		AppliedFrom    *xsdDateTime `xml:"urn:com.workday/bsvc Applied_From,omitempty"`
		AppliedThrough *xsdDateTime `xml:"urn:com.workday/bsvc Applied_Through,omitempty"`
		CreatedFrom    *xsdDateTime `xml:"urn:com.workday/bsvc Created_From,omitempty"`
		CreatedThrough *xsdDateTime `xml:"urn:com.workday/bsvc Created_Through,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AppliedFrom = (*xsdDateTime)(layout.T.AppliedFrom)
	layout.AppliedThrough = (*xsdDateTime)(layout.T.AppliedThrough)
	layout.CreatedFrom = (*xsdDateTime)(layout.T.CreatedFrom)
	layout.CreatedThrough = (*xsdDateTime)(layout.T.CreatedThrough)
	return e.EncodeElement(layout, start)
}
func (t *CandidateRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CandidateRequestCriteriaType
	var overlay struct {
		*T
		AppliedFrom    *xsdDateTime `xml:"urn:com.workday/bsvc Applied_From,omitempty"`
		AppliedThrough *xsdDateTime `xml:"urn:com.workday/bsvc Applied_Through,omitempty"`
		CreatedFrom    *xsdDateTime `xml:"urn:com.workday/bsvc Created_From,omitempty"`
		CreatedThrough *xsdDateTime `xml:"urn:com.workday/bsvc Created_Through,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AppliedFrom = (*xsdDateTime)(overlay.T.AppliedFrom)
	overlay.AppliedThrough = (*xsdDateTime)(overlay.T.AppliedThrough)
	overlay.CreatedFrom = (*xsdDateTime)(overlay.T.CreatedFrom)
	overlay.CreatedThrough = (*xsdDateTime)(overlay.T.CreatedThrough)
	return d.DecodeElement(&overlay, &start)
}

// Contains candidate references.
type CandidateRequestReferencesType struct {
	CandidateReference       []CandidateObjectType `xml:"urn:com.workday/bsvc Candidate_Reference"`
	SkipNonExistingInstances bool                  `xml:"urn:com.workday/bsvc Skip_Non_Existing_Instances,attr,omitempty"`
}

// Contains Candidate data.
type CandidateResponseDataType struct {
	Candidate []CandidateType `xml:"urn:com.workday/bsvc Candidate,omitempty"`
}

// Element containing Candidate response grouping options
type CandidateResponseGroupType struct {
	IncludeReference      *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	ExcludeAllAttachments *bool `xml:"urn:com.workday/bsvc Exclude_All_Attachments,omitempty"`
}

// Contains the candidate resume (skills, experience, and education). Do not remove the Candidate_Resume_Data section when the Enforce Required in Web Service check box is selected for a section on a Job Application Template.
type CandidateResumeDataType struct {
	Summary                   string                                   `xml:"urn:com.workday/bsvc Summary,omitempty"`
	SkillData                 []CandidateSkillDataType                 `xml:"urn:com.workday/bsvc Skill_Data,omitempty"`
	ExperienceData            []CandidateExperienceDataType            `xml:"urn:com.workday/bsvc Experience_Data,omitempty"`
	EducationData             []CandidateEducationDataType             `xml:"urn:com.workday/bsvc Education_Data,omitempty"`
	LanguageData              []CandidateLanguageSkillDataType         `xml:"urn:com.workday/bsvc Language_Data,omitempty"`
	QuestionnaireResponseData []CandidateQuestionnaireResponseDataType `xml:"urn:com.workday/bsvc Questionnaire_Response_Data,omitempty"`
}

// Contains skill items for the candidate.
type CandidateSkillDataType struct {
	SkillReference *SkillItemTenantedObjectType `xml:"urn:com.workday/bsvc Skill_Reference,omitempty"`
	SkillName      string                       `xml:"urn:com.workday/bsvc Skill_Name,omitempty"`
}

// Contains social media account data for a candidate.
type CandidateSocialMediaAccountDataType struct {
	SocialNetworkTypeReference   *SocialNetworkMetaTypeObjectType `xml:"urn:com.workday/bsvc Social_Network_Type_Reference,omitempty"`
	SocialNetworkAccountURL      string                           `xml:"urn:com.workday/bsvc Social_Network_Account_URL,omitempty"`
	SocialNetworkAccountUserName string                           `xml:"urn:com.workday/bsvc Social_Network_Account_User_Name,omitempty"`
}

// Contains candidate statuses.
type CandidateStatusDataType struct {
	DoNotHire *bool `xml:"urn:com.workday/bsvc Do_Not_Hire,omitempty"`
	Withdrawn *bool `xml:"urn:com.workday/bsvc Withdrawn,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CandidateTagObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CandidateTagObjectType struct {
	ID         []CandidateTagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains Candidate data.
type CandidateType struct {
	CandidateReference *CandidateObjectType `xml:"urn:com.workday/bsvc Candidate_Reference,omitempty"`
	CandidateData      *CandidateDataType   `xml:"urn:com.workday/bsvc Candidate_Data,omitempty"`
}

// Add websites for the Candidate.
type CandidateWebsitesDataType struct {
	URLAddress string `xml:"urn:com.workday/bsvc URL_Address,omitempty"`
}

// Wrapper element for Certification information.
type CertificationAchievementDataType struct {
	CertificationID          string                                  `xml:"urn:com.workday/bsvc Certification_ID,omitempty"`
	RemoveCertification      *bool                                   `xml:"urn:com.workday/bsvc Remove_Certification,omitempty"`
	CertificationReference   *CertificationObjectType                `xml:"urn:com.workday/bsvc Certification_Reference,omitempty"`
	CountryReference         *CountryObjectType                      `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CertificationName        string                                  `xml:"urn:com.workday/bsvc Certification_Name,omitempty"`
	Issuer                   string                                  `xml:"urn:com.workday/bsvc Issuer,omitempty"`
	CertificationNumber      string                                  `xml:"urn:com.workday/bsvc Certification_Number,omitempty"`
	IssuedDate               *time.Time                              `xml:"urn:com.workday/bsvc Issued_Date,omitempty"`
	ExpirationDate           *time.Time                              `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	ExaminationScore         string                                  `xml:"urn:com.workday/bsvc Examination_Score,omitempty"`
	ExaminationDate          *time.Time                              `xml:"urn:com.workday/bsvc Examination_Date,omitempty"`
	SpecialtyAchievementData []SpecialtyAchievementDatawithDatesType `xml:"urn:com.workday/bsvc Specialty_Achievement_Data,omitempty"`
	WorkerDocumentData       []CertificationAttachmentDataType       `xml:"urn:com.workday/bsvc Worker_Document_Data,omitempty"`
}

func (t *CertificationAchievementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CertificationAchievementDataType
	var layout struct {
		*T
		IssuedDate      *xsdDate `xml:"urn:com.workday/bsvc Issued_Date,omitempty"`
		ExpirationDate  *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
		ExaminationDate *xsdDate `xml:"urn:com.workday/bsvc Examination_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.IssuedDate = (*xsdDate)(layout.T.IssuedDate)
	layout.ExpirationDate = (*xsdDate)(layout.T.ExpirationDate)
	layout.ExaminationDate = (*xsdDate)(layout.T.ExaminationDate)
	return e.EncodeElement(layout, start)
}
func (t *CertificationAchievementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CertificationAchievementDataType
	var overlay struct {
		*T
		IssuedDate      *xsdDate `xml:"urn:com.workday/bsvc Issued_Date,omitempty"`
		ExpirationDate  *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
		ExaminationDate *xsdDate `xml:"urn:com.workday/bsvc Examination_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.IssuedDate = (*xsdDate)(overlay.T.IssuedDate)
	overlay.ExpirationDate = (*xsdDate)(overlay.T.ExpirationDate)
	overlay.ExaminationDate = (*xsdDate)(overlay.T.ExaminationDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for Certification information.
type CertificationAchievementType struct {
	CertificationReference *CertificationSkillObjectType      `xml:"urn:com.workday/bsvc Certification_Reference,omitempty"`
	CertificationData      []CertificationAchievementDataType `xml:"urn:com.workday/bsvc Certification_Data"`
}

// Attachments for Certification.
type CertificationAttachmentDataType struct {
	FileName                  string                         `xml:"urn:com.workday/bsvc File_Name"`
	Comment                   string                         `xml:"urn:com.workday/bsvc Comment,omitempty"`
	File                      *[]byte                        `xml:"urn:com.workday/bsvc File,omitempty"`
	DocumentCategoryReference *DocumentCategoryAllObjectType `xml:"urn:com.workday/bsvc Document_Category_Reference"`
	ContentType               string                         `xml:"urn:com.workday/bsvc Content_Type,omitempty"`
}

func (t *CertificationAttachmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CertificationAttachmentDataType
	var layout struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	layout.T = (*T)(t)
	layout.File = (*xsdBase64Binary)(layout.T.File)
	return e.EncodeElement(layout, start)
}
func (t *CertificationAttachmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CertificationAttachmentDataType
	var overlay struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.File = (*xsdBase64Binary)(overlay.T.File)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type CertificationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CertificationObjectType struct {
	ID         []CertificationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the certification profile's information for a position.
type CertificationProfileforJobDataType struct {
	CountryReference             *CountryObjectType                   `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CertificationReference       *CertificationObjectType             `xml:"urn:com.workday/bsvc Certification_Reference,omitempty"`
	CertificationName            string                               `xml:"urn:com.workday/bsvc Certification_Name,omitempty"`
	CertificationIssuer          string                               `xml:"urn:com.workday/bsvc Certification_Issuer,omitempty"`
	Required                     *bool                                `xml:"urn:com.workday/bsvc Required,omitempty"`
	QualificationSourceReference *SkillQualificationEnabledObjectType `xml:"urn:com.workday/bsvc Qualification_Source_Reference,omitempty"`
	SpecialtyAchievementData     []SpecialtyAchievementDataType       `xml:"urn:com.workday/bsvc Specialty_Achievement_Data,omitempty"`
}

// Contains the position's certification information.
type CertificationProfileforJobType struct {
	CertificationProfileReference *UniqueIdentifierObjectType         `xml:"urn:com.workday/bsvc Certification_Profile_Reference,omitempty"`
	CertificationProfileData      *CertificationProfileforJobDataType `xml:"urn:com.workday/bsvc Certification_Profile_Data,omitempty"`
}

// Replacement element containing Certification Qualifications for the Job Profile
// When updating a Job Profile, all Certifications for the Job Profile will be replaced by the submitted data. If no data is submitted, then the existing Certifications are not changed.
type CertificationQualificationProfileReplacementDataType struct {
	CountryReference              *CountryObjectType             `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CertificationReference        *CertificationObjectType       `xml:"urn:com.workday/bsvc Certification_Reference,omitempty"`
	Name                          string                         `xml:"urn:com.workday/bsvc Name,omitempty"`
	Issuer                        string                         `xml:"urn:com.workday/bsvc Issuer,omitempty"`
	Required                      *bool                          `xml:"urn:com.workday/bsvc Required,omitempty"`
	SpecialtyAchievementReference []SpecialtyAchievementDataType `xml:"urn:com.workday/bsvc Specialty_Achievement_Reference,omitempty"`
}

// Wrapper element for Certification Qualifications. Allows all certification qualifications for a Job Profile or Position Restriction to be removed - or to replace all existing certification qualifications with those sent in the web service
type CertificationQualificationReplacementType struct {
	CertificationQualificationReplacementData []CertificationQualificationProfileReplacementDataType `xml:"urn:com.workday/bsvc Certification_Qualification_Replacement_Data,omitempty"`
	Delete                                    bool                                                   `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CertificationSkillObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CertificationSkillObjectType struct {
	ID         []CertificationSkillObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Top element used for Check Position Budget as a sub business process.
type CheckPositionBudgetSubBusinessProcessType struct {
	BusinessSubProcessParameters *FinancialsBusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
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

// Wrapper element for the closing of an Evergreen Requisition
type CloseEvergreenRequisitionDataType struct {
	EvergreenRequisitionReference            *EvergreenJobRequisitionObjectType        `xml:"urn:com.workday/bsvc Evergreen_Requisition_Reference"`
	CloseEvergreenRequisitionReasonReference *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Close_Evergreen_Requisition_Reason_Reference"`
	EvergreenRequisitionCloseDate            time.Time                                 `xml:"urn:com.workday/bsvc Evergreen_Requisition_Close_Date"`
	UnpostJobSubBusinessProcess              *UnpostJobSubBusinessProcessType          `xml:"urn:com.workday/bsvc Unpost_Job_Sub_Business_Process,omitempty"`
}

func (t *CloseEvergreenRequisitionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CloseEvergreenRequisitionDataType
	var layout struct {
		*T
		EvergreenRequisitionCloseDate *xsdDate `xml:"urn:com.workday/bsvc Evergreen_Requisition_Close_Date"`
	}
	layout.T = (*T)(t)
	layout.EvergreenRequisitionCloseDate = (*xsdDate)(&layout.T.EvergreenRequisitionCloseDate)
	return e.EncodeElement(layout, start)
}
func (t *CloseEvergreenRequisitionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CloseEvergreenRequisitionDataType
	var overlay struct {
		*T
		EvergreenRequisitionCloseDate *xsdDate `xml:"urn:com.workday/bsvc Evergreen_Requisition_Close_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EvergreenRequisitionCloseDate = (*xsdDate)(&overlay.T.EvergreenRequisitionCloseDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper for Close Evergreen Requisition Web Service
type CloseEvergreenRequisitionRequestType struct {
	BusinessProcessParameters     *BusinessProcessParametersType     `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	CloseEvergreenRequisitionData *CloseEvergreenRequisitionDataType `xml:"urn:com.workday/bsvc Close_Evergreen_Requisition_Data"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the event id for the Close Evergreen Job Requisition Event and the Evergreen Requisition Reference.
type CloseEvergreenRequisitionResponseType struct {
	EventReference                *UniqueIdentifierObjectType        `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	EvergreenRequisitionReference *EvergreenJobRequisitionObjectType `xml:"urn:com.workday/bsvc Evergreen_Requisition_Reference,omitempty"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the closing of a job requisition.
type CloseJobRequisitionDataType struct {
	JobRequisitionReference            *JobRequisitionObjectType                 `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
	CloseJobRequisitionReasonReference *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Close_Job_Requisition_Reason_Reference"`
	JobRequisitionCloseDate            time.Time                                 `xml:"urn:com.workday/bsvc Job_Requisition_Close_Date"`
	CloseUnfilledPositions             *bool                                     `xml:"urn:com.workday/bsvc Close_Unfilled_Positions,omitempty"`
	PositionRestrictionReference       []PositionRestrictionsObjectType          `xml:"urn:com.workday/bsvc Position_Restriction_Reference,omitempty"`
	UnpostJobSubBusinessProcess        *UnpostJobSubBusinessProcessType          `xml:"urn:com.workday/bsvc Unpost_Job_Sub_Business_Process,omitempty"`
}

func (t *CloseJobRequisitionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CloseJobRequisitionDataType
	var layout struct {
		*T
		JobRequisitionCloseDate *xsdDate `xml:"urn:com.workday/bsvc Job_Requisition_Close_Date"`
	}
	layout.T = (*T)(t)
	layout.JobRequisitionCloseDate = (*xsdDate)(&layout.T.JobRequisitionCloseDate)
	return e.EncodeElement(layout, start)
}
func (t *CloseJobRequisitionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CloseJobRequisitionDataType
	var overlay struct {
		*T
		JobRequisitionCloseDate *xsdDate `xml:"urn:com.workday/bsvc Job_Requisition_Close_Date"`
	}
	overlay.T = (*T)(t)
	overlay.JobRequisitionCloseDate = (*xsdDate)(&overlay.T.JobRequisitionCloseDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper for Close Job Requisition Web Service
type CloseJobRequisitionRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	CloseJobRequisitionData   *CloseJobRequisitionDataType   `xml:"urn:com.workday/bsvc Close_Job_Requisition_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the event id for the Close Job Requisition Event and the Job Requisition Reference.
type CloseJobRequisitionResponseType struct {
	EventReference          *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	JobRequisitionReference *JobRequisitionObjectType   `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	Version                 string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CommissionPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CommissionPlanObjectType struct {
	ID         []CommissionPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CommonBooleanEnumerationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CommonBooleanEnumerationObjectType struct {
	ID         []CommonBooleanEnumerationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CommonYesNoObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CommonYesNoObjectType struct {
	ID         []CommonYesNoObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Data element containing the compensation guidelines for a compensation change.
type CompensatableDefaultGuidelinesDataType struct {
	CompensationPackageReference      *CompensationPackageObjectType      `xml:"urn:com.workday/bsvc Compensation_Package_Reference,omitempty"`
	CompensationGradeReference        *CompensationGradeObjectType        `xml:"urn:com.workday/bsvc Compensation_Grade_Reference,omitempty"`
	CompensationGradeProfileReference *CompensationGradeProfileObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Profile_Reference,omitempty"`
	CompensationStepReference         *CompensationStepObjectType         `xml:"urn:com.workday/bsvc Compensation_Step_Reference,omitempty"`
}

// Data element containing the compensation guidelines for a compensation change.
type CompensatableGuidelinesDataType struct {
	CompensationPackageReference      *CompensationPackageObjectType      `xml:"urn:com.workday/bsvc Compensation_Package_Reference,omitempty"`
	CompensationGradeReference        *CompensationGradeObjectType        `xml:"urn:com.workday/bsvc Compensation_Grade_Reference,omitempty"`
	CompensationGradeProfileReference *CompensationGradeProfileObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Profile_Reference,omitempty"`
	CompensationStepReference         *CompensationStepObjectType         `xml:"urn:com.workday/bsvc Compensation_Step_Reference,omitempty"`
	ProgressionStartDate              *time.Time                          `xml:"urn:com.workday/bsvc Progression_Start_Date,omitempty"`
}

func (t *CompensatableGuidelinesDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CompensatableGuidelinesDataType
	var layout struct {
		*T
		ProgressionStartDate *xsdDate `xml:"urn:com.workday/bsvc Progression_Start_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ProgressionStartDate = (*xsdDate)(layout.T.ProgressionStartDate)
	return e.EncodeElement(layout, start)
}
func (t *CompensatableGuidelinesDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompensatableGuidelinesDataType
	var overlay struct {
		*T
		ProgressionStartDate *xsdDate `xml:"urn:com.workday/bsvc Progression_Start_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ProgressionStartDate = (*xsdDate)(overlay.T.ProgressionStartDate)
	return d.DecodeElement(&overlay, &start)
}

// Data element that contains the compensation default information.  If a section is to be provided, it must be provided fully as it will fully replace the current default compensation.
type CompensationDefaultDataType struct {
	PrimaryCompensationBasis              float64                                               `xml:"urn:com.workday/bsvc Primary_Compensation_Basis,omitempty"`
	PrimaryCompensationBasisAmountChange  float64                                               `xml:"urn:com.workday/bsvc Primary_Compensation_Basis_Amount_Change,omitempty"`
	PrimaryCompensationBasisPercentChange float64                                               `xml:"urn:com.workday/bsvc Primary_Compensation_Basis_Percent_Change,omitempty"`
	GuidelinesData                        *CompensatableDefaultGuidelinesDataType               `xml:"urn:com.workday/bsvc Guidelines_Data,omitempty"`
	PayPlanData                           *ProposedBasePayPlanAssignmentContainerDataType       `xml:"urn:com.workday/bsvc Pay_Plan_Data,omitempty"`
	UnitSalaryPlanData                    *ProposedSalaryUnitPlanAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Unit_Salary_Plan_Data,omitempty"`
	AllowancePlanNonUnitBasedData         *ProposedAllowancePlanAssignmentContainerDataType     `xml:"urn:com.workday/bsvc Allowance_Plan_Non-Unit_Based_Data,omitempty"`
	AllowancePlanUnitBasedData            *ProposedAllowanceUnitPlanAssignmentContainerDataType `xml:"urn:com.workday/bsvc Allowance_Plan_Unit_Based_Data,omitempty"`
	BonusPlanData                         *ProposedBonusPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Bonus_Plan_Data,omitempty"`
	MeritPlanData                         *ProposedMeritPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Merit_Plan_Data,omitempty"`
	CommissionPlanData                    *ProposedCommissionPlanAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Commission_Plan_Data,omitempty"`
	StockPlanData                         *ProposedStockPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Stock_Plan_Data,omitempty"`
	PeriodSalaryPlanData                  *ProposedPeriodSalaryPlanAssignmentContainerDataType  `xml:"urn:com.workday/bsvc Period_Salary_Plan_Data,omitempty"`
	CalculatedPlanData                    *ProposedCalculatedPlanAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Calculated_Plan_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompensationGradeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationGradeObjectType struct {
	ID         []CompensationGradeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompensationGradeProfileObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationGradeProfileObjectType struct {
	ID         []CompensationGradeProfileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompensationPackageObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationPackageObjectType struct {
	ID         []CompensationPackageObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompensationPeriodObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationPeriodObjectType struct {
	ID         []CompensationPeriodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Data element for the Propose Compensation for Hire business process.
type CompensationProposedForEmploymentDataType struct {
	PrimaryCompensationBasis   float64                                               `xml:"urn:com.workday/bsvc Primary_Compensation_Basis,omitempty"`
	CompensationGuidelinesData *CompensatableGuidelinesDataType                      `xml:"urn:com.workday/bsvc Compensation_Guidelines_Data,omitempty"`
	PayPlanData                *ProposedBasePayPlanJobAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Pay_Plan_Data,omitempty"`
	UnitSalaryPlanData         *ProposedSalaryUnitPlanAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Unit_Salary_Plan_Data,omitempty"`
	AllowancePlanData          *ProposedAllowancePlanAssignmentContainerDataType     `xml:"urn:com.workday/bsvc Allowance_Plan_Data,omitempty"`
	UnitAllowancePlanData      *ProposedAllowanceUnitPlanAssignmentContainerDataType `xml:"urn:com.workday/bsvc Unit_Allowance_Plan_Data,omitempty"`
	BonusPlanData              *ProposedBonusPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Bonus_Plan_Data,omitempty"`
	MeritPlanData              *ProposedMeritPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Merit_Plan_Data,omitempty"`
	CommissionPlanData         *ProposedCommissionPlanAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Commission_Plan_Data,omitempty"`
	StockPlanData              *ProposedStockPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Stock_Plan_Data,omitempty"`
	PeriodSalaryPlanData       *ProposedPeriodSalaryPlanAssignmentContainerDataType  `xml:"urn:com.workday/bsvc Period_Salary_Plan_Data,omitempty"`
	CalculatedPlanData         *ProposedCalculatedPlanAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Calculated_Plan_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompensationStepObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationStepObjectType struct {
	ID         []CompensationStepObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for Competency Data.
type CompetencyAchievementDatav10Type struct {
	CompetencyLevelBehaviorReference  *ProficiencyRatingBehaviorObjectType `xml:"urn:com.workday/bsvc Competency_Level_Behavior_Reference,omitempty"`
	CompetencyLevelBehaviorDescriptor string                               `xml:"urn:com.workday/bsvc Competency_Level_Behavior_Descriptor,omitempty"`
	AssessmentComment                 string                               `xml:"urn:com.workday/bsvc Assessment_Comment,omitempty"`
	AssessedOn                        *time.Time                           `xml:"urn:com.workday/bsvc Assessed_On,omitempty"`
	AssessedByPersonReference         *RoleObjectType                      `xml:"urn:com.workday/bsvc Assessed_By_Person_Reference,omitempty"`
	AssessedByWorkerDescriptor        string                               `xml:"urn:com.workday/bsvc Assessed_By_Worker_Descriptor,omitempty"`
	CompetencyReference               *CompetencyObjectType                `xml:"urn:com.workday/bsvc Competency_Reference"`
	CompetencyDescriptor              string                               `xml:"urn:com.workday/bsvc Competency_Descriptor,omitempty"`
}

func (t *CompetencyAchievementDatav10Type) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CompetencyAchievementDatav10Type
	var layout struct {
		*T
		AssessedOn *xsdDate `xml:"urn:com.workday/bsvc Assessed_On,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssessedOn = (*xsdDate)(layout.T.AssessedOn)
	return e.EncodeElement(layout, start)
}
func (t *CompetencyAchievementDatav10Type) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompetencyAchievementDatav10Type
	var overlay struct {
		*T
		AssessedOn *xsdDate `xml:"urn:com.workday/bsvc Assessed_On,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssessedOn = (*xsdDate)(overlay.T.AssessedOn)
	return d.DecodeElement(&overlay, &start)
}

// Competency rating and comments.
type CompetencyAssessmentDataType struct {
	CompetencyReference        *CompetencyObjectType        `xml:"urn:com.workday/bsvc Competency_Reference,omitempty"`
	ProficiencyRatingReference *ProficiencyRatingObjectType `xml:"urn:com.workday/bsvc Proficiency_Rating_Reference"`
	Comment                    string                       `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompetencyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompetencyObjectType struct {
	ID         []CompetencyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the competency profile's information for a position.
type CompetencyProfileforJobDataType struct {
	CompetencyReference          *CompetencyObjectType                `xml:"urn:com.workday/bsvc Competency_Reference,omitempty"`
	ProficiencyRatingReference   *ProficiencyRatingObjectType         `xml:"urn:com.workday/bsvc Proficiency_Rating_Reference,omitempty"`
	Required                     *bool                                `xml:"urn:com.workday/bsvc Required,omitempty"`
	QualificationSourceReference *SkillQualificationEnabledObjectType `xml:"urn:com.workday/bsvc Qualification_Source_Reference,omitempty"`
}

// Contains the position's competency information.
type CompetencyProfileforJobType struct {
	CompetencyProfileReference *UniqueIdentifierObjectType      `xml:"urn:com.workday/bsvc Competency_Profile_Reference,omitempty"`
	CompetencyProfileData      *CompetencyProfileforJobDataType `xml:"urn:com.workday/bsvc Competency_Profile_Data,omitempty"`
}

// Replacement element containing Competency Qualifications for the Job Profile
// When updating a Job Profile, all Competencies for the Job Profile will be replaced by the submitted data. If no data is submitted, then the existing Competencies are not changed.
type CompetencyQualificationProfileReplacementDataType struct {
	CompetencyReference        *CompetencyObjectType        `xml:"urn:com.workday/bsvc Competency_Reference"`
	ProficiencyRatingReference *ProficiencyRatingObjectType `xml:"urn:com.workday/bsvc Proficiency_Rating_Reference,omitempty"`
	Required                   *bool                        `xml:"urn:com.workday/bsvc Required,omitempty"`
}

// Wrapper element for Competency Qualifications. Allows all competency qualifications for a Job Profile or Position Restriction to be removed - or to replace all existing competency qualifications with those sent in the web service
type CompetencyQualificationReplacementType struct {
	CompetencyQualificationReplacementData []CompetencyQualificationProfileReplacementDataType `xml:"urn:com.workday/bsvc Competency_Qualification_Replacement_Data,omitempty"`
	Delete                                 bool                                                `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// All of the person's contact data (address, phone, email, instant messenger, web address).
type ContactInformationDataType struct {
	AddressData          []AddressInformationDataType          `xml:"urn:com.workday/bsvc Address_Data,omitempty"`
	PhoneData            []PhoneInformationDataType            `xml:"urn:com.workday/bsvc Phone_Data,omitempty"`
	EmailAddressData     []EmailAddressInformationDataType     `xml:"urn:com.workday/bsvc Email_Address_Data,omitempty"`
	InstantMessengerData []InstantMessengerInformationDataType `xml:"urn:com.workday/bsvc Instant_Messenger_Data,omitempty"`
	WebAddressData       []WebAddressInformationDataType       `xml:"urn:com.workday/bsvc Web_Address_Data,omitempty"`
}

// This element contains all information related to hiring/contracting a contingent worker.  This information only needs to be filled in if the worker type is for an employee.
type ContingentWorkerCostInformationDataType struct {
	RequesterReference               *WorkerObjectType                    `xml:"urn:com.workday/bsvc Requester_Reference,omitempty"`
	CompanyforPurchaseOrderReference *CompanyObjectType                   `xml:"urn:com.workday/bsvc Company_for_Purchase_Order_Reference,omitempty"`
	SupplierReference                *SupplierObjectType                  `xml:"urn:com.workday/bsvc Supplier_Reference,omitempty"`
	CurrencyReference                *CurrencyObjectType                  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	SpendCategoryReference           *SpendCategoryObjectType             `xml:"urn:com.workday/bsvc Spend_Category_Reference,omitempty"`
	PayRate                          float64                              `xml:"urn:com.workday/bsvc Pay_Rate,omitempty"`
	FrequencyReference               *FrequencyObjectType                 `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	MaximumAmount                    float64                              `xml:"urn:com.workday/bsvc Maximum_Amount,omitempty"`
	WorktagsReference                []AuditedAccountingWorktagObjectType `xml:"urn:com.workday/bsvc Worktags_Reference,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type CostingAllocationLevelObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CostingAllocationLevelObjectType struct {
	ID         []CostingAllocationLevelObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type CountryPhoneCodeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CountryPhoneCodeObjectType struct {
	ID         []CountryPhoneCodeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type CountryRegionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CountryRegionObjectType struct {
	ID         []CountryRegionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CountrySubregionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CountrySubregionObjectType struct {
	ID         []CountrySubregionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for the creation of a new evergreen job requisition.
type CreateEvergreenJobRequisitionDataType struct {
	SupervisoryOrganizationReference  *SupervisoryOrganizationObjectType               `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference,omitempty"`
	EvergreenRequisitionData          *EvergreenJobRequisitionDataforCreateandEditType `xml:"urn:com.workday/bsvc Evergreen_Requisition_Data"`
	PostJobSubProcess                 *PostJobSubBusinessProcessType                   `xml:"urn:com.workday/bsvc Post_Job_Sub_Process,omitempty"`
	AssignOrganizationRolesSubProcess *AssignOrganizationRolesSubBusinessProcessType   `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Sub_Process,omitempty"`
}

// Wrapper for Create Evergreen Requisition Web Service
type CreateEvergreenRequisitionRequestType struct {
	BusinessProcessParameters      *BusinessProcessParametersType         `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	CreateEvergreenRequisitionData *CreateEvergreenJobRequisitionDataType `xml:"urn:com.workday/bsvc Create_Evergreen_Requisition_Data"`
	Version                        string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the event id for the Create Evergreen Job Requisition Event and the Evergreen Job Requisition Reference.
type CreateEvergreenRequisitionResponseType struct {
	EventReference                *UniqueIdentifierObjectType        `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	EvergreenRequisitionReference *EvergreenJobRequisitionObjectType `xml:"urn:com.workday/bsvc Evergreen_Requisition_Reference,omitempty"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the creation of a new job requisition for a position in a supervisory organization using position management.
type CreateJobRequisitionDataType struct {
	PositionRestrictionReference             *PositionGroupObjectType                       `xml:"urn:com.workday/bsvc Position_Restriction_Reference"`
	SupervisoryOrganizationReference         *SupervisoryOrganizationObjectType             `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference"`
	CreateJobRequisitionReasonReference      *EventClassificationSubcategoryObjectType      `xml:"urn:com.workday/bsvc Create_Job_Requisition_Reason_Reference"`
	NumberofOpenings                         float64                                        `xml:"urn:com.workday/bsvc Number_of_Openings,omitempty"`
	JobRequisitionData                       *JobRequisitionDataforCreateandEditType        `xml:"urn:com.workday/bsvc Job_Requisition_Data"`
	CheckPositionBudgetSubProcess            *CheckPositionBudgetSubBusinessProcessType     `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process,omitempty"`
	PostJobSubProcess                        *PostJobSubBusinessProcessType                 `xml:"urn:com.workday/bsvc Post_Job_Sub_Process,omitempty"`
	AssignOrganizationRolesSubProcess        *AssignOrganizationRolesSubBusinessProcessType `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Sub_Process,omitempty"`
	RequestRequisitionCompensationSubProcess *RequestRequisitionCompensationSubProcessType  `xml:"urn:com.workday/bsvc Request_Requisition_Compensation_Sub_Process,omitempty"`
}

// Responds with the event id for the Create Job Requisition Event and the Job Requisition Reference.
type CreateJobRequisitionResponseType struct {
	EventReference          *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	JobRequisitionReference *JobRequisitionObjectType   `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	Version                 string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the creation of a new position opening for a supervisory organization using position management.
type CreatePositionDataType struct {
	SupervisoryOrganizationReference     *SupervisoryOrganizationObjectType                           `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference"`
	PositionRequestReasonReference       *EventClassificationSubcategoryObjectType                    `xml:"urn:com.workday/bsvc Position_Request_Reason_Reference,omitempty"`
	PositionData                         *PositionDefinitionDataType                                  `xml:"urn:com.workday/bsvc Position_Data"`
	QualificationReplacementData         *QualificationDataforPositionRestrictionorJobProfileType     `xml:"urn:com.workday/bsvc Qualification_Replacement_Data,omitempty"`
	PositionGroupRestrictionsData        *PositionGroupRestrictionDataType                            `xml:"urn:com.workday/bsvc Position_Group_Restrictions_Data,omitempty"`
	EditAssignOrganizationSubProcess     *EditAssignPositionOrganizationSubBusinessProcessType        `xml:"urn:com.workday/bsvc Edit_Assign_Organization_Sub_Process,omitempty"`
	RequestDefaultCompensationSubProcess *RequestCompensationDefaultSubBusinessProcessType            `xml:"urn:com.workday/bsvc Request_Default_Compensation_Sub_Process,omitempty"`
	AssignPayGroupSubProcess             *AssignPayGroupforPositionRestrictionsSubBusinessProcessType `xml:"urn:com.workday/bsvc Assign_Pay_Group_Sub_Process,omitempty"`
	AssignCostingAllocationSubProcess    *AssignCostingAllocationSubBusinessProcessType               `xml:"urn:com.workday/bsvc Assign_Costing_Allocation_Sub_Process,omitempty"`
}

// Wrapper for Create Open Position Web Service.
type CreatePositionRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	CreatePositionData        *CreatePositionDataType        `xml:"urn:com.workday/bsvc Create_Position_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the event id for the Position Request Event and the Position Opening Reference (if the event was completed).
type CreatePositionResponseType struct {
	EventReference    *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	PositionReference *PositionGroupObjectType    `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	Version           string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper for Create Job Requisition Web Service
type CreateRequisitionRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	CreateJobRequisitionData  *CreateJobRequisitionDataType  `xml:"urn:com.workday/bsvc Create_Job_Requisition_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Allows user to add new or delete existing custom organizations.
// If a new Unique custom organization is added - the prior org of that type will be deleted. (The old organization does not have to be submitted with delete flag true).
// If a new non-Unique custom organization is added - the existing org(s) of that type will be retained (unless they are sent with the delete flag true).
// Custom Organizations that are not being modified do not need to be submitted.
type CustomOrganizationAssignmentDataType struct {
	CustomOrganizationAssignmentReference *CustomOrganizationObjectType `xml:"urn:com.workday/bsvc Custom_Organization_Assignment_Reference"`
	Delete                                bool                          `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type DayoftheMonthObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DayoftheMonthObjectType struct {
	ID         []DayoftheMonthObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DayoftheWeekObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DayoftheWeekObjectType struct {
	ID         []DayoftheWeekObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for  the default organization assignments for the Position.
type DefaultPositionOrganizationAssignmentsDataType struct {
	CompanyAssignmentsReference            []CompanyObjectType            `xml:"urn:com.workday/bsvc Company_Assignments_Reference,omitempty"`
	CostCenterAssignmentsReference         []CostCenterObjectType         `xml:"urn:com.workday/bsvc Cost_Center_Assignments_Reference,omitempty"`
	RegionAssignmentsReference             []RegionObjectType             `xml:"urn:com.workday/bsvc Region_Assignments_Reference,omitempty"`
	CustomOrganizationAssignmentsReference []CustomOrganizationObjectType `xml:"urn:com.workday/bsvc Custom_Organization_Assignments_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DegreeCompletedObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DegreeCompletedObjectType struct {
	ID         []DegreeCompletedObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DegreeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DegreeObjectType struct {
	ID         []DegreeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DifficultytoFillObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DifficultytoFillObjectType struct {
	ID         []DifficultytoFillObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DisabilityCertificationAuthorityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DisabilityCertificationAuthorityObjectType struct {
	ID         []DisabilityCertificationAuthorityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DisabilityCertificationBasisObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DisabilityCertificationBasisObjectType struct {
	ID         []DisabilityCertificationBasisObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type DisabilityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DisabilityObjectType struct {
	ID         []DisabilityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type DocumentCategoryAllObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DocumentCategoryAllObjectType struct {
	ID         []DocumentCategoryAllObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Result of the evaluation of an External Field based on a contextual instance.
type DocumentFieldResultDataType struct {
	FieldReference *IntegrationDocumentFieldObjectType `xml:"urn:com.workday/bsvc Field_Reference,omitempty"`
	Value          string                              `xml:"urn:com.workday/bsvc Value,omitempty"`
}

// Container for the processing options for a dynamic business process. If no options are submitted, the process does not move to the next stage or disposition.
type DynamicBusinessProcessParametersType struct {
	NextStepReference        *WorkflowStepObjectType          `xml:"urn:com.workday/bsvc Next_Step_Reference"`
	DispositionStepReference *WorkflowStepObjectType          `xml:"urn:com.workday/bsvc Disposition_Step_Reference"`
	CommentData              []BusinessProcessCommentDataType `xml:"urn:com.workday/bsvc Comment_Data,omitempty"`
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

// Wrapper element for the Assign Organizations to Position sub business process.
// If this element is submitted (Auto or Manual) as part of a Staffing event and the Position_Organization_Assignment  wrapper is not submitted, then the position's organization assignments will default from the supervisory organization.
type EditAssignPositionOrganizationSubBusinessProcessType struct {
	BusinessSubProcessParameters        *BusinessSubProcessParametersType        `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	PositionOrganizationAssignmentsData *AssignPositionOrganizationEventDataType `xml:"urn:com.workday/bsvc Position_Organization_Assignments_Data,omitempty"`
}

// Wrapper for Edit Evergreen Requisition, Data, and Sub Processes
type EditEvergreenRequisitionDataType struct {
	EvergreenRequisitionReference     *EvergreenJobRequisitionObjectType             `xml:"urn:com.workday/bsvc Evergreen_Requisition_Reference"`
	EditEvergreenRequisitionEventData *EditEvergreenRequisitionEventDataType         `xml:"urn:com.workday/bsvc Edit_Evergreen_Requisition_Event_Data,omitempty"`
	UnpostJobSubBusinessProcess       *UnpostJobSubBusinessProcessType               `xml:"urn:com.workday/bsvc Unpost_Job_Sub_Business_Process,omitempty"`
	PostJobSubBusinessProcess         *PostJobSubBusinessProcessType                 `xml:"urn:com.workday/bsvc Post_Job_Sub_Business_Process,omitempty"`
	AssignOrganizationRolesSubProcess *AssignOrganizationRolesSubBusinessProcessType `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Sub_Process,omitempty"`
}

// Wrapper element for reference to evergreen job requisition and evergreen job requisition data element.
type EditEvergreenRequisitionEventDataType struct {
	SupervisoryOrganizationReference        *SupervisoryOrganizationObjectType                `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference,omitempty"`
	EditEvergreenRequisitionReasonReference *EventClassificationSubcategoryObjectType         `xml:"urn:com.workday/bsvc Edit_Evergreen_Requisition_Reason_Reference"`
	EvergreenRequisitionData                []EvergreenJobRequisitionDataforCreateandEditType `xml:"urn:com.workday/bsvc Evergreen_Requisition_Data"`
}

// Wrapper for Edit Evergreen Requisition Request
type EditEvergreenRequisitionRequestType struct {
	BusinessProcessParameters    *BusinessProcessParametersType    `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EditEvergreenRequisitionData *EditEvergreenRequisitionDataType `xml:"urn:com.workday/bsvc Edit_Evergreen_Requisition_Data"`
	Version                      string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top-level response element for Edit Evergreen Requisition operation.
type EditEvergreenRequisitionResponseType struct {
	JobRequisitionReference *JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	EventReference          *UniqueIdentifierObjectType      `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version                 string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Allows updating of effective-dated custom objects for job requisitions.  When updating custom objects that allow multiple instances, data for existing instances must appear in the request, otherwise it will be removed.
type EditJobRequisitionAdditionalDataRequestType struct {
	BusinessProcessParameters      *BusinessProcessParametersType      `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	JobRequisitionCustomObjectData *JobRequisitionCustomObjectDataType `xml:"urn:com.workday/bsvc Job_Requisition_Custom_Object_Data"`
	Version                        string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Job Requisition effective-dated custom object data.
type EditJobRequisitionAdditionalDataResponseType struct {
	EventReference          *UniqueIdentifierObjectType                  `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	JobRequisitionReference *JobRequisitionEnabledObjectType             `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	EffectiveDate           *time.Time                                   `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	AdditionalData          []EffectiveDatedWebServiceAdditionalDataType `xml:"urn:com.workday/bsvc Additional_Data,omitempty"`
	Version                 string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

func (t *EditJobRequisitionAdditionalDataResponseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EditJobRequisitionAdditionalDataResponseType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EditJobRequisitionAdditionalDataResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EditJobRequisitionAdditionalDataResponseType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper for Edit Job Requisition Reason, Data, and Sub Processes
type EditJobRequisitionDataType struct {
	JobRequisitionReference                  *JobRequisitionObjectType                      `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
	EditJobRequisitionEventData              *EditJobRequisitionEventDataType               `xml:"urn:com.workday/bsvc Edit_Job_Requisition_Event_Data,omitempty"`
	CheckPositionBudgetSubProcess            *CheckPositionBudgetSubBusinessProcessType     `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process,omitempty"`
	UnpostJobSubBusinessProcess              *UnpostJobSubBusinessProcessType               `xml:"urn:com.workday/bsvc Unpost_Job_Sub_Business_Process,omitempty"`
	PostJobSubBusinessProcess                *PostJobSubBusinessProcessType                 `xml:"urn:com.workday/bsvc Post_Job_Sub_Business_Process,omitempty"`
	AssignOrganizationRolesSubProcess        *AssignOrganizationRolesSubBusinessProcessType `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Sub_Process,omitempty"`
	RequestRequisitionCompensationSubProcess *RequestRequisitionCompensationSubProcessType  `xml:"urn:com.workday/bsvc Request_Requisition_Compensation_Sub_Process,omitempty"`
}

// Wrapper element for reference to job requisition and job requisition data element.
type EditJobRequisitionEventDataType struct {
	NumberofOpenings                  float64                                   `xml:"urn:com.workday/bsvc Number_of_Openings,omitempty"`
	EditJobRequisitionReasonReference *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Edit_Job_Requisition_Reason_Reference"`
	JobRequisitionData                []JobRequisitionDataforCreateandEditType  `xml:"urn:com.workday/bsvc Job_Requisition_Data"`
}

// Wrapper for Edit Job Requisition Request
type EditJobRequisitionRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EditJobRequisitionData    *EditJobRequisitionDataType    `xml:"urn:com.workday/bsvc Edit_Job_Requisition_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top-level response element for Edit Job Requisition operation.
type EditJobRequisitionResponseType struct {
	JobRequisitionReference *JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	EventReference          *UniqueIdentifierObjectType      `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version                 string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper for the details of the Edit Position Restriction transaction.
type EditPositionRestrictionEventDataType struct {
	PositionEditReasonReference   *EventClassificationSubcategoryObjectType                `xml:"urn:com.workday/bsvc Position_Edit_Reason_Reference,omitempty"`
	PositionData                  *PositionDefinitionDataType                              `xml:"urn:com.workday/bsvc Position_Data,omitempty"`
	PositionGroupRestrictionsData *PositionGroupRestrictionDataType                        `xml:"urn:com.workday/bsvc Position_Group_Restrictions_Data,omitempty"`
	QualificationReplacementData  *QualificationDataforPositionRestrictionorJobProfileType `xml:"urn:com.workday/bsvc Qualification_Replacement_Data,omitempty"`
}

// Allows updating of effective-dated custom objects for position restrictions.  When updating custom objects that allow multiple instances, data for existing instances must appear in the request, otherwise it will be removed.
type EditPositionRestrictionsAdditionalDataRequestType struct {
	BusinessProcessParameters            *BusinessProcessParametersType            `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	PositionRestrictionsCustomObjectData *PositionRestrictionsCustomObjectDataType `xml:"urn:com.workday/bsvc Position_Restrictions_Custom_Object_Data"`
	Version                              string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The additional data for the position restriction.
type EditPositionRestrictionsAdditionalDataResponseType struct {
	EventReference                *UniqueIdentifierObjectType                  `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	PositionRestrictionsReference *PositionGroupObjectType                     `xml:"urn:com.workday/bsvc Position_Restrictions_Reference,omitempty"`
	EffectiveDate                 *time.Time                                   `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	AdditionalData                []EffectiveDatedWebServiceAdditionalDataType `xml:"urn:com.workday/bsvc Additional_Data,omitempty"`
	Version                       string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

func (t *EditPositionRestrictionsAdditionalDataResponseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EditPositionRestrictionsAdditionalDataResponseType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EditPositionRestrictionsAdditionalDataResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EditPositionRestrictionsAdditionalDataResponseType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for the editing of a position restriction.
type EditPositionRestrictionsDataType struct {
	PositionReference                    *PositionRestrictionsObjectType                              `xml:"urn:com.workday/bsvc Position_Reference"`
	EditPositionRestrictionEventData     *EditPositionRestrictionEventDataType                        `xml:"urn:com.workday/bsvc Edit_Position_Restriction_Event_Data,omitempty"`
	EditAssignOrganizationSubProcess     *EditAssignPositionOrganizationSubBusinessProcessType        `xml:"urn:com.workday/bsvc Edit_Assign_Organization_Sub_Process,omitempty"`
	RequestDefaultCompensationSubProcess *RequestCompensationDefaultSubBusinessProcessType            `xml:"urn:com.workday/bsvc Request_Default_Compensation_Sub_Process,omitempty"`
	AssignPayGroupSubProcess             *AssignPayGroupforPositionRestrictionsSubBusinessProcessType `xml:"urn:com.workday/bsvc Assign_Pay_Group_Sub_Process,omitempty"`
}

// Wrapper element for Edit Position Restrictions Web Service.
type EditPositionRestrictionsRequestType struct {
	BusinessProcessParameters    *BusinessProcessParametersType    `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EditPositionRestrictionsData *EditPositionRestrictionsDataType `xml:"urn:com.workday/bsvc Edit_Position_Restrictions_Data"`
	Version                      string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the event ID of the Position Restrictions Edit Event, the Position Restriction being edited and the workflow state of the event.
type EditPositonRestrictionResponseType struct {
	EventReference    *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	PositionReference *PositionGroupObjectType    `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	Version           string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for Education information.
type EducationAchievementDataType struct {
	EducationID               string                     `xml:"urn:com.workday/bsvc Education_ID,omitempty"`
	RemoveEducation           *bool                      `xml:"urn:com.workday/bsvc Remove_Education,omitempty"`
	CountryReference          *CountryObjectType         `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	SchoolReference           *SchoolObjectType          `xml:"urn:com.workday/bsvc School_Reference,omitempty"`
	SchoolName                string                     `xml:"urn:com.workday/bsvc School_Name,omitempty"`
	SchoolTypeReference       *SchoolTypeObjectType      `xml:"urn:com.workday/bsvc School_Type_Reference,omitempty"`
	Location                  string                     `xml:"urn:com.workday/bsvc Location,omitempty"`
	DegreeReference           *DegreeObjectType          `xml:"urn:com.workday/bsvc Degree_Reference,omitempty"`
	DegreeCompletedReference  *DegreeCompletedObjectType `xml:"urn:com.workday/bsvc Degree_Completed_Reference,omitempty"`
	DateDegreeReceived        *time.Time                 `xml:"urn:com.workday/bsvc Date_Degree_Received,omitempty"`
	FieldOfStudyReference     *FieldofStudyObjectType    `xml:"urn:com.workday/bsvc Field_Of_Study_Reference,omitempty"`
	GradeAverage              string                     `xml:"urn:com.workday/bsvc Grade_Average,omitempty"`
	FirstYearAttended         *time.Time                 `xml:"urn:com.workday/bsvc First_Year_Attended,omitempty"`
	LastYearAttended          *time.Time                 `xml:"urn:com.workday/bsvc Last_Year_Attended,omitempty"`
	IsHighestLevelofEducation *bool                      `xml:"urn:com.workday/bsvc Is_Highest_Level_of_Education,omitempty"`
	FirstDayAttended          *time.Time                 `xml:"urn:com.workday/bsvc First_Day_Attended,omitempty"`
	LastDayAttended           *time.Time                 `xml:"urn:com.workday/bsvc Last_Day_Attended,omitempty"`
}

func (t *EducationAchievementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EducationAchievementDataType
	var layout struct {
		*T
		DateDegreeReceived *xsdDate `xml:"urn:com.workday/bsvc Date_Degree_Received,omitempty"`
		FirstYearAttended  *xsdDate `xml:"urn:com.workday/bsvc First_Year_Attended,omitempty"`
		LastYearAttended   *xsdDate `xml:"urn:com.workday/bsvc Last_Year_Attended,omitempty"`
		FirstDayAttended   *xsdDate `xml:"urn:com.workday/bsvc First_Day_Attended,omitempty"`
		LastDayAttended    *xsdDate `xml:"urn:com.workday/bsvc Last_Day_Attended,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateDegreeReceived = (*xsdDate)(layout.T.DateDegreeReceived)
	layout.FirstYearAttended = (*xsdDate)(layout.T.FirstYearAttended)
	layout.LastYearAttended = (*xsdDate)(layout.T.LastYearAttended)
	layout.FirstDayAttended = (*xsdDate)(layout.T.FirstDayAttended)
	layout.LastDayAttended = (*xsdDate)(layout.T.LastDayAttended)
	return e.EncodeElement(layout, start)
}
func (t *EducationAchievementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EducationAchievementDataType
	var overlay struct {
		*T
		DateDegreeReceived *xsdDate `xml:"urn:com.workday/bsvc Date_Degree_Received,omitempty"`
		FirstYearAttended  *xsdDate `xml:"urn:com.workday/bsvc First_Year_Attended,omitempty"`
		LastYearAttended   *xsdDate `xml:"urn:com.workday/bsvc Last_Year_Attended,omitempty"`
		FirstDayAttended   *xsdDate `xml:"urn:com.workday/bsvc First_Day_Attended,omitempty"`
		LastDayAttended    *xsdDate `xml:"urn:com.workday/bsvc Last_Day_Attended,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateDegreeReceived = (*xsdDate)(overlay.T.DateDegreeReceived)
	overlay.FirstYearAttended = (*xsdDate)(overlay.T.FirstYearAttended)
	overlay.LastYearAttended = (*xsdDate)(overlay.T.LastYearAttended)
	overlay.FirstDayAttended = (*xsdDate)(overlay.T.FirstDayAttended)
	overlay.LastDayAttended = (*xsdDate)(overlay.T.LastDayAttended)
	return d.DecodeElement(&overlay, &start)
}

// Contains the education profile's information for a position.
type EducationProfileforJobDataType struct {
	DegreeReference              *DegreeObjectType                    `xml:"urn:com.workday/bsvc Degree_Reference"`
	FieldOfStudyReference        *FieldofStudyObjectType              `xml:"urn:com.workday/bsvc Field_Of_Study_Reference,omitempty"`
	Required                     *bool                                `xml:"urn:com.workday/bsvc Required,omitempty"`
	QualificationSourceReference *SkillQualificationEnabledObjectType `xml:"urn:com.workday/bsvc Qualification_Source_Reference,omitempty"`
}

// Contains the position's education information.
type EducationProfileforJobType struct {
	EducationProfileReference *UniqueIdentifierObjectType     `xml:"urn:com.workday/bsvc Education_Profile_Reference,omitempty"`
	EducationProfileData      *EducationProfileforJobDataType `xml:"urn:com.workday/bsvc Education_Profile_Data,omitempty"`
}

// Replacement element containing Education Qualifications for the Job Profile
// When updating a Job Profile, all Education Qualifications for the Job Profile will be replaced by the submitted data. If no data is submitted, then the existing Education Qualifications are not changed.
type EducationQualificationProfileReplacementDataType struct {
	DegreeReference       *DegreeObjectType       `xml:"urn:com.workday/bsvc Degree_Reference,omitempty"`
	FieldOfStudyReference *FieldofStudyObjectType `xml:"urn:com.workday/bsvc Field_Of_Study_Reference,omitempty"`
	Required              *bool                   `xml:"urn:com.workday/bsvc Required,omitempty"`
}

// Wrapper element for Education Qualifications. Allows all education qualifications for a Job Profile or Position Restriction to be removed - or to replace all existing education qualifications with those sent in the web service.
type EducationQualificationReplacementType struct {
	EducationQualificationReplacementData []EducationQualificationProfileReplacementDataType `xml:"urn:com.workday/bsvc Education_Qualification_Replacement_Data,omitempty"`
	Delete                                bool                                               `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EducationSkillObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EducationSkillObjectType struct {
	ID         []EducationSkillObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for Education information.
type EducationType struct {
	EducationReference *EducationSkillObjectType      `xml:"urn:com.workday/bsvc Education_Reference,omitempty"`
	EducationData      []EducationAchievementDataType `xml:"urn:com.workday/bsvc Education_Data,omitempty"`
}

// Element contains effective and updated date/time data.
type EffectiveAndUpdatedDateTimeDataType struct {
	UpdatedFrom      *time.Time `xml:"urn:com.workday/bsvc Updated_From,omitempty"`
	UpdatedThrough   *time.Time `xml:"urn:com.workday/bsvc Updated_Through,omitempty"`
	EffectiveFrom    *time.Time `xml:"urn:com.workday/bsvc Effective_From,omitempty"`
	EffectiveThrough *time.Time `xml:"urn:com.workday/bsvc Effective_Through,omitempty"`
}

func (t *EffectiveAndUpdatedDateTimeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EffectiveAndUpdatedDateTimeDataType
	var layout struct {
		*T
		UpdatedFrom      *xsdDateTime `xml:"urn:com.workday/bsvc Updated_From,omitempty"`
		UpdatedThrough   *xsdDateTime `xml:"urn:com.workday/bsvc Updated_Through,omitempty"`
		EffectiveFrom    *xsdDateTime `xml:"urn:com.workday/bsvc Effective_From,omitempty"`
		EffectiveThrough *xsdDateTime `xml:"urn:com.workday/bsvc Effective_Through,omitempty"`
	}
	layout.T = (*T)(t)
	layout.UpdatedFrom = (*xsdDateTime)(layout.T.UpdatedFrom)
	layout.UpdatedThrough = (*xsdDateTime)(layout.T.UpdatedThrough)
	layout.EffectiveFrom = (*xsdDateTime)(layout.T.EffectiveFrom)
	layout.EffectiveThrough = (*xsdDateTime)(layout.T.EffectiveThrough)
	return e.EncodeElement(layout, start)
}
func (t *EffectiveAndUpdatedDateTimeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EffectiveAndUpdatedDateTimeDataType
	var overlay struct {
		*T
		UpdatedFrom      *xsdDateTime `xml:"urn:com.workday/bsvc Updated_From,omitempty"`
		UpdatedThrough   *xsdDateTime `xml:"urn:com.workday/bsvc Updated_Through,omitempty"`
		EffectiveFrom    *xsdDateTime `xml:"urn:com.workday/bsvc Effective_From,omitempty"`
		EffectiveThrough *xsdDateTime `xml:"urn:com.workday/bsvc Effective_Through,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.UpdatedFrom = (*xsdDateTime)(overlay.T.UpdatedFrom)
	overlay.UpdatedThrough = (*xsdDateTime)(overlay.T.UpdatedThrough)
	overlay.EffectiveFrom = (*xsdDateTime)(overlay.T.EffectiveFrom)
	overlay.EffectiveThrough = (*xsdDateTime)(overlay.T.EffectiveThrough)
	return d.DecodeElement(&overlay, &start)
}

// Web Service Additional Data
type EffectiveDatedWebServiceAdditionalDataType []string

func (a EffectiveDatedWebServiceAdditionalDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ArrayType string   `xml:"http://schemas.xmlsoap.org/wsdl/ arrayType,attr"`
		Items     []string `xml:" item"`
	}
	output.Items = []string(a)
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{"", "xmlns:ns1"}, Value: "http://www.w3.org/2001/XMLSchema"})
	output.ArrayType = "ns1:anyType[]"
	return e.EncodeElement(&output, start)
}
func (a *EffectiveDatedWebServiceAdditionalDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var tok xml.Token
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if tok, ok := tok.(xml.StartElement); ok {
			var item string
			if err = d.DecodeElement(&item, &tok); err == nil {
				*a = append(*a, item)
			}
		}
		if _, ok := tok.(xml.EndElement); ok {
			break
		}
	}
	return err
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
type EthnicityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EthnicityObjectType struct {
	ID         []EthnicityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type EventClassificationSubcategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EventClassificationSubcategoryObjectType struct {
	ID         []EventClassificationSubcategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Wrapper Element for an Evergreen Requisition
type EvergreenJobRequisitionData2Type struct {
	JobRequisitionStatusReference    *JobRequisitionStatusObjectType                 `xml:"urn:com.workday/bsvc Job_Requisition_Status_Reference,omitempty"`
	LinkedJobRequisitionsReference   []JobRequisitionObjectType                      `xml:"urn:com.workday/bsvc Linked_Job_Requisitions_Reference,omitempty"`
	JobRequisitionDetailData         *JobRequisitionDefinitionDataType               `xml:"urn:com.workday/bsvc Job_Requisition_Detail_Data,omitempty"`
	HiringRequirementData            *JobRequisitionRestrictionsDataType             `xml:"urn:com.workday/bsvc Hiring_Requirement_Data,omitempty"`
	JobRequisitionAttachments        *JobRequisitionAttachmentsType                  `xml:"urn:com.workday/bsvc Job_Requisition_Attachments,omitempty"`
	QualificationData                *QualificationsfromPositionRestrictionsDataType `xml:"urn:com.workday/bsvc Qualification_Data,omitempty"`
	QuestionnaireReference           *JobRequisitionQuestionnaireDataType            `xml:"urn:com.workday/bsvc Questionnaire_Reference,omitempty"`
	AssessmentData                   *JobRequisitionAssessmentDataType               `xml:"urn:com.workday/bsvc Assessment_Data,omitempty"`
	SupervisoryOrganizationReference *SupervisoryOrganizationObjectType              `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference,omitempty"`
	OrganizationData                 []JobRequisitionOrganizationsDataType           `xml:"urn:com.workday/bsvc Organization_Data,omitempty"`
	RoleAssignmentData               []RoleAssignmentEffectiveDataType               `xml:"urn:com.workday/bsvc Role_Assignment_Data,omitempty"`
}

// Wrapper for Job Requisition data common to create and edit requests.
type EvergreenJobRequisitionDataforCreateandEditType struct {
	JobRequisitionID                      string                                                          `xml:"urn:com.workday/bsvc Job_Requisition_ID,omitempty"`
	JobPostingTitle                       string                                                          `xml:"urn:com.workday/bsvc Job_Posting_Title,omitempty"`
	RecruitingInstructionReference        *RecruitingInstructionObjectType                                `xml:"urn:com.workday/bsvc Recruiting_Instruction_Reference,omitempty"`
	JobDescriptionSummary                 string                                                          `xml:"urn:com.workday/bsvc Job_Description_Summary,omitempty"`
	JobDescription                        string                                                          `xml:"urn:com.workday/bsvc Job_Description,omitempty"`
	AdditionalJobDescription              string                                                          `xml:"urn:com.workday/bsvc Additional_Job_Description,omitempty"`
	Justification                         string                                                          `xml:"urn:com.workday/bsvc Justification,omitempty"`
	JobRequisitionAttachments             []JobRequisitionAttachmentsType                                 `xml:"urn:com.workday/bsvc Job_Requisition_Attachments,omitempty"`
	RecruitingStartDate                   *time.Time                                                      `xml:"urn:com.workday/bsvc Recruiting_Start_Date,omitempty"`
	TargetHireDate                        *time.Time                                                      `xml:"urn:com.workday/bsvc Target_Hire_Date,omitempty"`
	TargetEndDate                         *time.Time                                                      `xml:"urn:com.workday/bsvc Target_End_Date,omitempty"`
	JobProfileReference                   *JobProfileObjectType                                           `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	AdditionalJobProfilesReference        []JobProfileObjectType                                          `xml:"urn:com.workday/bsvc Additional_Job_Profiles_Reference,omitempty"`
	WorkerTypeReference                   *WorkerTypeObjectType                                           `xml:"urn:com.workday/bsvc Worker_Type_Reference,omitempty"`
	WorkerSubTypeReference                *PositionWorkerTypeObjectType                                   `xml:"urn:com.workday/bsvc Worker_Sub-Type_Reference,omitempty"`
	PrimaryLocationReference              *LocationObjectType                                             `xml:"urn:com.workday/bsvc Primary_Location_Reference,omitempty"`
	PrimaryJobPostingLocationReference    *LocationObjectType                                             `xml:"urn:com.workday/bsvc Primary_Job_Posting_Location_Reference,omitempty"`
	AdditionalLocationsReference          []LocationObjectType                                            `xml:"urn:com.workday/bsvc Additional_Locations_Reference,omitempty"`
	AdditionalJobPostingLocationReference []LocationObjectType                                            `xml:"urn:com.workday/bsvc Additional_Job_Posting_Location_Reference,omitempty"`
	PositionTimeTypeReference             *PositionTimeTypeObjectType                                     `xml:"urn:com.workday/bsvc Position_Time_Type_Reference,omitempty"`
	ScheduledWeeklyHours                  float64                                                         `xml:"urn:com.workday/bsvc Scheduled_Weekly_Hours,omitempty"`
	WorkShiftReference                    *WorkShiftObjectType                                            `xml:"urn:com.workday/bsvc Work_Shift_Reference,omitempty"`
	SpotlightJob                          *bool                                                           `xml:"urn:com.workday/bsvc Spotlight_Job,omitempty"`
	JobRequisitionLinkstoAddReference     []JobRequisitionObjectType                                      `xml:"urn:com.workday/bsvc Job_Requisition_Links_to_Add_Reference,omitempty"`
	JobRequisitionLinkstoRemoveReference  []JobRequisitionObjectType                                      `xml:"urn:com.workday/bsvc Job_Requisition_Links_to_Remove_Reference,omitempty"`
	QuestionnaireData                     *JobRequisitionQuestionnaireDataType                            `xml:"urn:com.workday/bsvc Questionnaire_Data,omitempty"`
	AssessmentData                        *JobRequisitionAssessmentDataType                               `xml:"urn:com.workday/bsvc Assessment_Data,omitempty"`
	QualificationReplacementData          *QualificationDataforPositionRestrictionorJobProfileType        `xml:"urn:com.workday/bsvc Qualification_Replacement_Data,omitempty"`
	OrganizationData                      *EvergreenJobRequisitionOrganizationsDataforBusinessProcessType `xml:"urn:com.workday/bsvc Organization_Data,omitempty"`
	UseUpdatedTemplateVersion             *bool                                                           `xml:"urn:com.workday/bsvc Use_Updated_Template_Version,omitempty"`
	JobApplicationTemplateReference       *JobApplicationTemplateObjectType                               `xml:"urn:com.workday/bsvc Job_Application_Template_Reference,omitempty"`
}

func (t *EvergreenJobRequisitionDataforCreateandEditType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EvergreenJobRequisitionDataforCreateandEditType
	var layout struct {
		*T
		RecruitingStartDate *xsdDate `xml:"urn:com.workday/bsvc Recruiting_Start_Date,omitempty"`
		TargetHireDate      *xsdDate `xml:"urn:com.workday/bsvc Target_Hire_Date,omitempty"`
		TargetEndDate       *xsdDate `xml:"urn:com.workday/bsvc Target_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.RecruitingStartDate = (*xsdDate)(layout.T.RecruitingStartDate)
	layout.TargetHireDate = (*xsdDate)(layout.T.TargetHireDate)
	layout.TargetEndDate = (*xsdDate)(layout.T.TargetEndDate)
	return e.EncodeElement(layout, start)
}
func (t *EvergreenJobRequisitionDataforCreateandEditType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EvergreenJobRequisitionDataforCreateandEditType
	var overlay struct {
		*T
		RecruitingStartDate *xsdDate `xml:"urn:com.workday/bsvc Recruiting_Start_Date,omitempty"`
		TargetHireDate      *xsdDate `xml:"urn:com.workday/bsvc Target_Hire_Date,omitempty"`
		TargetEndDate       *xsdDate `xml:"urn:com.workday/bsvc Target_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.RecruitingStartDate = (*xsdDate)(overlay.T.RecruitingStartDate)
	overlay.TargetHireDate = (*xsdDate)(overlay.T.TargetHireDate)
	overlay.TargetEndDate = (*xsdDate)(overlay.T.TargetEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type EvergreenJobRequisitionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EvergreenJobRequisitionObjectType struct {
	ID         []EvergreenJobRequisitionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Containing element for the cost center and custom organizations related to the job requisition.
type EvergreenJobRequisitionOrganizationsDataforBusinessProcessType struct {
	OrganizationAssignmentsforJobRequisitionData *OrganizationAssignmentsforJobRequisitionDataType `xml:"urn:com.workday/bsvc Organization_Assignments_for_Job_Requisition_Data,omitempty"`
	ReplaceAll                                   bool                                              `xml:"urn:com.workday/bsvc Replace_All,attr,omitempty"`
	Delete                                       bool                                              `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Utilize the Request Criteria element to filter instance(s) of Evergreen Requisitions by status or supervisory org.
type EvergreenJobRequisitionRequestCriteriaType struct {
	TransactionLogCriteriaData       []TransactionLogCriteriaType        `xml:"urn:com.workday/bsvc Transaction_Log_Criteria_Data,omitempty"`
	JobRequisitionStatusReference    []JobRequisitionStatusObjectType    `xml:"urn:com.workday/bsvc Job_Requisition_Status_Reference,omitempty"`
	SupervisoryOrganizationReference []SupervisoryOrganizationObjectType `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference,omitempty"`
	PrimaryLocationReference         []LocationObjectType                `xml:"urn:com.workday/bsvc Primary_Location_Reference,omitempty"`
	AdditionalLocationsReference     []LocationObjectType                `xml:"urn:com.workday/bsvc Additional_Locations_Reference,omitempty"`
}

// Utilize the Request References element to retrieve a specific instance(s) of Evergreen Requisition and its associated data.
type EvergreenJobRequisitionRequestReferencesType struct {
	EvergreenJobRequisitionReference []EvergreenJobRequisitionObjectType `xml:"urn:com.workday/bsvc Evergreen_Job_Requisition_Reference"`
	SkipNonExistingInstances         bool                                `xml:"urn:com.workday/bsvc Skip_Non_Existing_Instances,attr,omitempty"`
}

// Contains each Evergreen Requisition based on the Request References or Request Criteria.
type EvergreenJobRequisitionResponseDataType struct {
	EvergreenJobRequsition []EvergreenJobRequisitionType `xml:"urn:com.workday/bsvc Evergreen_Job_Requsition,omitempty"`
}

// The response group allows for the response data to be tailored to only included elements that the user is looking for. If no response group is provided in the request then only the following elements will be returned: Reference, Evergreen Requisition Definition Data, and Evergreen Requisition Restrictions Data.
type EvergreenJobRequisitionResponseGroupType struct {
	IncludeReference                     *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeJobRequisitionDefinitionData  *bool `xml:"urn:com.workday/bsvc Include_Job_Requisition_Definition_Data,omitempty"`
	CreateJobRequisitionRestrictionsData *bool `xml:"urn:com.workday/bsvc Create_Job_Requisition_Restrictions_Data,omitempty"`
	IncludeQualifications                *bool `xml:"urn:com.workday/bsvc Include_Qualifications,omitempty"`
	IncludeJobRequisitionAttachments     *bool `xml:"urn:com.workday/bsvc Include_Job_Requisition_Attachments,omitempty"`
	IncludeOrganizations                 *bool `xml:"urn:com.workday/bsvc Include_Organizations,omitempty"`
	IncludeRoles                         *bool `xml:"urn:com.workday/bsvc Include_Roles,omitempty"`
}

// Contains each Evergreen Requisition based on the Request References or Request Criteria.
type EvergreenJobRequisitionType struct {
	EvergreenJobRequisitionReference *EvergreenJobRequisitionObjectType `xml:"urn:com.workday/bsvc Evergreen_Job_Requisition_Reference,omitempty"`
	EvergreenJobRequisitionData      *EvergreenJobRequisitionData2Type  `xml:"urn:com.workday/bsvc Evergreen_Job_Requisition_Data,omitempty"`
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

// Integration ID Help Text
type ExternalIntegrationIDDataType struct {
	ID []IDType `xml:"urn:com.workday/bsvc ID"`
}

// Contains a unique identifier for an instance of an object.
type ExternalURLObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ExternalURLObjectType struct {
	ID         []ExternalURLObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Field Result Request Criteria
type FieldAndParameterCriteriaDataType struct {
	ProviderReference []ExternalFieldandParameterInitializationProviderObjectType `xml:"urn:com.workday/bsvc Provider_Reference"`
}

// Contains a unique identifier for an instance of an object.
type FieldofStudyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type FieldofStudyObjectType struct {
	ID         []FieldofStudyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Container for the processing options for sub-business processes within a business process. If no options are submitted (or the options are submitted as 'false') then the sub-business process is simply initiated as if it where submitted on-line with approvals, reviews, notifications and to-do's in place. If the Initiator is an Integration System User, any validations you configured on the Initiation step are ignored.
type FinancialsBusinessSubProcessParametersType struct {
	Skip                       *bool                           `xml:"urn:com.workday/bsvc Skip,omitempty"`
	BusinessProcessCommentData *BusinessProcessCommentDataType `xml:"urn:com.workday/bsvc Business_Process_Comment_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type FormerWorkerObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type FormerWorkerObjectType struct {
	ID         []FormerWorkerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type GenderIdentityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type GenderIdentityObjectType struct {
	ID         []GenderIdentityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type GenericJobPostingSiteObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type GenericJobPostingSiteObjectType struct {
	ID         []GenericJobPostingSiteObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for Applicant Request.
type GetApplicantsRequestType struct {
	RequestReferences *ApplicantRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *ApplicantRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *ApplicantResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Applicant Information returned as a result of a Get Applicant request.
type GetApplicantsResponseType struct {
	RequestReferences *ApplicantRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *ApplicantRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *ApplicantResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType            `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *ApplicantsResponseDataType     `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// A request for information on a Candidate Assessment.
type GetAssessCandidateRequestType struct {
	RequestReferences *AssessCandidateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *AssessCandidateRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response to a web service call to return information on a Candidate Assessment.
type GetAssessCandidateResponseType struct {
	RequestReferences *AssessCandidateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *AssessCandidateRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                  `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *AssessCandidateResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top level request for Recruiting Self-Schedule Calendars
type GetAssignRecruitingSelfScheduleCalendarsRequestType struct {
	RequestReferences *AssignRecruitingSelfScheduleCalendarRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top level response for the Assign Recruiting Self-Schedule Event Calendars Web Service
type GetAssignRecruitingSelfScheduleCalendarsResponseType struct {
	RequestReferences *AssignRecruitingSelfScheduleCalendarRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                                       `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *AssignRecruitingSelfScheduleCalendarResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Background Check Packages data.
type GetBackgroundCheckPackagesRequestType struct {
	RequestReferences *BackgroundCheckPackageRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BackgroundCheckPackageResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element containing Background Check Package Response Data for requested references or criteria and for requested response group.
type GetBackgroundCheckPackagesResponseType struct {
	RequestReferences *BackgroundCheckPackageRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BackgroundCheckPackageResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                         `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BackgroundCheckPackageResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// All the data for requesting details for Background Check Events. This web service operation assumes the Background Check Events use Background Check Package.
type GetBackgroundCheckRequestType struct {
	RequestReferences *BackgroundCheckRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BackgroundCheckRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Data for Background Check Events. Used for Background Check Events that have Background Check Packages.
type GetBackgroundCheckResponseType struct {
	RequestReferences *BackgroundCheckRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BackgroundCheckRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                  `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BackgroundCheckResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Candidate attachment data.
type GetCandidateAttachmentsRequestType struct {
	RequestReferences *CandidateAttachmentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CandidateAttachmentRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CandidateAttachmentResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element containing Candidate Attachment Response Data for requested references or criteria and for requested response group.
type GetCandidateAttachmentsResponseType struct {
	RequestReferences *CandidateAttachmentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CandidateAttachmentRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CandidateAttachmentResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                      `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *CandidateAttachmentResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Candidate photo data.
type GetCandidatePhotosRequestType struct {
	RequestReferences *CandidateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CandidateResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element containing Candidate Photo Response Data for requested references or criteria and for requested response group.
type GetCandidatePhotosResponseType struct {
	RequestReferences []CandidateRequestReferencesType  `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    []ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     []CandidateResponseGroupType      `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   []ResponseResultsType             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []CandidatePhotosResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Candidate data.
type GetCandidatesRequestType struct {
	RequestReferences *CandidateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CandidateRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CandidateResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element containing Candidate Response Data for requested references or criteria and for requested response group.
type GetCandidatesResponseType struct {
	RequestReferences *CandidateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CandidateRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CandidateResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType            `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *CandidateResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to get evergreen requisitions and their associated data.
type GetEvergreenRequisitionsRequestType struct {
	RequestReferences *EvergreenJobRequisitionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *EvergreenJobRequisitionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                           `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *EvergreenJobRequisitionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing instances of Evergreen Requisitions and their associated data.
type GetEvergreenRequisitionsResponseType struct {
	RequestReferences *EvergreenJobRequisitionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *EvergreenJobRequisitionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                           `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *EvergreenJobRequisitionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                          `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *EvergreenJobRequisitionResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting interview feedback data.
type GetInterviewFeedbacksRequestType struct {
	RequestReferences *InterviewFeedbackRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *InterviewFeedbackRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *InterviewFeedbackResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element containing Interview Feedback Response Data for requested references or criteria and for requested response group.
type GetInterviewFeedbacksResponseType struct {
	RequestReferences *InterviewFeedbackRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *InterviewFeedbackRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *InterviewFeedbackResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *InterviewFeedbackResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting interview data.
type GetInterviewsRequestType struct {
	RequestReferences *InterviewRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *InterviewRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *InterviewResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element containing Interview Response Data for requested references or criteria and for requested response group.
type GetInterviewsResponseType struct {
	RequestReferences *InterviewRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *InterviewRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *InterviewResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType            `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *InterviewResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Job Application Additional data.
type GetJobApplicationAdditionalDataRequestType struct {
	RequestReferences *JobApplicationRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *JobApplicationRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobApplicationResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element containing Candidate Job Application Additional Data for requested references or criteria and for requested response group.
type GetJobApplicationAdditionalDataResponseType struct {
	RequestReferences *JobApplicationRequestReferencesType          `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *JobApplicationRequestCriteriaType            `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                           `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobApplicationResponseGroupType              `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                          `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *JobApplicationAdditionalDataResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Outer Container which holds all data and criteria for retrieving Job Posting Site(s).
type GetJobPostingSitesRequestType struct {
	RequestReferences *JobPostingSiteRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *JobPostingSiteRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobPostingSiteResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response for Job Posting Site
type GetJobPostingSitesResponseType struct {
	RequestReferences *JobPostingSiteRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *JobPostingSiteRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobPostingSiteResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *JobPostingSiteResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to get Job Postings and their associated data.
type GetJobPostingsRequestType struct {
	RequestReferences *JobPostingRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *JobPostingRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobPostingResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing instances of Job Postings and their associated data.
type GetJobPostingsResponseType struct {
	RequestReferences *JobPostingRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *JobPostingRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobPostingResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *JobPostingResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter, and response group for getting job requisition interview team data.
type GetJobRequisitionInterviewTeamsRequestType struct {
	RequestReferences *JobRequisitionInterviewTeamRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobRequisitionInterviewTeamResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element containing Job Requisition Interview Team data for requested references or criteria and for requested response group.
type GetJobRequisitionInterviewTeamsResponseType struct {
	RequestReferences *JobRequisitionInterviewTeamRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobRequisitionInterviewTeamResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                              `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *JobRequisitionInterviewTeamResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to get job requisitions and their associated data.
type GetJobRequisitionsRequestType struct {
	RequestReferences *JobRequisitionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *JobRequisitionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobRequisitionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing instances of Job Requisitions and their associated data.
type GetJobRequisitionsResponseType struct {
	RequestReferences *JobRequisitionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *JobRequisitionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobRequisitionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *JobRequisitionResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Root element for Get Organizations operation request
type GetOrganizationsRequestType struct {
	RequestReferences *OrganizationRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *OrganizationRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *OrganizationResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains each Organization based on the Request References or Request Criteria.  The data returned is the current information as of the dates in the response filter, and does not include all historical information about the Organization.
type GetOrganizationsResponseType struct {
	RequestReferences *OrganizationRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *OrganizationRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *OrganizationResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType               `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *OrganizationResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element  to get position management positions and their associated data.
type GetPositionsRequestType struct {
	RequestReferences *PositionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PositionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PositionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing instances of positions and their associated data.
type GetPositionsResponseType struct {
	RequestReferences          []PositionRequestReferencesType  `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria            []PositionRequestCriteriaType    `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter             []ResponseFilterType             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup              []PositionResponseGroupType      `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults            []ResponseResultsType            `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData               []PositionsResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	InvalidReferenceIDResponse []InvalidReferenceIDResponseType `xml:"urn:com.workday/bsvc Invalid_Reference_ID_Response,omitempty"`
	Version                    string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Recruiting Agency data.
type GetRecruitingAgenciesRequestType struct {
	RequestReferences *RecruitingAgencyRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *RecruitingAgencyRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Recruiting Agencies Response
type GetRecruitingAgenciesResponseType struct {
	RequestReferences *RecruitingAgencyRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *RecruitingAgencyRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *RecruitingAgencyResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Outer Container which holds all data and criteria for retrieving Recruiting Agency  additional data.
type GetRecruitingAgencyAdditionalDataRequestType struct {
	RequestReferences *RecruitingAgencyRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *RecruitingAgencyRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *RecruitingAgencyUserResponseGroupType `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Recruiting Agency Additional Data Response
type GetRecruitingAgencyAdditionalDataResponseType struct {
	RequestReferences []RecruitingAgencyRequestReferencesType          `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   []RecruitingAgencyRequestCriteriaType            `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    []ResponseFilterType                             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     []RecruitingAgencyUserResponseGroupType          `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   []ResponseResultsType                            `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []RecruitingAgencyAdditionalDataResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Recruiting Agency User data.
type GetRecruitingAgencyUsersRequestType struct {
	RequestReferences *RecruitingAgencyUserRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *RecruitingAgencyUserResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element containing Recruiting Agency Users Response Data for requested references or criteria and for requested response group.
type GetRecruitingAgencyUsersResponseType struct {
	RequestReferences *RecruitingAgencyUserRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *RecruitingAgencyUserResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                       `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *RecruitingAgencyUserResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Retrieves Recruiting Assessment Categories
type GetRecruitingAssessmentCategoriesRequestType struct {
	RequestReferences *RecruitingAssessmentCategoryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     []RecruitingAssessmentCategoryResponseGroupType    `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Recruiting Assessment Categories Response Container
type GetRecruitingAssessmentCategoriesResponseType struct {
	RequestReferences *RecruitingAssessmentCategoryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     []RecruitingAssessmentCategoryResponseGroupType    `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                               `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *RecruitingAssessmentCategoryResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the request to retrieve Recruiting Assessment Category Security Segment(s).
type GetRecruitingAssessmentCategorySecuritySegmentsRequestType struct {
	RequestReferences *RecruitingAssessmentCategorySecuritySegmentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *RecruitingAssessmentCategorySecuritySegmentRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *RecruitingAssessmentCategorySecuritySegmentResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the Recruiting Assessment Category Security Segment(s) retrieved by the web service.
type GetRecruitingAssessmentCategorySecuritySegmentsResponseType struct {
	RequestReferences *RecruitingAssessmentCategorySecuritySegmentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *RecruitingAssessmentCategorySecuritySegmentRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *RecruitingAssessmentCategorySecuritySegmentResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                                              `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *RecruitingAssessmentCategorySecuritySegmentResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the request to retrieve Recruiting Self-Schedule Calendar Type(s).
type GetRecruitingSelfScheduleCalendarTypesRequestType struct {
	RequestReferences *RecruitingSelfScheduleCalendarTypeRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *RecruitingSelfScheduleCalendarTypeRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *RecruitingSelfScheduleCalendarTypeResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the Recruiting Self-Schedule Calendar Type(s) retrieved by the web service.
type GetRecruitingSelfScheduleCalendarTypesResponseType struct {
	RequestReferences *RecruitingSelfScheduleCalendarTypeRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *RecruitingSelfScheduleCalendarTypeRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *RecruitingSelfScheduleCalendarTypeResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                                     `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *RecruitingSelfScheduleCalendarTypeResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the request to retrieve Recruiting Self-Schedule Calendar(s).
type GetRecruitingSelfScheduleCalendarsRequestType struct {
	RequestReferences *RecruitingSelfScheduleCalendarRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *RecruitingSelfScheduleCalendarRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *RecruitingSelfScheduleCalendarResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the Recruiting Self-Schedule Calendar(s) retrieved by the web service.
type GetRecruitingSelfScheduleCalendarsResponseType struct {
	RequestReferences *RecruitingSelfScheduleCalendarRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *RecruitingSelfScheduleCalendarRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *RecruitingSelfScheduleCalendarResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *RecruitingSelfScheduleCalendarResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, and filter for getting Veteran Status data.
type GetVeteranStatusesRequestType struct {
	RequestReferences *VeteranStatusRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *VeteranStatusRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Veteran Status data.
type GetVeteranStatusesResponseType struct {
	RequestReferences *VeteranStatusRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *VeteranStatusRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *VeteranStatusResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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
type GlobalSetupDataMappingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type GlobalSetupDataMappingObjectType struct {
	ID         []GlobalSetupDataMappingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type GrantObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type GrantObjectType struct {
	ID         []GrantObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type HukouTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type HukouTypeObjectType struct {
	ID         []HukouTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// External ID that uniquely identifies the integratable object within the context of the integration system identified by the System ID attribute.
type IDType struct {
	Value    string `xml:",chardata"`
	SystemID string `xml:"urn:com.workday/bsvc System_ID,attr,omitempty"`
}

// Documentation Import Applicant Request element
type ImportApplicantRequestType struct {
	ID        string                       `xml:"urn:com.workday/bsvc ID,omitempty"`
	Applicant []ApplicantInformationHVType `xml:"urn:com.workday/bsvc Applicant,omitempty"`
	Version   string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The data used to create or modify a Candidate Assessment
type InlineAssessmentDataType struct {
	JobApplicationReference      *JobApplicationObjectType        `xml:"urn:com.workday/bsvc Job_Application_Reference,omitempty"`
	CandidateReference           *CandidateObjectType             `xml:"urn:com.workday/bsvc Candidate_Reference,omitempty"`
	JobRequisitionReference      *JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	JobApplicationEventReference *UniqueIdentifierObjectType      `xml:"urn:com.workday/bsvc Job_Application_Event_Reference,omitempty"`
	CandidateAssessmentData      *RecruitingAssessmentDataType    `xml:"urn:com.workday/bsvc Candidate_Assessment_Data,omitempty"`
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

// Wrapper element for the Internal Project Experience information.
type InternalProjectExperienceDataType struct {
	InternalProjectExperienceID     string     `xml:"urn:com.workday/bsvc Internal_Project_Experience_ID,omitempty"`
	RemoveInternalProjectExperience *bool      `xml:"urn:com.workday/bsvc Remove_Internal_Project_Experience,omitempty"`
	InternalProject                 string     `xml:"urn:com.workday/bsvc Internal_Project,omitempty"`
	Description                     string     `xml:"urn:com.workday/bsvc Description,omitempty"`
	StartDate                       *time.Time `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                         *time.Time `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	ProjectLeader                   string     `xml:"urn:com.workday/bsvc Project_Leader,omitempty"`
}

func (t *InternalProjectExperienceDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InternalProjectExperienceDataType
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
func (t *InternalProjectExperienceDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InternalProjectExperienceDataType
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

// Contains a unique identifier for an instance of an object.
type InternalProjectExperienceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InternalProjectExperienceObjectType struct {
	ID         []InternalProjectExperienceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for Internal Project Experience information.
type InternalProjectExperienceType struct {
	InternalProjectExperienceReference *InternalProjectExperienceObjectType `xml:"urn:com.workday/bsvc Internal_Project_Experience_Reference,omitempty"`
	InternalProjectExperienceData      []InternalProjectExperienceDataType  `xml:"urn:com.workday/bsvc Internal_Project_Experience_Data,omitempty"`
}

// Contains one of the following: Interview Event Reference, Job Application Reference, or the Candidate Interview Criteria Data.
type InterviewDataType struct {
	InterviewEventReference        *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Interview_Event_Reference"`
	JobApplicationReference        *JobApplicationObjectType   `xml:"urn:com.workday/bsvc Job_Application_Reference"`
	CandidateInterviewCriteriaData *MoveCandidateCriteriaType  `xml:"urn:com.workday/bsvc Candidate_Interview_Criteria_Data"`
	InterviewSessionData           []InterviewSessionDataType  `xml:"urn:com.workday/bsvc Interview_Session_Data,omitempty"`
	OverallComment                 string                      `xml:"urn:com.workday/bsvc Overall_Comment,omitempty"`
	TimezoneReference              *TimeZoneObjectType         `xml:"urn:com.workday/bsvc Timezone_Reference"`
}

// Contains one of the following: Interview Event Reference, Job Application Reference, or the Candidate Interview Criteria Data.
type InterviewFeedbackDataType struct {
	InterviewEventReference        *UniqueIdentifierObjectType        `xml:"urn:com.workday/bsvc Interview_Event_Reference"`
	JobApplicationReference        *JobApplicationObjectType          `xml:"urn:com.workday/bsvc Job_Application_Reference"`
	CandidateInterviewCriteriaData *MoveCandidateCriteriaType         `xml:"urn:com.workday/bsvc Candidate_Interview_Criteria_Data"`
	InterviewFeedbackDetailsData   []InterviewFeedbackDetailsDataType `xml:"urn:com.workday/bsvc Interview_Feedback_Details_Data,omitempty"`
}

// Contains the interviewer reference and its feedback data.
type InterviewFeedbackDetailsDataType struct {
	InterviewerReference     *WorkerObjectType              `xml:"urn:com.workday/bsvc Interviewer_Reference,omitempty"`
	InterviewFeedbackDetails []InterviewFeedbackDetailsType `xml:"urn:com.workday/bsvc Interview_Feedback_Details,omitempty"`
}

// Contains the overall rating and comment for the interview feedback.
type InterviewFeedbackDetailsType struct {
	CompetencyAssessmentData         []CompetencyAssessmentDataType     `xml:"urn:com.workday/bsvc Competency_Assessment_Data,omitempty"`
	InterviewQuestionnaireData       []InterviewQuestionnaireDataType   `xml:"urn:com.workday/bsvc Interview_Questionnaire_Data,omitempty"`
	InterviewFeedbackRatingReference *InterviewFeedbackRatingObjectType `xml:"urn:com.workday/bsvc Interview_Feedback_Rating_Reference,omitempty"`
	OverallComment                   string                             `xml:"urn:com.workday/bsvc Overall_Comment,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type InterviewFeedbackRatingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InterviewFeedbackRatingObjectType struct {
	ID         []InterviewFeedbackRatingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Criteria used to determine which interview to return.
type InterviewFeedbackRequestCriteriaType struct {
	JobApplicationReference        *JobApplicationObjectType  `xml:"urn:com.workday/bsvc Job_Application_Reference"`
	CandidateInterviewCriteriaData *MoveCandidateCriteriaType `xml:"urn:com.workday/bsvc Candidate_Interview_Criteria_Data"`
}

// Contains the interview reference.
type InterviewFeedbackRequestReferencesType struct {
	InterviewEventReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Interview_Event_Reference"`
}

// Contains interview feedback data.
type InterviewFeedbackResponseDataType struct {
	InterviewFeedback []InterviewFeedbackType `xml:"urn:com.workday/bsvc Interview_Feedback,omitempty"`
}

// Element containing Interview Feedback response group options.
type InterviewFeedbackResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Contains interview feedback data.
type InterviewFeedbackType struct {
	InterviewFeedbackData []InterviewFeedbackDataType `xml:"urn:com.workday/bsvc Interview_Feedback_Data,omitempty"`
}

// Interview Feedback Details
type InterviewQuestionnaireDataType struct {
	QuestionnairesResponseData []QuestionnairesResponseDataType `xml:"urn:com.workday/bsvc Questionnaires_Response_Data,omitempty"`
}

// Criteria used to determine which interview to return.
type InterviewRequestCriteriaType struct {
	JobApplicationReference        *JobApplicationObjectType  `xml:"urn:com.workday/bsvc Job_Application_Reference"`
	CandidateInterviewCriteriaData *MoveCandidateCriteriaType `xml:"urn:com.workday/bsvc Candidate_Interview_Criteria_Data"`
}

// Contains the interview reference.
type InterviewRequestReferencesType struct {
	InterviewReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Interview_Reference"`
}

// Contains interview data.
type InterviewResponseDataType struct {
	Interview []InterviewType `xml:"urn:com.workday/bsvc Interview,omitempty"`
}

// Element containing Interview response group options.
type InterviewResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Contains the interview session reference and its data.
type InterviewSessionDataType struct {
	InterviewSessionReference   *UniqueIdentifierObjectType       `xml:"urn:com.workday/bsvc Interview_Session_Reference,omitempty"`
	InterviewSessionDetailsData []InterviewSessionDetailsDataType `xml:"urn:com.workday/bsvc Interview_Session_Details_Data,omitempty"`
	Delete                      bool                              `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains the interviewers, duration in minutes, interview start time, interview end time, interview type, and comment data.
type InterviewSessionDetailsDataType struct {
	InterviewersReference            []WorkerObjectType        `xml:"urn:com.workday/bsvc Interviewers_Reference"`
	InterviewStartTime               time.Time                 `xml:"urn:com.workday/bsvc Interview_Start_Time"`
	InterviewEndTime                 time.Time                 `xml:"urn:com.workday/bsvc Interview_End_Time"`
	InterviewLocationReference       *LocationObjectType       `xml:"urn:com.workday/bsvc Interview_Location_Reference,omitempty"`
	InterviewTypeReference           *InterviewTypeObjectType  `xml:"urn:com.workday/bsvc Interview_Type_Reference,omitempty"`
	InterviewCompetenciesReference   []CompetencyObjectType    `xml:"urn:com.workday/bsvc Interview_Competencies_Reference,omitempty"`
	InterviewQuestionnairesReference []QuestionnaireObjectType `xml:"urn:com.workday/bsvc Interview_Questionnaires_Reference,omitempty"`
	Comment                          string                    `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

func (t *InterviewSessionDetailsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InterviewSessionDetailsDataType
	var layout struct {
		*T
		InterviewStartTime *xsdDateTime `xml:"urn:com.workday/bsvc Interview_Start_Time"`
		InterviewEndTime   *xsdDateTime `xml:"urn:com.workday/bsvc Interview_End_Time"`
	}
	layout.T = (*T)(t)
	layout.InterviewStartTime = (*xsdDateTime)(&layout.T.InterviewStartTime)
	layout.InterviewEndTime = (*xsdDateTime)(&layout.T.InterviewEndTime)
	return e.EncodeElement(layout, start)
}
func (t *InterviewSessionDetailsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InterviewSessionDetailsDataType
	var overlay struct {
		*T
		InterviewStartTime *xsdDateTime `xml:"urn:com.workday/bsvc Interview_Start_Time"`
		InterviewEndTime   *xsdDateTime `xml:"urn:com.workday/bsvc Interview_End_Time"`
	}
	overlay.T = (*T)(t)
	overlay.InterviewStartTime = (*xsdDateTime)(&overlay.T.InterviewStartTime)
	overlay.InterviewEndTime = (*xsdDateTime)(&overlay.T.InterviewEndTime)
	return d.DecodeElement(&overlay, &start)
}

// Contains interview data.
type InterviewType struct {
	InterviewData []InterviewDataType `xml:"urn:com.workday/bsvc Interview_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type InterviewTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InterviewTypeObjectType struct {
	ID         []InterviewTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Represents a reference ID that was submitted with data not found in Workday
type InvalidReferenceIDDataType struct {
	InvalidReferenceID     string `xml:"urn:com.workday/bsvc Invalid_Reference_ID,omitempty"`
	InvalidReferenceIDType string `xml:"urn:com.workday/bsvc Invalid_Reference_ID_Type,omitempty"`
}

// Invalid Reference ID Response
type InvalidReferenceIDResponseType struct {
	InvalidReference []InvalidReferenceIDDataType `xml:"urn:com.workday/bsvc Invalid_Reference,omitempty"`
}

// Contains Job Application Additional data
type JobApplicationAdditionalDataResponseDataType struct {
	JobApplicationCustomObjectData []JobApplicationType `xml:"urn:com.workday/bsvc Job_Application_Custom_Object_Data,omitempty"`
}

// Contains Information about the Job Application to update and Attachment to be added.
type JobApplicationAttachmentDataType struct {
	JobRequisitionReference *JobRequisitionObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	JobApplicationReference *JobApplicationObjectType `xml:"urn:com.workday/bsvc Job_Application_Reference,omitempty"`
	AttachmentData          *AttachmentWWSDataType    `xml:"urn:com.workday/bsvc Attachment_Data"`
}

// Contains information about the Job Application the Attachment was added to.
type JobApplicationAttachmentResponseType struct {
	JobApplicationReference           *JobApplicationObjectType    `xml:"urn:com.workday/bsvc Job_Application_Reference,omitempty"`
	JobApplicationAttachmentReference []ResumeAttachmentObjectType `xml:"urn:com.workday/bsvc Job_Application_Attachment_Reference,omitempty"`
	JobApplicationsUpdated            string                       `xml:"urn:com.workday/bsvc Job_Applications_Updated,omitempty"`
}

// Job Application Additional Data
type JobApplicationCustomObjectDataType struct {
	JobApplicationReference      *JobApplicationObjectType                      `xml:"urn:com.workday/bsvc Job_Application_Reference"`
	BusinessObjectAdditionalData *NonEffectiveDatedWebServiceAdditionalDataType `xml:"urn:com.workday/bsvc Business_Object_Additional_Data"`
}

// Contains information about the Job Application for Candidate.
type JobApplicationDataType struct {
	JobAppliedToData     []CandidateJobAppliedToDataType `xml:"urn:com.workday/bsvc Job_Applied_To_Data"`
	ResumeAttachmentData []AttachmentWWSDataType         `xml:"urn:com.workday/bsvc Resume_Attachment_Data,omitempty"`
	ResumeData           *CandidateResumeDataType        `xml:"urn:com.workday/bsvc Resume_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JobApplicationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobApplicationObjectType struct {
	ID         []JobApplicationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Utilize the Request Criteria element to filter instance(s) of Job Application Additional Data by Candidate
type JobApplicationRequestCriteriaType struct {
	CandidateReference *CandidateObjectType `xml:"urn:com.workday/bsvc Candidate_Reference,omitempty"`
}

// Job Application Reference
type JobApplicationRequestReferencesType struct {
	JobApplicationReference []JobApplicationObjectType `xml:"urn:com.workday/bsvc Job_Application_Reference"`
}

// Job Application Additional Data response group
type JobApplicationResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JobApplicationTemplateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobApplicationTemplateObjectType struct {
	ID         []JobApplicationTemplateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Job Application Additional Data
type JobApplicationType struct {
	JobApplicationReference      *JobApplicationObjectType                      `xml:"urn:com.workday/bsvc Job_Application_Reference,omitempty"`
	BusinessObjectAdditionalData *NonEffectiveDatedWebServiceAdditionalDataType `xml:"urn:com.workday/bsvc Business_Object_Additional_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JobCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobCategoryObjectType struct {
	ID         []JobCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JobFamilyBaseObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobFamilyBaseObjectType struct {
	ID         []JobFamilyBaseObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JobFamilyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobFamilyObjectType struct {
	ID         []JobFamilyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for the job history information.
type JobHistoryAchievementDataType struct {
	JobHistoryID                    string                       `xml:"urn:com.workday/bsvc Job_History_ID,omitempty"`
	RemoveJobHistory                *bool                        `xml:"urn:com.workday/bsvc Remove_Job_History,omitempty"`
	JobTitle                        string                       `xml:"urn:com.workday/bsvc Job_Title,omitempty"`
	Company                         string                       `xml:"urn:com.workday/bsvc Company,omitempty"`
	JobHistoryCompanyReference      *JobHistoryCompanyObjectType `xml:"urn:com.workday/bsvc Job_History_Company_Reference,omitempty"`
	StartDate                       *time.Time                   `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                         *time.Time                   `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	ResponsibilitiesandAchievements string                       `xml:"urn:com.workday/bsvc Responsibilities_and_Achievements,omitempty"`
	Location                        string                       `xml:"urn:com.workday/bsvc Location,omitempty"`
	JobReference                    string                       `xml:"urn:com.workday/bsvc Job_Reference,omitempty"`
	Contact                         string                       `xml:"urn:com.workday/bsvc Contact,omitempty"`
}

func (t *JobHistoryAchievementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobHistoryAchievementDataType
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
func (t *JobHistoryAchievementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobHistoryAchievementDataType
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

// Contains a unique identifier for an instance of an object.
type JobHistoryCompanyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobHistoryCompanyObjectType struct {
	ID         []JobHistoryCompanyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JobHistorySkillObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobHistorySkillObjectType struct {
	ID         []JobHistorySkillObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for Job History information.
type JobHistoryType struct {
	JobHistoryReference *JobHistorySkillObjectType      `xml:"urn:com.workday/bsvc Job_History_Reference,omitempty"`
	JobHistoryData      []JobHistoryAchievementDataType `xml:"urn:com.workday/bsvc Job_History_Data"`
}

// Contains References and the External Name value for a Job Posting's additional locations.
type JobPostingAdditionalLocationDataType struct {
	AdditionalLocationReference *LocationObjectType `xml:"urn:com.workday/bsvc Additional_Location_Reference,omitempty"`
	ExternalName                string              `xml:"urn:com.workday/bsvc External_Name,omitempty"`
}

// Data for individual Job Postings.
type JobPostingDataType struct {
	JobPostingReference     *JobPostingObjectType     `xml:"urn:com.workday/bsvc Job_Posting_Reference,omitempty"`
	JobPostingSiteReference *JobPostingSiteObjectType `xml:"urn:com.workday/bsvc Job_Posting_Site_Reference,omitempty"`
	PrimaryPosting          *bool                     `xml:"urn:com.workday/bsvc Primary_Posting,omitempty"`
}

// Wrapper Element for a Job Posting.
type JobPostingDataWWSType struct {
	JobPostingID                      string                                          `xml:"urn:com.workday/bsvc Job_Posting_ID,omitempty"`
	JobPostingTitle                   string                                          `xml:"urn:com.workday/bsvc Job_Posting_Title,omitempty"`
	JobPostingDescription             string                                          `xml:"urn:com.workday/bsvc Job_Posting_Description,omitempty"`
	JobPostingSiteReference           *JobPostingSiteObjectType                       `xml:"urn:com.workday/bsvc Job_Posting_Site_Reference,omitempty"`
	ExternalJobPath                   string                                          `xml:"urn:com.workday/bsvc External_Job_Path,omitempty"`
	ExternalApplyURL                  string                                          `xml:"urn:com.workday/bsvc External_Apply_URL,omitempty"`
	JobRequisitionReference           *JobRequisitionEnabledObjectType                `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	JobRequisitionID                  string                                          `xml:"urn:com.workday/bsvc Job_Requisition_ID,omitempty"`
	JobPostingLocationData            *JobPostingLocationDataType                     `xml:"urn:com.workday/bsvc Job_Posting_Location_Data,omitempty"`
	JobPostingStartDate               *time.Time                                      `xml:"urn:com.workday/bsvc Job_Posting_Start_Date,omitempty"`
	JobPostingEndDate                 *time.Time                                      `xml:"urn:com.workday/bsvc Job_Posting_End_Date,omitempty"`
	PrimaryPosting                    *bool                                           `xml:"urn:com.workday/bsvc Primary_Posting,omitempty"`
	ForecastedPayout                  float64                                         `xml:"urn:com.workday/bsvc Forecasted_Payout,omitempty"`
	ForecastedPayoutCurrencyReference *CurrencyObjectType                             `xml:"urn:com.workday/bsvc Forecasted_Payout_Currency_Reference,omitempty"`
	JobFamilyReference                []JobFamilyObjectType                           `xml:"urn:com.workday/bsvc Job_Family_Reference,omitempty"`
	TimeTypeReference                 *PositionTimeTypeObjectType                     `xml:"urn:com.workday/bsvc Time_Type_Reference,omitempty"`
	JobTypeReference                  *PositionWorkerTypeObjectType                   `xml:"urn:com.workday/bsvc Job_Type_Reference,omitempty"`
	SupervisoryOrganizationReference  *SupervisoryOrganizationObjectType              `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference,omitempty"`
	StudentAwardSourceReference       []StudentAwardSourceObjectType                  `xml:"urn:com.workday/bsvc Student_Award_Source_Reference,omitempty"`
	UnionReference                    []UnionObjectType                               `xml:"urn:com.workday/bsvc Union_Reference,omitempty"`
	SimilarJobsReference              []JobPostingObjectType                          `xml:"urn:com.workday/bsvc Similar_Jobs_Reference,omitempty"`
	HiringRequirementData             *JobRequisitionRestrictionsDataType             `xml:"urn:com.workday/bsvc Hiring_Requirement_Data,omitempty"`
	JobRequisitionDetailData          *JobRequisitionDefinitionDataType               `xml:"urn:com.workday/bsvc Job_Requisition_Detail_Data,omitempty"`
	QualificationData                 *QualificationsfromPositionRestrictionsDataType `xml:"urn:com.workday/bsvc Qualification_Data,omitempty"`
	JobRequisitionAttachments         *JobRequisitionAttachmentsType                  `xml:"urn:com.workday/bsvc Job_Requisition_Attachments,omitempty"`
	DocumentFieldResultData           []DocumentFieldResultDataType                   `xml:"urn:com.workday/bsvc Document_Field_Result_Data,omitempty"`
}

func (t *JobPostingDataWWSType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobPostingDataWWSType
	var layout struct {
		*T
		JobPostingStartDate *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_Start_Date,omitempty"`
		JobPostingEndDate   *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.JobPostingStartDate = (*xsdDate)(layout.T.JobPostingStartDate)
	layout.JobPostingEndDate = (*xsdDate)(layout.T.JobPostingEndDate)
	return e.EncodeElement(layout, start)
}
func (t *JobPostingDataWWSType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobPostingDataWWSType
	var overlay struct {
		*T
		JobPostingStartDate *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_Start_Date,omitempty"`
		JobPostingEndDate   *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.JobPostingStartDate = (*xsdDate)(overlay.T.JobPostingStartDate)
	overlay.JobPostingEndDate = (*xsdDate)(overlay.T.JobPostingEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains information about a Job Posting's primary and additional Locations.
type JobPostingLocationDataType struct {
	PrimaryLocationReference         *LocationObjectType                    `xml:"urn:com.workday/bsvc Primary_Location_Reference,omitempty"`
	ExternalName                     string                                 `xml:"urn:com.workday/bsvc External_Name,omitempty"`
	JobPostingAdditionalLocationData []JobPostingAdditionalLocationDataType `xml:"urn:com.workday/bsvc Job_Posting_Additional_Location_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JobPostingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobPostingObjectType struct {
	ID         []JobPostingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Data used to identify the Job Posting when no Job Posting Reference is supplied.
type JobPostingReferenceDataType struct {
	JobPostingSiteName      string                           `xml:"urn:com.workday/bsvc Job_Posting_Site_Name"`
	JobPostingSiteReference *JobPostingSiteObjectType        `xml:"urn:com.workday/bsvc Job_Posting_Site_Reference"`
	JobRequisitionReference *JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
}

// Utilize the Request Criteria element to filter instance(s) of Job Postings by date, posting site, supervisory organization, job requisition, and active status.
type JobPostingRequestCriteriaType struct {
	JobPostingStartDate              *time.Time                          `xml:"urn:com.workday/bsvc Job_Posting_Start_Date,omitempty"`
	JobPostingEndDate                *time.Time                          `xml:"urn:com.workday/bsvc Job_Posting_End_Date,omitempty"`
	JobPostingSiteName               string                              `xml:"urn:com.workday/bsvc Job_Posting_Site_Name,omitempty"`
	JobPostingSiteReference          []JobPostingSiteObjectType          `xml:"urn:com.workday/bsvc Job_Posting_Site_Reference,omitempty"`
	SupervisoryOrganizationReference []SupervisoryOrganizationObjectType `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference,omitempty"`
	JobRequisitionReference          []JobRequisitionEnabledObjectType   `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	ShowOnlyActiveJobPostings        *bool                               `xml:"urn:com.workday/bsvc Show_Only_Active_Job_Postings,omitempty"`
	FieldAndParameterCriteriaData    *FieldAndParameterCriteriaDataType  `xml:"urn:com.workday/bsvc Field_And_Parameter_Criteria_Data,omitempty"`
}

func (t *JobPostingRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobPostingRequestCriteriaType
	var layout struct {
		*T
		JobPostingStartDate *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_Start_Date,omitempty"`
		JobPostingEndDate   *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.JobPostingStartDate = (*xsdDate)(layout.T.JobPostingStartDate)
	layout.JobPostingEndDate = (*xsdDate)(layout.T.JobPostingEndDate)
	return e.EncodeElement(layout, start)
}
func (t *JobPostingRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobPostingRequestCriteriaType
	var overlay struct {
		*T
		JobPostingStartDate *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_Start_Date,omitempty"`
		JobPostingEndDate   *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.JobPostingStartDate = (*xsdDate)(overlay.T.JobPostingStartDate)
	overlay.JobPostingEndDate = (*xsdDate)(overlay.T.JobPostingEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Utilize the Request References element to retrieve a specific instance(s) of Job Posting and its associated data.
type JobPostingRequestReferencesType struct {
	JobPostingReference []JobPostingObjectType `xml:"urn:com.workday/bsvc Job_Posting_Reference"`
}

// Contains each Job Posting based on the Request References or Request Criteria.
type JobPostingResponseDataType struct {
	JobPosting []JobPostingType `xml:"urn:com.workday/bsvc Job_Posting,omitempty"`
}

// The response group allows for the response data to be tailored to only included elements that the user is looking for.  If no response group is provided in the request then only the following elements will be returned:  Reference, Job Requisition Definition Data, and Job Requisition Restrictions Data.
type JobPostingResponseGroupType struct {
	IncludeReference                      *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeJobRequisitionRestrictionsData *bool `xml:"urn:com.workday/bsvc Include_Job_Requisition_Restrictions_Data,omitempty"`
	IncludeJobRequisitionDefinitionData   *bool `xml:"urn:com.workday/bsvc Include_Job_Requisition_Definition_Data,omitempty"`
	IncludeQualifications                 *bool `xml:"urn:com.workday/bsvc Include_Qualifications,omitempty"`
	IncludeJobRequisitionAttachments      *bool `xml:"urn:com.workday/bsvc Include_Job_Requisition_Attachments,omitempty"`
}

// Container for Job Posting Site Data.
type JobPostingSiteDataType struct {
	ID                      string                    `xml:"urn:com.workday/bsvc ID,omitempty"`
	JobPostingSiteName      string                    `xml:"urn:com.workday/bsvc Job_Posting_Site_Name"`
	JobPostingSiteReference *JobPostingSiteObjectType `xml:"urn:com.workday/bsvc Job_Posting_Site_Reference"`
	JobPostingStartDate     time.Time                 `xml:"urn:com.workday/bsvc Job_Posting_Start_Date"`
	JobPostingEndDate       *time.Time                `xml:"urn:com.workday/bsvc Job_Posting_End_Date,omitempty"`
	PrimaryPosting          *bool                     `xml:"urn:com.workday/bsvc Primary_Posting,omitempty"`
}

func (t *JobPostingSiteDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobPostingSiteDataType
	var layout struct {
		*T
		JobPostingStartDate *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_Start_Date"`
		JobPostingEndDate   *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.JobPostingStartDate = (*xsdDate)(&layout.T.JobPostingStartDate)
	layout.JobPostingEndDate = (*xsdDate)(layout.T.JobPostingEndDate)
	return e.EncodeElement(layout, start)
}
func (t *JobPostingSiteDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobPostingSiteDataType
	var overlay struct {
		*T
		JobPostingStartDate *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_Start_Date"`
		JobPostingEndDate   *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.JobPostingStartDate = (*xsdDate)(&overlay.T.JobPostingStartDate)
	overlay.JobPostingEndDate = (*xsdDate)(overlay.T.JobPostingEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Holds all relevant details for a Job Posting Site.
type JobPostingSiteDetailDataType struct {
	JobPostingSiteName          string                        `xml:"urn:com.workday/bsvc Job_Posting_Site_Name"`
	JobPostingSiteID            string                        `xml:"urn:com.workday/bsvc Job_Posting_Site_ID,omitempty"`
	ContractID                  string                        `xml:"urn:com.workday/bsvc Contract_ID,omitempty"`
	JobPostingCost              float64                       `xml:"urn:com.workday/bsvc Job_Posting_Cost,omitempty"`
	CurrencyReference           *CurrencyObjectType           `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	RecruitingSourceReference   *ApplicantSourceObjectType    `xml:"urn:com.workday/bsvc Recruiting_Source_Reference,omitempty"`
	JobPostingTemplateReference *JobPostingTemplateObjectType `xml:"urn:com.workday/bsvc Job_Posting_Template_Reference"`
	Inactive                    *bool                         `xml:"urn:com.workday/bsvc Inactive,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JobPostingSiteObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobPostingSiteObjectType struct {
	ID         []JobPostingSiteObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Get Job Posting Site via following criteria:      Posting Site Name or All Generic Job Posting Sites
type JobPostingSiteRequestCriteriaType struct {
	PostingSiteName        string `xml:"urn:com.workday/bsvc Posting_Site_Name,omitempty"`
	GenericJobPostingSites *bool  `xml:"urn:com.workday/bsvc Generic_Job_Posting_Sites,omitempty"`
}

// Get Job Posting Site via the Reference ID
type JobPostingSiteRequestReferencesType struct {
	JobPostingSiteReference []JobPostingSiteObjectType `xml:"urn:com.workday/bsvc Job_Posting_Site_Reference"`
}

// Container which holds Job Posting Sites returned by the GET
type JobPostingSiteResponseDataType struct {
	JobPostingSite []JobPostingSiteType `xml:"urn:com.workday/bsvc Job_Posting_Site,omitempty"`
}

// WWS responds with All Job Posting Sites if no reference ID and no criteria specified
type JobPostingSiteResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Container which holds each individual Job Posting Site Data
type JobPostingSiteType struct {
	JobPostingSiteReference *JobPostingSiteObjectType     `xml:"urn:com.workday/bsvc Job_Posting_Site_Reference,omitempty"`
	JobPostingSiteData      *JobPostingSiteDetailDataType `xml:"urn:com.workday/bsvc Job_Posting_Site_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JobPostingTemplateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobPostingTemplateObjectType struct {
	ID         []JobPostingTemplateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains each Job Posting based on the Request References or Request Criteria.
type JobPostingType struct {
	JobPostingReference *JobPostingObjectType  `xml:"urn:com.workday/bsvc Job_Posting_Reference,omitempty"`
	JobPostingData      *JobPostingDataWWSType `xml:"urn:com.workday/bsvc Job_Posting_Data,omitempty"`
}

// Element containing the Job Profile Exempt data for a Job Profile.
type JobProfileExemptDataType struct {
	LocationContextReference *LocationContextObjectType `xml:"urn:com.workday/bsvc Location_Context_Reference"`
	JobExempt                *bool                      `xml:"urn:com.workday/bsvc Job_Exempt,omitempty"`
	Delete                   bool                       `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
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

// Contains the details of a Job Profile.
type JobProfileSummaryDataType struct {
	JobProfileReference       *JobProfileObjectType       `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	JobProfileName            string                      `xml:"urn:com.workday/bsvc Job_Profile_Name,omitempty"`
	ManagementLevelReference  *ManagementLevelObjectType  `xml:"urn:com.workday/bsvc Management_Level_Reference,omitempty"`
	JobCategoryReference      *JobCategoryObjectType      `xml:"urn:com.workday/bsvc Job_Category_Reference,omitempty"`
	JobFamilyReference        []JobFamilyBaseObjectType   `xml:"urn:com.workday/bsvc Job_Family_Reference,omitempty"`
	WorkShiftRequired         *bool                       `xml:"urn:com.workday/bsvc Work_Shift_Required,omitempty"`
	JobProfileExemptData      []JobProfileExemptDataType  `xml:"urn:com.workday/bsvc Job_Profile_Exempt_Data,omitempty"`
	CriticalJob               *bool                       `xml:"urn:com.workday/bsvc Critical_Job,omitempty"`
	DifficultytoFillReference *DifficultytoFillObjectType `xml:"urn:com.workday/bsvc Difficulty_to_Fill_Reference,omitempty"`
}

// Assessments to be used at the beginning or during the Job Application process. For use by Recruiting customers only.
type JobRequisitionAssessmentDataType struct {
	InlineAssessmentTestReference   *RecruitingAssessmentTestObjectType  `xml:"urn:com.workday/bsvc Inline_Assessment_Test_Reference,omitempty"`
	DefaultAssessmentTestsReference []RecruitingAssessmentTestObjectType `xml:"urn:com.workday/bsvc Default_Assessment_Tests_Reference,omitempty"`
}

// Wrapper element for a signle file attachment
type JobRequisitionAttachmentDataType struct {
	FileName string  `xml:"urn:com.workday/bsvc File_Name"`
	Comment  string  `xml:"urn:com.workday/bsvc Comment,omitempty"`
	File     *[]byte `xml:"urn:com.workday/bsvc File,omitempty"`
}

func (t *JobRequisitionAttachmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobRequisitionAttachmentDataType
	var layout struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	layout.T = (*T)(t)
	layout.File = (*xsdBase64Binary)(layout.T.File)
	return e.EncodeElement(layout, start)
}
func (t *JobRequisitionAttachmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobRequisitionAttachmentDataType
	var overlay struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.File = (*xsdBase64Binary)(overlay.T.File)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for file Attachments for Job Requisition
type JobRequisitionAttachmentsType struct {
	JobRequisitionAttachmentData []JobRequisitionAttachmentDataType `xml:"urn:com.workday/bsvc Job_Requisition_Attachment_Data,omitempty"`
}

// Effective dated additional data for a job requisition.
type JobRequisitionCustomObjectDataType struct {
	EffectiveDate                *time.Time                                   `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	JobRequisitionReference      *JobRequisitionEnabledObjectType             `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	BusinessObjectAdditionalData []EffectiveDatedWebServiceAdditionalDataType `xml:"urn:com.workday/bsvc Business_Object_Additional_Data,omitempty"`
}

func (t *JobRequisitionCustomObjectDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobRequisitionCustomObjectDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *JobRequisitionCustomObjectDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobRequisitionCustomObjectDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper Element for a Job Requisition
type JobRequisitionData2Type struct {
	JobRequisitionStatusReference    *JobRequisitionStatusObjectType                 `xml:"urn:com.workday/bsvc Job_Requisition_Status_Reference,omitempty"`
	JobRequisitionDetailData         *JobRequisitionDefinitionDataType               `xml:"urn:com.workday/bsvc Job_Requisition_Detail_Data,omitempty"`
	HiringRequirementData            *JobRequisitionRestrictionsDataType             `xml:"urn:com.workday/bsvc Hiring_Requirement_Data,omitempty"`
	JobRequisitionAttachments        *JobRequisitionAttachmentsType                  `xml:"urn:com.workday/bsvc Job_Requisition_Attachments,omitempty"`
	QualificationData                *QualificationsfromPositionRestrictionsDataType `xml:"urn:com.workday/bsvc Qualification_Data,omitempty"`
	PositionData                     []PositionsDataType                             `xml:"urn:com.workday/bsvc Position_Data,omitempty"`
	QuestionnaireReference           *JobRequisitionQuestionnaireDataType            `xml:"urn:com.workday/bsvc Questionnaire_Reference,omitempty"`
	AssessmentData                   *JobRequisitionAssessmentDataType               `xml:"urn:com.workday/bsvc Assessment_Data,omitempty"`
	SupervisoryOrganizationReference *SupervisoryOrganizationObjectType              `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference,omitempty"`
	IntegrationFieldOverrideData     []DocumentFieldResultDataType                   `xml:"urn:com.workday/bsvc Integration_Field_Override_Data,omitempty"`
	OrganizationData                 []JobRequisitionOrganizationsDataType           `xml:"urn:com.workday/bsvc Organization_Data,omitempty"`
	RoleAssignmentData               []RoleAssignmentEffectiveDataType               `xml:"urn:com.workday/bsvc Role_Assignment_Data,omitempty"`
}

// Wrapper Element for a Job Requisition
type JobRequisitionDataType struct {
	JobRequisitionReference       *JobRequisitionObjectType                                `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	JobRequisitionID              string                                                   `xml:"urn:com.workday/bsvc Job_Requisition_ID,omitempty"`
	JobRequisitionStatusReference *JobRequisitionStatusObjectType                          `xml:"urn:com.workday/bsvc Job_Requisition_Status_Reference,omitempty"`
	JobPostingTitle               string                                                   `xml:"urn:com.workday/bsvc Job_Posting_Title,omitempty"`
	RecruitingInstructionData     *RecruitingInstructionDataType                           `xml:"urn:com.workday/bsvc Recruiting_Instruction_Data,omitempty"`
	AcademicTenureEligible        *bool                                                    `xml:"urn:com.workday/bsvc Academic_Tenure_Eligible,omitempty"`
	NumberofOpenings              float64                                                  `xml:"urn:com.workday/bsvc Number_of_Openings,omitempty"`
	JobDescriptionSummary         string                                                   `xml:"urn:com.workday/bsvc Job_Description_Summary,omitempty"`
	Justification                 string                                                   `xml:"urn:com.workday/bsvc Justification,omitempty"`
	JobRequisitionAttachments     *JobRequisitionAttachmentsType                           `xml:"urn:com.workday/bsvc Job_Requisition_Attachments,omitempty"`
	RecruitingStartDate           *time.Time                                               `xml:"urn:com.workday/bsvc Recruiting_Start_Date,omitempty"`
	TargetHireDate                *time.Time                                               `xml:"urn:com.workday/bsvc Target_Hire_Date,omitempty"`
	TargetEndDate                 *time.Time                                               `xml:"urn:com.workday/bsvc Target_End_Date,omitempty"`
	JobProfileReference           *JobProfileObjectType                                    `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	WorkerTypeReference           *WorkerTypeObjectType                                    `xml:"urn:com.workday/bsvc Worker_Type_Reference,omitempty"`
	PositionWorkerTypeReference   *PositionWorkerTypeObjectType                            `xml:"urn:com.workday/bsvc Position_Worker_Type_Reference,omitempty"`
	PrimaryLocationReference      *LocationObjectType                                      `xml:"urn:com.workday/bsvc Primary_Location_Reference,omitempty"`
	AdditionalLocationsReference  []LocationObjectType                                     `xml:"urn:com.workday/bsvc Additional_Locations_Reference,omitempty"`
	PositionTimeTypeReference     *PositionTimeTypeObjectType                              `xml:"urn:com.workday/bsvc Position_Time_Type_Reference,omitempty"`
	ScheduledWeeklyHours          float64                                                  `xml:"urn:com.workday/bsvc Scheduled_Weekly_Hours,omitempty"`
	WorkShiftReference            *WorkShiftObjectType                                     `xml:"urn:com.workday/bsvc Work_Shift_Reference,omitempty"`
	SpotlightJob                  *bool                                                    `xml:"urn:com.workday/bsvc Spotlight_Job,omitempty"`
	QualificationData             *QualificationDataforPositionRestrictionorJobProfileType `xml:"urn:com.workday/bsvc Qualification_Data,omitempty"`
	ReplacementforWorkerReference *WorkerObjectType                                        `xml:"urn:com.workday/bsvc Replacement_for_Worker_Reference,omitempty"`
}

func (t *JobRequisitionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobRequisitionDataType
	var layout struct {
		*T
		RecruitingStartDate *xsdDate `xml:"urn:com.workday/bsvc Recruiting_Start_Date,omitempty"`
		TargetHireDate      *xsdDate `xml:"urn:com.workday/bsvc Target_Hire_Date,omitempty"`
		TargetEndDate       *xsdDate `xml:"urn:com.workday/bsvc Target_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.RecruitingStartDate = (*xsdDate)(layout.T.RecruitingStartDate)
	layout.TargetHireDate = (*xsdDate)(layout.T.TargetHireDate)
	layout.TargetEndDate = (*xsdDate)(layout.T.TargetEndDate)
	return e.EncodeElement(layout, start)
}
func (t *JobRequisitionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobRequisitionDataType
	var overlay struct {
		*T
		RecruitingStartDate *xsdDate `xml:"urn:com.workday/bsvc Recruiting_Start_Date,omitempty"`
		TargetHireDate      *xsdDate `xml:"urn:com.workday/bsvc Target_Hire_Date,omitempty"`
		TargetEndDate       *xsdDate `xml:"urn:com.workday/bsvc Target_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.RecruitingStartDate = (*xsdDate)(overlay.T.RecruitingStartDate)
	overlay.TargetHireDate = (*xsdDate)(overlay.T.TargetHireDate)
	overlay.TargetEndDate = (*xsdDate)(overlay.T.TargetEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper for Job Requisition data common to create and edit requests.
type JobRequisitionDataforCreateandEditType struct {
	JobRequisitionID                      string                                                   `xml:"urn:com.workday/bsvc Job_Requisition_ID,omitempty"`
	JobPostingTitle                       string                                                   `xml:"urn:com.workday/bsvc Job_Posting_Title,omitempty"`
	RecruitingInstructionReference        *RecruitingInstructionObjectType                         `xml:"urn:com.workday/bsvc Recruiting_Instruction_Reference,omitempty"`
	JobDescriptionSummary                 string                                                   `xml:"urn:com.workday/bsvc Job_Description_Summary,omitempty"`
	JobDescription                        string                                                   `xml:"urn:com.workday/bsvc Job_Description,omitempty"`
	AdditionalJobDescription              string                                                   `xml:"urn:com.workday/bsvc Additional_Job_Description,omitempty"`
	Justification                         string                                                   `xml:"urn:com.workday/bsvc Justification,omitempty"`
	JobRequisitionAttachments             []JobRequisitionAttachmentsType                          `xml:"urn:com.workday/bsvc Job_Requisition_Attachments,omitempty"`
	RecruitingStartDate                   *time.Time                                               `xml:"urn:com.workday/bsvc Recruiting_Start_Date,omitempty"`
	TargetHireDate                        *time.Time                                               `xml:"urn:com.workday/bsvc Target_Hire_Date,omitempty"`
	TargetEndDate                         *time.Time                                               `xml:"urn:com.workday/bsvc Target_End_Date,omitempty"`
	JobProfileReference                   *JobProfileObjectType                                    `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	AdditionalJobProfilesReference        []JobProfileObjectType                                   `xml:"urn:com.workday/bsvc Additional_Job_Profiles_Reference,omitempty"`
	ReferralPaymentPlanReference          *OneTimePaymentPlanObjectType                            `xml:"urn:com.workday/bsvc Referral_Payment_Plan_Reference,omitempty"`
	WorkerTypeReference                   *WorkerTypeObjectType                                    `xml:"urn:com.workday/bsvc Worker_Type_Reference,omitempty"`
	WorkerSubTypeReference                *PositionWorkerTypeObjectType                            `xml:"urn:com.workday/bsvc Worker_Sub-Type_Reference,omitempty"`
	PrimaryLocationReference              *LocationObjectType                                      `xml:"urn:com.workday/bsvc Primary_Location_Reference,omitempty"`
	PrimaryJobPostingLocationReference    *LocationObjectType                                      `xml:"urn:com.workday/bsvc Primary_Job_Posting_Location_Reference,omitempty"`
	AdditionalLocationReference           []LocationObjectType                                     `xml:"urn:com.workday/bsvc Additional_Location_Reference,omitempty"`
	AdditionalJobPostingLocationReference []LocationObjectType                                     `xml:"urn:com.workday/bsvc Additional_Job_Posting_Location_Reference,omitempty"`
	PositionTimeTypeReference             *PositionTimeTypeObjectType                              `xml:"urn:com.workday/bsvc Position_Time_Type_Reference,omitempty"`
	ScheduledWeeklyHours                  float64                                                  `xml:"urn:com.workday/bsvc Scheduled_Weekly_Hours,omitempty"`
	WorkShiftReference                    *WorkShiftObjectType                                     `xml:"urn:com.workday/bsvc Work_Shift_Reference,omitempty"`
	SpotlightJob                          *bool                                                    `xml:"urn:com.workday/bsvc Spotlight_Job,omitempty"`
	ConfidentialJobRequisition            *bool                                                    `xml:"urn:com.workday/bsvc Confidential_Job_Requisition,omitempty"`
	LinktoEvergreenRequisitionReference   *EvergreenJobRequisitionObjectType                       `xml:"urn:com.workday/bsvc Link_to_Evergreen_Requisition_Reference,omitempty"`
	QuestionnaireData                     *JobRequisitionQuestionnaireDataType                     `xml:"urn:com.workday/bsvc Questionnaire_Data,omitempty"`
	AssessmentData                        *JobRequisitionAssessmentDataType                        `xml:"urn:com.workday/bsvc Assessment_Data,omitempty"`
	QualificationReplacementData          *QualificationDataforPositionRestrictionorJobProfileType `xml:"urn:com.workday/bsvc Qualification_Replacement_Data,omitempty"`
	ContingentWorkerCostInformation       *ContingentWorkerCostInformationDataType                 `xml:"urn:com.workday/bsvc Contingent_Worker_Cost_Information,omitempty"`
	ReplacementforWorkerReference         *WorkerObjectType                                        `xml:"urn:com.workday/bsvc Replacement_for_Worker_Reference,omitempty"`
	OrganizationData                      *JobRequisitionOrganizationsDataforBusinessProcessType   `xml:"urn:com.workday/bsvc Organization_Data,omitempty"`
	UseUpdatedTemplateVersion             *bool                                                    `xml:"urn:com.workday/bsvc Use_Updated_Template_Version,omitempty"`
	JobApplicationTemplateReference       *JobApplicationTemplateObjectType                        `xml:"urn:com.workday/bsvc Job_Application_Template_Reference,omitempty"`
}

func (t *JobRequisitionDataforCreateandEditType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobRequisitionDataforCreateandEditType
	var layout struct {
		*T
		RecruitingStartDate *xsdDate `xml:"urn:com.workday/bsvc Recruiting_Start_Date,omitempty"`
		TargetHireDate      *xsdDate `xml:"urn:com.workday/bsvc Target_Hire_Date,omitempty"`
		TargetEndDate       *xsdDate `xml:"urn:com.workday/bsvc Target_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.RecruitingStartDate = (*xsdDate)(layout.T.RecruitingStartDate)
	layout.TargetHireDate = (*xsdDate)(layout.T.TargetHireDate)
	layout.TargetEndDate = (*xsdDate)(layout.T.TargetEndDate)
	return e.EncodeElement(layout, start)
}
func (t *JobRequisitionDataforCreateandEditType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobRequisitionDataforCreateandEditType
	var overlay struct {
		*T
		RecruitingStartDate *xsdDate `xml:"urn:com.workday/bsvc Recruiting_Start_Date,omitempty"`
		TargetHireDate      *xsdDate `xml:"urn:com.workday/bsvc Target_Hire_Date,omitempty"`
		TargetEndDate       *xsdDate `xml:"urn:com.workday/bsvc Target_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.RecruitingStartDate = (*xsdDate)(overlay.T.RecruitingStartDate)
	overlay.TargetHireDate = (*xsdDate)(overlay.T.TargetHireDate)
	overlay.TargetEndDate = (*xsdDate)(overlay.T.TargetEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for Job Requisition Details
type JobRequisitionDefinitionDataType struct {
	JobPostingTitle                 string                            `xml:"urn:com.workday/bsvc Job_Posting_Title,omitempty"`
	PositionsAllocated              float64                           `xml:"urn:com.workday/bsvc Positions_Allocated,omitempty"`
	PositionsAvailable              float64                           `xml:"urn:com.workday/bsvc Positions_Available,omitempty"`
	AcademicTenureEligible          *bool                             `xml:"urn:com.workday/bsvc Academic_Tenure_Eligible,omitempty"`
	JobDescriptionSummary           string                            `xml:"urn:com.workday/bsvc Job_Description_Summary,omitempty"`
	JobDescription                  string                            `xml:"urn:com.workday/bsvc Job_Description,omitempty"`
	AdditionalJobDescription        string                            `xml:"urn:com.workday/bsvc Additional_Job_Description,omitempty"`
	Justification                   string                            `xml:"urn:com.workday/bsvc Justification,omitempty"`
	RecruitingInstructionData       *RecruitingInstructionDataType    `xml:"urn:com.workday/bsvc Recruiting_Instruction_Data,omitempty"`
	AvailableforRecruiting          *bool                             `xml:"urn:com.workday/bsvc Available_for_Recruiting,omitempty"`
	ReplacementforWorkerReference   *WorkerObjectType                 `xml:"urn:com.workday/bsvc Replacement_for_Worker_Reference,omitempty"`
	ConfidentialJobRequisition      *bool                             `xml:"urn:com.workday/bsvc Confidential_Job_Requisition,omitempty"`
	JobApplicationTemplateReference *JobApplicationTemplateObjectType `xml:"urn:com.workday/bsvc Job_Application_Template_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JobRequisitionEnabledObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobRequisitionEnabledObjectType struct {
	ID         []JobRequisitionEnabledObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JobRequisitionInterviewSessionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobRequisitionInterviewSessionObjectType struct {
	ID         []JobRequisitionInterviewSessionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains delete attribute, which marks a job requisition interview session for delete, a reference to a job requisition interview session, and job requisition interview session data.
type JobRequisitionInterviewSessionType struct {
	JobRequisitionInterviewSessionReference *JobRequisitionInterviewSessionObjectType    `xml:"urn:com.workday/bsvc Job_Requisition_Interview_Session_Reference,omitempty"`
	JobRequisitionInterviewSessionData      []JobRequisitionInterviewTeamSessionDataType `xml:"urn:com.workday/bsvc Job_Requisition_Interview_Session_Data,omitempty"`
	Delete                                  bool                                         `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains job requisition interview sessions data and the overall comment for the job requisition interview team.
type JobRequisitionInterviewTeamDataType struct {
	JobRequisitionInterviewTeamID  string                               `xml:"urn:com.workday/bsvc Job_Requisition_Interview_Team_ID,omitempty"`
	JobRequisitionInterviewSession []JobRequisitionInterviewSessionType `xml:"urn:com.workday/bsvc Job_Requisition_Interview_Session,omitempty"`
	OverallComment                 string                               `xml:"urn:com.workday/bsvc Overall_Comment,omitempty"`
}

// Wrapper element containing references to specific job requisitions.
type JobRequisitionInterviewTeamRequestReferencesType struct {
	JobRequisitionInterviewTeamReference []JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Interview_Team_Reference"`
}

// Contains job requisition interview team data.
type JobRequisitionInterviewTeamResponseDataType struct {
	JobRequisitionInterviewTeam []JobRequisitionInterviewTeamType `xml:"urn:com.workday/bsvc Job_Requisition_Interview_Team,omitempty"`
}

// Element containing job requisition interview team response group options.
type JobRequisitionInterviewTeamResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Contains the reference id, order, interviewers, duration, interview type, alternate interviewers, and comment data for the job requisition interview team session.
type JobRequisitionInterviewTeamSessionDataType struct {
	JobRequisitionInterviewSessionID string                    `xml:"urn:com.workday/bsvc Job_Requisition_Interview_Session_ID,omitempty"`
	Order                            string                    `xml:"urn:com.workday/bsvc Order,omitempty"`
	InterviewerReference             []WorkerObjectType        `xml:"urn:com.workday/bsvc Interviewer_Reference"`
	DurationinMinutes                float64                   `xml:"urn:com.workday/bsvc Duration_in_Minutes,omitempty"`
	InterviewTypeReference           *InterviewTypeObjectType  `xml:"urn:com.workday/bsvc Interview_Type_Reference,omitempty"`
	AlternateInterviewerReference    []WorkerObjectType        `xml:"urn:com.workday/bsvc Alternate_Interviewer_Reference,omitempty"`
	CompetenciesReference            []CompetencyObjectType    `xml:"urn:com.workday/bsvc Competencies_Reference,omitempty"`
	QuestionnairesReference          []QuestionnaireObjectType `xml:"urn:com.workday/bsvc Questionnaires_Reference,omitempty"`
	Comment                          string                    `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

// Contains the job requisition reference and job requisition interview team data.
type JobRequisitionInterviewTeamType struct {
	JobRequisitionInterviewTeamReference *JobRequisitionEnabledObjectType     `xml:"urn:com.workday/bsvc Job_Requisition_Interview_Team_Reference,omitempty"`
	JobRequisitionInterviewTeamData      *JobRequisitionInterviewTeamDataType `xml:"urn:com.workday/bsvc Job_Requisition_Interview_Team_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JobRequisitionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobRequisitionObjectType struct {
	ID         []JobRequisitionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Containing element for the cost center and custom organizations related to the job requisition.
type JobRequisitionOrganizationsDataType struct {
	OrganizationAssignmentsData []OrganizationAssignmentsforJobRequisitionDataType `xml:"urn:com.workday/bsvc Organization_Assignments_Data,omitempty"`
}

// Containing element for the cost center and custom organizations related to the job requisition.
type JobRequisitionOrganizationsDataforBusinessProcessType struct {
	OrganizationAssignmentsforJobRequisitionData *OrganizationAssignmentsforJobRequisitionDataType `xml:"urn:com.workday/bsvc Organization_Assignments_for_Job_Requisition_Data,omitempty"`
	ReplaceAll                                   bool                                              `xml:"urn:com.workday/bsvc Replace_All,attr,omitempty"`
	Delete                                       bool                                              `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Questionnaires to be used by Job Requisition for Internal and External Career Site Job Postings. For use by Recruiting customers only.
type JobRequisitionQuestionnaireDataType struct {
	QuestionnaireforInternalCareerSiteReference          *QuestionnaireObjectType `xml:"urn:com.workday/bsvc Questionnaire_for_Internal_Career_Site_Reference,omitempty"`
	SecondaryQuestionnaireforInternalCareerSiteReference *QuestionnaireObjectType `xml:"urn:com.workday/bsvc Secondary_Questionnaire_for_Internal_Career_Site_Reference,omitempty"`
	QuestionnaireforExternalCareerSitesReference         *QuestionnaireObjectType `xml:"urn:com.workday/bsvc Questionnaire_for_External_Career_Sites_Reference,omitempty"`
	SecondaryQuestionnaireforExternalCareerSiteReference *QuestionnaireObjectType `xml:"urn:com.workday/bsvc Secondary_Questionnaire_for_External_Career_Site_Reference,omitempty"`
}

// Utilize the Request Criteria element to filter instance(s) of Job Requisitions by status or supervisory org.
type JobRequisitionRequestCriteriaType struct {
	TransactionLogCriteriaData       []TransactionLogCriteriaType        `xml:"urn:com.workday/bsvc Transaction_Log_Criteria_Data,omitempty"`
	JobRequisitionStatusReference    []JobRequisitionStatusObjectType    `xml:"urn:com.workday/bsvc Job_Requisition_Status_Reference,omitempty"`
	SupervisoryOrganizationReference []SupervisoryOrganizationObjectType `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference,omitempty"`
	FieldAndParameterCriteriaData    *FieldAndParameterCriteriaDataType  `xml:"urn:com.workday/bsvc Field_And_Parameter_Criteria_Data,omitempty"`
	PrimaryLocationReference         []LocationObjectType                `xml:"urn:com.workday/bsvc Primary_Location_Reference,omitempty"`
	AdditionalLocationsReference     []LocationObjectType                `xml:"urn:com.workday/bsvc Additional_Locations_Reference,omitempty"`
}

// Utilize the Request References element to retrieve a specific instance(s) of Job Requisition and its associated data.
type JobRequisitionRequestReferencesType struct {
	JobRequisitionReference  []JobRequisitionObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
	SkipNonExistingInstances bool                       `xml:"urn:com.workday/bsvc Skip_Non_Existing_Instances,attr,omitempty"`
}

// Contains each Job Requisition based on the Request References or Request Criteria.
type JobRequisitionResponseDataType struct {
	JobRequisition []JobRequisitionType `xml:"urn:com.workday/bsvc Job_Requisition,omitempty"`
}

// The response group allows for the response data to be tailored to only included elements that the user is looking for.  If no response group is provided in the request then only the following elements will be returned:  Reference, Job Requisition Definition Data, and Job Requisition Restrictions Data.
type JobRequisitionResponseGroupType struct {
	IncludeReference                      *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeJobRequisitionDefinitionData   *bool `xml:"urn:com.workday/bsvc Include_Job_Requisition_Definition_Data,omitempty"`
	IncludeJobRequisitionRestrictionsData *bool `xml:"urn:com.workday/bsvc Include_Job_Requisition_Restrictions_Data,omitempty"`
	IncludeQualifications                 *bool `xml:"urn:com.workday/bsvc Include_Qualifications,omitempty"`
	IncludeJobRequisitionAttachments      *bool `xml:"urn:com.workday/bsvc Include_Job_Requisition_Attachments,omitempty"`
	IncludeOrganizations                  *bool `xml:"urn:com.workday/bsvc Include_Organizations,omitempty"`
	IncludeRoles                          *bool `xml:"urn:com.workday/bsvc Include_Roles,omitempty"`
}

// Wrapper element for Hiring Requirements
type JobRequisitionRestrictionsDataType struct {
	RecruitingStartDate                   *time.Time                    `xml:"urn:com.workday/bsvc Recruiting_Start_Date,omitempty"`
	TargetHireDate                        *time.Time                    `xml:"urn:com.workday/bsvc Target_Hire_Date,omitempty"`
	TargetEndDate                         *time.Time                    `xml:"urn:com.workday/bsvc Target_End_Date,omitempty"`
	JobProfileReference                   *JobProfileObjectType         `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	AdditionalJobProfilesReference        []JobProfileObjectType        `xml:"urn:com.workday/bsvc Additional_Job_Profiles_Reference,omitempty"`
	WorkerTypeReference                   *WorkerTypeObjectType         `xml:"urn:com.workday/bsvc Worker_Type_Reference,omitempty"`
	PositionWorkerTypeReference           *PositionWorkerTypeObjectType `xml:"urn:com.workday/bsvc Position_Worker_Type_Reference,omitempty"`
	PrimaryLocationReference              *LocationObjectType           `xml:"urn:com.workday/bsvc Primary_Location_Reference,omitempty"`
	PrimaryJobPostingLocationReference    *LocationObjectType           `xml:"urn:com.workday/bsvc Primary_Job_Posting_Location_Reference,omitempty"`
	AdditionalLocationsReference          []LocationObjectType          `xml:"urn:com.workday/bsvc Additional_Locations_Reference,omitempty"`
	AdditionalJobPostingLocationReference []LocationObjectType          `xml:"urn:com.workday/bsvc Additional_Job_Posting_Location_Reference,omitempty"`
	TimeTypeReference                     *PositionTimeTypeObjectType   `xml:"urn:com.workday/bsvc Time_Type_Reference,omitempty"`
	ScheduledWeeklyHours                  float64                       `xml:"urn:com.workday/bsvc Scheduled_Weekly_Hours,omitempty"`
	WorkShiftReference                    *WorkShiftObjectType          `xml:"urn:com.workday/bsvc Work_Shift_Reference,omitempty"`
	SpotlightJob                          *bool                         `xml:"urn:com.workday/bsvc Spotlight_Job,omitempty"`
	ReferralPaymentPlanReference          *OneTimePaymentPlanObjectType `xml:"urn:com.workday/bsvc Referral_Payment_Plan_Reference,omitempty"`
}

func (t *JobRequisitionRestrictionsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobRequisitionRestrictionsDataType
	var layout struct {
		*T
		RecruitingStartDate *xsdDate `xml:"urn:com.workday/bsvc Recruiting_Start_Date,omitempty"`
		TargetHireDate      *xsdDate `xml:"urn:com.workday/bsvc Target_Hire_Date,omitempty"`
		TargetEndDate       *xsdDate `xml:"urn:com.workday/bsvc Target_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.RecruitingStartDate = (*xsdDate)(layout.T.RecruitingStartDate)
	layout.TargetHireDate = (*xsdDate)(layout.T.TargetHireDate)
	layout.TargetEndDate = (*xsdDate)(layout.T.TargetEndDate)
	return e.EncodeElement(layout, start)
}
func (t *JobRequisitionRestrictionsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobRequisitionRestrictionsDataType
	var overlay struct {
		*T
		RecruitingStartDate *xsdDate `xml:"urn:com.workday/bsvc Recruiting_Start_Date,omitempty"`
		TargetHireDate      *xsdDate `xml:"urn:com.workday/bsvc Target_Hire_Date,omitempty"`
		TargetEndDate       *xsdDate `xml:"urn:com.workday/bsvc Target_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.RecruitingStartDate = (*xsdDate)(overlay.T.RecruitingStartDate)
	overlay.TargetHireDate = (*xsdDate)(overlay.T.TargetHireDate)
	overlay.TargetEndDate = (*xsdDate)(overlay.T.TargetEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type JobRequisitionStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobRequisitionStatusObjectType struct {
	ID         []JobRequisitionStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains each Job Requisition based on the Request References or Request Criteria.
type JobRequisitionType struct {
	JobRequisitionReference *JobRequisitionObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	JobRequisitionData      *JobRequisitionData2Type  `xml:"urn:com.workday/bsvc Job_Requisition_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LGBTIdentificationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LGBTIdentificationObjectType struct {
	ID         []LGBTIdentificationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the language ability information.
type LanguageAbilityDataType struct {
	LanguageProficiencyReference *LanguageProficiencyObjectType `xml:"urn:com.workday/bsvc Language_Proficiency_Reference,omitempty"`
	LanguageAbilityTypeReference *LanguageAbilityTypeObjectType `xml:"urn:com.workday/bsvc Language_Ability_Type_Reference"`
}

// Element containing the language ability type and proficiency for a Job Profile Language Qualfication language.
type LanguageAbilityProfileDataType struct {
	LanguageAbilityTypeReference *LanguageAbilityTypeObjectType `xml:"urn:com.workday/bsvc Language_Ability_Type_Reference"`
	LanguageProficiencyReference *LanguageProficiencyObjectType `xml:"urn:com.workday/bsvc Language_Proficiency_Reference,omitempty"`
}

// Wrapper element for Language Ability
type LanguageAbilityType struct {
	LanguageAbilityData []LanguageAbilityDataType `xml:"urn:com.workday/bsvc Language_Ability_Data"`
}

// Contains a unique identifier for an instance of an object.
type LanguageAbilityTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LanguageAbilityTypeObjectType struct {
	ID         []LanguageAbilityTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for Language Achievement information
type LanguageAchievementDataType struct {
	LanguageReference         *LanguageObjectType   `xml:"urn:com.workday/bsvc Language_Reference"`
	RemoveLanguage            *bool                 `xml:"urn:com.workday/bsvc Remove_Language,omitempty"`
	NativeLanguage            *bool                 `xml:"urn:com.workday/bsvc Native_Language,omitempty"`
	LanguageAbility           []LanguageAbilityType `xml:"urn:com.workday/bsvc Language_Ability,omitempty"`
	AssessedOn                *time.Time            `xml:"urn:com.workday/bsvc Assessed_On,omitempty"`
	Note                      string                `xml:"urn:com.workday/bsvc Note,omitempty"`
	AssessedbyWorkerReference *RoleObjectType       `xml:"urn:com.workday/bsvc Assessed_by_Worker_Reference,omitempty"`
}

func (t *LanguageAchievementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LanguageAchievementDataType
	var layout struct {
		*T
		AssessedOn *xsdDate `xml:"urn:com.workday/bsvc Assessed_On,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssessedOn = (*xsdDate)(layout.T.AssessedOn)
	return e.EncodeElement(layout, start)
}
func (t *LanguageAchievementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LanguageAchievementDataType
	var overlay struct {
		*T
		AssessedOn *xsdDate `xml:"urn:com.workday/bsvc Assessed_On,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssessedOn = (*xsdDate)(overlay.T.AssessedOn)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type LanguageObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LanguageObjectType struct {
	ID         []LanguageObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LanguageProficiencyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LanguageProficiencyObjectType struct {
	ID         []LanguageProficiencyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the language profile's information for a position.
type LanguageProfileforJobDataType struct {
	LanguageReference            *LanguageObjectType                  `xml:"urn:com.workday/bsvc Language_Reference,omitempty"`
	LanguageAbilityData          []LanguageAbilityDataType            `xml:"urn:com.workday/bsvc Language_Ability_Data,omitempty"`
	Required                     *bool                                `xml:"urn:com.workday/bsvc Required,omitempty"`
	QualificationSourceReference *SkillQualificationEnabledObjectType `xml:"urn:com.workday/bsvc Qualification_Source_Reference,omitempty"`
}

// Contains the position's language information.
type LanguageProfileforJobType struct {
	LanguageProfileReference *UniqueIdentifierObjectType    `xml:"urn:com.workday/bsvc Language_Profile_Reference,omitempty"`
	LanguageProfileData      *LanguageProfileforJobDataType `xml:"urn:com.workday/bsvc Language_Profile_Data,omitempty"`
}

// Replacement element containing Language Qualifications for the Job Profile. When updating a Job Profile, all Languages for the Job Profile will be replaced by the submitted data. If no data is submitted, then the existing Languages are not changed.
type LanguageQualificationProfileReplacementDataType struct {
	LanguageReference          *LanguageObjectType              `xml:"urn:com.workday/bsvc Language_Reference"`
	LanguageAbilityProfileData []LanguageAbilityProfileDataType `xml:"urn:com.workday/bsvc Language_Ability_Profile_Data,omitempty"`
	Required                   *bool                            `xml:"urn:com.workday/bsvc Required,omitempty"`
}

// Wrapper element for Language Qualifications. Allows all language qualifications for a Job Profile or Position Restriction to be removed - or to replace all existing language qualifications with those sent in the web service.
type LanguageQualificationReplacementType struct {
	LanguageQualificationReplacementData []LanguageQualificationProfileReplacementDataType `xml:"urn:com.workday/bsvc Language_Qualification_Replacement_Data,omitempty"`
	Delete                               bool                                              `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains the legal name for a person.  A person must name one and only one legal name.
type LegalNameDataType struct {
	NameDetailData *PersonNameDetailDataType `xml:"urn:com.workday/bsvc Name_Detail_Data"`
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
type LocationContextObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type LocationContextObjectType struct {
	ID         []LocationContextObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Wrapper element for the freezing/unfreezing of a job requisition.
type ManageJobRequisitionFreezeDataType struct {
	JobRequisitionReference       *JobRequisitionEnabledObjectType          `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
	FreezeUnfreezeReasonReference *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Freeze_Unfreeze_Reason_Reference"`
	FreezeUnfreezeDate            time.Time                                 `xml:"urn:com.workday/bsvc Freeze_Unfreeze_Date"`
	Frozen                        *bool                                     `xml:"urn:com.workday/bsvc Frozen,omitempty"`
	UnpostJobSubBusinessProcess   *UnpostJobSubBusinessProcessType          `xml:"urn:com.workday/bsvc Unpost_Job_Sub_Business_Process,omitempty"`
	PostJobSubProcess             *PostJobSubBusinessProcessType            `xml:"urn:com.workday/bsvc Post_Job_Sub_Process,omitempty"`
}

func (t *ManageJobRequisitionFreezeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ManageJobRequisitionFreezeDataType
	var layout struct {
		*T
		FreezeUnfreezeDate *xsdDate `xml:"urn:com.workday/bsvc Freeze_Unfreeze_Date"`
	}
	layout.T = (*T)(t)
	layout.FreezeUnfreezeDate = (*xsdDate)(&layout.T.FreezeUnfreezeDate)
	return e.EncodeElement(layout, start)
}
func (t *ManageJobRequisitionFreezeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ManageJobRequisitionFreezeDataType
	var overlay struct {
		*T
		FreezeUnfreezeDate *xsdDate `xml:"urn:com.workday/bsvc Freeze_Unfreeze_Date"`
	}
	overlay.T = (*T)(t)
	overlay.FreezeUnfreezeDate = (*xsdDate)(&overlay.T.FreezeUnfreezeDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper for Manage Job Requisition Freeze Web Service
type ManageJobRequisitionFreezeRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType      `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	ManageJobRequisitionData  *ManageJobRequisitionFreezeDataType `xml:"urn:com.workday/bsvc Manage_Job_Requisition_Data"`
	Version                   string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the event id for the Freeze Job Requisition Event, the Job Requisition Reference, and the frozen/unfrozen state of the job requisition.
type ManageJobRequisitionFreezeResponseType struct {
	EventReference          *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	JobRequisitionReference *JobRequisitionObjectType   `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	Frozen                  *bool                       `xml:"urn:com.workday/bsvc Frozen,omitempty"`
	Version                 string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ManagementLevelObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ManagementLevelObjectType struct {
	ID         []ManagementLevelObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MaritalStatusObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type MaritalStatusObjectType struct {
	ID         []MaritalStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MeritPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MeritPlanObjectType struct {
	ID         []MeritPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MilitaryRankObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MilitaryRankObjectType struct {
	ID         []MilitaryRankObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MilitaryServiceReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MilitaryServiceReferenceObjectType struct {
	ID         []MilitaryServiceReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Container for military service data for the Change Personal Information business process.
type MilitaryServiceSubDataType struct {
	MilitaryStatusReference      *MilitaryStatusObjectType      `xml:"urn:com.workday/bsvc Military_Status_Reference"`
	MilitaryDischargeDate        *time.Time                     `xml:"urn:com.workday/bsvc Military_Discharge_Date,omitempty"`
	MilitaryStatusBeginDate      *time.Time                     `xml:"urn:com.workday/bsvc Military_Status_Begin_Date,omitempty"`
	MilitaryServiceTypeReference *MilitaryServiceTypeObjectType `xml:"urn:com.workday/bsvc Military_Service_Type_Reference,omitempty"`
	MilitaryRankReference        *MilitaryRankObjectType        `xml:"urn:com.workday/bsvc Military_Rank_Reference,omitempty"`
	Notes                        string                         `xml:"urn:com.workday/bsvc Notes,omitempty"`
}

func (t *MilitaryServiceSubDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MilitaryServiceSubDataType
	var layout struct {
		*T
		MilitaryDischargeDate   *xsdDate `xml:"urn:com.workday/bsvc Military_Discharge_Date,omitempty"`
		MilitaryStatusBeginDate *xsdDate `xml:"urn:com.workday/bsvc Military_Status_Begin_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.MilitaryDischargeDate = (*xsdDate)(layout.T.MilitaryDischargeDate)
	layout.MilitaryStatusBeginDate = (*xsdDate)(layout.T.MilitaryStatusBeginDate)
	return e.EncodeElement(layout, start)
}
func (t *MilitaryServiceSubDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MilitaryServiceSubDataType
	var overlay struct {
		*T
		MilitaryDischargeDate   *xsdDate `xml:"urn:com.workday/bsvc Military_Discharge_Date,omitempty"`
		MilitaryStatusBeginDate *xsdDate `xml:"urn:com.workday/bsvc Military_Status_Begin_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.MilitaryDischargeDate = (*xsdDate)(overlay.T.MilitaryDischargeDate)
	overlay.MilitaryStatusBeginDate = (*xsdDate)(overlay.T.MilitaryStatusBeginDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type MilitaryServiceTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MilitaryServiceTypeObjectType struct {
	ID         []MilitaryServiceTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MilitaryStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MilitaryStatusObjectType struct {
	ID         []MilitaryStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type MonthObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MonthObjectType struct {
	ID         []MonthObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Collects the Candidate reference and Job Requisition reference
type MoveCandidateCriteriaType struct {
	CandidateReference      *CandidateObjectType             `xml:"urn:com.workday/bsvc Candidate_Reference"`
	JobRequisitionReference *JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
}

// Reference to the Candidate Job Application to be moved from a stage to next possible stage or to a disposition stage.
type MoveCandidateDataType struct {
	EventReference          *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference"`
	JobApplicationReference *JobApplicationObjectType   `xml:"urn:com.workday/bsvc Job_Application_Reference"`
	CandidateCriteriaData   *MoveCandidateCriteriaType  `xml:"urn:com.workday/bsvc Candidate_Criteria_Data"`
}

// A request to move a Candidate from a stage to next possible stage or to a disposition stage.
type MoveCandidateRequestType struct {
	DynamicBusinessProcessParameters *DynamicBusinessProcessParametersType `xml:"urn:com.workday/bsvc Dynamic_Business_Process_Parameters,omitempty"`
	MoveCandidateData                *MoveCandidateDataType                `xml:"urn:com.workday/bsvc Move_Candidate_Data"`
	Version                          string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The result of an attempt to move a Candidate from a stage to next possible stage or to a disposition stage.
type MoveCandidateResponseType struct {
	EventReference *RecruitingEventObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to move a candidate from a job requisition to a linked ~Evergreen Requistion~
type MoveCandidatetoLinkedEvergreenRequisitionRequestType struct {
	EvergreenRequisitionReference *EvergreenJobRequisitionObjectType `xml:"urn:com.workday/bsvc Evergreen_Requisition_Reference"`
	JobRequisitionReference       *JobRequisitionObjectType          `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
	JobApplicationReference       []JobApplicationObjectType         `xml:"urn:com.workday/bsvc Job_Application_Reference"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response element for the Move Candidate to Linked Evergreen Requisition Task.
type MoveCandidatetoLinkedEvergreenRequisitionResponseType struct {
	EvergreenRequisitionReference *EvergreenJobRequisitionObjectType `xml:"urn:com.workday/bsvc Evergreen_Requisition_Reference"`
	JobRequisitionReference       *JobRequisitionObjectType          `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
	JobApplicationReference       []JobApplicationObjectType         `xml:"urn:com.workday/bsvc Job_Application_Reference"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The request element for the Link To Job Requisition Web Service Task.
type MoveCandidatetoLinkedJobRequisitionRequestType struct {
	JobRequisitionReference       *JobRequisitionObjectType          `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
	EvergreenRequisitionReference *EvergreenJobRequisitionObjectType `xml:"urn:com.workday/bsvc Evergreen_Requisition_Reference"`
	JobApplicationReference       []JobApplicationObjectType         `xml:"urn:com.workday/bsvc Job_Application_Reference"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response of the Move To Linked Job Requisition Web Service Task.
type MoveCandidatetoLinkedJobRequisitionResponseType struct {
	JobRequisitionReference       *JobRequisitionObjectType          `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	EvergreenRequisitionReference *EvergreenJobRequisitionObjectType `xml:"urn:com.workday/bsvc Evergreen_Requisition_Reference,omitempty"`
	JobApplicationReference       []JobApplicationObjectType         `xml:"urn:com.workday/bsvc Job_Application_Reference,omitempty"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Specify the Job Requisition to be moved and the Job Requisition's new Supervisory Organization.
type MoveJobRequisitionDataType struct {
	JobRequisitiontobeMovedReference    *JobRequisitionObjectType          `xml:"urn:com.workday/bsvc Job_Requisition_to_be_Moved_Reference"`
	NewSupervisoryOrganizationReference *SupervisoryOrganizationObjectType `xml:"urn:com.workday/bsvc New_Supervisory_Organization_Reference"`
}

// Move a job requisition from one Job Management Organization to another Job Management Organization using the Move Job Requisition business process.
type MoveJobRequisitionRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	MoveJobRequisitionData    *MoveJobRequisitionDataType    `xml:"urn:com.workday/bsvc Move_Job_Requisition_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response to a Move Job Requisition Request. The response contains the ID of a Move Job Requisition Event if the business process was initiated successfully.
type MoveJobRequisitionResponseType struct {
	MoveJobRequisitionReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Move_Job_Requisition_Reference,omitempty"`
	Version                     string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains multiple choice answer data.
type MultipleChoiceAnswerDataType struct {
	MultipleChoiceAnswerReference *QuestionMultipleChoiceAnswerObjectType `xml:"urn:com.workday/bsvc Multiple_Choice_Answer_Reference"`
	MultipleChoiceAnswerText      string                                  `xml:"urn:com.workday/bsvc Multiple_Choice_Answer_Text"`
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

// Non Effective Dated Web Service Additional Data
type NonEffectiveDatedWebServiceAdditionalDataType []string

func (a NonEffectiveDatedWebServiceAdditionalDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ArrayType string   `xml:"http://schemas.xmlsoap.org/wsdl/ arrayType,attr"`
		Items     []string `xml:" item"`
	}
	output.Items = []string(a)
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{"", "xmlns:ns1"}, Value: "http://www.w3.org/2001/XMLSchema"})
	output.ArrayType = "ns1:anyType[]"
	return e.EncodeElement(&output, start)
}
func (a *NonEffectiveDatedWebServiceAdditionalDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var tok xml.Token
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if tok, ok := tok.(xml.StartElement); ok {
			var item string
			if err = d.DecodeElement(&item, &tok); err == nil {
				*a = append(*a, item)
			}
		}
		if _, ok := tok.(xml.EndElement); ok {
			break
		}
	}
	return err
}

// Contains the details of the offer and compensation
type OfferDataType struct {
	JobApplicationEventReference    *UniqueIdentifierObjectType                             `xml:"urn:com.workday/bsvc Job_Application_Event_Reference"`
	JobApplicationReference         *JobApplicationObjectType                               `xml:"urn:com.workday/bsvc Job_Application_Reference"`
	CandidateCriteriaData           *MoveCandidateCriteriaType                              `xml:"urn:com.workday/bsvc Candidate_Criteria_Data"`
	OfferEventData                  *OfferEventDataType                                     `xml:"urn:com.workday/bsvc Offer_Event_Data"`
	ProposeCompensationSubProcess   *ProposeCompensationForEmploymentSubBusinessProcessType `xml:"urn:com.workday/bsvc Propose_Compensation_Sub_Process,omitempty"`
	RequestOneTimePaymentSubProcess *RequestOneTimePaymentSubBusinessProcessType            `xml:"urn:com.workday/bsvc Request_One_Time_Payment_Sub_Process,omitempty"`
	RequestStockGrantSubProcess     *RequestStockOfferSubBusinessProcessType                `xml:"urn:com.workday/bsvc Request_Stock_Grant_Sub_Process,omitempty"`
}

// Contains offer details
type OfferEventDataType struct {
	HireDate                  time.Time                                 `xml:"urn:com.workday/bsvc Hire_Date"`
	HireReasonReference       *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Hire_Reason_Reference,omitempty"`
	LocationReference         *LocationObjectType                       `xml:"urn:com.workday/bsvc Location_Reference"`
	ProposedEndDate           *time.Time                                `xml:"urn:com.workday/bsvc Proposed_End_Date,omitempty"`
	DocumentLanguageReference *UserLanguageObjectType                   `xml:"urn:com.workday/bsvc Document_Language_Reference"`
	JobProfileReference       *JobProfileObjectType                     `xml:"urn:com.workday/bsvc Job_Profile_Reference"`
	BusinessTitle             string                                    `xml:"urn:com.workday/bsvc Business_Title"`
	DefaultWeeklyHours        float64                                   `xml:"urn:com.workday/bsvc Default_Weekly_Hours,omitempty"`
	ScheduledWeeklyHours      float64                                   `xml:"urn:com.workday/bsvc Scheduled_Weekly_Hours,omitempty"`
}

func (t *OfferEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T OfferEventDataType
	var layout struct {
		*T
		HireDate        *xsdDate `xml:"urn:com.workday/bsvc Hire_Date"`
		ProposedEndDate *xsdDate `xml:"urn:com.workday/bsvc Proposed_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.HireDate = (*xsdDate)(&layout.T.HireDate)
	layout.ProposedEndDate = (*xsdDate)(layout.T.ProposedEndDate)
	return e.EncodeElement(layout, start)
}
func (t *OfferEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T OfferEventDataType
	var overlay struct {
		*T
		HireDate        *xsdDate `xml:"urn:com.workday/bsvc Hire_Date"`
		ProposedEndDate *xsdDate `xml:"urn:com.workday/bsvc Proposed_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.HireDate = (*xsdDate)(&overlay.T.HireDate)
	overlay.ProposedEndDate = (*xsdDate)(overlay.T.ProposedEndDate)
	return d.DecodeElement(&overlay, &start)
}

// A request to initiate offer
type OfferRequestType struct {
	DynamicBusinessProcessParameters *DynamicBusinessProcessParametersType `xml:"urn:com.workday/bsvc Dynamic_Business_Process_Parameters,omitempty"`
	OfferData                        *OfferDataType                        `xml:"urn:com.workday/bsvc Offer_Data"`
	Version                          string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for Offer response.
type OfferResponseType struct {
	OfferEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Offer_Event_Reference,omitempty"`
	Version             string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Data Element that contains the one-time payment information.
type OneTimePaymentDataType struct {
	OneTimePaymentPlanReference     *OneTimePaymentPlanObjectType      `xml:"urn:com.workday/bsvc One_Time_Payment_Plan_Reference"`
	ScheduledPaymentDate            *time.Time                         `xml:"urn:com.workday/bsvc Scheduled_Payment_Date,omitempty"`
	CoverageStartDate               *time.Time                         `xml:"urn:com.workday/bsvc Coverage_Start_Date,omitempty"`
	CoverageEndDate                 *time.Time                         `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	Amount                          float64                            `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Percent                         float64                            `xml:"urn:com.workday/bsvc Percent,omitempty"`
	CurrencyReference               *CurrencyObjectType                `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	Comment                         string                             `xml:"urn:com.workday/bsvc Comment,omitempty"`
	DoNotPay                        *bool                              `xml:"urn:com.workday/bsvc Do_Not_Pay,omitempty"`
	OneTimePaymentWorktagsReference []TenantedPayrollWorktagObjectType `xml:"urn:com.workday/bsvc One_Time_Payment_Worktags_Reference,omitempty"`
}

func (t *OneTimePaymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T OneTimePaymentDataType
	var layout struct {
		*T
		ScheduledPaymentDate *xsdDate `xml:"urn:com.workday/bsvc Scheduled_Payment_Date,omitempty"`
		CoverageStartDate    *xsdDate `xml:"urn:com.workday/bsvc Coverage_Start_Date,omitempty"`
		CoverageEndDate      *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ScheduledPaymentDate = (*xsdDate)(layout.T.ScheduledPaymentDate)
	layout.CoverageStartDate = (*xsdDate)(layout.T.CoverageStartDate)
	layout.CoverageEndDate = (*xsdDate)(layout.T.CoverageEndDate)
	return e.EncodeElement(layout, start)
}
func (t *OneTimePaymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T OneTimePaymentDataType
	var overlay struct {
		*T
		ScheduledPaymentDate *xsdDate `xml:"urn:com.workday/bsvc Scheduled_Payment_Date,omitempty"`
		CoverageStartDate    *xsdDate `xml:"urn:com.workday/bsvc Coverage_Start_Date,omitempty"`
		CoverageEndDate      *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ScheduledPaymentDate = (*xsdDate)(overlay.T.ScheduledPaymentDate)
	overlay.CoverageStartDate = (*xsdDate)(overlay.T.CoverageStartDate)
	overlay.CoverageEndDate = (*xsdDate)(overlay.T.CoverageEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type OneTimePaymentPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type OneTimePaymentPlanObjectType struct {
	ID         []OneTimePaymentPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for the default organization assignments for this organization.
type OrganizationAssignmentsDataType struct {
	OrganizationTypeReference *OrganizationTypeObjectType `xml:"urn:com.workday/bsvc Organization_Type_Reference"`
	OrganizationReference     *OrganizationObjectType     `xml:"urn:com.workday/bsvc Organization_Reference"`
}

// Wrapper element for the default organization assignments for this job requisition.
type OrganizationAssignmentsforJobRequisitionDataType struct {
	CompanyAssignmentReference            *CompanyObjectType             `xml:"urn:com.workday/bsvc Company_Assignment_Reference,omitempty"`
	CostCenterAssignmentReference         *CostCenterObjectType          `xml:"urn:com.workday/bsvc Cost_Center_Assignment_Reference,omitempty"`
	RegionAssignmentReference             *RegionObjectType              `xml:"urn:com.workday/bsvc Region_Assignment_Reference,omitempty"`
	BusinessUnitAssignmentReference       *BusinessUnitObjectType        `xml:"urn:com.workday/bsvc Business_Unit_Assignment_Reference,omitempty"`
	GrantAssignmentReference              *GrantObjectType               `xml:"urn:com.workday/bsvc Grant_Assignment_Reference,omitempty"`
	ProgramAssignmentReference            *ProgramObjectType             `xml:"urn:com.workday/bsvc Program_Assignment_Reference,omitempty"`
	FundAssignmentReference               *FundObjectType                `xml:"urn:com.workday/bsvc Fund_Assignment_Reference,omitempty"`
	GiftAssignmentReference               *GiftObjectType                `xml:"urn:com.workday/bsvc Gift_Assignment_Reference,omitempty"`
	CustomOrganizationAssignmentReference []CustomOrganizationObjectType `xml:"urn:com.workday/bsvc Custom_Organization_Assignment_Reference,omitempty"`
}

// Contains information regarding organizations just adjacent in the hierarchy.
type OrganizationHierarchyDataType struct {
	TopLevelOrganizationReference    *OrganizationObjectType  `xml:"urn:com.workday/bsvc Top-Level_Organization_Reference,omitempty"`
	SuperiorOrganizationReference    *OrganizationObjectType  `xml:"urn:com.workday/bsvc Superior_Organization_Reference,omitempty"`
	SubordinateOrganizationReference []OrganizationObjectType `xml:"urn:com.workday/bsvc Subordinate_Organization_Reference,omitempty"`
	IncludedOrganizationReference    []OrganizationObjectType `xml:"urn:com.workday/bsvc Included_Organization_Reference,omitempty"`
	IncludedInOrganizationReference  []OrganizationObjectType `xml:"urn:com.workday/bsvc Included_In_Organization_Reference,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type OrganizationOwnerObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type OrganizationOwnerObjectType struct {
	ID         []OrganizationOwnerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This element contains criteria to filter the Organizations returned.
type OrganizationRequestCriteriaType struct {
	OrganizationTypeReference     []OrganizationTypeObjectType       `xml:"urn:com.workday/bsvc Organization_Type_Reference,omitempty"`
	IncludeInactive               *bool                              `xml:"urn:com.workday/bsvc Include_Inactive,omitempty"`
	TransactionLogCriteria        []TransactionLogCriteriaType       `xml:"urn:com.workday/bsvc Transaction_Log_Criteria,omitempty"`
	FieldAndParameterCriteriaData *FieldAndParameterCriteriaDataType `xml:"urn:com.workday/bsvc Field_And_Parameter_Criteria_Data,omitempty"`
}

// References to specific Organizations to return
type OrganizationRequestReferencesType struct {
	OrganizationReference    []OrganizationObjectType `xml:"urn:com.workday/bsvc Organization_Reference"`
	SkipNonExistingInstances bool                     `xml:"urn:com.workday/bsvc Skip_Non_Existing_Instances,attr,omitempty"`
}

// Response element containing an instance of Organization and its associated data.
type OrganizationResponseDataType struct {
	Organization []OrganizationWWSType `xml:"urn:com.workday/bsvc Organization,omitempty"`
}

// The response group allows for the response data to be tailored to only included elements that the user is looking for.  If no response group is provided in the request, then only the following elements will be returned:  Reference, Organization Data, Hierarchy Data and if a Supervisory Organization Type the Supervisory Data.
type OrganizationResponseGroupType struct {
	IncludeRolesData                *bool `xml:"urn:com.workday/bsvc Include_Roles_Data,omitempty"`
	IncludeHierarchyData            *bool `xml:"urn:com.workday/bsvc Include_Hierarchy_Data,omitempty"`
	IncludeSupervisoryData          *bool `xml:"urn:com.workday/bsvc Include_Supervisory_Data,omitempty"`
	IncludeStaffingRestrictionsData *bool `xml:"urn:com.workday/bsvc Include_Staffing_Restrictions_Data,omitempty"`
}

// Wrapper element for a Organization Role Assignment
type OrganizationRoleAssignmentDataType struct {
	RoleAssignerReference            *RoleAssignerObjectType   `xml:"urn:com.workday/bsvc Role_Assigner_Reference,omitempty"`
	OrganizationRoleReference        *AssignableRoleObjectType `xml:"urn:com.workday/bsvc Organization_Role_Reference"`
	RoleAssigneeReference            []RoleeObjectType         `xml:"urn:com.workday/bsvc Role_Assignee_Reference,omitempty"`
	SingleAssignmentManagerReference *RoleeObjectType          `xml:"urn:com.workday/bsvc Single_Assignment_Manager_Reference,omitempty"`
}

// Contains information about organization role assignments.
type OrganizationRoleAssignmentWWSDataType struct {
	RoleReference   *AssignableRoleObjectType `xml:"urn:com.workday/bsvc Role_Reference,omitempty"`
	WorkerReference []WorkerObjectType        `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
}

// Contains information about role assignments for an organization.
type OrganizationRolesWWSDataType struct {
	OrganizationRoleData []OrganizationRoleAssignmentWWSDataType `xml:"urn:com.workday/bsvc Organization_Role_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type OrganizationSubtypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type OrganizationSubtypeObjectType struct {
	ID         []OrganizationSubtypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains the detailed information about an Organization.
type OrganizationWWSDataType struct {
	ReferenceID                     string                         `xml:"urn:com.workday/bsvc Reference_ID,omitempty"`
	Name                            string                         `xml:"urn:com.workday/bsvc Name,omitempty"`
	Description                     string                         `xml:"urn:com.workday/bsvc Description,omitempty"`
	OrganizationCode                string                         `xml:"urn:com.workday/bsvc Organization_Code,omitempty"`
	IncludeManagerinName            *bool                          `xml:"urn:com.workday/bsvc Include_Manager_in_Name,omitempty"`
	IncludeOrganizationCodeinName   *bool                          `xml:"urn:com.workday/bsvc Include_Organization_Code_in_Name,omitempty"`
	OrganizationTypeReference       *OrganizationTypeObjectType    `xml:"urn:com.workday/bsvc Organization_Type_Reference,omitempty"`
	OrganizationSubtypeReference    *OrganizationSubtypeObjectType `xml:"urn:com.workday/bsvc Organization_Subtype_Reference,omitempty"`
	AvailibilityDate                *time.Time                     `xml:"urn:com.workday/bsvc Availibility_Date,omitempty"`
	LastUpdatedDateTime             *time.Time                     `xml:"urn:com.workday/bsvc Last_Updated_DateTime,omitempty"`
	Inactive                        *bool                          `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	InactiveDate                    *time.Time                     `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
	ManagerReference                *WorkerObjectType              `xml:"urn:com.workday/bsvc Manager_Reference,omitempty"`
	LeadershipReference             []WorkerObjectType             `xml:"urn:com.workday/bsvc Leadership_Reference,omitempty"`
	OrganizationOwnerReference      *OrganizationOwnerObjectType   `xml:"urn:com.workday/bsvc Organization_Owner_Reference,omitempty"`
	OrganizationVisibilityReference *UniqueIdentifierObjectType    `xml:"urn:com.workday/bsvc Organization_Visibility_Reference,omitempty"`
	ExternalURLReference            *ExternalURLObjectType         `xml:"urn:com.workday/bsvc External_URL_Reference,omitempty"`
	ExternalIDsData                 *ExternalIntegrationIDDataType `xml:"urn:com.workday/bsvc External_IDs_Data,omitempty"`
	RolesData                       *OrganizationRolesWWSDataType  `xml:"urn:com.workday/bsvc Roles_Data,omitempty"`
	HierarchyData                   *OrganizationHierarchyDataType `xml:"urn:com.workday/bsvc Hierarchy_Data,omitempty"`
	SupervisoryData                 *SupervisoryOrgDataType        `xml:"urn:com.workday/bsvc Supervisory_Data,omitempty"`
	IntegrationFieldOverrideData    []DocumentFieldResultDataType  `xml:"urn:com.workday/bsvc Integration_Field_Override_Data,omitempty"`
}

func (t *OrganizationWWSDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T OrganizationWWSDataType
	var layout struct {
		*T
		AvailibilityDate    *xsdDateTime `xml:"urn:com.workday/bsvc Availibility_Date,omitempty"`
		LastUpdatedDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Last_Updated_DateTime,omitempty"`
		InactiveDate        *xsdDate     `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AvailibilityDate = (*xsdDateTime)(layout.T.AvailibilityDate)
	layout.LastUpdatedDateTime = (*xsdDateTime)(layout.T.LastUpdatedDateTime)
	layout.InactiveDate = (*xsdDate)(layout.T.InactiveDate)
	return e.EncodeElement(layout, start)
}
func (t *OrganizationWWSDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T OrganizationWWSDataType
	var overlay struct {
		*T
		AvailibilityDate    *xsdDateTime `xml:"urn:com.workday/bsvc Availibility_Date,omitempty"`
		LastUpdatedDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Last_Updated_DateTime,omitempty"`
		InactiveDate        *xsdDate     `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AvailibilityDate = (*xsdDateTime)(overlay.T.AvailibilityDate)
	overlay.LastUpdatedDateTime = (*xsdDateTime)(overlay.T.LastUpdatedDateTime)
	overlay.InactiveDate = (*xsdDate)(overlay.T.InactiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Response element containing an instance of Organization and its associated data.
type OrganizationWWSType struct {
	OrganizationReference *OrganizationObjectType  `xml:"urn:com.workday/bsvc Organization_Reference,omitempty"`
	OrganizationData      *OrganizationWWSDataType `xml:"urn:com.workday/bsvc Organization_Data,omitempty"`
}

// Details of the Background Check Package.
type PackageReferenceDataType struct {
	PackageReference   *BackgroundCheckPackageObjectType `xml:"urn:com.workday/bsvc Package_Reference,omitempty"`
	StatusReference    *BackgroundCheckStatusObjectType  `xml:"urn:com.workday/bsvc Status_Reference"`
	ResultsURL         string                            `xml:"urn:com.workday/bsvc Results_URL,omitempty"`
	PackageComment     string                            `xml:"urn:com.workday/bsvc Package_Comment,omitempty"`
	PackageName        string                            `xml:"urn:com.workday/bsvc Package_Name,attr,omitempty"`
	PackageDescription string                            `xml:"urn:com.workday/bsvc Package_Description,attr,omitempty"`
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
type PeriodSalaryPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PeriodSalaryPlanObjectType struct {
	ID         []PeriodSalaryPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains Candidate Personal Info: Gender, Date of Birth, Country of Birth, Place of Birth Region, Place of Birth City,  Ethnicity, Social Benefits Locality, Marital Status,  Hispanic/Latino, Military Service, Religion,  Citizenship Status ,  Primary Nationality, Additional Nationalities,  Relative Name, Disabilities, Sexual Orientation Gender Identity
type PersonBiographicandDemographicInformationDataType struct {
	GenderReference                 *GenderObjectType                             `xml:"urn:com.workday/bsvc Gender_Reference,omitempty"`
	DateofBirth                     *time.Time                                    `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	BirthCountryReference           *CountryObjectType                            `xml:"urn:com.workday/bsvc Birth_Country_Reference,omitempty"`
	BirthRegionReference            *CountryRegionObjectType                      `xml:"urn:com.workday/bsvc Birth_Region_Reference,omitempty"`
	CityofBirth                     string                                        `xml:"urn:com.workday/bsvc City_of_Birth,omitempty"`
	MaritalStatusReference          *MaritalStatusObjectType                      `xml:"urn:com.workday/bsvc Marital_Status_Reference,omitempty"`
	HispanicorLatino                *bool                                         `xml:"urn:com.workday/bsvc Hispanic_or_Latino,omitempty"`
	EthnicityReference              []EthnicityObjectType                         `xml:"urn:com.workday/bsvc Ethnicity_Reference,omitempty"`
	ReligionReference               *ReligionObjectType                           `xml:"urn:com.workday/bsvc Religion_Reference,omitempty"`
	CitizenshipReference            []CitizenshipStatusObjectType                 `xml:"urn:com.workday/bsvc Citizenship_Reference,omitempty"`
	NationalityReference            *CountryObjectType                            `xml:"urn:com.workday/bsvc Nationality_Reference,omitempty"`
	AdditionalNationalityReference  []CountryObjectType                           `xml:"urn:com.workday/bsvc Additional_Nationality_Reference,omitempty"`
	SocialBenefitsLocalityReference *SocialBenefitsLocalityObjectType             `xml:"urn:com.workday/bsvc Social_Benefits_Locality_Reference,omitempty"`
	LGBTIdentificationReference     []LGBTIdentificationObjectType                `xml:"urn:com.workday/bsvc LGBT_Identification_Reference,omitempty"`
	SexualOrientationReference      *SexualOrientationObjectType                  `xml:"urn:com.workday/bsvc Sexual_Orientation_Reference,omitempty"`
	GenderIdentityReference         *GenderIdentityObjectType                     `xml:"urn:com.workday/bsvc Gender_Identity_Reference,omitempty"`
	PronounReference                *PronounObjectType                            `xml:"urn:com.workday/bsvc Pronoun_Reference,omitempty"`
	DisabilityStatusData            []CandidateDisabilityStatusDataType           `xml:"urn:com.workday/bsvc Disability_Status_Data,omitempty"`
	MilitaryServiceInformationData  []CandidateMilitaryServiceInformationDataType `xml:"urn:com.workday/bsvc Military_Service_Information_Data,omitempty"`
}

func (t *PersonBiographicandDemographicInformationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PersonBiographicandDemographicInformationDataType
	var layout struct {
		*T
		DateofBirth *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateofBirth = (*xsdDate)(layout.T.DateofBirth)
	return e.EncodeElement(layout, start)
}
func (t *PersonBiographicandDemographicInformationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PersonBiographicandDemographicInformationDataType
	var overlay struct {
		*T
		DateofBirth *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateofBirth = (*xsdDate)(overlay.T.DateofBirth)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for the disability status information for the person.
type PersonDisabilityStatusDataType struct {
	DisabilityReference                       *DisabilityObjectType                       `xml:"urn:com.workday/bsvc Disability_Reference"`
	DisabilityStatusDate                      *time.Time                                  `xml:"urn:com.workday/bsvc Disability_Status_Date,omitempty"`
	DisabilityDateKnown                       *time.Time                                  `xml:"urn:com.workday/bsvc Disability_Date_Known,omitempty"`
	DisabilityEndDate                         *time.Time                                  `xml:"urn:com.workday/bsvc Disability_End_Date,omitempty"`
	DisabilityGradeReference                  *DisabilityGradeObjectType                  `xml:"urn:com.workday/bsvc Disability_Grade_Reference,omitempty"`
	DisabilityDegree                          float64                                     `xml:"urn:com.workday/bsvc Disability_Degree,omitempty"`
	DisabilityRemainingCapacity               float64                                     `xml:"urn:com.workday/bsvc Disability_Remaining_Capacity,omitempty"`
	DisabilityCertificationAuthorityReference *DisabilityCertificationAuthorityObjectType `xml:"urn:com.workday/bsvc Disability_Certification_Authority_Reference,omitempty"`
	DisabilityCertificationAuthority          string                                      `xml:"urn:com.workday/bsvc Disability_Certification_Authority,omitempty"`
	DisabilityCertifiedAt                     string                                      `xml:"urn:com.workday/bsvc Disability_Certified_At,omitempty"`
	DisabilityCertificationID                 string                                      `xml:"urn:com.workday/bsvc Disability_Certification_ID,omitempty"`
	DisabilityCertificationBasisReference     *DisabilityCertificationBasisObjectType     `xml:"urn:com.workday/bsvc Disability_Certification_Basis_Reference,omitempty"`
	DisabilitySeverityRecognitionDate         *time.Time                                  `xml:"urn:com.workday/bsvc Disability_Severity_Recognition_Date,omitempty"`
	DisabilityFTETowardQuota                  float64                                     `xml:"urn:com.workday/bsvc Disability_FTE_Toward_Quota,omitempty"`
	DisabilityWorkRestrictions                string                                      `xml:"urn:com.workday/bsvc Disability_Work_Restrictions,omitempty"`
	DisabilityAccommodationsRequested         string                                      `xml:"urn:com.workday/bsvc Disability_Accommodations_Requested,omitempty"`
	DisabilityAccommodationsProvided          string                                      `xml:"urn:com.workday/bsvc Disability_Accommodations_Provided,omitempty"`
	DisabilityRehabilitationRequested         string                                      `xml:"urn:com.workday/bsvc Disability_Rehabilitation_Requested,omitempty"`
	DisabilityRehabilitationProvided          string                                      `xml:"urn:com.workday/bsvc Disability_Rehabilitation_Provided,omitempty"`
	Note                                      string                                      `xml:"urn:com.workday/bsvc Note,omitempty"`
	WorkerDocumentReference                   []WorkerDocumentObjectType                  `xml:"urn:com.workday/bsvc Worker_Document_Reference,omitempty"`
	DisabilityStatusReference                 *DisabilityStatusReferenceObjectType        `xml:"urn:com.workday/bsvc Disability_Status_Reference,omitempty"`
}

func (t *PersonDisabilityStatusDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PersonDisabilityStatusDataType
	var layout struct {
		*T
		DisabilityStatusDate              *xsdDate `xml:"urn:com.workday/bsvc Disability_Status_Date,omitempty"`
		DisabilityDateKnown               *xsdDate `xml:"urn:com.workday/bsvc Disability_Date_Known,omitempty"`
		DisabilityEndDate                 *xsdDate `xml:"urn:com.workday/bsvc Disability_End_Date,omitempty"`
		DisabilitySeverityRecognitionDate *xsdDate `xml:"urn:com.workday/bsvc Disability_Severity_Recognition_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DisabilityStatusDate = (*xsdDate)(layout.T.DisabilityStatusDate)
	layout.DisabilityDateKnown = (*xsdDate)(layout.T.DisabilityDateKnown)
	layout.DisabilityEndDate = (*xsdDate)(layout.T.DisabilityEndDate)
	layout.DisabilitySeverityRecognitionDate = (*xsdDate)(layout.T.DisabilitySeverityRecognitionDate)
	return e.EncodeElement(layout, start)
}
func (t *PersonDisabilityStatusDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PersonDisabilityStatusDataType
	var overlay struct {
		*T
		DisabilityStatusDate              *xsdDate `xml:"urn:com.workday/bsvc Disability_Status_Date,omitempty"`
		DisabilityDateKnown               *xsdDate `xml:"urn:com.workday/bsvc Disability_Date_Known,omitempty"`
		DisabilityEndDate                 *xsdDate `xml:"urn:com.workday/bsvc Disability_End_Date,omitempty"`
		DisabilitySeverityRecognitionDate *xsdDate `xml:"urn:com.workday/bsvc Disability_Severity_Recognition_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DisabilityStatusDate = (*xsdDate)(overlay.T.DisabilityStatusDate)
	overlay.DisabilityDateKnown = (*xsdDate)(overlay.T.DisabilityDateKnown)
	overlay.DisabilityEndDate = (*xsdDate)(overlay.T.DisabilityEndDate)
	overlay.DisabilitySeverityRecognitionDate = (*xsdDate)(overlay.T.DisabilitySeverityRecognitionDate)
	return d.DecodeElement(&overlay, &start)
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

// Wrapper element for the military service information for the person.
type PersonMilitaryServiceDataType struct {
	StatusReference              *MilitaryStatusObjectType           `xml:"urn:com.workday/bsvc Status_Reference"`
	DischargeDate                *time.Time                          `xml:"urn:com.workday/bsvc Discharge_Date,omitempty"`
	StatusBeginDate              *time.Time                          `xml:"urn:com.workday/bsvc Status_Begin_Date,omitempty"`
	MilitaryServiceTypeReference *MilitaryServiceTypeObjectType      `xml:"urn:com.workday/bsvc Military_Service_Type_Reference,omitempty"`
	MilitaryRankReference        *MilitaryRankObjectType             `xml:"urn:com.workday/bsvc Military_Rank_Reference,omitempty"`
	Notes                        string                              `xml:"urn:com.workday/bsvc Notes,omitempty"`
	MilitaryServiceReference     *MilitaryServiceReferenceObjectType `xml:"urn:com.workday/bsvc Military_Service_Reference,omitempty"`
}

func (t *PersonMilitaryServiceDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PersonMilitaryServiceDataType
	var layout struct {
		*T
		DischargeDate   *xsdDate `xml:"urn:com.workday/bsvc Discharge_Date,omitempty"`
		StatusBeginDate *xsdDate `xml:"urn:com.workday/bsvc Status_Begin_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DischargeDate = (*xsdDate)(layout.T.DischargeDate)
	layout.StatusBeginDate = (*xsdDate)(layout.T.StatusBeginDate)
	return e.EncodeElement(layout, start)
}
func (t *PersonMilitaryServiceDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PersonMilitaryServiceDataType
	var overlay struct {
		*T
		DischargeDate   *xsdDate `xml:"urn:com.workday/bsvc Discharge_Date,omitempty"`
		StatusBeginDate *xsdDate `xml:"urn:com.workday/bsvc Status_Begin_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DischargeDate = (*xsdDate)(overlay.T.DischargeDate)
	overlay.StatusBeginDate = (*xsdDate)(overlay.T.StatusBeginDate)
	return d.DecodeElement(&overlay, &start)
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

// Contains the worker's skills and experience.
//
// Security Note: This element is secured to the following domains:  Self Service: Skills and Experience; Worker: Skills and Experience
type PersonQualificationDataType struct {
	ExternalJobHistory        []JobHistoryType                   `xml:"urn:com.workday/bsvc External_Job_History,omitempty"`
	Competency                []CompetencyAchievementDatav10Type `xml:"urn:com.workday/bsvc Competency,omitempty"`
	Certification             []CertificationAchievementType     `xml:"urn:com.workday/bsvc Certification,omitempty"`
	Training                  []TrainingType                     `xml:"urn:com.workday/bsvc Training,omitempty"`
	AwardandActivity          []AwardandActivityType             `xml:"urn:com.workday/bsvc Award_and_Activity,omitempty"`
	OrganizationMembership    []ProfessionalAffiliationSkillType `xml:"urn:com.workday/bsvc Organization_Membership,omitempty"`
	Education                 []EducationType                    `xml:"urn:com.workday/bsvc Education,omitempty"`
	WorkExperience            []WorkExperienceDataType           `xml:"urn:com.workday/bsvc Work_Experience,omitempty"`
	Language                  []LanguageAchievementDataType      `xml:"urn:com.workday/bsvc Language,omitempty"`
	InternalProjectExperience []InternalProjectExperienceType    `xml:"urn:com.workday/bsvc Internal_Project_Experience,omitempty"`
}

// Wrapper element for Personal Data.
type PersonalInformationDataType struct {
	UniversalID                     string                            `xml:"urn:com.workday/bsvc Universal_ID,omitempty"`
	NameData                        *PersonNameDataType               `xml:"urn:com.workday/bsvc Name_Data,omitempty"`
	GenderReference                 *GenderObjectType                 `xml:"urn:com.workday/bsvc Gender_Reference,omitempty"`
	BirthDate                       *time.Time                        `xml:"urn:com.workday/bsvc Birth_Date,omitempty"`
	DateofDeath                     *time.Time                        `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
	CountryofBirthReference         *CountryObjectType                `xml:"urn:com.workday/bsvc Country_of_Birth_Reference,omitempty"`
	RegionofBirthReference          *CountryRegionObjectType          `xml:"urn:com.workday/bsvc Region_of_Birth_Reference,omitempty"`
	RegionofBirthDescriptor         string                            `xml:"urn:com.workday/bsvc Region_of_Birth_Descriptor,omitempty"`
	CityofBirth                     string                            `xml:"urn:com.workday/bsvc City_of_Birth,omitempty"`
	CityofBirthReference            *CountryCityObjectType            `xml:"urn:com.workday/bsvc City_of_Birth_Reference,omitempty"`
	MaritalStatusReference          *MaritalStatusObjectType          `xml:"urn:com.workday/bsvc Marital_Status_Reference,omitempty"`
	MaritalStatusDate               *time.Time                        `xml:"urn:com.workday/bsvc Marital_Status_Date,omitempty"`
	ReligionReference               *ReligionObjectType               `xml:"urn:com.workday/bsvc Religion_Reference,omitempty"`
	DisabilityStatusData            []PersonDisabilityStatusDataType  `xml:"urn:com.workday/bsvc Disability_Status_Data,omitempty"`
	EthnicityReference              []EthnicityObjectType             `xml:"urn:com.workday/bsvc Ethnicity_Reference,omitempty"`
	HispanicorLatino                *bool                             `xml:"urn:com.workday/bsvc Hispanic_or_Latino,omitempty"`
	CitizenshipStatusReference      []CitizenshipStatusObjectType     `xml:"urn:com.workday/bsvc Citizenship_Status_Reference,omitempty"`
	PrimaryNationalityReference     *CountryObjectType                `xml:"urn:com.workday/bsvc Primary_Nationality_Reference,omitempty"`
	AdditionalNationalityReference  []CountryObjectType               `xml:"urn:com.workday/bsvc Additional_Nationality_Reference,omitempty"`
	HukouRegionReference            *CountryRegionObjectType          `xml:"urn:com.workday/bsvc Hukou_Region_Reference,omitempty"`
	HukouSubregionReference         *CountrySubregionObjectType       `xml:"urn:com.workday/bsvc Hukou_Subregion_Reference,omitempty"`
	HukouLocality                   string                            `xml:"urn:com.workday/bsvc Hukou_Locality,omitempty"`
	HukouPostalCode                 string                            `xml:"urn:com.workday/bsvc Hukou_Postal_Code,omitempty"`
	HukouTypeReference              *HukouTypeObjectType              `xml:"urn:com.workday/bsvc Hukou_Type_Reference,omitempty"`
	LocalHukou                      *bool                             `xml:"urn:com.workday/bsvc Local_Hukou,omitempty"`
	NativeRegionReference           *CountryRegionObjectType          `xml:"urn:com.workday/bsvc Native_Region_Reference,omitempty"`
	NativeRegionDescriptor          string                            `xml:"urn:com.workday/bsvc Native_Region_Descriptor,omitempty"`
	PersonnelFileAgencyforPerson    string                            `xml:"urn:com.workday/bsvc Personnel_File_Agency_for_Person,omitempty"`
	LastMedicalExamDate             *time.Time                        `xml:"urn:com.workday/bsvc Last_Medical_Exam_Date,omitempty"`
	LastMedicalExamValidTo          *time.Time                        `xml:"urn:com.workday/bsvc Last_Medical_Exam_Valid_To,omitempty"`
	MedicalExamNotes                string                            `xml:"urn:com.workday/bsvc Medical_Exam_Notes,omitempty"`
	BloodTypeReference              *BloodTypeObjectType              `xml:"urn:com.workday/bsvc Blood_Type_Reference,omitempty"`
	MilitaryServiceData             []PersonMilitaryServiceDataType   `xml:"urn:com.workday/bsvc Military_Service_Data,omitempty"`
	IdentificationData              *PersonIdentificationDataType     `xml:"urn:com.workday/bsvc Identification_Data,omitempty"`
	ContactData                     *ContactInformationDataType       `xml:"urn:com.workday/bsvc Contact_Data,omitempty"`
	TobaccoUse                      *bool                             `xml:"urn:com.workday/bsvc Tobacco_Use,omitempty"`
	PoliticalAffiliationReference   *PoliticalAffiliationObjectType   `xml:"urn:com.workday/bsvc Political_Affiliation_Reference,omitempty"`
	SocialBenefitsLocalityReference *SocialBenefitsLocalityObjectType `xml:"urn:com.workday/bsvc Social_Benefits_Locality_Reference,omitempty"`
	RelativeNameInformationData     *RelativeNameInformationDataType  `xml:"urn:com.workday/bsvc Relative_Name_Information_Data,omitempty"`
}

func (t *PersonalInformationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PersonalInformationDataType
	var layout struct {
		*T
		BirthDate              *xsdDate `xml:"urn:com.workday/bsvc Birth_Date,omitempty"`
		DateofDeath            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
		MaritalStatusDate      *xsdDate `xml:"urn:com.workday/bsvc Marital_Status_Date,omitempty"`
		LastMedicalExamDate    *xsdDate `xml:"urn:com.workday/bsvc Last_Medical_Exam_Date,omitempty"`
		LastMedicalExamValidTo *xsdDate `xml:"urn:com.workday/bsvc Last_Medical_Exam_Valid_To,omitempty"`
	}
	layout.T = (*T)(t)
	layout.BirthDate = (*xsdDate)(layout.T.BirthDate)
	layout.DateofDeath = (*xsdDate)(layout.T.DateofDeath)
	layout.MaritalStatusDate = (*xsdDate)(layout.T.MaritalStatusDate)
	layout.LastMedicalExamDate = (*xsdDate)(layout.T.LastMedicalExamDate)
	layout.LastMedicalExamValidTo = (*xsdDate)(layout.T.LastMedicalExamValidTo)
	return e.EncodeElement(layout, start)
}
func (t *PersonalInformationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PersonalInformationDataType
	var overlay struct {
		*T
		BirthDate              *xsdDate `xml:"urn:com.workday/bsvc Birth_Date,omitempty"`
		DateofDeath            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
		MaritalStatusDate      *xsdDate `xml:"urn:com.workday/bsvc Marital_Status_Date,omitempty"`
		LastMedicalExamDate    *xsdDate `xml:"urn:com.workday/bsvc Last_Medical_Exam_Date,omitempty"`
		LastMedicalExamValidTo *xsdDate `xml:"urn:com.workday/bsvc Last_Medical_Exam_Valid_To,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.BirthDate = (*xsdDate)(overlay.T.BirthDate)
	overlay.DateofDeath = (*xsdDate)(overlay.T.DateofDeath)
	overlay.MaritalStatusDate = (*xsdDate)(overlay.T.MaritalStatusDate)
	overlay.LastMedicalExamDate = (*xsdDate)(overlay.T.LastMedicalExamDate)
	overlay.LastMedicalExamValidTo = (*xsdDate)(overlay.T.LastMedicalExamValidTo)
	return d.DecodeElement(&overlay, &start)
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
type PoliticalAffiliationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PoliticalAffiliationObjectType struct {
	ID         []PoliticalAffiliationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the detailed information about a Position.
type PositionDataType struct {
	SupervisoryOrganizationReference           *SupervisoryOrganizationObjectType              `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference"`
	EffectiveDate                              *time.Time                                      `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	PositionDefinitionData                     *PositionDetailsDataType                        `xml:"urn:com.workday/bsvc Position_Definition_Data,omitempty"`
	PositionRestrictionsData                   *PositionGroupRestrictionSummaryDataType        `xml:"urn:com.workday/bsvc Position_Restrictions_Data,omitempty"`
	DefaultCompensationData                    *CompensationDefaultDataType                    `xml:"urn:com.workday/bsvc Default_Compensation_Data,omitempty"`
	DefaultPositionOrganizationAssignmentsData *DefaultPositionOrganizationAssignmentsDataType `xml:"urn:com.workday/bsvc Default_Position_Organization_Assignments_Data,omitempty"`
	WorkerForFilledPosition                    *WorkerForFilledPositionType                    `xml:"urn:com.workday/bsvc Worker_For_Filled_Position,omitempty"`
	QualificationData                          *QualificationsfromPositionRestrictionsDataType `xml:"urn:com.workday/bsvc Qualification_Data,omitempty"`
	IntegrationFieldOverrideData               []DocumentFieldResultDataType                   `xml:"urn:com.workday/bsvc Integration_Field_Override_Data,omitempty"`
	JobRequisitionData                         *JobRequisitionDataType                         `xml:"urn:com.workday/bsvc Job_Requisition_Data,omitempty"`
	Closed                                     *bool                                           `xml:"urn:com.workday/bsvc Closed,omitempty"`
}

func (t *PositionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PositionDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PositionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PositionDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for general data pertaining to a position opening.
type PositionDefinitionDataType struct {
	PositionID                string                      `xml:"urn:com.workday/bsvc Position_ID,omitempty"`
	JobPostingTitle           string                      `xml:"urn:com.workday/bsvc Job_Posting_Title,omitempty"`
	JobDescriptionSummary     string                      `xml:"urn:com.workday/bsvc Job_Description_Summary,omitempty"`
	JobDescription            string                      `xml:"urn:com.workday/bsvc Job_Description,omitempty"`
	CriticalJob               *bool                       `xml:"urn:com.workday/bsvc Critical_Job,omitempty"`
	DifficultytoFillReference *DifficultytoFillObjectType `xml:"urn:com.workday/bsvc Difficulty_to_Fill_Reference,omitempty"`
}

// Wrapper element for general data pertaining to a position opening.
type PositionDetailsDataType struct {
	PositionID                string                      `xml:"urn:com.workday/bsvc Position_ID,omitempty"`
	JobPostingTitle           string                      `xml:"urn:com.workday/bsvc Job_Posting_Title,omitempty"`
	AcademicTenureEligible    *bool                       `xml:"urn:com.workday/bsvc Academic_Tenure_Eligible,omitempty"`
	JobDescriptionSummary     string                      `xml:"urn:com.workday/bsvc Job_Description_Summary,omitempty"`
	JobDescription            string                      `xml:"urn:com.workday/bsvc Job_Description,omitempty"`
	PositionStatusReference   []WorkerEventObjectType     `xml:"urn:com.workday/bsvc Position_Status_Reference,omitempty"`
	AvailableForHire          *bool                       `xml:"urn:com.workday/bsvc Available_For_Hire,omitempty"`
	AvailableforRecruiting    *bool                       `xml:"urn:com.workday/bsvc Available_for_Recruiting,omitempty"`
	HiringFreeze              *bool                       `xml:"urn:com.workday/bsvc Hiring_Freeze,omitempty"`
	WorkShiftRequired         *bool                       `xml:"urn:com.workday/bsvc Work_Shift_Required,omitempty"`
	AvailableforOverlap       *bool                       `xml:"urn:com.workday/bsvc Available_for_Overlap,omitempty"`
	EarliestOverlapDate       *time.Time                  `xml:"urn:com.workday/bsvc Earliest_Overlap_Date,omitempty"`
	CriticalJob               *bool                       `xml:"urn:com.workday/bsvc Critical_Job,omitempty"`
	DifficultytoFillReference *DifficultytoFillObjectType `xml:"urn:com.workday/bsvc Difficulty_to_Fill_Reference,omitempty"`
}

func (t *PositionDetailsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PositionDetailsDataType
	var layout struct {
		*T
		EarliestOverlapDate *xsdDate `xml:"urn:com.workday/bsvc Earliest_Overlap_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EarliestOverlapDate = (*xsdDate)(layout.T.EarliestOverlapDate)
	return e.EncodeElement(layout, start)
}
func (t *PositionDetailsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PositionDetailsDataType
	var overlay struct {
		*T
		EarliestOverlapDate *xsdDate `xml:"urn:com.workday/bsvc Earliest_Overlap_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EarliestOverlapDate = (*xsdDate)(overlay.T.EarliestOverlapDate)
	return d.DecodeElement(&overlay, &start)
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
type PositionGroupObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PositionGroupObjectType struct {
	ID         []PositionGroupObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper for restriction data for openings for all staffing models.
type PositionGroupRestrictionDataType struct {
	AvailabilityDate            *time.Time                     `xml:"urn:com.workday/bsvc Availability_Date,omitempty"`
	EarliestHireDate            *time.Time                     `xml:"urn:com.workday/bsvc Earliest_Hire_Date,omitempty"`
	JobFamilyReference          []JobFamilyBaseObjectType      `xml:"urn:com.workday/bsvc Job_Family_Reference,omitempty"`
	JobProfileReference         []JobProfileObjectType         `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	LocationReference           []LocationObjectType           `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	WorkerTypeReference         *WorkerTypeObjectType          `xml:"urn:com.workday/bsvc Worker_Type_Reference,omitempty"`
	TimeTypeReference           *PositionTimeTypeObjectType    `xml:"urn:com.workday/bsvc Time_Type_Reference,omitempty"`
	PositionWorkerTypeReference []PositionWorkerTypeObjectType `xml:"urn:com.workday/bsvc Position_Worker_Type_Reference,omitempty"`
}

func (t *PositionGroupRestrictionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PositionGroupRestrictionDataType
	var layout struct {
		*T
		AvailabilityDate *xsdDate `xml:"urn:com.workday/bsvc Availability_Date,omitempty"`
		EarliestHireDate *xsdDate `xml:"urn:com.workday/bsvc Earliest_Hire_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AvailabilityDate = (*xsdDate)(layout.T.AvailabilityDate)
	layout.EarliestHireDate = (*xsdDate)(layout.T.EarliestHireDate)
	return e.EncodeElement(layout, start)
}
func (t *PositionGroupRestrictionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PositionGroupRestrictionDataType
	var overlay struct {
		*T
		AvailabilityDate *xsdDate `xml:"urn:com.workday/bsvc Availability_Date,omitempty"`
		EarliestHireDate *xsdDate `xml:"urn:com.workday/bsvc Earliest_Hire_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AvailabilityDate = (*xsdDate)(overlay.T.AvailabilityDate)
	overlay.EarliestHireDate = (*xsdDate)(overlay.T.EarliestHireDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper for restriction data for openings for all staffing models.
type PositionGroupRestrictionSummaryDataType struct {
	AvailabilityDate                 *time.Time                     `xml:"urn:com.workday/bsvc Availability_Date,omitempty"`
	EarliestHireDate                 *time.Time                     `xml:"urn:com.workday/bsvc Earliest_Hire_Date,omitempty"`
	JobFamilyReference               []JobFamilyBaseObjectType      `xml:"urn:com.workday/bsvc Job_Family_Reference,omitempty"`
	JobProfileRestrictionSummaryData []JobProfileSummaryDataType    `xml:"urn:com.workday/bsvc Job_Profile_Restriction_Summary_Data,omitempty"`
	LocationReference                []LocationObjectType           `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	WorkerTypeReference              *WorkerTypeObjectType          `xml:"urn:com.workday/bsvc Worker_Type_Reference,omitempty"`
	TimeTypeReference                *PositionTimeTypeObjectType    `xml:"urn:com.workday/bsvc Time_Type_Reference,omitempty"`
	PositionWorkerTypeReference      []PositionWorkerTypeObjectType `xml:"urn:com.workday/bsvc Position_Worker_Type_Reference,omitempty"`
}

func (t *PositionGroupRestrictionSummaryDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PositionGroupRestrictionSummaryDataType
	var layout struct {
		*T
		AvailabilityDate *xsdDate `xml:"urn:com.workday/bsvc Availability_Date,omitempty"`
		EarliestHireDate *xsdDate `xml:"urn:com.workday/bsvc Earliest_Hire_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AvailabilityDate = (*xsdDate)(layout.T.AvailabilityDate)
	layout.EarliestHireDate = (*xsdDate)(layout.T.EarliestHireDate)
	return e.EncodeElement(layout, start)
}
func (t *PositionGroupRestrictionSummaryDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PositionGroupRestrictionSummaryDataType
	var overlay struct {
		*T
		AvailabilityDate *xsdDate `xml:"urn:com.workday/bsvc Availability_Date,omitempty"`
		EarliestHireDate *xsdDate `xml:"urn:com.workday/bsvc Earliest_Hire_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AvailabilityDate = (*xsdDate)(overlay.T.AvailabilityDate)
	overlay.EarliestHireDate = (*xsdDate)(overlay.T.EarliestHireDate)
	return d.DecodeElement(&overlay, &start)
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

// Utilize the Request Criteria element to search the transaction log for specific instance(s) of a Position.
type PositionRequestCriteriaType struct {
	TransactionLogCriteriaData       []TransactionLogCriteriaType        `xml:"urn:com.workday/bsvc Transaction_Log_Criteria_Data,omitempty"`
	SupervisoryOrganizationReference []SupervisoryOrganizationObjectType `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference,omitempty"`
	JobRequisitionReference          []JobRequisitionObjectType          `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	ExcludeFilledPositions           *bool                               `xml:"urn:com.workday/bsvc Exclude_Filled_Positions,omitempty"`
	ExcludeNonRecruitablePositions   *bool                               `xml:"urn:com.workday/bsvc Exclude_Non-Recruitable_Positions,omitempty"`
	IgnoreFuturePositions            *bool                               `xml:"urn:com.workday/bsvc Ignore_Future_Positions,omitempty"`
	IncludeClosedPositions           *bool                               `xml:"urn:com.workday/bsvc Include_Closed_Positions,omitempty"`
	ShowOnlyClosedPositions          *bool                               `xml:"urn:com.workday/bsvc Show_Only_Closed_Positions,omitempty"`
	EventReference                   *EventObjectType                    `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	FieldAndParameterCriteriaData    *FieldAndParameterCriteriaDataType  `xml:"urn:com.workday/bsvc Field_And_Parameter_Criteria_Data,omitempty"`
	LocationReference                []LocationObjectType                `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	AdditionalLocationsReference     []LocationObjectType                `xml:"urn:com.workday/bsvc Additional_Locations_Reference,omitempty"`
}

// Utilize the Request References element to retrieve a specific instance(s) of Position and its associated data.
type PositionRequestReferencesType struct {
	PositionsReference       []PositionElementObjectType `xml:"urn:com.workday/bsvc Positions_Reference"`
	SkipNonExistingInstances bool                        `xml:"urn:com.workday/bsvc Skip_Non_Existing_Instances,attr,omitempty"`
	IgnoreInvalidReferences  bool                        `xml:"urn:com.workday/bsvc Ignore_Invalid_References,attr,omitempty"`
}

// The response group allows for the response data to be tailored to only included elements that the user is looking for.  If no response group is provided in the request the only the following elements will be returned:  Reference and Position Definition Data
type PositionResponseGroupType struct {
	IncludeReference                                 *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludePositionDefinitionData                    *bool `xml:"urn:com.workday/bsvc Include_Position_Definition_Data,omitempty"`
	IncludePositionRestrictionsData                  *bool `xml:"urn:com.workday/bsvc Include_Position_Restrictions_Data,omitempty"`
	IncludeDefaultCompensationData                   *bool `xml:"urn:com.workday/bsvc Include_Default_Compensation_Data,omitempty"`
	IncludeDefaultPositionOrganizationAssignmentData *bool `xml:"urn:com.workday/bsvc Include_Default_Position_Organization_Assignment_Data,omitempty"`
	IncludeWorkerForFilledPositionsData              *bool `xml:"urn:com.workday/bsvc Include_Worker_For_Filled_Positions_Data,omitempty"`
	IncludeQualifications                            *bool `xml:"urn:com.workday/bsvc Include_Qualifications,omitempty"`
	IncludeJobRequisitionData                        *bool `xml:"urn:com.workday/bsvc Include_Job_Requisition_Data,omitempty"`
	IncludeJobRequisitionAttachments                 *bool `xml:"urn:com.workday/bsvc Include_Job_Requisition_Attachments,omitempty"`
}

// Contains each Position based on the Request References or Request Criteria.  The data returned is the current information as of the dates in the response filter, and does not include all historical information about the  position.
type PositionRestrictionType struct {
	PositionReference *PositionRestrictionsObjectType `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	PositionData      *PositionDataType               `xml:"urn:com.workday/bsvc Position_Data,omitempty"`
}

// Effective dated additional data for position restrictions.
type PositionRestrictionsCustomObjectDataType struct {
	EffectiveDate                 *time.Time                                   `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	PositionRestrictionsReference *PositionRestrictionsObjectType              `xml:"urn:com.workday/bsvc Position_Restrictions_Reference,omitempty"`
	BusinessObjectAdditionalData  []EffectiveDatedWebServiceAdditionalDataType `xml:"urn:com.workday/bsvc Business_Object_Additional_Data,omitempty"`
}

func (t *PositionRestrictionsCustomObjectDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PositionRestrictionsCustomObjectDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PositionRestrictionsCustomObjectDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PositionRestrictionsCustomObjectDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
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

// Contains a unique identifier for an instance of an object.
type PositionTimeTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PositionTimeTypeObjectType struct {
	ID         []PositionTimeTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PositionWorkerTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PositionWorkerTypeObjectType struct {
	ID         []PositionWorkerTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Data on positions related to the job requisition.
type PositionsDataType struct {
	PositionReference []PositionGroupObjectType `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
}

// Contains each Position based on the Request References or Request Criteria.  The data returned is the current information as of the dates in the response filter, and does not include all historical information about the  position.
type PositionsResponseDataType struct {
	Position []PositionRestrictionType `xml:"urn:com.workday/bsvc Position,omitempty"`
}

// Wrapper for Post Job data.
type PostJobBusinessProcessDataType struct {
	JobRequisitionReference *JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
	JobPostingSiteData      []JobPostingSiteDataType         `xml:"urn:com.workday/bsvc Job_Posting_Site_Data"`
}

// Post a job to a Job Posting Site.
type PostJobRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType  `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	PostJobData               *PostJobBusinessProcessDataType `xml:"urn:com.workday/bsvc Post_Job_Data"`
	Version                   string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Event ID for the Job Posting Event.
type PostJobResponseType struct {
	EventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	JobPostingData []JobPostingDataType        `xml:"urn:com.workday/bsvc Job_Posting_Data,omitempty"`
	Version        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Run the Post Job Sub Business Process to post this job requisition to one or more Job Posting Sites. To use, Post Job must be enabled in the Workflow Definition for Create Job Requisition.
type PostJobSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	PostJobData                  []JobPostingSiteDataType          `xml:"urn:com.workday/bsvc Post_Job_Data,omitempty"`
}

// The Wrapper element for the Applicant Data Web Service.
type PreHireDataWWSType struct {
	ApplicantID               string                                 `xml:"urn:com.workday/bsvc Applicant_ID,omitempty"`
	PersonalData              *PersonalInformationDataType           `xml:"urn:com.workday/bsvc Personal_Data,omitempty"`
	QualificationData         *PersonQualificationDataType           `xml:"urn:com.workday/bsvc Qualification_Data,omitempty"`
	RecruitingData            *RecruitingDataType                    `xml:"urn:com.workday/bsvc Recruiting_Data,omitempty"`
	ResumeData                *PreHireResumeDataType                 `xml:"urn:com.workday/bsvc Resume_Data,omitempty"`
	BackgroundCheckData       []BackgroundCheckOverallStatusDataType `xml:"urn:com.workday/bsvc Background_Check_Data,omitempty"`
	ExternalIntegrationIDData *ExternalIntegrationIDDataType         `xml:"urn:com.workday/bsvc External_Integration_ID_Data,omitempty"`
	DocumentFieldResultData   []DocumentFieldResultDataType          `xml:"urn:com.workday/bsvc Document_Field_Result_Data,omitempty"`
}

// Contains the applicant's resume(s).
type PreHireResumeDataType struct {
	Resume []PreHireResumeDetailDataType `xml:"urn:com.workday/bsvc Resume,omitempty"`
}

// The resume for an applicant.
type PreHireResumeDetailDataType struct {
	ResumeReference *ApplicantResumeObjectType `xml:"urn:com.workday/bsvc Resume_Reference,omitempty"`
	ResumeData      *AttachmentDataWWSType     `xml:"urn:com.workday/bsvc Resume_Data,omitempty"`
}

// Contains the preferred name for a person.  If no preferred name is returned then the legal name is assumed to be the preferred name.  If a preferred name is not provided in a request then the legal name is assumed to be the preferred name.
type PreferredNameDataType struct {
	NameDetailData *PersonNameDetailDataType `xml:"urn:com.workday/bsvc Name_Detail_Data"`
}

type ProcessingFaultType struct {
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProcessmaintainedRoleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProcessmaintainedRoleObjectType struct {
	ID         []ProcessmaintainedRoleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for the Organization Professional Affiliation information.
type ProfessionalAffiliationOrganizationDataType struct {
	ProfessionalAffiliationID                        string                                             `xml:"urn:com.workday/bsvc Professional_Affiliation_ID,omitempty"`
	RemoveProfessionalAffiliation                    *bool                                              `xml:"urn:com.workday/bsvc Remove_Professional_Affiliation,omitempty"`
	ProfessionalAffiliationReference                 *ProfessionalAffiliationReferenceObjectType        `xml:"urn:com.workday/bsvc Professional_Affiliation_Reference,omitempty"`
	ProfessionalAffiliation                          string                                             `xml:"urn:com.workday/bsvc Professional_Affiliation,omitempty"`
	ProfessionalAffiliationTypeReference             *ProfessionalAffiliationTypeAllObjectType          `xml:"urn:com.workday/bsvc Professional_Affiliation_Type_Reference,omitempty"`
	Affiliation                                      string                                             `xml:"urn:com.workday/bsvc Affiliation,omitempty"`
	ProfessionalAffiliationRelationshipTypeReference *ProfessionalAffiliationRelationshipTypeObjectType `xml:"urn:com.workday/bsvc Professional_Affiliation_Relationship_Type_Reference,omitempty"`
	BeginDate                                        *time.Time                                         `xml:"urn:com.workday/bsvc Begin_Date,omitempty"`
	EndDate                                          *time.Time                                         `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	ContactInformationData                           []ContactInformationDataType                       `xml:"urn:com.workday/bsvc Contact_Information_Data,omitempty"`
}

func (t *ProfessionalAffiliationOrganizationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProfessionalAffiliationOrganizationDataType
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
func (t *ProfessionalAffiliationOrganizationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProfessionalAffiliationOrganizationDataType
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
type ProfessionalAffiliationReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProfessionalAffiliationReferenceObjectType struct {
	ID         []ProfessionalAffiliationReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProfessionalAffiliationRelationshipTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProfessionalAffiliationRelationshipTypeObjectType struct {
	ID         []ProfessionalAffiliationRelationshipTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProfessionalAffiliationSkillObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProfessionalAffiliationSkillObjectType struct {
	ID         []ProfessionalAffiliationSkillObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for the Organization Professional Affiliation information.
type ProfessionalAffiliationSkillType struct {
	OrganizationProfessionalAffiliationReference *ProfessionalAffiliationSkillObjectType       `xml:"urn:com.workday/bsvc Organization_Professional_Affiliation_Reference,omitempty"`
	OrganizationProfessionalAffiliationData      []ProfessionalAffiliationOrganizationDataType `xml:"urn:com.workday/bsvc Organization_Professional_Affiliation_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProfessionalAffiliationTypeAllObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProfessionalAffiliationTypeAllObjectType struct {
	ID         []ProfessionalAffiliationTypeAllObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProficiencyRatingBehaviorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProficiencyRatingBehaviorObjectType struct {
	ID         []ProficiencyRatingBehaviorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProficiencyRatingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProficiencyRatingObjectType struct {
	ID         []ProficiencyRatingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type PronounObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PronounObjectType struct {
	ID         []PronounObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for the Propose Compensation for Hire and Offer sub business processes.  If any errors are found during processing, the Auto Complete boolean will be set to False and manual processing will occur for these business processes.
type ProposeCompensationForEmploymentSubBusinessProcessType struct {
	BusinessSubProcessParameters   *BusinessSubProcessParametersType          `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	ProposeCompensationforHireData *CompensationProposedForEmploymentDataType `xml:"urn:com.workday/bsvc Propose_Compensation_for_Hire_Data,omitempty"`
}

// Encapsulating element containing all Allowance Plan Compensation data.
type ProposedAllowancePlanAssignmentContainerDataType struct {
	AllowancePlanSubData []ProposedAllowancePlanAssignmentDataType `xml:"urn:com.workday/bsvc Allowance_Plan_Sub_Data,omitempty"`
	Replace              bool                                      `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}

// Encapsulating element containing all Allowance Plan Compensation data.
type ProposedAllowancePlanAssignmentDataType struct {
	AllowancePlanReference                  *AllowanceValuePlanObjectType `xml:"urn:com.workday/bsvc Allowance_Plan_Reference,omitempty"`
	Percent                                 float64                       `xml:"urn:com.workday/bsvc Percent,omitempty"`
	Amount                                  float64                       `xml:"urn:com.workday/bsvc Amount,omitempty"`
	ManagebyCompensationBasisOverrideAmount float64                       `xml:"urn:com.workday/bsvc Manage_by_Compensation_Basis_Override_Amount,omitempty"`
	CurrencyReference                       *CurrencyObjectType           `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference                      *FrequencyObjectType          `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	ExpectedEndDate                         *time.Time                    `xml:"urn:com.workday/bsvc Expected_End_Date,omitempty"`
	ReimbursementStartDate                  *time.Time                    `xml:"urn:com.workday/bsvc Reimbursement_Start_Date,omitempty"`
	ActualEndDate                           *time.Time                    `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	FixedforManagebyBasisTotal              *bool                         `xml:"urn:com.workday/bsvc Fixed_for_Manage_by_Basis_Total,omitempty"`
}

func (t *ProposedAllowancePlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProposedAllowancePlanAssignmentDataType
	var layout struct {
		*T
		ExpectedEndDate        *xsdDate `xml:"urn:com.workday/bsvc Expected_End_Date,omitempty"`
		ReimbursementStartDate *xsdDate `xml:"urn:com.workday/bsvc Reimbursement_Start_Date,omitempty"`
		ActualEndDate          *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ExpectedEndDate = (*xsdDate)(layout.T.ExpectedEndDate)
	layout.ReimbursementStartDate = (*xsdDate)(layout.T.ReimbursementStartDate)
	layout.ActualEndDate = (*xsdDate)(layout.T.ActualEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ProposedAllowancePlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProposedAllowancePlanAssignmentDataType
	var overlay struct {
		*T
		ExpectedEndDate        *xsdDate `xml:"urn:com.workday/bsvc Expected_End_Date,omitempty"`
		ReimbursementStartDate *xsdDate `xml:"urn:com.workday/bsvc Reimbursement_Start_Date,omitempty"`
		ActualEndDate          *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ExpectedEndDate = (*xsdDate)(overlay.T.ExpectedEndDate)
	overlay.ReimbursementStartDate = (*xsdDate)(overlay.T.ReimbursementStartDate)
	overlay.ActualEndDate = (*xsdDate)(overlay.T.ActualEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Unit Allowance Plan Compensation data.
type ProposedAllowanceUnitPlanAssignmentContainerDataType struct {
	UnitAllowancePlanSubData []ProposedAllowanceUnitPlanAssignmentDataType `xml:"urn:com.workday/bsvc Unit_Allowance_Plan_Sub_Data,omitempty"`
	Replace                  bool                                          `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}

// Encapsulating element containing all Unit Allowance Plan Compensation data.
type ProposedAllowanceUnitPlanAssignmentDataType struct {
	UnitAllowancePlanReference *AllowanceUnitPlanObjectType `xml:"urn:com.workday/bsvc Unit_Allowance_Plan_Reference,omitempty"`
	NumberofUnits              float64                      `xml:"urn:com.workday/bsvc Number_of_Units,omitempty"`
	FrequencyReference         *FrequencyObjectType         `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	PerUnitAmount              float64                      `xml:"urn:com.workday/bsvc Per_Unit_Amount,omitempty"`
	CurrencyReference          *CurrencyObjectType          `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	ReimbursementStartDate     *time.Time                   `xml:"urn:com.workday/bsvc Reimbursement_Start_Date,omitempty"`
	ActualEndDate              *time.Time                   `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	FixedforManagebyBasisTotal *bool                        `xml:"urn:com.workday/bsvc Fixed_for_Manage_by_Basis_Total,omitempty"`
}

func (t *ProposedAllowanceUnitPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProposedAllowanceUnitPlanAssignmentDataType
	var layout struct {
		*T
		ReimbursementStartDate *xsdDate `xml:"urn:com.workday/bsvc Reimbursement_Start_Date,omitempty"`
		ActualEndDate          *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ReimbursementStartDate = (*xsdDate)(layout.T.ReimbursementStartDate)
	layout.ActualEndDate = (*xsdDate)(layout.T.ActualEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ProposedAllowanceUnitPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProposedAllowanceUnitPlanAssignmentDataType
	var overlay struct {
		*T
		ReimbursementStartDate *xsdDate `xml:"urn:com.workday/bsvc Reimbursement_Start_Date,omitempty"`
		ActualEndDate          *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ReimbursementStartDate = (*xsdDate)(overlay.T.ReimbursementStartDate)
	overlay.ActualEndDate = (*xsdDate)(overlay.T.ActualEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Salary or Hourly Plan Compensation data.
type ProposedBasePayPlanAssignmentContainerDataType struct {
	PayPlanSubData []ProposedBasePayPlanAssignmentDataType `xml:"urn:com.workday/bsvc Pay_Plan_Sub_Data,omitempty"`
	Replace        bool                                    `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}

// Encapsulating element containing all Salary or Hourly Plan Compensation data.
type ProposedBasePayPlanAssignmentDataType struct {
	PayPlanReference   *SalaryPayPlanObjectType `xml:"urn:com.workday/bsvc Pay_Plan_Reference,omitempty"`
	Amount             float64                  `xml:"urn:com.workday/bsvc Amount"`
	PercentChange      float64                  `xml:"urn:com.workday/bsvc Percent_Change"`
	AmountChange       float64                  `xml:"urn:com.workday/bsvc Amount_Change"`
	CurrencyReference  *CurrencyObjectType      `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference *FrequencyObjectType     `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	ExpectedEndDate    *time.Time               `xml:"urn:com.workday/bsvc Expected_End_Date,omitempty"`
	ActualEndDate      *time.Time               `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
}

func (t *ProposedBasePayPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProposedBasePayPlanAssignmentDataType
	var layout struct {
		*T
		ExpectedEndDate *xsdDate `xml:"urn:com.workday/bsvc Expected_End_Date,omitempty"`
		ActualEndDate   *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ExpectedEndDate = (*xsdDate)(layout.T.ExpectedEndDate)
	layout.ActualEndDate = (*xsdDate)(layout.T.ActualEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ProposedBasePayPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProposedBasePayPlanAssignmentDataType
	var overlay struct {
		*T
		ExpectedEndDate *xsdDate `xml:"urn:com.workday/bsvc Expected_End_Date,omitempty"`
		ActualEndDate   *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ExpectedEndDate = (*xsdDate)(overlay.T.ExpectedEndDate)
	overlay.ActualEndDate = (*xsdDate)(overlay.T.ActualEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Salary or Hourly Plan Compensation data.
type ProposedBasePayPlanJobAssignmentContainerDataType struct {
	PayPlanSubData []ProposedBasePayPlanJobAssignmentDataType `xml:"urn:com.workday/bsvc Pay_Plan_Sub_Data,omitempty"`
	Replace        bool                                       `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}

// Encapsulating element containing all Salary or Hourly Plan Compensation data.
type ProposedBasePayPlanJobAssignmentDataType struct {
	PayPlanReference   *SalaryPayPlanObjectType `xml:"urn:com.workday/bsvc Pay_Plan_Reference,omitempty"`
	Amount             float64                  `xml:"urn:com.workday/bsvc Amount"`
	CurrencyReference  *CurrencyObjectType      `xml:"urn:com.workday/bsvc Currency_Reference"`
	FrequencyReference *FrequencyObjectType     `xml:"urn:com.workday/bsvc Frequency_Reference"`
	ExpectedEndDate    *time.Time               `xml:"urn:com.workday/bsvc Expected_End_Date,omitempty"`
	ActualEndDate      *time.Time               `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
}

func (t *ProposedBasePayPlanJobAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProposedBasePayPlanJobAssignmentDataType
	var layout struct {
		*T
		ExpectedEndDate *xsdDate `xml:"urn:com.workday/bsvc Expected_End_Date,omitempty"`
		ActualEndDate   *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ExpectedEndDate = (*xsdDate)(layout.T.ExpectedEndDate)
	layout.ActualEndDate = (*xsdDate)(layout.T.ActualEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ProposedBasePayPlanJobAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProposedBasePayPlanJobAssignmentDataType
	var overlay struct {
		*T
		ExpectedEndDate *xsdDate `xml:"urn:com.workday/bsvc Expected_End_Date,omitempty"`
		ActualEndDate   *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ExpectedEndDate = (*xsdDate)(overlay.T.ExpectedEndDate)
	overlay.ActualEndDate = (*xsdDate)(overlay.T.ActualEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Bonus Plan Compensation data.
type ProposedBonusPlanAssignmentContainerDataType struct {
	BonusPlanSubData []ProposedBonusPlanAssignmentDataType `xml:"urn:com.workday/bsvc Bonus_Plan_Sub_Data,omitempty"`
	Replace          bool                                  `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}

// Encapsulating element containing all Bonus Plan Compensation data.
type ProposedBonusPlanAssignmentDataType struct {
	BonusPlanReference                      *BonusPlanObjectType `xml:"urn:com.workday/bsvc Bonus_Plan_Reference,omitempty"`
	IndividualTargetAmount                  float64              `xml:"urn:com.workday/bsvc Individual_Target_Amount,omitempty"`
	IndividualTargetPercent                 float64              `xml:"urn:com.workday/bsvc Individual_Target_Percent,omitempty"`
	ManagebyCompensationBasisOverrideAmount float64              `xml:"urn:com.workday/bsvc Manage_by_Compensation_Basis_Override_Amount,omitempty"`
	GuaranteedMinimum                       *bool                `xml:"urn:com.workday/bsvc Guaranteed_Minimum,omitempty"`
	PercentAssigned                         float64              `xml:"urn:com.workday/bsvc Percent_Assigned,omitempty"`
	ActualEndDate                           *time.Time           `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	FixedforManagebyBasisTotal              *bool                `xml:"urn:com.workday/bsvc Fixed_for_Manage_by_Basis_Total,omitempty"`
}

func (t *ProposedBonusPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProposedBonusPlanAssignmentDataType
	var layout struct {
		*T
		ActualEndDate *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ActualEndDate = (*xsdDate)(layout.T.ActualEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ProposedBonusPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProposedBonusPlanAssignmentDataType
	var overlay struct {
		*T
		ActualEndDate *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ActualEndDate = (*xsdDate)(overlay.T.ActualEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Calculated Plan Compensation data.
type ProposedCalculatedPlanAssignmentContainerDataType struct {
	CalculatedPlanSubData []ProposedCalculatedPlanAssignmentDataType `xml:"urn:com.workday/bsvc Calculated_Plan_Sub_Data,omitempty"`
	Replace               bool                                       `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}

// Encapsulating element containing all Calculated Plan Compensation data.
type ProposedCalculatedPlanAssignmentDataType struct {
	CalculatedPlanReference                 *CalculatedPlanObjectType `xml:"urn:com.workday/bsvc Calculated_Plan_Reference,omitempty"`
	ManagebyCompensationBasisOverrideAmount float64                   `xml:"urn:com.workday/bsvc Manage_by_Compensation_Basis_Override_Amount,omitempty"`
	CurrencyReference                       *CurrencyObjectType       `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference                      *FrequencyObjectType      `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	ActualEndDate                           *time.Time                `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
}

func (t *ProposedCalculatedPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProposedCalculatedPlanAssignmentDataType
	var layout struct {
		*T
		ActualEndDate *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ActualEndDate = (*xsdDate)(layout.T.ActualEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ProposedCalculatedPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProposedCalculatedPlanAssignmentDataType
	var overlay struct {
		*T
		ActualEndDate *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ActualEndDate = (*xsdDate)(overlay.T.ActualEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Commission Plan Compensation data.
type ProposedCommissionPlanAssignmentContainerDataType struct {
	CommissionPlanSubData []ProposedCommissionPlanAssignmentDataType `xml:"urn:com.workday/bsvc Commission_Plan_Sub_Data,omitempty"`
	Replace               bool                                       `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}

// Encapsulating element containing all Commission Plan Compensation data.
type ProposedCommissionPlanAssignmentDataType struct {
	CommissionPlanReference    *CommissionPlanObjectType `xml:"urn:com.workday/bsvc Commission_Plan_Reference,omitempty"`
	TargetAmount               float64                   `xml:"urn:com.workday/bsvc Target_Amount,omitempty"`
	CurrencyReference          *CurrencyObjectType       `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference         *FrequencyObjectType      `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	DrawAmount                 float64                   `xml:"urn:com.workday/bsvc Draw_Amount,omitempty"`
	DrawFrequencyReference     *FrequencyObjectType      `xml:"urn:com.workday/bsvc Draw_Frequency_Reference,omitempty"`
	DrawDuration               string                    `xml:"urn:com.workday/bsvc Draw_Duration,omitempty"`
	Recoverable                *bool                     `xml:"urn:com.workday/bsvc Recoverable,omitempty"`
	ActualEndDate              *time.Time                `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	FixedforManagebyBasisTotal *bool                     `xml:"urn:com.workday/bsvc Fixed_for_Manage_by_Basis_Total,omitempty"`
}

func (t *ProposedCommissionPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProposedCommissionPlanAssignmentDataType
	var layout struct {
		*T
		ActualEndDate *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ActualEndDate = (*xsdDate)(layout.T.ActualEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ProposedCommissionPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProposedCommissionPlanAssignmentDataType
	var overlay struct {
		*T
		ActualEndDate *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ActualEndDate = (*xsdDate)(overlay.T.ActualEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Merit Plan Compensation data.
type ProposedMeritPlanAssignmentContainerDataType struct {
	MeritPlanSubData []ProposedMeritPlanAssignmentDataType `xml:"urn:com.workday/bsvc Merit_Plan_Sub_Data,omitempty"`
	Replace          bool                                  `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}

// Encapsulating element containing all Merit Plan Compensation data.
type ProposedMeritPlanAssignmentDataType struct {
	MeritPlanReference      *MeritPlanObjectType `xml:"urn:com.workday/bsvc Merit_Plan_Reference,omitempty"`
	IndividualTargetAmount  float64              `xml:"urn:com.workday/bsvc Individual_Target_Amount,omitempty"`
	IndividualTargetPercent float64              `xml:"urn:com.workday/bsvc Individual_Target_Percent,omitempty"`
	GuaranteedMinimum       *bool                `xml:"urn:com.workday/bsvc Guaranteed_Minimum,omitempty"`
	ActualEndDate           *time.Time           `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
}

func (t *ProposedMeritPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProposedMeritPlanAssignmentDataType
	var layout struct {
		*T
		ActualEndDate *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ActualEndDate = (*xsdDate)(layout.T.ActualEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ProposedMeritPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProposedMeritPlanAssignmentDataType
	var overlay struct {
		*T
		ActualEndDate *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ActualEndDate = (*xsdDate)(overlay.T.ActualEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Period Salary Plan Compensation data.
type ProposedPeriodSalaryPlanAssignmentContainerDataType struct {
	PeriodSalaryPlanSubData []ProposedPeriodSalaryPlanAssignmentDataType `xml:"urn:com.workday/bsvc Period_Salary_Plan_Sub_Data,omitempty"`
	Replace                 bool                                         `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}

// Encapsulating element containing all Period Salary Plan Compensation data.
type ProposedPeriodSalaryPlanAssignmentDataType struct {
	PeriodSalaryPlanReference               *PeriodSalaryPlanObjectType   `xml:"urn:com.workday/bsvc Period_Salary_Plan_Reference,omitempty"`
	CompensationPeriodReference             *CompensationPeriodObjectType `xml:"urn:com.workday/bsvc Compensation_Period_Reference,omitempty"`
	ManagebyCompensationBasisOverrideAmount float64                       `xml:"urn:com.workday/bsvc Manage_by_Compensation_Basis_Override_Amount,omitempty"`
	CurrencyReference                       *CurrencyObjectType           `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	CompensationPeriodMultiplier            float64                       `xml:"urn:com.workday/bsvc Compensation_Period_Multiplier,omitempty"`
	FrequencyReference                      *FrequencyObjectType          `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	ActualEndDate                           *time.Time                    `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
}

func (t *ProposedPeriodSalaryPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProposedPeriodSalaryPlanAssignmentDataType
	var layout struct {
		*T
		ActualEndDate *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ActualEndDate = (*xsdDate)(layout.T.ActualEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ProposedPeriodSalaryPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProposedPeriodSalaryPlanAssignmentDataType
	var overlay struct {
		*T
		ActualEndDate *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ActualEndDate = (*xsdDate)(overlay.T.ActualEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Unit Salary Plan Compensation data.
type ProposedSalaryUnitPlanAssignmentContainerDataType struct {
	UnitSalaryPlanSubData []ProposedSalaryUnitPlanAssignmentDataType `xml:"urn:com.workday/bsvc Unit_Salary_Plan_Sub_Data,omitempty"`
	Replace               bool                                       `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}

// Encapsulating element containing all Unit Salary Plan Compensation data.
type ProposedSalaryUnitPlanAssignmentDataType struct {
	UnitSalaryPlanReference *SalaryUnitPlanObjectType `xml:"urn:com.workday/bsvc Unit_Salary_Plan_Reference,omitempty"`
	PerUnitAmount           float64                   `xml:"urn:com.workday/bsvc Per_Unit_Amount,omitempty"`
	CurrencyReference       *CurrencyObjectType       `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	DefaultUnits            float64                   `xml:"urn:com.workday/bsvc Default_Units,omitempty"`
	FrequencyReference      *FrequencyObjectType      `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	ActualEndDate           *time.Time                `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
}

func (t *ProposedSalaryUnitPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProposedSalaryUnitPlanAssignmentDataType
	var layout struct {
		*T
		ActualEndDate *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ActualEndDate = (*xsdDate)(layout.T.ActualEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ProposedSalaryUnitPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProposedSalaryUnitPlanAssignmentDataType
	var overlay struct {
		*T
		ActualEndDate *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ActualEndDate = (*xsdDate)(overlay.T.ActualEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Stock Plan Compensation data.
type ProposedStockPlanAssignmentContainerDataType struct {
	StockPlanSubData []ProposedStockPlanAssignmentDataType `xml:"urn:com.workday/bsvc Stock_Plan_Sub_Data,omitempty"`
	Replace          bool                                  `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}

// Encapsulating element containing all Stock Plan Compensation data.
type ProposedStockPlanAssignmentDataType struct {
	StockPlanReference                      *StockPlanObjectType `xml:"urn:com.workday/bsvc Stock_Plan_Reference,omitempty"`
	IndividualTargetShares                  float64              `xml:"urn:com.workday/bsvc Individual_Target_Shares,omitempty"`
	IndividualTargetPercent                 float64              `xml:"urn:com.workday/bsvc Individual_Target_Percent,omitempty"`
	IndividualTargetAmount                  float64              `xml:"urn:com.workday/bsvc Individual_Target_Amount,omitempty"`
	ManagebyCompensationBasisOverrideAmount float64              `xml:"urn:com.workday/bsvc Manage_by_Compensation_Basis_Override_Amount,omitempty"`
	CurrencyReference                       *CurrencyObjectType  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	ActualEndDate                           *time.Time           `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	FixedforManagebyBasisTotal              *bool                `xml:"urn:com.workday/bsvc Fixed_for_Manage_by_Basis_Total,omitempty"`
}

func (t *ProposedStockPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProposedStockPlanAssignmentDataType
	var layout struct {
		*T
		ActualEndDate *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ActualEndDate = (*xsdDate)(layout.T.ActualEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ProposedStockPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProposedStockPlanAssignmentDataType
	var overlay struct {
		*T
		ActualEndDate *xsdDate `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ActualEndDate = (*xsdDate)(overlay.T.ActualEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Resume/Cover Letter Attachments for prospect
type ProspectAttachmentsDataType struct {
	ResumeAttachments []AttachmentWWSDataType `xml:"urn:com.workday/bsvc Resume_Attachments,omitempty"`
}

// Wrapper element for the Put Applicant Request Information.
type PutApplicantRequestType struct {
	ApplicantReference *ApplicantObjectType `xml:"urn:com.workday/bsvc Applicant_Reference,omitempty"`
	ApplicantData      *PreHireDataWWSType  `xml:"urn:com.workday/bsvc Applicant_Data"`
	Version            string               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Put Applicant Returned data.
type PutApplicantResponseType struct {
	ApplicantReference     *ApplicantObjectType                           `xml:"urn:com.workday/bsvc Applicant_Reference,omitempty"`
	ExceptionsResponseData []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing request attributes, Background Check Package reference, and test details.
type PutBackgroundCheckPackageRequestType struct {
	BackgroundCheckPackageReference *BackgroundCheckPackageObjectType `xml:"urn:com.workday/bsvc Background_Check_Package_Reference,omitempty"`
	BackgroundCheckPackageData      *BackgroundCheckPackageDataType   `xml:"urn:com.workday/bsvc Background_Check_Package_Data"`
	AddOnly                         bool                              `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                         string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the response of the request to create or update a Background Check Package.
type PutBackgroundCheckPackageResponseType struct {
	BackgroundCheckPackageReference *BackgroundCheckPackageObjectType `xml:"urn:com.workday/bsvc Background_Check_Package_Reference,omitempty"`
	Version                         string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request to update the status of a Background Check on a person. This operation assumes that the Background Check is using a Background Check Package.
type PutBackgroundCheckRequestType struct {
	BusinessProcessParameters        *BusinessProcessParametersType        `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	DynamicBusinessProcessParameters *DynamicBusinessProcessParametersType `xml:"urn:com.workday/bsvc Dynamic_Business_Process_Parameters,omitempty"`
	BackgroundCheckData              *BackgroundCheckDataType              `xml:"urn:com.workday/bsvc Background_Check_Data"`
	Version                          string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The Event modified in this operation.
type PutBackgroundCheckResponseType struct {
	EventReference *BackgroundCheckEventObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a Candidate and/or Job Application to add Attachments to.
type PutCandidateAttachmentRequestType struct {
	CandidateReference            *CandidateObjectType               `xml:"urn:com.workday/bsvc Candidate_Reference,omitempty"`
	ProspectResumeAttachmentsData *ProspectAttachmentsDataType       `xml:"urn:com.workday/bsvc Prospect_Resume_Attachments_Data,omitempty"`
	CandidateAttachmentData       *CandidateAttachmentDataType       `xml:"urn:com.workday/bsvc Candidate_Attachment_Data,omitempty"`
	JobApplicationAttachmentData  []JobApplicationAttachmentDataType `xml:"urn:com.workday/bsvc Job_Application_Attachment_Data,omitempty"`
	AddOnly                       bool                               `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains information about about the Candidate and/or Job Application the attachment was added to.
type PutCandidateAttachmentResponseType struct {
	CandidateReference           *CandidateObjectType                   `xml:"urn:com.workday/bsvc Candidate_Reference,omitempty"`
	AttachmentReference          []CandidateAttachmentObjectType        `xml:"urn:com.workday/bsvc Attachment_Reference,omitempty"`
	JobApplicationAttachmentData []JobApplicationAttachmentResponseType `xml:"urn:com.workday/bsvc Job_Application_Attachment_Data,omitempty"`
	ProspectResumeAttachmentData *ProspectAttachmentsDataType           `xml:"urn:com.workday/bsvc Prospect_Resume_Attachment_Data,omitempty"`
	Version                      string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains information about the candidate photo.
type PutCandidatePhotoRequestType struct {
	CandidateReference *CandidateObjectType    `xml:"urn:com.workday/bsvc Candidate_Reference,omitempty"`
	CandidatePhotoData *CandidatePhotoDataType `xml:"urn:com.workday/bsvc Candidate_Photo_Data,omitempty"`
	Version            string                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains information about the candidate and their photo.
type PutCandidatePhotosResponseType struct {
	CandidateReference *CandidateObjectType `xml:"urn:com.workday/bsvc Candidate_Reference,omitempty"`
	Version            string               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update a candidate.
type PutCandidateRequestType struct {
	CandidateReference *CandidateObjectType `xml:"urn:com.workday/bsvc Candidate_Reference,omitempty"`
	CandidateData      *CandidateDataType   `xml:"urn:com.workday/bsvc Candidate_Data"`
	AddOnly            bool                 `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version            string               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for Put Candidate response.
type PutCandidateResponseType struct {
	CandidateReference          *CandidateObjectType              `xml:"urn:com.workday/bsvc Candidate_Reference,omitempty"`
	CandidateJobApplicationData []CandidateJobApplicationDataType `xml:"urn:com.workday/bsvc Candidate_Job_Application_Data,omitempty"`
	Version                     string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Import Process Response element
type PutImportProcessResponseType struct {
	ImportProcessReference  *WebServiceBackgroundProcessRuntimeObjectType `xml:"urn:com.workday/bsvc Import_Process_Reference,omitempty"`
	HeaderInstanceReference *InstanceObjectType                           `xml:"urn:com.workday/bsvc Header_Instance_Reference,omitempty"`
	Version                 string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// A request to submit interview feedback for one or more interviewers.
type PutInterviewFeedbackRequestType struct {
	DynamicBusinessProcessParameters *DynamicBusinessProcessParametersType `xml:"urn:com.workday/bsvc Dynamic_Business_Process_Parameters,omitempty"`
	InterviewFeedbackData            *InterviewFeedbackDataType            `xml:"urn:com.workday/bsvc Interview_Feedback_Data"`
	Version                          string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for Put Interview Feedback response.
type PutInterviewFeedbackResponseType struct {
	InterviewEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Interview_Event_Reference,omitempty"`
	Version                 string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// A request to schedule an interview.
type PutInterviewRequestType struct {
	InterviewData *InterviewDataType `xml:"urn:com.workday/bsvc Interview_Data"`
	Version       string             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for Put Interview response.
type PutInterviewResponseType struct {
	InterviewReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Interview_Reference,omitempty"`
	Version            string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains Job Application Additional data
type PutJobApplicationAdditionalDataRequestType struct {
	JobApplicationCustomObjectData *JobApplicationCustomObjectDataType `xml:"urn:com.workday/bsvc Job_Application_Custom_Object_Data"`
	Version                        string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains Job Application Additional data
type PutJobApplicationAdditionalDataResponseType struct {
	JobApplicationAdditionalDataReference *JobApplicationObjectType                      `xml:"urn:com.workday/bsvc Job_Application_Additional_Data_Reference,omitempty"`
	JobApplicationAdditionalData          *NonEffectiveDatedWebServiceAdditionalDataType `xml:"urn:com.workday/bsvc Job_Application_Additional_Data,omitempty"`
	Version                               string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Container which holds the request data for a Job Posting Site Put service.
type PutJobPostingSiteRequestType struct {
	JobPostingSiteReference *GenericJobPostingSiteObjectType `xml:"urn:com.workday/bsvc Job_Posting_Site_Reference,omitempty"`
	JobPostingSiteData      *JobPostingSiteDetailDataType    `xml:"urn:com.workday/bsvc Job_Posting_Site_Data"`
	AddOnly                 bool                             `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                 string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// After a PUT operation this element will be the response. It holds the updated Job Posting Site.
type PutJobPostingSiteResponseType struct {
	JobPostingSiteReference *JobPostingSiteObjectType `xml:"urn:com.workday/bsvc Job_Posting_Site_Reference,omitempty"`
	Version                 string                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the job requisition reference and job requisition interview team data.
type PutJobRequisitionInterviewTeamRequestType struct {
	JobRequisitionInterviewTeamReference *JobRequisitionEnabledObjectType     `xml:"urn:com.workday/bsvc Job_Requisition_Interview_Team_Reference,omitempty"`
	JobRequisitionInterviewTeamData      *JobRequisitionInterviewTeamDataType `xml:"urn:com.workday/bsvc Job_Requisition_Interview_Team_Data,omitempty"`
	Version                              string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains job requisition interview team data.
type PutJobRequisitionInterviewTeamResponseType struct {
	JobRequisitionReference *JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	Version                 string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request information to make a Job Posting the Primary Posting
type PutPrimaryPostingRequestType struct {
	JobRequisitionReference *JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
	JobPostingReference     *JobPostingObjectType            `xml:"urn:com.workday/bsvc Job_Posting_Reference"`
	ForecastedPayout        float64                          `xml:"urn:com.workday/bsvc Forecasted_Payout,omitempty"`
	AgencyCurrencyReference *CurrencyObjectType              `xml:"urn:com.workday/bsvc Agency_Currency_Reference,omitempty"`
	Version                 string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for Primary Posting Response.
type PutPrimaryPostingResponseType struct {
	PrimaryPostingData []JobPostingDataType `xml:"urn:com.workday/bsvc Primary_Posting_Data,omitempty"`
	Version            string               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Custom Object Information for Recruiting Agency.
type PutRecruitingAgencyAdditionalDataRequestType struct {
	RecruitingAgencyCustomObjectData *RecruitingAgencyCustomObjectDataType `xml:"urn:com.workday/bsvc Recruiting_Agency_Custom_Object_Data"`
	Version                          string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Custom Object Information for Recruiting Agency.
type PutRecruitingAgencyAdditionalDataResponseType struct {
	RecruitingAgencyAdditionalDataReference *RecruitingAgencySiteObjectType                `xml:"urn:com.workday/bsvc Recruiting_Agency_Additional_Data_Reference,omitempty"`
	RecruitingAgencyAdditionalData          *NonEffectiveDatedWebServiceAdditionalDataType `xml:"urn:com.workday/bsvc Recruiting_Agency_Additional_Data,omitempty"`
	Version                                 string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update a Recruiting Agency
type PutRecruitingAgencyRequestType struct {
	RecruitingAgencyReference *RecruitingAgencySiteObjectType `xml:"urn:com.workday/bsvc Recruiting_Agency_Reference,omitempty"`
	RecruitingAgencyData      *RecruitingAgencyDataType       `xml:"urn:com.workday/bsvc Recruiting_Agency_Data"`
	AddOnly                   bool                            `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                   string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Recruiting Agency Response
type PutRecruitingAgencyResponseType struct {
	RecruitingAgencyReference *RecruitingAgencySiteObjectType `xml:"urn:com.workday/bsvc Recruiting_Agency_Reference,omitempty"`
	Version                   string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update a Recruiting Agency User
type PutRecruitingAgencyUserRequestType struct {
	RecruitingAgencyUserReference *RecruitingAgencyUserObjectType `xml:"urn:com.workday/bsvc Recruiting_Agency_User_Reference,omitempty"`
	RecruitingAgencyUserData      *RecruitingAgencyUserDataType   `xml:"urn:com.workday/bsvc Recruiting_Agency_User_Data"`
	AddOnly                       bool                            `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                       string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Returns the Recruiting Agency User
type PutRecruitingAgencyUserResponseType struct {
	RecruitingAgencyUserReference *RecruitingAgencyUserObjectType `xml:"urn:com.workday/bsvc Recruiting_Agency_User_Reference,omitempty"`
	Version                       string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Create or Update Recruiting Assessment Categories
type PutRecruitingAssessmentCategoryRequestType struct {
	RecruitingAssessmentCategoryReference *RecruitingAssessmentCategoryObjectType `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Reference,omitempty"`
	RecruitingAssessmentCategoryData      *RecruitingAssessmentCategoryDataType   `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Data,omitempty"`
	AddOnly                               bool                                    `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Delete                                bool                                    `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	Version                               string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response from Put Recruiting Assessment Category
type PutRecruitingAssessmentCategoryResponseType struct {
	RecruitingAssessmentCategoryReference *RecruitingAssessmentCategoryObjectType `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Reference,omitempty"`
	Version                               string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the request to add, update or delete a Recruiting Assessment Category Security Segment.
type PutRecruitingAssessmentCategorySecuritySegmentRequestType struct {
	RecruitingAssessmentCategorySecuritySegmentReference *RecruitingAssessmentCategorySecuritySegmentObjectType `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Security_Segment_Reference,omitempty"`
	RecruitingAssessmentCategorySecuritySegmentData      *RecruitingAssessmentCategorySecuritySegmentDataType   `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Security_Segment_Data,omitempty"`
	AddOnly                                              bool                                                   `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Delete                                               bool                                                   `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	Version                                              string                                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the Recruiting Assessment Category Security Segment created or modified by the web service.
type PutRecruitingAssessmentCategorySecuritySegmentResponseType struct {
	RecruitingAssessmentCategorySecuritySegmentReference *RecruitingAssessmentCategorySecuritySegmentObjectType `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Security_Segment_Reference,omitempty"`
	Version                                              string                                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the request to add, update or delete a Recruiting Self-Schedule Calendar.
type PutRecruitingSelfScheduleCalendarRequestType struct {
	RecruitingSelfScheduleCalendarReference *RecruitingSelfScheduleCalendarObjectType `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_Reference,omitempty"`
	RecruitingSelfScheduleCalendarData      *RecruitingSelfScheduleCalendarDataType   `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_Data,omitempty"`
	AddOnly                                 bool                                      `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Delete                                  bool                                      `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	Version                                 string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The Recruiting Self-Schedule Calendar created or modified by the web service.
type PutRecruitingSelfScheduleCalendarResponseType struct {
	RecruitingSelfScheduleCalendarReference *RecruitingSelfScheduleCalendarObjectType `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_Reference,omitempty"`
	Version                                 string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the request to add, update or delete a Recruiting Self-Schedule Calendar Type.
type PutRecruitingSelfScheduleCalendarTypeRequestType struct {
	RecruitingSelfScheduleCalendarTypeReference *RecruitingSelfScheduleCalendarTypeObjectType `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_Type_Reference,omitempty"`
	RecruitingSelfScheduleCalendarTypeData      *RecruitingSelfScheduleCalendarTypeDataType   `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_Type_Data,omitempty"`
	AddOnly                                     bool                                          `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Delete                                      bool                                          `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	Version                                     string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the Recruiting Self-Schedule Calendar Type created or modified by the web service.
type PutRecruitingSelfScheduleCalendarTypeResponseType struct {
	RecruitingSelfScheduleCalendarTypeReference *RecruitingSelfScheduleCalendarTypeObjectType `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_Type_Reference,omitempty"`
	Version                                     string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// A request to undo the move from a staffing process back to the originating job application process.
type PutUndoMovefromHireRequestType struct {
	JobApplicationEventReference            *UniqueIdentifierObjectType               `xml:"urn:com.workday/bsvc Job_Application_Event_Reference"`
	StaffingEventReference                  *UniqueIdentifierObjectType               `xml:"urn:com.workday/bsvc Staffing_Event_Reference"`
	JobApplicationReference                 *JobApplicationObjectType                 `xml:"urn:com.workday/bsvc Job_Application_Reference"`
	EventClassificationSubcategoryReference *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Event_Classification_Subcategory_Reference"`
	Comment                                 string                                    `xml:"urn:com.workday/bsvc Comment,omitempty"`
	Version                                 string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The job application event that was included in this undo move.
type PutUndoMovefromHireResponseType struct {
	JobApplicationEventReference *UniqueIdentifierObjectType               `xml:"urn:com.workday/bsvc Job_Application_Event_Reference,omitempty"`
	StageReference               *RecruitingStageWorkdayOwnedObjectType    `xml:"urn:com.workday/bsvc Stage_Reference,omitempty"`
	ReasonReference              *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
	Version                      string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains all the request data to add or update a Veteran Status.
type PutVeteranStatusRequestType struct {
	VeteranStatusReference *VeteranStatusObjectType `xml:"urn:com.workday/bsvc Veteran_Status_Reference,omitempty"`
	VeteranStatusData      *VeteranStatusDataType   `xml:"urn:com.workday/bsvc Veteran_Status_Data,omitempty"`
	AddOnly                bool                     `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Delete                 bool                     `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	Version                string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for Put Veteran Status response.
type PutVeteranStatusResponseType struct {
	VeteranStatusReference *VeteranStatusObjectType `xml:"urn:com.workday/bsvc Veteran_Status_Reference,omitempty"`
	Version                string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the qualification data for a Position Restriction or Job Profile.
type QualificationDataforPositionRestrictionorJobProfileType struct {
	ResponsibilityQualificationReplacement *ResponsibilityQualificationReplacementType `xml:"urn:com.workday/bsvc Responsibility_Qualification_Replacement,omitempty"`
	WorkExperienceQualificationReplacement *WorkExperienceQualificationReplacementType `xml:"urn:com.workday/bsvc Work_Experience_Qualification_Replacement,omitempty"`
	EducationQualificationReplacement      *EducationQualificationReplacementType      `xml:"urn:com.workday/bsvc Education_Qualification_Replacement,omitempty"`
	LanguageQualificationReplacement       *LanguageQualificationReplacementType       `xml:"urn:com.workday/bsvc Language_Qualification_Replacement,omitempty"`
	CompetencyQualificationReplacement     *CompetencyQualificationReplacementType     `xml:"urn:com.workday/bsvc Competency_Qualification_Replacement,omitempty"`
	CertificationQualificationReplacement  *CertificationQualificationReplacementType  `xml:"urn:com.workday/bsvc Certification_Qualification_Replacement,omitempty"`
	TrainingQualificationReplacement       *TrainingQualificationReplacementType       `xml:"urn:com.workday/bsvc Training_Qualification_Replacement,omitempty"`
	SkillQualificationReplacement          *SkillQualificationReplacementType          `xml:"urn:com.workday/bsvc Skill_Qualification_Replacement,omitempty"`
}

// Contains the position's and it's job profile's qualification information which includes the following: Competency, Certification, Education, Language, Responsibility, Training, and Work Experience
type QualificationsfromPositionRestrictionsDataType struct {
	CompetencyData     []CompetencyProfileforJobType     `xml:"urn:com.workday/bsvc Competency_Data,omitempty"`
	CertificationData  []CertificationProfileforJobType  `xml:"urn:com.workday/bsvc Certification_Data,omitempty"`
	EducationData      []EducationProfileforJobType      `xml:"urn:com.workday/bsvc Education_Data,omitempty"`
	LanguageData       []LanguageProfileforJobType       `xml:"urn:com.workday/bsvc Language_Data,omitempty"`
	ResponsibilityData []ResponsibilityProfileforJobType `xml:"urn:com.workday/bsvc Responsibility_Data,omitempty"`
	TrainingData       []TrainingProfileforJobType       `xml:"urn:com.workday/bsvc Training_Data,omitempty"`
	WorkExperienceData []WorkExperienceforJobType        `xml:"urn:com.workday/bsvc Work_Experience_Data,omitempty"`
	SkillData          []SkillProfileforJobType          `xml:"urn:com.workday/bsvc Skill_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type QuestionMultipleChoiceAnswerObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type QuestionMultipleChoiceAnswerObjectType struct {
	ID         []QuestionMultipleChoiceAnswerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type QuestionSetupObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type QuestionSetupObjectType struct {
	ID         []QuestionSetupObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the questions and answers on the questionnaire response.
type QuestionnaireAnswerDataType struct {
	QuestionSetupReference   *QuestionSetupObjectType                `xml:"urn:com.workday/bsvc Question_Setup_Reference"`
	MultipleChoiceAnswerData []MultipleChoiceAnswerDataType          `xml:"urn:com.workday/bsvc Multiple_Choice_Answer_Data,omitempty"`
	AnswerDate               *time.Time                              `xml:"urn:com.workday/bsvc Answer_Date,omitempty"`
	AnswerNumeric            float64                                 `xml:"urn:com.workday/bsvc Answer_Numeric,omitempty"`
	AnswerText               string                                  `xml:"urn:com.workday/bsvc Answer_Text,omitempty"`
	AttachmentAnswerData     []QuestionnaireAttachmentAnswerDataType `xml:"urn:com.workday/bsvc Attachment_Answer_Data,omitempty"`
}

func (t *QuestionnaireAnswerDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T QuestionnaireAnswerDataType
	var layout struct {
		*T
		AnswerDate *xsdDate `xml:"urn:com.workday/bsvc Answer_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AnswerDate = (*xsdDate)(layout.T.AnswerDate)
	return e.EncodeElement(layout, start)
}
func (t *QuestionnaireAnswerDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T QuestionnaireAnswerDataType
	var overlay struct {
		*T
		AnswerDate *xsdDate `xml:"urn:com.workday/bsvc Answer_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AnswerDate = (*xsdDate)(overlay.T.AnswerDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains attachment answer data.
type QuestionnaireAttachmentAnswerDataType struct {
	AnswerAttachment *AttachmentWWSDataType `xml:"urn:com.workday/bsvc Answer_Attachment"`
}

// Contains a unique identifier for an instance of an object.
type QuestionnaireObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type QuestionnaireObjectType struct {
	ID         []QuestionnaireObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the questionnaire response.
type QuestionnairesResponseDataType struct {
	QuestionnaireReference  *QuestionnaireObjectType      `xml:"urn:com.workday/bsvc Questionnaire_Reference"`
	QuestionnaireAnswerData []QuestionnaireAnswerDataType `xml:"urn:com.workday/bsvc Questionnaire_Answer_Data"`
}

// Contains Recruiting Agency Additional Data
type RecruitingAgencyAdditionalDataResponseDataType struct {
	RecruitingAgencyCustomObjectData []RecruitingAgencyAdditionalDataType `xml:"urn:com.workday/bsvc Recruiting_Agency_Custom_Object_Data,omitempty"`
}

// Recruiting Agency Additional Data
type RecruitingAgencyAdditionalDataType struct {
	RecruitingAgencyReference *RecruitingAgencySiteObjectType                `xml:"urn:com.workday/bsvc Recruiting_Agency_Reference,omitempty"`
	RecruitingAgencyData      []RecruitingAgencyDataType                     `xml:"urn:com.workday/bsvc Recruiting_Agency_Data,omitempty"`
	AdditionalData            *NonEffectiveDatedWebServiceAdditionalDataType `xml:"urn:com.workday/bsvc Additional_Data,omitempty"`
}

// Recruiting Agency attachment information.
type RecruitingAgencyAttachmentDataType struct {
	AttachmentData []AttachmentWWSDataType `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
}

// Contains information about a Attachment to add to a Candidate.
type RecruitingAgencyCandidateAttachmentDataType struct {
	AttachmentData *AttachmentWWSDataType `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
}

// Contains Candidate Submission data for recruiting agency candidate process.
type RecruitingAgencyCandidateSubmissionDataType struct {
	AgencyReference                 *RecruitingAgencySiteObjectType               `xml:"urn:com.workday/bsvc Agency_Reference"`
	AgencyUserReference             *RecruitingAgencyUserObjectType               `xml:"urn:com.workday/bsvc Agency_User_Reference"`
	JobPostingReference             *JobPostingObjectType                         `xml:"urn:com.workday/bsvc Job_Posting_Reference,omitempty"`
	JobRequisitionReference         *JobRequisitionEnabledObjectType              `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	CandidateNameData               *CandidateNameDataType                        `xml:"urn:com.workday/bsvc Candidate_Name_Data,omitempty"`
	CandidateContactData            *CandidateContactDataType                     `xml:"urn:com.workday/bsvc Candidate_Contact_Data,omitempty"`
	CandidateSocialMediaAccountData []CandidateSocialMediaAccountDataType         `xml:"urn:com.workday/bsvc Candidate_Social_Media_Account_Data,omitempty"`
	CandidateResumeData             *CandidateResumeDataType                      `xml:"urn:com.workday/bsvc Candidate_Resume_Data,omitempty"`
	CandidateAttachmentData         []RecruitingAgencyCandidateAttachmentDataType `xml:"urn:com.workday/bsvc Candidate_Attachment_Data,omitempty"`
}

// Recruiting Agency data.
type RecruitingAgencyCareerSiteConfigurationDataType struct {
	SourceReference              *ApplicantSourceObjectType    `xml:"urn:com.workday/bsvc Source_Reference"`
	OwnershipPeriod              float64                       `xml:"urn:com.workday/bsvc Ownership_Period,omitempty"`
	JobPostingTemplateReference  *JobPostingTemplateObjectType `xml:"urn:com.workday/bsvc Job_Posting_Template_Reference"`
	RestricttoLocationsReference []LocationObjectType          `xml:"urn:com.workday/bsvc Restrict_to_Locations_Reference,omitempty"`
	TermsandConditions           string                        `xml:"urn:com.workday/bsvc Terms_and_Conditions,omitempty"`
	Requireesignature            *bool                         `xml:"urn:com.workday/bsvc Require_e-signature,omitempty"`
}

// Custom Object Information for Recruiting Agency.
type RecruitingAgencyCustomObjectDataType struct {
	RecruitingAgencyReference *RecruitingAgencySiteObjectType                `xml:"urn:com.workday/bsvc Recruiting_Agency_Reference"`
	AdditionalData            *NonEffectiveDatedWebServiceAdditionalDataType `xml:"urn:com.workday/bsvc Additional_Data"`
}

// Recruiting Agency data.
type RecruitingAgencyDataType struct {
	ID                                          string                                           `xml:"urn:com.workday/bsvc ID,omitempty"`
	Name                                        string                                           `xml:"urn:com.workday/bsvc Name"`
	Inactive                                    *bool                                            `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	Description                                 string                                           `xml:"urn:com.workday/bsvc Description,omitempty"`
	TypeReference                               []RecruitingAgencyTypeObjectType                 `xml:"urn:com.workday/bsvc Type_Reference,omitempty"`
	SuperiorReference                           *RecruitingAgencySiteObjectType                  `xml:"urn:com.workday/bsvc Superior_Reference,omitempty"`
	ContactInformationData                      *ContactInformationDataType                      `xml:"urn:com.workday/bsvc Contact_Information_Data,omitempty"`
	RecruitingAgencyCareerSiteConfigurationData *RecruitingAgencyCareerSiteConfigurationDataType `xml:"urn:com.workday/bsvc Recruiting_Agency_Career_Site_Configuration_Data"`
	RecruitingAgencyPayoutTermsData             *RecruitingAgencyPayoutTermsDataType             `xml:"urn:com.workday/bsvc Recruiting_Agency_Payout_Terms_Data,omitempty"`
	RecruitingAgencyAttachmentData              []RecruitingAgencyAttachmentDataType             `xml:"urn:com.workday/bsvc Recruiting_Agency_Attachment_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingAgencyFeeTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingAgencyFeeTypeObjectType struct {
	ID         []RecruitingAgencyFeeTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Recruiting Agency Payout Terms Data
type RecruitingAgencyPayoutTermsDataType struct {
	RecruitingAgencyFeeTypeReference *RecruitingAgencyFeeTypeObjectType `xml:"urn:com.workday/bsvc Recruiting_Agency_Fee_Type_Reference"`
	FeeValue                         float64                            `xml:"urn:com.workday/bsvc Fee_Value"`
	CurrencyReference                *CurrencyObjectType                `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
}

// Utilize the Request Criteria element to filter instance(s) of Recruiting Agency Additional Data by Name
type RecruitingAgencyRequestCriteriaType struct {
	Name string `xml:"urn:com.workday/bsvc Name,omitempty"`
}

// Recruiting Agency Reference.
type RecruitingAgencyRequestReferencesType struct {
	RecruitingAgencyReference []RecruitingAgencySiteObjectType `xml:"urn:com.workday/bsvc Recruiting_Agency_Reference"`
}

// Wrapper Element for Recruiting Agency Data
type RecruitingAgencyResponseDataType struct {
	RecruitingAgency []RecruitingAgencyType `xml:"urn:com.workday/bsvc Recruiting_Agency,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingAgencySiteObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingAgencySiteObjectType struct {
	ID         []RecruitingAgencySiteObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains details about a Recruiting Agency Site
type RecruitingAgencyType struct {
	RecruitingAgencyReference *RecruitingAgencySiteObjectType `xml:"urn:com.workday/bsvc Recruiting_Agency_Reference,omitempty"`
	RecruitingAgencyData      []RecruitingAgencyDataType      `xml:"urn:com.workday/bsvc Recruiting_Agency_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingAgencyTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingAgencyTypeObjectType struct {
	ID         []RecruitingAgencyTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Recruiting Agency User Data Wrapper
type RecruitingAgencyUserAttachmentDataType struct {
	AttachmentData *AttachmentWWSDataType `xml:"urn:com.workday/bsvc Attachment_Data"`
}

// Recruiting Agency User Data
type RecruitingAgencyUserDataType struct {
	RecruitingAgencyReference          []RecruitingAgencySiteObjectType         `xml:"urn:com.workday/bsvc Recruiting_Agency_Reference"`
	Inactive                           *bool                                    `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	NameData                           *LegalNameDataType                       `xml:"urn:com.workday/bsvc Name_Data,omitempty"`
	ContactInformationData             *ContactInformationDataType              `xml:"urn:com.workday/bsvc Contact_Information_Data,omitempty"`
	LinkedInURL                        string                                   `xml:"urn:com.workday/bsvc LinkedIn_URL,omitempty"`
	TwitterUserName                    string                                   `xml:"urn:com.workday/bsvc Twitter_User_Name,omitempty"`
	FacebookURL                        string                                   `xml:"urn:com.workday/bsvc Facebook_URL,omitempty"`
	GoogleURL                          string                                   `xml:"urn:com.workday/bsvc Google__URL,omitempty"`
	RecruitingAgencyUserAttachmentData []RecruitingAgencyUserAttachmentDataType `xml:"urn:com.workday/bsvc Recruiting_Agency_User_Attachment_Data,omitempty"`
	ID                                 string                                   `xml:"urn:com.workday/bsvc ID,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingAgencyUserObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingAgencyUserObjectType struct {
	ID         []RecruitingAgencyUserObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains Recruiting Agency User references.
type RecruitingAgencyUserRequestReferencesType struct {
	RecruitingAgencyUserReference []RecruitingAgencyUserObjectType `xml:"urn:com.workday/bsvc Recruiting_Agency_User_Reference"`
}

// Contains Recruiting Agency User data.
type RecruitingAgencyUserResponseDataType struct {
	RecruitingAgencyUser []RecruitingAgencyUserType `xml:"urn:com.workday/bsvc Recruiting_Agency_User,omitempty"`
}

// Recruiting Agency Additional Data response group
type RecruitingAgencyUserResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Contains Recruiting Agency User data.
type RecruitingAgencyUserType struct {
	RecruitingAgencyUserReference *RecruitingAgencyUserObjectType `xml:"urn:com.workday/bsvc Recruiting_Agency_User_Reference,omitempty"`
	RecruitingAgencyUserData      *RecruitingAgencyUserDataType   `xml:"urn:com.workday/bsvc Recruiting_Agency_User_Data,omitempty"`
}

// Container for Assessment Category Data
type RecruitingAssessmentCategoryDataType struct {
	ID                     string `xml:"urn:com.workday/bsvc ID"`
	AssessmentCategoryName string `xml:"urn:com.workday/bsvc Assessment_Category_Name"`
	Description            string `xml:"urn:com.workday/bsvc Description,omitempty"`
	Inactive               *bool  `xml:"urn:com.workday/bsvc Inactive,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingAssessmentCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingAssessmentCategoryObjectType struct {
	ID         []RecruitingAssessmentCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Container with reference to Recruiting Assessment Category
type RecruitingAssessmentCategoryRequestReferencesType struct {
	RecruitingAssessmentCategoryReference []RecruitingAssessmentCategoryObjectType `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Reference"`
}

// Recruiting Assessment Category Response Data
type RecruitingAssessmentCategoryResponseDataType struct {
	RecruitingAssessmentCategory []RecruitingAssessmentCategoryType `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category,omitempty"`
}

// Container for Determining whether or not to include reference
type RecruitingAssessmentCategoryResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing data for the Recruiting Assessment Category Security Segment.
type RecruitingAssessmentCategorySecuritySegmentDataType struct {
	ID                                              string                                   `xml:"urn:com.workday/bsvc ID,omitempty"`
	RecruitingAssessmentCategorySecuritySegmentName string                                   `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Security_Segment_Name"`
	RecruitingAssessmentCategoryReference           []RecruitingAssessmentCategoryObjectType `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Reference"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingAssessmentCategorySecuritySegmentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingAssessmentCategorySecuritySegmentObjectType struct {
	ID         []RecruitingAssessmentCategorySecuritySegmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing filtering logic for the Recruiting Assessment Category Security Segments.
type RecruitingAssessmentCategorySecuritySegmentRequestCriteriaType struct {
}

// Element containing references to Recruiting Assessment Category Security Segments to retrieve.
type RecruitingAssessmentCategorySecuritySegmentRequestReferencesType struct {
	RecruitingAssessmentCategorySecuritySegmentReference []RecruitingAssessmentCategorySecuritySegmentObjectType `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Security_Segment_Reference"`
}

// Element containing data for the Recruiting Assessment Category Security Segment.
type RecruitingAssessmentCategorySecuritySegmentResponseDataType struct {
	RecruitingAssessmentCategorySecuritySegment []RecruitingAssessmentCategorySecuritySegmentType `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Security_Segment,omitempty"`
}

// Element containing Recruiting Assessment Category Security Segment response grouping options.
type RecruitingAssessmentCategorySecuritySegmentResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing data for the Recruiting Assessment Category Security Segment.
type RecruitingAssessmentCategorySecuritySegmentType struct {
	RecruitingAssessmentCategorySecuritySegmentReference *RecruitingAssessmentCategorySecuritySegmentObjectType `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Security_Segment_Reference,omitempty"`
	RecruitingAssessmentCategorySecuritySegmentData      *RecruitingAssessmentCategorySecuritySegmentDataType   `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Security_Segment_Data,omitempty"`
}

// Recruiting Assessment Category Response Container
type RecruitingAssessmentCategoryType struct {
	RecruitingAssessmentCategoryReference *RecruitingAssessmentCategoryObjectType `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Reference,omitempty"`
	RecruitingAssessmentCategoryData      []RecruitingAssessmentCategoryDataType  `xml:"urn:com.workday/bsvc Recruiting_Assessment_Category_Data,omitempty"`
}

// The data used to create or modify a Candidate Assessment.
type RecruitingAssessmentDataType struct {
	AssessCandidateReference      string                                   `xml:"urn:com.workday/bsvc Assess_Candidate_Reference,omitempty"`
	AssessedByReference           *WorkerObjectType                        `xml:"urn:com.workday/bsvc Assessed_By_Reference,omitempty"`
	AssessedOnDate                time.Time                                `xml:"urn:com.workday/bsvc Assessed_On_Date"`
	AssessmentStatusReference     *RecruitingAssessmentStatusObjectType    `xml:"urn:com.workday/bsvc Assessment_Status_Reference"`
	OverallAssessmentComment      string                                   `xml:"urn:com.workday/bsvc Overall_Assessment_Comment,omitempty"`
	IsInlineAssessment            *bool                                    `xml:"urn:com.workday/bsvc Is_Inline_Assessment,omitempty"`
	AssessmentTestURL             string                                   `xml:"urn:com.workday/bsvc Assessment_Test_URL,omitempty"`
	AssessCandidateTestResultData []RecruitingAssessmentTestResultDataType `xml:"urn:com.workday/bsvc Assess_Candidate_Test_Result_Data,omitempty"`
}

func (t *RecruitingAssessmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RecruitingAssessmentDataType
	var layout struct {
		*T
		AssessedOnDate *xsdDate `xml:"urn:com.workday/bsvc Assessed_On_Date"`
	}
	layout.T = (*T)(t)
	layout.AssessedOnDate = (*xsdDate)(&layout.T.AssessedOnDate)
	return e.EncodeElement(layout, start)
}
func (t *RecruitingAssessmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RecruitingAssessmentDataType
	var overlay struct {
		*T
		AssessedOnDate *xsdDate `xml:"urn:com.workday/bsvc Assessed_On_Date"`
	}
	overlay.T = (*T)(t)
	overlay.AssessedOnDate = (*xsdDate)(&overlay.T.AssessedOnDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type RecruitingAssessmentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingAssessmentObjectType struct {
	ID         []RecruitingAssessmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingAssessmentStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingAssessmentStatusObjectType struct {
	ID         []RecruitingAssessmentStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingAssessmentTestObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingAssessmentTestObjectType struct {
	ID         []RecruitingAssessmentTestObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// A collection of information on the Assessment Test Results used in the Candidate Assessment.
type RecruitingAssessmentTestResultDataType struct {
	AssessmentTestResultReference *RecruitingAssessmentTestResultObjectType `xml:"urn:com.workday/bsvc Assessment_Test_Result_Reference,omitempty"`
	AssessmentTestReference       *RecruitingAssessmentTestObjectType       `xml:"urn:com.workday/bsvc Assessment_Test_Reference"`
	AssessmentTestScore           float64                                   `xml:"urn:com.workday/bsvc Assessment_Test_Score,omitempty"`
	AssessmentTestStatusReference *RecruitingAssessmentStatusObjectType     `xml:"urn:com.workday/bsvc Assessment_Test_Status_Reference,omitempty"`
	AssessmentTestDate            *time.Time                                `xml:"urn:com.workday/bsvc Assessment_Test_Date,omitempty"`
	Comment                       string                                    `xml:"urn:com.workday/bsvc Comment,omitempty"`
	AssessmentTestResultsURL      string                                    `xml:"urn:com.workday/bsvc Assessment_Test_Results_URL,omitempty"`
}

func (t *RecruitingAssessmentTestResultDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RecruitingAssessmentTestResultDataType
	var layout struct {
		*T
		AssessmentTestDate *xsdDate `xml:"urn:com.workday/bsvc Assessment_Test_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssessmentTestDate = (*xsdDate)(layout.T.AssessmentTestDate)
	return e.EncodeElement(layout, start)
}
func (t *RecruitingAssessmentTestResultDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RecruitingAssessmentTestResultDataType
	var overlay struct {
		*T
		AssessmentTestDate *xsdDate `xml:"urn:com.workday/bsvc Assessment_Test_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssessmentTestDate = (*xsdDate)(overlay.T.AssessmentTestDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type RecruitingAssessmentTestResultObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingAssessmentTestResultObjectType struct {
	ID         []RecruitingAssessmentTestResultObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for Recruiting Information.
type RecruitingDataType struct {
	ApplicantEnteredDate                *time.Time                          `xml:"urn:com.workday/bsvc Applicant_Entered_Date,omitempty"`
	ApplicantComments                   string                              `xml:"urn:com.workday/bsvc Applicant_Comments,omitempty"`
	EligibleForHireReference            *CommonYesNoObjectType              `xml:"urn:com.workday/bsvc Eligible_For_Hire_Reference,omitempty"`
	EligibleforRehireComments           string                              `xml:"urn:com.workday/bsvc Eligible_for_Rehire_Comments,omitempty"`
	ApplicantHasMarkedasNoShowReference *CommonBooleanEnumerationObjectType `xml:"urn:com.workday/bsvc Applicant_Has_Marked_as_No_Show_Reference,omitempty"`
	ApplicantSourceReference            *ApplicantSourceObjectType          `xml:"urn:com.workday/bsvc Applicant_Source_Reference,omitempty"`
	ReferredbyWorkerReference           []WorkerObjectType                  `xml:"urn:com.workday/bsvc Referred_by_Worker_Reference,omitempty"`
	PositionsConsideredforReference     []PositionGroupObjectType           `xml:"urn:com.workday/bsvc Positions_Considered_for_Reference,omitempty"`
}

func (t *RecruitingDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RecruitingDataType
	var layout struct {
		*T
		ApplicantEnteredDate *xsdDate `xml:"urn:com.workday/bsvc Applicant_Entered_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ApplicantEnteredDate = (*xsdDate)(layout.T.ApplicantEnteredDate)
	return e.EncodeElement(layout, start)
}
func (t *RecruitingDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RecruitingDataType
	var overlay struct {
		*T
		ApplicantEnteredDate *xsdDate `xml:"urn:com.workday/bsvc Applicant_Entered_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ApplicantEnteredDate = (*xsdDate)(overlay.T.ApplicantEnteredDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type RecruitingDispositionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingDispositionObjectType struct {
	ID         []RecruitingDispositionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingEventObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingEventObjectType struct {
	ID         []RecruitingEventObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for Job Requisition Recruiting Instructions or posting instructions.
type RecruitingInstructionDataType struct {
	RecruitingInstructionReference *RecruitingInstructionObjectType `xml:"urn:com.workday/bsvc Recruiting_Instruction_Reference,omitempty"`
	DoNotSendToRecruitingSystem    *bool                            `xml:"urn:com.workday/bsvc Do_Not_Send_To_Recruiting_System,omitempty"`
	Name                           string                           `xml:"urn:com.workday/bsvc Name,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingInstructionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingInstructionObjectType struct {
	ID         []RecruitingInstructionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains Prospect Data: Level, Status, and Type
type RecruitingProspectDataType struct {
	Prospect                  *bool                               `xml:"urn:com.workday/bsvc Prospect,omitempty"`
	Confidential              *bool                               `xml:"urn:com.workday/bsvc Confidential,omitempty"`
	ReferralConsentGiven      *bool                               `xml:"urn:com.workday/bsvc Referral_Consent_Given,omitempty"`
	LevelReference            *ManagementLevelObjectType          `xml:"urn:com.workday/bsvc Level_Reference,omitempty"`
	ProspectStatusReference   *RecruitingProspectStatusObjectType `xml:"urn:com.workday/bsvc Prospect_Status_Reference,omitempty"`
	ProspectTypeReference     *RecruitingProspectTypeObjectType   `xml:"urn:com.workday/bsvc Prospect_Type_Reference,omitempty"`
	SourceReference           *ApplicantSourceObjectType          `xml:"urn:com.workday/bsvc Source_Reference,omitempty"`
	ReferredByWorkerReference *WorkerObjectType                   `xml:"urn:com.workday/bsvc Referred_By_Worker_Reference,omitempty"`
	AddedByWorkerReference    *ProcessmaintainedRoleObjectType    `xml:"urn:com.workday/bsvc Added_By_Worker_Reference,omitempty"`
	ProspectAttachmentData    *ProspectAttachmentsDataType        `xml:"urn:com.workday/bsvc Prospect_Attachment_Data,omitempty"`
	ResumeData                *CandidateResumeDataType            `xml:"urn:com.workday/bsvc Resume_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingProspectStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingProspectStatusObjectType struct {
	ID         []RecruitingProspectStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingProspectTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingProspectTypeObjectType struct {
	ID         []RecruitingProspectTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing data for the Recruiting Self-Schedule Calendar.
type RecruitingSelfScheduleCalendarDataType struct {
	ID                                         string                                          `xml:"urn:com.workday/bsvc ID,omitempty"`
	Name                                       string                                          `xml:"urn:com.workday/bsvc Name"`
	RecruitingSelfScheduleCalendarExternalName string                                          `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_External_Name,omitempty"`
	CalendarTypeReference                      *RecruitingSelfScheduleCalendarTypeObjectType   `xml:"urn:com.workday/bsvc Calendar_Type_Reference"`
	Duration                                   float64                                         `xml:"urn:com.workday/bsvc Duration,omitempty"`
	EarliestVisibility                         float64                                         `xml:"urn:com.workday/bsvc Earliest_Visibility,omitempty"`
	LatestVisibility                           float64                                         `xml:"urn:com.workday/bsvc Latest_Visibility,omitempty"`
	MinimumNumberofDatestoDisplay              float64                                         `xml:"urn:com.workday/bsvc Minimum_Number_of_Dates_to_Display,omitempty"`
	RescheduleLockHours                        float64                                         `xml:"urn:com.workday/bsvc Reschedule_Lock_Hours,omitempty"`
	Inactive                                   *bool                                           `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	Schedules                                  []RecruitingSelfServiceScheduleType             `xml:"urn:com.workday/bsvc Schedules,omitempty"`
	ExcludedDays                               []RecruitingSelfServiceScheduleExcludedDaysType `xml:"urn:com.workday/bsvc Excluded_Days,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingSelfScheduleCalendarObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingSelfScheduleCalendarObjectType struct {
	ID         []RecruitingSelfScheduleCalendarObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The "Request Criteria" element for the web service operation contains the filtering logic to limit the data that is returned in the response.
type RecruitingSelfScheduleCalendarRequestCriteriaType struct {
}

// References to Recruiting Self-Schedule Calendars to retrieve.
type RecruitingSelfScheduleCalendarRequestReferencesType struct {
	RecruitingSelfScheduleCalendarReference []RecruitingSelfScheduleCalendarObjectType `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_Reference"`
}

// Element containing data for the Recruiting Self-Schedule Calendar.
type RecruitingSelfScheduleCalendarResponseDataType struct {
	RecruitingSelfScheduleCalendar []RecruitingSelfScheduleCalendarType `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar,omitempty"`
}

// The Response Group allows the request to specify which data attributes should be returned in the Response, such as whether to include reference elements, attachments, etc.
type RecruitingSelfScheduleCalendarResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing data for the Recruiting Self-Schedule Calendar.
type RecruitingSelfScheduleCalendarType struct {
	RecruitingSelfScheduleCalendarReference *RecruitingSelfScheduleCalendarObjectType `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_Reference,omitempty"`
	RecruitingSelfScheduleCalendarData      *RecruitingSelfScheduleCalendarDataType   `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_Data,omitempty"`
}

// Element containing data for the Recruiting Self-Schedule Calendar Type.
type RecruitingSelfScheduleCalendarTypeDataType struct {
	ID                                                 string                                 `xml:"urn:com.workday/bsvc ID,omitempty"`
	Order                                              string                                 `xml:"urn:com.workday/bsvc Order,omitempty"`
	CalendarTypeName                                   string                                 `xml:"urn:com.workday/bsvc Calendar_Type_Name"`
	MapstoRecruitingStageReference                     *RecruitingStageWorkdayOwnedObjectType `xml:"urn:com.workday/bsvc Maps_to_Recruiting_Stage_Reference"`
	CandidateSelfScheduleTaskMessage                   string                                 `xml:"urn:com.workday/bsvc Candidate_Self-Schedule_Task_Message"`
	CandidateMessageWhenNoTimeSlotsareAvailable        string                                 `xml:"urn:com.workday/bsvc Candidate_Message_When_No_Time_Slots_are_Available"`
	CandidateMessageWhenNoDatesorTimesWorkforCandidate string                                 `xml:"urn:com.workday/bsvc Candidate_Message_When_No_Dates_or_Times_Work_for_Candidate"`
	Inactive                                           *bool                                  `xml:"urn:com.workday/bsvc Inactive,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingSelfScheduleCalendarTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingSelfScheduleCalendarTypeObjectType struct {
	ID         []RecruitingSelfScheduleCalendarTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing filtering logic for the Recruiting Self-Schedule Calendar Types.
type RecruitingSelfScheduleCalendarTypeRequestCriteriaType struct {
}

// Element containing references to Recruiting Self-Schedule Calendar Types to retrieve.
type RecruitingSelfScheduleCalendarTypeRequestReferencesType struct {
	RecruitingSelfScheduleCalendarTypeReference []RecruitingSelfScheduleCalendarTypeObjectType `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_Type_Reference"`
}

// Element containing data for the Recruiting Self-Schedule Calendar Type.
type RecruitingSelfScheduleCalendarTypeResponseDataType struct {
	RecruitingSelfScheduleCalendarType []RecruitingSelfScheduleCalendarTypeType `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_Type,omitempty"`
}

// Element containing Recruiting Self-Schedule Calendar Type response grouping options.
type RecruitingSelfScheduleCalendarTypeResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing data for the Recruiting Self-Schedule Calendar Type.
type RecruitingSelfScheduleCalendarTypeType struct {
	RecruitingSelfScheduleCalendarTypeReference *RecruitingSelfScheduleCalendarTypeObjectType `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_Type_Reference,omitempty"`
	RecruitingSelfScheduleCalendarTypeData      *RecruitingSelfScheduleCalendarTypeDataType   `xml:"urn:com.workday/bsvc Recruiting_Self-Schedule_Calendar_Type_Data,omitempty"`
}

// Section containing data for the Recruiting Calendar Schedule Role Assignments.
type RecruitingSelfServiceScheduleAssignRoleDataType struct {
	OrganizationRoleReference *AssignableRoleObjectType `xml:"urn:com.workday/bsvc Organization_Role_Reference"`
	RoleAssigneeReference     []RoleeObjectType         `xml:"urn:com.workday/bsvc Role_Assignee_Reference,omitempty"`
	Delete                    bool                      `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	Update                    bool                      `xml:"urn:com.workday/bsvc Update,attr,omitempty"`
}

// Section containing data for the Recruiting Calendar Schedule.
type RecruitingSelfServiceScheduleDataType struct {
	StartDate                    time.Time                                         `xml:"urn:com.workday/bsvc Start_Date"`
	EndDate                      *time.Time                                        `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	WeeklyRecurrence             float64                                           `xml:"urn:com.workday/bsvc Weekly_Recurrence"`
	MaxNumberofParticipants      float64                                           `xml:"urn:com.workday/bsvc Max_Number_of_Participants"`
	LocationReference            *LocationObjectType                               `xml:"urn:com.workday/bsvc Location_Reference"`
	TimeZoneReference            *TimeZoneObjectType                               `xml:"urn:com.workday/bsvc Time_Zone_Reference"`
	CandidateConfirmationMessage string                                            `xml:"urn:com.workday/bsvc Candidate_Confirmation_Message"`
	AssignRoleData               []RecruitingSelfServiceScheduleAssignRoleDataType `xml:"urn:com.workday/bsvc Assign_Role_Data,omitempty"`
	TimeBlocks                   []RecruitingSelfServiceScheduleTimeBlockType      `xml:"urn:com.workday/bsvc Time_Blocks,omitempty"`
}

func (t *RecruitingSelfServiceScheduleDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RecruitingSelfServiceScheduleDataType
	var layout struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(&layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *RecruitingSelfServiceScheduleDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RecruitingSelfServiceScheduleDataType
	var overlay struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(&overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Section containing data for the Recruiting Calendar Excluded Days setting.
type RecruitingSelfServiceScheduleExcludedDaysDataType struct {
	StartDate                    time.Time                   `xml:"urn:com.workday/bsvc Start_Date"`
	EndDate                      *time.Time                  `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	MonthReference               []MonthObjectType           `xml:"urn:com.workday/bsvc Month_Reference"`
	DaysoftheMonthReference      []DayoftheMonthObjectType   `xml:"urn:com.workday/bsvc Days_of_the_Month_Reference,omitempty"`
	DayoftheWeekinMonthReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Day_of_the_Week_in_Month_Reference,omitempty"`
}

func (t *RecruitingSelfServiceScheduleExcludedDaysDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RecruitingSelfServiceScheduleExcludedDaysDataType
	var layout struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(&layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *RecruitingSelfServiceScheduleExcludedDaysDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RecruitingSelfServiceScheduleExcludedDaysDataType
	var overlay struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(&overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Section to add, update or delete Excluded Days for Recruiting Calendars.
type RecruitingSelfServiceScheduleExcludedDaysType struct {
	ExcludedDaysReference *RecruitingSelfServiceTimeBlockObjectType          `xml:"urn:com.workday/bsvc Excluded_Days_Reference,omitempty"`
	ExcludedDaysData      *RecruitingSelfServiceScheduleExcludedDaysDataType `xml:"urn:com.workday/bsvc Excluded_Days_Data,omitempty"`
	AddOnly               bool                                               `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Delete                bool                                               `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingSelfServiceScheduleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingSelfServiceScheduleObjectType struct {
	ID         []RecruitingSelfServiceScheduleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Section containing data for the Recruiting Calendar Schedule Time Block.
type RecruitingSelfServiceScheduleTimeBlockDataType struct {
	DayoftheWeekReference *DayoftheWeekObjectType `xml:"urn:com.workday/bsvc Day_of_the_Week_Reference"`
	StartTime             time.Time               `xml:"urn:com.workday/bsvc Start_Time"`
	EndTime               time.Time               `xml:"urn:com.workday/bsvc End_Time"`
}

func (t *RecruitingSelfServiceScheduleTimeBlockDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RecruitingSelfServiceScheduleTimeBlockDataType
	var layout struct {
		*T
		StartTime *xsdTime `xml:"urn:com.workday/bsvc Start_Time"`
		EndTime   *xsdTime `xml:"urn:com.workday/bsvc End_Time"`
	}
	layout.T = (*T)(t)
	layout.StartTime = (*xsdTime)(&layout.T.StartTime)
	layout.EndTime = (*xsdTime)(&layout.T.EndTime)
	return e.EncodeElement(layout, start)
}
func (t *RecruitingSelfServiceScheduleTimeBlockDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RecruitingSelfServiceScheduleTimeBlockDataType
	var overlay struct {
		*T
		StartTime *xsdTime `xml:"urn:com.workday/bsvc Start_Time"`
		EndTime   *xsdTime `xml:"urn:com.workday/bsvc End_Time"`
	}
	overlay.T = (*T)(t)
	overlay.StartTime = (*xsdTime)(&overlay.T.StartTime)
	overlay.EndTime = (*xsdTime)(&overlay.T.EndTime)
	return d.DecodeElement(&overlay, &start)
}

// Section to add, update or delete Time Blocks for the Recruiting Calendar Schedule.
type RecruitingSelfServiceScheduleTimeBlockType struct {
	TimeBlockReference *RecruitingSelfServiceTimeBlockObjectType       `xml:"urn:com.workday/bsvc Time_Block_Reference,omitempty"`
	TimeBlockData      *RecruitingSelfServiceScheduleTimeBlockDataType `xml:"urn:com.workday/bsvc Time_Block_Data,omitempty"`
	AddOnly            bool                                            `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Delete             bool                                            `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Section to add, update or delete Schedules for Recruiting Calendars.
type RecruitingSelfServiceScheduleType struct {
	ScheduleReference *RecruitingSelfServiceScheduleObjectType `xml:"urn:com.workday/bsvc Schedule_Reference,omitempty"`
	ScheduleData      *RecruitingSelfServiceScheduleDataType   `xml:"urn:com.workday/bsvc Schedule_Data,omitempty"`
	AddOnly           bool                                     `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Delete            bool                                     `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingSelfServiceTimeBlockObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingSelfServiceTimeBlockObjectType struct {
	ID         []RecruitingSelfServiceTimeBlockObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RecruitingStageWorkdayOwnedObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RecruitingStageWorkdayOwnedObjectType struct {
	ID         []RecruitingStageWorkdayOwnedObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains information about a Attachment to add to a Candidate.
type ReferaCandidateAttachmentDataType struct {
	ResumeAttachment *AttachmentWWSDataType `xml:"urn:com.workday/bsvc Resume_Attachment,omitempty"`
}

// Contains all the request data to Refer a Candidate.
type ReferaCandidateRequestType struct {
	BusinessProcessParameters       *BusinessProcessParametersType       `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	ReferralCandidateSubmissionData *ReferralCandidateSubmissionDataType `xml:"urn:com.workday/bsvc Referral_Candidate_Submission_Data"`
	Version                         string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response for Refer a Candidate
type ReferaCandidateResponseType struct {
	ReferralEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Referral_Event_Reference,omitempty"`
	Version                string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains referral submission data for Refer a Candidate process.
type ReferralCandidateSubmissionDataType struct {
	ReferredByWorkerReference *WorkerObjectType                     `xml:"urn:com.workday/bsvc Referred_By_Worker_Reference"`
	NameData                  *CandidateNameDataType                `xml:"urn:com.workday/bsvc Name_Data,omitempty"`
	ContactData               *CandidateContactDataType             `xml:"urn:com.workday/bsvc Contact_Data,omitempty"`
	ReferralDetailsData       *ReferralDetailsDataType              `xml:"urn:com.workday/bsvc Referral_Details_Data,omitempty"`
	SocialMediaAccountData    []CandidateSocialMediaAccountDataType `xml:"urn:com.workday/bsvc Social_Media_Account_Data,omitempty"`
	ResumeAttachmentData      []ReferaCandidateAttachmentDataType   `xml:"urn:com.workday/bsvc Resume_Attachment_Data,omitempty"`
}

// Contains referral details data for a candidate.
type ReferralDetailsDataType struct {
	JobRequisitionReference       []JobRequisitionEnabledObjectType `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	JobAreasReference             []JobFamilyObjectType             `xml:"urn:com.workday/bsvc Job_Areas_Reference,omitempty"`
	ReferralRelationshipReference *ReferralRelationshipObjectType   `xml:"urn:com.workday/bsvc Referral_Relationship_Reference,omitempty"`
	Comment                       string                            `xml:"urn:com.workday/bsvc Comment,omitempty"`
	ReferralConsentGiven          *bool                             `xml:"urn:com.workday/bsvc Referral_Consent_Given,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ReferralRelationshipObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ReferralRelationshipObjectType struct {
	ID         []ReferralRelationshipObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// contains relative names for a relative of a person
type RelativeNameDataType struct {
	RelativeNameReferenceReference *RelativeNameReferenceObjectType `xml:"urn:com.workday/bsvc Relative_Name_Reference_Reference,omitempty"`
	RelativeTypeReference          *RelativeTypeObjectType          `xml:"urn:com.workday/bsvc Relative_Type_Reference,omitempty"`
	NameDetailData                 *PersonNameDetailDataType        `xml:"urn:com.workday/bsvc Name_Detail_Data,omitempty"`
	Delete                         bool                             `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Wrapper element for all relative names of a person
type RelativeNameInformationDataType struct {
	RelativeNameData []RelativeNameDataType `xml:"urn:com.workday/bsvc Relative_Name_Data,omitempty"`
	ReplaceAll       bool                   `xml:"urn:com.workday/bsvc Replace_All,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RelativeNameReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RelativeNameReferenceObjectType struct {
	ID         []RelativeNameReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RelativeTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RelativeTypeObjectType struct {
	ID         []RelativeTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ReligionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ReligionObjectType struct {
	ID         []ReligionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for the Request Default Compensation sub business process.
type RequestCompensationDefaultSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	CompensationDefaultData      *CompensationDefaultDataType      `xml:"urn:com.workday/bsvc Compensation_Default_Data,omitempty"`
}

// Grant an employee a one-time payment such as a referral, spot bonus, or severance pay. Uses the Request One Time Payment business event through the web service.
// Can be skipped, processed automatically or processed manually.
type RequestOneTimePaymentSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType         `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	RequestOneTimePaymentData    []OneTimePaymentDataType                  `xml:"urn:com.workday/bsvc Request_One_Time_Payment_Data,omitempty"`
	EffectiveDate                *time.Time                                `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	ReasonReference              *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
}

func (t *RequestOneTimePaymentSubBusinessProcessType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RequestOneTimePaymentSubBusinessProcessType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *RequestOneTimePaymentSubBusinessProcessType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RequestOneTimePaymentSubBusinessProcessType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper for Request Requisition Compensation sub business process.
type RequestRequisitionCompensationSubProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	RequisitionCompensationData  *RequisitionCompensationDataType  `xml:"urn:com.workday/bsvc Requisition_Compensation_Data,omitempty"`
}

// Data element for the stock grant request
type RequestStockGrantOfferDataType struct {
	EffectiveDate      time.Time                                 `xml:"urn:com.workday/bsvc Effective_Date"`
	StockPlanReference *StockPlanObjectType                      `xml:"urn:com.workday/bsvc Stock_Plan_Reference"`
	ReasonReference    *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
	SharesGranted      float64                                   `xml:"urn:com.workday/bsvc Shares_Granted"`
	GrantPercent       float64                                   `xml:"urn:com.workday/bsvc Grant_Percent"`
	GrantAmount        float64                                   `xml:"urn:com.workday/bsvc Grant_Amount"`
}

func (t *RequestStockGrantOfferDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RequestStockGrantOfferDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *RequestStockGrantOfferDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RequestStockGrantOfferDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for the Request Stock Grant sub business process.
type RequestStockOfferSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters"`
	RequestStockGrantData        *RequestStockGrantOfferDataType   `xml:"urn:com.workday/bsvc Request_Stock_Grant_Data,omitempty"`
}

// Data element that contains the requisition compensation information.
type RequisitionCompensationDataType struct {
	CompensatableGuidelinesData   *CompensatableDefaultGuidelinesDataType               `xml:"urn:com.workday/bsvc Compensatable_Guidelines_Data,omitempty"`
	PayPlanData                   *ProposedBasePayPlanAssignmentContainerDataType       `xml:"urn:com.workday/bsvc Pay_Plan_Data,omitempty"`
	UnitSalaryPlanData            *ProposedSalaryUnitPlanAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Unit_Salary_Plan_Data,omitempty"`
	AllowancePlanNonUnitBasedData *ProposedAllowancePlanAssignmentContainerDataType     `xml:"urn:com.workday/bsvc Allowance_Plan_Non-Unit_Based_Data,omitempty"`
	AllowancePlanUnitBasedData    *ProposedAllowanceUnitPlanAssignmentContainerDataType `xml:"urn:com.workday/bsvc Allowance_Plan_Unit_Based_Data,omitempty"`
	BonusPlanData                 *ProposedBonusPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Bonus_Plan_Data,omitempty"`
	MeritPlanData                 *ProposedMeritPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Merit_Plan_Data,omitempty"`
	CommissionPlanData            *ProposedCommissionPlanAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Commission_Plan_Data,omitempty"`
	StockPlanData                 *ProposedStockPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Stock_Plan_Data,omitempty"`
	PeriodSalaryPlanData          *ProposedPeriodSalaryPlanAssignmentContainerDataType  `xml:"urn:com.workday/bsvc Period_Salary_Plan_Data,omitempty"`
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

// Contains the responsibility profile's information for a position.
type ResponsibilityProfileforJobDataType struct {
	Responsibility               string                               `xml:"urn:com.workday/bsvc Responsibility,omitempty"`
	Required                     *bool                                `xml:"urn:com.workday/bsvc Required,omitempty"`
	QualificationSourceReference *SkillQualificationEnabledObjectType `xml:"urn:com.workday/bsvc Qualification_Source_Reference,omitempty"`
}

// Contains the position's responsibility information.
type ResponsibilityProfileforJobType struct {
	ResponsibilityProfileReference *UniqueIdentifierObjectType          `xml:"urn:com.workday/bsvc Responsibility_Profile_Reference,omitempty"`
	ResponsibilityProfileData      *ResponsibilityProfileforJobDataType `xml:"urn:com.workday/bsvc Responsibility_Profile_Data,omitempty"`
}

// Replacement Element containing Responsibility Qualifications for the Job Profile.
// When updating a Job Profile, all Responsibilities for the Job Profile will be replaced by the data in being submitted. If no data is submitted, then the existing Responsibilities are not changed.
type ResponsibilityQualificationProfileReplacementDataType struct {
	ResponsibilityDescription string `xml:"urn:com.workday/bsvc Responsibility_Description"`
	Required                  *bool  `xml:"urn:com.workday/bsvc Required,omitempty"`
}

// Wrapper element for Responsibility Qualifications. Allows all responsibility qualifications for a Job Profile or Position Restriction to be removed - or to replace all existing qualifications with those sent in the web service.
type ResponsibilityQualificationReplacementType struct {
	ResponsibilityQualificationReplacementData []ResponsibilityQualificationProfileReplacementDataType `xml:"urn:com.workday/bsvc Responsibility_Qualification_Replacement_Data,omitempty"`
	Delete                                     bool                                                    `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ResumeAttachmentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ResumeAttachmentObjectType struct {
	ID         []ResumeAttachmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RoleAssignerObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RoleAssignerObjectType struct {
	ID         []RoleAssignerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for a Organization Role Assignment
type RoleAssignmentEffectiveDataType struct {
	RoleAssignmentReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Role_Assignment_Reference,omitempty"`
	EffectiveDate           *time.Time                  `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	RoleReference           *AssignableRoleObjectType   `xml:"urn:com.workday/bsvc Role_Reference,omitempty"`
	RoleAssigneeReference   []RoleeObjectType           `xml:"urn:com.workday/bsvc Role_Assignee_Reference,omitempty"`
}

func (t *RoleAssignmentEffectiveDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RoleAssignmentEffectiveDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *RoleAssignmentEffectiveDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RoleAssignmentEffectiveDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating Element for all Organization Role Assignment data.
type RoleAssignmentType struct {
	RoleAssignmentReference *UniqueIdentifierObjectType         `xml:"urn:com.workday/bsvc Role_Assignment_Reference,omitempty"`
	RoleAssignmentData      *OrganizationRoleAssignmentDataType `xml:"urn:com.workday/bsvc Role_Assignment_Data"`
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

// Contains a unique identifier for an instance of an object.
type RoleeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RoleeObjectType struct {
	ID         []RoleeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Details of the Salary over the Cap allocation, (e.g., a set of allocation dimensions and percentages)
type SalaryOvertheCapCostingAllocationDetailDataType struct {
	SOCOrder                             string                             `xml:"urn:com.workday/bsvc SOC_Order"`
	SOCDefaultfromOrganizationAssignment bool                               `xml:"urn:com.workday/bsvc SOC_Default_from_Organization_Assignment"`
	SOCOverrideWorktagReference          []TenantedPayrollWorktagObjectType `xml:"urn:com.workday/bsvc SOC_Override_Worktag_Reference"`
	SOCDistributionPercent               float64                            `xml:"urn:com.workday/bsvc SOC_Distribution_Percent"`
}

// Contains a unique identifier for an instance of an object.
type SalaryPayPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SalaryPayPlanObjectType struct {
	ID         []SalaryPayPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SalaryUnitPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SalaryUnitPlanObjectType struct {
	ID         []SalaryUnitPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SchoolObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SchoolObjectType struct {
	ID         []SchoolObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SchoolTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SchoolTypeObjectType struct {
	ID         []SchoolTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SelfIdentificationofDisabilityStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SelfIdentificationofDisabilityStatusObjectType struct {
	ID         []SelfIdentificationofDisabilityStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Utilize the following element to retrieve the current DateTime of the Workday system.
type ServerTimestampGetType struct {
	Version string `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing the current DateTime of the Workday system.
type ServerTimestampType struct {
	ServerTimestampData []time.Time `xml:"urn:com.workday/bsvc Server_Timestamp_Data,omitempty"`
	Version             string      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

//func (t *ServerTimestampType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
//	type T ServerTimestampType
//	var layout struct {
//		*T
//		ServerTimestampData *xsdDateTime `xml:"urn:com.workday/bsvc Server_Timestamp_Data,omitempty"`
//	}
//	layout.T = (*T)(t)
//	layout.ServerTimestampData = (*xsdDateTime)(&layout.T.ServerTimestampData)
//	return e.EncodeElement(layout, start)
//}
//func (t *ServerTimestampType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
//	type T ServerTimestampType
//	var overlay struct {
//		*T
//		ServerTimestampData *xsdDateTime `xml:"urn:com.workday/bsvc Server_Timestamp_Data,omitempty"`
//	}
//	overlay.T = (*T)(t)
//	overlay.ServerTimestampData = (*xsdDateTime)(&overlay.T.ServerTimestampData)
//	return d.DecodeElement(&overlay, &start)
//}

// Contains a unique identifier for an instance of an object.
type SexualOrientationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SexualOrientationObjectType struct {
	ID         []SexualOrientationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SkillItemMixedObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SkillItemMixedObjectType struct {
	ID         []SkillItemMixedObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SkillItemTenantedObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SkillItemTenantedObjectType struct {
	ID         []SkillItemTenantedObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the Skill profile's information for a position.
type SkillProfileforJobDataType struct {
	SkillReference               *SkillItemMixedObjectType            `xml:"urn:com.workday/bsvc Skill_Reference,omitempty"`
	Name                         string                               `xml:"urn:com.workday/bsvc Name,omitempty"`
	Required                     *bool                                `xml:"urn:com.workday/bsvc Required,omitempty"`
	QualificationSourceReference *SkillQualificationEnabledObjectType `xml:"urn:com.workday/bsvc Qualification_Source_Reference,omitempty"`
}

// Contains the Position's Skill information.
type SkillProfileforJobType struct {
	SkillProfileReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Skill_Profile_Reference,omitempty"`
	SkillProfileData      *SkillProfileforJobDataType `xml:"urn:com.workday/bsvc Skill_Profile_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SkillQualificationEnabledObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SkillQualificationEnabledObjectType struct {
	ID         []SkillQualificationEnabledObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for Skill Qualifications. Allows all Skill Qualifications for a Job Profile or Position Restriction to be removed - or to replace all existing Skill Qualifications with those sent in the web service.
type SkillQualificationProfileReplacementDataType struct {
	SkillReference *SkillItemTenantedObjectType `xml:"urn:com.workday/bsvc Skill_Reference,omitempty"`
	Name           string                       `xml:"urn:com.workday/bsvc Name,omitempty"`
	Required       *bool                        `xml:"urn:com.workday/bsvc Required,omitempty"`
}

// Wrapper element for Skill Qualifications. Allows all Skill Qualifications for a Job Profile or Position Restriction to be removed - or to replace all existing Skill Qualifications with those sent in the web service.
type SkillQualificationReplacementType struct {
	SkillQualificationReplacementData []SkillQualificationProfileReplacementDataType `xml:"urn:com.workday/bsvc Skill_Qualification_Replacement_Data,omitempty"`
	Delete                            bool                                           `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SocialBenefitsLocalityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SocialBenefitsLocalityObjectType struct {
	ID         []SocialBenefitsLocalityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SocialNetworkMetaTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SocialNetworkMetaTypeObjectType struct {
	ID         []SocialNetworkMetaTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This element contains data about each Specialty-Subspecialty combination associated with the certification achievement
type SpecialtyAchievementDataType struct {
	SpecialtyReference    *SpecialtyParentObjectType `xml:"urn:com.workday/bsvc Specialty_Reference,omitempty"`
	SubspecialtyReference []SpecialtyChildObjectType `xml:"urn:com.workday/bsvc Subspecialty_Reference,omitempty"`
}

// This element contains data about each Specialty-Subspecialty combination associated with the certification achievement
type SpecialtyAchievementDatawithDatesType struct {
	SpecialtyReference    *SpecialtyParentObjectType `xml:"urn:com.workday/bsvc Specialty_Reference,omitempty"`
	StartDate             *time.Time                 `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate               *time.Time                 `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	SubspecialtyReference []SpecialtyChildObjectType `xml:"urn:com.workday/bsvc Subspecialty_Reference,omitempty"`
}

func (t *SpecialtyAchievementDatawithDatesType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T SpecialtyAchievementDatawithDatesType
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
func (t *SpecialtyAchievementDatawithDatesType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SpecialtyAchievementDatawithDatesType
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

// Contains a unique identifier for an instance of an object.
type SpecialtyChildObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SpecialtyChildObjectType struct {
	ID         []SpecialtyChildObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SpecialtyParentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SpecialtyParentObjectType struct {
	ID         []SpecialtyParentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SpendCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SpendCategoryObjectType struct {
	ID         []SpendCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StaticCandidatePoolObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StaticCandidatePoolObjectType struct {
	ID         []StaticCandidatePoolObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StockPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StockPlanObjectType struct {
	ID         []StockPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StudentAwardSourceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StudentAwardSourceObjectType struct {
	ID         []StudentAwardSourceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains all the request data to submit a candidate via recruiting agency candidate process.
type SubmitRecruitingAgencyCandidateRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType               `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	CandidateSubmissionData   *RecruitingAgencyCandidateSubmissionDataType `xml:"urn:com.workday/bsvc Candidate_Submission_Data"`
	Version                   string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for Submit Recruiting Agency Candidate  Response.
type SubmitRecruitingAgencyCandidateResponseType struct {
	RecruitingAgencyCandidateEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Recruiting_Agency_Candidate_Event_Reference,omitempty"`
	Version                                 string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type SubscriberObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SubscriberObjectType struct {
	ID         []SubscriberObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains details specific to a Supervisory Organization.
type SupervisoryOrgDataType struct {
	StaffingModel               string                                   `xml:"urn:com.workday/bsvc Staffing_Model,omitempty"`
	LocationReference           *LocationObjectType                      `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	OrganizationAssignmentsData []OrganizationAssignmentsDataType        `xml:"urn:com.workday/bsvc Organization_Assignments_Data,omitempty"`
	StaffingRestrictionsData    *PositionGroupRestrictionSummaryDataType `xml:"urn:com.workday/bsvc Staffing_Restrictions_Data,omitempty"`
	AvailableForHire            *bool                                    `xml:"urn:com.workday/bsvc Available_For_Hire,omitempty"`
	HiringFreeze                *bool                                    `xml:"urn:com.workday/bsvc Hiring_Freeze,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type SupplierObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SupplierObjectType struct {
	ID         []SupplierObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Data for Background Check Package Test results.
type TestReferenceDataType struct {
	BackgroundCheckPackageTestReference *BackgroundCheckPackageTestObjectType `xml:"urn:com.workday/bsvc Background_Check_Package_Test_Reference"`
	StatusReference                     *BackgroundCheckStatusObjectType      `xml:"urn:com.workday/bsvc Status_Reference"`
	TestName                            string                                `xml:"urn:com.workday/bsvc Test_Name,attr,omitempty"`
	TestDescription                     string                                `xml:"urn:com.workday/bsvc Test_Description,attr,omitempty"`
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

// Wrapper element for Training information.
type TrainingAchievementDataType struct {
	TrainingID            string                  `xml:"urn:com.workday/bsvc Training_ID,omitempty"`
	RemoveTraining        *bool                   `xml:"urn:com.workday/bsvc Remove_Training,omitempty"`
	Training              string                  `xml:"urn:com.workday/bsvc Training,omitempty"`
	Description           string                  `xml:"urn:com.workday/bsvc Description,omitempty"`
	TrainingTypeReference *TrainingTypeObjectType `xml:"urn:com.workday/bsvc Training_Type_Reference,omitempty"`
	CompletionDate        *time.Time              `xml:"urn:com.workday/bsvc Completion_Date,omitempty"`
	TrainingDuration      string                  `xml:"urn:com.workday/bsvc Training_Duration,omitempty"`
}

func (t *TrainingAchievementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T TrainingAchievementDataType
	var layout struct {
		*T
		CompletionDate *xsdDate `xml:"urn:com.workday/bsvc Completion_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CompletionDate = (*xsdDate)(layout.T.CompletionDate)
	return e.EncodeElement(layout, start)
}
func (t *TrainingAchievementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T TrainingAchievementDataType
	var overlay struct {
		*T
		CompletionDate *xsdDate `xml:"urn:com.workday/bsvc Completion_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CompletionDate = (*xsdDate)(overlay.T.CompletionDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type TrainingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TrainingObjectType struct {
	ID         []TrainingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the training profile's information for a position.
type TrainingProfileforJobDataType struct {
	TrainingName                 string                               `xml:"urn:com.workday/bsvc Training_Name,omitempty"`
	Description                  string                               `xml:"urn:com.workday/bsvc Description,omitempty"`
	TrainingTypeReference        *TrainingTypeObjectType              `xml:"urn:com.workday/bsvc Training_Type_Reference,omitempty"`
	Required                     *bool                                `xml:"urn:com.workday/bsvc Required,omitempty"`
	QualificationSourceReference *SkillQualificationEnabledObjectType `xml:"urn:com.workday/bsvc Qualification_Source_Reference,omitempty"`
}

// Contains the position's training information.
type TrainingProfileforJobType struct {
	TrainingProfileReference *UniqueIdentifierObjectType    `xml:"urn:com.workday/bsvc Training_Profile_Reference,omitempty"`
	TrainingProfileData      *TrainingProfileforJobDataType `xml:"urn:com.workday/bsvc Training_Profile_Data,omitempty"`
}

// Replacement element containing Training Qualifications for the Job Profile
// When updating a Job Profile, all Training Qualifications for the Job Profile will be replaced by the submitted data. If no data is submitted, then the existing Training Qualifications are not changed.
type TrainingQualificationProfileReplacementDataType struct {
	TrainingName          string                  `xml:"urn:com.workday/bsvc Training_Name,omitempty"`
	Description           string                  `xml:"urn:com.workday/bsvc Description,omitempty"`
	TrainingTypeReference *TrainingTypeObjectType `xml:"urn:com.workday/bsvc Training_Type_Reference,omitempty"`
	Required              *bool                   `xml:"urn:com.workday/bsvc Required,omitempty"`
}

// Wrapper element for Training Qualifications. Allows all training qualifications for a Job Profile or Position Restriction to be removed - or to replace all existing training qualifications with those sent in the web service.
type TrainingQualificationReplacementType struct {
	TrainingQualificationReplacementData []TrainingQualificationProfileReplacementDataType `xml:"urn:com.workday/bsvc Training_Qualification_Replacement_Data,omitempty"`
	Delete                               bool                                              `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Wrapper element for Training information.
type TrainingType struct {
	TrainingReference *TrainingObjectType           `xml:"urn:com.workday/bsvc Training_Reference,omitempty"`
	TrainingData      []TrainingAchievementDataType `xml:"urn:com.workday/bsvc Training_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TrainingTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TrainingTypeObjectType struct {
	ID         []TrainingTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Searches the transaction log for business processes and event lites to find specific events that occurred during a period of time.
type TransactionLogCriteriaType struct {
	TransactionDateRangeData  *EffectiveAndUpdatedDateTimeDataType `xml:"urn:com.workday/bsvc Transaction_Date_Range_Data,omitempty"`
	TransactionTypeReferences *TransactionTypeReferencesType       `xml:"urn:com.workday/bsvc Transaction_Type_References,omitempty"`
	SubscriberReference       *SubscriberObjectType                `xml:"urn:com.workday/bsvc Subscriber_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TransactionLogTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TransactionLogTypeObjectType struct {
	ID         []TransactionLogTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element containing Transaction Types.
type TransactionTypeReferencesType struct {
	TransactionTypeReference []TransactionLogTypeObjectType `xml:"urn:com.workday/bsvc Transaction_Type_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type UnionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type UnionObjectType struct {
	ID         []UnionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Data for unposting a job.
type UnpostJobBusinessProcessDataType struct {
	JobPostingReference *JobPostingObjectType        `xml:"urn:com.workday/bsvc Job_Posting_Reference"`
	JobPostingSiteData  *JobPostingReferenceDataType `xml:"urn:com.workday/bsvc Job_Posting_Site_Data"`
}

// Unposts a Jop Posting.
type UnpostJobRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType    `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	UnpostJobData             *UnpostJobBusinessProcessDataType `xml:"urn:com.workday/bsvc Unpost_Job_Data"`
	Version                   string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response for Unpost Job Request
type UnpostJobResponseType struct {
	EventReference      *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	JobPostingReference []JobPostingObjectType      `xml:"urn:com.workday/bsvc Job_Posting_Reference,omitempty"`
	Version             string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Run the Update Job Posting sub business process to unpost the job postings associated with this Job Requisition.
type UnpostJobSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
}

// Run an Update Job Posting Event.
type UpdateJobPostingBusinessProcessDataType struct {
	JobPostingReference     *JobPostingObjectType        `xml:"urn:com.workday/bsvc Job_Posting_Reference"`
	JobPostingReferenceData *JobPostingReferenceDataType `xml:"urn:com.workday/bsvc Job_Posting_Reference_Data"`
	JobPostingData          []UpdateJobPostingDataType   `xml:"urn:com.workday/bsvc Job_Posting_Data"`
}

// Data for individual Job Postings.
type UpdateJobPostingDataType struct {
	JobPostingStartDate time.Time  `xml:"urn:com.workday/bsvc Job_Posting_Start_Date"`
	JobPostingEndDate   *time.Time `xml:"urn:com.workday/bsvc Job_Posting_End_Date,omitempty"`
}

func (t *UpdateJobPostingDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T UpdateJobPostingDataType
	var layout struct {
		*T
		JobPostingStartDate *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_Start_Date"`
		JobPostingEndDate   *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.JobPostingStartDate = (*xsdDate)(&layout.T.JobPostingStartDate)
	layout.JobPostingEndDate = (*xsdDate)(layout.T.JobPostingEndDate)
	return e.EncodeElement(layout, start)
}
func (t *UpdateJobPostingDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T UpdateJobPostingDataType
	var overlay struct {
		*T
		JobPostingStartDate *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_Start_Date"`
		JobPostingEndDate   *xsdDate `xml:"urn:com.workday/bsvc Job_Posting_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.JobPostingStartDate = (*xsdDate)(&overlay.T.JobPostingStartDate)
	overlay.JobPostingEndDate = (*xsdDate)(overlay.T.JobPostingEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Run an Update Job Posting Event.
type UpdateJobPostingRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType           `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	UpdateJobPostingData      *UpdateJobPostingBusinessProcessDataType `xml:"urn:com.workday/bsvc Update_Job_Posting_Data"`
	Version                   string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Event ID for the Job Posting Event.
type UpdateJobPostingResponseType struct {
	EventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	JobPostingData []JobPostingDataType        `xml:"urn:com.workday/bsvc Job_Posting_Data,omitempty"`
	Version        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type UserLanguageObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type UserLanguageObjectType struct {
	ID         []UserLanguageObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type ValidationErrorType struct {
	Message       string `xml:"urn:com.workday/bsvc Message,omitempty"`
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
	Xpath         string `xml:"urn:com.workday/bsvc Xpath,omitempty"`
}

type ValidationFaultType struct {
	ValidationError []ValidationErrorType `xml:"urn:com.workday/bsvc Validation_Error,omitempty"`
}

// Contains ID, Order, Veteran Status Name, Description, Inactive, and Global Setup Data Mapping.
type VeteranStatusDataType struct {
	ID                              string                            `xml:"urn:com.workday/bsvc ID,omitempty"`
	Order                           string                            `xml:"urn:com.workday/bsvc Order,omitempty"`
	VeteranStatusName               string                            `xml:"urn:com.workday/bsvc Veteran_Status_Name,omitempty"`
	Description                     string                            `xml:"urn:com.workday/bsvc Description,omitempty"`
	Inactive                        *bool                             `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	GlobalSetupDataMappingReference *GlobalSetupDataMappingObjectType `xml:"urn:com.workday/bsvc Global_Setup_Data_Mapping_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type VeteranStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type VeteranStatusObjectType struct {
	ID         []VeteranStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Utilize the following criteria options to search for Veteran Statuses within the Workday system. The Veteran Status references that are returned are those that satisfy ALL criteria included in the request. Therefore, the result set will become more limited with every criterium that is populated.
type VeteranStatusRequestCriteriaType struct {
}

// Contains Veteran Status references.
type VeteranStatusRequestReferencesType struct {
	VeteranStatusReference []VeteranStatusObjectType `xml:"urn:com.workday/bsvc Veteran_Status_Reference"`
}

// Contains Veteran Status data.
type VeteranStatusResponseDataType struct {
	VeteranStatus []VeteranStatusType `xml:"urn:com.workday/bsvc Veteran_Status,omitempty"`
}

// Contains Veteran Status data.
type VeteranStatusType struct {
	VeteranStatusReference *VeteranStatusObjectType `xml:"urn:com.workday/bsvc Veteran_Status_Reference,omitempty"`
	VeteranStatusData      []VeteranStatusDataType  `xml:"urn:com.workday/bsvc Veteran_Status_Data,omitempty"`
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

// Wrapper element for the work experience.
type WorkExperienceDataType struct {
	ExperienceReference       *WorkExperienceSkillObjectType  `xml:"urn:com.workday/bsvc Experience_Reference,omitempty"`
	RemoveExperience          *bool                           `xml:"urn:com.workday/bsvc Remove_Experience,omitempty"`
	ExperienceRatingReference *WorkExperienceRatingObjectType `xml:"urn:com.workday/bsvc Experience_Rating_Reference,omitempty"`
	ExperienceComment         string                          `xml:"urn:com.workday/bsvc Experience_Comment,omitempty"`
}

// Contains the work experience profile's information for a position.
type WorkExperienceProfileforJobDataType struct {
	WorkExperienceReference       *WorkExperienceSkillObjectType       `xml:"urn:com.workday/bsvc Work_Experience_Reference,omitempty"`
	WorkExperienceRatingReference *WorkExperienceRatingObjectType      `xml:"urn:com.workday/bsvc Work_Experience_Rating_Reference,omitempty"`
	Required                      *bool                                `xml:"urn:com.workday/bsvc Required,omitempty"`
	QualificationSourceReference  *SkillQualificationEnabledObjectType `xml:"urn:com.workday/bsvc Qualification_Source_Reference,omitempty"`
}

// Replacement element containing Work Experience Qualifications for the Job Profile
// When updating a Job Profile, all Work Experiences for the Job Profile will be replaced by the submitted data. If no data is submitted, then the existing Work Experiences are not changed.
type WorkExperienceQualificationProfileReplacementDataType struct {
	WorkerExperienceReference     *WorkExperienceSkillObjectType  `xml:"urn:com.workday/bsvc Worker_Experience_Reference"`
	WorkExperienceRatingReference *WorkExperienceRatingObjectType `xml:"urn:com.workday/bsvc Work_Experience_Rating_Reference,omitempty"`
	Required                      *bool                           `xml:"urn:com.workday/bsvc Required,omitempty"`
}

// Wrapper element for Work Experience Qualifications. Allows all work experience qualifications for a Job Profile or Position Restriction to be removed - or to replace all existing work experience qualifications with those sent in the web service.
type WorkExperienceQualificationReplacementType struct {
	WorkExperienceQualificationReplacementData []WorkExperienceQualificationProfileReplacementDataType `xml:"urn:com.workday/bsvc Work_Experience_Qualification_Replacement_Data,omitempty"`
	Delete                                     bool                                                    `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WorkExperienceRatingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkExperienceRatingObjectType struct {
	ID         []WorkExperienceRatingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WorkExperienceSkillObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkExperienceSkillObjectType struct {
	ID         []WorkExperienceSkillObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the position's work experience information.
type WorkExperienceforJobType struct {
	WorkerExperienceProfileReference *UniqueIdentifierObjectType          `xml:"urn:com.workday/bsvc Worker_Experience_Profile_Reference,omitempty"`
	WorkExperienceData               *WorkExperienceProfileforJobDataType `xml:"urn:com.workday/bsvc Work_Experience_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WorkShiftObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkShiftObjectType struct {
	ID         []WorkShiftObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type WorkdayCommonHeaderType struct {
	IncludeReferenceDescriptorsInResponse *bool `xml:"urn:com.workday/bsvc Include_Reference_Descriptors_In_Response,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WorkerDocumentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkerDocumentObjectType struct {
	ID         []WorkerDocumentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WorkerEventObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkerEventObjectType struct {
	ID         []WorkerEventObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for  Worker Filling Position.
type WorkerForFilledPositionType struct {
	WorkerReference   *WorkerObjectType   `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	PositionReference *PositionObjectType `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
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
type WorkerTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkerTypeObjectType struct {
	ID         []WorkerTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
