
// Package recruiting
//
// The Recruiting Web Service contains operations that expose Workday Human Capital Management Business Services data for integration with Talent Management and Applicant Tracking systems.
package recruiting

import (
	"context"

	"github.com/ianlopshire/go-workday/pkg/wws"
)

const (
	CurrentVersion = "v31.2"
	ServiceName = "Recruiting"
)

type Client struct {
	*wws.Client
}


// GetServerTimestamp (Get_Server_Timestamp)
// 
// This operation retrieves Workday&#39;s current system datetime.
func (c *Client) GetServerTimestamp(ctx context.Context, input *GetServerTimestampInput) (output *GetServerTimestampOutput, err error) {
	output = &GetServerTimestampOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetServerTimestampInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Server_Timestamp_Get"`
	ServerTimestampGetType
}

type GetServerTimestampOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Server_Timestamp"`
	ServerTimestampType
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

// CreateJobRequisition (Create_Job_Requisition)
// 
// This operation will create a job requisition for a position in a position management supervisory organization.  The position must not have any open requisitions, and the tenant must have the Enable Job Requisitions flag set to true.  Requisitions cannot be reloaded into the system.
func (c *Client) CreateJobRequisition(ctx context.Context, input *CreateJobRequisitionInput) (output *CreateJobRequisitionOutput, err error) {
	output = &CreateJobRequisitionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CreateJobRequisitionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Create_Requisition_Request"`
	CreateRequisitionRequestType
}

type CreateJobRequisitionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Create_Job_Requisition_Response"`
	CreateJobRequisitionResponseType
}

// EditJobRequisition (Edit_Job_Requisition)
// 
// This operation will edit an existing job requisition.  The job requisition must be open.
func (c *Client) EditJobRequisition(ctx context.Context, input *EditJobRequisitionInput) (output *EditJobRequisitionOutput, err error) {
	output = &EditJobRequisitionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EditJobRequisitionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Job_Requisition_Request"`
	EditJobRequisitionRequestType
}

type EditJobRequisitionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Job_Requisition_Response"`
	EditJobRequisitionResponseType
}

// GetJobRequisitions (Get_Job_Requisitions)
// 
// Returns information for job requisitions and provides a reference to the related position.
func (c *Client) GetJobRequisitions(ctx context.Context, input *GetJobRequisitionsInput) (output *GetJobRequisitionsOutput, err error) {
	output = &GetJobRequisitionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetJobRequisitionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Requisitions_Request"`
	GetJobRequisitionsRequestType
}

type GetJobRequisitionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Requisitions_Response"`
	GetJobRequisitionsResponseType
}

// PutCandidate (Put_Candidate)
// 
// Adds or updates a candidate with the information supplied in the request.
func (c *Client) PutCandidate(ctx context.Context, input *PutCandidateInput) (output *PutCandidateOutput, err error) {
	output = &PutCandidateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCandidateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Candidate_Request"`
	PutCandidateRequestType
}

type PutCandidateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Candidate_Response"`
	PutCandidateResponseType
}

// GetCandidates (Get_Candidates)
// 
// Returns information for candidates specified in the request.
func (c *Client) GetCandidates(ctx context.Context, input *GetCandidatesInput) (output *GetCandidatesOutput, err error) {
	output = &GetCandidatesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCandidatesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Candidates_Request"`
	GetCandidatesRequestType
}

type GetCandidatesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Candidates_Response"`
	GetCandidatesResponseType
}

// UnpostJob (Unpost_Job)
// 
// This web service operation will unpost a job posting.  The operation does not support criteria based requests with an As_Of_Entry_DateTime in the past, unless for paging consistency the As_Of_Entry_DateTime matches the page 1 request.
func (c *Client) UnpostJob(ctx context.Context, input *UnpostJobInput) (output *UnpostJobOutput, err error) {
	output = &UnpostJobOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type UnpostJobInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Unpost_Job_Request"`
	UnpostJobRequestType
}

type UnpostJobOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Unpost_Job_Response"`
	UnpostJobResponseType
}

// PostJob (Post_Job)
// 
// This web service operation will post a job to a career site configured in a tenant.
func (c *Client) PostJob(ctx context.Context, input *PostJobInput) (output *PostJobOutput, err error) {
	output = &PostJobOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PostJobInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Post_Job_Request"`
	PostJobRequestType
}

type PostJobOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Post_Job_Response"`
	PostJobResponseType
}

// UpdateJobPosting (Update_Job_Posting)
// 
// Updates Job Posting start date and end date.
func (c *Client) UpdateJobPosting(ctx context.Context, input *UpdateJobPostingInput) (output *UpdateJobPostingOutput, err error) {
	output = &UpdateJobPostingOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type UpdateJobPostingInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Update_Job_Posting_Request"`
	UpdateJobPostingRequestType
}

type UpdateJobPostingOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Update_Job_Posting_Response"`
	UpdateJobPostingResponseType
}

