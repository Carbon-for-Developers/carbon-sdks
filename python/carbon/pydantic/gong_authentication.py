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
from pydantic import BaseModel, Field, RootModel, ConfigDict


class GongAuthentication(BaseModel):
    source: Literal["GONG"] = Field(alias='source')

    access_token: str = Field(alias='access_token')

    gong_account_email: str = Field(alias='gong_account_email')

    refresh_token: typing.Optional[typing.Optional[str]] = Field(None, alias='refresh_token')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
