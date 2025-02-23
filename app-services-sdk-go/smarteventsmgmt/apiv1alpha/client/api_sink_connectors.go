/*
Red Hat Openshift SmartEvents Fleet Manager V2

The API exposed by the fleet manager of the SmartEvents service.

API version: 0.0.1
Contact: openbridge-dev@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smarteventsmgmtclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


type SinkConnectorsApi interface {

	/*
	SinkConnectorsAPICreateSinkConnector Create a Sink Connector for a Bridge instance

	Create a Sink Connector of a Bridge instance for the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bridgeId
	@return ApiSinkConnectorsAPICreateSinkConnectorRequest
	*/
	SinkConnectorsAPICreateSinkConnector(ctx context.Context, bridgeId string) ApiSinkConnectorsAPICreateSinkConnectorRequest

	// SinkConnectorsAPICreateSinkConnectorExecute executes the request
	//  @return SinkConnectorResponse
	SinkConnectorsAPICreateSinkConnectorExecute(r ApiSinkConnectorsAPICreateSinkConnectorRequest) (*SinkConnectorResponse, *http.Response, error)

	/*
	SinkConnectorsAPIDeleteSinkConnector Delete a Sink Connector of a Bridge instance

	Delete a Sink Connector of a Bridge instance for the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bridgeId
	@param sinkId
	@return ApiSinkConnectorsAPIDeleteSinkConnectorRequest
	*/
	SinkConnectorsAPIDeleteSinkConnector(ctx context.Context, bridgeId string, sinkId string) ApiSinkConnectorsAPIDeleteSinkConnectorRequest

	// SinkConnectorsAPIDeleteSinkConnectorExecute executes the request
	SinkConnectorsAPIDeleteSinkConnectorExecute(r ApiSinkConnectorsAPIDeleteSinkConnectorRequest) (*http.Response, error)

	/*
	SinkConnectorsAPIGetSinkConnector Get a Sink Connector of a Bridge instance

	Get a Sink Connector of a Bridge instance for the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bridgeId
	@param sinkId
	@return ApiSinkConnectorsAPIGetSinkConnectorRequest
	*/
	SinkConnectorsAPIGetSinkConnector(ctx context.Context, bridgeId string, sinkId string) ApiSinkConnectorsAPIGetSinkConnectorRequest

	// SinkConnectorsAPIGetSinkConnectorExecute executes the request
	//  @return SinkConnectorResponse
	SinkConnectorsAPIGetSinkConnectorExecute(r ApiSinkConnectorsAPIGetSinkConnectorRequest) (*SinkConnectorResponse, *http.Response, error)

	/*
	SinkConnectorsAPIGetSinkConnectors Get the list of Sink Connectors for a Bridge

	Get the list of Sink Connector instances of a Bridge instance for the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bridgeId
	@return ApiSinkConnectorsAPIGetSinkConnectorsRequest
	*/
	SinkConnectorsAPIGetSinkConnectors(ctx context.Context, bridgeId string) ApiSinkConnectorsAPIGetSinkConnectorsRequest

	// SinkConnectorsAPIGetSinkConnectorsExecute executes the request
	//  @return SinkConnectorListResponse
	SinkConnectorsAPIGetSinkConnectorsExecute(r ApiSinkConnectorsAPIGetSinkConnectorsRequest) (*SinkConnectorListResponse, *http.Response, error)

	/*
	SinkConnectorsAPIUpdateSinkConnector Update a Sink Connector instance.

	Update a Sink Connector instance for the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bridgeId
	@param sinkId
	@return ApiSinkConnectorsAPIUpdateSinkConnectorRequest
	*/
	SinkConnectorsAPIUpdateSinkConnector(ctx context.Context, bridgeId string, sinkId string) ApiSinkConnectorsAPIUpdateSinkConnectorRequest

	// SinkConnectorsAPIUpdateSinkConnectorExecute executes the request
	//  @return SinkConnectorResponse
	SinkConnectorsAPIUpdateSinkConnectorExecute(r ApiSinkConnectorsAPIUpdateSinkConnectorRequest) (*SinkConnectorResponse, *http.Response, error)
}

// SinkConnectorsApiService SinkConnectorsApi service
type SinkConnectorsApiService service

type ApiSinkConnectorsAPICreateSinkConnectorRequest struct {
	ctx context.Context
	ApiService SinkConnectorsApi
	bridgeId string
	connectorRequest *ConnectorRequest
}

func (r ApiSinkConnectorsAPICreateSinkConnectorRequest) ConnectorRequest(connectorRequest ConnectorRequest) ApiSinkConnectorsAPICreateSinkConnectorRequest {
	r.connectorRequest = &connectorRequest
	return r
}

