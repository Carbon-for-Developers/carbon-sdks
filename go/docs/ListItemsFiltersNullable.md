# ListItemsFiltersNullable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalIds** | Pointer to **[]string** |  | [optional] 
**Ids** | Pointer to **[]int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**RootFilesOnly** | Pointer to **NullableBool** |  | [optional] 
**FileFormats** | Pointer to [**[]StorageFileFormats**](StorageFileFormats.md) |  | [optional] 
**ItemTypes** | Pointer to [**[]SIHFNZCGABTypes1**](SIHFNZCGABTypes1.md) |  | [optional] 

## Methods

### NewListItemsFiltersNullable

`func NewListItemsFiltersNullable() *ListItemsFiltersNullable`

NewListItemsFiltersNullable instantiates a new ListItemsFiltersNullable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListItemsFiltersNullableWithDefaults

`func NewListItemsFiltersNullableWithDefaults() *ListItemsFiltersNullable`

NewListItemsFiltersNullableWithDefaults instantiates a new ListItemsFiltersNullable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalIds

`func (o *ListItemsFiltersNullable) GetExternalIds() []string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *ListItemsFiltersNullable) GetExternalIdsOk() (*[]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *ListItemsFiltersNullable) SetExternalIds(v []string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *ListItemsFiltersNullable) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### SetExternalIdsNil

`func (o *ListItemsFiltersNullable) SetExternalIdsNil(b bool)`

 SetExternalIdsNil sets the value for ExternalIds to be an explicit nil

### UnsetExternalIds
`func (o *ListItemsFiltersNullable) UnsetExternalIds()`

UnsetExternalIds ensures that no value is present for ExternalIds, not even an explicit nil
### GetIds

`func (o *ListItemsFiltersNullable) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListItemsFiltersNullable) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListItemsFiltersNullable) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListItemsFiltersNullable) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *ListItemsFiltersNullable) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *ListItemsFiltersNullable) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetName

`func (o *ListItemsFiltersNullable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListItemsFiltersNullable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListItemsFiltersNullable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListItemsFiltersNullable) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ListItemsFiltersNullable) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ListItemsFiltersNullable) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRootFilesOnly

`func (o *ListItemsFiltersNullable) GetRootFilesOnly() bool`

GetRootFilesOnly returns the RootFilesOnly field if non-nil, zero value otherwise.

### GetRootFilesOnlyOk

`func (o *ListItemsFiltersNullable) GetRootFilesOnlyOk() (*bool, bool)`

GetRootFilesOnlyOk returns a tuple with the RootFilesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFilesOnly

`func (o *ListItemsFiltersNullable) SetRootFilesOnly(v bool)`

SetRootFilesOnly sets RootFilesOnly field to given value.

### HasRootFilesOnly

`func (o *ListItemsFiltersNullable) HasRootFilesOnly() bool`

HasRootFilesOnly returns a boolean if a field has been set.

### SetRootFilesOnlyNil

`func (o *ListItemsFiltersNullable) SetRootFilesOnlyNil(b bool)`

 SetRootFilesOnlyNil sets the value for RootFilesOnly to be an explicit nil

### UnsetRootFilesOnly
`func (o *ListItemsFiltersNullable) UnsetRootFilesOnly()`

UnsetRootFilesOnly ensures that no value is present for RootFilesOnly, not even an explicit nil
### GetFileFormats

`func (o *ListItemsFiltersNullable) GetFileFormats() []StorageFileFormats`

GetFileFormats returns the FileFormats field if non-nil, zero value otherwise.

### GetFileFormatsOk

`func (o *ListItemsFiltersNullable) GetFileFormatsOk() (*[]StorageFileFormats, bool)`

GetFileFormatsOk returns a tuple with the FileFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormats

`func (o *ListItemsFiltersNullable) SetFileFormats(v []StorageFileFormats)`

SetFileFormats sets FileFormats field to given value.

### HasFileFormats

`func (o *ListItemsFiltersNullable) HasFileFormats() bool`

HasFileFormats returns a boolean if a field has been set.

### SetFileFormatsNil

`func (o *ListItemsFiltersNullable) SetFileFormatsNil(b bool)`

 SetFileFormatsNil sets the value for FileFormats to be an explicit nil

### UnsetFileFormats
`func (o *ListItemsFiltersNullable) UnsetFileFormats()`

UnsetFileFormats ensures that no value is present for FileFormats, not even an explicit nil
### GetItemTypes

`func (o *ListItemsFiltersNullable) GetItemTypes() []SIHFNZCGABTypes1`

GetItemTypes returns the ItemTypes field if non-nil, zero value otherwise.

### GetItemTypesOk

`func (o *ListItemsFiltersNullable) GetItemTypesOk() (*[]SIHFNZCGABTypes1, bool)`

GetItemTypesOk returns a tuple with the ItemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTypes

`func (o *ListItemsFiltersNullable) SetItemTypes(v []SIHFNZCGABTypes1)`

SetItemTypes sets ItemTypes field to given value.

### HasItemTypes

`func (o *ListItemsFiltersNullable) HasItemTypes() bool`

HasItemTypes returns a boolean if a field has been set.

### SetItemTypesNil

`func (o *ListItemsFiltersNullable) SetItemTypesNil(b bool)`

 SetItemTypesNil sets the value for ItemTypes to be an explicit nil

### UnsetItemTypes
`func (o *ListItemsFiltersNullable) UnsetItemTypes()`

UnsetItemTypes ensures that no value is present for ItemTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