// CloseJobRequisition (Close_Job_Requisition)
// 
// This operation will close a job requisition. The requisition must not have any pending events or be filled.
func (c *Client) CloseJobRequisition(ctx context.Context, input *CloseJobRequisitionInput) (output *CloseJobRequisitionOutput, err error) {
	output = &CloseJobRequisitionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CloseJobRequisitionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Close_Job_Requisition_Request"`
	CloseJobRequisitionRequestType
}

type CloseJobRequisitionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Close_Job_Requisition_Response"`
	CloseJobRequisitionResponseType
}

// PutCandidateAttachment (Put_Candidate_Attachment)
// 
// This operation adds an attachment to a Candidate and/or a Candidate&#39;s Job Application.
func (c *Client) PutCandidateAttachment(ctx context.Context, input *PutCandidateAttachmentInput) (output *PutCandidateAttachmentOutput, err error) {
	output = &PutCandidateAttachmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCandidateAttachmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Candidate_Attachment_Request"`
	PutCandidateAttachmentRequestType
}

type PutCandidateAttachmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Candidate_Attachment_Response"`
	PutCandidateAttachmentResponseType
}

// GetCandidateAttachments (Get_Candidate_Attachments)
// 
// This operation returns attachments for a Candidate and/or a Candidate&#39;s Job Application.
func (c *Client) GetCandidateAttachments(ctx context.Context, input *GetCandidateAttachmentsInput) (output *GetCandidateAttachmentsOutput, err error) {
	output = &GetCandidateAttachmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCandidateAttachmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Candidate_Attachments_Request"`
	GetCandidateAttachmentsRequestType
}

type GetCandidateAttachmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Candidate_Attachments_Response"`
	GetCandidateAttachmentsResponseType
}

// GetJobPostingSites (Get_Job_Posting_Sites)
// 
// Provides an Integration Partner the ability to GET a posting site details in the Posting Site Setup table through a webservice.
func (c *Client) GetJobPostingSites(ctx context.Context, input *GetJobPostingSitesInput) (output *GetJobPostingSitesOutput, err error) {
	output = &GetJobPostingSitesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetJobPostingSitesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Posting_Sites_Request"`
	GetJobPostingSitesRequestType
}

type GetJobPostingSitesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Posting_Sites_Response"`
	GetJobPostingSitesResponseType
}

// PutJobPostingSite (Put_Job_Posting_Site)
// 
// Adds a generic job posting board to the tenant.
func (c *Client) PutJobPostingSite(ctx context.Context, input *PutJobPostingSiteInput) (output *PutJobPostingSiteOutput, err error) {
	output = &PutJobPostingSiteOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutJobPostingSiteInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Posting_Site_Request"`
	PutJobPostingSiteRequestType
}

type PutJobPostingSiteOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Posting_Site_Response"`
	PutJobPostingSiteResponseType
}

// PutCandidatePhoto (Put_Candidate_Photo)
// 
// This operation adds an image to a Candidate.
func (c *Client) PutCandidatePhoto(ctx context.Context, input *PutCandidatePhotoInput) (output *PutCandidatePhotoOutput, err error) {
	output = &PutCandidatePhotoOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCandidatePhotoInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Candidate_Photo_Request"`
	PutCandidatePhotoRequestType
}

type PutCandidatePhotoOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Candidate_Photos_Response"`
	PutCandidatePhotosResponseType
}

// GetCandidatePhotos (Get_Candidate_Photos)
// 
// This operation returns the image for a Candidate.
func (c *Client) GetCandidatePhotos(ctx context.Context, input *GetCandidatePhotosInput) (output *GetCandidatePhotosOutput, err error) {
	output = &GetCandidatePhotosOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCandidatePhotosInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Candidate_Photos_Request"`
	GetCandidatePhotosRequestType
}

type GetCandidatePhotosOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Candidate_Photos_Response"`
	GetCandidatePhotosResponseType
}

// PutBackgroundCheckPackage (Put_Background_Check_Package)
// 
// Web service task to create or update background check packages.
func (c *Client) PutBackgroundCheckPackage(ctx context.Context, input *PutBackgroundCheckPackageInput) (output *PutBackgroundCheckPackageOutput, err error) {
	output = &PutBackgroundCheckPackageOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBackgroundCheckPackageInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Background_Check_Package_Request"`
	PutBackgroundCheckPackageRequestType
}

type PutBackgroundCheckPackageOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Background_Check_Package_Response"`
	PutBackgroundCheckPackageResponseType
}

// GetBackgroundCheckPackages (Get_Background_Check_Packages)
// 
// Returns information for Background Check Packages specified in the request.
func (c *Client) GetBackgroundCheckPackages(ctx context.Context, input *GetBackgroundCheckPackagesInput) (output *GetBackgroundCheckPackagesOutput, err error) {
	output = &GetBackgroundCheckPackagesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBackgroundCheckPackagesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Background_Check_Packages_Request"`
	GetBackgroundCheckPackagesRequestType
}

type GetBackgroundCheckPackagesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Background_Check_Packages_Response"`
	GetBackgroundCheckPackagesResponseType
}

