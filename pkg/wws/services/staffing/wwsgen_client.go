
// Package staffing
//
// The Staffing Web Service operations expose Workday Human Capital Management Business Services and data. These services pertain to staffing transactions for both employees and contingent workers, such as bringing employees and contingent workers on board.
package staffing

import (
	"context"

	"github.com/ianlopshire/go-workday/pkg/wws"
)

const (
	CurrentVersion = "v31.2"
	ServiceName = "Staffing"
)

type Client struct {
	*wws.Client
}


// TerminateEmployee (Terminate_Employee)
// 
// Terminates an employee.
func (c *Client) TerminateEmployee(ctx context.Context, input *TerminateEmployeeInput) (output *TerminateEmployeeOutput, err error) {
	output = &TerminateEmployeeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type TerminateEmployeeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Terminate_Employee_Request"`
	TerminateEmployeeRequestType
}

type TerminateEmployeeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Terminate_Employee_Event_Response"`
	TerminateEmployeeEventResponseType
}

// GetWorkers (Get_Workers)
// 
// Returns public and private information for specified workers.
func (c *Client) GetWorkers(ctx context.Context, input *GetWorkersInput) (output *GetWorkersOutput, err error) {
	output = &GetWorkersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetWorkersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Workers_Request"`
	GetWorkersRequestType
}

type GetWorkersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Workers_Response"`
	GetWorkersResponseRootType
}

// EditPosition (Edit_Position)
// 
// Edits a filled position. Uses the Edit Position business process.
func (c *Client) EditPosition(ctx context.Context, input *EditPositionInput) (output *EditPositionOutput, err error) {
	output = &EditPositionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EditPositionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Position_Request"`
	EditPositionRequestType
}

type EditPositionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Position_Event_Response"`
	EditPositionEventResponseType
}

// HireEmployee (Hire_Employee)
// 
// Hires a pre-hire (new or existing) into an employee position, headcount, or job. Uses the Hire Employee business process. Note: pre-hire was previously called applicant. However, the web service operation and its elements are not yet renamed, so that existing integrations continue to work.
func (c *Client) HireEmployee(ctx context.Context, input *HireEmployeeInput) (output *HireEmployeeOutput, err error) {
	output = &HireEmployeeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type HireEmployeeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Hire_Employee_Request"`
	HireEmployeeRequestType
}

type HireEmployeeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Hire_Employee_Event_Response"`
	HireEmployeeEventResponseType
}

// GetApplicants (Get_Applicants)
// 
// Returns information for pre-hires specified in the request. If the request does not specify a pre-hire, this operation returns information for all pre-hires. Note: pre-hire was previously called applicant. However, the web service operation and its elements are not yet renamed, so that existing integrations continue to work.
func (c *Client) GetApplicants(ctx context.Context, input *GetApplicantsInput) (output *GetApplicantsOutput, err error) {
	output = &GetApplicantsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetApplicantsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Applicants_Request"`
	GetApplicantsRequestType
}

type GetApplicantsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Applicants_Response"`
	GetApplicantsResponseType
}

// AssignOrganization (Assign_Organization)
// 
// Assigns company, cost center, region, fund, grant, business unit, program, and custom organizations configured for staffing usage to a staffing position. Uses the Change Organization Assignments for Worker business process.
func (c *Client) AssignOrganization(ctx context.Context, input *AssignOrganizationInput) (output *AssignOrganizationOutput, err error) {
	output = &AssignOrganizationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AssignOrganizationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Organization_Request"`
	AssignOrganizationRequestType
}

type AssignOrganizationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Organization_Response"`
	AssignOrganizationResponseType
}

// PutApplicant (Put_Applicant)
// 
// Adds a new pre-hire (or updates an existing pre-hire) with the information supplied in the request. Note: pre-hire was previously called applicant. However, the web service operation and its elements are not yet renamed, so that existing integrations continue to work.
func (c *Client) PutApplicant(ctx context.Context, input *PutApplicantInput) (output *PutApplicantOutput, err error) {
	output = &PutApplicantOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutApplicantInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Applicant_Request"`
	PutApplicantRequestType
}

type PutApplicantOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Applicant_Response"`
	PutApplicantResponseType
}

// EndContingentWorkerContract (End_Contingent_Worker_Contract)
// 
// Ends a contingent worker&#39;s contract. Uses the End Contingent Worker Contract Business Process.
func (c *Client) EndContingentWorkerContract(ctx context.Context, input *EndContingentWorkerContractInput) (output *EndContingentWorkerContractOutput, err error) {
	output = &EndContingentWorkerContractOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EndContingentWorkerContractInput struct {
	XMLName string `xml:"urn:com.workday/bsvc End_Contingent_Worker_Contract_Request"`
	EndContingentWorkerContractRequestType
}

type EndContingentWorkerContractOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc End_Contingent_Worker_Contract_Response"`
	EndContingentWorkerContractResponseType
}

// ContractContingentWorker (Contract_Contingent_Worker)
// 
// Contracts an existing pre-hire into a contingent worker position or job. Uses the Contract Contingent Worker business process. Note: pre-hire was previously called applicant. However, the web service operation and its elements are not yet renamed, so that existing integrations continue to work.
func (c *Client) ContractContingentWorker(ctx context.Context, input *ContractContingentWorkerInput) (output *ContractContingentWorkerOutput, err error) {
	output = &ContractContingentWorkerOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ContractContingentWorkerInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Contract_Contingent_Worker_Request"`
	ContractContingentWorkerRequestType
}

type ContractContingentWorkerOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Contract_Contingent_Worker_Response"`
	ContractContingentWorkerResponseType
}

// PutDependent (Put_Dependent)
// 
// DEPRECATED: 	Adds or updates a dependent
func (c *Client) PutDependent(ctx context.Context, input *PutDependentInput) (output *PutDependentOutput, err error) {
	output = &PutDependentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutDependentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Dependent_Request"`
	PutDependentRequestType
}

type PutDependentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Dependent_Response"`
	PutDependentResponseType
}

// CreatePosition (Create_Position)
// 
// Creates/opens a position for a supervisory organization using the position management staffing model. Uses the Create Position business process. There are several sub-operations within this operation.
// 
// Existing positions cannot be re-loaded into the system. You must use the Edit Position Restrictions operation to change a created position.
func (c *Client) CreatePosition(ctx context.Context, input *CreatePositionInput) (output *CreatePositionOutput, err error) {
	output = &CreatePositionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CreatePositionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Create_Position_Request"`
	CreatePositionRequestType
}

type CreatePositionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Create_Position_Response"`
	CreatePositionResponseType
}

// GetPositions (Get_Positions)
// 
// Returns information for position management positions. If a position is filled, identifies the worker filling the position.
func (c *Client) GetPositions(ctx context.Context, input *GetPositionsInput) (output *GetPositionsOutput, err error) {
	output = &GetPositionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPositionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Positions_Request"`
	GetPositionsRequestType
}

type GetPositionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Positions_Response"`
	GetPositionsResponseType
}

// EditServiceDates (Edit_Service_Dates)
// 
// Sets a worker&#39;s service dates. Uses the Edit Service Dates business process.
func (c *Client) EditServiceDates(ctx context.Context, input *EditServiceDatesInput) (output *EditServiceDatesOutput, err error) {
	output = &EditServiceDatesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EditServiceDatesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Service_Dates_Request"`
	EditServiceDatesRequestType
}

type EditServiceDatesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Service_Dates_Response"`
	EditServiceDatesResponseType
}

// GetOrganizations (Get_Organizations)
// 
// Returns organization information for a type of organization. If the request does not specify an organization, the operation returns information for all organizations.
func (c *Client) GetOrganizations(ctx context.Context, input *GetOrganizationsInput) (output *GetOrganizationsOutput, err error) {
	output = &GetOrganizationsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetOrganizationsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Organizations_Request"`
	GetOrganizationsRequestType
}

type GetOrganizationsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Organizations_Response"`
	GetOrganizationsResponseType
}

// EndAdditionalJob (End_Additional_Job)
// 
// Ends an employee&#39;s additional job. Uses the End Additional Job business process.
func (c *Client) EndAdditionalJob(ctx context.Context, input *EndAdditionalJobInput) (output *EndAdditionalJobOutput, err error) {
	output = &EndAdditionalJobOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EndAdditionalJobInput struct {
	XMLName string `xml:"urn:com.workday/bsvc End_Additional_Job_Request"`
	EndAdditionalJobRequestType
}

type EndAdditionalJobOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc End_Additional_Employee_Job_Event_Response"`
	EndAdditionalEmployeeJobEventResponseType
}

// AddAdditionalJob (Add_Additional_Job)
// 
// Adds an additional job for an existing employee. Uses the Add Additional Job business process.
func (c *Client) AddAdditionalJob(ctx context.Context, input *AddAdditionalJobInput) (output *AddAdditionalJobOutput, err error) {
	output = &AddAdditionalJobOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AddAdditionalJobInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Add_Additional_Job_Request"`
	AddAdditionalJobRequestType
}

type AddAdditionalJobOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Add_Additional_Job_Event_Response"`
	AddAdditionalJobEventResponseType
}

// AddRetireeStatus (Add_Retiree_Status)
// 
// Changes the status of a previously terminated worker to retiree. Uses the Add Retiree Status business process.
func (c *Client) AddRetireeStatus(ctx context.Context, input *AddRetireeStatusInput) (output *AddRetireeStatusOutput, err error) {
	output = &AddRetireeStatusOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AddRetireeStatusInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Add_Retiree_Status_Request"`
	AddRetireeStatusRequestType
}

type AddRetireeStatusOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Add_Retiree_Status_Response"`
	AddRetireeStatusResponseType
}

// RemoveRetireeStatus (Remove_Retiree_Status)
// 
// Removes the retirement status from a retired employee. Uses the Remove Retiree Status business process.
func (c *Client) RemoveRetireeStatus(ctx context.Context, input *RemoveRetireeStatusInput) (output *RemoveRetireeStatusOutput, err error) {
	output = &RemoveRetireeStatusOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type RemoveRetireeStatusInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Remove_Retiree_Status_Request"`
	RemoveRetireeStatusRequestType
}

type RemoveRetireeStatusOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Remove_Retiree_Status_Response"`
	RemoveRetireeStatusResponseType
}

// MaintainEmployeeContracts (Maintain_Employee_Contracts)
// 
// Adds or updates employee contracts. Uses the Maintain Employee Contracts business process.
func (c *Client) MaintainEmployeeContracts(ctx context.Context, input *MaintainEmployeeContractsInput) (output *MaintainEmployeeContractsOutput, err error) {
	output = &MaintainEmployeeContractsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type MaintainEmployeeContractsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Maintain_Employee_Contracts_Request"`
	MaintainEmployeeContractsRequestType
}

type MaintainEmployeeContractsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Maintain_Employee_Contracts_Response"`
	MaintainEmployeeContractsResponseType
}

// GetCitizenshipStatuses (Get_Citizenship_Statuses)
// 
// Returns citizenship status details for the requested status. If the request does not specify a citizenship status, the operations returns details for all citizenship statuses.
func (c *Client) GetCitizenshipStatuses(ctx context.Context, input *GetCitizenshipStatusesInput) (output *GetCitizenshipStatusesOutput, err error) {
	output = &GetCitizenshipStatusesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCitizenshipStatusesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Citizenship_Statuses_Request"`
	GetCitizenshipStatusesRequestType
}

type GetCitizenshipStatusesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Citizenship_Statuses_Response"`
	GetCitizenshipStatusesResponseType
}

// PutCitizenshipStatus (Put_Citizenship_Status)
// 
// Adds or updates citizenship status.
func (c *Client) PutCitizenshipStatus(ctx context.Context, input *PutCitizenshipStatusInput) (output *PutCitizenshipStatusOutput, err error) {
	output = &PutCitizenshipStatusOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCitizenshipStatusInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Citizenship_Status_Request"`
	PutCitizenshipStatusRequestType
}

type PutCitizenshipStatusOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Citizenship_Status_Response"`
	PutCitizenshipStatusResponseType
}

// EditPositionRestrictions (Edit_Position_Restrictions)
// 
// Edits an existing position restriction. Uses the Edit Position Restrictions business process.
func (c *Client) EditPositionRestrictions(ctx context.Context, input *EditPositionRestrictionsInput) (output *EditPositionRestrictionsOutput, err error) {
	output = &EditPositionRestrictionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EditPositionRestrictionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Position_Restrictions_Request"`
	EditPositionRestrictionsRequestType
}

type EditPositionRestrictionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Positon_Restriction_Response"`
	EditPositonRestrictionResponseType
}

// GetJobClassificationGroups (Get_Job_Classification_Groups)
// 
// Returns job classification group data.
func (c *Client) GetJobClassificationGroups(ctx context.Context, input *GetJobClassificationGroupsInput) (output *GetJobClassificationGroupsOutput, err error) {
	output = &GetJobClassificationGroupsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetJobClassificationGroupsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Classification_Groups_Request"`
	GetJobClassificationGroupsRequestType
}

type GetJobClassificationGroupsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Classification_Groups_Response"`
	GetJobClassificationGroupsResponseType
}

// PutJobClassificationGroup (Put_Job_Classification_Group)
// 
// Adds or updates a job classification group and/or job classifications.
func (c *Client) PutJobClassificationGroup(ctx context.Context, input *PutJobClassificationGroupInput) (output *PutJobClassificationGroupOutput, err error) {
	output = &PutJobClassificationGroupOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutJobClassificationGroupInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Classification_Group_Request"`
	PutJobClassificationGroupRequestType
}

type PutJobClassificationGroupOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Classification_Group_Response"`
	PutJobClassificationGroupResponseType
}

