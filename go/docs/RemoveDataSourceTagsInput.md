# RemoveDataSourceTagsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | **int32** |  | 
**TagsToRemove** | Pointer to **[]string** |  | [optional] [default to []]
**RemoveAllTags** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewRemoveDataSourceTagsInput

`func NewRemoveDataSourceTagsInput(dataSourceId int32, ) *RemoveDataSourceTagsInput`

NewRemoveDataSourceTagsInput instantiates a new RemoveDataSourceTagsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveDataSourceTagsInputWithDefaults

`func NewRemoveDataSourceTagsInputWithDefaults() *RemoveDataSourceTagsInput`

NewRemoveDataSourceTagsInputWithDefaults instantiates a new RemoveDataSourceTagsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *RemoveDataSourceTagsInput) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *RemoveDataSourceTagsInput) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *RemoveDataSourceTagsInput) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.


### GetTagsToRemove

`func (o *RemoveDataSourceTagsInput) GetTagsToRemove() []string`

GetTagsToRemove returns the TagsToRemove field if non-nil, zero value otherwise.

### GetTagsToRemoveOk

`func (o *RemoveDataSourceTagsInput) GetTagsToRemoveOk() (*[]string, bool)`

GetTagsToRemoveOk returns a tuple with the TagsToRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsToRemove

`func (o *RemoveDataSourceTagsInput) SetTagsToRemove(v []string)`

SetTagsToRemove sets TagsToRemove field to given value.

### HasTagsToRemove

`func (o *RemoveDataSourceTagsInput) HasTagsToRemove() bool`

HasTagsToRemove returns a boolean if a field has been set.

### GetRemoveAllTags

`func (o *RemoveDataSourceTagsInput) GetRemoveAllTags() bool`

GetRemoveAllTags returns the RemoveAllTags field if non-nil, zero value otherwise.

### GetRemoveAllTagsOk

`func (o *RemoveDataSourceTagsInput) GetRemoveAllTagsOk() (*bool, bool)`

GetRemoveAllTagsOk returns a tuple with the RemoveAllTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllTags

`func (o *RemoveDataSourceTagsInput) SetRemoveAllTags(v bool)`

SetRemoveAllTags sets RemoveAllTags field to given value.

### HasRemoveAllTags

`func (o *RemoveDataSourceTagsInput) HasRemoveAllTags() bool`

HasRemoveAllTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


