// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CIDRBlockIdentity CIDRBlockIdentity
// swagger:model CIDRBlockIdentity
type CIDRBlockIdentity struct {

	// A range of IPv4 addresses, in CIDR format.
	// Required: true
	CidrBlock *string `json:"cidr_block"`
}

// Validate validates this c ID r block identity
func (m *CIDRBlockIdentity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidrBlock(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CIDRBlockIdentity) validateCidrBlock(formats strfmt.Registry) error {

	if err := validate.Required("cidr_block", "body", m.CidrBlock); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CIDRBlockIdentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CIDRBlockIdentity) UnmarshalBinary(b []byte) error {
	var res CIDRBlockIdentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
