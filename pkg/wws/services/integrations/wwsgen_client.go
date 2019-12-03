
// Package integrations
//
// This Web Service contains operations related to all Integrations within the Workday system.
package integrations

import (
	"context"

	"github.com/ianlopshire/go-workday/pkg/wws"
)

const (
	CurrentVersion = "v31.2"
	ServiceName = "Integrations"
)

type Client struct {
	*wws.Client
}


// GetIntegrationSystems (Get_Integration_Systems)
// 
// This operation will retrieve data related to an Integration System.  Data returned within the response includes Services associated and enabled, Attribute/Map values, Service configurations and Custom Services Attributes, Maps and Launch Parameters.
func (c *Client) GetIntegrationSystems(ctx context.Context, input *GetIntegrationSystemsInput) (output *GetIntegrationSystemsOutput, err error) {
	output = &GetIntegrationSystemsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetIntegrationSystemsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Integration_Systems_Request"`
	GetIntegrationSystemsRequestType
}

type GetIntegrationSystemsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Integration_Systems_Response"`
	GetIntegrationSystemsResponseType
}

// PutIntegrationMessage (Put_Integration_Message)
// 
// This operation adds a message to either an Integration System or Integration Event which allows for status, information and error details to be communicated back to Workday.
func (c *Client) PutIntegrationMessage(ctx context.Context, input *PutIntegrationMessageInput) (output *PutIntegrationMessageOutput, err error) {
	output = &PutIntegrationMessageOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutIntegrationMessageInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Integration_Message_Request"`
	PutIntegrationMessageRequestType
}

type PutIntegrationMessageOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Integration_Message_Response"`
	PutIntegrationMessageResponseType
}

// PutIntegrationEvent (Put_Integration_Event)
// 
// This operation will create and/or update data related to an Integration Event.  Data included will be Integration System, Launch datetime, Triggering User, Status and Runtime Parameters.
func (c *Client) PutIntegrationEvent(ctx context.Context, input *PutIntegrationEventInput) (output *PutIntegrationEventOutput, err error) {
	output = &PutIntegrationEventOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutIntegrationEventInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Integration_Event_Request"`
	PutIntegrationEventRequestType
}

type PutIntegrationEventOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Integration_Event_Response"`
	PutIntegrationEventResponseType
}

// GetEventDetail (Get_Event_Detail)
// 
// This operation will retrieve data (e.g. Creation Date, Event State, Event Target, etc..) related to Workday Business Process Event (e.g. Hire Employee, Enter Time Off, etc...)
func (c *Client) GetEventDetail(ctx context.Context, input *GetEventDetailInput) (output *GetEventDetailOutput, err error) {
	output = &GetEventDetailOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEventDetailInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Event_Details_Request"`
	GetEventDetailsRequestType
}

type GetEventDetailOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Event_Details_Response"`
	GetEventDetailsResponseType
}

// GetSequenceGenerators (Get_Sequence_Generators)
// 
// This operation will retrieve data related to a Sequence Generator (e.g ID Definition).  Data associated to a Sequence Generator includes the formatting, syntax, last date used, increment value, etc.  Sequence Generators are typically associated with Integrations utilizing a Sequence Generator Service.
func (c *Client) GetSequenceGenerators(ctx context.Context, input *GetSequenceGeneratorsInput) (output *GetSequenceGeneratorsOutput, err error) {
	output = &GetSequenceGeneratorsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetSequenceGeneratorsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Sequence_Generators_Request"`
	GetSequenceGeneratorsRequestType
}

type GetSequenceGeneratorsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Sequence_Generators_Response"`
	GetSequenceGeneratorsResponseType
}

// PutSequenceGenerator (Put_Sequence_Generator)
// 
// This operation will create and/or update data related to a Sequence Generator (e.g ID Definition).  Data associated to a Sequence Generator includes the formatting, syntax, last date used, increment value, etc.  Sequence Generators are typically associated with Integrations utilizing a Sequence Generator Service.
func (c *Client) PutSequenceGenerator(ctx context.Context, input *PutSequenceGeneratorInput) (output *PutSequenceGeneratorOutput, err error) {
	output = &PutSequenceGeneratorOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutSequenceGeneratorInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Sequence_Generator_Request"`
	PutSequenceGeneratorRequestType
}

type PutSequenceGeneratorOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Sequence_Generator_Response"`
	PutSequenceGeneratorResponseType
}

// IncrementSequenceGenerator (Increment_Sequence_Generator)
// 
// This operation will increments a Sequence Generator and persists the data within Workday.  The response is the fully formatted value of the Sequence Generator after the increment.
func (c *Client) IncrementSequenceGenerator(ctx context.Context, input *IncrementSequenceGeneratorInput) (output *IncrementSequenceGeneratorOutput, err error) {
	output = &IncrementSequenceGeneratorOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type IncrementSequenceGeneratorInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Increment_Sequence_Generator_Request"`
	IncrementSequenceGeneratorRequestType
}

type IncrementSequenceGeneratorOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Increment_Sequence_Generator_Response"`
	IncrementSequenceGeneratorResponseType
}

// PutIntegrationSystem (Put_Integration_System)
// 
// This operation will create and/or update data related to an Integration System.  Data returned within the response includes Services associated and enabled, Attribute/Map values, Service configurations and Custom Services Attributes, Maps and Launch Parameters.
func (c *Client) PutIntegrationSystem(ctx context.Context, input *PutIntegrationSystemInput) (output *PutIntegrationSystemOutput, err error) {
	output = &PutIntegrationSystemOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutIntegrationSystemInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Integration_System_Request"`
	PutIntegrationSystemRequestType
}

type PutIntegrationSystemOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Integration_System_Response"`
	PutIntegrationSystemResponseType
}

// GetIntegrationEvents (Get_Integration_Events)
// 
// This operation will retrieve data related to an execution of an Integration Event.  Data included will be Integration System, Launch datetime, Triggering User, Status and Runtime Parameters.
func (c *Client) GetIntegrationEvents(ctx context.Context, input *GetIntegrationEventsInput) (output *GetIntegrationEventsOutput, err error) {
	output = &GetIntegrationEventsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetIntegrationEventsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Integration_Events_Request"`
	GetIntegrationEventsRequestType
}

type GetIntegrationEventsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Integration_Events_Response"`
	GetIntegrationEventsResponseType
}

// CancelBusinessProcess (Cancel_Business_Process)
// 
// This operation cancels a business process, if cancelable, by the user
func (c *Client) CancelBusinessProcess(ctx context.Context, input *CancelBusinessProcessInput) (output *CancelBusinessProcessOutput, err error) {
	output = &CancelBusinessProcessOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CancelBusinessProcessInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Cancel_Business_Process_Request"`
	CancelBusinessProcessRequestType
}

type CancelBusinessProcessOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Cancel_Business_Process_Response"`
	CancelBusinessProcessResponseType
}

// LaunchIntegration (Launch_Integration)
// 
// This operation will Launch an Integration.  The request has the ability to accept Runtime Parameters and the Workday Account that should be used within the integration processing.  The result of this operation will be the creation of an Integration Event along with the processing of the integration itself.
func (c *Client) LaunchIntegration(ctx context.Context, input *LaunchIntegrationInput) (output *LaunchIntegrationOutput, err error) {
	output = &LaunchIntegrationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type LaunchIntegrationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Launch_Integration_Event_Request"`
	LaunchIntegrationEventRequestType
}

type LaunchIntegrationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Launch_Integration_Event_Response"`
	LaunchIntegrationEventResponseType
}

// RescindBusinessProcess (Rescind_Business_Process)
// 
// This operation rescinds a business process, if rescindable by the user
func (c *Client) RescindBusinessProcess(ctx context.Context, input *RescindBusinessProcessInput) (output *RescindBusinessProcessOutput, err error) {
	output = &RescindBusinessProcessOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type RescindBusinessProcessInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Rescind_Business_Process_Request"`
	RescindBusinessProcessRequestType
}

type RescindBusinessProcessOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Rescind_Business_Process_Response"`
	RescindBusinessProcessResponseType
}

// GetIntegrationSystemUsers (Get_Integration_System_Users)
// 
// This operation will retrieve data related to Workday Accounts associated to an Integration System.
func (c *Client) GetIntegrationSystemUsers(ctx context.Context, input *GetIntegrationSystemUsersInput) (output *GetIntegrationSystemUsersOutput, err error) {
	output = &GetIntegrationSystemUsersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetIntegrationSystemUsersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Integration_System_Users_Request"`
	GetIntegrationSystemUsersRequestType
}

type GetIntegrationSystemUsersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Integration_System_Users_Response"`
	GetIntegrationSystemUsersResponseType
}

// PutIntegrationSystemUser (Put_Integration_System_User)
// 
// This operation will create and/or update data related to Workday Accounts associated to an Integration System.
func (c *Client) PutIntegrationSystemUser(ctx context.Context, input *PutIntegrationSystemUserInput) (output *PutIntegrationSystemUserOutput, err error) {
	output = &PutIntegrationSystemUserOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutIntegrationSystemUserInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Integration_System_User_Request"`
	PutIntegrationSystemUserRequestType
}

type PutIntegrationSystemUserOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Integration_System_User_Response"`
	PutIntegrationSystemUserResponseType
}

// GetSubscriptions (Get_Subscriptions)
// 
// This operation will retrieve data related to Subscriptions, both for an Integration System and for Transaction Log service configuration.  A Subscription includes an association with 1 or more Transaction/Business Process Types.
func (c *Client) GetSubscriptions(ctx context.Context, input *GetSubscriptionsInput) (output *GetSubscriptionsOutput, err error) {
	output = &GetSubscriptionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetSubscriptionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Subscriptions_Request"`
	GetSubscriptionsRequestType
}

type GetSubscriptionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Subscriptions_Response"`
	GetSubscriptionsResponseType
}

// PutSubscription (Put_Subscription)
// 
// This operation will create and/or update data related to Subscriptions, both for an Integration System and for Transaction Log service configuration.  A Subscription includes an association with 1 or more Transaction/Business Process Types.
func (c *Client) PutSubscription(ctx context.Context, input *PutSubscriptionInput) (output *PutSubscriptionOutput, err error) {
	output = &PutSubscriptionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutSubscriptionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Subscription_Request"`
	PutSubscriptionRequestType
}

type PutSubscriptionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Subscription_Response"`
	PutSubscriptionResponseType
}

// LaunchEIB (Launch_EIB)
// 
// This operation will Launch an EIB Integration.  The request has the ability to accept override Runtime Parameters that should be used within the integration processing.  If no override is provided within the request, the System Default, which has been defined on the EIB Definition, will be used.  The result of this operation will be the creation of an Integration Event along with the processing of the integration itself.
func (c *Client) LaunchEIB(ctx context.Context, input *LaunchEIBInput) (output *LaunchEIBOutput, err error) {
	output = &LaunchEIBOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type LaunchEIBInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Launch_EIB_Request"`
	LaunchEIBRequestType
}

type LaunchEIBOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Launch_EIB_Response"`
	LaunchEIBResponseType
}

// GetReferences (Get_References)
// 
// This operation provide a list of values for a business object in Workday given its &#34;Reference ID Type&#34;. The &#34;Reference ID Type&#34; for each referenced business object can be retrieved within our API Documentation. For example, the &#34;Reference ID Type&#34; for a Location reference is &#34;Location_ID&#34;. So, if we call &#34;Get References&#34; with &#34;Location_ID&#34; you will get back the list of all locations (and their Workday and Reference IDs).
func (c *Client) GetReferences(ctx context.Context, input *GetReferencesInput) (output *GetReferencesOutput, err error) {
	output = &GetReferencesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetReferencesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_References_Request"`
	GetReferencesRequestType
}

type GetReferencesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_References_Response"`
	GetReferencesResponseType
}

// GetEventDocuments (Get_Event_Documents)
// 
// This operation will retrieve data related to any Documents that may be available to an Integration Event.
func (c *Client) GetEventDocuments(ctx context.Context, input *GetEventDocumentsInput) (output *GetEventDocumentsOutput, err error) {
	output = &GetEventDocumentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEventDocumentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Event_Documents_Request"`
	GetEventDocumentsRequestType
}

type GetEventDocumentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Event_Documents_Response"`
	GetEventDocumentsResponseType
}

// PutReference (Put_Reference)
// 
// This operation updates the Reference IDs on a business object for a Reference ID Type.  Currently, updating Reference IDs for most, but not all Reference ID Types is supported.
func (c *Client) PutReference(ctx context.Context, input *PutReferenceInput) (output *PutReferenceOutput, err error) {
	output = &PutReferenceOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutReferenceInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Reference_Request"`
	PutReferenceRequestType
}

type PutReferenceOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Reference_Response"`
	PutReferenceResponseType
}

// ApproveBusinessProcess (Approve_Business_Process)
// 
// This operation approves a business process, if approvable, by the user. If Disable Comment is configured in the coressponding Business Process Policy, comments provided in the request would not be saved.
func (c *Client) ApproveBusinessProcess(ctx context.Context, input *ApproveBusinessProcessInput) (output *ApproveBusinessProcessOutput, err error) {
	output = &ApproveBusinessProcessOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ApproveBusinessProcessInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Approve_Business_Process_Request"`
	ApproveBusinessProcessRequestType
}

type ApproveBusinessProcessOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Approve_Business_Process_Response"`
	ApproveBusinessProcessResponseType
}

// DenyBusinessProcess (Deny_Business_Process)
// 
// This operation denies a business process, if deniable, by the user
func (c *Client) DenyBusinessProcess(ctx context.Context, input *DenyBusinessProcessInput) (output *DenyBusinessProcessOutput, err error) {
	output = &DenyBusinessProcessOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type DenyBusinessProcessInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Deny_Business_Process_Request"`
	DenyBusinessProcessRequestType
}

type DenyBusinessProcessOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Deny_Business_Process_Response"`
	DenyBusinessProcessResponseType
}

// GetImportProcessMessages (Get_Import_Process_Messages)
// 
// Gets messages associated to a specific Import Process event.
func (c *Client) GetImportProcessMessages(ctx context.Context, input *GetImportProcessMessagesInput) (output *GetImportProcessMessagesOutput, err error) {
	output = &GetImportProcessMessagesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetImportProcessMessagesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Import_Process_Messages_Request"`
	GetImportProcessMessagesRequestType
}

type GetImportProcessMessagesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Import_Process_Messages_Response"`
	GetImportProcessMessagesResponseType
}

// GetImportProcesses (Get_Import_Processes)
// 
// Get status and correlation information for Import Process load
func (c *Client) GetImportProcesses(ctx context.Context, input *GetImportProcessesInput) (output *GetImportProcessesOutput, err error) {
	output = &GetImportProcessesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetImportProcessesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Import_Processes_Request"`
	GetImportProcessesRequestType
}

type GetImportProcessesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Import_Processes_Response"`
	GetImportProcessesResponseType
}

// ReassignBusinessProcessStep (Reassign_Business_Process_Step)
// 
// This operation reassign a business process step, if reassignable by the processing user and the new assignee has the required Roles to complete the step.
func (c *Client) ReassignBusinessProcessStep(ctx context.Context, input *ReassignBusinessProcessStepInput) (output *ReassignBusinessProcessStepOutput, err error) {
	output = &ReassignBusinessProcessStepOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ReassignBusinessProcessStepInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Reassign_Business_Process_Step_Request"`
	ReassignBusinessProcessStepRequestType
}

type ReassignBusinessProcessStepOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Reassign_Business_Process_Step_Response"`
	ReassignBusinessProcessStepResponseType
}

