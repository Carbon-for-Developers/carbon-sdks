# Carbon::DeleteFilesQueryInput

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **file_ids** | **Array&lt;Integer&gt;** |  | [optional] |
| **sync_statuses** | [**Array&lt;ExternalFileSyncStatuses&gt;**](ExternalFileSyncStatuses.md) |  | [optional] |
| **delete_non_synced_only** | **Boolean** |  | [optional][default to false] |

## Example

```ruby
require 'carbon'

instance = Carbon::DeleteFilesQueryInput.new(
  file_ids: null,
  sync_statuses: null,
  delete_non_synced_only: false
)
```

