# Carbon::ChunkPropertiesNullable

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **set_page_as_boundary** | **Boolean** |  | [default to false] |
| **prepend_filename_to_chunks** | **Boolean** |  | [default to false] |
| **max_items_per_chunk** | **Integer** |  |  |

## Example

```ruby
require 'carbon'

instance = Carbon::ChunkPropertiesNullable.new(
  set_page_as_boundary: false,
  prepend_filename_to_chunks: false,
  max_items_per_chunk: null
)
```

