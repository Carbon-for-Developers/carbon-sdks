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
 * @interface PullRequestExtended
 */
export interface PullRequestExtended {
    /**
     * 
     * @type {string}
     * @memberof PullRequestExtended
     */
    'title': string;
    /**
     * 
     * @type {number}
     * @memberof PullRequestExtended
     */
    'id': number;
    /**
     * 
     * @type {string}
     * @memberof PullRequestExtended
     */
    'url': string;
    /**
     * 
     * @type {number}
     * @memberof PullRequestExtended
     */
    'number': number;
    /**
     * 
     * @type {PRState}
     * @memberof PullRequestExtended
     */
    'state': PRState;
    /**
     * 
     * @type {User}
     * @memberof PullRequestExtended
     */
    'user': User;
    /**
     * 
     * @type {string}
     * @memberof PullRequestExtended
     */
    'created_at': string;
    /**
     * 
     * @type {string}
     * @memberof PullRequestExtended
     */
    'updated_at': string;
    /**
     * 
     * @type {string}
     * @memberof PullRequestExtended
     */
    'closed_at': string | null;
    /**
     * 
     * @type {string}
     * @memberof PullRequestExtended
     */
    'merged_at': string | null;
    /**
     * 
     * @type {Array<User>}
     * @memberof PullRequestExtended
     */
    'requested_reviewers': Array<User>;
    /**
     * 
     * @type {Array<Team>}
     * @memberof PullRequestExtended
     */
    'requested_teams': Array<Team>;
    /**
     * 
     * @type {Array<Label>}
     * @memberof PullRequestExtended
     */
    'labels': Array<Label>;
    /**
     * 
     * @type {boolean}
     * @memberof PullRequestExtended
     */
    'draft': boolean;
    /**
     * 
     * @type {HeadOrBase}
     * @memberof PullRequestExtended
     */
    'head': HeadOrBase;
    /**
     * 
     * @type {HeadOrBase}
     * @memberof PullRequestExtended
     */
    'base': HeadOrBase;
    /**
     * 
     * @type {object}
     * @memberof PullRequestExtended
     */
    'remote_data': object | null;
    /**
     * 
     * @type {boolean}
     * @memberof PullRequestExtended
     */
    'merged': boolean;
    /**
     * 
     * @type {number}
     * @memberof PullRequestExtended
     */
    'num_comments': number;
    /**
     * 
     * @type {number}
     * @memberof PullRequestExtended
     */
    'num_review_comments': number;
    /**
     * 
     * @type {number}
     * @memberof PullRequestExtended
     */
    'num_commits': number;
    /**
     * 
     * @type {number}
     * @memberof PullRequestExtended
     */
    'num_additions': number;
    /**
     * 
     * @type {number}
     * @memberof PullRequestExtended
     */
    'num_deletions': number;
    /**
     * 
     * @type {number}
     * @memberof PullRequestExtended
     */
    'num_changed_files': number;
}

