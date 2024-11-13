package com.konfigthis.carbonai.client.api;

import com.konfigthis.carbonai.client.ApiClient;
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
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class DataSourcesApi extends DataSourcesApiGenerated {

    public DataSourcesApi(ApiClient apiClient) throws IllegalArgumentException {
        super(apiClient);
    }
    public class AddTagsRequestBuilder extends AddTagsRequestBuilderGenerated {
        public AddTagsRequestBuilder(Object tags, Integer dataSourceId) {
            super(tags, dataSourceId);
        }
    }
    public class QueryRequestBuilder extends QueryRequestBuilderGenerated {
        public QueryRequestBuilder() {
            super();
        }
    }
    public class QueryUserDataSourcesRequestBuilder extends QueryUserDataSourcesRequestBuilderGenerated {
        public QueryUserDataSourcesRequestBuilder() {
            super();
        }
    }
    public class RemoveTagsRequestBuilder extends RemoveTagsRequestBuilderGenerated {
        public RemoveTagsRequestBuilder(Integer dataSourceId) {
            super(dataSourceId);
        }
    }
    public class RevokeAccessTokenRequestBuilder extends RevokeAccessTokenRequestBuilderGenerated {
        public RevokeAccessTokenRequestBuilder(Integer dataSourceId) {
            super(dataSourceId);
        }
    }
}
