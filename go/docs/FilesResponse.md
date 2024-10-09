# FilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PRFile**](PRFile.md) |  | 
**NextCursor** | **NullableString** |  | 

## Methods

### NewFilesResponse

`func NewFilesResponse(data []PRFile, nextCursor NullableString, ) *FilesResponse`

NewFilesResponse instantiates a new FilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesResponseWithDefaults

`func NewFilesResponseWithDefaults() *FilesResponse`

NewFilesResponseWithDefaults instantiates a new FilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FilesResponse) GetData() []PRFile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FilesResponse) GetDataOk() (*[]PRFile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FilesResponse) SetData(v []PRFile)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *FilesResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *FilesResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *FilesResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *FilesResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *FilesResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


