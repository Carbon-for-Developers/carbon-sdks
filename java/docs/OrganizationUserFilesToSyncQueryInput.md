

# OrganizationUserFilesToSyncQueryInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**pagination** | [**Pagination**](Pagination.md) |  |  [optional] |
|**orderBy** | **OrganizationUserFilesToSyncOrderByTypes** |  |  [optional] |
|**orderDir** | **OrderDir** |  |  [optional] |
|**filters** | [**OrganizationUserFilesToSyncFilters**](OrganizationUserFilesToSyncFilters.md) |  |  [optional] |
|**includeRawFile** | **Boolean** | If true, the query will return presigned URLs for the raw file. Only relevant for the /user_files_v2 endpoint. |  [optional] |
|**includeParsedTextFile** | **Boolean** | If true, the query will return presigned URLs for the parsed text file. Only relevant for the /user_files_v2 endpoint. |  [optional] |
|**includeAdditionalFiles** | **Boolean** | If true, the query will return presigned URLs for additional files. Only relevant for the /user_files_v2 endpoint. |  [optional] |
|**presignedUrlExpiryTimeSeconds** | **Integer** | The expiry time for the presigned URLs. Only relevant for the /user_files_v2 endpoint. |  [optional] |



