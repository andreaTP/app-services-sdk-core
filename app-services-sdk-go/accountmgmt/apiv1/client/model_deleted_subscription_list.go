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

// DeletedSubscriptionList struct for DeletedSubscriptionList
type DeletedSubscriptionList struct {
	Kind string `json:"kind"`
	Page int32 `json:"page"`
	Size int32 `json:"size"`
	Total int32 `json:"total"`
	Items []DeletedSubscription `json:"items"`
}

// NewDeletedSubscriptionList instantiates a new DeletedSubscriptionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletedSubscriptionList(kind string, page int32, size int32, total int32, items []DeletedSubscription) *DeletedSubscriptionList {
	this := DeletedSubscriptionList{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	this.Items = items
	return &this
}

// NewDeletedSubscriptionListWithDefaults instantiates a new DeletedSubscriptionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletedSubscriptionListWithDefaults() *DeletedSubscriptionList {
	this := DeletedSubscriptionList{}
	return &this
}

// GetKind returns the Kind field value
func (o *DeletedSubscriptionList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *DeletedSubscriptionList) GetKindOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *DeletedSubscriptionList) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value
func (o *DeletedSubscriptionList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *DeletedSubscriptionList) GetPageOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *DeletedSubscriptionList) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *DeletedSubscriptionList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *DeletedSubscriptionList) GetSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *DeletedSubscriptionList) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *DeletedSubscriptionList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *DeletedSubscriptionList) GetTotalOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *DeletedSubscriptionList) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *DeletedSubscriptionList) GetItems() []DeletedSubscription {
	if o == nil {
		var ret []DeletedSubscription
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *DeletedSubscriptionList) GetItemsOk() ([]DeletedSubscription, bool) {
	if o == nil {
    return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *DeletedSubscriptionList) SetItems(v []DeletedSubscription) {
	o.Items = v
}

func (o DeletedSubscriptionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableDeletedSubscriptionList struct {
	value *DeletedSubscriptionList
	isSet bool
}

func (v NullableDeletedSubscriptionList) Get() *DeletedSubscriptionList {
	return v.value
}

func (v *NullableDeletedSubscriptionList) Set(val *DeletedSubscriptionList) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletedSubscriptionList) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletedSubscriptionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletedSubscriptionList(val *DeletedSubscriptionList) *NullableDeletedSubscriptionList {
	return &NullableDeletedSubscriptionList{value: val, isSet: true}
}

func (v NullableDeletedSubscriptionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletedSubscriptionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


