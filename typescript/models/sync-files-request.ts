/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { EmbeddingGeneratorsNullable } from './embedding-generators-nullable';

/**
 * 
 * @export
 * @interface SyncFilesRequest
 */
export interface SyncFilesRequest {
    /**
     * 
     * @type {object}
     * @memberof SyncFilesRequest
     */
    'tags'?: object | null;
    /**
     * 
     * @type {number}
     * @memberof SyncFilesRequest
     */
    'data_source_id': number;
    /**
     * 
     * @type {Array<string>}
     * @memberof SyncFilesRequest
     */
    'ids': Array<string>;
    /**
     * 
     * @type {number}
     * @memberof SyncFilesRequest
     */
    'chunk_size'?: number | null;
    /**
     * 
     * @type {number}
     * @memberof SyncFilesRequest
     */
    'chunk_overlap'?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof SyncFilesRequest
     */
    'skip_embedding_generation'?: boolean | null;
    /**
     * 
     * @type {EmbeddingGeneratorsNullable}
     * @memberof SyncFilesRequest
     */
    'embedding_model'?: EmbeddingGeneratorsNullable | null;
    /**
     * 
     * @type {boolean}
     * @memberof SyncFilesRequest
     */
    'generate_sparse_vectors'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof SyncFilesRequest
     */
    'prepend_filename_to_chunks'?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof SyncFilesRequest
     */
    'max_items_per_chunk'?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof SyncFilesRequest
     */
    'set_page_as_boundary'?: boolean;
}

