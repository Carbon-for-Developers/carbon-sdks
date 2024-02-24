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
from carbon.type.single_chunks_and_embeddings_upload_input import SingleChunksAndEmbeddingsUploadInput

class RequiredChunksAndEmbeddingsUploadInput(TypedDict):
    embedding_model: EmbeddingGenerators

    chunks_and_embeddings: typing.List[SingleChunksAndEmbeddingsUploadInput]

class OptionalChunksAndEmbeddingsUploadInput(TypedDict, total=False):
    overwrite_existing: bool

class ChunksAndEmbeddingsUploadInput(RequiredChunksAndEmbeddingsUploadInput, OptionalChunksAndEmbeddingsUploadInput):
    pass
