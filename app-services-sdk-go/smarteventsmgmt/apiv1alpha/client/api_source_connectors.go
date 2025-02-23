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


type SourceConnectorsApi interface {

	/*
	SourceConnectorsAPICreateSourceConnector Create a Source Connector for a Bridge instance

	Create a Source Connector of a Bridge instance for the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bridgeId
	@return ApiSourceConnectorsAPICreateSourceConnectorRequest
	*/
	SourceConnectorsAPICreateSourceConnector(ctx context.Context, bridgeId string) ApiSourceConnectorsAPICreateSourceConnectorRequest

	// SourceConnectorsAPICreateSourceConnectorExecute executes the request
	//  @return SourceConnectorResponse
	SourceConnectorsAPICreateSourceConnectorExecute(r ApiSourceConnectorsAPICreateSourceConnectorRequest) (*SourceConnectorResponse, *http.Response, error)

	/*
	SourceConnectorsAPIDeleteSourceConnector Delete a Source Connector of a Bridge instance

	Delete a Source Connector of a Bridge instance for the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bridgeId
	@param sourceId
	@return ApiSourceConnectorsAPIDeleteSourceConnectorRequest
	*/
	SourceConnectorsAPIDeleteSourceConnector(ctx context.Context, bridgeId string, sourceId string) ApiSourceConnectorsAPIDeleteSourceConnectorRequest

	// SourceConnectorsAPIDeleteSourceConnectorExecute executes the request
	SourceConnectorsAPIDeleteSourceConnectorExecute(r ApiSourceConnectorsAPIDeleteSourceConnectorRequest) (*http.Response, error)

	/*
	SourceConnectorsAPIGetSourceConnector Get a Source Connector of a Bridge instance

	Get a Source Connector of a Bridge instance for the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bridgeId
	@param sourceId
	@return ApiSourceConnectorsAPIGetSourceConnectorRequest
	*/
	SourceConnectorsAPIGetSourceConnector(ctx context.Context, bridgeId string, sourceId string) ApiSourceConnectorsAPIGetSourceConnectorRequest

	// SourceConnectorsAPIGetSourceConnectorExecute executes the request
	//  @return SourceConnectorResponse
	SourceConnectorsAPIGetSourceConnectorExecute(r ApiSourceConnectorsAPIGetSourceConnectorRequest) (*SourceConnectorResponse, *http.Response, error)

	/*
	SourceConnectorsAPIGetSourceConnectors Get the list of Sink Connectors for a Bridge

	Get the list of Source Connector instances of a Bridge instance instance for the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bridgeId
	@return ApiSourceConnectorsAPIGetSourceConnectorsRequest
	*/
	SourceConnectorsAPIGetSourceConnectors(ctx context.Context, bridgeId string) ApiSourceConnectorsAPIGetSourceConnectorsRequest

	// SourceConnectorsAPIGetSourceConnectorsExecute executes the request
	//  @return SourceConnectorListResponse
	SourceConnectorsAPIGetSourceConnectorsExecute(r ApiSourceConnectorsAPIGetSourceConnectorsRequest) (*SourceConnectorListResponse, *http.Response, error)

	/*
	SourceConnectorsAPIUpdateSourceConnector Update a Source Connector instance.

	Update a Source Connector instance for the authenticated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bridgeId
	@param sourceId
	@return ApiSourceConnectorsAPIUpdateSourceConnectorRequest
	*/
	SourceConnectorsAPIUpdateSourceConnector(ctx context.Context, bridgeId string, sourceId string) ApiSourceConnectorsAPIUpdateSourceConnectorRequest

	// SourceConnectorsAPIUpdateSourceConnectorExecute executes the request
	//  @return SourceConnectorResponse
	SourceConnectorsAPIUpdateSourceConnectorExecute(r ApiSourceConnectorsAPIUpdateSourceConnectorRequest) (*SourceConnectorResponse, *http.Response, error)
}

// SourceConnectorsApiService SourceConnectorsApi service
type SourceConnectorsApiService service

type ApiSourceConnectorsAPICreateSourceConnectorRequest struct {
	ctx context.Context
	ApiService SourceConnectorsApi
	bridgeId string
	connectorRequest *ConnectorRequest
}

func (r ApiSourceConnectorsAPICreateSourceConnectorRequest) ConnectorRequest(connectorRequest ConnectorRequest) ApiSourceConnectorsAPICreateSourceConnectorRequest {
	r.connectorRequest = &connectorRequest
	return r
}

