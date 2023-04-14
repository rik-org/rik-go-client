# WorkloadMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**Workload**](Workload.md) |  | [optional] 

## Methods

### NewWorkloadMeta

`func NewWorkloadMeta() *WorkloadMeta`

NewWorkloadMeta instantiates a new WorkloadMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadMetaWithDefaults

`func NewWorkloadMetaWithDefaults() *WorkloadMeta`

NewWorkloadMetaWithDefaults instantiates a new WorkloadMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkloadMeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkloadMeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkloadMeta) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkloadMeta) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkloadMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *WorkloadMeta) GetValue() Workload`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkloadMeta) GetValueOk() (*Workload, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkloadMeta) SetValue(v Workload)`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkloadMeta) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


