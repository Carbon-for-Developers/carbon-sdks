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


class LeadsOrderBy(
    schemas.EnumBase,
    schemas.StrSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        enum_value_to_name = {
            "created_at": "CREATED_AT",
            "updated_at": "UPDATED_AT",
            "name": "NAME",
            "last_activity_at": "LAST_ACTIVITY_AT",
        }
    
    @schemas.classproperty
    def CREATED_AT(cls):
        return cls("created_at")
    
    @schemas.classproperty
    def UPDATED_AT(cls):
        return cls("updated_at")
    
    @schemas.classproperty
    def NAME(cls):
        return cls("name")
    
    @schemas.classproperty
    def LAST_ACTIVITY_AT(cls):
        return cls("last_activity_at")
