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

// TokenAuthorizationResponse struct for TokenAuthorizationResponse
type TokenAuthorizationResponse struct {
	Account *Account `json:"account,omitempty"`
}

// NewTokenAuthorizationResponse instantiates a new TokenAuthorizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenAuthorizationResponse() *TokenAuthorizationResponse {
	this := TokenAuthorizationResponse{}
	return &this
}

// NewTokenAuthorizationResponseWithDefaults instantiates a new TokenAuthorizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenAuthorizationResponseWithDefaults() *TokenAuthorizationResponse {
	this := TokenAuthorizationResponse{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TokenAuthorizationResponse) GetAccount() Account {
	if o == nil || isNil(o.Account) {
		var ret Account
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenAuthorizationResponse) GetAccountOk() (*Account, bool) {
	if o == nil || isNil(o.Account) {
    return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TokenAuthorizationResponse) HasAccount() bool {
	if o != nil && !isNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Account and assigns it to the Account field.
func (o *TokenAuthorizationResponse) SetAccount(v Account) {
	o.Account = &v
}

func (o TokenAuthorizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableTokenAuthorizationResponse struct {
	value *TokenAuthorizationResponse
	isSet bool
}

func (v NullableTokenAuthorizationResponse) Get() *TokenAuthorizationResponse {
	return v.value
}

func (v *NullableTokenAuthorizationResponse) Set(val *TokenAuthorizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenAuthorizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenAuthorizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenAuthorizationResponse(val *TokenAuthorizationResponse) *NullableTokenAuthorizationResponse {
	return &NullableTokenAuthorizationResponse{value: val, isSet: true}
}

func (v NullableTokenAuthorizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenAuthorizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


