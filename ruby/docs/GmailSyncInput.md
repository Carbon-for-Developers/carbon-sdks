# Carbon::GmailSyncInput

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **tags** | **Object** |  | [optional] |
| **filters** | **Object** |  |  |
| **chunk_size** | **Integer** |  | [optional][default to 1500] |
| **chunk_overlap** | **Integer** |  | [optional][default to 20] |
| **skip_embedding_generation** | **Boolean** |  | [optional][default to false] |
| **embedding_model** | [**EmbeddingGenerators**](EmbeddingGenerators.md) |  | [optional][default to &#39;OPENAI&#39;] |
| **generate_sparse_vectors** | **Boolean** |  | [optional][default to false] |
| **prepend_filename_to_chunks** | **Boolean** |  | [optional][default to false] |

## Example

```ruby
require 'carbon'

instance = Carbon::GmailSyncInput.new(
  tags: {},
  filters: null,
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: null,
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false
)
```

