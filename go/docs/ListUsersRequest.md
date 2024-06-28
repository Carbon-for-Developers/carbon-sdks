# ListUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Filters** | Pointer to [**ListUsersFilters**](ListUsersFilters.md) |  | [optional] 
**OrderBy** | Pointer to [**ListUsersOrderByTypes**](ListUsersOrderByTypes.md) |  | [optional] 
**OrderDir** | Pointer to [**OrderDirV2**](OrderDirV2.md) |  | [optional] 
**IncludeCount** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewListUsersRequest

`func NewListUsersRequest() *ListUsersRequest`

NewListUsersRequest instantiates a new ListUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsersRequestWithDefaults

`func NewListUsersRequestWithDefaults() *ListUsersRequest`

NewListUsersRequestWithDefaults instantiates a new ListUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ListUsersRequest) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListUsersRequest) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListUsersRequest) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListUsersRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetFilters

`func (o *ListUsersRequest) GetFilters() ListUsersFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListUsersRequest) GetFiltersOk() (*ListUsersFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListUsersRequest) SetFilters(v ListUsersFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListUsersRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetOrderBy

`func (o *ListUsersRequest) GetOrderBy() ListUsersOrderByTypes`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ListUsersRequest) GetOrderByOk() (*ListUsersOrderByTypes, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ListUsersRequest) SetOrderBy(v ListUsersOrderByTypes)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ListUsersRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetOrderDir

`func (o *ListUsersRequest) GetOrderDir() OrderDirV2`

GetOrderDir returns the OrderDir field if non-nil, zero value otherwise.

### GetOrderDirOk

`func (o *ListUsersRequest) GetOrderDirOk() (*OrderDirV2, bool)`

GetOrderDirOk returns a tuple with the OrderDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDir

`func (o *ListUsersRequest) SetOrderDir(v OrderDirV2)`

SetOrderDir sets OrderDir field to given value.

### HasOrderDir

`func (o *ListUsersRequest) HasOrderDir() bool`

HasOrderDir returns a boolean if a field has been set.

### GetIncludeCount

`func (o *ListUsersRequest) GetIncludeCount() bool`

GetIncludeCount returns the IncludeCount field if non-nil, zero value otherwise.

### GetIncludeCountOk

`func (o *ListUsersRequest) GetIncludeCountOk() (*bool, bool)`

GetIncludeCountOk returns a tuple with the IncludeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCount

`func (o *ListUsersRequest) SetIncludeCount(v bool)`

SetIncludeCount sets IncludeCount field to given value.

### HasIncludeCount

`func (o *ListUsersRequest) HasIncludeCount() bool`

HasIncludeCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


