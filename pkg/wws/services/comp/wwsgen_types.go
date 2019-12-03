package comp

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// Contains a unique identifier for an instance of an object.
type AcademicPeriodObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AcademicPeriodObjectType struct {
	ID         []AcademicPeriodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Data element for the Request Stock Grant business process.
type AddIndividualStockGrantDataType struct {
	StockGrantReferenceID        string                          `xml:"urn:com.workday/bsvc Stock_Grant_Reference_ID,omitempty"`
	GrantTypeReference           *StockGrantTypeObjectType       `xml:"urn:com.workday/bsvc Grant_Type_Reference"`
	GrantID                      string                          `xml:"urn:com.workday/bsvc Grant_ID,omitempty"`
	GrantDate                    *time.Time                      `xml:"urn:com.workday/bsvc Grant_Date,omitempty"`
	VestingScheduleReference     *StockVestingScheduleObjectType `xml:"urn:com.workday/bsvc Vesting_Schedule_Reference,omitempty"`
	VestFromDate                 *time.Time                      `xml:"urn:com.workday/bsvc Vest_From_Date,omitempty"`
	ExpirationDate               *time.Time                      `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	SharesGranted                float64                         `xml:"urn:com.workday/bsvc Shares_Granted,omitempty"`
	GrantPercent                 float64                         `xml:"urn:com.workday/bsvc Grant_Percent,omitempty"`
	GrantAmount                  float64                         `xml:"urn:com.workday/bsvc Grant_Amount,omitempty"`
	GrantAmountCurrencyReference *CurrencyObjectType             `xml:"urn:com.workday/bsvc Grant_Amount_Currency_Reference,omitempty"`
	GrantPrice                   float64                         `xml:"urn:com.workday/bsvc Grant_Price,omitempty"`
	GrantPriceCurrencyReference  *CurrencyObjectType             `xml:"urn:com.workday/bsvc Grant_Price_Currency_Reference,omitempty"`
	OptionsPricingFactor         float64                         `xml:"urn:com.workday/bsvc Options_Pricing_Factor,omitempty"`
	BoardApproved                *bool                           `xml:"urn:com.workday/bsvc Board_Approved,omitempty"`
	Comments                     string                          `xml:"urn:com.workday/bsvc Comments,omitempty"`
	SharesVested                 float64                         `xml:"urn:com.workday/bsvc Shares_Vested,omitempty"`
	SharesUnvested               float64                         `xml:"urn:com.workday/bsvc Shares_Unvested,omitempty"`
	LongTermCashAmountVested     float64                         `xml:"urn:com.workday/bsvc Long_Term_Cash_Amount_Vested,omitempty"`
	LongTermCashAmountUnvested   float64                         `xml:"urn:com.workday/bsvc Long_Term_Cash_Amount_Unvested,omitempty"`
	VestingPrice                 float64                         `xml:"urn:com.workday/bsvc Vesting_Price,omitempty"`
	VestedAsOf                   *time.Time                      `xml:"urn:com.workday/bsvc Vested_As_Of,omitempty"`
}

func (t *AddIndividualStockGrantDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AddIndividualStockGrantDataType
	var layout struct {
		*T
		GrantDate      *xsdDate `xml:"urn:com.workday/bsvc Grant_Date,omitempty"`
		VestFromDate   *xsdDate `xml:"urn:com.workday/bsvc Vest_From_Date,omitempty"`
		ExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
		VestedAsOf     *xsdDate `xml:"urn:com.workday/bsvc Vested_As_Of,omitempty"`
	}
	layout.T = (*T)(t)
	layout.GrantDate = (*xsdDate)(layout.T.GrantDate)
	layout.VestFromDate = (*xsdDate)(layout.T.VestFromDate)
	layout.ExpirationDate = (*xsdDate)(layout.T.ExpirationDate)
	layout.VestedAsOf = (*xsdDate)(layout.T.VestedAsOf)
	return e.EncodeElement(layout, start)
}
func (t *AddIndividualStockGrantDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AddIndividualStockGrantDataType
	var overlay struct {
		*T
		GrantDate      *xsdDate `xml:"urn:com.workday/bsvc Grant_Date,omitempty"`
		VestFromDate   *xsdDate `xml:"urn:com.workday/bsvc Vest_From_Date,omitempty"`
		ExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
		VestedAsOf     *xsdDate `xml:"urn:com.workday/bsvc Vested_As_Of,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.GrantDate = (*xsdDate)(overlay.T.GrantDate)
	overlay.VestFromDate = (*xsdDate)(overlay.T.VestFromDate)
	overlay.ExpirationDate = (*xsdDate)(overlay.T.ExpirationDate)
	overlay.VestedAsOf = (*xsdDate)(overlay.T.VestedAsOf)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper element for the Request Stock Grant business process.
type AddStockGrantDataType struct {
	EffectiveDate      time.Time                                 `xml:"urn:com.workday/bsvc Effective_Date"`
	EmployeeReference  *EmployeeObjectType                       `xml:"urn:com.workday/bsvc Employee_Reference"`
	PositionReference  *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	StockPlanReference *StockPlanObjectType                      `xml:"urn:com.workday/bsvc Stock_Plan_Reference"`
	ReasonReference    *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
	StockGrantData     []AddIndividualStockGrantDataType         `xml:"urn:com.workday/bsvc Stock_Grant_Data"`
}

func (t *AddStockGrantDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AddStockGrantDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *AddStockGrantDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AddStockGrantDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// This web service operation is designed to grant a stock grant to an employee using the Request Stock Grant business process.
type AddStockGrantRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AddStockGrantData         *AddStockGrantDataType         `xml:"urn:com.workday/bsvc Add_Stock_Grant_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the Event ID for the Request Stock Grant Event.
type AddStockGrantResponseType struct {
	StockPackageEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Stock_Package_Event_Reference,omitempty"`
	Version                    string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains costing allocation for Period Activity Pay assignment line.
type AllocationDetailforPeriodPayDataType struct {
	Order                   string                             `xml:"urn:com.workday/bsvc Order,omitempty"`
	CostingWorktagReference []TenantedPayrollWorktagObjectType `xml:"urn:com.workday/bsvc Costing_Worktag_Reference,omitempty"`
	DistributionPercent     float64                            `xml:"urn:com.workday/bsvc Distribution_Percent,omitempty"`
	DistributionAmount      float64                            `xml:"urn:com.workday/bsvc Distribution_Amount,omitempty"`
}

// This is a wrapper for the Allowance Plan Amount Based Profile Data; this is used to display or replace Allowance Plan Amount Based Profile Data.
type AllowancePlanAmountBasedProfileReplacementDataType struct {
	EligibilityRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	Amount                   float64                  `xml:"urn:com.workday/bsvc Amount,omitempty"`
	CurrencyReference        *CurrencyObjectType      `xml:"urn:com.workday/bsvc Currency_Reference"`
}

// Contains Allowance Plan Amount Data.
type AllowancePlanAmountDataType struct {
	Amount                                         float64                                              `xml:"urn:com.workday/bsvc Amount,omitempty"`
	AllowancePlanAmountBasedProfileReplacementData []AllowancePlanAmountBasedProfileReplacementDataType `xml:"urn:com.workday/bsvc Allowance_Plan_Amount_Based_Profile_Replacement_Data,omitempty"`
}

// Allowance Plan Data consists of common information from the following Allowance Plans: Amount, Percent, and Unit.  It also must contain specific Amount, Percent, or Unit data information.
type AllowancePlanDataType struct {
	CompensationElementReference  *CompensationPayEarningObjectType            `xml:"urn:com.workday/bsvc Compensation_Element_Reference,omitempty"`
	SetupSecuritySegmentReference []CompensationSetupSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Setup_Security_Segment_Reference,omitempty"`
	CurrencyReference             *CurrencyObjectType                          `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference            *FrequencyObjectType                         `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	NoIndividualOverride          *bool                                        `xml:"urn:com.workday/bsvc No_Individual_Override,omitempty"`
	ApplyFullTimeEquivalent       *bool                                        `xml:"urn:com.workday/bsvc Apply_Full_Time_Equivalent,omitempty"`
	ExcludefromMerit              *bool                                        `xml:"urn:com.workday/bsvc Exclude_from_Merit,omitempty"`
	AmountData                    *AllowancePlanAmountDataType                 `xml:"urn:com.workday/bsvc Amount_Data"`
	PercentData                   *AllowancePlanPercentDataType                `xml:"urn:com.workday/bsvc Percent_Data"`
	UnitData                      *AllowancePlanUnitDataType                   `xml:"urn:com.workday/bsvc Unit_Data"`
	ReimbursableData              *AllowancePlanReimbursableDataType           `xml:"urn:com.workday/bsvc Reimbursable_Data,omitempty"`
}

// This is a wrapper for the Allowance Plan Percent Based Profile Data; this is used to display or replace Allowance Plan Percent Based Profile Data.
type AllowancePlanPercentBasedProfileReplacementDataType struct {
	EligibilityRuleReference   *ConditionRuleObjectType     `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	Percent                    float64                      `xml:"urn:com.workday/bsvc Percent,omitempty"`
	CompensationBasisReference *CompensationBasisObjectType `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
	CeilingAmount              float64                      `xml:"urn:com.workday/bsvc Ceiling_Amount,omitempty"`
	CeilingCurrencyReference   *CurrencyObjectType          `xml:"urn:com.workday/bsvc Ceiling_Currency_Reference,omitempty"`
}

// Contains Allowance Plan Percent data.
type AllowancePlanPercentDataType struct {
	Percentage                                      float64                                               `xml:"urn:com.workday/bsvc Percentage,omitempty"`
	CompensationBasisReference                      *CompensationBasisObjectType                          `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
	CeilingAmount                                   float64                                               `xml:"urn:com.workday/bsvc Ceiling_Amount,omitempty"`
	CeilingCurrencyReference                        *CurrencyObjectType                                   `xml:"urn:com.workday/bsvc Ceiling_Currency_Reference,omitempty"`
	AllowancePlanPercentBasedProfileReplacementData []AllowancePlanPercentBasedProfileReplacementDataType `xml:"urn:com.workday/bsvc Allowance_Plan_Percent_Based_Profile_Replacement_Data,omitempty"`
}

// Allowance Plan Reimbursable Data consists of information pertinent only to reimbursable allowance plans.
type AllowancePlanReimbursableDataType struct {
	ExpenseItemReference        *ExpenseItemObjectType        `xml:"urn:com.workday/bsvc Expense_Item_Reference"`
	ExpenseAccumulatorReference *ExpenseAccumulatorObjectType `xml:"urn:com.workday/bsvc Expense_Accumulator_Reference"`
}

// This is a wrapper for the Allowance Plan Unit Based Profile Data; this is used to display or replace Allowance Plan Unit Based Profile Data.
type AllowancePlanUnitBasedProfileReplacementDataType struct {
	EligibilityRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	NumberofUnits            float64                  `xml:"urn:com.workday/bsvc Number_of_Units,omitempty"`
	PerUnitAmount            float64                  `xml:"urn:com.workday/bsvc Per_Unit_Amount,omitempty"`
	CurrencyReference        *CurrencyObjectType      `xml:"urn:com.workday/bsvc Currency_Reference"`
}

// Contains Allowance Plan Unit data.
type AllowancePlanUnitDataType struct {
	PerUnitAmount                                float64                                            `xml:"urn:com.workday/bsvc Per_Unit_Amount,omitempty"`
	DefaultUnits                                 float64                                            `xml:"urn:com.workday/bsvc Default_Units,omitempty"`
	UnitOfMeasureReference                       *UnitofMeasureObjectType                           `xml:"urn:com.workday/bsvc Unit_Of_Measure_Reference"`
	AllowancePlanUnitBasedProfileReplacementData []AllowancePlanUnitBasedProfileReplacementDataType `xml:"urn:com.workday/bsvc Allowance_Plan_Unit_Based_Profile_Replacement_Data,omitempty"`
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

// The alternative compensation pay range.
type AlternatePayRangeDataType struct {
	CompensationBasisReference *CompensationBasisConfigurableObjectType `xml:"urn:com.workday/bsvc Compensation_Basis_Reference"`
	Minimum                    float64                                  `xml:"urn:com.workday/bsvc Minimum,omitempty"`
	Midpoint                   float64                                  `xml:"urn:com.workday/bsvc Midpoint,omitempty"`
	Maximum                    float64                                  `xml:"urn:com.workday/bsvc Maximum,omitempty"`
	Spread                     float64                                  `xml:"urn:com.workday/bsvc Spread,omitempty"`
	Segment1Top                float64                                  `xml:"urn:com.workday/bsvc Segment_1_Top,omitempty"`
	Segment2Top                float64                                  `xml:"urn:com.workday/bsvc Segment_2_Top,omitempty"`
	Segment3Top                float64                                  `xml:"urn:com.workday/bsvc Segment_3_Top,omitempty"`
	Segment4Top                float64                                  `xml:"urn:com.workday/bsvc Segment_4_Top,omitempty"`
	Segment5Top                float64                                  `xml:"urn:com.workday/bsvc Segment_5_Top,omitempty"`
}

// Contains an alternative pay range and its associated data.
type AlternativePayRangeType struct {
	Delete                *bool                      `xml:"urn:com.workday/bsvc Delete,omitempty"`
	AlternatePayRangeData *AlternatePayRangeDataType `xml:"urn:com.workday/bsvc Alternate_Pay_Range_Data"`
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

// Wrapper Element for the Assign Employee's ligible Activities business process.
type AssignEligiblePeriodActivitiesforEmployeeDataType struct {
	EmployeeReference               *EmployeeObjectType                       `xml:"urn:com.workday/bsvc Employee_Reference"`
	PositionReference               *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	EligibleActivityDate            time.Time                                 `xml:"urn:com.workday/bsvc Eligible_Activity_Date"`
	ReasonReference                 *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference"`
	EligiblePeriodActivitiesSubData []PeriodActivityEligibilityEntryDataType  `xml:"urn:com.workday/bsvc Eligible_Period_Activities_Sub_Data,omitempty"`
}

func (t *AssignEligiblePeriodActivitiesforEmployeeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AssignEligiblePeriodActivitiesforEmployeeDataType
	var layout struct {
		*T
		EligibleActivityDate *xsdDate `xml:"urn:com.workday/bsvc Eligible_Activity_Date"`
	}
	layout.T = (*T)(t)
	layout.EligibleActivityDate = (*xsdDate)(&layout.T.EligibleActivityDate)
	return e.EncodeElement(layout, start)
}
func (t *AssignEligiblePeriodActivitiesforEmployeeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AssignEligiblePeriodActivitiesforEmployeeDataType
	var overlay struct {
		*T
		EligibleActivityDate *xsdDate `xml:"urn:com.workday/bsvc Eligible_Activity_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EligibleActivityDate = (*xsdDate)(&overlay.T.EligibleActivityDate)
	return d.DecodeElement(&overlay, &start)
}

// Web service operation to assign a period activity for an employee.
type AssignEligiblePeriodActivitiesforEmployeeRequestType struct {
	BusinessProcessParameters                     *BusinessProcessParametersType                     `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AssignEligiblePeriodActivitiesforEmployeeData *AssignEligiblePeriodActivitiesforEmployeeDataType `xml:"urn:com.workday/bsvc Assign_Eligible_Period_Activities_for_Employee_Data"`
	Version                                       string                                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for the Assign Employee's Eligible Period Activities business process.
type AssignEligiblePeriodActivitiesforEmployeeResponseType struct {
	PeriodActivityEligibilityEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Period_Activity_Eligibility_Event_Reference,omitempty"`
	Version                                 string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Base Pay Plan Data consists of common information from the following Base Pay Plans: Salary, Hourly, and Unit.
type BasePayPlanDataType struct {
	CompensationElementReference  *CompensationPayEarningObjectType            `xml:"urn:com.workday/bsvc Compensation_Element_Reference"`
	SetupSecuritySegmentReference []CompensationSetupSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Setup_Security_Segment_Reference,omitempty"`
	CurrencyReference             *CurrencyObjectType                          `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference            *FrequencyObjectType                         `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	ExcludefromMerit              *bool                                        `xml:"urn:com.workday/bsvc Exclude_from_Merit,omitempty"`
	SalaryData                    *SalaryPlanDataType                          `xml:"urn:com.workday/bsvc Salary_Data"`
	HourlyData                    *HourlyPlanDataType                          `xml:"urn:com.workday/bsvc Hourly_Data"`
	UnitData                      *UnitSalaryPlanDataType                      `xml:"urn:com.workday/bsvc Unit_Data"`
}

// Benchmark Job Composite Data
type BenchmarkJobCompositeDataType struct {
	BenchmarkCompositeCategoryReference        *CompensationBenchmarkCompositeCategoryObjectType `xml:"urn:com.workday/bsvc Benchmark_Composite_Category_Reference"`
	CompetitiveMarketTargetPercentileReference *CompensationBenchmarkPercentileObjectType        `xml:"urn:com.workday/bsvc Competitive_Market_Target_Percentile_Reference,omitempty"`
	TargetSpread                               float64                                           `xml:"urn:com.workday/bsvc Target_Spread__,omitempty"`
	BenchmarkAmountReplacementData             []CompensationBenchmarkAmountReplacmentDataType   `xml:"urn:com.workday/bsvc Benchmark_Amount_Replacement_Data"`
	Delete                                     bool                                              `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Benchmark Job Data
type BenchmarkJobDataType struct {
	BenchmarkJobID                       string                                 `xml:"urn:com.workday/bsvc Benchmark_Job_ID,omitempty"`
	EffectiveDate                        *time.Time                             `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	Inactive                             *bool                                  `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	Name                                 string                                 `xml:"urn:com.workday/bsvc Name,omitempty"`
	SurveyDate                           *time.Time                             `xml:"urn:com.workday/bsvc Survey_Date,omitempty"`
	Description                          string                                 `xml:"urn:com.workday/bsvc Description,omitempty"`
	JobProfileReference                  []JobProfileObjectType                 `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	CurrencyReference                    *CurrencyObjectType                    `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference                   *FrequencyObjectType                   `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	BenchmarkJobCompositeReplacementData []BenchmarkJobCompositeDataType        `xml:"urn:com.workday/bsvc Benchmark_Job_Composite_Replacement_Data,omitempty"`
	BenchmarkProfileData                 []CompensationBenchmarkProfileDataType `xml:"urn:com.workday/bsvc Benchmark_Profile_Data,omitempty"`
}

func (t *BenchmarkJobDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BenchmarkJobDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		SurveyDate    *xsdDate `xml:"urn:com.workday/bsvc Survey_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	layout.SurveyDate = (*xsdDate)(layout.T.SurveyDate)
	return e.EncodeElement(layout, start)
}
func (t *BenchmarkJobDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BenchmarkJobDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		SurveyDate    *xsdDate `xml:"urn:com.workday/bsvc Survey_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	overlay.SurveyDate = (*xsdDate)(overlay.T.SurveyDate)
	return d.DecodeElement(&overlay, &start)
}

// Benchmark Job Request Criteria
type BenchmarkJobRequestCriteriaType struct {
	IncludeInactive *bool `xml:"urn:com.workday/bsvc Include_Inactive,omitempty"`
}

// Benchmark Job Request References
type BenchmarkJobRequestReferencesType struct {
	BenchmarkJobReference []CompensationBenchmarkDefaultObjectType `xml:"urn:com.workday/bsvc Benchmark_Job_Reference"`
}

// Benchmark Job Response Data
type BenchmarkJobResponseDataType struct {
	BenchmarkJob []BenchmarkJobType `xml:"urn:com.workday/bsvc Benchmark_Job,omitempty"`
}

// Benchmark Job Response Group
type BenchmarkJobResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Benchmark Job
type BenchmarkJobType struct {
	BenchmarkJobReference *CompensationBenchmarkDefaultObjectType `xml:"urn:com.workday/bsvc Benchmark_Job_Reference,omitempty"`
	BenchmarkJobData      []BenchmarkJobDataType                  `xml:"urn:com.workday/bsvc Benchmark_Job_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BenefitMultiplierOrderObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BenefitMultiplierOrderObjectType struct {
	ID         []BenefitMultiplierOrderObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Data for Employee's one-time ad-hoc bonus payment information
type BonusPaymentDataType struct {
	BonusPlanReference *BonusPlanObjectType `xml:"urn:com.workday/bsvc Bonus_Plan_Reference"`
	Percent            float64              `xml:"urn:com.workday/bsvc Percent,omitempty"`
	Amount             float64              `xml:"urn:com.workday/bsvc Amount,omitempty"`
	CurrencyReference  *CurrencyObjectType  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	Comment            string               `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BonusPercentPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BonusPercentPlanObjectType struct {
	ID         []BonusPercentPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains Bonus Plan Amount data.
type BonusPlanAmountDataType struct {
	Amount                           float64                                     `xml:"urn:com.workday/bsvc Amount,omitempty"`
	AmountPlanProfileReplacementData []BonusPlanProfileReplacementAmountDataType `xml:"urn:com.workday/bsvc Amount_Plan_Profile_Replacement_Data,omitempty"`
}

// Bonus Plan Data consists of common information from the following Bonus Plans: Amount and Percent.
type BonusPlanDataType struct {
	CompensationElementReference                          *CompensationPayEarningObjectType            `xml:"urn:com.workday/bsvc Compensation_Element_Reference"`
	SetupSecuritySegmentReference                         []CompensationSetupSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Setup_Security_Segment_Reference,omitempty"`
	CurrencyReference                                     *CurrencyObjectType                          `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference                                    *FrequencyObjectType                         `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	AllowIndividualTargets                                *bool                                        `xml:"urn:com.workday/bsvc Allow_Individual_Targets,omitempty"`
	WaitingPeriodReference                                *EligibilityWaitingPeriodObjectType          `xml:"urn:com.workday/bsvc Waiting_Period_Reference,omitempty"`
	TimeProrationRuleReference                            *TimeProrationRuleObjectType                 `xml:"urn:com.workday/bsvc Time_Proration_Rule_Reference,omitempty"`
	IncludeActiveEmployeesinWaitingPeriod                 *bool                                        `xml:"urn:com.workday/bsvc Include_Active_Employees_in_Waiting_Period,omitempty"`
	RoundingRuleReference                                 *CompensationRoundingRuleObjectType          `xml:"urn:com.workday/bsvc Rounding_Rule_Reference,omitempty"`
	CompanyPerformanceScorecardReference                  *DefaultScorecardObjectType                  `xml:"urn:com.workday/bsvc Company_Performance_Scorecard_Reference,omitempty"`
	BonusModifierReference                                *DefaultScorecardObjectType                  `xml:"urn:com.workday/bsvc Bonus_Modifier_Reference,omitempty"`
	IncludeActiveEmployeesAssignedPlanDuringProcessPeriod *bool                                        `xml:"urn:com.workday/bsvc Include_Active_Employees_Assigned_Plan_During_Process_Period,omitempty"`
	PerformanceMatrixData                                 *PerformanceMatrixDataType                   `xml:"urn:com.workday/bsvc Performance_Matrix_Data,omitempty"`
	PerformanceFactorsData                                *PerformanceFactorsDataType                  `xml:"urn:com.workday/bsvc Performance_Factors_Data,omitempty"`
	CompensationTrancheReplacementData                    []CompensationTrancheReplacementDataType     `xml:"urn:com.workday/bsvc Compensation_Tranche_Replacement_Data,omitempty"`
	AmountData                                            *BonusPlanAmountDataType                     `xml:"urn:com.workday/bsvc Amount_Data"`
	PercentData                                           *BonusPlanPercentDataType                    `xml:"urn:com.workday/bsvc Percent_Data"`
	HideTarget                                            *bool                                        `xml:"urn:com.workday/bsvc Hide_Target,omitempty"`
	DeferredCompensation                                  *DeferredCompensationDataType                `xml:"urn:com.workday/bsvc Deferred_Compensation,omitempty"`
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

// Contains Bonus Plan Percent data.
type BonusPlanPercentDataType struct {
	Percentage                        float64                                          `xml:"urn:com.workday/bsvc Percentage,omitempty"`
	CompensationBasisReference        *CompensationBasisObjectType                     `xml:"urn:com.workday/bsvc Compensation_Basis_Reference"`
	PercentPlanProfileReplacementData []BonusPlanPercentPlanProfileReplacementDataType `xml:"urn:com.workday/bsvc Percent_Plan_Profile_Replacement_Data,omitempty"`
}

// Used to display or replace Percent Plan Profile Data for a Percent Bonus Plan.
type BonusPlanPercentPlanProfileReplacementDataType struct {
	EligibilityRuleReference   *ConditionRuleObjectType     `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	Percentage                 float64                      `xml:"urn:com.workday/bsvc Percentage,omitempty"`
	CompensationBasisReference *CompensationBasisObjectType `xml:"urn:com.workday/bsvc Compensation_Basis_Reference"`
}

// This is a wrapper for the Bonus Plan Amount; this is used to display or replace Amount Plan Profile Data; it is required and mutually exclusive with Bonus Plan Percentage.
type BonusPlanProfileReplacementAmountDataType struct {
	EligibilityRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	Amount                   float64                  `xml:"urn:com.workday/bsvc Amount,omitempty"`
	CurrencyReference        *CurrencyObjectType      `xml:"urn:com.workday/bsvc Currency_Reference"`
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

// Data for creation of a Calculated Plan
type CalculatedPlanDataType struct {
	CompensationElementReference  *CompensationPayEarningObjectType            `xml:"urn:com.workday/bsvc Compensation_Element_Reference,omitempty"`
	SetupSecuritySegmentReference []CompensationSetupSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Setup_Security_Segment_Reference,omitempty"`
	CalculationReference          *CalculationObjectType                       `xml:"urn:com.workday/bsvc Calculation_Reference,omitempty"`
	CurrencyReference             *CurrencyObjectType                          `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference            *FrequencyObjectType                         `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
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
type CalculationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CalculationObjectType struct {
	ID         []CalculationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Top element used for Check Position Budget as a sub business process.
type CheckPositionBudgetSubBusinessProcessType struct {
	BusinessSubProcessParameters *FinancialsBusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
}

// Commission Plan Data
type CommissionPlanDataType struct {
	CompensationElementReference         *CompensationPayEarningObjectType            `xml:"urn:com.workday/bsvc Compensation_Element_Reference"`
	SetupSecuritySegmentReference        []CompensationSetupSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Setup_Security_Segment_Reference,omitempty"`
	CurrencyReference                    *CurrencyObjectType                          `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference                   *FrequencyObjectType                         `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	TargetAmount                         float64                                      `xml:"urn:com.workday/bsvc Target_Amount,omitempty"`
	DrawAmount                           float64                                      `xml:"urn:com.workday/bsvc Draw_Amount,omitempty"`
	DrawFrequencyReference               *FrequencyObjectType                         `xml:"urn:com.workday/bsvc Draw_Frequency_Reference,omitempty"`
	DrawDuration                         string                                       `xml:"urn:com.workday/bsvc Draw_Duration,omitempty"`
	Recoverable                          *bool                                        `xml:"urn:com.workday/bsvc Recoverable,omitempty"`
	CommissionPlanProfileReplacementData []CommissionPlanProfileReplacementDataType   `xml:"urn:com.workday/bsvc Commission_Plan_Profile_Replacement_Data,omitempty"`
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

// This is a wrapper for the Commission Plan Profile Data; this is used to display or replace Commission Plan Profile Data.
type CommissionPlanProfileReplacementDataType struct {
	EligibilityRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	Amount                   float64                  `xml:"urn:com.workday/bsvc Amount,omitempty"`
	DrawAmount               float64                  `xml:"urn:com.workday/bsvc Draw_Amount,omitempty"`
	CurrencyReference        *CurrencyObjectType      `xml:"urn:com.workday/bsvc Currency_Reference"`
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

// Contains a unique identifier for an instance of an object.
type CompaRatioRangeSegmentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompaRatioRangeSegmentObjectType struct {
	ID         []CompaRatioRangeSegmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type CompensationAssignableComponentTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationAssignableComponentTypeObjectType struct {
	ID         []CompensationAssignableComponentTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompensationAssignablePlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationAssignablePlanObjectType struct {
	ID         []CompensationAssignablePlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompensationBasisConfigurableObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationBasisConfigurableObjectType struct {
	ID         []CompensationBasisConfigurableObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompensationBasisObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationBasisObjectType struct {
	ID         []CompensationBasisObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Benchmark Job Composite Data
type CompensationBenchmarkAmountReplacmentDataType struct {
	PercentileReference *CompensationBenchmarkPercentileObjectType `xml:"urn:com.workday/bsvc Percentile_Reference"`
	Amount              float64                                    `xml:"urn:com.workday/bsvc Amount"`
}

// Contains a unique identifier for an instance of an object.
type CompensationBenchmarkCompositeCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationBenchmarkCompositeCategoryObjectType struct {
	ID         []CompensationBenchmarkCompositeCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompensationBenchmarkDefaultObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationBenchmarkDefaultObjectType struct {
	ID         []CompensationBenchmarkDefaultObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompensationBenchmarkPercentileObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationBenchmarkPercentileObjectType struct {
	ID         []CompensationBenchmarkPercentileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Benchmark Job Data
type CompensationBenchmarkProfileDataType struct {
	CompensationEligibilityRuleReference *ConditionRuleObjectType        `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule_Reference"`
	CurrencyReference                    *CurrencyObjectType             `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference                   *FrequencyObjectType            `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	BenchmarkJobCompositeReplacementData []BenchmarkJobCompositeDataType `xml:"urn:com.workday/bsvc Benchmark_Job_Composite_Replacement_Data,omitempty"`
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

// The compensation duration severance matrix data element.
type CompensationDurationSeveranceMatrixDataType struct {
	CompensationBasisReference                   *CompensationBasisObjectType                       `xml:"urn:com.workday/bsvc Compensation_Basis_Reference"`
	CalculateCompensationBasisAsOfDateReference  *SeveranceServiceDateObjectType                    `xml:"urn:com.workday/bsvc Calculate_Compensation_Basis_As_Of_Date_Reference"`
	CurrencyReference                            *CurrencyObjectType                                `xml:"urn:com.workday/bsvc Currency_Reference"`
	FrequencyReference                           *FrequencyObjectType                               `xml:"urn:com.workday/bsvc Frequency_Reference"`
	CompensationDurationSeveranceMatrixEntryData []CompensationDurationSeveranceMatrixEntryDataType `xml:"urn:com.workday/bsvc Compensation_Duration_Severance_Matrix_Entry_Data"`
}

// The data for the matrix entry.
type CompensationDurationSeveranceMatrixEntryDataType struct {
	MinimumCompensationBasisRange float64                  `xml:"urn:com.workday/bsvc Minimum_Compensation_Basis_Range,omitempty"`
	MaximumCompensationBasisRange float64                  `xml:"urn:com.workday/bsvc Maximum_Compensation_Basis_Range,omitempty"`
	EligibilityRuleReference      *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	Duration                      float64                  `xml:"urn:com.workday/bsvc Duration,omitempty"`
	MinimumDuration               float64                  `xml:"urn:com.workday/bsvc Minimum_Duration,omitempty"`
	MaximumDuration               float64                  `xml:"urn:com.workday/bsvc Maximum_Duration,omitempty"`
}

// Data for this Compensation Eligibility Rule
type CompensationEligibilityRuleDataType struct {
	EffectiveDate                     *time.Time                           `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	ConditionRuleData                 *ConditionRuleDataWWSType            `xml:"urn:com.workday/bsvc Condition_Rule_Data"`
	CompensationPackageReference      []CompensationPackageObjectType      `xml:"urn:com.workday/bsvc Compensation_Package_Reference,omitempty"`
	CompensationGradeReference        []CompensationGradeObjectType        `xml:"urn:com.workday/bsvc Compensation_Grade_Reference,omitempty"`
	CompensationGradeProfileReference []CompensationGradeProfileObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Profile_Reference,omitempty"`
	CompensationPlanReference         []CompensationPlanObjectType         `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
}

func (t *CompensationEligibilityRuleDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CompensationEligibilityRuleDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *CompensationEligibilityRuleDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompensationEligibilityRuleDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Request Criteria for this request. Note that there are none available.
type CompensationEligibilityRuleRequestCriteriaType struct {
}

// Compensation Eligibility Rule References for rules to return
type CompensationEligibilityRuleRequestReferencesType struct {
	CompensationEligibilityRuleReference []ConditionRuleObjectType `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule_Reference"`
}

// Compensation Eligibility Rule Response Data
type CompensationEligibilityRuleResponseDataType struct {
	CompensationEligibilityRule []CompensationEligibilityRuleType `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule,omitempty"`
}

// Provides the filters for the web service operation.
type CompensationEligibilityRuleResponseGroupType struct {
	HideConfigurationDependencies *bool `xml:"urn:com.workday/bsvc Hide_Configuration_Dependencies,omitempty"`
}

// Response data element for the Compensation Eligibility Rule
type CompensationEligibilityRuleType struct {
	CompensationEligibilityRuleReference *ConditionRuleObjectType              `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule_Reference,omitempty"`
	CompensationEligibilityRuleData      []CompensationEligibilityRuleDataType `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule_Data,omitempty"`
}

// Contains compensation grade data.
type CompensationGradeDataType struct {
	CompensationGradeID                       string                                       `xml:"urn:com.workday/bsvc Compensation_Grade_ID,omitempty"`
	EffectiveDate                             *time.Time                                   `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	Name                                      string                                       `xml:"urn:com.workday/bsvc Name"`
	Description                               string                                       `xml:"urn:com.workday/bsvc Description,omitempty"`
	CompensationElementReference              []CompensationPayEarningObjectType           `xml:"urn:com.workday/bsvc Compensation_Element_Reference"`
	EligibilityRuleReference                  []ConditionRuleObjectType                    `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference,omitempty"`
	CompensationPayRange                      *CompensationPayRangeDataType                `xml:"urn:com.workday/bsvc Compensation_Pay_Range,omitempty"`
	CompensationGradeProfile                  []CompensationGradeProfileType               `xml:"urn:com.workday/bsvc Compensation_Grade_Profile,omitempty"`
	SetupSecuritySegmentReference             []CompensationSetupSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Setup_Security_Segment_Reference,omitempty"`
	AssignFirstStepDuringCompensationProposal *bool                                        `xml:"urn:com.workday/bsvc Assign_First_Step_During_Compensation_Proposal,omitempty"`
	PayLevelReference                         *CompplObjectType                            `xml:"urn:com.workday/bsvc Pay_Level_Reference,omitempty"`
	Inactive                                  *bool                                        `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	AlternativePayRange                       []AlternativePayRangeType                    `xml:"urn:com.workday/bsvc Alternative_Pay_Range,omitempty"`
}

func (t *CompensationGradeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CompensationGradeDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *CompensationGradeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompensationGradeDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

type CompensationGradeHierarchyDataType struct {
	EffectiveDate          time.Time                    `xml:"urn:com.workday/bsvc Effective_Date"`
	CompensationGradeLevel []CompensationGradeLevelType `xml:"urn:com.workday/bsvc Compensation_Grade_Level,omitempty"`
}

func (t *CompensationGradeHierarchyDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CompensationGradeHierarchyDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *CompensationGradeHierarchyDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompensationGradeHierarchyDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Utilize the compensation grade hierarchy element to retrieve specific instance(s) of compensation grade hierarchy data.
type CompensationGradeHierarchyRequestReferencesType struct {
	EffectiveDate time.Time `xml:"urn:com.workday/bsvc Effective_Date"`
}

func (t *CompensationGradeHierarchyRequestReferencesType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CompensationGradeHierarchyRequestReferencesType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *CompensationGradeHierarchyRequestReferencesType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompensationGradeHierarchyRequestReferencesType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

type CompensationGradeHierarchyResponseDataType struct {
	CompensationGradeHierarchy []CompensationGradeHierarchyType `xml:"urn:com.workday/bsvc Compensation_Grade_Hierarchy,omitempty"`
}

type CompensationGradeHierarchyType struct {
	CompensationGradeHierarchyData *CompensationGradeHierarchyDataType `xml:"urn:com.workday/bsvc Compensation_Grade_Hierarchy_Data,omitempty"`
}

type CompensationGradeLevelDataType struct {
	ID                         string                        `xml:"urn:com.workday/bsvc ID,omitempty"`
	Rank                       float64                       `xml:"urn:com.workday/bsvc Rank"`
	Name                       string                        `xml:"urn:com.workday/bsvc Name"`
	CompensationGradeReference []CompensationGradeObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Reference"`
}

// Contains a unique identifier for an instance of an object.
type CompensationGradeLevelObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationGradeLevelObjectType struct {
	ID         []CompensationGradeLevelObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type CompensationGradeLevelType struct {
	CompensationGradeLevelReference *CompensationGradeLevelObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Level_Reference,omitempty"`
	CompensationGradeLevelData      *CompensationGradeLevelDataType   `xml:"urn:com.workday/bsvc Compensation_Grade_Level_Data,omitempty"`
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

// Contains compensation grade profile data.
type CompensationGradeProfileDataType struct {
	CompensationGradeProfileID    string                                       `xml:"urn:com.workday/bsvc Compensation_Grade_Profile_ID,omitempty"`
	EffectiveDate                 *time.Time                                   `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	Name                          string                                       `xml:"urn:com.workday/bsvc Name"`
	Description                   string                                       `xml:"urn:com.workday/bsvc Description,omitempty"`
	CompensationElementReference  []CompensationPayEarningObjectType           `xml:"urn:com.workday/bsvc Compensation_Element_Reference"`
	EligibilityRuleReference      []ConditionRuleObjectType                    `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference,omitempty"`
	Inactive                      *bool                                        `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	CompensationPayRangeData      *CompensationPayRangeDataType                `xml:"urn:com.workday/bsvc Compensation_Pay_Range_Data,omitempty"`
	SetupSecuritySegmentReference []CompensationSetupSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Setup_Security_Segment_Reference,omitempty"`
	AlternativePayRange           []AlternativePayRangeType                    `xml:"urn:com.workday/bsvc Alternative_Pay_Range,omitempty"`
}

func (t *CompensationGradeProfileDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CompensationGradeProfileDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *CompensationGradeProfileDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompensationGradeProfileDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
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

// Contains a compensation grade profile and its associated data.
type CompensationGradeProfileType struct {
	CompensationGradeProfileReference *CompensationGradeProfileObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Profile_Reference,omitempty"`
	CompensationGradeProfileData      *CompensationGradeProfileDataType   `xml:"urn:com.workday/bsvc Compensation_Grade_Profile_Data,omitempty"`
	Delete                            bool                                `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Get Compensation Grades Request Criteria
type CompensationGradeRequestCriteriaType struct {
	IncludeInactive *bool `xml:"urn:com.workday/bsvc Include_Inactive,omitempty"`
}

// Utilize the compensation grade element to retrieve specific instance(s) of compensation grade and its associated data.
type CompensationGradeRequestReferencesType struct {
	CompensationGradeReference []CompensationGradeObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Reference"`
}

// Contains compensation grade information based on Request References or Request Criteria.
type CompensationGradeResponseDataType struct {
	CompensationGrade []CompensationGradeType `xml:"urn:com.workday/bsvc Compensation_Grade,omitempty"`
}

// Provides the filters for the web service operation.
type CompensationGradeResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Contains a compensation grade and its associated data.
type CompensationGradeType struct {
	CompensationGradeReference *CompensationGradeObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Reference,omitempty"`
	CompensationGradeData      *CompensationGradeDataType   `xml:"urn:com.workday/bsvc Compensation_Grade_Data,omitempty"`
}

// The compensation length of service severance matrix data element.
type CompensationLengthofServiceSeveranceMatrixDataType struct {
	CompensationBasisReference                          *CompensationBasisObjectType                              `xml:"urn:com.workday/bsvc Compensation_Basis_Reference"`
	CalculateCompensationBasisAsOfDateReference         *SeveranceServiceDateObjectType                           `xml:"urn:com.workday/bsvc Calculate_Compensation_Basis_As_Of_Date_Reference"`
	CurrencyReference                                   *CurrencyObjectType                                       `xml:"urn:com.workday/bsvc Currency_Reference"`
	FrequencyReference                                  *FrequencyObjectType                                      `xml:"urn:com.workday/bsvc Frequency_Reference"`
	CompensationRoundingRuleReference                   *CompensationRoundingRuleObjectType                       `xml:"urn:com.workday/bsvc Compensation_Rounding_Rule_Reference"`
	MultiplierOrderReference                            *BenefitMultiplierOrderObjectType                         `xml:"urn:com.workday/bsvc Multiplier_Order_Reference"`
	CompensationLengthofServiceSeveranceMatrixEntryData []CompensationLengthofServiceSeveranceMatrixEntryDataType `xml:"urn:com.workday/bsvc Compensation_Length_of_Service_Severance_Matrix_Entry_Data"`
}

// The data for the matrix entry.
type CompensationLengthofServiceSeveranceMatrixEntryDataType struct {
	MinimumCompensationBasisRange float64                  `xml:"urn:com.workday/bsvc Minimum_Compensation_Basis_Range,omitempty"`
	MaximumCompensationBasisRange float64                  `xml:"urn:com.workday/bsvc Maximum_Compensation_Basis_Range,omitempty"`
	EligibilityRuleReference      *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	LengthofServiceMultiplier     float64                  `xml:"urn:com.workday/bsvc Length_of_Service_Multiplier,omitempty"`
	MinimumDuration               float64                  `xml:"urn:com.workday/bsvc Minimum_Duration,omitempty"`
	MaximumDuration               float64                  `xml:"urn:com.workday/bsvc Maximum_Duration,omitempty"`
}

// The amount based data for this compensation matrix.
type CompensationMatrixAmountBasedDataType struct {
	DefaultFrequencyReference              *FrequencyObjectType                         `xml:"urn:com.workday/bsvc Default_Frequency_Reference"`
	DefaultCurrencyReference               *CurrencyObjectType                          `xml:"urn:com.workday/bsvc Default_Currency_Reference,omitempty"`
	CompensationMatrixTargetRulesReference []ConditionRuleObjectType                    `xml:"urn:com.workday/bsvc Compensation_Matrix_Target_Rules_Reference,omitempty"`
	CompensationMatrixEntryAmountBasedData []CompensationMatrixEntryAmountBasedDataType `xml:"urn:com.workday/bsvc Compensation_Matrix_Entry_Amount_Based_Data,omitempty"`
}

// All data for this compensation matrix.
type CompensationMatrixDataType struct {
	CompensationMatrixMatrixID                    string                                              `xml:"urn:com.workday/bsvc Compensation_Matrix_Matrix_ID,omitempty"`
	EffectiveDate                                 *time.Time                                          `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	CompensationMatrixName                        string                                              `xml:"urn:com.workday/bsvc Compensation_Matrix_Name"`
	CompensationMatrixDescription                 string                                              `xml:"urn:com.workday/bsvc Compensation_Matrix_Description,omitempty"`
	RatingScaleReference                          *ReviewRatingScaleObjectType                        `xml:"urn:com.workday/bsvc Rating_Scale_Reference,omitempty"`
	CompensationMatrixNonweightedPercentBasedData []CompensationMatrixNonweightedPercentBasedDataType `xml:"urn:com.workday/bsvc Compensation_Matrix_Non-weighted_Percent_Based_Data"`
	CompensationMatrixAmountBasedData             []CompensationMatrixAmountBasedDataType             `xml:"urn:com.workday/bsvc Compensation_Matrix_Amount_Based_Data"`
	CompensationMatrixforWeightedPercentBasedData []CompensationMatrixWeightedPercentBasedDataType    `xml:"urn:com.workday/bsvc Compensation_Matrix_for_Weighted_Percent_Based_Data"`
}

func (t *CompensationMatrixDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CompensationMatrixDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *CompensationMatrixDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompensationMatrixDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Data for a single compensation matrix entry.
type CompensationMatrixEntryAmountBasedDataType struct {
	CompensationEntryTargetRuleReference *ConditionRuleObjectType          `xml:"urn:com.workday/bsvc Compensation_Entry_Target_Rule_Reference,omitempty"`
	PerformanceEvaluationRatingReference *ReviewRatingObjectType           `xml:"urn:com.workday/bsvc Performance_Evaluation_Rating_Reference,omitempty"`
	QuartilePlacementReference           *PayRangeQuartileObjectType       `xml:"urn:com.workday/bsvc Quartile_Placement_Reference,omitempty"`
	RetentionReference                   *RetentionObjectType              `xml:"urn:com.workday/bsvc Retention_Reference,omitempty"`
	CompaRatioRangeSegmentReference      *CompaRatioRangeSegmentObjectType `xml:"urn:com.workday/bsvc Compa_Ratio_Range_Segment_Reference,omitempty"`
	PayRangeSegmentReference             *PayRangeSegmentTypeObjectType    `xml:"urn:com.workday/bsvc Pay_Range_Segment_Reference,omitempty"`
	PotentialReference                   *PotentialObjectType              `xml:"urn:com.workday/bsvc Potential_Reference,omitempty"`
	MinimumAmount                        float64                           `xml:"urn:com.workday/bsvc Minimum_Amount,omitempty"`
	MaximumAmount                        float64                           `xml:"urn:com.workday/bsvc Maximum_Amount,omitempty"`
	CurrencyReference                    *CurrencyObjectType               `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
}

// Data for a single entry in a percent based compensation matrix.
type CompensationMatrixEntryPercentBasedDataType struct {
	CompensationMatrixEntryEligibilityRuleReference *ConditionRuleObjectType          `xml:"urn:com.workday/bsvc Compensation_Matrix_Entry_Eligibility_Rule_Reference,omitempty"`
	PerformanceEvaluationRatingReference            *ReviewRatingObjectType           `xml:"urn:com.workday/bsvc Performance_Evaluation_Rating_Reference,omitempty"`
	QuartilePlacementReference                      *PayRangeQuartileObjectType       `xml:"urn:com.workday/bsvc Quartile_Placement_Reference,omitempty"`
	RetentionReference                              *RetentionObjectType              `xml:"urn:com.workday/bsvc Retention_Reference,omitempty"`
	CompaRatioRangeSegmentReference                 *CompaRatioRangeSegmentObjectType `xml:"urn:com.workday/bsvc Compa_Ratio_Range_Segment_Reference,omitempty"`
	PayRangeSegmentReference                        *PayRangeSegmentTypeObjectType    `xml:"urn:com.workday/bsvc Pay_Range_Segment_Reference,omitempty"`
	PotentialReference                              *PotentialObjectType              `xml:"urn:com.workday/bsvc Potential_Reference,omitempty"`
	MinimumPercent                                  float64                           `xml:"urn:com.workday/bsvc Minimum_Percent,omitempty"`
	MaximumPercent                                  float64                           `xml:"urn:com.workday/bsvc Maximum_Percent,omitempty"`
}

// Data for a non-weighted percent based compensation matrix.
type CompensationMatrixNonweightedPercentBasedDataType struct {
	CompensationMatrixTargetRulesReference  []ConditionRuleObjectType                     `xml:"urn:com.workday/bsvc Compensation_Matrix_Target_Rules_Reference,omitempty"`
	CompensationMatrixEntryPercentBasedData []CompensationMatrixEntryPercentBasedDataType `xml:"urn:com.workday/bsvc Compensation_Matrix_Entry_Percent_Based_Data,omitempty"`
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

// Request criteria
type CompensationMatrixRequestCriteriaType struct {
	Name                                      string `xml:"urn:com.workday/bsvc Name"`
	CompensationMatrixAmountBased             bool   `xml:"urn:com.workday/bsvc Compensation_Matrix_Amount_Based"`
	CompensationMatrixNonweightedPercentBased bool   `xml:"urn:com.workday/bsvc Compensation_Matrix_Non-weighted_Percent_Based"`
	CompensationMatrixWeightedPercentBased    bool   `xml:"urn:com.workday/bsvc Compensation_Matrix_Weighted_Percent_Based"`
}

// Request references element.
type CompensationMatrixRequestReferencesType struct {
	CompensationMatricesReference []CompensationMatrixObjectType `xml:"urn:com.workday/bsvc Compensation_Matrices_Reference,omitempty"`
}

// Response Data element.
type CompensationMatrixResponseDataType struct {
	CompensationMatrix []CompensationMatrixType `xml:"urn:com.workday/bsvc Compensation_Matrix,omitempty"`
}

// The response group.
type CompensationMatrixResponseGroupType struct {
	IncludeReference              *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeCompensationMatrixData *bool `xml:"urn:com.workday/bsvc Include_Compensation_Matrix_Data,omitempty"`
}

// Compensation Matrix element.
type CompensationMatrixType struct {
	CompensationMatrixReference *CompensationMatrixObjectType `xml:"urn:com.workday/bsvc Compensation_Matrix_Reference,omitempty"`
	CompensationMatrixData      []CompensationMatrixDataType  `xml:"urn:com.workday/bsvc Compensation_Matrix_Data,omitempty"`
}

// Data for this weight percent based compensation matrix.
type CompensationMatrixWeightedPercentBasedDataType struct {
	CompensationMatrixTargetRulesReference  []ConditionRuleObjectType                     `xml:"urn:com.workday/bsvc Compensation_Matrix_Target_Rules_Reference,omitempty"`
	CompensationMatrixEntryPercentBasedData []CompensationMatrixEntryPercentBasedDataType `xml:"urn:com.workday/bsvc Compensation_Matrix_Entry_Percent_Based_Data,omitempty"`
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

// The compensation pay range.
type CompensationPayRangeDataType struct {
	NumberofSegments    float64                  `xml:"urn:com.workday/bsvc Number_of_Segments,omitempty"`
	Minimum             float64                  `xml:"urn:com.workday/bsvc Minimum,omitempty"`
	Midpoint            float64                  `xml:"urn:com.workday/bsvc Midpoint,omitempty"`
	Maximum             float64                  `xml:"urn:com.workday/bsvc Maximum,omitempty"`
	Spread              float64                  `xml:"urn:com.workday/bsvc Spread,omitempty"`
	Segment1Top         float64                  `xml:"urn:com.workday/bsvc Segment_1_Top,omitempty"`
	Segment2Top         float64                  `xml:"urn:com.workday/bsvc Segment_2_Top,omitempty"`
	Segment3Top         float64                  `xml:"urn:com.workday/bsvc Segment_3_Top,omitempty"`
	Segment4Top         float64                  `xml:"urn:com.workday/bsvc Segment_4_Top,omitempty"`
	Segment5Top         float64                  `xml:"urn:com.workday/bsvc Segment_5_Top,omitempty"`
	CurrencyReference   *CurrencyObjectType      `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference  *FrequencyObjectType     `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	SalaryPlanReference *SalaryPayPlanObjectType `xml:"urn:com.workday/bsvc Salary_Plan_Reference,omitempty"`
	AllowOverride       *bool                    `xml:"urn:com.workday/bsvc Allow_Override,omitempty"`
	CompensationStep    []CompensationStepType   `xml:"urn:com.workday/bsvc Compensation_Step,omitempty"`
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

// Contains common information for the following Compensation Plans:  Future Payment, Bonus, Merit, Allowance, Commission, and Base Pay.  It also must contain a Future Play, Bonus, Merit, Allowance, Commission, or Bay Pay Plan data.
type CompensationPlanDataType struct {
	CompensationPlanID       string                     `xml:"urn:com.workday/bsvc Compensation_Plan_ID,omitempty"`
	EffectiveDate            *time.Time                 `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	Name                     string                     `xml:"urn:com.workday/bsvc Name"`
	Description              string                     `xml:"urn:com.workday/bsvc Description,omitempty"`
	EligibilityRuleReference []ConditionRuleObjectType  `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference,omitempty"`
	Inactive                 *bool                      `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	FuturePaymentPlanData    *FuturePaymentPlanDataType `xml:"urn:com.workday/bsvc Future_Payment_Plan_Data"`
	BonusPlanData            *BonusPlanDataType         `xml:"urn:com.workday/bsvc Bonus_Plan_Data"`
	MeritPlanData            *MeritPlanDataType         `xml:"urn:com.workday/bsvc Merit_Plan_Data"`
	StockPlanData            *StockPlanDataType         `xml:"urn:com.workday/bsvc Stock_Plan_Data"`
	AllowancePlanData        *AllowancePlanDataType     `xml:"urn:com.workday/bsvc Allowance_Plan_Data"`
	CommissionPlanData       *CommissionPlanDataType    `xml:"urn:com.workday/bsvc Commission_Plan_Data"`
	BasePayPlanData          *BasePayPlanDataType       `xml:"urn:com.workday/bsvc Base_Pay_Plan_Data"`
	PeriodPlanData           *PeriodPlanDataType        `xml:"urn:com.workday/bsvc Period_Plan_Data"`
	CalculatedPlanData       *CalculatedPlanDataType    `xml:"urn:com.workday/bsvc Calculated_Plan_Data"`
}

func (t *CompensationPlanDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CompensationPlanDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *CompensationPlanDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompensationPlanDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
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

// Values in this element can be used to filter the types of Compensation Plans returned
type CompensationPlanRequestCriteriaType struct {
	PlanTypeReference []CompensationAssignableComponentTypeObjectType `xml:"urn:com.workday/bsvc Plan_Type_Reference,omitempty"`
	IncludeInactive   *bool                                           `xml:"urn:com.workday/bsvc Include_Inactive,omitempty"`
}

// Use this element to request specific Compensation Plans given the reference ID values
type CompensationPlanRequestReferencesType struct {
	CompensationPlanReference []CompensationAssignablePlanObjectType `xml:"urn:com.workday/bsvc Compensation_Plan_Reference"`
}

// Response element containing instances of compensation plans and the associated data.
type CompensationPlanResponseDataType struct {
	CompensationPlan []CompensationPlanType `xml:"urn:com.workday/bsvc Compensation_Plan"`
}

// Use to limit the returned data for a given Compensation Plan
type CompensationPlanResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// An instance of a compensation plan and associated data.
type CompensationPlanType struct {
	CompensationPlanReference *CompensationAssignablePlanObjectType `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
	CompensationPlanData      *CompensationPlanDataType             `xml:"urn:com.workday/bsvc Compensation_Plan_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompensationPreviousSystemHistoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationPreviousSystemHistoryObjectType struct {
	ID         []CompensationPreviousSystemHistoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompensationRoundingRuleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationRoundingRuleObjectType struct {
	ID         []CompensationRoundingRuleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains scorecard data.
type CompensationScorecardDataType struct {
	ID                   string                        `xml:"urn:com.workday/bsvc ID,omitempty"`
	EffectiveDate        time.Time                     `xml:"urn:com.workday/bsvc Effective_Date"`
	ScorecardName        string                        `xml:"urn:com.workday/bsvc Scorecard_Name"`
	ScorecardDescription string                        `xml:"urn:com.workday/bsvc Scorecard_Description,omitempty"`
	ScorecardGoalsData   []PerformanceCriteriaDataType `xml:"urn:com.workday/bsvc Scorecard_Goals_Data"`
	ScorecardProfileData []ScorecardProfileDataType    `xml:"urn:com.workday/bsvc Scorecard_Profile_Data,omitempty"`
}

func (t *CompensationScorecardDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CompensationScorecardDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *CompensationScorecardDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompensationScorecardDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains additional request criteria to limit the information returned. Currently no additional criteria is supported.
type CompensationScorecardRequestCriteriaType struct {
}

// A container for one or more unique scorecard identifiers.
type CompensationScorecardRequestReferencesType struct {
	CompensationScorecardReference []DefaultScorecardObjectType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Reference"`
}

// Contains the scorecard response information. This includes the unique scorecard identifier and the detailed scorecard data which was added or updated.
type CompensationScorecardResponseDataType struct {
	CompensationScorecard []CompensationScorecardType `xml:"urn:com.workday/bsvc Compensation_Scorecard,omitempty"`
}

// Contains additional response criteria to enhance the information returned.
type CompensationScorecardResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Contains the detailed scorecard result information.
type CompensationScorecardResultDataType struct {
	ID                             string                                       `xml:"urn:com.workday/bsvc ID,omitempty"`
	EvaluationDate                 time.Time                                    `xml:"urn:com.workday/bsvc Evaluation_Date"`
	CompensationScorecardReference *DefaultScorecardObjectType                  `xml:"urn:com.workday/bsvc Compensation_Scorecard_Reference"`
	ScorecardResultData            []ScorecardResultDataType                    `xml:"urn:com.workday/bsvc Scorecard_Result_Data"`
	ProfileScorecardResultData     []ProfileCompensationScorecardResultDataType `xml:"urn:com.workday/bsvc Profile_Scorecard_Result_Data,omitempty"`
}

func (t *CompensationScorecardResultDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CompensationScorecardResultDataType
	var layout struct {
		*T
		EvaluationDate *xsdDate `xml:"urn:com.workday/bsvc Evaluation_Date"`
	}
	layout.T = (*T)(t)
	layout.EvaluationDate = (*xsdDate)(&layout.T.EvaluationDate)
	return e.EncodeElement(layout, start)
}
func (t *CompensationScorecardResultDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompensationScorecardResultDataType
	var overlay struct {
		*T
		EvaluationDate *xsdDate `xml:"urn:com.workday/bsvc Evaluation_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EvaluationDate = (*xsdDate)(&overlay.T.EvaluationDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains additional request criteria to limit the information returned.  Currently no additional criteria is supported.
type CompensationScorecardResultRequestCriteriaType struct {
}

// A container for one or more unique scorecard result identifiers.
type CompensationScorecardResultRequestReferencesType struct {
	CompensationScorecardResultReference []ScoresetContainerObjectType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Result_Reference"`
}

// Contains the scorecard result response information. This includes the unique scorecard result identifier and the detailed scorecard result data which was added or updated.
type CompensationScorecardResultResponseDataType struct {
	CompensationScorecardResult []CompensationScorecardResultType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Result,omitempty"`
}

// Contains additional response criteria to enhance the information returned.
type CompensationScorecardResultResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Contains the scorecard result response information. This includes the unique scorecard result identifier and the detailed scorecard result data which was added or updated.
type CompensationScorecardResultType struct {
	CompensationScorecardResultReference *ScoresetContainerObjectType          `xml:"urn:com.workday/bsvc Compensation_Scorecard_Result_Reference,omitempty"`
	CompensationScorecardResultData      []CompensationScorecardResultDataType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Result_Data,omitempty"`
}

// Contains scorecard data.
type CompensationScorecardType struct {
	CompensationScorecardReference *DefaultScorecardObjectType     `xml:"urn:com.workday/bsvc Compensation_Scorecard_Reference,omitempty"`
	CompensationScorecardData      []CompensationScorecardDataType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompensationSetupSecuritySegmentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationSetupSecuritySegmentObjectType struct {
	ID         []CompensationSetupSecuritySegmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// A compensation step
type CompensationStepDataType struct {
	CompensationStepID       string                       `xml:"urn:com.workday/bsvc Compensation_Step_ID,omitempty"`
	Sequence                 string                       `xml:"urn:com.workday/bsvc Sequence"`
	Name                     string                       `xml:"urn:com.workday/bsvc Name"`
	Amount                   float64                      `xml:"urn:com.workday/bsvc Amount"`
	Interval                 float64                      `xml:"urn:com.workday/bsvc Interval,omitempty"`
	PeriodReference          *FrequencyBehaviorObjectType `xml:"urn:com.workday/bsvc Period_Reference,omitempty"`
	ProgressionRuleReference *ConditionRuleObjectType     `xml:"urn:com.workday/bsvc Progression_Rule_Reference,omitempty"`
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

// Contains the pay range data.
type CompensationStepType struct {
	CompensationStepReference *CompensationStepObjectType `xml:"urn:com.workday/bsvc Compensation_Step_Reference,omitempty"`
	CompensationStepData      *CompensationStepDataType   `xml:"urn:com.workday/bsvc Compensation_Step_Data,omitempty"`
	Delete                    bool                        `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// This is a wrapper for the Compensation Tranche Replacement Data which is used to display and replace Compensation Tranche Data within Bonus Plan; it is optional.
type CompensationTrancheReplacementDataType struct {
	Name       float64 `xml:"urn:com.workday/bsvc Name"`
	Percentage float64 `xml:"urn:com.workday/bsvc Percentage"`
}

// Contains a unique identifier for an instance of an object.
type CompensationWeightedPercentMatrixObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompensationWeightedPercentMatrixObjectType struct {
	ID         []CompensationWeightedPercentMatrixObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompplObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompplObjectType struct {
	ID         []CompplObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ConditionRuleObjectType struct {
	ID         []ConditionRuleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Launch Employee Severance Worksheet Request
type CreateSeveranceWorksheetRequestType struct {
	BusinessProcessParameters      *BusinessProcessParametersType      `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EmployeeSeveranceWorksheetData *EmployeeSeveranceWorksheetDataType `xml:"urn:com.workday/bsvc Employee_Severance_Worksheet_Data"`
	Version                        string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Launch Employee Severance Worksheet Response
type CreateSeveranceWorksheetResponseType struct {
	EmployeeSeveranceWorksheetEventReference *EmployeeSeveranceWorksheetEventObjectType `xml:"urn:com.workday/bsvc Employee_Severance_Worksheet_Event_Reference,omitempty"`
	Version                                  string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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
type DefaultScorecardObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DefaultScorecardObjectType struct {
	ID         []DefaultScorecardObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DeferredBonusCalculationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DeferredBonusCalculationObjectType struct {
	ID         []DeferredBonusCalculationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Deferred Compensation
type DeferredCompensationDataType struct {
	DeferredCompensationEligibilityReference []ConditionRuleObjectType              `xml:"urn:com.workday/bsvc Deferred_Compensation_Eligibility_Reference"`
	DefaultCalculationReference              *DeferredBonusCalculationObjectType    `xml:"urn:com.workday/bsvc Default_Calculation_Reference"`
	DefaultStockPlanReference                *StockPercentPlanObjectType            `xml:"urn:com.workday/bsvc Default_Stock_Plan_Reference"`
	DeferredCompensationProfiles             []DeferredCompensationProfilesDataType `xml:"urn:com.workday/bsvc Deferred_Compensation_Profiles,omitempty"`
}

// Deferred Compensation Profiles
type DeferredCompensationProfilesDataType struct {
	EligibilityRuleReference []ConditionRuleObjectType           `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	CalculationReference     *DeferredBonusCalculationObjectType `xml:"urn:com.workday/bsvc Calculation_Reference"`
	StockPlanReference       *StockPercentPlanObjectType         `xml:"urn:com.workday/bsvc Stock_Plan_Reference"`
}

// Contains a unique identifier for an instance of an object.
type EligibilityWaitingPeriodObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EligibilityWaitingPeriodObjectType struct {
	ID         []EligibilityWaitingPeriodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Utilize the Request Criteria element to search for specific instance(s) of Eligible Earnings.
type EligibleEarningRequestCriteriaType struct {
	WorkerReference                 *WorkerObjectType                         `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	PositionReference               *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	EligibleEarningsPeriodReference *EligibleEarningsOverridePeriodObjectType `xml:"urn:com.workday/bsvc Eligible_Earnings_Period_Reference,omitempty"`
	BonusPlanReference              *BonusPercentPlanObjectType               `xml:"urn:com.workday/bsvc Bonus_Plan_Reference,omitempty"`
}

// Data element for the Put Eligible Earnings.
type EligibleEarningsDataType struct {
	EligibleEarningsID            string                                    `xml:"urn:com.workday/bsvc Eligible_Earnings_ID,omitempty"`
	EmployeeReference             *WorkerObjectType                         `xml:"urn:com.workday/bsvc Employee_Reference"`
	PositionReference             *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	PeriodReference               *EligibleEarningsOverridePeriodObjectType `xml:"urn:com.workday/bsvc Period_Reference"`
	ApplytoAllBonusPlans          bool                                      `xml:"urn:com.workday/bsvc Apply_to_All_Bonus_Plans"`
	RestricttoBonusPlansReference []BonusPercentPlanObjectType              `xml:"urn:com.workday/bsvc Restrict_to_Bonus_Plans_Reference"`
	Amount                        float64                                   `xml:"urn:com.workday/bsvc Amount,omitempty"`
	CurrencyReference             *CurrencyObjectType                       `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
}

// Data element for the Put Eligible Earnings.
type EligibleEarningsOverrideHVDataType struct {
	EligibleEarningsID            string                       `xml:"urn:com.workday/bsvc Eligible_Earnings_ID,omitempty"`
	EmployeeReference             *WorkerObjectType            `xml:"urn:com.workday/bsvc Employee_Reference"`
	PositionReference             *PositionElementObjectType   `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	ApplytoAllBonusPlans          bool                         `xml:"urn:com.workday/bsvc Apply_to_All_Bonus_Plans"`
	RestricttoBonusPlansReference []BonusPercentPlanObjectType `xml:"urn:com.workday/bsvc Restrict_to_Bonus_Plans_Reference"`
	Amount                        float64                      `xml:"urn:com.workday/bsvc Amount,omitempty"`
	CurrencyReference             *CurrencyObjectType          `xml:"urn:com.workday/bsvc Currency_Reference"`
}

// Contains a unique identifier for an instance of an object.
type EligibleEarningsOverrideObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EligibleEarningsOverrideObjectType struct {
	ID         []EligibleEarningsOverrideObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EligibleEarningsOverridePeriodObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EligibleEarningsOverridePeriodObjectType struct {
	ID         []EligibleEarningsOverridePeriodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Utilize the Request References element to retrieve specific instance(s) of Eligible Earnings and its associated data.
type EligibleEarningsRequestReferencesType struct {
	EligibleEarningsReference []EligibleEarningsOverrideObjectType `xml:"urn:com.workday/bsvc Eligible_Earnings_Reference"`
}

// Contains Eligible Earnings Override information based on Request References or Request Criteria.
type EligibleEarningsResponseDataType struct {
	EligibleEarnings []EligibleEarningsType `xml:"urn:com.workday/bsvc Eligible_Earnings,omitempty"`
}

// The response group allows for the response data to be tailored to only included elements that the user is looking for.  If no response group is provided in the request the following elements will be returned:  Reference, Eligible Earnings Data
type EligibleEarningsResponseGroupType struct {
	IncludeReference            *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeEligibleEarningsData *bool `xml:"urn:com.workday/bsvc Include_Eligible_Earnings_Data,omitempty"`
}

// Contains Eligible Earnings Override information based on Request References or Request Criteria.
type EligibleEarningsType struct {
	EligibleEarningsReference *EligibleEarningsOverrideObjectType `xml:"urn:com.workday/bsvc Eligible_Earnings_Reference,omitempty"`
	EligibleEarningsData      *EligibleEarningsDataType           `xml:"urn:com.workday/bsvc Eligible_Earnings_Data,omitempty"`
}

// Data element for the Request Employee Merit Adjustment business process.
type EmployeeMeritAdjustmentDataType struct {
	BasePayPlanReference *SalaryPayPlanObjectType `xml:"urn:com.workday/bsvc Base_Pay_Plan_Reference,omitempty"`
	PercentChange        float64                  `xml:"urn:com.workday/bsvc Percent_Change"`
	AmountChange         float64                  `xml:"urn:com.workday/bsvc Amount_Change"`
	NewAmount            float64                  `xml:"urn:com.workday/bsvc New_Amount"`
	CurrencyReference    *CurrencyObjectType      `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference   *FrequencyObjectType     `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	CompensationComment  string                   `xml:"urn:com.workday/bsvc Compensation_Comment,omitempty"`
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

// Subelement for entering a severance component type along with eligibility and a comment
type EmployeeSeverancePackagePayComponentDataType struct {
	EmployeeSeverancePackagePayComponentReference *SeverancePackageComponentTypeObjectType `xml:"urn:com.workday/bsvc Employee_Severance_Package_Pay_Component_Reference,omitempty"`
	AssignmentDuration                            float64                                  `xml:"urn:com.workday/bsvc Assignment_Duration,omitempty"`
	PayRate                                       float64                                  `xml:"urn:com.workday/bsvc Pay_Rate,omitempty"`
	Total                                         float64                                  `xml:"urn:com.workday/bsvc Total,omitempty"`
	PlanReference                                 *SalaryPlanObjectType                    `xml:"urn:com.workday/bsvc Plan_Reference,omitempty"`
	Comments                                      string                                   `xml:"urn:com.workday/bsvc Comments,omitempty"`
}

// Employee Severance Worksheet Data
type EmployeeSeveranceWorksheetDataType struct {
	EffectiveDate                 time.Time                                 `xml:"urn:com.workday/bsvc Effective_Date"`
	SeverancePackageReference     *SeverancePackageObjectType               `xml:"urn:com.workday/bsvc Severance_Package_Reference"`
	EmployeeReference             *EmployeeObjectType                       `xml:"urn:com.workday/bsvc Employee_Reference"`
	TerminationDate               time.Time                                 `xml:"urn:com.workday/bsvc Termination_Date"`
	LastDayofWork                 *time.Time                                `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
	PayThroughDate                *time.Time                                `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
	ReasonforSeveranceReference   *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_for_Severance_Reference"`
	ReasonforTerminationReference *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_for_Termination_Reference,omitempty"`
	Regrettable                   *bool                                     `xml:"urn:com.workday/bsvc Regrettable,omitempty"`
	UseUponTermination            *bool                                     `xml:"urn:com.workday/bsvc Use_Upon_Termination,omitempty"`
	NotifiedEmployeeDate          *time.Time                                `xml:"urn:com.workday/bsvc Notified_Employee_Date,omitempty"`
	SeveranceResponseDueDate      *time.Time                                `xml:"urn:com.workday/bsvc Severance_Response_Due_Date,omitempty"`
	WorksheetDeliveredDate        *time.Time                                `xml:"urn:com.workday/bsvc Worksheet_Delivered_Date,omitempty"`
	WorksheetSignedDate           *time.Time                                `xml:"urn:com.workday/bsvc Worksheet_Signed_Date,omitempty"`
	WorksheetDeclinedDate         *time.Time                                `xml:"urn:com.workday/bsvc Worksheet_Declined_Date,omitempty"`
	RevocationCutoffDate          *time.Time                                `xml:"urn:com.workday/bsvc Revocation_Cutoff_Date,omitempty"`
	SeverancePayComponentData     []SeverancePayComponentDataType           `xml:"urn:com.workday/bsvc Severance_Pay_Component_Data,omitempty"`
	SeverancePackageComponentData []SeverancePackageComponentDataType       `xml:"urn:com.workday/bsvc Severance_Package_Component_Data,omitempty"`
	PayRangeCurrencyReference     *CurrencyObjectType                       `xml:"urn:com.workday/bsvc Pay_Range_Currency_Reference,omitempty"`
}

func (t *EmployeeSeveranceWorksheetDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeSeveranceWorksheetDataType
	var layout struct {
		*T
		EffectiveDate            *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
		TerminationDate          *xsdDate `xml:"urn:com.workday/bsvc Termination_Date"`
		LastDayofWork            *xsdDate `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
		PayThroughDate           *xsdDate `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
		NotifiedEmployeeDate     *xsdDate `xml:"urn:com.workday/bsvc Notified_Employee_Date,omitempty"`
		SeveranceResponseDueDate *xsdDate `xml:"urn:com.workday/bsvc Severance_Response_Due_Date,omitempty"`
		WorksheetDeliveredDate   *xsdDate `xml:"urn:com.workday/bsvc Worksheet_Delivered_Date,omitempty"`
		WorksheetSignedDate      *xsdDate `xml:"urn:com.workday/bsvc Worksheet_Signed_Date,omitempty"`
		WorksheetDeclinedDate    *xsdDate `xml:"urn:com.workday/bsvc Worksheet_Declined_Date,omitempty"`
		RevocationCutoffDate     *xsdDate `xml:"urn:com.workday/bsvc Revocation_Cutoff_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	layout.TerminationDate = (*xsdDate)(&layout.T.TerminationDate)
	layout.LastDayofWork = (*xsdDate)(layout.T.LastDayofWork)
	layout.PayThroughDate = (*xsdDate)(layout.T.PayThroughDate)
	layout.NotifiedEmployeeDate = (*xsdDate)(layout.T.NotifiedEmployeeDate)
	layout.SeveranceResponseDueDate = (*xsdDate)(layout.T.SeveranceResponseDueDate)
	layout.WorksheetDeliveredDate = (*xsdDate)(layout.T.WorksheetDeliveredDate)
	layout.WorksheetSignedDate = (*xsdDate)(layout.T.WorksheetSignedDate)
	layout.WorksheetDeclinedDate = (*xsdDate)(layout.T.WorksheetDeclinedDate)
	layout.RevocationCutoffDate = (*xsdDate)(layout.T.RevocationCutoffDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeSeveranceWorksheetDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeSeveranceWorksheetDataType
	var overlay struct {
		*T
		EffectiveDate            *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
		TerminationDate          *xsdDate `xml:"urn:com.workday/bsvc Termination_Date"`
		LastDayofWork            *xsdDate `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
		PayThroughDate           *xsdDate `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
		NotifiedEmployeeDate     *xsdDate `xml:"urn:com.workday/bsvc Notified_Employee_Date,omitempty"`
		SeveranceResponseDueDate *xsdDate `xml:"urn:com.workday/bsvc Severance_Response_Due_Date,omitempty"`
		WorksheetDeliveredDate   *xsdDate `xml:"urn:com.workday/bsvc Worksheet_Delivered_Date,omitempty"`
		WorksheetSignedDate      *xsdDate `xml:"urn:com.workday/bsvc Worksheet_Signed_Date,omitempty"`
		WorksheetDeclinedDate    *xsdDate `xml:"urn:com.workday/bsvc Worksheet_Declined_Date,omitempty"`
		RevocationCutoffDate     *xsdDate `xml:"urn:com.workday/bsvc Revocation_Cutoff_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	overlay.TerminationDate = (*xsdDate)(&overlay.T.TerminationDate)
	overlay.LastDayofWork = (*xsdDate)(overlay.T.LastDayofWork)
	overlay.PayThroughDate = (*xsdDate)(overlay.T.PayThroughDate)
	overlay.NotifiedEmployeeDate = (*xsdDate)(overlay.T.NotifiedEmployeeDate)
	overlay.SeveranceResponseDueDate = (*xsdDate)(overlay.T.SeveranceResponseDueDate)
	overlay.WorksheetDeliveredDate = (*xsdDate)(overlay.T.WorksheetDeliveredDate)
	overlay.WorksheetSignedDate = (*xsdDate)(overlay.T.WorksheetSignedDate)
	overlay.WorksheetDeclinedDate = (*xsdDate)(overlay.T.WorksheetDeclinedDate)
	overlay.RevocationCutoffDate = (*xsdDate)(overlay.T.RevocationCutoffDate)
	return d.DecodeElement(&overlay, &start)
}

// Employee Severance Worksheet Event Data
type EmployeeSeveranceWorksheetEventDataType struct {
	EffectiveDate                            *time.Time                                     `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	SeverancePackageReference                *SeverancePackageObjectType                    `xml:"urn:com.workday/bsvc Severance_Package_Reference,omitempty"`
	EmployeeReference                        *WorkerObjectType                              `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	TerminationDate                          *time.Time                                     `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
	LastDayofWork                            *time.Time                                     `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
	PayThroughDate                           *time.Time                                     `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
	ReasonforSeveranceReference              *EventClassificationSubcategoryObjectType      `xml:"urn:com.workday/bsvc Reason_for_Severance_Reference,omitempty"`
	ReasonforTerminationReference            *EventClassificationSubcategoryObjectType      `xml:"urn:com.workday/bsvc Reason_for_Termination_Reference,omitempty"`
	RegrettableReference                     *CommonYesNoObjectType                         `xml:"urn:com.workday/bsvc Regrettable_Reference,omitempty"`
	UseUponTermination                       *bool                                          `xml:"urn:com.workday/bsvc Use_Upon_Termination,omitempty"`
	NotifiedEmployeeDate                     *time.Time                                     `xml:"urn:com.workday/bsvc Notified_Employee_Date,omitempty"`
	SeveranceResponseDate                    *time.Time                                     `xml:"urn:com.workday/bsvc Severance_Response_Date_,omitempty"`
	WorksheetDeliveredDate                   *time.Time                                     `xml:"urn:com.workday/bsvc Worksheet_Delivered_Date,omitempty"`
	WorksheetSignedDate                      *time.Time                                     `xml:"urn:com.workday/bsvc Worksheet_Signed_Date,omitempty"`
	WorksheetDeclinedDate                    *time.Time                                     `xml:"urn:com.workday/bsvc Worksheet_Declined_Date,omitempty"`
	RevocationCutoffDate                     *time.Time                                     `xml:"urn:com.workday/bsvc Revocation_Cutoff_Date,omitempty"`
	EmployeeSeverancePackagePayComponentData []EmployeeSeverancePackagePayComponentDataType `xml:"urn:com.workday/bsvc Employee_Severance_Package_Pay_Component_Data,omitempty"`
	SeverancePackageComponent                []SeverancePackageComponentDataType            `xml:"urn:com.workday/bsvc Severance_Package_Component,omitempty"`
}

func (t *EmployeeSeveranceWorksheetEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EmployeeSeveranceWorksheetEventDataType
	var layout struct {
		*T
		EffectiveDate          *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		TerminationDate        *xsdDate `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
		LastDayofWork          *xsdDate `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
		PayThroughDate         *xsdDate `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
		NotifiedEmployeeDate   *xsdDate `xml:"urn:com.workday/bsvc Notified_Employee_Date,omitempty"`
		SeveranceResponseDate  *xsdDate `xml:"urn:com.workday/bsvc Severance_Response_Date_,omitempty"`
		WorksheetDeliveredDate *xsdDate `xml:"urn:com.workday/bsvc Worksheet_Delivered_Date,omitempty"`
		WorksheetSignedDate    *xsdDate `xml:"urn:com.workday/bsvc Worksheet_Signed_Date,omitempty"`
		WorksheetDeclinedDate  *xsdDate `xml:"urn:com.workday/bsvc Worksheet_Declined_Date,omitempty"`
		RevocationCutoffDate   *xsdDate `xml:"urn:com.workday/bsvc Revocation_Cutoff_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	layout.TerminationDate = (*xsdDate)(layout.T.TerminationDate)
	layout.LastDayofWork = (*xsdDate)(layout.T.LastDayofWork)
	layout.PayThroughDate = (*xsdDate)(layout.T.PayThroughDate)
	layout.NotifiedEmployeeDate = (*xsdDate)(layout.T.NotifiedEmployeeDate)
	layout.SeveranceResponseDate = (*xsdDate)(layout.T.SeveranceResponseDate)
	layout.WorksheetDeliveredDate = (*xsdDate)(layout.T.WorksheetDeliveredDate)
	layout.WorksheetSignedDate = (*xsdDate)(layout.T.WorksheetSignedDate)
	layout.WorksheetDeclinedDate = (*xsdDate)(layout.T.WorksheetDeclinedDate)
	layout.RevocationCutoffDate = (*xsdDate)(layout.T.RevocationCutoffDate)
	return e.EncodeElement(layout, start)
}
func (t *EmployeeSeveranceWorksheetEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EmployeeSeveranceWorksheetEventDataType
	var overlay struct {
		*T
		EffectiveDate          *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
		TerminationDate        *xsdDate `xml:"urn:com.workday/bsvc Termination_Date,omitempty"`
		LastDayofWork          *xsdDate `xml:"urn:com.workday/bsvc Last_Day_of_Work,omitempty"`
		PayThroughDate         *xsdDate `xml:"urn:com.workday/bsvc Pay_Through_Date,omitempty"`
		NotifiedEmployeeDate   *xsdDate `xml:"urn:com.workday/bsvc Notified_Employee_Date,omitempty"`
		SeveranceResponseDate  *xsdDate `xml:"urn:com.workday/bsvc Severance_Response_Date_,omitempty"`
		WorksheetDeliveredDate *xsdDate `xml:"urn:com.workday/bsvc Worksheet_Delivered_Date,omitempty"`
		WorksheetSignedDate    *xsdDate `xml:"urn:com.workday/bsvc Worksheet_Signed_Date,omitempty"`
		WorksheetDeclinedDate  *xsdDate `xml:"urn:com.workday/bsvc Worksheet_Declined_Date,omitempty"`
		RevocationCutoffDate   *xsdDate `xml:"urn:com.workday/bsvc Revocation_Cutoff_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	overlay.TerminationDate = (*xsdDate)(overlay.T.TerminationDate)
	overlay.LastDayofWork = (*xsdDate)(overlay.T.LastDayofWork)
	overlay.PayThroughDate = (*xsdDate)(overlay.T.PayThroughDate)
	overlay.NotifiedEmployeeDate = (*xsdDate)(overlay.T.NotifiedEmployeeDate)
	overlay.SeveranceResponseDate = (*xsdDate)(overlay.T.SeveranceResponseDate)
	overlay.WorksheetDeliveredDate = (*xsdDate)(overlay.T.WorksheetDeliveredDate)
	overlay.WorksheetSignedDate = (*xsdDate)(overlay.T.WorksheetSignedDate)
	overlay.WorksheetDeclinedDate = (*xsdDate)(overlay.T.WorksheetDeclinedDate)
	overlay.RevocationCutoffDate = (*xsdDate)(overlay.T.RevocationCutoffDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type EmployeeSeveranceWorksheetEventObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmployeeSeveranceWorksheetEventObjectType struct {
	ID         []EmployeeSeveranceWorksheetEventObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Request Criteria to filter Employee Severance Worksheet Event results
type EmployeeSeveranceWorksheetEventRequestCriteriaType struct {
}

// Reference for Employee Severance Worksheet Event
type EmployeeSeveranceWorksheetEventRequestReferencesType struct {
	EmployeeSeveranceWorksheetEventReference []EmployeeSeveranceWorksheetEventObjectType `xml:"urn:com.workday/bsvc Employee_Severance_Worksheet_Event_Reference"`
}

// Response element for Retrieve Employee Severance Worksheet Events
type EmployeeSeveranceWorksheetEventResponseDataType struct {
	EmployeeSeveranceWorksheet []EmployeeSeveranceWorksheetEventType `xml:"urn:com.workday/bsvc Employee_Severance_Worksheet,omitempty"`
}

// Response Group for Employee Severance Worksheet Event
type EmployeeSeveranceWorksheetEventResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Contains reference to Employee Severance Worksheet Event and Employee Severance Worksheet Event Data
type EmployeeSeveranceWorksheetEventType struct {
	EmployeeSeveranceWorksheetEventReference *EmployeeSeveranceWorksheetEventObjectType `xml:"urn:com.workday/bsvc Employee_Severance_Worksheet_Event_Reference,omitempty"`
	EmployeeSeveranceWorksheetData           []EmployeeSeveranceWorksheetEventDataType  `xml:"urn:com.workday/bsvc Employee_Severance_Worksheet_Data,omitempty"`
}

// Element to specify the data for the type of severance matrix to put/update.
type EnhancedSeveranceMatrixDataType struct {
	EnhancedSeveranceMatrixID                      string                                              `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix_ID,omitempty"`
	EffectiveDate                                  time.Time                                           `xml:"urn:com.workday/bsvc Effective_Date"`
	EnhancedSeveranceMatrixName                    string                                              `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix_Name"`
	Inactive                                       *bool                                               `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	DurationUnitReference                          *CompensationPeriodObjectType                       `xml:"urn:com.workday/bsvc Duration_Unit_Reference"`
	ServiceDurationSeveranceMatrixData             *ServiceDurationSeveranceMatrixDataType             `xml:"urn:com.workday/bsvc Service_Duration_Severance_Matrix_Data,omitempty"`
	ServiceLengthSeveranceMatrixData               *ServiceLengthSeveranceMatrixDataType               `xml:"urn:com.workday/bsvc Service_Length_Severance_Matrix_Data,omitempty"`
	CompensationDurationSeveranceMatrixData        *CompensationDurationSeveranceMatrixDataType        `xml:"urn:com.workday/bsvc Compensation_Duration_Severance_Matrix_Data,omitempty"`
	CompensationLengthofServiceSeveranceMatrixData *CompensationLengthofServiceSeveranceMatrixDataType `xml:"urn:com.workday/bsvc Compensation_Length_of_Service_Severance_Matrix_Data,omitempty"`
}

func (t *EnhancedSeveranceMatrixDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T EnhancedSeveranceMatrixDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *EnhancedSeveranceMatrixDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EnhancedSeveranceMatrixDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Request Criteria
type EnhancedSeveranceMatrixRequestCriteriaType struct {
	Name                                 string                             `xml:"urn:com.workday/bsvc Name,omitempty"`
	EnhancedSeveranceMatrixFilterOptions []SeveranceMatrixFilterOptionsType `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix_Filter_Options,omitempty"`
	IncludeInactive                      *bool                              `xml:"urn:com.workday/bsvc Include_Inactive,omitempty"`
}

// Request references element.
type EnhancedSeveranceMatrixRequestReferencesType struct {
	EnhancedSeveranceMatrixReference []SeveranceMatrixabstractObjectType `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix_Reference"`
}

// The response for the Get Enhanced Severance Matrix web service.
type EnhancedSeveranceMatrixResponseDataType struct {
	EnhancedSeveranceMatrix []EnhancedSeveranceMatrixType `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix,omitempty"`
}

// The response group.
type EnhancedSeveranceMatrixResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// The Enhanced Severance Matrix element.
type EnhancedSeveranceMatrixType struct {
	EnhancedSeveranceMatrixReference *SeveranceMatrixabstractObjectType `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix_Reference,omitempty"`
	EnhancedSeveranceMatrixData      []EnhancedSeveranceMatrixDataType  `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix_Data,omitempty"`
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
type ExpenseAccumulatorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ExpenseAccumulatorObjectType struct {
	ID         []ExpenseAccumulatorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ExpenseItemObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ExpenseItemObjectType struct {
	ID         []ExpenseItemObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ExternalFieldObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
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

// Container for the processing options for sub-business processes within a business process. If no options are submitted (or the options are submitted as 'false') then the sub-business process is simply initiated as if it where submitted on-line with approvals, reviews, notifications and to-do's in place. If the Initiator is an Integration System User, any validations you configured on the Initiation step are ignored.
type FinancialsBusinessSubProcessParametersType struct {
	Skip                       *bool                           `xml:"urn:com.workday/bsvc Skip,omitempty"`
	BusinessProcessCommentData *BusinessProcessCommentDataType `xml:"urn:com.workday/bsvc Business_Process_Comment_Data,omitempty"`
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

// Future Payment Plan Assignment Request Criteria contains a set of zero or more Employees for whom Future Payment Plan Assignments are requested.
type FuturePaymentPlanAssignmentRequestCriteriaType struct {
	EmployeeReference []WorkerObjectType          `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	PositionReference []PositionElementObjectType `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
}

// Contains future payment plan assignment information based on Request Criteria.
type FuturePaymentPlanAssignmentResponseDataType struct {
	PositionFuturePaymentPlanAssignment []PositionFuturePaymentPlanAssignmentsType `xml:"urn:com.workday/bsvc Position_Future_Payment_Plan_Assignment,omitempty"`
}

// Future Payment Plan Data
type FuturePaymentPlanDataType struct {
	CompensationElementReference              *CompensationPayEarningObjectType            `xml:"urn:com.workday/bsvc Compensation_Element_Reference"`
	SetupSecuritySegmentReference             []CompensationSetupSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Setup_Security_Segment_Reference,omitempty"`
	CurrencyReference                         *CurrencyObjectType                          `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference                        *FrequencyObjectType                         `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	AllowIndividualTargets                    *bool                                        `xml:"urn:com.workday/bsvc Allow_Individual_Targets,omitempty"`
	IncludeAllBonusCompensationPlans          *bool                                        `xml:"urn:com.workday/bsvc Include_All_Bonus_Compensation_Plans,omitempty"`
	BonusCompensationPlanReference            []BonusPlanObjectType                        `xml:"urn:com.workday/bsvc Bonus_Compensation_Plan_Reference,omitempty"`
	IncludeAllOneTimePaymentCompensationPlans *bool                                        `xml:"urn:com.workday/bsvc Include_All_One-Time_Payment_Compensation_Plans,omitempty"`
	OneTimePaymentCompensationPlanReference   []OneTimePaymentPlanObjectType               `xml:"urn:com.workday/bsvc One-Time_Payment_Compensation_Plan_Reference,omitempty"`
	Amount                                    float64                                      `xml:"urn:com.workday/bsvc Amount"`
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

// Request element to get Benchmark Job
type GetBenchmarkJobsRequestType struct {
	RequestReferences *BenchmarkJobRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BenchmarkJobRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BenchmarkJobResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Benchmark Jobs Response
type GetBenchmarkJobsResponseType struct {
	RequestReferences *BenchmarkJobRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BenchmarkJobRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BenchmarkJobResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType               `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BenchmarkJobResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Compensation Eligibility Rules Request incoming element
type GetCompensationEligibilityRulesRequestType struct {
	RequestReferences *CompensationEligibilityRuleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CompensationEligibilityRuleRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CompensationEligibilityRuleResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Compensation Eligibility Rules Response element
type GetCompensationEligibilityRulesResponseType struct {
	RequestReferences *CompensationEligibilityRuleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CompensationEligibilityRuleRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                              `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *CompensationEligibilityRuleResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Compensation Eligibility Rules without Dependencies Request element
type GetCompensationEligibilityRuleswithoutDependenciesRequestType struct {
	RequestReferences *CompensationEligibilityRuleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CompensationEligibilityRuleRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to get compensation grade hierarchy information. If no criteria are specified in either the compensation grade hierarchy request references or compensation grade hierarchy request criteria elements, all instances will be returned.
type GetCompensationGradeHierarchyRequestType struct {
	RequestReferences *CompensationGradeHierarchyRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

type GetCompensationGradeHierarchyResponseType struct {
	RequestReferences *CompensationGradeHierarchyRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []CompensationGradeHierarchyResponseDataType     `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to get compensation grade and grade profile information.  If no criteria are specified in either the compensation grade request references or compensaiton grade request criteria elements, all instances of eligible earnings will be returned.
type GetCompensationGradesRequestType struct {
	RequestReferences *CompensationGradeRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CompensationGradeRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CompensationGradeResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing instances of compensation grades and their associated data.
type GetCompensationGradesResponseType struct {
	RequestReferences []CompensationGradeRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    []ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     []CompensationGradeResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   []ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []CompensationGradeResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Compensation Matrices Request Element
type GetCompensationMatricesRequestType struct {
	RequestReferences *CompensationMatrixRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CompensationMatrixRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CompensationMatrixResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Compensation Matrices Response element.
type GetCompensationMatricesResponseType struct {
	RequestReferences []CompensationMatrixRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   []CompensationMatrixRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    []ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     []CompensationMatrixResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   []ResponseResultsType                     `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []CompensationMatrixResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to get compensation plan information.  If no criteria is specified in the compensation plan request references or compensaiton plan request criteria elements, all compensation plans will be returned.
type GetCompensationPlansRequestType struct {
	RequestReferences *CompensationPlanRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CompensationPlanRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CompensationPlanResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing instances of compensation plans and the associated data.
type GetCompensationPlansResponseType struct {
	RequestReferences *CompensationPlanRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CompensationPlanRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CompensationPlanResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *CompensationPlanResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the request details to retrieve one or more scorecard results.
type GetCompensationScorecardResultsRequestType struct {
	RequestReferences *CompensationScorecardResultRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CompensationScorecardResultRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CompensationScorecardResultResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the response details to retrieve one or more scorecard results.
type GetCompensationScorecardResultsResponseType struct {
	RequestReferences *CompensationScorecardResultRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CompensationScorecardResultRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CompensationScorecardResultResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                              `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *CompensationScorecardResultResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the request details to retrieve one or more scorecards.
type GetCompensationScorecardsRequestType struct {
	RequestReferences *CompensationScorecardRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CompensationScorecardRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CompensationScorecardResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the response details to retrieve one or more scorecard.
type GetCompensationScorecardsResponseType struct {
	RequestReferences *CompensationScorecardRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CompensationScorecardRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CompensationScorecardResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                        `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *CompensationScorecardResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to get eligible earnings override information for an employee, eligible earnings override period or percent based bonus plan.  If no criteria are specified in either the Eligible Earnings Request References or Eligible Earnings Request Criteria elements, all instances of eligible earnings override will be returned.
type GetEligibleEarningsRequestType struct {
	RequestReferences *EligibleEarningsRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *EligibleEarningRequestCriteriaType    `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *EligibleEarningsResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing instances of Eligible Earnings and their associated data.
type GetEligibleEarningsResponseType struct {
	RequestReferences []EligibleEarningsRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   []EligibleEarningRequestCriteriaType    `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    []ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     []EligibleEarningsResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   []ResponseResultsType                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []EligibleEarningsResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Enhanced Severance Matrix request element.
type GetEnhancedSeveranceMatricesRequestType struct {
	RequestReferences *EnhancedSeveranceMatrixRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *EnhancedSeveranceMatrixRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                           `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *EnhancedSeveranceMatrixResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response of the Get Enhanced Severance Matrices web service
type GetEnhancedSeveranceMatricesResponseType struct {
	RequestReferences *EnhancedSeveranceMatrixRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *EnhancedSeveranceMatrixRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                           `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *EnhancedSeveranceMatrixResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                          `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *EnhancedSeveranceMatrixResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to get future payment plan assignments. If no criteria are specified then all assignments active for today will be returned, grouped by employee.
type GetFuturePaymentPlanAssignmentsRequestType struct {
	RequestCriteria *FuturePaymentPlanAssignmentRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version         string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Future Payment Plan Assignments Response element containing retrieved Future Payment Plan Assignments and associated web service request related criteria.
type GetFuturePaymentPlanAssignmentsResponseType struct {
	RequestCriteria *FuturePaymentPlanAssignmentRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter  *ResponseFilterType                             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults *ResponseResultsType                            `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData    *FuturePaymentPlanAssignmentResponseDataType    `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version         string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element
type GetOneTimePaymentPlanConfigurableCategoriesRequestType struct {
	RequestReferences *OneTimePaymentPlanConfigurableCategoryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *OneTimePaymentPlanConfigurableCategoryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *OneTimePaymentPlanConfigurableCategoryResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response
type GetOneTimePaymentPlanConfigurableCategoriesResponseType struct {
	RequestReferences *OneTimePaymentPlanConfigurableCategoryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *OneTimePaymentPlanConfigurableCategoryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *OneTimePaymentPlanConfigurableCategoryResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                                         `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *OneTimePaymentPlanConfigurableCategoryResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to get period activity rate matrix information. If no criteria is specified in either the period activity rate matrix request references or period activity rate matrix request criteria elements, all active instances will be returned.
type GetPeriodActivityRateMatricesRequestType struct {
	RequestReferences *PeriodActivityRateMatrixRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PeriodActivityRateMatrixRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Period Activity Rate Matrices response
type GetPeriodActivityRateMatricesResponseType struct {
	RequestReferences *PeriodActivityRateMatrixRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PeriodActivityRateMatrixRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                           `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PeriodActivityRateMatrixResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to get period activity task information.  If no criteria are specified in either the period activity task request references or period activity task request criteria elements, all active instances will be returned.
type GetPeriodActivityTasksRequestType struct {
	RequestReferences *PeriodActivityTaskRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PeriodActivityTaskRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PeriodActivityTaskResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing instances of period activity tasks and their associated data.
type GetPeriodActivityTasksResponseType struct {
	RequestReferences *PeriodActivityTaskRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PeriodActivityTaskRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PeriodActivityTaskResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                     `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PeriodActivityTaskResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element used to return a worker's compensation history loaded from a previous system.
// If no worker reference is requested, then all workers with previous system compensation data will be returned.
type GetPreviousSystemCompensationHistoryRequestType struct {
	RequestReferences *WorkerRequestReferencesType         `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *WorkerResponseGroupforReferenceType `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing the instances of workers and their previous system compensation history.
type GetPreviousSystemCompensationHistoryResponseType struct {
	RequestReferences []WorkerRequestReferencesType                       `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    []ResponseFilterType                                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   []ResponseResultsType                               `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseGroup     []WorkerResponseGroupforReferenceType               `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseData      []PreviousSystemCompensationHistoryResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to get stock participation rate table information.  If no criteria is specified in the request references or request criteria elements, then all stock participation rate tables will be returned.
type GetStockParticipationRateTablesRequestType struct {
	RequestReferences *StockParticipationRateTableRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *StockParticipationRateTableRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *StockParticipationRateTableResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing instances of stock participation rate tables and their associated data.
type GetStockParticipationRateTablesResponseType struct {
	RequestReferences *StockParticipationRateTableRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *StockParticipationRateTableRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *StockParticipationRateTableResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                              `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *StockParticipationRateTableResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Data for Grant Type Split.
type GrantTypeSplitReplacementDataType struct {
	StockGrantTypeReference       *StockGrantTypeObjectType       `xml:"urn:com.workday/bsvc Stock_Grant_Type_Reference"`
	StockVestingScheduleReference *StockVestingScheduleObjectType `xml:"urn:com.workday/bsvc Stock_Vesting_Schedule_Reference,omitempty"`
	StockDateRuleReference        *StockDateRuleObjectType        `xml:"urn:com.workday/bsvc Stock_Date_Rule_Reference,omitempty"`
	GrantSplitPercent             float64                         `xml:"urn:com.workday/bsvc Grant_Split_Percent,omitempty"`
}

// This is a wrapper for the Hourly Data.
type HourlyPlanDataType struct {
	Amount      float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	MinimumWage *bool   `xml:"urn:com.workday/bsvc Minimum_Wage,omitempty"`
}

// This element contains an Eligible Earnings Override.
type ImportEligibleEarningsOverrideHVRequestType struct {
	EligibleEarningsReference *EligibleEarningsOverrideObjectType `xml:"urn:com.workday/bsvc Eligible_Earnings_Reference,omitempty"`
	EligibleEarningsData      *EligibleEarningsOverrideHVDataType `xml:"urn:com.workday/bsvc Eligible_Earnings_Data"`
}

// Loads eligible earnings overrides for a specific period. To apply the override to a specific bonus plan, add the bonus plan to the operation filter. To apply the override to all current and future bonus plans, select "Apply to All Bonus Plans”. An in-progress Compensation Review Process immediately implements these overrides if the Participation Options are configured to support this. Unlike the Put Eligible Earnings operation, if the input template contains any errors, the entire operation fails.
type ImportEligibleEarningsOverrideRequestType struct {
	EligibleEarningsOverridePeriodReference *EligibleEarningsOverridePeriodObjectType     `xml:"urn:com.workday/bsvc Eligible_Earnings_Override_Period_Reference"`
	EligibleEarnings                        []ImportEligibleEarningsOverrideHVRequestType `xml:"urn:com.workday/bsvc Eligible_Earnings,omitempty"`
	Version                                 string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Root Import Request Compensation Change Request Element
type ImportRequestCompensationChangeRequestType struct {
	ID                                      string                                   `xml:"urn:com.workday/bsvc ID,omitempty"`
	RequestCompensationRequestInformationHV []RequestCompensationChangeRequestHVType `xml:"urn:com.workday/bsvc Request_Compensation_Request_Information_HV,omitempty"`
	Version                                 string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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
type JobRequisitionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JobRequisitionObjectType struct {
	ID         []JobRequisitionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This operation allows the management of period activity based assignments for a given worker and a bounding date range or period via the Manage Period Activity Pay Assignments business process.
type ManagePeriodActivityPayAssignmentsRequestType struct {
	BusinessProcessParameters    *BusinessProcessParametersType        `xml:"urn:com.workday/bsvc Business_Process_Parameters"`
	PeriodActivityPayAssignments *PeriodActivityPayAssignmentsDataType `xml:"urn:com.workday/bsvc Period_Activity_Pay_Assignments"`
	Version                      string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the detailed business process event information resulting from a successful Manage Period Activity Pay Assignments request.
type ManagePeriodActivityPayAssignmentsResponseType struct {
	PeriodActivityPayAssigmentEvent *PeriodActivityAssignmentEventDataType `xml:"urn:com.workday/bsvc Period_Activity_Pay_Assigment_Event,omitempty"`
	Version                         string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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
type MeritPercentPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type MeritPercentPlanObjectType struct {
	ID         []MeritPercentPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Data specific to a Merit Amount Plan
type MeritPlanAmountDataType struct {
	Amount                                float64                                     `xml:"urn:com.workday/bsvc Amount,omitempty"`
	CompensationBasisReference            *CompensationBasisObjectType                `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
	ManageByCompensationBasisRules        *bool                                       `xml:"urn:com.workday/bsvc Manage_By_Compensation_Basis_Rules,omitempty"`
	MeritPlanProfileAmountReplacementData []MeritPlanProfileAmountReplacementDataType `xml:"urn:com.workday/bsvc Merit_Plan_Profile_Amount_Replacement_Data,omitempty"`
}

// Merit Plan Data
type MeritPlanDataType struct {
	SetupSecuritySegmentReference                         []CompensationSetupSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Setup_Security_Segment_Reference,omitempty"`
	CurrencyReference                                     *CurrencyObjectType                          `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference                                    *FrequencyObjectType                         `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	SubjecttoStatutoryMinimum                             *bool                                        `xml:"urn:com.workday/bsvc Subject_to_Statutory_Minimum,omitempty"`
	AllowIndividualTargets                                *bool                                        `xml:"urn:com.workday/bsvc Allow_Individual_Targets,omitempty"`
	CompensationMatrixReference                           *CompensationMatrixObjectType                `xml:"urn:com.workday/bsvc Compensation_Matrix_Reference,omitempty"`
	UseMatrixAsReferenceOnly                              *bool                                        `xml:"urn:com.workday/bsvc Use_Matrix_As_Reference_Only,omitempty"`
	WaitingPeriodReference                                *EligibilityWaitingPeriodObjectType          `xml:"urn:com.workday/bsvc Waiting_Period_Reference,omitempty"`
	IncludeActiveEmployeesinWaitingPeriod                 *bool                                        `xml:"urn:com.workday/bsvc Include_Active_Employees_in_Waiting_Period,omitempty"`
	TimeProrationRuleReference                            *TimeProrationRuleObjectType                 `xml:"urn:com.workday/bsvc Time_Proration_Rule_Reference,omitempty"`
	RoundingRuleReference                                 *CompensationRoundingRuleObjectType          `xml:"urn:com.workday/bsvc Rounding_Rule_Reference,omitempty"`
	IncludeActiveEmployeesAssignedPlanDuringProcessPeriod *bool                                        `xml:"urn:com.workday/bsvc Include_Active_Employees_Assigned_Plan_During_Process_Period,omitempty"`
	PercentData                                           *MeritPlanPercentDataType                    `xml:"urn:com.workday/bsvc Percent_Data,omitempty"`
	AmountData                                            *MeritPlanAmountDataType                     `xml:"urn:com.workday/bsvc Amount_Data,omitempty"`
	HideTarget                                            *bool                                        `xml:"urn:com.workday/bsvc Hide_Target,omitempty"`
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

// Required Percent Data for Merit Plan; consists of default percent and/or  Merit Plan Profiles.
type MeritPlanPercentDataType struct {
	Percentage                             float64                                      `xml:"urn:com.workday/bsvc Percentage,omitempty"`
	ManageByCompensationBasisRules         *bool                                        `xml:"urn:com.workday/bsvc Manage_By_Compensation_Basis_Rules,omitempty"`
	CompensationBasisReference             *CompensationBasisObjectType                 `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
	MeritPlanProfilePercentReplacementData []MeritPlanProfilePercentReplacementDataType `xml:"urn:com.workday/bsvc Merit_Plan_Profile_Percent_Replacement_Data,omitempty"`
}

// Data for a Merit Amount plan profile.
type MeritPlanProfileAmountReplacementDataType struct {
	ProfileAmount              float64                      `xml:"urn:com.workday/bsvc Profile_Amount,omitempty"`
	CurrencyReference          *CurrencyObjectType          `xml:"urn:com.workday/bsvc Currency_Reference"`
	ConditionRuleReference     *ConditionRuleObjectType     `xml:"urn:com.workday/bsvc Condition_Rule_Reference"`
	CompensationBasisReference *CompensationBasisObjectType `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
}

// This is a wrapper for the Merit Plan Percent Plan Profile Percent Replacement; used to display or replace Percent Plan Profile Data.
type MeritPlanProfilePercentReplacementDataType struct {
	EligibilityRuleReference   *ConditionRuleObjectType     `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	Percentage                 float64                      `xml:"urn:com.workday/bsvc Percentage,omitempty"`
	CompensationBasisReference *CompensationBasisObjectType `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
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

// Data that is persisted
type OneTimePaymentPlanConfigurableCategoryDataType struct {
	ID       string `xml:"urn:com.workday/bsvc ID,omitempty"`
	Category string `xml:"urn:com.workday/bsvc Category"`
}

// Contains a unique identifier for an instance of an object.
type OneTimePaymentPlanConfigurableCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type OneTimePaymentPlanConfigurableCategoryObjectType struct {
	ID         []OneTimePaymentPlanConfigurableCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Request Criteria
type OneTimePaymentPlanConfigurableCategoryRequestCriteriaType struct {
}

// Request References
type OneTimePaymentPlanConfigurableCategoryRequestReferencesType struct {
	OneTimePaymentPlanConfigurableCategoryReference []OneTimePaymentPlanConfigurableCategoryObjectType `xml:"urn:com.workday/bsvc One-Time_Payment_Plan_Configurable_Category_Reference"`
}

// Response Element
type OneTimePaymentPlanConfigurableCategoryResponseDataType struct {
	OneTimePaymentPlanConfigurableCategory []OneTimePaymentPlanConfigurableCategoryType `xml:"urn:com.workday/bsvc One-Time_Payment_Plan_Configurable_Category,omitempty"`
}

// Response Group
type OneTimePaymentPlanConfigurableCategoryResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Response Data
type OneTimePaymentPlanConfigurableCategoryType struct {
	OneTimePaymentPlanConfigurableCategoryReference *OneTimePaymentPlanConfigurableCategoryObjectType `xml:"urn:com.workday/bsvc One-Time_Payment_Plan_Configurable_Category_Reference,omitempty"`
	OneTimePaymentPlanConfigurableCategoryData      *OneTimePaymentPlanConfigurableCategoryDataType   `xml:"urn:com.workday/bsvc One-Time_Payment_Plan_Configurable_Category_Data,omitempty"`
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

// Contains a unique identifier for an instance of an object.
type PayRangeQuartileObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayRangeQuartileObjectType struct {
	ID         []PayRangeQuartileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayRangeSegmentTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayRangeSegmentTypeObjectType struct {
	ID         []PayRangeSegmentTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains information on scorecard.
type PerformanceCriteriaDataType struct {
	GoalID          string  `xml:"urn:com.workday/bsvc Goal_ID,omitempty"`
	GoalName        string  `xml:"urn:com.workday/bsvc Goal_Name"`
	GoalDescription string  `xml:"urn:com.workday/bsvc Goal_Description,omitempty"`
	GoalWeight      float64 `xml:"urn:com.workday/bsvc Goal_Weight"`
}

// Contains a unique identifier for an instance of an object.
type PerformanceCriteriaObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PerformanceCriteriaObjectType struct {
	ID         []PerformanceCriteriaObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Performance Factor Compensation Matrix Weighted Data is part of Performance Factors Data within a Bonus Plan; there must be one Performance Factor Compensation Matrix Weighted Data for a Performance Factors Data.
type PerformanceFactorCompensationMatrixWeightedDataType struct {
	CompensationMatrixWeightedReference *CompensationWeightedPercentMatrixObjectType `xml:"urn:com.workday/bsvc Compensation_Matrix_Weighted_Reference,omitempty"`
	ApplyNetAttainment                  *bool                                        `xml:"urn:com.workday/bsvc Apply_Net_Attainment,omitempty"`
	Weight                              float64                                      `xml:"urn:com.workday/bsvc Weight"`
	BonusModifierReference              *DefaultScorecardObjectType                  `xml:"urn:com.workday/bsvc Bonus_Modifier_Reference,omitempty"`
}

// Performance Factor Scorecard Data is part of Performance Factors Data within a Bonus Plan.  There must be at least one Performance Factor Scorecard Data within the Performance Factors Data.
type PerformanceFactorScorecardDataType struct {
	ScorecardReference     *DefaultScorecardObjectType `xml:"urn:com.workday/bsvc Scorecard_Reference"`
	Weight                 float64                     `xml:"urn:com.workday/bsvc Weight"`
	BonusModifierReference *DefaultScorecardObjectType `xml:"urn:com.workday/bsvc Bonus_Modifier_Reference,omitempty"`
}

// Performance Factors Data is part of Bonus Plan and is mutually exclusive with Performance Matrix Data.  It consists of one or more Performance Factor Scorecard Data and one Performance Factor Compensation Matrix Weighted Data.
type PerformanceFactorsDataType struct {
	PerformanceFactorScorecardData                  []PerformanceFactorScorecardDataType                 `xml:"urn:com.workday/bsvc Performance_Factor_Scorecard_Data,omitempty"`
	PerformanceFactorCompensationMatrixWeightedData *PerformanceFactorCompensationMatrixWeightedDataType `xml:"urn:com.workday/bsvc Performance_Factor_Compensation_Matrix_Weighted_Data,omitempty"`
}

// This is a wrapper for the Bonus Plan Performance Matrix and is mutually exclusive with Performance Factors Data.
type PerformanceMatrixDataType struct {
	CompensationMatrixReference *CompensationMatrixObjectType `xml:"urn:com.workday/bsvc Compensation_Matrix_Reference,omitempty"`
	UseMatrixasReferenceOnly    *bool                         `xml:"urn:com.workday/bsvc Use_Matrix_as_Reference_Only,omitempty"`
	ApplyNetAttainment          *bool                         `xml:"urn:com.workday/bsvc Apply_Net_Attainment,omitempty"`
	BonusModifierReference      *DefaultScorecardObjectType   `xml:"urn:com.workday/bsvc Bonus_Modifier_Reference,omitempty"`
}

// Contains assignments which have been removed from the employee.  An assignment can only be removed if no pay activity has occurred.
type PeriodActivityAssignmentDeletedDataType struct {
	PeriodActivityAssignmentReference *PeriodActivityAssignmentObjectType `xml:"urn:com.workday/bsvc Period_Activity_Assignment_Reference,omitempty"`
	PeriodActivityReference           *PeriodActivityObjectType           `xml:"urn:com.workday/bsvc Period_Activity_Reference,omitempty"`
	PeriodActivityTaskReference       *PeriodActivityTaskObjectType       `xml:"urn:com.workday/bsvc Period_Activity_Task_Reference,omitempty"`
	PeriodActivityStartDate           *time.Time                          `xml:"urn:com.workday/bsvc Period_Activity_Start_Date,omitempty"`
	PeriodActivityEndDate             *time.Time                          `xml:"urn:com.workday/bsvc Period_Activity_End_Date,omitempty"`
	Comment                           string                              `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

func (t *PeriodActivityAssignmentDeletedDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PeriodActivityAssignmentDeletedDataType
	var layout struct {
		*T
		PeriodActivityStartDate *xsdDate `xml:"urn:com.workday/bsvc Period_Activity_Start_Date,omitempty"`
		PeriodActivityEndDate   *xsdDate `xml:"urn:com.workday/bsvc Period_Activity_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PeriodActivityStartDate = (*xsdDate)(layout.T.PeriodActivityStartDate)
	layout.PeriodActivityEndDate = (*xsdDate)(layout.T.PeriodActivityEndDate)
	return e.EncodeElement(layout, start)
}
func (t *PeriodActivityAssignmentDeletedDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PeriodActivityAssignmentDeletedDataType
	var overlay struct {
		*T
		PeriodActivityStartDate *xsdDate `xml:"urn:com.workday/bsvc Period_Activity_Start_Date,omitempty"`
		PeriodActivityEndDate   *xsdDate `xml:"urn:com.workday/bsvc Period_Activity_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PeriodActivityStartDate = (*xsdDate)(overlay.T.PeriodActivityStartDate)
	overlay.PeriodActivityEndDate = (*xsdDate)(overlay.T.PeriodActivityEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a reference to the Manage Period Activity Pay Assignments business process event and the assignment details.
type PeriodActivityAssignmentEventDataType struct {
	PeriodActivityAssignmentEventReference *UniqueIdentifierObjectType               `xml:"urn:com.workday/bsvc Period_Activity_Assignment_Event_Reference,omitempty"`
	EffectiveDate                          *time.Time                                `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	EmployeeReference                      *EmployeeObjectType                       `xml:"urn:com.workday/bsvc Employee__Reference,omitempty"`
	PositionReference                      *PositionObjectType                       `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	AcademicPeriodReference                *AcademicPeriodObjectType                 `xml:"urn:com.workday/bsvc Academic_Period_Reference,omitempty"`
	PeriodActivityRateMatrixReference      *PeriodActivityRateMatrixObjectType       `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Reference,omitempty"`
	ReasonReference                        *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
	PeriodActivityPayAssignment            []PeriodActivityAssignmentVersionDataType `xml:"urn:com.workday/bsvc Period_Activity_Pay_Assignment,omitempty"`
	PeriodActivityPayAssignmentRemoved     []PeriodActivityAssignmentDeletedDataType `xml:"urn:com.workday/bsvc Period_Activity_Pay_Assignment_Removed,omitempty"`
}

func (t *PeriodActivityAssignmentEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PeriodActivityAssignmentEventDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PeriodActivityAssignmentEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PeriodActivityAssignmentEventDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PeriodActivityAssignmentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PeriodActivityAssignmentObjectType struct {
	ID         []PeriodActivityAssignmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the individual assignment details of a new activity, changes to an existing activity, or an existing activity to be deleted.
type PeriodActivityAssignmentVersionDataType struct {
	PeriodActivityPayAssignmentReference *PeriodActivityAssignmentObjectType `xml:"urn:com.workday/bsvc Period_Activity_Pay_Assignment_Reference,omitempty"`
	PeriodActivityReference              *PeriodActivityObjectType           `xml:"urn:com.workday/bsvc Period_Activity_Reference,omitempty"`
	PeriodActivityTaskReference          *PeriodActivityTaskObjectType       `xml:"urn:com.workday/bsvc Period_Activity_Task_Reference,omitempty"`
	PeriodActivityStartDate              *time.Time                          `xml:"urn:com.workday/bsvc Period_Activity_Start_Date,omitempty"`
	PeriodActivityEndDate                *time.Time                          `xml:"urn:com.workday/bsvc Period_Activity_End_Date,omitempty"`
	CurrencyReference                    *CurrencyObjectType                 `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	AssignedAmount                       float64                             `xml:"urn:com.workday/bsvc Assigned_Amount,omitempty"`
	AssignedQuantity                     float64                             `xml:"urn:com.workday/bsvc Assigned_Quantity,omitempty"`
	AssignedUnitRate                     float64                             `xml:"urn:com.workday/bsvc Assigned_Unit_Rate,omitempty"`
	PaymentStartDate                     *time.Time                          `xml:"urn:com.workday/bsvc Payment_Start_Date,omitempty"`
	PaymentEndDate                       *time.Time                          `xml:"urn:com.workday/bsvc Payment_End_Date,omitempty"`
	DoNotPay                             *bool                               `xml:"urn:com.workday/bsvc Do_Not_Pay,omitempty"`
	Comment                              string                              `xml:"urn:com.workday/bsvc Comment,omitempty"`
	PeriodActivityPayCosting             *PeriodActivityPayCostingDataType   `xml:"urn:com.workday/bsvc Period_Activity_Pay_Costing,omitempty"`
}

func (t *PeriodActivityAssignmentVersionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PeriodActivityAssignmentVersionDataType
	var layout struct {
		*T
		PeriodActivityStartDate *xsdDate `xml:"urn:com.workday/bsvc Period_Activity_Start_Date,omitempty"`
		PeriodActivityEndDate   *xsdDate `xml:"urn:com.workday/bsvc Period_Activity_End_Date,omitempty"`
		PaymentStartDate        *xsdDate `xml:"urn:com.workday/bsvc Payment_Start_Date,omitempty"`
		PaymentEndDate          *xsdDate `xml:"urn:com.workday/bsvc Payment_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PeriodActivityStartDate = (*xsdDate)(layout.T.PeriodActivityStartDate)
	layout.PeriodActivityEndDate = (*xsdDate)(layout.T.PeriodActivityEndDate)
	layout.PaymentStartDate = (*xsdDate)(layout.T.PaymentStartDate)
	layout.PaymentEndDate = (*xsdDate)(layout.T.PaymentEndDate)
	return e.EncodeElement(layout, start)
}
func (t *PeriodActivityAssignmentVersionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PeriodActivityAssignmentVersionDataType
	var overlay struct {
		*T
		PeriodActivityStartDate *xsdDate `xml:"urn:com.workday/bsvc Period_Activity_Start_Date,omitempty"`
		PeriodActivityEndDate   *xsdDate `xml:"urn:com.workday/bsvc Period_Activity_End_Date,omitempty"`
		PaymentStartDate        *xsdDate `xml:"urn:com.workday/bsvc Payment_Start_Date,omitempty"`
		PaymentEndDate          *xsdDate `xml:"urn:com.workday/bsvc Payment_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PeriodActivityStartDate = (*xsdDate)(overlay.T.PeriodActivityStartDate)
	overlay.PeriodActivityEndDate = (*xsdDate)(overlay.T.PeriodActivityEndDate)
	overlay.PaymentStartDate = (*xsdDate)(overlay.T.PaymentStartDate)
	overlay.PaymentEndDate = (*xsdDate)(overlay.T.PaymentEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PeriodActivityCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PeriodActivityCategoryObjectType struct {
	ID         []PeriodActivityCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element containing eligible period activity data.
type PeriodActivityEligibilityEntryDataType struct {
	PeriodActivityReference     *PeriodActivityObjectType     `xml:"urn:com.workday/bsvc Period_Activity_Reference"`
	PeriodActivityTaskReference *PeriodActivityTaskObjectType `xml:"urn:com.workday/bsvc Period_Activity_Task_Reference,omitempty"`
	AssignedUnitQuantity        float64                       `xml:"urn:com.workday/bsvc Assigned_Unit_Quantity,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PeriodActivityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PeriodActivityObjectType struct {
	ID         []PeriodActivityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the individual assignment details of a new activity, changes to an existing activity, or an existing activity to be deleted.
type PeriodActivityPayAssignmentDataType struct {
	PeriodActivityPayAssignmentReference       *PeriodActivityAssignmentObjectType    `xml:"urn:com.workday/bsvc Period_Activity_Pay_Assignment_Reference,omitempty"`
	PeriodActivityReference                    *PeriodActivityObjectType              `xml:"urn:com.workday/bsvc Period_Activity_Reference,omitempty"`
	PeriodActivityTaskorCourseSectionReference *PeriodActivityTaskInterfaceObjectType `xml:"urn:com.workday/bsvc Period_Activity_Task_or_Course_Section_Reference,omitempty"`
	PeriodActivityStartDate                    *time.Time                             `xml:"urn:com.workday/bsvc Period_Activity_Start_Date,omitempty"`
	PeriodActivityEndDate                      *time.Time                             `xml:"urn:com.workday/bsvc Period_Activity_End_Date,omitempty"`
	CurrencyReference                          *CurrencyObjectType                    `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	AssignedAmount                             float64                                `xml:"urn:com.workday/bsvc Assigned_Amount,omitempty"`
	AssignedQuantity                           float64                                `xml:"urn:com.workday/bsvc Assigned_Quantity,omitempty"`
	AssignedUnitRate                           float64                                `xml:"urn:com.workday/bsvc Assigned_Unit_Rate,omitempty"`
	PaymentStartDate                           *time.Time                             `xml:"urn:com.workday/bsvc Payment_Start_Date,omitempty"`
	PaymentEndDate                             *time.Time                             `xml:"urn:com.workday/bsvc Payment_End_Date,omitempty"`
	DoNotPay                                   *bool                                  `xml:"urn:com.workday/bsvc Do_Not_Pay,omitempty"`
	Comment                                    string                                 `xml:"urn:com.workday/bsvc Comment,omitempty"`
	PeriodActivityPayCosting                   *PeriodActivityPayCostingDataType      `xml:"urn:com.workday/bsvc Period_Activity_Pay_Costing,omitempty"`
	Delete                                     bool                                   `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

func (t *PeriodActivityPayAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PeriodActivityPayAssignmentDataType
	var layout struct {
		*T
		PeriodActivityStartDate *xsdDate `xml:"urn:com.workday/bsvc Period_Activity_Start_Date,omitempty"`
		PeriodActivityEndDate   *xsdDate `xml:"urn:com.workday/bsvc Period_Activity_End_Date,omitempty"`
		PaymentStartDate        *xsdDate `xml:"urn:com.workday/bsvc Payment_Start_Date,omitempty"`
		PaymentEndDate          *xsdDate `xml:"urn:com.workday/bsvc Payment_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PeriodActivityStartDate = (*xsdDate)(layout.T.PeriodActivityStartDate)
	layout.PeriodActivityEndDate = (*xsdDate)(layout.T.PeriodActivityEndDate)
	layout.PaymentStartDate = (*xsdDate)(layout.T.PaymentStartDate)
	layout.PaymentEndDate = (*xsdDate)(layout.T.PaymentEndDate)
	return e.EncodeElement(layout, start)
}
func (t *PeriodActivityPayAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PeriodActivityPayAssignmentDataType
	var overlay struct {
		*T
		PeriodActivityStartDate *xsdDate `xml:"urn:com.workday/bsvc Period_Activity_Start_Date,omitempty"`
		PeriodActivityEndDate   *xsdDate `xml:"urn:com.workday/bsvc Period_Activity_End_Date,omitempty"`
		PaymentStartDate        *xsdDate `xml:"urn:com.workday/bsvc Payment_Start_Date,omitempty"`
		PaymentEndDate          *xsdDate `xml:"urn:com.workday/bsvc Payment_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PeriodActivityStartDate = (*xsdDate)(overlay.T.PeriodActivityStartDate)
	overlay.PeriodActivityEndDate = (*xsdDate)(overlay.T.PeriodActivityEndDate)
	overlay.PaymentStartDate = (*xsdDate)(overlay.T.PaymentStartDate)
	overlay.PaymentEndDate = (*xsdDate)(overlay.T.PaymentEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains Period Activity Assignment information for a given Worker and Period range in a format suitable for constructing a Manage Period Activity Pay operation request.
type PeriodActivityPayAssignmentsDataType struct {
	EffectiveDate                     time.Time                                 `xml:"urn:com.workday/bsvc Effective_Date"`
	EmployeeReference                 *EmployeeObjectType                       `xml:"urn:com.workday/bsvc Employee_Reference"`
	PositionReference                 *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	AcademicPeriodReference           *AcademicPeriodObjectType                 `xml:"urn:com.workday/bsvc Academic_Period_Reference"`
	PeriodActivityRateMatrixReference *PeriodActivityRateMatrixObjectType       `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Reference"`
	ReasonReference                   *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference"`
	PeriodActivityPayAssignment       []PeriodActivityPayAssignmentDataType     `xml:"urn:com.workday/bsvc Period_Activity_Pay_Assignment"`
}

func (t *PeriodActivityPayAssignmentsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PeriodActivityPayAssignmentsDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PeriodActivityPayAssignmentsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PeriodActivityPayAssignmentsDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains one or more costing allocation splits for the specific Period Activity Pay assignment line
type PeriodActivityPayCostingDataType struct {
	AllocationDetailforPeriodPay []AllocationDetailforPeriodPayDataType `xml:"urn:com.workday/bsvc Allocation_Detail_for_Period_Pay"`
}

// Period Activity Rate Matrix data
type PeriodActivityRateMatrixDataType struct {
	ID                                string                                  `xml:"urn:com.workday/bsvc ID,omitempty"`
	PeriodActivityRateTableName       string                                  `xml:"urn:com.workday/bsvc Period_Activity_Rate_Table_Name"`
	EffectiveDate                     *time.Time                              `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	ConditionRuleReference            []ConditionRuleObjectType               `xml:"urn:com.workday/bsvc Condition_Rule_Reference,omitempty"`
	CurrencyReference                 *CurrencyObjectType                     `xml:"urn:com.workday/bsvc Currency_Reference"`
	Inactive                          *bool                                   `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	PeriodActivityRateMatrixEntryData []PeriodActivityRateMatrixEntryDataType `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Entry_Data"`
}

func (t *PeriodActivityRateMatrixDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PeriodActivityRateMatrixDataType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PeriodActivityRateMatrixDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PeriodActivityRateMatrixDataType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Period Activity Rate Matrix entry data for each row in the matrix
type PeriodActivityRateMatrixEntryDataType struct {
	PeriodActivityCategoryReference           *PeriodActivityCategoryObjectType              `xml:"urn:com.workday/bsvc Period_Activity_Category_Reference"`
	PeriodActivityUnitReference               *PeriodActivityUnitObjectType                  `xml:"urn:com.workday/bsvc Period_Activity_Unit_Reference"`
	PeriodActivityRateMatrixEntrySequenceData *PeriodActivityRateMatrixEntrySequenceDataType `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Entry_Sequence_Data"`
}

// Period Activity Rate Matrix entry data for rate and accelerator
type PeriodActivityRateMatrixEntrySequenceDataType struct {
	DefaultUnitRate          float64                     `xml:"urn:com.workday/bsvc Default_Unit_Rate,omitempty"`
	AcceleratorAmount        float64                     `xml:"urn:com.workday/bsvc Accelerator_Amount,omitempty"`
	AcceleratorPercent       float64                     `xml:"urn:com.workday/bsvc Accelerator_Percent,omitempty"`
	AcceleratorBaseReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Accelerator_Base_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PeriodActivityRateMatrixObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PeriodActivityRateMatrixObjectType struct {
	ID         []PeriodActivityRateMatrixObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Request criteria
type PeriodActivityRateMatrixRequestCriteriaType struct {
	IncludeInactive *bool `xml:"urn:com.workday/bsvc Include_Inactive,omitempty"`
}

// Utilize the period activity rate matrix to retrieve specific instance(s) of period activity rate matrix and its associated data.
type PeriodActivityRateMatrixRequestReferencesType struct {
	PeriodActivityRateMatrixReference []PeriodActivityRateMatrixObjectType `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Reference"`
}

// Contains Period Activity Rate Matrix information
type PeriodActivityRateMatrixResponseDataType struct {
	PeriodActivityRateMatrix []PeriodActivityRateMatrixType `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix,omitempty"`
}

// A Period Activity Rate Matrix
type PeriodActivityRateMatrixType struct {
	PeriodActivityRateMatrixReference *PeriodActivityRateMatrixObjectType `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Reference,omitempty"`
	PeriodActivityRateMatrixData      []PeriodActivityRateMatrixDataType  `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Data"`
}

// Contains Period Activity Task data.
type PeriodActivityTaskDataType struct {
	ID                            string                        `xml:"urn:com.workday/bsvc ID,omitempty"`
	PeriodActivityTaskCode        string                        `xml:"urn:com.workday/bsvc Period_Activity_Task_Code"`
	PeriodActivityTaskName        string                        `xml:"urn:com.workday/bsvc Period_Activity_Task_Name"`
	PeriodActivityTaskDescription string                        `xml:"urn:com.workday/bsvc Period_Activity_Task_Description,omitempty"`
	PeriodActivityUnitReference   *PeriodActivityUnitObjectType `xml:"urn:com.workday/bsvc Period_Activity_Unit_Reference,omitempty"`
	AcademicPeriodReference       *AcademicPeriodObjectType     `xml:"urn:com.workday/bsvc Academic_Period_Reference,omitempty"`
	DefaultUnitQuantity           float64                       `xml:"urn:com.workday/bsvc Default_Unit_Quantity,omitempty"`
	AllowUnitQuantityOverride     *bool                         `xml:"urn:com.workday/bsvc Allow_Unit_Quantity_Override,omitempty"`
	Inactive                      *bool                         `xml:"urn:com.workday/bsvc Inactive,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PeriodActivityTaskInterfaceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PeriodActivityTaskInterfaceObjectType struct {
	ID         []PeriodActivityTaskInterfaceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PeriodActivityTaskObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PeriodActivityTaskObjectType struct {
	ID         []PeriodActivityTaskObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Utilize the period activity task to retrieve specific instance(s) of period activity task and its associated data.
type PeriodActivityTaskRequestCriteriaType struct {
	IncludeInactive *bool `xml:"urn:com.workday/bsvc Include_Inactive,omitempty"`
}

// Utilize the period activity task to retrieve specific instance(s) of period activity task and its associated data.
type PeriodActivityTaskRequestReferencesType struct {
	PeriodActivityTaskReference []PeriodActivityTaskObjectType `xml:"urn:com.workday/bsvc Period_Activity_Task_Reference"`
}

// Contains period activity task information based on Request References or Request Criteria.
type PeriodActivityTaskResponseDataType struct {
	PeriodActivityTask []PeriodActivityTaskType `xml:"urn:com.workday/bsvc Period_Activity_Task,omitempty"`
}

// Provides the filters for the web service operation.
type PeriodActivityTaskResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Contains a period activity task and its associated data.
type PeriodActivityTaskType struct {
	PeriodActivityTaskReference *PeriodActivityTaskObjectType `xml:"urn:com.workday/bsvc Period_Activity_Task_Reference,omitempty"`
	PeriodActivityTaskData      []PeriodActivityTaskDataType  `xml:"urn:com.workday/bsvc Period_Activity_Task_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PeriodActivityUnitObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PeriodActivityUnitObjectType struct {
	ID         []PeriodActivityUnitObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Period Plan Data
type PeriodPlanDataType struct {
	CompensationElementReference              *CompensationPayEarningObjectType            `xml:"urn:com.workday/bsvc Compensation_Element_Reference"`
	DefaultCompensationPeriodMultiplier       float64                                      `xml:"urn:com.workday/bsvc Default_Compensation_Period_Multiplier,omitempty"`
	UnitofDurationReference                   *CompensationPeriodObjectType                `xml:"urn:com.workday/bsvc Unit_of_Duration_Reference"`
	CompensationBasisReference                *CompensationBasisObjectType                 `xml:"urn:com.workday/bsvc Compensation_Basis_Reference"`
	PeriodPlanProfileReplacementData          []PeriodPlanProfileReplacementDataType       `xml:"urn:com.workday/bsvc Period_Plan_Profile_Replacement_Data,omitempty"`
	CompensationSetupSecuritySegmentReference []CompensationSetupSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Compensation_Setup_Security_Segment_Reference,omitempty"`
}

// Profile data for Period Salary Plan.
type PeriodPlanProfileReplacementDataType struct {
	EligibilityRuleReference            *ConditionRuleObjectType     `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	DefaultCompensationPeriodMultiplier float64                      `xml:"urn:com.workday/bsvc Default_Compensation_Period_Multiplier"`
	CompensationBasisReference          *CompensationBasisObjectType `xml:"urn:com.workday/bsvc Compensation_Basis_Reference"`
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

// Future payment plan assignments for a position effective for specific date.
type PositionFuturePaymentPlanAssignmentDataType struct {
	EmployeeReference               *WorkerObjectType                     `xml:"urn:com.workday/bsvc Employee_Reference"`
	PositionReference               *PositionElementObjectType            `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	AssignmentDate                  *time.Time                            `xml:"urn:com.workday/bsvc Assignment_Date,omitempty"`
	FuturePaymentPlanAssignmentData []FuturePaymentPlanAssignmentDataType `xml:"urn:com.workday/bsvc Future_Payment_Plan_Assignment_Data,omitempty"`
}

func (t *PositionFuturePaymentPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PositionFuturePaymentPlanAssignmentDataType
	var layout struct {
		*T
		AssignmentDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssignmentDate = (*xsdDate)(layout.T.AssignmentDate)
	return e.EncodeElement(layout, start)
}
func (t *PositionFuturePaymentPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PositionFuturePaymentPlanAssignmentDataType
	var overlay struct {
		*T
		AssignmentDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssignmentDate = (*xsdDate)(overlay.T.AssignmentDate)
	return d.DecodeElement(&overlay, &start)
}

// Groups an Employee's Future Payment Plan Assignments.
type PositionFuturePaymentPlanAssignmentsType struct {
	PositionFuturePaymentPlanAssignmentData []PositionFuturePaymentPlanAssignmentDataType `xml:"urn:com.workday/bsvc Position_Future_Payment_Plan_Assignment_Data,omitempty"`
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
type PotentialObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PotentialObjectType struct {
	ID         []PotentialObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Container for defining whether a previous system compensation history entry is being added, updated or deleted.
type PreviousSystemCompensationHistoryDataType struct {
	PreviousSystemCompensationHistoryReference  *CompensationPreviousSystemHistoryObjectType     `xml:"urn:com.workday/bsvc Previous_System_Compensation_History_Reference,omitempty"`
	PreviousSystemCompensationHistoryDetailData *PreviousSystemCompensationHistoryDetailDataType `xml:"urn:com.workday/bsvc Previous_System_Compensation_History_Detail_Data,omitempty"`
	AddOnly                                     bool                                             `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Delete                                      bool                                             `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}

// Container for the detailed data for a previous system compensation history entry.
type PreviousSystemCompensationHistoryDetailDataType struct {
	ID                 string               `xml:"urn:com.workday/bsvc ID,omitempty"`
	WorkerHistoryName  string               `xml:"urn:com.workday/bsvc Worker_History_Name,omitempty"`
	ActionDate         *time.Time           `xml:"urn:com.workday/bsvc Action_Date,omitempty"`
	Reason             string               `xml:"urn:com.workday/bsvc Reason,omitempty"`
	Amount             float64              `xml:"urn:com.workday/bsvc Amount,omitempty"`
	CurrencyReference  *CurrencyObjectType  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	FrequencyReference *FrequencyObjectType `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
	AmountChange       float64              `xml:"urn:com.workday/bsvc Amount_Change,omitempty"`
	Description        string               `xml:"urn:com.workday/bsvc Description,omitempty"`
}

func (t *PreviousSystemCompensationHistoryDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PreviousSystemCompensationHistoryDetailDataType
	var layout struct {
		*T
		ActionDate *xsdDate `xml:"urn:com.workday/bsvc Action_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ActionDate = (*xsdDate)(layout.T.ActionDate)
	return e.EncodeElement(layout, start)
}
func (t *PreviousSystemCompensationHistoryDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PreviousSystemCompensationHistoryDetailDataType
	var overlay struct {
		*T
		ActionDate *xsdDate `xml:"urn:com.workday/bsvc Action_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ActionDate = (*xsdDate)(overlay.T.ActionDate)
	return d.DecodeElement(&overlay, &start)
}

// Container for a worker and the previous system compensation history.
type PreviousSystemCompensationHistoryGetDataType struct {
	WorkerReference                       *WorkerObjectType                       `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	PreviousSystemCompensationHistoryData []PreviousSystemCompensationHistoryType `xml:"urn:com.workday/bsvc Previous_System_Compensation_History_Data,omitempty"`
}

// Contains the worker and their compensation history from a previous system.
type PreviousSystemCompensationHistoryResponseDataType struct {
	PreviousSystemCompensationHistory []PreviousSystemCompensationHistoryGetDataType `xml:"urn:com.workday/bsvc Previous_System_Compensation_History,omitempty"`
}

// Contains the data for adding, updating or deleting a previous system compensation history entry for a worker.
type PreviousSystemCompensationHistoryType struct {
	WorkerReference          *WorkerObjectType                           `xml:"urn:com.workday/bsvc Worker_Reference"`
	PreviousSystemJobHistory []PreviousSystemCompensationHistoryDataType `xml:"urn:com.workday/bsvc Previous_System_Job_History"`
}

type ProcessingFaultType struct {
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
}

// Contains a rule/profile delineated set of goals and goal values for this result.
type ProfileCompensationScorecardResultDataType struct {
	EligibilityRuleReference *ConditionRuleObjectType         `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	ScorecardResultData      []ProfileScorecardResultDataType `xml:"urn:com.workday/bsvc Scorecard_Result_Data"`
}

// Contains the set of goals and goal values for this result.
type ProfileScorecardResultDataType struct {
	GoalReference *PerformanceCriteriaObjectType `xml:"urn:com.workday/bsvc Goal_Reference"`
	Achievement   float64                        `xml:"urn:com.workday/bsvc Achievement,omitempty"`
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

// Request element to put Benchmark Job.
type PutBenchmarkJobRequestType struct {
	BenchmarkJobReference *CompensationBenchmarkDefaultObjectType `xml:"urn:com.workday/bsvc Benchmark_Job_Reference,omitempty"`
	BenchmarkJobData      *BenchmarkJobDataType                   `xml:"urn:com.workday/bsvc Benchmark_Job_Data"`
	AddOnly               bool                                    `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version               string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element to put Benchmark Job
type PutBenchmarkJobResponseType struct {
	BenchmarkJobReference *CompensationBenchmarkDefaultObjectType `xml:"urn:com.workday/bsvc Benchmark_Job_Reference,omitempty"`
	Version               string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Incoming Put Compensation Eligibility Rule Request
type PutCompensationEligibilityRuleRequestType struct {
	CompensationEligibilityRuleReference *ConditionRuleObjectType             `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule_Reference,omitempty"`
	CompensationEligibilityRuleData      *CompensationEligibilityRuleDataType `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule_Data,omitempty"`
	AddOnly                              bool                                 `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Delete                               bool                                 `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	Version                              string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Compensation Eligibility Rule Response element
type PutCompensationEligibilityRuleResponseType struct {
	CompensationEligibilityRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule_Reference,omitempty"`
	Version                              string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

type PutCompensationGradeHierarchyRequestType struct {
	CompensationGradeHierarchyData *CompensationGradeHierarchyDataType `xml:"urn:com.workday/bsvc Compensation_Grade_Hierarchy_Data,omitempty"`
	Version                        string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

type PutCompensationGradeHierarchyResponseType struct {
	EffectiveDate           *time.Time                              `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	CompensationGradeLevels []PutCompensationGradeLevelResponseType `xml:"urn:com.workday/bsvc Compensation_Grade_Levels,omitempty"`
	Version                 string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

func (t *PutCompensationGradeHierarchyResponseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PutCompensationGradeHierarchyResponseType
	var layout struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *PutCompensationGradeHierarchyResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PutCompensationGradeHierarchyResponseType
	var overlay struct {
		*T
		EffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

type PutCompensationGradeLevelResponseType struct {
	CompensationGradeLevelReference *CompensationGradeLevelObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Level_Reference,omitempty"`
}

// This web service operation is designed to load compensation grades and the associated compensation grade profiles.  A complete set of compensation grade profiles should be passed in with the compensation grade.  Any profiles not found in the set will be deleted, new profiles will be added, and the existing profiles will be updated.
type PutCompensationGradeRequestType struct {
	CompensationGradeReference *CompensationGradeObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Reference,omitempty"`
	CompensationGradeData      *CompensationGradeDataType   `xml:"urn:com.workday/bsvc Compensation_Grade_Data"`
	AddOnly                    bool                         `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the reference ID for the compensation grade.
type PutCompensationGradeResponseType struct {
	CompensationGradeReference *CompensationGradeObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Reference,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// A request object to put Compensation Matrix data.
type PutCompensationMatrixRequestType struct {
	CompensationMatrixReference *CompensationMatrixObjectType `xml:"urn:com.workday/bsvc Compensation_Matrix_Reference,omitempty"`
	CompensationMatrixData      *CompensationMatrixDataType   `xml:"urn:com.workday/bsvc Compensation_Matrix_Data"`
	Version                     string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response to a request to modify or create Compensation Matrix data.
type PutCompensationMatrixResponseType struct {
	MeritIncreaseMatrixReference *CompensationMatrixObjectType `xml:"urn:com.workday/bsvc Merit_Increase_Matrix_Reference,omitempty"`
	Version                      string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to put Compensation Plan (Bonus, Merit, Allowance, Future Payment, Commission, and Base Pay Plans).
type PutCompensationPlanRequestType struct {
	CompensationPlanReference *CompensationAssignablePlanObjectType `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
	CompensationPlanData      *CompensationPlanDataType             `xml:"urn:com.workday/bsvc Compensation_Plan_Data"`
	AddOnly                   bool                                  `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                   string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element to put Compensation Plan (Bonus, Merit, Allowance, Future Payment, Commission, and Base Pay Plans).
type PutCompensationPlanResponseType struct {
	CompensationPlanReference *CompensationAssignablePlanObjectType `xml:"urn:com.workday/bsvc Compensation_Plan_Reference"`
	Version                   string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains scorecard information.
type PutCompensationScorecardRequestType struct {
	CompensationScorecardReference *DefaultScorecardObjectType    `xml:"urn:com.workday/bsvc Compensation_Scorecard_Reference,omitempty"`
	CompensationScorecardData      *CompensationScorecardDataType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Data"`
	AddOnly                        bool                           `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                        string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Compensation Scorecard Response.
type PutCompensationScorecardResponseType struct {
	CompensationScorecardReference *DefaultScorecardObjectType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Reference,omitempty"`
	Version                        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the scorecard result request information.  This includes the detailed scorecard result data and an optional reference an existing scorecard result to update.
type PutCompensationScorecardResultRequestType struct {
	CompensationScorecardResultReference *ScoresetContainerObjectType         `xml:"urn:com.workday/bsvc Compensation_Scorecard_Result_Reference,omitempty"`
	CompensationScorecardResultData      *CompensationScorecardResultDataType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Result_Data"`
	AddOnly                              bool                                 `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                              string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the scorecard result identifier of the added or modified result.
type PutCompensationScorecardResultResponseType struct {
	CompensationScorecardResultReference *ScoresetContainerObjectType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Result_Reference,omitempty"`
	Version                              string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This web service operation is designed to load eligible earnings overrides for an employee for a specified eligible earnings override period.  If bonus plans are specified, then the eligible earnings override will apply only to those specific bonus plans.  Alternatively, selecting "Apply to All Bonus Plans" option will enable the eligible earnings override to be applied to all current and future bonus plans assigned to the employee for the specified period.
type PutEligibleEarningsRequestType struct {
	EligibleEarningsReference *EligibleEarningsOverrideObjectType `xml:"urn:com.workday/bsvc Eligible_Earnings_Reference,omitempty"`
	EligibleEarningsData      *EligibleEarningsDataType           `xml:"urn:com.workday/bsvc Eligible_Earnings_Data"`
	AddOnly                   bool                                `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                   string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the reference ID for the Eligible Earnings
type PutEligibleEarningsResponseType struct {
	EligibleEarningsReference *EligibleEarningsOverrideObjectType `xml:"urn:com.workday/bsvc Eligible_Earnings_Reference,omitempty"`
	Version                   string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The element containing the Put Enhanced Severance Matrix web service sent.
type PutEnhancedSeveranceMatrixRequestType struct {
	EnhancedSeveranceMatrixReference *SeveranceMatrixabstractObjectType `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix_Reference,omitempty"`
	EnhancedSeveranceMatrixData      *EnhancedSeveranceMatrixDataType   `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix_Data"`
	Version                          string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response for the Put Enhanced Severance Matrix web service.
type PutEnhancedSeveranceMatrixResponseType struct {
	SeveranceMatrixabstractReference *SeveranceMatrixabstractObjectType `xml:"urn:com.workday/bsvc Severance_Matrix__abstract__Reference,omitempty"`
	Version                          string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This web service assigns a future payment plan to an employee for consumption in the future payment true up bonus process.
type PutFuturePaymentPlanAssignmentRequestType struct {
	PositionFuturePaymentPlanAssignmentData *PositionFuturePaymentPlanAssignmentDataType `xml:"urn:com.workday/bsvc Position_Future_Payment_Plan_Assignment_Data"`
	Version                                 string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response to a put future payment plan assignment request.
type PutFuturePaymentPlanAssignmentResponseType struct {
	EmployeeReference *WorkerObjectType   `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
	PositionReference *PositionObjectType `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	AssignmentDate    *time.Time          `xml:"urn:com.workday/bsvc Assignment_Date,omitempty"`
	Version           string              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

func (t *PutFuturePaymentPlanAssignmentResponseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PutFuturePaymentPlanAssignmentResponseType
	var layout struct {
		*T
		AssignmentDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AssignmentDate = (*xsdDate)(layout.T.AssignmentDate)
	return e.EncodeElement(layout, start)
}
func (t *PutFuturePaymentPlanAssignmentResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PutFuturePaymentPlanAssignmentResponseType
	var overlay struct {
		*T
		AssignmentDate *xsdDate `xml:"urn:com.workday/bsvc Assignment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AssignmentDate = (*xsdDate)(overlay.T.AssignmentDate)
	return d.DecodeElement(&overlay, &start)
}

// Put Import Process Response element
type PutImportProcessResponseType struct {
	ImportProcessReference  *WebServiceBackgroundProcessRuntimeObjectType `xml:"urn:com.workday/bsvc Import_Process_Reference,omitempty"`
	HeaderInstanceReference *InstanceObjectType                           `xml:"urn:com.workday/bsvc Header_Instance_Reference,omitempty"`
	Version                 string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Incoming request element
type PutOneTimePaymentPlanConfigurableCategoryRequestType struct {
	OneTimePaymentPlanConfigurableCategoryReference *OneTimePaymentPlanConfigurableCategoryObjectType `xml:"urn:com.workday/bsvc One-Time_Payment_Plan_Configurable_Category_Reference,omitempty"`
	OneTimePaymentPlanConfigurableCategoryData      *OneTimePaymentPlanConfigurableCategoryDataType   `xml:"urn:com.workday/bsvc One-Time_Payment_Plan_Configurable_Category_Data"`
	AddOnly                                         bool                                              `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                                         string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response Element
type PutOneTimePaymentPlanConfigurableCategoryResponseType struct {
	OneTimePaymentPlanConfigurableCategoryReference *OneTimePaymentPlanConfigurableCategoryObjectType `xml:"urn:com.workday/bsvc One-Time_Payment_Plan_Configurable_Category_Reference,omitempty"`
	Version                                         string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request to create or update a Period Activity Rate Matrix
type PutPeriodActivityRateMatrixRequestType struct {
	PeriodActivityRateMatrixReference *PeriodActivityRateMatrixObjectType `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Reference,omitempty"`
	PeriodActivityRateMatrixData      *PeriodActivityRateMatrixDataType   `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Data"`
	AddOnly                           bool                                `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Period Activity Rate Matrix response
type PutPeriodActivityRateMatrixResponseType struct {
	PeriodActivityRateMatrixReference *PeriodActivityRateMatrixObjectType `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Reference,omitempty"`
	Version                           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This web service operation is designed to load period activity tasks.
type PutPeriodActivityTaskRequestType struct {
	PeriodActivityTaskReference *PeriodActivityTaskObjectType `xml:"urn:com.workday/bsvc Period_Activity_Task_Reference,omitempty"`
	PeriodActivityTaskData      *PeriodActivityTaskDataType   `xml:"urn:com.workday/bsvc Period_Activity_Task_Data"`
	AddOnly                     bool                          `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                     string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Responds with the reference ID for the period activity task.
type PutPeriodActivityTaskResponseType struct {
	PeriodActivityTaskReference *PeriodActivityTaskObjectType `xml:"urn:com.workday/bsvc Period_Activity_Task_Reference,omitempty"`
	Version                     string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the data for adding, updating or deleting a previous system compensation history entry for a worker.
type PutPreviousSystemCompensationHistoryRequestType struct {
	PreviousSystemCompensationHistoryData *PreviousSystemCompensationHistoryType `xml:"urn:com.workday/bsvc Previous_System_Compensation_History_Data"`
	Version                               string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response for Put Previous System Compensation History. Returns the worker.
type PutPreviousSystemCompensationHistoryResponseType struct {
	WorkerReference *WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	Version         string            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request element to put Stock Participation Rate Table.
type PutStockParticipationRateTableRequestType struct {
	StockParticipationRateTableReference *StockParticipationRateTableObjectType `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_Reference,omitempty"`
	StockParticipationRateTableData      *StockParticipationRateTableData20Type `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_Data,omitempty"`
	Delete                               bool                                   `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	Version                              string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Stock Participation Rate Table Response
type PutStockParticipationRateTableResponseType struct {
	StockParticipationRateTableReference *StockParticipationRateTableObjectType `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_Reference,omitempty"`
	Version                              string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Specific compensation plan assigned to the worker to be removed.
type RemoveCompensationPlanAssignmentDataType struct {
	CompensationPlanReference []CompensationPlanObjectType `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
}

// Wrapper element for the Request Bonus Payment business process.
type RequestBonusPaymentDataType struct {
	EmployeeReference                       *EmployeeObjectType                       `xml:"urn:com.workday/bsvc Employee_Reference"`
	PositionReference                       *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	EffectiveDate                           time.Time                                 `xml:"urn:com.workday/bsvc Effective_Date"`
	BonusCompensationSnapshotDate           *time.Time                                `xml:"urn:com.workday/bsvc Bonus_Compensation_Snapshot_Date,omitempty"`
	EligibleEarningsOverridePeriodReference *EligibleEarningsOverridePeriodObjectType `xml:"urn:com.workday/bsvc Eligible_Earnings_Override_Period_Reference,omitempty"`
	BonusReasonReference                    *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Bonus_Reason_Reference,omitempty"`
	BonusPaymentData                        *BonusPaymentDataType                     `xml:"urn:com.workday/bsvc Bonus_Payment_Data"`
	IgnorePlanAssignment                    *bool                                     `xml:"urn:com.workday/bsvc Ignore_Plan_Assignment,omitempty"`
}

func (t *RequestBonusPaymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RequestBonusPaymentDataType
	var layout struct {
		*T
		EffectiveDate                 *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
		BonusCompensationSnapshotDate *xsdDate `xml:"urn:com.workday/bsvc Bonus_Compensation_Snapshot_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	layout.BonusCompensationSnapshotDate = (*xsdDate)(layout.T.BonusCompensationSnapshotDate)
	return e.EncodeElement(layout, start)
}
func (t *RequestBonusPaymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RequestBonusPaymentDataType
	var overlay struct {
		*T
		EffectiveDate                 *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
		BonusCompensationSnapshotDate *xsdDate `xml:"urn:com.workday/bsvc Bonus_Compensation_Snapshot_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	overlay.BonusCompensationSnapshotDate = (*xsdDate)(overlay.T.BonusCompensationSnapshotDate)
	return d.DecodeElement(&overlay, &start)
}

// This web service operation loads the approved bonus for an employee assigned to a bonus plan using the Request Bonus Payment business process.
type RequestBonusPaymentRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	RequestBonusPaymentData   *RequestBonusPaymentDataType   `xml:"urn:com.workday/bsvc Request_Bonus_Payment_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing the reference to the Bonus Payment Request
type RequestBonusPaymentResponseType struct {
	BonusPaymentEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Bonus_Payment_Event_Reference,omitempty"`
	Version                    string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper Element for the Request Compensation Change business process.
type RequestCompensationChangeDataType struct {
	EmployeeReference                 *EmployeeObjectType                        `xml:"urn:com.workday/bsvc Employee_Reference"`
	PositionReference                 *PositionElementObjectType                 `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	CompensationChangeDate            time.Time                                  `xml:"urn:com.workday/bsvc Compensation_Change_Date"`
	CompensationChangeOnNextPayPeriod bool                                       `xml:"urn:com.workday/bsvc Compensation_Change_On_Next_Pay_Period"`
	EmployeeVisibilityDate            *time.Time                                 `xml:"urn:com.workday/bsvc Employee_Visibility_Date,omitempty"`
	CompensationChangeData            *CompensationChangeDataType                `xml:"urn:com.workday/bsvc Compensation_Change_Data"`
	CheckPositionBudgetSubProcess     *CheckPositionBudgetSubBusinessProcessType `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process,omitempty"`
}

func (t *RequestCompensationChangeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RequestCompensationChangeDataType
	var layout struct {
		*T
		CompensationChangeDate *xsdDate `xml:"urn:com.workday/bsvc Compensation_Change_Date"`
		EmployeeVisibilityDate *xsdDate `xml:"urn:com.workday/bsvc Employee_Visibility_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CompensationChangeDate = (*xsdDate)(&layout.T.CompensationChangeDate)
	layout.EmployeeVisibilityDate = (*xsdDate)(layout.T.EmployeeVisibilityDate)
	return e.EncodeElement(layout, start)
}
func (t *RequestCompensationChangeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RequestCompensationChangeDataType
	var overlay struct {
		*T
		CompensationChangeDate *xsdDate `xml:"urn:com.workday/bsvc Compensation_Change_Date"`
		EmployeeVisibilityDate *xsdDate `xml:"urn:com.workday/bsvc Employee_Visibility_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CompensationChangeDate = (*xsdDate)(&overlay.T.CompensationChangeDate)
	overlay.EmployeeVisibilityDate = (*xsdDate)(overlay.T.EmployeeVisibilityDate)
	return d.DecodeElement(&overlay, &start)
}

// Web service operation to request a compensation change for employee(s). This operation is asynchronous and specifically designed to support very large documents.
type RequestCompensationChangeRequestHVType struct {
	BusinessProcessParameters     *BusinessProcessParametersType     `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	RequestCompensationChangeData *RequestCompensationChangeDataType `xml:"urn:com.workday/bsvc Request_Compensation_Change_Data"`
}

// Web service operation to request a compensation change for an employee.
type RequestCompensationChangeRequestType struct {
	BusinessProcessParameters     *BusinessProcessParametersType     `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	RequestCompensationChangeData *RequestCompensationChangeDataType `xml:"urn:com.workday/bsvc Request_Compensation_Change_Data"`
	Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for the Request Compensation Change business process.
type RequestCompensationChangeResponseType struct {
	RequestCompensationChangeEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Request_Compensation_Change_Event_Reference,omitempty"`
	Version                                 string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper Element for the Request Employee Merit Adjustment business process.
type RequestEmployeeMeritAdjustmentDataType struct {
	EmployeeReference              *EmployeeObjectType                       `xml:"urn:com.workday/bsvc Employee_Reference"`
	PositionReference              *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	MeritIncreaseEffectiveDate     time.Time                                 `xml:"urn:com.workday/bsvc Merit_Increase_Effective_Date"`
	MeritPlanReference             *MeritPercentPlanObjectType               `xml:"urn:com.workday/bsvc Merit_Plan_Reference,omitempty"`
	MeritReasonReference           *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Merit_Reason_Reference,omitempty"`
	EmployeeMeritAdjustmentSubData []EmployeeMeritAdjustmentDataType         `xml:"urn:com.workday/bsvc Employee_Merit_Adjustment_Sub_Data"`
}

func (t *RequestEmployeeMeritAdjustmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RequestEmployeeMeritAdjustmentDataType
	var layout struct {
		*T
		MeritIncreaseEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Merit_Increase_Effective_Date"`
	}
	layout.T = (*T)(t)
	layout.MeritIncreaseEffectiveDate = (*xsdDate)(&layout.T.MeritIncreaseEffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *RequestEmployeeMeritAdjustmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RequestEmployeeMeritAdjustmentDataType
	var overlay struct {
		*T
		MeritIncreaseEffectiveDate *xsdDate `xml:"urn:com.workday/bsvc Merit_Increase_Effective_Date"`
	}
	overlay.T = (*T)(t)
	overlay.MeritIncreaseEffectiveDate = (*xsdDate)(&overlay.T.MeritIncreaseEffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// Web service operation to request a merit adjustment for an employee.
type RequestEmployeeMeritAdjustmentRequestType struct {
	BusinessProcessParameters   *BusinessProcessParametersType          `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	EmployeeMeritAdjustmentData *RequestEmployeeMeritAdjustmentDataType `xml:"urn:com.workday/bsvc Employee_Merit_Adjustment_Data"`
	Version                     string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for the Request Employee Merit Adjustment business process.
type RequestEmployeeMeritAdjustmentResponseType struct {
	EmployeeMeritAdjustmentEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Employee_Merit_Adjustment_Event_Reference,omitempty"`
	Version                               string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Wrapper element for the Request One-Time Payment business process.
type RequestOneTimePaymentDataType struct {
	EmployeeReference      *EmployeeObjectType                       `xml:"urn:com.workday/bsvc Employee_Reference"`
	PositionReference      *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
	EffectiveDate          time.Time                                 `xml:"urn:com.workday/bsvc Effective_Date"`
	EmployeeVisibilityDate *time.Time                                `xml:"urn:com.workday/bsvc Employee_Visibility_Date,omitempty"`
	ReasonReference        *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
	OneTimePaymentSubData  []OneTimePaymentDataType                  `xml:"urn:com.workday/bsvc One-Time_Payment_Sub_Data"`
}

func (t *RequestOneTimePaymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RequestOneTimePaymentDataType
	var layout struct {
		*T
		EffectiveDate          *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
		EmployeeVisibilityDate *xsdDate `xml:"urn:com.workday/bsvc Employee_Visibility_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	layout.EmployeeVisibilityDate = (*xsdDate)(layout.T.EmployeeVisibilityDate)
	return e.EncodeElement(layout, start)
}
func (t *RequestOneTimePaymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RequestOneTimePaymentDataType
	var overlay struct {
		*T
		EffectiveDate          *xsdDate `xml:"urn:com.workday/bsvc Effective_Date"`
		EmployeeVisibilityDate *xsdDate `xml:"urn:com.workday/bsvc Employee_Visibility_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	overlay.EmployeeVisibilityDate = (*xsdDate)(overlay.T.EmployeeVisibilityDate)
	return d.DecodeElement(&overlay, &start)
}

// This web service operation is designed to pay an employee a one-time payment such as a signing bonus using the Request One-Time Payment  business process.
type RequestOneTimePaymentRequestType struct {
	BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	OneTimePaymentData        *RequestOneTimePaymentDataType `xml:"urn:com.workday/bsvc One-Time_Payment_Data"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element containing the reference for the one time payment.
type RequestOneTimePaymentResponseType struct {
	OneTimePaymentEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc One-Time_Payment_Event_Reference,omitempty"`
	Version                      string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Web service operation to request a requisition compensation change for a job requisition
type RequestRequisitionCompensationChangeRequestType struct {
	BusinessProcessParameters         *BusinessProcessParametersType         `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	RequisitionCompensationChangeData *RequisitionCompensationChangeDataType `xml:"urn:com.workday/bsvc Requisition_Compensation_Change_Data"`
	Version                           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for the Request Requisition Compensation Change business process.
type RequestRequisitionCompensationChangeResponseType struct {
	CompensationEventforJobRequisitionReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Compensation_Event_for_Job_Requisition_Reference,omitempty"`
	Version                                     string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element for the Request Requisition Compensation Change business process.
type RequisitionCompensationChangeDataType struct {
	JobRequisitionReference     *JobRequisitionObjectType        `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
	CompensationChangeDate      time.Time                        `xml:"urn:com.workday/bsvc Compensation_Change_Date"`
	RequisitionCompensationData *RequisitionCompensationDataType `xml:"urn:com.workday/bsvc Requisition_Compensation_Data"`
}

func (t *RequisitionCompensationChangeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RequisitionCompensationChangeDataType
	var layout struct {
		*T
		CompensationChangeDate *xsdDate `xml:"urn:com.workday/bsvc Compensation_Change_Date"`
	}
	layout.T = (*T)(t)
	layout.CompensationChangeDate = (*xsdDate)(&layout.T.CompensationChangeDate)
	return e.EncodeElement(layout, start)
}
func (t *RequisitionCompensationChangeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RequisitionCompensationChangeDataType
	var overlay struct {
		*T
		CompensationChangeDate *xsdDate `xml:"urn:com.workday/bsvc Compensation_Change_Date"`
	}
	overlay.T = (*T)(t)
	overlay.CompensationChangeDate = (*xsdDate)(&overlay.T.CompensationChangeDate)
	return d.DecodeElement(&overlay, &start)
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

// Contains a unique identifier for an instance of an object.
type RetentionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RetentionObjectType struct {
	ID         []RetentionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Request for getting back Employee Severance Worksheets
type RetrieveSeveranceWorksheetsRequestType struct {
	RequestReferences *EmployeeSeveranceWorksheetEventRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *EmployeeSeveranceWorksheetEventRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *EmployeeSeveranceWorksheetEventResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response element for Retrieve Employee Severance Worksheet Events
type RetrieveSeveranceWorksheetsResponseType struct {
	RequestReferences *EmployeeSeveranceWorksheetEventRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *EmployeeSeveranceWorksheetEventRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *EmployeeSeveranceWorksheetEventResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                                  `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *EmployeeSeveranceWorksheetEventResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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
type ReviewRatingScaleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ReviewRatingScaleObjectType struct {
	ID         []ReviewRatingScaleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// This is a wrapper for Salary Plans.
type SalaryPlanDataType struct {
	ApplyFullTimeEquivalent *bool `xml:"urn:com.workday/bsvc Apply_Full_Time_Equivalent,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SalaryPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SalaryPlanObjectType struct {
	ID         []SalaryPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Contains scorecard data.
type ScorecardProfileDataType struct {
	ScorecardProfileID                  string                                        `xml:"urn:com.workday/bsvc Scorecard_Profile_ID,omitempty"`
	ScorecardProfileTargetRuleReference *ConditionRuleObjectType                      `xml:"urn:com.workday/bsvc Scorecard_Profile_Target__Rule_Reference"`
	ScorecardProfileGoalData            []ScorecardProfilePerformanceCriteriaDataType `xml:"urn:com.workday/bsvc Scorecard_Profile_Goal_Data"`
}

// Contains set of goals or criteria.
type ScorecardProfilePerformanceCriteriaDataType struct {
	ScorecardProfileGoalID          string  `xml:"urn:com.workday/bsvc Scorecard_Profile_Goal_ID,omitempty"`
	ScorecardProfileGoalName        string  `xml:"urn:com.workday/bsvc Scorecard_Profile_Goal_Name"`
	ScorecardProfileGoalDescription string  `xml:"urn:com.workday/bsvc Scorecard_Profile_Goal_Description,omitempty"`
	ScorecardProfileGoalWeight      float64 `xml:"urn:com.workday/bsvc Scorecard_Profile_Goal_Weight"`
}

// Contains the set of goals and goal values for this result.
type ScorecardResultDataType struct {
	GoalReference *PerformanceCriteriaObjectType `xml:"urn:com.workday/bsvc Goal_Reference"`
	Achievement   float64                        `xml:"urn:com.workday/bsvc Achievement,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ScoresetContainerObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type ScoresetContainerObjectType struct {
	ID         []ScoresetContainerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The service length specific duration severance matrix data element.
type ServiceDurationSeveranceMatrixDataType struct {
	LengthofServiceUnitReference            *CompensationPeriodObjectType                 `xml:"urn:com.workday/bsvc Length_of_Service_Unit_Reference"`
	ServiceDurationSeveranceMatrixEntryData []ServiceDurationSeveranceMatrixEntryDataType `xml:"urn:com.workday/bsvc Service_Duration_Severance_Matrix_Entry_Data,omitempty"`
}

// The data for the matrix entry
type ServiceDurationSeveranceMatrixEntryDataType struct {
	MinimumLengthofService   float64                  `xml:"urn:com.workday/bsvc Minimum_Length_of_Service,omitempty"`
	MaximumLengthofService   float64                  `xml:"urn:com.workday/bsvc Maximum_Length_of_Service,omitempty"`
	EligibilityRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	Duration                 float64                  `xml:"urn:com.workday/bsvc Duration,omitempty"`
	MinimumDuration          float64                  `xml:"urn:com.workday/bsvc Minimum_Duration,omitempty"`
	MaximumDuration          float64                  `xml:"urn:com.workday/bsvc Maximum_Duration,omitempty"`
}

// The Service Length Severance Matrix data element.
type ServiceLengthSeveranceMatrixDataType struct {
	LengthofServiceUnitReference          *CompensationPeriodObjectType               `xml:"urn:com.workday/bsvc Length_of_Service_Unit_Reference"`
	CompensationRoundingRuleReference     *CompensationRoundingRuleObjectType         `xml:"urn:com.workday/bsvc Compensation_Rounding_Rule_Reference"`
	MultiplierOrderReference              *BenefitMultiplierOrderObjectType           `xml:"urn:com.workday/bsvc Multiplier_Order_Reference"`
	ServiceLengthSeveranceMatrixEntryData []ServiceLengthSeveranceMatrixEntryDataType `xml:"urn:com.workday/bsvc Service_Length_Severance_Matrix_Entry_Data"`
}

// the data for the matrix entry.
type ServiceLengthSeveranceMatrixEntryDataType struct {
	MinimumLengthofService    float64                  `xml:"urn:com.workday/bsvc Minimum_Length_of_Service,omitempty"`
	MaximumLengthofService    float64                  `xml:"urn:com.workday/bsvc Maximum_Length_of_Service,omitempty"`
	EligibilityRuleReference  *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	LengthofServiceMultiplier float64                  `xml:"urn:com.workday/bsvc Length_of_Service_Multiplier,omitempty"`
	MinimumDuration           float64                  `xml:"urn:com.workday/bsvc Minimum_Duration,omitempty"`
	MaximumDuration           float64                  `xml:"urn:com.workday/bsvc Maximum_Duration,omitempty"`
}

// Severance Matrix Filter Options
type SeveranceMatrixFilterOptionsType struct {
	CompensationBasisRange         *bool `xml:"urn:com.workday/bsvc Compensation_Basis_Range,omitempty"`
	ServiceLengthBasisRange        *bool `xml:"urn:com.workday/bsvc Service_Length_Basis_Range,omitempty"`
	DurationMatrixEntryType        *bool `xml:"urn:com.workday/bsvc Duration_Matrix_Entry_Type,omitempty"`
	LengthofServiceMatrixEntryType *bool `xml:"urn:com.workday/bsvc Length_of_Service_Matrix_Entry_Type,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SeveranceMatrixabstractObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SeveranceMatrixabstractObjectType struct {
	ID         []SeveranceMatrixabstractObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element to input data for a severance package component on a worksheet. Fields for a component type, eligibility, and comment
type SeverancePackageComponentDataType struct {
	SeverancePackageComponentTypeReference *SeverancePackageComponentTypeObjectType `xml:"urn:com.workday/bsvc Severance_Package_Component_Type_Reference,omitempty"`
	Eligible                               *bool                                    `xml:"urn:com.workday/bsvc Eligible,omitempty"`
	Comments                               string                                   `xml:"urn:com.workday/bsvc Comments,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SeverancePackageComponentTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SeverancePackageComponentTypeObjectType struct {
	ID         []SeverancePackageComponentTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SeverancePackageObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SeverancePackageObjectType struct {
	ID         []SeverancePackageObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Subedit to enter Severance Pay Component Information. Fields entered here will override default amounts.
type SeverancePayComponentDataType struct {
	SeverancePackageComponentTypeReference *SeverancePackageComponentTypeObjectType `xml:"urn:com.workday/bsvc Severance_Package_Component_Type_Reference,omitempty"`
	Duration                               float64                                  `xml:"urn:com.workday/bsvc Duration,omitempty"`
	OverrideDuration                       *bool                                    `xml:"urn:com.workday/bsvc Override_Duration,omitempty"`
	PayRate                                float64                                  `xml:"urn:com.workday/bsvc Pay_Rate,omitempty"`
	OverridePayRate                        *bool                                    `xml:"urn:com.workday/bsvc Override_Pay_Rate,omitempty"`
	Total                                  float64                                  `xml:"urn:com.workday/bsvc Total,omitempty"`
	SeveranceSalaryPlanReference           *SalaryPlanObjectType                    `xml:"urn:com.workday/bsvc Severance_Salary_Plan_Reference,omitempty"`
	SeveranceComponentComment              string                                   `xml:"urn:com.workday/bsvc Severance_Component_Comment,omitempty"`
	CurrencyReference                      *CurrencyObjectType                      `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SeveranceServiceDateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SeveranceServiceDateObjectType struct {
	ID         []SeveranceServiceDateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StockDateRuleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StockDateRuleObjectType struct {
	ID         []StockDateRuleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StockGrantObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StockGrantObjectType struct {
	ID         []StockGrantObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// Stock Participation Rate Table Data
type StockParticipationRateTableData20Type struct {
	StockParticipationRateTableID        string                                       `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_ID,omitempty"`
	Inactive                             *bool                                        `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	Name                                 string                                       `xml:"urn:com.workday/bsvc Name,omitempty"`
	Description                          string                                       `xml:"urn:com.workday/bsvc Description,omitempty"`
	DefaultPercentageRate                float64                                      `xml:"urn:com.workday/bsvc Default_Percentage_Rate,omitempty"`
	StockParticipationRateTableEntryData []StockParticipationRateTableEntryData20Type `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_Entry_Data,omitempty"`
}

// Stock Participation Rate Table Entry Data
type StockParticipationRateTableEntryData20Type struct {
	Delete                     *bool                         `xml:"urn:com.workday/bsvc Delete,omitempty"`
	EntrySortOrder             string                        `xml:"urn:com.workday/bsvc Entry_Sort_Order"`
	ManagementLevelReference   []ManagementLevelObjectType   `xml:"urn:com.workday/bsvc Management_Level_Reference,omitempty"`
	JobProfileReference        []JobProfileObjectType        `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
	CountryReference           []CountryObjectType           `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CompensationGradeReference []CompensationGradeObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Reference,omitempty"`
	PercentageRate             float64                       `xml:"urn:com.workday/bsvc Percentage_Rate,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StockParticipationRateTableObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StockParticipationRateTableObjectType struct {
	ID         []StockParticipationRateTableObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Request Criteria
type StockParticipationRateTableRequestCriteriaType struct {
	IncludeInactive bool `xml:"urn:com.workday/bsvc Include_Inactive,attr,omitempty"`
}

// Request References
type StockParticipationRateTableRequestReferencesType struct {
	StockParticipationRateTableReference []StockParticipationRateTableObjectType `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_Reference"`
}

// Stock Participation Rate Table Response Data
type StockParticipationRateTableResponseDataType struct {
	StockParticipationRateTable []StockParticipationRateTableSubResponseDataType `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table,omitempty"`
}

// Response Group
type StockParticipationRateTableResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Stock Participation Rate Table Data
type StockParticipationRateTableSubResponseDataType struct {
	StockParticipationRateTableReference *StockParticipationRateTableObjectType  `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_Reference,omitempty"`
	StockParticipationRateTableData      []StockParticipationRateTableData20Type `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type StockPercentPlanObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type StockPercentPlanObjectType struct {
	ID         []StockPercentPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Data for Stock Amount Plan.
type StockPlanAmountDataType struct {
	RoundingRuleReference                 *CompensationRoundingRuleObjectType         `xml:"urn:com.workday/bsvc Rounding_Rule_Reference,omitempty"`
	UseTargetCurrency                     *bool                                       `xml:"urn:com.workday/bsvc Use_Target_Currency,omitempty"`
	TargetAmount                          float64                                     `xml:"urn:com.workday/bsvc Target_Amount,omitempty"`
	CurrencyReference                     *CurrencyObjectType                         `xml:"urn:com.workday/bsvc Currency_Reference"`
	GrantSplitReplacementData             []GrantTypeSplitReplacementDataType         `xml:"urn:com.workday/bsvc Grant_Split_Replacement_Data"`
	StockPlanAmountProfileReplacementData []StockPlanAmountProfileReplacementDataType `xml:"urn:com.workday/bsvc Stock_Plan_Amount_Profile_Replacement_Data,omitempty"`
}

// Profile data for Stock Amount Plan.
type StockPlanAmountProfileReplacementDataType struct {
	TargetAmount                          float64                                     `xml:"urn:com.workday/bsvc Target_Amount,omitempty"`
	CurrencyReference                     *CurrencyObjectType                         `xml:"urn:com.workday/bsvc Currency_Reference"`
	EligibilityRuleReference              *ConditionRuleObjectType                    `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	StockProfileGrantSplitReplacementData []StockProfileGrantSplitReplacementDataType `xml:"urn:com.workday/bsvc Stock_Profile_Grant_Split_Replacement_Data"`
}

// Data for Stock Plan.
type StockPlanDataType struct {
	AllowTargetOverride         *bool                         `xml:"urn:com.workday/bsvc Allow_Target_Override,omitempty"`
	CompensationMatrixReference *CompensationMatrixObjectType `xml:"urn:com.workday/bsvc Compensation_Matrix_Reference,omitempty"`
	UseasReferenceOnly          *bool                         `xml:"urn:com.workday/bsvc Use_as_Reference_Only,omitempty"`
	StockPlanUnitData           *StockPlanUnitDataType        `xml:"urn:com.workday/bsvc Stock_Plan_Unit_Data"`
	StockPlanPercentData        *StockPlanPercentDataType     `xml:"urn:com.workday/bsvc Stock_Plan_Percent_Data"`
	StockPlanAmountData         *StockPlanAmountDataType      `xml:"urn:com.workday/bsvc Stock_Plan_Amount_Data"`
	HideTarget                  *bool                         `xml:"urn:com.workday/bsvc Hide_Target,omitempty"`
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

// Data for Stock Percent Plan.
type StockPlanPercentDataType struct {
	RoundingRuleReference                         *CompensationRoundingRuleObjectType          `xml:"urn:com.workday/bsvc Rounding_Rule_Reference,omitempty"`
	TargetPercent                                 float64                                      `xml:"urn:com.workday/bsvc Target_Percent,omitempty"`
	CompensationBasisReference                    *CompensationBasisObjectType                 `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
	GrantSplitReplacementData                     []GrantTypeSplitReplacementDataType          `xml:"urn:com.workday/bsvc Grant_Split_Replacement_Data"`
	StockPlanPercentProfileDefaultReplacementData []StockPlanPercentProfileReplacementDataType `xml:"urn:com.workday/bsvc Stock_Plan_Percent_Profile_Default_Replacement_Data,omitempty"`
}

// Data for Stock Percent Plan.
type StockPlanPercentProfileReplacementDataType struct {
	TargetPercent                         float64                                     `xml:"urn:com.workday/bsvc Target_Percent,omitempty"`
	EligibilityRuleReference              *ConditionRuleObjectType                    `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	CompensationBasisReference            *CompensationBasisObjectType                `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
	StockProfileGrantSplitReplacementData []StockProfileGrantSplitReplacementDataType `xml:"urn:com.workday/bsvc Stock_Profile_Grant_Split_Replacement_Data"`
}

// Data for Stock Unit Plan.
type StockPlanUnitDataType struct {
	TargetShares                        float64                                   `xml:"urn:com.workday/bsvc Target_Shares,omitempty"`
	GrantSplitReplacementData           []GrantTypeSplitReplacementDataType       `xml:"urn:com.workday/bsvc Grant_Split_Replacement_Data"`
	StockPlanUnitProfileReplacementData []StockPlanUnitProfileReplacementDataType `xml:"urn:com.workday/bsvc Stock_Plan_Unit_Profile_Replacement_Data,omitempty"`
}

// Profile data for Stock Unit Plan.
type StockPlanUnitProfileReplacementDataType struct {
	TargetShares                          float64                                     `xml:"urn:com.workday/bsvc Target_Shares,omitempty"`
	EligibilityRuleReference              *ConditionRuleObjectType                    `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
	StockProfileGrantSplitReplacementData []StockProfileGrantSplitReplacementDataType `xml:"urn:com.workday/bsvc Stock_Profile_Grant_Split_Replacement_Data"`
}

// Grant Split Data for Stock Profile
type StockProfileGrantSplitReplacementDataType struct {
	StockGrantTypeReference       *StockGrantTypeObjectType       `xml:"urn:com.workday/bsvc Stock_Grant_Type_Reference"`
	StockVestingScheduleReference *StockVestingScheduleObjectType `xml:"urn:com.workday/bsvc Stock_Vesting_Schedule_Reference,omitempty"`
	StockDateRuleReference        *StockDateRuleObjectType        `xml:"urn:com.workday/bsvc Stock_Date_Rule_Reference,omitempty"`
	GrantSplitPercent             float64                         `xml:"urn:com.workday/bsvc Grant_Split_Percent,omitempty"`
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
type TimeProrationRuleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TimeProrationRuleObjectType struct {
	ID         []TimeProrationRuleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
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

// This is a wrapper for the Base Pay Unit Salaries.
type UnitSalaryPlanDataType struct {
	PerUnitAmount          float64                  `xml:"urn:com.workday/bsvc Per_Unit_Amount,omitempty"`
	DefaultUnits           float64                  `xml:"urn:com.workday/bsvc Default_Units,omitempty"`
	UnitOfMeasureReference *UnitofMeasureObjectType `xml:"urn:com.workday/bsvc Unit_Of_Measure_Reference"`
	NoIndividualOverride   *bool                    `xml:"urn:com.workday/bsvc No_Individual_Override,omitempty"`
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

// Stock Grant data to be updated.
type UpdateStockGrantDataType struct {
	GrantID                     string              `xml:"urn:com.workday/bsvc Grant_ID,omitempty"`
	GrantDate                   *time.Time          `xml:"urn:com.workday/bsvc Grant_Date,omitempty"`
	GrantPrice                  float64             `xml:"urn:com.workday/bsvc Grant_Price,omitempty"`
	GrantPriceCurrencyReference *CurrencyObjectType `xml:"urn:com.workday/bsvc Grant_Price_Currency_Reference,omitempty"`
	BoardApproved               *bool               `xml:"urn:com.workday/bsvc Board_Approved,omitempty"`
	VestFromDate                *time.Time          `xml:"urn:com.workday/bsvc Vest_From_Date,omitempty"`
	ExpirationDate              *time.Time          `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
	OptionsPricingFactor        float64             `xml:"urn:com.workday/bsvc Options_Pricing_Factor,omitempty"`
	SharesVested                float64             `xml:"urn:com.workday/bsvc Shares_Vested,omitempty"`
	SharesUnvested              float64             `xml:"urn:com.workday/bsvc Shares_Unvested,omitempty"`
	VestingPrice                float64             `xml:"urn:com.workday/bsvc Vesting_Price,omitempty"`
	VestedAsOf                  *time.Time          `xml:"urn:com.workday/bsvc Vested_As_Of,omitempty"`
	LongTermCashAmountVested    float64             `xml:"urn:com.workday/bsvc Long_Term_Cash_Amount_Vested,omitempty"`
	LongTermCashAmountUnvested  float64             `xml:"urn:com.workday/bsvc Long_Term_Cash_Amount_Unvested,omitempty"`
}

func (t *UpdateStockGrantDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T UpdateStockGrantDataType
	var layout struct {
		*T
		GrantDate      *xsdDate `xml:"urn:com.workday/bsvc Grant_Date,omitempty"`
		VestFromDate   *xsdDate `xml:"urn:com.workday/bsvc Vest_From_Date,omitempty"`
		ExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
		VestedAsOf     *xsdDate `xml:"urn:com.workday/bsvc Vested_As_Of,omitempty"`
	}
	layout.T = (*T)(t)
	layout.GrantDate = (*xsdDate)(layout.T.GrantDate)
	layout.VestFromDate = (*xsdDate)(layout.T.VestFromDate)
	layout.ExpirationDate = (*xsdDate)(layout.T.ExpirationDate)
	layout.VestedAsOf = (*xsdDate)(layout.T.VestedAsOf)
	return e.EncodeElement(layout, start)
}
func (t *UpdateStockGrantDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T UpdateStockGrantDataType
	var overlay struct {
		*T
		GrantDate      *xsdDate `xml:"urn:com.workday/bsvc Grant_Date,omitempty"`
		VestFromDate   *xsdDate `xml:"urn:com.workday/bsvc Vest_From_Date,omitempty"`
		ExpirationDate *xsdDate `xml:"urn:com.workday/bsvc Expiration_Date,omitempty"`
		VestedAsOf     *xsdDate `xml:"urn:com.workday/bsvc Vested_As_Of,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.GrantDate = (*xsdDate)(overlay.T.GrantDate)
	overlay.VestFromDate = (*xsdDate)(overlay.T.VestFromDate)
	overlay.ExpirationDate = (*xsdDate)(overlay.T.ExpirationDate)
	overlay.VestedAsOf = (*xsdDate)(overlay.T.VestedAsOf)
	return d.DecodeElement(&overlay, &start)
}

// Stock grant request
type UpdateStockGrantRequestType struct {
	StockGrantReference *StockGrantObjectType     `xml:"urn:com.workday/bsvc Stock_Grant_Reference"`
	StockGrantData      *UpdateStockGrantDataType `xml:"urn:com.workday/bsvc Stock_Grant_Data"`
	Version             string                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Update Stock Grant Response
type UpdateStockGrantResponseType struct {
	StockGrantReference *StockGrantObjectType `xml:"urn:com.workday/bsvc Stock_Grant_Reference,omitempty"`
	Version             string                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
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

// Utilize the Request References element to retrieve a specific instance(s) of Worker and its associated data.
type WorkerRequestReferencesType struct {
	WorkerReference          []WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference"`
	SkipNonExistingInstances bool               `xml:"urn:com.workday/bsvc Skip_Non_Existing_Instances,attr,omitempty"`
	IgnoreInvalidReferences  bool               `xml:"urn:com.workday/bsvc Ignore_Invalid_References,attr,omitempty"`
}

// The response group allows for the response data to be tailored to only included elements that the user is looking for.  If no response group is provided in the request, then all groups will be returned. If the Response Group element is returned, you can select which sections of data to include in the response.
type WorkerResponseGroupforReferenceType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
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
