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

from carbon.type.lead import Lead

class RequiredLeadsResponse(TypedDict):
    count: int

    next_cursor: typing.Optional[str]

    data: typing.List[Lead]

class OptionalLeadsResponse(TypedDict, total=False):
    pass

class LeadsResponse(RequiredLeadsResponse, OptionalLeadsResponse):
    pass
