# OANSCZGF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**OAuthAuthenticationSource**](OAuthAuthenticationSource.md) |  | 
**AccessToken** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **NullableString** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 
**TenantName** | Pointer to **string** |  | [optional] 
**SiteName** | Pointer to **string** |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**AccessTokenSecret** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**ZoteroId** | Pointer to **string** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**AccessKey** | Pointer to **string** |  | [optional] 
**AccessKeySecret** | Pointer to **string** |  | [optional] 
**EndpointUrl** | Pointer to **NullableString** | You can specify a Digital Ocean endpoint URL to connect a Digital Ocean Space through this endpoint.         The URL should be of format &lt;region&gt;.digitaloceanspaces.com. It&#39;s not required for S3 buckets. | [optional] 
**InstanceSubdomain** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**RedirectUri** | Pointer to **string** |  | [optional] 
**GongAccountEmail** | Pointer to **string** |  | [optional] 

## Methods

### NewOANSCZGF

`func NewOANSCZGF(source OAuthAuthenticationSource, ) *OANSCZGF`

NewOANSCZGF instantiates a new OANSCZGF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOANSCZGFWithDefaults

`func NewOANSCZGFWithDefaults() *OANSCZGF`

NewOANSCZGFWithDefaults instantiates a new OANSCZGF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *OANSCZGF) GetSource() OAuthAuthenticationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OANSCZGF) GetSourceOk() (*OAuthAuthenticationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OANSCZGF) SetSource(v OAuthAuthenticationSource)`

SetSource sets Source field to given value.


### GetAccessToken

`func (o *OANSCZGF) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OANSCZGF) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OANSCZGF) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *OANSCZGF) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *OANSCZGF) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OANSCZGF) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OANSCZGF) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OANSCZGF) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *OANSCZGF) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *OANSCZGF) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetWorkspaceId

`func (o *OANSCZGF) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *OANSCZGF) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *OANSCZGF) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *OANSCZGF) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetTenantName

`func (o *OANSCZGF) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *OANSCZGF) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *OANSCZGF) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *OANSCZGF) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### GetSiteName

`func (o *OANSCZGF) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *OANSCZGF) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *OANSCZGF) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *OANSCZGF) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSubdomain

`func (o *OANSCZGF) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *OANSCZGF) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *OANSCZGF) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *OANSCZGF) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetAccessTokenSecret

`func (o *OANSCZGF) GetAccessTokenSecret() string`

GetAccessTokenSecret returns the AccessTokenSecret field if non-nil, zero value otherwise.

### GetAccessTokenSecretOk

`func (o *OANSCZGF) GetAccessTokenSecretOk() (*string, bool)`

GetAccessTokenSecretOk returns a tuple with the AccessTokenSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenSecret

`func (o *OANSCZGF) SetAccessTokenSecret(v string)`

SetAccessTokenSecret sets AccessTokenSecret field to given value.

### HasAccessTokenSecret

`func (o *OANSCZGF) HasAccessTokenSecret() bool`

HasAccessTokenSecret returns a boolean if a field has been set.

### GetUsername

`func (o *OANSCZGF) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OANSCZGF) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OANSCZGF) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OANSCZGF) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetZoteroId

`func (o *OANSCZGF) GetZoteroId() string`

GetZoteroId returns the ZoteroId field if non-nil, zero value otherwise.

### GetZoteroIdOk

`func (o *OANSCZGF) GetZoteroIdOk() (*string, bool)`

GetZoteroIdOk returns a tuple with the ZoteroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoteroId

`func (o *OANSCZGF) SetZoteroId(v string)`

SetZoteroId sets ZoteroId field to given value.

### HasZoteroId

`func (o *OANSCZGF) HasZoteroId() bool`

HasZoteroId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *OANSCZGF) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OANSCZGF) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OANSCZGF) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *OANSCZGF) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetDomain

`func (o *OANSCZGF) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OANSCZGF) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OANSCZGF) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *OANSCZGF) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetApiKey

`func (o *OANSCZGF) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OANSCZGF) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OANSCZGF) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *OANSCZGF) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAccessKey

`func (o *OANSCZGF) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *OANSCZGF) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *OANSCZGF) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *OANSCZGF) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessKeySecret

`func (o *OANSCZGF) GetAccessKeySecret() string`

GetAccessKeySecret returns the AccessKeySecret field if non-nil, zero value otherwise.

### GetAccessKeySecretOk

`func (o *OANSCZGF) GetAccessKeySecretOk() (*string, bool)`

GetAccessKeySecretOk returns a tuple with the AccessKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeySecret

`func (o *OANSCZGF) SetAccessKeySecret(v string)`

SetAccessKeySecret sets AccessKeySecret field to given value.

### HasAccessKeySecret

`func (o *OANSCZGF) HasAccessKeySecret() bool`

HasAccessKeySecret returns a boolean if a field has been set.

### GetEndpointUrl

`func (o *OANSCZGF) GetEndpointUrl() string`

GetEndpointUrl returns the EndpointUrl field if non-nil, zero value otherwise.

### GetEndpointUrlOk

`func (o *OANSCZGF) GetEndpointUrlOk() (*string, bool)`

GetEndpointUrlOk returns a tuple with the EndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUrl

`func (o *OANSCZGF) SetEndpointUrl(v string)`

SetEndpointUrl sets EndpointUrl field to given value.

### HasEndpointUrl

`func (o *OANSCZGF) HasEndpointUrl() bool`

HasEndpointUrl returns a boolean if a field has been set.

### SetEndpointUrlNil

`func (o *OANSCZGF) SetEndpointUrlNil(b bool)`

 SetEndpointUrlNil sets the value for EndpointUrl to be an explicit nil

### UnsetEndpointUrl
`func (o *OANSCZGF) UnsetEndpointUrl()`

UnsetEndpointUrl ensures that no value is present for EndpointUrl, not even an explicit nil
### GetInstanceSubdomain

`func (o *OANSCZGF) GetInstanceSubdomain() string`

GetInstanceSubdomain returns the InstanceSubdomain field if non-nil, zero value otherwise.

### GetInstanceSubdomainOk

`func (o *OANSCZGF) GetInstanceSubdomainOk() (*string, bool)`

GetInstanceSubdomainOk returns a tuple with the InstanceSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSubdomain

`func (o *OANSCZGF) SetInstanceSubdomain(v string)`

SetInstanceSubdomain sets InstanceSubdomain field to given value.

### HasInstanceSubdomain

`func (o *OANSCZGF) HasInstanceSubdomain() bool`

HasInstanceSubdomain returns a boolean if a field has been set.

### GetClientId

`func (o *OANSCZGF) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OANSCZGF) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OANSCZGF) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OANSCZGF) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *OANSCZGF) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OANSCZGF) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OANSCZGF) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OANSCZGF) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetRedirectUri

`func (o *OANSCZGF) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OANSCZGF) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OANSCZGF) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *OANSCZGF) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetGongAccountEmail

`func (o *OANSCZGF) GetGongAccountEmail() string`

GetGongAccountEmail returns the GongAccountEmail field if non-nil, zero value otherwise.

### GetGongAccountEmailOk

`func (o *OANSCZGF) GetGongAccountEmailOk() (*string, bool)`

GetGongAccountEmailOk returns a tuple with the GongAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGongAccountEmail

`func (o *OANSCZGF) SetGongAccountEmail(v string)`

SetGongAccountEmail sets GongAccountEmail field to given value.

### HasGongAccountEmail

`func (o *OANSCZGF) HasGongAccountEmail() bool`

HasGongAccountEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


