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

// AcsClientResponseData struct for AcsClientResponseData
type AcsClientResponseData struct {
	ClientId *string `json:"clientId,omitempty"`
	Secret *string `json:"secret,omitempty"`
	Name *string `json:"name,omitempty"`
	CreatedAt *int64 `json:"createdAt,omitempty"`
}

// NewAcsClientResponseData instantiates a new AcsClientResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcsClientResponseData() *AcsClientResponseData {
	this := AcsClientResponseData{}
	return &this
}

// NewAcsClientResponseDataWithDefaults instantiates a new AcsClientResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcsClientResponseDataWithDefaults() *AcsClientResponseData {
	this := AcsClientResponseData{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AcsClientResponseData) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcsClientResponseData) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
    return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AcsClientResponseData) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AcsClientResponseData) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AcsClientResponseData) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcsClientResponseData) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
    return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AcsClientResponseData) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AcsClientResponseData) SetSecret(v string) {
	o.Secret = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AcsClientResponseData) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcsClientResponseData) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AcsClientResponseData) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AcsClientResponseData) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AcsClientResponseData) GetCreatedAt() int64 {
	if o == nil || isNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcsClientResponseData) GetCreatedAtOk() (*int64, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AcsClientResponseData) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *AcsClientResponseData) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

func (o AcsClientResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableAcsClientResponseData struct {
	value *AcsClientResponseData
	isSet bool
}

func (v NullableAcsClientResponseData) Get() *AcsClientResponseData {
	return v.value
}

func (v *NullableAcsClientResponseData) Set(val *AcsClientResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAcsClientResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAcsClientResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcsClientResponseData(val *AcsClientResponseData) *NullableAcsClientResponseData {
	return &NullableAcsClientResponseData{value: val, isSet: true}
}

func (v NullableAcsClientResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcsClientResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


