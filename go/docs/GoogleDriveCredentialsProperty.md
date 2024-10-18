# GoogleDriveCredentialsProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**RedirectUri** | **string** |  | 
**ClientSecret** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**FilePickerClientId** | Pointer to **NullableString** |  | [optional] 
**FilePickerRedirectUri** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGoogleDriveCredentialsProperty

`func NewGoogleDriveCredentialsProperty(clientId string, redirectUri string, ) *GoogleDriveCredentialsProperty`

NewGoogleDriveCredentialsProperty instantiates a new GoogleDriveCredentialsProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleDriveCredentialsPropertyWithDefaults

`func NewGoogleDriveCredentialsPropertyWithDefaults() *GoogleDriveCredentialsProperty`

NewGoogleDriveCredentialsPropertyWithDefaults instantiates a new GoogleDriveCredentialsProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *GoogleDriveCredentialsProperty) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GoogleDriveCredentialsProperty) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GoogleDriveCredentialsProperty) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetRedirectUri

`func (o *GoogleDriveCredentialsProperty) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *GoogleDriveCredentialsProperty) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *GoogleDriveCredentialsProperty) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetClientSecret

`func (o *GoogleDriveCredentialsProperty) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *GoogleDriveCredentialsProperty) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *GoogleDriveCredentialsProperty) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *GoogleDriveCredentialsProperty) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *GoogleDriveCredentialsProperty) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *GoogleDriveCredentialsProperty) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetApiKey

`func (o *GoogleDriveCredentialsProperty) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GoogleDriveCredentialsProperty) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GoogleDriveCredentialsProperty) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GoogleDriveCredentialsProperty) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *GoogleDriveCredentialsProperty) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *GoogleDriveCredentialsProperty) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetFilePickerClientId

`func (o *GoogleDriveCredentialsProperty) GetFilePickerClientId() string`

GetFilePickerClientId returns the FilePickerClientId field if non-nil, zero value otherwise.

### GetFilePickerClientIdOk

`func (o *GoogleDriveCredentialsProperty) GetFilePickerClientIdOk() (*string, bool)`

GetFilePickerClientIdOk returns a tuple with the FilePickerClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePickerClientId

`func (o *GoogleDriveCredentialsProperty) SetFilePickerClientId(v string)`

SetFilePickerClientId sets FilePickerClientId field to given value.

### HasFilePickerClientId

`func (o *GoogleDriveCredentialsProperty) HasFilePickerClientId() bool`

HasFilePickerClientId returns a boolean if a field has been set.

### SetFilePickerClientIdNil

`func (o *GoogleDriveCredentialsProperty) SetFilePickerClientIdNil(b bool)`

 SetFilePickerClientIdNil sets the value for FilePickerClientId to be an explicit nil

### UnsetFilePickerClientId
`func (o *GoogleDriveCredentialsProperty) UnsetFilePickerClientId()`

UnsetFilePickerClientId ensures that no value is present for FilePickerClientId, not even an explicit nil
### GetFilePickerRedirectUri

`func (o *GoogleDriveCredentialsProperty) GetFilePickerRedirectUri() string`

GetFilePickerRedirectUri returns the FilePickerRedirectUri field if non-nil, zero value otherwise.

### GetFilePickerRedirectUriOk

`func (o *GoogleDriveCredentialsProperty) GetFilePickerRedirectUriOk() (*string, bool)`

GetFilePickerRedirectUriOk returns a tuple with the FilePickerRedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePickerRedirectUri

`func (o *GoogleDriveCredentialsProperty) SetFilePickerRedirectUri(v string)`

SetFilePickerRedirectUri sets FilePickerRedirectUri field to given value.

### HasFilePickerRedirectUri

`func (o *GoogleDriveCredentialsProperty) HasFilePickerRedirectUri() bool`

HasFilePickerRedirectUri returns a boolean if a field has been set.

### SetFilePickerRedirectUriNil

`func (o *GoogleDriveCredentialsProperty) SetFilePickerRedirectUriNil(b bool)`

 SetFilePickerRedirectUriNil sets the value for FilePickerRedirectUri to be an explicit nil

### UnsetFilePickerRedirectUri
`func (o *GoogleDriveCredentialsProperty) UnsetFilePickerRedirectUri()`

UnsetFilePickerRedirectUri ensures that no value is present for FilePickerRedirectUri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


