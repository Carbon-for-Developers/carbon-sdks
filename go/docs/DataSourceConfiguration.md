# DataSourceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedFileFormats** | Pointer to **[]string** | List of allowed file formats for the organization. An empty list means all file formats are allowed. | [optional] [default to []]
**GoogleWorkspaceDocsSaveAs** | Pointer to **NullableString** | Used to set the format Google Workspace documents (Docs, Sheets, Slides) are stored as for raw files uploads in S3. The default format is TXT | [optional] 

## Methods

### NewDataSourceConfiguration

`func NewDataSourceConfiguration() *DataSourceConfiguration`

NewDataSourceConfiguration instantiates a new DataSourceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceConfigurationWithDefaults

`func NewDataSourceConfigurationWithDefaults() *DataSourceConfiguration`

NewDataSourceConfigurationWithDefaults instantiates a new DataSourceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedFileFormats

`func (o *DataSourceConfiguration) GetAllowedFileFormats() []string`

GetAllowedFileFormats returns the AllowedFileFormats field if non-nil, zero value otherwise.

### GetAllowedFileFormatsOk

`func (o *DataSourceConfiguration) GetAllowedFileFormatsOk() (*[]string, bool)`

GetAllowedFileFormatsOk returns a tuple with the AllowedFileFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedFileFormats

`func (o *DataSourceConfiguration) SetAllowedFileFormats(v []string)`

SetAllowedFileFormats sets AllowedFileFormats field to given value.

### HasAllowedFileFormats

`func (o *DataSourceConfiguration) HasAllowedFileFormats() bool`

HasAllowedFileFormats returns a boolean if a field has been set.

### GetGoogleWorkspaceDocsSaveAs

`func (o *DataSourceConfiguration) GetGoogleWorkspaceDocsSaveAs() string`

GetGoogleWorkspaceDocsSaveAs returns the GoogleWorkspaceDocsSaveAs field if non-nil, zero value otherwise.

### GetGoogleWorkspaceDocsSaveAsOk

`func (o *DataSourceConfiguration) GetGoogleWorkspaceDocsSaveAsOk() (*string, bool)`

GetGoogleWorkspaceDocsSaveAsOk returns a tuple with the GoogleWorkspaceDocsSaveAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleWorkspaceDocsSaveAs

`func (o *DataSourceConfiguration) SetGoogleWorkspaceDocsSaveAs(v string)`

SetGoogleWorkspaceDocsSaveAs sets GoogleWorkspaceDocsSaveAs field to given value.

### HasGoogleWorkspaceDocsSaveAs

`func (o *DataSourceConfiguration) HasGoogleWorkspaceDocsSaveAs() bool`

HasGoogleWorkspaceDocsSaveAs returns a boolean if a field has been set.

### SetGoogleWorkspaceDocsSaveAsNil

`func (o *DataSourceConfiguration) SetGoogleWorkspaceDocsSaveAsNil(b bool)`

 SetGoogleWorkspaceDocsSaveAsNil sets the value for GoogleWorkspaceDocsSaveAs to be an explicit nil

### UnsetGoogleWorkspaceDocsSaveAs
`func (o *DataSourceConfiguration) UnsetGoogleWorkspaceDocsSaveAs()`

UnsetGoogleWorkspaceDocsSaveAs ensures that no value is present for GoogleWorkspaceDocsSaveAs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


