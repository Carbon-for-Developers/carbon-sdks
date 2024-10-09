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


class RequiredFilesInput(TypedDict):
    data_source_id: int

    # Full name of the repository, denoted as {owner}/{repo}
    repository: str

    pull_number: int


class OptionalFilesInput(TypedDict, total=False):
    include_remote_data: bool

    page: int

    page_size: int

    next_cursor: typing.Optional[str]

class FilesInput(RequiredFilesInput, OptionalFilesInput):
    pass
