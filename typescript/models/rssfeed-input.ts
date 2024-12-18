/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { EmbeddingGenerators } from './embedding-generators';

/**
 * 
 * @export
 * @interface RSSFeedInput
 */
export interface RSSFeedInput {
    /**
     * 
     * @type {object}
     * @memberof RSSFeedInput
     */
    'tags'?: object | null;
    /**
     * 
     * @type {string}
     * @memberof RSSFeedInput
     */
    'url': string;
    /**
     * 
     * @type {number}
     * @memberof RSSFeedInput
     */
    'chunk_size'?: number | null;
    /**
     * 
     * @type {number}
     * @memberof RSSFeedInput
     */
    'chunk_overlap'?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof RSSFeedInput
     */
    'skip_embedding_generation'?: boolean | null;
    /**
     * 
     * @type {EmbeddingGenerators}
     * @memberof RSSFeedInput
     */
    'embedding_model'?: EmbeddingGenerators;
    /**
     * 
     * @type {boolean}
     * @memberof RSSFeedInput
     */
    'generate_sparse_vectors'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RSSFeedInput
     */
    'prepend_filename_to_chunks'?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof RSSFeedInput
     */
    'request_id'?: string | null;
    /**
     * Tags to be associated with the data source. If the data source already has tags set, then an upsert will be performed.
     * @type {object}
     * @memberof RSSFeedInput
     */
    'data_source_tags'?: object;
}

