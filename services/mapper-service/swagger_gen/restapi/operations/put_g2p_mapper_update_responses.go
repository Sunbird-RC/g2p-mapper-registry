// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutG2pMapperUpdateUnauthorizedCode is the HTTP code returned for type PutG2pMapperUpdateUnauthorized
const PutG2pMapperUpdateUnauthorizedCode int = 401

/*
PutG2pMapperUpdateUnauthorized HTTP layer error details

swagger:response putG2pMapperUpdateUnauthorized
*/
type PutG2pMapperUpdateUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *PutG2pMapperUpdateUnauthorizedBody `json:"body,omitempty"`
}

// NewPutG2pMapperUpdateUnauthorized creates PutG2pMapperUpdateUnauthorized with default headers values
func NewPutG2pMapperUpdateUnauthorized() *PutG2pMapperUpdateUnauthorized {

	return &PutG2pMapperUpdateUnauthorized{}
}

// WithPayload adds the payload to the put g2p mapper update unauthorized response
func (o *PutG2pMapperUpdateUnauthorized) WithPayload(payload *PutG2pMapperUpdateUnauthorizedBody) *PutG2pMapperUpdateUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put g2p mapper update unauthorized response
func (o *PutG2pMapperUpdateUnauthorized) SetPayload(payload *PutG2pMapperUpdateUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutG2pMapperUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutG2pMapperUpdateForbiddenCode is the HTTP code returned for type PutG2pMapperUpdateForbidden
const PutG2pMapperUpdateForbiddenCode int = 403

/*
PutG2pMapperUpdateForbidden HTTP layer error details

swagger:response putG2pMapperUpdateForbidden
*/
type PutG2pMapperUpdateForbidden struct {

	/*
	  In: Body
	*/
	Payload *PutG2pMapperUpdateForbiddenBody `json:"body,omitempty"`
}

// NewPutG2pMapperUpdateForbidden creates PutG2pMapperUpdateForbidden with default headers values
func NewPutG2pMapperUpdateForbidden() *PutG2pMapperUpdateForbidden {

	return &PutG2pMapperUpdateForbidden{}
}

// WithPayload adds the payload to the put g2p mapper update forbidden response
func (o *PutG2pMapperUpdateForbidden) WithPayload(payload *PutG2pMapperUpdateForbiddenBody) *PutG2pMapperUpdateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put g2p mapper update forbidden response
func (o *PutG2pMapperUpdateForbidden) SetPayload(payload *PutG2pMapperUpdateForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutG2pMapperUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutG2pMapperUpdateInternalServerErrorCode is the HTTP code returned for type PutG2pMapperUpdateInternalServerError
const PutG2pMapperUpdateInternalServerErrorCode int = 500

/*
PutG2pMapperUpdateInternalServerError HTTP layer error details

swagger:response putG2pMapperUpdateInternalServerError
*/
type PutG2pMapperUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PutG2pMapperUpdateInternalServerErrorBody `json:"body,omitempty"`
}

// NewPutG2pMapperUpdateInternalServerError creates PutG2pMapperUpdateInternalServerError with default headers values
func NewPutG2pMapperUpdateInternalServerError() *PutG2pMapperUpdateInternalServerError {

	return &PutG2pMapperUpdateInternalServerError{}
}

// WithPayload adds the payload to the put g2p mapper update internal server error response
func (o *PutG2pMapperUpdateInternalServerError) WithPayload(payload *PutG2pMapperUpdateInternalServerErrorBody) *PutG2pMapperUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put g2p mapper update internal server error response
func (o *PutG2pMapperUpdateInternalServerError) SetPayload(payload *PutG2pMapperUpdateInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutG2pMapperUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PutG2pMapperUpdateDefault Acknowledgement of message received after successful validation of message and signature

swagger:response putG2pMapperUpdateDefault
*/
type PutG2pMapperUpdateDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *PutG2pMapperUpdateDefaultBody `json:"body,omitempty"`
}

// NewPutG2pMapperUpdateDefault creates PutG2pMapperUpdateDefault with default headers values
func NewPutG2pMapperUpdateDefault(code int) *PutG2pMapperUpdateDefault {
	if code <= 0 {
		code = 500
	}

	return &PutG2pMapperUpdateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put g2p mapper update default response
func (o *PutG2pMapperUpdateDefault) WithStatusCode(code int) *PutG2pMapperUpdateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put g2p mapper update default response
func (o *PutG2pMapperUpdateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put g2p mapper update default response
func (o *PutG2pMapperUpdateDefault) WithPayload(payload *PutG2pMapperUpdateDefaultBody) *PutG2pMapperUpdateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put g2p mapper update default response
func (o *PutG2pMapperUpdateDefault) SetPayload(payload *PutG2pMapperUpdateDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutG2pMapperUpdateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
