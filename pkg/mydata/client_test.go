package mydata

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
	"github.com/ppapapetrou76/go-utils/pkg/types"

	"github.com/ppapapetrou76/go-mydata-aade/internal/httpmock"
	"github.com/ppapapetrou76/go-mydata-aade/pkg/models"
)

//go:embed testdata/request_response.xml
var response []byte

//go:embed testdata/cancel_success.xml
var cancelSuccessResp []byte

//go:embed testdata/cancel_error.xml
var cancelErrorResp []byte

//go:embed testdata/send_success.xml
var sendSuccessResp []byte

//go:embed testdata/send_errors.xml
var sendErrorResp []byte

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

func TestBodyProvider_Intercept(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		body := &models.InvoicesDoc{Invoices: []models.Invoice{
			expectedInvoice(),
		}}
		bodyProvider := BodyProvider{
			Body: body,
		}
		req := &http.Request{}
		err := bodyProvider.Intercept(context.Background(), req)

		assert.ThatError(t, err).IsNil()
		defer req.Body.Close()
		rawBody, err := io.ReadAll(req.Body)
		assert.ThatError(t, err).IsNil()

		expectedBody := &models.InvoicesDoc{}
		err = xml.Unmarshal(rawBody, expectedBody)
		assert.ThatError(t, err).IsNil()

		assert.That(t, expectedBody).IsEqualTo(body)
	})
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

func TestClient_CancelDoc(t *testing.T) {
	ft := assert.NewFluentT(t)
	client, err := NewClient(&Config{UserID: "123", SubscriptionKey: "321"})
	ft.AssertThat(err).IsNil()
	t.Run("should fail if api call fails", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{}, errors.New("host not found")))

		docs, err := client.CancelDoc(123)
		ft.AssertThat(docs).IsNil()
		assert.ThatError(t, err).HasExactMessage("canceling doc: api client failed: Post \"https://mydata-dev.azure-api.net/CancelInvoice?mark=123\": host not found")
	})

	t.Run("should fail if api call returns an unexpected error", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(bytes.NewBuffer([]byte("bad request"))),
			}, nil))

		docs, err := client.CancelDoc(123)
		ft.AssertThat(docs).IsNil()
		assert.ThatError(t, err).HasExactMessage("canceling doc: unexpected status code (400): bad request")
	})

	t.Run("should fail if api call returns a malformed XML doc", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBuffer([]byte("{}"))),
			}, nil))

		docs, err := client.CancelDoc(123)
		ft.AssertThat(docs).IsNil()
		assert.ThatError(t, err).HasExactMessage("canceling doc: xml parser failed EOF")
	})

	t.Run("should succeed with no errors", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBuffer(cancelSuccessResp)),
			}, nil))

		resp, err := client.CancelDoc(123)
		ft.AssertThat(err).IsNil()
		ft.AssertThat(resp).IsNotNil().IsEqualTo([]*models.Response{{
			StatusCode:       "Success",
			CancellationMark: types.NewUint64(400001868549300),
		}})
	})

	t.Run("should succeed with errors", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBuffer(cancelErrorResp)),
			}, nil))

		resp, err := client.CancelDoc(123)
		assert.ThatError(t, err).
			IsNotNil().
			HasExactMessage("canceling doc Code:251 - Message:Invoice with MARK 400001868549153 cannot be canceled because of being already canceled")
		ft.AssertThat(resp).IsNotNil().IsEqualTo([]*models.Response{{
			StatusCode: "ValidationError",
			Errors: models.Errors{Error: []models.Error{
				{
					Message: "Invoice with MARK 400001868549153 cannot be canceled because of being already canceled",
					Code:    "251",
				},
			}},
		}})
	})
}

func TestClient_SendDocs(t *testing.T) {
	ft := assert.NewFluentT(t)
	client, err := NewClient(&Config{UserID: "123", SubscriptionKey: "321"})
	ft.AssertThat(err).IsNil()
	t.Run("should fail if api call fails", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{}, errors.New("host not found")))

		docs, err := client.SendDocs(&models.InvoicesDoc{})
		ft.AssertThat(docs).IsNil()
		assert.ThatError(t, err).HasExactMessage("sending doc: Post \"https://mydata-dev.azure-api.net/SendInvoices\": host not found")
	})

	t.Run("should fail if api call returns an unexpected error", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(bytes.NewBuffer([]byte("bad request"))),
			}, nil))

		docs, err := client.SendDocs(&models.InvoicesDoc{})
		ft.AssertThat(docs).IsNil()
		assert.ThatError(t, err).HasExactMessage("sending doc: unexpected status code (400): bad request")
	})

	t.Run("should fail if api call returns a malformed XML doc", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBuffer([]byte("{}"))),
			}, nil))

		docs, err := client.SendDocs(&models.InvoicesDoc{})
		ft.AssertThat(docs).IsNil()
		assert.ThatError(t, err).HasExactMessage("sending doc: xml parser failed EOF")
	})

	t.Run("should succeed with no errors", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBuffer(sendSuccessResp)),
			}, nil))

		resp, err := client.SendDocs(&models.InvoicesDoc{})
		ft.AssertThat(err).IsNil()
		ft.AssertThat(resp).IsNotNil().IsEqualTo([]*models.Response{
			{
				StatusCode:  "Success",
				Index:       types.NewUint(1),
				InvoiceUID:  types.NewString("148689AB8B26D37C5B3202E81E1C5FCE1B2EEA7B"),
				InvoiceMark: types.NewUint64(400001868549304),
			},
			{
				StatusCode:  "Success",
				Index:       types.NewUint(2),
				InvoiceUID:  types.NewString("72684BA069C4CF9CDEDCF3EBA30E40D0AF5890C4"),
				InvoiceMark: types.NewUint64(400001868549305),
			},
		})
	})

	t.Run("should succeed with errors", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBuffer(sendErrorResp)),
			}, nil))

		resp, err := client.SendDocs(&models.InvoicesDoc{})
		assert.ThatError(t, err).
			IsNotNil().
			HasExactMessage("sending doc 4 errors:[ Code:204 - Message:AA is mandatory for this invoice type, " +
				"Code:212 - Message:AA element must be number (positive) for issuer from Greece, " +
				"Code:204 - Message:AA is mandatory for this invoice type, " +
				"Code:212 - Message:AA element must be number (positive) for issuer from Greece ]")
		ft.AssertThat(resp).IsNotNil().IsEqualTo([]*models.Response{
			{
				StatusCode: "ValidationError",
				Index:      types.NewUint(1),
				Errors: models.Errors{Error: []models.Error{
					{Message: "AA is mandatory for this invoice type", Code: "204"},
					{Message: "AA element must be number (positive) for issuer from Greece", Code: "212"},
				}},
			},
			{
				StatusCode: "ValidationError",
				Index:      types.NewUint(2),
				Errors: models.Errors{Error: []models.Error{
					{Message: "AA is mandatory for this invoice type", Code: "204"},
					{Message: "AA element must be number (positive) for issuer from Greece", Code: "212"},
				}},
			},
		})
	})
}

