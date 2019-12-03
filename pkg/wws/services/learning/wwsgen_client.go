
// Package learning
//
// The Learning Web Service contains operations for creating, editing and retrieving Workday Learning related data such Courses, Course Offerings and Enrollments.
package learning

import (
	"context"

	"github.com/ianlopshire/go-workday/pkg/wws"
)

const (
	CurrentVersion = "v31.2"
	ServiceName = "Learning"
)

type Client struct {
	*wws.Client
}


// GetSalesItems (Get_Sales_Items)
// 
// This service operation will add or update Sales Items. Sales items are goods and services which a company sells.  Sales Item data includes  Sales Item ID, Item Name, Revenue Category, Item Description, Item Identifier, Unit of Measure, Unit Price, Currency, Tax Applicability, Sales Item Group. The request criteria can be for a single transaction based on Reference, or all transaction can be retrieved if no criteria is specified.
func (c *Client) GetSalesItems(ctx context.Context, input *GetSalesItemsInput) (output *GetSalesItemsOutput, err error) {
	output = &GetSalesItemsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetSalesItemsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Sales_Items_Request"`
	GetSalesItemsRequestType
}

type GetSalesItemsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Sales_Items_Response"`
	GetSalesItemsResponseType
}

// GetCustomValidationRules (Get_Custom_Validation_Rules)
// 
// This web service operation will retrieve instances of Custom Validation Context. The response contains the reference to the Custom Validation Context returned, the Label, and Classification. Each Custom Valdiation Context can contain many Custom Validations,, which are made up of Condition Items that make up the logic used in the Custom Validations for each Rule type. On the web service request, you can specify a specific reference to a Custom Validation Context to return.
func (c *Client) GetCustomValidationRules(ctx context.Context, input *GetCustomValidationRulesInput) (output *GetCustomValidationRulesOutput, err error) {
	output = &GetCustomValidationRulesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCustomValidationRulesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Custom_Validation_Rules_Request"`
	GetCustomValidationRulesRequestType
}

type GetCustomValidationRulesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Custom_Validation_Rules_Response"`
	GetCustomValidationRulesResponseType
}

// PutCustomValidationRule (Put_Custom_Validation_Rule)
// 
// This web service operation will add or update Custom Validation Rule. A Custom Validation Rule contains a Label and Classification. Each Custom Valiation Rule can contain many Condition Rules, which are made up of Condition Items that make up the logic used in the Custom Validations for each Rule type.
func (c *Client) PutCustomValidationRule(ctx context.Context, input *PutCustomValidationRuleInput) (output *PutCustomValidationRuleOutput, err error) {
	output = &PutCustomValidationRuleOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCustomValidationRuleInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Custom_Validation_Rule_Request"`
	PutCustomValidationRuleRequestType
}

type PutCustomValidationRuleOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Custom_Validation_Rule_Response"`
	PutCustomValidationRuleResponseType
}

// ManageLearningCourse (Manage_Learning_Course)
// 
// Updates or creates blended and digital courses. Uses the Manage Course business process.
func (c *Client) ManageLearningCourse(ctx context.Context, input *ManageLearningCourseInput) (output *ManageLearningCourseOutput, err error) {
	output = &ManageLearningCourseOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ManageLearningCourseInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Learning_Course_Request"`
	ManageLearningCourseRequestType
}

type ManageLearningCourseOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Learning_Course_Response"`
	ManageLearningCourseResponseType
}

// GetLearningLessons (Get_Learning_Lessons)
// 
// Retrieves standalone lessons.
func (c *Client) GetLearningLessons(ctx context.Context, input *GetLearningLessonsInput) (output *GetLearningLessonsOutput, err error) {
	output = &GetLearningLessonsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLearningLessonsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Lessons_Request"`
	GetLearningLessonsRequestType
}

type GetLearningLessonsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Lessons_Response"`
	GetLearningLessonsResponseType
}

// ManageLesson (Manage_Lesson)
// 
// Updates or creates standalone lessons. Uses the Manage Lesson business process
func (c *Client) ManageLesson(ctx context.Context, input *ManageLessonInput) (output *ManageLessonOutput, err error) {
	output = &ManageLessonOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ManageLessonInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Lesson_Request"`
	ManageLessonRequestType
}

type ManageLessonOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Lesson_Response"`
	ManageLessonResponseType
}

// GetLearningBlendedCourses (Get_Learning_Blended_Courses)
// 
// Retrieves blended courses.
func (c *Client) GetLearningBlendedCourses(ctx context.Context, input *GetLearningBlendedCoursesInput) (output *GetLearningBlendedCoursesOutput, err error) {
	output = &GetLearningBlendedCoursesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLearningBlendedCoursesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Blended_Courses_Request"`
	GetLearningBlendedCoursesRequestType
}

type GetLearningBlendedCoursesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Blended_Courses_Response"`
	GetLearningBlendedCoursesResponseType
}

// GetLearningTopics (Get_Learning_Topics)
// 
// Retrieves topics.
func (c *Client) GetLearningTopics(ctx context.Context, input *GetLearningTopicsInput) (output *GetLearningTopicsOutput, err error) {
	output = &GetLearningTopicsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLearningTopicsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Topics_Request"`
	GetLearningTopicsRequestType
}

type GetLearningTopicsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Topics_Response"`
	GetLearningTopicsResponseType
}

// PutLearningTopic (Put_Learning_Topic)
// 
// Updates or creates topics.
func (c *Client) PutLearningTopic(ctx context.Context, input *PutLearningTopicInput) (output *PutLearningTopicOutput, err error) {
	output = &PutLearningTopicOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutLearningTopicInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Learning_Topic_Request"`
	PutLearningTopicRequestType
}

type PutLearningTopicOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Learning_Topic_Response"`
	PutLearningTopicResponseType
}

// GetLearningInstructors (Get_Learning_Instructors)
// 
// Retrieves instructors.
func (c *Client) GetLearningInstructors(ctx context.Context, input *GetLearningInstructorsInput) (output *GetLearningInstructorsOutput, err error) {
	output = &GetLearningInstructorsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLearningInstructorsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Instructors_Request"`
	GetLearningInstructorsRequestType
}

type GetLearningInstructorsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Instructors_Response"`
	GetLearningInstructorsResponseType
}

// PutLearningInstructor (Put_Learning_Instructor)
// 
// Updates or creates instructors.
func (c *Client) PutLearningInstructor(ctx context.Context, input *PutLearningInstructorInput) (output *PutLearningInstructorOutput, err error) {
	output = &PutLearningInstructorOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutLearningInstructorInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Learning_Instructors_Request"`
	PutLearningInstructorsRequestType
}

type PutLearningInstructorOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Learning_Instructors_Response"`
	PutLearningInstructorsResponseType
}

// GetLearningOtherUnitTypes (Get_Learning_Other_Unit_Types)
// 
// Retrieves other unit types for digital and blended courses.
func (c *Client) GetLearningOtherUnitTypes(ctx context.Context, input *GetLearningOtherUnitTypesInput) (output *GetLearningOtherUnitTypesOutput, err error) {
	output = &GetLearningOtherUnitTypesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLearningOtherUnitTypesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Other_Unit_Types_Request"`
	GetLearningOtherUnitTypesRequestType
}

type GetLearningOtherUnitTypesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Other_Unit_Types_Response"`
	GetLearningOtherUnitTypesResponseType
}

// PutLearningOtherUnitType (Put_Learning_Other_Unit_Type)
// 
// Updates or creates other unit types for digital and blended courses.
func (c *Client) PutLearningOtherUnitType(ctx context.Context, input *PutLearningOtherUnitTypeInput) (output *PutLearningOtherUnitTypeOutput, err error) {
	output = &PutLearningOtherUnitTypeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutLearningOtherUnitTypeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Learning_Other_Unit_Type_Request"`
	PutLearningOtherUnitTypeRequestType
}

type PutLearningOtherUnitTypeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Learning_Other_Unit_Type_Response"`
	PutLearningOtherUnitTypeResponseType
}

// GetLearningDigitalCourses (Get_Learning_Digital_Courses)
// 
// Retrieves digital courses.
func (c *Client) GetLearningDigitalCourses(ctx context.Context, input *GetLearningDigitalCoursesInput) (output *GetLearningDigitalCoursesOutput, err error) {
	output = &GetLearningDigitalCoursesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLearningDigitalCoursesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Digital_Courses_Request"`
	GetLearningDigitalCoursesRequestType
}

type GetLearningDigitalCoursesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Digital_Courses_Response"`
	GetLearningDigitalCoursesResponseType
}

// GetLearningCourseOfferings (Get_Learning_Course_Offerings)
// 
// Retrieves blended course offerings.
func (c *Client) GetLearningCourseOfferings(ctx context.Context, input *GetLearningCourseOfferingsInput) (output *GetLearningCourseOfferingsOutput, err error) {
	output = &GetLearningCourseOfferingsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLearningCourseOfferingsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Course_Offerings_Request"`
	GetLearningCourseOfferingsRequestType
}

type GetLearningCourseOfferingsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Course_Offerings_Response"`
	GetLearningCourseOfferingsResponseType
}

