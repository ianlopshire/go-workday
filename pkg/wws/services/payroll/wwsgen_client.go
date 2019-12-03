
// Package payroll
//
// The Workday Payroll Web Service contains operations that expose Workday Payroll data for integration with third parties (such as time and attendance).
package payroll

import (
	"context"

	"github.com/ianlopshire/go-workday/pkg/wws"
)

const (
	CurrentVersion = "v31.2"
	ServiceName = "Payroll"
)

type Client struct {
	*wws.Client
}


// GetPayrollResults (Get_Payroll_Results)
// 
// Return Payroll Results for selected Workers, Companies, Pay Calculations for specified Date Range.
func (c *Client) GetPayrollResults(ctx context.Context, input *GetPayrollResultsInput) (output *GetPayrollResultsOutput, err error) {
	output = &GetPayrollResultsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollResultsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Results_Request"`
	GetPayrollResultsRequestType
}

type GetPayrollResultsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Results_Response"`
	GetPayrollResultsResponseType
}

// GetPayrollBalances (Get_Payroll_Balances)
// 
// Return Payroll Balances for selected Workers, Companies, Pay Calculations for specified Balance Periods.
func (c *Client) GetPayrollBalances(ctx context.Context, input *GetPayrollBalancesInput) (output *GetPayrollBalancesOutput, err error) {
	output = &GetPayrollBalancesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollBalancesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Balances_Request"`
	GetPayrollBalancesRequestType
}

type GetPayrollBalancesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Balances_Response"`
	GetPayrollBalancesResponseType
}

// PutPayrollOffcyclePayment (Put_Payroll_Off-cycle_Payment)
// 
// This public web service operation is designed to add/update off cycle payments from external systems into Workday.
func (c *Client) PutPayrollOffcyclePayment(ctx context.Context, input *PutPayrollOffcyclePaymentInput) (output *PutPayrollOffcyclePaymentOutput, err error) {
	output = &PutPayrollOffcyclePaymentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollOffcyclePaymentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Off-cycle_Payment_Request"`
	PutPayrollOffcyclePaymentRequestType
}

type PutPayrollOffcyclePaymentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Off-cycle_Payment_Response"`
	PutPayrollOffcyclePaymentResponseType
}

// PutPayrollHistoryPayment (Put_Payroll_History_Payment)
// 
// This public web service operation is designed to add/update historical payroll payments from external systems into Workday Payroll.
func (c *Client) PutPayrollHistoryPayment(ctx context.Context, input *PutPayrollHistoryPaymentInput) (output *PutPayrollHistoryPaymentOutput, err error) {
	output = &PutPayrollHistoryPaymentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollHistoryPaymentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_History_Payment_Request"`
	PutPayrollHistoryPaymentRequestType
}

type PutPayrollHistoryPaymentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_History_Payment_Response"`
	PutPayrollHistoryPaymentResponseType
}

// SubmitPayrollInput (Submit_Payroll_Input)
// 
// This public web service operation is designed to add/update Payroll Input data from external systems into Workday.
func (c *Client) SubmitPayrollInput(ctx context.Context, input *SubmitPayrollInputInput) (output *SubmitPayrollInputOutput, err error) {
	output = &SubmitPayrollInputOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitPayrollInputInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Payroll_Input_Request"`
	SubmitPayrollInputRequestType
}

type SubmitPayrollInputOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Payroll_Input_Response"`
	SubmitPayrollInputResponseType
}

// GetPayrollHistoryPayments (Get_Payroll_History_Payments)
// 
// Retrieves Workday History Payments.  History Payments in Workday are historical payments imported from external and/or legacy payroll systems.
func (c *Client) GetPayrollHistoryPayments(ctx context.Context, input *GetPayrollHistoryPaymentsInput) (output *GetPayrollHistoryPaymentsOutput, err error) {
	output = &GetPayrollHistoryPaymentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollHistoryPaymentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_History_Payments_Request"`
	GetPayrollHistoryPaymentsRequestType
}

type GetPayrollHistoryPaymentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_History_Payments_Response"`
	GetPayrollHistoryPaymentsResponseType
}

// GetPeriodSchedules (Get_Period_Schedules)
// 
// Return Period Schedules
func (c *Client) GetPeriodSchedules(ctx context.Context, input *GetPeriodSchedulesInput) (output *GetPeriodSchedulesOutput, err error) {
	output = &GetPeriodSchedulesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPeriodSchedulesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Period_Schedules_Request"`
	GetPeriodSchedulesRequestType
}

type GetPeriodSchedulesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Period_Schedules_Response"`
	GetPeriodSchedulesResponseType
}

// PutPeriodSchedule (Put_Period_Schedule)
// 
// This public web service operation is designed to add/update Period Schedules
func (c *Client) PutPeriodSchedule(ctx context.Context, input *PutPeriodScheduleInput) (output *PutPeriodScheduleOutput, err error) {
	output = &PutPeriodScheduleOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPeriodScheduleInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Period_Schedule_Request"`
	PutPeriodScheduleRequestType
}

type PutPeriodScheduleOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Period_Schedule_Response"`
	PutPeriodScheduleResponseType
}

// GetSubmitPayrollInputs (Get_Submit_Payroll_Inputs)
// 
// This public web service operation is designed to retrieve Payroll Input data from the Workday system.
func (c *Client) GetSubmitPayrollInputs(ctx context.Context, input *GetSubmitPayrollInputsInput) (output *GetSubmitPayrollInputsOutput, err error) {
	output = &GetSubmitPayrollInputsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetSubmitPayrollInputsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Submit_Payroll_Inputs_Request"`
	GetSubmitPayrollInputsRequestType
}

type GetSubmitPayrollInputsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Submit_Payroll_Inputs_Response"`
	GetSubmitPayrollInputsResponseType
}

// AssignCostingAllocation (Assign_Costing_Allocation)
// 
// This service operation will submit requests to replace, add, update or delete costing allocations (distributions) for worker/position/earning and position restrictions criteria.  This service does not manage strictly earning-level costing overrides.
func (c *Client) AssignCostingAllocation(ctx context.Context, input *AssignCostingAllocationInput) (output *AssignCostingAllocationOutput, err error) {
	output = &AssignCostingAllocationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AssignCostingAllocationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Costing_Allocation_Request"`
	AssignCostingAllocationRequestType
}

type AssignCostingAllocationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Costing_Allocation_Response"`
	AssignCostingAllocationResponseType
}

// GetPayrollPayslips (Get_Payroll_Payslips)
// 
// Return payroll payslips for an outsourced payment group that was generated by a settlement run.
func (c *Client) GetPayrollPayslips(ctx context.Context, input *GetPayrollPayslipsInput) (output *GetPayrollPayslipsOutput, err error) {
	output = &GetPayrollPayslipsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollPayslipsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payslips_Request"`
	GetPayrollPayslipsRequestType
}

type GetPayrollPayslipsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payslips_Response"`
	GetPayrollPayslipsResponseType
}

// GetQuarterlyWorkerTaxFilingData (Get_Quarterly_Worker_Tax_Filing_Data)
// 
// This web service operation returns quarterly tax filing data based on specified criteria.
func (c *Client) GetQuarterlyWorkerTaxFilingData(ctx context.Context, input *GetQuarterlyWorkerTaxFilingDataInput) (output *GetQuarterlyWorkerTaxFilingDataOutput, err error) {
	output = &GetQuarterlyWorkerTaxFilingDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetQuarterlyWorkerTaxFilingDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Quarterly_Worker_Tax_Data_Request"`
	GetQuarterlyWorkerTaxDataRequestType
}

type GetQuarterlyWorkerTaxFilingDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Quarterly_Worker_Tax_Data_Response"`
	GetQuarterlyWorkerTaxDataResponseType
}

// PutPayrollPayeeFICA (Put_Payroll_Payee_FICA)
// 
// Public Web Service for updating Payroll Payee FICA Withholdings for specific worker.
func (c *Client) PutPayrollPayeeFICA(ctx context.Context, input *PutPayrollPayeeFICAInput) (output *PutPayrollPayeeFICAOutput, err error) {
	output = &PutPayrollPayeeFICAOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollPayeeFICAInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Payee_FICA_Request"`
	PutPayrollPayeeFICARequestType
}

type PutPayrollPayeeFICAOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Payee_FICA_Response"`
	PutPayrollPayeeFICAResponseType
}

// GetPeriodicWorkerTaxFilingData (Get_Periodic_Worker_Tax_Filing_Data)
// 
// This web service operation returns periodic tax filing data based on specified criteria.
func (c *Client) GetPeriodicWorkerTaxFilingData(ctx context.Context, input *GetPeriodicWorkerTaxFilingDataInput) (output *GetPeriodicWorkerTaxFilingDataOutput, err error) {
	output = &GetPeriodicWorkerTaxFilingDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPeriodicWorkerTaxFilingDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Periodic_Worker_Tax_Filing_Data_Request"`
	GetPeriodicWorkerTaxFilingDataRequestType
}

type GetPeriodicWorkerTaxFilingDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Periodic_Worker_Tax_Filing_Data_Response"`
	GetPeriodicWorkerTaxFilingDataResponseType
}

// GetPeriodicCompanyTaxFilingData (Get_Periodic_Company_Tax_Filing_Data)
// 
// This web service operation retrieves periodic tax filing data for a company.  Tax filing data is aggregated by Payment Date and Third Party Sick Pay indicator.  
// 
// Workday filters out $0 values in the web service response to improve performance, which can result in record counts being off.
func (c *Client) GetPeriodicCompanyTaxFilingData(ctx context.Context, input *GetPeriodicCompanyTaxFilingDataInput) (output *GetPeriodicCompanyTaxFilingDataOutput, err error) {
	output = &GetPeriodicCompanyTaxFilingDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPeriodicCompanyTaxFilingDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Periodic_Company_Tax_Filing_Data_Request"`
	GetPeriodicCompanyTaxFilingDataRequestType
}

type GetPeriodicCompanyTaxFilingDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Periodic_Company_Tax_Filing_Data_Response"`
	GetPeriodicCompanyTaxFilingDataResponseType
}

// GetPayrollPayeePT1s (Get_Payroll_Payee_PT1s)
// 
// This web service is used to return Payroll Payee PT1s
func (c *Client) GetPayrollPayeePT1s(ctx context.Context, input *GetPayrollPayeePT1sInput) (output *GetPayrollPayeePT1sOutput, err error) {
	output = &GetPayrollPayeePT1sOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollPayeePT1sInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payee_PT1s_Request"`
	GetPayrollPayeePT1sRequestType
}

type GetPayrollPayeePT1sOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payee_PT1s_Response"`
	GetPayrollPayeePT1sResponseType
}

// PutPayrollPayeePT1 (Put_Payroll_Payee_PT1)
// 
// This web service is used to load PT1&#39;s
func (c *Client) PutPayrollPayeePT1(ctx context.Context, input *PutPayrollPayeePT1Input) (output *PutPayrollPayeePT1Output, err error) {
	output = &PutPayrollPayeePT1Output{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollPayeePT1Input struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Payee_PT1_Request"`
	PutPayrollPayeePT1RequestType
}

type PutPayrollPayeePT1Output struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Payee_PT1_Response"`
	PutPayrollPayeePT1ResponseType
}

// GetPayrollPayeeFICAs (Get_Payroll_Payee_FICAs)
// 
// This operation returns FICA data for a Worker and Position.
func (c *Client) GetPayrollPayeeFICAs(ctx context.Context, input *GetPayrollPayeeFICAsInput) (output *GetPayrollPayeeFICAsOutput, err error) {
	output = &GetPayrollPayeeFICAsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollPayeeFICAsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payee_FICAs_Request"`
	GetPayrollPayeeFICAsRequestType
}

type GetPayrollPayeeFICAsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payee_FICAs_Response"`
	GetPayrollPayeeFICAsResponseType
}

// GetPayrollOffcyclePayments (Get_Payroll_Off-cycle_Payments)
// 
// This operation returns Payroll Off-cycle Payments (Manual and On-demand) for either a list of references or request criteria.
// 
// The user can specify the following request criteria:
// 
// 1 - Payment ID
// When specified, Payment ID alone will be used as a key to retrieve the off-cycle payment.  All other optional criteria will be ignored.
// 
// 2 - Employee Reference
// When specified, it will be used along with other optional criteria (except Payment ID) to return off-cycle payments.
// 
// 3 - Pay Group Reference
// When specified, it will be used along with other optional criteria (except Payment ID and Employee Reference) to return off-cycle payments.
// 
// All other optional criteria will be used with either Employee Reference or Pay Group Reference to retrieve off-cycle payments.
func (c *Client) GetPayrollOffcyclePayments(ctx context.Context, input *GetPayrollOffcyclePaymentsInput) (output *GetPayrollOffcyclePaymentsOutput, err error) {
	output = &GetPayrollOffcyclePaymentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollOffcyclePaymentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Off-cycle_Payments_Request"`
	GetPayrollOffcyclePaymentsRequestType
}

type GetPayrollOffcyclePaymentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Off-cycle_Payments_Response"`
	GetPayrollOffcyclePaymentsResponseType
}

// GetCompanyPaymentDates (Get_Company_Payment_Dates)
// 
// This operation returns a list of Payment Dates for a given Company and Date Range parms.
func (c *Client) GetCompanyPaymentDates(ctx context.Context, input *GetCompanyPaymentDatesInput) (output *GetCompanyPaymentDatesOutput, err error) {
	output = &GetCompanyPaymentDatesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCompanyPaymentDatesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Company_Payment_Dates_Request"`
	GetCompanyPaymentDatesRequestType
}

type GetCompanyPaymentDatesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Company_Payment_Dates_Response"`
	GetCompanyPaymentDatesResponseType
}

// GetPayrollPayeeTD1s (Get_Payroll_Payee_TD1s)
// 
// This web service is used to return Payroll Payee TD1s
func (c *Client) GetPayrollPayeeTD1s(ctx context.Context, input *GetPayrollPayeeTD1sInput) (output *GetPayrollPayeeTD1sOutput, err error) {
	output = &GetPayrollPayeeTD1sOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollPayeeTD1sInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payee_TD1s_Request"`
	GetPayrollPayeeTD1sRequestType
}

type GetPayrollPayeeTD1sOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payee_TD1s_Response"`
	GetPayrollPayeeTD1sResponseType
}

// PutPayrollPayeeTD1 (Put_Payroll_Payee_TD1)
// 
// This web service is used to load TD1&#39;s
func (c *Client) PutPayrollPayeeTD1(ctx context.Context, input *PutPayrollPayeeTD1Input) (output *PutPayrollPayeeTD1Output, err error) {
	output = &PutPayrollPayeeTD1Output{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollPayeeTD1Input struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Payee_TD1_Request"`
	PutPayrollPayeeTD1RequestType
}

type PutPayrollPayeeTD1Output struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Payee_T1_Response"`
	PutPayrollPayeeT1ResponseType
}

