# DeleteFilesV2QueryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**OrganizationUserFilesToSyncFilters**](OrganizationUserFilesToSyncFilters.md) |  | [optional] 
**SendWebhook** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewDeleteFilesV2QueryInput

`func NewDeleteFilesV2QueryInput() *DeleteFilesV2QueryInput`

NewDeleteFilesV2QueryInput instantiates a new DeleteFilesV2QueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteFilesV2QueryInputWithDefaults

`func NewDeleteFilesV2QueryInputWithDefaults() *DeleteFilesV2QueryInput`

NewDeleteFilesV2QueryInputWithDefaults instantiates a new DeleteFilesV2QueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *DeleteFilesV2QueryInput) GetFilters() OrganizationUserFilesToSyncFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DeleteFilesV2QueryInput) GetFiltersOk() (*OrganizationUserFilesToSyncFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DeleteFilesV2QueryInput) SetFilters(v OrganizationUserFilesToSyncFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DeleteFilesV2QueryInput) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSendWebhook

`func (o *DeleteFilesV2QueryInput) GetSendWebhook() bool`

GetSendWebhook returns the SendWebhook field if non-nil, zero value otherwise.

### GetSendWebhookOk

`func (o *DeleteFilesV2QueryInput) GetSendWebhookOk() (*bool, bool)`

GetSendWebhookOk returns a tuple with the SendWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendWebhook

`func (o *DeleteFilesV2QueryInput) SetSendWebhook(v bool)`

SetSendWebhook sets SendWebhook field to given value.

### HasSendWebhook

`func (o *DeleteFilesV2QueryInput) HasSendWebhook() bool`

HasSendWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


