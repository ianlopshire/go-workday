
// Package cashmgmt
//
// The Cash Management  Web Service contains operations that expose Workday Financials Cash Management data. It includes data relative to Banking.
package cashmgmt

import (
	"context"

	"github.com/ianlopshire/go-workday/pkg/wws"
)

const (
	CurrentVersion = "v31.2"
	ServiceName = "Cash_Management"
)

type Client struct {
	*wws.Client
}


// SubmitAdHocBankTransaction (Submit_Ad_Hoc_Bank_Transaction)
// 
// This service operation will add or update Ad hoc Bank Transactions and submit to the Ad hoc Bank Transaction business process. Ad hoc Bank Transaction data includes Ad hoc Bank Transaction ID, Submit for Approval flag, Locked in Workday flag, Transaction Date, Transaction Memo, Company, Currency, Bank Account, Transaction Amount, Deposit flag, Withdraw flag, Purpose, Transaction Reference Text and Transaction Line data. Transaction Lines data includes Inter Company affiliate, Spend Category, Revenue Category, Ledger Account, Line Amount, Line Memo, and Worktags.
// 
// Spend Category and Resource Category are synonymous and refer to the same business object.
func (c *Client) SubmitAdHocBankTransaction(ctx context.Context, input *SubmitAdHocBankTransactionInput) (output *SubmitAdHocBankTransactionOutput, err error) {
	output = &SubmitAdHocBankTransactionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitAdHocBankTransactionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Ad_hoc_Bank_Transaction_Request"`
	SubmitAdhocBankTransactionRequestType
}

type SubmitAdHocBankTransactionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Ad_hoc_Bank_Transaction_Response"`
	SubmitAdhocBankTransactionResponseType
}

// GetBankAccounts (Get_Bank_Accounts)
// 
// This service operation wil get Bank Accounts for the specified criteria. Bank Account data includes the Bank Account ID, Account Name, Financial Institution, Bank Branch, Account Closed flag, Routing Transit or Institution Number, Financial Account Number, Bank Identifier Code, IBAN, Payment Type, Check Print Layout, Advanced Mode flag, Automatic Reconciliation Type, Reconciliation Rule Set, First Notice Reconciliation, First Notice Rule Set, Bank Electronic Payments flag, Used by Cash flag, Used by Customer Payment flag, Used by Expense Payments flag, Used by Payroll flag, Used by Supplier Payments flag, Used by Intercompany payments flag, Last Check Number Used, Enable Positive Pay flag, and Lockbox data.  The request criteria can be for a single transaction based on Reference, or all transaction can be retrieved if no criteria is specified.
func (c *Client) GetBankAccounts(ctx context.Context, input *GetBankAccountsInput) (output *GetBankAccountsOutput, err error) {
	output = &GetBankAccountsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBankAccountsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Accounts_Request"`
	GetBankAccountsRequestType
}

type GetBankAccountsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Accounts_Response"`
	GetBankAccountsResponseType
}

// PutBankAccount (Put_Bank_Account)
// 
// This service operation will add or update Bank Accounts.  Bank Account data includes the Bank Account ID, Account Name, Financial Institution, Bank Branch, Account Closed flag, Routing Transit or Institution Number, Financial Account Number, Bank Identifier Code, IBAN, Payment Type, Check Print Layout, Advanced Mode flag, Automatic Reconciliation Type, Reconciliation Rule Set, First Notice Reconciliation, First Notice Rule Set, Bank Electronic Payments flag, Used by Cash flag, Used by Customer Payment flag, Used by Expense Payments flag, Used by Payroll flag, Used by Supplier Payments flag, Used by Intercompany payments flag, Last Check Number Used, Enable Positive Pay flag, and Lockbox data.
func (c *Client) PutBankAccount(ctx context.Context, input *PutBankAccountInput) (output *PutBankAccountOutput, err error) {
	output = &PutBankAccountOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBankAccountInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Account_Request"`
	PutBankAccountRequestType
}

type PutBankAccountOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Account_Response"`
	PutBankAccountResponseType
}

// GetFinancialInstitutions (Get_Financial_Institutions)
// 
// This service operation will get Get  Financial Institutions for the specified criteria. A Financial Institution is a Business Entity. Financial Institution data includes Financial Institution ID, Financial Institution Reference ID,  Financial Institution Name, Bank Identification Code, Bank Code, and Business Entity Data.  Business Entity data includes Name, Tax ID, External ID, Contact Information data, and Business Entity Logo. Contact Information data includes Address, Phone, Email, Instant Messenger and Web Address data. The request criteria can be for a single transaction based on Reference, or all transaction can be retrieved if no criteria is specified.
func (c *Client) GetFinancialInstitutions(ctx context.Context, input *GetFinancialInstitutionsInput) (output *GetFinancialInstitutionsOutput, err error) {
	output = &GetFinancialInstitutionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetFinancialInstitutionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Financial_Institutions_Request"`
	GetFinancialInstitutionsRequestType
}

type GetFinancialInstitutionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Financial_Institutions_Response"`
	GetFinancialInstitutionsResponseType
}

// PutFinancialInstitution (Put_Financial_Institution)
// 
// This service operation will add or update Financial Institutions.  A Financial Institution is a Business Entity. Financial Institution data includes Financial Institution ID, Financial Institution Reference ID,  Financial Institution Name, Bank Identification Code, Bank Code, and Business Entity Data.  Business Entity data includes Name, Tax ID, External ID, Contact Information data, and Business Entity Logo.
func (c *Client) PutFinancialInstitution(ctx context.Context, input *PutFinancialInstitutionInput) (output *PutFinancialInstitutionOutput, err error) {
	output = &PutFinancialInstitutionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutFinancialInstitutionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Financial_Institution_Request"`
	PutFinancialInstitutionRequestType
}

type PutFinancialInstitutionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Financial_Institution_Response"`
	PutFinancialInstitutionResponseType
}

// GetBankBranches (Get_Bank_Branches)
// 
// This service operation will get Get Bank Branches for the specified criteria.  A Bank Branch is a Business Entity.  Bank Branch data include Bank Branch ID, Financial Institution, Bank Branch Number, and Business Entity Data. Business Entity data includes Name, Tax ID, External ID, Contact Information data, and Business Entity Logo. Contact Information data includes Address, Phone, Email, Instant Messenger and Web Address data.  The request criteria can be for a single transaction based on Reference, or all transaction can be retrieved if no criteria is specified.
func (c *Client) GetBankBranches(ctx context.Context, input *GetBankBranchesInput) (output *GetBankBranchesOutput, err error) {
	output = &GetBankBranchesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBankBranchesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Branches_Request"`
	GetBankBranchesRequestType
}

type GetBankBranchesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Branches_Response"`
	GetBankBranchesResponseType
}

// PutBankBranch (Put_Bank_Branch)
// 
// This service operation will add or update Bank Branches.  A Bank Branch is a Business Entity.  Bank Branch data include Bank Branch ID, Financial Institution, Bank Branch Number, and Business Entity Data. Business Entity data includes Name, Tax ID, External ID, Contact Information data, and Business Entity Logo. Contact Information data includes Address, Phone, Email, Instant Messenger and Web Address data.
func (c *Client) PutBankBranch(ctx context.Context, input *PutBankBranchInput) (output *PutBankBranchOutput, err error) {
	output = &PutBankBranchOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBankBranchInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Branch_Request"`
	PutBankBranchRequestType
}

type PutBankBranchOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Branch_Response"`
	PutBankBranchResponseType
}

// GetPettyCashAccounts (Get_Petty_Cash_Accounts)
// 
// This service operation will get Get Petty Cash Accounts for the specified criteria. Petty Cash Account data includes Petty Cash Account ID, Account Name, Financial Party, Payment Type, Used by Cash flag, Used by Customer Payments flag, Used by Expense Payments flag, Used by Payroll flag, Used by Supplier Payments flag, and Used by Inter Company Payments flag.  The request criteria can be for a single transaction based on Reference, or all transaction can be retrieved if no criteria is specified.
func (c *Client) GetPettyCashAccounts(ctx context.Context, input *GetPettyCashAccountsInput) (output *GetPettyCashAccountsOutput, err error) {
	output = &GetPettyCashAccountsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPettyCashAccountsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Petty_Cash_Accounts_Request"`
	GetPettyCashAccountsRequestType
}

type GetPettyCashAccountsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Petty_Cash_Accounts_Response"`
	GetPettyCashAccountsResponseType
}

// PutPettyCashAccount (Put_Petty_Cash_Account)
// 
// This service operation will add or update Petty Cash Accounts.  Petty Cash Account data includes Petty Cash Account ID, Account Name, Financial Party, Payment Type, Used by Cash flag, Used by Customer Payments flag, Used by Expense Payments flag, Used by Payroll flag, Used by Supplier Payments flag, and Used by Inter Company Payments flag.
func (c *Client) PutPettyCashAccount(ctx context.Context, input *PutPettyCashAccountInput) (output *PutPettyCashAccountOutput, err error) {
	output = &PutPettyCashAccountOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPettyCashAccountInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Petty_Cash_Account_Request"`
	PutPettyCashAccountRequestType
}

type PutPettyCashAccountOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Petty_Cash_Account_Response"`
	PutPettyCashAccountResponseType
}

// CancelAdHocBankTransaction (Cancel_Ad_Hoc_Bank_Transaction)
// 
// This service operation will cancel an Ad Hoc Bank Transaction
func (c *Client) CancelAdHocBankTransaction(ctx context.Context, input *CancelAdHocBankTransactionInput) (output *CancelAdHocBankTransactionOutput, err error) {
	output = &CancelAdHocBankTransactionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CancelAdHocBankTransactionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Cancel_Ad_hoc_Bank_Transaction_Request"`
	CancelAdhocBankTransactionRequestType
}

type CancelAdHocBankTransactionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Cancel_Ad_hoc_Bank_Transaction_Response"`
	CancelAdhocBankTransactionResponseType
}

// GetBusinessEntityContacts (Get_Business_Entity_Contacts)
// 
// This service operation will get Get Business Entity Contacts for the specified criteria.  The data includes the Business Entity Contact ID, the Business Entity the Contact is for (Supplier, Customer, Financial Institution or Tax Authority), the Personal Data including Name Detail data and Contact Information data.  Name Detail data includes formatted Name, Country, Prefix, First Name, Middle Name, Last Name, Secondary Last Name, and Name Suffix.  Contact Information data includes Address, Phone, Email, Instance Messenger and Web Address data. The request criteria can be for a single transaction based on Reference, or all transactions can be retrieved if no criteria is specified.
func (c *Client) GetBusinessEntityContacts(ctx context.Context, input *GetBusinessEntityContactsInput) (output *GetBusinessEntityContactsOutput, err error) {
	output = &GetBusinessEntityContactsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBusinessEntityContactsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Business_Entity_Contacts_Request"`
	GetBusinessEntityContactsRequestType
}

type GetBusinessEntityContactsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Business_Entity_Contacts_Response"`
	GetBusinessEntityContactsResponseType
}

// PutBusinessEntityContact (Put_Business_Entity_Contact)
// 
// This service operation will add or update Business Entity Contacts.  The data includes the Business Entity Contact ID, the Business Entity the Contact is for (Supplier, Customer, Financial Institution or Tax Authority), the Personal Data including Name Detail data and Contact Information data.  Name Detail data includes formatted Name, Country, Prefix, First Name, Middle Name, Last Name, Secondary Last Name, and Name Suffix.  Contact Information data includes Address, Phone, Email, Instance Messenger and Web Address data.
func (c *Client) PutBusinessEntityContact(ctx context.Context, input *PutBusinessEntityContactInput) (output *PutBusinessEntityContactOutput, err error) {
	output = &PutBusinessEntityContactOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBusinessEntityContactInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Business_Entity_Contact_Request"`
	PutBusinessEntityContactRequestType
}

type PutBusinessEntityContactOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Business_Entity_Contact_Response"`
	PutBusinessEntityContactResponseType
}

// GetPaymentElectionEnrollments (Get_Payment_Election_Enrollments)
// 
// Get Payment Election Enrollments
func (c *Client) GetPaymentElectionEnrollments(ctx context.Context, input *GetPaymentElectionEnrollmentsInput) (output *GetPaymentElectionEnrollmentsOutput, err error) {
	output = &GetPaymentElectionEnrollmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPaymentElectionEnrollmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payment_Election_Enrollments_Request"`
	GetPaymentElectionEnrollmentsRequestType
}

type GetPaymentElectionEnrollmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payment_Election_Enrollments_Response"`
	GetPaymentElectionEnrollmentsResponseType
}

// SubmitPaymentElectionEnrollment (Submit_Payment_Election_Enrollment)
// 
// Replaces current payment elections with new elections.
func (c *Client) SubmitPaymentElectionEnrollment(ctx context.Context, input *SubmitPaymentElectionEnrollmentInput) (output *SubmitPaymentElectionEnrollmentOutput, err error) {
	output = &SubmitPaymentElectionEnrollmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitPaymentElectionEnrollmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Payment_Election_Enrollment_Request"`
	SubmitPaymentElectionEnrollmentRequestType
}

type SubmitPaymentElectionEnrollmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payment_Election_Enrollment_Response"`
	PutPaymentElectionEnrollmentResponseType
}

// GetAdHocBankTransactions (Get_Ad_Hoc_Bank_Transactions)
// 
// This service operation wil retrieve Ad hoc Bank Transactions for the specified criteria.  Ad hoc Bank Transaction data includes Ad hoc Bank Transaction ID, Submit for Approval flag, Locked in Workday flag, Transaction Date, Transaction Memo, Company, Currency, Bank Account, Transaction Amount, Deposit flag, Withdraw flag, Purpose, Transaction Reference Text and Transaction Line data. Transaction Lines data includes Inter Company affiliate, Spend Category, Revenue Category, Ledger Account, Line Amount, Line Memo, and Worktags.  The request criteria can be for a single transaction based on Reference, select transaction based on criteria of Company, Bank Account, or transaction start and end date, or all transaction can be retrieved if no criteria is specified.
func (c *Client) GetAdHocBankTransactions(ctx context.Context, input *GetAdHocBankTransactionsInput) (output *GetAdHocBankTransactionsOutput, err error) {
	output = &GetAdHocBankTransactionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAdHocBankTransactionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Ad_hoc_Bank_Transactions_Request"`
	GetAdhocBankTransactionsRequestType
}

type GetAdHocBankTransactionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Ad_hoc_Bank_Transactions_Response"`
	GetAdhocBankTransactionsResponseType
}

// PutPaymentAcknowledgementMessage (Put_Payment_Acknowledgement_Message)
// 
// This service operation will add or update Payment Acknowledgements or Payment File Acknowledgements providing updates to existing payment status and error code information.
func (c *Client) PutPaymentAcknowledgementMessage(ctx context.Context, input *PutPaymentAcknowledgementMessageInput) (output *PutPaymentAcknowledgementMessageOutput, err error) {
	output = &PutPaymentAcknowledgementMessageOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPaymentAcknowledgementMessageInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payment_Acknowledgement_Message_Request"`
	PutPaymentAcknowledgementMessageRequestType
}

type PutPaymentAcknowledgementMessageOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payment_Acknowledgement_Message_Response"`
	PutPaymentAcknowledgementMessageResponseType
}

// GetAdHocPayments (Get_Ad_Hoc_Payments)
// 
// This service operation wil get Ad Hoc Paymnets for the specified criteria.
func (c *Client) GetAdHocPayments(ctx context.Context, input *GetAdHocPaymentsInput) (output *GetAdHocPaymentsOutput, err error) {
	output = &GetAdHocPaymentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAdHocPaymentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Ad_hoc_Payments_Request"`
	GetAdhocPaymentsRequestType
}

type GetAdHocPaymentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Ad_hoc_Payments_Response"`
	GetAdhocPaymentsResponseType
}

// SubmitAdHocPayment (Submit_Ad_Hoc_Payment)
// 
// This service operation will add or update Ad Hoc Payment and submit to the Ad Hoc Payment business process
func (c *Client) SubmitAdHocPayment(ctx context.Context, input *SubmitAdHocPaymentInput) (output *SubmitAdHocPaymentOutput, err error) {
	output = &SubmitAdHocPaymentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitAdHocPaymentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Ad_hoc_Payment_Request"`
	SubmitAdhocPaymentRequestType
}

type SubmitAdHocPaymentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Ad_hoc_Payment_Response"`
	SubmitAdhocPaymentResponseType
}

// CancelAdHocPayment (Cancel_Ad_Hoc_Payment)
// 
// This service operation will cancel an Ad Hoc Payment
func (c *Client) CancelAdHocPayment(ctx context.Context, input *CancelAdHocPaymentInput) (output *CancelAdHocPaymentOutput, err error) {
	output = &CancelAdHocPaymentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CancelAdHocPaymentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Cancel_Ad_hoc_Payment_Request"`
	CancelAdhocPaymentRequestType
}

type CancelAdHocPaymentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Cancel_Ad_hoc_Payment_Response"`
	CancelAdhocPaymentResponseType
}

// GetPayments (Get_Payments)
// 
// This service operation will get Payments for the specified criteria. The request criteria can be for a single payment based on Reference, or all payments for a specified company.
func (c *Client) GetPayments(ctx context.Context, input *GetPaymentsInput) (output *GetPaymentsOutput, err error) {
	output = &GetPaymentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPaymentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payments_Request"`
	GetPaymentsRequestType
}

type GetPaymentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payments_Response"`
	GetPaymentsResponseType
}

// GetPaymentMessages (Get_Payment_Messages)
// 
// This service operation will get Payment Messages for the specified criteria. The request criteria can be for a single payment message based on Reference, or all payment messages for the specified criteria.
func (c *Client) GetPaymentMessages(ctx context.Context, input *GetPaymentMessagesInput) (output *GetPaymentMessagesOutput, err error) {
	output = &GetPaymentMessagesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPaymentMessagesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payment_Messages_Request"`
	GetPaymentMessagesRequestType
}

type GetPaymentMessagesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payment_Messages_Response"`
	GetPaymentMessagesResponseType
}

// GetPaymentElectionOptions (Get_Payment_Election_Options)
// 
// Web service to get the Payment Election Options by Reference IDs, or returns all Payment Election Options if no Reference ID is provided. Includes country and currency overrides for each payment election rule.
func (c *Client) GetPaymentElectionOptions(ctx context.Context, input *GetPaymentElectionOptionsInput) (output *GetPaymentElectionOptionsOutput, err error) {
	output = &GetPaymentElectionOptionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetPaymentElectionOptionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payment_Election_Options_Request"`
	GetPaymentElectionOptionsRequestType
}

type GetPaymentElectionOptionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Payment_Election_Options_Response"`
	GetPaymentElectionOptionsResponseType
}

// PutPaymentElectionOption (Put_Payment_Election_Option)
// 
// Web service to create a new Payment Election Option or updates an existing Payment Election Option by Reference ID. Includes country and currency overrides for each payment election rule.
func (c *Client) PutPaymentElectionOption(ctx context.Context, input *PutPaymentElectionOptionInput) (output *PutPaymentElectionOptionOutput, err error) {
	output = &PutPaymentElectionOptionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutPaymentElectionOptionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payment_Election_Option_Request"`
	PutPaymentElectionOptionRequestType
}

type PutPaymentElectionOptionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Payment_Election_Option_Response"`
	PutPaymentElectionOptionResponseType
}

// PutBankStatementFile (Put_Bank_Statement_File)
// 
// This service operation will add or update Bank Statement Files. Each file will load one or more Bank Statements.
func (c *Client) PutBankStatementFile(ctx context.Context, input *PutBankStatementFileInput) (output *PutBankStatementFileOutput, err error) {
	output = &PutBankStatementFileOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBankStatementFileInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Statement_File_Request"`
	PutBankStatementFileRequestType
}

type PutBankStatementFileOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Statement_File_Response"`
	PutBankStatementFileResponseType
}

// PutBankStatement (Put_Bank_Statement)
// 
// This service operation will add or update Bank Statement Data and attach it to an existing Bank Statement File.
func (c *Client) PutBankStatement(ctx context.Context, input *PutBankStatementInput) (output *PutBankStatementOutput, err error) {
	output = &PutBankStatementOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBankStatementInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Statement_Request"`
	PutBankStatementRequestType
}

type PutBankStatementOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Statement_Response"`
	PutBankStatementResponseType
}

// GetBankStatementFiles (Get_Bank_Statement_Files)
// 
// This service operation will get all Bank Statement Files for the specified criteria. The response will include specific information about the Bank Statement Files such as the Bank Statement File References, File Creation Date, File Identification Number, Receiver Identification, Sender Identification, and Bank Statement References.
func (c *Client) GetBankStatementFiles(ctx context.Context, input *GetBankStatementFilesInput) (output *GetBankStatementFilesOutput, err error) {
	output = &GetBankStatementFilesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBankStatementFilesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Statement_File_Request"`
	GetBankStatementFileRequestType
}

type GetBankStatementFilesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Statement_Files_Response"`
	GetBankStatementFilesResponseType
}

// GetBankStatements (Get_Bank_Statements)
// 
// This service operation will get all Bank Statement Files for the specified criteria. The response will include specific information about Bank Statements.
func (c *Client) GetBankStatements(ctx context.Context, input *GetBankStatementsInput) (output *GetBankStatementsOutput, err error) {
	output = &GetBankStatementsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBankStatementsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Statements_Request"`
	GetBankStatementsRequestType
}

type GetBankStatementsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Statements_Response"`
	GetBankStatementsResponseType
}

// GetDonors (Get_Donors)
// 
// This service operation will get Donors for the specified criteria.
func (c *Client) GetDonors(ctx context.Context, input *GetDonorsInput) (output *GetDonorsOutput, err error) {
	output = &GetDonorsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetDonorsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Donors_Request"`
	GetDonorsRequestType
}

type GetDonorsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Donors_Response"`
	GetDonorsResponseType
}

// PutDonor (Put_Donor)
// 
// This service operation will add or update Donor Data.
func (c *Client) PutDonor(ctx context.Context, input *PutDonorInput) (output *PutDonorOutput, err error) {
	output = &PutDonorOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutDonorInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Donor_Request"`
	PutDonorRequestType
}

type PutDonorOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Donor_Response"`
	PutDonorResponseType
}

// GetDonorContributions (Get_Donor_Contributions)
// 
// This service operation will get Donor Contributions for the specified criteria.
func (c *Client) GetDonorContributions(ctx context.Context, input *GetDonorContributionsInput) (output *GetDonorContributionsOutput, err error) {
	output = &GetDonorContributionsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetDonorContributionsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Donor_Contributions_Request"`
	GetDonorContributionsRequestType
}

type GetDonorContributionsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Donor_Contributions_Response"`
	GetDonorContributionsResponseType
}

// SubmitDonorContribution (Submit_Donor_Contribution)
// 
// This service operation will add or update Donor Contribution Data.
func (c *Client) SubmitDonorContribution(ctx context.Context, input *SubmitDonorContributionInput) (output *SubmitDonorContributionOutput, err error) {
	output = &SubmitDonorContributionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitDonorContributionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Donor_Contribution_Request"`
	SubmitDonorContributionRequestType
}

type SubmitDonorContributionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Donor_Contribution_Response"`
	SubmitDonorContributionResponseType
}

