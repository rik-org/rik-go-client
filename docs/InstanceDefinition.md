# InstanceDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**WorkloadId** | **string** |  | 
**Replicas** | Pointer to **int32** |  | [optional] 

## Methods

### NewInstanceDefinition

`func NewInstanceDefinition(name string, workloadId string, ) *InstanceDefinition`

NewInstanceDefinition instantiates a new InstanceDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceDefinitionWithDefaults

`func NewInstanceDefinitionWithDefaults() *InstanceDefinition`

NewInstanceDefinitionWithDefaults instantiates a new InstanceDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstanceDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetWorkloadId

`func (o *InstanceDefinition) GetWorkloadId() string`

GetWorkloadId returns the WorkloadId field if non-nil, zero value otherwise.

### GetWorkloadIdOk

`func (o *InstanceDefinition) GetWorkloadIdOk() (*string, bool)`

GetWorkloadIdOk returns a tuple with the WorkloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadId

`func (o *InstanceDefinition) SetWorkloadId(v string)`

SetWorkloadId sets WorkloadId field to given value.


### GetReplicas

`func (o *InstanceDefinition) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *InstanceDefinition) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *InstanceDefinition) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *InstanceDefinition) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


