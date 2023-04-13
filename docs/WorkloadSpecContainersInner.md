# WorkloadSpecContainersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Env** | Pointer to [**[]ContainerWorkloadContainersInnerEnvInner**](ContainerWorkloadContainersInnerEnvInner.md) |  | [optional] 
**Ports** | Pointer to [**ContainerWorkloadContainersInnerNetwork**](ContainerWorkloadContainersInnerNetwork.md) |  | [optional] 

## Methods

### NewWorkloadSpecContainersInner

`func NewWorkloadSpecContainersInner() *WorkloadSpecContainersInner`

NewWorkloadSpecContainersInner instantiates a new WorkloadSpecContainersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadSpecContainersInnerWithDefaults

`func NewWorkloadSpecContainersInnerWithDefaults() *WorkloadSpecContainersInner`

NewWorkloadSpecContainersInnerWithDefaults instantiates a new WorkloadSpecContainersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkloadSpecContainersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadSpecContainersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadSpecContainersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadSpecContainersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *WorkloadSpecContainersInner) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WorkloadSpecContainersInner) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WorkloadSpecContainersInner) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *WorkloadSpecContainersInner) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetEnv

`func (o *WorkloadSpecContainersInner) GetEnv() []ContainerWorkloadContainersInnerEnvInner`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *WorkloadSpecContainersInner) GetEnvOk() (*[]ContainerWorkloadContainersInnerEnvInner, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *WorkloadSpecContainersInner) SetEnv(v []ContainerWorkloadContainersInnerEnvInner)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *WorkloadSpecContainersInner) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetPorts

`func (o *WorkloadSpecContainersInner) GetPorts() ContainerWorkloadContainersInnerNetwork`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *WorkloadSpecContainersInner) GetPortsOk() (*ContainerWorkloadContainersInnerNetwork, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *WorkloadSpecContainersInner) SetPorts(v ContainerWorkloadContainersInnerNetwork)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *WorkloadSpecContainersInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


