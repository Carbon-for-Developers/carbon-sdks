# coding: utf-8

"""
    Carbon

    Connect external data to LLMs, no matter the source.

    The version of the OpenAPI document: 1.0.0
    Generated by: https://konfigthis.com
"""

from datetime import datetime, date
import typing
from enum import Enum
from typing_extensions import TypedDict, Literal, TYPE_CHECKING


class RequiredS3AuthRequest(TypedDict):
    access_key: str

    access_key_secret: str

class OptionalS3AuthRequest(TypedDict, total=False):
    pass

class S3AuthRequest(RequiredS3AuthRequest, OptionalS3AuthRequest):
    pass
