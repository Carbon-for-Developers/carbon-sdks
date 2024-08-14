# ModifyColdStorageParametersQueryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**OrganizationUserFilesToSyncFilters**](OrganizationUserFilesToSyncFilters.md) |  | [optional] 
**EnableColdStorage** | Pointer to **NullableBool** |  | [optional] 
**HotStorageTimeToLive** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewModifyColdStorageParametersQueryInput

`func NewModifyColdStorageParametersQueryInput() *ModifyColdStorageParametersQueryInput`

NewModifyColdStorageParametersQueryInput instantiates a new ModifyColdStorageParametersQueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyColdStorageParametersQueryInputWithDefaults

`func NewModifyColdStorageParametersQueryInputWithDefaults() *ModifyColdStorageParametersQueryInput`

NewModifyColdStorageParametersQueryInputWithDefaults instantiates a new ModifyColdStorageParametersQueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ModifyColdStorageParametersQueryInput) GetFilters() OrganizationUserFilesToSyncFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ModifyColdStorageParametersQueryInput) GetFiltersOk() (*OrganizationUserFilesToSyncFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ModifyColdStorageParametersQueryInput) SetFilters(v OrganizationUserFilesToSyncFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ModifyColdStorageParametersQueryInput) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetEnableColdStorage

`func (o *ModifyColdStorageParametersQueryInput) GetEnableColdStorage() bool`

GetEnableColdStorage returns the EnableColdStorage field if non-nil, zero value otherwise.

### GetEnableColdStorageOk

`func (o *ModifyColdStorageParametersQueryInput) GetEnableColdStorageOk() (*bool, bool)`

GetEnableColdStorageOk returns a tuple with the EnableColdStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableColdStorage

`func (o *ModifyColdStorageParametersQueryInput) SetEnableColdStorage(v bool)`

SetEnableColdStorage sets EnableColdStorage field to given value.

### HasEnableColdStorage

`func (o *ModifyColdStorageParametersQueryInput) HasEnableColdStorage() bool`

HasEnableColdStorage returns a boolean if a field has been set.

### SetEnableColdStorageNil

`func (o *ModifyColdStorageParametersQueryInput) SetEnableColdStorageNil(b bool)`

 SetEnableColdStorageNil sets the value for EnableColdStorage to be an explicit nil

### UnsetEnableColdStorage
`func (o *ModifyColdStorageParametersQueryInput) UnsetEnableColdStorage()`

UnsetEnableColdStorage ensures that no value is present for EnableColdStorage, not even an explicit nil
### GetHotStorageTimeToLive

`func (o *ModifyColdStorageParametersQueryInput) GetHotStorageTimeToLive() int32`

GetHotStorageTimeToLive returns the HotStorageTimeToLive field if non-nil, zero value otherwise.

### GetHotStorageTimeToLiveOk

`func (o *ModifyColdStorageParametersQueryInput) GetHotStorageTimeToLiveOk() (*int32, bool)`

GetHotStorageTimeToLiveOk returns a tuple with the HotStorageTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotStorageTimeToLive

`func (o *ModifyColdStorageParametersQueryInput) SetHotStorageTimeToLive(v int32)`

SetHotStorageTimeToLive sets HotStorageTimeToLive field to given value.

### HasHotStorageTimeToLive

`func (o *ModifyColdStorageParametersQueryInput) HasHotStorageTimeToLive() bool`

HasHotStorageTimeToLive returns a boolean if a field has been set.

### SetHotStorageTimeToLiveNil

`func (o *ModifyColdStorageParametersQueryInput) SetHotStorageTimeToLiveNil(b bool)`

 SetHotStorageTimeToLiveNil sets the value for HotStorageTimeToLive to be an explicit nil

### UnsetHotStorageTimeToLive
`func (o *ModifyColdStorageParametersQueryInput) UnsetHotStorageTimeToLive()`

UnsetHotStorageTimeToLive ensures that no value is present for HotStorageTimeToLive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


