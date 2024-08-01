/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { SentWebhookPayloadObject } from './sent-webhook-payload-object';

/**
 * 
 * @export
 * @interface SentWebhookPayload
 */
export interface SentWebhookPayload {
    /**
     * 
     * @type {string}
     * @memberof SentWebhookPayload
     */
    'webhook_type'?: SentWebhookPayloadWebhookTypeEnum;
    /**
     * 
     * @type {string}
     * @memberof SentWebhookPayload
     */
    'customer_id'?: string;
    /**
     * 
     * @type {string}
     * @memberof SentWebhookPayload
     */
    'timestamp'?: string;
    /**
     * 
     * @type {SentWebhookPayloadObject}
     * @memberof SentWebhookPayload
     */
    'object'?: SentWebhookPayloadObject;
}

type SentWebhookPayloadWebhookTypeEnum = 'ALL_UPLOADED_FILES_QUEUED' | 'CHECKUP' | 'DATA_SOURCE_READY' | 'FILES_CREATED' | 'FILES_SKIPPED' | 'FILE_DELETED' | 'FILE_ERROR' | 'FILE_READY' | 'FILE_STATISTICS_AGGREGATED' | 'FILE_SYNC_LIMIT_REACHED' | 'ORGANIZATION_USER_DELETED' | 'RATE_LIMIT_ERROR' | 'SPARSE_VECTOR_QUEUE_STATUS' | 'WEBPAGE_ERROR' | 'WEBPAGE_READY' | 'WEBSCRAPE_URLS_READY'


