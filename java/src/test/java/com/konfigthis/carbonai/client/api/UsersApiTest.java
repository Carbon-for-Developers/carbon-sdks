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

import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.model.DeleteUsersInput;
import com.konfigthis.carbonai.client.model.GenericSuccessResponse;
import com.konfigthis.carbonai.client.model.ListUsersFilters;
import com.konfigthis.carbonai.client.model.ListUsersOrderByTypes;
import com.konfigthis.carbonai.client.model.ListUsersRequest;
import com.konfigthis.carbonai.client.model.ModifyUserConfigurationInput;
import com.konfigthis.carbonai.client.model.OrderDirV2;
import com.konfigthis.carbonai.client.model.Pagination;
import com.konfigthis.carbonai.client.model.UpdateUsersInput;
import com.konfigthis.carbonai.client.model.UserListResponse;
import com.konfigthis.carbonai.client.model.UserRequestContent;
import com.konfigthis.carbonai.client.model.UserResponse;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.BeforeAll;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for UsersApi
 */
@Disabled
public class UsersApiTest {

    private static UsersApi api;

    
    @BeforeAll
    public static void beforeClass() {
        ApiClient apiClient = Configuration.getDefaultApiClient();
        api = new UsersApi(apiClient);
    }

    /**
     * Delete Users
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteTest() throws ApiException {
        List<String> customerIds = null;
        GenericSuccessResponse response = api.delete(customerIds)
                .execute();
        // TODO: test validations
    }

    /**
     * User Endpoint
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getTest() throws ApiException {
        String customerId = null;
        UserResponse response = api.get(customerId)
                .execute();
        // TODO: test validations
    }

    /**
     * List Users Endpoint
     *
     * List users within an organization
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listTest() throws ApiException {
        Pagination pagination = null;
        ListUsersFilters filters = null;
        ListUsersOrderByTypes orderBy = null;
        OrderDirV2 orderDir = null;
        Boolean includeCount = null;
        UserListResponse response = api.list()
                .pagination(pagination)
                .filters(filters)
                .orderBy(orderBy)
                .orderDir(orderDir)
                .includeCount(includeCount)
                .execute();
        // TODO: test validations
    }

    /**
     * Toggle User Features
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void toggleUserFeaturesTest() throws ApiException {
        String configurationKeyName = null;
        Object value = null;
        GenericSuccessResponse response = api.toggleUserFeatures(configurationKeyName, value)
                .execute();
        // TODO: test validations
    }

    /**
     * Update Users
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateUsersTest() throws ApiException {
        List<String> customerIds = null;
        Object autoSyncEnabledSources = null;
        Integer maxFiles = null;
        Integer maxFilesPerUpload = null;
        GenericSuccessResponse response = api.updateUsers(customerIds)
                .autoSyncEnabledSources(autoSyncEnabledSources)
                .maxFiles(maxFiles)
                .maxFilesPerUpload(maxFilesPerUpload)
                .execute();
        // TODO: test validations
    }

}