// GetInvestmentStatements (Get_Investment_Statements)
// 
// This service operation will add or update Investment Statement Data.
func (c *Client) GetInvestmentStatements(ctx context.Context, input *GetInvestmentStatementsInput) (output *GetInvestmentStatementsOutput, err error) {
	output = &GetInvestmentStatementsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetInvestmentStatementsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Investment_Statements_Request"`
	GetInvestmentStatementsRequestType
}

type GetInvestmentStatementsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Investment_Statements_Response"`
	GetInvestmentStatementsResponseType
}

// SubmitInvestmentStatement (Submit_Investment_Statement)
// 
// This service operation will add or update Investment Statement Data.
func (c *Client) SubmitInvestmentStatement(ctx context.Context, input *SubmitInvestmentStatementInput) (output *SubmitInvestmentStatementOutput, err error) {
	output = &SubmitInvestmentStatementOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitInvestmentStatementInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Investment_Statement_Request"`
	SubmitInvestmentStatementRequestType
}

type SubmitInvestmentStatementOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Investment_Statement_Response"`
	SubmitInvestmentStatementResponseType
}

// GetInvestmentPoolPurchases (Get_Investment_Pool_Purchases)
// 
// This service operation will get Investment Pool Purchases for the specified criteria.
func (c *Client) GetInvestmentPoolPurchases(ctx context.Context, input *GetInvestmentPoolPurchasesInput) (output *GetInvestmentPoolPurchasesOutput, err error) {
	output = &GetInvestmentPoolPurchasesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetInvestmentPoolPurchasesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Investment_Pool_Purchases_Request"`
	GetInvestmentPoolPurchasesRequestType
}

type GetInvestmentPoolPurchasesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Investment_Pool_Purchases_Response"`
	GetInvestmentPoolPurchasesResponseType
}

// SubmitInvestmentPoolPurchase (Submit_Investment_Pool_Purchase)
// 
// This service operation will add or update Investment Pool Purchases.
func (c *Client) SubmitInvestmentPoolPurchase(ctx context.Context, input *SubmitInvestmentPoolPurchaseInput) (output *SubmitInvestmentPoolPurchaseOutput, err error) {
	output = &SubmitInvestmentPoolPurchaseOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitInvestmentPoolPurchaseInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Investment_Pool_Purchase_Request"`
	SubmitInvestmentPoolPurchaseRequestType
}

type SubmitInvestmentPoolPurchaseOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Investment_Pool_Purchase_Response"`
	SubmitInvestmentPoolPurchaseResponseType
}

// GetInvestmentPoolTransfers (Get_Investment_Pool_Transfers)
// 
// This service operation will get Investment Pool Transfers for the specified criteria.
func (c *Client) GetInvestmentPoolTransfers(ctx context.Context, input *GetInvestmentPoolTransfersInput) (output *GetInvestmentPoolTransfersOutput, err error) {
	output = &GetInvestmentPoolTransfersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetInvestmentPoolTransfersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Investment_Pool_Transfers_Request"`
	GetInvestmentPoolTransfersRequestType
}

type GetInvestmentPoolTransfersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Investment_Pool_Transfers_Response"`
	GetInvestmentPoolTransfersResponseType
}

// GetAdHocPayees (Get_Ad_Hoc_Payees)
// 
// This service operation will get Ad Hoc Payees for the specified criteria.
func (c *Client) GetAdHocPayees(ctx context.Context, input *GetAdHocPayeesInput) (output *GetAdHocPayeesOutput, err error) {
	output = &GetAdHocPayeesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAdHocPayeesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Ad_Hoc_Payees_Request"`
	GetAdHocPayeesRequestType
}

type GetAdHocPayeesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Ad_Hoc_Payees_Response"`
	GetAdHocPayeesResponseType
}

// PutAdHocPayee (Put_Ad_Hoc_Payee)
// 
// This service operation will add or update Ad Hoc Payees. Ad Hoc Payee data includes Ad Hoc Payee Reference ID, Ad Hoc Payee Name, Tax Authority Form Type, TIN Type Reference, Tax ID, Address Data, Bank Data, and Payee Alternate Name Data.
func (c *Client) PutAdHocPayee(ctx context.Context, input *PutAdHocPayeeInput) (output *PutAdHocPayeeOutput, err error) {
	output = &PutAdHocPayeeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutAdHocPayeeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Ad_Hoc_Payee_Request"`
	PutAdHocPayeeRequestType
}

type PutAdHocPayeeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Ad_Hoc_Payee_Response"`
	PutAdHocPayeeResponseType
}

// SubmitInvestmentPoolTransfer (Submit_Investment_Pool_Transfer)
// 
// This service operation will submit Investment Pool Transfer.
func (c *Client) SubmitInvestmentPoolTransfer(ctx context.Context, input *SubmitInvestmentPoolTransferInput) (output *SubmitInvestmentPoolTransferOutput, err error) {
	output = &SubmitInvestmentPoolTransferOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitInvestmentPoolTransferInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Investment_Pool_Transfer_Request"`
	SubmitInvestmentPoolTransferRequestType
}

type SubmitInvestmentPoolTransferOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Investment_Pool_Transfer_Response"`
	SubmitInvestmentPoolTransferResponseType
}

// SubmitInvestmentPoolSale (Submit_Investment_Pool_Sale)
// 
// This service if for Selling Investment Pool Units
func (c *Client) SubmitInvestmentPoolSale(ctx context.Context, input *SubmitInvestmentPoolSaleInput) (output *SubmitInvestmentPoolSaleOutput, err error) {
	output = &SubmitInvestmentPoolSaleOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitInvestmentPoolSaleInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Investment_Pool_Sale_Request"`
	SubmitInvestmentPoolSaleRequestType
}

type SubmitInvestmentPoolSaleOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Investment_Pool_Sale_Response"`
	SubmitInvestmentPoolSaleResponseType
}

// GetInvestmentPoolSales (Get_Investment_Pool_Sales)
// 
// This service is to get the Investment Pool Sale Unit details
func (c *Client) GetInvestmentPoolSales(ctx context.Context, input *GetInvestmentPoolSalesInput) (output *GetInvestmentPoolSalesOutput, err error) {
	output = &GetInvestmentPoolSalesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetInvestmentPoolSalesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Investment_Pool_Sales_Request"`
	GetInvestmentPoolSalesRequestType
}

type GetInvestmentPoolSalesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Investment_Pool_Sales_Response"`
	GetInvestmentPoolSalesResponseType
}

// SubmitInvestmentPoolAdjustment (Submit_Investment_Pool_Adjustment)
// 
// This service operation will submit Investment Pool Adjustment.
func (c *Client) SubmitInvestmentPoolAdjustment(ctx context.Context, input *SubmitInvestmentPoolAdjustmentInput) (output *SubmitInvestmentPoolAdjustmentOutput, err error) {
	output = &SubmitInvestmentPoolAdjustmentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitInvestmentPoolAdjustmentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Investment_Pool_Adjustment_Request"`
	SubmitInvestmentPoolAdjustmentRequestType
}

type SubmitInvestmentPoolAdjustmentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Investment_Pool_Adjustment_Response"`
	SubmitInvestmentPoolAdjustmentResponseType
}

// GetInvestmentPoolAdjustments (Get_Investment_Pool_Adjustments)
// 
// This service operation will get Investment Pool Adjustments for the specified criteria.
func (c *Client) GetInvestmentPoolAdjustments(ctx context.Context, input *GetInvestmentPoolAdjustmentsInput) (output *GetInvestmentPoolAdjustmentsOutput, err error) {
	output = &GetInvestmentPoolAdjustmentsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetInvestmentPoolAdjustmentsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Investment_Pool_Adjustments_Request"`
	GetInvestmentPoolAdjustmentsRequestType
}

type GetInvestmentPoolAdjustmentsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Investment_Pool_Adjustments_Response"`
	GetInvestmentPoolAdjustmentsResponseType
}

// GetBankAccountTransfers (Get_Bank_Account_Transfers)
// 
// This service operation will get Bank Account Transfers for the specified criteria.
func (c *Client) GetBankAccountTransfers(ctx context.Context, input *GetBankAccountTransfersInput) (output *GetBankAccountTransfersOutput, err error) {
	output = &GetBankAccountTransfersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBankAccountTransfersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Account_Transfers_Request"`
	GetBankAccountTransfersRequestType
}

type GetBankAccountTransfersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Account_Transfers_Response"`
	GetBankAccountTransfersResponseType
}

// SubmitBankAccountTransfer (Submit_Bank_Account_Transfer)
// 
// Creates and submit a new Bank Account Transfer with the information supplied in the request
// This service operation will submit Bank Account Transfer
func (c *Client) SubmitBankAccountTransfer(ctx context.Context, input *SubmitBankAccountTransferInput) (output *SubmitBankAccountTransferOutput, err error) {
	output = &SubmitBankAccountTransferOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitBankAccountTransferInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Bank_Account_Transfer_Request"`
	SubmitBankAccountTransferRequestType
}

type SubmitBankAccountTransferOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Bank_Account_Transfer_Response"`
	SubmitBankAccountTransferResponseType
}

// GetBankStatementCustomCodeSets (Get_Bank_Statement_Custom_Code_Sets)
// 
// This service operation will get Bank Statement Custom Code Sets for the specified criteria. Bank Statement Custom Code Set data includes the Custom Code Set ID, Custom Code Set Name, Associated Bank Statement Format, Custom Code Set Description, Comment and Custom Balance Type Codes and Custom Transaction Type Codes. The Custom Balance Type Code data includes  Balance/Summary indicator, Type Code, Debit/Credit Flag, Type Description, Associated Bank Statement Type Category and whether its a opening balance or a closing balance if its a balance type code. The Custom Transaction Type Code data includes Type Code, Debit/Credit Flag, Type Description and Associated Bank Statement Type Category. The request criteria can be for a single transaction based on Reference, or all transaction can be retrieved if no criteria is specified.
func (c *Client) GetBankStatementCustomCodeSets(ctx context.Context, input *GetBankStatementCustomCodeSetsInput) (output *GetBankStatementCustomCodeSetsOutput, err error) {
	output = &GetBankStatementCustomCodeSetsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBankStatementCustomCodeSetsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Statement_Custom_Code_Sets_Request"`
	GetBankStatementCustomCodeSetsRequestType
}

type GetBankStatementCustomCodeSetsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Statement_Custom_Code_Sets_Response"`
	GetBankStatementCustomCodeSetsResponseType
}

// PutBankStatementCustomCodeSet (Put_Bank_Statement_Custom_Code_Set)
// 
// This service operation will add or update Bank Statement Custom Code Sets. Bank Statement Custom Code Set data includes the Custom Code Set ID, Custom Code Set Name, Associated Bank Statement Format, Custom Code Set Description, Comment and Custom Balance Type Codes and Custom Transaction Type Codes. The Custom Balance Type Code data includes Balance/Summary indicator, Type Code, Debit/Credit Flag, Type Description, Associated Bank Statement Type Category and whether its a opening balance or a closing balance if its a balance type code. The Custom Transaction Type Code data includes Type Code, Debit/Credit Flag, Type Description and Associated Bank Statement Type Category.
func (c *Client) PutBankStatementCustomCodeSet(ctx context.Context, input *PutBankStatementCustomCodeSetInput) (output *PutBankStatementCustomCodeSetOutput, err error) {
	output = &PutBankStatementCustomCodeSetOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBankStatementCustomCodeSetInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Statement_Custom_Code_Set_Request"`
	PutBankStatementCustomCodeSetRequestType
}

type PutBankStatementCustomCodeSetOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Statement_Custom_Code_Set_Response"`
	PutBankStatementCustomCodeSetResponseType
}

// GetIntradayBankStatements (Get_Intraday_Bank_Statements)
// 
// This service operation will get all Intraday Bank Statement for the specified criteria. The response will include specific information about Intraday Bank Statements.
func (c *Client) GetIntradayBankStatements(ctx context.Context, input *GetIntradayBankStatementsInput) (output *GetIntradayBankStatementsOutput, err error) {
	output = &GetIntradayBankStatementsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetIntradayBankStatementsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Intraday_Bank_Statements_Request"`
	GetIntradayBankStatementsRequestType
}

type GetIntradayBankStatementsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Intraday_Bank_Statements_Response"`
	GetIntradayBankStatementsResponseType
}

// PutInvestmentPoolCatchupPayoutCriteria (Put_Investment_Pool_Catch-up_Payout_Criteria)
// 
// Web service to submit request to initiate the Run Catch-up Payout with the information supplied in the request.
func (c *Client) PutInvestmentPoolCatchupPayoutCriteria(ctx context.Context, input *PutInvestmentPoolCatchupPayoutCriteriaInput) (output *PutInvestmentPoolCatchupPayoutCriteriaOutput, err error) {
	output = &PutInvestmentPoolCatchupPayoutCriteriaOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutInvestmentPoolCatchupPayoutCriteriaInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Investment_Pool_Catch-up_Payout_Criteria_Request"`
	PutInvestmentPoolCatchupPayoutCriteriaRequestType
}

type PutInvestmentPoolCatchupPayoutCriteriaOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Investment_Pool_Catch-up_Payout_Criteria_Response"`
	PutInvestmentPoolCatchupPayoutCriteriaResponseType
}

// GetBankFeeStatementFiles (Get_Bank_Fee_Statement_Files)
// 
// This service operation will get all Bank Fee Statement Files for the specified criteria. The response will include specific information about the Bank Fee Statement Files such as the Bank Fee Statement File References, File Creation Date, File Identification Number, Receiver Identification, Sender Identification, and Bank Fee Statement References.
func (c *Client) GetBankFeeStatementFiles(ctx context.Context, input *GetBankFeeStatementFilesInput) (output *GetBankFeeStatementFilesOutput, err error) {
	output = &GetBankFeeStatementFilesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBankFeeStatementFilesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Fee_Statement_Files_Request"`
	GetBankFeeStatementFilesRequestType
}

type GetBankFeeStatementFilesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Fee_Statement_Files_Response"`
	GetBankFeeStatementFilesResponseType
}

// PutBankFeeStatementFile (Put_Bank_Fee_Statement_File)
// 
// This service operation will add or update Bank Fee Statement Files. Each file will load one or more Bank Fee Statements.
func (c *Client) PutBankFeeStatementFile(ctx context.Context, input *PutBankFeeStatementFileInput) (output *PutBankFeeStatementFileOutput, err error) {
	output = &PutBankFeeStatementFileOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBankFeeStatementFileInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Fee_Statement_File_Request"`
	PutBankFeeStatementFileRequestType
}

type PutBankFeeStatementFileOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Fee_Statement_File_Response"`
	PutBankFeeStatementFileResponseType
}

// GetBankFeeStatements (Get_Bank_Fee_Statements)
// 
// Retrieves the Bank Fee Statements and filters by request criteria if any are given.
func (c *Client) GetBankFeeStatements(ctx context.Context, input *GetBankFeeStatementsInput) (output *GetBankFeeStatementsOutput, err error) {
	output = &GetBankFeeStatementsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBankFeeStatementsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Fee_Statements_Request"`
	GetBankFeeStatementsRequestType
}

type GetBankFeeStatementsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Fee_Statements_Response"`
	GetBankFeeStatementsResponseType
}

// CancelInvestmentPoolTransaction (Cancel_Investment_Pool_Transaction)
// 
// This service operation will cancel an Investment Pool Transaction
func (c *Client) CancelInvestmentPoolTransaction(ctx context.Context, input *CancelInvestmentPoolTransactionInput) (output *CancelInvestmentPoolTransactionOutput, err error) {
	output = &CancelInvestmentPoolTransactionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CancelInvestmentPoolTransactionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Cancel_Investment_Pool_Transaction_Request"`
	CancelInvestmentPoolTransactionRequestType
}

type CancelInvestmentPoolTransactionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Cancel_Investment_Pool_Transaction_Response"`
	CancelInvestmentPoolTransactionResponseType
}

