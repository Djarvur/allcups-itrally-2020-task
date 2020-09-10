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

// DigReader is a Reader for the Dig structure.
type DigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDigOK creates a DigOK with default headers values
func NewDigOK() *DigOK {
	return &DigOK{}
}

/*DigOK handles this case with default header values.

List of treasures found.
*/
type DigOK struct {
	Payload model.TreasureList
}

func (o *DigOK) Error() string {
	return fmt.Sprintf("[POST /dig][%d] digOK  %+v", 200, o.Payload)
}

func (o *DigOK) GetPayload() model.TreasureList {
	return o.Payload
}

func (o *DigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDigDefault creates a DigDefault with default headers values
func NewDigDefault(code int) *DigDefault {
	return &DigDefault{
		_statusCode: code,
	}
}

/*DigDefault handles this case with default header values.

General errors using same model as used by go-swagger for validation errors.
*/
type DigDefault struct {
	_statusCode int

	Payload *model.Error
}

// Code gets the status code for the dig default response
func (o *DigDefault) Code() int {
	return o._statusCode
}

func (o *DigDefault) Error() string {
	return fmt.Sprintf("[POST /dig][%d] dig default  %+v", o._statusCode, o.Payload)
}

func (o *DigDefault) GetPayload() *model.Error {
	return o.Payload
}

func (o *DigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
