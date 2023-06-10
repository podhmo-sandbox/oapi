// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: task API
//	Version: 0.0.0
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Error is an object. default error
type Error struct {
	// Message:
	Message string `json:"message" mapstructure:"message"`
}

// Validate implements basic validation for this model
func (m Error) Validate() error {
	return validation.Errors{}.Filter()
}

// GetMessage returns the Message property
func (m Error) GetMessage() string {
	return m.Message
}

// SetMessage sets the Message property
func (m *Error) SetMessage(val string) {
	m.Message = val
}
