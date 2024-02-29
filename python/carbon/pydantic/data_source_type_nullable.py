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


DataSourceTypeNullable = Literal["GOOGLE_DRIVE", "NOTION", "NOTION_DATABASE", "INTERCOM", "DROPBOX", "ONEDRIVE", "SHAREPOINT", "CONFLUENCE", "BOX", "ZENDESK", "ZOTERO", "S3", "GMAIL", "OUTLOOK", "TEXT", "CSV", "TSV", "PDF", "DOCX", "PPTX", "XLSX", "MD", "RTF", "JSON", "RAW_TEXT", "WEB_SCRAPE", "RSS_FEED", "FRESHDESK", "GITBOOK", "SALESFORCE", "JPG", "PNG", "MP3", "MP4", "MP2", "AAC", "WAV", "FLAC", "PCM", "M4A", "OGG", "OPUS", "WEBM"]
