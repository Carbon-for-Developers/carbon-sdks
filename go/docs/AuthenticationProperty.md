# AuthenticationProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **interface{}** |  | 
**AccessToken** | **string** |  | 
**RefreshToken** | Pointer to **NullableString** |  | [optional] 
**WorkspaceId** | **string** |  | 
**TenantName** | **string** |  | 
**SiteName** | **string** |  | 
**Subdomain** | **string** |  | 
**AccessTokenSecret** | **string** |  | 
**Username** | **string** |  | 
**ZoteroId** | **string** |  | 
**OrganizationName** | **string** |  | 
**Domain** | **string** |  | 
**ApiKey** | **string** |  | 
**AccessKey** | **string** |  | 
**AccessKeySecret** | **string** |  | 

## Methods

### NewAuthenticationProperty

`func NewAuthenticationProperty(source interface{}, accessToken string, workspaceId string, tenantName string, siteName string, subdomain string, accessTokenSecret string, username string, zoteroId string, organizationName string, domain string, apiKey string, accessKey string, accessKeySecret string, ) *AuthenticationProperty`

NewAuthenticationProperty instantiates a new AuthenticationProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationPropertyWithDefaults

`func NewAuthenticationPropertyWithDefaults() *AuthenticationProperty`

NewAuthenticationPropertyWithDefaults instantiates a new AuthenticationProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *AuthenticationProperty) GetSource() interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AuthenticationProperty) GetSourceOk() (*interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AuthenticationProperty) SetSource(v interface{})`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *AuthenticationProperty) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AuthenticationProperty) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetAccessToken

`func (o *AuthenticationProperty) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthenticationProperty) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthenticationProperty) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRefreshToken

`func (o *AuthenticationProperty) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *AuthenticationProperty) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *AuthenticationProperty) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *AuthenticationProperty) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *AuthenticationProperty) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *AuthenticationProperty) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetWorkspaceId

`func (o *AuthenticationProperty) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *AuthenticationProperty) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *AuthenticationProperty) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetTenantName

`func (o *AuthenticationProperty) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *AuthenticationProperty) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *AuthenticationProperty) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.


### GetSiteName

`func (o *AuthenticationProperty) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *AuthenticationProperty) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *AuthenticationProperty) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.


### GetSubdomain

`func (o *AuthenticationProperty) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *AuthenticationProperty) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *AuthenticationProperty) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.


### GetAccessTokenSecret

`func (o *AuthenticationProperty) GetAccessTokenSecret() string`

GetAccessTokenSecret returns the AccessTokenSecret field if non-nil, zero value otherwise.

### GetAccessTokenSecretOk

`func (o *AuthenticationProperty) GetAccessTokenSecretOk() (*string, bool)`

GetAccessTokenSecretOk returns a tuple with the AccessTokenSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenSecret

`func (o *AuthenticationProperty) SetAccessTokenSecret(v string)`

SetAccessTokenSecret sets AccessTokenSecret field to given value.


### GetUsername

`func (o *AuthenticationProperty) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AuthenticationProperty) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AuthenticationProperty) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetZoteroId

`func (o *AuthenticationProperty) GetZoteroId() string`

GetZoteroId returns the ZoteroId field if non-nil, zero value otherwise.

### GetZoteroIdOk

`func (o *AuthenticationProperty) GetZoteroIdOk() (*string, bool)`

GetZoteroIdOk returns a tuple with the ZoteroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoteroId

`func (o *AuthenticationProperty) SetZoteroId(v string)`

SetZoteroId sets ZoteroId field to given value.


### GetOrganizationName

`func (o *AuthenticationProperty) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AuthenticationProperty) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AuthenticationProperty) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetDomain

`func (o *AuthenticationProperty) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AuthenticationProperty) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AuthenticationProperty) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetApiKey

`func (o *AuthenticationProperty) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AuthenticationProperty) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AuthenticationProperty) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetAccessKey

`func (o *AuthenticationProperty) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AuthenticationProperty) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AuthenticationProperty) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetAccessKeySecret

`func (o *AuthenticationProperty) GetAccessKeySecret() string`

GetAccessKeySecret returns the AccessKeySecret field if non-nil, zero value otherwise.

### GetAccessKeySecretOk

`func (o *AuthenticationProperty) GetAccessKeySecretOk() (*string, bool)`

GetAccessKeySecretOk returns a tuple with the AccessKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeySecret

`func (o *AuthenticationProperty) SetAccessKeySecret(v string)`

SetAccessKeySecret sets AccessKeySecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


