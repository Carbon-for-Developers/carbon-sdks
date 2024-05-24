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

from carbon.pydantic.fetch_urls_response_urls import FetchURLsResponseUrls

class FetchURLsResponse(BaseModel):
    urls: FetchURLsResponseUrls = Field(alias='urls')

    html_content: str = Field(alias='html_content')

    error_message: typing.Optional[str] = Field(alias='error_message')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
