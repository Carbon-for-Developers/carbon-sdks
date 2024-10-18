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


class OneDriveSharepointCredentials(BaseModel):
    client_id: str = Field(alias='client_id')

    redirect_uri: str = Field(alias='redirect_uri')

    client_secret: typing.Optional[typing.Optional[str]] = Field(None, alias='client_secret')

    file_picker_client_id: typing.Optional[typing.Optional[str]] = Field(None, alias='file_picker_client_id')

    file_picker_redirect_uri: typing.Optional[typing.Optional[str]] = Field(None, alias='file_picker_redirect_uri')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
