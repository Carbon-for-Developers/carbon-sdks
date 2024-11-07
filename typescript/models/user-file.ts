/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { ChunkPropertiesNullable } from './chunk-properties-nullable';
import { DataSourceType } from './data-source-type';
import { EmbeddingProperties } from './embedding-properties';
import { EmbeddingStorageStatus } from './embedding-storage-status';
import { ExternalFileSyncStatuses } from './external-file-sync-statuses';
import { FileStatisticsNullable } from './file-statistics-nullable';

/**
 * 
 * @export
 * @interface UserFile
 */
export interface UserFile {
    /**
     * 
     * @type {object}
     * @memberof UserFile
     */
    'tags': object | null;
    /**
     * 
     * @type {number}
     * @memberof UserFile
     */
    'id': number;
    /**
     * 
     * @type {DataSourceType}
     * @memberof UserFile
     */
    'source': DataSourceType;
    /**
     * 
     * @type {number}
     * @memberof UserFile
     */
    'organization_id': number;
    /**
     * 
     * @type {number}
     * @memberof UserFile
     */
    'organization_user_id': number | null;
    /**
     * 
     * @type {string}
     * @memberof UserFile
     */
    'organization_supplied_user_id': string;
    /**
     * 
     * @type {number}
     * @memberof UserFile
     */
    'organization_user_data_source_id': number | null;
    /**
     * 
     * @type {string}
     * @memberof UserFile
     */
    'external_file_id': string;
    /**
     * 
     * @type {string}
     * @memberof UserFile
     */
    'external_url': string | null;
    /**
     * 
     * @type {ExternalFileSyncStatuses}
     * @memberof UserFile
     */
    'sync_status': ExternalFileSyncStatuses;
    /**
     * 
     * @type {string}
     * @memberof UserFile
     */
    'sync_error_message': string | null;
    /**
     * 
     * @type {string}
     * @memberof UserFile
     */
    'last_sync': string | null;
    /**
     * 
     * @type {FileStatisticsNullable}
     * @memberof UserFile
     */
    'file_statistics': FileStatisticsNullable | null;
    /**
     * 
     * @type {object}
     * @memberof UserFile
     */
    'file_metadata': object | null;
    /**
     * 
     * @type {{ [key: string]: EmbeddingProperties; }}
     * @memberof UserFile
     */
    'embedding_properties': { [key: string]: EmbeddingProperties; } | null;
    /**
     * 
     * @type {number}
     * @memberof UserFile
     */
    'chunk_size': number | null;
    /**
     * 
     * @type {number}
     * @memberof UserFile
     */
    'chunk_overlap': number | null;
    /**
     * 
     * @type {ChunkPropertiesNullable}
     * @memberof UserFile
     */
    'chunk_properties': ChunkPropertiesNullable | null;
    /**
     * 
     * @type {object}
     * @memberof UserFile
     */
    'ocr_properties': object;
    /**
     * 
     * @type {string}
     * @memberof UserFile
     */
    'ocr_job_started_at': string | null;
    /**
     * 
     * @type {string}
     * @memberof UserFile
     */
    'name': string | null;
    /**
     * 
     * @type {number}
     * @memberof UserFile
     */
    'parent_id': number | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserFile
     */
    'enable_auto_sync': boolean | null;
    /**
     * 
     * @type {string}
     * @memberof UserFile
     */
    'presigned_url': string | null;
    /**
     * 
     * @type {string}
     * @memberof UserFile
     */
    'parsed_text_url': string | null;
    /**
     * 
     * @type {object}
     * @memberof UserFile
     */
    'additional_presigned_urls': object | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserFile
     */
    'skip_embedding_generation': boolean;
    /**
     * 
     * @type {string}
     * @memberof UserFile
     */
    'source_created_at': string | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserFile
     */
    'generate_sparse_vectors': boolean | null;
    /**
     * 
     * @type {string}
     * @memberof UserFile
     */
    'request_id': string | null;
    /**
     * 
     * @type {string}
     * @memberof UserFile
     */
    'upload_id': string | null;
    /**
     * 
     * @type {object}
     * @memberof UserFile
     */
    'sync_properties': object;
    /**
     * 
     * @type {object}
     * @memberof UserFile
     */
    'messages_metadata': object;
    /**
     * 
     * @type {boolean}
     * @memberof UserFile
     */
    'file_contents_deleted': boolean;
    /**
     * 
     * @type {boolean}
     * @memberof UserFile
     */
    'supports_cold_storage': boolean;
    /**
     * 
     * @type {number}
     * @memberof UserFile
     */
    'hot_storage_time_to_live': number | null;
    /**
     * 
     * @type {EmbeddingStorageStatus}
     * @memberof UserFile
     */
    'embedding_storage_status': EmbeddingStorageStatus;
    /**
     * 
     * @type {string}
     * @memberof UserFile
     */
    'created_at': string;
    /**
     * 
     * @type {string}
     * @memberof UserFile
     */
    'updated_at': string;
}

