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

// NamedLogConfigurationAllOf struct for NamedLogConfigurationAllOf
type NamedLogConfigurationAllOf struct {
	// 
	Name string `json:"name"`
}

// NewNamedLogConfigurationAllOf instantiates a new NamedLogConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamedLogConfigurationAllOf(name string) *NamedLogConfigurationAllOf {
	this := NamedLogConfigurationAllOf{}
	this.Name = name
	return &this
}

// NewNamedLogConfigurationAllOfWithDefaults instantiates a new NamedLogConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamedLogConfigurationAllOfWithDefaults() *NamedLogConfigurationAllOf {
	this := NamedLogConfigurationAllOf{}
	return &this
}

// GetName returns the Name field value
func (o *NamedLogConfigurationAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NamedLogConfigurationAllOf) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NamedLogConfigurationAllOf) SetName(v string) {
	o.Name = v
}

func (o NamedLogConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNamedLogConfigurationAllOf struct {
	value *NamedLogConfigurationAllOf
	isSet bool
}

func (v NullableNamedLogConfigurationAllOf) Get() *NamedLogConfigurationAllOf {
	return v.value
}

func (v *NullableNamedLogConfigurationAllOf) Set(val *NamedLogConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNamedLogConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNamedLogConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamedLogConfigurationAllOf(val *NamedLogConfigurationAllOf) *NullableNamedLogConfigurationAllOf {
	return &NullableNamedLogConfigurationAllOf{value: val, isSet: true}
}

func (v NullableNamedLogConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamedLogConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


