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


class FetchURLsResponse(
    schemas.DictSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        required = {
            "html_content",
            "urls",
        }
        
        class properties:
        
            @staticmethod
            def urls() -> typing.Type['FetchURLsResponseUrls']:
                return FetchURLsResponseUrls
            html_content = schemas.StrSchema
            __annotations__ = {
                "urls": urls,
                "html_content": html_content,
            }
    
    html_content: MetaOapg.properties.html_content
    urls: 'FetchURLsResponseUrls'
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["urls"]) -> 'FetchURLsResponseUrls': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["html_content"]) -> MetaOapg.properties.html_content: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["urls", "html_content", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["urls"]) -> 'FetchURLsResponseUrls': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["html_content"]) -> MetaOapg.properties.html_content: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["urls", "html_content", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        html_content: typing.Union[MetaOapg.properties.html_content, str, ],
        urls: 'FetchURLsResponseUrls',
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'FetchURLsResponse':
        return super().__new__(
            cls,
            *args,
            html_content=html_content,
            urls=urls,
            _configuration=_configuration,
            **kwargs,
        )

from carbon.model.fetch_urls_response_urls import FetchURLsResponseUrls
