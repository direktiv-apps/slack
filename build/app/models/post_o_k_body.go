// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostOKBody post o k body
//
// swagger:model postOKBody
type PostOKBody struct {

	// slack
	Slack *PostOKBodySlack `json:"slack,omitempty"`
}

// Validate validates this post o k body
func (m *PostOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSlack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostOKBody) validateSlack(formats strfmt.Registry) error {
	if swag.IsZero(m.Slack) { // not required
		return nil
	}

	if m.Slack != nil {
		if err := m.Slack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slack")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post o k body based on the context it is used
func (m *PostOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSlack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostOKBody) contextValidateSlack(ctx context.Context, formats strfmt.Registry) error {

	if m.Slack != nil {
		if err := m.Slack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slack")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostOKBody) UnmarshalBinary(b []byte) error {
	var res PostOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
