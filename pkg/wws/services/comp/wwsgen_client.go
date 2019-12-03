// Package comp
//
// The Compensation Web Service contains operations that expose Workday Human Capital Management Business Services compensation-related data.
package comp

import (
	"context"

	"github.com/ianlopshire/go-workday/pkg/wws"
)

const (
	CurrentVersion = "v31.2"
	ServiceName    = "Compensation"
)

type Client struct {
	*wws.Client
}

// PutPreviousSystemCompensationHistory (Put_Previous_System_Compensation_History)
//
// This operation will load previous system compensation history for a given employee. The operation allows free-form text entry indication of changes to a workers compensation before the implementation in the Workday system.
func (c *Client) PutPreviousSystemCompensationHistory(ctx context.Context, input *PutPreviousSystemCompensationHistoryInput) (output *PutPreviousSystemCompensationHistoryOutput, err error) {
	output = &PutPreviousSystemCompensationHistoryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPreviousSystemCompensationHistoryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Previous_System_Compensation_History_Request"`
	PutPreviousSystemCompensationHistoryRequestType
}

type PutPreviousSystemCompensationHistoryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Previous_System_Compensation_History_Response"`
	PutPreviousSystemCompensationHistoryResponseType
}

// GetPreviousSystemCompensationHistory (Get_Previous_System_Compensation_History)
//
// This operation will return any previous system compensation history which has been imported for a given employee.
func (c *Client) GetPreviousSystemCompensationHistory(ctx context.Context, input *GetPreviousSystemCompensationHistoryInput) (output *GetPreviousSystemCompensationHistoryOutput, err error) {
	output = &GetPreviousSystemCompensationHistoryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPreviousSystemCompensationHistoryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Previous_System_Compensation_History_Request"`
	GetPreviousSystemCompensationHistoryRequestType
}

type GetPreviousSystemCompensationHistoryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Previous_System_Compensation_History_Response"`
	GetPreviousSystemCompensationHistoryResponseType
}

// RequestBonusPayment (Request_Bonus_Payment)
//
// This operation allows the requesting of a bonus payment for an employee via the Request Bonus Payment business process.
func (c *Client) RequestBonusPayment(ctx context.Context, input *RequestBonusPaymentInput) (output *RequestBonusPaymentOutput, err error) {
	output = &RequestBonusPaymentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type RequestBonusPaymentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Request_Bonus_Payment_Request"`
	RequestBonusPaymentRequestType
}

type RequestBonusPaymentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Request_Bonus_Payment_Response"`
	RequestBonusPaymentResponseType
}

// RequestEmployeeMeritAdjustment (Request_Employee_Merit_Adjustment)
//
// This operation allows the requesting of a merit adjustment for an employee via the Request Employee Merit Adjustment business process.
func (c *Client) RequestEmployeeMeritAdjustment(ctx context.Context, input *RequestEmployeeMeritAdjustmentInput) (output *RequestEmployeeMeritAdjustmentOutput, err error) {
	output = &RequestEmployeeMeritAdjustmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type RequestEmployeeMeritAdjustmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Request_Employee_Merit_Adjustment_Request"`
	RequestEmployeeMeritAdjustmentRequestType
}

type RequestEmployeeMeritAdjustmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Request_Employee_Merit_Adjustment_Response"`
	RequestEmployeeMeritAdjustmentResponseType
}

// RequestOneTimePayment (Request_One-Time_Payment)
//
// This operation allows the requesting of a one-time payment for an employee via the Request One-Time Payment business process.
func (c *Client) RequestOneTimePayment(ctx context.Context, input *RequestOneTimePaymentInput) (output *RequestOneTimePaymentOutput, err error) {
	output = &RequestOneTimePaymentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type RequestOneTimePaymentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Request_One-Time_Payment_Request"`
	RequestOneTimePaymentRequestType
}

type RequestOneTimePaymentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Request_One-Time_Payment_Response"`
	RequestOneTimePaymentResponseType
}

