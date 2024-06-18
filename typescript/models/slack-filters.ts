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
 * @interface SlackFilters
 */
export interface SlackFilters {
    /**
     * Conversation ID from slack, can be obtained from Carbon with slack conversations endpoint
     * @type {string}
     * @memberof SlackFilters
     */
    'conversation_id': string;
    /**
     * Only messages after this date will be included in results. If not set, all messages          will be included. Should be entered in YYYY/mm/dd format
     * @type {string}
     * @memberof SlackFilters
     */
    'after'?: string | null;
}

