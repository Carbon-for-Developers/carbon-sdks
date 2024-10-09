# FilesInput

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

## Methods

### NewFilesInput

`func NewFilesInput(dataSourceId int32, repository string, pullNumber int32, ) *FilesInput`

NewFilesInput instantiates a new FilesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesInputWithDefaults

`func NewFilesInputWithDefaults() *FilesInput`

NewFilesInputWithDefaults instantiates a new FilesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *FilesInput) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *FilesInput) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *FilesInput) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.


### GetIncludeRemoteData

`func (o *FilesInput) GetIncludeRemoteData() bool`

GetIncludeRemoteData returns the IncludeRemoteData field if non-nil, zero value otherwise.

### GetIncludeRemoteDataOk

`func (o *FilesInput) GetIncludeRemoteDataOk() (*bool, bool)`

GetIncludeRemoteDataOk returns a tuple with the IncludeRemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRemoteData

`func (o *FilesInput) SetIncludeRemoteData(v bool)`

SetIncludeRemoteData sets IncludeRemoteData field to given value.

### HasIncludeRemoteData

`func (o *FilesInput) HasIncludeRemoteData() bool`

HasIncludeRemoteData returns a boolean if a field has been set.

### GetRepository

`func (o *FilesInput) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *FilesInput) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *FilesInput) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetPage

`func (o *FilesInput) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *FilesInput) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *FilesInput) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *FilesInput) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *FilesInput) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *FilesInput) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *FilesInput) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *FilesInput) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetNextCursor

`func (o *FilesInput) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *FilesInput) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *FilesInput) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *FilesInput) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *FilesInput) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *FilesInput) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetPullNumber

`func (o *FilesInput) GetPullNumber() int32`

GetPullNumber returns the PullNumber field if non-nil, zero value otherwise.

### GetPullNumberOk

`func (o *FilesInput) GetPullNumberOk() (*int32, bool)`

GetPullNumberOk returns a tuple with the PullNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullNumber

`func (o *FilesInput) SetPullNumber(v int32)`

SetPullNumber sets PullNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


