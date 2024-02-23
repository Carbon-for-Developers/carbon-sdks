# Carbon::SyncFilesRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **tags** | **Object** |  | [optional] |
| **data_source_id** | **Integer** |  |  |
| **ids** | **Array&lt;String&gt;** |  |  |
| **chunk_size** | **Integer** |  | [optional][default to 1500] |
| **chunk_overlap** | **Integer** |  | [optional][default to 20] |
| **skip_embedding_generation** | **Boolean** |  | [optional][default to false] |
| **embedding_model** | [**EmbeddingGeneratorsNullable**](EmbeddingGeneratorsNullable.md) |  | [optional][default to &#39;OPENAI&#39;] |
| **generate_sparse_vectors** | **Boolean** |  | [optional][default to false] |
| **prepend_filename_to_chunks** | **Boolean** |  | [optional][default to false] |
| **max_items_per_chunk** | **Integer** |  | [optional] |

## Example

```ruby
require 'carbon'

instance = Carbon::SyncFilesRequest.new(
  tags: {},
  data_source_id: null,
  ids: null,
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: null,
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false,
  max_items_per_chunk: null
)
```

