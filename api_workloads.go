/*
RIK

RESTful public-facing API. The API is accessible through HTTP calls on specific URLs carrying JSON modeled data. 

API version: 0.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type WorkloadsApi interface {

	/*
	CreateWorkload Create a new workload

	Create a new workload

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return WorkloadsApiCreateWorkloadRequest
	*/
	CreateWorkload(ctx context.Context) WorkloadsApiCreateWorkloadRequest

	// CreateWorkloadExecute executes the request
	//  @return CreateWorkload200Response
	CreateWorkloadExecute(r WorkloadsApiCreateWorkloadRequest) (*CreateWorkload200Response, *http.Response, error)

	/*
	DeleteWorkload Method for DeleteWorkload

	Delete a workload

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return WorkloadsApiDeleteWorkloadRequest
	*/
	DeleteWorkload(ctx context.Context) WorkloadsApiDeleteWorkloadRequest

	// DeleteWorkloadExecute executes the request
	DeleteWorkloadExecute(r WorkloadsApiDeleteWorkloadRequest) (*http.Response, error)

	/*
	GetWorkloads List all the workloads available in the cluster

	Retrieve a list of all the active workloads in the cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return WorkloadsApiGetWorkloadsRequest
	*/
	GetWorkloads(ctx context.Context) WorkloadsApiGetWorkloadsRequest

	// GetWorkloadsExecute executes the request
	//  @return []GetWorkloadsResponseInner
	GetWorkloadsExecute(r WorkloadsApiGetWorkloadsRequest) ([]GetWorkloadsResponseInner, *http.Response, error)
}

// WorkloadsApiService WorkloadsApi service
type WorkloadsApiService service

type WorkloadsApiCreateWorkloadRequest struct {
	ctx context.Context
	ApiService WorkloadsApi
	body *Workload
}

func (r WorkloadsApiCreateWorkloadRequest) Body(body Workload) WorkloadsApiCreateWorkloadRequest {
	r.body = &body
	return r
}

func (r WorkloadsApiCreateWorkloadRequest) Execute() (*CreateWorkload200Response, *http.Response, error) {
	return r.ApiService.CreateWorkloadExecute(r)
}

/*
CreateWorkload Create a new workload

Create a new workload

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return WorkloadsApiCreateWorkloadRequest
*/
func (a *WorkloadsApiService) CreateWorkload(ctx context.Context) WorkloadsApiCreateWorkloadRequest {
	return WorkloadsApiCreateWorkloadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateWorkload200Response
func (a *WorkloadsApiService) CreateWorkloadExecute(r WorkloadsApiCreateWorkloadRequest) (*CreateWorkload200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateWorkload200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkloadsApiService.CreateWorkload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/workloads.create"

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
	localVarPostBody = r.body
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

type WorkloadsApiDeleteWorkloadRequest struct {
	ctx context.Context
	ApiService WorkloadsApi
	metadata *Metadata
}

func (r WorkloadsApiDeleteWorkloadRequest) Metadata(metadata Metadata) WorkloadsApiDeleteWorkloadRequest {
	r.metadata = &metadata
	return r
}

func (r WorkloadsApiDeleteWorkloadRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteWorkloadExecute(r)
}

/*
DeleteWorkload Method for DeleteWorkload

Delete a workload

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return WorkloadsApiDeleteWorkloadRequest
*/
func (a *WorkloadsApiService) DeleteWorkload(ctx context.Context) WorkloadsApiDeleteWorkloadRequest {
	return WorkloadsApiDeleteWorkloadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WorkloadsApiService) DeleteWorkloadExecute(r WorkloadsApiDeleteWorkloadRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkloadsApiService.DeleteWorkload")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/workloads.delete"

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
	localVarPostBody = r.metadata
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

type WorkloadsApiGetWorkloadsRequest struct {
	ctx context.Context
	ApiService WorkloadsApi
}

func (r WorkloadsApiGetWorkloadsRequest) Execute() ([]GetWorkloadsResponseInner, *http.Response, error) {
	return r.ApiService.GetWorkloadsExecute(r)
}

/*
GetWorkloads List all the workloads available in the cluster

Retrieve a list of all the active workloads in the cluster.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return WorkloadsApiGetWorkloadsRequest
*/
func (a *WorkloadsApiService) GetWorkloads(ctx context.Context) WorkloadsApiGetWorkloadsRequest {
	return WorkloadsApiGetWorkloadsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []GetWorkloadsResponseInner
func (a *WorkloadsApiService) GetWorkloadsExecute(r WorkloadsApiGetWorkloadsRequest) ([]GetWorkloadsResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetWorkloadsResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkloadsApiService.GetWorkloads")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/workloads.list"

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
