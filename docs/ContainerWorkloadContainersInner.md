# ContainerWorkloadContainersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Env** | Pointer to [**[]ContainerWorkloadContainersInnerEnvInner**](ContainerWorkloadContainersInnerEnvInner.md) |  | [optional] 
**Network** | Pointer to [**ContainerWorkloadContainersInnerNetwork**](ContainerWorkloadContainersInnerNetwork.md) |  | [optional] 

## Methods

### NewContainerWorkloadContainersInner

`func NewContainerWorkloadContainersInner() *ContainerWorkloadContainersInner`

NewContainerWorkloadContainersInner instantiates a new ContainerWorkloadContainersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerWorkloadContainersInnerWithDefaults

`func NewContainerWorkloadContainersInnerWithDefaults() *ContainerWorkloadContainersInner`

NewContainerWorkloadContainersInnerWithDefaults instantiates a new ContainerWorkloadContainersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerWorkloadContainersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerWorkloadContainersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerWorkloadContainersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContainerWorkloadContainersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *ContainerWorkloadContainersInner) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ContainerWorkloadContainersInner) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ContainerWorkloadContainersInner) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ContainerWorkloadContainersInner) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetEnv

`func (o *ContainerWorkloadContainersInner) GetEnv() []ContainerWorkloadContainersInnerEnvInner`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ContainerWorkloadContainersInner) GetEnvOk() (*[]ContainerWorkloadContainersInnerEnvInner, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ContainerWorkloadContainersInner) SetEnv(v []ContainerWorkloadContainersInnerEnvInner)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ContainerWorkloadContainersInner) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetNetwork

`func (o *ContainerWorkloadContainersInner) GetNetwork() ContainerWorkloadContainersInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ContainerWorkloadContainersInner) GetNetworkOk() (*ContainerWorkloadContainersInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ContainerWorkloadContainersInner) SetNetwork(v ContainerWorkloadContainersInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ContainerWorkloadContainersInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


