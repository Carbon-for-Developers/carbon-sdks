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

from carbon.type.chunks_and_embeddings_embedding import ChunksAndEmbeddingsEmbedding

class RequiredChunksAndEmbeddings(TypedDict):
    chunk_number: typing.Optional[int]

    chunk: str

    embedding: ChunksAndEmbeddingsEmbedding

class OptionalChunksAndEmbeddings(TypedDict, total=False):
    pass

class ChunksAndEmbeddings(RequiredChunksAndEmbeddings, OptionalChunksAndEmbeddings):
    pass
