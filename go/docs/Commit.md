# Commit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | [**NullableCommitUserNullable**](CommitUserNullable.md) |  | 
**Committer** | [**NullableCommitUserNullable**](CommitUserNullable.md) |  | 
**Message** | **string** |  | 
**Tree** | [**Tree**](Tree.md) |  | 
**Url** | **string** |  | 
**CommentCount** | **int32** |  | 

## Methods

### NewCommit

`func NewCommit(author NullableCommitUserNullable, committer NullableCommitUserNullable, message string, tree Tree, url string, commentCount int32, ) *Commit`

NewCommit instantiates a new Commit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitWithDefaults

`func NewCommitWithDefaults() *Commit`

NewCommitWithDefaults instantiates a new Commit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *Commit) GetAuthor() CommitUserNullable`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Commit) GetAuthorOk() (*CommitUserNullable, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Commit) SetAuthor(v CommitUserNullable)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *Commit) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *Commit) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetCommitter

`func (o *Commit) GetCommitter() CommitUserNullable`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *Commit) GetCommitterOk() (*CommitUserNullable, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *Commit) SetCommitter(v CommitUserNullable)`

SetCommitter sets Committer field to given value.


### SetCommitterNil

`func (o *Commit) SetCommitterNil(b bool)`

 SetCommitterNil sets the value for Committer to be an explicit nil

### UnsetCommitter
`func (o *Commit) UnsetCommitter()`

UnsetCommitter ensures that no value is present for Committer, not even an explicit nil
### GetMessage

`func (o *Commit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Commit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Commit) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTree

`func (o *Commit) GetTree() Tree`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *Commit) GetTreeOk() (*Tree, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *Commit) SetTree(v Tree)`

SetTree sets Tree field to given value.


### GetUrl

`func (o *Commit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Commit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Commit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCommentCount

`func (o *Commit) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *Commit) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *Commit) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


