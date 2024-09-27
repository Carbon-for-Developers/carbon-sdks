# Opportunity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** |  | 
**Id** | **string** |  | 
**Owner** | [**PartialOwner**](PartialOwner.md) |  | 
**Name** | **NullableString** |  | 
**Amount** | **NullableInt32** |  | 
**Account** | [**NullablePartialAccountNullable**](PartialAccountNullable.md) |  | 
**Contact** | [**NullablePartialContactNullable**](PartialContactNullable.md) |  | 
**Stage** | **NullableString** |  | 
**Status** | [**NullableOpportunityStatusNullable**](OpportunityStatusNullable.md) |  | 
**CloseDate** | **NullableString** |  | 
**LastActivityAt** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**IsDeleted** | **bool** |  | 
**Tasks** | [**[]Task**](Task.md) |  | 
**Events** | [**[]Event**](Event.md) |  | 
**RemoteData** | **map[string]interface{}** |  | 

## Methods

### NewOpportunity

`func NewOpportunity(description NullableString, id string, owner PartialOwner, name NullableString, amount NullableInt32, account NullablePartialAccountNullable, contact NullablePartialContactNullable, stage NullableString, status NullableOpportunityStatusNullable, closeDate NullableString, lastActivityAt NullableString, createdAt string, updatedAt string, isDeleted bool, tasks []Task, events []Event, remoteData map[string]interface{}, ) *Opportunity`

NewOpportunity instantiates a new Opportunity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunityWithDefaults

`func NewOpportunityWithDefaults() *Opportunity`

NewOpportunityWithDefaults instantiates a new Opportunity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Opportunity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Opportunity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Opportunity) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Opportunity) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Opportunity) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *Opportunity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Opportunity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Opportunity) SetId(v string)`

SetId sets Id field to given value.


### GetOwner

`func (o *Opportunity) GetOwner() PartialOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Opportunity) GetOwnerOk() (*PartialOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Opportunity) SetOwner(v PartialOwner)`

SetOwner sets Owner field to given value.


### GetName

`func (o *Opportunity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Opportunity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Opportunity) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Opportunity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Opportunity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAmount

`func (o *Opportunity) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Opportunity) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Opportunity) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### SetAmountNil

`func (o *Opportunity) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *Opportunity) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetAccount

`func (o *Opportunity) GetAccount() PartialAccountNullable`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Opportunity) GetAccountOk() (*PartialAccountNullable, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Opportunity) SetAccount(v PartialAccountNullable)`

SetAccount sets Account field to given value.


### SetAccountNil

`func (o *Opportunity) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *Opportunity) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetContact

`func (o *Opportunity) GetContact() PartialContactNullable`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Opportunity) GetContactOk() (*PartialContactNullable, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Opportunity) SetContact(v PartialContactNullable)`

SetContact sets Contact field to given value.


### SetContactNil

`func (o *Opportunity) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *Opportunity) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetStage

`func (o *Opportunity) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *Opportunity) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *Opportunity) SetStage(v string)`

SetStage sets Stage field to given value.


### SetStageNil

`func (o *Opportunity) SetStageNil(b bool)`

 SetStageNil sets the value for Stage to be an explicit nil

### UnsetStage
`func (o *Opportunity) UnsetStage()`

UnsetStage ensures that no value is present for Stage, not even an explicit nil
### GetStatus

`func (o *Opportunity) GetStatus() OpportunityStatusNullable`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Opportunity) GetStatusOk() (*OpportunityStatusNullable, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Opportunity) SetStatus(v OpportunityStatusNullable)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *Opportunity) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Opportunity) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCloseDate

`func (o *Opportunity) GetCloseDate() string`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *Opportunity) GetCloseDateOk() (*string, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *Opportunity) SetCloseDate(v string)`

SetCloseDate sets CloseDate field to given value.


### SetCloseDateNil

`func (o *Opportunity) SetCloseDateNil(b bool)`

 SetCloseDateNil sets the value for CloseDate to be an explicit nil

### UnsetCloseDate
`func (o *Opportunity) UnsetCloseDate()`

UnsetCloseDate ensures that no value is present for CloseDate, not even an explicit nil
### GetLastActivityAt

`func (o *Opportunity) GetLastActivityAt() string`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *Opportunity) GetLastActivityAtOk() (*string, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *Opportunity) SetLastActivityAt(v string)`

SetLastActivityAt sets LastActivityAt field to given value.


### SetLastActivityAtNil

`func (o *Opportunity) SetLastActivityAtNil(b bool)`

 SetLastActivityAtNil sets the value for LastActivityAt to be an explicit nil

### UnsetLastActivityAt
`func (o *Opportunity) UnsetLastActivityAt()`

UnsetLastActivityAt ensures that no value is present for LastActivityAt, not even an explicit nil
### GetCreatedAt

`func (o *Opportunity) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Opportunity) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Opportunity) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Opportunity) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Opportunity) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Opportunity) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeleted

`func (o *Opportunity) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Opportunity) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Opportunity) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetTasks

`func (o *Opportunity) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Opportunity) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Opportunity) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.


### SetTasksNil

`func (o *Opportunity) SetTasksNil(b bool)`

 SetTasksNil sets the value for Tasks to be an explicit nil

### UnsetTasks
`func (o *Opportunity) UnsetTasks()`

UnsetTasks ensures that no value is present for Tasks, not even an explicit nil
### GetEvents

`func (o *Opportunity) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Opportunity) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Opportunity) SetEvents(v []Event)`

SetEvents sets Events field to given value.


### SetEventsNil

`func (o *Opportunity) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *Opportunity) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetRemoteData

`func (o *Opportunity) GetRemoteData() map[string]interface{}`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Opportunity) GetRemoteDataOk() (*map[string]interface{}, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Opportunity) SetRemoteData(v map[string]interface{})`

SetRemoteData sets RemoteData field to given value.


### SetRemoteDataNil

`func (o *Opportunity) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Opportunity) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


