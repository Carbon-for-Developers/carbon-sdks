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

from carbon.pydantic.simple_o_auth_data_sources import SimpleOAuthDataSources

class OAuthAuthentication(BaseModel):
    source: SimpleOAuthDataSources = Field(alias='source')

    access_token: str = Field(alias='access_token')

    refresh_token: typing.Optional[typing.Optional[str]] = Field(None, alias='refresh_token')
    class Config:
        arbitrary_types_allowed = True
