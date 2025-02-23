/*
Kafka Instance API

API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs

API version: 0.13.0-SNAPSHOT
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// ListDeprecatedAllOf struct for ListDeprecatedAllOf
type ListDeprecatedAllOf struct {
	// Offset of the first record returned, zero-based
	// Deprecated
	Offset *int32 `json:"offset,omitempty"`
	// Maximum number of records to return, from request
	// Deprecated
	Limit *int32 `json:"limit,omitempty"`
	// Total number of entries in the full result set
	// Deprecated
	Count *int32 `json:"count,omitempty"`
}

// NewListDeprecatedAllOf instantiates a new ListDeprecatedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDeprecatedAllOf() *ListDeprecatedAllOf {
	this := ListDeprecatedAllOf{}
	return &this
}

// NewListDeprecatedAllOfWithDefaults instantiates a new ListDeprecatedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDeprecatedAllOfWithDefaults() *ListDeprecatedAllOf {
	this := ListDeprecatedAllOf{}
	return &this
}

// GetOffset returns the Offset field value if set, zero value otherwise.
// Deprecated
func (o *ListDeprecatedAllOf) GetOffset() int32 {
	if o == nil || isNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ListDeprecatedAllOf) GetOffsetOk() (*int32, bool) {
	if o == nil || isNil(o.Offset) {
    return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListDeprecatedAllOf) HasOffset() bool {
	if o != nil && !isNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
// Deprecated
func (o *ListDeprecatedAllOf) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
// Deprecated
func (o *ListDeprecatedAllOf) GetLimit() int32 {
	if o == nil || isNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ListDeprecatedAllOf) GetLimitOk() (*int32, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListDeprecatedAllOf) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
// Deprecated
func (o *ListDeprecatedAllOf) SetLimit(v int32) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
// Deprecated
func (o *ListDeprecatedAllOf) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ListDeprecatedAllOf) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListDeprecatedAllOf) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
// Deprecated
func (o *ListDeprecatedAllOf) SetCount(v int32) {
	o.Count = &v
}

func (o ListDeprecatedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableListDeprecatedAllOf struct {
	value *ListDeprecatedAllOf
	isSet bool
}

func (v NullableListDeprecatedAllOf) Get() *ListDeprecatedAllOf {
	return v.value
}

func (v *NullableListDeprecatedAllOf) Set(val *ListDeprecatedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListDeprecatedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListDeprecatedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDeprecatedAllOf(val *ListDeprecatedAllOf) *NullableListDeprecatedAllOf {
	return &NullableListDeprecatedAllOf{value: val, isSet: true}
}

func (v NullableListDeprecatedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDeprecatedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


