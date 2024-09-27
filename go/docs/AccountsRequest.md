# AccountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | **int32** |  | 
**IncludeRemoteData** | Pointer to **bool** |  | [optional] [default to false]
**NextCursor** | Pointer to **NullableString** |  | [optional] 
**PageSize** | Pointer to **NullableInt32** |  | [optional] 
**OrderDir** | Pointer to [**NullableOrderDirV2Nullable**](OrderDirV2Nullable.md) |  | [optional] [default to ORDERDIRV2NULLABLE_ASC]
**Includes** | Pointer to [**[]BaseIncludes**](BaseIncludes.md) |  | [optional] [default to []]
**Filters** | Pointer to [**AccountFilters**](AccountFilters.md) |  | [optional] 
**OrderBy** | Pointer to [**NullableAccountsOrderByNullable**](AccountsOrderByNullable.md) |  | [optional] 

## Methods

### NewAccountsRequest

`func NewAccountsRequest(dataSourceId int32, ) *AccountsRequest`

NewAccountsRequest instantiates a new AccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsRequestWithDefaults

`func NewAccountsRequestWithDefaults() *AccountsRequest`

NewAccountsRequestWithDefaults instantiates a new AccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *AccountsRequest) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *AccountsRequest) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *AccountsRequest) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.


### GetIncludeRemoteData

`func (o *AccountsRequest) GetIncludeRemoteData() bool`

GetIncludeRemoteData returns the IncludeRemoteData field if non-nil, zero value otherwise.

### GetIncludeRemoteDataOk

`func (o *AccountsRequest) GetIncludeRemoteDataOk() (*bool, bool)`

GetIncludeRemoteDataOk returns a tuple with the IncludeRemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRemoteData

`func (o *AccountsRequest) SetIncludeRemoteData(v bool)`

SetIncludeRemoteData sets IncludeRemoteData field to given value.

### HasIncludeRemoteData

`func (o *AccountsRequest) HasIncludeRemoteData() bool`

HasIncludeRemoteData returns a boolean if a field has been set.

### GetNextCursor

`func (o *AccountsRequest) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *AccountsRequest) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *AccountsRequest) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *AccountsRequest) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *AccountsRequest) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *AccountsRequest) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetPageSize

`func (o *AccountsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *AccountsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *AccountsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *AccountsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *AccountsRequest) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *AccountsRequest) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetOrderDir

`func (o *AccountsRequest) GetOrderDir() OrderDirV2Nullable`

GetOrderDir returns the OrderDir field if non-nil, zero value otherwise.

### GetOrderDirOk

`func (o *AccountsRequest) GetOrderDirOk() (*OrderDirV2Nullable, bool)`

GetOrderDirOk returns a tuple with the OrderDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDir

`func (o *AccountsRequest) SetOrderDir(v OrderDirV2Nullable)`

SetOrderDir sets OrderDir field to given value.

### HasOrderDir

`func (o *AccountsRequest) HasOrderDir() bool`

HasOrderDir returns a boolean if a field has been set.

### SetOrderDirNil

`func (o *AccountsRequest) SetOrderDirNil(b bool)`

 SetOrderDirNil sets the value for OrderDir to be an explicit nil

### UnsetOrderDir
`func (o *AccountsRequest) UnsetOrderDir()`

UnsetOrderDir ensures that no value is present for OrderDir, not even an explicit nil
### GetIncludes

`func (o *AccountsRequest) GetIncludes() []BaseIncludes`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *AccountsRequest) GetIncludesOk() (*[]BaseIncludes, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *AccountsRequest) SetIncludes(v []BaseIncludes)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *AccountsRequest) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### GetFilters

`func (o *AccountsRequest) GetFilters() AccountFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AccountsRequest) GetFiltersOk() (*AccountFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AccountsRequest) SetFilters(v AccountFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AccountsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetOrderBy

`func (o *AccountsRequest) GetOrderBy() AccountsOrderByNullable`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *AccountsRequest) GetOrderByOk() (*AccountsOrderByNullable, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *AccountsRequest) SetOrderBy(v AccountsOrderByNullable)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *AccountsRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *AccountsRequest) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *AccountsRequest) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


