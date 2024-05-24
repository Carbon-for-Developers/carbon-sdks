/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { EmbeddingGenerators } from './embedding-generators';
import { FileSyncConfigNullable } from './file-sync-config-nullable';
import { S3GetFileInput } from './s3-get-file-input';

/**
 * 
 * @export
 * @interface S3FileSyncInput
 */
export interface S3FileSyncInput {
    /**
     * 
     * @type {object}
     * @memberof S3FileSyncInput
     */
    'tags'?: object | null;
    /**
     * 
     * @type {Array<S3GetFileInput>}
     * @memberof S3FileSyncInput
     */
    'ids': Array<S3GetFileInput>;
    /**
     * 
     * @type {number}
     * @memberof S3FileSyncInput
     */
    'chunk_size'?: number | null;
    /**
     * 
     * @type {number}
     * @memberof S3FileSyncInput
     */
    'chunk_overlap'?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof S3FileSyncInput
     */
    'skip_embedding_generation'?: boolean | null;
    /**
     * 
     * @type {EmbeddingGenerators}
     * @memberof S3FileSyncInput
     */
    'embedding_model'?: EmbeddingGenerators;
    /**
     * 
     * @type {boolean}
     * @memberof S3FileSyncInput
     */
    'generate_sparse_vectors'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof S3FileSyncInput
     */
    'prepend_filename_to_chunks'?: boolean | null;
    /**
     * Number of objects per chunk. For csv, tsv, xlsx, and json files only.
     * @type {number}
     * @memberof S3FileSyncInput
     */
    'max_items_per_chunk'?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof S3FileSyncInput
     */
    'set_page_as_boundary'?: boolean;
    /**
     * 
     * @type {number}
     * @memberof S3FileSyncInput
     */
    'data_source_id'?: number | null;
    /**
     * 
     * @type {string}
     * @memberof S3FileSyncInput
     */
    'request_id'?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof S3FileSyncInput
     */
    'use_ocr'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof S3FileSyncInput
     */
    'parse_pdf_tables_with_ocr'?: boolean | null;
    /**
     * 
     * @type {FileSyncConfigNullable}
     * @memberof S3FileSyncInput
     */
    'file_sync_config'?: FileSyncConfigNullable | null;
}