// PutLearningCourseOffering (Put_Learning_Course_Offering)
// 
// Updates or creates blended course offerings.
func (c *Client) PutLearningCourseOffering(ctx context.Context, input *PutLearningCourseOfferingInput) (output *PutLearningCourseOfferingOutput, err error) {
	output = &PutLearningCourseOfferingOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutLearningCourseOfferingInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Learning_Course_Offering_Request"`
	PutLearningCourseOfferingRequestType
}

type PutLearningCourseOfferingOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Learning_Course_Offering_Response"`
	PutLearningCourseOfferingResponseType
}

// GetLearningEnrollments (Get_Learning_Enrollments)
// 
// Retrieves enrollments.
func (c *Client) GetLearningEnrollments(ctx context.Context, input *GetLearningEnrollmentsInput) (output *GetLearningEnrollmentsOutput, err error) {
	output = &GetLearningEnrollmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLearningEnrollmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Enrollments_Request"`
	GetLearningEnrollmentsRequestType
}

type GetLearningEnrollmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Enrollments_Response"`
	GetLearningEnrollmentsResponseType
}

// PutLearningEnrollment (Put_Learning_Enrollment)
// 
// Updates or creates legacy enrollments.
func (c *Client) PutLearningEnrollment(ctx context.Context, input *PutLearningEnrollmentInput) (output *PutLearningEnrollmentOutput, err error) {
	output = &PutLearningEnrollmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutLearningEnrollmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Learning_Enrollment_Request"`
	PutLearningEnrollmentRequestType
}

type PutLearningEnrollmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Learning_Enrollment_Response"`
	PutLearningEnrollmentResponseType
}

// GetLearningSecurityCategories (Get_Learning_Security_Categories)
// 
// Returns Learning Security Category by Name. Returns all Learning Security Category if no Name is provided.
func (c *Client) GetLearningSecurityCategories(ctx context.Context, input *GetLearningSecurityCategoriesInput) (output *GetLearningSecurityCategoriesOutput, err error) {
	output = &GetLearningSecurityCategoriesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLearningSecurityCategoriesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Security_Categories_Request"`
	GetLearningSecurityCategoriesRequestType
}

type GetLearningSecurityCategoriesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Security_Categories_Response"`
	GetLearningSecurityCategoriesResponseType
}

// PutLearningSecurityCategory (Put_Learning_Security_Category)
// 
// Creates a new Learning Security Category (or updates an existing Learning Security Category) with the information supplied in the request.
func (c *Client) PutLearningSecurityCategory(ctx context.Context, input *PutLearningSecurityCategoryInput) (output *PutLearningSecurityCategoryOutput, err error) {
	output = &PutLearningSecurityCategoryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutLearningSecurityCategoryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Learning_Security_Category_Request"`
	PutLearningSecurityCategoryRequestType
}

type PutLearningSecurityCategoryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Learning_Security_Category_Response"`
	PutLearningSecurityCategoryResponseType
}

// GetLearningTopicSecuritySegments (Get_Learning_Topic_Security_Segments)
// 
// Returns Learning Topic Security Segments by Reference ID, or returns all Learning Topic Security Segments if no Reference ID is provided
func (c *Client) GetLearningTopicSecuritySegments(ctx context.Context, input *GetLearningTopicSecuritySegmentsInput) (output *GetLearningTopicSecuritySegmentsOutput, err error) {
	output = &GetLearningTopicSecuritySegmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLearningTopicSecuritySegmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Topic_Security_Segments_Request"`
	GetLearningTopicSecuritySegmentsRequestType
}

type GetLearningTopicSecuritySegmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Topic_Security_Segments_Response"`
	GetLearningTopicSecuritySegmentsResponseType
}

// PutLearningTopicSecuritySegment (Put_Learning_Topic_Security_Segment)
// 
// Creates a new Learning Topic Security Segment (or updates an existing Learning Topic Security Segment) with the information supplied in the request.
func (c *Client) PutLearningTopicSecuritySegment(ctx context.Context, input *PutLearningTopicSecuritySegmentInput) (output *PutLearningTopicSecuritySegmentOutput, err error) {
	output = &PutLearningTopicSecuritySegmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutLearningTopicSecuritySegmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Learning_Topic_Security_Segment_Request"`
	PutLearningTopicSecuritySegmentRequestType
}

type PutLearningTopicSecuritySegmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Learning_Topic_Security_Segment_Response"`
	PutLearningTopicSecuritySegmentResponseType
}

