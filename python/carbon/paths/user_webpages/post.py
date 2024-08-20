# coding: utf-8

"""
    Carbon

    Connect external data to LLMs, no matter the source.

    The version of the OpenAPI document: 1.0.0
    Generated by: https://konfigthis.com
"""

from dataclasses import dataclass
import typing_extensions
import urllib3
from pydantic import RootModel
from carbon.request_before_hook import request_before_hook
import json
from urllib3._collections import HTTPHeaderDict

from carbon.api_response import AsyncGeneratorResponse
from carbon import api_client, exceptions
from datetime import date, datetime  # noqa: F401
import decimal  # noqa: F401
import functools  # noqa: F401
import io  # noqa: F401
import re  # noqa: F401
import typing  # noqa: F401
import typing_extensions  # noqa: F401
import uuid  # noqa: F401

import frozendict  # noqa: F401

from carbon import schemas  # noqa: F401

from carbon.model.pagination import Pagination as PaginationSchema
from carbon.model.http_validation_error import HTTPValidationError as HTTPValidationErrorSchema
from carbon.model.user_web_pages_filters import UserWebPagesFilters as UserWebPagesFiltersSchema
from carbon.model.order_dir_v2 import OrderDirV2 as OrderDirV2Schema
from carbon.model.user_web_pages_request import UserWebPagesRequest as UserWebPagesRequestSchema
from carbon.model.user_web_page_order_by_types import UserWebPageOrderByTypes as UserWebPageOrderByTypesSchema

from carbon.type.user_web_pages_request import UserWebPagesRequest
from carbon.type.http_validation_error import HTTPValidationError
from carbon.type.user_web_page_order_by_types import UserWebPageOrderByTypes
from carbon.type.order_dir_v2 import OrderDirV2
from carbon.type.pagination import Pagination
from carbon.type.user_web_pages_filters import UserWebPagesFilters

from ...api_client import Dictionary
from carbon.pydantic.user_web_pages_filters import UserWebPagesFilters as UserWebPagesFiltersPydantic
from carbon.pydantic.order_dir_v2 import OrderDirV2 as OrderDirV2Pydantic
from carbon.pydantic.user_web_page_order_by_types import UserWebPageOrderByTypes as UserWebPageOrderByTypesPydantic
from carbon.pydantic.http_validation_error import HTTPValidationError as HTTPValidationErrorPydantic
from carbon.pydantic.pagination import Pagination as PaginationPydantic
from carbon.pydantic.user_web_pages_request import UserWebPagesRequest as UserWebPagesRequestPydantic

from . import path

# body param
SchemaForRequestBodyApplicationJson = UserWebPagesRequestSchema


request_body_user_web_pages_request = api_client.RequestBody(
    content={
        'application/json': api_client.MediaType(
            schema=SchemaForRequestBodyApplicationJson),
    },
    required=True,
)
_auth = [
    'accessToken',
    'apiKey',
    'customerId',
]
SchemaFor200ResponseBodyApplicationJson = schemas.DictSchema


@dataclass
class ApiResponseFor200(api_client.ApiResponse):
    body: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]


@dataclass
class ApiResponseFor200Async(api_client.AsyncApiResponse):
    body: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]


_response_for_200 = api_client.OpenApiResponse(
    response_cls=ApiResponseFor200,
    response_cls_async=ApiResponseFor200Async,
    content={
        'application/json': api_client.MediaType(
            schema=SchemaFor200ResponseBodyApplicationJson),
    },
)
SchemaFor422ResponseBodyApplicationJson = HTTPValidationErrorSchema


@dataclass
class ApiResponseFor422(api_client.ApiResponse):
    body: HTTPValidationError


@dataclass
class ApiResponseFor422Async(api_client.AsyncApiResponse):
    body: HTTPValidationError


_response_for_422 = api_client.OpenApiResponse(
    response_cls=ApiResponseFor422,
    response_cls_async=ApiResponseFor422Async,
    content={
        'application/json': api_client.MediaType(
            schema=SchemaFor422ResponseBodyApplicationJson),
    },
)
_status_code_to_response = {
    '200': _response_for_200,
    '422': _response_for_422,
}
_all_accept_content_types = (
    'application/json',
)


