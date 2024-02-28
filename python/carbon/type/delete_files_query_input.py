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

from carbon.type.delete_files_query_input_file_ids import DeleteFilesQueryInputFileIds
from carbon.type.external_file_sync_statuses import ExternalFileSyncStatuses

class RequiredDeleteFilesQueryInput(TypedDict):
    pass

class OptionalDeleteFilesQueryInput(TypedDict, total=False):
    file_ids: DeleteFilesQueryInputFileIds

    sync_statuses: typing.Optional[typing.List[ExternalFileSyncStatuses]]

    delete_non_synced_only: bool

    send_webhook: bool

class DeleteFilesQueryInput(RequiredDeleteFilesQueryInput, OptionalDeleteFilesQueryInput):
    pass
