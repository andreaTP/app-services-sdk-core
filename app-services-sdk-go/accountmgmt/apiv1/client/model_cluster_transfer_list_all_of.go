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

// ClusterTransferListAllOf struct for ClusterTransferListAllOf
type ClusterTransferListAllOf struct {
	Items []ClusterTransfer `json:"items,omitempty"`
}

// NewClusterTransferListAllOf instantiates a new ClusterTransferListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterTransferListAllOf() *ClusterTransferListAllOf {
	this := ClusterTransferListAllOf{}
	return &this
}

// NewClusterTransferListAllOfWithDefaults instantiates a new ClusterTransferListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterTransferListAllOfWithDefaults() *ClusterTransferListAllOf {
	this := ClusterTransferListAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ClusterTransferListAllOf) GetItems() []ClusterTransfer {
	if o == nil || isNil(o.Items) {
		var ret []ClusterTransfer
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTransferListAllOf) GetItemsOk() ([]ClusterTransfer, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ClusterTransferListAllOf) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ClusterTransfer and assigns it to the Items field.
func (o *ClusterTransferListAllOf) SetItems(v []ClusterTransfer) {
	o.Items = v
}

func (o ClusterTransferListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableClusterTransferListAllOf struct {
	value *ClusterTransferListAllOf
	isSet bool
}

func (v NullableClusterTransferListAllOf) Get() *ClusterTransferListAllOf {
	return v.value
}

func (v *NullableClusterTransferListAllOf) Set(val *ClusterTransferListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterTransferListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterTransferListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterTransferListAllOf(val *ClusterTransferListAllOf) *NullableClusterTransferListAllOf {
	return &NullableClusterTransferListAllOf{value: val, isSet: true}
}

func (v NullableClusterTransferListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterTransferListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


