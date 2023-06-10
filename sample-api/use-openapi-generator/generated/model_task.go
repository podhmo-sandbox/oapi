// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: task API
//	Version: 0.0.0
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

// Task is an object.
type Task struct {
	// CreatedAt:
	CreatedAt DateTime `json:"createdAt" mapstructure:"createdAt"`
	// Done:
	Done bool `json:"done" mapstructure:"done"`
	// Name: name of something
	Name Name `json:"name" mapstructure:"name"`
}

// Validate implements basic validation for this model
func (m Task) Validate() error {
	return validation.Errors{
		"createdAt": validation.Validate(
			m.CreatedAt, validation.Required, validation.Date(time.RFC3339),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 0),
		),
	}.Filter()
}

// GetCreatedAt returns the CreatedAt property
func (m Task) GetCreatedAt() DateTime {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m *Task) SetCreatedAt(val DateTime) {
	m.CreatedAt = val
}

// GetDone returns the Done property
func (m Task) GetDone() bool {
	return m.Done
}

// SetDone sets the Done property
func (m *Task) SetDone(val bool) {
	m.Done = val
}

// GetName returns the Name property
func (m Task) GetName() Name {
	return m.Name
}

// SetName sets the Name property
func (m *Task) SetName(val Name) {
	m.Name = val
}
