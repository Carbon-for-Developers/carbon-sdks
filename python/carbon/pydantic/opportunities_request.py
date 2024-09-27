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

from carbon.pydantic.base_includes import BaseIncludes
from carbon.pydantic.opportunities_order_by_nullable import OpportunitiesOrderByNullable
from carbon.pydantic.opportunity_filters import OpportunityFilters
from carbon.pydantic.order_dir_v2_nullable import OrderDirV2Nullable

class OpportunitiesRequest(BaseModel):
    data_source_id: int = Field(alias='data_source_id')

    include_remote_data: typing.Optional[bool] = Field(None, alias='include_remote_data')

    next_cursor: typing.Optional[typing.Optional[str]] = Field(None, alias='next_cursor')

    page_size: typing.Optional[typing.Optional[int]] = Field(None, alias='page_size')

    order_dir: typing.Optional[OrderDirV2Nullable] = Field(None, alias='order_dir')

    includes: typing.Optional[typing.List[BaseIncludes]] = Field(None, alias='includes')

    filters: typing.Optional[OpportunityFilters] = Field(None, alias='filters')

    order_by: typing.Optional[OpportunitiesOrderByNullable] = Field(None, alias='order_by')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
