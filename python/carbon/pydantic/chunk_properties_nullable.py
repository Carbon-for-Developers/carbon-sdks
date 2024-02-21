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


class ChunkPropertiesNullable(BaseModel):
    set_page_as_boundary: bool = Field(alias='set_page_as_boundary')

    prepend_filename_to_chunks: bool = Field(alias='prepend_filename_to_chunks')

    max_items_per_chunk: typing.Optional[int] = Field(alias='max_items_per_chunk')
