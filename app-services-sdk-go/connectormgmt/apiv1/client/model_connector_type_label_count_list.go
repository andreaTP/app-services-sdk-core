/*
Connector Management API

Connector Management API is a REST API to manage connectors.

API version: 0.1.0
Contact: rhosak-support@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
)

// ConnectorTypeLabelCountList struct for ConnectorTypeLabelCountList
type ConnectorTypeLabelCountList struct {
	Items []ConnectorTypeLabelCount `json:"items,omitempty"`
}

// NewConnectorTypeLabelCountList instantiates a new ConnectorTypeLabelCountList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorTypeLabelCountList() *ConnectorTypeLabelCountList {
	this := ConnectorTypeLabelCountList{}
	return &this
}

// NewConnectorTypeLabelCountListWithDefaults instantiates a new ConnectorTypeLabelCountList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorTypeLabelCountListWithDefaults() *ConnectorTypeLabelCountList {
	this := ConnectorTypeLabelCountList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ConnectorTypeLabelCountList) GetItems() []ConnectorTypeLabelCount {
	if o == nil || isNil(o.Items) {
		var ret []ConnectorTypeLabelCount
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorTypeLabelCountList) GetItemsOk() ([]ConnectorTypeLabelCount, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConnectorTypeLabelCountList) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ConnectorTypeLabelCount and assigns it to the Items field.
func (o *ConnectorTypeLabelCountList) SetItems(v []ConnectorTypeLabelCount) {
	o.Items = v
}

func (o ConnectorTypeLabelCountList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorTypeLabelCountList struct {
	value *ConnectorTypeLabelCountList
	isSet bool
}

func (v NullableConnectorTypeLabelCountList) Get() *ConnectorTypeLabelCountList {
	return v.value
}

func (v *NullableConnectorTypeLabelCountList) Set(val *ConnectorTypeLabelCountList) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorTypeLabelCountList) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorTypeLabelCountList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorTypeLabelCountList(val *ConnectorTypeLabelCountList) *NullableConnectorTypeLabelCountList {
	return &NullableConnectorTypeLabelCountList{value: val, isSet: true}
}

func (v NullableConnectorTypeLabelCountList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorTypeLabelCountList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


