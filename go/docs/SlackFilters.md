# SlackFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | **string** | Conversation ID from slack, can be obtained from Carbon with slack conversations endpoint | 
**After** | Pointer to **NullableString** | Only messages after this date will be included in results. If not set, all messages          will be included. Should be entered in YYYY/mm/dd format | [optional] 

## Methods

### NewSlackFilters

`func NewSlackFilters(conversationId string, ) *SlackFilters`

NewSlackFilters instantiates a new SlackFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackFiltersWithDefaults

`func NewSlackFiltersWithDefaults() *SlackFilters`

NewSlackFiltersWithDefaults instantiates a new SlackFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *SlackFilters) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *SlackFilters) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *SlackFilters) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetAfter

`func (o *SlackFilters) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *SlackFilters) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *SlackFilters) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *SlackFilters) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *SlackFilters) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *SlackFilters) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