// GetJobFamilies (Get_Job_Families)
// 
// Returns job family data.
func (c *Client) GetJobFamilies(ctx context.Context, input *GetJobFamiliesInput) (output *GetJobFamiliesOutput, err error) {
	output = &GetJobFamiliesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetJobFamiliesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Families_Request"`
	GetJobFamiliesRequestType
}

type GetJobFamiliesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Families_Response"`
	GetJobFamiliesResponseType
}

// PutJobFamily (Put_Job_Family)
// 
// Adds or updates a job family.
func (c *Client) PutJobFamily(ctx context.Context, input *PutJobFamilyInput) (output *PutJobFamilyOutput, err error) {
	output = &PutJobFamilyOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutJobFamilyInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Family_Request"`
	PutJobFamilyRequestType
}

type PutJobFamilyOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Family_Response"`
	PutJobFamilyResponseType
}

// GetJobFamilyGroups (Get_Job_Family_Groups)
// 
// Returns job family and job family group data.
func (c *Client) GetJobFamilyGroups(ctx context.Context, input *GetJobFamilyGroupsInput) (output *GetJobFamilyGroupsOutput, err error) {
	output = &GetJobFamilyGroupsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetJobFamilyGroupsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Family_Groups_Request"`
	GetJobFamilyGroupsRequestType
}

type GetJobFamilyGroupsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Family_Groups_Response"`
	GetJobFamilyGroupsResponseType
}

// PutJobFamilyGroup (Put_Job_Family_Group)
// 
// Adds or updates a job family group. Also adds or removes a group&#39;s associated job families.
func (c *Client) PutJobFamilyGroup(ctx context.Context, input *PutJobFamilyGroupInput) (output *PutJobFamilyGroupOutput, err error) {
	output = &PutJobFamilyGroupOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutJobFamilyGroupInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Family_Group_Request"`
	PutJobFamilyGroupRequestType
}

type PutJobFamilyGroupOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Family_Group_Response"`
	PutJobFamilyGroupResponseType
}

// StartInternationalAssignment (Start_International_Assignment)
// 
// Start international assignment for an employee
func (c *Client) StartInternationalAssignment(ctx context.Context, input *StartInternationalAssignmentInput) (output *StartInternationalAssignmentOutput, err error) {
	output = &StartInternationalAssignmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type StartInternationalAssignmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Start_International_Assignment_Request"`
	StartInternationalAssignmentRequestType
}

type StartInternationalAssignmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Start_International_Assignment_Response"`
	StartInternationalAssignmentResponseType
}

// EndInternationalAssignment (End_International_Assignment)
// 
// End international assignment for an employee
func (c *Client) EndInternationalAssignment(ctx context.Context, input *EndInternationalAssignmentInput) (output *EndInternationalAssignmentOutput, err error) {
	output = &EndInternationalAssignmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EndInternationalAssignmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc End_International_Assignment_for_Employee_Event_Request"`
	EndInternationalAssignmentforEmployeeEventRequestType
}

type EndInternationalAssignmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc End_International_Assignment_for_Employee_Event_Response"`
	EndInternationalAssignmentforEmployeeEventResponseType
}

// AssignOrganizationRoles (Assign_Organization_Roles)
// 
// DEPRECATED: This web service operation is deprecated. Please use the Assign_Roles web service operation instead. If you do use this web service operation, note that all new and existing assignees must be specified as role assignees. Any existing role assignees that are not included in the set of role assignees will be removed. Later-dated role assignments will not be updated with new or removed assignees.
func (c *Client) AssignOrganizationRoles(ctx context.Context, input *AssignOrganizationRolesInput) (output *AssignOrganizationRolesOutput, err error) {
	output = &AssignOrganizationRolesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AssignOrganizationRolesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Request"`
	AssignOrganizationRolesRequestType
}

type AssignOrganizationRolesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Organization_Roles_Response"`
	AssignOrganizationRolesResponseType
}

// PutWorkerDocument (Put_Worker_Document)
// 
// Adds or updates a worker document. The operation adds documents not associated with events and those documents are not secured by the events.
func (c *Client) PutWorkerDocument(ctx context.Context, input *PutWorkerDocumentInput) (output *PutWorkerDocumentOutput, err error) {
	output = &PutWorkerDocumentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutWorkerDocumentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Worker_Document_Request"`
	PutWorkerDocumentRequestType
}

type PutWorkerDocumentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Worker_Document_Response"`
	PutWorkerDocumentResponseType
}

// GetWorkerDocuments (Get_Worker_Documents)
// 
// Returns worker document data.
func (c *Client) GetWorkerDocuments(ctx context.Context, input *GetWorkerDocumentsInput) (output *GetWorkerDocumentsOutput, err error) {
	output = &GetWorkerDocumentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetWorkerDocumentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Worker_Documents_Request"`
	GetWorkerDocumentsRequestType
}

type GetWorkerDocumentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Worker_Documents_Response"`
	GetWorkerDocumentsResponseType
}

// ChangeJob (Change_Job)
// 
// Perform a job change on an employee or contingent worker. Uses the Change Job business process. The types of changes include transfer, promotion, demotion, lateral moves and any other change of data on the job.
func (c *Client) ChangeJob(ctx context.Context, input *ChangeJobInput) (output *ChangeJobOutput, err error) {
	output = &ChangeJobOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeJobInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Job_Request"`
	ChangeJobRequestType
}

type ChangeJobOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Job_Response"`
	ChangeJobResponseType
}

// SwitchPrimaryJob (Switch_Primary_Job)
// 
// This operation will invoke the business process to switch a workers primary job with one of their additional jobs. Uses the Switch Primary Job business process.
func (c *Client) SwitchPrimaryJob(ctx context.Context, input *SwitchPrimaryJobInput) (output *SwitchPrimaryJobOutput, err error) {
	output = &SwitchPrimaryJobOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SwitchPrimaryJobInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Switch_Primary_Job_Request"`
	SwitchPrimaryJobRequestType
}

type SwitchPrimaryJobOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Switch_Primary_Job_Event_Response"`
	SwitchPrimaryJobEventResponseType
}

// ClosePosition (Close_Position)
// 
// This operation will invoke the Business Process to Close a Position.
func (c *Client) ClosePosition(ctx context.Context, input *ClosePositionInput) (output *ClosePositionOutput, err error) {
	output = &ClosePositionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ClosePositionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Close_Position_Request"`
	ClosePositionRequestType
}

type ClosePositionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Close_Position_Response"`
	ClosePositionResponseType
}

// GetChangeGovernmentIDs (Get_Change_Government_IDs)
// 
// This public web service request gets government IDs for a worker. The response can be used as input to the web service request Change Government IDs.
func (c *Client) GetChangeGovernmentIDs(ctx context.Context, input *GetChangeGovernmentIDsInput) (output *GetChangeGovernmentIDsOutput, err error) {
	output = &GetChangeGovernmentIDsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetChangeGovernmentIDsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Government_IDs_Request"`
	GetChangeGovernmentIDsRequestType
}

type GetChangeGovernmentIDsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Government_IDs_Response"`
	GetChangeGovernmentIDsResponseType
}

// GetChangePassportsandVisas (Get_Change_Passports_and_Visas)
// 
// This public web service request gets passport and visas for a worker. The response can be used as input to the web service request Change Passports and Visas.
func (c *Client) GetChangePassportsandVisas(ctx context.Context, input *GetChangePassportsandVisasInput) (output *GetChangePassportsandVisasOutput, err error) {
	output = &GetChangePassportsandVisasOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetChangePassportsandVisasInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Passports_and_Visas_Request"`
	GetChangePassportsandVisasRequestType
}

type GetChangePassportsandVisasOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Passports_and_Visas_Response"`
	GetChangePassportsandVisasResponseType
}

// GetChangeLicenses (Get_Change_Licenses)
// 
// This public web service request gets licenses for a worker. The response can be used as input to the web service request Change Licenses.
func (c *Client) GetChangeLicenses(ctx context.Context, input *GetChangeLicensesInput) (output *GetChangeLicensesOutput, err error) {
	output = &GetChangeLicensesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetChangeLicensesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Licenses_Request"`
	GetChangeLicensesRequestType
}

type GetChangeLicensesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Licenses_Response"`
	GetChangeLicensesResponseType
}

// GetChangeOtherIDs (Get_Change_Other_IDs)
// 
// This public web service request gets custom IDs for a worker. The response can be used as input to the web service request Change Other IDs.
func (c *Client) GetChangeOtherIDs(ctx context.Context, input *GetChangeOtherIDsInput) (output *GetChangeOtherIDsOutput, err error) {
	output = &GetChangeOtherIDsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetChangeOtherIDsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Other_IDs_Request"`
	GetChangeOtherIDsRequestType
}

type GetChangeOtherIDsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Other_IDs_Response"`
	GetChangeOtherIDsResponseType
}

// GetChangePreferredName (Get_Change_Preferred_Name)
// 
// This public web service request gets preferred name for a worker. The response can be used as input to the web service request Change Preferred Name. All response fields are secured to their specific granular domains.
func (c *Client) GetChangePreferredName(ctx context.Context, input *GetChangePreferredNameInput) (output *GetChangePreferredNameOutput, err error) {
	output = &GetChangePreferredNameOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetChangePreferredNameInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Preferred_Name_Request"`
	GetChangePreferredNameRequestType
}

type GetChangePreferredNameOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Preferred_Name_Response"`
	GetChangePreferredNameResponseType
}

// GetChangeLegalName (Get_Change_Legal_Name)
// 
// This public web service request gets the legal name for a worker. The response can be used as input to the web service request Change Legal Name. All response fields are secured to their specific granular domains.
func (c *Client) GetChangeLegalName(ctx context.Context, input *GetChangeLegalNameInput) (output *GetChangeLegalNameOutput, err error) {
	output = &GetChangeLegalNameOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetChangeLegalNameInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Legal_Name_Request"`
	GetChangeLegalNameRequestType
}

type GetChangeLegalNameOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Legal_Name_Response"`
	GetChangeLegalNameResponseType
}

// SetHiringRestrictions (Set_Hiring_Restrictions)
// 
// This operation will create the hiring restrictions for a job management supervisory organization.
func (c *Client) SetHiringRestrictions(ctx context.Context, input *SetHiringRestrictionsInput) (output *SetHiringRestrictionsOutput, err error) {
	output = &SetHiringRestrictionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SetHiringRestrictionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Set_Hiring_Restrictions_Request"`
	SetHiringRestrictionsRequestType
}

type SetHiringRestrictionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Set_Hiring_Restrictions_Response"`
	SetHiringRestrictionsResponseType
}

// EditHiringRestrictions (Edit_Hiring_Restrictions)
// 
// This operation will edit the hiring restrictions for a job management supervisory organization.
func (c *Client) EditHiringRestrictions(ctx context.Context, input *EditHiringRestrictionsInput) (output *EditHiringRestrictionsOutput, err error) {
	output = &EditHiringRestrictionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EditHiringRestrictionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Hiring_Restrictions_Request"`
	EditHiringRestrictionsRequestType
}

type EditHiringRestrictionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Hiring_Restrictions_Response"`
	EditHiringRestrictionsResponseType
}

// PutHireEventProposedWorkerID (Put_Hire_Event_Proposed_Worker_ID)
// 
// This web service imports a Worker ID which will override the Workday generated ID.
func (c *Client) PutHireEventProposedWorkerID(ctx context.Context, input *PutHireEventProposedWorkerIDInput) (output *PutHireEventProposedWorkerIDOutput, err error) {
	output = &PutHireEventProposedWorkerIDOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutHireEventProposedWorkerIDInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Hire_Event_Proposed_Worker_ID_Request"`
	PutHireEventProposedWorkerIDRequestType
}

type PutHireEventProposedWorkerIDOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Hire_Event_Proposed_Worker_ID_Response"`
	PutHireEventProposedWorkerIDResponseType
}

// EditWorkerAdditionalData (Edit_Worker_Additional_Data)
// 
// Allows updating of effective dated custom objects for a Worker.  When updating custom objects that allow multiple instances, data for existing instances must appear in the request, otherwise it will be removed.
func (c *Client) EditWorkerAdditionalData(ctx context.Context, input *EditWorkerAdditionalDataInput) (output *EditWorkerAdditionalDataOutput, err error) {
	output = &EditWorkerAdditionalDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EditWorkerAdditionalDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Worker_Additional_Data_Request"`
	EditWorkerAdditionalDataRequestType
}

type EditWorkerAdditionalDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Worker_Additional_Data_Response"`
	EditWorkerAdditionalDataResponseType
}

// EditPositionRestrictionsAdditionalData (Edit_Position_Restrictions_Additional_Data)
// 
// Allows updating of effective-dated custom objects for position restrictions.  When updating custom objects that allow multiple instances, data for existing instances must appear in the request, otherwise it will be removed.
func (c *Client) EditPositionRestrictionsAdditionalData(ctx context.Context, input *EditPositionRestrictionsAdditionalDataInput) (output *EditPositionRestrictionsAdditionalDataOutput, err error) {
	output = &EditPositionRestrictionsAdditionalDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EditPositionRestrictionsAdditionalDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Position_Restrictions_Additional_Data_Request"`
	EditPositionRestrictionsAdditionalDataRequestType
}

type EditPositionRestrictionsAdditionalDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Position_Restrictions_Additional_Data_Response"`
	EditPositionRestrictionsAdditionalDataResponseType
}

