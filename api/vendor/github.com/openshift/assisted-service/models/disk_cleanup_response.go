// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DiskCleanupResponse disk cleanup response
//
// swagger:model disk_cleanup_response
type DiskCleanupResponse struct {

	// The device path.
	Path string `json:"path,omitempty"`

	// Result of the cleanup operation.
	Successful bool `json:"successful,omitempty"`
}

// Validate validates this disk cleanup response
func (m *DiskCleanupResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this disk cleanup response based on context it is used
func (m *DiskCleanupResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DiskCleanupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiskCleanupResponse) UnmarshalBinary(b []byte) error {
	var res DiskCleanupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
