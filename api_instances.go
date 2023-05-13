/*
RIK

RESTful public-facing API. The API is accessible through HTTP calls on specific URLs carrying JSON modeled data. 

API version: 0.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type InstancesApi interface {

	/*
	CreateWorkloadInstance Create a new instance for a given workload

	Create a new instance for a given workload.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return InstancesApiCreateWorkloadInstanceRequest
	*/
	CreateWorkloadInstance(ctx context.Context) InstancesApiCreateWorkloadInstanceRequest

	// CreateWorkloadInstanceExecute executes the request
	//  @return []string
	CreateWorkloadInstanceExecute(r InstancesApiCreateWorkloadInstanceRequest) ([]string, *http.Response, error)

	/*
	DeleteInstance Method for DeleteInstance

	Delete an instance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return InstancesApiDeleteInstanceRequest
	*/
	DeleteInstance(ctx context.Context) InstancesApiDeleteInstanceRequest

	// DeleteInstanceExecute executes the request
	DeleteInstanceExecute(r InstancesApiDeleteInstanceRequest) (*http.Response, error)

	/*
	GetInstances Method for GetInstances

	List all instances

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return InstancesApiGetInstancesRequest
	*/
	GetInstances(ctx context.Context) InstancesApiGetInstancesRequest

	// GetInstancesExecute executes the request
	//  @return GetInstances200Response
	GetInstancesExecute(r InstancesApiGetInstancesRequest) (*GetInstances200Response, *http.Response, error)

	/*
	GetWorkloadInstances Get all the instances for the given workload

	Retrieve the list of instances for the given workload.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param workloadId
	@return InstancesApiGetWorkloadInstancesRequest
	*/
	GetWorkloadInstances(ctx context.Context, workloadId string) InstancesApiGetWorkloadInstancesRequest

	// GetWorkloadInstancesExecute executes the request
	//  @return GetWorkloadInstancesResponse
	GetWorkloadInstancesExecute(r InstancesApiGetWorkloadInstancesRequest) (*GetWorkloadInstancesResponse, *http.Response, error)
}

// InstancesApiService InstancesApi service
type InstancesApiService service

type InstancesApiCreateWorkloadInstanceRequest struct {
	ctx context.Context
	ApiService InstancesApi
	createInstanceRequest *CreateInstanceRequest
}

func (r InstancesApiCreateWorkloadInstanceRequest) CreateInstanceRequest(createInstanceRequest CreateInstanceRequest) InstancesApiCreateWorkloadInstanceRequest {
	r.createInstanceRequest = &createInstanceRequest
	return r
}

func (r InstancesApiCreateWorkloadInstanceRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.CreateWorkloadInstanceExecute(r)
}

/*
CreateWorkloadInstance Create a new instance for a given workload

Create a new instance for a given workload.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return InstancesApiCreateWorkloadInstanceRequest
*/
func (a *InstancesApiService) CreateWorkloadInstance(ctx context.Context) InstancesApiCreateWorkloadInstanceRequest {
	return InstancesApiCreateWorkloadInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []string
func (a *InstancesApiService) CreateWorkloadInstanceExecute(r InstancesApiCreateWorkloadInstanceRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.CreateWorkloadInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/instances.create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createInstanceRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type InstancesApiDeleteInstanceRequest struct {
	ctx context.Context
	ApiService InstancesApi
	deleteInstanceRequest *DeleteInstanceRequest
}

func (r InstancesApiDeleteInstanceRequest) DeleteInstanceRequest(deleteInstanceRequest DeleteInstanceRequest) InstancesApiDeleteInstanceRequest {
	r.deleteInstanceRequest = &deleteInstanceRequest
	return r
}

func (r InstancesApiDeleteInstanceRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteInstanceExecute(r)
}

/*
DeleteInstance Method for DeleteInstance

Delete an instance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return InstancesApiDeleteInstanceRequest
*/
func (a *InstancesApiService) DeleteInstance(ctx context.Context) InstancesApiDeleteInstanceRequest {
	return InstancesApiDeleteInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *InstancesApiService) DeleteInstanceExecute(r InstancesApiDeleteInstanceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.DeleteInstance")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/instances.delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.deleteInstanceRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type InstancesApiGetInstancesRequest struct {
	ctx context.Context
	ApiService InstancesApi
}

func (r InstancesApiGetInstancesRequest) Execute() (*GetInstances200Response, *http.Response, error) {
	return r.ApiService.GetInstancesExecute(r)
}

/*
GetInstances Method for GetInstances

List all instances

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return InstancesApiGetInstancesRequest
*/
func (a *InstancesApiService) GetInstances(ctx context.Context) InstancesApiGetInstancesRequest {
	return InstancesApiGetInstancesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetInstances200Response
func (a *InstancesApiService) GetInstancesExecute(r InstancesApiGetInstancesRequest) (*GetInstances200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetInstances200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.GetInstances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/instances.list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type InstancesApiGetWorkloadInstancesRequest struct {
	ctx context.Context
	ApiService InstancesApi
	workloadId string
}

func (r InstancesApiGetWorkloadInstancesRequest) Execute() (*GetWorkloadInstancesResponse, *http.Response, error) {
	return r.ApiService.GetWorkloadInstancesExecute(r)
}

/*
GetWorkloadInstances Get all the instances for the given workload

Retrieve the list of instances for the given workload.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workloadId
 @return InstancesApiGetWorkloadInstancesRequest
*/
func (a *InstancesApiService) GetWorkloadInstances(ctx context.Context, workloadId string) InstancesApiGetWorkloadInstancesRequest {
	return InstancesApiGetWorkloadInstancesRequest{
		ApiService: a,
		ctx: ctx,
		workloadId: workloadId,
	}
}

// Execute executes the request
//  @return GetWorkloadInstancesResponse
func (a *InstancesApiService) GetWorkloadInstancesExecute(r InstancesApiGetWorkloadInstancesRequest) (*GetWorkloadInstancesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetWorkloadInstancesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.GetWorkloadInstances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/workloads.instances/{workloadId}"
	localVarPath = strings.Replace(localVarPath, "{"+"workloadId"+"}", url.PathEscape(parameterValueToString(r.workloadId, "workloadId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
