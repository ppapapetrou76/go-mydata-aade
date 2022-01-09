package mydata

import (
	"bytes"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"

	"github.com/ppapapetrou76/go-mydata-aade/internal/httpmock"
	"github.com/ppapapetrou76/go-mydata-aade/pkg/types"
)

//go:embed testdata/request_response.xml
var response []byte

func TestAuthHeaderProvider_Intercept(t *testing.T) {
	authHeaderProvider := AuthHeaderProvider{
		userID:          "userID",
		subscriptionKey: "subKey",
	}
	req := &http.Request{Header: http.Header{}}
	err := authHeaderProvider.Intercept(context.Background(), req)

	assert.ThatError(t, err).IsNil()
	assert.That(t, req.Header.Get("aade-user-id")).IsEqualTo("userID")
	assert.That(t, req.Header.Get("Ocp-Apim-Subscription-Key")).IsEqualTo("subKey")
}

func TestNewClient(t *testing.T) {
	ft := assert.NewFluentT(t)

	t.Run("should fail if config is nil", func(t *testing.T) {
		client, err := NewClient(nil)
		ft.AssertThat(client).IsNil()
		ft.AssertThat(err).IsNotNil().IsEqualTo(errors.New("api config is required"))
	})

	t.Run("should fail if config validation fails", func(t *testing.T) {
		client, err := NewClient(&Config{})
		ft.AssertThat(client).IsNil()
		assert.ThatError(t, err).
			IsNotNil().
			HasExactMessage("api config validation 2 errors:[ subscription key is empty, user id is empty ]")
	})

	t.Run("should return new client with default values", func(t *testing.T) {
		client, err := NewClient(&Config{
			UserID:          "userID",
			SubscriptionKey: "subKey",
		})
		ft.AssertThat(err).IsNil()
		ft.AssertThat(client.timeout).IsEqualTo(defaultTimeout)
		ft.AssertThat(client.myDataClient.Server).IsEqualTo(defaultHost)
		ft.AssertThat(client.myDataClient.Client).IsEqualTo(&http.Client{Timeout: defaultTimeout})
		ft.AssertThatSlice(client.myDataClient.RequestEditors).HasSize(1)
	})
}

func TestClient_RequestDocs(t *testing.T) {
	ft := assert.NewFluentT(t)
	client, err := NewClient(&Config{UserID: "123", SubscriptionKey: "321"})
	ft.AssertThat(err).IsNil()

	runRequestTest(t, client, client.RequestDocs, "RequestDocs", "getting docs")
}

func TestClient_RequestTransmittedDocs(t *testing.T) {
	ft := assert.NewFluentT(t)
	client, err := NewClient(&Config{UserID: "123", SubscriptionKey: "321"})
	ft.AssertThat(err).IsNil()

	runRequestTest(t, client, client.RequestTransmittedDocs, "RequestTransmittedDocs", "getting transmitted docs")
}

func runRequestTest(t *testing.T, client *Client, f func() (*types.RequestedDoc, error), path, errPrefix string) {
	t.Helper()
	ft := assert.NewFluentT(t)
	t.Run("should fail if api call fails", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{}, errors.New("host not found")))

		docs, err := f()
		ft.AssertThat(docs).IsNil()
		assert.ThatError(t, err).HasExactMessage(
			fmt.Sprintf("%s: api client failed: Get \"https://mydata-dev.azure-api.net/%s?mark=\": host not found", errPrefix, path))
	})

	t.Run("should fail if api call returns an unexpected error", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(bytes.NewBuffer([]byte("bad request"))),
			}, nil))

		docs, err := f()
		ft.AssertThat(docs).IsNil()
		assert.ThatError(t, err).HasExactMessage(
			fmt.Sprintf("%s: unexpected status code (400): bad request", errPrefix))
	})

	t.Run("should fail if api call returns a malformed XML doc", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBuffer([]byte("{}"))),
			}, nil))

		docs, err := f()
		ft.AssertThat(docs).IsNil()
		assert.ThatError(t, err).HasExactMessage(
			fmt.Sprintf("%s: xml parser failed EOF", errPrefix))
	})

	t.Run("should succeed", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBuffer(response)),
			}, nil))

		docs, err := f()
		ft.AssertThat(err).IsNil()
		ft.AssertThatSlice(docs.InvoicesDoc.Invoices).HasSize(1).Contains(expectedInvoice())
	})
}

func expectedInvoice() types.Invoice {
	return types.Invoice{
		UID:  "123",
		Mark: "mark123",
		Issuer: types.Issuer{
			VatNumber: "123456789",
			Country:   "GR",
			Branch:    "0",
		},
		Counterpart: types.Counterpart{
			VatNumber: "987654321",
			Country:   "GR",
			Branch:    "0",
			Name:      "",
			Address: types.Address{
				Street:     "Some address",
				PostalCode: "11111",
				City:       "Athens",
			},
		},
		InvoiceHeader: types.InvoiceHeader{
			Series:               "~",
			Aa:                   "23",
			IssueDate:            "2021-02-10",
			InvoiceType:          "2.1",
			VatPaymentSuspension: "false",
			Currency:             "EUR",
		},
		PaymentMethods: types.PaymentMethods{
			PaymentMethodDetails: types.PaymentMethodDetails{
				Type:   "5",
				Amount: "62",
			},
		},
		InvoiceDetails: types.InvoiceDetails{
			LineNumber:  "1",
			NetValue:    "50",
			VatCategory: "1",
			VatAmount:   "12",
		},
		InvoiceSummary: types.InvoiceSummary{
			TotalNetValue:         "50",
			TotalVatAmount:        "12",
			TotalWithheldAmount:   "0",
			TotalFeesAmount:       "0",
			TotalStampDutyAmount:  "0",
			TotalOtherTaxesAmount: "0",
			TotalDeductionsAmount: "0",
			TotalGrossValue:       "62",
			IncomeClassification:  types.IncomeClassification{},
		},
	}
}
