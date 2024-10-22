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

from carbon.type.remove_data_source_tags_input_tags_to_remove import RemoveDataSourceTagsInputTagsToRemove

class RequiredRemoveDataSourceTagsInput(TypedDict):
    data_source_id: int


class OptionalRemoveDataSourceTagsInput(TypedDict, total=False):
    tags_to_remove: RemoveDataSourceTagsInputTagsToRemove

    remove_all_tags: bool

class RemoveDataSourceTagsInput(RequiredRemoveDataSourceTagsInput, OptionalRemoveDataSourceTagsInput):
    pass