// GetBankFeeServiceCodes (Get_Bank_Fee_Service_Codes)
// 
// Get Bank Fee Service Codes
func (c *Client) GetBankFeeServiceCodes(ctx context.Context, input *GetBankFeeServiceCodesInput) (output *GetBankFeeServiceCodesOutput, err error) {
	output = &GetBankFeeServiceCodesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBankFeeServiceCodesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Fee_Service_Codes_Request"`
	GetBankFeeServiceCodesRequestType
}

type GetBankFeeServiceCodesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Fee_Service_Codes_Response"`
	GetBankFeeServiceCodesResponseType
}

// PutBankFeeServiceCode (Put_Bank_Fee_Service_Code)
// 
// Put Bank Fee Service Code
func (c *Client) PutBankFeeServiceCode(ctx context.Context, input *PutBankFeeServiceCodeInput) (output *PutBankFeeServiceCodeOutput, err error) {
	output = &PutBankFeeServiceCodeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBankFeeServiceCodeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Fee_Service_Code_Request"`
	PutBankFeeServiceCodeRequestType
}

type PutBankFeeServiceCodeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Fee_Service_Code_Response"`
	PutBankFeeServiceCodeResponseType
}

// GetSignatureMethods (Get_Signature_Methods)
// 
// Retrieves the Signature Methods used within Bank Account Signatories.
func (c *Client) GetSignatureMethods(ctx context.Context, input *GetSignatureMethodsInput) (output *GetSignatureMethodsOutput, err error) {
	output = &GetSignatureMethodsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetSignatureMethodsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Signature_Methods_Request"`
	GetSignatureMethodsRequestType
}

type GetSignatureMethodsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Signature_Methods_Response"`
	GetSignatureMethodsResponseType
}

// GetAuthorityTypes (Get_Authority_Types)
// 
// Retrieves the Authority Types used within Bank Account Signatories.
func (c *Client) GetAuthorityTypes(ctx context.Context, input *GetAuthorityTypesInput) (output *GetAuthorityTypesOutput, err error) {
	output = &GetAuthorityTypesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAuthorityTypesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Authority_Types_Request"`
	GetAuthorityTypesRequestType
}

type GetAuthorityTypesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Authority_Types_Response"`
	GetAuthorityTypesResponseType
}

// PutSignatureMethod (Put_Signature_Method)
// 
// Adds or updates Signature Methods.
func (c *Client) PutSignatureMethod(ctx context.Context, input *PutSignatureMethodInput) (output *PutSignatureMethodOutput, err error) {
	output = &PutSignatureMethodOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutSignatureMethodInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Signature_Method_Request"`
	PutSignatureMethodRequestType
}

type PutSignatureMethodOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Signature_Method_Response"`
	PutSignatureMethodResponseType
}

// PutAuthorityType (Put_Authority_Type)
// 
// Adds or updates Authority Types.
func (c *Client) PutAuthorityType(ctx context.Context, input *PutAuthorityTypeInput) (output *PutAuthorityTypeOutput, err error) {
	output = &PutAuthorityTypeOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutAuthorityTypeInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Authority_Type_Request"`
	PutAuthorityTypeRequestType
}

type PutAuthorityTypeOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Authority_Type_Response"`
	PutAuthorityTypeResponseType
}

// CorrectAdHocPayment (Correct_Ad_Hoc_Payment)
// 
// Correct Ad Hoc Payment Web Service for small changes after the payment has been settled.
func (c *Client) CorrectAdHocPayment(ctx context.Context, input *CorrectAdHocPaymentInput) (output *CorrectAdHocPaymentOutput, err error) {
	output = &CorrectAdHocPaymentOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CorrectAdHocPaymentInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Correct_Ad_Hoc_Payment_Request"`
	CorrectAdHocPaymentRequestType
}

type CorrectAdHocPaymentOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Correct_Ad_Hoc_Payment_Response"`
	CorrectAdHocPaymentResponseType
}

// PutSigner (Put_Signer)
// 
// Adds Signers.
func (c *Client) PutSigner(ctx context.Context, input *PutSignerInput) (output *PutSignerOutput, err error) {
	output = &PutSignerOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutSignerInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Signer_Request"`
	PutSignerRequestType
}

type PutSignerOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Signer_Response"`
	PutSignerResponseType
}

// GetSigners (Get_Signers)
// 
// Retrieves the Signers used within Bank Account Signatories.
func (c *Client) GetSigners(ctx context.Context, input *GetSignersInput) (output *GetSignersOutput, err error) {
	output = &GetSignersOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetSignersInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Signers_Request"`
	GetSignersRequestType
}

type GetSignersOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Signers_Response"`
	GetSignersResponseType
}

// SubmitBankAccountSignatory (Submit_Bank_Account_Signatory)
// 
// Adds, updates, or removes Bank Account Signatories.
func (c *Client) SubmitBankAccountSignatory(ctx context.Context, input *SubmitBankAccountSignatoryInput) (output *SubmitBankAccountSignatoryOutput, err error) {
	output = &SubmitBankAccountSignatoryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitBankAccountSignatoryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Bank_Account_Signatory_Request"`
	SubmitBankAccountSignatoryRequestType
}

type SubmitBankAccountSignatoryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Bank_Account_Signatory_Response"`
	SubmitBankAccountSignatoryResponseType
}

// GetBankAccountSignatories (Get_Bank_Account_Signatories)
// 
// Retrieves the Bank Account and its Signatories.
func (c *Client) GetBankAccountSignatories(ctx context.Context, input *GetBankAccountSignatoriesInput) (output *GetBankAccountSignatoriesOutput, err error) {
	output = &GetBankAccountSignatoriesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBankAccountSignatoriesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Account_Signatories_Request"`
	GetBankAccountSignatoriesRequestType
}

type GetBankAccountSignatoriesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Account_Signatories_Response"`
	GetBankAccountSignatoriesResponseType
}

// RecalculateBankStatementBalance (Recalculate_Bank_Statement_Balance)
// 
// Webservice used  to recalculate beginning balance. This should be called  if default beginning balance is selected for the bank account and multiple bank statements are loaded as single file.
func (c *Client) RecalculateBankStatementBalance(ctx context.Context, input *RecalculateBankStatementBalanceInput) (output *RecalculateBankStatementBalanceOutput, err error) {
	output = &RecalculateBankStatementBalanceOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type RecalculateBankStatementBalanceInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Recalculate_Bank_Statement_Balance_Request"`
	RecalculateBankStatementBalanceRequestType
}

type RecalculateBankStatementBalanceOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Recalculate_Bank_Statement_Balance_Response"`
	RecalculateBankStatementBalanceResponseType
}

// CancelDonorContribution (Cancel_Donor_Contribution)
// 
// This Service Operation will cancel an eligible Donor Contribution.
func (c *Client) CancelDonorContribution(ctx context.Context, input *CancelDonorContributionInput) (output *CancelDonorContributionOutput, err error) {
	output = &CancelDonorContributionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type CancelDonorContributionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Cancel_Donor_Contribution_Request"`
	CancelDonorContributionRequestType
}

type CancelDonorContributionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Cancel_Donor_Contribution_Response"`
	CancelDonorContributionResponseType
}

// GetExternalCashActivities (Get_External_Cash_Activities)
// 
// This operation will get External Cash Activities for the specified criteria. Includes details such as Company, Bank Account, Currency, Amount, Expiration Date, Worktag &amp; Counterparty etc.
func (c *Client) GetExternalCashActivities(ctx context.Context, input *GetExternalCashActivitiesInput) (output *GetExternalCashActivitiesOutput, err error) {
	output = &GetExternalCashActivitiesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetExternalCashActivitiesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_External_Cash_Activities_Request"`
	GetExternalCashActivitiesRequestType
}

type GetExternalCashActivitiesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_External_Cash_Activities_Response"`
	GetExternalCashActivitiesResponseType
}

