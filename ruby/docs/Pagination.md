# Carbon::Pagination

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **limit** | **Integer** |  | [optional][default to 10] |
| **offset** | **Integer** |  | [optional][default to 0] |

## Example

```ruby
require 'carbon'

instance = Carbon::Pagination.new(
  limit: 10,
  offset: 0
)
```

