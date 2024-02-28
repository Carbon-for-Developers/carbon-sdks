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

from carbon.type.embedding_generators import EmbeddingGenerators
from carbon.type.gitbook_sync_request_space_ids import GitbookSyncRequestSpaceIds

class RequiredGitbookSyncRequest(TypedDict):
    space_ids: GitbookSyncRequestSpaceIds

    data_source_id: int

class OptionalGitbookSyncRequest(TypedDict, total=False):
    tags: typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]]

    chunk_size: typing.Optional[int]

    chunk_overlap: typing.Optional[int]

    skip_embedding_generation: typing.Optional[bool]

    embedding_model: EmbeddingGenerators

    generate_sparse_vectors: typing.Optional[bool]

    prepend_filename_to_chunks: typing.Optional[bool]

class GitbookSyncRequest(RequiredGitbookSyncRequest, OptionalGitbookSyncRequest):
    pass
