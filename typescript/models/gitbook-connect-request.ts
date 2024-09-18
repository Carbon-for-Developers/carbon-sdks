/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { EmbeddingGenerators } from './embedding-generators';
import { FileSyncConfigNullable } from './file-sync-config-nullable';

/**
 * 
 * @export
 * @interface GitbookConnectRequest
 */
export interface GitbookConnectRequest {
    /**
     * 
     * @type {object}
     * @memberof GitbookConnectRequest
     */
    'tags'?: object | null;
    /**
     * 
     * @type {string}
     * @memberof GitbookConnectRequest
     */
    'organization': string;
    /**
     * 
     * @type {string}
     * @memberof GitbookConnectRequest
     */
    'access_token': string;
    /**
     * 
     * @type {number}
     * @memberof GitbookConnectRequest
     */
    'chunk_size'?: number | null;
    /**
     * 
     * @type {number}
     * @memberof GitbookConnectRequest
     */
    'chunk_overlap'?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof GitbookConnectRequest
     */
    'skip_embedding_generation'?: boolean | null;
    /**
     * 
     * @type {EmbeddingGenerators}
     * @memberof GitbookConnectRequest
     */
    'embedding_model'?: EmbeddingGenerators;
    /**
     * 
     * @type {boolean}
     * @memberof GitbookConnectRequest
     */
    'generate_sparse_vectors'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof GitbookConnectRequest
     */
    'prepend_filename_to_chunks'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof GitbookConnectRequest
     */
    'sync_files_on_connection'?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof GitbookConnectRequest
     */
    'request_id'?: string | null;
    /**
     * Enabling this flag will fetch all available content from the source to be listed via list items endpoint
     * @type {boolean}
     * @memberof GitbookConnectRequest
     */
    'sync_source_items'?: boolean;
    /**
     * 
     * @type {FileSyncConfigNullable}
     * @memberof GitbookConnectRequest
     */
    'file_sync_config'?: FileSyncConfigNullable | null;
}