class BaseApi(api_client.Api):

    def _user_webpages_mapped_args(
        self,
        filters: typing.Optional[UserWebPagesFilters] = None,
        pagination: typing.Optional[Pagination] = None,
        order_by: typing.Optional[UserWebPageOrderByTypes] = None,
        order_dir: typing.Optional[OrderDirV2] = None,
    ) -> api_client.MappedArgs:
        args: api_client.MappedArgs = api_client.MappedArgs()
        _body = {}
        if filters is not None:
            _body["filters"] = filters
        if pagination is not None:
            _body["pagination"] = pagination
        if order_by is not None:
            _body["order_by"] = order_by
        if order_dir is not None:
            _body["order_dir"] = order_dir
        args.body = _body
        return args

    async def _auser_webpages_oapg(
        self,
        body: typing.Any = None,
        skip_deserialization: bool = True,
        timeout: typing.Optional[typing.Union[float, typing.Tuple]] = None,
        accept_content_types: typing.Tuple[str] = _all_accept_content_types,
        content_type: str = 'application/json',
        stream: bool = False,
        **kwargs,
    ) -> typing.Union[
        ApiResponseFor200Async,
        api_client.ApiResponseWithoutDeserializationAsync,
        AsyncGeneratorResponse,
    ]:
        """
        User Web Pages
        :param skip_deserialization: If true then api_response.response will be set but
            api_response.body and api_response.headers will not be deserialized into schema
            class instances
        """
        used_path = path.value
    
        _headers = HTTPHeaderDict()
        # TODO add cookie handling
        if accept_content_types:
            for accept_content_type in accept_content_types:
                _headers.add('Accept', accept_content_type)
        method = 'post'.upper()
        _headers.add('Content-Type', content_type)
    
        if body is schemas.unset:
            raise exceptions.ApiValueError(
                'The required body parameter has an invalid value of: unset. Set a valid value instead')
        _fields = None
        _body = None
        request_before_hook(
            resource_path=used_path,
            method=method,
            configuration=self.api_client.configuration,
            path_template='/user_webpages',
            body=body,
            auth_settings=_auth,
            headers=_headers,
        )
        serialized_data = request_body_user_web_pages_request.serialize(body, content_type)
        if 'fields' in serialized_data:
            _fields = serialized_data['fields']
        elif 'body' in serialized_data:
            _body = serialized_data['body']
    
        response = await self.api_client.async_call_api(
            resource_path=used_path,
            method=method,
            headers=_headers,
            fields=_fields,
            serialized_body=_body,
            body=body,
            auth_settings=_auth,
            timeout=timeout,
            **kwargs
        )
    
        if stream:
            if not 200 <= response.http_response.status <= 299:
                body = (await response.http_response.content.read()).decode("utf-8")
                raise exceptions.ApiStreamingException(
                    status=response.http_response.status,
                    reason=response.http_response.reason,
                    body=body,
                )
    
            async def stream_iterator():
                """
                iterates over response.http_response.content and closes connection once iteration has finished
                """
                async for line in response.http_response.content:
                    if line == b'\r\n':
                        continue
                    yield line
                response.http_response.close()
                await response.session.close()
            return AsyncGeneratorResponse(
                content=stream_iterator(),
                headers=response.http_response.headers,
                status=response.http_response.status,
                response=response.http_response
            )
    
        response_for_status = _status_code_to_response.get(str(response.http_response.status))
        if response_for_status:
            api_response = await response_for_status.deserialize_async(
                                                    response,
                                                    self.api_client.configuration,
                                                    skip_deserialization=skip_deserialization
                                                )
        else:
            # If response data is JSON then deserialize for SDK consumer convenience
            is_json = api_client.JSONDetector._content_type_is_json(response.http_response.headers.get('Content-Type', ''))
            api_response = api_client.ApiResponseWithoutDeserializationAsync(
                body=await response.http_response.json() if is_json else await response.http_response.text(),
                response=response.http_response,
                round_trip_time=response.round_trip_time,
                status=response.http_response.status,
                headers=response.http_response.headers,
            )
    
        if not 200 <= api_response.status <= 299:
            raise exceptions.ApiException(api_response=api_response)
    
        # cleanup session / response
        response.http_response.close()
        await response.session.close()
    
        return api_response


    def _user_webpages_oapg(
        self,
        body: typing.Any = None,
        skip_deserialization: bool = True,
        timeout: typing.Optional[typing.Union[float, typing.Tuple]] = None,
        accept_content_types: typing.Tuple[str] = _all_accept_content_types,
        content_type: str = 'application/json',
        stream: bool = False,
    ) -> typing.Union[
        ApiResponseFor200,
        api_client.ApiResponseWithoutDeserialization,
    ]:
        """
        User Web Pages
        :param skip_deserialization: If true then api_response.response will be set but
            api_response.body and api_response.headers will not be deserialized into schema
            class instances
        """
        used_path = path.value
    
        _headers = HTTPHeaderDict()
        # TODO add cookie handling
        if accept_content_types:
            for accept_content_type in accept_content_types:
                _headers.add('Accept', accept_content_type)
        method = 'post'.upper()
        _headers.add('Content-Type', content_type)
    
        if body is schemas.unset:
            raise exceptions.ApiValueError(
                'The required body parameter has an invalid value of: unset. Set a valid value instead')
        _fields = None
        _body = None
        request_before_hook(
            resource_path=used_path,
            method=method,
            configuration=self.api_client.configuration,
            path_template='/user_webpages',
            body=body,
            auth_settings=_auth,
            headers=_headers,
        )
        serialized_data = request_body_user_web_pages_request.serialize(body, content_type)
        if 'fields' in serialized_data:
            _fields = serialized_data['fields']
        elif 'body' in serialized_data:
            _body = serialized_data['body']
    
        response = self.api_client.call_api(
            resource_path=used_path,
            method=method,
            headers=_headers,
            fields=_fields,
            serialized_body=_body,
            body=body,
            auth_settings=_auth,
            timeout=timeout,
        )
    
        response_for_status = _status_code_to_response.get(str(response.http_response.status))
        if response_for_status:
            api_response = response_for_status.deserialize(
                                                    response,
                                                    self.api_client.configuration,
                                                    skip_deserialization=skip_deserialization
                                                )
        else:
            # If response data is JSON then deserialize for SDK consumer convenience
            is_json = api_client.JSONDetector._content_type_is_json(response.http_response.headers.get('Content-Type', ''))
            api_response = api_client.ApiResponseWithoutDeserialization(
                body=json.loads(response.http_response.data) if is_json else response.http_response.data,
                response=response.http_response,
                round_trip_time=response.round_trip_time,
                status=response.http_response.status,
                headers=response.http_response.headers,
            )
    
        if not 200 <= api_response.status <= 299:
            raise exceptions.ApiException(api_response=api_response)
    
        return api_response