func (r ApiSourceConnectorsAPICreateSourceConnectorRequest) Execute() (*SourceConnectorResponse, *http.Response, error) {
	return r.ApiService.SourceConnectorsAPICreateSourceConnectorExecute(r)
}

/*
SourceConnectorsAPICreateSourceConnector Create a Source Connector for a Bridge instance

Create a Source Connector of a Bridge instance for the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bridgeId
 @return ApiSourceConnectorsAPICreateSourceConnectorRequest
*/
func (a *SourceConnectorsApiService) SourceConnectorsAPICreateSourceConnector(ctx context.Context, bridgeId string) ApiSourceConnectorsAPICreateSourceConnectorRequest {
	return ApiSourceConnectorsAPICreateSourceConnectorRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
	}
}

// Execute executes the request
//  @return SourceConnectorResponse
func (a *SourceConnectorsApiService) SourceConnectorsAPICreateSourceConnectorExecute(r ApiSourceConnectorsAPICreateSourceConnectorRequest) (*SourceConnectorResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SourceConnectorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SourceConnectorsApiService.SourceConnectorsAPICreateSourceConnector")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/sources"
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

type ApiSourceConnectorsAPIDeleteSourceConnectorRequest struct {
	ctx context.Context
	ApiService SourceConnectorsApi
	bridgeId string
	sourceId string
}

func (r ApiSourceConnectorsAPIDeleteSourceConnectorRequest) Execute() (*http.Response, error) {
	return r.ApiService.SourceConnectorsAPIDeleteSourceConnectorExecute(r)
}

/*
SourceConnectorsAPIDeleteSourceConnector Delete a Source Connector of a Bridge instance

Delete a Source Connector of a Bridge instance for the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bridgeId
 @param sourceId
 @return ApiSourceConnectorsAPIDeleteSourceConnectorRequest
*/
func (a *SourceConnectorsApiService) SourceConnectorsAPIDeleteSourceConnector(ctx context.Context, bridgeId string, sourceId string) ApiSourceConnectorsAPIDeleteSourceConnectorRequest {
	return ApiSourceConnectorsAPIDeleteSourceConnectorRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
		sourceId: sourceId,
	}
}

