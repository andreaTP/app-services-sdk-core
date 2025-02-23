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

// MetricsListAllOf struct for MetricsListAllOf
type MetricsListAllOf struct {
	Items []Metric `json:"items,omitempty"`
}

// NewMetricsListAllOf instantiates a new MetricsListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsListAllOf() *MetricsListAllOf {
	this := MetricsListAllOf{}
	return &this
}

// NewMetricsListAllOfWithDefaults instantiates a new MetricsListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsListAllOfWithDefaults() *MetricsListAllOf {
	this := MetricsListAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *MetricsListAllOf) GetItems() []Metric {
	if o == nil || isNil(o.Items) {
		var ret []Metric
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsListAllOf) GetItemsOk() ([]Metric, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *MetricsListAllOf) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Metric and assigns it to the Items field.
func (o *MetricsListAllOf) SetItems(v []Metric) {
	o.Items = v
}

func (o MetricsListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableMetricsListAllOf struct {
	value *MetricsListAllOf
	isSet bool
}

func (v NullableMetricsListAllOf) Get() *MetricsListAllOf {
	return v.value
}

func (v *NullableMetricsListAllOf) Set(val *MetricsListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsListAllOf(val *MetricsListAllOf) *NullableMetricsListAllOf {
	return &NullableMetricsListAllOf{value: val, isSet: true}
}

func (v NullableMetricsListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


