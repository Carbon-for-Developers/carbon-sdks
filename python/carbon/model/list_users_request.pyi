# coding: utf-8

"""
    Carbon

    Connect external data to LLMs, no matter the source.

    The version of the OpenAPI document: 1.0.0
    Generated by: https://konfigthis.com
"""

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


class ListUsersRequest(
    schemas.DictSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        
        class properties:
        
            @staticmethod
            def pagination() -> typing.Type['Pagination']:
                return Pagination
        
            @staticmethod
            def filters() -> typing.Type['ListUsersFilters']:
                return ListUsersFilters
        
            @staticmethod
            def order_by() -> typing.Type['ListUsersOrderByTypes']:
                return ListUsersOrderByTypes
        
            @staticmethod
            def order_dir() -> typing.Type['OrderDirV2']:
                return OrderDirV2
            include_count = schemas.BoolSchema
            __annotations__ = {
                "pagination": pagination,
                "filters": filters,
                "order_by": order_by,
                "order_dir": order_dir,
                "include_count": include_count,
            }
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["pagination"]) -> 'Pagination': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["filters"]) -> 'ListUsersFilters': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["order_by"]) -> 'ListUsersOrderByTypes': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["order_dir"]) -> 'OrderDirV2': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["include_count"]) -> MetaOapg.properties.include_count: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["pagination", "filters", "order_by", "order_dir", "include_count", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["pagination"]) -> typing.Union['Pagination', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["filters"]) -> typing.Union['ListUsersFilters', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["order_by"]) -> typing.Union['ListUsersOrderByTypes', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["order_dir"]) -> typing.Union['OrderDirV2', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["include_count"]) -> typing.Union[MetaOapg.properties.include_count, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["pagination", "filters", "order_by", "order_dir", "include_count", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        pagination: typing.Union['Pagination', schemas.Unset] = schemas.unset,
        filters: typing.Union['ListUsersFilters', schemas.Unset] = schemas.unset,
        order_by: typing.Union['ListUsersOrderByTypes', schemas.Unset] = schemas.unset,
        order_dir: typing.Union['OrderDirV2', schemas.Unset] = schemas.unset,
        include_count: typing.Union[MetaOapg.properties.include_count, bool, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'ListUsersRequest':
        return super().__new__(
            cls,
            *args,
            pagination=pagination,
            filters=filters,
            order_by=order_by,
            order_dir=order_dir,
            include_count=include_count,
            _configuration=_configuration,
            **kwargs,
        )

from carbon.model.list_users_filters import ListUsersFilters
from carbon.model.list_users_order_by_types import ListUsersOrderByTypes
from carbon.model.order_dir_v2 import OrderDirV2
from carbon.model.pagination import Pagination