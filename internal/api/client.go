// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

const (
	ApiKeyHeaderScopes = "apiKeyHeader.Scopes"
	ApiKeyQueryScopes  = "apiKeyQuery.Scopes"
)

// PostCancelinvoiceParams defines parameters for PostCancelinvoice.
type PostCancelinvoiceParams struct {
	// Μοναδικός αριθμός καταχώρησης
	// παραστατικού προς ακύρωση
	Mark PostCancelinvoiceParamsMark `json:"mark"`

	// User Id
	AadeUserId PostCancelinvoiceParamsAadeUserId `json:"aade-user-id"`
}

// PostCancelinvoiceParamsMark defines parameters for PostCancelinvoice.
type PostCancelinvoiceParamsMark string

// PostCancelinvoiceParamsAadeUserId defines parameters for PostCancelinvoice.
type PostCancelinvoiceParamsAadeUserId string

// GetRequestdocsParams defines parameters for GetRequestdocs.
type GetRequestdocsParams struct {
	// Μοναδικός Αριθμός Καταχώρησης
	Mark *GetRequestdocsParamsMark `json:"mark,omitempty"`

	// Παράμετρος για την τμηματική λήψη των αποτελεσμάτων
	NextPartitionKey *GetRequestdocsParamsNextPartitionKey `json:"nextPartitionKey,omitempty"`

	// Παράμετρος για την τμηματική λήψη των αποτελεσμάτων
	NextRowKey *GetRequestdocsParamsNextRowKey `json:"nextRowKey,omitempty"`

	// aade-user-id
	AadeUserId GetRequestdocsParamsAadeUserId `json:"aade-user-id"`
}

// GetRequestdocsParamsMark defines parameters for GetRequestdocs.
type GetRequestdocsParamsMark string

// GetRequestdocsParamsNextPartitionKey defines parameters for GetRequestdocs.
type GetRequestdocsParamsNextPartitionKey string

// GetRequestdocsParamsNextRowKey defines parameters for GetRequestdocs.
type GetRequestdocsParamsNextRowKey string

// GetRequestdocsParamsAadeUserId defines parameters for GetRequestdocs.
type GetRequestdocsParamsAadeUserId string

// GetRequesttransmitteddocsParams defines parameters for GetRequesttransmitteddocs.
type GetRequesttransmitteddocsParams struct {
	// Μοναδικός Αριθμός Καταχώρησης
	Mark GetRequesttransmitteddocsParamsMark `json:"mark"`

	// Παράμετρος για την τμηματική λήψη των αποτελεσμάτων
	NextPartitionKey *GetRequesttransmitteddocsParamsNextPartitionKey `json:"nextPartitionKey,omitempty"`

	// Παράμετρος για την τμηματική λήψη των αποτελεσμάτων
	NextRowKey *GetRequesttransmitteddocsParamsNextRowKey `json:"nextRowKey,omitempty"`

	// User Id
	AadeUserId GetRequesttransmitteddocsParamsAadeUserId `json:"aade-user-id"`
}

// GetRequesttransmitteddocsParamsMark defines parameters for GetRequesttransmitteddocs.
type GetRequesttransmitteddocsParamsMark string

// GetRequesttransmitteddocsParamsNextPartitionKey defines parameters for GetRequesttransmitteddocs.
type GetRequesttransmitteddocsParamsNextPartitionKey string

// GetRequesttransmitteddocsParamsNextRowKey defines parameters for GetRequesttransmitteddocs.
type GetRequesttransmitteddocsParamsNextRowKey string

// GetRequesttransmitteddocsParamsAadeUserId defines parameters for GetRequesttransmitteddocs.
type GetRequesttransmitteddocsParamsAadeUserId string

// PostSendexpensesclassificationParams defines parameters for PostSendexpensesclassification.
type PostSendexpensesclassificationParams struct {
	// User Id
	AadeUserId PostSendexpensesclassificationParamsAadeUserId `json:"{aade-user-id}"`
}

