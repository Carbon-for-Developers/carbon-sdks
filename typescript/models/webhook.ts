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
 * @interface Webhook
 */
export interface Webhook {
    /**
     * 
     * @type {number}
     * @memberof Webhook
     */
    'id': number;
    /**
     * 
     * @type {number}
     * @memberof Webhook
     */
    'organization_id': number;
    /**
     * 
     * @type {string}
     * @memberof Webhook
     */
    'url': string;
    /**
     * 
     * @type {string}
     * @memberof Webhook
     */
    'signing_key': string;
    /**
     * 
     * @type {string}
     * @memberof Webhook
     */
    'created_at': string;
    /**
     * 
     * @type {string}
     * @memberof Webhook
     */
    'updated_at': string;
}

