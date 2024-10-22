# ColdStorageProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableColdStorage** | Pointer to **bool** | Enable cold storage for the file. If set to true, the file will be moved to cold storage after a certain period of inactivity. Default is false. | [optional] [default to false]
**HotStorageTimeToLive** | Pointer to **NullableInt32** | Time in days after which the file will be moved to cold storage. Valid values are [1, 3, 7, 14, 30] | [optional] 

## Methods

### NewColdStorageProps

`func NewColdStorageProps() *ColdStorageProps`

NewColdStorageProps instantiates a new ColdStorageProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColdStoragePropsWithDefaults

`func NewColdStoragePropsWithDefaults() *ColdStorageProps`

NewColdStoragePropsWithDefaults instantiates a new ColdStorageProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableColdStorage

`func (o *ColdStorageProps) GetEnableColdStorage() bool`

GetEnableColdStorage returns the EnableColdStorage field if non-nil, zero value otherwise.

### GetEnableColdStorageOk

`func (o *ColdStorageProps) GetEnableColdStorageOk() (*bool, bool)`

GetEnableColdStorageOk returns a tuple with the EnableColdStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableColdStorage

`func (o *ColdStorageProps) SetEnableColdStorage(v bool)`

SetEnableColdStorage sets EnableColdStorage field to given value.

### HasEnableColdStorage

`func (o *ColdStorageProps) HasEnableColdStorage() bool`

HasEnableColdStorage returns a boolean if a field has been set.

### GetHotStorageTimeToLive

`func (o *ColdStorageProps) GetHotStorageTimeToLive() int32`

GetHotStorageTimeToLive returns the HotStorageTimeToLive field if non-nil, zero value otherwise.

### GetHotStorageTimeToLiveOk

`func (o *ColdStorageProps) GetHotStorageTimeToLiveOk() (*int32, bool)`

GetHotStorageTimeToLiveOk returns a tuple with the HotStorageTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotStorageTimeToLive

`func (o *ColdStorageProps) SetHotStorageTimeToLive(v int32)`

SetHotStorageTimeToLive sets HotStorageTimeToLive field to given value.

### HasHotStorageTimeToLive

`func (o *ColdStorageProps) HasHotStorageTimeToLive() bool`

HasHotStorageTimeToLive returns a boolean if a field has been set.

### SetHotStorageTimeToLiveNil

`func (o *ColdStorageProps) SetHotStorageTimeToLiveNil(b bool)`

 SetHotStorageTimeToLiveNil sets the value for HotStorageTimeToLive to be an explicit nil

### UnsetHotStorageTimeToLive
`func (o *ColdStorageProps) UnsetHotStorageTimeToLive()`

UnsetHotStorageTimeToLive ensures that no value is present for HotStorageTimeToLive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


