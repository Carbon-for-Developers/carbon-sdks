package com.konfigthis.carbonai.client.api;

import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.model.Account;
import com.konfigthis.carbonai.client.model.AccountFilters;
import com.konfigthis.carbonai.client.model.AccountResponse;
import com.konfigthis.carbonai.client.model.AccountsOrderByNullable;
import com.konfigthis.carbonai.client.model.AccountsRequest;
import com.konfigthis.carbonai.client.model.BaseIncludes;
import com.konfigthis.carbonai.client.model.Contact;
import com.konfigthis.carbonai.client.model.ContactFilters;
import com.konfigthis.carbonai.client.model.ContactsOrderByNullable;
import com.konfigthis.carbonai.client.model.ContactsRequest;
import com.konfigthis.carbonai.client.model.ContactsResponse;
import com.konfigthis.carbonai.client.model.Lead;
import com.konfigthis.carbonai.client.model.LeadFilters;
import com.konfigthis.carbonai.client.model.LeadsOrderByNullable;
import com.konfigthis.carbonai.client.model.LeadsRequest;
import com.konfigthis.carbonai.client.model.LeadsResponse;
import com.konfigthis.carbonai.client.model.OpportunitiesOrderByNullable;
import com.konfigthis.carbonai.client.model.OpportunitiesRequest;
import com.konfigthis.carbonai.client.model.OpportunitiesResponse;
import com.konfigthis.carbonai.client.model.Opportunity;
import com.konfigthis.carbonai.client.model.OpportunityFilters;
import com.konfigthis.carbonai.client.model.OrderDirV2Nullable;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class CrmApi extends CrmApiGenerated {

    public CrmApi(ApiClient apiClient) throws IllegalArgumentException {
        super(apiClient);
    }
    public class GetAccountRequestBuilder extends GetAccountRequestBuilderGenerated {
        public GetAccountRequestBuilder(String id, Integer dataSourceId) {
            super(id, dataSourceId);
        }
    }
    public class GetAccountsRequestBuilder extends GetAccountsRequestBuilderGenerated {
        public GetAccountsRequestBuilder(Integer dataSourceId) {
            super(dataSourceId);
        }
    }
    public class GetContactRequestBuilder extends GetContactRequestBuilderGenerated {
        public GetContactRequestBuilder(String id, Integer dataSourceId) {
            super(id, dataSourceId);
        }
    }
    public class GetContactsRequestBuilder extends GetContactsRequestBuilderGenerated {
        public GetContactsRequestBuilder(Integer dataSourceId) {
            super(dataSourceId);
        }
    }
    public class GetLeadRequestBuilder extends GetLeadRequestBuilderGenerated {
        public GetLeadRequestBuilder(String id, Integer dataSourceId) {
            super(id, dataSourceId);
        }
    }
    public class GetLeadsRequestBuilder extends GetLeadsRequestBuilderGenerated {
        public GetLeadsRequestBuilder(Integer dataSourceId) {
            super(dataSourceId);
        }
    }
    public class GetOpportunitiesRequestBuilder extends GetOpportunitiesRequestBuilderGenerated {
        public GetOpportunitiesRequestBuilder(Integer dataSourceId) {
            super(dataSourceId);
        }
    }
    public class GetOpportunityRequestBuilder extends GetOpportunityRequestBuilderGenerated {
        public GetOpportunityRequestBuilder(String id, Integer dataSourceId) {
            super(id, dataSourceId);
        }
    }
}
