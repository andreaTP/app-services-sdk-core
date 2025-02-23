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

// LabelAllOf struct for LabelAllOf
type LabelAllOf struct {
	AccountId *string `json:"account_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Internal bool `json:"internal"`
	Key string `json:"key"`
	OrganizationId *string `json:"organization_id,omitempty"`
	SubscriptionId *string `json:"subscription_id,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Value string `json:"value"`
}

// NewLabelAllOf instantiates a new LabelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelAllOf(internal bool, key string, value string) *LabelAllOf {
	this := LabelAllOf{}
	this.Internal = internal
	this.Key = key
	this.Value = value
	return &this
}

// NewLabelAllOfWithDefaults instantiates a new LabelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelAllOfWithDefaults() *LabelAllOf {
	this := LabelAllOf{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *LabelAllOf) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelAllOf) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
    return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *LabelAllOf) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *LabelAllOf) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LabelAllOf) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LabelAllOf) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LabelAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetInternal returns the Internal field value
func (o *LabelAllOf) GetInternal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Internal
}

// GetInternalOk returns a tuple with the Internal field value
// and a boolean to check if the value has been set.
func (o *LabelAllOf) GetInternalOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Internal, true
}

// SetInternal sets field value
func (o *LabelAllOf) SetInternal(v bool) {
	o.Internal = v
}

// GetKey returns the Key field value
func (o *LabelAllOf) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *LabelAllOf) GetKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *LabelAllOf) SetKey(v string) {
	o.Key = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *LabelAllOf) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelAllOf) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *LabelAllOf) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *LabelAllOf) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *LabelAllOf) GetSubscriptionId() string {
	if o == nil || isNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelAllOf) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || isNil(o.SubscriptionId) {
    return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *LabelAllOf) HasSubscriptionId() bool {
	if o != nil && !isNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *LabelAllOf) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LabelAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LabelAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LabelAllOf) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LabelAllOf) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelAllOf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LabelAllOf) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *LabelAllOf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetValue returns the Value field value
func (o *LabelAllOf) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *LabelAllOf) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *LabelAllOf) SetValue(v string) {
	o.Value = v
}

func (o LabelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["internal"] = o.Internal
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.SubscriptionId) {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableLabelAllOf struct {
	value *LabelAllOf
	isSet bool
}

func (v NullableLabelAllOf) Get() *LabelAllOf {
	return v.value
}

func (v *NullableLabelAllOf) Set(val *LabelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelAllOf(val *LabelAllOf) *NullableLabelAllOf {
	return &NullableLabelAllOf{value: val, isSet: true}
}

func (v NullableLabelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


