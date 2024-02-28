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


FileFormatsNullable = Literal["TXT", "CSV", "TSV", "PDF", "DOCX", "PPTX", "XLSX", "MD", "RTF", "JSON", "NOTION", "GOOGLE_DRIVE", "GOOGLE_DOCS", "GOOGLE_SHEETS", "GOOGLE_SLIDES", "INTERCOM", "DROPBOX", "ONEDRIVE", "SHAREPOINT", "CONFLUENCE", "RSS_FEED", "GMAIL", "OUTLOOK", "BOX", "ZENDESK", "FRESHDESK", "WEB_SCRAPE", "GITBOOK", "JPG", "JPG", "PNG", "MP3", "MP4", "MP2", "AAC", "WAV", "FLAC", "PCM", "M4A", "OGG", "OPUS", "WEBM"]
