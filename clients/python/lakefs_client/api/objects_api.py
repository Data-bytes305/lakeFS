"""
    lakeFS API

    lakeFS HTTP API  # noqa: E501

    The version of the OpenAPI document: 0.1.0
    Contact: services@treeverse.io
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from lakefs_client.api_client import ApiClient, Endpoint as _Endpoint
from lakefs_client.model_utils import (  # noqa: F401
    check_allowed_values,
    check_validations,
    date,
    datetime,
    file_type,
    none_type,
    validate_and_convert_types
)
from lakefs_client.model.error import Error
from lakefs_client.model.object_stage_creation import ObjectStageCreation
from lakefs_client.model.object_stats import ObjectStats
from lakefs_client.model.object_stats_list import ObjectStatsList
from lakefs_client.model.underlying_object_properties import UnderlyingObjectProperties


class ObjectsApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client

        def __delete_object(
            self,
            repository,
            branch,
            path,
            **kwargs
        ):
            """delete object  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.delete_object(repository, branch, path, async_req=True)
            >>> result = thread.get()

            Args:
                repository (str):
                branch (str):
                path (str):

            Keyword Args:
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
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
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['repository'] = \
                repository
            kwargs['branch'] = \
                branch
            kwargs['path'] = \
                path
            return self.call_with_http_info(**kwargs)

        self.delete_object = _Endpoint(
            settings={
                'response_type': None,
                'auth': [
                    'basic_auth',
                    'cookie_auth',
                    'jwt_token'
                ],
                'endpoint_path': '/repositories/{repository}/branches/{branch}/objects',
                'operation_id': 'delete_object',
                'http_method': 'DELETE',
                'servers': None,
            },
            params_map={
                'all': [
                    'repository',
                    'branch',
                    'path',
                ],
                'required': [
                    'repository',
                    'branch',
                    'path',
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
                    'repository':
                        (str,),
                    'branch':
                        (str,),
                    'path':
                        (str,),
                },
                'attribute_map': {
                    'repository': 'repository',
                    'branch': 'branch',
                    'path': 'path',
                },
                'location_map': {
                    'repository': 'path',
                    'branch': 'path',
                    'path': 'query',
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
            api_client=api_client,
            callable=__delete_object
        )

        def __get_object(
            self,
            repository,
            ref,
            path,
            **kwargs
        ):
            """get object content  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.get_object(repository, ref, path, async_req=True)
            >>> result = thread.get()

            Args:
                repository (str):
                ref (str): a reference (could be either a branch or a commit ID)
                path (str):

            Keyword Args:
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                file_type
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
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['repository'] = \
                repository
            kwargs['ref'] = \
                ref
            kwargs['path'] = \
                path
            return self.call_with_http_info(**kwargs)

        self.get_object = _Endpoint(
            settings={
                'response_type': (file_type,),
                'auth': [
                    'basic_auth',
                    'cookie_auth',
                    'jwt_token'
                ],
                'endpoint_path': '/repositories/{repository}/refs/{ref}/objects',
                'operation_id': 'get_object',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'repository',
                    'ref',
                    'path',
                ],
                'required': [
                    'repository',
                    'ref',
                    'path',
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
                    'repository':
                        (str,),
                    'ref':
                        (str,),
                    'path':
                        (str,),
                },
                'attribute_map': {
                    'repository': 'repository',
                    'ref': 'ref',
                    'path': 'path',
                },
                'location_map': {
                    'repository': 'path',
                    'ref': 'path',
                    'path': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/octet-stream',
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__get_object
        )

        def __get_underlying_properties(
            self,
            repository,
            ref,
            path,
            **kwargs
        ):
            """get object properties on underlying storage  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.get_underlying_properties(repository, ref, path, async_req=True)
            >>> result = thread.get()

            Args:
                repository (str):
                ref (str): a reference (could be either a branch or a commit ID)
                path (str):

            Keyword Args:
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                UnderlyingObjectProperties
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
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['repository'] = \
                repository
            kwargs['ref'] = \
                ref
            kwargs['path'] = \
                path
            return self.call_with_http_info(**kwargs)

        self.get_underlying_properties = _Endpoint(
            settings={
                'response_type': (UnderlyingObjectProperties,),
                'auth': [
                    'basic_auth',
                    'cookie_auth',
                    'jwt_token'
                ],
                'endpoint_path': '/repositories/{repository}/refs/{ref}/objects/underlyingProperties',
                'operation_id': 'get_underlying_properties',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'repository',
                    'ref',
                    'path',
                ],
                'required': [
                    'repository',
                    'ref',
                    'path',
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
                    'repository':
                        (str,),
                    'ref':
                        (str,),
                    'path':
                        (str,),
                },
                'attribute_map': {
                    'repository': 'repository',
                    'ref': 'ref',
                    'path': 'path',
                },
                'location_map': {
                    'repository': 'path',
                    'ref': 'path',
                    'path': 'query',
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
            api_client=api_client,
            callable=__get_underlying_properties
        )

        def __list_objects(
            self,
            repository,
            ref,
            **kwargs
        ):
            """list objects under a given prefix  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.list_objects(repository, ref, async_req=True)
            >>> result = thread.get()

            Args:
                repository (str):
                ref (str): a reference (could be either a branch or a commit ID)

            Keyword Args:
                after (str): return items after this value. [optional]
                amount (int): how many items to return. [optional] if omitted the server will use the default value of 100
                delimiter (str): delimiter used to group common prefixes by. [optional]
                prefix (str): return items prefixed with this value. [optional]
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                ObjectStatsList
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
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['repository'] = \
                repository
            kwargs['ref'] = \
                ref
            return self.call_with_http_info(**kwargs)

        self.list_objects = _Endpoint(
            settings={
                'response_type': (ObjectStatsList,),
                'auth': [
                    'basic_auth',
                    'cookie_auth',
                    'jwt_token'
                ],
                'endpoint_path': '/repositories/{repository}/refs/{ref}/objects/ls',
                'operation_id': 'list_objects',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'repository',
                    'ref',
                    'after',
                    'amount',
                    'delimiter',
                    'prefix',
                ],
                'required': [
                    'repository',
                    'ref',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'amount',
                ]
            },
            root_map={
                'validations': {
                    ('amount',): {

                        'inclusive_maximum': 1000,
                        'inclusive_minimum': -1,
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'repository':
                        (str,),
                    'ref':
                        (str,),
                    'after':
                        (str,),
                    'amount':
                        (int,),
                    'delimiter':
                        (str,),
                    'prefix':
                        (str,),
                },
                'attribute_map': {
                    'repository': 'repository',
                    'ref': 'ref',
                    'after': 'after',
                    'amount': 'amount',
                    'delimiter': 'delimiter',
                    'prefix': 'prefix',
                },
                'location_map': {
                    'repository': 'path',
                    'ref': 'path',
                    'after': 'query',
                    'amount': 'query',
                    'delimiter': 'query',
                    'prefix': 'query',
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
            api_client=api_client,
            callable=__list_objects
        )

        def __stage_object(
            self,
            repository,
            branch,
            path,
            object_stage_creation,
            **kwargs
        ):
            """stage an object\"s metadata for the given branch  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.stage_object(repository, branch, path, object_stage_creation, async_req=True)
            >>> result = thread.get()

            Args:
                repository (str):
                branch (str):
                path (str):
                object_stage_creation (ObjectStageCreation):

            Keyword Args:
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                ObjectStats
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
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['repository'] = \
                repository
            kwargs['branch'] = \
                branch
            kwargs['path'] = \
                path
            kwargs['object_stage_creation'] = \
                object_stage_creation
            return self.call_with_http_info(**kwargs)

        self.stage_object = _Endpoint(
            settings={
                'response_type': (ObjectStats,),
                'auth': [
                    'basic_auth',
                    'cookie_auth',
                    'jwt_token'
                ],
                'endpoint_path': '/repositories/{repository}/branches/{branch}/objects',
                'operation_id': 'stage_object',
                'http_method': 'PUT',
                'servers': None,
            },
            params_map={
                'all': [
                    'repository',
                    'branch',
                    'path',
                    'object_stage_creation',
                ],
                'required': [
                    'repository',
                    'branch',
                    'path',
                    'object_stage_creation',
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
                    'repository':
                        (str,),
                    'branch':
                        (str,),
                    'path':
                        (str,),
                    'object_stage_creation':
                        (ObjectStageCreation,),
                },
                'attribute_map': {
                    'repository': 'repository',
                    'branch': 'branch',
                    'path': 'path',
                },
                'location_map': {
                    'repository': 'path',
                    'branch': 'path',
                    'path': 'query',
                    'object_stage_creation': 'body',
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
            api_client=api_client,
            callable=__stage_object
        )

        def __stat_object(
            self,
            repository,
            ref,
            path,
            **kwargs
        ):
            """get object metadata  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.stat_object(repository, ref, path, async_req=True)
            >>> result = thread.get()

            Args:
                repository (str):
                ref (str): a reference (could be either a branch or a commit ID)
                path (str):

            Keyword Args:
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                ObjectStats
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
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['repository'] = \
                repository
            kwargs['ref'] = \
                ref
            kwargs['path'] = \
                path
            return self.call_with_http_info(**kwargs)

        self.stat_object = _Endpoint(
            settings={
                'response_type': (ObjectStats,),
                'auth': [
                    'basic_auth',
                    'cookie_auth',
                    'jwt_token'
                ],
                'endpoint_path': '/repositories/{repository}/refs/{ref}/objects/stat',
                'operation_id': 'stat_object',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'repository',
                    'ref',
                    'path',
                ],
                'required': [
                    'repository',
                    'ref',
                    'path',
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
                    'repository':
                        (str,),
                    'ref':
                        (str,),
                    'path':
                        (str,),
                },
                'attribute_map': {
                    'repository': 'repository',
                    'ref': 'ref',
                    'path': 'path',
                },
                'location_map': {
                    'repository': 'path',
                    'ref': 'path',
                    'path': 'query',
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
            api_client=api_client,
            callable=__stat_object
        )

        def __upload_object(
            self,
            repository,
            branch,
            path,
            **kwargs
        ):
            """upload_object  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.upload_object(repository, branch, path, async_req=True)
            >>> result = thread.get()

            Args:
                repository (str):
                branch (str):
                path (str):

            Keyword Args:
                storage_class (str): [optional]
                if_none_match (str): Currently supports only \"*\" to allow uploading an object only if one doesn't exist yet. [optional]
                content (file_type): Object content to upload. [optional]
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                ObjectStats
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
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['repository'] = \
                repository
            kwargs['branch'] = \
                branch
            kwargs['path'] = \
                path
            return self.call_with_http_info(**kwargs)

        self.upload_object = _Endpoint(
            settings={
                'response_type': (ObjectStats,),
                'auth': [
                    'basic_auth',
                    'cookie_auth',
                    'jwt_token'
                ],
                'endpoint_path': '/repositories/{repository}/branches/{branch}/objects',
                'operation_id': 'upload_object',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'repository',
                    'branch',
                    'path',
                    'storage_class',
                    'if_none_match',
                    'content',
                ],
                'required': [
                    'repository',
                    'branch',
                    'path',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'if_none_match',
                ]
            },
            root_map={
                'validations': {
                    ('if_none_match',): {

                        'regex': {
                            'pattern': r'^\*$',  # noqa: E501
                        },
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'repository':
                        (str,),
                    'branch':
                        (str,),
                    'path':
                        (str,),
                    'storage_class':
                        (str,),
                    'if_none_match':
                        (str,),
                    'content':
                        (file_type,),
                },
                'attribute_map': {
                    'repository': 'repository',
                    'branch': 'branch',
                    'path': 'path',
                    'storage_class': 'storageClass',
                    'if_none_match': 'If-None-Match',
                    'content': 'content',
                },
                'location_map': {
                    'repository': 'path',
                    'branch': 'path',
                    'path': 'query',
                    'storage_class': 'query',
                    'if_none_match': 'header',
                    'content': 'form',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [
                    'multipart/form-data'
                ]
            },
            api_client=api_client,
            callable=__upload_object
        )
