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
 * @interface S3AuthRequest
 */
export interface S3AuthRequest {
    /**
     * 
     * @type {string}
     * @memberof S3AuthRequest
     */
    'access_key': string;
    /**
     * 
     * @type {string}
     * @memberof S3AuthRequest
     */
    'access_key_secret': string;
    /**
     * Enabling this flag will fetch all available content from the source to be listed via list items endpoint
     * @type {boolean}
     * @memberof S3AuthRequest
     */
    'sync_source_items'?: boolean;
    /**
     * You can specify a Digital Ocean endpoint URL to connect a Digital Ocean Space through this endpoint.         The URL should be of format <region>.digitaloceanspaces.com. It\'s not required for S3 buckets.
     * @type {string}
     * @memberof S3AuthRequest
     */
    'endpoint_url'?: string | null;
    /**
     * Tags to be associated with the data source. If the data source already has tags set, then an upsert will be performed.
     * @type {object}
     * @memberof S3AuthRequest
     */
    'data_source_tags'?: object;
}

