/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { IssuePRNullable } from './issue-prnullable';
import { Label } from './label';
import { PRState } from './prstate';
import { User } from './user';
import { UserNullable } from './user-nullable';

/**
 * 
 * @export
 * @interface Issue
 */
export interface Issue {
    /**
     * 
     * @type {string}
     * @memberof Issue
     */
    'title': string;
    /**
     * 
     * @type {number}
     * @memberof Issue
     */
    'id': number;
    /**
     * 
     * @type {string}
     * @memberof Issue
     */
    'url': string;
    /**
     * 
     * @type {string}
     * @memberof Issue
     */
    'html_url': string;
    /**
     * 
     * @type {number}
     * @memberof Issue
     */
    'number': number;
    /**
     * 
     * @type {User}
     * @memberof Issue
     */
    'user': User;
    /**
     * 
     * @type {Array<Label>}
     * @memberof Issue
     */
    'labels': Array<Label>;
    /**
     * 
     * @type {PRState}
     * @memberof Issue
     */
    'state': PRState;
    /**
     * 
     * @type {boolean}
     * @memberof Issue
     */
    'locked': boolean;
    /**
     * 
     * @type {number}
     * @memberof Issue
     */
    'num_comments': number;
    /**
     * 
     * @type {string}
     * @memberof Issue
     */
    'created_at': string;
    /**
     * 
     * @type {string}
     * @memberof Issue
     */
    'updated_at': string;
    /**
     * 
     * @type {string}
     * @memberof Issue
     */
    'closed_at': string | null;
    /**
     * 
     * @type {boolean}
     * @memberof Issue
     */
    'draft': boolean;
    /**
     * 
     * @type {IssuePRNullable}
     * @memberof Issue
     */
    'pull_request': IssuePRNullable | null;
    /**
     * 
     * @type {string}
     * @memberof Issue
     */
    'body': string;
    /**
     * 
     * @type {UserNullable}
     * @memberof Issue
     */
    'closed_by': UserNullable | null;
}

