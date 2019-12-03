package staffing

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// Contains a unique identifier for an instance of an object.
type AcademicAffiliateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AcademicAffiliateObjectType struct {
	ID         []AcademicAffiliateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AcademicAppointeeEnabledObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AcademicAppointeeEnabledObjectType struct {
	ID         []AcademicAppointeeEnabledObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AcademicAppointmentIdentifierObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AcademicAppointmentIdentifierObjectType struct {
	ID         []AcademicAppointmentIdentifierObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing the specifics for the academic appointment being  added or updated. An academic appointment tracks an academic worker’s time at a university or college.
type AcademicAppointmentSnapshotDataType struct {
	AcademicAppointmentTrackReference *AcademicAppointmentTrackObjectType      `xml:"urn:com.workday/bsvc Academic_Appointment_Track_Reference,omitempty"`
	AppointmentTrackID                string                                   `xml:"urn:com.workday/bsvc Appointment_Track_ID,omitempty"`
	TrackTypeReference                *AcademicTrackTypeObjectType             `xml:"urn:com.workday/bsvc Track_Type_Reference,omitempty"`
	AppointmentIdentifierReference    *AcademicAppointmentIdentifierObjectType `xml:"urn:com.workday/bsvc Appointment_Identifier_Reference"`
	PositionReference                 *PositionElementObjectType               `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	AcademicUnitReference             *AcademicUnitObjectType                  `xml:"urn:com.workday/bsvc Academic_Unit_Reference"`
	RosterPercent                     float64                                  `xml:"urn:com.workday/bsvc Roster_Percent,omitempty"`
	AppointmentStartDate              time.Time                                `xml:"urn:com.workday/bsvc Appointment_Start_Date"`
	AppointmentEndDate                *time.Time                               `xml:"urn:com.workday/bsvc Appointment_End_Date,omitempty"`
	TrackStartDateOverride            *time.Time                               `xml:"urn:com.workday/bsvc Track_Start_Date_Override,omitempty"`
	RankReference                     *AcademicRankObjectType                  `xml:"urn:com.workday/bsvc Rank_Reference,omitempty"`
	NamedProfessorshipReference       *NamedProfessorshipObjectType            `xml:"urn:com.workday/bsvc Named_Professorship_Reference,omitempty"`
	AppointmentSpecialtyReference     *AppointmentSpecialtyObjectType          `xml:"urn:com.workday/bsvc Appointment_Specialty_Reference,omitempty"`
	ConstructedTitle                  string                                   `xml:"urn:com.workday/bsvc Constructed_Title,omitempty"`
	AppointmentTitle                  string                                   `xml:"urn:com.workday/bsvc Appointment_Title"`
	AdjustedTitleStartDate            *time.Time                               `xml:"urn:com.workday/bsvc Adjusted_Title_Start_Date,omitempty"`
	AcademicReviewDate                *time.Time                               `xml:"urn:com.workday/bsvc Academic_Review_Date,omitempty"`
	RelatedAcademicUnitReference      *AcademicUnitObjectType                  `xml:"urn:com.workday/bsvc Related_Academic_Unit_Reference,omitempty"`
	TenureHomeReference               *AcademicUnitObjectType                  `xml:"urn:com.workday/bsvc Tenure_Home_Reference,omitempty"`
	TenureStatusReference             *AcademicTenureStatusObjectType          `xml:"urn:com.workday/bsvc Tenure_Status_Reference,omitempty"`
	ProbationaryEndDate               *time.Time                               `xml:"urn:com.workday/bsvc Probationary_End_Date,omitempty"`
	TenureAwardDate                   *time.Time                               `xml:"urn:com.workday/bsvc Tenure_Award_Date,omitempty"`
}

func (t *AcademicAppointmentSnapshotDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AcademicAppointmentSnapshotDataType
	var layout struct {
		*T
		AppointmentStartDate   *xsdDate `xml:"urn:com.workday/bsvc Appointment_Start_Date"`
		AppointmentEndDate     *xsdDate `xml:"urn:com.workday/bsvc Appointment_End_Date,omitempty"`
		TrackStartDateOverride *xsdDate `xml:"urn:com.workday/bsvc Track_Start_Date_Override,omitempty"`
		AdjustedTitleStartDate *xsdDate `xml:"urn:com.workday/bsvc Adjusted_Title_Start_Date,omitempty"`
		AcademicReviewDate     *xsdDate `xml:"urn:com.workday/bsvc Academic_Review_Date,omitempty"`
		ProbationaryEndDate    *xsdDate `xml:"urn:com.workday/bsvc Probationary_End_Date,omitempty"`
		TenureAwardDate        *xsdDate `xml:"urn:com.workday/bsvc Tenure_Award_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AppointmentStartDate = (*xsdDate)(&layout.T.AppointmentStartDate)
	layout.AppointmentEndDate = (*xsdDate)(layout.T.AppointmentEndDate)
	layout.TrackStartDateOverride = (*xsdDate)(layout.T.TrackStartDateOverride)
	layout.AdjustedTitleStartDate = (*xsdDate)(layout.T.AdjustedTitleStartDate)
	layout.AcademicReviewDate = (*xsdDate)(layout.T.AcademicReviewDate)
	layout.ProbationaryEndDate = (*xsdDate)(layout.T.ProbationaryEndDate)
	layout.TenureAwardDate = (*xsdDate)(layout.T.TenureAwardDate)
	return e.EncodeElement(layout, start)
}
func (t *AcademicAppointmentSnapshotDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AcademicAppointmentSnapshotDataType
	var overlay struct {
		*T
		AppointmentStartDate   *xsdDate `xml:"urn:com.workday/bsvc Appointment_Start_Date"`
		AppointmentEndDate     *xsdDate `xml:"urn:com.workday/bsvc Appointment_End_Date,omitempty"`
		TrackStartDateOverride *xsdDate `xml:"urn:com.workday/bsvc Track_Start_Date_Override,omitempty"`
		AdjustedTitleStartDate *xsdDate `xml:"urn:com.workday/bsvc Adjusted_Title_Start_Date,omitempty"`
		AcademicReviewDate     *xsdDate `xml:"urn:com.workday/bsvc Academic_Review_Date,omitempty"`
		ProbationaryEndDate    *xsdDate `xml:"urn:com.workday/bsvc Probationary_End_Date,omitempty"`
		TenureAwardDate        *xsdDate `xml:"urn:com.workday/bsvc Tenure_Award_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AppointmentStartDate = (*xsdDate)(&overlay.T.AppointmentStartDate)
	overlay.AppointmentEndDate = (*xsdDate)(overlay.T.AppointmentEndDate)
	overlay.TrackStartDateOverride = (*xsdDate)(overlay.T.TrackStartDateOverride)
	overlay.AdjustedTitleStartDate = (*xsdDate)(overlay.T.AdjustedTitleStartDate)
	overlay.AcademicReviewDate = (*xsdDate)(overlay.T.AcademicReviewDate)
	overlay.ProbationaryEndDate = (*xsdDate)(overlay.T.ProbationaryEndDate)
	overlay.TenureAwardDate = (*xsdDate)(overlay.T.TenureAwardDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type AcademicAppointmentTrackObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AcademicAppointmentTrackObjectType struct {
	ID         []AcademicAppointmentTrackObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AcademicCurricularDivisionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AcademicCurricularDivisionObjectType struct {
	ID         []AcademicCurricularDivisionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AcademicLevelObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AcademicLevelObjectType struct {
	ID         []AcademicLevelObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AcademicPayPeriodObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AcademicPayPeriodObjectType struct {
	ID         []AcademicPayPeriodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element that contains position's academic pay setup related data.
type AcademicPaySetupDataType struct {
	AnnualWorkPeriodWorkPercentofYear float64    `xml:"urn:com.workday/bsvc Annual_Work_Period_Work_Percent_of_Year,omitempty"`
	AnnualWorkPeriodStartDate         *time.Time `xml:"urn:com.workday/bsvc Annual_Work_Period_Start_Date,omitempty"`
	AnnualWorkPeriodEndDate           *time.Time `xml:"urn:com.workday/bsvc Annual_Work_Period_End_Date,omitempty"`
	DisbursementPlanPeriodStartDate   *time.Time `xml:"urn:com.workday/bsvc Disbursement_Plan_Period_Start_Date,omitempty"`
	DisbursementPlanPeriodEndDate     *time.Time `xml:"urn:com.workday/bsvc Disbursement_Plan_Period_End_Date,omitempty"`
}

func (t *AcademicPaySetupDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AcademicPaySetupDataType
	var layout struct {
		*T
		AnnualWorkPeriodStartDate       *xsdDate `xml:"urn:com.workday/bsvc Annual_Work_Period_Start_Date,omitempty"`
		AnnualWorkPeriodEndDate         *xsdDate `xml:"urn:com.workday/bsvc Annual_Work_Period_End_Date,omitempty"`
		DisbursementPlanPeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Disbursement_Plan_Period_Start_Date,omitempty"`
		DisbursementPlanPeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Disbursement_Plan_Period_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AnnualWorkPeriodStartDate = (*xsdDate)(layout.T.AnnualWorkPeriodStartDate)
	layout.AnnualWorkPeriodEndDate = (*xsdDate)(layout.T.AnnualWorkPeriodEndDate)
	layout.DisbursementPlanPeriodStartDate = (*xsdDate)(layout.T.DisbursementPlanPeriodStartDate)
	layout.DisbursementPlanPeriodEndDate = (*xsdDate)(layout.T.DisbursementPlanPeriodEndDate)
	return e.EncodeElement(layout, start)
}
func (t *AcademicPaySetupDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AcademicPaySetupDataType
	var overlay struct {
		*T
		AnnualWorkPeriodStartDate       *xsdDate `xml:"urn:com.workday/bsvc Annual_Work_Period_Start_Date,omitempty"`
		AnnualWorkPeriodEndDate         *xsdDate `xml:"urn:com.workday/bsvc Annual_Work_Period_End_Date,omitempty"`
		DisbursementPlanPeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Disbursement_Plan_Period_Start_Date,omitempty"`
		DisbursementPlanPeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Disbursement_Plan_Period_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AnnualWorkPeriodStartDate = (*xsdDate)(overlay.T.AnnualWorkPeriodStartDate)
	overlay.AnnualWorkPeriodEndDate = (*xsdDate)(overlay.T.AnnualWorkPeriodEndDate)
	overlay.DisbursementPlanPeriodStartDate = (*xsdDate)(overlay.T.DisbursementPlanPeriodStartDate)
	overlay.DisbursementPlanPeriodEndDate = (*xsdDate)(overlay.T.DisbursementPlanPeriodEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type AcademicPeriodObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AcademicPeriodObjectType struct {
	ID         []AcademicPeriodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AcademicPeriodTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AcademicPeriodTypeObjectType struct {
	ID         []AcademicPeriodTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AcademicRankObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AcademicRankObjectType struct {
	ID         []AcademicRankObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AcademicTenureStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AcademicTenureStatusObjectType struct {
	ID         []AcademicTenureStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AcademicTrackTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AcademicTrackTypeObjectType struct {
	ID         []AcademicTrackTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AcademicUnitObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AcademicUnitObjectType struct {
	ID         []AcademicUnitObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AchievableLevelObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AchievableLevelObjectType struct {
	ID         []AchievableLevelObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Wrapper element for the Add Academic Appointment business process.
type AddAcademicAppointmentDataType struct {
	ReasonReference                                                     *GeneralEventSubcategoryObjectType                        `xml:"urn:com.workday/bsvc Reason_Reference"`
	AddAcademicAffiliateStatus                                          *bool                                                     `xml:"urn:com.workday/bsvc Add_Academic_Affiliate_Status,omitempty"`
	AcademicAppointeeReference                                          *AcademicAppointeeEnabledObjectType                       `xml:"urn:com.workday/bsvc Academic_Appointee_Reference,omitempty"`
	CreateAcademicAffiliateData                                         *CreateAcademicAffiliateDataType                          `xml:"urn:com.workday/bsvc Create_Academic_Affiliate_Data,omitempty"`
	AcademicAppointmentData                                             *AcademicAppointmentSnapshotDataType                      `xml:"urn:com.workday/bsvc Academic_Appointment_Data"`
	CreateWorkdayAccountSubBusinessProcessforAcademicAffiliate          *CreateWorkdayAccountSubBusinessProcessType               `xml:"urn:com.workday/bsvc Create_Workday_Account_Sub_Business_Process_for_Academic_Affiliate,omitempty"`
	ManageProfessionalAffiliationSubBusinessProcessforAcademicAffiliate []ManageProfessionalAffiliationSubBusinessProcessDataType `xml:"urn:com.workday/bsvc Manage_Professional_Affiliation_Sub_Business_Process_for_Academic_Affiliate,omitempty"`
	ManageEducationSubBusinessProcessforAcademicAffiliate               []ManageEducationSubBusinessProcessDataType               `xml:"urn:com.workday/bsvc Manage_Education_Sub_Business_Process_for_Academic_Affiliate,omitempty"`
	ManageInstructorEligibilitySubBusinessProcessforAcademicAffiliate   []ManageInstructorEligibilitySubBusinessProcessDataType   `xml:"urn:com.workday/bsvc Manage_Instructor_Eligibility_Sub_Business_Process_for_Academic_Affiliate,omitempty"`
}

// Wrapper for the Add Academic Appointment sub business process that is part of the Hire and Edit Position business processes.
type AddAcademicAppointmentSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	AddAcademicAppointmentData   *AddAcademicAppointmentDataType   `xml:"urn:com.workday/bsvc Add_Academic_Appointment_Data,omitempty"`
}

// Wrapper Element for the Add Additional Job business process web service and its sub business processes.
type AddAdditionalJobDataType struct {
	EmployeeReference                               *EmployeeObjectType                                      `xml:"urn:com.workday/bsvc Employee_Reference"`
	OrganizationReference                           *SupervisoryOrganizationObjectType                       `xml:"urn:com.workday/bsvc Organization_Reference"`
	PositionReference                               *PositionRestrictionsObjectType                          `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	JobRequisitionReference                         *JobRequisitionObjectType                                `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	ExcludefromHeadcount                            *bool                                                    `xml:"urn:com.workday/bsvc Exclude_from_Headcount,omitempty"`
	AddAdditionalJobEventData                       *AddAdditionalJobEventDataType                           `xml:"urn:com.workday/bsvc Add_Additional_Job_Event_Data"`
	EditAssignOrganizationSubProcess                *EditAssignPositionOrganizationSubBusinessProcessType    `xml:"urn:com.workday/bsvc Edit_Assign_Organization_Sub_Process,omitempty"`
	AssignMatrixOrganizationSubProcess              *AssignMatrixOrganizationSubBusinessProcessType          `xml:"urn:com.workday/bsvc Assign_Matrix_Organization_Sub_Process,omitempty"`
	AssignPayGroupSubProcess                        *AssignPayGroupSubBusinessProcessType                    `xml:"urn:com.workday/bsvc Assign_Pay_Group_Sub_Process,omitempty"`
	ProposeCompensationforAdditionalJobSubProcess   *ProposeCompensationForEmploymentSubBusinessProcessType  `xml:"urn:com.workday/bsvc Propose_Compensation_for_Additional_Job_Sub_Process,omitempty"`
	CheckPositionBudgetSubProcess                   *CheckPositionBudgetSubBusinessProcessType               `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process,omitempty"`
	OnboardingSetupSubProcess                       *OnboardingSetupSubBusinessProcessType                   `xml:"urn:com.workday/bsvc Onboarding_Setup_Sub_Process,omitempty"`
	AssignCostingAllocationSubProcess               *AssignCostingAllocationSubBusinessProcessType           `xml:"urn:com.workday/bsvc Assign_Costing_Allocation_Sub_Process,omitempty"`
	SwitchPrimaryJobSubProcess                      *SwitchPrimaryJobSubProcessType                          `xml:"urn:com.workday/bsvc Switch_Primary_Job_Sub_Process,omitempty"`
	AddAcademicAppointmentSubProcess                *AddAcademicAppointmentSubBusinessProcessType            `xml:"urn:com.workday/bsvc Add_Academic_Appointment_Sub_Process,omitempty"`
	StudentEmploymentEligibilitySubProcess          *StudentEmploymentEligibilitySubBusinessProcessType      `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Sub_Process,omitempty"`
	ManageUnionMembershipSubProcess                 *ManageUnionMembershipSubBusinessProcessType             `xml:"urn:com.workday/bsvc Manage_Union_Membership_Sub_Process,omitempty"`
	AssignCollectiveAgreementSubProcess             *AssignEmployeeCollectiveAgreementSubBusinessProcessType `xml:"urn:com.workday/bsvc Assign_Collective_Agreement_Sub_Process,omitempty"`
	MaintainEmployeeContractsSubProcess             *MaintainEmployeeContractsSubBusinessProcessType         `xml:"urn:com.workday/bsvc Maintain_Employee_Contracts_Sub_Process,omitempty"`
	ManageEmployeeProbationPeriodSubBusinessProcess *ManageProbationPeriodSubBusinessProcessType             `xml:"urn:com.workday/bsvc Manage_Employee_Probation_Period_Sub_Business_Process,omitempty"`
	StartDate                                       time.Time                                                `xml:"urn:com.workday/bsvc Start_Date,attr,omitempty"`
}

func (t *AddAdditionalJobDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AddAdditionalJobDataType
	var layout struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(&layout.T.StartDate)
	return e.EncodeElement(layout, start)
}
func (t *AddAdditionalJobDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AddAdditionalJobDataType
	var overlay struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(&overlay.T.StartDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for the details of an Add Additional Job for an employee.
type AddAdditionalJobEventDataType struct {
	PositionID                   string                                    `xml:"urn:com.workday/bsvc Position_ID,omitempty"`
	AdditionalJobReasonReference *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Additional_Job_Reason_Reference,omitempty"`
	EmployeeTypeReference        *PositionWorkerTypeObjectType             `xml:"urn:com.workday/bsvc Employee_Type_Reference,omitempty"`
	FirstDayofWork               *time.Time                                `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
	ProbationStartDate           *time.Time                                `xml:"urn:com.workday/bsvc Probation_Start_Date,omitempty"`
	ProbationEndDate             *time.Time                                `xml:"urn:com.workday/bsvc Probation_End_Date,omitempty"`
	EndEmploymentDate            *time.Time                                `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
	ConversionPositionStartDate  *time.Time                                `xml:"urn:com.workday/bsvc Conversion_Position_Start_Date,omitempty"`
	PositionDetails              *PositionDetailsSubDataType               `xml:"urn:com.workday/bsvc Position_Details"`
}

func (t *AddAdditionalJobEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AddAdditionalJobEventDataType
	var layout struct {
		*T
		FirstDayofWork              *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
		ProbationStartDate          *xsdDate `xml:"urn:com.workday/bsvc Probation_Start_Date,omitempty"`
		ProbationEndDate            *xsdDate `xml:"urn:com.workday/bsvc Probation_End_Date,omitempty"`
		EndEmploymentDate           *xsdDate `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
		ConversionPositionStartDate *xsdDate `xml:"urn:com.workday/bsvc Conversion_Position_Start_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FirstDayofWork = (*xsdDate)(layout.T.FirstDayofWork)
	layout.ProbationStartDate = (*xsdDate)(layout.T.ProbationStartDate)
	layout.ProbationEndDate = (*xsdDate)(layout.T.ProbationEndDate)
	layout.EndEmploymentDate = (*xsdDate)(layout.T.EndEmploymentDate)
	layout.ConversionPositionStartDate = (*xsdDate)(layout.T.ConversionPositionStartDate)
	return e.EncodeElement(layout, start)
}
func (t *AddAdditionalJobEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AddAdditionalJobEventDataType
	var overlay struct {
		*T
		FirstDayofWork              *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
		ProbationStartDate          *xsdDate `xml:"urn:com.workday/bsvc Probation_Start_Date,omitempty"`
		ProbationEndDate            *xsdDate `xml:"urn:com.workday/bsvc Probation_End_Date,omitempty"`
		EndEmploymentDate           *xsdDate `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
		ConversionPositionStartDate *xsdDate `xml:"urn:com.workday/bsvc Conversion_Position_Start_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FirstDayofWork = (*xsdDate)(overlay.T.FirstDayofWork)
	overlay.ProbationStartDate = (*xsdDate)(overlay.T.ProbationStartDate)
	overlay.ProbationEndDate = (*xsdDate)(overlay.T.ProbationEndDate)
	overlay.EndEmploymentDate = (*xsdDate)(overlay.T.EndEmploymentDate)
	overlay.ConversionPositionStartDate = (*xsdDate)(overlay.T.ConversionPositionStartDate)
	return d.DecodeElement(&overlay, &start)
}

// Responds with the Event ID for the Add Additional Employee Job Event and the job reference.
type AddAdditionalJobEventResponseType struct {
	EventReference    *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	EmployeeReference *WorkerObjectType           `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	JobReference      *PositionElementObjectType  `xml:"urn:com.workday/bsvc Job_Reference,omitempty"`
	Version           string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This web service operation is designed to all an additional job to be added to an existing active employee using the Add Additional Job business process.
type AddAdditionalJobRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AddAdditionalJobData      *AddAdditionalJobDataType      `xml:"urn:com.workday/bsvc Add_Additional_Job_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper Element for the Add Retiree Status business process web service and its sub business processes.
type AddRetireeStatusDataType struct {
	EmployeeReference         *EmployeeObjectType            `xml:"urn:com.workday/bsvc Employee_Reference"`
	RetireeStatusDate         time.Time                      `xml:"urn:com.workday/bsvc Retiree_Status_Date"`
	AddRetireeStatusEventData *AddRetireeStatusEventDataType `xml:"urn:com.workday/bsvc Add_Retiree_Status_Event_Data"`
}

func (t *AddRetireeStatusDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AddRetireeStatusDataType
	var layout struct {
		*T
		RetireeStatusDate *xsdDate `xml:"urn:com.workday/bsvc Retiree_Status_Date"`
	}
	layout.T = (*T)(t)
	layout.RetireeStatusDate = (*xsdDate)(&layout.T.RetireeStatusDate)
	return e.EncodeElement(layout, start)
}
func (t *AddRetireeStatusDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AddRetireeStatusDataType
	var overlay struct {
		*T
		RetireeStatusDate *xsdDate `xml:"urn:com.workday/bsvc Retiree_Status_Date"`
	}
	overlay.T = (*T)(t)
	overlay.RetireeStatusDate = (*xsdDate)(&overlay.T.RetireeStatusDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains retirement reason and retiree organization references for Add Retiree Status.
type AddRetireeStatusEventDataType struct {
	ReasonReference              *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference"`
	RetireeOrganizationReference *RetireeOrganizationObjectType            `xml:"urn:com.workday/bsvc Retiree_Organization_Reference"`
}

// This web service operation is designed to add a terminated worker as a retiree into a retiree organization using the Add Retiree Status business process.
type AddRetireeStatusRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AddRetireeStatusData      *AddRetireeStatusDataType      `xml:"urn:com.workday/bsvc Add_Retiree_Status_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This element contains information in Add Retiree Status event.
type AddRetireeStatusResponseType struct {
	EventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element that contains details for invoking Add Retiree Status as sub process.
type AddRetireeStatusSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType      `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	AddRetireeStatusDetails      *AddRetireeStatusSubProcessDetailsType `xml:"urn:com.workday/bsvc Add_Retiree_Status_Details,omitempty"`
}

// Contains details for Add Retiree Status as sub process.
type AddRetireeStatusSubProcessDetailsType struct {
	RetirementDate               time.Time                                 `xml:"urn:com.workday/bsvc Retirement_Date"`
	RetirementReasonReference    *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Retirement_Reason_Reference"`
	RetireeOrganizationReference *RetireeOrganizationObjectType            `xml:"urn:com.workday/bsvc Retiree_Organization_Reference"`
}

func (t *AddRetireeStatusSubProcessDetailsType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AddRetireeStatusSubProcessDetailsType
	var layout struct {
		*T
		RetirementDate *xsdDate `xml:"urn:com.workday/bsvc Retirement_Date"`
	}
	layout.T = (*T)(t)
	layout.RetirementDate = (*xsdDate)(&layout.T.RetirementDate)
	return e.EncodeElement(layout, start)
}
func (t *AddRetireeStatusSubProcessDetailsType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AddRetireeStatusSubProcessDetailsType
	var overlay struct {
		*T
		RetirementDate *xsdDate `xml:"urn:com.workday/bsvc Retirement_Date"`
	}
	overlay.T = (*T)(t)
	overlay.RetirementDate = (*xsdDate)(&overlay.T.RetirementDate)
	return d.DecodeElement(&overlay, &start)
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

// Contains a unique identifier for an instance of an object.
type AndOrOperatorsObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AndOrOperatorsObjectType struct {
	ID         []AndOrOperatorsObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The annual contribution for the spending account election.
type AnnualContributionDataType struct {
	AnnualContributionAmount float64 `xml:"urn:com.workday/bsvc Annual_Contribution_Amount"`
	PriorContributionAmount  float64 `xml:"urn:com.workday/bsvc Prior_Contribution_Amount"`
	RemainingPeriods         float64 `xml:"urn:com.workday/bsvc Remaining_Periods"`
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

// Contains a unique identifier for an instance of an object.
type AppointmentSpecialtyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AppointmentSpecialtyObjectType struct {
	ID         []AppointmentSpecialtyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Wrapper to hold the business process configuration and data for Assign Employee Collective Agreement sub business process.
type AssignEmployeeCollectiveAgreementSubBusinessProcessType struct {
	BusinessSubProcessParameters          *BusinessSubProcessParametersType             `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	AssignEmployeeCollectiveAgreementData *AssignEmployeeCollectiveAgreementSubDataType `xml:"urn:com.workday/bsvc Assign_Employee_Collective_Agreement_Data,omitempty"`
}

// Wrapper element for Assign Employee Collective Agreement event used as a sub bp.
type AssignEmployeeCollectiveAgreementSubDataType struct {
	CollectiveAgreementReference *CollectiveAgreementObjectType            `xml:"urn:com.workday/bsvc Collective_Agreement_Reference"`
	FactorOptions                *CollectiveAgreementFactorOptionsDataType `xml:"urn:com.workday/bsvc Factor_Options,omitempty"`
}

// Worker, position, and matrix organization data
type AssignMatrixOrganizationBusinessProcessDataType struct {
	WorkerReference                *WorkerObjectType                   `xml:"urn:com.workday/bsvc Worker_Reference"`
	PositionMatrixOrganizationData *PositionMatrixOrganizationDataType `xml:"urn:com.workday/bsvc Position_Matrix_Organization_Data"`
}

// Wrapper for Assign Matrix Organizations event data which is usable for the standalone business process web service and when it is used as a sub business process.
type AssignMatrixOrganizationDataType struct {
	MatrixOrganizationReference []MatrixOrganizationObjectType `xml:"urn:com.workday/bsvc Matrix_Organization_Reference,omitempty"`
}

// Request for each worker, position, matrix org in Assign Matrix Organization
type AssignMatrixOrganizationRequestHVType struct {
	BusinessProcessParameters                   *BusinessProcessParametersType                   `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AssignMatrixOrganizationBusinessProcessData *AssignMatrixOrganizationBusinessProcessDataType `xml:"urn:com.workday/bsvc Assign_Matrix_Organization_Business_Process_Data"`
}

// Encapsulates the Assign Matrix Organization sub business process.
type AssignMatrixOrganizationSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	AssignMatrixOrganizationData *AssignMatrixOrganizationDataType `xml:"urn:com.workday/bsvc Assign_Matrix_Organization_Data,omitempty"`
}

// Wrapper element for the Assign Organization web service.
type AssignOrganizationEventDataType struct {
	PositionReference             *StaffingInterfaceObjectType               `xml:"urn:com.workday/bsvc Position_Reference"`
	WorkerReference               *WorkerObjectType                          `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	OrganizationReference         []OrganizationObjectType                   `xml:"urn:com.workday/bsvc Organization_Reference,omitempty"`
	FundReference                 []FundObjectType                           `xml:"urn:com.workday/bsvc Fund_Reference,omitempty"`
	GrantReference                []GrantObjectType                          `xml:"urn:com.workday/bsvc Grant_Reference,omitempty"`
	ProgramReference              []ProgramObjectType                        `xml:"urn:com.workday/bsvc Program_Reference,omitempty"`
	BusinessUnitReference         []BusinessUnitObjectType                   `xml:"urn:com.workday/bsvc Business_Unit_Reference,omitempty"`
	GiftReference                 []GiftObjectType                           `xml:"urn:com.workday/bsvc Gift_Reference,omitempty"`
	CheckPositionBudgetSubProcess *CheckPositionBudgetSubBusinessProcessType `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process,omitempty"`
	AsOfEffectiveDate             time.Time                                  `xml:"urn:com.workday/bsvc As_Of_Effective_Date,attr,omitempty"`
}

func (t *AssignOrganizationEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AssignOrganizationEventDataType
	var layout struct {
		*T
		AsOfEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc As_Of_Effective_Date,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AsOfEffectiveDate = (*xsdDate)(&layout.T.AsOfEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *AssignOrganizationEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AssignOrganizationEventDataType
	var overlay struct {
		*T
		AsOfEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc As_Of_Effective_Date,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AsOfEffectiveDate = (*xsdDate)(&overlay.T.AsOfEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// This operation will assign company, cost center, region, and custom organizations (that are configured for staffing usage) to a position using the Assign Organization business process event.
type AssignOrganizationRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType   `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AssignOrganizationData    *AssignOrganizationEventDataType `xml:"urn:com.workday/bsvc Assign_Organization_Data"`
	Version                   string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Event ID for the Assign Organization Event.
type AssignOrganizationResponseType struct {
	EventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// This operation will assign organization roles to one ore more workers using the Assign Organization Roles business process event.
type AssignOrganizationRolesRequestType struct {
	BusinessProcessParameters   *BusinessProcessParametersType        `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AssignOrganizationRolesData *AssignOrganizationRolesEventDataType `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Data"`
	Version                     string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Event ID for the Assign Organization Roles business process event.
type AssignOrganizationRolesResponseType struct {
	AssignRolesEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Assign_Roles_Event_Reference,omitempty"`
	Version                   string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper to hold the business process configuration and organization role assignment details. For fields that accept worker, worker's position as of specified effective date will be determined when event is submitted, not when it is completed. This means if worker is being assigned a new position and worker is specified, role will be assigned to their existing position and not their new position. Workday recommends you use 'Assign_Roles_Subprocess' instead to reduce the risk of your role assignments being inadvertently reversed due to another role assignment event being processed at the same time or later-dated role assignments.
type AssignOrganizationRolesSubBusinessProcessType struct {
	BusinessSubProcessParameters     *BusinessSubProcessParametersType     `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	AssignOrganizationRolesEventData *AssignOrganizationRolesEventDataType `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Event_Data,omitempty"`
}

// Wrapper element for Assign Pay Group business process. The Pay Group cannot be removed from a worker once it is set. It can only be changed to another pay group.
type AssignPayGroupDataType struct {
	PayGroupReference                                 *PayGroupObjectType         `xml:"urn:com.workday/bsvc Pay_Group_Reference,omitempty"`
	AdditionalPositionsforPayGroupAssignmentReference []PositionElementObjectType `xml:"urn:com.workday/bsvc Additional_Positions_for_Pay_Group_Assignment_Reference,omitempty"`
}

// Wrapper element for Assign Pay Group sub business process.
type AssignPayGroupSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	AssignPayGroupData           *AssignPayGroupDataType           `xml:"urn:com.workday/bsvc Assign_Pay_Group_Data,omitempty"`
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

// Wrapper Element for the Assign Roles web service
type AssignRolesEventDataType struct {
	EffectiveDate                                  *time.Time                          `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	EffectiveTimezoneReference                     *TimeZoneObjectType                 `xml:"urn:com.workday/bsvc Effective_Timezone_Reference,omitempty"`
	EventTargetAssigneeReference                   *RoleeSelectorObjectType            `xml:"urn:com.workday/bsvc Event_Target_Assignee_Reference"`
	RemoveAllRoleAssignmentsforEventTargetAssignee *bool                               `xml:"urn:com.workday/bsvc Remove_All_Role_Assignments_for_Event_Target_Assignee,omitempty"`
	AssignRolesRoleAssignmentData                  []AssignRolesRoleAssignmentDataType `xml:"urn:com.workday/bsvc Assign_Roles_Role_Assignment_Data,omitempty"`
}

func (t *AssignRolesEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AssignRolesEventDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *AssignRolesEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AssignRolesEventDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// This operation will assign organization roles to one ore more workers or positions.
type AssignRolesRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AssignRolesEventData      *AssignRolesEventDataType      `xml:"urn:com.workday/bsvc Assign_Roles_Event_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Event ID for the Assign Organization Roles business process event.
type AssignRolesResponseType struct {
	AssignRolesEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Assign_Roles_Event_Reference,omitempty"`
	Version                   string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Encapsulating Element for all Organization Role Assignment data.
type AssignRolesRoleAssignmentDataType struct {
	RoleAssignerReference                                   *RoleAssignerObjectType   `xml:"urn:com.workday/bsvc Role_Assigner_Reference"`
	AssignableRoleReference                                 *AssignableRoleObjectType `xml:"urn:com.workday/bsvc Assignable_Role_Reference"`
	RemoveExistingAssigneesforAssignableRoleonRoleAssigner  *bool                     `xml:"urn:com.workday/bsvc Remove_Existing_Assignees_for_Assignable_Role_on_Role_Assigner,omitempty"`
	UpdateLaterDatedAssignments                             *bool                     `xml:"urn:com.workday/bsvc Update_Later_Dated_Assignments,omitempty"`
	AssigneestoAddReference                                 []RoleeSelectorObjectType `xml:"urn:com.workday/bsvc Assignees_to_Add_Reference,omitempty"`
	AssigneestoRemoveReference                              []RoleeSelectorObjectType `xml:"urn:com.workday/bsvc Assignees_to_Remove_Reference,omitempty"`
	SupervisoryOrganizationSingleAssignmentManagerReference *RoleeSelectorObjectType  `xml:"urn:com.workday/bsvc Supervisory_Organization_Single_Assignment_Manager_Reference,omitempty"`
	RemoveSupervisoryOrganizationSingleAssignmentManager    *bool                     `xml:"urn:com.workday/bsvc Remove_Supervisory_Organization_Single_Assignment_Manager,omitempty"`
}

// Container for the Assign Roles and related Role Assignments data
type AssignRolesSubProcessDataType struct {
	RemoveAllRoleAssignmentsforEventTargetAssignee *bool                               `xml:"urn:com.workday/bsvc Remove_All_Role_Assignments_for_Event_Target_Assignee,omitempty"`
	AssignRolesRoleAssignmentData                  []AssignRolesRoleAssignmentDataType `xml:"urn:com.workday/bsvc Assign_Roles_Role_Assignment_Data,omitempty"`
}

// Container for Organization Assign Roles and options for this sub-business process. If a worker is specified, position to which role is assigned is determined when event is submitted, not completed. If worker has multiple jobs as of specified effective date, role will be assigned to international assignment (IA) position with earliest start date, otherwise to worker's primary position. Specify a position if this behavior for worker is not desired.
type AssignRolesSubProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	AssignRolesSubProcessData    *AssignRolesSubProcessDataType    `xml:"urn:com.workday/bsvc Assign_Roles_Sub_Process_Data,omitempty"`
}

// Assign Superior Organization Sub Process
type AssignSuperiorOrganizationDataType struct {
	ProposedSupervisoryOrganizationReference     *HierarchicAssignerObjectType       `xml:"urn:com.workday/bsvc Proposed_Supervisory_Organization_Reference"`
	SubordinateSupervisoryOrganizationsReference []SupervisoryOrganizationObjectType `xml:"urn:com.workday/bsvc Subordinate_Supervisory_Organizations_Reference"`
}

// Wrapper for  Assign Superior Organization
type AssignSuperiorOrganizationSubProcessType struct {
	BusinessSubProcessParameters   *BusinessSubProcessParametersType   `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	AssignSuperiorOrganizationData *AssignSuperiorOrganizationDataType `xml:"urn:com.workday/bsvc Assign_Superior_Organization_Data,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type BackgroundCheckStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BackgroundCheckStatusObjectType struct {
	ID         []BackgroundCheckStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the detailed allocation amounts for the beneficiaries of the elected coverage.
type BeneficiaryAllocationDataType struct {
	AllocationType    string              `xml:"urn:com.workday/bsvc Allocation_Type"`
	Amount            float64             `xml:"urn:com.workday/bsvc Amount"`
	AmountType        string              `xml:"urn:com.workday/bsvc Amount_Type"`
	CurrencyReference *CurrencyObjectType `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
}

// Contains the details about the beneficiary.
type BeneficiaryDataType struct {
	BeneficiaryID string                                               `xml:"urn:com.workday/bsvc Beneficiary_ID,omitempty"`
	Irrevocable   *bool                                                `xml:"urn:com.workday/bsvc Irrevocable,omitempty"`
	InactiveDate  *time.Time                                           `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
	CourtOrder    []QualifiedDomesticRelationsOrderReplacementDataType `xml:"urn:com.workday/bsvc Court_Order,omitempty"`
}

func (t *BeneficiaryDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BeneficiaryDataType
	var layout struct {
		*T
		InactiveDate *xsdDate `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.InactiveDate = (*xsdDate)(layout.T.InactiveDate)
	return e.EncodeElement(layout, start)
}
func (t *BeneficiaryDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BeneficiaryDataType
	var overlay struct {
		*T
		InactiveDate *xsdDate `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.InactiveDate = (*xsdDate)(overlay.T.InactiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains the beneficiary that the coverage allocation is for and the details about the allocation.
type BeneficiaryDesignationDataType struct {
	BeneficiaryReference      *BeneficiaryObjectType         `xml:"urn:com.workday/bsvc Beneficiary_Reference"`
	OriginalCoverageBeginDate time.Time                      `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date"`
	CoverageEndDate           *time.Time                     `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	BeneficiaryAllocationData *BeneficiaryAllocationDataType `xml:"urn:com.workday/bsvc Beneficiary_Allocation_Data"`
}

func (t *BeneficiaryDesignationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BeneficiaryDesignationDataType
	var layout struct {
		*T
		OriginalCoverageBeginDate *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date"`
		CoverageEndDate           *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.OriginalCoverageBeginDate = (*xsdDate)(&layout.T.OriginalCoverageBeginDate)
	layout.CoverageEndDate = (*xsdDate)(layout.T.CoverageEndDate)
	return e.EncodeElement(layout, start)
}
func (t *BeneficiaryDesignationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BeneficiaryDesignationDataType
	var overlay struct {
		*T
		OriginalCoverageBeginDate *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date"`
		CoverageEndDate           *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.OriginalCoverageBeginDate = (*xsdDate)(&overlay.T.OriginalCoverageBeginDate)
	overlay.CoverageEndDate = (*xsdDate)(overlay.T.CoverageEndDate)
	return d.DecodeElement(&overlay, &start)
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

// Encapsulating element containing all Beneficiary data.
type BeneficiaryType struct {
	BeneficiaryReference *BeneficiaryObjectType `xml:"urn:com.workday/bsvc Beneficiary_Reference"`
	BeneficiaryData      *BeneficiaryDataType   `xml:"urn:com.workday/bsvc Beneficiary_Data"`
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

// Contains a unique identifier for an instance of an object.
type BenefitLifeEventTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitLifeEventTypeObjectType struct {
	ID         []BenefitLifeEventTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains the summary information about the benefit provider of the benefit plan.
type BenefitPlanSummaryDataType struct {
	BenefitPlanReference              *BenefitPlanObjectType              `xml:"urn:com.workday/bsvc Benefit_Plan_Reference,omitempty"`
	BenefitPlanName                   string                              `xml:"urn:com.workday/bsvc Benefit_Plan_Name,omitempty"`
	GroupIdentifier                   string                              `xml:"urn:com.workday/bsvc Group_Identifier,omitempty"`
	BenefitCoverageTypeReference      *BenefitCoverageTypeObjectType      `xml:"urn:com.workday/bsvc Benefit_Coverage_Type_Reference,omitempty"`
	CurrencyReference                 *CurrencyObjectType                 `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference                *FrequencyObjectType                `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	HealthCareClassificationReference *HealthCareClassificationObjectType `xml:"urn:com.workday/bsvc Health_Care_Classification_Reference,omitempty"`
	InsuranceCoverageTargetReference  *InsuranceCoverageTargetObjectType  `xml:"urn:com.workday/bsvc Insurance_Coverage_Target_Reference,omitempty"`
	BenefitProviderSummaryData        *BenefitProviderSummaryDataType     `xml:"urn:com.workday/bsvc Benefit_Provider_Summary_Data,omitempty"`
	ExternalIntegrationIDData         *ExternalIntegrationIDDataType      `xml:"urn:com.workday/bsvc External_Integration_ID_Data,omitempty"`
	PayComponentReference             []PayComponentObjectType            `xml:"urn:com.workday/bsvc Pay_Component_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BenefitProviderIdentifierTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitProviderIdentifierTypeObjectType struct {
	ID         []BenefitProviderIdentifierTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BenefitProviderObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitProviderObjectType struct {
	ID         []BenefitProviderObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the summary information about the benefit provider of the benefit plan.
type BenefitProviderSummaryDataType struct {
	BenefitProviderReference               *BenefitProviderObjectType               `xml:"urn:com.workday/bsvc Benefit_Provider_Reference,omitempty"`
	BenefitProviderName                    string                                   `xml:"urn:com.workday/bsvc Benefit_Provider_Name,omitempty"`
	BenefitProviderIdentifierTypeReference *BenefitProviderIdentifierTypeObjectType `xml:"urn:com.workday/bsvc Benefit_Provider_Identifier_Type_Reference,omitempty"`
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

// Contains COBRA Eligibility detail for a participant.  If this element is not populated then the participant is assumed to not be COBRA eligible.
type COBRAEligibilityDataType struct {
	COBRAEligibilityReasonReference *COBRAEligibilityReasonObjectType `xml:"urn:com.workday/bsvc COBRA_Eligibility_Reason_Reference,omitempty"`
	EligibleDate                    *time.Time                        `xml:"urn:com.workday/bsvc Eligible_Date,omitempty"`
	LossofCoverageDate              *time.Time                        `xml:"urn:com.workday/bsvc Loss_of_Coverage_Date,omitempty"`
	CoverageEndDate                 *time.Time                        `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	BenefitPlanReference            *BenefitPlanObjectType            `xml:"urn:com.workday/bsvc Benefit_Plan_Reference,omitempty"`
}

func (t *COBRAEligibilityDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T COBRAEligibilityDataType
	var layout struct {
		*T
		EligibleDate       *xsdDate `xml:"urn:com.workday/bsvc Eligible_Date,omitempty"`
		LossofCoverageDate *xsdDate `xml:"urn:com.workday/bsvc Loss_of_Coverage_Date,omitempty"`
		CoverageEndDate    *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EligibleDate = (*xsdDate)(layout.T.EligibleDate)
	layout.LossofCoverageDate = (*xsdDate)(layout.T.LossofCoverageDate)
	layout.CoverageEndDate = (*xsdDate)(layout.T.CoverageEndDate)
	return e.EncodeElement(layout, start)
}
func (t *COBRAEligibilityDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T COBRAEligibilityDataType
	var overlay struct {
		*T
		EligibleDate       *xsdDate `xml:"urn:com.workday/bsvc Eligible_Date,omitempty"`
		LossofCoverageDate *xsdDate `xml:"urn:com.workday/bsvc Loss_of_Coverage_Date,omitempty"`
		CoverageEndDate    *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EligibleDate = (*xsdDate)(overlay.T.EligibleDate)
	overlay.LossofCoverageDate = (*xsdDate)(overlay.T.LossofCoverageDate)
	overlay.CoverageEndDate = (*xsdDate)(overlay.T.CoverageEndDate)
	return d.DecodeElement(&overlay, &start)
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
type CalculatedPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CalculatedPlanObjectType struct {
	ID         []CalculatedPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CareerPreferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CareerPreferenceObjectType struct {
	ID         []CareerPreferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Employees Career Interest Data.
type CareerPreferencesDataType struct {
	CareerInterestReference []CareerPreferenceObjectType `xml:"urn:com.workday/bsvc Career_Interest_Reference,omitempty"`
	CareerInterests         string                       `xml:"urn:com.workday/bsvc Career_Interests,omitempty"`
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

// Wrapper element for the emergency contact
type ChangeEmergencyContactDataType struct {
	EmergencyContactID                      string                                       `xml:"urn:com.workday/bsvc Emergency_Contact_ID,omitempty"`
	Primary                                 *bool                                        `xml:"urn:com.workday/bsvc Primary,omitempty"`
	Priority                                float64                                      `xml:"urn:com.workday/bsvc Priority"`
	RelatedPersonRelationshipReference      []RelatedPersonRelationshipObjectType        `xml:"urn:com.workday/bsvc Related_Person_Relationship_Reference"`
	LanguageReference                       []LanguageObjectType                         `xml:"urn:com.workday/bsvc Language_Reference,omitempty"`
	EmergencyContactPersonalInformationData *EmergencyContactPersonalInformationDataType `xml:"urn:com.workday/bsvc Emergency_Contact_Personal_Information_Data,omitempty"`
}

// Wrapper element to hold reference and data for an emergency contact.
type ChangeEmergencyContactsDataType struct {
	EmergencyContactReference *EmergencyContactObjectType     `xml:"urn:com.workday/bsvc Emergency_Contact_Reference,omitempty"`
	Delete                    *bool                           `xml:"urn:com.workday/bsvc Delete,omitempty"`
	EmergencyContactData      *ChangeEmergencyContactDataType `xml:"urn:com.workday/bsvc Emergency_Contact_Data,omitempty"`
}

// Wrapper for the Change Emergency Contacts sub business process that is part of the Hire and Edit Position business processes.
type ChangeEmergencyContactsSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType   `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	ChangeEmergencyContactsData  *ChangeEmergencyContactsSubDataType `xml:"urn:com.workday/bsvc Change_Emergency_Contacts_Data,omitempty"`
}

// Container for data used in Change Emergency Contact sub business process
type ChangeEmergencyContactsSubDataType struct {
	ReplaceAll                     *bool                             `xml:"urn:com.workday/bsvc Replace_All,omitempty"`
	EmergencyContactsReferenceData []ChangeEmergencyContactsDataType `xml:"urn:com.workday/bsvc Emergency_Contacts_Reference_Data,omitempty"`
}

// Allows retrieving government IDs based on request criteria. Request Criteria is only used when government ID reference is not provided
type ChangeGovernmentIDRequestCriteriaType struct {
	PersonTypeCriteriaData *PersonTypeCriteriaType `xml:"urn:com.workday/bsvc Person_Type_Criteria_Data,omitempty"`
}

// Wrapper element for the Government Identifier data.
type ChangeGovernmentIDsBusinessProcessDataType struct {
	PersonReference              *RoleObjectType                   `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	UniversalIDReference         *UniversalIdentifierObjectType    `xml:"urn:com.workday/bsvc Universal_ID_Reference,omitempty"`
	GovernmentIdentificationdata *GovernmentIdentificationDataType `xml:"urn:com.workday/bsvc Government_Identification_data"`
}

// Element containing the worker data for the Get Change Government IDs response.
type ChangeGovernmentIDsResponseDataType struct {
	ChangeGovernmentIDs []ChangeGovernmentIDsResponseWrapperType `xml:"urn:com.workday/bsvc Change_Government_IDs,omitempty"`
}

// Wrapper element for the Change Government IDs Business Process Data element. This is the element that contains the data to load.
type ChangeGovernmentIDsResponseWrapperType struct {
	PersonReference         *RoleObjectType                              `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	ChangeGovernmentIDsData []ChangeGovernmentIDsBusinessProcessDataType `xml:"urn:com.workday/bsvc Change_Government_IDs_Data,omitempty"`
}

// Contains the Change Job web service and its sub business processes.
type ChangeJobDataType struct {
	WorkerReference                                 *WorkerObjectType                                        `xml:"urn:com.workday/bsvc Worker_Reference"`
	PositionReference                               *PositionElementObjectType                               `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	EffectiveDate                                   time.Time                                                `xml:"urn:com.workday/bsvc Effective_Date"`
	ChangeJobDetailData                             *ChangeJobDetailDataType                                 `xml:"urn:com.workday/bsvc Change_Job_Detail_Data"`
	ProposeCompensationSubProcess                   *ProposeCompensationForPositionSubBusinessProcessType    `xml:"urn:com.workday/bsvc Propose_Compensation_Sub_Process,omitempty"`
	RequestOneTimePaymentSubProcess                 *RequestOneTimePaymentSubBusinessProcessType             `xml:"urn:com.workday/bsvc Request_One_Time_Payment_Sub_Process,omitempty"`
	RequestStockGrantSubProcess                     *RequestStockSubBusinessProcessType                      `xml:"urn:com.workday/bsvc Request_Stock_Grant_Sub_Process,omitempty"`
	ChangeOrganizationSubProcess                    *EditAssignPositionOrganizationSubBusinessProcessType    `xml:"urn:com.workday/bsvc Change_Organization_Sub_Process,omitempty"`
	AssignPayGroupSubProcess                        *AssignPayGroupSubBusinessProcessType                    `xml:"urn:com.workday/bsvc Assign_Pay_Group_Sub_Process,omitempty"`
	ReviewPayrollInterfaceSubProcess                *ReviewPayrollInterfaceDataSubBusinessProcessType        `xml:"urn:com.workday/bsvc Review_Payroll_Interface_Sub_Process,omitempty"`
	AssignMatrixOrganizationSubProcess              *AssignMatrixOrganizationSubBusinessProcessType          `xml:"urn:com.workday/bsvc Assign_Matrix_Organization_Sub_Process,omitempty"`
	ChangePersonalInformationSubProcess             *ChangePersonalInformationSubBusinessProcessType         `xml:"urn:com.workday/bsvc Change_Personal_Information_Sub_Process,omitempty"`
	MaintainEmployeeContractsSubProcess             *MaintainEmployeeContractsSubBusinessProcessType         `xml:"urn:com.workday/bsvc Maintain_Employee_Contracts_Sub_Process,omitempty"`
	AssignOrganizationRolesSubProcess               *AssignOrganizationRolesSubBusinessProcessType           `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Sub_Process,omitempty"`
	AssignRolesSubProcess                           *AssignRolesSubProcessType                               `xml:"urn:com.workday/bsvc Assign_Roles_Sub_Process,omitempty"`
	AssignSuperiorOrganizationSubProcess            *AssignSuperiorOrganizationSubProcessType                `xml:"urn:com.workday/bsvc Assign_Superior_Organization_Sub_Process,omitempty"`
	CreateJobRequisitionSubProcess                  *CreateJobRequisitionSubProcessType                      `xml:"urn:com.workday/bsvc Create_Job_Requisition_Sub_Process,omitempty"`
	CheckPositionBudgetSubProcess                   *CheckPositionBudgetSubBusinessProcessType               `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process,omitempty"`
	SwitchPrimaryJobSubProcess                      *SwitchPrimaryJobSubProcessType                          `xml:"urn:com.workday/bsvc Switch_Primary_Job_Sub_Process,omitempty"`
	UpdateAcademicAppointmentSubProcess             *UpdateAcademicAppointmentSubProcessType                 `xml:"urn:com.workday/bsvc Update_Academic_Appointment_Sub_Process,omitempty"`
	AssignCostingAllocationSubProcess               *AssignCostingAllocationSubBusinessProcessType           `xml:"urn:com.workday/bsvc Assign_Costing_Allocation_Sub_Process,omitempty"`
	AssignEmployeeCollectiveAgreementSubProcess     *AssignEmployeeCollectiveAgreementSubBusinessProcessType `xml:"urn:com.workday/bsvc Assign_Employee_Collective_Agreement_Sub_Process,omitempty"`
	CreateSubordinateSubProcess                     *CreateSubordinateSubBusinessProcessType                 `xml:"urn:com.workday/bsvc Create_Subordinate_Sub_Process,omitempty"`
	StudentEmploymentEligibilitySubProcess          *StudentEmploymentEligibilitySubBusinessProcessType      `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Sub_Process,omitempty"`
	OnboardingSetupSubProcess                       *OnboardingSetupSubBusinessProcessType                   `xml:"urn:com.workday/bsvc Onboarding_Setup_Sub_Process,omitempty"`
	ManageUnionMembershipSubProcess                 *ManageUnionMembershipSubBusinessProcessType             `xml:"urn:com.workday/bsvc Manage_Union_Membership_Sub_Process,omitempty"`
	ManageEmployeeProbationPeriodSubBusinessProcess *ManageProbationPeriodSubBusinessProcessType             `xml:"urn:com.workday/bsvc Manage_Employee_Probation_Period_Sub_Business_Process,omitempty"`
	EditNoticePeriodsSubProcess                     *EditNoticePeriodsSubBusinessProcessType                 `xml:"urn:com.workday/bsvc Edit_Notice_Periods_Sub_Process,omitempty"`
}

func (t *ChangeJobDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ChangeJobDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *ChangeJobDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ChangeJobDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains the Change Job web service and its sub business processes.
type ChangeJobDataforMassPositionSwapType struct {
	WorkerReference                      *WorkerObjectType                                     `xml:"urn:com.workday/bsvc Worker_Reference"`
	CurrentPositionReference             *PositionElementObjectType                            `xml:"urn:com.workday/bsvc Current_Position_Reference"`
	ProposedPositionReference            *PositionElementObjectType                            `xml:"urn:com.workday/bsvc Proposed_Position_Reference"`
	ProposeCompensationSubProcess        *ProposeCompensationForPositionSubBusinessProcessType `xml:"urn:com.workday/bsvc Propose_Compensation_Sub_Process,omitempty"`
	RequestOneTimePaymentSubProcess      *RequestOneTimePaymentSubBusinessProcessType          `xml:"urn:com.workday/bsvc Request_One_Time_Payment_Sub_Process,omitempty"`
	RequestStockGrantSubProcess          *RequestStockSubBusinessProcessType                   `xml:"urn:com.workday/bsvc Request_Stock_Grant_Sub_Process,omitempty"`
	AssignPayGroupSubProcess             *AssignPayGroupSubBusinessProcessType                 `xml:"urn:com.workday/bsvc Assign_Pay_Group_Sub_Process,omitempty"`
	AssignMatrixOrganizationSubProcess   *AssignMatrixOrganizationSubBusinessProcessType       `xml:"urn:com.workday/bsvc Assign_Matrix_Organization_Sub_Process,omitempty"`
	ChangePersonalInformationSubProcess  *ChangePersonalInformationSubBusinessProcessType      `xml:"urn:com.workday/bsvc Change_Personal_Information_Sub_Process,omitempty"`
	MaintainEmployeeContractsSubProcess  *MaintainEmployeeContractsSubBusinessProcessType      `xml:"urn:com.workday/bsvc Maintain_Employee_Contracts_Sub_Process,omitempty"`
	AssignOrganizationRolesSubProcess    *AssignOrganizationRolesSubBusinessProcessType        `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Sub_Process,omitempty"`
	AssignRolesSubProcess                *AssignRolesSubProcessType                            `xml:"urn:com.workday/bsvc Assign_Roles_Sub_Process,omitempty"`
	AssignSuperiorOrganizationSubProcess *AssignSuperiorOrganizationSubProcessType             `xml:"urn:com.workday/bsvc Assign_Superior_Organization_Sub_Process,omitempty"`
	AssignCostingAllocationSubProcess    *AssignCostingAllocationSubBusinessProcessType        `xml:"urn:com.workday/bsvc Assign_Costing_Allocation_Sub_Process,omitempty"`
	CreateSubordinateSubProcess          *CreateSubordinateSubBusinessProcessType              `xml:"urn:com.workday/bsvc Create_Subordinate_Sub_Process,omitempty"`
	OnboardingSetupSubProcess            *OnboardingSetupSubBusinessProcessType                `xml:"urn:com.workday/bsvc Onboarding_Setup_Sub_Process,omitempty"`
	ManageUnionMembershipSubProcess      *ManageUnionMembershipSubBusinessProcessType          `xml:"urn:com.workday/bsvc Manage_Union_Membership_Sub_Process,omitempty"`
}

// Contains the details of the change job business process including the reason.
type ChangeJobDetailDataType struct {
	ReasonReference                      *ChangeJobSubcategoryObjectType          `xml:"urn:com.workday/bsvc Reason_Reference"`
	SupervisoryOrganizationReference     *SupervisoryOrganizationObjectType       `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference,omitempty"`
	HeadcountOptionReference             *HeadcountOptionsObjectType              `xml:"urn:com.workday/bsvc Headcount_Option_Reference,omitempty"`
	JobOverlapAllowed                    *bool                                    `xml:"urn:com.workday/bsvc Job_Overlap_Allowed,omitempty"`
	ProposedPositionReference            *PositionRestrictionsObjectType          `xml:"urn:com.workday/bsvc Proposed_Position_Reference,omitempty"`
	ProposedJobRequisitionReference      *JobRequisitionObjectType                `xml:"urn:com.workday/bsvc Proposed_Job_Requisition_Reference,omitempty"`
	CreatePosition                       *bool                                    `xml:"urn:com.workday/bsvc Create_Position,omitempty"`
	EmployeeTypeReference                *EmployeeTypeObjectType                  `xml:"urn:com.workday/bsvc Employee_Type_Reference,omitempty"`
	ContingentWorkerTypeReference        *ContingentWorkerTypeObjectType          `xml:"urn:com.workday/bsvc Contingent_Worker_Type_Reference,omitempty"`
	JobDetailsData                       *PositionDetailsSubDataType              `xml:"urn:com.workday/bsvc Job_Details_Data,omitempty"`
	InternationalAssignmentTypeReference *InternationalAssignmentTypeObjectType   `xml:"urn:com.workday/bsvc International_Assignment_Type_Reference,omitempty"`
	ExpectedAssignmentEndDate            *time.Time                               `xml:"urn:com.workday/bsvc Expected_Assignment_End_Date,omitempty"`
	EndEmploymentDate                    *time.Time                               `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
	ContractEndDate                      *time.Time                               `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
	FirstDayofWork                       *time.Time                               `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
	NotifyWorkerBy                       *time.Time                               `xml:"urn:com.workday/bsvc Notify_Worker_By,omitempty"`
	NewPositionID                        string                                   `xml:"urn:com.workday/bsvc New_Position_ID,omitempty"`
	ContractDetailsData                  *ContractDetailsSubDataType              `xml:"urn:com.workday/bsvc Contract_Details_Data,omitempty"`
	WorkerDocumentData                   []WorkerDocumentforStaffingEventDataType `xml:"urn:com.workday/bsvc Worker_Document_Data,omitempty"`
}

func (t *ChangeJobDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ChangeJobDetailDataType
	var layout struct {
		*T
		ExpectedAssignmentEndDate *xsdDate `xml:"urn:com.workday/bsvc Expected_Assignment_End_Date,omitempty"`
		EndEmploymentDate         *xsdDate `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
		ContractEndDate           *xsdDate `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
		FirstDayofWork            *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
		NotifyWorkerBy            *xsdDate `xml:"urn:com.workday/bsvc Notify_Worker_By,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ExpectedAssignmentEndDate = (*xsdDate)(layout.T.ExpectedAssignmentEndDate)
	layout.EndEmploymentDate = (*xsdDate)(layout.T.EndEmploymentDate)
	layout.ContractEndDate = (*xsdDate)(layout.T.ContractEndDate)
	layout.FirstDayofWork = (*xsdDate)(layout.T.FirstDayofWork)
	layout.NotifyWorkerBy = (*xsdDate)(layout.T.NotifyWorkerBy)
	return e.EncodeElement(layout, start)
}
func (t *ChangeJobDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ChangeJobDetailDataType
	var overlay struct {
		*T
		ExpectedAssignmentEndDate *xsdDate `xml:"urn:com.workday/bsvc Expected_Assignment_End_Date,omitempty"`
		EndEmploymentDate         *xsdDate `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
		ContractEndDate           *xsdDate `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
		FirstDayofWork            *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
		NotifyWorkerBy            *xsdDate `xml:"urn:com.workday/bsvc Notify_Worker_By,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ExpectedAssignmentEndDate = (*xsdDate)(overlay.T.ExpectedAssignmentEndDate)
	overlay.EndEmploymentDate = (*xsdDate)(overlay.T.EndEmploymentDate)
	overlay.ContractEndDate = (*xsdDate)(overlay.T.ContractEndDate)
	overlay.FirstDayofWork = (*xsdDate)(overlay.T.FirstDayofWork)
	overlay.NotifyWorkerBy = (*xsdDate)(overlay.T.NotifyWorkerBy)
	return d.DecodeElement(&overlay, &start)
}

// This web service operation is designed to perform a job change on an employee or contingent worker using the Change Job business process. The types of changes include transfer, promotion, demotion, lateral moves and any other change of data on the job.
type ChangeJobRequestHVType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	ChangeJobData             *ChangeJobDataType             `xml:"urn:com.workday/bsvc Change_Job_Data"`
}

// Business Process for doing a Change Job on a worker.
type ChangeJobRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	ChangeJobData             *ChangeJobDataType             `xml:"urn:com.workday/bsvc Change_Job_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Event ID of the Change Job and any exceptions.
type ChangeJobResponseType struct {
	EventReference         *UniqueIdentifierObjectType                    `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	ExceptionsResponseData []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ChangeJobSubcategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ChangeJobSubcategoryObjectType struct {
	ID         []ChangeJobSubcategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for the Legal Name Change business process web service.
type ChangeLegalNameBusinessProcessDataType struct {
	PersonReference      *RoleObjectType                `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	UniversalIDReference *UniversalIdentifierObjectType `xml:"urn:com.workday/bsvc Universal_ID_Reference,omitempty"`
	EffectiveDate        *time.Time                     `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	NameData             *PersonNameDetailDataType      `xml:"urn:com.workday/bsvc Name_Data"`
}

func (t *ChangeLegalNameBusinessProcessDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ChangeLegalNameBusinessProcessDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *ChangeLegalNameBusinessProcessDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ChangeLegalNameBusinessProcessDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Allows retrieving Legal Name based on request criteria.
type ChangeLegalNameRequestCriteriaType struct {
	PersonTypeCriteriaData *PersonTypeCriteriaType `xml:"urn:com.workday/bsvc Person_Type_Criteria_Data,omitempty"`
}

// Element containing the worker data for the Get Change Legal Name response
type ChangeLegalNameResponseDataType struct {
	ChangeLegalName []ChangeLegalNameResponseWrapperType `xml:"urn:com.workday/bsvc Change_Legal_Name,omitempty"`
}

// Wrapper element for the Change Legal Name Business Process Data element. This is the element that contains data to load.
type ChangeLegalNameResponseWrapperType struct {
	PersonReference     *RoleObjectType                         `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	ChangeLegalNameData *ChangeLegalNameBusinessProcessDataType `xml:"urn:com.workday/bsvc Change_Legal_Name_Data,omitempty"`
}

// Allows retrieving License IDs based on request criteria.
type ChangeLicenseIDRequestCriteriaType struct {
	PersonTypeCriteriaData *PersonTypeCriteriaType `xml:"urn:com.workday/bsvc Person_Type_Criteria_Data,omitempty"`
}

// Wrapper element for the License Identifier data.
type ChangeLicensesBusinessProcessDataType struct {
	PersonReference           *RoleObjectType                `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	UniversalIDReference      *UniversalIdentifierObjectType `xml:"urn:com.workday/bsvc Universal_ID_Reference,omitempty"`
	LicenseIdentificationData *LicenseIdentificationDataType `xml:"urn:com.workday/bsvc License_Identification_Data"`
}

// Element containing the worker data for the Get Change Licenses response.
type ChangeLicensesResponseDataType struct {
	ChangeLicenses []ChangeLicensesResponseWrapperType `xml:"urn:com.workday/bsvc Change_Licenses,omitempty"`
}

// Wrapper element for the Change Licenses Business Process Data element. This is the element that contains the data to load.
type ChangeLicensesResponseWrapperType struct {
	PersonReference    *RoleObjectType                         `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	ChangeLicensesData []ChangeLicensesBusinessProcessDataType `xml:"urn:com.workday/bsvc Change_Licenses_Data,omitempty"`
}

// Wrapper element for the Change Organization Assignments web service.
type ChangeOrganizationAssignmentsDataType struct {
	PositionReference                   *StaffingInterfaceObjectType               `xml:"urn:com.workday/bsvc Position_Reference"`
	WorkerReference                     *WorkerObjectType                          `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	PositionOrganizationAssignmentsData *AssignPositionOrganizationEventDataType   `xml:"urn:com.workday/bsvc Position_Organization_Assignments_Data"`
	CheckPositionBudgetSubProcess       *CheckPositionBudgetSubBusinessProcessType `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process,omitempty"`
	EffectiveDate                       time.Time                                  `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
}

func (t *ChangeOrganizationAssignmentsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ChangeOrganizationAssignmentsDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *ChangeOrganizationAssignmentsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ChangeOrganizationAssignmentsDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// This operation will assign company, cost center, region, and custom organizations (that are configured for staffing usage) to a position using the Change Organization Assignments business process.
type ChangeOrganizationAssignmentsRequestType struct {
	BusinessProcessParameters         *BusinessProcessParametersType         `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	ChangeOrganizationAssignmentsData *ChangeOrganizationAssignmentsDataType `xml:"urn:com.workday/bsvc Change_Organization_Assignments_Data"`
	Version                           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Event ID for the Change Organization Assignments Event.
type ChangeOrganizationAssignmentsResponseType struct {
	ChangeOrganizationAssignmentsEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Change_Organization_Assignments_Event_Reference,omitempty"`
	Version                                     string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Allows retrieving Passport and Visa IDs based on request criteria.
type ChangeOtherIDRequestCriteriaType struct {
	PersonTypeCriteriaData *PersonTypeCriteriaType `xml:"urn:com.workday/bsvc Person_Type_Criteria_Data,omitempty"`
}

// Wrapper element for the Other Identifier Data.
type ChangeOtherIDsBusinessProcessDataType struct {
	PersonReference          *RoleObjectType                `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	UniversalIDReference     *UniversalIdentifierObjectType `xml:"urn:com.workday/bsvc Universal_ID_Reference,omitempty"`
	CustomIdentificationData *CustomIdentificationDataType  `xml:"urn:com.workday/bsvc Custom_Identification_Data"`
}

// element containing the worker data for the Get Change Other IDs response.
type ChangeOtherIDsResponseDataType struct {
	ChangeOtherIDs []ChangeOtherIDsResponseWrapperType `xml:"urn:com.workday/bsvc Change_Other_IDs,omitempty"`
}

// Wrapper element for the Change Other IDs Business Process Data element. This is the element that contains the data to load.
type ChangeOtherIDsResponseWrapperType struct {
	PersonReference    *RoleObjectType                         `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	ChangeOtherIDsData []ChangeOtherIDsBusinessProcessDataType `xml:"urn:com.workday/bsvc Change_Other_IDs_Data,omitempty"`
}

// Allows retrieving Passport and Visa IDs based on request criteria.
type ChangePassportandVisaRequestCriteriaType struct {
	PersonTypeCriteriaData *PersonTypeCriteriaType `xml:"urn:com.workday/bsvc Person_Type_Criteria_Data,omitempty"`
}

// Wrapper element for the Passport and Visa Identifier data.
type ChangePassportsandVisasBusinessProcessDataType struct {
	PersonReference                     *RoleObjectType                          `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	UniversalIDReference                *UniversalIdentifierObjectType           `xml:"urn:com.workday/bsvc Universal_ID_Reference,omitempty"`
	PassportsandVisasIdentificationData *PassportsandVisasIdentificationDataType `xml:"urn:com.workday/bsvc Passports_and_Visas_Identification_Data"`
}

// Element containing the worker data for the Get Change Passports and Visas response.
type ChangePassportsandVisasResponseDataType struct {
	ChangePassportsandVisas []ChangePassportsandVisasResponseWrapperType `xml:"urn:com.workday/bsvc Change_Passports_and_Visas,omitempty"`
}

// Wrapper element for the Change Passports and Visas Business Process Data element. This is the element that contains the data to load.
type ChangePassportsandVisasResponseWrapperType struct {
	PersonReference             *RoleObjectType                                  `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	ChangePassportsandVisasData []ChangePassportsandVisasBusinessProcessDataType `xml:"urn:com.workday/bsvc Change_Passports_and_Visas_Data,omitempty"`
}

// Container for the data changed in the Change Personal Information business process.
type ChangePersonalInformationDataType struct {
	DateofBirth                     *time.Time                        `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	BirthCountryReference           *CountryObjectType                `xml:"urn:com.workday/bsvc Birth_Country_Reference,omitempty"`
	BirthRegionReference            *CountryRegionObjectType          `xml:"urn:com.workday/bsvc Birth_Region_Reference,omitempty"`
	GenderReference                 *GenderObjectType                 `xml:"urn:com.workday/bsvc Gender_Reference,omitempty"`
	DisabilityInformationData       *DisabilityInformationDataType    `xml:"urn:com.workday/bsvc Disability_Information_Data,omitempty"`
	MaritalStatusReference          *MaritalStatusObjectType          `xml:"urn:com.workday/bsvc Marital_Status_Reference,omitempty"`
	CitizenshipReference            []CitizenshipStatusObjectType     `xml:"urn:com.workday/bsvc Citizenship_Reference,omitempty"`
	PrimaryNationalityReference     *CountryObjectType                `xml:"urn:com.workday/bsvc Primary_Nationality_Reference,omitempty"`
	AdditionalNationalityReference  []CountryObjectType               `xml:"urn:com.workday/bsvc Additional_Nationality_Reference,omitempty"`
	EthnicityReference              []EthnicityObjectType             `xml:"urn:com.workday/bsvc Ethnicity_Reference,omitempty"`
	HispanicorLatino                *bool                             `xml:"urn:com.workday/bsvc Hispanic_or_Latino,omitempty"`
	VisualSurveyEthnicityReference  []EthnicityObjectType             `xml:"urn:com.workday/bsvc Visual_Survey_Ethnicity_Reference,omitempty"`
	HispanicorLatinoforVisualSurvey *bool                             `xml:"urn:com.workday/bsvc Hispanic_or_Latino_for_Visual_Survey,omitempty"`
	ReligionReference               *ReligionObjectType               `xml:"urn:com.workday/bsvc Religion_Reference,omitempty"`
	HukouRegionReference            *CountryRegionObjectType          `xml:"urn:com.workday/bsvc Hukou_Region_Reference,omitempty"`
	HukouSubregionReference         *CountrySubregionObjectType       `xml:"urn:com.workday/bsvc Hukou_Subregion_Reference,omitempty"`
	HukouLocality                   string                            `xml:"urn:com.workday/bsvc Hukou_Locality,omitempty"`
	HukouPostalCode                 string                            `xml:"urn:com.workday/bsvc Hukou_Postal_Code,omitempty"`
	HukouTypeReference              *HukouTypeObjectType              `xml:"urn:com.workday/bsvc Hukou_Type_Reference,omitempty"`
	NativeRegionReference           *CountryRegionObjectType          `xml:"urn:com.workday/bsvc Native_Region_Reference,omitempty"`
	PersonnelFileAgency             string                            `xml:"urn:com.workday/bsvc Personnel_File_Agency,omitempty"`
	MilitaryInformationData         *MilitaryInformationDataType      `xml:"urn:com.workday/bsvc Military_Information_Data,omitempty"`
	PoliticalAffiliationReference   *PoliticalAffiliationObjectType   `xml:"urn:com.workday/bsvc Political_Affiliation_Reference,omitempty"`
	DateofDeath                     *time.Time                        `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
	CityofBirth                     string                            `xml:"urn:com.workday/bsvc City_of_Birth,omitempty"`
	CityofBirthReference            *CountryCityObjectType            `xml:"urn:com.workday/bsvc City_of_Birth_Reference,omitempty"`
	MaritalStatusDate               *time.Time                        `xml:"urn:com.workday/bsvc Marital_Status_Date,omitempty"`
	LastMedicalExamDate             *time.Time                        `xml:"urn:com.workday/bsvc Last_Medical_Exam_Date,omitempty"`
	LastMedicalExamValidTo          *time.Time                        `xml:"urn:com.workday/bsvc Last_Medical_Exam_Valid_To,omitempty"`
	MedicalExamNotes                string                            `xml:"urn:com.workday/bsvc Medical_Exam_Notes,omitempty"`
	BloodTypeReference              *BloodTypeObjectType              `xml:"urn:com.workday/bsvc Blood_Type_Reference,omitempty"`
	UsesTobacco                     *bool                             `xml:"urn:com.workday/bsvc Uses_Tobacco,omitempty"`
	SocialBenefitsLocalityReference *SocialBenefitsLocalityObjectType `xml:"urn:com.workday/bsvc Social_Benefits_Locality_Reference,omitempty"`
	LGBTIdentificationReference     []LGBTIdentificationObjectType    `xml:"urn:com.workday/bsvc LGBT_Identification_Reference,omitempty"`
	SexualOrientationReference      *SexualOrientationObjectType      `xml:"urn:com.workday/bsvc Sexual_Orientation_Reference,omitempty"`
	GenderIdentityReference         *GenderIdentityObjectType         `xml:"urn:com.workday/bsvc Gender_Identity_Reference,omitempty"`
	PronounReference                *PronounObjectType                `xml:"urn:com.workday/bsvc Pronoun_Reference,omitempty"`
	RelativeNameInformationData     *RelativeNameInformationDataType  `xml:"urn:com.workday/bsvc Relative_Name_Information_Data,omitempty"`
}

func (t *ChangePersonalInformationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ChangePersonalInformationDataType
	var layout struct {
		*T
		DateofBirth            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
		DateofDeath            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
		MaritalStatusDate      *xsdDate `xml:"urn:com.workday/bsvc Marital_Status_Date,omitempty"`
		LastMedicalExamDate    *xsdDate `xml:"urn:com.workday/bsvc Last_Medical_Exam_Date,omitempty"`
		LastMedicalExamValidTo *xsdDate `xml:"urn:com.workday/bsvc Last_Medical_Exam_Valid_To,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateofBirth = (*xsdDate)(layout.T.DateofBirth)
	layout.DateofDeath = (*xsdDate)(layout.T.DateofDeath)
	layout.MaritalStatusDate = (*xsdDate)(layout.T.MaritalStatusDate)
	layout.LastMedicalExamDate = (*xsdDate)(layout.T.LastMedicalExamDate)
	layout.LastMedicalExamValidTo = (*xsdDate)(layout.T.LastMedicalExamValidTo)
	return e.EncodeElement(layout, start)
}
func (t *ChangePersonalInformationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ChangePersonalInformationDataType
	var overlay struct {
		*T
		DateofBirth            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
		DateofDeath            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
		MaritalStatusDate      *xsdDate `xml:"urn:com.workday/bsvc Marital_Status_Date,omitempty"`
		LastMedicalExamDate    *xsdDate `xml:"urn:com.workday/bsvc Last_Medical_Exam_Date,omitempty"`
		LastMedicalExamValidTo *xsdDate `xml:"urn:com.workday/bsvc Last_Medical_Exam_Valid_To,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateofBirth = (*xsdDate)(overlay.T.DateofBirth)
	overlay.DateofDeath = (*xsdDate)(overlay.T.DateofDeath)
	overlay.MaritalStatusDate = (*xsdDate)(overlay.T.MaritalStatusDate)
	overlay.LastMedicalExamDate = (*xsdDate)(overlay.T.LastMedicalExamDate)
	overlay.LastMedicalExamValidTo = (*xsdDate)(overlay.T.LastMedicalExamValidTo)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper for the Change Personal data sub business process that is part of the Hire and Edit Position business processes.
type ChangePersonalInformationSubBusinessProcessType struct {
	BusinessSubProcessParameters  *BusinessSubProcessParametersType  `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	ChangePersonalInformationData *ChangePersonalInformationDataType `xml:"urn:com.workday/bsvc Change_Personal_Information_Data,omitempty"`
}

// Wrapper element for the Preferred Name Change business process web service.
type ChangePreferredNameBusinessProcessDataType struct {
	PersonReference             *RoleObjectType                `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	UniversalIDReference        *UniversalIdentifierObjectType `xml:"urn:com.workday/bsvc Universal_ID_Reference,omitempty"`
	UseLegalNameAsPreferredName *bool                          `xml:"urn:com.workday/bsvc Use_Legal_Name_As_Preferred_Name,omitempty"`
	NameData                    *PersonNameDetailDataType      `xml:"urn:com.workday/bsvc Name_Data,omitempty"`
}

// Allows retrieving Preferred Name based on request criteria
type ChangePreferredNameRequestCriteriaType struct {
	PersonTypeCriteriaData *PersonTypeCriteriaType `xml:"urn:com.workday/bsvc Person_Type_Criteria_Data,omitempty"`
}

// Element containing the worker data for the Get Change Preferred Name response
type ChangePreferredNameResponseDataType struct {
	ChangePreferredName []ChangePreferredNameResponseWrapperType `xml:"urn:com.workday/bsvc Change_Preferred_Name,omitempty"`
}

// Wrapper element for the Change Preferred Name Business Process Data element. This is the element that contains data to load.
type ChangePreferredNameResponseWrapperType struct {
	PersonReference         *RoleObjectType                             `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	ChangePreferredNameData *ChangePreferredNameBusinessProcessDataType `xml:"urn:com.workday/bsvc Change_Preferred_Name_Data,omitempty"`
}

// This web service allows updates to the work space location of a filled Position.
type ChangeWorkSpaceRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	WorkSpaceChangeEventData  *WorkSpaceChangeEventDataType  `xml:"urn:com.workday/bsvc Work_Space_Change_Event_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with Event ID for the Change Work Space event.
type ChangeWorkSpaceResponseType struct {
	WorkSpaceChangeEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Work_Space_Change_Event_Reference,omitempty"`
	Version                       string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top element used for Check Position Budget as a sub business process.
type CheckPositionBudgetSubBusinessProcessType struct {
	BusinessSubProcessParameters *FinancialsBusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
}

// Citizenship Status Details
type CitizenshipStatusDataType struct {
	CitizenshipStatusID               string                            `xml:"urn:com.workday/bsvc Citizenship_Status_ID,omitempty"`
	CitizenshipStatusName             string                            `xml:"urn:com.workday/bsvc Citizenship_Status_Name,omitempty"`
	CitizenshipStatusDescription      string                            `xml:"urn:com.workday/bsvc Citizenship_Status_Description,omitempty"`
	CountryReference                  *CountryObjectType                `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CitizenshipStatusMappingReference *GlobalSetupDataMappingObjectType `xml:"urn:com.workday/bsvc Citizenship_Status_Mapping_Reference,omitempty"`
	Citizen                           *bool                             `xml:"urn:com.workday/bsvc Citizen,omitempty"`
	Inactive                          *bool                             `xml:"urn:com.workday/bsvc Inactive,omitempty"`
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

// Element containing the references of the Citizenship Statuses requested.
type CitizenshipStatusRequestReferencesType struct {
	CitizenshipStatusReference []CitizenshipStatusObjectType `xml:"urn:com.workday/bsvc Citizenship_Status_Reference"`
}

// Response Data for the Citizenship Status.
type CitizenshipStatusResponseDataType struct {
	CitizenshipStatus []CitizenshipStatusType `xml:"urn:com.workday/bsvc Citizenship_Status,omitempty"`
}

// Citizenship Status Details.
type CitizenshipStatusType struct {
	CitizenshipStatusReference *CitizenshipStatusObjectType `xml:"urn:com.workday/bsvc Citizenship_Status_Reference,omitempty"`
	CitizenshipStatusData      []CitizenshipStatusDataType  `xml:"urn:com.workday/bsvc Citizenship_Status_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ClassStandingDefinitionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ClassStandingDefinitionObjectType struct {
	ID         []ClassStandingDefinitionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for the Close Position business process data including the position to close and the event data.
type ClosePositionDataType struct {
	PositionReference *PositionRestrictionsObjectType `xml:"urn:com.workday/bsvc Position_Reference"`
	CloseEventData    []ClosePositionEventDataType    `xml:"urn:com.workday/bsvc Close_Event_Data"`
}

// Contains the detailed data for the closure.
type ClosePositionEventDataType struct {
	ReasonReference *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
	CloseDate       time.Time                                 `xml:"urn:com.workday/bsvc Close_Date"`
}

func (t *ClosePositionEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ClosePositionEventDataType
	var layout struct {
		*T
		CloseDate *xsdDate `xml:"urn:com.workday/bsvc Close_Date"`
	}
	layout.T = (*T)(t)
	layout.CloseDate = (*xsdDate)(&layout.T.CloseDate)
	return e.EncodeElement(layout, start)
}
func (t *ClosePositionEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ClosePositionEventDataType
	var overlay struct {
		*T
		CloseDate *xsdDate `xml:"urn:com.workday/bsvc Close_Date"`
	}
	overlay.T = (*T)(t)
	overlay.CloseDate = (*xsdDate)(&overlay.T.CloseDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for the Close Position business process.
type ClosePositionRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	CloseRestrictionData      *ClosePositionDataType         `xml:"urn:com.workday/bsvc Close_Restriction_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the id of the Event and the reference of the Position Restriction being closed.
type ClosePositionResponseType struct {
	EventReference               *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	PositionorHeadcountReference *PositionGroupObjectType    `xml:"urn:com.workday/bsvc Position_or_Headcount_Reference,omitempty"`
	Version                      string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Encapsulating element containing all Collective Agreement data.
type CollectiveAgreementDetailDataType struct {
	AssignEmployeeCollectiveAgreementEventReference *UniqueIdentifierObjectType           `xml:"urn:com.workday/bsvc Assign_Employee_Collective_Agreement_Event_Reference,omitempty"`
	EffectiveDate                                   *time.Time                            `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	EndAssignmentDate                               *time.Time                            `xml:"urn:com.workday/bsvc End_Assignment_Date,omitempty"`
	CollectiveAgreementData                         []CollectiveAgreementSnapshotDataType `xml:"urn:com.workday/bsvc Collective_Agreement_Data,omitempty"`
}

func (t *CollectiveAgreementDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CollectiveAgreementDetailDataType
	var layout struct {
		*T
		EffectiveDate     *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		EndAssignmentDate *xsdDate `xml:"urn:com.workday/bsvc End_Assignment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	layout.EndAssignmentDate = (*xsdDate)(layout.T.EndAssignmentDate)
	return e.EncodeElement(layout, start)
}
func (t *CollectiveAgreementDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CollectiveAgreementDetailDataType
	var overlay struct {
		*T
		EffectiveDate     *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		EndAssignmentDate *xsdDate `xml:"urn:com.workday/bsvc End_Assignment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	overlay.EndAssignmentDate = (*xsdDate)(overlay.T.EndAssignmentDate)
	return d.DecodeElement(&overlay, &start)
}

// Container for the collective agreement factor 1 and factor option 1 data.
type CollectiveAgreementFactor1DataType struct {
	CollectiveAgreementFactor1Reference       *CollectiveAgreementFactorObjectType       `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_1_Reference,omitempty"`
	CollectiveAgreementFactorOption1Reference *CollectiveAgreementFactorOptionObjectType `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_Option_1_Reference,omitempty"`
}

// Container for the collective agreement factor 2 and factor option 2 data.
type CollectiveAgreementFactor2DataType struct {
	CollectiveAgreementFactor2Reference       *CollectiveAgreementFactorObjectType       `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_2_Reference,omitempty"`
	CollectiveAgreementFactorOption2Reference *CollectiveAgreementFactorOptionObjectType `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_Option_2_Reference,omitempty"`
}

// Container for the collective agreement factor 3 and factor option 3 data.
type CollectiveAgreementFactor3DataType struct {
	CollectiveAgreementFactor3Reference       *CollectiveAgreementFactorObjectType       `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_3_Reference,omitempty"`
	CollectiveAgreementFactorOption3Reference *CollectiveAgreementFactorOptionObjectType `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_Option_3_Reference,omitempty"`
}

// Container for the collective agreement factor 4 and factor option 4 data.
type CollectiveAgreementFactor4DataType struct {
	CollectiveAgreementFactor4Reference       *CollectiveAgreementFactorObjectType       `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_4_Reference,omitempty"`
	CollectiveAgreementFactorOption4Reference *CollectiveAgreementFactorOptionObjectType `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_Option_4_Reference,omitempty"`
}

// Container for the collective agreement factor 5 and factor option 5 data.
type CollectiveAgreementFactor5DataType struct {
	CollectiveAgreementFactor5Reference       *CollectiveAgreementFactorObjectType       `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_5_Reference,omitempty"`
	CollectiveAgreementFactorOption5Reference *CollectiveAgreementFactorOptionObjectType `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_Option_5_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CollectiveAgreementFactorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CollectiveAgreementFactorObjectType struct {
	ID         []CollectiveAgreementFactorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CollectiveAgreementFactorOptionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CollectiveAgreementFactorOptionObjectType struct {
	ID         []CollectiveAgreementFactorOptionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper for collective agreement factors and factor options.
type CollectiveAgreementFactorOptionsDataType struct {
	CollectiveAgreementFactor1Data *CollectiveAgreementFactor1DataType `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_1_Data,omitempty"`
	CollectiveAgreementFactor2Data *CollectiveAgreementFactor2DataType `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_2_Data,omitempty"`
	CollectiveAgreementFactor3Data *CollectiveAgreementFactor3DataType `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_3_Data,omitempty"`
	CollectiveAgreementFactor4Data *CollectiveAgreementFactor4DataType `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_4_Data,omitempty"`
	CollectiveAgreementFactor5Data *CollectiveAgreementFactor5DataType `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_5_Data,omitempty"`
}

// Encapsulating Element for the Collective Agreement Snapshot Data.
type CollectiveAgreementFactorParameterDataType struct {
	Order                                    string                                      `xml:"urn:com.workday/bsvc Order,omitempty"`
	CollectiveAgreementFactorReference       *CollectiveAgreementFactorObjectType        `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_Reference,omitempty"`
	CollectiveAgreementFactorOptionReference []CollectiveAgreementFactorOptionObjectType `xml:"urn:com.workday/bsvc Collective_Agreement_Factor_Option_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CollectiveAgreementObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CollectiveAgreementObjectType struct {
	ID         []CollectiveAgreementObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating Element containing the Collective Agreement Classification Parameter Data.
type CollectiveAgreementSnapshotDataType struct {
	CollectiveAgreementReference *CollectiveAgreementObjectType               `xml:"urn:com.workday/bsvc Collective_Agreement_Reference,omitempty"`
	CollectiveAgreementFactor    []CollectiveAgreementFactorParameterDataType `xml:"urn:com.workday/bsvc Collective_Agreement_Factor,omitempty"`
}

// Encapsulating element containing all Collective Agreement data. Including Corrected data.
type CollectiveAgreementSummaryDataType struct {
	CollectiveAgreementData *CollectiveAgreementDetailDataType `xml:"urn:com.workday/bsvc Collective_Agreement_Data,omitempty"`
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
type CompanyInsiderTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompanyInsiderTypeObjectType struct {
	ID         []CompanyInsiderTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Data element containing annualized summary compensation information for a worker.
type CompensatableSummaryAmountAnnualizedDataType struct {
	TotalBasePay             float64              `xml:"urn:com.workday/bsvc Total_Base_Pay,omitempty"`
	TotalSalaryandAllowances float64              `xml:"urn:com.workday/bsvc Total_Salary_and_Allowances,omitempty"`
	PrimaryCompensationBasis float64              `xml:"urn:com.workday/bsvc Primary_Compensation_Basis,omitempty"`
	CurrencyReference        *CurrencyObjectType  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference       *FrequencyObjectType `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
}

// Data element containing annualized summary compensation information for a worker in the default currency.
type CompensatableSummaryAmountAnnualizedInReportingCurrencyDataType struct {
	TotalBasePay             float64              `xml:"urn:com.workday/bsvc Total_Base_Pay,omitempty"`
	TotalSalaryandAllowances float64              `xml:"urn:com.workday/bsvc Total_Salary_and_Allowances,omitempty"`
	PrimaryCompensationBasis float64              `xml:"urn:com.workday/bsvc Primary_Compensation_Basis,omitempty"`
	CurrencyReference        *CurrencyObjectType  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference       *FrequencyObjectType `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
}

// Data element containing summary compensation information for the worker.
type CompensatableSummaryAmountDataType struct {
	TotalBasePay             float64              `xml:"urn:com.workday/bsvc Total_Base_Pay,omitempty"`
	TotalSalaryandAllowances float64              `xml:"urn:com.workday/bsvc Total_Salary_and_Allowances,omitempty"`
	PrimaryCompensationBasis float64              `xml:"urn:com.workday/bsvc Primary_Compensation_Basis,omitempty"`
	CurrencyReference        *CurrencyObjectType  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference       *FrequencyObjectType `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
}

// Data element containing summary compensation information for the worker in an hourly frequency when the worker has at least one hourly pay plan assigned.
type CompensatableSummaryAmountHourlyDataType struct {
	TotalBasePay       float64              `xml:"urn:com.workday/bsvc Total_Base_Pay,omitempty"`
	CurrencyReference  *CurrencyObjectType  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference *FrequencyObjectType `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
}

// Data element containing summary compensation information for a worker converted to their pay group frequency.
type CompensatableSummaryAmountInTargetFrequencyDataType struct {
	TotalBasePay             float64              `xml:"urn:com.workday/bsvc Total_Base_Pay,omitempty"`
	TotalSalaryandAllowances float64              `xml:"urn:com.workday/bsvc Total_Salary_and_Allowances,omitempty"`
	PrimaryCompensationBasis float64              `xml:"urn:com.workday/bsvc Primary_Compensation_Basis,omitempty"`
	CurrencyReference        *CurrencyObjectType  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference       *FrequencyObjectType `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
}

// Wrapper element containing summary compensation information for a position.  The information is in the position's compensation currency and frequency, annualized, and annualized in the default currency.
type CompensatableSummaryDataType struct {
	EmployeeCompensationSummaryData          *CompensatableSummaryAmountDataType                              `xml:"urn:com.workday/bsvc Employee_Compensation_Summary_Data,omitempty"`
	AnnualizedSummaryData                    *CompensatableSummaryAmountAnnualizedDataType                    `xml:"urn:com.workday/bsvc Annualized_Summary_Data,omitempty"`
	SummaryDatainPayGroupFrequency           *CompensatableSummaryAmountInTargetFrequencyDataType             `xml:"urn:com.workday/bsvc Summary_Data_in_Pay_Group_Frequency,omitempty"`
	AnnualizedinReportingCurrencySummaryData *CompensatableSummaryAmountAnnualizedInReportingCurrencyDataType `xml:"urn:com.workday/bsvc Annualized_in_Reporting_Currency_Summary_Data,omitempty"`
	SummaryDatainHourlyFrequency             *CompensatableSummaryAmountHourlyDataType                        `xml:"urn:com.workday/bsvc Summary_Data_in_Hourly_Frequency,omitempty"`
}

// Data element that contains the compensation change information.
type CompensationChangeDataType struct {
	ReasonReference                       *EventClassificationSubcategoryObjectType             `xml:"urn:com.workday/bsvc Reason_Reference"`
	OverrideCompensationBasisCalculation  *bool                                                 `xml:"urn:com.workday/bsvc Override_Compensation_Basis_Calculation,omitempty"`
	PrimaryCompensationBasis              float64                                               `xml:"urn:com.workday/bsvc Primary_Compensation_Basis,omitempty"`
	PrimaryCompensationBasisAmountChange  float64                                               `xml:"urn:com.workday/bsvc Primary_Compensation_Basis_Amount_Change,omitempty"`
	PrimaryCompensationBasisPercentChange float64                                               `xml:"urn:com.workday/bsvc Primary_Compensation_Basis_Percent_Change,omitempty"`
	CompensationGuidelinesData            *CompensatableGuidelinesDataType                      `xml:"urn:com.workday/bsvc Compensation_Guidelines_Data,omitempty"`
	PayPlanData                           *ProposedBasePayPlanAssignmentContainerDataType       `xml:"urn:com.workday/bsvc Pay_Plan_Data,omitempty"`
	UnitSalaryPlanData                    *ProposedSalaryUnitPlanAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Unit_Salary_Plan_Data,omitempty"`
	AllowancePlanData                     *ProposedAllowancePlanAssignmentContainerDataType     `xml:"urn:com.workday/bsvc Allowance_Plan_Data,omitempty"`
	UnitAllowancePlanData                 *ProposedAllowanceUnitPlanAssignmentContainerDataType `xml:"urn:com.workday/bsvc Unit_Allowance_Plan_Data,omitempty"`
	BonusPlanData                         *ProposedBonusPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Bonus_Plan_Data,omitempty"`
	MeritPlanData                         *ProposedMeritPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Merit_Plan_Data,omitempty"`
	CommissionPlanData                    *ProposedCommissionPlanAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Commission_Plan_Data,omitempty"`
	StockPlanData                         *ProposedStockPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Stock_Plan_Data,omitempty"`
	RemovePlanData                        []RemoveCompensationPlanAssignmentDataType            `xml:"urn:com.workday/bsvc Remove_Plan_Data,omitempty"`
	PeriodSalaryPlanData                  *ProposedPeriodSalaryPlanAssignmentContainerDataType  `xml:"urn:com.workday/bsvc Period_Salary_Plan_Data,omitempty"`
	CalculatedPlanData                    *ProposedCalculatedPlanAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Calculated_Plan_Data,omitempty"`
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
type CompensationMatrixObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationMatrixObjectType struct {
	ID         []CompensationMatrixObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type CompensationPayEarningObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationPayEarningObjectType struct {
	ID         []CompensationPayEarningObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type CompensationPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationPlanObjectType struct {
	ID         []CompensationPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Data element for the Propose Compensation for Transfer business process.
type CompensationProposedForPositionDataType struct {
	InitializeUsingDefaultingLogic *bool                                                 `xml:"urn:com.workday/bsvc Initialize_Using_Defaulting_Logic,omitempty"`
	EmployeeVisibilityDate         *time.Time                                            `xml:"urn:com.workday/bsvc Employee_Visibility_Date,omitempty"`
	PrimaryCompensationBasis       float64                                               `xml:"urn:com.workday/bsvc Primary_Compensation_Basis,omitempty"`
	CompensationGuidelinesData     *CompensatableGuidelinesDataType                      `xml:"urn:com.workday/bsvc Compensation_Guidelines_Data,omitempty"`
	PayPlanData                    *ProposedBasePayPlanJobAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Pay_Plan_Data,omitempty"`
	UnitSalaryPlanData             *ProposedSalaryUnitPlanAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Unit_Salary_Plan_Data,omitempty"`
	AllowancePlanData              *ProposedAllowancePlanAssignmentContainerDataType     `xml:"urn:com.workday/bsvc Allowance_Plan_Data,omitempty"`
	UnitAllowancePlanData          *ProposedAllowanceUnitPlanAssignmentContainerDataType `xml:"urn:com.workday/bsvc Unit_Allowance_Plan_Data,omitempty"`
	BonusPlanData                  *ProposedBonusPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Bonus_Plan_Data,omitempty"`
	MeritPlanData                  *ProposedMeritPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Merit_Plan_Data,omitempty"`
	CommissionPlanData             *ProposedCommissionPlanAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Commission_Plan_Data,omitempty"`
	StockPlanData                  *ProposedStockPlanAssignmentContainerDataType         `xml:"urn:com.workday/bsvc Stock_Plan_Data,omitempty"`
	PeriodSalaryPlanData           *ProposedPeriodSalaryPlanAssignmentContainerDataType  `xml:"urn:com.workday/bsvc Period_Salary_Plan_Data,omitempty"`
	CalculatedPlanData             *ProposedCalculatedPlanAssignmentContainerDataType    `xml:"urn:com.workday/bsvc Calculated_Plan_Data,omitempty"`
}

func (t *CompensationProposedForPositionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CompensationProposedForPositionDataType
	var layout struct {
		*T
		EmployeeVisibilityDate *xsdDate `xml:"urn:com.workday/bsvc Employee_Visibility_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EmployeeVisibilityDate = (*xsdDate)(layout.T.EmployeeVisibilityDate)
	return e.EncodeElement(layout, start)
}
func (t *CompensationProposedForPositionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompensationProposedForPositionDataType
	var overlay struct {
		*T
		EmployeeVisibilityDate *xsdDate `xml:"urn:com.workday/bsvc Employee_Visibility_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EmployeeVisibilityDate = (*xsdDate)(overlay.T.EmployeeVisibilityDate)
	return d.DecodeElement(&overlay, &start)
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

// Contains a unique identifier for an instance of an object.
type ComponentCompletionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ComponentCompletionObjectType struct {
	ID         []ComponentCompletionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type ContingentWorkerObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ContingentWorkerObjectType struct {
	ID         []ContingentWorkerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contingent Worker Tax Authority Form Information
type ContingentWorkerTaxAuthorityFormInformationDataType struct {
	TaxAuthorityFormTypeReference *TaxAuthorityFormTypeObjectType `xml:"urn:com.workday/bsvc Tax_Authority_Form_Type_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ContingentWorkerTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ContingentWorkerTypeObjectType struct {
	ID         []ContingentWorkerTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for the Contract Contingent Worker business process web service and its sub business processes.
type ContractContingentWorkerBusinessProcessDataType struct {
	ApplicantReference                  *ApplicantObjectType                                  `xml:"urn:com.workday/bsvc Applicant_Reference"`
	FormerWorkerReference               *FormerWorkerObjectType                               `xml:"urn:com.workday/bsvc Former_Worker_Reference"`
	ApplicantData                       *CreateApplicantDataType                              `xml:"urn:com.workday/bsvc Applicant_Data"`
	OrganizationReference               *SupervisoryOrganizationObjectType                    `xml:"urn:com.workday/bsvc Organization_Reference,omitempty"`
	PositionReference                   *PositionRestrictionsObjectType                       `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	JobRequisitionReference             *JobRequisitionObjectType                             `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	ContractStartDate                   time.Time                                             `xml:"urn:com.workday/bsvc Contract_Start_Date"`
	ContractContingentWorkerEventData   *ContractContingentWorkerEventDataType                `xml:"urn:com.workday/bsvc Contract_Contingent_Worker_Event_Data"`
	UpdateIDInformationSubProcess       *UpdateIDInformationSubBusinessProcessType            `xml:"urn:com.workday/bsvc Update_ID_Information_Sub_Process,omitempty"`
	EditGovernmentIDsSubProcess         *EditGovernmentIDsSubBusinessProcessType              `xml:"urn:com.workday/bsvc Edit_Government_IDs_Sub_Process,omitempty"`
	EditPassportsandVisasSubProcess     *EditPassportsandVisasSubBusinessProcessType          `xml:"urn:com.workday/bsvc Edit_Passports_and_Visas_Sub_Process,omitempty"`
	EditLicenseSubProcess               *EditLicensesSubBusinessProcessType                   `xml:"urn:com.workday/bsvc Edit_License_Sub_Process,omitempty"`
	EditCustomIDsSubProcess             *EditCustomIDsSubBusinessProcessType                  `xml:"urn:com.workday/bsvc Edit_Custom_IDs_Sub_Process,omitempty"`
	EditAssignOrganizationSubProcess    *EditAssignPositionOrganizationSubBusinessProcessType `xml:"urn:com.workday/bsvc Edit_Assign_Organization_Sub_Process,omitempty"`
	CreateWorkdayAccountSubProcess      *CreateWorkdayAccountSubBusinessProcessType           `xml:"urn:com.workday/bsvc Create_Workday_Account_Sub_Process,omitempty"`
	AssignMatrixOrganizationSubProcess  *AssignMatrixOrganizationSubBusinessProcessType       `xml:"urn:com.workday/bsvc Assign_Matrix_Organization_Sub_Process,omitempty"`
	ChangePersonalInformationSubProcess *ChangePersonalInformationSubBusinessProcessType      `xml:"urn:com.workday/bsvc Change_Personal_Information_Sub_Process,omitempty"`
	EditServiceDatesSubProcess          *EditServiceDatesSubBusinessProcessType               `xml:"urn:com.workday/bsvc Edit_Service_Dates_Sub_Process,omitempty"`
	RemoveRetireeStatusSubProcess       *RemoveRetireeStatusSubBusinessProcessType            `xml:"urn:com.workday/bsvc Remove_Retiree_Status_Sub_Process,omitempty"`
	OnboardingSetupSubProcess           *OnboardingSetupSubBusinessProcessType                `xml:"urn:com.workday/bsvc Onboarding_Setup_Sub_Process,omitempty"`
}

func (t *ContractContingentWorkerBusinessProcessDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ContractContingentWorkerBusinessProcessDataType
	var layout struct {
		*T
		ContractStartDate *xsdDate `xml:"urn:com.workday/bsvc Contract_Start_Date"`
	}
	layout.T = (*T)(t)
	layout.ContractStartDate = (*xsdDate)(&layout.T.ContractStartDate)
	return e.EncodeElement(layout, start)
}
func (t *ContractContingentWorkerBusinessProcessDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ContractContingentWorkerBusinessProcessDataType
	var overlay struct {
		*T
		ContractStartDate *xsdDate `xml:"urn:com.workday/bsvc Contract_Start_Date"`
	}
	overlay.T = (*T)(t)
	overlay.ContractStartDate = (*xsdDate)(&overlay.T.ContractStartDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for the details of a contingent worker contract.
type ContractContingentWorkerEventDataType struct {
	ContingentWorkerID               string                                    `xml:"urn:com.workday/bsvc Contingent_Worker_ID,omitempty"`
	PositionID                       string                                    `xml:"urn:com.workday/bsvc Position_ID,omitempty"`
	ContractWorkerReasonReference    *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Contract_Worker_Reason_Reference,omitempty"`
	ContractWorkerTypeReference      *ContingentWorkerTypeObjectType           `xml:"urn:com.workday/bsvc Contract_Worker_Type_Reference,omitempty"`
	CreatePurchaseOrder              *bool                                     `xml:"urn:com.workday/bsvc Create_Purchase_Order,omitempty"`
	ContractEndDate                  *time.Time                                `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
	FirstDayofWork                   *time.Time                                `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
	ConversionPositionStartDate      *time.Time                                `xml:"urn:com.workday/bsvc Conversion_Position_Start_Date,omitempty"`
	SupplierReference                *SupplierObjectType                       `xml:"urn:com.workday/bsvc Supplier_Reference,omitempty"`
	DefaultPaymentTermReference      *PaymentTermsObjectType                   `xml:"urn:com.workday/bsvc Default_Payment_Term_Reference,omitempty"`
	PositionDetails                  *PositionDetailsSubDataType               `xml:"urn:com.workday/bsvc Position_Details"`
	ContractDetailsData              []ContractDetailsSubDataType              `xml:"urn:com.workday/bsvc Contract_Details_Data,omitempty"`
	PurchaseOrderContractDetailsData *PurchaseOrderContractDetailsDataType     `xml:"urn:com.workday/bsvc Purchase_Order_Contract_Details_Data,omitempty"`
	ContingentWorkerExternalIDData   *ExternalIDDataType                       `xml:"urn:com.workday/bsvc Contingent_Worker_External_ID_Data,omitempty"`
}

func (t *ContractContingentWorkerEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ContractContingentWorkerEventDataType
	var layout struct {
		*T
		ContractEndDate             *xsdDate `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
		FirstDayofWork              *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
		ConversionPositionStartDate *xsdDate `xml:"urn:com.workday/bsvc Conversion_Position_Start_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ContractEndDate = (*xsdDate)(layout.T.ContractEndDate)
	layout.FirstDayofWork = (*xsdDate)(layout.T.FirstDayofWork)
	layout.ConversionPositionStartDate = (*xsdDate)(layout.T.ConversionPositionStartDate)
	return e.EncodeElement(layout, start)
}
func (t *ContractContingentWorkerEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ContractContingentWorkerEventDataType
	var overlay struct {
		*T
		ContractEndDate             *xsdDate `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
		FirstDayofWork              *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
		ConversionPositionStartDate *xsdDate `xml:"urn:com.workday/bsvc Conversion_Position_Start_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ContractEndDate = (*xsdDate)(overlay.T.ContractEndDate)
	overlay.FirstDayofWork = (*xsdDate)(overlay.T.FirstDayofWork)
	overlay.ConversionPositionStartDate = (*xsdDate)(overlay.T.ConversionPositionStartDate)
	return d.DecodeElement(&overlay, &start)
}

// This web service operation is designed to contract an existing applicant into an existing Position.
type ContractContingentWorkerRequestHVType struct {
	BusinessProcessParameters    *BusinessProcessParametersType                   `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	ContractContingentWorkerData *ContractContingentWorkerBusinessProcessDataType `xml:"urn:com.workday/bsvc Contract_Contingent_Worker_Data"`
}

// This web service operation is designed to contract an existing applicant into an existing Position or Job Group.
type ContractContingentWorkerRequestType struct {
	BusinessProcessParameters    *BusinessProcessParametersType                   `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	ContractContingentWorkerData *ContractContingentWorkerBusinessProcessDataType `xml:"urn:com.workday/bsvc Contract_Contingent_Worker_Data"`
	Version                      string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the event id of the contract contingent worker event.
type ContractContingentWorkerResponseType struct {
	EventReference            *UniqueIdentifierObjectType                    `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	ContingentWorkerReference *WorkerObjectType                              `xml:"urn:com.workday/bsvc Contingent_Worker_Reference,omitempty"`
	PositionReference         *PositionElementObjectType                     `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	ExceptionsResponseData    []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                   string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for contingent worker contract details for a filled position.
// Only use for contingent workers.
// If the element is not sent, the existing contract detail data will be kept.
type ContractDetailsSubDataType struct {
	ContractPayRate           float64              `xml:"urn:com.workday/bsvc Contract_Pay_Rate,omitempty"`
	CurrencyReference         *CurrencyObjectType  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference        *FrequencyObjectType `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	ContractAssignmentDetails string               `xml:"urn:com.workday/bsvc Contract_Assignment_Details,omitempty"`
}

// Contains the contribution amount for the retirement savings election.
type ContributionAmountDataType struct {
	ContributionAmount float64              `xml:"urn:com.workday/bsvc Contribution_Amount"`
	FrequencyReference *FrequencyObjectType `xml:"urn:com.workday/bsvc Frequency_Reference"`
}

// Contains the contribtion information for the spending account election.
type ContributionDataType struct {
	ContributionAmount float64              `xml:"urn:com.workday/bsvc Contribution_Amount,omitempty"`
	FrequencyReference *FrequencyObjectType `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
}

// Transaction Data for an event that has been rescinded or corrected.
type CorrectedOrRescindedTransactionDataType struct {
	MainTransactionData     []MainTransactionLogEntryDataType              `xml:"urn:com.workday/bsvc Main_Transaction_Data,omitempty"`
	CorrectionOrRescindData []TransactionLogRescindAndCorrectEventDataType `xml:"urn:com.workday/bsvc Correction_Or_Rescind_Data,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type CourseDefinitionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CourseDefinitionObjectType struct {
	ID         []CourseDefinitionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CourseSubjectObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CourseSubjectObjectType struct {
	ID         []CourseSubjectObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the detailed insurance coverage level information for an election.
type CoverageLevelValueDataType struct {
	CoverageLevelValue      float64             `xml:"urn:com.workday/bsvc Coverage_Level_Value"`
	BuyUpAmount             float64             `xml:"urn:com.workday/bsvc Buy_Up_Amount,omitempty"`
	GuaranteeIssueAmount    float64             `xml:"urn:com.workday/bsvc Guarantee_Issue_Amount,omitempty"`
	CoverageLevelMultiplier float64             `xml:"urn:com.workday/bsvc Coverage_Level_Multiplier"`
	CoverageLevelType       string              `xml:"urn:com.workday/bsvc Coverage_Level_Type"`
	CurrencyReference       *CurrencyObjectType `xml:"urn:com.workday/bsvc Currency_Reference"`
}

// Data for a new Academic Affiliate if an affiliate is to be created as part of the appointment.
type CreateAcademicAffiliateDataType struct {
	ID                        string                              `xml:"urn:com.workday/bsvc ID,omitempty"`
	PersonNameData            *PersonNameDataType                 `xml:"urn:com.workday/bsvc Person_Name_Data,omitempty"`
	ContactInformationData    *ContactInformationDataType         `xml:"urn:com.workday/bsvc Contact_Information_Data,omitempty"`
	PersonalInformationData   *GenericPersonalInformationDataType `xml:"urn:com.workday/bsvc Personal_Information_Data,omitempty"`
	PersonIdentificationData  *PersonIdentificationDataType       `xml:"urn:com.workday/bsvc Person_Identification_Data,omitempty"`
	PersonPhotoData           *PersonPhotoDataType                `xml:"urn:com.workday/bsvc Person_Photo_Data,omitempty"`
	LocationContextsReference []LocationContextObjectType         `xml:"urn:com.workday/bsvc Location_Contexts_Reference,omitempty"`
}

// Data for a new Applicant if an applicant is to be created as part of the hire.
type CreateApplicantDataType struct {
	ApplicantID               string                         `xml:"urn:com.workday/bsvc Applicant_ID,omitempty"`
	PersonalData              *PersonalInformationDataType   `xml:"urn:com.workday/bsvc Personal_Data,omitempty"`
	ExternalIntegrationIDData *ExternalIntegrationIDDataType `xml:"urn:com.workday/bsvc External_Integration_ID_Data,omitempty"`
}

// Container for the data in the Create Benefit Life Event sub business process.
type CreateBenefitLifeEventDataType struct {
	LifeEventTypeReference *BenefitLifeEventTypeObjectType `xml:"urn:com.workday/bsvc Life_Event_Type_Reference"`
	LifeEventDate          *time.Time                      `xml:"urn:com.workday/bsvc Life_Event_Date,omitempty"`
	SubmitElectionsBy      *time.Time                      `xml:"urn:com.workday/bsvc Submit_Elections_By,omitempty"`
}

func (t *CreateBenefitLifeEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CreateBenefitLifeEventDataType
	var layout struct {
		*T
		LifeEventDate     *xsdDate `xml:"urn:com.workday/bsvc Life_Event_Date,omitempty"`
		SubmitElectionsBy *xsdDate `xml:"urn:com.workday/bsvc Submit_Elections_By,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LifeEventDate = (*xsdDate)(layout.T.LifeEventDate)
	layout.SubmitElectionsBy = (*xsdDate)(layout.T.SubmitElectionsBy)
	return e.EncodeElement(layout, start)
}
func (t *CreateBenefitLifeEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CreateBenefitLifeEventDataType
	var overlay struct {
		*T
		LifeEventDate     *xsdDate `xml:"urn:com.workday/bsvc Life_Event_Date,omitempty"`
		SubmitElectionsBy *xsdDate `xml:"urn:com.workday/bsvc Submit_Elections_By,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LifeEventDate = (*xsdDate)(overlay.T.LifeEventDate)
	overlay.SubmitElectionsBy = (*xsdDate)(overlay.T.SubmitElectionsBy)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for the Create Benefit Life Event sub business process.
type CreateBenefitLifeEventSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	CreateBenefitLifeEventData   *CreateBenefitLifeEventDataType   `xml:"urn:com.workday/bsvc Create_Benefit_Life_Event_Data,omitempty"`
}

// Wrapper element for the sub-process data and Job Requisition data.
type CreateJobRequisitionSubProcessDataType struct {
	CreateJobRequisitionReasonReference                  *EventClassificationSubcategoryObjectType  `xml:"urn:com.workday/bsvc Create_Job_Requisition_Reason_Reference,omitempty"`
	JobRequisitionData                                   *JobRequisitionDataforCreateandEditType    `xml:"urn:com.workday/bsvc Job_Requisition_Data,omitempty"`
	CheckPositionBudgetSubProcessforCreateJobRequisition *CheckPositionBudgetSubBusinessProcessType `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process_for_Create_Job_Requisition,omitempty"`
}

// Wrapper element for the Create Job Requisition Business Process Parameters and Data.
type CreateJobRequisitionSubProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType       `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	CreateJobRequisitionData     *CreateJobRequisitionSubProcessDataType `xml:"urn:com.workday/bsvc Create_Job_Requisition_Data,omitempty"`
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

// Launch the create provisioning system.
// Uses the Create Provisioning System business process.
// This subprocess can be skipped, processed manually or processed automatically from the web service.
type CreateProvisioningEventSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
}

// Contains information for the subordinate organization
type CreateSubordinateEventDataType struct {
	AvailabilityDate              *time.Time                     `xml:"urn:com.workday/bsvc Availability_Date,omitempty"`
	OrganizationName              string                         `xml:"urn:com.workday/bsvc Organization_Name,omitempty"`
	IncludeOrganizationIDinName   *bool                          `xml:"urn:com.workday/bsvc Include_Organization_ID_in_Name,omitempty"`
	OrganizationCode              string                         `xml:"urn:com.workday/bsvc Organization_Code,omitempty"`
	IncludeOrganizationCodeinName *bool                          `xml:"urn:com.workday/bsvc Include_Organization_Code_in_Name,omitempty"`
	IncludeManagerLeaderinName    *bool                          `xml:"urn:com.workday/bsvc Include_Manager_Leader_in_Name,omitempty"`
	OrganizationSubtypeReference  *OrganizationSubtypeObjectType `xml:"urn:com.workday/bsvc Organization_Subtype_Reference"`
	ExternalURLReference          *ExternalURLObjectType         `xml:"urn:com.workday/bsvc External_URL_Reference,omitempty"`
	PrimaryLocationReference      *LocationObjectType            `xml:"urn:com.workday/bsvc Primary_Location_Reference,omitempty"`
}

func (t *CreateSubordinateEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CreateSubordinateEventDataType
	var layout struct {
		*T
		AvailabilityDate *xsdDate `xml:"urn:com.workday/bsvc Availability_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AvailabilityDate = (*xsdDate)(layout.T.AvailabilityDate)
	return e.EncodeElement(layout, start)
}
func (t *CreateSubordinateEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CreateSubordinateEventDataType
	var overlay struct {
		*T
		AvailabilityDate *xsdDate `xml:"urn:com.workday/bsvc Availability_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AvailabilityDate = (*xsdDate)(overlay.T.AvailabilityDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper for the Create Subordinate sub process
type CreateSubordinateSubBusinessProcessType struct {
	BusinessSubProcessParameters []BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	CreateSubordinateEventData   []CreateSubordinateEventDataType   `xml:"urn:com.workday/bsvc Create_Subordinate_Event_Data,omitempty"`
}

// Wrapper to hold the business process configuration and user account details.
type CreateWorkdayAccountSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	CreateWorkdayAccountData     *SystemUserDataType               `xml:"urn:com.workday/bsvc Create_Workday_Account_Data,omitempty"`
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

// Wrapper for Custom Identification Data. Includes Custom Identifiers.
type CustomIdentificationDataType struct {
	CustomID   []CustomIDType `xml:"urn:com.workday/bsvc Custom_ID,omitempty"`
	ReplaceAll bool           `xml:"urn:com.workday/bsvc Replace_All,attr,omitempty"`
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
type DeliveryModeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DeliveryModeObjectType struct {
	ID         []DeliveryModeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the dependents covered for the election.
type DependentCoverageDataType struct {
	DependentReference                              *DependentObjectType       `xml:"urn:com.workday/bsvc Dependent_Reference,omitempty"`
	OriginalCoverageBeginDate                       *time.Time                 `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
	CoverageEndDate                                 *time.Time                 `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	OriginalCoverageBeginDateforBenefitCoverageType *time.Time                 `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date_for_Benefit_Coverage_Type,omitempty"`
	FirstDateCoveredbyBenefitPlan                   *time.Time                 `xml:"urn:com.workday/bsvc First_Date_Covered_by_Benefit_Plan,omitempty"`
	FirstDateCoveredbyBenefitProvider               *time.Time                 `xml:"urn:com.workday/bsvc First_Date_Covered_by_Benefit_Provider,omitempty"`
	ProviderID                                      string                     `xml:"urn:com.workday/bsvc Provider_ID,omitempty"`
	COBRAEligibilityData                            []COBRAEligibilityDataType `xml:"urn:com.workday/bsvc COBRA_Eligibility_Data,omitempty"`
}

func (t *DependentCoverageDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DependentCoverageDataType
	var layout struct {
		*T
		OriginalCoverageBeginDate                       *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
		CoverageEndDate                                 *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
		OriginalCoverageBeginDateforBenefitCoverageType *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date_for_Benefit_Coverage_Type,omitempty"`
		FirstDateCoveredbyBenefitPlan                   *xsdDate `xml:"urn:com.workday/bsvc First_Date_Covered_by_Benefit_Plan,omitempty"`
		FirstDateCoveredbyBenefitProvider               *xsdDate `xml:"urn:com.workday/bsvc First_Date_Covered_by_Benefit_Provider,omitempty"`
	}
	layout.T = (*T)(t)
	layout.OriginalCoverageBeginDate = (*xsdDate)(layout.T.OriginalCoverageBeginDate)
	layout.CoverageEndDate = (*xsdDate)(layout.T.CoverageEndDate)
	layout.OriginalCoverageBeginDateforBenefitCoverageType = (*xsdDate)(layout.T.OriginalCoverageBeginDateforBenefitCoverageType)
	layout.FirstDateCoveredbyBenefitPlan = (*xsdDate)(layout.T.FirstDateCoveredbyBenefitPlan)
	layout.FirstDateCoveredbyBenefitProvider = (*xsdDate)(layout.T.FirstDateCoveredbyBenefitProvider)
	return e.EncodeElement(layout, start)
}
func (t *DependentCoverageDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DependentCoverageDataType
	var overlay struct {
		*T
		OriginalCoverageBeginDate                       *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
		CoverageEndDate                                 *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
		OriginalCoverageBeginDateforBenefitCoverageType *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date_for_Benefit_Coverage_Type,omitempty"`
		FirstDateCoveredbyBenefitPlan                   *xsdDate `xml:"urn:com.workday/bsvc First_Date_Covered_by_Benefit_Plan,omitempty"`
		FirstDateCoveredbyBenefitProvider               *xsdDate `xml:"urn:com.workday/bsvc First_Date_Covered_by_Benefit_Provider,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.OriginalCoverageBeginDate = (*xsdDate)(overlay.T.OriginalCoverageBeginDate)
	overlay.CoverageEndDate = (*xsdDate)(overlay.T.CoverageEndDate)
	overlay.OriginalCoverageBeginDateforBenefitCoverageType = (*xsdDate)(overlay.T.OriginalCoverageBeginDateforBenefitCoverageType)
	overlay.FirstDateCoveredbyBenefitPlan = (*xsdDate)(overlay.T.FirstDateCoveredbyBenefitPlan)
	overlay.FirstDateCoveredbyBenefitProvider = (*xsdDate)(overlay.T.FirstDateCoveredbyBenefitProvider)
	return d.DecodeElement(&overlay, &start)
}

// Contains the detailed information for a dependent.
type DependentDataType struct {
	DependentID                   string                                               `xml:"urn:com.workday/bsvc Dependent_ID,omitempty"`
	FulltimeStudent               *bool                                                `xml:"urn:com.workday/bsvc Full-time_Student,omitempty"`
	StudentStatusStartDate        *time.Time                                           `xml:"urn:com.workday/bsvc Student_Status_Start_Date,omitempty"`
	StudentStatusEndDate          *time.Time                                           `xml:"urn:com.workday/bsvc Student_Status_End_Date,omitempty"`
	Disabled                      *bool                                                `xml:"urn:com.workday/bsvc Disabled,omitempty"`
	InactiveDate                  *time.Time                                           `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
	DependentforPayrollPurposes   *bool                                                `xml:"urn:com.workday/bsvc Dependent_for_Payroll_Purposes,omitempty"`
	CitizenshipStatusReference    []CitizenshipStatusObjectType                        `xml:"urn:com.workday/bsvc Citizenship_Status_Reference,omitempty"`
	CountryofNationalityReference *CountryObjectType                                   `xml:"urn:com.workday/bsvc Country_of_Nationality_Reference,omitempty"`
	CountryofBirthReference       *CountryObjectType                                   `xml:"urn:com.workday/bsvc Country_of_Birth_Reference,omitempty"`
	RegionofBirthReference        *CountryRegionObjectType                             `xml:"urn:com.workday/bsvc Region_of_Birth_Reference,omitempty"`
	CityofBirth                   string                                               `xml:"urn:com.workday/bsvc City_of_Birth,omitempty"`
	CourtOrder                    []QualifiedDomesticRelationsOrderReplacementDataType `xml:"urn:com.workday/bsvc Court_Order,omitempty"`
	LivesWithWorkerData           []LivesWithWorkerDataType                            `xml:"urn:com.workday/bsvc Lives_With_Worker_Data,omitempty"`
	HasHealthInsuranceData        []HasHealthInsuranceDataType                         `xml:"urn:com.workday/bsvc Has_Health_Insurance_Data,omitempty"`
	AllowedforTaxDeductionData    []AllowedforTaxDeductionDataType                     `xml:"urn:com.workday/bsvc Allowed_for_Tax_Deduction_Data,omitempty"`
	AnnualIncomeData              []AnnualIncomeDataType                               `xml:"urn:com.workday/bsvc Annual_Income_Data,omitempty"`
	OccupationData                []OccupationDataType                                 `xml:"urn:com.workday/bsvc Occupation_Data,omitempty"`
	DisabilityData                []DisabilityInformationDataforRelatedPersonType      `xml:"urn:com.workday/bsvc Disability_Data,omitempty"`
}

func (t *DependentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DependentDataType
	var layout struct {
		*T
		StudentStatusStartDate *xsdDate `xml:"urn:com.workday/bsvc Student_Status_Start_Date,omitempty"`
		StudentStatusEndDate   *xsdDate `xml:"urn:com.workday/bsvc Student_Status_End_Date,omitempty"`
		InactiveDate           *xsdDate `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StudentStatusStartDate = (*xsdDate)(layout.T.StudentStatusStartDate)
	layout.StudentStatusEndDate = (*xsdDate)(layout.T.StudentStatusEndDate)
	layout.InactiveDate = (*xsdDate)(layout.T.InactiveDate)
	return e.EncodeElement(layout, start)
}
func (t *DependentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DependentDataType
	var overlay struct {
		*T
		StudentStatusStartDate *xsdDate `xml:"urn:com.workday/bsvc Student_Status_Start_Date,omitempty"`
		StudentStatusEndDate   *xsdDate `xml:"urn:com.workday/bsvc Student_Status_End_Date,omitempty"`
		InactiveDate           *xsdDate `xml:"urn:com.workday/bsvc Inactive_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StudentStatusStartDate = (*xsdDate)(overlay.T.StudentStatusStartDate)
	overlay.StudentStatusEndDate = (*xsdDate)(overlay.T.StudentStatusEndDate)
	overlay.InactiveDate = (*xsdDate)(overlay.T.InactiveDate)
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

// Encapsulating element containing all Dependent data.
type DependentType struct {
	DependentReference *DependentObjectType `xml:"urn:com.workday/bsvc Dependent_Reference"`
	DependentData      *DependentDataType   `xml:"urn:com.workday/bsvc Dependent_Data"`
}

// Contains a unique identifier for an instance of an object.
type DevelopmentItemCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DevelopmentItemCategoryObjectType struct {
	ID         []DevelopmentItemCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DevelopmentItemObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DevelopmentItemObjectType struct {
	ID         []DevelopmentItemObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DevelopmentItemStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DevelopmentItemStatusObjectType struct {
	ID         []DevelopmentItemStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains detailed data for the referenced Development Item
type DevelopmentItemVersionDataType struct {
	ID                            string                              `xml:"urn:com.workday/bsvc ID,omitempty"`
	DevelopmentItem               string                              `xml:"urn:com.workday/bsvc Development_Item"`
	AdditionalInformation         string                              `xml:"urn:com.workday/bsvc Additional_Information,omitempty"`
	CategoryReference             []DevelopmentItemCategoryObjectType `xml:"urn:com.workday/bsvc Category_Reference,omitempty"`
	DevelopmentItemStartDate      *time.Time                          `xml:"urn:com.workday/bsvc Development_Item_Start_Date,omitempty"`
	DevelopmentItemCompletionDate *time.Time                          `xml:"urn:com.workday/bsvc Development_Item_Completion_Date,omitempty"`
	StatusReference               *DevelopmentItemStatusObjectType    `xml:"urn:com.workday/bsvc Status_Reference"`
	StatusNote                    string                              `xml:"urn:com.workday/bsvc Status_Note,omitempty"`
	UpdatedbyWorkerReference      *WorkerObjectType                   `xml:"urn:com.workday/bsvc Updated_by_Worker_Reference,omitempty"`
	RelatesToReference            []TalentTagObjectType               `xml:"urn:com.workday/bsvc Relates_To_Reference,omitempty"`
}

func (t *DevelopmentItemVersionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DevelopmentItemVersionDataType
	var layout struct {
		*T
		DevelopmentItemStartDate      *xsdDate `xml:"urn:com.workday/bsvc Development_Item_Start_Date,omitempty"`
		DevelopmentItemCompletionDate *xsdDate `xml:"urn:com.workday/bsvc Development_Item_Completion_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DevelopmentItemStartDate = (*xsdDate)(layout.T.DevelopmentItemStartDate)
	layout.DevelopmentItemCompletionDate = (*xsdDate)(layout.T.DevelopmentItemCompletionDate)
	return e.EncodeElement(layout, start)
}
func (t *DevelopmentItemVersionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DevelopmentItemVersionDataType
	var overlay struct {
		*T
		DevelopmentItemStartDate      *xsdDate `xml:"urn:com.workday/bsvc Development_Item_Start_Date,omitempty"`
		DevelopmentItemCompletionDate *xsdDate `xml:"urn:com.workday/bsvc Development_Item_Completion_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DevelopmentItemStartDate = (*xsdDate)(overlay.T.DevelopmentItemStartDate)
	overlay.DevelopmentItemCompletionDate = (*xsdDate)(overlay.T.DevelopmentItemCompletionDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains the latest completed development plan for the employee.
type DevelopmentPlanDataType struct {
	DevelopmentPlanReference *DevelopmentPlanObjectType     `xml:"urn:com.workday/bsvc Development_Plan_Reference,omitempty"`
	ReviewData               *EmployeeReviewDetailsDataType `xml:"urn:com.workday/bsvc Review_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DevelopmentPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DevelopmentPlanObjectType struct {
	ID         []DevelopmentPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Wrapper element for all Disability Status data for the worker.
type DisabilityInformationDataType struct {
	DisabilityStatusInformationData []DisabilityStatusInformationDataType `xml:"urn:com.workday/bsvc Disability_Status_Information_Data,omitempty"`
	ReplaceAll                      bool                                  `xml:"urn:com.workday/bsvc Replace_All,attr,omitempty"`
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

// Wrapper element for each disability status entry.
type DisabilityStatusInformationDataType struct {
	DisabilityStatusReference *DisabilityStatusReferenceObjectType `xml:"urn:com.workday/bsvc Disability_Status_Reference,omitempty"`
	DisabilityStatusData      *DisabilityStatusSubDataType         `xml:"urn:com.workday/bsvc Disability_Status_Data,omitempty"`
	Delete                    bool                                 `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
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
type DisabilityStatusSubDataType struct {
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
}

func (t *DisabilityStatusSubDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DisabilityStatusSubDataType
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
func (t *DisabilityStatusSubDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DisabilityStatusSubDataType
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

// Contains the latest completed disciplinary action for the employee.
type DisciplinaryActionDataType struct {
	DisciplinaryActionReference          *DisciplinaryActionObjectType        `xml:"urn:com.workday/bsvc Disciplinary_Action_Reference,omitempty"`
	DisciplinaryActionReasonReference    []DisciplinaryActionReasonObjectType `xml:"urn:com.workday/bsvc Disciplinary_Action_Reason_Reference,omitempty"`
	DisciplinaryActionRelatedToReference []DisciplinaryActionObjectType       `xml:"urn:com.workday/bsvc Disciplinary_Action_Related_To_Reference,omitempty"`
	ReviewData                           *EmployeeReviewDetailsDataType       `xml:"urn:com.workday/bsvc Review_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DisciplinaryActionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DisciplinaryActionObjectType struct {
	ID         []DisciplinaryActionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DisciplinaryActionReasonObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DisciplinaryActionReasonObjectType struct {
	ID         []DisciplinaryActionReasonObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Container for the data changed in the Edit Background Check business process.
type EditBackgroundCheckDataType struct {
	StatusDate      time.Time                        `xml:"urn:com.workday/bsvc Status_Date"`
	StatusReference *BackgroundCheckStatusObjectType `xml:"urn:com.workday/bsvc Status_Reference"`
	StatusComment   string                           `xml:"urn:com.workday/bsvc Status_Comment,omitempty"`
}

func (t *EditBackgroundCheckDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EditBackgroundCheckDataType
	var layout struct {
		*T
		StatusDate *xsdDate `xml:"urn:com.workday/bsvc Status_Date"`
	}
	layout.T = (*T)(t)
	layout.StatusDate = (*xsdDate)(&layout.T.StatusDate)
	return e.EncodeElement(layout, start)
}
func (t *EditBackgroundCheckDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EditBackgroundCheckDataType
	var overlay struct {
		*T
		StatusDate *xsdDate `xml:"urn:com.workday/bsvc Status_Date"`
	}
	overlay.T = (*T)(t)
	overlay.StatusDate = (*xsdDate)(&overlay.T.StatusDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper for the Edit Background Check sub business process that is part of the Hire and Edit Position business processes.
type EditBackgroundCheckSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	EditBackgroundCheckData      *EditBackgroundCheckDataType      `xml:"urn:com.workday/bsvc Edit_Background_Check_Data,omitempty"`
}

// The main element for the Edit Other IDs Business Process. This contains all the Custom IDs and Business Processing processing parameters.
type EditCustomIDsSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	CustomIdentificationData     []CustomIdentificationDataType    `xml:"urn:com.workday/bsvc Custom_Identification_Data,omitempty"`
}

// Allows updating of effective dated custom objects for an Employee Contract.  When updating custom objects that allow multiple instances, data for existing instances must appear in the request, otherwise it will be removed.
type EditEmployeeContractAdditionalDataRequestType struct {
	BusinessProcessParameters        *BusinessProcessParametersType         `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EmployeeContractCustomObjectData []EmployeeContractCustomObjectDataType `xml:"urn:com.workday/bsvc Employee_Contract_Custom_Object_Data,omitempty"`
	Version                          string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Employee Contract effective-dated custom object data
type EditEmployeeContractAdditionalDataResponseType struct {
	EventReference            *UniqueIdentifierObjectType                  `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	EmployeeContractReference *EmployeeContractObjectType                  `xml:"urn:com.workday/bsvc Employee_Contract_Reference,omitempty"`
	EffectiveDate             *time.Time                                   `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	AdditionalData            []EffectiveDatedWebServiceAdditionalDataType `xml:"urn:com.workday/bsvc Additional_Data,omitempty"`
	Version                   string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

func (t *EditEmployeeContractAdditionalDataResponseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EditEmployeeContractAdditionalDataResponseType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EditEmployeeContractAdditionalDataResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EditEmployeeContractAdditionalDataResponseType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// The main element for the Edit Government IDs Business Process. This contains all the Government IDs information  and Business Processing processing parameters.
type EditGovernmentIDsSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType  `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	GovernmentIdentificationData []GovernmentIdentificationDataType `xml:"urn:com.workday/bsvc Government_Identification_Data,omitempty"`
}

// Wrapper element for editting hiring restrictions
type EditHiringRestrictionsDataType struct {
	SupervisoryOrganizationReference      *SupervisoryOrganizationObjectType                `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference"`
	HiringRestrictionsEditReasonReference *EventClassificationSubcategoryObjectType         `xml:"urn:com.workday/bsvc Hiring_Restrictions_Edit_Reason_Reference,omitempty"`
	PositionGroupRestrictionsData         *PositionGroupRestrictionDataType                 `xml:"urn:com.workday/bsvc Position_Group_Restrictions_Data"`
	RequestDefaultCompensationSubProcess  *RequestCompensationDefaultSubBusinessProcessType `xml:"urn:com.workday/bsvc Request_Default_Compensation_Sub_Process,omitempty"`
}

// Wrapper element for edit hiring restrictions web service
type EditHiringRestrictionsRequestType struct {
	BusinessProcessParameters  *BusinessProcessParametersType  `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EditHiringRestrictionsData *EditHiringRestrictionsDataType `xml:"urn:com.workday/bsvc Edit_Hiring_Restrictions_Data"`
	Version                    string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the event ID for the Job Restrictions Edit Event, and the Supervisory Organization reference.
type EditHiringRestrictionsResponseType struct {
	EventReference                   *UniqueIdentifierObjectType        `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	SupervisoryOrganizationReference *SupervisoryOrganizationObjectType `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference,omitempty"`
	Version                          string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The main element for the Edit Licenses Business Process. This contains all the License IDs and Business Processing processing parameters.
type EditLicensesSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	LicenseIdentificationData    []LicenseIdentificationDataType   `xml:"urn:com.workday/bsvc License_Identification_Data,omitempty"`
}

// Contains Notice Period Data
type EditNoticePeriodsDataType struct {
	EmployerNoticePeriodData *EmployerNoticePeriodDataType `xml:"urn:com.workday/bsvc Employer_Notice_Period_Data,omitempty"`
	EmployeeNoticePeriodData *EmployeeNoticePeriodDataType `xml:"urn:com.workday/bsvc Employee_Notice_Period_Data,omitempty"`
}

// Wrapper element for Edit Notice Periods sub business process
type EditNoticePeriodsSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	EditNoticePeriodsData        *EditNoticePeriodsDataType        `xml:"urn:com.workday/bsvc Edit_Notice_Periods_Data,omitempty"`
}

// The main element for the Edit Passports and Visas Business Process. This contains all the Passports and Visas IDs.
type EditPassportsandVisasSubBusinessProcessType struct {
	BusinessSubProcessParameters        *BusinessSubProcessParametersType         `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	PassportsandVisasIdentificationData []PassportsandVisasIdentificationDataType `xml:"urn:com.workday/bsvc Passports_and_Visas_Identification_Data,omitempty"`
}

// Wrapper for the Edit Position Business Process Web Service and its sub business processes.
type EditPositionDataType struct {
	WorkerReference                             *WorkerObjectType                                         `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	PositionReference                           *PositionElementObjectType                                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	EffectiveDate                               time.Time                                                 `xml:"urn:com.workday/bsvc Effective_Date"`
	EditPositionEventData                       *EditPositionEventDataType                                `xml:"urn:com.workday/bsvc Edit_Position_Event_Data"`
	RequestCompensationChangeSubProcess         *RequestCompensationForEditPositionSubBusinessProcessType `xml:"urn:com.workday/bsvc Request_Compensation_Change_Sub_Process,omitempty"`
	EditAssignOrganizationSubProcess            *EditAssignPositionOrganizationSubBusinessProcessType     `xml:"urn:com.workday/bsvc Edit_Assign_Organization_Sub_Process,omitempty"`
	AssignPayGroupSubProcess                    *AssignPayGroupSubBusinessProcessType                     `xml:"urn:com.workday/bsvc Assign_Pay_Group_Sub_Process,omitempty"`
	ReviewPayrollInterfaceSubProcess            *ReviewPayrollInterfaceDataSubBusinessProcessType         `xml:"urn:com.workday/bsvc Review_Payroll_Interface_Sub_Process,omitempty"`
	AssignMatrixOrganizationSubProcess          *AssignMatrixOrganizationSubBusinessProcessType           `xml:"urn:com.workday/bsvc Assign_Matrix_Organization_Sub_Process,omitempty"`
	ChangePersonalInformationSubProcess         *ChangePersonalInformationSubBusinessProcessType          `xml:"urn:com.workday/bsvc Change_Personal_Information_Sub_Process,omitempty"`
	RequestDefaultCompensationSubProcess        *RequestCompensationDefaultSubBusinessProcessType         `xml:"urn:com.workday/bsvc Request_Default_Compensation_Sub_Process,omitempty"`
	EditServiceDatesSubProcess                  *EditServiceDatesSubBusinessProcessType                   `xml:"urn:com.workday/bsvc Edit_Service_Dates_Sub_Process,omitempty"`
	RemoveRetireeStatusSubProcess               *RemoveRetireeStatusSubBusinessProcessType                `xml:"urn:com.workday/bsvc Remove_Retiree_Status_Sub_Process,omitempty"`
	MaintainEmployeeContractsSubBusinessProcess *MaintainEmployeeContractsSubBusinessProcessType          `xml:"urn:com.workday/bsvc Maintain_Employee_Contracts_Sub_Business_Process,omitempty"`
	CheckPositionBudgetSubProcess               *CheckPositionBudgetSubBusinessProcessType                `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process,omitempty"`
	UpdateAcademicAppointmentSubProcess         *UpdateAcademicAppointmentSubProcessType                  `xml:"urn:com.workday/bsvc Update_Academic_Appointment_Sub_Process,omitempty"`
	AssignCostingAllocationSubProcess           *AssignCostingAllocationSubBusinessProcessType            `xml:"urn:com.workday/bsvc Assign_Costing_Allocation_Sub_Process,omitempty"`
	ManageUnionMembershipSubProcess             *ManageUnionMembershipSubBusinessProcessType              `xml:"urn:com.workday/bsvc Manage_Union_Membership_Sub_Process,omitempty"`
}

func (t *EditPositionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EditPositionDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EditPositionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EditPositionDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for the Edit Position Event.
type EditPositionEventDataType struct {
	PositionChangeReasonReference *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Position_Change_Reason_Reference,omitempty"`
	FilledPositionEditDetailsData *FilledPositionEditDetailsDataType        `xml:"urn:com.workday/bsvc Filled_Position_Edit_Details_Data"`
}

// Responds with the Event ID for the Edit Position Event.
type EditPositionEventResponseType struct {
	EventReference         *UniqueIdentifierObjectType                    `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	ExceptionsResponseData []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This web service operation is designed to edit a filled Position (Position, Position Group, Job) using the Edit Position Business Process.
type EditPositionRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EditPositionData          *EditPositionDataType          `xml:"urn:com.workday/bsvc Edit_Position_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Wrapper element for the Edit Service Dates business process web service.
type EditServiceDatesBusinessProcessDataType struct {
	WorkerReference           *WorkerObjectType              `xml:"urn:com.workday/bsvc Worker_Reference"`
	EditServiceDatesEventData *EditServiceDatesEventDataType `xml:"urn:com.workday/bsvc Edit_Service_Dates_Event_Data"`
}

// Wrapper element for the Edit Service Dates details.
type EditServiceDatesEventDataType struct {
	OriginalHireDate                *time.Time `xml:"urn:com.workday/bsvc Original_Hire_Date,omitempty"`
	ContinuousServiceDate           *time.Time `xml:"urn:com.workday/bsvc Continuous_Service_Date,omitempty"`
	ExpectedRetirementDate          *time.Time `xml:"urn:com.workday/bsvc Expected_Retirement_Date,omitempty"`
	RetirementEligibilityDate       *time.Time `xml:"urn:com.workday/bsvc Retirement_Eligibility_Date,omitempty"`
	EndEmploymentDate               *time.Time `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
	SeniorityDate                   *time.Time `xml:"urn:com.workday/bsvc Seniority_Date,omitempty"`
	SeveranceDate                   *time.Time `xml:"urn:com.workday/bsvc Severance_Date,omitempty"`
	ContractEndDate                 *time.Time `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
	BenefitsServiceDate             *time.Time `xml:"urn:com.workday/bsvc Benefits_Service_Date,omitempty"`
	CompanyServiceDate              *time.Time `xml:"urn:com.workday/bsvc Company_Service_Date,omitempty"`
	TimeOffServiceDate              *time.Time `xml:"urn:com.workday/bsvc Time_Off_Service_Date,omitempty"`
	VestingDate                     *time.Time `xml:"urn:com.workday/bsvc Vesting_Date,omitempty"`
	DateEnteredWorkforce            *time.Time `xml:"urn:com.workday/bsvc Date_Entered_Workforce,omitempty"`
	DaysUnemployed                  float64    `xml:"urn:com.workday/bsvc Days_Unemployed,omitempty"`
	MonthsContinuousPriorEmployment float64    `xml:"urn:com.workday/bsvc Months_Continuous_Prior_Employment,omitempty"`
}

func (t *EditServiceDatesEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EditServiceDatesEventDataType
	var layout struct {
		*T
		OriginalHireDate          *xsdDate `xml:"urn:com.workday/bsvc Original_Hire_Date,omitempty"`
		ContinuousServiceDate     *xsdDate `xml:"urn:com.workday/bsvc Continuous_Service_Date,omitempty"`
		ExpectedRetirementDate    *xsdDate `xml:"urn:com.workday/bsvc Expected_Retirement_Date,omitempty"`
		RetirementEligibilityDate *xsdDate `xml:"urn:com.workday/bsvc Retirement_Eligibility_Date,omitempty"`
		EndEmploymentDate         *xsdDate `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
		SeniorityDate             *xsdDate `xml:"urn:com.workday/bsvc Seniority_Date,omitempty"`
		SeveranceDate             *xsdDate `xml:"urn:com.workday/bsvc Severance_Date,omitempty"`
		ContractEndDate           *xsdDate `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
		BenefitsServiceDate       *xsdDate `xml:"urn:com.workday/bsvc Benefits_Service_Date,omitempty"`
		CompanyServiceDate        *xsdDate `xml:"urn:com.workday/bsvc Company_Service_Date,omitempty"`
		TimeOffServiceDate        *xsdDate `xml:"urn:com.workday/bsvc Time_Off_Service_Date,omitempty"`
		VestingDate               *xsdDate `xml:"urn:com.workday/bsvc Vesting_Date,omitempty"`
		DateEnteredWorkforce      *xsdDate `xml:"urn:com.workday/bsvc Date_Entered_Workforce,omitempty"`
	}
	layout.T = (*T)(t)
	layout.OriginalHireDate = (*xsdDate)(layout.T.OriginalHireDate)
	layout.ContinuousServiceDate = (*xsdDate)(layout.T.ContinuousServiceDate)
	layout.ExpectedRetirementDate = (*xsdDate)(layout.T.ExpectedRetirementDate)
	layout.RetirementEligibilityDate = (*xsdDate)(layout.T.RetirementEligibilityDate)
	layout.EndEmploymentDate = (*xsdDate)(layout.T.EndEmploymentDate)
	layout.SeniorityDate = (*xsdDate)(layout.T.SeniorityDate)
	layout.SeveranceDate = (*xsdDate)(layout.T.SeveranceDate)
	layout.ContractEndDate = (*xsdDate)(layout.T.ContractEndDate)
	layout.BenefitsServiceDate = (*xsdDate)(layout.T.BenefitsServiceDate)
	layout.CompanyServiceDate = (*xsdDate)(layout.T.CompanyServiceDate)
	layout.TimeOffServiceDate = (*xsdDate)(layout.T.TimeOffServiceDate)
	layout.VestingDate = (*xsdDate)(layout.T.VestingDate)
	layout.DateEnteredWorkforce = (*xsdDate)(layout.T.DateEnteredWorkforce)
	return e.EncodeElement(layout, start)
}
func (t *EditServiceDatesEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EditServiceDatesEventDataType
	var overlay struct {
		*T
		OriginalHireDate          *xsdDate `xml:"urn:com.workday/bsvc Original_Hire_Date,omitempty"`
		ContinuousServiceDate     *xsdDate `xml:"urn:com.workday/bsvc Continuous_Service_Date,omitempty"`
		ExpectedRetirementDate    *xsdDate `xml:"urn:com.workday/bsvc Expected_Retirement_Date,omitempty"`
		RetirementEligibilityDate *xsdDate `xml:"urn:com.workday/bsvc Retirement_Eligibility_Date,omitempty"`
		EndEmploymentDate         *xsdDate `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
		SeniorityDate             *xsdDate `xml:"urn:com.workday/bsvc Seniority_Date,omitempty"`
		SeveranceDate             *xsdDate `xml:"urn:com.workday/bsvc Severance_Date,omitempty"`
		ContractEndDate           *xsdDate `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
		BenefitsServiceDate       *xsdDate `xml:"urn:com.workday/bsvc Benefits_Service_Date,omitempty"`
		CompanyServiceDate        *xsdDate `xml:"urn:com.workday/bsvc Company_Service_Date,omitempty"`
		TimeOffServiceDate        *xsdDate `xml:"urn:com.workday/bsvc Time_Off_Service_Date,omitempty"`
		VestingDate               *xsdDate `xml:"urn:com.workday/bsvc Vesting_Date,omitempty"`
		DateEnteredWorkforce      *xsdDate `xml:"urn:com.workday/bsvc Date_Entered_Workforce,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.OriginalHireDate = (*xsdDate)(overlay.T.OriginalHireDate)
	overlay.ContinuousServiceDate = (*xsdDate)(overlay.T.ContinuousServiceDate)
	overlay.ExpectedRetirementDate = (*xsdDate)(overlay.T.ExpectedRetirementDate)
	overlay.RetirementEligibilityDate = (*xsdDate)(overlay.T.RetirementEligibilityDate)
	overlay.EndEmploymentDate = (*xsdDate)(overlay.T.EndEmploymentDate)
	overlay.SeniorityDate = (*xsdDate)(overlay.T.SeniorityDate)
	overlay.SeveranceDate = (*xsdDate)(overlay.T.SeveranceDate)
	overlay.ContractEndDate = (*xsdDate)(overlay.T.ContractEndDate)
	overlay.BenefitsServiceDate = (*xsdDate)(overlay.T.BenefitsServiceDate)
	overlay.CompanyServiceDate = (*xsdDate)(overlay.T.CompanyServiceDate)
	overlay.TimeOffServiceDate = (*xsdDate)(overlay.T.TimeOffServiceDate)
	overlay.VestingDate = (*xsdDate)(overlay.T.VestingDate)
	overlay.DateEnteredWorkforce = (*xsdDate)(overlay.T.DateEnteredWorkforce)
	return d.DecodeElement(&overlay, &start)
}

// This web service operation allows updates to the service dates for a worker.
type EditServiceDatesRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType           `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EditServiceDatesData      *EditServiceDatesBusinessProcessDataType `xml:"urn:com.workday/bsvc Edit_Service_Dates_Data"`
	Version                   string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Event ID for the Service Dates Change Event.
type EditServiceDatesResponseType struct {
	EditServiceDatesEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Edit_Service_Dates_Event_Reference,omitempty"`
	Version                        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This web service operation allows updates to the service dates for a worker.
type EditServiceDatesSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	EditServiceDatesEventData    *EditServiceDatesEventDataType    `xml:"urn:com.workday/bsvc Edit_Service_Dates_Event_Data,omitempty"`
}

// Allows updating of effective dated custom objects for a worker.  When updating custom objects that allow multiple instances, data for existing instances must appear in the request, otherwise it will be removed.
type EditWorkerAdditionalDataRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	WorkerCustomObjectData    *WorkerCustomObjectDataType    `xml:"urn:com.workday/bsvc Worker_Custom_Object_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Woker effective-dated custom object data
type EditWorkerAdditionalDataResponseType struct {
	EventReference  *UniqueIdentifierObjectType                  `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	WorkerReference *WorkerObjectType                            `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	EffectiveDate   *time.Time                                   `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	AdditionalData  []EffectiveDatedWebServiceAdditionalDataType `xml:"urn:com.workday/bsvc Additional_Data,omitempty"`
	Version         string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

func (t *EditWorkerAdditionalDataResponseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EditWorkerAdditionalDataResponseType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EditWorkerAdditionalDataResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EditWorkerAdditionalDataResponseType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
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

// Contains a unique identifier for an instance of an object.
type EducationalTaxonomyCodeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EducationalTaxonomyCodeObjectType struct {
	ID         []EducationalTaxonomyCodeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Eligibility Criteria Data.  Eligibility Criteria element is used in conjunction with Workday delivered integrations.
type EligibilityCriteriaDataType struct {
	FieldReference    *IntegrationDocumentFieldObjectType `xml:"urn:com.workday/bsvc Field_Reference"`
	AsOfEffectiveDate *time.Time                          `xml:"urn:com.workday/bsvc As_Of_Effective_Date,omitempty"`
	AsOfEntryDateTime *time.Time                          `xml:"urn:com.workday/bsvc As_Of_Entry_DateTime,omitempty"`
}

func (t *EligibilityCriteriaDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EligibilityCriteriaDataType
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
func (t *EligibilityCriteriaDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EligibilityCriteriaDataType
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

// Contains the details about an emergency contact.
type EmergencyContactDataType struct {
	EmergencyContactID                string                              `xml:"urn:com.workday/bsvc Emergency_Contact_ID,omitempty"`
	LanguageReference                 []LanguageObjectType                `xml:"urn:com.workday/bsvc Language_Reference,omitempty"`
	EmergencyContactPriorityReference *EmergencyContactPriorityObjectType `xml:"urn:com.workday/bsvc Emergency_Contact_Priority_Reference,omitempty"`
	Primary                           bool                                `xml:"urn:com.workday/bsvc Primary,attr,omitempty"`
	Priority                          float64                             `xml:"urn:com.workday/bsvc Priority,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EmergencyContactObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmergencyContactObjectType struct {
	ID         []EmergencyContactObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element to hold the name and contact information of the emergency contact
type EmergencyContactPersonalInformationDataType struct {
	PersonNameData         []PersonNameDataType         `xml:"urn:com.workday/bsvc Person_Name_Data,omitempty"`
	ContactInformationData []ContactInformationDataType `xml:"urn:com.workday/bsvc Contact_Information_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EmergencyContactPriorityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmergencyContactPriorityObjectType struct {
	ID         []EmergencyContactPriorityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element containing all Allowance Plan Compensation data.
type EmployeeAllowancePlanAssignmentDataType struct {
	CompensationPlanReference    *AllowanceValuePlanObjectType     `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
	CompensationElementReference *CompensationPayEarningObjectType `xml:"urn:com.workday/bsvc Compensation_Element_Reference,omitempty"`
	Percent                      float64                           `xml:"urn:com.workday/bsvc Percent,omitempty"`
	Amount                       float64                           `xml:"urn:com.workday/bsvc Amount,omitempty"`
	CurrencyReference            *CurrencyObjectType               `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference           *FrequencyObjectType              `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	AssignmentEffectiveDate      *time.Time                        `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
}

func (t *EmployeeAllowancePlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeAllowancePlanAssignmentDataType
	var layout struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssignmentEffectiveDate = (*xsdDate)(layout.T.AssignmentEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeAllowancePlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeAllowancePlanAssignmentDataType
	var overlay struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssignmentEffectiveDate = (*xsdDate)(overlay.T.AssignmentEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Unit Allowance Plan Compensation data.
type EmployeeAllowanceUnitPlanAssignmentDataType struct {
	CompensationPlanReference    *AllowanceUnitPlanObjectType      `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
	CompensationElementReference *CompensationPayEarningObjectType `xml:"urn:com.workday/bsvc Compensation_Element_Reference,omitempty"`
	NumberofUnits                float64                           `xml:"urn:com.workday/bsvc Number_of_Units,omitempty"`
	UnitReference                *UnitofMeasureObjectType          `xml:"urn:com.workday/bsvc Unit_Reference,omitempty"`
	FrequencyReference           *FrequencyObjectType              `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	PerUnitAmount                float64                           `xml:"urn:com.workday/bsvc Per_Unit_Amount,omitempty"`
	CurrencyReference            *CurrencyObjectType               `xml:"urn:com.workday/bsvc Currency_Reference"`
	AssignmentEffectiveDate      *time.Time                        `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
}

func (t *EmployeeAllowanceUnitPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeAllowanceUnitPlanAssignmentDataType
	var layout struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssignmentEffectiveDate = (*xsdDate)(layout.T.AssignmentEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeAllowanceUnitPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeAllowanceUnitPlanAssignmentDataType
	var overlay struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssignmentEffectiveDate = (*xsdDate)(overlay.T.AssignmentEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Salary or Hourly Plan Compensation data.
type EmployeeBasePayPlanAssignmentDataType struct {
	CompensationPlanReference    *SalaryPayPlanObjectType          `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
	CompensationElementReference *CompensationPayEarningObjectType `xml:"urn:com.workday/bsvc Compensation_Element_Reference,omitempty"`
	Amount                       float64                           `xml:"urn:com.workday/bsvc Amount,omitempty"`
	CurrencyReference            *CurrencyObjectType               `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference           *FrequencyObjectType              `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	AssignmentEffectiveDate      *time.Time                        `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
}

func (t *EmployeeBasePayPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeBasePayPlanAssignmentDataType
	var layout struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssignmentEffectiveDate = (*xsdDate)(layout.T.AssignmentEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeBasePayPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeBasePayPlanAssignmentDataType
	var overlay struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssignmentEffectiveDate = (*xsdDate)(overlay.T.AssignmentEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Bonus Plan Compensation data.
type EmployeeBonusPlanAssignmentDataType struct {
	CompensationPlanReference    *BonusPlanObjectType              `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
	CompensationElementReference *CompensationPayEarningObjectType `xml:"urn:com.workday/bsvc Compensation_Element_Reference,omitempty"`
	IndividualTargetAmount       float64                           `xml:"urn:com.workday/bsvc Individual_Target_Amount,omitempty"`
	DefaultTargetAmount          float64                           `xml:"urn:com.workday/bsvc Default_Target_Amount,omitempty"`
	IndividualTargetPercent      float64                           `xml:"urn:com.workday/bsvc Individual_Target_Percent,omitempty"`
	DefaultTargetPercent         float64                           `xml:"urn:com.workday/bsvc Default_Target_Percent,omitempty"`
	GuaranteedMinimum            *bool                             `xml:"urn:com.workday/bsvc Guaranteed_Minimum,omitempty"`
	CurrencyReference            *CurrencyObjectType               `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference           *FrequencyObjectType              `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	PercentAssigned              float64                           `xml:"urn:com.workday/bsvc Percent_Assigned,omitempty"`
	AssignmentEffectiveDate      *time.Time                        `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
}

func (t *EmployeeBonusPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeBonusPlanAssignmentDataType
	var layout struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssignmentEffectiveDate = (*xsdDate)(layout.T.AssignmentEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeBonusPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeBonusPlanAssignmentDataType
	var overlay struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssignmentEffectiveDate = (*xsdDate)(overlay.T.AssignmentEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Commission Plan Compensation data.
type EmployeeCommissionPlanAssignmentDataType struct {
	CompensationPlanReference    *CommissionPlanObjectType         `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
	CompensationElementReference *CompensationPayEarningObjectType `xml:"urn:com.workday/bsvc Compensation_Element_Reference,omitempty"`
	TargetAmount                 float64                           `xml:"urn:com.workday/bsvc Target_Amount,omitempty"`
	CurrencyReference            *CurrencyObjectType               `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference           *FrequencyObjectType              `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	DrawAmount                   float64                           `xml:"urn:com.workday/bsvc Draw_Amount,omitempty"`
	DrawFrequencyReference       *FrequencyObjectType              `xml:"urn:com.workday/bsvc Draw_Frequency_Reference,omitempty"`
	DrawDuration                 string                            `xml:"urn:com.workday/bsvc Draw_Duration,omitempty"`
	Recoverable                  *bool                             `xml:"urn:com.workday/bsvc Recoverable,omitempty"`
	AssignmentEffectiveDate      *time.Time                        `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
}

func (t *EmployeeCommissionPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeCommissionPlanAssignmentDataType
	var layout struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssignmentEffectiveDate = (*xsdDate)(layout.T.AssignmentEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeCommissionPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeCommissionPlanAssignmentDataType
	var overlay struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssignmentEffectiveDate = (*xsdDate)(overlay.T.AssignmentEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Effective dated additional data for a employee contract.  When updating custom objects that allow multiple instances, data for existing instances must appear in the request, otherwise it will be removed.
type EmployeeContractCustomObjectDataType struct {
	EffectiveDate                time.Time                                   `xml:"urn:com.workday/bsvc Effective_Date"`
	EmployeeContractReference    *EmployeeContractObjectType                 `xml:"urn:com.workday/bsvc Employee_Contract_Reference"`
	BusinessObjectAdditionalData *EffectiveDatedWebServiceAdditionalDataType `xml:"urn:com.workday/bsvc Business_Object_Additional_Data"`
}

func (t *EmployeeContractCustomObjectDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeContractCustomObjectDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeContractCustomObjectDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeContractCustomObjectDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains data for an Employee Contract.
type EmployeeContractDataType struct {
	EmployeeContractReference                    *EmployeeContractObjectType        `xml:"urn:com.workday/bsvc Employee_Contract_Reference,omitempty"`
	EmployeeContractReasonReference              []EmployeeContractReasonObjectType `xml:"urn:com.workday/bsvc Employee_Contract_Reason_Reference,omitempty"`
	EmployeeContractID                           string                             `xml:"urn:com.workday/bsvc Employee_Contract_ID,omitempty"`
	EffectiveDate                                time.Time                          `xml:"urn:com.workday/bsvc Effective_Date"`
	ContractID                                   string                             `xml:"urn:com.workday/bsvc Contract_ID,omitempty"`
	ContractTypeReference                        *EmployeeContractTypeObjectType    `xml:"urn:com.workday/bsvc Contract_Type_Reference,omitempty"`
	ContractStartDate                            time.Time                          `xml:"urn:com.workday/bsvc Contract_Start_Date"`
	ContractEndDate                              *time.Time                         `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
	EmployeeContractCollectiveAgreementReference []CollectiveAgreementObjectType    `xml:"urn:com.workday/bsvc Employee_Contract_Collective_Agreement_Reference,omitempty"`
	MaximumWeeklyHours                           float64                            `xml:"urn:com.workday/bsvc Maximum_Weekly_Hours,omitempty"`
	MinimumWeeklyHours                           float64                            `xml:"urn:com.workday/bsvc Minimum_Weekly_Hours,omitempty"`
	ContractStatusReference                      *EmployeeContractStatusObjectType  `xml:"urn:com.workday/bsvc Contract_Status_Reference"`
	PositionReference                            *PositionElementObjectType         `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	ContractDescription                          string                             `xml:"urn:com.workday/bsvc Contract_Description,omitempty"`
	DateEmployeeSigned                           *time.Time                         `xml:"urn:com.workday/bsvc Date_Employee_Signed,omitempty"`
	DateEmployerSigned                           *time.Time                         `xml:"urn:com.workday/bsvc Date_Employer_Signed,omitempty"`
	WorkerDocumentReference                      []WorkerDocumentObjectType         `xml:"urn:com.workday/bsvc Worker_Document_Reference,omitempty"`
}

func (t *EmployeeContractDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeContractDataType
	var layout struct {
		*T
		EffectiveDate      *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
		ContractStartDate  *xsdDate `xml:"urn:com.workday/bsvc Contract_Start_Date"`
		ContractEndDate    *xsdDate `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
		DateEmployeeSigned *xsdDate `xml:"urn:com.workday/bsvc Date_Employee_Signed,omitempty"`
		DateEmployerSigned *xsdDate `xml:"urn:com.workday/bsvc Date_Employer_Signed,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	layout.ContractStartDate = (*xsdDate)(&layout.T.ContractStartDate)
	layout.ContractEndDate = (*xsdDate)(layout.T.ContractEndDate)
	layout.DateEmployeeSigned = (*xsdDate)(layout.T.DateEmployeeSigned)
	layout.DateEmployerSigned = (*xsdDate)(layout.T.DateEmployerSigned)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeContractDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeContractDataType
	var overlay struct {
		*T
		EffectiveDate      *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
		ContractStartDate  *xsdDate `xml:"urn:com.workday/bsvc Contract_Start_Date"`
		ContractEndDate    *xsdDate `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
		DateEmployeeSigned *xsdDate `xml:"urn:com.workday/bsvc Date_Employee_Signed,omitempty"`
		DateEmployerSigned *xsdDate `xml:"urn:com.workday/bsvc Date_Employer_Signed,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	overlay.ContractStartDate = (*xsdDate)(&overlay.T.ContractStartDate)
	overlay.ContractEndDate = (*xsdDate)(overlay.T.ContractEndDate)
	overlay.DateEmployeeSigned = (*xsdDate)(overlay.T.DateEmployeeSigned)
	overlay.DateEmployerSigned = (*xsdDate)(overlay.T.DateEmployerSigned)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type EmployeeContractObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmployeeContractObjectType struct {
	ID         []EmployeeContractObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EmployeeContractReasonObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmployeeContractReasonObjectType struct {
	ID         []EmployeeContractReasonObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EmployeeContractStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmployeeContractStatusObjectType struct {
	ID         []EmployeeContractStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EmployeeContractTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmployeeContractTypeObjectType struct {
	ID         []EmployeeContractTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the employee's contract information.
type EmployeeContractsDataType struct {
	EmployeeContractData []EmployeeContractDataType `xml:"urn:com.workday/bsvc Employee_Contract_Data,omitempty"`
}

// Contains the retirement savings amount information.
type EmployeeContributionAmountDataType struct {
	ContributionAmountData                 *ContributionAmountDataType                 `xml:"urn:com.workday/bsvc Contribution_Amount_Data"`
	PayrollInterfaceContributionAmountData *PayrollInterfaceContributionAmountDataType `xml:"urn:com.workday/bsvc Payroll_Interface_Contribution_Amount_Data,omitempty"`
	ContributionAmountMaximum              float64                                     `xml:"urn:com.workday/bsvc Contribution_Amount_Maximum"`
	CurrencyReference                      *CurrencyObjectType                         `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
}

// Contains the spending account election information including the contribution, payroll interface contribution, goal, and annual contribution information.
type EmployeeContributionDataType struct {
	ContributionData                 *ContributionDataType                 `xml:"urn:com.workday/bsvc Contribution_Data,omitempty"`
	PayrollInterfaceContributionData *PayrollInterfaceContributionDataType `xml:"urn:com.workday/bsvc Payroll_Interface_Contribution_Data,omitempty"`
	GoalData                         *GoalDataType                         `xml:"urn:com.workday/bsvc Goal_Data,omitempty"`
	AnnualContributionData           *AnnualContributionDataType           `xml:"urn:com.workday/bsvc Annual_Contribution_Data,omitempty"`
	CurrencyReference                *CurrencyObjectType                   `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
}

// This section holds the employee's election and the maximum the employee can elect when the benefit plan uses percentages as elections.
type EmployeeContributionPercentageDataType struct {
	ElectionPercentage            float64 `xml:"urn:com.workday/bsvc Election_Percentage"`
	ContributionPercentageMaximum float64 `xml:"urn:com.workday/bsvc Contribution_Percentage_Maximum"`
}

// Encapsulating element containing all Employee Image data.
type EmployeeImageDataType struct {
	Filename string  `xml:"urn:com.workday/bsvc Filename"`
	Image    *[]byte `xml:"urn:com.workday/bsvc Image,omitempty"`
}

func (t *EmployeeImageDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeImageDataType
	var layout struct {
		*T
		Image *xsdBase64Binary `xml:"urn:com.workday/bsvc Image,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Image = (*xsdBase64Binary)(layout.T.Image)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeImageDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeImageDataType
	var overlay struct {
		*T
		Image *xsdBase64Binary `xml:"urn:com.workday/bsvc Image,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Image = (*xsdBase64Binary)(overlay.T.Image)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Merit Plan Compensation data.
type EmployeeMeritPlanAssignmentDataType struct {
	CompensationPlanReference    *MeritPercentPlanObjectType   `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
	IndividualTargetPercent      float64                       `xml:"urn:com.workday/bsvc Individual_Target_Percent,omitempty"`
	DefaultTargetPercent         float64                       `xml:"urn:com.workday/bsvc Default_Target_Percent,omitempty"`
	MeritIncreaseMatrixReference *CompensationMatrixObjectType `xml:"urn:com.workday/bsvc Merit_Increase_Matrix_Reference,omitempty"`
	GuaranteedMinimum            *bool                         `xml:"urn:com.workday/bsvc Guaranteed_Minimum,omitempty"`
	AssignmentEffectiveDate      *time.Time                    `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
}

func (t *EmployeeMeritPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeMeritPlanAssignmentDataType
	var layout struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssignmentEffectiveDate = (*xsdDate)(layout.T.AssignmentEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeMeritPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeMeritPlanAssignmentDataType
	var overlay struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssignmentEffectiveDate = (*xsdDate)(overlay.T.AssignmentEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Adds or updates notice period data for an employee's voluntary termination.
type EmployeeNoticePeriodDataType struct {
	DeriveNoticePeriod       bool                          `xml:"urn:com.workday/bsvc Derive_Notice_Period"`
	OverrideNoticePeriodData *OverrideNoticePeriodDataType `xml:"urn:com.workday/bsvc Override_Notice_Period_Data"`
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

// Encapsulating element containing all Period Salary Plan Compensation data.
type EmployeePeriodSalaryPlanAssignmentDataType struct {
	CompensationPlanReference    *PeriodSalaryPlanObjectType       `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
	CompensationElementReference *CompensationPayEarningObjectType `xml:"urn:com.workday/bsvc Compensation_Element_Reference,omitempty"`
	CompensationPeriodReference  *CompensationPeriodObjectType     `xml:"urn:com.workday/bsvc Compensation_Period_Reference,omitempty"`
	CurrencyReference            *CurrencyObjectType               `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	CompensationPeriodMultiplier float64                           `xml:"urn:com.workday/bsvc Compensation_Period_Multiplier,omitempty"`
	FrequencyReference           *FrequencyObjectType              `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	AssignmentEffectiveDate      *time.Time                        `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
}

func (t *EmployeePeriodSalaryPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeePeriodSalaryPlanAssignmentDataType
	var layout struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssignmentEffectiveDate = (*xsdDate)(layout.T.AssignmentEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeePeriodSalaryPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeePeriodSalaryPlanAssignmentDataType
	var overlay struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssignmentEffectiveDate = (*xsdDate)(overlay.T.AssignmentEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Probation Period data. Including Corrected data.
type EmployeeProbationPeriodDetailDataType struct {
	EmployeeProbationPeriodEventReference *UniqueIdentifierObjectType              `xml:"urn:com.workday/bsvc Employee_Probation_Period_Event_Reference,omitempty"`
	EmployeeProbationPeriodReference      *EmployeeProbationPeriodObjectType       `xml:"urn:com.workday/bsvc Employee_Probation_Period_Reference,omitempty"`
	EffectiveDate                         *time.Time                               `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	StartDate                             *time.Time                               `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                               *time.Time                               `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	ProbationTypeReference                *EmployeeProbationPeriodTypeObjectType   `xml:"urn:com.workday/bsvc Probation_Type_Reference,omitempty"`
	ProbationReasonReference              *EmployeeProbationPeriodReasonObjectType `xml:"urn:com.workday/bsvc Probation_Reason_Reference,omitempty"`
	ExtendedEndDate                       *time.Time                               `xml:"urn:com.workday/bsvc Extended_End_Date,omitempty"`
	Note                                  string                                   `xml:"urn:com.workday/bsvc Note,omitempty"`
}

func (t *EmployeeProbationPeriodDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeProbationPeriodDetailDataType
	var layout struct {
		*T
		EffectiveDate   *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		StartDate       *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate         *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
		ExtendedEndDate *xsdDate `xml:"urn:com.workday/bsvc Extended_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	layout.StartDate = (*xsdDate)(layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	layout.ExtendedEndDate = (*xsdDate)(layout.T.ExtendedEndDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeProbationPeriodDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeProbationPeriodDetailDataType
	var overlay struct {
		*T
		EffectiveDate   *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		StartDate       *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate         *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
		ExtendedEndDate *xsdDate `xml:"urn:com.workday/bsvc Extended_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	overlay.StartDate = (*xsdDate)(overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	overlay.ExtendedEndDate = (*xsdDate)(overlay.T.ExtendedEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type EmployeeProbationPeriodObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmployeeProbationPeriodObjectType struct {
	ID         []EmployeeProbationPeriodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EmployeeProbationPeriodReasonObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmployeeProbationPeriodReasonObjectType struct {
	ID         []EmployeeProbationPeriodReasonObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element containing all Probation Period data.
type EmployeeProbationPeriodSummaryDataType struct {
	EmployeeProbationPeriodDetailData []EmployeeProbationPeriodDetailDataType `xml:"urn:com.workday/bsvc Employee_Probation_Period_Detail_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EmployeeProbationPeriodTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmployeeProbationPeriodTypeObjectType struct {
	ID         []EmployeeProbationPeriodTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the latest completed performance review for the employee.
type EmployeeReviewDataType struct {
	PerformanceReviewData          *PerformanceReviewDataType          `xml:"urn:com.workday/bsvc Performance_Review_Data,omitempty"`
	PerformanceImprovementPlanData *PerformanceImprovementPlanDataType `xml:"urn:com.workday/bsvc Performance_Improvement_Plan_Data,omitempty"`
	DevelopmentPlanData            *DevelopmentPlanDataType            `xml:"urn:com.workday/bsvc Development_Plan_Data,omitempty"`
	DisciplinaryActionData         *DisciplinaryActionDataType         `xml:"urn:com.workday/bsvc Disciplinary_Action_Data,omitempty"`
}

// Contains the information about the review for the employee.
type EmployeeReviewDetailsDataType struct {
	ManagerReference         *WorkerObjectType                 `xml:"urn:com.workday/bsvc Manager_Reference,omitempty"`
	MultipleManagerReference []WorkerObjectType                `xml:"urn:com.workday/bsvc Multiple_Manager_Reference,omitempty"`
	ReviewTypeReference      *ReviewTypeObjectType             `xml:"urn:com.workday/bsvc Review_Type_Reference,omitempty"`
	ReviewTemplateReference  *EmployeeReviewTemplateObjectType `xml:"urn:com.workday/bsvc Review_Template_Reference,omitempty"`
	ReviewInitiatedDate      *time.Time                        `xml:"urn:com.workday/bsvc Review_Initiated_Date,omitempty"`
	PeriodStartDate          *time.Time                        `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
	PeriodEndDate            *time.Time                        `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	SelfEvaluationData       *SelfEvaluationDetailDataType     `xml:"urn:com.workday/bsvc Self_Evaluation_Data,omitempty"`
	ManagerEvaluationData    *ManagerEvaluationDetailDataType  `xml:"urn:com.workday/bsvc Manager_Evaluation_Data,omitempty"`
}

func (t *EmployeeReviewDetailsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeReviewDetailsDataType
	var layout struct {
		*T
		ReviewInitiatedDate *xsdDateTime `xml:"urn:com.workday/bsvc Review_Initiated_Date,omitempty"`
		PeriodStartDate     *xsdDate     `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
		PeriodEndDate       *xsdDate     `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ReviewInitiatedDate = (*xsdDateTime)(layout.T.ReviewInitiatedDate)
	layout.PeriodStartDate = (*xsdDate)(layout.T.PeriodStartDate)
	layout.PeriodEndDate = (*xsdDate)(layout.T.PeriodEndDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeReviewDetailsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeReviewDetailsDataType
	var overlay struct {
		*T
		ReviewInitiatedDate *xsdDateTime `xml:"urn:com.workday/bsvc Review_Initiated_Date,omitempty"`
		PeriodStartDate     *xsdDate     `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
		PeriodEndDate       *xsdDate     `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ReviewInitiatedDate = (*xsdDateTime)(overlay.T.ReviewInitiatedDate)
	overlay.PeriodStartDate = (*xsdDate)(overlay.T.PeriodStartDate)
	overlay.PeriodEndDate = (*xsdDate)(overlay.T.PeriodEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type EmployeeReviewTemplateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmployeeReviewTemplateObjectType struct {
	ID         []EmployeeReviewTemplateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element containing all Unit Salary Plan Compensation data.
type EmployeeSalaryUnitPlanAssignmentDataType struct {
	CompensationPlanReference    *SalaryUnitPlanObjectType         `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
	CompensationElementReference *CompensationPayEarningObjectType `xml:"urn:com.workday/bsvc Compensation_Element_Reference,omitempty"`
	UnitReference                *UnitofMeasureObjectType          `xml:"urn:com.workday/bsvc Unit_Reference,omitempty"`
	PerUnitAmount                float64                           `xml:"urn:com.workday/bsvc Per_Unit_Amount,omitempty"`
	CurrencyReference            *CurrencyObjectType               `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	NumberofUnits                float64                           `xml:"urn:com.workday/bsvc Number_of_Units,omitempty"`
	FrequencyReference           *FrequencyObjectType              `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	AssignmentEffectiveDate      *time.Time                        `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
}

func (t *EmployeeSalaryUnitPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeSalaryUnitPlanAssignmentDataType
	var layout struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssignmentEffectiveDate = (*xsdDate)(layout.T.AssignmentEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeSalaryUnitPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeSalaryUnitPlanAssignmentDataType
	var overlay struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssignmentEffectiveDate = (*xsdDate)(overlay.T.AssignmentEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Stock Plan Compensation data.
type EmployeeStockPlanAssignmentDataType struct {
	CompensationPlanReference   *StockPlanObjectType            `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
	PercentofTargetasOption     float64                         `xml:"urn:com.workday/bsvc Percent_of_Target_as_Option,omitempty"`
	PercentofTargetasStock      float64                         `xml:"urn:com.workday/bsvc Percent_of_Target_as_Stock,omitempty"`
	TargetShares                float64                         `xml:"urn:com.workday/bsvc Target_Shares,omitempty"`
	IndividualTargetShares      float64                         `xml:"urn:com.workday/bsvc Individual_Target_Shares,omitempty"`
	TargetPercent               float64                         `xml:"urn:com.workday/bsvc Target_Percent,omitempty"`
	IndividualTargetAmount      float64                         `xml:"urn:com.workday/bsvc Individual_Target_Amount,omitempty"`
	TargetAmount                float64                         `xml:"urn:com.workday/bsvc Target_Amount,omitempty"`
	IndividualTargetPercent     float64                         `xml:"urn:com.workday/bsvc Individual_Target_Percent,omitempty"`
	CurrencyReference           *CurrencyObjectType             `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	VestingScheduleReference    *StockVestingScheduleObjectType `xml:"urn:com.workday/bsvc Vesting_Schedule_Reference,omitempty"`
	CompensationMatrixReference *CompensationMatrixObjectType   `xml:"urn:com.workday/bsvc Compensation_Matrix_Reference,omitempty"`
	AssignmentEffectiveDate     *time.Time                      `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
}

func (t *EmployeeStockPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeStockPlanAssignmentDataType
	var layout struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssignmentEffectiveDate = (*xsdDate)(layout.T.AssignmentEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeStockPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeStockPlanAssignmentDataType
	var overlay struct {
		*T
		AssignmentEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssignmentEffectiveDate = (*xsdDate)(overlay.T.AssignmentEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains the information about the employee's potential.
type EmployeeTalentAssessmentDataType struct {
	PotentialReference       *PotentialObjectType       `xml:"urn:com.workday/bsvc Potential_Reference,omitempty"`
	AchievableLevelReference *AchievableLevelObjectType `xml:"urn:com.workday/bsvc Achievable_Level_Reference,omitempty"`
	RetentionRiskReference   *RetentionObjectType       `xml:"urn:com.workday/bsvc Retention_Risk_Reference,omitempty"`
	LossImpactRiskReference  *LossImpactObjectType      `xml:"urn:com.workday/bsvc Loss_Impact_Risk_Reference,omitempty"`
	Notes                    string                     `xml:"urn:com.workday/bsvc Notes,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EmployeeTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmployeeTypeObjectType struct {
	ID         []EmployeeTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the employer contribution amount for the retirement savings election.
type EmployerContributionAmountDataType struct {
	ContributionAmountData                 *ContributionAmountDataType                 `xml:"urn:com.workday/bsvc Contribution_Amount_Data"`
	PayrollInterfaceContributionAmountData *PayrollInterfaceContributionAmountDataType `xml:"urn:com.workday/bsvc Payroll_Interface_Contribution_Amount_Data,omitempty"`
	CurrencyReference                      *CurrencyObjectType                         `xml:"urn:com.workday/bsvc Currency_Reference"`
}

// Contains the employer contribution amount information for the spending account election.
type EmployerContributionDataType struct {
	ContributionData                 *ContributionDataType                 `xml:"urn:com.workday/bsvc Contribution_Data,omitempty"`
	PayrollInterfaceContributionData *PayrollInterfaceContributionDataType `xml:"urn:com.workday/bsvc Payroll_Interface_Contribution_Data,omitempty"`
	GoalData                         *GoalDataType                         `xml:"urn:com.workday/bsvc Goal_Data,omitempty"`
	CurrencyReference                *CurrencyObjectType                   `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
}

// This section holds the employer contribution if the plan includes employer contributions and the contributions are stated in percentages.
type EmployerContributionPercentageDataType struct {
	ElectionPercentage float64 `xml:"urn:com.workday/bsvc Election_Percentage,omitempty"`
}

// Adds or updates an employer's notice period data for an employee's involuntary termination.
type EmployerNoticePeriodDataType struct {
	DeriveNoticePeriod       bool                          `xml:"urn:com.workday/bsvc Derive_Notice_Period"`
	OverrideNoticePeriodData *OverrideNoticePeriodDataType `xml:"urn:com.workday/bsvc Override_Notice_Period_Data"`
}

// Wrapper element for the End Academic Appointment business process.
type EndAcademicAppointmentDataType struct {
	AcademicAppointeeReference        *AcademicAppointeeEnabledObjectType `xml:"urn:com.workday/bsvc Academic_Appointee_Reference,omitempty"`
	AcademicAppointmentTrackReference *AcademicAppointmentTrackObjectType `xml:"urn:com.workday/bsvc Academic_Appointment_Track_Reference,omitempty"`
	ReasonReference                   *GeneralEventSubcategoryObjectType  `xml:"urn:com.workday/bsvc Reason_Reference"`
	TrackEndDate                      time.Time                           `xml:"urn:com.workday/bsvc Track_End_Date"`
}

func (t *EndAcademicAppointmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EndAcademicAppointmentDataType
	var layout struct {
		*T
		TrackEndDate *xsdDate `xml:"urn:com.workday/bsvc Track_End_Date"`
	}
	layout.T = (*T)(t)
	layout.TrackEndDate = (*xsdDate)(&layout.T.TrackEndDate)
	return e.EncodeElement(layout, start)
}
func (t *EndAcademicAppointmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EndAcademicAppointmentDataType
	var overlay struct {
		*T
		TrackEndDate *xsdDate `xml:"urn:com.workday/bsvc Track_End_Date"`
	}
	overlay.T = (*T)(t)
	overlay.TrackEndDate = (*xsdDate)(&overlay.T.TrackEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper for the End Academic Appointment sub business process that is part of the Terminate Employee business processes.
type EndAcademicAppointmentSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	EndAcademicAppointmentData   *EndAcademicAppointmentDataType   `xml:"urn:com.workday/bsvc End_Academic_Appointment_Data,omitempty"`
}

// Responds with the Event ID for the End Additional Job Event.
type EndAdditionalEmployeeJobEventResponseType struct {
	EventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This web service is used to End Additional Job for an Employee. It will run the End Additional Job Business Process.
type EndAdditionalJobDataType struct {
	EmployeeReference                 *EmployeeObjectType                            `xml:"urn:com.workday/bsvc Employee_Reference"`
	PositionReference                 *PositionElementObjectType                     `xml:"urn:com.workday/bsvc Position_Reference"`
	EndAdditionalJobDate              time.Time                                      `xml:"urn:com.workday/bsvc End_Additional_Job_Date"`
	EndAdditionalJobEventData         *EndAdditionalJobEventDataType                 `xml:"urn:com.workday/bsvc End_Additional_Job_Event_Data"`
	AssignOrganizationRolesSubProcess *AssignOrganizationRolesSubBusinessProcessType `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Sub_Process,omitempty"`
	CreateJobRequisitionSubProcess    *CreateJobRequisitionSubProcessType            `xml:"urn:com.workday/bsvc Create_Job_Requisition_Sub_Process,omitempty"`
	CheckPositionBudgetSubProcess     *CheckPositionBudgetSubBusinessProcessType     `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process,omitempty"`
	EndAcademicAppointmentSubProcess  *EndAcademicAppointmentSubBusinessProcessType  `xml:"urn:com.workday/bsvc End_Academic_Appointment_Sub_Process,omitempty"`
	ManageUnionMembershipSubProcess   *ManageUnionMembershipSubBusinessProcessType   `xml:"urn:com.workday/bsvc Manage_Union_Membership_Sub_Process,omitempty"`
}

func (t *EndAdditionalJobDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EndAdditionalJobDataType
	var layout struct {
		*T
		EndAdditionalJobDate *xsdDate `xml:"urn:com.workday/bsvc End_Additional_Job_Date"`
	}
	layout.T = (*T)(t)
	layout.EndAdditionalJobDate = (*xsdDate)(&layout.T.EndAdditionalJobDate)
	return e.EncodeElement(layout, start)
}
func (t *EndAdditionalJobDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EndAdditionalJobDataType
	var overlay struct {
		*T
		EndAdditionalJobDate *xsdDate `xml:"urn:com.workday/bsvc End_Additional_Job_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EndAdditionalJobDate = (*xsdDate)(&overlay.T.EndAdditionalJobDate)
	return d.DecodeElement(&overlay, &start)
}

// Data Element that contains the basic information to end an additional job for an employee. All required information must be included. The End Additional Job business process (for the employee) will be initiated from this information.
type EndAdditionalJobEventDataType struct {
	LastDayofWork        time.Time                                 `xml:"urn:com.workday/bsvc Last_Day_of_Work"`
	ReasonReference      *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference"`
	PayThroughDate       *time.Time                                `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
	NotifyEmployeeByDate *time.Time                                `xml:"urn:com.workday/bsvc Notify_Employee_By_Date,omitempty"`
	ClosePosition        *bool                                     `xml:"urn:com.workday/bsvc Close_Position,omitempty"`
	JobOverlapAllowed    *bool                                     `xml:"urn:com.workday/bsvc Job_Overlap_Allowed,omitempty"`
}

func (t *EndAdditionalJobEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EndAdditionalJobEventDataType
	var layout struct {
		*T
		LastDayofWork        *xsdDate `xml:"urn:com.workday/bsvc Last_Day_of_Work"`
		PayThroughDate       *xsdDate `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
		NotifyEmployeeByDate *xsdDate `xml:"urn:com.workday/bsvc Notify_Employee_By_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LastDayofWork = (*xsdDate)(&layout.T.LastDayofWork)
	layout.PayThroughDate = (*xsdDate)(layout.T.PayThroughDate)
	layout.NotifyEmployeeByDate = (*xsdDate)(layout.T.NotifyEmployeeByDate)
	return e.EncodeElement(layout, start)
}
func (t *EndAdditionalJobEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EndAdditionalJobEventDataType
	var overlay struct {
		*T
		LastDayofWork        *xsdDate `xml:"urn:com.workday/bsvc Last_Day_of_Work"`
		PayThroughDate       *xsdDate `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
		NotifyEmployeeByDate *xsdDate `xml:"urn:com.workday/bsvc Notify_Employee_By_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LastDayofWork = (*xsdDate)(&overlay.T.LastDayofWork)
	overlay.PayThroughDate = (*xsdDate)(overlay.T.PayThroughDate)
	overlay.NotifyEmployeeByDate = (*xsdDate)(overlay.T.NotifyEmployeeByDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper for End Additional Job Web Service.
type EndAdditionalJobRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EndAdditionalJobData      *EndAdditionalJobDataType      `xml:"urn:com.workday/bsvc End_Additional_Job_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This web service is used to end a contract for a contingent worker. This web service will run the End Contingent Worker Contract business process and any of its sub-processes.
type EndContingentWorkerContractDataType struct {
	ContingentWorkerReference         *ContingentWorkerObjectType                    `xml:"urn:com.workday/bsvc Contingent_Worker_Reference"`
	ContractEndDate                   time.Time                                      `xml:"urn:com.workday/bsvc Contract_End_Date"`
	EndContractEventData              *EndContingentWorkerContractEventDataType      `xml:"urn:com.workday/bsvc End_Contract_Event_Data"`
	AssignOrganizationRolesSubProcess *AssignOrganizationRolesSubBusinessProcessType `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Sub_Process,omitempty"`
	CreateJobRequisitionSubProcess    []CreateJobRequisitionSubProcessType           `xml:"urn:com.workday/bsvc Create_Job_Requisition_Sub_Process,omitempty"`
}

func (t *EndContingentWorkerContractDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EndContingentWorkerContractDataType
	var layout struct {
		*T
		ContractEndDate *xsdDate `xml:"urn:com.workday/bsvc Contract_End_Date"`
	}
	layout.T = (*T)(t)
	layout.ContractEndDate = (*xsdDate)(&layout.T.ContractEndDate)
	return e.EncodeElement(layout, start)
}
func (t *EndContingentWorkerContractDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EndContingentWorkerContractDataType
	var overlay struct {
		*T
		ContractEndDate *xsdDate `xml:"urn:com.workday/bsvc Contract_End_Date"`
	}
	overlay.T = (*T)(t)
	overlay.ContractEndDate = (*xsdDate)(&overlay.T.ContractEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Data Element that contains the basic information to end the contract of a contingent worker. All required information must be included. The End Contingent Worker Contract business process will be initiated from this information.
type EndContingentWorkerContractEventDataType struct {
	LastDayofWork                   *time.Time                                `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
	PrimaryReasonReference          *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Primary_Reason_Reference"`
	SecondaryReasonReference        []TerminationSubcategoryObjectType        `xml:"urn:com.workday/bsvc Secondary_Reason_Reference,omitempty"`
	LocalTerminationReasonReference *LocalTerminationReasonObjectType         `xml:"urn:com.workday/bsvc Local_Termination_Reason_Reference,omitempty"`
	NotifyWorkerByDate              *time.Time                                `xml:"urn:com.workday/bsvc Notify_Worker_By_Date,omitempty"`
	RegrettableReference            *CommonYesNoObjectType                    `xml:"urn:com.workday/bsvc Regrettable_Reference,omitempty"`
	ClosePosition                   *bool                                     `xml:"urn:com.workday/bsvc Close_Position,omitempty"`
	JobOverlapAllowed               *bool                                     `xml:"urn:com.workday/bsvc Job_Overlap_Allowed,omitempty"`
}

func (t *EndContingentWorkerContractEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EndContingentWorkerContractEventDataType
	var layout struct {
		*T
		LastDayofWork      *xsdDate `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
		NotifyWorkerByDate *xsdDate `xml:"urn:com.workday/bsvc Notify_Worker_By_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LastDayofWork = (*xsdDate)(layout.T.LastDayofWork)
	layout.NotifyWorkerByDate = (*xsdDate)(layout.T.NotifyWorkerByDate)
	return e.EncodeElement(layout, start)
}
func (t *EndContingentWorkerContractEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EndContingentWorkerContractEventDataType
	var overlay struct {
		*T
		LastDayofWork      *xsdDate `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
		NotifyWorkerByDate *xsdDate `xml:"urn:com.workday/bsvc Notify_Worker_By_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LastDayofWork = (*xsdDate)(overlay.T.LastDayofWork)
	overlay.NotifyWorkerByDate = (*xsdDate)(overlay.T.NotifyWorkerByDate)
	return d.DecodeElement(&overlay, &start)
}

// This web service operation is designed to end an existing Contingent Worker Contract
type EndContingentWorkerContractRequestHVType struct {
	BusinessProcessParameters       *BusinessProcessParametersType       `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EndContingentWorkerContractData *EndContingentWorkerContractDataType `xml:"urn:com.workday/bsvc End_Contingent_Worker_Contract_Data"`
}

// Wrapper for End Contingent Worker Contract web service and its sub-processes.
type EndContingentWorkerContractRequestType struct {
	BusinessProcessParameters       *BusinessProcessParametersType       `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EndContingentWorkerContractData *EndContingentWorkerContractDataType `xml:"urn:com.workday/bsvc End_Contingent_Worker_Contract_Data"`
	Version                         string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Reponds with the event id for the End Contract Event.
type EndContingentWorkerContractResponseType struct {
	EventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This web service is used to End International Assignment for an Employee. It will run the End International Assignment Business Process.
type EndInternationalAssignmentDataType struct {
	EmployeeReference                   *EmployeeObjectType                        `xml:"urn:com.workday/bsvc Employee_Reference"`
	PositionReference                   *PositionElementObjectType                 `xml:"urn:com.workday/bsvc Position_Reference"`
	EndInternationalAssignmentDate      time.Time                                  `xml:"urn:com.workday/bsvc End_International_Assignment_Date"`
	EndInternationalAssignmentEventData *EndInternationalAssignmentEventDataType   `xml:"urn:com.workday/bsvc End_International_Assignment_Event_Data"`
	CreateJobRequisitionSubProcess      *CreateJobRequisitionSubProcessType        `xml:"urn:com.workday/bsvc Create_Job_Requisition_Sub_Process,omitempty"`
	CheckPositionBudgetSubProcess       *CheckPositionBudgetSubBusinessProcessType `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process,omitempty"`
}

func (t *EndInternationalAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EndInternationalAssignmentDataType
	var layout struct {
		*T
		EndInternationalAssignmentDate *xsdDate `xml:"urn:com.workday/bsvc End_International_Assignment_Date"`
	}
	layout.T = (*T)(t)
	layout.EndInternationalAssignmentDate = (*xsdDate)(&layout.T.EndInternationalAssignmentDate)
	return e.EncodeElement(layout, start)
}
func (t *EndInternationalAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EndInternationalAssignmentDataType
	var overlay struct {
		*T
		EndInternationalAssignmentDate *xsdDate `xml:"urn:com.workday/bsvc End_International_Assignment_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EndInternationalAssignmentDate = (*xsdDate)(&overlay.T.EndInternationalAssignmentDate)
	return d.DecodeElement(&overlay, &start)
}

// Data Element that contains the basic information to end an international assignment job for an employee. All required information must be included. The End International Assignment business process (for the employee) will be initiated from this information.
type EndInternationalAssignmentEventDataType struct {
	LastDayofWork        time.Time                                 `xml:"urn:com.workday/bsvc Last_Day_of_Work"`
	ReasonReference      *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference"`
	PayThroughDate       *time.Time                                `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
	NotifyEmployeeByDate *time.Time                                `xml:"urn:com.workday/bsvc Notify_Employee_By_Date,omitempty"`
	ClosePosition        *bool                                     `xml:"urn:com.workday/bsvc Close_Position,omitempty"`
	JobOverlapAllowed    *bool                                     `xml:"urn:com.workday/bsvc Job_Overlap_Allowed,omitempty"`
}

func (t *EndInternationalAssignmentEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EndInternationalAssignmentEventDataType
	var layout struct {
		*T
		LastDayofWork        *xsdDate `xml:"urn:com.workday/bsvc Last_Day_of_Work"`
		PayThroughDate       *xsdDate `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
		NotifyEmployeeByDate *xsdDate `xml:"urn:com.workday/bsvc Notify_Employee_By_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LastDayofWork = (*xsdDate)(&layout.T.LastDayofWork)
	layout.PayThroughDate = (*xsdDate)(layout.T.PayThroughDate)
	layout.NotifyEmployeeByDate = (*xsdDate)(layout.T.NotifyEmployeeByDate)
	return e.EncodeElement(layout, start)
}
func (t *EndInternationalAssignmentEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EndInternationalAssignmentEventDataType
	var overlay struct {
		*T
		LastDayofWork        *xsdDate `xml:"urn:com.workday/bsvc Last_Day_of_Work"`
		PayThroughDate       *xsdDate `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
		NotifyEmployeeByDate *xsdDate `xml:"urn:com.workday/bsvc Notify_Employee_By_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LastDayofWork = (*xsdDate)(&overlay.T.LastDayofWork)
	overlay.PayThroughDate = (*xsdDate)(overlay.T.PayThroughDate)
	overlay.NotifyEmployeeByDate = (*xsdDate)(overlay.T.NotifyEmployeeByDate)
	return d.DecodeElement(&overlay, &start)
}

// This web service operation is designed to end an existing international assignment
type EndInternationalAssignmentforEmployeeEventRequestType struct {
	BusinessProcessParameters      *BusinessProcessParametersType      `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EndInternationalAssignmentData *EndInternationalAssignmentDataType `xml:"urn:com.workday/bsvc End_International_Assignment_Data"`
	Version                        string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with a reference to the End International Assignment Employee Event
type EndInternationalAssignmentforEmployeeEventResponseType struct {
	EndInternationalAssignmentforEmployeeEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc End_International_Assignment_for_Employee_Event_Reference,omitempty"`
	Version                                             string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the benefit plan year information.
type EnrollmentPeriodDataType struct {
	BenefitProgramName string     `xml:"urn:com.workday/bsvc Benefit_Program_Name"`
	PlanYear           *time.Time `xml:"urn:com.workday/bsvc Plan_Year,omitempty"`
	BeginDate          time.Time  `xml:"urn:com.workday/bsvc Begin_Date"`
	EndDate            time.Time  `xml:"urn:com.workday/bsvc End_Date"`
}

func (t *EnrollmentPeriodDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EnrollmentPeriodDataType
	var layout struct {
		*T
		PlanYear  *xsdDate `xml:"urn:com.workday/bsvc Plan_Year,omitempty"`
		BeginDate *xsdDate `xml:"urn:com.workday/bsvc Begin_Date"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date"`
	}
	layout.T = (*T)(t)
	layout.PlanYear = (*xsdDate)(layout.T.PlanYear)
	layout.BeginDate = (*xsdDate)(&layout.T.BeginDate)
	layout.EndDate = (*xsdDate)(&layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *EnrollmentPeriodDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EnrollmentPeriodDataType
	var overlay struct {
		*T
		PlanYear  *xsdDate `xml:"urn:com.workday/bsvc Plan_Year,omitempty"`
		BeginDate *xsdDate `xml:"urn:com.workday/bsvc Begin_Date"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date"`
	}
	overlay.T = (*T)(t)
	overlay.PlanYear = (*xsdDate)(overlay.T.PlanYear)
	overlay.BeginDate = (*xsdDate)(&overlay.T.BeginDate)
	overlay.EndDate = (*xsdDate)(&overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
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
type EventClassificationCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EventClassificationCategoryObjectType struct {
	ID         []EventClassificationCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Encapsulating element containing all transaction log entries.
type EventTargetTransactionLogEntryDataType struct {
	TransactionLogEntry []TransactionLogEntryType `xml:"urn:com.workday/bsvc Transaction_Log_Entry,omitempty"`
}

// Encapsulating element containing transaction log entries that have been rescinded or corrected.
type EventTargetTransactionLogRescindAndCorrectDataType struct {
	CorrectedOrRescindedTransactionLogData []CorrectedOrRescindedTransactionDataType `xml:"urn:com.workday/bsvc Corrected_Or_Rescinded_Transaction_Log_Data,omitempty"`
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

// Exception (Errors and Warning) associated with the transaction.
type ExceptionDataType struct {
	Classification string `xml:"urn:com.workday/bsvc Classification,omitempty"`
	Message        string `xml:"urn:com.workday/bsvc Message,omitempty"`
}

// Contains extended data for an employee contract.
type ExtendedEmployeeContractDataType struct {
	EmployeeContractData []EmployeeContractDataType `xml:"urn:com.workday/bsvc Employee_Contract_Data,omitempty"`
}

// Contains the extended employee's contract information.
type ExtendedEmployeeContractsDataType struct {
	ExtendedEmployeeContractData []ExtendedEmployeeContractDataType `xml:"urn:com.workday/bsvc Extended_Employee_Contract_Data,omitempty"`
}

// Element containing External Student Record data
type ExternalAcademicRecordDataType struct {
	StudentActive                    *bool                                 `xml:"urn:com.workday/bsvc Student_Active,omitempty"`
	DoNotReleaseDirectoryInformation *bool                                 `xml:"urn:com.workday/bsvc Do_Not_Release_Directory_Information,omitempty"`
	AcademicUnitReference            *AcademicCurricularDivisionObjectType `xml:"urn:com.workday/bsvc Academic_Unit_Reference"`
	AcademicLevelReference           *AcademicLevelObjectType              `xml:"urn:com.workday/bsvc Academic_Level_Reference,omitempty"`
	ExpectedGraduationDate           *time.Time                            `xml:"urn:com.workday/bsvc Expected_Graduation_Date,omitempty"`
	WorkStudyEligible                *bool                                 `xml:"urn:com.workday/bsvc Work-Study_Eligible,omitempty"`
	GPA                              float64                               `xml:"urn:com.workday/bsvc GPA,omitempty"`
	AcademicStandingReference        *StudentAcademicStandingObjectType    `xml:"urn:com.workday/bsvc Academic_Standing_Reference,omitempty"`
	StudentLoadStatusReference       *StudentLoadStatusObjectType          `xml:"urn:com.workday/bsvc Student_Load_Status_Reference,omitempty"`
	EnrolledUnits                    float64                               `xml:"urn:com.workday/bsvc Enrolled_Units,omitempty"`
	EnrolledUnitTypeReference        *OtherUnitTypeObjectType              `xml:"urn:com.workday/bsvc Enrolled_Unit_Type_Reference,omitempty"`
	ClassStandingReference           *ClassStandingDefinitionObjectType    `xml:"urn:com.workday/bsvc Class_Standing_Reference,omitempty"`
	ExternalStudentStudentData       *ExternalStudentStudentDataType       `xml:"urn:com.workday/bsvc External_Student_Student_Data"`
}

func (t *ExternalAcademicRecordDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ExternalAcademicRecordDataType
	var layout struct {
		*T
		ExpectedGraduationDate *xsdDate `xml:"urn:com.workday/bsvc Expected_Graduation_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ExpectedGraduationDate = (*xsdDate)(layout.T.ExpectedGraduationDate)
	return e.EncodeElement(layout, start)
}
func (t *ExternalAcademicRecordDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ExternalAcademicRecordDataType
	var overlay struct {
		*T
		ExpectedGraduationDate *xsdDate `xml:"urn:com.workday/bsvc Expected_Graduation_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ExpectedGraduationDate = (*xsdDate)(overlay.T.ExpectedGraduationDate)
	return d.DecodeElement(&overlay, &start)
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
type ExternalFieldandParameterInitializationProviderObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ExternalFieldandParameterInitializationProviderObjectType struct {
	ID         []ExternalFieldandParameterInitializationProviderObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Integration ID Help Text
type ExternalIDDataType struct {
	ExternalID []ExternalIDType `xml:"urn:com.workday/bsvc External_ID,omitempty"`
}

// External ID that uniquely identifies the integratable object within the context of the integration system identified by the System ID attribute.
type ExternalIDType struct {
	ExternalID string `xml:"urn:com.workday/bsvc External_ID"`
	SystemID   string `xml:"urn:com.workday/bsvc System_ID,attr,omitempty"`
}

// Integration ID Help Text
type ExternalIntegrationIDDataType struct {
	ID []IDType `xml:"urn:com.workday/bsvc ID"`
}

// Contains a unique identifier for an instance of an object.
type ExternalPayGroupObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type ExternalPayGroupObjectType struct {
	ID         []ExternalPayGroupObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ExternalPayrollEmployeeTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ExternalPayrollEmployeeTypeObjectType struct {
	ID         []ExternalPayrollEmployeeTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ExternalPayrollEntityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ExternalPayrollEntityObjectType struct {
	ID         []ExternalPayrollEntityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing the Person data for the Student.
type ExternalStudentPersonDataType struct {
	PersonNameData           *PersonNameDataType                                          `xml:"urn:com.workday/bsvc Person_Name_Data,omitempty"`
	ContactInformationData   *ContactInformationDataType                                  `xml:"urn:com.workday/bsvc Contact_Information_Data,omitempty"`
	PersonalProfileData      *StudentPersonalProfileDataType                              `xml:"urn:com.workday/bsvc Personal_Profile_Data,omitempty"`
	PersonIdentificationData *PersonIdentificationDataType                                `xml:"urn:com.workday/bsvc Person_Identification_Data,omitempty"`
	PersonalInformationData  *StudentApplicationDetailsPersonalInformationSubeditDataType `xml:"urn:com.workday/bsvc Personal_Information_Data,omitempty"`
}

// Element containing information about the Student to whom this External Student Record belongs
type ExternalStudentStudentDataType struct {
	StudentReference                  *StudentObjectType                     `xml:"urn:com.workday/bsvc Student_Reference,omitempty"`
	ExternalStudentStudentSubeditData *ExternalStudentStudentSubeditDataType `xml:"urn:com.workday/bsvc External_Student_Student_Subedit_Data,omitempty"`
}

// Element containing Student data
type ExternalStudentStudentSubeditDataType struct {
	StudentID         string                         `xml:"urn:com.workday/bsvc Student_ID"`
	WorkerReference   *WorkerObjectType              `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	PreHireReference  *ApplicantObjectType           `xml:"urn:com.workday/bsvc Pre_Hire_Reference,omitempty"`
	StudentPersonData *ExternalStudentPersonDataType `xml:"urn:com.workday/bsvc Student_Person_Data,omitempty"`
	StudentPhotoData  *StudentPhotoDataType          `xml:"urn:com.workday/bsvc Student_Photo_Data,omitempty"`
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

// Contains the employee's feedback received
type FeedbackReceivedDataType struct {
	FeedbackReceived []FeedbackReceivedType `xml:"urn:com.workday/bsvc Feedback_Received,omitempty"`
}

// Contains the informational components of a feedback (comment, date, question asked, sender, requester, and type).
type FeedbackReceivedType struct {
	From                       string                 `xml:"urn:com.workday/bsvc From,omitempty"`
	RequestedByWorkerReference *WorkerObjectType      `xml:"urn:com.workday/bsvc Requested_By_Worker_Reference,omitempty"`
	FeedbackType               string                 `xml:"urn:com.workday/bsvc Feedback_Type,omitempty"`
	Date                       *time.Time             `xml:"urn:com.workday/bsvc Date,omitempty"`
	FeedbackResponseData       []FeedbackResponseType `xml:"urn:com.workday/bsvc Feedback_Response_Data,omitempty"`
}

func (t *FeedbackReceivedType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T FeedbackReceivedType
	var layout struct {
		*T
		Date *xsdDate `xml:"urn:com.workday/bsvc Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Date = (*xsdDate)(layout.T.Date)
	return e.EncodeElement(layout, start)
}
func (t *FeedbackReceivedType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T FeedbackReceivedType
	var overlay struct {
		*T
		Date *xsdDate `xml:"urn:com.workday/bsvc Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Date = (*xsdDate)(overlay.T.Date)
	return d.DecodeElement(&overlay, &start)
}

// Element for Feedback Comments and Questions.  There can be multiple questions per request, thus multiple comments per feedback given.
type FeedbackResponseType struct {
	FeedbackQuestion string `xml:"urn:com.workday/bsvc Feedback_Question,omitempty"`
	FeedbackComment  string `xml:"urn:com.workday/bsvc Feedback_Comment,omitempty"`
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

// Contains information for a filled position.
// If this Position has already been filled, this element can be used to change this data. If the position has not been entered, an error will be returned if this element is supplied.
type FilledPositionEditDetailsDataType struct {
	PositionID                           string                                 `xml:"urn:com.workday/bsvc Position_ID,omitempty"`
	EmployeeTypeReference                *EmployeeTypeObjectType                `xml:"urn:com.workday/bsvc Employee_Type_Reference,omitempty"`
	ContingentWorkerTypeReference        *ContingentWorkerTypeObjectType        `xml:"urn:com.workday/bsvc Contingent_Worker_Type_Reference,omitempty"`
	PositionDetailsSubData               *PositionDetailsSubDataType            `xml:"urn:com.workday/bsvc Position_Details_Sub_Data,omitempty"`
	EndEmploymentDate                    *time.Time                             `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
	ContractEndDate                      *time.Time                             `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
	ContractDetailsData                  *ContractDetailsSubDataType            `xml:"urn:com.workday/bsvc Contract_Details_Data,omitempty"`
	ExcludefromHeadcount                 *bool                                  `xml:"urn:com.workday/bsvc Exclude_from_Headcount,omitempty"`
	ExpectedAssignmentEndDate            *time.Time                             `xml:"urn:com.workday/bsvc Expected_Assignment_End_Date,omitempty"`
	InternationalAssignmentTypeReference *InternationalAssignmentTypeObjectType `xml:"urn:com.workday/bsvc International_Assignment_Type_Reference,omitempty"`
}

func (t *FilledPositionEditDetailsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T FilledPositionEditDetailsDataType
	var layout struct {
		*T
		EndEmploymentDate         *xsdDate `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
		ContractEndDate           *xsdDate `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
		ExpectedAssignmentEndDate *xsdDate `xml:"urn:com.workday/bsvc Expected_Assignment_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EndEmploymentDate = (*xsdDate)(layout.T.EndEmploymentDate)
	layout.ContractEndDate = (*xsdDate)(layout.T.ContractEndDate)
	layout.ExpectedAssignmentEndDate = (*xsdDate)(layout.T.ExpectedAssignmentEndDate)
	return e.EncodeElement(layout, start)
}
func (t *FilledPositionEditDetailsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T FilledPositionEditDetailsDataType
	var overlay struct {
		*T
		EndEmploymentDate         *xsdDate `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
		ContractEndDate           *xsdDate `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
		ExpectedAssignmentEndDate *xsdDate `xml:"urn:com.workday/bsvc Expected_Assignment_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EndEmploymentDate = (*xsdDate)(overlay.T.EndEmploymentDate)
	overlay.ContractEndDate = (*xsdDate)(overlay.T.ContractEndDate)
	overlay.ExpectedAssignmentEndDate = (*xsdDate)(overlay.T.ExpectedAssignmentEndDate)
	return d.DecodeElement(&overlay, &start)
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

// Wrapper element for the Freeze Position business process.
type FreezePositionRequestType struct {
	FreezePositionData        *PositionGroupFreezeDataType   `xml:"urn:com.workday/bsvc Freeze_Position_Data"`
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the reference of the Position Restriction being frozen.
type FreezePositionResponseType struct {
	PositionGroupFreezeEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Position_Group_Freeze_Event_Reference,omitempty"`
	Version                           string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Future payment plan assignment data submitted by the put future payment plan assignment request.
type FuturePaymentPlanAssignmentDataType struct {
	CompensationPlanReference         *FuturePaymentPlanObjectType `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
	IndividualTargetAmount            float64                      `xml:"urn:com.workday/bsvc Individual_Target_Amount,omitempty"`
	IndividualTargetCurrencyReference *CurrencyObjectType          `xml:"urn:com.workday/bsvc Individual_Target_Currency_Reference,omitempty"`
	IndividualTargetPaymentDate       *time.Time                   `xml:"urn:com.workday/bsvc Individual_Target_Payment_Date,omitempty"`
	Comment                           string                       `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

func (t *FuturePaymentPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T FuturePaymentPlanAssignmentDataType
	var layout struct {
		*T
		IndividualTargetPaymentDate *xsdDate `xml:"urn:com.workday/bsvc Individual_Target_Payment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.IndividualTargetPaymentDate = (*xsdDate)(layout.T.IndividualTargetPaymentDate)
	return e.EncodeElement(layout, start)
}
func (t *FuturePaymentPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T FuturePaymentPlanAssignmentDataType
	var overlay struct {
		*T
		IndividualTargetPaymentDate *xsdDate `xml:"urn:com.workday/bsvc Individual_Target_Payment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.IndividualTargetPaymentDate = (*xsdDate)(overlay.T.IndividualTargetPaymentDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type FuturePaymentPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type FuturePaymentPlanObjectType struct {
	ID         []FuturePaymentPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type GeneralEventSubcategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type GeneralEventSubcategoryObjectType struct {
	ID         []GeneralEventSubcategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for all Disability Status data for the person.
type GenericDisabilityInformationDataType struct {
	DisabilityStatusInformationData []GenericDisabilityStatusInformationDataType `xml:"urn:com.workday/bsvc Disability_Status_Information_Data,omitempty"`
	ReplaceAll                      bool                                         `xml:"urn:com.workday/bsvc Replace_All,attr,omitempty"`
}

// Wrapper element for each disability status entry.
type GenericDisabilityStatusInformationDataType struct {
	DisabilityStatusReference *DisabilityStatusReferenceObjectType `xml:"urn:com.workday/bsvc Disability_Status_Reference,omitempty"`
	DisabilityStatusData      *GenericDisabilityStatusSubDataType  `xml:"urn:com.workday/bsvc Disability_Status_Data,omitempty"`
	Delete                    bool                                 `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Container for disability status data.
type GenericDisabilityStatusSubDataType struct {
	DisabilityReference                       *DisabilityObjectType                       `xml:"urn:com.workday/bsvc Disability_Reference"`
	DisabilityStatusDate                      *time.Time                                  `xml:"urn:com.workday/bsvc Disability_Status_Date,omitempty"`
	DisabilityDateKnown                       *time.Time                                  `xml:"urn:com.workday/bsvc Disability_Date_Known,omitempty"`
	DisabilityEndDate                         *time.Time                                  `xml:"urn:com.workday/bsvc Disability_End_Date,omitempty"`
	DisabilityGradeReference                  *DisabilityGradeObjectType                  `xml:"urn:com.workday/bsvc Disability_Grade_Reference,omitempty"`
	DisabilityDegree                          float64                                     `xml:"urn:com.workday/bsvc Disability_Degree,omitempty"`
	DisabilityRemainingCapacity               float64                                     `xml:"urn:com.workday/bsvc Disability_Remaining_Capacity,omitempty"`
	DisabilityCertificationAuthorityReference *DisabilityCertificationAuthorityObjectType `xml:"urn:com.workday/bsvc Disability_Certification_Authority_Reference,omitempty"`
	DisabilityAuthority                       string                                      `xml:"urn:com.workday/bsvc Disability_Authority,omitempty"`
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
}

func (t *GenericDisabilityStatusSubDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T GenericDisabilityStatusSubDataType
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
func (t *GenericDisabilityStatusSubDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T GenericDisabilityStatusSubDataType
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

// Wrapper element for all Military Service data for the person.
type GenericMilitaryInformationDataType struct {
	MilitaryServiceInformationData []GenericMilitaryServiceInformationDataType `xml:"urn:com.workday/bsvc Military_Service_Information_Data,omitempty"`
	ReplaceAll                     bool                                        `xml:"urn:com.workday/bsvc Replace_All,attr,omitempty"`
}

// Wrapper element for each Military Service entry.
type GenericMilitaryServiceInformationDataType struct {
	MilitaryServiceReference *MilitaryServiceReferenceObjectType `xml:"urn:com.workday/bsvc Military_Service_Reference,omitempty"`
	MilitaryServiceData      *GenericMilitaryServiceSubDataType  `xml:"urn:com.workday/bsvc Military_Service_Data,omitempty"`
	Delete                   bool                                `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Container for military service data for the Person
type GenericMilitaryServiceSubDataType struct {
	MilitaryStatusReference      *MilitaryStatusObjectType      `xml:"urn:com.workday/bsvc Military_Status_Reference"`
	MilitaryDischargeDate        *time.Time                     `xml:"urn:com.workday/bsvc Military_Discharge_Date,omitempty"`
	MilitaryStatusBeginDate      *time.Time                     `xml:"urn:com.workday/bsvc Military_Status_Begin_Date,omitempty"`
	MilitaryServiceTypeReference *MilitaryServiceTypeObjectType `xml:"urn:com.workday/bsvc Military_Service_Type_Reference,omitempty"`
	MilitaryRankReference        *MilitaryRankObjectType        `xml:"urn:com.workday/bsvc Military_Rank_Reference,omitempty"`
	Notes                        string                         `xml:"urn:com.workday/bsvc Notes,omitempty"`
}

func (t *GenericMilitaryServiceSubDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T GenericMilitaryServiceSubDataType
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
func (t *GenericMilitaryServiceSubDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T GenericMilitaryServiceSubDataType
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

// Container for the Personal Information data.
type GenericPersonalInformationDataType struct {
	GenderReference                 *GenderObjectType                      `xml:"urn:com.workday/bsvc Gender_Reference,omitempty"`
	DateofBirth                     *time.Time                             `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	DateofDeath                     *time.Time                             `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
	BirthCountryReference           *CountryObjectType                     `xml:"urn:com.workday/bsvc Birth_Country_Reference,omitempty"`
	BirthRegionReference            *CountryRegionObjectType               `xml:"urn:com.workday/bsvc Birth_Region_Reference,omitempty"`
	CityofBirth                     string                                 `xml:"urn:com.workday/bsvc City_of_Birth,omitempty"`
	MaritalStatusReference          *MaritalStatusObjectType               `xml:"urn:com.workday/bsvc Marital_Status_Reference,omitempty"`
	MaritalStatusDate               *time.Time                             `xml:"urn:com.workday/bsvc Marital_Status_Date,omitempty"`
	EthnicityReference              []EthnicityObjectType                  `xml:"urn:com.workday/bsvc Ethnicity_Reference,omitempty"`
	HispanicorLatino                *bool                                  `xml:"urn:com.workday/bsvc Hispanic_or_Latino,omitempty"`
	ReligionReference               *ReligionObjectType                    `xml:"urn:com.workday/bsvc Religion_Reference,omitempty"`
	CitizenshipReference            []CitizenshipStatusObjectType          `xml:"urn:com.workday/bsvc Citizenship_Reference,omitempty"`
	NationalityReference            *CountryObjectType                     `xml:"urn:com.workday/bsvc Nationality_Reference,omitempty"`
	HukouRegionReference            *CountryRegionObjectType               `xml:"urn:com.workday/bsvc Hukou_Region_Reference,omitempty"`
	HukouSubregionReference         *CountrySubregionObjectType            `xml:"urn:com.workday/bsvc Hukou_Subregion_Reference,omitempty"`
	HukouLocality                   string                                 `xml:"urn:com.workday/bsvc Hukou_Locality,omitempty"`
	HukouPostalCode                 string                                 `xml:"urn:com.workday/bsvc Hukou_Postal_Code,omitempty"`
	HukouTypeReference              *HukouTypeObjectType                   `xml:"urn:com.workday/bsvc Hukou_Type_Reference,omitempty"`
	NativeRegionReference           *CountryRegionObjectType               `xml:"urn:com.workday/bsvc Native_Region_Reference,omitempty"`
	PersonnelFileAgency             string                                 `xml:"urn:com.workday/bsvc Personnel_File_Agency,omitempty"`
	PoliticalAffiliationReference   *PoliticalAffiliationObjectType        `xml:"urn:com.workday/bsvc Political_Affiliation_Reference,omitempty"`
	SocialBenefitsLocalityReference *SocialBenefitsLocalityObjectType      `xml:"urn:com.workday/bsvc Social_Benefits_Locality_Reference,omitempty"`
	LastMedicalExamDate             *time.Time                             `xml:"urn:com.workday/bsvc Last_Medical_Exam_Date,omitempty"`
	LastMedicalExamValidTo          *time.Time                             `xml:"urn:com.workday/bsvc Last_Medical_Exam_Valid_To,omitempty"`
	MedicalExamNotes                string                                 `xml:"urn:com.workday/bsvc Medical_Exam_Notes,omitempty"`
	DisabilityInformationData       []GenericDisabilityInformationDataType `xml:"urn:com.workday/bsvc Disability_Information_Data,omitempty"`
	MilitaryInformationData         []GenericMilitaryInformationDataType   `xml:"urn:com.workday/bsvc Military_Information_Data,omitempty"`
	UsesTobacco                     *bool                                  `xml:"urn:com.workday/bsvc Uses_Tobacco,omitempty"`
}

func (t *GenericPersonalInformationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T GenericPersonalInformationDataType
	var layout struct {
		*T
		DateofBirth            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
		DateofDeath            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
		MaritalStatusDate      *xsdDate `xml:"urn:com.workday/bsvc Marital_Status_Date,omitempty"`
		LastMedicalExamDate    *xsdDate `xml:"urn:com.workday/bsvc Last_Medical_Exam_Date,omitempty"`
		LastMedicalExamValidTo *xsdDate `xml:"urn:com.workday/bsvc Last_Medical_Exam_Valid_To,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateofBirth = (*xsdDate)(layout.T.DateofBirth)
	layout.DateofDeath = (*xsdDate)(layout.T.DateofDeath)
	layout.MaritalStatusDate = (*xsdDate)(layout.T.MaritalStatusDate)
	layout.LastMedicalExamDate = (*xsdDate)(layout.T.LastMedicalExamDate)
	layout.LastMedicalExamValidTo = (*xsdDate)(layout.T.LastMedicalExamValidTo)
	return e.EncodeElement(layout, start)
}
func (t *GenericPersonalInformationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T GenericPersonalInformationDataType
	var overlay struct {
		*T
		DateofBirth            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
		DateofDeath            *xsdDate `xml:"urn:com.workday/bsvc Date_of_Death,omitempty"`
		MaritalStatusDate      *xsdDate `xml:"urn:com.workday/bsvc Marital_Status_Date,omitempty"`
		LastMedicalExamDate    *xsdDate `xml:"urn:com.workday/bsvc Last_Medical_Exam_Date,omitempty"`
		LastMedicalExamValidTo *xsdDate `xml:"urn:com.workday/bsvc Last_Medical_Exam_Valid_To,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateofBirth = (*xsdDate)(overlay.T.DateofBirth)
	overlay.DateofDeath = (*xsdDate)(overlay.T.DateofDeath)
	overlay.MaritalStatusDate = (*xsdDate)(overlay.T.MaritalStatusDate)
	overlay.LastMedicalExamDate = (*xsdDate)(overlay.T.LastMedicalExamDate)
	overlay.LastMedicalExamValidTo = (*xsdDate)(overlay.T.LastMedicalExamValidTo)
	return d.DecodeElement(&overlay, &start)
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

// Utilize the Request References element to retrieve a specific  instance or instances of Worker and its associated IDs.
type GetChangeGovernmentIDsRequestReferencesType struct {
	PersonReference      []RoleObjectType                `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	UniversalIDReference []UniversalIdentifierObjectType `xml:"urn:com.workday/bsvc Universal_ID_Reference,omitempty"`
}

// Request element used to get government IDs for a worker. The response is formatted to be usable as input to the Change Government IDs web service request.
type GetChangeGovernmentIDsRequestType struct {
	RequestReferences   *GetChangeGovernmentIDsRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter      *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	RequestCriteriaData *ChangeGovernmentIDRequestCriteriaType       `xml:"urn:com.workday/bsvc Request_Criteria_Data,omitempty"`
	Version             string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the Get Change Government IDs response elements including echoed request data and request references.
type GetChangeGovernmentIDsResponseType struct {
	RequestReferences []GetChangeGovernmentIDsRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    []ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   []ResponseResultsType                         `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []ChangeGovernmentIDsResponseDataType         `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Utilize the Request References element to retrieve a specific instance(s) of Worker and its associated data.
type GetChangeLegalNameRequestReferencesType struct {
	PersonReference      []RoleObjectType                `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	UniversalIDReference []UniversalIdentifierObjectType `xml:"urn:com.workday/bsvc Universal_ID_Reference,omitempty"`
}

// Request element used to get legal name for a worker. The response is formatted to be usable as input to the Change Legal Name web service request.
type GetChangeLegalNameRequestType struct {
	RequestReferences   *GetChangeLegalNameRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter      *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	RequestCriteriaData *ChangeLegalNameRequestCriteriaType      `xml:"urn:com.workday/bsvc Request_Criteria_Data,omitempty"`
	Version             string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the Get Change Legal Name response elements including echoed request data and request references.
type GetChangeLegalNameResponseType struct {
	RequestReferences *GetChangeLegalNameRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                     `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *ChangeLegalNameResponseDataType         `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Utilize the Request References element to retrieve a specific  instance or instances of Worker and its associated IDs.
type GetChangeLicensesRequestReferencesType struct {
	PersonReference      []RoleObjectType                `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	UniversalIDReference []UniversalIdentifierObjectType `xml:"urn:com.workday/bsvc Universal_ID_Reference,omitempty"`
}

// Request element used to get licenses for a worker. The response is formatted to be usable as input to the Change Licenses web service request.
type GetChangeLicensesRequestType struct {
	RequestReferences   *GetChangeLicensesRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter      *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	RequestCriteriaData *ChangeLicenseIDRequestCriteriaType     `xml:"urn:com.workday/bsvc Request_Criteria_Data,omitempty"`
	Version             string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing a worker's licenses.
type GetChangeLicensesResponseType struct {
	RequestReferences []GetChangeLicensesRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    []ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   []ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []ChangeLicensesResponseDataType         `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Utilize the Request References element to retrieve a specific  instance or instances of Worker and its associated IDs.
type GetChangeOtherIDsRequestReferencesType struct {
	PersonReference      []RoleObjectType                `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	UniversalIDReference []UniversalIdentifierObjectType `xml:"urn:com.workday/bsvc Universal_ID_Reference,omitempty"`
}

// Request element used to get other IDs for a worker. The response is formatted to be usable as input to the Change Other IDs web service request.
type GetChangeOtherIDsRequestType struct {
	RequestReferences   *GetChangeOtherIDsRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter      *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	RequestCriteriaData *ChangeOtherIDRequestCriteriaType       `xml:"urn:com.workday/bsvc Request_Criteria_Data,omitempty"`
	Version             string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing a worker's other IDs.
type GetChangeOtherIDsResponseType struct {
	RequestReferences []GetChangeOtherIDsRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    []ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   []ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []ChangeOtherIDsResponseDataType         `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Utilize the Request References element to retrieve a specific  instance or instances of Worker and its associated IDs.
type GetChangePassportsandVisasRequestReferencesType struct {
	PersonReference      []RoleObjectType                `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	UniversalIDReference []UniversalIdentifierObjectType `xml:"urn:com.workday/bsvc Universal_ID_Reference,omitempty"`
}

// Request element used to get passports and visas for a worker. The response is formatted to be usable as input to the Change Passports and Visas web service request.
type GetChangePassportsandVisasRequestType struct {
	RequestReferences   *GetChangePassportsandVisasRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter      *ResponseFilterType                              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	RequestCriteriaData *ChangePassportandVisaRequestCriteriaType        `xml:"urn:com.workday/bsvc Request_Criteria_Data,omitempty"`
	Version             string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing a worker's passports and visas.
type GetChangePassportsandVisasResponseType struct {
	RequestReferences []GetChangePassportsandVisasRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    []ResponseFilterType                              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   []ResponseResultsType                             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []ChangePassportsandVisasResponseDataType         `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Utilize the Request References element to retrieve a specific instance(s) of Worker and its associated data.
type GetChangePreferredNameRequestReferencesType struct {
	PersonReference      []RoleObjectType                `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	UniversalIDReference []UniversalIdentifierObjectType `xml:"urn:com.workday/bsvc Universal_ID_Reference,omitempty"`
}

// Request element used to get preferred name for a worker. The response is formatted to be usable as input to the Change Preferred Name web service request.
type GetChangePreferredNameRequestType struct {
	RequestReferences   *GetChangePreferredNameRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter      *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	RequestCriteriaData *ChangePreferredNameRequestCriteriaType      `xml:"urn:com.workday/bsvc Request_Criteria_Data,omitempty"`
	Version             string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the Get Change Preferred Name response elements including echoed request data and request references.
type GetChangePreferredNameResponseType struct {
	RequestReferences *GetChangePreferredNameRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                         `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *ChangePreferredNameResponseDataType         `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the references of the Citizenship Statuses requested.
type GetCitizenshipStatusesRequestType struct {
	RequestReferences *CitizenshipStatusRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response Element for Get Citizenship Statuses.
type GetCitizenshipStatusesResponseType struct {
	RequestReferences *CitizenshipStatusRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *CitizenshipStatusResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request Element for Get Job Classification
type GetJobClassificationGroupsRequestType struct {
	RequestReferences *JobClassificationGroupRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobClassificationGroupResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response Element for Get Job Classification Groups
type GetJobClassificationGroupsResponseType struct {
	RequestReferences *JobClassificationGroupRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobClassificationGroupResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                         `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *JobClassificationGroupResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for Get Job Families
type GetJobFamiliesRequestType struct {
	RequestReferences *JobFamilyRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *JobFamiliesRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobFamilyResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Get Job Families.
type GetJobFamiliesResponseType struct {
	RequestReferences *JobFamilyRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	RequestCriteria   *JobFamiliesRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseGroup     *JobFamilyResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType            `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *JobFamilyResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for Get Job Family Groups
type GetJobFamilyGroupsRequestType struct {
	RequestReferences *JobFamilyGroupRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobFamilyGroupResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for the Get Job Family Group request
type GetJobFamilyGroupsResponseType struct {
	RequestReferences *JobFamilyGroupRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *JobFamilyGroupResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *JobFamilyGroupResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Notice Periods for a single or multiple countries. Get all Notice Periods if Country is empty.
type GetMaintainNoticePeriodsForCountryRequestType struct {
	RequestReferences *NoticePeriodsForCountryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *NoticePeriodsForCountryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                           `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response contain all Notice Periods what met Request Criteria.
type GetMaintainNoticePeriodsForCountryResponseType struct {
	RequestReferences *NoticePeriodsForCountryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *NoticePeriodsForCountryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseData      *NoticePeriodsForCountryResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	ResponseFilter    *ResponseFilterType                           `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                          `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	Version           string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Probation Periods for a single or multiple countries
type GetMaintainProbationPeriodsForCountryRequestType struct {
	RequestReferences *ProbationPeriodsForCountryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *ProbationPeriodsForCountryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response contains all Probation Periods that met Request Criteria
type GetMaintainProbationPeriodsForCountryResponseType struct {
	RequestReferences *ProbationPeriodsForCountryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *ProbationPeriodsForCountryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseData      *ProbationPeriodsForCountryResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	ResponseFilter    *ResponseFilterType                              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	Version           string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Notice Period Eligibility Rules Request
type GetNoticePeriodEligibilityRulesRequestType struct {
	RequestReferences *NoticePeriodEligibilityRuleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Notice Period Eligibility Rules Response
type GetNoticePeriodEligibilityRulesResponseType struct {
	RequestReferences *NoticePeriodEligibilityRuleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                              `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *NoticePeriodEligibilityRuleResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a single reference to a Notice Period Target.
type GetNoticePeriodsRequestReferencesType struct {
	NoticePeriodTargetReference *NoticePeriodTargetObjectType `xml:"urn:com.workday/bsvc Notice_Period_Target_Reference"`
}

// Get Notice Periods for a single Notice Period Target (Employee). Returns employer and employee notice periods.
type GetNoticePeriodsRequestType struct {
	RequestReferences *GetNoticePeriodsRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the employer/employee Notice Period response data specific to the Notice Period Target request and As Of Effective Date.
type GetNoticePeriodsResponseDataType struct {
	NoticePeriodDataforNoticePeriodTarget []NoticePeriodDataforNoticePeriodTargetType `xml:"urn:com.workday/bsvc Notice_Period_Data_for_Notice_Period_Target,omitempty"`
}

// Contains the employer/employee Notice Period response data specific to the Notice Period Target request and As Of Effective Date.
type GetNoticePeriodsforNoticePeriodTargetResponseType struct {
	RequestReferences *GetNoticePeriodsRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *GetNoticePeriodsResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Get Probation Period Eligibility Rules Request Element
type GetProbationPeriodEligibilityRulesRequestType struct {
	RequestReferences *ProbationPeriodEligibilityRuleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Probation Period Eligibility Rules Response Element
type GetProbationPeriodEligibilityRulesResponseType struct {
	RequestReferences *ProbationPeriodEligibilityRuleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *ProbationPeriodEligibilityRuleResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for Get Probation Period Outcomes
type GetProbationPeriodOutcomesRequestType struct {
	RequestReferences *ProbationPeriodOutcomesForCountryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *ProbationPeriodOutcomesForCountryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response containing all Probation Period Outcomes which met the Request Criteria
type GetProbationPeriodOutcomesResponseType struct {
	RequestReferences *ProbationPeriodOutcomesForCountryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *ProbationPeriodOutcomesForCountryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseData      *ProbationPeriodOutcomesForCountryResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	ResponseFilter    *ResponseFilterType                                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	Version           string                                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains Probation Periods For Worker
type GetProbationPeriodsForWorkerResponseDataType struct {
	ProbationPeriodsForWorkers []ProbationPeriodsForWorkersType `xml:"urn:com.workday/bsvc Probation_Periods_For_Workers,omitempty"`
}

// Get Probation Periods for a single Worker
type GetProbationPeriodsForWorkersRequestType struct {
	RequestReferences *ProbationPeriodsForWorkersRequestReferencesType `xml:"urn:com.workday/bsvc Request_References"`
	ResponseFilter    *ResponseFilterType                              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Holds the Get Probation Periods For a Worker Response
type GetProbationPeriodsForWorkersResponseType struct {
	RequestReferences *ProbationPeriodsForWorkersRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *GetProbationPeriodsForWorkerResponseDataType    `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Gets Role Assignments for Role Assigners Request
type GetRoleAssignmentsForRoleAssignersRequestType struct {
	RequestReferences *RoleAssignerRequestReferencesType  `xml:"urn:com.workday/bsvc Request_References"`
	RequestCriteria   *RoleAssignmentsRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *RoleAssignerResponseGroupType      `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Role Assignments Response.
type GetRoleAssignmentsForRoleAssignersResponseType struct {
	RequestReferences *RoleAssignerRequestReferencesType  `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *RoleAssignmentsRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *RoleAssignerResponseGroupType      `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *RoleAssignerResponseDataType       `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to provide additional supporting information about a Student Employment Eligibility Event
type GetStudentEmploymentEligibilityDataRequestType struct {
	RequestReferences *StudentEmploymentEligibilityEventRequestReferencesType `xml:"urn:com.workday/bsvc Request_References"`
	RequestCriteria   []StudentEmploymentEligibilityEventRequestCriteriaType  `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *StudentEmploymentEligibilityEventResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response of Student Employment Eligibility with supporting information
type GetStudentEmploymentEligibilityDataResponseType struct {
	RequestReferences *StudentEmploymentEligibilityEventRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *StudentEmploymentEligibilityEventResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *StudentEmploymentEligibilityEventResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Container for Get Student Employment Eligibility Rule Sets Request
type GetStudentEmploymentEligibilityRuleSetRequestType struct {
	RequestReferences *StudentEmploymentEligibilityRuleSetRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *StudentEmploymentEligibilityRuleSetResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Container for Get Student Employment Eligibility Rule Sets Response
type GetStudentEmploymentEligibilityRuleSetResponseType struct {
	RequestCriteria *StudentEmploymentEligibilityRuleSetRequestReferencesType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *StudentEmploymentEligibilityRuleSetResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults *ResponseResultsType                                      `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *StudentEmploymentEligibilityRuleSetResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper for requesting Student Employment Eligibility Rule Data
type GetStudentEmploymentEligibilityRulesRequestType struct {
	RequestReferences *StudentEmploymentEligibilityRuleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *StudentEmploymentEligibilityRuleResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Container for the Student Employment Eligibility Rules Response
type GetStudentEmploymentEligibilityRulesResponseType struct {
	RequestReferences *StudentEmploymentEligibilityRuleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *StudentEmploymentEligibilityRuleResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *StudentEmploymentEligibilityRuleResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for Get Worker Documents.
type GetWorkerDocumentsRequestType struct {
	RequestReferences *WorkerDocumentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *WorkerDocumentResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Get Worker Documents.
type GetWorkerDocumentsResponseType struct {
	RequestReferences *WorkerDocumentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *WorkerDocumentResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *WorkerDocumentResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element used to find and get workers and their associated data.
type GetWorkersRequestType struct {
	RequestReferences *WorkerRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *WorkerRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *WorkerResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

type GetWorkersResponseRootType struct {
	RequestReferences          *WorkerRequestReferencesType     `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria            *WorkerRequestCriteriaType       `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter             *ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup              *WorkerResponseGroupType         `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults            *ResponseResultsType             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData               *WorkersResponseDataType         `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	InvalidReferenceIDResponse []InvalidReferenceIDResponseType `xml:"urn:com.workday/bsvc Invalid_Reference_ID_Response,omitempty"`
	Version                    string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for Get Workers Response.  The response element has to match the Operation Name.
type GetWorkersResponseType struct {
	RequestReferences          *WorkerRequestReferencesType     `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria            *WorkerRequestCriteriaType       `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter             *ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup              *WorkerResponseGroupType         `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults            *ResponseResultsType             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData               *WorkersResponseDataType         `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	InvalidReferenceIDResponse []InvalidReferenceIDResponseType `xml:"urn:com.workday/bsvc Invalid_Reference_ID_Response,omitempty"`
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

// Contains the target contribution amount for the spending account election.
type GoalDataType struct {
	GoalAmount         float64              `xml:"urn:com.workday/bsvc Goal_Amount,omitempty"`
	FrequencyReference *FrequencyObjectType `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
}

// Contains a single Goal Detail (Review Goal or Worker Goal Detail) and the Notes and History.
type GoalDetailDataType struct {
	GoalReferenceID           string                         `xml:"urn:com.workday/bsvc Goal_Reference_ID,omitempty"`
	Goal                      string                         `xml:"urn:com.workday/bsvc Goal,omitempty"`
	Description               string                         `xml:"urn:com.workday/bsvc Description,omitempty"`
	OrganizationGoalReference *OrganizationGoalObjectType    `xml:"urn:com.workday/bsvc Organization_Goal_Reference,omitempty"`
	DueDate                   *time.Time                     `xml:"urn:com.workday/bsvc Due_Date,omitempty"`
	CompletionStatusReference *ComponentCompletionObjectType `xml:"urn:com.workday/bsvc Completion_Status_Reference,omitempty"`
	CompletionDate            *time.Time                     `xml:"urn:com.workday/bsvc Completion_Date,omitempty"`
}

func (t *GoalDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T GoalDetailDataType
	var layout struct {
		*T
		DueDate        *xsdDate `xml:"urn:com.workday/bsvc Due_Date,omitempty"`
		CompletionDate *xsdDate `xml:"urn:com.workday/bsvc Completion_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DueDate = (*xsdDate)(layout.T.DueDate)
	layout.CompletionDate = (*xsdDate)(layout.T.CompletionDate)
	return e.EncodeElement(layout, start)
}
func (t *GoalDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T GoalDetailDataType
	var overlay struct {
		*T
		DueDate        *xsdDate `xml:"urn:com.workday/bsvc Due_Date,omitempty"`
		CompletionDate *xsdDate `xml:"urn:com.workday/bsvc Completion_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DueDate = (*xsdDate)(overlay.T.DueDate)
	overlay.CompletionDate = (*xsdDate)(overlay.T.CompletionDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type GoalObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type GoalObjectType struct {
	ID         []GoalObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Wrapper for Government Identification Data. Includes National Identifiers and Government Identifiers.
type GovernmentIdentificationDataType struct {
	NationalID   []NationalIDType   `xml:"urn:com.workday/bsvc National_ID,omitempty"`
	GovernmentID []GovernmentIDType `xml:"urn:com.workday/bsvc Government_ID,omitempty"`
	ReplaceAll   bool               `xml:"urn:com.workday/bsvc Replace_All,attr,omitempty"`
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
type HeadcountOptionsObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type HeadcountOptionsObjectType struct {
	ID         []HeadcountOptionsObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type HeadcountRestrictionsObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type HeadcountRestrictionsObjectType struct {
	ID         []HeadcountRestrictionsObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type HealthCareClassificationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type HealthCareClassificationObjectType struct {
	ID         []HealthCareClassificationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains the spending account coverage information (elections) for an employee.
type HealthSavingsAccountCoverageDataType struct {
	BenefitElectionData                  *WorkerBenefitElectionDataType            `xml:"urn:com.workday/bsvc Benefit_Election_Data,omitempty"`
	HealthSavingsAccountElectionInfoData *HealthSavingsAccountElectionInfoDataType `xml:"urn:com.workday/bsvc Health_Savings_Account_Election_Info_Data,omitempty"`
	EmployeeContributionData             *EmployeeContributionDataType             `xml:"urn:com.workday/bsvc Employee_Contribution_Data,omitempty"`
	EmployerContributionData             *EmployerContributionDataType             `xml:"urn:com.workday/bsvc Employer_Contribution_Data,omitempty"`
}

// Benefit Plan Earnings Deduction Details Data
type HealthSavingsAccountElectionInfoDataType struct {
	CoverageTargetName              string              `xml:"urn:com.workday/bsvc Coverage_Target_Name,omitempty"`
	MaximumAnnualContributionAmount float64             `xml:"urn:com.workday/bsvc Maximum_Annual_Contribution_Amount,omitempty"`
	CurrencyReference               *CurrencyObjectType `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type HierarchicAssignerObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type HierarchicAssignerObjectType struct {
	ID         []HierarchicAssignerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper Element for the Hire Employee business process web service and its sub business processes.
type HireEmployeeBusinessProcessDataType struct {
	ApplicantReference                              *ApplicantObjectType                                     `xml:"urn:com.workday/bsvc Applicant_Reference"`
	FormerWorkerReference                           *FormerWorkerObjectType                                  `xml:"urn:com.workday/bsvc Former_Worker_Reference"`
	StudentReference                                *StudentObjectType                                       `xml:"urn:com.workday/bsvc Student_Reference"`
	AcademicAffiliateReference                      *AcademicAffiliateObjectType                             `xml:"urn:com.workday/bsvc Academic_Affiliate_Reference"`
	ApplicantData                                   *CreateApplicantDataType                                 `xml:"urn:com.workday/bsvc Applicant_Data"`
	OrganizationReference                           *SupervisoryOrganizationObjectType                       `xml:"urn:com.workday/bsvc Organization_Reference,omitempty"`
	PositionReference                               *PositionRestrictionsObjectType                          `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	JobRequisitionReference                         *JobRequisitionObjectType                                `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	HireDate                                        *time.Time                                               `xml:"urn:com.workday/bsvc Hire_Date,omitempty"`
	HireEmployeeEventData                           *HireEmployeeEventDataType                               `xml:"urn:com.workday/bsvc Hire_Employee_Event_Data"`
	ProposeCompensationforHireSubProcess            *ProposeCompensationForEmploymentSubBusinessProcessType  `xml:"urn:com.workday/bsvc Propose_Compensation_for_Hire_Sub_Process,omitempty"`
	UpdateIDInformationSubProcess                   *UpdateIDInformationSubBusinessProcessType               `xml:"urn:com.workday/bsvc Update_ID_Information_Sub_Process,omitempty"`
	EditGovernmentIDsSubProcess                     *EditGovernmentIDsSubBusinessProcessType                 `xml:"urn:com.workday/bsvc Edit_Government_IDs_Sub_Process,omitempty"`
	EditPassportsandVisasSubProcess                 *EditPassportsandVisasSubBusinessProcessType             `xml:"urn:com.workday/bsvc Edit_Passports_and_Visas_Sub_Process,omitempty"`
	EditLicenseSubProcess                           *EditLicensesSubBusinessProcessType                      `xml:"urn:com.workday/bsvc Edit_License_Sub_Process,omitempty"`
	EditCustomIDsSubProcess                         *EditCustomIDsSubBusinessProcessType                     `xml:"urn:com.workday/bsvc Edit_Custom_IDs_Sub_Process,omitempty"`
	EditAssignOrganizationSubProcess                *EditAssignPositionOrganizationSubBusinessProcessType    `xml:"urn:com.workday/bsvc Edit_Assign_Organization_Sub_Process,omitempty"`
	AssignPayGroupSubProcess                        *AssignPayGroupSubBusinessProcessType                    `xml:"urn:com.workday/bsvc Assign_Pay_Group_Sub_Process,omitempty"`
	ReviewPayrollInterfaceSubProcess                *ReviewPayrollInterfaceDataSubBusinessProcessType        `xml:"urn:com.workday/bsvc Review_Payroll_Interface_Sub_Process,omitempty"`
	RequestOneTimePaymentSubProcess                 *RequestOneTimePaymentSubBusinessProcessType             `xml:"urn:com.workday/bsvc Request_One_Time_Payment_Sub_Process,omitempty"`
	RequestOneTimePaymentforReferralSubProcess      *OneTimePaymentforReferralSubBusinessProcessType         `xml:"urn:com.workday/bsvc Request_One_Time_Payment_for_Referral_Sub_Process,omitempty"`
	RequestStockGrantSubProcess                     *RequestStockSubBusinessProcessType                      `xml:"urn:com.workday/bsvc Request_Stock_Grant_Sub_Process,omitempty"`
	CreateWorkdayAccountSubProcess                  *CreateWorkdayAccountSubBusinessProcessType              `xml:"urn:com.workday/bsvc Create_Workday_Account_Sub_Process,omitempty"`
	AssignMatrixOrganizationSubProcess              *AssignMatrixOrganizationSubBusinessProcessType          `xml:"urn:com.workday/bsvc Assign_Matrix_Organization_Sub_Process,omitempty"`
	ChangePersonalInformationSubProcess             *ChangePersonalInformationSubBusinessProcessType         `xml:"urn:com.workday/bsvc Change_Personal_Information_Sub_Process,omitempty"`
	CreateProvisioningEventSubProcess               *CreateProvisioningEventSubBusinessProcessType           `xml:"urn:com.workday/bsvc Create_Provisioning_Event_Sub_Process,omitempty"`
	CreateBenefitLifeEventSubProcess                *CreateBenefitLifeEventSubBusinessProcessType            `xml:"urn:com.workday/bsvc Create_Benefit_Life_Event_Sub_Process,omitempty"`
	MaintainEmployeeContractsSubBusinessProcess     *MaintainEmployeeContractsSubBusinessProcessType         `xml:"urn:com.workday/bsvc Maintain_Employee_Contracts_Sub_Business_Process,omitempty"`
	EditServiceDatesSubProcess                      *EditServiceDatesSubBusinessProcessType                  `xml:"urn:com.workday/bsvc Edit_Service_Dates_Sub_Process,omitempty"`
	RemoveRetireeStatusSubProcess                   *RemoveRetireeStatusSubBusinessProcessType               `xml:"urn:com.workday/bsvc Remove_Retiree_Status_Sub_Process,omitempty"`
	CheckPositionBudgetSubProcess                   *CheckPositionBudgetSubBusinessProcessType               `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process,omitempty"`
	AssignCostingAllocationSubProcess               *AssignCostingAllocationSubBusinessProcessType           `xml:"urn:com.workday/bsvc Assign_Costing_Allocation_Sub_Process,omitempty"`
	EditBackgroundCheckSubProcess                   *EditBackgroundCheckSubBusinessProcessType               `xml:"urn:com.workday/bsvc Edit_Background_Check_Sub_Process,omitempty"`
	AddAcademicAppointmentSubProcess                *AddAcademicAppointmentSubBusinessProcessType            `xml:"urn:com.workday/bsvc Add_Academic_Appointment_Sub_Process,omitempty"`
	AssignEmployeeCollectiveAgreementSubProcess     *AssignEmployeeCollectiveAgreementSubBusinessProcessType `xml:"urn:com.workday/bsvc Assign_Employee_Collective_Agreement_Sub_Process,omitempty"`
	ManageEmployeeProbationPeriodSubBusinessProcess *ManageProbationPeriodSubBusinessProcessType             `xml:"urn:com.workday/bsvc Manage_Employee_Probation_Period_Sub_Business_Process,omitempty"`
	EmergencyContactsSubProcess                     *ChangeEmergencyContactsSubBusinessProcessType           `xml:"urn:com.workday/bsvc Emergency_Contacts_Sub_Process,omitempty"`
	OnboardingSetupSubProcess                       *OnboardingSetupSubBusinessProcessType                   `xml:"urn:com.workday/bsvc Onboarding_Setup_Sub_Process,omitempty"`
	StudentEmploymentEligibilitySubProcess          *StudentEmploymentEligibilitySubBusinessProcessType      `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Sub_Process,omitempty"`
	ManageUnionMembershipSubProcess                 *ManageUnionMembershipSubBusinessProcessType             `xml:"urn:com.workday/bsvc Manage_Union_Membership_Sub_Process,omitempty"`
	EditNoticePeriodsSubProcess                     *EditNoticePeriodsSubBusinessProcessType                 `xml:"urn:com.workday/bsvc Edit_Notice_Periods_Sub_Process,omitempty"`
}

func (t *HireEmployeeBusinessProcessDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T HireEmployeeBusinessProcessDataType
	var layout struct {
		*T
		HireDate *xsdDate `xml:"urn:com.workday/bsvc Hire_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.HireDate = (*xsdDate)(layout.T.HireDate)
	return e.EncodeElement(layout, start)
}
func (t *HireEmployeeBusinessProcessDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T HireEmployeeBusinessProcessDataType
	var overlay struct {
		*T
		HireDate *xsdDate `xml:"urn:com.workday/bsvc Hire_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.HireDate = (*xsdDate)(overlay.T.HireDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for the details for an Employee Hire.
type HireEmployeeEventDataType struct {
	EmployeeID                  string                                    `xml:"urn:com.workday/bsvc Employee_ID,omitempty"`
	PositionID                  string                                    `xml:"urn:com.workday/bsvc Position_ID,omitempty"`
	HireReasonReference         *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Hire_Reason_Reference,omitempty"`
	EmployeeTypeReference       *PositionWorkerTypeObjectType             `xml:"urn:com.workday/bsvc Employee_Type_Reference,omitempty"`
	FirstDayofWork              *time.Time                                `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
	TimeofHire                  *time.Time                                `xml:"urn:com.workday/bsvc Time_of_Hire,omitempty"`
	ContinuousServiceDate       *time.Time                                `xml:"urn:com.workday/bsvc Continuous_Service_Date,omitempty"`
	ProbationStartDate          *time.Time                                `xml:"urn:com.workday/bsvc Probation_Start_Date,omitempty"`
	ProbationEndDate            *time.Time                                `xml:"urn:com.workday/bsvc Probation_End_Date,omitempty"`
	EmploymentEndDate           *time.Time                                `xml:"urn:com.workday/bsvc Employment_End_Date,omitempty"`
	BenefitsServiceDate         *time.Time                                `xml:"urn:com.workday/bsvc Benefits_Service_Date,omitempty"`
	CompanyServiceDate          *time.Time                                `xml:"urn:com.workday/bsvc Company_Service_Date,omitempty"`
	ConversionPositionStartDate *time.Time                                `xml:"urn:com.workday/bsvc Conversion_Position_Start_Date,omitempty"`
	PositionDetails             *PositionDetailsSubDataType               `xml:"urn:com.workday/bsvc Position_Details"`
	EmployeeExternalIDData      *ExternalIDDataType                       `xml:"urn:com.workday/bsvc Employee_External_ID_Data,omitempty"`
}

func (t *HireEmployeeEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T HireEmployeeEventDataType
	var layout struct {
		*T
		FirstDayofWork              *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
		TimeofHire                  *xsdTime `xml:"urn:com.workday/bsvc Time_of_Hire,omitempty"`
		ContinuousServiceDate       *xsdDate `xml:"urn:com.workday/bsvc Continuous_Service_Date,omitempty"`
		ProbationStartDate          *xsdDate `xml:"urn:com.workday/bsvc Probation_Start_Date,omitempty"`
		ProbationEndDate            *xsdDate `xml:"urn:com.workday/bsvc Probation_End_Date,omitempty"`
		EmploymentEndDate           *xsdDate `xml:"urn:com.workday/bsvc Employment_End_Date,omitempty"`
		BenefitsServiceDate         *xsdDate `xml:"urn:com.workday/bsvc Benefits_Service_Date,omitempty"`
		CompanyServiceDate          *xsdDate `xml:"urn:com.workday/bsvc Company_Service_Date,omitempty"`
		ConversionPositionStartDate *xsdDate `xml:"urn:com.workday/bsvc Conversion_Position_Start_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FirstDayofWork = (*xsdDate)(layout.T.FirstDayofWork)
	layout.TimeofHire = (*xsdTime)(layout.T.TimeofHire)
	layout.ContinuousServiceDate = (*xsdDate)(layout.T.ContinuousServiceDate)
	layout.ProbationStartDate = (*xsdDate)(layout.T.ProbationStartDate)
	layout.ProbationEndDate = (*xsdDate)(layout.T.ProbationEndDate)
	layout.EmploymentEndDate = (*xsdDate)(layout.T.EmploymentEndDate)
	layout.BenefitsServiceDate = (*xsdDate)(layout.T.BenefitsServiceDate)
	layout.CompanyServiceDate = (*xsdDate)(layout.T.CompanyServiceDate)
	layout.ConversionPositionStartDate = (*xsdDate)(layout.T.ConversionPositionStartDate)
	return e.EncodeElement(layout, start)
}
func (t *HireEmployeeEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T HireEmployeeEventDataType
	var overlay struct {
		*T
		FirstDayofWork              *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
		TimeofHire                  *xsdTime `xml:"urn:com.workday/bsvc Time_of_Hire,omitempty"`
		ContinuousServiceDate       *xsdDate `xml:"urn:com.workday/bsvc Continuous_Service_Date,omitempty"`
		ProbationStartDate          *xsdDate `xml:"urn:com.workday/bsvc Probation_Start_Date,omitempty"`
		ProbationEndDate            *xsdDate `xml:"urn:com.workday/bsvc Probation_End_Date,omitempty"`
		EmploymentEndDate           *xsdDate `xml:"urn:com.workday/bsvc Employment_End_Date,omitempty"`
		BenefitsServiceDate         *xsdDate `xml:"urn:com.workday/bsvc Benefits_Service_Date,omitempty"`
		CompanyServiceDate          *xsdDate `xml:"urn:com.workday/bsvc Company_Service_Date,omitempty"`
		ConversionPositionStartDate *xsdDate `xml:"urn:com.workday/bsvc Conversion_Position_Start_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FirstDayofWork = (*xsdDate)(overlay.T.FirstDayofWork)
	overlay.TimeofHire = (*xsdTime)(overlay.T.TimeofHire)
	overlay.ContinuousServiceDate = (*xsdDate)(overlay.T.ContinuousServiceDate)
	overlay.ProbationStartDate = (*xsdDate)(overlay.T.ProbationStartDate)
	overlay.ProbationEndDate = (*xsdDate)(overlay.T.ProbationEndDate)
	overlay.EmploymentEndDate = (*xsdDate)(overlay.T.EmploymentEndDate)
	overlay.BenefitsServiceDate = (*xsdDate)(overlay.T.BenefitsServiceDate)
	overlay.CompanyServiceDate = (*xsdDate)(overlay.T.CompanyServiceDate)
	overlay.ConversionPositionStartDate = (*xsdDate)(overlay.T.ConversionPositionStartDate)
	return d.DecodeElement(&overlay, &start)
}

// Responds with the Event ID for the Hire Employee Event.
type HireEmployeeEventResponseType struct {
	EventReference         *UniqueIdentifierObjectType                    `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	EmployeeReference      *WorkerObjectType                              `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	PositionReference      *PositionElementObjectType                     `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	ApplicantReference     *ApplicantObjectType                           `xml:"urn:com.workday/bsvc Applicant_Reference,omitempty"`
	ExceptionsResponseData []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This web service operation is designed to hire an existing applicant into an Employee position
type HireEmployeeRequestHVType struct {
	BusinessProcessParameters *BusinessProcessParametersType       `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	HireEmployeeData          *HireEmployeeBusinessProcessDataType `xml:"urn:com.workday/bsvc Hire_Employee_Data"`
}

// This web service operation is designed to hire an existing applicant into an Employee position or job using the Hire Employee business process.
type HireEmployeeRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType       `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	HireEmployeeData          *HireEmployeeBusinessProcessDataType `xml:"urn:com.workday/bsvc Hire_Employee_Data"`
	Version                   string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// High Volume request to process each worker, position, and matrix organization
type ImportAssignMatrixOrganizationRequestType struct {
	ID                                    string                                  `xml:"urn:com.workday/bsvc ID,omitempty"`
	AssignMatrixOrganizationInformationHV []AssignMatrixOrganizationRequestHVType `xml:"urn:com.workday/bsvc Assign_Matrix_Organization_Information_HV,omitempty"`
	Version                               string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Root Import Change Job Request Element
type ImportChangeJobRequestType struct {
	ID                     string                   `xml:"urn:com.workday/bsvc ID,omitempty"`
	ChangeJobInformationHV []ChangeJobRequestHVType `xml:"urn:com.workday/bsvc Change_Job_Information_HV,omitempty"`
	Version                string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Root Import Contract Contingent Worker Request Element
type ImportContractContingentWorkerRequestType struct {
	ID                                    string                                  `xml:"urn:com.workday/bsvc ID,omitempty"`
	ContractContingentWorkerInformationHV []ContractContingentWorkerRequestHVType `xml:"urn:com.workday/bsvc Contract_Contingent_Worker_Information_HV,omitempty"`
	Version                               string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Root Import End Contingent Worker Contract Request Element
type ImportEndContingentWorkerContractRequestType struct {
	ID                                       string                                     `xml:"urn:com.workday/bsvc ID,omitempty"`
	EndContingentWorkerContractInformationHV []EndContingentWorkerContractRequestHVType `xml:"urn:com.workday/bsvc End_Contingent_Worker_Contract_Information_HV,omitempty"`
	Version                                  string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Header element of the Import External Student Information web service task
type ImportExternalStudentRequestType struct {
	ID                  string                           `xml:"urn:com.workday/bsvc ID,omitempty"`
	ExternalStudentData []ExternalAcademicRecordDataType `xml:"urn:com.workday/bsvc External_Student_Data,omitempty"`
	Version             string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Root Import Hire Employee Request Element
type ImportHireEmployeeRequestType struct {
	ID                        string                      `xml:"urn:com.workday/bsvc ID,omitempty"`
	HireEmployeeInformationHV []HireEmployeeRequestHVType `xml:"urn:com.workday/bsvc Hire_Employee_Information_HV,omitempty"`
	Version                   string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request for each worker, position, matrix org in Remove Matrix Organization
type ImportRemoveMatrixOrganizationRequestHVType struct {
	BusinessProcessParameters                   *BusinessProcessParametersType                   `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	RemoveMatrixOrganizationBusinessProcessData *RemoveMatrixOrganizationBusinessProcessDataType `xml:"urn:com.workday/bsvc Remove_Matrix_Organization_Business_Process_Data"`
}

// High Volume request to process each worker, position, and matrix organization
type ImportRemoveMatrixOrganizationRequestType struct {
	ID                                    string                                        `xml:"urn:com.workday/bsvc ID,omitempty"`
	RemoveMatrixOrganizationInformationHV []ImportRemoveMatrixOrganizationRequestHVType `xml:"urn:com.workday/bsvc Remove_Matrix_Organization_Information_HV,omitempty"`
	Version                               string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Business Process for swapping positions between workers
type ImportSwapPositionsRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	MassSwapPositionsData     *MassSwapPositionsDataType     `xml:"urn:com.workday/bsvc Mass_Swap_Positions_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Root Import Terminate Employee Request Element
type ImportTerminateEmployeeRequestType struct {
	ID                             string                           `xml:"urn:com.workday/bsvc ID,omitempty"`
	TerminateEmployeeInformationHV []TerminateEmployeeRequestHVType `xml:"urn:com.workday/bsvc Terminate_Employee_Information_HV,omitempty"`
	Version                        string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Data element for the Request Stock Grant business process.
type IndividualStockGrantDataType struct {
	SharesGranted                float64                         `xml:"urn:com.workday/bsvc Shares_Granted,omitempty"`
	GrantPercent                 float64                         `xml:"urn:com.workday/bsvc Grant_Percent,omitempty"`
	GrantAmount                  float64                         `xml:"urn:com.workday/bsvc Grant_Amount,omitempty"`
	GrantTypeReference           *StockGrantTypeObjectType       `xml:"urn:com.workday/bsvc Grant_Type_Reference"`
	GrantAmountCurrencyReference *CurrencyObjectType             `xml:"urn:com.workday/bsvc Grant_Amount_Currency_Reference,omitempty"`
	OptionPricingFactor          float64                         `xml:"urn:com.workday/bsvc Option_Pricing_Factor,omitempty"`
	Comments                     string                          `xml:"urn:com.workday/bsvc Comments,omitempty"`
	VestingScheduleReference     *StockVestingScheduleObjectType `xml:"urn:com.workday/bsvc Vesting_Schedule_Reference,omitempty"`
	GrantID                      string                          `xml:"urn:com.workday/bsvc Grant_ID,omitempty"`
	GrantDate                    *time.Time                      `xml:"urn:com.workday/bsvc Grant_Date,omitempty"`
	VestFromDate                 *time.Time                      `xml:"urn:com.workday/bsvc Vest_From_Date,omitempty"`
	ExpirationDate               *time.Time                      `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	GrantPrice                   float64                         `xml:"urn:com.workday/bsvc Grant_Price,omitempty"`
	CurrencyReference            *CurrencyObjectType             `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	BoardApproved                *bool                           `xml:"urn:com.workday/bsvc Board_Approved,omitempty"`
}

func (t *IndividualStockGrantDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T IndividualStockGrantDataType
	var layout struct {
		*T
		GrantDate      *xsdDate `xml:"urn:com.workday/bsvc Grant_Date,omitempty"`
		VestFromDate   *xsdDate `xml:"urn:com.workday/bsvc Vest_From_Date,omitempty"`
		ExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.GrantDate = (*xsdDate)(layout.T.GrantDate)
	layout.VestFromDate = (*xsdDate)(layout.T.VestFromDate)
	layout.ExpirationDate = (*xsdDate)(layout.T.ExpirationDate)
	return e.EncodeElement(layout, start)
}
func (t *IndividualStockGrantDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T IndividualStockGrantDataType
	var overlay struct {
		*T
		GrantDate      *xsdDate `xml:"urn:com.workday/bsvc Grant_Date,omitempty"`
		VestFromDate   *xsdDate `xml:"urn:com.workday/bsvc Vest_From_Date,omitempty"`
		ExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.GrantDate = (*xsdDate)(overlay.T.GrantDate)
	overlay.VestFromDate = (*xsdDate)(overlay.T.VestFromDate)
	overlay.ExpirationDate = (*xsdDate)(overlay.T.ExpirationDate)
	return d.DecodeElement(&overlay, &start)
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

// Contains Instructor Eligibility Line Data
type InstructorEligibilityLineDataType struct {
	AcademicUnitReference            *AcademicUnitObjectType             `xml:"urn:com.workday/bsvc Academic_Unit_Reference,omitempty"`
	AcademicLevelReference           []AcademicLevelObjectType           `xml:"urn:com.workday/bsvc Academic_Level_Reference,omitempty"`
	CourseSubjectReference           []CourseSubjectObjectType           `xml:"urn:com.workday/bsvc Course_Subject_Reference,omitempty"`
	CourseReference                  []CourseDefinitionObjectType        `xml:"urn:com.workday/bsvc Course_Reference,omitempty"`
	DeliveryModeReference            []DeliveryModeObjectType            `xml:"urn:com.workday/bsvc Delivery_Mode_Reference,omitempty"`
	CourseTagReference               []StudentCourseTagObjectType        `xml:"urn:com.workday/bsvc Course_Tag_Reference,omitempty"`
	LocationReference                []LocationObjectType                `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	EducationalTaxonomyCodeReference []EducationalTaxonomyCodeObjectType `xml:"urn:com.workday/bsvc Educational_Taxonomy_Code_Reference,omitempty"`
}

// Contains the coverage level detailed information and volume for the insurance election.
type InsuranceCoverageLevelDataType struct {
	CoverageLevelValueData *CoverageLevelValueDataType `xml:"urn:com.workday/bsvc Coverage_Level_Value_Data"`
	VolumeData             *VolumeDataType             `xml:"urn:com.workday/bsvc Volume_Data"`
}

// Contains a unique identifier for an instance of an object.
type InsuranceCoverageTargetObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InsuranceCoverageTargetObjectType struct {
	ID         []InsuranceCoverageTargetObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Wrapper element that contains position's international assignment information
type InternationalAssignmentDataType struct {
	InternationalAssignmentTypeReference        *InternationalAssignmentTypeObjectType    `xml:"urn:com.workday/bsvc International_Assignment_Type_Reference,omitempty"`
	StartInternationalAssignmentReasonReference *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Start_International_Assignment_Reason_Reference,omitempty"`
	ExpectedAssignmentEndDate                   *time.Time                                `xml:"urn:com.workday/bsvc Expected_Assignment_End_Date,omitempty"`
	HostCountryReference                        *CountryObjectType                        `xml:"urn:com.workday/bsvc Host_Country_Reference,omitempty"`
	HomeCountryReference                        *CountryObjectType                        `xml:"urn:com.workday/bsvc Home_Country_Reference,omitempty"`
}

func (t *InternationalAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InternationalAssignmentDataType
	var layout struct {
		*T
		ExpectedAssignmentEndDate *xsdDate `xml:"urn:com.workday/bsvc Expected_Assignment_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ExpectedAssignmentEndDate = (*xsdDate)(layout.T.ExpectedAssignmentEndDate)
	return e.EncodeElement(layout, start)
}
func (t *InternationalAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InternationalAssignmentDataType
	var overlay struct {
		*T
		ExpectedAssignmentEndDate *xsdDate `xml:"urn:com.workday/bsvc Expected_Assignment_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ExpectedAssignmentEndDate = (*xsdDate)(overlay.T.ExpectedAssignmentEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element that contains basic information of worker's international assignment if any.
type InternationalAssignmentSummaryDataType struct {
	HasInternationalAssignment *bool               `xml:"urn:com.workday/bsvc Has_International_Assignment,omitempty"`
	HostCountriesReference     []CountryObjectType `xml:"urn:com.workday/bsvc Host_Countries_Reference,omitempty"`
	HomeCountryReference       *CountryObjectType  `xml:"urn:com.workday/bsvc Home_Country_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type InternationalAssignmentTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InternationalAssignmentTypeObjectType struct {
	ID         []InternationalAssignmentTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type JobApplicationTemplateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobApplicationTemplateObjectType struct {
	ID         []JobApplicationTemplateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Job Classification Data Element
type JobClassificationDataType struct {
	ID                                  string                      `xml:"urn:com.workday/bsvc ID,omitempty"`
	JobClassificationID                 string                      `xml:"urn:com.workday/bsvc Job_Classification_ID,omitempty"`
	Description                         string                      `xml:"urn:com.workday/bsvc Description,omitempty"`
	JobClassificationMappingIDReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Job_Classification_Mapping_ID_Reference,omitempty"`
	Inactive                            *bool                       `xml:"urn:com.workday/bsvc Inactive,omitempty"`
}

// Job Classification Group Data Element
type JobClassificationGroupDataType struct {
	ID                              string                            `xml:"urn:com.workday/bsvc ID,omitempty"`
	EffectiveDate                   *time.Time                        `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	JobClassificationGroupName      string                            `xml:"urn:com.workday/bsvc Job_Classification_Group_Name,omitempty"`
	GlobalSetupDataMappingReference *GlobalSetupDataMappingObjectType `xml:"urn:com.workday/bsvc Global_Setup_Data_Mapping_Reference,omitempty"`
	LocationReference               *LocationContextObjectType        `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	Inactive                        *bool                             `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	JobClassification               []JobClassificationType           `xml:"urn:com.workday/bsvc Job_Classification,omitempty"`
}

func (t *JobClassificationGroupDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobClassificationGroupDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *JobClassificationGroupDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobClassificationGroupDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Response element containing an instance of Job Classification Group and its associated data.
type JobClassificationGroupNewType struct {
	JobClassificationGroupReference *JobClassificationGroupObjectType         `xml:"urn:com.workday/bsvc Job_Classification_Group_Reference,omitempty"`
	JobClassificationGroupNameData  []JobClassificationGroupReferenceDataType `xml:"urn:com.workday/bsvc Job_Classification_Group_Name_Data,omitempty"`
	JobClassificationGroupData      []JobClassificationGroupDataType          `xml:"urn:com.workday/bsvc Job_Classification_Group_Data,omitempty"`
	AsOfDate                        time.Time                                 `xml:"urn:com.workday/bsvc As_Of_Date,attr,omitempty"`
	AsOfMoment                      time.Time                                 `xml:"urn:com.workday/bsvc As_Of_Moment,attr,omitempty"`
}

func (t *JobClassificationGroupNewType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobClassificationGroupNewType
	var layout struct {
		*T
		AsOfDate   *xsdDate     `xml:"urn:com.workday/bsvc As_Of_Date,attr,omitempty"`
		AsOfMoment *xsdDateTime `xml:"urn:com.workday/bsvc As_Of_Moment,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AsOfDate = (*xsdDate)(&layout.T.AsOfDate)
	layout.AsOfMoment = (*xsdDateTime)(&layout.T.AsOfMoment)
	return e.EncodeElement(layout, start)
}
func (t *JobClassificationGroupNewType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobClassificationGroupNewType
	var overlay struct {
		*T
		AsOfDate   *xsdDate     `xml:"urn:com.workday/bsvc As_Of_Date,attr,omitempty"`
		AsOfMoment *xsdDateTime `xml:"urn:com.workday/bsvc As_Of_Moment,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AsOfDate = (*xsdDate)(&overlay.T.AsOfDate)
	overlay.AsOfMoment = (*xsdDateTime)(&overlay.T.AsOfMoment)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type JobClassificationGroupObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobClassificationGroupObjectType struct {
	ID         []JobClassificationGroupObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// A unique identifier used to reference a Job Classification Group.
type JobClassificationGroupReferenceDataType struct {
	JobClassificationGroupName string `xml:"urn:com.workday/bsvc Job_Classification_Group_Name"`
}

// Utilize the Request References element to retrieve a specific instance(s) of Job Classification Group and its associated data.
type JobClassificationGroupRequestReferencesType struct {
	JobClassificationGroupReference []JobClassificationGroupObjectType `xml:"urn:com.workday/bsvc Job_Classification_Group_Reference"`
}

// Element containing the Job Classification Group response data
type JobClassificationGroupResponseDataType struct {
	JobClassificationGroup []JobClassificationGroupNewType `xml:"urn:com.workday/bsvc Job_Classification_Group,omitempty"`
}

// Get Job Classification Group Response Group
type JobClassificationGroupResponseGroupType struct {
	IncludeReference          *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeJobClassifications *bool `xml:"urn:com.workday/bsvc Include_Job_Classifications,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JobClassificationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobClassificationObjectType struct {
	ID         []JobClassificationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the job classifications from the job profile for the position and the additional job classifications specified for the position.
type JobClassificationSummaryDataType struct {
	JobClassificationReference *JobClassificationObjectType      `xml:"urn:com.workday/bsvc Job_Classification_Reference,omitempty"`
	JobGroupReference          *JobClassificationGroupObjectType `xml:"urn:com.workday/bsvc Job_Group_Reference,omitempty"`
	Additional                 bool                              `xml:"urn:com.workday/bsvc Additional,attr,omitempty"`
}

// Job Classification Group Data
type JobClassificationType struct {
	JobClassificationReference *JobClassificationObjectType `xml:"urn:com.workday/bsvc Job_Classification_Reference,omitempty"`
	JobClassificationData      *JobClassificationDataType   `xml:"urn:com.workday/bsvc Job_Classification_Data,omitempty"`
	Delete                     bool                         `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Utilize the Request Criteria element to limit the Job Families returned based on the specific criteria.  Job Families that are returned must meet all of the criteria specified.
type JobFamiliesRequestCriteriaType struct {
	FieldAndParameterCriteriaData *FieldAndParameterCriteriaDataType `xml:"urn:com.workday/bsvc Field_And_Parameter_Criteria_Data,omitempty"`
	IncludeInActiveJobFamilies    *bool                              `xml:"urn:com.workday/bsvc Include_InActive_Job_Families,omitempty"`
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

// Encapsulating element containing all Job Family data.
type JobFamilyDataType struct {
	ID                           string                           `xml:"urn:com.workday/bsvc ID,omitempty"`
	EffectiveDate                *time.Time                       `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	Name                         string                           `xml:"urn:com.workday/bsvc Name,omitempty"`
	Summary                      string                           `xml:"urn:com.workday/bsvc Summary,omitempty"`
	Inactive                     *bool                            `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	JobProfileData               []JobProfileforJobFamilyDataType `xml:"urn:com.workday/bsvc Job_Profile_Data,omitempty"`
	IntegrationFieldOverrideData []DocumentFieldResultDataType    `xml:"urn:com.workday/bsvc Integration_Field_Override_Data,omitempty"`
}

func (t *JobFamilyDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobFamilyDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *JobFamilyDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobFamilyDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Job Family Group data
type JobFamilyGroupDataType struct {
	ID            string                        `xml:"urn:com.workday/bsvc ID,omitempty"`
	EffectiveDate *time.Time                    `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	Name          string                        `xml:"urn:com.workday/bsvc Name,omitempty"`
	Summary       string                        `xml:"urn:com.workday/bsvc Summary,omitempty"`
	Inactive      *bool                         `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	JobFamilyData []JobFamilyGroupJobFamilyType `xml:"urn:com.workday/bsvc Job_Family_Data,omitempty"`
}

func (t *JobFamilyGroupDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobFamilyGroupDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *JobFamilyGroupDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobFamilyGroupDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Info only data about a Job Family
type JobFamilyGroupJobFamilyDataType struct {
	JobFamilyName      string                      `xml:"urn:com.workday/bsvc Job_Family_Name,omitempty"`
	JobFamilySummary   string                      `xml:"urn:com.workday/bsvc Job_Family_Summary,omitempty"`
	Inactive           *bool                       `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	JobProfileInfoData []JobProfileSummaryDataType `xml:"urn:com.workday/bsvc Job_Profile_Info_Data,omitempty"`
}

// Job Family Data.  Contains the references to Job Families contained in the group.  Optionally contains other data about the contained families, however this extra data is for information only and cannot be updated here.
type JobFamilyGroupJobFamilyType struct {
	JobFamilyReference *JobFamilyObjectType             `xml:"urn:com.workday/bsvc Job_Family_Reference"`
	JobFamilyInfoData  *JobFamilyGroupJobFamilyDataType `xml:"urn:com.workday/bsvc Job_Family_Info_Data,omitempty"`
	Delete             bool                             `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Encapsulating element containing all Job Family Group data.
type JobFamilyGroupNewType struct {
	JobFamilyGroupReference *JobFamilyObjectType    `xml:"urn:com.workday/bsvc Job_Family_Group_Reference,omitempty"`
	JobFamilyGroupData      *JobFamilyGroupDataType `xml:"urn:com.workday/bsvc Job_Family_Group_Data,omitempty"`
	AsOfDate                time.Time               `xml:"urn:com.workday/bsvc As_Of_Date,attr,omitempty"`
	AsOfMoment              time.Time               `xml:"urn:com.workday/bsvc As_Of_Moment,attr,omitempty"`
}

func (t *JobFamilyGroupNewType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobFamilyGroupNewType
	var layout struct {
		*T
		AsOfDate   *xsdDate     `xml:"urn:com.workday/bsvc As_Of_Date,attr,omitempty"`
		AsOfMoment *xsdDateTime `xml:"urn:com.workday/bsvc As_Of_Moment,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AsOfDate = (*xsdDate)(&layout.T.AsOfDate)
	layout.AsOfMoment = (*xsdDateTime)(&layout.T.AsOfMoment)
	return e.EncodeElement(layout, start)
}
func (t *JobFamilyGroupNewType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobFamilyGroupNewType
	var overlay struct {
		*T
		AsOfDate   *xsdDate     `xml:"urn:com.workday/bsvc As_Of_Date,attr,omitempty"`
		AsOfMoment *xsdDateTime `xml:"urn:com.workday/bsvc As_Of_Moment,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AsOfDate = (*xsdDate)(&overlay.T.AsOfDate)
	overlay.AsOfMoment = (*xsdDateTime)(&overlay.T.AsOfMoment)
	return d.DecodeElement(&overlay, &start)
}

// Utilize the Request References element to retrieve a specific instance(s) of Job Family Group and its associated data.
type JobFamilyGroupRequestReferencesType struct {
	JobFamilyGroupReference  []JobFamilyObjectType `xml:"urn:com.workday/bsvc Job_Family_Group_Reference"`
	SkipNonExistingInstances bool                  `xml:"urn:com.workday/bsvc Skip_Non_Existing_Instances,attr,omitempty"`
}

// Response element for the Get Job Family request
type JobFamilyGroupResponseDataType struct {
	JobFamilyGroup []JobFamilyGroupNewType `xml:"urn:com.workday/bsvc Job_Family_Group,omitempty"`
}

// Get Job Family Group Response Group
type JobFamilyGroupResponseGroupType struct {
	IncludeReference          *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeJobProfileInfoData *bool `xml:"urn:com.workday/bsvc Include_Job_Profile_Info_Data,omitempty"`
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

// The Job Family reference(s) to be retrieved. Does not support job family groups.
type JobFamilyRequestReferencesType struct {
	JobFamilyReference       []JobFamilyObjectType `xml:"urn:com.workday/bsvc Job_Family_Reference"`
	SkipNonExistingInstances bool                  `xml:"urn:com.workday/bsvc Skip_Non_Existing_Instances,attr,omitempty"`
}

// Job Family Response Data containing the retrieved Job Family information
type JobFamilyResponseDataType struct {
	JobFamily []JobFamilyType `xml:"urn:com.workday/bsvc Job_Family,omitempty"`
}

// Job Family Response Group
type JobFamilyResponseGroupType struct {
	IncludeReference          *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeJobProfileInfoData *bool `xml:"urn:com.workday/bsvc Include_Job_Profile_Info_Data,omitempty"`
}

// Contains Job Family information
type JobFamilyType struct {
	JobFamilyReference *JobFamilyObjectType `xml:"urn:com.workday/bsvc Job_Family_Reference,omitempty"`
	JobFamilyData      *JobFamilyDataType   `xml:"urn:com.workday/bsvc Job_Family_Data,omitempty"`
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

// Contains the information about the employee's  job interests.
type JobInterestsDataType struct {
	JobProfileReference []JobProfileObjectType `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
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

// Element containing the Job Profile reference for a Job Family.
type JobProfileforJobFamilyDataType struct {
	JobProfileReference *JobProfileObjectType `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	Delete              bool                  `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Contains the job profile for the position.
type JobProfileinPositionSummaryDataType struct {
	JobProfileReference       *JobProfileObjectType       `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	JobExempt                 *bool                       `xml:"urn:com.workday/bsvc Job_Exempt,omitempty"`
	ManagementLevelReference  *ManagementLevelObjectType  `xml:"urn:com.workday/bsvc Management_Level_Reference,omitempty"`
	JobCategoryReference      *JobCategoryObjectType      `xml:"urn:com.workday/bsvc Job_Category_Reference,omitempty"`
	JobFamilyReference        []JobFamilyBaseObjectType   `xml:"urn:com.workday/bsvc Job_Family_Reference,omitempty"`
	JobProfileName            string                      `xml:"urn:com.workday/bsvc Job_Profile_Name,omitempty"`
	WorkShiftRequired         *bool                       `xml:"urn:com.workday/bsvc Work_Shift_Required,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type JobRequisitionStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobRequisitionStatusObjectType struct {
	ID         []JobRequisitionStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Leave Request Additional Fields
type LeaveRequestAdditionalFieldsType struct {
	LastDateforWhichPaid            *time.Time `xml:"urn:com.workday/bsvc Last_Date_for_Which_Paid,omitempty"`
	ExpectedDueDate                 *time.Time `xml:"urn:com.workday/bsvc Expected_Due_Date,omitempty"`
	ChildsBirthDate                 *time.Time `xml:"urn:com.workday/bsvc Child_s_Birth_Date,omitempty"`
	StillbirthBabyDeceased          *bool      `xml:"urn:com.workday/bsvc Stillbirth_Baby_Deceased,omitempty"`
	DateBabyArrivedHomeFromHospital *time.Time `xml:"urn:com.workday/bsvc Date_Baby_Arrived_Home_From_Hospital,omitempty"`
	AdoptionPlacementDate           *time.Time `xml:"urn:com.workday/bsvc Adoption_Placement_Date,omitempty"`
	AdoptionNotificationDate        *time.Time `xml:"urn:com.workday/bsvc Adoption_Notification_Date,omitempty"`
	DateChildEnteredCountry         *time.Time `xml:"urn:com.workday/bsvc Date_Child_Entered_Country,omitempty"`
	MultipleChildIndicator          *bool      `xml:"urn:com.workday/bsvc Multiple_Child_Indicator,omitempty"`
	NumberofBabiesAdoptedChildren   float64    `xml:"urn:com.workday/bsvc Number_of_Babies_Adopted_Children,omitempty"`
	NumberofPreviousBirths          float64    `xml:"urn:com.workday/bsvc Number_of_Previous_Births,omitempty"`
	NumberofPreviousMaternityLeaves float64    `xml:"urn:com.workday/bsvc Number_of_Previous_Maternity_Leaves,omitempty"`
	NumberofChildDependents         float64    `xml:"urn:com.workday/bsvc Number_of_Child_Dependents,omitempty"`
	SingleParentIndicator           *bool      `xml:"urn:com.workday/bsvc Single_Parent_Indicator,omitempty"`
	AgeofDependent                  float64    `xml:"urn:com.workday/bsvc Age_of_Dependent,omitempty"`
	WorkRelated                     *bool      `xml:"urn:com.workday/bsvc Work_Related,omitempty"`
	StopPaymentDate                 *time.Time `xml:"urn:com.workday/bsvc Stop_Payment_Date,omitempty"`
	SocialSecurityDisabilityCode    string     `xml:"urn:com.workday/bsvc Social_Security_Disability_Code,omitempty"`
	LocationDuringLeave             string     `xml:"urn:com.workday/bsvc Location_During_Leave,omitempty"`
	CaesareanSectionBirth           *bool      `xml:"urn:com.workday/bsvc Caesarean_Section_Birth,omitempty"`
	LeavePercentage                 float64    `xml:"urn:com.workday/bsvc Leave_Percentage,omitempty"`
	WeekofConfinement               *time.Time `xml:"urn:com.workday/bsvc Week_of_Confinement,omitempty"`
	LeaveEntitlementOverride        float64    `xml:"urn:com.workday/bsvc Leave_Entitlement_Override,omitempty"`
	DateofRecall                    *time.Time `xml:"urn:com.workday/bsvc Date_of_Recall,omitempty"`
}

func (t *LeaveRequestAdditionalFieldsType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LeaveRequestAdditionalFieldsType
	var layout struct {
		*T
		LastDateforWhichPaid            *xsdDate `xml:"urn:com.workday/bsvc Last_Date_for_Which_Paid,omitempty"`
		ExpectedDueDate                 *xsdDate `xml:"urn:com.workday/bsvc Expected_Due_Date,omitempty"`
		ChildsBirthDate                 *xsdDate `xml:"urn:com.workday/bsvc Child_s_Birth_Date,omitempty"`
		DateBabyArrivedHomeFromHospital *xsdDate `xml:"urn:com.workday/bsvc Date_Baby_Arrived_Home_From_Hospital,omitempty"`
		AdoptionPlacementDate           *xsdDate `xml:"urn:com.workday/bsvc Adoption_Placement_Date,omitempty"`
		AdoptionNotificationDate        *xsdDate `xml:"urn:com.workday/bsvc Adoption_Notification_Date,omitempty"`
		DateChildEnteredCountry         *xsdDate `xml:"urn:com.workday/bsvc Date_Child_Entered_Country,omitempty"`
		StopPaymentDate                 *xsdDate `xml:"urn:com.workday/bsvc Stop_Payment_Date,omitempty"`
		WeekofConfinement               *xsdDate `xml:"urn:com.workday/bsvc Week_of_Confinement,omitempty"`
		DateofRecall                    *xsdDate `xml:"urn:com.workday/bsvc Date_of_Recall,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LastDateforWhichPaid = (*xsdDate)(layout.T.LastDateforWhichPaid)
	layout.ExpectedDueDate = (*xsdDate)(layout.T.ExpectedDueDate)
	layout.ChildsBirthDate = (*xsdDate)(layout.T.ChildsBirthDate)
	layout.DateBabyArrivedHomeFromHospital = (*xsdDate)(layout.T.DateBabyArrivedHomeFromHospital)
	layout.AdoptionPlacementDate = (*xsdDate)(layout.T.AdoptionPlacementDate)
	layout.AdoptionNotificationDate = (*xsdDate)(layout.T.AdoptionNotificationDate)
	layout.DateChildEnteredCountry = (*xsdDate)(layout.T.DateChildEnteredCountry)
	layout.StopPaymentDate = (*xsdDate)(layout.T.StopPaymentDate)
	layout.WeekofConfinement = (*xsdDate)(layout.T.WeekofConfinement)
	layout.DateofRecall = (*xsdDate)(layout.T.DateofRecall)
	return e.EncodeElement(layout, start)
}
func (t *LeaveRequestAdditionalFieldsType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LeaveRequestAdditionalFieldsType
	var overlay struct {
		*T
		LastDateforWhichPaid            *xsdDate `xml:"urn:com.workday/bsvc Last_Date_for_Which_Paid,omitempty"`
		ExpectedDueDate                 *xsdDate `xml:"urn:com.workday/bsvc Expected_Due_Date,omitempty"`
		ChildsBirthDate                 *xsdDate `xml:"urn:com.workday/bsvc Child_s_Birth_Date,omitempty"`
		DateBabyArrivedHomeFromHospital *xsdDate `xml:"urn:com.workday/bsvc Date_Baby_Arrived_Home_From_Hospital,omitempty"`
		AdoptionPlacementDate           *xsdDate `xml:"urn:com.workday/bsvc Adoption_Placement_Date,omitempty"`
		AdoptionNotificationDate        *xsdDate `xml:"urn:com.workday/bsvc Adoption_Notification_Date,omitempty"`
		DateChildEnteredCountry         *xsdDate `xml:"urn:com.workday/bsvc Date_Child_Entered_Country,omitempty"`
		StopPaymentDate                 *xsdDate `xml:"urn:com.workday/bsvc Stop_Payment_Date,omitempty"`
		WeekofConfinement               *xsdDate `xml:"urn:com.workday/bsvc Week_of_Confinement,omitempty"`
		DateofRecall                    *xsdDate `xml:"urn:com.workday/bsvc Date_of_Recall,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LastDateforWhichPaid = (*xsdDate)(overlay.T.LastDateforWhichPaid)
	overlay.ExpectedDueDate = (*xsdDate)(overlay.T.ExpectedDueDate)
	overlay.ChildsBirthDate = (*xsdDate)(overlay.T.ChildsBirthDate)
	overlay.DateBabyArrivedHomeFromHospital = (*xsdDate)(overlay.T.DateBabyArrivedHomeFromHospital)
	overlay.AdoptionPlacementDate = (*xsdDate)(overlay.T.AdoptionPlacementDate)
	overlay.AdoptionNotificationDate = (*xsdDate)(overlay.T.AdoptionNotificationDate)
	overlay.DateChildEnteredCountry = (*xsdDate)(overlay.T.DateChildEnteredCountry)
	overlay.StopPaymentDate = (*xsdDate)(overlay.T.StopPaymentDate)
	overlay.WeekofConfinement = (*xsdDate)(overlay.T.WeekofConfinement)
	overlay.DateofRecall = (*xsdDate)(overlay.T.DateofRecall)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type LeaveRequestEventObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LeaveRequestEventObjectType struct {
	ID         []LeaveRequestEventObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element containing all Leave Requests that have corrected data.
type LeaveRequestsCorrectedDetailDataType struct {
	LeaveRequestEventReference     *LeaveRequestEventObjectType      `xml:"urn:com.workday/bsvc Leave_Request_Event_Reference,omitempty"`
	LeaveRequestDescription        string                            `xml:"urn:com.workday/bsvc Leave_Request_Description,omitempty"`
	LeaveReturnEventReference      *UniqueIdentifierObjectType       `xml:"urn:com.workday/bsvc Leave_Return_Event_Reference,omitempty"`
	OnLeave                        *bool                             `xml:"urn:com.workday/bsvc On_Leave,omitempty"`
	LeaveStartDate                 *time.Time                        `xml:"urn:com.workday/bsvc Leave_Start_Date,omitempty"`
	EstimatedLeaveEndDate          *time.Time                        `xml:"urn:com.workday/bsvc Estimated_Leave_End_Date,omitempty"`
	LeaveEndDate                   *time.Time                        `xml:"urn:com.workday/bsvc Leave_End_Date,omitempty"`
	FirstDayofWork                 *time.Time                        `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
	LeaveLastDayofWork             *time.Time                        `xml:"urn:com.workday/bsvc Leave_Last_Day_of_Work,omitempty"`
	LeaveofAbsenceTypeReference    *LeaveofAbsenceTypeObjectType     `xml:"urn:com.workday/bsvc Leave_of_Absence_Type_Reference,omitempty"`
	LinksBacktoPriorEventReference *LeaveRequestEventObjectType      `xml:"urn:com.workday/bsvc Links_Back_to_Prior_Event_Reference,omitempty"`
	BenefitsEffect                 *bool                             `xml:"urn:com.workday/bsvc Benefits_Effect,omitempty"`
	PayrollEffect                  *bool                             `xml:"urn:com.workday/bsvc Payroll_Effect,omitempty"`
	PaidTimeOffAccrualEffect       *bool                             `xml:"urn:com.workday/bsvc Paid_Time_Off_Accrual_Effect,omitempty"`
	ContinuousServiceAccrualEffect *bool                             `xml:"urn:com.workday/bsvc Continuous_Service_Accrual_Effect,omitempty"`
	StockVestingEffect             *bool                             `xml:"urn:com.workday/bsvc Stock_Vesting_Effect,omitempty"`
	LeaveTypeReasonReference       *LeaveTypeReasonObjectType        `xml:"urn:com.workday/bsvc Leave_Type_Reason_Reference,omitempty"`
	LeaveRequestAdditionalFields   *LeaveRequestAdditionalFieldsType `xml:"urn:com.workday/bsvc Leave_Request_Additional_Fields,omitempty"`
}

func (t *LeaveRequestsCorrectedDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LeaveRequestsCorrectedDetailDataType
	var layout struct {
		*T
		LeaveStartDate        *xsdDate `xml:"urn:com.workday/bsvc Leave_Start_Date,omitempty"`
		EstimatedLeaveEndDate *xsdDate `xml:"urn:com.workday/bsvc Estimated_Leave_End_Date,omitempty"`
		LeaveEndDate          *xsdDate `xml:"urn:com.workday/bsvc Leave_End_Date,omitempty"`
		FirstDayofWork        *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
		LeaveLastDayofWork    *xsdDate `xml:"urn:com.workday/bsvc Leave_Last_Day_of_Work,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LeaveStartDate = (*xsdDate)(layout.T.LeaveStartDate)
	layout.EstimatedLeaveEndDate = (*xsdDate)(layout.T.EstimatedLeaveEndDate)
	layout.LeaveEndDate = (*xsdDate)(layout.T.LeaveEndDate)
	layout.FirstDayofWork = (*xsdDate)(layout.T.FirstDayofWork)
	layout.LeaveLastDayofWork = (*xsdDate)(layout.T.LeaveLastDayofWork)
	return e.EncodeElement(layout, start)
}
func (t *LeaveRequestsCorrectedDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LeaveRequestsCorrectedDetailDataType
	var overlay struct {
		*T
		LeaveStartDate        *xsdDate `xml:"urn:com.workday/bsvc Leave_Start_Date,omitempty"`
		EstimatedLeaveEndDate *xsdDate `xml:"urn:com.workday/bsvc Estimated_Leave_End_Date,omitempty"`
		LeaveEndDate          *xsdDate `xml:"urn:com.workday/bsvc Leave_End_Date,omitempty"`
		FirstDayofWork        *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
		LeaveLastDayofWork    *xsdDate `xml:"urn:com.workday/bsvc Leave_Last_Day_of_Work,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LeaveStartDate = (*xsdDate)(overlay.T.LeaveStartDate)
	overlay.EstimatedLeaveEndDate = (*xsdDate)(overlay.T.EstimatedLeaveEndDate)
	overlay.LeaveEndDate = (*xsdDate)(overlay.T.LeaveEndDate)
	overlay.FirstDayofWork = (*xsdDate)(overlay.T.FirstDayofWork)
	overlay.LeaveLastDayofWork = (*xsdDate)(overlay.T.LeaveLastDayofWork)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containg all Leave Status data.
type LeaveStatusDetailDataType struct {
	LeaveRequestEventReference     *LeaveRequestEventObjectType      `xml:"urn:com.workday/bsvc Leave_Request_Event_Reference,omitempty"`
	LeaveRequestDescription        string                            `xml:"urn:com.workday/bsvc Leave_Request_Description,omitempty"`
	LeaveReturnEventReference      *UniqueIdentifierObjectType       `xml:"urn:com.workday/bsvc Leave_Return_Event_Reference,omitempty"`
	OnLeave                        *bool                             `xml:"urn:com.workday/bsvc On_Leave,omitempty"`
	LeaveStartDate                 *time.Time                        `xml:"urn:com.workday/bsvc Leave_Start_Date,omitempty"`
	EstimatedLeaveEndDate          *time.Time                        `xml:"urn:com.workday/bsvc Estimated_Leave_End_Date,omitempty"`
	LeaveEndDate                   *time.Time                        `xml:"urn:com.workday/bsvc Leave_End_Date,omitempty"`
	FirstDayOfWork                 *time.Time                        `xml:"urn:com.workday/bsvc First_Day_Of_Work,omitempty"`
	LeaveLastDayofWork             *time.Time                        `xml:"urn:com.workday/bsvc Leave_Last_Day_of_Work,omitempty"`
	LeaveofAbsenceTypeReference    *LeaveofAbsenceTypeObjectType     `xml:"urn:com.workday/bsvc Leave_of_Absence_Type_Reference,omitempty"`
	BenefitsEffect                 *bool                             `xml:"urn:com.workday/bsvc Benefits_Effect,omitempty"`
	PayrollEffect                  *bool                             `xml:"urn:com.workday/bsvc Payroll_Effect,omitempty"`
	PaidTimeOffAccrualEffect       *bool                             `xml:"urn:com.workday/bsvc Paid_Time_Off_Accrual_Effect,omitempty"`
	ContinuousServiceAccrualEffect *bool                             `xml:"urn:com.workday/bsvc Continuous_Service_Accrual_Effect,omitempty"`
	StockVestingEffect             *bool                             `xml:"urn:com.workday/bsvc Stock_Vesting_Effect,omitempty"`
	LeaveTypeReasonReference       *LeaveTypeReasonObjectType        `xml:"urn:com.workday/bsvc Leave_Type_Reason_Reference,omitempty"`
	LeaveRequestAdditionalFields   *LeaveRequestAdditionalFieldsType `xml:"urn:com.workday/bsvc Leave_Request_Additional_Fields,omitempty"`
}

func (t *LeaveStatusDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LeaveStatusDetailDataType
	var layout struct {
		*T
		LeaveStartDate        *xsdDate `xml:"urn:com.workday/bsvc Leave_Start_Date,omitempty"`
		EstimatedLeaveEndDate *xsdDate `xml:"urn:com.workday/bsvc Estimated_Leave_End_Date,omitempty"`
		LeaveEndDate          *xsdDate `xml:"urn:com.workday/bsvc Leave_End_Date,omitempty"`
		FirstDayOfWork        *xsdDate `xml:"urn:com.workday/bsvc First_Day_Of_Work,omitempty"`
		LeaveLastDayofWork    *xsdDate `xml:"urn:com.workday/bsvc Leave_Last_Day_of_Work,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LeaveStartDate = (*xsdDate)(layout.T.LeaveStartDate)
	layout.EstimatedLeaveEndDate = (*xsdDate)(layout.T.EstimatedLeaveEndDate)
	layout.LeaveEndDate = (*xsdDate)(layout.T.LeaveEndDate)
	layout.FirstDayOfWork = (*xsdDate)(layout.T.FirstDayOfWork)
	layout.LeaveLastDayofWork = (*xsdDate)(layout.T.LeaveLastDayofWork)
	return e.EncodeElement(layout, start)
}
func (t *LeaveStatusDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LeaveStatusDetailDataType
	var overlay struct {
		*T
		LeaveStartDate        *xsdDate `xml:"urn:com.workday/bsvc Leave_Start_Date,omitempty"`
		EstimatedLeaveEndDate *xsdDate `xml:"urn:com.workday/bsvc Estimated_Leave_End_Date,omitempty"`
		LeaveEndDate          *xsdDate `xml:"urn:com.workday/bsvc Leave_End_Date,omitempty"`
		FirstDayOfWork        *xsdDate `xml:"urn:com.workday/bsvc First_Day_Of_Work,omitempty"`
		LeaveLastDayofWork    *xsdDate `xml:"urn:com.workday/bsvc Leave_Last_Day_of_Work,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LeaveStartDate = (*xsdDate)(overlay.T.LeaveStartDate)
	overlay.EstimatedLeaveEndDate = (*xsdDate)(overlay.T.EstimatedLeaveEndDate)
	overlay.LeaveEndDate = (*xsdDate)(overlay.T.LeaveEndDate)
	overlay.FirstDayOfWork = (*xsdDate)(overlay.T.FirstDayOfWork)
	overlay.LeaveLastDayofWork = (*xsdDate)(overlay.T.LeaveLastDayofWork)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type LeaveTypeReasonObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LeaveTypeReasonObjectType struct {
	ID         []LeaveTypeReasonObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LeaveofAbsenceTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LeaveofAbsenceTypeObjectType struct {
	ID         []LeaveofAbsenceTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Wrapper for License Identification Data. Includes License Identifiers.
type LicenseIdentificationDataType struct {
	LicenseID  []LicenseIDType `xml:"urn:com.workday/bsvc License_ID,omitempty"`
	ReplaceAll bool            `xml:"urn:com.workday/bsvc Replace_All,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type LocalTerminationReasonObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LocalTerminationReasonObjectType struct {
	ID         []LocalTerminationReasonObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LocaleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LocaleObjectType struct {
	ID         []LocaleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Encapsulating element containing a brief summary of Location data.
type LocationSummaryDataType struct {
	LocationReference        *LocationObjectType          `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	Name                     string                       `xml:"urn:com.workday/bsvc Name,omitempty"`
	LocationTypeReference    []LocationTypeObjectType     `xml:"urn:com.workday/bsvc Location_Type_Reference,omitempty"`
	LocalReference           *LocaleObjectType            `xml:"urn:com.workday/bsvc Local_Reference,omitempty"`
	DisplayLanguageReference *UserLanguageObjectType      `xml:"urn:com.workday/bsvc Display_Language_Reference,omitempty"`
	TimeProfileReference     *TimeProfileObjectType       `xml:"urn:com.workday/bsvc Time_Profile_Reference,omitempty"`
	ScheduledWeeklyHours     float64                      `xml:"urn:com.workday/bsvc Scheduled_Weekly_Hours,omitempty"`
	AddressData              []AddressInformationDataType `xml:"urn:com.workday/bsvc Address_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LocationTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LocationTypeObjectType struct {
	ID         []LocationTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LossImpactObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LossImpactObjectType struct {
	ID         []LossImpactObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The details of the transaction that has been rescinded or corrected.
type MainTransactionLogEntryDataType struct {
	TransactionLogReference *TransactionLogObjectType `xml:"urn:com.workday/bsvc Transaction_Log_Reference,omitempty"`
	TransactionLogData      *TransactionLogDataType   `xml:"urn:com.workday/bsvc Transaction_Log_Data,omitempty"`
}

// Contains employee contracts for an employee.
type MaintainEmployeeContractsDataType struct {
	EmployeeReference    *EmployeeObjectType        `xml:"urn:com.workday/bsvc Employee_Reference"`
	EmployeeContractData []EmployeeContractDataType `xml:"urn:com.workday/bsvc Employee_Contract_Data"`
}

// Add or update employee contracts.
type MaintainEmployeeContractsRequestType struct {
	BusinessProcessParameters     *BusinessProcessParametersType     `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	MaintainEmployeeContractsData *MaintainEmployeeContractsDataType `xml:"urn:com.workday/bsvc Maintain_Employee_Contracts_Data"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Employee Contract that was added or updated.
type MaintainEmployeeContractsResponseType struct {
	EmployeeContractEventReference *UniqueIdentifierObjectType  `xml:"urn:com.workday/bsvc Employee_Contract_Event_Reference,omitempty"`
	EmployeeContractReference      []EmployeeContractObjectType `xml:"urn:com.workday/bsvc Employee_Contract_Reference,omitempty"`
	Version                        string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Maintain Employee Contracts sub business process.
type MaintainEmployeeContractsSubBusinessProcessType struct {
	BusinessSubProcessParameters  *BusinessSubProcessParametersType     `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	MaintainEmployeeContractsData *MaintainEmployeeContractsSubDataType `xml:"urn:com.workday/bsvc Maintain_Employee_Contracts_Data,omitempty"`
}

// Contains data for the Employee Contract Event.
type MaintainEmployeeContractsSubDataType struct {
	EmployeeContractData []EmployeeContractDataType `xml:"urn:com.workday/bsvc Employee_Contract_Data,omitempty"`
}

// Contains Maintain Probation Period for Country data
type MaintainProbationPeriodsForCountryDataType struct {
	CountryReference              *CountryObjectType                  `xml:"urn:com.workday/bsvc Country_Reference"`
	ProbationPeriodForCountryData []ProbationPeriodForCountryDataType `xml:"urn:com.workday/bsvc Probation_Period_For_Country_Data,omitempty"`
}

// Wrapper element for Education Data.
type ManageEducationDataforRoleType struct {
	RoleReference   *RoleObjectType                      `xml:"urn:com.workday/bsvc Role_Reference,omitempty"`
	SourceReference *PersonSkillSourceCategoryObjectType `xml:"urn:com.workday/bsvc Source_Reference,omitempty"`
	Education       []EducationType                      `xml:"urn:com.workday/bsvc Education,omitempty"`
}

// Manage Education Sub Business Process Data
type ManageEducationSubBusinessProcessDataType struct {
	BusinessSubProcessParameters []BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	ManageEducationData          []ManageEducationDataforRoleType   `xml:"urn:com.workday/bsvc Manage_Education_Data,omitempty"`
}

// Details of the Probation Period being added to the Worker
type ManageEmployeeProbationPeriodSubDataType struct {
	ProbationPeriodReference      *EmployeeProbationPeriodObjectType       `xml:"urn:com.workday/bsvc Probation_Period_Reference,omitempty"`
	StartDate                     time.Time                                `xml:"urn:com.workday/bsvc Start_Date"`
	EndDate                       *time.Time                               `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	ProbationTypeReference        *EmployeeProbationPeriodTypeObjectType   `xml:"urn:com.workday/bsvc Probation_Type_Reference,omitempty"`
	DeriveProbationPeriodfromRule *bool                                    `xml:"urn:com.workday/bsvc Derive_Probation_Period_from_Rule,omitempty"`
	Length                        *ProbationPeriodLengthDataType           `xml:"urn:com.workday/bsvc Length,omitempty"`
	ProbationReasonReference      *EmployeeProbationPeriodReasonObjectType `xml:"urn:com.workday/bsvc Probation_Reason_Reference,omitempty"`
	ExtendedEndDate               *time.Time                               `xml:"urn:com.workday/bsvc Extended_End_Date,omitempty"`
	Note                          string                                   `xml:"urn:com.workday/bsvc Note,omitempty"`
}

func (t *ManageEmployeeProbationPeriodSubDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ManageEmployeeProbationPeriodSubDataType
	var layout struct {
		*T
		StartDate       *xsdDate `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate         *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
		ExtendedEndDate *xsdDate `xml:"urn:com.workday/bsvc Extended_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(&layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	layout.ExtendedEndDate = (*xsdDate)(layout.T.ExtendedEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ManageEmployeeProbationPeriodSubDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ManageEmployeeProbationPeriodSubDataType
	var overlay struct {
		*T
		StartDate       *xsdDate `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate         *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
		ExtendedEndDate *xsdDate `xml:"urn:com.workday/bsvc Extended_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(&overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	overlay.ExtendedEndDate = (*xsdDate)(overlay.T.ExtendedEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Manage Instructor Eligibility Sub Business Process Data
type ManageInstructorEligibilitySubBusinessProcessDataType struct {
	BusinessSubProcessParameters  []BusinessSubProcessParametersType  `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	InstructorEligibilityLineData []InstructorEligibilityLineDataType `xml:"urn:com.workday/bsvc Instructor_Eligibility_Line_Data,omitempty"`
}

// Details of the Probation Period being added to the Worker
type ManageProbationPeriodSubBusinessProcessType struct {
	BusinessSubProcessParameters      *BusinessSubProcessParametersType         `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	ManageEmployeeProbationPeriodData *ManageEmployeeProbationPeriodSubDataType `xml:"urn:com.workday/bsvc Manage_Employee_Probation_Period_Data,omitempty"`
}

// Wrapper element for Professional Affiliation Data.
type ManageProfessionalAffiliationDataforRoleType struct {
	RoleReference           *RoleObjectType                      `xml:"urn:com.workday/bsvc Role_Reference,omitempty"`
	SourceReference         *PersonSkillSourceCategoryObjectType `xml:"urn:com.workday/bsvc Source_Reference,omitempty"`
	ProfessionalAffiliation []ProfessionalAffiliationSkillType   `xml:"urn:com.workday/bsvc Professional_Affiliation,omitempty"`
}

// Manage Professional Affiliation Sub Business Process Data
type ManageProfessionalAffiliationSubBusinessProcessDataType struct {
	BusinessSubProcessParameters      []BusinessSubProcessParametersType             `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	ManageProfessionalAffiliationData []ManageProfessionalAffiliationDataforRoleType `xml:"urn:com.workday/bsvc Manage_Professional_Affiliation_Data,omitempty"`
}

// Wrapper for the Manage Union Membership sub business process that is part of staffing event business processes.
type ManageUnionMembershipSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	UnionMemberData              *UnionMemberDataType              `xml:"urn:com.workday/bsvc Union_Member_Data,omitempty"`
}

// Contains the organization reference and details about an organization.
type ManagementChainDataType struct {
	OrganizationReference *OrganizationObjectType `xml:"urn:com.workday/bsvc Organization_Reference,omitempty"`
	ManagerReference      []WorkerObjectType      `xml:"urn:com.workday/bsvc Manager_Reference,omitempty"`
	Manager               []WorkerBasicDataType   `xml:"urn:com.workday/bsvc Manager,omitempty"`
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

// Contains the evaluation information that was filled out by the manager.
type ManagerEvaluationDetailDataType struct {
	OverallData *OverallEvaluationDetailDataType `xml:"urn:com.workday/bsvc Overall_Data,omitempty"`
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

// Contains the Position Swap details and the Change Job data
type MassSwapPositionsDataType struct {
	Name                          string                                 `xml:"urn:com.workday/bsvc Name"`
	Description                   string                                 `xml:"urn:com.workday/bsvc Description,omitempty"`
	EffectiveDate                 time.Time                              `xml:"urn:com.workday/bsvc Effective_Date"`
	ChangeJobReasonReference      *ChangeJobSubcategoryObjectType        `xml:"urn:com.workday/bsvc Change_Job_Reason_Reference"`
	ChangeJobDataforSwapPositions []ChangeJobDataforMassPositionSwapType `xml:"urn:com.workday/bsvc Change_Job_Data_for_Swap_Positions,omitempty"`
}

func (t *MassSwapPositionsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MassSwapPositionsDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *MassSwapPositionsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MassSwapPositionsDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type MatrixOrganizationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MatrixOrganizationObjectType struct {
	ID         []MatrixOrganizationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Maximum Probation Period Data
type MaximumProbationPeriodDataType struct {
	MaximumProbationDuration      float64                     `xml:"urn:com.workday/bsvc Maximum_Probation_Duration,omitempty"`
	MaximumProbationUnitReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Maximum_Probation_Unit_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type MeritPercentPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MeritPercentPlanObjectType struct {
	ID         []MeritPercentPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Wrapper element for all Military Service data for the worker.
type MilitaryInformationDataType struct {
	MilitaryServiceInformationData []MilitaryServiceInformationDataType `xml:"urn:com.workday/bsvc Military_Service_Information_Data,omitempty"`
	ReplaceAll                     bool                                 `xml:"urn:com.workday/bsvc Replace_All,attr,omitempty"`
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

// Wrapper element for each Military Service entry.
type MilitaryServiceInformationDataType struct {
	MilitaryServiceReference *MilitaryServiceReferenceObjectType `xml:"urn:com.workday/bsvc Military_Service_Reference,omitempty"`
	MilitaryServiceData      *MilitaryServiceSubDataType         `xml:"urn:com.workday/bsvc Military_Service_Data,omitempty"`
	Delete                   bool                                `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
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
type MobilityChoiceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MobilityChoiceObjectType struct {
	ID         []MobilityChoiceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the Move Workers (By Organization) web service.
type MoveWorkersByOrganizationDataType struct {
	EffectiveDate                         time.Time                                   `xml:"urn:com.workday/bsvc Effective_Date"`
	FromOrganizationReference             *StaffObjectType                            `xml:"urn:com.workday/bsvc From_Organization_Reference"`
	MoveWorkersByOrganizationPositionData []MoveWorkersByOrganizationPositionDataType `xml:"urn:com.workday/bsvc Move_Workers_By_Organization_Position_Data"`
	ToOrganizationReference               *StaffObjectType                            `xml:"urn:com.workday/bsvc To_Organization_Reference"`
}

func (t *MoveWorkersByOrganizationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MoveWorkersByOrganizationDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *MoveWorkersByOrganizationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MoveWorkersByOrganizationDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains the Move Workers (By Organization) position and worker reference.
type MoveWorkersByOrganizationPositionDataType struct {
	PositiontoMoveReference    *StaffingInterfaceObjectType `xml:"urn:com.workday/bsvc Position_to_Move_Reference"`
	WorkerforPositionReference *WorkerObjectType            `xml:"urn:com.workday/bsvc Worker_for_Position_Reference,omitempty"`
}

// Business process for performing Move Workers (By Organization).
type MoveWorkersByOrganizationRequestType struct {
	BusinessProcessParameters     *BusinessProcessParametersType     `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	MoveWorkersByOrganizationData *MoveWorkersByOrganizationDataType `xml:"urn:com.workday/bsvc Move_Workers_By_Organization_Data"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Event ID of the Move Workers (By Organization) event.
type MoveWorkersByOrganizationResponseType struct {
	MoveWorkersByOrganizationReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Move_Workers_By_Organization_Reference,omitempty"`
	Version                            string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type NamedProfessorshipObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type NamedProfessorshipObjectType struct {
	ID         []NamedProfessorshipObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Web service for a No Show Employee. This web service will run the No Show Business Process.
type NoShowEmployeeDataType struct {
	EmployeeReference   *EmployeeObjectType                 `xml:"urn:com.workday/bsvc Employee_Reference"`
	MarkNoShowReference *CommonBooleanEnumerationObjectType `xml:"urn:com.workday/bsvc Mark_No_Show_Reference,omitempty"`
}

// Responds with the Event ID for the No Show Event
type NoShowEventResponseType struct {
	EventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper for No Show Web Service.
type NoShowRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	NoShowEmployeeData        *NoShowEmployeeDataType        `xml:"urn:com.workday/bsvc No_Show_Employee_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Notice Period data specific to the Employee.
type NoticePeriodDataEmployeeType struct {
	NoticePeriodData []NoticePeriodDataType `xml:"urn:com.workday/bsvc Notice_Period_Data,omitempty"`
}

// Notice Period data for an employer.
type NoticePeriodDataEmployerType struct {
	NoticePeriodData []NoticePeriodDataType `xml:"urn:com.workday/bsvc Notice_Period_Data,omitempty"`
}

// Detail notice period data for employer and employee.
type NoticePeriodDataType struct {
	NoticePeriodReference           *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Notice_Period_Reference,omitempty"`
	NoticePeriodDuration            float64                     `xml:"urn:com.workday/bsvc Notice_Period_Duration,omitempty"`
	NoticePeriodUnitReference       *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Notice_Period_Unit_Reference,omitempty"`
	DateAndTimeAdjustmentReference  *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Date_And_Time_Adjustment_Reference,omitempty"`
	NoticePeriodIsCustom            *bool                       `xml:"urn:com.workday/bsvc Notice_Period_Is_Custom,omitempty"`
	OverriddenNoticePeriodReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Overridden_Notice_Period_Reference,omitempty"`
}

// Contains the employer/employee Notice Period response data specific to the Notice Period Target request and As Of Effective Date.
type NoticePeriodDataforNoticePeriodTargetType struct {
	NoticePeriodTargetReference *NoticePeriodTargetObjectType  `xml:"urn:com.workday/bsvc Notice_Period_Target_Reference,omitempty"`
	LocationReference           *LocationObjectType            `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	EmployerNoticePeriodData    []NoticePeriodDataEmployerType `xml:"urn:com.workday/bsvc Employer_Notice_Period_Data,omitempty"`
	EmployeeNoticePeriodData    []NoticePeriodDataEmployeeType `xml:"urn:com.workday/bsvc Employee_Notice_Period_Data,omitempty"`
}

// Notice Period Eligibility Rule
type NoticePeriodEligibilityRulePrev312Type struct {
	NoticePeriodEligibilityRuleReference *ConditionRuleObjectType   `xml:"urn:com.workday/bsvc Notice_Period_Eligibility_Rule_Reference,omitempty"`
	NoticePeriodEligibilityRuleData      []ConditionRuleDataWWSType `xml:"urn:com.workday/bsvc Notice_Period_Eligibility_Rule_Data,omitempty"`
}

// Notice Period Eligibility Rules Request References
type NoticePeriodEligibilityRuleRequestReferencesType struct {
	NoticePeriodEligibilityRuleReference []ConditionRuleObjectType `xml:"urn:com.workday/bsvc Notice_Period_Eligibility_Rule_Reference"`
}

// Notice Period Eligibility Rule Response Data
type NoticePeriodEligibilityRuleResponseDataType struct {
	NoticePeriodEligibilityRule []NoticePeriodEligibilityRulePrev312Type `xml:"urn:com.workday/bsvc Notice_Period_Eligibility_Rule,omitempty"`
}

// Adds a row to the table.
type NoticePeriodLineDataType struct {
	Order                    string                      `xml:"urn:com.workday/bsvc Order"`
	EligibilityRuleReference *ConditionRuleObjectType    `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	HasNoticePeriod          *bool                       `xml:"urn:com.workday/bsvc Has_Notice_Period,omitempty"`
	ForEmployer              *bool                       `xml:"urn:com.workday/bsvc For_Employer,omitempty"`
	ForEmployee              *bool                       `xml:"urn:com.workday/bsvc For__Employee_,omitempty"`
	Duration                 float64                     `xml:"urn:com.workday/bsvc Duration,omitempty"`
	UnitReference            *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Unit_Reference,omitempty"`
	AdjustmentReference      *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Adjustment_Reference,omitempty"`
	Inactive                 *bool                       `xml:"urn:com.workday/bsvc Inactive,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type NoticePeriodTargetObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type NoticePeriodTargetObjectType struct {
	ID         []NoticePeriodTargetObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Data element containing Notice Period Lines and Country Reference.
type NoticePeriodsForCountryDataType struct {
	CountryReference     *CountryObjectType         `xml:"urn:com.workday/bsvc Country_Reference"`
	NoticePeriodLineData []NoticePeriodLineDataType `xml:"urn:com.workday/bsvc Notice_Period_Line_Data,omitempty"`
}

// Data element for Get Maintain Notice Periods for Country web service
type NoticePeriodsForCountryGetDataType struct {
	MaintainNoticePeriodsForCountryReference *NoticePeriodsSetupTableObjectType `xml:"urn:com.workday/bsvc Maintain_Notice_Periods_For_Country_Reference,omitempty"`
	MaintainNoticePeriodsForCountryData      *NoticePeriodsForCountryDataType   `xml:"urn:com.workday/bsvc Maintain_Notice_Periods_For_Country_Data,omitempty"`
}

// Request Criteria
type NoticePeriodsForCountryRequestCriteriaType struct {
	CountryReference []CountryObjectType `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
}

// Use this element to request specific notice periods setup tables given reference ID values.
type NoticePeriodsForCountryRequestReferencesType struct {
	MaintainNoticePeriodsForCountryReference []NoticePeriodsSetupTableObjectType `xml:"urn:com.workday/bsvc Maintain_Notice_Periods_For_Country_Reference"`
}

// Response Data element for Get Maintain Notice Periods for Country web service.
type NoticePeriodsForCountryResponseDataType struct {
	MaintainNoticePeriodsForCountry []NoticePeriodsForCountryGetDataType `xml:"urn:com.workday/bsvc Maintain_Notice_Periods_For_Country,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type NoticePeriodsSetupTableObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type NoticePeriodsSetupTableObjectType struct {
	ID         []NoticePeriodsSetupTableObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type NotificationCategorizableObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type NotificationCategorizableObjectType struct {
	ID         []NotificationCategorizableObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Notification Type configurations for a particular user
type NotificationSubCategoryConfigurationsforUsersType struct {
	NotificationSubTypeConfiguration []NotificationSubTypeConfigurationType `xml:"urn:com.workday/bsvc Notification_Sub_Type_Configuration,omitempty"`
}

// Notification sub type configuration for each notification sub type
type NotificationSubTypeConfigurationType struct {
	NotificationSubTypeReference  *NotificationCategorizableObjectType `xml:"urn:com.workday/bsvc Notification_Sub_Type_Reference,omitempty"`
	NotificationchannelsReference []UniqueIdentifierObjectType         `xml:"urn:com.workday/bsvc Notification_channels_Reference,omitempty"`
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

// Encapsulating element containing all Emergency Contact data.
type OldEmergencyContactType struct {
	EmergencyContactReference *EmergencyContactObjectType `xml:"urn:com.workday/bsvc Emergency_Contact_Reference"`
	EmergencyContactData      *EmergencyContactDataType   `xml:"urn:com.workday/bsvc Emergency_Contact_Data"`
}

// Container for data used in Onboarding Setup sub business process
type OnboardingSetupDataType struct {
	OnboardingSetupTemplateReference *OnboardingSetupTemplateObjectType `xml:"urn:com.workday/bsvc Onboarding_Setup_Template_Reference,omitempty"`
}

// Wrapper for the Onboarding Setup sub business process that is part of the Hire business process.
type OnboardingSetupSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	OnboardingSetupData          *OnboardingSetupDataType          `xml:"urn:com.workday/bsvc Onboarding_Setup_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type OnboardingSetupTemplateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type OnboardingSetupTemplateObjectType struct {
	ID         []OnboardingSetupTemplateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Sub process for creating a one-time payment paid to the person who referred the hired worker.
type OneTimePaymentforReferralSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType         `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	EffectiveDate                *time.Time                                `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	ReasonReference              *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
}

func (t *OneTimePaymentforReferralSubBusinessProcessType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T OneTimePaymentforReferralSubBusinessProcessType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *OneTimePaymentforReferralSubBusinessProcessType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T OneTimePaymentforReferralSubBusinessProcessType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
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

// Contains a unique identifier for an instance of an object.
type OrganizationGoalObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type OrganizationGoalObjectType struct {
	ID         []OrganizationGoalObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains the organzation role information.
type OrganizationRoleAssignmentforWorkerDataType struct {
	RoleAssignerReference *RoleAssignerObjectType `xml:"urn:com.workday/bsvc Role_Assigner_Reference,omitempty"`
	EffectiveDate         *time.Time              `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	AssignmentFrom        string                  `xml:"urn:com.workday/bsvc Assignment_From,omitempty"`
}

func (t *OrganizationRoleAssignmentforWorkerDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T OrganizationRoleAssignmentforWorkerDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *OrganizationRoleAssignmentforWorkerDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T OrganizationRoleAssignmentforWorkerDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains the organization role and the organizations that the worker supports that role for.
type OrganizationRoleforWorkerDataType struct {
	OrganizationRoleReference *AssignableRoleObjectType                     `xml:"urn:com.workday/bsvc Organization_Role_Reference,omitempty"`
	OrganizationRoleData      []OrganizationRoleAssignmentforWorkerDataType `xml:"urn:com.workday/bsvc Organization_Role_Data,omitempty"`
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

// Contains the details about the organization.
type OrganizationSummaryDataType struct {
	OrganizationReferenceID             string                          `xml:"urn:com.workday/bsvc Organization_Reference_ID,omitempty"`
	OrganizationCode                    string                          `xml:"urn:com.workday/bsvc Organization_Code,omitempty"`
	IntegrationIDData                   []ExternalIntegrationIDDataType `xml:"urn:com.workday/bsvc Integration_ID_Data,omitempty"`
	OrganizationName                    string                          `xml:"urn:com.workday/bsvc Organization_Name,omitempty"`
	OrganizationTypeReference           *OrganizationTypeObjectType     `xml:"urn:com.workday/bsvc Organization_Type_Reference,omitempty"`
	OrganizationSubtypeReference        *OrganizationSubtypeObjectType  `xml:"urn:com.workday/bsvc Organization_Subtype_Reference,omitempty"`
	PrimaryBusinessSiteReference        *LocationObjectType             `xml:"urn:com.workday/bsvc Primary_Business_Site_Reference,omitempty"`
	OrganizationSupportRoleData         *OrganizationSupportRoleType    `xml:"urn:com.workday/bsvc Organization_Support_Role_Data,omitempty"`
	DateofPayGroupAssignment            *time.Time                      `xml:"urn:com.workday/bsvc Date_of_Pay_Group_Assignment,omitempty"`
	UsedinChangeOrganizationAssignments *bool                           `xml:"urn:com.workday/bsvc Used_in_Change_Organization_Assignments,omitempty"`
}

func (t *OrganizationSummaryDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T OrganizationSummaryDataType
	var layout struct {
		*T
		DateofPayGroupAssignment *xsdDate `xml:"urn:com.workday/bsvc Date_of_Pay_Group_Assignment,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateofPayGroupAssignment = (*xsdDate)(layout.T.DateofPayGroupAssignment)
	return e.EncodeElement(layout, start)
}
func (t *OrganizationSummaryDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T OrganizationSummaryDataType
	var overlay struct {
		*T
		DateofPayGroupAssignment *xsdDate `xml:"urn:com.workday/bsvc Date_of_Pay_Group_Assignment,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateofPayGroupAssignment = (*xsdDate)(overlay.T.DateofPayGroupAssignment)
	return d.DecodeElement(&overlay, &start)
}

// Contains the organization role refeference and the information about that role.
type OrganizationSupportRoleType struct {
	OrganizationSupportRole []OrganizationSupportingRoleDataType `xml:"urn:com.workday/bsvc Organization_Support_Role,omitempty"`
}

// The organizational roles which provide support for the worker.
type OrganizationSupportingRoleDataType struct {
	OrganizationRoleReference *AssignableRoleObjectType                  `xml:"urn:com.workday/bsvc Organization_Role_Reference,omitempty"`
	OrganizationRoleData      []WorkerOrganizationRoleAssignmentDataType `xml:"urn:com.workday/bsvc Organization_Role_Data,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type OtherUnitTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type OtherUnitTypeObjectType struct {
	ID         []OtherUnitTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The overall section evaluation information (rating, comments) for the evaluator.
type OverallEvaluationDetailDataType struct {
	RatingReference *ReviewRatingObjectType `xml:"urn:com.workday/bsvc Rating_Reference,omitempty"`
	Comment         string                  `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

// Adds or updates the unit and duration for a custom notice period.
type OverrideNoticePeriodDataType struct {
	Duration            float64                     `xml:"urn:com.workday/bsvc Duration,omitempty"`
	UnitReference       *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Unit_Reference"`
	AdjustmentReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Adjustment_Reference,omitempty"`
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

// Wrapper for Passports and Visas Identification Data. Includes Passport Identifiers and Visa Identifiers
type PassportsandVisasIdentificationDataType struct {
	PassportID []PassportIDType `xml:"urn:com.workday/bsvc Passport_ID,omitempty"`
	VisaID     []VisaIDType     `xml:"urn:com.workday/bsvc Visa_ID,omitempty"`
	ReplaceAll bool             `xml:"urn:com.workday/bsvc Replace_All,attr,omitempty"`
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

// This element reports Pay Group Assignment rescind events.
type PayGroupAssignmentCorrectorRescindedOrganizationDataType struct {
	EffectiveDate             *time.Time          `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	CompletedOn               *time.Time          `xml:"urn:com.workday/bsvc Completed_On,omitempty"`
	EventCorrected            *bool               `xml:"urn:com.workday/bsvc Event_Corrected,omitempty"`
	EventRescinded            *bool               `xml:"urn:com.workday/bsvc Event_Rescinded,omitempty"`
	OriginalPayGroupReference *PayGroupObjectType `xml:"urn:com.workday/bsvc Original_Pay_Group_Reference,omitempty"`
	ProposedPayGroupReference *PayGroupObjectType `xml:"urn:com.workday/bsvc Proposed_Pay_Group_Reference,omitempty"`
}

func (t *PayGroupAssignmentCorrectorRescindedOrganizationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayGroupAssignmentCorrectorRescindedOrganizationDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate     `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		CompletedOn   *xsdDateTime `xml:"urn:com.workday/bsvc Completed_On,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	layout.CompletedOn = (*xsdDateTime)(layout.T.CompletedOn)
	return e.EncodeElement(layout, start)
}
func (t *PayGroupAssignmentCorrectorRescindedOrganizationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayGroupAssignmentCorrectorRescindedOrganizationDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate     `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		CompletedOn   *xsdDateTime `xml:"urn:com.workday/bsvc Completed_On,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	overlay.CompletedOn = (*xsdDateTime)(overlay.T.CompletedOn)
	return d.DecodeElement(&overlay, &start)
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
type PayRateTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayRateTypeObjectType struct {
	ID         []PayRateTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PaymentTermsObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentTermsObjectType struct {
	ID         []PaymentTermsObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the payroll interface contribution amount for the retirement savings election.
type PayrollInterfaceContributionAmountDataType struct {
	ContributionAmount float64              `xml:"urn:com.workday/bsvc Contribution_Amount"`
	FrequencyReference *FrequencyObjectType `xml:"urn:com.workday/bsvc Frequency_Reference"`
}

// Contains the payroll interface contribution information for the spending account election.
type PayrollInterfaceContributionDataType struct {
	ContributionAmount float64              `xml:"urn:com.workday/bsvc Contribution_Amount,omitempty"`
	FrequencyReference *FrequencyObjectType `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type PayrollReportingTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollReportingTypeObjectType struct {
	ID         []PayrollReportingTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the latest completed performance improvement plan for the employee.
type PerformanceImprovementPlanDataType struct {
	PerformanceImprovementPlanReference *PerformanceImprovementPlanObjectType `xml:"urn:com.workday/bsvc Performance_Improvement_Plan_Reference,omitempty"`
	ReviewData                          *EmployeeReviewDetailsDataType        `xml:"urn:com.workday/bsvc Review_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PerformanceImprovementPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PerformanceImprovementPlanObjectType struct {
	ID         []PerformanceImprovementPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the latest completed performance review for the employee.
type PerformanceReviewDataType struct {
	PerformanceReviewReference *PerformanceReviewObjectType   `xml:"urn:com.workday/bsvc Performance_Review_Reference,omitempty"`
	ReviewData                 *EmployeeReviewDetailsDataType `xml:"urn:com.workday/bsvc Review_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PerformanceReviewObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PerformanceReviewObjectType struct {
	ID         []PerformanceReviewObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// This element contains all provisioning group assignments for a person.
type PersonAccountProvisioningDataType struct {
	ProvisioningGroupAssignmentData []ProvisioningGroupAssignmentforPersonDataType `xml:"urn:com.workday/bsvc Provisioning_Group_Assignment_Data,omitempty"`
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

// Person's Photo Information
type PersonPhotoDataType struct {
	Filename string `xml:"urn:com.workday/bsvc Filename,omitempty"`
	File     []byte `xml:"urn:com.workday/bsvc File"`
}

func (t *PersonPhotoDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PersonPhotoDataType
	var layout struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File"`
	}
	layout.T = (*T)(t)
	layout.File = (*xsdBase64Binary)(&layout.T.File)
	return e.EncodeElement(layout, start)
}
func (t *PersonPhotoDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PersonPhotoDataType
	var overlay struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File"`
	}
	overlay.T = (*T)(t)
	overlay.File = (*xsdBase64Binary)(&overlay.T.File)
	return d.DecodeElement(&overlay, &start)
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

// Contains a unique identifier for an instance of an object.
type PersonSkillSourceCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PersonSkillSourceCategoryObjectType struct {
	ID         []PersonSkillSourceCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Provides selection for filtering the response
type PersonTypeCriteriaType struct {
	IncludeAcademicAffiliates       *bool `xml:"urn:com.workday/bsvc Include_Academic_Affiliates,omitempty"`
	IncludeExternalCommitteeMembers *bool `xml:"urn:com.workday/bsvc Include_External_Committee_Members,omitempty"`
	IncludeExternalStudentRecords   *bool `xml:"urn:com.workday/bsvc Include_External_Student_Records,omitempty"`
	IncludeStudentProspectRecords   *bool `xml:"urn:com.workday/bsvc Include_Student_Prospect_Records,omitempty"`
	IncludeStudentRecords           *bool `xml:"urn:com.workday/bsvc Include_Student_Records,omitempty"`
	IncludeWorkers                  *bool `xml:"urn:com.workday/bsvc Include_Workers,omitempty"`
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

// Information about the Earliest and Earliest Contiguous Eligiblility Dates for Retirement Plans which are Part of a Plan Year
type PlanEligiblityDatesDataType struct {
	BenefitPlanReference              *BenefitPlanObjectType `xml:"urn:com.workday/bsvc Benefit_Plan_Reference,omitempty"`
	EarliestEligibilityDate           *time.Time             `xml:"urn:com.workday/bsvc Earliest_Eligibility_Date,omitempty"`
	EarliestContiguousEligibilityDate *time.Time             `xml:"urn:com.workday/bsvc Earliest_Contiguous_Eligibility_Date,omitempty"`
}

func (t *PlanEligiblityDatesDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PlanEligiblityDatesDataType
	var layout struct {
		*T
		EarliestEligibilityDate           *xsdDate `xml:"urn:com.workday/bsvc Earliest_Eligibility_Date,omitempty"`
		EarliestContiguousEligibilityDate *xsdDate `xml:"urn:com.workday/bsvc Earliest_Contiguous_Eligibility_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EarliestEligibilityDate = (*xsdDate)(layout.T.EarliestEligibilityDate)
	layout.EarliestContiguousEligibilityDate = (*xsdDate)(layout.T.EarliestContiguousEligibilityDate)
	return e.EncodeElement(layout, start)
}
func (t *PlanEligiblityDatesDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PlanEligiblityDatesDataType
	var overlay struct {
		*T
		EarliestEligibilityDate           *xsdDate `xml:"urn:com.workday/bsvc Earliest_Eligibility_Date,omitempty"`
		EarliestContiguousEligibilityDate *xsdDate `xml:"urn:com.workday/bsvc Earliest_Contiguous_Eligibility_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EarliestEligibilityDate = (*xsdDate)(overlay.T.EarliestEligibilityDate)
	overlay.EarliestContiguousEligibilityDate = (*xsdDate)(overlay.T.EarliestContiguousEligibilityDate)
	return d.DecodeElement(&overlay, &start)
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

// Encapsulating element containg all Payroll Interface Processing data.
type PositionDetailDataType struct {
	PositionReference                             *PositionElementObjectType                 `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	HeadcountReference                            *HeadcountRestrictionsObjectType           `xml:"urn:com.workday/bsvc Headcount_Reference,omitempty"`
	PositionID                                    string                                     `xml:"urn:com.workday/bsvc Position_ID,omitempty"`
	PositionTitle                                 string                                     `xml:"urn:com.workday/bsvc Position_Title,omitempty"`
	BusinessTitle                                 string                                     `xml:"urn:com.workday/bsvc Business_Title,omitempty"`
	StartDate                                     *time.Time                                 `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndEmploymentDate                             *time.Time                                 `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
	EndEmploymentReasonReference                  []EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc End_Employment_Reason_Reference,omitempty"`
	WorkerTypeReference                           *PositionWorkerTypeObjectType              `xml:"urn:com.workday/bsvc Worker_Type_Reference,omitempty"`
	PositionTimeTypeReference                     *PositionTimeTypeObjectType                `xml:"urn:com.workday/bsvc Position_Time_Type_Reference,omitempty"`
	JobExempt                                     *bool                                      `xml:"urn:com.workday/bsvc Job_Exempt,omitempty"`
	ScheduledWeeklyHours                          float64                                    `xml:"urn:com.workday/bsvc Scheduled_Weekly_Hours,omitempty"`
	DefaultWeeklyHours                            float64                                    `xml:"urn:com.workday/bsvc Default_Weekly_Hours,omitempty"`
	WorkingTimeFrequencyReference                 *FrequencyObjectType                       `xml:"urn:com.workday/bsvc Working_Time_Frequency_Reference,omitempty"`
	WorkingTimeUnitReference                      *WorkingTimeUnitObjectType                 `xml:"urn:com.workday/bsvc Working_Time_Unit_Reference,omitempty"`
	WorkingTimeValue                              float64                                    `xml:"urn:com.workday/bsvc Working_Time_Value,omitempty"`
	FullTimeEquivalentPercentage                  float64                                    `xml:"urn:com.workday/bsvc Full_Time_Equivalent_Percentage,omitempty"`
	ExcludefromHeadcount                          *bool                                      `xml:"urn:com.workday/bsvc Exclude_from_Headcount,omitempty"`
	PayRateTypeReference                          *PayRateTypeObjectType                     `xml:"urn:com.workday/bsvc Pay_Rate_Type_Reference,omitempty"`
	JobClassificationSummaryData                  []JobClassificationSummaryDataType         `xml:"urn:com.workday/bsvc Job_Classification_Summary_Data,omitempty"`
	CompanyInsiderReference                       []CompanyInsiderTypeObjectType             `xml:"urn:com.workday/bsvc Company_Insider_Reference,omitempty"`
	WorkShiftReference                            *WorkShiftObjectType                       `xml:"urn:com.workday/bsvc Work_Shift_Reference,omitempty"`
	WorkHoursProfilesReference                    *WorkHoursProfileObjectType                `xml:"urn:com.workday/bsvc Work_Hours_Profiles_Reference,omitempty"`
	FederalWithholdingFEIN                        string                                     `xml:"urn:com.workday/bsvc Federal_Withholding_FEIN,omitempty"`
	WorkerCompensationCodeData                    []WorkerCompensationCodeDataType           `xml:"urn:com.workday/bsvc Worker_Compensation_Code_Data,omitempty"`
	PositionPayrollReportingCodeData              []PositionPayrollReportingCodeDataType     `xml:"urn:com.workday/bsvc Position_Payroll_Reporting_Code_Data,omitempty"`
	JobProfileSummaryData                         *JobProfileinPositionSummaryDataType       `xml:"urn:com.workday/bsvc Job_Profile_Summary_Data,omitempty"`
	BusinessSiteSummaryData                       *LocationSummaryDataType                   `xml:"urn:com.workday/bsvc Business_Site_Summary_Data,omitempty"`
	PayrollInterfaceProcessingData                *PositionPayrollInterfaceDetailDataType    `xml:"urn:com.workday/bsvc Payroll_Interface_Processing_Data,omitempty"`
	RegularPaidEquivalentHours                    float64                                    `xml:"urn:com.workday/bsvc Regular_Paid_Equivalent_Hours,omitempty"`
	WorkerHoursProfileClassification              string                                     `xml:"urn:com.workday/bsvc Worker_Hours_Profile_Classification,omitempty"`
	InternationalAssignmentData                   *InternationalAssignmentDataType           `xml:"urn:com.workday/bsvc International_Assignment_Data,omitempty"`
	WorkSpaceReference                            *LocationObjectType                        `xml:"urn:com.workday/bsvc Work_Space__Reference,omitempty"`
	AcademicPaySetupData                          *AcademicPaySetupDataType                  `xml:"urn:com.workday/bsvc Academic_Pay_Setup_Data,omitempty"`
	EndDate                                       *time.Time                                 `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	PayThroughDate                                *time.Time                                 `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
	CollectiveAgreementSummaryData                []CollectiveAgreementSummaryDataType       `xml:"urn:com.workday/bsvc Collective_Agreement_Summary_Data,omitempty"`
	EmployeeProbationPeriodSummaryData            *EmployeeProbationPeriodSummaryDataType    `xml:"urn:com.workday/bsvc Employee_Probation_Period_Summary_Data,omitempty"`
	ManagerasoflastdetectedmanagerchangeReference []WorkerObjectType                         `xml:"urn:com.workday/bsvc Manager_as_of_last_detected_manager_change_Reference,omitempty"`
	EffectiveDate                                 time.Time                                  `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
}

func (t *PositionDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PositionDetailDataType
	var layout struct {
		*T
		StartDate         *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndEmploymentDate *xsdDate `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
		EndDate           *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
		PayThroughDate    *xsdDate `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
		EffectiveDate     *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(layout.T.StartDate)
	layout.EndEmploymentDate = (*xsdDate)(layout.T.EndEmploymentDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	layout.PayThroughDate = (*xsdDate)(layout.T.PayThroughDate)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PositionDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PositionDetailDataType
	var overlay struct {
		*T
		StartDate         *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndEmploymentDate *xsdDate `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
		EndDate           *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
		PayThroughDate    *xsdDate `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
		EffectiveDate     *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(overlay.T.StartDate)
	overlay.EndEmploymentDate = (*xsdDate)(overlay.T.EndEmploymentDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	overlay.PayThroughDate = (*xsdDate)(overlay.T.PayThroughDate)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
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

// Wrapper element for filled position data.
type PositionDetailsSubDataType struct {
	JobProfileReference                      *JobProfileObjectType              `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	PositionTitle                            string                             `xml:"urn:com.workday/bsvc Position_Title,omitempty"`
	BusinessTitle                            string                             `xml:"urn:com.workday/bsvc Business_Title,omitempty"`
	LocationReference                        *LocationObjectType                `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	WorkSpaceReference                       *LocationObjectType                `xml:"urn:com.workday/bsvc Work_Space_Reference,omitempty"`
	PositionTimeTypeReference                *PositionTimeTypeObjectType        `xml:"urn:com.workday/bsvc Position_Time_Type_Reference,omitempty"`
	WorkShiftReference                       *WorkShiftObjectType               `xml:"urn:com.workday/bsvc Work_Shift_Reference,omitempty"`
	WorkHoursProfileReference                *WorkHoursProfileObjectType        `xml:"urn:com.workday/bsvc Work_Hours_Profile_Reference,omitempty"`
	DefaultHours                             float64                            `xml:"urn:com.workday/bsvc Default_Hours,omitempty"`
	ScheduledHours                           float64                            `xml:"urn:com.workday/bsvc Scheduled_Hours,omitempty"`
	WorkingTimeFrequencyReference            *FrequencyObjectType               `xml:"urn:com.workday/bsvc Working_Time_Frequency_Reference,omitempty"`
	WorkingTimeUnitReference                 *WorkingTimeUnitObjectType         `xml:"urn:com.workday/bsvc Working_Time_Unit_Reference,omitempty"`
	WorkingTimeValue                         float64                            `xml:"urn:com.workday/bsvc Working_Time_Value,omitempty"`
	PayRateTypeReference                     *PayRateTypeObjectType             `xml:"urn:com.workday/bsvc Pay_Rate_Type_Reference,omitempty"`
	JobClassificationReference               []JobClassificationObjectType      `xml:"urn:com.workday/bsvc Job_Classification_Reference,omitempty"`
	CompanyInsiderTypeReference              []CompanyInsiderTypeObjectType     `xml:"urn:com.workday/bsvc Company_Insider_Type_Reference,omitempty"`
	AnnualWorkPeriodReference                *AcademicPayPeriodObjectType       `xml:"urn:com.workday/bsvc Annual_Work_Period_Reference,omitempty"`
	DisbursementPlanPeriodReference          *AcademicPayPeriodObjectType       `xml:"urn:com.workday/bsvc Disbursement_Plan_Period_Reference,omitempty"`
	WorkersCompensationCodeOverrideReference *WorkersCompensationCodeObjectType `xml:"urn:com.workday/bsvc Workers__Compensation_Code_Override_Reference,omitempty"`
	PositionExternalIDData                   *ExternalIDDataType                `xml:"urn:com.workday/bsvc Position_External_ID_Data,omitempty"`
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

// Wrapper element for the Freeze Position business process data including the position to close and the event data.
type PositionGroupFreezeDataType struct {
	PositionReference *PositionRestrictionsObjectType   `xml:"urn:com.workday/bsvc Position_Reference"`
	FreezeEventData   *PositionGroupFreezeEventDataType `xml:"urn:com.workday/bsvc Freeze_Event_Data"`
}

// Contains the detailed data for the freeze.
type PositionGroupFreezeEventDataType struct {
	ReasonReference *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
	FreezeDate      time.Time                                 `xml:"urn:com.workday/bsvc Freeze_Date"`
	FrozenPosition  *bool                                     `xml:"urn:com.workday/bsvc Frozen_Position,omitempty"`
}

func (t *PositionGroupFreezeEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PositionGroupFreezeEventDataType
	var layout struct {
		*T
		FreezeDate *xsdDate `xml:"urn:com.workday/bsvc Freeze_Date"`
	}
	layout.T = (*T)(t)
	layout.FreezeDate = (*xsdDate)(&layout.T.FreezeDate)
	return e.EncodeElement(layout, start)
}
func (t *PositionGroupFreezeEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PositionGroupFreezeEventDataType
	var overlay struct {
		*T
		FreezeDate *xsdDate `xml:"urn:com.workday/bsvc Freeze_Date"`
	}
	overlay.T = (*T)(t)
	overlay.FreezeDate = (*xsdDate)(&overlay.T.FreezeDate)
	return d.DecodeElement(&overlay, &start)
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

// Contains the worker's organizations that they are a member of related to the position.
type PositionManagementChainsDataType struct {
	PositionSupervisoryManagementChainData *WorkerSupervisoryManagementChainDataType `xml:"urn:com.workday/bsvc Position_Supervisory_Management_Chain_Data,omitempty"`
	PositionMatrixManagementChainData      *WorkerMatrixManagementChainDataType      `xml:"urn:com.workday/bsvc Position_Matrix_Management_Chain_Data,omitempty"`
}

// Contains matrix organization assignment for each position
type PositionMatrixOrganizationDataType struct {
	PositionReference                    *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference"`
	PositionMatrixOrganizationDetailData *PositionMatrixOrganizationDetailDataType `xml:"urn:com.workday/bsvc Position_Matrix_Organization_Detail_Data"`
}

// Effective date and Matrix Organization for Matrix Organization Assignment
type PositionMatrixOrganizationDetailDataType struct {
	EffectiveDate               time.Time                     `xml:"urn:com.workday/bsvc Effective_Date"`
	MatrixOrganizationReference *MatrixOrganizationObjectType `xml:"urn:com.workday/bsvc Matrix_Organization_Reference"`
}

func (t *PositionMatrixOrganizationDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PositionMatrixOrganizationDetailDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PositionMatrixOrganizationDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PositionMatrixOrganizationDetailDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
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

// Contains the position's organization that it is part of
type PositionOrganizationDataType struct {
	PositionOrganizationData               []WorkerOrganizationMembershipDataType                     `xml:"urn:com.workday/bsvc Position_Organization_Data,omitempty"`
	PayGroupAssignmentCorrectorRescindData []PayGroupAssignmentCorrectorRescindedOrganizationDataType `xml:"urn:com.workday/bsvc Pay_Group_Assignment_Correct_or_Rescind_Data,omitempty"`
}

// Encapsulating element containg all Payroll Interface Processing data.
type PositionPayrollInterfaceDetailDataType struct {
	EffectiveDate                 *time.Time                             `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	PayRateTypeReference          *PayRateTypeObjectType                 `xml:"urn:com.workday/bsvc Pay_Rate_Type_Reference,omitempty"`
	FrequencyReference            *FrequencyObjectType                   `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	PayGroupReference             *ExternalPayGroupObjectType            `xml:"urn:com.workday/bsvc Pay_Group_Reference,omitempty"`
	PayrollEntityReference        *ExternalPayrollEntityObjectType       `xml:"urn:com.workday/bsvc Payroll_Entity_Reference,omitempty"`
	ExternalEmployeeTypeReference *ExternalPayrollEmployeeTypeObjectType `xml:"urn:com.workday/bsvc External_Employee_Type_Reference,omitempty"`
	PayrollFileNumber             string                                 `xml:"urn:com.workday/bsvc Payroll_File_Number,omitempty"`
}

func (t *PositionPayrollInterfaceDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PositionPayrollInterfaceDetailDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PositionPayrollInterfaceDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PositionPayrollInterfaceDetailDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains each payroll reporting code for the position.
type PositionPayrollReportingCodeDataType struct {
	PayrollReportingCodeReference *PayrollReportingCodeAllObjectType `xml:"urn:com.workday/bsvc Payroll_Reporting_Code_Reference"`
	Code                          string                             `xml:"urn:com.workday/bsvc Code"`
	FormattedCode                 string                             `xml:"urn:com.workday/bsvc Formatted_Code,omitempty"`
	Name                          string                             `xml:"urn:com.workday/bsvc Name"`
	PayrollReportingTypeReference *PayrollReportingTypeObjectType    `xml:"urn:com.workday/bsvc Payroll_Reporting_Type_Reference"`
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
type PositionSetObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PositionSetObjectType struct {
	ID         []PositionSetObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains each Position based on the Request References or Request Criteria.  The data returned is the current information as of the dates in the response filter, and does not include all historical information about the  position.
type PositionsResponseDataType struct {
	Position []PositionRestrictionType `xml:"urn:com.workday/bsvc Position,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PotentialObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PotentialObjectType struct {
	ID         []PotentialObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type ProbationPeriodActionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProbationPeriodActionObjectType struct {
	ID         []ProbationPeriodActionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Probation Period Eligibility Rule Data Element
type ProbationPeriodEligibilityRuleDataType struct {
	ConditionRuleData *ConditionRuleDataWWSType `xml:"urn:com.workday/bsvc Condition_Rule_Data,omitempty"`
}

// Probation Period Eligibility Rule Get Data Element
type ProbationPeriodEligibilityRuleGetDataType struct {
	ProbationPeriodEligibilityRuleReference *ConditionRuleObjectType                 `xml:"urn:com.workday/bsvc Probation_Period_Eligibility_Rule_Reference,omitempty"`
	ProbationPeriodEligibilityRuleData      []ProbationPeriodEligibilityRuleDataType `xml:"urn:com.workday/bsvc Probation_Period_Eligibility_Rule_Data,omitempty"`
}

// Probation Period Eligibility Rule Request References Element
type ProbationPeriodEligibilityRuleRequestReferencesType struct {
	ProbationPeriodEligibilityRuleReference []ConditionRuleObjectType `xml:"urn:com.workday/bsvc Probation_Period_Eligibility_Rule_Reference,omitempty"`
}

// Probation Period Eligibility Rule Response Data
type ProbationPeriodEligibilityRuleResponseDataType struct {
	ProbationPeriodEligibilityRule []ProbationPeriodEligibilityRuleGetDataType `xml:"urn:com.workday/bsvc Probation_Period_Eligibility_Rule,omitempty"`
}

// Probation Period For Country Data
type ProbationPeriodForCountryDataType struct {
	Order                                   string                           `xml:"urn:com.workday/bsvc Order,omitempty"`
	ProbationPeriodEligibilityRuleReference *ConditionRuleObjectType         `xml:"urn:com.workday/bsvc Probation_Period_Eligibility_Rule_Reference"`
	ProbationPeriodDuration                 float64                          `xml:"urn:com.workday/bsvc Probation_Period_Duration,omitempty"`
	ProbationPeriodUnitReference            *UniqueIdentifierObjectType      `xml:"urn:com.workday/bsvc Probation_Period_Unit_Reference,omitempty"`
	MaximumProbationPeriodData              []MaximumProbationPeriodDataType `xml:"urn:com.workday/bsvc Maximum_Probation_Period_Data,omitempty"`
	ReviewProbationPeriodData               []ReviewProbationPeriodDataType  `xml:"urn:com.workday/bsvc Review_Probation_Period_Data,omitempty"`
	Inactive                                *bool                            `xml:"urn:com.workday/bsvc Inactive,omitempty"`
}

// Probation Period Length
type ProbationPeriodLengthDataType struct {
	Duration      float64                     `xml:"urn:com.workday/bsvc Duration"`
	UnitReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Unit_Reference"`
}

// Adds a probation period outcome row.
type ProbationPeriodOutcomeDataType struct {
	ProbationPeriodOutcomeReference *ProbationPeriodOutcomeObjectType `xml:"urn:com.workday/bsvc Probation_Period_Outcome_Reference,omitempty"`
	ProbationPeriodOutcomeName      string                            `xml:"urn:com.workday/bsvc Probation_Period_Outcome_Name"`
	ProbationPeriodActionsReference []ProbationPeriodActionObjectType `xml:"urn:com.workday/bsvc Probation_Period_Actions_Reference"`
	Inactive                        *bool                             `xml:"urn:com.workday/bsvc Inactive,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProbationPeriodOutcomeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProbationPeriodOutcomeObjectType struct {
	ID         []ProbationPeriodOutcomeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Response Data element for Get Probation Period Outcomes web service.
type ProbationPeriodOutcomeResponseDataType struct {
	ProbationPeriodOutcomesForCountryReference *ProbationPeriodOutcomesSetupTableObjectType `xml:"urn:com.workday/bsvc Probation_Period_Outcomes_For_Country_Reference,omitempty"`
	ProbationPeriodOutcomesData                []ProbationPeriodOutcomesForCountryDataType  `xml:"urn:com.workday/bsvc Probation_Period_Outcomes_Data,omitempty"`
}

// Data element containing Probation Period Outcomes and Country Reference.
type ProbationPeriodOutcomesForCountryDataType struct {
	CountryReference           *CountryObjectType               `xml:"urn:com.workday/bsvc Country_Reference"`
	ProbationPeriodOutcomeData []ProbationPeriodOutcomeDataType `xml:"urn:com.workday/bsvc Probation_Period_Outcome_Data,omitempty"`
}

// Request element for Country References
type ProbationPeriodOutcomesForCountryRequestCriteriaType struct {
	CountryReference []CountryObjectType `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
}

// Request element for Probation Period Outcome References
type ProbationPeriodOutcomesForCountryRequestReferencesType struct {
	ProbationPeriodOutcomesForCountryReference []ProbationPeriodOutcomesSetupTableObjectType `xml:"urn:com.workday/bsvc Probation_Period_Outcomes_For_Country_Reference"`
}

// Probation Period Outcomes For Country Response Data Element
type ProbationPeriodOutcomesForCountryResponseDataType struct {
	ProbationPeriodOutcomes []ProbationPeriodOutcomeResponseDataType `xml:"urn:com.workday/bsvc Probation_Period_Outcomes,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProbationPeriodOutcomesSetupTableObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProbationPeriodOutcomesSetupTableObjectType struct {
	ID         []ProbationPeriodOutcomesSetupTableObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProbationPeriodRuleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProbationPeriodRuleObjectType struct {
	ID         []ProbationPeriodRuleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProbationPeriodRulesTableObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProbationPeriodRulesTableObjectType struct {
	ID         []ProbationPeriodRulesTableObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Probation Periods For Country Get Data element.
type ProbationPeriodsForCountryGetDataType struct {
	MaintainProbationPeriodsForCountryReference *ProbationPeriodRulesTableObjectType        `xml:"urn:com.workday/bsvc Maintain_Probation_Periods_For_Country_Reference,omitempty"`
	MaintainProbationPeriodsForCountryData      *MaintainProbationPeriodsForCountryDataType `xml:"urn:com.workday/bsvc Maintain_Probation_Periods_For_Country_Data,omitempty"`
}

// Probation Periods For Country Request Criteria
type ProbationPeriodsForCountryRequestCriteriaType struct {
	CountryReference []CountryObjectType `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
}

// Probation Periods For Country References
type ProbationPeriodsForCountryRequestReferencesType struct {
	MaintainProbationPeriodsForCountryReference []ProbationPeriodRulesTableObjectType `xml:"urn:com.workday/bsvc Maintain_Probation_Periods_For_Country_Reference"`
}

// Probation Periods For Country Response Data element.
type ProbationPeriodsForCountryResponseDataType struct {
	MaintainProbationPeriodsForCountry []ProbationPeriodsForCountryGetDataType `xml:"urn:com.workday/bsvc Maintain_Probation_Periods_For_Country,omitempty"`
}

// Contains the probation period data for a worker
type ProbationPeriodsForWorkerDataType struct {
	WorkerReference                          *WorkerObjectType                        `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	PositionReference                        *PositionElementObjectType               `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	StartDate                                *time.Time                               `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                                  *time.Time                               `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	ExtendedEndDate                          *time.Time                               `xml:"urn:com.workday/bsvc Extended_End_Date,omitempty"`
	EmployeeProbationPeriodTypeReference     *EmployeeProbationPeriodTypeObjectType   `xml:"urn:com.workday/bsvc Employee_Probation_Period_Type_Reference,omitempty"`
	ProbationPeriodReasonReference           *EmployeeProbationPeriodReasonObjectType `xml:"urn:com.workday/bsvc Probation_Period_Reason_Reference,omitempty"`
	Note                                     string                                   `xml:"urn:com.workday/bsvc Note,omitempty"`
	CustomProbationPeriodDuration            float64                                  `xml:"urn:com.workday/bsvc Custom_Probation_Period_Duration,omitempty"`
	CustomProbationPeriodUnitReference       *UniqueIdentifierObjectType              `xml:"urn:com.workday/bsvc Custom_Probation_Period_Unit_Reference,omitempty"`
	ExtendedProbationPeriodDuration          float64                                  `xml:"urn:com.workday/bsvc Extended_Probation_Period_Duration,omitempty"`
	ExtendedProbationPeriodUnitReference     *UniqueIdentifierObjectType              `xml:"urn:com.workday/bsvc Extended_Probation_Period_Unit_Reference,omitempty"`
	CustomReviewProbationPeriodDate          *time.Time                               `xml:"urn:com.workday/bsvc Custom_Review_Probation_Period_Date,omitempty"`
	CustomReviewProbationPeriodDuration      float64                                  `xml:"urn:com.workday/bsvc Custom_Review_Probation_Period_Duration,omitempty"`
	CustomReviewProbationPeriodUnitReference *UniqueIdentifierObjectType              `xml:"urn:com.workday/bsvc Custom_Review_Probation_Period_Unit_Reference,omitempty"`
}

func (t *ProbationPeriodsForWorkerDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProbationPeriodsForWorkerDataType
	var layout struct {
		*T
		StartDate                       *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate                         *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
		ExtendedEndDate                 *xsdDate `xml:"urn:com.workday/bsvc Extended_End_Date,omitempty"`
		CustomReviewProbationPeriodDate *xsdDate `xml:"urn:com.workday/bsvc Custom_Review_Probation_Period_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	layout.ExtendedEndDate = (*xsdDate)(layout.T.ExtendedEndDate)
	layout.CustomReviewProbationPeriodDate = (*xsdDate)(layout.T.CustomReviewProbationPeriodDate)
	return e.EncodeElement(layout, start)
}
func (t *ProbationPeriodsForWorkerDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProbationPeriodsForWorkerDataType
	var overlay struct {
		*T
		StartDate                       *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate                         *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
		ExtendedEndDate                 *xsdDate `xml:"urn:com.workday/bsvc Extended_End_Date,omitempty"`
		CustomReviewProbationPeriodDate *xsdDate `xml:"urn:com.workday/bsvc Custom_Review_Probation_Period_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	overlay.ExtendedEndDate = (*xsdDate)(overlay.T.ExtendedEndDate)
	overlay.CustomReviewProbationPeriodDate = (*xsdDate)(overlay.T.CustomReviewProbationPeriodDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a single reference to a worker
type ProbationPeriodsForWorkersRequestReferencesType struct {
	WorkersReference []WorkerObjectType `xml:"urn:com.workday/bsvc Workers_Reference,omitempty"`
}

// Contains Probation Periods For Worker
type ProbationPeriodsForWorkersType struct {
	ProbationPeriodsForWorkerReference *EmployeeProbationPeriodObjectType  `xml:"urn:com.workday/bsvc Probation_Periods_For_Worker_Reference,omitempty"`
	ProbationPeriodForWorkerData       []ProbationPeriodsForWorkerDataType `xml:"urn:com.workday/bsvc Probation_Period_For_Worker_Data,omitempty"`
}

type ProcessingFaultType struct {
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
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

// Wrapper element for the Propose Compensation for Transfer sub business process.  If any errors are found during processing, the Auto Complete boolean will be set to False and manual processing will occur for this business process.
type ProposeCompensationForPositionSubBusinessProcessType struct {
	BusinessSubProcessParameters       *BusinessSubProcessParametersType        `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	ProposeCompensationforPositionData *CompensationProposedForPositionDataType `xml:"urn:com.workday/bsvc Propose_Compensation_for_Position_Data,omitempty"`
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

// This element contains all data defining the assignment of the person to a specific provisioning group at the current moment. It does not include the assignment's history.
type ProvisioningGroupAssignmentforPersonDataType struct {
	ProvisioningGroup string     `xml:"urn:com.workday/bsvc Provisioning_Group,omitempty"`
	Status            string     `xml:"urn:com.workday/bsvc Status,omitempty"`
	LastChanged       *time.Time `xml:"urn:com.workday/bsvc Last_Changed,omitempty"`
}

func (t *ProvisioningGroupAssignmentforPersonDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProvisioningGroupAssignmentforPersonDataType
	var layout struct {
		*T
		LastChanged *xsdDateTime `xml:"urn:com.workday/bsvc Last_Changed,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LastChanged = (*xsdDateTime)(layout.T.LastChanged)
	return e.EncodeElement(layout, start)
}
func (t *ProvisioningGroupAssignmentforPersonDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProvisioningGroupAssignmentforPersonDataType
	var overlay struct {
		*T
		LastChanged *xsdDateTime `xml:"urn:com.workday/bsvc Last_Changed,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LastChanged = (*xsdDateTime)(overlay.T.LastChanged)
	return d.DecodeElement(&overlay, &start)
}

// This element is used to group all data needed to create a purchase order if purchase orders are used.
type PurchaseOrderContractDetailsDataType struct {
	CompanyforPurchaseOrderReference *CompanyObjectType                   `xml:"urn:com.workday/bsvc Company_for_Purchase_Order_Reference,omitempty"`
	RequesterReference               *WorkerObjectType                    `xml:"urn:com.workday/bsvc Requester_Reference,omitempty"`
	SpendCategoryReference           *SpendCategoryObjectType             `xml:"urn:com.workday/bsvc Spend_Category_Reference,omitempty"`
	ContractAmount                   float64                              `xml:"urn:com.workday/bsvc Contract_Amount,omitempty"`
	WorktagsReference                []AuditedAccountingWorktagObjectType `xml:"urn:com.workday/bsvc Worktags_Reference,omitempty"`
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

// Assigns a matrix organization to a worker position
type PutAssignMatrixOrganizationRequestType struct {
	BusinessProcessParameters                   *BusinessProcessParametersType                   `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AssignMatrixOrganizationBusinessProcessData *AssignMatrixOrganizationBusinessProcessDataType `xml:"urn:com.workday/bsvc Assign_Matrix_Organization_Business_Process_Data"`
	Version                                     string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response for assign worker matrix organization request
type PutAssignMatrixOrganizationResponseType struct {
	MatrixOrganizationEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Matrix_Organization_Event_Reference,omitempty"`
	Version                          string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for the Put Citizenship Status operation
type PutCitizenshipStatusRequestType struct {
	CitizenshipStatusReference *CitizenshipStatusObjectType `xml:"urn:com.workday/bsvc Citizenship_Status_Reference,omitempty"`
	CitizenshipStatusData      *CitizenshipStatusDataType   `xml:"urn:com.workday/bsvc Citizenship_Status_Data,omitempty"`
	AddOnly                    bool                         `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for the Put Citizenship Status operation.
type PutCitizenshipStatusResponseType struct {
	CitizenshipStatusReference *CitizenshipStatusObjectType `xml:"urn:com.workday/bsvc Citizenship_Status_Reference,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Contains data for the Edit Notice Periods event.
type PutEditNoticePeriodsEventDataType struct {
	EmployeeReference        *EmployeeObjectType           `xml:"urn:com.workday/bsvc Employee_Reference"`
	EffectiveDate            time.Time                     `xml:"urn:com.workday/bsvc Effective_Date"`
	EmployerNoticePeriodData *EmployerNoticePeriodDataType `xml:"urn:com.workday/bsvc Employer_Notice_Period_Data,omitempty"`
	EmployeeNoticePeriodData *EmployeeNoticePeriodDataType `xml:"urn:com.workday/bsvc Employee_Notice_Period_Data,omitempty"`
}

func (t *PutEditNoticePeriodsEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PutEditNoticePeriodsEventDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PutEditNoticePeriodsEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PutEditNoticePeriodsEventDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Request element for Put Edit Notice Periods Event.
type PutEditNoticePeriodsEventRequestType struct {
	BusinessProcessParameters     *BusinessProcessParametersType     `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	PutEditNoticePeriodsEventData *PutEditNoticePeriodsEventDataType `xml:"urn:com.workday/bsvc Put_Edit_Notice_Periods_Event_Data"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Put Edit Notice Periods Event.
type PutEditNoticePeriodsEventResponseType struct {
	EditNoticePeriodsEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Edit_Notice_Periods_Event_Reference,omitempty"`
	Version                         string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This web service imports a Worker ID which will override the Workday generated ID.
type PutHireEventProposedWorkerIDRequestType struct {
	EventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference"`
	WorkerID       string                      `xml:"urn:com.workday/bsvc Worker_ID"`
	Version        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Event ID for the Hire Event.
type PutHireEventProposedWorkerIDResponseType struct {
	HireEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Hire_Event_Reference,omitempty"`
	WorkerID           string                      `xml:"urn:com.workday/bsvc Worker_ID,omitempty"`
	Version            string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Import Process Response element
type PutImportProcessResponseType struct {
	ImportProcessReference  *WebServiceBackgroundProcessRuntimeObjectType `xml:"urn:com.workday/bsvc Import_Process_Reference,omitempty"`
	HeaderInstanceReference *InstanceObjectType                           `xml:"urn:com.workday/bsvc Header_Instance_Reference,omitempty"`
	Version                 string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request Element for Put Job Classification Group
type PutJobClassificationGroupRequestType struct {
	JobClassificationGroupReference *JobClassificationGroupObjectType `xml:"urn:com.workday/bsvc Job_Classification_Group_Reference,omitempty"`
	JobClassificationGroupData      *JobClassificationGroupDataType   `xml:"urn:com.workday/bsvc Job_Classification_Group_Data,omitempty"`
	AddOnly                         bool                              `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                         string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Put Job Classification
type PutJobClassificationGroupResponseType struct {
	JobClassificationGroupReference *JobClassificationGroupObjectType `xml:"urn:com.workday/bsvc Job_Classification_Group_Reference,omitempty"`
	Version                         string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for  Put Job Family Group
type PutJobFamilyGroupRequestType struct {
	JobFamilyGroupReference *JobFamilyObjectType    `xml:"urn:com.workday/bsvc Job_Family_Group_Reference,omitempty"`
	JobFamilyGroupData      *JobFamilyGroupDataType `xml:"urn:com.workday/bsvc Job_Family_Group_Data,omitempty"`
	AddOnly                 bool                    `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                 string                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for the Put Job Family Group
type PutJobFamilyGroupResponseType struct {
	JobFamilyGroupsReference *JobFamilyObjectType `xml:"urn:com.workday/bsvc Job_Family_Groups_Reference,omitempty"`
	Version                  string               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for Put Job Family
type PutJobFamilyRequestType struct {
	JobFamilyReference *JobFamilyObjectType `xml:"urn:com.workday/bsvc Job_Family_Reference,omitempty"`
	JobFamilyData      *JobFamilyDataType   `xml:"urn:com.workday/bsvc Job_Family_Data,omitempty"`
	AddOnly            bool                 `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version            string               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Put Job Family
type PutJobFamilyResponseType struct {
	JobFamilyReference *JobFamilyObjectType `xml:"urn:com.workday/bsvc Job_Family_Reference,omitempty"`
	Version            string               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for Put Maintain Notice Periods for Country.
type PutMaintainNoticePeriodsforCountryRequestType struct {
	MaintainNoticePeriodsForCountryReference *NoticePeriodsSetupTableObjectType `xml:"urn:com.workday/bsvc Maintain_Notice_Periods_For_Country_Reference,omitempty"`
	MaintainNoticePeriodsForCountryData      []NoticePeriodsForCountryDataType  `xml:"urn:com.workday/bsvc Maintain_Notice_Periods_For_Country_Data,omitempty"`
	AddOnly                                  bool                               `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                                  string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Put Maintain Notice Periods for Country.
type PutMaintainNoticePeriodsforCountryResponseType struct {
	CountryReference                         []CountryObjectType                `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	MaintainNoticePeriodsForCountryReference *NoticePeriodsSetupTableObjectType `xml:"urn:com.workday/bsvc Maintain_Notice_Periods_For_Country_Reference,omitempty"`
	Version                                  string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for Put Maintain Probation Periods for Country
type PutMaintainProbationPeriodsForCountryRequestType struct {
	MaintainProbationPeriodsForCountryReference *ProbationPeriodRulesTableObjectType        `xml:"urn:com.workday/bsvc Maintain_Probation_Periods_For_Country_Reference,omitempty"`
	MaintainProbationPeriodsForCountryData      *MaintainProbationPeriodsForCountryDataType `xml:"urn:com.workday/bsvc Maintain_Probation_Periods_For_Country_Data"`
	AddOnly                                     bool                                        `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                                     string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Put Maintain Probation Periods for Country.
type PutMaintainProbationPeriodsForCountryResponseType struct {
	ProbationPeriodForCountryReference []ProbationPeriodRuleObjectType                `xml:"urn:com.workday/bsvc Probation_Period_For_Country_Reference,omitempty"`
	ExceptionsResponseData             []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                            string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for Put Notice Period Eligibility Rule.
type PutNoticePeriodEligibilityRuleRequestType struct {
	NoticePeriodEligibilityRuleReference *ConditionRuleObjectType  `xml:"urn:com.workday/bsvc Notice_Period_Eligibility_Rule_Reference,omitempty"`
	NoticePeriodEligibilityRuleData      *ConditionRuleDataWWSType `xml:"urn:com.workday/bsvc Notice_Period_Eligibility_Rule_Data,omitempty"`
	AddOnly                              bool                      `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Delete                               bool                      `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	Version                              string                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Put Notice Period Eligibility Rule.
type PutNoticePeriodEligibilityRuleResponseType struct {
	NoticePeriodEligibilityRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Notice_Period_Eligibility_Rule_Reference,omitempty"`
	Version                              string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Probation Period Eligibility Rule Request Element
type PutProbationPeriodEligibilityRuleRequestType struct {
	ProbationPeriodEligibilityRuleReference *ConditionRuleObjectType                `xml:"urn:com.workday/bsvc Probation_Period_Eligibility_Rule_Reference,omitempty"`
	ProbationPeriodEligibilityRuleData      *ProbationPeriodEligibilityRuleDataType `xml:"urn:com.workday/bsvc Probation_Period_Eligibility_Rule_Data,omitempty"`
	AddOnly                                 bool                                    `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Delete                                  bool                                    `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	Version                                 string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Probation Period Eligibility Rule Response Element
type PutProbationPeriodEligibilityRuleResponseType struct {
	ProbationPeriodEligibilityRuleReference  *ConditionRuleObjectType                       `xml:"urn:com.workday/bsvc Probation_Period_Eligibility_Rule_Reference,omitempty"`
	ApplicationInstanceRelatedExceptionsData []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Application_Instance_Related_Exceptions_Data,omitempty"`
	Version                                  string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for Put Probation Period Outcomes For Country
type PutProbationPeriodOutcomesRequestType struct {
	ProbationPeriodOutcomesForCountryReference *ProbationPeriodOutcomesSetupTableObjectType `xml:"urn:com.workday/bsvc Probation_Period_Outcomes_For_Country_Reference,omitempty"`
	ProbationPeriodOutcomesData                *ProbationPeriodOutcomesForCountryDataType   `xml:"urn:com.workday/bsvc Probation_Period_Outcomes_Data,omitempty"`
	AddOnly                                    bool                                         `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                                    string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Put Probation Period Outcomes
type PutProbationPeriodOutcomesResponseType struct {
	ProbationPeriodOutcomeReference            []ProbationPeriodOutcomeObjectType           `xml:"urn:com.workday/bsvc Probation_Period_Outcome_Reference,omitempty"`
	CountryReference                           *CountryObjectType                           `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	ProbationPeriodOutcomesForCountryReference *ProbationPeriodOutcomesSetupTableObjectType `xml:"urn:com.workday/bsvc Probation_Period_Outcomes_For_Country_Reference,omitempty"`
	Version                                    string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Removes a matrix organization from a worker position
type PutRemoveMatrixOrganizationRequestType struct {
	BusinessProcessParameters                   *BusinessProcessParametersType                   `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	RemoveMatrixOrganizationBusinessProcessData *RemoveMatrixOrganizationBusinessProcessDataType `xml:"urn:com.workday/bsvc Remove_Matrix_Organization_Business_Process_Data"`
	Version                                     string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response for remove worker matrix organization request
type PutRemoveMatrixOrganizationResponseType struct {
	RemovefromMatrixOrganizationEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Remove_from_Matrix_Organization_Event_Reference,omitempty"`
	Version                                    string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Creates a Student Employment Eligibility Rule to be used in a Student Employment Eligibility Rule Set
type PutStudentEmploymentEligibilityRuleRequestType struct {
	StudentEmploymentEligibilityRuleReference *ConditionRuleObjectType                  `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Reference,omitempty"`
	StudentEmploymentEligibilityRuleData      *StudentEmploymentEligibilityRuleDataType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Data"`
	AddOnly                                   bool                                      `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                                   string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Container for the Student Employment Eligibility Rule
type PutStudentEmploymentEligibilityRuleResponseType struct {
	StudentEmploymentEligibilityRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Reference,omitempty"`
	Version                                   string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Container for Put Student Employment Eligibility Rule Set Request
type PutStudentEmploymentEligibilityRuleSetRequestType struct {
	StudentEmploymentEligibilityRuleSetReference *StudentEmploymentEligibilityRuleSetObjectType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Set_Reference,omitempty"`
	StudentEmploymentEligibilityRuleSetData      *StudentEmploymentEligibilityRuleSetDataType   `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Set_Data"`
	AddOnly                                      bool                                           `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                                      string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Container for Student Employment Eligibility Rule Set Data
type PutStudentEmploymentEligibilityRuleSetResponseType struct {
	StudentEmploymentEligibilityRuleSetReference *StudentEmploymentEligibilityRuleSetObjectType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Set_Reference,omitempty"`
	StudentEmploymentEligibilityRuleSetData      []StudentEmploymentEligibilityRuleSetDataType  `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Set_Data,omitempty"`
	Version                                      string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Update request for Student Employment Eligibility Event
type PutStudentEmploymentEligibilityStatusRequestType struct {
	StudentEmploymentEligibilityEventReference *StudentEmploymentEligibilityEventObjectType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Event_Reference"`
	StudentEmploymentEligibilityEventData      *StudentEmploymentEligibilityEventDataType   `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Event_Data,omitempty"`
	Version                                    string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Student Emploment Eligibility Reference Updated
type PutStudentEmploymentEligibilityStatusResponseType struct {
	StudentEmploymentEligibilityEventReference *StudentEmploymentEligibilityEventObjectType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Event_Reference,omitempty"`
	Version                                    string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element for Put Worker Document/
type PutWorkerDocumentRequestType struct {
	WorkerDocumentReference *WorkerDocumentObjectType `xml:"urn:com.workday/bsvc Worker_Document_Reference,omitempty"`
	WorkerDocumentData      *WorkerDocumentDataType   `xml:"urn:com.workday/bsvc Worker_Document_Data"`
	AddOnly                 bool                      `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                 string                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Put Worker Document.
type PutWorkerDocumentResponseType struct {
	WorkerDocumentReference *WorkerDocumentObjectType `xml:"urn:com.workday/bsvc Worker_Document_Reference,omitempty"`
	Version                 string                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type QuestionnaireObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type QuestionnaireObjectType struct {
	ID         []QuestionnaireObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type RelatedPersonRelationshipObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RelatedPersonRelationshipObjectType struct {
	ID         []RelatedPersonRelationshipObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element containing all Related Person data.
type RelatedPersonType struct {
	RelatedPersonRelationshipReference []RelatedPersonRelationshipObjectType `xml:"urn:com.workday/bsvc Related_Person_Relationship_Reference,omitempty"`
	PersonReference                    *UniqueIdentifierObjectType           `xml:"urn:com.workday/bsvc Person_Reference,omitempty"`
	PersonalData                       *PersonalInformationDataType          `xml:"urn:com.workday/bsvc Personal_Data,omitempty"`
	Dependent                          *DependentType                        `xml:"urn:com.workday/bsvc Dependent,omitempty"`
	Beneficiary                        *BeneficiaryType                      `xml:"urn:com.workday/bsvc Beneficiary,omitempty"`
	EmergencyContact                   *OldEmergencyContactType              `xml:"urn:com.workday/bsvc Emergency_Contact,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type RelationalOperatorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RelationalOperatorObjectType struct {
	ID         []RelationalOperatorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type RelocationAreaObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RelocationAreaObjectType struct {
	ID         []RelocationAreaObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the information about the employee's willingness to relocate as part of their job.
type RelocationPreferenceDataType struct {
	ShortTermRelocationReference     *CommonBooleanEnumerationObjectType `xml:"urn:com.workday/bsvc Short_Term_Relocation_Reference,omitempty"`
	ShortTermRelocationAreaReference []RelocationAreaObjectType          `xml:"urn:com.workday/bsvc Short_Term_Relocation_Area__Reference,omitempty"`
	LongTermRelocationReference      *CommonBooleanEnumerationObjectType `xml:"urn:com.workday/bsvc Long_Term_Relocation_Reference,omitempty"`
	LongTermRelocationAreaReference  []RelocationAreaObjectType          `xml:"urn:com.workday/bsvc Long_Term_Relocation_Area__Reference,omitempty"`
	AdditionalInformation            string                              `xml:"urn:com.workday/bsvc Additional_Information,omitempty"`
}

// Specific compensation plan assigned to the worker to be removed.
type RemoveCompensationPlanAssignmentDataType struct {
	CompensationPlanReference []CompensationPlanObjectType `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
}

// Worker, position, and matrix organization data
type RemoveMatrixOrganizationBusinessProcessDataType struct {
	WorkerReference                      *WorkerObjectType                         `xml:"urn:com.workday/bsvc Worker_Reference"`
	RemoveMatrixOrganizationPositionData *RemoveMatrixOrganizationPositionDataType `xml:"urn:com.workday/bsvc Remove_Matrix_Organization_Position_Data"`
}

// Effective date and Matrix Organization for Matrix Organization Removal
type RemoveMatrixOrganizationDetailDataType struct {
	EffectiveDate               time.Time                     `xml:"urn:com.workday/bsvc Effective_Date"`
	MatrixOrganizationReference *MatrixOrganizationObjectType `xml:"urn:com.workday/bsvc Matrix_Organization_Reference"`
}

func (t *RemoveMatrixOrganizationDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RemoveMatrixOrganizationDetailDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *RemoveMatrixOrganizationDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RemoveMatrixOrganizationDetailDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains matrix organization assignment for each position
type RemoveMatrixOrganizationPositionDataType struct {
	PositionReference                  *PositionElementObjectType              `xml:"urn:com.workday/bsvc Position_Reference"`
	RemoveMatrixOrganizationDetailData *RemoveMatrixOrganizationDetailDataType `xml:"urn:com.workday/bsvc Remove_Matrix_Organization_Detail_Data"`
}

// Wrapper Element for the Remove Retiree Status business process web service.
type RemoveRetireeStatusDataType struct {
	EmployeeReference            *EmployeeObjectType               `xml:"urn:com.workday/bsvc Employee_Reference"`
	EndRetirementDate            time.Time                         `xml:"urn:com.workday/bsvc End_Retirement_Date"`
	RemoveRetireeStatusEventData *RemoveRetireeStatusEventDataType `xml:"urn:com.workday/bsvc Remove_Retiree_Status_Event_Data"`
}

func (t *RemoveRetireeStatusDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RemoveRetireeStatusDataType
	var layout struct {
		*T
		EndRetirementDate *xsdDate `xml:"urn:com.workday/bsvc End_Retirement_Date"`
	}
	layout.T = (*T)(t)
	layout.EndRetirementDate = (*xsdDate)(&layout.T.EndRetirementDate)
	return e.EncodeElement(layout, start)
}
func (t *RemoveRetireeStatusDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RemoveRetireeStatusDataType
	var overlay struct {
		*T
		EndRetirementDate *xsdDate `xml:"urn:com.workday/bsvc End_Retirement_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EndRetirementDate = (*xsdDate)(&overlay.T.EndRetirementDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains reason for removing retirement status and retiree organization that the worker is in.
type RemoveRetireeStatusEventDataType struct {
	EndRetirementReasonReference *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc End_Retirement_Reason_Reference"`
}

// This web service operation is designed to remove the retiree designation from a retiree using the Remove Retiree Status business process.
type RemoveRetireeStatusRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters"`
	RemoveRetireeStatusData   *RemoveRetireeStatusDataType   `xml:"urn:com.workday/bsvc Remove_Retiree_Status_Data,omitempty"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This element contains information on Remove Retiree Status event.
type RemoveRetireeStatusResponseType struct {
	EventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element that contains details for invoking Remove Retiree Status as sub process.
type RemoveRetireeStatusSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType         `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	RemoveRetireeStatusDetails   *RemoveRetireeStatusSubProcessDetailsType `xml:"urn:com.workday/bsvc Remove_Retiree_Status_Details,omitempty"`
}

// Contains details for Remove Retiree Status as sub process.
type RemoveRetireeStatusSubProcessDetailsType struct {
	EndRetirementDate time.Time                                 `xml:"urn:com.workday/bsvc End_Retirement_Date"`
	ReasonReference   *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference"`
}

func (t *RemoveRetireeStatusSubProcessDetailsType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RemoveRetireeStatusSubProcessDetailsType
	var layout struct {
		*T
		EndRetirementDate *xsdDate `xml:"urn:com.workday/bsvc End_Retirement_Date"`
	}
	layout.T = (*T)(t)
	layout.EndRetirementDate = (*xsdDate)(&layout.T.EndRetirementDate)
	return e.EncodeElement(layout, start)
}
func (t *RemoveRetireeStatusSubProcessDetailsType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RemoveRetireeStatusSubProcessDetailsType
	var overlay struct {
		*T
		EndRetirementDate *xsdDate `xml:"urn:com.workday/bsvc End_Retirement_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EndRetirementDate = (*xsdDate)(&overlay.T.EndRetirementDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for the Request Default Compensation sub business process.
type RequestCompensationDefaultSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	CompensationDefaultData      *CompensationDefaultDataType      `xml:"urn:com.workday/bsvc Compensation_Default_Data,omitempty"`
}

// Wrapper element for the Request Compensation change from the Edit Position business process.
type RequestCompensationForEditPositionSubBusinessProcessType struct {
	BusinessSubProcessParameters  *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	RequestCompensationChangeData *CompensationChangeDataType       `xml:"urn:com.workday/bsvc Request_Compensation_Change_Data,omitempty"`
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

// Data element for the stock grant request
type RequestStockGrantSubdataType struct {
	StockPlanReference       *StockPlanObjectType                      `xml:"urn:com.workday/bsvc Stock_Plan_Reference"`
	ReasonReference          *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
	IndividualStockGrantData []IndividualStockGrantDataType            `xml:"urn:com.workday/bsvc Individual_Stock_Grant_Data"`
}

// Wrapper element for the Request Stock Grant sub business process.
type RequestStockSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters"`
	RequestStockGrantData        *RequestStockGrantSubdataType     `xml:"urn:com.workday/bsvc Request_Stock_Grant_Data,omitempty"`
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
type RetentionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RetentionObjectType struct {
	ID         []RetentionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RetireeOrganizationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RetireeOrganizationObjectType struct {
	ID         []RetireeOrganizationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the retirement savings election information for the benefit plan year period.
type RetirementSavingsDataType struct {
	RetirementSavingsPeriodData []WorkerRetirementSavingsPeriodDataType `xml:"urn:com.workday/bsvc Retirement_Savings_Period_Data,omitempty"`
}

// Contains COBRA Eligibility detail for a participant.  If this element is not populated then the participant is assumed to not be COBRA eligible.
type ReviewCOBRAEligibilityDataType struct {
	COBRAEligible        *bool                             `xml:"urn:com.workday/bsvc COBRA_Eligible,omitempty"`
	ReasonReference      *COBRAEligibilityReasonObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
	COBRAEligibleDate    *time.Time                        `xml:"urn:com.workday/bsvc COBRA_Eligible_Date,omitempty"`
	QualifyingEventDate  *time.Time                        `xml:"urn:com.workday/bsvc Qualifying_Event_Date,omitempty"`
	CoverageEndDate      *time.Time                        `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	ParticipantReference []COBRAParticipantObjectType      `xml:"urn:com.workday/bsvc Participant_Reference,omitempty"`
	BenefitPlanReference *BenefitPlanObjectType            `xml:"urn:com.workday/bsvc Benefit_Plan_Reference,omitempty"`
}

func (t *ReviewCOBRAEligibilityDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ReviewCOBRAEligibilityDataType
	var layout struct {
		*T
		COBRAEligibleDate   *xsdDate `xml:"urn:com.workday/bsvc COBRA_Eligible_Date,omitempty"`
		QualifyingEventDate *xsdDate `xml:"urn:com.workday/bsvc Qualifying_Event_Date,omitempty"`
		CoverageEndDate     *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.COBRAEligibleDate = (*xsdDate)(layout.T.COBRAEligibleDate)
	layout.QualifyingEventDate = (*xsdDate)(layout.T.QualifyingEventDate)
	layout.CoverageEndDate = (*xsdDate)(layout.T.CoverageEndDate)
	return e.EncodeElement(layout, start)
}
func (t *ReviewCOBRAEligibilityDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ReviewCOBRAEligibilityDataType
	var overlay struct {
		*T
		COBRAEligibleDate   *xsdDate `xml:"urn:com.workday/bsvc COBRA_Eligible_Date,omitempty"`
		QualifyingEventDate *xsdDate `xml:"urn:com.workday/bsvc Qualifying_Event_Date,omitempty"`
		CoverageEndDate     *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.COBRAEligibleDate = (*xsdDate)(overlay.T.COBRAEligibleDate)
	overlay.QualifyingEventDate = (*xsdDate)(overlay.T.QualifyingEventDate)
	overlay.CoverageEndDate = (*xsdDate)(overlay.T.CoverageEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Enter COBRA information for a United States employee who is terminated or requesting a leave of absence.  Uses the Review COBRA Eligibility business process. The subprocess can be skipped, processed automatically or processed manually from the web service.
type ReviewCOBRAEligibilitySubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	ReviewCOBRAEligibilityData   []ReviewCOBRAEligibilityDataType  `xml:"urn:com.workday/bsvc Review_COBRA_Eligibility_Data,omitempty"`
}

// Launch the payroll integration.
// Uses the Review Payroll Data Integration Extract business process.
// This subprocess can be skipped, processed manually or processed automatically from the web service.
type ReviewPayrollInterfaceDataExtractBusinessSubProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
}

// Wrapper for the Edit Position data for Payroll Interface sub business process.
type ReviewPayrollInterfaceDataSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	ReviewPayrollInterfaceData   *ReviewPayrollInterfaceDataType   `xml:"urn:com.workday/bsvc Review_Payroll_Interface_Data,omitempty"`
}

// Wrapper element for Review Payroll Interface business process event via a web service.
type ReviewPayrollInterfaceDataType struct {
	ExternalPayGroupReference             *ExternalPayGroupObjectType                              `xml:"urn:com.workday/bsvc External_Pay_Group_Reference,omitempty"`
	PayFrequencyReference                 *FrequencyObjectType                                     `xml:"urn:com.workday/bsvc Pay_Frequency_Reference,omitempty"`
	ExternalPayrollEmployeeTypeReference  *ExternalPayrollEmployeeTypeObjectType                   `xml:"urn:com.workday/bsvc External_Payroll_Employee_Type_Reference,omitempty"`
	PayrollFileNumber                     string                                                   `xml:"urn:com.workday/bsvc Payroll_File_Number,omitempty"`
	ReviewPayrollInterfaceEventSubProcess *ReviewPayrollInterfaceDataExtractBusinessSubProcessType `xml:"urn:com.workday/bsvc Review_Payroll_Interface_Event_Sub_Process,omitempty"`
}

// Review Probation Period Data
type ReviewProbationPeriodDataType struct {
	ReviewProbationDuration      float64                     `xml:"urn:com.workday/bsvc Review_Probation_Duration,omitempty"`
	ReviewProbationUnitReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Review_Probation_Unit_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ReviewRatingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ReviewRatingObjectType struct {
	ID         []ReviewRatingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ReviewTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ReviewTypeObjectType struct {
	ID         []ReviewTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Role Assigner Reference and Data.
type RoleAssignerDataType struct {
	EffectiveTimezoneReference *TimeZoneObjectType      `xml:"urn:com.workday/bsvc Effective_Timezone_Reference,omitempty"`
	RoleAssignerReference      *RoleAssignerObjectType  `xml:"urn:com.workday/bsvc Role_Assigner_Reference"`
	RoleAssignmentData         []RoleAssignmentDataType `xml:"urn:com.workday/bsvc Role_Assignment_Data,omitempty"`
	EffectiveDate              time.Time                `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
}

func (t *RoleAssignerDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RoleAssignerDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *RoleAssignerDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RoleAssignerDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
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

// Role Assigner References.
type RoleAssignerRequestReferencesType struct {
	RoleAssignerReference []RoleAssignerObjectType `xml:"urn:com.workday/bsvc Role_Assigner_Reference"`
}

// Role Assignments for Role Assigner References.
type RoleAssignerResponseDataType struct {
	RoleAssigner []RoleAssignerType `xml:"urn:com.workday/bsvc Role_Assigner,omitempty"`
}

// Parameters for Response Group for Role Assigner References.
type RoleAssignerResponseGroupType struct {
	IncludeReference           *bool               `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	EffectiveTimezoneReference *TimeZoneObjectType `xml:"urn:com.workday/bsvc Effective_Timezone_Reference,omitempty"`
}

// Role Assigner Reference and Data.
type RoleAssignerType struct {
	RoleAssignerReference *RoleAssignerObjectType `xml:"urn:com.workday/bsvc Role_Assigner_Reference,omitempty"`
	RoleAssignerData      *RoleAssignerDataType   `xml:"urn:com.workday/bsvc Role_Assigner_Data,omitempty"`
}

// Role Assignment Data
type RoleAssignmentDataType struct {
	Delete                    *bool                     `xml:"urn:com.workday/bsvc Delete,omitempty"`
	Update                    *bool                     `xml:"urn:com.workday/bsvc Update,omitempty"`
	OrganizationRoleReference *AssignableRoleObjectType `xml:"urn:com.workday/bsvc Organization_Role_Reference"`
	RoleAssigneeReference     []RoleeObjectType         `xml:"urn:com.workday/bsvc Role_Assignee_Reference,omitempty"`
}

// Encapsulating Element for all Organization Role Assignment data.
type RoleAssignmentType struct {
	RoleAssignmentReference *UniqueIdentifierObjectType         `xml:"urn:com.workday/bsvc Role_Assignment_Reference,omitempty"`
	RoleAssignmentData      *OrganizationRoleAssignmentDataType `xml:"urn:com.workday/bsvc Role_Assignment_Data"`
}

// Request Criteria for Role Assignments of Role Assigner References.
type RoleAssignmentsRequestCriteriaType struct {
	AssignableRoleReference []AssignableRoleObjectType `xml:"urn:com.workday/bsvc Assignable_Role_Reference,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type RoleeSelectorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RoleeSelectorObjectType struct {
	ID         []RoleeSelectorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// The secondary reasons for the worker's most recent termination.
type SecondaryTerminationDataType struct {
	SecondaryTerminationReasonReference         *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Secondary_Termination_Reason_Reference,omitempty"`
	SecondaryTerminationReasonCategoryReference *EventClassificationCategoryObjectType    `xml:"urn:com.workday/bsvc Secondary_Termination_Reason_Category_Reference,omitempty"`
}

// Contains the evaluation information that was filled out by the employee.
type SelfEvaluationDetailDataType struct {
	OverallData *OverallEvaluationDetailDataType `xml:"urn:com.workday/bsvc Overall_Data,omitempty"`
}

// Wrapper element for setting hiring restrictions
type SetHiringRestrictionsDataType struct {
	SupervisoryOrganizationReference         *SupervisoryOrganizationObjectType                `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference"`
	HiringRestrictionsRequestReasonReference *EventClassificationSubcategoryObjectType         `xml:"urn:com.workday/bsvc Hiring_Restrictions_Request_Reason_Reference,omitempty"`
	PositionGroupRestrictionsData            *PositionGroupRestrictionDataType                 `xml:"urn:com.workday/bsvc Position_Group_Restrictions_Data"`
	RequestDefaultCompensationSubProcess     *RequestCompensationDefaultSubBusinessProcessType `xml:"urn:com.workday/bsvc Request_Default_Compensation_Sub_Process,omitempty"`
}

// Wrapper element for set hiring restrictions web service
type SetHiringRestrictionsRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	SetHiringRestrictionsData *SetHiringRestrictionsDataType `xml:"urn:com.workday/bsvc Set_Hiring_Restrictions_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the event id for the Job Restrictions Request Event, and the Supervisory Organization Reference.
type SetHiringRestrictionsResponseType struct {
	EventReference                   *UniqueIdentifierObjectType        `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	SupervisoryOrganizationReference *SupervisoryOrganizationObjectType `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference,omitempty"`
	Version                          string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

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
type SkillItemCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SkillItemCategoryObjectType struct {
	ID         []SkillItemCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains all the information about the Skill Item.
type SkillItemDataType struct {
	CategoryReference *SkillItemCategoryObjectType `xml:"urn:com.workday/bsvc Category_Reference,omitempty"`
	ID                string                       `xml:"urn:com.workday/bsvc ID,omitempty"`
	Name              string                       `xml:"urn:com.workday/bsvc Name,attr,omitempty"`
	Inactive          bool                         `xml:"urn:com.workday/bsvc Inactive,attr,omitempty"`
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

// Wrapper element for Skill Items.
type SkillType struct {
	SkillReference *SkillItemTenantedObjectType `xml:"urn:com.workday/bsvc Skill_Reference,omitempty"`
	SkillData      *SkillItemDataType           `xml:"urn:com.workday/bsvc Skill_Data,omitempty"`
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

// Contains the spending account coverage information (elections) for an employee.
type SpendingAccountCoverageDataType struct {
	BenefitElectionData      *WorkerBenefitElectionDataType `xml:"urn:com.workday/bsvc Benefit_Election_Data,omitempty"`
	EmployeeContributionData *EmployeeContributionDataType  `xml:"urn:com.workday/bsvc Employee_Contribution_Data,omitempty"`
	EmployerContributionData *EmployerContributionDataType  `xml:"urn:com.workday/bsvc Employer_Contribution_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StaffObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StaffObjectType struct {
	ID         []StaffObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StaffingInterfaceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StaffingInterfaceObjectType struct {
	ID         []StaffingInterfaceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper Element for the Start International Assignment business process web service and its sub business process
type StartInternationalAssignmentDataType struct {
	EmployeeReference                             *EmployeeObjectType                                     `xml:"urn:com.workday/bsvc Employee_Reference"`
	OrganizationReference                         *SupervisoryOrganizationObjectType                      `xml:"urn:com.workday/bsvc Organization_Reference"`
	PositionReference                             *PositionRestrictionsObjectType                         `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	JobRequisitionReference                       *JobRequisitionObjectType                               `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	StartInternationalAssignmentEventData         *StartInternationalAssignmentEventDataType              `xml:"urn:com.workday/bsvc Start_International_Assignment_Event_Data"`
	EditAssignOrganizationSubProcess              *EditAssignPositionOrganizationSubBusinessProcessType   `xml:"urn:com.workday/bsvc Edit_Assign_Organization_Sub_Process,omitempty"`
	AssignMatrixOrganizationSubProcess            *AssignMatrixOrganizationSubBusinessProcessType         `xml:"urn:com.workday/bsvc Assign_Matrix_Organization_Sub_Process,omitempty"`
	AssignPayGroupSubProcess                      *AssignPayGroupSubBusinessProcessType                   `xml:"urn:com.workday/bsvc Assign_Pay_Group_Sub_Process,omitempty"`
	ProposeCompensationforAdditionalJobSubProcess *ProposeCompensationForEmploymentSubBusinessProcessType `xml:"urn:com.workday/bsvc Propose_Compensation_for_Additional_Job_Sub_Process,omitempty"`
	MaintainEmployeeContractsSubBusinessProcess   *MaintainEmployeeContractsSubBusinessProcessType        `xml:"urn:com.workday/bsvc Maintain_Employee_Contracts_Sub_Business_Process,omitempty"`
	UpdateIDInformationSubProcess                 *UpdateIDInformationSubBusinessProcessType              `xml:"urn:com.workday/bsvc Update_ID_Information_Sub_Process,omitempty"`
	EditGovernmentIDsSubProcess                   *EditGovernmentIDsSubBusinessProcessType                `xml:"urn:com.workday/bsvc Edit_Government_IDs_Sub_Process,omitempty"`
	EditPassportsandVisasSubProcess               *EditPassportsandVisasSubBusinessProcessType            `xml:"urn:com.workday/bsvc Edit_Passports_and_Visas_Sub_Process,omitempty"`
	EditLicenseSubProcess                         *EditLicensesSubBusinessProcessType                     `xml:"urn:com.workday/bsvc Edit_License_Sub_Process,omitempty"`
	EditCustomIDsSubProcess                       *EditCustomIDsSubBusinessProcessType                    `xml:"urn:com.workday/bsvc Edit_Custom_IDs_Sub_Process,omitempty"`
	CheckPositionBudgetSubProcess                 *CheckPositionBudgetSubBusinessProcessType              `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process,omitempty"`
	OnboardingSetupSubProcess                     *OnboardingSetupSubBusinessProcessType                  `xml:"urn:com.workday/bsvc Onboarding_Setup_Sub_Process,omitempty"`
	AssignCostingAllocationSubProcess             *AssignCostingAllocationSubBusinessProcessType          `xml:"urn:com.workday/bsvc Assign_Costing_Allocation_Sub_Process,omitempty"`
	StartDate                                     time.Time                                               `xml:"urn:com.workday/bsvc Start_Date,attr,omitempty"`
}

func (t *StartInternationalAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T StartInternationalAssignmentDataType
	var layout struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(&layout.T.StartDate)
	return e.EncodeElement(layout, start)
}
func (t *StartInternationalAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T StartInternationalAssignmentDataType
	var overlay struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(&overlay.T.StartDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element that contains international assignment job related information
type StartInternationalAssignmentEventDataType struct {
	PositionID                                  string                                    `xml:"urn:com.workday/bsvc Position_ID,omitempty"`
	InternationalAssignmentTypeReference        *InternationalAssignmentTypeObjectType    `xml:"urn:com.workday/bsvc International_Assignment_Type_Reference"`
	ExpectedAssignmentEndDate                   *time.Time                                `xml:"urn:com.workday/bsvc Expected_Assignment_End_Date,omitempty"`
	StartInternationalAssignmentReasonReference *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Start_International_Assignment_Reason_Reference"`
	EmployeeTypeReference                       *PositionWorkerTypeObjectType             `xml:"urn:com.workday/bsvc Employee_Type_Reference"`
	FirstDayofWork                              *time.Time                                `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
	EndEmploymentDate                           *time.Time                                `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
	PositionDetails                             *PositionDetailsSubDataType               `xml:"urn:com.workday/bsvc Position_Details,omitempty"`
}

func (t *StartInternationalAssignmentEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T StartInternationalAssignmentEventDataType
	var layout struct {
		*T
		ExpectedAssignmentEndDate *xsdDate `xml:"urn:com.workday/bsvc Expected_Assignment_End_Date,omitempty"`
		FirstDayofWork            *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
		EndEmploymentDate         *xsdDate `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ExpectedAssignmentEndDate = (*xsdDate)(layout.T.ExpectedAssignmentEndDate)
	layout.FirstDayofWork = (*xsdDate)(layout.T.FirstDayofWork)
	layout.EndEmploymentDate = (*xsdDate)(layout.T.EndEmploymentDate)
	return e.EncodeElement(layout, start)
}
func (t *StartInternationalAssignmentEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T StartInternationalAssignmentEventDataType
	var overlay struct {
		*T
		ExpectedAssignmentEndDate *xsdDate `xml:"urn:com.workday/bsvc Expected_Assignment_End_Date,omitempty"`
		FirstDayofWork            *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
		EndEmploymentDate         *xsdDate `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ExpectedAssignmentEndDate = (*xsdDate)(overlay.T.ExpectedAssignmentEndDate)
	overlay.FirstDayofWork = (*xsdDate)(overlay.T.FirstDayofWork)
	overlay.EndEmploymentDate = (*xsdDate)(overlay.T.EndEmploymentDate)
	return d.DecodeElement(&overlay, &start)
}

// This web service operation is designed to allow an international assignment to be added to an existing active employee
type StartInternationalAssignmentRequestType struct {
	BusinessProcessParameters        *BusinessProcessParametersType        `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	StartInternationalAssignmentData *StartInternationalAssignmentDataType `xml:"urn:com.workday/bsvc Start_International_Assignment_Data"`
	Version                          string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Event ID for the Start International Assignment for Employee Event and the job reference.
type StartInternationalAssignmentResponseType struct {
	StartInternationalAssignmentReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Start_International_Assignment_Reference,omitempty"`
	EmployeeReference                     *WorkerObjectType           `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	JobReference                          *PositionElementObjectType  `xml:"urn:com.workday/bsvc Job_Reference,omitempty"`
	Version                               string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StockGrantTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StockGrantTypeObjectType struct {
	ID         []StockGrantTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type StockVestingScheduleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StockVestingScheduleObjectType struct {
	ID         []StockVestingScheduleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StudentAcademicStandingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StudentAcademicStandingObjectType struct {
	ID         []StudentAcademicStandingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Personal Information data for a Student Applicant
type StudentApplicationDetailsPersonalInformationSubeditDataType struct {
	MaritalStatusReference     []MaritalStatusObjectType            `xml:"urn:com.workday/bsvc Marital_Status_Reference,omitempty"`
	CountryofBirthReference    *CountryObjectType                   `xml:"urn:com.workday/bsvc Country_of_Birth_Reference,omitempty"`
	RegionofBirthReference     *CountryRegionObjectType             `xml:"urn:com.workday/bsvc Region_of_Birth_Reference,omitempty"`
	CityofBirth                string                               `xml:"urn:com.workday/bsvc City_of_Birth,omitempty"`
	GenericMilitaryServiceData []GenericMilitaryInformationDataType `xml:"urn:com.workday/bsvc Generic_Military_Service_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StudentCourseTagObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StudentCourseTagObjectType struct {
	ID         []StudentCourseTagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Request element to update Student Employment Eligibility Event
type StudentEmploymentEligibilityEventDataType struct {
	Eligible                               *bool                                        `xml:"urn:com.workday/bsvc Eligible,omitempty"`
	StudentEmploymentEligibilityReasonData []StudentEmploymentEligibilityReasonDataType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Reason_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StudentEmploymentEligibilityEventObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StudentEmploymentEligibilityEventObjectType struct {
	ID         []StudentEmploymentEligibilityEventObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Additional data to be included with response
type StudentEmploymentEligibilityEventRequestCriteriaType struct {
}

// Container Element for included references
type StudentEmploymentEligibilityEventRequestReferencesType struct {
	StudentEmploymentEligibilityEventReference []StudentEmploymentEligibilityEventObjectType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Event_Reference"`
}

// Container for Student Employment Eligibility Event data
type StudentEmploymentEligibilityEventResponseDataType struct {
	StudentEmploymentEligibilityEvent []StudentEmploymentEligibilityEventType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Event,omitempty"`
}

// Container for response behavior
type StudentEmploymentEligibilityEventResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Container for individual student employment eligibility data
type StudentEmploymentEligibilityEventType struct {
	StudentEmploymentEligibilityEventReference *StudentEmploymentEligibilityEventObjectType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Event_Reference,omitempty"`
	PreHireReference                           *ApplicantObjectType                         `xml:"urn:com.workday/bsvc Pre-Hire_Reference,omitempty"`
	WorkerReference                            *WorkerObjectType                            `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	ProposedWorkerID                           string                                       `xml:"urn:com.workday/bsvc Proposed_Worker_ID,omitempty"`
	ParentEventReference                       *ActionEventObjectType                       `xml:"urn:com.workday/bsvc Parent_Event_Reference,omitempty"`
	EmployeeTypeReference                      *EmployeeTypeObjectType                      `xml:"urn:com.workday/bsvc Employee_Type_Reference,omitempty"`
	SupervisoryOrganizationReference           *SupervisoryOrganizationObjectType           `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference,omitempty"`
	ManagerReference                           *WorkerObjectType                            `xml:"urn:com.workday/bsvc Manager_Reference,omitempty"`
	EventEffectiveDate                         *time.Time                                   `xml:"urn:com.workday/bsvc Event_Effective_Date,omitempty"`
	TotalScheduledHours                        float64                                      `xml:"urn:com.workday/bsvc Total_Scheduled_Hours,omitempty"`
	JobRequisitionReference                    *JobRequisitionObjectType                    `xml:"urn:com.workday/bsvc Job_Requisition_Reference,omitempty"`
	PositionRestrictionReference               *PositionRestrictionsObjectType              `xml:"urn:com.workday/bsvc Position_Restriction_Reference,omitempty"`
	PositionDetailsData                        []PositionDetailsSubDataType                 `xml:"urn:com.workday/bsvc Position_Details_Data,omitempty"`
}

func (t *StudentEmploymentEligibilityEventType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T StudentEmploymentEligibilityEventType
	var layout struct {
		*T
		EventEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Event_Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EventEffectiveDate = (*xsdDate)(layout.T.EventEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *StudentEmploymentEligibilityEventType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T StudentEmploymentEligibilityEventType
	var overlay struct {
		*T
		EventEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Event_Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EventEffectiveDate = (*xsdDate)(overlay.T.EventEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Reason and description of why event is ineligible
type StudentEmploymentEligibilityReasonDataType struct {
	ReasonforIneligibility string `xml:"urn:com.workday/bsvc Reason_for_Ineligibility,omitempty"`
	Critical               *bool  `xml:"urn:com.workday/bsvc Critical,omitempty"`
}

// Contains the Data for Student Employment Eligibility Rule
type StudentEmploymentEligibilityRuleDataType struct {
	ConditionRuleData []ConditionRuleDataWWSType `xml:"urn:com.workday/bsvc Condition_Rule_Data"`
}

// Container for providing a specific Student Employment Eligibility Rule Reference
type StudentEmploymentEligibilityRuleRequestReferencesType struct {
	StudentEmploymentEligibilityRuleReference []ConditionRuleObjectType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Reference"`
}

// Container for Student Employment Eligibility Rule Data
type StudentEmploymentEligibilityRuleResponseDataType struct {
	StudentEmploymentEligibilityRule []StudentEmploymentEligibilityRuleType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule,omitempty"`
}

// Container for determining what is included in the Response
type StudentEmploymentEligibilityRuleResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Container for Put Student Employment Eligibility Rule Set Response
type StudentEmploymentEligibilityRuleSetDataType struct {
	ID                                             string                                            `xml:"urn:com.workday/bsvc ID,omitempty"`
	StudentEmploymentEligibilityRuleSetName        string                                            `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Set_Name"`
	StudentEmploymentEligibilityRuleSetDescription string                                            `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Set_Description,omitempty"`
	Inactive                                       *bool                                             `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	JobProfileReference                            []JobProfileObjectType                            `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	JobFamilyReference                             []JobFamilyObjectType                             `xml:"urn:com.workday/bsvc Job_Family_Reference,omitempty"`
	AcademicLevelReference                         []AcademicLevelObjectType                         `xml:"urn:com.workday/bsvc Academic_Level_Reference,omitempty"`
	AcademicPeriodReference                        []AcademicPeriodObjectType                        `xml:"urn:com.workday/bsvc Academic_Period_Reference,omitempty"`
	AcademicPeriodTypeReference                    []AcademicPeriodTypeObjectType                    `xml:"urn:com.workday/bsvc Academic_Period_Type_Reference,omitempty"`
	StudentEmploymentEligibilityRuleSubeditData    []StudentEmploymentEligibilityRuleSubeditDataType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Subedit_Data"`
}

// Contains a unique identifier for an instance of an object.
type StudentEmploymentEligibilityRuleSetObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StudentEmploymentEligibilityRuleSetObjectType struct {
	ID         []StudentEmploymentEligibilityRuleSetObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Reference  Student Employment Eligibility Rule Set Request
type StudentEmploymentEligibilityRuleSetRequestReferencesType struct {
	StudentEmploymentEligibilityRuleSetRequestReference []StudentEmploymentEligibilityRuleSetObjectType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Set_Request_Reference"`
}

// Container Student Employment Eligibility Rule Set Response Data
type StudentEmploymentEligibilityRuleSetResponseDataType struct {
	StudentEmploymentEligibilityRuleSet []StudentEmploymentEligibilityRuleSetType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Set,omitempty"`
}

// Container for Student Employment Eligibility Rule Set Response Group
type StudentEmploymentEligibilityRuleSetResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Container for Student Employment Eligibility Rule Set Data
type StudentEmploymentEligibilityRuleSetType struct {
	StudentEmploymentEligibilityRuleSetReference *StudentEmploymentEligibilityRuleSetObjectType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Set_Reference,omitempty"`
	StudentEmploymentEligibilityRuleSetData      []StudentEmploymentEligibilityRuleSetDataType  `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Set_Data,omitempty"`
}

// Container for Student Employment Eligibility Rule Subedit Data
type StudentEmploymentEligibilityRuleSubeditDataType struct {
	ConditionRuleReference                    *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Condition_Rule_Reference"`
	StudentEmploymentEligibilityRuleReference string                   `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Reference,omitempty"`
	ReasonforIneligibility                    string                   `xml:"urn:com.workday/bsvc Reason_for_Ineligibility"`
	CriticalSeverity                          *bool                    `xml:"urn:com.workday/bsvc Critical_Severity,omitempty"`
}

// Container for Student Employment Eligibility Condition Rule Data
type StudentEmploymentEligibilityRuleType struct {
	StudentEmploymentEligibilityRuleReference *ConditionRuleObjectType                   `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Reference,omitempty"`
	StudentEmploymentEligibilityRuleData      []StudentEmploymentEligibilityRuleDataType `xml:"urn:com.workday/bsvc Student_Employment_Eligibility_Rule_Data,omitempty"`
}

// Wrapper element for the Student Employment Eligibility sub business process.
type StudentEmploymentEligibilitySubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StudentLoadStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StudentLoadStatusObjectType struct {
	ID         []StudentLoadStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StudentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StudentObjectType struct {
	ID         []StudentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains personal information of the Student Prospect.
type StudentPersonalProfileDataType struct {
	DateofBirth                *time.Time                    `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	GenderReference            *GenderObjectType             `xml:"urn:com.workday/bsvc Gender_Reference,omitempty"`
	EthnicityReference         []EthnicityObjectType         `xml:"urn:com.workday/bsvc Ethnicity_Reference,omitempty"`
	CitizenshipStatusReference []CitizenshipStatusObjectType `xml:"urn:com.workday/bsvc Citizenship_Status_Reference,omitempty"`
	HispanicorLatino           *bool                         `xml:"urn:com.workday/bsvc Hispanic_or_Latino,omitempty"`
}

func (t *StudentPersonalProfileDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T StudentPersonalProfileDataType
	var layout struct {
		*T
		DateofBirth *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateofBirth = (*xsdDate)(layout.T.DateofBirth)
	return e.EncodeElement(layout, start)
}
func (t *StudentPersonalProfileDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T StudentPersonalProfileDataType
	var overlay struct {
		*T
		DateofBirth *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateofBirth = (*xsdDate)(overlay.T.DateofBirth)
	return d.DecodeElement(&overlay, &start)
}

// Element containing the photo data for the student.
type StudentPhotoDataType struct {
	Filename string `xml:"urn:com.workday/bsvc Filename"`
	File     []byte `xml:"urn:com.workday/bsvc File"`
}

func (t *StudentPhotoDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T StudentPhotoDataType
	var layout struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File"`
	}
	layout.T = (*T)(t)
	layout.File = (*xsdBase64Binary)(&layout.T.File)
	return e.EncodeElement(layout, start)
}
func (t *StudentPhotoDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T StudentPhotoDataType
	var overlay struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File"`
	}
	overlay.T = (*T)(t)
	overlay.File = (*xsdBase64Binary)(&overlay.T.File)
	return d.DecodeElement(&overlay, &start)
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

// Contains information about the succession plan candidacy for the employee.
type SuccessionPlanCandidateDataType struct {
	PositionElementReference     *PositionElementObjectType     `xml:"urn:com.workday/bsvc Position_Element_Reference,omitempty"`
	SuccessionReadinessReference *SuccessionReadinessObjectType `xml:"urn:com.workday/bsvc Succession_Readiness_Reference,omitempty"`
	TopCandidate                 *bool                          `xml:"urn:com.workday/bsvc Top_Candidate,omitempty"`
	Notes                        string                         `xml:"urn:com.workday/bsvc Notes,omitempty"`
	LastUpdated                  string                         `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
}

// Contains information about all the succession plans that this employee is on.
type SuccessionProfileDataType struct {
	CandidateData []SuccessionPlanCandidateDataType `xml:"urn:com.workday/bsvc Candidate_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SuccessionReadinessObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SuccessionReadinessObjectType struct {
	ID         []SuccessionReadinessObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Wrapper element for the Switch Primary Job business process.
type SwitchPrimaryJobDataType struct {
	ReasonReference             *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference"`
	ProposedPrimaryJobReference *PositionElementObjectType                `xml:"urn:com.workday/bsvc Proposed_Primary_Job_Reference,omitempty"`
}

// Responds with the Event ID for the Switch Primary Job business process.
type SwitchPrimaryJobEventResponseType struct {
	EventReference         *UniqueIdentifierObjectType                    `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	ExceptionsResponseData []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Switch Primary Job business process.
type SwitchPrimaryJobRequestDataType struct {
	WorkerReference      *WorkerObjectType         `xml:"urn:com.workday/bsvc Worker_Reference"`
	EffectiveDate        time.Time                 `xml:"urn:com.workday/bsvc Effective_Date"`
	SwitchPrimaryJobData *SwitchPrimaryJobDataType `xml:"urn:com.workday/bsvc Switch_Primary_Job_Data"`
}

func (t *SwitchPrimaryJobRequestDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T SwitchPrimaryJobRequestDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *SwitchPrimaryJobRequestDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SwitchPrimaryJobRequestDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// This web service operation is designed to switch primary jobs for a worker using the Switch Primary Job Business Process
type SwitchPrimaryJobRequestType struct {
	BusinessProcessParameters        *BusinessProcessParametersType   `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	SwitchPrimaryPositionRequestData *SwitchPrimaryJobRequestDataType `xml:"urn:com.workday/bsvc Switch_Primary_Position_Request_Data"`
	Version                          string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Switch Primary Job sub process.
type SwitchPrimaryJobSubProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	SwitchPrimaryJobData         *SwitchPrimaryJobDataType         `xml:"urn:com.workday/bsvc Switch_Primary_Job_Data,omitempty"`
}

// The data necessary to create or update a user's account in the Workday system.
type SystemUserDataType struct {
	UserName                                          string                                              `xml:"urn:com.workday/bsvc User_Name"`
	SessionTimeoutMinutes                             float64                                             `xml:"urn:com.workday/bsvc Session_Timeout_Minutes,omitempty"`
	AccountDisabled                                   *bool                                               `xml:"urn:com.workday/bsvc Account_Disabled,omitempty"`
	AccountExpirationDate                             *time.Time                                          `xml:"urn:com.workday/bsvc Account_Expiration_Date,omitempty"`
	AccountLocked                                     *bool                                               `xml:"urn:com.workday/bsvc Account_Locked,omitempty"`
	RequiredNewPasswordAtNextLogin                    *bool                                               `xml:"urn:com.workday/bsvc Required_New_Password_At_Next_Login,omitempty"`
	ShowUserNameinBrowserWindow                       *bool                                               `xml:"urn:com.workday/bsvc Show_User_Name_in_Browser_Window,omitempty"`
	DisplayXMLIcononReports                           *bool                                               `xml:"urn:com.workday/bsvc Display_XML_Icon_on_Reports,omitempty"`
	EnableWorkbox                                     *bool                                               `xml:"urn:com.workday/bsvc Enable_Workbox,omitempty"`
	LocaleReference                                   *LocaleObjectType                                   `xml:"urn:com.workday/bsvc Locale_Reference,omitempty"`
	UserLanguageReference                             *UserLanguageObjectType                             `xml:"urn:com.workday/bsvc User_Language_Reference,omitempty"`
	PreferredSearchScopeReference                     *UniqueIdentifierObjectType                         `xml:"urn:com.workday/bsvc Preferred_Search_Scope_Reference,omitempty"`
	DelegatedAuthenticationIntegrationSystemReference *IntegrationSystemAuditedObjectType                 `xml:"urn:com.workday/bsvc Delegated_Authentication_Integration_System_Reference,omitempty"`
	AllowMixedLanguageTransactions                    *bool                                               `xml:"urn:com.workday/bsvc Allow_Mixed-Language_Transactions,omitempty"`
	NotificationSubTypeConfigurations                 []NotificationSubCategoryConfigurationsforUsersType `xml:"urn:com.workday/bsvc Notification_Sub_Type_Configurations,omitempty"`
	Password                                          string                                              `xml:"urn:com.workday/bsvc Password,omitempty"`
	GenerateRandomPassword                            *bool                                               `xml:"urn:com.workday/bsvc Generate_Random_Password,omitempty"`
	ExemptfromDelegatedAuthentication                 *bool                                               `xml:"urn:com.workday/bsvc Exempt_from_Delegated_Authentication,omitempty"`
	PasscodeExempt                                    *bool                                               `xml:"urn:com.workday/bsvc Passcode_Exempt,omitempty"`
	PasscodeGracePeriodEnabled                        *bool                                               `xml:"urn:com.workday/bsvc Passcode_Grace_Period_Enabled,omitempty"`
	PasscodeGracePeriodLoginRemainingCount            float64                                             `xml:"urn:com.workday/bsvc Passcode_Grace_Period_Login_Remaining_Count,omitempty"`
	OpenIDIdentifier                                  string                                              `xml:"urn:com.workday/bsvc OpenID_Identifier,omitempty"`
	OpenIDInternalIdentifier                          string                                              `xml:"urn:com.workday/bsvc OpenID_Internal_Identifier,omitempty"`
	OpenIDConnectInternalIdentifier                   string                                              `xml:"urn:com.workday/bsvc OpenID_Connect_Internal_Identifier,omitempty"`
	SimplifiedView                                    *bool                                               `xml:"urn:com.workday/bsvc Simplified_View,omitempty"`
}

func (t *SystemUserDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T SystemUserDataType
	var layout struct {
		*T
		AccountExpirationDate *xsdDateTime `xml:"urn:com.workday/bsvc Account_Expiration_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AccountExpirationDate = (*xsdDateTime)(layout.T.AccountExpirationDate)
	return e.EncodeElement(layout, start)
}
func (t *SystemUserDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SystemUserDataType
	var overlay struct {
		*T
		AccountExpirationDate *xsdDateTime `xml:"urn:com.workday/bsvc Account_Expiration_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AccountExpirationDate = (*xsdDateTime)(overlay.T.AccountExpirationDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains the worker's user account information (user name, preferred language).
type SystemUserforWorkerDataType struct {
	UserName                                string                  `xml:"urn:com.workday/bsvc User_Name"`
	UserLanguageReference                   *UserLanguageObjectType `xml:"urn:com.workday/bsvc User_Language__Reference,omitempty"`
	PreferredCommunicationLanguageReference *UserLanguageObjectType `xml:"urn:com.workday/bsvc Preferred_Communication_Language_Reference,omitempty"`
	LocaleReference                         *LocaleObjectType       `xml:"urn:com.workday/bsvc Locale__Reference,omitempty"`
	CurrencyReference                       *CurrencyObjectType     `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	TimeZoneReference                       *TimeZoneObjectType     `xml:"urn:com.workday/bsvc Time_Zone_Reference,omitempty"`
	DefaultDisplayLanguageReference         *UserLanguageObjectType `xml:"urn:com.workday/bsvc Default_Display_Language_Reference,omitempty"`
	SimplifiedView                          *bool                   `xml:"urn:com.workday/bsvc Simplified_View,omitempty"`
}

// Contains the potential assessment for the employee.
type TalentAssessmentDataType struct {
	EmployeePotentialData *EmployeeTalentAssessmentDataType `xml:"urn:com.workday/bsvc Employee_Potential_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TalentTagObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TalentTagObjectType struct {
	ID         []TalentTagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TaxAuthorityFormTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TaxAuthorityFormTypeObjectType struct {
	ID         []TaxAuthorityFormTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// This web service is used to Terminate an Employee. It will run the Terminate Employee Business Process and (optionally) its subprocesses.
type TerminateEmployeeDataType struct {
	EmployeeReference                           *EmployeeObjectType                                      `xml:"urn:com.workday/bsvc Employee_Reference"`
	TerminationDate                             time.Time                                                `xml:"urn:com.workday/bsvc Termination_Date"`
	TerminateEventData                          *TerminateEventDataType                                  `xml:"urn:com.workday/bsvc Terminate_Event_Data"`
	RequestOneTimePaymentSubProcess             *RequestOneTimePaymentSubBusinessProcessType             `xml:"urn:com.workday/bsvc Request_One_Time_Payment_Sub_Process,omitempty"`
	ReviewCOBRAEligibilitySubProcess            *ReviewCOBRAEligibilitySubBusinessProcessType            `xml:"urn:com.workday/bsvc Review_COBRA_Eligibility_Sub_Process,omitempty"`
	ReviewPayrollInterfaceEventSubProcess       *ReviewPayrollInterfaceDataExtractBusinessSubProcessType `xml:"urn:com.workday/bsvc Review_Payroll_Interface_Event_Sub_Process,omitempty"`
	EditServiceDatesSubProcess                  *EditServiceDatesSubBusinessProcessType                  `xml:"urn:com.workday/bsvc Edit_Service_Dates_Sub_Process,omitempty"`
	AddRetireeStatusSubProcess                  *AddRetireeStatusSubBusinessProcessType                  `xml:"urn:com.workday/bsvc Add_Retiree_Status_Sub_Process,omitempty"`
	MaintainEmployeeContractsSubBusinessProcess *MaintainEmployeeContractsSubBusinessProcessType         `xml:"urn:com.workday/bsvc Maintain_Employee_Contracts_Sub_Business_Process,omitempty"`
	CheckPositionBudgetSubProcess               *CheckPositionBudgetSubBusinessProcessType               `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process,omitempty"`
	AssignOrganizationRolesSubProcess           *AssignOrganizationRolesSubBusinessProcessType           `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Sub_Process,omitempty"`
	CreateJobRequisitionSubProcess              *CreateJobRequisitionSubProcessType                      `xml:"urn:com.workday/bsvc Create_Job_Requisition_Sub_Process,omitempty"`
	EndAcademicAppointmentSubProcess            *EndAcademicAppointmentSubBusinessProcessType            `xml:"urn:com.workday/bsvc End_Academic_Appointment_Sub_Process,omitempty"`
	AssignCostingAllocationSubProcess           *AssignCostingAllocationSubBusinessProcessType           `xml:"urn:com.workday/bsvc Assign_Costing_Allocation_Sub_Process,omitempty"`
	ManageUnionMembershipSubProcess             *ManageUnionMembershipSubBusinessProcessType             `xml:"urn:com.workday/bsvc Manage_Union_Membership_Sub_Process,omitempty"`
}

func (t *TerminateEmployeeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T TerminateEmployeeDataType
	var layout struct {
		*T
		TerminationDate *xsdDate `xml:"urn:com.workday/bsvc Termination_Date"`
	}
	layout.T = (*T)(t)
	layout.TerminationDate = (*xsdDate)(&layout.T.TerminationDate)
	return e.EncodeElement(layout, start)
}
func (t *TerminateEmployeeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T TerminateEmployeeDataType
	var overlay struct {
		*T
		TerminationDate *xsdDate `xml:"urn:com.workday/bsvc Termination_Date"`
	}
	overlay.T = (*T)(t)
	overlay.TerminationDate = (*xsdDate)(&overlay.T.TerminationDate)
	return d.DecodeElement(&overlay, &start)
}

// Responds with the Event ID for the Terminate Employee Event along with the worker reference.
type TerminateEmployeeEventResponseType struct {
	EventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This web service operation is designed to terminate an existing Employee from a Position
type TerminateEmployeeRequestHVType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	TerminateEmployeeData     *TerminateEmployeeDataType     `xml:"urn:com.workday/bsvc Terminate_Employee_Data"`
}

// Wrapper for Terminate Employee Web Service and its sub-processes.
type TerminateEmployeeRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	TerminateEmployeeData     *TerminateEmployeeDataType     `xml:"urn:com.workday/bsvc Terminate_Employee_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Data Element that contains the basic information to terminate an employee. All required information must be included. The Terminate Employee business process (for the employee) will be initiated from this information.
type TerminateEventDataType struct {
	LastDayofWork                   *time.Time                                `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
	PrimaryReasonReference          *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Primary_Reason_Reference"`
	SecondaryReasonReference        []TerminationSubcategoryObjectType        `xml:"urn:com.workday/bsvc Secondary_Reason_Reference,omitempty"`
	LocalTerminationReasonReference *LocalTerminationReasonObjectType         `xml:"urn:com.workday/bsvc Local_Termination_Reason_Reference,omitempty"`
	PayThroughDate                  *time.Time                                `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
	ResignationDate                 *time.Time                                `xml:"urn:com.workday/bsvc Resignation_Date,omitempty"`
	NotificationDate                *time.Time                                `xml:"urn:com.workday/bsvc Notification_Date,omitempty"`
	RegrettableReference            *CommonYesNoObjectType                    `xml:"urn:com.workday/bsvc Regrettable_Reference,omitempty"`
	EligibleforHireReference        *CommonYesNoObjectType                    `xml:"urn:com.workday/bsvc Eligible_for_Hire_Reference,omitempty"`
	ClosePosition                   *bool                                     `xml:"urn:com.workday/bsvc Close_Position,omitempty"`
	JobOverlapAllowed               *bool                                     `xml:"urn:com.workday/bsvc Job_Overlap_Allowed,omitempty"`
	LastDateforWhichPaid            *time.Time                                `xml:"urn:com.workday/bsvc Last_Date_for_Which_Paid,omitempty"`
	ExpectedDateofReturn            *time.Time                                `xml:"urn:com.workday/bsvc Expected_Date_of_Return,omitempty"`
	NotReturning                    *bool                                     `xml:"urn:com.workday/bsvc Not_Returning,omitempty"`
	ReturnUnknown                   *bool                                     `xml:"urn:com.workday/bsvc Return_Unknown,omitempty"`
	WorkerDocumentData              []WorkerDocumentforStaffingEventDataType  `xml:"urn:com.workday/bsvc Worker_Document_Data,omitempty"`
}

func (t *TerminateEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T TerminateEventDataType
	var layout struct {
		*T
		LastDayofWork        *xsdDate `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
		PayThroughDate       *xsdDate `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
		ResignationDate      *xsdDate `xml:"urn:com.workday/bsvc Resignation_Date,omitempty"`
		NotificationDate     *xsdDate `xml:"urn:com.workday/bsvc Notification_Date,omitempty"`
		LastDateforWhichPaid *xsdDate `xml:"urn:com.workday/bsvc Last_Date_for_Which_Paid,omitempty"`
		ExpectedDateofReturn *xsdDate `xml:"urn:com.workday/bsvc Expected_Date_of_Return,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LastDayofWork = (*xsdDate)(layout.T.LastDayofWork)
	layout.PayThroughDate = (*xsdDate)(layout.T.PayThroughDate)
	layout.ResignationDate = (*xsdDate)(layout.T.ResignationDate)
	layout.NotificationDate = (*xsdDate)(layout.T.NotificationDate)
	layout.LastDateforWhichPaid = (*xsdDate)(layout.T.LastDateforWhichPaid)
	layout.ExpectedDateofReturn = (*xsdDate)(layout.T.ExpectedDateofReturn)
	return e.EncodeElement(layout, start)
}
func (t *TerminateEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T TerminateEventDataType
	var overlay struct {
		*T
		LastDayofWork        *xsdDate `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
		PayThroughDate       *xsdDate `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
		ResignationDate      *xsdDate `xml:"urn:com.workday/bsvc Resignation_Date,omitempty"`
		NotificationDate     *xsdDate `xml:"urn:com.workday/bsvc Notification_Date,omitempty"`
		LastDateforWhichPaid *xsdDate `xml:"urn:com.workday/bsvc Last_Date_for_Which_Paid,omitempty"`
		ExpectedDateofReturn *xsdDate `xml:"urn:com.workday/bsvc Expected_Date_of_Return,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LastDayofWork = (*xsdDate)(overlay.T.LastDayofWork)
	overlay.PayThroughDate = (*xsdDate)(overlay.T.PayThroughDate)
	overlay.ResignationDate = (*xsdDate)(overlay.T.ResignationDate)
	overlay.NotificationDate = (*xsdDate)(overlay.T.NotificationDate)
	overlay.LastDateforWhichPaid = (*xsdDate)(overlay.T.LastDateforWhichPaid)
	overlay.ExpectedDateofReturn = (*xsdDate)(overlay.T.ExpectedDateofReturn)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type TerminationSubcategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TerminationSubcategoryObjectType struct {
	ID         []TerminationSubcategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TimeProfileObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimeProfileObjectType struct {
	ID         []TimeProfileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Encapsulating element containing the data of the transaction log entry.
type TransactionLogDataType struct {
	TransactionLogDescription   string                         `xml:"urn:com.workday/bsvc Transaction_Log_Description,omitempty"`
	TransactionEffectiveMoment  *time.Time                     `xml:"urn:com.workday/bsvc Transaction_Effective_Moment,omitempty"`
	TransactionEntryMoment      *time.Time                     `xml:"urn:com.workday/bsvc Transaction_Entry_Moment,omitempty"`
	TransactionLogTypeReference []TransactionLogTypeObjectType `xml:"urn:com.workday/bsvc Transaction_Log_Type_Reference,omitempty"`
	TransactionTargetReference  []EventTargetObjectType        `xml:"urn:com.workday/bsvc Transaction_Target_Reference,omitempty"`
	IsRescindOrRescinded        *bool                          `xml:"urn:com.workday/bsvc Is_Rescind_Or_Rescinded,omitempty"`
	IsCorrectionOrCorrected     *bool                          `xml:"urn:com.workday/bsvc Is_Correction_Or_Corrected,omitempty"`
}

func (t *TransactionLogDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T TransactionLogDataType
	var layout struct {
		*T
		TransactionEffectiveMoment *xsdDateTime `xml:"urn:com.workday/bsvc Transaction_Effective_Moment,omitempty"`
		TransactionEntryMoment     *xsdDateTime `xml:"urn:com.workday/bsvc Transaction_Entry_Moment,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TransactionEffectiveMoment = (*xsdDateTime)(layout.T.TransactionEffectiveMoment)
	layout.TransactionEntryMoment = (*xsdDateTime)(layout.T.TransactionEntryMoment)
	return e.EncodeElement(layout, start)
}
func (t *TransactionLogDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T TransactionLogDataType
	var overlay struct {
		*T
		TransactionEffectiveMoment *xsdDateTime `xml:"urn:com.workday/bsvc Transaction_Effective_Moment,omitempty"`
		TransactionEntryMoment     *xsdDateTime `xml:"urn:com.workday/bsvc Transaction_Entry_Moment,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionEffectiveMoment = (*xsdDateTime)(overlay.T.TransactionEffectiveMoment)
	overlay.TransactionEntryMoment = (*xsdDateTime)(overlay.T.TransactionEntryMoment)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing the transaction log entry.
type TransactionLogEntryType struct {
	TransactionLogReference *TransactionLogObjectType `xml:"urn:com.workday/bsvc Transaction_Log_Reference,omitempty"`
	TransactionLogData      *TransactionLogDataType   `xml:"urn:com.workday/bsvc Transaction_Log_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TransactionLogObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TransactionLogObjectType struct {
	ID         []TransactionLogObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Details on the rescind or correct of the event described in the main transaction data section.
type TransactionLogRescindAndCorrectEventDataType struct {
	TransactionLogReference   *TransactionLogObjectType `xml:"urn:com.workday/bsvc Transaction_Log_Reference,omitempty"`
	TransactionLogDescription string                    `xml:"urn:com.workday/bsvc Transaction_Log_Description,omitempty"`
	IsCorrection              *bool                     `xml:"urn:com.workday/bsvc Is_Correction,omitempty"`
	IsRescind                 *bool                     `xml:"urn:com.workday/bsvc Is_Rescind,omitempty"`
	PriorEffectiveMoment      *time.Time                `xml:"urn:com.workday/bsvc Prior_Effective_Moment,omitempty"`
	TransactionEntryMoment    *time.Time                `xml:"urn:com.workday/bsvc Transaction_Entry_Moment,omitempty"`
}

func (t *TransactionLogRescindAndCorrectEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T TransactionLogRescindAndCorrectEventDataType
	var layout struct {
		*T
		PriorEffectiveMoment   *xsdDate     `xml:"urn:com.workday/bsvc Prior_Effective_Moment,omitempty"`
		TransactionEntryMoment *xsdDateTime `xml:"urn:com.workday/bsvc Transaction_Entry_Moment,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PriorEffectiveMoment = (*xsdDate)(layout.T.PriorEffectiveMoment)
	layout.TransactionEntryMoment = (*xsdDateTime)(layout.T.TransactionEntryMoment)
	return e.EncodeElement(layout, start)
}
func (t *TransactionLogRescindAndCorrectEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T TransactionLogRescindAndCorrectEventDataType
	var overlay struct {
		*T
		PriorEffectiveMoment   *xsdDate     `xml:"urn:com.workday/bsvc Prior_Effective_Moment,omitempty"`
		TransactionEntryMoment *xsdDateTime `xml:"urn:com.workday/bsvc Transaction_Entry_Moment,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PriorEffectiveMoment = (*xsdDate)(overlay.T.PriorEffectiveMoment)
	overlay.TransactionEntryMoment = (*xsdDateTime)(overlay.T.TransactionEntryMoment)
	return d.DecodeElement(&overlay, &start)
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
type TravelAmountObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TravelAmountObjectType struct {
	ID         []TravelAmountObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the information about the employees willingness to travel as part of their job.
type TravelPreferenceDataType struct {
	WillingtoTravelReference    *MobilityChoiceObjectType `xml:"urn:com.workday/bsvc Willing_to_Travel_Reference,omitempty"`
	TravelAmountReference       *TravelAmountObjectType   `xml:"urn:com.workday/bsvc Travel_Amount_Reference,omitempty"`
	TravelAdditionalInformation string                    `xml:"urn:com.workday/bsvc Travel_Additional_Information,omitempty"`
}

// Element  containing a single union member with associated membership details.
type UnionMemberDataType struct {
	WorkerReference     *WorkerObjectType         `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	UnionReference      *UnionObjectType          `xml:"urn:com.workday/bsvc Union_Reference"`
	SeniorityDate       *time.Time                `xml:"urn:com.workday/bsvc Seniority_Date,omitempty"`
	UnionMembershipData []UnionMembershipDataType `xml:"urn:com.workday/bsvc Union_Membership_Data"`
}

func (t *UnionMemberDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T UnionMemberDataType
	var layout struct {
		*T
		SeniorityDate *xsdDate `xml:"urn:com.workday/bsvc Seniority_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.SeniorityDate = (*xsdDate)(layout.T.SeniorityDate)
	return e.EncodeElement(layout, start)
}
func (t *UnionMemberDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T UnionMemberDataType
	var overlay struct {
		*T
		SeniorityDate *xsdDate `xml:"urn:com.workday/bsvc Seniority_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.SeniorityDate = (*xsdDate)(overlay.T.SeniorityDate)
	return d.DecodeElement(&overlay, &start)
}

// Element containing specific union membership dates and details.
type UnionMembershipDataType struct {
	StartDate                time.Time                      `xml:"urn:com.workday/bsvc Start_Date"`
	EndDate                  *time.Time                     `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	MembershipTypeReference  *UnionMembershipTypeObjectType `xml:"urn:com.workday/bsvc Membership_Type_Reference"`
	RelatedPositionReference *StaffingInterfaceObjectType   `xml:"urn:com.workday/bsvc Related_Position_Reference,omitempty"`
	Notes                    string                         `xml:"urn:com.workday/bsvc Notes,omitempty"`
}

func (t *UnionMembershipDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T UnionMembershipDataType
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
func (t *UnionMembershipDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T UnionMembershipDataType
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

// Contains a unique identifier for an instance of an object.
type UnionMembershipTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type UnionMembershipTypeObjectType struct {
	ID         []UnionMembershipTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type UnitofMeasureObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type UnitofMeasureObjectType struct {
	ID         []UnitofMeasureObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type UniversalIdentifierObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type UniversalIdentifierObjectType struct {
	ID         []UniversalIdentifierObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for the Update Academic Appointment business process.
type UpdateAcademicAppointmentDataType struct {
	ReasonReference            *GeneralEventSubcategoryObjectType   `xml:"urn:com.workday/bsvc Reason_Reference"`
	AcademicAppointeeReference *AcademicAppointeeEnabledObjectType  `xml:"urn:com.workday/bsvc Academic_Appointee_Reference"`
	AcademicAppointmentData    *AcademicAppointmentSnapshotDataType `xml:"urn:com.workday/bsvc Academic_Appointment_Data"`
}

// Wrapper element for the Update Academic Appointment sub business process.
type UpdateAcademicAppointmentSubProcessType struct {
	BusinessSubProcessParameters  *BusinessSubProcessParametersType  `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	UpdateAcademicAppointmentData *UpdateAcademicAppointmentDataType `xml:"urn:com.workday/bsvc Update_Academic_Appointment_Data,omitempty"`
}

// The main element for the Change ID Business Process. This contains all the ID information  and Business Processing processing parameters.
type UpdateIDInformationSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	PersonIdentificationData     []PersonIdentificationDataType    `xml:"urn:com.workday/bsvc Person_Identification_Data,omitempty"`
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

// The volume information for the coverage level selected.
type VolumeDataType struct {
	Volume            float64             `xml:"urn:com.workday/bsvc Volume"`
	CurrencyReference *CurrencyObjectType `xml:"urn:com.workday/bsvc Currency_Reference"`
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
type WorkHoursProfileObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkHoursProfileObjectType struct {
	ID         []WorkHoursProfileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Sub element used to change work space on a Position.
type WorkSpaceChangeEventDataType struct {
	PositionReference  *PositionElementObjectType `xml:"urn:com.workday/bsvc Position_Reference"`
	EffectiveDate      time.Time                  `xml:"urn:com.workday/bsvc Effective_Date"`
	WorkSpaceReference *LocationObjectType        `xml:"urn:com.workday/bsvc Work_Space_Reference"`
}

func (t *WorkSpaceChangeEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkSpaceChangeEventDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *WorkSpaceChangeEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkSpaceChangeEventDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

type WorkdayCommonHeaderType struct {
	IncludeReferenceDescriptorsInResponse *bool `xml:"urn:com.workday/bsvc Include_Reference_Descriptors_In_Response,omitempty"`
}

// Contains the additional benefits coverage information (elections) for an employee.
type WorkerAdditionalBenefitsCoverageDataType struct {
	AdditionalBenefitsCoverageTargetReference  *AdditionalBenefitsCoverageTargetObjectType `xml:"urn:com.workday/bsvc Additional_Benefits_Coverage_Target_Reference,omitempty"`
	OriginalCoverageBeginDateforCoverageTarget *time.Time                                  `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date_for_Coverage_Target,omitempty"`
	BenefitElectionData                        *WorkerBenefitElectionDataType              `xml:"urn:com.workday/bsvc Benefit_Election_Data,omitempty"`
	PercentContributionData                    *EmployeeContributionPercentageDataType     `xml:"urn:com.workday/bsvc Percent_Contribution_Data,omitempty"`
	AmountContributionData                     []EmployeeContributionAmountDataType        `xml:"urn:com.workday/bsvc Amount_Contribution_Data,omitempty"`
}

func (t *WorkerAdditionalBenefitsCoverageDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerAdditionalBenefitsCoverageDataType
	var layout struct {
		*T
		OriginalCoverageBeginDateforCoverageTarget *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date_for_Coverage_Target,omitempty"`
	}
	layout.T = (*T)(t)
	layout.OriginalCoverageBeginDateforCoverageTarget = (*xsdDate)(layout.T.OriginalCoverageBeginDateforCoverageTarget)
	return e.EncodeElement(layout, start)
}
func (t *WorkerAdditionalBenefitsCoverageDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerAdditionalBenefitsCoverageDataType
	var overlay struct {
		*T
		OriginalCoverageBeginDateforCoverageTarget *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date_for_Coverage_Target,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.OriginalCoverageBeginDateforCoverageTarget = (*xsdDate)(overlay.T.OriginalCoverageBeginDateforCoverageTarget)
	return d.DecodeElement(&overlay, &start)
}

// Contains the additional benefits data for an employee.
type WorkerAdditionalBenefitsDataType struct {
	AdditionalBenefitsPeriodData []WorkerAdditionalBenefitsPeriodDataType `xml:"urn:com.workday/bsvc Additional_Benefits_Period_Data,omitempty"`
}

// Contains the additional benefits period data for an employee based on the benefit plan year.
type WorkerAdditionalBenefitsPeriodDataType struct {
	EnrollmentPeriodData           *EnrollmentPeriodDataType                  `xml:"urn:com.workday/bsvc Enrollment_Period_Data,omitempty"`
	AdditionalBenefitsCoverageData []WorkerAdditionalBenefitsCoverageDataType `xml:"urn:com.workday/bsvc Additional_Benefits_Coverage_Data,omitempty"`
}

// Worker Reference and Name of the Worker
type WorkerBasicDataType struct {
	WorkerReference  *WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	WorkerDescriptor string            `xml:"urn:com.workday/bsvc Worker_Descriptor,omitempty"`
}

// The election information for the coverage.
type WorkerBenefitElectionDataType struct {
	CoverageBeginDate                               *time.Time                  `xml:"urn:com.workday/bsvc Coverage_Begin_Date,omitempty"`
	CoverageEndDate                                 *time.Time                  `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
	ElectionCoverageBeginDate                       *time.Time                  `xml:"urn:com.workday/bsvc Election_Coverage_Begin_Date,omitempty"`
	OriginalCoverageBeginDate                       *time.Time                  `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
	OriginalCoverageBeginDateforBenefitCoverageType *time.Time                  `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date_for_Benefit_Coverage_Type,omitempty"`
	DeductionBeginDate                              *time.Time                  `xml:"urn:com.workday/bsvc Deduction_Begin_Date,omitempty"`
	DeductionEndDate                                *time.Time                  `xml:"urn:com.workday/bsvc Deduction_End_Date,omitempty"`
	ElectionStatus                                  string                      `xml:"urn:com.workday/bsvc Election_Status,omitempty"`
	EnrollmentSignatureDate                         *time.Time                  `xml:"urn:com.workday/bsvc Enrollment_Signature_Date,omitempty"`
	LatestEnrollmentSignatureDate                   *time.Time                  `xml:"urn:com.workday/bsvc Latest_Enrollment_Signature_Date,omitempty"`
	PassiveEnrollment                               *bool                       `xml:"urn:com.workday/bsvc Passive_Enrollment,omitempty"`
	BenefitPlanSummaryData                          *BenefitPlanSummaryDataType `xml:"urn:com.workday/bsvc Benefit_Plan_Summary_Data,omitempty"`
	OriginalPlanEnrollmentDate                      *time.Time                  `xml:"urn:com.workday/bsvc Original_Plan_Enrollment_Date,omitempty"`
	OriginalBenefitProviderEnrollmentDate           *time.Time                  `xml:"urn:com.workday/bsvc Original_Benefit_Provider_Enrollment_Date,omitempty"`
	IsCorrectedorRescinded                          bool                        `xml:"urn:com.workday/bsvc Is_Corrected_or_Rescinded,attr,omitempty"`
}

func (t *WorkerBenefitElectionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerBenefitElectionDataType
	var layout struct {
		*T
		CoverageBeginDate                               *xsdDate `xml:"urn:com.workday/bsvc Coverage_Begin_Date,omitempty"`
		CoverageEndDate                                 *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
		ElectionCoverageBeginDate                       *xsdDate `xml:"urn:com.workday/bsvc Election_Coverage_Begin_Date,omitempty"`
		OriginalCoverageBeginDate                       *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
		OriginalCoverageBeginDateforBenefitCoverageType *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date_for_Benefit_Coverage_Type,omitempty"`
		DeductionBeginDate                              *xsdDate `xml:"urn:com.workday/bsvc Deduction_Begin_Date,omitempty"`
		DeductionEndDate                                *xsdDate `xml:"urn:com.workday/bsvc Deduction_End_Date,omitempty"`
		EnrollmentSignatureDate                         *xsdDate `xml:"urn:com.workday/bsvc Enrollment_Signature_Date,omitempty"`
		LatestEnrollmentSignatureDate                   *xsdDate `xml:"urn:com.workday/bsvc Latest_Enrollment_Signature_Date,omitempty"`
		OriginalPlanEnrollmentDate                      *xsdDate `xml:"urn:com.workday/bsvc Original_Plan_Enrollment_Date,omitempty"`
		OriginalBenefitProviderEnrollmentDate           *xsdDate `xml:"urn:com.workday/bsvc Original_Benefit_Provider_Enrollment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CoverageBeginDate = (*xsdDate)(layout.T.CoverageBeginDate)
	layout.CoverageEndDate = (*xsdDate)(layout.T.CoverageEndDate)
	layout.ElectionCoverageBeginDate = (*xsdDate)(layout.T.ElectionCoverageBeginDate)
	layout.OriginalCoverageBeginDate = (*xsdDate)(layout.T.OriginalCoverageBeginDate)
	layout.OriginalCoverageBeginDateforBenefitCoverageType = (*xsdDate)(layout.T.OriginalCoverageBeginDateforBenefitCoverageType)
	layout.DeductionBeginDate = (*xsdDate)(layout.T.DeductionBeginDate)
	layout.DeductionEndDate = (*xsdDate)(layout.T.DeductionEndDate)
	layout.EnrollmentSignatureDate = (*xsdDate)(layout.T.EnrollmentSignatureDate)
	layout.LatestEnrollmentSignatureDate = (*xsdDate)(layout.T.LatestEnrollmentSignatureDate)
	layout.OriginalPlanEnrollmentDate = (*xsdDate)(layout.T.OriginalPlanEnrollmentDate)
	layout.OriginalBenefitProviderEnrollmentDate = (*xsdDate)(layout.T.OriginalBenefitProviderEnrollmentDate)
	return e.EncodeElement(layout, start)
}
func (t *WorkerBenefitElectionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerBenefitElectionDataType
	var overlay struct {
		*T
		CoverageBeginDate                               *xsdDate `xml:"urn:com.workday/bsvc Coverage_Begin_Date,omitempty"`
		CoverageEndDate                                 *xsdDate `xml:"urn:com.workday/bsvc Coverage_End_Date,omitempty"`
		ElectionCoverageBeginDate                       *xsdDate `xml:"urn:com.workday/bsvc Election_Coverage_Begin_Date,omitempty"`
		OriginalCoverageBeginDate                       *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date,omitempty"`
		OriginalCoverageBeginDateforBenefitCoverageType *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date_for_Benefit_Coverage_Type,omitempty"`
		DeductionBeginDate                              *xsdDate `xml:"urn:com.workday/bsvc Deduction_Begin_Date,omitempty"`
		DeductionEndDate                                *xsdDate `xml:"urn:com.workday/bsvc Deduction_End_Date,omitempty"`
		EnrollmentSignatureDate                         *xsdDate `xml:"urn:com.workday/bsvc Enrollment_Signature_Date,omitempty"`
		LatestEnrollmentSignatureDate                   *xsdDate `xml:"urn:com.workday/bsvc Latest_Enrollment_Signature_Date,omitempty"`
		OriginalPlanEnrollmentDate                      *xsdDate `xml:"urn:com.workday/bsvc Original_Plan_Enrollment_Date,omitempty"`
		OriginalBenefitProviderEnrollmentDate           *xsdDate `xml:"urn:com.workday/bsvc Original_Benefit_Provider_Enrollment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CoverageBeginDate = (*xsdDate)(overlay.T.CoverageBeginDate)
	overlay.CoverageEndDate = (*xsdDate)(overlay.T.CoverageEndDate)
	overlay.ElectionCoverageBeginDate = (*xsdDate)(overlay.T.ElectionCoverageBeginDate)
	overlay.OriginalCoverageBeginDate = (*xsdDate)(overlay.T.OriginalCoverageBeginDate)
	overlay.OriginalCoverageBeginDateforBenefitCoverageType = (*xsdDate)(overlay.T.OriginalCoverageBeginDateforBenefitCoverageType)
	overlay.DeductionBeginDate = (*xsdDate)(overlay.T.DeductionBeginDate)
	overlay.DeductionEndDate = (*xsdDate)(overlay.T.DeductionEndDate)
	overlay.EnrollmentSignatureDate = (*xsdDate)(overlay.T.EnrollmentSignatureDate)
	overlay.LatestEnrollmentSignatureDate = (*xsdDate)(overlay.T.LatestEnrollmentSignatureDate)
	overlay.OriginalPlanEnrollmentDate = (*xsdDate)(overlay.T.OriginalPlanEnrollmentDate)
	overlay.OriginalBenefitProviderEnrollmentDate = (*xsdDate)(overlay.T.OriginalBenefitProviderEnrollmentDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains the plans that the worker is eligible for
type WorkerBenefitEligibilityDataType struct {
	BenefitPlanReference     []BenefitPlanObjectType       `xml:"urn:com.workday/bsvc Benefit_Plan_Reference,omitempty"`
	PlanEligibilityDatesData []PlanEligiblityDatesDataType `xml:"urn:com.workday/bsvc Plan_Eligibility_Dates_Data,omitempty"`
}

// Contains the worker's benefits information.
//
// Security Note: This element is secured to the following domains:  Self Service: Benefit Elections; Worker Data: Benefit Elections
type WorkerBenefitEnrollmentDataType struct {
	HealthCareData           *WorkerHealthCareDataType           `xml:"urn:com.workday/bsvc Health_Care_Data,omitempty"`
	HealthSavingsAccountData *WorkerHealthSavingsAccountDataType `xml:"urn:com.workday/bsvc Health_Savings_Account_Data,omitempty"`
	SpendingAccountData      *WorkerSpendingAccountDataType      `xml:"urn:com.workday/bsvc Spending_Account_Data,omitempty"`
	InsuranceData            *WorkerInsuranceDataType            `xml:"urn:com.workday/bsvc Insurance_Data,omitempty"`
	RetirementSavingsData    *RetirementSavingsDataType          `xml:"urn:com.workday/bsvc Retirement_Savings_Data,omitempty"`
	AdditionalBenefitsData   *WorkerAdditionalBenefitsDataType   `xml:"urn:com.workday/bsvc Additional_Benefits_Data,omitempty"`
	COBRAEligibilityData     []COBRAEligibilityDataType          `xml:"urn:com.workday/bsvc COBRA_Eligibility_Data,omitempty"`
}

// Contains the current career information for an employee. This includes the travel and relocation preferences, job profiles the employee is interested in, and their career interests
type WorkerCareerDataType struct {
	RelocationData      *RelocationPreferenceDataType `xml:"urn:com.workday/bsvc Relocation_Data,omitempty"`
	TravelData          *TravelPreferenceDataType     `xml:"urn:com.workday/bsvc Travel_Data,omitempty"`
	JobInterestsData    *JobInterestsDataType         `xml:"urn:com.workday/bsvc Job_Interests_Data,omitempty"`
	CareerInterestsData *CareerPreferencesDataType    `xml:"urn:com.workday/bsvc Career_Interests_Data,omitempty"`
}

// Wrapper element for the worker compensation code. Contains the reference id and the worker comp code.
type WorkerCompensationCodeDataType struct {
	WorkersCompensationCodeReference *WorkersCompensationCodeObjectType `xml:"urn:com.workday/bsvc Workers_Compensation_Code_Reference,omitempty"`
	WorkersCompensationCode          string                             `xml:"urn:com.workday/bsvc Workers_Compensation_Code,omitempty"`
}

// Contains the worker's compensation information.
//
// Security Note: This element is secured to the following domains:  Self Service: Compensation; Worker Data: Compensation by Organization
type WorkerCompensationDataType struct {
	CompensationEffectiveDate       *time.Time                                    `xml:"urn:com.workday/bsvc Compensation_Effective_Date,omitempty"`
	ReasonReference                 *EventClassificationSubcategoryObjectType     `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
	CompensationGuidelinesData      []CompensatableGuidelinesDataType             `xml:"urn:com.workday/bsvc Compensation_Guidelines_Data,omitempty"`
	SalaryandHourlyData             []EmployeeBasePayPlanAssignmentDataType       `xml:"urn:com.workday/bsvc Salary_and_Hourly_Data,omitempty"`
	UnitSalaryPlanData              []EmployeeSalaryUnitPlanAssignmentDataType    `xml:"urn:com.workday/bsvc Unit_Salary_Plan_Data,omitempty"`
	AllowancePlanData               []EmployeeAllowancePlanAssignmentDataType     `xml:"urn:com.workday/bsvc Allowance_Plan_Data,omitempty"`
	UnitAllowancePlanData           []EmployeeAllowanceUnitPlanAssignmentDataType `xml:"urn:com.workday/bsvc Unit_Allowance_Plan_Data,omitempty"`
	BonusPlanData                   []EmployeeBonusPlanAssignmentDataType         `xml:"urn:com.workday/bsvc Bonus_Plan_Data,omitempty"`
	MeritPlanData                   []EmployeeMeritPlanAssignmentDataType         `xml:"urn:com.workday/bsvc Merit_Plan_Data,omitempty"`
	CommissionPlanData              []EmployeeCommissionPlanAssignmentDataType    `xml:"urn:com.workday/bsvc Commission_Plan_Data,omitempty"`
	StockPlanData                   []EmployeeStockPlanAssignmentDataType         `xml:"urn:com.workday/bsvc Stock_Plan_Data,omitempty"`
	FuturePaymentPlanData           []FuturePaymentPlanAssignmentDataType         `xml:"urn:com.workday/bsvc Future_Payment_Plan_Data,omitempty"`
	PeriodSalaryPlanData            []EmployeePeriodSalaryPlanAssignmentDataType  `xml:"urn:com.workday/bsvc Period_Salary_Plan_Data,omitempty"`
	EmployeeCompensationSummaryData *CompensatableSummaryDataType                 `xml:"urn:com.workday/bsvc Employee_Compensation_Summary_Data,omitempty"`
}

func (t *WorkerCompensationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerCompensationDataType
	var layout struct {
		*T
		CompensationEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Compensation_Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CompensationEffectiveDate = (*xsdDate)(layout.T.CompensationEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *WorkerCompensationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerCompensationDataType
	var overlay struct {
		*T
		CompensationEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Compensation_Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CompensationEffectiveDate = (*xsdDate)(overlay.T.CompensationEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing all Contract data for a Worker.
type WorkerContractDetailDataType struct {
	ContractPayRate           float64              `xml:"urn:com.workday/bsvc Contract_Pay_Rate,omitempty"`
	CurrencyReference         *CurrencyObjectType  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference        *FrequencyObjectType `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	ContractAssignmentDetails string               `xml:"urn:com.workday/bsvc Contract_Assignment_Details,omitempty"`
	ContractEndDate           *time.Time           `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
	SupplierReference         *SupplierObjectType  `xml:"urn:com.workday/bsvc Supplier_Reference,omitempty"`
}

func (t *WorkerContractDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerContractDetailDataType
	var layout struct {
		*T
		ContractEndDate *xsdDate `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ContractEndDate = (*xsdDate)(layout.T.ContractEndDate)
	return e.EncodeElement(layout, start)
}
func (t *WorkerContractDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerContractDetailDataType
	var overlay struct {
		*T
		ContractEndDate *xsdDate `xml:"urn:com.workday/bsvc Contract_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ContractEndDate = (*xsdDate)(overlay.T.ContractEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Effective dated additional data for a worker.  When updating custom objects that allow multiple instances, data for existing instances must appear in the request, otherwise it will be removed.
type WorkerCustomObjectDataType struct {
	EffectiveDate                time.Time                                   `xml:"urn:com.workday/bsvc Effective_Date"`
	WorkerReference              *WorkerObjectType                           `xml:"urn:com.workday/bsvc Worker_Reference"`
	BusinessObjectAdditionalData *EffectiveDatedWebServiceAdditionalDataType `xml:"urn:com.workday/bsvc Business_Object_Additional_Data"`
}

func (t *WorkerCustomObjectDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerCustomObjectDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *WorkerCustomObjectDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerCustomObjectDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains the detailed information about a worker.
type WorkerDataType struct {
	WorkerID                                 string                                                `xml:"urn:com.workday/bsvc Worker_ID,omitempty"`
	UserID                                   string                                                `xml:"urn:com.workday/bsvc User_ID,omitempty"`
	UniversalID                              string                                                `xml:"urn:com.workday/bsvc Universal_ID,omitempty"`
	PersonalData                             *PersonalInformationDataType                          `xml:"urn:com.workday/bsvc Personal_Data,omitempty"`
	EmploymentData                           *WorkerEmploymentInformationDataType                  `xml:"urn:com.workday/bsvc Employment_Data,omitempty"`
	CompensationData                         *WorkerCompensationDataType                           `xml:"urn:com.workday/bsvc Compensation_Data,omitempty"`
	OrganizationData                         *WorkerOrganizationDataType                           `xml:"urn:com.workday/bsvc Organization_Data,omitempty"`
	RoleData                                 *WorkerRoleDataType                                   `xml:"urn:com.workday/bsvc Role_Data,omitempty"`
	ManagementChainData                      *WorkerManagementChainDataType                        `xml:"urn:com.workday/bsvc Management_Chain_Data,omitempty"`
	BenefitEnrollmentData                    *WorkerBenefitEnrollmentDataType                      `xml:"urn:com.workday/bsvc Benefit_Enrollment_Data,omitempty"`
	BenefitEligibilityData                   *WorkerBenefitEligibilityDataType                     `xml:"urn:com.workday/bsvc Benefit_Eligibility_Data,omitempty"`
	RelatedPersonData                        *WorkerRelatedPersonsDataType                         `xml:"urn:com.workday/bsvc Related_Person_Data,omitempty"`
	QualificationData                        *PersonQualificationDataType                          `xml:"urn:com.workday/bsvc Qualification_Data,omitempty"`
	EmployeeReviewData                       *EmployeeReviewDataType                               `xml:"urn:com.workday/bsvc Employee_Review_Data,omitempty"`
	PhotoData                                *EmployeeImageDataType                                `xml:"urn:com.workday/bsvc Photo_Data,omitempty"`
	WorkerDocumentData                       *WorkerDocumentDataWWSType                            `xml:"urn:com.workday/bsvc Worker_Document_Data,omitempty"`
	IntegrationFieldOverrideData             []DocumentFieldResultDataType                         `xml:"urn:com.workday/bsvc Integration_Field_Override_Data,omitempty"`
	TransactionLogEntryData                  *EventTargetTransactionLogEntryDataType               `xml:"urn:com.workday/bsvc Transaction_Log_Entry_Data,omitempty"`
	TransactionLogCorrectedAndRescindedData  []EventTargetTransactionLogRescindAndCorrectDataType  `xml:"urn:com.workday/bsvc Transaction_Log_Corrected_And_Rescinded_Data,omitempty"`
	SuccessionProfileData                    *SuccessionProfileDataType                            `xml:"urn:com.workday/bsvc Succession_Profile_Data,omitempty"`
	TalentAssessmentData                     *TalentAssessmentDataType                             `xml:"urn:com.workday/bsvc Talent_Assessment_Data,omitempty"`
	WorkerGoalData                           *WorkerGoalDataType                                   `xml:"urn:com.workday/bsvc Worker_Goal_Data,omitempty"`
	DevelopmentItemData                      *WorkerDevelopmentItemDataType                        `xml:"urn:com.workday/bsvc Development_Item_Data,omitempty"`
	SkillData                                *WorkerSkillItemDataType                              `xml:"urn:com.workday/bsvc Skill_Data,omitempty"`
	EmployeeContractsData                    *EmployeeContractsDataType                            `xml:"urn:com.workday/bsvc Employee_Contracts_Data,omitempty"`
	ExtendedEmployeeContractsData            *ExtendedEmployeeContractsDataType                    `xml:"urn:com.workday/bsvc Extended_Employee_Contracts_Data,omitempty"`
	FeedbackReceivedData                     *FeedbackReceivedDataType                             `xml:"urn:com.workday/bsvc Feedback_Received_Data,omitempty"`
	UserAccountData                          *SystemUserforWorkerDataType                          `xml:"urn:com.workday/bsvc User_Account_Data,omitempty"`
	CareerData                               *WorkerCareerDataType                                 `xml:"urn:com.workday/bsvc Career_Data,omitempty"`
	AccountProvisioningData                  *PersonAccountProvisioningDataType                    `xml:"urn:com.workday/bsvc Account_Provisioning_Data,omitempty"`
	BackgroundCheckData                      []BackgroundCheckOverallStatusDataType                `xml:"urn:com.workday/bsvc Background_Check_Data,omitempty"`
	ContingentWorkerTaxAuthorityFormTypeData []ContingentWorkerTaxAuthorityFormInformationDataType `xml:"urn:com.workday/bsvc Contingent_Worker_Tax_Authority_Form_Type_Data,omitempty"`
}

// Contains all Development Items for the referenced Worker
type WorkerDevelopmentItemDataType struct {
	WorkerDevelopmentItem []WorkerDevelopmentItemType `xml:"urn:com.workday/bsvc Worker_Development_Item,omitempty"`
}

// Contains the reference to the Development Item and the Development Item details
type WorkerDevelopmentItemType struct {
	DevelopmentItemReference *DevelopmentItemObjectType      `xml:"urn:com.workday/bsvc Development_Item_Reference,omitempty"`
	DevelopmentItemData      *DevelopmentItemVersionDataType `xml:"urn:com.workday/bsvc Development_Item_Data,omitempty"`
}

// The information about the worker document, such as the category and file.
type WorkerDocumentDataType struct {
	ID                        string                         `xml:"urn:com.workday/bsvc ID,omitempty"`
	Filename                  string                         `xml:"urn:com.workday/bsvc Filename"`
	Comment                   string                         `xml:"urn:com.workday/bsvc Comment,omitempty"`
	File                      *[]byte                        `xml:"urn:com.workday/bsvc File,omitempty"`
	DocumentCategoryReference *DocumentCategoryAllObjectType `xml:"urn:com.workday/bsvc Document_Category_Reference"`
	WorkerReference           *WorkerObjectType              `xml:"urn:com.workday/bsvc Worker_Reference"`
	ContentType               string                         `xml:"urn:com.workday/bsvc Content_Type,omitempty"`
}

func (t *WorkerDocumentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerDocumentDataType
	var layout struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	layout.T = (*T)(t)
	layout.File = (*xsdBase64Binary)(layout.T.File)
	return e.EncodeElement(layout, start)
}
func (t *WorkerDocumentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerDocumentDataType
	var overlay struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.File = (*xsdBase64Binary)(overlay.T.File)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for documents associated with a worker.
type WorkerDocumentDataWWSType struct {
	WorkerDocument []WorkerDocumentWWSType `xml:"urn:com.workday/bsvc Worker_Document,omitempty"`
}

// Wrapper element for the details of a worker document.
type WorkerDocumentDetailsDataType struct {
	DocumentCategoryReference *DocumentCategoryAllObjectType `xml:"urn:com.workday/bsvc Document_Category_Reference"`
	Filename                  string                         `xml:"urn:com.workday/bsvc Filename"`
	Comment                   string                         `xml:"urn:com.workday/bsvc Comment,omitempty"`
	File                      *[]byte                        `xml:"urn:com.workday/bsvc File,omitempty"`
}

func (t *WorkerDocumentDetailsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerDocumentDetailsDataType
	var layout struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	layout.T = (*T)(t)
	layout.File = (*xsdBase64Binary)(layout.T.File)
	return e.EncodeElement(layout, start)
}
func (t *WorkerDocumentDetailsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerDocumentDetailsDataType
	var overlay struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.File = (*xsdBase64Binary)(overlay.T.File)
	return d.DecodeElement(&overlay, &start)
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

// The Worker Document reference(s) to be retrieved.
type WorkerDocumentRequestReferencesType struct {
	WorkerDocumentReference []WorkerDocumentObjectType `xml:"urn:com.workday/bsvc Worker_Document_Reference"`
}

// Worker Document Response Data containing the retrieved Worker Document information.
type WorkerDocumentResponseDataType struct {
	WorkerDocument []WorkerDocumentType `xml:"urn:com.workday/bsvc Worker_Document,omitempty"`
}

// Worker Document Response Group which controls the type and amount of data returned.
type WorkerDocumentResponseGroupType struct {
	IncludeReference          *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeWorkerDocumentData *bool `xml:"urn:com.workday/bsvc Include_Worker_Document_Data,omitempty"`
}

// Contains the retrieved Worker Document information.
type WorkerDocumentType struct {
	WorkerDocumentReference *WorkerDocumentObjectType `xml:"urn:com.workday/bsvc Worker_Document_Reference,omitempty"`
	WorkerDocumentData      *WorkerDocumentDataType   `xml:"urn:com.workday/bsvc Worker_Document_Data,omitempty"`
}

// Wrapper element for documents associated with a worker.
type WorkerDocumentWWSType struct {
	WorkerDocumentReference  *WorkerDocumentObjectType      `xml:"urn:com.workday/bsvc Worker_Document_Reference"`
	WorkerDocumentDetailData *WorkerDocumentDetailsDataType `xml:"urn:com.workday/bsvc Worker_Document_Detail_Data"`
}

// Attachments for the Staffing Event entered via a web service.
type WorkerDocumentforStaffingEventDataType struct {
	FileName                  string                         `xml:"urn:com.workday/bsvc File_Name"`
	WorkerDocumentComment     string                         `xml:"urn:com.workday/bsvc Worker_Document_Comment,omitempty"`
	DocumentCategoryReference *DocumentCategoryAllObjectType `xml:"urn:com.workday/bsvc Document_Category_Reference"`
	File                      *[]byte                        `xml:"urn:com.workday/bsvc File,omitempty"`
	ContentType               string                         `xml:"urn:com.workday/bsvc Content_Type,omitempty"`
}

func (t *WorkerDocumentforStaffingEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerDocumentforStaffingEventDataType
	var layout struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	layout.T = (*T)(t)
	layout.File = (*xsdBase64Binary)(layout.T.File)
	return e.EncodeElement(layout, start)
}
func (t *WorkerDocumentforStaffingEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerDocumentforStaffingEventDataType
	var overlay struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.File = (*xsdBase64Binary)(overlay.T.File)
	return d.DecodeElement(&overlay, &start)
}

// Contains the worker's employment information, such as their position and job information.
//
// Security Note: This element is secured to the following domains:  Self-Service: Current Staffing Information, Worker Data: Current Staffing Information
type WorkerEmploymentInformationDataType struct {
	WorkerJobData                      []WorkerJobDataType                     `xml:"urn:com.workday/bsvc Worker_Job_Data,omitempty"`
	WorkerStatusData                   *WorkerStatusDetailDataType             `xml:"urn:com.workday/bsvc Worker_Status_Data,omitempty"`
	WorkerContractData                 *WorkerContractDetailDataType           `xml:"urn:com.workday/bsvc Worker_Contract_Data,omitempty"`
	InternationalAssignmentSummaryData *InternationalAssignmentSummaryDataType `xml:"urn:com.workday/bsvc International_Assignment_Summary_Data,omitempty"`
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

// Contains all the Goal Details (Review Goals and Worker Goal Details) for a Worker.
type WorkerGoalDataType struct {
	Goal []WorkerGoalType `xml:"urn:com.workday/bsvc Goal,omitempty"`
}

// Contains the reference to the Goal Detail and all the Goal Data payload.
type WorkerGoalType struct {
	GoalReference *GoalObjectType      `xml:"urn:com.workday/bsvc Goal_Reference,omitempty"`
	GoalData      []GoalDetailDataType `xml:"urn:com.workday/bsvc Goal_Data,omitempty"`
}

// Contains the health care coverage information (elections) for an employee.
type WorkerHealthCareCoverageDataType struct {
	HealthCareCoverageTargetReference          *HealthCareCoverageTargetObjectType `xml:"urn:com.workday/bsvc Health_Care_Coverage_Target_Reference,omitempty"`
	OriginalCoverageBeginDateforCoverageTarget *time.Time                          `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date_for_Coverage_Target,omitempty"`
	ProviderID                                 string                              `xml:"urn:com.workday/bsvc Provider_ID,omitempty"`
	BenefitElectionData                        *WorkerBenefitElectionDataType      `xml:"urn:com.workday/bsvc Benefit_Election_Data,omitempty"`
	DependentCoverageData                      []DependentCoverageDataType         `xml:"urn:com.workday/bsvc Dependent_Coverage_Data,omitempty"`
}

func (t *WorkerHealthCareCoverageDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerHealthCareCoverageDataType
	var layout struct {
		*T
		OriginalCoverageBeginDateforCoverageTarget *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date_for_Coverage_Target,omitempty"`
	}
	layout.T = (*T)(t)
	layout.OriginalCoverageBeginDateforCoverageTarget = (*xsdDate)(layout.T.OriginalCoverageBeginDateforCoverageTarget)
	return e.EncodeElement(layout, start)
}
func (t *WorkerHealthCareCoverageDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerHealthCareCoverageDataType
	var overlay struct {
		*T
		OriginalCoverageBeginDateforCoverageTarget *xsdDate `xml:"urn:com.workday/bsvc Original_Coverage_Begin_Date_for_Coverage_Target,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.OriginalCoverageBeginDateforCoverageTarget = (*xsdDate)(overlay.T.OriginalCoverageBeginDateforCoverageTarget)
	return d.DecodeElement(&overlay, &start)
}

// Contains the health care data for an employee.
type WorkerHealthCareDataType struct {
	HealthCarePeriodData []WorkerHealthCarePeriodDataType `xml:"urn:com.workday/bsvc Health_Care_Period_Data,omitempty"`
}

// Contains the health care period data for an employee based on the benefit plan year.
type WorkerHealthCarePeriodDataType struct {
	EnrollmentPeriodData   *EnrollmentPeriodDataType          `xml:"urn:com.workday/bsvc Enrollment_Period_Data,omitempty"`
	HealthCareCoverageData []WorkerHealthCareCoverageDataType `xml:"urn:com.workday/bsvc Health_Care_Coverage_Data,omitempty"`
}

// Contains the health savings account information for an employee.
type WorkerHealthSavingsAccountDataType struct {
	HealthSavingsAccountPeriodData []WorkerHealthSavingsAccountPeriodDataType `xml:"urn:com.workday/bsvc Health_Savings_Account_Period_Data,omitempty"`
}

// Contains the health saving account period data for an employee based on the benefit plan year.
type WorkerHealthSavingsAccountPeriodDataType struct {
	EnrollmentPeriodData             *EnrollmentPeriodDataType              `xml:"urn:com.workday/bsvc Enrollment_Period_Data,omitempty"`
	HealthSavingsAccountCoverageData []HealthSavingsAccountCoverageDataType `xml:"urn:com.workday/bsvc Health_Savings_Account_Coverage_Data,omitempty"`
}

// Contains the insurance election information for the employee.
type WorkerInsuranceCoverageDataType struct {
	InsuranceCoverageLevelData *InsuranceCoverageLevelDataType  `xml:"urn:com.workday/bsvc Insurance_Coverage_Level_Data"`
	BenefitElectionData        *WorkerBenefitElectionDataType   `xml:"urn:com.workday/bsvc Benefit_Election_Data"`
	DependentCoverageData      []DependentCoverageDataType      `xml:"urn:com.workday/bsvc Dependent_Coverage_Data,omitempty"`
	BeneficiaryDesignationData []BeneficiaryDesignationDataType `xml:"urn:com.workday/bsvc Beneficiary_Designation_Data,omitempty"`
}

// Contains the insurance information for an employee.
type WorkerInsuranceDataType struct {
	InsurancePeriodData []WorkerInsurancePeriodDataType `xml:"urn:com.workday/bsvc Insurance_Period_Data,omitempty"`
}

// Contains the insurance period data for an employee based on the benefit plan year.
type WorkerInsurancePeriodDataType struct {
	EnrollmentPeriodData  *EnrollmentPeriodDataType         `xml:"urn:com.workday/bsvc Enrollment_Period_Data,omitempty"`
	InsuranceCoverageData []WorkerInsuranceCoverageDataType `xml:"urn:com.workday/bsvc Insurance_Coverage_Data,omitempty"`
}

// Contains the Position-level data for a Worker.  The primary position is always the first element.
type WorkerJobDataType struct {
	PositionData                 *PositionDetailDataType           `xml:"urn:com.workday/bsvc Position_Data,omitempty"`
	PositionOrganizationsData    *PositionOrganizationDataType     `xml:"urn:com.workday/bsvc Position_Organizations_Data,omitempty"`
	PositionManagementChainsData *PositionManagementChainsDataType `xml:"urn:com.workday/bsvc Position_Management_Chains_Data,omitempty"`
	PrimaryJob                   bool                              `xml:"urn:com.workday/bsvc Primary_Job,attr,omitempty"`
}

// Contains the worker's management chain for supervisory and matrix organizations.
type WorkerManagementChainDataType struct {
	WorkerSupervisoryManagementChainData *WorkerSupervisoryManagementChainDataType `xml:"urn:com.workday/bsvc Worker_Supervisory_Management_Chain_Data,omitempty"`
	WorkerMatrixManagementChainData      *WorkerMatrixManagementChainDataType      `xml:"urn:com.workday/bsvc Worker_Matrix_Management_Chain_Data,omitempty"`
}

// Contains the worker's matrix management chain.
type WorkerMatrixManagementChainDataType struct {
	ManagementChainData []ManagementChainDataType `xml:"urn:com.workday/bsvc Management_Chain_Data,omitempty"`
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

// Contains the worker's organizations that they are a member of.
type WorkerOrganizationDataType struct {
	WorkerOrganizationData []WorkerOrganizationMembershipDataType `xml:"urn:com.workday/bsvc Worker_Organization_Data,omitempty"`
}

// Contains the organization reference and details about an organization.
type WorkerOrganizationMembershipDataType struct {
	OrganizationReference *PositionSetObjectType       `xml:"urn:com.workday/bsvc Organization_Reference"`
	OrganizationData      *OrganizationSummaryDataType `xml:"urn:com.workday/bsvc Organization_Data"`
}

// Contains the workers that fill the organization role.
type WorkerOrganizationRoleAssignmentDataType struct {
	WorkerReference *WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	AssignmentFrom  string            `xml:"urn:com.workday/bsvc Assignment_From,omitempty"`
}

// Contains the organization role and the organizations that the worker supports that role for.
type WorkerOrganizationRoleDataType struct {
	OrganizationRole []OrganizationRoleforWorkerDataType `xml:"urn:com.workday/bsvc Organization_Role,omitempty"`
}

// Contains the person's that are related to the worker as a dependent, beneficiary or emergency contact.
//
// Security Note: This element is secured to the following domains:  Self Service: Benefit Elections; Worker Data: Benefit Elections
type WorkerRelatedPersonsDataType struct {
	NumberofPayrollDependents float64             `xml:"urn:com.workday/bsvc Number_of_Payroll_Dependents,omitempty"`
	RelatedPerson             []RelatedPersonType `xml:"urn:com.workday/bsvc Related_Person,omitempty"`
}

// The Request Criteria element lets you apply additional criteria to identify the specific instance(s) of a Worker.
type WorkerRequestCriteriaType struct {
	TransactionLogCriteriaData      []TransactionLogCriteriaType            `xml:"urn:com.workday/bsvc Transaction_Log_Criteria_Data,omitempty"`
	OrganizationReference           []OrganizationObjectType                `xml:"urn:com.workday/bsvc Organization_Reference,omitempty"`
	CountryReference                []CountryObjectType                     `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	IncludeSubordinateOrganizations *bool                                   `xml:"urn:com.workday/bsvc Include_Subordinate_Organizations,omitempty"`
	PositionReference               []PositionElementObjectType             `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	EventReference                  *TransactionLogObjectType               `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	BenefitPlanReference            []BenefitPlanObjectType                 `xml:"urn:com.workday/bsvc Benefit_Plan_Reference,omitempty"`
	FieldAndParameterCriteriaData   *FieldAndParameterCriteriaDataType      `xml:"urn:com.workday/bsvc Field_And_Parameter_Criteria_Data,omitempty"`
	NationalIDCriteriaData          []WorkerbyNationalIDRequestCriteriaType `xml:"urn:com.workday/bsvc National_ID_Criteria_Data,omitempty"`
	ExcludeInactiveWorkers          *bool                                   `xml:"urn:com.workday/bsvc Exclude_Inactive_Workers,omitempty"`
	ExcludeEmployees                *bool                                   `xml:"urn:com.workday/bsvc Exclude_Employees,omitempty"`
	ExcludeContingentWorkers        *bool                                   `xml:"urn:com.workday/bsvc Exclude_Contingent_Workers,omitempty"`
	EligibilityCriteriaData         []EligibilityCriteriaDataType           `xml:"urn:com.workday/bsvc Eligibility_Criteria_Data,omitempty"`
	UniversalIDReference            []UniversalIdentifierObjectType         `xml:"urn:com.workday/bsvc Universal_ID_Reference,omitempty"`
}

// Utilize the Request References element to retrieve a specific instance(s) of Worker and its associated data.
type WorkerRequestReferencesType struct {
	WorkerReference          []WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference"`
	SkipNonExistingInstances bool               `xml:"urn:com.workday/bsvc Skip_Non_Existing_Instances,attr,omitempty"`
	IgnoreInvalidReferences  bool               `xml:"urn:com.workday/bsvc Ignore_Invalid_References,attr,omitempty"`
}

// Use the response group to limit the response to the data you are interested in. If the request does not set any values for the response group, then the response by default returns the following elements: Reference, Personal Data, Employment Data, Compensation Data, Organization Data, and Role Data.
type WorkerResponseGroupType struct {
	IncludeReference                                   *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludePersonalInformation                         *bool `xml:"urn:com.workday/bsvc Include_Personal_Information,omitempty"`
	IncludeAdditionalJobs                              *bool `xml:"urn:com.workday/bsvc Include_Additional_Jobs,omitempty"`
	IncludeEmploymentInformation                       *bool `xml:"urn:com.workday/bsvc Include_Employment_Information,omitempty"`
	IncludeCompensation                                *bool `xml:"urn:com.workday/bsvc Include_Compensation,omitempty"`
	IncludeOrganizations                               *bool `xml:"urn:com.workday/bsvc Include_Organizations,omitempty"`
	ExcludeOrganizationSupportRoleData                 *bool `xml:"urn:com.workday/bsvc Exclude_Organization_Support_Role_Data,omitempty"`
	ExcludeLocationHierarchies                         *bool `xml:"urn:com.workday/bsvc Exclude_Location_Hierarchies,omitempty"`
	ExcludeCostCenters                                 *bool `xml:"urn:com.workday/bsvc Exclude_Cost_Centers,omitempty"`
	ExcludeCostCenterHierarchies                       *bool `xml:"urn:com.workday/bsvc Exclude_Cost_Center_Hierarchies,omitempty"`
	ExcludeCompanies                                   *bool `xml:"urn:com.workday/bsvc Exclude_Companies,omitempty"`
	ExcludeCompanyHierarchies                          *bool `xml:"urn:com.workday/bsvc Exclude_Company_Hierarchies,omitempty"`
	ExcludeMatrixOrganizations                         *bool `xml:"urn:com.workday/bsvc Exclude_Matrix_Organizations,omitempty"`
	ExcludePayGroups                                   *bool `xml:"urn:com.workday/bsvc Exclude_Pay_Groups,omitempty"`
	ExcludeRegions                                     *bool `xml:"urn:com.workday/bsvc Exclude_Regions,omitempty"`
	ExcludeRegionHierarchies                           *bool `xml:"urn:com.workday/bsvc Exclude_Region_Hierarchies,omitempty"`
	ExcludeSupervisoryOrganizations                    *bool `xml:"urn:com.workday/bsvc Exclude_Supervisory_Organizations,omitempty"`
	ExcludeTeams                                       *bool `xml:"urn:com.workday/bsvc Exclude_Teams,omitempty"`
	ExcludeCustomOrganizations                         *bool `xml:"urn:com.workday/bsvc Exclude_Custom_Organizations,omitempty"`
	IncludeRoles                                       *bool `xml:"urn:com.workday/bsvc Include_Roles,omitempty"`
	IncludeManagementChainData                         *bool `xml:"urn:com.workday/bsvc Include_Management_Chain_Data,omitempty"`
	IncludeMultipleManagersinManagementChainData       *bool `xml:"urn:com.workday/bsvc Include_Multiple_Managers_in_Management_Chain_Data,omitempty"`
	IncludeBenefitEnrollments                          *bool `xml:"urn:com.workday/bsvc Include_Benefit_Enrollments,omitempty"`
	IncludeBenefitEligibility                          *bool `xml:"urn:com.workday/bsvc Include_Benefit_Eligibility,omitempty"`
	IncludeRelatedPersons                              *bool `xml:"urn:com.workday/bsvc Include_Related_Persons,omitempty"`
	IncludeQualifications                              *bool `xml:"urn:com.workday/bsvc Include_Qualifications,omitempty"`
	IncludeEmployeeReview                              *bool `xml:"urn:com.workday/bsvc Include_Employee_Review,omitempty"`
	IncludeGoals                                       *bool `xml:"urn:com.workday/bsvc Include_Goals,omitempty"`
	IncludeDevelopmentItems                            *bool `xml:"urn:com.workday/bsvc Include_Development_Items,omitempty"`
	IncludeSkills                                      *bool `xml:"urn:com.workday/bsvc Include_Skills,omitempty"`
	IncludePhoto                                       *bool `xml:"urn:com.workday/bsvc Include_Photo,omitempty"`
	IncludeWorkerDocuments                             *bool `xml:"urn:com.workday/bsvc Include_Worker_Documents,omitempty"`
	IncludeTransactionLogData                          *bool `xml:"urn:com.workday/bsvc Include_Transaction_Log_Data,omitempty"`
	IncludeSubeventsforCorrectedTransaction            *bool `xml:"urn:com.workday/bsvc Include_Subevents_for_Corrected_Transaction,omitempty"`
	IncludeSubeventsforRescindedTransaction            *bool `xml:"urn:com.workday/bsvc Include_Subevents_for_Rescinded_Transaction,omitempty"`
	IncludeSuccessionProfile                           *bool `xml:"urn:com.workday/bsvc Include_Succession_Profile,omitempty"`
	IncludeTalentAssessment                            *bool `xml:"urn:com.workday/bsvc Include_Talent_Assessment,omitempty"`
	IncludeEmployeeContractData                        *bool `xml:"urn:com.workday/bsvc Include_Employee_Contract_Data,omitempty"`
	IncludeContractsforTerminatedWorkers               *bool `xml:"urn:com.workday/bsvc Include_Contracts_for_Terminated_Workers,omitempty"`
	IncludeCollectiveAgreementData                     *bool `xml:"urn:com.workday/bsvc Include_Collective_Agreement_Data,omitempty"`
	IncludeProbationPeriodData                         *bool `xml:"urn:com.workday/bsvc Include_Probation_Period_Data,omitempty"`
	IncludeExtendedEmployeeContractDetails             *bool `xml:"urn:com.workday/bsvc Include_Extended_Employee_Contract_Details,omitempty"`
	IncludeFeedbackReceived                            *bool `xml:"urn:com.workday/bsvc Include_Feedback_Received,omitempty"`
	IncludeUserAccount                                 *bool `xml:"urn:com.workday/bsvc Include_User_Account,omitempty"`
	IncludeCareer                                      *bool `xml:"urn:com.workday/bsvc Include_Career,omitempty"`
	IncludeAccountProvisioning                         *bool `xml:"urn:com.workday/bsvc Include_Account_Provisioning,omitempty"`
	IncludeBackgroundCheckData                         *bool `xml:"urn:com.workday/bsvc Include_Background_Check_Data,omitempty"`
	IncludeContingentWorkerTaxAuthorityFormInformation *bool `xml:"urn:com.workday/bsvc Include_Contingent_Worker_Tax_Authority_Form_Information,omitempty"`
	ExcludeFunds                                       *bool `xml:"urn:com.workday/bsvc Exclude_Funds,omitempty"`
	ExcludeFundHierarchies                             *bool `xml:"urn:com.workday/bsvc Exclude_Fund_Hierarchies,omitempty"`
	ExcludeGrants                                      *bool `xml:"urn:com.workday/bsvc Exclude_Grants,omitempty"`
	ExcludeGrantHierarchies                            *bool `xml:"urn:com.workday/bsvc Exclude_Grant_Hierarchies,omitempty"`
	ExcludeBusinessUnits                               *bool `xml:"urn:com.workday/bsvc Exclude_Business_Units,omitempty"`
	ExcludeBusinessUnitHierarchies                     *bool `xml:"urn:com.workday/bsvc Exclude_Business_Unit_Hierarchies,omitempty"`
	ExcludePrograms                                    *bool `xml:"urn:com.workday/bsvc Exclude_Programs,omitempty"`
	ExcludeProgramHierarchies                          *bool `xml:"urn:com.workday/bsvc Exclude_Program_Hierarchies,omitempty"`
	ExcludeGifts                                       *bool `xml:"urn:com.workday/bsvc Exclude_Gifts,omitempty"`
	ExcludeGiftHierarchies                             *bool `xml:"urn:com.workday/bsvc Exclude_Gift_Hierarchies,omitempty"`
}

// Contains the retirement savings election information for the employee.
type WorkerRetirementSavingsDataType struct {
	BenefitElectionData                *WorkerBenefitElectionDataType          `xml:"urn:com.workday/bsvc Benefit_Election_Data"`
	EmployeeContributionPercentageData *EmployeeContributionPercentageDataType `xml:"urn:com.workday/bsvc Employee_Contribution_Percentage_Data,omitempty"`
	EmployeeContributionAmountData     *EmployeeContributionAmountDataType     `xml:"urn:com.workday/bsvc Employee_Contribution_Amount_Data,omitempty"`
	EmployerContributionPercentageData *EmployerContributionPercentageDataType `xml:"urn:com.workday/bsvc Employer_Contribution_Percentage_Data,omitempty"`
	EmployerContributionAmountData     *EmployerContributionAmountDataType     `xml:"urn:com.workday/bsvc Employer_Contribution_Amount_Data,omitempty"`
	BeneficiaryDesignationData         []BeneficiaryDesignationDataType        `xml:"urn:com.workday/bsvc Beneficiary_Designation_Data,omitempty"`
}

// Contains the retirement savings election information for the benefit plan year period.
type WorkerRetirementSavingsPeriodDataType struct {
	EnrollmentPeriodData          *EnrollmentPeriodDataType         `xml:"urn:com.workday/bsvc Enrollment_Period_Data,omitempty"`
	RetirementSavingsCoverageData []WorkerRetirementSavingsDataType `xml:"urn:com.workday/bsvc Retirement_Savings_Coverage_Data,omitempty"`
}

// Contains the roles that a worker holds.
type WorkerRoleDataType struct {
	OrganizationRoleData *WorkerOrganizationRoleDataType `xml:"urn:com.workday/bsvc Organization_Role_Data,omitempty"`
}

// Creates all Skills for the referenced Worker
type WorkerSkillItemDataType struct {
	WorkerSkillItem []SkillType `xml:"urn:com.workday/bsvc Worker_Skill_Item,omitempty"`
}

// Contains the spending account information for an employee.
type WorkerSpendingAccountDataType struct {
	SpendingAccountPeriodData []WorkerSpendingAccountPeriodDataType `xml:"urn:com.workday/bsvc Spending_Account_Period_Data,omitempty"`
}

// Contains the spending account period data for an employee based on the benefit plan year.
type WorkerSpendingAccountPeriodDataType struct {
	EnrollmentPeriodData        *EnrollmentPeriodDataType         `xml:"urn:com.workday/bsvc Enrollment_Period_Data,omitempty"`
	SpendingAccountCoverageData []SpendingAccountCoverageDataType `xml:"urn:com.workday/bsvc Spending_Account_Coverage_Data,omitempty"`
}

// Encapsulating element containing all Status data for a Worker.
type WorkerStatusDetailDataType struct {
	Active                                        *bool                                  `xml:"urn:com.workday/bsvc Active,omitempty"`
	ActiveStatusDate                              *time.Time                             `xml:"urn:com.workday/bsvc Active_Status_Date,omitempty"`
	HireDate                                      *time.Time                             `xml:"urn:com.workday/bsvc Hire_Date,omitempty"`
	OriginalHireDate                              *time.Time                             `xml:"urn:com.workday/bsvc Original_Hire_Date,omitempty"`
	HireReasonReference                           *GeneralEventSubcategoryObjectType     `xml:"urn:com.workday/bsvc Hire_Reason_Reference,omitempty"`
	EndEmploymentDate                             *time.Time                             `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
	ContinuousServiceDate                         *time.Time                             `xml:"urn:com.workday/bsvc Continuous_Service_Date,omitempty"`
	FirstDayofWork                                *time.Time                             `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
	ExpectedRetirementDate                        *time.Time                             `xml:"urn:com.workday/bsvc Expected_Retirement_Date,omitempty"`
	RetirementEligibilityDate                     *time.Time                             `xml:"urn:com.workday/bsvc Retirement_Eligibility_Date,omitempty"`
	Retired                                       *bool                                  `xml:"urn:com.workday/bsvc Retired,omitempty"`
	RetirementDate                                *time.Time                             `xml:"urn:com.workday/bsvc Retirement_Date,omitempty"`
	SeniorityDate                                 *time.Time                             `xml:"urn:com.workday/bsvc Seniority_Date,omitempty"`
	SeveranceDate                                 *time.Time                             `xml:"urn:com.workday/bsvc Severance_Date,omitempty"`
	BenefitsServiceDate                           *time.Time                             `xml:"urn:com.workday/bsvc Benefits_Service_Date,omitempty"`
	CompanyServiceDate                            *time.Time                             `xml:"urn:com.workday/bsvc Company_Service_Date,omitempty"`
	TimeOffServiceDate                            *time.Time                             `xml:"urn:com.workday/bsvc Time_Off_Service_Date,omitempty"`
	VestingDate                                   *time.Time                             `xml:"urn:com.workday/bsvc Vesting_Date,omitempty"`
	DateEnteredWorkforce                          *time.Time                             `xml:"urn:com.workday/bsvc Date_Entered_Workforce,omitempty"`
	DaysUnemployed                                float64                                `xml:"urn:com.workday/bsvc Days_Unemployed,omitempty"`
	MonthsContinuousPriorEmployment               float64                                `xml:"urn:com.workday/bsvc Months_Continuous_Prior_Employment,omitempty"`
	Terminated                                    *bool                                  `xml:"urn:com.workday/bsvc Terminated,omitempty"`
	TerminationDate                               *time.Time                             `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
	PayThroughDate                                *time.Time                             `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
	PrimaryTerminationReasonReference             *TerminationSubcategoryObjectType      `xml:"urn:com.workday/bsvc Primary_Termination_Reason_Reference,omitempty"`
	PrimaryTerminationCategoryReference           *EventClassificationCategoryObjectType `xml:"urn:com.workday/bsvc Primary_Termination_Category_Reference,omitempty"`
	TerminationInvoluntary                        *bool                                  `xml:"urn:com.workday/bsvc Termination_Involuntary,omitempty"`
	SecondaryTerminationReasonsData               []SecondaryTerminationDataType         `xml:"urn:com.workday/bsvc Secondary_Termination_Reasons_Data,omitempty"`
	LocalTerminationReasonReference               *LocalTerminationReasonObjectType      `xml:"urn:com.workday/bsvc Local_Termination_Reason_Reference,omitempty"`
	EligibleforHireReference                      *CommonYesNoObjectType                 `xml:"urn:com.workday/bsvc Eligible_for_Hire_Reference,omitempty"`
	RegrettableTermination                        *bool                                  `xml:"urn:com.workday/bsvc Regrettable_Termination,omitempty"`
	EligibleforRehireonLatestTerminationReference *CommonYesNoObjectType                 `xml:"urn:com.workday/bsvc Eligible_for_Rehire_on_Latest_Termination_Reference,omitempty"`
	HireRescinded                                 *bool                                  `xml:"urn:com.workday/bsvc Hire_Rescinded,omitempty"`
	TerminationLastDayofWork                      *time.Time                             `xml:"urn:com.workday/bsvc Termination_Last_Day_of_Work,omitempty"`
	ResignationDate                               *time.Time                             `xml:"urn:com.workday/bsvc Resignation_Date,omitempty"`
	LastDateforWhichPaid                          *time.Time                             `xml:"urn:com.workday/bsvc Last_Date_for_Which_Paid,omitempty"`
	ExpectedDateofReturn                          *time.Time                             `xml:"urn:com.workday/bsvc Expected_Date_of_Return,omitempty"`
	NotReturning                                  *bool                                  `xml:"urn:com.workday/bsvc Not_Returning,omitempty"`
	ReturnUnknown                                 *bool                                  `xml:"urn:com.workday/bsvc Return_Unknown,omitempty"`
	ProbationStartDate                            *time.Time                             `xml:"urn:com.workday/bsvc Probation_Start_Date,omitempty"`
	ProbationEndDate                              *time.Time                             `xml:"urn:com.workday/bsvc Probation_End_Date,omitempty"`
	LeaveStatusData                               []LeaveStatusDetailDataType            `xml:"urn:com.workday/bsvc Leave_Status_Data,omitempty"`
	LeaveRequestsCorrectedData                    []LeaveRequestsCorrectedDetailDataType `xml:"urn:com.workday/bsvc Leave_Requests_Corrected_Data,omitempty"`
	AcademicTenureDate                            *time.Time                             `xml:"urn:com.workday/bsvc Academic_Tenure_Date,omitempty"`
	Rehire                                        *bool                                  `xml:"urn:com.workday/bsvc Rehire,omitempty"`
}

func (t *WorkerStatusDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkerStatusDetailDataType
	var layout struct {
		*T
		ActiveStatusDate          *xsdDate `xml:"urn:com.workday/bsvc Active_Status_Date,omitempty"`
		HireDate                  *xsdDate `xml:"urn:com.workday/bsvc Hire_Date,omitempty"`
		OriginalHireDate          *xsdDate `xml:"urn:com.workday/bsvc Original_Hire_Date,omitempty"`
		EndEmploymentDate         *xsdDate `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
		ContinuousServiceDate     *xsdDate `xml:"urn:com.workday/bsvc Continuous_Service_Date,omitempty"`
		FirstDayofWork            *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
		ExpectedRetirementDate    *xsdDate `xml:"urn:com.workday/bsvc Expected_Retirement_Date,omitempty"`
		RetirementEligibilityDate *xsdDate `xml:"urn:com.workday/bsvc Retirement_Eligibility_Date,omitempty"`
		RetirementDate            *xsdDate `xml:"urn:com.workday/bsvc Retirement_Date,omitempty"`
		SeniorityDate             *xsdDate `xml:"urn:com.workday/bsvc Seniority_Date,omitempty"`
		SeveranceDate             *xsdDate `xml:"urn:com.workday/bsvc Severance_Date,omitempty"`
		BenefitsServiceDate       *xsdDate `xml:"urn:com.workday/bsvc Benefits_Service_Date,omitempty"`
		CompanyServiceDate        *xsdDate `xml:"urn:com.workday/bsvc Company_Service_Date,omitempty"`
		TimeOffServiceDate        *xsdDate `xml:"urn:com.workday/bsvc Time_Off_Service_Date,omitempty"`
		VestingDate               *xsdDate `xml:"urn:com.workday/bsvc Vesting_Date,omitempty"`
		DateEnteredWorkforce      *xsdDate `xml:"urn:com.workday/bsvc Date_Entered_Workforce,omitempty"`
		TerminationDate           *xsdDate `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
		PayThroughDate            *xsdDate `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
		TerminationLastDayofWork  *xsdDate `xml:"urn:com.workday/bsvc Termination_Last_Day_of_Work,omitempty"`
		ResignationDate           *xsdDate `xml:"urn:com.workday/bsvc Resignation_Date,omitempty"`
		LastDateforWhichPaid      *xsdDate `xml:"urn:com.workday/bsvc Last_Date_for_Which_Paid,omitempty"`
		ExpectedDateofReturn      *xsdDate `xml:"urn:com.workday/bsvc Expected_Date_of_Return,omitempty"`
		ProbationStartDate        *xsdDate `xml:"urn:com.workday/bsvc Probation_Start_Date,omitempty"`
		ProbationEndDate          *xsdDate `xml:"urn:com.workday/bsvc Probation_End_Date,omitempty"`
		AcademicTenureDate        *xsdDate `xml:"urn:com.workday/bsvc Academic_Tenure_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ActiveStatusDate = (*xsdDate)(layout.T.ActiveStatusDate)
	layout.HireDate = (*xsdDate)(layout.T.HireDate)
	layout.OriginalHireDate = (*xsdDate)(layout.T.OriginalHireDate)
	layout.EndEmploymentDate = (*xsdDate)(layout.T.EndEmploymentDate)
	layout.ContinuousServiceDate = (*xsdDate)(layout.T.ContinuousServiceDate)
	layout.FirstDayofWork = (*xsdDate)(layout.T.FirstDayofWork)
	layout.ExpectedRetirementDate = (*xsdDate)(layout.T.ExpectedRetirementDate)
	layout.RetirementEligibilityDate = (*xsdDate)(layout.T.RetirementEligibilityDate)
	layout.RetirementDate = (*xsdDate)(layout.T.RetirementDate)
	layout.SeniorityDate = (*xsdDate)(layout.T.SeniorityDate)
	layout.SeveranceDate = (*xsdDate)(layout.T.SeveranceDate)
	layout.BenefitsServiceDate = (*xsdDate)(layout.T.BenefitsServiceDate)
	layout.CompanyServiceDate = (*xsdDate)(layout.T.CompanyServiceDate)
	layout.TimeOffServiceDate = (*xsdDate)(layout.T.TimeOffServiceDate)
	layout.VestingDate = (*xsdDate)(layout.T.VestingDate)
	layout.DateEnteredWorkforce = (*xsdDate)(layout.T.DateEnteredWorkforce)
	layout.TerminationDate = (*xsdDate)(layout.T.TerminationDate)
	layout.PayThroughDate = (*xsdDate)(layout.T.PayThroughDate)
	layout.TerminationLastDayofWork = (*xsdDate)(layout.T.TerminationLastDayofWork)
	layout.ResignationDate = (*xsdDate)(layout.T.ResignationDate)
	layout.LastDateforWhichPaid = (*xsdDate)(layout.T.LastDateforWhichPaid)
	layout.ExpectedDateofReturn = (*xsdDate)(layout.T.ExpectedDateofReturn)
	layout.ProbationStartDate = (*xsdDate)(layout.T.ProbationStartDate)
	layout.ProbationEndDate = (*xsdDate)(layout.T.ProbationEndDate)
	layout.AcademicTenureDate = (*xsdDate)(layout.T.AcademicTenureDate)
	return e.EncodeElement(layout, start)
}
func (t *WorkerStatusDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkerStatusDetailDataType
	var overlay struct {
		*T
		ActiveStatusDate          *xsdDate `xml:"urn:com.workday/bsvc Active_Status_Date,omitempty"`
		HireDate                  *xsdDate `xml:"urn:com.workday/bsvc Hire_Date,omitempty"`
		OriginalHireDate          *xsdDate `xml:"urn:com.workday/bsvc Original_Hire_Date,omitempty"`
		EndEmploymentDate         *xsdDate `xml:"urn:com.workday/bsvc End_Employment_Date,omitempty"`
		ContinuousServiceDate     *xsdDate `xml:"urn:com.workday/bsvc Continuous_Service_Date,omitempty"`
		FirstDayofWork            *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Work,omitempty"`
		ExpectedRetirementDate    *xsdDate `xml:"urn:com.workday/bsvc Expected_Retirement_Date,omitempty"`
		RetirementEligibilityDate *xsdDate `xml:"urn:com.workday/bsvc Retirement_Eligibility_Date,omitempty"`
		RetirementDate            *xsdDate `xml:"urn:com.workday/bsvc Retirement_Date,omitempty"`
		SeniorityDate             *xsdDate `xml:"urn:com.workday/bsvc Seniority_Date,omitempty"`
		SeveranceDate             *xsdDate `xml:"urn:com.workday/bsvc Severance_Date,omitempty"`
		BenefitsServiceDate       *xsdDate `xml:"urn:com.workday/bsvc Benefits_Service_Date,omitempty"`
		CompanyServiceDate        *xsdDate `xml:"urn:com.workday/bsvc Company_Service_Date,omitempty"`
		TimeOffServiceDate        *xsdDate `xml:"urn:com.workday/bsvc Time_Off_Service_Date,omitempty"`
		VestingDate               *xsdDate `xml:"urn:com.workday/bsvc Vesting_Date,omitempty"`
		DateEnteredWorkforce      *xsdDate `xml:"urn:com.workday/bsvc Date_Entered_Workforce,omitempty"`
		TerminationDate           *xsdDate `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
		PayThroughDate            *xsdDate `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
		TerminationLastDayofWork  *xsdDate `xml:"urn:com.workday/bsvc Termination_Last_Day_of_Work,omitempty"`
		ResignationDate           *xsdDate `xml:"urn:com.workday/bsvc Resignation_Date,omitempty"`
		LastDateforWhichPaid      *xsdDate `xml:"urn:com.workday/bsvc Last_Date_for_Which_Paid,omitempty"`
		ExpectedDateofReturn      *xsdDate `xml:"urn:com.workday/bsvc Expected_Date_of_Return,omitempty"`
		ProbationStartDate        *xsdDate `xml:"urn:com.workday/bsvc Probation_Start_Date,omitempty"`
		ProbationEndDate          *xsdDate `xml:"urn:com.workday/bsvc Probation_End_Date,omitempty"`
		AcademicTenureDate        *xsdDate `xml:"urn:com.workday/bsvc Academic_Tenure_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ActiveStatusDate = (*xsdDate)(overlay.T.ActiveStatusDate)
	overlay.HireDate = (*xsdDate)(overlay.T.HireDate)
	overlay.OriginalHireDate = (*xsdDate)(overlay.T.OriginalHireDate)
	overlay.EndEmploymentDate = (*xsdDate)(overlay.T.EndEmploymentDate)
	overlay.ContinuousServiceDate = (*xsdDate)(overlay.T.ContinuousServiceDate)
	overlay.FirstDayofWork = (*xsdDate)(overlay.T.FirstDayofWork)
	overlay.ExpectedRetirementDate = (*xsdDate)(overlay.T.ExpectedRetirementDate)
	overlay.RetirementEligibilityDate = (*xsdDate)(overlay.T.RetirementEligibilityDate)
	overlay.RetirementDate = (*xsdDate)(overlay.T.RetirementDate)
	overlay.SeniorityDate = (*xsdDate)(overlay.T.SeniorityDate)
	overlay.SeveranceDate = (*xsdDate)(overlay.T.SeveranceDate)
	overlay.BenefitsServiceDate = (*xsdDate)(overlay.T.BenefitsServiceDate)
	overlay.CompanyServiceDate = (*xsdDate)(overlay.T.CompanyServiceDate)
	overlay.TimeOffServiceDate = (*xsdDate)(overlay.T.TimeOffServiceDate)
	overlay.VestingDate = (*xsdDate)(overlay.T.VestingDate)
	overlay.DateEnteredWorkforce = (*xsdDate)(overlay.T.DateEnteredWorkforce)
	overlay.TerminationDate = (*xsdDate)(overlay.T.TerminationDate)
	overlay.PayThroughDate = (*xsdDate)(overlay.T.PayThroughDate)
	overlay.TerminationLastDayofWork = (*xsdDate)(overlay.T.TerminationLastDayofWork)
	overlay.ResignationDate = (*xsdDate)(overlay.T.ResignationDate)
	overlay.LastDateforWhichPaid = (*xsdDate)(overlay.T.LastDateforWhichPaid)
	overlay.ExpectedDateofReturn = (*xsdDate)(overlay.T.ExpectedDateofReturn)
	overlay.ProbationStartDate = (*xsdDate)(overlay.T.ProbationStartDate)
	overlay.ProbationEndDate = (*xsdDate)(overlay.T.ProbationEndDate)
	overlay.AcademicTenureDate = (*xsdDate)(overlay.T.AcademicTenureDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains the worker's supervisory management chain.
type WorkerSupervisoryManagementChainDataType struct {
	ManagementChainData []ManagementChainDataType `xml:"urn:com.workday/bsvc Management_Chain_Data,omitempty"`
}

// Encapsulating element containing all Worker Profile data.
type WorkerType struct {
	WorkerReference              *WorkerObjectType              `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	WorkerDescriptor             string                         `xml:"urn:com.workday/bsvc Worker_Descriptor,omitempty"`
	UniversalIdentifierReference *UniversalIdentifierObjectType `xml:"urn:com.workday/bsvc Universal_Identifier_Reference,omitempty"`
	WorkerData                   *WorkerDataType                `xml:"urn:com.workday/bsvc Worker_Data,omitempty"`
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

// National Id request criteria.
type WorkerbyNationalIDRequestCriteriaType struct {
	IdentifierID            string                    `xml:"urn:com.workday/bsvc Identifier_ID"`
	NationalIDTypeReference *NationalIDTypeObjectType `xml:"urn:com.workday/bsvc National_ID_Type_Reference"`
	CountryReference        *CountryObjectType        `xml:"urn:com.workday/bsvc Country_Reference"`
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

// Contains each Worker based on the Request References or Request Criteria.  The data returned is the current information as of the dates in the response filter, and does not include all historical information about the worker.
type WorkersResponseDataType struct {
	Worker []WorkerType `xml:"urn:com.workday/bsvc Worker,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WorkingTimeUnitObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkingTimeUnitObjectType struct {
	ID         []WorkingTimeUnitObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
