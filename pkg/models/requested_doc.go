package models

type RequestedDoc struct {
	ContinuationToken      ContinuationToken          `xml:"continuationToken"`
	InvoicesDoc            InvoicesDoc                `xml:"invoicesDoc"`
	CancelledInvoicesDoc   CancelledInvoicesDoc       `xml:"cancelledInvoicesDoc"`
	IncomeClassification   IncomeClassificationType   `xml:"incomeClassificationsDoc"`
	ExpensesClassification ExpensesClassificationType `xml:"expensesClassification"`
}

type InvoicesDoc struct {
	Xmlns    string    `xml:"xmlns,attr"`
	Invoices []Invoice `xml:"invoice"`
}

type ContinuationToken struct {
	NextPartitionKey string `xml:"nextPartitionKey"`
	NextRowKey       string `xml:"nextRowKey"`
}

type CancelledInvoicesDoc struct {
	CancelledInvoices []CancelledInvoice `xml:"cancelledInvoice"`
}

type CancelledInvoice struct {
	InvoiceMark      uint64 `xml:"invoiceMark"`
	CancellationMark uint64 `xml:"cancellationMark"`
	CancellationDate string `xml:"cancellationDate"`
}

type Address struct {
	Street     *string `xml:"street"`
	Number     *string `xml:"number"`
	PostalCode string  `xml:"postalCode"`
	City       string  `xml:"city"`
}

type PartyType struct {
	VatNumber string   `xml:"vatNumber"`
	Country   string   `xml:"country"`
	Branch    uint64   `xml:"branch"`
	Name      *string  `xml:"name"`
	Address   *Address `xml:"address"`
}

type InvoiceHeader struct {
	Series               string   `xml:"series"`
	Aa                   string   `xml:"aa"`
	IssueDate            string   `xml:"issueDate"`
	InvoiceType          string   `xml:"invoiceType"`
	VatPaymentSuspension *bool    `xml:"vatPaymentSuspension"`
	Currency             string   `xml:"currency"`
	ExchangeRate         *float64 `xml:"exchangeRate"`
	SelfPricing          *bool    `xml:"selfPricing"`
	CorrelatedInvoices   *uint    `xml:"correlatedInvoices"`
	DispatchDate         *string  `xml:"dispatchDate"`
	DispatchTime         *string  `xml:"dispatchTime"`
	VehicleNumber        *string  `xml:"vehicleNumber"`
	MovePurpose          *uint    `xml:"movePurpose"`
	FuelInvoice          *bool    `xml:"fuelInvoice"`
}

type IncomeClassificationType struct {
	ClassificationType     string `xml:"classificationType"`
	ClassificationCategory string `xml:"classificationCategory"`
	Amount                 string `xml:"amount"`
	ID                     *byte  `xml:"id"`
}

type ExpensesClassificationType struct {
	ClassificationType     string `xml:"classificationType"`
	ClassificationCategory string `xml:"classificationCategory"`
	Amount                 string `xml:"amount"`
	ID                     *byte  `xml:"id"`
}

type InvoiceDetails struct {
	LineNumber                uint                        `xml:"lineNumber"`
	RecType                   *uint                       `xml:"recType"`
	Quantity                  *float64                    `xml:"quantity"`
	MeasurementUnit           *uint                       `xml:"measurementUnit"`
	InvoiceDetailType         *uint                       `xml:"invoiceDetailType"`
	NetValue                  float64                     `xml:"netValue"`
	VatCategory               uint                        `xml:"vatCategory"`
	VatAmount                 float64                     `xml:"vatAmount"`
	VatExemptionCategory      *uint                       `xml:"vatExemptionCategory"`
	Dienergia                 *ShipType                   `xml:"dienergia"`
	DiscountOption            *bool                       `xml:"discountOption"`
	WithheldAmount            *float64                    `xml:"withheldAmount"`
	WithheldPercentCategory   *uint                       `xml:"withheldPercentCategory"`
	StampDutyAmount           *float64                    `xml:"stampDutyAmount"`
	StampDutyPercentCategory  *uint                       `xml:"stampDutyPercentCategory"`
	FeesAmount                *float64                    `xml:"feesAmount"`
	FeesPercentCategory       *uint                       `xml:"feesPercentCategory"`
	OtherTaxesPercentCategory *uint                       `xml:"otherTaxesPercentCategory"`
	OtherTaxesAmount          *float64                    `xml:"otherTaxesAmount"`
	DeductionsAmount          *float64                    `xml:"deductionsAmount"`
	IncomeClassification      *IncomeClassificationType   `xml:"incomeClassification"`
	ExpensesClassification    *ExpensesClassificationType `xml:"expensesClassification"`
}

type ShipType struct {
	ApplicationID   string `xml:"applicationId"`
	ApplicationDate string `xml:"applicationDate"`
	Doy             string `xml:"doy"`
	ShipID          string `xml:"shipID"`
}

type InvoiceSummary struct {
	TotalNetValue          float64                     `xml:"totalNetValue"`
	TotalVatAmount         float64                     `xml:"totalVatAmount"`
	TotalWithheldAmount    float64                     `xml:"totalWithheldAmount"`
	TotalFeesAmount        float64                     `xml:"totalFeesAmount"`
	TotalStampDutyAmount   float64                     `xml:"totalStampDutyAmount"`
	TotalOtherTaxesAmount  float64                     `xml:"totalOtherTaxesAmount"`
	TotalDeductionsAmount  float64                     `xml:"totalDeductionsAmount"`
	TotalGrossValue        float64                     `xml:"totalGrossValue"`
	IncomeClassification   *IncomeClassificationType   `xml:"incomeClassification"`
	ExpensesClassification *ExpensesClassificationType `xml:"expensesClassification"`
}

type PaymentMethodDetails struct {
	Type              uint    `xml:"type"`
	Amount            float64 `xml:"amount"`
	PaymentMethodInfo string  `xml:"paymentMethodInfo"`
}

type PaymentMethods struct {
	PaymentMethodDetails *PaymentMethodDetails `xml:"paymentMethodDetails"`
}

type Taxes struct {
	TaxType         byte     `xml:"taxType"`
	TaxCategory     *byte    `xml:"taxCategory"`
	UnderlyingValue *float64 `xml:"underlyingValue"`
	TaxAmount       float64  `xml:"taxAmount"`
	ID              *byte    `xml:"id"`
}

type TaxesTotals struct {
	Taxes Taxes `xml:"taxes"`
}

type Invoice struct {
	UID                 *string           `xml:"uid"`
	Mark                *uint64           `xml:"mark"`
	CancelledByMark     *uint64           `xml:"cancelledByMark"`
	AuthenticationCode  *string           `xml:"authenticationCode"`
	TransmissionFailure *byte             `xml:"transmissionFailure"`
	Issuer              *PartyType        `xml:"issuer"`
	Counterpart         *PartyType        `xml:"counterpart"`
	InvoiceHeader       *InvoiceHeader    `xml:"invoiceHeader"`
	PaymentMethods      *PaymentMethods   `xml:"paymentMethods"`
	InvoiceDetails      []*InvoiceDetails `xml:"invoiceDetails"`
	TaxesTotals         *TaxesTotals      `xml:"taxesTotals"`
	InvoiceSummary      *InvoiceSummary   `xml:"invoiceSummary"`
}
