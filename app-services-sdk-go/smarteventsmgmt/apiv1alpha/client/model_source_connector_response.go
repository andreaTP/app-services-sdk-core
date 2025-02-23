/*
Red Hat Openshift SmartEvents Fleet Manager V2

The API exposed by the fleet manager of the SmartEvents service.

API version: 0.0.1
Contact: openbridge-dev@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smarteventsmgmtclient

import (
	"encoding/json"
	"time"
)

// SourceConnectorResponse struct for SourceConnectorResponse
type SourceConnectorResponse struct {
	// The kind (type) of this resource
	Kind string `json:"kind"`
	// The unique identifier of this resource
	Id string `json:"id"`
	// The name of this resource
	Name string `json:"name"`
	// The URL of this resource, without the protocol
	Href string `json:"href"`
	SubmittedAt time.Time `json:"submitted_at"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	Status ManagedResourceStatus `json:"status"`
	// The user that owns this resource
	Owner string `json:"owner"`
	// The connector type
	ConnectorTypeId string `json:"connector_type_id"`
	// The Connector configuration payload
	Connector map[string]interface{} `json:"connector"`
	// A detailed status message in case there is a problem with the connector
	StatusMessage *string `json:"status_message,omitempty"`
}

// NewSourceConnectorResponse instantiates a new SourceConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceConnectorResponse(kind string, id string, name string, href string, submittedAt time.Time, status ManagedResourceStatus, owner string, connectorTypeId string, connector map[string]interface{}) *SourceConnectorResponse {
	this := SourceConnectorResponse{}
	this.Kind = kind
	this.Id = id
	this.Name = name
	this.Href = href
	this.SubmittedAt = submittedAt
	this.Status = status
	this.Owner = owner
	this.ConnectorTypeId = connectorTypeId
	this.Connector = connector
	return &this
}

// NewSourceConnectorResponseWithDefaults instantiates a new SourceConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceConnectorResponseWithDefaults() *SourceConnectorResponse {
	this := SourceConnectorResponse{}
	return &this
}

// GetKind returns the Kind field value
func (o *SourceConnectorResponse) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *SourceConnectorResponse) GetKindOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *SourceConnectorResponse) SetKind(v string) {
	o.Kind = v
}

// GetId returns the Id field value
func (o *SourceConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SourceConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SourceConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SourceConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetHref returns the Href field value
func (o *SourceConnectorResponse) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *SourceConnectorResponse) GetHrefOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *SourceConnectorResponse) SetHref(v string) {
	o.Href = v
}

// GetSubmittedAt returns the SubmittedAt field value
func (o *SourceConnectorResponse) GetSubmittedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SubmittedAt
}

// GetSubmittedAtOk returns a tuple with the SubmittedAt field value
// and a boolean to check if the value has been set.
func (o *SourceConnectorResponse) GetSubmittedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SubmittedAt, true
}

// SetSubmittedAt sets field value
func (o *SourceConnectorResponse) SetSubmittedAt(v time.Time) {
	o.SubmittedAt = v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *SourceConnectorResponse) GetPublishedAt() time.Time {
	if o == nil || isNil(o.PublishedAt) {
		var ret time.Time
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceConnectorResponse) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.PublishedAt) {
    return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *SourceConnectorResponse) HasPublishedAt() bool {
	if o != nil && !isNil(o.PublishedAt) {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given time.Time and assigns it to the PublishedAt field.
func (o *SourceConnectorResponse) SetPublishedAt(v time.Time) {
	o.PublishedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SourceConnectorResponse) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceConnectorResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SourceConnectorResponse) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *SourceConnectorResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetStatus returns the Status field value
func (o *SourceConnectorResponse) GetStatus() ManagedResourceStatus {
	if o == nil {
		var ret ManagedResourceStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SourceConnectorResponse) GetStatusOk() (*ManagedResourceStatus, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SourceConnectorResponse) SetStatus(v ManagedResourceStatus) {
	o.Status = v
}

// GetOwner returns the Owner field value
func (o *SourceConnectorResponse) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *SourceConnectorResponse) GetOwnerOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *SourceConnectorResponse) SetOwner(v string) {
	o.Owner = v
}

// GetConnectorTypeId returns the ConnectorTypeId field value
func (o *SourceConnectorResponse) GetConnectorTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorTypeId
}

// GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field value
// and a boolean to check if the value has been set.
func (o *SourceConnectorResponse) GetConnectorTypeIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConnectorTypeId, true
}

// SetConnectorTypeId sets field value
func (o *SourceConnectorResponse) SetConnectorTypeId(v string) {
	o.ConnectorTypeId = v
}

// GetConnector returns the Connector field value
func (o *SourceConnectorResponse) GetConnector() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value
// and a boolean to check if the value has been set.
func (o *SourceConnectorResponse) GetConnectorOk() (map[string]interface{}, bool) {
	if o == nil {
    return map[string]interface{}{}, false
	}
	return o.Connector, true
}

// SetConnector sets field value
func (o *SourceConnectorResponse) SetConnector(v map[string]interface{}) {
	o.Connector = v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *SourceConnectorResponse) GetStatusMessage() string {
	if o == nil || isNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceConnectorResponse) GetStatusMessageOk() (*string, bool) {
	if o == nil || isNil(o.StatusMessage) {
    return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *SourceConnectorResponse) HasStatusMessage() bool {
	if o != nil && !isNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *SourceConnectorResponse) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

func (o SourceConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["submitted_at"] = o.SubmittedAt
	}
	if !isNil(o.PublishedAt) {
		toSerialize["published_at"] = o.PublishedAt
	}
	if !isNil(o.ModifiedAt) {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["owner"] = o.Owner
	}
	if true {
		toSerialize["connector_type_id"] = o.ConnectorTypeId
	}
	if true {
		toSerialize["connector"] = o.Connector
	}
	if !isNil(o.StatusMessage) {
		toSerialize["status_message"] = o.StatusMessage
	}
	return json.Marshal(toSerialize)
}

type NullableSourceConnectorResponse struct {
	value *SourceConnectorResponse
	isSet bool
}

func (v NullableSourceConnectorResponse) Get() *SourceConnectorResponse {
	return v.value
}

func (v *NullableSourceConnectorResponse) Set(val *SourceConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceConnectorResponse(val *SourceConnectorResponse) *NullableSourceConnectorResponse {
	return &NullableSourceConnectorResponse{value: val, isSet: true}
}

func (v NullableSourceConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


