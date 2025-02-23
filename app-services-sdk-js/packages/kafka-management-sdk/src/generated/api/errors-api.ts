/* tslint:disable */
/* eslint-disable */
/**
 * Kafka Management API
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * The version of the OpenAPI document: 1.14.0
 * Contact: rhosak-support@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import globalAxios, { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from '../common';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { ErrorList } from '../model';
/**
 * ErrorsApi - axios parameter creator
 * @export
 */
export const ErrorsApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * Returns the error by Id
         * @param {string} id The ID of record
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getErrorById: async (id: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('getErrorById', 'id', id)
            const localVarPath = `/api/kafkas_mgmt/v1/errors/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * Returns the list of possible API errors
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getErrors: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/kafkas_mgmt/v1/errors`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * ErrorsApi - functional programming interface
 * @export
 */
export const ErrorsApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = ErrorsApiAxiosParamCreator(configuration)
    return {
        /**
         * Returns the error by Id
         * @param {string} id The ID of record
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getErrorById(id: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Error>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getErrorById(id, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * Returns the list of possible API errors
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getErrors(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ErrorList>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getErrors(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * ErrorsApi - factory interface
 * @export
 */
export const ErrorsApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = ErrorsApiFp(configuration)
    return {
        /**
         * Returns the error by Id
         * @param {string} id The ID of record
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getErrorById(id: string, options?: any): AxiosPromise<Error> {
            return localVarFp.getErrorById(id, options).then((request) => request(axios, basePath));
        },
        /**
         * Returns the list of possible API errors
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getErrors(options?: any): AxiosPromise<ErrorList> {
            return localVarFp.getErrors(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * ErrorsApi - interface
 * @export
 * @interface ErrorsApi
 */
export interface ErrorsApiInterface {
    /**
     * Returns the error by Id
     * @param {string} id The ID of record
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ErrorsApiInterface
     */
    getErrorById(id: string, options?: AxiosRequestConfig): AxiosPromise<Error>;

    /**
     * Returns the list of possible API errors
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ErrorsApiInterface
     */
    getErrors(options?: AxiosRequestConfig): AxiosPromise<ErrorList>;

}

/**
 * ErrorsApi - object-oriented interface
 * @export
 * @class ErrorsApi
 * @extends {BaseAPI}
 */
export class ErrorsApi extends BaseAPI implements ErrorsApiInterface {
    /**
     * Returns the error by Id
     * @param {string} id The ID of record
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ErrorsApi
     */
    public getErrorById(id: string, options?: AxiosRequestConfig) {
        return ErrorsApiFp(this.configuration).getErrorById(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * Returns the list of possible API errors
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ErrorsApi
     */
    public getErrors(options?: AxiosRequestConfig) {
        return ErrorsApiFp(this.configuration).getErrors(options).then((request) => request(this.axios, this.basePath));
    }
}
