package models

type ExpensesClassificationsDoc struct {
	ExpensesInvoiceClassification []ExpensesInvoiceClassification `xml:"ExpensesInvoiceClassification"`
}

type ExpensesInvoiceClassification struct {
	InvoiceMark                           uint64                                  `xml:"invoiceMark"`
	ClassificationMark                    *uint64                                 `xml:"classificationMark"`
	EntityVatNumber                       *string                                 `xml:"entityVatNumber"`
	TransactionMode                       *uint8                                  `xml:"transactionMode"`
	InvoicesExpensesClassificationDetails []InvoicesExpensesClassificationDetails `xml:"invoicesExpensesClassificationDetails"`
}

type InvoicesExpensesClassificationDetails struct {
	LineNumber                       uint                         `xml:"lineNumber"`
	ExpensesClassificationDetailData []ExpensesClassificationType `xml:"ExpensesClassificationDetailData"`
}
