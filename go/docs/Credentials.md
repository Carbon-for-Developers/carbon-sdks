# Credentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**RedirectUri** | **string** |  | 
**ClientSecret** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCredentials

`func NewCredentials(clientId string, redirectUri string, ) *Credentials`

NewCredentials instantiates a new Credentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsWithDefaults

`func NewCredentialsWithDefaults() *Credentials`

NewCredentialsWithDefaults instantiates a new Credentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *Credentials) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Credentials) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Credentials) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetRedirectUri

`func (o *Credentials) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *Credentials) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *Credentials) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetClientSecret

`func (o *Credentials) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Credentials) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Credentials) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *Credentials) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *Credentials) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *Credentials) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


