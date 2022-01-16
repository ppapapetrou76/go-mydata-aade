package models

import (
	"fmt"

	"github.com/ppapapetrou76/go-utils/pkg/multierror"
)

type ResponseDoc struct {
	Response []*Response `xml:"response"`
}

type Response struct {
	StatusCode         string  `xml:"statusCode"`
	Errors             Errors  `xml:"errors"`
	Index              *uint   `xml:"index"`
	InvoiceUID         *string `xml:"invoiceUid"`
	InvoiceMark        *uint64 `xml:"invoiceMark"`
	ClassificationMark *uint64 `xml:"classificationMark"`
	AuthenticationCode *string `xml:"authenticationCode"`
	CancellationMark   *uint64 `xml:"cancellationMark"`
}

type Error struct {
	Message string `xml:"message"`
	Code    string `xml:"code"`
}

// Error implements the error interface and returns a human-readable representation of the given error.
func (e Error) Error() string {
	return fmt.Sprintf("Code:%s - Message:%s", e.Code, e.Message)
}

type Errors struct {
	Error []Error `xml:"error"`
}

// HasErrors returns true if the response contains at least one errors, else false.
func (rDoc ResponseDoc) HasErrors() bool {
	for _, resp := range rDoc.Response {
		if len(resp.Errors.Error) > 0 {
			return true
		}
	}

	return false
}

// Errors returns all errors of the response doc wrapped with a given prefix.
// If no errors found then it returns nil.
func (rDoc ResponseDoc) Errors(prefix string) error {
	errs := multierror.NewPrefixed(prefix)
	for _, resp := range rDoc.Response {
		for _, respErr := range resp.Errors.Error {
			errs = errs.Append(respErr)
		}
	}

	return errs.ErrorOrNil()
}
