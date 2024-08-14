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

from carbon.type.helpdesk_file_types import HelpdeskFileTypes
from carbon.type.transcription_service_nullable import TranscriptionServiceNullable

class RequiredFileSyncConfig(TypedDict):
    pass

class OptionalFileSyncConfig(TypedDict, total=False):
    # File types to automatically sync when the data source connects. Only a subset of file types can be          controlled. If not supported, then they will always be synced
    auto_synced_source_types: typing.List[HelpdeskFileTypes]

    # Automatically sync attachments from files where supported. Currently applies to Helpdesk Tickets
    sync_attachments: bool

    # Detect audio language before transcription for audio files
    detect_audio_language: bool

    transcription_service: typing.Optional[TranscriptionServiceNullable]

    # Detect multiple speakers and label segments of speech by speaker for audio files.
    include_speaker_labels: bool

    # Whether to split tabular rows into chunks. Currently only valid for CSV, TSV, and XLSX files.
    split_rows: bool

class FileSyncConfig(RequiredFileSyncConfig, OptionalFileSyncConfig):
    pass
