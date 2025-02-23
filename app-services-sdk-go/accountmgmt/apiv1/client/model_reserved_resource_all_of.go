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

// ReservedResourceAllOf struct for ReservedResourceAllOf
type ReservedResourceAllOf struct {
	AvailabilityZoneType *string `json:"availability_zone_type,omitempty"`
	BillingMarketplaceAccount *string `json:"billing_marketplace_account,omitempty"`
	BillingModel *string `json:"billing_model,omitempty"`
	Byoc bool `json:"byoc"`
	Cluster *bool `json:"cluster,omitempty"`
	Count *int32 `json:"count,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ResourceName *string `json:"resource_name,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
	Subscription *ObjectReference `json:"subscription,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewReservedResourceAllOf instantiates a new ReservedResourceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservedResourceAllOf(byoc bool) *ReservedResourceAllOf {
	this := ReservedResourceAllOf{}
	this.Byoc = byoc
	return &this
}

// NewReservedResourceAllOfWithDefaults instantiates a new ReservedResourceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservedResourceAllOfWithDefaults() *ReservedResourceAllOf {
	this := ReservedResourceAllOf{}
	return &this
}

// GetAvailabilityZoneType returns the AvailabilityZoneType field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetAvailabilityZoneType() string {
	if o == nil || isNil(o.AvailabilityZoneType) {
		var ret string
		return ret
	}
	return *o.AvailabilityZoneType
}

// GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetAvailabilityZoneTypeOk() (*string, bool) {
	if o == nil || isNil(o.AvailabilityZoneType) {
    return nil, false
	}
	return o.AvailabilityZoneType, true
}

// HasAvailabilityZoneType returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasAvailabilityZoneType() bool {
	if o != nil && !isNil(o.AvailabilityZoneType) {
		return true
	}

	return false
}

// SetAvailabilityZoneType gets a reference to the given string and assigns it to the AvailabilityZoneType field.
func (o *ReservedResourceAllOf) SetAvailabilityZoneType(v string) {
	o.AvailabilityZoneType = &v
}

// GetBillingMarketplaceAccount returns the BillingMarketplaceAccount field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetBillingMarketplaceAccount() string {
	if o == nil || isNil(o.BillingMarketplaceAccount) {
		var ret string
		return ret
	}
	return *o.BillingMarketplaceAccount
}

// GetBillingMarketplaceAccountOk returns a tuple with the BillingMarketplaceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetBillingMarketplaceAccountOk() (*string, bool) {
	if o == nil || isNil(o.BillingMarketplaceAccount) {
    return nil, false
	}
	return o.BillingMarketplaceAccount, true
}

// HasBillingMarketplaceAccount returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasBillingMarketplaceAccount() bool {
	if o != nil && !isNil(o.BillingMarketplaceAccount) {
		return true
	}

	return false
}

// SetBillingMarketplaceAccount gets a reference to the given string and assigns it to the BillingMarketplaceAccount field.
func (o *ReservedResourceAllOf) SetBillingMarketplaceAccount(v string) {
	o.BillingMarketplaceAccount = &v
}

// GetBillingModel returns the BillingModel field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetBillingModel() string {
	if o == nil || isNil(o.BillingModel) {
		var ret string
		return ret
	}
	return *o.BillingModel
}

// GetBillingModelOk returns a tuple with the BillingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetBillingModelOk() (*string, bool) {
	if o == nil || isNil(o.BillingModel) {
    return nil, false
	}
	return o.BillingModel, true
}

// HasBillingModel returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasBillingModel() bool {
	if o != nil && !isNil(o.BillingModel) {
		return true
	}

	return false
}

// SetBillingModel gets a reference to the given string and assigns it to the BillingModel field.
func (o *ReservedResourceAllOf) SetBillingModel(v string) {
	o.BillingModel = &v
}

// GetByoc returns the Byoc field value
func (o *ReservedResourceAllOf) GetByoc() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Byoc
}

// GetByocOk returns a tuple with the Byoc field value
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetByocOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Byoc, true
}

// SetByoc sets field value
func (o *ReservedResourceAllOf) SetByoc(v bool) {
	o.Byoc = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetCluster() bool {
	if o == nil || isNil(o.Cluster) {
		var ret bool
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetClusterOk() (*bool, bool) {
	if o == nil || isNil(o.Cluster) {
    return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasCluster() bool {
	if o != nil && !isNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given bool and assigns it to the Cluster field.
func (o *ReservedResourceAllOf) SetCluster(v bool) {
	o.Cluster = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ReservedResourceAllOf) SetCount(v int32) {
	o.Count = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ReservedResourceAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetResourceName() string {
	if o == nil || isNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetResourceNameOk() (*string, bool) {
	if o == nil || isNil(o.ResourceName) {
    return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasResourceName() bool {
	if o != nil && !isNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *ReservedResourceAllOf) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetResourceType() string {
	if o == nil || isNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetResourceTypeOk() (*string, bool) {
	if o == nil || isNil(o.ResourceType) {
    return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasResourceType() bool {
	if o != nil && !isNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ReservedResourceAllOf) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetSubscription() ObjectReference {
	if o == nil || isNil(o.Subscription) {
		var ret ObjectReference
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetSubscriptionOk() (*ObjectReference, bool) {
	if o == nil || isNil(o.Subscription) {
    return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasSubscription() bool {
	if o != nil && !isNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given ObjectReference and assigns it to the Subscription field.
func (o *ReservedResourceAllOf) SetSubscription(v ObjectReference) {
	o.Subscription = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ReservedResourceAllOf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ReservedResourceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AvailabilityZoneType) {
		toSerialize["availability_zone_type"] = o.AvailabilityZoneType
	}
	if !isNil(o.BillingMarketplaceAccount) {
		toSerialize["billing_marketplace_account"] = o.BillingMarketplaceAccount
	}
	if !isNil(o.BillingModel) {
		toSerialize["billing_model"] = o.BillingModel
	}
	if true {
		toSerialize["byoc"] = o.Byoc
	}
	if !isNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.ResourceName) {
		toSerialize["resource_name"] = o.ResourceName
	}
	if !isNil(o.ResourceType) {
		toSerialize["resource_type"] = o.ResourceType
	}
	if !isNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableReservedResourceAllOf struct {
	value *ReservedResourceAllOf
	isSet bool
}

func (v NullableReservedResourceAllOf) Get() *ReservedResourceAllOf {
	return v.value
}

func (v *NullableReservedResourceAllOf) Set(val *ReservedResourceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReservedResourceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReservedResourceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservedResourceAllOf(val *ReservedResourceAllOf) *NullableReservedResourceAllOf {
	return &NullableReservedResourceAllOf{value: val, isSet: true}
}

func (v NullableReservedResourceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservedResourceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


