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


class EmbeddingAndChunk(
    schemas.DictSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        required = {
            "chunk_index",
            "source_content",
            "user_file_id",
            "embedding",
        }
        
        class properties:
            user_file_id = schemas.IntSchema
            
            
            class chunk_index(
                schemas.IntBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneDecimalMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, decimal.Decimal, int, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'chunk_index':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            source_content = schemas.StrSchema
        
            @staticmethod
            def embedding() -> typing.Type['EmbeddingAndChunkEmbedding']:
                return EmbeddingAndChunkEmbedding
            __annotations__ = {
                "user_file_id": user_file_id,
                "chunk_index": chunk_index,
                "source_content": source_content,
                "embedding": embedding,
            }
    
    chunk_index: MetaOapg.properties.chunk_index
    source_content: MetaOapg.properties.source_content
    user_file_id: MetaOapg.properties.user_file_id
    embedding: 'EmbeddingAndChunkEmbedding'
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["user_file_id"]) -> MetaOapg.properties.user_file_id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["chunk_index"]) -> MetaOapg.properties.chunk_index: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["source_content"]) -> MetaOapg.properties.source_content: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["embedding"]) -> 'EmbeddingAndChunkEmbedding': ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["user_file_id", "chunk_index", "source_content", "embedding", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["user_file_id"]) -> MetaOapg.properties.user_file_id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["chunk_index"]) -> MetaOapg.properties.chunk_index: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["source_content"]) -> MetaOapg.properties.source_content: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["embedding"]) -> 'EmbeddingAndChunkEmbedding': ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["user_file_id", "chunk_index", "source_content", "embedding", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        chunk_index: typing.Union[MetaOapg.properties.chunk_index, None, decimal.Decimal, int, ],
        source_content: typing.Union[MetaOapg.properties.source_content, str, ],
        user_file_id: typing.Union[MetaOapg.properties.user_file_id, decimal.Decimal, int, ],
        embedding: 'EmbeddingAndChunkEmbedding',
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'EmbeddingAndChunk':
        return super().__new__(
            cls,
            *args,
            chunk_index=chunk_index,
            source_content=source_content,
            user_file_id=user_file_id,
            embedding=embedding,
            _configuration=_configuration,
            **kwargs,
        )

from carbon.model.embedding_and_chunk_embedding import EmbeddingAndChunkEmbedding
