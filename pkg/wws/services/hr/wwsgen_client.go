
// Package hr
//
// The Human Resources Web Service contains operations that expose Workday Human Capital Management Business Services data, including Employee, Contingent Worker and Organization information. This Web Service can be used for integration with enterprise systems including corporate directories, data analysis tools, email or other provisioning sub-systems, or any other systems needing Worker and/or Organization data.
package hr

import (
	"context"

	"github.com/ianlopshire/go-workday/pkg/wws"
)

const (
	CurrentVersion = "v31.2"
	ServiceName = "Human_Resources"
)

type Client struct {
	*wws.Client
}


// PutLocation (Put_Location)
// 
// This operation adds or updates a Location. Location Addresses, Usages, Email Addresses, Phone Numbers, Time Profile data, Location Attributes, Location Superior and Location Hierarchies to include location in can also be added and updated with this operation. Effective date must be blank for new locations to be consistent with new locations created in the UI. When effective date is blank, the value is effective as of the beginning of time. The Effective Date applies to these fields only: Location Name, Inactive, and Location Hierarchy Reference. If Location Hierarchy Reference is left blank for an existing location, location will be removed from all location hierarchies it is included in as of the effective date specified. If these fields are left blank, existing values will be removed: Altitude, Latitude, Longitude, Trade Name, Worksite Identification Code, Locale, User Language, Location Attribute, Location Type, Time Profile, Time Zone, Superior Location.
func (c *Client) PutLocation(ctx context.Context, input *PutLocationInput) (output *PutLocationOutput, err error) {
	output = &PutLocationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutLocationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Location_Request"`
	PutLocationRequestType
}

type PutLocationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Location_Response"`
	PutLocationResponseType
}

// FindBusinessSite (Find_Business_Site)
// 
// DEPRECATED: This web service operation is deprecated. Please use the Get Locations web service operation instead.  
// 
// This operation responds with a set of references to Business Sites that match the criteria specified in the request element.
func (c *Client) FindBusinessSite(ctx context.Context, input *FindBusinessSiteInput) (output *FindBusinessSiteOutput, err error) {
	output = &FindBusinessSiteOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type FindBusinessSiteInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Business_Site_Find"`
	BusinessSiteFindType
}

type FindBusinessSiteOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Business_Site_References"`
	BusinessSiteReferencesRootType
}

// GetOrganization (Get_Organization)
// 
// This operation retrieves data related to an Organization (e.g. Staffing Configuration, Structure, etc.).
func (c *Client) GetOrganization(ctx context.Context, input *GetOrganizationInput) (output *GetOrganizationOutput, err error) {
	output = &GetOrganizationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetOrganizationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Organization_Get"`
	OrganizationGetType
}

type GetOrganizationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Organization"`
	OrganizationType
}

// AddUpdateOrganization (Add_Update_Organization)
// 
// This operation adds a new Organization (or updates an existing Organization) with the supplied information. A new effective-dated organization name is automatically created if any attributes used for organization name are different as of specified effective date.
func (c *Client) AddUpdateOrganization(ctx context.Context, input *AddUpdateOrganizationInput) (output *AddUpdateOrganizationOutput, err error) {
	output = &AddUpdateOrganizationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AddUpdateOrganizationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Organization_Add_Update"`
	OrganizationAddUpdateType
}

type AddUpdateOrganizationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Organization_Reference"`
	OrganizationReferenceRootType
}

// GetEmployeePersonalInfo (Get_Employee_Personal_Info)
// 
// This operation retrieves data related to an Employee and his/her Personal (e.g. Biographic, Demographic, etc.) information.
func (c *Client) GetEmployeePersonalInfo(ctx context.Context, input *GetEmployeePersonalInfoInput) (output *GetEmployeePersonalInfoOutput, err error) {
	output = &GetEmployeePersonalInfoOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEmployeePersonalInfoInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Employee_Personal_Info_Get"`
	EmployeePersonalInfoGetType
}

type GetEmployeePersonalInfoOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Employee_Personal_Info"`
	EmployeePersonalInfoType
}

// GetContingentWorkerContractInfo (Get_Contingent_Worker_Contract_Info)
// 
// This operation retrieves data related to a Contingent Worker and his/her Contract information.
func (c *Client) GetContingentWorkerContractInfo(ctx context.Context, input *GetContingentWorkerContractInfoInput) (output *GetContingentWorkerContractInfoOutput, err error) {
	output = &GetContingentWorkerContractInfoOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetContingentWorkerContractInfoInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Contingent_Worker_Contract_Info_Get"`
	ContingentWorkerContractInfoGetType
}

type GetContingentWorkerContractInfoOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Contingent_Worker_Contract_Info"`
	ContingentWorkerContractInfoType
}

// GetWorkerProfile (Get_Worker_Profile)
// 
// This operation retrieves a subset of data related to a Worker (e.g. Employee and Contingent Worker) and his/her Employment/Contract, Personal and Compensation information.
func (c *Client) GetWorkerProfile(ctx context.Context, input *GetWorkerProfileInput) (output *GetWorkerProfileOutput, err error) {
	output = &GetWorkerProfileOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetWorkerProfileInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Worker_Profile_Get"`
	WorkerProfileGetType
}

type GetWorkerProfileOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Worker_Profile"`
	WorkerProfileType
}

// PutPreviousSystemJobHistory (Put_Previous_System_Job_History)
// 
// Loads history for a Worker whose history did not get loaded during the initial implementation period.  This operation allows free-form text entry of changes that occurred between the Worker&#39;s original hire date and the date that the implementation occurred.  Existing entries can be deleted or updated.
func (c *Client) PutPreviousSystemJobHistory(ctx context.Context, input *PutPreviousSystemJobHistoryInput) (output *PutPreviousSystemJobHistoryOutput, err error) {
	output = &PutPreviousSystemJobHistoryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPreviousSystemJobHistoryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Previous_System_Job_History_Request"`
	PutPreviousSystemJobHistoryRequestType
}

type PutPreviousSystemJobHistoryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Previous_System_Job_History_Response"`
	PutPreviousSystemJobHistoryResponseType
}

// DissolveOrganizationStructure (Dissolve_Organization_Structure)
// 
// This operation will dissolve an Organization structure (e.g. disconnect the subordinate and contained organizations).
func (c *Client) DissolveOrganizationStructure(ctx context.Context, input *DissolveOrganizationStructureInput) (output *DissolveOrganizationStructureOutput, err error) {
	output = &DissolveOrganizationStructureOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type DissolveOrganizationStructureInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Organization_Structure_Dissolve"`
	OrganizationStructureDissolveType
}

type DissolveOrganizationStructureOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc "`
	
}

// FindOrganization (Find_Organization)
// 
// This operation responds with a set of references to Organizations that match the criteria specified in the request element.
func (c *Client) FindOrganization(ctx context.Context, input *FindOrganizationInput) (output *FindOrganizationOutput, err error) {
	output = &FindOrganizationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type FindOrganizationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Organization_Find"`
	OrganizationFindType
}

type FindOrganizationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Organization_References"`
	OrganizationReferencesRootType
}

// UpdateEmployeePersonalInfo (Update_Employee_Personal_Info)
// 
// This operation updates Personal (e.g. Biographic, Demographic, etc.) information related to an Employee.
func (c *Client) UpdateEmployeePersonalInfo(ctx context.Context, input *UpdateEmployeePersonalInfoInput) (output *UpdateEmployeePersonalInfoOutput, err error) {
	output = &UpdateEmployeePersonalInfoOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type UpdateEmployeePersonalInfoInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Employee_Personal_Info_Update"`
	EmployeePersonalInfoUpdateType
}

type UpdateEmployeePersonalInfoOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc "`
	
}

// GetBusinessSite (Get_Business_Site)
// 
// DEPRECATED: This web service operation is deprecated. Please use the Get Locations web service operation instead.  
// 
// This operation retrieves data related to an Business Site.
func (c *Client) GetBusinessSite(ctx context.Context, input *GetBusinessSiteInput) (output *GetBusinessSiteOutput, err error) {
	output = &GetBusinessSiteOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBusinessSiteInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Business_Site_Get"`
	BusinessSiteGetType
}

type GetBusinessSiteOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Business_Site"`
	BusinessSiteType
}

// GetEmployeeEmploymentInfo (Get_Employee_Employment_Info)
// 
// This operation retrieves data related to an Employee and his/her Employment (e.g. Position, Job, Status, etc.) information.
func (c *Client) GetEmployeeEmploymentInfo(ctx context.Context, input *GetEmployeeEmploymentInfoInput) (output *GetEmployeeEmploymentInfoOutput, err error) {
	output = &GetEmployeeEmploymentInfoOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEmployeeEmploymentInfoInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Employee_Employment_Info_Get"`
	EmployeeEmploymentInfoGetType
}

type GetEmployeeEmploymentInfoOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Employee_Employment_Info"`
	EmployeeEmploymentInfoType
}

// GetContingentWorkerPersonalInfo (Get_Contingent_Worker_Personal_Info)
// 
// This operation retrieves data related to a Contingent Worker and his/her Personal (e.g. Biographic, Demographic, etc.) information.
func (c *Client) GetContingentWorkerPersonalInfo(ctx context.Context, input *GetContingentWorkerPersonalInfoInput) (output *GetContingentWorkerPersonalInfoOutput, err error) {
	output = &GetContingentWorkerPersonalInfoOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetContingentWorkerPersonalInfoInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Contingent_Worker_Personal_Info_Get"`
	ContingentWorkerPersonalInfoGetType
}

type GetContingentWorkerPersonalInfoOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Contingent_Worker_Personal_Info"`
	ContingentWorkerPersonalInfoType
}

// UpdateContingentWorkerPersonalInfo (Update_Contingent_Worker_Personal_Info)
// 
// This operation updates Personal (e.g. Biographic, Demographic, etc.) information related to a Contingent Worker.
func (c *Client) UpdateContingentWorkerPersonalInfo(ctx context.Context, input *UpdateContingentWorkerPersonalInfoInput) (output *UpdateContingentWorkerPersonalInfoOutput, err error) {
	output = &UpdateContingentWorkerPersonalInfoOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type UpdateContingentWorkerPersonalInfoInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Contingent_Worker_Personal_Info_Update"`
	ContingentWorkerPersonalInfoUpdateType
}

type UpdateContingentWorkerPersonalInfoOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc "`
	
}

// InactivateOrganization (Inactivate_Organization)
// 
// This operation will inactivate an Organization. If you do not choose to &#39;Keep in Hierarchy&#39; then the organization being inactivated will be moved out of the hierarchy. Default behavior for inactive subordinates / included organizations is that they remain as is. Default behavior for active subordinates is that they are moved to the superior unless you specify another organization to move them to. If the organization has subordinates or included content that you do not want to move to its superior or another organization, you can use the DISSOLVE_ORGANIZATION_STRUCTURE web service to disconnect it from the hierarchy and remove its subordinates and included organizations. In this web service, the Integration_ID_Reference must be the external system id.
func (c *Client) InactivateOrganization(ctx context.Context, input *InactivateOrganizationInput) (output *InactivateOrganizationOutput, err error) {
	output = &InactivateOrganizationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type InactivateOrganizationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Organization_Inactivate"`
	OrganizationInactivateType
}

type InactivateOrganizationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc "`
	
}

// GetEmployeeRelatedPersons (Get_Employee_Related_Persons)
// 
// This operation retrieves data related to an Employee&#39;s Related Persons (e.g. Dependents, Beneficiaries, etc.) and each of his/her Personal (e.g. Biographic, Demographic, etc.) information.
func (c *Client) GetEmployeeRelatedPersons(ctx context.Context, input *GetEmployeeRelatedPersonsInput) (output *GetEmployeeRelatedPersonsOutput, err error) {
	output = &GetEmployeeRelatedPersonsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEmployeeRelatedPersonsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Employee_Related_Persons_Get"`
	EmployeeRelatedPersonsGetType
}

type GetEmployeeRelatedPersonsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Employee_Related_Persons"`
	EmployeeRelatedPersonsType
}

// GetLocations (Get_Locations)
// 
// This operation retrieves data related to a Location for the specified criteria. The request criteria can be for a single entry based on a Reference ID, Location Name or all Locations will be retrieved if no criteria is specified.
func (c *Client) GetLocations(ctx context.Context, input *GetLocationsInput) (output *GetLocationsOutput, err error) {
	output = &GetLocationsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLocationsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Locations_Request"`
	GetLocationsRequestType
}

type GetLocationsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Locations_Response"`
	GetLocationsResponseType
}

// GetPreviousSystemJobHistory (Get_Previous_System_Job_History)
// 
// Returns a worker&#39;s job or position history imported from a previous system; that is, the user&#39;s system prior to migrating to Workday.
func (c *Client) GetPreviousSystemJobHistory(ctx context.Context, input *GetPreviousSystemJobHistoryInput) (output *GetPreviousSystemJobHistoryOutput, err error) {
	output = &GetPreviousSystemJobHistoryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPreviousSystemJobHistoryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Previous_System_Job_History_Request"`
	GetPreviousSystemJobHistoryRequestType
}

type GetPreviousSystemJobHistoryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Previous_System_Job_History_Response"`
	GetPreviousSystemJobHistoryResponseType
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

// UpdateEmployeeImage (Update_Employee_Image)
// 
// This operation updates an image (e.g. picture) related to a employee.
func (c *Client) UpdateEmployeeImage(ctx context.Context, input *UpdateEmployeeImageInput) (output *UpdateEmployeeImageOutput, err error) {
	output = &UpdateEmployeeImageOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type UpdateEmployeeImageInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Employee_Image_Update"`
	EmployeeImageUpdateType
}

type UpdateEmployeeImageOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc "`
	
}

// GetEmployee (Get_Employee)
// 
// This operation retrieves data related to an Employee and his/her Employment, Personal and Compensation.
func (c *Client) GetEmployee(ctx context.Context, input *GetEmployeeInput) (output *GetEmployeeOutput, err error) {
	output = &GetEmployeeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEmployeeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Employee_Get"`
	EmployeeGetType
}

type GetEmployeeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Employee"`
	EmployeeType
}

// GetContingentWorker (Get_Contingent_Worker)
// 
// This operation retrieves data related to a Contingent Worker and his/her Contract and Personal information.
func (c *Client) GetContingentWorker(ctx context.Context, input *GetContingentWorkerInput) (output *GetContingentWorkerOutput, err error) {
	output = &GetContingentWorkerOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetContingentWorkerInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Contingent_Worker_Get"`
	ContingentWorkerGetType
}

type GetContingentWorkerOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Contingent_Worker"`
	ContingentWorkerType
}

// GetWorkerEventHistory (Get_Worker_Event_History)
// 
// This operation retrieves references to all Events (created through workflow) associated with a Worker based on the Event Type and Date parameters.
func (c *Client) GetWorkerEventHistory(ctx context.Context, input *GetWorkerEventHistoryInput) (output *GetWorkerEventHistoryOutput, err error) {
	output = &GetWorkerEventHistoryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetWorkerEventHistoryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Worker_Event_History_Get"`
	WorkerEventHistoryGetType
}

type GetWorkerEventHistoryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Worker_Event_History"`
	WorkerEventHistoryType
}

// UpdateWorkdayAccount (Update_Workday_Account)
// 
// This operation updates an existing Workday Account (e.g. User Name / Password) with the supplied information.
func (c *Client) UpdateWorkdayAccount(ctx context.Context, input *UpdateWorkdayAccountInput) (output *UpdateWorkdayAccountOutput, err error) {
	output = &UpdateWorkdayAccountOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type UpdateWorkdayAccountInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Workday_Account_for_Worker_Update"`
	WorkdayAccountforWorkerUpdateType
}

type UpdateWorkdayAccountOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc "`
	
}

// AddWorkdayAccount (Add_Workday_Account)
// 
// This operation adds a new Workday Account (e.g. User Name / Password) with the supplied information.
func (c *Client) AddWorkdayAccount(ctx context.Context, input *AddWorkdayAccountInput) (output *AddWorkdayAccountOutput, err error) {
	output = &AddWorkdayAccountOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AddWorkdayAccountInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Workday_Account_for_Worker_Add"`
	WorkdayAccountforWorkerAddType
}

type AddWorkdayAccountOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc "`
	
}

// PutCompanyInsiderType (Put_Company_Insider_Type)
// 
// This operation adds or updates Company Insider Types.
func (c *Client) PutCompanyInsiderType(ctx context.Context, input *PutCompanyInsiderTypeInput) (output *PutCompanyInsiderTypeOutput, err error) {
	output = &PutCompanyInsiderTypeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCompanyInsiderTypeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Company_Insider_Type_Request"`
	PutCompanyInsiderTypeRequestType
}

type PutCompanyInsiderTypeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Company_Insider_Type_Response"`
	PutCompanyInsiderTypeResponseType
}

// GetCompanyInsiderTypes (Get_Company_Insider_Types)
// 
// This operation will get Company Insider Types for the specified criteria. The request criteria can be for a single entry based on a Reference ID or all Company Insider Types will be retrieved if no criteria is specified.  Company Insider Type data includes the Reference ID, the name and the description.
func (c *Client) GetCompanyInsiderTypes(ctx context.Context, input *GetCompanyInsiderTypesInput) (output *GetCompanyInsiderTypesOutput, err error) {
	output = &GetCompanyInsiderTypesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCompanyInsiderTypesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Company_Insider_Types_Request"`
	GetCompanyInsiderTypesRequestType
}

type GetCompanyInsiderTypesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Company_Insider_Types_Response"`
	GetCompanyInsiderTypesResponseType
}

// GetEmployeeImage (Get_Employee_Image)
// 
// Get a photographic image of this employee.
func (c *Client) GetEmployeeImage(ctx context.Context, input *GetEmployeeImageInput) (output *GetEmployeeImageOutput, err error) {
	output = &GetEmployeeImageOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEmployeeImageInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Employee_Image_Get"`
	EmployeeImageGetType
}

type GetEmployeeImageOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Employee_Image"`
	EmployeeImageType
}

// PutWorkShift (Put_Work_Shift)
// 
// This operation adds or updates a Work Shift.
func (c *Client) PutWorkShift(ctx context.Context, input *PutWorkShiftInput) (output *PutWorkShiftOutput, err error) {
	output = &PutWorkShiftOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutWorkShiftInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Work_Shift_Request"`
	PutWorkShiftRequestType
}

type PutWorkShiftOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Work_Shift_Response"`
	PutWorkShiftResponseType
}

// GetWorkShifts (Get_Work_Shifts)
// 
// This operation retrieves data related to a Work Shift.
func (c *Client) GetWorkShifts(ctx context.Context, input *GetWorkShiftsInput) (output *GetWorkShiftsOutput, err error) {
	output = &GetWorkShiftsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetWorkShiftsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Work_Shifts_Request"`
	GetWorkShiftsRequestType
}

type GetWorkShiftsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Work_Shifts_Response"`
	GetWorkShiftsResponseType
}

// PutJobCategory (Put_Job_Category)
// 
// This operation adds or updates a Job Category.
func (c *Client) PutJobCategory(ctx context.Context, input *PutJobCategoryInput) (output *PutJobCategoryOutput, err error) {
	output = &PutJobCategoryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutJobCategoryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Category_Request"`
	PutJobCategoryRequestType
}

type PutJobCategoryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Category_Response"`
	PutJobCategoryResponseType
}

// GetJobCategories (Get_Job_Categories)
// 
// This operation will get Job Categories for the specified criteria. The request criteria can be for a single entry based on a Reference ID or all Job Categories will be retrieved if no criteria is specified.  Job Category data includes the Reference ID, the name and the description and inactive flag.
func (c *Client) GetJobCategories(ctx context.Context, input *GetJobCategoriesInput) (output *GetJobCategoriesOutput, err error) {
	output = &GetJobCategoriesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetJobCategoriesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Categories_Request"`
	GetJobCategoriesRequestType
}

type GetJobCategoriesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Categories_Response"`
	GetJobCategoriesResponseType
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

// PutWorkersCompensationCode (Put_Workers_Compensation_Code)
// 
// This operation adds or updates a Workers&#39; Compensation Code.
func (c *Client) PutWorkersCompensationCode(ctx context.Context, input *PutWorkersCompensationCodeInput) (output *PutWorkersCompensationCodeOutput, err error) {
	output = &PutWorkersCompensationCodeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutWorkersCompensationCodeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Workers_Compensation_Code_Request"`
	PutWorkersCompensationCodeRequestType
}

type PutWorkersCompensationCodeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Workers_Compensation_Code_Response"`
	PutWorkersCompensationCodeResponseType
}

// GetWorkersCompensationCodes (Get_Workers_Compensation_Codes)
// 
// This operation will get Workers&#39; Compensation Codes for the specified criteria. The request criteria can be for a single entry based on a Reference ID; a specified country, country region, or business site; or all Workers&#39; Compensation Codes will be retrieved if no criteria is specified.
// Workers&#39; Compensation Code data includes the code, name, business site, country, country region and inactive flag.
func (c *Client) GetWorkersCompensationCodes(ctx context.Context, input *GetWorkersCompensationCodesInput) (output *GetWorkersCompensationCodesOutput, err error) {
	output = &GetWorkersCompensationCodesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetWorkersCompensationCodesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Workers_Compensation_Codes_Request"`
	GetWorkersCompensationCodesRequestType
}

type GetWorkersCompensationCodesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Workers_Compensation_Codes_Response"`
	GetWorkersCompensationCodesResponseType
}

// GetHolidayCalendars (Get_Holiday_Calendars)
// 
// Returns Holiday Calendars based on criteria.
func (c *Client) GetHolidayCalendars(ctx context.Context, input *GetHolidayCalendarsInput) (output *GetHolidayCalendarsOutput, err error) {
	output = &GetHolidayCalendarsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetHolidayCalendarsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Holiday_Calendars_Request"`
	GetHolidayCalendarsRequestType
}

type GetHolidayCalendarsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Holiday_Calendars_Response"`
	GetHolidayCalendarsResponseType
}

// PutHolidayCalendar (Put_Holiday_Calendar)
// 
// Adds or updates Holiday Calendar.
func (c *Client) PutHolidayCalendar(ctx context.Context, input *PutHolidayCalendarInput) (output *PutHolidayCalendarOutput, err error) {
	output = &PutHolidayCalendarOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutHolidayCalendarInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Holiday_Calendar_Request"`
	PutHolidayCalendarRequestType
}

type PutHolidayCalendarOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Holiday_Calendar_Response"`
	PutHolidayCalendarResponseType
}

// GetWorkScheduleCalendars (Get_Work_Schedule_Calendars)
// 
// View Work Schedule Calendars.
func (c *Client) GetWorkScheduleCalendars(ctx context.Context, input *GetWorkScheduleCalendarsInput) (output *GetWorkScheduleCalendarsOutput, err error) {
	output = &GetWorkScheduleCalendarsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetWorkScheduleCalendarsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Work_Schedule_Calendars_Request"`
	GetWorkScheduleCalendarsRequestType
}

type GetWorkScheduleCalendarsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Work_Schedule_Calendars_Response"`
	GetWorkScheduleCalendarsResponseType
}

// PutWorkScheduleCalendar (Put_Work_Schedule_Calendar)
// 
// Adds or updates Work Schedule Calendar
func (c *Client) PutWorkScheduleCalendar(ctx context.Context, input *PutWorkScheduleCalendarInput) (output *PutWorkScheduleCalendarOutput, err error) {
	output = &PutWorkScheduleCalendarOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutWorkScheduleCalendarInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Work_Schedule_Calendar_Request"`
	PutWorkScheduleCalendarRequestType
}

type PutWorkScheduleCalendarOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Work_Schedule_Calendar_Response"`
	PutWorkScheduleCalendarResponseType
}

// GetSupervisoryOrganizationAssignmentRestrictions (Get_Supervisory_Organization_Assignment_Restrictions)
// 
// This operation retrieves the Organization Assignment Default Values and Allowed Values for Supervisory Organizations.
func (c *Client) GetSupervisoryOrganizationAssignmentRestrictions(ctx context.Context, input *GetSupervisoryOrganizationAssignmentRestrictionsInput) (output *GetSupervisoryOrganizationAssignmentRestrictionsOutput, err error) {
	output = &GetSupervisoryOrganizationAssignmentRestrictionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetSupervisoryOrganizationAssignmentRestrictionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Organization_Assignment_Restrictions_Request"`
	GetOrganizationAssignmentRestrictionsRequestType
}

type GetSupervisoryOrganizationAssignmentRestrictionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Organization_Assignment_Restrictions_Response"`
	GetOrganizationAssignmentRestrictionsResponseType
}

// PutSupervisoryOrganizationAssignmentRestrictions (Put_Supervisory_Organization_Assignment_Restrictions)
// 
// This operation updates the Organization Assignment Default Value and Allowed Values for a Supervisory Organization. Updates can be made by organization type.  Note, the default value of Replace_All attribute is FALSE. This is different from most other webservices with this attribute because in this case, FALSE is more conservative. Default and allowed values for all organization types are inherited from the superior when no values have been specified for any organization types. Once a default or allowed value is specified for any organization type, the inheritance is severed. In the UI, inherited values are persisted when inheritance is severed. In order to mimic this behavior in the web service, inherited values will be persisted unless the &#39;Replace All&#39; or &#39;Delete&#39; options are set to TRUE.
func (c *Client) PutSupervisoryOrganizationAssignmentRestrictions(ctx context.Context, input *PutSupervisoryOrganizationAssignmentRestrictionsInput) (output *PutSupervisoryOrganizationAssignmentRestrictionsOutput, err error) {
	output = &PutSupervisoryOrganizationAssignmentRestrictionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutSupervisoryOrganizationAssignmentRestrictionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Organization_Assignment_Restrictions_Request"`
	PutOrganizationAssignmentRestrictionsRequestType
}

type PutSupervisoryOrganizationAssignmentRestrictionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Organization_Assignment_Restrictions_Response"`
	PutOrganizationAssignmentRestrictionsResponseType
}

// GetFrequencies (Get_Frequencies)
// 
// This operation retrieves data related to Frequencies.
func (c *Client) GetFrequencies(ctx context.Context, input *GetFrequenciesInput) (output *GetFrequenciesOutput, err error) {
	output = &GetFrequenciesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetFrequenciesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Frequencies_Request"`
	GetFrequenciesRequestType
}

type GetFrequenciesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Frequencies_Response"`
	GetFrequenciesResponseType
}

// PutFrequency (Put_Frequency)
// 
// This operation adds or updates a Frequency.
func (c *Client) PutFrequency(ctx context.Context, input *PutFrequencyInput) (output *PutFrequencyOutput, err error) {
	output = &PutFrequencyOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutFrequencyInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Frequency_Request"`
	PutFrequencyRequestType
}

type PutFrequencyOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Frequency_Response"`
	PutFrequencyResponseType
}

// GetDisabilities (Get_Disabilities)
// 
// This operation retrieves the disabilities that are currently defined.
func (c *Client) GetDisabilities(ctx context.Context, input *GetDisabilitiesInput) (output *GetDisabilitiesOutput, err error) {
	output = &GetDisabilitiesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetDisabilitiesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Disabilities_Request"`
	GetDisabilitiesRequestType
}

type GetDisabilitiesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Disabilities_Response"`
	GetDisabilitiesResponseType
}

// PutDisability (Put_Disability)
// 
// This operation adds or updates a Disability.
func (c *Client) PutDisability(ctx context.Context, input *PutDisabilityInput) (output *PutDisabilityOutput, err error) {
	output = &PutDisabilityOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutDisabilityInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Disability_Request"`
	PutDisabilityRequestType
}

type PutDisabilityOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Disability_Response"`
	PutDisabilityResponseType
}

// PutJobProfile (Put_Job_Profile)
// 
// This operations adds or updates a Job Profile.
func (c *Client) PutJobProfile(ctx context.Context, input *PutJobProfileInput) (output *PutJobProfileOutput, err error) {
	output = &PutJobProfileOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutJobProfileInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Profile_Request"`
	PutJobProfileRequestType
}

type PutJobProfileOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Job_Profile_Response"`
	PutJobProfileResponseType
}

// GetJobProfiles (Get_Job_Profiles)
// 
// This operation retrieves data related to a Job Profile for the specified criteria. The request criteria can be for a single entry based on a Reference ID,  management level, job level, job family, job category, company insider type, job classification, workers compensation code, pay rate type , Job Exempt location context, work hours profiles or all Job Profiles will be retrieved if no criteria is specified. 
// The data returned is organized into different response groups which can be include or not.
func (c *Client) GetJobProfiles(ctx context.Context, input *GetJobProfilesInput) (output *GetJobProfilesOutput, err error) {
	output = &GetJobProfilesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetJobProfilesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Profiles_Request"`
	GetJobProfilesRequestType
}

type GetJobProfilesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Job_Profiles_Response"`
	GetJobProfilesResponseType
}

