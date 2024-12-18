/*
Carbon

Connect external data to LLMs, no matter the source.

API version: 1.0.0
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package carbon

import (
	"encoding/json"
	"time"
)

// ExternalSourceItem struct for ExternalSourceItem
type ExternalSourceItem struct {
	Id int32 `json:"id"`
	ExternalId string `json:"external_id"`
	Source DataSourceType `json:"source"`
	Name string `json:"name"`
	SyncedAt time.Time `json:"synced_at"`
	IsSelectable NullableBool `json:"is_selectable"`
	IsExpandable NullableBool `json:"is_expandable"`
	OrganizationId int32 `json:"organization_id"`
	OrganizationSuppliedUserId string `json:"organization_supplied_user_id"`
	OrganizationUserId int32 `json:"organization_user_id"`
	OrganizationUserDataSourceId int32 `json:"organization_user_data_source_id"`
	OrganizationUserFileToSyncId NullableInt32 `json:"organization_user_file_to_sync_id"`
	ParentExternalId NullableString `json:"parent_external_id"`
	ItemType NullableString `json:"item_type"`
	RootExternalId NullableString `json:"root_external_id"`
	ExternalUrl NullableString `json:"external_url"`
	FileFormat NullableString `json:"file_format"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewExternalSourceItem instantiates a new ExternalSourceItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalSourceItem(id int32, externalId string, source DataSourceType, name string, syncedAt time.Time, isSelectable NullableBool, isExpandable NullableBool, organizationId int32, organizationSuppliedUserId string, organizationUserId int32, organizationUserDataSourceId int32, organizationUserFileToSyncId NullableInt32, parentExternalId NullableString, itemType NullableString, rootExternalId NullableString, externalUrl NullableString, fileFormat NullableString, createdAt time.Time, updatedAt time.Time) *ExternalSourceItem {
	this := ExternalSourceItem{}
	this.Id = id
	this.ExternalId = externalId
	this.Source = source
	this.Name = name
	this.SyncedAt = syncedAt
	this.IsSelectable = isSelectable
	this.IsExpandable = isExpandable
	this.OrganizationId = organizationId
	this.OrganizationSuppliedUserId = organizationSuppliedUserId
	this.OrganizationUserId = organizationUserId
	this.OrganizationUserDataSourceId = organizationUserDataSourceId
	this.OrganizationUserFileToSyncId = organizationUserFileToSyncId
	this.ParentExternalId = parentExternalId
	this.ItemType = itemType
	this.RootExternalId = rootExternalId
	this.ExternalUrl = externalUrl
	this.FileFormat = fileFormat
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewExternalSourceItemWithDefaults instantiates a new ExternalSourceItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalSourceItemWithDefaults() *ExternalSourceItem {
	this := ExternalSourceItem{}
	return &this
}

// GetId returns the Id field value
func (o *ExternalSourceItem) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalSourceItem) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalSourceItem) SetId(v int32) {
	o.Id = v
}

// GetExternalId returns the ExternalId field value
func (o *ExternalSourceItem) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *ExternalSourceItem) GetExternalIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *ExternalSourceItem) SetExternalId(v string) {
	o.ExternalId = v
}

// GetSource returns the Source field value
func (o *ExternalSourceItem) GetSource() DataSourceType {
	if o == nil {
		var ret DataSourceType
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ExternalSourceItem) GetSourceOk() (*DataSourceType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ExternalSourceItem) SetSource(v DataSourceType) {
	o.Source = v
}

// GetName returns the Name field value
func (o *ExternalSourceItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalSourceItem) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalSourceItem) SetName(v string) {
	o.Name = v
}

// GetSyncedAt returns the SyncedAt field value
func (o *ExternalSourceItem) GetSyncedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SyncedAt
}

// GetSyncedAtOk returns a tuple with the SyncedAt field value
// and a boolean to check if the value has been set.
func (o *ExternalSourceItem) GetSyncedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SyncedAt, true
}

// SetSyncedAt sets field value
func (o *ExternalSourceItem) SetSyncedAt(v time.Time) {
	o.SyncedAt = v
}

// GetIsSelectable returns the IsSelectable field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *ExternalSourceItem) GetIsSelectable() bool {
	if o == nil || o.IsSelectable.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsSelectable.Get()
}

// GetIsSelectableOk returns a tuple with the IsSelectable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalSourceItem) GetIsSelectableOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsSelectable.Get(), o.IsSelectable.IsSet()
}

// SetIsSelectable sets field value
func (o *ExternalSourceItem) SetIsSelectable(v bool) {
	o.IsSelectable.Set(&v)
}

// GetIsExpandable returns the IsExpandable field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *ExternalSourceItem) GetIsExpandable() bool {
	if o == nil || o.IsExpandable.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsExpandable.Get()
}

// GetIsExpandableOk returns a tuple with the IsExpandable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalSourceItem) GetIsExpandableOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsExpandable.Get(), o.IsExpandable.IsSet()
}

// SetIsExpandable sets field value
func (o *ExternalSourceItem) SetIsExpandable(v bool) {
	o.IsExpandable.Set(&v)
}

// GetOrganizationId returns the OrganizationId field value
func (o *ExternalSourceItem) GetOrganizationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *ExternalSourceItem) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *ExternalSourceItem) SetOrganizationId(v int32) {
	o.OrganizationId = v
}

// GetOrganizationSuppliedUserId returns the OrganizationSuppliedUserId field value
func (o *ExternalSourceItem) GetOrganizationSuppliedUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationSuppliedUserId
}

// GetOrganizationSuppliedUserIdOk returns a tuple with the OrganizationSuppliedUserId field value
// and a boolean to check if the value has been set.
func (o *ExternalSourceItem) GetOrganizationSuppliedUserIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrganizationSuppliedUserId, true
}

// SetOrganizationSuppliedUserId sets field value
func (o *ExternalSourceItem) SetOrganizationSuppliedUserId(v string) {
	o.OrganizationSuppliedUserId = v
}

// GetOrganizationUserId returns the OrganizationUserId field value
func (o *ExternalSourceItem) GetOrganizationUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrganizationUserId
}

// GetOrganizationUserIdOk returns a tuple with the OrganizationUserId field value
// and a boolean to check if the value has been set.
func (o *ExternalSourceItem) GetOrganizationUserIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrganizationUserId, true
}

// SetOrganizationUserId sets field value
func (o *ExternalSourceItem) SetOrganizationUserId(v int32) {
	o.OrganizationUserId = v
}

// GetOrganizationUserDataSourceId returns the OrganizationUserDataSourceId field value
func (o *ExternalSourceItem) GetOrganizationUserDataSourceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrganizationUserDataSourceId
}

// GetOrganizationUserDataSourceIdOk returns a tuple with the OrganizationUserDataSourceId field value
// and a boolean to check if the value has been set.
func (o *ExternalSourceItem) GetOrganizationUserDataSourceIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrganizationUserDataSourceId, true
}

// SetOrganizationUserDataSourceId sets field value
func (o *ExternalSourceItem) SetOrganizationUserDataSourceId(v int32) {
	o.OrganizationUserDataSourceId = v
}

// GetOrganizationUserFileToSyncId returns the OrganizationUserFileToSyncId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ExternalSourceItem) GetOrganizationUserFileToSyncId() int32 {
	if o == nil || o.OrganizationUserFileToSyncId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.OrganizationUserFileToSyncId.Get()
}

// GetOrganizationUserFileToSyncIdOk returns a tuple with the OrganizationUserFileToSyncId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalSourceItem) GetOrganizationUserFileToSyncIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.OrganizationUserFileToSyncId.Get(), o.OrganizationUserFileToSyncId.IsSet()
}

// SetOrganizationUserFileToSyncId sets field value
func (o *ExternalSourceItem) SetOrganizationUserFileToSyncId(v int32) {
	o.OrganizationUserFileToSyncId.Set(&v)
}

// GetParentExternalId returns the ParentExternalId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExternalSourceItem) GetParentExternalId() string {
	if o == nil || o.ParentExternalId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ParentExternalId.Get()
}

// GetParentExternalIdOk returns a tuple with the ParentExternalId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalSourceItem) GetParentExternalIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ParentExternalId.Get(), o.ParentExternalId.IsSet()
}

// SetParentExternalId sets field value
func (o *ExternalSourceItem) SetParentExternalId(v string) {
	o.ParentExternalId.Set(&v)
}

// GetItemType returns the ItemType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExternalSourceItem) GetItemType() string {
	if o == nil || o.ItemType.Get() == nil {
		var ret string
		return ret
	}

	return *o.ItemType.Get()
}

// GetItemTypeOk returns a tuple with the ItemType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalSourceItem) GetItemTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ItemType.Get(), o.ItemType.IsSet()
}

// SetItemType sets field value
func (o *ExternalSourceItem) SetItemType(v string) {
	o.ItemType.Set(&v)
}

// GetRootExternalId returns the RootExternalId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExternalSourceItem) GetRootExternalId() string {
	if o == nil || o.RootExternalId.Get() == nil {
		var ret string
		return ret
	}

	return *o.RootExternalId.Get()
}

// GetRootExternalIdOk returns a tuple with the RootExternalId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalSourceItem) GetRootExternalIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RootExternalId.Get(), o.RootExternalId.IsSet()
}

// SetRootExternalId sets field value
func (o *ExternalSourceItem) SetRootExternalId(v string) {
	o.RootExternalId.Set(&v)
}

// GetExternalUrl returns the ExternalUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExternalSourceItem) GetExternalUrl() string {
	if o == nil || o.ExternalUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalUrl.Get()
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalSourceItem) GetExternalUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ExternalUrl.Get(), o.ExternalUrl.IsSet()
}

// SetExternalUrl sets field value
func (o *ExternalSourceItem) SetExternalUrl(v string) {
	o.ExternalUrl.Set(&v)
}

// GetFileFormat returns the FileFormat field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExternalSourceItem) GetFileFormat() string {
	if o == nil || o.FileFormat.Get() == nil {
		var ret string
		return ret
	}

	return *o.FileFormat.Get()
}

// GetFileFormatOk returns a tuple with the FileFormat field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalSourceItem) GetFileFormatOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.FileFormat.Get(), o.FileFormat.IsSet()
}

// SetFileFormat sets field value
func (o *ExternalSourceItem) SetFileFormat(v string) {
	o.FileFormat.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *ExternalSourceItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ExternalSourceItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ExternalSourceItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ExternalSourceItem) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ExternalSourceItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ExternalSourceItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ExternalSourceItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["external_id"] = o.ExternalId
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["synced_at"] = o.SyncedAt
	}
	if true {
		toSerialize["is_selectable"] = o.IsSelectable.Get()
	}
	if true {
		toSerialize["is_expandable"] = o.IsExpandable.Get()
	}
	if true {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if true {
		toSerialize["organization_supplied_user_id"] = o.OrganizationSuppliedUserId
	}
	if true {
		toSerialize["organization_user_id"] = o.OrganizationUserId
	}
	if true {
		toSerialize["organization_user_data_source_id"] = o.OrganizationUserDataSourceId
	}
	if true {
		toSerialize["organization_user_file_to_sync_id"] = o.OrganizationUserFileToSyncId.Get()
	}
	if true {
		toSerialize["parent_external_id"] = o.ParentExternalId.Get()
	}
	if true {
		toSerialize["item_type"] = o.ItemType.Get()
	}
	if true {
		toSerialize["root_external_id"] = o.RootExternalId.Get()
	}
	if true {
		toSerialize["external_url"] = o.ExternalUrl.Get()
	}
	if true {
		toSerialize["file_format"] = o.FileFormat.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableExternalSourceItem struct {
	value *ExternalSourceItem
	isSet bool
}

func (v NullableExternalSourceItem) Get() *ExternalSourceItem {
	return v.value
}

func (v *NullableExternalSourceItem) Set(val *ExternalSourceItem) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalSourceItem) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalSourceItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalSourceItem(val *ExternalSourceItem) *NullableExternalSourceItem {
	return &NullableExternalSourceItem{value: val, isSet: true}
}

func (v NullableExternalSourceItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalSourceItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


