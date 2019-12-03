
// Package benadmin
//
// The Benefits Administration Web Service contains operations that expose Workday Human Capital Management Business Services benefits-related data.
package benadmin

import (
	"context"

	"github.com/ianlopshire/go-workday/pkg/wws"
)

const (
	CurrentVersion = "v31.2"
	ServiceName = "Benefits_Administration"
)

type Client struct {
	*wws.Client
}


// GetBenefitAnnualRates (Get_Benefit_Annual_Rates)
// 
// This service operation retrieves Benefits Annual Rates for one or more employees.
func (c *Client) GetBenefitAnnualRates(ctx context.Context, input *GetBenefitAnnualRatesInput) (output *GetBenefitAnnualRatesOutput, err error) {
	output = &GetBenefitAnnualRatesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBenefitAnnualRatesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Benefit_Annual_Rates_Request"`
	GetBenefitAnnualRatesRequestType
}

type GetBenefitAnnualRatesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Benefit_Annual_Rates_Response"`
	GetBenefitAnnualRatesResponseType
}

// PutBenefitAnnualRate (Put_Benefit_Annual_Rate)
// 
// This service operation creates a Benefit Annual Rate for one or more employees.
func (c *Client) PutBenefitAnnualRate(ctx context.Context, input *PutBenefitAnnualRateInput) (output *PutBenefitAnnualRateOutput, err error) {
	output = &PutBenefitAnnualRateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBenefitAnnualRateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Benefit_Annual_Rate_Request"`
	PutBenefitAnnualRateRequestType
}

type PutBenefitAnnualRateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Benefit_Annual_Rate_Response"`
	PutBenefitAnnualRateResponseType
}

// PutEmployeeDefinedContributionElections (Put_Employee_Defined_Contribution_Elections)
// 
// DEPRECATED:  The functionality within this operation has been replaced by new functionality within the &#34;Change Benefits for Life Event&#34; operation.
// 
// This service operation previously adds, updates or deletes Defined Contribution Elections of an employee.
func (c *Client) PutEmployeeDefinedContributionElections(ctx context.Context, input *PutEmployeeDefinedContributionElectionsInput) (output *PutEmployeeDefinedContributionElectionsOutput, err error) {
	output = &PutEmployeeDefinedContributionElectionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutEmployeeDefinedContributionElectionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Employee_Defined_Contribution_Elections_Request"`
	PutEmployeeDefinedContributionElectionsRequestType
}

type PutEmployeeDefinedContributionElectionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Employee_Defined_Contribution_Elections_Response"`
	PutEmployeeDefinedContributionElectionsResponseType
}

// GetEmployeeDefinedContributionElections (Get_Employee_Defined_Contribution_Elections)
// 
// DEPRECATED: The functionality within this operation has been replaced by new functionality within the &#34;Get Change Benefits for Life Event&#34; operation.
//  
// This service operation previously retrieved Defined Contribution Elections for one or more employees.
func (c *Client) GetEmployeeDefinedContributionElections(ctx context.Context, input *GetEmployeeDefinedContributionElectionsInput) (output *GetEmployeeDefinedContributionElectionsOutput, err error) {
	output = &GetEmployeeDefinedContributionElectionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetEmployeeDefinedContributionElectionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Employee_Defined_Contribution_Elections_Request"`
	GetEmployeeDefinedContributionElectionsRequestType
}

type GetEmployeeDefinedContributionElectionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Employee_Defined_Contribution_Elections_Response"`
	GetEmployeeDefinedContributionElectionsResponseType
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

// GetHealthCareRates (Get_Health_Care_Rates)
// 
// This web service retrieves Health Care Rates.
func (c *Client) GetHealthCareRates(ctx context.Context, input *GetHealthCareRatesInput) (output *GetHealthCareRatesOutput, err error) {
	output = &GetHealthCareRatesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetHealthCareRatesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Health_Care_Rates_Request"`
	GetHealthCareRatesRequestType
}

type GetHealthCareRatesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Health_Care_Rates_Response"`
	GetHealthCareRatesResponseType
}

// PutHealthCareRate (Put_Health_Care_Rate)
// 
// This Web Service puts Health Care Rates.
func (c *Client) PutHealthCareRate(ctx context.Context, input *PutHealthCareRateInput) (output *PutHealthCareRateOutput, err error) {
	output = &PutHealthCareRateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutHealthCareRateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Health_Care_Rate_Request"`
	PutHealthCareRateRequestType
}

type PutHealthCareRateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Health_Care_Rate_Response"`
	PutHealthCareRateResponseType
}

// PutPostalCodeSet (Put_Postal_Code_Set)
// 
// Use this operation to add new postal codes to an existing postal code set, or to create a new postal code set.
func (c *Client) PutPostalCodeSet(ctx context.Context, input *PutPostalCodeSetInput) (output *PutPostalCodeSetOutput, err error) {
	output = &PutPostalCodeSetOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPostalCodeSetInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Postal_Code_Set_Request"`
	PutPostalCodeSetRequestType
}

type PutPostalCodeSetOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Postal_Code_Set_Response"`
	PutPostalCodeSetResponseType
}

// GetBenefitIndividualRates (Get_Benefit_Individual_Rates)
// 
// This service operation retrieves Benefits Individual Rates for one or more employees.
func (c *Client) GetBenefitIndividualRates(ctx context.Context, input *GetBenefitIndividualRatesInput) (output *GetBenefitIndividualRatesOutput, err error) {
	output = &GetBenefitIndividualRatesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBenefitIndividualRatesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Benefit_Individual_Rates_Request"`
	GetBenefitIndividualRatesRequestType
}

type GetBenefitIndividualRatesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Benefit_Individual_Rates_Response"`
	GetBenefitIndividualRatesResponseType
}

// PutBenefitIndividualRate (Put_Benefit_Individual_Rate)
// 
// This service operation updates Benefits Individual Rates for one or more employees.
func (c *Client) PutBenefitIndividualRate(ctx context.Context, input *PutBenefitIndividualRateInput) (output *PutBenefitIndividualRateOutput, err error) {
	output = &PutBenefitIndividualRateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBenefitIndividualRateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Benefit_Individual_Rate_Request"`
	PutBenefitIndividualRateRequestType
}

type PutBenefitIndividualRateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Benefit_Individual_Rate_Response"`
	PutBenefitIndividualRateResponseType
}

// ChangeBenefits (Change_Benefits)
// 
// This service operation initiates the &#34;Change Benefits for Life Event&#34; business process allowing changes to benefit elections for an employee.  This service does not support incremental updates or defaults.
func (c *Client) ChangeBenefits(ctx context.Context, input *ChangeBenefitsInput) (output *ChangeBenefitsOutput, err error) {
	output = &ChangeBenefitsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeBenefitsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Benefits_Request"`
	ChangeBenefitsRequestType
}

type ChangeBenefitsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Benefits_Response"`
	ChangeBenefitsResponseType
}

// EnrollinRetirementSavingsPlans (Enroll_in_Retirement_Savings_Plans)
// 
// This service operation allows for enrollment into Retirement Savings Plans. It updates incrementally from existing elections and carries forward prior elections unless they are explicitly changed or waived.
func (c *Client) EnrollinRetirementSavingsPlans(ctx context.Context, input *EnrollinRetirementSavingsPlansInput) (output *EnrollinRetirementSavingsPlansOutput, err error) {
	output = &EnrollinRetirementSavingsPlansOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EnrollinRetirementSavingsPlansInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Enroll_in_Retirement_Savings_Plans_Request"`
	EnrollinRetirementSavingsPlansRequestType
}

type EnrollinRetirementSavingsPlansOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Enroll_in_Retirement_Savings_Plans_Response"`
	EnrollinRetirementSavingsPlansResponseType
}

// GrantCOBRAEligibility (Grant_COBRA_Eligibility)
// 
// A transaction to modify existing COBRA Eligibility or add new COBRA Eligibility Entries
func (c *Client) GrantCOBRAEligibility(ctx context.Context, input *GrantCOBRAEligibilityInput) (output *GrantCOBRAEligibilityOutput, err error) {
	output = &GrantCOBRAEligibilityOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GrantCOBRAEligibilityInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Grant_COBRA_Eligibility_Request"`
	GrantCOBRAEligibilityRequestType
}

type GrantCOBRAEligibilityOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Grant_COBRA_Eligibility_Response"`
	GrantCOBRAEligibilityResponseType
}

