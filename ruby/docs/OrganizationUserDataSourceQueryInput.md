# Carbon::OrganizationUserDataSourceQueryInput

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **pagination** | [**Pagination**](Pagination.md) |  | [optional] |
| **order_by** | [**OrganizationUserDataSourceOrderByColumns**](OrganizationUserDataSourceOrderByColumns.md) |  | [optional][default to &#39;updated_at&#39;] |
| **order_dir** | [**OrderDir**](OrderDir.md) |  | [optional][default to &#39;asc&#39;] |
| **filters** | [**OrganizationUserDataSourceFilters**](OrganizationUserDataSourceFilters.md) |  | [optional] |

## Example

```ruby
require 'carbon'

instance = Carbon::OrganizationUserDataSourceQueryInput.new(
  pagination: null,
  order_by: null,
  order_dir: null,
  filters: null
)
```

