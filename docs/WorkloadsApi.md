# \WorkloadsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkload**](WorkloadsApi.md#CreateWorkload) | **Post** /api/v0/workloads.create | Create a new workload
[**DeleteWorkload**](WorkloadsApi.md#DeleteWorkload) | **Post** /api/v0/workloads.delete | 
[**GetWorkloadInstances**](WorkloadsApi.md#GetWorkloadInstances) | **Get** /api/v0/workloads.instances/{workloadId} | Get all the instances for the given workload
[**GetWorkloads**](WorkloadsApi.md#GetWorkloads) | **Get** /api/v0/workloads.list | List all the workloads available in the cluster



## CreateWorkload

> WorkloadName CreateWorkload(ctx).Workload(workload).Execute()

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
    workload := *openapiclient.NewWorkload() // Workload |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkloadsApi.CreateWorkload(context.Background()).Workload(workload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsApi.CreateWorkload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkload`: WorkloadName
    fmt.Fprintf(os.Stdout, "Response from `WorkloadsApi.CreateWorkload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workload** | [**Workload**](Workload.md) |  | 

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


## DeleteWorkload

> DeleteWorkload(ctx).WorkloadName(workloadName).Execute()





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
    r, err := apiClient.WorkloadsApi.DeleteWorkload(context.Background()).WorkloadName(workloadName).Execute()
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


## GetWorkloadInstances

> GetWorkloadInstances200Response GetWorkloadInstances(ctx, workloadId).Execute()

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
    workloadId := "28dcac69-33ef-4b13-a42f-0d07c7acc1a6" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkloadsApi.GetWorkloadInstances(context.Background(), workloadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsApi.GetWorkloadInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkloadInstances`: GetWorkloadInstances200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkloadsApi.GetWorkloadInstances`: %v\n", resp)
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

[**GetWorkloadInstances200Response**](GetWorkloadInstances200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkloads

> []WorkloadMeta GetWorkloads(ctx).Execute()

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
    // response from `GetWorkloads`: []WorkloadMeta
    fmt.Fprintf(os.Stdout, "Response from `WorkloadsApi.GetWorkloads`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadsRequest struct via the builder pattern


### Return type

[**[]WorkloadMeta**](WorkloadMeta.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