// GetCashActivityCategories (Get_Cash_Activity_Categories)
// 
// Retrieves the Cash Activity Categories.
func (c *Client) GetCashActivityCategories(ctx context.Context, input *GetCashActivityCategoriesInput) (output *GetCashActivityCategoriesOutput, err error) {
	output = &GetCashActivityCategoriesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCashActivityCategoriesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Cash_Activity_Categories_Request"`
	GetCashActivityCategoriesRequestType
}

type GetCashActivityCategoriesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Cash_Activity_Categories_Response"`
	GetCashActivityCategoriesResponseType
}

// PutCashActivityCategory (Put_Cash_Activity_Category)
// 
// Adds or updates Cash Activity Categories.
func (c *Client) PutCashActivityCategory(ctx context.Context, input *PutCashActivityCategoryInput) (output *PutCashActivityCategoryOutput, err error) {
	output = &PutCashActivityCategoryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCashActivityCategoryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Cash_Activity_Category_Request"`
	PutCashActivityCategoryRequestType
}

type PutCashActivityCategoryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Cash_Activity_Category_Response"`
	PutCashActivityCategoryResponseType
}

// GetBankAccountTransfersforSettlement (Get_Bank_Account_Transfers_for_Settlement)
// 
// This service operation will get Bank Account Transfers for Settlement that satisfied the specified criteria
func (c *Client) GetBankAccountTransfersforSettlement(ctx context.Context, input *GetBankAccountTransfersforSettlementInput) (output *GetBankAccountTransfersforSettlementOutput, err error) {
	output = &GetBankAccountTransfersforSettlementOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBankAccountTransfersforSettlementInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Account_Transfers_for_Settlement_Request"`
	GetBankAccountTransfersforSettlementRequestType
}

type GetBankAccountTransfersforSettlementOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Account_Transfers_for_Settlement_Response"`
	GetBankAccountTransfersforSettlementResponseType
}

// SubmitBankAccountTransferforSettlement (Submit_Bank_Account_Transfer_for_Settlement)
// 
// Create or update a Bank Account Transfer for Settlement with the information supplied in the request.  The service operation will submit the request to the Bank Account Transfer for Settlement business process
func (c *Client) SubmitBankAccountTransferforSettlement(ctx context.Context, input *SubmitBankAccountTransferforSettlementInput) (output *SubmitBankAccountTransferforSettlementOutput, err error) {
	output = &SubmitBankAccountTransferforSettlementOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitBankAccountTransferforSettlementInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Bank_Account_Transfer_for_Settlement_Request"`
	SubmitBankAccountTransferforSettlementRequestType
}

type SubmitBankAccountTransferforSettlementOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Bank_Account_Transfer_for_Settlement_Response"`
	SubmitBankAccountTransferforSettlementResponseType
}

// PutBankAccountSignatory (Put_Bank_Account_Signatory)
// 
// Same as Submit Bank Account Signatory, but with Auto Complete enabled for mass update via iLoad.
func (c *Client) PutBankAccountSignatory(ctx context.Context, input *PutBankAccountSignatoryInput) (output *PutBankAccountSignatoryOutput, err error) {
	output = &PutBankAccountSignatoryOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBankAccountSignatoryInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Account_Signatory_Request"`
	PutBankAccountSignatoryRequestType
}

type PutBankAccountSignatoryOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Account_Signatory_Response"`
	PutBankAccountSignatoryResponseType
}

// GetBankFeeServiceContracts (Get_Bank_Fee_Service_Contracts)
// 
// Get Bank Fee Service Contracts
func (c *Client) GetBankFeeServiceContracts(ctx context.Context, input *GetBankFeeServiceContractsInput) (output *GetBankFeeServiceContractsOutput, err error) {
	output = &GetBankFeeServiceContractsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBankFeeServiceContractsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Fee_Service_Contracts_Request"`
	GetBankFeeServiceContractsRequestType
}

type GetBankFeeServiceContractsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Fee_Service_Contracts_Response"`
	GetBankFeeServiceContractsResponseType
}

// PutBankFeeServiceContract (Put_Bank_Fee_Service_Contract)
// 
// Creates a new Bank Fee Service Contract with the information supplied in the request.
func (c *Client) PutBankFeeServiceContract(ctx context.Context, input *PutBankFeeServiceContractInput) (output *PutBankFeeServiceContractOutput, err error) {
	output = &PutBankFeeServiceContractOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutBankFeeServiceContractInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Fee_Service_Contract_Request"`
	PutBankFeeServiceContractRequestType
}

type PutBankFeeServiceContractOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Bank_Fee_Service_Contract_Response"`
	PutBankFeeServiceContractResponseType
}

// GetBankAccountTransferTemplates (Get_Bank_Account_Transfer_Templates)
// 
// This operation will get Bank Account Transfer Templates for the specified criteria.
func (c *Client) GetBankAccountTransferTemplates(ctx context.Context, input *GetBankAccountTransferTemplatesInput) (output *GetBankAccountTransferTemplatesOutput, err error) {
	output = &GetBankAccountTransferTemplatesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBankAccountTransferTemplatesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Account_Transfer_Templates_Request"`
	GetBankAccountTransferTemplatesRequestType
}

type GetBankAccountTransferTemplatesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Account_Transfer_Templates_Response"`
	GetBankAccountTransferTemplatesResponseType
}

// GetCashPools (Get_Cash_Pools)
// 
// Retrieves the referenced Cash Pool. If there is no referenced Cash Pool, retrieves all Cash Pools.
func (c *Client) GetCashPools(ctx context.Context, input *GetCashPoolsInput) (output *GetCashPoolsOutput, err error) {
	output = &GetCashPoolsOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetCashPoolsInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Cash_Pools_Request"`
	GetCashPoolsRequestType
}

type GetCashPoolsOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Cash_Pools_Response"`
	GetCashPoolsResponseType
}

// PutCashPool (Put_Cash_Pool)
// 
// Adds or updates Cash Pools.
func (c *Client) PutCashPool(ctx context.Context, input *PutCashPoolInput) (output *PutCashPoolOutput, err error) {
	output = &PutCashPoolOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type PutCashPoolInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Cash_Pool_Request"`
	PutCashPoolRequestType
}

type PutCashPoolOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Cash_Pool_Response"`
	PutCashPoolResponseType
}

// SubmitBankAccountTransferTemplate (Submit_Bank_Account_Transfer_Template)
// 
// Creates and submit a new Bank Account Transfer Template with the information supplied in the request.
// This service operation will submit Bank Account Transfer Template
func (c *Client) SubmitBankAccountTransferTemplate(ctx context.Context, input *SubmitBankAccountTransferTemplateInput) (output *SubmitBankAccountTransferTemplateOutput, err error) {
	output = &SubmitBankAccountTransferTemplateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitBankAccountTransferTemplateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Bank_Account_Transfer_Template_Request"`
	SubmitBankAccountTransferTemplateRequestType
}

type SubmitBankAccountTransferTemplateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Bank_Account_Transfer_Template_Response"`
	SubmitBankAccountTransferTemplateResponseType
}

// SubmitAdHocPaymentTemplate (Submit_Ad_Hoc_Payment_Template)
// 
// This operation will add or update an Ad Hoc Payment Template and submit to the Ad Hoc Payment Template business process.
func (c *Client) SubmitAdHocPaymentTemplate(ctx context.Context, input *SubmitAdHocPaymentTemplateInput) (output *SubmitAdHocPaymentTemplateOutput, err error) {
	output = &SubmitAdHocPaymentTemplateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitAdHocPaymentTemplateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Ad_Hoc_Payment_Template_Request"`
	SubmitAdHocPaymentTemplateRequestType
}

type SubmitAdHocPaymentTemplateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Ad_Hoc_Payment_Template_Response"`
	SubmitAdHocPaymentTemplateResponseType
}

