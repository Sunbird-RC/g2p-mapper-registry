// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KeyValueInfo Key Value info
//
// swagger:model KeyValueInfo
type KeyValueInfo struct {

	// key
	// Required: true
	// Max Length: 99
	Key *string `json:"key"`

	// value
	// Required: true
	// Max Length: 999
	Value *string `json:"value"`
}

// Validate validates this key value info
func (m *KeyValueInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyValueInfo) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	if err := validate.MaxLength("key", "body", *m.Key, 99); err != nil {
		return err
	}

	return nil
}

func (m *KeyValueInfo) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if err := validate.MaxLength("value", "body", *m.Value, 999); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this key value info based on context it is used
func (m *KeyValueInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KeyValueInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyValueInfo) UnmarshalBinary(b []byte) error {
	var res KeyValueInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}