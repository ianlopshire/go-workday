

# comp
`import "github.com/ianlopshire/go-workday/pkg/wws/services/comp"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package comp

The Compensation Web Service contains operations that expose Workday Human Capital Management Business Services compensation-related data.




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [type AcademicPeriodObjectIDType](#AcademicPeriodObjectIDType)
* [type AcademicPeriodObjectType](#AcademicPeriodObjectType)
* [type AddIndividualStockGrantDataType](#AddIndividualStockGrantDataType)
  * [func (t *AddIndividualStockGrantDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#AddIndividualStockGrantDataType.MarshalXML)
  * [func (t *AddIndividualStockGrantDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#AddIndividualStockGrantDataType.UnmarshalXML)
* [type AddStockGrantDataType](#AddStockGrantDataType)
  * [func (t *AddStockGrantDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#AddStockGrantDataType.MarshalXML)
  * [func (t *AddStockGrantDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#AddStockGrantDataType.UnmarshalXML)
* [type AddStockGrantInput](#AddStockGrantInput)
* [type AddStockGrantOutput](#AddStockGrantOutput)
* [type AddStockGrantRequestType](#AddStockGrantRequestType)
* [type AddStockGrantResponseType](#AddStockGrantResponseType)
* [type AllocationDetailforPeriodPayDataType](#AllocationDetailforPeriodPayDataType)
* [type AllowancePlanAmountBasedProfileReplacementDataType](#AllowancePlanAmountBasedProfileReplacementDataType)
* [type AllowancePlanAmountDataType](#AllowancePlanAmountDataType)
* [type AllowancePlanDataType](#AllowancePlanDataType)
* [type AllowancePlanPercentBasedProfileReplacementDataType](#AllowancePlanPercentBasedProfileReplacementDataType)
* [type AllowancePlanPercentDataType](#AllowancePlanPercentDataType)
* [type AllowancePlanReimbursableDataType](#AllowancePlanReimbursableDataType)
* [type AllowancePlanUnitBasedProfileReplacementDataType](#AllowancePlanUnitBasedProfileReplacementDataType)
* [type AllowancePlanUnitDataType](#AllowancePlanUnitDataType)
* [type AllowanceUnitPlanObjectIDType](#AllowanceUnitPlanObjectIDType)
* [type AllowanceUnitPlanObjectType](#AllowanceUnitPlanObjectType)
* [type AllowanceValuePlanObjectIDType](#AllowanceValuePlanObjectIDType)
* [type AllowanceValuePlanObjectType](#AllowanceValuePlanObjectType)
* [type AlternatePayRangeDataType](#AlternatePayRangeDataType)
* [type AlternativePayRangeType](#AlternativePayRangeType)
* [type AndOrOperatorsObjectIDType](#AndOrOperatorsObjectIDType)
* [type AndOrOperatorsObjectType](#AndOrOperatorsObjectType)
* [type AssignEligiblePeriodActivitiesforEmployeeDataType](#AssignEligiblePeriodActivitiesforEmployeeDataType)
  * [func (t *AssignEligiblePeriodActivitiesforEmployeeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#AssignEligiblePeriodActivitiesforEmployeeDataType.MarshalXML)
  * [func (t *AssignEligiblePeriodActivitiesforEmployeeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#AssignEligiblePeriodActivitiesforEmployeeDataType.UnmarshalXML)
* [type AssignEligiblePeriodActivitiesforEmployeeInput](#AssignEligiblePeriodActivitiesforEmployeeInput)
* [type AssignEligiblePeriodActivitiesforEmployeeOutput](#AssignEligiblePeriodActivitiesforEmployeeOutput)
* [type AssignEligiblePeriodActivitiesforEmployeeRequestType](#AssignEligiblePeriodActivitiesforEmployeeRequestType)
* [type AssignEligiblePeriodActivitiesforEmployeeResponseType](#AssignEligiblePeriodActivitiesforEmployeeResponseType)
* [type BasePayPlanDataType](#BasePayPlanDataType)
* [type BenchmarkJobCompositeDataType](#BenchmarkJobCompositeDataType)
* [type BenchmarkJobDataType](#BenchmarkJobDataType)
  * [func (t *BenchmarkJobDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#BenchmarkJobDataType.MarshalXML)
  * [func (t *BenchmarkJobDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#BenchmarkJobDataType.UnmarshalXML)
* [type BenchmarkJobRequestCriteriaType](#BenchmarkJobRequestCriteriaType)
* [type BenchmarkJobRequestReferencesType](#BenchmarkJobRequestReferencesType)
* [type BenchmarkJobResponseDataType](#BenchmarkJobResponseDataType)
* [type BenchmarkJobResponseGroupType](#BenchmarkJobResponseGroupType)
* [type BenchmarkJobType](#BenchmarkJobType)
* [type BenefitMultiplierOrderObjectIDType](#BenefitMultiplierOrderObjectIDType)
* [type BenefitMultiplierOrderObjectType](#BenefitMultiplierOrderObjectType)
* [type BonusPaymentDataType](#BonusPaymentDataType)
* [type BonusPercentPlanObjectIDType](#BonusPercentPlanObjectIDType)
* [type BonusPercentPlanObjectType](#BonusPercentPlanObjectType)
* [type BonusPlanAmountDataType](#BonusPlanAmountDataType)
* [type BonusPlanDataType](#BonusPlanDataType)
* [type BonusPlanObjectIDType](#BonusPlanObjectIDType)
* [type BonusPlanObjectType](#BonusPlanObjectType)
* [type BonusPlanPercentDataType](#BonusPlanPercentDataType)
* [type BonusPlanPercentPlanProfileReplacementDataType](#BonusPlanPercentPlanProfileReplacementDataType)
* [type BonusPlanProfileReplacementAmountDataType](#BonusPlanProfileReplacementAmountDataType)
* [type BusinessProcessAttachmentDataType](#BusinessProcessAttachmentDataType)
  * [func (t *BusinessProcessAttachmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#BusinessProcessAttachmentDataType.MarshalXML)
  * [func (t *BusinessProcessAttachmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#BusinessProcessAttachmentDataType.UnmarshalXML)
* [type BusinessProcessCommentDataType](#BusinessProcessCommentDataType)
* [type BusinessProcessParametersType](#BusinessProcessParametersType)
* [type CalculatedPlanDataType](#CalculatedPlanDataType)
* [type CalculatedPlanObjectIDType](#CalculatedPlanObjectIDType)
* [type CalculatedPlanObjectType](#CalculatedPlanObjectType)
* [type CalculationObjectIDType](#CalculationObjectIDType)
* [type CalculationObjectType](#CalculationObjectType)
* [type CheckPositionBudgetSubBusinessProcessType](#CheckPositionBudgetSubBusinessProcessType)
* [type Client](#Client)
  * [func (c *Client) AddStockGrant(ctx context.Context, input *AddStockGrantInput) (output *AddStockGrantOutput, err error)](#Client.AddStockGrant)
  * [func (c *Client) AssignEligiblePeriodActivitiesforEmployee(ctx context.Context, input *AssignEligiblePeriodActivitiesforEmployeeInput) (output *AssignEligiblePeriodActivitiesforEmployeeOutput, err error)](#Client.AssignEligiblePeriodActivitiesforEmployee)
  * [func (c *Client) CreateSeveranceWorksheet(ctx context.Context, input *CreateSeveranceWorksheetInput) (output *CreateSeveranceWorksheetOutput, err error)](#Client.CreateSeveranceWorksheet)
  * [func (c *Client) GetBenchmarkJobs(ctx context.Context, input *GetBenchmarkJobsInput) (output *GetBenchmarkJobsOutput, err error)](#Client.GetBenchmarkJobs)
  * [func (c *Client) GetCompensationEligibilityRules(ctx context.Context, input *GetCompensationEligibilityRulesInput) (output *GetCompensationEligibilityRulesOutput, err error)](#Client.GetCompensationEligibilityRules)
  * [func (c *Client) GetCompensationEligibilityRuleswithoutDependencies(ctx context.Context, input *GetCompensationEligibilityRuleswithoutDependenciesInput) (output *GetCompensationEligibilityRuleswithoutDependenciesOutput, err error)](#Client.GetCompensationEligibilityRuleswithoutDependencies)
  * [func (c *Client) GetCompensationGradeHierarchy(ctx context.Context, input *GetCompensationGradeHierarchyInput) (output *GetCompensationGradeHierarchyOutput, err error)](#Client.GetCompensationGradeHierarchy)
  * [func (c *Client) GetCompensationGrades(ctx context.Context, input *GetCompensationGradesInput) (output *GetCompensationGradesOutput, err error)](#Client.GetCompensationGrades)
  * [func (c *Client) GetCompensationMatrices(ctx context.Context, input *GetCompensationMatricesInput) (output *GetCompensationMatricesOutput, err error)](#Client.GetCompensationMatrices)
  * [func (c *Client) GetCompensationPlans(ctx context.Context, input *GetCompensationPlansInput) (output *GetCompensationPlansOutput, err error)](#Client.GetCompensationPlans)
  * [func (c *Client) GetCompensationScorecardResults(ctx context.Context, input *GetCompensationScorecardResultsInput) (output *GetCompensationScorecardResultsOutput, err error)](#Client.GetCompensationScorecardResults)
  * [func (c *Client) GetCompensationScorecards(ctx context.Context, input *GetCompensationScorecardsInput) (output *GetCompensationScorecardsOutput, err error)](#Client.GetCompensationScorecards)
  * [func (c *Client) GetEligibleEarnings(ctx context.Context, input *GetEligibleEarningsInput) (output *GetEligibleEarningsOutput, err error)](#Client.GetEligibleEarnings)
  * [func (c *Client) GetEnhancedSeveranceMatrices(ctx context.Context, input *GetEnhancedSeveranceMatricesInput) (output *GetEnhancedSeveranceMatricesOutput, err error)](#Client.GetEnhancedSeveranceMatrices)
  * [func (c *Client) GetFuturePaymentPlanAssignments(ctx context.Context, input *GetFuturePaymentPlanAssignmentsInput) (output *GetFuturePaymentPlanAssignmentsOutput, err error)](#Client.GetFuturePaymentPlanAssignments)
  * [func (c *Client) GetOneTimePaymentPlanConfigurableCategories(ctx context.Context, input *GetOneTimePaymentPlanConfigurableCategoriesInput) (output *GetOneTimePaymentPlanConfigurableCategoriesOutput, err error)](#Client.GetOneTimePaymentPlanConfigurableCategories)
  * [func (c *Client) GetPeriodActivityRateMatrices(ctx context.Context, input *GetPeriodActivityRateMatricesInput) (output *GetPeriodActivityRateMatricesOutput, err error)](#Client.GetPeriodActivityRateMatrices)
  * [func (c *Client) GetPeriodActivityTasks(ctx context.Context, input *GetPeriodActivityTasksInput) (output *GetPeriodActivityTasksOutput, err error)](#Client.GetPeriodActivityTasks)
  * [func (c *Client) GetPreviousSystemCompensationHistory(ctx context.Context, input *GetPreviousSystemCompensationHistoryInput) (output *GetPreviousSystemCompensationHistoryOutput, err error)](#Client.GetPreviousSystemCompensationHistory)
  * [func (c *Client) GetStockParticipationRateTables(ctx context.Context, input *GetStockParticipationRateTablesInput) (output *GetStockParticipationRateTablesOutput, err error)](#Client.GetStockParticipationRateTables)
  * [func (c *Client) ImportEligibleEarningsOverride(ctx context.Context, input *ImportEligibleEarningsOverrideInput) (output *ImportEligibleEarningsOverrideOutput, err error)](#Client.ImportEligibleEarningsOverride)
  * [func (c *Client) ImportRequestCompensationChange(ctx context.Context, input *ImportRequestCompensationChangeInput) (output *ImportRequestCompensationChangeOutput, err error)](#Client.ImportRequestCompensationChange)
  * [func (c *Client) ManagePeriodActivityPayAssignments(ctx context.Context, input *ManagePeriodActivityPayAssignmentsInput) (output *ManagePeriodActivityPayAssignmentsOutput, err error)](#Client.ManagePeriodActivityPayAssignments)
  * [func (c *Client) PutBenchmarkJob(ctx context.Context, input *PutBenchmarkJobInput) (output *PutBenchmarkJobOutput, err error)](#Client.PutBenchmarkJob)
  * [func (c *Client) PutCompensationEligibilityRule(ctx context.Context, input *PutCompensationEligibilityRuleInput) (output *PutCompensationEligibilityRuleOutput, err error)](#Client.PutCompensationEligibilityRule)
  * [func (c *Client) PutCompensationGrade(ctx context.Context, input *PutCompensationGradeInput) (output *PutCompensationGradeOutput, err error)](#Client.PutCompensationGrade)
  * [func (c *Client) PutCompensationGradeHierarchy(ctx context.Context, input *PutCompensationGradeHierarchyInput) (output *PutCompensationGradeHierarchyOutput, err error)](#Client.PutCompensationGradeHierarchy)
  * [func (c *Client) PutCompensationMatrix(ctx context.Context, input *PutCompensationMatrixInput) (output *PutCompensationMatrixOutput, err error)](#Client.PutCompensationMatrix)
  * [func (c *Client) PutCompensationPlans(ctx context.Context, input *PutCompensationPlansInput) (output *PutCompensationPlansOutput, err error)](#Client.PutCompensationPlans)
  * [func (c *Client) PutCompensationScorecard(ctx context.Context, input *PutCompensationScorecardInput) (output *PutCompensationScorecardOutput, err error)](#Client.PutCompensationScorecard)
  * [func (c *Client) PutCompensationScorecardResult(ctx context.Context, input *PutCompensationScorecardResultInput) (output *PutCompensationScorecardResultOutput, err error)](#Client.PutCompensationScorecardResult)
  * [func (c *Client) PutEligibleEarnings(ctx context.Context, input *PutEligibleEarningsInput) (output *PutEligibleEarningsOutput, err error)](#Client.PutEligibleEarnings)
  * [func (c *Client) PutEnhancedSeveranceMatrix(ctx context.Context, input *PutEnhancedSeveranceMatrixInput) (output *PutEnhancedSeveranceMatrixOutput, err error)](#Client.PutEnhancedSeveranceMatrix)
  * [func (c *Client) PutFuturePaymentPlanAssignment(ctx context.Context, input *PutFuturePaymentPlanAssignmentInput) (output *PutFuturePaymentPlanAssignmentOutput, err error)](#Client.PutFuturePaymentPlanAssignment)
  * [func (c *Client) PutOneTimePaymentPlanConfigurableCategory(ctx context.Context, input *PutOneTimePaymentPlanConfigurableCategoryInput) (output *PutOneTimePaymentPlanConfigurableCategoryOutput, err error)](#Client.PutOneTimePaymentPlanConfigurableCategory)
  * [func (c *Client) PutPeriodActivityRateMatrix(ctx context.Context, input *PutPeriodActivityRateMatrixInput) (output *PutPeriodActivityRateMatrixOutput, err error)](#Client.PutPeriodActivityRateMatrix)
  * [func (c *Client) PutPeriodActivityTask(ctx context.Context, input *PutPeriodActivityTaskInput) (output *PutPeriodActivityTaskOutput, err error)](#Client.PutPeriodActivityTask)
  * [func (c *Client) PutPreviousSystemCompensationHistory(ctx context.Context, input *PutPreviousSystemCompensationHistoryInput) (output *PutPreviousSystemCompensationHistoryOutput, err error)](#Client.PutPreviousSystemCompensationHistory)
  * [func (c *Client) PutStockParticipationRateTable(ctx context.Context, input *PutStockParticipationRateTableInput) (output *PutStockParticipationRateTableOutput, err error)](#Client.PutStockParticipationRateTable)
  * [func (c *Client) RequestBonusPayment(ctx context.Context, input *RequestBonusPaymentInput) (output *RequestBonusPaymentOutput, err error)](#Client.RequestBonusPayment)
  * [func (c *Client) RequestCompensationChange(ctx context.Context, input *RequestCompensationChangeInput) (output *RequestCompensationChangeOutput, err error)](#Client.RequestCompensationChange)
  * [func (c *Client) RequestEmployeeMeritAdjustment(ctx context.Context, input *RequestEmployeeMeritAdjustmentInput) (output *RequestEmployeeMeritAdjustmentOutput, err error)](#Client.RequestEmployeeMeritAdjustment)
  * [func (c *Client) RequestOneTimePayment(ctx context.Context, input *RequestOneTimePaymentInput) (output *RequestOneTimePaymentOutput, err error)](#Client.RequestOneTimePayment)
  * [func (c *Client) RequestRequisitionCompensationChange(ctx context.Context, input *RequestRequisitionCompensationChangeInput) (output *RequestRequisitionCompensationChangeOutput, err error)](#Client.RequestRequisitionCompensationChange)
  * [func (c *Client) RetrieveSeveranceWorksheet(ctx context.Context, input *RetrieveSeveranceWorksheetInput) (output *RetrieveSeveranceWorksheetOutput, err error)](#Client.RetrieveSeveranceWorksheet)
  * [func (c *Client) UpdateStockGrant(ctx context.Context, input *UpdateStockGrantInput) (output *UpdateStockGrantOutput, err error)](#Client.UpdateStockGrant)
* [type CommissionPlanDataType](#CommissionPlanDataType)
* [type CommissionPlanObjectIDType](#CommissionPlanObjectIDType)
* [type CommissionPlanObjectType](#CommissionPlanObjectType)
* [type CommissionPlanProfileReplacementDataType](#CommissionPlanProfileReplacementDataType)
* [type CommonYesNoObjectIDType](#CommonYesNoObjectIDType)
* [type CommonYesNoObjectType](#CommonYesNoObjectType)
* [type CompaRatioRangeSegmentObjectIDType](#CompaRatioRangeSegmentObjectIDType)
* [type CompaRatioRangeSegmentObjectType](#CompaRatioRangeSegmentObjectType)
* [type CompensatableDefaultGuidelinesDataType](#CompensatableDefaultGuidelinesDataType)
* [type CompensatableGuidelinesDataType](#CompensatableGuidelinesDataType)
  * [func (t *CompensatableGuidelinesDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#CompensatableGuidelinesDataType.MarshalXML)
  * [func (t *CompensatableGuidelinesDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#CompensatableGuidelinesDataType.UnmarshalXML)
* [type CompensationAssignableComponentTypeObjectIDType](#CompensationAssignableComponentTypeObjectIDType)
* [type CompensationAssignableComponentTypeObjectType](#CompensationAssignableComponentTypeObjectType)
* [type CompensationAssignablePlanObjectIDType](#CompensationAssignablePlanObjectIDType)
* [type CompensationAssignablePlanObjectType](#CompensationAssignablePlanObjectType)
* [type CompensationBasisConfigurableObjectIDType](#CompensationBasisConfigurableObjectIDType)
* [type CompensationBasisConfigurableObjectType](#CompensationBasisConfigurableObjectType)
* [type CompensationBasisObjectIDType](#CompensationBasisObjectIDType)
* [type CompensationBasisObjectType](#CompensationBasisObjectType)
* [type CompensationBenchmarkAmountReplacmentDataType](#CompensationBenchmarkAmountReplacmentDataType)
* [type CompensationBenchmarkCompositeCategoryObjectIDType](#CompensationBenchmarkCompositeCategoryObjectIDType)
* [type CompensationBenchmarkCompositeCategoryObjectType](#CompensationBenchmarkCompositeCategoryObjectType)
* [type CompensationBenchmarkDefaultObjectIDType](#CompensationBenchmarkDefaultObjectIDType)
* [type CompensationBenchmarkDefaultObjectType](#CompensationBenchmarkDefaultObjectType)
* [type CompensationBenchmarkPercentileObjectIDType](#CompensationBenchmarkPercentileObjectIDType)
* [type CompensationBenchmarkPercentileObjectType](#CompensationBenchmarkPercentileObjectType)
* [type CompensationBenchmarkProfileDataType](#CompensationBenchmarkProfileDataType)
* [type CompensationChangeDataType](#CompensationChangeDataType)
* [type CompensationDurationSeveranceMatrixDataType](#CompensationDurationSeveranceMatrixDataType)
* [type CompensationDurationSeveranceMatrixEntryDataType](#CompensationDurationSeveranceMatrixEntryDataType)
* [type CompensationEligibilityRuleDataType](#CompensationEligibilityRuleDataType)
  * [func (t *CompensationEligibilityRuleDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#CompensationEligibilityRuleDataType.MarshalXML)
  * [func (t *CompensationEligibilityRuleDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#CompensationEligibilityRuleDataType.UnmarshalXML)
* [type CompensationEligibilityRuleRequestCriteriaType](#CompensationEligibilityRuleRequestCriteriaType)
* [type CompensationEligibilityRuleRequestReferencesType](#CompensationEligibilityRuleRequestReferencesType)
* [type CompensationEligibilityRuleResponseDataType](#CompensationEligibilityRuleResponseDataType)
* [type CompensationEligibilityRuleResponseGroupType](#CompensationEligibilityRuleResponseGroupType)
* [type CompensationEligibilityRuleType](#CompensationEligibilityRuleType)
* [type CompensationGradeDataType](#CompensationGradeDataType)
  * [func (t *CompensationGradeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#CompensationGradeDataType.MarshalXML)
  * [func (t *CompensationGradeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#CompensationGradeDataType.UnmarshalXML)
* [type CompensationGradeHierarchyDataType](#CompensationGradeHierarchyDataType)
  * [func (t *CompensationGradeHierarchyDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#CompensationGradeHierarchyDataType.MarshalXML)
  * [func (t *CompensationGradeHierarchyDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#CompensationGradeHierarchyDataType.UnmarshalXML)
* [type CompensationGradeHierarchyRequestReferencesType](#CompensationGradeHierarchyRequestReferencesType)
  * [func (t *CompensationGradeHierarchyRequestReferencesType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#CompensationGradeHierarchyRequestReferencesType.MarshalXML)
  * [func (t *CompensationGradeHierarchyRequestReferencesType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#CompensationGradeHierarchyRequestReferencesType.UnmarshalXML)
* [type CompensationGradeHierarchyResponseDataType](#CompensationGradeHierarchyResponseDataType)
* [type CompensationGradeHierarchyType](#CompensationGradeHierarchyType)
* [type CompensationGradeLevelDataType](#CompensationGradeLevelDataType)
* [type CompensationGradeLevelObjectIDType](#CompensationGradeLevelObjectIDType)
* [type CompensationGradeLevelObjectType](#CompensationGradeLevelObjectType)
* [type CompensationGradeLevelType](#CompensationGradeLevelType)
* [type CompensationGradeObjectIDType](#CompensationGradeObjectIDType)
* [type CompensationGradeObjectType](#CompensationGradeObjectType)
* [type CompensationGradeProfileDataType](#CompensationGradeProfileDataType)
  * [func (t *CompensationGradeProfileDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#CompensationGradeProfileDataType.MarshalXML)
  * [func (t *CompensationGradeProfileDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#CompensationGradeProfileDataType.UnmarshalXML)
* [type CompensationGradeProfileObjectIDType](#CompensationGradeProfileObjectIDType)
* [type CompensationGradeProfileObjectType](#CompensationGradeProfileObjectType)
* [type CompensationGradeProfileType](#CompensationGradeProfileType)
* [type CompensationGradeRequestCriteriaType](#CompensationGradeRequestCriteriaType)
* [type CompensationGradeRequestReferencesType](#CompensationGradeRequestReferencesType)
* [type CompensationGradeResponseDataType](#CompensationGradeResponseDataType)
* [type CompensationGradeResponseGroupType](#CompensationGradeResponseGroupType)
* [type CompensationGradeType](#CompensationGradeType)
* [type CompensationLengthofServiceSeveranceMatrixDataType](#CompensationLengthofServiceSeveranceMatrixDataType)
* [type CompensationLengthofServiceSeveranceMatrixEntryDataType](#CompensationLengthofServiceSeveranceMatrixEntryDataType)
* [type CompensationMatrixAmountBasedDataType](#CompensationMatrixAmountBasedDataType)
* [type CompensationMatrixDataType](#CompensationMatrixDataType)
  * [func (t *CompensationMatrixDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#CompensationMatrixDataType.MarshalXML)
  * [func (t *CompensationMatrixDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#CompensationMatrixDataType.UnmarshalXML)
* [type CompensationMatrixEntryAmountBasedDataType](#CompensationMatrixEntryAmountBasedDataType)
* [type CompensationMatrixEntryPercentBasedDataType](#CompensationMatrixEntryPercentBasedDataType)
* [type CompensationMatrixNonweightedPercentBasedDataType](#CompensationMatrixNonweightedPercentBasedDataType)
* [type CompensationMatrixObjectIDType](#CompensationMatrixObjectIDType)
* [type CompensationMatrixObjectType](#CompensationMatrixObjectType)
* [type CompensationMatrixRequestCriteriaType](#CompensationMatrixRequestCriteriaType)
* [type CompensationMatrixRequestReferencesType](#CompensationMatrixRequestReferencesType)
* [type CompensationMatrixResponseDataType](#CompensationMatrixResponseDataType)
* [type CompensationMatrixResponseGroupType](#CompensationMatrixResponseGroupType)
* [type CompensationMatrixType](#CompensationMatrixType)
* [type CompensationMatrixWeightedPercentBasedDataType](#CompensationMatrixWeightedPercentBasedDataType)
* [type CompensationPackageObjectIDType](#CompensationPackageObjectIDType)
* [type CompensationPackageObjectType](#CompensationPackageObjectType)
* [type CompensationPayEarningObjectIDType](#CompensationPayEarningObjectIDType)
* [type CompensationPayEarningObjectType](#CompensationPayEarningObjectType)
* [type CompensationPayRangeDataType](#CompensationPayRangeDataType)
* [type CompensationPeriodObjectIDType](#CompensationPeriodObjectIDType)
* [type CompensationPeriodObjectType](#CompensationPeriodObjectType)
* [type CompensationPlanDataType](#CompensationPlanDataType)
  * [func (t *CompensationPlanDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#CompensationPlanDataType.MarshalXML)
  * [func (t *CompensationPlanDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#CompensationPlanDataType.UnmarshalXML)
* [type CompensationPlanObjectIDType](#CompensationPlanObjectIDType)
* [type CompensationPlanObjectType](#CompensationPlanObjectType)
* [type CompensationPlanRequestCriteriaType](#CompensationPlanRequestCriteriaType)
* [type CompensationPlanRequestReferencesType](#CompensationPlanRequestReferencesType)
* [type CompensationPlanResponseDataType](#CompensationPlanResponseDataType)
* [type CompensationPlanResponseGroupType](#CompensationPlanResponseGroupType)
* [type CompensationPlanType](#CompensationPlanType)
* [type CompensationPreviousSystemHistoryObjectIDType](#CompensationPreviousSystemHistoryObjectIDType)
* [type CompensationPreviousSystemHistoryObjectType](#CompensationPreviousSystemHistoryObjectType)
* [type CompensationRoundingRuleObjectIDType](#CompensationRoundingRuleObjectIDType)
* [type CompensationRoundingRuleObjectType](#CompensationRoundingRuleObjectType)
* [type CompensationScorecardDataType](#CompensationScorecardDataType)
  * [func (t *CompensationScorecardDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#CompensationScorecardDataType.MarshalXML)
  * [func (t *CompensationScorecardDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#CompensationScorecardDataType.UnmarshalXML)
* [type CompensationScorecardRequestCriteriaType](#CompensationScorecardRequestCriteriaType)
* [type CompensationScorecardRequestReferencesType](#CompensationScorecardRequestReferencesType)
* [type CompensationScorecardResponseDataType](#CompensationScorecardResponseDataType)
* [type CompensationScorecardResponseGroupType](#CompensationScorecardResponseGroupType)
* [type CompensationScorecardResultDataType](#CompensationScorecardResultDataType)
  * [func (t *CompensationScorecardResultDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#CompensationScorecardResultDataType.MarshalXML)
  * [func (t *CompensationScorecardResultDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#CompensationScorecardResultDataType.UnmarshalXML)
* [type CompensationScorecardResultRequestCriteriaType](#CompensationScorecardResultRequestCriteriaType)
* [type CompensationScorecardResultRequestReferencesType](#CompensationScorecardResultRequestReferencesType)
* [type CompensationScorecardResultResponseDataType](#CompensationScorecardResultResponseDataType)
* [type CompensationScorecardResultResponseGroupType](#CompensationScorecardResultResponseGroupType)
* [type CompensationScorecardResultType](#CompensationScorecardResultType)
* [type CompensationScorecardType](#CompensationScorecardType)
* [type CompensationSetupSecuritySegmentObjectIDType](#CompensationSetupSecuritySegmentObjectIDType)
* [type CompensationSetupSecuritySegmentObjectType](#CompensationSetupSecuritySegmentObjectType)
* [type CompensationStepDataType](#CompensationStepDataType)
* [type CompensationStepObjectIDType](#CompensationStepObjectIDType)
* [type CompensationStepObjectType](#CompensationStepObjectType)
* [type CompensationStepType](#CompensationStepType)
* [type CompensationTrancheReplacementDataType](#CompensationTrancheReplacementDataType)
* [type CompensationWeightedPercentMatrixObjectIDType](#CompensationWeightedPercentMatrixObjectIDType)
* [type CompensationWeightedPercentMatrixObjectType](#CompensationWeightedPercentMatrixObjectType)
* [type CompplObjectIDType](#CompplObjectIDType)
* [type CompplObjectType](#CompplObjectType)
* [type ConditionEntryOptionObjectIDType](#ConditionEntryOptionObjectIDType)
* [type ConditionEntryOptionObjectType](#ConditionEntryOptionObjectType)
* [type ConditionItemDataWWSType](#ConditionItemDataWWSType)
  * [func (t *ConditionItemDataWWSType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#ConditionItemDataWWSType.MarshalXML)
  * [func (t *ConditionItemDataWWSType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#ConditionItemDataWWSType.UnmarshalXML)
* [type ConditionRuleCategoryObjectIDType](#ConditionRuleCategoryObjectIDType)
* [type ConditionRuleCategoryObjectType](#ConditionRuleCategoryObjectType)
* [type ConditionRuleDataWWSType](#ConditionRuleDataWWSType)
* [type ConditionRuleObjectIDType](#ConditionRuleObjectIDType)
* [type ConditionRuleObjectType](#ConditionRuleObjectType)
* [type CountryObjectIDType](#CountryObjectIDType)
* [type CountryObjectType](#CountryObjectType)
* [type CreateSeveranceWorksheetInput](#CreateSeveranceWorksheetInput)
* [type CreateSeveranceWorksheetOutput](#CreateSeveranceWorksheetOutput)
* [type CreateSeveranceWorksheetRequestType](#CreateSeveranceWorksheetRequestType)
* [type CreateSeveranceWorksheetResponseType](#CreateSeveranceWorksheetResponseType)
* [type CurrencyObjectIDType](#CurrencyObjectIDType)
* [type CurrencyObjectType](#CurrencyObjectType)
* [type DefaultScorecardObjectIDType](#DefaultScorecardObjectIDType)
* [type DefaultScorecardObjectType](#DefaultScorecardObjectType)
* [type DeferredBonusCalculationObjectIDType](#DeferredBonusCalculationObjectIDType)
* [type DeferredBonusCalculationObjectType](#DeferredBonusCalculationObjectType)
* [type DeferredCompensationDataType](#DeferredCompensationDataType)
* [type DeferredCompensationProfilesDataType](#DeferredCompensationProfilesDataType)
* [type EligibilityWaitingPeriodObjectIDType](#EligibilityWaitingPeriodObjectIDType)
* [type EligibilityWaitingPeriodObjectType](#EligibilityWaitingPeriodObjectType)
* [type EligibleEarningRequestCriteriaType](#EligibleEarningRequestCriteriaType)
* [type EligibleEarningsDataType](#EligibleEarningsDataType)
* [type EligibleEarningsOverrideHVDataType](#EligibleEarningsOverrideHVDataType)
* [type EligibleEarningsOverrideObjectIDType](#EligibleEarningsOverrideObjectIDType)
* [type EligibleEarningsOverrideObjectType](#EligibleEarningsOverrideObjectType)
* [type EligibleEarningsOverridePeriodObjectIDType](#EligibleEarningsOverridePeriodObjectIDType)
* [type EligibleEarningsOverridePeriodObjectType](#EligibleEarningsOverridePeriodObjectType)
* [type EligibleEarningsRequestReferencesType](#EligibleEarningsRequestReferencesType)
* [type EligibleEarningsResponseDataType](#EligibleEarningsResponseDataType)
* [type EligibleEarningsResponseGroupType](#EligibleEarningsResponseGroupType)
* [type EligibleEarningsType](#EligibleEarningsType)
* [type EmployeeMeritAdjustmentDataType](#EmployeeMeritAdjustmentDataType)
* [type EmployeeObjectIDType](#EmployeeObjectIDType)
* [type EmployeeObjectType](#EmployeeObjectType)
* [type EmployeeSeverancePackagePayComponentDataType](#EmployeeSeverancePackagePayComponentDataType)
* [type EmployeeSeveranceWorksheetDataType](#EmployeeSeveranceWorksheetDataType)
  * [func (t *EmployeeSeveranceWorksheetDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#EmployeeSeveranceWorksheetDataType.MarshalXML)
  * [func (t *EmployeeSeveranceWorksheetDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#EmployeeSeveranceWorksheetDataType.UnmarshalXML)
* [type EmployeeSeveranceWorksheetEventDataType](#EmployeeSeveranceWorksheetEventDataType)
  * [func (t *EmployeeSeveranceWorksheetEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#EmployeeSeveranceWorksheetEventDataType.MarshalXML)
  * [func (t *EmployeeSeveranceWorksheetEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#EmployeeSeveranceWorksheetEventDataType.UnmarshalXML)
* [type EmployeeSeveranceWorksheetEventObjectIDType](#EmployeeSeveranceWorksheetEventObjectIDType)
* [type EmployeeSeveranceWorksheetEventObjectType](#EmployeeSeveranceWorksheetEventObjectType)
* [type EmployeeSeveranceWorksheetEventRequestCriteriaType](#EmployeeSeveranceWorksheetEventRequestCriteriaType)
* [type EmployeeSeveranceWorksheetEventRequestReferencesType](#EmployeeSeveranceWorksheetEventRequestReferencesType)
* [type EmployeeSeveranceWorksheetEventResponseDataType](#EmployeeSeveranceWorksheetEventResponseDataType)
* [type EmployeeSeveranceWorksheetEventResponseGroupType](#EmployeeSeveranceWorksheetEventResponseGroupType)
* [type EmployeeSeveranceWorksheetEventType](#EmployeeSeveranceWorksheetEventType)
* [type EnhancedSeveranceMatrixDataType](#EnhancedSeveranceMatrixDataType)
  * [func (t *EnhancedSeveranceMatrixDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#EnhancedSeveranceMatrixDataType.MarshalXML)
  * [func (t *EnhancedSeveranceMatrixDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#EnhancedSeveranceMatrixDataType.UnmarshalXML)
* [type EnhancedSeveranceMatrixRequestCriteriaType](#EnhancedSeveranceMatrixRequestCriteriaType)
* [type EnhancedSeveranceMatrixRequestReferencesType](#EnhancedSeveranceMatrixRequestReferencesType)
* [type EnhancedSeveranceMatrixResponseDataType](#EnhancedSeveranceMatrixResponseDataType)
* [type EnhancedSeveranceMatrixResponseGroupType](#EnhancedSeveranceMatrixResponseGroupType)
* [type EnhancedSeveranceMatrixType](#EnhancedSeveranceMatrixType)
* [type EventAttachmentCategoryObjectIDType](#EventAttachmentCategoryObjectIDType)
* [type EventAttachmentCategoryObjectType](#EventAttachmentCategoryObjectType)
* [type EventClassificationSubcategoryObjectIDType](#EventClassificationSubcategoryObjectIDType)
* [type EventClassificationSubcategoryObjectType](#EventClassificationSubcategoryObjectType)
* [type ExpenseAccumulatorObjectIDType](#ExpenseAccumulatorObjectIDType)
* [type ExpenseAccumulatorObjectType](#ExpenseAccumulatorObjectType)
* [type ExpenseItemObjectIDType](#ExpenseItemObjectIDType)
* [type ExpenseItemObjectType](#ExpenseItemObjectType)
* [type ExternalFieldObjectIDType](#ExternalFieldObjectIDType)
* [type ExternalFieldObjectType](#ExternalFieldObjectType)
* [type FilterDateTimeZoneDataType](#FilterDateTimeZoneDataType)
  * [func (t *FilterDateTimeZoneDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#FilterDateTimeZoneDataType.MarshalXML)
  * [func (t *FilterDateTimeZoneDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#FilterDateTimeZoneDataType.UnmarshalXML)
* [type FinancialsBusinessSubProcessParametersType](#FinancialsBusinessSubProcessParametersType)
* [type FrequencyBehaviorObjectIDType](#FrequencyBehaviorObjectIDType)
* [type FrequencyBehaviorObjectType](#FrequencyBehaviorObjectType)
* [type FrequencyObjectIDType](#FrequencyObjectIDType)
* [type FrequencyObjectType](#FrequencyObjectType)
* [type FuturePaymentPlanAssignmentDataType](#FuturePaymentPlanAssignmentDataType)
  * [func (t *FuturePaymentPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#FuturePaymentPlanAssignmentDataType.MarshalXML)
  * [func (t *FuturePaymentPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#FuturePaymentPlanAssignmentDataType.UnmarshalXML)
* [type FuturePaymentPlanAssignmentRequestCriteriaType](#FuturePaymentPlanAssignmentRequestCriteriaType)
* [type FuturePaymentPlanAssignmentResponseDataType](#FuturePaymentPlanAssignmentResponseDataType)
* [type FuturePaymentPlanDataType](#FuturePaymentPlanDataType)
* [type FuturePaymentPlanObjectIDType](#FuturePaymentPlanObjectIDType)
* [type FuturePaymentPlanObjectType](#FuturePaymentPlanObjectType)
* [type GetBenchmarkJobsInput](#GetBenchmarkJobsInput)
* [type GetBenchmarkJobsOutput](#GetBenchmarkJobsOutput)
* [type GetBenchmarkJobsRequestType](#GetBenchmarkJobsRequestType)
* [type GetBenchmarkJobsResponseType](#GetBenchmarkJobsResponseType)
* [type GetCompensationEligibilityRulesInput](#GetCompensationEligibilityRulesInput)
* [type GetCompensationEligibilityRulesOutput](#GetCompensationEligibilityRulesOutput)
* [type GetCompensationEligibilityRulesRequestType](#GetCompensationEligibilityRulesRequestType)
* [type GetCompensationEligibilityRulesResponseType](#GetCompensationEligibilityRulesResponseType)
* [type GetCompensationEligibilityRuleswithoutDependenciesInput](#GetCompensationEligibilityRuleswithoutDependenciesInput)
* [type GetCompensationEligibilityRuleswithoutDependenciesOutput](#GetCompensationEligibilityRuleswithoutDependenciesOutput)
* [type GetCompensationEligibilityRuleswithoutDependenciesRequestType](#GetCompensationEligibilityRuleswithoutDependenciesRequestType)
* [type GetCompensationGradeHierarchyInput](#GetCompensationGradeHierarchyInput)
* [type GetCompensationGradeHierarchyOutput](#GetCompensationGradeHierarchyOutput)
* [type GetCompensationGradeHierarchyRequestType](#GetCompensationGradeHierarchyRequestType)
* [type GetCompensationGradeHierarchyResponseType](#GetCompensationGradeHierarchyResponseType)
* [type GetCompensationGradesInput](#GetCompensationGradesInput)
* [type GetCompensationGradesOutput](#GetCompensationGradesOutput)
* [type GetCompensationGradesRequestType](#GetCompensationGradesRequestType)
* [type GetCompensationGradesResponseType](#GetCompensationGradesResponseType)
* [type GetCompensationMatricesInput](#GetCompensationMatricesInput)
* [type GetCompensationMatricesOutput](#GetCompensationMatricesOutput)
* [type GetCompensationMatricesRequestType](#GetCompensationMatricesRequestType)
* [type GetCompensationMatricesResponseType](#GetCompensationMatricesResponseType)
* [type GetCompensationPlansInput](#GetCompensationPlansInput)
* [type GetCompensationPlansOutput](#GetCompensationPlansOutput)
* [type GetCompensationPlansRequestType](#GetCompensationPlansRequestType)
* [type GetCompensationPlansResponseType](#GetCompensationPlansResponseType)
* [type GetCompensationScorecardResultsInput](#GetCompensationScorecardResultsInput)
* [type GetCompensationScorecardResultsOutput](#GetCompensationScorecardResultsOutput)
* [type GetCompensationScorecardResultsRequestType](#GetCompensationScorecardResultsRequestType)
* [type GetCompensationScorecardResultsResponseType](#GetCompensationScorecardResultsResponseType)
* [type GetCompensationScorecardsInput](#GetCompensationScorecardsInput)
* [type GetCompensationScorecardsOutput](#GetCompensationScorecardsOutput)
* [type GetCompensationScorecardsRequestType](#GetCompensationScorecardsRequestType)
* [type GetCompensationScorecardsResponseType](#GetCompensationScorecardsResponseType)
* [type GetEligibleEarningsInput](#GetEligibleEarningsInput)
* [type GetEligibleEarningsOutput](#GetEligibleEarningsOutput)
* [type GetEligibleEarningsRequestType](#GetEligibleEarningsRequestType)
* [type GetEligibleEarningsResponseType](#GetEligibleEarningsResponseType)
* [type GetEnhancedSeveranceMatricesInput](#GetEnhancedSeveranceMatricesInput)
* [type GetEnhancedSeveranceMatricesOutput](#GetEnhancedSeveranceMatricesOutput)
* [type GetEnhancedSeveranceMatricesRequestType](#GetEnhancedSeveranceMatricesRequestType)
* [type GetEnhancedSeveranceMatricesResponseType](#GetEnhancedSeveranceMatricesResponseType)
* [type GetFuturePaymentPlanAssignmentsInput](#GetFuturePaymentPlanAssignmentsInput)
* [type GetFuturePaymentPlanAssignmentsOutput](#GetFuturePaymentPlanAssignmentsOutput)
* [type GetFuturePaymentPlanAssignmentsRequestType](#GetFuturePaymentPlanAssignmentsRequestType)
* [type GetFuturePaymentPlanAssignmentsResponseType](#GetFuturePaymentPlanAssignmentsResponseType)
* [type GetOneTimePaymentPlanConfigurableCategoriesInput](#GetOneTimePaymentPlanConfigurableCategoriesInput)
* [type GetOneTimePaymentPlanConfigurableCategoriesOutput](#GetOneTimePaymentPlanConfigurableCategoriesOutput)
* [type GetOneTimePaymentPlanConfigurableCategoriesRequestType](#GetOneTimePaymentPlanConfigurableCategoriesRequestType)
* [type GetOneTimePaymentPlanConfigurableCategoriesResponseType](#GetOneTimePaymentPlanConfigurableCategoriesResponseType)
* [type GetPeriodActivityRateMatricesInput](#GetPeriodActivityRateMatricesInput)
* [type GetPeriodActivityRateMatricesOutput](#GetPeriodActivityRateMatricesOutput)
* [type GetPeriodActivityRateMatricesRequestType](#GetPeriodActivityRateMatricesRequestType)
* [type GetPeriodActivityRateMatricesResponseType](#GetPeriodActivityRateMatricesResponseType)
* [type GetPeriodActivityTasksInput](#GetPeriodActivityTasksInput)
* [type GetPeriodActivityTasksOutput](#GetPeriodActivityTasksOutput)
* [type GetPeriodActivityTasksRequestType](#GetPeriodActivityTasksRequestType)
* [type GetPeriodActivityTasksResponseType](#GetPeriodActivityTasksResponseType)
* [type GetPreviousSystemCompensationHistoryInput](#GetPreviousSystemCompensationHistoryInput)
* [type GetPreviousSystemCompensationHistoryOutput](#GetPreviousSystemCompensationHistoryOutput)
* [type GetPreviousSystemCompensationHistoryRequestType](#GetPreviousSystemCompensationHistoryRequestType)
* [type GetPreviousSystemCompensationHistoryResponseType](#GetPreviousSystemCompensationHistoryResponseType)
* [type GetStockParticipationRateTablesInput](#GetStockParticipationRateTablesInput)
* [type GetStockParticipationRateTablesOutput](#GetStockParticipationRateTablesOutput)
* [type GetStockParticipationRateTablesRequestType](#GetStockParticipationRateTablesRequestType)
* [type GetStockParticipationRateTablesResponseType](#GetStockParticipationRateTablesResponseType)
* [type GrantTypeSplitReplacementDataType](#GrantTypeSplitReplacementDataType)
* [type HourlyPlanDataType](#HourlyPlanDataType)
* [type ImportEligibleEarningsOverrideHVRequestType](#ImportEligibleEarningsOverrideHVRequestType)
* [type ImportEligibleEarningsOverrideInput](#ImportEligibleEarningsOverrideInput)
* [type ImportEligibleEarningsOverrideOutput](#ImportEligibleEarningsOverrideOutput)
* [type ImportEligibleEarningsOverrideRequestType](#ImportEligibleEarningsOverrideRequestType)
* [type ImportRequestCompensationChangeInput](#ImportRequestCompensationChangeInput)
* [type ImportRequestCompensationChangeOutput](#ImportRequestCompensationChangeOutput)
* [type ImportRequestCompensationChangeRequestType](#ImportRequestCompensationChangeRequestType)
* [type InstanceIDType](#InstanceIDType)
* [type InstanceObjectType](#InstanceObjectType)
* [type JobProfileObjectIDType](#JobProfileObjectIDType)
* [type JobProfileObjectType](#JobProfileObjectType)
* [type JobRequisitionObjectIDType](#JobRequisitionObjectIDType)
* [type JobRequisitionObjectType](#JobRequisitionObjectType)
* [type ManagePeriodActivityPayAssignmentsInput](#ManagePeriodActivityPayAssignmentsInput)
* [type ManagePeriodActivityPayAssignmentsOutput](#ManagePeriodActivityPayAssignmentsOutput)
* [type ManagePeriodActivityPayAssignmentsRequestType](#ManagePeriodActivityPayAssignmentsRequestType)
* [type ManagePeriodActivityPayAssignmentsResponseType](#ManagePeriodActivityPayAssignmentsResponseType)
* [type ManagementLevelObjectIDType](#ManagementLevelObjectIDType)
* [type ManagementLevelObjectType](#ManagementLevelObjectType)
* [type MeritPercentPlanObjectIDType](#MeritPercentPlanObjectIDType)
* [type MeritPercentPlanObjectType](#MeritPercentPlanObjectType)
* [type MeritPlanAmountDataType](#MeritPlanAmountDataType)
* [type MeritPlanDataType](#MeritPlanDataType)
* [type MeritPlanObjectIDType](#MeritPlanObjectIDType)
* [type MeritPlanObjectType](#MeritPlanObjectType)
* [type MeritPlanPercentDataType](#MeritPlanPercentDataType)
* [type MeritPlanProfileAmountReplacementDataType](#MeritPlanProfileAmountReplacementDataType)
* [type MeritPlanProfilePercentReplacementDataType](#MeritPlanProfilePercentReplacementDataType)
* [type OneTimePaymentDataType](#OneTimePaymentDataType)
  * [func (t *OneTimePaymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#OneTimePaymentDataType.MarshalXML)
  * [func (t *OneTimePaymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#OneTimePaymentDataType.UnmarshalXML)
* [type OneTimePaymentPlanConfigurableCategoryDataType](#OneTimePaymentPlanConfigurableCategoryDataType)
* [type OneTimePaymentPlanConfigurableCategoryObjectIDType](#OneTimePaymentPlanConfigurableCategoryObjectIDType)
* [type OneTimePaymentPlanConfigurableCategoryObjectType](#OneTimePaymentPlanConfigurableCategoryObjectType)
* [type OneTimePaymentPlanConfigurableCategoryRequestCriteriaType](#OneTimePaymentPlanConfigurableCategoryRequestCriteriaType)
* [type OneTimePaymentPlanConfigurableCategoryRequestReferencesType](#OneTimePaymentPlanConfigurableCategoryRequestReferencesType)
* [type OneTimePaymentPlanConfigurableCategoryResponseDataType](#OneTimePaymentPlanConfigurableCategoryResponseDataType)
* [type OneTimePaymentPlanConfigurableCategoryResponseGroupType](#OneTimePaymentPlanConfigurableCategoryResponseGroupType)
* [type OneTimePaymentPlanConfigurableCategoryType](#OneTimePaymentPlanConfigurableCategoryType)
* [type OneTimePaymentPlanObjectIDType](#OneTimePaymentPlanObjectIDType)
* [type OneTimePaymentPlanObjectType](#OneTimePaymentPlanObjectType)
* [type PayRangeQuartileObjectIDType](#PayRangeQuartileObjectIDType)
* [type PayRangeQuartileObjectType](#PayRangeQuartileObjectType)
* [type PayRangeSegmentTypeObjectIDType](#PayRangeSegmentTypeObjectIDType)
* [type PayRangeSegmentTypeObjectType](#PayRangeSegmentTypeObjectType)
* [type PerformanceCriteriaDataType](#PerformanceCriteriaDataType)
* [type PerformanceCriteriaObjectIDType](#PerformanceCriteriaObjectIDType)
* [type PerformanceCriteriaObjectType](#PerformanceCriteriaObjectType)
* [type PerformanceFactorCompensationMatrixWeightedDataType](#PerformanceFactorCompensationMatrixWeightedDataType)
* [type PerformanceFactorScorecardDataType](#PerformanceFactorScorecardDataType)
* [type PerformanceFactorsDataType](#PerformanceFactorsDataType)
* [type PerformanceMatrixDataType](#PerformanceMatrixDataType)
* [type PeriodActivityAssignmentDeletedDataType](#PeriodActivityAssignmentDeletedDataType)
  * [func (t *PeriodActivityAssignmentDeletedDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#PeriodActivityAssignmentDeletedDataType.MarshalXML)
  * [func (t *PeriodActivityAssignmentDeletedDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#PeriodActivityAssignmentDeletedDataType.UnmarshalXML)
* [type PeriodActivityAssignmentEventDataType](#PeriodActivityAssignmentEventDataType)
  * [func (t *PeriodActivityAssignmentEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#PeriodActivityAssignmentEventDataType.MarshalXML)
  * [func (t *PeriodActivityAssignmentEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#PeriodActivityAssignmentEventDataType.UnmarshalXML)
* [type PeriodActivityAssignmentObjectIDType](#PeriodActivityAssignmentObjectIDType)
* [type PeriodActivityAssignmentObjectType](#PeriodActivityAssignmentObjectType)
* [type PeriodActivityAssignmentVersionDataType](#PeriodActivityAssignmentVersionDataType)
  * [func (t *PeriodActivityAssignmentVersionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#PeriodActivityAssignmentVersionDataType.MarshalXML)
  * [func (t *PeriodActivityAssignmentVersionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#PeriodActivityAssignmentVersionDataType.UnmarshalXML)
* [type PeriodActivityCategoryObjectIDType](#PeriodActivityCategoryObjectIDType)
* [type PeriodActivityCategoryObjectType](#PeriodActivityCategoryObjectType)
* [type PeriodActivityEligibilityEntryDataType](#PeriodActivityEligibilityEntryDataType)
* [type PeriodActivityObjectIDType](#PeriodActivityObjectIDType)
* [type PeriodActivityObjectType](#PeriodActivityObjectType)
* [type PeriodActivityPayAssignmentDataType](#PeriodActivityPayAssignmentDataType)
  * [func (t *PeriodActivityPayAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#PeriodActivityPayAssignmentDataType.MarshalXML)
  * [func (t *PeriodActivityPayAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#PeriodActivityPayAssignmentDataType.UnmarshalXML)
* [type PeriodActivityPayAssignmentsDataType](#PeriodActivityPayAssignmentsDataType)
  * [func (t *PeriodActivityPayAssignmentsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#PeriodActivityPayAssignmentsDataType.MarshalXML)
  * [func (t *PeriodActivityPayAssignmentsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#PeriodActivityPayAssignmentsDataType.UnmarshalXML)
* [type PeriodActivityPayCostingDataType](#PeriodActivityPayCostingDataType)
* [type PeriodActivityRateMatrixDataType](#PeriodActivityRateMatrixDataType)
  * [func (t *PeriodActivityRateMatrixDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#PeriodActivityRateMatrixDataType.MarshalXML)
  * [func (t *PeriodActivityRateMatrixDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#PeriodActivityRateMatrixDataType.UnmarshalXML)
* [type PeriodActivityRateMatrixEntryDataType](#PeriodActivityRateMatrixEntryDataType)
* [type PeriodActivityRateMatrixEntrySequenceDataType](#PeriodActivityRateMatrixEntrySequenceDataType)
* [type PeriodActivityRateMatrixObjectIDType](#PeriodActivityRateMatrixObjectIDType)
* [type PeriodActivityRateMatrixObjectType](#PeriodActivityRateMatrixObjectType)
* [type PeriodActivityRateMatrixRequestCriteriaType](#PeriodActivityRateMatrixRequestCriteriaType)
* [type PeriodActivityRateMatrixRequestReferencesType](#PeriodActivityRateMatrixRequestReferencesType)
* [type PeriodActivityRateMatrixResponseDataType](#PeriodActivityRateMatrixResponseDataType)
* [type PeriodActivityRateMatrixType](#PeriodActivityRateMatrixType)
* [type PeriodActivityTaskDataType](#PeriodActivityTaskDataType)
* [type PeriodActivityTaskInterfaceObjectIDType](#PeriodActivityTaskInterfaceObjectIDType)
* [type PeriodActivityTaskInterfaceObjectType](#PeriodActivityTaskInterfaceObjectType)
* [type PeriodActivityTaskObjectIDType](#PeriodActivityTaskObjectIDType)
* [type PeriodActivityTaskObjectType](#PeriodActivityTaskObjectType)
* [type PeriodActivityTaskRequestCriteriaType](#PeriodActivityTaskRequestCriteriaType)
* [type PeriodActivityTaskRequestReferencesType](#PeriodActivityTaskRequestReferencesType)
* [type PeriodActivityTaskResponseDataType](#PeriodActivityTaskResponseDataType)
* [type PeriodActivityTaskResponseGroupType](#PeriodActivityTaskResponseGroupType)
* [type PeriodActivityTaskType](#PeriodActivityTaskType)
* [type PeriodActivityUnitObjectIDType](#PeriodActivityUnitObjectIDType)
* [type PeriodActivityUnitObjectType](#PeriodActivityUnitObjectType)
* [type PeriodPlanDataType](#PeriodPlanDataType)
* [type PeriodPlanProfileReplacementDataType](#PeriodPlanProfileReplacementDataType)
* [type PeriodSalaryPlanObjectIDType](#PeriodSalaryPlanObjectIDType)
* [type PeriodSalaryPlanObjectType](#PeriodSalaryPlanObjectType)
* [type PositionElementObjectIDType](#PositionElementObjectIDType)
* [type PositionElementObjectType](#PositionElementObjectType)
* [type PositionFuturePaymentPlanAssignmentDataType](#PositionFuturePaymentPlanAssignmentDataType)
  * [func (t *PositionFuturePaymentPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#PositionFuturePaymentPlanAssignmentDataType.MarshalXML)
  * [func (t *PositionFuturePaymentPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#PositionFuturePaymentPlanAssignmentDataType.UnmarshalXML)
* [type PositionFuturePaymentPlanAssignmentsType](#PositionFuturePaymentPlanAssignmentsType)
* [type PositionObjectIDType](#PositionObjectIDType)
* [type PositionObjectType](#PositionObjectType)
* [type PotentialObjectIDType](#PotentialObjectIDType)
* [type PotentialObjectType](#PotentialObjectType)
* [type PreviousSystemCompensationHistoryDataType](#PreviousSystemCompensationHistoryDataType)
* [type PreviousSystemCompensationHistoryDetailDataType](#PreviousSystemCompensationHistoryDetailDataType)
  * [func (t *PreviousSystemCompensationHistoryDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#PreviousSystemCompensationHistoryDetailDataType.MarshalXML)
  * [func (t *PreviousSystemCompensationHistoryDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#PreviousSystemCompensationHistoryDetailDataType.UnmarshalXML)
* [type PreviousSystemCompensationHistoryGetDataType](#PreviousSystemCompensationHistoryGetDataType)
* [type PreviousSystemCompensationHistoryResponseDataType](#PreviousSystemCompensationHistoryResponseDataType)
* [type PreviousSystemCompensationHistoryType](#PreviousSystemCompensationHistoryType)
* [type ProcessingFaultType](#ProcessingFaultType)
* [type ProfileCompensationScorecardResultDataType](#ProfileCompensationScorecardResultDataType)
* [type ProfileScorecardResultDataType](#ProfileScorecardResultDataType)
* [type ProposedAllowancePlanAssignmentContainerDataType](#ProposedAllowancePlanAssignmentContainerDataType)
* [type ProposedAllowancePlanAssignmentDataType](#ProposedAllowancePlanAssignmentDataType)
  * [func (t *ProposedAllowancePlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#ProposedAllowancePlanAssignmentDataType.MarshalXML)
  * [func (t *ProposedAllowancePlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#ProposedAllowancePlanAssignmentDataType.UnmarshalXML)
* [type ProposedAllowanceUnitPlanAssignmentContainerDataType](#ProposedAllowanceUnitPlanAssignmentContainerDataType)
* [type ProposedAllowanceUnitPlanAssignmentDataType](#ProposedAllowanceUnitPlanAssignmentDataType)
  * [func (t *ProposedAllowanceUnitPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#ProposedAllowanceUnitPlanAssignmentDataType.MarshalXML)
  * [func (t *ProposedAllowanceUnitPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#ProposedAllowanceUnitPlanAssignmentDataType.UnmarshalXML)
* [type ProposedBasePayPlanAssignmentContainerDataType](#ProposedBasePayPlanAssignmentContainerDataType)
* [type ProposedBasePayPlanAssignmentDataType](#ProposedBasePayPlanAssignmentDataType)
  * [func (t *ProposedBasePayPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#ProposedBasePayPlanAssignmentDataType.MarshalXML)
  * [func (t *ProposedBasePayPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#ProposedBasePayPlanAssignmentDataType.UnmarshalXML)
* [type ProposedBonusPlanAssignmentContainerDataType](#ProposedBonusPlanAssignmentContainerDataType)
* [type ProposedBonusPlanAssignmentDataType](#ProposedBonusPlanAssignmentDataType)
  * [func (t *ProposedBonusPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#ProposedBonusPlanAssignmentDataType.MarshalXML)
  * [func (t *ProposedBonusPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#ProposedBonusPlanAssignmentDataType.UnmarshalXML)
* [type ProposedCalculatedPlanAssignmentContainerDataType](#ProposedCalculatedPlanAssignmentContainerDataType)
* [type ProposedCalculatedPlanAssignmentDataType](#ProposedCalculatedPlanAssignmentDataType)
  * [func (t *ProposedCalculatedPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#ProposedCalculatedPlanAssignmentDataType.MarshalXML)
  * [func (t *ProposedCalculatedPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#ProposedCalculatedPlanAssignmentDataType.UnmarshalXML)
* [type ProposedCommissionPlanAssignmentContainerDataType](#ProposedCommissionPlanAssignmentContainerDataType)
* [type ProposedCommissionPlanAssignmentDataType](#ProposedCommissionPlanAssignmentDataType)
  * [func (t *ProposedCommissionPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#ProposedCommissionPlanAssignmentDataType.MarshalXML)
  * [func (t *ProposedCommissionPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#ProposedCommissionPlanAssignmentDataType.UnmarshalXML)
* [type ProposedMeritPlanAssignmentContainerDataType](#ProposedMeritPlanAssignmentContainerDataType)
* [type ProposedMeritPlanAssignmentDataType](#ProposedMeritPlanAssignmentDataType)
  * [func (t *ProposedMeritPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#ProposedMeritPlanAssignmentDataType.MarshalXML)
  * [func (t *ProposedMeritPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#ProposedMeritPlanAssignmentDataType.UnmarshalXML)
* [type ProposedPeriodSalaryPlanAssignmentContainerDataType](#ProposedPeriodSalaryPlanAssignmentContainerDataType)
* [type ProposedPeriodSalaryPlanAssignmentDataType](#ProposedPeriodSalaryPlanAssignmentDataType)
  * [func (t *ProposedPeriodSalaryPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#ProposedPeriodSalaryPlanAssignmentDataType.MarshalXML)
  * [func (t *ProposedPeriodSalaryPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#ProposedPeriodSalaryPlanAssignmentDataType.UnmarshalXML)
* [type ProposedSalaryUnitPlanAssignmentContainerDataType](#ProposedSalaryUnitPlanAssignmentContainerDataType)
* [type ProposedSalaryUnitPlanAssignmentDataType](#ProposedSalaryUnitPlanAssignmentDataType)
  * [func (t *ProposedSalaryUnitPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#ProposedSalaryUnitPlanAssignmentDataType.MarshalXML)
  * [func (t *ProposedSalaryUnitPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#ProposedSalaryUnitPlanAssignmentDataType.UnmarshalXML)
* [type ProposedStockPlanAssignmentContainerDataType](#ProposedStockPlanAssignmentContainerDataType)
* [type ProposedStockPlanAssignmentDataType](#ProposedStockPlanAssignmentDataType)
  * [func (t *ProposedStockPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#ProposedStockPlanAssignmentDataType.MarshalXML)
  * [func (t *ProposedStockPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#ProposedStockPlanAssignmentDataType.UnmarshalXML)
* [type PutBenchmarkJobInput](#PutBenchmarkJobInput)
* [type PutBenchmarkJobOutput](#PutBenchmarkJobOutput)
* [type PutBenchmarkJobRequestType](#PutBenchmarkJobRequestType)
* [type PutBenchmarkJobResponseType](#PutBenchmarkJobResponseType)
* [type PutCompensationEligibilityRuleInput](#PutCompensationEligibilityRuleInput)
* [type PutCompensationEligibilityRuleOutput](#PutCompensationEligibilityRuleOutput)
* [type PutCompensationEligibilityRuleRequestType](#PutCompensationEligibilityRuleRequestType)
* [type PutCompensationEligibilityRuleResponseType](#PutCompensationEligibilityRuleResponseType)
* [type PutCompensationGradeHierarchyInput](#PutCompensationGradeHierarchyInput)
* [type PutCompensationGradeHierarchyOutput](#PutCompensationGradeHierarchyOutput)
* [type PutCompensationGradeHierarchyRequestType](#PutCompensationGradeHierarchyRequestType)
* [type PutCompensationGradeHierarchyResponseType](#PutCompensationGradeHierarchyResponseType)
  * [func (t *PutCompensationGradeHierarchyResponseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#PutCompensationGradeHierarchyResponseType.MarshalXML)
  * [func (t *PutCompensationGradeHierarchyResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#PutCompensationGradeHierarchyResponseType.UnmarshalXML)
* [type PutCompensationGradeInput](#PutCompensationGradeInput)
* [type PutCompensationGradeLevelResponseType](#PutCompensationGradeLevelResponseType)
* [type PutCompensationGradeOutput](#PutCompensationGradeOutput)
* [type PutCompensationGradeRequestType](#PutCompensationGradeRequestType)
* [type PutCompensationGradeResponseType](#PutCompensationGradeResponseType)
* [type PutCompensationMatrixInput](#PutCompensationMatrixInput)
* [type PutCompensationMatrixOutput](#PutCompensationMatrixOutput)
* [type PutCompensationMatrixRequestType](#PutCompensationMatrixRequestType)
* [type PutCompensationMatrixResponseType](#PutCompensationMatrixResponseType)
* [type PutCompensationPlanRequestType](#PutCompensationPlanRequestType)
* [type PutCompensationPlanResponseType](#PutCompensationPlanResponseType)
* [type PutCompensationPlansInput](#PutCompensationPlansInput)
* [type PutCompensationPlansOutput](#PutCompensationPlansOutput)
* [type PutCompensationScorecardInput](#PutCompensationScorecardInput)
* [type PutCompensationScorecardOutput](#PutCompensationScorecardOutput)
* [type PutCompensationScorecardRequestType](#PutCompensationScorecardRequestType)
* [type PutCompensationScorecardResponseType](#PutCompensationScorecardResponseType)
* [type PutCompensationScorecardResultInput](#PutCompensationScorecardResultInput)
* [type PutCompensationScorecardResultOutput](#PutCompensationScorecardResultOutput)
* [type PutCompensationScorecardResultRequestType](#PutCompensationScorecardResultRequestType)
* [type PutCompensationScorecardResultResponseType](#PutCompensationScorecardResultResponseType)
* [type PutEligibleEarningsInput](#PutEligibleEarningsInput)
* [type PutEligibleEarningsOutput](#PutEligibleEarningsOutput)
* [type PutEligibleEarningsRequestType](#PutEligibleEarningsRequestType)
* [type PutEligibleEarningsResponseType](#PutEligibleEarningsResponseType)
* [type PutEnhancedSeveranceMatrixInput](#PutEnhancedSeveranceMatrixInput)
* [type PutEnhancedSeveranceMatrixOutput](#PutEnhancedSeveranceMatrixOutput)
* [type PutEnhancedSeveranceMatrixRequestType](#PutEnhancedSeveranceMatrixRequestType)
* [type PutEnhancedSeveranceMatrixResponseType](#PutEnhancedSeveranceMatrixResponseType)
* [type PutFuturePaymentPlanAssignmentInput](#PutFuturePaymentPlanAssignmentInput)
* [type PutFuturePaymentPlanAssignmentOutput](#PutFuturePaymentPlanAssignmentOutput)
* [type PutFuturePaymentPlanAssignmentRequestType](#PutFuturePaymentPlanAssignmentRequestType)
* [type PutFuturePaymentPlanAssignmentResponseType](#PutFuturePaymentPlanAssignmentResponseType)
  * [func (t *PutFuturePaymentPlanAssignmentResponseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#PutFuturePaymentPlanAssignmentResponseType.MarshalXML)
  * [func (t *PutFuturePaymentPlanAssignmentResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#PutFuturePaymentPlanAssignmentResponseType.UnmarshalXML)
* [type PutImportProcessResponseType](#PutImportProcessResponseType)
* [type PutOneTimePaymentPlanConfigurableCategoryInput](#PutOneTimePaymentPlanConfigurableCategoryInput)
* [type PutOneTimePaymentPlanConfigurableCategoryOutput](#PutOneTimePaymentPlanConfigurableCategoryOutput)
* [type PutOneTimePaymentPlanConfigurableCategoryRequestType](#PutOneTimePaymentPlanConfigurableCategoryRequestType)
* [type PutOneTimePaymentPlanConfigurableCategoryResponseType](#PutOneTimePaymentPlanConfigurableCategoryResponseType)
* [type PutPeriodActivityRateMatrixInput](#PutPeriodActivityRateMatrixInput)
* [type PutPeriodActivityRateMatrixOutput](#PutPeriodActivityRateMatrixOutput)
* [type PutPeriodActivityRateMatrixRequestType](#PutPeriodActivityRateMatrixRequestType)
* [type PutPeriodActivityRateMatrixResponseType](#PutPeriodActivityRateMatrixResponseType)
* [type PutPeriodActivityTaskInput](#PutPeriodActivityTaskInput)
* [type PutPeriodActivityTaskOutput](#PutPeriodActivityTaskOutput)
* [type PutPeriodActivityTaskRequestType](#PutPeriodActivityTaskRequestType)
* [type PutPeriodActivityTaskResponseType](#PutPeriodActivityTaskResponseType)
* [type PutPreviousSystemCompensationHistoryInput](#PutPreviousSystemCompensationHistoryInput)
* [type PutPreviousSystemCompensationHistoryOutput](#PutPreviousSystemCompensationHistoryOutput)
* [type PutPreviousSystemCompensationHistoryRequestType](#PutPreviousSystemCompensationHistoryRequestType)
* [type PutPreviousSystemCompensationHistoryResponseType](#PutPreviousSystemCompensationHistoryResponseType)
* [type PutStockParticipationRateTableInput](#PutStockParticipationRateTableInput)
* [type PutStockParticipationRateTableOutput](#PutStockParticipationRateTableOutput)
* [type PutStockParticipationRateTableRequestType](#PutStockParticipationRateTableRequestType)
* [type PutStockParticipationRateTableResponseType](#PutStockParticipationRateTableResponseType)
* [type RelationalOperatorObjectIDType](#RelationalOperatorObjectIDType)
* [type RelationalOperatorObjectType](#RelationalOperatorObjectType)
* [type RemoveCompensationPlanAssignmentDataType](#RemoveCompensationPlanAssignmentDataType)
* [type RequestBonusPaymentDataType](#RequestBonusPaymentDataType)
  * [func (t *RequestBonusPaymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#RequestBonusPaymentDataType.MarshalXML)
  * [func (t *RequestBonusPaymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#RequestBonusPaymentDataType.UnmarshalXML)
* [type RequestBonusPaymentInput](#RequestBonusPaymentInput)
* [type RequestBonusPaymentOutput](#RequestBonusPaymentOutput)
* [type RequestBonusPaymentRequestType](#RequestBonusPaymentRequestType)
* [type RequestBonusPaymentResponseType](#RequestBonusPaymentResponseType)
* [type RequestCompensationChangeDataType](#RequestCompensationChangeDataType)
  * [func (t *RequestCompensationChangeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#RequestCompensationChangeDataType.MarshalXML)
  * [func (t *RequestCompensationChangeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#RequestCompensationChangeDataType.UnmarshalXML)
* [type RequestCompensationChangeInput](#RequestCompensationChangeInput)
* [type RequestCompensationChangeOutput](#RequestCompensationChangeOutput)
* [type RequestCompensationChangeRequestHVType](#RequestCompensationChangeRequestHVType)
* [type RequestCompensationChangeRequestType](#RequestCompensationChangeRequestType)
* [type RequestCompensationChangeResponseType](#RequestCompensationChangeResponseType)
* [type RequestEmployeeMeritAdjustmentDataType](#RequestEmployeeMeritAdjustmentDataType)
  * [func (t *RequestEmployeeMeritAdjustmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#RequestEmployeeMeritAdjustmentDataType.MarshalXML)
  * [func (t *RequestEmployeeMeritAdjustmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#RequestEmployeeMeritAdjustmentDataType.UnmarshalXML)
* [type RequestEmployeeMeritAdjustmentInput](#RequestEmployeeMeritAdjustmentInput)
* [type RequestEmployeeMeritAdjustmentOutput](#RequestEmployeeMeritAdjustmentOutput)
* [type RequestEmployeeMeritAdjustmentRequestType](#RequestEmployeeMeritAdjustmentRequestType)
* [type RequestEmployeeMeritAdjustmentResponseType](#RequestEmployeeMeritAdjustmentResponseType)
* [type RequestOneTimePaymentDataType](#RequestOneTimePaymentDataType)
  * [func (t *RequestOneTimePaymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#RequestOneTimePaymentDataType.MarshalXML)
  * [func (t *RequestOneTimePaymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#RequestOneTimePaymentDataType.UnmarshalXML)
* [type RequestOneTimePaymentInput](#RequestOneTimePaymentInput)
* [type RequestOneTimePaymentOutput](#RequestOneTimePaymentOutput)
* [type RequestOneTimePaymentRequestType](#RequestOneTimePaymentRequestType)
* [type RequestOneTimePaymentResponseType](#RequestOneTimePaymentResponseType)
* [type RequestRequisitionCompensationChangeInput](#RequestRequisitionCompensationChangeInput)
* [type RequestRequisitionCompensationChangeOutput](#RequestRequisitionCompensationChangeOutput)
* [type RequestRequisitionCompensationChangeRequestType](#RequestRequisitionCompensationChangeRequestType)
* [type RequestRequisitionCompensationChangeResponseType](#RequestRequisitionCompensationChangeResponseType)
* [type RequisitionCompensationChangeDataType](#RequisitionCompensationChangeDataType)
  * [func (t *RequisitionCompensationChangeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#RequisitionCompensationChangeDataType.MarshalXML)
  * [func (t *RequisitionCompensationChangeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#RequisitionCompensationChangeDataType.UnmarshalXML)
* [type RequisitionCompensationDataType](#RequisitionCompensationDataType)
* [type ResponseFilterType](#ResponseFilterType)
  * [func (t *ResponseFilterType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#ResponseFilterType.MarshalXML)
  * [func (t *ResponseFilterType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#ResponseFilterType.UnmarshalXML)
* [type ResponseResultsType](#ResponseResultsType)
* [type RetentionObjectIDType](#RetentionObjectIDType)
* [type RetentionObjectType](#RetentionObjectType)
* [type RetrieveSeveranceWorksheetInput](#RetrieveSeveranceWorksheetInput)
* [type RetrieveSeveranceWorksheetOutput](#RetrieveSeveranceWorksheetOutput)
* [type RetrieveSeveranceWorksheetsRequestType](#RetrieveSeveranceWorksheetsRequestType)
* [type RetrieveSeveranceWorksheetsResponseType](#RetrieveSeveranceWorksheetsResponseType)
* [type ReviewRatingObjectIDType](#ReviewRatingObjectIDType)
* [type ReviewRatingObjectType](#ReviewRatingObjectType)
* [type ReviewRatingScaleObjectIDType](#ReviewRatingScaleObjectIDType)
* [type ReviewRatingScaleObjectType](#ReviewRatingScaleObjectType)
* [type SalaryPayPlanObjectIDType](#SalaryPayPlanObjectIDType)
* [type SalaryPayPlanObjectType](#SalaryPayPlanObjectType)
* [type SalaryPlanDataType](#SalaryPlanDataType)
* [type SalaryPlanObjectIDType](#SalaryPlanObjectIDType)
* [type SalaryPlanObjectType](#SalaryPlanObjectType)
* [type SalaryUnitPlanObjectIDType](#SalaryUnitPlanObjectIDType)
* [type SalaryUnitPlanObjectType](#SalaryUnitPlanObjectType)
* [type ScorecardProfileDataType](#ScorecardProfileDataType)
* [type ScorecardProfilePerformanceCriteriaDataType](#ScorecardProfilePerformanceCriteriaDataType)
* [type ScorecardResultDataType](#ScorecardResultDataType)
* [type ScoresetContainerObjectIDType](#ScoresetContainerObjectIDType)
* [type ScoresetContainerObjectType](#ScoresetContainerObjectType)
* [type ServiceDurationSeveranceMatrixDataType](#ServiceDurationSeveranceMatrixDataType)
* [type ServiceDurationSeveranceMatrixEntryDataType](#ServiceDurationSeveranceMatrixEntryDataType)
* [type ServiceLengthSeveranceMatrixDataType](#ServiceLengthSeveranceMatrixDataType)
* [type ServiceLengthSeveranceMatrixEntryDataType](#ServiceLengthSeveranceMatrixEntryDataType)
* [type SeveranceMatrixFilterOptionsType](#SeveranceMatrixFilterOptionsType)
* [type SeveranceMatrixabstractObjectIDType](#SeveranceMatrixabstractObjectIDType)
* [type SeveranceMatrixabstractObjectType](#SeveranceMatrixabstractObjectType)
* [type SeverancePackageComponentDataType](#SeverancePackageComponentDataType)
* [type SeverancePackageComponentTypeObjectIDType](#SeverancePackageComponentTypeObjectIDType)
* [type SeverancePackageComponentTypeObjectType](#SeverancePackageComponentTypeObjectType)
* [type SeverancePackageObjectIDType](#SeverancePackageObjectIDType)
* [type SeverancePackageObjectType](#SeverancePackageObjectType)
* [type SeverancePayComponentDataType](#SeverancePayComponentDataType)
* [type SeveranceServiceDateObjectIDType](#SeveranceServiceDateObjectIDType)
* [type SeveranceServiceDateObjectType](#SeveranceServiceDateObjectType)
* [type StockDateRuleObjectIDType](#StockDateRuleObjectIDType)
* [type StockDateRuleObjectType](#StockDateRuleObjectType)
* [type StockGrantObjectIDType](#StockGrantObjectIDType)
* [type StockGrantObjectType](#StockGrantObjectType)
* [type StockGrantTypeObjectIDType](#StockGrantTypeObjectIDType)
* [type StockGrantTypeObjectType](#StockGrantTypeObjectType)
* [type StockParticipationRateTableData20Type](#StockParticipationRateTableData20Type)
* [type StockParticipationRateTableEntryData20Type](#StockParticipationRateTableEntryData20Type)
* [type StockParticipationRateTableObjectIDType](#StockParticipationRateTableObjectIDType)
* [type StockParticipationRateTableObjectType](#StockParticipationRateTableObjectType)
* [type StockParticipationRateTableRequestCriteriaType](#StockParticipationRateTableRequestCriteriaType)
* [type StockParticipationRateTableRequestReferencesType](#StockParticipationRateTableRequestReferencesType)
* [type StockParticipationRateTableResponseDataType](#StockParticipationRateTableResponseDataType)
* [type StockParticipationRateTableResponseGroupType](#StockParticipationRateTableResponseGroupType)
* [type StockParticipationRateTableSubResponseDataType](#StockParticipationRateTableSubResponseDataType)
* [type StockPercentPlanObjectIDType](#StockPercentPlanObjectIDType)
* [type StockPercentPlanObjectType](#StockPercentPlanObjectType)
* [type StockPlanAmountDataType](#StockPlanAmountDataType)
* [type StockPlanAmountProfileReplacementDataType](#StockPlanAmountProfileReplacementDataType)
* [type StockPlanDataType](#StockPlanDataType)
* [type StockPlanObjectIDType](#StockPlanObjectIDType)
* [type StockPlanObjectType](#StockPlanObjectType)
* [type StockPlanPercentDataType](#StockPlanPercentDataType)
* [type StockPlanPercentProfileReplacementDataType](#StockPlanPercentProfileReplacementDataType)
* [type StockPlanUnitDataType](#StockPlanUnitDataType)
* [type StockPlanUnitProfileReplacementDataType](#StockPlanUnitProfileReplacementDataType)
* [type StockProfileGrantSplitReplacementDataType](#StockProfileGrantSplitReplacementDataType)
* [type StockVestingScheduleObjectIDType](#StockVestingScheduleObjectIDType)
* [type StockVestingScheduleObjectType](#StockVestingScheduleObjectType)
* [type TenantedPayrollWorktagObjectIDType](#TenantedPayrollWorktagObjectIDType)
* [type TenantedPayrollWorktagObjectType](#TenantedPayrollWorktagObjectType)
* [type TimeProrationRuleObjectIDType](#TimeProrationRuleObjectIDType)
* [type TimeProrationRuleObjectType](#TimeProrationRuleObjectType)
* [type TimeZoneObjectIDType](#TimeZoneObjectIDType)
* [type TimeZoneObjectType](#TimeZoneObjectType)
* [type UniqueIdentifierObjectIDType](#UniqueIdentifierObjectIDType)
* [type UniqueIdentifierObjectType](#UniqueIdentifierObjectType)
* [type UnitSalaryPlanDataType](#UnitSalaryPlanDataType)
* [type UnitofMeasureObjectIDType](#UnitofMeasureObjectIDType)
* [type UnitofMeasureObjectType](#UnitofMeasureObjectType)
* [type UpdateStockGrantDataType](#UpdateStockGrantDataType)
  * [func (t *UpdateStockGrantDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error](#UpdateStockGrantDataType.MarshalXML)
  * [func (t *UpdateStockGrantDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error](#UpdateStockGrantDataType.UnmarshalXML)
* [type UpdateStockGrantInput](#UpdateStockGrantInput)
* [type UpdateStockGrantOutput](#UpdateStockGrantOutput)
* [type UpdateStockGrantRequestType](#UpdateStockGrantRequestType)
* [type UpdateStockGrantResponseType](#UpdateStockGrantResponseType)
* [type ValidationErrorType](#ValidationErrorType)
* [type ValidationFaultType](#ValidationFaultType)
* [type WebServiceBackgroundProcessRuntimeObjectIDType](#WebServiceBackgroundProcessRuntimeObjectIDType)
* [type WebServiceBackgroundProcessRuntimeObjectType](#WebServiceBackgroundProcessRuntimeObjectType)
* [type WorkdayCommonHeaderType](#WorkdayCommonHeaderType)
* [type WorkerObjectIDType](#WorkerObjectIDType)
* [type WorkerObjectType](#WorkerObjectType)
* [type WorkerRequestReferencesType](#WorkerRequestReferencesType)
* [type WorkerResponseGroupforReferenceType](#WorkerResponseGroupforReferenceType)


#### <a name="pkg-files">Package files</a>
[wwsgen.go](/src/github.com/ianlopshire/go-workday/pkg/wws/services/comp/wwsgen.go) [wwsgen_client.go](/src/github.com/ianlopshire/go-workday/pkg/wws/services/comp/wwsgen_client.go) [wwsgen_types.go](/src/github.com/ianlopshire/go-workday/pkg/wws/services/comp/wwsgen_types.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    CurrentVersion = "v31.2"
    ServiceName    = "Compensation"
)
```




## <a name="AcademicPeriodObjectIDType">type</a> [AcademicPeriodObjectIDType](/src/target/wwsgen_types.go?s=140:267#L11)
``` go
type AcademicPeriodObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="AcademicPeriodObjectType">type</a> [AcademicPeriodObjectType](/src/target/wwsgen_types.go?s=269:488#L16)
``` go
type AcademicPeriodObjectType struct {
    ID         []AcademicPeriodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="AddIndividualStockGrantDataType">type</a> [AddIndividualStockGrantDataType](/src/target/wwsgen_types.go?s=552:3216#L22)
``` go
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
```
Data element for the Request Stock Grant business process.










### <a name="AddIndividualStockGrantDataType.MarshalXML">func</a> (\*AddIndividualStockGrantDataType) [MarshalXML](/src/target/wwsgen_types.go?s=3218:3316#L47)
``` go
func (t *AddIndividualStockGrantDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="AddIndividualStockGrantDataType.UnmarshalXML">func</a> (\*AddIndividualStockGrantDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=3986:4086#L63)
``` go
func (t *AddIndividualStockGrantDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="AddStockGrantDataType">type</a> [AddStockGrantDataType](/src/target/wwsgen_types.go?s=4835:5546#L81)
``` go
type AddStockGrantDataType struct {
    EffectiveDate      time.Time                                 `xml:"urn:com.workday/bsvc Effective_Date"`
    EmployeeReference  *EmployeeObjectType                       `xml:"urn:com.workday/bsvc Employee_Reference"`
    PositionReference  *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
    StockPlanReference *StockPlanObjectType                      `xml:"urn:com.workday/bsvc Stock_Plan_Reference"`
    ReasonReference    *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
    StockGrantData     []AddIndividualStockGrantDataType         `xml:"urn:com.workday/bsvc Stock_Grant_Data"`
}
```
Wrapper element for the Request Stock Grant business process.










### <a name="AddStockGrantDataType.MarshalXML">func</a> (\*AddStockGrantDataType) [MarshalXML](/src/target/wwsgen_types.go?s=5548:5636#L90)
``` go
func (t *AddStockGrantDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="AddStockGrantDataType.UnmarshalXML">func</a> (\*AddStockGrantDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=5888:5978#L100)
``` go
func (t *AddStockGrantDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="AddStockGrantInput">type</a> [AddStockGrantInput](/src/target/wwsgen_client.go?s=18245:18374#L495)
``` go
type AddStockGrantInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Add_Stock_Grant_Request"`
    AddStockGrantRequestType
}
```









## <a name="AddStockGrantOutput">type</a> [AddStockGrantOutput](/src/target/wwsgen_client.go?s=18376:18508#L500)
``` go
type AddStockGrantOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Add_Stock_Grant_Response"`
    AddStockGrantResponseType
}
```









## <a name="AddStockGrantRequestType">type</a> [AddStockGrantRequestType](/src/target/wwsgen_types.go?s=6366:6749#L112)
``` go
type AddStockGrantRequestType struct {
    BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
    AddStockGrantData         *AddStockGrantDataType         `xml:"urn:com.workday/bsvc Add_Stock_Grant_Data"`
    Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
This web service operation is designed to grant a stock grant to an employee using the Request Stock Grant business process.










## <a name="AddStockGrantResponseType">type</a> [AddStockGrantResponseType](/src/target/wwsgen_types.go?s=6816:7090#L119)
``` go
type AddStockGrantResponseType struct {
    StockPackageEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Stock_Package_Event_Reference,omitempty"`
    Version                    string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Responds with the Event ID for the Request Stock Grant Event.










## <a name="AllocationDetailforPeriodPayDataType">type</a> [AllocationDetailforPeriodPayDataType](/src/target/wwsgen_types.go?s=7164:7685#L125)
``` go
type AllocationDetailforPeriodPayDataType struct {
    Order                   string                             `xml:"urn:com.workday/bsvc Order,omitempty"`
    CostingWorktagReference []TenantedPayrollWorktagObjectType `xml:"urn:com.workday/bsvc Costing_Worktag_Reference,omitempty"`
    DistributionPercent     float64                            `xml:"urn:com.workday/bsvc Distribution_Percent,omitempty"`
    DistributionAmount      float64                            `xml:"urn:com.workday/bsvc Distribution_Amount,omitempty"`
}
```
Contains costing allocation for Period Activity Pay assignment line.










## <a name="AllowancePlanAmountBasedProfileReplacementDataType">type</a> [AllowancePlanAmountBasedProfileReplacementDataType](/src/target/wwsgen_types.go?s=7835:8204#L133)
``` go
type AllowancePlanAmountBasedProfileReplacementDataType struct {
    EligibilityRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    Amount                   float64                  `xml:"urn:com.workday/bsvc Amount,omitempty"`
    CurrencyReference        *CurrencyObjectType      `xml:"urn:com.workday/bsvc Currency_Reference"`
}
```
This is a wrapper for the Allowance Plan Amount Based Profile Data; this is used to display or replace Allowance Plan Amount Based Profile Data.










## <a name="AllowancePlanAmountDataType">type</a> [AllowancePlanAmountDataType](/src/target/wwsgen_types.go?s=8246:8629#L140)
``` go
type AllowancePlanAmountDataType struct {
    Amount                                         float64                                              `xml:"urn:com.workday/bsvc Amount,omitempty"`
    AllowancePlanAmountBasedProfileReplacementData []AllowancePlanAmountBasedProfileReplacementDataType `xml:"urn:com.workday/bsvc Allowance_Plan_Amount_Based_Profile_Replacement_Data,omitempty"`
}
```
Contains Allowance Plan Amount Data.










## <a name="AllowancePlanDataType">type</a> [AllowancePlanDataType](/src/target/wwsgen_types.go?s=8822:10319#L146)
``` go
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
```
Allowance Plan Data consists of common information from the following Allowance Plans: Amount, Percent, and Unit.  It also must contain specific Amount, Percent, or Unit data information.










## <a name="AllowancePlanPercentBasedProfileReplacementDataType">type</a> [AllowancePlanPercentBasedProfileReplacementDataType](/src/target/wwsgen_types.go?s=10471:11114#L161)
``` go
type AllowancePlanPercentBasedProfileReplacementDataType struct {
    EligibilityRuleReference   *ConditionRuleObjectType     `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    Percent                    float64                      `xml:"urn:com.workday/bsvc Percent,omitempty"`
    CompensationBasisReference *CompensationBasisObjectType `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
    CeilingAmount              float64                      `xml:"urn:com.workday/bsvc Ceiling_Amount,omitempty"`
    CeilingCurrencyReference   *CurrencyObjectType          `xml:"urn:com.workday/bsvc Ceiling_Currency_Reference,omitempty"`
}
```
This is a wrapper for the Allowance Plan Percent Based Profile Data; this is used to display or replace Allowance Plan Percent Based Profile Data.










## <a name="AllowancePlanPercentDataType">type</a> [AllowancePlanPercentDataType](/src/target/wwsgen_types.go?s=11157:12047#L170)
``` go
type AllowancePlanPercentDataType struct {
    Percentage                                      float64                                               `xml:"urn:com.workday/bsvc Percentage,omitempty"`
    CompensationBasisReference                      *CompensationBasisObjectType                          `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
    CeilingAmount                                   float64                                               `xml:"urn:com.workday/bsvc Ceiling_Amount,omitempty"`
    CeilingCurrencyReference                        *CurrencyObjectType                                   `xml:"urn:com.workday/bsvc Ceiling_Currency_Reference,omitempty"`
    AllowancePlanPercentBasedProfileReplacementData []AllowancePlanPercentBasedProfileReplacementDataType `xml:"urn:com.workday/bsvc Allowance_Plan_Percent_Based_Profile_Replacement_Data,omitempty"`
}
```
Contains Allowance Plan Percent data.










## <a name="AllowancePlanReimbursableDataType">type</a> [AllowancePlanReimbursableDataType](/src/target/wwsgen_types.go?s=12157:12435#L179)
``` go
type AllowancePlanReimbursableDataType struct {
    ExpenseItemReference        *ExpenseItemObjectType        `xml:"urn:com.workday/bsvc Expense_Item_Reference"`
    ExpenseAccumulatorReference *ExpenseAccumulatorObjectType `xml:"urn:com.workday/bsvc Expense_Accumulator_Reference"`
}
```
Allowance Plan Reimbursable Data consists of information pertinent only to reimbursable allowance plans.










## <a name="AllowancePlanUnitBasedProfileReplacementDataType">type</a> [AllowancePlanUnitBasedProfileReplacementDataType](/src/target/wwsgen_types.go?s=12581:13063#L185)
``` go
type AllowancePlanUnitBasedProfileReplacementDataType struct {
    EligibilityRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    NumberofUnits            float64                  `xml:"urn:com.workday/bsvc Number_of_Units,omitempty"`
    PerUnitAmount            float64                  `xml:"urn:com.workday/bsvc Per_Unit_Amount,omitempty"`
    CurrencyReference        *CurrencyObjectType      `xml:"urn:com.workday/bsvc Currency_Reference"`
}
```
This is a wrapper for the Allowance Plan Unit Based Profile Data; this is used to display or replace Allowance Plan Unit Based Profile Data.










## <a name="AllowancePlanUnitDataType">type</a> [AllowancePlanUnitDataType](/src/target/wwsgen_types.go?s=13103:13785#L193)
``` go
type AllowancePlanUnitDataType struct {
    PerUnitAmount                                float64                                            `xml:"urn:com.workday/bsvc Per_Unit_Amount,omitempty"`
    DefaultUnits                                 float64                                            `xml:"urn:com.workday/bsvc Default_Units,omitempty"`
    UnitOfMeasureReference                       *UnitofMeasureObjectType                           `xml:"urn:com.workday/bsvc Unit_Of_Measure_Reference"`
    AllowancePlanUnitBasedProfileReplacementData []AllowancePlanUnitBasedProfileReplacementDataType `xml:"urn:com.workday/bsvc Allowance_Plan_Unit_Based_Profile_Replacement_Data,omitempty"`
}
```
Contains Allowance Plan Unit data.










## <a name="AllowanceUnitPlanObjectIDType">type</a> [AllowanceUnitPlanObjectIDType](/src/target/wwsgen_types.go?s=13849:13979#L201)
``` go
type AllowanceUnitPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="AllowanceUnitPlanObjectType">type</a> [AllowanceUnitPlanObjectType](/src/target/wwsgen_types.go?s=13981:14209#L206)
``` go
type AllowanceUnitPlanObjectType struct {
    ID         []AllowanceUnitPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="AllowanceValuePlanObjectIDType">type</a> [AllowanceValuePlanObjectIDType](/src/target/wwsgen_types.go?s=14273:14404#L212)
``` go
type AllowanceValuePlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="AllowanceValuePlanObjectType">type</a> [AllowanceValuePlanObjectType](/src/target/wwsgen_types.go?s=14406:14637#L217)
``` go
type AllowanceValuePlanObjectType struct {
    ID         []AllowanceValuePlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="AlternatePayRangeDataType">type</a> [AlternatePayRangeDataType](/src/target/wwsgen_types.go?s=14682:15924#L223)
``` go
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
```
The alternative compensation pay range.










## <a name="AlternativePayRangeType">type</a> [AlternativePayRangeType](/src/target/wwsgen_types.go?s=15988:16227#L237)
``` go
type AlternativePayRangeType struct {
    Delete                *bool                      `xml:"urn:com.workday/bsvc Delete,omitempty"`
    AlternatePayRangeData *AlternatePayRangeDataType `xml:"urn:com.workday/bsvc Alternate_Pay_Range_Data"`
}
```
Contains an alternative pay range and its associated data.










## <a name="AndOrOperatorsObjectIDType">type</a> [AndOrOperatorsObjectIDType](/src/target/wwsgen_types.go?s=16291:16418#L243)
``` go
type AndOrOperatorsObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="AndOrOperatorsObjectType">type</a> [AndOrOperatorsObjectType](/src/target/wwsgen_types.go?s=16420:16639#L248)
``` go
type AndOrOperatorsObjectType struct {
    ID         []AndOrOperatorsObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="AssignEligiblePeriodActivitiesforEmployeeDataType">type</a> [AssignEligiblePeriodActivitiesforEmployeeDataType](/src/target/wwsgen_types.go?s=16723:17442#L254)
``` go
type AssignEligiblePeriodActivitiesforEmployeeDataType struct {
    EmployeeReference               *EmployeeObjectType                       `xml:"urn:com.workday/bsvc Employee_Reference"`
    PositionReference               *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
    EligibleActivityDate            time.Time                                 `xml:"urn:com.workday/bsvc Eligible_Activity_Date"`
    ReasonReference                 *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference"`
    EligiblePeriodActivitiesSubData []PeriodActivityEligibilityEntryDataType  `xml:"urn:com.workday/bsvc Eligible_Period_Activities_Sub_Data,omitempty"`
}
```
Wrapper Element for the Assign Employee's ligible Activities business process.










### <a name="AssignEligiblePeriodActivitiesforEmployeeDataType.MarshalXML">func</a> (\*AssignEligiblePeriodActivitiesforEmployeeDataType) [MarshalXML](/src/target/wwsgen_types.go?s=17444:17560#L262)
``` go
func (t *AssignEligiblePeriodActivitiesforEmployeeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="AssignEligiblePeriodActivitiesforEmployeeDataType.UnmarshalXML">func</a> (\*AssignEligiblePeriodActivitiesforEmployeeDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=17869:17987#L272)
``` go
func (t *AssignEligiblePeriodActivitiesforEmployeeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="AssignEligiblePeriodActivitiesforEmployeeInput">type</a> [AssignEligiblePeriodActivitiesforEmployeeInput](/src/target/wwsgen_client.go?s=28187:28403#L737)
``` go
type AssignEligiblePeriodActivitiesforEmployeeInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Assign_Eligible_Period_Activities_for_Employee_Request"`
    AssignEligiblePeriodActivitiesforEmployeeRequestType
}
```









## <a name="AssignEligiblePeriodActivitiesforEmployeeOutput">type</a> [AssignEligiblePeriodActivitiesforEmployeeOutput](/src/target/wwsgen_client.go?s=28405:28624#L742)
``` go
type AssignEligiblePeriodActivitiesforEmployeeOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Assign_Eligible_Period_Activities_for_Employee_Response"`
    AssignEligiblePeriodActivitiesforEmployeeResponseType
}
```









## <a name="AssignEligiblePeriodActivitiesforEmployeeRequestType">type</a> [AssignEligiblePeriodActivitiesforEmployeeRequestType](/src/target/wwsgen_types.go?s=18374:18936#L284)
``` go
type AssignEligiblePeriodActivitiesforEmployeeRequestType struct {
    BusinessProcessParameters                     *BusinessProcessParametersType                     `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
    AssignEligiblePeriodActivitiesforEmployeeData *AssignEligiblePeriodActivitiesforEmployeeDataType `xml:"urn:com.workday/bsvc Assign_Eligible_Period_Activities_for_Employee_Data"`
    Version                                       string                                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Web service operation to assign a period activity for an employee.










## <a name="AssignEligiblePeriodActivitiesforEmployeeResponseType">type</a> [AssignEligiblePeriodActivitiesforEmployeeResponseType](/src/target/wwsgen_types.go?s=19029:19371#L291)
``` go
type AssignEligiblePeriodActivitiesforEmployeeResponseType struct {
    PeriodActivityEligibilityEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Period_Activity_Eligibility_Event_Reference,omitempty"`
    Version                                 string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element for the Assign Employee's Eligible Period Activities business process.










## <a name="BasePayPlanDataType">type</a> [BasePayPlanDataType](/src/target/wwsgen_types.go?s=19487:20558#L297)
``` go
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
```
Base Pay Plan Data consists of common information from the following Base Pay Plans: Salary, Hourly, and Unit.










## <a name="BenchmarkJobCompositeDataType">type</a> [BenchmarkJobCompositeDataType](/src/target/wwsgen_types.go?s=20592:21430#L309)
``` go
type BenchmarkJobCompositeDataType struct {
    BenchmarkCompositeCategoryReference        *CompensationBenchmarkCompositeCategoryObjectType `xml:"urn:com.workday/bsvc Benchmark_Composite_Category_Reference"`
    CompetitiveMarketTargetPercentileReference *CompensationBenchmarkPercentileObjectType        `xml:"urn:com.workday/bsvc Competitive_Market_Target_Percentile_Reference,omitempty"`
    TargetSpread                               float64                                           `xml:"urn:com.workday/bsvc Target_Spread__,omitempty"`
    BenchmarkAmountReplacementData             []CompensationBenchmarkAmountReplacmentDataType   `xml:"urn:com.workday/bsvc Benchmark_Amount_Replacement_Data"`
    Delete                                     bool                                              `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}
```
Benchmark Job Composite Data










## <a name="BenchmarkJobDataType">type</a> [BenchmarkJobDataType](/src/target/wwsgen_types.go?s=21454:22961#L318)
``` go
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
```
Benchmark Job Data










### <a name="BenchmarkJobDataType.MarshalXML">func</a> (\*BenchmarkJobDataType) [MarshalXML](/src/target/wwsgen_types.go?s=22963:23050#L332)
``` go
func (t *BenchmarkJobDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="BenchmarkJobDataType.UnmarshalXML">func</a> (\*BenchmarkJobDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=23439:23528#L344)
``` go
func (t *BenchmarkJobDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="BenchmarkJobRequestCriteriaType">type</a> [BenchmarkJobRequestCriteriaType](/src/target/wwsgen_types.go?s=23961:24087#L358)
``` go
type BenchmarkJobRequestCriteriaType struct {
    IncludeInactive *bool `xml:"urn:com.workday/bsvc Include_Inactive,omitempty"`
}
```
Benchmark Job Request Criteria










## <a name="BenchmarkJobRequestReferencesType">type</a> [BenchmarkJobRequestReferencesType](/src/target/wwsgen_types.go?s=24125:24291#L363)
``` go
type BenchmarkJobRequestReferencesType struct {
    BenchmarkJobReference []CompensationBenchmarkDefaultObjectType `xml:"urn:com.workday/bsvc Benchmark_Job_Reference"`
}
```
Benchmark Job Request References










## <a name="BenchmarkJobResponseDataType">type</a> [BenchmarkJobResponseDataType](/src/target/wwsgen_types.go?s=24324:24454#L368)
``` go
type BenchmarkJobResponseDataType struct {
    BenchmarkJob []BenchmarkJobType `xml:"urn:com.workday/bsvc Benchmark_Job,omitempty"`
}
```
Benchmark Job Response Data










## <a name="BenchmarkJobResponseGroupType">type</a> [BenchmarkJobResponseGroupType](/src/target/wwsgen_types.go?s=24488:24614#L373)
``` go
type BenchmarkJobResponseGroupType struct {
    IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}
```
Benchmark Job Response Group










## <a name="BenchmarkJobType">type</a> [BenchmarkJobType](/src/target/wwsgen_types.go?s=24633:24912#L378)
``` go
type BenchmarkJobType struct {
    BenchmarkJobReference *CompensationBenchmarkDefaultObjectType `xml:"urn:com.workday/bsvc Benchmark_Job_Reference,omitempty"`
    BenchmarkJobData      []BenchmarkJobDataType                  `xml:"urn:com.workday/bsvc Benchmark_Job_Data,omitempty"`
}
```
Benchmark Job










## <a name="BenefitMultiplierOrderObjectIDType">type</a> [BenefitMultiplierOrderObjectIDType](/src/target/wwsgen_types.go?s=24976:25111#L384)
``` go
type BenefitMultiplierOrderObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="BenefitMultiplierOrderObjectType">type</a> [BenefitMultiplierOrderObjectType](/src/target/wwsgen_types.go?s=25113:25356#L389)
``` go
type BenefitMultiplierOrderObjectType struct {
    ID         []BenefitMultiplierOrderObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="BonusPaymentDataType">type</a> [BonusPaymentDataType](/src/target/wwsgen_types.go?s=25423:25912#L395)
``` go
type BonusPaymentDataType struct {
    BonusPlanReference *BonusPlanObjectType `xml:"urn:com.workday/bsvc Bonus_Plan_Reference"`
    Percent            float64              `xml:"urn:com.workday/bsvc Percent,omitempty"`
    Amount             float64              `xml:"urn:com.workday/bsvc Amount,omitempty"`
    CurrencyReference  *CurrencyObjectType  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
    Comment            string               `xml:"urn:com.workday/bsvc Comment,omitempty"`
}
```
Data for Employee's one-time ad-hoc bonus payment information










## <a name="BonusPercentPlanObjectIDType">type</a> [BonusPercentPlanObjectIDType](/src/target/wwsgen_types.go?s=25976:26105#L404)
``` go
type BonusPercentPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="BonusPercentPlanObjectType">type</a> [BonusPercentPlanObjectType](/src/target/wwsgen_types.go?s=26107:26332#L409)
``` go
type BonusPercentPlanObjectType struct {
    ID         []BonusPercentPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="BonusPlanAmountDataType">type</a> [BonusPlanAmountDataType](/src/target/wwsgen_types.go?s=26370:26687#L415)
``` go
type BonusPlanAmountDataType struct {
    Amount                           float64                                     `xml:"urn:com.workday/bsvc Amount,omitempty"`
    AmountPlanProfileReplacementData []BonusPlanProfileReplacementAmountDataType `xml:"urn:com.workday/bsvc Amount_Plan_Profile_Replacement_Data,omitempty"`
}
```
Contains Bonus Plan Amount data.










## <a name="BonusPlanDataType">type</a> [BonusPlanDataType](/src/target/wwsgen_types.go?s=26791:29957#L421)
``` go
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
```
Bonus Plan Data consists of common information from the following Bonus Plans: Amount and Percent.










## <a name="BonusPlanObjectIDType">type</a> [BonusPlanObjectIDType](/src/target/wwsgen_types.go?s=30021:30143#L444)
``` go
type BonusPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="BonusPlanObjectType">type</a> [BonusPlanObjectType](/src/target/wwsgen_types.go?s=30145:30349#L449)
``` go
type BonusPlanObjectType struct {
    ID         []BonusPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="BonusPlanPercentDataType">type</a> [BonusPlanPercentDataType](/src/target/wwsgen_types.go?s=30388:30865#L455)
``` go
type BonusPlanPercentDataType struct {
    Percentage                        float64                                          `xml:"urn:com.workday/bsvc Percentage,omitempty"`
    CompensationBasisReference        *CompensationBasisObjectType                     `xml:"urn:com.workday/bsvc Compensation_Basis_Reference"`
    PercentPlanProfileReplacementData []BonusPlanPercentPlanProfileReplacementDataType `xml:"urn:com.workday/bsvc Percent_Plan_Profile_Replacement_Data,omitempty"`
}
```
Contains Bonus Plan Percent data.










## <a name="BonusPlanPercentPlanProfileReplacementDataType">type</a> [BonusPlanPercentPlanProfileReplacementDataType](/src/target/wwsgen_types.go?s=30949:31346#L462)
``` go
type BonusPlanPercentPlanProfileReplacementDataType struct {
    EligibilityRuleReference   *ConditionRuleObjectType     `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    Percentage                 float64                      `xml:"urn:com.workday/bsvc Percentage,omitempty"`
    CompensationBasisReference *CompensationBasisObjectType `xml:"urn:com.workday/bsvc Compensation_Basis_Reference"`
}
```
Used to display or replace Percent Plan Profile Data for a Percent Bonus Plan.










## <a name="BonusPlanProfileReplacementAmountDataType">type</a> [BonusPlanProfileReplacementAmountDataType](/src/target/wwsgen_types.go?s=31523:31883#L469)
``` go
type BonusPlanProfileReplacementAmountDataType struct {
    EligibilityRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    Amount                   float64                  `xml:"urn:com.workday/bsvc Amount,omitempty"`
    CurrencyReference        *CurrencyObjectType      `xml:"urn:com.workday/bsvc Currency_Reference"`
}
```
This is a wrapper for the Bonus Plan Amount; this is used to display or replace Amount Plan Profile Data; it is required and mutually exclusive with Bonus Plan Percentage.










## <a name="BusinessProcessAttachmentDataType">type</a> [BusinessProcessAttachmentDataType](/src/target/wwsgen_types.go?s=31969:32641#L476)
``` go
type BusinessProcessAttachmentDataType struct {
    FileName                         string                             `xml:"urn:com.workday/bsvc File_Name"`
    EventAttachmentDescription       string                             `xml:"urn:com.workday/bsvc Event_Attachment_Description,omitempty"`
    EventAttachmentCategoryReference *EventAttachmentCategoryObjectType `xml:"urn:com.workday/bsvc Event_Attachment_Category_Reference,omitempty"`
    File                             *[]byte                            `xml:"urn:com.workday/bsvc File,omitempty"`
    ContentType                      string                             `xml:"urn:com.workday/bsvc Content_Type,omitempty"`
}
```
Element for the attachments pertaining to a Event entered through a web service.










### <a name="BusinessProcessAttachmentDataType.MarshalXML">func</a> (\*BusinessProcessAttachmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=32643:32743#L484)
``` go
func (t *BusinessProcessAttachmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="BusinessProcessAttachmentDataType.UnmarshalXML">func</a> (\*BusinessProcessAttachmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=32995:33097#L494)
``` go
func (t *BusinessProcessAttachmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="BusinessProcessCommentDataType">type</a> [BusinessProcessCommentDataType](/src/target/wwsgen_types.go?s=33405:33624#L506)
``` go
type BusinessProcessCommentDataType struct {
    Comment         string            `xml:"urn:com.workday/bsvc Comment,omitempty"`
    WorkerReference *WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
}
```
Captures a comment for the Business Process.










## <a name="BusinessProcessParametersType">type</a> [BusinessProcessParametersType](/src/target/wwsgen_types.go?s=34015:34552#L512)
``` go
type BusinessProcessParametersType struct {
    AutoComplete                  *bool                               `xml:"urn:com.workday/bsvc Auto_Complete,omitempty"`
    RunNow                        *bool                               `xml:"urn:com.workday/bsvc Run_Now,omitempty"`
    CommentData                   *BusinessProcessCommentDataType     `xml:"urn:com.workday/bsvc Comment_Data,omitempty"`
    BusinessProcessAttachmentData []BusinessProcessAttachmentDataType `xml:"urn:com.workday/bsvc Business_Process_Attachment_Data,omitempty"`
}
```
Container for the processing options for a business process. If no options are submitted (or the options are submitted as 'false') then the business process is simply initiated as if it where submitted on-line with approvals, reviews, notifications and to-do's in place. If the Initiator is an Integration System User, any validations you configured on the Initiation step are ignored.










## <a name="CalculatedPlanDataType">type</a> [CalculatedPlanDataType](/src/target/wwsgen_types.go?s=34596:35334#L520)
``` go
type CalculatedPlanDataType struct {
    CompensationElementReference  *CompensationPayEarningObjectType            `xml:"urn:com.workday/bsvc Compensation_Element_Reference,omitempty"`
    SetupSecuritySegmentReference []CompensationSetupSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Setup_Security_Segment_Reference,omitempty"`
    CalculationReference          *CalculationObjectType                       `xml:"urn:com.workday/bsvc Calculation_Reference,omitempty"`
    CurrencyReference             *CurrencyObjectType                          `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
    FrequencyReference            *FrequencyObjectType                         `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
}
```
Data for creation of a Calculated Plan










## <a name="CalculatedPlanObjectIDType">type</a> [CalculatedPlanObjectIDType](/src/target/wwsgen_types.go?s=35398:35525#L529)
``` go
type CalculatedPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CalculatedPlanObjectType">type</a> [CalculatedPlanObjectType](/src/target/wwsgen_types.go?s=35527:35746#L534)
``` go
type CalculatedPlanObjectType struct {
    ID         []CalculatedPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CalculationObjectIDType">type</a> [CalculationObjectIDType](/src/target/wwsgen_types.go?s=35810:35934#L540)
``` go
type CalculationObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CalculationObjectType">type</a> [CalculationObjectType](/src/target/wwsgen_types.go?s=35936:36146#L545)
``` go
type CalculationObjectType struct {
    ID         []CalculationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CheckPositionBudgetSubBusinessProcessType">type</a> [CheckPositionBudgetSubBusinessProcessType](/src/target/wwsgen_types.go?s=36221:36423#L551)
``` go
type CheckPositionBudgetSubBusinessProcessType struct {
    BusinessSubProcessParameters *FinancialsBusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
}
```
Top element used for Check Position Budget as a sub business process.










## <a name="Client">type</a> [Client](/src/target/wwsgen_client.go?s=314:349#L17)
``` go
type Client struct {
    *wws.Client
}
```









### <a name="Client.AddStockGrant">func</a> (\*Client) [AddStockGrant](/src/target/wwsgen_client.go?s=17968:18087#L486)
``` go
func (c *Client) AddStockGrant(ctx context.Context, input *AddStockGrantInput) (output *AddStockGrantOutput, err error)
```
AddStockGrant (Add_Stock_Grant)

This operation allows the addition of stock grants to an employee via the Request Stock Option Grant business process.




### <a name="Client.AssignEligiblePeriodActivitiesforEmployee">func</a> (\*Client) [AssignEligiblePeriodActivitiesforEmployee](/src/target/wwsgen_client.go?s=27798:28001#L728)
``` go
func (c *Client) AssignEligiblePeriodActivitiesforEmployee(ctx context.Context, input *AssignEligiblePeriodActivitiesforEmployeeInput) (output *AssignEligiblePeriodActivitiesforEmployeeOutput, err error)
```
AssignEligiblePeriodActivitiesforEmployee (Assign_Eligible_Period_Activities_for_Employee)

This operation allows the requesting of a period activity assignment for an employee via the Assign Eligible Period Activities for Employee business process.




### <a name="Client.CreateSeveranceWorksheet">func</a> (\*Client) [CreateSeveranceWorksheet](/src/target/wwsgen_client.go?s=23294:23446#L618)
``` go
func (c *Client) CreateSeveranceWorksheet(ctx context.Context, input *CreateSeveranceWorksheetInput) (output *CreateSeveranceWorksheetOutput, err error)
```
CreateSeveranceWorksheet (Create_Severance_Worksheet)

This operation allows the launching of an employee severance worksheet via the Severance Worksheet business process.




### <a name="Client.GetBenchmarkJobs">func</a> (\*Client) [GetBenchmarkJobs](/src/target/wwsgen_client.go?s=16521:16649#L442)
``` go
func (c *Client) GetBenchmarkJobs(ctx context.Context, input *GetBenchmarkJobsInput) (output *GetBenchmarkJobsOutput, err error)
```
GetBenchmarkJobs (Get_Benchmark_Jobs)

This operation allows the retrieval of benchmark job information.




### <a name="Client.GetCompensationEligibilityRules">func</a> (\*Client) [GetCompensationEligibilityRules](/src/target/wwsgen_client.go?s=19517:19690#L530)
``` go
func (c *Client) GetCompensationEligibilityRules(ctx context.Context, input *GetCompensationEligibilityRulesInput) (output *GetCompensationEligibilityRulesOutput, err error)
```
GetCompensationEligibilityRules (Get_Compensation_Eligibility_Rules)

This operation returns Compensation Eligibility Rules




### <a name="Client.GetCompensationEligibilityRuleswithoutDependencies">func</a> (\*Client) [GetCompensationEligibilityRuleswithoutDependencies](/src/target/wwsgen_client.go?s=20473:20703#L552)
``` go
func (c *Client) GetCompensationEligibilityRuleswithoutDependencies(ctx context.Context, input *GetCompensationEligibilityRuleswithoutDependenciesInput) (output *GetCompensationEligibilityRuleswithoutDependenciesOutput, err error)
```
GetCompensationEligibilityRuleswithoutDependencies (Get_Compensation_Eligibility_Rules_without_Dependencies)

This operation returns Compensation Eligibility Rules without dependent Compensation Eligibility Rule references.




### <a name="Client.GetCompensationGradeHierarchy">func</a> (\*Client) [GetCompensationGradeHierarchy](/src/target/wwsgen_client.go?s=24201:24368#L640)
``` go
func (c *Client) GetCompensationGradeHierarchy(ctx context.Context, input *GetCompensationGradeHierarchyInput) (output *GetCompensationGradeHierarchyOutput, err error)
```
GetCompensationGradeHierarchy (Get_Compensation_Grade_Hierarchy)

This operation allows the retrieval of compensation grade hierarchy information. You must have access to the &#34;Set Up: Compensation General&#34; domain to access this operation.




### <a name="Client.GetCompensationGrades">func</a> (\*Client) [GetCompensationGrades](/src/target/wwsgen_client.go?s=9236:9379#L244)
``` go
func (c *Client) GetCompensationGrades(ctx context.Context, input *GetCompensationGradesInput) (output *GetCompensationGradesOutput, err error)
```
GetCompensationGrades (Get_Compensation_Grades)

This operation allows the retrieval of compensation grade and grade profile information.  You must have access to the &#34;Set Up: Compensation Packages&#34; domain to access this operation.




### <a name="Client.GetCompensationMatrices">func</a> (\*Client) [GetCompensationMatrices](/src/target/wwsgen_client.go?s=7603:7752#L200)
``` go
func (c *Client) GetCompensationMatrices(ctx context.Context, input *GetCompensationMatricesInput) (output *GetCompensationMatricesOutput, err error)
```
GetCompensationMatrices (Get_Compensation_Matrices)

This operation retrieves Merit Increase Matrices (MIMs) information for use by merit or bonus plans.




### <a name="Client.GetCompensationPlans">func</a> (\*Client) [GetCompensationPlans](/src/target/wwsgen_client.go?s=10855:10995#L288)
``` go
func (c *Client) GetCompensationPlans(ctx context.Context, input *GetCompensationPlansInput) (output *GetCompensationPlansOutput, err error)
```
GetCompensationPlans (Get_Compensation_Plans)

This operation allows the retrieval of detailed compensation plan information.




### <a name="Client.GetCompensationScorecardResults">func</a> (\*Client) [GetCompensationScorecardResults](/src/target/wwsgen_client.go?s=22390:22563#L596)
``` go
func (c *Client) GetCompensationScorecardResults(ctx context.Context, input *GetCompensationScorecardResultsInput) (output *GetCompensationScorecardResultsOutput, err error)
```
GetCompensationScorecardResults (Get_Compensation_Scorecard_Results)

This operation allows the retrieval of one or more scorecard results




### <a name="Client.GetCompensationScorecards">func</a> (\*Client) [GetCompensationScorecards](/src/target/wwsgen_client.go?s=14279:14434#L376)
``` go
func (c *Client) GetCompensationScorecards(ctx context.Context, input *GetCompensationScorecardsInput) (output *GetCompensationScorecardsOutput, err error)
```
GetCompensationScorecards (Get_Compensation_Scorecards)

This operation allows the retrieval of one or more scorecards.




### <a name="Client.GetEligibleEarnings">func</a> (\*Client) [GetEligibleEarnings](/src/target/wwsgen_client.go?s=6839:6976#L178)
``` go
func (c *Client) GetEligibleEarnings(ctx context.Context, input *GetEligibleEarningsInput) (output *GetEligibleEarningsOutput, err error)
```
GetEligibleEarnings (Get_Eligible_Earnings)

This operation retrieves the eligible earnings override values that are associated with a given employee.




### <a name="Client.GetEnhancedSeveranceMatrices">func</a> (\*Client) [GetEnhancedSeveranceMatrices](/src/target/wwsgen_client.go?s=37794:37958#L970)
``` go
func (c *Client) GetEnhancedSeveranceMatrices(ctx context.Context, input *GetEnhancedSeveranceMatricesInput) (output *GetEnhancedSeveranceMatricesOutput, err error)
```
GetEnhancedSeveranceMatrices (Get_Enhanced_Severance_Matrices)

This service retrieves the Enhanced Severance Matrices from the system.




### <a name="Client.GetFuturePaymentPlanAssignments">func</a> (\*Client) [GetFuturePaymentPlanAssignments](/src/target/wwsgen_client.go?s=12475:12648#L332)
``` go
func (c *Client) GetFuturePaymentPlanAssignments(ctx context.Context, input *GetFuturePaymentPlanAssignmentsInput) (output *GetFuturePaymentPlanAssignmentsOutput, err error)
```
GetFuturePaymentPlanAssignments (Get_Future_Payment_Plan_Assignments)

This operation allows the retrieval of future payment plan assignments for one or more employees. You must have access to the Worker:Compensation by Organization domain.




### <a name="Client.GetOneTimePaymentPlanConfigurableCategories">func</a> (\*Client) [GetOneTimePaymentPlanConfigurableCategories](/src/target/wwsgen_client.go?s=33268:33477#L860)
``` go
func (c *Client) GetOneTimePaymentPlanConfigurableCategories(ctx context.Context, input *GetOneTimePaymentPlanConfigurableCategoriesInput) (output *GetOneTimePaymentPlanConfigurableCategoriesOutput, err error)
```
GetOneTimePaymentPlanConfigurableCategories (Get_One-Time_Payment_Plan_Configurable_Categories)

Web Service for Get One-Time Payment Plan Configurable Categories




### <a name="Client.GetPeriodActivityRateMatrices">func</a> (\*Client) [GetPeriodActivityRateMatrices](/src/target/wwsgen_client.go?s=29587:29754#L772)
``` go
func (c *Client) GetPeriodActivityRateMatrices(ctx context.Context, input *GetPeriodActivityRateMatricesInput) (output *GetPeriodActivityRateMatricesOutput, err error)
```
GetPeriodActivityRateMatrices (Get_Period_Activity_Rate_Matrices)

This operation returns the Period Activity Rate Matrices created




### <a name="Client.GetPeriodActivityTasks">func</a> (\*Client) [GetPeriodActivityTasks](/src/target/wwsgen_client.go?s=26067:26213#L684)
``` go
func (c *Client) GetPeriodActivityTasks(ctx context.Context, input *GetPeriodActivityTasksInput) (output *GetPeriodActivityTasksOutput, err error)
```
GetPeriodActivityTasks (Get_Period_Activity_Tasks)

This operation allows the retrieval of period activity tasks.  You must have access to the &#34;Set Up: Work Activity for Pay&#34; domain to access this operation.




### <a name="Client.GetPreviousSystemCompensationHistory">func</a> (\*Client) [GetPreviousSystemCompensationHistory](/src/target/wwsgen_client.go?s=1642:1830#L46)
``` go
func (c *Client) GetPreviousSystemCompensationHistory(ctx context.Context, input *GetPreviousSystemCompensationHistoryInput) (output *GetPreviousSystemCompensationHistoryOutput, err error)
```
GetPreviousSystemCompensationHistory (Get_Previous_System_Compensation_History)

This operation will return any previous system compensation history which has been imported for a given employee.




### <a name="Client.GetStockParticipationRateTables">func</a> (\*Client) [GetStockParticipationRateTables](/src/target/wwsgen_client.go?s=32371:32544#L838)
``` go
func (c *Client) GetStockParticipationRateTables(ctx context.Context, input *GetStockParticipationRateTablesInput) (output *GetStockParticipationRateTablesOutput, err error)
```
GetStockParticipationRateTables (Get_Stock_Participation_Rate_Tables)

This operation allows you to retrieve Stock Participation Rate Table information.




### <a name="Client.ImportEligibleEarningsOverride">func</a> (\*Client) [ImportEligibleEarningsOverride](/src/target/wwsgen_client.go?s=38807:38977#L992)
``` go
func (c *Client) ImportEligibleEarningsOverride(ctx context.Context, input *ImportEligibleEarningsOverrideInput) (output *ImportEligibleEarningsOverrideOutput, err error)
```
ImportEligibleEarningsOverride (Import_Eligible_Earnings_Override)

This operation adds or updates eligible earnings override information for a given employee. If there is an in-progress Bonus Process with configured Participation Rules, then that process will react to the eligible earnings in this operation.




### <a name="Client.ImportRequestCompensationChange">func</a> (\*Client) [ImportRequestCompensationChange](/src/target/wwsgen_client.go?s=39646:39819#L1014)
``` go
func (c *Client) ImportRequestCompensationChange(ctx context.Context, input *ImportRequestCompensationChangeInput) (output *ImportRequestCompensationChangeOutput, err error)
```
ImportRequestCompensationChange (Import_Request_Compensation_Change)

High-volume web service task to Request Compensation Change for employee(s)




### <a name="Client.ManagePeriodActivityPayAssignments">func</a> (\*Client) [ManagePeriodActivityPayAssignments](/src/target/wwsgen_client.go?s=30572:30754#L794)
``` go
func (c *Client) ManagePeriodActivityPayAssignments(ctx context.Context, input *ManagePeriodActivityPayAssignmentsInput) (output *ManagePeriodActivityPayAssignmentsOutput, err error)
```
ManagePeriodActivityPayAssignments (Manage_Period_Activity_Pay_Assignments)

This operation allows the management of period activity based assignments for a given worker and a bounding date range or period via the Manage Period Activity Pay Assignments business process.




### <a name="Client.PutBenchmarkJob">func</a> (\*Client) [PutBenchmarkJob](/src/target/wwsgen_client.go?s=15846:15971#L420)
``` go
func (c *Client) PutBenchmarkJob(ctx context.Context, input *PutBenchmarkJobInput) (output *PutBenchmarkJobOutput, err error)
```
PutBenchmarkJob (Put_Benchmark_Job)

This operation allows the addition and updating of benchmark job information.




### <a name="Client.PutCompensationEligibilityRule">func</a> (\*Client) [PutCompensationEligibilityRule](/src/target/wwsgen_client.go?s=18671:18841#L508)
``` go
func (c *Client) PutCompensationEligibilityRule(ctx context.Context, input *PutCompensationEligibilityRuleInput) (output *PutCompensationEligibilityRuleOutput, err error)
```
PutCompensationEligibilityRule (Put_Compensation_Eligibility_Rule)

This operation allows for the creation or update of a Compensation Eligibility Rule.




### <a name="Client.PutCompensationGrade">func</a> (\*Client) [PutCompensationGrade](/src/target/wwsgen_client.go?s=10109:10249#L266)
``` go
func (c *Client) PutCompensationGrade(ctx context.Context, input *PutCompensationGradeInput) (output *PutCompensationGradeOutput, err error)
```
PutCompensationGrade (Put_Compensation_Grade)

This operation adds or updates compensation grade and compensation grade profile information.  You must have access to the &#34;Set Up: Compensation Packages&#34; domain to access this operation.




### <a name="Client.PutCompensationGradeHierarchy">func</a> (\*Client) [PutCompensationGradeHierarchy](/src/target/wwsgen_client.go?s=25139:25306#L662)
``` go
func (c *Client) PutCompensationGradeHierarchy(ctx context.Context, input *PutCompensationGradeHierarchyInput) (output *PutCompensationGradeHierarchyOutput, err error)
```
PutCompensationGradeHierarchy (Put_Compensation_Grade_Hierarchy)

This operation adds or updates compensation grade hierarchy You must have access to the &#34;Set Up: Compensation General&#34; domain to access this operation.




### <a name="Client.PutCompensationMatrix">func</a> (\*Client) [PutCompensationMatrix](/src/target/wwsgen_client.go?s=8366:8509#L222)
``` go
func (c *Client) PutCompensationMatrix(ctx context.Context, input *PutCompensationMatrixInput) (output *PutCompensationMatrixOutput, err error)
```
PutCompensationMatrix (Put_Compensation_Matrix)

This operation adds or updates compensation matrix information.




### <a name="Client.PutCompensationPlans">func</a> (\*Client) [PutCompensationPlans](/src/target/wwsgen_client.go?s=11618:11758#L310)
``` go
func (c *Client) PutCompensationPlans(ctx context.Context, input *PutCompensationPlansInput) (output *PutCompensationPlansOutput, err error)
```
PutCompensationPlans (Put_Compensation_Plans)

This operation supports the addition or modification of detailed compensation plan information.




### <a name="Client.PutCompensationScorecard">func</a> (\*Client) [PutCompensationScorecard](/src/target/wwsgen_client.go?s=15071:15223#L398)
``` go
func (c *Client) PutCompensationScorecard(ctx context.Context, input *PutCompensationScorecardInput) (output *PutCompensationScorecardOutput, err error)
```
PutCompensationScorecard (Put_Compensation_Scorecard)

This operation allows the creation or modification of a scorecard.




### <a name="Client.PutCompensationScorecardResult">func</a> (\*Client) [PutCompensationScorecardResult](/src/target/wwsgen_client.go?s=21529:21699#L574)
``` go
func (c *Client) PutCompensationScorecardResult(ctx context.Context, input *PutCompensationScorecardResultInput) (output *PutCompensationScorecardResultOutput, err error)
```
PutCompensationScorecardResult (Put_Compensation_Scorecard_Result)

This operation allows the creation or modification of a scorecard result against a specific scorecard




### <a name="Client.PutEligibleEarnings">func</a> (\*Client) [PutEligibleEarnings](/src/target/wwsgen_client.go?s=6078:6215#L156)
``` go
func (c *Client) PutEligibleEarnings(ctx context.Context, input *PutEligibleEarningsInput) (output *PutEligibleEarningsOutput, err error)
```
PutEligibleEarnings (Put_Eligible_Earnings)

Adds or updates eligible earnings override information for an employee. This operation doesn’t follow participation rules. Use Import Eligible Earnings instead to follow those rules. This operation might affect in-progress bonus reviews.




### <a name="Client.PutEnhancedSeveranceMatrix">func</a> (\*Client) [PutEnhancedSeveranceMatrix](/src/target/wwsgen_client.go?s=36976:37134#L948)
``` go
func (c *Client) PutEnhancedSeveranceMatrix(ctx context.Context, input *PutEnhancedSeveranceMatrixInput) (output *PutEnhancedSeveranceMatrixOutput, err error)
```
PutEnhancedSeveranceMatrix (Put_Enhanced_Severance_Matrix)

Create or modify data for an Enhanced Severance Matrix.




### <a name="Client.PutFuturePaymentPlanAssignment">func</a> (\*Client) [PutFuturePaymentPlanAssignment](/src/target/wwsgen_client.go?s=13435:13605#L354)
``` go
func (c *Client) PutFuturePaymentPlanAssignment(ctx context.Context, input *PutFuturePaymentPlanAssignmentInput) (output *PutFuturePaymentPlanAssignmentOutput, err error)
```
PutFuturePaymentPlanAssignment (Put_Future_Payment_Plan_Assignment)

This operation adds or updates future payment plan assignments for a given employee. You must have access to the Worker:Compensation by Organization domain.




### <a name="Client.PutOneTimePaymentPlanConfigurableCategory">func</a> (\*Client) [PutOneTimePaymentPlanConfigurableCategory](/src/target/wwsgen_client.go?s=34279:34482#L882)
``` go
func (c *Client) PutOneTimePaymentPlanConfigurableCategory(ctx context.Context, input *PutOneTimePaymentPlanConfigurableCategoryInput) (output *PutOneTimePaymentPlanConfigurableCategoryOutput, err error)
```
PutOneTimePaymentPlanConfigurableCategory (Put_One-Time_Payment_Plan_Configurable_Category)

Web Service for One-Time Payment Plan Configurable Category




### <a name="Client.PutPeriodActivityRateMatrix">func</a> (\*Client) [PutPeriodActivityRateMatrix](/src/target/wwsgen_client.go?s=28761:28922#L750)
``` go
func (c *Client) PutPeriodActivityRateMatrix(ctx context.Context, input *PutPeriodActivityRateMatrixInput) (output *PutPeriodActivityRateMatrixOutput, err error)
```
PutPeriodActivityRateMatrix (Put_Period_Activity_Rate_Matrix)

This operation creates or updates a Period Activity Rate Matrix




### <a name="Client.PutPeriodActivityTask">func</a> (\*Client) [PutPeriodActivityTask](/src/target/wwsgen_client.go?s=26916:27059#L706)
``` go
func (c *Client) PutPeriodActivityTask(ctx context.Context, input *PutPeriodActivityTaskInput) (output *PutPeriodActivityTaskOutput, err error)
```
PutPeriodActivityTask (Put_Period_Activity_Task)

This operation adds or updates a period activity task.  You must have access to the &#34;Set Up: Work Activity for Pay&#34; domain to access this operation.




### <a name="Client.PutPreviousSystemCompensationHistory">func</a> (\*Client) [PutPreviousSystemCompensationHistory](/src/target/wwsgen_client.go?s=663:851#L24)
``` go
func (c *Client) PutPreviousSystemCompensationHistory(ctx context.Context, input *PutPreviousSystemCompensationHistoryInput) (output *PutPreviousSystemCompensationHistoryOutput, err error)
```
PutPreviousSystemCompensationHistory (Put_Previous_System_Compensation_History)

This operation will load previous system compensation history for a given employee. The operation allows free-form text entry indication of changes to a workers compensation before the implementation in the Workday system.




### <a name="Client.PutStockParticipationRateTable">func</a> (\*Client) [PutStockParticipationRateTable](/src/target/wwsgen_client.go?s=31494:31664#L816)
``` go
func (c *Client) PutStockParticipationRateTable(ctx context.Context, input *PutStockParticipationRateTableInput) (output *PutStockParticipationRateTableOutput, err error)
```
PutStockParticipationRateTable (Put_Stock_Participation_Rate_Table)

This operation allows you to create, update, or delete a Stock Participation Rate Table.




### <a name="Client.RequestBonusPayment">func</a> (\*Client) [RequestBonusPayment](/src/target/wwsgen_client.go?s=2591:2728#L68)
``` go
func (c *Client) RequestBonusPayment(ctx context.Context, input *RequestBonusPaymentInput) (output *RequestBonusPaymentOutput, err error)
```
RequestBonusPayment (Request_Bonus_Payment)

This operation allows the requesting of a bonus payment for an employee via the Request Bonus Payment business process.




### <a name="Client.RequestCompensationChange">func</a> (\*Client) [RequestCompensationChange](/src/target/wwsgen_client.go?s=5123:5278#L134)
``` go
func (c *Client) RequestCompensationChange(ctx context.Context, input *RequestCompensationChangeInput) (output *RequestCompensationChangeOutput, err error)
```
RequestCompensationChange (Request_Compensation_Change)

This operation allows the requesting of a compensation change for an employee via the Request Compensation Change business process.




### <a name="Client.RequestEmployeeMeritAdjustment">func</a> (\*Client) [RequestEmployeeMeritAdjustment](/src/target/wwsgen_client.go?s=3404:3574#L90)
``` go
func (c *Client) RequestEmployeeMeritAdjustment(ctx context.Context, input *RequestEmployeeMeritAdjustmentInput) (output *RequestEmployeeMeritAdjustmentOutput, err error)
```
RequestEmployeeMeritAdjustment (Request_Employee_Merit_Adjustment)

This operation allows the requesting of a merit adjustment for an employee via the Request Employee Merit Adjustment business process.




### <a name="Client.RequestOneTimePayment">func</a> (\*Client) [RequestOneTimePayment](/src/target/wwsgen_client.go?s=4302:4445#L112)
``` go
func (c *Client) RequestOneTimePayment(ctx context.Context, input *RequestOneTimePaymentInput) (output *RequestOneTimePaymentOutput, err error)
```
RequestOneTimePayment (Request_One-Time_Payment)

This operation allows the requesting of a one-time payment for an employee via the Request One-Time Payment business process.




### <a name="Client.RequestRequisitionCompensationChange">func</a> (\*Client) [RequestRequisitionCompensationChange](/src/target/wwsgen_client.go?s=35295:35483#L904)
``` go
func (c *Client) RequestRequisitionCompensationChange(ctx context.Context, input *RequestRequisitionCompensationChangeInput) (output *RequestRequisitionCompensationChangeOutput, err error)
```
RequestRequisitionCompensationChange (Request_Requisition_Compensation_Change)

Web service operation to request an ad-hoc requisition compensation change for a job requisition.




### <a name="Client.RetrieveSeveranceWorksheet">func</a> (\*Client) [RetrieveSeveranceWorksheet](/src/target/wwsgen_client.go?s=36176:36334#L926)
``` go
func (c *Client) RetrieveSeveranceWorksheet(ctx context.Context, input *RetrieveSeveranceWorksheetInput) (output *RetrieveSeveranceWorksheetOutput, err error)
```
RetrieveSeveranceWorksheet (Retrieve_Severance_Worksheet)

Returns an Employee Severance Worksheet




### <a name="Client.UpdateStockGrant">func</a> (\*Client) [UpdateStockGrant](/src/target/wwsgen_client.go?s=17236:17364#L464)
``` go
func (c *Client) UpdateStockGrant(ctx context.Context, input *UpdateStockGrantInput) (output *UpdateStockGrantOutput, err error)
```
UpdateStockGrant (Update_Stock_Grant)

This operation allows the updating of an existing stock grant which has been given to a worker.




## <a name="CommissionPlanDataType">type</a> [CommissionPlanDataType](/src/target/wwsgen_types.go?s=36449:37918#L556)
``` go
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
```
Commission Plan Data










## <a name="CommissionPlanObjectIDType">type</a> [CommissionPlanObjectIDType](/src/target/wwsgen_types.go?s=37982:38109#L570)
``` go
type CommissionPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CommissionPlanObjectType">type</a> [CommissionPlanObjectType](/src/target/wwsgen_types.go?s=38111:38330#L575)
``` go
type CommissionPlanObjectType struct {
    ID         []CommissionPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CommissionPlanProfileReplacementDataType">type</a> [CommissionPlanProfileReplacementDataType](/src/target/wwsgen_types.go?s=38456:38917#L581)
``` go
type CommissionPlanProfileReplacementDataType struct {
    EligibilityRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    Amount                   float64                  `xml:"urn:com.workday/bsvc Amount,omitempty"`
    DrawAmount               float64                  `xml:"urn:com.workday/bsvc Draw_Amount,omitempty"`
    CurrencyReference        *CurrencyObjectType      `xml:"urn:com.workday/bsvc Currency_Reference"`
}
```
This is a wrapper for the Commission Plan Profile Data; this is used to display or replace Commission Plan Profile Data.










## <a name="CommonYesNoObjectIDType">type</a> [CommonYesNoObjectIDType](/src/target/wwsgen_types.go?s=38981:39105#L589)
``` go
type CommonYesNoObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CommonYesNoObjectType">type</a> [CommonYesNoObjectType](/src/target/wwsgen_types.go?s=39107:39317#L594)
``` go
type CommonYesNoObjectType struct {
    ID         []CommonYesNoObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompaRatioRangeSegmentObjectIDType">type</a> [CompaRatioRangeSegmentObjectIDType](/src/target/wwsgen_types.go?s=39381:39516#L600)
``` go
type CompaRatioRangeSegmentObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompaRatioRangeSegmentObjectType">type</a> [CompaRatioRangeSegmentObjectType](/src/target/wwsgen_types.go?s=39518:39761#L605)
``` go
type CompaRatioRangeSegmentObjectType struct {
    ID         []CompaRatioRangeSegmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensatableDefaultGuidelinesDataType">type</a> [CompensatableDefaultGuidelinesDataType](/src/target/wwsgen_types.go?s=39845:40464#L611)
``` go
type CompensatableDefaultGuidelinesDataType struct {
    CompensationPackageReference      *CompensationPackageObjectType      `xml:"urn:com.workday/bsvc Compensation_Package_Reference,omitempty"`
    CompensationGradeReference        *CompensationGradeObjectType        `xml:"urn:com.workday/bsvc Compensation_Grade_Reference,omitempty"`
    CompensationGradeProfileReference *CompensationGradeProfileObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Profile_Reference,omitempty"`
    CompensationStepReference         *CompensationStepObjectType         `xml:"urn:com.workday/bsvc Compensation_Step_Reference,omitempty"`
}
```
Data element containing the compensation guidelines for a compensation change.










## <a name="CompensatableGuidelinesDataType">type</a> [CompensatableGuidelinesDataType](/src/target/wwsgen_types.go?s=40548:41293#L619)
``` go
type CompensatableGuidelinesDataType struct {
    CompensationPackageReference      *CompensationPackageObjectType      `xml:"urn:com.workday/bsvc Compensation_Package_Reference,omitempty"`
    CompensationGradeReference        *CompensationGradeObjectType        `xml:"urn:com.workday/bsvc Compensation_Grade_Reference,omitempty"`
    CompensationGradeProfileReference *CompensationGradeProfileObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Profile_Reference,omitempty"`
    CompensationStepReference         *CompensationStepObjectType         `xml:"urn:com.workday/bsvc Compensation_Step_Reference,omitempty"`
    ProgressionStartDate              *time.Time                          `xml:"urn:com.workday/bsvc Progression_Start_Date,omitempty"`
}
```
Data element containing the compensation guidelines for a compensation change.










### <a name="CompensatableGuidelinesDataType.MarshalXML">func</a> (\*CompensatableGuidelinesDataType) [MarshalXML](/src/target/wwsgen_types.go?s=41295:41393#L627)
``` go
func (t *CompensatableGuidelinesDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="CompensatableGuidelinesDataType.UnmarshalXML">func</a> (\*CompensatableGuidelinesDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=41693:41793#L637)
``` go
func (t *CompensatableGuidelinesDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="CompensationAssignableComponentTypeObjectIDType">type</a> [CompensationAssignableComponentTypeObjectIDType](/src/target/wwsgen_types.go?s=42163:42311#L649)
``` go
type CompensationAssignableComponentTypeObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationAssignableComponentTypeObjectType">type</a> [CompensationAssignableComponentTypeObjectType](/src/target/wwsgen_types.go?s=42313:42595#L654)
``` go
type CompensationAssignableComponentTypeObjectType struct {
    ID         []CompensationAssignableComponentTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationAssignablePlanObjectIDType">type</a> [CompensationAssignablePlanObjectIDType](/src/target/wwsgen_types.go?s=42659:42798#L660)
``` go
type CompensationAssignablePlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationAssignablePlanObjectType">type</a> [CompensationAssignablePlanObjectType](/src/target/wwsgen_types.go?s=42800:43055#L665)
``` go
type CompensationAssignablePlanObjectType struct {
    ID         []CompensationAssignablePlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationBasisConfigurableObjectIDType">type</a> [CompensationBasisConfigurableObjectIDType](/src/target/wwsgen_types.go?s=43119:43261#L671)
``` go
type CompensationBasisConfigurableObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationBasisConfigurableObjectType">type</a> [CompensationBasisConfigurableObjectType](/src/target/wwsgen_types.go?s=43263:43527#L676)
``` go
type CompensationBasisConfigurableObjectType struct {
    ID         []CompensationBasisConfigurableObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationBasisObjectIDType">type</a> [CompensationBasisObjectIDType](/src/target/wwsgen_types.go?s=43591:43721#L682)
``` go
type CompensationBasisObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationBasisObjectType">type</a> [CompensationBasisObjectType](/src/target/wwsgen_types.go?s=43723:43951#L687)
``` go
type CompensationBasisObjectType struct {
    ID         []CompensationBasisObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationBenchmarkAmountReplacmentDataType">type</a> [CompensationBenchmarkAmountReplacmentDataType](/src/target/wwsgen_types.go?s=43985:44260#L693)
``` go
type CompensationBenchmarkAmountReplacmentDataType struct {
    PercentileReference *CompensationBenchmarkPercentileObjectType `xml:"urn:com.workday/bsvc Percentile_Reference"`
    Amount              float64                                    `xml:"urn:com.workday/bsvc Amount"`
}
```
Benchmark Job Composite Data










## <a name="CompensationBenchmarkCompositeCategoryObjectIDType">type</a> [CompensationBenchmarkCompositeCategoryObjectIDType](/src/target/wwsgen_types.go?s=44324:44475#L699)
``` go
type CompensationBenchmarkCompositeCategoryObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationBenchmarkCompositeCategoryObjectType">type</a> [CompensationBenchmarkCompositeCategoryObjectType](/src/target/wwsgen_types.go?s=44477:44768#L704)
``` go
type CompensationBenchmarkCompositeCategoryObjectType struct {
    ID         []CompensationBenchmarkCompositeCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationBenchmarkDefaultObjectIDType">type</a> [CompensationBenchmarkDefaultObjectIDType](/src/target/wwsgen_types.go?s=44832:44973#L710)
``` go
type CompensationBenchmarkDefaultObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationBenchmarkDefaultObjectType">type</a> [CompensationBenchmarkDefaultObjectType](/src/target/wwsgen_types.go?s=44975:45236#L715)
``` go
type CompensationBenchmarkDefaultObjectType struct {
    ID         []CompensationBenchmarkDefaultObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationBenchmarkPercentileObjectIDType">type</a> [CompensationBenchmarkPercentileObjectIDType](/src/target/wwsgen_types.go?s=45300:45444#L721)
``` go
type CompensationBenchmarkPercentileObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationBenchmarkPercentileObjectType">type</a> [CompensationBenchmarkPercentileObjectType](/src/target/wwsgen_types.go?s=45446:45716#L726)
``` go
type CompensationBenchmarkPercentileObjectType struct {
    ID         []CompensationBenchmarkPercentileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationBenchmarkProfileDataType">type</a> [CompensationBenchmarkProfileDataType](/src/target/wwsgen_types.go?s=45740:46338#L732)
``` go
type CompensationBenchmarkProfileDataType struct {
    CompensationEligibilityRuleReference *ConditionRuleObjectType        `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule_Reference"`
    CurrencyReference                    *CurrencyObjectType             `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
    FrequencyReference                   *FrequencyObjectType            `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
    BenchmarkJobCompositeReplacementData []BenchmarkJobCompositeDataType `xml:"urn:com.workday/bsvc Benchmark_Job_Composite_Replacement_Data,omitempty"`
}
```
Benchmark Job Data










## <a name="CompensationChangeDataType">type</a> [CompensationChangeDataType](/src/target/wwsgen_types.go?s=46407:49091#L740)
``` go
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
```
Data element that contains the compensation change information.










## <a name="CompensationDurationSeveranceMatrixDataType">type</a> [CompensationDurationSeveranceMatrixDataType](/src/target/wwsgen_types.go?s=49153:50010#L761)
``` go
type CompensationDurationSeveranceMatrixDataType struct {
    CompensationBasisReference                   *CompensationBasisObjectType                       `xml:"urn:com.workday/bsvc Compensation_Basis_Reference"`
    CalculateCompensationBasisAsOfDateReference  *SeveranceServiceDateObjectType                    `xml:"urn:com.workday/bsvc Calculate_Compensation_Basis_As_Of_Date_Reference"`
    CurrencyReference                            *CurrencyObjectType                                `xml:"urn:com.workday/bsvc Currency_Reference"`
    FrequencyReference                           *FrequencyObjectType                               `xml:"urn:com.workday/bsvc Frequency_Reference"`
    CompensationDurationSeveranceMatrixEntryData []CompensationDurationSeveranceMatrixEntryDataType `xml:"urn:com.workday/bsvc Compensation_Duration_Severance_Matrix_Entry_Data"`
}
```
The compensation duration severance matrix data element.










## <a name="CompensationDurationSeveranceMatrixEntryDataType">type</a> [CompensationDurationSeveranceMatrixEntryDataType](/src/target/wwsgen_types.go?s=50046:50806#L770)
``` go
type CompensationDurationSeveranceMatrixEntryDataType struct {
    MinimumCompensationBasisRange float64                  `xml:"urn:com.workday/bsvc Minimum_Compensation_Basis_Range,omitempty"`
    MaximumCompensationBasisRange float64                  `xml:"urn:com.workday/bsvc Maximum_Compensation_Basis_Range,omitempty"`
    EligibilityRuleReference      *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    Duration                      float64                  `xml:"urn:com.workday/bsvc Duration,omitempty"`
    MinimumDuration               float64                  `xml:"urn:com.workday/bsvc Minimum_Duration,omitempty"`
    MaximumDuration               float64                  `xml:"urn:com.workday/bsvc Maximum_Duration,omitempty"`
}
```
The data for the matrix entry.










## <a name="CompensationEligibilityRuleDataType">type</a> [CompensationEligibilityRuleDataType](/src/target/wwsgen_types.go?s=50855:51722#L780)
``` go
type CompensationEligibilityRuleDataType struct {
    EffectiveDate                     *time.Time                           `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
    ConditionRuleData                 *ConditionRuleDataWWSType            `xml:"urn:com.workday/bsvc Condition_Rule_Data"`
    CompensationPackageReference      []CompensationPackageObjectType      `xml:"urn:com.workday/bsvc Compensation_Package_Reference,omitempty"`
    CompensationGradeReference        []CompensationGradeObjectType        `xml:"urn:com.workday/bsvc Compensation_Grade_Reference,omitempty"`
    CompensationGradeProfileReference []CompensationGradeProfileObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Profile_Reference,omitempty"`
    CompensationPlanReference         []CompensationPlanObjectType         `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
}
```
Data for this Compensation Eligibility Rule










### <a name="CompensationEligibilityRuleDataType.MarshalXML">func</a> (\*CompensationEligibilityRuleDataType) [MarshalXML](/src/target/wwsgen_types.go?s=51724:51826#L789)
``` go
func (t *CompensationEligibilityRuleDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="CompensationEligibilityRuleDataType.UnmarshalXML">func</a> (\*CompensationEligibilityRuleDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=52101:52205#L799)
``` go
func (t *CompensationEligibilityRuleDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="CompensationEligibilityRuleRequestCriteriaType">type</a> [CompensationEligibilityRuleRequestCriteriaType](/src/target/wwsgen_types.go?s=52562:52624#L811)
``` go
type CompensationEligibilityRuleRequestCriteriaType struct {
}
```
Request Criteria for this request. Note that there are none available.










## <a name="CompensationEligibilityRuleRequestReferencesType">type</a> [CompensationEligibilityRuleRequestReferencesType](/src/target/wwsgen_types.go?s=52690:52887#L815)
``` go
type CompensationEligibilityRuleRequestReferencesType struct {
    CompensationEligibilityRuleReference []ConditionRuleObjectType `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule_Reference"`
}
```
Compensation Eligibility Rule References for rules to return










## <a name="CompensationEligibilityRuleResponseDataType">type</a> [CompensationEligibilityRuleResponseDataType](/src/target/wwsgen_types.go?s=52936:53127#L820)
``` go
type CompensationEligibilityRuleResponseDataType struct {
    CompensationEligibilityRule []CompensationEligibilityRuleType `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule,omitempty"`
}
```
Compensation Eligibility Rule Response Data










## <a name="CompensationEligibilityRuleResponseGroupType">type</a> [CompensationEligibilityRuleResponseGroupType](/src/target/wwsgen_types.go?s=53184:53352#L825)
``` go
type CompensationEligibilityRuleResponseGroupType struct {
    HideConfigurationDependencies *bool `xml:"urn:com.workday/bsvc Hide_Configuration_Dependencies,omitempty"`
}
```
Provides the filters for the web service operation.










## <a name="CompensationEligibilityRuleType">type</a> [CompensationEligibilityRuleType](/src/target/wwsgen_types.go?s=53417:53769#L830)
``` go
type CompensationEligibilityRuleType struct {
    CompensationEligibilityRuleReference *ConditionRuleObjectType              `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule_Reference,omitempty"`
    CompensationEligibilityRuleData      []CompensationEligibilityRuleDataType `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule_Data,omitempty"`
}
```
Response data element for the Compensation Eligibility Rule










## <a name="CompensationGradeDataType">type</a> [CompensationGradeDataType](/src/target/wwsgen_types.go?s=53808:55773#L836)
``` go
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
```
Contains compensation grade data.










### <a name="CompensationGradeDataType.MarshalXML">func</a> (\*CompensationGradeDataType) [MarshalXML](/src/target/wwsgen_types.go?s=55775:55867#L852)
``` go
func (t *CompensationGradeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="CompensationGradeDataType.UnmarshalXML">func</a> (\*CompensationGradeDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=56132:56226#L862)
``` go
func (t *CompensationGradeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="CompensationGradeHierarchyDataType">type</a> [CompensationGradeHierarchyDataType](/src/target/wwsgen_types.go?s=56499:56763#L873)
``` go
type CompensationGradeHierarchyDataType struct {
    EffectiveDate          time.Time                    `xml:"urn:com.workday/bsvc Effective_Date"`
    CompensationGradeLevel []CompensationGradeLevelType `xml:"urn:com.workday/bsvc Compensation_Grade_Level,omitempty"`
}
```









### <a name="CompensationGradeHierarchyDataType.MarshalXML">func</a> (\*CompensationGradeHierarchyDataType) [MarshalXML](/src/target/wwsgen_types.go?s=56765:56866#L878)
``` go
func (t *CompensationGradeHierarchyDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="CompensationGradeHierarchyDataType.UnmarshalXML">func</a> (\*CompensationGradeHierarchyDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=57131:57234#L888)
``` go
func (t *CompensationGradeHierarchyDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="CompensationGradeHierarchyRequestReferencesType">type</a> [CompensationGradeHierarchyRequestReferencesType](/src/target/wwsgen_types.go?s=57630:57762#L900)
``` go
type CompensationGradeHierarchyRequestReferencesType struct {
    EffectiveDate time.Time `xml:"urn:com.workday/bsvc Effective_Date"`
}
```
Utilize the compensation grade hierarchy element to retrieve specific instance(s) of compensation grade hierarchy data.










### <a name="CompensationGradeHierarchyRequestReferencesType.MarshalXML">func</a> (\*CompensationGradeHierarchyRequestReferencesType) [MarshalXML](/src/target/wwsgen_types.go?s=57764:57878#L904)
``` go
func (t *CompensationGradeHierarchyRequestReferencesType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="CompensationGradeHierarchyRequestReferencesType.UnmarshalXML">func</a> (\*CompensationGradeHierarchyRequestReferencesType) [UnmarshalXML](/src/target/wwsgen_types.go?s=58156:58272#L914)
``` go
func (t *CompensationGradeHierarchyRequestReferencesType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="CompensationGradeHierarchyResponseDataType">type</a> [CompensationGradeHierarchyResponseDataType](/src/target/wwsgen_types.go?s=58558:58745#L925)
``` go
type CompensationGradeHierarchyResponseDataType struct {
    CompensationGradeHierarchy []CompensationGradeHierarchyType `xml:"urn:com.workday/bsvc Compensation_Grade_Hierarchy,omitempty"`
}
```









## <a name="CompensationGradeHierarchyType">type</a> [CompensationGradeHierarchyType](/src/target/wwsgen_types.go?s=58747:58934#L929)
``` go
type CompensationGradeHierarchyType struct {
    CompensationGradeHierarchyData *CompensationGradeHierarchyDataType `xml:"urn:com.workday/bsvc Compensation_Grade_Hierarchy_Data,omitempty"`
}
```









## <a name="CompensationGradeLevelDataType">type</a> [CompensationGradeLevelDataType](/src/target/wwsgen_types.go?s=58936:59382#L933)
``` go
type CompensationGradeLevelDataType struct {
    ID                         string                        `xml:"urn:com.workday/bsvc ID,omitempty"`
    Rank                       float64                       `xml:"urn:com.workday/bsvc Rank"`
    Name                       string                        `xml:"urn:com.workday/bsvc Name"`
    CompensationGradeReference []CompensationGradeObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Reference"`
}
```









## <a name="CompensationGradeLevelObjectIDType">type</a> [CompensationGradeLevelObjectIDType](/src/target/wwsgen_types.go?s=59446:59581#L941)
``` go
type CompensationGradeLevelObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationGradeLevelObjectType">type</a> [CompensationGradeLevelObjectType](/src/target/wwsgen_types.go?s=59583:59826#L946)
``` go
type CompensationGradeLevelObjectType struct {
    ID         []CompensationGradeLevelObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationGradeLevelType">type</a> [CompensationGradeLevelType](/src/target/wwsgen_types.go?s=59828:60147#L951)
``` go
type CompensationGradeLevelType struct {
    CompensationGradeLevelReference *CompensationGradeLevelObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Level_Reference,omitempty"`
    CompensationGradeLevelData      *CompensationGradeLevelDataType   `xml:"urn:com.workday/bsvc Compensation_Grade_Level_Data,omitempty"`
}
```









## <a name="CompensationGradeObjectIDType">type</a> [CompensationGradeObjectIDType](/src/target/wwsgen_types.go?s=60211:60341#L957)
``` go
type CompensationGradeObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationGradeObjectType">type</a> [CompensationGradeObjectType](/src/target/wwsgen_types.go?s=60343:60571#L962)
``` go
type CompensationGradeObjectType struct {
    ID         []CompensationGradeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationGradeProfileDataType">type</a> [CompensationGradeProfileDataType](/src/target/wwsgen_types.go?s=60618:62008#L968)
``` go
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
```
Contains compensation grade profile data.










### <a name="CompensationGradeProfileDataType.MarshalXML">func</a> (\*CompensationGradeProfileDataType) [MarshalXML](/src/target/wwsgen_types.go?s=62010:62109#L981)
``` go
func (t *CompensationGradeProfileDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="CompensationGradeProfileDataType.UnmarshalXML">func</a> (\*CompensationGradeProfileDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=62381:62482#L991)
``` go
func (t *CompensationGradeProfileDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="CompensationGradeProfileObjectIDType">type</a> [CompensationGradeProfileObjectIDType](/src/target/wwsgen_types.go?s=62824:62961#L1003)
``` go
type CompensationGradeProfileObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationGradeProfileObjectType">type</a> [CompensationGradeProfileObjectType](/src/target/wwsgen_types.go?s=62963:63212#L1008)
``` go
type CompensationGradeProfileObjectType struct {
    ID         []CompensationGradeProfileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationGradeProfileType">type</a> [CompensationGradeProfileType](/src/target/wwsgen_types.go?s=63280:63735#L1014)
``` go
type CompensationGradeProfileType struct {
    CompensationGradeProfileReference *CompensationGradeProfileObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Profile_Reference,omitempty"`
    CompensationGradeProfileData      *CompensationGradeProfileDataType   `xml:"urn:com.workday/bsvc Compensation_Grade_Profile_Data,omitempty"`
    Delete                            bool                                `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}
```
Contains a compensation grade profile and its associated data.










## <a name="CompensationGradeRequestCriteriaType">type</a> [CompensationGradeRequestCriteriaType](/src/target/wwsgen_types.go?s=63781:63912#L1021)
``` go
type CompensationGradeRequestCriteriaType struct {
    IncludeInactive *bool `xml:"urn:com.workday/bsvc Include_Inactive,omitempty"`
}
```
Get Compensation Grades Request Criteria










## <a name="CompensationGradeRequestReferencesType">type</a> [CompensationGradeRequestReferencesType](/src/target/wwsgen_types.go?s=64036:64206#L1026)
``` go
type CompensationGradeRequestReferencesType struct {
    CompensationGradeReference []CompensationGradeObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Reference"`
}
```
Utilize the compensation grade element to retrieve specific instance(s) of compensation grade and its associated data.










## <a name="CompensationGradeResponseDataType">type</a> [CompensationGradeResponseDataType](/src/target/wwsgen_types.go?s=64300:64450#L1031)
``` go
type CompensationGradeResponseDataType struct {
    CompensationGrade []CompensationGradeType `xml:"urn:com.workday/bsvc Compensation_Grade,omitempty"`
}
```
Contains compensation grade information based on Request References or Request Criteria.










## <a name="CompensationGradeResponseGroupType">type</a> [CompensationGradeResponseGroupType](/src/target/wwsgen_types.go?s=64507:64638#L1036)
``` go
type CompensationGradeResponseGroupType struct {
    IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}
```
Provides the filters for the web service operation.










## <a name="CompensationGradeType">type</a> [CompensationGradeType](/src/target/wwsgen_types.go?s=64698:64980#L1041)
``` go
type CompensationGradeType struct {
    CompensationGradeReference *CompensationGradeObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Reference,omitempty"`
    CompensationGradeData      *CompensationGradeDataType   `xml:"urn:com.workday/bsvc Compensation_Grade_Data,omitempty"`
}
```
Contains a compensation grade and its associated data.










## <a name="CompensationLengthofServiceSeveranceMatrixDataType">type</a> [CompensationLengthofServiceSeveranceMatrixDataType](/src/target/wwsgen_types.go?s=65051:66338#L1047)
``` go
type CompensationLengthofServiceSeveranceMatrixDataType struct {
    CompensationBasisReference                          *CompensationBasisObjectType                              `xml:"urn:com.workday/bsvc Compensation_Basis_Reference"`
    CalculateCompensationBasisAsOfDateReference         *SeveranceServiceDateObjectType                           `xml:"urn:com.workday/bsvc Calculate_Compensation_Basis_As_Of_Date_Reference"`
    CurrencyReference                                   *CurrencyObjectType                                       `xml:"urn:com.workday/bsvc Currency_Reference"`
    FrequencyReference                                  *FrequencyObjectType                                      `xml:"urn:com.workday/bsvc Frequency_Reference"`
    CompensationRoundingRuleReference                   *CompensationRoundingRuleObjectType                       `xml:"urn:com.workday/bsvc Compensation_Rounding_Rule_Reference"`
    MultiplierOrderReference                            *BenefitMultiplierOrderObjectType                         `xml:"urn:com.workday/bsvc Multiplier_Order_Reference"`
    CompensationLengthofServiceSeveranceMatrixEntryData []CompensationLengthofServiceSeveranceMatrixEntryDataType `xml:"urn:com.workday/bsvc Compensation_Length_of_Service_Severance_Matrix_Entry_Data"`
}
```
The compensation length of service severance matrix data element.










## <a name="CompensationLengthofServiceSeveranceMatrixEntryDataType">type</a> [CompensationLengthofServiceSeveranceMatrixEntryDataType](/src/target/wwsgen_types.go?s=66374:67161#L1058)
``` go
type CompensationLengthofServiceSeveranceMatrixEntryDataType struct {
    MinimumCompensationBasisRange float64                  `xml:"urn:com.workday/bsvc Minimum_Compensation_Basis_Range,omitempty"`
    MaximumCompensationBasisRange float64                  `xml:"urn:com.workday/bsvc Maximum_Compensation_Basis_Range,omitempty"`
    EligibilityRuleReference      *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    LengthofServiceMultiplier     float64                  `xml:"urn:com.workday/bsvc Length_of_Service_Multiplier,omitempty"`
    MinimumDuration               float64                  `xml:"urn:com.workday/bsvc Minimum_Duration,omitempty"`
    MaximumDuration               float64                  `xml:"urn:com.workday/bsvc Maximum_Duration,omitempty"`
}
```
The data for the matrix entry.










## <a name="CompensationMatrixAmountBasedDataType">type</a> [CompensationMatrixAmountBasedDataType](/src/target/wwsgen_types.go?s=67218:67899#L1068)
``` go
type CompensationMatrixAmountBasedDataType struct {
    DefaultFrequencyReference              *FrequencyObjectType                         `xml:"urn:com.workday/bsvc Default_Frequency_Reference"`
    DefaultCurrencyReference               *CurrencyObjectType                          `xml:"urn:com.workday/bsvc Default_Currency_Reference,omitempty"`
    CompensationMatrixTargetRulesReference []ConditionRuleObjectType                    `xml:"urn:com.workday/bsvc Compensation_Matrix_Target_Rules_Reference,omitempty"`
    CompensationMatrixEntryAmountBasedData []CompensationMatrixEntryAmountBasedDataType `xml:"urn:com.workday/bsvc Compensation_Matrix_Entry_Amount_Based_Data,omitempty"`
}
```
The amount based data for this compensation matrix.










## <a name="CompensationMatrixDataType">type</a> [CompensationMatrixDataType](/src/target/wwsgen_types.go?s=67943:69316#L1076)
``` go
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
```
All data for this compensation matrix.










### <a name="CompensationMatrixDataType.MarshalXML">func</a> (\*CompensationMatrixDataType) [MarshalXML](/src/target/wwsgen_types.go?s=69318:69411#L1087)
``` go
func (t *CompensationMatrixDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="CompensationMatrixDataType.UnmarshalXML">func</a> (\*CompensationMatrixDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=69677:69772#L1097)
``` go
func (t *CompensationMatrixDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="CompensationMatrixEntryAmountBasedDataType">type</a> [CompensationMatrixEntryAmountBasedDataType](/src/target/wwsgen_types.go?s=70094:71525#L1109)
``` go
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
```
Data for a single compensation matrix entry.










## <a name="CompensationMatrixEntryPercentBasedDataType">type</a> [CompensationMatrixEntryPercentBasedDataType](/src/target/wwsgen_types.go?s=71594:73009#L1123)
``` go
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
```
Data for a single entry in a percent based compensation matrix.










## <a name="CompensationMatrixNonweightedPercentBasedDataType">type</a> [CompensationMatrixNonweightedPercentBasedDataType](/src/target/wwsgen_types.go?s=73073:73478#L1136)
``` go
type CompensationMatrixNonweightedPercentBasedDataType struct {
    CompensationMatrixTargetRulesReference  []ConditionRuleObjectType                     `xml:"urn:com.workday/bsvc Compensation_Matrix_Target_Rules_Reference,omitempty"`
    CompensationMatrixEntryPercentBasedData []CompensationMatrixEntryPercentBasedDataType `xml:"urn:com.workday/bsvc Compensation_Matrix_Entry_Percent_Based_Data,omitempty"`
}
```
Data for a non-weighted percent based compensation matrix.










## <a name="CompensationMatrixObjectIDType">type</a> [CompensationMatrixObjectIDType](/src/target/wwsgen_types.go?s=73542:73673#L1142)
``` go
type CompensationMatrixObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationMatrixObjectType">type</a> [CompensationMatrixObjectType](/src/target/wwsgen_types.go?s=73675:73906#L1147)
``` go
type CompensationMatrixObjectType struct {
    ID         []CompensationMatrixObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationMatrixRequestCriteriaType">type</a> [CompensationMatrixRequestCriteriaType](/src/target/wwsgen_types.go?s=73928:74425#L1153)
``` go
type CompensationMatrixRequestCriteriaType struct {
    Name                                      string `xml:"urn:com.workday/bsvc Name"`
    CompensationMatrixAmountBased             bool   `xml:"urn:com.workday/bsvc Compensation_Matrix_Amount_Based"`
    CompensationMatrixNonweightedPercentBased bool   `xml:"urn:com.workday/bsvc Compensation_Matrix_Non-weighted_Percent_Based"`
    CompensationMatrixWeightedPercentBased    bool   `xml:"urn:com.workday/bsvc Compensation_Matrix_Weighted_Percent_Based"`
}
```
Request criteria










## <a name="CompensationMatrixRequestReferencesType">type</a> [CompensationMatrixRequestReferencesType](/src/target/wwsgen_types.go?s=74458:74646#L1161)
``` go
type CompensationMatrixRequestReferencesType struct {
    CompensationMatricesReference []CompensationMatrixObjectType `xml:"urn:com.workday/bsvc Compensation_Matrices_Reference,omitempty"`
}
```
Request references element.










## <a name="CompensationMatrixResponseDataType">type</a> [CompensationMatrixResponseDataType](/src/target/wwsgen_types.go?s=74674:74828#L1166)
``` go
type CompensationMatrixResponseDataType struct {
    CompensationMatrix []CompensationMatrixType `xml:"urn:com.workday/bsvc Compensation_Matrix,omitempty"`
}
```
Response Data element.










## <a name="CompensationMatrixResponseGroupType">type</a> [CompensationMatrixResponseGroupType](/src/target/wwsgen_types.go?s=74853:75107#L1171)
``` go
type CompensationMatrixResponseGroupType struct {
    IncludeReference              *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
    IncludeCompensationMatrixData *bool `xml:"urn:com.workday/bsvc Include_Compensation_Matrix_Data,omitempty"`
}
```
The response group.










## <a name="CompensationMatrixType">type</a> [CompensationMatrixType](/src/target/wwsgen_types.go?s=75141:75430#L1177)
``` go
type CompensationMatrixType struct {
    CompensationMatrixReference *CompensationMatrixObjectType `xml:"urn:com.workday/bsvc Compensation_Matrix_Reference,omitempty"`
    CompensationMatrixData      []CompensationMatrixDataType  `xml:"urn:com.workday/bsvc Compensation_Matrix_Data,omitempty"`
}
```
Compensation Matrix element.










## <a name="CompensationMatrixWeightedPercentBasedDataType">type</a> [CompensationMatrixWeightedPercentBasedDataType](/src/target/wwsgen_types.go?s=75491:75893#L1183)
``` go
type CompensationMatrixWeightedPercentBasedDataType struct {
    CompensationMatrixTargetRulesReference  []ConditionRuleObjectType                     `xml:"urn:com.workday/bsvc Compensation_Matrix_Target_Rules_Reference,omitempty"`
    CompensationMatrixEntryPercentBasedData []CompensationMatrixEntryPercentBasedDataType `xml:"urn:com.workday/bsvc Compensation_Matrix_Entry_Percent_Based_Data,omitempty"`
}
```
Data for this weight percent based compensation matrix.










## <a name="CompensationPackageObjectIDType">type</a> [CompensationPackageObjectIDType](/src/target/wwsgen_types.go?s=75957:76089#L1189)
``` go
type CompensationPackageObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationPackageObjectType">type</a> [CompensationPackageObjectType](/src/target/wwsgen_types.go?s=76091:76325#L1194)
``` go
type CompensationPackageObjectType struct {
    ID         []CompensationPackageObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationPayEarningObjectIDType">type</a> [CompensationPayEarningObjectIDType](/src/target/wwsgen_types.go?s=76389:76524#L1200)
``` go
type CompensationPayEarningObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationPayEarningObjectType">type</a> [CompensationPayEarningObjectType](/src/target/wwsgen_types.go?s=76526:76769#L1205)
``` go
type CompensationPayEarningObjectType struct {
    ID         []CompensationPayEarningObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationPayRangeDataType">type</a> [CompensationPayRangeDataType](/src/target/wwsgen_types.go?s=76802:78336#L1211)
``` go
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
```
The compensation pay range.










## <a name="CompensationPeriodObjectIDType">type</a> [CompensationPeriodObjectIDType](/src/target/wwsgen_types.go?s=78400:78531#L1230)
``` go
type CompensationPeriodObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationPeriodObjectType">type</a> [CompensationPeriodObjectType](/src/target/wwsgen_types.go?s=78533:78764#L1235)
``` go
type CompensationPeriodObjectType struct {
    ID         []CompensationPeriodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationPlanDataType">type</a> [CompensationPlanDataType](/src/target/wwsgen_types.go?s=78999:80579#L1241)
``` go
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
```
Contains common information for the following Compensation Plans:  Future Payment, Bonus, Merit, Allowance, Commission, and Base Pay.  It also must contain a Future Play, Bonus, Merit, Allowance, Commission, or Bay Pay Plan data.










### <a name="CompensationPlanDataType.MarshalXML">func</a> (\*CompensationPlanDataType) [MarshalXML](/src/target/wwsgen_types.go?s=80581:80672#L1259)
``` go
func (t *CompensationPlanDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="CompensationPlanDataType.UnmarshalXML">func</a> (\*CompensationPlanDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=80936:81029#L1269)
``` go
func (t *CompensationPlanDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="CompensationPlanObjectIDType">type</a> [CompensationPlanObjectIDType](/src/target/wwsgen_types.go?s=81363:81492#L1281)
``` go
type CompensationPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationPlanObjectType">type</a> [CompensationPlanObjectType](/src/target/wwsgen_types.go?s=81494:81719#L1286)
``` go
type CompensationPlanObjectType struct {
    ID         []CompensationPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationPlanRequestCriteriaType">type</a> [CompensationPlanRequestCriteriaType](/src/target/wwsgen_types.go?s=81810:82110#L1292)
``` go
type CompensationPlanRequestCriteriaType struct {
    PlanTypeReference []CompensationAssignableComponentTypeObjectType `xml:"urn:com.workday/bsvc Plan_Type_Reference,omitempty"`
    IncludeInactive   *bool                                           `xml:"urn:com.workday/bsvc Include_Inactive,omitempty"`
}
```
Values in this element can be used to filter the types of Compensation Plans returned










## <a name="CompensationPlanRequestReferencesType">type</a> [CompensationPlanRequestReferencesType](/src/target/wwsgen_types.go?s=82201:82377#L1298)
``` go
type CompensationPlanRequestReferencesType struct {
    CompensationPlanReference []CompensationAssignablePlanObjectType `xml:"urn:com.workday/bsvc Compensation_Plan_Reference"`
}
```
Use this element to request specific Compensation Plans given the reference ID values










## <a name="CompensationPlanResponseDataType">type</a> [CompensationPlanResponseDataType](/src/target/wwsgen_types.go?s=82467:82603#L1303)
``` go
type CompensationPlanResponseDataType struct {
    CompensationPlan []CompensationPlanType `xml:"urn:com.workday/bsvc Compensation_Plan"`
}
```
Response element containing instances of compensation plans and the associated data.










## <a name="CompensationPlanResponseGroupType">type</a> [CompensationPlanResponseGroupType](/src/target/wwsgen_types.go?s=82669:82799#L1308)
``` go
type CompensationPlanResponseGroupType struct {
    IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}
```
Use to limit the returned data for a given Compensation Plan










## <a name="CompensationPlanType">type</a> [CompensationPlanType](/src/target/wwsgen_types.go?s=82860:83155#L1313)
``` go
type CompensationPlanType struct {
    CompensationPlanReference *CompensationAssignablePlanObjectType `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
    CompensationPlanData      *CompensationPlanDataType             `xml:"urn:com.workday/bsvc Compensation_Plan_Data,omitempty"`
}
```
An instance of a compensation plan and associated data.










## <a name="CompensationPreviousSystemHistoryObjectIDType">type</a> [CompensationPreviousSystemHistoryObjectIDType](/src/target/wwsgen_types.go?s=83219:83365#L1319)
``` go
type CompensationPreviousSystemHistoryObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationPreviousSystemHistoryObjectType">type</a> [CompensationPreviousSystemHistoryObjectType](/src/target/wwsgen_types.go?s=83367:83643#L1324)
``` go
type CompensationPreviousSystemHistoryObjectType struct {
    ID         []CompensationPreviousSystemHistoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationRoundingRuleObjectIDType">type</a> [CompensationRoundingRuleObjectIDType](/src/target/wwsgen_types.go?s=83707:83844#L1330)
``` go
type CompensationRoundingRuleObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationRoundingRuleObjectType">type</a> [CompensationRoundingRuleObjectType](/src/target/wwsgen_types.go?s=83846:84095#L1335)
``` go
type CompensationRoundingRuleObjectType struct {
    ID         []CompensationRoundingRuleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationScorecardDataType">type</a> [CompensationScorecardDataType](/src/target/wwsgen_types.go?s=84125:84785#L1341)
``` go
type CompensationScorecardDataType struct {
    ID                   string                        `xml:"urn:com.workday/bsvc ID,omitempty"`
    EffectiveDate        time.Time                     `xml:"urn:com.workday/bsvc Effective_Date"`
    ScorecardName        string                        `xml:"urn:com.workday/bsvc Scorecard_Name"`
    ScorecardDescription string                        `xml:"urn:com.workday/bsvc Scorecard_Description,omitempty"`
    ScorecardGoalsData   []PerformanceCriteriaDataType `xml:"urn:com.workday/bsvc Scorecard_Goals_Data"`
    ScorecardProfileData []ScorecardProfileDataType    `xml:"urn:com.workday/bsvc Scorecard_Profile_Data,omitempty"`
}
```
Contains scorecard data.










### <a name="CompensationScorecardDataType.MarshalXML">func</a> (\*CompensationScorecardDataType) [MarshalXML](/src/target/wwsgen_types.go?s=84787:84883#L1350)
``` go
func (t *CompensationScorecardDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="CompensationScorecardDataType.UnmarshalXML">func</a> (\*CompensationScorecardDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=85143:85241#L1360)
``` go
func (t *CompensationScorecardDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="CompensationScorecardRequestCriteriaType">type</a> [CompensationScorecardRequestCriteriaType](/src/target/wwsgen_types.go?s=85631:85687#L1372)
``` go
type CompensationScorecardRequestCriteriaType struct {
}
```
Contains additional request criteria to limit the information returned. Currently no additional criteria is supported.










## <a name="CompensationScorecardRequestReferencesType">type</a> [CompensationScorecardRequestReferencesType](/src/target/wwsgen_types.go?s=85750:85931#L1376)
``` go
type CompensationScorecardRequestReferencesType struct {
    CompensationScorecardReference []DefaultScorecardObjectType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Reference"`
}
```
A container for one or more unique scorecard identifiers.










## <a name="CompensationScorecardResponseDataType">type</a> [CompensationScorecardResponseDataType](/src/target/wwsgen_types.go?s=86087:86253#L1381)
``` go
type CompensationScorecardResponseDataType struct {
    CompensationScorecard []CompensationScorecardType `xml:"urn:com.workday/bsvc Compensation_Scorecard,omitempty"`
}
```
Contains the scorecard response information. This includes the unique scorecard identifier and the detailed scorecard data which was added or updated.










## <a name="CompensationScorecardResponseGroupType">type</a> [CompensationScorecardResponseGroupType](/src/target/wwsgen_types.go?s=86333:86468#L1386)
``` go
type CompensationScorecardResponseGroupType struct {
    IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}
```
Contains additional response criteria to enhance the information returned.










## <a name="CompensationScorecardResultDataType">type</a> [CompensationScorecardResultDataType](/src/target/wwsgen_types.go?s=86525:87230#L1391)
``` go
type CompensationScorecardResultDataType struct {
    ID                             string                                       `xml:"urn:com.workday/bsvc ID,omitempty"`
    EvaluationDate                 time.Time                                    `xml:"urn:com.workday/bsvc Evaluation_Date"`
    CompensationScorecardReference *DefaultScorecardObjectType                  `xml:"urn:com.workday/bsvc Compensation_Scorecard_Reference"`
    ScorecardResultData            []ScorecardResultDataType                    `xml:"urn:com.workday/bsvc Scorecard_Result_Data"`
    ProfileScorecardResultData     []ProfileCompensationScorecardResultDataType `xml:"urn:com.workday/bsvc Profile_Scorecard_Result_Data,omitempty"`
}
```
Contains the detailed scorecard result information.










### <a name="CompensationScorecardResultDataType.MarshalXML">func</a> (\*CompensationScorecardResultDataType) [MarshalXML](/src/target/wwsgen_types.go?s=87232:87334#L1399)
``` go
func (t *CompensationScorecardResultDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="CompensationScorecardResultDataType.UnmarshalXML">func</a> (\*CompensationScorecardResultDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=87604:87708#L1409)
``` go
func (t *CompensationScorecardResultDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="CompensationScorecardResultRequestCriteriaType">type</a> [CompensationScorecardResultRequestCriteriaType](/src/target/wwsgen_types.go?s=88109:88171#L1421)
``` go
type CompensationScorecardResultRequestCriteriaType struct {
}
```
Contains additional request criteria to limit the information returned.  Currently no additional criteria is supported.










## <a name="CompensationScorecardResultRequestReferencesType">type</a> [CompensationScorecardResultRequestReferencesType](/src/target/wwsgen_types.go?s=88241:88442#L1425)
``` go
type CompensationScorecardResultRequestReferencesType struct {
    CompensationScorecardResultReference []ScoresetContainerObjectType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Result_Reference"`
}
```
A container for one or more unique scorecard result identifiers.










## <a name="CompensationScorecardResultResponseDataType">type</a> [CompensationScorecardResultResponseDataType](/src/target/wwsgen_types.go?s=88619:88810#L1430)
``` go
type CompensationScorecardResultResponseDataType struct {
    CompensationScorecardResult []CompensationScorecardResultType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Result,omitempty"`
}
```
Contains the scorecard result response information. This includes the unique scorecard result identifier and the detailed scorecard result data which was added or updated.










## <a name="CompensationScorecardResultResponseGroupType">type</a> [CompensationScorecardResultResponseGroupType](/src/target/wwsgen_types.go?s=88890:89031#L1435)
``` go
type CompensationScorecardResultResponseGroupType struct {
    IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}
```
Contains additional response criteria to enhance the information returned.










## <a name="CompensationScorecardResultType">type</a> [CompensationScorecardResultType](/src/target/wwsgen_types.go?s=89208:89560#L1440)
``` go
type CompensationScorecardResultType struct {
    CompensationScorecardResultReference *ScoresetContainerObjectType          `xml:"urn:com.workday/bsvc Compensation_Scorecard_Result_Reference,omitempty"`
    CompensationScorecardResultData      []CompensationScorecardResultDataType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Result_Data,omitempty"`
}
```
Contains the scorecard result response information. This includes the unique scorecard result identifier and the detailed scorecard result data which was added or updated.










## <a name="CompensationScorecardType">type</a> [CompensationScorecardType](/src/target/wwsgen_types.go?s=89590:89898#L1446)
``` go
type CompensationScorecardType struct {
    CompensationScorecardReference *DefaultScorecardObjectType     `xml:"urn:com.workday/bsvc Compensation_Scorecard_Reference,omitempty"`
    CompensationScorecardData      []CompensationScorecardDataType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Data,omitempty"`
}
```
Contains scorecard data.










## <a name="CompensationSetupSecuritySegmentObjectIDType">type</a> [CompensationSetupSecuritySegmentObjectIDType](/src/target/wwsgen_types.go?s=89962:90107#L1452)
``` go
type CompensationSetupSecuritySegmentObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationSetupSecuritySegmentObjectType">type</a> [CompensationSetupSecuritySegmentObjectType](/src/target/wwsgen_types.go?s=90109:90382#L1457)
``` go
type CompensationSetupSecuritySegmentObjectType struct {
    ID         []CompensationSetupSecuritySegmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationStepDataType">type</a> [CompensationStepDataType](/src/target/wwsgen_types.go?s=90407:91170#L1463)
``` go
type CompensationStepDataType struct {
    CompensationStepID       string                       `xml:"urn:com.workday/bsvc Compensation_Step_ID,omitempty"`
    Sequence                 string                       `xml:"urn:com.workday/bsvc Sequence"`
    Name                     string                       `xml:"urn:com.workday/bsvc Name"`
    Amount                   float64                      `xml:"urn:com.workday/bsvc Amount"`
    Interval                 float64                      `xml:"urn:com.workday/bsvc Interval,omitempty"`
    PeriodReference          *FrequencyBehaviorObjectType `xml:"urn:com.workday/bsvc Period_Reference,omitempty"`
    ProgressionRuleReference *ConditionRuleObjectType     `xml:"urn:com.workday/bsvc Progression_Rule_Reference,omitempty"`
}
```
A compensation step










## <a name="CompensationStepObjectIDType">type</a> [CompensationStepObjectIDType](/src/target/wwsgen_types.go?s=91234:91363#L1474)
``` go
type CompensationStepObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationStepObjectType">type</a> [CompensationStepObjectType](/src/target/wwsgen_types.go?s=91365:91590#L1479)
``` go
type CompensationStepObjectType struct {
    ID         []CompensationStepObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompensationStepType">type</a> [CompensationStepType](/src/target/wwsgen_types.go?s=91624:92005#L1485)
``` go
type CompensationStepType struct {
    CompensationStepReference *CompensationStepObjectType `xml:"urn:com.workday/bsvc Compensation_Step_Reference,omitempty"`
    CompensationStepData      *CompensationStepDataType   `xml:"urn:com.workday/bsvc Compensation_Step_Data,omitempty"`
    Delete                    bool                        `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}
```
Contains the pay range data.










## <a name="CompensationTrancheReplacementDataType">type</a> [CompensationTrancheReplacementDataType](/src/target/wwsgen_types.go?s=92172:92340#L1492)
``` go
type CompensationTrancheReplacementDataType struct {
    Name       float64 `xml:"urn:com.workday/bsvc Name"`
    Percentage float64 `xml:"urn:com.workday/bsvc Percentage"`
}
```
This is a wrapper for the Compensation Tranche Replacement Data which is used to display and replace Compensation Tranche Data within Bonus Plan; it is optional.










## <a name="CompensationWeightedPercentMatrixObjectIDType">type</a> [CompensationWeightedPercentMatrixObjectIDType](/src/target/wwsgen_types.go?s=92404:92550#L1498)
``` go
type CompensationWeightedPercentMatrixObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompensationWeightedPercentMatrixObjectType">type</a> [CompensationWeightedPercentMatrixObjectType](/src/target/wwsgen_types.go?s=92552:92828#L1503)
``` go
type CompensationWeightedPercentMatrixObjectType struct {
    ID         []CompensationWeightedPercentMatrixObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CompplObjectIDType">type</a> [CompplObjectIDType](/src/target/wwsgen_types.go?s=92892:93011#L1509)
``` go
type CompplObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CompplObjectType">type</a> [CompplObjectType](/src/target/wwsgen_types.go?s=93013:93208#L1514)
``` go
type CompplObjectType struct {
    ID         []CompplObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="ConditionEntryOptionObjectIDType">type</a> [ConditionEntryOptionObjectIDType](/src/target/wwsgen_types.go?s=93272:93405#L1520)
``` go
type ConditionEntryOptionObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="ConditionEntryOptionObjectType">type</a> [ConditionEntryOptionObjectType](/src/target/wwsgen_types.go?s=93407:93644#L1525)
``` go
type ConditionEntryOptionObjectType struct {
    ID         []ConditionEntryOptionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="ConditionItemDataWWSType">type</a> [ConditionItemDataWWSType](/src/target/wwsgen_types.go?s=93706:95685#L1531)
``` go
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
```
Wrapper element containing data for each Condition Item.










### <a name="ConditionItemDataWWSType.MarshalXML">func</a> (\*ConditionItemDataWWSType) [MarshalXML](/src/target/wwsgen_types.go?s=95687:95778#L1550)
``` go
func (t *ConditionItemDataWWSType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="ConditionItemDataWWSType.UnmarshalXML">func</a> (\*ConditionItemDataWWSType) [UnmarshalXML](/src/target/wwsgen_types.go?s=96156:96249#L1562)
``` go
func (t *ConditionItemDataWWSType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="ConditionRuleCategoryObjectIDType">type</a> [ConditionRuleCategoryObjectIDType](/src/target/wwsgen_types.go?s=96699:96833#L1576)
``` go
type ConditionRuleCategoryObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="ConditionRuleCategoryObjectType">type</a> [ConditionRuleCategoryObjectType](/src/target/wwsgen_types.go?s=96835:97075#L1581)
``` go
type ConditionRuleCategoryObjectType struct {
    ID         []ConditionRuleCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="ConditionRuleDataWWSType">type</a> [ConditionRuleDataWWSType](/src/target/wwsgen_types.go?s=97124:97891#L1587)
``` go
type ConditionRuleDataWWSType struct {
    ConditionRuleID                string                            `xml:"urn:com.workday/bsvc Condition_Rule_ID,omitempty"`
    RuleDescription                string                            `xml:"urn:com.workday/bsvc Rule_Description"`
    Comment                        string                            `xml:"urn:com.workday/bsvc Comment,omitempty"`
    ConditionRuleCategoryReference []ConditionRuleCategoryObjectType `xml:"urn:com.workday/bsvc Condition_Rule_Category_Reference,omitempty"`
    ConditionItemData              []ConditionItemDataWWSType        `xml:"urn:com.workday/bsvc Condition_Item_Data"`
    CountriesReference             []CountryObjectType               `xml:"urn:com.workday/bsvc Countries_Reference,omitempty"`
}
```
Wrapper containing data for Condition Rule.










## <a name="ConditionRuleObjectIDType">type</a> [ConditionRuleObjectIDType](/src/target/wwsgen_types.go?s=97955:98081#L1597)
``` go
type ConditionRuleObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="ConditionRuleObjectType">type</a> [ConditionRuleObjectType](/src/target/wwsgen_types.go?s=98083:98299#L1602)
``` go
type ConditionRuleObjectType struct {
    ID         []ConditionRuleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CountryObjectIDType">type</a> [CountryObjectIDType](/src/target/wwsgen_types.go?s=98363:98483#L1608)
``` go
type CountryObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CountryObjectType">type</a> [CountryObjectType](/src/target/wwsgen_types.go?s=98485:98683#L1613)
``` go
type CountryObjectType struct {
    ID         []CountryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="CreateSeveranceWorksheetInput">type</a> [CreateSeveranceWorksheetInput](/src/target/wwsgen_client.go?s=23615:23777#L627)
``` go
type CreateSeveranceWorksheetInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Create_Severance_Worksheet_Request"`
    CreateSeveranceWorksheetRequestType
}
```









## <a name="CreateSeveranceWorksheetOutput">type</a> [CreateSeveranceWorksheetOutput](/src/target/wwsgen_client.go?s=23779:23944#L632)
``` go
type CreateSeveranceWorksheetOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Create_Severance_Worksheet_Response"`
    CreateSeveranceWorksheetResponseType
}
```









## <a name="CreateSeveranceWorksheetRequestType">type</a> [CreateSeveranceWorksheetRequestType](/src/target/wwsgen_types.go?s=98732:99169#L1619)
``` go
type CreateSeveranceWorksheetRequestType struct {
    BusinessProcessParameters      *BusinessProcessParametersType      `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
    EmployeeSeveranceWorksheetData *EmployeeSeveranceWorksheetDataType `xml:"urn:com.workday/bsvc Employee_Severance_Worksheet_Data"`
    Version                        string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Launch Employee Severance Worksheet Request










## <a name="CreateSeveranceWorksheetResponseType">type</a> [CreateSeveranceWorksheetResponseType](/src/target/wwsgen_types.go?s=99219:99577#L1626)
``` go
type CreateSeveranceWorksheetResponseType struct {
    EmployeeSeveranceWorksheetEventReference *EmployeeSeveranceWorksheetEventObjectType `xml:"urn:com.workday/bsvc Employee_Severance_Worksheet_Event_Reference,omitempty"`
    Version                                  string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Launch Employee Severance Worksheet Response










## <a name="CurrencyObjectIDType">type</a> [CurrencyObjectIDType](/src/target/wwsgen_types.go?s=99641:99762#L1632)
``` go
type CurrencyObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="CurrencyObjectType">type</a> [CurrencyObjectType](/src/target/wwsgen_types.go?s=99764:99965#L1637)
``` go
type CurrencyObjectType struct {
    ID         []CurrencyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="DefaultScorecardObjectIDType">type</a> [DefaultScorecardObjectIDType](/src/target/wwsgen_types.go?s=100029:100158#L1643)
``` go
type DefaultScorecardObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="DefaultScorecardObjectType">type</a> [DefaultScorecardObjectType](/src/target/wwsgen_types.go?s=100160:100385#L1648)
``` go
type DefaultScorecardObjectType struct {
    ID         []DefaultScorecardObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="DeferredBonusCalculationObjectIDType">type</a> [DeferredBonusCalculationObjectIDType](/src/target/wwsgen_types.go?s=100449:100586#L1654)
``` go
type DeferredBonusCalculationObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="DeferredBonusCalculationObjectType">type</a> [DeferredBonusCalculationObjectType](/src/target/wwsgen_types.go?s=100588:100837#L1659)
``` go
type DeferredBonusCalculationObjectType struct {
    ID         []DeferredBonusCalculationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="DeferredCompensationDataType">type</a> [DeferredCompensationDataType](/src/target/wwsgen_types.go?s=100864:101492#L1665)
``` go
type DeferredCompensationDataType struct {
    DeferredCompensationEligibilityReference []ConditionRuleObjectType              `xml:"urn:com.workday/bsvc Deferred_Compensation_Eligibility_Reference"`
    DefaultCalculationReference              *DeferredBonusCalculationObjectType    `xml:"urn:com.workday/bsvc Default_Calculation_Reference"`
    DefaultStockPlanReference                *StockPercentPlanObjectType            `xml:"urn:com.workday/bsvc Default_Stock_Plan_Reference"`
    DeferredCompensationProfiles             []DeferredCompensationProfilesDataType `xml:"urn:com.workday/bsvc Deferred_Compensation_Profiles,omitempty"`
}
```
Deferred Compensation










## <a name="DeferredCompensationProfilesDataType">type</a> [DeferredCompensationProfilesDataType](/src/target/wwsgen_types.go?s=101528:101923#L1673)
``` go
type DeferredCompensationProfilesDataType struct {
    EligibilityRuleReference []ConditionRuleObjectType           `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    CalculationReference     *DeferredBonusCalculationObjectType `xml:"urn:com.workday/bsvc Calculation_Reference"`
    StockPlanReference       *StockPercentPlanObjectType         `xml:"urn:com.workday/bsvc Stock_Plan_Reference"`
}
```
Deferred Compensation Profiles










## <a name="EligibilityWaitingPeriodObjectIDType">type</a> [EligibilityWaitingPeriodObjectIDType](/src/target/wwsgen_types.go?s=101987:102124#L1680)
``` go
type EligibilityWaitingPeriodObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="EligibilityWaitingPeriodObjectType">type</a> [EligibilityWaitingPeriodObjectType](/src/target/wwsgen_types.go?s=102126:102375#L1685)
``` go
type EligibilityWaitingPeriodObjectType struct {
    ID         []EligibilityWaitingPeriodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="EligibleEarningRequestCriteriaType">type</a> [EligibleEarningRequestCriteriaType](/src/target/wwsgen_types.go?s=102474:103072#L1691)
``` go
type EligibleEarningRequestCriteriaType struct {
    WorkerReference                 *WorkerObjectType                         `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
    PositionReference               *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
    EligibleEarningsPeriodReference *EligibleEarningsOverridePeriodObjectType `xml:"urn:com.workday/bsvc Eligible_Earnings_Period_Reference,omitempty"`
    BonusPlanReference              *BonusPercentPlanObjectType               `xml:"urn:com.workday/bsvc Bonus_Plan_Reference,omitempty"`
}
```
Utilize the Request Criteria element to search for specific instance(s) of Eligible Earnings.










## <a name="EligibleEarningsDataType">type</a> [EligibleEarningsDataType](/src/target/wwsgen_types.go?s=103121:104178#L1699)
``` go
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
```
Data element for the Put Eligible Earnings.










## <a name="EligibleEarningsOverrideHVDataType">type</a> [EligibleEarningsOverrideHVDataType](/src/target/wwsgen_types.go?s=104227:105074#L1711)
``` go
type EligibleEarningsOverrideHVDataType struct {
    EligibleEarningsID            string                       `xml:"urn:com.workday/bsvc Eligible_Earnings_ID,omitempty"`
    EmployeeReference             *WorkerObjectType            `xml:"urn:com.workday/bsvc Employee_Reference"`
    PositionReference             *PositionElementObjectType   `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
    ApplytoAllBonusPlans          bool                         `xml:"urn:com.workday/bsvc Apply_to_All_Bonus_Plans"`
    RestricttoBonusPlansReference []BonusPercentPlanObjectType `xml:"urn:com.workday/bsvc Restrict_to_Bonus_Plans_Reference"`
    Amount                        float64                      `xml:"urn:com.workday/bsvc Amount,omitempty"`
    CurrencyReference             *CurrencyObjectType          `xml:"urn:com.workday/bsvc Currency_Reference"`
}
```
Data element for the Put Eligible Earnings.










## <a name="EligibleEarningsOverrideObjectIDType">type</a> [EligibleEarningsOverrideObjectIDType](/src/target/wwsgen_types.go?s=105138:105275#L1722)
``` go
type EligibleEarningsOverrideObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="EligibleEarningsOverrideObjectType">type</a> [EligibleEarningsOverrideObjectType](/src/target/wwsgen_types.go?s=105277:105526#L1727)
``` go
type EligibleEarningsOverrideObjectType struct {
    ID         []EligibleEarningsOverrideObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="EligibleEarningsOverridePeriodObjectIDType">type</a> [EligibleEarningsOverridePeriodObjectIDType](/src/target/wwsgen_types.go?s=105590:105733#L1733)
``` go
type EligibleEarningsOverridePeriodObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="EligibleEarningsOverridePeriodObjectType">type</a> [EligibleEarningsOverridePeriodObjectType](/src/target/wwsgen_types.go?s=105735:106002#L1738)
``` go
type EligibleEarningsOverridePeriodObjectType struct {
    ID         []EligibleEarningsOverridePeriodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="EligibleEarningsRequestReferencesType">type</a> [EligibleEarningsRequestReferencesType](/src/target/wwsgen_types.go?s=106125:106299#L1744)
``` go
type EligibleEarningsRequestReferencesType struct {
    EligibleEarningsReference []EligibleEarningsOverrideObjectType `xml:"urn:com.workday/bsvc Eligible_Earnings_Reference"`
}
```
Utilize the Request References element to retrieve specific instance(s) of Eligible Earnings and its associated data.










## <a name="EligibleEarningsResponseDataType">type</a> [EligibleEarningsResponseDataType](/src/target/wwsgen_types.go?s=106401:106547#L1749)
``` go
type EligibleEarningsResponseDataType struct {
    EligibleEarnings []EligibleEarningsType `xml:"urn:com.workday/bsvc Eligible_Earnings,omitempty"`
}
```
Contains Eligible Earnings Override information based on Request References or Request Criteria.










## <a name="EligibleEarningsResponseGroupType">type</a> [EligibleEarningsResponseGroupType](/src/target/wwsgen_types.go?s=106796:107042#L1754)
``` go
type EligibleEarningsResponseGroupType struct {
    IncludeReference            *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
    IncludeEligibleEarningsData *bool `xml:"urn:com.workday/bsvc Include_Eligible_Earnings_Data,omitempty"`
}
```
The response group allows for the response data to be tailored to only included elements that the user is looking for.  If no response group is provided in the request the following elements will be returned:  Reference, Eligible Earnings Data










## <a name="EligibleEarningsType">type</a> [EligibleEarningsType](/src/target/wwsgen_types.go?s=107144:107435#L1760)
``` go
type EligibleEarningsType struct {
    EligibleEarningsReference *EligibleEarningsOverrideObjectType `xml:"urn:com.workday/bsvc Eligible_Earnings_Reference,omitempty"`
    EligibleEarningsData      *EligibleEarningsDataType           `xml:"urn:com.workday/bsvc Eligible_Earnings_Data,omitempty"`
}
```
Contains Eligible Earnings Override information based on Request References or Request Criteria.










## <a name="EmployeeMeritAdjustmentDataType">type</a> [EmployeeMeritAdjustmentDataType](/src/target/wwsgen_types.go?s=107513:108256#L1766)
``` go
type EmployeeMeritAdjustmentDataType struct {
    BasePayPlanReference *SalaryPayPlanObjectType `xml:"urn:com.workday/bsvc Base_Pay_Plan_Reference,omitempty"`
    PercentChange        float64                  `xml:"urn:com.workday/bsvc Percent_Change"`
    AmountChange         float64                  `xml:"urn:com.workday/bsvc Amount_Change"`
    NewAmount            float64                  `xml:"urn:com.workday/bsvc New_Amount"`
    CurrencyReference    *CurrencyObjectType      `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
    FrequencyReference   *FrequencyObjectType     `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
    CompensationComment  string                   `xml:"urn:com.workday/bsvc Compensation_Comment,omitempty"`
}
```
Data element for the Request Employee Merit Adjustment business process.










## <a name="EmployeeObjectIDType">type</a> [EmployeeObjectIDType](/src/target/wwsgen_types.go?s=108320:108441#L1777)
``` go
type EmployeeObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="EmployeeObjectType">type</a> [EmployeeObjectType](/src/target/wwsgen_types.go?s=108443:108644#L1782)
``` go
type EmployeeObjectType struct {
    ID         []EmployeeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="EmployeeSeverancePackagePayComponentDataType">type</a> [EmployeeSeverancePackagePayComponentDataType](/src/target/wwsgen_types.go?s=108737:109669#L1788)
``` go
type EmployeeSeverancePackagePayComponentDataType struct {
    EmployeeSeverancePackagePayComponentReference *SeverancePackageComponentTypeObjectType `xml:"urn:com.workday/bsvc Employee_Severance_Package_Pay_Component_Reference,omitempty"`
    AssignmentDuration                            float64                                  `xml:"urn:com.workday/bsvc Assignment_Duration,omitempty"`
    PayRate                                       float64                                  `xml:"urn:com.workday/bsvc Pay_Rate,omitempty"`
    Total                                         float64                                  `xml:"urn:com.workday/bsvc Total,omitempty"`
    PlanReference                                 *SalaryPlanObjectType                    `xml:"urn:com.workday/bsvc Plan_Reference,omitempty"`
    Comments                                      string                                   `xml:"urn:com.workday/bsvc Comments,omitempty"`
}
```
Subelement for entering a severance component type along with eligibility and a comment










## <a name="EmployeeSeveranceWorksheetDataType">type</a> [EmployeeSeveranceWorksheetDataType](/src/target/wwsgen_types.go?s=109708:112282#L1798)
``` go
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
```
Employee Severance Worksheet Data










### <a name="EmployeeSeveranceWorksheetDataType.MarshalXML">func</a> (\*EmployeeSeveranceWorksheetDataType) [MarshalXML](/src/target/wwsgen_types.go?s=112284:112385#L1820)
``` go
func (t *EmployeeSeveranceWorksheetDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="EmployeeSeveranceWorksheetDataType.UnmarshalXML">func</a> (\*EmployeeSeveranceWorksheetDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=114156:114259#L1848)
``` go
func (t *EmployeeSeveranceWorksheetDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="EmployeeSeveranceWorksheetEventDataType">type</a> [EmployeeSeveranceWorksheetEventDataType](/src/target/wwsgen_types.go?s=116099:118894#L1878)
``` go
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
```
Employee Severance Worksheet Event Data










### <a name="EmployeeSeveranceWorksheetEventDataType.MarshalXML">func</a> (\*EmployeeSeveranceWorksheetEventDataType) [MarshalXML](/src/target/wwsgen_types.go?s=118896:119002#L1899)
``` go
func (t *EmployeeSeveranceWorksheetEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="EmployeeSeveranceWorksheetEventDataType.UnmarshalXML">func</a> (\*EmployeeSeveranceWorksheetEventDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=120767:120875#L1927)
``` go
func (t *EmployeeSeveranceWorksheetEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="EmployeeSeveranceWorksheetEventObjectIDType">type</a> [EmployeeSeveranceWorksheetEventObjectIDType](/src/target/wwsgen_types.go?s=122728:122872#L1957)
``` go
type EmployeeSeveranceWorksheetEventObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="EmployeeSeveranceWorksheetEventObjectType">type</a> [EmployeeSeveranceWorksheetEventObjectType](/src/target/wwsgen_types.go?s=122874:123144#L1962)
``` go
type EmployeeSeveranceWorksheetEventObjectType struct {
    ID         []EmployeeSeveranceWorksheetEventObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="EmployeeSeveranceWorksheetEventRequestCriteriaType">type</a> [EmployeeSeveranceWorksheetEventRequestCriteriaType](/src/target/wwsgen_types.go?s=123219:123285#L1968)
``` go
type EmployeeSeveranceWorksheetEventRequestCriteriaType struct {
}
```
Request Criteria to filter Employee Severance Worksheet Event results










## <a name="EmployeeSeveranceWorksheetEventRequestReferencesType">type</a> [EmployeeSeveranceWorksheetEventRequestReferencesType](/src/target/wwsgen_types.go?s=123339:123567#L1972)
``` go
type EmployeeSeveranceWorksheetEventRequestReferencesType struct {
    EmployeeSeveranceWorksheetEventReference []EmployeeSeveranceWorksheetEventObjectType `xml:"urn:com.workday/bsvc Employee_Severance_Worksheet_Event_Reference"`
}
```
Reference for Employee Severance Worksheet Event










## <a name="EmployeeSeveranceWorksheetEventResponseDataType">type</a> [EmployeeSeveranceWorksheetEventResponseDataType](/src/target/wwsgen_types.go?s=123638:123835#L1977)
``` go
type EmployeeSeveranceWorksheetEventResponseDataType struct {
    EmployeeSeveranceWorksheet []EmployeeSeveranceWorksheetEventType `xml:"urn:com.workday/bsvc Employee_Severance_Worksheet,omitempty"`
}
```
Response element for Retrieve Employee Severance Worksheet Events










## <a name="EmployeeSeveranceWorksheetEventResponseGroupType">type</a> [EmployeeSeveranceWorksheetEventResponseGroupType](/src/target/wwsgen_types.go?s=123894:124039#L1982)
``` go
type EmployeeSeveranceWorksheetEventResponseGroupType struct {
    IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}
```
Response Group for Employee Severance Worksheet Event










## <a name="EmployeeSeveranceWorksheetEventType">type</a> [EmployeeSeveranceWorksheetEventType](/src/target/wwsgen_types.go?s=124145:124523#L1987)
``` go
type EmployeeSeveranceWorksheetEventType struct {
    EmployeeSeveranceWorksheetEventReference *EmployeeSeveranceWorksheetEventObjectType `xml:"urn:com.workday/bsvc Employee_Severance_Worksheet_Event_Reference,omitempty"`
    EmployeeSeveranceWorksheetData           []EmployeeSeveranceWorksheetEventDataType  `xml:"urn:com.workday/bsvc Employee_Severance_Worksheet_Data,omitempty"`
}
```
Contains reference to Employee Severance Worksheet Event and Employee Severance Worksheet Event Data










## <a name="EnhancedSeveranceMatrixDataType">type</a> [EnhancedSeveranceMatrixDataType](/src/target/wwsgen_types.go?s=124604:126153#L1993)
``` go
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
```
Element to specify the data for the type of severance matrix to put/update.










### <a name="EnhancedSeveranceMatrixDataType.MarshalXML">func</a> (\*EnhancedSeveranceMatrixDataType) [MarshalXML](/src/target/wwsgen_types.go?s=126155:126253#L2005)
``` go
func (t *EnhancedSeveranceMatrixDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="EnhancedSeveranceMatrixDataType.UnmarshalXML">func</a> (\*EnhancedSeveranceMatrixDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=126515:126615#L2015)
``` go
func (t *EnhancedSeveranceMatrixDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="EnhancedSeveranceMatrixRequestCriteriaType">type</a> [EnhancedSeveranceMatrixRequestCriteriaType](/src/target/wwsgen_types.go?s=126905:127362#L2027)
``` go
type EnhancedSeveranceMatrixRequestCriteriaType struct {
    Name                                 string                             `xml:"urn:com.workday/bsvc Name,omitempty"`
    EnhancedSeveranceMatrixFilterOptions []SeveranceMatrixFilterOptionsType `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix_Filter_Options,omitempty"`
    IncludeInactive                      *bool                              `xml:"urn:com.workday/bsvc Include_Inactive,omitempty"`
}
```
Request Criteria










## <a name="EnhancedSeveranceMatrixRequestReferencesType">type</a> [EnhancedSeveranceMatrixRequestReferencesType](/src/target/wwsgen_types.go?s=127395:127590#L2034)
``` go
type EnhancedSeveranceMatrixRequestReferencesType struct {
    EnhancedSeveranceMatrixReference []SeveranceMatrixabstractObjectType `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix_Reference"`
}
```
Request references element.










## <a name="EnhancedSeveranceMatrixResponseDataType">type</a> [EnhancedSeveranceMatrixResponseDataType](/src/target/wwsgen_types.go?s=127659:127834#L2039)
``` go
type EnhancedSeveranceMatrixResponseDataType struct {
    EnhancedSeveranceMatrix []EnhancedSeveranceMatrixType `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix,omitempty"`
}
```
The response for the Get Enhanced Severance Matrix web service.










## <a name="EnhancedSeveranceMatrixResponseGroupType">type</a> [EnhancedSeveranceMatrixResponseGroupType](/src/target/wwsgen_types.go?s=127859:127996#L2044)
``` go
type EnhancedSeveranceMatrixResponseGroupType struct {
    IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}
```
The response group.










## <a name="EnhancedSeveranceMatrixType">type</a> [EnhancedSeveranceMatrixType](/src/target/wwsgen_types.go?s=128040:128366#L2049)
``` go
type EnhancedSeveranceMatrixType struct {
    EnhancedSeveranceMatrixReference *SeveranceMatrixabstractObjectType `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix_Reference,omitempty"`
    EnhancedSeveranceMatrixData      []EnhancedSeveranceMatrixDataType  `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix_Data,omitempty"`
}
```
The Enhanced Severance Matrix element.










## <a name="EventAttachmentCategoryObjectIDType">type</a> [EventAttachmentCategoryObjectIDType](/src/target/wwsgen_types.go?s=128430:128566#L2055)
``` go
type EventAttachmentCategoryObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="EventAttachmentCategoryObjectType">type</a> [EventAttachmentCategoryObjectType](/src/target/wwsgen_types.go?s=128568:128814#L2060)
``` go
type EventAttachmentCategoryObjectType struct {
    ID         []EventAttachmentCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="EventClassificationSubcategoryObjectIDType">type</a> [EventClassificationSubcategoryObjectIDType](/src/target/wwsgen_types.go?s=128878:129021#L2066)
``` go
type EventClassificationSubcategoryObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="EventClassificationSubcategoryObjectType">type</a> [EventClassificationSubcategoryObjectType](/src/target/wwsgen_types.go?s=129023:129290#L2071)
``` go
type EventClassificationSubcategoryObjectType struct {
    ID         []EventClassificationSubcategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="ExpenseAccumulatorObjectIDType">type</a> [ExpenseAccumulatorObjectIDType](/src/target/wwsgen_types.go?s=129354:129485#L2077)
``` go
type ExpenseAccumulatorObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="ExpenseAccumulatorObjectType">type</a> [ExpenseAccumulatorObjectType](/src/target/wwsgen_types.go?s=129487:129718#L2082)
``` go
type ExpenseAccumulatorObjectType struct {
    ID         []ExpenseAccumulatorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="ExpenseItemObjectIDType">type</a> [ExpenseItemObjectIDType](/src/target/wwsgen_types.go?s=129782:129906#L2088)
``` go
type ExpenseItemObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="ExpenseItemObjectType">type</a> [ExpenseItemObjectType](/src/target/wwsgen_types.go?s=129908:130118#L2093)
``` go
type ExpenseItemObjectType struct {
    ID         []ExpenseItemObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="ExternalFieldObjectIDType">type</a> [ExternalFieldObjectIDType](/src/target/wwsgen_types.go?s=130182:130308#L2099)
``` go
type ExternalFieldObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="ExternalFieldObjectType">type</a> [ExternalFieldObjectType](/src/target/wwsgen_types.go?s=130310:130526#L2104)
``` go
type ExternalFieldObjectType struct {
    ID         []ExternalFieldObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="FilterDateTimeZoneDataType">type</a> [FilterDateTimeZoneDataType](/src/target/wwsgen_types.go?s=130562:130814#L2110)
``` go
type FilterDateTimeZoneDataType struct {
    FilterDateTime          *time.Time          `xml:"urn:com.workday/bsvc Filter_DateTime,omitempty"`
    FilterTimeZoneReference *TimeZoneObjectType `xml:"urn:com.workday/bsvc Filter_TimeZone_Reference,omitempty"`
}
```
DateTimeZone value for filter.










### <a name="FilterDateTimeZoneDataType.MarshalXML">func</a> (\*FilterDateTimeZoneDataType) [MarshalXML](/src/target/wwsgen_types.go?s=130816:130909#L2115)
``` go
func (t *FilterDateTimeZoneDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="FilterDateTimeZoneDataType.UnmarshalXML">func</a> (\*FilterDateTimeZoneDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=131187:131282#L2125)
``` go
func (t *FilterDateTimeZoneDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="FinancialsBusinessSubProcessParametersType">type</a> [FinancialsBusinessSubProcessParametersType](/src/target/wwsgen_types.go?s=131991:132282#L2137)
``` go
type FinancialsBusinessSubProcessParametersType struct {
    Skip                       *bool                           `xml:"urn:com.workday/bsvc Skip,omitempty"`
    BusinessProcessCommentData *BusinessProcessCommentDataType `xml:"urn:com.workday/bsvc Business_Process_Comment_Data,omitempty"`
}
```
Container for the processing options for sub-business processes within a business process. If no options are submitted (or the options are submitted as 'false') then the sub-business process is simply initiated as if it where submitted on-line with approvals, reviews, notifications and to-do's in place. If the Initiator is an Integration System User, any validations you configured on the Initiation step are ignored.










## <a name="FrequencyBehaviorObjectIDType">type</a> [FrequencyBehaviorObjectIDType](/src/target/wwsgen_types.go?s=132346:132476#L2143)
``` go
type FrequencyBehaviorObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="FrequencyBehaviorObjectType">type</a> [FrequencyBehaviorObjectType](/src/target/wwsgen_types.go?s=132478:132706#L2148)
``` go
type FrequencyBehaviorObjectType struct {
    ID         []FrequencyBehaviorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="FrequencyObjectIDType">type</a> [FrequencyObjectIDType](/src/target/wwsgen_types.go?s=132770:132892#L2154)
``` go
type FrequencyObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="FrequencyObjectType">type</a> [FrequencyObjectType](/src/target/wwsgen_types.go?s=132894:133098#L2159)
``` go
type FrequencyObjectType struct {
    ID         []FrequencyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="FuturePaymentPlanAssignmentDataType">type</a> [FuturePaymentPlanAssignmentDataType](/src/target/wwsgen_types.go?s=133200:133895#L2165)
``` go
type FuturePaymentPlanAssignmentDataType struct {
    CompensationPlanReference         *FuturePaymentPlanObjectType `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
    IndividualTargetAmount            float64                      `xml:"urn:com.workday/bsvc Individual_Target_Amount,omitempty"`
    IndividualTargetCurrencyReference *CurrencyObjectType          `xml:"urn:com.workday/bsvc Individual_Target_Currency_Reference,omitempty"`
    IndividualTargetPaymentDate       *time.Time                   `xml:"urn:com.workday/bsvc Individual_Target_Payment_Date,omitempty"`
    Comment                           string                       `xml:"urn:com.workday/bsvc Comment,omitempty"`
}
```
Future payment plan assignment data submitted by the put future payment plan assignment request.










### <a name="FuturePaymentPlanAssignmentDataType.MarshalXML">func</a> (\*FuturePaymentPlanAssignmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=133897:133999#L2173)
``` go
func (t *FuturePaymentPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="FuturePaymentPlanAssignmentDataType.UnmarshalXML">func</a> (\*FuturePaymentPlanAssignmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=134332:134436#L2183)
``` go
func (t *FuturePaymentPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="FuturePaymentPlanAssignmentRequestCriteriaType">type</a> [FuturePaymentPlanAssignmentRequestCriteriaType](/src/target/wwsgen_types.go?s=134925:135197#L2195)
``` go
type FuturePaymentPlanAssignmentRequestCriteriaType struct {
    EmployeeReference []WorkerObjectType          `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
    PositionReference []PositionElementObjectType `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
}
```
Future Payment Plan Assignment Request Criteria contains a set of zero or more Employees for whom Future Payment Plan Assignments are requested.










## <a name="FuturePaymentPlanAssignmentResponseDataType">type</a> [FuturePaymentPlanAssignmentResponseDataType](/src/target/wwsgen_types.go?s=135281:135499#L2201)
``` go
type FuturePaymentPlanAssignmentResponseDataType struct {
    PositionFuturePaymentPlanAssignment []PositionFuturePaymentPlanAssignmentsType `xml:"urn:com.workday/bsvc Position_Future_Payment_Plan_Assignment,omitempty"`
}
```
Contains future payment plan assignment information based on Request Criteria.










## <a name="FuturePaymentPlanDataType">type</a> [FuturePaymentPlanDataType](/src/target/wwsgen_types.go?s=135529:137119#L2206)
``` go
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
```
Future Payment Plan Data










## <a name="FuturePaymentPlanObjectIDType">type</a> [FuturePaymentPlanObjectIDType](/src/target/wwsgen_types.go?s=137183:137313#L2220)
``` go
type FuturePaymentPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="FuturePaymentPlanObjectType">type</a> [FuturePaymentPlanObjectType](/src/target/wwsgen_types.go?s=137315:137543#L2225)
``` go
type FuturePaymentPlanObjectType struct {
    ID         []FuturePaymentPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="GetBenchmarkJobsInput">type</a> [GetBenchmarkJobsInput](/src/target/wwsgen_client.go?s=16810:16948#L451)
``` go
type GetBenchmarkJobsInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Benchmark_Jobs_Request"`
    GetBenchmarkJobsRequestType
}
```









## <a name="GetBenchmarkJobsOutput">type</a> [GetBenchmarkJobsOutput](/src/target/wwsgen_client.go?s=16950:17091#L456)
``` go
type GetBenchmarkJobsOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Benchmark_Jobs_Response"`
    GetBenchmarkJobsResponseType
}
```









## <a name="GetBenchmarkJobsRequestType">type</a> [GetBenchmarkJobsRequestType](/src/target/wwsgen_types.go?s=137585:138173#L2231)
``` go
type GetBenchmarkJobsRequestType struct {
    RequestReferences *BenchmarkJobRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *BenchmarkJobRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *BenchmarkJobResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request element to get Benchmark Job










## <a name="GetBenchmarkJobsResponseType">type</a> [GetBenchmarkJobsResponseType](/src/target/wwsgen_types.go?s=138206:139012#L2240)
``` go
type GetBenchmarkJobsResponseType struct {
    RequestReferences *BenchmarkJobRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *BenchmarkJobRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *BenchmarkJobResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    ResponseResults   *ResponseResultsType               `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      *BenchmarkJobResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Get Benchmark Jobs Response










## <a name="GetCompensationEligibilityRulesInput">type</a> [GetCompensationEligibilityRulesInput](/src/target/wwsgen_client.go?s=19866:20050#L539)
``` go
type GetCompensationEligibilityRulesInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Eligibility_Rules_Request"`
    GetCompensationEligibilityRulesRequestType
}
```









## <a name="GetCompensationEligibilityRulesOutput">type</a> [GetCompensationEligibilityRulesOutput](/src/target/wwsgen_client.go?s=20052:20239#L544)
``` go
type GetCompensationEligibilityRulesOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Eligibility_Rules_Response"`
    GetCompensationEligibilityRulesResponseType
}
```









## <a name="GetCompensationEligibilityRulesRequestType">type</a> [GetCompensationEligibilityRulesRequestType](/src/target/wwsgen_types.go?s=139077:139755#L2251)
``` go
type GetCompensationEligibilityRulesRequestType struct {
    RequestReferences *CompensationEligibilityRuleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *CompensationEligibilityRuleRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *CompensationEligibilityRuleResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Get Compensation Eligibility Rules Request incoming element










## <a name="GetCompensationEligibilityRulesResponseType">type</a> [GetCompensationEligibilityRulesResponseType](/src/target/wwsgen_types.go?s=139812:140615#L2260)
``` go
type GetCompensationEligibilityRulesResponseType struct {
    RequestReferences *CompensationEligibilityRuleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *CompensationEligibilityRuleRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseResults   *ResponseResultsType                              `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      *CompensationEligibilityRuleResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Get Compensation Eligibility Rules Response element










## <a name="GetCompensationEligibilityRuleswithoutDependenciesInput">type</a> [GetCompensationEligibilityRuleswithoutDependenciesInput](/src/target/wwsgen_client.go?s=20898:21141#L561)
``` go
type GetCompensationEligibilityRuleswithoutDependenciesInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Eligibility_Rules_without_Dependencies_Request"`
    GetCompensationEligibilityRuleswithoutDependenciesRequestType
}
```









## <a name="GetCompensationEligibilityRuleswithoutDependenciesOutput">type</a> [GetCompensationEligibilityRuleswithoutDependenciesOutput](/src/target/wwsgen_client.go?s=21143:21349#L566)
``` go
type GetCompensationEligibilityRuleswithoutDependenciesOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Eligibility_Rules_Response"`
    GetCompensationEligibilityRulesResponseType
}
```









## <a name="GetCompensationEligibilityRuleswithoutDependenciesRequestType">type</a> [GetCompensationEligibilityRuleswithoutDependenciesRequestType](/src/target/wwsgen_types.go?s=140692:141266#L2270)
``` go
type GetCompensationEligibilityRuleswithoutDependenciesRequestType struct {
    RequestReferences *CompensationEligibilityRuleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *CompensationEligibilityRuleRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Get Compensation Eligibility Rules without Dependencies Request element










## <a name="GetCompensationGradeHierarchyInput">type</a> [GetCompensationGradeHierarchyInput](/src/target/wwsgen_client.go?s=24542:24720#L649)
``` go
type GetCompensationGradeHierarchyInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Grade_Hierarchy_Request"`
    GetCompensationGradeHierarchyRequestType
}
```









## <a name="GetCompensationGradeHierarchyOutput">type</a> [GetCompensationGradeHierarchyOutput](/src/target/wwsgen_client.go?s=24722:24903#L654)
``` go
type GetCompensationGradeHierarchyOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Grade_Hierarchy_Response"`
    GetCompensationGradeHierarchyResponseType
}
```









## <a name="GetCompensationGradeHierarchyRequestType">type</a> [GetCompensationGradeHierarchyRequestType](/src/target/wwsgen_types.go?s=141518:141943#L2278)
``` go
type GetCompensationGradeHierarchyRequestType struct {
    RequestReferences *CompensationGradeHierarchyRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    ResponseFilter    *ResponseFilterType                              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    Version           string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request element to get compensation grade hierarchy information. If no criteria are specified in either the compensation grade hierarchy request references or compensation grade hierarchy request criteria elements, all instances will be returned.










## <a name="GetCompensationGradeHierarchyResponseType">type</a> [GetCompensationGradeHierarchyResponseType](/src/target/wwsgen_types.go?s=141945:142616#L2284)
``` go
type GetCompensationGradeHierarchyResponseType struct {
    RequestReferences *CompensationGradeHierarchyRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    ResponseFilter    *ResponseFilterType                              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseResults   *ResponseResultsType                             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      []CompensationGradeHierarchyResponseDataType     `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```









## <a name="GetCompensationGradesInput">type</a> [GetCompensationGradesInput](/src/target/wwsgen_client.go?s=9545:9698#L253)
``` go
type GetCompensationGradesInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Grades_Request"`
    GetCompensationGradesRequestType
}
```









## <a name="GetCompensationGradesOutput">type</a> [GetCompensationGradesOutput](/src/target/wwsgen_client.go?s=9700:9856#L258)
``` go
type GetCompensationGradesOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Grades_Response"`
    GetCompensationGradesResponseType
}
```









## <a name="GetCompensationGradesRequestType">type</a> [GetCompensationGradesRequestType](/src/target/wwsgen_types.go?s=142878:143496#L2293)
``` go
type GetCompensationGradesRequestType struct {
    RequestReferences *CompensationGradeRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *CompensationGradeRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *CompensationGradeResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request element to get compensation grade and grade profile information.  If no criteria are specified in either the compensation grade request references or compensaiton grade request criteria elements, all instances of eligible earnings will be returned.










## <a name="GetCompensationGradesResponseType">type</a> [GetCompensationGradesResponseType](/src/target/wwsgen_types.go?s=143589:144326#L2302)
``` go
type GetCompensationGradesResponseType struct {
    RequestReferences []CompensationGradeRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    ResponseFilter    []ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     []CompensationGradeResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    ResponseResults   []ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      []CompensationGradeResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element containing instances of compensation grades and their associated data.










## <a name="GetCompensationMatricesInput">type</a> [GetCompensationMatricesInput](/src/target/wwsgen_client.go?s=7920:8079#L209)
``` go
type GetCompensationMatricesInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Matrices_Request"`
    GetCompensationMatricesRequestType
}
```









## <a name="GetCompensationMatricesOutput">type</a> [GetCompensationMatricesOutput](/src/target/wwsgen_client.go?s=8081:8243#L214)
``` go
type GetCompensationMatricesOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Matrices_Response"`
    GetCompensationMatricesResponseType
}
```









## <a name="GetCompensationMatricesRequestType">type</a> [GetCompensationMatricesRequestType](/src/target/wwsgen_types.go?s=144373:144998#L2312)
``` go
type GetCompensationMatricesRequestType struct {
    RequestReferences *CompensationMatrixRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *CompensationMatrixRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *CompensationMatrixResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Get Compensation Matrices Request Element










## <a name="GetCompensationMatricesResponseType">type</a> [GetCompensationMatricesResponseType](/src/target/wwsgen_types.go?s=145047:145909#L2321)
``` go
type GetCompensationMatricesResponseType struct {
    RequestReferences []CompensationMatrixRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   []CompensationMatrixRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    []ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     []CompensationMatrixResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    ResponseResults   []ResponseResultsType                     `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      []CompensationMatrixResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Get Compensation Matrices Response element.










## <a name="GetCompensationPlansInput">type</a> [GetCompensationPlansInput](/src/target/wwsgen_client.go?s=11160:11310#L297)
``` go
type GetCompensationPlansInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Plans_Request"`
    GetCompensationPlansRequestType
}
```









## <a name="GetCompensationPlansOutput">type</a> [GetCompensationPlansOutput](/src/target/wwsgen_client.go?s=11312:11465#L302)
``` go
type GetCompensationPlansOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Plans_Response"`
    GetCompensationPlansResponseType
}
```









## <a name="GetCompensationPlansRequestType">type</a> [GetCompensationPlansRequestType](/src/target/wwsgen_types.go?s=146130:146742#L2332)
``` go
type GetCompensationPlansRequestType struct {
    RequestReferences *CompensationPlanRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *CompensationPlanRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *CompensationPlanResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request element to get compensation plan information.  If no criteria is specified in the compensation plan request references or compensaiton plan request criteria elements, all compensation plans will be returned.










## <a name="GetCompensationPlansResponseType">type</a> [GetCompensationPlansResponseType](/src/target/wwsgen_types.go?s=146832:147670#L2341)
``` go
type GetCompensationPlansResponseType struct {
    RequestReferences *CompensationPlanRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *CompensationPlanRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *CompensationPlanResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    ResponseResults   *ResponseResultsType                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      *CompensationPlanResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element containing instances of compensation plans and the associated data.










## <a name="GetCompensationScorecardResultsInput">type</a> [GetCompensationScorecardResultsInput](/src/target/wwsgen_client.go?s=22739:22923#L605)
``` go
type GetCompensationScorecardResultsInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Scorecard_Results_Request"`
    GetCompensationScorecardResultsRequestType
}
```









## <a name="GetCompensationScorecardResultsOutput">type</a> [GetCompensationScorecardResultsOutput](/src/target/wwsgen_client.go?s=22925:23112#L610)
``` go
type GetCompensationScorecardResultsOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Scorecard_Results_Response"`
    GetCompensationScorecardResultsResponseType
}
```









## <a name="GetCompensationScorecardResultsRequestType">type</a> [GetCompensationScorecardResultsRequestType](/src/target/wwsgen_types.go?s=147747:148425#L2352)
``` go
type GetCompensationScorecardResultsRequestType struct {
    RequestReferences *CompensationScorecardResultRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *CompensationScorecardResultRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *CompensationScorecardResultResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Contains the request details to retrieve one or more scorecard results.










## <a name="GetCompensationScorecardResultsResponseType">type</a> [GetCompensationScorecardResultsResponseType](/src/target/wwsgen_types.go?s=148503:149429#L2361)
``` go
type GetCompensationScorecardResultsResponseType struct {
    RequestReferences *CompensationScorecardResultRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *CompensationScorecardResultRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *CompensationScorecardResultResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    ResponseResults   *ResponseResultsType                              `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      *CompensationScorecardResultResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Contains the response details to retrieve one or more scorecard results.










## <a name="GetCompensationScorecardsInput">type</a> [GetCompensationScorecardsInput](/src/target/wwsgen_client.go?s=14604:14769#L385)
``` go
type GetCompensationScorecardsInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Scorecards_Request"`
    GetCompensationScorecardsRequestType
}
```









## <a name="GetCompensationScorecardsOutput">type</a> [GetCompensationScorecardsOutput](/src/target/wwsgen_client.go?s=14771:14939#L390)
``` go
type GetCompensationScorecardsOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Compensation_Scorecards_Response"`
    GetCompensationScorecardsResponseType
}
```









## <a name="GetCompensationScorecardsRequestType">type</a> [GetCompensationScorecardsRequestType](/src/target/wwsgen_types.go?s=149499:150141#L2372)
``` go
type GetCompensationScorecardsRequestType struct {
    RequestReferences *CompensationScorecardRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *CompensationScorecardRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *CompensationScorecardResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Contains the request details to retrieve one or more scorecards.










## <a name="GetCompensationScorecardsResponseType">type</a> [GetCompensationScorecardsResponseType](/src/target/wwsgen_types.go?s=150211:151089#L2381)
``` go
type GetCompensationScorecardsResponseType struct {
    RequestReferences *CompensationScorecardRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *CompensationScorecardRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *CompensationScorecardResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    ResponseResults   *ResponseResultsType                        `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      *CompensationScorecardResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Contains the response details to retrieve one or more scorecard.










## <a name="GetEligibleEarningsInput">type</a> [GetEligibleEarningsInput](/src/target/wwsgen_client.go?s=7140:7287#L187)
``` go
type GetEligibleEarningsInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Eligible_Earnings_Request"`
    GetEligibleEarningsRequestType
}
```









## <a name="GetEligibleEarningsOutput">type</a> [GetEligibleEarningsOutput](/src/target/wwsgen_client.go?s=7289:7439#L192)
``` go
type GetEligibleEarningsOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Eligible_Earnings_Response"`
    GetEligibleEarningsResponseType
}
```









## <a name="GetEligibleEarningsRequestType">type</a> [GetEligibleEarningsRequestType](/src/target/wwsgen_types.go?s=151427:152038#L2392)
``` go
type GetEligibleEarningsRequestType struct {
    RequestReferences *EligibleEarningsRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *EligibleEarningRequestCriteriaType    `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *EligibleEarningsResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request element to get eligible earnings override information for an employee, eligible earnings override period or percent based bonus plan.  If no criteria are specified in either the Eligible Earnings Request References or Eligible Earnings Request Criteria elements, all instances of eligible earnings override will be returned.










## <a name="GetEligibleEarningsResponseType">type</a> [GetEligibleEarningsResponseType](/src/target/wwsgen_types.go?s=152129:152973#L2401)
``` go
type GetEligibleEarningsResponseType struct {
    RequestReferences []EligibleEarningsRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   []EligibleEarningRequestCriteriaType    `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    []ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     []EligibleEarningsResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    ResponseResults   []ResponseResultsType                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      []EligibleEarningsResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element containing instances of Eligible Earnings and their associated data.










## <a name="GetEnhancedSeveranceMatricesInput">type</a> [GetEnhancedSeveranceMatricesInput](/src/target/wwsgen_client.go?s=38131:38306#L979)
``` go
type GetEnhancedSeveranceMatricesInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Enhanced_Severance_Matrices_Request"`
    GetEnhancedSeveranceMatricesRequestType
}
```









## <a name="GetEnhancedSeveranceMatricesOutput">type</a> [GetEnhancedSeveranceMatricesOutput](/src/target/wwsgen_client.go?s=38308:38486#L984)
``` go
type GetEnhancedSeveranceMatricesOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Enhanced_Severance_Matrices_Response"`
    GetEnhancedSeveranceMatricesResponseType
}
```









## <a name="GetEnhancedSeveranceMatricesRequestType">type</a> [GetEnhancedSeveranceMatricesRequestType](/src/target/wwsgen_types.go?s=153025:153680#L2412)
``` go
type GetEnhancedSeveranceMatricesRequestType struct {
    RequestReferences *EnhancedSeveranceMatrixRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *EnhancedSeveranceMatrixRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                           `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *EnhancedSeveranceMatrixResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Get Enhanced Severance Matrix request element.










## <a name="GetEnhancedSeveranceMatricesResponseType">type</a> [GetEnhancedSeveranceMatricesResponseType](/src/target/wwsgen_types.go?s=153749:154644#L2421)
``` go
type GetEnhancedSeveranceMatricesResponseType struct {
    RequestReferences *EnhancedSeveranceMatrixRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *EnhancedSeveranceMatrixRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                           `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *EnhancedSeveranceMatrixResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    ResponseResults   *ResponseResultsType                          `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      *EnhancedSeveranceMatrixResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
The response of the Get Enhanced Severance Matrices web service










## <a name="GetFuturePaymentPlanAssignmentsInput">type</a> [GetFuturePaymentPlanAssignmentsInput](/src/target/wwsgen_client.go?s=12824:13009#L341)
``` go
type GetFuturePaymentPlanAssignmentsInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Future_Payment_Plan_Assignments_Request"`
    GetFuturePaymentPlanAssignmentsRequestType
}
```









## <a name="GetFuturePaymentPlanAssignmentsOutput">type</a> [GetFuturePaymentPlanAssignmentsOutput](/src/target/wwsgen_client.go?s=13011:13199#L346)
``` go
type GetFuturePaymentPlanAssignmentsOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Future_Payment_Plan_Assignments_Response"`
    GetFuturePaymentPlanAssignmentsResponseType
}
```









## <a name="GetFuturePaymentPlanAssignmentsRequestType">type</a> [GetFuturePaymentPlanAssignmentsRequestType](/src/target/wwsgen_types.go?s=154811:155227#L2432)
``` go
type GetFuturePaymentPlanAssignmentsRequestType struct {
    RequestCriteria *FuturePaymentPlanAssignmentRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter  *ResponseFilterType                             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    Version         string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request element to get future payment plan assignments. If no criteria are specified then all assignments active for today will be returned, grouped by employee.










## <a name="GetFuturePaymentPlanAssignmentsResponseType">type</a> [GetFuturePaymentPlanAssignmentsResponseType](/src/target/wwsgen_types.go?s=155391:156047#L2439)
``` go
type GetFuturePaymentPlanAssignmentsResponseType struct {
    RequestCriteria *FuturePaymentPlanAssignmentRequestCriteriaType `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter  *ResponseFilterType                             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseResults *ResponseResultsType                            `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData    *FuturePaymentPlanAssignmentResponseDataType    `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version         string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Get Future Payment Plan Assignments Response element containing retrieved Future Payment Plan Assignments and associated web service request related criteria.










## <a name="GetOneTimePaymentPlanConfigurableCategoriesInput">type</a> [GetOneTimePaymentPlanConfigurableCategoriesInput](/src/target/wwsgen_client.go?s=33665:33888#L869)
``` go
type GetOneTimePaymentPlanConfigurableCategoriesInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_One-Time_Payment_Plan_Configurable_Categories_Request"`
    GetOneTimePaymentPlanConfigurableCategoriesRequestType
}
```









## <a name="GetOneTimePaymentPlanConfigurableCategoriesOutput">type</a> [GetOneTimePaymentPlanConfigurableCategoriesOutput](/src/target/wwsgen_client.go?s=33890:34116#L874)
``` go
type GetOneTimePaymentPlanConfigurableCategoriesOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_One-Time_Payment_Plan_Configurable_Categories_Response"`
    GetOneTimePaymentPlanConfigurableCategoriesResponseType
}
```









## <a name="GetOneTimePaymentPlanConfigurableCategoriesRequestType">type</a> [GetOneTimePaymentPlanConfigurableCategoriesRequestType](/src/target/wwsgen_types.go?s=156068:156813#L2448)
``` go
type GetOneTimePaymentPlanConfigurableCategoriesRequestType struct {
    RequestReferences *OneTimePaymentPlanConfigurableCategoryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *OneTimePaymentPlanConfigurableCategoryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *OneTimePaymentPlanConfigurableCategoryResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request element










## <a name="GetOneTimePaymentPlanConfigurableCategoriesResponseType">type</a> [GetOneTimePaymentPlanConfigurableCategoriesResponseType](/src/target/wwsgen_types.go?s=156827:157842#L2457)
``` go
type GetOneTimePaymentPlanConfigurableCategoriesResponseType struct {
    RequestReferences *OneTimePaymentPlanConfigurableCategoryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *OneTimePaymentPlanConfigurableCategoryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *OneTimePaymentPlanConfigurableCategoryResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    ResponseResults   *ResponseResultsType                                         `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      *OneTimePaymentPlanConfigurableCategoryResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response










## <a name="GetPeriodActivityRateMatricesInput">type</a> [GetPeriodActivityRateMatricesInput](/src/target/wwsgen_client.go?s=29928:30107#L781)
``` go
type GetPeriodActivityRateMatricesInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Period_Activity_Rate_Matrices_Request"`
    GetPeriodActivityRateMatricesRequestType
}
```









## <a name="GetPeriodActivityRateMatricesOutput">type</a> [GetPeriodActivityRateMatricesOutput](/src/target/wwsgen_client.go?s=30109:30291#L786)
``` go
type GetPeriodActivityRateMatricesOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Period_Activity_Rate_Matrices_Response"`
    GetPeriodActivityRateMatricesResponseType
}
```









## <a name="GetPeriodActivityRateMatricesRequestType">type</a> [GetPeriodActivityRateMatricesRequestType](/src/target/wwsgen_types.go?s=158097:158638#L2468)
``` go
type GetPeriodActivityRateMatricesRequestType struct {
    RequestReferences *PeriodActivityRateMatrixRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *PeriodActivityRateMatrixRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    Version           string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request element to get period activity rate matrix information. If no criteria is specified in either the period activity rate matrix request references or period activity rate matrix request criteria elements, all active instances will be returned.










## <a name="GetPeriodActivityRateMatricesResponseType">type</a> [GetPeriodActivityRateMatricesResponseType](/src/target/wwsgen_types.go?s=158682:159465#L2476)
``` go
type GetPeriodActivityRateMatricesResponseType struct {
    RequestReferences *PeriodActivityRateMatrixRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *PeriodActivityRateMatrixRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseResults   *ResponseResultsType                           `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      *PeriodActivityRateMatrixResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Period Activity Rate Matrices response










## <a name="GetPeriodActivityTasksInput">type</a> [GetPeriodActivityTasksInput](/src/target/wwsgen_client.go?s=26380:26537#L693)
``` go
type GetPeriodActivityTasksInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Period_Activity_Tasks_Request"`
    GetPeriodActivityTasksRequestType
}
```









## <a name="GetPeriodActivityTasksOutput">type</a> [GetPeriodActivityTasksOutput](/src/target/wwsgen_client.go?s=26539:26699#L698)
``` go
type GetPeriodActivityTasksOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Period_Activity_Tasks_Response"`
    GetPeriodActivityTasksResponseType
}
```









## <a name="GetPeriodActivityTasksRequestType">type</a> [GetPeriodActivityTasksRequestType](/src/target/wwsgen_types.go?s=159701:160325#L2486)
``` go
type GetPeriodActivityTasksRequestType struct {
    RequestReferences *PeriodActivityTaskRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *PeriodActivityTaskRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *PeriodActivityTaskResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request element to get period activity task information.  If no criteria are specified in either the period activity task request references or period activity task request criteria elements, all active instances will be returned.










## <a name="GetPeriodActivityTasksResponseType">type</a> [GetPeriodActivityTasksResponseType](/src/target/wwsgen_types.go?s=160420:161274#L2495)
``` go
type GetPeriodActivityTasksResponseType struct {
    RequestReferences *PeriodActivityTaskRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *PeriodActivityTaskRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *PeriodActivityTaskResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    ResponseResults   *ResponseResultsType                     `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      *PeriodActivityTaskResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element containing instances of period activity tasks and their associated data.










## <a name="GetPreviousSystemCompensationHistoryInput">type</a> [GetPreviousSystemCompensationHistoryInput](/src/target/wwsgen_client.go?s=2011:2211#L55)
``` go
type GetPreviousSystemCompensationHistoryInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Previous_System_Compensation_History_Request"`
    GetPreviousSystemCompensationHistoryRequestType
}
```









## <a name="GetPreviousSystemCompensationHistoryOutput">type</a> [GetPreviousSystemCompensationHistoryOutput](/src/target/wwsgen_client.go?s=2213:2416#L60)
``` go
type GetPreviousSystemCompensationHistoryOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Previous_System_Compensation_History_Response"`
    GetPreviousSystemCompensationHistoryResponseType
}
```









## <a name="GetPreviousSystemCompensationHistoryRequestType">type</a> [GetPreviousSystemCompensationHistoryRequestType](/src/target/wwsgen_types.go?s=161487:161993#L2507)
``` go
type GetPreviousSystemCompensationHistoryRequestType struct {
    RequestReferences *WorkerRequestReferencesType         `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *WorkerResponseGroupforReferenceType `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request element used to return a worker's compensation history loaded from a previous system.
If no worker reference is requested, then all workers with previous system compensation data will be returned.










## <a name="GetPreviousSystemCompensationHistoryResponseType">type</a> [GetPreviousSystemCompensationHistoryResponseType](/src/target/wwsgen_types.go?s=162099:162917#L2515)
``` go
type GetPreviousSystemCompensationHistoryResponseType struct {
    RequestReferences []WorkerRequestReferencesType                       `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    ResponseFilter    []ResponseFilterType                                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseResults   []ResponseResultsType                               `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseGroup     []WorkerResponseGroupforReferenceType               `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    ResponseData      []PreviousSystemCompensationHistoryResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element containing the instances of workers and their previous system compensation history.










## <a name="GetStockParticipationRateTablesInput">type</a> [GetStockParticipationRateTablesInput](/src/target/wwsgen_client.go?s=32720:32905#L847)
``` go
type GetStockParticipationRateTablesInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Stock_Participation_Rate_Tables_Request"`
    GetStockParticipationRateTablesRequestType
}
```









## <a name="GetStockParticipationRateTablesOutput">type</a> [GetStockParticipationRateTablesOutput](/src/target/wwsgen_client.go?s=32907:33095#L852)
``` go
type GetStockParticipationRateTablesOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Get_Stock_Participation_Rate_Tables_Response"`
    GetStockParticipationRateTablesResponseType
}
```









## <a name="GetStockParticipationRateTablesRequestType">type</a> [GetStockParticipationRateTablesRequestType](/src/target/wwsgen_types.go?s=163133:163811#L2525)
``` go
type GetStockParticipationRateTablesRequestType struct {
    RequestReferences *StockParticipationRateTableRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *StockParticipationRateTableRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *StockParticipationRateTableResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request element to get stock participation rate table information.  If no criteria is specified in the request references or request criteria elements, then all stock participation rate tables will be returned.










## <a name="GetStockParticipationRateTablesResponseType">type</a> [GetStockParticipationRateTablesResponseType](/src/target/wwsgen_types.go?s=163916:164842#L2534)
``` go
type GetStockParticipationRateTablesResponseType struct {
    RequestReferences *StockParticipationRateTableRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *StockParticipationRateTableRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *StockParticipationRateTableResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    ResponseResults   *ResponseResultsType                              `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      *StockParticipationRateTableResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element containing instances of stock participation rate tables and their associated data.










## <a name="GrantTypeSplitReplacementDataType">type</a> [GrantTypeSplitReplacementDataType](/src/target/wwsgen_types.go?s=164874:165427#L2545)
``` go
type GrantTypeSplitReplacementDataType struct {
    StockGrantTypeReference       *StockGrantTypeObjectType       `xml:"urn:com.workday/bsvc Stock_Grant_Type_Reference"`
    StockVestingScheduleReference *StockVestingScheduleObjectType `xml:"urn:com.workday/bsvc Stock_Vesting_Schedule_Reference,omitempty"`
    StockDateRuleReference        *StockDateRuleObjectType        `xml:"urn:com.workday/bsvc Stock_Date_Rule_Reference,omitempty"`
    GrantSplitPercent             float64                         `xml:"urn:com.workday/bsvc Grant_Split_Percent,omitempty"`
}
```
Data for Grant Type Split.










## <a name="HourlyPlanDataType">type</a> [HourlyPlanDataType](/src/target/wwsgen_types.go?s=165471:165645#L2553)
``` go
type HourlyPlanDataType struct {
    Amount      float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
    MinimumWage *bool   `xml:"urn:com.workday/bsvc Minimum_Wage,omitempty"`
}
```
This is a wrapper for the Hourly Data.










## <a name="ImportEligibleEarningsOverrideHVRequestType">type</a> [ImportEligibleEarningsOverrideHVRequestType](/src/target/wwsgen_types.go?s=165703:166007#L2559)
``` go
type ImportEligibleEarningsOverrideHVRequestType struct {
    EligibleEarningsReference *EligibleEarningsOverrideObjectType `xml:"urn:com.workday/bsvc Eligible_Earnings_Reference,omitempty"`
    EligibleEarningsData      *EligibleEarningsOverrideHVDataType `xml:"urn:com.workday/bsvc Eligible_Earnings_Data"`
}
```
This element contains an Eligible Earnings Override.










## <a name="ImportEligibleEarningsOverrideInput">type</a> [ImportEligibleEarningsOverrideInput](/src/target/wwsgen_client.go?s=39152:39333#L1001)
``` go
type ImportEligibleEarningsOverrideInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Import_Eligible_Earnings_Override_Request"`
    ImportEligibleEarningsOverrideRequestType
}
```









## <a name="ImportEligibleEarningsOverrideOutput">type</a> [ImportEligibleEarningsOverrideOutput](/src/target/wwsgen_client.go?s=39335:39490#L1006)
``` go
type ImportEligibleEarningsOverrideOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
    PutImportProcessResponseType
}
```









## <a name="ImportEligibleEarningsOverrideRequestType">type</a> [ImportEligibleEarningsOverrideRequestType](/src/target/wwsgen_types.go?s=166517:167017#L2565)
``` go
type ImportEligibleEarningsOverrideRequestType struct {
    EligibleEarningsOverridePeriodReference *EligibleEarningsOverridePeriodObjectType     `xml:"urn:com.workday/bsvc Eligible_Earnings_Override_Period_Reference"`
    EligibleEarnings                        []ImportEligibleEarningsOverrideHVRequestType `xml:"urn:com.workday/bsvc Eligible_Earnings,omitempty"`
    Version                                 string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Loads eligible earnings overrides for a specific period. To apply the override to a specific bonus plan, add the bonus plan to the operation filter. To apply the override to all current and future bonus plans, select "Apply to All Bonus Plans”. An in-progress Compensation Review Process immediately implements these overrides if the Participation Options are configured to support this. Unlike the Put Eligible Earnings operation, if the input template contains any errors, the entire operation fails.










## <a name="ImportRequestCompensationChangeInput">type</a> [ImportRequestCompensationChangeInput](/src/target/wwsgen_client.go?s=39995:40179#L1023)
``` go
type ImportRequestCompensationChangeInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Import_Request_Compensation_Change_Request"`
    ImportRequestCompensationChangeRequestType
}
```









## <a name="ImportRequestCompensationChangeOutput">type</a> [ImportRequestCompensationChangeOutput](/src/target/wwsgen_client.go?s=40181:40337#L1028)
``` go
type ImportRequestCompensationChangeOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
    PutImportProcessResponseType
}
```









## <a name="ImportRequestCompensationChangeRequestType">type</a> [ImportRequestCompensationChangeRequestType](/src/target/wwsgen_types.go?s=167078:167559#L2572)
``` go
type ImportRequestCompensationChangeRequestType struct {
    ID                                      string                                   `xml:"urn:com.workday/bsvc ID,omitempty"`
    RequestCompensationRequestInformationHV []RequestCompensationChangeRequestHVType `xml:"urn:com.workday/bsvc Request_Compensation_Request_Information_HV,omitempty"`
    Version                                 string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Root Import Request Compensation Change Request Element










## <a name="InstanceIDType">type</a> [InstanceIDType](/src/target/wwsgen_types.go?s=167561:167844#L2578)
``` go
type InstanceIDType struct {
    Value      string `xml:",chardata"`
    Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
    Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
    Type       string `xml:"urn:com.workday/bsvc type,attr,omitempty"`
}
```









## <a name="InstanceObjectType">type</a> [InstanceObjectType](/src/target/wwsgen_types.go?s=167846:168035#L2585)
``` go
type InstanceObjectType struct {
    ID         []InstanceIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="JobProfileObjectIDType">type</a> [JobProfileObjectIDType](/src/target/wwsgen_types.go?s=168099:168222#L2591)
``` go
type JobProfileObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="JobProfileObjectType">type</a> [JobProfileObjectType](/src/target/wwsgen_types.go?s=168224:168431#L2596)
``` go
type JobProfileObjectType struct {
    ID         []JobProfileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="JobRequisitionObjectIDType">type</a> [JobRequisitionObjectIDType](/src/target/wwsgen_types.go?s=168495:168622#L2602)
``` go
type JobRequisitionObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="JobRequisitionObjectType">type</a> [JobRequisitionObjectType](/src/target/wwsgen_types.go?s=168624:168843#L2607)
``` go
type JobRequisitionObjectType struct {
    ID         []JobRequisitionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="ManagePeriodActivityPayAssignmentsInput">type</a> [ManagePeriodActivityPayAssignmentsInput](/src/target/wwsgen_client.go?s=30933:31127#L803)
``` go
type ManagePeriodActivityPayAssignmentsInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Manage_Period_Activity_Pay_Assignments_Request"`
    ManagePeriodActivityPayAssignmentsRequestType
}
```









## <a name="ManagePeriodActivityPayAssignmentsOutput">type</a> [ManagePeriodActivityPayAssignmentsOutput](/src/target/wwsgen_client.go?s=31129:31326#L808)
``` go
type ManagePeriodActivityPayAssignmentsOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Manage_Period_Activity_Pay_Assignments_Response"`
    ManagePeriodActivityPayAssignmentsResponseType
}
```









## <a name="ManagePeriodActivityPayAssignmentsRequestType">type</a> [ManagePeriodActivityPayAssignmentsRequestType](/src/target/wwsgen_types.go?s=169042:169477#L2613)
``` go
type ManagePeriodActivityPayAssignmentsRequestType struct {
    BusinessProcessParameters    *BusinessProcessParametersType        `xml:"urn:com.workday/bsvc Business_Process_Parameters"`
    PeriodActivityPayAssignments *PeriodActivityPayAssignmentsDataType `xml:"urn:com.workday/bsvc Period_Activity_Pay_Assignments"`
    Version                      string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
This operation allows the management of period activity based assignments for a given worker and a bounding date range or period via the Manage Period Activity Pay Assignments business process.










## <a name="ManagePeriodActivityPayAssignmentsResponseType">type</a> [ManagePeriodActivityPayAssignmentsResponseType](/src/target/wwsgen_types.go?s=169615:169948#L2620)
``` go
type ManagePeriodActivityPayAssignmentsResponseType struct {
    PeriodActivityPayAssigmentEvent *PeriodActivityAssignmentEventDataType `xml:"urn:com.workday/bsvc Period_Activity_Pay_Assigment_Event,omitempty"`
    Version                         string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Contains the detailed business process event information resulting from a successful Manage Period Activity Pay Assignments request.










## <a name="ManagementLevelObjectIDType">type</a> [ManagementLevelObjectIDType](/src/target/wwsgen_types.go?s=170012:170140#L2626)
``` go
type ManagementLevelObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="ManagementLevelObjectType">type</a> [ManagementLevelObjectType](/src/target/wwsgen_types.go?s=170142:170364#L2631)
``` go
type ManagementLevelObjectType struct {
    ID         []ManagementLevelObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="MeritPercentPlanObjectIDType">type</a> [MeritPercentPlanObjectIDType](/src/target/wwsgen_types.go?s=170428:170557#L2637)
``` go
type MeritPercentPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="MeritPercentPlanObjectType">type</a> [MeritPercentPlanObjectType](/src/target/wwsgen_types.go?s=170559:170784#L2642)
``` go
type MeritPercentPlanObjectType struct {
    ID         []MeritPercentPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="MeritPlanAmountDataType">type</a> [MeritPlanAmountDataType](/src/target/wwsgen_types.go?s=170826:171467#L2648)
``` go
type MeritPlanAmountDataType struct {
    Amount                                float64                                     `xml:"urn:com.workday/bsvc Amount,omitempty"`
    CompensationBasisReference            *CompensationBasisObjectType                `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
    ManageByCompensationBasisRules        *bool                                       `xml:"urn:com.workday/bsvc Manage_By_Compensation_Basis_Rules,omitempty"`
    MeritPlanProfileAmountReplacementData []MeritPlanProfileAmountReplacementDataType `xml:"urn:com.workday/bsvc Merit_Plan_Profile_Amount_Replacement_Data,omitempty"`
}
```
Data specific to a Merit Amount Plan










## <a name="MeritPlanDataType">type</a> [MeritPlanDataType](/src/target/wwsgen_types.go?s=171488:174011#L2656)
``` go
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
```
Merit Plan Data










## <a name="MeritPlanObjectIDType">type</a> [MeritPlanObjectIDType](/src/target/wwsgen_types.go?s=174075:174197#L2675)
``` go
type MeritPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="MeritPlanObjectType">type</a> [MeritPlanObjectType](/src/target/wwsgen_types.go?s=174199:174403#L2680)
``` go
type MeritPlanObjectType struct {
    ID         []MeritPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="MeritPlanPercentDataType">type</a> [MeritPlanPercentDataType](/src/target/wwsgen_types.go?s=174503:175158#L2686)
``` go
type MeritPlanPercentDataType struct {
    Percentage                             float64                                      `xml:"urn:com.workday/bsvc Percentage,omitempty"`
    ManageByCompensationBasisRules         *bool                                        `xml:"urn:com.workday/bsvc Manage_By_Compensation_Basis_Rules,omitempty"`
    CompensationBasisReference             *CompensationBasisObjectType                 `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
    MeritPlanProfilePercentReplacementData []MeritPlanProfilePercentReplacementDataType `xml:"urn:com.workday/bsvc Merit_Plan_Profile_Percent_Replacement_Data,omitempty"`
}
```
Required Percent Data for Merit Plan; consists of default percent and/or  Merit Plan Profiles.










## <a name="MeritPlanProfileAmountReplacementDataType">type</a> [MeritPlanProfileAmountReplacementDataType](/src/target/wwsgen_types.go?s=175201:175710#L2694)
``` go
type MeritPlanProfileAmountReplacementDataType struct {
    ProfileAmount              float64                      `xml:"urn:com.workday/bsvc Profile_Amount,omitempty"`
    CurrencyReference          *CurrencyObjectType          `xml:"urn:com.workday/bsvc Currency_Reference"`
    ConditionRuleReference     *ConditionRuleObjectType     `xml:"urn:com.workday/bsvc Condition_Rule_Reference"`
    CompensationBasisReference *CompensationBasisObjectType `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
}
```
Data for a Merit Amount plan profile.










## <a name="MeritPlanProfilePercentReplacementDataType">type</a> [MeritPlanProfilePercentReplacementDataType](/src/target/wwsgen_types.go?s=175848:176251#L2702)
``` go
type MeritPlanProfilePercentReplacementDataType struct {
    EligibilityRuleReference   *ConditionRuleObjectType     `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    Percentage                 float64                      `xml:"urn:com.workday/bsvc Percentage,omitempty"`
    CompensationBasisReference *CompensationBasisObjectType `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
}
```
This is a wrapper for the Merit Plan Percent Plan Profile Percent Replacement; used to display or replace Percent Plan Profile Data.










## <a name="OneTimePaymentDataType">type</a> [OneTimePaymentDataType](/src/target/wwsgen_types.go?s=176317:177597#L2709)
``` go
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
```
Data Element that contains the one-time payment information.










### <a name="OneTimePaymentDataType.MarshalXML">func</a> (\*OneTimePaymentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=177599:177688#L2722)
``` go
func (t *OneTimePaymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="OneTimePaymentDataType.UnmarshalXML">func</a> (\*OneTimePaymentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=178289:178380#L2736)
``` go
func (t *OneTimePaymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="OneTimePaymentPlanConfigurableCategoryDataType">type</a> [OneTimePaymentPlanConfigurableCategoryDataType](/src/target/wwsgen_types.go?s=179019:179195#L2752)
``` go
type OneTimePaymentPlanConfigurableCategoryDataType struct {
    ID       string `xml:"urn:com.workday/bsvc ID,omitempty"`
    Category string `xml:"urn:com.workday/bsvc Category"`
}
```
Data that is persisted










## <a name="OneTimePaymentPlanConfigurableCategoryObjectIDType">type</a> [OneTimePaymentPlanConfigurableCategoryObjectIDType](/src/target/wwsgen_types.go?s=179259:179410#L2758)
``` go
type OneTimePaymentPlanConfigurableCategoryObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="OneTimePaymentPlanConfigurableCategoryObjectType">type</a> [OneTimePaymentPlanConfigurableCategoryObjectType](/src/target/wwsgen_types.go?s=179412:179703#L2763)
``` go
type OneTimePaymentPlanConfigurableCategoryObjectType struct {
    ID         []OneTimePaymentPlanConfigurableCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="OneTimePaymentPlanConfigurableCategoryRequestCriteriaType">type</a> [OneTimePaymentPlanConfigurableCategoryRequestCriteriaType](/src/target/wwsgen_types.go?s=179725:179798#L2769)
``` go
type OneTimePaymentPlanConfigurableCategoryRequestCriteriaType struct {
}
```
Request Criteria










## <a name="OneTimePaymentPlanConfigurableCategoryRequestReferencesType">type</a> [OneTimePaymentPlanConfigurableCategoryRequestReferencesType](/src/target/wwsgen_types.go?s=179822:180080#L2773)
``` go
type OneTimePaymentPlanConfigurableCategoryRequestReferencesType struct {
    OneTimePaymentPlanConfigurableCategoryReference []OneTimePaymentPlanConfigurableCategoryObjectType `xml:"urn:com.workday/bsvc One-Time_Payment_Plan_Configurable_Category_Reference"`
}
```
Request References










## <a name="OneTimePaymentPlanConfigurableCategoryResponseDataType">type</a> [OneTimePaymentPlanConfigurableCategoryResponseDataType](/src/target/wwsgen_types.go?s=180102:180340#L2778)
``` go
type OneTimePaymentPlanConfigurableCategoryResponseDataType struct {
    OneTimePaymentPlanConfigurableCategory []OneTimePaymentPlanConfigurableCategoryType `xml:"urn:com.workday/bsvc One-Time_Payment_Plan_Configurable_Category,omitempty"`
}
```
Response Element










## <a name="OneTimePaymentPlanConfigurableCategoryResponseGroupType">type</a> [OneTimePaymentPlanConfigurableCategoryResponseGroupType](/src/target/wwsgen_types.go?s=180360:180512#L2783)
``` go
type OneTimePaymentPlanConfigurableCategoryResponseGroupType struct {
    IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}
```
Response Group










## <a name="OneTimePaymentPlanConfigurableCategoryType">type</a> [OneTimePaymentPlanConfigurableCategoryType](/src/target/wwsgen_types.go?s=180531:180968#L2788)
``` go
type OneTimePaymentPlanConfigurableCategoryType struct {
    OneTimePaymentPlanConfigurableCategoryReference *OneTimePaymentPlanConfigurableCategoryObjectType `xml:"urn:com.workday/bsvc One-Time_Payment_Plan_Configurable_Category_Reference,omitempty"`
    OneTimePaymentPlanConfigurableCategoryData      *OneTimePaymentPlanConfigurableCategoryDataType   `xml:"urn:com.workday/bsvc One-Time_Payment_Plan_Configurable_Category_Data,omitempty"`
}
```
Response Data










## <a name="OneTimePaymentPlanObjectIDType">type</a> [OneTimePaymentPlanObjectIDType](/src/target/wwsgen_types.go?s=181032:181163#L2794)
``` go
type OneTimePaymentPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="OneTimePaymentPlanObjectType">type</a> [OneTimePaymentPlanObjectType](/src/target/wwsgen_types.go?s=181165:181396#L2799)
``` go
type OneTimePaymentPlanObjectType struct {
    ID         []OneTimePaymentPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PayRangeQuartileObjectIDType">type</a> [PayRangeQuartileObjectIDType](/src/target/wwsgen_types.go?s=181460:181589#L2805)
``` go
type PayRangeQuartileObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="PayRangeQuartileObjectType">type</a> [PayRangeQuartileObjectType](/src/target/wwsgen_types.go?s=181591:181816#L2810)
``` go
type PayRangeQuartileObjectType struct {
    ID         []PayRangeQuartileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PayRangeSegmentTypeObjectIDType">type</a> [PayRangeSegmentTypeObjectIDType](/src/target/wwsgen_types.go?s=181880:182012#L2816)
``` go
type PayRangeSegmentTypeObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="PayRangeSegmentTypeObjectType">type</a> [PayRangeSegmentTypeObjectType](/src/target/wwsgen_types.go?s=182014:182248#L2821)
``` go
type PayRangeSegmentTypeObjectType struct {
    ID         []PayRangeSegmentTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PerformanceCriteriaDataType">type</a> [PerformanceCriteriaDataType](/src/target/wwsgen_types.go?s=182288:182614#L2827)
``` go
type PerformanceCriteriaDataType struct {
    GoalID          string  `xml:"urn:com.workday/bsvc Goal_ID,omitempty"`
    GoalName        string  `xml:"urn:com.workday/bsvc Goal_Name"`
    GoalDescription string  `xml:"urn:com.workday/bsvc Goal_Description,omitempty"`
    GoalWeight      float64 `xml:"urn:com.workday/bsvc Goal_Weight"`
}
```
Contains information on scorecard.










## <a name="PerformanceCriteriaObjectIDType">type</a> [PerformanceCriteriaObjectIDType](/src/target/wwsgen_types.go?s=182678:182810#L2835)
``` go
type PerformanceCriteriaObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="PerformanceCriteriaObjectType">type</a> [PerformanceCriteriaObjectType](/src/target/wwsgen_types.go?s=182812:183046#L2840)
``` go
type PerformanceCriteriaObjectType struct {
    ID         []PerformanceCriteriaObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PerformanceFactorCompensationMatrixWeightedDataType">type</a> [PerformanceFactorCompensationMatrixWeightedDataType](/src/target/wwsgen_types.go?s=183264:183897#L2846)
``` go
type PerformanceFactorCompensationMatrixWeightedDataType struct {
    CompensationMatrixWeightedReference *CompensationWeightedPercentMatrixObjectType `xml:"urn:com.workday/bsvc Compensation_Matrix_Weighted_Reference,omitempty"`
    ApplyNetAttainment                  *bool                                        `xml:"urn:com.workday/bsvc Apply_Net_Attainment,omitempty"`
    Weight                              float64                                      `xml:"urn:com.workday/bsvc Weight"`
    BonusModifierReference              *DefaultScorecardObjectType                  `xml:"urn:com.workday/bsvc Bonus_Modifier_Reference,omitempty"`
}
```
Performance Factor Compensation Matrix Weighted Data is part of Performance Factors Data within a Bonus Plan; there must be one Performance Factor Compensation Matrix Weighted Data for a Performance Factors Data.










## <a name="PerformanceFactorScorecardDataType">type</a> [PerformanceFactorScorecardDataType](/src/target/wwsgen_types.go?s=184092:184447#L2854)
``` go
type PerformanceFactorScorecardDataType struct {
    ScorecardReference     *DefaultScorecardObjectType `xml:"urn:com.workday/bsvc Scorecard_Reference"`
    Weight                 float64                     `xml:"urn:com.workday/bsvc Weight"`
    BonusModifierReference *DefaultScorecardObjectType `xml:"urn:com.workday/bsvc Bonus_Modifier_Reference,omitempty"`
}
```
Performance Factor Scorecard Data is part of Performance Factors Data within a Bonus Plan.  There must be at least one Performance Factor Scorecard Data within the Performance Factors Data.










## <a name="PerformanceFactorsDataType">type</a> [PerformanceFactorsDataType](/src/target/wwsgen_types.go?s=184679:185090#L2861)
``` go
type PerformanceFactorsDataType struct {
    PerformanceFactorScorecardData                  []PerformanceFactorScorecardDataType                 `xml:"urn:com.workday/bsvc Performance_Factor_Scorecard_Data,omitempty"`
    PerformanceFactorCompensationMatrixWeightedData *PerformanceFactorCompensationMatrixWeightedDataType `xml:"urn:com.workday/bsvc Performance_Factor_Compensation_Matrix_Weighted_Data,omitempty"`
}
```
Performance Factors Data is part of Bonus Plan and is mutually exclusive with Performance Matrix Data.  It consists of one or more Performance Factor Scorecard Data and one Performance Factor Compensation Matrix Weighted Data.










## <a name="PerformanceMatrixDataType">type</a> [PerformanceMatrixDataType](/src/target/wwsgen_types.go?s=185208:185746#L2867)
``` go
type PerformanceMatrixDataType struct {
    CompensationMatrixReference *CompensationMatrixObjectType `xml:"urn:com.workday/bsvc Compensation_Matrix_Reference,omitempty"`
    UseMatrixasReferenceOnly    *bool                         `xml:"urn:com.workday/bsvc Use_Matrix_as_Reference_Only,omitempty"`
    ApplyNetAttainment          *bool                         `xml:"urn:com.workday/bsvc Apply_Net_Attainment,omitempty"`
    BonusModifierReference      *DefaultScorecardObjectType   `xml:"urn:com.workday/bsvc Bonus_Modifier_Reference,omitempty"`
}
```
This is a wrapper for the Bonus Plan Performance Matrix and is mutually exclusive with Performance Factors Data.










## <a name="PeriodActivityAssignmentDeletedDataType">type</a> [PeriodActivityAssignmentDeletedDataType](/src/target/wwsgen_types.go?s=185883:186752#L2875)
``` go
type PeriodActivityAssignmentDeletedDataType struct {
    PeriodActivityAssignmentReference *PeriodActivityAssignmentObjectType `xml:"urn:com.workday/bsvc Period_Activity_Assignment_Reference,omitempty"`
    PeriodActivityReference           *PeriodActivityObjectType           `xml:"urn:com.workday/bsvc Period_Activity_Reference,omitempty"`
    PeriodActivityTaskReference       *PeriodActivityTaskObjectType       `xml:"urn:com.workday/bsvc Period_Activity_Task_Reference,omitempty"`
    PeriodActivityStartDate           *time.Time                          `xml:"urn:com.workday/bsvc Period_Activity_Start_Date,omitempty"`
    PeriodActivityEndDate             *time.Time                          `xml:"urn:com.workday/bsvc Period_Activity_End_Date,omitempty"`
    Comment                           string                              `xml:"urn:com.workday/bsvc Comment,omitempty"`
}
```
Contains assignments which have been removed from the employee.  An assignment can only be removed if no pay activity has occurred.










### <a name="PeriodActivityAssignmentDeletedDataType.MarshalXML">func</a> (\*PeriodActivityAssignmentDeletedDataType) [MarshalXML](/src/target/wwsgen_types.go?s=186754:186860#L2884)
``` go
func (t *PeriodActivityAssignmentDeletedDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="PeriodActivityAssignmentDeletedDataType.UnmarshalXML">func</a> (\*PeriodActivityAssignmentDeletedDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=187355:187463#L2896)
``` go
func (t *PeriodActivityAssignmentDeletedDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="PeriodActivityAssignmentEventDataType">type</a> [PeriodActivityAssignmentEventDataType](/src/target/wwsgen_types.go?s=188089:189479#L2910)
``` go
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
```
Contains a reference to the Manage Period Activity Pay Assignments business process event and the assignment details.










### <a name="PeriodActivityAssignmentEventDataType.MarshalXML">func</a> (\*PeriodActivityAssignmentEventDataType) [MarshalXML](/src/target/wwsgen_types.go?s=189481:189585#L2922)
``` go
func (t *PeriodActivityAssignmentEventDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="PeriodActivityAssignmentEventDataType.UnmarshalXML">func</a> (\*PeriodActivityAssignmentEventDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=189862:189968#L2932)
``` go
func (t *PeriodActivityAssignmentEventDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="PeriodActivityAssignmentObjectIDType">type</a> [PeriodActivityAssignmentObjectIDType](/src/target/wwsgen_types.go?s=190315:190452#L2944)
``` go
type PeriodActivityAssignmentObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="PeriodActivityAssignmentObjectType">type</a> [PeriodActivityAssignmentObjectType](/src/target/wwsgen_types.go?s=190454:190703#L2949)
``` go
type PeriodActivityAssignmentObjectType struct {
    ID         []PeriodActivityAssignmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PeriodActivityAssignmentVersionDataType">type</a> [PeriodActivityAssignmentVersionDataType](/src/target/wwsgen_types.go?s=190842:192784#L2955)
``` go
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
```
Contains the individual assignment details of a new activity, changes to an existing activity, or an existing activity to be deleted.










### <a name="PeriodActivityAssignmentVersionDataType.MarshalXML">func</a> (\*PeriodActivityAssignmentVersionDataType) [MarshalXML](/src/target/wwsgen_types.go?s=192786:192892#L2972)
``` go
func (t *PeriodActivityAssignmentVersionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="PeriodActivityAssignmentVersionDataType.UnmarshalXML">func</a> (\*PeriodActivityAssignmentVersionDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=193697:193805#L2988)
``` go
func (t *PeriodActivityAssignmentVersionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="PeriodActivityCategoryObjectIDType">type</a> [PeriodActivityCategoryObjectIDType](/src/target/wwsgen_types.go?s=194686:194821#L3006)
``` go
type PeriodActivityCategoryObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="PeriodActivityCategoryObjectType">type</a> [PeriodActivityCategoryObjectType](/src/target/wwsgen_types.go?s=194823:195066#L3011)
``` go
type PeriodActivityCategoryObjectType struct {
    ID         []PeriodActivityCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PeriodActivityEligibilityEntryDataType">type</a> [PeriodActivityEligibilityEntryDataType](/src/target/wwsgen_types.go?s=195135:195553#L3017)
``` go
type PeriodActivityEligibilityEntryDataType struct {
    PeriodActivityReference     *PeriodActivityObjectType     `xml:"urn:com.workday/bsvc Period_Activity_Reference"`
    PeriodActivityTaskReference *PeriodActivityTaskObjectType `xml:"urn:com.workday/bsvc Period_Activity_Task_Reference,omitempty"`
    AssignedUnitQuantity        float64                       `xml:"urn:com.workday/bsvc Assigned_Unit_Quantity,omitempty"`
}
```
Encapsulating element containing eligible period activity data.










## <a name="PeriodActivityObjectIDType">type</a> [PeriodActivityObjectIDType](/src/target/wwsgen_types.go?s=195617:195744#L3024)
``` go
type PeriodActivityObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="PeriodActivityObjectType">type</a> [PeriodActivityObjectType](/src/target/wwsgen_types.go?s=195746:195965#L3029)
``` go
type PeriodActivityObjectType struct {
    ID         []PeriodActivityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PeriodActivityPayAssignmentDataType">type</a> [PeriodActivityPayAssignmentDataType](/src/target/wwsgen_types.go?s=196104:198320#L3035)
``` go
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
```
Contains the individual assignment details of a new activity, changes to an existing activity, or an existing activity to be deleted.










### <a name="PeriodActivityPayAssignmentDataType.MarshalXML">func</a> (\*PeriodActivityPayAssignmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=198322:198424#L3053)
``` go
func (t *PeriodActivityPayAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="PeriodActivityPayAssignmentDataType.UnmarshalXML">func</a> (\*PeriodActivityPayAssignmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=199225:199329#L3069)
``` go
func (t *PeriodActivityPayAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="PeriodActivityPayAssignmentsDataType">type</a> [PeriodActivityPayAssignmentsDataType](/src/target/wwsgen_types.go?s=200317:201286#L3087)
``` go
type PeriodActivityPayAssignmentsDataType struct {
    EffectiveDate                     time.Time                                 `xml:"urn:com.workday/bsvc Effective_Date"`
    EmployeeReference                 *EmployeeObjectType                       `xml:"urn:com.workday/bsvc Employee_Reference"`
    PositionReference                 *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
    AcademicPeriodReference           *AcademicPeriodObjectType                 `xml:"urn:com.workday/bsvc Academic_Period_Reference"`
    PeriodActivityRateMatrixReference *PeriodActivityRateMatrixObjectType       `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Reference"`
    ReasonReference                   *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference"`
    PeriodActivityPayAssignment       []PeriodActivityPayAssignmentDataType     `xml:"urn:com.workday/bsvc Period_Activity_Pay_Assignment"`
}
```
Contains Period Activity Assignment information for a given Worker and Period range in a format suitable for constructing a Manage Period Activity Pay operation request.










### <a name="PeriodActivityPayAssignmentsDataType.MarshalXML">func</a> (\*PeriodActivityPayAssignmentsDataType) [MarshalXML](/src/target/wwsgen_types.go?s=201288:201391#L3097)
``` go
func (t *PeriodActivityPayAssignmentsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="PeriodActivityPayAssignmentsDataType.UnmarshalXML">func</a> (\*PeriodActivityPayAssignmentsDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=201658:201763#L3107)
``` go
func (t *PeriodActivityPayAssignmentsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="PeriodActivityPayCostingDataType">type</a> [PeriodActivityPayCostingDataType](/src/target/wwsgen_types.go?s=202141:202320#L3119)
``` go
type PeriodActivityPayCostingDataType struct {
    AllocationDetailforPeriodPay []AllocationDetailforPeriodPayDataType `xml:"urn:com.workday/bsvc Allocation_Detail_for_Period_Pay"`
}
```
Contains one or more costing allocation splits for the specific Period Activity Pay assignment line










## <a name="PeriodActivityRateMatrixDataType">type</a> [PeriodActivityRateMatrixDataType](/src/target/wwsgen_types.go?s=202358:203316#L3124)
``` go
type PeriodActivityRateMatrixDataType struct {
    ID                                string                                  `xml:"urn:com.workday/bsvc ID,omitempty"`
    PeriodActivityRateTableName       string                                  `xml:"urn:com.workday/bsvc Period_Activity_Rate_Table_Name"`
    EffectiveDate                     *time.Time                              `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
    ConditionRuleReference            []ConditionRuleObjectType               `xml:"urn:com.workday/bsvc Condition_Rule_Reference,omitempty"`
    CurrencyReference                 *CurrencyObjectType                     `xml:"urn:com.workday/bsvc Currency_Reference"`
    Inactive                          *bool                                   `xml:"urn:com.workday/bsvc Inactive,omitempty"`
    PeriodActivityRateMatrixEntryData []PeriodActivityRateMatrixEntryDataType `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Entry_Data"`
}
```
Period Activity Rate Matrix data










### <a name="PeriodActivityRateMatrixDataType.MarshalXML">func</a> (\*PeriodActivityRateMatrixDataType) [MarshalXML](/src/target/wwsgen_types.go?s=203318:203417#L3134)
``` go
func (t *PeriodActivityRateMatrixDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="PeriodActivityRateMatrixDataType.UnmarshalXML">func</a> (\*PeriodActivityRateMatrixDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=203689:203790#L3144)
``` go
func (t *PeriodActivityRateMatrixDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="PeriodActivityRateMatrixEntryDataType">type</a> [PeriodActivityRateMatrixEntryDataType](/src/target/wwsgen_types.go?s=204139:204663#L3156)
``` go
type PeriodActivityRateMatrixEntryDataType struct {
    PeriodActivityCategoryReference           *PeriodActivityCategoryObjectType              `xml:"urn:com.workday/bsvc Period_Activity_Category_Reference"`
    PeriodActivityUnitReference               *PeriodActivityUnitObjectType                  `xml:"urn:com.workday/bsvc Period_Activity_Unit_Reference"`
    PeriodActivityRateMatrixEntrySequenceData *PeriodActivityRateMatrixEntrySequenceDataType `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Entry_Sequence_Data"`
}
```
Period Activity Rate Matrix entry data for each row in the matrix










## <a name="PeriodActivityRateMatrixEntrySequenceDataType">type</a> [PeriodActivityRateMatrixEntrySequenceDataType](/src/target/wwsgen_types.go?s=204732:205249#L3163)
``` go
type PeriodActivityRateMatrixEntrySequenceDataType struct {
    DefaultUnitRate          float64                     `xml:"urn:com.workday/bsvc Default_Unit_Rate,omitempty"`
    AcceleratorAmount        float64                     `xml:"urn:com.workday/bsvc Accelerator_Amount,omitempty"`
    AcceleratorPercent       float64                     `xml:"urn:com.workday/bsvc Accelerator_Percent,omitempty"`
    AcceleratorBaseReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Accelerator_Base_Reference,omitempty"`
}
```
Period Activity Rate Matrix entry data for rate and accelerator










## <a name="PeriodActivityRateMatrixObjectIDType">type</a> [PeriodActivityRateMatrixObjectIDType](/src/target/wwsgen_types.go?s=205313:205450#L3171)
``` go
type PeriodActivityRateMatrixObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="PeriodActivityRateMatrixObjectType">type</a> [PeriodActivityRateMatrixObjectType](/src/target/wwsgen_types.go?s=205452:205701#L3176)
``` go
type PeriodActivityRateMatrixObjectType struct {
    ID         []PeriodActivityRateMatrixObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PeriodActivityRateMatrixRequestCriteriaType">type</a> [PeriodActivityRateMatrixRequestCriteriaType](/src/target/wwsgen_types.go?s=205723:205861#L3182)
``` go
type PeriodActivityRateMatrixRequestCriteriaType struct {
    IncludeInactive *bool `xml:"urn:com.workday/bsvc Include_Inactive,omitempty"`
}
```
Request criteria










## <a name="PeriodActivityRateMatrixRequestReferencesType">type</a> [PeriodActivityRateMatrixRequestReferencesType](/src/target/wwsgen_types.go?s=205995:206195#L3187)
``` go
type PeriodActivityRateMatrixRequestReferencesType struct {
    PeriodActivityRateMatrixReference []PeriodActivityRateMatrixObjectType `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Reference"`
}
```
Utilize the period activity rate matrix to retrieve specific instance(s) of period activity rate matrix and its associated data.










## <a name="PeriodActivityRateMatrixResponseDataType">type</a> [PeriodActivityRateMatrixResponseDataType](/src/target/wwsgen_types.go?s=206249:206429#L3192)
``` go
type PeriodActivityRateMatrixResponseDataType struct {
    PeriodActivityRateMatrix []PeriodActivityRateMatrixType `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix,omitempty"`
}
```
Contains Period Activity Rate Matrix information










## <a name="PeriodActivityRateMatrixType">type</a> [PeriodActivityRateMatrixType](/src/target/wwsgen_types.go?s=206464:206789#L3197)
``` go
type PeriodActivityRateMatrixType struct {
    PeriodActivityRateMatrixReference *PeriodActivityRateMatrixObjectType `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Reference,omitempty"`
    PeriodActivityRateMatrixData      []PeriodActivityRateMatrixDataType  `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Data"`
}
```
A Period Activity Rate Matrix










## <a name="PeriodActivityTaskDataType">type</a> [PeriodActivityTaskDataType](/src/target/wwsgen_types.go?s=206830:207957#L3203)
``` go
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
```
Contains Period Activity Task data.










## <a name="PeriodActivityTaskInterfaceObjectIDType">type</a> [PeriodActivityTaskInterfaceObjectIDType](/src/target/wwsgen_types.go?s=208021:208161#L3216)
``` go
type PeriodActivityTaskInterfaceObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="PeriodActivityTaskInterfaceObjectType">type</a> [PeriodActivityTaskInterfaceObjectType](/src/target/wwsgen_types.go?s=208163:208421#L3221)
``` go
type PeriodActivityTaskInterfaceObjectType struct {
    ID         []PeriodActivityTaskInterfaceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PeriodActivityTaskObjectIDType">type</a> [PeriodActivityTaskObjectIDType](/src/target/wwsgen_types.go?s=208485:208616#L3227)
``` go
type PeriodActivityTaskObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="PeriodActivityTaskObjectType">type</a> [PeriodActivityTaskObjectType](/src/target/wwsgen_types.go?s=208618:208849#L3232)
``` go
type PeriodActivityTaskObjectType struct {
    ID         []PeriodActivityTaskObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PeriodActivityTaskRequestCriteriaType">type</a> [PeriodActivityTaskRequestCriteriaType](/src/target/wwsgen_types.go?s=208969:209101#L3238)
``` go
type PeriodActivityTaskRequestCriteriaType struct {
    IncludeInactive *bool `xml:"urn:com.workday/bsvc Include_Inactive,omitempty"`
}
```
Utilize the period activity task to retrieve specific instance(s) of period activity task and its associated data.










## <a name="PeriodActivityTaskRequestReferencesType">type</a> [PeriodActivityTaskRequestReferencesType](/src/target/wwsgen_types.go?s=209221:209396#L3243)
``` go
type PeriodActivityTaskRequestReferencesType struct {
    PeriodActivityTaskReference []PeriodActivityTaskObjectType `xml:"urn:com.workday/bsvc Period_Activity_Task_Reference"`
}
```
Utilize the period activity task to retrieve specific instance(s) of period activity task and its associated data.










## <a name="PeriodActivityTaskResponseDataType">type</a> [PeriodActivityTaskResponseDataType](/src/target/wwsgen_types.go?s=209492:209647#L3248)
``` go
type PeriodActivityTaskResponseDataType struct {
    PeriodActivityTask []PeriodActivityTaskType `xml:"urn:com.workday/bsvc Period_Activity_Task,omitempty"`
}
```
Contains period activity task information based on Request References or Request Criteria.










## <a name="PeriodActivityTaskResponseGroupType">type</a> [PeriodActivityTaskResponseGroupType](/src/target/wwsgen_types.go?s=209704:209836#L3253)
``` go
type PeriodActivityTaskResponseGroupType struct {
    IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}
```
Provides the filters for the web service operation.










## <a name="PeriodActivityTaskType">type</a> [PeriodActivityTaskType](/src/target/wwsgen_types.go?s=209898:210189#L3258)
``` go
type PeriodActivityTaskType struct {
    PeriodActivityTaskReference *PeriodActivityTaskObjectType `xml:"urn:com.workday/bsvc Period_Activity_Task_Reference,omitempty"`
    PeriodActivityTaskData      []PeriodActivityTaskDataType  `xml:"urn:com.workday/bsvc Period_Activity_Task_Data,omitempty"`
}
```
Contains a period activity task and its associated data.










## <a name="PeriodActivityUnitObjectIDType">type</a> [PeriodActivityUnitObjectIDType](/src/target/wwsgen_types.go?s=210253:210384#L3264)
``` go
type PeriodActivityUnitObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="PeriodActivityUnitObjectType">type</a> [PeriodActivityUnitObjectType](/src/target/wwsgen_types.go?s=210386:210617#L3269)
``` go
type PeriodActivityUnitObjectType struct {
    ID         []PeriodActivityUnitObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PeriodPlanDataType">type</a> [PeriodPlanDataType](/src/target/wwsgen_types.go?s=210639:211614#L3275)
``` go
type PeriodPlanDataType struct {
    CompensationElementReference              *CompensationPayEarningObjectType            `xml:"urn:com.workday/bsvc Compensation_Element_Reference"`
    DefaultCompensationPeriodMultiplier       float64                                      `xml:"urn:com.workday/bsvc Default_Compensation_Period_Multiplier,omitempty"`
    UnitofDurationReference                   *CompensationPeriodObjectType                `xml:"urn:com.workday/bsvc Unit_of_Duration_Reference"`
    CompensationBasisReference                *CompensationBasisObjectType                 `xml:"urn:com.workday/bsvc Compensation_Basis_Reference"`
    PeriodPlanProfileReplacementData          []PeriodPlanProfileReplacementDataType       `xml:"urn:com.workday/bsvc Period_Plan_Profile_Replacement_Data,omitempty"`
    CompensationSetupSecuritySegmentReference []CompensationSetupSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Compensation_Setup_Security_Segment_Reference,omitempty"`
}
```
Period Plan Data










## <a name="PeriodPlanProfileReplacementDataType">type</a> [PeriodPlanProfileReplacementDataType](/src/target/wwsgen_types.go?s=211656:212088#L3285)
``` go
type PeriodPlanProfileReplacementDataType struct {
    EligibilityRuleReference            *ConditionRuleObjectType     `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    DefaultCompensationPeriodMultiplier float64                      `xml:"urn:com.workday/bsvc Default_Compensation_Period_Multiplier"`
    CompensationBasisReference          *CompensationBasisObjectType `xml:"urn:com.workday/bsvc Compensation_Basis_Reference"`
}
```
Profile data for Period Salary Plan.










## <a name="PeriodSalaryPlanObjectIDType">type</a> [PeriodSalaryPlanObjectIDType](/src/target/wwsgen_types.go?s=212152:212281#L3292)
``` go
type PeriodSalaryPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="PeriodSalaryPlanObjectType">type</a> [PeriodSalaryPlanObjectType](/src/target/wwsgen_types.go?s=212283:212508#L3297)
``` go
type PeriodSalaryPlanObjectType struct {
    ID         []PeriodSalaryPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PositionElementObjectIDType">type</a> [PositionElementObjectIDType](/src/target/wwsgen_types.go?s=212572:212700#L3303)
``` go
type PositionElementObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="PositionElementObjectType">type</a> [PositionElementObjectType](/src/target/wwsgen_types.go?s=212702:212924#L3308)
``` go
type PositionElementObjectType struct {
    ID         []PositionElementObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PositionFuturePaymentPlanAssignmentDataType">type</a> [PositionFuturePaymentPlanAssignmentDataType](/src/target/wwsgen_types.go?s=213005:213584#L3314)
``` go
type PositionFuturePaymentPlanAssignmentDataType struct {
    EmployeeReference               *WorkerObjectType                     `xml:"urn:com.workday/bsvc Employee_Reference"`
    PositionReference               *PositionElementObjectType            `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
    AssignmentDate                  *time.Time                            `xml:"urn:com.workday/bsvc Assignment_Date,omitempty"`
    FuturePaymentPlanAssignmentData []FuturePaymentPlanAssignmentDataType `xml:"urn:com.workday/bsvc Future_Payment_Plan_Assignment_Data,omitempty"`
}
```
Future payment plan assignments for a position effective for specific date.










### <a name="PositionFuturePaymentPlanAssignmentDataType.MarshalXML">func</a> (\*PositionFuturePaymentPlanAssignmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=213586:213696#L3321)
``` go
func (t *PositionFuturePaymentPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="PositionFuturePaymentPlanAssignmentDataType.UnmarshalXML">func</a> (\*PositionFuturePaymentPlanAssignmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=213983:214095#L3331)
``` go
func (t *PositionFuturePaymentPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="PositionFuturePaymentPlanAssignmentsType">type</a> [PositionFuturePaymentPlanAssignmentsType](/src/target/wwsgen_types.go?s=214447:214674#L3343)
``` go
type PositionFuturePaymentPlanAssignmentsType struct {
    PositionFuturePaymentPlanAssignmentData []PositionFuturePaymentPlanAssignmentDataType `xml:"urn:com.workday/bsvc Position_Future_Payment_Plan_Assignment_Data,omitempty"`
}
```
Groups an Employee's Future Payment Plan Assignments.










## <a name="PositionObjectIDType">type</a> [PositionObjectIDType](/src/target/wwsgen_types.go?s=214738:214859#L3348)
``` go
type PositionObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="PositionObjectType">type</a> [PositionObjectType](/src/target/wwsgen_types.go?s=214861:215062#L3353)
``` go
type PositionObjectType struct {
    ID         []PositionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PotentialObjectIDType">type</a> [PotentialObjectIDType](/src/target/wwsgen_types.go?s=215126:215248#L3359)
``` go
type PotentialObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="PotentialObjectType">type</a> [PotentialObjectType](/src/target/wwsgen_types.go?s=215250:215454#L3364)
``` go
type PotentialObjectType struct {
    ID         []PotentialObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="PreviousSystemCompensationHistoryDataType">type</a> [PreviousSystemCompensationHistoryDataType](/src/target/wwsgen_types.go?s=215571:216282#L3370)
``` go
type PreviousSystemCompensationHistoryDataType struct {
    PreviousSystemCompensationHistoryReference  *CompensationPreviousSystemHistoryObjectType     `xml:"urn:com.workday/bsvc Previous_System_Compensation_History_Reference,omitempty"`
    PreviousSystemCompensationHistoryDetailData *PreviousSystemCompensationHistoryDetailDataType `xml:"urn:com.workday/bsvc Previous_System_Compensation_History_Detail_Data,omitempty"`
    AddOnly                                     bool                                             `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
    Delete                                      bool                                             `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
}
```
Container for defining whether a previous system compensation history entry is being added, updated or deleted.










## <a name="PreviousSystemCompensationHistoryDetailDataType">type</a> [PreviousSystemCompensationHistoryDetailDataType](/src/target/wwsgen_types.go?s=216369:217266#L3378)
``` go
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
```
Container for the detailed data for a previous system compensation history entry.










### <a name="PreviousSystemCompensationHistoryDetailDataType.MarshalXML">func</a> (\*PreviousSystemCompensationHistoryDetailDataType) [MarshalXML](/src/target/wwsgen_types.go?s=217268:217382#L3390)
``` go
func (t *PreviousSystemCompensationHistoryDetailDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="PreviousSystemCompensationHistoryDetailDataType.UnmarshalXML">func</a> (\*PreviousSystemCompensationHistoryDetailDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=217657:217773#L3400)
``` go
func (t *PreviousSystemCompensationHistoryDetailDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="PreviousSystemCompensationHistoryGetDataType">type</a> [PreviousSystemCompensationHistoryGetDataType](/src/target/wwsgen_types.go?s=218128:218483#L3412)
``` go
type PreviousSystemCompensationHistoryGetDataType struct {
    WorkerReference                       *WorkerObjectType                       `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
    PreviousSystemCompensationHistoryData []PreviousSystemCompensationHistoryType `xml:"urn:com.workday/bsvc Previous_System_Compensation_History_Data,omitempty"`
}
```
Container for a worker and the previous system compensation history.










## <a name="PreviousSystemCompensationHistoryResponseDataType">type</a> [PreviousSystemCompensationHistoryResponseDataType](/src/target/wwsgen_types.go?s=218563:218786#L3418)
``` go
type PreviousSystemCompensationHistoryResponseDataType struct {
    PreviousSystemCompensationHistory []PreviousSystemCompensationHistoryGetDataType `xml:"urn:com.workday/bsvc Previous_System_Compensation_History,omitempty"`
}
```
Contains the worker and their compensation history from a previous system.










## <a name="PreviousSystemCompensationHistoryType">type</a> [PreviousSystemCompensationHistoryType](/src/target/wwsgen_types.go?s=218901:219197#L3423)
``` go
type PreviousSystemCompensationHistoryType struct {
    WorkerReference          *WorkerObjectType                           `xml:"urn:com.workday/bsvc Worker_Reference"`
    PreviousSystemJobHistory []PreviousSystemCompensationHistoryDataType `xml:"urn:com.workday/bsvc Previous_System_Job_History"`
}
```
Contains the data for adding, updating or deleting a previous system compensation history entry for a worker.










## <a name="ProcessingFaultType">type</a> [ProcessingFaultType](/src/target/wwsgen_types.go?s=219199:219310#L3428)
``` go
type ProcessingFaultType struct {
    DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
}
```









## <a name="ProfileCompensationScorecardResultDataType">type</a> [ProfileCompensationScorecardResultDataType](/src/target/wwsgen_types.go?s=219396:219679#L3433)
``` go
type ProfileCompensationScorecardResultDataType struct {
    EligibilityRuleReference *ConditionRuleObjectType         `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    ScorecardResultData      []ProfileScorecardResultDataType `xml:"urn:com.workday/bsvc Scorecard_Result_Data"`
}
```
Contains a rule/profile delineated set of goals and goal values for this result.










## <a name="ProfileScorecardResultDataType">type</a> [ProfileScorecardResultDataType](/src/target/wwsgen_types.go?s=219743:219976#L3439)
``` go
type ProfileScorecardResultDataType struct {
    GoalReference *PerformanceCriteriaObjectType `xml:"urn:com.workday/bsvc Goal_Reference"`
    Achievement   float64                        `xml:"urn:com.workday/bsvc Achievement,omitempty"`
}
```
Contains the set of goals and goal values for this result.










## <a name="ProposedAllowancePlanAssignmentContainerDataType">type</a> [ProposedAllowancePlanAssignmentContainerDataType](/src/target/wwsgen_types.go?s=220052:220359#L3445)
``` go
type ProposedAllowancePlanAssignmentContainerDataType struct {
    AllowancePlanSubData []ProposedAllowancePlanAssignmentDataType `xml:"urn:com.workday/bsvc Allowance_Plan_Sub_Data,omitempty"`
    Replace              bool                                      `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}
```
Encapsulating element containing all Allowance Plan Compensation data.










## <a name="ProposedAllowancePlanAssignmentDataType">type</a> [ProposedAllowancePlanAssignmentDataType](/src/target/wwsgen_types.go?s=220435:221805#L3451)
``` go
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
```
Encapsulating element containing all Allowance Plan Compensation data.










### <a name="ProposedAllowancePlanAssignmentDataType.MarshalXML">func</a> (\*ProposedAllowancePlanAssignmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=221807:221913#L3464)
``` go
func (t *ProposedAllowancePlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="ProposedAllowancePlanAssignmentDataType.UnmarshalXML">func</a> (\*ProposedAllowancePlanAssignmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=222531:222639#L3478)
``` go
func (t *ProposedAllowancePlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="ProposedAllowanceUnitPlanAssignmentContainerDataType">type</a> [ProposedAllowanceUnitPlanAssignmentContainerDataType](/src/target/wwsgen_types.go?s=223348:223680#L3494)
``` go
type ProposedAllowanceUnitPlanAssignmentContainerDataType struct {
    UnitAllowancePlanSubData []ProposedAllowanceUnitPlanAssignmentDataType `xml:"urn:com.workday/bsvc Unit_Allowance_Plan_Sub_Data,omitempty"`
    Replace                  bool                                          `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}
```
Encapsulating element containing all Unit Allowance Plan Compensation data.










## <a name="ProposedAllowanceUnitPlanAssignmentDataType">type</a> [ProposedAllowanceUnitPlanAssignmentDataType](/src/target/wwsgen_types.go?s=223761:224762#L3500)
``` go
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
```
Encapsulating element containing all Unit Allowance Plan Compensation data.










### <a name="ProposedAllowanceUnitPlanAssignmentDataType.MarshalXML">func</a> (\*ProposedAllowanceUnitPlanAssignmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=224764:224874#L3511)
``` go
func (t *ProposedAllowanceUnitPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="ProposedAllowanceUnitPlanAssignmentDataType.UnmarshalXML">func</a> (\*ProposedAllowanceUnitPlanAssignmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=225342:225454#L3523)
``` go
func (t *ProposedAllowanceUnitPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="ProposedBasePayPlanAssignmentContainerDataType">type</a> [ProposedBasePayPlanAssignmentContainerDataType](/src/target/wwsgen_types.go?s=226013:226296#L3537)
``` go
type ProposedBasePayPlanAssignmentContainerDataType struct {
    PayPlanSubData []ProposedBasePayPlanAssignmentDataType `xml:"urn:com.workday/bsvc Pay_Plan_Sub_Data,omitempty"`
    Replace        bool                                    `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}
```
Encapsulating element containing all Salary or Hourly Plan Compensation data.










## <a name="ProposedBasePayPlanAssignmentDataType">type</a> [ProposedBasePayPlanAssignmentDataType](/src/target/wwsgen_types.go?s=226379:227202#L3543)
``` go
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
```
Encapsulating element containing all Salary or Hourly Plan Compensation data.










### <a name="ProposedBasePayPlanAssignmentDataType.MarshalXML">func</a> (\*ProposedBasePayPlanAssignmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=227204:227308#L3554)
``` go
func (t *ProposedBasePayPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="ProposedBasePayPlanAssignmentDataType.UnmarshalXML">func</a> (\*ProposedBasePayPlanAssignmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=227735:227841#L3566)
``` go
func (t *ProposedBasePayPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="ProposedBonusPlanAssignmentContainerDataType">type</a> [ProposedBonusPlanAssignmentContainerDataType](/src/target/wwsgen_types.go?s=228348:228631#L3580)
``` go
type ProposedBonusPlanAssignmentContainerDataType struct {
    BonusPlanSubData []ProposedBonusPlanAssignmentDataType `xml:"urn:com.workday/bsvc Bonus_Plan_Sub_Data,omitempty"`
    Replace          bool                                  `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}
```
Encapsulating element containing all Bonus Plan Compensation data.










## <a name="ProposedBonusPlanAssignmentDataType">type</a> [ProposedBonusPlanAssignmentDataType](/src/target/wwsgen_types.go?s=228703:229763#L3586)
``` go
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
```
Encapsulating element containing all Bonus Plan Compensation data.










### <a name="ProposedBonusPlanAssignmentDataType.MarshalXML">func</a> (\*ProposedBonusPlanAssignmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=229765:229867#L3597)
``` go
func (t *ProposedBonusPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="ProposedBonusPlanAssignmentDataType.UnmarshalXML">func</a> (\*ProposedBonusPlanAssignmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=230143:230247#L3607)
``` go
func (t *ProposedBonusPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="ProposedCalculatedPlanAssignmentContainerDataType">type</a> [ProposedCalculatedPlanAssignmentContainerDataType](/src/target/wwsgen_types.go?s=230606:230919#L3619)
``` go
type ProposedCalculatedPlanAssignmentContainerDataType struct {
    CalculatedPlanSubData []ProposedCalculatedPlanAssignmentDataType `xml:"urn:com.workday/bsvc Calculated_Plan_Sub_Data,omitempty"`
    Replace               bool                                       `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}
```
Encapsulating element containing all Calculated Plan Compensation data.










## <a name="ProposedCalculatedPlanAssignmentDataType">type</a> [ProposedCalculatedPlanAssignmentDataType](/src/target/wwsgen_types.go?s=230996:231708#L3625)
``` go
type ProposedCalculatedPlanAssignmentDataType struct {
    CalculatedPlanReference                 *CalculatedPlanObjectType `xml:"urn:com.workday/bsvc Calculated_Plan_Reference,omitempty"`
    ManagebyCompensationBasisOverrideAmount float64                   `xml:"urn:com.workday/bsvc Manage_by_Compensation_Basis_Override_Amount,omitempty"`
    CurrencyReference                       *CurrencyObjectType       `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
    FrequencyReference                      *FrequencyObjectType      `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
    ActualEndDate                           *time.Time                `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
}
```
Encapsulating element containing all Calculated Plan Compensation data.










### <a name="ProposedCalculatedPlanAssignmentDataType.MarshalXML">func</a> (\*ProposedCalculatedPlanAssignmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=231710:231817#L3633)
``` go
func (t *ProposedCalculatedPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="ProposedCalculatedPlanAssignmentDataType.UnmarshalXML">func</a> (\*ProposedCalculatedPlanAssignmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=232098:232207#L3643)
``` go
func (t *ProposedCalculatedPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="ProposedCommissionPlanAssignmentContainerDataType">type</a> [ProposedCommissionPlanAssignmentContainerDataType](/src/target/wwsgen_types.go?s=232571:232884#L3655)
``` go
type ProposedCommissionPlanAssignmentContainerDataType struct {
    CommissionPlanSubData []ProposedCommissionPlanAssignmentDataType `xml:"urn:com.workday/bsvc Commission_Plan_Sub_Data,omitempty"`
    Replace               bool                                       `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}
```
Encapsulating element containing all Commission Plan Compensation data.










## <a name="ProposedCommissionPlanAssignmentDataType">type</a> [ProposedCommissionPlanAssignmentDataType](/src/target/wwsgen_types.go?s=232961:234137#L3661)
``` go
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
```
Encapsulating element containing all Commission Plan Compensation data.










### <a name="ProposedCommissionPlanAssignmentDataType.MarshalXML">func</a> (\*ProposedCommissionPlanAssignmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=234139:234246#L3674)
``` go
func (t *ProposedCommissionPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="ProposedCommissionPlanAssignmentDataType.UnmarshalXML">func</a> (\*ProposedCommissionPlanAssignmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=234527:234636#L3684)
``` go
func (t *ProposedCommissionPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="ProposedMeritPlanAssignmentContainerDataType">type</a> [ProposedMeritPlanAssignmentContainerDataType](/src/target/wwsgen_types.go?s=234995:235278#L3696)
``` go
type ProposedMeritPlanAssignmentContainerDataType struct {
    MeritPlanSubData []ProposedMeritPlanAssignmentDataType `xml:"urn:com.workday/bsvc Merit_Plan_Sub_Data,omitempty"`
    Replace          bool                                  `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}
```
Encapsulating element containing all Merit Plan Compensation data.










## <a name="ProposedMeritPlanAssignmentDataType">type</a> [ProposedMeritPlanAssignmentDataType](/src/target/wwsgen_types.go?s=235350:235933#L3702)
``` go
type ProposedMeritPlanAssignmentDataType struct {
    MeritPlanReference      *MeritPlanObjectType `xml:"urn:com.workday/bsvc Merit_Plan_Reference,omitempty"`
    IndividualTargetAmount  float64              `xml:"urn:com.workday/bsvc Individual_Target_Amount,omitempty"`
    IndividualTargetPercent float64              `xml:"urn:com.workday/bsvc Individual_Target_Percent,omitempty"`
    GuaranteedMinimum       *bool                `xml:"urn:com.workday/bsvc Guaranteed_Minimum,omitempty"`
    ActualEndDate           *time.Time           `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
}
```
Encapsulating element containing all Merit Plan Compensation data.










### <a name="ProposedMeritPlanAssignmentDataType.MarshalXML">func</a> (\*ProposedMeritPlanAssignmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=235935:236037#L3710)
``` go
func (t *ProposedMeritPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="ProposedMeritPlanAssignmentDataType.UnmarshalXML">func</a> (\*ProposedMeritPlanAssignmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=236313:236417#L3720)
``` go
func (t *ProposedMeritPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="ProposedPeriodSalaryPlanAssignmentContainerDataType">type</a> [ProposedPeriodSalaryPlanAssignmentContainerDataType](/src/target/wwsgen_types.go?s=236779:237105#L3732)
``` go
type ProposedPeriodSalaryPlanAssignmentContainerDataType struct {
    PeriodSalaryPlanSubData []ProposedPeriodSalaryPlanAssignmentDataType `xml:"urn:com.workday/bsvc Period_Salary_Plan_Sub_Data,omitempty"`
    Replace                 bool                                         `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}
```
Encapsulating element containing all Period Salary Plan Compensation data.










## <a name="ProposedPeriodSalaryPlanAssignmentDataType">type</a> [ProposedPeriodSalaryPlanAssignmentDataType](/src/target/wwsgen_types.go?s=237185:238203#L3738)
``` go
type ProposedPeriodSalaryPlanAssignmentDataType struct {
    PeriodSalaryPlanReference               *PeriodSalaryPlanObjectType   `xml:"urn:com.workday/bsvc Period_Salary_Plan_Reference,omitempty"`
    CompensationPeriodReference             *CompensationPeriodObjectType `xml:"urn:com.workday/bsvc Compensation_Period_Reference,omitempty"`
    ManagebyCompensationBasisOverrideAmount float64                       `xml:"urn:com.workday/bsvc Manage_by_Compensation_Basis_Override_Amount,omitempty"`
    CurrencyReference                       *CurrencyObjectType           `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
    CompensationPeriodMultiplier            float64                       `xml:"urn:com.workday/bsvc Compensation_Period_Multiplier,omitempty"`
    FrequencyReference                      *FrequencyObjectType          `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
    ActualEndDate                           *time.Time                    `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
}
```
Encapsulating element containing all Period Salary Plan Compensation data.










### <a name="ProposedPeriodSalaryPlanAssignmentDataType.MarshalXML">func</a> (\*ProposedPeriodSalaryPlanAssignmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=238205:238314#L3748)
``` go
func (t *ProposedPeriodSalaryPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="ProposedPeriodSalaryPlanAssignmentDataType.UnmarshalXML">func</a> (\*ProposedPeriodSalaryPlanAssignmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=238597:238708#L3758)
``` go
func (t *ProposedPeriodSalaryPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="ProposedSalaryUnitPlanAssignmentContainerDataType">type</a> [ProposedSalaryUnitPlanAssignmentContainerDataType](/src/target/wwsgen_types.go?s=239075:239389#L3770)
``` go
type ProposedSalaryUnitPlanAssignmentContainerDataType struct {
    UnitSalaryPlanSubData []ProposedSalaryUnitPlanAssignmentDataType `xml:"urn:com.workday/bsvc Unit_Salary_Plan_Sub_Data,omitempty"`
    Replace               bool                                       `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}
```
Encapsulating element containing all Unit Salary Plan Compensation data.










## <a name="ProposedSalaryUnitPlanAssignmentDataType">type</a> [ProposedSalaryUnitPlanAssignmentDataType](/src/target/wwsgen_types.go?s=239467:240175#L3776)
``` go
type ProposedSalaryUnitPlanAssignmentDataType struct {
    UnitSalaryPlanReference *SalaryUnitPlanObjectType `xml:"urn:com.workday/bsvc Unit_Salary_Plan_Reference,omitempty"`
    PerUnitAmount           float64                   `xml:"urn:com.workday/bsvc Per_Unit_Amount,omitempty"`
    CurrencyReference       *CurrencyObjectType       `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
    DefaultUnits            float64                   `xml:"urn:com.workday/bsvc Default_Units,omitempty"`
    FrequencyReference      *FrequencyObjectType      `xml:"urn:com.workday/bsvc Frequency_Reference,omitempty"`
    ActualEndDate           *time.Time                `xml:"urn:com.workday/bsvc Actual_End_Date,omitempty"`
}
```
Encapsulating element containing all Unit Salary Plan Compensation data.










### <a name="ProposedSalaryUnitPlanAssignmentDataType.MarshalXML">func</a> (\*ProposedSalaryUnitPlanAssignmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=240177:240284#L3785)
``` go
func (t *ProposedSalaryUnitPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="ProposedSalaryUnitPlanAssignmentDataType.UnmarshalXML">func</a> (\*ProposedSalaryUnitPlanAssignmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=240565:240674#L3795)
``` go
func (t *ProposedSalaryUnitPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="ProposedStockPlanAssignmentContainerDataType">type</a> [ProposedStockPlanAssignmentContainerDataType](/src/target/wwsgen_types.go?s=241033:241316#L3807)
``` go
type ProposedStockPlanAssignmentContainerDataType struct {
    StockPlanSubData []ProposedStockPlanAssignmentDataType `xml:"urn:com.workday/bsvc Stock_Plan_Sub_Data,omitempty"`
    Replace          bool                                  `xml:"urn:com.workday/bsvc Replace,attr,omitempty"`
}
```
Encapsulating element containing all Stock Plan Compensation data.










## <a name="ProposedStockPlanAssignmentDataType">type</a> [ProposedStockPlanAssignmentDataType](/src/target/wwsgen_types.go?s=241388:242456#L3813)
``` go
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
```
Encapsulating element containing all Stock Plan Compensation data.










### <a name="ProposedStockPlanAssignmentDataType.MarshalXML">func</a> (\*ProposedStockPlanAssignmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=242458:242560#L3824)
``` go
func (t *ProposedStockPlanAssignmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="ProposedStockPlanAssignmentDataType.UnmarshalXML">func</a> (\*ProposedStockPlanAssignmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=242836:242940#L3834)
``` go
func (t *ProposedStockPlanAssignmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="PutBenchmarkJobInput">type</a> [PutBenchmarkJobInput](/src/target/wwsgen_client.go?s=16131:16266#L429)
``` go
type PutBenchmarkJobInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Benchmark_Job_Request"`
    PutBenchmarkJobRequestType
}
```









## <a name="PutBenchmarkJobOutput">type</a> [PutBenchmarkJobOutput](/src/target/wwsgen_client.go?s=16268:16406#L434)
``` go
type PutBenchmarkJobOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Benchmark_Job_Response"`
    PutBenchmarkJobResponseType
}
```









## <a name="PutBenchmarkJobRequestType">type</a> [PutBenchmarkJobRequestType](/src/target/wwsgen_types.go?s=243265:243775#L3846)
``` go
type PutBenchmarkJobRequestType struct {
    BenchmarkJobReference *CompensationBenchmarkDefaultObjectType `xml:"urn:com.workday/bsvc Benchmark_Job_Reference,omitempty"`
    BenchmarkJobData      *BenchmarkJobDataType                   `xml:"urn:com.workday/bsvc Benchmark_Job_Data"`
    AddOnly               bool                                    `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
    Version               string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request element to put Benchmark Job.










## <a name="PutBenchmarkJobResponseType">type</a> [PutBenchmarkJobResponseType](/src/target/wwsgen_types.go?s=243818:244102#L3854)
``` go
type PutBenchmarkJobResponseType struct {
    BenchmarkJobReference *CompensationBenchmarkDefaultObjectType `xml:"urn:com.workday/bsvc Benchmark_Job_Reference,omitempty"`
    Version               string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element to put Benchmark Job










## <a name="PutCompensationEligibilityRuleInput">type</a> [PutCompensationEligibilityRuleInput](/src/target/wwsgen_client.go?s=19016:19197#L517)
``` go
type PutCompensationEligibilityRuleInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Eligibility_Rule_Request"`
    PutCompensationEligibilityRuleRequestType
}
```









## <a name="PutCompensationEligibilityRuleOutput">type</a> [PutCompensationEligibilityRuleOutput](/src/target/wwsgen_client.go?s=19199:19383#L522)
``` go
type PutCompensationEligibilityRuleOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Eligibility_Rule_Response"`
    PutCompensationEligibilityRuleResponseType
}
```









## <a name="PutCompensationEligibilityRuleRequestType">type</a> [PutCompensationEligibilityRuleRequestType](/src/target/wwsgen_types.go?s=244158:244899#L3860)
``` go
type PutCompensationEligibilityRuleRequestType struct {
    CompensationEligibilityRuleReference *ConditionRuleObjectType             `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule_Reference,omitempty"`
    CompensationEligibilityRuleData      *CompensationEligibilityRuleDataType `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule_Data,omitempty"`
    AddOnly                              bool                                 `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
    Delete                               bool                                 `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
    Version                              string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Incoming Put Compensation Eligibility Rule Request










## <a name="PutCompensationEligibilityRuleResponseType">type</a> [PutCompensationEligibilityRuleResponseType](/src/target/wwsgen_types.go?s=244955:245270#L3869)
``` go
type PutCompensationEligibilityRuleResponseType struct {
    CompensationEligibilityRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Compensation_Eligibility_Rule_Reference,omitempty"`
    Version                              string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Put Compensation Eligibility Rule Response element










## <a name="PutCompensationGradeHierarchyInput">type</a> [PutCompensationGradeHierarchyInput](/src/target/wwsgen_client.go?s=25480:25658#L671)
``` go
type PutCompensationGradeHierarchyInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Grade_Hierarchy_Request"`
    PutCompensationGradeHierarchyRequestType
}
```









## <a name="PutCompensationGradeHierarchyOutput">type</a> [PutCompensationGradeHierarchyOutput](/src/target/wwsgen_client.go?s=25660:25841#L676)
``` go
type PutCompensationGradeHierarchyOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Grade_Hierarchy_Response"`
    PutCompensationGradeHierarchyResponseType
}
```









## <a name="PutCompensationGradeHierarchyRequestType">type</a> [PutCompensationGradeHierarchyRequestType](/src/target/wwsgen_types.go?s=245272:245589#L3874)
``` go
type PutCompensationGradeHierarchyRequestType struct {
    CompensationGradeHierarchyData *CompensationGradeHierarchyDataType `xml:"urn:com.workday/bsvc Compensation_Grade_Hierarchy_Data,omitempty"`
    Version                        string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```









## <a name="PutCompensationGradeHierarchyResponseType">type</a> [PutCompensationGradeHierarchyResponseType](/src/target/wwsgen_types.go?s=245591:246014#L3879)
``` go
type PutCompensationGradeHierarchyResponseType struct {
    EffectiveDate           *time.Time                              `xml:"urn:com.workday/bsvc Effective_Date,omitempty"`
    CompensationGradeLevels []PutCompensationGradeLevelResponseType `xml:"urn:com.workday/bsvc Compensation_Grade_Levels,omitempty"`
    Version                 string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```









### <a name="PutCompensationGradeHierarchyResponseType.MarshalXML">func</a> (\*PutCompensationGradeHierarchyResponseType) [MarshalXML](/src/target/wwsgen_types.go?s=246016:246124#L3885)
``` go
func (t *PutCompensationGradeHierarchyResponseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="PutCompensationGradeHierarchyResponseType.UnmarshalXML">func</a> (\*PutCompensationGradeHierarchyResponseType) [UnmarshalXML](/src/target/wwsgen_types.go?s=246405:246515#L3895)
``` go
func (t *PutCompensationGradeHierarchyResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="PutCompensationGradeInput">type</a> [PutCompensationGradeInput](/src/target/wwsgen_client.go?s=10414:10564#L275)
``` go
type PutCompensationGradeInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Grade_Request"`
    PutCompensationGradeRequestType
}
```









## <a name="PutCompensationGradeLevelResponseType">type</a> [PutCompensationGradeLevelResponseType](/src/target/wwsgen_types.go?s=246804:246998#L3906)
``` go
type PutCompensationGradeLevelResponseType struct {
    CompensationGradeLevelReference *CompensationGradeLevelObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Level_Reference,omitempty"`
}
```









## <a name="PutCompensationGradeOutput">type</a> [PutCompensationGradeOutput](/src/target/wwsgen_client.go?s=10566:10719#L280)
``` go
type PutCompensationGradeOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Grade_Response"`
    PutCompensationGradeResponseType
}
```









## <a name="PutCompensationGradeRequestType">type</a> [PutCompensationGradeRequestType](/src/target/wwsgen_types.go?s=247337:247838#L3911)
``` go
type PutCompensationGradeRequestType struct {
    CompensationGradeReference *CompensationGradeObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Reference,omitempty"`
    CompensationGradeData      *CompensationGradeDataType   `xml:"urn:com.workday/bsvc Compensation_Grade_Data"`
    AddOnly                    bool                         `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
    Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
This web service operation is designed to load compensation grades and the associated compensation grade profiles.  A complete set of compensation grade profiles should be passed in with the compensation grade.  Any profiles not found in the set will be deleted, new profiles will be added, and the existing profiles will be updated.










## <a name="PutCompensationGradeResponseType">type</a> [PutCompensationGradeResponseType](/src/target/wwsgen_types.go?s=247902:248184#L3919)
``` go
type PutCompensationGradeResponseType struct {
    CompensationGradeReference *CompensationGradeObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Reference,omitempty"`
    Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Responds with the reference ID for the compensation grade.










## <a name="PutCompensationMatrixInput">type</a> [PutCompensationMatrixInput](/src/target/wwsgen_client.go?s=8675:8828#L231)
``` go
type PutCompensationMatrixInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Matrix_Request"`
    PutCompensationMatrixRequestType
}
```









## <a name="PutCompensationMatrixOutput">type</a> [PutCompensationMatrixOutput](/src/target/wwsgen_client.go?s=8830:8986#L236)
``` go
type PutCompensationMatrixOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Matrix_Response"`
    PutCompensationMatrixResponseType
}
```









## <a name="PutCompensationMatrixRequestType">type</a> [PutCompensationMatrixRequestType](/src/target/wwsgen_types.go?s=248239:248639#L3925)
``` go
type PutCompensationMatrixRequestType struct {
    CompensationMatrixReference *CompensationMatrixObjectType `xml:"urn:com.workday/bsvc Compensation_Matrix_Reference,omitempty"`
    CompensationMatrixData      *CompensationMatrixDataType   `xml:"urn:com.workday/bsvc Compensation_Matrix_Data"`
    Version                     string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
A request object to put Compensation Matrix data.










## <a name="PutCompensationMatrixResponseType">type</a> [PutCompensationMatrixResponseType](/src/target/wwsgen_types.go?s=248716:249008#L3932)
``` go
type PutCompensationMatrixResponseType struct {
    MeritIncreaseMatrixReference *CompensationMatrixObjectType `xml:"urn:com.workday/bsvc Merit_Increase_Matrix_Reference,omitempty"`
    Version                      string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
The response to a request to modify or create Compensation Matrix data.










## <a name="PutCompensationPlanRequestType">type</a> [PutCompensationPlanRequestType](/src/target/wwsgen_types.go?s=249129:249659#L3938)
``` go
type PutCompensationPlanRequestType struct {
    CompensationPlanReference *CompensationAssignablePlanObjectType `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
    CompensationPlanData      *CompensationPlanDataType             `xml:"urn:com.workday/bsvc Compensation_Plan_Data"`
    AddOnly                   bool                                  `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
    Version                   string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request element to put Compensation Plan (Bonus, Merit, Allowance, Future Payment, Commission, and Base Pay Plans).










## <a name="PutCompensationPlanResponseType">type</a> [PutCompensationPlanResponseType](/src/target/wwsgen_types.go?s=249781:250067#L3946)
``` go
type PutCompensationPlanResponseType struct {
    CompensationPlanReference *CompensationAssignablePlanObjectType `xml:"urn:com.workday/bsvc Compensation_Plan_Reference"`
    Version                   string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element to put Compensation Plan (Bonus, Merit, Allowance, Future Payment, Commission, and Base Pay Plans).










## <a name="PutCompensationPlansInput">type</a> [PutCompensationPlansInput](/src/target/wwsgen_client.go?s=11923:12071#L319)
``` go
type PutCompensationPlansInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Plan_Request"`
    PutCompensationPlanRequestType
}
```









## <a name="PutCompensationPlansOutput">type</a> [PutCompensationPlansOutput](/src/target/wwsgen_client.go?s=12073:12224#L324)
``` go
type PutCompensationPlansOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Plan_Response"`
    PutCompensationPlanResponseType
}
```









## <a name="PutCompensationScorecardInput">type</a> [PutCompensationScorecardInput](/src/target/wwsgen_client.go?s=15392:15554#L407)
``` go
type PutCompensationScorecardInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Scorecard_Request"`
    PutCompensationScorecardRequestType
}
```









## <a name="PutCompensationScorecardOutput">type</a> [PutCompensationScorecardOutput](/src/target/wwsgen_client.go?s=15556:15721#L412)
``` go
type PutCompensationScorecardOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Scorecard_Response"`
    PutCompensationScorecardResponseType
}
```









## <a name="PutCompensationScorecardRequestType">type</a> [PutCompensationScorecardRequestType](/src/target/wwsgen_types.go?s=250104:250641#L3952)
``` go
type PutCompensationScorecardRequestType struct {
    CompensationScorecardReference *DefaultScorecardObjectType    `xml:"urn:com.workday/bsvc Compensation_Scorecard_Reference,omitempty"`
    CompensationScorecardData      *CompensationScorecardDataType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Data"`
    AddOnly                        bool                           `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
    Version                        string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Contains scorecard information.










## <a name="PutCompensationScorecardResponseType">type</a> [PutCompensationScorecardResponseType](/src/target/wwsgen_types.go?s=250679:250975#L3960)
``` go
type PutCompensationScorecardResponseType struct {
    CompensationScorecardReference *DefaultScorecardObjectType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Reference,omitempty"`
    Version                        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Compensation Scorecard Response.










## <a name="PutCompensationScorecardResultInput">type</a> [PutCompensationScorecardResultInput](/src/target/wwsgen_client.go?s=21874:22055#L583)
``` go
type PutCompensationScorecardResultInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Scorecard_Result_Request"`
    PutCompensationScorecardResultRequestType
}
```









## <a name="PutCompensationScorecardResultOutput">type</a> [PutCompensationScorecardResultOutput](/src/target/wwsgen_client.go?s=22057:22241#L588)
``` go
type PutCompensationScorecardResultOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Compensation_Scorecard_Result_Response"`
    PutCompensationScorecardResultResponseType
}
```









## <a name="PutCompensationScorecardResultRequestType">type</a> [PutCompensationScorecardResultRequestType](/src/target/wwsgen_types.go?s=251147:251752#L3966)
``` go
type PutCompensationScorecardResultRequestType struct {
    CompensationScorecardResultReference *ScoresetContainerObjectType         `xml:"urn:com.workday/bsvc Compensation_Scorecard_Result_Reference,omitempty"`
    CompensationScorecardResultData      *CompensationScorecardResultDataType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Result_Data"`
    AddOnly                              bool                                 `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
    Version                              string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Contains the scorecard result request information.  This includes the detailed scorecard result data and an optional reference an existing scorecard result to update.










## <a name="PutCompensationScorecardResultResponseType">type</a> [PutCompensationScorecardResultResponseType](/src/target/wwsgen_types.go?s=251831:252154#L3974)
``` go
type PutCompensationScorecardResultResponseType struct {
    CompensationScorecardResultReference *ScoresetContainerObjectType `xml:"urn:com.workday/bsvc Compensation_Scorecard_Result_Reference,omitempty"`
    Version                              string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Contains the scorecard result identifier of the added or modified result.










## <a name="PutEligibleEarningsInput">type</a> [PutEligibleEarningsInput](/src/target/wwsgen_client.go?s=6379:6526#L165)
``` go
type PutEligibleEarningsInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Eligible_Earnings_Request"`
    PutEligibleEarningsRequestType
}
```









## <a name="PutEligibleEarningsOutput">type</a> [PutEligibleEarningsOutput](/src/target/wwsgen_client.go?s=6528:6678#L170)
``` go
type PutEligibleEarningsOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Eligible_Earnings_Response"`
    PutEligibleEarningsResponseType
}
```









## <a name="PutEligibleEarningsRequestType">type</a> [PutEligibleEarningsRequestType](/src/target/wwsgen_types.go?s=252621:253143#L3980)
``` go
type PutEligibleEarningsRequestType struct {
    EligibleEarningsReference *EligibleEarningsOverrideObjectType `xml:"urn:com.workday/bsvc Eligible_Earnings_Reference,omitempty"`
    EligibleEarningsData      *EligibleEarningsDataType           `xml:"urn:com.workday/bsvc Eligible_Earnings_Data"`
    AddOnly                   bool                                `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
    Version                   string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
This web service operation is designed to load eligible earnings overrides for an employee for a specified eligible earnings override period.  If bonus plans are specified, then the eligible earnings override will apply only to those specific bonus plans.  Alternatively, selecting "Apply to All Bonus Plans" option will enable the eligible earnings override to be applied to all current and future bonus plans assigned to the employee for the specified period.










## <a name="PutEligibleEarningsResponseType">type</a> [PutEligibleEarningsResponseType](/src/target/wwsgen_types.go?s=253205:253497#L3988)
``` go
type PutEligibleEarningsResponseType struct {
    EligibleEarningsReference *EligibleEarningsOverrideObjectType `xml:"urn:com.workday/bsvc Eligible_Earnings_Reference,omitempty"`
    Version                   string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Responds with the reference ID for the Eligible Earnings










## <a name="PutEnhancedSeveranceMatrixInput">type</a> [PutEnhancedSeveranceMatrixInput](/src/target/wwsgen_client.go?s=37305:37474#L957)
``` go
type PutEnhancedSeveranceMatrixInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Enhanced_Severance_Matrix_Request"`
    PutEnhancedSeveranceMatrixRequestType
}
```









## <a name="PutEnhancedSeveranceMatrixOutput">type</a> [PutEnhancedSeveranceMatrixOutput](/src/target/wwsgen_client.go?s=37476:37648#L962)
``` go
type PutEnhancedSeveranceMatrixOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Enhanced_Severance_Matrix_Response"`
    PutEnhancedSeveranceMatrixResponseType
}
```









## <a name="PutEnhancedSeveranceMatrixRequestType">type</a> [PutEnhancedSeveranceMatrixRequestType](/src/target/wwsgen_types.go?s=253577:254024#L3994)
``` go
type PutEnhancedSeveranceMatrixRequestType struct {
    EnhancedSeveranceMatrixReference *SeveranceMatrixabstractObjectType `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix_Reference,omitempty"`
    EnhancedSeveranceMatrixData      *EnhancedSeveranceMatrixDataType   `xml:"urn:com.workday/bsvc Enhanced_Severance_Matrix_Data"`
    Version                          string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
The element containing the Put Enhanced Severance Matrix web service sent.










## <a name="PutEnhancedSeveranceMatrixResponseType">type</a> [PutEnhancedSeveranceMatrixResponseType](/src/target/wwsgen_types.go?s=254093:254414#L4001)
``` go
type PutEnhancedSeveranceMatrixResponseType struct {
    SeveranceMatrixabstractReference *SeveranceMatrixabstractObjectType `xml:"urn:com.workday/bsvc Severance_Matrix__abstract__Reference,omitempty"`
    Version                          string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
The response for the Put Enhanced Severance Matrix web service.










## <a name="PutFuturePaymentPlanAssignmentInput">type</a> [PutFuturePaymentPlanAssignmentInput](/src/target/wwsgen_client.go?s=13780:13962#L363)
``` go
type PutFuturePaymentPlanAssignmentInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Future_Payment_Plan_Assignment_Request"`
    PutFuturePaymentPlanAssignmentRequestType
}
```









## <a name="PutFuturePaymentPlanAssignmentOutput">type</a> [PutFuturePaymentPlanAssignmentOutput](/src/target/wwsgen_client.go?s=13964:14149#L368)
``` go
type PutFuturePaymentPlanAssignmentOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Future_Payment_Plan_Assignment_Response"`
    PutFuturePaymentPlanAssignmentResponseType
}
```









## <a name="PutFuturePaymentPlanAssignmentRequestType">type</a> [PutFuturePaymentPlanAssignmentRequestType](/src/target/wwsgen_types.go?s=254542:254897#L4007)
``` go
type PutFuturePaymentPlanAssignmentRequestType struct {
    PositionFuturePaymentPlanAssignmentData *PositionFuturePaymentPlanAssignmentDataType `xml:"urn:com.workday/bsvc Position_Future_Payment_Plan_Assignment_Data"`
    Version                                 string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
This web service assigns a future payment plan to an employee for consumption in the future payment true up bonus process.










## <a name="PutFuturePaymentPlanAssignmentResponseType">type</a> [PutFuturePaymentPlanAssignmentResponseType](/src/target/wwsgen_types.go?s=254960:255397#L4013)
``` go
type PutFuturePaymentPlanAssignmentResponseType struct {
    EmployeeReference *WorkerObjectType   `xml:"urn:com.workday/bsvc Employee_Reference,omitempty"`
    PositionReference *PositionObjectType `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
    AssignmentDate    *time.Time          `xml:"urn:com.workday/bsvc Assignment_Date,omitempty"`
    Version           string              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response to a put future payment plan assignment request.










### <a name="PutFuturePaymentPlanAssignmentResponseType.MarshalXML">func</a> (\*PutFuturePaymentPlanAssignmentResponseType) [MarshalXML](/src/target/wwsgen_types.go?s=255399:255508#L4020)
``` go
func (t *PutFuturePaymentPlanAssignmentResponseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="PutFuturePaymentPlanAssignmentResponseType.UnmarshalXML">func</a> (\*PutFuturePaymentPlanAssignmentResponseType) [UnmarshalXML](/src/target/wwsgen_types.go?s=255794:255905#L4030)
``` go
func (t *PutFuturePaymentPlanAssignmentResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="PutImportProcessResponseType">type</a> [PutImportProcessResponseType](/src/target/wwsgen_types.go?s=256238:256676#L4042)
``` go
type PutImportProcessResponseType struct {
    ImportProcessReference  *WebServiceBackgroundProcessRuntimeObjectType `xml:"urn:com.workday/bsvc Import_Process_Reference,omitempty"`
    HeaderInstanceReference *InstanceObjectType                           `xml:"urn:com.workday/bsvc Header_Instance_Reference,omitempty"`
    Version                 string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Put Import Process Response element










## <a name="PutOneTimePaymentPlanConfigurableCategoryInput">type</a> [PutOneTimePaymentPlanConfigurableCategoryInput](/src/target/wwsgen_client.go?s=34668:34885#L891)
``` go
type PutOneTimePaymentPlanConfigurableCategoryInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_One-Time_Payment_Plan_Configurable_Category_Request"`
    PutOneTimePaymentPlanConfigurableCategoryRequestType
}
```









## <a name="PutOneTimePaymentPlanConfigurableCategoryOutput">type</a> [PutOneTimePaymentPlanConfigurableCategoryOutput](/src/target/wwsgen_client.go?s=34887:35107#L896)
``` go
type PutOneTimePaymentPlanConfigurableCategoryOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_One-Time_Payment_Plan_Configurable_Category_Response"`
    PutOneTimePaymentPlanConfigurableCategoryResponseType
}
```









## <a name="PutOneTimePaymentPlanConfigurableCategoryRequestType">type</a> [PutOneTimePaymentPlanConfigurableCategoryRequestType](/src/target/wwsgen_types.go?s=256706:257446#L4049)
``` go
type PutOneTimePaymentPlanConfigurableCategoryRequestType struct {
    OneTimePaymentPlanConfigurableCategoryReference *OneTimePaymentPlanConfigurableCategoryObjectType `xml:"urn:com.workday/bsvc One-Time_Payment_Plan_Configurable_Category_Reference,omitempty"`
    OneTimePaymentPlanConfigurableCategoryData      *OneTimePaymentPlanConfigurableCategoryDataType   `xml:"urn:com.workday/bsvc One-Time_Payment_Plan_Configurable_Category_Data"`
    AddOnly                                         bool                                              `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
    Version                                         string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Incoming request element










## <a name="PutOneTimePaymentPlanConfigurableCategoryResponseType">type</a> [PutOneTimePaymentPlanConfigurableCategoryResponseType](/src/target/wwsgen_types.go?s=257468:257880#L4057)
``` go
type PutOneTimePaymentPlanConfigurableCategoryResponseType struct {
    OneTimePaymentPlanConfigurableCategoryReference *OneTimePaymentPlanConfigurableCategoryObjectType `xml:"urn:com.workday/bsvc One-Time_Payment_Plan_Configurable_Category_Reference,omitempty"`
    Version                                         string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response Element










## <a name="PutPeriodActivityRateMatrixInput">type</a> [PutPeriodActivityRateMatrixInput](/src/target/wwsgen_client.go?s=29094:29267#L759)
``` go
type PutPeriodActivityRateMatrixInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Period_Activity_Rate_Matrix_Request"`
    PutPeriodActivityRateMatrixRequestType
}
```









## <a name="PutPeriodActivityRateMatrixOutput">type</a> [PutPeriodActivityRateMatrixOutput](/src/target/wwsgen_client.go?s=29269:29445#L764)
``` go
type PutPeriodActivityRateMatrixOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Period_Activity_Rate_Matrix_Response"`
    PutPeriodActivityRateMatrixResponseType
}
```









## <a name="PutPeriodActivityRateMatrixRequestType">type</a> [PutPeriodActivityRateMatrixRequestType](/src/target/wwsgen_types.go?s=257943:258525#L4063)
``` go
type PutPeriodActivityRateMatrixRequestType struct {
    PeriodActivityRateMatrixReference *PeriodActivityRateMatrixObjectType `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Reference,omitempty"`
    PeriodActivityRateMatrixData      *PeriodActivityRateMatrixDataType   `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Data"`
    AddOnly                           bool                                `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
    Version                           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request to create or update a Period Activity Rate Matrix










## <a name="PutPeriodActivityRateMatrixResponseType">type</a> [PutPeriodActivityRateMatrixResponseType](/src/target/wwsgen_types.go?s=258571:258897#L4071)
``` go
type PutPeriodActivityRateMatrixResponseType struct {
    PeriodActivityRateMatrixReference *PeriodActivityRateMatrixObjectType `xml:"urn:com.workday/bsvc Period_Activity_Rate_Matrix_Reference,omitempty"`
    Version                           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Put Period Activity Rate Matrix response










## <a name="PutPeriodActivityTaskInput">type</a> [PutPeriodActivityTaskInput](/src/target/wwsgen_client.go?s=27225:27379#L715)
``` go
type PutPeriodActivityTaskInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Period_Activity_Task_Request"`
    PutPeriodActivityTaskRequestType
}
```









## <a name="PutPeriodActivityTaskOutput">type</a> [PutPeriodActivityTaskOutput](/src/target/wwsgen_client.go?s=27381:27538#L720)
``` go
type PutPeriodActivityTaskOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Period_Activity_Task_Response"`
    PutPeriodActivityTaskResponseType
}
```









## <a name="PutPeriodActivityTaskRequestType">type</a> [PutPeriodActivityTaskRequestType](/src/target/wwsgen_types.go?s=258972:259486#L4077)
``` go
type PutPeriodActivityTaskRequestType struct {
    PeriodActivityTaskReference *PeriodActivityTaskObjectType `xml:"urn:com.workday/bsvc Period_Activity_Task_Reference,omitempty"`
    PeriodActivityTaskData      *PeriodActivityTaskDataType   `xml:"urn:com.workday/bsvc Period_Activity_Task_Data"`
    AddOnly                     bool                          `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
    Version                     string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
This web service operation is designed to load period activity tasks.










## <a name="PutPeriodActivityTaskResponseType">type</a> [PutPeriodActivityTaskResponseType](/src/target/wwsgen_types.go?s=259552:259841#L4085)
``` go
type PutPeriodActivityTaskResponseType struct {
    PeriodActivityTaskReference *PeriodActivityTaskObjectType `xml:"urn:com.workday/bsvc Period_Activity_Task_Reference,omitempty"`
    Version                     string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Responds with the reference ID for the period activity task.










## <a name="PutPreviousSystemCompensationHistoryInput">type</a> [PutPreviousSystemCompensationHistoryInput](/src/target/wwsgen_client.go?s=1032:1232#L33)
``` go
type PutPreviousSystemCompensationHistoryInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Previous_System_Compensation_History_Request"`
    PutPreviousSystemCompensationHistoryRequestType
}
```









## <a name="PutPreviousSystemCompensationHistoryOutput">type</a> [PutPreviousSystemCompensationHistoryOutput](/src/target/wwsgen_client.go?s=1234:1437#L38)
``` go
type PutPreviousSystemCompensationHistoryOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Previous_System_Compensation_History_Response"`
    PutPreviousSystemCompensationHistoryResponseType
}
```









## <a name="PutPreviousSystemCompensationHistoryRequestType">type</a> [PutPreviousSystemCompensationHistoryRequestType](/src/target/wwsgen_types.go?s=259956:260298#L4091)
``` go
type PutPreviousSystemCompensationHistoryRequestType struct {
    PreviousSystemCompensationHistoryData *PreviousSystemCompensationHistoryType `xml:"urn:com.workday/bsvc Previous_System_Compensation_History_Data"`
    Version                               string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Contains the data for adding, updating or deleting a previous system compensation history entry for a worker.










## <a name="PutPreviousSystemCompensationHistoryResponseType">type</a> [PutPreviousSystemCompensationHistoryResponseType](/src/target/wwsgen_types.go?s=260378:260620#L4097)
``` go
type PutPreviousSystemCompensationHistoryResponseType struct {
    WorkerReference *WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
    Version         string            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response for Put Previous System Compensation History. Returns the worker.










## <a name="PutStockParticipationRateTableInput">type</a> [PutStockParticipationRateTableInput](/src/target/wwsgen_client.go?s=31839:32021#L825)
``` go
type PutStockParticipationRateTableInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Stock_Participation_Rate_Table_Request"`
    PutStockParticipationRateTableRequestType
}
```









## <a name="PutStockParticipationRateTableOutput">type</a> [PutStockParticipationRateTableOutput](/src/target/wwsgen_client.go?s=32023:32208#L830)
``` go
type PutStockParticipationRateTableOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Put_Stock_Participation_Rate_Table_Response"`
    PutStockParticipationRateTableResponseType
}
```









## <a name="PutStockParticipationRateTableRequestType">type</a> [PutStockParticipationRateTableRequestType](/src/target/wwsgen_types.go?s=260680:261303#L4103)
``` go
type PutStockParticipationRateTableRequestType struct {
    StockParticipationRateTableReference *StockParticipationRateTableObjectType `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_Reference,omitempty"`
    StockParticipationRateTableData      *StockParticipationRateTableData20Type `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_Data,omitempty"`
    Delete                               bool                                   `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
    Version                              string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request element to put Stock Participation Rate Table.










## <a name="PutStockParticipationRateTableResponseType">type</a> [PutStockParticipationRateTableResponseType](/src/target/wwsgen_types.go?s=261352:261696#L4111)
``` go
type PutStockParticipationRateTableResponseType struct {
    StockParticipationRateTableReference *StockParticipationRateTableObjectType `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_Reference,omitempty"`
    Version                              string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Put Stock Participation Rate Table Response










## <a name="RelationalOperatorObjectIDType">type</a> [RelationalOperatorObjectIDType](/src/target/wwsgen_types.go?s=261760:261891#L4117)
``` go
type RelationalOperatorObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="RelationalOperatorObjectType">type</a> [RelationalOperatorObjectType](/src/target/wwsgen_types.go?s=261893:262124#L4122)
``` go
type RelationalOperatorObjectType struct {
    ID         []RelationalOperatorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="RemoveCompensationPlanAssignmentDataType">type</a> [RemoveCompensationPlanAssignmentDataType](/src/target/wwsgen_types.go?s=262194:262373#L4128)
``` go
type RemoveCompensationPlanAssignmentDataType struct {
    CompensationPlanReference []CompensationPlanObjectType `xml:"urn:com.workday/bsvc Compensation_Plan_Reference,omitempty"`
}
```
Specific compensation plan assigned to the worker to be removed.










## <a name="RequestBonusPaymentDataType">type</a> [RequestBonusPaymentDataType](/src/target/wwsgen_types.go?s=262442:263626#L4133)
``` go
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
```
Wrapper element for the Request Bonus Payment business process.










### <a name="RequestBonusPaymentDataType.MarshalXML">func</a> (\*RequestBonusPaymentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=263628:263722#L4144)
``` go
func (t *RequestBonusPaymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="RequestBonusPaymentDataType.UnmarshalXML">func</a> (\*RequestBonusPaymentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=264200:264296#L4156)
``` go
func (t *RequestBonusPaymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="RequestBonusPaymentInput">type</a> [RequestBonusPaymentInput](/src/target/wwsgen_client.go?s=2892:3039#L77)
``` go
type RequestBonusPaymentInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Request_Bonus_Payment_Request"`
    RequestBonusPaymentRequestType
}
```









## <a name="RequestBonusPaymentOutput">type</a> [RequestBonusPaymentOutput](/src/target/wwsgen_client.go?s=3041:3191#L82)
``` go
type RequestBonusPaymentOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Request_Bonus_Payment_Response"`
    RequestBonusPaymentResponseType
}
```









## <a name="RequestBonusPaymentRequestType">type</a> [RequestBonusPaymentRequestType](/src/target/wwsgen_types.go?s=264930:265325#L4170)
``` go
type RequestBonusPaymentRequestType struct {
    BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
    RequestBonusPaymentData   *RequestBonusPaymentDataType   `xml:"urn:com.workday/bsvc Request_Bonus_Payment_Data"`
    Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
This web service operation loads the approved bonus for an employee assigned to a bonus plan using the Request Bonus Payment business process.










## <a name="RequestBonusPaymentResponseType">type</a> [RequestBonusPaymentResponseType](/src/target/wwsgen_types.go?s=265401:265681#L4177)
``` go
type RequestBonusPaymentResponseType struct {
    BonusPaymentEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Bonus_Payment_Event_Reference,omitempty"`
    Version                    string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element containing the reference to the Bonus Payment Request










## <a name="RequestCompensationChangeDataType">type</a> [RequestCompensationChangeDataType](/src/target/wwsgen_types.go?s=265756:266770#L4183)
``` go
type RequestCompensationChangeDataType struct {
    EmployeeReference                 *EmployeeObjectType                        `xml:"urn:com.workday/bsvc Employee_Reference"`
    PositionReference                 *PositionElementObjectType                 `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
    CompensationChangeDate            time.Time                                  `xml:"urn:com.workday/bsvc Compensation_Change_Date"`
    CompensationChangeOnNextPayPeriod bool                                       `xml:"urn:com.workday/bsvc Compensation_Change_On_Next_Pay_Period"`
    EmployeeVisibilityDate            *time.Time                                 `xml:"urn:com.workday/bsvc Employee_Visibility_Date,omitempty"`
    CompensationChangeData            *CompensationChangeDataType                `xml:"urn:com.workday/bsvc Compensation_Change_Data"`
    CheckPositionBudgetSubProcess     *CheckPositionBudgetSubBusinessProcessType `xml:"urn:com.workday/bsvc Check_Position_Budget_Sub_Process,omitempty"`
}
```
Wrapper Element for the Request Compensation Change business process.










### <a name="RequestCompensationChangeDataType.MarshalXML">func</a> (\*RequestCompensationChangeDataType) [MarshalXML](/src/target/wwsgen_types.go?s=266772:266872#L4193)
``` go
func (t *RequestCompensationChangeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="RequestCompensationChangeDataType.UnmarshalXML">func</a> (\*RequestCompensationChangeDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=267348:267450#L4205)
``` go
func (t *RequestCompensationChangeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="RequestCompensationChangeInput">type</a> [RequestCompensationChangeInput](/src/target/wwsgen_client.go?s=5448:5613#L143)
``` go
type RequestCompensationChangeInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Request_Compensation_Change_Request"`
    RequestCompensationChangeRequestType
}
```









## <a name="RequestCompensationChangeOutput">type</a> [RequestCompensationChangeOutput](/src/target/wwsgen_client.go?s=5615:5783#L148)
``` go
type RequestCompensationChangeOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Request_Compensation_Change_Response"`
    RequestCompensationChangeResponseType
}
```









## <a name="RequestCompensationChangeRequestHVType">type</a> [RequestCompensationChangeRequestHVType](/src/target/wwsgen_types.go?s=268101:268416#L4219)
``` go
type RequestCompensationChangeRequestHVType struct {
    BusinessProcessParameters     *BusinessProcessParametersType     `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
    RequestCompensationChangeData *RequestCompensationChangeDataType `xml:"urn:com.workday/bsvc Request_Compensation_Change_Data"`
}
```
Web service operation to request a compensation change for employee(s). This operation is asynchronous and specifically designed to support very large documents.










## <a name="RequestCompensationChangeRequestType">type</a> [RequestCompensationChangeRequestType](/src/target/wwsgen_types.go?s=268493:268924#L4225)
``` go
type RequestCompensationChangeRequestType struct {
    BusinessProcessParameters     *BusinessProcessParametersType     `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
    RequestCompensationChangeData *RequestCompensationChangeDataType `xml:"urn:com.workday/bsvc Request_Compensation_Change_Data"`
    Version                       string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Web service operation to request a compensation change for an employee.










## <a name="RequestCompensationChangeResponseType">type</a> [RequestCompensationChangeResponseType](/src/target/wwsgen_types.go?s=269000:269326#L4232)
``` go
type RequestCompensationChangeResponseType struct {
    RequestCompensationChangeEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Request_Compensation_Change_Event_Reference,omitempty"`
    Version                                 string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element for the Request Compensation Change business process.










## <a name="RequestEmployeeMeritAdjustmentDataType">type</a> [RequestEmployeeMeritAdjustmentDataType](/src/target/wwsgen_types.go?s=269407:270256#L4238)
``` go
type RequestEmployeeMeritAdjustmentDataType struct {
    EmployeeReference              *EmployeeObjectType                       `xml:"urn:com.workday/bsvc Employee_Reference"`
    PositionReference              *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
    MeritIncreaseEffectiveDate     time.Time                                 `xml:"urn:com.workday/bsvc Merit_Increase_Effective_Date"`
    MeritPlanReference             *MeritPercentPlanObjectType               `xml:"urn:com.workday/bsvc Merit_Plan_Reference,omitempty"`
    MeritReasonReference           *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Merit_Reason_Reference,omitempty"`
    EmployeeMeritAdjustmentSubData []EmployeeMeritAdjustmentDataType         `xml:"urn:com.workday/bsvc Employee_Merit_Adjustment_Sub_Data"`
}
```
Wrapper Element for the Request Employee Merit Adjustment business process.










### <a name="RequestEmployeeMeritAdjustmentDataType.MarshalXML">func</a> (\*RequestEmployeeMeritAdjustmentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=270258:270363#L4247)
``` go
func (t *RequestEmployeeMeritAdjustmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="RequestEmployeeMeritAdjustmentDataType.UnmarshalXML">func</a> (\*RequestEmployeeMeritAdjustmentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=270686:270793#L4257)
``` go
func (t *RequestEmployeeMeritAdjustmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="RequestEmployeeMeritAdjustmentInput">type</a> [RequestEmployeeMeritAdjustmentInput](/src/target/wwsgen_client.go?s=3749:3930#L99)
``` go
type RequestEmployeeMeritAdjustmentInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Request_Employee_Merit_Adjustment_Request"`
    RequestEmployeeMeritAdjustmentRequestType
}
```









## <a name="RequestEmployeeMeritAdjustmentOutput">type</a> [RequestEmployeeMeritAdjustmentOutput](/src/target/wwsgen_client.go?s=3932:4116#L104)
``` go
type RequestEmployeeMeritAdjustmentOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Request_Employee_Merit_Adjustment_Response"`
    RequestEmployeeMeritAdjustmentResponseType
}
```









## <a name="RequestEmployeeMeritAdjustmentRequestType">type</a> [RequestEmployeeMeritAdjustmentRequestType](/src/target/wwsgen_types.go?s=271196:271639#L4269)
``` go
type RequestEmployeeMeritAdjustmentRequestType struct {
    BusinessProcessParameters   *BusinessProcessParametersType          `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
    EmployeeMeritAdjustmentData *RequestEmployeeMeritAdjustmentDataType `xml:"urn:com.workday/bsvc Employee_Merit_Adjustment_Data"`
    Version                     string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Web service operation to request a merit adjustment for an employee.










## <a name="RequestEmployeeMeritAdjustmentResponseType">type</a> [RequestEmployeeMeritAdjustmentResponseType](/src/target/wwsgen_types.go?s=271721:272046#L4276)
``` go
type RequestEmployeeMeritAdjustmentResponseType struct {
    EmployeeMeritAdjustmentEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Employee_Merit_Adjustment_Event_Reference,omitempty"`
    Version                               string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element for the Request Employee Merit Adjustment business process.










## <a name="RequestOneTimePaymentDataType">type</a> [RequestOneTimePaymentDataType](/src/target/wwsgen_types.go?s=272118:272884#L4282)
``` go
type RequestOneTimePaymentDataType struct {
    EmployeeReference      *EmployeeObjectType                       `xml:"urn:com.workday/bsvc Employee_Reference"`
    PositionReference      *PositionElementObjectType                `xml:"urn:com.workday/bsvc Position_Reference,omitempty"`
    EffectiveDate          time.Time                                 `xml:"urn:com.workday/bsvc Effective_Date"`
    EmployeeVisibilityDate *time.Time                                `xml:"urn:com.workday/bsvc Employee_Visibility_Date,omitempty"`
    ReasonReference        *EventClassificationSubcategoryObjectType `xml:"urn:com.workday/bsvc Reason_Reference,omitempty"`
    OneTimePaymentSubData  []OneTimePaymentDataType                  `xml:"urn:com.workday/bsvc One-Time_Payment_Sub_Data"`
}
```
Wrapper element for the Request One-Time Payment business process.










### <a name="RequestOneTimePaymentDataType.MarshalXML">func</a> (\*RequestOneTimePaymentDataType) [MarshalXML](/src/target/wwsgen_types.go?s=272886:272982#L4291)
``` go
func (t *RequestOneTimePaymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="RequestOneTimePaymentDataType.UnmarshalXML">func</a> (\*RequestOneTimePaymentDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=273426:273524#L4303)
``` go
func (t *RequestOneTimePaymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="RequestOneTimePaymentInput">type</a> [RequestOneTimePaymentInput](/src/target/wwsgen_client.go?s=4611:4765#L121)
``` go
type RequestOneTimePaymentInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Request_One-Time_Payment_Request"`
    RequestOneTimePaymentRequestType
}
```









## <a name="RequestOneTimePaymentOutput">type</a> [RequestOneTimePaymentOutput](/src/target/wwsgen_client.go?s=4767:4924#L126)
``` go
type RequestOneTimePaymentOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Request_One-Time_Payment_Response"`
    RequestOneTimePaymentResponseType
}
```









## <a name="RequestOneTimePaymentRequestType">type</a> [RequestOneTimePaymentRequestType](/src/target/wwsgen_types.go?s=274136:274528#L4317)
``` go
type RequestOneTimePaymentRequestType struct {
    BusinessProcessParameters *BusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
    OneTimePaymentData        *RequestOneTimePaymentDataType `xml:"urn:com.workday/bsvc One-Time_Payment_Data"`
    Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
This web service operation is designed to pay an employee a one-time payment such as a signing bonus using the Request One-Time Payment  business process.










## <a name="RequestOneTimePaymentResponseType">type</a> [RequestOneTimePaymentResponseType](/src/target/wwsgen_types.go?s=274601:274890#L4324)
``` go
type RequestOneTimePaymentResponseType struct {
    OneTimePaymentEventReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc One-Time_Payment_Event_Reference,omitempty"`
    Version                      string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element containing the reference for the one time payment.










## <a name="RequestRequisitionCompensationChangeInput">type</a> [RequestRequisitionCompensationChangeInput](/src/target/wwsgen_client.go?s=35664:35863#L913)
``` go
type RequestRequisitionCompensationChangeInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Request_Requisition_Compensation_Change_Request"`
    RequestRequisitionCompensationChangeRequestType
}
```









## <a name="RequestRequisitionCompensationChangeOutput">type</a> [RequestRequisitionCompensationChangeOutput](/src/target/wwsgen_client.go?s=35865:36067#L918)
``` go
type RequestRequisitionCompensationChangeOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Request_Requisition_Compensation_Change_Response"`
    RequestRequisitionCompensationChangeResponseType
}
```









## <a name="RequestRequisitionCompensationChangeRequestType">type</a> [RequestRequisitionCompensationChangeRequestType](/src/target/wwsgen_types.go?s=274984:275454#L4330)
``` go
type RequestRequisitionCompensationChangeRequestType struct {
    BusinessProcessParameters         *BusinessProcessParametersType         `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
    RequisitionCompensationChangeData *RequisitionCompensationChangeDataType `xml:"urn:com.workday/bsvc Requisition_Compensation_Change_Data"`
    Version                           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Web service operation to request a requisition compensation change for a job requisition










## <a name="RequestRequisitionCompensationChangeResponseType">type</a> [RequestRequisitionCompensationChangeResponseType](/src/target/wwsgen_types.go?s=275542:275892#L4337)
``` go
type RequestRequisitionCompensationChangeResponseType struct {
    CompensationEventforJobRequisitionReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Compensation_Event_for_Job_Requisition_Reference,omitempty"`
    Version                                     string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element for the Request Requisition Compensation Change business process.










## <a name="RequisitionCompensationChangeDataType">type</a> [RequisitionCompensationChangeDataType](/src/target/wwsgen_types.go?s=275971:276378#L4343)
``` go
type RequisitionCompensationChangeDataType struct {
    JobRequisitionReference     *JobRequisitionObjectType        `xml:"urn:com.workday/bsvc Job_Requisition_Reference"`
    CompensationChangeDate      time.Time                        `xml:"urn:com.workday/bsvc Compensation_Change_Date"`
    RequisitionCompensationData *RequisitionCompensationDataType `xml:"urn:com.workday/bsvc Requisition_Compensation_Data"`
}
```
Element for the Request Requisition Compensation Change business process.










### <a name="RequisitionCompensationChangeDataType.MarshalXML">func</a> (\*RequisitionCompensationChangeDataType) [MarshalXML](/src/target/wwsgen_types.go?s=276380:276484#L4349)
``` go
func (t *RequisitionCompensationChangeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="RequisitionCompensationChangeDataType.UnmarshalXML">func</a> (\*RequisitionCompensationChangeDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=276789:276895#L4359)
``` go
func (t *RequisitionCompensationChangeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="RequisitionCompensationDataType">type</a> [RequisitionCompensationDataType](/src/target/wwsgen_types.go?s=277280:278792#L4371)
``` go
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
```
Data element that contains the requisition compensation information.










## <a name="ResponseFilterType">type</a> [ResponseFilterType](/src/target/wwsgen_types.go?s=278922:279285#L4385)
``` go
type ResponseFilterType struct {
    AsOfEffectiveDate *time.Time `xml:"urn:com.workday/bsvc As_Of_Effective_Date,omitempty"`
    AsOfEntryDateTime *time.Time `xml:"urn:com.workday/bsvc As_Of_Entry_DateTime,omitempty"`
    Page              float64    `xml:"urn:com.workday/bsvc Page,omitempty"`
    Count             float64    `xml:"urn:com.workday/bsvc Count,omitempty"`
}
```
Parameters that let you filter the data returned in the response. You can filter returned data by dates and page attributes.










### <a name="ResponseFilterType.MarshalXML">func</a> (\*ResponseFilterType) [MarshalXML](/src/target/wwsgen_types.go?s=279287:279372#L4392)
``` go
func (t *ResponseFilterType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="ResponseFilterType.UnmarshalXML">func</a> (\*ResponseFilterType) [UnmarshalXML](/src/target/wwsgen_types.go?s=279816:279903#L4404)
``` go
func (t *ResponseFilterType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="ResponseResultsType">type</a> [ResponseResultsType](/src/target/wwsgen_types.go?s=280551:280874#L4418)
``` go
type ResponseResultsType struct {
    TotalResults float64 `xml:"urn:com.workday/bsvc Total_Results,omitempty"`
    TotalPages   float64 `xml:"urn:com.workday/bsvc Total_Pages,omitempty"`
    PageResults  float64 `xml:"urn:com.workday/bsvc Page_Results,omitempty"`
    Page         float64 `xml:"urn:com.workday/bsvc Page,omitempty"`
}
```
The "Response_Results" element contains summary information about the data that has been returned from your request including "Total_Results", "Total_Pages", and the current "Page" returned.










## <a name="RetentionObjectIDType">type</a> [RetentionObjectIDType](/src/target/wwsgen_types.go?s=280938:281060#L4426)
``` go
type RetentionObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="RetentionObjectType">type</a> [RetentionObjectType](/src/target/wwsgen_types.go?s=281062:281266#L4431)
``` go
type RetentionObjectType struct {
    ID         []RetentionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="RetrieveSeveranceWorksheetInput">type</a> [RetrieveSeveranceWorksheetInput](/src/target/wwsgen_client.go?s=36505:36675#L935)
``` go
type RetrieveSeveranceWorksheetInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Retrieve_Severance_Worksheets_Request"`
    RetrieveSeveranceWorksheetsRequestType
}
```









## <a name="RetrieveSeveranceWorksheetOutput">type</a> [RetrieveSeveranceWorksheetOutput](/src/target/wwsgen_client.go?s=36677:36850#L940)
``` go
type RetrieveSeveranceWorksheetOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Retrieve_Severance_Worksheets_Response"`
    RetrieveSeveranceWorksheetsResponseType
}
```









## <a name="RetrieveSeveranceWorksheetsRequestType">type</a> [RetrieveSeveranceWorksheetsRequestType](/src/target/wwsgen_types.go?s=281326:282020#L4437)
``` go
type RetrieveSeveranceWorksheetsRequestType struct {
    RequestReferences *EmployeeSeveranceWorksheetEventRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *EmployeeSeveranceWorksheetEventRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *EmployeeSeveranceWorksheetEventResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    Version           string                                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Request for getting back Employee Severance Worksheets










## <a name="RetrieveSeveranceWorksheetsResponseType">type</a> [RetrieveSeveranceWorksheetsResponseType](/src/target/wwsgen_types.go?s=282091:283041#L4446)
``` go
type RetrieveSeveranceWorksheetsResponseType struct {
    RequestReferences *EmployeeSeveranceWorksheetEventRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
    RequestCriteria   *EmployeeSeveranceWorksheetEventRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
    ResponseFilter    *ResponseFilterType                                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
    ResponseGroup     *EmployeeSeveranceWorksheetEventResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
    ResponseResults   *ResponseResultsType                                  `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
    ResponseData      *EmployeeSeveranceWorksheetEventResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
    Version           string                                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Response element for Retrieve Employee Severance Worksheet Events










## <a name="ReviewRatingObjectIDType">type</a> [ReviewRatingObjectIDType](/src/target/wwsgen_types.go?s=283105:283230#L4457)
``` go
type ReviewRatingObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="ReviewRatingObjectType">type</a> [ReviewRatingObjectType](/src/target/wwsgen_types.go?s=283232:283445#L4462)
``` go
type ReviewRatingObjectType struct {
    ID         []ReviewRatingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="ReviewRatingScaleObjectIDType">type</a> [ReviewRatingScaleObjectIDType](/src/target/wwsgen_types.go?s=283509:283639#L4468)
``` go
type ReviewRatingScaleObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="ReviewRatingScaleObjectType">type</a> [ReviewRatingScaleObjectType](/src/target/wwsgen_types.go?s=283641:283869#L4473)
``` go
type ReviewRatingScaleObjectType struct {
    ID         []ReviewRatingScaleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="SalaryPayPlanObjectIDType">type</a> [SalaryPayPlanObjectIDType](/src/target/wwsgen_types.go?s=283933:284059#L4479)
``` go
type SalaryPayPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="SalaryPayPlanObjectType">type</a> [SalaryPayPlanObjectType](/src/target/wwsgen_types.go?s=284061:284277#L4484)
``` go
type SalaryPayPlanObjectType struct {
    ID         []SalaryPayPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="SalaryPlanDataType">type</a> [SalaryPlanDataType](/src/target/wwsgen_types.go?s=284318:284449#L4490)
``` go
type SalaryPlanDataType struct {
    ApplyFullTimeEquivalent *bool `xml:"urn:com.workday/bsvc Apply_Full_Time_Equivalent,omitempty"`
}
```
This is a wrapper for Salary Plans.










## <a name="SalaryPlanObjectIDType">type</a> [SalaryPlanObjectIDType](/src/target/wwsgen_types.go?s=284513:284636#L4495)
``` go
type SalaryPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="SalaryPlanObjectType">type</a> [SalaryPlanObjectType](/src/target/wwsgen_types.go?s=284638:284845#L4500)
``` go
type SalaryPlanObjectType struct {
    ID         []SalaryPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="SalaryUnitPlanObjectIDType">type</a> [SalaryUnitPlanObjectIDType](/src/target/wwsgen_types.go?s=284909:285036#L4506)
``` go
type SalaryUnitPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="SalaryUnitPlanObjectType">type</a> [SalaryUnitPlanObjectType](/src/target/wwsgen_types.go?s=285038:285257#L4511)
``` go
type SalaryUnitPlanObjectType struct {
    ID         []SalaryUnitPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="ScorecardProfileDataType">type</a> [ScorecardProfileDataType](/src/target/wwsgen_types.go?s=285287:285763#L4517)
``` go
type ScorecardProfileDataType struct {
    ScorecardProfileID                  string                                        `xml:"urn:com.workday/bsvc Scorecard_Profile_ID,omitempty"`
    ScorecardProfileTargetRuleReference *ConditionRuleObjectType                      `xml:"urn:com.workday/bsvc Scorecard_Profile_Target__Rule_Reference"`
    ScorecardProfileGoalData            []ScorecardProfilePerformanceCriteriaDataType `xml:"urn:com.workday/bsvc Scorecard_Profile_Goal_Data"`
}
```
Contains scorecard data.










## <a name="ScorecardProfilePerformanceCriteriaDataType">type</a> [ScorecardProfilePerformanceCriteriaDataType](/src/target/wwsgen_types.go?s=285803:286281#L4524)
``` go
type ScorecardProfilePerformanceCriteriaDataType struct {
    ScorecardProfileGoalID          string  `xml:"urn:com.workday/bsvc Scorecard_Profile_Goal_ID,omitempty"`
    ScorecardProfileGoalName        string  `xml:"urn:com.workday/bsvc Scorecard_Profile_Goal_Name"`
    ScorecardProfileGoalDescription string  `xml:"urn:com.workday/bsvc Scorecard_Profile_Goal_Description,omitempty"`
    ScorecardProfileGoalWeight      float64 `xml:"urn:com.workday/bsvc Scorecard_Profile_Goal_Weight"`
}
```
Contains set of goals or criteria.










## <a name="ScorecardResultDataType">type</a> [ScorecardResultDataType](/src/target/wwsgen_types.go?s=286345:286571#L4532)
``` go
type ScorecardResultDataType struct {
    GoalReference *PerformanceCriteriaObjectType `xml:"urn:com.workday/bsvc Goal_Reference"`
    Achievement   float64                        `xml:"urn:com.workday/bsvc Achievement,omitempty"`
}
```
Contains the set of goals and goal values for this result.










## <a name="ScoresetContainerObjectIDType">type</a> [ScoresetContainerObjectIDType](/src/target/wwsgen_types.go?s=286635:286923#L4538)
``` go
type ScoresetContainerObjectIDType struct {
    Value      string `xml:",chardata"`
    Type       string `xml:"urn:com.workday/bsvc type,attr"`
    Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
    Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="ScoresetContainerObjectType">type</a> [ScoresetContainerObjectType](/src/target/wwsgen_types.go?s=286925:287153#L4545)
``` go
type ScoresetContainerObjectType struct {
    ID         []ScoresetContainerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="ServiceDurationSeveranceMatrixDataType">type</a> [ServiceDurationSeveranceMatrixDataType](/src/target/wwsgen_types.go?s=287226:287600#L4551)
``` go
type ServiceDurationSeveranceMatrixDataType struct {
    LengthofServiceUnitReference            *CompensationPeriodObjectType                 `xml:"urn:com.workday/bsvc Length_of_Service_Unit_Reference"`
    ServiceDurationSeveranceMatrixEntryData []ServiceDurationSeveranceMatrixEntryDataType `xml:"urn:com.workday/bsvc Service_Duration_Severance_Matrix_Entry_Data,omitempty"`
}
```
The service length specific duration severance matrix data element.










## <a name="ServiceDurationSeveranceMatrixEntryDataType">type</a> [ServiceDurationSeveranceMatrixEntryDataType](/src/target/wwsgen_types.go?s=287635:288346#L4557)
``` go
type ServiceDurationSeveranceMatrixEntryDataType struct {
    MinimumLengthofService   float64                  `xml:"urn:com.workday/bsvc Minimum_Length_of_Service,omitempty"`
    MaximumLengthofService   float64                  `xml:"urn:com.workday/bsvc Maximum_Length_of_Service,omitempty"`
    EligibilityRuleReference *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    Duration                 float64                  `xml:"urn:com.workday/bsvc Duration,omitempty"`
    MinimumDuration          float64                  `xml:"urn:com.workday/bsvc Minimum_Duration,omitempty"`
    MaximumDuration          float64                  `xml:"urn:com.workday/bsvc Maximum_Duration,omitempty"`
}
```
The data for the matrix entry










## <a name="ServiceLengthSeveranceMatrixDataType">type</a> [ServiceLengthSeveranceMatrixDataType](/src/target/wwsgen_types.go?s=288401:289041#L4567)
``` go
type ServiceLengthSeveranceMatrixDataType struct {
    LengthofServiceUnitReference          *CompensationPeriodObjectType               `xml:"urn:com.workday/bsvc Length_of_Service_Unit_Reference"`
    CompensationRoundingRuleReference     *CompensationRoundingRuleObjectType         `xml:"urn:com.workday/bsvc Compensation_Rounding_Rule_Reference"`
    MultiplierOrderReference              *BenefitMultiplierOrderObjectType           `xml:"urn:com.workday/bsvc Multiplier_Order_Reference"`
    ServiceLengthSeveranceMatrixEntryData []ServiceLengthSeveranceMatrixEntryDataType `xml:"urn:com.workday/bsvc Service_Length_Severance_Matrix_Entry_Data"`
}
```
The Service Length Severance Matrix data element.










## <a name="ServiceLengthSeveranceMatrixEntryDataType">type</a> [ServiceLengthSeveranceMatrixEntryDataType](/src/target/wwsgen_types.go?s=289077:289812#L4575)
``` go
type ServiceLengthSeveranceMatrixEntryDataType struct {
    MinimumLengthofService    float64                  `xml:"urn:com.workday/bsvc Minimum_Length_of_Service,omitempty"`
    MaximumLengthofService    float64                  `xml:"urn:com.workday/bsvc Maximum_Length_of_Service,omitempty"`
    EligibilityRuleReference  *ConditionRuleObjectType `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    LengthofServiceMultiplier float64                  `xml:"urn:com.workday/bsvc Length_of_Service_Multiplier,omitempty"`
    MinimumDuration           float64                  `xml:"urn:com.workday/bsvc Minimum_Duration,omitempty"`
    MaximumDuration           float64                  `xml:"urn:com.workday/bsvc Maximum_Duration,omitempty"`
}
```
the data for the matrix entry.










## <a name="SeveranceMatrixFilterOptionsType">type</a> [SeveranceMatrixFilterOptionsType](/src/target/wwsgen_types.go?s=289849:290320#L4585)
``` go
type SeveranceMatrixFilterOptionsType struct {
    CompensationBasisRange         *bool `xml:"urn:com.workday/bsvc Compensation_Basis_Range,omitempty"`
    ServiceLengthBasisRange        *bool `xml:"urn:com.workday/bsvc Service_Length_Basis_Range,omitempty"`
    DurationMatrixEntryType        *bool `xml:"urn:com.workday/bsvc Duration_Matrix_Entry_Type,omitempty"`
    LengthofServiceMatrixEntryType *bool `xml:"urn:com.workday/bsvc Length_of_Service_Matrix_Entry_Type,omitempty"`
}
```
Severance Matrix Filter Options










## <a name="SeveranceMatrixabstractObjectIDType">type</a> [SeveranceMatrixabstractObjectIDType](/src/target/wwsgen_types.go?s=290384:290520#L4593)
``` go
type SeveranceMatrixabstractObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="SeveranceMatrixabstractObjectType">type</a> [SeveranceMatrixabstractObjectType](/src/target/wwsgen_types.go?s=290522:290768#L4598)
``` go
type SeveranceMatrixabstractObjectType struct {
    ID         []SeveranceMatrixabstractObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="SeverancePackageComponentDataType">type</a> [SeverancePackageComponentDataType](/src/target/wwsgen_types.go?s=290899:291369#L4604)
``` go
type SeverancePackageComponentDataType struct {
    SeverancePackageComponentTypeReference *SeverancePackageComponentTypeObjectType `xml:"urn:com.workday/bsvc Severance_Package_Component_Type_Reference,omitempty"`
    Eligible                               *bool                                    `xml:"urn:com.workday/bsvc Eligible,omitempty"`
    Comments                               string                                   `xml:"urn:com.workday/bsvc Comments,omitempty"`
}
```
Element to input data for a severance package component on a worksheet. Fields for a component type, eligibility, and comment










## <a name="SeverancePackageComponentTypeObjectIDType">type</a> [SeverancePackageComponentTypeObjectIDType](/src/target/wwsgen_types.go?s=291433:291575#L4611)
``` go
type SeverancePackageComponentTypeObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="SeverancePackageComponentTypeObjectType">type</a> [SeverancePackageComponentTypeObjectType](/src/target/wwsgen_types.go?s=291577:291841#L4616)
``` go
type SeverancePackageComponentTypeObjectType struct {
    ID         []SeverancePackageComponentTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="SeverancePackageObjectIDType">type</a> [SeverancePackageObjectIDType](/src/target/wwsgen_types.go?s=291905:292034#L4622)
``` go
type SeverancePackageObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="SeverancePackageObjectType">type</a> [SeverancePackageObjectType](/src/target/wwsgen_types.go?s=292036:292261#L4627)
``` go
type SeverancePackageObjectType struct {
    ID         []SeverancePackageObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="SeverancePayComponentDataType">type</a> [SeverancePayComponentDataType](/src/target/wwsgen_types.go?s=292371:293678#L4633)
``` go
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
```
Subedit to enter Severance Pay Component Information. Fields entered here will override default amounts.










## <a name="SeveranceServiceDateObjectIDType">type</a> [SeveranceServiceDateObjectIDType](/src/target/wwsgen_types.go?s=293742:293875#L4646)
``` go
type SeveranceServiceDateObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="SeveranceServiceDateObjectType">type</a> [SeveranceServiceDateObjectType](/src/target/wwsgen_types.go?s=293877:294114#L4651)
``` go
type SeveranceServiceDateObjectType struct {
    ID         []SeveranceServiceDateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="StockDateRuleObjectIDType">type</a> [StockDateRuleObjectIDType](/src/target/wwsgen_types.go?s=294178:294304#L4657)
``` go
type StockDateRuleObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="StockDateRuleObjectType">type</a> [StockDateRuleObjectType](/src/target/wwsgen_types.go?s=294306:294522#L4662)
``` go
type StockDateRuleObjectType struct {
    ID         []StockDateRuleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="StockGrantObjectIDType">type</a> [StockGrantObjectIDType](/src/target/wwsgen_types.go?s=294586:294709#L4668)
``` go
type StockGrantObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="StockGrantObjectType">type</a> [StockGrantObjectType](/src/target/wwsgen_types.go?s=294711:294918#L4673)
``` go
type StockGrantObjectType struct {
    ID         []StockGrantObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="StockGrantTypeObjectIDType">type</a> [StockGrantTypeObjectIDType](/src/target/wwsgen_types.go?s=294982:295109#L4679)
``` go
type StockGrantTypeObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="StockGrantTypeObjectType">type</a> [StockGrantTypeObjectType](/src/target/wwsgen_types.go?s=295111:295330#L4684)
``` go
type StockGrantTypeObjectType struct {
    ID         []StockGrantTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="StockParticipationRateTableData20Type">type</a> [StockParticipationRateTableData20Type](/src/target/wwsgen_types.go?s=295371:296282#L4690)
``` go
type StockParticipationRateTableData20Type struct {
    StockParticipationRateTableID        string                                       `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_ID,omitempty"`
    Inactive                             *bool                                        `xml:"urn:com.workday/bsvc Inactive,omitempty"`
    Name                                 string                                       `xml:"urn:com.workday/bsvc Name,omitempty"`
    Description                          string                                       `xml:"urn:com.workday/bsvc Description,omitempty"`
    DefaultPercentageRate                float64                                      `xml:"urn:com.workday/bsvc Default_Percentage_Rate,omitempty"`
    StockParticipationRateTableEntryData []StockParticipationRateTableEntryData20Type `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_Entry_Data,omitempty"`
}
```
Stock Participation Rate Table Data










## <a name="StockParticipationRateTableEntryData20Type">type</a> [StockParticipationRateTableEntryData20Type](/src/target/wwsgen_types.go?s=296329:297192#L4700)
``` go
type StockParticipationRateTableEntryData20Type struct {
    Delete                     *bool                         `xml:"urn:com.workday/bsvc Delete,omitempty"`
    EntrySortOrder             string                        `xml:"urn:com.workday/bsvc Entry_Sort_Order"`
    ManagementLevelReference   []ManagementLevelObjectType   `xml:"urn:com.workday/bsvc Management_Level_Reference,omitempty"`
    JobProfileReference        []JobProfileObjectType        `xml:"urn:com.workday/bsvc Job_Profile_Reference,omitempty"`
    CountryReference           []CountryObjectType           `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
    CompensationGradeReference []CompensationGradeObjectType `xml:"urn:com.workday/bsvc Compensation_Grade_Reference,omitempty"`
    PercentageRate             float64                       `xml:"urn:com.workday/bsvc Percentage_Rate,omitempty"`
}
```
Stock Participation Rate Table Entry Data










## <a name="StockParticipationRateTableObjectIDType">type</a> [StockParticipationRateTableObjectIDType](/src/target/wwsgen_types.go?s=297256:297396#L4711)
``` go
type StockParticipationRateTableObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="StockParticipationRateTableObjectType">type</a> [StockParticipationRateTableObjectType](/src/target/wwsgen_types.go?s=297398:297656#L4716)
``` go
type StockParticipationRateTableObjectType struct {
    ID         []StockParticipationRateTableObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="StockParticipationRateTableRequestCriteriaType">type</a> [StockParticipationRateTableRequestCriteriaType](/src/target/wwsgen_types.go?s=297678:297823#L4722)
``` go
type StockParticipationRateTableRequestCriteriaType struct {
    IncludeInactive bool `xml:"urn:com.workday/bsvc Include_Inactive,attr,omitempty"`
}
```
Request Criteria










## <a name="StockParticipationRateTableRequestReferencesType">type</a> [StockParticipationRateTableRequestReferencesType](/src/target/wwsgen_types.go?s=297847:298059#L4727)
``` go
type StockParticipationRateTableRequestReferencesType struct {
    StockParticipationRateTableReference []StockParticipationRateTableObjectType `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_Reference"`
}
```
Request References










## <a name="StockParticipationRateTableResponseDataType">type</a> [StockParticipationRateTableResponseDataType](/src/target/wwsgen_types.go?s=298109:298316#L4732)
``` go
type StockParticipationRateTableResponseDataType struct {
    StockParticipationRateTable []StockParticipationRateTableSubResponseDataType `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table,omitempty"`
}
```
Stock Participation Rate Table Response Data










## <a name="StockParticipationRateTableResponseGroupType">type</a> [StockParticipationRateTableResponseGroupType](/src/target/wwsgen_types.go?s=298336:298477#L4737)
``` go
type StockParticipationRateTableResponseGroupType struct {
    IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}
```
Response Group










## <a name="StockParticipationRateTableSubResponseDataType">type</a> [StockParticipationRateTableSubResponseDataType](/src/target/wwsgen_types.go?s=298518:298891#L4742)
``` go
type StockParticipationRateTableSubResponseDataType struct {
    StockParticipationRateTableReference *StockParticipationRateTableObjectType  `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_Reference,omitempty"`
    StockParticipationRateTableData      []StockParticipationRateTableData20Type `xml:"urn:com.workday/bsvc Stock_Participation_Rate_Table_Data,omitempty"`
}
```
Stock Participation Rate Table Data










## <a name="StockPercentPlanObjectIDType">type</a> [StockPercentPlanObjectIDType](/src/target/wwsgen_types.go?s=298955:299084#L4748)
``` go
type StockPercentPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="StockPercentPlanObjectType">type</a> [StockPercentPlanObjectType](/src/target/wwsgen_types.go?s=299086:299311#L4753)
``` go
type StockPercentPlanObjectType struct {
    ID         []StockPercentPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="StockPlanAmountDataType">type</a> [StockPlanAmountDataType](/src/target/wwsgen_types.go?s=299344:300244#L4759)
``` go
type StockPlanAmountDataType struct {
    RoundingRuleReference                 *CompensationRoundingRuleObjectType         `xml:"urn:com.workday/bsvc Rounding_Rule_Reference,omitempty"`
    UseTargetCurrency                     *bool                                       `xml:"urn:com.workday/bsvc Use_Target_Currency,omitempty"`
    TargetAmount                          float64                                     `xml:"urn:com.workday/bsvc Target_Amount,omitempty"`
    CurrencyReference                     *CurrencyObjectType                         `xml:"urn:com.workday/bsvc Currency_Reference"`
    GrantSplitReplacementData             []GrantTypeSplitReplacementDataType         `xml:"urn:com.workday/bsvc Grant_Split_Replacement_Data"`
    StockPlanAmountProfileReplacementData []StockPlanAmountProfileReplacementDataType `xml:"urn:com.workday/bsvc Stock_Plan_Amount_Profile_Replacement_Data,omitempty"`
}
```
Data for Stock Amount Plan.










## <a name="StockPlanAmountProfileReplacementDataType">type</a> [StockPlanAmountProfileReplacementDataType](/src/target/wwsgen_types.go?s=300285:300903#L4769)
``` go
type StockPlanAmountProfileReplacementDataType struct {
    TargetAmount                          float64                                     `xml:"urn:com.workday/bsvc Target_Amount,omitempty"`
    CurrencyReference                     *CurrencyObjectType                         `xml:"urn:com.workday/bsvc Currency_Reference"`
    EligibilityRuleReference              *ConditionRuleObjectType                    `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    StockProfileGrantSplitReplacementData []StockProfileGrantSplitReplacementDataType `xml:"urn:com.workday/bsvc Stock_Profile_Grant_Split_Replacement_Data"`
}
```
Profile data for Stock Amount Plan.










## <a name="StockPlanDataType">type</a> [StockPlanDataType](/src/target/wwsgen_types.go?s=300929:301772#L4777)
``` go
type StockPlanDataType struct {
    AllowTargetOverride         *bool                         `xml:"urn:com.workday/bsvc Allow_Target_Override,omitempty"`
    CompensationMatrixReference *CompensationMatrixObjectType `xml:"urn:com.workday/bsvc Compensation_Matrix_Reference,omitempty"`
    UseasReferenceOnly          *bool                         `xml:"urn:com.workday/bsvc Use_as_Reference_Only,omitempty"`
    StockPlanUnitData           *StockPlanUnitDataType        `xml:"urn:com.workday/bsvc Stock_Plan_Unit_Data"`
    StockPlanPercentData        *StockPlanPercentDataType     `xml:"urn:com.workday/bsvc Stock_Plan_Percent_Data"`
    StockPlanAmountData         *StockPlanAmountDataType      `xml:"urn:com.workday/bsvc Stock_Plan_Amount_Data"`
    HideTarget                  *bool                         `xml:"urn:com.workday/bsvc Hide_Target,omitempty"`
}
```
Data for Stock Plan.










## <a name="StockPlanObjectIDType">type</a> [StockPlanObjectIDType](/src/target/wwsgen_types.go?s=301836:301958#L4788)
``` go
type StockPlanObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="StockPlanObjectType">type</a> [StockPlanObjectType](/src/target/wwsgen_types.go?s=301960:302164#L4793)
``` go
type StockPlanObjectType struct {
    ID         []StockPlanObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="StockPlanPercentDataType">type</a> [StockPlanPercentDataType](/src/target/wwsgen_types.go?s=302198:303032#L4799)
``` go
type StockPlanPercentDataType struct {
    RoundingRuleReference                         *CompensationRoundingRuleObjectType          `xml:"urn:com.workday/bsvc Rounding_Rule_Reference,omitempty"`
    TargetPercent                                 float64                                      `xml:"urn:com.workday/bsvc Target_Percent,omitempty"`
    CompensationBasisReference                    *CompensationBasisObjectType                 `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
    GrantSplitReplacementData                     []GrantTypeSplitReplacementDataType          `xml:"urn:com.workday/bsvc Grant_Split_Replacement_Data"`
    StockPlanPercentProfileDefaultReplacementData []StockPlanPercentProfileReplacementDataType `xml:"urn:com.workday/bsvc Stock_Plan_Percent_Profile_Default_Replacement_Data,omitempty"`
}
```
Data for Stock Percent Plan.










## <a name="StockPlanPercentProfileReplacementDataType">type</a> [StockPlanPercentProfileReplacementDataType](/src/target/wwsgen_types.go?s=303066:303706#L4808)
``` go
type StockPlanPercentProfileReplacementDataType struct {
    TargetPercent                         float64                                     `xml:"urn:com.workday/bsvc Target_Percent,omitempty"`
    EligibilityRuleReference              *ConditionRuleObjectType                    `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    CompensationBasisReference            *CompensationBasisObjectType                `xml:"urn:com.workday/bsvc Compensation_Basis_Reference,omitempty"`
    StockProfileGrantSplitReplacementData []StockProfileGrantSplitReplacementDataType `xml:"urn:com.workday/bsvc Stock_Profile_Grant_Split_Replacement_Data"`
}
```
Data for Stock Percent Plan.










## <a name="StockPlanUnitDataType">type</a> [StockPlanUnitDataType](/src/target/wwsgen_types.go?s=303737:304202#L4816)
``` go
type StockPlanUnitDataType struct {
    TargetShares                        float64                                   `xml:"urn:com.workday/bsvc Target_Shares,omitempty"`
    GrantSplitReplacementData           []GrantTypeSplitReplacementDataType       `xml:"urn:com.workday/bsvc Grant_Split_Replacement_Data"`
    StockPlanUnitProfileReplacementData []StockPlanUnitProfileReplacementDataType `xml:"urn:com.workday/bsvc Stock_Plan_Unit_Profile_Replacement_Data,omitempty"`
}
```
Data for Stock Unit Plan.










## <a name="StockPlanUnitProfileReplacementDataType">type</a> [StockPlanUnitProfileReplacementDataType](/src/target/wwsgen_types.go?s=304241:304726#L4823)
``` go
type StockPlanUnitProfileReplacementDataType struct {
    TargetShares                          float64                                     `xml:"urn:com.workday/bsvc Target_Shares,omitempty"`
    EligibilityRuleReference              *ConditionRuleObjectType                    `xml:"urn:com.workday/bsvc Eligibility_Rule_Reference"`
    StockProfileGrantSplitReplacementData []StockProfileGrantSplitReplacementDataType `xml:"urn:com.workday/bsvc Stock_Profile_Grant_Split_Replacement_Data"`
}
```
Profile data for Stock Unit Plan.










## <a name="StockProfileGrantSplitReplacementDataType">type</a> [StockProfileGrantSplitReplacementDataType](/src/target/wwsgen_types.go?s=304766:305327#L4830)
``` go
type StockProfileGrantSplitReplacementDataType struct {
    StockGrantTypeReference       *StockGrantTypeObjectType       `xml:"urn:com.workday/bsvc Stock_Grant_Type_Reference"`
    StockVestingScheduleReference *StockVestingScheduleObjectType `xml:"urn:com.workday/bsvc Stock_Vesting_Schedule_Reference,omitempty"`
    StockDateRuleReference        *StockDateRuleObjectType        `xml:"urn:com.workday/bsvc Stock_Date_Rule_Reference,omitempty"`
    GrantSplitPercent             float64                         `xml:"urn:com.workday/bsvc Grant_Split_Percent,omitempty"`
}
```
Grant Split Data for Stock Profile










## <a name="StockVestingScheduleObjectIDType">type</a> [StockVestingScheduleObjectIDType](/src/target/wwsgen_types.go?s=305391:305524#L4838)
``` go
type StockVestingScheduleObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="StockVestingScheduleObjectType">type</a> [StockVestingScheduleObjectType](/src/target/wwsgen_types.go?s=305526:305763#L4843)
``` go
type StockVestingScheduleObjectType struct {
    ID         []StockVestingScheduleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="TenantedPayrollWorktagObjectIDType">type</a> [TenantedPayrollWorktagObjectIDType](/src/target/wwsgen_types.go?s=305827:306120#L4849)
``` go
type TenantedPayrollWorktagObjectIDType struct {
    Value      string `xml:",chardata"`
    Type       string `xml:"urn:com.workday/bsvc type,attr"`
    Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
    Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="TenantedPayrollWorktagObjectType">type</a> [TenantedPayrollWorktagObjectType](/src/target/wwsgen_types.go?s=306122:306365#L4856)
``` go
type TenantedPayrollWorktagObjectType struct {
    ID         []TenantedPayrollWorktagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="TimeProrationRuleObjectIDType">type</a> [TimeProrationRuleObjectIDType](/src/target/wwsgen_types.go?s=306429:306559#L4862)
``` go
type TimeProrationRuleObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="TimeProrationRuleObjectType">type</a> [TimeProrationRuleObjectType](/src/target/wwsgen_types.go?s=306561:306789#L4867)
``` go
type TimeProrationRuleObjectType struct {
    ID         []TimeProrationRuleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="TimeZoneObjectIDType">type</a> [TimeZoneObjectIDType](/src/target/wwsgen_types.go?s=306853:306974#L4873)
``` go
type TimeZoneObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="TimeZoneObjectType">type</a> [TimeZoneObjectType](/src/target/wwsgen_types.go?s=306976:307177#L4878)
``` go
type TimeZoneObjectType struct {
    ID         []TimeZoneObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="UniqueIdentifierObjectIDType">type</a> [UniqueIdentifierObjectIDType](/src/target/wwsgen_types.go?s=307241:307370#L4884)
``` go
type UniqueIdentifierObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="UniqueIdentifierObjectType">type</a> [UniqueIdentifierObjectType](/src/target/wwsgen_types.go?s=307372:307597#L4889)
``` go
type UniqueIdentifierObjectType struct {
    ID         []UniqueIdentifierObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="UnitSalaryPlanDataType">type</a> [UnitSalaryPlanDataType](/src/target/wwsgen_types.go?s=307652:308111#L4895)
``` go
type UnitSalaryPlanDataType struct {
    PerUnitAmount          float64                  `xml:"urn:com.workday/bsvc Per_Unit_Amount,omitempty"`
    DefaultUnits           float64                  `xml:"urn:com.workday/bsvc Default_Units,omitempty"`
    UnitOfMeasureReference *UnitofMeasureObjectType `xml:"urn:com.workday/bsvc Unit_Of_Measure_Reference"`
    NoIndividualOverride   *bool                    `xml:"urn:com.workday/bsvc No_Individual_Override,omitempty"`
}
```
This is a wrapper for the Base Pay Unit Salaries.










## <a name="UnitofMeasureObjectIDType">type</a> [UnitofMeasureObjectIDType](/src/target/wwsgen_types.go?s=308175:308301#L4903)
``` go
type UnitofMeasureObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="UnitofMeasureObjectType">type</a> [UnitofMeasureObjectType](/src/target/wwsgen_types.go?s=308303:308519#L4908)
``` go
type UnitofMeasureObjectType struct {
    ID         []UnitofMeasureObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="UpdateStockGrantDataType">type</a> [UpdateStockGrantDataType](/src/target/wwsgen_types.go?s=308556:310077#L4914)
``` go
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
```
Stock Grant data to be updated.










### <a name="UpdateStockGrantDataType.MarshalXML">func</a> (\*UpdateStockGrantDataType) [MarshalXML](/src/target/wwsgen_types.go?s=310079:310170#L4931)
``` go
func (t *UpdateStockGrantDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```



### <a name="UpdateStockGrantDataType.UnmarshalXML">func</a> (\*UpdateStockGrantDataType) [UnmarshalXML](/src/target/wwsgen_types.go?s=310833:310926#L4947)
``` go
func (t *UpdateStockGrantDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```



## <a name="UpdateStockGrantInput">type</a> [UpdateStockGrantInput](/src/target/wwsgen_client.go?s=17525:17663#L473)
``` go
type UpdateStockGrantInput struct {
    XMLName string `xml:"urn:com.workday/bsvc Update_Stock_Grant_Request"`
    UpdateStockGrantRequestType
}
```









## <a name="UpdateStockGrantOutput">type</a> [UpdateStockGrantOutput](/src/target/wwsgen_client.go?s=17665:17806#L478)
``` go
type UpdateStockGrantOutput struct {
    XMLName string `xml:"urn:com.workday/bsvc Update_Stock_Grant_Response"`
    UpdateStockGrantResponseType
}
```









## <a name="UpdateStockGrantRequestType">type</a> [UpdateStockGrantRequestType](/src/target/wwsgen_types.go?s=311626:311959#L4965)
``` go
type UpdateStockGrantRequestType struct {
    StockGrantReference *StockGrantObjectType     `xml:"urn:com.workday/bsvc Stock_Grant_Reference"`
    StockGrantData      *UpdateStockGrantDataType `xml:"urn:com.workday/bsvc Stock_Grant_Data"`
    Version             string                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Stock grant request










## <a name="UpdateStockGrantResponseType">type</a> [UpdateStockGrantResponseType](/src/target/wwsgen_types.go?s=311992:312235#L4972)
``` go
type UpdateStockGrantResponseType struct {
    StockGrantReference *StockGrantObjectType `xml:"urn:com.workday/bsvc Stock_Grant_Reference,omitempty"`
    Version             string                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}
```
Update Stock Grant Response










## <a name="ValidationErrorType">type</a> [ValidationErrorType](/src/target/wwsgen_types.go?s=312237:312484#L4977)
``` go
type ValidationErrorType struct {
    Message       string `xml:"urn:com.workday/bsvc Message,omitempty"`
    DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
    Xpath         string `xml:"urn:com.workday/bsvc Xpath,omitempty"`
}
```









## <a name="ValidationFaultType">type</a> [ValidationFaultType](/src/target/wwsgen_types.go?s=312486:312616#L4983)
``` go
type ValidationFaultType struct {
    ValidationError []ValidationErrorType `xml:"urn:com.workday/bsvc Validation_Error,omitempty"`
}
```









## <a name="WebServiceBackgroundProcessRuntimeObjectIDType">type</a> [WebServiceBackgroundProcessRuntimeObjectIDType](/src/target/wwsgen_types.go?s=312680:312827#L4988)
``` go
type WebServiceBackgroundProcessRuntimeObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="WebServiceBackgroundProcessRuntimeObjectType">type</a> [WebServiceBackgroundProcessRuntimeObjectType](/src/target/wwsgen_types.go?s=312829:313108#L4993)
``` go
type WebServiceBackgroundProcessRuntimeObjectType struct {
    ID         []WebServiceBackgroundProcessRuntimeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string                                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="WorkdayCommonHeaderType">type</a> [WorkdayCommonHeaderType](/src/target/wwsgen_types.go?s=313110:313275#L4998)
``` go
type WorkdayCommonHeaderType struct {
    IncludeReferenceDescriptorsInResponse *bool `xml:"urn:com.workday/bsvc Include_Reference_Descriptors_In_Response,omitempty"`
}
```









## <a name="WorkerObjectIDType">type</a> [WorkerObjectIDType](/src/target/wwsgen_types.go?s=313339:313458#L5003)
``` go
type WorkerObjectIDType struct {
    Value string `xml:",chardata"`
    Type  string `xml:"urn:com.workday/bsvc type,attr"`
}
```
Contains a unique identifier for an instance of an object.










## <a name="WorkerObjectType">type</a> [WorkerObjectType](/src/target/wwsgen_types.go?s=313460:313655#L5008)
``` go
type WorkerObjectType struct {
    ID         []WorkerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
    Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}
```









## <a name="WorkerRequestReferencesType">type</a> [WorkerRequestReferencesType](/src/target/wwsgen_types.go?s=313769:314135#L5014)
``` go
type WorkerRequestReferencesType struct {
    WorkerReference          []WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference"`
    SkipNonExistingInstances bool               `xml:"urn:com.workday/bsvc Skip_Non_Existing_Instances,attr,omitempty"`
    IgnoreInvalidReferences  bool               `xml:"urn:com.workday/bsvc Ignore_Invalid_References,attr,omitempty"`
}
```
Utilize the Request References element to retrieve a specific instance(s) of Worker and its associated data.










## <a name="WorkerResponseGroupforReferenceType">type</a> [WorkerResponseGroupforReferenceType](/src/target/wwsgen_types.go?s=314452:314584#L5021)
``` go
type WorkerResponseGroupforReferenceType struct {
    IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}
```
The response group allows for the response data to be tailored to only included elements that the user is looking for.  If no response group is provided in the request, then all groups will be returned. If the Response Group element is returned, you can select which sections of data to include in the response.














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
