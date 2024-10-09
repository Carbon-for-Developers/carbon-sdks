# PullRequestsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | **int32** |  | 
**IncludeRemoteData** | Pointer to **bool** |  | [optional] [default to false]
**Repository** | **string** | Full name of the repository, denoted as {owner}/{repo} | 
**Page** | Pointer to **int32** |  | [optional] [default to 1]
**PageSize** | Pointer to **int32** |  | [optional] [default to 30]
**NextCursor** | Pointer to **NullableString** |  | [optional] 
**Filters** | Pointer to [**PullRequestFilters**](PullRequestFilters.md) |  | [optional] 
**OrderBy** | Pointer to [**PROrderBy**](PROrderBy.md) |  | [optional] 
**OrderDir** | Pointer to [**NullableOrderDirV2Nullable**](OrderDirV2Nullable.md) |  | [optional] [default to ORDERDIRV2NULLABLE_ASC]

## Methods

### NewPullRequestsInput

`func NewPullRequestsInput(dataSourceId int32, repository string, ) *PullRequestsInput`

NewPullRequestsInput instantiates a new PullRequestsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestsInputWithDefaults

`func NewPullRequestsInputWithDefaults() *PullRequestsInput`

NewPullRequestsInputWithDefaults instantiates a new PullRequestsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *PullRequestsInput) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *PullRequestsInput) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *PullRequestsInput) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.


### GetIncludeRemoteData

`func (o *PullRequestsInput) GetIncludeRemoteData() bool`

GetIncludeRemoteData returns the IncludeRemoteData field if non-nil, zero value otherwise.

### GetIncludeRemoteDataOk

`func (o *PullRequestsInput) GetIncludeRemoteDataOk() (*bool, bool)`

GetIncludeRemoteDataOk returns a tuple with the IncludeRemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRemoteData

`func (o *PullRequestsInput) SetIncludeRemoteData(v bool)`

SetIncludeRemoteData sets IncludeRemoteData field to given value.

### HasIncludeRemoteData

`func (o *PullRequestsInput) HasIncludeRemoteData() bool`

HasIncludeRemoteData returns a boolean if a field has been set.

### GetRepository

`func (o *PullRequestsInput) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PullRequestsInput) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PullRequestsInput) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetPage

`func (o *PullRequestsInput) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PullRequestsInput) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PullRequestsInput) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PullRequestsInput) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *PullRequestsInput) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PullRequestsInput) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PullRequestsInput) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PullRequestsInput) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetNextCursor

`func (o *PullRequestsInput) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PullRequestsInput) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PullRequestsInput) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *PullRequestsInput) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *PullRequestsInput) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *PullRequestsInput) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetFilters

`func (o *PullRequestsInput) GetFilters() PullRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PullRequestsInput) GetFiltersOk() (*PullRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PullRequestsInput) SetFilters(v PullRequestFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PullRequestsInput) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetOrderBy

`func (o *PullRequestsInput) GetOrderBy() PROrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *PullRequestsInput) GetOrderByOk() (*PROrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *PullRequestsInput) SetOrderBy(v PROrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *PullRequestsInput) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetOrderDir

`func (o *PullRequestsInput) GetOrderDir() OrderDirV2Nullable`

GetOrderDir returns the OrderDir field if non-nil, zero value otherwise.

### GetOrderDirOk

`func (o *PullRequestsInput) GetOrderDirOk() (*OrderDirV2Nullable, bool)`

GetOrderDirOk returns a tuple with the OrderDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDir

`func (o *PullRequestsInput) SetOrderDir(v OrderDirV2Nullable)`

SetOrderDir sets OrderDir field to given value.

### HasOrderDir

`func (o *PullRequestsInput) HasOrderDir() bool`

HasOrderDir returns a boolean if a field has been set.

### SetOrderDirNil

`func (o *PullRequestsInput) SetOrderDirNil(b bool)`

 SetOrderDirNil sets the value for OrderDir to be an explicit nil

### UnsetOrderDir
`func (o *PullRequestsInput) UnsetOrderDir()`

UnsetOrderDir ensures that no value is present for OrderDir, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


