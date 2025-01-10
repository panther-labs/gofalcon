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

// DomainKestrelParams domain kestrel params
//
// swagger:model domain.KestrelParams
type DomainKestrelParams struct {

	// view id
	// Required: true
	ViewID *string `json:"view_id"`
}

// Validate validates this domain kestrel params
func (m *DomainKestrelParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateViewID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainKestrelParams) validateViewID(formats strfmt.Registry) error {

	if err := validate.Required("view_id", "body", m.ViewID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain kestrel params based on context it is used
func (m *DomainKestrelParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainKestrelParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainKestrelParams) UnmarshalBinary(b []byte) error {
	var res DomainKestrelParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