// EnrollInLearningCourse (Enroll_In_Learning_Course)
// 
// Uses the Enroll in Course business process to enroll learners into blended or digital courses.
func (c *Client) EnrollInLearningCourse(ctx context.Context, input *EnrollInLearningCourseInput) (output *EnrollInLearningCourseOutput, err error) {
	output = &EnrollInLearningCourseOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type EnrollInLearningCourseInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Enroll_In_Learning_Course_Request"`
	EnrollInLearningCourseRequestType
}

type EnrollInLearningCourseOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Enroll_In_Learning_Course_Response"`
	EnrollInLearningCourseResponseType
}

// PutExtendedEnterpriseLearner (Put_Extended_Enterprise_Learner)
// 
// Updates or creates Extended Enterprise Learners.
func (c *Client) PutExtendedEnterpriseLearner(ctx context.Context, input *PutExtendedEnterpriseLearnerInput) (output *PutExtendedEnterpriseLearnerOutput, err error) {
	output = &PutExtendedEnterpriseLearnerOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutExtendedEnterpriseLearnerInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Extended_Enterprise_Learner_Request"`
	PutExtendedEnterpriseLearnerRequestType
}

type PutExtendedEnterpriseLearnerOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Extended_Enterprise_Learner_Response"`
	PutExtendedEnterpriseLearnerResponseType
}

// GetExtendedEnterpriseLearners (Get_Extended_Enterprise_Learners)
// 
// Information about the specified Extended Enterprise Learners identified in the request. If the request doesn&#39;t specify the Extended Enterprise Learners to return, contains information for all External Learners in the system.
func (c *Client) GetExtendedEnterpriseLearners(ctx context.Context, input *GetExtendedEnterpriseLearnersInput) (output *GetExtendedEnterpriseLearnersOutput, err error) {
	output = &GetExtendedEnterpriseLearnersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetExtendedEnterpriseLearnersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Extended_Enterprise_Learners_Request"`
	GetExtendedEnterpriseLearnersRequestType
}

type GetExtendedEnterpriseLearnersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Extended_Enterprise_Learners_Response"`
	GetExtendedEnterpriseLearnersResponseType
}

// GetExtendedEnterpriseAffiliations (Get_Extended_Enterprise_Affiliations)
// 
// Retrieve Extended Enterprise Affiliation(s).
func (c *Client) GetExtendedEnterpriseAffiliations(ctx context.Context, input *GetExtendedEnterpriseAffiliationsInput) (output *GetExtendedEnterpriseAffiliationsOutput, err error) {
	output = &GetExtendedEnterpriseAffiliationsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetExtendedEnterpriseAffiliationsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Extended_Enterprise_Affiliations_Request"`
	GetExtendedEnterpriseAffiliationsRequestType
}

type GetExtendedEnterpriseAffiliationsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Extended_Enterprise_Affiliations_Response"`
	GetExtendedEnterpriseAffiliationsResponseType
}

// PutExtendedEnterpriseAffiliation (Put_Extended_Enterprise_Affiliation)
// 
// Create or Update an Extended Enterprise Affiliation.
func (c *Client) PutExtendedEnterpriseAffiliation(ctx context.Context, input *PutExtendedEnterpriseAffiliationInput) (output *PutExtendedEnterpriseAffiliationOutput, err error) {
	output = &PutExtendedEnterpriseAffiliationOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutExtendedEnterpriseAffiliationInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Extended_Enterprise_Affiliation_Request"`
	PutExtendedEnterpriseAffiliationRequestType
}

type PutExtendedEnterpriseAffiliationOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Extended_Enterprise_Affiliation_Response"`
	PutExtendedEnterpriseAffiliationResponseType
}

// GetLearningPrograms (Get_Learning_Programs)
// 
// Retrieves Learning Programs.
func (c *Client) GetLearningPrograms(ctx context.Context, input *GetLearningProgramsInput) (output *GetLearningProgramsOutput, err error) {
	output = &GetLearningProgramsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetLearningProgramsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Programs_Request"`
	GetLearningProgramsRequestType
}

type GetLearningProgramsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Learning_Programs_Response"`
	GetLearningProgramsResponseType
}

// ManageLearningProgram (Manage_Learning_Program)
// 
// Updates or creates Learning Programs. Uses the Manage Learning Program business process
func (c *Client) ManageLearningProgram(ctx context.Context, input *ManageLearningProgramInput) (output *ManageLearningProgramOutput, err error) {
	output = &ManageLearningProgramOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ManageLearningProgramInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Learning_Program_Request"`
	ManageLearningProgramRequestType
}

type ManageLearningProgramOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Manage_Learning_Program_Response"`
	ManageLearningProgramResponseType
}

