/*
Kafka Management API

Kafka Management API is a REST API to manage Kafka instances

API version: 1.15.0
Contact: rhosak-support@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// checks if the EnterpriseClusterAllOfCapacityInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnterpriseClusterAllOfCapacityInformation{}

// EnterpriseClusterAllOfCapacityInformation Returns the capacity related information
type EnterpriseClusterAllOfCapacityInformation struct {
	// The kafka machine pool node count provided during cluster registration
	KafkaMachinePoolNodeCount int32 `json:"kafka_machine_pool_node_count"`
	// The maximum number of Kafka streaming units that can be created on this cluster
	MaximumKafkaStreamingUnits int32 `json:"maximum_kafka_streaming_units"`
	// The remaining number of Kafka streaming units that can be still be created on this cluster
	RemainingKafkaStreamingUnits int32 `json:"remaining_kafka_streaming_units"`
	// The number of Kafka streaming units that have been consumed on this cluster
	ConsumedKafkaStreamingUnits int32 `json:"consumed_kafka_streaming_units"`
}

// NewEnterpriseClusterAllOfCapacityInformation instantiates a new EnterpriseClusterAllOfCapacityInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseClusterAllOfCapacityInformation(kafkaMachinePoolNodeCount int32, maximumKafkaStreamingUnits int32, remainingKafkaStreamingUnits int32, consumedKafkaStreamingUnits int32) *EnterpriseClusterAllOfCapacityInformation {
	this := EnterpriseClusterAllOfCapacityInformation{}
	this.KafkaMachinePoolNodeCount = kafkaMachinePoolNodeCount
	this.MaximumKafkaStreamingUnits = maximumKafkaStreamingUnits
	this.RemainingKafkaStreamingUnits = remainingKafkaStreamingUnits
	this.ConsumedKafkaStreamingUnits = consumedKafkaStreamingUnits
	return &this
}

// NewEnterpriseClusterAllOfCapacityInformationWithDefaults instantiates a new EnterpriseClusterAllOfCapacityInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseClusterAllOfCapacityInformationWithDefaults() *EnterpriseClusterAllOfCapacityInformation {
	this := EnterpriseClusterAllOfCapacityInformation{}
	return &this
}

// GetKafkaMachinePoolNodeCount returns the KafkaMachinePoolNodeCount field value
func (o *EnterpriseClusterAllOfCapacityInformation) GetKafkaMachinePoolNodeCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.KafkaMachinePoolNodeCount
}

// GetKafkaMachinePoolNodeCountOk returns a tuple with the KafkaMachinePoolNodeCount field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterAllOfCapacityInformation) GetKafkaMachinePoolNodeCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KafkaMachinePoolNodeCount, true
}

// SetKafkaMachinePoolNodeCount sets field value
func (o *EnterpriseClusterAllOfCapacityInformation) SetKafkaMachinePoolNodeCount(v int32) {
	o.KafkaMachinePoolNodeCount = v
}

// GetMaximumKafkaStreamingUnits returns the MaximumKafkaStreamingUnits field value
func (o *EnterpriseClusterAllOfCapacityInformation) GetMaximumKafkaStreamingUnits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaximumKafkaStreamingUnits
}

// GetMaximumKafkaStreamingUnitsOk returns a tuple with the MaximumKafkaStreamingUnits field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterAllOfCapacityInformation) GetMaximumKafkaStreamingUnitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumKafkaStreamingUnits, true
}

// SetMaximumKafkaStreamingUnits sets field value
func (o *EnterpriseClusterAllOfCapacityInformation) SetMaximumKafkaStreamingUnits(v int32) {
	o.MaximumKafkaStreamingUnits = v
}

// GetRemainingKafkaStreamingUnits returns the RemainingKafkaStreamingUnits field value
func (o *EnterpriseClusterAllOfCapacityInformation) GetRemainingKafkaStreamingUnits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RemainingKafkaStreamingUnits
}

// GetRemainingKafkaStreamingUnitsOk returns a tuple with the RemainingKafkaStreamingUnits field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterAllOfCapacityInformation) GetRemainingKafkaStreamingUnitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingKafkaStreamingUnits, true
}

// SetRemainingKafkaStreamingUnits sets field value
func (o *EnterpriseClusterAllOfCapacityInformation) SetRemainingKafkaStreamingUnits(v int32) {
	o.RemainingKafkaStreamingUnits = v
}

// GetConsumedKafkaStreamingUnits returns the ConsumedKafkaStreamingUnits field value
func (o *EnterpriseClusterAllOfCapacityInformation) GetConsumedKafkaStreamingUnits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConsumedKafkaStreamingUnits
}

// GetConsumedKafkaStreamingUnitsOk returns a tuple with the ConsumedKafkaStreamingUnits field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterAllOfCapacityInformation) GetConsumedKafkaStreamingUnitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumedKafkaStreamingUnits, true
}

// SetConsumedKafkaStreamingUnits sets field value
func (o *EnterpriseClusterAllOfCapacityInformation) SetConsumedKafkaStreamingUnits(v int32) {
	o.ConsumedKafkaStreamingUnits = v
}

func (o EnterpriseClusterAllOfCapacityInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnterpriseClusterAllOfCapacityInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kafka_machine_pool_node_count"] = o.KafkaMachinePoolNodeCount
	toSerialize["maximum_kafka_streaming_units"] = o.MaximumKafkaStreamingUnits
	toSerialize["remaining_kafka_streaming_units"] = o.RemainingKafkaStreamingUnits
	toSerialize["consumed_kafka_streaming_units"] = o.ConsumedKafkaStreamingUnits
	return toSerialize, nil
}

type NullableEnterpriseClusterAllOfCapacityInformation struct {
	value *EnterpriseClusterAllOfCapacityInformation
	isSet bool
}

func (v NullableEnterpriseClusterAllOfCapacityInformation) Get() *EnterpriseClusterAllOfCapacityInformation {
	return v.value
}

func (v *NullableEnterpriseClusterAllOfCapacityInformation) Set(val *EnterpriseClusterAllOfCapacityInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseClusterAllOfCapacityInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseClusterAllOfCapacityInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseClusterAllOfCapacityInformation(val *EnterpriseClusterAllOfCapacityInformation) *NullableEnterpriseClusterAllOfCapacityInformation {
	return &NullableEnterpriseClusterAllOfCapacityInformation{value: val, isSet: true}
}

func (v NullableEnterpriseClusterAllOfCapacityInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseClusterAllOfCapacityInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


