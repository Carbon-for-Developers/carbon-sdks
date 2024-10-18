/* tslint:disable */
/* eslint-disable */
/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/

import globalAxios, { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction, isBrowser } from '../common';
import { fromBuffer } from "file-type/browser"
const FormData = require("form-data")
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { Credentials } from '../models';
// @ts-ignore
import { DeleteWhiteLabelRequest } from '../models';
// @ts-ignore
import { HTTPValidationError } from '../models';
// @ts-ignore
import { ListWhiteLabelRequest } from '../models';
// @ts-ignore
import { OrderDir } from '../models';
// @ts-ignore
import { Pagination } from '../models';
// @ts-ignore
import { WhiteLabelCreateRequestInner } from '../models';
// @ts-ignore
import { WhiteLabelFilters } from '../models';
// @ts-ignore
import { WhiteLabelOrderByColumns } from '../models';
// @ts-ignore
import { WhiteLabelUpdateRequest } from '../models';
import { paginate } from "../pagination/paginate";
import type * as buffer from "buffer"
import { requestBeforeHook } from '../requestBeforeHook';
/**
 * WhiteLabelApi - axios parameter creator
 * @export
 */
export const WhiteLabelApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Create White Labels
         * @param {Array<WhiteLabelCreateRequestInner>} whiteLabelCreateRequestInner 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        create: async (whiteLabelCreateRequestInner: Array<WhiteLabelCreateRequestInner>, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'whiteLabelCreateRequestInner' is not null or undefined
            assertParamExists('create', 'whiteLabelCreateRequestInner', whiteLabelCreateRequestInner)
            const localVarPath = `/white_label/create`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions: AxiosRequestConfig = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = configuration && !isBrowser() ? { "User-Agent": configuration.userAgent } : {} as any;
            const localVarQueryParameter = {} as any;

            // authentication apiKey required
            await setApiKeyToObject({ object: localVarHeaderParameter, key: "authorization", keyParamName: "apiKey", configuration, prefix: "Bearer " })

    
            localVarHeaderParameter['Content-Type'] = 'application/json';


            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            requestBeforeHook({
                requestBody: whiteLabelCreateRequestInner,
                queryParameters: localVarQueryParameter,
                requestConfig: localVarRequestOptions,
                path: localVarPath,
                configuration,
                pathTemplate: '/white_label/create',
                httpMethod: 'POST'
            });
            localVarRequestOptions.data = serializeDataIfNeeded(whiteLabelCreateRequestInner, localVarRequestOptions, configuration)

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Delete White Labels
         * @param {DeleteWhiteLabelRequest} deleteWhiteLabelRequest 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        delete: async (deleteWhiteLabelRequest: DeleteWhiteLabelRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'deleteWhiteLabelRequest' is not null or undefined
            assertParamExists('delete', 'deleteWhiteLabelRequest', deleteWhiteLabelRequest)
            const localVarPath = `/white_label/delete`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions: AxiosRequestConfig = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = configuration && !isBrowser() ? { "User-Agent": configuration.userAgent } : {} as any;
            const localVarQueryParameter = {} as any;

            // authentication apiKey required
            await setApiKeyToObject({ object: localVarHeaderParameter, key: "authorization", keyParamName: "apiKey", configuration, prefix: "Bearer " })

    
            localVarHeaderParameter['Content-Type'] = 'application/json';


            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            requestBeforeHook({
                requestBody: deleteWhiteLabelRequest,
                queryParameters: localVarQueryParameter,
                requestConfig: localVarRequestOptions,
                path: localVarPath,
                configuration,
                pathTemplate: '/white_label/delete',
                httpMethod: 'POST'
            });
            localVarRequestOptions.data = serializeDataIfNeeded(deleteWhiteLabelRequest, localVarRequestOptions, configuration)

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary List White Labels
         * @param {ListWhiteLabelRequest} listWhiteLabelRequest 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        list: async (listWhiteLabelRequest: ListWhiteLabelRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'listWhiteLabelRequest' is not null or undefined
            assertParamExists('list', 'listWhiteLabelRequest', listWhiteLabelRequest)
            const localVarPath = `/white_label/list`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions: AxiosRequestConfig = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = configuration && !isBrowser() ? { "User-Agent": configuration.userAgent } : {} as any;
            const localVarQueryParameter = {} as any;

            // authentication apiKey required
            await setApiKeyToObject({ object: localVarHeaderParameter, key: "authorization", keyParamName: "apiKey", configuration, prefix: "Bearer " })

    
            localVarHeaderParameter['Content-Type'] = 'application/json';


            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            requestBeforeHook({
                requestBody: listWhiteLabelRequest,
                queryParameters: localVarQueryParameter,
                requestConfig: localVarRequestOptions,
                path: localVarPath,
                configuration,
                pathTemplate: '/white_label/list',
                httpMethod: 'POST'
            });
            localVarRequestOptions.data = serializeDataIfNeeded(listWhiteLabelRequest, localVarRequestOptions, configuration)

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Update White Label
         * @param {WhiteLabelUpdateRequest} whiteLabelUpdateRequest 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        update: async (whiteLabelUpdateRequest: WhiteLabelUpdateRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'whiteLabelUpdateRequest' is not null or undefined
            assertParamExists('update', 'whiteLabelUpdateRequest', whiteLabelUpdateRequest)
            const localVarPath = `/white_label/update`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions: AxiosRequestConfig = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = configuration && !isBrowser() ? { "User-Agent": configuration.userAgent } : {} as any;
            const localVarQueryParameter = {} as any;

            // authentication apiKey required
            await setApiKeyToObject({ object: localVarHeaderParameter, key: "authorization", keyParamName: "apiKey", configuration, prefix: "Bearer " })

    
            localVarHeaderParameter['Content-Type'] = 'application/json';


            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            requestBeforeHook({
                requestBody: whiteLabelUpdateRequest,
                queryParameters: localVarQueryParameter,
                requestConfig: localVarRequestOptions,
                path: localVarPath,
                configuration,
                pathTemplate: '/white_label/update',
                httpMethod: 'POST'
            });
            localVarRequestOptions.data = serializeDataIfNeeded(whiteLabelUpdateRequest, localVarRequestOptions, configuration)

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * WhiteLabelApi - functional programming interface
 * @export
 */
export const WhiteLabelApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = WhiteLabelApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @summary Create White Labels
         * @param {WhiteLabelApiCreateRequest} requestParameters Request parameters.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async create(requestParameters: WhiteLabelApiCreateRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<object>> {
            const whiteLabelCreateRequestInner: Array<WhiteLabelCreateRequestInner> = requestParameters;
            const localVarAxiosArgs = await localVarAxiosParamCreator.create(whiteLabelCreateRequestInner, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Delete White Labels
         * @param {WhiteLabelApiDeleteRequest} requestParameters Request parameters.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async delete(requestParameters: WhiteLabelApiDeleteRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<object>> {
            const deleteWhiteLabelRequest: DeleteWhiteLabelRequest = {
                ids: requestParameters.ids
            };
            const localVarAxiosArgs = await localVarAxiosParamCreator.delete(deleteWhiteLabelRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary List White Labels
         * @param {WhiteLabelApiListRequest} requestParameters Request parameters.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async list(requestParameters: WhiteLabelApiListRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<object>> {
            const listWhiteLabelRequest: ListWhiteLabelRequest = {
                pagination: requestParameters.pagination,
                order_by: requestParameters.order_by,
                order_dir: requestParameters.order_dir,
                filters: requestParameters.filters
            };
            const localVarAxiosArgs = await localVarAxiosParamCreator.list(listWhiteLabelRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Update White Label
         * @param {WhiteLabelApiUpdateRequest} requestParameters Request parameters.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async update(requestParameters: WhiteLabelApiUpdateRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<object>> {
            const whiteLabelUpdateRequest: WhiteLabelUpdateRequest = requestParameters;
            const localVarAxiosArgs = await localVarAxiosParamCreator.update(whiteLabelUpdateRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * WhiteLabelApi - factory interface
 * @export
 */
export const WhiteLabelApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = WhiteLabelApiFp(configuration)
    return {
        /**
         * 
         * @summary Create White Labels
         * @param {WhiteLabelApiCreateRequest} requestParameters Request parameters.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        create(requestParameters: WhiteLabelApiCreateRequest, options?: AxiosRequestConfig): AxiosPromise<object> {
            return localVarFp.create(requestParameters, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Delete White Labels
         * @param {WhiteLabelApiDeleteRequest} requestParameters Request parameters.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        delete(requestParameters: WhiteLabelApiDeleteRequest, options?: AxiosRequestConfig): AxiosPromise<object> {
            return localVarFp.delete(requestParameters, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary List White Labels
         * @param {WhiteLabelApiListRequest} requestParameters Request parameters.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        list(requestParameters: WhiteLabelApiListRequest, options?: AxiosRequestConfig): AxiosPromise<object> {
            return localVarFp.list(requestParameters, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Update White Label
         * @param {WhiteLabelApiUpdateRequest} requestParameters Request parameters.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        update(requestParameters: WhiteLabelApiUpdateRequest, options?: AxiosRequestConfig): AxiosPromise<object> {
            return localVarFp.update(requestParameters, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for create operation in WhiteLabelApi.
 * @export
 * @interface WhiteLabelApiCreateRequest
 */
export type WhiteLabelApiCreateRequest = Array<WhiteLabelCreateRequestInner>

/**
 * Request parameters for delete operation in WhiteLabelApi.
 * @export
 * @interface WhiteLabelApiDeleteRequest
 */
export type WhiteLabelApiDeleteRequest = {
    
} & DeleteWhiteLabelRequest

/**
 * Request parameters for list operation in WhiteLabelApi.
 * @export
 * @interface WhiteLabelApiListRequest
 */
export type WhiteLabelApiListRequest = {
    
} & ListWhiteLabelRequest

/**
 * Request parameters for update operation in WhiteLabelApi.
 * @export
 * @interface WhiteLabelApiUpdateRequest
 */
export type WhiteLabelApiUpdateRequest = {
    
} & WhiteLabelUpdateRequest

/**
 * WhiteLabelApiGenerated - object-oriented interface
 * @export
 * @class WhiteLabelApiGenerated
 * @extends {BaseAPI}
 */
export class WhiteLabelApiGenerated extends BaseAPI {
    /**
     * 
     * @summary Create White Labels
     * @param {WhiteLabelApiCreateRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof WhiteLabelApiGenerated
     */
    public create(requestParameters: WhiteLabelApiCreateRequest, options?: AxiosRequestConfig) {
        return WhiteLabelApiFp(this.configuration).create(requestParameters, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Delete White Labels
     * @param {WhiteLabelApiDeleteRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof WhiteLabelApiGenerated
     */
    public delete(requestParameters: WhiteLabelApiDeleteRequest, options?: AxiosRequestConfig) {
        return WhiteLabelApiFp(this.configuration).delete(requestParameters, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary List White Labels
     * @param {WhiteLabelApiListRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof WhiteLabelApiGenerated
     */
    public list(requestParameters: WhiteLabelApiListRequest, options?: AxiosRequestConfig) {
        return WhiteLabelApiFp(this.configuration).list(requestParameters, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Update White Label
     * @param {WhiteLabelApiUpdateRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof WhiteLabelApiGenerated
     */
    public update(requestParameters: WhiteLabelApiUpdateRequest, options?: AxiosRequestConfig) {
        return WhiteLabelApiFp(this.configuration).update(requestParameters, options).then((request) => request(this.axios, this.basePath));
    }
}
