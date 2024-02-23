# Carbon::RawTextInput

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **contents** | **String** |  |  |
| **name** | **String** |  | [optional] |
| **chunk_size** | **Integer** |  | [optional] |
| **chunk_overlap** | **Integer** |  | [optional] |
| **skip_embedding_generation** | **Boolean** |  | [optional][default to false] |
| **overwrite_file_id** | **Integer** |  | [optional] |
| **embedding_model** | [**EmbeddingGeneratorsNullable**](EmbeddingGeneratorsNullable.md) |  | [optional][default to &#39;OPENAI&#39;] |
| **generate_sparse_vectors** | **Boolean** |  | [optional][default to false] |

## Example

```ruby
require 'carbon'

instance = Carbon::RawTextInput.new(
  contents: null,
  name: null,
  chunk_size: null,
  chunk_overlap: null,
  skip_embedding_generation: false,
  overwrite_file_id: null,
  embedding_model: null,
  generate_sparse_vectors: false
)
```

