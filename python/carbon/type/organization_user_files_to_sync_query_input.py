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

from carbon.type.order_dir import OrderDir
from carbon.type.organization_user_files_to_sync_filters import OrganizationUserFilesToSyncFilters
from carbon.type.organization_user_files_to_sync_order_by_types import OrganizationUserFilesToSyncOrderByTypes
from carbon.type.pagination import Pagination

class RequiredOrganizationUserFilesToSyncQueryInput(TypedDict):
    pass

class OptionalOrganizationUserFilesToSyncQueryInput(TypedDict, total=False):
    pagination: Pagination

    order_by: OrganizationUserFilesToSyncOrderByTypes

    order_dir: OrderDir

    filters: OrganizationUserFilesToSyncFilters

    # If true, the query will return presigned URLs for the raw file. Only relevant for the /user_files_v2 endpoint.
    include_raw_file: typing.Optional[bool]

    # If true, the query will return presigned URLs for the parsed text file. Only relevant for the /user_files_v2 endpoint.
    include_parsed_text_file: typing.Optional[bool]

    # If true, the query will return presigned URLs for additional files. Only relevant for the /user_files_v2 endpoint.
    include_additional_files: typing.Optional[bool]

    # The expiry time for the presigned URLs. Only relevant for the /user_files_v2 endpoint.
    presigned_url_expiry_time_seconds: int

class OrganizationUserFilesToSyncQueryInput(RequiredOrganizationUserFilesToSyncQueryInput, OptionalOrganizationUserFilesToSyncQueryInput):
    pass
