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
 * @interface AzureBlobStorageAuthentication
 */
export interface AzureBlobStorageAuthentication {
    /**
     * 
     * @type {string}
     * @memberof AzureBlobStorageAuthentication
     */
    'source': AzureBlobStorageAuthenticationSourceEnum;
    /**
     * 
     * @type {string}
     * @memberof AzureBlobStorageAuthentication
     */
    'account_name': string;
    /**
     * 
     * @type {string}
     * @memberof AzureBlobStorageAuthentication
     */
    'account_key': string;
}

type AzureBlobStorageAuthenticationSourceEnum = 'AZURE_BLOB_STORAGE'


