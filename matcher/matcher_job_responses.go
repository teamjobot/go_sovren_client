// Code generated by go-swagger; DO NOT EDIT.

package matcher

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/teamjobot/go_sovren_models"
)

// MatcherJobReader is a Reader for the MatcherJob structure.
type MatcherJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MatcherJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMatcherJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMatcherJobOK creates a MatcherJobOK with default headers values
func NewMatcherJobOK() *MatcherJobOK {
	return &MatcherJobOK{}
}

/* MatcherJobOK describes a response with status code 200, with default header values.

OK
*/
type MatcherJobOK struct {
	Payload *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10V10MatchResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull
}

func (o *MatcherJobOK) Error() string {
	return fmt.Sprintf("[POST /v10/matcher/joborder][%d] matcherJobOK  %+v", 200, o.Payload)
}
func (o *MatcherJobOK) GetPayload() *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10V10MatchResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull {
	return o.Payload
}

func (o *MatcherJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenSaasDomainV10V10MatchResponseSovrenSaasDomainVersion9520CultureNeutralPublicKeyTokenNull)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
