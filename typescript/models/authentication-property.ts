/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { ConfluenceAuthentication } from './confluence-authentication';
import { FreskdeskAuthentication } from './freskdesk-authentication';
import { GitbookAuthetication } from './gitbook-authetication';
import { GithubAuthentication } from './github-authentication';
import { NotionAuthentication } from './notion-authentication';
import { OAuthAuthentication } from './oauth-authentication';
import { S3Authentication } from './s3-authentication';
import { SalesforceAuthentication } from './salesforce-authentication';
import { SharepointAuthentication } from './sharepoint-authentication';
import { ZendeskAuthentication } from './zendesk-authentication';
import { ZoteroAuthentication } from './zotero-authentication';

/**
 * @type AuthenticationProperty
 * @export
 */
export type AuthenticationProperty = OAuthAuthentication | NotionAuthentication | SharepointAuthentication | ConfluenceAuthentication | ZendeskAuthentication | ZoteroAuthentication | GitbookAuthetication | SalesforceAuthentication | FreskdeskAuthentication | S3Authentication | GithubAuthentication;


