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


class RequiredWebhookNoKey(TypedDict):
    id: int

    organization_id: int

    url: str

    created_at: datetime

    updated_at: datetime

class OptionalWebhookNoKey(TypedDict, total=False):
    pass

class WebhookNoKey(RequiredWebhookNoKey, OptionalWebhookNoKey):
    pass
