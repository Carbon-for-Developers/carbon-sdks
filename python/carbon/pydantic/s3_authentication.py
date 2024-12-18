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


class S3Authentication(BaseModel):
    source: Literal["S3"] = Field(alias='source')

    access_key: str = Field(alias='access_key')

    access_key_secret: str = Field(alias='access_key_secret')

    # You can specify a Digital Ocean endpoint URL to connect a Digital Ocean Space through this endpoint.         The URL should be of format <region>.digitaloceanspaces.com. It's not required for S3 buckets.
    endpoint_url: typing.Optional[typing.Optional[str]] = Field(None, alias='endpoint_url')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
