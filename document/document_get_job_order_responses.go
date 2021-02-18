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

// DocumentGetJobOrderReader is a Reader for the DocumentGetJobOrder structure.
type DocumentGetJobOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocumentGetJobOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDocumentGetJobOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDocumentGetJobOrderOK creates a DocumentGetJobOrderOK with default headers values
func NewDocumentGetJobOrderOK() *DocumentGetJobOrderOK {
	return &DocumentGetJobOrderOK{}
}

/* DocumentGetJobOrderOK describes a response with status code 200, with default header values.

OK
*/
type DocumentGetJobOrderOK struct {
	Payload *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenParserModelJobJobDataSovrenCommonVersion9520CultureNeutralPublicKeyTokenNull
}

func (o *DocumentGetJobOrderOK) Error() string {
	return fmt.Sprintf("[GET /v10/index/{indexId}/joborder/{documentId}][%d] documentGetJobOrderOK  %+v", 200, o.Payload)
}
func (o *DocumentGetJobOrderOK) GetPayload() *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenParserModelJobJobDataSovrenCommonVersion9520CultureNeutralPublicKeyTokenNull {
	return o.Payload
}

func (o *DocumentGetJobOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenParserModelJobJobDataSovrenCommonVersion9520CultureNeutralPublicKeyTokenNull)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}