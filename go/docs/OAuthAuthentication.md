# OAuthAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**AccessToken** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **NullableString** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 
**TenantName** | Pointer to **NullableString** |  | [optional] 
**SiteName** | Pointer to **NullableString** |  | [optional] 
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
**AccountName** | Pointer to **string** |  | [optional] 
**AccountKey** | Pointer to **string** |  | [optional] 
**InstanceSubdomain** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**RedirectUri** | Pointer to **string** |  | [optional] 
**GongAccountEmail** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuthAuthentication

`func NewOAuthAuthentication(source string, ) *OAuthAuthentication`

NewOAuthAuthentication instantiates a new OAuthAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthAuthenticationWithDefaults

`func NewOAuthAuthenticationWithDefaults() *OAuthAuthentication`

NewOAuthAuthenticationWithDefaults instantiates a new OAuthAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *OAuthAuthentication) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OAuthAuthentication) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OAuthAuthentication) SetSource(v string)`

SetSource sets Source field to given value.


### GetAccessToken

`func (o *OAuthAuthentication) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OAuthAuthentication) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OAuthAuthentication) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *OAuthAuthentication) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *OAuthAuthentication) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OAuthAuthentication) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OAuthAuthentication) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OAuthAuthentication) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *OAuthAuthentication) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *OAuthAuthentication) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetWorkspaceId

`func (o *OAuthAuthentication) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *OAuthAuthentication) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *OAuthAuthentication) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *OAuthAuthentication) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetTenantName

`func (o *OAuthAuthentication) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *OAuthAuthentication) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *OAuthAuthentication) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *OAuthAuthentication) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### SetTenantNameNil

`func (o *OAuthAuthentication) SetTenantNameNil(b bool)`

 SetTenantNameNil sets the value for TenantName to be an explicit nil

### UnsetTenantName
`func (o *OAuthAuthentication) UnsetTenantName()`

UnsetTenantName ensures that no value is present for TenantName, not even an explicit nil
### GetSiteName

`func (o *OAuthAuthentication) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *OAuthAuthentication) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *OAuthAuthentication) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *OAuthAuthentication) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteNameNil

`func (o *OAuthAuthentication) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *OAuthAuthentication) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetSubdomain

`func (o *OAuthAuthentication) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *OAuthAuthentication) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *OAuthAuthentication) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *OAuthAuthentication) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetAccessTokenSecret

`func (o *OAuthAuthentication) GetAccessTokenSecret() string`

GetAccessTokenSecret returns the AccessTokenSecret field if non-nil, zero value otherwise.

### GetAccessTokenSecretOk

`func (o *OAuthAuthentication) GetAccessTokenSecretOk() (*string, bool)`

GetAccessTokenSecretOk returns a tuple with the AccessTokenSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenSecret

`func (o *OAuthAuthentication) SetAccessTokenSecret(v string)`

SetAccessTokenSecret sets AccessTokenSecret field to given value.

### HasAccessTokenSecret

`func (o *OAuthAuthentication) HasAccessTokenSecret() bool`

HasAccessTokenSecret returns a boolean if a field has been set.

### GetUsername

`func (o *OAuthAuthentication) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OAuthAuthentication) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OAuthAuthentication) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OAuthAuthentication) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetZoteroId

`func (o *OAuthAuthentication) GetZoteroId() string`

GetZoteroId returns the ZoteroId field if non-nil, zero value otherwise.

### GetZoteroIdOk

`func (o *OAuthAuthentication) GetZoteroIdOk() (*string, bool)`

GetZoteroIdOk returns a tuple with the ZoteroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoteroId

`func (o *OAuthAuthentication) SetZoteroId(v string)`

SetZoteroId sets ZoteroId field to given value.

### HasZoteroId

`func (o *OAuthAuthentication) HasZoteroId() bool`

HasZoteroId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *OAuthAuthentication) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OAuthAuthentication) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OAuthAuthentication) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *OAuthAuthentication) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetDomain

`func (o *OAuthAuthentication) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OAuthAuthentication) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OAuthAuthentication) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *OAuthAuthentication) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetApiKey

`func (o *OAuthAuthentication) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OAuthAuthentication) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OAuthAuthentication) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *OAuthAuthentication) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAccessKey

