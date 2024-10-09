# GithubApi

All URIs are relative to *https://api.carbon.ai*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getIssue**](GithubApi.md#getIssue) | **GET** /integrations/data/github/issues/{issue_number} | Issue |
| [**getIssues**](GithubApi.md#getIssues) | **POST** /integrations/data/github/issues | Issues |
| [**getPr**](GithubApi.md#getPr) | **GET** /integrations/data/github/pull_requests/{pull_number} | Get Pr |
| [**getPrComments**](GithubApi.md#getPrComments) | **POST** /integrations/data/github/pull_requests/comments | Pr Comments |
| [**getPrCommits**](GithubApi.md#getPrCommits) | **POST** /integrations/data/github/pull_requests/commits | Pr Commits |
| [**getPrFiles**](GithubApi.md#getPrFiles) | **POST** /integrations/data/github/pull_requests/files | Pr Files |
| [**getPullRequests**](GithubApi.md#getPullRequests) | **POST** /integrations/data/github/pull_requests | Get Prs |


<a name="getIssue"></a>
# **getIssue**
> Issue getIssue(issueNumber).includeRemoteData(includeRemoteData).dataSourceId(dataSourceId).repository(repository).execute();

Issue

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.GithubApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "https://api.carbon.ai";
    
    configuration.accessToken  = "YOUR API KEY";
    
    configuration.apiKey  = "YOUR API KEY";
    
    configuration.customerId  = "YOUR API KEY";
    Carbon client = new Carbon(configuration);
    Integer issueNumber = 56;
    Boolean includeRemoteData = false;
    Integer dataSourceId = 56;
    String repository = "repository_example";
    try {
      Issue result = client
              .github
              .getIssue(issueNumber)
              .includeRemoteData(includeRemoteData)
              .dataSourceId(dataSourceId)
              .repository(repository)
              .execute();
      System.out.println(result);
      System.out.println(result.getTitle());
      System.out.println(result.getId());
      System.out.println(result.getUrl());
      System.out.println(result.getHtmlUrl());
      System.out.println(result.getNumber());
      System.out.println(result.getUser());
      System.out.println(result.getLabels());
      System.out.println(result.getState());
      System.out.println(result.getLocked());
      System.out.println(result.getNumComments());
      System.out.println(result.getCreatedAt());
      System.out.println(result.getUpdatedAt());
      System.out.println(result.getClosedAt());
      System.out.println(result.getDraft());
      System.out.println(result.getPullRequest());
      System.out.println(result.getBody());
      System.out.println(result.getClosedBy());
    } catch (ApiException e) {
      System.err.println("Exception when calling GithubApi#getIssue");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<Issue> response = client
              .github
              .getIssue(issueNumber)
              .includeRemoteData(includeRemoteData)
              .dataSourceId(dataSourceId)
              .repository(repository)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling GithubApi#getIssue");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}

```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **issueNumber** | **Integer**|  | |
| **includeRemoteData** | **Boolean**|  | [optional] [default to false] |
| **dataSourceId** | **Integer**|  | [optional] |
| **repository** | **String**|  | [optional] |

### Return type

[**Issue**](Issue.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="getIssues"></a>
# **getIssues**
> IssuesResponse getIssues(issuesInput).execute();

Issues

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.GithubApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "https://api.carbon.ai";
    
    configuration.accessToken  = "YOUR API KEY";
    
    configuration.apiKey  = "YOUR API KEY";
    
    configuration.customerId  = "YOUR API KEY";
    Carbon client = new Carbon(configuration);
    Integer dataSourceId = 56;
    String repository = "repository_example"; // Full name of the repository, denoted as {owner}/{repo}
    Boolean includeRemoteData = false;
    Integer page = 1;
    Integer pageSize = 30;
    String nextCursor = "nextCursor_example";
    IssuesFilter filters = new IssuesFilter();
    IssuesOrderBy orderBy = IssuesOrderBy.fromValue("created");
    OrderDirV2Nullable orderDir = OrderDirV2Nullable.fromValue("asc");
    try {
      IssuesResponse result = client
              .github
              .getIssues(dataSourceId, repository)
              .includeRemoteData(includeRemoteData)
              .page(page)
              .pageSize(pageSize)
              .nextCursor(nextCursor)
              .filters(filters)
              .orderBy(orderBy)
              .orderDir(orderDir)
              .execute();
      System.out.println(result);
      System.out.println(result.getData());
      System.out.println(result.getNextCursor());
    } catch (ApiException e) {
      System.err.println("Exception when calling GithubApi#getIssues");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<IssuesResponse> response = client
              .github
              .getIssues(dataSourceId, repository)
              .includeRemoteData(includeRemoteData)
              .page(page)
              .pageSize(pageSize)
              .nextCursor(nextCursor)
              .filters(filters)
              .orderBy(orderBy)
              .orderDir(orderDir)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling GithubApi#getIssues");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}

```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **issuesInput** | [**IssuesInput**](IssuesInput.md)|  | |

### Return type

[**IssuesResponse**](IssuesResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="getPr"></a>
# **getPr**
> PullRequestExtended getPr(pullNumber).includeRemoteData(includeRemoteData).dataSourceId(dataSourceId).repository(repository).execute();

Get Pr

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.GithubApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "https://api.carbon.ai";
    
    configuration.accessToken  = "YOUR API KEY";
    
    configuration.apiKey  = "YOUR API KEY";
    
    configuration.customerId  = "YOUR API KEY";
    Carbon client = new Carbon(configuration);
    Integer pullNumber = 56;
    Boolean includeRemoteData = false;
    Integer dataSourceId = 56;
    String repository = "repository_example";
    try {
      PullRequestExtended result = client
              .github
              .getPr(pullNumber)
              .includeRemoteData(includeRemoteData)
              .dataSourceId(dataSourceId)
              .repository(repository)
              .execute();
      System.out.println(result);
      System.out.println(result.getTitle());
      System.out.println(result.getId());
      System.out.println(result.getUrl());
      System.out.println(result.getNumber());
      System.out.println(result.getState());
      System.out.println(result.getUser());
      System.out.println(result.getCreatedAt());
      System.out.println(result.getUpdatedAt());
      System.out.println(result.getClosedAt());
      System.out.println(result.getMergedAt());
      System.out.println(result.getRequestedReviewers());
      System.out.println(result.getRequestedTeams());
      System.out.println(result.getLabels());
      System.out.println(result.getDraft());
      System.out.println(result.getHead());
      System.out.println(result.getBase());
      System.out.println(result.getRemoteData());
      System.out.println(result.getMerged());
      System.out.println(result.getNumComments());
      System.out.println(result.getNumReviewComments());
      System.out.println(result.getNumCommits());
      System.out.println(result.getNumAdditions());
      System.out.println(result.getNumDeletions());
      System.out.println(result.getNumChangedFiles());
    } catch (ApiException e) {
      System.err.println("Exception when calling GithubApi#getPr");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<PullRequestExtended> response = client
              .github
              .getPr(pullNumber)
              .includeRemoteData(includeRemoteData)
              .dataSourceId(dataSourceId)
              .repository(repository)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling GithubApi#getPr");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}

```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **pullNumber** | **Integer**|  | |
| **includeRemoteData** | **Boolean**|  | [optional] [default to false] |
| **dataSourceId** | **Integer**|  | [optional] |
| **repository** | **String**|  | [optional] |

### Return type

[**PullRequestExtended**](PullRequestExtended.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="getPrComments"></a>
# **getPrComments**
> CommentsResponse getPrComments(commentsInput).execute();

Pr Comments

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.GithubApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "https://api.carbon.ai";
    
    configuration.accessToken  = "YOUR API KEY";
    
    configuration.apiKey  = "YOUR API KEY";
    
    configuration.customerId  = "YOUR API KEY";
    Carbon client = new Carbon(configuration);
    Integer dataSourceId = 56;
    String repository = "repository_example"; // Full name of the repository, denoted as {owner}/{repo}
    Integer pullNumber = 56;
    Boolean includeRemoteData = false;
    Integer page = 1;
    Integer pageSize = 30;
    String nextCursor = "nextCursor_example";
    CommentsOrderBy orderBy = CommentsOrderBy.fromValue("created");
    OrderDirV2Nullable orderDir = OrderDirV2Nullable.fromValue("asc");
    try {
      CommentsResponse result = client
              .github
              .getPrComments(dataSourceId, repository, pullNumber)
              .includeRemoteData(includeRemoteData)
              .page(page)
              .pageSize(pageSize)
              .nextCursor(nextCursor)
              .orderBy(orderBy)
              .orderDir(orderDir)
              .execute();
      System.out.println(result);
      System.out.println(result.getData());
      System.out.println(result.getNextCursor());
    } catch (ApiException e) {
      System.err.println("Exception when calling GithubApi#getPrComments");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<CommentsResponse> response = client
              .github
              .getPrComments(dataSourceId, repository, pullNumber)
              .includeRemoteData(includeRemoteData)
              .page(page)
              .pageSize(pageSize)
              .nextCursor(nextCursor)
              .orderBy(orderBy)
              .orderDir(orderDir)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling GithubApi#getPrComments");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}

```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **commentsInput** | [**CommentsInput**](CommentsInput.md)|  | |

### Return type

[**CommentsResponse**](CommentsResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="getPrCommits"></a>
# **getPrCommits**
> CommitsResponse getPrCommits(commitsInput).execute();

Pr Commits

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.GithubApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "https://api.carbon.ai";
    
    configuration.accessToken  = "YOUR API KEY";
    
    configuration.apiKey  = "YOUR API KEY";
    
    configuration.customerId  = "YOUR API KEY";
    Carbon client = new Carbon(configuration);
    Integer dataSourceId = 56;
    String repository = "repository_example"; // Full name of the repository, denoted as {owner}/{repo}
    Integer pullNumber = 56;
    Boolean includeRemoteData = false;
    Integer page = 1;
    Integer pageSize = 30;
    String nextCursor = "nextCursor_example";
    try {
      CommitsResponse result = client
              .github
              .getPrCommits(dataSourceId, repository, pullNumber)
              .includeRemoteData(includeRemoteData)
              .page(page)
              .pageSize(pageSize)
              .nextCursor(nextCursor)
              .execute();
      System.out.println(result);
      System.out.println(result.getData());
      System.out.println(result.getNextCursor());
    } catch (ApiException e) {
      System.err.println("Exception when calling GithubApi#getPrCommits");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<CommitsResponse> response = client
              .github
              .getPrCommits(dataSourceId, repository, pullNumber)
              .includeRemoteData(includeRemoteData)
              .page(page)
              .pageSize(pageSize)
              .nextCursor(nextCursor)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling GithubApi#getPrCommits");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}

```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **commitsInput** | [**CommitsInput**](CommitsInput.md)|  | |

### Return type

[**CommitsResponse**](CommitsResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="getPrFiles"></a>
# **getPrFiles**
> FilesResponse getPrFiles(filesInput).execute();

Pr Files

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.GithubApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "https://api.carbon.ai";
    
    configuration.accessToken  = "YOUR API KEY";
    
    configuration.apiKey  = "YOUR API KEY";
    
    configuration.customerId  = "YOUR API KEY";
    Carbon client = new Carbon(configuration);
    Integer dataSourceId = 56;
    String repository = "repository_example"; // Full name of the repository, denoted as {owner}/{repo}
    Integer pullNumber = 56;
    Boolean includeRemoteData = false;
    Integer page = 1;
    Integer pageSize = 30;
    String nextCursor = "nextCursor_example";
    try {
      FilesResponse result = client
              .github
              .getPrFiles(dataSourceId, repository, pullNumber)
              .includeRemoteData(includeRemoteData)
              .page(page)
              .pageSize(pageSize)
              .nextCursor(nextCursor)
              .execute();
      System.out.println(result);
      System.out.println(result.getData());
      System.out.println(result.getNextCursor());
    } catch (ApiException e) {
      System.err.println("Exception when calling GithubApi#getPrFiles");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<FilesResponse> response = client
              .github
              .getPrFiles(dataSourceId, repository, pullNumber)
              .includeRemoteData(includeRemoteData)
              .page(page)
              .pageSize(pageSize)
              .nextCursor(nextCursor)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling GithubApi#getPrFiles");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}

```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **filesInput** | [**FilesInput**](FilesInput.md)|  | |

### Return type

[**FilesResponse**](FilesResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="getPullRequests"></a>
# **getPullRequests**
> PullRequestResponse getPullRequests(pullRequestsInput).execute();

Get Prs

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.GithubApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "https://api.carbon.ai";
    
    configuration.accessToken  = "YOUR API KEY";
    
    configuration.apiKey  = "YOUR API KEY";
    
    configuration.customerId  = "YOUR API KEY";
    Carbon client = new Carbon(configuration);
    Integer dataSourceId = 56;
    String repository = "repository_example"; // Full name of the repository, denoted as {owner}/{repo}
    Boolean includeRemoteData = false;
    Integer page = 1;
    Integer pageSize = 30;
    String nextCursor = "nextCursor_example";
    PullRequestFilters filters = new PullRequestFilters();
    PROrderBy orderBy = PROrderBy.fromValue("created");
    OrderDirV2Nullable orderDir = OrderDirV2Nullable.fromValue("asc");
    try {
      PullRequestResponse result = client
              .github
              .getPullRequests(dataSourceId, repository)
              .includeRemoteData(includeRemoteData)
              .page(page)
              .pageSize(pageSize)
              .nextCursor(nextCursor)
              .filters(filters)
              .orderBy(orderBy)
              .orderDir(orderDir)
              .execute();
      System.out.println(result);
      System.out.println(result.getData());
      System.out.println(result.getNextCursor());
    } catch (ApiException e) {
      System.err.println("Exception when calling GithubApi#getPullRequests");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<PullRequestResponse> response = client
              .github
              .getPullRequests(dataSourceId, repository)
              .includeRemoteData(includeRemoteData)
              .page(page)
              .pageSize(pageSize)
              .nextCursor(nextCursor)
              .filters(filters)
              .orderBy(orderBy)
              .orderDir(orderDir)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling GithubApi#getPullRequests");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}

```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **pullRequestsInput** | [**PullRequestsInput**](PullRequestsInput.md)|  | |

### Return type

[**PullRequestResponse**](PullRequestResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

