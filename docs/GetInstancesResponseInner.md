# GetInstancesResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Simple helper schema to define an UUID | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**Instance**](Instance.md) |  | [optional] 

## Methods

### NewGetInstancesResponseInner

`func NewGetInstancesResponseInner() *GetInstancesResponseInner`

NewGetInstancesResponseInner instantiates a new GetInstancesResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstancesResponseInnerWithDefaults

`func NewGetInstancesResponseInnerWithDefaults() *GetInstancesResponseInner`

NewGetInstancesResponseInnerWithDefaults instantiates a new GetInstancesResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInstancesResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstancesResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstancesResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstancesResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetInstancesResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInstancesResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInstancesResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInstancesResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *GetInstancesResponseInner) GetValue() Instance`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetInstancesResponseInner) GetValueOk() (*Instance, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetInstancesResponseInner) SetValue(v Instance)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetInstancesResponseInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


