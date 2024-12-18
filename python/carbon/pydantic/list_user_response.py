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

from carbon.pydantic.list_user_response_auto_sync_enabled_sources import ListUserResponseAutoSyncEnabledSources

class ListUserResponse(BaseModel):
    id: int = Field(alias='id')

    organization_id: int = Field(alias='organization_id')

    organization_supplied_user_id: str = Field(alias='organization_supplied_user_id')

    created_at: datetime = Field(alias='created_at')

    updated_at: datetime = Field(alias='updated_at')

    deleted_at: typing.Optional[datetime] = Field(alias='deleted_at')

    enabled_features: typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]] = Field(alias='enabled_features')

    custom_limits: typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]] = Field(None, alias='custom_limits')

    auto_sync_enabled_sources: typing.Optional[ListUserResponseAutoSyncEnabledSources] = Field(None, alias='auto_sync_enabled_sources')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
