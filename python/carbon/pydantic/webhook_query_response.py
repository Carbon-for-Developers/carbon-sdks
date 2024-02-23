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
from pydantic import BaseModel, Field, RootModel

from carbon.pydantic.webhook_no_key import WebhookNoKey

class WebhookQueryResponse(BaseModel):
    results: typing.List[WebhookNoKey] = Field(alias='results')

    count: int = Field(alias='count')
    class Config:
        arbitrary_types_allowed = True
