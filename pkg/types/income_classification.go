package types

type IncomeClassificationsDoc struct {
	IncomeInvoiceClassification struct {
		InvoiceMark        string `xml:"invoiceMark"`
		ClassificationMark string `xml:"classificationMark"`
		EntityVatNumber    string `xml:"entityVatNumber"`
		TransactionMode    string `xml:"transactionMode"`
	} `xml:"incomeInvoiceClassification"`
}
