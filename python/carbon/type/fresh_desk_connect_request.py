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

from carbon.type.embedding_generators_nullable import EmbeddingGeneratorsNullable
from carbon.type.file_sync_config_nullable import FileSyncConfigNullable

class RequiredFreshDeskConnectRequest(TypedDict):
    domain: str

    api_key: str


class OptionalFreshDeskConnectRequest(TypedDict, total=False):
    tags: typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]]

    chunk_size: typing.Optional[int]

    chunk_overlap: typing.Optional[int]

    skip_embedding_generation: typing.Optional[bool]

    embedding_model: typing.Optional[EmbeddingGeneratorsNullable]

    generate_sparse_vectors: typing.Optional[bool]

    prepend_filename_to_chunks: typing.Optional[bool]

    sync_files_on_connection: typing.Optional[bool]

    request_id: typing.Optional[str]

    # Enabling this flag will fetch all available content from the source to be listed via list items endpoint
    sync_source_items: bool

    file_sync_config: typing.Optional[FileSyncConfigNullable]

    # Tags to be associated with the data source. If the data source already has tags set, then an upsert will be performed.
    data_source_tags: typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]

class FreshDeskConnectRequest(RequiredFreshDeskConnectRequest, OptionalFreshDeskConnectRequest):
    pass
