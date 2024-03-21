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

from carbon.type.data_source_type_nullable import DataSourceTypeNullable
from carbon.type.document_response_tags import DocumentResponseTags
from carbon.type.document_response_vector import DocumentResponseVector

class RequiredDocumentResponse(TypedDict):
    tags: typing.Optional[DocumentResponseTags]

    content: str

    file_id: int

    source: typing.Optional[str]

    source_url: typing.Optional[str]

    source_type: typing.Optional[DataSourceTypeNullable]

    presigned_url: typing.Optional[str]

    vector: typing.Optional[DocumentResponseVector]

    score: typing.Optional[typing.Union[int, float]]

    rank: typing.Union[typing.Union[int, float], int]

    content_metadata: typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]]

class OptionalDocumentResponse(TypedDict, total=False):
    pass

class DocumentResponse(RequiredDocumentResponse, OptionalDocumentResponse):
    pass
