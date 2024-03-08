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


class EmbeddingGenerators(
    schemas.EnumBase,
    schemas.StrSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        enum_value_to_name = {
            "OPENAI": "OPENAI",
            "AZURE_OPENAI": "AZURE_OPENAI",
            "AZURE_ADA_LARGE_256": "AZURE_ADA_LARGE_256",
            "AZURE_ADA_LARGE_1024": "AZURE_ADA_LARGE_1024",
            "AZURE_ADA_LARGE_3072": "AZURE_ADA_LARGE_3072",
            "AZURE_ADA_SMALL_512": "AZURE_ADA_SMALL_512",
            "AZURE_ADA_SMALL_1536": "AZURE_ADA_SMALL_1536",
            "COHERE_MULTILINGUAL_V3": "COHERE_MULTILINGUAL_V3",
            "VERTEX_MULTIMODAL": "VERTEX_MULTIMODAL",
            "OPENAI_ADA_LARGE_256": "OPENAI_ADA_LARGE_256",
            "OPENAI_ADA_LARGE_1024": "OPENAI_ADA_LARGE_1024",
            "OPENAI_ADA_LARGE_3072": "OPENAI_ADA_LARGE_3072",
            "OPENAI_ADA_SMALL_512": "OPENAI_ADA_SMALL_512",
            "OPENAI_ADA_SMALL_1536": "OPENAI_ADA_SMALL_1536",
        }
    
    @schemas.classproperty
    def OPENAI(cls):
        return cls("OPENAI")
    
    @schemas.classproperty
    def AZURE_OPENAI(cls):
        return cls("AZURE_OPENAI")
    
    @schemas.classproperty
    def AZURE_ADA_LARGE_256(cls):
        return cls("AZURE_ADA_LARGE_256")
    
    @schemas.classproperty
    def AZURE_ADA_LARGE_1024(cls):
        return cls("AZURE_ADA_LARGE_1024")
    
    @schemas.classproperty
    def AZURE_ADA_LARGE_3072(cls):
        return cls("AZURE_ADA_LARGE_3072")
    
    @schemas.classproperty
    def AZURE_ADA_SMALL_512(cls):
        return cls("AZURE_ADA_SMALL_512")
    
    @schemas.classproperty
    def AZURE_ADA_SMALL_1536(cls):
        return cls("AZURE_ADA_SMALL_1536")
    
    @schemas.classproperty
    def COHERE_MULTILINGUAL_V3(cls):
        return cls("COHERE_MULTILINGUAL_V3")
    
    @schemas.classproperty
    def VERTEX_MULTIMODAL(cls):
        return cls("VERTEX_MULTIMODAL")
    
    @schemas.classproperty
    def OPENAI_ADA_LARGE_256(cls):
        return cls("OPENAI_ADA_LARGE_256")
    
    @schemas.classproperty
    def OPENAI_ADA_LARGE_1024(cls):
        return cls("OPENAI_ADA_LARGE_1024")
    
    @schemas.classproperty
    def OPENAI_ADA_LARGE_3072(cls):
        return cls("OPENAI_ADA_LARGE_3072")
    
    @schemas.classproperty
    def OPENAI_ADA_SMALL_512(cls):
        return cls("OPENAI_ADA_SMALL_512")
    
    @schemas.classproperty
    def OPENAI_ADA_SMALL_1536(cls):
        return cls("OPENAI_ADA_SMALL_1536")
