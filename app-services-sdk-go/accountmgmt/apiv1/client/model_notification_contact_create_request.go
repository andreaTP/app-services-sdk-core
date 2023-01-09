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

// NotificationContactCreateRequest struct for NotificationContactCreateRequest
type NotificationContactCreateRequest struct {
	AccountIdentifier *string `json:"account_identifier,omitempty"`
}

// NewNotificationContactCreateRequest instantiates a new NotificationContactCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationContactCreateRequest() *NotificationContactCreateRequest {
	this := NotificationContactCreateRequest{}
	return &this
}

// NewNotificationContactCreateRequestWithDefaults instantiates a new NotificationContactCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationContactCreateRequestWithDefaults() *NotificationContactCreateRequest {
	this := NotificationContactCreateRequest{}
	return &this
}

// GetAccountIdentifier returns the AccountIdentifier field value if set, zero value otherwise.
func (o *NotificationContactCreateRequest) GetAccountIdentifier() string {
	if o == nil || isNil(o.AccountIdentifier) {
		var ret string
		return ret
	}
	return *o.AccountIdentifier
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationContactCreateRequest) GetAccountIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.AccountIdentifier) {
    return nil, false
	}
	return o.AccountIdentifier, true
}

// HasAccountIdentifier returns a boolean if a field has been set.
func (o *NotificationContactCreateRequest) HasAccountIdentifier() bool {
	if o != nil && !isNil(o.AccountIdentifier) {
		return true
	}

	return false
}

// SetAccountIdentifier gets a reference to the given string and assigns it to the AccountIdentifier field.
func (o *NotificationContactCreateRequest) SetAccountIdentifier(v string) {
	o.AccountIdentifier = &v
}

func (o NotificationContactCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccountIdentifier) {
		toSerialize["account_identifier"] = o.AccountIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationContactCreateRequest struct {
	value *NotificationContactCreateRequest
	isSet bool
}

func (v NullableNotificationContactCreateRequest) Get() *NotificationContactCreateRequest {
	return v.value
}

func (v *NullableNotificationContactCreateRequest) Set(val *NotificationContactCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationContactCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationContactCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationContactCreateRequest(val *NotificationContactCreateRequest) *NullableNotificationContactCreateRequest {
	return &NullableNotificationContactCreateRequest{value: val, isSet: true}
}

func (v NullableNotificationContactCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationContactCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


