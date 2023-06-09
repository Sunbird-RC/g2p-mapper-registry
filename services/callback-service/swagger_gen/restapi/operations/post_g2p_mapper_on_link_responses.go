// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostG2pMapperOnLinkUnauthorizedCode is the HTTP code returned for type PostG2pMapperOnLinkUnauthorized
const PostG2pMapperOnLinkUnauthorizedCode int = 401

/*
PostG2pMapperOnLinkUnauthorized HTTP layer error details

swagger:response postG2pMapperOnLinkUnauthorized
*/
type PostG2pMapperOnLinkUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *PostG2pMapperOnLinkUnauthorizedBody `json:"body,omitempty"`
}

// NewPostG2pMapperOnLinkUnauthorized creates PostG2pMapperOnLinkUnauthorized with default headers values
func NewPostG2pMapperOnLinkUnauthorized() *PostG2pMapperOnLinkUnauthorized {

	return &PostG2pMapperOnLinkUnauthorized{}
}

// WithPayload adds the payload to the post g2p mapper on link unauthorized response
func (o *PostG2pMapperOnLinkUnauthorized) WithPayload(payload *PostG2pMapperOnLinkUnauthorizedBody) *PostG2pMapperOnLinkUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p mapper on link unauthorized response
func (o *PostG2pMapperOnLinkUnauthorized) SetPayload(payload *PostG2pMapperOnLinkUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pMapperOnLinkUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostG2pMapperOnLinkForbiddenCode is the HTTP code returned for type PostG2pMapperOnLinkForbidden
const PostG2pMapperOnLinkForbiddenCode int = 403

/*
PostG2pMapperOnLinkForbidden HTTP layer error details

swagger:response postG2pMapperOnLinkForbidden
*/
type PostG2pMapperOnLinkForbidden struct {

	/*
	  In: Body
	*/
	Payload *PostG2pMapperOnLinkForbiddenBody `json:"body,omitempty"`
}

// NewPostG2pMapperOnLinkForbidden creates PostG2pMapperOnLinkForbidden with default headers values
func NewPostG2pMapperOnLinkForbidden() *PostG2pMapperOnLinkForbidden {

	return &PostG2pMapperOnLinkForbidden{}
}

// WithPayload adds the payload to the post g2p mapper on link forbidden response
func (o *PostG2pMapperOnLinkForbidden) WithPayload(payload *PostG2pMapperOnLinkForbiddenBody) *PostG2pMapperOnLinkForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p mapper on link forbidden response
func (o *PostG2pMapperOnLinkForbidden) SetPayload(payload *PostG2pMapperOnLinkForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pMapperOnLinkForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostG2pMapperOnLinkInternalServerErrorCode is the HTTP code returned for type PostG2pMapperOnLinkInternalServerError
const PostG2pMapperOnLinkInternalServerErrorCode int = 500

/*
PostG2pMapperOnLinkInternalServerError HTTP layer error details

swagger:response postG2pMapperOnLinkInternalServerError
*/
type PostG2pMapperOnLinkInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PostG2pMapperOnLinkInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostG2pMapperOnLinkInternalServerError creates PostG2pMapperOnLinkInternalServerError with default headers values
func NewPostG2pMapperOnLinkInternalServerError() *PostG2pMapperOnLinkInternalServerError {

	return &PostG2pMapperOnLinkInternalServerError{}
}

// WithPayload adds the payload to the post g2p mapper on link internal server error response
func (o *PostG2pMapperOnLinkInternalServerError) WithPayload(payload *PostG2pMapperOnLinkInternalServerErrorBody) *PostG2pMapperOnLinkInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p mapper on link internal server error response
func (o *PostG2pMapperOnLinkInternalServerError) SetPayload(payload *PostG2pMapperOnLinkInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pMapperOnLinkInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PostG2pMapperOnLinkDefault Acknowledgement of message received after successful validation of message and signature

swagger:response postG2pMapperOnLinkDefault
*/
type PostG2pMapperOnLinkDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *PostG2pMapperOnLinkDefaultBody `json:"body,omitempty"`
}

// NewPostG2pMapperOnLinkDefault creates PostG2pMapperOnLinkDefault with default headers values
func NewPostG2pMapperOnLinkDefault(code int) *PostG2pMapperOnLinkDefault {
	if code <= 0 {
		code = 500
	}

	return &PostG2pMapperOnLinkDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post g2p mapper on link default response
func (o *PostG2pMapperOnLinkDefault) WithStatusCode(code int) *PostG2pMapperOnLinkDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post g2p mapper on link default response
func (o *PostG2pMapperOnLinkDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post g2p mapper on link default response
func (o *PostG2pMapperOnLinkDefault) WithPayload(payload *PostG2pMapperOnLinkDefaultBody) *PostG2pMapperOnLinkDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p mapper on link default response
func (o *PostG2pMapperOnLinkDefault) SetPayload(payload *PostG2pMapperOnLinkDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pMapperOnLinkDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
