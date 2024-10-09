# PRFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha** | **string** |  | 
**Filename** | **string** |  | 
**Status** | [**FileStatus**](FileStatus.md) |  | 
**Additions** | **int32** |  | 
**Deletions** | **int32** |  | 
**Changes** | **int32** |  | 
**BlobUrl** | **string** |  | 
**RawUrl** | **string** |  | 
**ContentsUrl** | **string** |  | 
**RemoteData** | **map[string]interface{}** |  | 

## Methods

### NewPRFile

`func NewPRFile(sha string, filename string, status FileStatus, additions int32, deletions int32, changes int32, blobUrl string, rawUrl string, contentsUrl string, remoteData map[string]interface{}, ) *PRFile`

NewPRFile instantiates a new PRFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPRFileWithDefaults

`func NewPRFileWithDefaults() *PRFile`

NewPRFileWithDefaults instantiates a new PRFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha

`func (o *PRFile) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *PRFile) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *PRFile) SetSha(v string)`

SetSha sets Sha field to given value.


### GetFilename

`func (o *PRFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PRFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PRFile) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetStatus

`func (o *PRFile) GetStatus() FileStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PRFile) GetStatusOk() (*FileStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PRFile) SetStatus(v FileStatus)`

SetStatus sets Status field to given value.


### GetAdditions

`func (o *PRFile) GetAdditions() int32`

GetAdditions returns the Additions field if non-nil, zero value otherwise.

### GetAdditionsOk

`func (o *PRFile) GetAdditionsOk() (*int32, bool)`

GetAdditionsOk returns a tuple with the Additions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditions

`func (o *PRFile) SetAdditions(v int32)`

SetAdditions sets Additions field to given value.


### GetDeletions

`func (o *PRFile) GetDeletions() int32`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *PRFile) GetDeletionsOk() (*int32, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *PRFile) SetDeletions(v int32)`

SetDeletions sets Deletions field to given value.


### GetChanges

`func (o *PRFile) GetChanges() int32`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *PRFile) GetChangesOk() (*int32, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *PRFile) SetChanges(v int32)`

SetChanges sets Changes field to given value.


### GetBlobUrl

`func (o *PRFile) GetBlobUrl() string`

GetBlobUrl returns the BlobUrl field if non-nil, zero value otherwise.

### GetBlobUrlOk

`func (o *PRFile) GetBlobUrlOk() (*string, bool)`

GetBlobUrlOk returns a tuple with the BlobUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobUrl

`func (o *PRFile) SetBlobUrl(v string)`

SetBlobUrl sets BlobUrl field to given value.


### GetRawUrl

`func (o *PRFile) GetRawUrl() string`

GetRawUrl returns the RawUrl field if non-nil, zero value otherwise.

### GetRawUrlOk

`func (o *PRFile) GetRawUrlOk() (*string, bool)`

GetRawUrlOk returns a tuple with the RawUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawUrl

`func (o *PRFile) SetRawUrl(v string)`

SetRawUrl sets RawUrl field to given value.


### GetContentsUrl

`func (o *PRFile) GetContentsUrl() string`

GetContentsUrl returns the ContentsUrl field if non-nil, zero value otherwise.

### GetContentsUrlOk

`func (o *PRFile) GetContentsUrlOk() (*string, bool)`

GetContentsUrlOk returns a tuple with the ContentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsUrl

`func (o *PRFile) SetContentsUrl(v string)`

SetContentsUrl sets ContentsUrl field to given value.


### GetRemoteData

`func (o *PRFile) GetRemoteData() map[string]interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *PRFile) GetRemoteDataOk() (*map[string]interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *PRFile) SetRemoteData(v map[string]interface{})`

SetRemoteData sets RemoteData field to given value.


### SetRemoteDataNil

`func (o *PRFile) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *PRFile) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


