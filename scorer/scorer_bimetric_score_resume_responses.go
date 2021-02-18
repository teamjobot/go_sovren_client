// Code generated by go-swagger; DO NOT EDIT.

package scorer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/teamjobot/go_sovren_models"
)

// ScorerBimetricScoreResumeReader is a Reader for the ScorerBimetricScoreResume structure.
type ScorerBimetricScoreResumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScorerBimetricScoreResumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScorerBimetricScoreResumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewScorerBimetricScoreResumeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewScorerBimetricScoreResumeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewScorerBimetricScoreResumeOK creates a ScorerBimetricScoreResumeOK with default headers values
func NewScorerBimetricScoreResumeOK() *ScorerBimetricScoreResumeOK {
	return &ScorerBimetricScoreResumeOK{}
}

/* ScorerBimetricScoreResumeOK describes a response with status code 200, with default header values.

OK
*/
type ScorerBimetricScoreResumeOK struct {
	Payload *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenV10BimetricMatchResponseSovrenCommonVersion9520CultureNeutralPublicKeyTokenNull
}

func (o *ScorerBimetricScoreResumeOK) Error() string {
	return fmt.Sprintf("[POST /v10/scorer/bimetric/resume][%d] scorerBimetricScoreResumeOK  %+v", 200, o.Payload)
}
func (o *ScorerBimetricScoreResumeOK) GetPayload() *go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenV10BimetricMatchResponseSovrenCommonVersion9520CultureNeutralPublicKeyTokenNull {
	return o.Payload
}

func (o *ScorerBimetricScoreResumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(go_sovren_models.SovrenSaasDomainV10ResponseModel1SovrenV10BimetricMatchResponseSovrenCommonVersion9520CultureNeutralPublicKeyTokenNull)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScorerBimetricScoreResumeBadRequest creates a ScorerBimetricScoreResumeBadRequest with default headers values
func NewScorerBimetricScoreResumeBadRequest() *ScorerBimetricScoreResumeBadRequest {
	return &ScorerBimetricScoreResumeBadRequest{}
}

/* ScorerBimetricScoreResumeBadRequest describes a response with status code 400, with default header values.

A message describing why the request couldn't be handled.
*/
type ScorerBimetricScoreResumeBadRequest struct {
	Payload string
}

func (o *ScorerBimetricScoreResumeBadRequest) Error() string {
	return fmt.Sprintf("[POST /v10/scorer/bimetric/resume][%d] scorerBimetricScoreResumeBadRequest  %+v", 400, o.Payload)
}
func (o *ScorerBimetricScoreResumeBadRequest) GetPayload() string {
	return o.Payload
}

func (o *ScorerBimetricScoreResumeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScorerBimetricScoreResumeUnauthorized creates a ScorerBimetricScoreResumeUnauthorized with default headers values
func NewScorerBimetricScoreResumeUnauthorized() *ScorerBimetricScoreResumeUnauthorized {
	return &ScorerBimetricScoreResumeUnauthorized{}
}

/* ScorerBimetricScoreResumeUnauthorized describes a response with status code 401, with default header values.

A message describing why the message was unauthorized.
*/
type ScorerBimetricScoreResumeUnauthorized struct {
	Payload string
}

func (o *ScorerBimetricScoreResumeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v10/scorer/bimetric/resume][%d] scorerBimetricScoreResumeUnauthorized  %+v", 401, o.Payload)
}
func (o *ScorerBimetricScoreResumeUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *ScorerBimetricScoreResumeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
