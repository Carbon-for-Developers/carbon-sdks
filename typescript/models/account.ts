/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { Address } from './address';
import { Event } from './event';
import { PartialOwnerNullable } from './partial-owner-nullable';
import { PhoneNumber } from './phone-number';
import { Task } from './task';

/**
 * 
 * @export
 * @interface Account
 */
export interface Account {
    /**
     * 
     * @type {string}
     * @memberof Account
     */
    'description': string | null;
    /**
     * 
     * @type {string}
     * @memberof Account
     */
    'id': string;
    /**
     * 
     * @type {PartialOwnerNullable}
     * @memberof Account
     */
    'owner': PartialOwnerNullable | null;
    /**
     * 
     * @type {string}
     * @memberof Account
     */
    'name': string | null;
    /**
     * 
     * @type {string}
     * @memberof Account
     */
    'industry': string | null;
    /**
     * 
     * @type {string}
     * @memberof Account
     */
    'website': string | null;
    /**
     * 
     * @type {number}
     * @memberof Account
     */
    'number_of_employees': number | null;
    /**
     * 
     * @type {Array<Address>}
     * @memberof Account
     */
    'addresses': Array<Address>;
    /**
     * 
     * @type {Array<PhoneNumber>}
     * @memberof Account
     */
    'phone_numbers': Array<PhoneNumber>;
    /**
     * 
     * @type {string}
     * @memberof Account
     */
    'last_activity_at': string | null;
    /**
     * 
     * @type {string}
     * @memberof Account
     */
    'created_at': string;
    /**
     * 
     * @type {string}
     * @memberof Account
     */
    'updated_at': string;
    /**
     * 
     * @type {boolean}
     * @memberof Account
     */
    'is_deleted': boolean;
    /**
     * 
     * @type {Array<Task>}
     * @memberof Account
     */
    'tasks': Array<Task> | null;
    /**
     * 
     * @type {Array<Event>}
     * @memberof Account
     */
    'events': Array<Event> | null;
    /**
     * 
     * @type {object}
     * @memberof Account
     */
    'remote_data': object | null;
}

