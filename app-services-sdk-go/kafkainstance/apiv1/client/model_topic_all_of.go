/*
Kafka Instance API

API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs

API version: 0.13.0-SNAPSHOT
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// TopicAllOf Kafka Topic (A feed where records are stored and published)
type TopicAllOf struct {
	// The name of the topic.
	Name *string `json:"name,omitempty"`
	IsInternal *bool `json:"isInternal,omitempty"`
	// Partitions for this topic.
	Partitions []Partition `json:"partitions,omitempty"`
	// Topic configuration entry.
	Config []ConfigEntry `json:"config,omitempty"`
}

// NewTopicAllOf instantiates a new TopicAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopicAllOf() *TopicAllOf {
	this := TopicAllOf{}
	return &this
}

// NewTopicAllOfWithDefaults instantiates a new TopicAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicAllOfWithDefaults() *TopicAllOf {
	this := TopicAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TopicAllOf) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicAllOf) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TopicAllOf) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TopicAllOf) SetName(v string) {
	o.Name = &v
}

// GetIsInternal returns the IsInternal field value if set, zero value otherwise.
func (o *TopicAllOf) GetIsInternal() bool {
	if o == nil || isNil(o.IsInternal) {
		var ret bool
		return ret
	}
	return *o.IsInternal
}

// GetIsInternalOk returns a tuple with the IsInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicAllOf) GetIsInternalOk() (*bool, bool) {
	if o == nil || isNil(o.IsInternal) {
    return nil, false
	}
	return o.IsInternal, true
}

// HasIsInternal returns a boolean if a field has been set.
func (o *TopicAllOf) HasIsInternal() bool {
	if o != nil && !isNil(o.IsInternal) {
		return true
	}

	return false
}

// SetIsInternal gets a reference to the given bool and assigns it to the IsInternal field.
func (o *TopicAllOf) SetIsInternal(v bool) {
	o.IsInternal = &v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *TopicAllOf) GetPartitions() []Partition {
	if o == nil || isNil(o.Partitions) {
		var ret []Partition
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicAllOf) GetPartitionsOk() ([]Partition, bool) {
	if o == nil || isNil(o.Partitions) {
    return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *TopicAllOf) HasPartitions() bool {
	if o != nil && !isNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []Partition and assigns it to the Partitions field.
func (o *TopicAllOf) SetPartitions(v []Partition) {
	o.Partitions = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *TopicAllOf) GetConfig() []ConfigEntry {
	if o == nil || isNil(o.Config) {
		var ret []ConfigEntry
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicAllOf) GetConfigOk() ([]ConfigEntry, bool) {
	if o == nil || isNil(o.Config) {
    return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *TopicAllOf) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given []ConfigEntry and assigns it to the Config field.
func (o *TopicAllOf) SetConfig(v []ConfigEntry) {
	o.Config = v
}

func (o TopicAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.IsInternal) {
		toSerialize["isInternal"] = o.IsInternal
	}
	if !isNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableTopicAllOf struct {
	value *TopicAllOf
	isSet bool
}

func (v NullableTopicAllOf) Get() *TopicAllOf {
	return v.value
}

func (v *NullableTopicAllOf) Set(val *TopicAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicAllOf(val *TopicAllOf) *NullableTopicAllOf {
	return &NullableTopicAllOf{value: val, isSet: true}
}

func (v NullableTopicAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopicAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


