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


class RequiredOneDriveSharepointCredentials(TypedDict):
    client_id: str

    redirect_uri: str


class OptionalOneDriveSharepointCredentials(TypedDict, total=False):
    client_secret: typing.Optional[str]

    file_picker_client_id: typing.Optional[str]

    file_picker_redirect_uri: typing.Optional[str]

class OneDriveSharepointCredentials(RequiredOneDriveSharepointCredentials, OptionalOneDriveSharepointCredentials):
    pass
