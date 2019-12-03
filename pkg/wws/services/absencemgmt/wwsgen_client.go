
// Package absencemgmt
//
// The Absence Management Web Service contains operations that expose absence-related data, including Employee Time Off and Absence Inputs for time off and accrual adjustments and overrides, and Leave Requests, in Workday Human Capital Management Business Services.
package absencemgmt

import (
	"context"

	"github.com/ianlopshire/go-workday/pkg/wws"
)

const (
	CurrentVersion = "v31.2"
	ServiceName = "Absence_Management"
)

type Client struct {
	*wws.Client
}


// AdjustTimeOff (Adjust_Time_Off)
// 
// This operation will adjust an existing time off entry. The adjusted units must be passed in (and not the corrected units). This operation uses the Correct Time Off business process event.
func (c *Client) AdjustTimeOff(ctx context.Context, input *AdjustTimeOffInput) (output *AdjustTimeOffOutput, err error) {
	output = &AdjustTimeOffOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AdjustTimeOffInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Adjust_Time_Off_Request"`
	AdjustTimeOffRequestType
}

type AdjustTimeOffOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Time_Off_Event_Response"`
	TimeOffEventResponseType
}

// EnterTimeOff (Enter_Time_Off)
// 
// This operation will add new time off entries. It uses the Request Time Off business process event.
func (c *Client) EnterTimeOff(ctx context.Context, input *EnterTimeOffInput) (output *EnterTimeOffOutput, err error) {
	output = &EnterTimeOffOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EnterTimeOffInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Enter_Time_Off_Request"`
	EnterTimeOffRequestType
}

type EnterTimeOffOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Time_Off_Event_Response"`
	TimeOffEventResponseType
}

// PutAbsenceInput (Put_Absence_Input)
// 
// This operation adds a new Accrual or Time Off Adjustment/Override for an Employee (or updates an existing Adjustment/Override) with the supplied information.
// 
// NOTE : This web service went through a breaking change for Workday 16.  In Workday 14 and 15, unbounded Absence_Input_Line_Data elements were allowed within an Absence_Input_Data element.  Going forward, Workday 16 does use Absence_Input_Line_Data elements and the data has been moved to the Absence_Input_Data parent element.  If the older Workday schema is submitted, the first Absence_Input_Line_Data element is used and transformed internally to the updated schema.  All nested Absence_Input_Line_Data elements beyond the first are ignored.
func (c *Client) PutAbsenceInput(ctx context.Context, input *PutAbsenceInputInput) (output *PutAbsenceInputOutput, err error) {
	output = &PutAbsenceInputOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutAbsenceInputInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Absence_Input_Request"`
	PutAbsenceInputRequestType
}

type PutAbsenceInputOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Absence_Input_Response"`
	PutAbsenceInputResponseType
}

// GetAbsenceInputs (Get_Absence_Inputs)
// 
// This operation will get information for Accrual or Time Off Adjustments/Overrides for employees and/or batch id&#39;s specified in the request or information for all if no criteria is specified in the request.
func (c *Client) GetAbsenceInputs(ctx context.Context, input *GetAbsenceInputsInput) (output *GetAbsenceInputsOutput, err error) {
	output = &GetAbsenceInputsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAbsenceInputsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Absence_Inputs_Request"`
	GetAbsenceInputsRequestType
}

type GetAbsenceInputsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Absence_Inputs_Response"`
	GetAbsenceInputsResponseType
}

// RequestLeaveofAbsence (Request_Leave_of_Absence)
// 
// This operation will add new leave of absence requests. It uses the Request Leave of Absence business process event and its sub business processes.
func (c *Client) RequestLeaveofAbsence(ctx context.Context, input *RequestLeaveofAbsenceInput) (output *RequestLeaveofAbsenceOutput, err error) {
	output = &RequestLeaveofAbsenceOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type RequestLeaveofAbsenceInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Request_Leave_of_Absence_Request"`
	RequestLeaveofAbsenceRequestType
}

type RequestLeaveofAbsenceOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Request_Leave_of_Absence_Response"`
	RequestLeaveofAbsenceResponseType
}

