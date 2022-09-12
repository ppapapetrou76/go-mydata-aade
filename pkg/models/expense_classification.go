package models

import (
	"encoding/xml"
	"fmt"
)

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

// MarshalXML transforms an ExpensesClassificationType to expensesClassificationType and serializes it in order to
// include the `ecls` namespace for every field.
func (classification ExpensesClassificationType) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type expensesClassificationType struct {
		ClassificationType     string  `xml:"ecls:classificationType"`
		ClassificationCategory string  `xml:"ecls:classificationCategory"`
		Amount                 float64 `xml:"ecls:amount"`
		ID                     *byte   `xml:"ecls:id"`
	}

	err := enc.EncodeElement(expensesClassificationType(classification), start)
	if err != nil {
		return fmt.Errorf("xml marshal expense classification: %w", err)
	}

	return nil
}
