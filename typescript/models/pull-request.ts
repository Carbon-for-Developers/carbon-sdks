/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { HeadOrBase } from './head-or-base';
import { Label } from './label';
import { PRState } from './prstate';
import { Team } from './team';
import { User } from './user';

/**
 * 
 * @export
 * @interface PullRequest
 */
export interface PullRequest {
    /**
     * 
     * @type {string}
     * @memberof PullRequest
     */
    'title': string;
    /**
     * 
     * @type {number}
     * @memberof PullRequest
     */
    'id': number;
    /**
     * 
     * @type {string}
     * @memberof PullRequest
     */
    'url': string;
    /**
     * 
     * @type {number}
     * @memberof PullRequest
     */
    'number': number;
    /**
     * 
     * @type {PRState}
     * @memberof PullRequest
     */
    'state': PRState;
    /**
     * 
     * @type {User}
     * @memberof PullRequest
     */
    'user': User;
    /**
     * 
     * @type {string}
     * @memberof PullRequest
     */
    'created_at': string;
    /**
     * 
     * @type {string}
     * @memberof PullRequest
     */
    'updated_at': string;
    /**
     * 
     * @type {string}
     * @memberof PullRequest
     */
    'closed_at': string | null;
    /**
     * 
     * @type {string}
     * @memberof PullRequest
     */
    'merged_at': string | null;
    /**
     * 
     * @type {Array<User>}
     * @memberof PullRequest
     */
    'requested_reviewers': Array<User>;
    /**
     * 
     * @type {Array<Team>}
     * @memberof PullRequest
     */
    'requested_teams': Array<Team>;
    /**
     * 
     * @type {Array<Label>}
     * @memberof PullRequest
     */
    'labels': Array<Label>;
    /**
     * 
     * @type {boolean}
     * @memberof PullRequest
     */
    'draft': boolean;
    /**
     * 
     * @type {HeadOrBase}
     * @memberof PullRequest
     */
    'head': HeadOrBase;
    /**
     * 
     * @type {HeadOrBase}
     * @memberof PullRequest
     */
    'base': HeadOrBase;
    /**
     * 
     * @type {object}
     * @memberof PullRequest
     */
    'remote_data': object | null;
}

