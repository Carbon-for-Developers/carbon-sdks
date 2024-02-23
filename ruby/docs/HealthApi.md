# Carbon::HealthApi

All URIs are relative to *https://api.carbon.ai*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**check**](HealthApi.md#check) | **GET** /health | Health |

## check

Health

### Example

```ruby
require 'carbon'
configuration = Carbon::Configuration.new
carbon = Carbon::Client.new(configuration)


begin
  # Health
  result = carbon.health.check
  p result
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Health.check: #{e}"
end
```

#### Using the check_with_http_info variant

This returns an Array which contains the response data, status code and headers.

```ruby

begin
  # Health
  data, status_code, headers, response = Carbon::Health.check_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Object
rescue Carbon::ApiError => e
  puts "Exception when calling Carbon::Health.check: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

**Object**

