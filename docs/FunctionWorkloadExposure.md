# FunctionWorkloadExposure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** |  | [optional] 
**TargetPort** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** | The type of the exposure | [optional] 

## Methods

### NewFunctionWorkloadExposure

`func NewFunctionWorkloadExposure() *FunctionWorkloadExposure`

NewFunctionWorkloadExposure instantiates a new FunctionWorkloadExposure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionWorkloadExposureWithDefaults

`func NewFunctionWorkloadExposureWithDefaults() *FunctionWorkloadExposure`

NewFunctionWorkloadExposureWithDefaults instantiates a new FunctionWorkloadExposure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *FunctionWorkloadExposure) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FunctionWorkloadExposure) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FunctionWorkloadExposure) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *FunctionWorkloadExposure) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTargetPort

`func (o *FunctionWorkloadExposure) GetTargetPort() int32`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *FunctionWorkloadExposure) GetTargetPortOk() (*int32, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *FunctionWorkloadExposure) SetTargetPort(v int32)`

SetTargetPort sets TargetPort field to given value.

### HasTargetPort

`func (o *FunctionWorkloadExposure) HasTargetPort() bool`

HasTargetPort returns a boolean if a field has been set.

### GetType

`func (o *FunctionWorkloadExposure) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionWorkloadExposure) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionWorkloadExposure) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FunctionWorkloadExposure) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


