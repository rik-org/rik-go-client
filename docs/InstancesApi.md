# \InstancesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV0InstancesCreatePost**](InstancesApi.md#ApiV0InstancesCreatePost) | **Post** /api/v0/instances.create | 
[**ApiV0InstancesDeletePost**](InstancesApi.md#ApiV0InstancesDeletePost) | **Post** /api/v0/instances.delete | 
[**ApiV0InstancesListGet**](InstancesApi.md#ApiV0InstancesListGet) | **Get** /api/v0/instances.list | 
[**ApiV0WorkloadsInstancesWorkloadIdGet**](InstancesApi.md#ApiV0WorkloadsInstancesWorkloadIdGet) | **Get** /api/v0/workloads.instances/{workloadId} | Get instances of a workload



## ApiV0InstancesCreatePost

> ApiV0InstancesCreatePost(ctx).InstanceDefinition(instanceDefinition).Execute()





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
    r, err := apiClient.InstancesApi.ApiV0InstancesCreatePost(context.Background()).InstanceDefinition(instanceDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ApiV0InstancesCreatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV0InstancesCreatePostRequest struct via the builder pattern


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


## ApiV0InstancesDeletePost

> ApiV0InstancesDeletePost(ctx).ApiV0InstancesDeletePostRequest(apiV0InstancesDeletePostRequest).Execute()





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
    apiV0InstancesDeletePostRequest := *openapiclient.NewApiV0InstancesDeletePostRequest() // ApiV0InstancesDeletePostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InstancesApi.ApiV0InstancesDeletePost(context.Background()).ApiV0InstancesDeletePostRequest(apiV0InstancesDeletePostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ApiV0InstancesDeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV0InstancesDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV0InstancesDeletePostRequest** | [**ApiV0InstancesDeletePostRequest**](ApiV0InstancesDeletePostRequest.md) |  | 

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


## ApiV0InstancesListGet

> ApiV0WorkloadsInstancesWorkloadIdGet200Response ApiV0InstancesListGet(ctx).Execute()





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
    resp, r, err := apiClient.InstancesApi.ApiV0InstancesListGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ApiV0InstancesListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV0InstancesListGet`: ApiV0WorkloadsInstancesWorkloadIdGet200Response
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ApiV0InstancesListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV0InstancesListGetRequest struct via the builder pattern


### Return type

[**ApiV0WorkloadsInstancesWorkloadIdGet200Response**](ApiV0WorkloadsInstancesWorkloadIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV0WorkloadsInstancesWorkloadIdGet

> ApiV0WorkloadsInstancesWorkloadIdGet200Response ApiV0WorkloadsInstancesWorkloadIdGet(ctx, workloadId).Execute()

Get instances of a workload



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
    workloadId := "28dcac69-33ef-4b13-a42f-0d07c7acc1a6" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstancesApi.ApiV0WorkloadsInstancesWorkloadIdGet(context.Background(), workloadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ApiV0WorkloadsInstancesWorkloadIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV0WorkloadsInstancesWorkloadIdGet`: ApiV0WorkloadsInstancesWorkloadIdGet200Response
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ApiV0WorkloadsInstancesWorkloadIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workloadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV0WorkloadsInstancesWorkloadIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiV0WorkloadsInstancesWorkloadIdGet200Response**](ApiV0WorkloadsInstancesWorkloadIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

