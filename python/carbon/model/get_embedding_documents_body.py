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


class GetEmbeddingDocumentsBody(
    schemas.DictSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        required = {
            "query",
            "k",
        }
        
        class properties:
            
            
            class query(
                schemas.StrSchema
            ):
            
            
                class MetaOapg:
                    min_length = 1
            k = schemas.IntSchema
        
            @staticmethod
            def tags() -> typing.Type['GetEmbeddingDocumentsBodyTags']:
                return GetEmbeddingDocumentsBodyTags
        
            @staticmethod
            def query_vector() -> typing.Type['GetEmbeddingDocumentsBodyQueryVector']:
                return GetEmbeddingDocumentsBodyQueryVector
        
            @staticmethod
            def file_ids() -> typing.Type['GetEmbeddingDocumentsBodyFileIds']:
                return GetEmbeddingDocumentsBodyFileIds
        
            @staticmethod
            def parent_file_ids() -> typing.Type['GetEmbeddingDocumentsBodyParentFileIds']:
                return GetEmbeddingDocumentsBodyParentFileIds
            
            
            class tags_v2(
                schemas.DictBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneFrozenDictMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[dict, frozendict.frozendict, None, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                    **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
                ) -> 'tags_v2':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                        **kwargs,
                    )
            
            
            class include_tags(
                schemas.BoolBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneBoolMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, bool, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'include_tags':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class include_vectors(
                schemas.BoolBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneBoolMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, bool, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'include_vectors':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class include_raw_file(
                schemas.BoolBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneBoolMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, bool, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'include_raw_file':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class hybrid_search(
                schemas.BoolBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneBoolMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, bool, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'hybrid_search':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
        
            @staticmethod
            def hybrid_search_tuning_parameters() -> typing.Type['HybridSearchTuningParamsNullable']:
                return HybridSearchTuningParamsNullable
        
            @staticmethod
            def media_type() -> typing.Type['FileContentTypesNullable']:
                return FileContentTypesNullable
        
            @staticmethod
            def embedding_model() -> typing.Type['EmbeddingGeneratorsNullable']:
                return EmbeddingGeneratorsNullable
            __annotations__ = {
                "query": query,
                "k": k,
                "tags": tags,
                "query_vector": query_vector,
                "file_ids": file_ids,
                "parent_file_ids": parent_file_ids,
                "tags_v2": tags_v2,
                "include_tags": include_tags,
                "include_vectors": include_vectors,
                "include_raw_file": include_raw_file,
                "hybrid_search": hybrid_search,
                "hybrid_search_tuning_parameters": hybrid_search_tuning_parameters,
                "media_type": media_type,
                "embedding_model": embedding_model,
            }
    
    query: MetaOapg.properties.query
    k: MetaOapg.properties.k
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["query"]) -> MetaOapg.properties.query: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["k"]) -> MetaOapg.properties.k: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["tags"]) -> 'GetEmbeddingDocumentsBodyTags': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["query_vector"]) -> 'GetEmbeddingDocumentsBodyQueryVector': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["file_ids"]) -> 'GetEmbeddingDocumentsBodyFileIds': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["parent_file_ids"]) -> 'GetEmbeddingDocumentsBodyParentFileIds': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["tags_v2"]) -> MetaOapg.properties.tags_v2: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["include_tags"]) -> MetaOapg.properties.include_tags: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["include_vectors"]) -> MetaOapg.properties.include_vectors: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["include_raw_file"]) -> MetaOapg.properties.include_raw_file: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["hybrid_search"]) -> MetaOapg.properties.hybrid_search: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["hybrid_search_tuning_parameters"]) -> 'HybridSearchTuningParamsNullable': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["media_type"]) -> 'FileContentTypesNullable': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["embedding_model"]) -> 'EmbeddingGeneratorsNullable': ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["query", "k", "tags", "query_vector", "file_ids", "parent_file_ids", "tags_v2", "include_tags", "include_vectors", "include_raw_file", "hybrid_search", "hybrid_search_tuning_parameters", "media_type", "embedding_model", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["query"]) -> MetaOapg.properties.query: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["k"]) -> MetaOapg.properties.k: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["tags"]) -> typing.Union['GetEmbeddingDocumentsBodyTags', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["query_vector"]) -> typing.Union['GetEmbeddingDocumentsBodyQueryVector', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["file_ids"]) -> typing.Union['GetEmbeddingDocumentsBodyFileIds', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["parent_file_ids"]) -> typing.Union['GetEmbeddingDocumentsBodyParentFileIds', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["tags_v2"]) -> typing.Union[MetaOapg.properties.tags_v2, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["include_tags"]) -> typing.Union[MetaOapg.properties.include_tags, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["include_vectors"]) -> typing.Union[MetaOapg.properties.include_vectors, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["include_raw_file"]) -> typing.Union[MetaOapg.properties.include_raw_file, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["hybrid_search"]) -> typing.Union[MetaOapg.properties.hybrid_search, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["hybrid_search_tuning_parameters"]) -> typing.Union['HybridSearchTuningParamsNullable', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["media_type"]) -> typing.Union['FileContentTypesNullable', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["embedding_model"]) -> typing.Union['EmbeddingGeneratorsNullable', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["query", "k", "tags", "query_vector", "file_ids", "parent_file_ids", "tags_v2", "include_tags", "include_vectors", "include_raw_file", "hybrid_search", "hybrid_search_tuning_parameters", "media_type", "embedding_model", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        query: typing.Union[MetaOapg.properties.query, str, ],
        k: typing.Union[MetaOapg.properties.k, decimal.Decimal, int, ],
        tags: typing.Union['GetEmbeddingDocumentsBodyTags', schemas.Unset] = schemas.unset,
        query_vector: typing.Union['GetEmbeddingDocumentsBodyQueryVector', schemas.Unset] = schemas.unset,
        file_ids: typing.Union['GetEmbeddingDocumentsBodyFileIds', schemas.Unset] = schemas.unset,
        parent_file_ids: typing.Union['GetEmbeddingDocumentsBodyParentFileIds', schemas.Unset] = schemas.unset,
        tags_v2: typing.Union[MetaOapg.properties.tags_v2, dict, frozendict.frozendict, None, schemas.Unset] = schemas.unset,
        include_tags: typing.Union[MetaOapg.properties.include_tags, None, bool, schemas.Unset] = schemas.unset,
        include_vectors: typing.Union[MetaOapg.properties.include_vectors, None, bool, schemas.Unset] = schemas.unset,
        include_raw_file: typing.Union[MetaOapg.properties.include_raw_file, None, bool, schemas.Unset] = schemas.unset,
        hybrid_search: typing.Union[MetaOapg.properties.hybrid_search, None, bool, schemas.Unset] = schemas.unset,
        hybrid_search_tuning_parameters: typing.Union['HybridSearchTuningParamsNullable', schemas.Unset] = schemas.unset,
        media_type: typing.Union['FileContentTypesNullable', schemas.Unset] = schemas.unset,
        embedding_model: typing.Union['EmbeddingGeneratorsNullable', schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'GetEmbeddingDocumentsBody':
        return super().__new__(
            cls,
            *args,
            query=query,
            k=k,
            tags=tags,
            query_vector=query_vector,
            file_ids=file_ids,
            parent_file_ids=parent_file_ids,
            tags_v2=tags_v2,
            include_tags=include_tags,
            include_vectors=include_vectors,
            include_raw_file=include_raw_file,
            hybrid_search=hybrid_search,
            hybrid_search_tuning_parameters=hybrid_search_tuning_parameters,
            media_type=media_type,
            embedding_model=embedding_model,
            _configuration=_configuration,
            **kwargs,
        )

from carbon.model.embedding_generators_nullable import EmbeddingGeneratorsNullable
from carbon.model.file_content_types_nullable import FileContentTypesNullable
from carbon.model.get_embedding_documents_body_file_ids import GetEmbeddingDocumentsBodyFileIds
from carbon.model.get_embedding_documents_body_parent_file_ids import GetEmbeddingDocumentsBodyParentFileIds
from carbon.model.get_embedding_documents_body_query_vector import GetEmbeddingDocumentsBodyQueryVector
from carbon.model.get_embedding_documents_body_tags import GetEmbeddingDocumentsBodyTags
from carbon.model.hybrid_search_tuning_params_nullable import HybridSearchTuningParamsNullable
