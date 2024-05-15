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
from carbon.pydantic.embedding_generators_nullable import EmbeddingGeneratorsNullable
from carbon.pydantic.helpdesk_file_sync_config_nullable import HelpdeskFileSyncConfigNullable

class OAuthURLRequest(BaseModel):
    service: DataSourceType = Field(alias='service')

    tags: typing.Optional[typing.Union[bool, date, datetime, dict, float, int, list, str, None]] = Field(None, alias='tags')

    scope: typing.Optional[typing.Optional[str]] = Field(None, alias='scope')

    chunk_size: typing.Optional[typing.Optional[int]] = Field(None, alias='chunk_size')

    chunk_overlap: typing.Optional[typing.Optional[int]] = Field(None, alias='chunk_overlap')

    skip_embedding_generation: typing.Optional[typing.Optional[bool]] = Field(None, alias='skip_embedding_generation')

    embedding_model: typing.Optional[EmbeddingGeneratorsNullable] = Field(None, alias='embedding_model')

    zendesk_subdomain: typing.Optional[typing.Optional[str]] = Field(None, alias='zendesk_subdomain')

    microsoft_tenant: typing.Optional[typing.Optional[str]] = Field(None, alias='microsoft_tenant')

    sharepoint_site_name: typing.Optional[typing.Optional[str]] = Field(None, alias='sharepoint_site_name')

    confluence_subdomain: typing.Optional[typing.Optional[str]] = Field(None, alias='confluence_subdomain')

    generate_sparse_vectors: typing.Optional[typing.Optional[bool]] = Field(None, alias='generate_sparse_vectors')

    prepend_filename_to_chunks: typing.Optional[typing.Optional[bool]] = Field(None, alias='prepend_filename_to_chunks')

    # Number of objects per chunk. For csv, tsv, xlsx, and json files only.
    max_items_per_chunk: typing.Optional[typing.Optional[int]] = Field(None, alias='max_items_per_chunk')

    salesforce_domain: typing.Optional[typing.Optional[str]] = Field(None, alias='salesforce_domain')

    # Used to specify whether Carbon should attempt to sync all your files automatically when authorization         is complete. This is only supported for a subset of connectors and will be ignored for the rest. Supported         connectors: Intercom, Zendesk, Gitbook, Confluence, Salesforce, Freshdesk
    sync_files_on_connection: typing.Optional[typing.Optional[bool]] = Field(None, alias='sync_files_on_connection')

    set_page_as_boundary: typing.Optional[bool] = Field(None, alias='set_page_as_boundary')

    # Used to specify a data source to sync from if you have multiple connected. It can be skipped if          you only have one data source of that type connected or are connecting a new account.
    data_source_id: typing.Optional[typing.Optional[int]] = Field(None, alias='data_source_id')

    # Used to connect a new data source. If not specified, we will attempt to create a sync URL         for an existing data source based on type and ID.
    connecting_new_account: typing.Optional[typing.Optional[bool]] = Field(None, alias='connecting_new_account')

    # This request id will be added to all files that get synced using the generated OAuth URL
    request_id: typing.Optional[str] = Field(None, alias='request_id')

    # Enable OCR for files that support it. Supported formats: pdf
    use_ocr: typing.Optional[typing.Optional[bool]] = Field(None, alias='use_ocr')

    parse_pdf_tables_with_ocr: typing.Optional[typing.Optional[bool]] = Field(None, alias='parse_pdf_tables_with_ocr')

    # Enable integration's file picker for sources that support it. Supported sources: GOOGLE_DRIVE, ONEDRIVE, BOX, SHAREPOINT, DROPBOX
    enable_file_picker: typing.Optional[bool] = Field(None, alias='enable_file_picker')

    # Enabling this flag will fetch all available content from the source to be listed via list items endpoint
    sync_source_items: typing.Optional[bool] = Field(None, alias='sync_source_items')

    # Only sync files if they have not already been synced or if the embedding properties have changed.         This flag is currently supported by ONEDRIVE, GOOGLE_DRIVE, BOX, DROPBOX. It will be ignored for other data sources.
    incremental_sync: typing.Optional[bool] = Field(None, alias='incremental_sync')

    file_sync_config: typing.Optional[HelpdeskFileSyncConfigNullable] = Field(None, alias='file_sync_config')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
