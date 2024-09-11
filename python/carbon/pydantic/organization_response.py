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


class OrganizationResponse(BaseModel):
    id: int = Field(alias='id')

    name: str = Field(alias='name')

    nickname: typing.Optional[str] = Field(alias='nickname')

    remove_branding: bool = Field(alias='remove_branding')

    custom_branding: typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]] = Field(alias='custom_branding')

    custom_limits: typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]] = Field(alias='custom_limits')

    aggregate_file_size: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]] = Field(alias='aggregate_file_size')

    aggregate_num_characters: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]] = Field(alias='aggregate_num_characters')

    aggregate_num_tokens: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]] = Field(alias='aggregate_num_tokens')

    aggregate_num_embeddings: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]] = Field(alias='aggregate_num_embeddings')

    aggregate_num_files_by_source: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]] = Field(alias='aggregate_num_files_by_source')

    aggregate_num_files_by_file_format: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]] = Field(alias='aggregate_num_files_by_file_format')

    file_statistics_aggregated_at: typing.Optional[datetime] = Field(alias='file_statistics_aggregated_at')

    period_ends_at: typing.Optional[datetime] = Field(alias='period_ends_at')

    cancel_at_period_end: typing.Optional[bool] = Field(alias='cancel_at_period_end')

    connector_settings: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]] = Field(alias='connector_settings')

    global_user_config: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]] = Field(alias='global_user_config')

    file_sync_usage: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]] = Field(alias='file_sync_usage')

    logging_settings: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]] = Field(alias='logging_settings')

    created_at: datetime = Field(alias='created_at')

    updated_at: datetime = Field(alias='updated_at')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