// Execute executes the request
func (a *SourceConnectorsApiService) SourceConnectorsAPIDeleteSourceConnectorExecute(r ApiSourceConnectorsAPIDeleteSourceConnectorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SourceConnectorsApiService.SourceConnectorsAPIDeleteSourceConnector")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/sources/{sourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", url.PathEscape(parameterToString(r.bridgeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sourceId"+"}", url.PathEscape(parameterToString(r.sourceId, "")), -1)

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

type ApiSourceConnectorsAPIGetSourceConnectorRequest struct {
	ctx context.Context
	ApiService SourceConnectorsApi
	bridgeId string
	sourceId string
}

func (r ApiSourceConnectorsAPIGetSourceConnectorRequest) Execute() (*SourceConnectorResponse, *http.Response, error) {
	return r.ApiService.SourceConnectorsAPIGetSourceConnectorExecute(r)
}

/*
SourceConnectorsAPIGetSourceConnector Get a Source Connector of a Bridge instance

Get a Source Connector of a Bridge instance for the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bridgeId
 @param sourceId
 @return ApiSourceConnectorsAPIGetSourceConnectorRequest
*/
func (a *SourceConnectorsApiService) SourceConnectorsAPIGetSourceConnector(ctx context.Context, bridgeId string, sourceId string) ApiSourceConnectorsAPIGetSourceConnectorRequest {
	return ApiSourceConnectorsAPIGetSourceConnectorRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
		sourceId: sourceId,
	}
}

// Execute executes the request
//  @return SourceConnectorResponse
func (a *SourceConnectorsApiService) SourceConnectorsAPIGetSourceConnectorExecute(r ApiSourceConnectorsAPIGetSourceConnectorRequest) (*SourceConnectorResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SourceConnectorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SourceConnectorsApiService.SourceConnectorsAPIGetSourceConnector")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/sources/{sourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", url.PathEscape(parameterToString(r.bridgeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sourceId"+"}", url.PathEscape(parameterToString(r.sourceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.bridgeId) < 1 {
		return localVarReturnValue, nil, reportError("bridgeId must have at least 1 elements")
	}
	if strlen(r.sourceId) < 1 {
		return localVarReturnValue, nil, reportError("sourceId must have at least 1 elements")
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

type ApiSourceConnectorsAPIGetSourceConnectorsRequest struct {
	ctx context.Context
	ApiService SourceConnectorsApi
	bridgeId string
	name *string
	page *int32
	size *int32
	status *[]ManagedResourceStatus
}

func (r ApiSourceConnectorsAPIGetSourceConnectorsRequest) Name(name string) ApiSourceConnectorsAPIGetSourceConnectorsRequest {
	r.name = &name
	return r
}

func (r ApiSourceConnectorsAPIGetSourceConnectorsRequest) Page(page int32) ApiSourceConnectorsAPIGetSourceConnectorsRequest {
	r.page = &page
	return r
}

func (r ApiSourceConnectorsAPIGetSourceConnectorsRequest) Size(size int32) ApiSourceConnectorsAPIGetSourceConnectorsRequest {
	r.size = &size
	return r
}

func (r ApiSourceConnectorsAPIGetSourceConnectorsRequest) Status(status []ManagedResourceStatus) ApiSourceConnectorsAPIGetSourceConnectorsRequest {
	r.status = &status
	return r
}

func (r ApiSourceConnectorsAPIGetSourceConnectorsRequest) Execute() (*SourceConnectorListResponse, *http.Response, error) {
	return r.ApiService.SourceConnectorsAPIGetSourceConnectorsExecute(r)
}

/*
SourceConnectorsAPIGetSourceConnectors Get the list of Sink Connectors for a Bridge

Get the list of Source Connector instances of a Bridge instance instance for the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bridgeId
 @return ApiSourceConnectorsAPIGetSourceConnectorsRequest
*/
func (a *SourceConnectorsApiService) SourceConnectorsAPIGetSourceConnectors(ctx context.Context, bridgeId string) ApiSourceConnectorsAPIGetSourceConnectorsRequest {
	return ApiSourceConnectorsAPIGetSourceConnectorsRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
	}
}

// Execute executes the request
//  @return SourceConnectorListResponse
func (a *SourceConnectorsApiService) SourceConnectorsAPIGetSourceConnectorsExecute(r ApiSourceConnectorsAPIGetSourceConnectorsRequest) (*SourceConnectorListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SourceConnectorListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SourceConnectorsApiService.SourceConnectorsAPIGetSourceConnectors")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/sources"
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

type ApiSourceConnectorsAPIUpdateSourceConnectorRequest struct {
	ctx context.Context
	ApiService SourceConnectorsApi
	bridgeId string
	sourceId string
	connectorRequest *ConnectorRequest
}

func (r ApiSourceConnectorsAPIUpdateSourceConnectorRequest) ConnectorRequest(connectorRequest ConnectorRequest) ApiSourceConnectorsAPIUpdateSourceConnectorRequest {
	r.connectorRequest = &connectorRequest
	return r
}

func (r ApiSourceConnectorsAPIUpdateSourceConnectorRequest) Execute() (*SourceConnectorResponse, *http.Response, error) {
	return r.ApiService.SourceConnectorsAPIUpdateSourceConnectorExecute(r)
}

/*
SourceConnectorsAPIUpdateSourceConnector Update a Source Connector instance.

Update a Source Connector instance for the authenticated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bridgeId
 @param sourceId
 @return ApiSourceConnectorsAPIUpdateSourceConnectorRequest
*/
func (a *SourceConnectorsApiService) SourceConnectorsAPIUpdateSourceConnector(ctx context.Context, bridgeId string, sourceId string) ApiSourceConnectorsAPIUpdateSourceConnectorRequest {
	return ApiSourceConnectorsAPIUpdateSourceConnectorRequest{
		ApiService: a,
		ctx: ctx,
		bridgeId: bridgeId,
		sourceId: sourceId,
	}
}

// Execute executes the request
//  @return SourceConnectorResponse
func (a *SourceConnectorsApiService) SourceConnectorsAPIUpdateSourceConnectorExecute(r ApiSourceConnectorsAPIUpdateSourceConnectorRequest) (*SourceConnectorResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SourceConnectorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SourceConnectorsApiService.SourceConnectorsAPIUpdateSourceConnector")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/smartevents_mgmt/v2/bridges/{bridgeId}/sources/{sourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeId"+"}", url.PathEscape(parameterToString(r.bridgeId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sourceId"+"}", url.PathEscape(parameterToString(r.sourceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.bridgeId) < 1 {
		return localVarReturnValue, nil, reportError("bridgeId must have at least 1 elements")
	}
	if strlen(r.sourceId) < 1 {
		return localVarReturnValue, nil, reportError("sourceId must have at least 1 elements")
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
