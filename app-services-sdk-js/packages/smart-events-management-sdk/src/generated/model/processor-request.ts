/* tslint:disable */
/* eslint-disable */
/**
 * Red Hat Openshift SmartEvents Fleet Manager V2
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * The version of the OpenAPI document: 0.0.1
 * Contact: openbridge-dev@redhat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */



/**
 * 
 * @export
 * @interface ProcessorRequest
 */
export interface ProcessorRequest {
    /**
     * The name of the processor
     * @type {string}
     * @memberof ProcessorRequest
     */
    'name': string;
    /**
     * The Camel YAML DSL code, formatted as JSON, that defines the flows in the processor
     * @type {object}
     * @memberof ProcessorRequest
     */
    'flows': object;
}

