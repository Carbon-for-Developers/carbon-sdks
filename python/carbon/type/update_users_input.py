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

from carbon.type.data_source_type import DataSourceType
from carbon.type.update_users_input_customer_ids import UpdateUsersInputCustomerIds

class RequiredUpdateUsersInput(TypedDict):
    customer_ids: UpdateUsersInputCustomerIds


class OptionalUpdateUsersInput(TypedDict, total=False):
    # List of data source types to enable auto sync for. Empty array will remove all sources          and the string \"ALL\" will enable it for all data sources
    auto_sync_enabled_sources: typing.Union[typing.List[DataSourceType], str]

    # Custom file upload limit for the user over *all* user's files across all uploads.          If set, then the user will not be allowed to upload more files than this limit. If not set, or if set to -1,         then the user will have no limit.
    max_files: typing.Optional[int]

    # Custom file upload limit for the user across a single upload.         If set, then the user will not be allowed to upload more files than this limit in a single upload. If not set,         or if set to -1, then the user will have no limit.
    max_files_per_upload: typing.Optional[int]

class UpdateUsersInput(RequiredUpdateUsersInput, OptionalUpdateUsersInput):
    pass
