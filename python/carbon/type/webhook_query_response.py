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

from carbon.type.webhook_no_key import WebhookNoKey

class RequiredWebhookQueryResponse(TypedDict):
    results: typing.List[WebhookNoKey]

    count: int

class OptionalWebhookQueryResponse(TypedDict, total=False):
    pass

class WebhookQueryResponse(RequiredWebhookQueryResponse, OptionalWebhookQueryResponse):
    pass
