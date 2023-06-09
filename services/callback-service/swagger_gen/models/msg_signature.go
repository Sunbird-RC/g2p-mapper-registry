// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// MsgSignature Signature of {header}+{message} body verified using sender's signing public key
// Example: Signature:  namespace=\"g2p\", kidId=\"{sender_id}|{unique_key_id}|{algorithm}\", algorithm=\"ed25519\", created=\"1606970629\", expires=\"1607030629\", headers=\"(created) (expires) digest\", signature=\"Base64(signing content)
//
// swagger:model MsgSignature
type MsgSignature string

// Validate validates this msg signature
func (m MsgSignature) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg signature based on context it is used
func (m MsgSignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
