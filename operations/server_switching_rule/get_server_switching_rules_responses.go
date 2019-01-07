// Code generated by go-swagger; DO NOT EDIT.

package server_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetServerSwitchingRulesOKCode is the HTTP code returned for type GetServerSwitchingRulesOK
const GetServerSwitchingRulesOKCode int = 200

/*GetServerSwitchingRulesOK Successful operation

swagger:response getServerSwitchingRulesOK
*/
type GetServerSwitchingRulesOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetServerSwitchingRulesOKBody `json:"body,omitempty"`
}

// NewGetServerSwitchingRulesOK creates GetServerSwitchingRulesOK with default headers values
func NewGetServerSwitchingRulesOK() *GetServerSwitchingRulesOK {

	return &GetServerSwitchingRulesOK{}
}

// WithPayload adds the payload to the get server switching rules o k response
func (o *GetServerSwitchingRulesOK) WithPayload(payload *models.GetServerSwitchingRulesOKBody) *GetServerSwitchingRulesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server switching rules o k response
func (o *GetServerSwitchingRulesOK) SetPayload(payload *models.GetServerSwitchingRulesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerSwitchingRulesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetServerSwitchingRulesDefault General Error

swagger:response getServerSwitchingRulesDefault
*/
type GetServerSwitchingRulesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServerSwitchingRulesDefault creates GetServerSwitchingRulesDefault with default headers values
func NewGetServerSwitchingRulesDefault(code int) *GetServerSwitchingRulesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetServerSwitchingRulesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get server switching rules default response
func (o *GetServerSwitchingRulesDefault) WithStatusCode(code int) *GetServerSwitchingRulesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get server switching rules default response
func (o *GetServerSwitchingRulesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get server switching rules default response
func (o *GetServerSwitchingRulesDefault) WithPayload(payload *models.Error) *GetServerSwitchingRulesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server switching rules default response
func (o *GetServerSwitchingRulesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerSwitchingRulesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}