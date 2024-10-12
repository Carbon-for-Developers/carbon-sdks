/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { PartialAccountNullable } from './partial-account-nullable';
import { PartialContactNullable } from './partial-contact-nullable';
import { PartialOwner } from './partial-owner';

/**
 * 
 * @export
 * @interface Event
 */
export interface Event {
    /**
     * 
     * @type {string}
     * @memberof Event
     */
    'description': string | null;
    /**
     * 
     * @type {string}
     * @memberof Event
     */
    'id': string;
    /**
     * 
     * @type {PartialOwner}
     * @memberof Event
     */
    'owner': PartialOwner;
    /**
     * 
     * @type {string}
     * @memberof Event
     */
    'subject': string | null;
    /**
     * 
     * @type {string}
     * @memberof Event
     */
    'location': string | null;
    /**
     * 
     * @type {boolean}
     * @memberof Event
     */
    'is_all_day': boolean;
    /**
     * 
     * @type {string}
     * @memberof Event
     */
    'start_date': string | null;
    /**
     * 
     * @type {string}
     * @memberof Event
     */
    'end_date': string | null;
    /**
     * 
     * @type {PartialAccountNullable}
     * @memberof Event
     */
    'account': PartialAccountNullable | null;
    /**
     * 
     * @type {PartialContactNullable}
     * @memberof Event
     */
    'contact': PartialContactNullable | null;
    /**
     * 
     * @type {string}
     * @memberof Event
     */
    'created_at': string;
    /**
     * 
     * @type {string}
     * @memberof Event
     */
    'updated_at': string;
    /**
     * 
     * @type {boolean}
     * @memberof Event
     */
    'is_deleted': boolean;
    /**
     * 
     * @type {object}
     * @memberof Event
     */
    'remote_data': object | null;
}
