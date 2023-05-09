// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SearchResponse Digital credential
//
// swagger:model SearchResponse
type SearchResponse struct {

	// locale
	Locale LanguageCode `json:"locale,omitempty"`

	// mapper
	Mapper []*ResolveResponse `json:"mapper"`

	// reference id
	// Required: true
	ReferenceID *ReferenceID `json:"reference_id"`

	// status
	// Required: true
	Status *RequestStatus `json:"status"`

	// status reason code
	StatusReasonCode SearchStatusReasonCode `json:"status_reason_code,omitempty"`

	// Status reason code message. Helps actionanble messaging for systems/end users
	// Max Length: 999
	StatusReasonMessage string `json:"status_reason_message,omitempty"`

	// timestamp
	// Required: true
	// Format: date-time
	Timestamp *Timestamp `json:"timestamp"`
}

// Validate validates this search response
func (m *SearchResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapper(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusReasonCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusReasonMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchResponse) validateLocale(formats strfmt.Registry) error {
	if swag.IsZero(m.Locale) { // not required
		return nil
	}

	if err := m.Locale.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("locale")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("locale")
		}
		return err
	}

	return nil
}

func (m *SearchResponse) validateMapper(formats strfmt.Registry) error {
	if swag.IsZero(m.Mapper) { // not required
		return nil
	}

	for i := 0; i < len(m.Mapper); i++ {
		if swag.IsZero(m.Mapper[i]) { // not required
			continue
		}

		if m.Mapper[i] != nil {
			if err := m.Mapper[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mapper" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mapper" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SearchResponse) validateReferenceID(formats strfmt.Registry) error {

	if err := validate.Required("reference_id", "body", m.ReferenceID); err != nil {
		return err
	}

	if err := validate.Required("reference_id", "body", m.ReferenceID); err != nil {
		return err
	}

	if m.ReferenceID != nil {
		if err := m.ReferenceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reference_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reference_id")
			}
			return err
		}
	}

	return nil
}

func (m *SearchResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *SearchResponse) validateStatusReasonCode(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusReasonCode) { // not required
		return nil
	}

	if err := m.StatusReasonCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status_reason_code")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status_reason_code")
		}
		return err
	}

	return nil
}

func (m *SearchResponse) validateStatusReasonMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusReasonMessage) { // not required
		return nil
	}

	if err := validate.MaxLength("status_reason_message", "body", m.StatusReasonMessage, 999); err != nil {
		return err
	}

	return nil
}

func (m *SearchResponse) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if m.Timestamp != nil {
		if err := m.Timestamp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timestamp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timestamp")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this search response based on the context it is used
func (m *SearchResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocale(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMapper(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReferenceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusReasonCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchResponse) contextValidateLocale(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Locale.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("locale")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("locale")
		}
		return err
	}

	return nil
}

func (m *SearchResponse) contextValidateMapper(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Mapper); i++ {

		if m.Mapper[i] != nil {
			if err := m.Mapper[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mapper" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mapper" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SearchResponse) contextValidateReferenceID(ctx context.Context, formats strfmt.Registry) error {

	if m.ReferenceID != nil {
		if err := m.ReferenceID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reference_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reference_id")
			}
			return err
		}
	}

	return nil
}

func (m *SearchResponse) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *SearchResponse) contextValidateStatusReasonCode(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StatusReasonCode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status_reason_code")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status_reason_code")
		}
		return err
	}

	return nil
}

func (m *SearchResponse) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if m.Timestamp != nil {
		if err := m.Timestamp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timestamp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timestamp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchResponse) UnmarshalBinary(b []byte) error {
	var res SearchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
