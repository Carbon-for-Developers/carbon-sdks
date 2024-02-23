# Carbon::DataSourcesApi

All URIs are relative to *https://api.carbon.ai*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**query_user_data_sources**](DataSourcesApi.md#query_user_data_sources) | **POST** /user_data_sources | User Data Sources |
| [**revoke_access_token**](DataSourcesApi.md#revoke_access_token) | **POST** /revoke_access_token | Revoke Access Token |

## query_user_data_sources

User Data Sources

### Example

```ruby
require 'carbon'

Carbon. = 'YOUR API KEY'
Carbon. = 'YOUR API KEY'
Carbon. = 'YOUR API KEY'

pagination = {
        "limit" => 10,
        "offset" => 0,
    }
order_by = "created_at"
order_dir = "desc"
filters = {
        "source" => "GOOGLE_DRIVE",
    }

begin
  # User Data Sources
  result = Carbon::DataSources.query_user_data_sources(
                                                          pagination: pagination,
                                                          order_by: order_by,
                                                          order_dir: order_dir,
                                                          filters: filters,
                                                        )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::DataSources.query_user_data_sources: #{e}"
end
```

#### Using the query_user_data_sources_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
pagination = {
        "limit" => 10,
        "offset" => 0,
    }
order_by = "created_at"
order_dir = "desc"
filters = {
        "source" => "GOOGLE_DRIVE",
    }

begin
  # User Data Sources
  data, status_code, headers, response = Carbon::DataSources.query_user_data_sources_with_http_info(
                                                                                                       pagination: pagination,
                                                                                                       order_by: order_by,
                                                                                                       order_dir: order_dir,
                                                                                                       filters: filters,
                                                                                                     )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => OrganizationUserDataSourceResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::DataSources.query_user_data_sources: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **organization_user_data_source_query_input** | [**OrganizationUserDataSourceQueryInput**](OrganizationUserDataSourceQueryInput.md) |  |  |

### Return type

[**OrganizationUserDataSourceResponse**](OrganizationUserDataSourceResponse.md)

## revoke_access_token

Revoke Access Token

### Example

```ruby
require 'carbon'

Carbon. = 'YOUR API KEY'
Carbon. = 'YOUR API KEY'
Carbon. = 'YOUR API KEY'

data_source_id = 1

begin
  # Revoke Access Token
  result = Carbon::DataSources.revoke_access_token(
                                                      data_source_id: data_source_id,
                                                    )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::DataSources.revoke_access_token: #{e}"
end
```

#### Using the revoke_access_token_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
data_source_id = 1

begin
  # Revoke Access Token
  data, status_code, headers, response = Carbon::DataSources.revoke_access_token_with_http_info(
                                                                                                   data_source_id: data_source_id,
                                                                                                 )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => GenericSuccessResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::DataSources.revoke_access_token: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **revoke_access_token_input** | [**RevokeAccessTokenInput**](RevokeAccessTokenInput.md) |  |  |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)

