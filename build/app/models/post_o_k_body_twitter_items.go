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

// PostOKBodyTwitterItems post o k body twitter items
//
// swagger:model postOKBodyTwitterItems
type PostOKBodyTwitterItems struct {

	// result
	// Required: true
	Result interface{} `json:"result"`

	// success
	// Required: true
	Success *bool `json:"success"`
}

// Validate validates this post o k body twitter items
func (m *PostOKBodyTwitterItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostOKBodyTwitterItems) validateResult(formats strfmt.Registry) error {

	if m.Result == nil {
		return errors.Required("result", "body", nil)
	}

	return nil
}

func (m *PostOKBodyTwitterItems) validateSuccess(formats strfmt.Registry) error {

	if err := validate.Required("success", "body", m.Success); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post o k body twitter items based on context it is used
func (m *PostOKBodyTwitterItems) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostOKBodyTwitterItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostOKBodyTwitterItems) UnmarshalBinary(b []byte) error {
	var res PostOKBodyTwitterItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}