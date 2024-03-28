# ListDataSourceItemsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | **int32** |  | 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Filters** | Pointer to [**NullableListItemsFiltersNullable**](ListItemsFiltersNullable.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListDataSourceItemsRequest

`func NewListDataSourceItemsRequest(dataSourceId int32, ) *ListDataSourceItemsRequest`

NewListDataSourceItemsRequest instantiates a new ListDataSourceItemsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDataSourceItemsRequestWithDefaults

`func NewListDataSourceItemsRequestWithDefaults() *ListDataSourceItemsRequest`

NewListDataSourceItemsRequestWithDefaults instantiates a new ListDataSourceItemsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *ListDataSourceItemsRequest) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *ListDataSourceItemsRequest) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *ListDataSourceItemsRequest) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.


### GetParentId

`func (o *ListDataSourceItemsRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ListDataSourceItemsRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ListDataSourceItemsRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ListDataSourceItemsRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *ListDataSourceItemsRequest) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *ListDataSourceItemsRequest) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetFilters

`func (o *ListDataSourceItemsRequest) GetFilters() ListItemsFiltersNullable`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListDataSourceItemsRequest) GetFiltersOk() (*ListItemsFiltersNullable, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListDataSourceItemsRequest) SetFilters(v ListItemsFiltersNullable)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListDataSourceItemsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *ListDataSourceItemsRequest) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *ListDataSourceItemsRequest) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetPagination

`func (o *ListDataSourceItemsRequest) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListDataSourceItemsRequest) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListDataSourceItemsRequest) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListDataSourceItemsRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