// MoveWorkersByOrganization (Move_Workers_By_Organization)
// 
// The following Organizations are valid for this Web Service: Company, Cost Center, Region, Custom Orgs that are Unique and are part of Change Organization Assignment, Supervisory Organizations (except Headcount Management). Job Management Supervisory Organizations must have Hiring Restrictions defined.
// 
// If the Organization is Supervisory, Position Management Positions, Job Management Positions, Position Restrictions can be moved (except Position Restrictions (and the filled positions) that are overlapped).
// 
// If the Organization is not Supervisory, Position Restrictions, Headcount Restrictions, and all filled positions of any staffing model can be moved.
func (c *Client) MoveWorkersByOrganization(ctx context.Context, input *MoveWorkersByOrganizationInput) (output *MoveWorkersByOrganizationOutput, err error) {
	output = &MoveWorkersByOrganizationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type MoveWorkersByOrganizationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Move_Workers_By_Organization_Request"`
	MoveWorkersByOrganizationRequestType
}

type MoveWorkersByOrganizationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Move_Workers_By_Organization_Response"`
	MoveWorkersByOrganizationResponseType
}

// ChangeOrganizationAssignments (Change_Organization_Assignments)
// 
// Assigns company, cost center, region, fund, grant, business unit, program, gift and custom organizations configured for staffing usage to a filled position or position restriction. Uses the Change Organization Assignments for Worker business process. This is to be used as a replacement for the web service Assign Organization.
func (c *Client) ChangeOrganizationAssignments(ctx context.Context, input *ChangeOrganizationAssignmentsInput) (output *ChangeOrganizationAssignmentsOutput, err error) {
	output = &ChangeOrganizationAssignmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeOrganizationAssignmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Organization_Assignments_Request"`
	ChangeOrganizationAssignmentsRequestType
}

type ChangeOrganizationAssignmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Organization_Assignments_Response"`
	ChangeOrganizationAssignmentsResponseType
}

// AssignRoles (Assign_Roles)
// 
// Assigns roles to assignees and uses the Assign Roles business process. You may specify a worker or position. If you specify a worker as the event target or an assignee to add and the worker has multiple positions as of the role assignment effective date, the international assignment (IA) position will take precedence over the worker&#39;s primary position. Specifying a worker with multiple positions as an assignee to remove will remove all the positions for that worker for that role assignment.
func (c *Client) AssignRoles(ctx context.Context, input *AssignRolesInput) (output *AssignRolesOutput, err error) {
	output = &AssignRolesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AssignRolesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Roles_Request"`
	AssignRolesRequestType
}

type AssignRolesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Roles_Response"`
	AssignRolesResponseType
}

// PutStudentEmploymentEligibilityStatus (Put_Student_Employment_Eligibility_Status)
// 
// Updates employment eligibility status and related reasons on the Student Employment Eligibility event, based on results from an external system. Not intended for EIB use.
func (c *Client) PutStudentEmploymentEligibilityStatus(ctx context.Context, input *PutStudentEmploymentEligibilityStatusInput) (output *PutStudentEmploymentEligibilityStatusOutput, err error) {
	output = &PutStudentEmploymentEligibilityStatusOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutStudentEmploymentEligibilityStatusInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Student_Employment_Eligibility_Status_Request"`
	PutStudentEmploymentEligibilityStatusRequestType
}

type PutStudentEmploymentEligibilityStatusOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Student_Employment_Eligibility_Status_Response"`
	PutStudentEmploymentEligibilityStatusResponseType
}

// GetStudentEmploymentEligibilityData (Get_Student_Employment_Eligibility_Data)
// 
// Gets data from Workday for the Verify Student Employment Eligibility business process, which is used to evaluate student employment eligibility in an external system. Not intended for EIB use.
func (c *Client) GetStudentEmploymentEligibilityData(ctx context.Context, input *GetStudentEmploymentEligibilityDataInput) (output *GetStudentEmploymentEligibilityDataOutput, err error) {
	output = &GetStudentEmploymentEligibilityDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetStudentEmploymentEligibilityDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Student_Employment_Eligibility_Data_Request"`
	GetStudentEmploymentEligibilityDataRequestType
}

type GetStudentEmploymentEligibilityDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Student_Employment_Eligibility_Data_Response"`
	GetStudentEmploymentEligibilityDataResponseType
}

// PutNoticePeriodEligibilityRule (Put_Notice_Period_Eligibility_Rule)
// 
// Adds or updates a notice period eligibility rule.
func (c *Client) PutNoticePeriodEligibilityRule(ctx context.Context, input *PutNoticePeriodEligibilityRuleInput) (output *PutNoticePeriodEligibilityRuleOutput, err error) {
	output = &PutNoticePeriodEligibilityRuleOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutNoticePeriodEligibilityRuleInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Notice_Period_Eligibility_Rule_Request"`
	PutNoticePeriodEligibilityRuleRequestType
}

type PutNoticePeriodEligibilityRuleOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Notice_Period_Eligibility_Rule_Response"`
	PutNoticePeriodEligibilityRuleResponseType
}

// PutMaintainNoticePeriodsforCountry (Put_Maintain_Notice_Periods_for_Country)
// 
// Maintain the Notice Periods table for a single country.
func (c *Client) PutMaintainNoticePeriodsforCountry(ctx context.Context, input *PutMaintainNoticePeriodsforCountryInput) (output *PutMaintainNoticePeriodsforCountryOutput, err error) {
	output = &PutMaintainNoticePeriodsforCountryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutMaintainNoticePeriodsforCountryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Maintain_Notice_Periods_for_Country_Request"`
	PutMaintainNoticePeriodsforCountryRequestType
}

type PutMaintainNoticePeriodsforCountryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Maintain_Notice_Periods_for_Country_Response"`
	PutMaintainNoticePeriodsforCountryResponseType
}

// GetNoticePeriodEligibilityRules (Get_Notice_Period_Eligibility_Rules)
// 
// Get Notice Period Eligibility Rules Web Service Task. This task returns all of the Notice Period Eligibility Rules specified in the request
func (c *Client) GetNoticePeriodEligibilityRules(ctx context.Context, input *GetNoticePeriodEligibilityRulesInput) (output *GetNoticePeriodEligibilityRulesOutput, err error) {
	output = &GetNoticePeriodEligibilityRulesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetNoticePeriodEligibilityRulesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Notice_Period_Eligibility_Rules_Request"`
	GetNoticePeriodEligibilityRulesRequestType
}

type GetNoticePeriodEligibilityRulesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Notice_Period_Eligibility_Rules_Response"`
	GetNoticePeriodEligibilityRulesResponseType
}

// GetMaintainNoticePeriodsForCountry (Get_Maintain_Notice_Periods_For_Country)
// 
// Get the Notice Period Rules for a single or many countries.
func (c *Client) GetMaintainNoticePeriodsForCountry(ctx context.Context, input *GetMaintainNoticePeriodsForCountryInput) (output *GetMaintainNoticePeriodsForCountryOutput, err error) {
	output = &GetMaintainNoticePeriodsForCountryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetMaintainNoticePeriodsForCountryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Maintain_Notice_Periods_For_Country_Request"`
	GetMaintainNoticePeriodsForCountryRequestType
}

type GetMaintainNoticePeriodsForCountryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Maintain_Notice_Periods_For_Country_Response"`
	GetMaintainNoticePeriodsForCountryResponseType
}

// PutEditNoticePeriodsEvent (Put_Edit_Notice_Periods_Event)
// 
// Adds or updates notice period details. Uses the Edit Notice Periods business process.
func (c *Client) PutEditNoticePeriodsEvent(ctx context.Context, input *PutEditNoticePeriodsEventInput) (output *PutEditNoticePeriodsEventOutput, err error) {
	output = &PutEditNoticePeriodsEventOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutEditNoticePeriodsEventInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Edit_Notice_Periods_Event_Request"`
	PutEditNoticePeriodsEventRequestType
}

type PutEditNoticePeriodsEventOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Edit_Notice_Periods_Event_Response"`
	PutEditNoticePeriodsEventResponseType
}

// GetNoticePeriods (Get_Notice_Periods)
// 
// Returns employer and employee notice periods for a notice period target
func (c *Client) GetNoticePeriods(ctx context.Context, input *GetNoticePeriodsInput) (output *GetNoticePeriodsOutput, err error) {
	output = &GetNoticePeriodsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetNoticePeriodsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Notice_Periods_Request"`
	GetNoticePeriodsRequestType
}

type GetNoticePeriodsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Notice_Periods_for_Notice_Period_Target_Response"`
	GetNoticePeriodsforNoticePeriodTargetResponseType
}

// GetMaintainProbationPeriodsForCountry (Get_Maintain_Probation_Periods_For_Country)
// 
// Get the Probation Period Rules for a single or many countries.
func (c *Client) GetMaintainProbationPeriodsForCountry(ctx context.Context, input *GetMaintainProbationPeriodsForCountryInput) (output *GetMaintainProbationPeriodsForCountryOutput, err error) {
	output = &GetMaintainProbationPeriodsForCountryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetMaintainProbationPeriodsForCountryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Maintain_Probation_Periods_For_Country_Request"`
	GetMaintainProbationPeriodsForCountryRequestType
}

type GetMaintainProbationPeriodsForCountryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Maintain_Probation_Periods_For_Country_Response"`
	GetMaintainProbationPeriodsForCountryResponseType
}

// PutMaintainProbationPeriodsForCountry (Put_Maintain_Probation_Periods_For_Country)
// 
// Maintain the Probation Periods table for a single country.
func (c *Client) PutMaintainProbationPeriodsForCountry(ctx context.Context, input *PutMaintainProbationPeriodsForCountryInput) (output *PutMaintainProbationPeriodsForCountryOutput, err error) {
	output = &PutMaintainProbationPeriodsForCountryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutMaintainProbationPeriodsForCountryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Maintain_Probation_Periods_For_Country_Request"`
	PutMaintainProbationPeriodsForCountryRequestType
}

type PutMaintainProbationPeriodsForCountryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Maintain_Probation_Periods_For_Country_Response"`
	PutMaintainProbationPeriodsForCountryResponseType
}

// PutStudentEmploymentEligibilityRule (Put_Student_Employment_Eligibility_Rule)
// 
// Creates a new or updates an existing Student Employment Eligibility Rule to be used in Student Employment Eligibility Rule Sets.
func (c *Client) PutStudentEmploymentEligibilityRule(ctx context.Context, input *PutStudentEmploymentEligibilityRuleInput) (output *PutStudentEmploymentEligibilityRuleOutput, err error) {
	output = &PutStudentEmploymentEligibilityRuleOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutStudentEmploymentEligibilityRuleInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Student_Employment_Eligibility_Rule_Request"`
	PutStudentEmploymentEligibilityRuleRequestType
}

type PutStudentEmploymentEligibilityRuleOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Student_Employment_Eligibility_Rule_Response"`
	PutStudentEmploymentEligibilityRuleResponseType
}

// GetStudentEmploymentEligibilityRules (Get_Student_Employment_Eligibility_Rules)
// 
// Get Student Employment Eligibility Rule Data
func (c *Client) GetStudentEmploymentEligibilityRules(ctx context.Context, input *GetStudentEmploymentEligibilityRulesInput) (output *GetStudentEmploymentEligibilityRulesOutput, err error) {
	output = &GetStudentEmploymentEligibilityRulesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetStudentEmploymentEligibilityRulesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Student_Employment_Eligibility_Rules_Request"`
	GetStudentEmploymentEligibilityRulesRequestType
}

type GetStudentEmploymentEligibilityRulesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Student_Employment_Eligibility_Rules_Response"`
	GetStudentEmploymentEligibilityRulesResponseType
}

// GetProbationPeriodEligibilityRules (Get_Probation_Period_Eligibility_Rules)
// 
// Get Probation Period Eligibility Rules Web Service Task. This task returns all of the Probation Period Eligibility Rules specified in the request
func (c *Client) GetProbationPeriodEligibilityRules(ctx context.Context, input *GetProbationPeriodEligibilityRulesInput) (output *GetProbationPeriodEligibilityRulesOutput, err error) {
	output = &GetProbationPeriodEligibilityRulesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetProbationPeriodEligibilityRulesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Probation_Period_Eligibility_Rules_Request"`
	GetProbationPeriodEligibilityRulesRequestType
}

type GetProbationPeriodEligibilityRulesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Probation_Period_Eligibility_Rules_Response"`
	GetProbationPeriodEligibilityRulesResponseType
}

// PutProbationPeriodEligibilityRule (Put_Probation_Period_Eligibility_Rule)
// 
// Adds, updates or delete a probation period eligibility rule.
func (c *Client) PutProbationPeriodEligibilityRule(ctx context.Context, input *PutProbationPeriodEligibilityRuleInput) (output *PutProbationPeriodEligibilityRuleOutput, err error) {
	output = &PutProbationPeriodEligibilityRuleOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutProbationPeriodEligibilityRuleInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Probation_Period_Eligibility_Rule_Request"`
	PutProbationPeriodEligibilityRuleRequestType
}

type PutProbationPeriodEligibilityRuleOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Probation_Period_Eligibility_Rule_Response"`
	PutProbationPeriodEligibilityRuleResponseType
}

// PutStudentEmploymentEligibilityRuleSet (Put_Student_Employment_Eligibility_Rule_Set)
// 
// Creates a new or updates an existing Student Employment Eligibility Rule Set to be used in Eligibility Verification
func (c *Client) PutStudentEmploymentEligibilityRuleSet(ctx context.Context, input *PutStudentEmploymentEligibilityRuleSetInput) (output *PutStudentEmploymentEligibilityRuleSetOutput, err error) {
	output = &PutStudentEmploymentEligibilityRuleSetOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutStudentEmploymentEligibilityRuleSetInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Student_Employment_Eligibility_Rule_Set_Request"`
	PutStudentEmploymentEligibilityRuleSetRequestType
}