// RequestCompensationChange (Request_Compensation_Change)
//
// This operation allows the requesting of a compensation change for an employee via the Request Compensation Change business process.
func (c *Client) RequestCompensationChange(ctx context.Context, input *RequestCompensationChangeInput) (output *RequestCompensationChangeOutput, err error) {
	output = &RequestCompensationChangeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type RequestCompensationChangeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Request_Compensation_Change_Request"`
	RequestCompensationChangeRequestType
}

type RequestCompensationChangeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Request_Compensation_Change_Response"`
	RequestCompensationChangeResponseType
}

// PutEligibleEarnings (Put_Eligible_Earnings)
//
// Adds or updates eligible earnings override information for an employee. This operation doesn’t follow participation rules. Use Import Eligible Earnings instead to follow those rules. This operation might affect in-progress bonus reviews.
func (c *Client) PutEligibleEarnings(ctx context.Context, input *PutEligibleEarningsInput) (output *PutEligibleEarningsOutput, err error) {
	output = &PutEligibleEarningsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutEligibleEarningsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Eligible_Earnings_Request"`
	PutEligibleEarningsRequestType
}

type PutEligibleEarningsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Eligible_Earnings_Response"`
	PutEligibleEarningsResponseType
}

// GetEligibleEarnings (Get_Eligible_Earnings)
//
// This operation retrieves the eligible earnings override values that are associated with a given employee.
func (c *Client) GetEligibleEarnings(ctx context.Context, input *GetEligibleEarningsInput) (output *GetEligibleEarningsOutput, err error) {
	output = &GetEligibleEarningsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEligibleEarningsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Eligible_Earnings_Request"`
	GetEligibleEarningsRequestType
}

type GetEligibleEarningsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Eligible_Earnings_Response"`
	GetEligibleEarningsResponseType
}

// GetCompensationMatrices (Get_Compensation_Matrices)
//
// This operation retrieves Merit Increase Matrices (MIMs) information for use by merit or bonus plans.
func (c *Client) GetCompensationMatrices(ctx context.Context, input *GetCompensationMatricesInput) (output *GetCompensationMatricesOutput, err error) {
	output = &GetCompensationMatricesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCompensationMatricesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Matrices_Request"`
	GetCompensationMatricesRequestType
}

type GetCompensationMatricesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Matrices_Response"`
	GetCompensationMatricesResponseType
}

// PutCompensationMatrix (Put_Compensation_Matrix)
//
// This operation adds or updates compensation matrix information.
func (c *Client) PutCompensationMatrix(ctx context.Context, input *PutCompensationMatrixInput) (output *PutCompensationMatrixOutput, err error) {
	output = &PutCompensationMatrixOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCompensationMatrixInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Matrix_Request"`
	PutCompensationMatrixRequestType
}

type PutCompensationMatrixOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Matrix_Response"`
	PutCompensationMatrixResponseType
}

// GetCompensationGrades (Get_Compensation_Grades)
//
// This operation allows the retrieval of compensation grade and grade profile information.  You must have access to the &#34;Set Up: Compensation Packages&#34; domain to access this operation.
func (c *Client) GetCompensationGrades(ctx context.Context, input *GetCompensationGradesInput) (output *GetCompensationGradesOutput, err error) {
	output = &GetCompensationGradesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCompensationGradesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Grades_Request"`
	GetCompensationGradesRequestType
}

type GetCompensationGradesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Grades_Response"`
	GetCompensationGradesResponseType
}

// PutCompensationGrade (Put_Compensation_Grade)
//
// This operation adds or updates compensation grade and compensation grade profile information.  You must have access to the &#34;Set Up: Compensation Packages&#34; domain to access this operation.
func (c *Client) PutCompensationGrade(ctx context.Context, input *PutCompensationGradeInput) (output *PutCompensationGradeOutput, err error) {
	output = &PutCompensationGradeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCompensationGradeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Grade_Request"`
	PutCompensationGradeRequestType
}

type PutCompensationGradeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Grade_Response"`
	PutCompensationGradeResponseType
}

// GetCompensationPlans (Get_Compensation_Plans)
//
// This operation allows the retrieval of detailed compensation plan information.
func (c *Client) GetCompensationPlans(ctx context.Context, input *GetCompensationPlansInput) (output *GetCompensationPlansOutput, err error) {
	output = &GetCompensationPlansOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCompensationPlansInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Plans_Request"`
	GetCompensationPlansRequestType
}

type GetCompensationPlansOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Plans_Response"`
	GetCompensationPlansResponseType
}

// PutCompensationPlans (Put_Compensation_Plans)
//
// This operation supports the addition or modification of detailed compensation plan information.
func (c *Client) PutCompensationPlans(ctx context.Context, input *PutCompensationPlansInput) (output *PutCompensationPlansOutput, err error) {
	output = &PutCompensationPlansOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCompensationPlansInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Plan_Request"`
	PutCompensationPlanRequestType
}

type PutCompensationPlansOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Plan_Response"`
	PutCompensationPlanResponseType
}

// GetFuturePaymentPlanAssignments (Get_Future_Payment_Plan_Assignments)
//
// This operation allows the retrieval of future payment plan assignments for one or more employees. You must have access to the Worker:Compensation by Organization domain.
func (c *Client) GetFuturePaymentPlanAssignments(ctx context.Context, input *GetFuturePaymentPlanAssignmentsInput) (output *GetFuturePaymentPlanAssignmentsOutput, err error) {
	output = &GetFuturePaymentPlanAssignmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetFuturePaymentPlanAssignmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Future_Payment_Plan_Assignments_Request"`
	GetFuturePaymentPlanAssignmentsRequestType
}

type GetFuturePaymentPlanAssignmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Future_Payment_Plan_Assignments_Response"`
	GetFuturePaymentPlanAssignmentsResponseType
}

// PutFuturePaymentPlanAssignment (Put_Future_Payment_Plan_Assignment)
//
// This operation adds or updates future payment plan assignments for a given employee. You must have access to the Worker:Compensation by Organization domain.
func (c *Client) PutFuturePaymentPlanAssignment(ctx context.Context, input *PutFuturePaymentPlanAssignmentInput) (output *PutFuturePaymentPlanAssignmentOutput, err error) {
	output = &PutFuturePaymentPlanAssignmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutFuturePaymentPlanAssignmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Future_Payment_Plan_Assignment_Request"`
	PutFuturePaymentPlanAssignmentRequestType
}

type PutFuturePaymentPlanAssignmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Future_Payment_Plan_Assignment_Response"`
	PutFuturePaymentPlanAssignmentResponseType
}

// GetCompensationScorecards (Get_Compensation_Scorecards)
//
// This operation allows the retrieval of one or more scorecards.
func (c *Client) GetCompensationScorecards(ctx context.Context, input *GetCompensationScorecardsInput) (output *GetCompensationScorecardsOutput, err error) {
	output = &GetCompensationScorecardsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCompensationScorecardsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Scorecards_Request"`
	GetCompensationScorecardsRequestType
}

type GetCompensationScorecardsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Scorecards_Response"`
	GetCompensationScorecardsResponseType
}

// PutCompensationScorecard (Put_Compensation_Scorecard)
//
// This operation allows the creation or modification of a scorecard.
func (c *Client) PutCompensationScorecard(ctx context.Context, input *PutCompensationScorecardInput) (output *PutCompensationScorecardOutput, err error) {
	output = &PutCompensationScorecardOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCompensationScorecardInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Scorecard_Request"`
	PutCompensationScorecardRequestType
}

type PutCompensationScorecardOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Scorecard_Response"`
	PutCompensationScorecardResponseType
}

