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


class TokenResponse(
    schemas.DictSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        required = {
            "access_token",
            "refresh_token",
        }
        
        class properties:
            access_token = schemas.StrSchema
            refresh_token = schemas.StrSchema
            __annotations__ = {
                "access_token": access_token,
                "refresh_token": refresh_token,
            }
    
    access_token: MetaOapg.properties.access_token
    refresh_token: MetaOapg.properties.refresh_token
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["access_token"]) -> MetaOapg.properties.access_token: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["refresh_token"]) -> MetaOapg.properties.refresh_token: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["access_token", "refresh_token", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["access_token"]) -> MetaOapg.properties.access_token: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["refresh_token"]) -> MetaOapg.properties.refresh_token: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["access_token", "refresh_token", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        access_token: typing.Union[MetaOapg.properties.access_token, str, ],
        refresh_token: typing.Union[MetaOapg.properties.refresh_token, str, ],
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'TokenResponse':
        return super().__new__(
            cls,
            *args,
            access_token=access_token,
            refresh_token=refresh_token,
            _configuration=_configuration,
            **kwargs,
        )