// ManageJobRequisitionFreeze (Manage_Job_Requisition_Freeze)
// 
// This operation will freeze or unfreeze a job requisition. The requisition must not be filled or closed.
func (c *Client) ManageJobRequisitionFreeze(ctx context.Context, input *ManageJobRequisitionFreezeInput) (output *ManageJobRequisitionFreezeOutput, err error) {
	output = &ManageJobRequisitionFreezeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ManageJobRequisitionFreezeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Job_Requisition_Freeze_Request"`
	ManageJobRequisitionFreezeRequestType
}

type ManageJobRequisitionFreezeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Job_Requisition_Freeze_Response"`
	ManageJobRequisitionFreezeResponseType
}

// PutBackgroundCheck (Put_Background_Check)
// 
// Load background check results for a Background Check Event.
func (c *Client) PutBackgroundCheck(ctx context.Context, input *PutBackgroundCheckInput) (output *PutBackgroundCheckOutput, err error) {
	output = &PutBackgroundCheckOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBackgroundCheckInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Background_Check_Request"`
	PutBackgroundCheckRequestType
}

type PutBackgroundCheckOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Background_Check_Response"`
	PutBackgroundCheckResponseType
}

// GetBackgroundCheck (Get_Background_Check)
// 
// Retrieve background check results for a Background Check Event.
func (c *Client) GetBackgroundCheck(ctx context.Context, input *GetBackgroundCheckInput) (output *GetBackgroundCheckOutput, err error) {
	output = &GetBackgroundCheckOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBackgroundCheckInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Background_Check_Request"`
	GetBackgroundCheckRequestType
}

type GetBackgroundCheckOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Background_Check_Response"`
	GetBackgroundCheckResponseType
}

// GetJobPostings (Get_Job_Postings)
// 
// Returns information for Job Postings specified in the requests.
func (c *Client) GetJobPostings(ctx context.Context, input *GetJobPostingsInput) (output *GetJobPostingsOutput, err error) {
	output = &GetJobPostingsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetJobPostingsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Postings_Request"`
	GetJobPostingsRequestType
}

type GetJobPostingsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Postings_Response"`
	GetJobPostingsResponseType
}

// GetAssessCandidate (Get_Assess_Candidate)
// 
// Retrieve information about Candidate Assessments.
func (c *Client) GetAssessCandidate(ctx context.Context, input *GetAssessCandidateInput) (output *GetAssessCandidateOutput, err error) {
	output = &GetAssessCandidateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAssessCandidateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Assess_Candidate_Request"`
	GetAssessCandidateRequestType
}

type GetAssessCandidateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Assess_Candidate_Response"`
	GetAssessCandidateResponseType
}

// AssessCandidate (Assess_Candidate)
// 
// Create or modify a Candidate Assessment.  The Job Application business process relevant to the candidate must be in the Assess Candidate state.
func (c *Client) AssessCandidate(ctx context.Context, input *AssessCandidateInput) (output *AssessCandidateOutput, err error) {
	output = &AssessCandidateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AssessCandidateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assess_Candidate_Request"`
	AssessCandidateRequestType
}

type AssessCandidateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assess_Candidate_Response"`
	AssessCandidateResponseType
}

// CreateEvergreenRequisition (Create_Evergreen_Requisition)
// 
// This operation will create an evergreen job requisition.
func (c *Client) CreateEvergreenRequisition(ctx context.Context, input *CreateEvergreenRequisitionInput) (output *CreateEvergreenRequisitionOutput, err error) {
	output = &CreateEvergreenRequisitionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CreateEvergreenRequisitionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Create_Evergreen_Requisition_Request"`
	CreateEvergreenRequisitionRequestType
}

type CreateEvergreenRequisitionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Create_Evergreen_Requisition_Response"`
	CreateEvergreenRequisitionResponseType
}

// GetEvergreenRequisitions (Get_Evergreen_Requisitions)
// 
// Returns information for evergreen requisitions and provides a reference to the related position.
func (c *Client) GetEvergreenRequisitions(ctx context.Context, input *GetEvergreenRequisitionsInput) (output *GetEvergreenRequisitionsOutput, err error) {
	output = &GetEvergreenRequisitionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEvergreenRequisitionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Evergreen_Requisitions_Request"`
	GetEvergreenRequisitionsRequestType
}

type GetEvergreenRequisitionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Evergreen_Requisitions_Response"`
	GetEvergreenRequisitionsResponseType
}

// EditEvergreenRequisition (Edit_Evergreen_Requisition)
// 
// This operation will edit an evergreen job requisition.
func (c *Client) EditEvergreenRequisition(ctx context.Context, input *EditEvergreenRequisitionInput) (output *EditEvergreenRequisitionOutput, err error) {
	output = &EditEvergreenRequisitionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EditEvergreenRequisitionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Evergreen_Requisition_Request"`
	EditEvergreenRequisitionRequestType
}

type EditEvergreenRequisitionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Evergreen_Requisition_Response"`
	EditEvergreenRequisitionResponseType
}

