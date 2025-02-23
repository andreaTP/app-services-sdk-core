/*
Connector Management API

Connector Management API is a REST API to manage connectors.

API version: 0.1.0
Contact: rhosak-support@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
)

// ConnectorNamespaceQuota struct for ConnectorNamespaceQuota
type ConnectorNamespaceQuota struct {
	Connectors *int32 `json:"connectors,omitempty"`
	// Memory quota for limits or requests
	MemoryRequests *string `json:"memory_requests,omitempty"`
	// Memory quota for limits or requests
	MemoryLimits *string `json:"memory_limits,omitempty"`
	// CPU quota for limits or requests
	CpuRequests *string `json:"cpu_requests,omitempty"`
	// CPU quota for limits or requests
	CpuLimits *string `json:"cpu_limits,omitempty"`
}

// NewConnectorNamespaceQuota instantiates a new ConnectorNamespaceQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorNamespaceQuota() *ConnectorNamespaceQuota {
	this := ConnectorNamespaceQuota{}
	return &this
}

// NewConnectorNamespaceQuotaWithDefaults instantiates a new ConnectorNamespaceQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorNamespaceQuotaWithDefaults() *ConnectorNamespaceQuota {
	this := ConnectorNamespaceQuota{}
	return &this
}

// GetConnectors returns the Connectors field value if set, zero value otherwise.
func (o *ConnectorNamespaceQuota) GetConnectors() int32 {
	if o == nil || isNil(o.Connectors) {
		var ret int32
		return ret
	}
	return *o.Connectors
}

// GetConnectorsOk returns a tuple with the Connectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceQuota) GetConnectorsOk() (*int32, bool) {
	if o == nil || isNil(o.Connectors) {
    return nil, false
	}
	return o.Connectors, true
}

// HasConnectors returns a boolean if a field has been set.
func (o *ConnectorNamespaceQuota) HasConnectors() bool {
	if o != nil && !isNil(o.Connectors) {
		return true
	}

	return false
}

// SetConnectors gets a reference to the given int32 and assigns it to the Connectors field.
func (o *ConnectorNamespaceQuota) SetConnectors(v int32) {
	o.Connectors = &v
}

// GetMemoryRequests returns the MemoryRequests field value if set, zero value otherwise.
func (o *ConnectorNamespaceQuota) GetMemoryRequests() string {
	if o == nil || isNil(o.MemoryRequests) {
		var ret string
		return ret
	}
	return *o.MemoryRequests
}

// GetMemoryRequestsOk returns a tuple with the MemoryRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceQuota) GetMemoryRequestsOk() (*string, bool) {
	if o == nil || isNil(o.MemoryRequests) {
    return nil, false
	}
	return o.MemoryRequests, true
}

// HasMemoryRequests returns a boolean if a field has been set.
func (o *ConnectorNamespaceQuota) HasMemoryRequests() bool {
	if o != nil && !isNil(o.MemoryRequests) {
		return true
	}

	return false
}

// SetMemoryRequests gets a reference to the given string and assigns it to the MemoryRequests field.
func (o *ConnectorNamespaceQuota) SetMemoryRequests(v string) {
	o.MemoryRequests = &v
}

// GetMemoryLimits returns the MemoryLimits field value if set, zero value otherwise.
func (o *ConnectorNamespaceQuota) GetMemoryLimits() string {
	if o == nil || isNil(o.MemoryLimits) {
		var ret string
		return ret
	}
	return *o.MemoryLimits
}

// GetMemoryLimitsOk returns a tuple with the MemoryLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceQuota) GetMemoryLimitsOk() (*string, bool) {
	if o == nil || isNil(o.MemoryLimits) {
    return nil, false
	}
	return o.MemoryLimits, true
}

// HasMemoryLimits returns a boolean if a field has been set.
func (o *ConnectorNamespaceQuota) HasMemoryLimits() bool {
	if o != nil && !isNil(o.MemoryLimits) {
		return true
	}

	return false
}

// SetMemoryLimits gets a reference to the given string and assigns it to the MemoryLimits field.
func (o *ConnectorNamespaceQuota) SetMemoryLimits(v string) {
	o.MemoryLimits = &v
}

// GetCpuRequests returns the CpuRequests field value if set, zero value otherwise.
func (o *ConnectorNamespaceQuota) GetCpuRequests() string {
	if o == nil || isNil(o.CpuRequests) {
		var ret string
		return ret
	}
	return *o.CpuRequests
}

// GetCpuRequestsOk returns a tuple with the CpuRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceQuota) GetCpuRequestsOk() (*string, bool) {
	if o == nil || isNil(o.CpuRequests) {
    return nil, false
	}
	return o.CpuRequests, true
}

// HasCpuRequests returns a boolean if a field has been set.
func (o *ConnectorNamespaceQuota) HasCpuRequests() bool {
	if o != nil && !isNil(o.CpuRequests) {
		return true
	}

	return false
}

// SetCpuRequests gets a reference to the given string and assigns it to the CpuRequests field.
func (o *ConnectorNamespaceQuota) SetCpuRequests(v string) {
	o.CpuRequests = &v
}

// GetCpuLimits returns the CpuLimits field value if set, zero value otherwise.
func (o *ConnectorNamespaceQuota) GetCpuLimits() string {
	if o == nil || isNil(o.CpuLimits) {
		var ret string
		return ret
	}
	return *o.CpuLimits
}

// GetCpuLimitsOk returns a tuple with the CpuLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceQuota) GetCpuLimitsOk() (*string, bool) {
	if o == nil || isNil(o.CpuLimits) {
    return nil, false
	}
	return o.CpuLimits, true
}

// HasCpuLimits returns a boolean if a field has been set.
func (o *ConnectorNamespaceQuota) HasCpuLimits() bool {
	if o != nil && !isNil(o.CpuLimits) {
		return true
	}

	return false
}

// SetCpuLimits gets a reference to the given string and assigns it to the CpuLimits field.
func (o *ConnectorNamespaceQuota) SetCpuLimits(v string) {
	o.CpuLimits = &v
}

func (o ConnectorNamespaceQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Connectors) {
		toSerialize["connectors"] = o.Connectors
	}
	if !isNil(o.MemoryRequests) {
		toSerialize["memory_requests"] = o.MemoryRequests
	}
	if !isNil(o.MemoryLimits) {
		toSerialize["memory_limits"] = o.MemoryLimits
	}
	if !isNil(o.CpuRequests) {
		toSerialize["cpu_requests"] = o.CpuRequests
	}
	if !isNil(o.CpuLimits) {
		toSerialize["cpu_limits"] = o.CpuLimits
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorNamespaceQuota struct {
	value *ConnectorNamespaceQuota
	isSet bool
}

func (v NullableConnectorNamespaceQuota) Get() *ConnectorNamespaceQuota {
	return v.value
}

func (v *NullableConnectorNamespaceQuota) Set(val *ConnectorNamespaceQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorNamespaceQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorNamespaceQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorNamespaceQuota(val *ConnectorNamespaceQuota) *NullableConnectorNamespaceQuota {
	return &NullableConnectorNamespaceQuota{value: val, isSet: true}
}

func (v NullableConnectorNamespaceQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorNamespaceQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


