// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
)

// CashOKCode is the HTTP code returned for type CashOK
const CashOKCode int = 200

/*CashOK Payment for treasure.

swagger:response cashOK
*/
type CashOK struct {

	/*
	  In: Body
	*/
	Payload model.Wallet `json:"body,omitempty"`
}

// NewCashOK creates CashOK with default headers values
func NewCashOK() *CashOK {

	return &CashOK{}
}

// WithPayload adds the payload to the cash o k response
func (o *CashOK) WithPayload(payload model.Wallet) *CashOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the cash o k response
func (o *CashOK) SetPayload(payload model.Wallet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CashOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = model.Wallet{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *CashOK) CashResponder() {}

/*CashDefault - 409.1003: treasure is not digged


swagger:response cashDefault
*/
type CashDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewCashDefault creates CashDefault with default headers values
func NewCashDefault(code int) *CashDefault {
	if code <= 0 {
		code = 500
	}

	return &CashDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the cash default response
func (o *CashDefault) WithStatusCode(code int) *CashDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the cash default response
func (o *CashDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the cash default response
func (o *CashDefault) WithPayload(payload interface{}) *CashDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the cash default response
func (o *CashDefault) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CashDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *CashDefault) CashResponder() {}

type CashNotImplementedResponder struct {
	middleware.Responder
}

func (*CashNotImplementedResponder) CashResponder() {}

func CashNotImplemented() CashResponder {
	return &CashNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Cash has not yet been implemented",
		),
	}
}

type CashResponder interface {
	middleware.Responder
	CashResponder()
}
