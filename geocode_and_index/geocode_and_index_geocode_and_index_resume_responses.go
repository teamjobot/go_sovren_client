// Code generated by go-swagger; DO NOT EDIT.

package geocode_and_index

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/teamjobot/go_sovren_models"
)

// GeocodeAndIndexGeocodeAndIndexResumeReader is a Reader for the GeocodeAndIndexGeocodeAndIndexResume structure.
type GeocodeAndIndexGeocodeAndIndexResumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GeocodeAndIndexGeocodeAndIndexResumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGeocodeAndIndexGeocodeAndIndexResumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGeocodeAndIndexGeocodeAndIndexResumeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGeocodeAndIndexGeocodeAndIndexResumeOK creates a GeocodeAndIndexGeocodeAndIndexResumeOK with default headers values
func NewGeocodeAndIndexGeocodeAndIndexResumeOK() *GeocodeAndIndexGeocodeAndIndexResumeOK {
	return &GeocodeAndIndexGeocodeAndIndexResumeOK{}
}

/* GeocodeAndIndexGeocodeAndIndexResumeOK describes a response with status code 200, with default header values.

OK
*/
type GeocodeAndIndexGeocodeAndIndexResumeOK struct {
	Payload *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10GeocodeAndIndexResumeResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull
}

func (o *GeocodeAndIndexGeocodeAndIndexResumeOK) Error() string {
	return fmt.Sprintf("[POST /v10/geocodeAndIndex/resume][%d] geocodeAndIndexGeocodeAndIndexResumeOK  %+v", 200, o.Payload)
}
func (o *GeocodeAndIndexGeocodeAndIndexResumeOK) GetPayload() *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10GeocodeAndIndexResumeResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull {
	return o.Payload
}

func (o *GeocodeAndIndexGeocodeAndIndexResumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10GeocodeAndIndexResumeResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeocodeAndIndexGeocodeAndIndexResumeUnauthorized creates a GeocodeAndIndexGeocodeAndIndexResumeUnauthorized with default headers values
func NewGeocodeAndIndexGeocodeAndIndexResumeUnauthorized() *GeocodeAndIndexGeocodeAndIndexResumeUnauthorized {
	return &GeocodeAndIndexGeocodeAndIndexResumeUnauthorized{}
}

/* GeocodeAndIndexGeocodeAndIndexResumeUnauthorized describes a response with status code 401, with default header values.

A message describing why the message was unauthorized.
*/
type GeocodeAndIndexGeocodeAndIndexResumeUnauthorized struct {
	Payload string
}

func (o *GeocodeAndIndexGeocodeAndIndexResumeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v10/geocodeAndIndex/resume][%d] geocodeAndIndexGeocodeAndIndexResumeUnauthorized  %+v", 401, o.Payload)
}
func (o *GeocodeAndIndexGeocodeAndIndexResumeUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *GeocodeAndIndexGeocodeAndIndexResumeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
