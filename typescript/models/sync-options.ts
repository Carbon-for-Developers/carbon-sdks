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
 * @interface SyncOptions
 */
export interface SyncOptions {
    /**
     * 
     * @type {any}
     * @memberof SyncOptions
     */
    'tags'?: any | null;
    /**
     * 
     * @type {number}
     * @memberof SyncOptions
     */
    'chunk_size'?: number | null;
    /**
     * 
     * @type {number}
     * @memberof SyncOptions
     */
    'chunk_overlap'?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof SyncOptions
     */
    'skip_embedding_generation'?: boolean | null;
    /**
     * 
     * @type {EmbeddingGeneratorsNullable}
     * @memberof SyncOptions
     */
    'embedding_model'?: EmbeddingGeneratorsNullable | null;
    /**
     * 
     * @type {boolean}
     * @memberof SyncOptions
     */
    'generate_sparse_vectors'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof SyncOptions
     */
    'prepend_filename_to_chunks'?: boolean | null;
    /**
     * Number of objects per chunk. For csv, tsv, xlsx, and json files only.
     * @type {number}
     * @memberof SyncOptions
     */
    'max_items_per_chunk'?: number | null;
    /**
     * Used to specify whether Carbon should attempt to sync all your files automatically when authorization         is complete. This is only supported for a subset of connectors and will be ignored for the rest. Supported         connectors: Intercom, Zendesk, Gitbook, Confluence, Salesforce, Freshdesk
     * @type {boolean}
     * @memberof SyncOptions
     */
    'sync_files_on_connection'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof SyncOptions
     */
    'set_page_as_boundary'?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof SyncOptions
     */
    'enable_file_picker'?: boolean;
}

