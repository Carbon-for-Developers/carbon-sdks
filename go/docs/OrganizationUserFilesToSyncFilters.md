# OrganizationUserFilesToSyncFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**map[string]Tags1**](Tags1.md) |  | [optional] 
**Source** | Pointer to [**NullableSourceProperty**](SourceProperty.md) |  | [optional] 
**Name** | Pointer to **NullableString** | The name of the file. The query will return files with names that contain this string. | [optional] 
**TagsV2** | Pointer to **map[string]interface{}** |          Tags to filter by. Supports logical AND and OR operations. Input should be like below:         {             \&quot;OR\&quot;: [                 {                 \&quot;key\&quot;: \&quot;subject\&quot;,                 \&quot;value\&quot;: \&quot;holy-bible\&quot;,                 \&quot;negate\&quot;: false                 },                 {                     \&quot;key\&quot;: \&quot;person-of-interest\&quot;,                     \&quot;value\&quot;: \&quot;jesus christ\&quot;,                     \&quot;negate\&quot;: false                 },                 {                     \&quot;key\&quot;: \&quot;genre\&quot;,                     \&quot;value\&quot;: \&quot;fiction\&quot;,                     \&quot;negate\&quot;: true                 }                 {                     \&quot;AND\&quot;: [                         {                             \&quot;key\&quot;: \&quot;subject\&quot;,                             \&quot;value\&quot;: \&quot;tao-te-ching\&quot;,                             \&quot;negate\&quot;: true                         },                         {                             \&quot;key\&quot;: \&quot;author\&quot;,                             \&quot;value\&quot;: \&quot;lao-tzu\&quot;,                             \&quot;negate\&quot;: false                         }                     ]                 }             ]         }         For a single filter, the filter block can be placed within either an \&quot;AND\&quot; or \&quot;OR\&quot; block.          | [optional] 
**Ids** | Pointer to **[]int32** | The IDs of the files. The query will return files with these IDs. | [optional] 
**ExternalFileIds** | Pointer to **[]string** | The external file IDs of the files. The query will return files with these external file IDs. | [optional] 
**SyncStatuses** | Pointer to [**[]ExternalFileSyncStatuses**](ExternalFileSyncStatuses.md) | The sync statuses of the files. The query will return files with these sync statuses. | [optional] 
**ParentFileIds** | Pointer to **[]int32** |  | [optional] 
**OrganizationUserDataSourceId** | Pointer to **[]int32** | The organization user data source IDs of the files. The query will return files with these organization user data source IDs. | [optional] 
**EmbeddingGenerators** | Pointer to [**[]EmbeddingGenerators**](EmbeddingGenerators.md) | The embedding generators of the files. The query will return files with these embedding generators. | [optional] 
**RootFilesOnly** | Pointer to **NullableBool** | If true, the query will return only root files. Cannot be true if parent_file_ids or include_all_children is specified. | [optional] 
**IncludeAllChildren** | Pointer to **bool** | If true, the query will return all descendents of the specified parent_file_ids. | [optional] [default to false]
**NonSyncedOnly** | Pointer to **bool** | If true, the query will return only files that have not been synced yet. | [optional] [default to false]
**RequestIds** | Pointer to **[]string** | Filter by request ID(s) which were used to sync the files | [optional] 
**SyncErrorMessage** | Pointer to **NullableString** | The error message of the file. The query will return files with error messages that contain this string. To search for files with no error message, use an empty string. | [optional] 
**IncludeContainers** | Pointer to **NullableBool** | If true, the query will return containers in the response. Containers are files that group other files together and have no content themselves. Default behavior is to include containers. | [optional] 
**ExternalUrls** | Pointer to **[]string** | The external URLs of the files. The query will return files with these external URLs. | [optional] 
**FileTypesAtSource** | Pointer to [**[]HSNFileTypes2**](HSNFileTypes2.md) | Filter files based on their type at the source (for example help center tickets and articles) | [optional] 

## Methods

### NewOrganizationUserFilesToSyncFilters

`func NewOrganizationUserFilesToSyncFilters() *OrganizationUserFilesToSyncFilters`

NewOrganizationUserFilesToSyncFilters instantiates a new OrganizationUserFilesToSyncFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationUserFilesToSyncFiltersWithDefaults

`func NewOrganizationUserFilesToSyncFiltersWithDefaults() *OrganizationUserFilesToSyncFilters`

