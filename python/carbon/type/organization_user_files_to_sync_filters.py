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
from carbon.type.embedding_generators import EmbeddingGenerators
from carbon.type.external_file_sync_statuses import ExternalFileSyncStatuses
from carbon.type.organization_user_files_to_sync_filters_external_file_ids import OrganizationUserFilesToSyncFiltersExternalFileIds
from carbon.type.organization_user_files_to_sync_filters_ids import OrganizationUserFilesToSyncFiltersIds
from carbon.type.organization_user_files_to_sync_filters_organization_user_data_source_id import OrganizationUserFilesToSyncFiltersOrganizationUserDataSourceId
from carbon.type.organization_user_files_to_sync_filters_parent_file_ids import OrganizationUserFilesToSyncFiltersParentFileIds
from carbon.type.organization_user_files_to_sync_filters_tags import OrganizationUserFilesToSyncFiltersTags

class RequiredOrganizationUserFilesToSyncFilters(TypedDict):
    pass

class OptionalOrganizationUserFilesToSyncFilters(TypedDict, total=False):
    # WARNING: This property is deprecated
    tags: typing.Optional[OrganizationUserFilesToSyncFiltersTags]

    # The source of the file. If a list is provided, the query will return files from any of the sources in the list.
    source: typing.Union[DataSourceType, typing.List[DataSourceType]]

    # The name of the file. The query will return files with names that contain this string.
    name: typing.Optional[str]

    #          Tags to filter by. Supports logical AND and OR operations. Input should be like below:         {             \"OR\": [                 {                 \"key\": \"subject\",                 \"value\": \"holy-bible\",                 \"negate\": false                 },                 {                     \"key\": \"person-of-interest\",                     \"value\": \"jesus christ\",                     \"negate\": false                 },                 {                     \"key\": \"genre\",                     \"value\": \"fiction\",                     \"negate\": true                 }                 {                     \"AND\": [                         {                             \"key\": \"subject\",                             \"value\": \"tao-te-ching\",                             \"negate\": true                         },                         {                             \"key\": \"author\",                             \"value\": \"lao-tzu\",                             \"negate\": false                         }                     ]                 }             ]         }         For a single filter, the filter block can be placed within either an \"AND\" or \"OR\" block.         
    tags_v2: typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]]

    ids: typing.Optional[OrganizationUserFilesToSyncFiltersIds]

    external_file_ids: typing.Optional[OrganizationUserFilesToSyncFiltersExternalFileIds]

    # The sync statuses of the files. The query will return files with these sync statuses.
    sync_statuses: typing.Optional[typing.List[ExternalFileSyncStatuses]]

    # WARNING: This property is deprecated
    parent_file_ids: typing.Optional[OrganizationUserFilesToSyncFiltersParentFileIds]

    organization_user_data_source_id: typing.Optional[OrganizationUserFilesToSyncFiltersOrganizationUserDataSourceId]

    # The embedding generators of the files. The query will return files with these embedding generators.
    embedding_generators: typing.Optional[typing.List[EmbeddingGenerators]]

    # If true, the query will return only root files. Cannot be true if parent_file_ids or include_all_children is specified.
    root_files_only: typing.Optional[bool]

    # If true, the query will return all descendents of the specified parent_file_ids.
    include_all_children: bool

    # If true, the query will return only files that have not been synced yet.
    non_synced_only: bool

class OrganizationUserFilesToSyncFilters(RequiredOrganizationUserFilesToSyncFilters, OptionalOrganizationUserFilesToSyncFilters):
    pass
