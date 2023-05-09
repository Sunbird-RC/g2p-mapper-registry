// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Ack 1. ACK: If the request is valid (for basic checks) and async callback (i.e webhook) will be invoked by reciever back to the sender.
// 2. NACK: If the request is valid (for basic checks) and there is no futher updates from reciever back to the sender.
// 3. ERR: If the reuqest is invalid and reciver can't process the request. error object holds error code, message.
//
// swagger:model Ack
type Ack string

func NewAck(value Ack) *Ack {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Ack.
func (m Ack) Pointer() *Ack {
	return &m
}

const (

	// AckACK captures enum value "ACK"
	AckACK Ack = "ACK"

	// AckNACK captures enum value "NACK"
	AckNACK Ack = "NACK"

	// AckERR captures enum value "ERR"
	AckERR Ack = "ERR"
)

// for schema
var ackEnum []interface{}

func init() {
	var res []Ack
	if err := json.Unmarshal([]byte(`["ACK","NACK","ERR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ackEnum = append(ackEnum, v)
	}
}

func (m Ack) validateAckEnum(path, location string, value Ack) error {
	if err := validate.EnumCase(path, location, value, ackEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this ack
func (m Ack) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAckEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this ack based on context it is used
func (m Ack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
