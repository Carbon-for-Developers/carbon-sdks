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


import com.konfigthis.carbonai.client.model.DeleteUsersInput;
import com.konfigthis.carbonai.client.model.GenericSuccessResponse;
import com.konfigthis.carbonai.client.model.ModifyUserConfigurationInput;
import com.konfigthis.carbonai.client.model.UpdateUsersInput;
import com.konfigthis.carbonai.client.model.UserRequestContent;
import com.konfigthis.carbonai.client.model.UserResponse;

import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import javax.ws.rs.core.GenericType;

public class UsersApiGenerated {
    private ApiClient localVarApiClient;
    private int localHostIndex;
    private String localCustomBaseUrl;

    public UsersApiGenerated() throws IllegalArgumentException {
        this(Configuration.getDefaultApiClient());
    }

    public UsersApiGenerated(ApiClient apiClient) throws IllegalArgumentException {
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

    private okhttp3.Call deleteCall(DeleteUsersInput deleteUsersInput, final ApiCallback _callback) throws ApiException {
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

        Object localVarPostBody = deleteUsersInput;

        // create path and map variables
        String localVarPath = "/delete_users";

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
    private okhttp3.Call deleteValidateBeforeCall(DeleteUsersInput deleteUsersInput, final ApiCallback _callback) throws ApiException {
        // verify the required parameter 'deleteUsersInput' is set
        if (deleteUsersInput == null) {
            throw new ApiException("Missing the required parameter 'deleteUsersInput' when calling delete(Async)");
        }

        return deleteCall(deleteUsersInput, _callback);

    }


    private ApiResponse<GenericSuccessResponse> deleteWithHttpInfo(DeleteUsersInput deleteUsersInput) throws ApiException {
        okhttp3.Call localVarCall = deleteValidateBeforeCall(deleteUsersInput, null);
        Type localVarReturnType = new TypeToken<GenericSuccessResponse>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    private okhttp3.Call deleteAsync(DeleteUsersInput deleteUsersInput, final ApiCallback<GenericSuccessResponse> _callback) throws ApiException {

        okhttp3.Call localVarCall = deleteValidateBeforeCall(deleteUsersInput, _callback);
        Type localVarReturnType = new TypeToken<GenericSuccessResponse>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }

    public class DeleteRequestBuilder {
        private final List<String> customerIds;

        private DeleteRequestBuilder(List<String> customerIds) {
            this.customerIds = customerIds;
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
            DeleteUsersInput deleteUsersInput = buildBodyParams();
            return deleteCall(deleteUsersInput, _callback);
        }

        private DeleteUsersInput buildBodyParams() {
            DeleteUsersInput deleteUsersInput = new DeleteUsersInput();
            deleteUsersInput.customerIds(this.customerIds);
            return deleteUsersInput;
        }

        /**
         * Execute delete request
         * @return GenericSuccessResponse
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public GenericSuccessResponse execute() throws ApiException {
            DeleteUsersInput deleteUsersInput = buildBodyParams();
            ApiResponse<GenericSuccessResponse> localVarResp = deleteWithHttpInfo(deleteUsersInput);
            return localVarResp.getResponseBody();
        }

        /**
         * Execute delete request with HTTP info returned
         * @return ApiResponse&lt;GenericSuccessResponse&gt;
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public ApiResponse<GenericSuccessResponse> executeWithHttpInfo() throws ApiException {
            DeleteUsersInput deleteUsersInput = buildBodyParams();
            return deleteWithHttpInfo(deleteUsersInput);
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
        public okhttp3.Call executeAsync(final ApiCallback<GenericSuccessResponse> _callback) throws ApiException {
            DeleteUsersInput deleteUsersInput = buildBodyParams();
            return deleteAsync(deleteUsersInput, _callback);
        }
    }

    /**
     * Delete Users
     * 
     * @param deleteUsersInput  (required)
     * @return DeleteRequestBuilder
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
     </table>
     */
    public DeleteRequestBuilder delete(List<String> customerIds) throws IllegalArgumentException {
        if (customerIds == null) throw new IllegalArgumentException("\"customerIds\" is required but got null");
        return new DeleteRequestBuilder(customerIds);
    }
    private okhttp3.Call getCall(UserRequestContent userRequestContent, final ApiCallback _callback) throws ApiException {
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

        Object localVarPostBody = userRequestContent;

        // create path and map variables
        String localVarPath = "/user";

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
    private okhttp3.Call getValidateBeforeCall(UserRequestContent userRequestContent, final ApiCallback _callback) throws ApiException {
        // verify the required parameter 'userRequestContent' is set
        if (userRequestContent == null) {
            throw new ApiException("Missing the required parameter 'userRequestContent' when calling get(Async)");
        }

        return getCall(userRequestContent, _callback);

    }


    private ApiResponse<UserResponse> getWithHttpInfo(UserRequestContent userRequestContent) throws ApiException {
        okhttp3.Call localVarCall = getValidateBeforeCall(userRequestContent, null);
        Type localVarReturnType = new TypeToken<UserResponse>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    private okhttp3.Call getAsync(UserRequestContent userRequestContent, final ApiCallback<UserResponse> _callback) throws ApiException {

        okhttp3.Call localVarCall = getValidateBeforeCall(userRequestContent, _callback);
        Type localVarReturnType = new TypeToken<UserResponse>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }

    public class GetRequestBuilder {
        private final String customerId;

        private GetRequestBuilder(String customerId) {
            this.customerId = customerId;
        }

        /**
         * Build call for get
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
            UserRequestContent userRequestContent = buildBodyParams();
            return getCall(userRequestContent, _callback);
        }

        private UserRequestContent buildBodyParams() {
            UserRequestContent userRequestContent = new UserRequestContent();
            userRequestContent.customerId(this.customerId);
            return userRequestContent;
        }

        /**
         * Execute get request
         * @return UserResponse
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public UserResponse execute() throws ApiException {
            UserRequestContent userRequestContent = buildBodyParams();
            ApiResponse<UserResponse> localVarResp = getWithHttpInfo(userRequestContent);
            return localVarResp.getResponseBody();
        }

        /**
         * Execute get request with HTTP info returned
         * @return ApiResponse&lt;UserResponse&gt;
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public ApiResponse<UserResponse> executeWithHttpInfo() throws ApiException {
            UserRequestContent userRequestContent = buildBodyParams();
            return getWithHttpInfo(userRequestContent);
        }

        /**
         * Execute get request (asynchronously)
         * @param _callback The callback to be executed when the API call finishes
         * @return The request call
         * @throws ApiException If fail to process the API call, e.g. serializing the request body object
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public okhttp3.Call executeAsync(final ApiCallback<UserResponse> _callback) throws ApiException {
            UserRequestContent userRequestContent = buildBodyParams();
            return getAsync(userRequestContent, _callback);
        }
    }

    /**
     * User Endpoint
     * 
     * @param userRequestContent  (required)
     * @return GetRequestBuilder
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
     </table>
     */
    public GetRequestBuilder get(String customerId) throws IllegalArgumentException {
        if (customerId == null) throw new IllegalArgumentException("\"customerId\" is required but got null");
            

        return new GetRequestBuilder(customerId);
    }
    private okhttp3.Call toggleUserFeaturesCall(ModifyUserConfigurationInput modifyUserConfigurationInput, final ApiCallback _callback) throws ApiException {
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

        Object localVarPostBody = modifyUserConfigurationInput;

        // create path and map variables
        String localVarPath = "/modify_user_configuration";

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

        String[] localVarAuthNames = new String[] { "accessToken", "apiKey", "customerId" };
        return localVarApiClient.buildCall(basePath, localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @Deprecated
    @SuppressWarnings("rawtypes")
    private okhttp3.Call toggleUserFeaturesValidateBeforeCall(ModifyUserConfigurationInput modifyUserConfigurationInput, final ApiCallback _callback) throws ApiException {
        // verify the required parameter 'modifyUserConfigurationInput' is set
        if (modifyUserConfigurationInput == null) {
            throw new ApiException("Missing the required parameter 'modifyUserConfigurationInput' when calling toggleUserFeatures(Async)");
        }

        return toggleUserFeaturesCall(modifyUserConfigurationInput, _callback);

    }


    private ApiResponse<GenericSuccessResponse> toggleUserFeaturesWithHttpInfo(ModifyUserConfigurationInput modifyUserConfigurationInput) throws ApiException {
        okhttp3.Call localVarCall = toggleUserFeaturesValidateBeforeCall(modifyUserConfigurationInput, null);
        Type localVarReturnType = new TypeToken<GenericSuccessResponse>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    private okhttp3.Call toggleUserFeaturesAsync(ModifyUserConfigurationInput modifyUserConfigurationInput, final ApiCallback<GenericSuccessResponse> _callback) throws ApiException {

        okhttp3.Call localVarCall = toggleUserFeaturesValidateBeforeCall(modifyUserConfigurationInput, _callback);
        Type localVarReturnType = new TypeToken<GenericSuccessResponse>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }

    public class ToggleUserFeaturesRequestBuilder {
        private final String configurationKeyName;
        private final Object value;

        private ToggleUserFeaturesRequestBuilder(String configurationKeyName, Object value) {
            this.configurationKeyName = configurationKeyName;
            this.value = value;
        }

        /**
         * Build call for toggleUserFeatures
         * @param _callback ApiCallback API callback
         * @return Call to execute
         * @throws ApiException If fail to serialize the request body object
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         * @deprecated
         */
        @Deprecated
        public okhttp3.Call buildCall(final ApiCallback _callback) throws ApiException {
            ModifyUserConfigurationInput modifyUserConfigurationInput = buildBodyParams();
            return toggleUserFeaturesCall(modifyUserConfigurationInput, _callback);
        }

        private ModifyUserConfigurationInput buildBodyParams() {
            ModifyUserConfigurationInput modifyUserConfigurationInput = new ModifyUserConfigurationInput();
            modifyUserConfigurationInput.configurationKeyName(this.configurationKeyName);
            modifyUserConfigurationInput.value(this.value);
            return modifyUserConfigurationInput;
        }

        /**
         * Execute toggleUserFeatures request
         * @return GenericSuccessResponse
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         * @deprecated
         */
        @Deprecated
        public GenericSuccessResponse execute() throws ApiException {
            ModifyUserConfigurationInput modifyUserConfigurationInput = buildBodyParams();
            ApiResponse<GenericSuccessResponse> localVarResp = toggleUserFeaturesWithHttpInfo(modifyUserConfigurationInput);
            return localVarResp.getResponseBody();
        }

        /**
         * Execute toggleUserFeatures request with HTTP info returned
         * @return ApiResponse&lt;GenericSuccessResponse&gt;
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         * @deprecated
         */
        @Deprecated
        public ApiResponse<GenericSuccessResponse> executeWithHttpInfo() throws ApiException {
            ModifyUserConfigurationInput modifyUserConfigurationInput = buildBodyParams();
            return toggleUserFeaturesWithHttpInfo(modifyUserConfigurationInput);
        }

        /**
         * Execute toggleUserFeatures request (asynchronously)
         * @param _callback The callback to be executed when the API call finishes
         * @return The request call
         * @throws ApiException If fail to process the API call, e.g. serializing the request body object
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         * @deprecated
         */
        @Deprecated
        public okhttp3.Call executeAsync(final ApiCallback<GenericSuccessResponse> _callback) throws ApiException {
            ModifyUserConfigurationInput modifyUserConfigurationInput = buildBodyParams();
            return toggleUserFeaturesAsync(modifyUserConfigurationInput, _callback);
        }
    }

    /**
     * Toggle User Features
     * 
     * @param modifyUserConfigurationInput  (required)
     * @return ToggleUserFeaturesRequestBuilder
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
     </table>
     * @deprecated
     */
    @Deprecated
    public ToggleUserFeaturesRequestBuilder toggleUserFeatures(String configurationKeyName, Object value) throws IllegalArgumentException {
        if (configurationKeyName == null) throw new IllegalArgumentException("\"configurationKeyName\" is required but got null");
            

        if (value == null) throw new IllegalArgumentException("\"value\" is required but got null");
        return new ToggleUserFeaturesRequestBuilder(configurationKeyName, value);
    }
    private okhttp3.Call updateUsersCall(UpdateUsersInput updateUsersInput, final ApiCallback _callback) throws ApiException {
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

        Object localVarPostBody = updateUsersInput;

        // create path and map variables
        String localVarPath = "/update_users";

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
    private okhttp3.Call updateUsersValidateBeforeCall(UpdateUsersInput updateUsersInput, final ApiCallback _callback) throws ApiException {
        // verify the required parameter 'updateUsersInput' is set
        if (updateUsersInput == null) {
            throw new ApiException("Missing the required parameter 'updateUsersInput' when calling updateUsers(Async)");
        }

        return updateUsersCall(updateUsersInput, _callback);

    }


    private ApiResponse<GenericSuccessResponse> updateUsersWithHttpInfo(UpdateUsersInput updateUsersInput) throws ApiException {
        okhttp3.Call localVarCall = updateUsersValidateBeforeCall(updateUsersInput, null);
        Type localVarReturnType = new TypeToken<GenericSuccessResponse>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    private okhttp3.Call updateUsersAsync(UpdateUsersInput updateUsersInput, final ApiCallback<GenericSuccessResponse> _callback) throws ApiException {

        okhttp3.Call localVarCall = updateUsersValidateBeforeCall(updateUsersInput, _callback);
        Type localVarReturnType = new TypeToken<GenericSuccessResponse>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }

    public class UpdateUsersRequestBuilder {
        private final List<String> customerIds;
        private Object autoSyncEnabledSources;
        private Integer maxFiles;
        private Integer maxFilesPerUpload;

        private UpdateUsersRequestBuilder(List<String> customerIds) {
            this.customerIds = customerIds;
        }

        /**
         * Set autoSyncEnabledSources
         * @param autoSyncEnabledSources  (optional)
         * @return UpdateUsersRequestBuilder
         */
        public UpdateUsersRequestBuilder autoSyncEnabledSources(Object autoSyncEnabledSources) {
            this.autoSyncEnabledSources = autoSyncEnabledSources;
            return this;
        }
        
        /**
         * Set maxFiles
         * @param maxFiles Custom file upload limit for the user over *all* user&#39;s files across all uploads.          If set, then the user will not be allowed to upload more files than this limit. If not set, or if set to -1,         then the user will have no limit. (optional)
         * @return UpdateUsersRequestBuilder
         */
        public UpdateUsersRequestBuilder maxFiles(Integer maxFiles) {
            this.maxFiles = maxFiles;
            return this;
        }
        
        /**
         * Set maxFilesPerUpload
         * @param maxFilesPerUpload Custom file upload limit for the user across a single upload.         If set, then the user will not be allowed to upload more files than this limit in a single upload. If not set,         or if set to -1, then the user will have no limit. (optional)
         * @return UpdateUsersRequestBuilder
         */
        public UpdateUsersRequestBuilder maxFilesPerUpload(Integer maxFilesPerUpload) {
            this.maxFilesPerUpload = maxFilesPerUpload;
            return this;
        }
        
        /**
         * Build call for updateUsers
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
            UpdateUsersInput updateUsersInput = buildBodyParams();
            return updateUsersCall(updateUsersInput, _callback);
        }

        private UpdateUsersInput buildBodyParams() {
            UpdateUsersInput updateUsersInput = new UpdateUsersInput();
            updateUsersInput.autoSyncEnabledSources(this.autoSyncEnabledSources);
            updateUsersInput.maxFiles(this.maxFiles);
            updateUsersInput.maxFilesPerUpload(this.maxFilesPerUpload);
            updateUsersInput.customerIds(this.customerIds);
            return updateUsersInput;
        }

        /**
         * Execute updateUsers request
         * @return GenericSuccessResponse
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public GenericSuccessResponse execute() throws ApiException {
            UpdateUsersInput updateUsersInput = buildBodyParams();
            ApiResponse<GenericSuccessResponse> localVarResp = updateUsersWithHttpInfo(updateUsersInput);
            return localVarResp.getResponseBody();
        }

        /**
         * Execute updateUsers request with HTTP info returned
         * @return ApiResponse&lt;GenericSuccessResponse&gt;
         * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public ApiResponse<GenericSuccessResponse> executeWithHttpInfo() throws ApiException {
            UpdateUsersInput updateUsersInput = buildBodyParams();
            return updateUsersWithHttpInfo(updateUsersInput);
        }

        /**
         * Execute updateUsers request (asynchronously)
         * @param _callback The callback to be executed when the API call finishes
         * @return The request call
         * @throws ApiException If fail to process the API call, e.g. serializing the request body object
         * @http.response.details
         <table summary="Response Details" border="1">
            <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
            <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
         </table>
         */
        public okhttp3.Call executeAsync(final ApiCallback<GenericSuccessResponse> _callback) throws ApiException {
            UpdateUsersInput updateUsersInput = buildBodyParams();
            return updateUsersAsync(updateUsersInput, _callback);
        }
    }

    /**
     * Update Users
     * 
     * @param updateUsersInput  (required)
     * @return UpdateUsersRequestBuilder
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Successful Response </td><td>  -  </td></tr>
     </table>
     */
    public UpdateUsersRequestBuilder updateUsers(List<String> customerIds) throws IllegalArgumentException {
        if (customerIds == null) throw new IllegalArgumentException("\"customerIds\" is required but got null");
        return new UpdateUsersRequestBuilder(customerIds);
    }
}
