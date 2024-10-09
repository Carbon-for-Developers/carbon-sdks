/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { PRComment } from './prcomment';

/**
 * 
 * @export
 * @interface CommentsResponse
 */
export interface CommentsResponse {
    /**
     * 
     * @type {Array<PRComment>}
     * @memberof CommentsResponse
     */
    'data': Array<PRComment>;
    /**
     * 
     * @type {string}
     * @memberof CommentsResponse
     */
    'next_cursor': string | null;
}