// PutBenchmarkJob (Put_Benchmark_Job)
//
// This operation allows the addition and updating of benchmark job information.
func (c *Client) PutBenchmarkJob(ctx context.Context, input *PutBenchmarkJobInput) (output *PutBenchmarkJobOutput, err error) {
	output = &PutBenchmarkJobOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBenchmarkJobInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Benchmark_Job_Request"`
	PutBenchmarkJobRequestType
}

type PutBenchmarkJobOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Benchmark_Job_Response"`
	PutBenchmarkJobResponseType
}

// GetBenchmarkJobs (Get_Benchmark_Jobs)
//
// This operation allows the retrieval of benchmark job information.
func (c *Client) GetBenchmarkJobs(ctx context.Context, input *GetBenchmarkJobsInput) (output *GetBenchmarkJobsOutput, err error) {
	output = &GetBenchmarkJobsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBenchmarkJobsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Benchmark_Jobs_Request"`
	GetBenchmarkJobsRequestType
}

type GetBenchmarkJobsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Benchmark_Jobs_Response"`
	GetBenchmarkJobsResponseType
}

// UpdateStockGrant (Update_Stock_Grant)
//
// This operation allows the updating of an existing stock grant which has been given to a worker.
func (c *Client) UpdateStockGrant(ctx context.Context, input *UpdateStockGrantInput) (output *UpdateStockGrantOutput, err error) {
	output = &UpdateStockGrantOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type UpdateStockGrantInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Update_Stock_Grant_Request"`
	UpdateStockGrantRequestType
}

type UpdateStockGrantOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Update_Stock_Grant_Response"`
	UpdateStockGrantResponseType
}

// AddStockGrant (Add_Stock_Grant)
//
// This operation allows the addition of stock grants to an employee via the Request Stock Option Grant business process.
func (c *Client) AddStockGrant(ctx context.Context, input *AddStockGrantInput) (output *AddStockGrantOutput, err error) {
	output = &AddStockGrantOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AddStockGrantInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Add_Stock_Grant_Request"`
	AddStockGrantRequestType
}

type AddStockGrantOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Add_Stock_Grant_Response"`
	AddStockGrantResponseType
}

// PutCompensationEligibilityRule (Put_Compensation_Eligibility_Rule)
//
// This operation allows for the creation or update of a Compensation Eligibility Rule.
func (c *Client) PutCompensationEligibilityRule(ctx context.Context, input *PutCompensationEligibilityRuleInput) (output *PutCompensationEligibilityRuleOutput, err error) {
	output = &PutCompensationEligibilityRuleOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCompensationEligibilityRuleInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Eligibility_Rule_Request"`
	PutCompensationEligibilityRuleRequestType
}

type PutCompensationEligibilityRuleOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Eligibility_Rule_Response"`
	PutCompensationEligibilityRuleResponseType
}

// GetCompensationEligibilityRules (Get_Compensation_Eligibility_Rules)
//
// This operation returns Compensation Eligibility Rules
func (c *Client) GetCompensationEligibilityRules(ctx context.Context, input *GetCompensationEligibilityRulesInput) (output *GetCompensationEligibilityRulesOutput, err error) {
	output = &GetCompensationEligibilityRulesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCompensationEligibilityRulesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Eligibility_Rules_Request"`
	GetCompensationEligibilityRulesRequestType
}

type GetCompensationEligibilityRulesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Eligibility_Rules_Response"`
	GetCompensationEligibilityRulesResponseType
}

// GetCompensationEligibilityRuleswithoutDependencies (Get_Compensation_Eligibility_Rules_without_Dependencies)
//
// This operation returns Compensation Eligibility Rules without dependent Compensation Eligibility Rule references.
func (c *Client) GetCompensationEligibilityRuleswithoutDependencies(ctx context.Context, input *GetCompensationEligibilityRuleswithoutDependenciesInput) (output *GetCompensationEligibilityRuleswithoutDependenciesOutput, err error) {
	output = &GetCompensationEligibilityRuleswithoutDependenciesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCompensationEligibilityRuleswithoutDependenciesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Eligibility_Rules_without_Dependencies_Request"`
	GetCompensationEligibilityRuleswithoutDependenciesRequestType
}

