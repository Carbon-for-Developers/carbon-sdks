/*
Carbon

Connect external data to LLMs, no matter the source.

API version: 1.0.0
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package carbon

import (
	"encoding/json"
)

// OrganizationUserFilesToSyncFilters struct for OrganizationUserFilesToSyncFilters
type OrganizationUserFilesToSyncFilters struct {
	// Deprecated
	Tags map[string]Tags1 `json:"tags,omitempty"`
	Source NullableSourceProperty `json:"source,omitempty"`
	// The name of the file. The query will return files with names that contain this string.
	Name NullableString `json:"name,omitempty"`
	//          Tags to filter by. Supports logical AND and OR operations. Input should be like below:         {             \"OR\": [                 {                 \"key\": \"subject\",                 \"value\": \"holy-bible\",                 \"negate\": false                 },                 {                     \"key\": \"person-of-interest\",                     \"value\": \"jesus christ\",                     \"negate\": false                 },                 {                     \"key\": \"genre\",                     \"value\": \"fiction\",                     \"negate\": true                 }                 {                     \"AND\": [                         {                             \"key\": \"subject\",                             \"value\": \"tao-te-ching\",                             \"negate\": true                         },                         {                             \"key\": \"author\",                             \"value\": \"lao-tzu\",                             \"negate\": false                         }                     ]                 }             ]         }         For a single filter, the filter block can be placed within either an \"AND\" or \"OR\" block.         
	TagsV2 map[string]interface{} `json:"tags_v2,omitempty"`
	// The IDs of the files. The query will return files with these IDs.
	Ids []int32 `json:"ids,omitempty"`
	// The external file IDs of the files. The query will return files with these external file IDs.
	ExternalFileIds []string `json:"external_file_ids,omitempty"`
	// The sync statuses of the files. The query will return files with these sync statuses.
	SyncStatuses []ExternalFileSyncStatuses `json:"sync_statuses,omitempty"`
	// Deprecated
	ParentFileIds []int32 `json:"parent_file_ids,omitempty"`
	// The organization user data source IDs of the files. The query will return files with these organization user data source IDs.
	OrganizationUserDataSourceId []int32 `json:"organization_user_data_source_id,omitempty"`
	// The embedding generators of the files. The query will return files with these embedding generators.
	EmbeddingGenerators []EmbeddingGenerators `json:"embedding_generators,omitempty"`
	// If true, the query will return only root files. Cannot be true if parent_file_ids or include_all_children is specified.
	RootFilesOnly NullableBool `json:"root_files_only,omitempty"`
	// If true, the query will return all descendents of the specified parent_file_ids.
	IncludeAllChildren *bool `json:"include_all_children,omitempty"`
	// If true, the query will return only files that have not been synced yet.
	NonSyncedOnly *bool `json:"non_synced_only,omitempty"`
	// Filter by request ID(s) which were used to sync the files
	RequestIds []string `json:"request_ids,omitempty"`
	// Filter by upload ID(s) which were used to sync the files
	UploadIds []string `json:"upload_ids,omitempty"`
	// The error message of the file. The query will return files with error messages that contain this string. To search for files with no error message, use an empty string.
	SyncErrorMessage NullableString `json:"sync_error_message,omitempty"`
	// If true, the query will return containers in the response. Containers are files that group other files together and have no content themselves. Default behavior is to include containers.
	IncludeContainers NullableBool `json:"include_containers,omitempty"`
	// The external URLs of the files. The query will return files with these external URLs.
	ExternalUrls []string `json:"external_urls,omitempty"`
	// Filter files based on their type at the source (for example help center tickets and articles)
	FileTypesAtSource []HSNFileTypes2 `json:"file_types_at_source,omitempty"`
}

// NewOrganizationUserFilesToSyncFilters instantiates a new OrganizationUserFilesToSyncFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationUserFilesToSyncFilters() *OrganizationUserFilesToSyncFilters {
	this := OrganizationUserFilesToSyncFilters{}
	var includeAllChildren bool = false
	this.IncludeAllChildren = &includeAllChildren
	var nonSyncedOnly bool = false
	this.NonSyncedOnly = &nonSyncedOnly
	return &this
}

// NewOrganizationUserFilesToSyncFiltersWithDefaults instantiates a new OrganizationUserFilesToSyncFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationUserFilesToSyncFiltersWithDefaults() *OrganizationUserFilesToSyncFilters {
	this := OrganizationUserFilesToSyncFilters{}
	var includeAllChildren bool = false
	this.IncludeAllChildren = &includeAllChildren
	var nonSyncedOnly bool = false
	this.NonSyncedOnly = &nonSyncedOnly
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *OrganizationUserFilesToSyncFilters) GetTags() map[string]Tags1 {
	if o == nil {
		var ret map[string]Tags1
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *OrganizationUserFilesToSyncFilters) GetTagsOk() (*map[string]Tags1, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]Tags1 and assigns it to the Tags field.
// Deprecated
func (o *OrganizationUserFilesToSyncFilters) SetTags(v map[string]Tags1) {
	o.Tags = v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetSource() SourceProperty {
	if o == nil || isNil(o.Source.Get()) {
		var ret SourceProperty
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetSourceOk() (*SourceProperty, bool) {
	if o == nil {
    return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasSource() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableSourceProperty and assigns it to the Source field.
func (o *OrganizationUserFilesToSyncFilters) SetSource(v SourceProperty) {
	o.Source.Set(&v)
}
// SetSourceNil sets the value for Source to be an explicit nil
func (o *OrganizationUserFilesToSyncFilters) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *OrganizationUserFilesToSyncFilters) UnsetSource() {
	o.Source.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OrganizationUserFilesToSyncFilters) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *OrganizationUserFilesToSyncFilters) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OrganizationUserFilesToSyncFilters) UnsetName() {
	o.Name.Unset()
}

// GetTagsV2 returns the TagsV2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetTagsV2() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TagsV2
}

// GetTagsV2Ok returns a tuple with the TagsV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetTagsV2Ok() (map[string]interface{}, bool) {
	if o == nil || isNil(o.TagsV2) {
    return map[string]interface{}{}, false
	}
	return o.TagsV2, true
}

// HasTagsV2 returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasTagsV2() bool {
	if o != nil && isNil(o.TagsV2) {
		return true
	}

	return false
}

// SetTagsV2 gets a reference to the given map[string]interface{} and assigns it to the TagsV2 field.
func (o *OrganizationUserFilesToSyncFilters) SetTagsV2(v map[string]interface{}) {
	o.TagsV2 = v
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasIds() bool {
	if o != nil && isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *OrganizationUserFilesToSyncFilters) SetIds(v []int32) {
	o.Ids = v
}

// GetExternalFileIds returns the ExternalFileIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetExternalFileIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExternalFileIds
}

// GetExternalFileIdsOk returns a tuple with the ExternalFileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetExternalFileIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ExternalFileIds) {
    return nil, false
	}
	return o.ExternalFileIds, true
}

// HasExternalFileIds returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasExternalFileIds() bool {
	if o != nil && isNil(o.ExternalFileIds) {
		return true
	}

	return false
}

// SetExternalFileIds gets a reference to the given []string and assigns it to the ExternalFileIds field.
func (o *OrganizationUserFilesToSyncFilters) SetExternalFileIds(v []string) {
	o.ExternalFileIds = v
}

// GetSyncStatuses returns the SyncStatuses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetSyncStatuses() []ExternalFileSyncStatuses {
	if o == nil {
		var ret []ExternalFileSyncStatuses
		return ret
	}
	return o.SyncStatuses
}

// GetSyncStatusesOk returns a tuple with the SyncStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetSyncStatusesOk() ([]ExternalFileSyncStatuses, bool) {
	if o == nil || isNil(o.SyncStatuses) {
    return nil, false
	}
	return o.SyncStatuses, true
}

// HasSyncStatuses returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasSyncStatuses() bool {
	if o != nil && isNil(o.SyncStatuses) {
		return true
	}

	return false
}

// SetSyncStatuses gets a reference to the given []ExternalFileSyncStatuses and assigns it to the SyncStatuses field.
func (o *OrganizationUserFilesToSyncFilters) SetSyncStatuses(v []ExternalFileSyncStatuses) {
	o.SyncStatuses = v
}

// GetParentFileIds returns the ParentFileIds field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *OrganizationUserFilesToSyncFilters) GetParentFileIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.ParentFileIds
}

// GetParentFileIdsOk returns a tuple with the ParentFileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *OrganizationUserFilesToSyncFilters) GetParentFileIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.ParentFileIds) {
    return nil, false
	}
	return o.ParentFileIds, true
}

// HasParentFileIds returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasParentFileIds() bool {
	if o != nil && isNil(o.ParentFileIds) {
		return true
	}

	return false
}

// SetParentFileIds gets a reference to the given []int32 and assigns it to the ParentFileIds field.
// Deprecated
func (o *OrganizationUserFilesToSyncFilters) SetParentFileIds(v []int32) {
	o.ParentFileIds = v
}

// GetOrganizationUserDataSourceId returns the OrganizationUserDataSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetOrganizationUserDataSourceId() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.OrganizationUserDataSourceId
}

// GetOrganizationUserDataSourceIdOk returns a tuple with the OrganizationUserDataSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetOrganizationUserDataSourceIdOk() ([]int32, bool) {
	if o == nil || isNil(o.OrganizationUserDataSourceId) {
    return nil, false
	}
	return o.OrganizationUserDataSourceId, true
}

// HasOrganizationUserDataSourceId returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasOrganizationUserDataSourceId() bool {
	if o != nil && isNil(o.OrganizationUserDataSourceId) {
		return true
	}

	return false
}

// SetOrganizationUserDataSourceId gets a reference to the given []int32 and assigns it to the OrganizationUserDataSourceId field.
func (o *OrganizationUserFilesToSyncFilters) SetOrganizationUserDataSourceId(v []int32) {
	o.OrganizationUserDataSourceId = v
}

// GetEmbeddingGenerators returns the EmbeddingGenerators field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetEmbeddingGenerators() []EmbeddingGenerators {
	if o == nil {
		var ret []EmbeddingGenerators
		return ret
	}
	return o.EmbeddingGenerators
}

// GetEmbeddingGeneratorsOk returns a tuple with the EmbeddingGenerators field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetEmbeddingGeneratorsOk() ([]EmbeddingGenerators, bool) {
	if o == nil || isNil(o.EmbeddingGenerators) {
    return nil, false
	}
	return o.EmbeddingGenerators, true
}

// HasEmbeddingGenerators returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasEmbeddingGenerators() bool {
	if o != nil && isNil(o.EmbeddingGenerators) {
		return true
	}

	return false
}

// SetEmbeddingGenerators gets a reference to the given []EmbeddingGenerators and assigns it to the EmbeddingGenerators field.
func (o *OrganizationUserFilesToSyncFilters) SetEmbeddingGenerators(v []EmbeddingGenerators) {
	o.EmbeddingGenerators = v
}

// GetRootFilesOnly returns the RootFilesOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetRootFilesOnly() bool {
	if o == nil || isNil(o.RootFilesOnly.Get()) {
		var ret bool
		return ret
	}
	return *o.RootFilesOnly.Get()
}

// GetRootFilesOnlyOk returns a tuple with the RootFilesOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetRootFilesOnlyOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.RootFilesOnly.Get(), o.RootFilesOnly.IsSet()
}

// HasRootFilesOnly returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasRootFilesOnly() bool {
	if o != nil && o.RootFilesOnly.IsSet() {
		return true
	}

	return false
}

// SetRootFilesOnly gets a reference to the given NullableBool and assigns it to the RootFilesOnly field.
func (o *OrganizationUserFilesToSyncFilters) SetRootFilesOnly(v bool) {
	o.RootFilesOnly.Set(&v)
}
// SetRootFilesOnlyNil sets the value for RootFilesOnly to be an explicit nil
func (o *OrganizationUserFilesToSyncFilters) SetRootFilesOnlyNil() {
	o.RootFilesOnly.Set(nil)
}

// UnsetRootFilesOnly ensures that no value is present for RootFilesOnly, not even an explicit nil
func (o *OrganizationUserFilesToSyncFilters) UnsetRootFilesOnly() {
	o.RootFilesOnly.Unset()
}

// GetIncludeAllChildren returns the IncludeAllChildren field value if set, zero value otherwise.
func (o *OrganizationUserFilesToSyncFilters) GetIncludeAllChildren() bool {
	if o == nil || isNil(o.IncludeAllChildren) {
		var ret bool
		return ret
	}
	return *o.IncludeAllChildren
}

// GetIncludeAllChildrenOk returns a tuple with the IncludeAllChildren field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserFilesToSyncFilters) GetIncludeAllChildrenOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeAllChildren) {
    return nil, false
	}
	return o.IncludeAllChildren, true
}

// HasIncludeAllChildren returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasIncludeAllChildren() bool {
	if o != nil && !isNil(o.IncludeAllChildren) {
		return true
	}

	return false
}

// SetIncludeAllChildren gets a reference to the given bool and assigns it to the IncludeAllChildren field.
func (o *OrganizationUserFilesToSyncFilters) SetIncludeAllChildren(v bool) {
	o.IncludeAllChildren = &v
}

// GetNonSyncedOnly returns the NonSyncedOnly field value if set, zero value otherwise.
func (o *OrganizationUserFilesToSyncFilters) GetNonSyncedOnly() bool {
	if o == nil || isNil(o.NonSyncedOnly) {
		var ret bool
		return ret
	}
	return *o.NonSyncedOnly
}

// GetNonSyncedOnlyOk returns a tuple with the NonSyncedOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserFilesToSyncFilters) GetNonSyncedOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.NonSyncedOnly) {
    return nil, false
	}
	return o.NonSyncedOnly, true
}

// HasNonSyncedOnly returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasNonSyncedOnly() bool {
	if o != nil && !isNil(o.NonSyncedOnly) {
		return true
	}

	return false
}

// SetNonSyncedOnly gets a reference to the given bool and assigns it to the NonSyncedOnly field.
func (o *OrganizationUserFilesToSyncFilters) SetNonSyncedOnly(v bool) {
	o.NonSyncedOnly = &v
}

// GetRequestIds returns the RequestIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetRequestIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RequestIds
}

// GetRequestIdsOk returns a tuple with the RequestIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetRequestIdsOk() ([]string, bool) {
	if o == nil || isNil(o.RequestIds) {
    return nil, false
	}
	return o.RequestIds, true
}

// HasRequestIds returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasRequestIds() bool {
	if o != nil && isNil(o.RequestIds) {
		return true
	}

	return false
}

// SetRequestIds gets a reference to the given []string and assigns it to the RequestIds field.
func (o *OrganizationUserFilesToSyncFilters) SetRequestIds(v []string) {
	o.RequestIds = v
}

// GetUploadIds returns the UploadIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetUploadIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UploadIds
}

// GetUploadIdsOk returns a tuple with the UploadIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetUploadIdsOk() ([]string, bool) {
	if o == nil || isNil(o.UploadIds) {
    return nil, false
	}
	return o.UploadIds, true
}

// HasUploadIds returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasUploadIds() bool {
	if o != nil && isNil(o.UploadIds) {
		return true
	}

	return false
}

// SetUploadIds gets a reference to the given []string and assigns it to the UploadIds field.
func (o *OrganizationUserFilesToSyncFilters) SetUploadIds(v []string) {
	o.UploadIds = v
}

// GetSyncErrorMessage returns the SyncErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetSyncErrorMessage() string {
	if o == nil || isNil(o.SyncErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.SyncErrorMessage.Get()
}

// GetSyncErrorMessageOk returns a tuple with the SyncErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetSyncErrorMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SyncErrorMessage.Get(), o.SyncErrorMessage.IsSet()
}

// HasSyncErrorMessage returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasSyncErrorMessage() bool {
	if o != nil && o.SyncErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetSyncErrorMessage gets a reference to the given NullableString and assigns it to the SyncErrorMessage field.
func (o *OrganizationUserFilesToSyncFilters) SetSyncErrorMessage(v string) {
	o.SyncErrorMessage.Set(&v)
}
// SetSyncErrorMessageNil sets the value for SyncErrorMessage to be an explicit nil
func (o *OrganizationUserFilesToSyncFilters) SetSyncErrorMessageNil() {
	o.SyncErrorMessage.Set(nil)
}

// UnsetSyncErrorMessage ensures that no value is present for SyncErrorMessage, not even an explicit nil
func (o *OrganizationUserFilesToSyncFilters) UnsetSyncErrorMessage() {
	o.SyncErrorMessage.Unset()
}

// GetIncludeContainers returns the IncludeContainers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetIncludeContainers() bool {
	if o == nil || isNil(o.IncludeContainers.Get()) {
		var ret bool
		return ret
	}
	return *o.IncludeContainers.Get()
}

// GetIncludeContainersOk returns a tuple with the IncludeContainers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetIncludeContainersOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IncludeContainers.Get(), o.IncludeContainers.IsSet()
}

// HasIncludeContainers returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasIncludeContainers() bool {
	if o != nil && o.IncludeContainers.IsSet() {
		return true
	}

	return false
}

// SetIncludeContainers gets a reference to the given NullableBool and assigns it to the IncludeContainers field.
func (o *OrganizationUserFilesToSyncFilters) SetIncludeContainers(v bool) {
	o.IncludeContainers.Set(&v)
}
// SetIncludeContainersNil sets the value for IncludeContainers to be an explicit nil
func (o *OrganizationUserFilesToSyncFilters) SetIncludeContainersNil() {
	o.IncludeContainers.Set(nil)
}

// UnsetIncludeContainers ensures that no value is present for IncludeContainers, not even an explicit nil
func (o *OrganizationUserFilesToSyncFilters) UnsetIncludeContainers() {
	o.IncludeContainers.Unset()
}

// GetExternalUrls returns the ExternalUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetExternalUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetExternalUrlsOk() ([]string, bool) {
	if o == nil || isNil(o.ExternalUrls) {
    return nil, false
	}
	return o.ExternalUrls, true
}

// HasExternalUrls returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasExternalUrls() bool {
	if o != nil && isNil(o.ExternalUrls) {
		return true
	}

	return false
}

// SetExternalUrls gets a reference to the given []string and assigns it to the ExternalUrls field.
func (o *OrganizationUserFilesToSyncFilters) SetExternalUrls(v []string) {
	o.ExternalUrls = v
}

// GetFileTypesAtSource returns the FileTypesAtSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationUserFilesToSyncFilters) GetFileTypesAtSource() []HSNFileTypes2 {
	if o == nil {
		var ret []HSNFileTypes2
		return ret
	}
	return o.FileTypesAtSource
}

// GetFileTypesAtSourceOk returns a tuple with the FileTypesAtSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationUserFilesToSyncFilters) GetFileTypesAtSourceOk() ([]HSNFileTypes2, bool) {
	if o == nil || isNil(o.FileTypesAtSource) {
    return nil, false
	}
	return o.FileTypesAtSource, true
}

// HasFileTypesAtSource returns a boolean if a field has been set.
func (o *OrganizationUserFilesToSyncFilters) HasFileTypesAtSource() bool {
	if o != nil && isNil(o.FileTypesAtSource) {
		return true
	}

	return false
}

// SetFileTypesAtSource gets a reference to the given []HSNFileTypes2 and assigns it to the FileTypesAtSource field.
func (o *OrganizationUserFilesToSyncFilters) SetFileTypesAtSource(v []HSNFileTypes2) {
	o.FileTypesAtSource = v
}

func (o OrganizationUserFilesToSyncFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Source.IsSet() {
		toSerialize["source"] = o.Source.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.TagsV2 != nil {
		toSerialize["tags_v2"] = o.TagsV2
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.ExternalFileIds != nil {
		toSerialize["external_file_ids"] = o.ExternalFileIds
	}
	if o.SyncStatuses != nil {
		toSerialize["sync_statuses"] = o.SyncStatuses
	}
	if o.ParentFileIds != nil {
		toSerialize["parent_file_ids"] = o.ParentFileIds
	}
	if o.OrganizationUserDataSourceId != nil {
		toSerialize["organization_user_data_source_id"] = o.OrganizationUserDataSourceId
	}
	if o.EmbeddingGenerators != nil {
		toSerialize["embedding_generators"] = o.EmbeddingGenerators
	}
	if o.RootFilesOnly.IsSet() {
		toSerialize["root_files_only"] = o.RootFilesOnly.Get()
	}
	if !isNil(o.IncludeAllChildren) {
		toSerialize["include_all_children"] = o.IncludeAllChildren
	}
	if !isNil(o.NonSyncedOnly) {
		toSerialize["non_synced_only"] = o.NonSyncedOnly
	}
	if o.RequestIds != nil {
		toSerialize["request_ids"] = o.RequestIds
	}
	if o.UploadIds != nil {
		toSerialize["upload_ids"] = o.UploadIds
	}
	if o.SyncErrorMessage.IsSet() {
		toSerialize["sync_error_message"] = o.SyncErrorMessage.Get()
	}
	if o.IncludeContainers.IsSet() {
		toSerialize["include_containers"] = o.IncludeContainers.Get()
	}
	if o.ExternalUrls != nil {
		toSerialize["external_urls"] = o.ExternalUrls
	}
	if o.FileTypesAtSource != nil {
		toSerialize["file_types_at_source"] = o.FileTypesAtSource
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationUserFilesToSyncFilters struct {
	value *OrganizationUserFilesToSyncFilters
	isSet bool
}

func (v NullableOrganizationUserFilesToSyncFilters) Get() *OrganizationUserFilesToSyncFilters {
	return v.value
}

func (v *NullableOrganizationUserFilesToSyncFilters) Set(val *OrganizationUserFilesToSyncFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationUserFilesToSyncFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationUserFilesToSyncFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationUserFilesToSyncFilters(val *OrganizationUserFilesToSyncFilters) *NullableOrganizationUserFilesToSyncFilters {
	return &NullableOrganizationUserFilesToSyncFilters{value: val, isSet: true}
}

func (v NullableOrganizationUserFilesToSyncFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationUserFilesToSyncFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