// CloseEvergreenRequisition (Close_Evergreen_Requisition)
// 
// This operation will close an evergreen requisition. The requisition must not have any pending events.
func (c *Client) CloseEvergreenRequisition(ctx context.Context, input *CloseEvergreenRequisitionInput) (output *CloseEvergreenRequisitionOutput, err error) {
	output = &CloseEvergreenRequisitionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CloseEvergreenRequisitionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Close_Evergreen_Requisition_Request"`
	CloseEvergreenRequisitionRequestType
}

type CloseEvergreenRequisitionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Close_Evergreen_Requisition_Response"`
	CloseEvergreenRequisitionResponseType
}

// MoveJobRequisition (Move_Job_Requisition)
// 
// Move a job requisition from one Job Management Organization to another Job Management Organization using the Move Job Requisition business process.
func (c *Client) MoveJobRequisition(ctx context.Context, input *MoveJobRequisitionInput) (output *MoveJobRequisitionOutput, err error) {
	output = &MoveJobRequisitionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type MoveJobRequisitionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Move_Job_Requisition_Request"`
	MoveJobRequisitionRequestType
}

type MoveJobRequisitionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Move_Job_Requisition_Response"`
	MoveJobRequisitionResponseType
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

// EditJobRequisitionAdditionalData (Edit_Job_Requisition_Additional_Data)
// 
// Allows updating of effective-dated custom objects for job requisitions.  When updating custom objects that allow multiple instances, data for existing instances must appear in the request, otherwise it will be removed.
func (c *Client) EditJobRequisitionAdditionalData(ctx context.Context, input *EditJobRequisitionAdditionalDataInput) (output *EditJobRequisitionAdditionalDataOutput, err error) {
	output = &EditJobRequisitionAdditionalDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EditJobRequisitionAdditionalDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Job_Requisition_Additional_Data_Request"`
	EditJobRequisitionAdditionalDataRequestType
}

type EditJobRequisitionAdditionalDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Job_Requisition_Additional_Data_Response"`
	EditJobRequisitionAdditionalDataResponseType
}

// GetRecruitingAgencyUsers (Get_Recruiting_Agency_Users)
// 
// Get Recruiting Agency Users.
func (c *Client) GetRecruitingAgencyUsers(ctx context.Context, input *GetRecruitingAgencyUsersInput) (output *GetRecruitingAgencyUsersOutput, err error) {
	output = &GetRecruitingAgencyUsersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetRecruitingAgencyUsersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Recruiting_Agency_Users_Request"`
	GetRecruitingAgencyUsersRequestType
}

type GetRecruitingAgencyUsersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Recruiting_Agency_Users_Response"`
	GetRecruitingAgencyUsersResponseType
}

// PutRecruitingAgencyUser (Put_Recruiting_Agency_User)
// 
// Adds or updates a recruiting agency user with the information supplied in the request.
func (c *Client) PutRecruitingAgencyUser(ctx context.Context, input *PutRecruitingAgencyUserInput) (output *PutRecruitingAgencyUserOutput, err error) {
	output = &PutRecruitingAgencyUserOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutRecruitingAgencyUserInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Recruiting_Agency_User_Request"`
	PutRecruitingAgencyUserRequestType
}

type PutRecruitingAgencyUserOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Recruiting_Agency_User_Response"`
	PutRecruitingAgencyUserResponseType
}

// MoveCandidatetoLinkedJobRequisition (Move_Candidate_to_Linked_Job_Requisition)
// 
// This operation will move a Job Application from an Evergreen Requisition to a Job Requisition.
func (c *Client) MoveCandidatetoLinkedJobRequisition(ctx context.Context, input *MoveCandidatetoLinkedJobRequisitionInput) (output *MoveCandidatetoLinkedJobRequisitionOutput, err error) {
	output = &MoveCandidatetoLinkedJobRequisitionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type MoveCandidatetoLinkedJobRequisitionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Move_Candidate_to_Linked_Job_Requisition_Request"`
	MoveCandidatetoLinkedJobRequisitionRequestType
}

type MoveCandidatetoLinkedJobRequisitionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Move_Candidate_to_Linked_Job_Requisition_Response"`
	MoveCandidatetoLinkedJobRequisitionResponseType
}

// PutPrimaryPosting (Put_Primary_Posting)
// 
// Updates an existing Job Posting to make it the Primary Posting, used in social share and invite to apply.
func (c *Client) PutPrimaryPosting(ctx context.Context, input *PutPrimaryPostingInput) (output *PutPrimaryPostingOutput, err error) {
	output = &PutPrimaryPostingOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPrimaryPostingInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Primary_Posting_Request"`
	PutPrimaryPostingRequestType
}

type PutPrimaryPostingOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Primary_Posting_Response"`
	PutPrimaryPostingResponseType
}

// MoveCandidatetoLinkedEvergreenRequisition (Move_Candidate_to_Linked_Evergreen_Requisition)
// 
// Move candidates from a Job Requisition to a linked ~Evergreen Requisition~.
func (c *Client) MoveCandidatetoLinkedEvergreenRequisition(ctx context.Context, input *MoveCandidatetoLinkedEvergreenRequisitionInput) (output *MoveCandidatetoLinkedEvergreenRequisitionOutput, err error) {
	output = &MoveCandidatetoLinkedEvergreenRequisitionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type MoveCandidatetoLinkedEvergreenRequisitionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Move_Candidate_to_Linked_Evergreen_Requisition_Request"`
	MoveCandidatetoLinkedEvergreenRequisitionRequestType
}

