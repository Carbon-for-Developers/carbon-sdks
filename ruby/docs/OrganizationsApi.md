# Carbon::OrganizationsApi

All URIs are relative to *https://api.carbon.ai*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get**](OrganizationsApi.md#get) | **GET** /organization | Get Organization |

## get

Get Organization

### Example

```ruby
require 'carbon'

Carbon. = 'YOUR API KEY'


begin
  # Get Organization
  result = Carbon::Organizations.get
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Organizations.get: #{e}"
end
```

#### Using the get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby

begin
  # Get Organization
  data, status_code, headers, response = Carbon::Organizations.get_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => OrganizationResponse
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Organizations.get: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**OrganizationResponse**](OrganizationResponse.md)

