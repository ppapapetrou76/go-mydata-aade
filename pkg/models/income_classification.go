package models

import (
	"encoding/xml"
	"fmt"
)

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

// MarshalXML transforms an IncomeClassificationType to incomeClassificationType and serializes it in order to include
// the `icls` namespace for every field.
func (classification IncomeClassificationType) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	type incomeClassificationType struct {
		ClassificationType     string  `xml:"icls:classificationType"`
		ClassificationCategory string  `xml:"icls:classificationCategory"`
		Amount                 float64 `xml:"icls:amount"`
		ID                     *byte   `xml:"icls:id"`
	}

	err := enc.EncodeElement(incomeClassificationType(classification), start)
	if err != nil {
		return fmt.Errorf("xml marshal income classification: %w", err)
	}

	return nil
}
