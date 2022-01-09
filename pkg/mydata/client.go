package mydata

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ppapapetrou76/go-mydata-aade/internal/api"
	"github.com/ppapapetrou76/go-mydata-aade/pkg/types"
)

const (
	defaultTimeout = time.Second * 30
	defaultHost    = "https://mydata-dev.azure-api.net/"
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

func (c Client) RequestTransmittedDocs() (*types.RequestedDoc, error) {
	const errPrefix = "getting transmitted docs"

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.myDataClient.GetRequesttransmitteddocs(ctx, &api.GetRequesttransmitteddocsParams{Mark: ""})
	if err != nil {
		return nil, fmt.Errorf("%s: api client failed: %w", errPrefix, err)
	}

	return c.readResponseBody(resp, errPrefix)
}

func (c Client) RequestDocs() (*types.RequestedDoc, error) {
	const errPrefix = "getting docs"
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	mark := api.GetRequestdocsParamsMark("")

	resp, err := c.myDataClient.GetRequestdocs(ctx, &api.GetRequestdocsParams{Mark: &mark})
	if err != nil {
		return nil, fmt.Errorf("%s: api client failed: %w", errPrefix, err)
	}

	return c.readResponseBody(resp, errPrefix)
}

func (c Client) readResponseBody(resp *http.Response, errPrefix string) (*types.RequestedDoc, error) {
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: reading response body failed: %w", errPrefix, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: unexpected status code (%d): %s", errPrefix, resp.StatusCode, respBody)
	}
	docs := &types.RequestedDoc{}
	if err := xml.Unmarshal(respBody, docs); err != nil {
		return nil, fmt.Errorf("%s: xml parser failed %w", errPrefix, err)
	}

	return docs, nil
}