// GetCanadianRecordofEmploymentData (Get_Canadian_Record_of_Employment_Data)
// 
// Return Approved Canadian Record of Employment records for outbound ROE processing.
func (c *Client) GetCanadianRecordofEmploymentData(ctx context.Context, input *GetCanadianRecordofEmploymentDataInput) (output *GetCanadianRecordofEmploymentDataOutput, err error) {
	output = &GetCanadianRecordofEmploymentDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCanadianRecordofEmploymentDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Canadian_Record_of_Employment_Data_Request"`
	GetCanadianRecordofEmploymentDataRequestType
}

type GetCanadianRecordofEmploymentDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Canadian_Record_of_Employment_Data_Response"`
	GetCanadianRecordofEmploymentDataResponseType
}

// PutROEHistoryData (Put_ROE_History_Data)
// 
// This Web service is used to update the ROE History instances statuses, issued date, and issued number when the file is imported from Service Canada into Workday System.
func (c *Client) PutROEHistoryData(ctx context.Context, input *PutROEHistoryDataInput) (output *PutROEHistoryDataOutput, err error) {
	output = &PutROEHistoryDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutROEHistoryDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_ROE_History_Data_Request"`
	PutROEHistoryDataRequestType
}

type PutROEHistoryDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_ROE_History_Data_Response"`
	PutROEHistoryDataResponseType
}

// GetROEHistoryData (Get_ROE_History_Data)
// 
// Return Canadian Record of Employment history data for a given Record of Employment.
func (c *Client) GetROEHistoryData(ctx context.Context, input *GetROEHistoryDataInput) (output *GetROEHistoryDataOutput, err error) {
	output = &GetROEHistoryDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetROEHistoryDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_ROE_History_Data_Request"`
	GetROEHistoryDataRequestType
}

type GetROEHistoryDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_ROE_History_Data_Response"`
	GetROEHistoryDataResponseType
}

// PutCanadianRecordofEmploymentData (Put_Canadian_Record_of_Employment_Data)
// 
// The Web services is used to update the ROE instances statuses, issued date, and issued number when the file is imported from Service Canada into Workday System.
func (c *Client) PutCanadianRecordofEmploymentData(ctx context.Context, input *PutCanadianRecordofEmploymentDataInput) (output *PutCanadianRecordofEmploymentDataOutput, err error) {
	output = &PutCanadianRecordofEmploymentDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCanadianRecordofEmploymentDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Canadian_Record_of_Employment_Data_Request"`
	PutCanadianRecordofEmploymentDataRequestType
}

type PutCanadianRecordofEmploymentDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Canadian_Record_of_Employment_Data_Response"`
	PutCanadianRecordofEmploymentDataResponseType
}

// GetPayrollFederalW4TaxElection (Get_Payroll_Federal_W-4_Tax_Election)
// 
// This web service operation is designed to get worker federal W-4 tax election data from either the Reference ID or the Worker Reference/Company Reference/Effective Date.
func (c *Client) GetPayrollFederalW4TaxElection(ctx context.Context, input *GetPayrollFederalW4TaxElectionInput) (output *GetPayrollFederalW4TaxElectionOutput, err error) {
	output = &GetPayrollFederalW4TaxElectionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollFederalW4TaxElectionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Federal_W-4_Tax_Elections_Request"`
	GetPayrollFederalW4TaxElectionsRequestType
}

type GetPayrollFederalW4TaxElectionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Federal_W-4_Tax_Elections_Response"`
	GetPayrollFederalW4TaxElectionsResponseType
}

// PutPayrollFederalW4TaxElection (Put_Payroll_Federal_W-4_Tax_Election)
// 
// This web service operation is designed to put worker federal W-4 tax election data.
func (c *Client) PutPayrollFederalW4TaxElection(ctx context.Context, input *PutPayrollFederalW4TaxElectionInput) (output *PutPayrollFederalW4TaxElectionOutput, err error) {
	output = &PutPayrollFederalW4TaxElectionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollFederalW4TaxElectionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Federal_W-4_Tax_Election_Request"`
	PutPayrollFederalW4TaxElectionRequestType
}

type PutPayrollFederalW4TaxElectionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Federal_W-4_Tax_Election_Response"`
	PutPayrollFederalW4TaxElectionResponseType
}

// PutPayrollWorkerTaxTreaty (Put_Payroll_Worker_Tax_Treaty)
// 
// This web service method allows external clients to add or update a Payroll Worker Tax Treaty
func (c *Client) PutPayrollWorkerTaxTreaty(ctx context.Context, input *PutPayrollWorkerTaxTreatyInput) (output *PutPayrollWorkerTaxTreatyOutput, err error) {
	output = &PutPayrollWorkerTaxTreatyOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollWorkerTaxTreatyInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Worker_Tax_Treaty_Request"`
	PutPayrollWorkerTaxTreatyRequestType
}

type PutPayrollWorkerTaxTreatyOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Worker_Tax_Treaty_Response"`
	PutWorkerTaxTreatyResponseType
}

// GetPayrollWorkerTaxTreaties (Get_Payroll_Worker_Tax_Treaties)
// 
// Allows an external client to retrieve Worker Tax Treaty data.
func (c *Client) GetPayrollWorkerTaxTreaties(ctx context.Context, input *GetPayrollWorkerTaxTreatiesInput) (output *GetPayrollWorkerTaxTreatiesOutput, err error) {
	output = &GetPayrollWorkerTaxTreatiesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollWorkerTaxTreatiesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Worker_Tax_Treaties_Request"`
	GetPayrollWorkerTaxTreatiesRequestType
}

type GetPayrollWorkerTaxTreatiesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Worker_Tax_Treaties_Response"`
	GetWorkerTaxTreatiesResponseType
}

// GetPayrollInvoluntaryWithholdingOrders (Get_Payroll_Involuntary_Withholding_Orders)
// 
// This web service will extract all types of Involuntary Withholding Orders:
//  1-Bankruptcy Order
//  2-Creditor Garnishment
//  3-Federal Administrative Wage Garnishment
//  4-Federal Student Loan
//  5-Federal Tax Levy
//  6-State Tax Levy
//  7-Wage Assignment
//  8-Support Order
//  9-Support Order (Lump Sum)
// 
// For loading the above Withholding Orders, there is a corresponding web service: Put Payroll Involuntary Withholding Order.
func (c *Client) GetPayrollInvoluntaryWithholdingOrders(ctx context.Context, input *GetPayrollInvoluntaryWithholdingOrdersInput) (output *GetPayrollInvoluntaryWithholdingOrdersOutput, err error) {
	output = &GetPayrollInvoluntaryWithholdingOrdersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollInvoluntaryWithholdingOrdersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Involuntary_Withholding_Orders_Request"`
	GetPayrollInvoluntaryWithholdingOrdersRequestType
}

type GetPayrollInvoluntaryWithholdingOrdersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Involuntary_Withholding_Orders_Response"`
	GetPayrollInvoluntaryWithholdingOrdersResponseType
}

// PutPayrollInvoluntaryWithholdingOrder (Put_Payroll_Involuntary_Withholding_Order)
// 
// This web service will load all types of Withholding Orders:
//  1-Bankruptcy Order
//  2-Creditor Garnishment
//  3-Federal Administrative Wage Garnishment
//  4-Federal Student Loan
//  5-Federal Tax Levy
//  6-State Tax Levy
//  7-Wage Assignment
//  8-Support Order (Lump Sum)
//  9-Support Order
// 
// This method also supports Amend/Terminate Support Order.
// 
// For extracting the above Withholding Orders, there is a corresponding web service: Get Payroll Involuntary Withholding Orders.
func (c *Client) PutPayrollInvoluntaryWithholdingOrder(ctx context.Context, input *PutPayrollInvoluntaryWithholdingOrderInput) (output *PutPayrollInvoluntaryWithholdingOrderOutput, err error) {
	output = &PutPayrollInvoluntaryWithholdingOrderOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollInvoluntaryWithholdingOrderInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Involuntary_Withholding_Order_Request"`
	PutPayrollInvoluntaryWithholdingOrderRequestType
}

