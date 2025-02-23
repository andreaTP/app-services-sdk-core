/*
Account Management Service API

Manage user subscriptions and clusters

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"encoding/json"
	"time"
)

// RegistryCredentialAllOf struct for RegistryCredentialAllOf
type RegistryCredentialAllOf struct {
	Account *ObjectReference `json:"account,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ExternalResourceId *string `json:"external_resource_id,omitempty"`
	Registry *ObjectReference `json:"registry,omitempty"`
	Token *string `json:"token,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewRegistryCredentialAllOf instantiates a new RegistryCredentialAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryCredentialAllOf() *RegistryCredentialAllOf {
	this := RegistryCredentialAllOf{}
	return &this
}

// NewRegistryCredentialAllOfWithDefaults instantiates a new RegistryCredentialAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryCredentialAllOfWithDefaults() *RegistryCredentialAllOf {
	this := RegistryCredentialAllOf{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *RegistryCredentialAllOf) GetAccount() ObjectReference {
	if o == nil || isNil(o.Account) {
		var ret ObjectReference
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCredentialAllOf) GetAccountOk() (*ObjectReference, bool) {
	if o == nil || isNil(o.Account) {
    return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *RegistryCredentialAllOf) HasAccount() bool {
	if o != nil && !isNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ObjectReference and assigns it to the Account field.
func (o *RegistryCredentialAllOf) SetAccount(v ObjectReference) {
	o.Account = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RegistryCredentialAllOf) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCredentialAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RegistryCredentialAllOf) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RegistryCredentialAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExternalResourceId returns the ExternalResourceId field value if set, zero value otherwise.
func (o *RegistryCredentialAllOf) GetExternalResourceId() string {
	if o == nil || isNil(o.ExternalResourceId) {
		var ret string
		return ret
	}
	return *o.ExternalResourceId
}

// GetExternalResourceIdOk returns a tuple with the ExternalResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCredentialAllOf) GetExternalResourceIdOk() (*string, bool) {
	if o == nil || isNil(o.ExternalResourceId) {
    return nil, false
	}
	return o.ExternalResourceId, true
}

// HasExternalResourceId returns a boolean if a field has been set.
func (o *RegistryCredentialAllOf) HasExternalResourceId() bool {
	if o != nil && !isNil(o.ExternalResourceId) {
		return true
	}

	return false
}

// SetExternalResourceId gets a reference to the given string and assigns it to the ExternalResourceId field.
func (o *RegistryCredentialAllOf) SetExternalResourceId(v string) {
	o.ExternalResourceId = &v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *RegistryCredentialAllOf) GetRegistry() ObjectReference {
	if o == nil || isNil(o.Registry) {
		var ret ObjectReference
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCredentialAllOf) GetRegistryOk() (*ObjectReference, bool) {
	if o == nil || isNil(o.Registry) {
    return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *RegistryCredentialAllOf) HasRegistry() bool {
	if o != nil && !isNil(o.Registry) {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given ObjectReference and assigns it to the Registry field.
func (o *RegistryCredentialAllOf) SetRegistry(v ObjectReference) {
	o.Registry = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RegistryCredentialAllOf) GetToken() string {
	if o == nil || isNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCredentialAllOf) GetTokenOk() (*string, bool) {
	if o == nil || isNil(o.Token) {
    return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RegistryCredentialAllOf) HasToken() bool {
	if o != nil && !isNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RegistryCredentialAllOf) SetToken(v string) {
	o.Token = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RegistryCredentialAllOf) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCredentialAllOf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RegistryCredentialAllOf) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RegistryCredentialAllOf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *RegistryCredentialAllOf) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCredentialAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *RegistryCredentialAllOf) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *RegistryCredentialAllOf) SetUsername(v string) {
	o.Username = &v
}

func (o RegistryCredentialAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.ExternalResourceId) {
		toSerialize["external_resource_id"] = o.ExternalResourceId
	}
	if !isNil(o.Registry) {
		toSerialize["registry"] = o.Registry
	}
	if !isNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableRegistryCredentialAllOf struct {
	value *RegistryCredentialAllOf
	isSet bool
}

func (v NullableRegistryCredentialAllOf) Get() *RegistryCredentialAllOf {
	return v.value
}

func (v *NullableRegistryCredentialAllOf) Set(val *RegistryCredentialAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryCredentialAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryCredentialAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryCredentialAllOf(val *RegistryCredentialAllOf) *NullableRegistryCredentialAllOf {
	return &NullableRegistryCredentialAllOf{value: val, isSet: true}
}

func (v NullableRegistryCredentialAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryCredentialAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