// AddDependent (Add_Dependent)
// 
// This operation adds a Dependent.
func (c *Client) AddDependent(ctx context.Context, input *AddDependentInput) (output *AddDependentOutput, err error) {
	output = &AddDependentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type AddDependentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Add_Dependent_Request"`
	AddDependentRequestType
}

type AddDependentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Add_Dependent_Response"`
	AddDependentResponseType
}

// EditDependent (Edit_Dependent)
// 
// This operation updates a dependent.
func (c *Client) EditDependent(ctx context.Context, input *EditDependentInput) (output *EditDependentOutput, err error) {
	output = &EditDependentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EditDependentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Dependent_Request"`
	EditDependentRequestType
}

type EditDependentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Edit_Dependent_Response"`
	EditDependentResponseType
}

// ChangeBenefitJobs (Change_Benefit_Jobs)
// 
// Specifies the benefit jobs for a worker.
func (c *Client) ChangeBenefitJobs(ctx context.Context, input *ChangeBenefitJobsInput) (output *ChangeBenefitJobsOutput, err error) {
	output = &ChangeBenefitJobsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeBenefitJobsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Benefit_Jobs_Request"`
	ChangeBenefitJobsRequestType
}

type ChangeBenefitJobsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Benefit_Jobs_Response"`
	ChangeBenefitJobsResponseType
}

// GetBenefitAnnualCredits (Get_Benefit_Annual_Credits)
// 
// This service operation retrieves Benefits Annual Credits for one or more employees.
func (c *Client) GetBenefitAnnualCredits(ctx context.Context, input *GetBenefitAnnualCreditsInput) (output *GetBenefitAnnualCreditsOutput, err error) {
	output = &GetBenefitAnnualCreditsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBenefitAnnualCreditsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Benefit_Annual_Credits_Request"`
	GetBenefitAnnualCreditsRequestType
}

type GetBenefitAnnualCreditsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Benefit_Annual_Credits_Response"`
	GetBenefitAnnualCreditsResponseType
}

// PutBenefitAnnualCredit (Put_Benefit_Annual_Credit)
// 
// This service operation creates a Benefit Annual Credit for one or more employees.
func (c *Client) PutBenefitAnnualCredit(ctx context.Context, input *PutBenefitAnnualCreditInput) (output *PutBenefitAnnualCreditOutput, err error) {
	output = &PutBenefitAnnualCreditOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBenefitAnnualCreditInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Benefit_Annual_Credit_Request"`
	PutBenefitAnnualCreditRequestType
}

type PutBenefitAnnualCreditOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Benefit_Annual_Credit_Response"`
	PutBenefitAnnualCreditResponseType
}

// PutWorkerWellnessData (Put_Worker_Wellness_Data)
// 
// This service operation creates/updates Wellness Data for a given worker.
func (c *Client) PutWorkerWellnessData(ctx context.Context, input *PutWorkerWellnessDataInput) (output *PutWorkerWellnessDataOutput, err error) {
	output = &PutWorkerWellnessDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutWorkerWellnessDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Worker_Wellness_Data_Request"`
	PutWorkerWellnessDataRequestType
}

type PutWorkerWellnessDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Worker_Wellness_Data_Response"`
	PutWorkerWellnessDataResponseType
}

// PutAffordableCareActWorkerHoursAndWages (Put_Affordable_Care_Act_Worker_Hours_And_Wages)
// 
// This service operation creates/updates Affordable Care Act Worker Hours for a given worker.
func (c *Client) PutAffordableCareActWorkerHoursAndWages(ctx context.Context, input *PutAffordableCareActWorkerHoursAndWagesInput) (output *PutAffordableCareActWorkerHoursAndWagesOutput, err error) {
	output = &PutAffordableCareActWorkerHoursAndWagesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutAffordableCareActWorkerHoursAndWagesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Affordable_Care_Act_Worker_Hours_And_Wages_Request"`
	PutAffordableCareActWorkerHoursAndWagesRequestType
}

type PutAffordableCareActWorkerHoursAndWagesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Affordable_Care_Act_Worker_Hours_And_Wages_Response"`
	PutAffordableCareActWorkerHoursAndWagesResponseType
}