// PostSendexpensesclassificationParamsAadeUserId defines parameters for PostSendexpensesclassification.
type PostSendexpensesclassificationParamsAadeUserId string

// PostSendincomeclassificationParams defines parameters for PostSendincomeclassification.
type PostSendincomeclassificationParams struct {
	// User Id
	AadeUserId PostSendincomeclassificationParamsAadeUserId `json:"aade-user-id"`
}

// PostSendincomeclassificationParamsAadeUserId defines parameters for PostSendincomeclassification.
type PostSendincomeclassificationParamsAadeUserId string

// PostSendinvoicesParams defines parameters for PostSendinvoices.
type PostSendinvoicesParams struct {
	// User Id
	AadeUserId PostSendinvoicesParamsAadeUserId `json:"aade-user-id"`
}

// PostSendinvoicesParamsAadeUserId defines parameters for PostSendinvoices.
type PostSendinvoicesParamsAadeUserId string

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// PostCancelinvoice request
	PostCancelinvoice(ctx context.Context, params *PostCancelinvoiceParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetRequestdocs request
	GetRequestdocs(ctx context.Context, params *GetRequestdocsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetRequesttransmitteddocs request
	GetRequesttransmitteddocs(ctx context.Context, params *GetRequesttransmitteddocsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostSendexpensesclassification request
	PostSendexpensesclassification(ctx context.Context, params *PostSendexpensesclassificationParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostSendincomeclassification request
	PostSendincomeclassification(ctx context.Context, params *PostSendincomeclassificationParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostSendinvoices request
	PostSendinvoices(ctx context.Context, params *PostSendinvoicesParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) PostCancelinvoice(ctx context.Context, params *PostCancelinvoiceParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostCancelinvoiceRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetRequestdocs(ctx context.Context, params *GetRequestdocsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRequestdocsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetRequesttransmitteddocs(ctx context.Context, params *GetRequesttransmitteddocsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRequesttransmitteddocsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostSendexpensesclassification(ctx context.Context, params *PostSendexpensesclassificationParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostSendexpensesclassificationRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostSendincomeclassification(ctx context.Context, params *PostSendincomeclassificationParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostSendincomeclassificationRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostSendinvoices(ctx context.Context, params *PostSendinvoicesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostSendinvoicesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewPostCancelinvoiceRequest generates requests for PostCancelinvoice
func NewPostCancelinvoiceRequest(server string, params *PostCancelinvoiceParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/CancelInvoice")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "mark", runtime.ParamLocationQuery, params.Mark); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var headerParam0 string

	headerParam0, err = runtime.StyleParamWithLocation("simple", false, "aade-user-id", runtime.ParamLocationHeader, params.AadeUserId)
	if err != nil {
		return nil, err
	}

	req.Header.Set("aade-user-id", headerParam0)

	return req, nil
}

// NewGetRequestdocsRequest generates requests for GetRequestdocs
func NewGetRequestdocsRequest(server string, params *GetRequestdocsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/RequestDocs")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Mark != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "mark", runtime.ParamLocationQuery, *params.Mark); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.NextPartitionKey != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "nextPartitionKey", runtime.ParamLocationQuery, *params.NextPartitionKey); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.NextRowKey != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "nextRowKey", runtime.ParamLocationQuery, *params.NextRowKey); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var headerParam0 string

	headerParam0, err = runtime.StyleParamWithLocation("simple", false, "aade-user-id", runtime.ParamLocationHeader, params.AadeUserId)
	if err != nil {
		return nil, err
	}

	req.Header.Set("aade-user-id", headerParam0)

	return req, nil
}

// NewGetRequesttransmitteddocsRequest generates requests for GetRequesttransmitteddocs
func NewGetRequesttransmitteddocsRequest(server string, params *GetRequesttransmitteddocsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/RequestTransmittedDocs")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "mark", runtime.ParamLocationQuery, params.Mark); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.NextPartitionKey != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "nextPartitionKey", runtime.ParamLocationQuery, *params.NextPartitionKey); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.NextRowKey != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "nextRowKey", runtime.ParamLocationQuery, *params.NextRowKey); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var headerParam0 string

	headerParam0, err = runtime.StyleParamWithLocation("simple", false, "aade-user-id", runtime.ParamLocationHeader, params.AadeUserId)
	if err != nil {
		return nil, err
	}

	req.Header.Set("aade-user-id", headerParam0)

	return req, nil
}

// NewPostSendexpensesclassificationRequest generates requests for PostSendexpensesclassification
func NewPostSendexpensesclassificationRequest(server string, params *PostSendexpensesclassificationParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/SendExpensesClassification")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var headerParam0 string

	headerParam0, err = runtime.StyleParamWithLocation("simple", false, "{aade-user-id}", runtime.ParamLocationHeader, params.AadeUserId)
	if err != nil {
		return nil, err
	}

	req.Header.Set("{aade-user-id}", headerParam0)

	return req, nil
}

// NewPostSendincomeclassificationRequest generates requests for PostSendincomeclassification
func NewPostSendincomeclassificationRequest(server string, params *PostSendincomeclassificationParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/SendIncomeClassification")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var headerParam0 string

	headerParam0, err = runtime.StyleParamWithLocation("simple", false, "aade-user-id", runtime.ParamLocationHeader, params.AadeUserId)
	if err != nil {
		return nil, err
	}

	req.Header.Set("aade-user-id", headerParam0)

	return req, nil
}

// NewPostSendinvoicesRequest generates requests for PostSendinvoices
func NewPostSendinvoicesRequest(server string, params *PostSendinvoicesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/SendInvoices")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var headerParam0 string

	headerParam0, err = runtime.StyleParamWithLocation("simple", false, "aade-user-id", runtime.ParamLocationHeader, params.AadeUserId)
	if err != nil {
		return nil, err
	}

	req.Header.Set("aade-user-id", headerParam0)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// PostCancelinvoice request
	PostCancelinvoiceWithResponse(ctx context.Context, params *PostCancelinvoiceParams, reqEditors ...RequestEditorFn) (*PostCancelinvoiceResponse, error)

	// GetRequestdocs request
	GetRequestdocsWithResponse(ctx context.Context, params *GetRequestdocsParams, reqEditors ...RequestEditorFn) (*GetRequestdocsResponse, error)

	// GetRequesttransmitteddocs request
	GetRequesttransmitteddocsWithResponse(ctx context.Context, params *GetRequesttransmitteddocsParams, reqEditors ...RequestEditorFn) (*GetRequesttransmitteddocsResponse, error)

	// PostSendexpensesclassification request
	PostSendexpensesclassificationWithResponse(ctx context.Context, params *PostSendexpensesclassificationParams, reqEditors ...RequestEditorFn) (*PostSendexpensesclassificationResponse, error)

	// PostSendincomeclassification request
	PostSendincomeclassificationWithResponse(ctx context.Context, params *PostSendincomeclassificationParams, reqEditors ...RequestEditorFn) (*PostSendincomeclassificationResponse, error)

	// PostSendinvoices request
	PostSendinvoicesWithResponse(ctx context.Context, params *PostSendinvoicesParams, reqEditors ...RequestEditorFn) (*PostSendinvoicesResponse, error)
}

type PostCancelinvoiceResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostCancelinvoiceResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostCancelinvoiceResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRequestdocsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetRequestdocsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetRequestdocsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRequesttransmitteddocsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetRequesttransmitteddocsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetRequesttransmitteddocsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostSendexpensesclassificationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostSendexpensesclassificationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostSendexpensesclassificationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostSendincomeclassificationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostSendincomeclassificationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostSendincomeclassificationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostSendinvoicesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostSendinvoicesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostSendinvoicesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostCancelinvoiceWithResponse request returning *PostCancelinvoiceResponse
func (c *ClientWithResponses) PostCancelinvoiceWithResponse(ctx context.Context, params *PostCancelinvoiceParams, reqEditors ...RequestEditorFn) (*PostCancelinvoiceResponse, error) {
	rsp, err := c.PostCancelinvoice(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostCancelinvoiceResponse(rsp)
}

// GetRequestdocsWithResponse request returning *GetRequestdocsResponse
func (c *ClientWithResponses) GetRequestdocsWithResponse(ctx context.Context, params *GetRequestdocsParams, reqEditors ...RequestEditorFn) (*GetRequestdocsResponse, error) {
	rsp, err := c.GetRequestdocs(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRequestdocsResponse(rsp)
}

// GetRequesttransmitteddocsWithResponse request returning *GetRequesttransmitteddocsResponse
func (c *ClientWithResponses) GetRequesttransmitteddocsWithResponse(ctx context.Context, params *GetRequesttransmitteddocsParams, reqEditors ...RequestEditorFn) (*GetRequesttransmitteddocsResponse, error) {
	rsp, err := c.GetRequesttransmitteddocs(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRequesttransmitteddocsResponse(rsp)
}

// PostSendexpensesclassificationWithResponse request returning *PostSendexpensesclassificationResponse
func (c *ClientWithResponses) PostSendexpensesclassificationWithResponse(ctx context.Context, params *PostSendexpensesclassificationParams, reqEditors ...RequestEditorFn) (*PostSendexpensesclassificationResponse, error) {
	rsp, err := c.PostSendexpensesclassification(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostSendexpensesclassificationResponse(rsp)
}

// PostSendincomeclassificationWithResponse request returning *PostSendincomeclassificationResponse
func (c *ClientWithResponses) PostSendincomeclassificationWithResponse(ctx context.Context, params *PostSendincomeclassificationParams, reqEditors ...RequestEditorFn) (*PostSendincomeclassificationResponse, error) {
	rsp, err := c.PostSendincomeclassification(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostSendincomeclassificationResponse(rsp)
}

// PostSendinvoicesWithResponse request returning *PostSendinvoicesResponse
func (c *ClientWithResponses) PostSendinvoicesWithResponse(ctx context.Context, params *PostSendinvoicesParams, reqEditors ...RequestEditorFn) (*PostSendinvoicesResponse, error) {
	rsp, err := c.PostSendinvoices(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostSendinvoicesResponse(rsp)
}

// ParsePostCancelinvoiceResponse parses an HTTP response from a PostCancelinvoiceWithResponse call
func ParsePostCancelinvoiceResponse(rsp *http.Response) (*PostCancelinvoiceResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostCancelinvoiceResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetRequestdocsResponse parses an HTTP response from a GetRequestdocsWithResponse call
func ParseGetRequestdocsResponse(rsp *http.Response) (*GetRequestdocsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetRequestdocsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetRequesttransmitteddocsResponse parses an HTTP response from a GetRequesttransmitteddocsWithResponse call
func ParseGetRequesttransmitteddocsResponse(rsp *http.Response) (*GetRequesttransmitteddocsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetRequesttransmitteddocsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParsePostSendexpensesclassificationResponse parses an HTTP response from a PostSendexpensesclassificationWithResponse call
func ParsePostSendexpensesclassificationResponse(rsp *http.Response) (*PostSendexpensesclassificationResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostSendexpensesclassificationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParsePostSendincomeclassificationResponse parses an HTTP response from a PostSendincomeclassificationWithResponse call
func ParsePostSendincomeclassificationResponse(rsp *http.Response) (*PostSendincomeclassificationResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostSendincomeclassificationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParsePostSendinvoicesResponse parses an HTTP response from a PostSendinvoicesWithResponse call
func ParsePostSendinvoicesResponse(rsp *http.Response) (*PostSendinvoicesResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostSendinvoicesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
