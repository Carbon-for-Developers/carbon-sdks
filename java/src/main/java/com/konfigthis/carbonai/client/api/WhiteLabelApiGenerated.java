/*
 * Carbon
 * Connect external data to LLMs, no matter the source.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by Konfig (https://konfigthis.com).
 * Do not edit the class manually.
 */


package com.konfigthis.carbonai.client.api;

import com.konfigthis.carbonai.client.ApiCallback;
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.Pair;
import com.konfigthis.carbonai.client.ProgressRequestBody;
import com.konfigthis.carbonai.client.ProgressResponseBody;

import com.google.gson.reflect.TypeToken;

import java.io.IOException;


import com.konfigthis.carbonai.client.model.DeleteWhiteLabelRequest;
import com.konfigthis.carbonai.client.model.GDOSWhiteLabelInput;
import com.konfigthis.carbonai.client.model.GDOSWhiteLabelInput1;
import com.konfigthis.carbonai.client.model.GoogleDriveCredentialsProperty;
import com.konfigthis.carbonai.client.model.ListWhiteLabelRequest;
import com.konfigthis.carbonai.client.model.OrderDir;
import com.konfigthis.carbonai.client.model.Pagination;
import com.konfigthis.carbonai.client.model.WhiteLabelFilters;
import com.konfigthis.carbonai.client.model.WhiteLabelOrderByColumns;

import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import javax.ws.rs.core.GenericType;

public class WhiteLabelApiGenerated {
    private ApiClient localVarApiClient;
    private int localHostIndex;
    private String localCustomBaseUrl;

    public WhiteLabelApiGenerated() throws IllegalArgumentException {
        this(Configuration.getDefaultApiClient());
    }

    public WhiteLabelApiGenerated(ApiClient apiClient) throws IllegalArgumentException {
        this.localVarApiClient = apiClient;
    }

    public ApiClient getApiClient() {
        return localVarApiClient;
    }

