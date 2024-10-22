package com.konfigthis.carbonai.client.api;

import com.konfigthis.carbonai.client.ApiClient;
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
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class UsersApi extends UsersApiGenerated {

    public UsersApi(ApiClient apiClient) throws IllegalArgumentException {
        super(apiClient);
    }
    public class DeleteRequestBuilder extends DeleteRequestBuilderGenerated {
        public DeleteRequestBuilder(List<String> customerIds) {
            super(customerIds);
        }
    }
    public class GetRequestBuilder extends GetRequestBuilderGenerated {
        public GetRequestBuilder(String customerId) {
            super(customerId);
        }
    }
    public class ListRequestBuilder extends ListRequestBuilderGenerated {
        public ListRequestBuilder() {
            super();
        }
    }
    public class ToggleUserFeaturesRequestBuilder extends ToggleUserFeaturesRequestBuilderGenerated {
        public ToggleUserFeaturesRequestBuilder(String configurationKeyName, Object value) {
            super(configurationKeyName, value);
        }
    }
    public class UpdateUsersRequestBuilder extends UpdateUsersRequestBuilderGenerated {
        public UpdateUsersRequestBuilder(List<String> customerIds) {
            super(customerIds);
        }
    }
    public class WhoAmIRequestBuilder extends WhoAmIRequestBuilderGenerated {
        public WhoAmIRequestBuilder() {
            super();
        }
    }
}