// GetAdHocPaymentTemplates (Get_Ad_Hoc_Payment_Templates)
// 
// This operation will get Ad Hoc Payment Templates for the specified criteria.
func (c *Client) GetAdHocPaymentTemplates(ctx context.Context, input *GetAdHocPaymentTemplatesInput) (output *GetAdHocPaymentTemplatesOutput, err error) {
	output = &GetAdHocPaymentTemplatesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetAdHocPaymentTemplatesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Ad_Hoc_Payment_Templates_Request"`
	GetAdHocPaymentTemplatesRequestType
}

type GetAdHocPaymentTemplatesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Ad_Hoc_Payment_Templates_Response"`
	GetAdHocPaymentTemplatesResponseType
}

// SubmitAdHocPaymentfromTemplate (Submit_Ad_Hoc_Payment_from_Template)
// 
// This service operation will create an Ad Hoc Payment using an Ad Hoc Payment Template and submit to the Ad Hoc Payment business process.
func (c *Client) SubmitAdHocPaymentfromTemplate(ctx context.Context, input *SubmitAdHocPaymentfromTemplateInput) (output *SubmitAdHocPaymentfromTemplateOutput, err error) {
	output = &SubmitAdHocPaymentfromTemplateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitAdHocPaymentfromTemplateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Ad_Hoc_Payment_from_Template_Request"`
	SubmitAdHocPaymentfromTemplateRequestType
}

type SubmitAdHocPaymentfromTemplateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Ad_Hoc_Payment_from_Template_Response"`
	SubmitAdHocPaymentfromTemplateResponseType
}

// SubmitBankAccount (Submit_Bank_Account)
// 
// This service operation will add or update Bank Accounts.  Bank Account data includes the Bank Account ID, Account Name, Financial Institution, Bank Branch, Account Closed flag, Routing Transit or Institution Number, Financial Account Number, Bank Identifier Code, IBAN, Payment Type, Check Print Layout, Advanced Mode flag, Automatic Reconciliation Type, Reconciliation Rule Set, First Notice Reconciliation, First Notice Rule Set, Bank Electronic Payments flag, Used by Cash flag, Used by Customer Payment flag, Used by Expense Payments flag, Used by Payroll flag, Used by Supplier Payments flag, Used by Intercompany payments flag, Last Check Number Used, Enable Positive Pay flag, and Lockbox data.
func (c *Client) SubmitBankAccount(ctx context.Context, input *SubmitBankAccountInput) (output *SubmitBankAccountOutput, err error) {
	output = &SubmitBankAccountOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitBankAccountInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Bank_Account_Request"`
	SubmitBankAccountRequestType
}

type SubmitBankAccountOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Bank_Account_Response"`
	SubmitBankAccountResponseType
}

// SubmitBankAccountTransferforSettlementTemplate (Submit_Bank_Account_Transfer_for_Settlement_Template)
// 
// This operation will add or update a Bank Account Transfer for Settlement Template and submit to the Bank Account Transfer for Settlement Template business process.
func (c *Client) SubmitBankAccountTransferforSettlementTemplate(ctx context.Context, input *SubmitBankAccountTransferforSettlementTemplateInput) (output *SubmitBankAccountTransferforSettlementTemplateOutput, err error) {
	output = &SubmitBankAccountTransferforSettlementTemplateOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type SubmitBankAccountTransferforSettlementTemplateInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Bank_Account_Transfer_for_Settlement_Template_Request"`
	SubmitBankAccountTransferforSettlementTemplateRequestType
}

type SubmitBankAccountTransferforSettlementTemplateOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Submit_Bank_Account_Transfer_for_Settlement_Template_Response"`
	SubmitBankAccountTransferforSettlementTemplateResponseType
}

// GetBankAccountTransferforSettlementTemplates (Get_Bank_Account_Transfer_for_Settlement_Templates)
// 
// This operation will get Bank Account Transfer for Settlement Templates for the specified criteria.
func (c *Client) GetBankAccountTransferforSettlementTemplates(ctx context.Context, input *GetBankAccountTransferforSettlementTemplatesInput) (output *GetBankAccountTransferforSettlementTemplatesOutput, err error) {
	output = &GetBankAccountTransferforSettlementTemplatesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type GetBankAccountTransferforSettlementTemplatesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Account_Transfer_for_Settlement_Templates_Request"`
	GetBankAccountTransferforSettlementTemplatesRequestType
}

type GetBankAccountTransferforSettlementTemplatesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Get_Bank_Account_Transfer_for_Settlement_Templates_Response"`
	GetBankAccountTransferforSettlementTemplatesResponseType
}

// ImportBankStatement (Import_Bank_Statement)
// 
// This service operation will add or update Bank Statement Data and attach it to an existing Bank Statement File.
func (c *Client) ImportBankStatement(ctx context.Context, input *ImportBankStatementInput) (output *ImportBankStatementOutput, err error) {
	output = &ImportBankStatementOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportBankStatementInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Bank_Statement_Request"`
	ImportBankStatementRequestType
}

type ImportBankStatementOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportAdhocBankTransaction (Import_Ad_hoc_Bank_Transaction)
// 
// This service operation will add or update Ad hoc Bank Transactions and submit to the Ad hoc Bank Transaction business process. Ad hoc Bank Transaction data includes Ad hoc Bank Transaction ID, Submit for Approval flag, Locked in Workday flag, Transaction Date, Transaction Memo, Company, Currency, Bank Account, Transaction Amount, Deposit flag, Withdraw flag, Purpose, Transaction Reference Text and Transaction Line data. Transaction Lines data includes Inter Company affiliate, Spend Category, Revenue Category, Ledger Account, Line Amount, Line Memo, and Worktags.
// 
// Spend Category and Resource Category are synonymous and refer to the same business object.
func (c *Client) ImportAdhocBankTransaction(ctx context.Context, input *ImportAdhocBankTransactionInput) (output *ImportAdhocBankTransactionOutput, err error) {
	output = &ImportAdhocBankTransactionOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportAdhocBankTransactionInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Ad_hoc_Bank_Transaction_Request"`
	ImportAdhocBankTransactionRequestType
}

type ImportAdhocBankTransactionOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportIntradayBankStatement (Import_Intraday_Bank_Statement)
// 
// This will be used to Create or Edit Intraday Bank Statements.
func (c *Client) ImportIntradayBankStatement(ctx context.Context, input *ImportIntradayBankStatementInput) (output *ImportIntradayBankStatementOutput, err error) {
	output = &ImportIntradayBankStatementOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportIntradayBankStatementInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Intraday_Bank_Statement_Request"`
	ImportIntradayBankStatementRequestType
}

type ImportIntradayBankStatementOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportBankFeeStatement (Import_Bank_Fee_Statement)
// 
// This service operation will add or update Bank Fee Statement Data and attach it to an existing Bank Fee Statement File.
func (c *Client) ImportBankFeeStatement(ctx context.Context, input *ImportBankFeeStatementInput) (output *ImportBankFeeStatementOutput, err error) {
	output = &ImportBankFeeStatementOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportBankFeeStatementInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_Bank_Fee_Statement_Request"`
	ImportBankFeeStatementRequestType
}

type ImportBankFeeStatementOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

// ImportExternalCashActivities (Import_External_Cash_Activities)
// 
// This service operation will add or update External Cash Activity Data. The External Cash Activity data includes Company, Bank Account, Date, Currency, Amount, Debit/Credit, Counterparty, Worktags, Expiration Date and Description.
func (c *Client) ImportExternalCashActivities(ctx context.Context, input *ImportExternalCashActivitiesInput) (output *ImportExternalCashActivitiesOutput, err error) {
	output = &ImportExternalCashActivitiesOutput{}
	err = c.Do(ctx, ServiceName, CurrentVersion, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

type ImportExternalCashActivitiesInput struct {
	XMLName string `xml:"urn:com.workday/bsvc Import_External_Cash_Activity_Request"`
	ImportExternalCashActivityRequestType
}

type ImportExternalCashActivitiesOutput struct {
	XMLName string `xml:"urn:com.workday/bsvc Put_Import_Process_Response"`
	PutImportProcessResponseType
}

