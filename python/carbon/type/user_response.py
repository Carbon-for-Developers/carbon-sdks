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

from carbon.type.user_response_unique_file_tags import UserResponseUniqueFileTags

class RequiredUserResponse(TypedDict):
    id: int

    organization_id: int

    organization_supplied_user_id: str

    created_at: datetime

    updated_at: datetime

    deleted_at: typing.Optional[datetime]

    num_files_synced: int

    num_characters_synced: int

    num_tokens_synced: int

    unique_file_tags: UserResponseUniqueFileTags

    enabled_features: typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]]

class OptionalUserResponse(TypedDict, total=False):
    pass

class UserResponse(RequiredUserResponse, OptionalUserResponse):
    pass
