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


class RequiredLeadFilters(TypedDict):
    pass

class OptionalLeadFilters(TypedDict, total=False):
    owner_id: typing.Optional[str]

    name: typing.Optional[str]

    converted_account_id: typing.Optional[str]

    converted_contact_id: typing.Optional[str]

class LeadFilters(RequiredLeadFilters, OptionalLeadFilters):
    pass
