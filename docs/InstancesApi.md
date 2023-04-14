# \InstancesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkloadInstance**](InstancesApi.md#CreateWorkloadInstance) | **Post** /api/v0/instances.create | Create a new instance for a given workload
[**DeleteInstance**](InstancesApi.md#DeleteInstance) | **Post** /api/v0/instances.delete | 
[**GetInstances**](InstancesApi.md#GetInstances) | **Get** /api/v0/instances.list | 



## CreateWorkloadInstance

> CreateWorkloadInstance(ctx).InstanceDefinition(instanceDefinition).Execute()

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
    instanceDefinition := *openapiclient.NewInstanceDefinition("cobra-citer-1923", "c63f1351-d371-4700-81a4-ac97359bf5a3") // InstanceDefinition |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InstancesApi.CreateWorkloadInstance(context.Background()).InstanceDefinition(instanceDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.CreateWorkloadInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkloadInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceDefinition** | [**InstanceDefinition**](InstanceDefinition.md) |  | 

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

> GetWorkloadInstances200Response GetInstances(ctx).Execute()





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
    // response from `GetInstances`: GetWorkloadInstances200Response
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.GetInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstancesRequest struct via the builder pattern


### Return type

[**GetWorkloadInstances200Response**](GetWorkloadInstances200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

