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

from carbon.pydantic.data_source_type_nullable import DataSourceTypeNullable
from carbon.pydantic.organization_user_data_source_filters_ids import OrganizationUserDataSourceFiltersIds

class OrganizationUserDataSourceFilters(BaseModel):
    source: typing.Optional[DataSourceTypeNullable] = Field(None, alias='source')

    ids: typing.Optional[OrganizationUserDataSourceFiltersIds] = Field(None, alias='ids')

    revoked_access: typing.Optional[typing.Optional[bool]] = Field(None, alias='revoked_access')
    class Config:
        arbitrary_types_allowed = True
