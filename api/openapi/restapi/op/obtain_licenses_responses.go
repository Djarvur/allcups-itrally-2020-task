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

// ObtainLicensesOKCode is the HTTP code returned for type ObtainLicensesOK
const ObtainLicensesOKCode int = 200

/*ObtainLicensesOK list of licenses issued

swagger:response obtainLicensesOK
*/
type ObtainLicensesOK struct {

	/*
	  In: Body
	*/
	Payload model.Licenses `json:"body,omitempty"`
}

// NewObtainLicensesOK creates ObtainLicensesOK with default headers values
func NewObtainLicensesOK() *ObtainLicensesOK {

	return &ObtainLicensesOK{}
}

// WithPayload adds the payload to the obtain licenses o k response
func (o *ObtainLicensesOK) WithPayload(payload model.Licenses) *ObtainLicensesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the obtain licenses o k response
func (o *ObtainLicensesOK) SetPayload(payload model.Licenses) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObtainLicensesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = model.Licenses{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *ObtainLicensesOK) ObtainLicensesResponder() {}

// ObtainLicensesPaymentRequiredCode is the HTTP code returned for type ObtainLicensesPaymentRequired
const ObtainLicensesPaymentRequiredCode int = 402

/*ObtainLicensesPaymentRequired Not enough money

swagger:response obtainLicensesPaymentRequired
*/
type ObtainLicensesPaymentRequired struct {
}

// NewObtainLicensesPaymentRequired creates ObtainLicensesPaymentRequired with default headers values
func NewObtainLicensesPaymentRequired() *ObtainLicensesPaymentRequired {

	return &ObtainLicensesPaymentRequired{}
}

// WriteResponse to the client
func (o *ObtainLicensesPaymentRequired) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(402)
}

func (o *ObtainLicensesPaymentRequired) ObtainLicensesResponder() {}

// ObtainLicensesInternalServerErrorCode is the HTTP code returned for type ObtainLicensesInternalServerError
const ObtainLicensesInternalServerErrorCode int = 500

/*ObtainLicensesInternalServerError Internal Server Error

swagger:response obtainLicensesInternalServerError
*/
type ObtainLicensesInternalServerError struct {
}

// NewObtainLicensesInternalServerError creates ObtainLicensesInternalServerError with default headers values
func NewObtainLicensesInternalServerError() *ObtainLicensesInternalServerError {

	return &ObtainLicensesInternalServerError{}
}

// WriteResponse to the client
func (o *ObtainLicensesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

func (o *ObtainLicensesInternalServerError) ObtainLicensesResponder() {}

type ObtainLicensesNotImplementedResponder struct {
	middleware.Responder
}

func (*ObtainLicensesNotImplementedResponder) ObtainLicensesResponder() {}

func ObtainLicensesNotImplemented() ObtainLicensesResponder {
	return &ObtainLicensesNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.ObtainLicenses has not yet been implemented",
		),
	}
}

type ObtainLicensesResponder interface {
	middleware.Responder
	ObtainLicensesResponder()
}