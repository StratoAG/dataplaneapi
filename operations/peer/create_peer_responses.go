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

package peer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models/v2"
)

// CreatePeerCreatedCode is the HTTP code returned for type CreatePeerCreated
const CreatePeerCreatedCode int = 201

/*CreatePeerCreated Peer created

swagger:response createPeerCreated
*/
type CreatePeerCreated struct {

	/*
	  In: Body
	*/
	Payload *models.PeerSection `json:"body,omitempty"`
}

// NewCreatePeerCreated creates CreatePeerCreated with default headers values
func NewCreatePeerCreated() *CreatePeerCreated {

	return &CreatePeerCreated{}
}

// WithPayload adds the payload to the create peer created response
func (o *CreatePeerCreated) WithPayload(payload *models.PeerSection) *CreatePeerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create peer created response
func (o *CreatePeerCreated) SetPayload(payload *models.PeerSection) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePeerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePeerAcceptedCode is the HTTP code returned for type CreatePeerAccepted
const CreatePeerAcceptedCode int = 202

/*CreatePeerAccepted Configuration change accepted and reload requested

swagger:response createPeerAccepted
*/
type CreatePeerAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.PeerSection `json:"body,omitempty"`
}

// NewCreatePeerAccepted creates CreatePeerAccepted with default headers values
func NewCreatePeerAccepted() *CreatePeerAccepted {

	return &CreatePeerAccepted{}
}

// WithReloadID adds the reloadId to the create peer accepted response
func (o *CreatePeerAccepted) WithReloadID(reloadID string) *CreatePeerAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create peer accepted response
func (o *CreatePeerAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create peer accepted response
func (o *CreatePeerAccepted) WithPayload(payload *models.PeerSection) *CreatePeerAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create peer accepted response
func (o *CreatePeerAccepted) SetPayload(payload *models.PeerSection) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePeerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePeerBadRequestCode is the HTTP code returned for type CreatePeerBadRequest
const CreatePeerBadRequestCode int = 400

/*CreatePeerBadRequest Bad request

swagger:response createPeerBadRequest
*/
type CreatePeerBadRequest struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreatePeerBadRequest creates CreatePeerBadRequest with default headers values
func NewCreatePeerBadRequest() *CreatePeerBadRequest {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreatePeerBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the create peer bad request response
func (o *CreatePeerBadRequest) WithConfigurationVersion(configurationVersion int64) *CreatePeerBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create peer bad request response
func (o *CreatePeerBadRequest) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create peer bad request response
func (o *CreatePeerBadRequest) WithPayload(payload *models.Error) *CreatePeerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create peer bad request response
func (o *CreatePeerBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePeerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePeerConflictCode is the HTTP code returned for type CreatePeerConflict
const CreatePeerConflictCode int = 409

/*CreatePeerConflict The specified resource already exists

swagger:response createPeerConflict
*/
type CreatePeerConflict struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreatePeerConflict creates CreatePeerConflict with default headers values
func NewCreatePeerConflict() *CreatePeerConflict {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreatePeerConflict{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the create peer conflict response
func (o *CreatePeerConflict) WithConfigurationVersion(configurationVersion int64) *CreatePeerConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create peer conflict response
func (o *CreatePeerConflict) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create peer conflict response
func (o *CreatePeerConflict) WithPayload(payload *models.Error) *CreatePeerConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create peer conflict response
func (o *CreatePeerConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePeerConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreatePeerDefault General Error

swagger:response createPeerDefault
*/
type CreatePeerDefault struct {
	_statusCode int
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreatePeerDefault creates CreatePeerDefault with default headers values
func NewCreatePeerDefault(code int) *CreatePeerDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreatePeerDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the create peer default response
func (o *CreatePeerDefault) WithStatusCode(code int) *CreatePeerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create peer default response
func (o *CreatePeerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create peer default response
func (o *CreatePeerDefault) WithConfigurationVersion(configurationVersion int64) *CreatePeerDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create peer default response
func (o *CreatePeerDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create peer default response
func (o *CreatePeerDefault) WithPayload(payload *models.Error) *CreatePeerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create peer default response
func (o *CreatePeerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePeerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}