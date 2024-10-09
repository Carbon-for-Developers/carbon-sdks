# IssuesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | **int32** |  | 
**IncludeRemoteData** | Pointer to **bool** |  | [optional] [default to false]
**Repository** | **string** | Full name of the repository, denoted as {owner}/{repo} | 
**Page** | Pointer to **int32** |  | [optional] [default to 1]
**PageSize** | Pointer to **int32** |  | [optional] [default to 30]
**NextCursor** | Pointer to **NullableString** |  | [optional] 
**Filters** | Pointer to [**IssuesFilter**](IssuesFilter.md) |  | [optional] 
**OrderBy** | Pointer to [**IssuesOrderBy**](IssuesOrderBy.md) |  | [optional] 
**OrderDir** | Pointer to [**NullableOrderDirV2Nullable**](OrderDirV2Nullable.md) |  | [optional] [default to ORDERDIRV2NULLABLE_ASC]

## Methods

### NewIssuesInput

`func NewIssuesInput(dataSourceId int32, repository string, ) *IssuesInput`

NewIssuesInput instantiates a new IssuesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesInputWithDefaults

`func NewIssuesInputWithDefaults() *IssuesInput`

NewIssuesInputWithDefaults instantiates a new IssuesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *IssuesInput) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *IssuesInput) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *IssuesInput) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.


### GetIncludeRemoteData

`func (o *IssuesInput) GetIncludeRemoteData() bool`

GetIncludeRemoteData returns the IncludeRemoteData field if non-nil, zero value otherwise.

### GetIncludeRemoteDataOk

`func (o *IssuesInput) GetIncludeRemoteDataOk() (*bool, bool)`

GetIncludeRemoteDataOk returns a tuple with the IncludeRemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRemoteData

`func (o *IssuesInput) SetIncludeRemoteData(v bool)`

SetIncludeRemoteData sets IncludeRemoteData field to given value.

### HasIncludeRemoteData

`func (o *IssuesInput) HasIncludeRemoteData() bool`

HasIncludeRemoteData returns a boolean if a field has been set.

### GetRepository

`func (o *IssuesInput) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *IssuesInput) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *IssuesInput) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetPage

`func (o *IssuesInput) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *IssuesInput) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *IssuesInput) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *IssuesInput) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *IssuesInput) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *IssuesInput) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *IssuesInput) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *IssuesInput) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetNextCursor

`func (o *IssuesInput) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *IssuesInput) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *IssuesInput) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *IssuesInput) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *IssuesInput) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *IssuesInput) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetFilters

`func (o *IssuesInput) GetFilters() IssuesFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *IssuesInput) GetFiltersOk() (*IssuesFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *IssuesInput) SetFilters(v IssuesFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *IssuesInput) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetOrderBy

`func (o *IssuesInput) GetOrderBy() IssuesOrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *IssuesInput) GetOrderByOk() (*IssuesOrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *IssuesInput) SetOrderBy(v IssuesOrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *IssuesInput) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetOrderDir

`func (o *IssuesInput) GetOrderDir() OrderDirV2Nullable`

GetOrderDir returns the OrderDir field if non-nil, zero value otherwise.

### GetOrderDirOk

`func (o *IssuesInput) GetOrderDirOk() (*OrderDirV2Nullable, bool)`

GetOrderDirOk returns a tuple with the OrderDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDir

`func (o *IssuesInput) SetOrderDir(v OrderDirV2Nullable)`

SetOrderDir sets OrderDir field to given value.

### HasOrderDir

`func (o *IssuesInput) HasOrderDir() bool`

HasOrderDir returns a boolean if a field has been set.

### SetOrderDirNil

`func (o *IssuesInput) SetOrderDirNil(b bool)`

 SetOrderDirNil sets the value for OrderDir to be an explicit nil

### UnsetOrderDir
`func (o *IssuesInput) UnsetOrderDir()`

UnsetOrderDir ensures that no value is present for OrderDir, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