func runRequestTest(t *testing.T, client *Client, f func(mark string, nextPartitionKey, nextRowKey *string) (*models.RequestedDoc, error), path, errPrefix string) {
	t.Helper()
	ft := assert.NewFluentT(t)
	var mark string
	var nextPartitionKey, nextRowKey *string = nil, nil
	t.Run("should fail if api call fails", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{}, errors.New("host not found")))

		docs, err := f(mark, nextPartitionKey, nextRowKey)
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

		docs, err := f(mark, nextPartitionKey, nextRowKey)
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

		docs, err := f(mark, nextPartitionKey, nextRowKey)
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

		docs, err := f(mark, nextPartitionKey, nextRowKey)
		ft.AssertThat(err).IsNil()
		ft.AssertThatStruct(docs.InvoicesDoc.Invoices[0]).IsEqualTo(expectedInvoice())
		ft.AssertThatSlice(docs.InvoicesDoc.Invoices).HasSize(1).Contains(expectedInvoice())
	})

	t.Run("should include next partition and row key in the URL", func(t *testing.T) {
		client.myDataClient.Client = httpmock.NewMockClient(
			httpmock.NewMockResponse(&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBuffer(response)),
			}, nil))

		want := map[string][]string{
			"mark":             {"123"},
			"nextPartitionKey": {"1"},
			"nextRowKey":       {"2"},
		}

		checkParameters := func(_ context.Context, req *http.Request) error {
			values := req.URL.Query()
			for k, got := range values {
				ft.AssertThat(got).IsEqualTo(want[k])
			}

			return nil
		}
		client.myDataClient.RequestEditors = append(client.myDataClient.RequestEditors, checkParameters)

		mark = "123"
		nextPartitionKeyValue := "1"
		nextRowKeyValue := "2"
		_, err := f(mark, &nextPartitionKeyValue, &nextRowKeyValue)
		ft.AssertThat(err).IsNil()
	})
}

func expectedInvoice() models.Invoice {
	return models.Invoice{
		UID:  types.NewString("123"),
		Mark: types.NewUint64(321),
		Issuer: &models.PartyType{
			VatNumber: "123456789",
			Country:   "GR",
			Branch:    0,
		},
		Counterpart: &models.PartyType{
			VatNumber: "987654321",
			Country:   "GR",
			Branch:    0,
			Address: &models.Address{
				Street:     types.NewString("Some address"),
				Number:     types.NewString("100"),
				PostalCode: "11111",
				City:       "Athens",
			},
		},
		InvoiceHeader: &models.InvoiceHeader{
			Series:               "~",
			Aa:                   "23",
			IssueDate:            "2021-02-10",
			InvoiceType:          "2.1",
			Currency:             "EUR",
			VatPaymentSuspension: types.NewBool(false),
		},
		PaymentMethods: &models.PaymentMethods{
			PaymentMethodDetails: &models.PaymentMethodDetails{
				Type:              5,
				Amount:            62.00,
				PaymentMethodInfo: "some payment info",
			},
		},
		InvoiceDetails: []*models.InvoiceDetails{
			{
				LineNumber:  1,
				NetValue:    50.00,
				VatCategory: 1,
				VatAmount:   12,
				IncomeClassification: []*models.IncomeClassificationType{
					{
						ClassificationType:     "E3_561_001",
						ClassificationCategory: "category1_2",
						Amount:                 15,
						ID:                     nil,
					},
					{
						ClassificationType:     "E3_561_001",
						ClassificationCategory: "category1_3",
						Amount:                 20,
						ID:                     nil,
					},
				},
			},
		},
		InvoiceSummary: &models.InvoiceSummary{
			TotalNetValue:         50,
			TotalVatAmount:        12,
			TotalWithheldAmount:   0,
			TotalFeesAmount:       0,
			TotalStampDutyAmount:  0,
			TotalOtherTaxesAmount: 0,
			TotalDeductionsAmount: 0,
			TotalGrossValue:       62,
		},
	}
}
