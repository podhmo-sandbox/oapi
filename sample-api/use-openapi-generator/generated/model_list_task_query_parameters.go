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

// ListTaskQueryParameters is an object.
type ListTaskQueryParameters struct {
	// Sort:
	Sort string `json:"sort" mapstructure:"sort"`
}

// Validate implements basic validation for this model
func (m ListTaskQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetSort returns the Sort property
func (m ListTaskQueryParameters) GetSort() string {
	return m.Sort
}

// SetSort sets the Sort property
func (m *ListTaskQueryParameters) SetSort(val string) {
	m.Sort = val
}
