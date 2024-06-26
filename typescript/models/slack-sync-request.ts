/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { EmbeddingGenerators } from './embedding-generators';
import { SlackFilters } from './slack-filters';

/**
 * 
 * @export
 * @interface SlackSyncRequest
 */
export interface SlackSyncRequest {
    /**
     * 
     * @type {object}
     * @memberof SlackSyncRequest
     */
    'tags'?: object | null;
    /**
     * 
     * @type {SlackFilters}
     * @memberof SlackSyncRequest
     */
    'filters': SlackFilters;
    /**
     * 
     * @type {number}
     * @memberof SlackSyncRequest
     */
    'chunk_size'?: number | null;
    /**
     * 
     * @type {number}
     * @memberof SlackSyncRequest
     */
    'chunk_overlap'?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof SlackSyncRequest
     */
    'skip_embedding_generation'?: boolean | null;
    /**
     * 
     * @type {EmbeddingGenerators}
     * @memberof SlackSyncRequest
     */
    'embedding_model'?: EmbeddingGenerators;
    /**
     * 
     * @type {boolean}
     * @memberof SlackSyncRequest
     */
    'generate_sparse_vectors'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof SlackSyncRequest
     */
    'prepend_filename_to_chunks'?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof SlackSyncRequest
     */
    'data_source_id'?: number | null;
    /**
     * 
     * @type {string}
     * @memberof SlackSyncRequest
     */
    'request_id'?: string | null;
}

