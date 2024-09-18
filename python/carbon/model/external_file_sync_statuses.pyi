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


class ExternalFileSyncStatuses(
    schemas.EnumBase,
    schemas.StrSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """
    
    @schemas.classproperty
    def DELAYED(cls):
        return cls("DELAYED")
    
    @schemas.classproperty
    def QUEUED_FOR_SYNC(cls):
        return cls("QUEUED_FOR_SYNC")
    
    @schemas.classproperty
    def SYNCING(cls):
        return cls("SYNCING")
    
    @schemas.classproperty
    def READY(cls):
        return cls("READY")
    
    @schemas.classproperty
    def SYNC_ERROR(cls):
        return cls("SYNC_ERROR")
    
    @schemas.classproperty
    def EVALUATING_RESYNC(cls):
        return cls("EVALUATING_RESYNC")
    
    @schemas.classproperty
    def RATE_LIMITED(cls):
        return cls("RATE_LIMITED")
    
    @schemas.classproperty
    def SYNC_ABORTED(cls):
        return cls("SYNC_ABORTED")
    
    @schemas.classproperty
    def QUEUED_FOR_OCR(cls):
        return cls("QUEUED_FOR_OCR")
    
    @schemas.classproperty
    def READY_TO_SYNC(cls):
        return cls("READY_TO_SYNC")
