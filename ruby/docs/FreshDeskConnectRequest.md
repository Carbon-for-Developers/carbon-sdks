# Carbon::FreshDeskConnectRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **tags** | **Object** |  | [optional] |
| **domain** | **String** |  |  |
| **api_key** | **String** |  |  |
| **chunk_size** | **Integer** |  | [optional][default to 1500] |
| **chunk_overlap** | **Integer** |  | [optional][default to 20] |
| **skip_embedding_generation** | **Boolean** |  | [optional][default to false] |
| **embedding_model** | [**EmbeddingGeneratorsNullable**](EmbeddingGeneratorsNullable.md) |  | [optional][default to &#39;OPENAI&#39;] |
| **generate_sparse_vectors** | **Boolean** |  | [optional][default to false] |
| **prepend_filename_to_chunks** | **Boolean** |  | [optional][default to false] |

## Example

```ruby
require 'carbon'

instance = Carbon::FreshDeskConnectRequest.new(
  tags: {},
  domain: null,
  api_key: null,
  chunk_size: 1500,
  chunk_overlap: 20,
  skip_embedding_generation: false,
  embedding_model: OPENAI,
  generate_sparse_vectors: false,
  prepend_filename_to_chunks: false
)
```

