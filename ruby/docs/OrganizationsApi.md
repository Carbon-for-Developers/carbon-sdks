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
configuration = Carbon::Configuration.new
configuration.access_token = 'YOUR API KEY'
carbon = Carbon::Client.new(configuration)


begin
  # Get Organization
  result = carbon.organizations.get
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

