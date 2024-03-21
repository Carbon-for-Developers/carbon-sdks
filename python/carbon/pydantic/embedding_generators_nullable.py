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


EmbeddingGeneratorsNullable = Literal["OPENAI", "AZURE_OPENAI", "AZURE_ADA_LARGE_256", "AZURE_ADA_LARGE_1024", "AZURE_ADA_LARGE_3072", "AZURE_ADA_SMALL_512", "AZURE_ADA_SMALL_1536", "COHERE_MULTILINGUAL_V3", "VERTEX_MULTIMODAL", "OPENAI_ADA_LARGE_256", "OPENAI_ADA_LARGE_1024", "OPENAI_ADA_LARGE_3072", "OPENAI_ADA_SMALL_512", "OPENAI_ADA_SMALL_1536"]
