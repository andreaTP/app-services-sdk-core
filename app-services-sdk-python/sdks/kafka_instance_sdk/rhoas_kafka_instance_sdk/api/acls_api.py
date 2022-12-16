"""
    Kafka Instance API

    API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs  # noqa: E501

    The version of the OpenAPI document: 0.13.0-SNAPSHOT
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from rhoas_kafka_instance_sdk.api_client import ApiClient, Endpoint as _Endpoint
from rhoas_kafka_instance_sdk.model_utils import (  # noqa: F401
    check_allowed_values,
    check_validations,
    date,
    datetime,
    file_type,
    none_type,
    validate_and_convert_types
)
from rhoas_kafka_instance_sdk.model.acl_binding import AclBinding
from rhoas_kafka_instance_sdk.model.acl_binding_list_page import AclBindingListPage
from rhoas_kafka_instance_sdk.model.error import Error
from rhoas_kafka_instance_sdk.model.sort_direction import SortDirection


class AclsApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.create_acl_endpoint = _Endpoint(
            settings={
                'response_type': None,
                'auth': [
                    'Bearer',
                    'OAuth2'
                ],
                'endpoint_path': '/api/v1/acls',
                'operation_id': 'create_acl',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'acl_binding',
                ],
                'required': [
                    'acl_binding',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'acl_binding':
                        (AclBinding,),
                },
                'attribute_map': {
                },
                'location_map': {
                    'acl_binding': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [
                    'application/json'
                ]
            },
            api_client=api_client
        )
        self.delete_acls_endpoint = _Endpoint(
            settings={
                'response_type': (AclBindingListPage,),
                'auth': [
                    'Bearer',
                    'OAuth2'
                ],
                'endpoint_path': '/api/v1/acls',
                'operation_id': 'delete_acls',
                'http_method': 'DELETE',
                'servers': None,
            },
            params_map={
                'all': [
                    'resource_type',
                    'resource_name',
                    'pattern_type',
                    'principal',
                    'operation',
                    'permission',
                ],
                'required': [],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'resource_type':
                        (bool, date, datetime, dict, float, int, list, str, none_type,),
                    'resource_name':
                        (str,),
                    'pattern_type':
                        (bool, date, datetime, dict, float, int, list, str, none_type,),
                    'principal':
                        (str,),
                    'operation':
                        (bool, date, datetime, dict, float, int, list, str, none_type,),
                    'permission':
                        (bool, date, datetime, dict, float, int, list, str, none_type,),
                },
                'attribute_map': {
                    'resource_type': 'resourceType',
                    'resource_name': 'resourceName',
                    'pattern_type': 'patternType',
                    'principal': 'principal',
                    'operation': 'operation',
                    'permission': 'permission',
                },
                'location_map': {
                    'resource_type': 'query',
                    'resource_name': 'query',
                    'pattern_type': 'query',
                    'principal': 'query',
                    'operation': 'query',
                    'permission': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )
        self.get_acl_resource_operations_endpoint = _Endpoint(
            settings={
                'response_type': ({str: ([str],)},),
                'auth': [
                    'Bearer',
                    'OAuth2'
                ],
                'endpoint_path': '/api/v1/acls/resource-operations',
                'operation_id': 'get_acl_resource_operations',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                ],
                'required': [],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                },
                'attribute_map': {
                },
                'location_map': {
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )
        self.get_acls_endpoint = _Endpoint(
            settings={
                'response_type': (AclBindingListPage,),
                'auth': [
                    'Bearer',
                    'OAuth2'
                ],
                'endpoint_path': '/api/v1/acls',
                'operation_id': 'get_acls',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'resource_type',
                    'resource_name',
                    'pattern_type',
                    'principal',
                    'operation',
                    'permission',
                    'page',
                    'size',
                    'order',
                    'order_key',
                ],
                'required': [],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'page',
                    'size',
                ]
            },
            root_map={
                'validations': {
                    ('page',): {

                        'inclusive_minimum': 1,
                    },
                    ('size',): {

                        'inclusive_minimum': 1,
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'resource_type':
                        (bool, date, datetime, dict, float, int, list, str, none_type,),
                    'resource_name':
                        (str,),
                    'pattern_type':
                        (bool, date, datetime, dict, float, int, list, str, none_type,),
                    'principal':
                        (str,),
                    'operation':
                        (bool, date, datetime, dict, float, int, list, str, none_type,),
                    'permission':
                        (bool, date, datetime, dict, float, int, list, str, none_type,),
                    'page':
                        (int,),
                    'size':
                        (int,),
                    'order':
                        (SortDirection,),
                    'order_key':
                        (bool, date, datetime, dict, float, int, list, str, none_type,),
                },
                'attribute_map': {
                    'resource_type': 'resourceType',
                    'resource_name': 'resourceName',
                    'pattern_type': 'patternType',
                    'principal': 'principal',
                    'operation': 'operation',
                    'permission': 'permission',
                    'page': 'page',
                    'size': 'size',
                    'order': 'order',
                    'order_key': 'orderKey',
                },
                'location_map': {
                    'resource_type': 'query',
                    'resource_name': 'query',
                    'pattern_type': 'query',
                    'principal': 'query',
                    'operation': 'query',
                    'permission': 'query',
                    'page': 'query',
                    'size': 'query',
                    'order': 'query',
                    'order_key': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )

    def create_acl(
        self,
        acl_binding,
        **kwargs
    ):
        """Create ACL binding  # noqa: E501

        Creates a new ACL binding for a Kafka instance.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.create_acl(acl_binding, async_req=True)
        >>> result = thread.get()

        Args:
            acl_binding (AclBinding): ACL to create.

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            None
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        kwargs['acl_binding'] = \
            acl_binding
        return self.create_acl_endpoint.call_with_http_info(**kwargs)

    def delete_acls(
        self,
        **kwargs
    ):
        """Delete ACL bindings  # noqa: E501

        Deletes ACL bindings that match the query parameters.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.delete_acls(async_req=True)
        >>> result = thread.get()


        Keyword Args:
            resource_type (bool, date, datetime, dict, float, int, list, str, none_type): ACL Resource Type Filter. [optional]
            resource_name (str): ACL Resource Name Filter. [optional]
            pattern_type (bool, date, datetime, dict, float, int, list, str, none_type): ACL Pattern Type Filter. [optional]
            principal (str): ACL Principal Filter. Either a specific user or the wildcard user `User:*` may be provided. - When fetching by a specific user, the results will also include ACL bindings that apply to all users. - When deleting, ACL bindings to be delete must match the provided `principal` exactly.. [optional] if omitted the server will use the default value of ""
            operation (bool, date, datetime, dict, float, int, list, str, none_type): ACL Operation Filter. The ACL binding operation provided should be valid for the resource type in the request, if not `ANY`.. [optional]
            permission (bool, date, datetime, dict, float, int, list, str, none_type): ACL Permission Type Filter. [optional]
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            AclBindingListPage
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        return self.delete_acls_endpoint.call_with_http_info(**kwargs)

    def get_acl_resource_operations(
        self,
        **kwargs
    ):
        """Retrieve allowed ACL resources and operations  # noqa: E501

        Retrieve the resources and associated operations that may have ACLs configured.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.get_acl_resource_operations(async_req=True)
        >>> result = thread.get()


        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            {str: ([str],)}
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        return self.get_acl_resource_operations_endpoint.call_with_http_info(**kwargs)

    def get_acls(
        self,
        **kwargs
    ):
        """List ACL bindings  # noqa: E501

        Returns a list of all of the available ACL bindings, or the list of bindings that meet the user's URL query parameters. If no parameters are specified, all ACL bindings known to the system will be returned (with paging).  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.get_acls(async_req=True)
        >>> result = thread.get()


        Keyword Args:
            resource_type (bool, date, datetime, dict, float, int, list, str, none_type): ACL Resource Type Filter. [optional]
            resource_name (str): ACL Resource Name Filter. [optional]
            pattern_type (bool, date, datetime, dict, float, int, list, str, none_type): ACL Pattern Type Filter. [optional]
            principal (str): ACL Principal Filter. Either a specific user or the wildcard user `User:*` may be provided. - When fetching by a specific user, the results will also include ACL bindings that apply to all users. - When deleting, ACL bindings to be delete must match the provided `principal` exactly.. [optional] if omitted the server will use the default value of ""
            operation (bool, date, datetime, dict, float, int, list, str, none_type): ACL Operation Filter. The ACL binding operation provided should be valid for the resource type in the request, if not `ANY`.. [optional]
            permission (bool, date, datetime, dict, float, int, list, str, none_type): ACL Permission Type Filter. [optional]
            page (int): Page number. [optional]
            size (int): Number of records per page. [optional]
            order (SortDirection): Order items are sorted. [optional]
            order_key (bool, date, datetime, dict, float, int, list, str, none_type): [optional]
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            AclBindingListPage
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        return self.get_acls_endpoint.call_with_http_info(**kwargs)

