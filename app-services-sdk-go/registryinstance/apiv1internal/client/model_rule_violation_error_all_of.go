/*
Service Registry API

Service Registry Instance API  NOTE: This API cannot be called directly from the portal.

API version: 2.2.5.Final
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
)

// RuleViolationErrorAllOf struct for RuleViolationErrorAllOf
type RuleViolationErrorAllOf struct {
	// List of rule violation causes.
	Causes []RuleViolationCause `json:"causes"`
}

// NewRuleViolationErrorAllOf instantiates a new RuleViolationErrorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleViolationErrorAllOf(causes []RuleViolationCause) *RuleViolationErrorAllOf {
	this := RuleViolationErrorAllOf{}
	this.Causes = causes
	return &this
}

// NewRuleViolationErrorAllOfWithDefaults instantiates a new RuleViolationErrorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleViolationErrorAllOfWithDefaults() *RuleViolationErrorAllOf {
	this := RuleViolationErrorAllOf{}
	return &this
}

// GetCauses returns the Causes field value
func (o *RuleViolationErrorAllOf) GetCauses() []RuleViolationCause {
	if o == nil {
		var ret []RuleViolationCause
		return ret
	}

	return o.Causes
}

// GetCausesOk returns a tuple with the Causes field value
// and a boolean to check if the value has been set.
func (o *RuleViolationErrorAllOf) GetCausesOk() ([]RuleViolationCause, bool) {
	if o == nil {
    return nil, false
	}
	return o.Causes, true
}

// SetCauses sets field value
func (o *RuleViolationErrorAllOf) SetCauses(v []RuleViolationCause) {
	o.Causes = v
}

func (o RuleViolationErrorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["causes"] = o.Causes
	}
	return json.Marshal(toSerialize)
}

type NullableRuleViolationErrorAllOf struct {
	value *RuleViolationErrorAllOf
	isSet bool
}

func (v NullableRuleViolationErrorAllOf) Get() *RuleViolationErrorAllOf {
	return v.value
}

func (v *NullableRuleViolationErrorAllOf) Set(val *RuleViolationErrorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleViolationErrorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleViolationErrorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleViolationErrorAllOf(val *RuleViolationErrorAllOf) *NullableRuleViolationErrorAllOf {
	return &NullableRuleViolationErrorAllOf{value: val, isSet: true}
}

func (v NullableRuleViolationErrorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleViolationErrorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


