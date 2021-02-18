// Code generated by go-swagger; DO NOT EDIT.

package index

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/teamjobot/go_sovren_models"
)

// IndexGetIndexDocumentCountReader is a Reader for the IndexGetIndexDocumentCount structure.
type IndexGetIndexDocumentCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexGetIndexDocumentCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexGetIndexDocumentCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexGetIndexDocumentCountOK creates a IndexGetIndexDocumentCountOK with default headers values
func NewIndexGetIndexDocumentCountOK() *IndexGetIndexDocumentCountOK {
	return &IndexGetIndexDocumentCountOK{}
}

/* IndexGetIndexDocumentCountOK describes a response with status code 200, with default header values.

OK
*/
type IndexGetIndexDocumentCountOK struct {
	Payload *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10V10IndexDocumentCountResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull
}

func (o *IndexGetIndexDocumentCountOK) Error() string {
	return fmt.Sprintf("[GET /v10/index/{indexId}/count][%d] indexGetIndexDocumentCountOK  %+v", 200, o.Payload)
}
func (o *IndexGetIndexDocumentCountOK) GetPayload() *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10V10IndexDocumentCountResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull {
	return o.Payload
}

func (o *IndexGetIndexDocumentCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10V10IndexDocumentCountResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
