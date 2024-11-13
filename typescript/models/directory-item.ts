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
 * @interface DirectoryItem
 */
export interface DirectoryItem {
    /**
     * 
     * @type {string}
     * @memberof DirectoryItem
     */
    'id': string;
    /**
     * 
     * @type {string}
     * @memberof DirectoryItem
     */
    'name': string;
    /**
     * 
     * @type {boolean}
     * @memberof DirectoryItem
     */
    'is_synced'?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof DirectoryItem
     */
    'has_children'?: boolean;
}

