// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ReferenceID Unique reference_id set by txn initiating system for each request in a batch
// Example: 12345678901234567890
//
// swagger:model ReferenceId
type ReferenceID string

// Validate validates this reference Id
func (m ReferenceID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this reference Id based on context it is used
func (m ReferenceID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
