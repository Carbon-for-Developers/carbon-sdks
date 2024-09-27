# OpportunitiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | **int32** |  | 
**IncludeRemoteData** | Pointer to **bool** |  | [optional] [default to false]
**NextCursor** | Pointer to **NullableString** |  | [optional] 
**PageSize** | Pointer to **NullableInt32** |  | [optional] 
**OrderDir** | Pointer to [**NullableOrderDirV2Nullable**](OrderDirV2Nullable.md) |  | [optional] [default to ORDERDIRV2NULLABLE_ASC]
**Includes** | Pointer to [**[]BaseIncludes**](BaseIncludes.md) |  | [optional] [default to []]
**Filters** | Pointer to [**OpportunityFilters**](OpportunityFilters.md) |  | [optional] 
**OrderBy** | Pointer to [**NullableOpportunitiesOrderByNullable**](OpportunitiesOrderByNullable.md) |  | [optional] 

## Methods

### NewOpportunitiesRequest

`func NewOpportunitiesRequest(dataSourceId int32, ) *OpportunitiesRequest`

NewOpportunitiesRequest instantiates a new OpportunitiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunitiesRequestWithDefaults

`func NewOpportunitiesRequestWithDefaults() *OpportunitiesRequest`

NewOpportunitiesRequestWithDefaults instantiates a new OpportunitiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *OpportunitiesRequest) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *OpportunitiesRequest) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *OpportunitiesRequest) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.


### GetIncludeRemoteData

`func (o *OpportunitiesRequest) GetIncludeRemoteData() bool`

GetIncludeRemoteData returns the IncludeRemoteData field if non-nil, zero value otherwise.

### GetIncludeRemoteDataOk

`func (o *OpportunitiesRequest) GetIncludeRemoteDataOk() (*bool, bool)`

GetIncludeRemoteDataOk returns a tuple with the IncludeRemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRemoteData

`func (o *OpportunitiesRequest) SetIncludeRemoteData(v bool)`

SetIncludeRemoteData sets IncludeRemoteData field to given value.

### HasIncludeRemoteData

`func (o *OpportunitiesRequest) HasIncludeRemoteData() bool`

HasIncludeRemoteData returns a boolean if a field has been set.

### GetNextCursor

`func (o *OpportunitiesRequest) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *OpportunitiesRequest) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *OpportunitiesRequest) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *OpportunitiesRequest) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *OpportunitiesRequest) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *OpportunitiesRequest) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetPageSize

`func (o *OpportunitiesRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OpportunitiesRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OpportunitiesRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *OpportunitiesRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *OpportunitiesRequest) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *OpportunitiesRequest) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetOrderDir

`func (o *OpportunitiesRequest) GetOrderDir() OrderDirV2Nullable`

GetOrderDir returns the OrderDir field if non-nil, zero value otherwise.

### GetOrderDirOk

`func (o *OpportunitiesRequest) GetOrderDirOk() (*OrderDirV2Nullable, bool)`

GetOrderDirOk returns a tuple with the OrderDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDir

`func (o *OpportunitiesRequest) SetOrderDir(v OrderDirV2Nullable)`

SetOrderDir sets OrderDir field to given value.

### HasOrderDir

`func (o *OpportunitiesRequest) HasOrderDir() bool`

HasOrderDir returns a boolean if a field has been set.

### SetOrderDirNil

`func (o *OpportunitiesRequest) SetOrderDirNil(b bool)`

 SetOrderDirNil sets the value for OrderDir to be an explicit nil

### UnsetOrderDir
`func (o *OpportunitiesRequest) UnsetOrderDir()`

UnsetOrderDir ensures that no value is present for OrderDir, not even an explicit nil
### GetIncludes

`func (o *OpportunitiesRequest) GetIncludes() []BaseIncludes`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *OpportunitiesRequest) GetIncludesOk() (*[]BaseIncludes, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *OpportunitiesRequest) SetIncludes(v []BaseIncludes)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *OpportunitiesRequest) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### GetFilters

`func (o *OpportunitiesRequest) GetFilters() OpportunityFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *OpportunitiesRequest) GetFiltersOk() (*OpportunityFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *OpportunitiesRequest) SetFilters(v OpportunityFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *OpportunitiesRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetOrderBy

`func (o *OpportunitiesRequest) GetOrderBy() OpportunitiesOrderByNullable`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *OpportunitiesRequest) GetOrderByOk() (*OpportunitiesOrderByNullable, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *OpportunitiesRequest) SetOrderBy(v OpportunitiesOrderByNullable)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *OpportunitiesRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *OpportunitiesRequest) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *OpportunitiesRequest) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


