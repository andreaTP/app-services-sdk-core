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

// RegistryCredentialPatchRequest struct for RegistryCredentialPatchRequest
type RegistryCredentialPatchRequest struct {
	AccountId *string `json:"account_id,omitempty"`
	ExternalResourceId *string `json:"external_resource_id,omitempty"`
	RegistryId *string `json:"registry_id,omitempty"`
	Token *string `json:"token,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewRegistryCredentialPatchRequest instantiates a new RegistryCredentialPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryCredentialPatchRequest() *RegistryCredentialPatchRequest {
	this := RegistryCredentialPatchRequest{}
	return &this
}

// NewRegistryCredentialPatchRequestWithDefaults instantiates a new RegistryCredentialPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryCredentialPatchRequestWithDefaults() *RegistryCredentialPatchRequest {
	this := RegistryCredentialPatchRequest{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *RegistryCredentialPatchRequest) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCredentialPatchRequest) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
    return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *RegistryCredentialPatchRequest) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *RegistryCredentialPatchRequest) SetAccountId(v string) {
	o.AccountId = &v
}

// GetExternalResourceId returns the ExternalResourceId field value if set, zero value otherwise.
func (o *RegistryCredentialPatchRequest) GetExternalResourceId() string {
	if o == nil || isNil(o.ExternalResourceId) {
		var ret string
		return ret
	}
	return *o.ExternalResourceId
}

// GetExternalResourceIdOk returns a tuple with the ExternalResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCredentialPatchRequest) GetExternalResourceIdOk() (*string, bool) {
	if o == nil || isNil(o.ExternalResourceId) {
    return nil, false
	}
	return o.ExternalResourceId, true
}

// HasExternalResourceId returns a boolean if a field has been set.
func (o *RegistryCredentialPatchRequest) HasExternalResourceId() bool {
	if o != nil && !isNil(o.ExternalResourceId) {
		return true
	}

	return false
}

// SetExternalResourceId gets a reference to the given string and assigns it to the ExternalResourceId field.
func (o *RegistryCredentialPatchRequest) SetExternalResourceId(v string) {
	o.ExternalResourceId = &v
}

// GetRegistryId returns the RegistryId field value if set, zero value otherwise.
func (o *RegistryCredentialPatchRequest) GetRegistryId() string {
	if o == nil || isNil(o.RegistryId) {
		var ret string
		return ret
	}
	return *o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCredentialPatchRequest) GetRegistryIdOk() (*string, bool) {
	if o == nil || isNil(o.RegistryId) {
    return nil, false
	}
	return o.RegistryId, true
}

// HasRegistryId returns a boolean if a field has been set.
func (o *RegistryCredentialPatchRequest) HasRegistryId() bool {
	if o != nil && !isNil(o.RegistryId) {
		return true
	}

	return false
}

// SetRegistryId gets a reference to the given string and assigns it to the RegistryId field.
func (o *RegistryCredentialPatchRequest) SetRegistryId(v string) {
	o.RegistryId = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RegistryCredentialPatchRequest) GetToken() string {
	if o == nil || isNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCredentialPatchRequest) GetTokenOk() (*string, bool) {
	if o == nil || isNil(o.Token) {
    return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RegistryCredentialPatchRequest) HasToken() bool {
	if o != nil && !isNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RegistryCredentialPatchRequest) SetToken(v string) {
	o.Token = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *RegistryCredentialPatchRequest) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCredentialPatchRequest) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *RegistryCredentialPatchRequest) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *RegistryCredentialPatchRequest) SetUsername(v string) {
	o.Username = &v
}

func (o RegistryCredentialPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !isNil(o.ExternalResourceId) {
		toSerialize["external_resource_id"] = o.ExternalResourceId
	}
	if !isNil(o.RegistryId) {
		toSerialize["registry_id"] = o.RegistryId
	}
	if !isNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableRegistryCredentialPatchRequest struct {
	value *RegistryCredentialPatchRequest
	isSet bool
}

func (v NullableRegistryCredentialPatchRequest) Get() *RegistryCredentialPatchRequest {
	return v.value
}

func (v *NullableRegistryCredentialPatchRequest) Set(val *RegistryCredentialPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryCredentialPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryCredentialPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryCredentialPatchRequest(val *RegistryCredentialPatchRequest) *NullableRegistryCredentialPatchRequest {
	return &NullableRegistryCredentialPatchRequest{value: val, isSet: true}
}

func (v NullableRegistryCredentialPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryCredentialPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


