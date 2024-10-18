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
import com.konfigthis.carbonai.client.model.DeleteWhiteLabelRequest;
import com.konfigthis.carbonai.client.model.GDOSWhiteLabelInput;
import com.konfigthis.carbonai.client.model.GDOSWhiteLabelInput1;
import com.konfigthis.carbonai.client.model.GoogleDriveCredentialsProperty;
import com.konfigthis.carbonai.client.model.ListWhiteLabelRequest;
import com.konfigthis.carbonai.client.model.OrderDir;
import com.konfigthis.carbonai.client.model.Pagination;
import com.konfigthis.carbonai.client.model.WhiteLabelFilters;
import com.konfigthis.carbonai.client.model.WhiteLabelOrderByColumns;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.BeforeAll;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for WhiteLabelApi
 */
@Disabled
public class WhiteLabelApiTest {

    private static WhiteLabelApi api;

    
    @BeforeAll
    public static void beforeClass() {
        ApiClient apiClient = Configuration.getDefaultApiClient();
        api = new WhiteLabelApi(apiClient);
    }

    /**
     * Create White Labels
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createTest() throws ApiException {
        Object response = api.create()
                .execute();
        // TODO: test validations
    }

    /**
     * Delete White Labels
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteTest() throws ApiException {
        List<Integer> ids = null;
        Object response = api.delete(ids)
                .execute();
        // TODO: test validations
    }

    /**
     * List White Labels
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listTest() throws ApiException {
        Pagination pagination = null;
        WhiteLabelOrderByColumns orderBy = null;
        OrderDir orderDir = null;
        WhiteLabelFilters filters = null;
        Object response = api.list()
                .pagination(pagination)
                .orderBy(orderBy)
                .orderDir(orderDir)
                .filters(filters)
                .execute();
        // TODO: test validations
    }

    /**
     * Update White Label
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateTest() throws ApiException {
        Object dataSourceType = null;
        GoogleDriveCredentialsProperty credentials = null;
        Object response = api.update()
                .dataSourceType(dataSourceType)
                .credentials(credentials)
                .execute();
        // TODO: test validations
    }

}
