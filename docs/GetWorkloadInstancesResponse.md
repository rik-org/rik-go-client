# GetWorkloadInstancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**[]Instance**](Instance.md) |  | [optional] 

## Methods

### NewGetWorkloadInstancesResponse

`func NewGetWorkloadInstancesResponse() *GetWorkloadInstancesResponse`

NewGetWorkloadInstancesResponse instantiates a new GetWorkloadInstancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWorkloadInstancesResponseWithDefaults

`func NewGetWorkloadInstancesResponseWithDefaults() *GetWorkloadInstancesResponse`

NewGetWorkloadInstancesResponseWithDefaults instantiates a new GetWorkloadInstancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *GetWorkloadInstancesResponse) GetInstances() []Instance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *GetWorkloadInstancesResponse) GetInstancesOk() (*[]Instance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *GetWorkloadInstancesResponse) SetInstances(v []Instance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *GetWorkloadInstancesResponse) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


