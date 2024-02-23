# Carbon::OrganizationUserDataSourceFilters

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **source** | [**DataSourceTypeNullable**](DataSourceTypeNullable.md) |  | [optional] |
| **ids** | **Array&lt;Integer&gt;** |  | [optional] |
| **revoked_access** | **Boolean** |  | [optional] |

## Example

```ruby
require 'carbon'

instance = Carbon::OrganizationUserDataSourceFilters.new(
  source: null,
  ids: null,
  revoked_access: null
)
```

