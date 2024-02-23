# Carbon::UsersApi

All URIs are relative to *https://api.carbon.ai*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get**](UsersApi.md#get) | **POST** /user | User Endpoint |
| [**toggle_user_features**](UsersApi.md#toggle_user_features) | **POST** /modify_user_configuration | Toggle User Features |

## get

User Endpoint

### Example

```ruby
require 'carbon'

Carbon. = 'YOUR API KEY'

customer_id = "string_example"

begin
  # User Endpoint
  result = Carbon::Users.get(
                                customer_id: customer_id,
                              )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Users.get: #{e}"
end
```

#### Using the get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
customer_id = "string_example"

begin
  # User Endpoint
  data, status_code, headers, response = Carbon::Users.get_with_http_info(
                                                                             customer_id: customer_id,
                                                                           )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => UserResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Users.get: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **user_request_content** | [**UserRequestContent**](UserRequestContent.md) |  |  |

### Return type

[**UserResponse**](UserResponse.md)

## toggle_user_features

Toggle User Features

### Example

```ruby
require 'carbon'

Carbon. = 'YOUR API KEY'
Carbon. = 'YOUR API KEY'
Carbon. = 'YOUR API KEY'

configuration_key_name = "string_example"
value = {}

begin
  # Toggle User Features
  result = Carbon::Users.toggle_user_features(
                                                 configuration_key_name: configuration_key_name,
                                                 value: value,
                                               )
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Users.toggle_user_features: #{e}"
end
```

#### Using the toggle_user_features_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby
configuration_key_name = "string_example"
value = {}

begin
  # Toggle User Features
  data, status_code, headers, response = Carbon::Users.toggle_user_features_with_http_info(
                                                                                              configuration_key_name: configuration_key_name,
                                                                                              value: value,
                                                                                            )
  p status_code # => 2xx
  p headers # => { ... }
  p data # => GenericSuccessResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Users.toggle_user_features: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **modify_user_configuration_input** | [**ModifyUserConfigurationInput**](ModifyUserConfigurationInput.md) |  |  |

### Return type

[**GenericSuccessResponse**](GenericSuccessResponse.md)

