# Carbon::ResyncFileQueryInput

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **file_id** | **Integer** |  |  |
| **chunk_size** | **Integer** |  | [optional] |
| **chunk_overlap** | **Integer** |  | [optional] |

## Example

```ruby
require 'carbon'

instance = Carbon::ResyncFileQueryInput.new(
  file_id: null,
  chunk_size: null,
  chunk_overlap: null
)
```

