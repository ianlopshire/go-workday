package wws

const (
	// Predefined XML namespaces
	XMLNSWSS          = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	XMLNSWSU          = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	XMLNSPasswordText = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"
)

const (
	Version27   = "v27.0" // Version string for WWS version 27.0
	Version27_1 = "v27.1" // Version string for WWS version 27.1
	Version28   = "v29.0" // Version string for WWS version 28.0
	Version28_1 = "v28.1" // Version string for WWS version 28.1
	Version28_2 = "v28.2" // Version string for WWS version 28.2
	Version29   = "v29.0" // Version string for WWS version 29.0
	Version29_1 = "v29.1" // Version string for WWS version 29.1
	Version29_2 = "v29.2" // Version string for WWS version 29.2
	Version30   = "v30.0" // Version string for WWS version 30.0
	Version30_1 = "v30.1" // Version string for WWS version 30.1
	Version30_2 = "v30.2" // Version string for WWS version 30.2
	Version31   = "v31.0" // Version string for WWS version 31.0
	Version31_1 = "v31.1" // Version string for WWS version 31.1
	Version31_2 = "v31.2" // Version string for WWS version 31.2
)

const (
	ServiceAbsenceManagement              = "Absence_Management"
	ServiceServiceAcademicAdvising        = "Academic_Advising"                // Web Service for retrieving objects related to Academic Advising, such as Student Academic Requirement Assignments.
	ServiceAcademicFoundation             = "Academic_Foundation"              // Public Web Service for creating, editing and retrieving the foundational objects for the Student System, such as Programs of Study, Educational Institutions, and other other objects with cross-module uses.
	ServiceAdmissions                     = "Admissions"                       // Web Service for creating, editing and retrieving objects related to Admissions, such as Applicants.
	ServiceAdoption                       = "Adoption"                         // Adoption Web Service Operations.
	ServiceBenefitsAdministration         = "Benefits_Administration"          // The Benefits Administration Web Service contains operations that expose Workday Human Capital Management Business Services benefits-related data.
	ServiceCampusEngagement               = "Campus_Engagement"                // Public Web Service for creating, editing and retrieving objects related to planning and organizing communication between Student Prospects and Recruiters.
	ServiceCashManagement                 = "Cash_Management"                  // The Cash Management Web Service contains operations that expose Workday Financials Cash Management data. It includes data relative to Banking.
	ServiceCompensation                   = "Compensation"                     // The Compensation Web Service contains operations that expose Workday Human Capital Management Business Services compensation-related data.
	ServiceCompensationReview             = "Compensation_Review"              // The Compensation Review Web Service contains operations that expose Workday Human Capital Management Business Services compensation review related data.
	ServiceDynamicDocumentGeneration      = "Dynamic_Document_Generation"      // Web service for creating, editing and retrieving objects related to Document Templates, such as Text Blocks and Text Block Categories.
	ServiceExternalIntegrations           = "External_Integrations"            // The External Integrations Web Service provides an operation that informs external systems of Integration Events that have been triggered from within Workday. The WSDL for this service provides the structure that needs to be implemented by an external system to receive event "Launch" information.
	ServiceFinancialAid                   = "Financial_Aid"                    // This web service is for the Financial Aid module
	ServiceFinancialManagement            = "Financial_Management"             // The Financial Management Web Service contains operations that expose Workday Financials data. It includes data relative to Accounts, Accounting, Business Plans, Financial Reporting, Tax, Financial Organizations, Basic Worktags, Related Worktags and more.
	ServiceHumanResources                 = "Human_Resources"                  // The Human Resources Web Service contains operations that expose Workday Human Capital Management Business Services data, including Employee, Contingent Worker and Organization information. This Web Service can be used for integration with enterprise systems including corporate directories, data analysis tools, email or other provisioning sub-systems, or any other systems needing Worker and/or Organization data.
	ServiceIdentityManagement             = "Identity_Management"              // The Identity Management Web Service contains operations that relate to Workday Identity and Access Management.
	ServiceIntegrations                   = "Integrations"                     // This Web Service contains operations related to all Integrations within the Workday system.
	ServiceInterviewFeedback              = "Interview_Feedback"               // Submits Interview Feedback for Interviewers on the Interview Schedule. Also, allows you to move a candidate to the next stage of the business process.
	ServiceInventory                      = "Inventory"                        // The Inventory Web Service contains operations that expose Workday Financials Inventory data.
	ServiceLearning                       = "Learning"                         // The Learning Web Service contains operations for creating, editing and retrieving Workday Learning related data such Courses, Course Offerings and Enrollments.
	ServiceNotification                   = "Notification"                     // The Notification Web Service provides an operation that informs external systems of business events that occur within Workday. The WSDL for this service provides the structure that needs to be implemented by an external system to receive notifications for their subscribed Workday business events. Subscriptions and Notification details (i.e. Endpoint and security information) are defined within the Workday application.
	ServicePayroll                        = "Payroll"                          // The Workday Payroll Web Service contains operations that expose Workday Payroll data for integration with third parties (such as time and attendance).
	ServicePayrollCAN                     = "Payroll_CAN"                      // The Workday Payroll CAN Web Service contains operations that expose Workday Payroll CAN data for integration with third parties (such as time and attendance).
	ServicePayrollFRA                     = "Payroll_FRA"                      // The Workday Payroll FRA Web Service contains operations that expose Workday Payroll FRA data for integration with third parties
	ServicePayrollGBR                     = "Payroll_GBR"                      // The Workday Payroll GBR Web Service contains operations that expose Workday UK Payroll data for integration with third parties.
	ServicePayrollInterface               = "Payroll_Interface"                // The Payroll Interface Web Service contains operations that expose Workday Human Capital Management Business Services data for integration with external payroll systems.
	ServicePerformanceManagement          = "Performance_Management"           // The Performance Management Web Service contains operations that expose Workday Employee Performance Management Business Services data. This Web Service can be used for integration with other employee performance management systems.
	ServiceProfessionalServicesAutomation = "Professional_Services_Automation" // The Professional Services Automation Web Service contains operations that expose Workday Financials Business Services data for integration with Professional Services Automation (PSA) systems.
	ServiceRecruiting                     = "Recruiting"                       // The Recruiting Web Service contains operations that expose Workday Human Capital Management Business Services data for integration with Talent Management and Applicant Tracking systems.
	ServiceResourceManagement             = "Resource_Management"              // The Resource Management Web Service contains operations that expose Workday Financials Resource Management data. It includes data relative to Suppliers, Supplier Accounts, Expenses, Business Assets and Projects.
	ServiceRevenueManagement              = "Revenue_Management"               // The Revenue Management Web Service contains operations that expose Workday Financials Revenue Management data. It includes data relative to Customers, Customer Accounts, Prospect and Opportunities.
	ServiceSettlementServices             = "Settlement_Services"              // Web Service used for settlement management and services.
	ServiceStaffing                       = "Staffing"                         // The Staffing Web Service operations expose Workday Human Capital Management Business Services and data. These services pertain to staffing transactions for both employees and contingent workers, such as bringing employees and contingent workers on board.
	ServiceStudentFinance                 = "Student_Finance"                  // Web service for creating, editing and retrieving objects related to Student Finance, such as charges and payments for Students.
	ServiceStudentRecords                 = "Student_Records"                  // Web Service for creating, editing and retrieving objects related to Student Records, such as Student Courses, Sections etc.
	ServiceStudentRecruiting              = "Student_Recruiting"               // Web Service for creating, editing and retrieving objects related to Student Recruiting, such as Student Recruiting Events, Campaigns, Cycles, Recruiters, and Prospects.
	ServiceTalent                         = "Talent"                           // This web service consists of operations for interfacing with the Workday Talent Management web service operations.
	ServiceTenantDataTranslation          = "Tenant_Data_Translation"          // Public Web Service to export and import translatable tenant data
	ServiceTimeTracking                   = "Time_Tracking"                    // Exposes Workday Time Tracking service to ingest time data from external systems.
	ServiceWorkdayConnect                 = "Workday_Connect"                  // Get and put web services used for communication capabilities across applications.
	ServiceWorkdayExtensibility           = "Workday_Extensibility"            // Public Web Service for extensibility features across applications.
	ServiceWorkforcePlanning              = "Workforce_Planning"               // The Workforce Planning Web Service operations expose Workforce Planning Business Services and data.
)
