# Partition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partition** | **int32** | The partition id, unique among partitions of the same topic | 
**Replicas** | Pointer to [**[]Node**](Node.md) | List of replicas for the partition | [optional] 
**Isr** | Pointer to [**[]Node**](Node.md) | List in-sync replicas for this partition. | [optional] 
**Leader** | Pointer to [**PartitionLeader**](PartitionLeader.md) |  | [optional] 
**Id** | Pointer to **int32** | Unique id for the partition (deprecated, use &#x60;partition&#x60; instead) | [optional] 

## Methods

### NewPartition

`func NewPartition(partition int32, ) *Partition`

NewPartition instantiates a new Partition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionWithDefaults

`func NewPartitionWithDefaults() *Partition`

NewPartitionWithDefaults instantiates a new Partition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartition

`func (o *Partition) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *Partition) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *Partition) SetPartition(v int32)`

SetPartition sets Partition field to given value.


### GetReplicas

`func (o *Partition) GetReplicas() []Node`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *Partition) GetReplicasOk() (*[]Node, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *Partition) SetReplicas(v []Node)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *Partition) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetIsr

`func (o *Partition) GetIsr() []Node`

GetIsr returns the Isr field if non-nil, zero value otherwise.

### GetIsrOk

`func (o *Partition) GetIsrOk() (*[]Node, bool)`

GetIsrOk returns a tuple with the Isr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsr

`func (o *Partition) SetIsr(v []Node)`

SetIsr sets Isr field to given value.

### HasIsr

`func (o *Partition) HasIsr() bool`

HasIsr returns a boolean if a field has been set.

### GetLeader

`func (o *Partition) GetLeader() PartitionLeader`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *Partition) GetLeaderOk() (*PartitionLeader, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *Partition) SetLeader(v PartitionLeader)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *Partition) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetId

`func (o *Partition) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Partition) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Partition) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Partition) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


