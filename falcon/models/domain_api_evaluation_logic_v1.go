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

// DomainAPIEvaluationLogicV1 domain API evaluation logic v1
//
// swagger:model domain.APIEvaluationLogicV1
type DomainAPIEvaluationLogicV1 struct {

	// aid
	Aid string `json:"aid,omitempty"`

	// cid
	Cid string `json:"cid,omitempty"`

	// complex check operator
	ComplexCheckOperator string `json:"complex_check_operator,omitempty"`

	// created timestamp
	CreatedTimestamp string `json:"created_timestamp,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// logic
	Logic []*DomainAPIEvaluationLogicItemV1 `json:"logic"`

	// updated timestamp
	UpdatedTimestamp string `json:"updated_timestamp,omitempty"`
}

// Validate validates this domain API evaluation logic v1
func (m *DomainAPIEvaluationLogicV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIEvaluationLogicV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIEvaluationLogicV1) validateLogic(formats strfmt.Registry) error {
	if swag.IsZero(m.Logic) { // not required
		return nil
	}

	for i := 0; i < len(m.Logic); i++ {
		if swag.IsZero(m.Logic[i]) { // not required
			continue
		}

		if m.Logic[i] != nil {
			if err := m.Logic[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logic" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logic" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this domain API evaluation logic v1 based on the context it is used
func (m *DomainAPIEvaluationLogicV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIEvaluationLogicV1) contextValidateLogic(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Logic); i++ {

		if m.Logic[i] != nil {

			if swag.IsZero(m.Logic[i]) { // not required
				return nil
			}

			if err := m.Logic[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logic" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logic" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainAPIEvaluationLogicV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAPIEvaluationLogicV1) UnmarshalBinary(b []byte) error {
	var res DomainAPIEvaluationLogicV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
