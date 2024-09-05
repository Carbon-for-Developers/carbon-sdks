<div align="center">

[![Visit Carbon](https://raw.githubusercontent.com/Carbon-for-Developers/carbon-sdks/HEAD/java/header.png)](https://carbon.ai)

# [Carbon](https://carbon.ai)

Connect external data to LLMs, no matter the source.

[![Maven Central](https://img.shields.io/badge/Maven%20Central-v0.1.30-blue)](https://central.sonatype.com/artifact/com.konfigthis.carbonai/carbonai-java-sdk/0.1.30)

</div>

## Table of Contents

<!-- toc -->

## Requirements

Building the API client library requires:

1. Java 1.8+
2. Maven (3.8.3+)/Gradle (7.2+)

If you are adding this library to an Android Application or Library:

3. Android 8.0+ (API Level 26+)

## Installation

To install the API client library to your local Maven repository, simply execute:

```shell
mvn clean install
```

To deploy it to a remote Maven repository instead, configure the settings of the repository and execute:

```shell
mvn clean deploy
```

Refer to the [OSSRH Guide](http://central.sonatype.org/pages/ossrh-guide.html) for more information.

### Maven users

Add this dependency to your project's POM:

```xml
<dependency>
  <groupId>com.konfigthis.carbonai</groupId>
  <artifactId>carbonai-java-sdk</artifactId>
  <version>0.1.30</version>
  <scope>compile</scope>
</dependency>
```

### Gradle users

Add this dependency to your `build.gradle`:

```groovy
// build.gradle
repositories {
  mavenCentral()
}

dependencies {
   implementation "com.konfigthis.carbonai:carbonai-java-sdk:0.1.30"
}
```

### Android users

Make sure your `build.gradle` file as a `minSdk` version of at least 26:
```groovy
// build.gradle
android {
    defaultConfig {
        minSdk 26
    }
}
```

Also make sure your library or application has internet permissions in your `AndroidManifest.xml`:

```xml
<!--AndroidManifest.xml-->
<?xml version="1.0" encoding="utf-8"?>
<manifest xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:tools="http://schemas.android.com/tools">
    <uses-permission android:name="android.permission.INTERNET"/>
</manifest>
```

### Others

At first generate the JAR by executing:

```shell
mvn clean package
```

Then manually install the following JARs:

* `target/carbonai-java-sdk-0.1.30.jar`
* `target/lib/*.jar`

## Getting Started

Please follow the [installation](#installation) instruction and execute the following Java code:

```java
import com.konfigthis.carbonai.*;

public class GettingStarted {
    public static void main(String[] args) {
        // 1) Get an access token for a customer
        Configuration configuration = new Configuration();
        configuration.apiKey = "YOUR API KEY";
        configuration.customerId = "YOUR CUSTOMER ID";
        Carbon carbon = new Carbon(configuration);

        TokenResponse token = null;
        try {
            token = carbon.auth.getAccessToken().execute();
        } catch (ApiException e) {
            e.printStackTrace();
        }

        // 2) Use the access token to authenticate moving forward
        Configuration configuration2 = new Configuration();
        configuration2.accessToken = token.getAccessToken();
        Carbon carbon2 = new Carbon(configuration2);

        // use SDK as usual
        WhiteLabelingResponse whiteLabeling = null;
        try {
            whiteLabeling = carbon2.auth.getWhiteLabeling().execute();
        } catch (ApiException e) {
            e.printStackTrace();
        }
        // etc.
    }
}
```
## Reference


### `carbon.auth.getAccessToken`

Get Access Token

#### 🛠️ Usage

```java
TokenResponse result = client
        .auth
        .getAccessToken()
        .execute();
```

#### 🔄 Return

[TokenResponse](./src/main/java/com/konfigthis/client/model/TokenResponse.java)

#### 🌐 Endpoint

`/auth/v1/access_token` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.auth.getWhiteLabeling`

Returns whether or not the organization is white labeled and which integrations are white labeled

:param current_user: the current user
:param db: the database session
:return: a WhiteLabelingResponse

#### 🛠️ Usage

```java
WhiteLabelingResponse result = client
        .auth
        .getWhiteLabeling()
        .execute();
```

#### 🔄 Return

[WhiteLabelingResponse](./src/main/java/com/konfigthis/client/model/WhiteLabelingResponse.java)

#### 🌐 Endpoint

`/auth/v1/white_labeling` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.dataSources.queryUserDataSources`

User Data Sources

#### 🛠️ Usage

```java
OrganizationUserDataSourceResponse result = client
        .dataSources
        .queryUserDataSources()
        .pagination(pagination)
        .orderBy(orderBy)
        .orderDir(orderDir)
        .filters(filters)
        .execute();
```

#### ⚙️ Parameters

##### pagination: [`Pagination`](./src/main/java/com/konfigthis/client/model/Pagination.java)

##### order_by:

##### order_dir:

##### filters: [`OrganizationUserDataSourceFilters`](./src/main/java/com/konfigthis/client/model/OrganizationUserDataSourceFilters.java)

#### 🔄 Return

[OrganizationUserDataSourceResponse](./src/main/java/com/konfigthis/client/model/OrganizationUserDataSourceResponse.java)

#### 🌐 Endpoint

`/user_data_sources` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.dataSources.revokeAccessToken`

Revoke Access Token

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .dataSources
        .revokeAccessToken(dataSourceId)
        .execute();
```

#### ⚙️ Parameters

##### data_source_id: `Integer`

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/revoke_access_token` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.embeddings.getDocuments`

For pre-filtering documents, using `tags_v2` is preferred to using `tags` (which is now deprecated). If both `tags_v2`
and `tags` are specified, `tags` is ignored. `tags_v2` enables
building complex filters through the use of "AND", "OR", and negation logic. Take the below input as an example:
```json
{
    "OR": [
        {
            "key": "subject",
            "value": "holy-bible",
            "negate": false
        },
        {
            "key": "person-of-interest",
            "value": "jesus christ",
            "negate": false
        },
        {
            "key": "genre",
            "value": "religion",
            "negate": true
        }
        {
            "AND": [
                {
                    "key": "subject",
                    "value": "tao-te-ching",
                    "negate": false
                },
                {
                    "key": "author",
                    "value": "lao-tzu",
                    "negate": false
                }
            ]
        }
    ]
}
```
In this case, files will be filtered such that:
1. "subject" = "holy-bible" OR
2. "person-of-interest" = "jesus christ" OR
3. "genre" != "religion" OR
4. "subject" = "tao-te-ching" AND "author" = "lao-tzu"

Note that the top level of the query must be either an "OR" or "AND" array. Currently, nesting is limited to 3.
For tag blocks (those with "key", "value", and "negate" keys), the following typing rules apply:
1. "key" isn't optional and must be a `string`
2. "value" isn't optional and can be `any` or list[`any`]
3. "negate" is optional and must be `true` or `false`. If present and `true`, then the filter block is negated in
the resulting query. It is `false` by default.


When querying embeddings, you can optionally specify the `media_type` parameter in your request. By default (if
not set), it is equal to "TEXT". This means that the query will be performed over files that have
been parsed as text (for now, this covers all files except image files). If it is equal to "IMAGE",
the query will be performed over image files (for now, `.jpg` and `.png` files). You can think of this
field as an additional filter on top of any filters set in `file_ids` and


When `hybrid_search` is set to true, a combination of keyword search and semantic search are used to rank
and select candidate embeddings during information retrieval. By default, these search methods are weighted
equally during the ranking process. To adjust the weight (or "importance") of each search method, you can use
the `hybrid_search_tuning_parameters` property. The description for the different tuning parameters are:
- `weight_a`: weight to assign to semantic search
- `weight_b`: weight to assign to keyword search

You must ensure that `sum(weight_a, weight_b,..., weight_n)` for all *n* weights is equal to 1. The equality
has an error tolerance of 0.001 to account for possible floating point issues.

In order to use hybrid search for a customer across a set of documents, two flags need to be enabled:
1. Use the `/modify_user_configuration` endpoint to to enable `sparse_vectors` for the customer. The payload
body for this request is below:
```
{
  "configuration_key_name": "sparse_vectors",
  "value": {
    "enabled": true
  }
}
```
2. Make sure hybrid search is enabled for the documents across which you want to perform the search. For the
`/uploadfile` endpoint, this can be done by setting the following query parameter: `generate_sparse_vectors=true`


Carbon supports multiple models for use in generating embeddings for files. For images, we support Vertex AI's
multimodal model; for text, we support OpenAI's `text-embedding-ada-002` and Cohere's embed-multilingual-v3.0.
The model can be specified via the `embedding_model` parameter (in the POST body for `/embeddings`, and a query 
parameter in `/uploadfile`). If no model is supplied, the `text-embedding-ada-002` is used by default. When performing
embedding queries, embeddings from files that used the specified model will be considered in the query.
For example, if files A and B have embeddings generated with `OPENAI`, and files C and D have embeddings generated with
`COHERE_MULTILINGUAL_V3`, then by default, queries will only consider files A and B. If `COHERE_MULTILINGUAL_V3` is
specified as the `embedding_model` in `/embeddings`, then only files C and D will be considered. Make sure that
the set of all files you want considered for a query have embeddings generated via the same model. For now, **do not**
set `VERTEX_MULTIMODAL` as an `embedding_model`. This model is used automatically by Carbon when it detects an image file.

#### 🛠️ Usage

```java
DocumentResponseList result = client
        .embeddings
        .getDocuments(query, k)
        .tags(tags)
        .queryVector(queryVector)
        .fileIds(fileIds)
        .parentFileIds(parentFileIds)
        .includeAllChildren(includeAllChildren)
        .tagsV2(tagsV2)
        .includeTags(includeTags)
        .includeVectors(includeVectors)
        .includeRawFile(includeRawFile)
        .hybridSearch(hybridSearch)
        .hybridSearchTuningParameters(hybridSearchTuningParameters)
        .mediaType(mediaType)
        .embeddingModel(embeddingModel)
        .includeFileLevelMetadata(includeFileLevelMetadata)
        .highAccuracy(highAccuracy)
        .rerank(rerank)
        .fileTypesAtSource(fileTypesAtSource)
        .excludeColdStorageFiles(excludeColdStorageFiles)
        .execute();
```

#### ⚙️ Parameters

##### query: `String`

Query for which to get related chunks and embeddings.

##### k: `Integer`

Number of related chunks to return.

##### tags: `Map<String, Object>`

A set of tags to limit the search to. Deprecated and may be removed in the future.

##### query_vector: List<`Double`>

Optional query vector for which to get related chunks and embeddings. It must have been         generated by the same model used to generate the embeddings across which the search is being conducted. Cannot         provide both `query` and `query_vector`.

##### file_ids: List<`Integer`>

Optional list of file IDs to limit the search to

##### parent_file_ids: List<`Integer`>

Optional list of parent file IDs to limit the search to. A parent file describes a file to which         another file belongs (e.g. a folder)

##### include_all_children: `Boolean`

Flag to control whether or not to include all children of filtered files in the embedding search.

##### tags_v2: `Object`

A set of tags to limit the search to. Use this instead of `tags`, which is deprecated.

##### include_tags: `Boolean`

Flag to control whether or not to include tags for each chunk in the response.

##### include_vectors: `Boolean`

Flag to control whether or not to include embedding vectors in the response.

##### include_raw_file: `Boolean`

Flag to control whether or not to include a signed URL to the raw file containing each chunk         in the response.

##### hybrid_search: `Boolean`

Flag to control whether or not to perform hybrid search.

##### hybrid_search_tuning_parameters: [`HybridSearchTuningParamsNullable`](./src/main/java/com/konfigthis/client/model/HybridSearchTuningParamsNullable.java)

##### media_type:

##### embedding_model:

##### include_file_level_metadata: `Boolean`

Flag to control whether or not to include file-level metadata in the response. This metadata         will be included in the `content_metadata` field of each document along with chunk/embedding level metadata.

##### high_accuracy: `Boolean`

Flag to control whether or not to perform a high accuracy embedding search. By default, this is set to false.         If true, the search may return more accurate results, but may take longer to complete.

##### rerank: [`RerankParamsNullable`](./src/main/java/com/konfigthis/client/model/RerankParamsNullable.java)

##### file_types_at_source: List>

Filter files based on their type at the source (for example help center tickets and articles)

##### exclude_cold_storage_files: `Boolean`

Flag to control whether or not to exclude files that are not in hot storage. If set to False, then an error will be returned if any filtered         files are in cold storage.

#### 🔄 Return

[DocumentResponseList](./src/main/java/com/konfigthis/client/model/DocumentResponseList.java)

#### 🌐 Endpoint

`/embeddings` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.embeddings.getEmbeddingsAndChunks`

Retrieve Embeddings And Content

#### 🛠️ Usage

```java
EmbeddingsAndChunksResponse result = client
        .embeddings
        .getEmbeddingsAndChunks(filters)
        .pagination(pagination)
        .orderBy(orderBy)
        .orderDir(orderDir)
        .includeVectors(includeVectors)
        .execute();
```

#### ⚙️ Parameters

##### filters: [`EmbeddingsAndChunksFilters`](./src/main/java/com/konfigthis/client/model/EmbeddingsAndChunksFilters.java)

##### pagination: [`Pagination`](./src/main/java/com/konfigthis/client/model/Pagination.java)

##### order_by:

##### order_dir:

##### include_vectors: `Boolean`

#### 🔄 Return

[EmbeddingsAndChunksResponse](./src/main/java/com/konfigthis/client/model/EmbeddingsAndChunksResponse.java)

#### 🌐 Endpoint

`/text_chunks` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.embeddings.list`

Retrieve Embeddings And Content V2

#### 🛠️ Usage

```java
EmbeddingsAndChunksResponse result = client
        .embeddings
        .list(filters)
        .pagination(pagination)
        .orderBy(orderBy)
        .orderDir(orderDir)
        .includeVectors(includeVectors)
        .execute();
```

#### ⚙️ Parameters

##### filters: [`OrganizationUserFilesToSyncFilters`](./src/main/java/com/konfigthis/client/model/OrganizationUserFilesToSyncFilters.java)

##### pagination: [`Pagination`](./src/main/java/com/konfigthis/client/model/Pagination.java)

##### order_by:

##### order_dir:

##### include_vectors: `Boolean`

#### 🔄 Return

[EmbeddingsAndChunksResponse](./src/main/java/com/konfigthis/client/model/EmbeddingsAndChunksResponse.java)

#### 🌐 Endpoint

`/list_chunks_and_embeddings` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.embeddings.uploadChunksAndEmbeddings`

Upload Chunks And Embeddings

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .embeddings
        .uploadChunksAndEmbeddings(embeddingModel, chunksAndEmbeddings)
        .overwriteExisting(overwriteExisting)
        .chunksOnly(chunksOnly)
        .customCredentials(customCredentials)
        .execute();
```

#### ⚙️ Parameters

##### embedding_model:

##### chunks_and_embeddings: List<[`SingleChunksAndEmbeddingsUploadInput`](./src/main/java/com/konfigthis/client/model/SingleChunksAndEmbeddingsUploadInput.java)>

##### overwrite_existing: `Boolean`

##### chunks_only: `Boolean`

##### custom_credentials: `Map<String, Object>`

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/upload_chunks_and_embeddings` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.createUserFileTags`

A tag is a key-value pair that can be added to a file. This pair can then be used
for searches (e.g. embedding searches) in order to narrow down the scope of the search.
A file can have any number of tags. The following are reserved keys that cannot be used:
- db_embedding_id
- organization_id
- user_id
- organization_user_file_id

Carbon currently supports two data types for tag values - `string` and `list<string>`.
Keys can only be `string`. If values other than `string` and `list<string>` are used,
they're automatically converted to strings (e.g. 4 will become "4").

#### 🛠️ Usage

```java
UserFile result = client
        .files
        .createUserFileTags(tags, organizationUserFileId)
        .execute();
```

#### ⚙️ Parameters

##### tags: `Map<String, Object>`

##### organization_user_file_id: `Integer`

#### 🔄 Return

[UserFile](./src/main/java/com/konfigthis/client/model/UserFile.java)

#### 🌐 Endpoint

`/create_user_file_tags` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.delete`
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

Delete File Endpoint

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .files
        .delete(fileId)
        .execute();
```

#### ⚙️ Parameters

##### fileId: `Integer`

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/deletefile/{file_id}` `DELETE`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.deleteFileTags`

Delete File Tags

#### 🛠️ Usage

```java
UserFile result = client
        .files
        .deleteFileTags(tags, organizationUserFileId)
        .execute();
```

#### ⚙️ Parameters

##### tags: List<`String`>

##### organization_user_file_id: `Integer`

#### 🔄 Return

[UserFile](./src/main/java/com/konfigthis/client/model/UserFile.java)

#### 🌐 Endpoint

`/delete_user_file_tags` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.deleteMany`
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

Delete Files Endpoint

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .files
        .deleteMany()
        .fileIds(fileIds)
        .syncStatuses(syncStatuses)
        .deleteNonSyncedOnly(deleteNonSyncedOnly)
        .sendWebhook(sendWebhook)
        .deleteChildFiles(deleteChildFiles)
        .execute();
```

#### ⚙️ Parameters

##### file_ids: List<`Integer`>

##### sync_statuses: List>

##### delete_non_synced_only: `Boolean`

##### send_webhook: `Boolean`

##### delete_child_files: `Boolean`

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/delete_files` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.deleteV2`

Delete Files V2 Endpoint

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .files
        .deleteV2()
        .filters(filters)
        .sendWebhook(sendWebhook)
        .preserveFileRecord(preserveFileRecord)
        .execute();
```

#### ⚙️ Parameters

##### filters: [`OrganizationUserFilesToSyncFilters`](./src/main/java/com/konfigthis/client/model/OrganizationUserFilesToSyncFilters.java)

##### send_webhook: `Boolean`

##### preserve_file_record: `Boolean`

Whether or not to delete all data related to the file from the database, BUT to preserve the file metadata, allowing for         resyncs. By default `preserve_file_record` is false, which means that all data related to the file *as well as* its metadata will be deleted. Note that         even if `preserve_file_record` is true, raw files uploaded via the `uploadfile` endpoint still cannot be resynced.

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/delete_files_v2` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.getParsedFile`
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

This route is deprecated. Use `/user_files_v2` instead.

#### 🛠️ Usage

```java
PresignedURLResponse result = client
        .files
        .getParsedFile(fileId)
        .execute();
```

#### ⚙️ Parameters

##### fileId: `Integer`

#### 🔄 Return

[PresignedURLResponse](./src/main/java/com/konfigthis/client/model/PresignedURLResponse.java)

#### 🌐 Endpoint

`/parsed_file/{file_id}` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.getRawFile`
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

This route is deprecated. Use `/user_files_v2` instead.

#### 🛠️ Usage

```java
PresignedURLResponse result = client
        .files
        .getRawFile(fileId)
        .execute();
```

#### ⚙️ Parameters

##### fileId: `Integer`

#### 🔄 Return

[PresignedURLResponse](./src/main/java/com/konfigthis/client/model/PresignedURLResponse.java)

#### 🌐 Endpoint

`/raw_file/{file_id}` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.modifyColdStorageParameters`

Modify Cold Storage Parameters

#### 🛠️ Usage

```java
Boolean result = client
        .files
        .modifyColdStorageParameters()
        .filters(filters)
        .enableColdStorage(enableColdStorage)
        .hotStorageTimeToLive(hotStorageTimeToLive)
        .execute();
```

#### ⚙️ Parameters

##### filters: [`OrganizationUserFilesToSyncFilters`](./src/main/java/com/konfigthis/client/model/OrganizationUserFilesToSyncFilters.java)

##### enable_cold_storage: `Boolean`

##### hot_storage_time_to_live: `Integer`

#### 🌐 Endpoint

`/modify_cold_storage_parameters` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.moveToHotStorage`

Move To Hot Storage

#### 🛠️ Usage

```java
Boolean result = client
        .files
        .moveToHotStorage()
        .filters(filters)
        .execute();
```

#### ⚙️ Parameters

##### filters: [`OrganizationUserFilesToSyncFilters`](./src/main/java/com/konfigthis/client/model/OrganizationUserFilesToSyncFilters.java)

#### 🌐 Endpoint

`/move_to_hot_storage` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.queryUserFiles`

For pre-filtering documents, using `tags_v2` is preferred to using `tags` (which is now deprecated). If both `tags_v2`
and `tags` are specified, `tags` is ignored. `tags_v2` enables
building complex filters through the use of "AND", "OR", and negation logic. Take the below input as an example:
```json
{
    "OR": [
        {
            "key": "subject",
            "value": "holy-bible",
            "negate": false
        },
        {
            "key": "person-of-interest",
            "value": "jesus christ",
            "negate": false
        },
        {
            "key": "genre",
            "value": "religion",
            "negate": true
        }
        {
            "AND": [
                {
                    "key": "subject",
                    "value": "tao-te-ching",
                    "negate": false
                },
                {
                    "key": "author",
                    "value": "lao-tzu",
                    "negate": false
                }
            ]
        }
    ]
}
```
In this case, files will be filtered such that:
1. "subject" = "holy-bible" OR
2. "person-of-interest" = "jesus christ" OR
3. "genre" != "religion" OR
4. "subject" = "tao-te-ching" AND "author" = "lao-tzu"

Note that the top level of the query must be either an "OR" or "AND" array. Currently, nesting is limited to 3.
For tag blocks (those with "key", "value", and "negate" keys), the following typing rules apply:
1. "key" isn't optional and must be a `string`
2. "value" isn't optional and can be `any` or list[`any`]
3. "negate" is optional and must be `true` or `false`. If present and `true`, then the filter block is negated in
the resulting query. It is `false` by default.

#### 🛠️ Usage

```java
UserFilesV2 result = client
        .files
        .queryUserFiles()
        .pagination(pagination)
        .orderBy(orderBy)
        .orderDir(orderDir)
        .filters(filters)
        .includeRawFile(includeRawFile)
        .includeParsedTextFile(includeParsedTextFile)
        .includeAdditionalFiles(includeAdditionalFiles)
        .execute();
```

#### ⚙️ Parameters

##### pagination: [`Pagination`](./src/main/java/com/konfigthis/client/model/Pagination.java)

##### order_by:

##### order_dir:

##### filters: [`OrganizationUserFilesToSyncFilters`](./src/main/java/com/konfigthis/client/model/OrganizationUserFilesToSyncFilters.java)

##### include_raw_file: `Boolean`

##### include_parsed_text_file: `Boolean`

##### include_additional_files: `Boolean`

#### 🔄 Return

[UserFilesV2](./src/main/java/com/konfigthis/client/model/UserFilesV2.java)

#### 🌐 Endpoint

`/user_files_v2` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.queryUserFilesDeprecated`
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

This route is deprecated. Use `/user_files_v2` instead.

#### 🛠️ Usage

```java
List<UserFile> result = client
        .files
        .queryUserFilesDeprecated()
        .pagination(pagination)
        .orderBy(orderBy)
        .orderDir(orderDir)
        .filters(filters)
        .includeRawFile(includeRawFile)
        .includeParsedTextFile(includeParsedTextFile)
        .includeAdditionalFiles(includeAdditionalFiles)
        .execute();
```

#### ⚙️ Parameters

##### pagination: [`Pagination`](./src/main/java/com/konfigthis/client/model/Pagination.java)

##### order_by:

##### order_dir:

##### filters: [`OrganizationUserFilesToSyncFilters`](./src/main/java/com/konfigthis/client/model/OrganizationUserFilesToSyncFilters.java)

##### include_raw_file: `Boolean`

##### include_parsed_text_file: `Boolean`

##### include_additional_files: `Boolean`

#### 🔄 Return

[UserFile](./src/main/java/com/konfigthis/client/model/UserFile.java)

#### 🌐 Endpoint

`/user_files` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.resync`

Resync File

#### 🛠️ Usage

```java
UserFile result = client
        .files
        .resync(fileId)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .forceEmbeddingGeneration(forceEmbeddingGeneration)
        .execute();
```

#### ⚙️ Parameters

##### file_id: `Integer`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### force_embedding_generation: `Boolean`

#### 🔄 Return

[UserFile](./src/main/java/com/konfigthis/client/model/UserFile.java)

#### 🌐 Endpoint

`/resync_file` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.upload`

This endpoint is used to directly upload local files to Carbon. The `POST` request should be a multipart form request.
Note that the `set_page_as_boundary` query parameter is applicable only to PDFs for now. When this value is set,
PDF chunks are at most one page long. Additional information can be retrieved for each chunk, however, namely the coordinates
of the bounding box around the chunk (this can be used for things like text highlighting). Following is a description
of all possible query parameters:
- `chunk_size`: the chunk size (in tokens) applied when splitting the document
- `chunk_overlap`: the chunk overlap (in tokens) applied when splitting the document
- `skip_embedding_generation`: whether or not to skip the generation of chunks and embeddings
- `set_page_as_boundary`: described above
- `embedding_model`: the model used to generate embeddings for the document chunks
- `use_ocr`: whether or not to use OCR as a preprocessing step prior to generating chunks. Valid for PDFs, JPEGs, and PNGs
- `generate_sparse_vectors`: whether or not to generate sparse vectors for the file. Required for hybrid search.
- `prepend_filename_to_chunks`: whether or not to prepend the filename to the chunk text


Carbon supports multiple models for use in generating embeddings for files. For images, we support Vertex AI's
multimodal model; for text, we support OpenAI's `text-embedding-ada-002` and Cohere's embed-multilingual-v3.0.
The model can be specified via the `embedding_model` parameter (in the POST body for `/embeddings`, and a query 
parameter in `/uploadfile`). If no model is supplied, the `text-embedding-ada-002` is used by default. When performing
embedding queries, embeddings from files that used the specified model will be considered in the query.
For example, if files A and B have embeddings generated with `OPENAI`, and files C and D have embeddings generated with
`COHERE_MULTILINGUAL_V3`, then by default, queries will only consider files A and B. If `COHERE_MULTILINGUAL_V3` is
specified as the `embedding_model` in `/embeddings`, then only files C and D will be considered. Make sure that
the set of all files you want considered for a query have embeddings generated via the same model. For now, **do not**
set `VERTEX_MULTIMODAL` as an `embedding_model`. This model is used automatically by Carbon when it detects an image file.

#### 🛠️ Usage

```java
UserFile result = client
        .files
        .upload(_file)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .setPageAsBoundary(setPageAsBoundary)
        .embeddingModel(embeddingModel)
        .useOcr(useOcr)
        .generateSparseVectors(generateSparseVectors)
        .prependFilenameToChunks(prependFilenameToChunks)
        .maxItemsPerChunk(maxItemsPerChunk)
        .parsePdfTablesWithOcr(parsePdfTablesWithOcr)
        .detectAudioLanguage(detectAudioLanguage)
        .transcriptionService(transcriptionService)
        .includeSpeakerLabels(includeSpeakerLabels)
        .mediaType(mediaType)
        .splitRows(splitRows)
        .enableColdStorage(enableColdStorage)
        .hotStorageTimeToLive(hotStorageTimeToLive)
        .generateChunksOnly(generateChunksOnly)
        .execute();
```

#### ⚙️ Parameters

##### _file: `File`

##### chunkSize: `Integer`

Chunk size in tiktoken tokens to be used when processing file.

##### chunkOverlap: `Integer`

Chunk overlap in tiktoken tokens to be used when processing file.

##### skipEmbeddingGeneration: `Boolean`

Flag to control whether or not embeddings should be generated and stored             when processing file.

##### setPageAsBoundary: `Boolean`

Flag to control whether or not to set the a page's worth of content as the maximum             amount of content that can appear in a chunk. Only valid for PDFs. See description route description for             more information.

##### embeddingModel:

Embedding model that will be used to embed file chunks.

##### useOcr: `Boolean`

Whether or not to use OCR when processing files. Valid for PDFs, JPEGs, and PNGs. Useful for documents with             tables, images, and/or scanned text.

##### generateSparseVectors: `Boolean`

Whether or not to generate sparse vectors for the file. This is *required* for the file to be a             candidate for hybrid search.

##### prependFilenameToChunks: `Boolean`

Whether or not to prepend the file's name to chunks.

##### maxItemsPerChunk: `Integer`

Number of objects per chunk. For csv, tsv, xlsx, and json files only.

##### parsePdfTablesWithOcr: `Boolean`

Whether to use rich table parsing when `use_ocr` is enabled.

##### detectAudioLanguage: `Boolean`

Whether to automatically detect the language of the uploaded audio file.

##### transcriptionService:

The transcription service to use for audio files. If no service is specified, 'deepgram' will be used.

##### includeSpeakerLabels: `Boolean`

Detect multiple speakers and label segments of speech by speaker for audio files.

##### mediaType:

The media type of the file. If not provided, it will be inferred from the file extension.

##### splitRows: `Boolean`

Whether to split tabular rows into chunks. Currently only valid for CSV, TSV, and XLSX files.

##### enableColdStorage: `Boolean`

Enable cold storage for the file. If set to true, the file will be moved to cold storage after a certain period of inactivity. Default is false.

##### hotStorageTimeToLive: `Integer`

Time in seconds after which the file will be moved to cold storage.

##### generateChunksOnly: `Boolean`

If this flag is enabled, the file will be chunked and stored with Carbon,             but no embeddings will be generated. This overrides the skip_embedding_generation flag.

#### 🔄 Return

[UserFile](./src/main/java/com/konfigthis/client/model/UserFile.java)

#### 🌐 Endpoint

`/uploadfile` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.uploadFromUrl`

Create Upload File From Url

#### 🛠️ Usage

```java
UserFile result = client
        .files
        .uploadFromUrl(url)
        .fileName(fileName)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .setPageAsBoundary(setPageAsBoundary)
        .embeddingModel(embeddingModel)
        .generateSparseVectors(generateSparseVectors)
        .useTextract(useTextract)
        .prependFilenameToChunks(prependFilenameToChunks)
        .maxItemsPerChunk(maxItemsPerChunk)
        .parsePdfTablesWithOcr(parsePdfTablesWithOcr)
        .detectAudioLanguage(detectAudioLanguage)
        .transcriptionService(transcriptionService)
        .includeSpeakerLabels(includeSpeakerLabels)
        .mediaType(mediaType)
        .splitRows(splitRows)
        .coldStorageParams(coldStorageParams)
        .generateChunksOnly(generateChunksOnly)
        .execute();
```

#### ⚙️ Parameters

##### url: `String`

##### file_name: `String`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### set_page_as_boundary: `Boolean`

##### embedding_model:

##### generate_sparse_vectors: `Boolean`

##### use_textract: `Boolean`

##### prepend_filename_to_chunks: `Boolean`

##### max_items_per_chunk: `Integer`

Number of objects per chunk. For csv, tsv, xlsx, and json files only.

##### parse_pdf_tables_with_ocr: `Boolean`

##### detect_audio_language: `Boolean`

##### transcription_service:

##### include_speaker_labels: `Boolean`

##### media_type:

##### split_rows: `Boolean`

##### cold_storage_params: [`ColdStorageProps`](./src/main/java/com/konfigthis/client/model/ColdStorageProps.java)

##### generate_chunks_only: `Boolean`

If this flag is enabled, the file will be chunked and stored with Carbon,         but no embeddings will be generated. This overrides the skip_embedding_generation flag.

#### 🔄 Return

[UserFile](./src/main/java/com/konfigthis/client/model/UserFile.java)

#### 🌐 Endpoint

`/upload_file_from_url` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.files.uploadText`

Carbon supports multiple models for use in generating embeddings for files. For images, we support Vertex AI's
multimodal model; for text, we support OpenAI's `text-embedding-ada-002` and Cohere's embed-multilingual-v3.0.
The model can be specified via the `embedding_model` parameter (in the POST body for `/embeddings`, and a query 
parameter in `/uploadfile`). If no model is supplied, the `text-embedding-ada-002` is used by default. When performing
embedding queries, embeddings from files that used the specified model will be considered in the query.
For example, if files A and B have embeddings generated with `OPENAI`, and files C and D have embeddings generated with
`COHERE_MULTILINGUAL_V3`, then by default, queries will only consider files A and B. If `COHERE_MULTILINGUAL_V3` is
specified as the `embedding_model` in `/embeddings`, then only files C and D will be considered. Make sure that
the set of all files you want considered for a query have embeddings generated via the same model. For now, **do not**
set `VERTEX_MULTIMODAL` as an `embedding_model`. This model is used automatically by Carbon when it detects an image file.

#### 🛠️ Usage

```java
UserFile result = client
        .files
        .uploadText(contents)
        .name(name)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .overwriteFileId(overwriteFileId)
        .embeddingModel(embeddingModel)
        .generateSparseVectors(generateSparseVectors)
        .coldStorageParams(coldStorageParams)
        .generateChunksOnly(generateChunksOnly)
        .execute();
```

#### ⚙️ Parameters

##### contents: `String`

##### name: `String`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### overwrite_file_id: `Integer`

##### embedding_model:

##### generate_sparse_vectors: `Boolean`

##### cold_storage_params: [`ColdStorageProps`](./src/main/java/com/konfigthis/client/model/ColdStorageProps.java)

##### generate_chunks_only: `Boolean`

If this flag is enabled, the file will be chunked and stored with Carbon,         but no embeddings will be generated. This overrides the skip_embedding_generation flag.

#### 🔄 Return

[UserFile](./src/main/java/com/konfigthis/client/model/UserFile.java)

#### 🌐 Endpoint

`/upload_text` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.cancel`

Cancel Data Source Items Sync

#### 🛠️ Usage

```java
OrganizationUserDataSourceAPI result = client
        .integrations
        .cancel(dataSourceId)
        .execute();
```

#### ⚙️ Parameters

##### data_source_id: `Integer`

#### 🔄 Return

[OrganizationUserDataSourceAPI](./src/main/java/com/konfigthis/client/model/OrganizationUserDataSourceAPI.java)

#### 🌐 Endpoint

`/integrations/items/sync/cancel` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.connectDataSource`

Connect Data Source

#### 🛠️ Usage

```java
ConnectDataSourceResponse result = client
        .integrations
        .connectDataSource(authentication)
        .syncOptions(syncOptions)
        .execute();
```

#### ⚙️ Parameters

##### authentication: [`OANSCZGF`](./src/main/java/com/konfigthis/client/model/OANSCZGF.java)

##### sync_options: [`SyncOptions`](./src/main/java/com/konfigthis/client/model/SyncOptions.java)

#### 🔄 Return

[ConnectDataSourceResponse](./src/main/java/com/konfigthis/client/model/ConnectDataSourceResponse.java)

#### 🌐 Endpoint

`/integrations/connect` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.connectFreshdesk`

Refer this article to obtain an API key https://support.freshdesk.com/en/support/solutions/articles/215517.
Make sure that your API key has the permission to read solutions from your account and you are on a <b>paid</b> plan.
Once you have an API key, you can make a request to this endpoint along with your freshdesk domain. This will 
trigger an automatic sync of the articles in your "solutions" tab. Additional parameters below can be used to associate 
data with the synced articles or modify the sync behavior.

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .integrations
        .connectFreshdesk(domain, apiKey)
        .tags(tags)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .embeddingModel(embeddingModel)
        .generateSparseVectors(generateSparseVectors)
        .prependFilenameToChunks(prependFilenameToChunks)
        .syncFilesOnConnection(syncFilesOnConnection)
        .requestId(requestId)
        .syncSourceItems(syncSourceItems)
        .fileSyncConfig(fileSyncConfig)
        .execute();
```

#### ⚙️ Parameters

##### domain: `String`

##### api_key: `String`

##### tags: `Object`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### embedding_model:

##### generate_sparse_vectors: `Boolean`

##### prepend_filename_to_chunks: `Boolean`

##### sync_files_on_connection: `Boolean`

##### request_id: `String`

##### sync_source_items: `Boolean`

Enabling this flag will fetch all available content from the source to be listed via list items endpoint

##### file_sync_config: [`FileSyncConfigNullable`](./src/main/java/com/konfigthis/client/model/FileSyncConfigNullable.java)

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/integrations/freshdesk` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.connectGitbook`

You will need an access token to connect your Gitbook account. Note that the permissions will be defined by the user 
generating access token so make sure you have the permission to access spaces you will be syncing. 
Refer this article for more details https://developer.gitbook.com/gitbook-api/authentication. Additionally, you
need to specify the name of organization you will be syncing data from.

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .integrations
        .connectGitbook(organization, accessToken)
        .tags(tags)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .embeddingModel(embeddingModel)
        .generateSparseVectors(generateSparseVectors)
        .prependFilenameToChunks(prependFilenameToChunks)
        .syncFilesOnConnection(syncFilesOnConnection)
        .requestId(requestId)
        .syncSourceItems(syncSourceItems)
        .execute();
```

#### ⚙️ Parameters

##### organization: `String`

##### access_token: `String`

##### tags: `Object`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### embedding_model:

##### generate_sparse_vectors: `Boolean`

##### prepend_filename_to_chunks: `Boolean`

##### sync_files_on_connection: `Boolean`

##### request_id: `String`

##### sync_source_items: `Boolean`

Enabling this flag will fetch all available content from the source to be listed via list items endpoint

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/integrations/gitbook` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.connectGuru`

You will need an access token to connect your Guru account. To obtain an access token, follow the steps highlighted here
https://help.getguru.com/docs/gurus-api#obtaining-a-user-token. The username should be your Guru username.

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .integrations
        .connectGuru(username, accessToken)
        .tags(tags)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .embeddingModel(embeddingModel)
        .generateSparseVectors(generateSparseVectors)
        .prependFilenameToChunks(prependFilenameToChunks)
        .syncFilesOnConnection(syncFilesOnConnection)
        .requestId(requestId)
        .syncSourceItems(syncSourceItems)
        .fileSyncConfig(fileSyncConfig)
        .execute();
```

#### ⚙️ Parameters

##### username: `String`

##### access_token: `String`

##### tags: `Object`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### embedding_model:

##### generate_sparse_vectors: `Boolean`

##### prepend_filename_to_chunks: `Boolean`

##### sync_files_on_connection: `Boolean`

##### request_id: `String`

##### sync_source_items: `Boolean`

Enabling this flag will fetch all available content from the source to be listed via list items endpoint

##### file_sync_config: [`FileSyncConfigNullable`](./src/main/java/com/konfigthis/client/model/FileSyncConfigNullable.java)

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/integrations/guru` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.createAwsIamUser`

This endpoint can be used to connect S3 as well as Digital Ocean Spaces (S3 compatible)  
For S3, create a new IAM user with permissions to:
<ol>
<li>List all buckets.</li>
<li>Read from the specific buckets and objects to sync with Carbon. Ensure any future buckets or objects carry 
the same permissions.</li>
</ol>
Once created, generate an access key for this user and share the credentials with us. We recommend testing this key beforehand.  
For Digital Ocean Spaces, generate the above credentials in your Applications and API page here https://cloud.digitalocean.com/account/api/spaces.
Endpoint URL is required to connect Digital Ocean Spaces. It should look like <<region>>.digitaloceanspaces.com

#### 🛠️ Usage

```java
OrganizationUserDataSourceAPI result = client
        .integrations
        .createAwsIamUser(accessKey, accessKeySecret)
        .syncSourceItems(syncSourceItems)
        .endpointUrl(endpointUrl)
        .execute();
```

#### ⚙️ Parameters

##### access_key: `String`

##### access_key_secret: `String`

##### sync_source_items: `Boolean`

Enabling this flag will fetch all available content from the source to be listed via list items endpoint

##### endpoint_url: `String`

You can specify a Digital Ocean endpoint URL to connect a Digital Ocean Space through this endpoint.         The URL should be of format <region>.digitaloceanspaces.com. It's not required for S3 buckets.

#### 🔄 Return

[OrganizationUserDataSourceAPI](./src/main/java/com/konfigthis/client/model/OrganizationUserDataSourceAPI.java)

#### 🌐 Endpoint

`/integrations/s3` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.getOauthUrl`

This endpoint can be used to generate the following URLs
- An OAuth URL for OAuth based connectors
- A file syncing URL which skips the OAuth flow if the user already has a valid access token and takes them to the
success state.

#### 🛠️ Usage

```java
OuthURLResponse result = client
        .integrations
        .getOauthUrl(service)
        .tags(tags)
        .scope(scope)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .embeddingModel(embeddingModel)
        .zendeskSubdomain(zendeskSubdomain)
        .microsoftTenant(microsoftTenant)
        .sharepointSiteName(sharepointSiteName)
        .confluenceSubdomain(confluenceSubdomain)
        .generateSparseVectors(generateSparseVectors)
        .prependFilenameToChunks(prependFilenameToChunks)
        .maxItemsPerChunk(maxItemsPerChunk)
        .salesforceDomain(salesforceDomain)
        .syncFilesOnConnection(syncFilesOnConnection)
        .setPageAsBoundary(setPageAsBoundary)
        .dataSourceId(dataSourceId)
        .connectingNewAccount(connectingNewAccount)
        .requestId(requestId)
        .useOcr(useOcr)
        .parsePdfTablesWithOcr(parsePdfTablesWithOcr)
        .enableFilePicker(enableFilePicker)
        .syncSourceItems(syncSourceItems)
        .incrementalSync(incrementalSync)
        .fileSyncConfig(fileSyncConfig)
        .automaticallyOpenFilePicker(automaticallyOpenFilePicker)
        .servicenowCredentials(servicenowCredentials)
        .execute();
```

#### ⚙️ Parameters

##### service:

##### tags: `Object`

##### scope: `String`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### embedding_model:

##### zendesk_subdomain: `String`

##### microsoft_tenant: `String`

##### sharepoint_site_name: `String`

##### confluence_subdomain: `String`

##### generate_sparse_vectors: `Boolean`

##### prepend_filename_to_chunks: `Boolean`

##### max_items_per_chunk: `Integer`

Number of objects per chunk. For csv, tsv, xlsx, and json files only.

##### salesforce_domain: `String`

##### sync_files_on_connection: `Boolean`

Used to specify whether Carbon should attempt to sync all your files automatically when authorization         is complete. This is only supported for a subset of connectors and will be ignored for the rest. Supported         connectors: Intercom, Zendesk, Gitbook, Confluence, Salesforce, Freshdesk

##### set_page_as_boundary: `Boolean`

##### data_source_id: `Integer`

Used to specify a data source to sync from if you have multiple connected. It can be skipped if          you only have one data source of that type connected or are connecting a new account.

##### connecting_new_account: `Boolean`

Used to connect a new data source. If not specified, we will attempt to create a sync URL         for an existing data source based on type and ID.

##### request_id: `String`

This request id will be added to all files that get synced using the generated OAuth URL

##### use_ocr: `Boolean`

Enable OCR for files that support it. Supported formats: pdf, png, jpg

##### parse_pdf_tables_with_ocr: `Boolean`

##### enable_file_picker: `Boolean`

Enable integration's file picker for sources that support it. Supported sources: BOX, DROPBOX, GOOGLE_DRIVE, ONEDRIVE, SHAREPOINT

##### sync_source_items: `Boolean`

Enabling this flag will fetch all available content from the source to be listed via list items endpoint

##### incremental_sync: `Boolean`

Only sync files if they have not already been synced or if the embedding properties have changed.         This flag is currently supported by ONEDRIVE, GOOGLE_DRIVE, BOX, DROPBOX, INTERCOM, GMAIL, OUTLOOK, ZENDESK, CONFLUENCE, NOTION, SHAREPOINT, SERVICENOW. It will be ignored for other data sources.

##### file_sync_config: [`FileSyncConfigNullable`](./src/main/java/com/konfigthis/client/model/FileSyncConfigNullable.java)

##### automatically_open_file_picker: `Boolean`

Automatically open source file picker after the OAuth flow is complete. This flag is currently supported by         BOX, DROPBOX, GOOGLE_DRIVE, ONEDRIVE, SHAREPOINT. It will be ignored for other data sources.

##### servicenow_credentials: [`ServiceNowCredentialsNullable`](./src/main/java/com/konfigthis/client/model/ServiceNowCredentialsNullable.java)

#### 🔄 Return

[OuthURLResponse](./src/main/java/com/konfigthis/client/model/OuthURLResponse.java)

#### 🌐 Endpoint

`/integrations/oauth_url` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.listConfluencePages`
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

This endpoint has been deprecated. Use /integrations/items/list instead.

To begin listing a user's Confluence pages, at least a `data_source_id` of a connected
Confluence account must be specified. This base request returns a list of root pages for
every space the user has access to in a Confluence instance. To traverse further down
the user's page directory, additional requests to this endpoint can be made with the same
`data_source_id` and with `parent_id` set to the id of page from a previous request. For
convenience, the `has_children` property in each directory item in the response list will
flag which pages will return non-empty lists of pages when set as the `parent_id`.

#### 🛠️ Usage

```java
ListResponse result = client
        .integrations
        .listConfluencePages(dataSourceId)
        .parentId(parentId)
        .execute();
```

#### ⚙️ Parameters

##### data_source_id: `Integer`

##### parent_id: `String`

#### 🔄 Return

[ListResponse](./src/main/java/com/konfigthis/client/model/ListResponse.java)

#### 🌐 Endpoint

`/integrations/confluence/list` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.listConversations`

List all of your public and private channels, DMs, and Group DMs. The ID from response 
can be used as a filter to sync messages to Carbon   
types: Comma separated list of types. Available types are im (DMs), mpim (group DMs), public_channel, and private_channel.
Defaults to public_channel.    
cursor: Used for pagination. If next_cursor is returned in response, you need to pass it as the cursor in the next request    
data_source_id: Data source needs to be specified if you have linked multiple slack accounts  
exclude_archived: Should archived conversations be excluded, defaults to true

#### 🛠️ Usage

```java
Object result = client
        .integrations
        .listConversations()
        .types(types)
        .cursor(cursor)
        .dataSourceId(dataSourceId)
        .excludeArchived(excludeArchived)
        .execute();
```

#### ⚙️ Parameters

##### types: `String`

##### cursor: `String`

##### dataSourceId: `Integer`

##### excludeArchived: `Boolean`

#### 🌐 Endpoint

`/integrations/slack/conversations` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.listDataSourceItems`

List Data Source Items

#### 🛠️ Usage

```java
ListDataSourceItemsResponse result = client
        .integrations
        .listDataSourceItems(dataSourceId)
        .parentId(parentId)
        .filters(filters)
        .pagination(pagination)
        .orderBy(orderBy)
        .orderDir(orderDir)
        .execute();
```

#### ⚙️ Parameters

##### data_source_id: `Integer`

##### parent_id: `String`

##### filters: [`ListItemsFiltersNullable`](./src/main/java/com/konfigthis/client/model/ListItemsFiltersNullable.java)

##### pagination: [`Pagination`](./src/main/java/com/konfigthis/client/model/Pagination.java)

##### order_by:

##### order_dir:

#### 🔄 Return

[ListDataSourceItemsResponse](./src/main/java/com/konfigthis/client/model/ListDataSourceItemsResponse.java)

#### 🌐 Endpoint

`/integrations/items/list` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.listFolders`

After connecting your Outlook account, you can use this endpoint to list all of your folders on outlook. This includes 
both system folders like "inbox" and user created folders.

#### 🛠️ Usage

```java
Object result = client
        .integrations
        .listFolders()
        .dataSourceId(dataSourceId)
        .execute();
```

#### ⚙️ Parameters

##### dataSourceId: `Integer`

#### 🌐 Endpoint

`/integrations/outlook/user_folders` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.listGitbookSpaces`

After connecting your Gitbook account, you can use this endpoint to list all of your spaces under current organization.

#### 🛠️ Usage

```java
Object result = client
        .integrations
        .listGitbookSpaces(dataSourceId)
        .execute();
```

#### ⚙️ Parameters

##### dataSourceId: `Integer`

#### 🌐 Endpoint

`/integrations/gitbook/spaces` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.listLabels`

After connecting your Gmail account, you can use this endpoint to list all of your labels. User created labels
will have the type "user" and Gmail's default labels will have the type "system"

#### 🛠️ Usage

```java
Object result = client
        .integrations
        .listLabels()
        .dataSourceId(dataSourceId)
        .execute();
```

#### ⚙️ Parameters

##### dataSourceId: `Integer`

#### 🌐 Endpoint

`/integrations/gmail/user_labels` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.listOutlookCategories`

After connecting your Outlook account, you can use this endpoint to list all of your categories on outlook. We currently
support listing up to 250 categories.

#### 🛠️ Usage

```java
Object result = client
        .integrations
        .listOutlookCategories()
        .dataSourceId(dataSourceId)
        .execute();
```

#### ⚙️ Parameters

##### dataSourceId: `Integer`

#### 🌐 Endpoint

`/integrations/outlook/user_categories` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.listRepos`

Once you have connected your GitHub account, you can use this endpoint to list the 
    repositories your account has access to. You can use a data source ID or username to fetch from a specific account.

#### 🛠️ Usage

```java
Object result = client
        .integrations
        .listRepos()
        .perPage(perPage)
        .page(page)
        .dataSourceId(dataSourceId)
        .execute();
```

#### ⚙️ Parameters

##### perPage: `Integer`

##### page: `Integer`

##### dataSourceId: `Integer`

#### 🌐 Endpoint

`/integrations/github/repos` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.syncConfluence`
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

This endpoint has been deprecated. Use /integrations/files/sync instead.

After listing pages in a user's Confluence account, the set of selected page `ids` and the
connected account's `data_source_id` can be passed into this endpoint to sync them into
Carbon. Additional parameters listed below can be used to associate data to the selected
pages or alter the behavior of the sync.

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .integrations
        .syncConfluence(dataSourceId, ids)
        .tags(tags)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .embeddingModel(embeddingModel)
        .generateSparseVectors(generateSparseVectors)
        .prependFilenameToChunks(prependFilenameToChunks)
        .maxItemsPerChunk(maxItemsPerChunk)
        .setPageAsBoundary(setPageAsBoundary)
        .requestId(requestId)
        .useOcr(useOcr)
        .parsePdfTablesWithOcr(parsePdfTablesWithOcr)
        .incrementalSync(incrementalSync)
        .fileSyncConfig(fileSyncConfig)
        .execute();
```

#### ⚙️ Parameters

##### data_source_id: `Integer`

##### ids: List<`Object`
                `Object`>

##### tags: `Object`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### embedding_model:

##### generate_sparse_vectors: `Boolean`

##### prepend_filename_to_chunks: `Boolean`

##### max_items_per_chunk: `Integer`

Number of objects per chunk. For csv, tsv, xlsx, and json files only.

##### set_page_as_boundary: `Boolean`

##### request_id: `String`

##### use_ocr: `Boolean`

##### parse_pdf_tables_with_ocr: `Boolean`

##### incremental_sync: `Boolean`

Only sync files if they have not already been synced or if the embedding properties have changed.         This flag is currently supported by ONEDRIVE, GOOGLE_DRIVE, BOX, DROPBOX, INTERCOM, GMAIL, OUTLOOK, ZENDESK, CONFLUENCE, NOTION, SHAREPOINT, SERVICENOW. It will be ignored for other data sources.

##### file_sync_config: [`FileSyncConfigNullable`](./src/main/java/com/konfigthis/client/model/FileSyncConfigNullable.java)

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/integrations/confluence/sync` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.syncDataSourceItems`

Sync Data Source Items

#### 🛠️ Usage

```java
OrganizationUserDataSourceAPI result = client
        .integrations
        .syncDataSourceItems(dataSourceId)
        .execute();
```

#### ⚙️ Parameters

##### data_source_id: `Integer`

#### 🔄 Return

[OrganizationUserDataSourceAPI](./src/main/java/com/konfigthis/client/model/OrganizationUserDataSourceAPI.java)

#### 🌐 Endpoint

`/integrations/items/sync` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.syncFiles`

After listing files and folders via /integrations/items/sync and integrations/items/list, use the selected items' external ids 
as the ids in this endpoint to sync them into Carbon. Sharepoint items take an additional parameter root_id, which identifies
the drive the file or folder is in and is stored in root_external_id. That additional paramter is optional and excluding it will
tell the sync to assume the item is stored in the default Documents drive.

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .integrations
        .syncFiles(dataSourceId, ids)
        .tags(tags)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .embeddingModel(embeddingModel)
        .generateSparseVectors(generateSparseVectors)
        .prependFilenameToChunks(prependFilenameToChunks)
        .maxItemsPerChunk(maxItemsPerChunk)
        .setPageAsBoundary(setPageAsBoundary)
        .requestId(requestId)
        .useOcr(useOcr)
        .parsePdfTablesWithOcr(parsePdfTablesWithOcr)
        .incrementalSync(incrementalSync)
        .fileSyncConfig(fileSyncConfig)
        .execute();
```

#### ⚙️ Parameters

##### data_source_id: `Integer`

##### ids: List<`Object`
                `Object`>

##### tags: `Object`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### embedding_model:

##### generate_sparse_vectors: `Boolean`

##### prepend_filename_to_chunks: `Boolean`

##### max_items_per_chunk: `Integer`

Number of objects per chunk. For csv, tsv, xlsx, and json files only.

##### set_page_as_boundary: `Boolean`

##### request_id: `String`

##### use_ocr: `Boolean`

##### parse_pdf_tables_with_ocr: `Boolean`

##### incremental_sync: `Boolean`

Only sync files if they have not already been synced or if the embedding properties have changed.         This flag is currently supported by ONEDRIVE, GOOGLE_DRIVE, BOX, DROPBOX, INTERCOM, GMAIL, OUTLOOK, ZENDESK, CONFLUENCE, NOTION, SHAREPOINT, SERVICENOW. It will be ignored for other data sources.

##### file_sync_config: [`FileSyncConfigNullable`](./src/main/java/com/konfigthis/client/model/FileSyncConfigNullable.java)

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/integrations/files/sync` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.syncGitHub`

Refer this article to obtain an access token https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens.
Make sure that your access token has the permission to read content from your desired repos. Note that if your access token
expires you will need to manually update it through this endpoint.

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .integrations
        .syncGitHub(username, accessToken)
        .syncSourceItems(syncSourceItems)
        .execute();
```

#### ⚙️ Parameters

##### username: `String`

##### access_token: `String`

##### sync_source_items: `Boolean`

Enabling this flag will fetch all available content from the source to be listed via list items endpoint

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/integrations/github` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.syncGitbook`

You can sync upto 20 Gitbook spaces at a time using this endpoint. Additional parameters below can be used to associate 
data with the synced pages or modify the sync behavior.

#### 🛠️ Usage

```java
Object result = client
        .integrations
        .syncGitbook(spaceIds, dataSourceId)
        .tags(tags)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .embeddingModel(embeddingModel)
        .generateSparseVectors(generateSparseVectors)
        .prependFilenameToChunks(prependFilenameToChunks)
        .requestId(requestId)
        .execute();
```

#### ⚙️ Parameters

##### space_ids: List<`String`>

##### data_source_id: `Integer`

##### tags: `Object`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### embedding_model:

##### generate_sparse_vectors: `Boolean`

##### prepend_filename_to_chunks: `Boolean`

##### request_id: `String`

#### 🌐 Endpoint

`/integrations/gitbook/sync` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.syncGmail`

Once you have successfully connected your gmail account, you can choose which emails to sync with us
using the filters parameter. Filters is a JSON object with key value pairs. It also supports AND and OR operations.
For now, we support a limited set of keys listed below.

<b>label</b>: Inbuilt Gmail labels, for example "Important" or a custom label you created.  
<b>after</b> or <b>before</b>: A date in YYYY/mm/dd format (example 2023/12/31). Gets emails after/before a certain date.
You can also use them in combination to get emails from a certain period.  
<b>is</b>: Can have the following values - starred, important, snoozed, and unread  
<b>from</b>: Email address of the sender  
<b>to</b>: Email address of the recipient  
<b>in</b>: Can have the following values - sent (sync emails sent by the user)  
<b>has</b>: Can have the following values - attachment (sync emails that have attachments)  

Using keys or values outside of the specified values can lead to unexpected behaviour.

An example of a basic query with filters can be
```json
{
    "filters": {
            "key": "label",
            "value": "Test"
        }
}
```
Which will list all emails that have the label "Test".

You can use AND and OR operation in the following way:
```json
{
    "filters": {
        "AND": [
            {
                "key": "after",
                "value": "2024/01/07"
            },
            {
                "OR": [
                    {
                        "key": "label",
                        "value": "Personal"
                    },
                    {
                        "key": "is",
                        "value": "starred"
                    }
                ]
            }
        ]
    }
}
```
This will return emails after 7th of Jan that are either starred or have the label "Personal". 
Note that this is the highest level of nesting we support, i.e. you can't add more AND/OR filters within the OR filter
in the above example.

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .integrations
        .syncGmail(filters)
        .tags(tags)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .embeddingModel(embeddingModel)
        .generateSparseVectors(generateSparseVectors)
        .prependFilenameToChunks(prependFilenameToChunks)
        .dataSourceId(dataSourceId)
        .requestId(requestId)
        .syncAttachments(syncAttachments)
        .fileSyncConfig(fileSyncConfig)
        .incrementalSync(incrementalSync)
        .execute();
```

#### ⚙️ Parameters

##### filters: `Object`

##### tags: `Object`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### embedding_model:

##### generate_sparse_vectors: `Boolean`

##### prepend_filename_to_chunks: `Boolean`

##### data_source_id: `Integer`

##### request_id: `String`

##### sync_attachments: `Boolean`

##### file_sync_config: [`FileSyncConfigNullable`](./src/main/java/com/konfigthis/client/model/FileSyncConfigNullable.java)

##### incremental_sync: `Boolean`

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/integrations/gmail/sync` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.syncOutlook`

Once you have successfully connected your Outlook account, you can choose which emails to sync with us
using the filters and folder parameter. "folder" should be the folder you want to sync from Outlook. By default
we get messages from your inbox folder.  
Filters is a JSON object with key value pairs. It also supports AND and OR operations.
For now, we support a limited set of keys listed below.

<b>category</b>: Custom categories that you created in Outlook.  
<b>after</b> or <b>before</b>: A date in YYYY/mm/dd format (example 2023/12/31). Gets emails after/before a certain date. You can also use them in combination to get emails from a certain period.    
<b>is</b>: Can have the following values: flagged  
<b>from</b>: Email address of the sender   

An example of a basic query with filters can be
```json
{
    "filters": {
            "key": "category",
            "value": "Test"
        }
}
```
Which will list all emails that have the category "Test".  

Specifying a custom folder in the same query
```json
{
    "folder": "Folder Name",
    "filters": {
            "key": "category",
            "value": "Test"
        }
}
```

You can use AND and OR operation in the following way:
```json
{
    "filters": {
        "AND": [
            {
                "key": "after",
                "value": "2024/01/07"
            },
            {
                "OR": [
                    {
                        "key": "category",
                        "value": "Personal"
                    },
                    {
                        "key": "category",
                        "value": "Test"
                    },
                ]
            }
        ]
    }
}
```
This will return emails after 7th of Jan that have either Personal or Test as category. 
Note that this is the highest level of nesting we support, i.e. you can't add more AND/OR filters within the OR filter
in the above example.

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .integrations
        .syncOutlook(filters)
        .tags(tags)
        .folder(folder)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .embeddingModel(embeddingModel)
        .generateSparseVectors(generateSparseVectors)
        .prependFilenameToChunks(prependFilenameToChunks)
        .dataSourceId(dataSourceId)
        .requestId(requestId)
        .syncAttachments(syncAttachments)
        .fileSyncConfig(fileSyncConfig)
        .incrementalSync(incrementalSync)
        .execute();
```

#### ⚙️ Parameters

##### filters: `Object`

##### tags: `Object`

##### folder: `String`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### embedding_model:

##### generate_sparse_vectors: `Boolean`

##### prepend_filename_to_chunks: `Boolean`

##### data_source_id: `Integer`

##### request_id: `String`

##### sync_attachments: `Boolean`

##### file_sync_config: [`FileSyncConfigNullable`](./src/main/java/com/konfigthis/client/model/FileSyncConfigNullable.java)

##### incremental_sync: `Boolean`

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/integrations/outlook/sync` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.syncRepos`

You can retreive repos your token has access to using /integrations/github/repos and sync their content. 
You can also pass full name of any public repository (username/repo-name). This will store the repo content with 
carbon which can be accessed through /integrations/items/list endpoint. Maximum of 25 repositories are accepted per request.

#### 🛠️ Usage

```java
Object result = client
        .integrations
        .syncRepos(repos)
        .dataSourceId(dataSourceId)
        .execute();
```

#### ⚙️ Parameters

##### repos: List<`String`>

##### data_source_id: `Integer`

#### 🌐 Endpoint

`/integrations/github/sync_repos` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.syncRssFeed`

Rss Feed

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .integrations
        .syncRssFeed(url)
        .tags(tags)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .embeddingModel(embeddingModel)
        .generateSparseVectors(generateSparseVectors)
        .prependFilenameToChunks(prependFilenameToChunks)
        .requestId(requestId)
        .execute();
```

#### ⚙️ Parameters

##### url: `String`

##### tags: `Object`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### embedding_model:

##### generate_sparse_vectors: `Boolean`

##### prepend_filename_to_chunks: `Boolean`

##### request_id: `String`

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/integrations/rss_feed` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.syncS3Files`

After optionally loading the items via /integrations/items/sync and integrations/items/list, use the bucket name 
and object key as the ID in this endpoint to sync them into Carbon. Additional parameters below can associate 
data with the selected items or modify the sync behavior

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .integrations
        .syncS3Files(ids)
        .tags(tags)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .embeddingModel(embeddingModel)
        .generateSparseVectors(generateSparseVectors)
        .prependFilenameToChunks(prependFilenameToChunks)
        .maxItemsPerChunk(maxItemsPerChunk)
        .setPageAsBoundary(setPageAsBoundary)
        .dataSourceId(dataSourceId)
        .requestId(requestId)
        .useOcr(useOcr)
        .parsePdfTablesWithOcr(parsePdfTablesWithOcr)
        .fileSyncConfig(fileSyncConfig)
        .execute();
```

#### ⚙️ Parameters

##### ids: List<[`S3GetFileInput`](./src/main/java/com/konfigthis/client/model/S3GetFileInput.java)>

##### tags: `Object`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### embedding_model:

##### generate_sparse_vectors: `Boolean`

##### prepend_filename_to_chunks: `Boolean`

##### max_items_per_chunk: `Integer`

Number of objects per chunk. For csv, tsv, xlsx, and json files only.

##### set_page_as_boundary: `Boolean`

##### data_source_id: `Integer`

##### request_id: `String`

##### use_ocr: `Boolean`

##### parse_pdf_tables_with_ocr: `Boolean`

##### file_sync_config: [`FileSyncConfigNullable`](./src/main/java/com/konfigthis/client/model/FileSyncConfigNullable.java)

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/integrations/s3/files` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.integrations.syncSlack`

You can list all conversations using the endpoint /integrations/slack/conversations. The ID of 
conversation will be used as an input for this endpoint with timestamps as optional filters.

#### 🛠️ Usage

```java
Object result = client
        .integrations
        .syncSlack(filters)
        .tags(tags)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .embeddingModel(embeddingModel)
        .generateSparseVectors(generateSparseVectors)
        .prependFilenameToChunks(prependFilenameToChunks)
        .dataSourceId(dataSourceId)
        .requestId(requestId)
        .execute();
```

#### ⚙️ Parameters

##### filters: [`SlackFilters`](./src/main/java/com/konfigthis/client/model/SlackFilters.java)

##### tags: `Object`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### embedding_model:

##### generate_sparse_vectors: `Boolean`

##### prepend_filename_to_chunks: `Boolean`

##### data_source_id: `Integer`

##### request_id: `String`

#### 🌐 Endpoint

`/integrations/slack/sync` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.organizations.get`

Get Organization

#### 🛠️ Usage

```java
OrganizationResponse result = client
        .organizations
        .get()
        .execute();
```

#### 🔄 Return

[OrganizationResponse](./src/main/java/com/konfigthis/client/model/OrganizationResponse.java)

#### 🌐 Endpoint

`/organization` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.organizations.update`

Update Organization

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .organizations
        .update()
        .globalUserConfig(globalUserConfig)
        .dataSourceConfigs(dataSourceConfigs)
        .execute();
```

#### ⚙️ Parameters

##### global_user_config: [`UserConfigurationNullable`](./src/main/java/com/konfigthis/client/model/UserConfigurationNullable.java)

##### data_source_configs: Map<string, [`DataSourceConfiguration`](./src/main/java/com/konfigthis/client/model/DataSourceConfiguration.java)>

Used to set organization level defaults for configuration related to data sources.

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/organization/update` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.organizations.updateStats`

Use this endpoint to reaggregate the statistics for an organization, for example aggregate_file_size. The reaggregation
process is asyncronous so a webhook will be sent with the event type being FILE_STATISTICS_AGGREGATED to notify when the
process is complee. After this aggregation is complete, the updated statistics can be retrieved using the /organization
endpoint. The response of /organization willalso contain a timestamp of the last time the statistics were reaggregated.

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .organizations
        .updateStats()
        .execute();
```

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/organization/statistics` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.users.delete`

Delete Users

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .users
        .delete(customerIds)
        .execute();
```

#### ⚙️ Parameters

##### customer_ids: List<`String`>

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/delete_users` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.users.get`

User Endpoint

#### 🛠️ Usage

```java
UserResponse result = client
        .users
        .get(customerId)
        .execute();
```

#### ⚙️ Parameters

##### customer_id: `String`

#### 🔄 Return

[UserResponse](./src/main/java/com/konfigthis/client/model/UserResponse.java)

#### 🌐 Endpoint

`/user` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.users.list`

List users within an organization

#### 🛠️ Usage

```java
UserListResponse result = client
        .users
        .list()
        .pagination(pagination)
        .filters(filters)
        .orderBy(orderBy)
        .orderDir(orderDir)
        .includeCount(includeCount)
        .execute();
```

#### ⚙️ Parameters

##### pagination: [`Pagination`](./src/main/java/com/konfigthis/client/model/Pagination.java)

##### filters: [`ListUsersFilters`](./src/main/java/com/konfigthis/client/model/ListUsersFilters.java)

##### order_by:

##### order_dir:

##### include_count: `Boolean`

#### 🔄 Return

[UserListResponse](./src/main/java/com/konfigthis/client/model/UserListResponse.java)

#### 🌐 Endpoint

`/list_users` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.users.toggleUserFeatures`
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

Toggle User Features

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .users
        .toggleUserFeatures(configurationKeyName, value)
        .execute();
```

#### ⚙️ Parameters

##### configuration_key_name: `String`

##### value: `Object`

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/modify_user_configuration` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.users.updateUsers`

Update Users

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .users
        .updateUsers(customerIds)
        .autoSyncEnabledSources(autoSyncEnabledSources)
        .maxFiles(maxFiles)
        .maxFilesPerUpload(maxFilesPerUpload)
        .execute();
```

#### ⚙️ Parameters

##### customer_ids: List<`String`>

List of organization supplied user IDs

##### auto_sync_enabled_sources: `Object`

##### max_files: `Integer`

Custom file upload limit for the user over *all* user's files across all uploads.          If set, then the user will not be allowed to upload more files than this limit. If not set, or if set to -1,         then the user will have no limit.

##### max_files_per_upload: `Integer`

Custom file upload limit for the user across a single upload.         If set, then the user will not be allowed to upload more files than this limit in a single upload. If not set,         or if set to -1, then the user will have no limit.

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/update_users` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.fetchUrls`
![Deprecated](https://img.shields.io/badge/deprecated-yellow)

Extracts all URLs from a webpage. 

Args:
    url (str): URL of the webpage

Returns:
    FetchURLsResponse: A response object with a list of URLs extracted from the webpage and the webpage content.

#### 🛠️ Usage

```java
FetchURLsResponse result = client
        .utilities
        .fetchUrls(url)
        .execute();
```

#### ⚙️ Parameters

##### url: `String`

#### 🔄 Return

[FetchURLsResponse](./src/main/java/com/konfigthis/client/model/FetchURLsResponse.java)

#### 🌐 Endpoint

`/fetch_urls` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.fetchWebpage`

Fetch Urls V2

#### 🛠️ Usage

```java
Object result = client
        .utilities
        .fetchWebpage(url)
        .execute();
```

#### ⚙️ Parameters

##### url: `String`

#### 🌐 Endpoint

`/fetch_webpage` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.fetchYoutubeTranscripts`

Fetches english transcripts from YouTube videos.

Args:
    id (str): The ID of the YouTube video. 
    raw (bool): Whether to return the raw transcript or not. Defaults to False.

Returns:
    dict: A dictionary with the transcript of the YouTube video.

#### 🛠️ Usage

```java
YoutubeTranscriptResponse result = client
        .utilities
        .fetchYoutubeTranscripts(id)
        .raw(raw)
        .execute();
```

#### ⚙️ Parameters

##### id: `String`

##### raw: `Boolean`

#### 🔄 Return

[YoutubeTranscriptResponse](./src/main/java/com/konfigthis/client/model/YoutubeTranscriptResponse.java)

#### 🌐 Endpoint

`/fetch_youtube_transcript` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.processSitemap`

Retrieves all URLs from a sitemap, which can subsequently be utilized with our `web_scrape` endpoint.

<!--Args:
    url (str): URL of the sitemap

Returns:
    dict: A dictionary with a list of URLs extracted from the sitemap.-->

#### 🛠️ Usage

```java
Object result = client
        .utilities
        .processSitemap(url)
        .execute();
```

#### ⚙️ Parameters

##### url: `String`

#### 🌐 Endpoint

`/process_sitemap` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.scrapeSitemap`

Extracts all URLs from a sitemap and performs a web scrape on each of them.

Args:
    sitemap_url (str): URL of the sitemap

Returns:
    dict: A response object with the status of the scraping job message.-->

#### 🛠️ Usage

```java
Object result = client
        .utilities
        .scrapeSitemap(url)
        .tags(tags)
        .maxPagesToScrape(maxPagesToScrape)
        .chunkSize(chunkSize)
        .chunkOverlap(chunkOverlap)
        .skipEmbeddingGeneration(skipEmbeddingGeneration)
        .enableAutoSync(enableAutoSync)
        .generateSparseVectors(generateSparseVectors)
        .prependFilenameToChunks(prependFilenameToChunks)
        .htmlTagsToSkip(htmlTagsToSkip)
        .cssClassesToSkip(cssClassesToSkip)
        .cssSelectorsToSkip(cssSelectorsToSkip)
        .embeddingModel(embeddingModel)
        .urlPathsToInclude(urlPathsToInclude)
        .urlPathsToExclude(urlPathsToExclude)
        .urlsToScrape(urlsToScrape)
        .downloadCssAndMedia(downloadCssAndMedia)
        .generateChunksOnly(generateChunksOnly)
        .execute();
```

#### ⚙️ Parameters

##### url: `String`

##### tags: `Map<String, Object>`

##### max_pages_to_scrape: `Integer`

##### chunk_size: `Integer`

##### chunk_overlap: `Integer`

##### skip_embedding_generation: `Boolean`

##### enable_auto_sync: `Boolean`

##### generate_sparse_vectors: `Boolean`

##### prepend_filename_to_chunks: `Boolean`

##### html_tags_to_skip: List<`String`>

##### css_classes_to_skip: List<`String`>

##### css_selectors_to_skip: List<`String`>

##### embedding_model:

##### url_paths_to_include: List<`String`>

URL subpaths or directories that you want to include. For example if you want to only include         URLs that start with /questions in stackoverflow.com, you will add /questions/ in this input

##### url_paths_to_exclude: List<`String`>

URL subpaths or directories that you want to exclude. For example if you want to exclude         URLs that start with /questions in stackoverflow.com, you will add /questions/ in this input

##### urls_to_scrape: List<`String`>

You can submit a subset of URLs from the sitemap that should be scraped. To get the list of URLs,           you can check out /process_sitemap endpoint. If left empty, all URLs from the sitemap will be scraped.

##### download_css_and_media: `Boolean`

Whether the scraper should download css and media from the page (images, fonts, etc). Scrapes          might take longer to finish with this flag enabled, but the success rate is improved.

##### generate_chunks_only: `Boolean`

If this flag is enabled, the file will be chunked and stored with Carbon,           but no embeddings will be generated. This overrides the skip_embedding_generation flag.

#### 🌐 Endpoint

`/scrape_sitemap` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.scrapeWeb`

Conduct a web scrape on a given webpage URL. Our web scraper is fully compatible with JavaScript and supports recursion depth, enabling you to efficiently extract all content from the target website.

<!--Args:
    scraping_requests (List[WebscrapeRequest]): A list of WebscrapeRequest objects.

    
Returns:
    dict: A response object with the status of the scraping job message.-->

#### 🛠️ Usage

```java
Object result = client
        .utilities
        .scrapeWeb()
        .execute();
```

#### ⚙️ Request Body

List<[`WebscrapeRequest`](./src/main/java/com/konfigthis/client/model/WebscrapeRequest.java)>

#### 🌐 Endpoint

`/web_scrape` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.searchUrls`

Perform a web search and obtain a list of relevant URLs.

As an illustration, when you perform a search for “content related to MRNA,” you will receive a list of links such as the following:

    - https://tomrenz.substack.com/p/mrna-and-why-it-matters

    - https://www.statnews.com/2020/11/10/the-story-of-mrna-how-a-once-dismissed-idea-became-a-leading-technology-in-the-covid-vaccine-race/

    - https://www.statnews.com/2022/11/16/covid-19-vaccines-were-a-success-but-mrna-still-has-a-delivery-problem/
    
    - https://joomi.substack.com/p/were-still-being-misled-about-how

Subsequently, you can submit these links to the web_scrape endpoint in order to retrieve the content of the respective web pages.

Args:
    query (str): Query to search for

Returns:
    FetchURLsResponse: A response object with a list of URLs for a given search query.

#### 🛠️ Usage

```java
FetchURLsResponse result = client
        .utilities
        .searchUrls(query)
        .execute();
```

#### ⚙️ Parameters

##### query: `String`

#### 🔄 Return

[FetchURLsResponse](./src/main/java/com/konfigthis/client/model/FetchURLsResponse.java)

#### 🌐 Endpoint

`/search_urls` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.utilities.userWebpages`

User Web Pages

#### 🛠️ Usage

```java
Object result = client
        .utilities
        .userWebpages()
        .filters(filters)
        .pagination(pagination)
        .orderBy(orderBy)
        .orderDir(orderDir)
        .execute();
```

#### ⚙️ Parameters

##### filters: [`UserWebPagesFilters`](./src/main/java/com/konfigthis/client/model/UserWebPagesFilters.java)

##### pagination: [`Pagination`](./src/main/java/com/konfigthis/client/model/Pagination.java)

##### order_by:

##### order_dir:

#### 🌐 Endpoint

`/user_webpages` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.webhooks.addUrl`

Add Webhook Url

#### 🛠️ Usage

```java
Webhook result = client
        .webhooks
        .addUrl(url)
        .execute();
```

#### ⚙️ Parameters

##### url: `String`

#### 🔄 Return

[Webhook](./src/main/java/com/konfigthis/client/model/Webhook.java)

#### 🌐 Endpoint

`/add_webhook` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.webhooks.deleteUrl`

Delete Webhook Url

#### 🛠️ Usage

```java
GenericSuccessResponse result = client
        .webhooks
        .deleteUrl(webhookId)
        .execute();
```

#### ⚙️ Parameters

##### webhookId: `Integer`

#### 🔄 Return

[GenericSuccessResponse](./src/main/java/com/konfigthis/client/model/GenericSuccessResponse.java)

#### 🌐 Endpoint

`/delete_webhook/{webhook_id}` `DELETE`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `carbon.webhooks.urls`

Webhook Urls

#### 🛠️ Usage

```java
WebhookQueryResponse result = client
        .webhooks
        .urls()
        .pagination(pagination)
        .orderBy(orderBy)
        .orderDir(orderDir)
        .filters(filters)
        .execute();
```

#### ⚙️ Parameters

##### pagination: [`Pagination`](./src/main/java/com/konfigthis/client/model/Pagination.java)

##### order_by:

##### order_dir:

##### filters: [`WebhookFilters`](./src/main/java/com/konfigthis/client/model/WebhookFilters.java)

#### 🔄 Return

[WebhookQueryResponse](./src/main/java/com/konfigthis/client/model/WebhookQueryResponse.java)

#### 🌐 Endpoint

`/webhooks` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


## Author
This Java package is automatically generated by [Konfig](https://konfigthis.com)