type MoveCandidatetoLinkedEvergreenRequisitionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Move_Candidate_to_Linked_Evergreen_Requisition_Response"`
	MoveCandidatetoLinkedEvergreenRequisitionResponseType
}

// GetJobApplicationAdditionalData (Get_Job_Application_Additional_Data)
// 
// Get Additional data for Job Application
func (c *Client) GetJobApplicationAdditionalData(ctx context.Context, input *GetJobApplicationAdditionalDataInput) (output *GetJobApplicationAdditionalDataOutput, err error) {
	output = &GetJobApplicationAdditionalDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetJobApplicationAdditionalDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Application_Additional_Data_Request"`
	GetJobApplicationAdditionalDataRequestType
}

type GetJobApplicationAdditionalDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Application_Additional_Data_Response"`
	GetJobApplicationAdditionalDataResponseType
}

// PutJobApplicationAdditionalData (Put_Job_Application_Additional_Data)
// 
// Adds or updates Job Application with the Additional Data information supplied in the request.
func (c *Client) PutJobApplicationAdditionalData(ctx context.Context, input *PutJobApplicationAdditionalDataInput) (output *PutJobApplicationAdditionalDataOutput, err error) {
	output = &PutJobApplicationAdditionalDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutJobApplicationAdditionalDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Application_Additional_Data_Request"`
	PutJobApplicationAdditionalDataRequestType
}

type PutJobApplicationAdditionalDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Application_Additional_Data_Response"`
	PutJobApplicationAdditionalDataResponseType
}

// GetRecruitingAgencies (Get_Recruiting_Agencies)
// 
// Get Recruiting Agencies
func (c *Client) GetRecruitingAgencies(ctx context.Context, input *GetRecruitingAgenciesInput) (output *GetRecruitingAgenciesOutput, err error) {
	output = &GetRecruitingAgenciesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetRecruitingAgenciesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Recruiting_Agencies_Request"`
	GetRecruitingAgenciesRequestType
}

type GetRecruitingAgenciesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Recruiting_Agencies_Response"`
	GetRecruitingAgenciesResponseType
}

// PutRecruitingAgency (Put_Recruiting_Agency)
// 
// Add or update a Recruiting Agency specified with the information supplied in the request.
func (c *Client) PutRecruitingAgency(ctx context.Context, input *PutRecruitingAgencyInput) (output *PutRecruitingAgencyOutput, err error) {
	output = &PutRecruitingAgencyOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutRecruitingAgencyInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Recruiting_Agency_Request"`
	PutRecruitingAgencyRequestType
}

type PutRecruitingAgencyOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Recruiting_Agency_Response"`
	PutRecruitingAgencyResponseType
}

// SubmitRecruitingAgencyCandidate (Submit_Recruiting_Agency_Candidate)
// 
// Submit a candidate as recruiting agency via recruiting agency candidate process with the information provided in the request.
func (c *Client) SubmitRecruitingAgencyCandidate(ctx context.Context, input *SubmitRecruitingAgencyCandidateInput) (output *SubmitRecruitingAgencyCandidateOutput, err error) {
	output = &SubmitRecruitingAgencyCandidateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitRecruitingAgencyCandidateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Recruiting_Agency_Candidate_Request"`
	SubmitRecruitingAgencyCandidateRequestType
}

type SubmitRecruitingAgencyCandidateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Recruiting_Agency_Candidate__Response"`
	SubmitRecruitingAgencyCandidateResponseType
}

// MoveCandidate (Move_Candidate)
// 
// Move Candidate from any Recruiting Stage to next possible stage or to a disposition stage.
func (c *Client) MoveCandidate(ctx context.Context, input *MoveCandidateInput) (output *MoveCandidateOutput, err error) {
	output = &MoveCandidateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type MoveCandidateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Move_Candidate_Request"`
	MoveCandidateRequestType
}

type MoveCandidateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Move_Candidate_Response"`
	MoveCandidateResponseType
}

// GetRecruitingAgencyAdditionalData (Get_Recruiting_Agency_Additional_Data)
// 
// Retrieves additional data associated with the supplied Recruiting Agency.
func (c *Client) GetRecruitingAgencyAdditionalData(ctx context.Context, input *GetRecruitingAgencyAdditionalDataInput) (output *GetRecruitingAgencyAdditionalDataOutput, err error) {
	output = &GetRecruitingAgencyAdditionalDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetRecruitingAgencyAdditionalDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Recruiting_Agency_Additional_Data_Request"`
	GetRecruitingAgencyAdditionalDataRequestType
}

type GetRecruitingAgencyAdditionalDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Recruiting_Agency_Additional_Data_Response"`
	GetRecruitingAgencyAdditionalDataResponseType
}

