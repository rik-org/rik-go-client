# WorkloadDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** | Determine the kind of object you want to create | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Spec** | Pointer to [**WorkloadDefinitionSpec**](WorkloadDefinitionSpec.md) |  | [optional] 

## Methods

### NewWorkloadDefinition

`func NewWorkloadDefinition() *WorkloadDefinition`

NewWorkloadDefinition instantiates a new WorkloadDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadDefinitionWithDefaults

`func NewWorkloadDefinitionWithDefaults() *WorkloadDefinition`

NewWorkloadDefinitionWithDefaults instantiates a new WorkloadDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *WorkloadDefinition) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *WorkloadDefinition) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *WorkloadDefinition) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *WorkloadDefinition) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *WorkloadDefinition) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WorkloadDefinition) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WorkloadDefinition) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WorkloadDefinition) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *WorkloadDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpec

`func (o *WorkloadDefinition) GetSpec() WorkloadDefinitionSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *WorkloadDefinition) GetSpecOk() (*WorkloadDefinitionSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *WorkloadDefinition) SetSpec(v WorkloadDefinitionSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *WorkloadDefinition) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


