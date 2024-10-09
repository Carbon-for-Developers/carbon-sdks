# CommitsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PRCommit**](PRCommit.md) |  | 
**NextCursor** | **NullableString** |  | 

## Methods

### NewCommitsResponse

`func NewCommitsResponse(data []PRCommit, nextCursor NullableString, ) *CommitsResponse`

NewCommitsResponse instantiates a new CommitsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitsResponseWithDefaults

`func NewCommitsResponseWithDefaults() *CommitsResponse`

NewCommitsResponseWithDefaults instantiates a new CommitsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CommitsResponse) GetData() []PRCommit`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CommitsResponse) GetDataOk() (*[]PRCommit, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CommitsResponse) SetData(v []PRCommit)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *CommitsResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *CommitsResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *CommitsResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *CommitsResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *CommitsResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


