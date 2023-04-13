# WorkloadDefinitionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | Pointer to [**[]ContainerWorkloadContainersInner**](ContainerWorkloadContainersInner.md) |  | [optional] 
**Function** | Pointer to [**FunctionWorkloadFunction**](FunctionWorkloadFunction.md) |  | [optional] 

## Methods

### NewWorkloadDefinitionSpec

`func NewWorkloadDefinitionSpec() *WorkloadDefinitionSpec`

NewWorkloadDefinitionSpec instantiates a new WorkloadDefinitionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadDefinitionSpecWithDefaults

`func NewWorkloadDefinitionSpecWithDefaults() *WorkloadDefinitionSpec`

NewWorkloadDefinitionSpecWithDefaults instantiates a new WorkloadDefinitionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *WorkloadDefinitionSpec) GetContainers() []ContainerWorkloadContainersInner`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *WorkloadDefinitionSpec) GetContainersOk() (*[]ContainerWorkloadContainersInner, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *WorkloadDefinitionSpec) SetContainers(v []ContainerWorkloadContainersInner)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *WorkloadDefinitionSpec) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetFunction

`func (o *WorkloadDefinitionSpec) GetFunction() FunctionWorkloadFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *WorkloadDefinitionSpec) GetFunctionOk() (*FunctionWorkloadFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *WorkloadDefinitionSpec) SetFunction(v FunctionWorkloadFunction)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *WorkloadDefinitionSpec) HasFunction() bool`

HasFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

