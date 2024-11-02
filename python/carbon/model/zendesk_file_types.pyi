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


class ZendeskFileTypes(
    schemas.EnumBase,
    schemas.StrSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """
    
    @schemas.classproperty
    def TICKET(cls):
        return cls("TICKET")
    
    @schemas.classproperty
    def ARTICLE(cls):
        return cls("ARTICLE")
    
    @schemas.classproperty
    def TICKET_ATTACHMENT(cls):
        return cls("TICKET_ATTACHMENT")
    
    @schemas.classproperty
    def SECTION(cls):
        return cls("SECTION")
    
    @schemas.classproperty
    def CATEGORY(cls):
        return cls("CATEGORY")
