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

from carbon.model.http_validation_error import HTTPValidationError as HTTPValidationErrorSchema

from carbon.type.http_validation_error import HTTPValidationError

from ...api_client import Dictionary
from carbon.pydantic.http_validation_error import HTTPValidationError as HTTPValidationErrorPydantic

# Query params
TypesSchema = schemas.StrSchema


class CursorSchema(
    schemas.StrBase,
    schemas.NoneBase,
    schemas.Schema,
    schemas.NoneStrMixin
):


    def __new__(
        cls,
        *args: typing.Union[None, str, ],
        _configuration: typing.Optional[schemas.Configuration] = None,
    ) -> 'CursorSchema':
        return super().__new__(
            cls,
            *args,
            _configuration=_configuration,
        )


class DataSourceIdSchema(
    schemas.IntBase,
    schemas.NoneBase,
    schemas.Schema,
    schemas.NoneDecimalMixin
):


    def __new__(
        cls,
        *args: typing.Union[None, decimal.Decimal, int, ],
        _configuration: typing.Optional[schemas.Configuration] = None,
    ) -> 'DataSourceIdSchema':
        return super().__new__(
            cls,
            *args,
            _configuration=_configuration,
        )
ExcludeArchivedSchema = schemas.BoolSchema
RequestRequiredQueryParams = typing_extensions.TypedDict(
    'RequestRequiredQueryParams',
    {
    }
)
RequestOptionalQueryParams = typing_extensions.TypedDict(
    'RequestOptionalQueryParams',
    {
        'types': typing.Union[TypesSchema, str, ],
        'cursor': typing.Union[CursorSchema, None, str, ],
        'data_source_id': typing.Union[DataSourceIdSchema, None, decimal.Decimal, int, ],
        'exclude_archived': typing.Union[ExcludeArchivedSchema, bool, ],
    },
    total=False
)


class RequestQueryParams(RequestRequiredQueryParams, RequestOptionalQueryParams):
    pass


request_query_types = api_client.QueryParameter(
    name="types",
    style=api_client.ParameterStyle.FORM,
    schema=TypesSchema,
    explode=True,
)
request_query_cursor = api_client.QueryParameter(
    name="cursor",
    style=api_client.ParameterStyle.FORM,
    schema=CursorSchema,
    explode=True,
)
request_query_data_source_id = api_client.QueryParameter(
    name="data_source_id",
    style=api_client.ParameterStyle.FORM,
    schema=DataSourceIdSchema,
    explode=True,
)
request_query_exclude_archived = api_client.QueryParameter(
    name="exclude_archived",
    style=api_client.ParameterStyle.FORM,
    schema=ExcludeArchivedSchema,
    explode=True,
)
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
_all_accept_content_types = (
    'application/json',
)


