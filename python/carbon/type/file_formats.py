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


FileFormats = Literal["TXT", "CSV", "TSV", "PDF", "DOCX", "PPTX", "XLSX", "XLSM", "MD", "RTF", "JSON", "HTML", "NOTION", "GOOGLE_DOCS", "GOOGLE_SHEETS", "GOOGLE_SLIDES", "INTERCOM", "CONFLUENCE", "RSS_FEED", "GMAIL", "OUTLOOK", "ZENDESK", "FRESHDESK", "WEB_SCRAPE", "GITBOOK", "SALESFORCE", "GITHUB", "SLACK", "GURU", "SERVICENOW", "GONG", "DOCUMENT360", "JPG", "PNG", "MP3", "MP2", "AAC", "WAV", "FLAC", "PCM", "M4A", "OGG", "OPUS", "MPEG", "MPG", "MP4", "WMV", "AVI", "MOV", "MKV", "FLV", "WEBM", "EML", "MSG"]
