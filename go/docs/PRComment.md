# PRComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**PullRequestReviewId** | **NullableInt32** |  | 
**Url** | **string** |  | 
**DiffHunk** | **string** |  | 
**Path** | **string** |  | 
**User** | [**User**](User.md) |  | 
**Body** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**StartLine** | **NullableInt32** |  | 
**Line** | **NullableInt32** |  | 
**RemoteData** | **map[string]interface{}** |  | 

## Methods

### NewPRComment

`func NewPRComment(id int32, pullRequestReviewId NullableInt32, url string, diffHunk string, path string, user User, body string, createdAt string, updatedAt string, startLine NullableInt32, line NullableInt32, remoteData map[string]interface{}, ) *PRComment`

NewPRComment instantiates a new PRComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPRCommentWithDefaults

`func NewPRCommentWithDefaults() *PRComment`

NewPRCommentWithDefaults instantiates a new PRComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PRComment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PRComment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PRComment) SetId(v int32)`

SetId sets Id field to given value.


### GetPullRequestReviewId

`func (o *PRComment) GetPullRequestReviewId() int32`

GetPullRequestReviewId returns the PullRequestReviewId field if non-nil, zero value otherwise.

### GetPullRequestReviewIdOk

`func (o *PRComment) GetPullRequestReviewIdOk() (*int32, bool)`

GetPullRequestReviewIdOk returns a tuple with the PullRequestReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestReviewId

`func (o *PRComment) SetPullRequestReviewId(v int32)`

SetPullRequestReviewId sets PullRequestReviewId field to given value.


### SetPullRequestReviewIdNil

`func (o *PRComment) SetPullRequestReviewIdNil(b bool)`

 SetPullRequestReviewIdNil sets the value for PullRequestReviewId to be an explicit nil

### UnsetPullRequestReviewId
`func (o *PRComment) UnsetPullRequestReviewId()`

UnsetPullRequestReviewId ensures that no value is present for PullRequestReviewId, not even an explicit nil
### GetUrl

`func (o *PRComment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PRComment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PRComment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDiffHunk

`func (o *PRComment) GetDiffHunk() string`

GetDiffHunk returns the DiffHunk field if non-nil, zero value otherwise.

### GetDiffHunkOk

`func (o *PRComment) GetDiffHunkOk() (*string, bool)`

GetDiffHunkOk returns a tuple with the DiffHunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffHunk

`func (o *PRComment) SetDiffHunk(v string)`

SetDiffHunk sets DiffHunk field to given value.


### GetPath

`func (o *PRComment) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PRComment) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PRComment) SetPath(v string)`

SetPath sets Path field to given value.


### GetUser

`func (o *PRComment) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PRComment) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PRComment) SetUser(v User)`

SetUser sets User field to given value.


### GetBody

`func (o *PRComment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PRComment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PRComment) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreatedAt

`func (o *PRComment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PRComment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PRComment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PRComment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PRComment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PRComment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStartLine

`func (o *PRComment) GetStartLine() int32`

GetStartLine returns the StartLine field if non-nil, zero value otherwise.

### GetStartLineOk

`func (o *PRComment) GetStartLineOk() (*int32, bool)`

GetStartLineOk returns a tuple with the StartLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLine

`func (o *PRComment) SetStartLine(v int32)`

SetStartLine sets StartLine field to given value.


### SetStartLineNil

`func (o *PRComment) SetStartLineNil(b bool)`

 SetStartLineNil sets the value for StartLine to be an explicit nil

### UnsetStartLine
`func (o *PRComment) UnsetStartLine()`

UnsetStartLine ensures that no value is present for StartLine, not even an explicit nil
### GetLine

`func (o *PRComment) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *PRComment) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *PRComment) SetLine(v int32)`

SetLine sets Line field to given value.


### SetLineNil

`func (o *PRComment) SetLineNil(b bool)`

 SetLineNil sets the value for Line to be an explicit nil

### UnsetLine
`func (o *PRComment) UnsetLine()`

UnsetLine ensures that no value is present for Line, not even an explicit nil
### GetRemoteData

`func (o *PRComment) GetRemoteData() map[string]interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *PRComment) GetRemoteDataOk() (*map[string]interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *PRComment) SetRemoteData(v map[string]interface{})`

SetRemoteData sets RemoteData field to given value.


### SetRemoteDataNil

`func (o *PRComment) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *PRComment) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


