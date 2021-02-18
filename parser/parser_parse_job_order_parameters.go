// Code generated by go-swagger; DO NOT EDIT.

package parser

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

// NewParserParseJobOrderParams creates a new ParserParseJobOrderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewParserParseJobOrderParams() *ParserParseJobOrderParams {
	return &ParserParseJobOrderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewParserParseJobOrderParamsWithTimeout creates a new ParserParseJobOrderParams object
// with the ability to set a timeout on a request.
func NewParserParseJobOrderParamsWithTimeout(timeout time.Duration) *ParserParseJobOrderParams {
	return &ParserParseJobOrderParams{
		timeout: timeout,
	}
}

// NewParserParseJobOrderParamsWithContext creates a new ParserParseJobOrderParams object
// with the ability to set a context for a request.
func NewParserParseJobOrderParamsWithContext(ctx context.Context) *ParserParseJobOrderParams {
	return &ParserParseJobOrderParams{
		Context: ctx,
	}
}

// NewParserParseJobOrderParamsWithHTTPClient creates a new ParserParseJobOrderParams object
// with the ability to set a custom HTTPClient for a request.
func NewParserParseJobOrderParamsWithHTTPClient(client *http.Client) *ParserParseJobOrderParams {
	return &ParserParseJobOrderParams{
		HTTPClient: client,
	}
}

/* ParserParseJobOrderParams contains all the parameters to send to the API endpoint
   for the parser parse job order operation.

   Typically these are written to a http.Request.
*/
type ParserParseJobOrderParams struct {

	/* SovrenAccountID.

	   Provided Account Id
	*/
	SovrenAccountID string

	/* SovrenServiceKey.

	   Provided Service Key
	*/
	SovrenServiceKey string

	/* Request.

	   The ParseResumeRequest object.
	*/
	Request *go_sovren_models.SovrenSaasServicesRestDomainV10ParseJobOrderRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the parser parse job order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ParserParseJobOrderParams) WithDefaults() *ParserParseJobOrderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the parser parse job order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ParserParseJobOrderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the parser parse job order params
func (o *ParserParseJobOrderParams) WithTimeout(timeout time.Duration) *ParserParseJobOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the parser parse job order params
func (o *ParserParseJobOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the parser parse job order params
func (o *ParserParseJobOrderParams) WithContext(ctx context.Context) *ParserParseJobOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the parser parse job order params
func (o *ParserParseJobOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the parser parse job order params
func (o *ParserParseJobOrderParams) WithHTTPClient(client *http.Client) *ParserParseJobOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the parser parse job order params
func (o *ParserParseJobOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSovrenAccountID adds the sovrenAccountID to the parser parse job order params
func (o *ParserParseJobOrderParams) WithSovrenAccountID(sovrenAccountID string) *ParserParseJobOrderParams {
	o.SetSovrenAccountID(sovrenAccountID)
	return o
}

// SetSovrenAccountID adds the sovrenAccountId to the parser parse job order params
func (o *ParserParseJobOrderParams) SetSovrenAccountID(sovrenAccountID string) {
	o.SovrenAccountID = sovrenAccountID
}

// WithSovrenServiceKey adds the sovrenServiceKey to the parser parse job order params
func (o *ParserParseJobOrderParams) WithSovrenServiceKey(sovrenServiceKey string) *ParserParseJobOrderParams {
	o.SetSovrenServiceKey(sovrenServiceKey)
	return o
}

// SetSovrenServiceKey adds the sovrenServiceKey to the parser parse job order params
func (o *ParserParseJobOrderParams) SetSovrenServiceKey(sovrenServiceKey string) {
	o.SovrenServiceKey = sovrenServiceKey
}

// WithRequest adds the request to the parser parse job order params
func (o *ParserParseJobOrderParams) WithRequest(request *go_sovren_models.SovrenSaasServicesRestDomainV10ParseJobOrderRequest) *ParserParseJobOrderParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the parser parse job order params
func (o *ParserParseJobOrderParams) SetRequest(request *go_sovren_models.SovrenSaasServicesRestDomainV10ParseJobOrderRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *ParserParseJobOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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