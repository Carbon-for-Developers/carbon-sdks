# Carbon::OrganizationUserFilesToSyncFilters

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **tags** | [**Hash&lt;String, Tags1&gt;**](Tags1.md) |  | [optional] |
| **source** | [**SourceProperty**](SourceProperty.md) |  | [optional] |
| **name** | **String** |  | [optional] |
| **tags_v2** | **Object** |  | [optional] |
| **ids** | **Array&lt;Integer&gt;** |  | [optional] |
| **external_file_ids** | **Array&lt;String&gt;** |  | [optional] |
| **sync_statuses** | [**Array&lt;ExternalFileSyncStatuses&gt;**](ExternalFileSyncStatuses.md) |  | [optional] |
| **parent_file_ids** | **Array&lt;Integer&gt;** |  | [optional] |
| **organization_user_data_source_id** | **Array&lt;Integer&gt;** |  | [optional] |
| **embedding_generators** | [**Array&lt;EmbeddingGenerators&gt;**](EmbeddingGenerators.md) |  | [optional] |

## Example

```ruby
require 'carbon'

instance = Carbon::OrganizationUserFilesToSyncFilters.new(
  tags: null,
  source: null,
  name: null,
  tags_v2: null,
  ids: null,
  external_file_ids: null,
  sync_statuses: null,
  parent_file_ids: null,
  organization_user_data_source_id: null,
  embedding_generators: null
)
```

