/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"


/**
 * 
 * @export
 * @interface AzureBlobAuthRequest
 */
export interface AzureBlobAuthRequest {
    /**
     * 
     * @type {string}
     * @memberof AzureBlobAuthRequest
     */
    'account_name': string;
    /**
     * 
     * @type {string}
     * @memberof AzureBlobAuthRequest
     */
    'account_key': string;
    /**
     * 
     * @type {boolean}
     * @memberof AzureBlobAuthRequest
     */
    'sync_source_items'?: boolean;
    /**
     * Tags to be associated with the data source. If the data source already has tags set, then an upsert will be performed.
     * @type {object}
     * @memberof AzureBlobAuthRequest
     */
    'data_source_tags'?: object;
}

