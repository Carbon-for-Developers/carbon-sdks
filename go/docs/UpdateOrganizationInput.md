# UpdateOrganizationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalUserConfig** | Pointer to [**NullableUserConfigurationNullable**](UserConfigurationNullable.md) |  | [optional] 
**DataSourceConfigs** | Pointer to [**map[string]DataSourceConfiguration**](DataSourceConfiguration.md) | Used to set organization level defaults for configuration related to data sources. | [optional] 

## Methods

### NewUpdateOrganizationInput

`func NewUpdateOrganizationInput() *UpdateOrganizationInput`

NewUpdateOrganizationInput instantiates a new UpdateOrganizationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationInputWithDefaults

`func NewUpdateOrganizationInputWithDefaults() *UpdateOrganizationInput`

NewUpdateOrganizationInputWithDefaults instantiates a new UpdateOrganizationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalUserConfig

`func (o *UpdateOrganizationInput) GetGlobalUserConfig() UserConfigurationNullable`

GetGlobalUserConfig returns the GlobalUserConfig field if non-nil, zero value otherwise.

### GetGlobalUserConfigOk

`func (o *UpdateOrganizationInput) GetGlobalUserConfigOk() (*UserConfigurationNullable, bool)`

GetGlobalUserConfigOk returns a tuple with the GlobalUserConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalUserConfig

`func (o *UpdateOrganizationInput) SetGlobalUserConfig(v UserConfigurationNullable)`

SetGlobalUserConfig sets GlobalUserConfig field to given value.

### HasGlobalUserConfig

`func (o *UpdateOrganizationInput) HasGlobalUserConfig() bool`

HasGlobalUserConfig returns a boolean if a field has been set.

### SetGlobalUserConfigNil

`func (o *UpdateOrganizationInput) SetGlobalUserConfigNil(b bool)`

 SetGlobalUserConfigNil sets the value for GlobalUserConfig to be an explicit nil

### UnsetGlobalUserConfig
`func (o *UpdateOrganizationInput) UnsetGlobalUserConfig()`

UnsetGlobalUserConfig ensures that no value is present for GlobalUserConfig, not even an explicit nil
### GetDataSourceConfigs

`func (o *UpdateOrganizationInput) GetDataSourceConfigs() map[string]DataSourceConfiguration`

GetDataSourceConfigs returns the DataSourceConfigs field if non-nil, zero value otherwise.

### GetDataSourceConfigsOk

`func (o *UpdateOrganizationInput) GetDataSourceConfigsOk() (*map[string]DataSourceConfiguration, bool)`

GetDataSourceConfigsOk returns a tuple with the DataSourceConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceConfigs

`func (o *UpdateOrganizationInput) SetDataSourceConfigs(v map[string]DataSourceConfiguration)`

SetDataSourceConfigs sets DataSourceConfigs field to given value.

### HasDataSourceConfigs

`func (o *UpdateOrganizationInput) HasDataSourceConfigs() bool`

HasDataSourceConfigs returns a boolean if a field has been set.

### SetDataSourceConfigsNil

`func (o *UpdateOrganizationInput) SetDataSourceConfigsNil(b bool)`

 SetDataSourceConfigsNil sets the value for DataSourceConfigs to be an explicit nil

### UnsetDataSourceConfigs
`func (o *UpdateOrganizationInput) UnsetDataSourceConfigs()`

UnsetDataSourceConfigs ensures that no value is present for DataSourceConfigs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