NewOrganizationUserFilesToSyncFiltersWithDefaults instantiates a new OrganizationUserFilesToSyncFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *OrganizationUserFilesToSyncFilters) GetTags() map[string]Tags1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrganizationUserFilesToSyncFilters) GetTagsOk() (*map[string]Tags1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrganizationUserFilesToSyncFilters) SetTags(v map[string]Tags1)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrganizationUserFilesToSyncFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *OrganizationUserFilesToSyncFilters) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *OrganizationUserFilesToSyncFilters) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetSource

`func (o *OrganizationUserFilesToSyncFilters) GetSource() SourceProperty`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OrganizationUserFilesToSyncFilters) GetSourceOk() (*SourceProperty, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OrganizationUserFilesToSyncFilters) SetSource(v SourceProperty)`

SetSource sets Source field to given value.

### HasSource

`func (o *OrganizationUserFilesToSyncFilters) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *OrganizationUserFilesToSyncFilters) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *OrganizationUserFilesToSyncFilters) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetName

`func (o *OrganizationUserFilesToSyncFilters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationUserFilesToSyncFilters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationUserFilesToSyncFilters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationUserFilesToSyncFilters) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganizationUserFilesToSyncFilters) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationUserFilesToSyncFilters) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTagsV2

`func (o *OrganizationUserFilesToSyncFilters) GetTagsV2() map[string]interface{}`

GetTagsV2 returns the TagsV2 field if non-nil, zero value otherwise.

### GetTagsV2Ok

`func (o *OrganizationUserFilesToSyncFilters) GetTagsV2Ok() (*map[string]interface{}, bool)`

GetTagsV2Ok returns a tuple with the TagsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsV2

`func (o *OrganizationUserFilesToSyncFilters) SetTagsV2(v map[string]interface{})`

SetTagsV2 sets TagsV2 field to given value.

### HasTagsV2

`func (o *OrganizationUserFilesToSyncFilters) HasTagsV2() bool`

HasTagsV2 returns a boolean if a field has been set.

### SetTagsV2Nil

`func (o *OrganizationUserFilesToSyncFilters) SetTagsV2Nil(b bool)`

 SetTagsV2Nil sets the value for TagsV2 to be an explicit nil

### UnsetTagsV2
`func (o *OrganizationUserFilesToSyncFilters) UnsetTagsV2()`

UnsetTagsV2 ensures that no value is present for TagsV2, not even an explicit nil
### GetIds

`func (o *OrganizationUserFilesToSyncFilters) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *OrganizationUserFilesToSyncFilters) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *OrganizationUserFilesToSyncFilters) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *OrganizationUserFilesToSyncFilters) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *OrganizationUserFilesToSyncFilters) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *OrganizationUserFilesToSyncFilters) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetExternalFileIds

`func (o *OrganizationUserFilesToSyncFilters) GetExternalFileIds() []string`

GetExternalFileIds returns the ExternalFileIds field if non-nil, zero value otherwise.

### GetExternalFileIdsOk

`func (o *OrganizationUserFilesToSyncFilters) GetExternalFileIdsOk() (*[]string, bool)`

GetExternalFileIdsOk returns a tuple with the ExternalFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFileIds

`func (o *OrganizationUserFilesToSyncFilters) SetExternalFileIds(v []string)`

SetExternalFileIds sets ExternalFileIds field to given value.

### HasExternalFileIds

`func (o *OrganizationUserFilesToSyncFilters) HasExternalFileIds() bool`

HasExternalFileIds returns a boolean if a field has been set.

### SetExternalFileIdsNil

`func (o *OrganizationUserFilesToSyncFilters) SetExternalFileIdsNil(b bool)`

 SetExternalFileIdsNil sets the value for ExternalFileIds to be an explicit nil

### UnsetExternalFileIds
`func (o *OrganizationUserFilesToSyncFilters) UnsetExternalFileIds()`

UnsetExternalFileIds ensures that no value is present for ExternalFileIds, not even an explicit nil
### GetSyncStatuses

`func (o *OrganizationUserFilesToSyncFilters) GetSyncStatuses() []ExternalFileSyncStatuses`

GetSyncStatuses returns the SyncStatuses field if non-nil, zero value otherwise.

### GetSyncStatusesOk

`func (o *OrganizationUserFilesToSyncFilters) GetSyncStatusesOk() (*[]ExternalFileSyncStatuses, bool)`

GetSyncStatusesOk returns a tuple with the SyncStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatuses

`func (o *OrganizationUserFilesToSyncFilters) SetSyncStatuses(v []ExternalFileSyncStatuses)`

SetSyncStatuses sets SyncStatuses field to given value.

### HasSyncStatuses