type PutStudentEmploymentEligibilityRuleSetOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Student_Employment_Eligibility_Rule_Set_Response"`
	PutStudentEmploymentEligibilityRuleSetResponseType
}

// GetStudentEmploymentEligibilityRuleSet (Get_Student_Employment_Eligibility_Rule_Set)
// 
// Get Student Employment Eligibility Rule Set Data
func (c *Client) GetStudentEmploymentEligibilityRuleSet(ctx context.Context, input *GetStudentEmploymentEligibilityRuleSetInput) (output *GetStudentEmploymentEligibilityRuleSetOutput, err error) {
	output = &GetStudentEmploymentEligibilityRuleSetOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetStudentEmploymentEligibilityRuleSetInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Student_Employment_Eligibility_Rule_Set_Request"`
	GetStudentEmploymentEligibilityRuleSetRequestType
}

type GetStudentEmploymentEligibilityRuleSetOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Student_Employment_Eligibility_Rule_Set_Response"`
	GetStudentEmploymentEligibilityRuleSetResponseType
}

// ChangeWorkSpace (Change_Work_Space)
// 
// Assign Work Space Location to a Position. Uses the Change Work Space business process.
func (c *Client) ChangeWorkSpace(ctx context.Context, input *ChangeWorkSpaceInput) (output *ChangeWorkSpaceOutput, err error) {
	output = &ChangeWorkSpaceOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeWorkSpaceInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Work_Space_Request"`
	ChangeWorkSpaceRequestType
}

type ChangeWorkSpaceOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Work_Space_Response"`
	ChangeWorkSpaceResponseType
}

// FreezePosition (Freeze_Position)
// 
// Invokes the Business Process to Freeze a Position
func (c *Client) FreezePosition(ctx context.Context, input *FreezePositionInput) (output *FreezePositionOutput, err error) {
	output = &FreezePositionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type FreezePositionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Freeze_Position_Request"`
	FreezePositionRequestType
}

type FreezePositionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Freeze_Position_Response"`
	FreezePositionResponseType
}

// GetProbationPeriodsForWorkers (Get_Probation_Periods_For_Workers)
// 
// Return Employees Probation Periods
func (c *Client) GetProbationPeriodsForWorkers(ctx context.Context, input *GetProbationPeriodsForWorkersInput) (output *GetProbationPeriodsForWorkersOutput, err error) {
	output = &GetProbationPeriodsForWorkersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetProbationPeriodsForWorkersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Probation_Periods_For_Workers_Request"`
	GetProbationPeriodsForWorkersRequestType
}

type GetProbationPeriodsForWorkersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Probation_Periods_For_Workers_Response"`
	GetProbationPeriodsForWorkersResponseType
}

// NoShow (No_Show)
// 
// Performs a No Show on an employee
func (c *Client) NoShow(ctx context.Context, input *NoShowInput) (output *NoShowOutput, err error) {
	output = &NoShowOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type NoShowInput struct {
	XMLName string `xml:"urn:com.workday/bsvc No_Show_Request"`
	NoShowRequestType
}

type NoShowOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc No_Show_Event_Response"`
	NoShowEventResponseType
}

// PutProbationPeriodOutcomes (Put_Probation_Period_Outcomes)
// 
// Put Probation Period Outcomes for a single country.
func (c *Client) PutProbationPeriodOutcomes(ctx context.Context, input *PutProbationPeriodOutcomesInput) (output *PutProbationPeriodOutcomesOutput, err error) {
	output = &PutProbationPeriodOutcomesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutProbationPeriodOutcomesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Probation_Period_Outcomes_Request"`
	PutProbationPeriodOutcomesRequestType
}

type PutProbationPeriodOutcomesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Probation_Period_Outcomes_Response"`
	PutProbationPeriodOutcomesResponseType
}

// PutAssignMatrixOrganization (Put_Assign_Matrix_Organization)
// 
// Assign ~workers~ to matrix organizations.
func (c *Client) PutAssignMatrixOrganization(ctx context.Context, input *PutAssignMatrixOrganizationInput) (output *PutAssignMatrixOrganizationOutput, err error) {
	output = &PutAssignMatrixOrganizationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutAssignMatrixOrganizationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Assign_Matrix_Organization_Request"`
	PutAssignMatrixOrganizationRequestType
}

type PutAssignMatrixOrganizationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Assign_Matrix_Organization_Response"`
	PutAssignMatrixOrganizationResponseType
}

// GetProbationPeriodOutcomes (Get_Probation_Period_Outcomes)
// 
// Get the Probation Period Outcomes for one or more countries.
func (c *Client) GetProbationPeriodOutcomes(ctx context.Context, input *GetProbationPeriodOutcomesInput) (output *GetProbationPeriodOutcomesOutput, err error) {
	output = &GetProbationPeriodOutcomesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetProbationPeriodOutcomesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Probation_Period_Outcomes_Request"`
	GetProbationPeriodOutcomesRequestType
}

type GetProbationPeriodOutcomesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Probation_Period_Outcomes_Response"`
	GetProbationPeriodOutcomesResponseType
}

// EditEmployeeContractAdditionalData (Edit_Employee_Contract_Additional_Data)
// 
// Allows updating of effective dated custom objects for an ~Employee~ Contract. When updating custom objects that allow multiple instances, data for existing instances must appear in the request, otherwise it will be removed.
func (c *Client) EditEmployeeContractAdditionalData(ctx context.Context, input *EditEmployeeContractAdditionalDataInput) (output *EditEmployeeContractAdditionalDataOutput, err error) {
	output = &EditEmployeeContractAdditionalDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EditEmployeeContractAdditionalDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Employee_Contract_Additional_Data_Request"`
	EditEmployeeContractAdditionalDataRequestType
}

type EditEmployeeContractAdditionalDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Employee_Contract_Additional_Data_Response"`
	EditEmployeeContractAdditionalDataResponseType
}

// PutRemoveMatrixOrganization (Put_Remove_Matrix_Organization)
// 
// Remove ~workers~ from matrix organizations.
func (c *Client) PutRemoveMatrixOrganization(ctx context.Context, input *PutRemoveMatrixOrganizationInput) (output *PutRemoveMatrixOrganizationOutput, err error) {
	output = &PutRemoveMatrixOrganizationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutRemoveMatrixOrganizationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Remove_Matrix_Organization_Request"`
	PutRemoveMatrixOrganizationRequestType
}

type PutRemoveMatrixOrganizationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Remove_Matrix_Organization_Response"`
	PutRemoveMatrixOrganizationResponseType
}

