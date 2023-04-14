# WorkloadSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Function** | Pointer to [**Function**](Function.md) |  | [optional] 
**Containers** | Pointer to [**[]Container**](Container.md) |  | [optional] 

## Methods

### NewWorkloadSpec

`func NewWorkloadSpec() *WorkloadSpec`

NewWorkloadSpec instantiates a new WorkloadSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadSpecWithDefaults

`func NewWorkloadSpecWithDefaults() *WorkloadSpec`

NewWorkloadSpecWithDefaults instantiates a new WorkloadSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunction

`func (o *WorkloadSpec) GetFunction() Function`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *WorkloadSpec) GetFunctionOk() (*Function, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *WorkloadSpec) SetFunction(v Function)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *WorkloadSpec) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetContainers

`func (o *WorkloadSpec) GetContainers() []Container`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *WorkloadSpec) GetContainersOk() (*[]Container, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *WorkloadSpec) SetContainers(v []Container)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *WorkloadSpec) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