`func (o *OrganizationUserFilesToSyncFilters) HasSyncStatuses() bool`

HasSyncStatuses returns a boolean if a field has been set.

### SetSyncStatusesNil

`func (o *OrganizationUserFilesToSyncFilters) SetSyncStatusesNil(b bool)`

 SetSyncStatusesNil sets the value for SyncStatuses to be an explicit nil

### UnsetSyncStatuses
`func (o *OrganizationUserFilesToSyncFilters) UnsetSyncStatuses()`

UnsetSyncStatuses ensures that no value is present for SyncStatuses, not even an explicit nil
### GetParentFileIds

`func (o *OrganizationUserFilesToSyncFilters) GetParentFileIds() []int32`

GetParentFileIds returns the ParentFileIds field if non-nil, zero value otherwise.

### GetParentFileIdsOk

`func (o *OrganizationUserFilesToSyncFilters) GetParentFileIdsOk() (*[]int32, bool)`

GetParentFileIdsOk returns a tuple with the ParentFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFileIds

`func (o *OrganizationUserFilesToSyncFilters) SetParentFileIds(v []int32)`

SetParentFileIds sets ParentFileIds field to given value.

### HasParentFileIds

`func (o *OrganizationUserFilesToSyncFilters) HasParentFileIds() bool`

HasParentFileIds returns a boolean if a field has been set.

### SetParentFileIdsNil

`func (o *OrganizationUserFilesToSyncFilters) SetParentFileIdsNil(b bool)`

 SetParentFileIdsNil sets the value for ParentFileIds to be an explicit nil

### UnsetParentFileIds
`func (o *OrganizationUserFilesToSyncFilters) UnsetParentFileIds()`

UnsetParentFileIds ensures that no value is present for ParentFileIds, not even an explicit nil
### GetOrganizationUserDataSourceId

`func (o *OrganizationUserFilesToSyncFilters) GetOrganizationUserDataSourceId() []int32`

GetOrganizationUserDataSourceId returns the OrganizationUserDataSourceId field if non-nil, zero value otherwise.

### GetOrganizationUserDataSourceIdOk

`func (o *OrganizationUserFilesToSyncFilters) GetOrganizationUserDataSourceIdOk() (*[]int32, bool)`

GetOrganizationUserDataSourceIdOk returns a tuple with the OrganizationUserDataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUserDataSourceId

`func (o *OrganizationUserFilesToSyncFilters) SetOrganizationUserDataSourceId(v []int32)`

SetOrganizationUserDataSourceId sets OrganizationUserDataSourceId field to given value.

### HasOrganizationUserDataSourceId

`func (o *OrganizationUserFilesToSyncFilters) HasOrganizationUserDataSourceId() bool`

HasOrganizationUserDataSourceId returns a boolean if a field has been set.

### SetOrganizationUserDataSourceIdNil

`func (o *OrganizationUserFilesToSyncFilters) SetOrganizationUserDataSourceIdNil(b bool)`

 SetOrganizationUserDataSourceIdNil sets the value for OrganizationUserDataSourceId to be an explicit nil

### UnsetOrganizationUserDataSourceId
`func (o *OrganizationUserFilesToSyncFilters) UnsetOrganizationUserDataSourceId()`

UnsetOrganizationUserDataSourceId ensures that no value is present for OrganizationUserDataSourceId, not even an explicit nil
### GetEmbeddingGenerators

`func (o *OrganizationUserFilesToSyncFilters) GetEmbeddingGenerators() []EmbeddingGenerators`

GetEmbeddingGenerators returns the EmbeddingGenerators field if non-nil, zero value otherwise.

### GetEmbeddingGeneratorsOk

`func (o *OrganizationUserFilesToSyncFilters) GetEmbeddingGeneratorsOk() (*[]EmbeddingGenerators, bool)`

GetEmbeddingGeneratorsOk returns a tuple with the EmbeddingGenerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingGenerators

`func (o *OrganizationUserFilesToSyncFilters) SetEmbeddingGenerators(v []EmbeddingGenerators)`

SetEmbeddingGenerators sets EmbeddingGenerators field to given value.

### HasEmbeddingGenerators

`func (o *OrganizationUserFilesToSyncFilters) HasEmbeddingGenerators() bool`

HasEmbeddingGenerators returns a boolean if a field has been set.

### SetEmbeddingGeneratorsNil

`func (o *OrganizationUserFilesToSyncFilters) SetEmbeddingGeneratorsNil(b bool)`

 SetEmbeddingGeneratorsNil sets the value for EmbeddingGenerators to be an explicit nil

