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

from carbon.type.cold_storage_props import ColdStorageProps
from carbon.type.embedding_generators_nullable import EmbeddingGeneratorsNullable

class RequiredRawTextInput(TypedDict):
    contents: str


class OptionalRawTextInput(TypedDict, total=False):
    name: typing.Optional[str]

    chunk_size: typing.Optional[int]

    chunk_overlap: typing.Optional[int]

    skip_embedding_generation: bool

    overwrite_file_id: typing.Optional[int]

    embedding_model: typing.Optional[EmbeddingGeneratorsNullable]

    generate_sparse_vectors: typing.Optional[bool]

    cold_storage_params: ColdStorageProps

    # If this flag is enabled, the file will be chunked and stored with Carbon,         but no embeddings will be generated. This overrides the skip_embedding_generation flag.
    generate_chunks_only: bool

    # If this flag is enabled, the file will be stored with Carbon, but no processing will be done.
    store_file_only: bool

class RawTextInput(RequiredRawTextInput, OptionalRawTextInput):
    pass
