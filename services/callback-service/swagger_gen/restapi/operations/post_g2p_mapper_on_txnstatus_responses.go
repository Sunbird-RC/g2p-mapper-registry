// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostG2pMapperOnTxnstatusUnauthorizedCode is the HTTP code returned for type PostG2pMapperOnTxnstatusUnauthorized
const PostG2pMapperOnTxnstatusUnauthorizedCode int = 401

/*
PostG2pMapperOnTxnstatusUnauthorized HTTP layer error details

swagger:response postG2pMapperOnTxnstatusUnauthorized
*/
type PostG2pMapperOnTxnstatusUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *PostG2pMapperOnTxnstatusUnauthorizedBody `json:"body,omitempty"`
}

// NewPostG2pMapperOnTxnstatusUnauthorized creates PostG2pMapperOnTxnstatusUnauthorized with default headers values
func NewPostG2pMapperOnTxnstatusUnauthorized() *PostG2pMapperOnTxnstatusUnauthorized {

	return &PostG2pMapperOnTxnstatusUnauthorized{}
}

// WithPayload adds the payload to the post g2p mapper on txnstatus unauthorized response
func (o *PostG2pMapperOnTxnstatusUnauthorized) WithPayload(payload *PostG2pMapperOnTxnstatusUnauthorizedBody) *PostG2pMapperOnTxnstatusUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p mapper on txnstatus unauthorized response
func (o *PostG2pMapperOnTxnstatusUnauthorized) SetPayload(payload *PostG2pMapperOnTxnstatusUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pMapperOnTxnstatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostG2pMapperOnTxnstatusForbiddenCode is the HTTP code returned for type PostG2pMapperOnTxnstatusForbidden
const PostG2pMapperOnTxnstatusForbiddenCode int = 403

/*
PostG2pMapperOnTxnstatusForbidden HTTP layer error details

swagger:response postG2pMapperOnTxnstatusForbidden
*/
type PostG2pMapperOnTxnstatusForbidden struct {

	/*
	  In: Body
	*/
	Payload *PostG2pMapperOnTxnstatusForbiddenBody `json:"body,omitempty"`
}

// NewPostG2pMapperOnTxnstatusForbidden creates PostG2pMapperOnTxnstatusForbidden with default headers values
func NewPostG2pMapperOnTxnstatusForbidden() *PostG2pMapperOnTxnstatusForbidden {

	return &PostG2pMapperOnTxnstatusForbidden{}
}

// WithPayload adds the payload to the post g2p mapper on txnstatus forbidden response
func (o *PostG2pMapperOnTxnstatusForbidden) WithPayload(payload *PostG2pMapperOnTxnstatusForbiddenBody) *PostG2pMapperOnTxnstatusForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p mapper on txnstatus forbidden response
func (o *PostG2pMapperOnTxnstatusForbidden) SetPayload(payload *PostG2pMapperOnTxnstatusForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pMapperOnTxnstatusForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostG2pMapperOnTxnstatusInternalServerErrorCode is the HTTP code returned for type PostG2pMapperOnTxnstatusInternalServerError
const PostG2pMapperOnTxnstatusInternalServerErrorCode int = 500

/*
PostG2pMapperOnTxnstatusInternalServerError HTTP layer error details

swagger:response postG2pMapperOnTxnstatusInternalServerError
*/
type PostG2pMapperOnTxnstatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PostG2pMapperOnTxnstatusInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostG2pMapperOnTxnstatusInternalServerError creates PostG2pMapperOnTxnstatusInternalServerError with default headers values
func NewPostG2pMapperOnTxnstatusInternalServerError() *PostG2pMapperOnTxnstatusInternalServerError {

	return &PostG2pMapperOnTxnstatusInternalServerError{}
}

// WithPayload adds the payload to the post g2p mapper on txnstatus internal server error response
func (o *PostG2pMapperOnTxnstatusInternalServerError) WithPayload(payload *PostG2pMapperOnTxnstatusInternalServerErrorBody) *PostG2pMapperOnTxnstatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p mapper on txnstatus internal server error response
func (o *PostG2pMapperOnTxnstatusInternalServerError) SetPayload(payload *PostG2pMapperOnTxnstatusInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pMapperOnTxnstatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PostG2pMapperOnTxnstatusDefault Acknowledgement of message received after successful validation of message and signature

swagger:response postG2pMapperOnTxnstatusDefault
*/
type PostG2pMapperOnTxnstatusDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *PostG2pMapperOnTxnstatusDefaultBody `json:"body,omitempty"`
}

// NewPostG2pMapperOnTxnstatusDefault creates PostG2pMapperOnTxnstatusDefault with default headers values
func NewPostG2pMapperOnTxnstatusDefault(code int) *PostG2pMapperOnTxnstatusDefault {
	if code <= 0 {
		code = 500
	}

	return &PostG2pMapperOnTxnstatusDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post g2p mapper on txnstatus default response
func (o *PostG2pMapperOnTxnstatusDefault) WithStatusCode(code int) *PostG2pMapperOnTxnstatusDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post g2p mapper on txnstatus default response
func (o *PostG2pMapperOnTxnstatusDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post g2p mapper on txnstatus default response
func (o *PostG2pMapperOnTxnstatusDefault) WithPayload(payload *PostG2pMapperOnTxnstatusDefaultBody) *PostG2pMapperOnTxnstatusDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p mapper on txnstatus default response
func (o *PostG2pMapperOnTxnstatusDefault) SetPayload(payload *PostG2pMapperOnTxnstatusDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pMapperOnTxnstatusDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