// PutEvidenceofInsurability (Put_Evidence_of_Insurability)
// 
// This Public Web Service is developed to  update Pending EOI status in Workday from external systems.
func (c *Client) PutEvidenceofInsurability(ctx context.Context, input *PutEvidenceofInsurabilityInput) (output *PutEvidenceofInsurabilityOutput, err error) {
	output = &PutEvidenceofInsurabilityOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutEvidenceofInsurabilityInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Evidence_of_Insurability_Request"`
	PutEvidenceofInsurabilityRequestType
}

type PutEvidenceofInsurabilityOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Evidence_of_Insurability_Response"`
	PutEvidenceofInsurabilityResponseType
}

// ChangeBeneficiary (Change_Beneficiary)
// 
// This web service allows the updating of a beneficiary for a person.
func (c *Client) ChangeBeneficiary(ctx context.Context, input *ChangeBeneficiaryInput) (output *ChangeBeneficiaryOutput, err error) {
	output = &ChangeBeneficiaryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ChangeBeneficiaryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Beneficiary_Request"`
	ChangeBeneficiaryRequestType
}

type ChangeBeneficiaryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Change_Beneficiary_Response"`
	ChangeBeneficiaryResponseType
}

// GetACA1095CFormsData (Get_ACA_1095-C_Forms_Data)
// 
// Get ACA 1095-C Forms Data
func (c *Client) GetACA1095CFormsData(ctx context.Context, input *GetACA1095CFormsDataInput) (output *GetACA1095CFormsDataOutput, err error) {
	output = &GetACA1095CFormsDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetACA1095CFormsDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_ACA_1095-C_Forms_Data_Request"`
	GetACA1095CFormsDataRequestType
}

type GetACA1095CFormsDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_ACA_1095-C_Forms_Data_Response"`
	GetACA1095CFormsDataResponseType
}

// ManageMedicareInformation (Manage_Medicare_Information)
// 
// This operation invokes the Manage Medicare Information business process.
func (c *Client) ManageMedicareInformation(ctx context.Context, input *ManageMedicareInformationInput) (output *ManageMedicareInformationOutput, err error) {
	output = &ManageMedicareInformationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ManageMedicareInformationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Medicare_Information_Request"`
	ManageMedicareInformationRequestType
}

type ManageMedicareInformationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Medicare_Information_Response"`
	ManageMedicareInformationResponseType
}

// GetMedicarePartDEGWP (Get_Medicare_Part_D_EGWP)
// 
// Get Medicare Part D EGWP Data, currently will retrieve all Medicare Part D Data for all persons
func (c *Client) GetMedicarePartDEGWP(ctx context.Context, input *GetMedicarePartDEGWPInput) (output *GetMedicarePartDEGWPOutput, err error) {
	output = &GetMedicarePartDEGWPOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetMedicarePartDEGWPInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Medicare_Part_D_EGWP_Request"`
	GetMedicarePartDEGWPRequestType
}

type GetMedicarePartDEGWPOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Medicare_Part_D_EGWP_Response"`
	GetMedicarePartDEGWPResponseType
}

// Import1095CFormRecipientsData (Import_1095-C_Form_Recipients_Data)
// 
// This operation updates 1095-C Form Data for a given 1095-C Recipient (or) adds 1095-C Form and External 1095-C Recipient Data for a new 1095-C Recipient
func (c *Client) Import1095CFormRecipientsData(ctx context.Context, input *Import1095CFormRecipientsDataInput) (output *Import1095CFormRecipientsDataOutput, err error) {
	output = &Import1095CFormRecipientsDataOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type Import1095CFormRecipientsDataInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_ACA_1095-C_Forms_Data_Request"`
	ImportACA1095CFormsDataRequestType
}

type Import1095CFormRecipientsDataOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportMedicarePartDEGWP (Import_Medicare_Part_D_EGWP)
// 
// Imports Medicare Part D EGWP Data using HICN Medicare ID
func (c *Client) ImportMedicarePartDEGWP(ctx context.Context, input *ImportMedicarePartDEGWPInput) (output *ImportMedicarePartDEGWPOutput, err error) {
	output = &ImportMedicarePartDEGWPOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportMedicarePartDEGWPInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Medicare_Part_D_EGWP_Request"`
	ImportMedicarePartDEGWPRequestType
}

type ImportMedicarePartDEGWPOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

