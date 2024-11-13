# Lead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **NullableString** |  | 
**Description** | **NullableString** |  | 
**Id** | **string** |  | 
**Owner** | [**PartialOwner**](PartialOwner.md) |  | 
**Source** | **NullableString** |  | 
**Status** | **NullableString** |  | 
**Company** | **NullableString** |  | 
**FirstName** | **NullableString** |  | 
**LastName** | **NullableString** |  | 
**Addresses** | [**[]Address**](Address.md) |  | 
**PhoneNumbers** | [**[]PhoneNumber**](PhoneNumber.md) |  | 
**Emails** | [**[]Email**](Email.md) |  | 
**ConvertedAt** | **NullableString** |  | 
**ConvertedAccount** | [**NullablePartialAccountNullable**](PartialAccountNullable.md) |  | 
**ConvertedContact** | [**NullablePartialContactNullable**](PartialContactNullable.md) |  | 
**LastActivityAt** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**IsDeleted** | **bool** |  | 
**Tasks** | Pointer to [**[]Task**](Task.md) |  | [optional] 
**Events** | Pointer to [**[]Event**](Event.md) |  | [optional] 
**RemoteData** | **map[string]interface{}** |  | 

## Methods

### NewLead

`func NewLead(title NullableString, description NullableString, id string, owner PartialOwner, source NullableString, status NullableString, company NullableString, firstName NullableString, lastName NullableString, addresses []Address, phoneNumbers []PhoneNumber, emails []Email, convertedAt NullableString, convertedAccount NullablePartialAccountNullable, convertedContact NullablePartialContactNullable, lastActivityAt NullableString, createdAt string, updatedAt string, isDeleted bool, remoteData map[string]interface{}, ) *Lead`

NewLead instantiates a new Lead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeadWithDefaults

`func NewLeadWithDefaults() *Lead`

NewLeadWithDefaults instantiates a new Lead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Lead) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Lead) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Lead) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *Lead) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Lead) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *Lead) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Lead) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Lead) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Lead) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Lead) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *Lead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Lead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Lead) SetId(v string)`

SetId sets Id field to given value.


### GetOwner

`func (o *Lead) GetOwner() PartialOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Lead) GetOwnerOk() (*PartialOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Lead) SetOwner(v PartialOwner)`

SetOwner sets Owner field to given value.


### GetSource

`func (o *Lead) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Lead) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Lead) SetSource(v string)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *Lead) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *Lead) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetStatus

`func (o *Lead) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Lead) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Lead) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *Lead) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Lead) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCompany

`func (o *Lead) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Lead) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Lead) SetCompany(v string)`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *Lead) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *Lead) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetFirstName

`func (o *Lead) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Lead) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Lead) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *Lead) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *Lead) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *Lead) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Lead) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Lead) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *Lead) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *Lead) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetAddresses

`func (o *Lead) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Lead) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Lead) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.


### GetPhoneNumbers

`func (o *Lead) GetPhoneNumbers() []PhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *Lead) GetPhoneNumbersOk() (*[]PhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *Lead) SetPhoneNumbers(v []PhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.


### GetEmails

`func (o *Lead) GetEmails() []Email`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *Lead) GetEmailsOk() (*[]Email, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *Lead) SetEmails(v []Email)`

SetEmails sets Emails field to given value.


### GetConvertedAt

`func (o *Lead) GetConvertedAt() string`

GetConvertedAt returns the ConvertedAt field if non-nil, zero value otherwise.

### GetConvertedAtOk

`func (o *Lead) GetConvertedAtOk() (*string, bool)`

GetConvertedAtOk returns a tuple with the ConvertedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedAt

`func (o *Lead) SetConvertedAt(v string)`

SetConvertedAt sets ConvertedAt field to given value.


### SetConvertedAtNil

`func (o *Lead) SetConvertedAtNil(b bool)`

 SetConvertedAtNil sets the value for ConvertedAt to be an explicit nil

### UnsetConvertedAt
`func (o *Lead) UnsetConvertedAt()`

UnsetConvertedAt ensures that no value is present for ConvertedAt, not even an explicit nil
### GetConvertedAccount

`func (o *Lead) GetConvertedAccount() PartialAccountNullable`

GetConvertedAccount returns the ConvertedAccount field if non-nil, zero value otherwise.

### GetConvertedAccountOk

`func (o *Lead) GetConvertedAccountOk() (*PartialAccountNullable, bool)`

GetConvertedAccountOk returns a tuple with the ConvertedAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedAccount

`func (o *Lead) SetConvertedAccount(v PartialAccountNullable)`

SetConvertedAccount sets ConvertedAccount field to given value.


### SetConvertedAccountNil

`func (o *Lead) SetConvertedAccountNil(b bool)`

 SetConvertedAccountNil sets the value for ConvertedAccount to be an explicit nil

### UnsetConvertedAccount
`func (o *Lead) UnsetConvertedAccount()`

UnsetConvertedAccount ensures that no value is present for ConvertedAccount, not even an explicit nil
### GetConvertedContact

`func (o *Lead) GetConvertedContact() PartialContactNullable`

GetConvertedContact returns the ConvertedContact field if non-nil, zero value otherwise.

### GetConvertedContactOk

`func (o *Lead) GetConvertedContactOk() (*PartialContactNullable, bool)`

GetConvertedContactOk returns a tuple with the ConvertedContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedContact

`func (o *Lead) SetConvertedContact(v PartialContactNullable)`

SetConvertedContact sets ConvertedContact field to given value.


### SetConvertedContactNil

`func (o *Lead) SetConvertedContactNil(b bool)`

 SetConvertedContactNil sets the value for ConvertedContact to be an explicit nil

### UnsetConvertedContact
`func (o *Lead) UnsetConvertedContact()`

UnsetConvertedContact ensures that no value is present for ConvertedContact, not even an explicit nil
### GetLastActivityAt

`func (o *Lead) GetLastActivityAt() string`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *Lead) GetLastActivityAtOk() (*string, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *Lead) SetLastActivityAt(v string)`

SetLastActivityAt sets LastActivityAt field to given value.


### SetLastActivityAtNil

`func (o *Lead) SetLastActivityAtNil(b bool)`

 SetLastActivityAtNil sets the value for LastActivityAt to be an explicit nil

### UnsetLastActivityAt
`func (o *Lead) UnsetLastActivityAt()`

UnsetLastActivityAt ensures that no value is present for LastActivityAt, not even an explicit nil
### GetCreatedAt

`func (o *Lead) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Lead) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Lead) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Lead) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Lead) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Lead) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *Lead) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Lead) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Lead) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetTasks

`func (o *Lead) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Lead) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Lead) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *Lead) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasksNil

`func (o *Lead) SetTasksNil(b bool)`

 SetTasksNil sets the value for Tasks to be an explicit nil

### UnsetTasks
`func (o *Lead) UnsetTasks()`

UnsetTasks ensures that no value is present for Tasks, not even an explicit nil
### GetEvents

`func (o *Lead) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Lead) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Lead) SetEvents(v []Event)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Lead) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *Lead) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *Lead) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetRemoteData

`func (o *Lead) GetRemoteData() map[string]interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Lead) GetRemoteDataOk() (*map[string]interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Lead) SetRemoteData(v map[string]interface{})`

SetRemoteData sets RemoteData field to given value.


### SetRemoteDataNil

`func (o *Lead) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Lead) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


