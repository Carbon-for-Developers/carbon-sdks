/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { OrderDirV2 } from './order-dir-v2';
import { Pagination } from './pagination';
import { UserWebPageOrderByTypes } from './user-web-page-order-by-types';
import { UserWebPagesFilters } from './user-web-pages-filters';

/**
 * 
 * @export
 * @interface UserWebPagesRequest
 */
export interface UserWebPagesRequest {
    /**
     * 
     * @type {UserWebPagesFilters}
     * @memberof UserWebPagesRequest
     */
    'filters'?: UserWebPagesFilters;
    /**
     * 
     * @type {Pagination}
     * @memberof UserWebPagesRequest
     */
    'pagination'?: Pagination;
    /**
     * 
     * @type {UserWebPageOrderByTypes}
     * @memberof UserWebPagesRequest
     */
    'order_by'?: UserWebPageOrderByTypes;
    /**
     * 
     * @type {OrderDirV2}
     * @memberof UserWebPagesRequest
     */
    'order_dir'?: OrderDirV2;
}

