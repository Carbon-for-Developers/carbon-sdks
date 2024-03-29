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

from carbon.pydantic.user_response_unique_file_tags import UserResponseUniqueFileTags

class UserResponse(BaseModel):
    id: int = Field(alias='id')

    organization_id: int = Field(alias='organization_id')

    organization_supplied_user_id: str = Field(alias='organization_supplied_user_id')

    created_at: datetime = Field(alias='created_at')

    updated_at: datetime = Field(alias='updated_at')

    deleted_at: typing.Optional[datetime] = Field(alias='deleted_at')

    num_files_synced: int = Field(alias='num_files_synced')

    num_characters_synced: int = Field(alias='num_characters_synced')

    num_tokens_synced: int = Field(alias='num_tokens_synced')

    unique_file_tags: UserResponseUniqueFileTags = Field(alias='unique_file_tags')

    enabled_features: typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]] = Field(alias='enabled_features')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
