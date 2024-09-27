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

from carbon.type.base_includes import BaseIncludes
from carbon.type.lead_filters import LeadFilters
from carbon.type.leads_order_by_nullable import LeadsOrderByNullable
from carbon.type.order_dir_v2_nullable import OrderDirV2Nullable

class RequiredLeadsRequest(TypedDict):
    data_source_id: int


class OptionalLeadsRequest(TypedDict, total=False):
    include_remote_data: bool

    next_cursor: typing.Optional[str]

    page_size: typing.Optional[int]

    order_dir: typing.Optional[OrderDirV2Nullable]

    includes: typing.List[BaseIncludes]

    filters: LeadFilters

    order_by: typing.Optional[LeadsOrderByNullable]

class LeadsRequest(RequiredLeadsRequest, OptionalLeadsRequest):
    pass
