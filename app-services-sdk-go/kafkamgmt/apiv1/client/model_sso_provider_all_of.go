/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.16.0
 * Contact: rhosak-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// SsoProviderAllOf struct for SsoProviderAllOf
type SsoProviderAllOf struct {
	// name of the sso provider
	Name *string `json:"name,omitempty"`
	// base url
	BaseUrl *string `json:"base_url,omitempty"`
	TokenUrl *string `json:"token_url,omitempty"`
	Jwks *string `json:"jwks,omitempty"`
	ValidIssuer *string `json:"valid_issuer,omitempty"`
}

// NewSsoProviderAllOf instantiates a new SsoProviderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoProviderAllOf() *SsoProviderAllOf {
	this := SsoProviderAllOf{}
	return &this
}

// NewSsoProviderAllOfWithDefaults instantiates a new SsoProviderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoProviderAllOfWithDefaults() *SsoProviderAllOf {
	this := SsoProviderAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SsoProviderAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoProviderAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SsoProviderAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SsoProviderAllOf) SetName(v string) {
	o.Name = &v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *SsoProviderAllOf) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoProviderAllOf) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *SsoProviderAllOf) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *SsoProviderAllOf) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *SsoProviderAllOf) GetTokenUrl() string {
	if o == nil || o.TokenUrl == nil {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoProviderAllOf) GetTokenUrlOk() (*string, bool) {
	if o == nil || o.TokenUrl == nil {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *SsoProviderAllOf) HasTokenUrl() bool {
	if o != nil && o.TokenUrl != nil {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *SsoProviderAllOf) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

// GetJwks returns the Jwks field value if set, zero value otherwise.
func (o *SsoProviderAllOf) GetJwks() string {
	if o == nil || o.Jwks == nil {
		var ret string
		return ret
	}
	return *o.Jwks
}

// GetJwksOk returns a tuple with the Jwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoProviderAllOf) GetJwksOk() (*string, bool) {
	if o == nil || o.Jwks == nil {
		return nil, false
	}
	return o.Jwks, true
}

// HasJwks returns a boolean if a field has been set.
func (o *SsoProviderAllOf) HasJwks() bool {
	if o != nil && o.Jwks != nil {
		return true
	}

	return false
}

// SetJwks gets a reference to the given string and assigns it to the Jwks field.
func (o *SsoProviderAllOf) SetJwks(v string) {
	o.Jwks = &v
}

// GetValidIssuer returns the ValidIssuer field value if set, zero value otherwise.
func (o *SsoProviderAllOf) GetValidIssuer() string {
	if o == nil || o.ValidIssuer == nil {
		var ret string
		return ret
	}
	return *o.ValidIssuer
}

// GetValidIssuerOk returns a tuple with the ValidIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoProviderAllOf) GetValidIssuerOk() (*string, bool) {
	if o == nil || o.ValidIssuer == nil {
		return nil, false
	}
	return o.ValidIssuer, true
}

// HasValidIssuer returns a boolean if a field has been set.
func (o *SsoProviderAllOf) HasValidIssuer() bool {
	if o != nil && o.ValidIssuer != nil {
		return true
	}

	return false
}

// SetValidIssuer gets a reference to the given string and assigns it to the ValidIssuer field.
func (o *SsoProviderAllOf) SetValidIssuer(v string) {
	o.ValidIssuer = &v
}

func (o SsoProviderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.BaseUrl != nil {
		toSerialize["base_url"] = o.BaseUrl
	}
	if o.TokenUrl != nil {
		toSerialize["token_url"] = o.TokenUrl
	}
	if o.Jwks != nil {
		toSerialize["jwks"] = o.Jwks
	}
	if o.ValidIssuer != nil {
		toSerialize["valid_issuer"] = o.ValidIssuer
	}
	return json.Marshal(toSerialize)
}

type NullableSsoProviderAllOf struct {
	value *SsoProviderAllOf
	isSet bool
}

func (v NullableSsoProviderAllOf) Get() *SsoProviderAllOf {
	return v.value
}

func (v *NullableSsoProviderAllOf) Set(val *SsoProviderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoProviderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoProviderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoProviderAllOf(val *SsoProviderAllOf) *NullableSsoProviderAllOf {
	return &NullableSsoProviderAllOf{value: val, isSet: true}
}

func (v NullableSsoProviderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoProviderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


