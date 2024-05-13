/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { HelpdeskFileTypes } from './helpdesk-file-types';

/**
 * 
 * @export
 * @interface HelpdeskFileSyncConfig
 */
export interface HelpdeskFileSyncConfig {
    /**
     * 
     * @type {Array<HelpdeskFileTypes>}
     * @memberof HelpdeskFileSyncConfig
     */
    'file_types'?: Array<HelpdeskFileTypes>;
    /**
     * 
     * @type {boolean}
     * @memberof HelpdeskFileSyncConfig
     */
    'sync_attachments'?: boolean;
}
