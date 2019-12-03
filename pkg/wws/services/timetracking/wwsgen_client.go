
// Package timetracking
//
// Operations for importing and exporting time and work schedule information.
package timetracking

import (
	"context"

	"github.com/ianlopshire/go-workday/pkg/wws"
)

const (
	CurrentVersion = "v31.2"
	ServiceName = "Time_Tracking"
)

type Client struct {
	*wws.Client
}


// PutTimeClockEvents (Put_Time_Clock_Events)
// 
// Add time clock events from a 3rd party time collection system.  Workday Time Tracking will process these time clock events into reported and calculated time blocks.  The Time Clock Event Time and Clock Event Type Reference will be verified immediately.  All other attributes will be persisted, but will be validated by subsequent, non-web service, business processes.
func (c *Client) PutTimeClockEvents(ctx context.Context, input *PutTimeClockEventsInput) (output *PutTimeClockEventsOutput, err error) {
	output = &PutTimeClockEventsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutTimeClockEventsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Time_Clock_Events_Request"`
	PutTimeClockEventsRequestType
}

type PutTimeClockEventsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Time_Clock_Events_Response"`
	PutTimeClockEventsResponseType
}

// AssignWorkSchedule (Assign_Work_Schedule)
// 
// Imports work schedule calendar assignments. Workday initiates an Assign Work Schedule business process event for each successfully imported assignment.
func (c *Client) AssignWorkSchedule(ctx context.Context, input *AssignWorkScheduleInput) (output *AssignWorkScheduleOutput, err error) {
	output = &AssignWorkScheduleOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AssignWorkScheduleInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Work_Schedule_Request"`
	AssignWorkScheduleRequestType
}

type AssignWorkScheduleOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Work_Schedule_Response"`
	AssignWorkScheduleResponseType
}

// GetCalculatedTimeBlocks (Get_Calculated_Time_Blocks)
// 
// Returns calculated time blocks by Workday ID, Date Range, Worker, Status, or Supervisory Organization. Returns all calculated time blocks if no Workday ID or Date Range is provided.
func (c *Client) GetCalculatedTimeBlocks(ctx context.Context, input *GetCalculatedTimeBlocksInput) (output *GetCalculatedTimeBlocksOutput, err error) {
	output = &GetCalculatedTimeBlocksOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCalculatedTimeBlocksInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Calculated_Time_Blocks_Request"`
	GetCalculatedTimeBlocksRequestType
}

type GetCalculatedTimeBlocksOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Calculated_Time_Blocks_Response"`
	GetCalculatedTimeBlocksResponseType
}

// ImportReportedTimeBlocks (Import_Reported_Time_Blocks)
// 
// Import time blocks from a 3rd party system. Workday Time Tracking will process these time blocks into reported and calculated time blocks.
func (c *Client) ImportReportedTimeBlocks(ctx context.Context, input *ImportReportedTimeBlocksInput) (output *ImportReportedTimeBlocksOutput, err error) {
	output = &ImportReportedTimeBlocksOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportReportedTimeBlocksInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Reported_Time_Blocks_Request"`
	ImportReportedTimeBlocksRequestType
}

type ImportReportedTimeBlocksOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportTimeClockEvents (Import_Time_Clock_Events)
// 
// Load large batches of time clock events on an infrequent basis.
func (c *Client) ImportTimeClockEvents(ctx context.Context, input *ImportTimeClockEventsInput) (output *ImportTimeClockEventsOutput, err error) {
	output = &ImportTimeClockEventsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportTimeClockEventsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Time_Clock_Events_Request"`
	ImportTimeClockEventsRequestType
}

type ImportTimeClockEventsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportAdHocSchedules (Import_Ad_Hoc_Schedules)
// 
// Load large batches of Schedule Blocks on an infrequent basis.
func (c *Client) ImportAdHocSchedules(ctx context.Context, input *ImportAdHocSchedulesInput) (output *ImportAdHocSchedulesOutput, err error) {
	output = &ImportAdHocSchedulesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportAdHocSchedulesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Ad_Hoc_Schedules_Request"`
	ImportAdHocSchedulesRequestType
}

type ImportAdHocSchedulesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

