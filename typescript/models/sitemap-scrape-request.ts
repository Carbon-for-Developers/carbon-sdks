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
 * @interface SitemapScrapeRequest
 */
export interface SitemapScrapeRequest {
    /**
     * 
     * @type {{ [key: string]: Tags1; }}
     * @memberof SitemapScrapeRequest
     */
    'tags'?: { [key: string]: Tags1; } | null;
    /**
     * 
     * @type {string}
     * @memberof SitemapScrapeRequest
     */
    'url': string;
    /**
     * 
     * @type {number}
     * @memberof SitemapScrapeRequest
     */
    'max_pages_to_scrape'?: number | null;
    /**
     * 
     * @type {number}
     * @memberof SitemapScrapeRequest
     */
    'chunk_size'?: number | null;
    /**
     * 
     * @type {number}
     * @memberof SitemapScrapeRequest
     */
    'chunk_overlap'?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof SitemapScrapeRequest
     */
    'skip_embedding_generation'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof SitemapScrapeRequest
     */
    'enable_auto_sync'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof SitemapScrapeRequest
     */
    'generate_sparse_vectors'?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof SitemapScrapeRequest
     */
    'prepend_filename_to_chunks'?: boolean | null;
    /**
     * 
     * @type {Array<string>}
     * @memberof SitemapScrapeRequest
     */
    'html_tags_to_skip'?: Array<string> | null;
    /**
     * 
     * @type {Array<string>}
     * @memberof SitemapScrapeRequest
     */
    'css_classes_to_skip'?: Array<string> | null;
    /**
     * 
     * @type {Array<string>}
     * @memberof SitemapScrapeRequest
     */
    'css_selectors_to_skip'?: Array<string> | null;
    /**
     * 
     * @type {EmbeddingGenerators}
     * @memberof SitemapScrapeRequest
     */
    'embedding_model'?: EmbeddingGenerators;
    /**
     * URL subpaths or directories that you want to include. For example if you want to only include         URLs that start with /questions in stackoverflow.com, you will add /questions/ in this input
     * @type {Array<string>}
     * @memberof SitemapScrapeRequest
     */
    'url_paths_to_include'?: Array<string> | null;
    /**
     * URL subpaths or directories that you want to exclude. For example if you want to exclude         URLs that start with /questions in stackoverflow.com, you will add /questions/ in this input
     * @type {Array<string>}
     * @memberof SitemapScrapeRequest
     */
    'url_paths_to_exclude'?: Array<string> | null;
    /**
     * You can submit a subset of URLs from the sitemap that should be scraped. To get the list of URLs,           you can check out /process_sitemap endpoint. If left empty, all URLs from the sitemap will be scraped.
     * @type {Array<string>}
     * @memberof SitemapScrapeRequest
     */
    'urls_to_scrape'?: Array<string> | null;
    /**
     * Whether the scraper should download css and media from the page (images, fonts, etc). Scrapes          might take longer to finish with this flag enabled, but the success rate is improved.
     * @type {boolean}
     * @memberof SitemapScrapeRequest
     */
    'download_css_and_media'?: boolean | null;
    /**
     * If this flag is enabled, the file will be chunked and stored with Carbon,           but no embeddings will be generated. This overrides the skip_embedding_generation flag.
     * @type {boolean}
     * @memberof SitemapScrapeRequest
     */
    'generate_chunks_only'?: boolean;
    /**
     * If this flag is enabled, the file will be stored with Carbon, but no processing will be done.
     * @type {boolean}
     * @memberof SitemapScrapeRequest
     */
    'store_file_only'?: boolean;
    /**
     * If the default proxies are blocked and not returning results, this flag can be enabled to use              alternate proxies (residential and office). Scrapes might take longer to finish with this flag enabled.         
     * @type {boolean}
     * @memberof SitemapScrapeRequest
     */
    'use_premium_proxies'?: boolean;
}