type PutPayrollInvoluntaryWithholdingOrderOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Involuntary_Withholding_Order_Response"`
	PutPayrollInvoluntaryWithholdingOrderResponseType
}

// ChangeNoRetroProcessingPriorTo (Change_No_Retro_Processing_Prior_To)
// 
// This Web Service is to allow clients to push out the &#34;No Retro Processing Prior to&#34; dates for workers.
func (c *Client) ChangeNoRetroProcessingPriorTo(ctx context.Context, input *ChangeNoRetroProcessingPriorToInput) (output *ChangeNoRetroProcessingPriorToOutput, err error) {
	output = &ChangeNoRetroProcessingPriorToOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeNoRetroProcessingPriorToInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_No_Retro_Processing_Prior_To_Request"`
	ChangeNoRetroProcessingPriorToRequestType
}

type ChangeNoRetroProcessingPriorToOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_No_Retro_Processing_Prior_To_Response"`
	ChangeNoRetroProcessingPriorToResponseType
}

// GetMultipleWorksiteReport (Get_Multiple_Worksite_Report)
// 
// This web service retrieves multiple worksite report data for a company and quarter.  The report is used to collect statistical information to supplement statewide data provided through employer&#39;s SUI reports.
func (c *Client) GetMultipleWorksiteReport(ctx context.Context, input *GetMultipleWorksiteReportInput) (output *GetMultipleWorksiteReportOutput, err error) {
	output = &GetMultipleWorksiteReportOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetMultipleWorksiteReportInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Multiple_Worksite_Report_Request"`
	GetMultipleWorksiteReportRequestType
}

type GetMultipleWorksiteReportOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Multiple_Worksite_Report_Response"`
	GetMultipleWorksiteReportResponseType
}

// GetPayrollLimitOverrides (Get_Payroll_Limit_Overrides)
// 
// This web service will return Payroll Limit Override from either Payroll Limit Override Reference, or Worker Reference/Pay Component/Start Date/End Date filter criteria.
func (c *Client) GetPayrollLimitOverrides(ctx context.Context, input *GetPayrollLimitOverridesInput) (output *GetPayrollLimitOverridesOutput, err error) {
	output = &GetPayrollLimitOverridesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollLimitOverridesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Limit_Overrides_Request"`
	GetPayrollLimitOverridesRequestType
}

type GetPayrollLimitOverridesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Limit_Overrides_Response"`
	GetPayrollLimitOverridesResponseType
}

// PutPayrollLimitOverride (Put_Payroll_Limit_Override)
// 
// This web service is for adding or updating Payroll Limit Override for Worker.
func (c *Client) PutPayrollLimitOverride(ctx context.Context, input *PutPayrollLimitOverrideInput) (output *PutPayrollLimitOverrideOutput, err error) {
	output = &PutPayrollLimitOverrideOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollLimitOverrideInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Limit_Override_Request"`
	PutPayrollLimitOverrideRequestType
}

type PutPayrollLimitOverrideOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Limit_Override_Response"`
	PutPayrollLimitOverrideResponseType
}

// GetWorkerCostingAllocations (Get_Worker_Costing_Allocations)
// 
// Get Worker- or Position Restrictions-based costing allocations, either for specific criteria or for all workers or positions.  Can include only the external integration references or the full allocation details.
func (c *Client) GetWorkerCostingAllocations(ctx context.Context, input *GetWorkerCostingAllocationsInput) (output *GetWorkerCostingAllocationsOutput, err error) {
	output = &GetWorkerCostingAllocationsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetWorkerCostingAllocationsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Worker_Costing_Allocations_Request"`
	GetWorkerCostingAllocationsRequestType
}

type GetWorkerCostingAllocationsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Worker_Costing_Allocations_Response"`
	GetWorkerCostingAllocationsResponseType
}

// GetPayrollPayeeRPPorDPSPRegistrationNumbers (Get_Payroll_Payee_RPP_or_DPSP_Registration_Numbers)
// 
// This web service is used to return RPP and DPSP registration numbers
func (c *Client) GetPayrollPayeeRPPorDPSPRegistrationNumbers(ctx context.Context, input *GetPayrollPayeeRPPorDPSPRegistrationNumbersInput) (output *GetPayrollPayeeRPPorDPSPRegistrationNumbersOutput, err error) {
	output = &GetPayrollPayeeRPPorDPSPRegistrationNumbersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollPayeeRPPorDPSPRegistrationNumbersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payee_RPP_or_DPSP_Registration_Numbers_Request"`
	GetPayrollPayeeRPPorDPSPRegistrationNumbersRequestType
}

type GetPayrollPayeeRPPorDPSPRegistrationNumbersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payee_RPP_or_DPSP_Registration_Numbers_Response"`
	GetPayrollPayeeRPPorDPSPRegistrationNumbersResponseType
}

// PutPayrollPayeeRPPorDPSPRegistrationNumber (Put_Payroll_Payee_RPP_or_DPSP_Registration_Number)
// 
// This web service is used to load RPP and DPSP Registration Numbers
func (c *Client) PutPayrollPayeeRPPorDPSPRegistrationNumber(ctx context.Context, input *PutPayrollPayeeRPPorDPSPRegistrationNumberInput) (output *PutPayrollPayeeRPPorDPSPRegistrationNumberOutput, err error) {
	output = &PutPayrollPayeeRPPorDPSPRegistrationNumberOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollPayeeRPPorDPSPRegistrationNumberInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Payee_RPP_or_DPSP_Registration_Number_Request"`
	PutPayrollPayeeRPPorDPSPRegistrationNumberRequestType
}

type PutPayrollPayeeRPPorDPSPRegistrationNumberOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Payee_RPP_or_DPSP_Registration_Number_Response"`
	PutPayrollPayeeRPPorDPSPRegistrationNumberResponseType
}

// GetPayrollUSAStateandLocalTaxElections (Get_Payroll_USA_State_and_Local_Tax_Elections)
// 
// Retrieve Worker Tax Elections for State and Local Tax Authorities
// 
// NOTE: The As Of Entry Moment and As Of Effective date fields in the Response Filter are not supported by this method - any values you pass in via these fields will be ignored. Use the Effective Date field in the Request Criteria instead to retrieve elections effective as of a given date.  The operation does not support criteria based requests with an As_Of_Entry_DateTime in the past, unless for paging consistency the As_Of_Entry_DateTime matches the page 1 request.
func (c *Client) GetPayrollUSAStateandLocalTaxElections(ctx context.Context, input *GetPayrollUSAStateandLocalTaxElectionsInput) (output *GetPayrollUSAStateandLocalTaxElectionsOutput, err error) {
	output = &GetPayrollUSAStateandLocalTaxElectionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollUSAStateandLocalTaxElectionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_USA_State_and_Local_Tax_Elections_Request"`
	GetPayrollUSAStateandLocalTaxElectionsRequestType
}

type GetPayrollUSAStateandLocalTaxElectionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_USA_State_and_Local_Tax_Elections_Response"`
	GetPayrollUSAStateandLocalTaxElectionsResponseType
}

