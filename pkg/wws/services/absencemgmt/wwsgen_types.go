package absencemgmt

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// Contains a unique identifier for an instance of an object.
type AbsenceComponentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AbsenceComponentObjectType struct {
	ID         []AbsenceComponentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Absence Input Data
type AbsenceInputDataType struct {
	AbsenceInputID            string                      `xml:"urn:com.workday/bsvc Absence_Input_ID,omitempty"`
	BatchID                   string                      `xml:"urn:com.workday/bsvc Batch_ID,omitempty"`
	WorkerReference           *WorkerObjectType           `xml:"urn:com.workday/bsvc Worker_Reference"`
	PositionReference         *PositionElementObjectType  `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	AbsenceComponentReference *AbsenceComponentObjectType `xml:"urn:com.workday/bsvc Absence_Component_Reference"`
	StartDate                 time.Time                   `xml:"urn:com.workday/bsvc Start_Date"`
	EndDate                   *time.Time                  `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	ReferenceDate             time.Time                   `xml:"urn:com.workday/bsvc Reference_Date"`
	Hours                     float64                     `xml:"urn:com.workday/bsvc Hours,omitempty"`
	Adjustment                *bool                       `xml:"urn:com.workday/bsvc Adjustment,omitempty"`
	Comment                   string                      `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

func (t *AbsenceInputDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AbsenceInputDataType
	var layout struct {
		*T
		StartDate     *xsdDate `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate       *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
		ReferenceDate *xsdDate `xml:"urn:com.workday/bsvc Reference_Date"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(&layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	layout.ReferenceDate = (*xsdDate)(&layout.T.ReferenceDate)
	return e.EncodeElement(layout, start)
}
func (t *AbsenceInputDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AbsenceInputDataType
	var overlay struct {
		*T
		StartDate     *xsdDate `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate       *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
		ReferenceDate *xsdDate `xml:"urn:com.workday/bsvc Reference_Date"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(&overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	overlay.ReferenceDate = (*xsdDate)(&overlay.T.ReferenceDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type AbsenceInputObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AbsenceInputObjectType struct {
	ID         []AbsenceInputObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Absence Input Request Criteria
type AbsenceInputRequestCriteriaType struct {
	EmployeeReference *WorkerObjectType           `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	BatchReference    *ApplicationBatchObjectType `xml:"urn:com.workday/bsvc Batch_Reference,omitempty"`
}

// Absence Input Request References
type AbsenceInputRequestReferencesType struct {
	AbsenceInputReference []AbsenceInputObjectType `xml:"urn:com.workday/bsvc Absence_Input_Reference"`
}

// Absence Input Response Data
type AbsenceInputResponseDataType struct {
	AbsenceInput []AbsenceInputType `xml:"urn:com.workday/bsvc Absence_Input,omitempty"`
}

// Absence Input Response Group
type AbsenceInputResponseGroupType struct {
	IncludeReference        *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeAbsenceInputData *bool `xml:"urn:com.workday/bsvc Include_Absence_Input_Data,omitempty"`
}

// Absence Input
type AbsenceInputType struct {
	AbsenceInputReference *AbsenceInputObjectType `xml:"urn:com.workday/bsvc Absence_Input_Reference,omitempty"`
	AbsenceInputData      *AbsenceInputDataType   `xml:"urn:com.workday/bsvc Absence_Input_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AbsencePlanAllObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AbsencePlanAllObjectType struct {
	ID         []AbsencePlanAllObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Absence Plan Reference
type AbsencePlanReferenceType struct {
	TimeOffPlanName string `xml:"urn:com.workday/bsvc Time_Off_Plan_Name"`
}

// Contains a unique identifier for an instance of an object.
type AbsenceTableObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AbsenceTableObjectType struct {
	ID         []AbsenceTableObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains adjustment data for the Adjust Time Off business process
type AdjustTimeOffDataType struct {
	RunTimeOffValidations  *bool                        `xml:"urn:com.workday/bsvc Run_Time_Off_Validations,omitempty"`
	WorkerReference        *WorkerObjectType            `xml:"urn:com.workday/bsvc Worker_Reference"`
	AdjustTimeOffEntryData []AdjustTimeOffEntryDataType `xml:"urn:com.workday/bsvc Adjust_Time_Off_Entry_Data"`
}

// Contains data for adjusting time off
type AdjustTimeOffEntryDataType struct {
	TimeOffEntryID        string                  `xml:"urn:com.workday/bsvc Time_Off_Entry_ID,omitempty"`
	TimeOffEntryReference *TimeOffEntryObjectType `xml:"urn:com.workday/bsvc Time_Off_Entry_Reference"`
	AdjustmenttoRequested float64                 `xml:"urn:com.workday/bsvc Adjustment_to_Requested"`
	Comment               string                  `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

// The purpose of this spreadsheet is to provide a facility to upload data from Microsoft Excel into your Workday system. By performing this upload process, the customer acknowledges that they have already performed the required business process tasks and approvals to confirm the worker's new data values. This spreadsheet contains a worksheet for Correct Time Off business process. Fill out this worksheet with the data to match your business processes. It is understood that different supervisory organizations may have a different business processes and therefore not every worker will necessarily go through all the optional related processes. When doing the Correct Time Off business process via the user interface, you enter the 'correction' to requested units (example - if original request was for 8 hours and it should have been 6, you would enter 6 as entering 'corrected' units). However, when using this spreadsheet you need to enter the 'adjustment' to the requested units (example - you would enter -2 as 'adjustment' to requested units).
type AdjustTimeOffRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AdjustTimeOffData         *AdjustTimeOffDataType         `xml:"urn:com.workday/bsvc Adjust_Time_Off_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type ApplicationBatchObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ApplicationBatchObjectType struct {
	ID         []ApplicationBatchObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type AssignableRoleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AssignableRoleObjectType struct {
	ID         []AssignableRoleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type CarryoverExpirationLimitOverrideObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CarryoverExpirationLimitOverrideObjectType struct {
	ID         []CarryoverExpirationLimitOverrideObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Override Balance Units Data
type CarryoverOverrideBalanceUnitsDataType struct {
	CarryoverExpirationDate       *time.Time `xml:"urn:com.workday/bsvc Carryover_Expiration_Date,omitempty"`
	CarryoverOverrideBalanceUnits float64    `xml:"urn:com.workday/bsvc Carryover_Override_Balance_Units,omitempty"`
}

func (t *CarryoverOverrideBalanceUnitsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CarryoverOverrideBalanceUnitsDataType
	var layout struct {
		*T
		CarryoverExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Carryover_Expiration_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CarryoverExpirationDate = (*xsdDate)(layout.T.CarryoverExpirationDate)
	return e.EncodeElement(layout, start)
}
func (t *CarryoverOverrideBalanceUnitsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CarryoverOverrideBalanceUnitsDataType
	var overlay struct {
		*T
		CarryoverExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Carryover_Expiration_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CarryoverExpirationDate = (*xsdDate)(overlay.T.CarryoverExpirationDate)
	return d.DecodeElement(&overlay, &start)
}

// Carryover Override
type CarryoverOverrideDataType struct {
	ID                      string                     `xml:"urn:com.workday/bsvc ID,omitempty"`
	XMLNAMEWorkerReference  *WorkerObjectType          `xml:"urn:com.workday/bsvc XMLNAME__Worker__Reference"`
	AbsencePlanReference    *AbsencePlanAllObjectType  `xml:"urn:com.workday/bsvc Absence_Plan_Reference"`
	CarryoverDate           time.Time                  `xml:"urn:com.workday/bsvc Carryover_Date"`
	PositionReference       *PositionElementObjectType `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	CarryoverOverrideAmount float64                    `xml:"urn:com.workday/bsvc Carryover_Override_Amount,omitempty"`
	CarryoverExpirationDate time.Time                  `xml:"urn:com.workday/bsvc Carryover_Expiration_Date"`
	Comments                string                     `xml:"urn:com.workday/bsvc Comments,omitempty"`
}

func (t *CarryoverOverrideDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CarryoverOverrideDataType
	var layout struct {
		*T
		CarryoverDate           *xsdDate `xml:"urn:com.workday/bsvc Carryover_Date"`
		CarryoverExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Carryover_Expiration_Date"`
	}
	layout.T = (*T)(t)
	layout.CarryoverDate = (*xsdDate)(&layout.T.CarryoverDate)
	layout.CarryoverExpirationDate = (*xsdDate)(&layout.T.CarryoverExpirationDate)
	return e.EncodeElement(layout, start)
}
func (t *CarryoverOverrideDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CarryoverOverrideDataType
	var overlay struct {
		*T
		CarryoverDate           *xsdDate `xml:"urn:com.workday/bsvc Carryover_Date"`
		CarryoverExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Carryover_Expiration_Date"`
	}
	overlay.T = (*T)(t)
	overlay.CarryoverDate = (*xsdDate)(&overlay.T.CarryoverDate)
	overlay.CarryoverExpirationDate = (*xsdDate)(&overlay.T.CarryoverExpirationDate)
	return d.DecodeElement(&overlay, &start)
}

// Carryover Override Request Criteria
type CarryoverOverrideRequestCriteriaType struct {
	XMLNAMEWorkerReference *WorkerObjectType `xml:"urn:com.workday/bsvc XMLNAME__Worker__Reference,omitempty"`
}

// Carryover Override Request References
type CarryoverOverrideRequestReferencesType struct {
	CarryoverOverrideReference []CarryoverExpirationLimitOverrideObjectType `xml:"urn:com.workday/bsvc Carryover_Override_Reference"`
}

// Get Carryover Override Response
type CarryoverOverrideResponseDataType struct {
	CarryoverOverride []CarryoverOverrideType `xml:"urn:com.workday/bsvc Carryover_Override,omitempty"`
}

// Carryover Override Response Group
type CarryoverOverrideResponseGroupType struct {
	IncludeReference             *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeCarryoverOverrideData *bool `xml:"urn:com.workday/bsvc Include_Carryover_Override_Data,omitempty"`
}

// Carryover Override Response Data
type CarryoverOverrideType struct {
	CarryoverOverrideReference *CarryoverExpirationLimitOverrideObjectType `xml:"urn:com.workday/bsvc Carryover_Override_Reference,omitempty"`
	CarryoverOverrideData      *CarryoverOverrideDataType                  `xml:"urn:com.workday/bsvc Carryover_Override_Data,omitempty"`
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
type CurrencyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CurrencyObjectType struct {
	ID         []CurrencyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type EarningAllObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EarningAllObjectType struct {
	ID         []EarningAllObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type EditServiceDatesSubBusinessProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	EditServiceDatesEventData    *EditServiceDatesEventDataType    `xml:"urn:com.workday/bsvc Edit_Service_Dates_Event_Data,omitempty"`
}

// Contains Request Time Off data
type EnterTimeOffDataType struct {
	TurnOffRunTimeCalculationswithTimeOffApproval *bool                       `xml:"urn:com.workday/bsvc Turn_Off_Run_Time_Calculations_with_Time_Off_Approval,omitempty"`
	RunTimeOffValidations                         *bool                       `xml:"urn:com.workday/bsvc Run_Time_Off_Validations,omitempty"`
	WorkerReference                               *WorkerObjectType           `xml:"urn:com.workday/bsvc Worker_Reference"`
	EnterTimeOffEntryData                         []EnterTimeOffEntryDataType `xml:"urn:com.workday/bsvc Enter_Time_Off_Entry_Data"`
}

// Data for Employee Paid TIme Off
type EnterTimeOffEntryDataType struct {
	TimeOffEntryID         string                     `xml:"urn:com.workday/bsvc Time_Off_Entry_ID,omitempty"`
	Date                   time.Time                  `xml:"urn:com.workday/bsvc Date"`
	Requested              float64                    `xml:"urn:com.workday/bsvc Requested"`
	TimeOffTypeReference   *TimeOffTypeObjectType     `xml:"urn:com.workday/bsvc Time_Off_Type_Reference,omitempty"`
	TimeOffReference       *TimeOffObjectType         `xml:"urn:com.workday/bsvc Time_Off_Reference,omitempty"`
	AbsenceTableReference  *AbsenceTableObjectType    `xml:"urn:com.workday/bsvc Absence_Table_Reference,omitempty"`
	PositionReference      *PositionElementObjectType `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	TimeOffReasonReference *TimeOffReasonObjectType   `xml:"urn:com.workday/bsvc Time_Off_Reason_Reference,omitempty"`
	Comment                string                     `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

func (t *EnterTimeOffEntryDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EnterTimeOffEntryDataType
	var layout struct {
		*T
		Date *xsdDate `xml:"urn:com.workday/bsvc Date"`
	}
	layout.T = (*T)(t)
	layout.Date = (*xsdDate)(&layout.T.Date)
	return e.EncodeElement(layout, start)
}
func (t *EnterTimeOffEntryDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EnterTimeOffEntryDataType
	var overlay struct {
		*T
		Date *xsdDate `xml:"urn:com.workday/bsvc Date"`
	}
	overlay.T = (*T)(t)
	overlay.Date = (*xsdDate)(&overlay.T.Date)
	return d.DecodeElement(&overlay, &start)
}

// The purpose of this spreadsheet is to provide a facility to upload data from Microsoft Excel into your Workday system. By performing this upload process, the customer acknowledges that they have already performed the required business process tasks and approvals to confirm the worker's new data values. This spreadsheet contains worksheets for "Enter Time Off"; both the main process and available related processes. Fill out these worksheets with the data to match your business processes. It is understood that different supervisory organizations may have a different business processes and therefore not every worker will necessarily go through all the optional related processes.
type EnterTimeOffRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EnterTimeOffData          *EnterTimeOffDataType          `xml:"urn:com.workday/bsvc Enter_Time_Off_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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
type FrequencyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type FrequencyObjectType struct {
	ID         []FrequencyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Get Absence Inputs Request
type GetAbsenceInputsRequestType struct {
	RequestReferences *AbsenceInputRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *AbsenceInputRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *AbsenceInputResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Absence Inputs Response
type GetAbsenceInputsResponseType struct {
	RequestReferences *AbsenceInputRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *AbsenceInputRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *AbsenceInputResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType               `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *AbsenceInputResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Carryover Override Request
type GetCarryoverOverrideRequestType struct {
	RequestReferences *CarryoverOverrideRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CarryoverOverrideRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CarryoverOverrideResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Carryover Override Response
type GetCarryoverOverrideResponseType struct {
	RequestReferences *CarryoverOverrideRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CarryoverOverrideRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroups    *CarryoverOverrideResponseGroupType     `xml:"urn:com.workday/bsvc Response_Groups,omitempty"`
	ResponseResults   *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *CarryoverOverrideResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Time Off Plan Override Balances Request
type GetOverrideBalancesRequestType struct {
	RequestReferences *OverrideBalanceRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *OverrideBalanceRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *OverrideBalanceResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Time Off Plan Override Balances Response
type GetOverrideBalancesResponseType struct {
	RequestReferences *OverrideBalanceRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *OverrideBalanceRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroups    *OverrideBalanceResponseGroupType     `xml:"urn:com.workday/bsvc Response_Groups,omitempty"`
	ResponseResults   *ResponseResultsType                  `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *OverrideBalanceResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Time Off Plan Balances Request
type GetTimeOffPlanBalancesRequestType struct {
	RequestCriteria *TimeOffPlanBalanceRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *TimeOffPlanBalanceResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version         string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Time Off Plan Balances Response
type GetTimeOffPlanBalancesResponseType struct {
	RequestCriteria *TimeOffPlanBalanceRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup   *TimeOffPlanBalanceResponseGroupType   `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults []ResponseResultsType                  `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    []TimeOffPlanBalanceResponseDataType   `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element which contains collection of Enter Time Off Data elements needed for a Time Off Request.
type ImportTimeOffRequestEventBatchRequestType struct {
	EnterTimeOffData []EnterTimeOffDataType `xml:"urn:com.workday/bsvc Enter_Time_Off_Data,omitempty"`
	Version          string                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Element containing the Request Leave of Absence Data for an employee.
type LeaveRequestDataType struct {
	RequestLeaveofAbsenceID         string                        `xml:"urn:com.workday/bsvc Request_Leave_of_Absence_ID,omitempty"`
	Correction                      *bool                         `xml:"urn:com.workday/bsvc Correction,omitempty"`
	LeaveofAbsenceTypeReference     *LeaveofAbsenceTypeObjectType `xml:"urn:com.workday/bsvc Leave_of_Absence_Type_Reference"`
	LeaveReasonReference            *LeaveTypeReasonObjectType    `xml:"urn:com.workday/bsvc Leave_Reason_Reference,omitempty"`
	PositionReference               *PositionElementObjectType    `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	FirstDayofLeave                 time.Time                     `xml:"urn:com.workday/bsvc First_Day_of_Leave"`
	LastDayofWork                   *time.Time                    `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
	EstimatedLastDayofLeave         time.Time                     `xml:"urn:com.workday/bsvc Estimated_Last_Day_of_Leave"`
	LinksBacktoPriorEventReference  *LeaveRequestEventObjectType  `xml:"urn:com.workday/bsvc Links_Back_to_Prior_Event_Reference,omitempty"`
	DependentReference              *DependentObjectType          `xml:"urn:com.workday/bsvc Dependent_Reference,omitempty"`
	LastDateforWhichPaid            *time.Time                    `xml:"urn:com.workday/bsvc Last_Date_for_Which_Paid,omitempty"`
	ExpectedDueDate                 *time.Time                    `xml:"urn:com.workday/bsvc Expected_Due_Date,omitempty"`
	ChildsBirthDate                 *time.Time                    `xml:"urn:com.workday/bsvc Child_s_Birth_Date,omitempty"`
	StillbirthBabyDeceased          *bool                         `xml:"urn:com.workday/bsvc Stillbirth_Baby_Deceased,omitempty"`
	DateBabyArrivedHomeFromHospital *time.Time                    `xml:"urn:com.workday/bsvc Date_Baby_Arrived_Home_From_Hospital,omitempty"`
	AdoptionPlacementDate           *time.Time                    `xml:"urn:com.workday/bsvc Adoption_Placement_Date,omitempty"`
	AdoptionNotificationDate        *time.Time                    `xml:"urn:com.workday/bsvc Adoption_Notification_Date,omitempty"`
	DateChildEnteredCountry         *time.Time                    `xml:"urn:com.workday/bsvc Date_Child_Entered_Country,omitempty"`
	MultipleChildIndicator          *bool                         `xml:"urn:com.workday/bsvc Multiple_Child_Indicator,omitempty"`
	NumberofBabiesAdoptedChildren   float64                       `xml:"urn:com.workday/bsvc Number_of_Babies_Adopted_Children,omitempty"`
	NumberofPreviousBirths          float64                       `xml:"urn:com.workday/bsvc Number_of_Previous_Births,omitempty"`
	NumberofPreviousMaternityLeaves float64                       `xml:"urn:com.workday/bsvc Number_of_Previous_Maternity_Leaves,omitempty"`
	NumberofChildDependents         float64                       `xml:"urn:com.workday/bsvc Number_of_Child_Dependents,omitempty"`
	SingleParentIndicator           *bool                         `xml:"urn:com.workday/bsvc Single_Parent_Indicator,omitempty"`
	AgeofDependent                  float64                       `xml:"urn:com.workday/bsvc Age_of_Dependent,omitempty"`
	WorkRelated                     *bool                         `xml:"urn:com.workday/bsvc Work_Related,omitempty"`
	StopPaymentDate                 *time.Time                    `xml:"urn:com.workday/bsvc Stop_Payment_Date,omitempty"`
	SocialSecurityDisabilityCode    string                        `xml:"urn:com.workday/bsvc Social_Security_Disability_Code,omitempty"`
	LocationDuringLeave             string                        `xml:"urn:com.workday/bsvc Location_During_Leave,omitempty"`
	CaesareanSectionBirth           *bool                         `xml:"urn:com.workday/bsvc Caesarean_Section_Birth,omitempty"`
	LeavePercentage                 float64                       `xml:"urn:com.workday/bsvc Leave_Percentage,omitempty"`
	WeekofConfinement               *time.Time                    `xml:"urn:com.workday/bsvc Week_of_Confinement,omitempty"`
	LeaveEntitlementOverride        float64                       `xml:"urn:com.workday/bsvc Leave_Entitlement_Override,omitempty"`
	DateofRecall                    *time.Time                    `xml:"urn:com.workday/bsvc Date_of_Recall,omitempty"`
}

func (t *LeaveRequestDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LeaveRequestDataType
	var layout struct {
		*T
		FirstDayofLeave                 *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Leave"`
		LastDayofWork                   *xsdDate `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
		EstimatedLastDayofLeave         *xsdDate `xml:"urn:com.workday/bsvc Estimated_Last_Day_of_Leave"`
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
	layout.FirstDayofLeave = (*xsdDate)(&layout.T.FirstDayofLeave)
	layout.LastDayofWork = (*xsdDate)(layout.T.LastDayofWork)
	layout.EstimatedLastDayofLeave = (*xsdDate)(&layout.T.EstimatedLastDayofLeave)
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
func (t *LeaveRequestDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LeaveRequestDataType
	var overlay struct {
		*T
		FirstDayofLeave                 *xsdDate `xml:"urn:com.workday/bsvc First_Day_of_Leave"`
		LastDayofWork                   *xsdDate `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
		EstimatedLastDayofLeave         *xsdDate `xml:"urn:com.workday/bsvc Estimated_Last_Day_of_Leave"`
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
	overlay.FirstDayofLeave = (*xsdDate)(&overlay.T.FirstDayofLeave)
	overlay.LastDayofWork = (*xsdDate)(overlay.T.LastDayofWork)
	overlay.EstimatedLastDayofLeave = (*xsdDate)(&overlay.T.EstimatedLastDayofLeave)
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

// Element containing the Request Return from Leave of Absence Data for an employee.
type LeavesReturningfromDataType struct {
	RequestLeaveofAbsenceReference *LeaveRequestEventObjectType `xml:"urn:com.workday/bsvc Request_Leave_of_Absence_Reference"`
	Correction                     *bool                        `xml:"urn:com.workday/bsvc Correction,omitempty"`
	ActualLastDayofLeave           time.Time                    `xml:"urn:com.workday/bsvc Actual_Last_Day_of_Leave"`
}

func (t *LeavesReturningfromDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T LeavesReturningfromDataType
	var layout struct {
		*T
		ActualLastDayofLeave *xsdDate `xml:"urn:com.workday/bsvc Actual_Last_Day_of_Leave"`
	}
	layout.T = (*T)(t)
	layout.ActualLastDayofLeave = (*xsdDate)(&layout.T.ActualLastDayofLeave)
	return e.EncodeElement(layout, start)
}
func (t *LeavesReturningfromDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T LeavesReturningfromDataType
	var overlay struct {
		*T
		ActualLastDayofLeave *xsdDate `xml:"urn:com.workday/bsvc Actual_Last_Day_of_Leave"`
	}
	overlay.T = (*T)(t)
	overlay.ActualLastDayofLeave = (*xsdDate)(&overlay.T.ActualLastDayofLeave)
	return d.DecodeElement(&overlay, &start)
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
type OrganizationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type OrganizationObjectType struct {
	ID         []OrganizationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for a Organization Role Assignment
type OrganizationRoleAssignmentDataType struct {
	RoleAssignerReference            *RoleAssignerObjectType   `xml:"urn:com.workday/bsvc Role_Assigner_Reference,omitempty"`
	OrganizationRoleReference        *AssignableRoleObjectType `xml:"urn:com.workday/bsvc Organization_Role_Reference"`
	RoleAssigneeReference            []RoleeObjectType         `xml:"urn:com.workday/bsvc Role_Assignee_Reference,omitempty"`
	SingleAssignmentManagerReference *RoleeObjectType          `xml:"urn:com.workday/bsvc Single_Assignment_Manager_Reference,omitempty"`
}

// Time Off Plan Override Balance Data
type OverrideBalanceDataType struct {
	ID                       string                                  `xml:"urn:com.workday/bsvc ID,omitempty"`
	BatchID                  string                                  `xml:"urn:com.workday/bsvc Batch_ID,omitempty"`
	WorkerReference          *WorkerObjectType                       `xml:"urn:com.workday/bsvc Worker_Reference"`
	PositionReference        *PositionElementObjectType              `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	TimeOffPlanReference     *AbsencePlanReferenceType               `xml:"urn:com.workday/bsvc Time_Off_Plan_Reference"`
	OverrideBalanceDate      time.Time                               `xml:"urn:com.workday/bsvc Override_Balance_Date"`
	OverrideBalanceUnits     float64                                 `xml:"urn:com.workday/bsvc Override_Balance_Units,omitempty"`
	OverrideBalanceUnitsData []CarryoverOverrideBalanceUnitsDataType `xml:"urn:com.workday/bsvc Override_Balance_Units_Data,omitempty"`
	Comments                 string                                  `xml:"urn:com.workday/bsvc Comments,omitempty"`
}

func (t *OverrideBalanceDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T OverrideBalanceDataType
	var layout struct {
		*T
		OverrideBalanceDate *xsdDate `xml:"urn:com.workday/bsvc Override_Balance_Date"`
	}
	layout.T = (*T)(t)
	layout.OverrideBalanceDate = (*xsdDate)(&layout.T.OverrideBalanceDate)
	return e.EncodeElement(layout, start)
}
func (t *OverrideBalanceDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T OverrideBalanceDataType
	var overlay struct {
		*T
		OverrideBalanceDate *xsdDate `xml:"urn:com.workday/bsvc Override_Balance_Date"`
	}
	overlay.T = (*T)(t)
	overlay.OverrideBalanceDate = (*xsdDate)(&overlay.T.OverrideBalanceDate)
	return d.DecodeElement(&overlay, &start)
}

// Time Off Plan Override Balance Request Criteria
type OverrideBalanceRequestCriteriaType struct {
	EmployeeReference *WorkerObjectType           `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	BatchReference    *ApplicationBatchObjectType `xml:"urn:com.workday/bsvc Batch_Reference,omitempty"`
}

// Time Off Plan Override Balance Request References
type OverrideBalanceRequestReferencesType struct {
	OverrideBalanceReference []TimeOffPlanBalanceObjectType `xml:"urn:com.workday/bsvc Override_Balance_Reference"`
}

// Time Off Plan Override Balance Response Data
type OverrideBalanceResponseDataType struct {
	OverrideBalance []OverrideBalanceType `xml:"urn:com.workday/bsvc Override_Balance,omitempty"`
}

// Time Off Plan Override Balance Response Group
type OverrideBalanceResponseGroupType struct {
	IncludeReference           *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeOverrideBalanceData *bool `xml:"urn:com.workday/bsvc Include_Override_Balance_Data,omitempty"`
}

// Time Off Plan Override Balance
type OverrideBalanceType struct {
	OverrideBalanceReference *TimeOffPlanBalanceObjectType `xml:"urn:com.workday/bsvc Override_Balance_Reference,omitempty"`
	OverrideBalanceData      *OverrideBalanceDataType      `xml:"urn:com.workday/bsvc Override_Balance_Data,omitempty"`
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

type ProcessingFaultType struct {
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
}

// Wrapper element for the Propose Compensation for Leave of Absence sub business process.  If any errors are found during processing, the Auto Complete boolean will be set to False and manual processing will occur for this business process.
type ProposeCompensationForLeaveofAbsenceSubBusinessProcessType struct {
	PositionElementforCompensationTransactionReference *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Element_for_Compensation_Transaction_Reference,omitempty"`
	ReasonReference                                    *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference"`
	ProposeCompensationforLeaveofAbsenceData           *CompensationProposedForPositionDataType  `xml:"urn:com.workday/bsvc Propose_Compensation_for_Leave_of_Absence_Data,omitempty"`
}

// Wrapper element for the Propose Compensation for Leave of Absence sub business process
type ProposeCompensationforLeaveofAbsenceSubProcessType struct {
	BusinessSubProcessParameters             *BusinessSubProcessParametersType                            `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
	ProposeCompensationforLeaveofAbsenceData []ProposeCompensationForLeaveofAbsenceSubBusinessProcessType `xml:"urn:com.workday/bsvc Propose_Compensation_for_Leave_of_Absence_Data,omitempty"`
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

// Absence Input Request
type PutAbsenceInputRequestType struct {
	AbsenceInputReference *AbsenceInputObjectType `xml:"urn:com.workday/bsvc Absence_Input_Reference,omitempty"`
	AbsenceInputData      *AbsenceInputDataType   `xml:"urn:com.workday/bsvc Absence_Input_Data"`
	Version               string                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Absence Input Response
type PutAbsenceInputResponseType struct {
	AbsenceInputReference *AbsenceInputObjectType `xml:"urn:com.workday/bsvc Absence_Input_Reference,omitempty"`
	Version               string                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Carryover Override Request
type PutCarryoverOverrideRequestType struct {
	CarryoverExpirationLimitOverrideReference *CarryoverExpirationLimitOverrideObjectType `xml:"urn:com.workday/bsvc Carryover_Expiration_Limit_Override_Reference,omitempty"`
	CarryoverOverrideData                     *CarryoverOverrideDataType                  `xml:"urn:com.workday/bsvc Carryover_Override_Data"`
	AddOnly                                   bool                                        `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                                   string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Carryover Override Response
type PutCarryoverOverrideResponseType struct {
	CarryoverExpirationLimitOverrideReference *CarryoverExpirationLimitOverrideObjectType `xml:"urn:com.workday/bsvc Carryover_Expiration_Limit_Override_Reference,omitempty"`
	Version                                   string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Import Process Response element
type PutImportProcessResponseType struct {
	ImportProcessReference  *WebServiceBackgroundProcessRuntimeObjectType `xml:"urn:com.workday/bsvc Import_Process_Reference,omitempty"`
	HeaderInstanceReference *InstanceObjectType                           `xml:"urn:com.workday/bsvc Header_Instance_Reference,omitempty"`
	Version                 string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Override Balance Request
type PutOverrideBalanceRequestType struct {
	OverrideBalanceReference *TimeOffPlanBalanceObjectType `xml:"urn:com.workday/bsvc Override_Balance_Reference,omitempty"`
	OverrideBalanceData      *OverrideBalanceDataType      `xml:"urn:com.workday/bsvc Override_Balance_Data"`
	Version                  string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Override Balance Response
type PutOverrideBalanceResponseType struct {
	OverrideBalanceReference *TimeOffPlanBalanceObjectType `xml:"urn:com.workday/bsvc Override_Balance_Reference,omitempty"`
	Version                  string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains Request Leave of Absence data
type RequestLeaveofAbsenceDataType struct {
	WorkerReference                                *WorkerObjectType                                        `xml:"urn:com.workday/bsvc Worker_Reference"`
	LeaveRequestData                               *LeaveRequestDataType                                    `xml:"urn:com.workday/bsvc Leave_Request_Data"`
	ReviewCOBRAEligibilitySubProcess               *ReviewCOBRAEligibilitySubBusinessProcessType            `xml:"urn:com.workday/bsvc Review_COBRA_Eligibility_Sub_Process,omitempty"`
	ReviewPayrollInterfaceEventSubProcess          *ReviewPayrollInterfaceDataExtractBusinessSubProcessType `xml:"urn:com.workday/bsvc Review_Payroll_Interface_Event_Sub_Process,omitempty"`
	ProposeCompensationforLeaveofAbsenceSubProcess *ProposeCompensationforLeaveofAbsenceSubProcessType      `xml:"urn:com.workday/bsvc Propose_Compensation_for_Leave_of_Absence_Sub_Process,omitempty"`
	AssignOrganizationRolesSubProcess              *AssignOrganizationRolesSubBusinessProcessType           `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Sub_Process,omitempty"`
	AssignCostingAllocationSubProcess              *AssignCostingAllocationSubBusinessProcessType           `xml:"urn:com.workday/bsvc Assign_Costing_Allocation_Sub_Process,omitempty"`
}

// The purpose of this spreadsheet is to provide a facility to upload data from Microsoft Excel into your Workday system. By performing this upload process, the customer acknowledges that they have already performed the required business process tasks and approvals to confirm the worker's new data values. This spreadsheet contains worksheets for "Request Leave of Absence"; both the main process and available related processes. Fill out these worksheets with the data to match your business processes. It is understood that different supervisory organizations may have a different business processes and therefore not every worker will necessarily go through all the optional related processes.
type RequestLeaveofAbsenceRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	RequestLeaveofAbsenceData *RequestLeaveofAbsenceDataType `xml:"urn:com.workday/bsvc Request_Leave_of_Absence_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Event ID for the Leave Request Event along with the worker reference.
type RequestLeaveofAbsenceResponseType struct {
	EventReference *LeaveRequestEventObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains Request Return from Leave of Absence data
type RequestReturnfromLeaveofAbsenceDataType struct {
	WorkerReference                                *WorkerObjectType                                        `xml:"urn:com.workday/bsvc Worker_Reference"`
	FirstDayBackatWork                             *time.Time                                               `xml:"urn:com.workday/bsvc First_Day_Back_at_Work,omitempty"`
	LeavesReturningfromData                        []LeavesReturningfromDataType                            `xml:"urn:com.workday/bsvc Leaves_Returning_from_Data"`
	EditServiceDatesSubProcess                     *EditServiceDatesSubBusinessProcessType                  `xml:"urn:com.workday/bsvc Edit_Service_Dates_Sub_Process,omitempty"`
	ReviewPayrollInterfaceEventSubProcess          *ReviewPayrollInterfaceDataExtractBusinessSubProcessType `xml:"urn:com.workday/bsvc Review_Payroll_Interface_Event_Sub_Process,omitempty"`
	ProposeCompensationforLeaveofAbsenceSubProcess *ProposeCompensationforLeaveofAbsenceSubProcessType      `xml:"urn:com.workday/bsvc Propose_Compensation_for_Leave_of_Absence_Sub_Process,omitempty"`
	AssignOrganizationRolesSubProcess              *AssignOrganizationRolesSubBusinessProcessType           `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Sub_Process,omitempty"`
}

func (t *RequestReturnfromLeaveofAbsenceDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RequestReturnfromLeaveofAbsenceDataType
	var layout struct {
		*T
		FirstDayBackatWork *xsdDate `xml:"urn:com.workday/bsvc First_Day_Back_at_Work,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FirstDayBackatWork = (*xsdDate)(layout.T.FirstDayBackatWork)
	return e.EncodeElement(layout, start)
}
func (t *RequestReturnfromLeaveofAbsenceDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RequestReturnfromLeaveofAbsenceDataType
	var overlay struct {
		*T
		FirstDayBackatWork *xsdDate `xml:"urn:com.workday/bsvc First_Day_Back_at_Work,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FirstDayBackatWork = (*xsdDate)(overlay.T.FirstDayBackatWork)
	return d.DecodeElement(&overlay, &start)
}

// The purpose of this spreadsheet is to provide a facility to upload data from Microsoft Excel into your Workday system. By performing this upload process, the customer acknowledges that they have already performed the required business process tasks and approvals to confirm the worker's new data values. This spreadsheet contains worksheets for "Request Return from Leave of Absence"; both the main process and available related processes. Fill out these worksheets with the data to match your business processes. It is understood that different supervisory organizations may have a different business processes and therefore not every worker will necessarily go through all the optional related processes.
type RequestReturnfromLeaveofAbsenceRequestType struct {
	BusinessProcessParameters           *BusinessProcessParametersType           `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	RequestReturnfromLeaveofAbsenceData *RequestReturnfromLeaveofAbsenceDataType `xml:"urn:com.workday/bsvc Request_Return_from_Leave_of_Absence_Data"`
	Version                             string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Event ID for the Leave Return Event along with the worker reference.
type RequestReturnfromLeaveofAbsenceResponseType struct {
	EventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Event_Reference,omitempty"`
	Version        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type RoleAssignerObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RoleAssignerObjectType struct {
	ID         []RoleAssignerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating Element for all Organization Role Assignment data.
type RoleAssignmentType struct {
	RoleAssignmentReference *UniqueIdentifierObjectType         `xml:"urn:com.workday/bsvc Role_Assignment_Reference,omitempty"`
	RoleAssignmentData      *OrganizationRoleAssignmentDataType `xml:"urn:com.workday/bsvc Role_Assignment_Data"`
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
type StockPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StockPlanObjectType struct {
	ID         []StockPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type TimeOffEntryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimeOffEntryObjectType struct {
	ID         []TimeOffEntryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Time Off Entry References
type TimeOffEntryReferencesType struct {
	TimeOffEntryReference []TimeOffEntryObjectType `xml:"urn:com.workday/bsvc Time_Off_Entry_Reference,omitempty"`
}

// Root Response Element
type TimeOffEventResponseType struct {
	TimeOffEventReference  *UniqueIdentifierObjectType  `xml:"urn:com.workday/bsvc Time_Off_Event_Reference,omitempty"`
	TimeOffEntryReferences []TimeOffEntryReferencesType `xml:"urn:com.workday/bsvc Time_Off_Entry_References,omitempty"`
	Version                string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TimeOffObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimeOffObjectType struct {
	ID         []TimeOffObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Time Off Plan Balance Data
type TimeOffPlanBalanceDataType struct {
	TimeOffPlanBalanceRecord []TimeOffPlanBalanceRecordType `xml:"urn:com.workday/bsvc Time_Off_Plan_Balance_Record,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TimeOffPlanBalanceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimeOffPlanBalanceObjectType struct {
	ID         []TimeOffPlanBalanceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Time Off Plan Balance Position Record
type TimeOffPlanBalancePositionRecordType struct {
	PositionReference  *PositionObjectType `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	TimeOffPlanBalance float64             `xml:"urn:com.workday/bsvc Time_Off_Plan_Balance"`
}

// Time Off Plan Balance Record
type TimeOffPlanBalanceRecordType struct {
	TimeOffPlanReference             *AbsencePlanAllObjectType              `xml:"urn:com.workday/bsvc Time_Off_Plan_Reference"`
	UnitofTimeReference              *UnitofTimeObjectType                  `xml:"urn:com.workday/bsvc Unit_of_Time_Reference"`
	TimeOffPlanBalancePositionRecord []TimeOffPlanBalancePositionRecordType `xml:"urn:com.workday/bsvc Time_Off_Plan_Balance_Position_Record"`
}

// Time Off Plan Balance Request Criteria.  The intersection of criteria passed in this element is used to create a working set.
type TimeOffPlanBalanceRequestCriteriaType struct {
	EmployeeReference     *WorkerObjectType         `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	TimeOffPlanReference  *AbsencePlanAllObjectType `xml:"urn:com.workday/bsvc Time_Off_Plan_Reference,omitempty"`
	OrganizationReference []OrganizationObjectType  `xml:"urn:com.workday/bsvc Organization_Reference,omitempty"`
}

// Time Off Plan Balance Response Data
type TimeOffPlanBalanceResponseDataType struct {
	TimeOffPlanBalance []TimeOffPlanBalanceType `xml:"urn:com.workday/bsvc Time_Off_Plan_Balance,omitempty"`
}

// Time Off Plan Balance Response Group
type TimeOffPlanBalanceResponseGroupType struct {
	IncludeReference              *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeTimeOffPlanBalanceData *bool `xml:"urn:com.workday/bsvc Include_Time_Off_Plan_Balance_Data,omitempty"`
}

// Time Off Plan Balance
type TimeOffPlanBalanceType struct {
	EmployeeReference      *WorkerObjectType           `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	TimeOffPlanBalanceData *TimeOffPlanBalanceDataType `xml:"urn:com.workday/bsvc Time_Off_Plan_Balance_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TimeOffReasonObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimeOffReasonObjectType struct {
	ID         []TimeOffReasonObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TimeOffTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimeOffTypeObjectType struct {
	ID         []TimeOffTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type UnitofTimeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type UnitofTimeObjectType struct {
	ID         []UnitofTimeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
