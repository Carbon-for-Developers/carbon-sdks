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


class ZoteroAuthentication(BaseModel):
    source: typing.Union[bool, date, datetime, dict, float, int, list, str, None] = Field(alias='source')

    access_token: str = Field(alias='access_token')

    access_token_secret: str = Field(alias='access_token_secret')

    username: str = Field(alias='username')

    zotero_id: str = Field(alias='zotero_id')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