// PutPayrollUSAStateandLocalTaxElection (Put_Payroll_USA_State_and_Local_Tax_Election)
// 
// Create Worker Tax Elections for State and Local Tax Authorities  The operation does not support criteria based requests with an As_Of_Entry_DateTime in the past, unless for paging consistency the As_Of_Entry_DateTime matches the page 1 request.
func (c *Client) PutPayrollUSAStateandLocalTaxElection(ctx context.Context, input *PutPayrollUSAStateandLocalTaxElectionInput) (output *PutPayrollUSAStateandLocalTaxElectionOutput, err error) {
	output = &PutPayrollUSAStateandLocalTaxElectionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollUSAStateandLocalTaxElectionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_USA_State_and_Local_Tax_Election_Request"`
	PutPayrollUSAStateandLocalTaxElectionRequestType
}

type PutPayrollUSAStateandLocalTaxElectionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_USA_State_and_Local_Tax_Election_Response"`
	PutPayrollUSAStateandLocalTaxElectionResponseType
}

// GetPayrollPayeeOngoingWorkJurisdictionTaxElection (Get_Payroll_Payee_Ongoing_Work_Jurisdiction_Tax_Election)
// 
// This web service operation is designed to get worker ongoing work jurisdiction tax election data from either the Reference ID or the Worker Reference/Company Reference/Effective Date.
func (c *Client) GetPayrollPayeeOngoingWorkJurisdictionTaxElection(ctx context.Context, input *GetPayrollPayeeOngoingWorkJurisdictionTaxElectionInput) (output *GetPayrollPayeeOngoingWorkJurisdictionTaxElectionOutput, err error) {
	output = &GetPayrollPayeeOngoingWorkJurisdictionTaxElectionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollPayeeOngoingWorkJurisdictionTaxElectionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payee_Ongoing_Work_Jurisdiction_Tax_Election_Request"`
	GetPayrollPayeeOngoingWorkJurisdictionTaxElectionRequestType
}

type GetPayrollPayeeOngoingWorkJurisdictionTaxElectionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payee_Ongoing_Work_Jurisdiction_Tax_Election_Response"`
	GetPayrollPayeeOngoingWorkJurisdictionTaxElectionResponseType
}

// PutPayrollPayeeOngoingWorkJurisdictionTaxElection (Put_Payroll_Payee_Ongoing_Work_Jurisdiction_Tax_Election)
// 
// This web service operation is designed to put worker ongoing work jurisdiction tax election data into Workday Payroll
func (c *Client) PutPayrollPayeeOngoingWorkJurisdictionTaxElection(ctx context.Context, input *PutPayrollPayeeOngoingWorkJurisdictionTaxElectionInput) (output *PutPayrollPayeeOngoingWorkJurisdictionTaxElectionOutput, err error) {
	output = &PutPayrollPayeeOngoingWorkJurisdictionTaxElectionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollPayeeOngoingWorkJurisdictionTaxElectionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Payee_Ongoing_Work_Jurisdiction_Tax_Election_Request"`
	PutPayrollPayeeOngoingWorkJurisdictionTaxElectionRequestType
}

type PutPayrollPayeeOngoingWorkJurisdictionTaxElectionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Payee_Ongoing_Work_Jurisdiction_Tax_Election_Response"`
	PutPayrollPayeeOngoingWorkJurisdictionTaxElectionResponseType
}

// GetMonthlyWorkerTaxFilingData (Get_Monthly_Worker_Tax_Filing_Data)
// 
// This web service operation returns monthly tax filing data based on request criteria.
func (c *Client) GetMonthlyWorkerTaxFilingData(ctx context.Context, input *GetMonthlyWorkerTaxFilingDataInput) (output *GetMonthlyWorkerTaxFilingDataOutput, err error) {
	output = &GetMonthlyWorkerTaxFilingDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetMonthlyWorkerTaxFilingDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Monthly_Worker_Tax_Filing_Data_Request"`
	GetMonthlyWorkerTaxFilingDataRequestType
}

type GetMonthlyWorkerTaxFilingDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Monthly_Worker_Tax_Filing_Data_Response"`
	GetMonthlyWorkerTaxFilingDataResponseType
}

// GetSuccessorEmployers (Get_Successor_Employers)
// 
// This public web service operation is designed to retrieve Successor Employer relationships. If an effective date is provided, it will retrieve Successor Employer relationships that are active as of the effective date. Otherwise, all Successor Employer relationships will be retrieved.
func (c *Client) GetSuccessorEmployers(ctx context.Context, input *GetSuccessorEmployersInput) (output *GetSuccessorEmployersOutput, err error) {
	output = &GetSuccessorEmployersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetSuccessorEmployersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Successor_Employers_Request"`
	GetSuccessorEmployersRequestType
}

type GetSuccessorEmployersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Successor_Employers_Response"`
	GetSuccessorEmployersResponseType
}

// PutSuccessorEmployer (Put_Successor_Employer)
// 
// This public web service operation is designed to add/update Successor Employer relationships.
func (c *Client) PutSuccessorEmployer(ctx context.Context, input *PutSuccessorEmployerInput) (output *PutSuccessorEmployerOutput, err error) {
	output = &PutSuccessorEmployerOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutSuccessorEmployerInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Successor_Employer_Request"`
	PutSuccessorEmployerRequestType
}

type PutSuccessorEmployerOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Successor_Employer_Response"`
	PutSuccessorEmployerResponseType
}

// PutPayrollPayeeFUTA (Put_Payroll_Payee_FUTA)
// 
// Web Service for entering and modifying FUTA records
func (c *Client) PutPayrollPayeeFUTA(ctx context.Context, input *PutPayrollPayeeFUTAInput) (output *PutPayrollPayeeFUTAOutput, err error) {
	output = &PutPayrollPayeeFUTAOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollPayeeFUTAInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Payee_FUTA_Request"`
	PutPayrollPayeeFUTARequestType
}

type PutPayrollPayeeFUTAOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Payee_FUTA_Response"`
	PutPayrollPayeeFUTAResponseType
}

// GetPayrollPayeeFUTAs (Get_Payroll_Payee_FUTAs)
// 
// This web service operation is designed to get worker FUTA tax election data from either the Reference ID or the Worker Reference/Company Reference/Effective Date.
func (c *Client) GetPayrollPayeeFUTAs(ctx context.Context, input *GetPayrollPayeeFUTAsInput) (output *GetPayrollPayeeFUTAsOutput, err error) {
	output = &GetPayrollPayeeFUTAsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollPayeeFUTAsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payee_FUTAs_Request"`
	GetPayrollPayeeFUTAsRequestType
}

type GetPayrollPayeeFUTAsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payee_FUTAs_Response"`
	GetPayrollPayeeFUTAsResponseType
}

// GetPayrollDeductionRecipients (Get_Payroll_Deduction_Recipients)
// 
// This web service operation is designed to retrieve all the Payroll Deduction Recipients matching the specified criteria. The request criteria can be for a single transaction based on Reference, or all Deduction Recipients can be retrieved if no criteria is specified.
func (c *Client) GetPayrollDeductionRecipients(ctx context.Context, input *GetPayrollDeductionRecipientsInput) (output *GetPayrollDeductionRecipientsOutput, err error) {
	output = &GetPayrollDeductionRecipientsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollDeductionRecipientsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Deduction_Recipients_Request"`
	GetPayrollDeductionRecipientsRequestType
}

type GetPayrollDeductionRecipientsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Deduction_Recipients_Response"`
	GetPayrollDeductionRecipientsResponseType
}

