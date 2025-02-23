/*
Account Management Service API

Manage user subscriptions and clusters

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"encoding/json"
)

// ExportControlReview struct for ExportControlReview
type ExportControlReview struct {
	Restricted bool `json:"restricted"`
}

// NewExportControlReview instantiates a new ExportControlReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportControlReview(restricted bool) *ExportControlReview {
	this := ExportControlReview{}
	this.Restricted = restricted
	return &this
}

// NewExportControlReviewWithDefaults instantiates a new ExportControlReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportControlReviewWithDefaults() *ExportControlReview {
	this := ExportControlReview{}
	return &this
}

// GetRestricted returns the Restricted field value
func (o *ExportControlReview) GetRestricted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Restricted
}

// GetRestrictedOk returns a tuple with the Restricted field value
// and a boolean to check if the value has been set.
func (o *ExportControlReview) GetRestrictedOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Restricted, true
}

// SetRestricted sets field value
func (o *ExportControlReview) SetRestricted(v bool) {
	o.Restricted = v
}

func (o ExportControlReview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["restricted"] = o.Restricted
	}
	return json.Marshal(toSerialize)
}

type NullableExportControlReview struct {
	value *ExportControlReview
	isSet bool
}

func (v NullableExportControlReview) Get() *ExportControlReview {
	return v.value
}

func (v *NullableExportControlReview) Set(val *ExportControlReview) {
	v.value = val
	v.isSet = true
}

func (v NullableExportControlReview) IsSet() bool {
	return v.isSet
}

func (v *NullableExportControlReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportControlReview(val *ExportControlReview) *NullableExportControlReview {
	return &NullableExportControlReview{value: val, isSet: true}
}

func (v NullableExportControlReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportControlReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


