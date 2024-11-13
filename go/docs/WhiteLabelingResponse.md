# WhiteLabelingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoveBranding** | **bool** |  | 
**Integrations** | **map[string]interface{}** |  | 
**CustomLimits** | **map[string]interface{}** |  | 
**ConnectorSettings** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]

## Methods

### NewWhiteLabelingResponse

`func NewWhiteLabelingResponse(removeBranding bool, integrations map[string]interface{}, customLimits map[string]interface{}, ) *WhiteLabelingResponse`

NewWhiteLabelingResponse instantiates a new WhiteLabelingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhiteLabelingResponseWithDefaults

`func NewWhiteLabelingResponseWithDefaults() *WhiteLabelingResponse`

NewWhiteLabelingResponseWithDefaults instantiates a new WhiteLabelingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoveBranding

`func (o *WhiteLabelingResponse) GetRemoveBranding() bool`

GetRemoveBranding returns the RemoveBranding field if non-nil, zero value otherwise.

### GetRemoveBrandingOk

`func (o *WhiteLabelingResponse) GetRemoveBrandingOk() (*bool, bool)`

GetRemoveBrandingOk returns a tuple with the RemoveBranding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveBranding

`func (o *WhiteLabelingResponse) SetRemoveBranding(v bool)`

SetRemoveBranding sets RemoveBranding field to given value.


### GetIntegrations

`func (o *WhiteLabelingResponse) GetIntegrations() map[string]interface{}`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *WhiteLabelingResponse) GetIntegrationsOk() (*map[string]interface{}, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *WhiteLabelingResponse) SetIntegrations(v map[string]interface{})`

SetIntegrations sets Integrations field to given value.


### GetCustomLimits

`func (o *WhiteLabelingResponse) GetCustomLimits() map[string]interface{}`

GetCustomLimits returns the CustomLimits field if non-nil, zero value otherwise.

### GetCustomLimitsOk

`func (o *WhiteLabelingResponse) GetCustomLimitsOk() (*map[string]interface{}, bool)`

GetCustomLimitsOk returns a tuple with the CustomLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLimits

`func (o *WhiteLabelingResponse) SetCustomLimits(v map[string]interface{})`

SetCustomLimits sets CustomLimits field to given value.


### GetConnectorSettings

`func (o *WhiteLabelingResponse) GetConnectorSettings() map[string]interface{}`

GetConnectorSettings returns the ConnectorSettings field if non-nil, zero value otherwise.

### GetConnectorSettingsOk

`func (o *WhiteLabelingResponse) GetConnectorSettingsOk() (*map[string]interface{}, bool)`

GetConnectorSettingsOk returns a tuple with the ConnectorSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorSettings

`func (o *WhiteLabelingResponse) SetConnectorSettings(v map[string]interface{})`

SetConnectorSettings sets ConnectorSettings field to given value.

### HasConnectorSettings

`func (o *WhiteLabelingResponse) HasConnectorSettings() bool`

HasConnectorSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


