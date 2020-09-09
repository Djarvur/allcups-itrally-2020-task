// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
)

// CheckCubeReader is a Reader for the CheckCube structure.
type CheckCubeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckCubeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckCubeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 402:
		result := NewCheckCubePaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckCubeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckCubeOK creates a CheckCubeOK with default headers values
func NewCheckCubeOK() *CheckCubeOK {
	return &CheckCubeOK{}
}

/*CheckCubeOK handles this case with default header values.

number of objects (treasures or even rocks) found
*/
type CheckCubeOK struct {
	Payload model.Number
}

func (o *CheckCubeOK) Error() string {
	return fmt.Sprintf("[POST /check][%d] checkCubeOK  %+v", 200, o.Payload)
}

func (o *CheckCubeOK) GetPayload() model.Number {
	return o.Payload
}

func (o *CheckCubeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckCubePaymentRequired creates a CheckCubePaymentRequired with default headers values
func NewCheckCubePaymentRequired() *CheckCubePaymentRequired {
	return &CheckCubePaymentRequired{}
}

/*CheckCubePaymentRequired handles this case with default header values.

Not enough money
*/
type CheckCubePaymentRequired struct {
}

func (o *CheckCubePaymentRequired) Error() string {
	return fmt.Sprintf("[POST /check][%d] checkCubePaymentRequired ", 402)
}

func (o *CheckCubePaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCheckCubeInternalServerError creates a CheckCubeInternalServerError with default headers values
func NewCheckCubeInternalServerError() *CheckCubeInternalServerError {
	return &CheckCubeInternalServerError{}
}

/*CheckCubeInternalServerError handles this case with default header values.

Internal Server Error
*/
type CheckCubeInternalServerError struct {
}

func (o *CheckCubeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /check][%d] checkCubeInternalServerError ", 500)
}

func (o *CheckCubeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
