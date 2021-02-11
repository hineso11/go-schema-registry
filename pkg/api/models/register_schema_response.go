// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RegisterSchemaResponse register schema response
//
// swagger:model RegisterSchemaResponse
type RegisterSchemaResponse struct {

	// id
	ID int32 `json:"id,omitempty"`
}

// Validate validates this register schema response
func (m *RegisterSchemaResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this register schema response based on context it is used
func (m *RegisterSchemaResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegisterSchemaResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisterSchemaResponse) UnmarshalBinary(b []byte) error {
	var res RegisterSchemaResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
