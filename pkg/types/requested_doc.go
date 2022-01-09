package types

type RequestedDoc struct {
	InvoicesDoc struct {
		Invoices []Invoice `xml:"invoice"`
	} `xml:"invoicesDoc"`
}

type Issuer struct {
	VatNumber string `xml:"vatNumber"`
	Country   string `xml:"country"`
	Branch    string `xml:"branch"`
}

type Address struct {
	Street     string `xml:"street"`
	PostalCode string `xml:"postalCode"`
	City       string `xml:"city"`
}

type Counterpart struct {
	VatNumber string  `xml:"vatNumber"`
	Country   string  `xml:"country"`
	Branch    string  `xml:"branch"`
	Name      string  `xml:"name"`
	Address   Address `xml:"address"`
}

type InvoiceHeader struct {
	Series               string `xml:"series"`
	Aa                   string `xml:"aa"`
	IssueDate            string `xml:"issueDate"`
	InvoiceType          string `xml:"invoiceType"`
	VatPaymentSuspension string `xml:"vatPaymentSuspension"`
	Currency             string `xml:"currency"`
	ExchangeRate         string `xml:"exchangeRate"`
}

type IncomeClassification struct {
	ClassificationType     string `xml:"classificationType"`
	ClassificationCategory string `xml:"classificationCategory"`
	Amount                 string `xml:"amount"`
}

type InvoiceDetails struct {
	LineNumber           string               `xml:"lineNumber"`
	NetValue             string               `xml:"netValue"`
	VatCategory          string               `xml:"vatCategory"`
	VatAmount            string               `xml:"vatAmount"`
	VatExemptionCategory string               `xml:"vatExemptionCategory"`
	IncomeClassification IncomeClassification `xml:"incomeClassification"`
}

type InvoiceSummary struct {
	TotalNetValue         string               `xml:"totalNetValue"`
	TotalVatAmount        string               `xml:"totalVatAmount"`
	TotalWithheldAmount   string               `xml:"totalWithheldAmount"`
	TotalFeesAmount       string               `xml:"totalFeesAmount"`
	TotalStampDutyAmount  string               `xml:"totalStampDutyAmount"`
	TotalOtherTaxesAmount string               `xml:"totalOtherTaxesAmount"`
	TotalDeductionsAmount string               `xml:"totalDeductionsAmount"`
	TotalGrossValue       string               `xml:"totalGrossValue"`
	IncomeClassification  IncomeClassification `xml:"incomeClassification"`
}

type PaymentMethodDetails struct {
	Type   string `xml:"type"`
	Amount string `xml:"amount"`
}

type PaymentMethods struct {
	PaymentMethodDetails PaymentMethodDetails `xml:"paymentMethodDetails"`
}

type Invoice struct {
	UID            string         `xml:"uid"`
	Mark           string         `xml:"mark"`
	Issuer         Issuer         `xml:"issuer"`
	Counterpart    Counterpart    `xml:"counterpart"`
	InvoiceHeader  InvoiceHeader  `xml:"invoiceHeader"`
	PaymentMethods PaymentMethods `xml:"paymentMethods"`
	InvoiceDetails InvoiceDetails `xml:"invoiceDetails"`
	InvoiceSummary InvoiceSummary `xml:"invoiceSummary"`
}
