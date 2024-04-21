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

from carbon.type.data_source_last_sync_actions import DataSourceLastSyncActions
from carbon.type.data_source_sync_statuses import DataSourceSyncStatuses
from carbon.type.data_source_type import DataSourceType

class RequiredOrganizationUserDataSourceAPI(TypedDict):
    id: int

    data_source_external_id: typing.Optional[str]

    data_source_type: DataSourceType

    token: typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]]

    sync_status: DataSourceSyncStatuses

    source_items_synced_at: typing.Optional[datetime]

    organization_user_id: int

    organization_id: int

    organization_supplied_user_id: str

    revoked_access: bool

    last_synced_at: datetime

    last_sync_action: DataSourceLastSyncActions

    enable_auto_sync: typing.Optional[bool]

    created_at: datetime

    updated_at: datetime

    files_synced_at: typing.Optional[datetime]

class OptionalOrganizationUserDataSourceAPI(TypedDict, total=False):
    pass

class OrganizationUserDataSourceAPI(RequiredOrganizationUserDataSourceAPI, OptionalOrganizationUserDataSourceAPI):
    pass