// GetRoleAssignmentsForRoleAssigners (Get_Role_Assignments_For_Role_Assigners)
// 
// This web service returns role assignments (the combination of assignable role and role assignees) for the role assigner(s) specified in the request. You must specify at least 1 role assigner. Only explicitly assigned role assignees are returned. Default and inherited role assignees are not returned.
// 
// Optional request parameters include Assignable Roles, Effective Date and Time Zone. When no value is specified for &#39;Assignable Roles&#39;, role assignments for all roles are returned. When at least one value is specified for &#39;Assignable Roles&#39;, role assignments for only the specified roles are returned. When no effective date is specified, role assignments as of the current date are returned. When an effective date is specified, role assignments in effect as of the specified effective date are returned. You cannot specify a time zone in the request unless the &#39;Role Assignment Time Zone Option&#39; field in Tenant Setup - System has a value. When no time zone is specified, role assignments as of the effective date and Pacific Standard time (PST) are returned. When a time zone is specified, role assignments as of the effective date and the specified time zone are returned.
// 
// The count in the response filter applies to the number of role assigners returned.
// 
// For the supervisory organization assignable role linked to the Workday Role of &#39;Manager&#39;, only the single assignment manager role assignee is returned when this role has multiple role assignees.
func (c *Client) GetRoleAssignmentsForRoleAssigners(ctx context.Context, input *GetRoleAssignmentsForRoleAssignersInput) (output *GetRoleAssignmentsForRoleAssignersOutput, err error) {
	output = &GetRoleAssignmentsForRoleAssignersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetRoleAssignmentsForRoleAssignersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Role_Assignments_For_Role_Assigners_Request"`
	GetRoleAssignmentsForRoleAssignersRequestType
}

type GetRoleAssignmentsForRoleAssignersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Role_Assignments_For_Role_Assigners_Response"`
	GetRoleAssignmentsForRoleAssignersResponseType
}

// ImportExternalStudentInformation (Import_External_Student_Information)
// 
// High Volume Web Service task for creating or updating a large volume of External Student Records for new or existing Students.  This task can create a new Student with Personal Information only, and a reference to a Worker can be provided to create a new Student Role for an existing Worker.  If a Worker reference is provided, personal information (name, contact information, IDs, etc.) cannot be supplied.
// 
// This task is only intended for use in tenants that do not have Workday Student.  External Student Records should not be created in a tenant in which Workday Student has been implemented.
func (c *Client) ImportExternalStudentInformation(ctx context.Context, input *ImportExternalStudentInformationInput) (output *ImportExternalStudentInformationOutput, err error) {
	output = &ImportExternalStudentInformationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportExternalStudentInformationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_External_Student_Request"`
	ImportExternalStudentRequestType
}

type ImportExternalStudentInformationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportHireEmployee (Import_Hire_Employee)
// 
// This high-volume web service task hires employees.
func (c *Client) ImportHireEmployee(ctx context.Context, input *ImportHireEmployeeInput) (output *ImportHireEmployeeOutput, err error) {
	output = &ImportHireEmployeeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportHireEmployeeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Hire_Employee_Request"`
	ImportHireEmployeeRequestType
}

type ImportHireEmployeeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportContractContingentWorker (Import_Contract_Contingent_Worker)
// 
// This high-volume web service task contracts contingent workers.
func (c *Client) ImportContractContingentWorker(ctx context.Context, input *ImportContractContingentWorkerInput) (output *ImportContractContingentWorkerOutput, err error) {
	output = &ImportContractContingentWorkerOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportContractContingentWorkerInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Contract_Contingent_Worker_Request"`
	ImportContractContingentWorkerRequestType
}

type ImportContractContingentWorkerOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportApplicant (Import_Applicant)
// 
// This high-volume web service task creates Applicant data.
func (c *Client) ImportApplicant(ctx context.Context, input *ImportApplicantInput) (output *ImportApplicantOutput, err error) {
	output = &ImportApplicantOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportApplicantInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Applicant_Request"`
	ImportApplicantRequestType
}

type ImportApplicantOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportSwapPositions (Import_Swap_Positions)
// 
// Swap positions between ~employees~ or ~contingent workers~. Uses the Position Swap business process and creates the Change Job events on each ~worker~.
func (c *Client) ImportSwapPositions(ctx context.Context, input *ImportSwapPositionsInput) (output *ImportSwapPositionsOutput, err error) {
	output = &ImportSwapPositionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportSwapPositionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Swap_Positions_Request"`
	ImportSwapPositionsRequestType
}

type ImportSwapPositionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportTerminateEmployee (Import_Terminate_Employee)
// 
// This high-volume web service task terminates employees.
func (c *Client) ImportTerminateEmployee(ctx context.Context, input *ImportTerminateEmployeeInput) (output *ImportTerminateEmployeeOutput, err error) {
	output = &ImportTerminateEmployeeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportTerminateEmployeeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Terminate_Employee_Request"`
	ImportTerminateEmployeeRequestType
}

type ImportTerminateEmployeeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportEndContingentWorkerContract (Import_End_Contingent_Worker_Contract)
// 
// This high-volume web service task ends contingent worker contracts.
func (c *Client) ImportEndContingentWorkerContract(ctx context.Context, input *ImportEndContingentWorkerContractInput) (output *ImportEndContingentWorkerContractOutput, err error) {
	output = &ImportEndContingentWorkerContractOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportEndContingentWorkerContractInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_End_Contingent_Worker_Contract_Request"`
	ImportEndContingentWorkerContractRequestType
}

type ImportEndContingentWorkerContractOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportChangeJob (Import_Change_Job)
// 
// This high-volume web service task performs a job change on an employee or contingent worker using the Change Job business process. The types of changes include transfer, promotion, demotion, lateral moves and any other change of data on the job.
func (c *Client) ImportChangeJob(ctx context.Context, input *ImportChangeJobInput) (output *ImportChangeJobOutput, err error) {
	output = &ImportChangeJobOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportChangeJobInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Change_Job_Request"`
	ImportChangeJobRequestType
}

type ImportChangeJobOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportAssignMatrixOrganization (Import_Assign_Matrix_Organization)
// 
// Assign ~workers~ to matrix organizations.
func (c *Client) ImportAssignMatrixOrganization(ctx context.Context, input *ImportAssignMatrixOrganizationInput) (output *ImportAssignMatrixOrganizationOutput, err error) {
	output = &ImportAssignMatrixOrganizationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportAssignMatrixOrganizationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Assign_Matrix_Organization_Request"`
	ImportAssignMatrixOrganizationRequestType
}

type ImportAssignMatrixOrganizationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportRemoveMatrixOrganization (Import_Remove_Matrix_Organization)
// 
// Remove ~workers~ from matrix organizations.
func (c *Client) ImportRemoveMatrixOrganization(ctx context.Context, input *ImportRemoveMatrixOrganizationInput) (output *ImportRemoveMatrixOrganizationOutput, err error) {
	output = &ImportRemoveMatrixOrganizationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportRemoveMatrixOrganizationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Remove_Matrix_Organization_Request"`
	ImportRemoveMatrixOrganizationRequestType
}

type ImportRemoveMatrixOrganizationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

