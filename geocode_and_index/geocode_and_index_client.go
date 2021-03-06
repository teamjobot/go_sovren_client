// Code generated by go-swagger; DO NOT EDIT.

package geocode_and_index

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new geocode and index API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for geocode and index API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GeocodeAndIndexGeocodeAndIndexJob(params *GeocodeAndIndexGeocodeAndIndexJobParams, opts ...ClientOption) (*GeocodeAndIndexGeocodeAndIndexJobOK, error)

	GeocodeAndIndexGeocodeAndIndexResume(params *GeocodeAndIndexGeocodeAndIndexResumeParams, opts ...ClientOption) (*GeocodeAndIndexGeocodeAndIndexResumeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GeocodeAndIndexGeocodeAndIndexJob geocodes and index an already parsed v10 job order
*/
func (a *Client) GeocodeAndIndexGeocodeAndIndexJob(params *GeocodeAndIndexGeocodeAndIndexJobParams, opts ...ClientOption) (*GeocodeAndIndexGeocodeAndIndexJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGeocodeAndIndexGeocodeAndIndexJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GeocodeAndIndex_GeocodeAndIndexJob",
		Method:             "POST",
		PathPattern:        "/v10/geocodeAndIndex/joborder",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GeocodeAndIndexGeocodeAndIndexJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GeocodeAndIndexGeocodeAndIndexJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GeocodeAndIndex_GeocodeAndIndexJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GeocodeAndIndexGeocodeAndIndexResume geocodes and index an already parsed v10 resume
*/
func (a *Client) GeocodeAndIndexGeocodeAndIndexResume(params *GeocodeAndIndexGeocodeAndIndexResumeParams, opts ...ClientOption) (*GeocodeAndIndexGeocodeAndIndexResumeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGeocodeAndIndexGeocodeAndIndexResumeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GeocodeAndIndex_GeocodeAndIndexResume",
		Method:             "POST",
		PathPattern:        "/v10/geocodeAndIndex/resume",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GeocodeAndIndexGeocodeAndIndexResumeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GeocodeAndIndexGeocodeAndIndexResumeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GeocodeAndIndex_GeocodeAndIndexResume: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
