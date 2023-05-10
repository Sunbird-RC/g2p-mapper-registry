// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FileInfo File info. Used in file upload feature using HTTPS
//
// swagger:model FileInfo
type FileInfo struct {

	// G2P Connect specific actions. Usually verb from the URI should go here to help store and fwd kind of processing requirements.
	// Required: true
	Action *string `json:"action"`

	// File content format. e.g json, csv, etc.,
	// Example: csv
	FileFormat *string `json:"fileFormat,omitempty"`

	// Disbursement instruction file representing Disburse or OnDisburse end point elements i.e signature/header/message entities as a file record
	// Required: true
	// Format: binary
	FileName io.ReadCloser `json:"fileName"`
}

// Validate validates this file info
func (m *FileInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileInfo) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *FileInfo) validateFileName(formats strfmt.Registry) error {

	if err := validate.Required("fileName", "body", io.ReadCloser(m.FileName)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this file info based on context it is used
func (m *FileInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileInfo) UnmarshalBinary(b []byte) error {
	var res FileInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}