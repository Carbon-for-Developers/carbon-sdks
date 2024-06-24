# UserWebPagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**UserWebPagesFilters**](UserWebPagesFilters.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**OrderBy** | Pointer to [**UserWebPageOrderByTypes**](UserWebPageOrderByTypes.md) |  | [optional] 
**OrderDir** | Pointer to [**OrderDirV2**](OrderDirV2.md) |  | [optional] 

## Methods

### NewUserWebPagesRequest

`func NewUserWebPagesRequest() *UserWebPagesRequest`

NewUserWebPagesRequest instantiates a new UserWebPagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWebPagesRequestWithDefaults

`func NewUserWebPagesRequestWithDefaults() *UserWebPagesRequest`

NewUserWebPagesRequestWithDefaults instantiates a new UserWebPagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *UserWebPagesRequest) GetFilters() UserWebPagesFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *UserWebPagesRequest) GetFiltersOk() (*UserWebPagesFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *UserWebPagesRequest) SetFilters(v UserWebPagesFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *UserWebPagesRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPagination

`func (o *UserWebPagesRequest) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *UserWebPagesRequest) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *UserWebPagesRequest) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *UserWebPagesRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetOrderBy

`func (o *UserWebPagesRequest) GetOrderBy() UserWebPageOrderByTypes`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *UserWebPagesRequest) GetOrderByOk() (*UserWebPageOrderByTypes, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *UserWebPagesRequest) SetOrderBy(v UserWebPageOrderByTypes)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *UserWebPagesRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetOrderDir

`func (o *UserWebPagesRequest) GetOrderDir() OrderDirV2`

GetOrderDir returns the OrderDir field if non-nil, zero value otherwise.

### GetOrderDirOk

`func (o *UserWebPagesRequest) GetOrderDirOk() (*OrderDirV2, bool)`

GetOrderDirOk returns a tuple with the OrderDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDir

`func (o *UserWebPagesRequest) SetOrderDir(v OrderDirV2)`

SetOrderDir sets OrderDir field to given value.

### HasOrderDir

`func (o *UserWebPagesRequest) HasOrderDir() bool`

HasOrderDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


