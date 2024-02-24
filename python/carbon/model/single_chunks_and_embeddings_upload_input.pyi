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


class SingleChunksAndEmbeddingsUploadInput(
    schemas.DictSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        required = {
            "chunks_and_embeddings",
            "file_id",
        }
        
        class properties:
            file_id = schemas.IntSchema
            
            
            class chunks_and_embeddings(
                schemas.ListSchema
            ):
            
            
                class MetaOapg:
                    
                    @staticmethod
                    def items() -> typing.Type['ChunksAndEmbeddings']:
                        return ChunksAndEmbeddings
            
                def __new__(
                    cls,
                    arg: typing.Union[typing.Tuple['ChunksAndEmbeddings'], typing.List['ChunksAndEmbeddings']],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'chunks_and_embeddings':
                    return super().__new__(
                        cls,
                        arg,
                        _configuration=_configuration,
                    )
            
                def __getitem__(self, i: int) -> 'ChunksAndEmbeddings':
                    return super().__getitem__(i)
            
            
            class chunk_size(
                schemas.IntBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneDecimalMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, decimal.Decimal, int, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'chunk_size':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class chunk_overlap(
                schemas.IntBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneDecimalMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, decimal.Decimal, int, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'chunk_overlap':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            __annotations__ = {
                "file_id": file_id,
                "chunks_and_embeddings": chunks_and_embeddings,
                "chunk_size": chunk_size,
                "chunk_overlap": chunk_overlap,
            }
    
    chunks_and_embeddings: MetaOapg.properties.chunks_and_embeddings
    file_id: MetaOapg.properties.file_id
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["file_id"]) -> MetaOapg.properties.file_id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["chunks_and_embeddings"]) -> MetaOapg.properties.chunks_and_embeddings: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["chunk_size"]) -> MetaOapg.properties.chunk_size: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["chunk_overlap"]) -> MetaOapg.properties.chunk_overlap: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["file_id", "chunks_and_embeddings", "chunk_size", "chunk_overlap", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["file_id"]) -> MetaOapg.properties.file_id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["chunks_and_embeddings"]) -> MetaOapg.properties.chunks_and_embeddings: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["chunk_size"]) -> typing.Union[MetaOapg.properties.chunk_size, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["chunk_overlap"]) -> typing.Union[MetaOapg.properties.chunk_overlap, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["file_id", "chunks_and_embeddings", "chunk_size", "chunk_overlap", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        chunks_and_embeddings: typing.Union[MetaOapg.properties.chunks_and_embeddings, list, tuple, ],
        file_id: typing.Union[MetaOapg.properties.file_id, decimal.Decimal, int, ],
        chunk_size: typing.Union[MetaOapg.properties.chunk_size, None, decimal.Decimal, int, schemas.Unset] = schemas.unset,
        chunk_overlap: typing.Union[MetaOapg.properties.chunk_overlap, None, decimal.Decimal, int, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'SingleChunksAndEmbeddingsUploadInput':
        return super().__new__(
            cls,
            *args,
            chunks_and_embeddings=chunks_and_embeddings,
            file_id=file_id,
            chunk_size=chunk_size,
            chunk_overlap=chunk_overlap,
            _configuration=_configuration,
            **kwargs,
        )

from carbon.model.chunks_and_embeddings import ChunksAndEmbeddings
