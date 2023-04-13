# InstanceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | Pointer to [**[]ContainerWorkload**](ContainerWorkload.md) |  | [optional] 
**Function** | Pointer to [**InstanceSpecFunction**](InstanceSpecFunction.md) |  | [optional] 

## Methods

### NewInstanceSpec

`func NewInstanceSpec() *InstanceSpec`

NewInstanceSpec instantiates a new InstanceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceSpecWithDefaults

`func NewInstanceSpecWithDefaults() *InstanceSpec`

NewInstanceSpecWithDefaults instantiates a new InstanceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *InstanceSpec) GetContainers() []ContainerWorkload`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *InstanceSpec) GetContainersOk() (*[]ContainerWorkload, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *InstanceSpec) SetContainers(v []ContainerWorkload)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *InstanceSpec) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetFunction

`func (o *InstanceSpec) GetFunction() InstanceSpecFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *InstanceSpec) GetFunctionOk() (*InstanceSpecFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *InstanceSpec) SetFunction(v InstanceSpecFunction)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *InstanceSpec) HasFunction() bool`

HasFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