// GetEthnicities (Get_Ethnicities)
// 
// This operation retrieves data related to Ethnicities.
func (c *Client) GetEthnicities(ctx context.Context, input *GetEthnicitiesInput) (output *GetEthnicitiesOutput, err error) {
	output = &GetEthnicitiesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEthnicitiesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Ethnicities_Request"`
	GetEthnicitiesRequestType
}

type GetEthnicitiesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Ethnicities_Response"`
	GetEthnicitiesResponseType
}

// PutEthnicity (Put_Ethnicity)
// 
// This operation adds or updates an Ethnicity.
func (c *Client) PutEthnicity(ctx context.Context, input *PutEthnicityInput) (output *PutEthnicityOutput, err error) {
	output = &PutEthnicityOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutEthnicityInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Ethnicity_Request"`
	PutEthnicityRequestType
}

type PutEthnicityOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Ethnicity_Response"`
	PutEthnicityResponseType
}

// GetTrainingTypes (Get_Training_Types)
// 
// This operation retrieves Training Types defined in the system.
func (c *Client) GetTrainingTypes(ctx context.Context, input *GetTrainingTypesInput) (output *GetTrainingTypesOutput, err error) {
	output = &GetTrainingTypesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetTrainingTypesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Training_Types_Request"`
	GetTrainingTypesRequestType
}

type GetTrainingTypesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Training_Types_Response"`
	GetTrainingTypesResponseType
}

// PutTrainingType (Put_Training_Type)
// 
// This operation adds or updates a Training Type
func (c *Client) PutTrainingType(ctx context.Context, input *PutTrainingTypeInput) (output *PutTrainingTypeOutput, err error) {
	output = &PutTrainingTypeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutTrainingTypeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Training_Type_Request"`
	PutTrainingTypeRequestType
}

type PutTrainingTypeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Training_Type_Response"`
	PutTrainingTypeResponseType
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

// PutUserBasedSecurityGroupAssignments (Put_User-Based_Security_Group_Assignments)
// 
// This operation performs a FULL REPLACEMENT of User-Based Security Groups. It does not add to existing entries. Assigns 1 or more User-Based Security Groups to a Workday Account.
func (c *Client) PutUserBasedSecurityGroupAssignments(ctx context.Context, input *PutUserBasedSecurityGroupAssignmentsInput) (output *PutUserBasedSecurityGroupAssignmentsOutput, err error) {
	output = &PutUserBasedSecurityGroupAssignmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutUserBasedSecurityGroupAssignmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Assign_User-Based_Security_Groups_Request"`
	PutAssignUserBasedSecurityGroupsRequestType
}

type PutUserBasedSecurityGroupAssignmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Assign_User-Based_Security_Groups_Response"`
	PutAssignUserBasedSecurityGroupsResponseType
}

// GetUserBasedSecurityGroupAssignments (Get_User-Based_Security_Group_Assignments)
// 
// Retrieves Workday Account and it&#39;s assigned user-based security groups.
func (c *Client) GetUserBasedSecurityGroupAssignments(ctx context.Context, input *GetUserBasedSecurityGroupAssignmentsInput) (output *GetUserBasedSecurityGroupAssignmentsOutput, err error) {
	output = &GetUserBasedSecurityGroupAssignmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetUserBasedSecurityGroupAssignmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Assign_User-Based_Security_Groups_Request"`
	GetAssignUserBasedSecurityGroupsRequestType
}

type GetUserBasedSecurityGroupAssignmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Assign_User-Based_Security_Groups_Response"`
	GetAssignUserBasedSecurityGroupsResponseType
}

// GetDifficultytoFill (Get_Difficulty_to_Fill)
// 
// Gets the difficulty level of filling a job profile.
func (c *Client) GetDifficultytoFill(ctx context.Context, input *GetDifficultytoFillInput) (output *GetDifficultytoFillOutput, err error) {
	output = &GetDifficultytoFillOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetDifficultytoFillInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Difficulty_to_Fill_Request"`
	GetDifficultytoFillRequestType
}

type GetDifficultytoFillOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Difficulty_to_Fill_Response"`
	GetDifficultytoFillResponseType
}

// PutDifficultytoFill (Put_Difficulty_to_Fill)
// 
// This operation adds or updates a job profile&#39;s Difficulty to Fill.
func (c *Client) PutDifficultytoFill(ctx context.Context, input *PutDifficultytoFillInput) (output *PutDifficultytoFillOutput, err error) {
	output = &PutDifficultytoFillOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutDifficultytoFillInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Difficulty_to_Fill_Request"`
	PutDifficultytoFillRequestType
}

type PutDifficultytoFillOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Difficulty_to_Fill_Response"`
	PutDifficultytoFillResponseType
}

// GetWorkerPhotos (Get_Worker_Photos)
// 
// Gets a Photo Image of the Worker
func (c *Client) GetWorkerPhotos(ctx context.Context, input *GetWorkerPhotosInput) (output *GetWorkerPhotosOutput, err error) {
	output = &GetWorkerPhotosOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetWorkerPhotosInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Worker_Photos_Request"`
	GetWorkerPhotosRequestType
}

type GetWorkerPhotosOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Worker_Photos_Response"`
	GetWorkerPhotosResponseType
}

// PutWorkerPhoto (Put_Worker_Photo)
// 
// This operation adds or updates a workers photo.
func (c *Client) PutWorkerPhoto(ctx context.Context, input *PutWorkerPhotoInput) (output *PutWorkerPhotoOutput, err error) {
	output = &PutWorkerPhotoOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutWorkerPhotoInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Worker_Photo_Request"`
	PutWorkerPhotoRequestType
}

type PutWorkerPhotoOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Worker_Photo_Response"`
	PutWorkerPhotoResponseType
}

// ManageUnionMembership (Manage_Union_Membership)
// 
// Operation that will add a new worker to a union.  Generates a Union Membership Event.
func (c *Client) ManageUnionMembership(ctx context.Context, input *ManageUnionMembershipInput) (output *ManageUnionMembershipOutput, err error) {
	output = &ManageUnionMembershipOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ManageUnionMembershipInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Union_Membership_Request"`
	ManageUnionMembershipRequestType
}

type ManageUnionMembershipOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Union_Membership_Response"`
	ManageUnionMembershipResponseType
}

// GetAcademicRanks (Get_Academic_Ranks)
// 
// Get Academic Rank Web Service
func (c *Client) GetAcademicRanks(ctx context.Context, input *GetAcademicRanksInput) (output *GetAcademicRanksOutput, err error) {
	output = &GetAcademicRanksOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAcademicRanksInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Academic_Ranks_Request"`
	GetAcademicRanksRequestType
}

type GetAcademicRanksOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Academic_Ranks_Response"`
	GetAcademicRanksResponseType
}

// PutAcademicRank (Put_Academic_Rank)
// 
// This operation adds or updates Academic Ranks
func (c *Client) PutAcademicRank(ctx context.Context, input *PutAcademicRankInput) (output *PutAcademicRankOutput, err error) {
	output = &PutAcademicRankOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutAcademicRankInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Academic_Rank_Request"`
	PutAcademicRankRequestType
}

type PutAcademicRankOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Academic_Rank_Response"`
	PutAcademicRankResponseType
}

// EndAcademicAppointment (End_Academic_Appointment)
// 
// This operation will invoke the Business Process to end an Academic Appointment.
func (c *Client) EndAcademicAppointment(ctx context.Context, input *EndAcademicAppointmentInput) (output *EndAcademicAppointmentOutput, err error) {
	output = &EndAcademicAppointmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EndAcademicAppointmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc End_Academic_Appointment_Request"`
	EndAcademicAppointmentRequestType
}

type EndAcademicAppointmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Academic_Appointment_Response"`
	AcademicAppointmentResponseType
}

// GetProvisioningGroups (Get_Provisioning_Groups)
// 
// This web service operation will get provisioning groups. It will return the provisioning groups and their associated information excluding any provisioning group assignments. The request can be for a single transaction based on reference, or all transactions can be retrieved if no reference is specified.
func (c *Client) GetProvisioningGroups(ctx context.Context, input *GetProvisioningGroupsInput) (output *GetProvisioningGroupsOutput, err error) {
	output = &GetProvisioningGroupsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetProvisioningGroupsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Provisioning_Groups_Request"`
	GetProvisioningGroupsRequestType
}

type GetProvisioningGroupsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Provisioning_Groups_Response"`
	ProvisioningGroupsResponseType
}

// AddAcademicAppointment (Add_Academic_Appointment)
// 
// This operation will invoke the Business Process to add an Academic Appointment.
func (c *Client) AddAcademicAppointment(ctx context.Context, input *AddAcademicAppointmentInput) (output *AddAcademicAppointmentOutput, err error) {
	output = &AddAcademicAppointmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AddAcademicAppointmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Add_Academic_Appointment_Request"`
	AddAcademicAppointmentRequestType
}

type AddAcademicAppointmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Academic_Appointment_Response"`
	AcademicAppointmentResponseType
}

// UpdateAcademicAppointment (Update_Academic_Appointment)
// 
// This operation will invoke the Business Process to update an Academic Appointment.
func (c *Client) UpdateAcademicAppointment(ctx context.Context, input *UpdateAcademicAppointmentInput) (output *UpdateAcademicAppointmentOutput, err error) {
	output = &UpdateAcademicAppointmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type UpdateAcademicAppointmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Update_Academic_Appointment_Request"`
	UpdateAcademicAppointmentRequestType
}

type UpdateAcademicAppointmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Academic_Appointment_Response"`
	AcademicAppointmentResponseType
}

// PutProvisioningGroup (Put_Provisioning_Group)
// 
// This web service operation updates an existing provisioning group if a reference is provided on the request, else it creates a new one. The provisioning group data to be specified as part of the request includes the name and description. Assignment information needs to be submitted using the operation Put Provisioning Group Assignment.
func (c *Client) PutProvisioningGroup(ctx context.Context, input *PutProvisioningGroupInput) (output *PutProvisioningGroupOutput, err error) {
	output = &PutProvisioningGroupOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutProvisioningGroupInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Provisioning_Group_Request"`
	PutProvisioningGroupRequestType
}

type PutProvisioningGroupOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Provisioning_Group_Response"`
	PutProvisioningGroupResponseType
}

// GetProvisioningGroupAssignments (Get_Provisioning_Group_Assignments)
// 
// This web service operation will get assignments of persons to provisioning groups. For each assignment a reference to the provisioning group, the person, the worker and status information will be returned. The request can be made for a list of persons as specified in the request criteria. Alternatively a list of workers can be specified there. Requests can also be made for individual assignments by specifying an assignment reference. If no request references or criteria are specified, the assignments for all persons will be returned.
func (c *Client) GetProvisioningGroupAssignments(ctx context.Context, input *GetProvisioningGroupAssignmentsInput) (output *GetProvisioningGroupAssignmentsOutput, err error) {
	output = &GetProvisioningGroupAssignmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetProvisioningGroupAssignmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Provisioning_Group_Assignments_Request"`
	GetProvisioningGroupAssignmentsRequestType
}

type GetProvisioningGroupAssignmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Provisioning_Group_Assignments_Response"`
	ProvisioningGroupAssignmentsResponseType
}

// PutProvisioningGroupAssignment (Put_Provisioning_Group_Assignment)
// 
// This operation creates or updates the assignment of a person to a provisioning group. If the an assignment between the person and the provisioning group exists, it will be updated with the data given in the request, else a new one will be created. The data to be specified include a reference to the person and the provisioning group. 
// Alternatively to the person a worker can be specified. In this case the operation will retrieve the person for that worker and create or update the assignment between that person and the group.
// The assignment status does not have to be specified on the request. If it is omitted, a default status of “Assigned” will be set.
func (c *Client) PutProvisioningGroupAssignment(ctx context.Context, input *PutProvisioningGroupAssignmentInput) (output *PutProvisioningGroupAssignmentOutput, err error) {
	output = &PutProvisioningGroupAssignmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutProvisioningGroupAssignmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Provisioning_Group_Assignment_Request"`
	PutProvisioningGroupAssignmentRequestType
}

type PutProvisioningGroupAssignmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Provisioning_Group_Assignment_Response"`
	PutProvisioningGroupAssignmentResponseType
}

// GetSearchSettings (Get_Search_Settings)
// 
// This operation will get information for search settings.
func (c *Client) GetSearchSettings(ctx context.Context, input *GetSearchSettingsInput) (output *GetSearchSettingsOutput, err error) {
	output = &GetSearchSettingsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetSearchSettingsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Search_Settings_Request"`
	GetSearchSettingsRequestType
}

type GetSearchSettingsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Search_Settings_Response"`
	GetSearchSettingsResponseType
}

// PutSearchSettings (Put_Search_Settings)
// 
// This operation will put search configurations into workday.
func (c *Client) PutSearchSettings(ctx context.Context, input *PutSearchSettingsInput) (output *PutSearchSettingsOutput, err error) {
	output = &PutSearchSettingsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutSearchSettingsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Search_Settings_Request"`
	PutSearchSettingsRequestType
}

type PutSearchSettingsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Search_Settings_Response"`
	PutSearchSettingsResponseType
}

// PutAddressesforCountryFormatExtension (Put_Addresses_for_Country_Format_Extension)
// 
// Use this operation ONLY for the country format extension and for local script addresses population. 
// 
// 1. Country Format Extension
// Since many countries require more granular address components than Workday currently supports, the Alternate Address Formats allow Workday customers to define a Basic and an Extended Address Format metadata for these countries.  Customers can then control address formats using Tenant Setup - Global when they want to use extended formats for a country.  This operation uploads addresses for a specified country after the addresses are modified to fit a new extended format.  Note that the attribute use Extended Address Format in Tenant Setup - Global must be enabled for the country beforeusing this operation to uploading addresses.
// 
// 2. Local Script Address Population
// Local script address components were added in W18 to allow address input in local script.  This operation uploads addresses with new local address components.
func (c *Client) PutAddressesforCountryFormatExtension(ctx context.Context, input *PutAddressesforCountryFormatExtensionInput) (output *PutAddressesforCountryFormatExtensionOutput, err error) {
	output = &PutAddressesforCountryFormatExtensionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutAddressesforCountryFormatExtensionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Addresses_for_Country_Format_Extension_Request"`
	PutAddressesforCountryFormatExtensionRequestType
}

type PutAddressesforCountryFormatExtensionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Addresses_for_Country_Format_Extension_Response"`
	PutAddressesforCountryFormatExtensionResponseType
}

