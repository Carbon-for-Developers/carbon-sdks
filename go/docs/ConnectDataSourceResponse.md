# ConnectDataSourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSource** | [**OrganizationUserDataSourceAPI**](OrganizationUserDataSourceAPI.md) |  | 
**SyncUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConnectDataSourceResponse

`func NewConnectDataSourceResponse(dataSource OrganizationUserDataSourceAPI, ) *ConnectDataSourceResponse`

NewConnectDataSourceResponse instantiates a new ConnectDataSourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectDataSourceResponseWithDefaults

`func NewConnectDataSourceResponseWithDefaults() *ConnectDataSourceResponse`

NewConnectDataSourceResponseWithDefaults instantiates a new ConnectDataSourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSource

`func (o *ConnectDataSourceResponse) GetDataSource() OrganizationUserDataSourceAPI`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *ConnectDataSourceResponse) GetDataSourceOk() (*OrganizationUserDataSourceAPI, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *ConnectDataSourceResponse) SetDataSource(v OrganizationUserDataSourceAPI)`

SetDataSource sets DataSource field to given value.


### GetSyncUrl

`func (o *ConnectDataSourceResponse) GetSyncUrl() string`

GetSyncUrl returns the SyncUrl field if non-nil, zero value otherwise.

### GetSyncUrlOk

`func (o *ConnectDataSourceResponse) GetSyncUrlOk() (*string, bool)`

GetSyncUrlOk returns a tuple with the SyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUrl

`func (o *ConnectDataSourceResponse) SetSyncUrl(v string)`

SetSyncUrl sets SyncUrl field to given value.

### HasSyncUrl

`func (o *ConnectDataSourceResponse) HasSyncUrl() bool`

HasSyncUrl returns a boolean if a field has been set.

### SetSyncUrlNil

`func (o *ConnectDataSourceResponse) SetSyncUrlNil(b bool)`

 SetSyncUrlNil sets the value for SyncUrl to be an explicit nil

### UnsetSyncUrl
`func (o *ConnectDataSourceResponse) UnsetSyncUrl()`

UnsetSyncUrl ensures that no value is present for SyncUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


