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

// PullSecretRequest struct for PullSecretRequest
type PullSecretRequest struct {
	ExternalResourceId string `json:"external_resource_id"`
}

// NewPullSecretRequest instantiates a new PullSecretRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullSecretRequest(externalResourceId string) *PullSecretRequest {
	this := PullSecretRequest{}
	this.ExternalResourceId = externalResourceId
	return &this
}

// NewPullSecretRequestWithDefaults instantiates a new PullSecretRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullSecretRequestWithDefaults() *PullSecretRequest {
	this := PullSecretRequest{}
	return &this
}

// GetExternalResourceId returns the ExternalResourceId field value
func (o *PullSecretRequest) GetExternalResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalResourceId
}

// GetExternalResourceIdOk returns a tuple with the ExternalResourceId field value
// and a boolean to check if the value has been set.
func (o *PullSecretRequest) GetExternalResourceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExternalResourceId, true
}

// SetExternalResourceId sets field value
func (o *PullSecretRequest) SetExternalResourceId(v string) {
	o.ExternalResourceId = v
}

func (o PullSecretRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["external_resource_id"] = o.ExternalResourceId
	}
	return json.Marshal(toSerialize)
}

type NullablePullSecretRequest struct {
	value *PullSecretRequest
	isSet bool
}

func (v NullablePullSecretRequest) Get() *PullSecretRequest {
	return v.value
}

func (v *NullablePullSecretRequest) Set(val *PullSecretRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePullSecretRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePullSecretRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullSecretRequest(val *PullSecretRequest) *NullablePullSecretRequest {
	return &NullablePullSecretRequest{value: val, isSet: true}
}

func (v NullablePullSecretRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullSecretRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