// MaintainContactInformation (Maintain_Contact_Information)
// 
// Creates or updates contact information for a person. Contact information includes email addresses, physical addresses, phone numbers, web addresses, and instant messenger contact information. Changes are routed through existing Contact Information for Person Event business process for necessary approvals.
// Submitted contact information by default replaces existing entries of the same type. For example, submitting a new email address  replaces the current email address. Set the Add Only flag to override this default behavior. When this flag is set to true, existing information is not replaced in cases where multiple entries are allowed.
func (c *Client) MaintainContactInformation(ctx context.Context, input *MaintainContactInformationInput) (output *MaintainContactInformationOutput, err error) {
	output = &MaintainContactInformationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type MaintainContactInformationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Maintain_Contact_Information_for_Person_Event_Request"`
	MaintainContactInformationforPersonEventRequestType
}

type MaintainContactInformationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Maintain_Contact_Information_for_Person_Event_Response"`
	MaintainContactInformationforPersonEventResponseType
}

// GetServiceCenterRepresentatives (Get_Service_Center_Representatives)
// 
// Get Service Center Representatives.
func (c *Client) GetServiceCenterRepresentatives(ctx context.Context, input *GetServiceCenterRepresentativesInput) (output *GetServiceCenterRepresentativesOutput, err error) {
	output = &GetServiceCenterRepresentativesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetServiceCenterRepresentativesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Service_Center_Representatives_Request"`
	GetServiceCenterRepresentativesRequestType
}

type GetServiceCenterRepresentativesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Service_Center_Representatives_Response"`
	GetServiceCenterRepresentativesResponseType
}

// PutServiceCenterRepresentative (Put_Service_Center_Representative)
// 
// Adds or updates a Service Center Representative.
func (c *Client) PutServiceCenterRepresentative(ctx context.Context, input *PutServiceCenterRepresentativeInput) (output *PutServiceCenterRepresentativeOutput, err error) {
	output = &PutServiceCenterRepresentativeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutServiceCenterRepresentativeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Service_Center_Representative_Request"`
	PutServiceCenterRepresentativeRequestType
}

type PutServiceCenterRepresentativeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Service_Center_Representative_Response"`
	PutServiceCenterRepresentativeResponseType
}

// GetServiceCenterRepresentativeWorkdayAccounts (Get_Service_Center_Representative_Workday_Accounts)
// 
// Get Service Center Representative Workday Accounts.
func (c *Client) GetServiceCenterRepresentativeWorkdayAccounts(ctx context.Context, input *GetServiceCenterRepresentativeWorkdayAccountsInput) (output *GetServiceCenterRepresentativeWorkdayAccountsOutput, err error) {
	output = &GetServiceCenterRepresentativeWorkdayAccountsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetServiceCenterRepresentativeWorkdayAccountsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Service_Center_Representative_Workday_Accounts_Request"`
	GetServiceCenterRepresentativeWorkdayAccountsRequestType
}

type GetServiceCenterRepresentativeWorkdayAccountsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Service_Center_Representative_Workday_Accounts_Response"`
	GetServiceCenterRepresentativeWorkdayAccountsResponseType
}

// PutServiceCenterRepresentativeWorkdayAccount (Put_Service_Center_Representative_Workday_Account)
// 
// Adds or updates a Service Center Representative Workday Account.
func (c *Client) PutServiceCenterRepresentativeWorkdayAccount(ctx context.Context, input *PutServiceCenterRepresentativeWorkdayAccountInput) (output *PutServiceCenterRepresentativeWorkdayAccountOutput, err error) {
	output = &PutServiceCenterRepresentativeWorkdayAccountOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutServiceCenterRepresentativeWorkdayAccountInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Service_Center_Representative_Workday_Account_Request"`
	PutServiceCenterRepresentativeWorkdayAccountRequestType
}

type PutServiceCenterRepresentativeWorkdayAccountOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Service_Center_Representative_Workday_Account_Response"`
	PutServiceCenterRepresentativeWorkdayAccountResponseType
}

// ChangeBackgroundCheckStatus (Change_Background_Check_Status)
// 
// Updates the employee&#39;s background check status by adding a new instance of Background Check Descriptor, which stores Status Date, Status, and Comment.
func (c *Client) ChangeBackgroundCheckStatus(ctx context.Context, input *ChangeBackgroundCheckStatusInput) (output *ChangeBackgroundCheckStatusOutput, err error) {
	output = &ChangeBackgroundCheckStatusOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeBackgroundCheckStatusInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Background_Check_Status_Request"`
	ChangeBackgroundCheckStatusRequestType
}

type ChangeBackgroundCheckStatusOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Background_Check_Status_Response"`
	ChangeBackgroundCheckStatusResponseType
}

// GetAcademicUnits (Get_Academic_Units)
// 
// Returns detailed information for Academic Units.
func (c *Client) GetAcademicUnits(ctx context.Context, input *GetAcademicUnitsInput) (output *GetAcademicUnitsOutput, err error) {
	output = &GetAcademicUnitsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAcademicUnitsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Academic_Units_Request"`
	GetAcademicUnitsRequestType
}

type GetAcademicUnitsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Academic_Units_Response"`
	GetAcademicUnitsResponseType
}

// PutAcademicUnit (Put_Academic_Unit)
// 
// This operation adds or updates Academic Units
func (c *Client) PutAcademicUnit(ctx context.Context, input *PutAcademicUnitInput) (output *PutAcademicUnitOutput, err error) {
	output = &PutAcademicUnitOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutAcademicUnitInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Academic_Unit_Request"`
	PutAcademicUnitRequestType
}

type PutAcademicUnitOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Academic_Unit_Response"`
	PutAcademicUnitResponseType
}

// GetAcademicUnitHierarchies (Get_Academic_Unit_Hierarchies)
// 
// This operation is used to retrieve Academic Unit Hierarchies.
func (c *Client) GetAcademicUnitHierarchies(ctx context.Context, input *GetAcademicUnitHierarchiesInput) (output *GetAcademicUnitHierarchiesOutput, err error) {
	output = &GetAcademicUnitHierarchiesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAcademicUnitHierarchiesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Academic_Unit_Hierarchies_Request"`
	GetAcademicUnitHierarchiesRequestType
}

type GetAcademicUnitHierarchiesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Academic_Unit_Hierarchies_Response"`
	GetAcademicUnitHierarchiesResponseType
}

// PutAcademicUnitHierarchy (Put_Academic_Unit_Hierarchy)
// 
// This operation is used to create an Academic Unit Hierarchy.
func (c *Client) PutAcademicUnitHierarchy(ctx context.Context, input *PutAcademicUnitHierarchyInput) (output *PutAcademicUnitHierarchyOutput, err error) {
	output = &PutAcademicUnitHierarchyOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutAcademicUnitHierarchyInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Academic_Unit_Hierarchy_Request"`
	PutAcademicUnitHierarchyRequestType
}

type PutAcademicUnitHierarchyOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Academic_Unit_Hierarchy_Response"`
	PutAcademicUnitHierarchyResponseType
}

// GetPoliticalAffiliations (Get_Political_Affiliations)
// 
// This operation retrieves the political affiliations that are currently defined.
func (c *Client) GetPoliticalAffiliations(ctx context.Context, input *GetPoliticalAffiliationsInput) (output *GetPoliticalAffiliationsOutput, err error) {
	output = &GetPoliticalAffiliationsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPoliticalAffiliationsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Political_Affiliations_Request"`
	GetPoliticalAffiliationsRequestType
}

type GetPoliticalAffiliationsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Political_Affiliations_Response"`
	GetPoliticalAffiliationsResponseType
}

// PutPoliticalAffiliation (Put_Political_Affiliation)
// 
// This operation adds or updates a Political Affiliation.
func (c *Client) PutPoliticalAffiliation(ctx context.Context, input *PutPoliticalAffiliationInput) (output *PutPoliticalAffiliationOutput, err error) {
	output = &PutPoliticalAffiliationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPoliticalAffiliationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Political_Affiliation_Request"`
	PutPoliticalAffiliationRequestType
}

type PutPoliticalAffiliationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Political_Affiliation_Response"`
	PutPoliticalAffiliationResponseType
}

// ChangePersonalInformation (Change_Personal_Information)
// 
// Sets a worker&#39;s personal information. Uses the Personal Information Change business process.
func (c *Client) ChangePersonalInformation(ctx context.Context, input *ChangePersonalInformationInput) (output *ChangePersonalInformationOutput, err error) {
	output = &ChangePersonalInformationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangePersonalInformationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Personal_Information_Request"`
	ChangePersonalInformationRequestType
}

type ChangePersonalInformationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Personal_Information_Response"`
	ChangePersonalInformationResponseType
}

// ChangeGovernmentIDs (Change_Government_IDs)
// 
// This web service allows the updating of Government IDs for a worker.
func (c *Client) ChangeGovernmentIDs(ctx context.Context, input *ChangeGovernmentIDsInput) (output *ChangeGovernmentIDsOutput, err error) {
	output = &ChangeGovernmentIDsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeGovernmentIDsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Government_IDs_Request"`
	ChangeGovernmentIDsRequestType
}

type ChangeGovernmentIDsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Government_IDs_Response"`
	ChangeGovernmentIDsResponseType
}

// ChangeLegalName (Change_Legal_Name)
// 
// Sets a worker&#39;s legal name. Uses the Legal Name Change business process.
func (c *Client) ChangeLegalName(ctx context.Context, input *ChangeLegalNameInput) (output *ChangeLegalNameOutput, err error) {
	output = &ChangeLegalNameOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeLegalNameInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Legal_Name_Request"`
	ChangeLegalNameRequestType
}

type ChangeLegalNameOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Legal_Name_Response"`
	ChangeLegalNameResponseType
}

// ChangePreferredName (Change_Preferred_Name)
// 
// Sets a worker&#39;s preferred name. Uses the Preferred Name Change business process.
func (c *Client) ChangePreferredName(ctx context.Context, input *ChangePreferredNameInput) (output *ChangePreferredNameOutput, err error) {
	output = &ChangePreferredNameOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangePreferredNameInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Preferred_Name_Request"`
	ChangePreferredNameRequestType
}

type ChangePreferredNameOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Preferred_Name_Response"`
	ChangePreferredNameResponseType
}

// ChangeAdditionalNames (Change_Additional_Names)
// 
// Sets a worker&#39;s additional names. Does not use an event.
func (c *Client) ChangeAdditionalNames(ctx context.Context, input *ChangeAdditionalNamesInput) (output *ChangeAdditionalNamesOutput, err error) {
	output = &ChangeAdditionalNamesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeAdditionalNamesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Additional_Names_Request"`
	ChangeAdditionalNamesRequestType
}

type ChangeAdditionalNamesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Additional_Names_Response"`
	ChangeAdditionalNamesResponseType
}

// ChangePassportsandVisas (Change_Passports_and_Visas)
// 
// This web service allows the updating of Passports and Visas for a worker.
func (c *Client) ChangePassportsandVisas(ctx context.Context, input *ChangePassportsandVisasInput) (output *ChangePassportsandVisasOutput, err error) {
	output = &ChangePassportsandVisasOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangePassportsandVisasInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Passports_and_Visas_Request"`
	ChangePassportsandVisasRequestType
}

type ChangePassportsandVisasOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Passports_and_Visas_Response"`
	ChangePassportsandVisasResponseType
}

// ChangeLicenses (Change_Licenses)
// 
// This web service allows the updating of Licenses for a worker.
func (c *Client) ChangeLicenses(ctx context.Context, input *ChangeLicensesInput) (output *ChangeLicensesOutput, err error) {
	output = &ChangeLicensesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeLicensesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Licenses_Request"`
	ChangeLicensesRequestType
}

type ChangeLicensesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Licenses_Response"`
	ChangeLicensesResponseType
}

// ChangeOtherIDs (Change_Other_IDs)
// 
// This web service allows the updating of Other IDs for a worker.
func (c *Client) ChangeOtherIDs(ctx context.Context, input *ChangeOtherIDsInput) (output *ChangeOtherIDsOutput, err error) {
	output = &ChangeOtherIDsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeOtherIDsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Other_IDs_Request"`
	ChangeOtherIDsRequestType
}

type ChangeOtherIDsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Other_IDs_Response"`
	ChangeOtherIDsResponseType
}

// PutFormerWorker (Put_Former_Worker)
// 
// This operation allows Former Workers to be added and edited.
// Former Worker Storage is for tracking data for individuals that have previously been employed but were not included as a Workday worker.  Data includes personal details, contact information, and job details as of termination date.
func (c *Client) PutFormerWorker(ctx context.Context, input *PutFormerWorkerInput) (output *PutFormerWorkerOutput, err error) {
	output = &PutFormerWorkerOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutFormerWorkerInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Former_Worker_Request"`
	PutFormerWorkerRequestType
}

type PutFormerWorkerOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Former_Worker_Response"`
	PutFormerWorkerResponseType
}

// GetFormerWorkers (Get_Former_Workers)
// 
// This operation retrieves Former Workers.
//  Former Worker Storage is for tracking data for individuals that have previously been employed but were not included as a Workday worker.  Data includes personal details, contact information, and job details as of termination date.
func (c *Client) GetFormerWorkers(ctx context.Context, input *GetFormerWorkersInput) (output *GetFormerWorkersOutput, err error) {
	output = &GetFormerWorkersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetFormerWorkersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Former_Workers_Request"`
	GetFormerWorkersRequestType
}

type GetFormerWorkersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Former_Workers_Response"`
	GetFormerWorkersResponseType
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

// GetSocialBenefitsLocalities (Get_Social_Benefits_Localities)
// 
// This operation retrieves the social benefits localities that are currently defined.
func (c *Client) GetSocialBenefitsLocalities(ctx context.Context, input *GetSocialBenefitsLocalitiesInput) (output *GetSocialBenefitsLocalitiesOutput, err error) {
	output = &GetSocialBenefitsLocalitiesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetSocialBenefitsLocalitiesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Social_Benefits_Localities_Request"`
	GetSocialBenefitsLocalitiesRequestType
}

type GetSocialBenefitsLocalitiesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Social_Benefits_Localities_Response"`
	GetSocialBenefitsLocalitiesResponseType
}

// PutSocialBenefitsLocality (Put_Social_Benefits_Locality)
// 
// This operation adds or updates a Social Benefits Locality.
func (c *Client) PutSocialBenefitsLocality(ctx context.Context, input *PutSocialBenefitsLocalityInput) (output *PutSocialBenefitsLocalityOutput, err error) {
	output = &PutSocialBenefitsLocalityOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutSocialBenefitsLocalityInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Social_Benefits_Locality_Request"`
	PutSocialBenefitsLocalityRequestType
}

type PutSocialBenefitsLocalityOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Social_Benefits_Locality_Response"`
	PutSocialBenefitsLocalityResponseType
}

// PutFormerWorkerDocument (Put_Former_Worker_Document)
// 
// Adds or updates a former worker document. The operation adds documents not associated with events and those documents are not secured by the events.
func (c *Client) PutFormerWorkerDocument(ctx context.Context, input *PutFormerWorkerDocumentInput) (output *PutFormerWorkerDocumentOutput, err error) {
	output = &PutFormerWorkerDocumentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutFormerWorkerDocumentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Former_Worker_Document_Request"`
	PutFormerWorkerDocumentRequestType
}

type PutFormerWorkerDocumentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Former_Worker_Document_Response"`
	PutFormerWorkerDocumentResponseType
}

// GetFormerWorkerDocuments (Get_Former_Worker_Documents)
// 
// Returns former worker document data
func (c *Client) GetFormerWorkerDocuments(ctx context.Context, input *GetFormerWorkerDocumentsInput) (output *GetFormerWorkerDocumentsOutput, err error) {
	output = &GetFormerWorkerDocumentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetFormerWorkerDocumentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Former_Worker_Documents_Request"`
	GetFormerWorkerDocumentsRequestType
}

type GetFormerWorkerDocumentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Former_Worker_Documents_Response"`
	GetFormerWorkerDocumentsResponseType
}

// CreateNamedProfessorship (Create_Named_Professorship)
// 
// This operation will invoke the Business Process to create an Named Professorship.
func (c *Client) CreateNamedProfessorship(ctx context.Context, input *CreateNamedProfessorshipInput) (output *CreateNamedProfessorshipOutput, err error) {
	output = &CreateNamedProfessorshipOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CreateNamedProfessorshipInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Create_Named_Professorship_Request"`
	CreateNamedProfessorshipRequestType
}

type CreateNamedProfessorshipOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Create_Named_Professorship_Response"`
	CreateNamedProfessorshipResponseType
}

// EditNamedProfessorship (Edit_Named_Professorship)
// 
// Put Task for Edit named Professorship
func (c *Client) EditNamedProfessorship(ctx context.Context, input *EditNamedProfessorshipInput) (output *EditNamedProfessorshipOutput, err error) {
	output = &EditNamedProfessorshipOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EditNamedProfessorshipInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Named_Professorship_Request"`
	EditNamedProfessorshipRequestType
}

type EditNamedProfessorshipOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Named_Professorship_Response"`
	EditNamedProfessorshipResponseType
}

// PutAssignUserBasedSecurityGroup (Put_Assign_User-Based_Security_Group)
// 
// This operation performs a FULL REPLACEMENT of User-Based Security Groups. It does not add to existing entries. Assigns 1 or more Workday Accounts to a User-Based Security Group.
func (c *Client) PutAssignUserBasedSecurityGroup(ctx context.Context, input *PutAssignUserBasedSecurityGroupInput) (output *PutAssignUserBasedSecurityGroupOutput, err error) {
	output = &PutAssignUserBasedSecurityGroupOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutAssignUserBasedSecurityGroupInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Assign_User-Based_Security_Group_Request"`
	PutAssignUserBasedSecurityGroupRequestType
}

type PutAssignUserBasedSecurityGroupOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Assign_User-Based_Security_Group_Response"`
	PutAssignUserBasedSecurityGroupResponseType
}

// GetAssignUserBasedSecurityGroup (Get_Assign_User-Based_Security_Group)
// 
// Retrieves User-Based Security Group and Workday Accounts assigned to the security group.
func (c *Client) GetAssignUserBasedSecurityGroup(ctx context.Context, input *GetAssignUserBasedSecurityGroupInput) (output *GetAssignUserBasedSecurityGroupOutput, err error) {
	output = &GetAssignUserBasedSecurityGroupOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAssignUserBasedSecurityGroupInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Assign_User-Based_Security_Group_Request"`
	GetAssignUserBasedSecurityGroupRequestType
}

type GetAssignUserBasedSecurityGroupOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Assign_User-Based_Security_Group_Response"`
	GetAssignUserBasedSecurityGroupResponseType
}

// AssignEmployeeCollectiveAgreementEvent (Assign_Employee_Collective_Agreement_Event)
// 
// Creates or Updates Collective Agreement for a Worker. The Collective Agreement includes Applicable Factors and Factor options.
func (c *Client) AssignEmployeeCollectiveAgreementEvent(ctx context.Context, input *AssignEmployeeCollectiveAgreementEventInput) (output *AssignEmployeeCollectiveAgreementEventOutput, err error) {
	output = &AssignEmployeeCollectiveAgreementEventOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AssignEmployeeCollectiveAgreementEventInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Employee_Collective_Agreement_Event_Request"`
	AssignEmployeeCollectiveAgreementEventRequestType
}

type AssignEmployeeCollectiveAgreementEventOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Employee_Collective_Agreement_Event_Response"`
	AssignEmployeeCollectiveAgreementEventResponseType
}

// ChangeEmergencyContacts (Change_Emergency_Contacts)
// 
// This web service allows the updating of emergency contacts for a person.
func (c *Client) ChangeEmergencyContacts(ctx context.Context, input *ChangeEmergencyContactsInput) (output *ChangeEmergencyContactsOutput, err error) {
	output = &ChangeEmergencyContactsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeEmergencyContactsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Emergency_Contacts_Request"`
	ChangeEmergencyContactsRequestType
}

type ChangeEmergencyContactsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Emergency_Contacts_Response"`
	ChangeEmergencyContactsResponseType
}

// ReassignSuperiortoInactiveOrganization (Reassign_Superior_to_Inactive_Organization)
// 
// Provides ability to assign the previous superior to an organization that is currently inactive and does not currently have a superior.
func (c *Client) ReassignSuperiortoInactiveOrganization(ctx context.Context, input *ReassignSuperiortoInactiveOrganizationInput) (output *ReassignSuperiortoInactiveOrganizationOutput, err error) {
	output = &ReassignSuperiortoInactiveOrganizationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ReassignSuperiortoInactiveOrganizationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Reassign_Superior_to_Inactive_Organization_Request"`
	ReassignSuperiortoInactiveOrganizationRequestType
}

type ReassignSuperiortoInactiveOrganizationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Reassign_Superior_to_Inactive_Organization_Response"`
	ReassignSuperiortoInactiveOrganizationResponseType
}

// PutLocationHierarchyOrganizationAssignments (Put_Location_Hierarchy_Organization_Assignments)
// 
// This operation adds, updates, replaces the Organization Assignment Allowed Values for a Location Hierarchy
func (c *Client) PutLocationHierarchyOrganizationAssignments(ctx context.Context, input *PutLocationHierarchyOrganizationAssignmentsInput) (output *PutLocationHierarchyOrganizationAssignmentsOutput, err error) {
	output = &PutLocationHierarchyOrganizationAssignmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutLocationHierarchyOrganizationAssignmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Location_Hierarchy_Organization_Assignments_Request"`
	PutLocationHierarchyOrganizationAssignmentsRequestType
}

type PutLocationHierarchyOrganizationAssignmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Location_Hierarchy_Organization_Assignments_Response"`
	PutLocationHierarchyOrganizationAssignmentsResponseType
}

// GetLocationHierarchyOrganizationAssignments (Get_Location_Hierarchy_Organization_Assignments)
// 
// Returns Organization Assignments for Location Hierarchies.
func (c *Client) GetLocationHierarchyOrganizationAssignments(ctx context.Context, input *GetLocationHierarchyOrganizationAssignmentsInput) (output *GetLocationHierarchyOrganizationAssignmentsOutput, err error) {
	output = &GetLocationHierarchyOrganizationAssignmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLocationHierarchyOrganizationAssignmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Location_Hierarchy_Organization_Assignments_Request"`
	GetLocationHierarchyOrganizationAssignmentsRequestType
}

type GetLocationHierarchyOrganizationAssignmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Location_Hierarchy_Organization_Assignments_Response"`
	GetLocationHierarchyOrganizationAssignmentsResponseType
}

// GetOrganizationReferenceIDs (Get_Organization_Reference_IDs)
// 
// This task can be used to retrieve Organization Reference ID for an Organization.
func (c *Client) GetOrganizationReferenceIDs(ctx context.Context, input *GetOrganizationReferenceIDsInput) (output *GetOrganizationReferenceIDsOutput, err error) {
	output = &GetOrganizationReferenceIDsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetOrganizationReferenceIDsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Organization_Reference_IDs_Request"`
	GetOrganizationReferenceIDsRequestType
}

type GetOrganizationReferenceIDsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Organization_Reference_IDs_Response"`
	GetOrganizationReferenceIDsResponseType
}

// PutOrganizationReferenceID (Put_Organization_Reference_ID)
// 
// This task allows changing Organization Reference ID for an Organization. For Supervisory Organizations
// - if a valid sequence generator is defined, and the Organization Reference ID field is set to empty, then ID will be replaced with an automatically generated ID
// - if no sequence generator is defined, and the Organization Reference ID field is set to empty, then ID will be generated follow default pattern (SUPERVISORY_ORGANIZATION_XX_YY)
// - an option &#39;Include Organization ID in Name&#39; can be set.
// A new effective-dated organization name is automatically created if any attributes used for organization name are different as of specified effective date.
func (c *Client) PutOrganizationReferenceID(ctx context.Context, input *PutOrganizationReferenceIDInput) (output *PutOrganizationReferenceIDOutput, err error) {
	output = &PutOrganizationReferenceIDOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutOrganizationReferenceIDInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Organization_Reference_ID_Request"`
	PutOrganizationReferenceIDRequestType
}

type PutOrganizationReferenceIDOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Organization_Reference_ID_Response"`
	PutOrganizationReferenceIDResponseType
}

// PutCommitteeType (Put_Committee_Type)
// 
// This operation adds or updates Committee Types
func (c *Client) PutCommitteeType(ctx context.Context, input *PutCommitteeTypeInput) (output *PutCommitteeTypeOutput, err error) {
	output = &PutCommitteeTypeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCommitteeTypeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Committee_Type_Request"`
	PutCommitteeTypeRequestType
}

type PutCommitteeTypeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Committee_Type_Response"`
	PutCommitteeTypeResponseType
}

// GetCommitteeTypes (Get_Committee_Types)
// 
// This operation retrieves Committee Types
func (c *Client) GetCommitteeTypes(ctx context.Context, input *GetCommitteeTypesInput) (output *GetCommitteeTypesOutput, err error) {
	output = &GetCommitteeTypesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCommitteeTypesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Committee_Types_Request"`
	GetCommitteeTypesRequestType
}

type GetCommitteeTypesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Committee_Types_Response"`
	GetCommitteeTypesResponseType
}

// PutCommitteeClassificationGroup (Put_Committee_Classification_Group)
// 
// This operation adds or updates Committee Classification Groups
func (c *Client) PutCommitteeClassificationGroup(ctx context.Context, input *PutCommitteeClassificationGroupInput) (output *PutCommitteeClassificationGroupOutput, err error) {
	output = &PutCommitteeClassificationGroupOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCommitteeClassificationGroupInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Committee_Classification_Group_Request"`
	PutCommitteeClassificationGroupRequestType
}

type PutCommitteeClassificationGroupOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Committee_Classification_Group_Response"`
	PutCommitteeClassificationGroupResponseType
}

// GetCommitteeClassificationGroups (Get_Committee_Classification_Groups)
// 
// This operation retrieves Committee Classification Groups
func (c *Client) GetCommitteeClassificationGroups(ctx context.Context, input *GetCommitteeClassificationGroupsInput) (output *GetCommitteeClassificationGroupsOutput, err error) {
	output = &GetCommitteeClassificationGroupsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCommitteeClassificationGroupsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Committee_Classification_Groups_Request"`
	GetCommitteeClassificationGroupsRequestType
}

type GetCommitteeClassificationGroupsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Committee_Classification_Groups_Response"`
	GetCommitteeClassificationGroupsResponseType
}

// PutCommitteeClassification (Put_Committee_Classification)
// 
// This operation adds or updates Committee Classification
func (c *Client) PutCommitteeClassification(ctx context.Context, input *PutCommitteeClassificationInput) (output *PutCommitteeClassificationOutput, err error) {
	output = &PutCommitteeClassificationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCommitteeClassificationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Committee_Classification_Request"`
	PutCommitteeClassificationRequestType
}

type PutCommitteeClassificationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Committee_Classification_Response"`
	PutCommitteeClassificationResponseType
}

// GetCommitteeClassifications (Get_Committee_Classifications)
// 
// This operation retrieves Committee Classifications
func (c *Client) GetCommitteeClassifications(ctx context.Context, input *GetCommitteeClassificationsInput) (output *GetCommitteeClassificationsOutput, err error) {
	output = &GetCommitteeClassificationsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCommitteeClassificationsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Committee_Classifications_Request"`
	GetCommitteeClassificationsRequestType
}

type GetCommitteeClassificationsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Committee_Classifications_Response"`
	GetCommitteeClassificationsResponseType
}

// GetCommitteeMembershipTypes (Get_Committee_Membership_Types)
// 
// This operation retrieves Committee Member Types.
func (c *Client) GetCommitteeMembershipTypes(ctx context.Context, input *GetCommitteeMembershipTypesInput) (output *GetCommitteeMembershipTypesOutput, err error) {
	output = &GetCommitteeMembershipTypesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCommitteeMembershipTypesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Committee_Membership_Types_Request"`
	GetCommitteeMembershipTypesRequestType
}

type GetCommitteeMembershipTypesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Committee_Membership_Types_Response"`
	GetCommitteeMembershipTypesResponseType
}

// PutCommitteeMembershipType (Put_Committee_Membership_Type)
// 
// This operation adds or updates Committee Member Type.
func (c *Client) PutCommitteeMembershipType(ctx context.Context, input *PutCommitteeMembershipTypeInput) (output *PutCommitteeMembershipTypeOutput, err error) {
	output = &PutCommitteeMembershipTypeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCommitteeMembershipTypeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Committee_Membership_Type_Request"`
	PutCommitteeMembershipTypeRequestType
}

type PutCommitteeMembershipTypeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Committee_Membership_Type_Response"`
	PutCommitteeMembershipTypeResponseType
}

// ManageCommitteeMembership (Manage_Committee_Membership)
// 
// This operation invokes the Manage Committee Membership business process.
func (c *Client) ManageCommitteeMembership(ctx context.Context, input *ManageCommitteeMembershipInput) (output *ManageCommitteeMembershipOutput, err error) {
	output = &ManageCommitteeMembershipOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ManageCommitteeMembershipInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Committee_Membership_Request"`
	ManageCommitteeMembershipRequestType
}

type ManageCommitteeMembershipOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Committee_Membership_Response"`
	ManageCommitteeMembershipResponseType
}

// ManageEmployeeProbationPeriodsEvent (Manage_Employee_Probation_Periods_Event)
// 
// Creates or Updates Probation Periods for a Worker.
func (c *Client) ManageEmployeeProbationPeriodsEvent(ctx context.Context, input *ManageEmployeeProbationPeriodsEventInput) (output *ManageEmployeeProbationPeriodsEventOutput, err error) {
	output = &ManageEmployeeProbationPeriodsEventOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ManageEmployeeProbationPeriodsEventInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Employee_Probation_Periods_Event_Request"`
	ManageEmployeeProbationPeriodsEventRequestType
}

type ManageEmployeeProbationPeriodsEventOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Employee_Probation_Periods_Event_Response"`
	ManageEmployeeProbationPeriodsEventResponseType
}

// MaintainCommitteeDefinition (Maintain_Committee_Definition)
// 
// This operation adds or updates a Committee Definition
func (c *Client) MaintainCommitteeDefinition(ctx context.Context, input *MaintainCommitteeDefinitionInput) (output *MaintainCommitteeDefinitionOutput, err error) {
	output = &MaintainCommitteeDefinitionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type MaintainCommitteeDefinitionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Maintain_Committee_Definition_Request"`
	MaintainCommitteeDefinitionRequestType
}

type MaintainCommitteeDefinitionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Maintain_Committee_Definition_Response"`
	MaintainCommitteeDefinitionResponseType
}

// GetCommitteeDefinition (Get_Committee_Definition)
// 
// This operation gets the current committee(s) definition
func (c *Client) GetCommitteeDefinition(ctx context.Context, input *GetCommitteeDefinitionInput) (output *GetCommitteeDefinitionOutput, err error) {
	output = &GetCommitteeDefinitionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCommitteeDefinitionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Committee_Definition_Request"`
	GetCommitteeDefinitionRequestType
}

type GetCommitteeDefinitionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Committee_Definition_Response"`
	GetCommitteeDefinitionResponseType
}

// GetLGBTIdentifications (Get_LGBT_Identifications)
// 
// This operation retrieves the LGBT Identifications that are currently defined.
func (c *Client) GetLGBTIdentifications(ctx context.Context, input *GetLGBTIdentificationsInput) (output *GetLGBTIdentificationsOutput, err error) {
	output = &GetLGBTIdentificationsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLGBTIdentificationsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_LGBT_Identifications_Request"`
	GetLGBTIdentificationsRequestType
}

type GetLGBTIdentificationsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_LGBT_Identifications_Response"`
	GetLGBTIdentificationsResponseType
}

// PutLGBTIdentification (Put_LGBT_Identification)
// 
// This operation adds or updates a LGBT Identification
func (c *Client) PutLGBTIdentification(ctx context.Context, input *PutLGBTIdentificationInput) (output *PutLGBTIdentificationOutput, err error) {
	output = &PutLGBTIdentificationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutLGBTIdentificationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_LGBT_Identification_Request"`
	PutLGBTIdentificationRequestType
}

type PutLGBTIdentificationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_LGBT_Identification_Response"`
	PutLGBTIdentificationResponseType
}

// PutAppointmentSpecialty (Put_Appointment_Specialty)
// 
// This operation adds or updates Appointment Specialty.
func (c *Client) PutAppointmentSpecialty(ctx context.Context, input *PutAppointmentSpecialtyInput) (output *PutAppointmentSpecialtyOutput, err error) {
	output = &PutAppointmentSpecialtyOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutAppointmentSpecialtyInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Appointment_Specialty_Request"`
	PutAppointmentSpecialtyRequestType
}

type PutAppointmentSpecialtyOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Appointment_Specialty_Response"`
	PutAppointmentSpecialtyResponseType
}

// GetAppointmentSpecialties (Get_Appointment_Specialties)
// 
// This operation retrieves Appointment Specialties.
func (c *Client) GetAppointmentSpecialties(ctx context.Context, input *GetAppointmentSpecialtiesInput) (output *GetAppointmentSpecialtiesOutput, err error) {
	output = &GetAppointmentSpecialtiesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAppointmentSpecialtiesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Appointment_Specialties_Request"`
	GetAppointmentSpecialtiesRequestType
}

type GetAppointmentSpecialtiesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Appointment_Specialties_Response"`
	GetAppointmentSpecialtiesResponseType
}

// EndCollectiveAgreementAssignment (End_Collective_Agreement_Assignment)
// 
// Ends the Collective Agreement Assignment for a ~Worker~ as of the effective data.
func (c *Client) EndCollectiveAgreementAssignment(ctx context.Context, input *EndCollectiveAgreementAssignmentInput) (output *EndCollectiveAgreementAssignmentOutput, err error) {
	output = &EndCollectiveAgreementAssignmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EndCollectiveAgreementAssignmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc End_Collective_Agreement_Assignment_Event_Request"`
	EndCollectiveAgreementAssignmentEventRequestType
}

type EndCollectiveAgreementAssignmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc End_Collective_Agreement_Assignment_Event_Response"`
	EndCollectiveAgreementAssignmentEventResponseType
}

// GetCommitteeMeetings (Get_Committee_Meetings)
// 
// This operation retrieves data about existing committee meetings.
func (c *Client) GetCommitteeMeetings(ctx context.Context, input *GetCommitteeMeetingsInput) (output *GetCommitteeMeetingsOutput, err error) {
	output = &GetCommitteeMeetingsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCommitteeMeetingsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Committee_Meetings_Request"`
	GetCommitteeMeetingsRequestType
}

type GetCommitteeMeetingsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Committee_Meetings_Response"`
	GetCommitteeMeetingsResponseType
}

// PutCommitteeMeeting (Put_Committee_Meeting)
// 
// This operation adds or updates a new Committee Meeting.
func (c *Client) PutCommitteeMeeting(ctx context.Context, input *PutCommitteeMeetingInput) (output *PutCommitteeMeetingOutput, err error) {
	output = &PutCommitteeMeetingOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCommitteeMeetingInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Committee_Meeting_Request"`
	PutCommitteeMeetingRequestType
}

type PutCommitteeMeetingOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Committee_Meeting_Response"`
	PutCommitteeMeetingResponseType
}

// GetEstablishments (Get_Establishments)
// 
// Get Establishments Data.
func (c *Client) GetEstablishments(ctx context.Context, input *GetEstablishmentsInput) (output *GetEstablishmentsOutput, err error) {
	output = &GetEstablishmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEstablishmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Establishments_Request"`
	GetEstablishmentsRequestType
}

type GetEstablishmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Establishments_Response"`
	GetEstablishmentsResponseType
}

// PutEstablishment (Put_Establishment)
// 
// Loads Establishment Data.
func (c *Client) PutEstablishment(ctx context.Context, input *PutEstablishmentInput) (output *PutEstablishmentOutput, err error) {
	output = &PutEstablishmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutEstablishmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Establishment_Request"`
	PutEstablishmentRequestType
}

type PutEstablishmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Establishment_Response"`
	PutEstablishmentResponseType
}

// AssignEstablishment (Assign_Establishment)
// 
// Manually assign establishments to workers&#39; positions
func (c *Client) AssignEstablishment(ctx context.Context, input *AssignEstablishmentInput) (output *AssignEstablishmentOutput, err error) {
	output = &AssignEstablishmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AssignEstablishmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Establishment_Request"`
	AssignEstablishmentRequestType
}

type AssignEstablishmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Establishment_Response"`
	AssignEstablishmentResponseType
}

// GetWorkdayAccount (Get_Workday_Account)
// 
// Get Workday Accounts for a user or set of users
func (c *Client) GetWorkdayAccount(ctx context.Context, input *GetWorkdayAccountInput) (output *GetWorkdayAccountOutput, err error) {
	output = &GetWorkdayAccountOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetWorkdayAccountInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Workday_Account_Request"`
	GetWorkdayAccountRequestType
}

type GetWorkdayAccountOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Workday_Account_Response"`
	GetWorkdayAccountResponseType
}

// ReactivateOrganization (Reactivate_Organization)
// 
// This operation reactivates an organization based on the specified reference ID or WID. The organization must already be inactive or be inactive in the future. The system automatically sets the reactivation effective date to the same date as the date of the last inactivation so that there are no gaps when an active organization was inactive. This web service can currently be used to reactivate these organization types: Supervisory, Company, Cost Center, Region, Matrix, Pay Group, Retiree, Custom Org, Company Hierarchy, Cost Center Hierarchy, Location Hierarchy, Region Hierarchy.
func (c *Client) ReactivateOrganization(ctx context.Context, input *ReactivateOrganizationInput) (output *ReactivateOrganizationOutput, err error) {
	output = &ReactivateOrganizationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ReactivateOrganizationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Reactivate_Organization_Request"`
	ReactivateOrganizationRequestType
}

type ReactivateOrganizationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Reactivate_Organization_Response"`
	ReactivateOrganizationResponseType
}

// ChangeVeteranStatusIdentification (Change_Veteran_Status_Identification)
// 
// This web service assigns Veteran Status Identifications for an Employee.
func (c *Client) ChangeVeteranStatusIdentification(ctx context.Context, input *ChangeVeteranStatusIdentificationInput) (output *ChangeVeteranStatusIdentificationOutput, err error) {
	output = &ChangeVeteranStatusIdentificationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeVeteranStatusIdentificationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Veteran_Status_Identification_Request"`
	ChangeVeteranStatusIdentificationRequestType
}

type ChangeVeteranStatusIdentificationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Veteran_Status_Identification_Response"`
	ChangeVeteranStatusIdentificationResponseType
}

// AssignMemberstoCustomOrganization (Assign_Members_to_Custom_Organization)
// 
// Directly assign or unassign worker positions as members to a custom organization. The custom organization type must be configured in &#39;Maintain Organization Types&#39; as &#39;Allow Reorganization Tasks (like Move Workers, Assign Workers)&#39; = Y and &#39;Position Unique&#39; = N. The Reorganization determines the effective date that a position is added as a member. In the UI, the equivalent task is executed via related action off the custom organization: Reorganization &gt; Assign Workers &gt; Select Workers by Individual.
func (c *Client) AssignMemberstoCustomOrganization(ctx context.Context, input *AssignMemberstoCustomOrganizationInput) (output *AssignMemberstoCustomOrganizationOutput, err error) {
	output = &AssignMemberstoCustomOrganizationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AssignMemberstoCustomOrganizationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Members_to_Custom_Organization_Request"`
	AssignMemberstoCustomOrganizationRequestType
}

type AssignMemberstoCustomOrganizationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Assign_Members_to_Custom_Organization_Response"`
	AssignMemberstoCustomOrganizationResponseType
}

// PutWorkStatusRuleSet (Put_Work_Status_Rule_Set)
// 
// Create/Edit a Work Status Rule Set via Web Services
func (c *Client) PutWorkStatusRuleSet(ctx context.Context, input *PutWorkStatusRuleSetInput) (output *PutWorkStatusRuleSetOutput, err error) {
	output = &PutWorkStatusRuleSetOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutWorkStatusRuleSetInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Work_Status_Rule_Set_Request"`
	PutWorkStatusRuleSetRequestType
}

type PutWorkStatusRuleSetOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Work_Status_Rule_Set_Response"`
	PutWorkStatusRuleSetResponseType
}

// GetWorkStatusRuleSets (Get_Work_Status_Rule_Sets)
// 
// Retrieves existing Work Status Rule Sets via Web Service
func (c *Client) GetWorkStatusRuleSets(ctx context.Context, input *GetWorkStatusRuleSetsInput) (output *GetWorkStatusRuleSetsOutput, err error) {
	output = &GetWorkStatusRuleSetsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetWorkStatusRuleSetsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Work_Status_Rule_Sets_Request"`
	GetWorkStatusRuleSetsRequestType
}

type GetWorkStatusRuleSetsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Work_Status_Rule_Sets_Response"`
	GetWorkStatusRuleSetsResponseType
}

// PutPeriodReportingCalendar (Put_Period_Reporting_Calendar)
// 
// This operation adds or updates a Period Reporting Calendar
func (c *Client) PutPeriodReportingCalendar(ctx context.Context, input *PutPeriodReportingCalendarInput) (output *PutPeriodReportingCalendarOutput, err error) {
	output = &PutPeriodReportingCalendarOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPeriodReportingCalendarInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Period_Reporting_Calendar_Request"`
	PutPeriodReportingCalendarRequestType
}

type PutPeriodReportingCalendarOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Period_Reporting_Calendar_Response"`
	PutPeriodReportingCalendarResponseType
}

// GetPeriodReportingCalendars (Get_Period_Reporting_Calendars)
// 
// Retrieves the information about the Period Reporting Calendars requested
func (c *Client) GetPeriodReportingCalendars(ctx context.Context, input *GetPeriodReportingCalendarsInput) (output *GetPeriodReportingCalendarsOutput, err error) {
	output = &GetPeriodReportingCalendarsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPeriodReportingCalendarsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Period_Reporting_Calendars_Request"`
	GetPeriodReportingCalendarsRequestType
}

type GetPeriodReportingCalendarsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Period_Reporting_Calendars_Response"`
	GetPeriodReportingCalendarsResponseType
}

// PutLocationAttribute (Put_Location_Attribute)
// 
// Public Web Service Operation for Adding/Editing Location Attributes
func (c *Client) PutLocationAttribute(ctx context.Context, input *PutLocationAttributeInput) (output *PutLocationAttributeOutput, err error) {
	output = &PutLocationAttributeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutLocationAttributeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Location_Attribute_Request"`
	PutLocationAttributeRequestType
}

type PutLocationAttributeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Location_Attribute_Response"`
	PutLocationAttributeResponseType
}

// GetLocationAttributes (Get_Location_Attributes)
// 
// Public Web Service Operation for Retrieving Location Attributes
func (c *Client) GetLocationAttributes(ctx context.Context, input *GetLocationAttributesInput) (output *GetLocationAttributesOutput, err error) {
	output = &GetLocationAttributesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLocationAttributesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Location_Attributes_Request"`
	GetLocationAttributesRequestType
}

type GetLocationAttributesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Location_Attributes_Response"`
	GetLocationAttributesResponseType
}

