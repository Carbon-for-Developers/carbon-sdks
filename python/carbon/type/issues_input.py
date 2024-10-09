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

from carbon.type.issues_filter import IssuesFilter
from carbon.type.issues_order_by import IssuesOrderBy
from carbon.type.order_dir_v2_nullable import OrderDirV2Nullable

class RequiredIssuesInput(TypedDict):
    data_source_id: int

    # Full name of the repository, denoted as {owner}/{repo}
    repository: str


class OptionalIssuesInput(TypedDict, total=False):
    include_remote_data: bool

    page: int

    page_size: int

    next_cursor: typing.Optional[str]

    filters: IssuesFilter

    order_by: IssuesOrderBy

    order_dir: typing.Optional[OrderDirV2Nullable]

class IssuesInput(RequiredIssuesInput, OptionalIssuesInput):
    pass
