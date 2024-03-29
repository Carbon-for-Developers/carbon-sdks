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
 * @interface UserResponse
 */
export interface UserResponse {
    /**
     * 
     * @type {number}
     * @memberof UserResponse
     */
    'id': number;
    /**
     * 
     * @type {number}
     * @memberof UserResponse
     */
    'organization_id': number;
    /**
     * 
     * @type {string}
     * @memberof UserResponse
     */
    'organization_supplied_user_id': string;
    /**
     * 
     * @type {string}
     * @memberof UserResponse
     */
    'created_at': string;
    /**
     * 
     * @type {string}
     * @memberof UserResponse
     */
    'updated_at': string;
    /**
     * 
     * @type {string}
     * @memberof UserResponse
     */
    'deleted_at': string | null;
    /**
     * 
     * @type {number}
     * @memberof UserResponse
     */
    'num_files_synced': number;
    /**
     * 
     * @type {number}
     * @memberof UserResponse
     */
    'num_characters_synced': number;
    /**
     * 
     * @type {number}
     * @memberof UserResponse
     */
    'num_tokens_synced': number;
    /**
     * 
     * @type {Array<object>}
     * @memberof UserResponse
     */
    'unique_file_tags': Array<object>;
    /**
     * 
     * @type {object}
     * @memberof UserResponse
     */
    'enabled_features': object | null;
}

