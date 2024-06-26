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
 * @interface ListItemsFilters
 */
export interface ListItemsFilters {
    /**
     * 
     * @type {Array<string>}
     * @memberof ListItemsFilters
     */
    'external_ids'?: Array<string> | null;
    /**
     * 
     * @type {Array<number>}
     * @memberof ListItemsFilters
     */
    'ids'?: Array<number> | null;
    /**
     * 
     * @type {string}
     * @memberof ListItemsFilters
     */
    'name'?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof ListItemsFilters
     */
    'root_files_only'?: boolean | null;
}

