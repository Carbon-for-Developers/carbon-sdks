# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** |  | 
**Id** | **string** |  | 
**Owner** | [**NullablePartialOwnerNullable**](PartialOwnerNullable.md) |  | 
**Name** | **NullableString** |  | 
**Industry** | **NullableString** |  | 
**Website** | **NullableString** |  | 
**NumberOfEmployees** | **NullableInt32** |  | 
**Addresses** | [**[]Address**](Address.md) |  | 
**PhoneNumbers** | [**[]PhoneNumber**](PhoneNumber.md) |  | 
**LastActivityAt** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**IsDeleted** | **bool** |  | 
**Tasks** | [**[]Task**](Task.md) |  | 
**Events** | [**[]Event**](Event.md) |  | 
**RemoteData** | **map[string]interface{}** |  | 

## Methods

### NewAccount

`func NewAccount(description NullableString, id string, owner NullablePartialOwnerNullable, name NullableString, industry NullableString, website NullableString, numberOfEmployees NullableInt32, addresses []Address, phoneNumbers []PhoneNumber, lastActivityAt NullableString, createdAt string, updatedAt string, isDeleted bool, tasks []Task, events []Event, remoteData map[string]interface{}, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Account) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Account) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Account) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Account) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Account) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.


### GetOwner

`func (o *Account) GetOwner() PartialOwnerNullable`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Account) GetOwnerOk() (*PartialOwnerNullable, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Account) SetOwner(v PartialOwnerNullable)`

SetOwner sets Owner field to given value.


### SetOwnerNil

`func (o *Account) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *Account) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Account) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Account) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIndustry

`func (o *Account) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *Account) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *Account) SetIndustry(v string)`

SetIndustry sets Industry field to given value.


### SetIndustryNil

`func (o *Account) SetIndustryNil(b bool)`

 SetIndustryNil sets the value for Industry to be an explicit nil

### UnsetIndustry
`func (o *Account) UnsetIndustry()`

UnsetIndustry ensures that no value is present for Industry, not even an explicit nil
### GetWebsite

`func (o *Account) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Account) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Account) SetWebsite(v string)`

SetWebsite sets Website field to given value.


### SetWebsiteNil

`func (o *Account) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *Account) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetNumberOfEmployees

`func (o *Account) GetNumberOfEmployees() int32`

GetNumberOfEmployees returns the NumberOfEmployees field if non-nil, zero value otherwise.

### GetNumberOfEmployeesOk

`func (o *Account) GetNumberOfEmployeesOk() (*int32, bool)`

GetNumberOfEmployeesOk returns a tuple with the NumberOfEmployees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfEmployees

`func (o *Account) SetNumberOfEmployees(v int32)`

SetNumberOfEmployees sets NumberOfEmployees field to given value.


### SetNumberOfEmployeesNil

`func (o *Account) SetNumberOfEmployeesNil(b bool)`

 SetNumberOfEmployeesNil sets the value for NumberOfEmployees to be an explicit nil

### UnsetNumberOfEmployees
`func (o *Account) UnsetNumberOfEmployees()`

UnsetNumberOfEmployees ensures that no value is present for NumberOfEmployees, not even an explicit nil
### GetAddresses

`func (o *Account) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Account) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Account) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.


### GetPhoneNumbers

`func (o *Account) GetPhoneNumbers() []PhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *Account) GetPhoneNumbersOk() (*[]PhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *Account) SetPhoneNumbers(v []PhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.


### GetLastActivityAt

`func (o *Account) GetLastActivityAt() string`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *Account) GetLastActivityAtOk() (*string, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *Account) SetLastActivityAt(v string)`

SetLastActivityAt sets LastActivityAt field to given value.


### SetLastActivityAtNil

`func (o *Account) SetLastActivityAtNil(b bool)`

 SetLastActivityAtNil sets the value for LastActivityAt to be an explicit nil

### UnsetLastActivityAt
`func (o *Account) UnsetLastActivityAt()`

UnsetLastActivityAt ensures that no value is present for LastActivityAt, not even an explicit nil
### GetCreatedAt

`func (o *Account) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Account) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Account) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Account) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Account) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Account) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *Account) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Account) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Account) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetTasks

`func (o *Account) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Account) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Account) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.


### SetTasksNil

`func (o *Account) SetTasksNil(b bool)`

 SetTasksNil sets the value for Tasks to be an explicit nil

### UnsetTasks
`func (o *Account) UnsetTasks()`

UnsetTasks ensures that no value is present for Tasks, not even an explicit nil
### GetEvents

`func (o *Account) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Account) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Account) SetEvents(v []Event)`

SetEvents sets Events field to given value.


### SetEventsNil

`func (o *Account) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *Account) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetRemoteData

`func (o *Account) GetRemoteData() map[string]interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Account) GetRemoteDataOk() (*map[string]interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Account) SetRemoteData(v map[string]interface{})`

SetRemoteData sets RemoteData field to given value.


### SetRemoteDataNil

`func (o *Account) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Account) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


