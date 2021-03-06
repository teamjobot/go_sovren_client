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

// NewMatcherPost0Params creates a new MatcherPost0Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMatcherPost0Params() *MatcherPost0Params {
	return &MatcherPost0Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewMatcherPost0ParamsWithTimeout creates a new MatcherPost0Params object
// with the ability to set a timeout on a request.
func NewMatcherPost0ParamsWithTimeout(timeout time.Duration) *MatcherPost0Params {
	return &MatcherPost0Params{
		timeout: timeout,
	}
}

// NewMatcherPost0ParamsWithContext creates a new MatcherPost0Params object
// with the ability to set a context for a request.
func NewMatcherPost0ParamsWithContext(ctx context.Context) *MatcherPost0Params {
	return &MatcherPost0Params{
		Context: ctx,
	}
}

// NewMatcherPost0ParamsWithHTTPClient creates a new MatcherPost0Params object
// with the ability to set a custom HTTPClient for a request.
func NewMatcherPost0ParamsWithHTTPClient(client *http.Client) *MatcherPost0Params {
	return &MatcherPost0Params{
		HTTPClient: client,
	}
}

/* MatcherPost0Params contains all the parameters to send to the API endpoint
   for the matcher post 0 operation.

   Typically these are written to a http.Request.
*/
type MatcherPost0Params struct {

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

	// Request.
	Request *go_sovren_models.SovrenV10MatchRequestBase

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the matcher post 0 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MatcherPost0Params) WithDefaults() *MatcherPost0Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the matcher post 0 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MatcherPost0Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the matcher post 0 params
func (o *MatcherPost0Params) WithTimeout(timeout time.Duration) *MatcherPost0Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the matcher post 0 params
func (o *MatcherPost0Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the matcher post 0 params
func (o *MatcherPost0Params) WithContext(ctx context.Context) *MatcherPost0Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the matcher post 0 params
func (o *MatcherPost0Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the matcher post 0 params
func (o *MatcherPost0Params) WithHTTPClient(client *http.Client) *MatcherPost0Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the matcher post 0 params
func (o *MatcherPost0Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSovrenAccountID adds the sovrenAccountID to the matcher post 0 params
func (o *MatcherPost0Params) WithSovrenAccountID(sovrenAccountID string) *MatcherPost0Params {
	o.SetSovrenAccountID(sovrenAccountID)
	return o
}

// SetSovrenAccountID adds the sovrenAccountId to the matcher post 0 params
func (o *MatcherPost0Params) SetSovrenAccountID(sovrenAccountID string) {
	o.SovrenAccountID = sovrenAccountID
}

// WithSovrenServiceKey adds the sovrenServiceKey to the matcher post 0 params
func (o *MatcherPost0Params) WithSovrenServiceKey(sovrenServiceKey string) *MatcherPost0Params {
	o.SetSovrenServiceKey(sovrenServiceKey)
	return o
}

// SetSovrenServiceKey adds the sovrenServiceKey to the matcher post 0 params
func (o *MatcherPost0Params) SetSovrenServiceKey(sovrenServiceKey string) {
	o.SovrenServiceKey = sovrenServiceKey
}

// WithDocumentID adds the documentID to the matcher post 0 params
func (o *MatcherPost0Params) WithDocumentID(documentID string) *MatcherPost0Params {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the matcher post 0 params
func (o *MatcherPost0Params) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WithIndexID adds the indexID to the matcher post 0 params
func (o *MatcherPost0Params) WithIndexID(indexID string) *MatcherPost0Params {
	o.SetIndexID(indexID)
	return o
}

// SetIndexID adds the indexId to the matcher post 0 params
func (o *MatcherPost0Params) SetIndexID(indexID string) {
	o.IndexID = indexID
}

// WithRequest adds the request to the matcher post 0 params
func (o *MatcherPost0Params) WithRequest(request *go_sovren_models.SovrenV10MatchRequestBase) *MatcherPost0Params {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the matcher post 0 params
func (o *MatcherPost0Params) SetRequest(request *go_sovren_models.SovrenV10MatchRequestBase) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *MatcherPost0Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
