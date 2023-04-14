# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Simple helper schema to define an UUID | [optional] 
**Kind** | Pointer to [**Kind**](Kind.md) |  | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**WorkloadId** | Pointer to **string** | Simple helper schema to define an UUID | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Spec** | Pointer to [**InstanceSpec**](InstanceSpec.md) |  | [optional] 

## Methods

### NewInstance

`func NewInstance() *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Instance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Instance) GetKind() Kind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Instance) GetKindOk() (*Kind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Instance) SetKind(v Kind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Instance) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetStatus

`func (o *Instance) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Instance) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Instance) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Instance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWorkloadId

`func (o *Instance) GetWorkloadId() string`

GetWorkloadId returns the WorkloadId field if non-nil, zero value otherwise.

### GetWorkloadIdOk

`func (o *Instance) GetWorkloadIdOk() (*string, bool)`

GetWorkloadIdOk returns a tuple with the WorkloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadId

`func (o *Instance) SetWorkloadId(v string)`

SetWorkloadId sets WorkloadId field to given value.

### HasWorkloadId

`func (o *Instance) HasWorkloadId() bool`

HasWorkloadId returns a boolean if a field has been set.

### GetNamespace

`func (o *Instance) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Instance) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Instance) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Instance) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSpec

`func (o *Instance) GetSpec() InstanceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Instance) GetSpecOk() (*InstanceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Instance) SetSpec(v InstanceSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *Instance) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


