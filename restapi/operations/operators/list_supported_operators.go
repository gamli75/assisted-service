// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListSupportedOperatorsHandlerFunc turns a function with the right signature into a list supported operators handler
type ListSupportedOperatorsHandlerFunc func(ListSupportedOperatorsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListSupportedOperatorsHandlerFunc) Handle(params ListSupportedOperatorsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListSupportedOperatorsHandler interface for that can handle valid list supported operators params
type ListSupportedOperatorsHandler interface {
	Handle(ListSupportedOperatorsParams, interface{}) middleware.Responder
}

// NewListSupportedOperators creates a new http.Handler for the list supported operators operation
func NewListSupportedOperators(ctx *middleware.Context, handler ListSupportedOperatorsHandler) *ListSupportedOperators {
	return &ListSupportedOperators{Context: ctx, Handler: handler}
}

/*ListSupportedOperators swagger:route GET /supported-operators operators listSupportedOperators

Retrieves the list of supported operators.

*/
type ListSupportedOperators struct {
	Context *middleware.Context
	Handler ListSupportedOperatorsHandler
}

func (o *ListSupportedOperators) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListSupportedOperatorsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