type GetCompensationEligibilityRuleswithoutDependenciesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Eligibility_Rules_Response"`
	GetCompensationEligibilityRulesResponseType
}

// PutCompensationScorecardResult (Put_Compensation_Scorecard_Result)
//
// This operation allows the creation or modification of a scorecard result against a specific scorecard
func (c *Client) PutCompensationScorecardResult(ctx context.Context, input *PutCompensationScorecardResultInput) (output *PutCompensationScorecardResultOutput, err error) {
	output = &PutCompensationScorecardResultOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCompensationScorecardResultInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Scorecard_Result_Request"`
	PutCompensationScorecardResultRequestType
}

type PutCompensationScorecardResultOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Scorecard_Result_Response"`
	PutCompensationScorecardResultResponseType
}

// GetCompensationScorecardResults (Get_Compensation_Scorecard_Results)
//
// This operation allows the retrieval of one or more scorecard results
func (c *Client) GetCompensationScorecardResults(ctx context.Context, input *GetCompensationScorecardResultsInput) (output *GetCompensationScorecardResultsOutput, err error) {
	output = &GetCompensationScorecardResultsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCompensationScorecardResultsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Scorecard_Results_Request"`
	GetCompensationScorecardResultsRequestType
}

type GetCompensationScorecardResultsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Scorecard_Results_Response"`
	GetCompensationScorecardResultsResponseType
}

// CreateSeveranceWorksheet (Create_Severance_Worksheet)
//
// This operation allows the launching of an employee severance worksheet via the Severance Worksheet business process.
func (c *Client) CreateSeveranceWorksheet(ctx context.Context, input *CreateSeveranceWorksheetInput) (output *CreateSeveranceWorksheetOutput, err error) {
	output = &CreateSeveranceWorksheetOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CreateSeveranceWorksheetInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Create_Severance_Worksheet_Request"`
	CreateSeveranceWorksheetRequestType
}

type CreateSeveranceWorksheetOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Create_Severance_Worksheet_Response"`
	CreateSeveranceWorksheetResponseType
}

// GetCompensationGradeHierarchy (Get_Compensation_Grade_Hierarchy)
//
// This operation allows the retrieval of compensation grade hierarchy information. You must have access to the &#34;Set Up: Compensation General&#34; domain to access this operation.
func (c *Client) GetCompensationGradeHierarchy(ctx context.Context, input *GetCompensationGradeHierarchyInput) (output *GetCompensationGradeHierarchyOutput, err error) {
	output = &GetCompensationGradeHierarchyOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCompensationGradeHierarchyInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Grade_Hierarchy_Request"`
	GetCompensationGradeHierarchyRequestType
}

type GetCompensationGradeHierarchyOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Grade_Hierarchy_Response"`
	GetCompensationGradeHierarchyResponseType
}

// PutCompensationGradeHierarchy (Put_Compensation_Grade_Hierarchy)
//
// This operation adds or updates compensation grade hierarchy You must have access to the &#34;Set Up: Compensation General&#34; domain to access this operation.
func (c *Client) PutCompensationGradeHierarchy(ctx context.Context, input *PutCompensationGradeHierarchyInput) (output *PutCompensationGradeHierarchyOutput, err error) {
	output = &PutCompensationGradeHierarchyOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCompensationGradeHierarchyInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Grade_Hierarchy_Request"`
	PutCompensationGradeHierarchyRequestType
}

type PutCompensationGradeHierarchyOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Grade_Hierarchy_Response"`
	PutCompensationGradeHierarchyResponseType
}

