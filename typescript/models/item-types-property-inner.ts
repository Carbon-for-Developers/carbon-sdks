/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { AzureBlobStorageTypes } from './azure-blob-storage-types';
import { ConfluenceFileTypes } from './confluence-file-types';
import { GitbookFileTypes } from './gitbook-file-types';
import { GithubFileTypes } from './github-file-types';
import { GongFileTypes } from './gong-file-types';
import { GuruFileTypes } from './guru-file-types';
import { HelpdeskFileTypes } from './helpdesk-file-types';
import { IntercomFileTypes } from './intercom-file-types';
import { ServiceNowFileTypes } from './service-now-file-types';
import { SharedSourceItemTypes } from './shared-source-item-types';
import { ZendeskFileTypes } from './zendesk-file-types';

/**
 * @type ItemTypesPropertyInner
 * @export
 */
export type ItemTypesPropertyInner = SharedSourceItemTypes | HelpdeskFileTypes | ServiceNowFileTypes | ZendeskFileTypes | IntercomFileTypes | ConfluenceFileTypes | GitbookFileTypes | GithubFileTypes | GuruFileTypes | GongFileTypes | AzureBlobStorageTypes;