    public void setApiClient(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    public int getHostIndex() {
        return localHostIndex;
    }

    public void setHostIndex(int hostIndex) {
        this.localHostIndex = hostIndex;
    }

    public String getCustomBaseUrl() {
        return localCustomBaseUrl;
    }

    public void setCustomBaseUrl(String customBaseUrl) {
        this.localCustomBaseUrl = customBaseUrl;
    }

    private okhttp3.Call createCall(List<GDOSWhiteLabelInput> gdOSWhiteLabelInput, final ApiCallback _callback) throws ApiException {
        String basePath = null;
        // Operation Servers
        String[] localBasePaths = new String[] {  };

        // Determine Base Path to Use
        if (localCustomBaseUrl != null){
            basePath = localCustomBaseUrl;
        } else if ( localBasePaths.length > 0 ) {
            basePath = localBasePaths[localHostIndex];
        } else {
            basePath = null;
        }

        Object localVarPostBody = gdOSWhiteLabelInput;

        // create path and map variables
        String localVarPath = "/white_label/create";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        if (localVarContentType != null) {
            localVarHeaderParams.put("Content-Type", localVarContentType);
        }

        String[] localVarAuthNames = new String[] { "apiKey" };
        return localVarApiClient.buildCall(basePath, localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call createValidateBeforeCall(List<GDOSWhiteLabelInput> gdOSWhiteLabelInput, final ApiCallback _callback) throws ApiException {
        // verify the required parameter 'gdOSWhiteLabelInput' is set
        if (gdOSWhiteLabelInput == null) {
            throw new ApiException("Missing the required parameter 'gdOSWhiteLabelInput' when calling create(Async)");
        }

        return createCall(gdOSWhiteLabelInput, _callback);

    }


    private ApiResponse<Object> createWithHttpInfo(List<GDOSWhiteLabelInput> gdOSWhiteLabelInput) throws ApiException {
        okhttp3.Call localVarCall = createValidateBeforeCall(gdOSWhiteLabelInput, null);
        Type localVarReturnType = new TypeToken<Object>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    private okhttp3.Call createAsync(List<GDOSWhiteLabelInput> gdOSWhiteLabelInput, final ApiCallback<Object> _callback) throws ApiException {

        okhttp3.Call localVarCall = createValidateBeforeCall(gdOSWhiteLabelInput, _callback);
        Type localVarReturnType = new TypeToken<Object>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }

    public abstract class CreateRequestBuilderGenerated {
        List<GDOSWhiteLabelInput> gdOSWhiteLabelInput;

        public CreateRequestBuilderGenerated() {
        }

        /**
         * Set gdOSWhiteLabelInput
         * @param gdOSWhiteLabelInput  (optional)
         * @return WhiteLabelApi.CreateRequestBuilder
         */
        public WhiteLabelApi.CreateRequestBuilder gdOSWhiteLabelInput(List<GDOSWhiteLabelInput> gdOSWhiteLabelInput) {
            this.gdOSWhiteLabelInput = gdOSWhiteLabelInput;
            return (WhiteLabelApi.CreateRequestBuilder) this;
        }

        /**
         * Build call for create
         * @param _callback ApiCallback API callback
         * @return Call to execute
         * @throws ApiException If fail to serialize the request body object
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public okhttp3.Call buildCall(final ApiCallback _callback) throws ApiException {
            List<GDOSWhiteLabelInput> gdOSWhiteLabelInput = buildBodyParams();
            return createCall(gdOSWhiteLabelInput, _callback);
        }

        private List<GDOSWhiteLabelInput> buildBodyParams() {
            return this.gdOSWhiteLabelInput;
        }

        /**
         * Execute create request
         * @return Object
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public Object execute() throws ApiException {
            List<GDOSWhiteLabelInput> gdOSWhiteLabelInput = buildBodyParams();
            ApiResponse<Object> localVarResp = createWithHttpInfo(gdOSWhiteLabelInput);
            return localVarResp.getResponseBody();
        }

        /**
         * Execute create request with HTTP info returned
         * @return ApiResponse&lt;Object&gt;
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public ApiResponse<Object> executeWithHttpInfo() throws ApiException {
            List<GDOSWhiteLabelInput> gdOSWhiteLabelInput = buildBodyParams();
            return createWithHttpInfo(gdOSWhiteLabelInput);
        }

        /**
         * Execute create request (asynchronously)
         * @param _callback The callback to be executed when the API call finishes
         * @return The request call
         * @throws ApiException If fail to process the API call, e.g. serializing the request body object
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public okhttp3.Call executeAsync(final ApiCallback<Object> _callback) throws ApiException {
            List<GDOSWhiteLabelInput> gdOSWhiteLabelInput = buildBodyParams();
            return createAsync(gdOSWhiteLabelInput, _callback);
        }
    }

    /**
     * Create White Labels
     * 
     * @param gdOSWhiteLabelInput  (required)
     * @return CreateRequestBuilder
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
     </table>
     */
    public WhiteLabelApi.CreateRequestBuilder create() throws IllegalArgumentException {
        return ((WhiteLabelApi) this).new CreateRequestBuilder();
    }
    private okhttp3.Call deleteCall(DeleteWhiteLabelRequest deleteWhiteLabelRequest, final ApiCallback _callback) throws ApiException {
        String basePath = null;
        // Operation Servers
        String[] localBasePaths = new String[] {  };

        // Determine Base Path to Use
        if (localCustomBaseUrl != null){
            basePath = localCustomBaseUrl;
        } else if ( localBasePaths.length > 0 ) {
            basePath = localBasePaths[localHostIndex];
        } else {
            basePath = null;
        }

        Object localVarPostBody = deleteWhiteLabelRequest;

        // create path and map variables
        String localVarPath = "/white_label/delete";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        if (localVarContentType != null) {
            localVarHeaderParams.put("Content-Type", localVarContentType);
        }

        String[] localVarAuthNames = new String[] { "apiKey" };
        return localVarApiClient.buildCall(basePath, localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call deleteValidateBeforeCall(DeleteWhiteLabelRequest deleteWhiteLabelRequest, final ApiCallback _callback) throws ApiException {
        // verify the required parameter 'deleteWhiteLabelRequest' is set
        if (deleteWhiteLabelRequest == null) {
            throw new ApiException("Missing the required parameter 'deleteWhiteLabelRequest' when calling delete(Async)");
        }

        return deleteCall(deleteWhiteLabelRequest, _callback);

    }


    private ApiResponse<Object> deleteWithHttpInfo(DeleteWhiteLabelRequest deleteWhiteLabelRequest) throws ApiException {
        okhttp3.Call localVarCall = deleteValidateBeforeCall(deleteWhiteLabelRequest, null);
        Type localVarReturnType = new TypeToken<Object>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    private okhttp3.Call deleteAsync(DeleteWhiteLabelRequest deleteWhiteLabelRequest, final ApiCallback<Object> _callback) throws ApiException {

        okhttp3.Call localVarCall = deleteValidateBeforeCall(deleteWhiteLabelRequest, _callback);
        Type localVarReturnType = new TypeToken<Object>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }

    public abstract class DeleteRequestBuilderGenerated {
        final List<Integer> ids;

        public DeleteRequestBuilderGenerated(List<Integer> ids) {
            this.ids = ids;
        }

        /**
         * Build call for delete
         * @param _callback ApiCallback API callback
         * @return Call to execute
         * @throws ApiException If fail to serialize the request body object
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public okhttp3.Call buildCall(final ApiCallback _callback) throws ApiException {
            DeleteWhiteLabelRequest deleteWhiteLabelRequest = buildBodyParams();
            return deleteCall(deleteWhiteLabelRequest, _callback);
        }

        private DeleteWhiteLabelRequest buildBodyParams() {
            DeleteWhiteLabelRequest deleteWhiteLabelRequest = new DeleteWhiteLabelRequest();
            deleteWhiteLabelRequest.ids(this.ids);
            return deleteWhiteLabelRequest;
        }

        /**
         * Execute delete request
         * @return Object
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public Object execute() throws ApiException {
            DeleteWhiteLabelRequest deleteWhiteLabelRequest = buildBodyParams();
            ApiResponse<Object> localVarResp = deleteWithHttpInfo(deleteWhiteLabelRequest);
            return localVarResp.getResponseBody();
        }

        /**
         * Execute delete request with HTTP info returned
         * @return ApiResponse&lt;Object&gt;
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public ApiResponse<Object> executeWithHttpInfo() throws ApiException {
            DeleteWhiteLabelRequest deleteWhiteLabelRequest = buildBodyParams();
            return deleteWithHttpInfo(deleteWhiteLabelRequest);
        }

        /**
         * Execute delete request (asynchronously)
         * @param _callback The callback to be executed when the API call finishes
         * @return The request call
         * @throws ApiException If fail to process the API call, e.g. serializing the request body object
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public okhttp3.Call executeAsync(final ApiCallback<Object> _callback) throws ApiException {
            DeleteWhiteLabelRequest deleteWhiteLabelRequest = buildBodyParams();
            return deleteAsync(deleteWhiteLabelRequest, _callback);
        }
    }

    /**
     * Delete White Labels
     * 
     * @param deleteWhiteLabelRequest  (required)
     * @return DeleteRequestBuilder
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
     </table>
     */
    public WhiteLabelApi.DeleteRequestBuilder delete(List<Integer> ids) throws IllegalArgumentException {
        if (ids == null) throw new IllegalArgumentException("\"ids\" is required but got null");
        return ((WhiteLabelApi) this).new DeleteRequestBuilder(ids);
    }
    private okhttp3.Call listCall(ListWhiteLabelRequest listWhiteLabelRequest, final ApiCallback _callback) throws ApiException {
        String basePath = null;
        // Operation Servers
        String[] localBasePaths = new String[] {  };

        // Determine Base Path to Use
        if (localCustomBaseUrl != null){
            basePath = localCustomBaseUrl;
        } else if ( localBasePaths.length > 0 ) {
            basePath = localBasePaths[localHostIndex];
        } else {
            basePath = null;
        }

        Object localVarPostBody = listWhiteLabelRequest;

        // create path and map variables
        String localVarPath = "/white_label/list";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        if (localVarContentType != null) {
            localVarHeaderParams.put("Content-Type", localVarContentType);
        }

        String[] localVarAuthNames = new String[] { "apiKey" };
        return localVarApiClient.buildCall(basePath, localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call listValidateBeforeCall(ListWhiteLabelRequest listWhiteLabelRequest, final ApiCallback _callback) throws ApiException {
        // verify the required parameter 'listWhiteLabelRequest' is set
        if (listWhiteLabelRequest == null) {
            throw new ApiException("Missing the required parameter 'listWhiteLabelRequest' when calling list(Async)");
        }

        return listCall(listWhiteLabelRequest, _callback);

    }


    private ApiResponse<Object> listWithHttpInfo(ListWhiteLabelRequest listWhiteLabelRequest) throws ApiException {
        okhttp3.Call localVarCall = listValidateBeforeCall(listWhiteLabelRequest, null);
        Type localVarReturnType = new TypeToken<Object>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    private okhttp3.Call listAsync(ListWhiteLabelRequest listWhiteLabelRequest, final ApiCallback<Object> _callback) throws ApiException {

        okhttp3.Call localVarCall = listValidateBeforeCall(listWhiteLabelRequest, _callback);
        Type localVarReturnType = new TypeToken<Object>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }

    public abstract class ListRequestBuilderGenerated {
        Pagination pagination;
        WhiteLabelOrderByColumns orderBy;
        OrderDir orderDir;
        WhiteLabelFilters filters;

        public ListRequestBuilderGenerated() {
        }

        /**
         * Set pagination
         * @param pagination  (optional)
         * @return WhiteLabelApi.ListRequestBuilder
         */
        public WhiteLabelApi.ListRequestBuilder pagination(Pagination pagination) {
            this.pagination = pagination;
            return (WhiteLabelApi.ListRequestBuilder) this;
        }
        
        /**
         * Set orderBy
         * @param orderBy  (optional)
         * @return WhiteLabelApi.ListRequestBuilder
         */
        public WhiteLabelApi.ListRequestBuilder orderBy(WhiteLabelOrderByColumns orderBy) {
            this.orderBy = orderBy;
            return (WhiteLabelApi.ListRequestBuilder) this;
        }
        
        /**
         * Set orderDir
         * @param orderDir  (optional)
         * @return WhiteLabelApi.ListRequestBuilder
         */
        public WhiteLabelApi.ListRequestBuilder orderDir(OrderDir orderDir) {
            this.orderDir = orderDir;
            return (WhiteLabelApi.ListRequestBuilder) this;
        }
        
        /**
         * Set filters
         * @param filters  (optional)
         * @return WhiteLabelApi.ListRequestBuilder
         */
        public WhiteLabelApi.ListRequestBuilder filters(WhiteLabelFilters filters) {
            this.filters = filters;
            return (WhiteLabelApi.ListRequestBuilder) this;
        }
        
        /**
         * Build call for list
         * @param _callback ApiCallback API callback
         * @return Call to execute
         * @throws ApiException If fail to serialize the request body object
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public okhttp3.Call buildCall(final ApiCallback _callback) throws ApiException {
            ListWhiteLabelRequest listWhiteLabelRequest = buildBodyParams();
            return listCall(listWhiteLabelRequest, _callback);
        }

        private ListWhiteLabelRequest buildBodyParams() {
            ListWhiteLabelRequest listWhiteLabelRequest = new ListWhiteLabelRequest();
            listWhiteLabelRequest.pagination(this.pagination);
            listWhiteLabelRequest.orderBy(this.orderBy);
            listWhiteLabelRequest.orderDir(this.orderDir);
            listWhiteLabelRequest.filters(this.filters);
            return listWhiteLabelRequest;
        }

        /**
         * Execute list request
         * @return Object
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public Object execute() throws ApiException {
            ListWhiteLabelRequest listWhiteLabelRequest = buildBodyParams();
            ApiResponse<Object> localVarResp = listWithHttpInfo(listWhiteLabelRequest);
            return localVarResp.getResponseBody();
        }

        /**
         * Execute list request with HTTP info returned
         * @return ApiResponse&lt;Object&gt;
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public ApiResponse<Object> executeWithHttpInfo() throws ApiException {
            ListWhiteLabelRequest listWhiteLabelRequest = buildBodyParams();
            return listWithHttpInfo(listWhiteLabelRequest);
        }

        /**
         * Execute list request (asynchronously)
         * @param _callback The callback to be executed when the API call finishes
         * @return The request call
         * @throws ApiException If fail to process the API call, e.g. serializing the request body object
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public okhttp3.Call executeAsync(final ApiCallback<Object> _callback) throws ApiException {
            ListWhiteLabelRequest listWhiteLabelRequest = buildBodyParams();
            return listAsync(listWhiteLabelRequest, _callback);
        }
    }

    /**
     * List White Labels
     * 
     * @param listWhiteLabelRequest  (required)
     * @return ListRequestBuilder
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
     </table>
     */
    public WhiteLabelApi.ListRequestBuilder list() throws IllegalArgumentException {
        return ((WhiteLabelApi) this).new ListRequestBuilder();
    }
    private okhttp3.Call updateCall(GDOSWhiteLabelInput1 body, final ApiCallback _callback) throws ApiException {
        String basePath = null;
        // Operation Servers
        String[] localBasePaths = new String[] {  };

        // Determine Base Path to Use
        if (localCustomBaseUrl != null){
            basePath = localCustomBaseUrl;
        } else if ( localBasePaths.length > 0 ) {
            basePath = localBasePaths[localHostIndex];
        } else {
            basePath = null;
        }

        Object localVarPostBody = body;

        // create path and map variables
        String localVarPath = "/white_label/update";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        if (localVarContentType != null) {
            localVarHeaderParams.put("Content-Type", localVarContentType);
        }

        String[] localVarAuthNames = new String[] { "apiKey" };
        return localVarApiClient.buildCall(basePath, localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call updateValidateBeforeCall(GDOSWhiteLabelInput1 body, final ApiCallback _callback) throws ApiException {
        // verify the required parameter 'body' is set
        if (body == null) {
            throw new ApiException("Missing the required parameter 'body' when calling update(Async)");
        }

        return updateCall(body, _callback);

    }


    private ApiResponse<Object> updateWithHttpInfo(GDOSWhiteLabelInput1 body) throws ApiException {
        okhttp3.Call localVarCall = updateValidateBeforeCall(body, null);
        Type localVarReturnType = new TypeToken<Object>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    private okhttp3.Call updateAsync(GDOSWhiteLabelInput1 body, final ApiCallback<Object> _callback) throws ApiException {

        okhttp3.Call localVarCall = updateValidateBeforeCall(body, _callback);
        Type localVarReturnType = new TypeToken<Object>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }

    public abstract class UpdateRequestBuilderGenerated {
        String dataSourceType;
        GoogleDriveCredentialsProperty credentials;

        public UpdateRequestBuilderGenerated() {
        }

        /**
         * Set dataSourceType
         * @param dataSourceType  (optional)
         * @return WhiteLabelApi.UpdateRequestBuilder
         */
        public WhiteLabelApi.UpdateRequestBuilder dataSourceType(String dataSourceType) {
            this.dataSourceType = dataSourceType;
            return (WhiteLabelApi.UpdateRequestBuilder) this;
        }
        
        /**
         * Set credentials
         * @param credentials  (optional)
         * @return WhiteLabelApi.UpdateRequestBuilder
         */
        public WhiteLabelApi.UpdateRequestBuilder credentials(GoogleDriveCredentialsProperty credentials) {
            this.credentials = credentials;
            return (WhiteLabelApi.UpdateRequestBuilder) this;
        }
        
        /**
         * Build call for update
         * @param _callback ApiCallback API callback
         * @return Call to execute
         * @throws ApiException If fail to serialize the request body object
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public okhttp3.Call buildCall(final ApiCallback _callback) throws ApiException {
            GDOSWhiteLabelInput1 body = buildBodyParams();
            return updateCall(body, _callback);
        }

        private GDOSWhiteLabelInput1 buildBodyParams() {
            GDOSWhiteLabelInput1 body = new GDOSWhiteLabelInput1();
            return body;
        }

        /**
         * Execute update request
         * @return Object
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public Object execute() throws ApiException {
            GDOSWhiteLabelInput1 body = buildBodyParams();
            ApiResponse<Object> localVarResp = updateWithHttpInfo(body);
            return localVarResp.getResponseBody();
        }

        /**
         * Execute update request with HTTP info returned
         * @return ApiResponse&lt;Object&gt;
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public ApiResponse<Object> executeWithHttpInfo() throws ApiException {
            GDOSWhiteLabelInput1 body = buildBodyParams();
            return updateWithHttpInfo(body);
        }

        /**
         * Execute update request (asynchronously)
         * @param _callback The callback to be executed when the API call finishes
         * @return The request call
         * @throws ApiException If fail to process the API call, e.g. serializing the request body object
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public okhttp3.Call executeAsync(final ApiCallback<Object> _callback) throws ApiException {
            GDOSWhiteLabelInput1 body = buildBodyParams();
            return updateAsync(body, _callback);
        }
    }

    /**
     * Update White Label
     * 
     * @param body  (required)
     * @return UpdateRequestBuilder
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
     </table>
     */
    public WhiteLabelApi.UpdateRequestBuilder update() throws IllegalArgumentException {
        return ((WhiteLabelApi) this).new UpdateRequestBuilder();
    }
}
