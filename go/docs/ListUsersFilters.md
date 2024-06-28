# ListUsersFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerIds** | Pointer to **[]string** | Filter by a list of customer IDs | [optional] 
**CarbonUserIds** | Pointer to **[]int32** | Filter by a list of carbon user IDs | [optional] 

## Methods

### NewListUsersFilters

`func NewListUsersFilters() *ListUsersFilters`

NewListUsersFilters instantiates a new ListUsersFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsersFiltersWithDefaults

`func NewListUsersFiltersWithDefaults() *ListUsersFilters`

NewListUsersFiltersWithDefaults instantiates a new ListUsersFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerIds

`func (o *ListUsersFilters) GetCustomerIds() []string`

GetCustomerIds returns the CustomerIds field if non-nil, zero value otherwise.

### GetCustomerIdsOk

`func (o *ListUsersFilters) GetCustomerIdsOk() (*[]string, bool)`

GetCustomerIdsOk returns a tuple with the CustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIds

`func (o *ListUsersFilters) SetCustomerIds(v []string)`

SetCustomerIds sets CustomerIds field to given value.

### HasCustomerIds

`func (o *ListUsersFilters) HasCustomerIds() bool`

HasCustomerIds returns a boolean if a field has been set.

### SetCustomerIdsNil

`func (o *ListUsersFilters) SetCustomerIdsNil(b bool)`

 SetCustomerIdsNil sets the value for CustomerIds to be an explicit nil

### UnsetCustomerIds
`func (o *ListUsersFilters) UnsetCustomerIds()`

UnsetCustomerIds ensures that no value is present for CustomerIds, not even an explicit nil
### GetCarbonUserIds

`func (o *ListUsersFilters) GetCarbonUserIds() []int32`

GetCarbonUserIds returns the CarbonUserIds field if non-nil, zero value otherwise.

### GetCarbonUserIdsOk

`func (o *ListUsersFilters) GetCarbonUserIdsOk() (*[]int32, bool)`

GetCarbonUserIdsOk returns a tuple with the CarbonUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarbonUserIds

`func (o *ListUsersFilters) SetCarbonUserIds(v []int32)`

SetCarbonUserIds sets CarbonUserIds field to given value.

### HasCarbonUserIds

`func (o *ListUsersFilters) HasCarbonUserIds() bool`

HasCarbonUserIds returns a boolean if a field has been set.

### SetCarbonUserIdsNil

`func (o *ListUsersFilters) SetCarbonUserIdsNil(b bool)`

 SetCarbonUserIdsNil sets the value for CarbonUserIds to be an explicit nil

### UnsetCarbonUserIds
`func (o *ListUsersFilters) UnsetCarbonUserIds()`

UnsetCarbonUserIds ensures that no value is present for CarbonUserIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


