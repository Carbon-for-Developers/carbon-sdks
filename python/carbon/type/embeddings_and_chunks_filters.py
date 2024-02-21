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

class RequiredEmbeddingsAndChunksFilters(TypedDict):
    user_file_id: int

class OptionalEmbeddingsAndChunksFilters(TypedDict, total=False):
    embedding_model: EmbeddingGeneratorsNullable

class EmbeddingsAndChunksFilters(RequiredEmbeddingsAndChunksFilters, OptionalEmbeddingsAndChunksFilters):
    pass
