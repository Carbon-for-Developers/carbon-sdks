# AddDataSourceTagsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | **map[string]interface{}** |  | 
**DataSourceId** | **int32** |  | 

## Methods

### NewAddDataSourceTagsInput

`func NewAddDataSourceTagsInput(tags map[string]interface{}, dataSourceId int32, ) *AddDataSourceTagsInput`

NewAddDataSourceTagsInput instantiates a new AddDataSourceTagsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDataSourceTagsInputWithDefaults

`func NewAddDataSourceTagsInputWithDefaults() *AddDataSourceTagsInput`

NewAddDataSourceTagsInputWithDefaults instantiates a new AddDataSourceTagsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *AddDataSourceTagsInput) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddDataSourceTagsInput) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddDataSourceTagsInput) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.


### GetDataSourceId

`func (o *AddDataSourceTagsInput) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *AddDataSourceTagsInput) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *AddDataSourceTagsInput) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


