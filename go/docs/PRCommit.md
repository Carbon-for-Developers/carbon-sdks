# PRCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | [**Commit**](Commit.md) |  | 
**Url** | **string** |  | 
**Author** | [**NullableUserNullable**](UserNullable.md) |  | 
**Committer** | [**NullableUserNullable**](UserNullable.md) |  | 
**RemoteData** | **map[string]interface{}** |  | 

## Methods

### NewPRCommit

`func NewPRCommit(commit Commit, url string, author NullableUserNullable, committer NullableUserNullable, remoteData map[string]interface{}, ) *PRCommit`

NewPRCommit instantiates a new PRCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPRCommitWithDefaults

`func NewPRCommitWithDefaults() *PRCommit`

NewPRCommitWithDefaults instantiates a new PRCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *PRCommit) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *PRCommit) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *PRCommit) SetCommit(v Commit)`

SetCommit sets Commit field to given value.


### GetUrl

`func (o *PRCommit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PRCommit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PRCommit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAuthor

`func (o *PRCommit) GetAuthor() UserNullable`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PRCommit) GetAuthorOk() (*UserNullable, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PRCommit) SetAuthor(v UserNullable)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *PRCommit) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *PRCommit) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetCommitter

`func (o *PRCommit) GetCommitter() UserNullable`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *PRCommit) GetCommitterOk() (*UserNullable, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *PRCommit) SetCommitter(v UserNullable)`

SetCommitter sets Committer field to given value.


### SetCommitterNil

`func (o *PRCommit) SetCommitterNil(b bool)`

 SetCommitterNil sets the value for Committer to be an explicit nil

### UnsetCommitter
`func (o *PRCommit) UnsetCommitter()`

UnsetCommitter ensures that no value is present for Committer, not even an explicit nil
### GetRemoteData

`func (o *PRCommit) GetRemoteData() map[string]interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *PRCommit) GetRemoteDataOk() (*map[string]interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *PRCommit) SetRemoteData(v map[string]interface{})`

SetRemoteData sets RemoteData field to given value.


### SetRemoteDataNil

`func (o *PRCommit) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *PRCommit) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


