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

from carbon.pydantic.embedding_generators import EmbeddingGenerators
from carbon.pydantic.file_sync_config_nullable import FileSyncConfigNullable
from carbon.pydantic.gitbook_sync_request_space_ids import GitbookSyncRequestSpaceIds

class GitbookSyncRequest(BaseModel):
    space_ids: GitbookSyncRequestSpaceIds = Field(alias='space_ids')

    data_source_id: int = Field(alias='data_source_id')

    tags: typing.Optional[typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]]] = Field(None, alias='tags')

    chunk_size: typing.Optional[typing.Optional[int]] = Field(None, alias='chunk_size')

    chunk_overlap: typing.Optional[typing.Optional[int]] = Field(None, alias='chunk_overlap')

    skip_embedding_generation: typing.Optional[typing.Optional[bool]] = Field(None, alias='skip_embedding_generation')

    embedding_model: typing.Optional[EmbeddingGenerators] = Field(None, alias='embedding_model')

    generate_sparse_vectors: typing.Optional[typing.Optional[bool]] = Field(None, alias='generate_sparse_vectors')

    prepend_filename_to_chunks: typing.Optional[typing.Optional[bool]] = Field(None, alias='prepend_filename_to_chunks')

    request_id: typing.Optional[typing.Optional[str]] = Field(None, alias='request_id')

    file_sync_config: typing.Optional[FileSyncConfigNullable] = Field(None, alias='file_sync_config')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
