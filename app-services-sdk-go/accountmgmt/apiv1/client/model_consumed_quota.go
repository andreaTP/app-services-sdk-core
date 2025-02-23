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

// ConsumedQuota struct for ConsumedQuota
type ConsumedQuota struct {
	Href *string `json:"href,omitempty"`
	Id *string `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	AvailabilityZoneType *string `json:"availability_zone_type,omitempty"`
	BillingModel *string `json:"billing_model,omitempty"`
	Byoc bool `json:"byoc"`
	CloudProviderId *string `json:"cloud_provider_id,omitempty"`
	Count int32 `json:"count"`
	OrganizationId *string `json:"organization_id,omitempty"`
	PlanId *string `json:"plan_id,omitempty"`
	ResourceName *string `json:"resource_name,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewConsumedQuota instantiates a new ConsumedQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumedQuota(byoc bool, count int32) *ConsumedQuota {
	this := ConsumedQuota{}
	this.Byoc = byoc
	this.Count = count
	return &this
}

// NewConsumedQuotaWithDefaults instantiates a new ConsumedQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumedQuotaWithDefaults() *ConsumedQuota {
	this := ConsumedQuota{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ConsumedQuota) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumedQuota) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
    return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ConsumedQuota) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ConsumedQuota) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConsumedQuota) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumedQuota) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConsumedQuota) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConsumedQuota) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ConsumedQuota) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumedQuota) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ConsumedQuota) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ConsumedQuota) SetKind(v string) {
	o.Kind = &v
}

// GetAvailabilityZoneType returns the AvailabilityZoneType field value if set, zero value otherwise.
func (o *ConsumedQuota) GetAvailabilityZoneType() string {
	if o == nil || isNil(o.AvailabilityZoneType) {
		var ret string
		return ret
	}
	return *o.AvailabilityZoneType
}

// GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumedQuota) GetAvailabilityZoneTypeOk() (*string, bool) {
	if o == nil || isNil(o.AvailabilityZoneType) {
    return nil, false
	}
	return o.AvailabilityZoneType, true
}

// HasAvailabilityZoneType returns a boolean if a field has been set.
func (o *ConsumedQuota) HasAvailabilityZoneType() bool {
	if o != nil && !isNil(o.AvailabilityZoneType) {
		return true
	}

	return false
}

// SetAvailabilityZoneType gets a reference to the given string and assigns it to the AvailabilityZoneType field.
func (o *ConsumedQuota) SetAvailabilityZoneType(v string) {
	o.AvailabilityZoneType = &v
}

// GetBillingModel returns the BillingModel field value if set, zero value otherwise.
func (o *ConsumedQuota) GetBillingModel() string {
	if o == nil || isNil(o.BillingModel) {
		var ret string
		return ret
	}
	return *o.BillingModel
}

// GetBillingModelOk returns a tuple with the BillingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumedQuota) GetBillingModelOk() (*string, bool) {
	if o == nil || isNil(o.BillingModel) {
    return nil, false
	}
	return o.BillingModel, true
}

// HasBillingModel returns a boolean if a field has been set.
func (o *ConsumedQuota) HasBillingModel() bool {
	if o != nil && !isNil(o.BillingModel) {
		return true
	}

	return false
}

// SetBillingModel gets a reference to the given string and assigns it to the BillingModel field.
func (o *ConsumedQuota) SetBillingModel(v string) {
	o.BillingModel = &v
}

// GetByoc returns the Byoc field value
func (o *ConsumedQuota) GetByoc() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Byoc
}

// GetByocOk returns a tuple with the Byoc field value
// and a boolean to check if the value has been set.
func (o *ConsumedQuota) GetByocOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Byoc, true
}

// SetByoc sets field value
func (o *ConsumedQuota) SetByoc(v bool) {
	o.Byoc = v
}

// GetCloudProviderId returns the CloudProviderId field value if set, zero value otherwise.
func (o *ConsumedQuota) GetCloudProviderId() string {
	if o == nil || isNil(o.CloudProviderId) {
		var ret string
		return ret
	}
	return *o.CloudProviderId
}

// GetCloudProviderIdOk returns a tuple with the CloudProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumedQuota) GetCloudProviderIdOk() (*string, bool) {
	if o == nil || isNil(o.CloudProviderId) {
    return nil, false
	}
	return o.CloudProviderId, true
}

// HasCloudProviderId returns a boolean if a field has been set.
func (o *ConsumedQuota) HasCloudProviderId() bool {
	if o != nil && !isNil(o.CloudProviderId) {
		return true
	}

	return false
}

// SetCloudProviderId gets a reference to the given string and assigns it to the CloudProviderId field.
func (o *ConsumedQuota) SetCloudProviderId(v string) {
	o.CloudProviderId = &v
}

// GetCount returns the Count field value
func (o *ConsumedQuota) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ConsumedQuota) GetCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ConsumedQuota) SetCount(v int32) {
	o.Count = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ConsumedQuota) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumedQuota) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ConsumedQuota) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ConsumedQuota) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *ConsumedQuota) GetPlanId() string {
	if o == nil || isNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumedQuota) GetPlanIdOk() (*string, bool) {
	if o == nil || isNil(o.PlanId) {
    return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *ConsumedQuota) HasPlanId() bool {
	if o != nil && !isNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *ConsumedQuota) SetPlanId(v string) {
	o.PlanId = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *ConsumedQuota) GetResourceName() string {
	if o == nil || isNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumedQuota) GetResourceNameOk() (*string, bool) {
	if o == nil || isNil(o.ResourceName) {
    return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *ConsumedQuota) HasResourceName() bool {
	if o != nil && !isNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *ConsumedQuota) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ConsumedQuota) GetResourceType() string {
	if o == nil || isNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumedQuota) GetResourceTypeOk() (*string, bool) {
	if o == nil || isNil(o.ResourceType) {
    return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ConsumedQuota) HasResourceType() bool {
	if o != nil && !isNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ConsumedQuota) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ConsumedQuota) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumedQuota) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ConsumedQuota) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ConsumedQuota) SetVersion(v string) {
	o.Version = &v
}

func (o ConsumedQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !isNil(o.AvailabilityZoneType) {
		toSerialize["availability_zone_type"] = o.AvailabilityZoneType
	}
	if !isNil(o.BillingModel) {
		toSerialize["billing_model"] = o.BillingModel
	}
	if true {
		toSerialize["byoc"] = o.Byoc
	}
	if !isNil(o.CloudProviderId) {
		toSerialize["cloud_provider_id"] = o.CloudProviderId
	}
	if true {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !isNil(o.ResourceName) {
		toSerialize["resource_name"] = o.ResourceName
	}
	if !isNil(o.ResourceType) {
		toSerialize["resource_type"] = o.ResourceType
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableConsumedQuota struct {
	value *ConsumedQuota
	isSet bool
}

func (v NullableConsumedQuota) Get() *ConsumedQuota {
	return v.value
}

func (v *NullableConsumedQuota) Set(val *ConsumedQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumedQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumedQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumedQuota(val *ConsumedQuota) *NullableConsumedQuota {
	return &NullableConsumedQuota{value: val, isSet: true}
}

func (v NullableConsumedQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumedQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


