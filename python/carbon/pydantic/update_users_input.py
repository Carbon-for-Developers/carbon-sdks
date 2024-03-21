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

from carbon.pydantic.data_source_type import DataSourceType
from carbon.pydantic.update_users_input_customer_ids import UpdateUsersInputCustomerIds

class UpdateUsersInput(BaseModel):
    customer_ids: UpdateUsersInputCustomerIds = Field(alias='customer_ids')

    # List of data source types to enable auto sync for. Empty array will remove all sources          and the string \"ALL\" will enable it for all data sources
    auto_sync_enabled_sources: typing.Optional[typing.Union[typing.List[DataSourceType], str]] = Field(None, alias='auto_sync_enabled_sources')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
