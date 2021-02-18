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

// DocumentPatchJobOrderReader is a Reader for the DocumentPatchJobOrder structure.
type DocumentPatchJobOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocumentPatchJobOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDocumentPatchJobOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDocumentPatchJobOrderOK creates a DocumentPatchJobOrderOK with default headers values
func NewDocumentPatchJobOrderOK() *DocumentPatchJobOrderOK {
	return &DocumentPatchJobOrderOK{}
}

/* DocumentPatchJobOrderOK describes a response with status code 200, with default header values.

OK
*/
type DocumentPatchJobOrderOK struct {
	Payload *go_sovren_models.SovrenSaasDomainV10ResponseModel1SystemObjectMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089
}

func (o *DocumentPatchJobOrderOK) Error() string {
	return fmt.Sprintf("[PATCH /v10/index/{indexId}/joborder/{documentId}][%d] documentPatchJobOrderOK  %+v", 200, o.Payload)
}
func (o *DocumentPatchJobOrderOK) GetPayload() *go_sovren_models.SovrenSaasDomainV10ResponseModel1SystemObjectMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 {
	return o.Payload
}

func (o *DocumentPatchJobOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(go_sovren_models.SovrenSaasDomainV10ResponseModel1SystemObjectMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