// PutRecruitingAgencyAdditionalData (Put_Recruiting_Agency_Additional_Data)
// 
// Add or update additional data to the Recruiting Agency specified. Additional data can only be added for custom objects already associated with Recruiting Agencies.
func (c *Client) PutRecruitingAgencyAdditionalData(ctx context.Context, input *PutRecruitingAgencyAdditionalDataInput) (output *PutRecruitingAgencyAdditionalDataOutput, err error) {
	output = &PutRecruitingAgencyAdditionalDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutRecruitingAgencyAdditionalDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Recruiting_Agency_Additional_Data_Request"`
	PutRecruitingAgencyAdditionalDataRequestType
}

type PutRecruitingAgencyAdditionalDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Recruiting_Agency_Additional_Data_Response"`
	PutRecruitingAgencyAdditionalDataResponseType
}

// GetJobRequisitionInterviewTeams (Get_Job_Requisition_Interview_Teams)
// 
// Web services for Manage Interview Team task.
func (c *Client) GetJobRequisitionInterviewTeams(ctx context.Context, input *GetJobRequisitionInterviewTeamsInput) (output *GetJobRequisitionInterviewTeamsOutput, err error) {
	output = &GetJobRequisitionInterviewTeamsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetJobRequisitionInterviewTeamsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Requisition_Interview_Teams_Request"`
	GetJobRequisitionInterviewTeamsRequestType
}

type GetJobRequisitionInterviewTeamsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Requisition_Interview_Teams_Response"`
	GetJobRequisitionInterviewTeamsResponseType
}

// PutJobRequisitionInterviewTeam (Put_Job_Requisition_Interview_Team)
// 
// Put web services for Manage Interview Team task.
func (c *Client) PutJobRequisitionInterviewTeam(ctx context.Context, input *PutJobRequisitionInterviewTeamInput) (output *PutJobRequisitionInterviewTeamOutput, err error) {
	output = &PutJobRequisitionInterviewTeamOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutJobRequisitionInterviewTeamInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Requisition_Interview_Team_Request"`
	PutJobRequisitionInterviewTeamRequestType
}

type PutJobRequisitionInterviewTeamOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Requisition_Interview_Team_Response"`
	PutJobRequisitionInterviewTeamResponseType
}

// PutInterview (Put_Interview)
// 
// Submits Interview Data for the Schedule Interview task.
func (c *Client) PutInterview(ctx context.Context, input *PutInterviewInput) (output *PutInterviewOutput, err error) {
	output = &PutInterviewOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutInterviewInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Interview_Request"`
	PutInterviewRequestType
}

type PutInterviewOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Interview_Response"`
	PutInterviewResponseType
}

// GetInterviews (Get_Interviews)
// 
// Retrieves Interview Data for the Schedule Interview task.
func (c *Client) GetInterviews(ctx context.Context, input *GetInterviewsInput) (output *GetInterviewsOutput, err error) {
	output = &GetInterviewsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetInterviewsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Interviews_Request"`
	GetInterviewsRequestType
}

type GetInterviewsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Interviews_Response"`
	GetInterviewsResponseType
}

// ReferaCandidate (Refer_a_Candidate)
// 
// Refer a Candidate with the information provided in the request.
func (c *Client) ReferaCandidate(ctx context.Context, input *ReferaCandidateInput) (output *ReferaCandidateOutput, err error) {
	output = &ReferaCandidateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ReferaCandidateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Refer_a_Candidate_Request"`
	ReferaCandidateRequestType
}

type ReferaCandidateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Refer_a_Candidate_Response"`
	ReferaCandidateResponseType
}

// PutInterviewFeedback (Put_Interview_Feedback)
// 
// Submits Interview Feedback for Interviewers on the Interview Schedule. Also, allows you to move a candidate to the next stage of the business process.
func (c *Client) PutInterviewFeedback(ctx context.Context, input *PutInterviewFeedbackInput) (output *PutInterviewFeedbackOutput, err error) {
	output = &PutInterviewFeedbackOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutInterviewFeedbackInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Interview_Feedback_Request"`
	PutInterviewFeedbackRequestType
}

type PutInterviewFeedbackOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Interview_Feedback_Response"`
	PutInterviewFeedbackResponseType
}

// GetInterviewFeedbacks (Get_Interview_Feedbacks)
// 
// Retrieves Interview Feedback Data for the Manage Interview Feedback task.
func (c *Client) GetInterviewFeedbacks(ctx context.Context, input *GetInterviewFeedbacksInput) (output *GetInterviewFeedbacksOutput, err error) {
	output = &GetInterviewFeedbacksOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetInterviewFeedbacksInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Interview_Feedbacks_Request"`
	GetInterviewFeedbacksRequestType
}

type GetInterviewFeedbacksOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Interview_Feedbacks_Response"`
	GetInterviewFeedbacksResponseType
}

// PutVeteranStatus (Put_Veteran_Status)
// 
// Creates a new Veteran Status with the information supplied in the request.
func (c *Client) PutVeteranStatus(ctx context.Context, input *PutVeteranStatusInput) (output *PutVeteranStatusOutput, err error) {
	output = &PutVeteranStatusOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutVeteranStatusInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Veteran_Status_Request"`
	PutVeteranStatusRequestType
}

type PutVeteranStatusOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Veteran_Status_Response"`
	PutVeteranStatusResponseType
}

// GetVeteranStatuses (Get_Veteran_Statuses)
// 
// Returns Veteran Statuses by Reference ID, or returns all Veteran Statuses if no Reference ID is provided.
func (c *Client) GetVeteranStatuses(ctx context.Context, input *GetVeteranStatusesInput) (output *GetVeteranStatusesOutput, err error) {
	output = &GetVeteranStatusesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetVeteranStatusesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Veteran_Statuses_Request"`
	GetVeteranStatusesRequestType
}

type GetVeteranStatusesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Veteran_Statuses_Response"`
	GetVeteranStatusesResponseType
}

// PutUndoMovefromHire (Put_Undo_Move_from_Hire)
// 
// Undo the move from a staffing process back to the originating job application process.
func (c *Client) PutUndoMovefromHire(ctx context.Context, input *PutUndoMovefromHireInput) (output *PutUndoMovefromHireOutput, err error) {
	output = &PutUndoMovefromHireOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutUndoMovefromHireInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Undo_Move_from_Hire_Request"`
	PutUndoMovefromHireRequestType
}

type PutUndoMovefromHireOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Undo_Move_from_Hire_Response"`
	PutUndoMovefromHireResponseType
}

// GetRecruitingSelfScheduleCalendarTypes (Get_Recruiting_Self-Schedule_Calendar_Types)
// 
// Web service operation to retrieve Recruiting Self-Schedule Calendar Types.
func (c *Client) GetRecruitingSelfScheduleCalendarTypes(ctx context.Context, input *GetRecruitingSelfScheduleCalendarTypesInput) (output *GetRecruitingSelfScheduleCalendarTypesOutput, err error) {
	output = &GetRecruitingSelfScheduleCalendarTypesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetRecruitingSelfScheduleCalendarTypesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Recruiting_Self-Schedule_Calendar_Types_Request"`
	GetRecruitingSelfScheduleCalendarTypesRequestType
}

type GetRecruitingSelfScheduleCalendarTypesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Recruiting_Self-Schedule_Calendar_Types_Response"`
	GetRecruitingSelfScheduleCalendarTypesResponseType
}

// PutRecruitingSelfScheduleCalendarType (Put_Recruiting_Self-Schedule_Calendar_Type)
// 
// Web service operation to add, update or delete a Recruiting Self-Schedule Calendar Type.
func (c *Client) PutRecruitingSelfScheduleCalendarType(ctx context.Context, input *PutRecruitingSelfScheduleCalendarTypeInput) (output *PutRecruitingSelfScheduleCalendarTypeOutput, err error) {
	output = &PutRecruitingSelfScheduleCalendarTypeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutRecruitingSelfScheduleCalendarTypeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Recruiting_Self-Schedule_Calendar_Type_Request"`
	PutRecruitingSelfScheduleCalendarTypeRequestType
}

type PutRecruitingSelfScheduleCalendarTypeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Recruiting_Self-Schedule_Calendar_Type_Response"`
	PutRecruitingSelfScheduleCalendarTypeResponseType
}

// GetAssignRecruitingSelfScheduleCalendars (Get_Assign_Recruiting_Self-Schedule_Calendars)
// 
// Web service operation to retrieve Recruiting Self-Schedule Calendar Assignments.
func (c *Client) GetAssignRecruitingSelfScheduleCalendars(ctx context.Context, input *GetAssignRecruitingSelfScheduleCalendarsInput) (output *GetAssignRecruitingSelfScheduleCalendarsOutput, err error) {
	output = &GetAssignRecruitingSelfScheduleCalendarsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAssignRecruitingSelfScheduleCalendarsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Assign_Recruiting_Self-Schedule_Calendars_Request"`
	GetAssignRecruitingSelfScheduleCalendarsRequestType
}

type GetAssignRecruitingSelfScheduleCalendarsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Assign_Recruiting_Self-Schedule_Calendars_Response"`
	GetAssignRecruitingSelfScheduleCalendarsResponseType
}

// AssignRecruitingSelfScheduleCalendar (Assign_Recruiting_Self-Schedule_Calendar)
// 
// Web service for assigning recruiting self-schedule calendars to job requisition enabled.
func (c *Client) AssignRecruitingSelfScheduleCalendar(ctx context.Context, input *AssignRecruitingSelfScheduleCalendarInput) (output *AssignRecruitingSelfScheduleCalendarOutput, err error) {
	output = &AssignRecruitingSelfScheduleCalendarOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AssignRecruitingSelfScheduleCalendarInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Recruiting_Self-Schedule_Calendar_Request"`
	AssignRecruitingSelfScheduleCalendarRequestType
}

type AssignRecruitingSelfScheduleCalendarOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Recruiting_Self-Schedule_Calendar_Response"`
	AssignRecruitingSelfScheduleCalendarResponseType
}

