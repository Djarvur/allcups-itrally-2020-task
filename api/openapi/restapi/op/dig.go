// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DigHandlerFunc turns a function with the right signature into a dig handler
type DigHandlerFunc func(DigParams) DigResponder

// Handle executing the request and returning a response
func (fn DigHandlerFunc) Handle(params DigParams) DigResponder {
	return fn(params)
}

// DigHandler interface for that can handle valid dig params
type DigHandler interface {
	Handle(DigParams) DigResponder
}

// NewDig creates a new http.Handler for the dig operation
func NewDig(ctx *middleware.Context, handler DigHandler) *Dig {
	return &Dig{Context: ctx, Handler: handler}
}

/*Dig swagger:route POST /dig dig

Dig at given point and depth, returns found treasures.

*/
type Dig struct {
	Context *middleware.Context
	Handler DigHandler
}

func (o *Dig) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDigParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}