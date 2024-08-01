// tslint:disable
/**
 * OpenAPI Spec for Solo APIs.
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: api_models
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * The response message containing the health status.
 * @export
 * @interface ApiModelsHealthCheckResponse
 */
export interface ApiModelsHealthCheckResponse  {
    /**
     * 
     * @type {boolean}
     * @memberof ApiModelsHealthCheckResponse
     */
    healthy?: boolean;
}

export function ApiModelsHealthCheckResponseFromJSON(json: any): ApiModelsHealthCheckResponse {
    return {
        'healthy': !exists(json, 'healthy') ? undefined : json['healthy'],
    };
}

export function ApiModelsHealthCheckResponseToJSON(value?: ApiModelsHealthCheckResponse): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'healthy': value.healthy,
    };
}