// RequestReturnfromLeaveofAbsence (Request_Return_from_Leave_of_Absence)
// 
// This operation will add new return from leave of absence requests. It uses the Request Return from Leave of Absence business process event and its sub business processes.
func (c *Client) RequestReturnfromLeaveofAbsence(ctx context.Context, input *RequestReturnfromLeaveofAbsenceInput) (output *RequestReturnfromLeaveofAbsenceOutput, err error) {
	output = &RequestReturnfromLeaveofAbsenceOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type RequestReturnfromLeaveofAbsenceInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Request_Return_from_Leave_of_Absence_Request"`
	RequestReturnfromLeaveofAbsenceRequestType
}

type RequestReturnfromLeaveofAbsenceOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Request_Return_from_Leave_of_Absence_Response"`
	RequestReturnfromLeaveofAbsenceResponseType
}

// GetOverrideBalances (Get_Override_Balances)
// 
// This operation sets Time Off Plan Override Balances for an Employee and Time Off Plan with the supplied information.
func (c *Client) GetOverrideBalances(ctx context.Context, input *GetOverrideBalancesInput) (output *GetOverrideBalancesOutput, err error) {
	output = &GetOverrideBalancesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetOverrideBalancesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Override_Balances_Request"`
	GetOverrideBalancesRequestType
}

type GetOverrideBalancesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Override_Balances_Response"`
	GetOverrideBalancesResponseType
}

// PutOverrideBalance (Put_Override_Balance)
// 
// This operation will add or update a new Time Off Plan Override Balance record for a given Employee and Time Off Plan.
func (c *Client) PutOverrideBalance(ctx context.Context, input *PutOverrideBalanceInput) (output *PutOverrideBalanceOutput, err error) {
	output = &PutOverrideBalanceOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutOverrideBalanceInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Override_Balance_Request"`
	PutOverrideBalanceRequestType
}

type PutOverrideBalanceOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Override_Balance_Response"`
	PutOverrideBalanceResponseType
}

// GetTimeOffPlanBalances (Get_Time_Off_Plan_Balances)
// 
// This operation retrieves the dynamic Time Off Plan Balances by Worker and Time Off Plan.
func (c *Client) GetTimeOffPlanBalances(ctx context.Context, input *GetTimeOffPlanBalancesInput) (output *GetTimeOffPlanBalancesOutput, err error) {
	output = &GetTimeOffPlanBalancesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetTimeOffPlanBalancesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Time_Off_Plan_Balances_Request"`
	GetTimeOffPlanBalancesRequestType
}

type GetTimeOffPlanBalancesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Time_Off_Plan_Balances_Response"`
	GetTimeOffPlanBalancesResponseType
}

// GetCarryoverOverride (Get_Carryover_Override)
// 
// This operation gets Carryover Expiration/Limit Override for a ~Worker~ and Time Off Plan with the supplied information.
func (c *Client) GetCarryoverOverride(ctx context.Context, input *GetCarryoverOverrideInput) (output *GetCarryoverOverrideOutput, err error) {
	output = &GetCarryoverOverrideOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCarryoverOverrideInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Carryover_Override_Request"`
	GetCarryoverOverrideRequestType
}

type GetCarryoverOverrideOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Carryover_Override_Response"`
	GetCarryoverOverrideResponseType
}

// PutCarryoverOverride (Put_Carryover_Override)
// 
// This operation will add or update a new Carryover Expiration/Limit Override record for a given ~Worker~ and Time Off Plan.
func (c *Client) PutCarryoverOverride(ctx context.Context, input *PutCarryoverOverrideInput) (output *PutCarryoverOverrideOutput, err error) {
	output = &PutCarryoverOverrideOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCarryoverOverrideInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Carryover_Override_Request"`
	PutCarryoverOverrideRequestType
}

type PutCarryoverOverrideOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Carryover_Override_Response"`
	PutCarryoverOverrideResponseType
}

// ImportTimeOffRequestEventBatch (Import_Time_Off_Request_Event_Batch)
// 
// Imports a batch of Time Off Request Events that are processed in a background process which persists only those that pass validation.
func (c *Client) ImportTimeOffRequestEventBatch(ctx context.Context, input *ImportTimeOffRequestEventBatchInput) (output *ImportTimeOffRequestEventBatchOutput, err error) {
	output = &ImportTimeOffRequestEventBatchOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportTimeOffRequestEventBatchInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Time_Off_Request_Event_Batch_Request"`
	ImportTimeOffRequestEventBatchRequestType
}

type ImportTimeOffRequestEventBatchOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

