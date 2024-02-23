# Carbon::WebhooksApi

All URIs are relative to *https://api.carbon.ai*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**add_url**](WebhooksApi.md#add_url) | **POST** /add_webhook | Add Webhook Url |
| [**delete_url**](WebhooksApi.md#delete_url) | **DELETE** /delete_webhook/{webhook_id} | Delete Webhook Url |
| [**urls**](WebhooksApi.md#urls) | **POST** /webhooks | Webhook Urls |

## add_url

Add Webhook Url

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'

url = "string_example"

begin
  # Add Webhook Url
  result = Carbon::Webhooks.add_url(
                                       url: url,
                                     )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Webhooks.add_url: #{e}"
end
```

#### Using the add_url_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
url = "string_example"

begin
  # Add Webhook Url
  data, status_code, headers, response = Carbon::Webhooks.add_url_with_http_info(
                                                                                    url: url,
                                                                                  )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Webhook
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Webhooks.add_url: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **add_webhook_props** | [**AddWebhookProps**](AddWebhookProps.md) |  |  |

### Return type

[**Webhook**](Webhook.md)

## delete_url

Delete Webhook Url

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'

webhook_id = 1

begin
  # Delete Webhook Url
  result = Carbon::Webhooks.delete_url(
                                          webhook_id: webhook_id,
                                        )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Webhooks.delete_url: #{e}"
end
```

#### Using the delete_url_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
webhook_id = 1

begin
  # Delete Webhook Url
  data, status_code, headers, response = Carbon::Webhooks.delete_url_with_http_info(
                                                                                       webhook_id: webhook_id,
                                                                                     )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => GenericSuccessResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Webhooks.delete_url: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **webhook_id** | **Integer** |  |  |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)

## urls

Webhook Urls

### Example

```ruby
require 'carbon'

Carbon.access_token = 'YOUR API KEY'

pagination = {
        "limit" => 10,
        "offset" => 0,
    }
order_by = "created_at"
order_dir = "desc"
filters = {
        "ids" => [],
    }

begin
  # Webhook Urls
  result = Carbon::Webhooks.urls(
                                    pagination: pagination,
                                    order_by: order_by,
                                    order_dir: order_dir,
                                    filters: filters,
                                  )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Webhooks.urls: #{e}"
end
```

#### Using the urls_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
pagination = {
        "limit" => 10,
        "offset" => 0,
    }
order_by = "created_at"
order_dir = "desc"
filters = {
        "ids" => [],
    }

begin
  # Webhook Urls
  data, status_code, headers, response = Carbon::Webhooks.urls_with_http_info(
                                                                                 pagination: pagination,
                                                                                 order_by: order_by,
                                                                                 order_dir: order_dir,
                                                                                 filters: filters,
                                                                               )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => WebhookQueryResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Webhooks.urls: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **webhook_query_input** | [**WebhookQueryInput**](WebhookQueryInput.md) |  |  |

### Return type

[**WebhookQueryResponse**](WebhookQueryResponse.md)

