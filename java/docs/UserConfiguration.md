

# UserConfiguration


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**autoSyncEnabledSources** | **Object** |  |  [optional] |
|**maxFiles** | **Integer** | Custom file upload limit for the user over *all* user&#39;s files across all uploads.          If set, then the user will not be allowed to upload more files than this limit. If not set, or if set to -1,         then the user will have no limit. |  [optional] |
|**maxFilesPerUpload** | **Integer** | Custom file upload limit for the user across a single upload.         If set, then the user will not be allowed to upload more files than this limit in a single upload. If not set,         or if set to -1, then the user will have no limit. |  [optional] |
|**maxCharacters** | **Integer** | Custom character upload limit for the user over *all* user&#39;s files across all uploads.          If set, then the user will not be allowed to upload more characters than this limit. If not set, or if set to -1,         then the user will have no limit. |  [optional] |
|**maxCharactersPerFile** | **Integer** | A single file upload from the user can not exceed this character limit.         If set, then the file will not be synced if it exceeds this limit. If not set, or if set to -1, then the          user will have no limit. |  [optional] |
|**maxCharactersPerUpload** | **Integer** | Custom character upload limit for the user across a single upload.         If set, then the user won&#39;t be able to sync more than this many characters in one upload.          If not set, or if set to -1, then the user will have no limit. |  [optional] |



