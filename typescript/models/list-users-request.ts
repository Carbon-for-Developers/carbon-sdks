/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { ListUsersFilters } from './list-users-filters';
import { ListUsersOrderByTypes } from './list-users-order-by-types';
import { OrderDirV2 } from './order-dir-v2';
import { Pagination } from './pagination';

/**
 * 
 * @export
 * @interface ListUsersRequest
 */
export interface ListUsersRequest {
    /**
     * 
     * @type {Pagination}
     * @memberof ListUsersRequest
     */
    'pagination'?: Pagination;
    /**
     * 
     * @type {ListUsersFilters}
     * @memberof ListUsersRequest
     */
    'filters'?: ListUsersFilters;
    /**
     * 
     * @type {ListUsersOrderByTypes}
     * @memberof ListUsersRequest
     */
    'order_by'?: ListUsersOrderByTypes;
    /**
     * 
     * @type {OrderDirV2}
     * @memberof ListUsersRequest
     */
    'order_dir'?: OrderDirV2;
    /**
     * 
     * @type {boolean}
     * @memberof ListUsersRequest
     */
    'include_count'?: boolean;
}

