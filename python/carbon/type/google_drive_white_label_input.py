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

from carbon.type.google_drive_credentials import GoogleDriveCredentials

class RequiredGoogleDriveWhiteLabelInput(TypedDict):
    data_source_type: typing.Union[bool, date, datetime, dict, float, int, list, str, None]

    credentials: GoogleDriveCredentials

class OptionalGoogleDriveWhiteLabelInput(TypedDict, total=False):
    pass

class GoogleDriveWhiteLabelInput(RequiredGoogleDriveWhiteLabelInput, OptionalGoogleDriveWhiteLabelInput):
    pass