### UnsetEmbeddingGenerators
`func (o *OrganizationUserFilesToSyncFilters) UnsetEmbeddingGenerators()`

UnsetEmbeddingGenerators ensures that no value is present for EmbeddingGenerators, not even an explicit nil
### GetRootFilesOnly

`func (o *OrganizationUserFilesToSyncFilters) GetRootFilesOnly() bool`

GetRootFilesOnly returns the RootFilesOnly field if non-nil, zero value otherwise.

### GetRootFilesOnlyOk

`func (o *OrganizationUserFilesToSyncFilters) GetRootFilesOnlyOk() (*bool, bool)`

GetRootFilesOnlyOk returns a tuple with the RootFilesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFilesOnly

`func (o *OrganizationUserFilesToSyncFilters) SetRootFilesOnly(v bool)`

SetRootFilesOnly sets RootFilesOnly field to given value.

### HasRootFilesOnly

`func (o *OrganizationUserFilesToSyncFilters) HasRootFilesOnly() bool`

HasRootFilesOnly returns a boolean if a field has been set.

### SetRootFilesOnlyNil

`func (o *OrganizationUserFilesToSyncFilters) SetRootFilesOnlyNil(b bool)`

 SetRootFilesOnlyNil sets the value for RootFilesOnly to be an explicit nil

### UnsetRootFilesOnly
`func (o *OrganizationUserFilesToSyncFilters) UnsetRootFilesOnly()`

UnsetRootFilesOnly ensures that no value is present for RootFilesOnly, not even an explicit nil
### GetIncludeAllChildren

`func (o *OrganizationUserFilesToSyncFilters) GetIncludeAllChildren() bool`

GetIncludeAllChildren returns the IncludeAllChildren field if non-nil, zero value otherwise.

### GetIncludeAllChildrenOk

`func (o *OrganizationUserFilesToSyncFilters) GetIncludeAllChildrenOk() (*bool, bool)`

GetIncludeAllChildrenOk returns a tuple with the IncludeAllChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllChildren

`func (o *OrganizationUserFilesToSyncFilters) SetIncludeAllChildren(v bool)`

SetIncludeAllChildren sets IncludeAllChildren field to given value.

### HasIncludeAllChildren

`func (o *OrganizationUserFilesToSyncFilters) HasIncludeAllChildren() bool`

HasIncludeAllChildren returns a boolean if a field has been set.

### GetNonSyncedOnly

`func (o *OrganizationUserFilesToSyncFilters) GetNonSyncedOnly() bool`

GetNonSyncedOnly returns the NonSyncedOnly field if non-nil, zero value otherwise.

### GetNonSyncedOnlyOk

`func (o *OrganizationUserFilesToSyncFilters) GetNonSyncedOnlyOk() (*bool, bool)`

GetNonSyncedOnlyOk returns a tuple with the NonSyncedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSyncedOnly

`func (o *OrganizationUserFilesToSyncFilters) SetNonSyncedOnly(v bool)`

SetNonSyncedOnly sets NonSyncedOnly field to given value.

### HasNonSyncedOnly

`func (o *OrganizationUserFilesToSyncFilters) HasNonSyncedOnly() bool`

HasNonSyncedOnly returns a boolean if a field has been set.

### GetRequestIds

`func (o *OrganizationUserFilesToSyncFilters) GetRequestIds() []string`

GetRequestIds returns the RequestIds field if non-nil, zero value otherwise.

### GetRequestIdsOk

`func (o *OrganizationUserFilesToSyncFilters) GetRequestIdsOk() (*[]string, bool)`

GetRequestIdsOk returns a tuple with the RequestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestIds

`func (o *OrganizationUserFilesToSyncFilters) SetRequestIds(v []string)`

SetRequestIds sets RequestIds field to given value.

### HasRequestIds

`func (o *OrganizationUserFilesToSyncFilters) HasRequestIds() bool`

HasRequestIds returns a boolean if a field has been set.

### SetRequestIdsNil

`func (o *OrganizationUserFilesToSyncFilters) SetRequestIdsNil(b bool)`

 SetRequestIdsNil sets the value for RequestIds to be an explicit nil

### UnsetRequestIds
`func (o *OrganizationUserFilesToSyncFilters) UnsetRequestIds()`

UnsetRequestIds ensures that no value is present for RequestIds, not even an explicit nil
### GetSyncErrorMessage

`func (o *OrganizationUserFilesToSyncFilters) GetSyncErrorMessage() string`

