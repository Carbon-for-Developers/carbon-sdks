# UsersApi

All URIs are relative to *https://api.carbon.ai*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**delete**](UsersApi.md#delete) | **POST** /delete_users | Delete Users |
| [**get**](UsersApi.md#get) | **POST** /user | User Endpoint |
| [**list**](UsersApi.md#list) | **POST** /list_users | List Users Endpoint |
| [**toggleUserFeatures**](UsersApi.md#toggleUserFeatures) | **POST** /modify_user_configuration | Toggle User Features |
| [**updateUsers**](UsersApi.md#updateUsers) | **POST** /update_users | Update Users |
| [**whoAmI**](UsersApi.md#whoAmI) | **GET** /whoami | Me Endpoint |


<a name="delete"></a>
# **delete**
> GenericSuccessResponse delete(deleteUsersInput).execute();

Delete Users

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.UsersApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "https://api.carbon.ai";
    
    configuration.apiKey  = "YOUR API KEY";
    Carbon client = new Carbon(configuration);
    List<String> customerIds = Arrays.asList();
    try {
      GenericSuccessResponse result = client
              .users
              .delete(customerIds)
              .execute();
      System.out.println(result);
      System.out.println(result.getSuccess());
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#delete");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<GenericSuccessResponse> response = client
              .users
              .delete(customerIds)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#delete");
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
| **deleteUsersInput** | [**DeleteUsersInput**](DeleteUsersInput.md)|  | |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="get"></a>
# **get**
> UserResponse get(userRequestContent).execute();

User Endpoint

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.UsersApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "https://api.carbon.ai";
    
    configuration.apiKey  = "YOUR API KEY";
    Carbon client = new Carbon(configuration);
    String customerId = "customerId_example";
    try {
      UserResponse result = client
              .users
              .get(customerId)
              .execute();
      System.out.println(result);
      System.out.println(result.getId());
      System.out.println(result.getOrganizationId());
      System.out.println(result.getOrganizationSuppliedUserId());
      System.out.println(result.getCreatedAt());
      System.out.println(result.getUpdatedAt());
      System.out.println(result.getDeletedAt());
      System.out.println(result.getNumFilesSynced());
      System.out.println(result.getNumCharactersSynced());
      System.out.println(result.getNumTokensSynced());
      System.out.println(result.getAggregateFileSize());
      System.out.println(result.getAggregateNumCharacters());
      System.out.println(result.getAggregateNumTokens());
      System.out.println(result.getAggregateNumEmbeddings());
      System.out.println(result.getAggregateNumFilesBySource());
      System.out.println(result.getAggregateNumFilesByFileFormat());
      System.out.println(result.getUniqueFileTags());
      System.out.println(result.getEnabledFeatures());
      System.out.println(result.getCustomLimits());
      System.out.println(result.getAutoSyncEnabledSources());
      System.out.println(result.getConnectorSettings());
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#get");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<UserResponse> response = client
              .users
              .get(customerId)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#get");
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
| **userRequestContent** | [**UserRequestContent**](UserRequestContent.md)|  | |

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="list"></a>
# **list**
> UserListResponse list(listUsersRequest).execute();

List Users Endpoint

List users within an organization

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.UsersApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "https://api.carbon.ai";
    
    configuration.apiKey  = "YOUR API KEY";
    Carbon client = new Carbon(configuration);
    Pagination pagination = new Pagination();
    ListUsersFilters filters = new ListUsersFilters();
    ListUsersOrderByTypes orderBy = ListUsersOrderByTypes.fromValue("created_at");
    OrderDirV2 orderDir = OrderDirV2.fromValue("asc");
    Boolean includeCount = false;
    try {
      UserListResponse result = client
              .users
              .list()
              .pagination(pagination)
              .filters(filters)
              .orderBy(orderBy)
              .orderDir(orderDir)
              .includeCount(includeCount)
              .execute();
      System.out.println(result);
      System.out.println(result.getUsers());
      System.out.println(result.getCount());
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#list");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<UserListResponse> response = client
              .users
              .list()
              .pagination(pagination)
              .filters(filters)
              .orderBy(orderBy)
              .orderDir(orderDir)
              .includeCount(includeCount)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#list");
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
| **listUsersRequest** | [**ListUsersRequest**](ListUsersRequest.md)|  | |

### Return type

[**UserListResponse**](UserListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="toggleUserFeatures"></a>
# **toggleUserFeatures**
> GenericSuccessResponse toggleUserFeatures(modifyUserConfigurationInput).execute();

Toggle User Features

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.UsersApi;
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
    ConfigurationKeys configurationKeyName = ConfigurationKeys.fromValue("sparse_vectors");
    Object value = null;
    try {
      GenericSuccessResponse result = client
              .users
              .toggleUserFeatures(configurationKeyName, value)
              .execute();
      System.out.println(result);
      System.out.println(result.getSuccess());
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#toggleUserFeatures");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<GenericSuccessResponse> response = client
              .users
              .toggleUserFeatures(configurationKeyName, value)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#toggleUserFeatures");
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
| **modifyUserConfigurationInput** | [**ModifyUserConfigurationInput**](ModifyUserConfigurationInput.md)|  | |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="updateUsers"></a>
# **updateUsers**
> GenericSuccessResponse updateUsers(updateUsersInput).execute();

Update Users

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.UsersApi;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public class Example {
  public static void main(String[] args) {
    Configuration configuration = new Configuration();
    configuration.host = "https://api.carbon.ai";
    
    configuration.apiKey  = "YOUR API KEY";
    Carbon client = new Carbon(configuration);
    List<String> customerIds = Arrays.asList(); // List of organization supplied user IDs
    Object autoSyncEnabledSources = null;
    Integer maxFiles = 56; // Custom file upload limit for the user over *all* user's files across all uploads.          If set, then the user will not be allowed to upload more files than this limit. If not set, or if set to -1,         then the user will have no limit.
    Integer maxFilesPerUpload = 56; // Custom file upload limit for the user across a single upload.         If set, then the user will not be allowed to upload more files than this limit in a single upload. If not set,         or if set to -1, then the user will have no limit.
    Integer maxCharacters = 56; // Custom character upload limit for the user over *all* user's files across all uploads.          If set, then the user will not be allowed to upload more characters than this limit. If not set, or if set to -1,         then the user will have no limit.
    Integer maxCharactersPerFile = 56; // A single file upload from the user can not exceed this character limit.         If set, then the file will not be synced if it exceeds this limit. If not set, or if set to -1, then the          user will have no limit.
    Integer maxCharactersPerUpload = 56; // Custom character upload limit for the user across a single upload.         If set, then the user won't be able to sync more than this many characters in one upload.          If not set, or if set to -1, then the user will have no limit.
    Integer autoSyncInterval = 56; // The interval in hours at which the user's data sources should be synced. If not set or set to -1,          the user will be synced at the organization level interval or default interval if that is also not set.          Must be one of [3, 6, 12, 24]
    try {
      GenericSuccessResponse result = client
              .users
              .updateUsers(customerIds)
              .autoSyncEnabledSources(autoSyncEnabledSources)
              .maxFiles(maxFiles)
              .maxFilesPerUpload(maxFilesPerUpload)
              .maxCharacters(maxCharacters)
              .maxCharactersPerFile(maxCharactersPerFile)
              .maxCharactersPerUpload(maxCharactersPerUpload)
              .autoSyncInterval(autoSyncInterval)
              .execute();
      System.out.println(result);
      System.out.println(result.getSuccess());
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#updateUsers");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<GenericSuccessResponse> response = client
              .users
              .updateUsers(customerIds)
              .autoSyncEnabledSources(autoSyncEnabledSources)
              .maxFiles(maxFiles)
              .maxFilesPerUpload(maxFilesPerUpload)
              .maxCharacters(maxCharacters)
              .maxCharactersPerFile(maxCharactersPerFile)
              .maxCharactersPerUpload(maxCharactersPerUpload)
              .autoSyncInterval(autoSyncInterval)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#updateUsers");
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
| **updateUsersInput** | [**UpdateUsersInput**](UpdateUsersInput.md)|  | |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="whoAmI"></a>
# **whoAmI**
> UserResponse whoAmI().execute();

Me Endpoint

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.UsersApi;
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
    try {
      UserResponse result = client
              .users
              .whoAmI()
              .execute();
      System.out.println(result);
      System.out.println(result.getId());
      System.out.println(result.getOrganizationId());
      System.out.println(result.getOrganizationSuppliedUserId());
      System.out.println(result.getCreatedAt());
      System.out.println(result.getUpdatedAt());
      System.out.println(result.getDeletedAt());
      System.out.println(result.getNumFilesSynced());
      System.out.println(result.getNumCharactersSynced());
      System.out.println(result.getNumTokensSynced());
      System.out.println(result.getAggregateFileSize());
      System.out.println(result.getAggregateNumCharacters());
      System.out.println(result.getAggregateNumTokens());
      System.out.println(result.getAggregateNumEmbeddings());
      System.out.println(result.getAggregateNumFilesBySource());
      System.out.println(result.getAggregateNumFilesByFileFormat());
      System.out.println(result.getUniqueFileTags());
      System.out.println(result.getEnabledFeatures());
      System.out.println(result.getCustomLimits());
      System.out.println(result.getAutoSyncEnabledSources());
      System.out.println(result.getConnectorSettings());
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#whoAmI");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<UserResponse> response = client
              .users
              .whoAmI()
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#whoAmI");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}

```

### Parameters
This endpoint does not need any parameter.

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

