package cashmgmt

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// Contains a unique identifier for an instance of an object.
type AbstractCreditCardTransactionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AbstractCreditCardTransactionObjectType struct {
	ID         []AbstractCreditCardTransactionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AccountingTreatmentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AccountingTreatmentObjectType struct {
	ID         []AccountingTreatmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AccountingWorktagObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AccountingWorktagObjectType struct {
	ID         []AccountingWorktagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AdHocBankTransactionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AdHocBankTransactionObjectType struct {
	ID         []AdHocBankTransactionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AdHocBankTransactionPurposeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AdHocBankTransactionPurposeObjectType struct {
	ID         []AdHocBankTransactionPurposeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Customer wrapper element that contains the Ad Hoc Payee data elements that are included in the response
type AdHocPayeeDataType struct {
	AdHocPayeeID                  string                                `xml:"urn:com.workday/bsvc Ad_Hoc_Payee_ID,omitempty"`
	AdHocPayeeName                string                                `xml:"urn:com.workday/bsvc Ad_Hoc_Payee_Name,omitempty"`
	Inactive                      *bool                                 `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	SingleUse                     *bool                                 `xml:"urn:com.workday/bsvc Single_Use,omitempty"`
	TaxAuthorityFormTypeReference []TaxAuthorityFormTypeObjectType      `xml:"urn:com.workday/bsvc Tax_Authority_Form_Type_Reference,omitempty"`
	TINTypeReference              *TaxpayerIDNumberTypeObjectType       `xml:"urn:com.workday/bsvc TIN_Type_Reference,omitempty"`
	TaxID                         string                                `xml:"urn:com.workday/bsvc Tax_ID,omitempty"`
	TaxDocumentDate               *time.Time                            `xml:"urn:com.workday/bsvc Tax_Document_Date,omitempty"`
	AddressData                   []AddressInformationDataType          `xml:"urn:com.workday/bsvc Address_Data,omitempty"`
	BankData                      []SettlementAccountWWSDataType        `xml:"urn:com.workday/bsvc Bank_Data,omitempty"`
	PayeeAlternateNameData        []BusinessEntityAlternateNameDataType `xml:"urn:com.workday/bsvc Payee_Alternate_Name_Data,omitempty"`
}

func (t *AdHocPayeeDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AdHocPayeeDataType
	var layout struct {
		*T
		TaxDocumentDate *xsdDate `xml:"urn:com.workday/bsvc Tax_Document_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TaxDocumentDate = (*xsdDate)(layout.T.TaxDocumentDate)
	return e.EncodeElement(layout, start)
}
func (t *AdHocPayeeDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AdHocPayeeDataType
	var overlay struct {
		*T
		TaxDocumentDate *xsdDate `xml:"urn:com.workday/bsvc Tax_Document_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TaxDocumentDate = (*xsdDate)(overlay.T.TaxDocumentDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type AdHocPayeeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AdHocPayeeObjectType struct {
	ID         []AdHocPayeeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element used to Request certain Ad Hoc Payees
type AdHocPayeeRequestCriteriaType struct {
	AdHocPayeeName string `xml:"urn:com.workday/bsvc Ad_Hoc_Payee_Name,omitempty"`
}

// Root element for the Response on the Get operation. Contains the instances returned by the Get operation and their accompanying data.
type AdHocPayeeRequestReferencesType struct {
	AdHocPayeeReference []AdHocPayeeObjectType `xml:"urn:com.workday/bsvc Ad_Hoc_Payee_Reference"`
}

// Customer wrapper element that contains the Ad Hoc Payee data elements that are included in the response
type AdHocPayeeResponseDataType struct {
	AdHocPayee []AdHocPayeeType `xml:"urn:com.workday/bsvc Ad_Hoc_Payee,omitempty"`
}

// Wrapper element around a list of elements representing the amount of data that should be included in the Ad Hoc Payee response.
type AdHocPayeeResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing Ad hoc Payee reference and data
type AdHocPayeeType struct {
	AdHocPayeeReference *AdHocPayeeObjectType `xml:"urn:com.workday/bsvc Ad_Hoc_Payee_Reference,omitempty"`
	AdHocPayeeData      []AdHocPayeeDataType  `xml:"urn:com.workday/bsvc Ad_Hoc_Payee_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AdHocPaymentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AdHocPaymentObjectType struct {
	ID         []AdHocPaymentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The data for the Ad Hoc Payment Template
type AdHocPaymentTemplateDataType struct {
	AdHocPaymentTemplateID             string                                `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_Template_ID,omitempty"`
	Name                               string                                `xml:"urn:com.workday/bsvc Name"`
	Inactive                           *bool                                 `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	CompanyReference                   *CompanyObjectType                    `xml:"urn:com.workday/bsvc Company_Reference"`
	BankAccountReference               *FinancialAccountObjectType           `xml:"urn:com.workday/bsvc Bank_Account_Reference"`
	PayeeReference                     *PayeeObjectType                      `xml:"urn:com.workday/bsvc Payee_Reference"`
	CreatenewAdhocPayeeName            string                                `xml:"urn:com.workday/bsvc Create_new_Ad_hoc_Payee_Name"`
	CurrencyReference                  *CurrencyObjectType                   `xml:"urn:com.workday/bsvc Currency_Reference"`
	EliminateForeignExchangeGainorLoss *bool                                 `xml:"urn:com.workday/bsvc Eliminate_Foreign_Exchange_Gain_or_Loss,omitempty"`
	PaymentTypeReference               *PaymentTypeObjectType                `xml:"urn:com.workday/bsvc Payment_Type_Reference"`
	Memo                               string                                `xml:"urn:com.workday/bsvc Memo,omitempty"`
	Addenda                            string                                `xml:"urn:com.workday/bsvc Addenda,omitempty"`
	TaxAuthorityFormTypeReference      *TaxAuthorityFormTypeObjectType       `xml:"urn:com.workday/bsvc Tax_Authority_Form_Type_Reference,omitempty"`
	TINTypeReference                   *TaxpayerIDNumberTypeObjectType       `xml:"urn:com.workday/bsvc TIN_Type_Reference,omitempty"`
	TaxID                              string                                `xml:"urn:com.workday/bsvc Tax_ID,omitempty"`
	TaxPayment                         *bool                                 `xml:"urn:com.workday/bsvc Tax_Payment,omitempty"`
	AdHocPaymentLineDefaults           *AdHocPaymentTemplateLineDataType     `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_Line_Defaults"`
	AddressData                        []AddressInformationDataType          `xml:"urn:com.workday/bsvc Address_Data,omitempty"`
	BankData                           []SettlementAccountWWSDataType        `xml:"urn:com.workday/bsvc Bank_Data,omitempty"`
	PayeeAlternateData                 []BusinessEntityAlternateNameDataType `xml:"urn:com.workday/bsvc Payee_Alternate_Data,omitempty"`
	AttachmentData                     []FinancialsAttachmentDataType        `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
}

// Data for Ad Hoc Payment Template Line Defaults
type AdHocPaymentTemplateLineDataType struct {
	LineCompanyReference      *CompanyObjectType            `xml:"urn:com.workday/bsvc Line_Company_Reference"`
	ItemReference             *ProcurementItemObjectType    `xml:"urn:com.workday/bsvc Item_Reference,omitempty"`
	ItemDescription           string                        `xml:"urn:com.workday/bsvc Item_Description,omitempty"`
	SpendCategoryReference    *SpendCategoryObjectType      `xml:"urn:com.workday/bsvc Spend_Category_Reference,omitempty"`
	TaxApplicabilityReference *TaxApplicabilityObjectType   `xml:"urn:com.workday/bsvc Tax_Applicability_Reference,omitempty"`
	WorktagsReference         []AccountingWorktagObjectType `xml:"urn:com.workday/bsvc Worktags_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AdHocPaymentTemplateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AdHocPaymentTemplateObjectType struct {
	ID         []AdHocPaymentTemplateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Request Criteria for Ad Hoc Payment Template
type AdHocPaymentTemplateRequestCriteriaType struct {
	CompanyReference     []CompanyObjectType          `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	BankAccountReference []FinancialAccountObjectType `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	PayeeReference       []PayeeObjectType            `xml:"urn:com.workday/bsvc Payee_Reference,omitempty"`
	CurrencyReference    []CurrencyObjectType         `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
}

// The reference instance for an Ad Hoc Payment Template
type AdHocPaymentTemplateRequestReferencesType struct {
	AdHocPaymentTemplateReference []AdHocPaymentTemplateObjectType `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_Template_Reference"`
}

// The response for an Ad Hoc Payment Template
type AdHocPaymentTemplateResponseDataType struct {
	AdHocPaymentTemplate []AdHocPaymentTemplateType `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_Template,omitempty"`
}

// The grouping options for the Ad hoc Payment Template response
type AdHocPaymentTemplateResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// The Ad Hoc Payment Template reference and its data
type AdHocPaymentTemplateType struct {
	AdHocPaymentTemplateReference *AdHocPaymentTemplateObjectType `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_Template_Reference,omitempty"`
	AdHocPaymentTemplateData      []AdHocPaymentTemplateDataType  `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_Template_Data,omitempty"`
}

// The data used to create an Ad Hoc Payment from Template.
type AdHocPaymentfromTemplateDataType struct {
	AdHocPaymentTemplateReference *AdHocPaymentTemplateObjectType `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_Template_Reference"`
	Submit                        *bool                           `xml:"urn:com.workday/bsvc Submit,omitempty"`
	AdHocPaymentID                string                          `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_ID,omitempty"`
	PaymentDate                   time.Time                       `xml:"urn:com.workday/bsvc Payment_Date"`
	Amount                        float64                         `xml:"urn:com.workday/bsvc Amount"`
	Memo                          string                          `xml:"urn:com.workday/bsvc Memo,omitempty"`
	ExternalReference             string                          `xml:"urn:com.workday/bsvc External_Reference,omitempty"`
	Addenda                       string                          `xml:"urn:com.workday/bsvc Addenda,omitempty"`
	FreightAmount                 float64                         `xml:"urn:com.workday/bsvc Freight_Amount,omitempty"`
	OtherCharges                  float64                         `xml:"urn:com.workday/bsvc Other_Charges,omitempty"`
	AttachmentData                []FinancialsAttachmentDataType  `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
}

func (t *AdHocPaymentfromTemplateDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AdHocPaymentfromTemplateDataType
	var layout struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date"`
	}
	layout.T = (*T)(t)
	layout.PaymentDate = (*xsdDate)(&layout.T.PaymentDate)
	return e.EncodeElement(layout, start)
}
func (t *AdHocPaymentfromTemplateDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AdHocPaymentfromTemplateDataType
	var overlay struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDate = (*xsdDate)(&overlay.T.PaymentDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type AdditionalReferenceTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AdditionalReferenceTypeObjectType struct {
	ID         []AdditionalReferenceTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Address information
type AddressInformationDataType struct {
	CountryReference             *CountryObjectType                            `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	LastModified                 *time.Time                                    `xml:"urn:com.workday/bsvc Last_Modified,omitempty"`
	AddressLineData              []AddressLineInformationDataType              `xml:"urn:com.workday/bsvc Address_Line_Data,omitempty"`
	Municipality                 string                                        `xml:"urn:com.workday/bsvc Municipality,omitempty"`
	CountryCityReference         *CountryCityObjectType                        `xml:"urn:com.workday/bsvc Country_City_Reference,omitempty"`
	SubmunicipalityData          []SubmunicipalityInformationDataType          `xml:"urn:com.workday/bsvc Submunicipality_Data,omitempty"`
	CountryRegionReference       *CountryRegionObjectType                      `xml:"urn:com.workday/bsvc Country_Region_Reference,omitempty"`
	CountryRegionDescriptor      string                                        `xml:"urn:com.workday/bsvc Country_Region_Descriptor,omitempty"`
	SubregionData                []SubregionInformationDataType                `xml:"urn:com.workday/bsvc Subregion_Data,omitempty"`
	PostalCode                   string                                        `xml:"urn:com.workday/bsvc Postal_Code,omitempty"`
	UsageData                    []CommunicationMethodUsageInformationDataType `xml:"urn:com.workday/bsvc Usage_Data,omitempty"`
	NumberofDays                 float64                                       `xml:"urn:com.workday/bsvc Number_of_Days,omitempty"`
	MunicipalityLocal            string                                        `xml:"urn:com.workday/bsvc Municipality_Local,omitempty"`
	AddressReference             *AddressReferenceObjectType                   `xml:"urn:com.workday/bsvc Address_Reference,omitempty"`
	AddressID                    string                                        `xml:"urn:com.workday/bsvc Address_ID,omitempty"`
	FormattedAddress             string                                        `xml:"urn:com.workday/bsvc Formatted_Address,attr,omitempty"`
	AddressFormatType            string                                        `xml:"urn:com.workday/bsvc Address_Format_Type,attr,omitempty"`
	DefaultedBusinessSiteAddress bool                                          `xml:"urn:com.workday/bsvc Defaulted_Business_Site_Address,attr,omitempty"`
	Delete                       bool                                          `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	DoNotReplaceAll              bool                                          `xml:"urn:com.workday/bsvc Do_Not_Replace_All,attr,omitempty"`
	EffectiveDate                time.Time                                     `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
}

func (t *AddressInformationDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AddressInformationDataType
	var layout struct {
		*T
		LastModified  *xsdDateTime `xml:"urn:com.workday/bsvc Last_Modified,omitempty"`
		EffectiveDate *xsdDate     `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LastModified = (*xsdDateTime)(layout.T.LastModified)
	layout.EffectiveDate = (*xsdDate)(&layout.T.EffectiveDate)
	return e.EncodeElement(layout, start)
}
func (t *AddressInformationDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AddressInformationDataType
	var overlay struct {
		*T
		LastModified  *xsdDateTime `xml:"urn:com.workday/bsvc Last_Modified,omitempty"`
		EffectiveDate *xsdDate     `xml:"urn:com.workday/bsvc Effective_Date,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LastModified = (*xsdDateTime)(overlay.T.LastModified)
	overlay.EffectiveDate = (*xsdDate)(&overlay.T.EffectiveDate)
	return d.DecodeElement(&overlay, &start)
}

// The address line for the address. This typically contains Street name, street number, apartment, suite number.
type AddressLineInformationDataType struct {
	Value      string `xml:",chardata"`
	Descriptor string `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
	Type       string `xml:"urn:com.workday/bsvc Type,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AddressReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AddressReferenceObjectType struct {
	ID         []AddressReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing all Ad hoc Bank Transaction data
type AdhocBankTransactionDataType struct {
	AdhocBankTransactionID                     string                                          `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_ID,omitempty"`
	AdHocBankTransactionNumber                 string                                          `xml:"urn:com.workday/bsvc Ad_Hoc_Bank_Transaction_Number,omitempty"`
	Submit                                     *bool                                           `xml:"urn:com.workday/bsvc Submit,omitempty"`
	LockedinWorkday                            *bool                                           `xml:"urn:com.workday/bsvc Locked_in_Workday,omitempty"`
	TransactionDate                            time.Time                                       `xml:"urn:com.workday/bsvc Transaction_Date"`
	TransactionMemo                            string                                          `xml:"urn:com.workday/bsvc Transaction_Memo"`
	CompanyReference                           *CompanyObjectType                              `xml:"urn:com.workday/bsvc Company_Reference"`
	CurrencyReference                          *CurrencyObjectType                             `xml:"urn:com.workday/bsvc Currency_Reference"`
	BankAccountReference                       *FinancialAccountObjectType                     `xml:"urn:com.workday/bsvc Bank_Account_Reference"`
	CurrencyRateTypeReference                  *CurrencyRateTypeObjectType                     `xml:"urn:com.workday/bsvc Currency_Rate_Type_Reference,omitempty"`
	CurrencyRateOverride                       float64                                         `xml:"urn:com.workday/bsvc Currency_Rate_Override,omitempty"`
	TransactionAmount                          float64                                         `xml:"urn:com.workday/bsvc Transaction_Amount"`
	Deposit                                    bool                                            `xml:"urn:com.workday/bsvc Deposit"`
	Withdrawal                                 bool                                            `xml:"urn:com.workday/bsvc Withdrawal"`
	AdhocBankTransactionPurposeReference       *AdHocBankTransactionPurposeObjectType          `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_Purpose_Reference,omitempty"`
	TransactionID                              string                                          `xml:"urn:com.workday/bsvc Transaction_ID,omitempty"`
	IncludeinIRS1099                           *bool                                           `xml:"urn:com.workday/bsvc Include_in_IRS_1099,omitempty"`
	JournalSourceReference                     *JournalSourceObjectType                        `xml:"urn:com.workday/bsvc Journal_Source_Reference,omitempty"`
	RemoveBankAccountWorktagonOffset           string                                          `xml:"urn:com.workday/bsvc Remove_Bank_Account_Worktag_on_Offset,omitempty"`
	EliminateFXGainLoss                        string                                          `xml:"urn:com.workday/bsvc Eliminate_FX_Gain_Loss,omitempty"`
	TransactionLineReplacementData             []AdhocBankTransactionLineDataType              `xml:"urn:com.workday/bsvc Transaction_Line_Replacement_Data,omitempty"`
	AttachmentData                             []FinancialsAttachmentDataType                  `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
	AdhocBankTransactionIntercompanySubProcess *AdhocBankTransactionIntercompanySubProcessType `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_Intercompany_Sub_Process,omitempty"`
}

func (t *AdhocBankTransactionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AdhocBankTransactionDataType
	var layout struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date"`
	}
	layout.T = (*T)(t)
	layout.TransactionDate = (*xsdDate)(&layout.T.TransactionDate)
	return e.EncodeElement(layout, start)
}
func (t *AdhocBankTransactionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AdhocBankTransactionDataType
	var overlay struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionDate = (*xsdDate)(&overlay.T.TransactionDate)
	return d.DecodeElement(&overlay, &start)
}

// Element containing all Ad hoc Bank Transaction data
type AdhocBankTransactionHVDataType struct {
	AdhocBankTransactionID                     string                                          `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_ID,omitempty"`
	Submit                                     *bool                                           `xml:"urn:com.workday/bsvc Submit,omitempty"`
	LockedinWorkday                            *bool                                           `xml:"urn:com.workday/bsvc Locked_in_Workday,omitempty"`
	TransactionDate                            time.Time                                       `xml:"urn:com.workday/bsvc Transaction_Date"`
	TransactionMemo                            string                                          `xml:"urn:com.workday/bsvc Transaction_Memo"`
	CompanyReference                           *CompanyObjectType                              `xml:"urn:com.workday/bsvc Company_Reference"`
	CurrencyReference                          *CurrencyObjectType                             `xml:"urn:com.workday/bsvc Currency_Reference"`
	BankAccountReference                       *FinancialAccountObjectType                     `xml:"urn:com.workday/bsvc Bank_Account_Reference"`
	CurrencyRateTypeReference                  *CurrencyRateTypeObjectType                     `xml:"urn:com.workday/bsvc Currency_Rate_Type_Reference,omitempty"`
	CurrencyRateOverride                       float64                                         `xml:"urn:com.workday/bsvc Currency_Rate_Override,omitempty"`
	TransactionAmount                          float64                                         `xml:"urn:com.workday/bsvc Transaction_Amount"`
	Deposit                                    bool                                            `xml:"urn:com.workday/bsvc Deposit"`
	Withdrawal                                 bool                                            `xml:"urn:com.workday/bsvc Withdrawal"`
	AdhocBankTransactionPurposeReference       *AdHocBankTransactionPurposeObjectType          `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_Purpose_Reference,omitempty"`
	TransactionID                              string                                          `xml:"urn:com.workday/bsvc Transaction_ID,omitempty"`
	IncludeinIRS1099                           *bool                                           `xml:"urn:com.workday/bsvc Include_in_IRS_1099,omitempty"`
	JournalSourceReference                     *JournalSourceObjectType                        `xml:"urn:com.workday/bsvc Journal_Source_Reference,omitempty"`
	RemoveBankAccountWorktagonOffset           string                                          `xml:"urn:com.workday/bsvc Remove_Bank_Account_Worktag_on_Offset,omitempty"`
	EliminateFXGainLoss                        string                                          `xml:"urn:com.workday/bsvc Eliminate_FX_Gain_Loss,omitempty"`
	TransactionLineReplacementData             []AdhocBankTransactionLineHVDataType            `xml:"urn:com.workday/bsvc Transaction_Line_Replacement_Data,omitempty"`
	AttachmentData                             []FinancialsAttachmentDataType                  `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
	AdhocBankTransactionIntercompanySubProcess *AdhocBankTransactionIntercompanySubProcessType `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_Intercompany_Sub_Process,omitempty"`
}

func (t *AdhocBankTransactionHVDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AdhocBankTransactionHVDataType
	var layout struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date"`
	}
	layout.T = (*T)(t)
	layout.TransactionDate = (*xsdDate)(&layout.T.TransactionDate)
	return e.EncodeElement(layout, start)
}
func (t *AdhocBankTransactionHVDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AdhocBankTransactionHVDataType
	var overlay struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionDate = (*xsdDate)(&overlay.T.TransactionDate)
	return d.DecodeElement(&overlay, &start)
}

// Wrapper for the Ad Hoc Bank Transaction Intercompany Sub Process.  It allows you to specify parameters for the sub process.
type AdhocBankTransactionIntercompanySubProcessType struct {
	BusinessSubProcessParameters *BusinessSubProcessParametersType `xml:"urn:com.workday/bsvc Business_Sub_Process_Parameters,omitempty"`
}

// Contains data for an Ad hoc Bank Transaction Line
type AdhocBankTransactionLineDataType struct {
	LineOrder                          string                        `xml:"urn:com.workday/bsvc Line_Order,omitempty"`
	IntercompanyAffiliateReference     *CompanyObjectType            `xml:"urn:com.workday/bsvc Intercompany_Affiliate_Reference,omitempty"`
	ResourceCategoryReference          *SpendCategoryObjectType      `xml:"urn:com.workday/bsvc Resource_Category_Reference,omitempty"`
	RevenueCategoryReference           *RevenueCategoryObjectType    `xml:"urn:com.workday/bsvc Revenue_Category_Reference,omitempty"`
	LedgerAccountReference             *LedgerAccountObjectType      `xml:"urn:com.workday/bsvc Ledger_Account_Reference,omitempty"`
	AlternateLedgerAccountReference    *LedgerAccountObjectType      `xml:"urn:com.workday/bsvc Alternate_Ledger_Account_Reference,omitempty"`
	LineAmount                         float64                       `xml:"urn:com.workday/bsvc Line_Amount"`
	LineMemo                           string                        `xml:"urn:com.workday/bsvc Line_Memo,omitempty"`
	WorktagsReference                  []AccountingWorktagObjectType `xml:"urn:com.workday/bsvc Worktags_Reference,omitempty"`
	BalancingWorktagAffiliateReference *BalancingWorktagObjectType   `xml:"urn:com.workday/bsvc Balancing_Worktag_Affiliate_Reference,omitempty"`
}

// Contains data for an Ad hoc Bank Transaction Line
type AdhocBankTransactionLineHVDataType struct {
	LineOrder                          string                        `xml:"urn:com.workday/bsvc Line_Order,omitempty"`
	IntercompanyAffiliateReference     *CompanyObjectType            `xml:"urn:com.workday/bsvc Intercompany_Affiliate_Reference,omitempty"`
	ResourceCategoryReference          *SpendCategoryObjectType      `xml:"urn:com.workday/bsvc Resource_Category_Reference,omitempty"`
	RevenueCategoryReference           *RevenueCategoryObjectType    `xml:"urn:com.workday/bsvc Revenue_Category_Reference,omitempty"`
	LedgerAccountReference             *LedgerAccountObjectType      `xml:"urn:com.workday/bsvc Ledger_Account_Reference,omitempty"`
	AlternateLedgerAccountReference    *LedgerAccountObjectType      `xml:"urn:com.workday/bsvc Alternate_Ledger_Account_Reference,omitempty"`
	LineAmount                         float64                       `xml:"urn:com.workday/bsvc Line_Amount"`
	LineMemo                           string                        `xml:"urn:com.workday/bsvc Line_Memo,omitempty"`
	WorktagsReference                  []AccountingWorktagObjectType `xml:"urn:com.workday/bsvc Worktags_Reference,omitempty"`
	BalancingWorktagAffiliateReference *BalancingWorktagObjectType   `xml:"urn:com.workday/bsvc Balancing_Worktag_Affiliate_Reference,omitempty"`
}

// Element containing ad hoc bank transaction request criteria
type AdhocBankTransactionRequestCriteriaType struct {
	CompanyReference           []OrganizationObjectType     `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	BankAccountReference       []FinancialAccountObjectType `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	TransactionOnorAfter       *time.Time                   `xml:"urn:com.workday/bsvc Transaction_On_or_After,omitempty"`
	TransactionOnorBefore      *time.Time                   `xml:"urn:com.workday/bsvc Transaction_On_or_Before,omitempty"`
	AdHocBankTransactionNumber string                       `xml:"urn:com.workday/bsvc Ad_Hoc_Bank_Transaction_Number,omitempty"`
}

func (t *AdhocBankTransactionRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AdhocBankTransactionRequestCriteriaType
	var layout struct {
		*T
		TransactionOnorAfter  *xsdDate `xml:"urn:com.workday/bsvc Transaction_On_or_After,omitempty"`
		TransactionOnorBefore *xsdDate `xml:"urn:com.workday/bsvc Transaction_On_or_Before,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TransactionOnorAfter = (*xsdDate)(layout.T.TransactionOnorAfter)
	layout.TransactionOnorBefore = (*xsdDate)(layout.T.TransactionOnorBefore)
	return e.EncodeElement(layout, start)
}
func (t *AdhocBankTransactionRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AdhocBankTransactionRequestCriteriaType
	var overlay struct {
		*T
		TransactionOnorAfter  *xsdDate `xml:"urn:com.workday/bsvc Transaction_On_or_After,omitempty"`
		TransactionOnorBefore *xsdDate `xml:"urn:com.workday/bsvc Transaction_On_or_Before,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionOnorAfter = (*xsdDate)(overlay.T.TransactionOnorAfter)
	overlay.TransactionOnorBefore = (*xsdDate)(overlay.T.TransactionOnorBefore)
	return d.DecodeElement(&overlay, &start)
}

// Element containing reference instances for an ad hoc bank transaction
type AdhocBankTransactionRequestReferencesType struct {
	AdhocBankTransactionReference []AdHocBankTransactionObjectType `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_Reference"`
}

// Wrapper Element that includes Ad hoc Bank Transaction Reference Instance and Bank Account Data
type AdhocBankTransactionResponseDataType struct {
	AdhocBankTransaction []AdhocBankTransactionType `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction,omitempty"`
}

// Element containing bank account response grouping options
type AdhocBankTransactionResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing Ad hoc Bank Transaction reference and data
type AdhocBankTransactionType struct {
	AdhocBankTransactionReference *AdHocBankTransactionObjectType `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_Reference,omitempty"`
	AdhocBankTransactionData      []AdhocBankTransactionDataType  `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_Data,omitempty"`
}

// Element containing all Ad hoc Payment data
type AdhocPaymentDataType struct {
	AdhocPaymentID                     string                                   `xml:"urn:com.workday/bsvc Ad_hoc_Payment_ID,omitempty"`
	Submit                             *bool                                    `xml:"urn:com.workday/bsvc Submit,omitempty"`
	CompanyReference                   *CompanyObjectType                       `xml:"urn:com.workday/bsvc Company_Reference"`
	BankAccountReference               *FinancialAccountObjectType              `xml:"urn:com.workday/bsvc Bank_Account_Reference"`
	PayeeReference                     *PayeeObjectType                         `xml:"urn:com.workday/bsvc Payee_Reference,omitempty"`
	CreatenewAdhocPayeeName            string                                   `xml:"urn:com.workday/bsvc Create_new_Ad_hoc_Payee_Name,omitempty"`
	SingleUse                          *bool                                    `xml:"urn:com.workday/bsvc Single_Use,omitempty"`
	TaxAuthorityFormTypeReference      *TaxAuthorityFormTypeObjectType          `xml:"urn:com.workday/bsvc Tax_Authority_Form_Type_Reference,omitempty"`
	IRS1099MISCPayee                   *bool                                    `xml:"urn:com.workday/bsvc IRS_1099_MISC_Payee,omitempty"`
	TINTypeReference                   *TaxpayerIDNumberTypeObjectType          `xml:"urn:com.workday/bsvc TIN_Type_Reference,omitempty"`
	TaxID                              string                                   `xml:"urn:com.workday/bsvc Tax_ID,omitempty"`
	CurrencyReference                  *CurrencyObjectType                      `xml:"urn:com.workday/bsvc Currency_Reference"`
	CurrencyRateTypeReference          *CurrencyRateTypeObjectType              `xml:"urn:com.workday/bsvc Currency_Rate_Type_Reference,omitempty"`
	CurrencyConversionRate             float64                                  `xml:"urn:com.workday/bsvc Currency_Conversion_Rate,omitempty"`
	EliminateForeignExchangeGainorLoss *bool                                    `xml:"urn:com.workday/bsvc Eliminate_Foreign_Exchange_Gain_or_Loss,omitempty"`
	TaxOptionReference                 *TaxOptionObjectType                     `xml:"urn:com.workday/bsvc Tax_Option_Reference,omitempty"`
	ShipToAddressReference             *UniqueIdentifierObjectType              `xml:"urn:com.workday/bsvc Ship-To_Address_Reference,omitempty"`
	TaxCodeReference                   *TaxCodeObjectType                       `xml:"urn:com.workday/bsvc Tax_Code_Reference,omitempty"`
	PaymentDate                        time.Time                                `xml:"urn:com.workday/bsvc Payment_Date"`
	PaymentTypeReference               *PaymentTypeObjectType                   `xml:"urn:com.workday/bsvc Payment_Type_Reference,omitempty"`
	Memo                               string                                   `xml:"urn:com.workday/bsvc Memo,omitempty"`
	ReferenceNumber                    string                                   `xml:"urn:com.workday/bsvc Reference_Number,omitempty"`
	ControlTotalAmount                 float64                                  `xml:"urn:com.workday/bsvc Control_Total_Amount,omitempty"`
	TaxAmount                          float64                                  `xml:"urn:com.workday/bsvc Tax_Amount,omitempty"`
	FreightAmount                      float64                                  `xml:"urn:com.workday/bsvc Freight_Amount,omitempty"`
	OtherCharges                       float64                                  `xml:"urn:com.workday/bsvc Other_Charges,omitempty"`
	HandlingCodeReference              *PaymentHandlingInstructionObjectType    `xml:"urn:com.workday/bsvc Handling_Code_Reference,omitempty"`
	AddendaLinesasText                 string                                   `xml:"urn:com.workday/bsvc Addenda_Lines_as_Text,omitempty"`
	ExternalReferenceasText            string                                   `xml:"urn:com.workday/bsvc External_Reference_as_Text,omitempty"`
	InvoiceLineReplacementData         []SupplierInvoiceLineReplacementDataType `xml:"urn:com.workday/bsvc Invoice_Line_Replacement_Data,omitempty"`
	TaxCodeData                        []TaxableCodeApplicationDataType         `xml:"urn:com.workday/bsvc Tax_Code_Data,omitempty"`
	AddressData                        []AddressInformationDataType             `xml:"urn:com.workday/bsvc Address_Data,omitempty"`
	BankData                           []SettlementAccountWWSDataType           `xml:"urn:com.workday/bsvc Bank_Data,omitempty"`
	PayeeAlternateNameData             []BusinessEntityAlternateNameDataType    `xml:"urn:com.workday/bsvc Payee_Alternate_Name_Data,omitempty"`
	AttachmentData                     []FinancialsAttachmentDataType           `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
}

func (t *AdhocPaymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AdhocPaymentDataType
	var layout struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date"`
	}
	layout.T = (*T)(t)
	layout.PaymentDate = (*xsdDate)(&layout.T.PaymentDate)
	return e.EncodeElement(layout, start)
}
func (t *AdhocPaymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AdhocPaymentDataType
	var overlay struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDate = (*xsdDate)(&overlay.T.PaymentDate)
	return d.DecodeElement(&overlay, &start)
}

// Element containing ad hoc payment request criteria
type AdhocPaymentRequestCriteriaType struct {
	CompanyReference          *OrganizationObjectType      `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	BankAccountReference      []FinancialAccountObjectType `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	PayeeReference            []PayeeObjectType            `xml:"urn:com.workday/bsvc Payee_Reference,omitempty"`
	PayeeHierarchyReference   []HierarchyObjectType        `xml:"urn:com.workday/bsvc Payee_Hierarchy_Reference,omitempty"`
	PaymentDateonDateorAfter  *time.Time                   `xml:"urn:com.workday/bsvc Payment_Date_on_Date_or_After,omitempty"`
	PaymentDateonDateorBefore *time.Time                   `xml:"urn:com.workday/bsvc Payment_Date_on_Date_or_Before,omitempty"`
	CurrencyReference         []CurrencyObjectType         `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
}

func (t *AdhocPaymentRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AdhocPaymentRequestCriteriaType
	var layout struct {
		*T
		PaymentDateonDateorAfter  *xsdDate `xml:"urn:com.workday/bsvc Payment_Date_on_Date_or_After,omitempty"`
		PaymentDateonDateorBefore *xsdDate `xml:"urn:com.workday/bsvc Payment_Date_on_Date_or_Before,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PaymentDateonDateorAfter = (*xsdDate)(layout.T.PaymentDateonDateorAfter)
	layout.PaymentDateonDateorBefore = (*xsdDate)(layout.T.PaymentDateonDateorBefore)
	return e.EncodeElement(layout, start)
}
func (t *AdhocPaymentRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AdhocPaymentRequestCriteriaType
	var overlay struct {
		*T
		PaymentDateonDateorAfter  *xsdDate `xml:"urn:com.workday/bsvc Payment_Date_on_Date_or_After,omitempty"`
		PaymentDateonDateorBefore *xsdDate `xml:"urn:com.workday/bsvc Payment_Date_on_Date_or_Before,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDateonDateorAfter = (*xsdDate)(overlay.T.PaymentDateonDateorAfter)
	overlay.PaymentDateonDateorBefore = (*xsdDate)(overlay.T.PaymentDateonDateorBefore)
	return d.DecodeElement(&overlay, &start)
}

// Element containing reference instances for an ad hoc payment
type AdhocPaymentRequestReferencesType struct {
	AdhocPaymentReference []AdHocPaymentObjectType `xml:"urn:com.workday/bsvc Ad_hoc_Payment_Reference"`
}

// Wrapper Element that includes Ad hoc Payment Reference Instance and Ad hoc Payment Data
type AdhocPaymentResponseDataType struct {
	AdhocPayment []AdhocPaymentType `xml:"urn:com.workday/bsvc Ad_hoc_Payment,omitempty"`
}

// Element containing Ad hoc Payment response grouping options
type AdhocPaymentResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing Ad hoc Payment reference and data
type AdhocPaymentType struct {
	AdhocPaymentReference *AdHocPaymentObjectType `xml:"urn:com.workday/bsvc Ad_hoc_Payment_Reference,omitempty"`
	AdhocPaymentData      []AdhocPaymentDataType  `xml:"urn:com.workday/bsvc Ad_hoc_Payment_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AlternateNameUsageObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AlternateNameUsageObjectType struct {
	ID         []AlternateNameUsageObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing application related exceptions data
type ApplicationInstanceExceptionsDataType struct {
	ExceptionData []ExceptionDataType `xml:"urn:com.workday/bsvc Exception_Data,omitempty"`
}

// Element containing Exceptions Data
type ApplicationInstanceRelatedExceptionsDataType struct {
	ExceptionsData []ApplicationInstanceExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Data,omitempty"`
}

// The details of a resume (resume, name, comments).
type AttachmentDataWWSType struct {
	FileID   string  `xml:"urn:com.workday/bsvc File_ID,omitempty"`
	File     *[]byte `xml:"urn:com.workday/bsvc File,omitempty"`
	FileName string  `xml:"urn:com.workday/bsvc FileName,omitempty"`
	Comments string  `xml:"urn:com.workday/bsvc Comments,omitempty"`
}

func (t *AttachmentDataWWSType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AttachmentDataWWSType
	var layout struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	layout.T = (*T)(t)
	layout.File = (*xsdBase64Binary)(layout.T.File)
	return e.EncodeElement(layout, start)
}
func (t *AttachmentDataWWSType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AttachmentDataWWSType
	var overlay struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.File = (*xsdBase64Binary)(overlay.T.File)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type AuthorityDesignationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AuthorityDesignationObjectType struct {
	ID         []AuthorityDesignationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Authority Type data.
type AuthorityTypeDataType struct {
	AuthorityTypeID               string                          `xml:"urn:com.workday/bsvc Authority_Type_ID,omitempty"`
	AuthorityType                 string                          `xml:"urn:com.workday/bsvc Authority_Type"`
	AuthorityDesignationReference *AuthorityDesignationObjectType `xml:"urn:com.workday/bsvc Authority_Designation_Reference"`
	Description                   string                          `xml:"urn:com.workday/bsvc Description,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AuthorityTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AuthorityTypeObjectType struct {
	ID         []AuthorityTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The search criteria for the request.
type AuthorityTypeRequestCriteriaType struct {
}

// Contains the reference to the Authority Type.
type AuthorityTypeRequestReferencesType struct {
	AuthorityTypeReference []AuthorityTypeObjectType `xml:"urn:com.workday/bsvc Authority_Type_Reference"`
}

// Wrapper containing the response data.
type AuthorityTypeResponseDataType struct {
	AuthorityType []AuthorityTypeType `xml:"urn:com.workday/bsvc Authority_Type,omitempty"`
}

// Contains a flag that determines whether the reference is included in the response.
type AuthorityTypeResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Response data containing the reference and data items.
type AuthorityTypeType struct {
	AuthorityTypeReference *AuthorityTypeObjectType `xml:"urn:com.workday/bsvc Authority_Type_Reference,omitempty"`
	AuthorityTypeData      []AuthorityTypeDataType  `xml:"urn:com.workday/bsvc Authority_Type_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type AutoReconInitiationTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type AutoReconInitiationTypeObjectType struct {
	ID         []AutoReconInitiationTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BalancingWorktagObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BalancingWorktagObjectType struct {
	ID         []BalancingWorktagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing all bank account data
type BankAccountDataType struct {
	BankAccountID                                      string                                          `xml:"urn:com.workday/bsvc Bank_Account_ID,omitempty"`
	AccountName                                        string                                          `xml:"urn:com.workday/bsvc Account_Name"`
	FinancialInstitutionReference                      *FinancialInstitutionObjectType                 `xml:"urn:com.workday/bsvc Financial_Institution_Reference,omitempty"`
	FinancialPartyReference                            *FinancialPartyObjectType                       `xml:"urn:com.workday/bsvc Financial_Party_Reference"`
	DefaultCurrencyReference                           *CurrencyObjectType                             `xml:"urn:com.workday/bsvc Default_Currency_Reference,omitempty"`
	AcceptedCurrenciesReference                        []CurrencyObjectType                            `xml:"urn:com.workday/bsvc Accepted_Currencies_Reference,omitempty"`
	BankBranchReference                                *BankBranchObjectType                           `xml:"urn:com.workday/bsvc Bank_Branch_Reference,omitempty"`
	BankAccountOpenDate                                *time.Time                                      `xml:"urn:com.workday/bsvc Bank_Account_Open_Date,omitempty"`
	AccountClosed                                      *bool                                           `xml:"urn:com.workday/bsvc Account_Closed,omitempty"`
	BankAccountCloseDate                               *time.Time                                      `xml:"urn:com.workday/bsvc Bank_Account_Close_Date,omitempty"`
	CountryReference                                   *CountryObjectType                              `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	RoutingTransitorInstitutionNumber                  string                                          `xml:"urn:com.workday/bsvc Routing_Transit_or_Institution_Number,omitempty"`
	FinancialAccountNumber                             string                                          `xml:"urn:com.workday/bsvc Financial_Account_Number,omitempty"`
	BankIdentifierCode                                 string                                          `xml:"urn:com.workday/bsvc Bank_Identifier_Code,omitempty"`
	IBAN                                               string                                          `xml:"urn:com.workday/bsvc IBAN,omitempty"`
	BankAccountFormatReference                         *BankStatementFormatObjectType                  `xml:"urn:com.workday/bsvc Bank_Account_Format_Reference,omitempty"`
	CustomCodeSetReference                             *BankStatementCustomCodeSetObjectType           `xml:"urn:com.workday/bsvc Custom_Code_Set_Reference,omitempty"`
	BranchName                                         string                                          `xml:"urn:com.workday/bsvc Branch_Name,omitempty"`
	CheckDigit                                         string                                          `xml:"urn:com.workday/bsvc Check_Digit,omitempty"`
	BankAccountName                                    string                                          `xml:"urn:com.workday/bsvc Bank_Account_Name,omitempty"`
	RollNumber                                         string                                          `xml:"urn:com.workday/bsvc Roll_Number,omitempty"`
	Fraction                                           string                                          `xml:"urn:com.workday/bsvc Fraction,omitempty"`
	FormattedMICR                                      string                                          `xml:"urn:com.workday/bsvc Formatted_MICR,omitempty"`
	TargetBalance                                      float64                                         `xml:"urn:com.workday/bsvc Target_Balance,omitempty"`
	PaymentTypeReference                               []PaymentTypeObjectType                         `xml:"urn:com.workday/bsvc Payment_Type_Reference"`
	DefaultPaymentTypeforAdHocPaymentsReference        *PaymentTypeObjectType                          `xml:"urn:com.workday/bsvc Default_Payment_Type_for_Ad_Hoc_Payments_Reference,omitempty"`
	BankAccountSecuritySegmentReference                []BankAccountSecuritySegmentObjectType          `xml:"urn:com.workday/bsvc Bank_Account_Security_Segment_Reference,omitempty"`
	CheckPrintLayoutReference                          *CheckPrintLayoutAbstractObjectType             `xml:"urn:com.workday/bsvc Check_Print_Layout_Reference,omitempty"`
	UseBranchAddress                                   *bool                                           `xml:"urn:com.workday/bsvc Use_Branch_Address,omitempty"`
	AdvancedMode                                       *bool                                           `xml:"urn:com.workday/bsvc Advanced_Mode,omitempty"`
	ParsingRuleSetReference                            *ParsingRuleSetObjectType                       `xml:"urn:com.workday/bsvc Parsing_Rule_Set_Reference,omitempty"`
	PerformAutomaticReconciliationReference            *AutoReconInitiationTypeObjectType              `xml:"urn:com.workday/bsvc Perform_Automatic_Reconciliation_Reference,omitempty"`
	ReconciliationRuleSetReference                     *ReconciliationRuleSetObjectType                `xml:"urn:com.workday/bsvc Reconciliation_Rule_Set_Reference,omitempty"`
	ReconciliationMatchingRuleSetReference             *BankingReconciliationMatchingRuleSetObjectType `xml:"urn:com.workday/bsvc Reconciliation_Matching_Rule_Set_Reference,omitempty"`
	PerformAutomaticFirstNoticeReconciliationReference *AutoReconInitiationTypeObjectType              `xml:"urn:com.workday/bsvc Perform_Automatic_First_Notice_Reconciliation_Reference,omitempty"`
	FirstNoticeRuleSetReference                        *ReconciliationRuleSetObjectType                `xml:"urn:com.workday/bsvc First_Notice_Rule_Set_Reference,omitempty"`
	BatchElectronicPayments                            *bool                                           `xml:"urn:com.workday/bsvc Batch_Electronic_Payments,omitempty"`
	BatchElectronicCustomerPaymentDeposits             *bool                                           `xml:"urn:com.workday/bsvc Batch_Electronic_Customer_Payment_Deposits,omitempty"`
	SubmitReconciledStatementsAutomatically            *bool                                           `xml:"urn:com.workday/bsvc Submit_Reconciled_Statements_Automatically,omitempty"`
	BankStatementWorktagRuleSetReference               *BankStatementWorktagRuleSetObjectType          `xml:"urn:com.workday/bsvc Bank_Statement_Worktag_Rule_Set_Reference,omitempty"`
	UsedbyCash                                         *bool                                           `xml:"urn:com.workday/bsvc Used_by_Cash,omitempty"`
	UsedbyCustomerPayments                             *bool                                           `xml:"urn:com.workday/bsvc Used_by_Customer_Payments,omitempty"`
	UsedbyCustomerRefunds                              *bool                                           `xml:"urn:com.workday/bsvc Used_by_Customer_Refunds,omitempty"`
	UsedbyExpensePayments                              *bool                                           `xml:"urn:com.workday/bsvc Used_by_Expense_Payments,omitempty"`
	UsedbyPayrollOnCycle                               *bool                                           `xml:"urn:com.workday/bsvc Used_by_Payroll_On_Cycle,omitempty"`
	UsedbySupplierPayments                             *bool                                           `xml:"urn:com.workday/bsvc Used_by_Supplier_Payments,omitempty"`
	UsedbyIntercompanyPayments                         *bool                                           `xml:"urn:com.workday/bsvc Used_by_Intercompany_Payments,omitempty"`
	UsedbyAdhocPayments                                *bool                                           `xml:"urn:com.workday/bsvc Used_by_Ad_hoc_Payments,omitempty"`
	UsedbyBankAccountTransfersforSettlement            *bool                                           `xml:"urn:com.workday/bsvc Used_by_Bank_Account_Transfers_for_Settlement,omitempty"`
	UsedByPayrollOffCycle                              *bool                                           `xml:"urn:com.workday/bsvc Used_By_Payroll_Off_Cycle,omitempty"`
	UsedByPrenotePayments                              *bool                                           `xml:"urn:com.workday/bsvc Used_By_Prenote_Payments,omitempty"`
	UsedByProcurementCardPayments                      *bool                                           `xml:"urn:com.workday/bsvc Used_By_Procurement_Card_Payments,omitempty"`
	UsedByTaxPayments                                  *bool                                           `xml:"urn:com.workday/bsvc Used_By_Tax_Payments,omitempty"`
	UsedbyCashAdvances                                 *bool                                           `xml:"urn:com.workday/bsvc Used_by_Cash_Advances,omitempty"`
	UsedbyExpenseCreditCardPayments                    *bool                                           `xml:"urn:com.workday/bsvc Used_by_Expense_Credit_Card_Payments,omitempty"`
	UsedbyStudentRefund                                *bool                                           `xml:"urn:com.workday/bsvc Used_by_Student_Refund,omitempty"`
	UsedByStudentPayment                               *bool                                           `xml:"urn:com.workday/bsvc Used_By_Student_Payment,omitempty"`
	LastCheckNumberUsed                                float64                                         `xml:"urn:com.workday/bsvc Last_Check_Number_Used,omitempty"`
	EnablePositivePay                                  *bool                                           `xml:"urn:com.workday/bsvc Enable_Positive_Pay,omitempty"`
	Outsourced                                         *bool                                           `xml:"urn:com.workday/bsvc Outsourced,omitempty"`
	AccountTypeReference                               *BankAccountTypeObjectType                      `xml:"urn:com.workday/bsvc Account_Type_Reference,omitempty"`
	AllowAdditionalUsage                               *bool                                           `xml:"urn:com.workday/bsvc Allow_Additional_Usage,omitempty"`
	DefaultBankStatementBeginningBalance               *bool                                           `xml:"urn:com.workday/bsvc Default_Bank_Statement_Beginning_Balance,omitempty"`
	OutsourcedBankProviderReference                    *UniqueIdentifierObjectType                     `xml:"urn:com.workday/bsvc Outsourced_Bank_Provider_Reference,omitempty"`
	PaymentIntegrationsData                            *PaymentIntegrationsDataType                    `xml:"urn:com.workday/bsvc Payment_Integrations_Data,omitempty"`
	LockboxData                                        []LockboxDataType                               `xml:"urn:com.workday/bsvc Lockbox_Data,omitempty"`
	CheckSortingSetupData                              *CheckSortingSetupDataType                      `xml:"urn:com.workday/bsvc Check_Sorting_Setup_Data,omitempty"`
	AttachmentData                                     []FinancialsAttachmentDataType                  `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
}

func (t *BankAccountDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankAccountDataType
	var layout struct {
		*T
		BankAccountOpenDate  *xsdDate `xml:"urn:com.workday/bsvc Bank_Account_Open_Date,omitempty"`
		BankAccountCloseDate *xsdDate `xml:"urn:com.workday/bsvc Bank_Account_Close_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.BankAccountOpenDate = (*xsdDate)(layout.T.BankAccountOpenDate)
	layout.BankAccountCloseDate = (*xsdDate)(layout.T.BankAccountCloseDate)
	return e.EncodeElement(layout, start)
}
func (t *BankAccountDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankAccountDataType
	var overlay struct {
		*T
		BankAccountOpenDate  *xsdDate `xml:"urn:com.workday/bsvc Bank_Account_Open_Date,omitempty"`
		BankAccountCloseDate *xsdDate `xml:"urn:com.workday/bsvc Bank_Account_Close_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.BankAccountOpenDate = (*xsdDate)(overlay.T.BankAccountOpenDate)
	overlay.BankAccountCloseDate = (*xsdDate)(overlay.T.BankAccountCloseDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type BankAccountObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankAccountObjectType struct {
	ID         []BankAccountObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing bank account request criteria
type BankAccountRequestCriteriaType struct {
}

// Element containing reference instances for a bank account
type BankAccountRequestReferencesType struct {
	BankAccountReference []BankAccountObjectType `xml:"urn:com.workday/bsvc Bank_Account_Reference"`
}

// Wrapper Element that includes Bank Account Reference Instance and Bank Account Data
type BankAccountResponseDataType struct {
	BankAccount []BankAccountType `xml:"urn:com.workday/bsvc Bank_Account,omitempty"`
}

// Element containing bank account response grouping options
type BankAccountResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankAccountSecuritySegmentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankAccountSecuritySegmentObjectType struct {
	ID         []BankAccountSecuritySegmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Response data containing the references and data items.
type BankAccountSignatoriesType struct {
	BankAccountReference     *BankAccountObjectType         `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	BankAccountSignatoryData []BankAccountSignatoryDataType `xml:"urn:com.workday/bsvc Bank_Account_Signatory_Data,omitempty"`
}

// Bank Account Signatory data.
type BankAccountSignatoryDataType struct {
	SignatoryReference        *SignatoryDetailsObjectType `xml:"urn:com.workday/bsvc Signatory_Reference,omitempty"`
	SignatoryID               string                      `xml:"urn:com.workday/bsvc Signatory_ID,omitempty"`
	OverrideCurrencyReference *CurrencyObjectType         `xml:"urn:com.workday/bsvc Override_Currency_Reference,omitempty"`
	MinimumCurrency           float64                     `xml:"urn:com.workday/bsvc Minimum_Currency,omitempty"`
	MaximumCurrency           float64                     `xml:"urn:com.workday/bsvc Maximum_Currency"`
	SignerReference           []SignerObjectType          `xml:"urn:com.workday/bsvc Signer_Reference"`
	PaymentTypeReference      []PaymentTypeObjectType     `xml:"urn:com.workday/bsvc Payment_Type_Reference"`
	AuthorityTypeReference    *AuthorityTypeObjectType    `xml:"urn:com.workday/bsvc Authority_Type_Reference"`
	SignatureMethodReference  []SignatureMethodObjectType `xml:"urn:com.workday/bsvc Signature_Method_Reference"`
	StartDate                 time.Time                   `xml:"urn:com.workday/bsvc Start_Date"`
	EndDate                   *time.Time                  `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	Notes                     string                      `xml:"urn:com.workday/bsvc Notes,omitempty"`
}

func (t *BankAccountSignatoryDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankAccountSignatoryDataType
	var layout struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(&layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *BankAccountSignatoryDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankAccountSignatoryDataType
	var overlay struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(&overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains the reference to the Signatory Details instance that is created, updated, or deleted.
type BankAccountSignatoryDetailsType struct {
	SignatoryReference *SignatoryDetailsObjectType `xml:"urn:com.workday/bsvc Signatory_Reference,omitempty"`
}

// Wrapper containing the response data.
type BankAccountSignatoryResponseDataType struct {
	BankAccountSignatory []BankAccountSignatoriesType `xml:"urn:com.workday/bsvc Bank_Account_Signatory,omitempty"`
}

// Bank Account Transfer Data contains essential information about Bank AccountTransfer
type BankAccountTransferDataType struct {
	BankAccountTransferID                         string                         `xml:"urn:com.workday/bsvc Bank_Account_Transfer_ID,omitempty"`
	Submit                                        *bool                          `xml:"urn:com.workday/bsvc Submit,omitempty"`
	Date                                          time.Time                      `xml:"urn:com.workday/bsvc Date"`
	CurrencyReference                             *CurrencyObjectType            `xml:"urn:com.workday/bsvc Currency_Reference"`
	Amount                                        float64                        `xml:"urn:com.workday/bsvc Amount"`
	EliminateForeignExchangeGainorLoss            *bool                          `xml:"urn:com.workday/bsvc Eliminate_Foreign_Exchange_Gain_or_Loss,omitempty"`
	Memo                                          string                         `xml:"urn:com.workday/bsvc Memo,omitempty"`
	Transaction                                   string                         `xml:"urn:com.workday/bsvc Transaction,omitempty"`
	CompanyReference                              *CompanyObjectType             `xml:"urn:com.workday/bsvc Company_Reference"`
	FromAccountReference                          *FinancialAccountObjectType    `xml:"urn:com.workday/bsvc From_Account_Reference"`
	FromAccountCurrencyReference                  *CurrencyObjectType            `xml:"urn:com.workday/bsvc From_Account_Currency_Reference,omitempty"`
	FromAccountCurrencyRateTypeReference          *CurrencyRateTypeObjectType    `xml:"urn:com.workday/bsvc From_Account_Currency_Rate_Type_Reference,omitempty"`
	FromAccountCurrencyConversionRate             float64                        `xml:"urn:com.workday/bsvc From_Account_Currency_Conversion_Rate,omitempty"`
	FromAccountBalancingWorktagReference          *BalancingWorktagObjectType    `xml:"urn:com.workday/bsvc From_Account_Balancing_Worktag_Reference,omitempty"`
	OptionalBalancingWorktagsfromAccountReference []AccountingWorktagObjectType  `xml:"urn:com.workday/bsvc Optional_Balancing_Worktags_from_Account_Reference,omitempty"`
	ToCompanyReference                            *CompanyObjectType             `xml:"urn:com.workday/bsvc To_Company_Reference"`
	ToAccountReference                            *FinancialAccountObjectType    `xml:"urn:com.workday/bsvc To_Account_Reference"`
	ToAccountCurrencyReference                    *CurrencyObjectType            `xml:"urn:com.workday/bsvc To_Account_Currency_Reference,omitempty"`
	ToAccountCurrencyRateTypeReference            *CurrencyRateTypeObjectType    `xml:"urn:com.workday/bsvc To_Account_Currency_Rate_Type_Reference,omitempty"`
	ToAccountCurrencyConversionRate               float64                        `xml:"urn:com.workday/bsvc To_Account_Currency_Conversion_Rate,omitempty"`
	ToAccountBalancingWorktagReference            *BalancingWorktagObjectType    `xml:"urn:com.workday/bsvc To_Account_Balancing_Worktag_Reference,omitempty"`
	OptionalBalancingWorktagsToAccountReference   []AccountingWorktagObjectType  `xml:"urn:com.workday/bsvc Optional_Balancing_Worktags_To_Account_Reference,omitempty"`
	AttachmentData                                []FinancialsAttachmentDataType `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
}

func (t *BankAccountTransferDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankAccountTransferDataType
	var layout struct {
		*T
		Date *xsdDate `xml:"urn:com.workday/bsvc Date"`
	}
	layout.T = (*T)(t)
	layout.Date = (*xsdDate)(&layout.T.Date)
	return e.EncodeElement(layout, start)
}
func (t *BankAccountTransferDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankAccountTransferDataType
	var overlay struct {
		*T
		Date *xsdDate `xml:"urn:com.workday/bsvc Date"`
	}
	overlay.T = (*T)(t)
	overlay.Date = (*xsdDate)(&overlay.T.Date)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type BankAccountTransferObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankAccountTransferObjectType struct {
	ID         []BankAccountTransferObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing search criteria for Bank Acount Transfers
type BankAccountTransferRequestCriteriaType struct {
	CompanyReference       []OrganizationObjectType     `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	BankAccountReference   []FinancialAccountObjectType `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	TransactionOnorAfter   *time.Time                   `xml:"urn:com.workday/bsvc Transaction_On_or_After,omitempty"`
	TransactionsOnorBefore *time.Time                   `xml:"urn:com.workday/bsvc Transactions_On_or_Before,omitempty"`
}

func (t *BankAccountTransferRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankAccountTransferRequestCriteriaType
	var layout struct {
		*T
		TransactionOnorAfter   *xsdDate `xml:"urn:com.workday/bsvc Transaction_On_or_After,omitempty"`
		TransactionsOnorBefore *xsdDate `xml:"urn:com.workday/bsvc Transactions_On_or_Before,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TransactionOnorAfter = (*xsdDate)(layout.T.TransactionOnorAfter)
	layout.TransactionsOnorBefore = (*xsdDate)(layout.T.TransactionsOnorBefore)
	return e.EncodeElement(layout, start)
}
func (t *BankAccountTransferRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankAccountTransferRequestCriteriaType
	var overlay struct {
		*T
		TransactionOnorAfter   *xsdDate `xml:"urn:com.workday/bsvc Transaction_On_or_After,omitempty"`
		TransactionsOnorBefore *xsdDate `xml:"urn:com.workday/bsvc Transactions_On_or_Before,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionOnorAfter = (*xsdDate)(overlay.T.TransactionOnorAfter)
	overlay.TransactionsOnorBefore = (*xsdDate)(overlay.T.TransactionsOnorBefore)
	return d.DecodeElement(&overlay, &start)
}

// References to existing Bank Account Transfers.
type BankAccountTransferRequestReferencesType struct {
	BankAccountTransferReference []BankAccountTransferObjectType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_Reference,omitempty"`
}

// WSDL Documentation Word Bucket Element containing data for existing Bank Account Transfers.
type BankAccountTransferResponseDataType struct {
	BankAccountTransfer []BankAccountTransferType `xml:"urn:com.workday/bsvc Bank_Account_Transfer,omitempty"`
}

// Element containing bank account response grouping options
type BankAccountTransferResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Bank Account Transfer Template data element
type BankAccountTransferTemplateDataType struct {
	BankAccountTransactionTemplateID              string                        `xml:"urn:com.workday/bsvc Bank_Account_Transaction_Template_ID,omitempty"`
	Name                                          string                        `xml:"urn:com.workday/bsvc Name"`
	CurrencyReference                             *CurrencyObjectType           `xml:"urn:com.workday/bsvc Currency_Reference"`
	EliminateForeignExchangeGainorLoss            *bool                         `xml:"urn:com.workday/bsvc Eliminate_Foreign_Exchange_Gain_or_Loss,omitempty"`
	Memo                                          string                        `xml:"urn:com.workday/bsvc Memo,omitempty"`
	Inactive                                      *bool                         `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	FromCompanyReference                          *CompanyObjectType            `xml:"urn:com.workday/bsvc From_Company_Reference"`
	FromFinancialAccountReference                 *FinancialAccountObjectType   `xml:"urn:com.workday/bsvc From_Financial_Account_Reference"`
	FromAccountCurrencyReference                  *CurrencyObjectType           `xml:"urn:com.workday/bsvc From_Account_Currency_Reference,omitempty"`
	FromCurrencyRateTypeReference                 *CurrencyRateTypeObjectType   `xml:"urn:com.workday/bsvc From_Currency_Rate_Type_Reference,omitempty"`
	FromBalancingWorktagReference                 *BalancingWorktagObjectType   `xml:"urn:com.workday/bsvc From_Balancing_Worktag_Reference,omitempty"`
	OptionalBalancingWorktagsfromAccountReference []AccountingWorktagObjectType `xml:"urn:com.workday/bsvc Optional_Balancing_Worktags_from_Account_Reference,omitempty"`
	ToCompanyReference                            *CompanyObjectType            `xml:"urn:com.workday/bsvc To_Company_Reference"`
	ToFinancialAccountReference                   *FinancialAccountObjectType   `xml:"urn:com.workday/bsvc To_Financial_Account_Reference"`
	ToAccountCurrencyReference                    *CurrencyObjectType           `xml:"urn:com.workday/bsvc To_Account_Currency_Reference,omitempty"`
	ToCurrencyRateTypeReference                   *CurrencyRateTypeObjectType   `xml:"urn:com.workday/bsvc To_Currency_Rate_Type_Reference,omitempty"`
	ToBalancingWorktagReference                   *BalancingWorktagObjectType   `xml:"urn:com.workday/bsvc To_Balancing_Worktag_Reference,omitempty"`
	OptionalBalancingWorktagsToAccountReference   []AccountingWorktagObjectType `xml:"urn:com.workday/bsvc Optional_Balancing_Worktags_To_Account_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankAccountTransferTemplateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankAccountTransferTemplateObjectType struct {
	ID         []BankAccountTransferTemplateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Bank Account Transfer Template request criteria element
type BankAccountTransferTemplateRequestCriteriaType struct {
}

// Bank Account Transfer Template request references element
type BankAccountTransferTemplateRequestReferencesType struct {
	BankAccountTransferTemplateReference []BankAccountTransferTemplateObjectType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_Template_Reference,omitempty"`
}

// Bank Account Transfer Template response data element
type BankAccountTransferTemplateResponseDataType struct {
	BankAccountTransferTemplate []BankAccountTransferTemplateType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_Template,omitempty"`
}

// Bank Account Transfer Template response group element
type BankAccountTransferTemplateResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Bank Account Transfer Template element
type BankAccountTransferTemplateType struct {
	BankAccountTransferTemplateReference *BankAccountTransferTemplateObjectType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_Template_Reference,omitempty"`
	BankAccountTransferTemplateData      *BankAccountTransferTemplateDataType   `xml:"urn:com.workday/bsvc Bank_Account_Transfer_Template_Data,omitempty"`
}

// Element containing reference to an existing Bank Account Transfer and data
type BankAccountTransferType struct {
	BankAccountTransferReference *BankAccountTransferObjectType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_Reference,omitempty"`
	BankAccountTransferData      []BankAccountTransferDataType  `xml:"urn:com.workday/bsvc Bank_Account_Transfer_Data,omitempty"`
}

// Element containing all Bank Account Transfer for Settlement data
type BankAccountTransferforSettlementDataType struct {
	BankAccountTransferforSettlementID     string                         `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_ID,omitempty"`
	Submit                                 *bool                          `xml:"urn:com.workday/bsvc Submit,omitempty"`
	Date                                   time.Time                      `xml:"urn:com.workday/bsvc Date"`
	CurrencyReference                      *CurrencyObjectType            `xml:"urn:com.workday/bsvc Currency_Reference"`
	Amount                                 float64                        `xml:"urn:com.workday/bsvc Amount"`
	PaymentTypeReference                   *PaymentTypeObjectType         `xml:"urn:com.workday/bsvc Payment_Type_Reference,omitempty"`
	EliminateFXGainorLoss                  *bool                          `xml:"urn:com.workday/bsvc Eliminate_FX_Gain_or_Loss,omitempty"`
	Memo                                   string                         `xml:"urn:com.workday/bsvc Memo,omitempty"`
	CompanyReference                       *CompanyObjectType             `xml:"urn:com.workday/bsvc Company_Reference"`
	FromAccountReference                   *FinancialAccountObjectType    `xml:"urn:com.workday/bsvc From_Account_Reference"`
	FromAccountCurrencyReference           *CurrencyObjectType            `xml:"urn:com.workday/bsvc From_Account_Currency_Reference,omitempty"`
	FromAccountCurrencyRateTypeReference   *CurrencyRateTypeObjectType    `xml:"urn:com.workday/bsvc From_Account_Currency_Rate_Type_Reference,omitempty"`
	FromAccountCurrencyConversionRate      float64                        `xml:"urn:com.workday/bsvc From_Account_Currency_Conversion_Rate,omitempty"`
	FromAccountBalancingWorktagReference   *BalancingWorktagObjectType    `xml:"urn:com.workday/bsvc From_Account_Balancing_Worktag_Reference,omitempty"`
	FromOptionalBalancingWorktagsReference []AccountingWorktagObjectType  `xml:"urn:com.workday/bsvc From_Optional_Balancing_Worktags_Reference,omitempty"`
	ToCompanyReference                     *CompanyObjectType             `xml:"urn:com.workday/bsvc To_Company_Reference"`
	ToAccountReference                     *BankAccountObjectType         `xml:"urn:com.workday/bsvc To_Account_Reference"`
	ToAccountCurrencyReference             *CurrencyObjectType            `xml:"urn:com.workday/bsvc To_Account_Currency_Reference,omitempty"`
	ToAccountCurrencyRateTypeReference     *CurrencyRateTypeObjectType    `xml:"urn:com.workday/bsvc To_Account_Currency_Rate_Type_Reference,omitempty"`
	ToAccountCurrencyConversionRate        float64                        `xml:"urn:com.workday/bsvc To_Account_Currency_Conversion_Rate,omitempty"`
	ToAccountBalancingWorktagReference     *BalancingWorktagObjectType    `xml:"urn:com.workday/bsvc To_Account_Balancing_Worktag_Reference,omitempty"`
	ToAccountOptionalWorktagsReference     []AccountingWorktagObjectType  `xml:"urn:com.workday/bsvc To_Account_Optional_Worktags_Reference,omitempty"`
	Attachment                             []FinancialsAttachmentDataType `xml:"urn:com.workday/bsvc Attachment,omitempty"`
}

func (t *BankAccountTransferforSettlementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankAccountTransferforSettlementDataType
	var layout struct {
		*T
		Date *xsdDate `xml:"urn:com.workday/bsvc Date"`
	}
	layout.T = (*T)(t)
	layout.Date = (*xsdDate)(&layout.T.Date)
	return e.EncodeElement(layout, start)
}
func (t *BankAccountTransferforSettlementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankAccountTransferforSettlementDataType
	var overlay struct {
		*T
		Date *xsdDate `xml:"urn:com.workday/bsvc Date"`
	}
	overlay.T = (*T)(t)
	overlay.Date = (*xsdDate)(&overlay.T.Date)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type BankAccountTransferforSettlementObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankAccountTransferforSettlementObjectType struct {
	ID         []BankAccountTransferforSettlementObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing search criteria for Get operation
type BankAccountTransferforSettlementRequestCriteriaType struct {
	CompanyReference          []OrganizationObjectType     `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	BankAccountReference      []FinancialAccountObjectType `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	TransactionDateOnorAfter  *time.Time                   `xml:"urn:com.workday/bsvc Transaction_Date_On_or_After,omitempty"`
	TransactionDateOnorBefore *time.Time                   `xml:"urn:com.workday/bsvc Transaction_Date_On_or_Before,omitempty"`
}

func (t *BankAccountTransferforSettlementRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankAccountTransferforSettlementRequestCriteriaType
	var layout struct {
		*T
		TransactionDateOnorAfter  *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date_On_or_After,omitempty"`
		TransactionDateOnorBefore *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date_On_or_Before,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TransactionDateOnorAfter = (*xsdDate)(layout.T.TransactionDateOnorAfter)
	layout.TransactionDateOnorBefore = (*xsdDate)(layout.T.TransactionDateOnorBefore)
	return e.EncodeElement(layout, start)
}
func (t *BankAccountTransferforSettlementRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankAccountTransferforSettlementRequestCriteriaType
	var overlay struct {
		*T
		TransactionDateOnorAfter  *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date_On_or_After,omitempty"`
		TransactionDateOnorBefore *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date_On_or_Before,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionDateOnorAfter = (*xsdDate)(overlay.T.TransactionDateOnorAfter)
	overlay.TransactionDateOnorBefore = (*xsdDate)(overlay.T.TransactionDateOnorBefore)
	return d.DecodeElement(&overlay, &start)
}

// Element specifies the Bank Account Transfer for Settlement reference for the operation
type BankAccountTransferforSettlementRequestReferencesType struct {
	BankAccountTransferforSettlementReference []BankAccountTransferforSettlementObjectType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Reference"`
}

// Element containing the Get operation response data
type BankAccountTransferforSettlementResponseDataType struct {
	BankAccountTransferforSettlement []BankAccountTransferforSettlementType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement,omitempty"`
}

// Element containing the response grouping options
type BankAccountTransferforSettlementResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// The data for the Bank Account Transfer for Settlement Template
type BankAccountTransferforSettlementTemplateDataType struct {
	BankAccountTransferforSettlementTemplateID    string                        `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Template_ID,omitempty"`
	Name                                          string                        `xml:"urn:com.workday/bsvc Name"`
	TransactionCurrencyReference                  *CurrencyObjectType           `xml:"urn:com.workday/bsvc Transaction_Currency_Reference"`
	PaymentTypeReference                          *PaymentTypeObjectType        `xml:"urn:com.workday/bsvc Payment_Type_Reference"`
	EliminateForeignExchangeGainorLoss            *bool                         `xml:"urn:com.workday/bsvc Eliminate_Foreign_Exchange_Gain_or_Loss,omitempty"`
	Memo                                          string                        `xml:"urn:com.workday/bsvc Memo,omitempty"`
	Inactive                                      *bool                         `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	FromCompanyReference                          *CompanyObjectType            `xml:"urn:com.workday/bsvc From_Company_Reference"`
	FromFinancialAccountReference                 *FinancialAccountObjectType   `xml:"urn:com.workday/bsvc From_Financial_Account_Reference"`
	FromAccountCurrencyReference                  *CurrencyObjectType           `xml:"urn:com.workday/bsvc From_Account_Currency_Reference,omitempty"`
	FromCurrencyRateTypeReference                 *CurrencyRateTypeObjectType   `xml:"urn:com.workday/bsvc From_Currency_Rate_Type_Reference,omitempty"`
	FromBalancingWorktagReference                 *BalancingWorktagObjectType   `xml:"urn:com.workday/bsvc From_Balancing_Worktag_Reference,omitempty"`
	OptionalBalancingWorktagsFromAccountReference []AccountingWorktagObjectType `xml:"urn:com.workday/bsvc Optional_Balancing_Worktags_From_Account_Reference,omitempty"`
	ToCompanyReference                            *CompanyObjectType            `xml:"urn:com.workday/bsvc To_Company_Reference"`
	ToFinancialAccountReference                   *FinancialAccountObjectType   `xml:"urn:com.workday/bsvc To_Financial_Account_Reference"`
	ToAccountCurrencyReference                    *CurrencyObjectType           `xml:"urn:com.workday/bsvc To_Account_Currency_Reference,omitempty"`
	ToCurrencyRateTypeReference                   *CurrencyRateTypeObjectType   `xml:"urn:com.workday/bsvc To_Currency_Rate_Type_Reference,omitempty"`
	ToBalancingWorktagReference                   *BalancingWorktagObjectType   `xml:"urn:com.workday/bsvc To_Balancing_Worktag_Reference,omitempty"`
	OptionalBalancingWorktagsToAccountReference   []AccountingWorktagObjectType `xml:"urn:com.workday/bsvc Optional_Balancing_Worktags_To_Account_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankAccountTransferforSettlementTemplateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankAccountTransferforSettlementTemplateObjectType struct {
	ID         []BankAccountTransferforSettlementTemplateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Request Criteria for Bank Account Transfer for Settlement Template
type BankAccountTransferforSettlementTemplateRequestCriteriaType struct {
	Name                 string                       `xml:"urn:com.workday/bsvc Name,omitempty"`
	CompanyReference     []CompanyObjectType          `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	BankAccountReference []FinancialAccountObjectType `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	CurrencyReference    []CurrencyObjectType         `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	StatusReference      []DocumentStatusObjectType   `xml:"urn:com.workday/bsvc Status_Reference,omitempty"`
}

// The reference instance for a Bank Account Transfer for Settlement Template
type BankAccountTransferforSettlementTemplateRequestReferencesType struct {
	BankAccountTransferforSettlementTemplateReference []BankAccountTransferforSettlementTemplateObjectType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Template_Reference"`
}

// The response for a Bank Account Transfer for Settlement Template
type BankAccountTransferforSettlementTemplateResponseDataType struct {
	BankAccountTransferforSettlementTemplate []BankAccountTransferforSettlementTemplateType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Template,omitempty"`
}

// The grouping options for the Bank Account Transfer for Settlement Template response
type BankAccountTransferforSettlementTemplateResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// The Bank Account Transfer for Settlement Template reference and its data
type BankAccountTransferforSettlementTemplateType struct {
	BankAccountTransferforSettlementTemplateReference *BankAccountTransferforSettlementTemplateObjectType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Template_Reference,omitempty"`
	BankAccountTransferforSettlementTemplateData      []BankAccountTransferforSettlementTemplateDataType  `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Template_Data,omitempty"`
}

// Element containing an existing Bank Account Transfer for Settlement
type BankAccountTransferforSettlementType struct {
	BankAccountTransferforSettlementReference *BankAccountTransferforSettlementObjectType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Reference,omitempty"`
	BankAccountTransferforSettlementData      []BankAccountTransferforSettlementDataType  `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Data,omitempty"`
}

// Element containing bank account reference and data
type BankAccountType struct {
	BankAccountReference *BankAccountObjectType `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	BankAccountData      []BankAccountDataType  `xml:"urn:com.workday/bsvc Bank_Account_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankAccountTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankAccountTypeObjectType struct {
	ID         []BankAccountTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing all Bank Branch data
type BankBranchDataType struct {
	BankBranchID                  string                          `xml:"urn:com.workday/bsvc Bank_Branch_ID,omitempty"`
	FinancialInstitutionReference *FinancialInstitutionObjectType `xml:"urn:com.workday/bsvc Financial_Institution_Reference"`
	BankBranchIDNumber            string                          `xml:"urn:com.workday/bsvc Bank_Branch_ID_Number,omitempty"`
	BusinessEntityData            *BusinessEntityWWSDataType      `xml:"urn:com.workday/bsvc Business_Entity_Data"`
}

// Contains a unique identifier for an instance of an object.
type BankBranchObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankBranchObjectType struct {
	ID         []BankBranchObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing Bank Account request criteria
type BankBranchRequestCriteriaType struct {
}

// Element containing reference instances for a Bank Branch
type BankBranchRequestReferencesType struct {
	BankBranchReference []BankBranchObjectType `xml:"urn:com.workday/bsvc Bank_Branch_Reference"`
}

// Wrapper Element that includes Bank Account Reference Instance and Bank Branch Data
type BankBranchResponseDataType struct {
	BankBranch []BankBranchType `xml:"urn:com.workday/bsvc Bank_Branch,omitempty"`
}

// Element containing Bank Branch response grouping options
type BankBranchResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing bank account reference and data
type BankBranchType struct {
	BankBranchReference *BankBranchObjectType `xml:"urn:com.workday/bsvc Bank_Branch_Reference,omitempty"`
	BankBranchData      []BankBranchDataType  `xml:"urn:com.workday/bsvc Bank_Branch_Data,omitempty"`
}

// Element containing the Compensation Code data.
type BankFeeCompensationCodeDataType struct {
	Description       string                             `xml:"urn:com.workday/bsvc Description,omitempty"`
	EDICode           string                             `xml:"urn:com.workday/bsvc EDI_Code,omitempty"`
	ISOCode           string                             `xml:"urn:com.workday/bsvc ISO_Code,omitempty"`
	CodeTypeReference *BankFeeCompensationTypeObjectType `xml:"urn:com.workday/bsvc Code_Type_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankFeeCompensationTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankFeeCompensationTypeObjectType struct {
	ID         []BankFeeCompensationTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankFeePricingTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankFeePricingTypeObjectType struct {
	ID         []BankFeePricingTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Bank Fee Service Code
type BankFeeServiceCodeDataType struct {
	BankFeeServiceCodeID          string                          `xml:"urn:com.workday/bsvc Bank_Fee_Service_Code_ID,omitempty"`
	FinancialInstitutionReference *FinancialInstitutionObjectType `xml:"urn:com.workday/bsvc Financial_Institution_Reference"`
	BankFeeServiceCode            string                          `xml:"urn:com.workday/bsvc Bank_Fee_Service_Code"`
	Description                   string                          `xml:"urn:com.workday/bsvc Description,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankFeeServiceCodeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankFeeServiceCodeObjectType struct {
	ID         []BankFeeServiceCodeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Bank Service Code Criteria
type BankFeeServiceCodeRequestCriteriaType struct {
}

// Bank Fee Service Code Request References
type BankFeeServiceCodeRequestReferencesType struct {
	BankFeeServiceCodeReference []BankFeeServiceCodeObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Service_Code_Reference"`
}

// Bank Fee Service Code Response Data
type BankFeeServiceCodeResponseDataType struct {
	BankFeeServiceCode []BankFeeServiceCodeType `xml:"urn:com.workday/bsvc Bank_Fee_Service_Code,omitempty"`
}

// Bank Fee Service Code Response Group
type BankFeeServiceCodeResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Bank Fee Service Code
type BankFeeServiceCodeType struct {
	BankFeeServiceCodeReference *BankFeeServiceCodeObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Service_Code_Reference,omitempty"`
	BankFeeServiceCodeData      []BankFeeServiceCodeDataType  `xml:"urn:com.workday/bsvc Bank_Fee_Service_Code_Data,omitempty"`
}

// Data for the incoming Service Contract.
type BankFeeServiceContractDataType struct {
	BankFeeServiceContractID        string                               `xml:"urn:com.workday/bsvc Bank_Fee_Service_Contract_ID,omitempty"`
	FinancialInstitutionReference   *FinancialInstitutionObjectType      `xml:"urn:com.workday/bsvc Financial_Institution_Reference"`
	BankAccountsReference           []BankAccountObjectType              `xml:"urn:com.workday/bsvc Bank_Accounts_Reference"`
	ServiceContractDescription      string                               `xml:"urn:com.workday/bsvc Service_Contract_Description"`
	BankFeeServiceContractLinesData []BankFeeServiceContractLineDataType `xml:"urn:com.workday/bsvc Bank_Fee_Service_Contract_Lines_Data,omitempty"`
	AttachmentData                  []FinancialsAttachmentDataType       `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
}

// The element container for the service contract line data.
type BankFeeServiceContractLineDataType struct {
	ServiceContractPricingReference         *BankFeeServiceContractPricingObjectType         `xml:"urn:com.workday/bsvc Service_Contract_Pricing_Reference,omitempty"`
	BankFeeServiceCodeReference             *BankFeeServiceCodeObjectType                    `xml:"urn:com.workday/bsvc Bank_Fee_Service_Code_Reference,omitempty"`
	ServiceCode                             string                                           `xml:"urn:com.workday/bsvc Service_Code"`
	ServiceCodeDescription                  string                                           `xml:"urn:com.workday/bsvc Service_Code_Description,omitempty"`
	BankFeePricingTypeReference             *BankFeePricingTypeObjectType                    `xml:"urn:com.workday/bsvc Bank_Fee_Pricing_Type_Reference"`
	ServiceContractPricingCurrencyReference *CurrencyObjectType                              `xml:"urn:com.workday/bsvc Service_Contract_Pricing_Currency_Reference"`
	BaseFee                                 float64                                          `xml:"urn:com.workday/bsvc Base_Fee,omitempty"`
	MinimumFee                              float64                                          `xml:"urn:com.workday/bsvc Minimum_Fee,omitempty"`
	MaximumFee                              float64                                          `xml:"urn:com.workday/bsvc Maximum_Fee,omitempty"`
	FixedPrice                              float64                                          `xml:"urn:com.workday/bsvc Fixed_Price,omitempty"`
	FlatPrice                               float64                                          `xml:"urn:com.workday/bsvc Flat_Price,omitempty"`
	ServiceContractTierPricingData          []BankFeeServiceContractPricingTierDataType      `xml:"urn:com.workday/bsvc Service_Contract_Tier_Pricing_Data,omitempty"`
	ServiceContractThresholdPricingData     []BankFeeServiceContractPricingThresholdDataType `xml:"urn:com.workday/bsvc Service_Contract_Threshold_Pricing_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankFeeServiceContractObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankFeeServiceContractObjectType struct {
	ID         []BankFeeServiceContractObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankFeeServiceContractPricingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankFeeServiceContractPricingObjectType struct {
	ID         []BankFeeServiceContractPricingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The element container for the Service Contract Threshold Data
type BankFeeServiceContractPricingThresholdDataType struct {
	ServiceContractThresholdPricingReference *BankFeeServiceContractThresholdPricingObjectType `xml:"urn:com.workday/bsvc Service_Contract_Threshold_Pricing_Reference,omitempty"`
	ThresholdMinimum                         float64                                           `xml:"urn:com.workday/bsvc Threshold_Minimum,omitempty"`
	ThresholdPrice                           float64                                           `xml:"urn:com.workday/bsvc Threshold_Price,omitempty"`
}

// The element container for the contract line.
type BankFeeServiceContractPricingTierDataType struct {
	ServiceContractTieredPricingReference    *BankFeeServiceContractTieredPricingObjectType `xml:"urn:com.workday/bsvc Service_Contract_Tiered_Pricing_Reference,omitempty"`
	ServiceContractPricingTierLevelReference *BankFeeTierPricingLevelObjectType             `xml:"urn:com.workday/bsvc Service_Contract_Pricing_Tier_Level_Reference,omitempty"`
	MinimumQuantity                          float64                                        `xml:"urn:com.workday/bsvc Minimum_Quantity,omitempty"`
	MaximumQuantity                          float64                                        `xml:"urn:com.workday/bsvc Maximum_Quantity,omitempty"`
	TierPrice                                float64                                        `xml:"urn:com.workday/bsvc Tier_Price,omitempty"`
}

// Bank Fee Service Contract Request Criteria
type BankFeeServiceContractRequestCriteriaType struct {
}

// Get Bank Fee Service Contract Reference
type BankFeeServiceContractRequestReferencesType struct {
	ServiceContractReference []BankFeeServiceContractObjectType `xml:"urn:com.workday/bsvc Service_Contract_Reference"`
}

// Get Bank Fee Service Contract Response
type BankFeeServiceContractResponseDataType struct {
	BankFeeServiceContract []BankFeeServiceContractType `xml:"urn:com.workday/bsvc Bank_Fee_Service_Contract,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankFeeServiceContractThresholdPricingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankFeeServiceContractThresholdPricingObjectType struct {
	ID         []BankFeeServiceContractThresholdPricingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankFeeServiceContractTieredPricingObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankFeeServiceContractTieredPricingObjectType struct {
	ID         []BankFeeServiceContractTieredPricingObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Get Bank Fee Service Contract Response
type BankFeeServiceContractType struct {
	ServiceContractReference   *BankFeeServiceContractObjectType `xml:"urn:com.workday/bsvc Service_Contract_Reference,omitempty"`
	BankFeeServiceContractData []BankFeeServiceContractDataType  `xml:"urn:com.workday/bsvc Bank_Fee_Service_Contract_Data,omitempty"`
}

// Element containing the compensation data, rate information.
type BankFeeStatementCompensationDataType struct {
	BankFeeCompensationCodeData []BankFeeCompensationCodeDataType `xml:"urn:com.workday/bsvc Bank_Fee_Compensation_Code_Data,omitempty"`
	CompensationAmount          float64                           `xml:"urn:com.workday/bsvc Compensation_Amount,omitempty"`
	CompensationMultiplier      float64                           `xml:"urn:com.workday/bsvc Compensation_Multiplier,omitempty"`
	DaysinPeriod                float64                           `xml:"urn:com.workday/bsvc Days_in_Period,omitempty"`
	DaysinYear                  float64                           `xml:"urn:com.workday/bsvc Days_in_Year,omitempty"`
}

// Element containing the compensation data, rate information.
type BankFeeStatementCompensationHVDataType struct {
	EDICode                string                             `xml:"urn:com.workday/bsvc EDI_Code,omitempty"`
	ISOCode                string                             `xml:"urn:com.workday/bsvc ISO_Code,omitempty"`
	CodeTypeReference      *BankFeeCompensationTypeObjectType `xml:"urn:com.workday/bsvc Code_Type_Reference"`
	CompensationAmount     float64                            `xml:"urn:com.workday/bsvc Compensation_Amount,omitempty"`
	CompensationMultiplier float64                            `xml:"urn:com.workday/bsvc Compensation_Multiplier,omitempty"`
	DaysinPeriod           float64                            `xml:"urn:com.workday/bsvc Days_in_Period,omitempty"`
	DaysinYear             float64                            `xml:"urn:com.workday/bsvc Days_in_Year,omitempty"`
}

// Specifies details related to currency exchange data.
type BankFeeStatementCurrencyConversionRateDataType struct {
	CurrencyRateTypeReference *CurrencyRateTypeObjectType `xml:"urn:com.workday/bsvc Currency_Rate_Type_Reference,omitempty"`
	SourceCurrencyReference   *CurrencyObjectType         `xml:"urn:com.workday/bsvc Source_Currency_Reference,omitempty"`
	TargetCurrencyReference   *CurrencyObjectType         `xml:"urn:com.workday/bsvc Target_Currency_Reference,omitempty"`
	CurrencyExchangeRateDate  *time.Time                  `xml:"urn:com.workday/bsvc Currency_Exchange_Rate_Date,omitempty"`
	CurrencyExchangeRate      float64                     `xml:"urn:com.workday/bsvc Currency_Exchange_Rate,omitempty"`
	Description               string                      `xml:"urn:com.workday/bsvc Description,omitempty"`
	Comments                  string                      `xml:"urn:com.workday/bsvc Comments,omitempty"`
}

func (t *BankFeeStatementCurrencyConversionRateDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankFeeStatementCurrencyConversionRateDataType
	var layout struct {
		*T
		CurrencyExchangeRateDate *xsdDate `xml:"urn:com.workday/bsvc Currency_Exchange_Rate_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CurrencyExchangeRateDate = (*xsdDate)(layout.T.CurrencyExchangeRateDate)
	return e.EncodeElement(layout, start)
}
func (t *BankFeeStatementCurrencyConversionRateDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankFeeStatementCurrencyConversionRateDataType
	var overlay struct {
		*T
		CurrencyExchangeRateDate *xsdDate `xml:"urn:com.workday/bsvc Currency_Exchange_Rate_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CurrencyExchangeRateDate = (*xsdDate)(overlay.T.CurrencyExchangeRateDate)
	return d.DecodeElement(&overlay, &start)
}

// Specifies details related to currency exchange data.
type BankFeeStatementCurrencyConversionRateHVDataType struct {
	SourceCurrencyReference  *CurrencyObjectType `xml:"urn:com.workday/bsvc Source_Currency_Reference"`
	TargetCurrencyReference  *CurrencyObjectType `xml:"urn:com.workday/bsvc Target_Currency_Reference"`
	CurrencyExchangeRateDate *time.Time          `xml:"urn:com.workday/bsvc Currency_Exchange_Rate_Date,omitempty"`
	CurrencyExchangeRate     float64             `xml:"urn:com.workday/bsvc Currency_Exchange_Rate"`
	Description              string              `xml:"urn:com.workday/bsvc Description,omitempty"`
	Comments                 string              `xml:"urn:com.workday/bsvc Comments,omitempty"`
}

func (t *BankFeeStatementCurrencyConversionRateHVDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankFeeStatementCurrencyConversionRateHVDataType
	var layout struct {
		*T
		CurrencyExchangeRateDate *xsdDate `xml:"urn:com.workday/bsvc Currency_Exchange_Rate_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CurrencyExchangeRateDate = (*xsdDate)(layout.T.CurrencyExchangeRateDate)
	return e.EncodeElement(layout, start)
}
func (t *BankFeeStatementCurrencyConversionRateHVDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankFeeStatementCurrencyConversionRateHVDataType
	var overlay struct {
		*T
		CurrencyExchangeRateDate *xsdDate `xml:"urn:com.workday/bsvc Currency_Exchange_Rate_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CurrencyExchangeRateDate = (*xsdDate)(overlay.T.CurrencyExchangeRateDate)
	return d.DecodeElement(&overlay, &start)
}

// Element containing all Bank Fee Statement data.
type BankFeeStatementDataType struct {
	BankFeeStatementID                         string                                           `xml:"urn:com.workday/bsvc Bank_Fee_Statement_ID,omitempty"`
	BankFeeStatementFileReference              *BankFeeStatementFileObjectType                  `xml:"urn:com.workday/bsvc Bank_Fee_Statement_File_Reference,omitempty"`
	StatementDate                              *time.Time                                       `xml:"urn:com.workday/bsvc Statement_Date,omitempty"`
	StatementBeginDate                         *time.Time                                       `xml:"urn:com.workday/bsvc Statement_Begin_Date,omitempty"`
	StatementEndDate                           *time.Time                                       `xml:"urn:com.workday/bsvc Statement_End_Date,omitempty"`
	CustomerAccountNumber                      string                                           `xml:"urn:com.workday/bsvc Customer_Account_Number,omitempty"`
	TaxCurrencyReference                       *CurrencyObjectType                              `xml:"urn:com.workday/bsvc Tax_Currency_Reference,omitempty"`
	SettlementCurrencyReference                *CurrencyObjectType                              `xml:"urn:com.workday/bsvc Settlement_Currency_Reference,omitempty"`
	TaxRegion                                  string                                           `xml:"urn:com.workday/bsvc Tax_Region,omitempty"`
	BankFeeStatementCurrencyConversionRateData []BankFeeStatementCurrencyConversionRateDataType `xml:"urn:com.workday/bsvc Bank_Fee_Statement_Currency_Conversion_Rate_Data,omitempty"`
	BankFeeStatementCompensationData           []BankFeeStatementCompensationDataType           `xml:"urn:com.workday/bsvc Bank_Fee_Statement_Compensation_Data,omitempty"`
	BankFeeStatementTransactionData            []BankFeeStatementTransactionDataType            `xml:"urn:com.workday/bsvc Bank_Fee_Statement_Transaction_Data,omitempty"`
}

func (t *BankFeeStatementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankFeeStatementDataType
	var layout struct {
		*T
		StatementDate      *xsdDate `xml:"urn:com.workday/bsvc Statement_Date,omitempty"`
		StatementBeginDate *xsdDate `xml:"urn:com.workday/bsvc Statement_Begin_Date,omitempty"`
		StatementEndDate   *xsdDate `xml:"urn:com.workday/bsvc Statement_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StatementDate = (*xsdDate)(layout.T.StatementDate)
	layout.StatementBeginDate = (*xsdDate)(layout.T.StatementBeginDate)
	layout.StatementEndDate = (*xsdDate)(layout.T.StatementEndDate)
	return e.EncodeElement(layout, start)
}
func (t *BankFeeStatementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankFeeStatementDataType
	var overlay struct {
		*T
		StatementDate      *xsdDate `xml:"urn:com.workday/bsvc Statement_Date,omitempty"`
		StatementBeginDate *xsdDate `xml:"urn:com.workday/bsvc Statement_Begin_Date,omitempty"`
		StatementEndDate   *xsdDate `xml:"urn:com.workday/bsvc Statement_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StatementDate = (*xsdDate)(overlay.T.StatementDate)
	overlay.StatementBeginDate = (*xsdDate)(overlay.T.StatementBeginDate)
	overlay.StatementEndDate = (*xsdDate)(overlay.T.StatementEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Element contailing an instance of Bank Fee Statement File Data
type BankFeeStatementFileDataType struct {
	ID                       string     `xml:"urn:com.workday/bsvc ID,omitempty"`
	SenderIdentification     string     `xml:"urn:com.workday/bsvc Sender_Identification,omitempty"`
	ReceiverIdentification   string     `xml:"urn:com.workday/bsvc Receiver_Identification,omitempty"`
	FileCreationDate         *time.Time `xml:"urn:com.workday/bsvc File_Creation_Date,omitempty"`
	FileIdentificationNumber string     `xml:"urn:com.workday/bsvc File_Identification_Number,omitempty"`
	SkipValidations          *bool      `xml:"urn:com.workday/bsvc Skip_Validations,omitempty"`
}

func (t *BankFeeStatementFileDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankFeeStatementFileDataType
	var layout struct {
		*T
		FileCreationDate *xsdDate `xml:"urn:com.workday/bsvc File_Creation_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FileCreationDate = (*xsdDate)(layout.T.FileCreationDate)
	return e.EncodeElement(layout, start)
}
func (t *BankFeeStatementFileDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankFeeStatementFileDataType
	var overlay struct {
		*T
		FileCreationDate *xsdDate `xml:"urn:com.workday/bsvc File_Creation_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FileCreationDate = (*xsdDate)(overlay.T.FileCreationDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type BankFeeStatementFileObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankFeeStatementFileObjectType struct {
	ID         []BankFeeStatementFileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Filter the returned Bank Fee Statement Files retrieved using the following criteria.
type BankFeeStatementFileRequestCriteriaType struct {
	FileCreationDateFrom *time.Time `xml:"urn:com.workday/bsvc File_Creation_Date_From,omitempty"`
	FileCreationDateTo   *time.Time `xml:"urn:com.workday/bsvc File_Creation_Date_To,omitempty"`
}

func (t *BankFeeStatementFileRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankFeeStatementFileRequestCriteriaType
	var layout struct {
		*T
		FileCreationDateFrom *xsdDate `xml:"urn:com.workday/bsvc File_Creation_Date_From,omitempty"`
		FileCreationDateTo   *xsdDate `xml:"urn:com.workday/bsvc File_Creation_Date_To,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FileCreationDateFrom = (*xsdDate)(layout.T.FileCreationDateFrom)
	layout.FileCreationDateTo = (*xsdDate)(layout.T.FileCreationDateTo)
	return e.EncodeElement(layout, start)
}
func (t *BankFeeStatementFileRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankFeeStatementFileRequestCriteriaType
	var overlay struct {
		*T
		FileCreationDateFrom *xsdDate `xml:"urn:com.workday/bsvc File_Creation_Date_From,omitempty"`
		FileCreationDateTo   *xsdDate `xml:"urn:com.workday/bsvc File_Creation_Date_To,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FileCreationDateFrom = (*xsdDate)(overlay.T.FileCreationDateFrom)
	overlay.FileCreationDateTo = (*xsdDate)(overlay.T.FileCreationDateTo)
	return d.DecodeElement(&overlay, &start)
}

// Element containing all of the requested Bank Fee Statement Files.
type BankFeeStatementFileRequestReferencesType struct {
	BankFeeStatementFileReference []BankFeeStatementFileObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Statement_File_Reference"`
}

// Element containing all instances of Bank Fee Statement File.
type BankFeeStatementFileResponseDataType struct {
	BankFeeStatementFile []BankFeeStatementFileType `xml:"urn:com.workday/bsvc Bank_Fee_Statement_File,omitempty"`
}

// Wrapper element around a list of elements representing the amount of data that should be included in.
type BankFeeStatementFileResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing an instance of Bank Fee Statement File.
type BankFeeStatementFileType struct {
	BankFeeStatementFileReference *BankFeeStatementFileObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Statement_File_Reference,omitempty"`
	BankFeeStatementFileData      []BankFeeStatementFileDataType  `xml:"urn:com.workday/bsvc Bank_Fee_Statement_File_Data,omitempty"`
}

// Element containing all Bank Fee Statement data.
type BankFeeStatementHVDataType struct {
	BankFeeStatementID                         string                                             `xml:"urn:com.workday/bsvc Bank_Fee_Statement_ID,omitempty"`
	BankFeeStatementFileReference              *BankFeeStatementFileObjectType                    `xml:"urn:com.workday/bsvc Bank_Fee_Statement_File_Reference,omitempty"`
	RoutingNumber                              string                                             `xml:"urn:com.workday/bsvc Routing_Number,omitempty"`
	CreationDate                               time.Time                                          `xml:"urn:com.workday/bsvc Creation_Date"`
	StatementBeginDate                         time.Time                                          `xml:"urn:com.workday/bsvc Statement_Begin_Date"`
	StatementEndDate                           time.Time                                          `xml:"urn:com.workday/bsvc Statement_End_Date"`
	CustomerAccountNumber                      string                                             `xml:"urn:com.workday/bsvc Customer_Account_Number,omitempty"`
	TaxRegion                                  string                                             `xml:"urn:com.workday/bsvc Tax_Region,omitempty"`
	TaxCurrencyReference                       *CurrencyObjectType                                `xml:"urn:com.workday/bsvc Tax_Currency_Reference,omitempty"`
	SettlementCurrencyReference                *CurrencyObjectType                                `xml:"urn:com.workday/bsvc Settlement_Currency_Reference,omitempty"`
	BankFeeStatementCurrencyConversionRateData []BankFeeStatementCurrencyConversionRateHVDataType `xml:"urn:com.workday/bsvc Bank_Fee_Statement_Currency_Conversion_Rate_Data,omitempty"`
	BankFeeStatementCompensationData           []BankFeeStatementCompensationHVDataType           `xml:"urn:com.workday/bsvc Bank_Fee_Statement_Compensation_Data,omitempty"`
	BankFeeStatementTransactionData            []BankFeeStatementTransactionHVDataType            `xml:"urn:com.workday/bsvc Bank_Fee_Statement_Transaction_Data,omitempty"`
	AttachmentData                             []FinancialsAttachmentDataType                     `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
	BankIdentificationCode                     string                                             `xml:"urn:com.workday/bsvc Bank_Identification_Code,omitempty"`
	IBAN                                       string                                             `xml:"urn:com.workday/bsvc IBAN,omitempty"`
}

func (t *BankFeeStatementHVDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankFeeStatementHVDataType
	var layout struct {
		*T
		CreationDate       *xsdDate `xml:"urn:com.workday/bsvc Creation_Date"`
		StatementBeginDate *xsdDate `xml:"urn:com.workday/bsvc Statement_Begin_Date"`
		StatementEndDate   *xsdDate `xml:"urn:com.workday/bsvc Statement_End_Date"`
	}
	layout.T = (*T)(t)
	layout.CreationDate = (*xsdDate)(&layout.T.CreationDate)
	layout.StatementBeginDate = (*xsdDate)(&layout.T.StatementBeginDate)
	layout.StatementEndDate = (*xsdDate)(&layout.T.StatementEndDate)
	return e.EncodeElement(layout, start)
}
func (t *BankFeeStatementHVDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankFeeStatementHVDataType
	var overlay struct {
		*T
		CreationDate       *xsdDate `xml:"urn:com.workday/bsvc Creation_Date"`
		StatementBeginDate *xsdDate `xml:"urn:com.workday/bsvc Statement_Begin_Date"`
		StatementEndDate   *xsdDate `xml:"urn:com.workday/bsvc Statement_End_Date"`
	}
	overlay.T = (*T)(t)
	overlay.CreationDate = (*xsdDate)(&overlay.T.CreationDate)
	overlay.StatementBeginDate = (*xsdDate)(&overlay.T.StatementBeginDate)
	overlay.StatementEndDate = (*xsdDate)(&overlay.T.StatementEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type BankFeeStatementObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankFeeStatementObjectType struct {
	ID         []BankFeeStatementObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing criteria used to filter Bank Fee Statements.
type BankFeeStatementRequestCriteriaType struct {
	CompanyReference          []OrganizationObjectType     `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	FinancialAccountReference []FinancialAccountObjectType `xml:"urn:com.workday/bsvc Financial_Account_Reference,omitempty"`
	StatementDateFrom         *time.Time                   `xml:"urn:com.workday/bsvc Statement_Date_From,omitempty"`
	StatementDateTo           *time.Time                   `xml:"urn:com.workday/bsvc Statement_Date_To,omitempty"`
}

func (t *BankFeeStatementRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankFeeStatementRequestCriteriaType
	var layout struct {
		*T
		StatementDateFrom *xsdDate `xml:"urn:com.workday/bsvc Statement_Date_From,omitempty"`
		StatementDateTo   *xsdDate `xml:"urn:com.workday/bsvc Statement_Date_To,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StatementDateFrom = (*xsdDate)(layout.T.StatementDateFrom)
	layout.StatementDateTo = (*xsdDate)(layout.T.StatementDateTo)
	return e.EncodeElement(layout, start)
}
func (t *BankFeeStatementRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankFeeStatementRequestCriteriaType
	var overlay struct {
		*T
		StatementDateFrom *xsdDate `xml:"urn:com.workday/bsvc Statement_Date_From,omitempty"`
		StatementDateTo   *xsdDate `xml:"urn:com.workday/bsvc Statement_Date_To,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StatementDateFrom = (*xsdDate)(overlay.T.StatementDateFrom)
	overlay.StatementDateTo = (*xsdDate)(overlay.T.StatementDateTo)
	return d.DecodeElement(&overlay, &start)
}

// Element containing the Bank Fee Statement reference.
type BankFeeStatementRequestReferencesType struct {
	BankFeeStatementReference []BankFeeStatementObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Statement_Reference"`
}

// Wrapper element containing the Bank Fee Statement reference and its data.
type BankFeeStatementResponseDataType struct {
	BankFeeStatement []BankFeeStatementType `xml:"urn:com.workday/bsvc Bank_Fee_Statement,omitempty"`
}

// Element containing flags that determine the amount of data that is included in the Bank Fee Statements response. If this element is not included then ALL of the Response Groups are included in the response.
type BankFeeStatementResponseGroupType struct {
	IncludeReference            *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeBankFeeStatementData *bool `xml:"urn:com.workday/bsvc Include_Bank_Fee_Statement_Data,omitempty"`
}

// Element containing the Bank Fee Statement line data.
type BankFeeStatementTransactionDataType struct {
	AFPCode                                   string                             `xml:"urn:com.workday/bsvc AFP_Code,omitempty"`
	AFPCodeDescription                        string                             `xml:"urn:com.workday/bsvc AFP_Code_Description,omitempty"`
	BankFeeServiceCodeReference               *BankFeeServiceCodeObjectType      `xml:"urn:com.workday/bsvc Bank_Fee_Service_Code_Reference,omitempty"`
	BankFeePricingTypeReference               *BankFeePricingTypeObjectType      `xml:"urn:com.workday/bsvc Bank_Fee_Pricing_Type_Reference,omitempty"`
	BankFeeTierPricingIfTierSelectedReference *BankFeeTierPricingLevelObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Tier_Pricing_If_Tier_Selected_Reference,omitempty"`
	ServiceUnitPrice                          float64                            `xml:"urn:com.workday/bsvc Service_Unit_Price,omitempty"`
	ServiceVolume                             float64                            `xml:"urn:com.workday/bsvc Service_Volume,omitempty"`
	ServiceCharge                             float64                            `xml:"urn:com.workday/bsvc Service_Charge,omitempty"`
	TaxesPerService                           float64                            `xml:"urn:com.workday/bsvc Taxes_Per_Service,omitempty"`
	TotalCharges                              float64                            `xml:"urn:com.workday/bsvc Total_Charges,omitempty"`
}

// Element containing the Bank Fee Statement line data.
type BankFeeStatementTransactionHVDataType struct {
	AFPCode                                   string                             `xml:"urn:com.workday/bsvc AFP_Code,omitempty"`
	AFPCodeDescription                        string                             `xml:"urn:com.workday/bsvc AFP_Code_Description,omitempty"`
	BankFeeServiceCode                        string                             `xml:"urn:com.workday/bsvc Bank_Fee_Service_Code"`
	BankFeeServiceCodeDescription             string                             `xml:"urn:com.workday/bsvc Bank_Fee_Service_Code_Description,omitempty"`
	BankFeePricingTypeReference               *BankFeePricingTypeObjectType      `xml:"urn:com.workday/bsvc Bank_Fee_Pricing_Type_Reference"`
	BankFeeTierPricingIfTierSelectedReference *BankFeeTierPricingLevelObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Tier_Pricing_If_Tier_Selected_Reference,omitempty"`
	PricingCurrencyReference                  *CurrencyObjectType                `xml:"urn:com.workday/bsvc Pricing_Currency_Reference,omitempty"`
	ServiceUnitPrice                          float64                            `xml:"urn:com.workday/bsvc Service_Unit_Price,omitempty"`
	ServiceVolume                             float64                            `xml:"urn:com.workday/bsvc Service_Volume,omitempty"`
	ServiceCharge                             float64                            `xml:"urn:com.workday/bsvc Service_Charge,omitempty"`
	TaxesPerService                           float64                            `xml:"urn:com.workday/bsvc Taxes_Per_Service,omitempty"`
	TotalCharges                              float64                            `xml:"urn:com.workday/bsvc Total_Charges,omitempty"`
}

// Wrapper element containing the Bank Fee Statement reference and its data.
type BankFeeStatementType struct {
	BankFeeStatementReference *BankFeeStatementObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Statement_Reference,omitempty"`
	BankFeeStatementData      []BankFeeStatementDataType  `xml:"urn:com.workday/bsvc Bank_Fee_Statement_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankFeeTierPricingLevelObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankFeeTierPricingLevelObjectType struct {
	ID         []BankFeeTierPricingLevelObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing Bank Statement Balance Type Data
type BankStatementBalanceTypeDataType struct {
	BankStatementBalanceTypeReference  *BankStatementBalanceTypeObjectType  `xml:"urn:com.workday/bsvc Bank_Statement_Balance_Type_Reference,omitempty"`
	ID                                 string                               `xml:"urn:com.workday/bsvc ID,omitempty"`
	SummaryBalance                     string                               `xml:"urn:com.workday/bsvc Summary_Balance"`
	TypeCode                           string                               `xml:"urn:com.workday/bsvc Type_Code"`
	DebitCredit                        string                               `xml:"urn:com.workday/bsvc Debit_Credit,omitempty"`
	Description                        string                               `xml:"urn:com.workday/bsvc Description"`
	BankStatementTypeCategoryReference *BankStatementTypeCategoryObjectType `xml:"urn:com.workday/bsvc Bank_Statement_Type_Category_Reference"`
	BeginningBalance                   string                               `xml:"urn:com.workday/bsvc Beginning_Balance,omitempty"`
	EndingBalance                      string                               `xml:"urn:com.workday/bsvc Ending_Balance,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankStatementBalanceTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankStatementBalanceTypeObjectType struct {
	ID         []BankStatementBalanceTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Bank Statement Balance Type Reference.
type BankStatementBalanceTypeReferenceType struct {
	TypeCode string `xml:"urn:com.workday/bsvc Type_Code,omitempty"`
}

// Criteria Parms for Bank statements
type BankStatementCriteriaDataType struct {
	WID string `xml:"urn:com.workday/bsvc WID,omitempty"`
}

// Element containing all Bank Statement Custom Code Set Data including the subelements balance type data & transaction type data
type BankStatementCustomCodeSetDataType struct {
	ID                               string                                 `xml:"urn:com.workday/bsvc ID,omitempty"`
	Name                             string                                 `xml:"urn:com.workday/bsvc Name"`
	BankStatementFormatReference     *BankStatementFormatObjectType         `xml:"urn:com.workday/bsvc Bank_Statement_Format_Reference"`
	Description                      string                                 `xml:"urn:com.workday/bsvc Description,omitempty"`
	Comment                          string                                 `xml:"urn:com.workday/bsvc Comment,omitempty"`
	BankStatementBalanceTypeData     []BankStatementBalanceTypeDataType     `xml:"urn:com.workday/bsvc Bank_Statement_Balance_Type_Data,omitempty"`
	BankStatementTransactionTypeData []BankStatementTransactionTypeDataType `xml:"urn:com.workday/bsvc Bank_Statement_Transaction_Type_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankStatementCustomCodeSetObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankStatementCustomCodeSetObjectType struct {
	ID         []BankStatementCustomCodeSetObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing bank statement custom code set request criteria
type BankStatementCustomCodeSetRequestCriteriaType struct {
}

// Element containing reference instances for a bank statement custom code set
type BankStatementCustomCodeSetRequestReferencesType struct {
	BankStatementCustomCodeSetReference []BankStatementCustomCodeSetObjectType `xml:"urn:com.workday/bsvc Bank_Statement_Custom_Code_Set_Reference"`
}

// Wrapper Element that includes Bank Statement Custom Code Set Reference Instance and Bank Statement Custom Code Set Data
type BankStatementCustomCodeSetResponseDataType struct {
	BankStatementCustomCodeSet []BankStatementCustomCodeSetType `xml:"urn:com.workday/bsvc Bank_Statement_Custom_Code_Set,omitempty"`
}

// Element containing bank statement custom code set response grouping options
type BankStatementCustomCodeSetResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing bank statement custom code set reference and data
type BankStatementCustomCodeSetType struct {
	BankStatementCustomCodeSetReference *BankStatementCustomCodeSetObjectType `xml:"urn:com.workday/bsvc Bank_Statement_Custom_Code_Set_Reference,omitempty"`
	BankStatementCustomCodeSetData      *BankStatementCustomCodeSetDataType   `xml:"urn:com.workday/bsvc Bank_Statement_Custom_Code_Set_Data,omitempty"`
}

// Element containing all Bank Statement data
type BankStatementDataType struct {
	BankStatementReferenceID     string                             `xml:"urn:com.workday/bsvc Bank_Statement_Reference_ID,omitempty"`
	BankStatementFileReference   *BankStatementFileObjectType       `xml:"urn:com.workday/bsvc Bank_Statement_File_Reference,omitempty"`
	CustomerAccountNumber        string                             `xml:"urn:com.workday/bsvc Customer_Account_Number,omitempty"`
	RoutingNumber                string                             `xml:"urn:com.workday/bsvc Routing_Number,omitempty"`
	BankAccountReference         *BankAccountObjectType             `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	StatementBeginDate           *time.Time                         `xml:"urn:com.workday/bsvc Statement_Begin_Date,omitempty"`
	StatementEndDate             *time.Time                         `xml:"urn:com.workday/bsvc Statement_End_Date,omitempty"`
	SameDayBankStatement         *bool                              `xml:"urn:com.workday/bsvc Same_Day_Bank_Statement,omitempty"`
	CurrencyReference            *CurrencyObjectType                `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	BankStatementSummaryData     []BankStatementSummaryDataType     `xml:"urn:com.workday/bsvc Bank_Statement_Summary_Data,omitempty"`
	BankStatementTransactionData []BankStatementTransactionDataType `xml:"urn:com.workday/bsvc Bank_Statement_Transaction_Data,omitempty"`
	BankIdentificationCode       string                             `xml:"urn:com.workday/bsvc Bank_Identification_Code,omitempty"`
	IBAN                         string                             `xml:"urn:com.workday/bsvc IBAN,omitempty"`
}

func (t *BankStatementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankStatementDataType
	var layout struct {
		*T
		StatementBeginDate *xsdDate `xml:"urn:com.workday/bsvc Statement_Begin_Date,omitempty"`
		StatementEndDate   *xsdDate `xml:"urn:com.workday/bsvc Statement_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StatementBeginDate = (*xsdDate)(layout.T.StatementBeginDate)
	layout.StatementEndDate = (*xsdDate)(layout.T.StatementEndDate)
	return e.EncodeElement(layout, start)
}
func (t *BankStatementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankStatementDataType
	var overlay struct {
		*T
		StatementBeginDate *xsdDate `xml:"urn:com.workday/bsvc Statement_Begin_Date,omitempty"`
		StatementEndDate   *xsdDate `xml:"urn:com.workday/bsvc Statement_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StatementBeginDate = (*xsdDate)(overlay.T.StatementBeginDate)
	overlay.StatementEndDate = (*xsdDate)(overlay.T.StatementEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Element contailing an instance of Bank Statement File Data
type BankStatementFileDataType struct {
	BankStatementFileReferenceID string     `xml:"urn:com.workday/bsvc Bank_Statement_File_Reference_ID,omitempty"`
	SenderIdentification         string     `xml:"urn:com.workday/bsvc Sender_Identification,omitempty"`
	ReceiverIdentification       string     `xml:"urn:com.workday/bsvc Receiver_Identification,omitempty"`
	FileCreationDate             *time.Time `xml:"urn:com.workday/bsvc File_Creation_Date,omitempty"`
	FileIdentificationNumber     string     `xml:"urn:com.workday/bsvc File_Identification_Number,omitempty"`
	SkipValidations              *bool      `xml:"urn:com.workday/bsvc Skip_Validations,omitempty"`
}

func (t *BankStatementFileDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankStatementFileDataType
	var layout struct {
		*T
		FileCreationDate *xsdDate `xml:"urn:com.workday/bsvc File_Creation_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FileCreationDate = (*xsdDate)(layout.T.FileCreationDate)
	return e.EncodeElement(layout, start)
}
func (t *BankStatementFileDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankStatementFileDataType
	var overlay struct {
		*T
		FileCreationDate *xsdDate `xml:"urn:com.workday/bsvc File_Creation_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FileCreationDate = (*xsdDate)(overlay.T.FileCreationDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type BankStatementFileObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankStatementFileObjectType struct {
	ID         []BankStatementFileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Filter the returned Bank Statement Files retrieved using the following criteria.
type BankStatementFileRequestCriteriaType struct {
	FileCreationDateFrom *time.Time `xml:"urn:com.workday/bsvc File_Creation_Date_From,omitempty"`
	FileCreationDateTo   *time.Time `xml:"urn:com.workday/bsvc File_Creation_Date_To,omitempty"`
}

func (t *BankStatementFileRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankStatementFileRequestCriteriaType
	var layout struct {
		*T
		FileCreationDateFrom *xsdDate `xml:"urn:com.workday/bsvc File_Creation_Date_From,omitempty"`
		FileCreationDateTo   *xsdDate `xml:"urn:com.workday/bsvc File_Creation_Date_To,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FileCreationDateFrom = (*xsdDate)(layout.T.FileCreationDateFrom)
	layout.FileCreationDateTo = (*xsdDate)(layout.T.FileCreationDateTo)
	return e.EncodeElement(layout, start)
}
func (t *BankStatementFileRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankStatementFileRequestCriteriaType
	var overlay struct {
		*T
		FileCreationDateFrom *xsdDate `xml:"urn:com.workday/bsvc File_Creation_Date_From,omitempty"`
		FileCreationDateTo   *xsdDate `xml:"urn:com.workday/bsvc File_Creation_Date_To,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FileCreationDateFrom = (*xsdDate)(overlay.T.FileCreationDateFrom)
	overlay.FileCreationDateTo = (*xsdDate)(overlay.T.FileCreationDateTo)
	return d.DecodeElement(&overlay, &start)
}

// Element constaining all of the requested Bank Statement Files.
type BankStatementFileRequestReferencesType struct {
	BankStatementFileRequestReference []BankStatementFileObjectType `xml:"urn:com.workday/bsvc Bank_Statement_File_Request_Reference"`
}

// Element containing all instances of Bank Statement File.
type BankStatementFileResponseDataType struct {
	BankStatementFile []BankStatementFileType `xml:"urn:com.workday/bsvc Bank_Statement_File,omitempty"`
}

// Wrapper element around a list of elements representing the amount of data that should be included in the Bank Statement File response. If this element is not included then ALL of the Response Groups are included in the response.
type BankStatementFileResponseGroupType struct {
	IncludeReference             *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeBankStatementFileData *bool `xml:"urn:com.workday/bsvc Include_Bank_Statement_File_Data,omitempty"`
}

// Element containing an instance of Bank Statement File.
type BankStatementFileType struct {
	BankStatementFileReference *BankStatementFileObjectType `xml:"urn:com.workday/bsvc Bank_Statement_File_Reference,omitempty"`
	BankStatementFileData      []BankStatementFileDataType  `xml:"urn:com.workday/bsvc Bank_Statement_File_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankStatementFormatObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankStatementFormatObjectType struct {
	ID         []BankStatementFormatObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing all Bank Statement data
type BankStatementHVDataType struct {
	BankStatementReferenceID     string                               `xml:"urn:com.workday/bsvc Bank_Statement_Reference_ID,omitempty"`
	BankStatementFileReference   *BankStatementFileObjectType         `xml:"urn:com.workday/bsvc Bank_Statement_File_Reference,omitempty"`
	CustomerAccountNumber        string                               `xml:"urn:com.workday/bsvc Customer_Account_Number,omitempty"`
	RoutingNumber                string                               `xml:"urn:com.workday/bsvc Routing_Number,omitempty"`
	BankAccountReference         *BankAccountObjectType               `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	StatementBeginDate           *time.Time                           `xml:"urn:com.workday/bsvc Statement_Begin_Date,omitempty"`
	StatementEndDate             *time.Time                           `xml:"urn:com.workday/bsvc Statement_End_Date,omitempty"`
	CurrencyReference            *CurrencyObjectType                  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	BankStatementSummaryData     []BankStatementSummaryDataType       `xml:"urn:com.workday/bsvc Bank_Statement_Summary_Data,omitempty"`
	BankStatementTransactionData []BankStatementTransactionHVDataType `xml:"urn:com.workday/bsvc Bank_Statement_Transaction_Data,omitempty"`
	BankIdentificationCode       string                               `xml:"urn:com.workday/bsvc Bank_Identification_Code,omitempty"`
	IBAN                         string                               `xml:"urn:com.workday/bsvc IBAN,omitempty"`
}

func (t *BankStatementHVDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankStatementHVDataType
	var layout struct {
		*T
		StatementBeginDate *xsdDate `xml:"urn:com.workday/bsvc Statement_Begin_Date,omitempty"`
		StatementEndDate   *xsdDate `xml:"urn:com.workday/bsvc Statement_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StatementBeginDate = (*xsdDate)(layout.T.StatementBeginDate)
	layout.StatementEndDate = (*xsdDate)(layout.T.StatementEndDate)
	return e.EncodeElement(layout, start)
}
func (t *BankStatementHVDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankStatementHVDataType
	var overlay struct {
		*T
		StatementBeginDate *xsdDate `xml:"urn:com.workday/bsvc Statement_Begin_Date,omitempty"`
		StatementEndDate   *xsdDate `xml:"urn:com.workday/bsvc Statement_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StatementBeginDate = (*xsdDate)(overlay.T.StatementBeginDate)
	overlay.StatementEndDate = (*xsdDate)(overlay.T.StatementEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type BankStatementObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankStatementObjectType struct {
	ID         []BankStatementObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This element content contains the Bank Statements Response Data element.
type BankStatementResponseDataType struct {
	BankStatement []BankStatementType `xml:"urn:com.workday/bsvc Bank_Statement,omitempty"`
}

// Bank Statement Summary Data.
type BankStatementSummaryDataType struct {
	BankStatementBalanceTypeReference *BankStatementBalanceTypeReferenceType `xml:"urn:com.workday/bsvc Bank_Statement_Balance_Type_Reference,omitempty"`
	Amount                            float64                                `xml:"urn:com.workday/bsvc Amount,omitempty"`
	ItemCount                         float64                                `xml:"urn:com.workday/bsvc Item_Count,omitempty"`
	FundsAvailabilityData             []FundsAvailabilityDataType            `xml:"urn:com.workday/bsvc Funds_Availability_Data,omitempty"`
	AlternateBalanceTypeCode          string                                 `xml:"urn:com.workday/bsvc Alternate_Balance_Type_Code,omitempty"`
	Debit                             *bool                                  `xml:"urn:com.workday/bsvc Debit,omitempty"`
}

// Bank Statement Transaction Data.
type BankStatementTransactionDataType struct {
	BankStatementTransactionTypeReference *BankStatementTransactionTypeReferenceType `xml:"urn:com.workday/bsvc Bank_Statement_Transaction_Type_Reference,omitempty"`
	TransactionDate                       *time.Time                                 `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
	Amount                                float64                                    `xml:"urn:com.workday/bsvc Amount,omitempty"`
	BankReferenceNumber                   string                                     `xml:"urn:com.workday/bsvc Bank_Reference_Number,omitempty"`
	CustomerReferenceNumber               string                                     `xml:"urn:com.workday/bsvc Customer_Reference_Number,omitempty"`
	Text                                  string                                     `xml:"urn:com.workday/bsvc Text,omitempty"`
	FundsAvailabilityData                 []FundsAvailabilityDataType                `xml:"urn:com.workday/bsvc Funds_Availability_Data,omitempty"`
	AlternateTransactionTypeCode          string                                     `xml:"urn:com.workday/bsvc Alternate_Transaction_Type_Code,omitempty"`
	OriginatingCurrencyAmount             float64                                    `xml:"urn:com.workday/bsvc Originating_Currency_Amount,omitempty"`
	OriginatingCurrency                   string                                     `xml:"urn:com.workday/bsvc Originating_Currency,omitempty"`
	CurrencyExchangeRate                  float64                                    `xml:"urn:com.workday/bsvc Currency_Exchange_Rate,omitempty"`
	Debit                                 *bool                                      `xml:"urn:com.workday/bsvc Debit,omitempty"`
	ValueDate                             *time.Time                                 `xml:"urn:com.workday/bsvc Value_Date,omitempty"`
	ReversalIndicator                     *bool                                      `xml:"urn:com.workday/bsvc Reversal_Indicator,omitempty"`
	AdditionalReferenceField1             string                                     `xml:"urn:com.workday/bsvc Additional_Reference_Field_1,omitempty"`
	AdditionalReferenceField2             string                                     `xml:"urn:com.workday/bsvc Additional_Reference_Field_2,omitempty"`
	AdditionalReferenceField3             string                                     `xml:"urn:com.workday/bsvc Additional_Reference_Field_3,omitempty"`
	AdditionalReferenceField4             string                                     `xml:"urn:com.workday/bsvc Additional_Reference_Field_4,omitempty"`
	AdditionalReferenceField5             string                                     `xml:"urn:com.workday/bsvc Additional_Reference_Field_5,omitempty"`
}

func (t *BankStatementTransactionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankStatementTransactionDataType
	var layout struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
		ValueDate       *xsdDate `xml:"urn:com.workday/bsvc Value_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TransactionDate = (*xsdDate)(layout.T.TransactionDate)
	layout.ValueDate = (*xsdDate)(layout.T.ValueDate)
	return e.EncodeElement(layout, start)
}
func (t *BankStatementTransactionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankStatementTransactionDataType
	var overlay struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
		ValueDate       *xsdDate `xml:"urn:com.workday/bsvc Value_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionDate = (*xsdDate)(overlay.T.TransactionDate)
	overlay.ValueDate = (*xsdDate)(overlay.T.ValueDate)
	return d.DecodeElement(&overlay, &start)
}

// Bank Statement Transaction Data.
type BankStatementTransactionHVDataType struct {
	BankStatementTransactionTypeReference *BankStatementTransactionTypeReferenceType `xml:"urn:com.workday/bsvc Bank_Statement_Transaction_Type_Reference,omitempty"`
	TransactionDate                       *time.Time                                 `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
	Amount                                float64                                    `xml:"urn:com.workday/bsvc Amount,omitempty"`
	BankReferenceNumber                   string                                     `xml:"urn:com.workday/bsvc Bank_Reference_Number,omitempty"`
	CustomerReferenceNumber               string                                     `xml:"urn:com.workday/bsvc Customer_Reference_Number,omitempty"`
	Text                                  string                                     `xml:"urn:com.workday/bsvc Text,omitempty"`
	FundsAvailabilityData                 []FundsAvailabilityDataType                `xml:"urn:com.workday/bsvc Funds_Availability_Data,omitempty"`
	AlternateTransactionTypeCode          string                                     `xml:"urn:com.workday/bsvc Alternate_Transaction_Type_Code,omitempty"`
	OriginatingCurrencyAmount             float64                                    `xml:"urn:com.workday/bsvc Originating_Currency_Amount,omitempty"`
	OriginatingCurrency                   string                                     `xml:"urn:com.workday/bsvc Originating_Currency,omitempty"`
	CurrencyExchangeRate                  float64                                    `xml:"urn:com.workday/bsvc Currency_Exchange_Rate,omitempty"`
	Debit                                 *bool                                      `xml:"urn:com.workday/bsvc Debit,omitempty"`
	ValueDate                             *time.Time                                 `xml:"urn:com.workday/bsvc Value_Date,omitempty"`
	ReversalIndicator                     *bool                                      `xml:"urn:com.workday/bsvc Reversal_Indicator,omitempty"`
	AdditionalReferenceField1             string                                     `xml:"urn:com.workday/bsvc Additional_Reference_Field_1,omitempty"`
	AdditionalReferenceField2             string                                     `xml:"urn:com.workday/bsvc Additional_Reference_Field_2,omitempty"`
	AdditionalReferenceField3             string                                     `xml:"urn:com.workday/bsvc Additional_Reference_Field_3,omitempty"`
	AdditionalReferenceField4             string                                     `xml:"urn:com.workday/bsvc Additional_Reference_Field_4,omitempty"`
	AdditionalReferenceField5             string                                     `xml:"urn:com.workday/bsvc Additional_Reference_Field_5,omitempty"`
}

func (t *BankStatementTransactionHVDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankStatementTransactionHVDataType
	var layout struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
		ValueDate       *xsdDate `xml:"urn:com.workday/bsvc Value_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TransactionDate = (*xsdDate)(layout.T.TransactionDate)
	layout.ValueDate = (*xsdDate)(layout.T.ValueDate)
	return e.EncodeElement(layout, start)
}
func (t *BankStatementTransactionHVDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankStatementTransactionHVDataType
	var overlay struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
		ValueDate       *xsdDate `xml:"urn:com.workday/bsvc Value_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionDate = (*xsdDate)(overlay.T.TransactionDate)
	overlay.ValueDate = (*xsdDate)(overlay.T.ValueDate)
	return d.DecodeElement(&overlay, &start)
}

// Element containing Bank Statement Transaction Type Data
type BankStatementTransactionTypeDataType struct {
	BankStatementTransactionTypeReference *BankStatementTransactionTypeObjectType `xml:"urn:com.workday/bsvc Bank_Statement_Transaction_Type_Reference,omitempty"`
	ID                                    string                                  `xml:"urn:com.workday/bsvc ID,omitempty"`
	TypeCode                              string                                  `xml:"urn:com.workday/bsvc Type_Code"`
	DebitCredit                           string                                  `xml:"urn:com.workday/bsvc Debit_Credit,omitempty"`
	Description                           string                                  `xml:"urn:com.workday/bsvc Description"`
	BankStatementTypeCategoryReference    *BankStatementTypeCategoryObjectType    `xml:"urn:com.workday/bsvc Bank_Statement_Type_Category_Reference"`
}

// Contains a unique identifier for an instance of an object.
type BankStatementTransactionTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankStatementTransactionTypeObjectType struct {
	ID         []BankStatementTransactionTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Bank Statement Transaction Type Reference.
type BankStatementTransactionTypeReferenceType struct {
	TypeCode string `xml:"urn:com.workday/bsvc Type_Code,omitempty"`
}

// Wrapper element that includes Bank Statement instances and data.
type BankStatementType struct {
	BankStatementReference *BankStatementObjectType `xml:"urn:com.workday/bsvc Bank_Statement_Reference,omitempty"`
	BankStatementData      []BankStatementDataType  `xml:"urn:com.workday/bsvc Bank_Statement_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankStatementTypeCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankStatementTypeCategoryObjectType struct {
	ID         []BankStatementTypeCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankStatementWorktagRuleSetObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankStatementWorktagRuleSetObjectType struct {
	ID         []BankStatementWorktagRuleSetObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This element is the wrapper around a list of elements representing the Bank Statements Request specific criteria needed to search for instances.  All of the elements are "AND" conditions. There are currently no elements.
type BankStatementsRequestCriteriaType struct {
	StatementDateFrom *time.Time `xml:"urn:com.workday/bsvc Statement_Date_From,omitempty"`
	StatementDateTo   *time.Time `xml:"urn:com.workday/bsvc Statement_Date_To,omitempty"`
}

func (t *BankStatementsRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BankStatementsRequestCriteriaType
	var layout struct {
		*T
		StatementDateFrom *xsdDate `xml:"urn:com.workday/bsvc Statement_Date_From,omitempty"`
		StatementDateTo   *xsdDate `xml:"urn:com.workday/bsvc Statement_Date_To,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StatementDateFrom = (*xsdDate)(layout.T.StatementDateFrom)
	layout.StatementDateTo = (*xsdDate)(layout.T.StatementDateTo)
	return e.EncodeElement(layout, start)
}
func (t *BankStatementsRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BankStatementsRequestCriteriaType
	var overlay struct {
		*T
		StatementDateFrom *xsdDate `xml:"urn:com.workday/bsvc Statement_Date_From,omitempty"`
		StatementDateTo   *xsdDate `xml:"urn:com.workday/bsvc Statement_Date_To,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StatementDateFrom = (*xsdDate)(overlay.T.StatementDateFrom)
	overlay.StatementDateTo = (*xsdDate)(overlay.T.StatementDateTo)
	return d.DecodeElement(&overlay, &start)
}

// Bank Statements Reference element contains the specific instance set containing the requested Bank Statements.
type BankStatementsRequestReferencesType struct {
	BankStatementsRequestReference []BankStatementObjectType `xml:"urn:com.workday/bsvc Bank_Statements_Request_Reference"`
}

// Wrapper element around a list of elements representing the amount of data that should be included in the Bank Statements response. If this element is not included then ALL of the Response Groups are included in the response.
type BankStatementsRequestResponseGroupType struct {
	IncludeReference         *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeBankStatementData *bool `xml:"urn:com.workday/bsvc Include_Bank_Statement_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BankingReconciliationMatchingRuleSetObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BankingReconciliationMatchingRuleSetObjectType struct {
	ID         []BankingReconciliationMatchingRuleSetObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BillableEntityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BillableEntityObjectType struct {
	ID         []BillableEntityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BusinessDocumentLineSplitObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BusinessDocumentLineSplitObjectType struct {
	ID         []BusinessDocumentLineSplitObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Business Entity Alternate Name Data
type BusinessEntityAlternateNameDataType struct {
	AlternateName               string                         `xml:"urn:com.workday/bsvc Alternate_Name"`
	AlternateNameUsageReference []AlternateNameUsageObjectType `xml:"urn:com.workday/bsvc Alternate_Name_Usage_Reference"`
}

// Element containing Business Entity Contact reference for update and all Business Entity Contact data items
type BusinessEntityContactDataType struct {
	PrimaryBillToContact              *bool                                  `xml:"urn:com.workday/bsvc Primary_Bill_To_Contact,omitempty"`
	DefaultBillToContact              *bool                                  `xml:"urn:com.workday/bsvc Default_Bill_To_Contact,omitempty"`
	BusinessEntityContactID           string                                 `xml:"urn:com.workday/bsvc Business_Entity_Contact_ID,omitempty"`
	SupplierReference                 *SupplierObjectType                    `xml:"urn:com.workday/bsvc Supplier_Reference"`
	BillableEntityReference           *BillableEntityObjectType              `xml:"urn:com.workday/bsvc Billable_Entity_Reference"`
	FinancialInstitutionReference     *FinancialInstitutionObjectType        `xml:"urn:com.workday/bsvc Financial_Institution_Reference"`
	TaxAuthorityReference             *TaxAuthorityObjectType                `xml:"urn:com.workday/bsvc Tax_Authority_Reference"`
	BusinessEntityContactPersonalData *BusinessEntityContactPersonalDataType `xml:"urn:com.workday/bsvc Business_Entity_Contact_Personal_Data"`
	ExternalIDData                    *ExternalIntegrationIDDataType         `xml:"urn:com.workday/bsvc External_ID_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type BusinessEntityContactObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BusinessEntityContactObjectType struct {
	ID         []BusinessEntityContactObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Business Entity Contact Data
type BusinessEntityContactPersonalDataType struct {
	NameData    []LegalNameDataType         `xml:"urn:com.workday/bsvc Name_Data"`
	ContactData *ContactInformationDataType `xml:"urn:com.workday/bsvc Contact_Data,omitempty"`
}

// Element containing Business Entity Contact request criteria
type BusinessEntityContactRequestCriteriaType struct {
}

// Element containing reference instances for a Business Entity Contact
type BusinessEntityContactRequestReferencesType struct {
	BusinessEntityContactReference []BusinessEntityContactObjectType `xml:"urn:com.workday/bsvc Business_Entity_Contact_Reference"`
}

// Wrapper Element that includes Business Entity Contact Instance and Data
type BusinessEntityContactResponseDataType struct {
	BusinessEntityContact []BusinessEntityContactType `xml:"urn:com.workday/bsvc Business_Entity_Contact,omitempty"`
}

// Element containing Business Entity Contact response grouping options
type BusinessEntityContactResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing Business Entity Contact reference and data
type BusinessEntityContactType struct {
	BusinessEntityContactReference *BusinessEntityContactObjectType `xml:"urn:com.workday/bsvc Business_Entity_Contact_Reference,omitempty"`
	BusinessEntityContactData      []BusinessEntityContactDataType  `xml:"urn:com.workday/bsvc Business_Entity_Contact_Data,omitempty"`
}

// Business Entity Logo Image Data is an element containing the logo data.
type BusinessEntityLogoImageDataType struct {
	Filename string  `xml:"urn:com.workday/bsvc Filename"`
	Image    *[]byte `xml:"urn:com.workday/bsvc Image,omitempty"`
}

func (t *BusinessEntityLogoImageDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BusinessEntityLogoImageDataType
	var layout struct {
		*T
		Image *xsdBase64Binary `xml:"urn:com.workday/bsvc Image,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Image = (*xsdBase64Binary)(layout.T.Image)
	return e.EncodeElement(layout, start)
}
func (t *BusinessEntityLogoImageDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BusinessEntityLogoImageDataType
	var overlay struct {
		*T
		Image *xsdBase64Binary `xml:"urn:com.workday/bsvc Image,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Image = (*xsdBase64Binary)(overlay.T.Image)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type BusinessEntityStatusValueObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type BusinessEntityStatusValueObjectType struct {
	ID         []BusinessEntityStatusValueObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Business Entity Data Element which is a wrapper for all business entity data of name and contract information
type BusinessEntityWWSDataType struct {
	BusinessEntityName          string                           `xml:"urn:com.workday/bsvc Business_Entity_Name"`
	BusinessEntityPhoneticName  string                           `xml:"urn:com.workday/bsvc Business_Entity_Phonetic_Name,omitempty"`
	BusinessEntityTaxID         string                           `xml:"urn:com.workday/bsvc Business_Entity_Tax_ID,omitempty"`
	ExternalEntityID            string                           `xml:"urn:com.workday/bsvc External_Entity_ID,omitempty"`
	ContactData                 *ContactInformationDataType      `xml:"urn:com.workday/bsvc Contact_Data,omitempty"`
	BusinessEntityLogoImageData *BusinessEntityLogoImageDataType `xml:"urn:com.workday/bsvc Business_Entity_Logo_Image_Data,omitempty"`
}

// Element for the attachments pertaining to a Event entered through a web service.
type BusinessProcessAttachmentDataType struct {
	FileName                         string                             `xml:"urn:com.workday/bsvc File_Name"`
	EventAttachmentDescription       string                             `xml:"urn:com.workday/bsvc Event_Attachment_Description,omitempty"`
	EventAttachmentCategoryReference *EventAttachmentCategoryObjectType `xml:"urn:com.workday/bsvc Event_Attachment_Category_Reference,omitempty"`
	File                             *[]byte                            `xml:"urn:com.workday/bsvc File,omitempty"`
	ContentType                      string                             `xml:"urn:com.workday/bsvc Content_Type,omitempty"`
}

func (t *BusinessProcessAttachmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BusinessProcessAttachmentDataType
	var layout struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	layout.T = (*T)(t)
	layout.File = (*xsdBase64Binary)(layout.T.File)
	return e.EncodeElement(layout, start)
}
func (t *BusinessProcessAttachmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BusinessProcessAttachmentDataType
	var overlay struct {
		*T
		File *xsdBase64Binary `xml:"urn:com.workday/bsvc File,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.File = (*xsdBase64Binary)(overlay.T.File)
	return d.DecodeElement(&overlay, &start)
}

// Captures a comment for the Business Process.
type BusinessProcessCommentDataType struct {
	Comment         string            `xml:"urn:com.workday/bsvc Comment,omitempty"`
	WorkerReference *WorkerObjectType `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
}

// Container for the processing options for sub-business processes within a business process. If no options are submitted (or the options are submitted as 'false') then the sub-business process is simply initiated as if it where submitted on-line with approvals, reviews, notifications and to-do's in place. If the Initiator is an Integration System User, any validations you configured on the Initiation step are ignored.
type BusinessSubProcessParametersType struct {
	AutoComplete                  *bool                               `xml:"urn:com.workday/bsvc Auto_Complete,omitempty"`
	Skip                          *bool                               `xml:"urn:com.workday/bsvc Skip,omitempty"`
	BusinessProcessCommentData    *BusinessProcessCommentDataType     `xml:"urn:com.workday/bsvc Business_Process_Comment_Data,omitempty"`
	BusinessProcessAttachmentData []BusinessProcessAttachmentDataType `xml:"urn:com.workday/bsvc Business_Process_Attachment_Data,omitempty"`
}

// Cancel Ad hoc Bank Transaction Request
type CancelAdhocBankTransactionRequestType struct {
	AdhocBankTransactionReference *AdHocBankTransactionObjectType `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_Reference"`
	Version                       string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Cancel Ad hoc Bank Transaction Response
type CancelAdhocBankTransactionResponseType struct {
	AdhocBankTransactionReference *AdHocBankTransactionObjectType `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_Reference,omitempty"`
	Version                       string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Cancel Ad hoc Payment Request
type CancelAdhocPaymentRequestType struct {
	AdhocPaymentReference *AdHocPaymentObjectType `xml:"urn:com.workday/bsvc Ad_hoc_Payment_Reference"`
	VoidCheck             *bool                   `xml:"urn:com.workday/bsvc Void_Check,omitempty"`
	ReasonforCancel       string                  `xml:"urn:com.workday/bsvc Reason_for_Cancel,omitempty"`
	Version               string                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Cancel Ad hoc Payment Response
type CancelAdhocPaymentResponseType struct {
	AdhocPaymentReference *AdHocPaymentObjectType `xml:"urn:com.workday/bsvc Ad_hoc_Payment_Reference,omitempty"`
	Version               string                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This is a Donor Contribution Reference ID that can be used in the Web Service Request/Response to indicate the Existing Donor Contribution that needs to be cancelled.
type CancelDonorContributionRequestType struct {
	DonorContributionReference *DonorContributionObjectType `xml:"urn:com.workday/bsvc Donor_Contribution_Reference"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This is a Donor Contribution Reference ID that can be used in the Web Service Request/Response to indicate the Existing Donor Contribution that needs to be cancelled.
type CancelDonorContributionResponseType struct {
	DonorContributionReference *DonorContributionObjectType `xml:"urn:com.workday/bsvc Donor_Contribution_Reference,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Cancel Investment Pool Transaction Request
type CancelInvestmentPoolTransactionRequestType struct {
	InvestmentPoolTransactionReference *InvestmentPoolTransactionObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Transaction_Reference"`
	Version                            string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Cancel Investment Pool Transaction Response
type CancelInvestmentPoolTransactionResponseType struct {
	InvestmentPoolTransactionReference *InvestmentPoolTransactionObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Transaction_Reference,omitempty"`
	Version                            string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Cash Activity Category data.
type CashActivityCategoryDataType struct {
	CashActivityCategoryID string `xml:"urn:com.workday/bsvc Cash_Activity_Category_ID,omitempty"`
	CashActivityCategory   string `xml:"urn:com.workday/bsvc Cash_Activity_Category"`
	Description            string `xml:"urn:com.workday/bsvc Description,omitempty"`
	Inactive               string `xml:"urn:com.workday/bsvc Inactive,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CashActivityCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CashActivityCategoryObjectType struct {
	ID         []CashActivityCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// The search criteria for the request.
type CashActivityCategoryRequestCriteriaType struct {
}

// Contains the reference to the Cash Activity Category.
type CashActivityCategoryRequestReferencesType struct {
	CashActivityCategoryReference []CashActivityCategoryObjectType `xml:"urn:com.workday/bsvc Cash_Activity_Category_Reference"`
}

// Wrapper containing the response data.
type CashActivityCategoryResponseDataType struct {
	CashActivityCategory []CashActivityCategoryType `xml:"urn:com.workday/bsvc Cash_Activity_Category,omitempty"`
}

// Contains a flag that determines whether the reference is included in the response.
type CashActivityCategoryResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Response data containing the reference and data items.
type CashActivityCategoryType struct {
	CashActivityCategoryReference *CashActivityCategoryObjectType `xml:"urn:com.workday/bsvc Cash_Activity_Category_Reference,omitempty"`
	CashActivityCategoryData      []CashActivityCategoryDataType  `xml:"urn:com.workday/bsvc Cash_Activity_Category_Data,omitempty"`
}

// Cash Out Address Data
type CashOutAddressDataType struct {
	CashOutAddressBusinessName    string                       `xml:"urn:com.workday/bsvc Cash_Out_Address_Business_Name,omitempty"`
	FormattedCashOutAddress       string                       `xml:"urn:com.workday/bsvc Formatted_Cash_Out_Address,omitempty"`
	CashOutAddressInformationData []AddressInformationDataType `xml:"urn:com.workday/bsvc Cash_Out_Address_Information_Data,omitempty"`
}

// Cash Pool data.
type CashPoolDataType struct {
	CashPoolID                 string                      `xml:"urn:com.workday/bsvc Cash_Pool_ID,omitempty"`
	Name                       string                      `xml:"urn:com.workday/bsvc Name"`
	Description                string                      `xml:"urn:com.workday/bsvc Description,omitempty"`
	Inactive                   *bool                       `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	MasterBankAccountReference *FinancialAccountObjectType `xml:"urn:com.workday/bsvc Master_Bank_Account_Reference"`
	SubBankAccountsData        []SubBankAccountsDataType   `xml:"urn:com.workday/bsvc Sub_Bank_Accounts_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CashPoolObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CashPoolObjectType struct {
	ID         []CashPoolObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the reference to the Cash Pool.
type CashPoolRequestReferencesType struct {
	CashPoolReference []CashPoolObjectType `xml:"urn:com.workday/bsvc Cash_Pool_Reference"`
}

// Wrapper containing the response data.
type CashPoolResponseDataType struct {
	CashPool []CashPoolType `xml:"urn:com.workday/bsvc Cash_Pool,omitempty"`
}

// Response data containing the reference and data items.
type CashPoolType struct {
	CashPoolReference *CashPoolObjectType `xml:"urn:com.workday/bsvc Cash_Pool_Reference,omitempty"`
	CashPoolData      []CashPoolDataType  `xml:"urn:com.workday/bsvc Cash_Pool_Data,omitempty"`
}

// Catch-up Payout Criteria Data
type CatchupPayoutCriteriaDataType struct {
	InvestmentPoolReference     *InvestmentPoolObjectType     `xml:"urn:com.workday/bsvc Investment_Pool_Reference"`
	TransactionDate             time.Time                     `xml:"urn:com.workday/bsvc Transaction_Date"`
	GiftReference               []GiftObjectType              `xml:"urn:com.workday/bsvc Gift_Reference"`
	FiscalTimeIntervalReference *FiscalTimeIntervalObjectType `xml:"urn:com.workday/bsvc Fiscal_Time_Interval_Reference"`
	FiscalYearReference         *FiscalYearObjectType         `xml:"urn:com.workday/bsvc Fiscal_Year_Reference"`
	Payoutperunit               float64                       `xml:"urn:com.workday/bsvc Payout_per_unit"`
	Administrativefeeperunit    float64                       `xml:"urn:com.workday/bsvc Administrative_fee_per_unit"`
}

func (t *CatchupPayoutCriteriaDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CatchupPayoutCriteriaDataType
	var layout struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date"`
	}
	layout.T = (*T)(t)
	layout.TransactionDate = (*xsdDate)(&layout.T.TransactionDate)
	return e.EncodeElement(layout, start)
}
func (t *CatchupPayoutCriteriaDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CatchupPayoutCriteriaDataType
	var overlay struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionDate = (*xsdDate)(&overlay.T.TransactionDate)
	return d.DecodeElement(&overlay, &start)
}

// Check Payments To Print Search Criteria
type CheckPaymentsToPrintCriteriaType struct {
	PaymentGroupReference                   *PaymentGroupObjectType          `xml:"urn:com.workday/bsvc Payment_Group_Reference,omitempty"`
	PaymentDateonDateOrAfter                *time.Time                       `xml:"urn:com.workday/bsvc Payment_Date_on_Date_Or_After,omitempty"`
	PaymentDateonDateOrBefore               *time.Time                       `xml:"urn:com.workday/bsvc Payment_Date_on_Date_Or_Before,omitempty"`
	PeriodsReference                        []PeriodObjectType               `xml:"urn:com.workday/bsvc Periods_Reference,omitempty"`
	PayRunGroupsandorPayGroupsReference     []PayRunGroupSelectionObjectType `xml:"urn:com.workday/bsvc Pay_Run_Groups_and_or_Pay_Groups_Reference,omitempty"`
	ExcludePaymentStatusesforCheckReference []PaymentStatusObjectType        `xml:"urn:com.workday/bsvc Exclude_Payment_Statuses_for_Check_Reference,omitempty"`
}

func (t *CheckPaymentsToPrintCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CheckPaymentsToPrintCriteriaType
	var layout struct {
		*T
		PaymentDateonDateOrAfter  *xsdDate `xml:"urn:com.workday/bsvc Payment_Date_on_Date_Or_After,omitempty"`
		PaymentDateonDateOrBefore *xsdDate `xml:"urn:com.workday/bsvc Payment_Date_on_Date_Or_Before,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PaymentDateonDateOrAfter = (*xsdDate)(layout.T.PaymentDateonDateOrAfter)
	layout.PaymentDateonDateOrBefore = (*xsdDate)(layout.T.PaymentDateonDateOrBefore)
	return e.EncodeElement(layout, start)
}
func (t *CheckPaymentsToPrintCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CheckPaymentsToPrintCriteriaType
	var overlay struct {
		*T
		PaymentDateonDateOrAfter  *xsdDate `xml:"urn:com.workday/bsvc Payment_Date_on_Date_Or_After,omitempty"`
		PaymentDateonDateOrBefore *xsdDate `xml:"urn:com.workday/bsvc Payment_Date_on_Date_Or_Before,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDateonDateOrAfter = (*xsdDate)(overlay.T.PaymentDateonDateOrAfter)
	overlay.PaymentDateonDateOrBefore = (*xsdDate)(overlay.T.PaymentDateonDateOrBefore)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type CheckPrintLayoutAbstractObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CheckPrintLayoutAbstractObjectType struct {
	ID         []CheckPrintLayoutAbstractObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Sets optional configuration for sorting checks.
type CheckSortingSetupDataType struct {
	CheckSortingSetupReference *UniqueIdentifierObjectType     `xml:"urn:com.workday/bsvc Check_Sorting_Setup_Reference,omitempty"`
	CheckSortingSetupLineData  []CheckSortingSetupLineDataType `xml:"urn:com.workday/bsvc Check_Sorting_Setup_Line_Data,omitempty"`
}

// Contains reference to Check Sorting Setup Line and associated data.
type CheckSortingSetupLineDataType struct {
	CheckSortingSetupLineReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Check_Sorting_Setup_Line_Reference,omitempty"`
	CheckSortingOrder              string                       `xml:"urn:com.workday/bsvc Check_Sorting_Order"`
	CheckSortingValueReference     *UniqueIdentifierObjectType  `xml:"urn:com.workday/bsvc Check_Sorting_Value_Reference,omitempty"`
	SplitBatch                     string                       `xml:"urn:com.workday/bsvc Split_Batch,omitempty"`
}

// Encapsulating element for all Communication Method Usage data.
type CommunicationMethodUsageDataType struct {
	TypeReference   []CommunicationUsageTypeReferenceType `xml:"urn:com.workday/bsvc Type_Reference"`
	UseForReference []string                              `xml:"urn:com.workday/bsvc Use_For_Reference,omitempty"`
	Comments        string                                `xml:"urn:com.workday/bsvc Comments,omitempty"`
	Public          bool                                  `xml:"urn:com.workday/bsvc Public,attr,omitempty"`
}

// Encapsulating element for all Communication Method Usage data.
type CommunicationMethodUsageInformationDataType struct {
	TypeData                []CommunicationUsageTypeDataType               `xml:"urn:com.workday/bsvc Type_Data"`
	UseForReference         []CommunicationUsageBehaviorObjectType         `xml:"urn:com.workday/bsvc Use_For_Reference,omitempty"`
	UseForTenantedReference []CommunicationUsageBehaviorTenantedObjectType `xml:"urn:com.workday/bsvc Use_For_Tenanted_Reference,omitempty"`
	Comments                string                                         `xml:"urn:com.workday/bsvc Comments,omitempty"`
	Public                  bool                                           `xml:"urn:com.workday/bsvc Public,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CommunicationUsageBehaviorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CommunicationUsageBehaviorObjectType struct {
	ID         []CommunicationUsageBehaviorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CommunicationUsageBehaviorTenantedObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CommunicationUsageBehaviorTenantedObjectType struct {
	ID         []CommunicationUsageBehaviorTenantedObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Reference ID for the communication usage type.
type CommunicationUsageTypeDataType struct {
	TypeReference *CommunicationUsageTypeObjectType `xml:"urn:com.workday/bsvc Type_Reference"`
	Primary       bool                              `xml:"urn:com.workday/bsvc Primary,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CommunicationUsageTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CommunicationUsageTypeObjectType struct {
	ID         []CommunicationUsageTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Reference element representing a unique instance of Communication Usage Type.
type CommunicationUsageTypeReferenceType struct {
	Value   string `xml:",chardata"`
	Primary bool   `xml:"urn:com.workday/bsvc Primary,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CompanyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CompanyObjectType struct {
	ID         []CompanyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// All of the person's contact data (address, phone, email, instant messenger, web address).
type ContactInformationDataType struct {
	AddressData          []AddressInformationDataType          `xml:"urn:com.workday/bsvc Address_Data,omitempty"`
	PhoneData            []PhoneInformationDataType            `xml:"urn:com.workday/bsvc Phone_Data,omitempty"`
	EmailAddressData     []EmailAddressInformationDataType     `xml:"urn:com.workday/bsvc Email_Address_Data,omitempty"`
	InstantMessengerData []InstantMessengerInformationDataType `xml:"urn:com.workday/bsvc Instant_Messenger_Data,omitempty"`
	WebAddressData       []WebAddressInformationDataType       `xml:"urn:com.workday/bsvc Web_Address_Data,omitempty"`
}

// Comment data for correction.
type CorrectAdHocPaymentCommentDataType struct {
	Comment string `xml:"urn:com.workday/bsvc Comment,omitempty"`
}

// Element containing an Ad hoc Payment data with corrections
type CorrectAdHocPaymentDataType struct {
	CompanyReference              *CompanyObjectType                                  `xml:"urn:com.workday/bsvc Company_Reference"`
	BankAccountReference          *FinancialAccountObjectType                         `xml:"urn:com.workday/bsvc Bank_Account_Reference"`
	PayeeReference                *PayeeObjectType                                    `xml:"urn:com.workday/bsvc Payee_Reference"`
	TaxAuthorityFormTypeReference *TaxAuthorityFormTypeObjectType                     `xml:"urn:com.workday/bsvc Tax_Authority_Form_Type_Reference,omitempty"`
	TINTypeReference              *TaxpayerIDNumberTypeObjectType                     `xml:"urn:com.workday/bsvc TIN_Type_Reference,omitempty"`
	TaxID                         string                                              `xml:"urn:com.workday/bsvc Tax_ID,omitempty"`
	TaxPayment                    *bool                                               `xml:"urn:com.workday/bsvc Tax_Payment,omitempty"`
	CurrencyReference             *CurrencyObjectType                                 `xml:"urn:com.workday/bsvc Currency_Reference"`
	CurrencyRateTypeReference     *CurrencyRateTypeObjectType                         `xml:"urn:com.workday/bsvc Currency_Rate_Type_Reference,omitempty"`
	CurrencyConversionRate        float64                                             `xml:"urn:com.workday/bsvc Currency_Conversion_Rate,omitempty"`
	EliminateFXGainLoss           *bool                                               `xml:"urn:com.workday/bsvc Eliminate_FX_Gain_Loss,omitempty"`
	TaxOptionReference            *TaxOptionObjectType                                `xml:"urn:com.workday/bsvc Tax_Option_Reference,omitempty"`
	ShipToAddressReference        *UniqueIdentifierObjectType                         `xml:"urn:com.workday/bsvc Ship-To_Address_Reference,omitempty"`
	TaxCodeReference              *TaxCodeObjectType                                  `xml:"urn:com.workday/bsvc Tax_Code_Reference,omitempty"`
	PaymentDate                   time.Time                                           `xml:"urn:com.workday/bsvc Payment_Date"`
	PaymentTypeReference          *PaymentTypeObjectType                              `xml:"urn:com.workday/bsvc Payment_Type_Reference"`
	HandlingCodeReference         *PaymentHandlingInstructionObjectType               `xml:"urn:com.workday/bsvc Handling_Code_Reference,omitempty"`
	Memo                          string                                              `xml:"urn:com.workday/bsvc Memo,omitempty"`
	Addenda                       string                                              `xml:"urn:com.workday/bsvc Addenda,omitempty"`
	TaxAmount                     float64                                             `xml:"urn:com.workday/bsvc Tax_Amount,omitempty"`
	FreightAmount                 float64                                             `xml:"urn:com.workday/bsvc Freight_Amount,omitempty"`
	OtherCharges                  float64                                             `xml:"urn:com.workday/bsvc Other_Charges,omitempty"`
	ExternalReference             string                                              `xml:"urn:com.workday/bsvc External_Reference,omitempty"`
	CorrectAdHocPaymentLineData   []CorrectAdHocPaymentLineDataType                   `xml:"urn:com.workday/bsvc Correct_Ad_Hoc_Payment_Line_Data,omitempty"`
	TaxCodeData                   []CorrectAdhocPaymentTaxableCodeApplicationDataType `xml:"urn:com.workday/bsvc Tax_Code_Data,omitempty"`
	CommentData                   []CorrectAdHocPaymentCommentDataType                `xml:"urn:com.workday/bsvc Comment_Data,omitempty"`
}

func (t *CorrectAdHocPaymentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CorrectAdHocPaymentDataType
	var layout struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date"`
	}
	layout.T = (*T)(t)
	layout.PaymentDate = (*xsdDate)(&layout.T.PaymentDate)
	return e.EncodeElement(layout, start)
}
func (t *CorrectAdHocPaymentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CorrectAdHocPaymentDataType
	var overlay struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDate = (*xsdDate)(&overlay.T.PaymentDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains corrections to Ad Hoc Payment Line.
type CorrectAdHocPaymentLineDataType struct {
	SupplierInvoiceLineID          string                        `xml:"urn:com.workday/bsvc Supplier_Invoice_Line_ID"`
	LineOrder                      string                        `xml:"urn:com.workday/bsvc Line_Order,omitempty"`
	IntercompanyAffiliateReference *CompanyObjectType            `xml:"urn:com.workday/bsvc Intercompany_Affiliate_Reference"`
	PurchaseItemReference          *ProcurementItemObjectType    `xml:"urn:com.workday/bsvc Purchase_Item_Reference,omitempty"`
	ItemDescription                string                        `xml:"urn:com.workday/bsvc Item_Description,omitempty"`
	SpendCategoryReference         *SpendCategoryObjectType      `xml:"urn:com.workday/bsvc Spend_Category_Reference,omitempty"`
	TaxApplicabilityReference      *TaxApplicabilityObjectType   `xml:"urn:com.workday/bsvc Tax_Applicability_Reference,omitempty"`
	TaxCodeReference               *TaxCodeObjectType            `xml:"urn:com.workday/bsvc Tax_Code_Reference,omitempty"`
	Quantity                       float64                       `xml:"urn:com.workday/bsvc Quantity"`
	UnitCost                       float64                       `xml:"urn:com.workday/bsvc Unit_Cost"`
	ExtendedAmount                 float64                       `xml:"urn:com.workday/bsvc Extended_Amount"`
	BusinessDocumentLineMemo       string                        `xml:"urn:com.workday/bsvc Business_Document_Line_Memo,omitempty"`
	WorktagsReference              []AccountingWorktagObjectType `xml:"urn:com.workday/bsvc Worktags_Reference,omitempty"`
}

// Correct Ad Hoc Payment Web Service Request
type CorrectAdHocPaymentRequestType struct {
	AdHocPaymentReference *AdHocPaymentObjectType      `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_Reference"`
	AdHocPaymentData      *CorrectAdHocPaymentDataType `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_Data,omitempty"`
	Version               string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Correct Ad Hoc Payment Response Data
type CorrectAdHocPaymentResponseType struct {
	AdhocPaymentReference *AdHocPaymentObjectType `xml:"urn:com.workday/bsvc Ad_hoc_Payment_Reference,omitempty"`
	Version               string                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element contain tax details data.
type CorrectAdhocPaymentTaxableCodeApplicationDataType struct {
	TaxApplicabilityReference *TaxApplicabilityObjectType `xml:"urn:com.workday/bsvc Tax_Applicability_Reference"`
	TaxCodeReference          *TaxCodeObjectType          `xml:"urn:com.workday/bsvc Tax_Code_Reference"`
	TaxbyTaxCode              float64                     `xml:"urn:com.workday/bsvc Tax_by_Tax_Code,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CostCenterObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CostCenterObjectType struct {
	ID         []CostCenterObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CounterpartyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CounterpartyObjectType struct {
	ID         []CounterpartyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CountryCityObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type CountryCityObjectType struct {
	ID         []CountryCityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CountryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CountryObjectType struct {
	ID         []CountryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CountryPredefinedPersonNameComponentValueObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CountryPredefinedPersonNameComponentValueObjectType struct {
	ID         []CountryPredefinedPersonNameComponentValueObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CountryRegionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CountryRegionObjectType struct {
	ID         []CountryRegionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CurrencyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CurrencyObjectType struct {
	ID         []CurrencyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CurrencyRateTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CurrencyRateTypeObjectType struct {
	ID         []CurrencyRateTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomerInvoiceLineObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomerInvoiceLineObjectType struct {
	ID         []CustomerInvoiceLineObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type CustomerObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type CustomerObjectType struct {
	ID         []CustomerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This is a name uniquely identifying a deduction already established in the Workday Payroll system.
type DeductionReferenceType struct {
	Code string `xml:"urn:com.workday/bsvc Code"`
}

// Result of the evaluation of an External Field based on a contextual instance.
type DocumentFieldResultDataType struct {
	FieldReference *IntegrationDocumentFieldObjectType `xml:"urn:com.workday/bsvc Field_Reference,omitempty"`
	Value          string                              `xml:"urn:com.workday/bsvc Value,omitempty"`
}

// Document Remittance Data
type DocumentRemittanceDataType struct {
	DocumentType                   string                                    `xml:"urn:com.workday/bsvc Document_Type,omitempty"`
	DocumentID                     string                                    `xml:"urn:com.workday/bsvc Document_ID,omitempty"`
	DocumentReference              string                                    `xml:"urn:com.workday/bsvc Document_Reference,omitempty"`
	AdditionalDocumentReference    string                                    `xml:"urn:com.workday/bsvc Additional_Document_Reference,omitempty"`
	DocumentDate                   *time.Time                                `xml:"urn:com.workday/bsvc Document_Date,omitempty"`
	TotalPayableAmount             float64                                   `xml:"urn:com.workday/bsvc Total_Payable_Amount,omitempty"`
	AmountPaid                     float64                                   `xml:"urn:com.workday/bsvc Amount_Paid,omitempty"`
	AmountDue                      float64                                   `xml:"urn:com.workday/bsvc Amount_Due,omitempty"`
	DiscountTaken                  float64                                   `xml:"urn:com.workday/bsvc Discount_Taken,omitempty"`
	TaxAmount                      float64                                   `xml:"urn:com.workday/bsvc Tax_Amount,omitempty"`
	WorkerReference                *WorkerObjectType                         `xml:"urn:com.workday/bsvc Worker_Reference,omitempty"`
	CreditCardTransactionReference []AbstractCreditCardTransactionObjectType `xml:"urn:com.workday/bsvc Credit_Card_Transaction_Reference,omitempty"`
	AdditionalInfo                 string                                    `xml:"urn:com.workday/bsvc Additional_Info,omitempty"`
	DocumentMemo                   string                                    `xml:"urn:com.workday/bsvc Document_Memo,omitempty"`
	WithheldTaxAmount              float64                                   `xml:"urn:com.workday/bsvc Withheld_Tax_Amount,omitempty"`
}

func (t *DocumentRemittanceDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DocumentRemittanceDataType
	var layout struct {
		*T
		DocumentDate *xsdDate `xml:"urn:com.workday/bsvc Document_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DocumentDate = (*xsdDate)(layout.T.DocumentDate)
	return e.EncodeElement(layout, start)
}
func (t *DocumentRemittanceDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DocumentRemittanceDataType
	var overlay struct {
		*T
		DocumentDate *xsdDate `xml:"urn:com.workday/bsvc Document_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DocumentDate = (*xsdDate)(overlay.T.DocumentDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type DocumentStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DocumentStatusObjectType struct {
	ID         []DocumentStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing essential Donor Contribution data
type DonorContributionDataType struct {
	ID                string                         `xml:"urn:com.workday/bsvc ID,omitempty"`
	Submit            *bool                          `xml:"urn:com.workday/bsvc Submit,omitempty"`
	DonorReference    *DonorObjectType               `xml:"urn:com.workday/bsvc Donor_Reference"`
	GiftReference     *GiftObjectType                `xml:"urn:com.workday/bsvc Gift_Reference"`
	DateReceived      time.Time                      `xml:"urn:com.workday/bsvc Date_Received"`
	DateRecorded      *time.Time                     `xml:"urn:com.workday/bsvc Date_Recorded,omitempty"`
	Amount            float64                        `xml:"urn:com.workday/bsvc Amount"`
	CurrencyReference *CurrencyObjectType            `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	AttachmentData    []FinancialsAttachmentDataType `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
}

func (t *DonorContributionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DonorContributionDataType
	var layout struct {
		*T
		DateReceived *xsdDate `xml:"urn:com.workday/bsvc Date_Received"`
		DateRecorded *xsdDate `xml:"urn:com.workday/bsvc Date_Recorded,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateReceived = (*xsdDate)(&layout.T.DateReceived)
	layout.DateRecorded = (*xsdDate)(layout.T.DateRecorded)
	return e.EncodeElement(layout, start)
}
func (t *DonorContributionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DonorContributionDataType
	var overlay struct {
		*T
		DateReceived *xsdDate `xml:"urn:com.workday/bsvc Date_Received"`
		DateRecorded *xsdDate `xml:"urn:com.workday/bsvc Date_Recorded,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateReceived = (*xsdDate)(&overlay.T.DateReceived)
	overlay.DateRecorded = (*xsdDate)(overlay.T.DateRecorded)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type DonorContributionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DonorContributionObjectType struct {
	ID         []DonorContributionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing search criteria for Donor Contributions.
type DonorContributionRequestCriteriaType struct {
}

// Element containing references to existing Donor Contributions.
type DonorContributionRequestReferencesType struct {
	DonorContributionReference []DonorContributionObjectType `xml:"urn:com.workday/bsvc Donor_Contribution_Reference"`
}

// Element containing data for existing Donor Contributions.
type DonorContributionResponseDataType struct {
	DonorContribution []DonorContributionType `xml:"urn:com.workday/bsvc Donor_Contribution,omitempty"`
}

// Element containing reference to existing Donor Contribution and data.
type DonorContributionType struct {
	DonorContributionReference *DonorContributionObjectType `xml:"urn:com.workday/bsvc Donor_Contribution_Reference,omitempty"`
	DonorContributionData      []DonorContributionDataType  `xml:"urn:com.workday/bsvc Donor_Contribution_Data,omitempty"`
}

// Element containing all Donor data
type DonorDataType struct {
	DonorID                       string                          `xml:"urn:com.workday/bsvc Donor_ID,omitempty"`
	DonorName                     string                          `xml:"urn:com.workday/bsvc Donor_Name,omitempty"`
	CustomerReference             *CustomerObjectType             `xml:"urn:com.workday/bsvc Customer_Reference,omitempty"`
	FinancialInstitutionReference *FinancialInstitutionObjectType `xml:"urn:com.workday/bsvc Financial_Institution_Reference,omitempty"`
	InvestorReference             *InvestorObjectType             `xml:"urn:com.workday/bsvc Investor_Reference,omitempty"`
	SponsorReference              *SponsorObjectType              `xml:"urn:com.workday/bsvc Sponsor_Reference,omitempty"`
	TaxAuthorityReference         *TaxAuthorityObjectType         `xml:"urn:com.workday/bsvc Tax_Authority_Reference,omitempty"`
	DonorTypeReference            *DonorTypeObjectType            `xml:"urn:com.workday/bsvc Donor_Type_Reference,omitempty"`
	ContactData                   *ContactInformationDataType     `xml:"urn:com.workday/bsvc Contact_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DonorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DonorObjectType struct {
	ID         []DonorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element for Donor Request Criteria contains Donor Request ID
type DonorRequestCriteriaType struct {
	DonorRequestID string `xml:"urn:com.workday/bsvc Donor_Request_ID,omitempty"`
}

// Element containing reference instances for a Donor
type DonorRequestReferencesType struct {
	DonorReference []DonorObjectType `xml:"urn:com.workday/bsvc Donor_Reference"`
}

// Element that includes Donor Instance
type DonorResponseDataType struct {
	Donor []DonorType `xml:"urn:com.workday/bsvc Donor,omitempty"`
}

// Wrapper Element that includes Donor Instance and Data
type DonorType struct {
	DonorReference *DonorObjectType `xml:"urn:com.workday/bsvc Donor_Reference,omitempty"`
	DonorData      *DonorDataType   `xml:"urn:com.workday/bsvc Donor_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type DonorTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type DonorTypeObjectType struct {
	ID         []DonorTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// This is a name uniquely identifying an earning already established in the Workday Payroll system.
type EarningReferenceType struct {
	Code string `xml:"urn:com.workday/bsvc Code"`
}

// Email Address Data
type EmailAddressDataType struct {
	EmailAddress string                             `xml:"urn:com.workday/bsvc Email_Address,omitempty"`
	EmailComment string                             `xml:"urn:com.workday/bsvc Email_Comment,omitempty"`
	UsageData    []CommunicationMethodUsageDataType `xml:"urn:com.workday/bsvc Usage_Data,omitempty"`
}

// Email Address Information
type EmailAddressInformationDataType struct {
	EmailAddress    string                                        `xml:"urn:com.workday/bsvc Email_Address,omitempty"`
	EmailComment    string                                        `xml:"urn:com.workday/bsvc Email_Comment,omitempty"`
	UsageData       []CommunicationMethodUsageInformationDataType `xml:"urn:com.workday/bsvc Usage_Data,omitempty"`
	EmailReference  *EmailReferenceObjectType                     `xml:"urn:com.workday/bsvc Email_Reference,omitempty"`
	ID              string                                        `xml:"urn:com.workday/bsvc ID,omitempty"`
	Delete          bool                                          `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	DoNotReplaceAll bool                                          `xml:"urn:com.workday/bsvc Do_Not_Replace_All,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EmailReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmailReferenceObjectType struct {
	ID         []EmailReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EmployeeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EmployeeObjectType struct {
	ID         []EmployeeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type EventAttachmentCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type EventAttachmentCategoryObjectType struct {
	ID         []EventAttachmentCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Exception (Errors and Warning) associated with the transaction.
type ExceptionDataType struct {
	Classification string `xml:"urn:com.workday/bsvc Classification,omitempty"`
	Message        string `xml:"urn:com.workday/bsvc Message,omitempty"`
}

// External cash activity data element
type ExternalCashActivityDataType struct {
	CompanyReference          *CompanyObjectType            `xml:"urn:com.workday/bsvc Company_Reference"`
	FinancialAccountReference *FinancialAccountObjectType   `xml:"urn:com.workday/bsvc Financial_Account_Reference"`
	CashActivityDate          time.Time                     `xml:"urn:com.workday/bsvc Cash_Activity_Date"`
	CurrencyReference         *CurrencyObjectType           `xml:"urn:com.workday/bsvc Currency_Reference"`
	CashActivityAmount        float64                       `xml:"urn:com.workday/bsvc Cash_Activity_Amount"`
	DebitCredit               string                        `xml:"urn:com.workday/bsvc Debit_Credit"`
	CounterpartyReference     *CounterpartyObjectType       `xml:"urn:com.workday/bsvc Counterparty_Reference,omitempty"`
	Processed                 *bool                         `xml:"urn:com.workday/bsvc Processed,omitempty"`
	Description               string                        `xml:"urn:com.workday/bsvc Description,omitempty"`
	WorktagsReference         []AccountingWorktagObjectType `xml:"urn:com.workday/bsvc Worktags_Reference,omitempty"`
}

func (t *ExternalCashActivityDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ExternalCashActivityDataType
	var layout struct {
		*T
		CashActivityDate *xsdDate `xml:"urn:com.workday/bsvc Cash_Activity_Date"`
	}
	layout.T = (*T)(t)
	layout.CashActivityDate = (*xsdDate)(&layout.T.CashActivityDate)
	return e.EncodeElement(layout, start)
}
func (t *ExternalCashActivityDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ExternalCashActivityDataType
	var overlay struct {
		*T
		CashActivityDate *xsdDate `xml:"urn:com.workday/bsvc Cash_Activity_Date"`
	}
	overlay.T = (*T)(t)
	overlay.CashActivityDate = (*xsdDate)(&overlay.T.CashActivityDate)
	return d.DecodeElement(&overlay, &start)
}

// External Cash Activity (HV) Data Element
type ExternalCashActivityHVDataType struct {
	ExternalCashActivityID string                        `xml:"urn:com.workday/bsvc External_Cash_Activity_ID,omitempty"`
	CompanyReference       *CompanyObjectType            `xml:"urn:com.workday/bsvc Company_Reference"`
	BankAccountReference   *FinancialAccountObjectType   `xml:"urn:com.workday/bsvc Bank_Account_Reference"`
	CashActivityDate       time.Time                     `xml:"urn:com.workday/bsvc Cash_Activity_Date"`
	CurrencyReference      *CurrencyObjectType           `xml:"urn:com.workday/bsvc Currency_Reference"`
	CashActivityAmount     float64                       `xml:"urn:com.workday/bsvc Cash_Activity_Amount"`
	DebitCredit            string                        `xml:"urn:com.workday/bsvc Debit_Credit"`
	CounterpartyReference  *CounterpartyObjectType       `xml:"urn:com.workday/bsvc Counterparty_Reference,omitempty"`
	AdHocCounterpartyName  string                        `xml:"urn:com.workday/bsvc Ad_Hoc_Counterparty_Name,omitempty"`
	Processed              *bool                         `xml:"urn:com.workday/bsvc Processed,omitempty"`
	Description            string                        `xml:"urn:com.workday/bsvc Description,omitempty"`
	WorktagsReference      []AccountingWorktagObjectType `xml:"urn:com.workday/bsvc Worktags_Reference,omitempty"`
}

func (t *ExternalCashActivityHVDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ExternalCashActivityHVDataType
	var layout struct {
		*T
		CashActivityDate *xsdDate `xml:"urn:com.workday/bsvc Cash_Activity_Date"`
	}
	layout.T = (*T)(t)
	layout.CashActivityDate = (*xsdDate)(&layout.T.CashActivityDate)
	return e.EncodeElement(layout, start)
}
func (t *ExternalCashActivityHVDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ExternalCashActivityHVDataType
	var overlay struct {
		*T
		CashActivityDate *xsdDate `xml:"urn:com.workday/bsvc Cash_Activity_Date"`
	}
	overlay.T = (*T)(t)
	overlay.CashActivityDate = (*xsdDate)(&overlay.T.CashActivityDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type ExternalCashActivityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ExternalCashActivityObjectType struct {
	ID         []ExternalCashActivityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// External cash activity request criteria element
type ExternalCashActivityRequestCriteriaType struct {
	CompanyReference     *CompanyObjectType          `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	BankAccountReference *FinancialAccountObjectType `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	StartDate            *time.Time                  `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate              *time.Time                  `xml:"urn:com.workday/bsvc End_Date,omitempty"`
}

func (t *ExternalCashActivityRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ExternalCashActivityRequestCriteriaType
	var layout struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *ExternalCashActivityRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ExternalCashActivityRequestCriteriaType
	var overlay struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// External cash activity request references element
type ExternalCashActivityRequestReferencesType struct {
	ExternalCashActivityReference []ExternalCashActivityObjectType `xml:"urn:com.workday/bsvc External_Cash_Activity_Reference,omitempty"`
}

// External cash activity response data element
type ExternalCashActivityResponseDataType struct {
	ExternalCashActivity []ExternalCashActivityType `xml:"urn:com.workday/bsvc External_Cash_Activity,omitempty"`
}

// External cash activity response group element
type ExternalCashActivityResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// External cash activity element
type ExternalCashActivityType struct {
	ExternalCashActivityReference *ExternalCashActivityObjectType `xml:"urn:com.workday/bsvc External_Cash_Activity_Reference,omitempty"`
	ExternalCashActivityData      []ExternalCashActivityDataType  `xml:"urn:com.workday/bsvc External_Cash_Activity_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ExternalCommitteeMemberObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ExternalCommitteeMemberObjectType struct {
	ID         []ExternalCommitteeMemberObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ExternalFieldandParameterInitializationProviderObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ExternalFieldandParameterInitializationProviderObjectType struct {
	ID         []ExternalFieldandParameterInitializationProviderObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Integration ID Help Text
type ExternalIntegrationIDDataType struct {
	ID []IDType `xml:"urn:com.workday/bsvc ID"`
}

// Field Result Request Criteria
type FieldAndParameterCriteriaDataType struct {
	ProviderReference []ExternalFieldandParameterInitializationProviderObjectType `xml:"urn:com.workday/bsvc Provider_Reference"`
}

// Contains a unique identifier for an instance of an object.
type FinancialAccountObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type FinancialAccountObjectType struct {
	ID         []FinancialAccountObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type FinancialInstitutionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type FinancialInstitutionObjectType struct {
	ID         []FinancialInstitutionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing Financial Institution request criteria
type FinancialInstitutionRequestCriteriaType struct {
}

// Element containing reference instances for an Financial Institution
type FinancialInstitutionRequestReferencesType struct {
	FinancialInstitutionReference []FinancialInstitutionObjectType `xml:"urn:com.workday/bsvc Financial_Institution_Reference"`
}

// Wrapper Element that includes Financial Institution Instance and Data
type FinancialInstitutionResponseDataType struct {
	FinancialInstitution []FinancialInstitutionType `xml:"urn:com.workday/bsvc Financial_Institution,omitempty"`
}

// Element containing Financial Institution response grouping options
type FinancialInstitutionResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing Financial Institution reference and data
type FinancialInstitutionType struct {
	FinancialInstitutionReference *FinancialInstitutionObjectType   `xml:"urn:com.workday/bsvc Financial_Institution_Reference,omitempty"`
	FinancialInstitutionData      []FinancialInstitutionWWSDataType `xml:"urn:com.workday/bsvc Financial_Institution_Data,omitempty"`
}

// Contains data for a Financial Institution
type FinancialInstitutionWWSDataType struct {
	FinancialInstitutionID                              string                         `xml:"urn:com.workday/bsvc Financial_Institution_ID,omitempty"`
	FinancialInstitutionReferenceID                     string                         `xml:"urn:com.workday/bsvc Financial_Institution_Reference_ID,omitempty"`
	FinancialInstitutionName                            string                         `xml:"urn:com.workday/bsvc Financial_Institution_Name,omitempty"`
	CreateFinancialInstitutionfromCustomerReference     *BillableEntityObjectType      `xml:"urn:com.workday/bsvc Create_Financial_Institution_from_Customer_Reference,omitempty"`
	CreateFinancialInstitutionfromTaxAuthorityReference *TaxAuthorityObjectType        `xml:"urn:com.workday/bsvc Create_Financial_Institution_from_Tax_Authority_Reference,omitempty"`
	CreateFinancialInstitutionfromSupplierReference     *SupplierObjectType            `xml:"urn:com.workday/bsvc Create_Financial_Institution_from_Supplier_Reference,omitempty"`
	CreateFinancialInstitutionfromInvestorReference     *InvestorObjectType            `xml:"urn:com.workday/bsvc Create_Financial_Institution_from_Investor_Reference,omitempty"`
	BankIdentifierCode                                  string                         `xml:"urn:com.workday/bsvc Bank_Identifier_Code,omitempty"`
	BankCode                                            string                         `xml:"urn:com.workday/bsvc Bank_Code,omitempty"`
	BusinessEntityData                                  *BusinessEntityWWSDataType     `xml:"urn:com.workday/bsvc Business_Entity_Data"`
	AttachmentData                                      []FinancialsAttachmentDataType `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type FinancialPartyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type FinancialPartyObjectType struct {
	ID         []FinancialPartyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element containing all Business Document Attachment data.
type FinancialsAttachmentDataType struct {
	FileContent *[]byte `xml:"urn:com.workday/bsvc File_Content,omitempty"`
	Comment     string  `xml:"urn:com.workday/bsvc Comment,omitempty"`
	ContentType string  `xml:"urn:com.workday/bsvc Content_Type,attr,omitempty"`
	Filename    string  `xml:"urn:com.workday/bsvc Filename,attr,omitempty"`
	Encoding    string  `xml:"urn:com.workday/bsvc Encoding,attr,omitempty"`
	Compressed  bool    `xml:"urn:com.workday/bsvc Compressed,attr,omitempty"`
}

func (t *FinancialsAttachmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T FinancialsAttachmentDataType
	var layout struct {
		*T
		FileContent *xsdBase64Binary `xml:"urn:com.workday/bsvc File_Content,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FileContent = (*xsdBase64Binary)(layout.T.FileContent)
	return e.EncodeElement(layout, start)
}
func (t *FinancialsAttachmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T FinancialsAttachmentDataType
	var overlay struct {
		*T
		FileContent *xsdBase64Binary `xml:"urn:com.workday/bsvc File_Content,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FileContent = (*xsdBase64Binary)(overlay.T.FileContent)
	return d.DecodeElement(&overlay, &start)
}

// Contains data for business processing
type FinancialsBusinessProcessParametersType struct {
	AutoComplete *bool                           `xml:"urn:com.workday/bsvc Auto_Complete,omitempty"`
	CommentData  *BusinessProcessCommentDataType `xml:"urn:com.workday/bsvc Comment_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type FiscalTimeIntervalObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type FiscalTimeIntervalObjectType struct {
	ID         []FiscalTimeIntervalObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type FiscalYearObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type FiscalYearObjectType struct {
	ID         []FiscalYearObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Funds Availability Data.
type FundsAvailabilityDataType struct {
	Amount                     float64    `xml:"urn:com.workday/bsvc Amount,omitempty"`
	AvailabilityDateTime       *time.Time `xml:"urn:com.workday/bsvc Availability_Date_Time,omitempty"`
	NumberofDaysuntilAvailable float64    `xml:"urn:com.workday/bsvc Number_of_Days_until_Available,omitempty"`
	UnknownNumberofDays        *bool      `xml:"urn:com.workday/bsvc Unknown_Number_of_Days,omitempty"`
}

func (t *FundsAvailabilityDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T FundsAvailabilityDataType
	var layout struct {
		*T
		AvailabilityDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Availability_Date_Time,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AvailabilityDateTime = (*xsdDateTime)(layout.T.AvailabilityDateTime)
	return e.EncodeElement(layout, start)
}
func (t *FundsAvailabilityDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T FundsAvailabilityDataType
	var overlay struct {
		*T
		AvailabilityDateTime *xsdDateTime `xml:"urn:com.workday/bsvc Availability_Date_Time,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AvailabilityDateTime = (*xsdDateTime)(overlay.T.AvailabilityDateTime)
	return d.DecodeElement(&overlay, &start)
}

// General Payment Search Criteria
type GeneralPaymentCriteriaType struct {
	CompanyReference                    []OrganizationObjectType              `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	BankAccountReference                []FinancialAccountObjectType          `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	PaymentCategoryReference            []PaymentCategoryObjectType           `xml:"urn:com.workday/bsvc Payment_Category_Reference,omitempty"`
	PayeeReference                      []PayeeObjectType                     `xml:"urn:com.workday/bsvc Payee_Reference,omitempty"`
	PayeeHierarchyReference             []HierarchyObjectType                 `xml:"urn:com.workday/bsvc Payee_Hierarchy_Reference,omitempty"`
	PaymentTypeReference                []PaymentTypeObjectType               `xml:"urn:com.workday/bsvc Payment_Type_Reference,omitempty"`
	PaymentStatusReference              []PaymentStatusObjectType             `xml:"urn:com.workday/bsvc Payment_Status_Reference,omitempty"`
	CustomerSupplierStatusReference     []BusinessEntityStatusValueObjectType `xml:"urn:com.workday/bsvc Customer_Supplier_Status_Reference,omitempty"`
	CurrencyReference                   []CurrencyObjectType                  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	PaymentDateonDateOrAfter            *time.Time                            `xml:"urn:com.workday/bsvc Payment_Date_on_Date_Or_After,omitempty"`
	PaymentDateonDateOrBefore           *time.Time                            `xml:"urn:com.workday/bsvc Payment_Date_on_Date_Or_Before,omitempty"`
	TransactionReferenceNumber          string                                `xml:"urn:com.workday/bsvc Transaction_Reference_Number,omitempty"`
	SettlementRunReference              []UniqueIdentifierObjectType          `xml:"urn:com.workday/bsvc Settlement_Run_Reference,omitempty"`
	SettlementRunNumber                 string                                `xml:"urn:com.workday/bsvc Settlement_Run_Number,omitempty"`
	SettlementRunName                   string                                `xml:"urn:com.workday/bsvc Settlement_Run_Name,omitempty"`
	CreatedbyWorkerReference            []WorkerObjectType                    `xml:"urn:com.workday/bsvc Created_by_Worker_Reference,omitempty"`
	PaymentAmountEqualTo                float64                               `xml:"urn:com.workday/bsvc Payment_Amount_Equal_To,omitempty"`
	PaymentAmountGreaterThan            float64                               `xml:"urn:com.workday/bsvc Payment_Amount_Greater_Than,omitempty"`
	PaymentAmountLessThan               float64                               `xml:"urn:com.workday/bsvc Payment_Amount_Less_Than,omitempty"`
	ExpensePayeeTypeReference           []UniqueIdentifierObjectType          `xml:"urn:com.workday/bsvc Expense_Payee_Type_Reference,omitempty"`
	TransactionIsIntercompany           *bool                                 `xml:"urn:com.workday/bsvc Transaction_Is_Intercompany,omitempty"`
	CompanyReceivingPaymentReference    []CompanyObjectType                   `xml:"urn:com.workday/bsvc Company_Receiving_Payment_Reference,omitempty"`
	PeriodsReference                    []PeriodObjectType                    `xml:"urn:com.workday/bsvc Periods_Reference,omitempty"`
	PayRunGroupsandorPayGroupsReference []PayRunGroupSelectionObjectType      `xml:"urn:com.workday/bsvc Pay_Run_Groups_and_or_Pay_Groups_Reference,omitempty"`
	EFTPaymentID                        string                                `xml:"urn:com.workday/bsvc EFT_Payment_ID,omitempty"`
	ReconciliationStatusReference       *ReconciliationStatusObjectType       `xml:"urn:com.workday/bsvc Reconciliation_Status_Reference,omitempty"`
}

func (t *GeneralPaymentCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T GeneralPaymentCriteriaType
	var layout struct {
		*T
		PaymentDateonDateOrAfter  *xsdDate `xml:"urn:com.workday/bsvc Payment_Date_on_Date_Or_After,omitempty"`
		PaymentDateonDateOrBefore *xsdDate `xml:"urn:com.workday/bsvc Payment_Date_on_Date_Or_Before,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PaymentDateonDateOrAfter = (*xsdDate)(layout.T.PaymentDateonDateOrAfter)
	layout.PaymentDateonDateOrBefore = (*xsdDate)(layout.T.PaymentDateonDateOrBefore)
	return e.EncodeElement(layout, start)
}
func (t *GeneralPaymentCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T GeneralPaymentCriteriaType
	var overlay struct {
		*T
		PaymentDateonDateOrAfter  *xsdDate `xml:"urn:com.workday/bsvc Payment_Date_on_Date_Or_After,omitempty"`
		PaymentDateonDateOrBefore *xsdDate `xml:"urn:com.workday/bsvc Payment_Date_on_Date_Or_Before,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDateonDateOrAfter = (*xsdDate)(overlay.T.PaymentDateonDateOrAfter)
	overlay.PaymentDateonDateOrBefore = (*xsdDate)(overlay.T.PaymentDateonDateOrBefore)
	return d.DecodeElement(&overlay, &start)
}

// This element is the top-level request element for all "Get" Ad Hoc Payee operations.
type GetAdHocPayeesRequestType struct {
	RequestReferences *AdHocPayeeRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *AdHocPayeeRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *AdHocPayeeResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top-level response element for all "Get" Ad Hoc Payees operations
type GetAdHocPayeesResponseType struct {
	RequestReferences *AdHocPayeeRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *AdHocPayeeRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *AdHocPayeeResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *AdHocPayeeResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request for Ad hoc Payment Template data
type GetAdHocPaymentTemplatesRequestType struct {
	RequestReferences *AdHocPaymentTemplateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *AdHocPaymentTemplateRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *AdHocPaymentTemplateResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response which consists of the input request and the data for the Ad Hoc Payment Templates
type GetAdHocPaymentTemplatesResponseType struct {
	RequestReferences *AdHocPaymentTemplateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *AdHocPaymentTemplateRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *AdHocPaymentTemplateResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                       `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *AdHocPaymentTemplateResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Root element for the Request on the Get operation. Contains references, criteria, filter and response group for specifying which instances to return in the Response.
type GetAdhocBankTransactionsRequestType struct {
	RequestReferences *AdhocBankTransactionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *AdhocBankTransactionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *AdhocBankTransactionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Ad hoc Bank Transactions response elements including received criteria, references, data based on criteria and grouping, results filter and result summary
type GetAdhocBankTransactionsResponseType struct {
	RequestReferences *AdhocBankTransactionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *AdhocBankTransactionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *AdhocBankTransactionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                       `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *AdhocBankTransactionResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Ad hoc Payment data
type GetAdhocPaymentsRequestType struct {
	RequestReferences *AdhocPaymentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *AdhocPaymentRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *AdhocPaymentResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Ad hoc Payments response elements including received criteria, references, data based on criteria and grouping, results filter and result summary
type GetAdhocPaymentsResponseType struct {
	RequestReferences *AdhocPaymentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *AdhocPaymentRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *AdhocPaymentResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType               `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *AdhocPaymentResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The request to retrieve Authority Types.
type GetAuthorityTypesRequestType struct {
	RequestReferences *AuthorityTypeRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *AuthorityTypeRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *AuthorityTypeResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response which contains the Authority Type data.
type GetAuthorityTypesResponseType struct {
	RequestReferences *AuthorityTypeRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *AuthorityTypeRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                 `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *AuthorityTypeResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *AuthorityTypeResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The request to retrieve Bank Account Signatories.
type GetBankAccountSignatoriesRequestType struct {
	RequestReferences *BankAccountRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response which contains the Bank Account Signatory data.
type GetBankAccountSignatoriesResponseType struct {
	RequestReferences *BankAccountRequestReferencesType     `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                  `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BankAccountSignatoryResponseDataType `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Bank Account Transfer Templates request element
type GetBankAccountTransferTemplatesRequestType struct {
	RequestReferences *BankAccountTransferTemplateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankAccountTransferTemplateRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankAccountTransferTemplateResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Bank Account Transfer Template response element
type GetBankAccountTransferTemplatesResponseType struct {
	RequestReferences []BankAccountTransferTemplateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   []BankAccountTransferTemplateRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    []ResponseFilterType                               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     []BankAccountTransferTemplateResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   []ResponseResultsType                              `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      []BankAccountTransferTemplateResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                             `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request for Bank Account Transfer for Settlement Template data
type GetBankAccountTransferforSettlementTemplatesRequestType struct {
	RequestReferences *BankAccountTransferforSettlementTemplateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankAccountTransferforSettlementTemplateRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankAccountTransferforSettlementTemplateResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response which consists of the input request and the data for the Bank Account Transfer for Settlement Templates
type GetBankAccountTransferforSettlementTemplatesResponseType struct {
	RequestReferences *BankAccountTransferforSettlementTemplateRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankAccountTransferforSettlementTemplateRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankAccountTransferforSettlementTemplateResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                                           `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BankAccountTransferforSettlementTemplateResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This element is the top-level request element for all Bank Accountl Transfer "Get" operations.
type GetBankAccountTransfersRequestType struct {
	RequestReferences *BankAccountTransferRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankAccountTransferRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankAccountTransferResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Bank Account Transfer response elements
type GetBankAccountTransfersResponseType struct {
	RequestReferences *BankAccountTransferRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankAccountTransferRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankAccountTransferResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                      `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BankAccountTransferResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top level request element for all Bank Account Transfers for Settlement "Get" operations
type GetBankAccountTransfersforSettlementRequestType struct {
	RequestReferences *BankAccountTransferforSettlementRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankAccountTransferforSettlementRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankAccountTransferforSettlementResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top level element containing Bank Account Transfer for Settlement response elements
type GetBankAccountTransfersforSettlementResponseType struct {
	RequestReferences *BankAccountTransferforSettlementRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankAccountTransferforSettlementRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankAccountTransferforSettlementResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BankAccountTransferforSettlementResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Bank Account data
type GetBankAccountsRequestType struct {
	RequestReferences *BankAccountRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankAccountRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankAccountResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Bank Account response elements including received criteria, references, data based on criteria and grouping, results filter and result summary
type GetBankAccountsResponseType struct {
	RequestReferences *BankAccountRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankAccountRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType               `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankAccountResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType              `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BankAccountResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Bank Branch data
type GetBankBranchesRequestType struct {
	RequestReferences *BankBranchRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankBranchRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankBranchResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Bank Branches response elements including received criteria, references, data based on criteria and grouping, results filter and result summary
type GetBankBranchesResponseType struct {
	RequestReferences *BankBranchRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankBranchRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankBranchResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BankBranchResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Bank Fee Service Codes Request
type GetBankFeeServiceCodesRequestType struct {
	RequestReferences *BankFeeServiceCodeRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankFeeServiceCodeRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankFeeServiceCodeResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Bank Fee Service Codes Response
type GetBankFeeServiceCodesResponseType struct {
	RequestReferences *BankFeeServiceCodeRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankFeeServiceCodeRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankFeeServiceCodeResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                     `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BankFeeServiceCodeResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Bank Fee Service Contracts Request
type GetBankFeeServiceContractsRequestType struct {
	RequestReferences *BankFeeServiceContractRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankFeeServiceContractRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Bank Fee Service Contract Response
type GetBankFeeServiceContractsResponseType struct {
	RequestReferences *BankFeeServiceContractRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankFeeServiceContractRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                         `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BankFeeServiceContractResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top level 'Get' Web Service to request Bank Fee Statement Files that have been posted.
type GetBankFeeStatementFilesRequestType struct {
	RequestReferences *BankFeeStatementFileRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankFeeStatementFileRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankFeeStatementFileResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top level 'Get' Web Service method that affords for queries against Bank Fee Statement File Data.
type GetBankFeeStatementFilesResponseType struct {
	RequestReferences *BankFeeStatementFileRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankFeeStatementFileRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankFeeStatementFileResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                       `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BankFeeStatementFileResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the request data for Bank Fee Statements.
type GetBankFeeStatementsRequestType struct {
	RequestReferences *BankFeeStatementRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankFeeStatementRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankFeeStatementResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing the Bank Fee Statement response data.
type GetBankFeeStatementsResponseType struct {
	RequestReferences *BankFeeStatementRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankFeeStatementRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankFeeStatementResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BankFeeStatementResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Bank Statement Custom Code Set data
type GetBankStatementCustomCodeSetsRequestType struct {
	RequestReferences *BankStatementCustomCodeSetRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankStatementCustomCodeSetRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankStatementCustomCodeSetResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Bank Statement Custom Code Set response elements including received criteria, references, data based on criteria and grouping, results filter and result summary
type GetBankStatementCustomCodeSetsResponseType struct {
	RequestReferences *BankStatementCustomCodeSetRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankStatementCustomCodeSetRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                              `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankStatementCustomCodeSetResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                             `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BankStatementCustomCodeSetResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top level 'Get' Web Service to request Bank Statement Files that have been posted.
type GetBankStatementFileRequestType struct {
	RequestReferences *BankStatementFileRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankStatementFileRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankStatementFileResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top level 'Get' Web Service method that affords for queries against Bank Statement File Data.
type GetBankStatementFilesResponseType struct {
	RequestReferences *BankStatementFileRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankStatementFileRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankStatementFileResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BankStatementFileResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top level 'Get' Web Service to request Bank Statements Files matching the filter criteria.
type GetBankStatementsRequestType struct {
	RequestReferences *BankStatementsRequestReferencesType    `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankStatementsRequestCriteriaType      `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankStatementsRequestResponseGroupType `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Bank Statements response elements including echoed request data and request result data.
type GetBankStatementsResponseType struct {
	RequestReferences *BankStatementsRequestReferencesType    `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BankStatementsRequestCriteriaType      `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BankStatementsRequestResponseGroupType `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BankStatementResponseDataType          `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Business Entity Contact data
type GetBusinessEntityContactsRequestType struct {
	RequestReferences *BusinessEntityContactRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BusinessEntityContactRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BusinessEntityContactResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Business Entity Contact response elements including received criteria, references, data based on criteria and grouping, results filter and result summary
type GetBusinessEntityContactsResponseType struct {
	RequestReferences *BusinessEntityContactRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *BusinessEntityContactRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *BusinessEntityContactResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                        `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *BusinessEntityContactResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The request to retrieve Cash Activity Categories.
type GetCashActivityCategoriesRequestType struct {
	RequestReferences *CashActivityCategoryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CashActivityCategoryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CashActivityCategoryResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response which contains the Cash Activity Category data.
type GetCashActivityCategoriesResponseType struct {
	RequestReferences *CashActivityCategoryRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *CashActivityCategoryRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *CashActivityCategoryResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                       `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *CashActivityCategoryResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The request to retrieve Cash Pools.
type GetCashPoolsRequestType struct {
	RequestReferences *CashPoolRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response which contains the Cash Pool data.
type GetCashPoolsResponseType struct {
	RequestReferences *CashPoolRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType           `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *CashPoolResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This element is the top-level request element for Get Donor Contribution..
type GetDonorContributionsRequestType struct {
	RequestReferences *DonorContributionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *DonorContributionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Donor Contribution response elements including received criteria, references, data based on criteria and grouping, results filter and result summary
type GetDonorContributionsResponseType struct {
	RequestReferences *DonorContributionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *DonorContributionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                     `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                    `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *DonorContributionResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                  `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request references and response filter.
type GetDonorsRequestType struct {
	RequestReferences *DonorRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *DonorRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Donor response elements including references, response filter, response results, and response data.
type GetDonorsResponseType struct {
	RequestReferences *DonorRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   []DonorRequestCriteriaType  `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType        `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *DonorResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get external cash activity request element
type GetExternalCashActivitiesRequestType struct {
	RequestReferences *ExternalCashActivityRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *ExternalCashActivityRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *ExternalCashActivityResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get external cash activity response element
type GetExternalCashActivitiesResponseType struct {
	RequestReferences *ExternalCashActivityRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *ExternalCashActivityRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *ExternalCashActivityResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                       `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *ExternalCashActivityResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Financial Institution data
type GetFinancialInstitutionsRequestType struct {
	RequestReferences *FinancialInstitutionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *FinancialInstitutionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *FinancialInstitutionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Financial Institution response elements including received criteria, references, data based on criteria and grouping, results filter and result summary
type GetFinancialInstitutionsResponseType struct {
	RequestReferences *FinancialInstitutionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *FinancialInstitutionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                        `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *FinancialInstitutionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                       `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *FinancialInstitutionResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains Request Element for Intraday Bank Statement
type GetIntradayBankStatementsRequestType struct {
	RequestReferences *IntradayBankStatementRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *IntradayBankStatementRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *IntradayBankStatementResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element Contains Intraday Bank Statement Response Data
type GetIntradayBankStatementsResponseType struct {
	RequestReferences *IntradayBankStatementRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *IntradayBankStatementRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *IntradayBankStatementResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                        `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *IntradayBankStatementResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This element is the top-level request element for all Investment Pool Adjustment "Get" operations.
type GetInvestmentPoolAdjustmentsRequestType struct {
	RequestReferences *InvestmentPoolAdjustmentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *InvestmentPoolAdjustmentRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *InvestmentPoolAdjustmentResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Investment Pool Adjustment response elements including echoed request data and request result
type GetInvestmentPoolAdjustmentsResponseType struct {
	RequestReferences *InvestmentPoolAdjustmentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *InvestmentPoolAdjustmentRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *InvestmentPoolAdjustmentResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                           `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *InvestmentPoolAdjustmentResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This element is the top-level request element for all Investment Pool Purchase "Get" operations.
type GetInvestmentPoolPurchasesRequestType struct {
	RequestReferences *InvestmentPoolPurchaseRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *InvestmentPoolPurchaseRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *InvestmentPoolPurchaseResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Investment Pool Purchase response elements including echoed request data and request result
type GetInvestmentPoolPurchasesResponseType struct {
	RequestReferences *InvestmentPoolPurchaseRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *InvestmentPoolPurchaseRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *InvestmentPoolPurchaseResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                         `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *InvestmentPoolPurchaseResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Investment Pool Sales Request
type GetInvestmentPoolSalesRequestType struct {
	RequestReferences *InvestmentPoolSaleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *InvestmentPoolSaleRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *InvestmentPoolSaleResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Investment Pool Sales Response
type GetInvestmentPoolSalesResponseType struct {
	RequestReferences *InvestmentPoolSaleRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *InvestmentPoolSaleRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                      `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *InvestmentPoolSaleResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                     `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *InvestmentPoolSaleResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This element is the top-level request element for all Investment Pool Transfer "Get" operations.
type GetInvestmentPoolTransfersRequestType struct {
	RequestReferences *InvestmentPoolTransferRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *InvestmentPoolTransferRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *InvestmentPoolTransferResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Investment Pool Transfer response elements including echoed request data and request result
type GetInvestmentPoolTransfersResponseType struct {
	RequestReferences *InvestmentPoolTransferRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *InvestmentPoolTransferRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *InvestmentPoolTransferResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                         `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *InvestmentPoolTransferResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This element is the top-level request element for Get Investment Statement.
type GetInvestmentStatementsRequestType struct {
	RequestReferences *InvestmentStatementRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *InvestmentStatementRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *InvestmentStatementResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Investment Statement response elements including received criteria, references, data based on criteria and grouping, results filter and result summary
type GetInvestmentStatementsResponseType struct {
	RequestReferences *InvestmentStatementRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *InvestmentStatementRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                       `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *InvestmentStatementResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                      `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *InvestmentStatementResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payment Election Enrollments Request
type GetPaymentElectionEnrollmentsRequestType struct {
	RequestReferences *PaymentElectionEnrollmentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PaymentElectionEnrollmentRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PaymentElectionEnrollmentResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payment Election Enrollments Response
type GetPaymentElectionEnrollmentsResponseType struct {
	RequestReferences *PaymentElectionEnrollmentRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PaymentElectionEnrollmentRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                             `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PaymentElectionEnrollmentResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                            `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PaymentElectionEnrollmentResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payment Election Options Request
type GetPaymentElectionOptionsRequestType struct {
	RequestReferences *PaymentElectionOptionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PaymentElectionOptionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PaymentElectionOptionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payment Election Options Response
type GetPaymentElectionOptionsResponseType struct {
	RequestReferences *PaymentElectionOptionRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PaymentElectionOptionRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                         `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PaymentElectionOptionResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                        `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PaymentElectionOptionResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payment Messages Request
type GetPaymentMessagesRequestType struct {
	RequestReferences *PaymentMessageRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PaymentMessageRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PaymentMessageResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Payment Messages Response Group
type GetPaymentMessagesResponseType struct {
	RequestReferences *PaymentMessageRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PaymentMessageRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                  `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PaymentMessageResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                 `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PaymentMessageResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                               `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payments Request
type GetPaymentsRequestType struct {
	RequestReferences *PaymentsRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PaymentsRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PaymentsResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Payment Response
type GetPaymentsResponseType struct {
	RequestReferences *PaymentsRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PaymentsRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType            `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PaymentsResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType           `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PaymentsResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains request reference, criteria, filter and response group for getting Petty Cash Account data
type GetPettyCashAccountsRequestType struct {
	RequestReferences *PettyCashAccountRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PettyCashAccountRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PettyCashAccountResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Petty Cash Account response elements including received criteria, references, data based on criteria and grouping, results filter and result summary
type GetPettyCashAccountsResponseType struct {
	RequestReferences *PettyCashAccountRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	RequestCriteria   *PettyCashAccountRequestCriteriaType   `xml:"urn:com.workday/bsvc Request_Criteria,omitempty"`
	ResponseFilter    *ResponseFilterType                    `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseGroup     *PettyCashAccountResponseGroupType     `xml:"urn:com.workday/bsvc Response_Group,omitempty"`
	ResponseResults   *ResponseResultsType                   `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *PettyCashAccountResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The request to retrieve Signature Methods.
type GetSignatureMethodsRequestType struct {
	RequestReferences *SignatureMethodRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response which contains the Signature Method data.
type GetSignatureMethodsResponseType struct {
	RequestReferences *SignatureMethodRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType                   `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType                  `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *SignatureMethodResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The request to retrieve Signers.
type GetSignersRequestType struct {
	RequestReferences *SignerRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	Version           string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response which contains the Signer data.
type GetSignersResponseType struct {
	RequestReferences *SignerRequestReferencesType `xml:"urn:com.workday/bsvc Request_References,omitempty"`
	ResponseFilter    *ResponseFilterType          `xml:"urn:com.workday/bsvc Response_Filter,omitempty"`
	ResponseResults   *ResponseResultsType         `xml:"urn:com.workday/bsvc Response_Results,omitempty"`
	ResponseData      *SignerResponseDataType      `xml:"urn:com.workday/bsvc Response_Data,omitempty"`
	Version           string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type GiftObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type GiftObjectType struct {
	ID         []GiftObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type HierarchyObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type HierarchyObjectType struct {
	ID         []HierarchyObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// External ID that uniquely identifies the integratable object within the context of the integration system identified by the System ID attribute.
type IDType struct {
	Value    string `xml:",chardata"`
	SystemID string `xml:"urn:com.workday/bsvc System_ID,attr,omitempty"`
}

// Contains data for creating or updating an ad hoc bank transaction and submitting for business processing
type ImportAdhocBankTransactionRequestType struct {
	AdhocBankTransactionReference *AdHocBankTransactionObjectType          `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_Reference,omitempty"`
	BusinessProcessParameters     *FinancialsBusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AdhocBankTransactionData      *AdhocBankTransactionHVDataType          `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_Data,omitempty"`
	AddOnly                       bool                                     `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                       string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Bank Fee Statement reference for update and all Bank Fee Statement data items.
type ImportBankFeeStatementRequestType struct {
	BankFeeStatementReference *BankFeeStatementObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Statement_Reference,omitempty"`
	BankFeeStatementData      *BankFeeStatementHVDataType `xml:"urn:com.workday/bsvc Bank_Fee_Statement_Data"`
	AddOnly                   bool                        `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                   string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Bank Statement reference for update and all Bank Statement data items.
type ImportBankStatementRequestType struct {
	BankStatementReference *BankStatementObjectType `xml:"urn:com.workday/bsvc Bank_Statement_Reference,omitempty"`
	BankStatementData      *BankStatementHVDataType `xml:"urn:com.workday/bsvc Bank_Statement_Data"`
	AddOnly                bool                     `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Import External Cash Activity Request Element
type ImportExternalCashActivityRequestType struct {
	LoadFileName             string                           `xml:"urn:com.workday/bsvc Load_File_Name,omitempty"`
	ExternalCashActivityData []ExternalCashActivityHVDataType `xml:"urn:com.workday/bsvc External_Cash_Activity_Data,omitempty"`
	AddOnly                  bool                             `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                  string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// This element contains all Bank Statement Reference and related Data for adding and updating Intraday Bank Statements.
type ImportIntradayBankStatementRequestType struct {
	IntradayBankStatementReference *IntradayBankStatementObjectType `xml:"urn:com.workday/bsvc Intraday_Bank_Statement_Reference,omitempty"`
	IntradayBankStatementData      *IntradayBankStatementDataType   `xml:"urn:com.workday/bsvc Intraday_Bank_Statement_Data"`
	AddOnly                        bool                             `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                        string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

type InstanceIDType struct {
	Value      string `xml:",chardata"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
	Type       string `xml:"urn:com.workday/bsvc type,attr,omitempty"`
}

type InstanceObjectType struct {
	ID         []InstanceIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Instant Messenger Address.
type InstantMessengerInformationDataType struct {
	InstantMessengerAddress       string                                        `xml:"urn:com.workday/bsvc Instant_Messenger_Address,omitempty"`
	InstantMessengerTypeReference *InstantMessengerTypeObjectType               `xml:"urn:com.workday/bsvc Instant_Messenger_Type_Reference,omitempty"`
	InstantMessengerComment       string                                        `xml:"urn:com.workday/bsvc Instant_Messenger_Comment,omitempty"`
	UsageData                     []CommunicationMethodUsageInformationDataType `xml:"urn:com.workday/bsvc Usage_Data,omitempty"`
	InstantMessengerReference     *InstantMessengerReferenceObjectType          `xml:"urn:com.workday/bsvc Instant_Messenger_Reference,omitempty"`
	ID                            string                                        `xml:"urn:com.workday/bsvc ID,omitempty"`
	Delete                        bool                                          `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	DoNotReplaceAll               bool                                          `xml:"urn:com.workday/bsvc Do_Not_Replace_All,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type InstantMessengerReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InstantMessengerReferenceObjectType struct {
	ID         []InstantMessengerReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type InstantMessengerTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InstantMessengerTypeObjectType struct {
	ID         []InstantMessengerTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationDocumentFieldObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type IntegrationDocumentFieldObjectType struct {
	ID         []IntegrationDocumentFieldObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type IntegrationSystemAuditedObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntegrationSystemAuditedObjectType struct {
	ID         []IntegrationSystemAuditedObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Intermediary Bank Account information including Order, Bank Account Type, Currency, Routing Transit Number, etc.
type IntermediaryBankAccountWWSDataType struct {
	IntermediaryBankAccountID   string                     `xml:"urn:com.workday/bsvc Intermediary_Bank_Account_ID,omitempty"`
	IntermediaryBankOrder       string                     `xml:"urn:com.workday/bsvc Intermediary_Bank_Order,omitempty"`
	CountryReference            *CountryObjectType         `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CurrencyReference           *CurrencyObjectType        `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	BankAccountNickname         string                     `xml:"urn:com.workday/bsvc Bank_Account_Nickname,omitempty"`
	BankAccountTypeReference    *BankAccountTypeObjectType `xml:"urn:com.workday/bsvc Bank_Account_Type_Reference,omitempty"`
	BankName                    string                     `xml:"urn:com.workday/bsvc Bank_Name,omitempty"`
	RoutingTransitNumber        string                     `xml:"urn:com.workday/bsvc Routing_Transit_Number,omitempty"`
	BranchID                    string                     `xml:"urn:com.workday/bsvc Branch_ID,omitempty"`
	BranchName                  string                     `xml:"urn:com.workday/bsvc Branch_Name,omitempty"`
	BankAccountNumber           string                     `xml:"urn:com.workday/bsvc Bank_Account_Number,omitempty"`
	CheckDigit                  string                     `xml:"urn:com.workday/bsvc Check_Digit,omitempty"`
	BankAccountName             string                     `xml:"urn:com.workday/bsvc Bank_Account_Name,omitempty"`
	RollNumber                  string                     `xml:"urn:com.workday/bsvc Roll_Number,omitempty"`
	IBAN                        string                     `xml:"urn:com.workday/bsvc IBAN,omitempty"`
	SWIFTBankIdentificationCode string                     `xml:"urn:com.workday/bsvc SWIFT_Bank_Identification_Code,omitempty"`
	Inactive                    *bool                      `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	BankInstructions            string                     `xml:"urn:com.workday/bsvc Bank_Instructions,omitempty"`
}

// Contains Type Code for Intraday Bank Statement Balance Code
type IntradayBankStatementBalanceTypeReferenceType struct {
	TypeCode string `xml:"urn:com.workday/bsvc Type_Code,omitempty"`
}

// This element contains all Intraday Bank Statement Data
type IntradayBankStatementDataType struct {
	BankStatementReferenceID             string                                     `xml:"urn:com.workday/bsvc Bank_Statement_Reference_ID,omitempty"`
	BankStatementFileReference           *BankStatementFileObjectType               `xml:"urn:com.workday/bsvc Bank_Statement_File_Reference,omitempty"`
	CustomerAccountNumber                string                                     `xml:"urn:com.workday/bsvc Customer_Account_Number,omitempty"`
	RoutingNumber                        string                                     `xml:"urn:com.workday/bsvc Routing_Number,omitempty"`
	BankIdentificationCode               string                                     `xml:"urn:com.workday/bsvc Bank_Identification_Code,omitempty"`
	IBAN                                 string                                     `xml:"urn:com.workday/bsvc IBAN,omitempty"`
	BankAccountReference                 *BankAccountObjectType                     `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	StatementDate                        *time.Time                                 `xml:"urn:com.workday/bsvc Statement_Date,omitempty"`
	CurrencyReference                    *CurrencyObjectType                        `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	IntradayBankStatementSummaryData     []IntradayBankStatementSummaryDataType     `xml:"urn:com.workday/bsvc Intraday_Bank_Statement_Summary_Data,omitempty"`
	IntradayBankStatementTransactionData []IntradayBankStatementTransactionDataType `xml:"urn:com.workday/bsvc Intraday_Bank_Statement_Transaction_Data,omitempty"`
}

func (t *IntradayBankStatementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T IntradayBankStatementDataType
	var layout struct {
		*T
		StatementDate *xsdDate `xml:"urn:com.workday/bsvc Statement_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StatementDate = (*xsdDate)(layout.T.StatementDate)
	return e.EncodeElement(layout, start)
}
func (t *IntradayBankStatementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T IntradayBankStatementDataType
	var overlay struct {
		*T
		StatementDate *xsdDate `xml:"urn:com.workday/bsvc Statement_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StatementDate = (*xsdDate)(overlay.T.StatementDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type IntradayBankStatementObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type IntradayBankStatementObjectType struct {
	ID         []IntradayBankStatementObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains Intraday Bank Statement Request Criteria
type IntradayBankStatementRequestCriteriaType struct {
	StatementDateFrom *time.Time `xml:"urn:com.workday/bsvc Statement_Date_From,omitempty"`
	StatementDateTo   *time.Time `xml:"urn:com.workday/bsvc Statement_Date_To,omitempty"`
}

func (t *IntradayBankStatementRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T IntradayBankStatementRequestCriteriaType
	var layout struct {
		*T
		StatementDateFrom *xsdDate `xml:"urn:com.workday/bsvc Statement_Date_From,omitempty"`
		StatementDateTo   *xsdDate `xml:"urn:com.workday/bsvc Statement_Date_To,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StatementDateFrom = (*xsdDate)(layout.T.StatementDateFrom)
	layout.StatementDateTo = (*xsdDate)(layout.T.StatementDateTo)
	return e.EncodeElement(layout, start)
}
func (t *IntradayBankStatementRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T IntradayBankStatementRequestCriteriaType
	var overlay struct {
		*T
		StatementDateFrom *xsdDate `xml:"urn:com.workday/bsvc Statement_Date_From,omitempty"`
		StatementDateTo   *xsdDate `xml:"urn:com.workday/bsvc Statement_Date_To,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StatementDateFrom = (*xsdDate)(overlay.T.StatementDateFrom)
	overlay.StatementDateTo = (*xsdDate)(overlay.T.StatementDateTo)
	return d.DecodeElement(&overlay, &start)
}

// Intraday Bank Statements Reference element contains the specific instance set containing the requested Intraday Bank Statements.
type IntradayBankStatementRequestReferencesType struct {
	IntradayBankStatementReference []IntradayBankStatementObjectType `xml:"urn:com.workday/bsvc Intraday_Bank_Statement_Reference"`
}

// contains Intraday Bank Statement Response Data
type IntradayBankStatementResponseDataType struct {
	IntradayBankStatement []IntradayBankStatementType `xml:"urn:com.workday/bsvc Intraday_Bank_Statement,omitempty"`
}

// Contains Include Reference Element
type IntradayBankStatementResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// This element holds the Intraday Bank Statement Summary Data.
type IntradayBankStatementSummaryDataType struct {
	IntradayBankStatementBalanceTypeReference *IntradayBankStatementBalanceTypeReferenceType `xml:"urn:com.workday/bsvc Intraday_Bank_Statement_Balance_Type_Reference,omitempty"`
	Amount                                    float64                                        `xml:"urn:com.workday/bsvc Amount,omitempty"`
	ItemCount                                 float64                                        `xml:"urn:com.workday/bsvc Item_Count,omitempty"`
	FundsAvailabilityData                     []FundsAvailabilityDataType                    `xml:"urn:com.workday/bsvc Funds_Availability_Data,omitempty"`
	AlternateBalanceTypeCode                  string                                         `xml:"urn:com.workday/bsvc Alternate_Balance_Type_Code,omitempty"`
	Debit                                     *bool                                          `xml:"urn:com.workday/bsvc Debit,omitempty"`
}

// This element contains Intraday Bank Statement transaction Data
type IntradayBankStatementTransactionDataType struct {
	IntradayBankStatementTransactionTypeReference *IntradayBankStatementTransactionTypeReferenceType `xml:"urn:com.workday/bsvc Intraday_Bank_Statement_Transaction_Type_Reference,omitempty"`
	TransactionDate                               *time.Time                                         `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
	Amount                                        float64                                            `xml:"urn:com.workday/bsvc Amount,omitempty"`
	BankReferenceNumber                           string                                             `xml:"urn:com.workday/bsvc Bank_Reference_Number,omitempty"`
	CustomerReferenceNumber                       string                                             `xml:"urn:com.workday/bsvc Customer_Reference_Number,omitempty"`
	Addenda                                       string                                             `xml:"urn:com.workday/bsvc Addenda,omitempty"`
	FundsAvailabilityData                         []FundsAvailabilityDataType                        `xml:"urn:com.workday/bsvc Funds_Availability_Data,omitempty"`
	AlternateTransactionTypeCode                  string                                             `xml:"urn:com.workday/bsvc Alternate_Transaction_Type_Code,omitempty"`
	OriginatingCurrencyAmount                     float64                                            `xml:"urn:com.workday/bsvc Originating_Currency_Amount,omitempty"`
	OriginatingCurrency                           string                                             `xml:"urn:com.workday/bsvc Originating_Currency,omitempty"`
	CurrencyExchangerate                          float64                                            `xml:"urn:com.workday/bsvc Currency_Exchange_rate,omitempty"`
	Debit                                         *bool                                              `xml:"urn:com.workday/bsvc Debit,omitempty"`
	ValueDate                                     *time.Time                                         `xml:"urn:com.workday/bsvc Value_Date,omitempty"`
	ReversalIndicator                             *bool                                              `xml:"urn:com.workday/bsvc Reversal_Indicator,omitempty"`
	AdditionalReferenceField1                     string                                             `xml:"urn:com.workday/bsvc Additional_Reference_Field_1,omitempty"`
	AdditionalReferenceField2                     string                                             `xml:"urn:com.workday/bsvc Additional_Reference_Field_2,omitempty"`
	AdditionalReferenceField3                     string                                             `xml:"urn:com.workday/bsvc Additional_Reference_Field_3,omitempty"`
	AdditionalReferenceField4                     string                                             `xml:"urn:com.workday/bsvc Additional_Reference_Field_4,omitempty"`
	AdditionalReferenceField5                     string                                             `xml:"urn:com.workday/bsvc Additional_Reference_Field_5,omitempty"`
	TransactionStatus                             string                                             `xml:"urn:com.workday/bsvc Transaction_Status,omitempty"`
}

func (t *IntradayBankStatementTransactionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T IntradayBankStatementTransactionDataType
	var layout struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
		ValueDate       *xsdDate `xml:"urn:com.workday/bsvc Value_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TransactionDate = (*xsdDate)(layout.T.TransactionDate)
	layout.ValueDate = (*xsdDate)(layout.T.ValueDate)
	return e.EncodeElement(layout, start)
}
func (t *IntradayBankStatementTransactionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T IntradayBankStatementTransactionDataType
	var overlay struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
		ValueDate       *xsdDate `xml:"urn:com.workday/bsvc Value_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionDate = (*xsdDate)(overlay.T.TransactionDate)
	overlay.ValueDate = (*xsdDate)(overlay.T.ValueDate)
	return d.DecodeElement(&overlay, &start)
}

// This Element contains Type code for Intraday Bank Statement Account Transactions.
type IntradayBankStatementTransactionTypeReferenceType struct {
	TypeCode string `xml:"urn:com.workday/bsvc Type_Code,omitempty"`
}

// contains Intraday Bank Statement Instance
type IntradayBankStatementType struct {
	IntradayBankStatementReference *IntradayBankStatementObjectType `xml:"urn:com.workday/bsvc Intraday_Bank_Statement_Reference,omitempty"`
	IntradayBankStatementData      []IntradayBankStatementDataType  `xml:"urn:com.workday/bsvc Intraday_Bank_Statement_Data,omitempty"`
}

// Investment Pool Transfer Data contains essential information about an Investment Pool Transfer
type InvestmentPoolAdjustmentDataType struct {
	TransactionDate                   time.Time                           `xml:"urn:com.workday/bsvc Transaction_Date"`
	ID                                string                              `xml:"urn:com.workday/bsvc ID,omitempty"`
	GiftReference                     *GiftObjectType                     `xml:"urn:com.workday/bsvc Gift_Reference"`
	UnitstoAdjust                     float64                             `xml:"urn:com.workday/bsvc Units_to_Adjust,omitempty"`
	StatusReference                   *DocumentStatusObjectType           `xml:"urn:com.workday/bsvc Status_Reference,omitempty"`
	AdjustmentFairMarketValue         float64                             `xml:"urn:com.workday/bsvc Adjustment_Fair_Market_Value,omitempty"`
	OriginalTotalUnits                float64                             `xml:"urn:com.workday/bsvc Original_Total_Units,omitempty"`
	CurrencyReference                 *CurrencyObjectType                 `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	InvestmentPoolReference           *InvestmentPoolObjectType           `xml:"urn:com.workday/bsvc Investment_Pool_Reference,omitempty"`
	UnitPrice                         float64                             `xml:"urn:com.workday/bsvc Unit_Price,omitempty"`
	InvestmentPoolAdjustmentReference *InvestmentPoolAdjustmentObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Adjustment_Reference,omitempty"`
}

func (t *InvestmentPoolAdjustmentDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InvestmentPoolAdjustmentDataType
	var layout struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date"`
	}
	layout.T = (*T)(t)
	layout.TransactionDate = (*xsdDate)(&layout.T.TransactionDate)
	return e.EncodeElement(layout, start)
}
func (t *InvestmentPoolAdjustmentDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InvestmentPoolAdjustmentDataType
	var overlay struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionDate = (*xsdDate)(&overlay.T.TransactionDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type InvestmentPoolAdjustmentObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InvestmentPoolAdjustmentObjectType struct {
	ID         []InvestmentPoolAdjustmentObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing search criteria for Investment Pool Adjustments.
type InvestmentPoolAdjustmentRequestCriteriaType struct {
}

// References to existing Investment Pool Adjustments.
type InvestmentPoolAdjustmentRequestReferencesType struct {
	InvestmentPoolAdjustmentReference []InvestmentPoolAdjustmentObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Adjustment_Reference"`
}

// Element containing data for existing Investment Pool Adjustments.
type InvestmentPoolAdjustmentResponseDataType struct {
	InvestmentPoolAdjustment []InvestmentPoolAdjustmentType `xml:"urn:com.workday/bsvc Investment_Pool_Adjustment,omitempty"`
}

// Element containing data for existing Investment Pool Adjustments.
type InvestmentPoolAdjustmentResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing reference to an existing Investment Pool Adjustmentr and data
type InvestmentPoolAdjustmentType struct {
	InvestmentPoolAdjustmentReference *InvestmentPoolAdjustmentObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Adjustment_Reference,omitempty"`
	InvestmentPoolAdjustmentData      []InvestmentPoolAdjustmentDataType  `xml:"urn:com.workday/bsvc Investment_Pool_Adjustment_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type InvestmentPoolObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InvestmentPoolObjectType struct {
	ID         []InvestmentPoolObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Investment Pool Payout Criteria Data
type InvestmentPoolPayoutCriteriaDataType struct {
	Payoutperunit                 float64                         `xml:"urn:com.workday/bsvc Payout_per_unit,omitempty"`
	AdministrativeFeeperunit      float64                         `xml:"urn:com.workday/bsvc Administrative_Fee_per_unit,omitempty"`
	EndDate                       *time.Time                      `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	InvestmentPoolPayoutReference *InvestmentPoolPayoutObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Payout_Reference,omitempty"`
}

func (t *InvestmentPoolPayoutCriteriaDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InvestmentPoolPayoutCriteriaDataType
	var layout struct {
		*T
		EndDate *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *InvestmentPoolPayoutCriteriaDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InvestmentPoolPayoutCriteriaDataType
	var overlay struct {
		*T
		EndDate *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type InvestmentPoolPayoutObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InvestmentPoolPayoutObjectType struct {
	ID         []InvestmentPoolPayoutObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Investment Pool Purchase Data contains essential information about an Investment Pool Purchase
type InvestmentPoolPurchaseDataType struct {
	ID                             string                               `xml:"urn:com.workday/bsvc ID,omitempty"`
	Submit                         *bool                                `xml:"urn:com.workday/bsvc Submit,omitempty"`
	GiftReference                  *GiftObjectType                      `xml:"urn:com.workday/bsvc Gift_Reference"`
	TransactionDate                time.Time                            `xml:"urn:com.workday/bsvc Transaction_Date"`
	PurchaseDate                   time.Time                            `xml:"urn:com.workday/bsvc Purchase_Date"`
	WorktagsReference              []AccountingWorktagObjectType        `xml:"urn:com.workday/bsvc Worktags_Reference,omitempty"`
	InvestmentPoolPurchaseLineData []InvestmentPoolPurchaseLineDataType `xml:"urn:com.workday/bsvc Investment_Pool_Purchase_Line_Data,omitempty"`
}

func (t *InvestmentPoolPurchaseDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InvestmentPoolPurchaseDataType
	var layout struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date"`
		PurchaseDate    *xsdDate `xml:"urn:com.workday/bsvc Purchase_Date"`
	}
	layout.T = (*T)(t)
	layout.TransactionDate = (*xsdDate)(&layout.T.TransactionDate)
	layout.PurchaseDate = (*xsdDate)(&layout.T.PurchaseDate)
	return e.EncodeElement(layout, start)
}
func (t *InvestmentPoolPurchaseDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InvestmentPoolPurchaseDataType
	var overlay struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date"`
		PurchaseDate    *xsdDate `xml:"urn:com.workday/bsvc Purchase_Date"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionDate = (*xsdDate)(&overlay.T.TransactionDate)
	overlay.PurchaseDate = (*xsdDate)(&overlay.T.PurchaseDate)
	return d.DecodeElement(&overlay, &start)
}

// Investment Pool Purchase Record Line Data contains essential information on one Investment Pool Purchase Record Line.
type InvestmentPoolPurchaseLineDataType struct {
	DonorContributionReference *DonorContributionObjectType `xml:"urn:com.workday/bsvc Donor_Contribution_Reference"`
	OverrideDefaultUnitPrice   *bool                        `xml:"urn:com.workday/bsvc Override_Default_Unit_Price,omitempty"`
	PurchaseDate               *time.Time                   `xml:"urn:com.workday/bsvc Purchase_Date,omitempty"`
	UnitPrice                  float64                      `xml:"urn:com.workday/bsvc Unit_Price,omitempty"`
}

func (t *InvestmentPoolPurchaseLineDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InvestmentPoolPurchaseLineDataType
	var layout struct {
		*T
		PurchaseDate *xsdDate `xml:"urn:com.workday/bsvc Purchase_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PurchaseDate = (*xsdDate)(layout.T.PurchaseDate)
	return e.EncodeElement(layout, start)
}
func (t *InvestmentPoolPurchaseLineDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InvestmentPoolPurchaseLineDataType
	var overlay struct {
		*T
		PurchaseDate *xsdDate `xml:"urn:com.workday/bsvc Purchase_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PurchaseDate = (*xsdDate)(overlay.T.PurchaseDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type InvestmentPoolPurchaseObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InvestmentPoolPurchaseObjectType struct {
	ID         []InvestmentPoolPurchaseObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing search criteria for Investment Pool Purchases.
type InvestmentPoolPurchaseRequestCriteriaType struct {
}

// Element containing references to existing Investment Pool Purchase Records.
type InvestmentPoolPurchaseRequestReferencesType struct {
	InvestmentPoolPurchaseRecordReference []InvestmentPoolPurchaseObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Purchase_Record_Reference"`
}

// Element containing data for existing Investment Pool Purchases.
type InvestmentPoolPurchaseResponseDataType struct {
	InvestmentPoolPurchase []InvestmentPoolPurchaseType `xml:"urn:com.workday/bsvc Investment_Pool_Purchase,omitempty"`
}

// Element containing data for existing Investment Pool Purchases.
type InvestmentPoolPurchaseResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing reference to an existing Investment Pool Purchase Record and data
type InvestmentPoolPurchaseType struct {
	InvestmentPoolPurchaseReference *InvestmentPoolPurchaseObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Purchase_Reference,omitempty"`
	InvestmentPoolPurchaseData      *InvestmentPoolPurchaseDataType   `xml:"urn:com.workday/bsvc Investment_Pool_Purchase_Data,omitempty"`
	InvestmentPoolPurchaseValue     *InvestmentPoolPurchaseValueType  `xml:"urn:com.workday/bsvc Investment_Pool_Purchase_Value,omitempty"`
}

// Element containing additional Investment Pool Purchase Record data.
type InvestmentPoolPurchaseValueType struct {
	Status                           string                      `xml:"urn:com.workday/bsvc Status,omitempty"`
	InvestmentPoolValuationReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Valuation_Reference,omitempty"`
	DefaultUnitPrice                 float64                     `xml:"urn:com.workday/bsvc Default_Unit_Price,omitempty"`
	UnitsPurchased                   float64                     `xml:"urn:com.workday/bsvc Units_Purchased,omitempty"`
	TotalAmountPurchased             float64                     `xml:"urn:com.workday/bsvc Total_Amount_Purchased,omitempty"`
	CurrencyReference                *CurrencyObjectType         `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	Reinvestment                     *bool                       `xml:"urn:com.workday/bsvc Reinvestment,omitempty"`
}

// Data Element having data for Selling Investment Pool Units
type InvestmentPoolSaleDataType struct {
	ID                               string                           `xml:"urn:com.workday/bsvc ID,omitempty"`
	TransactionDate                  time.Time                        `xml:"urn:com.workday/bsvc Transaction_Date"`
	StatusReference                  *DocumentStatusObjectType        `xml:"urn:com.workday/bsvc Status_Reference,omitempty"`
	InvestmentPoolReference          *InvestmentPoolObjectType        `xml:"urn:com.workday/bsvc Investment_Pool_Reference,omitempty"`
	InvestmentPoolValuationReference *UniqueIdentifierObjectType      `xml:"urn:com.workday/bsvc Investment_Pool_Valuation_Reference,omitempty"`
	CurrencyReference                *CurrencyObjectType              `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	GiftReference                    *GiftObjectType                  `xml:"urn:com.workday/bsvc Gift_Reference"`
	UnitstoSell                      float64                          `xml:"urn:com.workday/bsvc Units_to_Sell,omitempty"`
	UnitPrice                        float64                          `xml:"urn:com.workday/bsvc Unit_Price,omitempty"`
	NetBookValue                     float64                          `xml:"urn:com.workday/bsvc Net_Book_Value,omitempty"`
	UnitSelectionMethodReference     *UniqueIdentifierObjectType      `xml:"urn:com.workday/bsvc Unit_Selection_Method_Reference"`
	FairMarketValue                  float64                          `xml:"urn:com.workday/bsvc Fair_Market_Value,omitempty"`
	UnitSourceData                   []InvestmentPoolSaleLineDataType `xml:"urn:com.workday/bsvc Unit_Source_Data,omitempty"`
}

func (t *InvestmentPoolSaleDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InvestmentPoolSaleDataType
	var layout struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date"`
	}
	layout.T = (*T)(t)
	layout.TransactionDate = (*xsdDate)(&layout.T.TransactionDate)
	return e.EncodeElement(layout, start)
}
func (t *InvestmentPoolSaleDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InvestmentPoolSaleDataType
	var overlay struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionDate = (*xsdDate)(&overlay.T.TransactionDate)
	return d.DecodeElement(&overlay, &start)
}

// Investment Pool Sale Line Data Element
type InvestmentPoolSaleLineDataType struct {
	InvestmentPoolTransactionReference *InvestmentPoolTransactionObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Transaction_Reference,omitempty"`
	DonorContributionReference         *DonorContributionObjectType         `xml:"urn:com.workday/bsvc Donor_Contribution_Reference,omitempty"`
	Type                               string                               `xml:"urn:com.workday/bsvc Type,omitempty"`
	TransactionDate                    *time.Time                           `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
	Units                              float64                              `xml:"urn:com.workday/bsvc Units,omitempty"`
	UnitPrice                          float64                              `xml:"urn:com.workday/bsvc Unit_Price,omitempty"`
	Amount                             float64                              `xml:"urn:com.workday/bsvc Amount,omitempty"`
	AmountWithdrawn                    float64                              `xml:"urn:com.workday/bsvc Amount_Withdrawn,omitempty"`
}

func (t *InvestmentPoolSaleLineDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InvestmentPoolSaleLineDataType
	var layout struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TransactionDate = (*xsdDate)(layout.T.TransactionDate)
	return e.EncodeElement(layout, start)
}
func (t *InvestmentPoolSaleLineDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InvestmentPoolSaleLineDataType
	var overlay struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionDate = (*xsdDate)(overlay.T.TransactionDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type InvestmentPoolSaleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InvestmentPoolSaleObjectType struct {
	ID         []InvestmentPoolSaleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Investment Pool Sale Request Criteria
type InvestmentPoolSaleRequestCriteriaType struct {
}

// Investment Pool Sale Request References
type InvestmentPoolSaleRequestReferencesType struct {
	InvestmentPoolSaleReference []InvestmentPoolSaleObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Sale_Reference"`
}

// Investment Pool Sale Response Data
type InvestmentPoolSaleResponseDataType struct {
	InvestmentPoolSale []InvestmentPoolSaleType `xml:"urn:com.workday/bsvc Investment_Pool_Sale,omitempty"`
}

// Investment Pool Sale Response Group
type InvestmentPoolSaleResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Investment Pool Sale Element
type InvestmentPoolSaleType struct {
	InvestmentPoolSaleReference *InvestmentPoolSaleObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Sale_Reference,omitempty"`
	InvestmentPoolSaleData      []InvestmentPoolSaleDataType  `xml:"urn:com.workday/bsvc Investment_Pool_Sale_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type InvestmentPoolTransactionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InvestmentPoolTransactionObjectType struct {
	ID         []InvestmentPoolTransactionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Investment Pool Transfer Data contains essential information about an Investment Pool Transfer
type InvestmentPoolTransferDataType struct {
	ID                               string                               `xml:"urn:com.workday/bsvc ID,omitempty"`
	TransactionDate                  time.Time                            `xml:"urn:com.workday/bsvc Transaction_Date"`
	Status                           string                               `xml:"urn:com.workday/bsvc Status,omitempty"`
	SourceGiftReference              *GiftObjectType                      `xml:"urn:com.workday/bsvc Source_Gift_Reference"`
	DestinationGiftReference         *GiftObjectType                      `xml:"urn:com.workday/bsvc Destination_Gift_Reference"`
	InvestmentPoolReference          *InvestmentPoolObjectType            `xml:"urn:com.workday/bsvc Investment_Pool_Reference,omitempty"`
	InvestmentPoolValuationReference *UniqueIdentifierObjectType          `xml:"urn:com.workday/bsvc Investment_Pool_Valuation_Reference,omitempty"`
	CurrencyReference                *CurrencyObjectType                  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	UnitPrice                        float64                              `xml:"urn:com.workday/bsvc Unit_Price,omitempty"`
	TotalUnitsTransferred            float64                              `xml:"urn:com.workday/bsvc Total_Units_Transferred,omitempty"`
	FairMarketValue                  float64                              `xml:"urn:com.workday/bsvc Fair_Market_Value,omitempty"`
	NetBookValue                     float64                              `xml:"urn:com.workday/bsvc Net_Book_Value,omitempty"`
	UnitSelectionMethodReference     *UniqueIdentifierObjectType          `xml:"urn:com.workday/bsvc Unit_Selection_Method_Reference"`
	WorktagsReference                []AccountingWorktagObjectType        `xml:"urn:com.workday/bsvc Worktags_Reference,omitempty"`
	UnitSourceData                   []InvestmentPoolTransferLineDataType `xml:"urn:com.workday/bsvc Unit_Source_Data,omitempty"`
}

func (t *InvestmentPoolTransferDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InvestmentPoolTransferDataType
	var layout struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date"`
	}
	layout.T = (*T)(t)
	layout.TransactionDate = (*xsdDate)(&layout.T.TransactionDate)
	return e.EncodeElement(layout, start)
}
func (t *InvestmentPoolTransferDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InvestmentPoolTransferDataType
	var overlay struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionDate = (*xsdDate)(&overlay.T.TransactionDate)
	return d.DecodeElement(&overlay, &start)
}

// Unit Source Data contains essential information on one unit source for the Investment Pool Transfer
type InvestmentPoolTransferLineDataType struct {
	InvestmentPoolTransactionReference *InvestmentPoolTransactionObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Transaction_Reference,omitempty"`
	DonorContributionReference         *DonorContributionObjectType         `xml:"urn:com.workday/bsvc Donor_Contribution_Reference,omitempty"`
	Type                               string                               `xml:"urn:com.workday/bsvc Type,omitempty"`
	TransactionDate                    *time.Time                           `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
	UnitPrice                          float64                              `xml:"urn:com.workday/bsvc Unit_Price,omitempty"`
	Units                              float64                              `xml:"urn:com.workday/bsvc Units,omitempty"`
	Amount                             float64                              `xml:"urn:com.workday/bsvc Amount,omitempty"`
	AmountWithdrawn                    float64                              `xml:"urn:com.workday/bsvc Amount_Withdrawn,omitempty"`
}

func (t *InvestmentPoolTransferLineDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InvestmentPoolTransferLineDataType
	var layout struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TransactionDate = (*xsdDate)(layout.T.TransactionDate)
	return e.EncodeElement(layout, start)
}
func (t *InvestmentPoolTransferLineDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InvestmentPoolTransferLineDataType
	var overlay struct {
		*T
		TransactionDate *xsdDate `xml:"urn:com.workday/bsvc Transaction_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TransactionDate = (*xsdDate)(overlay.T.TransactionDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type InvestmentPoolTransferObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InvestmentPoolTransferObjectType struct {
	ID         []InvestmentPoolTransferObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing search criteria for Investment Pool Transfers.
type InvestmentPoolTransferRequestCriteriaType struct {
}

// References to existing Investment Pool Transfers.
type InvestmentPoolTransferRequestReferencesType struct {
	InvestmentPoolTransferReference []InvestmentPoolTransferObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Transfer_Reference"`
}

// Element containing data for existing Investment Pool Transfers.
type InvestmentPoolTransferResponseDataType struct {
	InvestmentPoolTransfer []InvestmentPoolTransferType `xml:"urn:com.workday/bsvc Investment_Pool_Transfer,omitempty"`
}

// Element containing data for existing Investment Pool Transfers.
type InvestmentPoolTransferResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing reference to an existing Investment Pool Transfer and data
type InvestmentPoolTransferType struct {
	InvestmentPoolTransferReference *InvestmentPoolTransferObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Transfer_Reference,omitempty"`
	InvestmentPoolTransferData      []InvestmentPoolTransferDataType  `xml:"urn:com.workday/bsvc Investment_Pool_Transfer_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type InvestmentProfileObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InvestmentProfileObjectType struct {
	ID         []InvestmentProfileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing essential Investment Statement data
type InvestmentStatementDataType struct {
	ID                          string                            `xml:"urn:com.workday/bsvc ID,omitempty"`
	Submit                      *bool                             `xml:"urn:com.workday/bsvc Submit,omitempty"`
	InvestmentProfileReference  *InvestmentProfileObjectType      `xml:"urn:com.workday/bsvc Investment_Profile_Reference"`
	StatementBeginningDate      *time.Time                        `xml:"urn:com.workday/bsvc Statement_Beginning_Date,omitempty"`
	StatementDate               time.Time                         `xml:"urn:com.workday/bsvc Statement_Date"`
	EndingBalance               float64                           `xml:"urn:com.workday/bsvc Ending_Balance,omitempty"`
	Memo                        string                            `xml:"urn:com.workday/bsvc Memo,omitempty"`
	InvestmentStatementLineData []InvestmentStatementLineDataType `xml:"urn:com.workday/bsvc Investment_Statement_Line_Data,omitempty"`
	AttachmentData              []FinancialsAttachmentDataType    `xml:"urn:com.workday/bsvc Attachment_Data,omitempty"`
}

func (t *InvestmentStatementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InvestmentStatementDataType
	var layout struct {
		*T
		StatementBeginningDate *xsdDate `xml:"urn:com.workday/bsvc Statement_Beginning_Date,omitempty"`
		StatementDate          *xsdDate `xml:"urn:com.workday/bsvc Statement_Date"`
	}
	layout.T = (*T)(t)
	layout.StatementBeginningDate = (*xsdDate)(layout.T.StatementBeginningDate)
	layout.StatementDate = (*xsdDate)(&layout.T.StatementDate)
	return e.EncodeElement(layout, start)
}
func (t *InvestmentStatementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InvestmentStatementDataType
	var overlay struct {
		*T
		StatementBeginningDate *xsdDate `xml:"urn:com.workday/bsvc Statement_Beginning_Date,omitempty"`
		StatementDate          *xsdDate `xml:"urn:com.workday/bsvc Statement_Date"`
	}
	overlay.T = (*T)(t)
	overlay.StatementBeginningDate = (*xsdDate)(overlay.T.StatementBeginningDate)
	overlay.StatementDate = (*xsdDate)(&overlay.T.StatementDate)
	return d.DecodeElement(&overlay, &start)
}

// Element containing essential Investment Statement Line data
type InvestmentStatementLineDataType struct {
	InvestmentStatementLineTypeReference *InvestmentStatementLineTypeObjectType `xml:"urn:com.workday/bsvc Investment_Statement_Line_Type_Reference"`
	Amount                               float64                                `xml:"urn:com.workday/bsvc Amount,omitempty"`
	Memo                                 string                                 `xml:"urn:com.workday/bsvc Memo,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type InvestmentStatementLineTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InvestmentStatementLineTypeObjectType struct {
	ID         []InvestmentStatementLineTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type InvestmentStatementObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InvestmentStatementObjectType struct {
	ID         []InvestmentStatementObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing search criteria for Investment Statements.
type InvestmentStatementRequestCriteriaType struct {
}

// Element containing references to existing Investment Statements.
type InvestmentStatementRequestReferencesType struct {
	InvestmentStatementReference []InvestmentStatementObjectType `xml:"urn:com.workday/bsvc Investment_Statement_Reference"`
}

// Element containing data for existing Investment Statements.
type InvestmentStatementResponseDataType struct {
	InvestmentStatement []InvestmentStatementType `xml:"urn:com.workday/bsvc Investment_Statement,omitempty"`
}

// Wrapper element around a list of elements representing the amount of data that should be included in
type InvestmentStatementResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing reference to existing Investment Statement and data.
type InvestmentStatementType struct {
	InvestmentStatementReference *InvestmentStatementObjectType `xml:"urn:com.workday/bsvc Investment_Statement_Reference,omitempty"`
	InvestmentStatementData      *InvestmentStatementDataType   `xml:"urn:com.workday/bsvc Investment_Statement_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type InvestorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type InvestorObjectType struct {
	ID         []InvestorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type JournalSourceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type JournalSourceObjectType struct {
	ID         []JournalSourceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LedgerAccountObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type LedgerAccountObjectType struct {
	ID         []LedgerAccountObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the legal name for a person.  A person must name one and only one legal name.
type LegalNameDataType struct {
	NameDetailData *PersonNameDetailDataType `xml:"urn:com.workday/bsvc Name_Detail_Data"`
}

// Contains the components of a name in local script, such as the First Name and Last Name, for supporting countries.
type LocalPersonNameDetailDataType struct {
	FirstName          string `xml:"urn:com.workday/bsvc First_Name,omitempty"`
	MiddleName         string `xml:"urn:com.workday/bsvc Middle_Name,omitempty"`
	LastName           string `xml:"urn:com.workday/bsvc Last_Name,omitempty"`
	SecondaryLastName  string `xml:"urn:com.workday/bsvc Secondary_Last_Name,omitempty"`
	FirstName2         string `xml:"urn:com.workday/bsvc First_Name_2,omitempty"`
	MiddleName2        string `xml:"urn:com.workday/bsvc Middle_Name_2,omitempty"`
	LastName2          string `xml:"urn:com.workday/bsvc Last_Name_2,omitempty"`
	SecondaryLastName2 string `xml:"urn:com.workday/bsvc Secondary_Last_Name_2,omitempty"`
	LocalName          string `xml:"urn:com.workday/bsvc Local_Name,attr,omitempty"`
	LocalScript        string `xml:"urn:com.workday/bsvc Local_Script,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type LocationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type LocationObjectType struct {
	ID         []LocationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing the lockbox data for the bank account
type LockboxDataType struct {
	LockboxID      string                      `xml:"urn:com.workday/bsvc Lockbox_ID,omitempty"`
	LockboxName    string                      `xml:"urn:com.workday/bsvc Lockbox_Name"`
	LockboxAddress *AddressInformationDataType `xml:"urn:com.workday/bsvc Lockbox_Address"`
}

// Mandate Data
type MandateDataType struct {
	Identification     string     `xml:"urn:com.workday/bsvc Identification,omitempty"`
	DateofSignature    *time.Time `xml:"urn:com.workday/bsvc Date_of_Signature,omitempty"`
	CreditorIdentifier string     `xml:"urn:com.workday/bsvc Creditor_Identifier,omitempty"`
	SequenceType       string     `xml:"urn:com.workday/bsvc Sequence_Type,omitempty"`
}

func (t *MandateDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MandateDataType
	var layout struct {
		*T
		DateofSignature *xsdDate `xml:"urn:com.workday/bsvc Date_of_Signature,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateofSignature = (*xsdDate)(layout.T.DateofSignature)
	return e.EncodeElement(layout, start)
}
func (t *MandateDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MandateDataType
	var overlay struct {
		*T
		DateofSignature *xsdDate `xml:"urn:com.workday/bsvc Date_of_Signature,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateofSignature = (*xsdDate)(overlay.T.DateofSignature)
	return d.DecodeElement(&overlay, &start)
}

// On Behalf Of Company Data
type OnBehalfOfCompanyWWSDataType struct {
	OnBehalfOfCompanyName               string                      `xml:"urn:com.workday/bsvc On_Behalf_Of_Company_Name,omitempty"`
	OnBehalfOfCompanyTaxID              string                      `xml:"urn:com.workday/bsvc On_Behalf_Of_Company_Tax_ID,omitempty"`
	OnBehalfOfCompanyAddressData        *AddressInformationDataType `xml:"urn:com.workday/bsvc On_Behalf_Of_Company_Address_Data,omitempty"`
	OnBehalfOfCompanyPrimaryPhoneNumber string                      `xml:"urn:com.workday/bsvc On_Behalf_Of_Company_Primary_Phone_Number,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type OrganizationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type OrganizationObjectType struct {
	ID         []OrganizationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Originating Party Bank Data. Can be conditionally returned by setting the "Include Originating Bank data" grouping flag.
type OriginatingPartyBankDataType struct {
	BankAccountReference          []BankAccountObjectType         `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	AccountNickname               string                          `xml:"urn:com.workday/bsvc Account_Nickname,omitempty"`
	AccountNumber                 string                          `xml:"urn:com.workday/bsvc Account_Number,omitempty"`
	AccountTypeCode               string                          `xml:"urn:com.workday/bsvc Account_Type_Code,omitempty"`
	FinancialInstitutionReference *FinancialInstitutionObjectType `xml:"urn:com.workday/bsvc Financial_Institution_Reference,omitempty"`
	IBAN                          string                          `xml:"urn:com.workday/bsvc IBAN,omitempty"`
	RoutingTransitNumber          string                          `xml:"urn:com.workday/bsvc Routing_Transit_Number,omitempty"`
	BankIdentificationCode        string                          `xml:"urn:com.workday/bsvc Bank_Identification_Code,omitempty"`
	BranchName                    string                          `xml:"urn:com.workday/bsvc Branch_Name,omitempty"`
	BranchIDNumber                string                          `xml:"urn:com.workday/bsvc Branch_ID_Number,omitempty"`
	CountryReference              *CountryObjectType              `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CurrencyReference             *CurrencyObjectType             `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	BankAccountName               string                          `xml:"urn:com.workday/bsvc Bank_Account_Name,omitempty"`
	CheckDigit                    string                          `xml:"urn:com.workday/bsvc Check_Digit,omitempty"`
	RollNumber                    string                          `xml:"urn:com.workday/bsvc Roll_Number,omitempty"`
	BankAddressData               []AddressInformationDataType    `xml:"urn:com.workday/bsvc Bank_Address_Data,omitempty"`
}

// Originating Party Data
type OriginatingPartyWWSDataType struct {
	OrganizationName               string                      `xml:"urn:com.workday/bsvc Organization_Name,omitempty"`
	TransactionTaxID               string                      `xml:"urn:com.workday/bsvc Transaction_Tax_ID,omitempty"`
	EmailAddressData               []EmailAddressDataType      `xml:"urn:com.workday/bsvc Email_Address_Data,omitempty"`
	AddressInformationData         *AddressInformationDataType `xml:"urn:com.workday/bsvc Address_Information_Data,omitempty"`
	OrganizationPrimaryPhoneNumber string                      `xml:"urn:com.workday/bsvc Organization_Primary_Phone_Number,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type OutsourcedPaymentGroupObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type OutsourcedPaymentGroupObjectType struct {
	ID         []OutsourcedPaymentGroupObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ParsingRuleSetObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ParsingRuleSetObjectType struct {
	ID         []ParsingRuleSetObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Pay Component Reference
type PayComponentReferenceType struct {
	Earning   *EarningReferenceType   `xml:"urn:com.workday/bsvc Earning"`
	Deduction *DeductionReferenceType `xml:"urn:com.workday/bsvc Deduction"`
}

// Contains a unique identifier for an instance of an object.
type PayGroupDetailObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayGroupDetailObjectType struct {
	ID         []PayGroupDetailObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayRateTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayRateTypeObjectType struct {
	ID         []PayRateTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayRunGroupSelectionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayRunGroupSelectionObjectType struct {
	ID         []PayRunGroupSelectionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Container element for Pay Type Payment Election Option
type PayTypePaymentElectionOptionDataType struct {
	PaymentElectionRuleReference                    *PaymentElectionRuleObjectType             `xml:"urn:com.workday/bsvc Payment_Election_Rule_Reference"`
	PayTypePaymentElectionOptionLineReplacementData []PayTypePaymentElectionOptionLineDataType `xml:"urn:com.workday/bsvc Pay_Type_Payment_Election_Option_Line_Replacement_Data,omitempty"`
}

// Container element for Pay Type Payment Election Option Line
type PayTypePaymentElectionOptionLineDataType struct {
	CountryReference     *CountryObjectType      `xml:"urn:com.workday/bsvc Country_Reference"`
	CurrencyReference    *CurrencyObjectType     `xml:"urn:com.workday/bsvc Currency_Reference"`
	PaymentTypeReference []PaymentTypeObjectType `xml:"urn:com.workday/bsvc Payment_Type_Reference"`
}

// Contains a unique identifier for an instance of an object.
type PayeeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayeeObjectType struct {
	ID         []PayeeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payment Acknowledgement Data add or update payment acknowledgements that allow the update of payment status and error code information.
type PaymentAcknowledgementDataType struct {
	PaymentProcessingID         string                                  `xml:"urn:com.workday/bsvc Payment_Processing_ID,omitempty"`
	PaymentMessageBatchID       string                                  `xml:"urn:com.workday/bsvc Payment_Message_Batch_ID,omitempty"`
	OriginalPaymentID           string                                  `xml:"urn:com.workday/bsvc Original_Payment_ID,omitempty"`
	AdditionalReconciliationID  string                                  `xml:"urn:com.workday/bsvc Additional_Reconciliation_ID,omitempty"`
	MerchantAccountID           string                                  `xml:"urn:com.workday/bsvc Merchant_Account_ID,omitempty"`
	PaymentStatusReference      *PaymentAcknowledgementStatusObjectType `xml:"urn:com.workday/bsvc Payment_Status_Reference,omitempty"`
	PaymentValueDate            *time.Time                              `xml:"urn:com.workday/bsvc Payment_Value_Date,omitempty"`
	StatusReasonCode            string                                  `xml:"urn:com.workday/bsvc Status_Reason_Code,omitempty"`
	StatusReasonCodeDescription string                                  `xml:"urn:com.workday/bsvc Status_Reason_Code_Description,omitempty"`
	ReasonCode                  string                                  `xml:"urn:com.workday/bsvc Reason_Code,omitempty"`
	ReasonCodeDescription       string                                  `xml:"urn:com.workday/bsvc Reason_Code_Description,omitempty"`
	OriginalPaymentAmount       float64                                 `xml:"urn:com.workday/bsvc Original_Payment_Amount,omitempty"`
	CurrencyReference           *CurrencyObjectType                     `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	OriginalPaymentDate         *time.Time                              `xml:"urn:com.workday/bsvc Original_Payment_Date,omitempty"`
}

func (t *PaymentAcknowledgementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PaymentAcknowledgementDataType
	var layout struct {
		*T
		PaymentValueDate    *xsdDate `xml:"urn:com.workday/bsvc Payment_Value_Date,omitempty"`
		OriginalPaymentDate *xsdDate `xml:"urn:com.workday/bsvc Original_Payment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PaymentValueDate = (*xsdDate)(layout.T.PaymentValueDate)
	layout.OriginalPaymentDate = (*xsdDate)(layout.T.OriginalPaymentDate)
	return e.EncodeElement(layout, start)
}
func (t *PaymentAcknowledgementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PaymentAcknowledgementDataType
	var overlay struct {
		*T
		PaymentValueDate    *xsdDate `xml:"urn:com.workday/bsvc Payment_Value_Date,omitempty"`
		OriginalPaymentDate *xsdDate `xml:"urn:com.workday/bsvc Original_Payment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentValueDate = (*xsdDate)(overlay.T.PaymentValueDate)
	overlay.OriginalPaymentDate = (*xsdDate)(overlay.T.OriginalPaymentDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains data for creating or updating an Payments Acknowledgment
type PaymentAcknowledgementMessageDataType struct {
	PaymentAcknowlegementMessageID            string                               `xml:"urn:com.workday/bsvc Payment_Acknowlegement_Message_ID,omitempty"`
	PaymentAcknowledgementMessageDate         *time.Time                           `xml:"urn:com.workday/bsvc Payment_Acknowledgement_Message_Date,omitempty"`
	OriginatingBankID                         string                               `xml:"urn:com.workday/bsvc Originating_Bank_ID,omitempty"`
	PaymentAcknowledgementMessageUpdate       *bool                                `xml:"urn:com.workday/bsvc Payment_Acknowledgement_Message_Update,omitempty"`
	MarkRejectedPrenotePaymentAsCanceled      *bool                                `xml:"urn:com.workday/bsvc Mark_Rejected_Prenote_Payment_As_Canceled,omitempty"`
	PaymentFileAcknowledgementReplacementData []PaymentFileAcknowledgementDataType `xml:"urn:com.workday/bsvc Payment_File_Acknowledgement_Replacement_Data,omitempty"`
	PaymentAcknowledgementData                []PaymentAcknowledgementDataType     `xml:"urn:com.workday/bsvc Payment_Acknowledgement_Data,omitempty"`
}

func (t *PaymentAcknowledgementMessageDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PaymentAcknowledgementMessageDataType
	var layout struct {
		*T
		PaymentAcknowledgementMessageDate *xsdDateTime `xml:"urn:com.workday/bsvc Payment_Acknowledgement_Message_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PaymentAcknowledgementMessageDate = (*xsdDateTime)(layout.T.PaymentAcknowledgementMessageDate)
	return e.EncodeElement(layout, start)
}
func (t *PaymentAcknowledgementMessageDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PaymentAcknowledgementMessageDataType
	var overlay struct {
		*T
		PaymentAcknowledgementMessageDate *xsdDateTime `xml:"urn:com.workday/bsvc Payment_Acknowledgement_Message_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentAcknowledgementMessageDate = (*xsdDateTime)(overlay.T.PaymentAcknowledgementMessageDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PaymentAcknowledgementStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentAcknowledgementStatusObjectType struct {
	ID         []PaymentAcknowledgementStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PaymentCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentCategoryObjectType struct {
	ID         []PaymentCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Container element for each Payment Election
type PaymentElectionDataType struct {
	Order                        float64                        `xml:"urn:com.workday/bsvc Order"`
	PaymentElectionRuleReference *PaymentElectionRuleObjectType `xml:"urn:com.workday/bsvc Payment_Election_Rule_Reference"`
	CountryReference             *CountryObjectType             `xml:"urn:com.workday/bsvc Country_Reference"`
	CurrencyReference            *CurrencyObjectType            `xml:"urn:com.workday/bsvc Currency_Reference"`
	PaymentTypeReference         *PaymentTypeObjectType         `xml:"urn:com.workday/bsvc Payment_Type_Reference"`
	WorkerBankAccountData        *WorkerBankAccountWWSDataType  `xml:"urn:com.workday/bsvc Worker_Bank_Account_Data,omitempty"`
	DistributionAmount           float64                        `xml:"urn:com.workday/bsvc Distribution_Amount"`
	DistributionPercentage       float64                        `xml:"urn:com.workday/bsvc Distribution_Percentage"`
	DistributionBalance          bool                           `xml:"urn:com.workday/bsvc Distribution_Balance"`
}

// Contains a unique identifier for an instance of an object.
type PaymentElectionEnrollableObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentElectionEnrollableObjectType struct {
	ID         []PaymentElectionEnrollableObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payment Election Enrollment Request Criteria
type PaymentElectionEnrollmentRequestCriteriaType struct {
}

// Payment Election Enrollment Request Reference
type PaymentElectionEnrollmentRequestReferencesType struct {
	PaymentElectionEnrollmentReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Payment_Election_Enrollment_Reference"`
}

// Get Payment Election Enrollments Response
type PaymentElectionEnrollmentResponseDataType struct {
	PaymentElectionEnrollment []PaymentElectionEnrollmentType `xml:"urn:com.workday/bsvc Payment_Election_Enrollment,omitempty"`
}

// Payment Election Enrollment Response Group
type PaymentElectionEnrollmentResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Container element for Payment Election Enrollments
type PaymentElectionEnrollmentType struct {
	PaymentElectionEnrollmentReference *UniqueIdentifierObjectType                  `xml:"urn:com.workday/bsvc Payment_Election_Enrollment_Reference,omitempty"`
	PaymentElectionEnrollmentData      []PaymentElectionEnrollmentforWorkerDataType `xml:"urn:com.workday/bsvc Payment_Election_Enrollment_Data,omitempty"`
}

// Container element for Payment Election Enrollments
type PaymentElectionEnrollmentforWorkerDataType struct {
	RoleIDReference              *PaymentElectionEnrollableObjectType `xml:"urn:com.workday/bsvc Role_ID_Reference"`
	CountryReference             *CountryObjectType                   `xml:"urn:com.workday/bsvc Country_Reference"`
	CurrencyReference            *CurrencyObjectType                  `xml:"urn:com.workday/bsvc Currency_Reference"`
	PaymentElectionRuleGroupData []PaymentElectionRuleGroupDataType   `xml:"urn:com.workday/bsvc Payment_Election_Rule_Group_Data,omitempty"`
}

// Container element for Payment Election Options
type PaymentElectionOptionDataType struct {
	RoleReference                    *PaymentElectionEnrollableObjectType   `xml:"urn:com.workday/bsvc Role_Reference"`
	DefaultCountryReference          *CountryObjectType                     `xml:"urn:com.workday/bsvc Default_Country_Reference,omitempty"`
	DefaultCurrencyReference         *CurrencyObjectType                    `xml:"urn:com.workday/bsvc Default_Currency_Reference,omitempty"`
	PayTypePaymentElectionOptionData []PayTypePaymentElectionOptionDataType `xml:"urn:com.workday/bsvc Pay_Type_Payment_Election_Option_Data,omitempty"`
}

// Payment Election Option Request Criteria
type PaymentElectionOptionRequestCriteriaType struct {
}

// Payment Election Option Request Reference
type PaymentElectionOptionRequestReferencesType struct {
	PaymentElectionOptionReference []UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Payment_Election_Option_Reference"`
}

// Get Payment Election Options Response
type PaymentElectionOptionResponseDataType struct {
	PaymentElectionOption []PaymentElectionOptionType `xml:"urn:com.workday/bsvc Payment_Election_Option,omitempty"`
}

// Payment Election Option Response Group
type PaymentElectionOptionResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Get Payment Election Options Response
type PaymentElectionOptionType struct {
	PaymentElectionOptionReference *UniqueIdentifierObjectType     `xml:"urn:com.workday/bsvc Payment_Election_Option_Reference,omitempty"`
	PaymentElectionOptionData      []PaymentElectionOptionDataType `xml:"urn:com.workday/bsvc Payment_Election_Option_Data,omitempty"`
}

// Container element for Payment Election Rule
type PaymentElectionRuleGroupDataType struct {
	PaymentElectionRuleReference *PaymentElectionRuleObjectType `xml:"urn:com.workday/bsvc Payment_Election_Rule_Reference"`
	PaymentElectionData          []PaymentElectionDataType      `xml:"urn:com.workday/bsvc Payment_Election_Data"`
}

// Contains a unique identifier for an instance of an object.
type PaymentElectionRuleObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentElectionRuleObjectType struct {
	ID         []PaymentElectionRuleObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payment File Acknowledgement Data add or update payment file acknowledgements that allow the update of payment file status and error code information.
type PaymentFileAcknowledgementDataType struct {
	OriginalPaymentMessageID    string                                  `xml:"urn:com.workday/bsvc Original_Payment_Message_ID,omitempty"`
	PaymentStatusReference      *PaymentAcknowledgementStatusObjectType `xml:"urn:com.workday/bsvc Payment_Status_Reference,omitempty"`
	StatusReasonCode            string                                  `xml:"urn:com.workday/bsvc Status_Reason_Code,omitempty"`
	StatusReasonCodeDescription string                                  `xml:"urn:com.workday/bsvc Status_Reason_Code_Description,omitempty"`
	OriginalPaymentMessageDate  *time.Time                              `xml:"urn:com.workday/bsvc Original_Payment_Message_Date,omitempty"`
}

func (t *PaymentFileAcknowledgementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PaymentFileAcknowledgementDataType
	var layout struct {
		*T
		OriginalPaymentMessageDate *xsdDate `xml:"urn:com.workday/bsvc Original_Payment_Message_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.OriginalPaymentMessageDate = (*xsdDate)(layout.T.OriginalPaymentMessageDate)
	return e.EncodeElement(layout, start)
}
func (t *PaymentFileAcknowledgementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PaymentFileAcknowledgementDataType
	var overlay struct {
		*T
		OriginalPaymentMessageDate *xsdDate `xml:"urn:com.workday/bsvc Original_Payment_Message_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.OriginalPaymentMessageDate = (*xsdDate)(overlay.T.OriginalPaymentMessageDate)
	return d.DecodeElement(&overlay, &start)
}

// Common payment data that can be part identified in payment groups. Can be conditionally returned by setting the "Include Payment Group Data" grouping flag.
type PaymentGroupDataType struct {
	PaymentGroupReference    *PaymentGroupObjectType    `xml:"urn:com.workday/bsvc Payment_Group_Reference,omitempty"`
	PaymentDate              *time.Time                 `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	CurrencyReference        *CurrencyObjectType        `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	PaymentTypeReference     *PaymentTypeObjectType     `xml:"urn:com.workday/bsvc Payment_Type_Reference,omitempty"`
	PaymentMethodReference   *PaymentMethodObjectType   `xml:"urn:com.workday/bsvc Payment_Method_Reference,omitempty"`
	PaymentCategoryReference *PaymentCategoryObjectType `xml:"urn:com.workday/bsvc Payment_Category_Reference,omitempty"`
}

func (t *PaymentGroupDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PaymentGroupDataType
	var layout struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PaymentDate = (*xsdDate)(layout.T.PaymentDate)
	return e.EncodeElement(layout, start)
}
func (t *PaymentGroupDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PaymentGroupDataType
	var overlay struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDate = (*xsdDate)(overlay.T.PaymentDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PaymentGroupObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentGroupObjectType struct {
	ID         []PaymentGroupObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PaymentHandlingInstructionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentHandlingInstructionObjectType struct {
	ID         []PaymentHandlingInstructionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                   `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing payment integration data
type PaymentIntegrationNewDataType struct {
	PaymentIntegrationReference *PaymentIntegrationObjectType `xml:"urn:com.workday/bsvc Payment_Integration_Reference"`
	Order                       string                        `xml:"urn:com.workday/bsvc Order"`
}

// Contains a unique identifier for an instance of an object.
type PaymentIntegrationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentIntegrationObjectType struct {
	ID         []PaymentIntegrationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing a set of payment integrations data
type PaymentIntegrationsDataType struct {
	PaymentIntegrationData []PaymentIntegrationNewDataType `xml:"urn:com.workday/bsvc Payment_Integration_Data,omitempty"`
}

// Payment Message Search Criteria
type PaymentMessageCriteriaType struct {
	PaymentMessageReference         *PaymentMessageObjectType `xml:"urn:com.workday/bsvc Payment_Message_Reference"`
	PaymentMessageGroup             string                    `xml:"urn:com.workday/bsvc Payment_Message_Group,omitempty"`
	ExcludePaymentStatusesReference []PaymentStatusObjectType `xml:"urn:com.workday/bsvc Exclude_Payment_Statuses_Reference,omitempty"`
}

// Payment Message Group Data
type PaymentMessageGroupWWSDataType struct {
	PaymentDate                  *time.Time                     `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	CurrencyReference            *CurrencyObjectType            `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	PaymentTypeReference         *PaymentTypeObjectType         `xml:"urn:com.workday/bsvc Payment_Type_Reference,omitempty"`
	PaymentMethodReference       *PaymentMethodObjectType       `xml:"urn:com.workday/bsvc Payment_Method_Reference,omitempty"`
	PaymentCategoryReference     *PaymentCategoryObjectType     `xml:"urn:com.workday/bsvc Payment_Category_Reference,omitempty"`
	OriginatingContactData       []OriginatingPartyWWSDataType  `xml:"urn:com.workday/bsvc Originating_Contact_Data,omitempty"`
	OriginatingBankData          []OriginatingPartyBankDataType `xml:"urn:com.workday/bsvc Originating_Bank_Data,omitempty"`
	GroupPaymentCount            float64                        `xml:"urn:com.workday/bsvc Group_Payment_Count,omitempty"`
	GroupRejectedPaymentCount    float64                        `xml:"urn:com.workday/bsvc Group_Rejected_Payment_Count,omitempty"`
	GroupReprocessedPaymentCount float64                        `xml:"urn:com.workday/bsvc Group_Reprocessed_Payment_Count,omitempty"`
	GroupPaymentSum              float64                        `xml:"urn:com.workday/bsvc Group_Payment_Sum,omitempty"`
	GroupID                      string                         `xml:"urn:com.workday/bsvc Group_ID,omitempty"`
	PaymentData                  []PaymentWWSDataType           `xml:"urn:com.workday/bsvc Payment_Data,omitempty"`
}

func (t *PaymentMessageGroupWWSDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PaymentMessageGroupWWSDataType
	var layout struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PaymentDate = (*xsdDate)(layout.T.PaymentDate)
	return e.EncodeElement(layout, start)
}
func (t *PaymentMessageGroupWWSDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PaymentMessageGroupWWSDataType
	var overlay struct {
		*T
		PaymentDate *xsdDate `xml:"urn:com.workday/bsvc Payment_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PaymentDate = (*xsdDate)(overlay.T.PaymentDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type PaymentMessageObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentMessageObjectType struct {
	ID         []PaymentMessageObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Get Payment Messages Request Criteria
type PaymentMessageRequestCriteriaType struct {
	OrganizationReference           []OrganizationObjectType          `xml:"urn:com.workday/bsvc Organization_Reference,omitempty"`
	BankAccountReference            *BankAccountObjectType            `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	OutsourcedPaymentGroupReference *OutsourcedPaymentGroupObjectType `xml:"urn:com.workday/bsvc Outsourced_Payment_Group_Reference,omitempty"`
	CreatedAfter                    *time.Time                        `xml:"urn:com.workday/bsvc Created_After,omitempty"`
	CreatedBefore                   *time.Time                        `xml:"urn:com.workday/bsvc Created_Before,omitempty"`
}

func (t *PaymentMessageRequestCriteriaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PaymentMessageRequestCriteriaType
	var layout struct {
		*T
		CreatedAfter  *xsdDate `xml:"urn:com.workday/bsvc Created_After,omitempty"`
		CreatedBefore *xsdDate `xml:"urn:com.workday/bsvc Created_Before,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAfter = (*xsdDate)(layout.T.CreatedAfter)
	layout.CreatedBefore = (*xsdDate)(layout.T.CreatedBefore)
	return e.EncodeElement(layout, start)
}
func (t *PaymentMessageRequestCriteriaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PaymentMessageRequestCriteriaType
	var overlay struct {
		*T
		CreatedAfter  *xsdDate `xml:"urn:com.workday/bsvc Created_After,omitempty"`
		CreatedBefore *xsdDate `xml:"urn:com.workday/bsvc Created_Before,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAfter = (*xsdDate)(overlay.T.CreatedAfter)
	overlay.CreatedBefore = (*xsdDate)(overlay.T.CreatedBefore)
	return d.DecodeElement(&overlay, &start)
}

// Payment Message Request References
type PaymentMessageRequestReferencesType struct {
	PaymentMessageReference []PaymentMessageObjectType `xml:"urn:com.workday/bsvc Payment_Message_Reference"`
}

// Payment Message Response Data
type PaymentMessageResponseDataType struct {
	PaymentMessage []PaymentMessageWWSType `xml:"urn:com.workday/bsvc Payment_Message,omitempty"`
}

// Payment Messages Response Group
type PaymentMessageResponseGroupType struct {
	IncludeReference             *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludePaymentData           *bool `xml:"urn:com.workday/bsvc Include_Payment_Data,omitempty"`
	IncludePayrollRemittanceData *bool `xml:"urn:com.workday/bsvc Include_Payroll_Remittance_Data,omitempty"`
}

// Payment Message Data
type PaymentMessageWWSDataType struct {
	PaymentMessageID           string                              `xml:"urn:com.workday/bsvc Payment_Message_ID,omitempty"`
	CreateDate                 *time.Time                          `xml:"urn:com.workday/bsvc Create_Date,omitempty"`
	CreateDateSequence         float64                             `xml:"urn:com.workday/bsvc Create_Date_Sequence,omitempty"`
	IntegrationSystemReference *IntegrationSystemAuditedObjectType `xml:"urn:com.workday/bsvc Integration_System_Reference,omitempty"`
	PaymentCount               float64                             `xml:"urn:com.workday/bsvc Payment_Count,omitempty"`
	PaymentsSum                float64                             `xml:"urn:com.workday/bsvc Payments_Sum,omitempty"`
	PaymentMessageGroupData    []PaymentMessageGroupWWSDataType    `xml:"urn:com.workday/bsvc Payment_Message_Group_Data,omitempty"`
}

func (t *PaymentMessageWWSDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PaymentMessageWWSDataType
	var layout struct {
		*T
		CreateDate *xsdDateTime `xml:"urn:com.workday/bsvc Create_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreateDate = (*xsdDateTime)(layout.T.CreateDate)
	return e.EncodeElement(layout, start)
}
func (t *PaymentMessageWWSDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PaymentMessageWWSDataType
	var overlay struct {
		*T
		CreateDate *xsdDateTime `xml:"urn:com.workday/bsvc Create_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreateDate = (*xsdDateTime)(overlay.T.CreateDate)
	return d.DecodeElement(&overlay, &start)
}

// Payment Message WWS
type PaymentMessageWWSType struct {
	PaymentMessageReference *PaymentMessageObjectType   `xml:"urn:com.workday/bsvc Payment_Message_Reference,omitempty"`
	PaymentMessageData      []PaymentMessageWWSDataType `xml:"urn:com.workday/bsvc Payment_Message_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PaymentMethodObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentMethodObjectType struct {
	ID         []PaymentMethodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PaymentStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentStatusObjectType struct {
	ID         []PaymentStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PaymentTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PaymentTypeObjectType struct {
	ID         []PaymentTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                    `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payment Data
type PaymentWWSDataType struct {
	PaymentGroupData             *PaymentGroupDataType                 `xml:"urn:com.workday/bsvc Payment_Group_Data,omitempty"`
	SettlementRunReference       *UniqueIdentifierObjectType           `xml:"urn:com.workday/bsvc Settlement_Run_Reference,omitempty"`
	PaymentStatus                string                                `xml:"urn:com.workday/bsvc Payment_Status,omitempty"`
	PaymentMemo                  string                                `xml:"urn:com.workday/bsvc Payment_Memo,omitempty"`
	AdditionalTypeReference      *AdditionalReferenceTypeObjectType    `xml:"urn:com.workday/bsvc Additional_Type_Reference,omitempty"`
	PaymentReferenceNumber       string                                `xml:"urn:com.workday/bsvc Payment_Reference_Number,omitempty"`
	PaymentAmount                float64                               `xml:"urn:com.workday/bsvc Payment_Amount,omitempty"`
	AmountasText                 string                                `xml:"urn:com.workday/bsvc Amount_as_Text,omitempty"`
	DiscountTaken                float64                               `xml:"urn:com.workday/bsvc Discount_Taken,omitempty"`
	CurrencyReference            *CurrencyObjectType                   `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	CurrencySymbol               string                                `xml:"urn:com.workday/bsvc Currency_Symbol,omitempty"`
	CurrencyDescription          string                                `xml:"urn:com.workday/bsvc Currency_Description,omitempty"`
	CheckorAdviceNumber          float64                               `xml:"urn:com.workday/bsvc Check_or_Advice_Number,omitempty"`
	CheckNumber                  string                                `xml:"urn:com.workday/bsvc Check_Number,omitempty"`
	PrenoteFlag                  *bool                                 `xml:"urn:com.workday/bsvc Prenote_Flag,omitempty"`
	PayrollPaymentDisplayOrder   float64                               `xml:"urn:com.workday/bsvc Payroll_Payment_Display_Order,omitempty"`
	HandlingCodeReference        *PaymentHandlingInstructionObjectType `xml:"urn:com.workday/bsvc Handling_Code_Reference,omitempty"`
	TaxCodeReference             *TaxCodeObjectType                    `xml:"urn:com.workday/bsvc Tax_Code_Reference,omitempty"`
	CompanyReference             *CompanyObjectType                    `xml:"urn:com.workday/bsvc Company_Reference,omitempty"`
	OnBehalfofCompanyReference   []CompanyObjectType                   `xml:"urn:com.workday/bsvc On_Behalf_of_Company_Reference,omitempty"`
	MandateData                  *MandateDataType                      `xml:"urn:com.workday/bsvc Mandate_Data,omitempty"`
	OriginatingContactData       []OriginatingPartyWWSDataType         `xml:"urn:com.workday/bsvc Originating_Contact_Data,omitempty"`
	OriginatingBankData          []OriginatingPartyBankDataType        `xml:"urn:com.workday/bsvc Originating_Bank_Data,omitempty"`
	OnBehalfOfCompanyData        []OnBehalfOfCompanyWWSDataType        `xml:"urn:com.workday/bsvc On_Behalf_Of_Company_Data,omitempty"`
	ReceivingPartyReference      *PayeeObjectType                      `xml:"urn:com.workday/bsvc Receiving_Party_Reference,omitempty"`
	ReceivingPartyContactData    []ReceivingPartyWWSDataType           `xml:"urn:com.workday/bsvc Receiving_Party_Contact_Data,omitempty"`
	ReceivingPartyBankData       []ReceivingPartyBankDataType          `xml:"urn:com.workday/bsvc Receiving_Party_Bank_Data,omitempty"`
	ReceivingPartyCreditCardData *ReceivingPartyCreditCardDataType     `xml:"urn:com.workday/bsvc Receiving_Party_Credit_Card_Data,omitempty"`
	RemittanceData               []RemittanceDataType                  `xml:"urn:com.workday/bsvc Remittance_Data,omitempty"`
	PayrollRemittanceData        []PayrollRemittanceDataType           `xml:"urn:com.workday/bsvc Payroll_Remittance_Data,omitempty"`
	PayrollResultGrossAmountData []PayrollResultGrossAmountDataType    `xml:"urn:com.workday/bsvc Payroll_Result_Gross_Amount_Data,omitempty"`
	PayrollResultCountryData     []PayrollResultCountryDataType        `xml:"urn:com.workday/bsvc Payroll_Result_Country_Data,omitempty"`
	CashOutAddressData           *CashOutAddressDataType               `xml:"urn:com.workday/bsvc Cash_Out_Address_Data,omitempty"`
	IntegrationFieldOverrideData []DocumentFieldResultDataType         `xml:"urn:com.workday/bsvc Integration_Field_Override_Data,omitempty"`
}

// Get Payments Request Criteria
type PaymentsRequestCriteriaType struct {
	GeneralPaymentCriteria         []GeneralPaymentCriteriaType       `xml:"urn:com.workday/bsvc General_Payment_Criteria,omitempty"`
	CheckPaymentsToPrintCriteria   []CheckPaymentsToPrintCriteriaType `xml:"urn:com.workday/bsvc Check_Payments_To_Print_Criteria,omitempty"`
	PaymentMessageCriteria         []PaymentMessageCriteriaType       `xml:"urn:com.workday/bsvc Payment_Message_Criteria,omitempty"`
	RemittanceFileCriteria         []RemittanceFileCriteriaType       `xml:"urn:com.workday/bsvc Remittance_File_Criteria,omitempty"`
	FieldAndParamenterCriteriaData *FieldAndParameterCriteriaDataType `xml:"urn:com.workday/bsvc Field_And_Paramenter_Criteria_Data,omitempty"`
}

// Payment Request References
type PaymentsRequestReferencesType struct {
	PaymentWWSReference []ReportingTransactionObjectType `xml:"urn:com.workday/bsvc Payment_WWS_Reference"`
}

// Payments Response Data
type PaymentsResponseDataType struct {
	Payment []PaymentsType `xml:"urn:com.workday/bsvc Payment,omitempty"`
}

// Payments Response Group
type PaymentsResponseGroupType struct {
	IncludeReference                *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
	IncludeOriginatingBankData      *bool `xml:"urn:com.workday/bsvc Include_Originating_Bank_Data,omitempty"`
	IncludePayrollRemittanceData    *bool `xml:"urn:com.workday/bsvc Include_Payroll_Remittance_Data,omitempty"`
	IncludePaymentGroupData         *bool `xml:"urn:com.workday/bsvc Include_Payment_Group_Data,omitempty"`
	IncludePayrollResultCountryData *bool `xml:"urn:com.workday/bsvc Include_Payroll_Result_Country_Data,omitempty"`
}

// Payment
type PaymentsType struct {
	PaymentReference *ReportingTransactionObjectType `xml:"urn:com.workday/bsvc Payment_Reference,omitempty"`
	PaymentData      []PaymentWWSDataType            `xml:"urn:com.workday/bsvc Payment_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollCalculationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollCalculationObjectType struct {
	ID         []PayrollCalculationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Payroll Remittance Absence Plan
type PayrollRemittanceAbsencePlanType struct {
	GBName  string  `xml:"urn:com.workday/bsvc GB_Name,omitempty"`
	Name    string  `xml:"urn:com.workday/bsvc Name,omitempty"`
	Accrued float64 `xml:"urn:com.workday/bsvc Accrued,omitempty"`
	Reduced float64 `xml:"urn:com.workday/bsvc Reduced,omitempty"`
	Balance float64 `xml:"urn:com.workday/bsvc Balance,omitempty"`
}

// Payroll Remittance Data. Will always be included in the Get Payroll Payslips web service.  Can be conditionally returned by setting the "Include Payroll Remittance Data" grouping flag in the Get Payments web service.
type PayrollRemittanceDataType struct {
	PeriodStartDate                      *time.Time                                      `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
	PeriodEndDate                        *time.Time                                      `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	FederalMaritalStatusReference        *PayrollWithholdingStatusObjectType             `xml:"urn:com.workday/bsvc Federal_Marital_Status_Reference,omitempty"`
	FederalAllowance                     float64                                         `xml:"urn:com.workday/bsvc Federal_Allowance,omitempty"`
	FederalAdditionalAllowance           float64                                         `xml:"urn:com.workday/bsvc Federal_Additional_Allowance,omitempty"`
	WorkStateMaritalStatusReference      *PayrollWithholdingStatusObjectType             `xml:"urn:com.workday/bsvc Work_State_Marital_Status_Reference,omitempty"`
	WorkStateAllowance                   float64                                         `xml:"urn:com.workday/bsvc Work_State_Allowance,omitempty"`
	WorkStateAdditionalAmount            float64                                         `xml:"urn:com.workday/bsvc Work_State_Additional_Amount,omitempty"`
	CalculatedFederalTotalClaimAmount    float64                                         `xml:"urn:com.workday/bsvc Calculated_Federal_Total_Claim_Amount,omitempty"`
	CalculatedProvincialTotalClaimAmount float64                                         `xml:"urn:com.workday/bsvc Calculated_Provincial_Total_Claim_Amount,omitempty"`
	FederalTotalClaimAmount              float64                                         `xml:"urn:com.workday/bsvc Federal_Total_Claim_Amount,omitempty"`
	ProvinceTotalClaimAmount             float64                                         `xml:"urn:com.workday/bsvc Province_Total_Claim_Amount,omitempty"`
	FederalLivingPrescribedZone          float64                                         `xml:"urn:com.workday/bsvc Federal_Living_Prescribed_Zone,omitempty"`
	ProvinceLivingPrescribedZone         float64                                         `xml:"urn:com.workday/bsvc Province_Living_Prescribed_Zone,omitempty"`
	FederalAdditionalTaxAmounts          float64                                         `xml:"urn:com.workday/bsvc Federal_Additional_Tax_Amounts,omitempty"`
	ProvinceAdditionalTaxAmounts         float64                                         `xml:"urn:com.workday/bsvc Province_Additional_Tax_Amounts,omitempty"`
	FederalAnnualDeductionsCredits       float64                                         `xml:"urn:com.workday/bsvc Federal_Annual_Deductions_Credits,omitempty"`
	ProvinceAnnualDeductionsCredits      float64                                         `xml:"urn:com.workday/bsvc Province_Annual_Deductions_Credits,omitempty"`
	FederalLabourSponsoredFund           float64                                         `xml:"urn:com.workday/bsvc Federal_Labour_Sponsored_Fund,omitempty"`
	ProvinceLabourSponsoredFund          float64                                         `xml:"urn:com.workday/bsvc Province_Labour_Sponsored_Fund,omitempty"`
	PayGroupDetailReference              *PayGroupDetailObjectType                       `xml:"urn:com.workday/bsvc Pay_Group_Detail_Reference,omitempty"`
	FrequencyName                        string                                          `xml:"urn:com.workday/bsvc Frequency_Name,omitempty"`
	GrossAmount                          float64                                         `xml:"urn:com.workday/bsvc Gross_Amount,omitempty"`
	TotalUnits                           float64                                         `xml:"urn:com.workday/bsvc Total_Units,omitempty"`
	TotalHoursWorked                     float64                                         `xml:"urn:com.workday/bsvc Total_Hours_Worked,omitempty"`
	YTDTotalHoursWorked                  float64                                         `xml:"urn:com.workday/bsvc YTD_Total_Hours_Worked,omitempty"`
	NetPay                               float64                                         `xml:"urn:com.workday/bsvc Net_Pay,omitempty"`
	TaxesDeductions                      float64                                         `xml:"urn:com.workday/bsvc Taxes_Deductions,omitempty"`
	PrintPayslip                         *bool                                           `xml:"urn:com.workday/bsvc Print_Payslip,omitempty"`
	PrintCheckonPayslip                  *bool                                           `xml:"urn:com.workday/bsvc Print_Check_on_Payslip,omitempty"`
	PayslipDistributionSortingCriteria   string                                          `xml:"urn:com.workday/bsvc Payslip_Distribution_Sorting_Criteria,omitempty"`
	OregonBusinessIdentificationNumber   string                                          `xml:"urn:com.workday/bsvc Oregon_Business_Identification_Number,omitempty"`
	PayRateTypeReference                 *PayRateTypeObjectType                          `xml:"urn:com.workday/bsvc Pay_Rate_Type_Reference,omitempty"`
	PayRate                              float64                                         `xml:"urn:com.workday/bsvc Pay_Rate,omitempty"`
	WorkerData                           []PayrollRemittanceWorkerDataType               `xml:"urn:com.workday/bsvc Worker_Data,omitempty"`
	WorkAddressData                      []AddressInformationDataType                    `xml:"urn:com.workday/bsvc Work_Address_Data,omitempty"`
	GrossAndNetData                      []PayrollRemittanceGrossAndNetDataType          `xml:"urn:com.workday/bsvc Gross_And_Net_Data,omitempty"`
	TotalData                            []PayrollRemittanceTotalDataType                `xml:"urn:com.workday/bsvc Total_Data,omitempty"`
	EarningsData                         []PayrollRemittanceEarningsDataType             `xml:"urn:com.workday/bsvc Earnings_Data,omitempty"`
	PreTaxDeductionsData                 []PayrollRemittancePreTaxDeductionDataType      `xml:"urn:com.workday/bsvc Pre_Tax_Deductions_Data,omitempty"`
	PostTaxDeductionsData                []PayrollRemittancePostTaxDataType              `xml:"urn:com.workday/bsvc Post_Tax_Deductions_Data,omitempty"`
	EmployeeTaxesData                    []PayrollRemittanceEmployeeTaxesDataType        `xml:"urn:com.workday/bsvc Employee_Taxes_Data,omitempty"`
	EmployerPaidBenefitsData             []PayrollRemittanceEmployerPaidBenefitsDataType `xml:"urn:com.workday/bsvc Employer_Paid_Benefits_Data,omitempty"`
	TaxableWagesData                     []PayrollRemittanceTaxableWagesDataType         `xml:"urn:com.workday/bsvc Taxable_Wages_Data,omitempty"`
	AbsencePlansData                     []PayrollRemittanceAbsencePlanType              `xml:"urn:com.workday/bsvc Absence_Plans_Data,omitempty"`
	PayslipMessageData                   []PayrollRemittancePayslipMessageDataType       `xml:"urn:com.workday/bsvc Payslip_Message_Data,omitempty"`
}

func (t *PayrollRemittanceDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollRemittanceDataType
	var layout struct {
		*T
		PeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
		PeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PeriodStartDate = (*xsdDate)(layout.T.PeriodStartDate)
	layout.PeriodEndDate = (*xsdDate)(layout.T.PeriodEndDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollRemittanceDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollRemittanceDataType
	var overlay struct {
		*T
		PeriodStartDate *xsdDate `xml:"urn:com.workday/bsvc Period_Start_Date,omitempty"`
		PeriodEndDate   *xsdDate `xml:"urn:com.workday/bsvc Period_End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PeriodStartDate = (*xsdDate)(overlay.T.PeriodStartDate)
	overlay.PeriodEndDate = (*xsdDate)(overlay.T.PeriodEndDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll remitance Earnings
type PayrollRemittanceEarningsDataType struct {
	PayComponentReference    []PayComponentReferenceType    `xml:"urn:com.workday/bsvc Pay_Component_Reference,omitempty"`
	GBName                   string                         `xml:"urn:com.workday/bsvc GB_Name,omitempty"`
	PayComponent             string                         `xml:"urn:com.workday/bsvc Pay_Component,omitempty"`
	StartDate                *time.Time                     `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                  *time.Time                     `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	PayrollWorktagsReference []PayrollWorktagObjectType     `xml:"urn:com.workday/bsvc Payroll_Worktags_Reference,omitempty"`
	Amount                   float64                        `xml:"urn:com.workday/bsvc Amount,omitempty"`
	YTDAmount                float64                        `xml:"urn:com.workday/bsvc YTD_Amount,omitempty"`
	UnitsTypeReference       []PayrollCalculationObjectType `xml:"urn:com.workday/bsvc Units_Type_Reference,omitempty"`
	UnitsValue               float64                        `xml:"urn:com.workday/bsvc Units_Value,omitempty"`
	YTDUnitsValue            float64                        `xml:"urn:com.workday/bsvc YTD_Units_Value,omitempty"`
	RateTypeReference        []PayrollCalculationObjectType `xml:"urn:com.workday/bsvc Rate_Type_Reference,omitempty"`
	Rate                     float64                        `xml:"urn:com.workday/bsvc Rate,omitempty"`
}

func (t *PayrollRemittanceEarningsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollRemittanceEarningsDataType
	var layout struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollRemittanceEarningsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollRemittanceEarningsDataType
	var overlay struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Remittance Employee Taxes Data
type PayrollRemittanceEmployeeTaxesDataType struct {
	GBName                   string                      `xml:"urn:com.workday/bsvc GB_Name,omitempty"`
	PayComponentReference    []PayComponentReferenceType `xml:"urn:com.workday/bsvc Pay_Component_Reference,omitempty"`
	PayComponent             string                      `xml:"urn:com.workday/bsvc Pay_Component,omitempty"`
	StartDate                *time.Time                  `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                  *time.Time                  `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	PayrollWorktagsReference []PayrollWorktagObjectType  `xml:"urn:com.workday/bsvc Payroll_Worktags_Reference,omitempty"`
	Amount                   float64                     `xml:"urn:com.workday/bsvc Amount,omitempty"`
	YTD                      float64                     `xml:"urn:com.workday/bsvc YTD,omitempty"`
}

func (t *PayrollRemittanceEmployeeTaxesDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollRemittanceEmployeeTaxesDataType
	var layout struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollRemittanceEmployeeTaxesDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollRemittanceEmployeeTaxesDataType
	var overlay struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Employer Paid Benefits
type PayrollRemittanceEmployerPaidBenefitsDataType struct {
	PayComponentReference    []PayComponentReferenceType `xml:"urn:com.workday/bsvc Pay_Component_Reference,omitempty"`
	GBName                   string                      `xml:"urn:com.workday/bsvc GB_Name,omitempty"`
	PayComponent             string                      `xml:"urn:com.workday/bsvc Pay_Component,omitempty"`
	StartDate                *time.Time                  `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                  *time.Time                  `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	PayrollWorktagsReference []PayrollWorktagObjectType  `xml:"urn:com.workday/bsvc Payroll_Worktags_Reference,omitempty"`
	Amount                   float64                     `xml:"urn:com.workday/bsvc Amount,omitempty"`
	YTD                      float64                     `xml:"urn:com.workday/bsvc YTD,omitempty"`
}

func (t *PayrollRemittanceEmployerPaidBenefitsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollRemittanceEmployerPaidBenefitsDataType
	var layout struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollRemittanceEmployerPaidBenefitsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollRemittanceEmployerPaidBenefitsDataType
	var overlay struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Remittance Gross And Net Data
type PayrollRemittanceGrossAndNetDataType struct {
	TypeReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Type_Reference,omitempty"`
	Amount        float64                     `xml:"urn:com.workday/bsvc Amount,omitempty"`
	YTD           float64                     `xml:"urn:com.workday/bsvc YTD,omitempty"`
}

// Payroll Remittance Payslip Message
type PayrollRemittancePayslipMessageDataType struct {
	PayslipMessage string `xml:"urn:com.workday/bsvc Payslip_Message,omitempty"`
}

// Payroll Remittance Post Tax Data
type PayrollRemittancePostTaxDataType struct {
	PayComponentReference    []PayComponentReferenceType `xml:"urn:com.workday/bsvc Pay_Component_Reference,omitempty"`
	GBName                   string                      `xml:"urn:com.workday/bsvc GB_Name,omitempty"`
	PayComponent             string                      `xml:"urn:com.workday/bsvc Pay_Component,omitempty"`
	StartDate                *time.Time                  `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                  *time.Time                  `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	PayrollWorktagsReference []PayrollWorktagObjectType  `xml:"urn:com.workday/bsvc Payroll_Worktags_Reference,omitempty"`
	Amount                   float64                     `xml:"urn:com.workday/bsvc Amount,omitempty"`
	YTD                      float64                     `xml:"urn:com.workday/bsvc YTD,omitempty"`
}

func (t *PayrollRemittancePostTaxDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollRemittancePostTaxDataType
	var layout struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollRemittancePostTaxDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollRemittancePostTaxDataType
	var overlay struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Remittance Pre Tax Deduction Data
type PayrollRemittancePreTaxDeductionDataType struct {
	PayComponentReference    []PayComponentReferenceType `xml:"urn:com.workday/bsvc Pay_Component_Reference,omitempty"`
	GBName                   string                      `xml:"urn:com.workday/bsvc GB_Name,omitempty"`
	PayComponent             string                      `xml:"urn:com.workday/bsvc Pay_Component,omitempty"`
	StartDate                *time.Time                  `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
	EndDate                  *time.Time                  `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	PayrollWorktagsReference []PayrollWorktagObjectType  `xml:"urn:com.workday/bsvc Payroll_Worktags_Reference,omitempty"`
	Amount                   float64                     `xml:"urn:com.workday/bsvc Amount,omitempty"`
	YTD                      float64                     `xml:"urn:com.workday/bsvc YTD,omitempty"`
}

func (t *PayrollRemittancePreTaxDeductionDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollRemittancePreTaxDeductionDataType
	var layout struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartDate = (*xsdDate)(layout.T.StartDate)
	layout.EndDate = (*xsdDate)(layout.T.EndDate)
	return e.EncodeElement(layout, start)
}
func (t *PayrollRemittancePreTaxDeductionDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollRemittancePreTaxDeductionDataType
	var overlay struct {
		*T
		StartDate *xsdDate `xml:"urn:com.workday/bsvc Start_Date,omitempty"`
		EndDate   *xsdDate `xml:"urn:com.workday/bsvc End_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartDate = (*xsdDate)(overlay.T.StartDate)
	overlay.EndDate = (*xsdDate)(overlay.T.EndDate)
	return d.DecodeElement(&overlay, &start)
}

// Payroll Remittance Taxable Wages
type PayrollRemittanceTaxableWagesDataType struct {
	GBName string  `xml:"urn:com.workday/bsvc GB_Name,omitempty"`
	Name   string  `xml:"urn:com.workday/bsvc Name,omitempty"`
	Amount float64 `xml:"urn:com.workday/bsvc Amount,omitempty"`
	YTD    float64 `xml:"urn:com.workday/bsvc YTD,omitempty"`
}

// Payroll Remittance Total Data
type PayrollRemittanceTotalDataType struct {
	Label         string  `xml:"urn:com.workday/bsvc Label,omitempty"`
	CurrentPeriod float64 `xml:"urn:com.workday/bsvc Current_Period,omitempty"`
	YTD           float64 `xml:"urn:com.workday/bsvc YTD,omitempty"`
}

// Payroll Remittance Data
type PayrollRemittanceWorkerDataType struct {
	EmployeeID                string                             `xml:"urn:com.workday/bsvc Employee_ID,omitempty"`
	EmployeeName              string                             `xml:"urn:com.workday/bsvc Employee_Name,omitempty"`
	PreferredName             string                             `xml:"urn:com.workday/bsvc Preferred_Name,omitempty"`
	LegalName                 string                             `xml:"urn:com.workday/bsvc Legal_Name,omitempty"`
	LegalFirstName            string                             `xml:"urn:com.workday/bsvc Legal_First_Name,omitempty"`
	LegalLastName             string                             `xml:"urn:com.workday/bsvc Legal_Last_Name,omitempty"`
	LegalMiddleName           string                             `xml:"urn:com.workday/bsvc Legal_Middle_Name,omitempty"`
	DateofBirth               *time.Time                         `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	SSN                       string                             `xml:"urn:com.workday/bsvc SSN,omitempty"`
	Company                   string                             `xml:"urn:com.workday/bsvc Company,omitempty"`
	Occupation                string                             `xml:"urn:com.workday/bsvc Occupation,omitempty"`
	Phone                     string                             `xml:"urn:com.workday/bsvc Phone,omitempty"`
	LocationReference         *LocationObjectType                `xml:"urn:com.workday/bsvc Location_Reference,omitempty"`
	DepartmentNameReference   *SupervisoryOrganizationObjectType `xml:"urn:com.workday/bsvc Department_Name_Reference,omitempty"`
	SupervisoryOrganizationID string                             `xml:"urn:com.workday/bsvc Supervisory_Organization_ID,omitempty"`
	RegionReference           []RegionObjectType                 `xml:"urn:com.workday/bsvc Region_Reference,omitempty"`
	RegionID                  string                             `xml:"urn:com.workday/bsvc Region_ID,omitempty"`
	CostCenterReference       []CostCenterObjectType             `xml:"urn:com.workday/bsvc Cost_Center_Reference,omitempty"`
	CostCenterID              string                             `xml:"urn:com.workday/bsvc Cost_Center_ID,omitempty"`
	ManagerReference          *WorkerObjectType                  `xml:"urn:com.workday/bsvc Manager_Reference,omitempty"`
	HomeAddressData           []AddressInformationDataType       `xml:"urn:com.workday/bsvc Home_Address_Data,omitempty"`
}

func (t *PayrollRemittanceWorkerDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PayrollRemittanceWorkerDataType
	var layout struct {
		*T
		DateofBirth *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	}
	layout.T = (*T)(t)
	layout.DateofBirth = (*xsdDate)(layout.T.DateofBirth)
	return e.EncodeElement(layout, start)
}
func (t *PayrollRemittanceWorkerDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PayrollRemittanceWorkerDataType
	var overlay struct {
		*T
		DateofBirth *xsdDate `xml:"urn:com.workday/bsvc Date_of_Birth,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.DateofBirth = (*xsdDate)(overlay.T.DateofBirth)
	return d.DecodeElement(&overlay, &start)
}

// Specific Country Data for the Payroll Result
type PayrollResultCountryDataType struct {
	CountryReference    []CountryObjectType       `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	PayrollResultUKData []PayrollResultUKDataType `xml:"urn:com.workday/bsvc Payroll_Result_UK_Data,omitempty"`
}

// Payroll Result Gross Amount Data. Will return Gross Amount for Payroll Results in the Get Payments web service. Only will return a value if the payment is for a Payroll Result and if "Include_Payroll_Remittance_Data" is false.
type PayrollResultGrossAmountDataType struct {
	PayrollResultGrossAmount float64 `xml:"urn:com.workday/bsvc Payroll_Result_Gross_Amount,omitempty"`
}

// UK Specific Data for Payroll Result
type PayrollResultUKDataType struct {
	RTISubReference string `xml:"urn:com.workday/bsvc RTI_Sub-Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollWithholdingStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PayrollWithholdingStatusObjectType struct {
	ID         []PayrollWithholdingStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PayrollWorktagObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type PayrollWorktagObjectType struct {
	ID         []PayrollWorktagObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PeriodObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PeriodObjectType struct {
	ID         []PeriodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the components of a name, such as the First Name and Last Name.
type PersonNameDetailDataType struct {
	CountryReference                *CountryObjectType             `xml:"urn:com.workday/bsvc Country_Reference"`
	PrefixData                      *PersonNamePrefixDataType      `xml:"urn:com.workday/bsvc Prefix_Data,omitempty"`
	FirstName                       string                         `xml:"urn:com.workday/bsvc First_Name,omitempty"`
	MiddleName                      string                         `xml:"urn:com.workday/bsvc Middle_Name,omitempty"`
	LastName                        string                         `xml:"urn:com.workday/bsvc Last_Name,omitempty"`
	SecondaryLastName               string                         `xml:"urn:com.workday/bsvc Secondary_Last_Name,omitempty"`
	TertiaryLastName                string                         `xml:"urn:com.workday/bsvc Tertiary_Last_Name,omitempty"`
	LocalNameDetailData             *LocalPersonNameDetailDataType `xml:"urn:com.workday/bsvc Local_Name_Detail_Data,omitempty"`
	SuffixData                      *PersonNameSuffixDataType      `xml:"urn:com.workday/bsvc Suffix_Data,omitempty"`
	FullNameforSingaporeandMalaysia string                         `xml:"urn:com.workday/bsvc Full_Name_for_Singapore_and_Malaysia,omitempty"`
	FormattedName                   string                         `xml:"urn:com.workday/bsvc Formatted_Name,attr,omitempty"`
	ReportingName                   string                         `xml:"urn:com.workday/bsvc Reporting_Name,attr,omitempty"`
}

// Contains the prefixes for a name.
type PersonNamePrefixDataType struct {
	TitleReference      *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Title_Reference,omitempty"`
	TitleDescriptor     string                                               `xml:"urn:com.workday/bsvc Title_Descriptor,omitempty"`
	SalutationReference *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Salutation_Reference,omitempty"`
}

// Contains the suffixes for a name.
type PersonNameSuffixDataType struct {
	SocialSuffixReference       *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Social_Suffix_Reference,omitempty"`
	SocialSuffixDescriptor      string                                               `xml:"urn:com.workday/bsvc Social_Suffix_Descriptor,omitempty"`
	AcademicSuffixReference     *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Academic_Suffix_Reference,omitempty"`
	HereditarySuffixReference   *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Hereditary_Suffix_Reference,omitempty"`
	HonorarySuffixReference     *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Honorary_Suffix_Reference,omitempty"`
	ProfessionalSuffixReference *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Professional_Suffix_Reference,omitempty"`
	ReligiousSuffixReference    *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Religious_Suffix_Reference,omitempty"`
	RoyalSuffixReference        *CountryPredefinedPersonNameComponentValueObjectType `xml:"urn:com.workday/bsvc Royal_Suffix_Reference,omitempty"`
}

// Wrapper Element that includes Petty Cash Account Instance and Data
type PettyCashAccountDataType struct {
	PettyCashAccountID                      string                                 `xml:"urn:com.workday/bsvc Petty_Cash_Account_ID,omitempty"`
	AccountName                             string                                 `xml:"urn:com.workday/bsvc Account_Name"`
	FinancialPartyReference                 *FinancialPartyObjectType              `xml:"urn:com.workday/bsvc Financial_Party_Reference"`
	DefaultCurrencyReference                *CurrencyObjectType                    `xml:"urn:com.workday/bsvc Default_Currency_Reference,omitempty"`
	AcceptedCurrenciesReference             []CurrencyObjectType                   `xml:"urn:com.workday/bsvc Accepted_Currencies_Reference,omitempty"`
	PaymentTypeReference                    []PaymentTypeObjectType                `xml:"urn:com.workday/bsvc Payment_Type_Reference"`
	UsedbyCash                              *bool                                  `xml:"urn:com.workday/bsvc Used_by_Cash,omitempty"`
	UsedbyCustomerPayments                  *bool                                  `xml:"urn:com.workday/bsvc Used_by_Customer_Payments,omitempty"`
	UsedbyExpensePayments                   *bool                                  `xml:"urn:com.workday/bsvc Used_by_Expense_Payments,omitempty"`
	UsedbyPayroll                           *bool                                  `xml:"urn:com.workday/bsvc Used_by_Payroll,omitempty"`
	UsedbySupplierPayments                  *bool                                  `xml:"urn:com.workday/bsvc Used_by_Supplier_Payments,omitempty"`
	UsedbyIntercompanyPayments              *bool                                  `xml:"urn:com.workday/bsvc Used_by_Intercompany_Payments,omitempty"`
	UsedbyAdhocPayments                     *bool                                  `xml:"urn:com.workday/bsvc Used_by_Ad_hoc_Payments,omitempty"`
	UsedbyBankAccountTransfersforSettlement *bool                                  `xml:"urn:com.workday/bsvc Used_by_Bank_Account_Transfers_for_Settlement,omitempty"`
	UsedbyOffCyclePayroll                   *bool                                  `xml:"urn:com.workday/bsvc Used_by_Off-Cycle_Payroll,omitempty"`
	UsedbyPrenotePayments                   *bool                                  `xml:"urn:com.workday/bsvc Used_by_Prenote_Payments,omitempty"`
	AllowAdditionalusage                    *bool                                  `xml:"urn:com.workday/bsvc Allow_Additional_usage,omitempty"`
	BankAccountSecuritySegmentReference     []BankAccountSecuritySegmentObjectType `xml:"urn:com.workday/bsvc Bank_Account_Security_Segment_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PettyCashAccountObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PettyCashAccountObjectType struct {
	ID         []PettyCashAccountObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing Petty Cash request criteria
type PettyCashAccountRequestCriteriaType struct {
}

// Element containing reference instances for a Petty Cash Account
type PettyCashAccountRequestReferencesType struct {
	PettyCashAccountReference []PettyCashAccountObjectType `xml:"urn:com.workday/bsvc Petty_Cash_Account_Reference"`
}

// Wrapper Element that includes Petty Cash Account Reference Instance
type PettyCashAccountResponseDataType struct {
	PettyCashAccount []PettyCashAccountType `xml:"urn:com.workday/bsvc Petty_Cash_Account,omitempty"`
}

// Element containing Petty Cash Account response grouping options
type PettyCashAccountResponseGroupType struct {
	IncludeReference *bool `xml:"urn:com.workday/bsvc Include_Reference,omitempty"`
}

// Element containing Petty Cash Account reference and data
type PettyCashAccountType struct {
	PettyCashAccountReference *PettyCashAccountObjectType `xml:"urn:com.workday/bsvc Petty_Cash_Account_Reference,omitempty"`
	PettyCashAccountData      []PettyCashAccountDataType  `xml:"urn:com.workday/bsvc Petty_Cash_Account_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PhoneDeviceTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PhoneDeviceTypeObjectType struct {
	ID         []PhoneDeviceTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Phone Information
type PhoneInformationDataType struct {
	CountryISOCode              string                                        `xml:"urn:com.workday/bsvc Country_ISO_Code,omitempty"`
	InternationalPhoneCode      string                                        `xml:"urn:com.workday/bsvc International_Phone_Code,omitempty"`
	PhoneNumber                 string                                        `xml:"urn:com.workday/bsvc Phone_Number,omitempty"`
	PhoneExtension              string                                        `xml:"urn:com.workday/bsvc Phone_Extension,omitempty"`
	PhoneDeviceTypeReference    *PhoneDeviceTypeObjectType                    `xml:"urn:com.workday/bsvc Phone_Device_Type_Reference,omitempty"`
	UsageData                   []CommunicationMethodUsageInformationDataType `xml:"urn:com.workday/bsvc Usage_Data,omitempty"`
	PhoneReference              *PhoneReferenceObjectType                     `xml:"urn:com.workday/bsvc Phone_Reference,omitempty"`
	ID                          string                                        `xml:"urn:com.workday/bsvc ID,omitempty"`
	AreaCode                    string                                        `xml:"urn:com.workday/bsvc Area_Code,attr,omitempty"`
	TenantFormattedPhone        string                                        `xml:"urn:com.workday/bsvc Tenant_Formatted_Phone,attr,omitempty"`
	InternationalFormattedPhone string                                        `xml:"urn:com.workday/bsvc International_Formatted_Phone,attr,omitempty"`
	PhoneNumberWithoutAreaCode  string                                        `xml:"urn:com.workday/bsvc Phone_Number_Without_Area_Code,attr,omitempty"`
	NationalFormattedPhone      string                                        `xml:"urn:com.workday/bsvc National_Formatted_Phone,attr,omitempty"`
	E164FormattedPhone          string                                        `xml:"urn:com.workday/bsvc E164_Formatted_Phone,attr,omitempty"`
	Delete                      bool                                          `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	DoNotReplaceAll             bool                                          `xml:"urn:com.workday/bsvc Do_Not_Replace_All,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PhoneReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type PhoneReferenceObjectType struct {
	ID         []PhoneReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type ProcessingFaultType struct {
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ProcurementItemObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ProcurementItemObjectType struct {
	ID         []ProcurementItemObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type PurchaseOrderLineObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type PurchaseOrderLineObjectType struct {
	ID         []PurchaseOrderLineObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Wrapper element for all Ad Hoc Payee request data
type PutAdHocPayeeRequestType struct {
	AdHocPayeeReference *AdHocPayeeObjectType `xml:"urn:com.workday/bsvc Ad_Hoc_Payee_Reference,omitempty"`
	AdHocPayeeData      *AdHocPayeeDataType   `xml:"urn:com.workday/bsvc Ad_Hoc_Payee_Data,omitempty"`
	AddOnly             bool                  `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version             string                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Ad Hoc Payee Response wrapper element
type PutAdHocPayeeResponseType struct {
	AdHocPayeeReference *AdHocPayeeObjectType `xml:"urn:com.workday/bsvc Ad_Hoc_Payee_Reference,omitempty"`
	Version             string                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The Authority Type reference for update and all Authority Type data items.
type PutAuthorityTypeRequestType struct {
	AuthorityTypeReference *AuthorityTypeObjectType `xml:"urn:com.workday/bsvc Authority_Type_Reference,omitempty"`
	AuthorityTypeData      *AuthorityTypeDataType   `xml:"urn:com.workday/bsvc Authority_Type_Data,omitempty"`
	AddOnly                bool                     `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the reference to the instance that is created or updated.
type PutAuthorityTypeResponseType struct {
	AuthorityTypeReference *AuthorityTypeObjectType `xml:"urn:com.workday/bsvc Authority_Type_Reference,omitempty"`
	Version                string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains data for adding or updating a Bank Account
type PutBankAccountRequestType struct {
	BankAccountReference *BankAccountObjectType `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	BankAccountData      *BankAccountDataType   `xml:"urn:com.workday/bsvc Bank_Account_Data"`
	AddOnly              bool                   `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version              string                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Put Ad hoc Bank Account Response Data
type PutBankAccountResponseType struct {
	BankAccountReference *BankAccountObjectType `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	Version              string                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The Bank Account reference for update and all of its Signatory data items.
type PutBankAccountSignatoryRequestType struct {
	BankAccountReference     *BankAccountObjectType         `xml:"urn:com.workday/bsvc Bank_Account_Reference"`
	BankAccountSignatoryData []BankAccountSignatoryDataType `xml:"urn:com.workday/bsvc Bank_Account_Signatory_Data,omitempty"`
	Version                  string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the reference to the Bank Account that is updated and the Signatories that are created, updated, or deleted.
type PutBankAccountSignatoryResponseType struct {
	BankAccountReference                 *BankAccountObjectType                         `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	BankAccountSignatoriesEventReference *UniqueIdentifierObjectType                    `xml:"urn:com.workday/bsvc Bank_Account_Signatories_Event_Reference,omitempty"`
	BankAccountSignatoryDetails          []BankAccountSignatoryDetailsType              `xml:"urn:com.workday/bsvc Bank_Account_Signatory_Details,omitempty"`
	ExceptionsResponseData               []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                              string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains data for creating a Bank Branch
type PutBankBranchRequestType struct {
	BankBranchReference *BankBranchObjectType `xml:"urn:com.workday/bsvc Bank_Branch_Reference,omitempty"`
	BankBranchData      *BankBranchDataType   `xml:"urn:com.workday/bsvc Bank_Branch_Data"`
	AddOnly             bool                  `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version             string                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Put Bank Branch Response Data
type PutBankBranchResponseType struct {
	BankBranchReference *BankBranchObjectType `xml:"urn:com.workday/bsvc Bank_Branch_Reference,omitempty"`
	Version             string                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Bank Fee Service Codes Request
type PutBankFeeServiceCodeRequestType struct {
	BankFeeServiceCodeReference *BankFeeServiceCodeObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Service_Code_Reference,omitempty"`
	BankFeeServiceCodeData      *BankFeeServiceCodeDataType   `xml:"urn:com.workday/bsvc Bank_Fee_Service_Code_Data,omitempty"`
	AddOnly                     bool                          `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                     string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Get Bank Fee Service Codes Response
type PutBankFeeServiceCodeResponseType struct {
	BankFeeServiceCodeReference *BankFeeServiceCodeObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Service_Code_Reference,omitempty"`
	Version                     string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The service contract request data.
type PutBankFeeServiceContractRequestType struct {
	BankFeeServiceContractReference *BankFeeServiceContractObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Service_Contract_Reference,omitempty"`
	BankFeeServiceContractData      *BankFeeServiceContractDataType   `xml:"urn:com.workday/bsvc Bank_Fee_Service_Contract_Data"`
	AddOnly                         bool                              `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                         string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The element container for the Service Contract Response
type PutBankFeeServiceContractResponseType struct {
	BankFeeServiceContractReference *BankFeeServiceContractObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Service_Contract_Reference,omitempty"`
	Version                         string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Bank Fee Statement File reference for update and all Bank Fee Statement File Data items.
type PutBankFeeStatementFileRequestType struct {
	BankFeeStatementFileReference *BankFeeStatementFileObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Statement_File_Reference,omitempty"`
	BankFeeStatementFileData      *BankFeeStatementFileDataType   `xml:"urn:com.workday/bsvc Bank_Fee_Statement_File_Data,omitempty"`
	AddOnly                       bool                            `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                       string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Put Bank Fee Statement File Response Data
type PutBankFeeStatementFileResponseType struct {
	BankFeeStatementFileReference *BankFeeStatementFileObjectType `xml:"urn:com.workday/bsvc Bank_Fee_Statement_File_Reference,omitempty"`
	Version                       string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains data for adding or updating a Bank Statement Custom Code Set
type PutBankStatementCustomCodeSetRequestType struct {
	BankStatementCustomCodeSetReference *BankStatementCustomCodeSetObjectType `xml:"urn:com.workday/bsvc Bank_Statement_Custom_Code_Set_Reference,omitempty"`
	BankStatementCustomCodeSetData      *BankStatementCustomCodeSetDataType   `xml:"urn:com.workday/bsvc Bank_Statement_Custom_Code_Set_Data"`
	AddOnly                             bool                                  `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                             string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Put Bank Statement Custom Code Set Response data
type PutBankStatementCustomCodeSetResponseType struct {
	BankStatementCustomCodeSetReference *BankStatementCustomCodeSetObjectType `xml:"urn:com.workday/bsvc Bank_Statement_Custom_Code_Set_Reference,omitempty"`
	Version                             string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Bank Statement File reference for update and all Bank Statement File Data items.
type PutBankStatementFileRequestType struct {
	BankStatementFileReference *BankStatementFileObjectType `xml:"urn:com.workday/bsvc Bank_Statement_File_Reference,omitempty"`
	BankStatementFileData      *BankStatementFileDataType   `xml:"urn:com.workday/bsvc Bank_Statement_File_Data"`
	AddOnly                    bool                         `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Put Bank Statement File Response Data
type PutBankStatementFileResponseType struct {
	BankStatementFileReference *BankStatementFileObjectType `xml:"urn:com.workday/bsvc Bank_Statement_File_Reference,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Bank Statement reference for update and all Bank Statement data items.
type PutBankStatementRequestType struct {
	BankStatementReference *BankStatementObjectType `xml:"urn:com.workday/bsvc Bank_Statement_Reference,omitempty"`
	BankStatementData      *BankStatementDataType   `xml:"urn:com.workday/bsvc Bank_Statement_Data"`
	AddOnly                bool                     `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top level 'Put' Web Service used to insert new instances of Bank Statement Response, confirm the response from the operation.
type PutBankStatementResponseType struct {
	BankStatementReference *BankStatementObjectType `xml:"urn:com.workday/bsvc Bank_Statement_Reference,omitempty"`
	Version                string                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains data for creating or updating Business Entity Contact data
type PutBusinessEntityContactRequestType struct {
	BusinessEntityContactReference *BusinessEntityContactObjectType `xml:"urn:com.workday/bsvc Business_Entity_Contact_Reference,omitempty"`
	BusinessEntityContactData      *BusinessEntityContactDataType   `xml:"urn:com.workday/bsvc Business_Entity_Contact_Data"`
	AddOnly                        bool                             `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                        string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Put Business Entity Contact  Response Data
type PutBusinessEntityContactResponseType struct {
	BusinessEntityContactReference *BusinessEntityContactObjectType `xml:"urn:com.workday/bsvc Business_Entity_Contact_Reference,omitempty"`
	Version                        string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The Cash Activity Category reference for update and all Cash Activity Category data items.
type PutCashActivityCategoryRequestType struct {
	CashActivityCategoryReference *CashActivityCategoryObjectType `xml:"urn:com.workday/bsvc Cash_Activity_Category_Reference,omitempty"`
	CashActivityCategoryData      *CashActivityCategoryDataType   `xml:"urn:com.workday/bsvc Cash_Activity_Category_Data,omitempty"`
	AddOnly                       bool                            `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                       string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the reference to the instance that is created or updated.
type PutCashActivityCategoryResponseType struct {
	CashActivityCategoryReference *CashActivityCategoryObjectType `xml:"urn:com.workday/bsvc Cash_Activity_Category_Reference,omitempty"`
	Version                       string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The Cash Pool reference for update and all Cash Pool data items.
type PutCashPoolRequestType struct {
	CashPoolReference *CashPoolObjectType `xml:"urn:com.workday/bsvc Cash_Pool_Reference,omitempty"`
	CashPoolData      *CashPoolDataType   `xml:"urn:com.workday/bsvc Cash_Pool_Data,omitempty"`
	AddOnly           bool                `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version           string              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the reference to the instance that is created or updated.
type PutCashPoolResponseType struct {
	CashPoolReference *CashPoolObjectType `xml:"urn:com.workday/bsvc Cash_Pool_Reference,omitempty"`
	Version           string              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Donor reference for update and all Donor data items
type PutDonorRequestType struct {
	DonorReference *DonorObjectType `xml:"urn:com.workday/bsvc Donor_Reference,omitempty"`
	DonorData      *DonorDataType   `xml:"urn:com.workday/bsvc Donor_Data"`
	AddOnly        bool             `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version        string           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Put Donor Response Data
type PutDonorResponseType struct {
	DonorReference *DonorObjectType `xml:"urn:com.workday/bsvc Donor_Reference,omitempty"`
	Version        string           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// contains data for creating or updating a Financial Institution
type PutFinancialInstitutionRequestType struct {
	FinancialInstitutionReference *FinancialInstitutionObjectType  `xml:"urn:com.workday/bsvc Financial_Institution_Reference,omitempty"`
	FinancialInstitutionData      *FinancialInstitutionWWSDataType `xml:"urn:com.workday/bsvc Financial_Institution_Data"`
	AddOnly                       bool                             `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                       string                           `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Put Financial Institution Response Data
type PutFinancialInstitutionResponseType struct {
	FinancialInstitutionReference *FinancialInstitutionObjectType `xml:"urn:com.workday/bsvc Financial_Institution_Reference,omitempty"`
	Version                       string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Import Process Response element
type PutImportProcessResponseType struct {
	ImportProcessReference  *WebServiceBackgroundProcessRuntimeObjectType `xml:"urn:com.workday/bsvc Import_Process_Reference,omitempty"`
	HeaderInstanceReference *InstanceObjectType                           `xml:"urn:com.workday/bsvc Header_Instance_Reference,omitempty"`
	Version                 string                                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Catch-up Payout Criteria Request
type PutInvestmentPoolCatchupPayoutCriteriaRequestType struct {
	CatchupPayoutCriteriaData *CatchupPayoutCriteriaDataType `xml:"urn:com.workday/bsvc Catch-up_Payout_Criteria_Data"`
	AddOnly                   bool                           `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                   string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Catch-up Payout Response
type PutInvestmentPoolCatchupPayoutCriteriaResponseType struct {
	InvestmentPoolPayoutCriteriaReference *UniqueIdentifierObjectType           `xml:"urn:com.workday/bsvc Investment_Pool_Payout_Criteria_Reference,omitempty"`
	InvestmentPoolPayoutCriteria          *InvestmentPoolPayoutCriteriaDataType `xml:"urn:com.workday/bsvc Investment_Pool_Payout_Criteria,omitempty"`
	Version                               string                                `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Payment Acknowledgement Message Request provide the abilty to add or update acknowledgements for payment file or payments providing status change and error code information.
type PutPaymentAcknowledgementMessageRequestType struct {
	PaymentAcknowledgementMessageReference *UniqueIdentifierObjectType            `xml:"urn:com.workday/bsvc Payment_Acknowledgement_Message_Reference,omitempty"`
	PaymentAcknowledgementMessageData      *PaymentAcknowledgementMessageDataType `xml:"urn:com.workday/bsvc Payment_Acknowledgement_Message_Data"`
	AddOnly                                bool                                   `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                                string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response to an add or updated Payment Acknowledgement Message
type PutPaymentAcknowledgementMessageResponseType struct {
	PaymentAcknowledgementMessageReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Payment_Acknowledgement_Message_Reference,omitempty"`
	Version                                string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Put Payment Election Enrollment Response
type PutPaymentElectionEnrollmentResponseType struct {
	PaymentElectionEnrollmentReference *UniqueIdentifierObjectType                    `xml:"urn:com.workday/bsvc Payment_Election_Enrollment_Reference,omitempty"`
	ExceptionsResponseData             []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                            string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Submit Payment Election Option Request
type PutPaymentElectionOptionRequestType struct {
	PaymentElectionOptionReference *UniqueIdentifierObjectType    `xml:"urn:com.workday/bsvc Payment_Election_Option_Reference,omitempty"`
	PaymentElectionOptionData      *PaymentElectionOptionDataType `xml:"urn:com.workday/bsvc Payment_Election_Option_Data,omitempty"`
	AddOnly                        bool                           `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                        string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Payment Election Option References
type PutPaymentElectionOptionResponseType struct {
	PaymentElectionOptionReference *UniqueIdentifierObjectType `xml:"urn:com.workday/bsvc Payment_Election_Option_Reference,omitempty"`
	Version                        string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains data for creating or updating a Petty Cash Account
type PutPettyCashAccountRequestType struct {
	PettyCashAccountReference *PettyCashAccountObjectType `xml:"urn:com.workday/bsvc Petty_Cash_Account_Reference,omitempty"`
	PettyCashAccountData      *PettyCashAccountDataType   `xml:"urn:com.workday/bsvc Petty_Cash_Account_Data"`
	AddOnly                   bool                        `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                   string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Put Petty Cash Response Data
type PutPettyCashAccountResponseType struct {
	PettyCashAccountReference *PettyCashAccountObjectType `xml:"urn:com.workday/bsvc Petty_Cash_Account_Reference,omitempty"`
	Version                   string                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The Signature Method reference for update and all Signature Method data items.
type PutSignatureMethodRequestType struct {
	SignatureMethodReference *SignatureMethodObjectType `xml:"urn:com.workday/bsvc Signature_Method_Reference,omitempty"`
	SignatureMethodData      *SignatureMethodDataType   `xml:"urn:com.workday/bsvc Signature_Method_Data,omitempty"`
	AddOnly                  bool                       `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                  string                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the reference to the instance that is created or updated.
type PutSignatureMethodResponseType struct {
	SignatureMethodReference *SignatureMethodObjectType `xml:"urn:com.workday/bsvc Signature_Method_Reference,omitempty"`
	Version                  string                     `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The Signer reference for update and all Signer data items.
type PutSignerRequestType struct {
	SignerReference *SignerObjectType `xml:"urn:com.workday/bsvc Signer_Reference,omitempty"`
	SignerData      *SignerDataType   `xml:"urn:com.workday/bsvc Signer_Data,omitempty"`
	AddOnly         bool              `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version         string            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the reference to the instance that is created or updated.
type PutSignerResponseType struct {
	SignerReference *SignerObjectType `xml:"urn:com.workday/bsvc Signer_Reference,omitempty"`
	Version         string            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains attributes required to recalculate the bank statement balance.
type RecalculateBankStatementBalanceDataType struct {
	BankStatementCriteriaData []BankStatementCriteriaDataType `xml:"urn:com.workday/bsvc Bank_Statement_Criteria_Data"`
}

// Contains parms required to recalculate bank statement beginning balance
type RecalculateBankStatementBalanceRequestType struct {
	RecalculateBankStatementBalanceData *RecalculateBankStatementBalanceDataType `xml:"urn:com.workday/bsvc Recalculate_Bank_Statement_Balance_Data"`
	Version                             string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains response of recalculated bank statement balance
type RecalculateBankStatementBalanceResponseType struct {
	RecalculatedBankStatement []RecalculatedBankStatementDataType `xml:"urn:com.workday/bsvc Recalculated_Bank_Statement,omitempty"`
	Version                   string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains Reference of all Bank Statements whose Beginning balance was recalculated.
type RecalculatedBankStatementDataType struct {
	BankStatementReference *BankStatementObjectType `xml:"urn:com.workday/bsvc Bank_Statement_Reference,omitempty"`
}

// Receiving Party Bank Data
type ReceivingPartyBankDataType struct {
	BankAccountNickname  string                               `xml:"urn:com.workday/bsvc Bank_Account_Nickname,omitempty"`
	AccountNumber        string                               `xml:"urn:com.workday/bsvc Account_Number,omitempty"`
	MaskedAccountNumber  string                               `xml:"urn:com.workday/bsvc Masked_Account_Number,omitempty"`
	AccountType          string                               `xml:"urn:com.workday/bsvc Account_Type,omitempty"`
	BankName             string                               `xml:"urn:com.workday/bsvc Bank_Name,omitempty"`
	IBAN                 string                               `xml:"urn:com.workday/bsvc IBAN,omitempty"`
	BankIDNumber         string                               `xml:"urn:com.workday/bsvc Bank_ID_Number,omitempty"`
	BIC                  string                               `xml:"urn:com.workday/bsvc BIC,omitempty"`
	BranchName           string                               `xml:"urn:com.workday/bsvc Branch_Name,omitempty"`
	BranchIDNumber       string                               `xml:"urn:com.workday/bsvc Branch_ID_Number,omitempty"`
	CountryReference     *CountryObjectType                   `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CurrencyReference    *InstanceObjectType                  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	BankAccountName      string                               `xml:"urn:com.workday/bsvc Bank_Account_Name,omitempty"`
	CheckDigit           string                               `xml:"urn:com.workday/bsvc Check_Digit,omitempty"`
	RollNumber           string                               `xml:"urn:com.workday/bsvc Roll_Number,omitempty"`
	BankInstructions     string                               `xml:"urn:com.workday/bsvc Bank_Instructions,omitempty"`
	IntermediaryBankData []IntermediaryBankAccountWWSDataType `xml:"urn:com.workday/bsvc Intermediary_Bank_Data,omitempty"`
}

// Receiving Party Cedit Card Data
type ReceivingPartyCreditCardDataType struct {
	MerchantAccountID         string `xml:"urn:com.workday/bsvc Merchant_Account_ID,omitempty"`
	CustomerProfileID         string `xml:"urn:com.workday/bsvc Customer_Profile_ID,omitempty"`
	CreditCardAuthorizationID string `xml:"urn:com.workday/bsvc Credit_Card_Authorization_ID,omitempty"`
	MerchantProcessingID      string `xml:"urn:com.workday/bsvc Merchant_Processing_ID,omitempty"`
	CreditCardCredit          *bool  `xml:"urn:com.workday/bsvc Credit_Card_Credit,omitempty"`
}

// Receiving Party Data
type ReceivingPartyWWSDataType struct {
	PayeeName                        string                                `xml:"urn:com.workday/bsvc Payee_Name"`
	PayeeLegalName                   string                                `xml:"urn:com.workday/bsvc Payee_Legal_Name"`
	PayeePreferredName               string                                `xml:"urn:com.workday/bsvc Payee_Preferred_Name"`
	PayeeID                          string                                `xml:"urn:com.workday/bsvc Payee_ID,omitempty"`
	ExpensePayeeEmployeeNumber       string                                `xml:"urn:com.workday/bsvc Expense_Payee_Employee_Number,omitempty"`
	TaxID                            string                                `xml:"urn:com.workday/bsvc Tax_ID,omitempty"`
	TaxIDType                        string                                `xml:"urn:com.workday/bsvc Tax_ID_Type,omitempty"`
	CorporateCreditCardAccountNumber string                                `xml:"urn:com.workday/bsvc Corporate_Credit_Card_Account_Number,omitempty"`
	MarketCode                       string                                `xml:"urn:com.workday/bsvc Market_Code,omitempty"`
	CountryReference                 *CountryObjectType                    `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	PayeeAlternateName               []BusinessEntityAlternateNameDataType `xml:"urn:com.workday/bsvc Payee_Alternate_Name,omitempty"`
	EmailAddressData                 []EmailAddressDataType                `xml:"urn:com.workday/bsvc Email_Address_Data,omitempty"`
	AddressInformationData           []AddressInformationDataType          `xml:"urn:com.workday/bsvc Address_Information_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ReconciliationRuleSetObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ReconciliationRuleSetObjectType struct {
	ID         []ReconciliationRuleSetObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                              `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ReconciliationStatusObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type ReconciliationStatusObjectType struct {
	ID         []ReconciliationStatusObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RegionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RegionObjectType struct {
	ID         []RegionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Remittance Data
type RemittanceDataType struct {
	UnstructuredRemittanceData []UnstructuredRemittanceDataType `xml:"urn:com.workday/bsvc Unstructured_Remittance_Data,omitempty"`
	DocumentRemittanceData     []DocumentRemittanceDataType     `xml:"urn:com.workday/bsvc Document_Remittance_Data,omitempty"`
}

// Remittance File Search Criteria
type RemittanceFileCriteriaType struct {
	RemittanceFileReference *RemittanceFileObjectType `xml:"urn:com.workday/bsvc Remittance_File_Reference"`
}

// Contains a unique identifier for an instance of an object.
type RemittanceFileObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RemittanceFileObjectType struct {
	ID         []RemittanceFileObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                       `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type ReportingTransactionObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type ReportingTransactionObjectType struct {
	ID         []ReportingTransactionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Parameters that let you filter the data returned in the response. You can filter returned data by dates and page attributes.
type ResponseFilterType struct {
	AsOfEffectiveDate *time.Time `xml:"urn:com.workday/bsvc As_Of_Effective_Date,omitempty"`
	AsOfEntryDateTime *time.Time `xml:"urn:com.workday/bsvc As_Of_Entry_DateTime,omitempty"`
	Page              float64    `xml:"urn:com.workday/bsvc Page,omitempty"`
	Count             float64    `xml:"urn:com.workday/bsvc Count,omitempty"`
}

func (t *ResponseFilterType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ResponseFilterType
	var layout struct {
		*T
		AsOfEffectiveDate *xsdDate     `xml:"urn:com.workday/bsvc As_Of_Effective_Date,omitempty"`
		AsOfEntryDateTime *xsdDateTime `xml:"urn:com.workday/bsvc As_Of_Entry_DateTime,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AsOfEffectiveDate = (*xsdDate)(layout.T.AsOfEffectiveDate)
	layout.AsOfEntryDateTime = (*xsdDateTime)(layout.T.AsOfEntryDateTime)
	return e.EncodeElement(layout, start)
}
func (t *ResponseFilterType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ResponseFilterType
	var overlay struct {
		*T
		AsOfEffectiveDate *xsdDate     `xml:"urn:com.workday/bsvc As_Of_Effective_Date,omitempty"`
		AsOfEntryDateTime *xsdDateTime `xml:"urn:com.workday/bsvc As_Of_Entry_DateTime,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AsOfEffectiveDate = (*xsdDate)(overlay.T.AsOfEffectiveDate)
	overlay.AsOfEntryDateTime = (*xsdDateTime)(overlay.T.AsOfEntryDateTime)
	return d.DecodeElement(&overlay, &start)
}

// The "Response_Results" element contains summary information about the data that has been returned from your request including "Total_Results", "Total_Pages", and the current "Page" returned.
type ResponseResultsType struct {
	TotalResults float64 `xml:"urn:com.workday/bsvc Total_Results,omitempty"`
	TotalPages   float64 `xml:"urn:com.workday/bsvc Total_Pages,omitempty"`
	PageResults  float64 `xml:"urn:com.workday/bsvc Page_Results,omitempty"`
	Page         float64 `xml:"urn:com.workday/bsvc Page,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type RevenueCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type RevenueCategoryObjectType struct {
	ID         []RevenueCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Settlement Account information including Bank Account Type, Currency, Routing Transit Number, etc.
//
// Bank Account Type becomes required when using v30.1 of the web service.
type SettlementAccountWWSDataType struct {
	SettlementBankAccountID      string                               `xml:"urn:com.workday/bsvc Settlement_Bank_Account_ID,omitempty"`
	CountryReference             *CountryObjectType                   `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CurrencyReference            *CurrencyObjectType                  `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	BankAccountNickname          string                               `xml:"urn:com.workday/bsvc Bank_Account_Nickname,omitempty"`
	BankAccountTypeReference     *BankAccountTypeObjectType           `xml:"urn:com.workday/bsvc Bank_Account_Type_Reference,omitempty"`
	BankName                     string                               `xml:"urn:com.workday/bsvc Bank_Name,omitempty"`
	RoutingTransitNumber         string                               `xml:"urn:com.workday/bsvc Routing_Transit_Number,omitempty"`
	BranchID                     string                               `xml:"urn:com.workday/bsvc Branch_ID,omitempty"`
	BranchName                   string                               `xml:"urn:com.workday/bsvc Branch_Name,omitempty"`
	BankAccountNumber            string                               `xml:"urn:com.workday/bsvc Bank_Account_Number,omitempty"`
	CheckDigit                   string                               `xml:"urn:com.workday/bsvc Check_Digit,omitempty"`
	BankAccountName              string                               `xml:"urn:com.workday/bsvc Bank_Account_Name,omitempty"`
	RollNumber                   string                               `xml:"urn:com.workday/bsvc Roll_Number,omitempty"`
	IBAN                         string                               `xml:"urn:com.workday/bsvc IBAN,omitempty"`
	SWIFTBankIdentificationCode  string                               `xml:"urn:com.workday/bsvc SWIFT_Bank_Identification_Code,omitempty"`
	AcceptsPaymentTypesReference []PaymentTypeObjectType              `xml:"urn:com.workday/bsvc Accepts_Payment_Types_Reference,omitempty"`
	PaymentTypesReference        []PaymentTypeObjectType              `xml:"urn:com.workday/bsvc Payment_Types_Reference,omitempty"`
	ForSupplierConnectionsOnly   *bool                                `xml:"urn:com.workday/bsvc For_Supplier_Connections_Only,omitempty"`
	RequiresPrenote              *bool                                `xml:"urn:com.workday/bsvc Requires_Prenote,omitempty"`
	PaymentTypePrenoteReference  *PaymentTypeObjectType               `xml:"urn:com.workday/bsvc Payment_Type_Prenote_Reference,omitempty"`
	Inactive                     *bool                                `xml:"urn:com.workday/bsvc Inactive,omitempty"`
	BankInstructions             string                               `xml:"urn:com.workday/bsvc Bank_Instructions,omitempty"`
	IntermediaryBankData         []IntermediaryBankAccountWWSDataType `xml:"urn:com.workday/bsvc Intermediary_Bank_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SignatoryDetailsObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SignatoryDetailsObjectType struct {
	ID         []SignatoryDetailsObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Signature Method data.
type SignatureMethodDataType struct {
	SignatureMethodID string `xml:"urn:com.workday/bsvc Signature_Method_ID,omitempty"`
	SignatureMethod   string `xml:"urn:com.workday/bsvc Signature_Method"`
	Description       string `xml:"urn:com.workday/bsvc Description,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SignatureMethodObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SignatureMethodObjectType struct {
	ID         []SignatureMethodObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                        `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the reference to the Signature Method.
type SignatureMethodRequestReferencesType struct {
	SignatureMethodReference []SignatureMethodObjectType `xml:"urn:com.workday/bsvc Signature_Method_Reference"`
}

// Wrapper containing the response data.
type SignatureMethodResponseDataType struct {
	SignatureMethod []SignatureMethodType `xml:"urn:com.workday/bsvc Signature_Method,omitempty"`
}

// Response data containing the reference and data items.
type SignatureMethodType struct {
	SignatureMethodReference *SignatureMethodObjectType `xml:"urn:com.workday/bsvc Signature_Method_Reference,omitempty"`
	SignatureMethodData      []SignatureMethodDataType  `xml:"urn:com.workday/bsvc Signature_Method_Data,omitempty"`
}

// Signer data.
type SignerDataType struct {
	ID                               string                             `xml:"urn:com.workday/bsvc ID,omitempty"`
	EmployeeReference                *EmployeeObjectType                `xml:"urn:com.workday/bsvc Employee_Reference"`
	ExternalCommitteeMemberReference *ExternalCommitteeMemberObjectType `xml:"urn:com.workday/bsvc External_Committee_Member_Reference"`
	SecuredAttachment                []AttachmentDataWWSType            `xml:"urn:com.workday/bsvc Secured_Attachment,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SignerObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SignerObjectType struct {
	ID         []SignerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains the reference to the Signer.
type SignerRequestReferencesType struct {
	SignerReference []SignerObjectType `xml:"urn:com.workday/bsvc Signer_Reference"`
}

// Wrapper containing the response data.
type SignerResponseDataType struct {
	Signer []SignerType `xml:"urn:com.workday/bsvc Signer,omitempty"`
}

// Response data containing the reference and data items.
type SignerType struct {
	SignerReference *SignerObjectType `xml:"urn:com.workday/bsvc Signer_Reference,omitempty"`
	SignerData      *SignerDataType   `xml:"urn:com.workday/bsvc Signer_Data,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SpendCategoryObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SpendCategoryObjectType struct {
	ID         []SpendCategoryObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SponsorObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SponsorObjectType struct {
	ID         []SponsorObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Sub Bank Accounts data.
type SubBankAccountsDataType struct {
	SubBankAccountReference *FinancialAccountObjectType `xml:"urn:com.workday/bsvc Sub_Bank_Account_Reference"`
}

// Request for Submit Ad Hoc Payment Template
type SubmitAdHocPaymentTemplateRequestType struct {
	Submit                        *bool                           `xml:"urn:com.workday/bsvc Submit,omitempty"`
	AdHocPaymentTemplateReference *AdHocPaymentTemplateObjectType `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_Template_Reference,omitempty"`
	BusinessProcessCommentData    *BusinessProcessCommentDataType `xml:"urn:com.workday/bsvc Business_Process_Comment_Data,omitempty"`
	AdHocPaymentTemplateData      *AdHocPaymentTemplateDataType   `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_Template_Data"`
	AddOnly                       bool                            `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                       string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response for Submit Bank Account Transfer Template
type SubmitAdHocPaymentTemplateResponseType struct {
	AdHocPaymentTemplateReference      *AdHocPaymentTemplateObjectType `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_Template_Reference,omitempty"`
	AdHocPaymentTemplateEventReference *UniqueIdentifierObjectType     `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_Template_Event_Reference,omitempty"`
	Version                            string                          `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request for Submit Ad Hoc Payment from Template
type SubmitAdHocPaymentfromTemplateRequestType struct {
	BusinessProcessParameters *FinancialsBusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AdHocPaymentData          *AdHocPaymentfromTemplateDataType        `xml:"urn:com.workday/bsvc Ad_Hoc_Payment_Data"`
	Version                   string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The response for Submit Ad hoc Payment from Template
type SubmitAdHocPaymentfromTemplateResponseType struct {
	AdhocPaymentReference  *AdHocPaymentObjectType                        `xml:"urn:com.workday/bsvc Ad_hoc_Payment_Reference,omitempty"`
	ExceptionsResponseData []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains data for creating or updating an ad hoc bank transaction and submitting for business processing
type SubmitAdhocBankTransactionRequestType struct {
	AdhocBankTransactionReference *AdHocBankTransactionObjectType          `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_Reference,omitempty"`
	BusinessProcessParameters     *FinancialsBusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AdhocBankTransactionData      *AdhocBankTransactionDataType            `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_Data,omitempty"`
	AddOnly                       bool                                     `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                       string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Put Ad hoc Bank Transaction Response Data
type SubmitAdhocBankTransactionResponseType struct {
	AdhocBankTransactionReference *AdHocBankTransactionObjectType                `xml:"urn:com.workday/bsvc Ad_hoc_Bank_Transaction_Reference,omitempty"`
	ExceptionsResponseData        []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                       string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains data for creating or updating an ad hoc payments and submitting for business processing
type SubmitAdhocPaymentRequestType struct {
	AdhocPaymentReference     *AdHocPaymentObjectType                  `xml:"urn:com.workday/bsvc Ad_hoc_Payment_Reference,omitempty"`
	BusinessProcessParameters *FinancialsBusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	AdhocPaymentData          *AdhocPaymentDataType                    `xml:"urn:com.workday/bsvc Ad_hoc_Payment_Data,omitempty"`
	AddOnly                   bool                                     `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                   string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Put Ad hoc Payment Response Data
type SubmitAdhocPaymentResponseType struct {
	AdhocPaymentReference  *AdHocPaymentObjectType                        `xml:"urn:com.workday/bsvc Ad_hoc_Payment_Reference,omitempty"`
	ExceptionsResponseData []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains data for adding or updating a Bank Account
type SubmitBankAccountRequestType struct {
	BankAccountReference      *BankAccountObjectType                   `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	BusinessProcessParameters *FinancialsBusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	BankAccountData           *BankAccountDataType                     `xml:"urn:com.workday/bsvc Bank_Account_Data"`
	AddOnly                   bool                                     `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                   string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Submit Bank Account Response data
type SubmitBankAccountResponseType struct {
	BankAccountReference *BankAccountObjectType `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	Version              string                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The Bank Account reference for update and all of its Signatory data items.
type SubmitBankAccountSignatoryRequestType struct {
	BankAccountReference      *BankAccountObjectType                   `xml:"urn:com.workday/bsvc Bank_Account_Reference"`
	BusinessProcessParameters *FinancialsBusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	BankAccountSignatoryData  []BankAccountSignatoryDataType           `xml:"urn:com.workday/bsvc Bank_Account_Signatory_Data,omitempty"`
	Version                   string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains the reference to the Bank Account that is updated and the Signatories that are created, updated, or deleted.
type SubmitBankAccountSignatoryResponseType struct {
	BankAccountReference                 *BankAccountObjectType                         `xml:"urn:com.workday/bsvc Bank_Account_Reference,omitempty"`
	BankAccountSignatoriesEventReference *UniqueIdentifierObjectType                    `xml:"urn:com.workday/bsvc Bank_Account_Signatories_Event_Reference,omitempty"`
	BankAccountSignatoryDetails          []BankAccountSignatoryDetailsType              `xml:"urn:com.workday/bsvc Bank_Account_Signatory_Details,omitempty"`
	ExceptionsResponseData               []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	Version                              string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The request object for Submit Bank Account Transfer request.
type SubmitBankAccountTransferRequestType struct {
	BankAccountTransferReference *BankAccountTransferObjectType            `xml:"urn:com.workday/bsvc Bank_Account_Transfer_Reference,omitempty"`
	BusinessProcessParameters    []FinancialsBusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	BankAccountTransferData      *BankAccountTransferDataType              `xml:"urn:com.workday/bsvc Bank_Account_Transfer_Data"`
	AddOnly                      bool                                      `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                      string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Submit bank account transfer response data containing reference to an existing bank account transfer.
type SubmitBankAccountTransferResponseType struct {
	BankAccountTransferReference *BankAccountTransferObjectType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_Reference,omitempty"`
	Version                      string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The request object for Submit Bank Account Transfer Template request.
type SubmitBankAccountTransferTemplateRequestType struct {
	Submit                               *bool                                  `xml:"urn:com.workday/bsvc Submit,omitempty"`
	BankAccountTransferTemplateReference *BankAccountTransferTemplateObjectType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_Template_Reference,omitempty"`
	BusinessProcessCommentData           *BusinessProcessCommentDataType        `xml:"urn:com.workday/bsvc Business_Process_Comment_Data,omitempty"`
	BankAccountTransferTemplateData      *BankAccountTransferTemplateDataType   `xml:"urn:com.workday/bsvc Bank_Account_Transfer_Template_Data"`
	AddOnly                              bool                                   `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                              string                                 `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Submit Bank Account Transfer Template Response element
type SubmitBankAccountTransferTemplateResponseType struct {
	BankAccountTransferTemplateReference      *BankAccountTransferTemplateObjectType         `xml:"urn:com.workday/bsvc Bank_Account_Transfer_Template_Reference,omitempty"`
	ExceptionsResponseData                    []ApplicationInstanceRelatedExceptionsDataType `xml:"urn:com.workday/bsvc Exceptions_Response_Data,omitempty"`
	BankAccountTransferTemplateEventReference *UniqueIdentifierObjectType                    `xml:"urn:com.workday/bsvc Bank_Account_Transfer_Template_Event_Reference,omitempty"`
	Version                                   string                                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The request object for Submit Bank Account Transfer for Settlement request containing data for create or update and submit for business processing
type SubmitBankAccountTransferforSettlementRequestType struct {
	BankAccountTransferforSettlementReference *BankAccountTransferforSettlementObjectType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Reference,omitempty"`
	BusinessProcessParameters                 *FinancialsBusinessProcessParametersType    `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	BankAccountTransferforSettlementData      *BankAccountTransferforSettlementDataType   `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Data"`
	AddOnly                                   bool                                        `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                                   string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Top level response element containing the created or updated bank account transfer for settlement.
type SubmitBankAccountTransferforSettlementResponseType struct {
	BankAccountTransferforSettlementReference *BankAccountTransferforSettlementObjectType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Reference,omitempty"`
	Version                                   string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request for Submit Bank Account Transfer for Settlement Template
type SubmitBankAccountTransferforSettlementTemplateRequestType struct {
	Submit                                            *bool                                               `xml:"urn:com.workday/bsvc Submit,omitempty"`
	BankAccountTransferforSettlementTemplateReference *BankAccountTransferforSettlementTemplateObjectType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Template_Reference,omitempty"`
	BusinessProcessCommentData                        *BusinessProcessCommentDataType                     `xml:"urn:com.workday/bsvc Business_Process_Comment_Data,omitempty"`
	BankAccountTransferforSettlementTemplateData      *BankAccountTransferforSettlementTemplateDataType   `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Template_Data"`
	AddOnly                                           bool                                                `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                                           string                                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response for Submit Bank Account Transfer for Settlement Template Response
type SubmitBankAccountTransferforSettlementTemplateResponseType struct {
	BankAccountTransferforSettlementTemplateReference      *BankAccountTransferforSettlementTemplateObjectType `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Template_Reference,omitempty"`
	BankAccountTransferforSettlementTemplateEventReference *UniqueIdentifierObjectType                         `xml:"urn:com.workday/bsvc Bank_Account_Transfer_for_Settlement_Template_Event_Reference,omitempty"`
	Version                                                string                                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains data for creating or updating a Donor Contribution and submitting for business processing
type SubmitDonorContributionRequestType struct {
	DonorContributionReference *DonorContributionObjectType              `xml:"urn:com.workday/bsvc Donor_Contribution_Reference,omitempty"`
	BusinessProcessParameters  []FinancialsBusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	DonorContributionData      *DonorContributionDataType                `xml:"urn:com.workday/bsvc Donor_Contribution_Data,omitempty"`
	AddOnly                    bool                                      `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                    string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing Submit Donor Contribution Response Data.
type SubmitDonorContributionResponseType struct {
	DonorContributionReference *DonorContributionObjectType `xml:"urn:com.workday/bsvc Donor_Contribution_Reference,omitempty"`
	Version                    string                       `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The Request Object for Submit Investment Pool Adjustment Web Service
type SubmitInvestmentPoolAdjustmentRequestType struct {
	InvestmentPoolAdjustmentReference *InvestmentPoolAdjustmentObjectType       `xml:"urn:com.workday/bsvc Investment_Pool_Adjustment_Reference,omitempty"`
	BusinessProcessParameters         []FinancialsBusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	InvestmentPoolAdjustmentData      *InvestmentPoolAdjustmentDataType         `xml:"urn:com.workday/bsvc Investment_Pool_Adjustment_Data"`
	AddOnly                           bool                                      `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                           string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The Response Object for Submit Investment Pool Transfer Web Service
type SubmitInvestmentPoolAdjustmentResponseType struct {
	InvestmentPoolAdjustmentReference *InvestmentPoolAdjustmentObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Adjustment_Reference,omitempty"`
	Version                           string                              `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains data for creating or updating an Investment Pool Purchase Record and submitting for business processing
type SubmitInvestmentPoolPurchaseRequestType struct {
	InvestmentPoolPurchaseReference *InvestmentPoolPurchaseObjectType         `xml:"urn:com.workday/bsvc Investment_Pool_Purchase_Reference,omitempty"`
	BusinessProcessParameters       []FinancialsBusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	InvestmentPoolPurchaseData      *InvestmentPoolPurchaseDataType           `xml:"urn:com.workday/bsvc Investment_Pool_Purchase_Data,omitempty"`
	AddOnly                         bool                                      `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                         string                                    `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Submit Investment Pool Purchase Record Response Data containing reference to existing Investment Pool Purchase Record.
type SubmitInvestmentPoolPurchaseResponseType struct {
	InvestmentPoolPurchaseReference *InvestmentPoolPurchaseObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Purchase_Reference,omitempty"`
	Version                         string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Request Element for Selling Investing Pool Units
type SubmitInvestmentPoolSaleRequestType struct {
	InvestmentPoolSaleReference *InvestmentPoolSaleObjectType            `xml:"urn:com.workday/bsvc Investment_Pool_Sale_Reference,omitempty"`
	BusinessProcessParameters   *FinancialsBusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	InvestmentPoolSaleData      *InvestmentPoolSaleDataType              `xml:"urn:com.workday/bsvc Investment_Pool_Sale_Data,omitempty"`
	AddOnly                     bool                                     `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                     string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Response Element for Selling Investment Pool Units
type SubmitInvestmentPoolSaleResponseType struct {
	InvestmentPoolSaleReference *InvestmentPoolSaleObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Sale_Reference,omitempty"`
	Version                     string                        `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The Request Object for Submit Investment Pool Transfer Web Service
type SubmitInvestmentPoolTransferRequestType struct {
	InvestmentPoolTransferReference *InvestmentPoolTransferObjectType        `xml:"urn:com.workday/bsvc Investment_Pool_Transfer_Reference,omitempty"`
	BusinessProcessParameters       *FinancialsBusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	InvestmentPoolTransferData      *InvestmentPoolTransferDataType          `xml:"urn:com.workday/bsvc Investment_Pool_Transfer_Data,omitempty"`
	AddOnly                         bool                                     `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                         string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The Response Object for Submit Investment Pool Transfer Web Service
type SubmitInvestmentPoolTransferResponseType struct {
	InvestmentPoolTransferReference *InvestmentPoolTransferObjectType `xml:"urn:com.workday/bsvc Investment_Pool_Transfer_Reference,omitempty"`
	Version                         string                            `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Contains data for creating or updating an Investment Statement and submitting for business processing
type SubmitInvestmentStatementRequestType struct {
	InvestmentStatementReference *InvestmentStatementObjectType           `xml:"urn:com.workday/bsvc Investment_Statement_Reference,omitempty"`
	BusinessProcessParameters    *FinancialsBusinessProcessParametersType `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	InvestmentStatementData      *InvestmentStatementDataType             `xml:"urn:com.workday/bsvc Investment_Statement_Data,omitempty"`
	AddOnly                      bool                                     `xml:"urn:com.workday/bsvc Add_Only,attr,omitempty"`
	Version                      string                                   `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Element containing a reference to an existing Investment Statement
type SubmitInvestmentStatementResponseType struct {
	InvestmentStatementReference *InvestmentStatementObjectType `xml:"urn:com.workday/bsvc Investment_Statement_Reference,omitempty"`
	Version                      string                         `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// Submit Payment Election Enrollment Request
type SubmitPaymentElectionEnrollmentRequestType struct {
	BusinessProcessParameters      *FinancialsBusinessProcessParametersType    `xml:"urn:com.workday/bsvc Business_Process_Parameters,omitempty"`
	PaymentElectionEnrollmentData  *PaymentElectionEnrollmentforWorkerDataType `xml:"urn:com.workday/bsvc Payment_Election_Enrollment_Data,omitempty"`
	RetainUnusedWorkerBankAccounts bool                                        `xml:"urn:com.workday/bsvc Retain_Unused_Worker_Bank_Accounts,attr,omitempty"`
	Version                        string                                      `xml:"urn:com.workday/bsvc version,attr,omitempty"`
}

// The submunicipality of the address.
type SubmunicipalityInformationDataType struct {
	Value                string `xml:",chardata"`
	AddressComponentName string `xml:"urn:com.workday/bsvc Address_Component_Name,attr,omitempty"`
	Type                 string `xml:"urn:com.workday/bsvc Type,attr,omitempty"`
}

// The subregion part of the address.
type SubregionInformationDataType struct {
	Value      string `xml:",chardata"`
	Descriptor string `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
	Type       string `xml:"urn:com.workday/bsvc Type,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SupervisoryOrganizationObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SupervisoryOrganizationObjectType struct {
	ID         []SupervisoryOrganizationObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SupplierContractLineObjectIDType struct {
	Value      string `xml:",chardata"`
	Type       string `xml:"urn:com.workday/bsvc type,attr"`
	Parentid   string `xml:"urn:com.workday/bsvc parent_id,attr,omitempty"`
	Parenttype string `xml:"urn:com.workday/bsvc parent_type,attr,omitempty"`
}

type SupplierContractLineObjectType struct {
	ID         []SupplierContractLineObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SupplierContractObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SupplierContractObjectType struct {
	ID         []SupplierContractObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element containing Supplier Invoice Line data.
type SupplierInvoiceLineReplacementDataType struct {
	SupplierInvoiceLineID          string                             `xml:"urn:com.workday/bsvc Supplier_Invoice_Line_ID,omitempty"`
	LineOrder                      string                             `xml:"urn:com.workday/bsvc Line_Order,omitempty"`
	IntercompanyAffiliateReference *CompanyObjectType                 `xml:"urn:com.workday/bsvc Intercompany_Affiliate_Reference,omitempty"`
	PurchaseItemReference          *ProcurementItemObjectType         `xml:"urn:com.workday/bsvc Purchase_Item_Reference,omitempty"`
	ItemDescription                string                             `xml:"urn:com.workday/bsvc Item_Description,omitempty"`
	PurchaseOrderLineReference     *PurchaseOrderLineObjectType       `xml:"urn:com.workday/bsvc Purchase_Order_Line_Reference,omitempty"`
	SupplierContractLineReference  *SupplierContractLineObjectType    `xml:"urn:com.workday/bsvc Supplier_Contract_Line_Reference,omitempty"`
	CustomerInvoiceLineReference   *CustomerInvoiceLineObjectType     `xml:"urn:com.workday/bsvc Customer_Invoice_Line_Reference,omitempty"`
	SpendCategoryReference         *SpendCategoryObjectType           `xml:"urn:com.workday/bsvc Spend_Category_Reference,omitempty"`
	ShipToAddressReference         *AddressReferenceObjectType        `xml:"urn:com.workday/bsvc Ship_To_Address_Reference,omitempty"`
	ShipToContactReference         *WorkerObjectType                  `xml:"urn:com.workday/bsvc Ship_To_Contact_Reference,omitempty"`
	AccountingTreatmentReference   *AccountingTreatmentObjectType     `xml:"urn:com.workday/bsvc Accounting_Treatment_Reference,omitempty"`
	TrackableItem                  *bool                              `xml:"urn:com.workday/bsvc Trackable_Item,omitempty"`
	TaxApplicabilityReference      *TaxApplicabilityObjectType        `xml:"urn:com.workday/bsvc Tax_Applicability_Reference,omitempty"`
	TaxCodeReference               *TaxCodeObjectType                 `xml:"urn:com.workday/bsvc Tax_Code_Reference,omitempty"`
	WithholdingTaxCodeReference    *TaxCodeObjectType                 `xml:"urn:com.workday/bsvc Withholding_Tax_Code_Reference,omitempty"`
	TaxData                        *TaxWidgetDataType                 `xml:"urn:com.workday/bsvc Tax_Data,omitempty"`
	TaxRateRecoverabilitiesData    *TaxRateOptionsDataType            `xml:"urn:com.workday/bsvc Tax_Rate_Recoverabilities_Data,omitempty"`
	PackagingString                string                             `xml:"urn:com.workday/bsvc Packaging_String,omitempty"`
	Quantity                       float64                            `xml:"urn:com.workday/bsvc Quantity,omitempty"`
	UnitofMeasureReference         *UnitofMeasureObjectType           `xml:"urn:com.workday/bsvc Unit_of_Measure_Reference,omitempty"`
	UnitCost                       float64                            `xml:"urn:com.workday/bsvc Unit_Cost,omitempty"`
	ExtendedAmount                 float64                            `xml:"urn:com.workday/bsvc Extended_Amount,omitempty"`
	RetentionAmount                float64                            `xml:"urn:com.workday/bsvc Retention_Amount,omitempty"`
	PaymentAmount                  float64                            `xml:"urn:com.workday/bsvc Payment_Amount,omitempty"`
	BudgetDate                     *time.Time                         `xml:"urn:com.workday/bsvc Budget_Date,omitempty"`
	Prepaid                        *bool                              `xml:"urn:com.workday/bsvc Prepaid,omitempty"`
	SupplierContractReference      *SupplierContractObjectType        `xml:"urn:com.workday/bsvc Supplier_Contract_Reference,omitempty"`
	Memo                           string                             `xml:"urn:com.workday/bsvc Memo,omitempty"`
	WorktagsReference              []AccountingWorktagObjectType      `xml:"urn:com.workday/bsvc Worktags_Reference,omitempty"`
	Billable                       *bool                              `xml:"urn:com.workday/bsvc Billable,omitempty"`
	SupplierInvoiceSplitLineData   []SupplierInvoiceLineSplitDataType `xml:"urn:com.workday/bsvc Supplier_Invoice_Split_Line_Data,omitempty"`
}

func (t *SupplierInvoiceLineReplacementDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T SupplierInvoiceLineReplacementDataType
	var layout struct {
		*T
		BudgetDate *xsdDate `xml:"urn:com.workday/bsvc Budget_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.BudgetDate = (*xsdDate)(layout.T.BudgetDate)
	return e.EncodeElement(layout, start)
}
func (t *SupplierInvoiceLineReplacementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SupplierInvoiceLineReplacementDataType
	var overlay struct {
		*T
		BudgetDate *xsdDate `xml:"urn:com.workday/bsvc Budget_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.BudgetDate = (*xsdDate)(overlay.T.BudgetDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing Supplier Invoice Line Split data.
type SupplierInvoiceLineSplitDataType struct {
	SupplierInvoiceLineSplitReference *SupplierInvoiceLineSplitObjectType  `xml:"urn:com.workday/bsvc Supplier_Invoice_Line_Split_Reference,omitempty"`
	SupplierInvoiceLineSplitID        string                               `xml:"urn:com.workday/bsvc Supplier_Invoice_Line_Split_ID,omitempty"`
	Quantity                          float64                              `xml:"urn:com.workday/bsvc Quantity,omitempty"`
	ExtendedAmount                    float64                              `xml:"urn:com.workday/bsvc Extended_Amount,omitempty"`
	Memo                              string                               `xml:"urn:com.workday/bsvc Memo,omitempty"`
	LineSplitAllocationReference      *BusinessDocumentLineSplitObjectType `xml:"urn:com.workday/bsvc Line_Split_Allocation_Reference,omitempty"`
	BudgetDate                        *time.Time                           `xml:"urn:com.workday/bsvc Budget_Date,omitempty"`
	WorktagReference                  []AccountingWorktagObjectType        `xml:"urn:com.workday/bsvc Worktag_Reference,omitempty"`
	Billable                          *bool                                `xml:"urn:com.workday/bsvc Billable,omitempty"`
}

func (t *SupplierInvoiceLineSplitDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T SupplierInvoiceLineSplitDataType
	var layout struct {
		*T
		BudgetDate *xsdDate `xml:"urn:com.workday/bsvc Budget_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.BudgetDate = (*xsdDate)(layout.T.BudgetDate)
	return e.EncodeElement(layout, start)
}
func (t *SupplierInvoiceLineSplitDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SupplierInvoiceLineSplitDataType
	var overlay struct {
		*T
		BudgetDate *xsdDate `xml:"urn:com.workday/bsvc Budget_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.BudgetDate = (*xsdDate)(overlay.T.BudgetDate)
	return d.DecodeElement(&overlay, &start)
}

// Contains a unique identifier for an instance of an object.
type SupplierInvoiceLineSplitObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SupplierInvoiceLineSplitObjectType struct {
	ID         []SupplierInvoiceLineSplitObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type SupplierObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type SupplierObjectType struct {
	ID         []SupplierObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                 `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TaxApplicabilityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TaxApplicabilityObjectType struct {
	ID         []TaxApplicabilityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TaxAuthorityFormTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TaxAuthorityFormTypeObjectType struct {
	ID         []TaxAuthorityFormTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TaxAuthorityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TaxAuthorityObjectType struct {
	ID         []TaxAuthorityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                     `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TaxCodeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TaxCodeObjectType struct {
	ID         []TaxCodeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TaxOptionObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TaxOptionObjectType struct {
	ID         []TaxOptionObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                  `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TaxPointDateTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TaxPointDateTypeObjectType struct {
	ID         []TaxPointDateTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Encapsulating element containing Tax Rate data.
type TaxRateApplicationDataType struct {
	TaxRateReference           *TaxRateObjectType           `xml:"urn:com.workday/bsvc Tax_Rate_Reference,omitempty"`
	TaxableAmount              float64                      `xml:"urn:com.workday/bsvc Taxable_Amount,omitempty"`
	TaxRecoverabilityReference *TaxRecoverabilityObjectType `xml:"urn:com.workday/bsvc Tax_Recoverability_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TaxRateObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TaxRateObjectType struct {
	ID         []TaxRateObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing Tax Rate Recoverabilities data for Tax Code.
type TaxRateOptionsDataType struct {
	TaxRate1Reference           *TaxRateObjectType           `xml:"urn:com.workday/bsvc Tax_Rate_1_Reference,omitempty"`
	TaxRecoverability1Reference *TaxRecoverabilityObjectType `xml:"urn:com.workday/bsvc Tax_Recoverability_1_Reference,omitempty"`
	TaxRate2Reference           []TaxRateObjectType          `xml:"urn:com.workday/bsvc Tax_Rate_2_Reference,omitempty"`
	TaxRecoverability2Reference *TaxRecoverabilityObjectType `xml:"urn:com.workday/bsvc Tax_Recoverability_2_Reference,omitempty"`
	TaxRate3Reference           []TaxRateObjectType          `xml:"urn:com.workday/bsvc Tax_Rate_3_Reference,omitempty"`
	TaxRecoverability3Reference *TaxRecoverabilityObjectType `xml:"urn:com.workday/bsvc Tax_Recoverability_3_Reference,omitempty"`
	TaxRate4Reference           []TaxRateObjectType          `xml:"urn:com.workday/bsvc Tax_Rate_4_Reference,omitempty"`
	TaxRecoverability4Reference *TaxRecoverabilityObjectType `xml:"urn:com.workday/bsvc Tax_Recoverability_4_Reference,omitempty"`
	TaxRate5Reference           []TaxRateObjectType          `xml:"urn:com.workday/bsvc Tax_Rate_5_Reference,omitempty"`
	TaxRecoverability5Reference *TaxRecoverabilityObjectType `xml:"urn:com.workday/bsvc Tax_Recoverability_5_Reference,omitempty"`
	TaxRate6Reference           []TaxRateObjectType          `xml:"urn:com.workday/bsvc Tax_Rate_6_Reference,omitempty"`
	TaxRecoverability6Reference *TaxRecoverabilityObjectType `xml:"urn:com.workday/bsvc Tax_Recoverability_6_Reference,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type TaxRecoverabilityObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TaxRecoverabilityObjectType struct {
	ID         []TaxRecoverabilityObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                          `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Element containing Tax Information
type TaxWidgetDataType struct {
	TaxPointDateTypeReference *TaxPointDateTypeObjectType `xml:"urn:com.workday/bsvc Tax_Point_Date_Type_Reference,omitempty"`
	TaxPointDate              *time.Time                  `xml:"urn:com.workday/bsvc Tax_Point_Date,omitempty"`
}

func (t *TaxWidgetDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T TaxWidgetDataType
	var layout struct {
		*T
		TaxPointDate *xsdDate `xml:"urn:com.workday/bsvc Tax_Point_Date,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TaxPointDate = (*xsdDate)(layout.T.TaxPointDate)
	return e.EncodeElement(layout, start)
}
func (t *TaxWidgetDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T TaxWidgetDataType
	var overlay struct {
		*T
		TaxPointDate *xsdDate `xml:"urn:com.workday/bsvc Tax_Point_Date,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TaxPointDate = (*xsdDate)(overlay.T.TaxPointDate)
	return d.DecodeElement(&overlay, &start)
}

// Encapsulating element containing Tax Details data
type TaxableCodeApplicationDataType struct {
	TaxApplicabilityReference *TaxApplicabilityObjectType  `xml:"urn:com.workday/bsvc Tax_Applicability_Reference"`
	TaxCodeReference          *TaxCodeObjectType           `xml:"urn:com.workday/bsvc Tax_Code_Reference"`
	TaxableAmount             float64                      `xml:"urn:com.workday/bsvc Taxable_Amount"`
	TaxRateData               []TaxRateApplicationDataType `xml:"urn:com.workday/bsvc Tax_Rate_Data"`
}

// Contains a unique identifier for an instance of an object.
type TaxpayerIDNumberTypeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type TaxpayerIDNumberTypeObjectType struct {
	ID         []TaxpayerIDNumberTypeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                             `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type UniqueIdentifierObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type UniqueIdentifierObjectType struct {
	ID         []UniqueIdentifierObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                         `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type UnitofMeasureObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type UnitofMeasureObjectType struct {
	ID         []UnitofMeasureObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                      `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Unstructured Remittance Data
type UnstructuredRemittanceDataType struct {
	AddendaLine string `xml:"urn:com.workday/bsvc Addenda_Line,omitempty"`
}

type ValidationErrorType struct {
	Message       string `xml:"urn:com.workday/bsvc Message,omitempty"`
	DetailMessage string `xml:"urn:com.workday/bsvc Detail_Message,omitempty"`
	Xpath         string `xml:"urn:com.workday/bsvc Xpath,omitempty"`
}

type ValidationFaultType struct {
	ValidationError []ValidationErrorType `xml:"urn:com.workday/bsvc Validation_Error,omitempty"`
}

// Web Address Information
type WebAddressInformationDataType struct {
	WebAddress          string                                        `xml:"urn:com.workday/bsvc Web_Address,omitempty"`
	WebAddressComment   string                                        `xml:"urn:com.workday/bsvc Web_Address_Comment,omitempty"`
	UsageData           []CommunicationMethodUsageInformationDataType `xml:"urn:com.workday/bsvc Usage_Data,omitempty"`
	WebAddressReference *WebAddressReferenceObjectType                `xml:"urn:com.workday/bsvc Web_Address_Reference,omitempty"`
	ID                  string                                        `xml:"urn:com.workday/bsvc ID,omitempty"`
	Delete              bool                                          `xml:"urn:com.workday/bsvc Delete,attr,omitempty"`
	DoNotReplaceAll     bool                                          `xml:"urn:com.workday/bsvc Do_Not_Replace_All,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WebAddressReferenceObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WebAddressReferenceObjectType struct {
	ID         []WebAddressReferenceObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                            `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WebServiceBackgroundProcessRuntimeObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WebServiceBackgroundProcessRuntimeObjectType struct {
	ID         []WebServiceBackgroundProcessRuntimeObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string                                           `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type WorkdayCommonHeaderType struct {
	IncludeReferenceDescriptorsInResponse *bool `xml:"urn:com.workday/bsvc Include_Reference_Descriptors_In_Response,omitempty"`
}

// Container element for Bank Account information
type WorkerBankAccountWWSDataType struct {
	CountryReference     *CountryObjectType         `xml:"urn:com.workday/bsvc Country_Reference,omitempty"`
	CurrencyReference    *CurrencyObjectType        `xml:"urn:com.workday/bsvc Currency_Reference,omitempty"`
	BankAccountNickname  string                     `xml:"urn:com.workday/bsvc Bank_Account_Nickname,omitempty"`
	BankAccountName      string                     `xml:"urn:com.workday/bsvc Bank_Account_Name,omitempty"`
	AccountNumber        string                     `xml:"urn:com.workday/bsvc Account_Number,omitempty"`
	RollNumber           string                     `xml:"urn:com.workday/bsvc Roll_Number,omitempty"`
	AccountTypeCode      string                     `xml:"urn:com.workday/bsvc Account_Type_Code,omitempty"`
	AccountTypeReference *BankAccountTypeObjectType `xml:"urn:com.workday/bsvc Account_Type_Reference,omitempty"`
	BankName             string                     `xml:"urn:com.workday/bsvc Bank_Name,omitempty"`
	IBAN                 string                     `xml:"urn:com.workday/bsvc IBAN,omitempty"`
	BankIDNumber         string                     `xml:"urn:com.workday/bsvc Bank_ID_Number,omitempty"`
	BIC                  string                     `xml:"urn:com.workday/bsvc BIC,omitempty"`
	BranchName           string                     `xml:"urn:com.workday/bsvc Branch_Name,omitempty"`
	BranchIDNumber       string                     `xml:"urn:com.workday/bsvc Branch_ID_Number,omitempty"`
	CheckDigit           string                     `xml:"urn:com.workday/bsvc Check_Digit,omitempty"`
}

// Contains a unique identifier for an instance of an object.
type WorkerObjectIDType struct {
	Value string `xml:",chardata"`
	Type  string `xml:"urn:com.workday/bsvc type,attr"`
}

type WorkerObjectType struct {
	ID         []WorkerObjectIDType `xml:"urn:com.workday/bsvc ID,omitempty"`
	Descriptor string               `xml:"urn:com.workday/bsvc Descriptor,attr,omitempty"`
}

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02")), nil
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
