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
from pydantic import BaseModel, Field, RootModel

from carbon.pydantic.chunks_and_embeddings_embedding import ChunksAndEmbeddingsEmbedding

class ChunksAndEmbeddings(BaseModel):
    chunk_number: typing.Optional[int] = Field(alias='chunk_number')

    chunk: str = Field(alias='chunk')

    embedding: ChunksAndEmbeddingsEmbedding = Field(alias='embedding')
