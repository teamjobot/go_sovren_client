// Code generated by go-swagger; DO NOT EDIT.

package matcher

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

// NewMatcherJobParams creates a new MatcherJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMatcherJobParams() *MatcherJobParams {
	return &MatcherJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMatcherJobParamsWithTimeout creates a new MatcherJobParams object
// with the ability to set a timeout on a request.
func NewMatcherJobParamsWithTimeout(timeout time.Duration) *MatcherJobParams {
	return &MatcherJobParams{
		timeout: timeout,
	}
}

// NewMatcherJobParamsWithContext creates a new MatcherJobParams object
// with the ability to set a context for a request.
func NewMatcherJobParamsWithContext(ctx context.Context) *MatcherJobParams {
	return &MatcherJobParams{
		Context: ctx,
	}
}

// NewMatcherJobParamsWithHTTPClient creates a new MatcherJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewMatcherJobParamsWithHTTPClient(client *http.Client) *MatcherJobParams {
	return &MatcherJobParams{
		HTTPClient: client,
	}
}

/* MatcherJobParams contains all the parameters to send to the API endpoint
   for the matcher job operation.

   Typically these are written to a http.Request.
*/
type MatcherJobParams struct {

	/* SovrenAccountID.

	   Provided Account Id
	*/
	SovrenAccountID string

	/* SovrenServiceKey.

	   Provided Service Key
	*/
	SovrenServiceKey string

	// Request.
	Request *go_sovren_models.SovrenMatchJobRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the matcher job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MatcherJobParams) WithDefaults() *MatcherJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the matcher job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MatcherJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the matcher job params
func (o *MatcherJobParams) WithTimeout(timeout time.Duration) *MatcherJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the matcher job params
func (o *MatcherJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the matcher job params
func (o *MatcherJobParams) WithContext(ctx context.Context) *MatcherJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the matcher job params
func (o *MatcherJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the matcher job params
func (o *MatcherJobParams) WithHTTPClient(client *http.Client) *MatcherJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the matcher job params
func (o *MatcherJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSovrenAccountID adds the sovrenAccountID to the matcher job params
func (o *MatcherJobParams) WithSovrenAccountID(sovrenAccountID string) *MatcherJobParams {
	o.SetSovrenAccountID(sovrenAccountID)
	return o
}

// SetSovrenAccountID adds the sovrenAccountId to the matcher job params
func (o *MatcherJobParams) SetSovrenAccountID(sovrenAccountID string) {
	o.SovrenAccountID = sovrenAccountID
}

// WithSovrenServiceKey adds the sovrenServiceKey to the matcher job params
func (o *MatcherJobParams) WithSovrenServiceKey(sovrenServiceKey string) *MatcherJobParams {
	o.SetSovrenServiceKey(sovrenServiceKey)
	return o
}

// SetSovrenServiceKey adds the sovrenServiceKey to the matcher job params
func (o *MatcherJobParams) SetSovrenServiceKey(sovrenServiceKey string) {
	o.SovrenServiceKey = sovrenServiceKey
}

// WithRequest adds the request to the matcher job params
func (o *MatcherJobParams) WithRequest(request *go_sovren_models.SovrenMatchJobRequest) *MatcherJobParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the matcher job params
func (o *MatcherJobParams) SetRequest(request *go_sovren_models.SovrenMatchJobRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *MatcherJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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