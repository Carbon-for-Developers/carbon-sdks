/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { ColdStorageProps } from './cold-storage-props';
import { EmbeddingGenerators } from './embedding-generators';
import { FileContentTypesNullable } from './file-content-types-nullable';
import { TranscriptionServiceNullable } from './transcription-service-nullable';

/**
 * 
 * @export
 * @interface UploadFileFromUrlInput
 */
export interface UploadFileFromUrlInput {
    /**
     * 
     * @type {string}
     * @memberof UploadFileFromUrlInput
     */
    'url': string;
    /**
     * 
     * @type {string}
     * @memberof UploadFileFromUrlInput
     */
    'file_name'?: string | null;
    /**
     * 
     * @type {number}
     * @memberof UploadFileFromUrlInput
     */
    'chunk_size'?: number | null;
    /**
     * 
     * @type {number}
     * @memberof UploadFileFromUrlInput
     */
    'chunk_overlap'?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof UploadFileFromUrlInput
     */
    'skip_embedding_generation'?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof UploadFileFromUrlInput
     */
    'set_page_as_boundary'?: boolean;
    /**
     * 
     * @type {EmbeddingGenerators}
     * @memberof UploadFileFromUrlInput
     */
    'embedding_model'?: EmbeddingGenerators;
    /**
     * 
     * @type {boolean}
     * @memberof UploadFileFromUrlInput
     */
    'generate_sparse_vectors'?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof UploadFileFromUrlInput
     */
    'use_textract'?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof UploadFileFromUrlInput
     */
    'prepend_filename_to_chunks'?: boolean;
    /**
     * Number of objects per chunk. For csv, tsv, xlsx, and json files only.
     * @type {number}
     * @memberof UploadFileFromUrlInput
     */
    'max_items_per_chunk'?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof UploadFileFromUrlInput
     */
    'parse_pdf_tables_with_ocr'?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof UploadFileFromUrlInput
     */
    'detect_audio_language'?: boolean;
    /**
     * 
     * @type {TranscriptionServiceNullable}
     * @memberof UploadFileFromUrlInput
     */
    'transcription_service'?: TranscriptionServiceNullable | null;
    /**
     * 
     * @type {boolean}
     * @memberof UploadFileFromUrlInput
     */
    'include_speaker_labels'?: boolean;
    /**
     * 
     * @type {FileContentTypesNullable}
     * @memberof UploadFileFromUrlInput
     */
    'media_type'?: FileContentTypesNullable | null;
    /**
     * 
     * @type {boolean}
     * @memberof UploadFileFromUrlInput
     */
    'split_rows'?: boolean;
    /**
     * 
     * @type {ColdStorageProps}
     * @memberof UploadFileFromUrlInput
     */
    'cold_storage_params'?: ColdStorageProps;
    /**
     * If this flag is enabled, the file will be chunked and stored with Carbon,         but no embeddings will be generated. This overrides the skip_embedding_generation flag.
     * @type {boolean}
     * @memberof UploadFileFromUrlInput
     */
    'generate_chunks_only'?: boolean;
    /**
     * If this flag is enabled, the file will be stored with Carbon, but no processing will be done.
     * @type {boolean}
     * @memberof UploadFileFromUrlInput
     */
    'store_file_only'?: boolean;
}

