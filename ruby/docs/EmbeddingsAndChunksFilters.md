# Carbon::EmbeddingsAndChunksFilters

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **user_file_id** | **Integer** |  |  |
| **embedding_model** | [**EmbeddingGeneratorsNullable**](EmbeddingGeneratorsNullable.md) |  | [optional][default to &#39;OPENAI&#39;] |

## Example

```ruby
require 'carbon'

instance = Carbon::EmbeddingsAndChunksFilters.new(
  user_file_id: null,
  embedding_model: OPENAI
)
```

