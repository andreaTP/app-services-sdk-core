/* tslint:disable */
/* eslint-disable */
/**
 * Kafka Management API
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * The version of the OpenAPI document: 1.16.0
 * Contact: rhosak-support@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */



/**
 * Schema for the request body sent to /kafkas POST
 * @export
 * @interface KafkaRequestPayload
 */
export interface KafkaRequestPayload {
    /**
     * The cloud provider where the Kafka cluster will be created in
     * @type {string}
     * @memberof KafkaRequestPayload
     */
    'cloud_provider'?: string;
    /**
     * The name of the Kafka cluster. It must consist of lower-case alphanumeric characters or \'-\', start with an alphabetic character, and end with an alphanumeric character, and can not be longer than 32 characters.
     * @type {string}
     * @memberof KafkaRequestPayload
     */
    'name': string;
    /**
     * The region where the Kafka cluster will be created in
     * @type {string}
     * @memberof KafkaRequestPayload
     */
    'region'?: string;
    /**
     * Whether connection reauthentication is enabled or not. If set to true, connection reauthentication on the Kafka instance will be required every 5 minutes. The default value is true
     * @type {boolean}
     * @memberof KafkaRequestPayload
     */
    'reauthentication_enabled'?: boolean | null;
    /**
     * kafka plan in a format of <instance_type>.<size_id>
     * @type {string}
     * @memberof KafkaRequestPayload
     */
    'plan'?: string;
    /**
     * cloud account id used to purchase the instance
     * @type {string}
     * @memberof KafkaRequestPayload
     */
    'billing_cloud_account_id'?: string | null;
    /**
     * marketplace where the instance is purchased on
     * @type {string}
     * @memberof KafkaRequestPayload
     */
    'marketplace'?: string | null;
    /**
     * billing model to use
     * @type {string}
     * @memberof KafkaRequestPayload
     */
    'billing_model'?: string | null;
    /**
     * enterprise OSD cluster ID to be used for kafka creation
     * @type {string}
     * @memberof KafkaRequestPayload
     */
    'cluster_id'?: string | null;
}

