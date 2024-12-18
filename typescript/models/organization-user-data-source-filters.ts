/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { DataSourceTypeNullable } from './data-source-type-nullable';

/**
 * 
 * @export
 * @interface OrganizationUserDataSourceFilters
 */
export interface OrganizationUserDataSourceFilters {
    /**
     *          Tags to filter by. Supports logical AND and OR operations. Input should be like below:         {             \"OR\": [                 {                 \"key\": \"subject\",                 \"value\": \"holy-bible\",                 \"negate\": false                 },                 {                     \"key\": \"person-of-interest\",                     \"value\": \"jesus christ\",                     \"negate\": false                 },                 {                     \"key\": \"genre\",                     \"value\": \"fiction\",                     \"negate\": true                 }                 {                     \"AND\": [                         {                             \"key\": \"subject\",                             \"value\": \"tao-te-ching\",                             \"negate\": true                         },                         {                             \"key\": \"author\",                             \"value\": \"lao-tzu\",                             \"negate\": false                         }                     ]                 }             ]         }         For a single filter, the filter block can be placed within either an \"AND\" or \"OR\" block.         
     * @type {object}
     * @memberof OrganizationUserDataSourceFilters
     */
    'tags'?: object | null;
    /**
     * 
     * @type {DataSourceTypeNullable}
     * @memberof OrganizationUserDataSourceFilters
     */
    'source'?: DataSourceTypeNullable | null;
    /**
     * 
     * @type {Array<number>}
     * @memberof OrganizationUserDataSourceFilters
     */
    'ids'?: Array<number> | null;
    /**
     * 
     * @type {boolean}
     * @memberof OrganizationUserDataSourceFilters
     */
    'revoked_access'?: boolean | null;
}