// GetPeriodActivityTasks (Get_Period_Activity_Tasks)
//
// This operation allows the retrieval of period activity tasks.  You must have access to the &#34;Set Up: Work Activity for Pay&#34; domain to access this operation.
func (c *Client) GetPeriodActivityTasks(ctx context.Context, input *GetPeriodActivityTasksInput) (output *GetPeriodActivityTasksOutput, err error) {
	output = &GetPeriodActivityTasksOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPeriodActivityTasksInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Period_Activity_Tasks_Request"`
	GetPeriodActivityTasksRequestType
}

type GetPeriodActivityTasksOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Period_Activity_Tasks_Response"`
	GetPeriodActivityTasksResponseType
}

// PutPeriodActivityTask (Put_Period_Activity_Task)
//
// This operation adds or updates a period activity task.  You must have access to the &#34;Set Up: Work Activity for Pay&#34; domain to access this operation.
func (c *Client) PutPeriodActivityTask(ctx context.Context, input *PutPeriodActivityTaskInput) (output *PutPeriodActivityTaskOutput, err error) {
	output = &PutPeriodActivityTaskOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPeriodActivityTaskInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Period_Activity_Task_Request"`
	PutPeriodActivityTaskRequestType
}

type PutPeriodActivityTaskOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Period_Activity_Task_Response"`
	PutPeriodActivityTaskResponseType
}

// AssignEligiblePeriodActivitiesforEmployee (Assign_Eligible_Period_Activities_for_Employee)
//
// This operation allows the requesting of a period activity assignment for an employee via the Assign Eligible Period Activities for Employee business process.
func (c *Client) AssignEligiblePeriodActivitiesforEmployee(ctx context.Context, input *AssignEligiblePeriodActivitiesforEmployeeInput) (output *AssignEligiblePeriodActivitiesforEmployeeOutput, err error) {
	output = &AssignEligiblePeriodActivitiesforEmployeeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AssignEligiblePeriodActivitiesforEmployeeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Eligible_Period_Activities_for_Employee_Request"`
	AssignEligiblePeriodActivitiesforEmployeeRequestType
}

type AssignEligiblePeriodActivitiesforEmployeeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Eligible_Period_Activities_for_Employee_Response"`
	AssignEligiblePeriodActivitiesforEmployeeResponseType
}

// PutPeriodActivityRateMatrix (Put_Period_Activity_Rate_Matrix)
//
// This operation creates or updates a Period Activity Rate Matrix
func (c *Client) PutPeriodActivityRateMatrix(ctx context.Context, input *PutPeriodActivityRateMatrixInput) (output *PutPeriodActivityRateMatrixOutput, err error) {
	output = &PutPeriodActivityRateMatrixOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPeriodActivityRateMatrixInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Period_Activity_Rate_Matrix_Request"`
	PutPeriodActivityRateMatrixRequestType
}

type PutPeriodActivityRateMatrixOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Period_Activity_Rate_Matrix_Response"`
	PutPeriodActivityRateMatrixResponseType
}

// GetPeriodActivityRateMatrices (Get_Period_Activity_Rate_Matrices)
//
// This operation returns the Period Activity Rate Matrices created
func (c *Client) GetPeriodActivityRateMatrices(ctx context.Context, input *GetPeriodActivityRateMatricesInput) (output *GetPeriodActivityRateMatricesOutput, err error) {
	output = &GetPeriodActivityRateMatricesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPeriodActivityRateMatricesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Period_Activity_Rate_Matrices_Request"`
	GetPeriodActivityRateMatricesRequestType
}

type GetPeriodActivityRateMatricesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Period_Activity_Rate_Matrices_Response"`
	GetPeriodActivityRateMatricesResponseType
}

