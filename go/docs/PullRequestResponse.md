# PullRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PullRequest**](PullRequest.md) |  | 
**NextCursor** | **NullableString** |  | 

## Methods

### NewPullRequestResponse

`func NewPullRequestResponse(data []PullRequest, nextCursor NullableString, ) *PullRequestResponse`

NewPullRequestResponse instantiates a new PullRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestResponseWithDefaults

`func NewPullRequestResponseWithDefaults() *PullRequestResponse`

NewPullRequestResponseWithDefaults instantiates a new PullRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PullRequestResponse) GetData() []PullRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PullRequestResponse) GetDataOk() (*[]PullRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PullRequestResponse) SetData(v []PullRequest)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *PullRequestResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PullRequestResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PullRequestResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *PullRequestResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *PullRequestResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


