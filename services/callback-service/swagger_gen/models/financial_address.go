// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// FinancialAddress <br> 1. Financial address is case insensitive normative represenation of a store of value account represented as id-type:id@provider <br> 2. Every payer/payee financial address must resolve to an actual store of value account number for processing the payment instruction <br> 3. It is recommended the mapping between id and store of value account details to be held only at final store of value entity and intermediaries can hold  3. Few examples: <br> - token@id-provider e.g token:12345@mosip <br> - uid@pymt-rail e.g uid:12345@mosip <br> - vid@id-provider e.g vid:12345@PhilID <br> - mobile@mobile-provider e.g mobile:12345@m-pesa <br> - account-id@bank-psp-code e.g account:12345@gtbank <br> - account-no@ifsc-code.ifsc.npci e.g account:12345@HDFC0000001.ifsc.npci <br> - user-id@psp-code e.g. joeuser@gtbank <br> - token@psp-code e.g token:123456@sbi <br> - code@purpose-code.voucher-provider e.g voucher:12345@food.sodexo <br> - cdbc-id@cdbc e.g. 12345@DCash
// Example: token:12345@gtbank
//
// swagger:model FinancialAddress
type FinancialAddress string

// Validate validates this financial address
func (m FinancialAddress) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this financial address based on context it is used
func (m FinancialAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
