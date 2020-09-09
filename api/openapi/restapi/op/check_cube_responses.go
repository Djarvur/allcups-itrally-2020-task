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

// CheckCubeOKCode is the HTTP code returned for type CheckCubeOK
const CheckCubeOKCode int = 200

/*CheckCubeOK number of objects (treasures or even rocks) found

swagger:response checkCubeOK
*/
type CheckCubeOK struct {

	/*
	  In: Body
	*/
	Payload model.Number `json:"body,omitempty"`
}

// NewCheckCubeOK creates CheckCubeOK with default headers values
func NewCheckCubeOK() *CheckCubeOK {

	return &CheckCubeOK{}
}

// WithPayload adds the payload to the check cube o k response
func (o *CheckCubeOK) WithPayload(payload model.Number) *CheckCubeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check cube o k response
func (o *CheckCubeOK) SetPayload(payload model.Number) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCubeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *CheckCubeOK) CheckCubeResponder() {}

// CheckCubePaymentRequiredCode is the HTTP code returned for type CheckCubePaymentRequired
const CheckCubePaymentRequiredCode int = 402

/*CheckCubePaymentRequired Not enough money

swagger:response checkCubePaymentRequired
*/
type CheckCubePaymentRequired struct {
}

// NewCheckCubePaymentRequired creates CheckCubePaymentRequired with default headers values
func NewCheckCubePaymentRequired() *CheckCubePaymentRequired {

	return &CheckCubePaymentRequired{}
}

// WriteResponse to the client
func (o *CheckCubePaymentRequired) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(402)
}

func (o *CheckCubePaymentRequired) CheckCubeResponder() {}

// CheckCubeInternalServerErrorCode is the HTTP code returned for type CheckCubeInternalServerError
const CheckCubeInternalServerErrorCode int = 500

/*CheckCubeInternalServerError Internal Server Error

swagger:response checkCubeInternalServerError
*/
type CheckCubeInternalServerError struct {
}

// NewCheckCubeInternalServerError creates CheckCubeInternalServerError with default headers values
func NewCheckCubeInternalServerError() *CheckCubeInternalServerError {

	return &CheckCubeInternalServerError{}
}

// WriteResponse to the client
func (o *CheckCubeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

func (o *CheckCubeInternalServerError) CheckCubeResponder() {}

type CheckCubeNotImplementedResponder struct {
	middleware.Responder
}

func (*CheckCubeNotImplementedResponder) CheckCubeResponder() {}

func CheckCubeNotImplemented() CheckCubeResponder {
	return &CheckCubeNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.CheckCube has not yet been implemented",
		),
	}
}

type CheckCubeResponder interface {
	middleware.Responder
	CheckCubeResponder()
}