// GetSafetyIncidentLocations (Get_Safety_Incident_Locations)
// 
// This web service allows a Safety Incident Location to be retrieved from Workday.
func (c *Client) GetSafetyIncidentLocations(ctx context.Context, input *GetSafetyIncidentLocationsInput) (output *GetSafetyIncidentLocationsOutput, err error) {
	output = &GetSafetyIncidentLocationsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetSafetyIncidentLocationsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Safety_Incident_Locations_Request"`
	GetSafetyIncidentLocationsRequestType
}

type GetSafetyIncidentLocationsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Safety_Incident_Locations_Response"`
	GetSafetyIncidentLocationsResponseType
}

// PutSafetyIncidentLocation (Put_Safety_Incident_Location)
// 
// This web service allows a Safety Incident Location to be added into Workday.
func (c *Client) PutSafetyIncidentLocation(ctx context.Context, input *PutSafetyIncidentLocationInput) (output *PutSafetyIncidentLocationOutput, err error) {
	output = &PutSafetyIncidentLocationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutSafetyIncidentLocationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Safety_Incident_Location_Request"`
	PutSafetyIncidentLocationRequestType
}

type PutSafetyIncidentLocationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Safety_Incident_Location_Response"`
	PutSafetyIncidentLocationResponseType
}

// PutExternalDisabilitySelfIdentificationRecord (Put_External_Disability_Self_Identification_Record)
// 
// Web service task to add, delete, or edit external disability self-identification records
func (c *Client) PutExternalDisabilitySelfIdentificationRecord(ctx context.Context, input *PutExternalDisabilitySelfIdentificationRecordInput) (output *PutExternalDisabilitySelfIdentificationRecordOutput, err error) {
	output = &PutExternalDisabilitySelfIdentificationRecordOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutExternalDisabilitySelfIdentificationRecordInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_External_Disability_Self_Identification_Record_Request"`
	PutExternalDisabilitySelfIdentificationRecordRequestType
}

type PutExternalDisabilitySelfIdentificationRecordOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_External_Disability_Self_Identification_Record_Response"`
	PutExternalDisabilitySelfIdentificationRecordResponseType
}

// GetExternalDisabilitySelfIdentificationRecords (Get_External_Disability_Self_Identification_Records)
// 
// This operation retrieves the External Disability Self-Identification Records that are currently defined
func (c *Client) GetExternalDisabilitySelfIdentificationRecords(ctx context.Context, input *GetExternalDisabilitySelfIdentificationRecordsInput) (output *GetExternalDisabilitySelfIdentificationRecordsOutput, err error) {
	output = &GetExternalDisabilitySelfIdentificationRecordsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetExternalDisabilitySelfIdentificationRecordsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_External_Disability_Self_Identification_Records_Request"`
	GetExternalDisabilitySelfIdentificationRecordsRequestType
}

type GetExternalDisabilitySelfIdentificationRecordsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_External_Disability_Self_Identification_Records_Response"`
	GetExternalDisabilitySelfIdentificationRecordsResponseType
}

// GetGenderIdentities (Get_Gender_Identities)
// 
// Returns Gender Identities by Reference ID, or returns all Gender Identities if no Reference ID is provided.
func (c *Client) GetGenderIdentities(ctx context.Context, input *GetGenderIdentitiesInput) (output *GetGenderIdentitiesOutput, err error) {
	output = &GetGenderIdentitiesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetGenderIdentitiesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Gender_Identities_Request"`
	GetGenderIdentitiesRequestType
}

type GetGenderIdentitiesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Gender_Identities_Response"`
	GetGenderIdentitiesResponseType
}

// PutGenderIdentity (Put_Gender_Identity)
// 
// Creates a new Gender Identity (or updates an existing Gender Identity) with the information supplied in the request.
func (c *Client) PutGenderIdentity(ctx context.Context, input *PutGenderIdentityInput) (output *PutGenderIdentityOutput, err error) {
	output = &PutGenderIdentityOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutGenderIdentityInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Gender_Identity_Request"`
	PutGenderIdentityRequestType
}

type PutGenderIdentityOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Gender_Identity_Response"`
	PutGenderIdentityResponseType
}

// GetSexualOrientations (Get_Sexual_Orientations)
// 
// Returns Sexual Orientations by Reference ID, or returns all Sexual Orientations if no Reference ID is provided.
func (c *Client) GetSexualOrientations(ctx context.Context, input *GetSexualOrientationsInput) (output *GetSexualOrientationsOutput, err error) {
	output = &GetSexualOrientationsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetSexualOrientationsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Sexual_Orientations_Request"`
	GetSexualOrientationsRequestType
}

type GetSexualOrientationsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Sexual_Orientations_Response"`
	GetSexualOrientationsResponseType
}

// PutSexualOrientation (Put_Sexual_Orientation)
// 
// Creates a new Sexual Orientation (or updates an existing Sexual Orientation) with the information supplied in the request.
func (c *Client) PutSexualOrientation(ctx context.Context, input *PutSexualOrientationInput) (output *PutSexualOrientationOutput, err error) {
	output = &PutSexualOrientationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutSexualOrientationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Sexual_Orientation_Request"`
	PutSexualOrientationRequestType
}

type PutSexualOrientationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Sexual_Orientation_Response"`
	PutSexualOrientationResponseType
}

// GetPronouns (Get_Pronouns)
// 
// Returns Pronouns by Reference ID, or returns all Pronouns if no Reference ID is provided.
func (c *Client) GetPronouns(ctx context.Context, input *GetPronounsInput) (output *GetPronounsOutput, err error) {
	output = &GetPronounsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPronounsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Pronouns_Request"`
	GetPronounsRequestType
}

type GetPronounsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Pronouns_Response"`
	GetPronounsResponseType
}

// PutPronoun (Put_Pronoun)
// 
// Creates a new Pronoun (or updates an existing Pronoun) with the information supplies in the request.
func (c *Client) PutPronoun(ctx context.Context, input *PutPronounInput) (output *PutPronounOutput, err error) {
	output = &PutPronounOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPronounInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Pronoun_Request"`
	PutPronounRequestType
}

type PutPronounOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Pronoun_Response"`
	PutPronounResponseType
}

// GetOrganizationsNames (Get_Organizations_Names)
// 
// Retrieves data related to the organization name or names in effect for a specified date range. If both From and To dates are empty, Workday retrieves data related to the organization name as of the current date. Retrieves all organizations unless Organization or Organization Type is specified.
func (c *Client) GetOrganizationsNames(ctx context.Context, input *GetOrganizationsNamesInput) (output *GetOrganizationsNamesOutput, err error) {
	output = &GetOrganizationsNamesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetOrganizationsNamesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Organizations_Names_Request"`
	GetOrganizationsNamesRequestType
}

type GetOrganizationsNamesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Organizations_Names_Response"`
	GetOrganizationsNamesResponseType
}

// PutOrganizationName (Put_Organization_Name)
// 
// Updates data related to an organization name as of the effective date. If no effective date is provided, Workday uses the current date.
func (c *Client) PutOrganizationName(ctx context.Context, input *PutOrganizationNameInput) (output *PutOrganizationNameOutput, err error) {
	output = &PutOrganizationNameOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutOrganizationNameInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Organization_Name_Request"`
	PutOrganizationNameRequestType
}

type PutOrganizationNameOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Organization_Name_Response"`
	PutOrganizationNameResponseType
}

// CreateReferenceLetterforWorker (Create_Reference_Letter_for_Worker)
// 
// Initiates the creation of a Reference Letter for a Worker. Uses the Reference Letter Request business process.
func (c *Client) CreateReferenceLetterforWorker(ctx context.Context, input *CreateReferenceLetterforWorkerInput) (output *CreateReferenceLetterforWorkerOutput, err error) {
	output = &CreateReferenceLetterforWorkerOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CreateReferenceLetterforWorkerInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Create_Reference_Letter_for_Worker_Request"`
	CreateReferenceLetterforWorkerRequestType
}

type CreateReferenceLetterforWorkerOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Create_Reference_Letter_for_Worker_Response"`
	CreateReferenceLetterforWorkerResponseType
}

// GetUniversalIdentifier (Get_Universal_Identifier)
// 
// Request for Universal Identifiers. Use a Person Reference to get the Universal ID for the Person or don&#39;t request anybody and get all Universal IDs.
func (c *Client) GetUniversalIdentifier(ctx context.Context, input *GetUniversalIdentifierInput) (output *GetUniversalIdentifierOutput, err error) {
	output = &GetUniversalIdentifierOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetUniversalIdentifierInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Universal_Identifiers_Request"`
	GetUniversalIdentifiersRequestType
}

type GetUniversalIdentifierOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Universal_Identifiers_Response"`
	GetUniversalIdentifiersResponseType
}

// PutUniversalIdentifier (Put_Universal_Identifier)
// 
// Put Request for Universal Identifiers. It&#39;s required to put a Person Reference in either Person Reference. Either place a manual Universal ID or no Universal ID, if no Universal ID is supplied it will use the Universal ID generator.
func (c *Client) PutUniversalIdentifier(ctx context.Context, input *PutUniversalIdentifierInput) (output *PutUniversalIdentifierOutput, err error) {
	output = &PutUniversalIdentifierOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutUniversalIdentifierInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Universal_Identifier_Request"`
	PutUniversalIdentifierRequestType
}

type PutUniversalIdentifierOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Universal_Identifier_Response"`
	PutUniversalIdentifierResponseType
}

// ChangeBusinessTitle (Change_Business_Title)
// 
// Sets a ~worker&#39;s~ personal business title. Uses the Change Business Title business process. If a ~worker~ has two jobs, one of the jobs must be specified for the ~worker~.
func (c *Client) ChangeBusinessTitle(ctx context.Context, input *ChangeBusinessTitleInput) (output *ChangeBusinessTitleOutput, err error) {
	output = &ChangeBusinessTitleOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeBusinessTitleInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Business_Title_Request"`
	ChangeBusinessTitleRequestType
}

type ChangeBusinessTitleOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Business_Title_Response"`
	ChangeBusinessTitleResponseType
}

// EditAcademicAppointmentTrackAdditionalData (Edit_Academic_Appointment_Track_Additional_Data)
// 
// This web service operation is best suited for the initial creation of the  effective-dated custom object data on the appointment track. Subsequent updates are best performed via the UI using the online Update Academic Appointment business process. If this operation is used for updates, it should be noted that it will require the full set of all the custom object data every time any sort of update is performed.
func (c *Client) EditAcademicAppointmentTrackAdditionalData(ctx context.Context, input *EditAcademicAppointmentTrackAdditionalDataInput) (output *EditAcademicAppointmentTrackAdditionalDataOutput, err error) {
	output = &EditAcademicAppointmentTrackAdditionalDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EditAcademicAppointmentTrackAdditionalDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Academic_Appointment_Track_Additional_Data_Request"`
	EditAcademicAppointmentTrackAdditionalDataRequestType
}

type EditAcademicAppointmentTrackAdditionalDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Academic_Appointment_Track_Additional_Data_Response"`
	EditAcademicAppointmentTrackAdditionalDataResponseType
}

// GetAcademicAppointmentTrackAdditionalData (Get_Academic_Appointment_Track_Additional_Data)
// 
// Retrieves effective-dated additional data for the specified Academic Appointment Tracks or for all Academic Appointment Tracks if none were specified. Use the effective response filter to specify the effective date, otherwise additional data as of the current system time will be retrieved.
func (c *Client) GetAcademicAppointmentTrackAdditionalData(ctx context.Context, input *GetAcademicAppointmentTrackAdditionalDataInput) (output *GetAcademicAppointmentTrackAdditionalDataOutput, err error) {
	output = &GetAcademicAppointmentTrackAdditionalDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAcademicAppointmentTrackAdditionalDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Academic_Appointment_Track_Additional_Data_Request"`
	GetAcademicAppointmentTrackAdditionalDataRequestType
}

type GetAcademicAppointmentTrackAdditionalDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Academic_Appointment_Track_Additional_Data_Response"`
	GetAcademicAppointmentTrackAdditionalDataResponseType
}

// GetAcademicAppointee (Get_Academic_Appointee)
// 
// Get get Academic Appointee information
func (c *Client) GetAcademicAppointee(ctx context.Context, input *GetAcademicAppointeeInput) (output *GetAcademicAppointeeOutput, err error) {
	output = &GetAcademicAppointeeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAcademicAppointeeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Academic_Appointee_Request"`
	GetAcademicAppointeeRequestType
}

type GetAcademicAppointeeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Academic_Appointee_Response"`
	GetAcademicAppointeeResponseType
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

// PutSortOrderEnabled (Put_Sort_Order_Enabled)
// 
// Web Service to create/maintain the order of orderable instances.
func (c *Client) PutSortOrderEnabled(ctx context.Context, input *PutSortOrderEnabledInput) (output *PutSortOrderEnabledOutput, err error) {
	output = &PutSortOrderEnabledOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutSortOrderEnabledInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Sort_Order_Enabled_Request"`
	PutSortOrderEnabledRequestType
}

type PutSortOrderEnabledOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Sort_Order_Enabled_Response"`
	PutSortOrderEnabledResponseType
}

// GetSortOrderEnableds (Get_Sort_Order_Enableds)
// 
// Web Service to get the order of orderable instances.
func (c *Client) GetSortOrderEnableds(ctx context.Context, input *GetSortOrderEnabledsInput) (output *GetSortOrderEnabledsOutput, err error) {
	output = &GetSortOrderEnabledsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetSortOrderEnabledsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Sort_Order_Enableds_Request"`
	GetSortOrderEnabledsRequestType
}

type GetSortOrderEnabledsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Sort_Order_Enableds_Response"`
	GetSortOrderEnabledsResponseType
}

// PutPreferredCommunicationLanguage (Put_Preferred_Communication_Language)
// 
// Put Preferred Communication Language configuration
func (c *Client) PutPreferredCommunicationLanguage(ctx context.Context, input *PutPreferredCommunicationLanguageInput) (output *PutPreferredCommunicationLanguageOutput, err error) {
	output = &PutPreferredCommunicationLanguageOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPreferredCommunicationLanguageInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Preferred_Communication_Language_Request"`
	PutPreferredCommunicationLanguageRequestType
}

type PutPreferredCommunicationLanguageOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Preferred_Communication_Language_Response"`
	PutPreferredCommunicationLanguageResponseType
}

// GetPreferredCommunicationLanguage (Get_Preferred_Communication_Language)
// 
// Get Preferred Communication language configuration
func (c *Client) GetPreferredCommunicationLanguage(ctx context.Context, input *GetPreferredCommunicationLanguageInput) (output *GetPreferredCommunicationLanguageOutput, err error) {
	output = &GetPreferredCommunicationLanguageOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPreferredCommunicationLanguageInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Preferred_Communication_Language_Request"`
	GetPreferredCommunicationLanguageRequestType
}

type GetPreferredCommunicationLanguageOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Preferred_Communication_Language_Response"`
	GetPreferredCommunicationLanguageResponseType
}

// GetWorkteams (Get_Workteams)
// 
// Returns Team information. If the request does not specify a Team, the operation returns information for all Teams.
func (c *Client) GetWorkteams(ctx context.Context, input *GetWorkteamsInput) (output *GetWorkteamsOutput, err error) {
	output = &GetWorkteamsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetWorkteamsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Workteams_Request"`
	GetWorkteamsRequestType
}

type GetWorkteamsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Workteams_Response"`
	GetWorkteamsResponseType
}

// PutWorkteam (Put_Workteam)
// 
// This operation adds or deletes workteam members, or updates workteam members&#39; membership status. Each membership status change is effective dated in order to capture membership status history. Each effective-dated membership status change is referred to as a snapshot.
// 
// When a workteam member is deleted, it is as if they were never a member and their entire membership status history is also deleted. If you want to retain membership status history after a member departs, Workday recommends that you change the member&#39;s status to Inactive instead.
// 
// When a worker does not exist as a workteam member as of the specified effective date, they are added as a member with the specified effective date and membership status. When a worker does exist as a workteam member as of the specified effective date, the operation updates their membership status. You can correct a membership status by using the same effective date as an existing membership status snapshot or correct the membership status to an earlier effective date by specifying the desired effective date and the same membership status as the next effective-dated snapshot.
// 
// Example: Membership status is Passive with effective date December 1, 2018. If you specify November 15, 2018 as the effective date and membership status as Passive, the system will change the effective date on the December 1, 2018 snapshot to November 15, 2018. If a workteam member is added too early, you can either delete the member entirely and add them again, or correct the membership status of the first snapshot to be Inactive and then add a new snapshot with the correct effective date and a membership status of Active or Passive.
func (c *Client) PutWorkteam(ctx context.Context, input *PutWorkteamInput) (output *PutWorkteamOutput, err error) {
	output = &PutWorkteamOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutWorkteamInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Workteam_Request"`
	PutWorkteamRequestType
}

type PutWorkteamOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Workteam_Response"`
	PutWorkteamResponseType
}

// ChangeWorkContactInformation (Change_Work_Contact_Information)
// 
// Creates or updates work contact information for a person. Contact information includes email addresses, physical addresses, phone numbers, web addresses, and instant messenger contact information. Changes are routed through existing Work Contact Change business process for necessary approvals.
func (c *Client) ChangeWorkContactInformation(ctx context.Context, input *ChangeWorkContactInformationInput) (output *ChangeWorkContactInformationOutput, err error) {
	output = &ChangeWorkContactInformationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeWorkContactInformationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Work_Contact_Information_Request"`
	ChangeWorkContactInformationRequestType
}

type ChangeWorkContactInformationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Work_Contact_Information_Response"`
	ChangeWorkContactInformationResponseType
}

// GetChangeWorkContactInformation (Get_Change_Work_Contact_Information)
// 
// This public web service request gets contact information for a person. The response can be used as input to the web service request Change Work Contact Information.
func (c *Client) GetChangeWorkContactInformation(ctx context.Context, input *GetChangeWorkContactInformationInput) (output *GetChangeWorkContactInformationOutput, err error) {
	output = &GetChangeWorkContactInformationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetChangeWorkContactInformationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Work_Contact_Information_Request"`
	GetChangeWorkContactInformationRequestType
}

type GetChangeWorkContactInformationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Work_Contact_Information_Response"`
	GetChangeWorkContactInformationResponseType
}

// ChangeHomeContactInformation (Change_Home_Contact_Information)
// 
// Creates or updates home contact information for a person. Contact information includes email addresses, physical addresses, phone numbers, web addresses, and instant messenger contact information. Changes are routed through existing Home Contact Change business process for necessary approvals.
func (c *Client) ChangeHomeContactInformation(ctx context.Context, input *ChangeHomeContactInformationInput) (output *ChangeHomeContactInformationOutput, err error) {
	output = &ChangeHomeContactInformationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeHomeContactInformationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Home_Contact_Information_Request"`
	ChangeHomeContactInformationRequestType
}

type ChangeHomeContactInformationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Home_Contact_Information_Response"`
	ChangeHomeContactInformationResponseType
}

// GetChangeHomeContactInformation (Get_Change_Home_Contact_Information)
// 
// This public web service request gets contact information for a person. The response can be used as input to the web service request Change Home Contact Information.
func (c *Client) GetChangeHomeContactInformation(ctx context.Context, input *GetChangeHomeContactInformationInput) (output *GetChangeHomeContactInformationOutput, err error) {
	output = &GetChangeHomeContactInformationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetChangeHomeContactInformationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Home_Contact_Information_Request"`
	GetChangeHomeContactInformationRequestType
}

type GetChangeHomeContactInformationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Home_Contact_Information_Response"`
	GetChangeHomeContactInformationResponseType
}

// GetChangeVeteranStatusIdentification (Get_Change_Veteran_Status_Identification)
// 
// This web service request gets veteran status identification for a worker. The response can be used as input to the web service request Change Veteran Status Identification.
func (c *Client) GetChangeVeteranStatusIdentification(ctx context.Context, input *GetChangeVeteranStatusIdentificationInput) (output *GetChangeVeteranStatusIdentificationOutput, err error) {
	output = &GetChangeVeteranStatusIdentificationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetChangeVeteranStatusIdentificationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Veteran_Status_Identification_Request"`
	GetChangeVeteranStatusIdentificationRequestType
}

type GetChangeVeteranStatusIdentificationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Change_Veteran_Status_Identification_Response"`
	GetChangeVeteranStatusIdentificationResponseType
}

// PutExternalFormI9 (Put_External_Form_I-9)
// 
// This web services allows a record of an External Form I-9 to be put into Workday with only critical field information.
func (c *Client) PutExternalFormI9(ctx context.Context, input *PutExternalFormI9Input) (output *PutExternalFormI9Output, err error) {
	output = &PutExternalFormI9Output{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutExternalFormI9Input struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_External_Form_I-9_Request"`
	PutExternalFormI9RequestType
}

type PutExternalFormI9Output struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_External_Form_I-9_Response"`
	PutExternalFormI9ResponseType
}

// GetExternalFormI9Source (Get_External_Form_I-9_Source)
// 
// This web services allows an External Form I-9 Source to be extracted from Workday.
func (c *Client) GetExternalFormI9Source(ctx context.Context, input *GetExternalFormI9SourceInput) (output *GetExternalFormI9SourceOutput, err error) {
	output = &GetExternalFormI9SourceOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetExternalFormI9SourceInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_External_Form_I-9_Source_Request"`
	GetExternalFormI9SourceRequestType
}

type GetExternalFormI9SourceOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_External_Form_I-9_Source_Response"`
	GetExternalFormI9SourceResponseType
}

// PutExternalFormI9Source (Put_External_Form_I-9_Source)
// 
// This web service allows an External Form I-9 Source to be put into Workday.
func (c *Client) PutExternalFormI9Source(ctx context.Context, input *PutExternalFormI9SourceInput) (output *PutExternalFormI9SourceOutput, err error) {
	output = &PutExternalFormI9SourceOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutExternalFormI9SourceInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_External_Form_I-9_Source_Request"`
	PutExternalFormI9SourceRequestType
}

type PutExternalFormI9SourceOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_External_Form_I-9_Source_Response"`
	PutExternalFormI9SourceResponseType
}

// ReportSafetyIncident (Report_Safety_Incident)
// 
// This web service creates a Safety Incident.
func (c *Client) ReportSafetyIncident(ctx context.Context, input *ReportSafetyIncidentInput) (output *ReportSafetyIncidentOutput, err error) {
	output = &ReportSafetyIncidentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ReportSafetyIncidentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Report_Safety_Incident_Request"`
	ReportSafetyIncidentRequestType
}

type ReportSafetyIncidentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Report_Safety_Incident_Response"`
	ReportSafetyIncidentResponseType
}

// GetExternalFormI9 (Get_External_Form_I-9)
// 
// This web services allows a record of an External Form I-9 to be extracted from Workday.
func (c *Client) GetExternalFormI9(ctx context.Context, input *GetExternalFormI9Input) (output *GetExternalFormI9Output, err error) {
	output = &GetExternalFormI9Output{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetExternalFormI9Input struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_External_Form_I-9_Request"`
	GetExternalFormI9RequestType
}

type GetExternalFormI9Output struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_External_Form_I-9_Response"`
	GetExternalFormI9ResponseType
}

// UpdateSafetyIncident (Update_Safety_Incident)
// 
// This web service updates a Safety Incident.
func (c *Client) UpdateSafetyIncident(ctx context.Context, input *UpdateSafetyIncidentInput) (output *UpdateSafetyIncidentOutput, err error) {
	output = &UpdateSafetyIncidentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type UpdateSafetyIncidentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Update_Safety_Incident_Request"`
	UpdateSafetyIncidentRequestType
}

type UpdateSafetyIncidentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Update_Safety_Incident_Response"`
	UpdateSafetyIncidentResponseType
}

// ChangePersonPhoto (Change_Person_Photo)
// 
// This web service adds or updates a person photo.
func (c *Client) ChangePersonPhoto(ctx context.Context, input *ChangePersonPhotoInput) (output *ChangePersonPhotoOutput, err error) {
	output = &ChangePersonPhotoOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangePersonPhotoInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Person_Photo_Request"`
	ChangePersonPhotoRequestType
}

type ChangePersonPhotoOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Person_Photo_Response"`
	ChangePersonPhotoResponseType
}

// GetPersonPhotos (Get_Person_Photos)
// 
// Returns the person photo by Person Reference or Universal ID. It will filter out any person instances that are not enabled for photos.
func (c *Client) GetPersonPhotos(ctx context.Context, input *GetPersonPhotosInput) (output *GetPersonPhotosOutput, err error) {
	output = &GetPersonPhotosOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPersonPhotosInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Person_Photos_Request"`
	GetPersonPhotosRequestType
}

type GetPersonPhotosOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Person_Photos_Response"`
	GetPersonPhotosResponseType
}

// PutExternalFormI9Section3 (Put_External_Form_I-9_Section_3)
// 
// This web services allows a record of an External Form I-9 Section 3 to be put into Workday with only critical field information.
func (c *Client) PutExternalFormI9Section3(ctx context.Context, input *PutExternalFormI9Section3Input) (output *PutExternalFormI9Section3Output, err error) {
	output = &PutExternalFormI9Section3Output{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutExternalFormI9Section3Input struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_External_Form_I-9_Section_3_Request"`
	PutExternalFormI9Section3RequestType
}

type PutExternalFormI9Section3Output struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_External_Form_I-9_Section_3_Response"`
	PutExternalFormI9Section3ResponseType
}

// GetReportSafetyIncident (Get_Report_Safety_Incident)
// 
// This web services allows a safety incident to be extracted from Workday.
func (c *Client) GetReportSafetyIncident(ctx context.Context, input *GetReportSafetyIncidentInput) (output *GetReportSafetyIncidentOutput, err error) {
	output = &GetReportSafetyIncidentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetReportSafetyIncidentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Report_Safety_Incident_Request"`
	GetReportSafetyIncidentRequestType
}

type GetReportSafetyIncidentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Report_Safety_Incident_Response"`
	GetReportSafetyIncidentResponseType
}

// GetExternalFormI9Section3 (Get_External_Form_I-9_Section_3)
// 
// This web services allows a record of an External Form I-9 Section 3 to be extracted from Workday.
func (c *Client) GetExternalFormI9Section3(ctx context.Context, input *GetExternalFormI9Section3Input) (output *GetExternalFormI9Section3Output, err error) {
	output = &GetExternalFormI9Section3Output{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetExternalFormI9Section3Input struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_External_Form_I-9_Section_3_Request"`
	GetExternalFormI9Section3RequestType
}

type GetExternalFormI9Section3Output struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_External_Form_I-9_Section_3_Response"`
	GetExternalFormI9Section3ResponseType
}

// PutWorkteamMembership (Put_Workteam_Membership)
// 
// Add, Change or Remove a Workteam Membership.
func (c *Client) PutWorkteamMembership(ctx context.Context, input *PutWorkteamMembershipInput) (output *PutWorkteamMembershipOutput, err error) {
	output = &PutWorkteamMembershipOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutWorkteamMembershipInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Workteam_Membership_Request"`
	PutWorkteamMembershipRequestType
}

type PutWorkteamMembershipOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Workteam_Membership_Response"`
	PutWorkteamMembershipResponseType
}

// ImportUniversalIdentifiers (Import_Universal_Identifiers)
// 
// Import Universal Identifiers. It&#39;s required to put a Person Reference. Either place a manual Universal ID or no Universal ID, if no Universal ID is supplied it will use the Universal ID generator.
func (c *Client) ImportUniversalIdentifiers(ctx context.Context, input *ImportUniversalIdentifiersInput) (output *ImportUniversalIdentifiersOutput, err error) {
	output = &ImportUniversalIdentifiersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportUniversalIdentifiersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Universal_Identifier_Request"`
	ImportUniversalIdentifierRequestType
}

type ImportUniversalIdentifiersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportChangeWorkContactInformation (Import_Change_Work_Contact_Information)
// 
// This is a high volume version of the Change Work Contact Information web service.  Creates or updates work contact information for a person or persons. Contact information includes email addresses, physical addresses, phone numbers, web addresses, and instant messenger contact information. Changes are routed through existing Work Contact Change business process for necessary approvals.
func (c *Client) ImportChangeWorkContactInformation(ctx context.Context, input *ImportChangeWorkContactInformationInput) (output *ImportChangeWorkContactInformationOutput, err error) {
	output = &ImportChangeWorkContactInformationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportChangeWorkContactInformationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Change_Work_Contact_Information_Request"`
	ImportChangeWorkContactInformationRequestType
}

type ImportChangeWorkContactInformationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportPersonPhotos (Import_Person_Photos)
// 
// This is a high volume version of the Change Person Photo web service.  Creates or updates the photo for a person or persons. Changes are routed through existing Photo Change business process for necessary approvals.
func (c *Client) ImportPersonPhotos(ctx context.Context, input *ImportPersonPhotosInput) (output *ImportPersonPhotosOutput, err error) {
	output = &ImportPersonPhotosOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportPersonPhotosInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Person_Photo_Request"`
	ImportPersonPhotoRequestType
}

type ImportPersonPhotosOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

