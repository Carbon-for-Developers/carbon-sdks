/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { OrganizationUserFilesToSyncFilters } from './organization-user-files-to-sync-filters';

/**
 * 
 * @export
 * @interface DeleteFilesV2QueryInput
 */
export interface DeleteFilesV2QueryInput {
    /**
     * 
     * @type {OrganizationUserFilesToSyncFilters}
     * @memberof DeleteFilesV2QueryInput
     */
    'filters'?: OrganizationUserFilesToSyncFilters;
    /**
     * 
     * @type {boolean}
     * @memberof DeleteFilesV2QueryInput
     */
    'send_webhook'?: boolean;
    /**
     * Whether or not to delete all data related to the file from the database, BUT to preserve the file metadata, allowing for         resyncs. By default `preserve_file_record` is false, which means that all data related to the file *as well as* its metadata will be deleted. Note that         even if `preserve_file_record` is true, raw files uploaded via the `uploadfile` endpoint still cannot be resynced.
     * @type {boolean}
     * @memberof DeleteFilesV2QueryInput
     */
    'preserve_file_record'?: boolean;
}

