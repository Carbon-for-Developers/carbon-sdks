# Carbon::DirectoryItem

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  |  |
| **name** | **String** |  |  |
| **is_synced** | **Boolean** |  | [default to false] |
| **has_children** | **Boolean** |  | [default to false] |

## Example

```ruby
require 'carbon'

instance = Carbon::DirectoryItem.new(
  id: null,
  name: null,
  is_synced: false,
  has_children: false
)
```

