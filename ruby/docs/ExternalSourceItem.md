# Carbon::ExternalSourceItem

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **external_id** | **String** |  |  |
| **source** | [**DataSourceType**](DataSourceType.md) |  |  |
| **name** | **String** |  |  |
| **synced_at** | **Time** |  |  |
| **is_selectable** | **Boolean** |  |  |
| **is_expandable** | **Boolean** |  |  |
| **organization_id** | **Integer** |  |  |
| **organization_supplied_user_id** | **String** |  |  |
| **organization_user_id** | **Integer** |  |  |
| **organization_user_data_source_id** | **Integer** |  |  |
| **organization_user_file_to_sync_id** | **Integer** |  |  |
| **parent_external_id** | **String** |  |  |
| **item_type** | **String** |  |  |
| **created_at** | **Time** |  |  |
| **updated_at** | **Time** |  |  |

## Example

```ruby
require 'carbon'

instance = Carbon::ExternalSourceItem.new(
  id: null,
  external_id: null,
  source: null,
  name: null,
  synced_at: null,
  is_selectable: null,
  is_expandable: null,
  organization_id: null,
  organization_supplied_user_id: null,
  organization_user_id: null,
  organization_user_data_source_id: null,
  organization_user_file_to_sync_id: null,
  parent_external_id: null,
  item_type: null,
  created_at: null,
  updated_at: null
)
```