GetSyncErrorMessage returns the SyncErrorMessage field if non-nil, zero value otherwise.

### GetSyncErrorMessageOk

`func (o *OrganizationUserFilesToSyncFilters) GetSyncErrorMessageOk() (*string, bool)`

GetSyncErrorMessageOk returns a tuple with the SyncErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncErrorMessage

`func (o *OrganizationUserFilesToSyncFilters) SetSyncErrorMessage(v string)`

SetSyncErrorMessage sets SyncErrorMessage field to given value.

### HasSyncErrorMessage

`func (o *OrganizationUserFilesToSyncFilters) HasSyncErrorMessage() bool`

HasSyncErrorMessage returns a boolean if a field has been set.

### SetSyncErrorMessageNil

`func (o *OrganizationUserFilesToSyncFilters) SetSyncErrorMessageNil(b bool)`

 SetSyncErrorMessageNil sets the value for SyncErrorMessage to be an explicit nil

### UnsetSyncErrorMessage
`func (o *OrganizationUserFilesToSyncFilters) UnsetSyncErrorMessage()`

UnsetSyncErrorMessage ensures that no value is present for SyncErrorMessage, not even an explicit nil
### GetIncludeContainers

`func (o *OrganizationUserFilesToSyncFilters) GetIncludeContainers() bool`

GetIncludeContainers returns the IncludeContainers field if non-nil, zero value otherwise.

### GetIncludeContainersOk

`func (o *OrganizationUserFilesToSyncFilters) GetIncludeContainersOk() (*bool, bool)`

GetIncludeContainersOk returns a tuple with the IncludeContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeContainers

`func (o *OrganizationUserFilesToSyncFilters) SetIncludeContainers(v bool)`

SetIncludeContainers sets IncludeContainers field to given value.

### HasIncludeContainers

`func (o *OrganizationUserFilesToSyncFilters) HasIncludeContainers() bool`

HasIncludeContainers returns a boolean if a field has been set.

### SetIncludeContainersNil

`func (o *OrganizationUserFilesToSyncFilters) SetIncludeContainersNil(b bool)`

 SetIncludeContainersNil sets the value for IncludeContainers to be an explicit nil

### UnsetIncludeContainers
`func (o *OrganizationUserFilesToSyncFilters) UnsetIncludeContainers()`

UnsetIncludeContainers ensures that no value is present for IncludeContainers, not even an explicit nil
### GetExternalUrls

`func (o *OrganizationUserFilesToSyncFilters) GetExternalUrls() []string`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *OrganizationUserFilesToSyncFilters) GetExternalUrlsOk() (*[]string, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *OrganizationUserFilesToSyncFilters) SetExternalUrls(v []string)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *OrganizationUserFilesToSyncFilters) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### SetExternalUrlsNil

`func (o *OrganizationUserFilesToSyncFilters) SetExternalUrlsNil(b bool)`

 SetExternalUrlsNil sets the value for ExternalUrls to be an explicit nil

### UnsetExternalUrls
`func (o *OrganizationUserFilesToSyncFilters) UnsetExternalUrls()`

UnsetExternalUrls ensures that no value is present for ExternalUrls, not even an explicit nil
### GetFileTypesAtSource

`func (o *OrganizationUserFilesToSyncFilters) GetFileTypesAtSource() []HSNFileTypes2`

GetFileTypesAtSource returns the FileTypesAtSource field if non-nil, zero value otherwise.

### GetFileTypesAtSourceOk

`func (o *OrganizationUserFilesToSyncFilters) GetFileTypesAtSourceOk() (*[]HSNFileTypes2, bool)`

GetFileTypesAtSourceOk returns a tuple with the FileTypesAtSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTypesAtSource

`func (o *OrganizationUserFilesToSyncFilters) SetFileTypesAtSource(v []HSNFileTypes2)`

SetFileTypesAtSource sets FileTypesAtSource field to given value.

### HasFileTypesAtSource

`func (o *OrganizationUserFilesToSyncFilters) HasFileTypesAtSource() bool`

HasFileTypesAtSource returns a boolean if a field has been set.

### SetFileTypesAtSourceNil

`func (o *OrganizationUserFilesToSyncFilters) SetFileTypesAtSourceNil(b bool)`

 SetFileTypesAtSourceNil sets the value for FileTypesAtSource to be an explicit nil

### UnsetFileTypesAtSource
`func (o *OrganizationUserFilesToSyncFilters) UnsetFileTypesAtSource()`

UnsetFileTypesAtSource ensures that no value is present for FileTypesAtSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


