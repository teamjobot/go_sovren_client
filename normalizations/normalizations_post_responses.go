// Code generated by go-swagger; DO NOT EDIT.

package normalizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/teamjobot/go_sovren_models"
)

// NormalizationsPostReader is a Reader for the NormalizationsPost structure.
type NormalizationsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NormalizationsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNormalizationsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNormalizationsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNormalizationsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNormalizationsPostOK creates a NormalizationsPostOK with default headers values
func NewNormalizationsPostOK() *NormalizationsPostOK {
	return &NormalizationsPostOK{}
}

/* NormalizationsPostOK describes a response with status code 200, with default header values.

A model of the Normalization created
*/
type NormalizationsPostOK struct {
	Payload *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainNormalizationResponseModelSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull
}

func (o *NormalizationsPostOK) Error() string {
	return fmt.Sprintf("[POST /v10/normalizations][%d] normalizationsPostOK  %+v", 200, o.Payload)
}
func (o *NormalizationsPostOK) GetPayload() *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainNormalizationResponseModelSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull {
	return o.Payload
}

func (o *NormalizationsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainNormalizationResponseModelSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNormalizationsPostBadRequest creates a NormalizationsPostBadRequest with default headers values
func NewNormalizationsPostBadRequest() *NormalizationsPostBadRequest {
	return &NormalizationsPostBadRequest{}
}

/* NormalizationsPostBadRequest describes a response with status code 400, with default header values.

A message describing why the message was in a bad state.
*/
type NormalizationsPostBadRequest struct {
	Payload string
}

func (o *NormalizationsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /v10/normalizations][%d] normalizationsPostBadRequest  %+v", 400, o.Payload)
}
func (o *NormalizationsPostBadRequest) GetPayload() string {
	return o.Payload
}

func (o *NormalizationsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNormalizationsPostUnauthorized creates a NormalizationsPostUnauthorized with default headers values
func NewNormalizationsPostUnauthorized() *NormalizationsPostUnauthorized {
	return &NormalizationsPostUnauthorized{}
}

/* NormalizationsPostUnauthorized describes a response with status code 401, with default header values.

A message describing why the message was unauthorized.
*/
type NormalizationsPostUnauthorized struct {
	Payload string
}

func (o *NormalizationsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v10/normalizations][%d] normalizationsPostUnauthorized  %+v", 401, o.Payload)
}
func (o *NormalizationsPostUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *NormalizationsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}