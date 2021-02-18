// Code generated by go-swagger; DO NOT EDIT.

package skills

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

// NewSkillsPostParams creates a new SkillsPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSkillsPostParams() *SkillsPostParams {
	return &SkillsPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSkillsPostParamsWithTimeout creates a new SkillsPostParams object
// with the ability to set a timeout on a request.
func NewSkillsPostParamsWithTimeout(timeout time.Duration) *SkillsPostParams {
	return &SkillsPostParams{
		timeout: timeout,
	}
}

// NewSkillsPostParamsWithContext creates a new SkillsPostParams object
// with the ability to set a context for a request.
func NewSkillsPostParamsWithContext(ctx context.Context) *SkillsPostParams {
	return &SkillsPostParams{
		Context: ctx,
	}
}

// NewSkillsPostParamsWithHTTPClient creates a new SkillsPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewSkillsPostParamsWithHTTPClient(client *http.Client) *SkillsPostParams {
	return &SkillsPostParams{
		HTTPClient: client,
	}
}

/* SkillsPostParams contains all the parameters to send to the API endpoint
   for the skills post operation.

   Typically these are written to a http.Request.
*/
type SkillsPostParams struct {

	/* SovrenAccountID.

	   Provided Account Id
	*/
	SovrenAccountID string

	/* SovrenServiceKey.

	   Provided Service Key
	*/
	SovrenServiceKey string

	/* Request.

	   The SetSkillRequest object.
	*/
	Request *go_sovren_models.SovrenSaasDomainCreateSkillRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the skills post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SkillsPostParams) WithDefaults() *SkillsPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the skills post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SkillsPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the skills post params
func (o *SkillsPostParams) WithTimeout(timeout time.Duration) *SkillsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the skills post params
func (o *SkillsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the skills post params
func (o *SkillsPostParams) WithContext(ctx context.Context) *SkillsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the skills post params
func (o *SkillsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the skills post params
func (o *SkillsPostParams) WithHTTPClient(client *http.Client) *SkillsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the skills post params
func (o *SkillsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSovrenAccountID adds the sovrenAccountID to the skills post params
func (o *SkillsPostParams) WithSovrenAccountID(sovrenAccountID string) *SkillsPostParams {
	o.SetSovrenAccountID(sovrenAccountID)
	return o
}

// SetSovrenAccountID adds the sovrenAccountId to the skills post params
func (o *SkillsPostParams) SetSovrenAccountID(sovrenAccountID string) {
	o.SovrenAccountID = sovrenAccountID
}

// WithSovrenServiceKey adds the sovrenServiceKey to the skills post params
func (o *SkillsPostParams) WithSovrenServiceKey(sovrenServiceKey string) *SkillsPostParams {
	o.SetSovrenServiceKey(sovrenServiceKey)
	return o
}

// SetSovrenServiceKey adds the sovrenServiceKey to the skills post params
func (o *SkillsPostParams) SetSovrenServiceKey(sovrenServiceKey string) {
	o.SovrenServiceKey = sovrenServiceKey
}

// WithRequest adds the request to the skills post params
func (o *SkillsPostParams) WithRequest(request *go_sovren_models.SovrenSaasDomainCreateSkillRequest) *SkillsPostParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the skills post params
func (o *SkillsPostParams) SetRequest(request *go_sovren_models.SovrenSaasDomainCreateSkillRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *SkillsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
