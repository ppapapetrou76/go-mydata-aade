package mydata

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/ppapapetrou76/go-mydata-aade/internal/api"
	"github.com/ppapapetrou76/go-mydata-aade/pkg/models"
)

const (
	defaultTimeout = time.Second * 30
	defaultHost    = "https://mydata-dev.azure-api.net/"
	xmlns          = "http://www.aade.gr/myDATA/invoice/v1.0"
)

// Client describes the myDATA Client.
type Client struct {
	myDataClient *api.Client
	timeout      time.Duration
}

// AuthHeaderProvider sends the required headers as part of the required authorization.
type AuthHeaderProvider struct {
	userID          string
	subscriptionKey string
}

// Intercept will attach the required headers to the request
// and ensures that the user-id and the subscription-key are set correctly.
func (s *AuthHeaderProvider) Intercept(_ context.Context, req *http.Request) error {
	req.Header.Set("aade-user-id", s.userID)
	req.Header.Set("Ocp-Apim-Subscription-Key", s.subscriptionKey)

	return nil
}

// BodyProvider attaches the request body to the request call.
type BodyProvider struct {
	Body *models.InvoicesDoc
}

// Intercept will attach the required body to the given request.
func (s *BodyProvider) Intercept(_ context.Context, req *http.Request) error {
	b, err := xml.Marshal(s.Body)
	if err != nil {
		return fmt.Errorf("BodyProvider failed to intercept: %w", err)
	}
	req.Body = io.NopCloser(bytes.NewBuffer(b))

	return nil
}

// NewClient returns a valid myDATA client or an error if the client cannot be constructed with the given configuration
// values.
func NewClient(config *Config) (*Client, error) {
	if config == nil {
		return nil, errors.New("api config is required")
	}
	if err := config.validate(); err != nil {
		return nil, err
	}
	config.defaults()

	authProvider := &AuthHeaderProvider{
		userID:          config.UserID,
		subscriptionKey: config.SubscriptionKey,
	}
	myDataClient, err := api.NewClient(config.Host,
		api.WithHTTPClient(&http.Client{Timeout: config.Timeout}),
		api.WithRequestEditorFn(authProvider.Intercept),
	)
	if err != nil {
		return nil, fmt.Errorf("new api client failed: %w", err)
	}

	return &Client{
		myDataClient: myDataClient,
		timeout:      config.Timeout,
	}, nil
}

// RequestTransmittedDocs returns the invoices,cancellations etc., issued by the entity associated to the
// authenticated user.
// If the api returns an error, or it is not accessible then the method returns an error with descriptive message.
func (c Client) RequestTransmittedDocs() (*models.RequestedDoc, error) {
	const errPrefix = "getting transmitted docs"

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.myDataClient.GetRequesttransmitteddocs(ctx, &api.GetRequesttransmitteddocsParams{Mark: ""})
	if err != nil {
		return nil, fmt.Errorf("%s: api client failed: %w", errPrefix, err)
	}

	requestDoc, err := c.readRequestedDoc(resp)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errPrefix, err)
	}

	return requestDoc, nil
}

// RequestDocs returns the invoices,cancellations etc., issued by third-party entities and are related to the entity
// associated to the authenticated user.
// If the api returns an error, or it is not accessible then the method returns an error with descriptive message.
func (c Client) RequestDocs() (*models.RequestedDoc, error) {
	const errPrefix = "getting docs"
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	mark := api.GetRequestdocsParamsMark("")

	resp, err := c.myDataClient.GetRequestdocs(ctx, &api.GetRequestdocsParams{Mark: &mark})
	if err != nil {
		return nil, fmt.Errorf("%s: api client failed: %w", errPrefix, err)
	}
	requestDoc, err := c.readRequestedDoc(resp)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errPrefix, err)
	}

	return requestDoc, nil
}

// CancelDoc cancels an issued invoice without issuing a new one.
// If the operation is successful, the response returns the status and a unique identification number of the
// processed cancellation.
// If the api returns an error, or it is not accessible then the method returns an error with descriptive message.
func (c Client) CancelDoc(mark uint64) ([]*models.Response, error) {
	const errPrefix = "canceling doc"
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.myDataClient.PostCancelinvoice(ctx, &api.PostCancelinvoiceParams{
		Mark: api.PostCancelinvoiceParamsMark(strconv.FormatUint(mark, 10)), //nolint:gomnd // we don't care
	})
	if err != nil {
		return nil, fmt.Errorf("%s: api client failed: %w", errPrefix, err)
	}

	responseDoc, err := c.readResponseDoc(resp)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errPrefix, err)
	}

	return responseDoc.Response, responseDoc.Errors(errPrefix) //nolint:wrapcheck // already prefixed
}

// SendDocs issues one or more invoices.
// If the operation is successful, the response returns for each processed invoice, the status, a unique identification
// number and the unique process number.
// If the api returns an error, or it is not accessible then the method returns an error with descriptive message.
func (c Client) SendDocs(invoiceDoc *models.InvoicesDoc) ([]*models.Response, error) {
	const errPrefix = "sending doc"
	invoiceDoc.Xmlns = xmlns

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	bodyProvider := BodyProvider{Body: invoiceDoc}
	resp, err := c.myDataClient.PostSendinvoices(ctx, &api.PostSendinvoicesParams{}, bodyProvider.Intercept)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errPrefix, err)
	}

	responseDoc, err := c.readResponseDoc(resp)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errPrefix, err)
	}

	return responseDoc.Response, responseDoc.Errors(errPrefix) //nolint:wrapcheck // already prefixed
}

func (c Client) readRequestedDoc(resp *http.Response) (*models.RequestedDoc, error) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code (%d): %s", resp.StatusCode, respBody)
	}
	docs := &models.RequestedDoc{}
	if err := xml.Unmarshal(respBody, docs); err != nil {
		return nil, fmt.Errorf("xml parser failed %w", err)
	}

	return docs, nil
}

func (c Client) readResponseDoc(resp *http.Response) (*models.ResponseDoc, error) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code (%d): %s", resp.StatusCode, respBody)
	}
	docs := &models.ResponseDoc{}
	if err := xml.Unmarshal(respBody, docs); err != nil {
		return nil, fmt.Errorf("xml parser failed %w", err)
	}

	return docs, nil
}
