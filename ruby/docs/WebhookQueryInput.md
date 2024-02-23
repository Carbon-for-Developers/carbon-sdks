# Carbon::WebhookQueryInput

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **pagination** | [**Pagination**](Pagination.md) |  | [optional] |
| **order_by** | [**WebhookOrderByColumns**](WebhookOrderByColumns.md) |  | [optional] |
| **order_dir** | [**OrderDir**](OrderDir.md) |  | [optional] |
| **filters** | [**WebhookFilters**](WebhookFilters.md) |  | [optional] |

## Example

```ruby
require 'carbon'

instance = Carbon::WebhookQueryInput.new(
  pagination: null,
  order_by: null,
  order_dir: null,
  filters: null
)
```

