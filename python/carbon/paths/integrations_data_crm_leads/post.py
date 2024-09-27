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

from carbon.model.base_includes import BaseIncludes as BaseIncludesSchema
from carbon.model.http_validation_error import HTTPValidationError as HTTPValidationErrorSchema
from carbon.model.leads_order_by_nullable import LeadsOrderByNullable as LeadsOrderByNullableSchema
from carbon.model.leads_response import LeadsResponse as LeadsResponseSchema
from carbon.model.order_dir_v2_nullable import OrderDirV2Nullable as OrderDirV2NullableSchema
from carbon.model.lead_filters import LeadFilters as LeadFiltersSchema
from carbon.model.leads_request import LeadsRequest as LeadsRequestSchema

from carbon.type.leads_request import LeadsRequest
from carbon.type.http_validation_error import HTTPValidationError
from carbon.type.leads_order_by_nullable import LeadsOrderByNullable
from carbon.type.order_dir_v2_nullable import OrderDirV2Nullable
from carbon.type.base_includes import BaseIncludes
from carbon.type.lead_filters import LeadFilters
from carbon.type.leads_response import LeadsResponse

from ...api_client import Dictionary
from carbon.pydantic.base_includes import BaseIncludes as BaseIncludesPydantic
from carbon.pydantic.order_dir_v2_nullable import OrderDirV2Nullable as OrderDirV2NullablePydantic
from carbon.pydantic.http_validation_error import HTTPValidationError as HTTPValidationErrorPydantic
from carbon.pydantic.leads_request import LeadsRequest as LeadsRequestPydantic
from carbon.pydantic.leads_response import LeadsResponse as LeadsResponsePydantic
from carbon.pydantic.lead_filters import LeadFilters as LeadFiltersPydantic
from carbon.pydantic.leads_order_by_nullable import LeadsOrderByNullable as LeadsOrderByNullablePydantic

from . import path

# body param
SchemaForRequestBodyApplicationJson = LeadsRequestSchema


