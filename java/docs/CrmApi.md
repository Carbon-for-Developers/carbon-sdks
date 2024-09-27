# CrmApi

All URIs are relative to *https://api.carbon.ai*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getAccount**](CrmApi.md#getAccount) | **GET** /integrations/data/crm/accounts/{id} | Get Account |
| [**getAccounts**](CrmApi.md#getAccounts) | **POST** /integrations/data/crm/accounts | Get Accounts |
| [**getContact**](CrmApi.md#getContact) | **GET** /integrations/data/crm/contacts/{id} | Get Contact |
| [**getContacts**](CrmApi.md#getContacts) | **POST** /integrations/data/crm/contacts | Get Contacts |
| [**getLead**](CrmApi.md#getLead) | **GET** /integrations/data/crm/leads/{id} | Get Lead |
| [**getLeads**](CrmApi.md#getLeads) | **POST** /integrations/data/crm/leads | Get Leads |
| [**getOpportunities**](CrmApi.md#getOpportunities) | **POST** /integrations/data/crm/opportunities | Get Opportunities |
| [**getOpportunity**](CrmApi.md#getOpportunity) | **GET** /integrations/data/crm/opportunities/{id} | Get Opportunity |


<a name="getAccount"></a>
# **getAccount**
> Account getAccount(id, dataSourceId).includeRemoteData(includeRemoteData).includes(includes).execute();

Get Account

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.CrmApi;
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
    String id = "id_example";
    Integer dataSourceId = 56;
    Boolean includeRemoteData = false;
    List<BaseIncludes> includes = Arrays.asList();
    try {
      Account result = client
              .crm
              .getAccount(id, dataSourceId)
              .includeRemoteData(includeRemoteData)
              .includes(includes)
              .execute();
      System.out.println(result);
      System.out.println(result.getDescription());
      System.out.println(result.getId());
      System.out.println(result.getOwner());
      System.out.println(result.getName());
      System.out.println(result.getIndustry());
      System.out.println(result.getWebsite());
      System.out.println(result.getNumberOfEmployees());
      System.out.println(result.getAddresses());
      System.out.println(result.getPhoneNumbers());
      System.out.println(result.getLastActivityAt());
      System.out.println(result.getCreatedAt());
      System.out.println(result.getUpdatedAt());
      System.out.println(result.getIsDeleted());
      System.out.println(result.getTasks());
      System.out.println(result.getEvents());
      System.out.println(result.getRemoteData());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getAccount");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<Account> response = client
              .crm
              .getAccount(id, dataSourceId)
              .includeRemoteData(includeRemoteData)
              .includes(includes)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getAccount");
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
| **id** | **String**|  | |
| **dataSourceId** | **Integer**|  | |
| **includeRemoteData** | **Boolean**|  | [optional] [default to false] |
| **includes** | [**List&lt;BaseIncludes&gt;**](BaseIncludes.md)|  | [optional] |

### Return type

[**Account**](Account.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="getAccounts"></a>
# **getAccounts**
> AccountResponse getAccounts(accountsRequest).execute();

Get Accounts

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.CrmApi;
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
    Boolean includeRemoteData = false;
    String nextCursor = "nextCursor_example";
    Integer pageSize = 56;
    OrderDirV2Nullable orderDir = OrderDirV2Nullable.fromValue("asc");
    List<BaseIncludes> includes = Arrays.asList();
    AccountFilters filters = new AccountFilters();
    AccountsOrderByNullable orderBy = AccountsOrderByNullable.fromValue("created_at");
    try {
      AccountResponse result = client
              .crm
              .getAccounts(dataSourceId)
              .includeRemoteData(includeRemoteData)
              .nextCursor(nextCursor)
              .pageSize(pageSize)
              .orderDir(orderDir)
              .includes(includes)
              .filters(filters)
              .orderBy(orderBy)
              .execute();
      System.out.println(result);
      System.out.println(result.getCount());
      System.out.println(result.getNextCursor());
      System.out.println(result.getData());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getAccounts");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<AccountResponse> response = client
              .crm
              .getAccounts(dataSourceId)
              .includeRemoteData(includeRemoteData)
              .nextCursor(nextCursor)
              .pageSize(pageSize)
              .orderDir(orderDir)
              .includes(includes)
              .filters(filters)
              .orderBy(orderBy)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getAccounts");
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
| **accountsRequest** | [**AccountsRequest**](AccountsRequest.md)|  | |

### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="getContact"></a>
# **getContact**
> Contact getContact(id, dataSourceId).includeRemoteData(includeRemoteData).includes(includes).execute();

Get Contact

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.CrmApi;
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
    String id = "id_example";
    Integer dataSourceId = 56;
    Boolean includeRemoteData = false;
    List<BaseIncludes> includes = Arrays.asList();
    try {
      Contact result = client
              .crm
              .getContact(id, dataSourceId)
              .includeRemoteData(includeRemoteData)
              .includes(includes)
              .execute();
      System.out.println(result);
      System.out.println(result.getTitle());
      System.out.println(result.getDescription());
      System.out.println(result.getId());
      System.out.println(result.getOwner());
      System.out.println(result.getFirstName());
      System.out.println(result.getLastName());
      System.out.println(result.getName());
      System.out.println(result.getDepartment());
      System.out.println(result.getAddresses());
      System.out.println(result.getPhoneNumbers());
      System.out.println(result.getEmails());
      System.out.println(result.getAccount());
      System.out.println(result.getLastActivityAt());
      System.out.println(result.getCreatedAt());
      System.out.println(result.getUpdatedAt());
      System.out.println(result.getIsDeleted());
      System.out.println(result.getTasks());
      System.out.println(result.getEvents());
      System.out.println(result.getRemoteData());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getContact");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<Contact> response = client
              .crm
              .getContact(id, dataSourceId)
              .includeRemoteData(includeRemoteData)
              .includes(includes)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getContact");
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
| **id** | **String**|  | |
| **dataSourceId** | **Integer**|  | |
| **includeRemoteData** | **Boolean**|  | [optional] [default to false] |
| **includes** | [**List&lt;BaseIncludes&gt;**](BaseIncludes.md)|  | [optional] |

### Return type

[**Contact**](Contact.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="getContacts"></a>
# **getContacts**
> ContactsResponse getContacts(contactsRequest).execute();

Get Contacts

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.CrmApi;
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
    Boolean includeRemoteData = false;
    String nextCursor = "nextCursor_example";
    Integer pageSize = 56;
    OrderDirV2Nullable orderDir = OrderDirV2Nullable.fromValue("asc");
    List<BaseIncludes> includes = Arrays.asList();
    ContactFilters filters = new ContactFilters();
    ContactsOrderByNullable orderBy = ContactsOrderByNullable.fromValue("created_at");
    try {
      ContactsResponse result = client
              .crm
              .getContacts(dataSourceId)
              .includeRemoteData(includeRemoteData)
              .nextCursor(nextCursor)
              .pageSize(pageSize)
              .orderDir(orderDir)
              .includes(includes)
              .filters(filters)
              .orderBy(orderBy)
              .execute();
      System.out.println(result);
      System.out.println(result.getCount());
      System.out.println(result.getNextCursor());
      System.out.println(result.getData());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getContacts");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<ContactsResponse> response = client
              .crm
              .getContacts(dataSourceId)
              .includeRemoteData(includeRemoteData)
              .nextCursor(nextCursor)
              .pageSize(pageSize)
              .orderDir(orderDir)
              .includes(includes)
              .filters(filters)
              .orderBy(orderBy)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getContacts");
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
| **contactsRequest** | [**ContactsRequest**](ContactsRequest.md)|  | |

### Return type

[**ContactsResponse**](ContactsResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="getLead"></a>
# **getLead**
> Lead getLead(id, dataSourceId).includeRemoteData(includeRemoteData).includes(includes).execute();

Get Lead

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.CrmApi;
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
    String id = "id_example";
    Integer dataSourceId = 56;
    Boolean includeRemoteData = false;
    List<BaseIncludes> includes = Arrays.asList();
    try {
      Lead result = client
              .crm
              .getLead(id, dataSourceId)
              .includeRemoteData(includeRemoteData)
              .includes(includes)
              .execute();
      System.out.println(result);
      System.out.println(result.getTitle());
      System.out.println(result.getDescription());
      System.out.println(result.getId());
      System.out.println(result.getOwner());
      System.out.println(result.getSource());
      System.out.println(result.getStatus());
      System.out.println(result.getCompany());
      System.out.println(result.getFirstName());
      System.out.println(result.getLastName());
      System.out.println(result.getAddresses());
      System.out.println(result.getPhoneNumbers());
      System.out.println(result.getEmails());
      System.out.println(result.getConvertedAt());
      System.out.println(result.getConvertedAccount());
      System.out.println(result.getConvertedContact());
      System.out.println(result.getLastActivityAt());
      System.out.println(result.getCreatedAt());
      System.out.println(result.getUpdatedAt());
      System.out.println(result.getIsDeleted());
      System.out.println(result.getTasks());
      System.out.println(result.getEvents());
      System.out.println(result.getRemoteData());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getLead");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<Lead> response = client
              .crm
              .getLead(id, dataSourceId)
              .includeRemoteData(includeRemoteData)
              .includes(includes)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getLead");
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
| **id** | **String**|  | |
| **dataSourceId** | **Integer**|  | |
| **includeRemoteData** | **Boolean**|  | [optional] [default to false] |
| **includes** | [**List&lt;BaseIncludes&gt;**](BaseIncludes.md)|  | [optional] |

### Return type

[**Lead**](Lead.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="getLeads"></a>
# **getLeads**
> LeadsResponse getLeads(leadsRequest).execute();

Get Leads

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.CrmApi;
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
    Boolean includeRemoteData = false;
    String nextCursor = "nextCursor_example";
    Integer pageSize = 56;
    OrderDirV2Nullable orderDir = OrderDirV2Nullable.fromValue("asc");
    List<BaseIncludes> includes = Arrays.asList();
    LeadFilters filters = new LeadFilters();
    LeadsOrderByNullable orderBy = LeadsOrderByNullable.fromValue("created_at");
    try {
      LeadsResponse result = client
              .crm
              .getLeads(dataSourceId)
              .includeRemoteData(includeRemoteData)
              .nextCursor(nextCursor)
              .pageSize(pageSize)
              .orderDir(orderDir)
              .includes(includes)
              .filters(filters)
              .orderBy(orderBy)
              .execute();
      System.out.println(result);
      System.out.println(result.getCount());
      System.out.println(result.getNextCursor());
      System.out.println(result.getData());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getLeads");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<LeadsResponse> response = client
              .crm
              .getLeads(dataSourceId)
              .includeRemoteData(includeRemoteData)
              .nextCursor(nextCursor)
              .pageSize(pageSize)
              .orderDir(orderDir)
              .includes(includes)
              .filters(filters)
              .orderBy(orderBy)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getLeads");
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
| **leadsRequest** | [**LeadsRequest**](LeadsRequest.md)|  | |

### Return type

[**LeadsResponse**](LeadsResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="getOpportunities"></a>
# **getOpportunities**
> OpportunitiesResponse getOpportunities(opportunitiesRequest).execute();

Get Opportunities

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.CrmApi;
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
    Boolean includeRemoteData = false;
    String nextCursor = "nextCursor_example";
    Integer pageSize = 56;
    OrderDirV2Nullable orderDir = OrderDirV2Nullable.fromValue("asc");
    List<BaseIncludes> includes = Arrays.asList();
    OpportunityFilters filters = new OpportunityFilters();
    OpportunitiesOrderByNullable orderBy = OpportunitiesOrderByNullable.fromValue("created_at");
    try {
      OpportunitiesResponse result = client
              .crm
              .getOpportunities(dataSourceId)
              .includeRemoteData(includeRemoteData)
              .nextCursor(nextCursor)
              .pageSize(pageSize)
              .orderDir(orderDir)
              .includes(includes)
              .filters(filters)
              .orderBy(orderBy)
              .execute();
      System.out.println(result);
      System.out.println(result.getCount());
      System.out.println(result.getNextCursor());
      System.out.println(result.getData());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getOpportunities");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<OpportunitiesResponse> response = client
              .crm
              .getOpportunities(dataSourceId)
              .includeRemoteData(includeRemoteData)
              .nextCursor(nextCursor)
              .pageSize(pageSize)
              .orderDir(orderDir)
              .includes(includes)
              .filters(filters)
              .orderBy(orderBy)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getOpportunities");
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
| **opportunitiesRequest** | [**OpportunitiesRequest**](OpportunitiesRequest.md)|  | |

### Return type

[**OpportunitiesResponse**](OpportunitiesResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

<a name="getOpportunity"></a>
# **getOpportunity**
> Opportunity getOpportunity(id, dataSourceId).includeRemoteData(includeRemoteData).includes(includes).execute();

Get Opportunity

### Example
```java
import com.konfigthis.carbonai.client.ApiClient;
import com.konfigthis.carbonai.client.ApiException;
import com.konfigthis.carbonai.client.ApiResponse;
import com.konfigthis.carbonai.client.Carbon;
import com.konfigthis.carbonai.client.Configuration;
import com.konfigthis.carbonai.client.auth.*;
import com.konfigthis.carbonai.client.model.*;
import com.konfigthis.carbonai.client.api.CrmApi;
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
    String id = "id_example";
    Integer dataSourceId = 56;
    Boolean includeRemoteData = false;
    List<BaseIncludes> includes = Arrays.asList();
    try {
      Opportunity result = client
              .crm
              .getOpportunity(id, dataSourceId)
              .includeRemoteData(includeRemoteData)
              .includes(includes)
              .execute();
      System.out.println(result);
      System.out.println(result.getDescription());
      System.out.println(result.getId());
      System.out.println(result.getOwner());
      System.out.println(result.getName());
      System.out.println(result.getAmount());
      System.out.println(result.getAccount());
      System.out.println(result.getContact());
      System.out.println(result.getStage());
      System.out.println(result.getStatus());
      System.out.println(result.getCloseDate());
      System.out.println(result.getLastActivityAt());
      System.out.println(result.getCreatedAt());
      System.out.println(result.getUpdatedAt());
      System.out.println(result.getIsDeleted());
      System.out.println(result.getTasks());
      System.out.println(result.getEvents());
      System.out.println(result.getRemoteData());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getOpportunity");
      System.err.println("Status code: " + e.getStatusCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }

    // Use .executeWithHttpInfo() to retrieve HTTP Status Code, Headers and Request
    try {
      ApiResponse<Opportunity> response = client
              .crm
              .getOpportunity(id, dataSourceId)
              .includeRemoteData(includeRemoteData)
              .includes(includes)
              .executeWithHttpInfo();
      System.out.println(response.getResponseBody());
      System.out.println(response.getResponseHeaders());
      System.out.println(response.getStatusCode());
      System.out.println(response.getRoundTripTime());
      System.out.println(response.getRequest());
    } catch (ApiException e) {
      System.err.println("Exception when calling CrmApi#getOpportunity");
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
| **id** | **String**|  | |
| **dataSourceId** | **Integer**|  | |
| **includeRemoteData** | **Boolean**|  | [optional] [default to false] |
| **includes** | [**List&lt;BaseIncludes&gt;**](BaseIncludes.md)|  | [optional] |

### Return type

[**Opportunity**](Opportunity.md)

### Authorization

[accessToken](../README.md#accessToken), [apiKey](../README.md#apiKey), [customerId](../README.md#customerId)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |

