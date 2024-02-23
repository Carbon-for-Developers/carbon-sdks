# Carbon::ListDataSourceItemsRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **data_source_id** | **Integer** |  |  |
| **parent_id** | **String** |  | [optional] |
| **pagination** | [**Pagination**](Pagination.md) |  | [optional] |

## Example

```ruby
require 'carbon'

instance = Carbon::ListDataSourceItemsRequest.new(
  data_source_id: null,
  parent_id: null,
  pagination: null
)
```

