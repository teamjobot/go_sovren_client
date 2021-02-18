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
)

// NewDocumentGetResumeParams creates a new DocumentGetResumeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDocumentGetResumeParams() *DocumentGetResumeParams {
	return &DocumentGetResumeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDocumentGetResumeParamsWithTimeout creates a new DocumentGetResumeParams object
// with the ability to set a timeout on a request.
func NewDocumentGetResumeParamsWithTimeout(timeout time.Duration) *DocumentGetResumeParams {
	return &DocumentGetResumeParams{
		timeout: timeout,
	}
}

// NewDocumentGetResumeParamsWithContext creates a new DocumentGetResumeParams object
// with the ability to set a context for a request.
func NewDocumentGetResumeParamsWithContext(ctx context.Context) *DocumentGetResumeParams {
	return &DocumentGetResumeParams{
		Context: ctx,
	}
}

// NewDocumentGetResumeParamsWithHTTPClient creates a new DocumentGetResumeParams object
// with the ability to set a custom HTTPClient for a request.
func NewDocumentGetResumeParamsWithHTTPClient(client *http.Client) *DocumentGetResumeParams {
	return &DocumentGetResumeParams{
		HTTPClient: client,
	}
}

/* DocumentGetResumeParams contains all the parameters to send to the API endpoint
   for the document get resume operation.

   Typically these are written to a http.Request.
*/
type DocumentGetResumeParams struct {

	/* SovrenAccountID.

	   Provided Account Id
	*/
	SovrenAccountID string

	/* SovrenServiceKey.

	   Provided Service Key
	*/
	SovrenServiceKey string

	// DocumentID.
	DocumentID string

	// IndexID.
	IndexID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the document get resume params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DocumentGetResumeParams) WithDefaults() *DocumentGetResumeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the document get resume params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DocumentGetResumeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the document get resume params
func (o *DocumentGetResumeParams) WithTimeout(timeout time.Duration) *DocumentGetResumeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the document get resume params
func (o *DocumentGetResumeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the document get resume params
func (o *DocumentGetResumeParams) WithContext(ctx context.Context) *DocumentGetResumeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the document get resume params
func (o *DocumentGetResumeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the document get resume params
func (o *DocumentGetResumeParams) WithHTTPClient(client *http.Client) *DocumentGetResumeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the document get resume params
func (o *DocumentGetResumeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSovrenAccountID adds the sovrenAccountID to the document get resume params
func (o *DocumentGetResumeParams) WithSovrenAccountID(sovrenAccountID string) *DocumentGetResumeParams {
	o.SetSovrenAccountID(sovrenAccountID)
	return o
}

// SetSovrenAccountID adds the sovrenAccountId to the document get resume params
func (o *DocumentGetResumeParams) SetSovrenAccountID(sovrenAccountID string) {
	o.SovrenAccountID = sovrenAccountID
}

// WithSovrenServiceKey adds the sovrenServiceKey to the document get resume params
func (o *DocumentGetResumeParams) WithSovrenServiceKey(sovrenServiceKey string) *DocumentGetResumeParams {
	o.SetSovrenServiceKey(sovrenServiceKey)
	return o
}

// SetSovrenServiceKey adds the sovrenServiceKey to the document get resume params
func (o *DocumentGetResumeParams) SetSovrenServiceKey(sovrenServiceKey string) {
	o.SovrenServiceKey = sovrenServiceKey
}

// WithDocumentID adds the documentID to the document get resume params
func (o *DocumentGetResumeParams) WithDocumentID(documentID string) *DocumentGetResumeParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the document get resume params
func (o *DocumentGetResumeParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WithIndexID adds the indexID to the document get resume params
func (o *DocumentGetResumeParams) WithIndexID(indexID string) *DocumentGetResumeParams {
	o.SetIndexID(indexID)
	return o
}

// SetIndexID adds the indexId to the document get resume params
func (o *DocumentGetResumeParams) SetIndexID(indexID string) {
	o.IndexID = indexID
}

// WriteToRequest writes these params to a swagger request
func (o *DocumentGetResumeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param documentId
	if err := r.SetPathParam("documentId", o.DocumentID); err != nil {
		return err
	}

	// path param indexId
	if err := r.SetPathParam("indexId", o.IndexID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}