class UserWebpagesRaw(BaseApi):
    # this class is used by api classes that refer to endpoints with operationId fn names

    async def auser_webpages(
        self,
        filters: typing.Optional[UserWebPagesFilters] = None,
        pagination: typing.Optional[Pagination] = None,
        order_by: typing.Optional[UserWebPageOrderByTypes] = None,
        order_dir: typing.Optional[OrderDirV2] = None,
        **kwargs,
    ) -> typing.Union[
        ApiResponseFor200Async,
        api_client.ApiResponseWithoutDeserializationAsync,
        AsyncGeneratorResponse,
    ]:
        args = self._user_webpages_mapped_args(
            filters=filters,
            pagination=pagination,
            order_by=order_by,
            order_dir=order_dir,
        )
        return await self._auser_webpages_oapg(
            body=args.body,
            **kwargs,
        )
    
    def user_webpages(
        self,
        filters: typing.Optional[UserWebPagesFilters] = None,
        pagination: typing.Optional[Pagination] = None,
        order_by: typing.Optional[UserWebPageOrderByTypes] = None,
        order_dir: typing.Optional[OrderDirV2] = None,
    ) -> typing.Union[
        ApiResponseFor200,
        api_client.ApiResponseWithoutDeserialization,
    ]:
        """  """
        args = self._user_webpages_mapped_args(
            filters=filters,
            pagination=pagination,
            order_by=order_by,
            order_dir=order_dir,
        )
        return self._user_webpages_oapg(
            body=args.body,
        )

class UserWebpages(BaseApi):

    async def auser_webpages(
        self,
        filters: typing.Optional[UserWebPagesFilters] = None,
        pagination: typing.Optional[Pagination] = None,
        order_by: typing.Optional[UserWebPageOrderByTypes] = None,
        order_dir: typing.Optional[OrderDirV2] = None,
        validate: bool = False,
        **kwargs,
    ) -> Dictionary:
        raw_response = await self.raw.auser_webpages(
            filters=filters,
            pagination=pagination,
            order_by=order_by,
            order_dir=order_dir,
            **kwargs,
        )
        if validate:
            return Dictionary(**raw_response.body)
        return api_client.construct_model_instance(Dictionary, raw_response.body)
    
    
    def user_webpages(
        self,
        filters: typing.Optional[UserWebPagesFilters] = None,
        pagination: typing.Optional[Pagination] = None,
        order_by: typing.Optional[UserWebPageOrderByTypes] = None,
        order_dir: typing.Optional[OrderDirV2] = None,
        validate: bool = False,
    ) -> Dictionary:
        raw_response = self.raw.user_webpages(
            filters=filters,
            pagination=pagination,
            order_by=order_by,
            order_dir=order_dir,
        )
        if validate:
            return Dictionary(**raw_response.body)
        return api_client.construct_model_instance(Dictionary, raw_response.body)


class ApiForpost(BaseApi):
    # this class is used by api classes that refer to endpoints by path and http method names

    async def apost(
        self,
        filters: typing.Optional[UserWebPagesFilters] = None,
        pagination: typing.Optional[Pagination] = None,
        order_by: typing.Optional[UserWebPageOrderByTypes] = None,
        order_dir: typing.Optional[OrderDirV2] = None,
        **kwargs,
    ) -> typing.Union[
        ApiResponseFor200Async,
        api_client.ApiResponseWithoutDeserializationAsync,
        AsyncGeneratorResponse,
    ]:
        args = self._user_webpages_mapped_args(
            filters=filters,
            pagination=pagination,
            order_by=order_by,
            order_dir=order_dir,
        )
        return await self._auser_webpages_oapg(
            body=args.body,
            **kwargs,
        )
    
    def post(
        self,
        filters: typing.Optional[UserWebPagesFilters] = None,
        pagination: typing.Optional[Pagination] = None,
        order_by: typing.Optional[UserWebPageOrderByTypes] = None,
        order_dir: typing.Optional[OrderDirV2] = None,
    ) -> typing.Union[
        ApiResponseFor200,
        api_client.ApiResponseWithoutDeserialization,
    ]:
        """  """
        args = self._user_webpages_mapped_args(
            filters=filters,
            pagination=pagination,
            order_by=order_by,
            order_dir=order_dir,
        )
        return self._user_webpages_oapg(
            body=args.body,
        )