request_body_leads_request = api_client.RequestBody(
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
SchemaFor200ResponseBodyApplicationJson = LeadsResponseSchema


@dataclass
class ApiResponseFor200(api_client.ApiResponse):
    body: LeadsResponse


@dataclass
class ApiResponseFor200Async(api_client.AsyncApiResponse):
    body: LeadsResponse


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

    def _get_leads_mapped_args(
        self,
        data_source_id: int,
        include_remote_data: typing.Optional[bool] = None,
        next_cursor: typing.Optional[typing.Optional[str]] = None,
        page_size: typing.Optional[typing.Optional[int]] = None,
        order_dir: typing.Optional[OrderDirV2Nullable] = None,
        includes: typing.Optional[typing.List[BaseIncludes]] = None,
        filters: typing.Optional[LeadFilters] = None,
        order_by: typing.Optional[LeadsOrderByNullable] = None,
    ) -> api_client.MappedArgs:
        args: api_client.MappedArgs = api_client.MappedArgs()
        _body = {}
        if data_source_id is not None:
            _body["data_source_id"] = data_source_id
        if include_remote_data is not None:
            _body["include_remote_data"] = include_remote_data
        if next_cursor is not None:
            _body["next_cursor"] = next_cursor
        if page_size is not None:
            _body["page_size"] = page_size
        if order_dir is not None:
            _body["order_dir"] = order_dir
        if includes is not None:
            _body["includes"] = includes
        if filters is not None:
            _body["filters"] = filters
        if order_by is not None:
            _body["order_by"] = order_by
        args.body = _body
        return args

    async def _aget_leads_oapg(
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
        Get Leads
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
            path_template='/integrations/data/crm/leads',
            body=body,
            auth_settings=_auth,
            headers=_headers,
        )
        serialized_data = request_body_leads_request.serialize(body, content_type)
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


    def _get_leads_oapg(
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
        Get Leads
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
            path_template='/integrations/data/crm/leads',
            body=body,
            auth_settings=_auth,
            headers=_headers,
        )
        serialized_data = request_body_leads_request.serialize(body, content_type)
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


class GetLeadsRaw(BaseApi):
    # this class is used by api classes that refer to endpoints with operationId fn names

    async def aget_leads(
        self,
        data_source_id: int,
        include_remote_data: typing.Optional[bool] = None,
        next_cursor: typing.Optional[typing.Optional[str]] = None,
        page_size: typing.Optional[typing.Optional[int]] = None,
        order_dir: typing.Optional[OrderDirV2Nullable] = None,
        includes: typing.Optional[typing.List[BaseIncludes]] = None,
        filters: typing.Optional[LeadFilters] = None,
        order_by: typing.Optional[LeadsOrderByNullable] = None,
        **kwargs,
    ) -> typing.Union[
        ApiResponseFor200Async,
        api_client.ApiResponseWithoutDeserializationAsync,
        AsyncGeneratorResponse,
    ]:
        args = self._get_leads_mapped_args(
            data_source_id=data_source_id,
            include_remote_data=include_remote_data,
            next_cursor=next_cursor,
            page_size=page_size,
            order_dir=order_dir,
            includes=includes,
            filters=filters,
            order_by=order_by,
        )
        return await self._aget_leads_oapg(
            body=args.body,
            **kwargs,
        )
    
    def get_leads(
        self,
        data_source_id: int,
        include_remote_data: typing.Optional[bool] = None,
        next_cursor: typing.Optional[typing.Optional[str]] = None,
        page_size: typing.Optional[typing.Optional[int]] = None,
        order_dir: typing.Optional[OrderDirV2Nullable] = None,
        includes: typing.Optional[typing.List[BaseIncludes]] = None,
        filters: typing.Optional[LeadFilters] = None,
        order_by: typing.Optional[LeadsOrderByNullable] = None,
    ) -> typing.Union[
        ApiResponseFor200,
        api_client.ApiResponseWithoutDeserialization,
    ]:
        """  """
        args = self._get_leads_mapped_args(
            data_source_id=data_source_id,
            include_remote_data=include_remote_data,
            next_cursor=next_cursor,
            page_size=page_size,
            order_dir=order_dir,
            includes=includes,
            filters=filters,
            order_by=order_by,
        )
        return self._get_leads_oapg(
            body=args.body,
        )

class GetLeads(BaseApi):

    async def aget_leads(
        self,
        data_source_id: int,
        include_remote_data: typing.Optional[bool] = None,
        next_cursor: typing.Optional[typing.Optional[str]] = None,
        page_size: typing.Optional[typing.Optional[int]] = None,
        order_dir: typing.Optional[OrderDirV2Nullable] = None,
        includes: typing.Optional[typing.List[BaseIncludes]] = None,
        filters: typing.Optional[LeadFilters] = None,
        order_by: typing.Optional[LeadsOrderByNullable] = None,
        validate: bool = False,
        **kwargs,
    ) -> LeadsResponsePydantic:
        raw_response = await self.raw.aget_leads(
            data_source_id=data_source_id,
            include_remote_data=include_remote_data,
            next_cursor=next_cursor,
            page_size=page_size,
            order_dir=order_dir,
            includes=includes,
            filters=filters,
            order_by=order_by,
            **kwargs,
        )
        if validate:
            return LeadsResponsePydantic(**raw_response.body)
        return api_client.construct_model_instance(LeadsResponsePydantic, raw_response.body)
    
    
    def get_leads(
        self,
        data_source_id: int,
        include_remote_data: typing.Optional[bool] = None,
        next_cursor: typing.Optional[typing.Optional[str]] = None,
        page_size: typing.Optional[typing.Optional[int]] = None,
        order_dir: typing.Optional[OrderDirV2Nullable] = None,
        includes: typing.Optional[typing.List[BaseIncludes]] = None,
        filters: typing.Optional[LeadFilters] = None,
        order_by: typing.Optional[LeadsOrderByNullable] = None,
        validate: bool = False,
    ) -> LeadsResponsePydantic:
        raw_response = self.raw.get_leads(
            data_source_id=data_source_id,
            include_remote_data=include_remote_data,
            next_cursor=next_cursor,
            page_size=page_size,
            order_dir=order_dir,
            includes=includes,
            filters=filters,
            order_by=order_by,
        )
        if validate:
            return LeadsResponsePydantic(**raw_response.body)
        return api_client.construct_model_instance(LeadsResponsePydantic, raw_response.body)


class ApiForpost(BaseApi):
    # this class is used by api classes that refer to endpoints by path and http method names

    async def apost(
        self,
        data_source_id: int,
        include_remote_data: typing.Optional[bool] = None,
        next_cursor: typing.Optional[typing.Optional[str]] = None,
        page_size: typing.Optional[typing.Optional[int]] = None,
        order_dir: typing.Optional[OrderDirV2Nullable] = None,
        includes: typing.Optional[typing.List[BaseIncludes]] = None,
        filters: typing.Optional[LeadFilters] = None,
        order_by: typing.Optional[LeadsOrderByNullable] = None,
        **kwargs,
    ) -> typing.Union[
        ApiResponseFor200Async,
        api_client.ApiResponseWithoutDeserializationAsync,
        AsyncGeneratorResponse,
    ]:
        args = self._get_leads_mapped_args(
            data_source_id=data_source_id,
            include_remote_data=include_remote_data,
            next_cursor=next_cursor,
            page_size=page_size,
            order_dir=order_dir,
            includes=includes,
            filters=filters,
            order_by=order_by,
        )
        return await self._aget_leads_oapg(
            body=args.body,
            **kwargs,
        )
    
    def post(
        self,
        data_source_id: int,
        include_remote_data: typing.Optional[bool] = None,
        next_cursor: typing.Optional[typing.Optional[str]] = None,
        page_size: typing.Optional[typing.Optional[int]] = None,
        order_dir: typing.Optional[OrderDirV2Nullable] = None,
        includes: typing.Optional[typing.List[BaseIncludes]] = None,
        filters: typing.Optional[LeadFilters] = None,
        order_by: typing.Optional[LeadsOrderByNullable] = None,
    ) -> typing.Union[
        ApiResponseFor200,
        api_client.ApiResponseWithoutDeserialization,
    ]:
        """  """
        args = self._get_leads_mapped_args(
            data_source_id=data_source_id,
            include_remote_data=include_remote_data,
            next_cursor=next_cursor,
            page_size=page_size,
            order_dir=order_dir,
            includes=includes,
            filters=filters,
            order_by=order_by,
        )
        return self._get_leads_oapg(
            body=args.body,
        )

