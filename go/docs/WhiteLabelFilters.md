# WhiteLabelFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int32** |  | [optional] [default to []]
**DataSourceType** | Pointer to [**[]OauthBasedConnectors**](OauthBasedConnectors.md) |  | [optional] [default to []]

## Methods

### NewWhiteLabelFilters

`func NewWhiteLabelFilters() *WhiteLabelFilters`

NewWhiteLabelFilters instantiates a new WhiteLabelFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhiteLabelFiltersWithDefaults

`func NewWhiteLabelFiltersWithDefaults() *WhiteLabelFilters`

NewWhiteLabelFiltersWithDefaults instantiates a new WhiteLabelFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *WhiteLabelFilters) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *WhiteLabelFilters) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *WhiteLabelFilters) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *WhiteLabelFilters) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *WhiteLabelFilters) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *WhiteLabelFilters) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetDataSourceType

`func (o *WhiteLabelFilters) GetDataSourceType() []OauthBasedConnectors`

GetDataSourceType returns the DataSourceType field if non-nil, zero value otherwise.

### GetDataSourceTypeOk

`func (o *WhiteLabelFilters) GetDataSourceTypeOk() (*[]OauthBasedConnectors, bool)`

GetDataSourceTypeOk returns a tuple with the DataSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceType

`func (o *WhiteLabelFilters) SetDataSourceType(v []OauthBasedConnectors)`

SetDataSourceType sets DataSourceType field to given value.

### HasDataSourceType

`func (o *WhiteLabelFilters) HasDataSourceType() bool`

HasDataSourceType returns a boolean if a field has been set.

### SetDataSourceTypeNil

`func (o *WhiteLabelFilters) SetDataSourceTypeNil(b bool)`

 SetDataSourceTypeNil sets the value for DataSourceType to be an explicit nil

### UnsetDataSourceType
`func (o *WhiteLabelFilters) UnsetDataSourceType()`

UnsetDataSourceType ensures that no value is present for DataSourceType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


