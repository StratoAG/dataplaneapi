// Code generated by go-swagger; DO NOT EDIT.

package stick_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteStickResponseRuleHandlerFunc turns a function with the right signature into a delete stick response rule handler
type DeleteStickResponseRuleHandlerFunc func(DeleteStickResponseRuleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteStickResponseRuleHandlerFunc) Handle(params DeleteStickResponseRuleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteStickResponseRuleHandler interface for that can handle valid delete stick response rule params
type DeleteStickResponseRuleHandler interface {
	Handle(DeleteStickResponseRuleParams, interface{}) middleware.Responder
}

// NewDeleteStickResponseRule creates a new http.Handler for the delete stick response rule operation
func NewDeleteStickResponseRule(ctx *middleware.Context, handler DeleteStickResponseRuleHandler) *DeleteStickResponseRule {
	return &DeleteStickResponseRule{Context: ctx, Handler: handler}
}

/*DeleteStickResponseRule swagger:route DELETE /services/haproxy/configuration/stick_response_rules/{id} StickResponseRule deleteStickResponseRule

Delete a Stick Response Rule

Deletes a Stick Response Rule configuration by it's ID from the specified backend.

*/
type DeleteStickResponseRule struct {
	Context *middleware.Context
	Handler DeleteStickResponseRuleHandler
}

func (o *DeleteStickResponseRule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteStickResponseRuleParams()

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