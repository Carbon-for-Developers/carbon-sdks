

# DeleteFilesV2QueryInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**filters** | [**OrganizationUserFilesToSyncFilters**](OrganizationUserFilesToSyncFilters.md) |  |  [optional] |
|**sendWebhook** | **Boolean** |  |  [optional] |
|**preserveFileRecord** | **Boolean** | Whether or not to delete all data related to the file from the database, BUT to preserve the file metadata, allowing for         resyncs. By default &#x60;preserve_file_record&#x60; is false, which means that all data related to the file *as well as* its metadata will be deleted. Note that         even if &#x60;preserve_file_record&#x60; is true, raw files uploaded via the &#x60;uploadfile&#x60; endpoint still cannot be resynced. |  [optional] |



