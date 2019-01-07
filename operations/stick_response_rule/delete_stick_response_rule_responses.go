// Code generated by go-swagger; DO NOT EDIT.

package stick_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// DeleteStickResponseRuleNoContentCode is the HTTP code returned for type DeleteStickResponseRuleNoContent
const DeleteStickResponseRuleNoContentCode int = 204

/*DeleteStickResponseRuleNoContent Stick Response Rule deleted

swagger:response deleteStickResponseRuleNoContent
*/
type DeleteStickResponseRuleNoContent struct {
}

// NewDeleteStickResponseRuleNoContent creates DeleteStickResponseRuleNoContent with default headers values
func NewDeleteStickResponseRuleNoContent() *DeleteStickResponseRuleNoContent {

	return &DeleteStickResponseRuleNoContent{}
}

// WriteResponse to the client
func (o *DeleteStickResponseRuleNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteStickResponseRuleNotFoundCode is the HTTP code returned for type DeleteStickResponseRuleNotFound
const DeleteStickResponseRuleNotFoundCode int = 404

/*DeleteStickResponseRuleNotFound The specified resource was not found

swagger:response deleteStickResponseRuleNotFound
*/
type DeleteStickResponseRuleNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteStickResponseRuleNotFound creates DeleteStickResponseRuleNotFound with default headers values
func NewDeleteStickResponseRuleNotFound() *DeleteStickResponseRuleNotFound {

	return &DeleteStickResponseRuleNotFound{}
}

// WithPayload adds the payload to the delete stick response rule not found response
func (o *DeleteStickResponseRuleNotFound) WithPayload(payload *models.Error) *DeleteStickResponseRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete stick response rule not found response
func (o *DeleteStickResponseRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStickResponseRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteStickResponseRuleDefault General Error

swagger:response deleteStickResponseRuleDefault
*/
type DeleteStickResponseRuleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteStickResponseRuleDefault creates DeleteStickResponseRuleDefault with default headers values
func NewDeleteStickResponseRuleDefault(code int) *DeleteStickResponseRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteStickResponseRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete stick response rule default response
func (o *DeleteStickResponseRuleDefault) WithStatusCode(code int) *DeleteStickResponseRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete stick response rule default response
func (o *DeleteStickResponseRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete stick response rule default response
func (o *DeleteStickResponseRuleDefault) WithPayload(payload *models.Error) *DeleteStickResponseRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete stick response rule default response
func (o *DeleteStickResponseRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStickResponseRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}