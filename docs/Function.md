# Function

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Execution** | Pointer to [**FunctionExecution**](FunctionExecution.md) |  | [optional] 

## Methods

### NewFunction

`func NewFunction() *Function`

NewFunction instantiates a new Function object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionWithDefaults

`func NewFunctionWithDefaults() *Function`

NewFunctionWithDefaults instantiates a new Function object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecution

`func (o *Function) GetExecution() FunctionExecution`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *Function) GetExecutionOk() (*FunctionExecution, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *Function) SetExecution(v FunctionExecution)`

SetExecution sets Execution field to given value.

### HasExecution

`func (o *Function) HasExecution() bool`

HasExecution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


