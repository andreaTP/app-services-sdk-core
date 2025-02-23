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



/**
 * 
 * @export
 * @interface ServiceAccountListItemAllOf
 */
export interface ServiceAccountListItemAllOf {
    /**
     * server generated unique id of the service account
     * @type {string}
     * @memberof ServiceAccountListItemAllOf
     */
    'id'?: string;
    /**
     * client id of the service account
     * @type {string}
     * @memberof ServiceAccountListItemAllOf
     */
    'client_id'?: string;
    /**
     * name of the service account
     * @type {string}
     * @memberof ServiceAccountListItemAllOf
     */
    'name'?: string;
    /**
     * owner of the service account
     * @type {string}
     * @memberof ServiceAccountListItemAllOf
     * @deprecated
     */
    'owner'?: string;
    /**
     * service account created by the user
     * @type {string}
     * @memberof ServiceAccountListItemAllOf
     */
    'created_by'?: string;
    /**
     * service account creation timestamp
     * @type {string}
     * @memberof ServiceAccountListItemAllOf
     */
    'created_at'?: string;
    /**
     * description of the service account
     * @type {string}
     * @memberof ServiceAccountListItemAllOf
     */
    'description'?: string;
}

