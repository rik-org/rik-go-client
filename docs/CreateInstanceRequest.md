# CreateInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkloadId** | **string** | Simple helper schema to define an UUID | 
**Name** | Pointer to **string** | The name of your instance. | [optional] 
**Replicas** | Pointer to **int32** | The number of instances to deploy. | [optional] [default to 1]

## Methods

### NewCreateInstanceRequest

`func NewCreateInstanceRequest(workloadId string, ) *CreateInstanceRequest`

NewCreateInstanceRequest instantiates a new CreateInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceRequestWithDefaults

`func NewCreateInstanceRequestWithDefaults() *CreateInstanceRequest`

NewCreateInstanceRequestWithDefaults instantiates a new CreateInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkloadId

`func (o *CreateInstanceRequest) GetWorkloadId() string`

GetWorkloadId returns the WorkloadId field if non-nil, zero value otherwise.

### GetWorkloadIdOk

`func (o *CreateInstanceRequest) GetWorkloadIdOk() (*string, bool)`

GetWorkloadIdOk returns a tuple with the WorkloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadId

`func (o *CreateInstanceRequest) SetWorkloadId(v string)`

SetWorkloadId sets WorkloadId field to given value.


### GetName

`func (o *CreateInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInstanceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateInstanceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReplicas

`func (o *CreateInstanceRequest) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *CreateInstanceRequest) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *CreateInstanceRequest) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *CreateInstanceRequest) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


