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

// ModelsDetectionEngineInfoType models detection engine info type
//
// swagger:model models.DetectionEngineInfoType
type ModelsDetectionEngineInfoType struct {

	// apk static version
	// Required: true
	ApkStaticVersion *string `json:"ApkStaticVersion"`

	// engine version
	// Required: true
	EngineVersion *string `json:"EngineVersion"`

	// performed at
	// Required: true
	// Format: date-time
	PerformedAt *strfmt.DateTime `json:"PerformedAt"`
}

// Validate validates this models detection engine info type
func (m *ModelsDetectionEngineInfoType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApkStaticVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEngineVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsDetectionEngineInfoType) validateApkStaticVersion(formats strfmt.Registry) error {

	if err := validate.Required("ApkStaticVersion", "body", m.ApkStaticVersion); err != nil {
		return err
	}

	return nil
}

func (m *ModelsDetectionEngineInfoType) validateEngineVersion(formats strfmt.Registry) error {

	if err := validate.Required("EngineVersion", "body", m.EngineVersion); err != nil {
		return err
	}

	return nil
}

func (m *ModelsDetectionEngineInfoType) validatePerformedAt(formats strfmt.Registry) error {

	if err := validate.Required("PerformedAt", "body", m.PerformedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("PerformedAt", "body", "date-time", m.PerformedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models detection engine info type based on context it is used
func (m *ModelsDetectionEngineInfoType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsDetectionEngineInfoType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsDetectionEngineInfoType) UnmarshalBinary(b []byte) error {
	var res ModelsDetectionEngineInfoType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
