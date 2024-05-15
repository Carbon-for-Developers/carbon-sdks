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
from carbon.type.embedding_generators_nullable import EmbeddingGeneratorsNullable
from carbon.type.helpdesk_file_sync_config_nullable import HelpdeskFileSyncConfigNullable

class RequiredOAuthURLRequest(TypedDict):
    service: DataSourceType

class OptionalOAuthURLRequest(TypedDict, total=False):
    tags: typing.Union[bool, date, datetime, dict, float, int, list, str, None]

    scope: typing.Optional[str]

    chunk_size: typing.Optional[int]

    chunk_overlap: typing.Optional[int]

    skip_embedding_generation: typing.Optional[bool]

    embedding_model: typing.Optional[EmbeddingGeneratorsNullable]

    zendesk_subdomain: typing.Optional[str]

    microsoft_tenant: typing.Optional[str]

    sharepoint_site_name: typing.Optional[str]

    confluence_subdomain: typing.Optional[str]

    generate_sparse_vectors: typing.Optional[bool]

    prepend_filename_to_chunks: typing.Optional[bool]

    # Number of objects per chunk. For csv, tsv, xlsx, and json files only.
    max_items_per_chunk: typing.Optional[int]

    salesforce_domain: typing.Optional[str]

    # Used to specify whether Carbon should attempt to sync all your files automatically when authorization         is complete. This is only supported for a subset of connectors and will be ignored for the rest. Supported         connectors: Intercom, Zendesk, Gitbook, Confluence, Salesforce, Freshdesk
    sync_files_on_connection: typing.Optional[bool]

    set_page_as_boundary: bool

    # Used to specify a data source to sync from if you have multiple connected. It can be skipped if          you only have one data source of that type connected or are connecting a new account.
    data_source_id: typing.Optional[int]

    # Used to connect a new data source. If not specified, we will attempt to create a sync URL         for an existing data source based on type and ID.
    connecting_new_account: typing.Optional[bool]

    # This request id will be added to all files that get synced using the generated OAuth URL
    request_id: str

    # Enable OCR for files that support it. Supported formats: pdf
    use_ocr: typing.Optional[bool]

    parse_pdf_tables_with_ocr: typing.Optional[bool]

    # Enable integration's file picker for sources that support it. Supported sources: GOOGLE_DRIVE, ONEDRIVE, BOX, SHAREPOINT, DROPBOX
    enable_file_picker: bool

    # Enabling this flag will fetch all available content from the source to be listed via list items endpoint
    sync_source_items: bool

    # Only sync files if they have not already been synced or if the embedding properties have changed.         This flag is currently supported by ONEDRIVE, GOOGLE_DRIVE, BOX, DROPBOX. It will be ignored for other data sources.
    incremental_sync: bool

    file_sync_config: typing.Optional[HelpdeskFileSyncConfigNullable]

class OAuthURLRequest(RequiredOAuthURLRequest, OptionalOAuthURLRequest):
    pass
