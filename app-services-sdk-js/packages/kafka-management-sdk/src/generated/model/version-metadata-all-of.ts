/* tslint:disable */
/* eslint-disable */
/**
 * Kafka Management API
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * The version of the OpenAPI document: 1.14.0
 * Contact: rhosak-support@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { ObjectReference } from './object-reference';

/**
 * 
 * @export
 * @interface VersionMetadataAllOf
 */
export interface VersionMetadataAllOf {
    /**
     * 
     * @type {string}
     * @memberof VersionMetadataAllOf
     */
    'server_version'?: string;
    /**
     * 
     * @type {Array<ObjectReference>}
     * @memberof VersionMetadataAllOf
     */
    'collections'?: Array<ObjectReference>;
}

