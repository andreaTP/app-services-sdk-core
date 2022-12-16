"""
    Service Registry API

    Service Registry Instance API  NOTE: This API cannot be called directly from the portal.  # noqa: E501

    The version of the OpenAPI document: 2.2.5.Final
    Contact: apicurio@lists.jboss.org
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from rhoas_registry_instance_sdk.api_client import ApiClient, Endpoint as _Endpoint
from rhoas_registry_instance_sdk.model_utils import (  # noqa: F401
    check_allowed_values,
    check_validations,
    date,
    datetime,
    file_type,
    none_type,
    validate_and_convert_types
)
from rhoas_registry_instance_sdk.model.artifact_search_results import ArtifactSearchResults
from rhoas_registry_instance_sdk.model.artifact_type import ArtifactType
from rhoas_registry_instance_sdk.model.error import Error
from rhoas_registry_instance_sdk.model.sort_by import SortBy
from rhoas_registry_instance_sdk.model.sort_order import SortOrder


class SearchApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.search_artifacts_endpoint = _Endpoint(
            settings={
                'response_type': (ArtifactSearchResults,),
                'auth': [],
                'endpoint_path': '/search/artifacts',
                'operation_id': 'search_artifacts',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'name',
                    'offset',
                    'limit',
                    'order',
                    'orderby',
                    'labels',
                    'properties',
                    'description',
                    'group',
                    'global_id',
                    'content_id',
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
                    'name':
                        (str,),
                    'offset':
                        (int,),
                    'limit':
                        (int,),
                    'order':
                        (SortOrder,),
                    'orderby':
                        (SortBy,),
                    'labels':
                        ([str],),
                    'properties':
                        ([str],),
                    'description':
                        (str,),
                    'group':
                        (str,),
                    'global_id':
                        (int,),
                    'content_id':
                        (int,),
                },
                'attribute_map': {
                    'name': 'name',
                    'offset': 'offset',
                    'limit': 'limit',
                    'order': 'order',
                    'orderby': 'orderby',
                    'labels': 'labels',
                    'properties': 'properties',
                    'description': 'description',
                    'group': 'group',
                    'global_id': 'globalId',
                    'content_id': 'contentId',
                },
                'location_map': {
                    'name': 'query',
                    'offset': 'query',
                    'limit': 'query',
                    'order': 'query',
                    'orderby': 'query',
                    'labels': 'query',
                    'properties': 'query',
                    'description': 'query',
                    'group': 'query',
                    'global_id': 'query',
                    'content_id': 'query',
                },
                'collection_format_map': {
                    'labels': 'multi',
                    'properties': 'multi',
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
        self.search_artifacts_by_content_endpoint = _Endpoint(
            settings={
                'response_type': (ArtifactSearchResults,),
                'auth': [],
                'endpoint_path': '/search/artifacts',
                'operation_id': 'search_artifacts_by_content',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'body',
                    'canonical',
                    'artifact_type',
                    'offset',
                    'limit',
                    'order',
                    'orderby',
                ],
                'required': [
                    'body',
                ],
                'nullable': [
                ],
                'enum': [
                    'order',
                    'orderby',
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                    ('order',): {

                        "ASC": "asc",
                        "DESC": "desc"
                    },
                    ('orderby',): {

                        "NAME": "name",
                        "CREATEDON": "createdOn"
                    },
                },
                'openapi_types': {
                    'body':
                        (file_type,),
                    'canonical':
                        (bool,),
                    'artifact_type':
                        (ArtifactType,),
                    'offset':
                        (int,),
                    'limit':
                        (int,),
                    'order':
                        (str,),
                    'orderby':
                        (str,),
                },
                'attribute_map': {
                    'canonical': 'canonical',
                    'artifact_type': 'artifactType',
                    'offset': 'offset',
                    'limit': 'limit',
                    'order': 'order',
                    'orderby': 'orderby',
                },
                'location_map': {
                    'body': 'body',
                    'canonical': 'query',
                    'artifact_type': 'query',
                    'offset': 'query',
                    'limit': 'query',
                    'order': 'query',
                    'orderby': 'query',
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

    def search_artifacts(
        self,
        **kwargs
    ):
        """Search for artifacts  # noqa: E501

        Returns a paginated list of all artifacts that match the provided filter criteria.   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.search_artifacts(async_req=True)
        >>> result = thread.get()


        Keyword Args:
            name (str): Filter by artifact name.. [optional]
            offset (int): The number of artifacts to skip before starting to collect the result set.  Defaults to 0.. [optional] if omitted the server will use the default value of 0
            limit (int): The number of artifacts to return.  Defaults to 20.. [optional] if omitted the server will use the default value of 20
            order (SortOrder): Sort order, ascending (`asc`) or descending (`desc`).. [optional]
            orderby (SortBy): The field to sort by.  Can be one of:  * `name` * `createdOn` . [optional]
            labels ([str]): Filter by label.  Include one or more label to only return artifacts containing all of the specified labels.. [optional]
            properties ([str]): Filter by one or more name/value property.  Separate each name/value pair using a colon.  For example `properties=foo:bar` will return only artifacts with a custom property named `foo` and value `bar`.. [optional]
            description (str): Filter by description.. [optional]
            group (str): Filter by artifact group.. [optional]
            global_id (int): Filter by globalId.. [optional]
            content_id (int): Filter by contentId.. [optional]
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
            ArtifactSearchResults
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
        return self.search_artifacts_endpoint.call_with_http_info(**kwargs)

    def search_artifacts_by_content(
        self,
        body,
        **kwargs
    ):
        """Search for artifacts by content  # noqa: E501

        Returns a paginated list of all artifacts with at least one version that matches the posted content.   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.search_artifacts_by_content(body, async_req=True)
        >>> result = thread.get()

        Args:
            body (file_type): The content to search for.

        Keyword Args:
            canonical (bool): Parameter that can be set to `true` to indicate that the server should \"canonicalize\" the content when searching for matching artifacts.  Canonicalization is unique to each artifact type, but typically involves removing any extra whitespace and formatting the content in a consistent manner.  Must be used along with the `artifactType` query parameter.. [optional]
            artifact_type (ArtifactType): Indicates the type of artifact represented by the content being used for the search.  This is only needed when using the `canonical` query parameter, so that the server knows how to canonicalize the content prior to searching for matching artifacts.. [optional]
            offset (int): The number of artifacts to skip before starting to collect the result set.  Defaults to 0.. [optional] if omitted the server will use the default value of 0
            limit (int): The number of artifacts to return.  Defaults to 20.. [optional] if omitted the server will use the default value of 20
            order (str): Sort order, ascending (`asc`) or descending (`desc`).. [optional]
            orderby (str): The field to sort by.  Can be one of:  * `name` * `createdOn` . [optional]
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
            ArtifactSearchResults
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
        kwargs['body'] = \
            body
        return self.search_artifacts_by_content_endpoint.call_with_http_info(**kwargs)

