/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { WebhookStatus } from './webhook-status';

/**
 * 
 * @export
 * @interface WebhookNoKey
 */
export interface WebhookNoKey {
    /**
     * 
     * @type {number}
     * @memberof WebhookNoKey
     */
    'id': number;
    /**
     * 
     * @type {number}
     * @memberof WebhookNoKey
     */
    'organization_id': number;
    /**
     * 
     * @type {string}
     * @memberof WebhookNoKey
     */
    'url': string;
    /**
     * 
     * @type {WebhookStatus}
     * @memberof WebhookNoKey
     */
    'status': WebhookStatus;
    /**
     * 
     * @type {string}
     * @memberof WebhookNoKey
     */
    'status_reason': string | null;
    /**
     * 
     * @type {string}
     * @memberof WebhookNoKey
     */
    'created_at': string;
    /**
     * 
     * @type {string}
     * @memberof WebhookNoKey
     */
    'updated_at': string;
}

