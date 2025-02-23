/*
sso.redhat.com API documentation

This is the API documentation for sso.redhat.com

API version: 5.0.19-SNAPSHOT
Contact: it-user-team-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceaccountsclient

import (
	"encoding/json"
)

// AuthenticationPolicy struct for AuthenticationPolicy
type AuthenticationPolicy struct {
	AuthenticationFactors *AuthenticationFactors `json:"authenticationFactors,omitempty"`
}

// NewAuthenticationPolicy instantiates a new AuthenticationPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationPolicy() *AuthenticationPolicy {
	this := AuthenticationPolicy{}
	return &this
}

// NewAuthenticationPolicyWithDefaults instantiates a new AuthenticationPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationPolicyWithDefaults() *AuthenticationPolicy {
	this := AuthenticationPolicy{}
	return &this
}

// GetAuthenticationFactors returns the AuthenticationFactors field value if set, zero value otherwise.
func (o *AuthenticationPolicy) GetAuthenticationFactors() AuthenticationFactors {
	if o == nil || isNil(o.AuthenticationFactors) {
		var ret AuthenticationFactors
		return ret
	}
	return *o.AuthenticationFactors
}

// GetAuthenticationFactorsOk returns a tuple with the AuthenticationFactors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicy) GetAuthenticationFactorsOk() (*AuthenticationFactors, bool) {
	if o == nil || isNil(o.AuthenticationFactors) {
    return nil, false
	}
	return o.AuthenticationFactors, true
}

// HasAuthenticationFactors returns a boolean if a field has been set.
func (o *AuthenticationPolicy) HasAuthenticationFactors() bool {
	if o != nil && !isNil(o.AuthenticationFactors) {
		return true
	}

	return false
}

// SetAuthenticationFactors gets a reference to the given AuthenticationFactors and assigns it to the AuthenticationFactors field.
func (o *AuthenticationPolicy) SetAuthenticationFactors(v AuthenticationFactors) {
	o.AuthenticationFactors = &v
}

func (o AuthenticationPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AuthenticationFactors) {
		toSerialize["authenticationFactors"] = o.AuthenticationFactors
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticationPolicy struct {
	value *AuthenticationPolicy
	isSet bool
}

func (v NullableAuthenticationPolicy) Get() *AuthenticationPolicy {
	return v.value
}

func (v *NullableAuthenticationPolicy) Set(val *AuthenticationPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationPolicy(val *AuthenticationPolicy) *NullableAuthenticationPolicy {
	return &NullableAuthenticationPolicy{value: val, isSet: true}
}

func (v NullableAuthenticationPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


