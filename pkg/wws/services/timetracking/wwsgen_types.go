package timetracking

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// Data related to the Ad Hoc Schedule Event to be processed.
type AdHocScheduleEventDataType struct {
	AdHocScheduleID               string                           `xml:"urn:com.workday/bsvc Ad_Hoc_Schedule_ID"`
	EmployeeReference             *WorkerObjectType                `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	PositionReference             *PositionElementObjectType       `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	DisplayName                   string                           `xml:"urn:com.workday/bsvc Display_Name,omitempty"`
	ScheduleStartDateTime         *time.Time                       `xml:"urn:com.workday/bsvc Schedule_Start_DateTime,omitempty"`
	ScheduleEndDateTime           *time.Time                       `xml:"urn:com.workday/bsvc Schedule_End_DateTime,omitempty"`
	MealData                      []AdHocScheduleEventMealDataType `xml:"urn:com.workday/bsvc Meal_Data,omitempty"`
	TimeTypeReference             *TimeTypeObjectType              `xml:"urn:com.workday/bsvc Time_Type_Reference,omitempty"`
	OverrideRate                  float64                          `xml:"urn:com.workday/bsvc Override_Rate,omitempty"`
	BusinessUnitReference         *BusinessUnitObjectType          `xml:"urn:com.workday/bsvc Business_Unit_Reference,omitempty"`
	CostCenterReference           *CostCenterObjectType            `xml:"urn:com.workday/bsvc Cost_Center_Reference,omitempty"`
	CustomOrganization01Reference *CustomOrganizationObjectType    `xml:"urn:com.workday/bsvc Custom_Organization_01_Reference,omitempty"`
	CustomOrganization02Reference *CustomOrganizationObjectType    `xml:"urn:com.workday/bsvc Custom_Organization_02_Reference,omitempty"`
	CustomOrganization03Reference *CustomOrganizationObjectType    `xml:"urn:com.workday/bsvc Custom_Organization_03_Reference,omitempty"`
	CustomOrganization04Reference *CustomOrganizationObjectType    `xml:"urn:com.workday/bsvc Custom_Organization_04_Reference,omitempty"`
	CustomOrganization05Reference *CustomOrganizationObjectType    `xml:"urn:com.workday/bsvc Custom_Organization_05_Reference,omitempty"`
	CustomOrganization06Reference *CustomOrganizationObjectType    `xml:"urn:com.workday/bsvc Custom_Organization_06_Reference,omitempty"`
	CustomOrganization07Reference *CustomOrganizationObjectType    `xml:"urn:com.workday/bsvc Custom_Organization_07_Reference,omitempty"`
	CustomOrganization08Reference *CustomOrganizationObjectType    `xml:"urn:com.workday/bsvc Custom_Organization_08_Reference,omitempty"`
	CustomOrganization09Reference *CustomOrganizationObjectType    `xml:"urn:com.workday/bsvc Custom_Organization_09_Reference,omitempty"`
	CustomOrganization10Reference *CustomOrganizationObjectType    `xml:"urn:com.workday/bsvc Custom_Organization_10_Reference,omitempty"`
	CustomWorktag01Reference      *CustomWorktag01ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_01_Reference,omitempty"`
	CustomWorktag02Reference      *CustomWorktag02ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_02_Reference,omitempty"`
	CustomWorktag03Reference      *CustomWorktag03ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_03_Reference,omitempty"`
	CustomWorktag04Reference      *CustomWorktag04ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_04_Reference,omitempty"`
	CustomWorktag05Reference      *CustomWorktag05ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_05_Reference,omitempty"`
	CustomWorktag06Reference      *CustomWorktag06ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_06_Reference,omitempty"`
	CustomWorktag07Reference      *CustomWorktag07ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_07_Reference,omitempty"`
	CustomWorktag08Reference      *CustomWorktag08ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_08_Reference,omitempty"`
	CustomWorktag09Reference      *CustomWorktag09ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_09_Reference,omitempty"`
	CustomWorktag10Reference      *CustomWorktag10ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_10_Reference,omitempty"`
	CustomWorktag11Reference      *CustomWorktag11ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_11_Reference,omitempty"`
	CustomWorktag12Reference      *CustomWorktag12ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_12_Reference,omitempty"`
	CustomWorktag13Reference      *CustomWorktag13ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_13_Reference,omitempty"`
	CustomWorktag14Reference      *CustomWorktag14ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_14_Reference,omitempty"`
	CustomWorktag15Reference      *CustomWorktag15ObjectType       `xml:"urn:com.workday/bsvc Custom_Worktag_15_Reference,omitempty"`
	FundReference                 *FundObjectType                  `xml:"urn:com.workday/bsvc Fund_Reference,omitempty"`
	GiftReference                 *GiftObjectType                  `xml:"urn:com.workday/bsvc Gift_Reference,omitempty"`
	GrantReference                *GrantObjectType                 `xml:"urn:com.workday/bsvc Grant_Reference,omitempty"`
	JobProfileReference           *JobProfileObjectType            `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	LocationReference             *LocationObjectType              `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	ProgramReference              *ProgramObjectType               `xml:"urn:com.workday/bsvc Program_Reference,omitempty"`
	RegionReference               *RegionObjectType                `xml:"urn:com.workday/bsvc Region_Reference,omitempty"`
	Comment                       string                           `xml:"urn:com.workday/bsvc Comment,omitempty"`
	Delete                        bool                             `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

func (t *AdHocScheduleEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AdHocScheduleEventDataType
	var layout struct {
		*T
		ScheduleStartDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Schedule_Start_DateTime,omitempty"`
		ScheduleEndDateTime   *xsdDateTime `xml:"urn:com.workday/bsvc Schedule_End_DateTime,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ScheduleStartDateTime = (*xsdDateTime)(layout.T.ScheduleStartDateTime)
	layout.ScheduleEndDateTime = (*xsdDateTime)(layout.T.ScheduleEndDateTime)
	return e.EncodeElement(layout, start)
}
func (t *AdHocScheduleEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AdHocScheduleEventDataType
	var overlay struct {
		*T
		ScheduleStartDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Schedule_Start_DateTime,omitempty"`
		ScheduleEndDateTime   *xsdDateTime `xml:"urn:com.workday/bsvc Schedule_End_DateTime,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ScheduleStartDateTime = (*xsdDateTime)(overlay.T.ScheduleStartDateTime)
	overlay.ScheduleEndDateTime = (*xsdDateTime)(overlay.T.ScheduleEndDateTime)
	return d.DecodeElement(&overlay, &start)
}

// Data related to the meal(s) to be processed.
type AdHocScheduleEventMealDataType struct {
	MealID  string     `xml:"urn:com.workday/bsvc Meal_ID,omitempty"`
	MealIn  *time.Time `xml:"urn:com.workday/bsvc Meal_In,omitempty"`
	MealOut *time.Time `xml:"urn:com.workday/bsvc Meal_Out,omitempty"`
}

func (t *AdHocScheduleEventMealDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AdHocScheduleEventMealDataType
	var layout struct {
		*T
		MealIn  *xsdTime `xml:"urn:com.workday/bsvc Meal_In,omitempty"`
		MealOut *xsdTime `xml:"urn:com.workday/bsvc Meal_Out,omitempty"`
	}
	layout.T = (*T)(t)
	layout.MealIn = (*xsdTime)(layout.T.MealIn)
	layout.MealOut = (*xsdTime)(layout.T.MealOut)
	return e.EncodeElement(layout, start)
}
func (t *AdHocScheduleEventMealDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AdHocScheduleEventMealDataType
	var overlay struct {
		*T
		MealIn  *xsdTime `xml:"urn:com.workday/bsvc Meal_In,omitempty"`
		MealOut *xsdTime `xml:"urn:com.workday/bsvc Meal_Out,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.MealIn = (*xsdTime)(overlay.T.MealIn)
	overlay.MealOut = (*xsdTime)(overlay.T.MealOut)
	return d.DecodeElement(&overlay, &start)
}

