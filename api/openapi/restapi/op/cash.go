// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CashHandlerFunc turns a function with the right signature into a cash handler
type CashHandlerFunc func(CashParams) CashResponder

// Handle executing the request and returning a response
func (fn CashHandlerFunc) Handle(params CashParams) CashResponder {
	return fn(params)
}

// CashHandler interface for that can handle valid cash params
type CashHandler interface {
	Handle(CashParams) CashResponder
}

// NewCash creates a new http.Handler for the cash operation
func NewCash(ctx *middleware.Context, handler CashHandler) *Cash {
	return &Cash{Context: ctx, Handler: handler}
}

/*Cash swagger:route POST /cash cash

Exchange provided treasure for money.

*/
type Cash struct {
	Context *middleware.Context
	Handler CashHandler
}

func (o *Cash) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	Params := NewCashParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)
}
