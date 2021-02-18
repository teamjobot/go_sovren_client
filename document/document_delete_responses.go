// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/teamjobot/go_sovren_models"
)

// DocumentDeleteReader is a Reader for the DocumentDelete structure.
type DocumentDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocumentDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDocumentDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDocumentDeleteOK creates a DocumentDeleteOK with default headers values
func NewDocumentDeleteOK() *DocumentDeleteOK {
	return &DocumentDeleteOK{}
}

/* DocumentDeleteOK describes a response with status code 200, with default header values.

OK
*/
type DocumentDeleteOK struct {
	Payload *go_sovren_models.SovrenSaasDomainV10ResponseModel1SystemObjectMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089
}

func (o *DocumentDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v10/index/{indexId}/documents/{documentId}][%d] documentDeleteOK  %+v", 200, o.Payload)
}
func (o *DocumentDeleteOK) GetPayload() *go_sovren_models.SovrenSaasDomainV10ResponseModel1SystemObjectMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 {
	return o.Payload
}

func (o *DocumentDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(go_sovren_models.SovrenSaasDomainV10ResponseModel1SystemObjectMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}