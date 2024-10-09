package com.konfigthis.carbonai.client.api;

import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.model.CommentsInput;
import com.konfigthis.carbonai.client.model.CommentsOrderBy;
import com.konfigthis.carbonai.client.model.CommentsResponse;
import com.konfigthis.carbonai.client.model.CommitsInput;
import com.konfigthis.carbonai.client.model.CommitsResponse;
import com.konfigthis.carbonai.client.model.FilesInput;
import com.konfigthis.carbonai.client.model.FilesResponse;
import com.konfigthis.carbonai.client.model.Issue;
import com.konfigthis.carbonai.client.model.IssuesFilter;
import com.konfigthis.carbonai.client.model.IssuesInput;
import com.konfigthis.carbonai.client.model.IssuesOrderBy;
import com.konfigthis.carbonai.client.model.IssuesResponse;
import com.konfigthis.carbonai.client.model.OrderDirV2Nullable;
import com.konfigthis.carbonai.client.model.PROrderBy;
import com.konfigthis.carbonai.client.model.PullRequestExtended;
import com.konfigthis.carbonai.client.model.PullRequestFilters;
import com.konfigthis.carbonai.client.model.PullRequestResponse;
import com.konfigthis.carbonai.client.model.PullRequestsInput;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class GithubApi extends GithubApiGenerated {

    public GithubApi(ApiClient apiClient) throws IllegalArgumentException {
        super(apiClient);
    }
    public class GetIssueRequestBuilder extends GetIssueRequestBuilderGenerated {
        public GetIssueRequestBuilder(Integer issueNumber) {
            super(issueNumber);
        }
    }
    public class GetIssuesRequestBuilder extends GetIssuesRequestBuilderGenerated {
        public GetIssuesRequestBuilder(Integer dataSourceId, String repository) {
            super(dataSourceId, repository);
        }
    }
    public class GetPrRequestBuilder extends GetPrRequestBuilderGenerated {
        public GetPrRequestBuilder(Integer pullNumber) {
            super(pullNumber);
        }
    }
    public class GetPrCommentsRequestBuilder extends GetPrCommentsRequestBuilderGenerated {
        public GetPrCommentsRequestBuilder(Integer dataSourceId, String repository, Integer pullNumber) {
            super(dataSourceId, repository, pullNumber);
        }
    }
    public class GetPrCommitsRequestBuilder extends GetPrCommitsRequestBuilderGenerated {
        public GetPrCommitsRequestBuilder(Integer dataSourceId, String repository, Integer pullNumber) {
            super(dataSourceId, repository, pullNumber);
        }
    }
    public class GetPrFilesRequestBuilder extends GetPrFilesRequestBuilderGenerated {
        public GetPrFilesRequestBuilder(Integer dataSourceId, String repository, Integer pullNumber) {
            super(dataSourceId, repository, pullNumber);
        }
    }
    public class GetPullRequestsRequestBuilder extends GetPullRequestsRequestBuilderGenerated {
        public GetPullRequestsRequestBuilder(Integer dataSourceId, String repository) {
            super(dataSourceId, repository);
        }
    }
}
