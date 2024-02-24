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

from carbon.type.file_formats_nullable import FileFormatsNullable

class RequiredFileStatisticsNullable(TypedDict):
    file_format: FileFormatsNullable

    file_size: typing.Optional[int]

    num_characters: typing.Optional[int]

    num_tokens: typing.Optional[int]

    num_embeddings: typing.Optional[int]

class OptionalFileStatisticsNullable(TypedDict, total=False):
    pass

class FileStatisticsNullable(RequiredFileStatisticsNullable, OptionalFileStatisticsNullable):
    pass
