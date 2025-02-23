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

// RegistryRequest struct for RegistryRequest
type RegistryRequest struct {
	CloudAlias *bool `json:"cloudAlias,omitempty"`
	Name *string `json:"name,omitempty"`
	OrgName *string `json:"org_name,omitempty"`
	TeamName *string `json:"team_name,omitempty"`
	Type *string `json:"type,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewRegistryRequest instantiates a new RegistryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryRequest() *RegistryRequest {
	this := RegistryRequest{}
	return &this
}

// NewRegistryRequestWithDefaults instantiates a new RegistryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryRequestWithDefaults() *RegistryRequest {
	this := RegistryRequest{}
	return &this
}

// GetCloudAlias returns the CloudAlias field value if set, zero value otherwise.
func (o *RegistryRequest) GetCloudAlias() bool {
	if o == nil || isNil(o.CloudAlias) {
		var ret bool
		return ret
	}
	return *o.CloudAlias
}

// GetCloudAliasOk returns a tuple with the CloudAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryRequest) GetCloudAliasOk() (*bool, bool) {
	if o == nil || isNil(o.CloudAlias) {
    return nil, false
	}
	return o.CloudAlias, true
}

// HasCloudAlias returns a boolean if a field has been set.
func (o *RegistryRequest) HasCloudAlias() bool {
	if o != nil && !isNil(o.CloudAlias) {
		return true
	}

	return false
}

// SetCloudAlias gets a reference to the given bool and assigns it to the CloudAlias field.
func (o *RegistryRequest) SetCloudAlias(v bool) {
	o.CloudAlias = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RegistryRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegistryRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegistryRequest) SetName(v string) {
	o.Name = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *RegistryRequest) GetOrgName() string {
	if o == nil || isNil(o.OrgName) {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryRequest) GetOrgNameOk() (*string, bool) {
	if o == nil || isNil(o.OrgName) {
    return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *RegistryRequest) HasOrgName() bool {
	if o != nil && !isNil(o.OrgName) {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *RegistryRequest) SetOrgName(v string) {
	o.OrgName = &v
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *RegistryRequest) GetTeamName() string {
	if o == nil || isNil(o.TeamName) {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryRequest) GetTeamNameOk() (*string, bool) {
	if o == nil || isNil(o.TeamName) {
    return nil, false
	}
	return o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *RegistryRequest) HasTeamName() bool {
	if o != nil && !isNil(o.TeamName) {
		return true
	}

	return false
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *RegistryRequest) SetTeamName(v string) {
	o.TeamName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RegistryRequest) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryRequest) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RegistryRequest) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RegistryRequest) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RegistryRequest) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryRequest) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RegistryRequest) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RegistryRequest) SetUrl(v string) {
	o.Url = &v
}

func (o RegistryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CloudAlias) {
		toSerialize["cloudAlias"] = o.CloudAlias
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.OrgName) {
		toSerialize["org_name"] = o.OrgName
	}
	if !isNil(o.TeamName) {
		toSerialize["team_name"] = o.TeamName
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableRegistryRequest struct {
	value *RegistryRequest
	isSet bool
}

func (v NullableRegistryRequest) Get() *RegistryRequest {
	return v.value
}

func (v *NullableRegistryRequest) Set(val *RegistryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryRequest(val *RegistryRequest) *NullableRegistryRequest {
	return &NullableRegistryRequest{value: val, isSet: true}
}

func (v NullableRegistryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


