# ServiceNowAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**AccessToken** | **string** |  | 
**RefreshToken** | Pointer to **NullableString** |  | [optional] 
**InstanceSubdomain** | **string** |  | 
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**RedirectUri** | **string** |  | 

## Methods

### NewServiceNowAuthentication

`func NewServiceNowAuthentication(source string, accessToken string, instanceSubdomain string, clientId string, clientSecret string, redirectUri string, ) *ServiceNowAuthentication`

NewServiceNowAuthentication instantiates a new ServiceNowAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceNowAuthenticationWithDefaults

`func NewServiceNowAuthenticationWithDefaults() *ServiceNowAuthentication`

NewServiceNowAuthenticationWithDefaults instantiates a new ServiceNowAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ServiceNowAuthentication) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ServiceNowAuthentication) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ServiceNowAuthentication) SetSource(v string)`

SetSource sets Source field to given value.


### GetAccessToken

`func (o *ServiceNowAuthentication) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ServiceNowAuthentication) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ServiceNowAuthentication) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRefreshToken

`func (o *ServiceNowAuthentication) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *ServiceNowAuthentication) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *ServiceNowAuthentication) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *ServiceNowAuthentication) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *ServiceNowAuthentication) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *ServiceNowAuthentication) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetInstanceSubdomain

`func (o *ServiceNowAuthentication) GetInstanceSubdomain() string`

GetInstanceSubdomain returns the InstanceSubdomain field if non-nil, zero value otherwise.

### GetInstanceSubdomainOk

`func (o *ServiceNowAuthentication) GetInstanceSubdomainOk() (*string, bool)`

GetInstanceSubdomainOk returns a tuple with the InstanceSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSubdomain

`func (o *ServiceNowAuthentication) SetInstanceSubdomain(v string)`

SetInstanceSubdomain sets InstanceSubdomain field to given value.


### GetClientId

`func (o *ServiceNowAuthentication) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ServiceNowAuthentication) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ServiceNowAuthentication) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *ServiceNowAuthentication) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ServiceNowAuthentication) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ServiceNowAuthentication) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetRedirectUri

`func (o *ServiceNowAuthentication) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *ServiceNowAuthentication) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *ServiceNowAuthentication) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


