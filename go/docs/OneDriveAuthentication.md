# OneDriveAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **interface{}** |  | 
**AccessToken** | **string** |  | 
**RefreshToken** | Pointer to **NullableString** |  | [optional] 
**TenantName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOneDriveAuthentication

`func NewOneDriveAuthentication(source interface{}, accessToken string, ) *OneDriveAuthentication`

NewOneDriveAuthentication instantiates a new OneDriveAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneDriveAuthenticationWithDefaults

`func NewOneDriveAuthenticationWithDefaults() *OneDriveAuthentication`

NewOneDriveAuthenticationWithDefaults instantiates a new OneDriveAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *OneDriveAuthentication) GetSource() interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OneDriveAuthentication) GetSourceOk() (*interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OneDriveAuthentication) SetSource(v interface{})`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *OneDriveAuthentication) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *OneDriveAuthentication) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetAccessToken

`func (o *OneDriveAuthentication) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OneDriveAuthentication) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OneDriveAuthentication) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRefreshToken

`func (o *OneDriveAuthentication) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OneDriveAuthentication) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OneDriveAuthentication) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OneDriveAuthentication) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *OneDriveAuthentication) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *OneDriveAuthentication) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetTenantName

`func (o *OneDriveAuthentication) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *OneDriveAuthentication) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *OneDriveAuthentication) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *OneDriveAuthentication) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### SetTenantNameNil

`func (o *OneDriveAuthentication) SetTenantNameNil(b bool)`

 SetTenantNameNil sets the value for TenantName to be an explicit nil

### UnsetTenantName
`func (o *OneDriveAuthentication) UnsetTenantName()`

UnsetTenantName ensures that no value is present for TenantName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


