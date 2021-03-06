// Code generated by go-swagger; DO NOT EDIT.

package geocoder

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

// NewGeocoderGeocodeJobParams creates a new GeocoderGeocodeJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGeocoderGeocodeJobParams() *GeocoderGeocodeJobParams {
	return &GeocoderGeocodeJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGeocoderGeocodeJobParamsWithTimeout creates a new GeocoderGeocodeJobParams object
// with the ability to set a timeout on a request.
func NewGeocoderGeocodeJobParamsWithTimeout(timeout time.Duration) *GeocoderGeocodeJobParams {
	return &GeocoderGeocodeJobParams{
		timeout: timeout,
	}
}

// NewGeocoderGeocodeJobParamsWithContext creates a new GeocoderGeocodeJobParams object
// with the ability to set a context for a request.
func NewGeocoderGeocodeJobParamsWithContext(ctx context.Context) *GeocoderGeocodeJobParams {
	return &GeocoderGeocodeJobParams{
		Context: ctx,
	}
}

// NewGeocoderGeocodeJobParamsWithHTTPClient creates a new GeocoderGeocodeJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewGeocoderGeocodeJobParamsWithHTTPClient(client *http.Client) *GeocoderGeocodeJobParams {
	return &GeocoderGeocodeJobParams{
		HTTPClient: client,
	}
}

/* GeocoderGeocodeJobParams contains all the parameters to send to the API endpoint
   for the geocoder geocode job operation.

   Typically these are written to a http.Request.
*/
type GeocoderGeocodeJobParams struct {

	/* SovrenAccountID.

	   Provided Account Id
	*/
	SovrenAccountID string

	/* SovrenServiceKey.

	   Provided Service Key
	*/
	SovrenServiceKey string

	/* RequestModel.

	   An object with a parsed resume and additional options.
	*/
	RequestModel *go_sovren_models.SovrenSaasDomainV10GeoCodeJobRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the geocoder geocode job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GeocoderGeocodeJobParams) WithDefaults() *GeocoderGeocodeJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the geocoder geocode job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GeocoderGeocodeJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the geocoder geocode job params
func (o *GeocoderGeocodeJobParams) WithTimeout(timeout time.Duration) *GeocoderGeocodeJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the geocoder geocode job params
func (o *GeocoderGeocodeJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the geocoder geocode job params
func (o *GeocoderGeocodeJobParams) WithContext(ctx context.Context) *GeocoderGeocodeJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the geocoder geocode job params
func (o *GeocoderGeocodeJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the geocoder geocode job params
func (o *GeocoderGeocodeJobParams) WithHTTPClient(client *http.Client) *GeocoderGeocodeJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the geocoder geocode job params
func (o *GeocoderGeocodeJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSovrenAccountID adds the sovrenAccountID to the geocoder geocode job params
func (o *GeocoderGeocodeJobParams) WithSovrenAccountID(sovrenAccountID string) *GeocoderGeocodeJobParams {
	o.SetSovrenAccountID(sovrenAccountID)
	return o
}

// SetSovrenAccountID adds the sovrenAccountId to the geocoder geocode job params
func (o *GeocoderGeocodeJobParams) SetSovrenAccountID(sovrenAccountID string) {
	o.SovrenAccountID = sovrenAccountID
}

// WithSovrenServiceKey adds the sovrenServiceKey to the geocoder geocode job params
func (o *GeocoderGeocodeJobParams) WithSovrenServiceKey(sovrenServiceKey string) *GeocoderGeocodeJobParams {
	o.SetSovrenServiceKey(sovrenServiceKey)
	return o
}

// SetSovrenServiceKey adds the sovrenServiceKey to the geocoder geocode job params
func (o *GeocoderGeocodeJobParams) SetSovrenServiceKey(sovrenServiceKey string) {
	o.SovrenServiceKey = sovrenServiceKey
}

// WithRequestModel adds the requestModel to the geocoder geocode job params
func (o *GeocoderGeocodeJobParams) WithRequestModel(requestModel *go_sovren_models.SovrenSaasDomainV10GeoCodeJobRequest) *GeocoderGeocodeJobParams {
	o.SetRequestModel(requestModel)
	return o
}

// SetRequestModel adds the requestModel to the geocoder geocode job params
func (o *GeocoderGeocodeJobParams) SetRequestModel(requestModel *go_sovren_models.SovrenSaasDomainV10GeoCodeJobRequest) {
	o.RequestModel = requestModel
}

// WriteToRequest writes these params to a swagger request
func (o *GeocoderGeocodeJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
