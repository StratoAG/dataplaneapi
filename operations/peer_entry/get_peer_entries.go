// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package peer_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/models/v2"
)

// GetPeerEntriesHandlerFunc turns a function with the right signature into a get peer entries handler
type GetPeerEntriesHandlerFunc func(GetPeerEntriesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPeerEntriesHandlerFunc) Handle(params GetPeerEntriesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetPeerEntriesHandler interface for that can handle valid get peer entries params
type GetPeerEntriesHandler interface {
	Handle(GetPeerEntriesParams, interface{}) middleware.Responder
}

// NewGetPeerEntries creates a new http.Handler for the get peer entries operation
func NewGetPeerEntries(ctx *middleware.Context, handler GetPeerEntriesHandler) *GetPeerEntries {
	return &GetPeerEntries{Context: ctx, Handler: handler}
}

/*GetPeerEntries swagger:route GET /services/haproxy/configuration/peer_entries PeerEntry getPeerEntries

Return an array of peer_entries

Returns an array of all peer_entries that are configured in specified peer section.

*/
type GetPeerEntries struct {
	Context *middleware.Context
	Handler GetPeerEntriesHandler
}

func (o *GetPeerEntries) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPeerEntriesParams()

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

// GetPeerEntriesOKBody get peer entries o k body
//
// swagger:model GetPeerEntriesOKBody
type GetPeerEntriesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.PeerEntries `json:"data"`
}

// Validate validates this get peer entries o k body
func (o *GetPeerEntriesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPeerEntriesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getPeerEntriesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getPeerEntriesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPeerEntriesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPeerEntriesOKBody) UnmarshalBinary(b []byte) error {
	var res GetPeerEntriesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}