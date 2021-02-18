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

// IndexGet0Reader is a Reader for the IndexGet0 structure.
type IndexGet0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexGet0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexGet0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexGet0OK creates a IndexGet0OK with default headers values
func NewIndexGet0OK() *IndexGet0OK {
	return &IndexGet0OK{}
}

/* IndexGet0OK describes a response with status code 200, with default header values.

OK
*/
type IndexGet0OK struct {
	Payload *go_sovren_models.SovrenSaasDomainV10ResponseModel1SystemCollectionsGenericList1SovrenIndexDetailSovrenCommonVersion9520CultureNeutralPublicKeyTokenNullMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089
}

func (o *IndexGet0OK) Error() string {
	return fmt.Sprintf("[GET /v10/index][%d] indexGet0OK  %+v", 200, o.Payload)
}
func (o *IndexGet0OK) GetPayload() *go_sovren_models.SovrenSaasDomainV10ResponseModel1SystemCollectionsGenericList1SovrenIndexDetailSovrenCommonVersion9520CultureNeutralPublicKeyTokenNullMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089 {
	return o.Payload
}

func (o *IndexGet0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(go_sovren_models.SovrenSaasDomainV10ResponseModel1SystemCollectionsGenericList1SovrenIndexDetailSovrenCommonVersion9520CultureNeutralPublicKeyTokenNullMscorlibVersion4000CultureNeutralPublicKeyTokenB77a5c561934e089)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
