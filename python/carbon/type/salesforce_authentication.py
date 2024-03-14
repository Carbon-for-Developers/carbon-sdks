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


class RequiredSalesforceAuthentication(TypedDict):
    source: typing.Union[bool, date, datetime, dict, float, int, list, str, None]

    access_token: str

    domain: str

class OptionalSalesforceAuthentication(TypedDict, total=False):
    refresh_token: typing.Optional[str]

class SalesforceAuthentication(RequiredSalesforceAuthentication, OptionalSalesforceAuthentication):
    pass
