/*
Kafka Management API

Kafka Management API is a REST API to manage Kafka instances

API version: 1.14.0
Contact: rhosak-support@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// ServiceAccountListItem 
type ServiceAccountListItem struct {
	// server generated unique id of the service account
	Id string `json:"id"`
	Kind string `json:"kind"`
	Href string `json:"href"`
	// client id of the service account
	ClientId *string `json:"client_id,omitempty"`
	// name of the service account
	Name *string `json:"name,omitempty"`
	// owner of the service account
	// Deprecated
	Owner *string `json:"owner,omitempty"`
	// service account created by the user
	CreatedBy *string `json:"created_by,omitempty"`
	// service account creation timestamp
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// description of the service account
	Description *string `json:"description,omitempty"`
}

// NewServiceAccountListItem instantiates a new ServiceAccountListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountListItem(id string, kind string, href string) *ServiceAccountListItem {
	this := ServiceAccountListItem{}
	this.Id = id
	this.Kind = kind
	this.Href = href
	return &this
}

// NewServiceAccountListItemWithDefaults instantiates a new ServiceAccountListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountListItemWithDefaults() *ServiceAccountListItem {
	this := ServiceAccountListItem{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceAccountListItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceAccountListItem) SetId(v string) {
	o.Id = v
}

// GetKind returns the Kind field value
func (o *ServiceAccountListItem) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetKindOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ServiceAccountListItem) SetKind(v string) {
	o.Kind = v
}

// GetHref returns the Href field value
func (o *ServiceAccountListItem) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetHrefOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ServiceAccountListItem) SetHref(v string) {
	o.Href = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ServiceAccountListItem) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
    return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ServiceAccountListItem) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ServiceAccountListItem) SetClientId(v string) {
	o.ClientId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceAccountListItem) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAccountListItem) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAccountListItem) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
// Deprecated
func (o *ServiceAccountListItem) GetOwner() string {
	if o == nil || isNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ServiceAccountListItem) GetOwnerOk() (*string, bool) {
	if o == nil || isNil(o.Owner) {
    return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ServiceAccountListItem) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
// Deprecated
func (o *ServiceAccountListItem) SetOwner(v string) {
	o.Owner = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ServiceAccountListItem) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetCreatedByOk() (*string, bool) {
	if o == nil || isNil(o.CreatedBy) {
    return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ServiceAccountListItem) HasCreatedBy() bool {
	if o != nil && !isNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ServiceAccountListItem) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceAccountListItem) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceAccountListItem) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServiceAccountListItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceAccountListItem) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItem) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccountListItem) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccountListItem) SetDescription(v string) {
	o.Description = &v
}

func (o ServiceAccountListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAccountListItem struct {
	value *ServiceAccountListItem
	isSet bool
}

func (v NullableServiceAccountListItem) Get() *ServiceAccountListItem {
	return v.value
}

func (v *NullableServiceAccountListItem) Set(val *ServiceAccountListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountListItem(val *ServiceAccountListItem) *NullableServiceAccountListItem {
	return &NullableServiceAccountListItem{value: val, isSet: true}
}

func (v NullableServiceAccountListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