func (r ApiSinkConnectorsAPICreateSinkConnectorRequest) Execute() (*SinkConnectorResponse, *http.Response, error) {
	return r.ApiService.SinkConnectorsAPICreateSinkConnectorExecute(r)
}

/*
SinkConnectorsAPICreateSinkConnector Create a Sink Connector for a Bridge instance

Create a Sink Connector of a Bridge instance for the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bridgeId
 @return ApiSinkConnectorsAPICreateSinkConnectorRequest
*/
func (a *SinkConnectorsApiService) SinkConnectorsAPICreateSinkConnector(ctx context.Context, bridgeId string) ApiSinkConnectorsAPICreateSinkConnectorRequest {
	return ApiSinkConnectorsAPICreateSinkConnectorRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
	}
}

// Execute executes the request
//  @return SinkConnectorResponse
func (a *SinkConnectorsApiService) SinkConnectorsAPICreateSinkConnectorExecute(r ApiSinkConnectorsAPICreateSinkConnectorRequest) (*SinkConnectorResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SinkConnectorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SinkConnectorsApiService.SinkConnectorsAPICreateSinkConnector")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", url.PathEscape(parameterToString(r.bridgeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.bridgeId) < 1 {
		return localVarReturnValue, nil, reportError("bridgeId must have at least 1 elements")
	}

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
	localVarPostBody = r.connectorRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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

type ApiSinkConnectorsAPIDeleteSinkConnectorRequest struct {
	ctx context.Context
	ApiService SinkConnectorsApi
	bridgeId string
	sinkId string
}

func (r ApiSinkConnectorsAPIDeleteSinkConnectorRequest) Execute() (*http.Response, error) {
	return r.ApiService.SinkConnectorsAPIDeleteSinkConnectorExecute(r)
}

/*
SinkConnectorsAPIDeleteSinkConnector Delete a Sink Connector of a Bridge instance

Delete a Sink Connector of a Bridge instance for the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bridgeId
 @param sinkId
 @return ApiSinkConnectorsAPIDeleteSinkConnectorRequest
*/
func (a *SinkConnectorsApiService) SinkConnectorsAPIDeleteSinkConnector(ctx context.Context, bridgeId string, sinkId string) ApiSinkConnectorsAPIDeleteSinkConnectorRequest {
	return ApiSinkConnectorsAPIDeleteSinkConnectorRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
		sinkId: sinkId,
	}
}

// Execute executes the request
func (a *SinkConnectorsApiService) SinkConnectorsAPIDeleteSinkConnectorExecute(r ApiSinkConnectorsAPIDeleteSinkConnectorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SinkConnectorsApiService.SinkConnectorsAPIDeleteSinkConnector")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks/{sinkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", url.PathEscape(parameterToString(r.bridgeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sinkId"+"}", url.PathEscape(parameterToString(r.sinkId, "")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSinkConnectorsAPIGetSinkConnectorRequest struct {
	ctx context.Context
	ApiService SinkConnectorsApi
	bridgeId string
	sinkId string
}

func (r ApiSinkConnectorsAPIGetSinkConnectorRequest) Execute() (*SinkConnectorResponse, *http.Response, error) {
	return r.ApiService.SinkConnectorsAPIGetSinkConnectorExecute(r)
}

/*
SinkConnectorsAPIGetSinkConnector Get a Sink Connector of a Bridge instance

Get a Sink Connector of a Bridge instance for the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bridgeId
 @param sinkId
 @return ApiSinkConnectorsAPIGetSinkConnectorRequest
*/
func (a *SinkConnectorsApiService) SinkConnectorsAPIGetSinkConnector(ctx context.Context, bridgeId string, sinkId string) ApiSinkConnectorsAPIGetSinkConnectorRequest {
	return ApiSinkConnectorsAPIGetSinkConnectorRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
		sinkId: sinkId,
	}
}

// Execute executes the request
//  @return SinkConnectorResponse
func (a *SinkConnectorsApiService) SinkConnectorsAPIGetSinkConnectorExecute(r ApiSinkConnectorsAPIGetSinkConnectorRequest) (*SinkConnectorResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SinkConnectorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SinkConnectorsApiService.SinkConnectorsAPIGetSinkConnector")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks/{sinkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", url.PathEscape(parameterToString(r.bridgeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sinkId"+"}", url.PathEscape(parameterToString(r.sinkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.bridgeId) < 1 {
		return localVarReturnValue, nil, reportError("bridgeId must have at least 1 elements")
	}
	if strlen(r.sinkId) < 1 {
		return localVarReturnValue, nil, reportError("sinkId must have at least 1 elements")
	}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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

type ApiSinkConnectorsAPIGetSinkConnectorsRequest struct {
	ctx context.Context
	ApiService SinkConnectorsApi
	bridgeId string
	name *string
	page *int32
	size *int32
	status *[]ManagedResourceStatus
}

func (r ApiSinkConnectorsAPIGetSinkConnectorsRequest) Name(name string) ApiSinkConnectorsAPIGetSinkConnectorsRequest {
	r.name = &name
	return r
}

func (r ApiSinkConnectorsAPIGetSinkConnectorsRequest) Page(page int32) ApiSinkConnectorsAPIGetSinkConnectorsRequest {
	r.page = &page
	return r
}

func (r ApiSinkConnectorsAPIGetSinkConnectorsRequest) Size(size int32) ApiSinkConnectorsAPIGetSinkConnectorsRequest {
	r.size = &size
	return r
}

func (r ApiSinkConnectorsAPIGetSinkConnectorsRequest) Status(status []ManagedResourceStatus) ApiSinkConnectorsAPIGetSinkConnectorsRequest {
	r.status = &status
	return r
}

func (r ApiSinkConnectorsAPIGetSinkConnectorsRequest) Execute() (*SinkConnectorListResponse, *http.Response, error) {
	return r.ApiService.SinkConnectorsAPIGetSinkConnectorsExecute(r)
}

/*
SinkConnectorsAPIGetSinkConnectors Get the list of Sink Connectors for a Bridge

Get the list of Sink Connector instances of a Bridge instance for the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bridgeId
 @return ApiSinkConnectorsAPIGetSinkConnectorsRequest
*/
func (a *SinkConnectorsApiService) SinkConnectorsAPIGetSinkConnectors(ctx context.Context, bridgeId string) ApiSinkConnectorsAPIGetSinkConnectorsRequest {
	return ApiSinkConnectorsAPIGetSinkConnectorsRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
	}
}

// Execute executes the request
//  @return SinkConnectorListResponse
func (a *SinkConnectorsApiService) SinkConnectorsAPIGetSinkConnectorsExecute(r ApiSinkConnectorsAPIGetSinkConnectorsRequest) (*SinkConnectorListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SinkConnectorListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SinkConnectorsApiService.SinkConnectorsAPIGetSinkConnectors")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", url.PathEscape(parameterToString(r.bridgeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.bridgeId) < 1 {
		return localVarReturnValue, nil, reportError("bridgeId must have at least 1 elements")
	}

	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.size != nil {
		localVarQueryParams.Add("size", parameterToString(*r.size, ""))
	}
	if r.status != nil {
		t := *r.status
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("status", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("status", parameterToString(t, "multi"))
		}
	}
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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

type ApiSinkConnectorsAPIUpdateSinkConnectorRequest struct {
	ctx context.Context
	ApiService SinkConnectorsApi
	bridgeId string
	sinkId string
	connectorRequest *ConnectorRequest
}

func (r ApiSinkConnectorsAPIUpdateSinkConnectorRequest) ConnectorRequest(connectorRequest ConnectorRequest) ApiSinkConnectorsAPIUpdateSinkConnectorRequest {
	r.connectorRequest = &connectorRequest
	return r
}

func (r ApiSinkConnectorsAPIUpdateSinkConnectorRequest) Execute() (*SinkConnectorResponse, *http.Response, error) {
	return r.ApiService.SinkConnectorsAPIUpdateSinkConnectorExecute(r)
}

/*
SinkConnectorsAPIUpdateSinkConnector Update a Sink Connector instance.

Update a Sink Connector instance for the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bridgeId
 @param sinkId
 @return ApiSinkConnectorsAPIUpdateSinkConnectorRequest
*/
func (a *SinkConnectorsApiService) SinkConnectorsAPIUpdateSinkConnector(ctx context.Context, bridgeId string, sinkId string) ApiSinkConnectorsAPIUpdateSinkConnectorRequest {
	return ApiSinkConnectorsAPIUpdateSinkConnectorRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
		sinkId: sinkId,
	}
}

// Execute executes the request
//  @return SinkConnectorResponse
func (a *SinkConnectorsApiService) SinkConnectorsAPIUpdateSinkConnectorExecute(r ApiSinkConnectorsAPIUpdateSinkConnectorRequest) (*SinkConnectorResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SinkConnectorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SinkConnectorsApiService.SinkConnectorsAPIUpdateSinkConnector")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks/{sinkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", url.PathEscape(parameterToString(r.bridgeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sinkId"+"}", url.PathEscape(parameterToString(r.sinkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.bridgeId) < 1 {
		return localVarReturnValue, nil, reportError("bridgeId must have at least 1 elements")
	}
	if strlen(r.sinkId) < 1 {
		return localVarReturnValue, nil, reportError("sinkId must have at least 1 elements")
	}

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
	localVarPostBody = r.connectorRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorsList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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
