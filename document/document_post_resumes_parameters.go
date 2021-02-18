// Code generated by go-swagger; DO NOT EDIT.

package document

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

// NewDocumentPostResumesParams creates a new DocumentPostResumesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDocumentPostResumesParams() *DocumentPostResumesParams {
	return &DocumentPostResumesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDocumentPostResumesParamsWithTimeout creates a new DocumentPostResumesParams object
// with the ability to set a timeout on a request.
func NewDocumentPostResumesParamsWithTimeout(timeout time.Duration) *DocumentPostResumesParams {
	return &DocumentPostResumesParams{
		timeout: timeout,
	}
}

// NewDocumentPostResumesParamsWithContext creates a new DocumentPostResumesParams object
// with the ability to set a context for a request.
func NewDocumentPostResumesParamsWithContext(ctx context.Context) *DocumentPostResumesParams {
	return &DocumentPostResumesParams{
		Context: ctx,
	}
}

// NewDocumentPostResumesParamsWithHTTPClient creates a new DocumentPostResumesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDocumentPostResumesParamsWithHTTPClient(client *http.Client) *DocumentPostResumesParams {
	return &DocumentPostResumesParams{
		HTTPClient: client,
	}
}

/* DocumentPostResumesParams contains all the parameters to send to the API endpoint
   for the document post resumes operation.

   Typically these are written to a http.Request.
*/
type DocumentPostResumesParams struct {

	/* SovrenAccountID.

	   Provided Account Id
	*/
	SovrenAccountID string

	/* SovrenServiceKey.

	   Provided Service Key
	*/
	SovrenServiceKey string

	// IndexID.
	IndexID string

	// Request.
	Request *go_sovren_models.SovrenSaasDomainV10AddMultipleResumesToIndexRequestModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the document post resumes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DocumentPostResumesParams) WithDefaults() *DocumentPostResumesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the document post resumes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DocumentPostResumesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the document post resumes params
func (o *DocumentPostResumesParams) WithTimeout(timeout time.Duration) *DocumentPostResumesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the document post resumes params
func (o *DocumentPostResumesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the document post resumes params
func (o *DocumentPostResumesParams) WithContext(ctx context.Context) *DocumentPostResumesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the document post resumes params
func (o *DocumentPostResumesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the document post resumes params
func (o *DocumentPostResumesParams) WithHTTPClient(client *http.Client) *DocumentPostResumesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the document post resumes params
func (o *DocumentPostResumesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSovrenAccountID adds the sovrenAccountID to the document post resumes params
func (o *DocumentPostResumesParams) WithSovrenAccountID(sovrenAccountID string) *DocumentPostResumesParams {
	o.SetSovrenAccountID(sovrenAccountID)
	return o
}

// SetSovrenAccountID adds the sovrenAccountId to the document post resumes params
func (o *DocumentPostResumesParams) SetSovrenAccountID(sovrenAccountID string) {
	o.SovrenAccountID = sovrenAccountID
}

// WithSovrenServiceKey adds the sovrenServiceKey to the document post resumes params
func (o *DocumentPostResumesParams) WithSovrenServiceKey(sovrenServiceKey string) *DocumentPostResumesParams {
	o.SetSovrenServiceKey(sovrenServiceKey)
	return o
}

// SetSovrenServiceKey adds the sovrenServiceKey to the document post resumes params
func (o *DocumentPostResumesParams) SetSovrenServiceKey(sovrenServiceKey string) {
	o.SovrenServiceKey = sovrenServiceKey
}

// WithIndexID adds the indexID to the document post resumes params
func (o *DocumentPostResumesParams) WithIndexID(indexID string) *DocumentPostResumesParams {
	o.SetIndexID(indexID)
	return o
}

// SetIndexID adds the indexId to the document post resumes params
func (o *DocumentPostResumesParams) SetIndexID(indexID string) {
	o.IndexID = indexID
}

// WithRequest adds the request to the document post resumes params
func (o *DocumentPostResumesParams) WithRequest(request *go_sovren_models.SovrenSaasDomainV10AddMultipleResumesToIndexRequestModel) *DocumentPostResumesParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the document post resumes params
func (o *DocumentPostResumesParams) SetRequest(request *go_sovren_models.SovrenSaasDomainV10AddMultipleResumesToIndexRequestModel) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *DocumentPostResumesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param indexId
	if err := r.SetPathParam("indexId", o.IndexID); err != nil {
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