class BaseApi(api_client.Api):

    def _list_conversations_mapped_args(
        self,
        types: typing.Optional[str] = None,
        cursor: typing.Optional[typing.Optional[str]] = None,
        data_source_id: typing.Optional[typing.Optional[int]] = None,
        exclude_archived: typing.Optional[bool] = None,
    ) -> api_client.MappedArgs:
        args: api_client.MappedArgs = api_client.MappedArgs()
        _query_params = {}
        if types is not None:
            _query_params["types"] = types
        if cursor is not None:
            _query_params["cursor"] = cursor
        if data_source_id is not None:
            _query_params["data_source_id"] = data_source_id
        if exclude_archived is not None:
            _query_params["exclude_archived"] = exclude_archived
        args.query = _query_params
        return args

    async def _alist_conversations_oapg(
        self,
            query_params: typing.Optional[dict] = {},
        skip_deserialization: bool = True,
        timeout: typing.Optional[typing.Union[float, typing.Tuple]] = None,
        accept_content_types: typing.Tuple[str] = _all_accept_content_types,
        stream: bool = False,
        **kwargs,
    ) -> typing.Union[
        ApiResponseFor200Async,
        api_client.ApiResponseWithoutDeserializationAsync,
        AsyncGeneratorResponse,
    ]:
        """
        Slack List Conversations
        :param skip_deserialization: If true then api_response.response will be set but
            api_response.body and api_response.headers will not be deserialized into schema
            class instances
        """
        self._verify_typed_dict_inputs_oapg(RequestQueryParams, query_params)
        used_path = path.value
    
        prefix_separator_iterator = None
        for parameter in (
            request_query_types,
            request_query_cursor,
            request_query_data_source_id,
            request_query_exclude_archived,
        ):
            parameter_data = query_params.get(parameter.name, schemas.unset)
            if parameter_data is schemas.unset:
                continue
            if prefix_separator_iterator is None:
                prefix_separator_iterator = parameter.get_prefix_separator_iterator()
            serialized_data = parameter.serialize(parameter_data, prefix_separator_iterator)
            for serialized_value in serialized_data.values():
                used_path += serialized_value
    
        _headers = HTTPHeaderDict()
        # TODO add cookie handling
        if accept_content_types:
            for accept_content_type in accept_content_types:
                _headers.add('Accept', accept_content_type)
        method = 'get'.upper()
        request_before_hook(
            resource_path=used_path,
            method=method,
            configuration=self.api_client.configuration,
            path_template='/integrations/slack/conversations',
            auth_settings=_auth,
            headers=_headers,
        )
    
        response = await self.api_client.async_call_api(
            resource_path=used_path,
            method=method,
            headers=_headers,
            auth_settings=_auth,
            prefix_separator_iterator=prefix_separator_iterator,
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


    def _list_conversations_oapg(
        self,
            query_params: typing.Optional[dict] = {},
        skip_deserialization: bool = True,
        timeout: typing.Optional[typing.Union[float, typing.Tuple]] = None,
        accept_content_types: typing.Tuple[str] = _all_accept_content_types,
        stream: bool = False,
    ) -> typing.Union[
        ApiResponseFor200,
        api_client.ApiResponseWithoutDeserialization,
    ]:
        """
        Slack List Conversations
        :param skip_deserialization: If true then api_response.response will be set but
            api_response.body and api_response.headers will not be deserialized into schema
            class instances
        """
        self._verify_typed_dict_inputs_oapg(RequestQueryParams, query_params)
        used_path = path.value
    
        prefix_separator_iterator = None
        for parameter in (
            request_query_types,
            request_query_cursor,
            request_query_data_source_id,
            request_query_exclude_archived,
        ):
            parameter_data = query_params.get(parameter.name, schemas.unset)
            if parameter_data is schemas.unset:
                continue
            if prefix_separator_iterator is None:
                prefix_separator_iterator = parameter.get_prefix_separator_iterator()
            serialized_data = parameter.serialize(parameter_data, prefix_separator_iterator)
            for serialized_value in serialized_data.values():
                used_path += serialized_value
    
        _headers = HTTPHeaderDict()
        # TODO add cookie handling
        if accept_content_types:
            for accept_content_type in accept_content_types:
                _headers.add('Accept', accept_content_type)
        method = 'get'.upper()
        request_before_hook(
            resource_path=used_path,
            method=method,
            configuration=self.api_client.configuration,
            path_template='/integrations/slack/conversations',
            auth_settings=_auth,
            headers=_headers,
        )
    
        response = self.api_client.call_api(
            resource_path=used_path,
            method=method,
            headers=_headers,
            auth_settings=_auth,
            prefix_separator_iterator=prefix_separator_iterator,
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


class ListConversationsRaw(BaseApi):
    # this class is used by api classes that refer to endpoints with operationId fn names

    async def alist_conversations(
        self,
        types: typing.Optional[str] = None,
        cursor: typing.Optional[typing.Optional[str]] = None,
        data_source_id: typing.Optional[typing.Optional[int]] = None,
        exclude_archived: typing.Optional[bool] = None,
        **kwargs,
    ) -> typing.Union[
        ApiResponseFor200Async,
        api_client.ApiResponseWithoutDeserializationAsync,
        AsyncGeneratorResponse,
    ]:
        args = self._list_conversations_mapped_args(
            types=types,
            cursor=cursor,
            data_source_id=data_source_id,
            exclude_archived=exclude_archived,
        )
        return await self._alist_conversations_oapg(
            query_params=args.query,
            **kwargs,
        )
    
    def list_conversations(
        self,
        types: typing.Optional[str] = None,
        cursor: typing.Optional[typing.Optional[str]] = None,
        data_source_id: typing.Optional[typing.Optional[int]] = None,
        exclude_archived: typing.Optional[bool] = None,
    ) -> typing.Union[
        ApiResponseFor200,
        api_client.ApiResponseWithoutDeserialization,
    ]:
        """ List all of your public and private channels, DMs, and Group DMs. The ID from response  can be used as a filter to sync messages to Carbon    types: Comma separated list of types. Available types are im (DMs), mpim (group DMs), public_channel, and private_channel. Defaults to public_channel.     cursor: Used for pagination. If next_cursor is returned in response, you need to pass it as the cursor in the next request     data_source_id: Data source needs to be specified if you have linked multiple slack accounts   exclude_archived: Should archived conversations be excluded, defaults to true """
        args = self._list_conversations_mapped_args(
            types=types,
            cursor=cursor,
            data_source_id=data_source_id,
            exclude_archived=exclude_archived,
        )
        return self._list_conversations_oapg(
            query_params=args.query,
        )

class ListConversations(BaseApi):

    async def alist_conversations(
        self,
        types: typing.Optional[str] = None,
        cursor: typing.Optional[typing.Optional[str]] = None,
        data_source_id: typing.Optional[typing.Optional[int]] = None,
        exclude_archived: typing.Optional[bool] = None,
        validate: bool = False,
        **kwargs,
    ) -> Dictionary:
        raw_response = await self.raw.alist_conversations(
            types=types,
            cursor=cursor,
            data_source_id=data_source_id,
            exclude_archived=exclude_archived,
            **kwargs,
        )
        if validate:
            return Dictionary(**raw_response.body)
        return api_client.construct_model_instance(Dictionary, raw_response.body)
    
    
    def list_conversations(
        self,
        types: typing.Optional[str] = None,
        cursor: typing.Optional[typing.Optional[str]] = None,
        data_source_id: typing.Optional[typing.Optional[int]] = None,
        exclude_archived: typing.Optional[bool] = None,
        validate: bool = False,
    ) -> Dictionary:
        raw_response = self.raw.list_conversations(
            types=types,
            cursor=cursor,
            data_source_id=data_source_id,
            exclude_archived=exclude_archived,
        )
        if validate:
            return Dictionary(**raw_response.body)
        return api_client.construct_model_instance(Dictionary, raw_response.body)


class ApiForget(BaseApi):
    # this class is used by api classes that refer to endpoints by path and http method names

    async def aget(
        self,
        types: typing.Optional[str] = None,
        cursor: typing.Optional[typing.Optional[str]] = None,
        data_source_id: typing.Optional[typing.Optional[int]] = None,
        exclude_archived: typing.Optional[bool] = None,
        **kwargs,
    ) -> typing.Union[
        ApiResponseFor200Async,
        api_client.ApiResponseWithoutDeserializationAsync,
        AsyncGeneratorResponse,
    ]:
        args = self._list_conversations_mapped_args(
            types=types,
            cursor=cursor,
            data_source_id=data_source_id,
            exclude_archived=exclude_archived,
        )
        return await self._alist_conversations_oapg(
            query_params=args.query,
            **kwargs,
        )
    
    def get(
        self,
        types: typing.Optional[str] = None,
        cursor: typing.Optional[typing.Optional[str]] = None,
        data_source_id: typing.Optional[typing.Optional[int]] = None,
        exclude_archived: typing.Optional[bool] = None,
    ) -> typing.Union[
        ApiResponseFor200,
        api_client.ApiResponseWithoutDeserialization,
    ]:
        """ List all of your public and private channels, DMs, and Group DMs. The ID from response  can be used as a filter to sync messages to Carbon    types: Comma separated list of types. Available types are im (DMs), mpim (group DMs), public_channel, and private_channel. Defaults to public_channel.     cursor: Used for pagination. If next_cursor is returned in response, you need to pass it as the cursor in the next request     data_source_id: Data source needs to be specified if you have linked multiple slack accounts   exclude_archived: Should archived conversations be excluded, defaults to true """
        args = self._list_conversations_mapped_args(
            types=types,
            cursor=cursor,
            data_source_id=data_source_id,
            exclude_archived=exclude_archived,
        )
        return self._list_conversations_oapg(
            query_params=args.query,
        )

