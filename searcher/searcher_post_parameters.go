// Code generated by go-swagger; DO NOT EDIT.

package searcher

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/teamjobot/go_sovren_models"
)

// NewSearcherPostParams creates a new SearcherPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearcherPostParams() *SearcherPostParams {
	return &SearcherPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearcherPostParamsWithTimeout creates a new SearcherPostParams object
// with the ability to set a timeout on a request.
func NewSearcherPostParamsWithTimeout(timeout time.Duration) *SearcherPostParams {
	return &SearcherPostParams{
		timeout: timeout,
	}
}

// NewSearcherPostParamsWithContext creates a new SearcherPostParams object
// with the ability to set a context for a request.
func NewSearcherPostParamsWithContext(ctx context.Context) *SearcherPostParams {
	return &SearcherPostParams{
		Context: ctx,
	}
}

// NewSearcherPostParamsWithHTTPClient creates a new SearcherPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearcherPostParamsWithHTTPClient(client *http.Client) *SearcherPostParams {
	return &SearcherPostParams{
		HTTPClient: client,
	}
}

/* SearcherPostParams contains all the parameters to send to the API endpoint
   for the searcher post operation.

   Typically these are written to a http.Request.
*/
type SearcherPostParams struct {

	/* SovrenAccountID.

	   Provided Account Id
	*/
	SovrenAccountID string

	/* SovrenServiceKey.

	   Provided Service Key
	*/
	SovrenServiceKey string

	// Request.
	Request *go_sovren_models.SovrenV10SearchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the searcher post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearcherPostParams) WithDefaults() *SearcherPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the searcher post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearcherPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the searcher post params
func (o *SearcherPostParams) WithTimeout(timeout time.Duration) *SearcherPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the searcher post params
func (o *SearcherPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the searcher post params
func (o *SearcherPostParams) WithContext(ctx context.Context) *SearcherPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the searcher post params
func (o *SearcherPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the searcher post params
func (o *SearcherPostParams) WithHTTPClient(client *http.Client) *SearcherPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the searcher post params
func (o *SearcherPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSovrenAccountID adds the sovrenAccountID to the searcher post params
func (o *SearcherPostParams) WithSovrenAccountID(sovrenAccountID string) *SearcherPostParams {
	o.SetSovrenAccountID(sovrenAccountID)
	return o
}

// SetSovrenAccountID adds the sovrenAccountId to the searcher post params
func (o *SearcherPostParams) SetSovrenAccountID(sovrenAccountID string) {
	o.SovrenAccountID = sovrenAccountID
}

// WithSovrenServiceKey adds the sovrenServiceKey to the searcher post params
func (o *SearcherPostParams) WithSovrenServiceKey(sovrenServiceKey string) *SearcherPostParams {
	o.SetSovrenServiceKey(sovrenServiceKey)
	return o
}

// SetSovrenServiceKey adds the sovrenServiceKey to the searcher post params
func (o *SearcherPostParams) SetSovrenServiceKey(sovrenServiceKey string) {
	o.SovrenServiceKey = sovrenServiceKey
}

// WithRequest adds the request to the searcher post params
func (o *SearcherPostParams) WithRequest(request *go_sovren_models.SovrenV10SearchRequest) *SearcherPostParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the searcher post params
func (o *SearcherPostParams) SetRequest(request *go_sovren_models.SovrenV10SearchRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *SearcherPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Sovren-AccountId
	if err := r.SetHeaderParam("Sovren-AccountId", o.SovrenAccountID); err != nil {
		return err
	}

	// header param Sovren-ServiceKey
	if err := r.SetHeaderParam("Sovren-ServiceKey", o.SovrenServiceKey); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}