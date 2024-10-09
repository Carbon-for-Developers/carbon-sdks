# CommitsInput

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

### NewCommitsInput

`func NewCommitsInput(dataSourceId int32, repository string, pullNumber int32, ) *CommitsInput`

NewCommitsInput instantiates a new CommitsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitsInputWithDefaults

`func NewCommitsInputWithDefaults() *CommitsInput`

NewCommitsInputWithDefaults instantiates a new CommitsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *CommitsInput) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *CommitsInput) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *CommitsInput) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.


### GetIncludeRemoteData

`func (o *CommitsInput) GetIncludeRemoteData() bool`

GetIncludeRemoteData returns the IncludeRemoteData field if non-nil, zero value otherwise.

### GetIncludeRemoteDataOk

`func (o *CommitsInput) GetIncludeRemoteDataOk() (*bool, bool)`

GetIncludeRemoteDataOk returns a tuple with the IncludeRemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRemoteData

`func (o *CommitsInput) SetIncludeRemoteData(v bool)`

SetIncludeRemoteData sets IncludeRemoteData field to given value.

### HasIncludeRemoteData

`func (o *CommitsInput) HasIncludeRemoteData() bool`

HasIncludeRemoteData returns a boolean if a field has been set.

### GetRepository

`func (o *CommitsInput) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CommitsInput) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CommitsInput) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetPage

`func (o *CommitsInput) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CommitsInput) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CommitsInput) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *CommitsInput) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *CommitsInput) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CommitsInput) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CommitsInput) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *CommitsInput) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetNextCursor

`func (o *CommitsInput) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *CommitsInput) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *CommitsInput) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *CommitsInput) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *CommitsInput) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *CommitsInput) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetPullNumber

`func (o *CommitsInput) GetPullNumber() int32`

GetPullNumber returns the PullNumber field if non-nil, zero value otherwise.

### GetPullNumberOk

`func (o *CommitsInput) GetPullNumberOk() (*int32, bool)`

GetPullNumberOk returns a tuple with the PullNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullNumber

`func (o *CommitsInput) SetPullNumber(v int32)`

SetPullNumber sets PullNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


