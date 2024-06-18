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

from carbon.pydantic.chunk_properties_nullable import ChunkPropertiesNullable
from carbon.pydantic.data_source_type import DataSourceType
from carbon.pydantic.external_file_sync_statuses import ExternalFileSyncStatuses
from carbon.pydantic.file_statistics_nullable import FileStatisticsNullable
from carbon.pydantic.user_file_embedding_properties import UserFileEmbeddingProperties

class UserFile(BaseModel):
    tags: typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]] = Field(alias='tags')

    id: int = Field(alias='id')

    source: DataSourceType = Field(alias='source')

    organization_id: int = Field(alias='organization_id')

    organization_supplied_user_id: str = Field(alias='organization_supplied_user_id')

    organization_user_data_source_id: typing.Optional[int] = Field(alias='organization_user_data_source_id')

    external_file_id: str = Field(alias='external_file_id')

    external_url: typing.Optional[str] = Field(alias='external_url')

    sync_status: ExternalFileSyncStatuses = Field(alias='sync_status')

    sync_error_message: typing.Optional[str] = Field(alias='sync_error_message')

    last_sync: typing.Optional[datetime] = Field(alias='last_sync')

    file_statistics: FileStatisticsNullable = Field(alias='file_statistics')

    file_metadata: typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]] = Field(alias='file_metadata')

    embedding_properties: UserFileEmbeddingProperties = Field(alias='embedding_properties')

    chunk_size: typing.Optional[int] = Field(alias='chunk_size')

    chunk_overlap: typing.Optional[int] = Field(alias='chunk_overlap')

    chunk_properties: ChunkPropertiesNullable = Field(alias='chunk_properties')

    ocr_properties: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]] = Field(alias='ocr_properties')

    ocr_job_started_at: typing.Optional[datetime] = Field(alias='ocr_job_started_at')

    name: typing.Optional[str] = Field(alias='name')

    parent_id: typing.Optional[int] = Field(alias='parent_id')

    enable_auto_sync: typing.Optional[bool] = Field(alias='enable_auto_sync')

    presigned_url: typing.Optional[str] = Field(alias='presigned_url')

    parsed_text_url: typing.Optional[str] = Field(alias='parsed_text_url')

    additional_presigned_urls: typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]] = Field(alias='additional_presigned_urls')

    skip_embedding_generation: bool = Field(alias='skip_embedding_generation')

    source_created_at: typing.Optional[datetime] = Field(alias='source_created_at')

    generate_sparse_vectors: typing.Optional[bool] = Field(alias='generate_sparse_vectors')

    request_id: typing.Optional[str] = Field(alias='request_id')

    sync_properties: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]] = Field(alias='sync_properties')

    messages_metadata: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]] = Field(alias='messages_metadata')

    created_at: datetime = Field(alias='created_at')

    updated_at: datetime = Field(alias='updated_at')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
