// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostG2pMapperResolveUnauthorizedCode is the HTTP code returned for type PostG2pMapperResolveUnauthorized
const PostG2pMapperResolveUnauthorizedCode int = 401

/*
PostG2pMapperResolveUnauthorized HTTP layer error details

swagger:response postG2pMapperResolveUnauthorized
*/
type PostG2pMapperResolveUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *PostG2pMapperResolveUnauthorizedBody `json:"body,omitempty"`
}

// NewPostG2pMapperResolveUnauthorized creates PostG2pMapperResolveUnauthorized with default headers values
func NewPostG2pMapperResolveUnauthorized() *PostG2pMapperResolveUnauthorized {

	return &PostG2pMapperResolveUnauthorized{}
}

// WithPayload adds the payload to the post g2p mapper resolve unauthorized response
func (o *PostG2pMapperResolveUnauthorized) WithPayload(payload *PostG2pMapperResolveUnauthorizedBody) *PostG2pMapperResolveUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p mapper resolve unauthorized response
func (o *PostG2pMapperResolveUnauthorized) SetPayload(payload *PostG2pMapperResolveUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pMapperResolveUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostG2pMapperResolveForbiddenCode is the HTTP code returned for type PostG2pMapperResolveForbidden
const PostG2pMapperResolveForbiddenCode int = 403

/*
PostG2pMapperResolveForbidden HTTP layer error details

swagger:response postG2pMapperResolveForbidden
*/
type PostG2pMapperResolveForbidden struct {

	/*
	  In: Body
	*/
	Payload *PostG2pMapperResolveForbiddenBody `json:"body,omitempty"`
}

// NewPostG2pMapperResolveForbidden creates PostG2pMapperResolveForbidden with default headers values
func NewPostG2pMapperResolveForbidden() *PostG2pMapperResolveForbidden {

	return &PostG2pMapperResolveForbidden{}
}

// WithPayload adds the payload to the post g2p mapper resolve forbidden response
func (o *PostG2pMapperResolveForbidden) WithPayload(payload *PostG2pMapperResolveForbiddenBody) *PostG2pMapperResolveForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p mapper resolve forbidden response
func (o *PostG2pMapperResolveForbidden) SetPayload(payload *PostG2pMapperResolveForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pMapperResolveForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostG2pMapperResolveInternalServerErrorCode is the HTTP code returned for type PostG2pMapperResolveInternalServerError
const PostG2pMapperResolveInternalServerErrorCode int = 500

/*
PostG2pMapperResolveInternalServerError HTTP layer error details

swagger:response postG2pMapperResolveInternalServerError
*/
type PostG2pMapperResolveInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PostG2pMapperResolveInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostG2pMapperResolveInternalServerError creates PostG2pMapperResolveInternalServerError with default headers values
func NewPostG2pMapperResolveInternalServerError() *PostG2pMapperResolveInternalServerError {

	return &PostG2pMapperResolveInternalServerError{}
}

// WithPayload adds the payload to the post g2p mapper resolve internal server error response
func (o *PostG2pMapperResolveInternalServerError) WithPayload(payload *PostG2pMapperResolveInternalServerErrorBody) *PostG2pMapperResolveInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p mapper resolve internal server error response
func (o *PostG2pMapperResolveInternalServerError) SetPayload(payload *PostG2pMapperResolveInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pMapperResolveInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PostG2pMapperResolveDefault Acknowledgement of message received after successful validation of message and signature

swagger:response postG2pMapperResolveDefault
*/
type PostG2pMapperResolveDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *PostG2pMapperResolveDefaultBody `json:"body,omitempty"`
}

// NewPostG2pMapperResolveDefault creates PostG2pMapperResolveDefault with default headers values
func NewPostG2pMapperResolveDefault(code int) *PostG2pMapperResolveDefault {
	if code <= 0 {
		code = 500
	}

	return &PostG2pMapperResolveDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post g2p mapper resolve default response
func (o *PostG2pMapperResolveDefault) WithStatusCode(code int) *PostG2pMapperResolveDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post g2p mapper resolve default response
func (o *PostG2pMapperResolveDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post g2p mapper resolve default response
func (o *PostG2pMapperResolveDefault) WithPayload(payload *PostG2pMapperResolveDefaultBody) *PostG2pMapperResolveDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p mapper resolve default response
func (o *PostG2pMapperResolveDefault) SetPayload(payload *PostG2pMapperResolveDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pMapperResolveDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