// PutPayrollDeductionRecipient (Put_Payroll_Deduction_Recipient)
// 
// This web service operation is designed to put Deduction Recipient data into Workday Payroll
func (c *Client) PutPayrollDeductionRecipient(ctx context.Context, input *PutPayrollDeductionRecipientInput) (output *PutPayrollDeductionRecipientOutput, err error) {
	output = &PutPayrollDeductionRecipientOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollDeductionRecipientInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Deduction_Recipient_Request"`
	PutPayrollDeductionRecipientRequestType
}

type PutPayrollDeductionRecipientOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Deduction_Recipient_Response"`
	PutPayrollDeductionRecipientResponseType
}

// GetPaycheckDeliveries (Get_Paycheck_Deliveries)
// 
// The Paycheck Deliveries operation returns all the Paycheck delivery information for Workers. Includes the Delivery Method, and also how to receive copies of their payslips (Electronic/Paper etc).
func (c *Client) GetPaycheckDeliveries(ctx context.Context, input *GetPaycheckDeliveriesInput) (output *GetPaycheckDeliveriesOutput, err error) {
	output = &GetPaycheckDeliveriesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPaycheckDeliveriesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Paycheck_Deliveries_Request"`
	GetPaycheckDeliveriesRequestType
}

type GetPaycheckDeliveriesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Paycheck_Deliveries_Response"`
	GetPaycheckDeliveriesResponseType
}

// PutPaycheckDeliveryPublic (Put_Paycheck_Delivery__Public_)
// 
// This operation allows the entry of a workers Paycheck Delivery and Payslip Printing options.
func (c *Client) PutPaycheckDeliveryPublic(ctx context.Context, input *PutPaycheckDeliveryPublicInput) (output *PutPaycheckDeliveryPublicOutput, err error) {
	output = &PutPaycheckDeliveryPublicOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPaycheckDeliveryPublicInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Paycheck_Delivery__Public__Request"`
	PutPaycheckDeliveryPublicRequestType
}

type PutPaycheckDeliveryPublicOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Paycheck_Delivery__Public__Response"`
	PutPaycheckDeliveryPublicResponseType
}

// PutSingleLegalEntity (Put_Single_Legal_Entity)
// 
// This web service is designed to add or update Single Legal Entity company relationships
func (c *Client) PutSingleLegalEntity(ctx context.Context, input *PutSingleLegalEntityInput) (output *PutSingleLegalEntityOutput, err error) {
	output = &PutSingleLegalEntityOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutSingleLegalEntityInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Single_Legal_Entity_Request"`
	PutSingleLegalEntityRequestType
}

type PutSingleLegalEntityOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Single_Legal_Entity_Response"`
	PutSingleLegalEntityResponseType
}

// GetSingleLegalEntities (Get_Single_Legal_Entities)
// 
// This web service retrieves Single Legal Entity company relationships.
// 
// If the domains for this task change, also change the domains of Secure Method Definition, Single Legal Entity@get Single Legal Entity References for Request Criteria(SS)*P*S[rsmb]&#43;PA / Process: Tax Filing/W-2s (Run) - USA (Secure Results by Domain), 5236$3343.
func (c *Client) GetSingleLegalEntities(ctx context.Context, input *GetSingleLegalEntitiesInput) (output *GetSingleLegalEntitiesOutput, err error) {
	output = &GetSingleLegalEntitiesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetSingleLegalEntitiesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Single_Legal_Entities_Request"`
	GetSingleLegalEntitiesRequestType
}

type GetSingleLegalEntitiesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Single_Legal_Entities_Response"`
	GetSingleLegalEntitiesResponseType
}

// GetW2W2CPrintingElection (Get_W-2_W-2C_Printing_Election)
// 
// This web service operation is designed to get worker W-2/W-2C printing elections from either the Reference ID or the Worker Reference/Company Reference.
// 
// If the domain for this task changes, also change the domain of Secure Method Definition, Worker@get Worker if in Request Otherwise Get all Workers for Company parm(SSC)*P*S[rsmb]&#43;PA / Worker Data: Payroll (Company Specific) - USA (Secure Results by Domain), 5236$3342.
func (c *Client) GetW2W2CPrintingElection(ctx context.Context, input *GetW2W2CPrintingElectionInput) (output *GetW2W2CPrintingElectionOutput, err error) {
	output = &GetW2W2CPrintingElectionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetW2W2CPrintingElectionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_W-2_W-2C_Printing_Election_Request"`
	GetW2W2CPrintingElectionRequestType
}

type GetW2W2CPrintingElectionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_W-2_W-2C_Printing_Election_Response"`
	GetW2W2CPrintingElectionResponseType
}

// PutW2W2CPrintingElection (Put_W-2_W-2C_Printing_Election)
// 
// This web service operation is designed to put W-2/W-2C Printing Election data.
func (c *Client) PutW2W2CPrintingElection(ctx context.Context, input *PutW2W2CPrintingElectionInput) (output *PutW2W2CPrintingElectionOutput, err error) {
	output = &PutW2W2CPrintingElectionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutW2W2CPrintingElectionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_W-2_W-2C_Printing_Election_Request"`
	PutW2W2CPrintingElectionRequestType
}

type PutW2W2CPrintingElectionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_W-2_W-2C_Printing_Election_for_Worker_Response"`
	PutW2W2CPrintingElectionforWorkerResponseType
}

// GetTaxLevyDeductionRestrictions (Get_Tax_Levy_Deduction_Restrictions)
// 
// This web service will extract all tax levy deduction restrictions.
// For loading the above tax levy deduction restrictions, there is a corresponding web service: Put Tax Levy Deduction Restriction.
func (c *Client) GetTaxLevyDeductionRestrictions(ctx context.Context, input *GetTaxLevyDeductionRestrictionsInput) (output *GetTaxLevyDeductionRestrictionsOutput, err error) {
	output = &GetTaxLevyDeductionRestrictionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetTaxLevyDeductionRestrictionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Tax_Levy_Deduction_Restrictions_Request"`
	GetTaxLevyDeductionRestrictionsRequestType
}

type GetTaxLevyDeductionRestrictionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Tax_Levy_Deduction_Restrictions_Response"`
	GetTaxLevyDeductionRestrictionsResponseType
}

// PutTaxLevyDeductionRestriction (Put_Tax_Levy_Deduction_Restriction)
// 
// This web service will load tax levy deduction restrictions.
// For extracting the above deduction restrictions, there is a corresponding web service: Get Tax Levy Deduction Restrictions.
func (c *Client) PutTaxLevyDeductionRestriction(ctx context.Context, input *PutTaxLevyDeductionRestrictionInput) (output *PutTaxLevyDeductionRestrictionOutput, err error) {
	output = &PutTaxLevyDeductionRestrictionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutTaxLevyDeductionRestrictionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Tax_Levy_Deduction_Restriction_Request"`
	PutTaxLevyDeductionRestrictionRequestType
}

type PutTaxLevyDeductionRestrictionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Tax_Levy_Deduction_Restriction_Response"`
	PutTaxLevyDeductionRestrictionResponseType
}

// PutTaxDocumentDelivery (Put_Tax_Document_Delivery)
// 
// This web service is designed to add or update the payroll year end form printing elections for Worker(s).  These elections determine whether the worker will receive a paper copy of their payroll year end form(s).
func (c *Client) PutTaxDocumentDelivery(ctx context.Context, input *PutTaxDocumentDeliveryInput) (output *PutTaxDocumentDeliveryOutput, err error) {
	output = &PutTaxDocumentDeliveryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutTaxDocumentDeliveryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Tax_Document_Delivery_Request"`
	PutTaxDocumentDeliveryRequestType
}

type PutTaxDocumentDeliveryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Tax_Document_Delivery_Response"`
	PutTaxDocumentDeliveryResponseType
}

// GetTaxDocumentDeliveries (Get_Tax_Document_Deliveries)
// 
// This web service is designed to get the payroll year end form printing elections for Company, Delivery Group and optionally Worker(s). These elections determine whether the worker will receive a paper copy of their payroll year end form(s).
func (c *Client) GetTaxDocumentDeliveries(ctx context.Context, input *GetTaxDocumentDeliveriesInput) (output *GetTaxDocumentDeliveriesOutput, err error) {
	output = &GetTaxDocumentDeliveriesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetTaxDocumentDeliveriesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Tax_Document_Deliveries_Request"`
	GetTaxDocumentDeliveriesRequestType
}

type GetTaxDocumentDeliveriesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Tax_Document_Deliveries_Response"`
	GetTaxDocumentDeliveriesResponseType
}

// GetPeriodicCompanyCANTaxRemittanceData (Get_Periodic_Company_CAN_Tax_Remittance_Data)
// 
// This web service operation retrieves the periodic tax remittance data for a company.  The tax remittance data is aggregated by Payment Date and Tax Account Number and Deduction Code.
func (c *Client) GetPeriodicCompanyCANTaxRemittanceData(ctx context.Context, input *GetPeriodicCompanyCANTaxRemittanceDataInput) (output *GetPeriodicCompanyCANTaxRemittanceDataOutput, err error) {
	output = &GetPeriodicCompanyCANTaxRemittanceDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPeriodicCompanyCANTaxRemittanceDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Periodic_Company_CAN_Tax_Remittance_Data_Request"`
	GetPeriodicCompanyCANTaxRemittanceDataRequestType
}

type GetPeriodicCompanyCANTaxRemittanceDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Periodic_Company_CAN_Tax_Remittance_Data_Response"`
	GetPeriodicCompanyCANTaxRemittanceDataResponseType
}

// PutWithholdingOrderLocalMinimumWageRate (Put_Withholding_Order_Local_Minimum_Wage_Rate)
// 
// This web service will load Withholding Order Local Minimum Wage(s) for existing Withholding Orders.
func (c *Client) PutWithholdingOrderLocalMinimumWageRate(ctx context.Context, input *PutWithholdingOrderLocalMinimumWageRateInput) (output *PutWithholdingOrderLocalMinimumWageRateOutput, err error) {
	output = &PutWithholdingOrderLocalMinimumWageRateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutWithholdingOrderLocalMinimumWageRateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Withholding_Order_Local_Minimum_Wage_Rate_Request"`
	PutWithholdingOrderLocalMinimumWageRateRequestType
}

type PutWithholdingOrderLocalMinimumWageRateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Withholding_Order_Local_Minimum_Wage_Rate_Response"`
	PutWithholdingOrderLocalMinimumWageRateResponseType
}

// GetWithholdingOrderLocalMinimumWageRate (Get_Withholding_Order_Local_Minimum_Wage_Rate)
// 
// This web service will return Withholding Order Local Minimum Wage(s) for either a company or ~workers~.
func (c *Client) GetWithholdingOrderLocalMinimumWageRate(ctx context.Context, input *GetWithholdingOrderLocalMinimumWageRateInput) (output *GetWithholdingOrderLocalMinimumWageRateOutput, err error) {
	output = &GetWithholdingOrderLocalMinimumWageRateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetWithholdingOrderLocalMinimumWageRateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Withholding_Order_Local_Minimum_Wage_Rate_Request"`
	GetWithholdingOrderLocalMinimumWageRateRequestType
}

type GetWithholdingOrderLocalMinimumWageRateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Withholding_Order_Local_Minimum_Wage_Rate_Response"`
	GetWithholdingOrderLocalMinimumWageRateResponseType
}

// GetPayrollTaxMappingsonLocation (Get_Payroll_Tax_Mappings_on_Location)
// 
// This web service operation is designed to get the Payroll Tax Mappings for a Location, given it has Payroll Tax usage enabled
func (c *Client) GetPayrollTaxMappingsonLocation(ctx context.Context, input *GetPayrollTaxMappingsonLocationInput) (output *GetPayrollTaxMappingsonLocationOutput, err error) {
	output = &GetPayrollTaxMappingsonLocationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollTaxMappingsonLocationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Tax_Mappings_on_Location_Request"`
	GetPayrollTaxMappingsonLocationRequestType
}

type GetPayrollTaxMappingsonLocationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Tax_Mappings_on_Location_Response"`
	GetPayrollTaxMappingsonLocationResponseType
}

// PutPayrollTaxMappingonLocation (Put_Payroll_Tax_Mapping_on_Location)
// 
// This web service operation is designed to put Payroll Tax Location Mapping data into Workday Payroll
func (c *Client) PutPayrollTaxMappingonLocation(ctx context.Context, input *PutPayrollTaxMappingonLocationInput) (output *PutPayrollTaxMappingonLocationOutput, err error) {
	output = &PutPayrollTaxMappingonLocationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollTaxMappingonLocationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Tax_Mapping_on_Location_Request"`
	PutPayrollTaxMappingonLocationRequestType
}

type PutPayrollTaxMappingonLocationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Tax_Mapping_on_Location_Response"`
	PutPayrollTaxMappingonLocationResponseType
}

// PutRunPayCalculation (Put_Run_Pay_Calculation)
// 
// This Web service initiates a Run Pay Calculation Job and returns its Reference ID as a response.
func (c *Client) PutRunPayCalculation(ctx context.Context, input *PutRunPayCalculationInput) (output *PutRunPayCalculationOutput, err error) {
	output = &PutRunPayCalculationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutRunPayCalculationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Run_Pay_Calculation_Request"`
	PutRunPayCalculationRequestType
}

type PutRunPayCalculationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Run_Pay_Calculation_Response"`
	PutRunPayCalculationResponseType
}

// GetEFW2YearEndWorkerFilingData (Get_EFW2_Year_End_Worker_Filing_Data)
// 
// Exposes the worker W-2 year end data for integration with the Social Security Administration (SSA) or other filing system. It is recommended that integrations do not reference box numbers when searching for specific box amounts. For example, search on the string &#34;Allocated tips&#34; instead of the string &#34;8 - Allocated tips&#34;. This is because over time box numbers may change.
func (c *Client) GetEFW2YearEndWorkerFilingData(ctx context.Context, input *GetEFW2YearEndWorkerFilingDataInput) (output *GetEFW2YearEndWorkerFilingDataOutput, err error) {
	output = &GetEFW2YearEndWorkerFilingDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEFW2YearEndWorkerFilingDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_EFW2_Year_End_Worker_Filing_Data_Request"`
	GetEFW2YearEndWorkerFilingDataRequestType
}

type GetEFW2YearEndWorkerFilingDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_EFW2_Year_End_Worker_Filing_Data_Response"`
	GetEFW2YearEndWorkerFilingDataResponseType
}

// GetPayrollPrePrintedPayslips (Get_Payroll_Pre-Printed_Payslips)
// 
// Public web service that returns Pre-Printed Payslip repository documents
func (c *Client) GetPayrollPrePrintedPayslips(ctx context.Context, input *GetPayrollPrePrintedPayslipsInput) (output *GetPayrollPrePrintedPayslipsOutput, err error) {
	output = &GetPayrollPrePrintedPayslipsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollPrePrintedPayslipsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Pre-Printed_Payslips_Request"`
	GetPayrollPrePrintedPayslipsRequestType
}

type GetPayrollPrePrintedPayslipsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Pre-Printed_Payslips_Response"`
	GetPayrollPrePrintedPayslipsResponseType
}

// GetEFW2YearEndEmployerFilingData (Get_EFW2_Year_End_Employer_Filing_Data)
// 
// Exposes the employer-level W-2 year end data for integration with the Social Security Administration (SSA) or other filing system. This includes employer totals for the W-2 box amounts as well as submitter information.
func (c *Client) GetEFW2YearEndEmployerFilingData(ctx context.Context, input *GetEFW2YearEndEmployerFilingDataInput) (output *GetEFW2YearEndEmployerFilingDataOutput, err error) {
	output = &GetEFW2YearEndEmployerFilingDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEFW2YearEndEmployerFilingDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_EFW2_Year_End_Employer_Filing_Data_Request"`
	GetEFW2YearEndEmployerFilingDataRequestType
}

type GetEFW2YearEndEmployerFilingDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_EFW2_Year_End_Employer_Filing_Data_Response"`
	GetEFW2YearEndEmployerFilingDataResponseType
}

// GetRunPayCalculation (Get_Run_Pay_Calculation)
// 
// Web Service to get Pay Calculation Background Processes.
func (c *Client) GetRunPayCalculation(ctx context.Context, input *GetRunPayCalculationInput) (output *GetRunPayCalculationOutput, err error) {
	output = &GetRunPayCalculationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetRunPayCalculationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Run_Pay_Calculation_Request"`
	GetRunPayCalculationRequestType
}

type GetRunPayCalculationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Run_Pay_Calculation_Response"`
	GetRunPayCalculationResponseType
}

// GetPayrollPayeeTaxLocationMappings (Get_Payroll_Payee_Tax_Location_Mappings)
// 
// This web service operation is designed to get worker tax location mapping data from either the Reference ID or the Worker Reference/Company Reference/Effective Date.
func (c *Client) GetPayrollPayeeTaxLocationMappings(ctx context.Context, input *GetPayrollPayeeTaxLocationMappingsInput) (output *GetPayrollPayeeTaxLocationMappingsOutput, err error) {
	output = &GetPayrollPayeeTaxLocationMappingsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollPayeeTaxLocationMappingsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payee_Tax_Location_Mappings_Request"`
	GetPayrollPayeeTaxLocationMappingsRequestType
}

type GetPayrollPayeeTaxLocationMappingsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Payee_Tax_Location_Mappings_Response"`
	GetPayrollPayeeTaxLocationMappingsResponseType
}

// PutPayrollPayeeTaxLocationMapping (Put_Payroll_Payee_Tax_Location_Mapping)
// 
// This web service operation is designed to put worker tax location mapping data into Workday Payroll.
func (c *Client) PutPayrollPayeeTaxLocationMapping(ctx context.Context, input *PutPayrollPayeeTaxLocationMappingInput) (output *PutPayrollPayeeTaxLocationMappingOutput, err error) {
	output = &PutPayrollPayeeTaxLocationMappingOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPayrollPayeeTaxLocationMappingInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Payee_Tax_Location_Mapping_Request"`
	PutPayrollPayeeTaxLocationMappingRequestType
}

type PutPayrollPayeeTaxLocationMappingOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payroll_Payee_Tax_Location_Mapping_Response"`
	PutPayrollPayeeTaxLocationMappingResponseType
}

// GetPayrollReportingCodesforWorkers (Get_Payroll_Reporting_Codes_for_Workers)
// 
// Returns effective payroll reporting codes associated with multiple workers.
func (c *Client) GetPayrollReportingCodesforWorkers(ctx context.Context, input *GetPayrollReportingCodesforWorkersInput) (output *GetPayrollReportingCodesforWorkersOutput, err error) {
	output = &GetPayrollReportingCodesforWorkersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPayrollReportingCodesforWorkersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Reporting_Codes_for_Workers_Request"`
	GetPayrollReportingCodesforWorkersRequestType
}

type GetPayrollReportingCodesforWorkersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payroll_Reporting_Codes_for_Workers_Response"`
	GetPayrollReportingCodesforWorkersResponseType
}

// ImportPayrollInput (Import_Payroll_Input)
// 
// This public web service operation is designed to bulk add/update Payroll Input data from external systems into Workday.
func (c *Client) ImportPayrollInput(ctx context.Context, input *ImportPayrollInputInput) (output *ImportPayrollInputOutput, err error) {
	output = &ImportPayrollInputOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportPayrollInputInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Payroll_Input_Request"`
	ImportPayrollInputRequestType
}

type ImportPayrollInputOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportPayrollOffCyclePayment (Import_Payroll_Off-Cycle_Payment)
// 
// Imports one or more Off-cycle Payments using an asynchronous request mode. Subsequent status updates for the import process can be retrieved using Get Import Processes (Web Service) and Get Import Process Messages (Web Service).
func (c *Client) ImportPayrollOffCyclePayment(ctx context.Context, input *ImportPayrollOffCyclePaymentInput) (output *ImportPayrollOffCyclePaymentOutput, err error) {
	output = &ImportPayrollOffCyclePaymentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportPayrollOffCyclePaymentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Payroll_Off-cycle_Payment_Request"`
	ImportPayrollOffcyclePaymentRequestType
}

type ImportPayrollOffCyclePaymentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportROEPriorPeriodHistoryResults (Import_ROE_Prior_Period_History_Results)
// 
// This Web Service is used to load the Canadian Employee Prior History Results for ROE processing.
func (c *Client) ImportROEPriorPeriodHistoryResults(ctx context.Context, input *ImportROEPriorPeriodHistoryResultsInput) (output *ImportROEPriorPeriodHistoryResultsOutput, err error) {
	output = &ImportROEPriorPeriodHistoryResultsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportROEPriorPeriodHistoryResultsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_ROE_Prior_Period_History_Results_Request"`
	ImportROEPriorPeriodHistoryResultsRequestType
}

type ImportROEPriorPeriodHistoryResultsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportMaintainPayrollReportingCodesforWorkers (Import_Maintain_Payroll_Reporting_Codes_for_Workers)
// 
// High Volume Web Service task for updating the payroll reporting codes for a large number of existing workers. This task can accept multiple elements, each of which takes in a worker, an effective date, and payroll reporting code(s). The web service then launches an Event that updates the relationship between the worker and their effective dated payroll reporting codes.
func (c *Client) ImportMaintainPayrollReportingCodesforWorkers(ctx context.Context, input *ImportMaintainPayrollReportingCodesforWorkersInput) (output *ImportMaintainPayrollReportingCodesforWorkersOutput, err error) {
	output = &ImportMaintainPayrollReportingCodesforWorkersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportMaintainPayrollReportingCodesforWorkersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Maintain_Payroll_Reporting_Codes_for_Worker_Request"`
	ImportMaintainPayrollReportingCodesforWorkerRequestType
}

type ImportMaintainPayrollReportingCodesforWorkersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportPayrollHistoryPayment (Import_Payroll_History_Payment)
// 
// This public web service operation is designed to bulk add/update historical payroll payments from external systems into Workday Payroll.
func (c *Client) ImportPayrollHistoryPayment(ctx context.Context, input *ImportPayrollHistoryPaymentInput) (output *ImportPayrollHistoryPaymentOutput, err error) {
	output = &ImportPayrollHistoryPaymentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportPayrollHistoryPaymentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Payroll_History_Payment_Request"`
	ImportPayrollHistoryPaymentRequestType
}

type ImportPayrollHistoryPaymentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