// ManagePeriodActivityPayAssignments (Manage_Period_Activity_Pay_Assignments)
//
// This operation allows the management of period activity based assignments for a given worker and a bounding date range or period via the Manage Period Activity Pay Assignments business process.
func (c *Client) ManagePeriodActivityPayAssignments(ctx context.Context, input *ManagePeriodActivityPayAssignmentsInput) (output *ManagePeriodActivityPayAssignmentsOutput, err error) {
	output = &ManagePeriodActivityPayAssignmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ManagePeriodActivityPayAssignmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Period_Activity_Pay_Assignments_Request"`
	ManagePeriodActivityPayAssignmentsRequestType
}

type ManagePeriodActivityPayAssignmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Period_Activity_Pay_Assignments_Response"`
	ManagePeriodActivityPayAssignmentsResponseType
}

// PutStockParticipationRateTable (Put_Stock_Participation_Rate_Table)
//
// This operation allows you to create, update, or delete a Stock Participation Rate Table.
func (c *Client) PutStockParticipationRateTable(ctx context.Context, input *PutStockParticipationRateTableInput) (output *PutStockParticipationRateTableOutput, err error) {
	output = &PutStockParticipationRateTableOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutStockParticipationRateTableInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Stock_Participation_Rate_Table_Request"`
	PutStockParticipationRateTableRequestType
}

type PutStockParticipationRateTableOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Stock_Participation_Rate_Table_Response"`
	PutStockParticipationRateTableResponseType
}

// GetStockParticipationRateTables (Get_Stock_Participation_Rate_Tables)
//
// This operation allows you to retrieve Stock Participation Rate Table information.
func (c *Client) GetStockParticipationRateTables(ctx context.Context, input *GetStockParticipationRateTablesInput) (output *GetStockParticipationRateTablesOutput, err error) {
	output = &GetStockParticipationRateTablesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetStockParticipationRateTablesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Stock_Participation_Rate_Tables_Request"`
	GetStockParticipationRateTablesRequestType
}

type GetStockParticipationRateTablesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Stock_Participation_Rate_Tables_Response"`
	GetStockParticipationRateTablesResponseType
}

// GetOneTimePaymentPlanConfigurableCategories (Get_One-Time_Payment_Plan_Configurable_Categories)
//
// Web Service for Get One-Time Payment Plan Configurable Categories
func (c *Client) GetOneTimePaymentPlanConfigurableCategories(ctx context.Context, input *GetOneTimePaymentPlanConfigurableCategoriesInput) (output *GetOneTimePaymentPlanConfigurableCategoriesOutput, err error) {
	output = &GetOneTimePaymentPlanConfigurableCategoriesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetOneTimePaymentPlanConfigurableCategoriesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_One-Time_Payment_Plan_Configurable_Categories_Request"`
	GetOneTimePaymentPlanConfigurableCategoriesRequestType
}

type GetOneTimePaymentPlanConfigurableCategoriesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_One-Time_Payment_Plan_Configurable_Categories_Response"`
	GetOneTimePaymentPlanConfigurableCategoriesResponseType
}

// PutOneTimePaymentPlanConfigurableCategory (Put_One-Time_Payment_Plan_Configurable_Category)
//
// Web Service for One-Time Payment Plan Configurable Category
func (c *Client) PutOneTimePaymentPlanConfigurableCategory(ctx context.Context, input *PutOneTimePaymentPlanConfigurableCategoryInput) (output *PutOneTimePaymentPlanConfigurableCategoryOutput, err error) {
	output = &PutOneTimePaymentPlanConfigurableCategoryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutOneTimePaymentPlanConfigurableCategoryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_One-Time_Payment_Plan_Configurable_Category_Request"`
	PutOneTimePaymentPlanConfigurableCategoryRequestType
}

type PutOneTimePaymentPlanConfigurableCategoryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_One-Time_Payment_Plan_Configurable_Category_Response"`
	PutOneTimePaymentPlanConfigurableCategoryResponseType
}

