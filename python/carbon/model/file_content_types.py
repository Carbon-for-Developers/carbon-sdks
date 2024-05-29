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


class FileContentTypes(
    schemas.EnumBase,
    schemas.StrSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        enum_value_to_name = {
            "TEXT": "TEXT",
            "IMAGE": "IMAGE",
            "AUDIO": "AUDIO",
            "VIDEO": "VIDEO",
        }
    
    @schemas.classproperty
    def TEXT(cls):
        return cls("TEXT")
    
    @schemas.classproperty
    def IMAGE(cls):
        return cls("IMAGE")
    
    @schemas.classproperty
    def AUDIO(cls):
        return cls("AUDIO")
    
    @schemas.classproperty
    def VIDEO(cls):
        return cls("VIDEO")
