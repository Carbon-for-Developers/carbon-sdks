package com.konfigthis.carbonai.client.api;

import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.model.DeleteWhiteLabelRequest;
import com.konfigthis.carbonai.client.model.GDOSWhiteLabelInput;
import com.konfigthis.carbonai.client.model.GDOSWhiteLabelInput1;
import com.konfigthis.carbonai.client.model.GoogleDriveCredentialsProperty;
import com.konfigthis.carbonai.client.model.ListWhiteLabelRequest;
import com.konfigthis.carbonai.client.model.OrderDir;
import com.konfigthis.carbonai.client.model.Pagination;
import com.konfigthis.carbonai.client.model.WhiteLabelFilters;
import com.konfigthis.carbonai.client.model.WhiteLabelOrderByColumns;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class WhiteLabelApi extends WhiteLabelApiGenerated {

    public WhiteLabelApi(ApiClient apiClient) throws IllegalArgumentException {
        super(apiClient);
    }
    public class CreateRequestBuilder extends CreateRequestBuilderGenerated {
        public CreateRequestBuilder() {
            super();
        }
    }
    public class DeleteRequestBuilder extends DeleteRequestBuilderGenerated {
        public DeleteRequestBuilder(List<Integer> ids) {
            super(ids);
        }
    }
    public class ListRequestBuilder extends ListRequestBuilderGenerated {
        public ListRequestBuilder() {
            super();
        }
    }
    public class UpdateRequestBuilder extends UpdateRequestBuilderGenerated {
        public UpdateRequestBuilder() {
            super();
        }
    }
}
