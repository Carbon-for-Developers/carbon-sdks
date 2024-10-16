/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/
import type * as buffer from "buffer"

import { OauthBasedConnectors } from './oauth-based-connectors';

/**
 * 
 * @export
 * @interface WhiteLabelFilters
 */
export interface WhiteLabelFilters {
    /**
     * 
     * @type {Array<number>}
     * @memberof WhiteLabelFilters
     */
    'ids'?: Array<number> | null;
    /**
     * 
     * @type {Array<OauthBasedConnectors>}
     * @memberof WhiteLabelFilters
     */
    'data_source_type'?: Array<OauthBasedConnectors> | null;
}

