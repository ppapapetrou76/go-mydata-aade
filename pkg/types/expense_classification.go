package types

type ExpensesClassificationsDoc struct {
	ExpensesInvoiceClassification struct {
		InvoiceMark        string `xml:"invoiceMark"`
		ClassificationMark string `xml:"classificationMark"`
		EntityVatNumber    string `xml:"entityVatNumber"`
		TransactionMode    string `xml:"transactionMode"`
	} `xml:"expensesInvoiceClassification"`
}
