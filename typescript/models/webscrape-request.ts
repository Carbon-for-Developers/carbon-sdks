/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { EmbeddingGenerators } from './embedding-generators';
import { Tags1 } from './tags1';

/**
 * 
 * @export
 * @interface WebscrapeRequest
 */
export interface WebscrapeRequest {
    /**
     * 
     * @type {{ [key: string]: Tags1; }}
     * @memberof WebscrapeRequest
     */
    'tags'?: { [key: string]: Tags1; } | null;
    /**
     * 
     * @type {string}
     * @memberof WebscrapeRequest
     */
    'url': string;
    /**
     * 
     * @type {number}
     * @memberof WebscrapeRequest
     */
    'recursion_depth'?: number | null;
    /**
     * 
     * @type {number}
     * @memberof WebscrapeRequest
     */
    'max_pages_to_scrape'?: number | null;
    /**
     * 
     * @type {number}
     * @memberof WebscrapeRequest
     */
    'chunk_size'?: number | null;
    /**
     * 
     * @type {number}
     * @memberof WebscrapeRequest
     */
    'chunk_overlap'?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof WebscrapeRequest
     */
    'skip_embedding_generation'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof WebscrapeRequest
     */
    'enable_auto_sync'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof WebscrapeRequest
     */
    'generate_sparse_vectors'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof WebscrapeRequest
     */
    'prepend_filename_to_chunks'?: boolean | null;
    /**
     * 
     * @type {Array<string>}
     * @memberof WebscrapeRequest
     */
    'html_tags_to_skip'?: Array<string> | null;
    /**
     * 
     * @type {Array<string>}
     * @memberof WebscrapeRequest
     */
    'css_classes_to_skip'?: Array<string> | null;
    /**
     * 
     * @type {Array<string>}
     * @memberof WebscrapeRequest
     */
    'css_selectors_to_skip'?: Array<string> | null;
    /**
     * 
     * @type {EmbeddingGenerators}
     * @memberof WebscrapeRequest
     */
    'embedding_model'?: EmbeddingGenerators;
    /**
     * URL subpaths or directories that you want to include. For example if you want to only include         URLs that start with /questions in stackoverflow.com, you will add /questions/ in this input
     * @type {Array<string>}
     * @memberof WebscrapeRequest
     */
    'url_paths_to_include'?: Array<string> | null;
    /**
     * Whether the scraper should download css and media from the page (images, fonts, etc). Scrapes          might take longer to finish with this flag enabled, but the success rate is improved.
     * @type {boolean}
     * @memberof WebscrapeRequest
     */
    'download_css_and_media'?: boolean | null;
    /**
     * If this flag is enabled, the file will be chunked and stored with Carbon,           but no embeddings will be generated. This overrides the skip_embedding_generation flag.
     * @type {boolean}
     * @memberof WebscrapeRequest
     */
    'generate_chunks_only'?: boolean;
}

