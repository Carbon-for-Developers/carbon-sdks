# SharepointAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **interface{}** |  | 
**AccessToken** | **string** |  | 
**RefreshToken** | Pointer to **NullableString** |  | [optional] 
**TenantName** | Pointer to **NullableString** |  | [optional] 
**SiteName** | **string** |  | 

## Methods

### NewSharepointAuthentication

`func NewSharepointAuthentication(source interface{}, accessToken string, siteName string, ) *SharepointAuthentication`

NewSharepointAuthentication instantiates a new SharepointAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharepointAuthenticationWithDefaults

`func NewSharepointAuthenticationWithDefaults() *SharepointAuthentication`

NewSharepointAuthenticationWithDefaults instantiates a new SharepointAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *SharepointAuthentication) GetSource() interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SharepointAuthentication) GetSourceOk() (*interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SharepointAuthentication) SetSource(v interface{})`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *SharepointAuthentication) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *SharepointAuthentication) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetAccessToken

`func (o *SharepointAuthentication) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *SharepointAuthentication) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *SharepointAuthentication) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRefreshToken

`func (o *SharepointAuthentication) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *SharepointAuthentication) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *SharepointAuthentication) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *SharepointAuthentication) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *SharepointAuthentication) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *SharepointAuthentication) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetTenantName

`func (o *SharepointAuthentication) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *SharepointAuthentication) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *SharepointAuthentication) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *SharepointAuthentication) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### SetTenantNameNil

`func (o *SharepointAuthentication) SetTenantNameNil(b bool)`

 SetTenantNameNil sets the value for TenantName to be an explicit nil

### UnsetTenantName
`func (o *SharepointAuthentication) UnsetTenantName()`

UnsetTenantName ensures that no value is present for TenantName, not even an explicit nil
### GetSiteName

`func (o *SharepointAuthentication) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *SharepointAuthentication) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *SharepointAuthentication) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


