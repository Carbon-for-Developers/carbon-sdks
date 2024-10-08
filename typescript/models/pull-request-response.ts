/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { PullRequest } from './pull-request';

/**
 * 
 * @export
 * @interface PullRequestResponse
 */
export interface PullRequestResponse {
    /**
     * 
     * @type {Array<PullRequest>}
     * @memberof PullRequestResponse
     */
    'data': Array<PullRequest>;
    /**
     * 
     * @type {string}
     * @memberof PullRequestResponse
     */
    'next_cursor': string | null;
}

