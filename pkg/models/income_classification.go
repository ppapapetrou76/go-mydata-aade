package models

type IncomeClassificationsDoc struct {
	IncomeInvoiceClassification []IncomeInvoiceClassification `xml:"incomeInvoiceClassification"`
}

type IncomeInvoiceClassification struct {
	InvoiceMark                         uint64                                `xml:"invoiceMark"`
	ClassificationMark                  *uint64                               `xml:"classificationMark"`
	EntityVatNumber                     *string                               `xml:"entityVatNumber"`
	TransactionMode                     *uint8                                `xml:"transactionMode"`
	InvoicesIncomeClassificationDetails []InvoicesIncomeClassificationDetails `xml:"invoicesIncomeClassificationDetails"`
}

type InvoicesIncomeClassificationDetails struct {
	LineNumber                     uint                       `xml:"lineNumber"`
	IncomeClassificationDetailData []IncomeClassificationType `xml:"incomeClassificationDetailData"`
}
