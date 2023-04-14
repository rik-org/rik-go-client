# \WorkloadsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkload**](WorkloadsApi.md#CreateWorkload) | **Post** /api/v0/workloads.create | Create a new workload
[**DeleteWorkload**](WorkloadsApi.md#DeleteWorkload) | **Post** /api/v0/workloads.delete | 
[**GetWorkloads**](WorkloadsApi.md#GetWorkloads) | **Get** /api/v0/workloads.list | List all the workloads available in the cluster



## CreateWorkload

> CreateWorkload200Response CreateWorkload(ctx).Body(body).Execute()

Create a new workload



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
    body := Workload(987) // Workload |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkloadsApi.CreateWorkload(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsApi.CreateWorkload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkload`: CreateWorkload200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkloadsApi.CreateWorkload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **Workload** |  | 

### Return type

[**CreateWorkload200Response**](CreateWorkload200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkload

> DeleteWorkload(ctx).Metadata(metadata).Execute()





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
    metadata := *openapiclient.NewMetadata() // Metadata |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkloadsApi.DeleteWorkload(context.Background()).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsApi.DeleteWorkload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metadata** | [**Metadata**](Metadata.md) |  | 

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


## GetWorkloads

> []GetWorkloadsResponseInner GetWorkloads(ctx).Execute()

List all the workloads available in the cluster



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
    resp, r, err := apiClient.WorkloadsApi.GetWorkloads(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsApi.GetWorkloads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkloads`: []GetWorkloadsResponseInner
    fmt.Fprintf(os.Stdout, "Response from `WorkloadsApi.GetWorkloads`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadsRequest struct via the builder pattern


### Return type

[**[]GetWorkloadsResponseInner**](GetWorkloadsResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

