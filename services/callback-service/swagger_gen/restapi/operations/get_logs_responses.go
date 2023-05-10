// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetLogsOKCode is the HTTP code returned for type GetLogsOK
const GetLogsOKCode int = 200

/*
GetLogsOK Health response

swagger:response getLogsOK
*/
type GetLogsOK struct {

	/*
	  In: Body
	*/
	Payload *GetLogsOKBody `json:"body,omitempty"`
}

// NewGetLogsOK creates GetLogsOK with default headers values
func NewGetLogsOK() *GetLogsOK {

	return &GetLogsOK{}
}

// WithPayload adds the payload to the get logs o k response
func (o *GetLogsOK) WithPayload(payload *GetLogsOKBody) *GetLogsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get logs o k response
func (o *GetLogsOK) SetPayload(payload *GetLogsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}