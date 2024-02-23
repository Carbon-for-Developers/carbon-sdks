# Carbon::OrganizationUserDataSourceAPI

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **Integer** |  |  |
| **data_source_external_id** | **String** |  |  |
| **data_source_type** | [**DataSourceType**](DataSourceType.md) |  |  |
| **sync_status** | [**DataSourceSyncStatuses**](DataSourceSyncStatuses.md) |  |  |
| **source_items_synced_at** | **Time** |  |  |
| **organization_user_id** | **Integer** |  |  |
| **organization_id** | **Integer** |  |  |
| **organization_supplied_user_id** | **String** |  |  |
| **revoked_access** | **Boolean** |  |  |
| **last_synced_at** | **Time** |  |  |
| **last_sync_action** | [**DataSourceLastSyncActions**](DataSourceLastSyncActions.md) |  |  |
| **created_at** | **Time** |  |  |
| **updated_at** | **Time** |  |  |

## Example

```ruby
require 'carbon'

instance = Carbon::OrganizationUserDataSourceAPI.new(
  id: null,
  data_source_external_id: null,
  data_source_type: null,
  sync_status: null,
  source_items_synced_at: null,
  organization_user_id: null,
  organization_id: null,
  organization_supplied_user_id: null,
  revoked_access: null,
  last_synced_at: null,
  last_sync_action: null,
  created_at: null,
  updated_at: null
)
```

