# ListUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**OrganizationId** | **int32** |  | 
**OrganizationSuppliedUserId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **NullableTime** |  | 
**EnabledFeatures** | **map[string]interface{}** |  | 
**CustomLimits** | **map[string]interface{}** |  | [default to {}]
**AutoSyncEnabledSources** | **[]interface{}** |  | [default to []]

## Methods

### NewListUserResponse

`func NewListUserResponse(id int32, organizationId int32, organizationSuppliedUserId string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, enabledFeatures map[string]interface{}, customLimits map[string]interface{}, autoSyncEnabledSources []interface{}, ) *ListUserResponse`

NewListUserResponse instantiates a new ListUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserResponseWithDefaults

`func NewListUserResponseWithDefaults() *ListUserResponse`

NewListUserResponseWithDefaults instantiates a new ListUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListUserResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListUserResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListUserResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *ListUserResponse) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ListUserResponse) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ListUserResponse) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationSuppliedUserId

`func (o *ListUserResponse) GetOrganizationSuppliedUserId() string`

GetOrganizationSuppliedUserId returns the OrganizationSuppliedUserId field if non-nil, zero value otherwise.

### GetOrganizationSuppliedUserIdOk

`func (o *ListUserResponse) GetOrganizationSuppliedUserIdOk() (*string, bool)`

GetOrganizationSuppliedUserIdOk returns a tuple with the OrganizationSuppliedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationSuppliedUserId

`func (o *ListUserResponse) SetOrganizationSuppliedUserId(v string)`

SetOrganizationSuppliedUserId sets OrganizationSuppliedUserId field to given value.


### GetCreatedAt

`func (o *ListUserResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListUserResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListUserResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListUserResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListUserResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListUserResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *ListUserResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ListUserResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ListUserResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *ListUserResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ListUserResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetEnabledFeatures

`func (o *ListUserResponse) GetEnabledFeatures() map[string]interface{}`

GetEnabledFeatures returns the EnabledFeatures field if non-nil, zero value otherwise.

### GetEnabledFeaturesOk

`func (o *ListUserResponse) GetEnabledFeaturesOk() (*map[string]interface{}, bool)`

GetEnabledFeaturesOk returns a tuple with the EnabledFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFeatures

`func (o *ListUserResponse) SetEnabledFeatures(v map[string]interface{})`

SetEnabledFeatures sets EnabledFeatures field to given value.


### SetEnabledFeaturesNil

`func (o *ListUserResponse) SetEnabledFeaturesNil(b bool)`

 SetEnabledFeaturesNil sets the value for EnabledFeatures to be an explicit nil

### UnsetEnabledFeatures
`func (o *ListUserResponse) UnsetEnabledFeatures()`

UnsetEnabledFeatures ensures that no value is present for EnabledFeatures, not even an explicit nil
### GetCustomLimits

`func (o *ListUserResponse) GetCustomLimits() map[string]interface{}`

GetCustomLimits returns the CustomLimits field if non-nil, zero value otherwise.

### GetCustomLimitsOk

`func (o *ListUserResponse) GetCustomLimitsOk() (*map[string]interface{}, bool)`

GetCustomLimitsOk returns a tuple with the CustomLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLimits

`func (o *ListUserResponse) SetCustomLimits(v map[string]interface{})`

SetCustomLimits sets CustomLimits field to given value.


### GetAutoSyncEnabledSources

`func (o *ListUserResponse) GetAutoSyncEnabledSources() []interface{}`

GetAutoSyncEnabledSources returns the AutoSyncEnabledSources field if non-nil, zero value otherwise.

### GetAutoSyncEnabledSourcesOk

`func (o *ListUserResponse) GetAutoSyncEnabledSourcesOk() (*[]interface{}, bool)`

GetAutoSyncEnabledSourcesOk returns a tuple with the AutoSyncEnabledSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncEnabledSources

`func (o *ListUserResponse) SetAutoSyncEnabledSources(v []interface{})`

SetAutoSyncEnabledSources sets AutoSyncEnabledSources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


