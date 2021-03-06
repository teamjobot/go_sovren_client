// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/teamjobot/go_sovren_models"
)

// AccountGetReader is a Reader for the AccountGet structure.
type AccountGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAccountGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAccountGetOK creates a AccountGetOK with default headers values
func NewAccountGetOK() *AccountGetOK {
	return &AccountGetOK{}
}

/* AccountGetOK describes a response with status code 200, with default header values.

OK
*/
type AccountGetOK struct {
	Payload *go_sovren_models.SovrenSaasDomainV10ResponseModel1SystemObjectMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089
}

func (o *AccountGetOK) Error() string {
	return fmt.Sprintf("[GET /v10/account][%d] accountGetOK  %+v", 200, o.Payload)
}
func (o *AccountGetOK) GetPayload() *go_sovren_models.SovrenSaasDomainV10ResponseModel1SystemObjectMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 {
	return o.Payload
}

func (o *AccountGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(go_sovren_models.SovrenSaasDomainV10ResponseModel1SystemObjectMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountGetUnauthorized creates a AccountGetUnauthorized with default headers values
func NewAccountGetUnauthorized() *AccountGetUnauthorized {
	return &AccountGetUnauthorized{}
}

/* AccountGetUnauthorized describes a response with status code 401, with default header values.

A message describing why the message was unauthorized.
*/
type AccountGetUnauthorized struct {
	Payload string
}

func (o *AccountGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v10/account][%d] accountGetUnauthorized  %+v", 401, o.Payload)
}
func (o *AccountGetUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *AccountGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
