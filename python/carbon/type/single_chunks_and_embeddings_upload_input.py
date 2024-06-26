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

from carbon.type.chunks_and_embeddings import ChunksAndEmbeddings

class RequiredSingleChunksAndEmbeddingsUploadInput(TypedDict):
    file_id: int

    chunks_and_embeddings: typing.List[ChunksAndEmbeddings]


class OptionalSingleChunksAndEmbeddingsUploadInput(TypedDict, total=False):
    chunk_size: typing.Optional[int]

    chunk_overlap: typing.Optional[int]

class SingleChunksAndEmbeddingsUploadInput(RequiredSingleChunksAndEmbeddingsUploadInput, OptionalSingleChunksAndEmbeddingsUploadInput):
    pass
