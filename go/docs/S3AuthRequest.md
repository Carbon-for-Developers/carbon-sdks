# S3AuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** |  | 
**AccessKeySecret** | **string** |  | 
**SyncSourceItems** | Pointer to **bool** | Enabling this flag will fetch all available content from the source to be listed via list items endpoint | [optional] [default to true]

## Methods

### NewS3AuthRequest

`func NewS3AuthRequest(accessKey string, accessKeySecret string, ) *S3AuthRequest`

NewS3AuthRequest instantiates a new S3AuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3AuthRequestWithDefaults

`func NewS3AuthRequestWithDefaults() *S3AuthRequest`

NewS3AuthRequestWithDefaults instantiates a new S3AuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *S3AuthRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *S3AuthRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *S3AuthRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetAccessKeySecret

`func (o *S3AuthRequest) GetAccessKeySecret() string`

GetAccessKeySecret returns the AccessKeySecret field if non-nil, zero value otherwise.

### GetAccessKeySecretOk

`func (o *S3AuthRequest) GetAccessKeySecretOk() (*string, bool)`

GetAccessKeySecretOk returns a tuple with the AccessKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeySecret

`func (o *S3AuthRequest) SetAccessKeySecret(v string)`

SetAccessKeySecret sets AccessKeySecret field to given value.


### GetSyncSourceItems

`func (o *S3AuthRequest) GetSyncSourceItems() bool`

GetSyncSourceItems returns the SyncSourceItems field if non-nil, zero value otherwise.

### GetSyncSourceItemsOk

`func (o *S3AuthRequest) GetSyncSourceItemsOk() (*bool, bool)`

GetSyncSourceItemsOk returns a tuple with the SyncSourceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSourceItems

`func (o *S3AuthRequest) SetSyncSourceItems(v bool)`

SetSyncSourceItems sets SyncSourceItems field to given value.

### HasSyncSourceItems

`func (o *S3AuthRequest) HasSyncSourceItems() bool`

HasSyncSourceItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


