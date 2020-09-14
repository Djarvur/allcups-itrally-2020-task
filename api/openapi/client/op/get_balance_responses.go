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

// GetBalanceReader is a Reader for the GetBalance structure.
type GetBalanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBalanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBalanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBalanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBalanceOK creates a GetBalanceOK with default headers values
func NewGetBalanceOK() *GetBalanceOK {
	return &GetBalanceOK{}
}

/*GetBalanceOK handles this case with default header values.

Current balance.
*/
type GetBalanceOK struct {
	Payload *model.Balance
}

func (o *GetBalanceOK) Error() string {
	return fmt.Sprintf("[GET /balance][%d] getBalanceOK  %+v", 200, o.Payload)
}

func (o *GetBalanceOK) GetPayload() *model.Balance {
	return o.Payload
}

func (o *GetBalanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Balance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBalanceDefault creates a GetBalanceDefault with default headers values
func NewGetBalanceDefault(code int) *GetBalanceDefault {
	return &GetBalanceDefault{
		_statusCode: code,
	}
}

/*GetBalanceDefault handles this case with default header values.

General errors using same model as used by go-swagger for validation errors.
*/
type GetBalanceDefault struct {
	_statusCode int

	Payload *model.Error
}

// Code gets the status code for the get balance default response
func (o *GetBalanceDefault) Code() int {
	return o._statusCode
}

func (o *GetBalanceDefault) Error() string {
	return fmt.Sprintf("[GET /balance][%d] getBalance default  %+v", o._statusCode, o.Payload)
}

func (o *GetBalanceDefault) GetPayload() *model.Error {
	return o.Payload
}

func (o *GetBalanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
