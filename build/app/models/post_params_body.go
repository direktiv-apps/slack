// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostParamsBody post params body
//
// swagger:model postParamsBody
type PostParamsBody struct {

	// content
	Content interface{} `json:"content,omitempty"`

	// URL for Slack's incoming webhook
	WebhookURL string `json:"webhook-url,omitempty"`
}

// Validate validates this post params body
func (m *PostParamsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post params body based on context it is used
func (m *PostParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostParamsBody) UnmarshalBinary(b []byte) error {
	var res PostParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}