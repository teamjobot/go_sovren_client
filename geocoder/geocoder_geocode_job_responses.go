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

// GeocoderGeocodeJobReader is a Reader for the GeocoderGeocodeJob structure.
type GeocoderGeocodeJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GeocoderGeocodeJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGeocoderGeocodeJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGeocoderGeocodeJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGeocoderGeocodeJobOK creates a GeocoderGeocodeJobOK with default headers values
func NewGeocoderGeocodeJobOK() *GeocoderGeocodeJobOK {
	return &GeocoderGeocodeJobOK{}
}

/* GeocoderGeocodeJobOK describes a response with status code 200, with default header values.

OK
*/
type GeocoderGeocodeJobOK struct {
	Payload *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10JobGeocoderResponseModelSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull
}

func (o *GeocoderGeocodeJobOK) Error() string {
	return fmt.Sprintf("[POST /v10/geocoder/joborder][%d] geocoderGeocodeJobOK  %+v", 200, o.Payload)
}
func (o *GeocoderGeocodeJobOK) GetPayload() *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10JobGeocoderResponseModelSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull {
	return o.Payload
}

func (o *GeocoderGeocodeJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10JobGeocoderResponseModelSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeocoderGeocodeJobUnauthorized creates a GeocoderGeocodeJobUnauthorized with default headers values
func NewGeocoderGeocodeJobUnauthorized() *GeocoderGeocodeJobUnauthorized {
	return &GeocoderGeocodeJobUnauthorized{}
}

/* GeocoderGeocodeJobUnauthorized describes a response with status code 401, with default header values.

A message describing why the message was unauthorized.
*/
type GeocoderGeocodeJobUnauthorized struct {
	Payload string
}

func (o *GeocoderGeocodeJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v10/geocoder/joborder][%d] geocoderGeocodeJobUnauthorized  %+v", 401, o.Payload)
}
func (o *GeocoderGeocodeJobUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *GeocoderGeocodeJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
