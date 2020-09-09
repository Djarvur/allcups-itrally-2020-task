// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Djarvur/allcups-itrally-2020-task/api/openapi/model"
)

// NewCheckCubeParams creates a new CheckCubeParams object
// no default values defined in spec.
func NewCheckCubeParams() CheckCubeParams {

	return CheckCubeParams{}
}

// CheckCubeParams contains all the bound params for the check cube operation
// typically these are obtained from a http.Request
//
// swagger:parameters checkCube
type CheckCubeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Cube to be checked with the license should be used to check. License available amount will be reduced according to the cube volume.
	  Required: true
	  In: body
	*/
	Request *model.Check
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCheckCubeParams() beforehand.
func (o *CheckCubeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body model.Check
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("request", "body", ""))
			} else {
				res = append(res, errors.NewParseError("request", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Request = &body
			}
		}
	} else {
		res = append(res, errors.Required("request", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
