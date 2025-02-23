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

// ConnectorClusterRequestMeta struct for ConnectorClusterRequestMeta
type ConnectorClusterRequestMeta struct {
	Name *string `json:"name,omitempty"`
	// Name-value string annotations for resource
	Annotations *map[string]string `json:"annotations,omitempty"`
}

// NewConnectorClusterRequestMeta instantiates a new ConnectorClusterRequestMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorClusterRequestMeta() *ConnectorClusterRequestMeta {
	this := ConnectorClusterRequestMeta{}
	return &this
}

// NewConnectorClusterRequestMetaWithDefaults instantiates a new ConnectorClusterRequestMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorClusterRequestMetaWithDefaults() *ConnectorClusterRequestMeta {
	this := ConnectorClusterRequestMeta{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectorClusterRequestMeta) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorClusterRequestMeta) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectorClusterRequestMeta) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectorClusterRequestMeta) SetName(v string) {
	o.Name = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *ConnectorClusterRequestMeta) GetAnnotations() map[string]string {
	if o == nil || isNil(o.Annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorClusterRequestMeta) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Annotations) {
    return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *ConnectorClusterRequestMeta) HasAnnotations() bool {
	if o != nil && !isNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *ConnectorClusterRequestMeta) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

func (o ConnectorClusterRequestMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorClusterRequestMeta struct {
	value *ConnectorClusterRequestMeta
	isSet bool
}

func (v NullableConnectorClusterRequestMeta) Get() *ConnectorClusterRequestMeta {
	return v.value
}

func (v *NullableConnectorClusterRequestMeta) Set(val *ConnectorClusterRequestMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorClusterRequestMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorClusterRequestMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorClusterRequestMeta(val *ConnectorClusterRequestMeta) *NullableConnectorClusterRequestMeta {
	return &NullableConnectorClusterRequestMeta{value: val, isSet: true}
}

func (v NullableConnectorClusterRequestMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorClusterRequestMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


