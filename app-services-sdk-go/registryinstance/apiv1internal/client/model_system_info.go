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

// SystemInfo 
type SystemInfo struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Version *string `json:"version,omitempty"`
	BuiltOn *string `json:"builtOn,omitempty"`
}

// NewSystemInfo instantiates a new SystemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInfo() *SystemInfo {
	this := SystemInfo{}
	return &this
}

// NewSystemInfoWithDefaults instantiates a new SystemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInfoWithDefaults() *SystemInfo {
	this := SystemInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SystemInfo) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SystemInfo) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SystemInfo) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SystemInfo) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SystemInfo) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SystemInfo) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SystemInfo) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SystemInfo) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SystemInfo) SetVersion(v string) {
	o.Version = &v
}

// GetBuiltOn returns the BuiltOn field value if set, zero value otherwise.
func (o *SystemInfo) GetBuiltOn() string {
	if o == nil || isNil(o.BuiltOn) {
		var ret string
		return ret
	}
	return *o.BuiltOn
}

// GetBuiltOnOk returns a tuple with the BuiltOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetBuiltOnOk() (*string, bool) {
	if o == nil || isNil(o.BuiltOn) {
    return nil, false
	}
	return o.BuiltOn, true
}

// HasBuiltOn returns a boolean if a field has been set.
func (o *SystemInfo) HasBuiltOn() bool {
	if o != nil && !isNil(o.BuiltOn) {
		return true
	}

	return false
}

// SetBuiltOn gets a reference to the given string and assigns it to the BuiltOn field.
func (o *SystemInfo) SetBuiltOn(v string) {
	o.BuiltOn = &v
}

func (o SystemInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.BuiltOn) {
		toSerialize["builtOn"] = o.BuiltOn
	}
	return json.Marshal(toSerialize)
}

type NullableSystemInfo struct {
	value *SystemInfo
	isSet bool
}

func (v NullableSystemInfo) Get() *SystemInfo {
	return v.value
}

func (v *NullableSystemInfo) Set(val *SystemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInfo(val *SystemInfo) *NullableSystemInfo {
	return &NullableSystemInfo{value: val, isSet: true}
}

func (v NullableSystemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


