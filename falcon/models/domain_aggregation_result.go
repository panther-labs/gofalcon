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

// DomainAggregationResult domain aggregation result
//
// swagger:model domain.AggregationResult
type DomainAggregationResult struct {

	// Aggregation buckets containing aggregated information
	// Required: true
	Buckets []*DomainAggregationResultItem `json:"buckets"`

	// Name of the aggregation, equal to the name of it from the AggregateQueryRequest name
	// Required: true
	Name *string `json:"name"`

	// sum count of the documents not included in the buckets
	SumOtherDocCount int64 `json:"sum_other_doc_count,omitempty"`
}

// Validate validates this domain aggregation result
func (m *DomainAggregationResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuckets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAggregationResult) validateBuckets(formats strfmt.Registry) error {

	if err := validate.Required("buckets", "body", m.Buckets); err != nil {
		return err
	}

	for i := 0; i < len(m.Buckets); i++ {
		if swag.IsZero(m.Buckets[i]) { // not required
			continue
		}

		if m.Buckets[i] != nil {
			if err := m.Buckets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("buckets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("buckets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainAggregationResult) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain aggregation result based on the context it is used
func (m *DomainAggregationResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuckets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAggregationResult) contextValidateBuckets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Buckets); i++ {

		if m.Buckets[i] != nil {

			if swag.IsZero(m.Buckets[i]) { // not required
				return nil
			}

			if err := m.Buckets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("buckets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("buckets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainAggregationResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAggregationResult) UnmarshalBinary(b []byte) error {
	var res DomainAggregationResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
