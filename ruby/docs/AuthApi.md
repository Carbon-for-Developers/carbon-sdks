# Carbon::AuthApi

All URIs are relative to *https://api.carbon.ai*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get_access_token**](AuthApi.md#get_access_token) | **GET** /auth/v1/access_token | Get Access Token |
| [**get_white_labeling**](AuthApi.md#get_white_labeling) | **GET** /auth/v1/white_labeling | Get White Labeling |

## get_access_token

Get Access Token

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
configuration.api_key = 'YOUR API KEY'
configuration.customer_id = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)


begin
  # Get Access Token
  result = carbon.auth.get_access_token
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Auth.get_access_token: #{e}"
end
```

#### Using the get_access_token_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby

begin
  # Get Access Token
  data, status_code, headers, response = Carbon::Auth.get_access_token_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => TokenResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Auth.get_access_token: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**TokenResponse**](TokenResponse.md)

## get_white_labeling

Get White Labeling

Returns whether or not the organization is white labeled and which integrations are white labeled  :param current_user: the current user :param db: the database session :return: a WhiteLabelingResponse

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)


begin
  # Get White Labeling
  result = carbon.auth.get_white_labeling
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Auth.get_white_labeling: #{e}"
end
```

#### Using the get_white_labeling_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby

begin
  # Get White Labeling
  data, status_code, headers, response = Carbon::Auth.get_white_labeling_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => WhiteLabelingResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Auth.get_white_labeling: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**WhiteLabelingResponse**](WhiteLabelingResponse.md)

