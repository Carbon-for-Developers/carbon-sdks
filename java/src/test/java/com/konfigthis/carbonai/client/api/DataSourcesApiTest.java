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
import com.konfigthis.carbonai.client.model.AddDataSourceTagsInput;
import com.konfigthis.carbonai.client.model.GenericSuccessResponse;
import com.konfigthis.carbonai.client.model.OrderDir;
import com.konfigthis.carbonai.client.model.OrganizationUserDataSourceAPI;
import com.konfigthis.carbonai.client.model.OrganizationUserDataSourceFilters;
import com.konfigthis.carbonai.client.model.OrganizationUserDataSourceOrderByColumns;
import com.konfigthis.carbonai.client.model.OrganizationUserDataSourceQueryInput;
import com.konfigthis.carbonai.client.model.OrganizationUserDataSourceResponse;
import com.konfigthis.carbonai.client.model.Pagination;
import com.konfigthis.carbonai.client.model.RemoveDataSourceTagsInput;
import com.konfigthis.carbonai.client.model.RevokeAccessTokenInput;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.BeforeAll;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for DataSourcesApi
 */
@Disabled
public class DataSourcesApiTest {

    private static DataSourcesApi api;

    
    @BeforeAll
    public static void beforeClass() {
        ApiClient apiClient = Configuration.getDefaultApiClient();
        api = new DataSourcesApi(apiClient);
    }

    /**
     * Add Data Source Tags
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void addTagsTest() throws ApiException {
        Object tags = null;
        Integer dataSourceId = null;
        OrganizationUserDataSourceAPI response = api.addTags(tags, dataSourceId)
                .execute();
        // TODO: test validations
    }

    /**
     * Data Sources
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void queryTest() throws ApiException {
        Pagination pagination = null;
        OrganizationUserDataSourceOrderByColumns orderBy = null;
        OrderDir orderDir = null;
        OrganizationUserDataSourceFilters filters = null;
        OrganizationUserDataSourceResponse response = api.query()
                .pagination(pagination)
                .orderBy(orderBy)
                .orderDir(orderDir)
                .filters(filters)
                .execute();
        // TODO: test validations
    }

    /**
     * User Data Sources
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void queryUserDataSourcesTest() throws ApiException {
        Pagination pagination = null;
        OrganizationUserDataSourceOrderByColumns orderBy = null;
        OrderDir orderDir = null;
        OrganizationUserDataSourceFilters filters = null;
        OrganizationUserDataSourceResponse response = api.queryUserDataSources()
                .pagination(pagination)
                .orderBy(orderBy)
                .orderDir(orderDir)
                .filters(filters)
                .execute();
        // TODO: test validations
    }

    /**
     * Remove Data Source Tags
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void removeTagsTest() throws ApiException {
        Integer dataSourceId = null;
        List<String> tagsToRemove = null;
        Boolean removeAllTags = null;
        OrganizationUserDataSourceAPI response = api.removeTags(dataSourceId)
                .tagsToRemove(tagsToRemove)
                .removeAllTags(removeAllTags)
                .execute();
        // TODO: test validations
    }

    /**
     * Revoke Access Token
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void revokeAccessTokenTest() throws ApiException {
        Integer dataSourceId = null;
        GenericSuccessResponse response = api.revokeAccessToken(dataSourceId)
                .execute();
        // TODO: test validations
    }

}