`func (o *OAuthAuthentication) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *OAuthAuthentication) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *OAuthAuthentication) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *OAuthAuthentication) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessKeySecret

`func (o *OAuthAuthentication) GetAccessKeySecret() string`

GetAccessKeySecret returns the AccessKeySecret field if non-nil, zero value otherwise.

### GetAccessKeySecretOk

`func (o *OAuthAuthentication) GetAccessKeySecretOk() (*string, bool)`

GetAccessKeySecretOk returns a tuple with the AccessKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeySecret

`func (o *OAuthAuthentication) SetAccessKeySecret(v string)`

SetAccessKeySecret sets AccessKeySecret field to given value.

### HasAccessKeySecret

`func (o *OAuthAuthentication) HasAccessKeySecret() bool`

HasAccessKeySecret returns a boolean if a field has been set.

### GetEndpointUrl

`func (o *OAuthAuthentication) GetEndpointUrl() string`

GetEndpointUrl returns the EndpointUrl field if non-nil, zero value otherwise.

### GetEndpointUrlOk

`func (o *OAuthAuthentication) GetEndpointUrlOk() (*string, bool)`

GetEndpointUrlOk returns a tuple with the EndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUrl

`func (o *OAuthAuthentication) SetEndpointUrl(v string)`

SetEndpointUrl sets EndpointUrl field to given value.

### HasEndpointUrl

`func (o *OAuthAuthentication) HasEndpointUrl() bool`

HasEndpointUrl returns a boolean if a field has been set.

### SetEndpointUrlNil

`func (o *OAuthAuthentication) SetEndpointUrlNil(b bool)`

 SetEndpointUrlNil sets the value for EndpointUrl to be an explicit nil

### UnsetEndpointUrl
`func (o *OAuthAuthentication) UnsetEndpointUrl()`

UnsetEndpointUrl ensures that no value is present for EndpointUrl, not even an explicit nil
### GetAccountName

`func (o *OAuthAuthentication) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *OAuthAuthentication) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *OAuthAuthentication) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *OAuthAuthentication) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountKey

`func (o *OAuthAuthentication) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *OAuthAuthentication) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *OAuthAuthentication) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.

### HasAccountKey

`func (o *OAuthAuthentication) HasAccountKey() bool`

HasAccountKey returns a boolean if a field has been set.

### GetInstanceSubdomain

`func (o *OAuthAuthentication) GetInstanceSubdomain() string`

GetInstanceSubdomain returns the InstanceSubdomain field if non-nil, zero value otherwise.

### GetInstanceSubdomainOk

`func (o *OAuthAuthentication) GetInstanceSubdomainOk() (*string, bool)`

GetInstanceSubdomainOk returns a tuple with the InstanceSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSubdomain

`func (o *OAuthAuthentication) SetInstanceSubdomain(v string)`

SetInstanceSubdomain sets InstanceSubdomain field to given value.

### HasInstanceSubdomain

`func (o *OAuthAuthentication) HasInstanceSubdomain() bool`

HasInstanceSubdomain returns a boolean if a field has been set.

### GetClientId

`func (o *OAuthAuthentication) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthAuthentication) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthAuthentication) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuthAuthentication) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *OAuthAuthentication) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthAuthentication) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthAuthentication) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuthAuthentication) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetRedirectUri

`func (o *OAuthAuthentication) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OAuthAuthentication) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OAuthAuthentication) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *OAuthAuthentication) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetGongAccountEmail

`func (o *OAuthAuthentication) GetGongAccountEmail() string`

GetGongAccountEmail returns the GongAccountEmail field if non-nil, zero value otherwise.

### GetGongAccountEmailOk

`func (o *OAuthAuthentication) GetGongAccountEmailOk() (*string, bool)`

GetGongAccountEmailOk returns a tuple with the GongAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGongAccountEmail

`func (o *OAuthAuthentication) SetGongAccountEmail(v string)`

SetGongAccountEmail sets GongAccountEmail field to given value.

### HasGongAccountEmail

`func (o *OAuthAuthentication) HasGongAccountEmail() bool`

HasGongAccountEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


