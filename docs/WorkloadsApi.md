# \WorkloadsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV0WorkloadsCreatePost**](WorkloadsApi.md#ApiV0WorkloadsCreatePost) | **Post** /api/v0/workloads.create | 
[**ApiV0WorkloadsDeletePost**](WorkloadsApi.md#ApiV0WorkloadsDeletePost) | **Post** /api/v0/workloads.delete | 
[**ApiV0WorkloadsListGet**](WorkloadsApi.md#ApiV0WorkloadsListGet) | **Get** /api/v0/workloads.list | 



## ApiV0WorkloadsCreatePost

> WorkloadName ApiV0WorkloadsCreatePost(ctx).WorkloadDefinition(workloadDefinition).Execute()





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
    workloadDefinition := *openapiclient.NewWorkloadDefinition() // WorkloadDefinition |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkloadsApi.ApiV0WorkloadsCreatePost(context.Background()).WorkloadDefinition(workloadDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsApi.ApiV0WorkloadsCreatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV0WorkloadsCreatePost`: WorkloadName
    fmt.Fprintf(os.Stdout, "Response from `WorkloadsApi.ApiV0WorkloadsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV0WorkloadsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workloadDefinition** | [**WorkloadDefinition**](WorkloadDefinition.md) |  | 

### Return type

[**WorkloadName**](WorkloadName.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV0WorkloadsDeletePost

> ApiV0WorkloadsDeletePost(ctx).WorkloadName(workloadName).Execute()





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
    workloadName := *openapiclient.NewWorkloadName() // WorkloadName |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkloadsApi.ApiV0WorkloadsDeletePost(context.Background()).WorkloadName(workloadName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsApi.ApiV0WorkloadsDeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV0WorkloadsDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workloadName** | [**WorkloadName**](WorkloadName.md) |  | 

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


## ApiV0WorkloadsListGet

> []Workload ApiV0WorkloadsListGet(ctx).Execute()



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
    resp, r, err := apiClient.WorkloadsApi.ApiV0WorkloadsListGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsApi.ApiV0WorkloadsListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV0WorkloadsListGet`: []Workload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadsApi.ApiV0WorkloadsListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV0WorkloadsListGetRequest struct via the builder pattern


### Return type

[**[]Workload**](Workload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

