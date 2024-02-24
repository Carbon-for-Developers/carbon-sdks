# coding: utf-8

"""
    Carbon

    Connect external data to LLMs, no matter the source.

    The version of the OpenAPI document: 1.0.0
    Generated by: https://konfigthis.com
"""

from datetime import date, datetime  # noqa: F401
import decimal  # noqa: F401
import functools  # noqa: F401
import io  # noqa: F401
import re  # noqa: F401
import typing  # noqa: F401
import typing_extensions  # noqa: F401
import uuid  # noqa: F401

import frozendict  # noqa: F401

from carbon import schemas  # noqa: F401


class YoutubeTranscriptResponseRawTranscript(
    schemas.ListBase,
    schemas.NoneBase,
    schemas.Schema,
    schemas.NoneTupleMixin
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        
        @staticmethod
        def items() -> typing.Type['YoutubeTranscriptResponseRawTranscriptItem']:
            return YoutubeTranscriptResponseRawTranscriptItem


    def __new__(
        cls,
        *args: typing.Union[list, tuple, None, ],
        _configuration: typing.Optional[schemas.Configuration] = None,
    ) -> 'YoutubeTranscriptResponseRawTranscript':
        return super().__new__(
            cls,
            *args,
            _configuration=_configuration,
        )

from carbon.model.youtube_transcript_response_raw_transcript_item import YoutubeTranscriptResponseRawTranscriptItem
