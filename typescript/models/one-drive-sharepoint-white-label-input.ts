/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { OneDriveSharepointCredentials } from './one-drive-sharepoint-credentials';

/**
 * 
 * @export
 * @interface OneDriveSharepointWhiteLabelInput
 */
export interface OneDriveSharepointWhiteLabelInput {
    /**
     * 
     * @type {string}
     * @memberof OneDriveSharepointWhiteLabelInput
     */
    'data_source_type': OneDriveSharepointWhiteLabelInputDataSourceTypeEnum;
    /**
     * 
     * @type {OneDriveSharepointCredentials}
     * @memberof OneDriveSharepointWhiteLabelInput
     */
    'credentials': OneDriveSharepointCredentials;
}

type OneDriveSharepointWhiteLabelInputDataSourceTypeEnum = 'ONEDRIVE' | 'SHAREPOINT'


