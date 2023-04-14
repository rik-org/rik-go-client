# \InstancesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkloadInstance**](InstancesApi.md#CreateWorkloadInstance) | **Post** /api/v0/instances.create | Create a new instance for a given workload
[**DeleteInstance**](InstancesApi.md#DeleteInstance) | **Post** /api/v0/instances.delete | 
[**GetInstances**](InstancesApi.md#GetInstances) | **Get** /api/v0/instances.list | 
[**GetWorkloadInstances**](InstancesApi.md#GetWorkloadInstances) | **Get** /api/v0/workloads.instances/{workloadId} | Get all the instances for the given workload



## CreateWorkloadInstance

> []string CreateWorkloadInstance(ctx).CreateInstanceRequest(createInstanceRequest).Execute()

Create a new instance for a given workload



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/rik-org/rik-go-client"
)

func main() {
    createInstanceRequest := *openapiclient.NewCreateInstanceRequest("c540eaf0-e41b-4de8-bbda-d1c815443b6e", "my-super-instance-of-my-workload") // CreateInstanceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstancesApi.CreateWorkloadInstance(context.Background()).CreateInstanceRequest(createInstanceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.CreateWorkloadInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkloadInstance`: []string
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.CreateWorkloadInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkloadInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInstanceRequest** | [**CreateInstanceRequest**](CreateInstanceRequest.md) |  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> DeleteInstance(ctx).DeleteInstanceRequest(deleteInstanceRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/rik-org/rik-go-client"
)

func main() {
    deleteInstanceRequest := *openapiclient.NewDeleteInstanceRequest() // DeleteInstanceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InstancesApi.DeleteInstance(context.Background()).DeleteInstanceRequest(deleteInstanceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.DeleteInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteInstanceRequest** | [**DeleteInstanceRequest**](DeleteInstanceRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstances

> GetInstances200Response GetInstances(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/rik-org/rik-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstancesApi.GetInstances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstances`: GetInstances200Response
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstancesRequest struct via the builder pattern


### Return type

[**GetInstances200Response**](GetInstances200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkloadInstances

> GetWorkloadInstancesResponse GetWorkloadInstances(ctx, workloadId).Execute()

Get all the instances for the given workload



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/rik-org/rik-go-client"
)

func main() {
    workloadId := "workloadId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstancesApi.GetWorkloadInstances(context.Background(), workloadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetWorkloadInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkloadInstances`: GetWorkloadInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetWorkloadInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workloadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWorkloadInstancesResponse**](GetWorkloadInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