// Contains each work schedule calendar assignment.
type AssignWorkScheduleDataType struct {
	WorkerReference               *WorkerObjectType               `xml:"urn:com.workday/bsvc Worker_Reference"`
	StartDate                     time.Time                       `xml:"urn:com.workday/bsvc Start_Date"`
	EndDate                       *time.Time                      `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	WorkScheduleCalendarReference *WorkScheduleCalendarObjectType `xml:"urn:com.workday/bsvc Work_Schedule_Calendar_Reference"`
}

func (t *AssignWorkScheduleDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AssignWorkScheduleDataType
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
func (t *AssignWorkScheduleDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AssignWorkScheduleDataType
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

// Imports work schedule calendar assignments.
type AssignWorkScheduleRequestType struct {
	BusinessProcessParametersData *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters_Data,omitempty"`
	AssignWorkScheduleData        []AssignWorkScheduleDataType   `xml:"urn:com.workday/bsvc Assign_Work_Schedule_Data"`
	Version                       string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element from importing a work schedule calendar assignment.
type AssignWorkScheduleResponseType struct {
	WorkScheduleAssignmentEventReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Work_Schedule_Assignment_Event_Reference,omitempty"`
	Version                              string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Calculated Time Block Response Data
type CalculatedTimeBlockDataType struct {
	WorkerReference               *WorkerObjectType                  `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	CalculatedDate                *time.Time                         `xml:"urn:com.workday/bsvc Calculated_Date,omitempty"`
	ShiftDate                     *time.Time                         `xml:"urn:com.workday/bsvc Shift_Date,omitempty"`
	InTime                        *time.Time                         `xml:"urn:com.workday/bsvc In_Time,omitempty"`
	OutTime                       *time.Time                         `xml:"urn:com.workday/bsvc Out_Time,omitempty"`
	CalculatedQuantity            float64                            `xml:"urn:com.workday/bsvc Calculated_Quantity,omitempty"`
	StatusReference               *TimeTrackingSetUpOptionObjectType `xml:"urn:com.workday/bsvc Status_Reference,omitempty"`
	IsDeleted                     *bool                              `xml:"urn:com.workday/bsvc Is_Deleted,omitempty"`
	CalculationTagReference       []TimeCalculationTagObjectType     `xml:"urn:com.workday/bsvc Calculation_Tag_Reference,omitempty"`
	LastUpdated                   *time.Time                         `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	BusinessUnitReference         *BusinessUnitObjectType            `xml:"urn:com.workday/bsvc Business_Unit_Reference,omitempty"`
	CostCenterReference           *CostCenterObjectType              `xml:"urn:com.workday/bsvc Cost_Center_Reference,omitempty"`
	FundReference                 *FundObjectType                    `xml:"urn:com.workday/bsvc Fund_Reference,omitempty"`
	GiftReference                 *GiftObjectType                    `xml:"urn:com.workday/bsvc Gift_Reference,omitempty"`
	GrantReference                *GrantObjectType                   `xml:"urn:com.workday/bsvc Grant_Reference,omitempty"`
	JobProfileReference           *JobProfileObjectType              `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	LocationReference             *LocationObjectType                `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	ProgramReference              *ProgramObjectType                 `xml:"urn:com.workday/bsvc Program_Reference,omitempty"`
	RegionReference               *RegionObjectType                  `xml:"urn:com.workday/bsvc Region_Reference,omitempty"`
	CustomWorktag01Reference      *CustomWorktag01ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_01_Reference,omitempty"`
	CustomWorktag02Reference      *CustomWorktag02ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_02_Reference,omitempty"`
	CustomWorktag03Reference      *CustomWorktag03ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_03_Reference,omitempty"`
	CustomWorktag04Reference      *CustomWorktag04ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_04_Reference,omitempty"`
	CustomWorktag05Reference      *CustomWorktag05ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_05_Reference,omitempty"`
	CustomWorktag06Reference      *CustomWorktag06ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_06_Reference,omitempty"`
	CustomWorktag07Reference      *CustomWorktag07ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_07_Reference,omitempty"`
	CustomWorktag08Reference      *CustomWorktag08ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_08_Reference,omitempty"`
	CustomWorktag09Reference      *CustomWorktag09ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_09_Reference,omitempty"`
	CustomWorktag10Reference      *CustomWorktag10ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_10_Reference,omitempty"`
	CustomWorktag11Reference      *CustomWorktag11ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_11_Reference,omitempty"`
	CustomWorktag12Reference      *CustomWorktag12ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_12_Reference,omitempty"`
	CustomWorktag13Reference      *CustomWorktag13ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_13_Reference,omitempty"`
	CustomWorktag14Reference      *CustomWorktag14ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_14_Reference,omitempty"`
	CustomWorktag15Reference      *CustomWorktag15ObjectType         `xml:"urn:com.workday/bsvc Custom_Worktag_15_Reference,omitempty"`
	CustomOrganization01Reference *CustomOrganizationObjectType      `xml:"urn:com.workday/bsvc Custom_Organization_01_Reference,omitempty"`
	CustomOrganization02Reference *CustomOrganizationObjectType      `xml:"urn:com.workday/bsvc Custom_Organization_02_Reference,omitempty"`
	CustomOrganization03Reference *CustomOrganizationObjectType      `xml:"urn:com.workday/bsvc Custom_Organization_03_Reference,omitempty"`
	CustomOrganization04Reference *CustomOrganizationObjectType      `xml:"urn:com.workday/bsvc Custom_Organization_04_Reference,omitempty"`
	CustomOrganization05Reference *CustomOrganizationObjectType      `xml:"urn:com.workday/bsvc Custom_Organization_05_Reference,omitempty"`
	CustomOrganization06Reference *CustomOrganizationObjectType      `xml:"urn:com.workday/bsvc Custom_Organization_06_Reference,omitempty"`
	CustomOrganization07Reference *CustomOrganizationObjectType      `xml:"urn:com.workday/bsvc Custom_Organization_07_Reference,omitempty"`
	CustomOrganization08Reference *CustomOrganizationObjectType      `xml:"urn:com.workday/bsvc Custom_Organization_08_Reference,omitempty"`
	CustomOrganization09Reference *CustomOrganizationObjectType      `xml:"urn:com.workday/bsvc Custom_Organization_09_Reference,omitempty"`
	CustomOrganization10Reference *CustomOrganizationObjectType      `xml:"urn:com.workday/bsvc Custom_Organization_10_Reference,omitempty"`
}

func (t *CalculatedTimeBlockDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CalculatedTimeBlockDataType
	var layout struct {
		*T
		CalculatedDate *xsdDate     `xml:"urn:com.workday/bsvc Calculated_Date,omitempty"`
		ShiftDate      *xsdDate     `xml:"urn:com.workday/bsvc Shift_Date,omitempty"`
		InTime         *xsdDateTime `xml:"urn:com.workday/bsvc In_Time,omitempty"`
		OutTime        *xsdDateTime `xml:"urn:com.workday/bsvc Out_Time,omitempty"`
		LastUpdated    *xsdDateTime `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CalculatedDate = (*xsdDate)(layout.T.CalculatedDate)
	layout.ShiftDate = (*xsdDate)(layout.T.ShiftDate)
	layout.InTime = (*xsdDateTime)(layout.T.InTime)
	layout.OutTime = (*xsdDateTime)(layout.T.OutTime)
	layout.LastUpdated = (*xsdDateTime)(layout.T.LastUpdated)
	return e.EncodeElement(layout, start)
}
func (t *CalculatedTimeBlockDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CalculatedTimeBlockDataType
	var overlay struct {
		*T
		CalculatedDate *xsdDate     `xml:"urn:com.workday/bsvc Calculated_Date,omitempty"`
		ShiftDate      *xsdDate     `xml:"urn:com.workday/bsvc Shift_Date,omitempty"`
		InTime         *xsdDateTime `xml:"urn:com.workday/bsvc In_Time,omitempty"`
		OutTime        *xsdDateTime `xml:"urn:com.workday/bsvc Out_Time,omitempty"`
		LastUpdated    *xsdDateTime `xml:"urn:com.workday/bsvc Last_Updated,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CalculatedDate = (*xsdDate)(overlay.T.CalculatedDate)
	overlay.ShiftDate = (*xsdDate)(overlay.T.ShiftDate)
	overlay.InTime = (*xsdDateTime)(overlay.T.InTime)
	overlay.OutTime = (*xsdDateTime)(overlay.T.OutTime)
	overlay.LastUpdated = (*xsdDateTime)(overlay.T.LastUpdated)
	return d.DecodeElement(&overlay, &start)
}

// Time Block Request Criteria
type CalculatedTimeBlockRequestCriteriaType struct {
	StartDate                        time.Time                           `xml:"urn:com.workday/bsvc Start_Date"`
	EndDate                          time.Time                           `xml:"urn:com.workday/bsvc End_Date"`
	WorkerReference                  []WorkerObjectType                  `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	SupervisoryOrganizationReference []SupervisoryOrganizationObjectType `xml:"urn:com.workday/bsvc Supervisory_Organization_Reference,omitempty"`
	CalculationTagReference          []TimeCalculationTagObjectType      `xml:"urn:com.workday/bsvc Calculation_Tag_Reference,omitempty"`
	StatusReference                  []TimeTrackingSetUpOptionObjectType `xml:"urn:com.workday/bsvc Status_Reference,omitempty"`
}

func (t *CalculatedTimeBlockRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CalculatedTimeBlockRequestCriteriaType
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
func (t *CalculatedTimeBlockRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CalculatedTimeBlockRequestCriteriaType
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

// Time Block Request References
type CalculatedTimeBlockRequestReferencesType struct {
	WorkerTimeBlockReference []WorkerTimeBlockObjectType `xml:"urn:com.workday/bsvc Worker_Time_Block_Reference"`
}

// Calculated Time Block Response Data
type CalculatedTimeBlockResponseDataType struct {
	CalculatedTimeBlock []CalculatedTimeBlockType `xml:"urn:com.workday/bsvc Calculated_Time_Block,omitempty"`
}

// Time Block Response Group
type CalculatedTimeBlockResponseGroupType struct {
	IncludeWorker             *bool `xml:"urn:com.workday/bsvc Include_Worker,omitempty"`
	IncludeDate               *bool `xml:"urn:com.workday/bsvc Include_Date,omitempty"`
	IncludeInOutTime          *bool `xml:"urn:com.workday/bsvc Include_In_Out_Time,omitempty"`
	IncludeCalculatedQuantity *bool `xml:"urn:com.workday/bsvc Include_Calculated_Quantity,omitempty"`
	IncludeStatus             *bool `xml:"urn:com.workday/bsvc Include_Status,omitempty"`
	IncludeDeleted            *bool `xml:"urn:com.workday/bsvc Include_Deleted,omitempty"`
	IncludeCalculationTags    *bool `xml:"urn:com.workday/bsvc Include_Calculation_Tags,omitempty"`
	IncludeLastUpdated        *bool `xml:"urn:com.workday/bsvc Include_Last_Updated,omitempty"`
	IncludeWorktags           *bool `xml:"urn:com.workday/bsvc Include_Worktags,omitempty"`
}

// Calculated Time Block
type CalculatedTimeBlockType struct {
	WorkerTimeBlockReference *WorkerTimeBlockObjectType   `xml:"urn:com.workday/bsvc Worker_Time_Block_Reference,omitempty"`
	CalculatedTimeBlockData  *CalculatedTimeBlockDataType `xml:"urn:com.workday/bsvc Calculated_Time_Block_Data,omitempty"`
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
type EventAttachmentCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EventAttachmentCategoryObjectType struct {
	ID         []EventAttachmentCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Get Calculated Time Blocks Request
type GetCalculatedTimeBlocksRequestType struct {
	RequestReferences *CalculatedTimeBlockRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CalculatedTimeBlockRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CalculatedTimeBlockResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Calculated Time Block Response
type GetCalculatedTimeBlocksResponseType struct {
	RequestReferences *CalculatedTimeBlockRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CalculatedTimeBlockRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CalculatedTimeBlockResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   []ResponseResultsType                     `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *CalculatedTimeBlockResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Load instances of Ad Hoc Schedule Events.
type ImportAdHocSchedulesRequestType struct {
	AdHocScheduleData []AdHocScheduleEventDataType `xml:"urn:com.workday/bsvc Ad-Hoc_Schedule_Data,omitempty"`
	Version           string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Load instances of Reported Time Blocks.
type ImportReportedTimeBlocksRequestType struct {
	ReportedTimeBlockData []ReportedTimeBlockDataType `xml:"urn:com.workday/bsvc Reported_Time_Block_Data,omitempty"`
	Version               string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Load instances of External Time Clock Events.
type ImportTimeClockEventsRequestType struct {
	TimeClockEventData []TimeClockEventDataType `xml:"urn:com.workday/bsvc Time_Clock_Event_Data,omitempty"`
	Version            string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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
type JobProfileObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobProfileObjectType struct {
	ID         []JobProfileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type PositionElementObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PositionElementObjectType struct {
	ID         []PositionElementObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
type ProjectPlanTaskObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProjectPlanTaskObjectType struct {
	ID         []ProjectPlanTaskObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Put Import Process Response element
type PutImportProcessResponseType struct {
	ImportProcessReference  *WebServiceBackgroundProcessRuntimeObjectType `xml:"urn:com.workday/bsvc Import_Process_Reference,omitempty"`
	HeaderInstanceReference *InstanceObjectType                           `xml:"urn:com.workday/bsvc Header_Instance_Reference,omitempty"`
	Version                 string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Load instances of External Time Clock Events.
type PutTimeClockEventsRequestType struct {
	TimeClockEventData []TimeClockEventDataType `xml:"urn:com.workday/bsvc Time_Clock_Event_Data"`
	Version            string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element from loading an instance of an External Time Clock Event
type PutTimeClockEventsResponseType struct {
	ResponseText string `xml:"urn:com.workday/bsvc Response_Text,omitempty"`
	Version      string `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Data related to the worker time block(s) to be processed.
type ReportedTimeBlockDataType struct {
	WorkerReference               *WorkerObjectType             `xml:"urn:com.workday/bsvc Worker_Reference"`
	TimeEntryCodeReference        *TimeEntryCodeObjectType      `xml:"urn:com.workday/bsvc Time_Entry_Code_Reference,omitempty"`
	ProjectReference              *ProjectObjectType            `xml:"urn:com.workday/bsvc Project_Reference,omitempty"`
	ProjectTaskReference          *ProjectPlanTaskObjectType    `xml:"urn:com.workday/bsvc Project_Task_Reference,omitempty"`
	Date                          *time.Time                    `xml:"urn:com.workday/bsvc Date,omitempty"`
	Quantity                      float64                       `xml:"urn:com.workday/bsvc Quantity,omitempty"`
	InDateTime                    *time.Time                    `xml:"urn:com.workday/bsvc In_Date_Time,omitempty"`
	OutDateTime                   *time.Time                    `xml:"urn:com.workday/bsvc Out_Date_Time,omitempty"`
	OutReasonReference            *TimePunchTypeObjectType      `xml:"urn:com.workday/bsvc Out_Reason_Reference,omitempty"`
	PositionReference             []PositionElementObjectType   `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	DoNotBill                     *bool                         `xml:"urn:com.workday/bsvc Do_Not_Bill,omitempty"`
	OverrideRate                  float64                       `xml:"urn:com.workday/bsvc Override_Rate,omitempty"`
	BusinessUnitReference         *BusinessUnitObjectType       `xml:"urn:com.workday/bsvc Business_Unit_Reference,omitempty"`
	CostCenterReference           *CostCenterObjectType         `xml:"urn:com.workday/bsvc Cost_Center_Reference,omitempty"`
	CustomOrganization01Reference *CustomOrganizationObjectType `xml:"urn:com.workday/bsvc Custom_Organization_01_Reference,omitempty"`
	CustomOrganization02Reference *CustomOrganizationObjectType `xml:"urn:com.workday/bsvc Custom_Organization_02_Reference,omitempty"`
	CustomOrganization03Reference *CustomOrganizationObjectType `xml:"urn:com.workday/bsvc Custom_Organization_03_Reference,omitempty"`
	CustomOrganization04Reference *CustomOrganizationObjectType `xml:"urn:com.workday/bsvc Custom_Organization_04_Reference,omitempty"`
	CustomOrganization05Reference *CustomOrganizationObjectType `xml:"urn:com.workday/bsvc Custom_Organization_05_Reference,omitempty"`
	CustomOrganization06Reference *CustomOrganizationObjectType `xml:"urn:com.workday/bsvc Custom_Organization_06_Reference,omitempty"`
	CustomOrganization07Reference *CustomOrganizationObjectType `xml:"urn:com.workday/bsvc Custom_Organization_07_Reference,omitempty"`
	CustomOrganization08Reference *CustomOrganizationObjectType `xml:"urn:com.workday/bsvc Custom_Organization_08_Reference,omitempty"`
	CustomOrganization09Reference *CustomOrganizationObjectType `xml:"urn:com.workday/bsvc Custom_Organization_09_Reference,omitempty"`
	CustomOrganization10Reference *CustomOrganizationObjectType `xml:"urn:com.workday/bsvc Custom_Organization_10_Reference,omitempty"`
	CustomWorktag01Reference      *CustomWorktag01ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_01_Reference,omitempty"`
	CustomWorktag02Reference      *CustomWorktag02ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_02_Reference,omitempty"`
	CustomWorktag03Reference      *CustomWorktag03ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_03_Reference,omitempty"`
	CustomWorktag04Reference      *CustomWorktag04ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_04_Reference,omitempty"`
	CustomWorktag05Reference      *CustomWorktag05ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_05_Reference,omitempty"`
	CustomWorktag06Reference      *CustomWorktag06ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_06_Reference,omitempty"`
	CustomWorktag07Reference      *CustomWorktag07ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_07_Reference,omitempty"`
	CustomWorktag08Reference      *CustomWorktag08ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_08_Reference,omitempty"`
	CustomWorktag09Reference      *CustomWorktag09ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_09_Reference,omitempty"`
	CustomWorktag10Reference      *CustomWorktag10ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_10_Reference,omitempty"`
	CustomWorktag11Reference      *CustomWorktag11ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_11_Reference,omitempty"`
	CustomWorktag12Reference      *CustomWorktag12ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_12_Reference,omitempty"`
	CustomWorktag13Reference      *CustomWorktag13ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_13_Reference,omitempty"`
	CustomWorktag14Reference      *CustomWorktag14ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_14_Reference,omitempty"`
	CustomWorktag15Reference      *CustomWorktag15ObjectType    `xml:"urn:com.workday/bsvc Custom_Worktag_15_Reference,omitempty"`
	FundReference                 *FundObjectType               `xml:"urn:com.workday/bsvc Fund_Reference,omitempty"`
	GiftReference                 *GiftObjectType               `xml:"urn:com.workday/bsvc Gift_Reference,omitempty"`
	GrantReference                *GrantObjectType              `xml:"urn:com.workday/bsvc Grant_Reference,omitempty"`
	JobProfileReference           *JobProfileObjectType         `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	LocationReference             *LocationObjectType           `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	ProgramReference              *ProgramObjectType            `xml:"urn:com.workday/bsvc Program_Reference,omitempty"`
	RegionReference               *RegionObjectType             `xml:"urn:com.workday/bsvc Region_Reference,omitempty"`
	Comment                       string                        `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

func (t *ReportedTimeBlockDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ReportedTimeBlockDataType
	var layout struct {
		*T
		Date        *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		InDateTime  *xsdDateTime `xml:"urn:com.workday/bsvc In_Date_Time,omitempty"`
		OutDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Out_Date_Time,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Date = (*xsdDate)(layout.T.Date)
	layout.InDateTime = (*xsdDateTime)(layout.T.InDateTime)
	layout.OutDateTime = (*xsdDateTime)(layout.T.OutDateTime)
	return e.EncodeElement(layout, start)
}
func (t *ReportedTimeBlockDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ReportedTimeBlockDataType
	var overlay struct {
		*T
		Date        *xsdDate     `xml:"urn:com.workday/bsvc Date,omitempty"`
		InDateTime  *xsdDateTime `xml:"urn:com.workday/bsvc In_Date_Time,omitempty"`
		OutDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Out_Date_Time,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Date = (*xsdDate)(overlay.T.Date)
	overlay.InDateTime = (*xsdDateTime)(overlay.T.InDateTime)
	overlay.OutDateTime = (*xsdDateTime)(overlay.T.OutDateTime)
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
type SupervisoryOrganizationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SupervisoryOrganizationObjectType struct {
	ID         []SupervisoryOrganizationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TimeCalculationTagObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimeCalculationTagObjectType struct {
	ID         []TimeCalculationTagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Data related to the time clock event(s) to be processed.
type TimeClockEventDataType struct {
	TimeClockEventID        string                   `xml:"urn:com.workday/bsvc Time_Clock_Event_ID,omitempty"`
	DeleteTimeClockEvent    *bool                    `xml:"urn:com.workday/bsvc Delete_Time_Clock_Event,omitempty"`
	EmployeeID              string                   `xml:"urn:com.workday/bsvc Employee_ID,omitempty"`
	PositionID              string                   `xml:"urn:com.workday/bsvc Position_ID,omitempty"`
	TimeClockEventDateTime  *time.Time               `xml:"urn:com.workday/bsvc Time_Clock_Event_Date_Time,omitempty"`
	TimeEntryCode           string                   `xml:"urn:com.workday/bsvc Time_Entry_Code,omitempty"`
	ProjectID               string                   `xml:"urn:com.workday/bsvc Project_ID,omitempty"`
	TaskID                  string                   `xml:"urn:com.workday/bsvc Task_ID,omitempty"`
	ClockEventTypeReference *TimePunchTypeObjectType `xml:"urn:com.workday/bsvc Clock_Event_Type_Reference,omitempty"`
	OverrideRate            float64                  `xml:"urn:com.workday/bsvc Override_Rate,omitempty"`
	BusinessUnit            string                   `xml:"urn:com.workday/bsvc Business_Unit,omitempty"`
	CostCenter              string                   `xml:"urn:com.workday/bsvc Cost_Center,omitempty"`
	CustomOrganization01    string                   `xml:"urn:com.workday/bsvc Custom_Organization_01,omitempty"`
	CustomOrganization02    string                   `xml:"urn:com.workday/bsvc Custom_Organization_02,omitempty"`
	CustomOrganization03    string                   `xml:"urn:com.workday/bsvc Custom_Organization_03,omitempty"`
	CustomOrganization04    string                   `xml:"urn:com.workday/bsvc Custom_Organization_04,omitempty"`
	CustomOrganization05    string                   `xml:"urn:com.workday/bsvc Custom_Organization_05,omitempty"`
	CustomOrganization06    string                   `xml:"urn:com.workday/bsvc Custom_Organization_06,omitempty"`
	CustomOrganization07    string                   `xml:"urn:com.workday/bsvc Custom_Organization_07,omitempty"`
	CustomOrganization08    string                   `xml:"urn:com.workday/bsvc Custom_Organization_08,omitempty"`
	CustomOrganization09    string                   `xml:"urn:com.workday/bsvc Custom_Organization_09,omitempty"`
	CustomOrganization10    string                   `xml:"urn:com.workday/bsvc Custom_Organization_10,omitempty"`
	CustomWorktag01         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_01,omitempty"`
	CustomWorktag02         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_02,omitempty"`
	CustomWorktag03         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_03,omitempty"`
	CustomWorktag04         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_04,omitempty"`
	CustomWorktag05         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_05,omitempty"`
	CustomWorktag06         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_06,omitempty"`
	CustomWorktag07         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_07,omitempty"`
	CustomWorktag08         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_08,omitempty"`
	CustomWorktag09         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_09,omitempty"`
	CustomWorktag10         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_10,omitempty"`
	CustomWorktag11         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_11,omitempty"`
	CustomWorktag12         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_12,omitempty"`
	CustomWorktag13         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_13,omitempty"`
	CustomWorktag14         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_14,omitempty"`
	CustomWorktag15         string                   `xml:"urn:com.workday/bsvc Custom_Worktag_15,omitempty"`
	Fund                    string                   `xml:"urn:com.workday/bsvc Fund,omitempty"`
	Gift                    string                   `xml:"urn:com.workday/bsvc Gift,omitempty"`
	Grant                   string                   `xml:"urn:com.workday/bsvc Grant,omitempty"`
	JobProfile              string                   `xml:"urn:com.workday/bsvc Job_Profile,omitempty"`
	Location                string                   `xml:"urn:com.workday/bsvc Location,omitempty"`
	Program                 string                   `xml:"urn:com.workday/bsvc Program,omitempty"`
	Region                  string                   `xml:"urn:com.workday/bsvc Region,omitempty"`
	Comment                 string                   `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

func (t *TimeClockEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T TimeClockEventDataType
	var layout struct {
		*T
		TimeClockEventDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Time_Clock_Event_Date_Time,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TimeClockEventDateTime = (*xsdDateTime)(layout.T.TimeClockEventDateTime)
	return e.EncodeElement(layout, start)
}
func (t *TimeClockEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T TimeClockEventDataType
	var overlay struct {
		*T
		TimeClockEventDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Time_Clock_Event_Date_Time,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TimeClockEventDateTime = (*xsdDateTime)(overlay.T.TimeClockEventDateTime)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type TimeEntryCodeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimeEntryCodeObjectType struct {
	ID         []TimeEntryCodeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TimePunchTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimePunchTypeObjectType struct {
	ID         []TimePunchTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TimeTrackingSetUpOptionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimeTrackingSetUpOptionObjectType struct {
	ID         []TimeTrackingSetUpOptionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TimeTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimeTypeObjectType struct {
	ID         []TimeTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type WorkScheduleCalendarObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkScheduleCalendarObjectType struct {
	ID         []WorkScheduleCalendarObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type WorkerTimeBlockObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkerTimeBlockObjectType struct {
	ID         []WorkerTimeBlockObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
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
