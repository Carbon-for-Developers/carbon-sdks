# CommentsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | **int32** |  | 
**IncludeRemoteData** | Pointer to **bool** |  | [optional] [default to false]
**Repository** | **string** | Full name of the repository, denoted as {owner}/{repo} | 
**Page** | Pointer to **int32** |  | [optional] [default to 1]
**PageSize** | Pointer to **int32** |  | [optional] [default to 30]
**NextCursor** | Pointer to **NullableString** |  | [optional] 
**PullNumber** | **int32** |  | 
**OrderBy** | Pointer to [**CommentsOrderBy**](CommentsOrderBy.md) |  | [optional] 
**OrderDir** | Pointer to [**NullableOrderDirV2Nullable**](OrderDirV2Nullable.md) |  | [optional] [default to ORDERDIRV2NULLABLE_ASC]

## Methods

### NewCommentsInput

`func NewCommentsInput(dataSourceId int32, repository string, pullNumber int32, ) *CommentsInput`

NewCommentsInput instantiates a new CommentsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentsInputWithDefaults

`func NewCommentsInputWithDefaults() *CommentsInput`

NewCommentsInputWithDefaults instantiates a new CommentsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *CommentsInput) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *CommentsInput) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *CommentsInput) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.


### GetIncludeRemoteData

`func (o *CommentsInput) GetIncludeRemoteData() bool`

GetIncludeRemoteData returns the IncludeRemoteData field if non-nil, zero value otherwise.

### GetIncludeRemoteDataOk

`func (o *CommentsInput) GetIncludeRemoteDataOk() (*bool, bool)`

GetIncludeRemoteDataOk returns a tuple with the IncludeRemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRemoteData

`func (o *CommentsInput) SetIncludeRemoteData(v bool)`

SetIncludeRemoteData sets IncludeRemoteData field to given value.

### HasIncludeRemoteData

`func (o *CommentsInput) HasIncludeRemoteData() bool`

HasIncludeRemoteData returns a boolean if a field has been set.

### GetRepository

`func (o *CommentsInput) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CommentsInput) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CommentsInput) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetPage

`func (o *CommentsInput) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CommentsInput) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CommentsInput) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *CommentsInput) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *CommentsInput) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CommentsInput) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CommentsInput) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *CommentsInput) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetNextCursor

`func (o *CommentsInput) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *CommentsInput) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *CommentsInput) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *CommentsInput) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *CommentsInput) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *CommentsInput) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetPullNumber

`func (o *CommentsInput) GetPullNumber() int32`

GetPullNumber returns the PullNumber field if non-nil, zero value otherwise.

### GetPullNumberOk

`func (o *CommentsInput) GetPullNumberOk() (*int32, bool)`

GetPullNumberOk returns a tuple with the PullNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullNumber

`func (o *CommentsInput) SetPullNumber(v int32)`

SetPullNumber sets PullNumber field to given value.


### GetOrderBy

`func (o *CommentsInput) GetOrderBy() CommentsOrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *CommentsInput) GetOrderByOk() (*CommentsOrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *CommentsInput) SetOrderBy(v CommentsOrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *CommentsInput) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetOrderDir

`func (o *CommentsInput) GetOrderDir() OrderDirV2Nullable`

GetOrderDir returns the OrderDir field if non-nil, zero value otherwise.

### GetOrderDirOk

`func (o *CommentsInput) GetOrderDirOk() (*OrderDirV2Nullable, bool)`

GetOrderDirOk returns a tuple with the OrderDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDir

`func (o *CommentsInput) SetOrderDir(v OrderDirV2Nullable)`

SetOrderDir sets OrderDir field to given value.

### HasOrderDir

`func (o *CommentsInput) HasOrderDir() bool`

HasOrderDir returns a boolean if a field has been set.

### SetOrderDirNil

`func (o *CommentsInput) SetOrderDirNil(b bool)`

 SetOrderDirNil sets the value for OrderDir to be an explicit nil

### UnsetOrderDir
`func (o *CommentsInput) UnsetOrderDir()`

UnsetOrderDir ensures that no value is present for OrderDir, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


