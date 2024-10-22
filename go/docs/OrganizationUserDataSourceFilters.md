# OrganizationUserDataSourceFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **map[string]interface{}** |          Tags to filter by. Supports logical AND and OR operations. Input should be like below:         {             \&quot;OR\&quot;: [                 {                 \&quot;key\&quot;: \&quot;subject\&quot;,                 \&quot;value\&quot;: \&quot;holy-bible\&quot;,                 \&quot;negate\&quot;: false                 },                 {                     \&quot;key\&quot;: \&quot;person-of-interest\&quot;,                     \&quot;value\&quot;: \&quot;jesus christ\&quot;,                     \&quot;negate\&quot;: false                 },                 {                     \&quot;key\&quot;: \&quot;genre\&quot;,                     \&quot;value\&quot;: \&quot;fiction\&quot;,                     \&quot;negate\&quot;: true                 }                 {                     \&quot;AND\&quot;: [                         {                             \&quot;key\&quot;: \&quot;subject\&quot;,                             \&quot;value\&quot;: \&quot;tao-te-ching\&quot;,                             \&quot;negate\&quot;: true                         },                         {                             \&quot;key\&quot;: \&quot;author\&quot;,                             \&quot;value\&quot;: \&quot;lao-tzu\&quot;,                             \&quot;negate\&quot;: false                         }                     ]                 }             ]         }         For a single filter, the filter block can be placed within either an \&quot;AND\&quot; or \&quot;OR\&quot; block.          | [optional] 
**Source** | Pointer to [**NullableDataSourceTypeNullable**](DataSourceTypeNullable.md) |  | [optional] 
**Ids** | Pointer to **[]int32** |  | [optional] 
**RevokedAccess** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewOrganizationUserDataSourceFilters

`func NewOrganizationUserDataSourceFilters() *OrganizationUserDataSourceFilters`

NewOrganizationUserDataSourceFilters instantiates a new OrganizationUserDataSourceFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationUserDataSourceFiltersWithDefaults

`func NewOrganizationUserDataSourceFiltersWithDefaults() *OrganizationUserDataSourceFilters`

NewOrganizationUserDataSourceFiltersWithDefaults instantiates a new OrganizationUserDataSourceFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *OrganizationUserDataSourceFilters) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrganizationUserDataSourceFilters) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrganizationUserDataSourceFilters) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrganizationUserDataSourceFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *OrganizationUserDataSourceFilters) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *OrganizationUserDataSourceFilters) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetSource

`func (o *OrganizationUserDataSourceFilters) GetSource() DataSourceTypeNullable`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OrganizationUserDataSourceFilters) GetSourceOk() (*DataSourceTypeNullable, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OrganizationUserDataSourceFilters) SetSource(v DataSourceTypeNullable)`

SetSource sets Source field to given value.

### HasSource

`func (o *OrganizationUserDataSourceFilters) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *OrganizationUserDataSourceFilters) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *OrganizationUserDataSourceFilters) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetIds

`func (o *OrganizationUserDataSourceFilters) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *OrganizationUserDataSourceFilters) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *OrganizationUserDataSourceFilters) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *OrganizationUserDataSourceFilters) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *OrganizationUserDataSourceFilters) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *OrganizationUserDataSourceFilters) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetRevokedAccess

`func (o *OrganizationUserDataSourceFilters) GetRevokedAccess() bool`

GetRevokedAccess returns the RevokedAccess field if non-nil, zero value otherwise.

### GetRevokedAccessOk

`func (o *OrganizationUserDataSourceFilters) GetRevokedAccessOk() (*bool, bool)`

GetRevokedAccessOk returns a tuple with the RevokedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAccess

`func (o *OrganizationUserDataSourceFilters) SetRevokedAccess(v bool)`

SetRevokedAccess sets RevokedAccess field to given value.

### HasRevokedAccess

`func (o *OrganizationUserDataSourceFilters) HasRevokedAccess() bool`

HasRevokedAccess returns a boolean if a field has been set.

### SetRevokedAccessNil

`func (o *OrganizationUserDataSourceFilters) SetRevokedAccessNil(b bool)`

 SetRevokedAccessNil sets the value for RevokedAccess to be an explicit nil

### UnsetRevokedAccess
`func (o *OrganizationUserDataSourceFilters) UnsetRevokedAccess()`

UnsetRevokedAccess ensures that no value is present for RevokedAccess, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