// GetRecruitingSelfScheduleCalendars (Get_Recruiting_Self-Schedule_Calendars)
// 
// Web service operation to retrieve Recruiting Self-Schedule Calendars.
func (c *Client) GetRecruitingSelfScheduleCalendars(ctx context.Context, input *GetRecruitingSelfScheduleCalendarsInput) (output *GetRecruitingSelfScheduleCalendarsOutput, err error) {
	output = &GetRecruitingSelfScheduleCalendarsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetRecruitingSelfScheduleCalendarsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Recruiting_Self-Schedule_Calendars_Request"`
	GetRecruitingSelfScheduleCalendarsRequestType
}

type GetRecruitingSelfScheduleCalendarsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Recruiting_Self-Schedule_Calendars_Response"`
	GetRecruitingSelfScheduleCalendarsResponseType
}

// PutRecruitingSelfScheduleCalendar (Put_Recruiting_Self-Schedule_Calendar)
// 
// Web service operation to add, update or delete a Recruiting Self-Schedule Calendar.
func (c *Client) PutRecruitingSelfScheduleCalendar(ctx context.Context, input *PutRecruitingSelfScheduleCalendarInput) (output *PutRecruitingSelfScheduleCalendarOutput, err error) {
	output = &PutRecruitingSelfScheduleCalendarOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutRecruitingSelfScheduleCalendarInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Recruiting_Self-Schedule_Calendar_Request"`
	PutRecruitingSelfScheduleCalendarRequestType
}

type PutRecruitingSelfScheduleCalendarOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Recruiting_Self-Schedule_Calendar_Response"`
	PutRecruitingSelfScheduleCalendarResponseType
}

// GetRecruitingAssessmentCategories (Get_Recruiting_Assessment_Categories)
// 
// Returns Recruiting Assessment Categories data.
func (c *Client) GetRecruitingAssessmentCategories(ctx context.Context, input *GetRecruitingAssessmentCategoriesInput) (output *GetRecruitingAssessmentCategoriesOutput, err error) {
	output = &GetRecruitingAssessmentCategoriesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetRecruitingAssessmentCategoriesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Recruiting_Assessment_Categories_Request"`
	GetRecruitingAssessmentCategoriesRequestType
}

type GetRecruitingAssessmentCategoriesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Recruiting_Assessment_Categories_Response"`
	GetRecruitingAssessmentCategoriesResponseType
}

// PutRecruitingAssessmentCategory (Put_Recruiting_Assessment_Category)
// 
// Create or Update Recruiting Assessment Catagories
func (c *Client) PutRecruitingAssessmentCategory(ctx context.Context, input *PutRecruitingAssessmentCategoryInput) (output *PutRecruitingAssessmentCategoryOutput, err error) {
	output = &PutRecruitingAssessmentCategoryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutRecruitingAssessmentCategoryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Recruiting_Assessment_Category_Request"`
	PutRecruitingAssessmentCategoryRequestType
}

type PutRecruitingAssessmentCategoryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Recruiting_Assessment_Category_Response"`
	PutRecruitingAssessmentCategoryResponseType
}

// GetRecruitingAssessmentCategorySecuritySegments (Get_Recruiting_Assessment_Category_Security_Segments)
// 
// Web service operation to retrieve Recruiting Assessment Category Security Segments.
func (c *Client) GetRecruitingAssessmentCategorySecuritySegments(ctx context.Context, input *GetRecruitingAssessmentCategorySecuritySegmentsInput) (output *GetRecruitingAssessmentCategorySecuritySegmentsOutput, err error) {
	output = &GetRecruitingAssessmentCategorySecuritySegmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetRecruitingAssessmentCategorySecuritySegmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Recruiting_Assessment_Category_Security_Segments_Request"`
	GetRecruitingAssessmentCategorySecuritySegmentsRequestType
}

type GetRecruitingAssessmentCategorySecuritySegmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Recruiting_Assessment_Category_Security_Segments_Response"`
	GetRecruitingAssessmentCategorySecuritySegmentsResponseType
}

// PutRecruitingAssessmentCategorySecuritySegment (Put_Recruiting_Assessment_Category_Security_Segment)
// 
// Web service operation to add, update or delete a Recruiting Assessment Category Security Segment.
func (c *Client) PutRecruitingAssessmentCategorySecuritySegment(ctx context.Context, input *PutRecruitingAssessmentCategorySecuritySegmentInput) (output *PutRecruitingAssessmentCategorySecuritySegmentOutput, err error) {
	output = &PutRecruitingAssessmentCategorySecuritySegmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutRecruitingAssessmentCategorySecuritySegmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Recruiting_Assessment_Category_Security_Segment_Request"`
	PutRecruitingAssessmentCategorySecuritySegmentRequestType
}

type PutRecruitingAssessmentCategorySecuritySegmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Recruiting_Assessment_Category_Security_Segment_Response"`
	PutRecruitingAssessmentCategorySecuritySegmentResponseType
}

// Offer (Offer)
// 
// Web service to initiate offer for a job application
func (c *Client) Offer(ctx context.Context, input *OfferInput) (output *OfferOutput, err error) {
	output = &OfferOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type OfferInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Offer_Request"`
	OfferRequestType
}

type OfferOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Offer_Response"`
	OfferResponseType
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

