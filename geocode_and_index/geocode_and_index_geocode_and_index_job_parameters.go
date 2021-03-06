// Code generated by go-swagger; DO NOT EDIT.

package geocode_and_index

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

// NewGeocodeAndIndexGeocodeAndIndexJobParams creates a new GeocodeAndIndexGeocodeAndIndexJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGeocodeAndIndexGeocodeAndIndexJobParams() *GeocodeAndIndexGeocodeAndIndexJobParams {
	return &GeocodeAndIndexGeocodeAndIndexJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGeocodeAndIndexGeocodeAndIndexJobParamsWithTimeout creates a new GeocodeAndIndexGeocodeAndIndexJobParams object
// with the ability to set a timeout on a request.
func NewGeocodeAndIndexGeocodeAndIndexJobParamsWithTimeout(timeout time.Duration) *GeocodeAndIndexGeocodeAndIndexJobParams {
	return &GeocodeAndIndexGeocodeAndIndexJobParams{
		timeout: timeout,
	}
}

// NewGeocodeAndIndexGeocodeAndIndexJobParamsWithContext creates a new GeocodeAndIndexGeocodeAndIndexJobParams object
// with the ability to set a context for a request.
func NewGeocodeAndIndexGeocodeAndIndexJobParamsWithContext(ctx context.Context) *GeocodeAndIndexGeocodeAndIndexJobParams {
	return &GeocodeAndIndexGeocodeAndIndexJobParams{
		Context: ctx,
	}
}

// NewGeocodeAndIndexGeocodeAndIndexJobParamsWithHTTPClient creates a new GeocodeAndIndexGeocodeAndIndexJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewGeocodeAndIndexGeocodeAndIndexJobParamsWithHTTPClient(client *http.Client) *GeocodeAndIndexGeocodeAndIndexJobParams {
	return &GeocodeAndIndexGeocodeAndIndexJobParams{
		HTTPClient: client,
	}
}

/* GeocodeAndIndexGeocodeAndIndexJobParams contains all the parameters to send to the API endpoint
   for the geocode and index geocode and index job operation.

   Typically these are written to a http.Request.
*/
type GeocodeAndIndexGeocodeAndIndexJobParams struct {

	/* SovrenAccountID.

	   Provided Account Id
	*/
	SovrenAccountID string

	/* SovrenServiceKey.

	   Provided Service Key
	*/
	SovrenServiceKey string

	// RequestModel.
	RequestModel *go_sovren_models.SovrenSaasDomainV10GeocodeAndIndexJobRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the geocode and index geocode and index job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) WithDefaults() *GeocodeAndIndexGeocodeAndIndexJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the geocode and index geocode and index job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the geocode and index geocode and index job params
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) WithTimeout(timeout time.Duration) *GeocodeAndIndexGeocodeAndIndexJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the geocode and index geocode and index job params
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the geocode and index geocode and index job params
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) WithContext(ctx context.Context) *GeocodeAndIndexGeocodeAndIndexJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the geocode and index geocode and index job params
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the geocode and index geocode and index job params
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) WithHTTPClient(client *http.Client) *GeocodeAndIndexGeocodeAndIndexJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the geocode and index geocode and index job params
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSovrenAccountID adds the sovrenAccountID to the geocode and index geocode and index job params
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) WithSovrenAccountID(sovrenAccountID string) *GeocodeAndIndexGeocodeAndIndexJobParams {
	o.SetSovrenAccountID(sovrenAccountID)
	return o
}

// SetSovrenAccountID adds the sovrenAccountId to the geocode and index geocode and index job params
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) SetSovrenAccountID(sovrenAccountID string) {
	o.SovrenAccountID = sovrenAccountID
}

// WithSovrenServiceKey adds the sovrenServiceKey to the geocode and index geocode and index job params
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) WithSovrenServiceKey(sovrenServiceKey string) *GeocodeAndIndexGeocodeAndIndexJobParams {
	o.SetSovrenServiceKey(sovrenServiceKey)
	return o
}

// SetSovrenServiceKey adds the sovrenServiceKey to the geocode and index geocode and index job params
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) SetSovrenServiceKey(sovrenServiceKey string) {
	o.SovrenServiceKey = sovrenServiceKey
}

// WithRequestModel adds the requestModel to the geocode and index geocode and index job params
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) WithRequestModel(requestModel *go_sovren_models.SovrenSaasDomainV10GeocodeAndIndexJobRequest) *GeocodeAndIndexGeocodeAndIndexJobParams {
	o.SetRequestModel(requestModel)
	return o
}

// SetRequestModel adds the requestModel to the geocode and index geocode and index job params
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) SetRequestModel(requestModel *go_sovren_models.SovrenSaasDomainV10GeocodeAndIndexJobRequest) {
	o.RequestModel = requestModel
}

// WriteToRequest writes these params to a swagger request
func (o *GeocodeAndIndexGeocodeAndIndexJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.RequestModel != nil {
		if err := r.SetBodyParam(o.RequestModel); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
