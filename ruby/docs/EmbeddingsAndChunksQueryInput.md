# Carbon::EmbeddingsAndChunksQueryInput

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **pagination** | [**Pagination**](Pagination.md) |  | [optional] |
| **order_by** | [**EmbeddingsAndChunksOrderByColumns**](EmbeddingsAndChunksOrderByColumns.md) |  | [optional] |
| **order_dir** | [**OrderDir**](OrderDir.md) |  | [optional] |
| **filters** | [**EmbeddingsAndChunksFilters**](EmbeddingsAndChunksFilters.md) |  |  |
| **include_vectors** | **Boolean** |  | [optional][default to false] |

## Example

```ruby
require 'carbon'

instance = Carbon::EmbeddingsAndChunksQueryInput.new(
  pagination: null,
  order_by: null,
  order_dir: null,
  filters: null,
  include_vectors: false
)
```

