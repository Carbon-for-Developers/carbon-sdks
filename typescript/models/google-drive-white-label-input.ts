/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { GoogleDriveCredentials } from './google-drive-credentials';

/**
 * 
 * @export
 * @interface GoogleDriveWhiteLabelInput
 */
export interface GoogleDriveWhiteLabelInput {
    /**
     * 
     * @type {any}
     * @memberof GoogleDriveWhiteLabelInput
     */
    'data_source_type': any;
    /**
     * 
     * @type {GoogleDriveCredentials}
     * @memberof GoogleDriveWhiteLabelInput
     */
    'credentials': GoogleDriveCredentials;
}

