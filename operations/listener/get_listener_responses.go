// Code generated by go-swagger; DO NOT EDIT.

package listener

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetListenerOKCode is the HTTP code returned for type GetListenerOK
const GetListenerOKCode int = 200

/*GetListenerOK Successful operation

swagger:response getListenerOK
*/
type GetListenerOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetListenerOKBody `json:"body,omitempty"`
}

// NewGetListenerOK creates GetListenerOK with default headers values
func NewGetListenerOK() *GetListenerOK {

	return &GetListenerOK{}
}

// WithPayload adds the payload to the get listener o k response
func (o *GetListenerOK) WithPayload(payload *models.GetListenerOKBody) *GetListenerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get listener o k response
func (o *GetListenerOK) SetPayload(payload *models.GetListenerOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetListenerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetListenerNotFoundCode is the HTTP code returned for type GetListenerNotFound
const GetListenerNotFoundCode int = 404

/*GetListenerNotFound The specified resource already exists

swagger:response getListenerNotFound
*/
type GetListenerNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetListenerNotFound creates GetListenerNotFound with default headers values
func NewGetListenerNotFound() *GetListenerNotFound {

	return &GetListenerNotFound{}
}

// WithPayload adds the payload to the get listener not found response
func (o *GetListenerNotFound) WithPayload(payload *models.Error) *GetListenerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get listener not found response
func (o *GetListenerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetListenerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetListenerDefault General Error

swagger:response getListenerDefault
*/
type GetListenerDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetListenerDefault creates GetListenerDefault with default headers values
func NewGetListenerDefault(code int) *GetListenerDefault {
	if code <= 0 {
		code = 500
	}

	return &GetListenerDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get listener default response
func (o *GetListenerDefault) WithStatusCode(code int) *GetListenerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get listener default response
func (o *GetListenerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get listener default response
func (o *GetListenerDefault) WithPayload(payload *models.Error) *GetListenerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get listener default response
func (o *GetListenerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetListenerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}