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

// DocumentGetResumeReader is a Reader for the DocumentGetResume structure.
type DocumentGetResumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocumentGetResumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDocumentGetResumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDocumentGetResumeOK creates a DocumentGetResumeOK with default headers values
func NewDocumentGetResumeOK() *DocumentGetResumeOK {
	return &DocumentGetResumeOK{}
}

/* DocumentGetResumeOK describes a response with status code 200, with default header values.

OK
*/
type DocumentGetResumeOK struct {
	Payload *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenParserModelResumeResumeDataSovrenCommonVersion9520CultureNeutralPublicKeyTokenNull
}

func (o *DocumentGetResumeOK) Error() string {
	return fmt.Sprintf("[GET /v10/index/{indexId}/resume/{documentId}][%d] documentGetResumeOK  %+v", 200, o.Payload)
}
func (o *DocumentGetResumeOK) GetPayload() *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenParserModelResumeResumeDataSovrenCommonVersion9520CultureNeutralPublicKeyTokenNull {
	return o.Payload
}

func (o *DocumentGetResumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenParserModelResumeResumeDataSovrenCommonVersion9520CultureNeutralPublicKeyTokenNull)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
