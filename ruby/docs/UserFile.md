# Carbon::UserFile

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **tags** | **Object** |  |  |
| **id** | **Integer** |  |  |
| **source** | [**DataSourceType**](DataSourceType.md) |  |  |
| **organization_id** | **Integer** |  |  |
| **organization_supplied_user_id** | **String** |  |  |
| **organization_user_data_source_id** | **Integer** |  |  |
| **external_file_id** | **String** |  |  |
| **external_url** | **String** |  |  |
| **sync_status** | [**ExternalFileSyncStatuses**](ExternalFileSyncStatuses.md) |  |  |
| **sync_error_message** | **String** |  |  |
| **last_sync** | **Time** |  |  |
| **file_statistics** | [**FileStatisticsNullable**](FileStatisticsNullable.md) |  |  |
| **file_metadata** | **Object** |  |  |
| **embedding_properties** | [**Hash&lt;String, EmbeddingProperties&gt;**](EmbeddingProperties.md) |  |  |
| **chunk_size** | **Integer** |  |  |
| **chunk_overlap** | **Integer** |  |  |
| **chunk_properties** | [**ChunkPropertiesNullable**](ChunkPropertiesNullable.md) |  |  |
| **name** | **String** |  |  |
| **parent_id** | **Integer** |  |  |
| **enable_auto_sync** | **Boolean** |  |  |
| **presigned_url** | **String** |  |  |
| **parsed_text_url** | **String** |  |  |
| **additional_presigned_urls** | **Object** |  |  |
| **skip_embedding_generation** | **Boolean** |  |  |
| **source_created_at** | **Time** |  |  |
| **generate_sparse_vectors** | **Boolean** |  |  |
| **created_at** | **Time** |  |  |
| **updated_at** | **Time** |  |  |

## Example

```ruby
require 'carbon'

instance = Carbon::UserFile.new(
  tags: null,
  id: null,
  source: null,
  organization_id: null,
  organization_supplied_user_id: null,
  organization_user_data_source_id: null,
  external_file_id: null,
  external_url: null,
  sync_status: null,
  sync_error_message: null,
  last_sync: null,
  file_statistics: null,
  file_metadata: null,
  embedding_properties: null,
  chunk_size: null,
  chunk_overlap: null,
  chunk_properties: null,
  name: null,
  parent_id: null,
  enable_auto_sync: null,
  presigned_url: null,
  parsed_text_url: null,
  additional_presigned_urls: null,
  skip_embedding_generation: null,
  source_created_at: null,
  generate_sparse_vectors: null,
  created_at: null,
  updated_at: null
)
```

