/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { User } from './user';

/**
 * 
 * @export
 * @interface PRComment
 */
export interface PRComment {
    /**
     * 
     * @type {number}
     * @memberof PRComment
     */
    'id': number;
    /**
     * 
     * @type {number}
     * @memberof PRComment
     */
    'pull_request_review_id': number | null;
    /**
     * 
     * @type {string}
     * @memberof PRComment
     */
    'url': string;
    /**
     * 
     * @type {string}
     * @memberof PRComment
     */
    'diff_hunk': string;
    /**
     * 
     * @type {string}
     * @memberof PRComment
     */
    'path': string;
    /**
     * 
     * @type {User}
     * @memberof PRComment
     */
    'user': User;
    /**
     * 
     * @type {string}
     * @memberof PRComment
     */
    'body': string;
    /**
     * 
     * @type {string}
     * @memberof PRComment
     */
    'created_at': string;
    /**
     * 
     * @type {string}
     * @memberof PRComment
     */
    'updated_at': string;
    /**
     * 
     * @type {number}
     * @memberof PRComment
     */
    'start_line': number | null;
    /**
     * 
     * @type {number}
     * @memberof PRComment
     */
    'line': number | null;
    /**
     * 
     * @type {object}
     * @memberof PRComment
     */
    'remote_data': object | null;
}

