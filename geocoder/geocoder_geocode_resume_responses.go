// Code generated by go-swagger; DO NOT EDIT.

package geocoder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/teamjobot/go_sovren_models"
)

// GeocoderGeocodeResumeReader is a Reader for the GeocoderGeocodeResume structure.
type GeocoderGeocodeResumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GeocoderGeocodeResumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGeocoderGeocodeResumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGeocoderGeocodeResumeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGeocoderGeocodeResumeOK creates a GeocoderGeocodeResumeOK with default headers values
func NewGeocoderGeocodeResumeOK() *GeocoderGeocodeResumeOK {
	return &GeocoderGeocodeResumeOK{}
}

/* GeocoderGeocodeResumeOK describes a response with status code 200, with default header values.

OK
*/
type GeocoderGeocodeResumeOK struct {
	Payload *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10ResumeGeocoderResponseModelSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull
}

func (o *GeocoderGeocodeResumeOK) Error() string {
	return fmt.Sprintf("[POST /v10/geocoder/resume][%d] geocoderGeocodeResumeOK  %+v", 200, o.Payload)
}
func (o *GeocoderGeocodeResumeOK) GetPayload() *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10ResumeGeocoderResponseModelSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull {
	return o.Payload
}

func (o *GeocoderGeocodeResumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10ResumeGeocoderResponseModelSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeocoderGeocodeResumeUnauthorized creates a GeocoderGeocodeResumeUnauthorized with default headers values
func NewGeocoderGeocodeResumeUnauthorized() *GeocoderGeocodeResumeUnauthorized {
	return &GeocoderGeocodeResumeUnauthorized{}
}

/* GeocoderGeocodeResumeUnauthorized describes a response with status code 401, with default header values.

A message describing why the message was unauthorized.
*/
type GeocoderGeocodeResumeUnauthorized struct {
	Payload string
}

func (o *GeocoderGeocodeResumeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v10/geocoder/resume][%d] geocoderGeocodeResumeUnauthorized  %+v", 401, o.Payload)
}
func (o *GeocoderGeocodeResumeUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *GeocoderGeocodeResumeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}