// RequestRequisitionCompensationChange (Request_Requisition_Compensation_Change)
//
// Web service operation to request an ad-hoc requisition compensation change for a job requisition.
func (c *Client) RequestRequisitionCompensationChange(ctx context.Context, input *RequestRequisitionCompensationChangeInput) (output *RequestRequisitionCompensationChangeOutput, err error) {
	output = &RequestRequisitionCompensationChangeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type RequestRequisitionCompensationChangeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Request_Requisition_Compensation_Change_Request"`
	RequestRequisitionCompensationChangeRequestType
}

type RequestRequisitionCompensationChangeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Request_Requisition_Compensation_Change_Response"`
	RequestRequisitionCompensationChangeResponseType
}

// RetrieveSeveranceWorksheet (Retrieve_Severance_Worksheet)
//
// Returns an Employee Severance Worksheet
func (c *Client) RetrieveSeveranceWorksheet(ctx context.Context, input *RetrieveSeveranceWorksheetInput) (output *RetrieveSeveranceWorksheetOutput, err error) {
	output = &RetrieveSeveranceWorksheetOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type RetrieveSeveranceWorksheetInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Retrieve_Severance_Worksheets_Request"`
	RetrieveSeveranceWorksheetsRequestType
}

type RetrieveSeveranceWorksheetOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Retrieve_Severance_Worksheets_Response"`
	RetrieveSeveranceWorksheetsResponseType
}

// PutEnhancedSeveranceMatrix (Put_Enhanced_Severance_Matrix)
//
// Create or modify data for an Enhanced Severance Matrix.
func (c *Client) PutEnhancedSeveranceMatrix(ctx context.Context, input *PutEnhancedSeveranceMatrixInput) (output *PutEnhancedSeveranceMatrixOutput, err error) {
	output = &PutEnhancedSeveranceMatrixOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutEnhancedSeveranceMatrixInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Enhanced_Severance_Matrix_Request"`
	PutEnhancedSeveranceMatrixRequestType
}

type PutEnhancedSeveranceMatrixOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Enhanced_Severance_Matrix_Response"`
	PutEnhancedSeveranceMatrixResponseType
}

// GetEnhancedSeveranceMatrices (Get_Enhanced_Severance_Matrices)
//
// This service retrieves the Enhanced Severance Matrices from the system.
func (c *Client) GetEnhancedSeveranceMatrices(ctx context.Context, input *GetEnhancedSeveranceMatricesInput) (output *GetEnhancedSeveranceMatricesOutput, err error) {
	output = &GetEnhancedSeveranceMatricesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEnhancedSeveranceMatricesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Enhanced_Severance_Matrices_Request"`
	GetEnhancedSeveranceMatricesRequestType
}

type GetEnhancedSeveranceMatricesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Enhanced_Severance_Matrices_Response"`
	GetEnhancedSeveranceMatricesResponseType
}

// ImportEligibleEarningsOverride (Import_Eligible_Earnings_Override)
//
// This operation adds or updates eligible earnings override information for a given employee. If there is an in-progress Bonus Process with configured Participation Rules, then that process will react to the eligible earnings in this operation.
func (c *Client) ImportEligibleEarningsOverride(ctx context.Context, input *ImportEligibleEarningsOverrideInput) (output *ImportEligibleEarningsOverrideOutput, err error) {
	output = &ImportEligibleEarningsOverrideOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportEligibleEarningsOverrideInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Eligible_Earnings_Override_Request"`
	ImportEligibleEarningsOverrideRequestType
}

type ImportEligibleEarningsOverrideOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportRequestCompensationChange (Import_Request_Compensation_Change)
//
// High-volume web service task to Request Compensation Change for employee(s)
func (c *Client) ImportRequestCompensationChange(ctx context.Context, input *ImportRequestCompensationChangeInput) (output *ImportRequestCompensationChangeOutput, err error) {
	output = &ImportRequestCompensationChangeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportRequestCompensationChangeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Request_Compensation_Change_Request"`
	ImportRequestCompensationChangeRequestType
}

type ImportRequestCompensationChangeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}
