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


class UserFile(
    schemas.DictSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        required = {
            "additional_presigned_urls",
            "chunk_size",
            "embedding_properties",
            "organization_user_data_source_id",
            "sync_properties",
            "created_at",
            "external_file_id",
            "file_contents_deleted",
            "file_metadata",
            "source",
            "organization_supplied_user_id",
            "external_url",
            "source_created_at",
            "ocr_job_started_at",
            "updated_at",
            "generate_sparse_vectors",
            "enable_auto_sync",
            "id",
            "chunk_properties",
            "last_sync",
            "ocr_properties",
            "presigned_url",
            "tags",
            "sync_error_message",
            "parent_id",
            "organization_id",
            "name",
            "file_statistics",
            "sync_status",
            "parsed_text_url",
            "chunk_overlap",
            "request_id",
            "skip_embedding_generation",
            "messages_metadata",
        }
        
        class properties:
            
            
            class tags(
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
                ) -> 'tags':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                        **kwargs,
                    )
            id = schemas.IntSchema
        
            @staticmethod
            def source() -> typing.Type['DataSourceType']:
                return DataSourceType
            organization_id = schemas.IntSchema
            organization_supplied_user_id = schemas.StrSchema
            
            
            class organization_user_data_source_id(
                schemas.IntBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneDecimalMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, decimal.Decimal, int, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'organization_user_data_source_id':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            external_file_id = schemas.StrSchema
            
            
            class external_url(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'external_url':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
        
            @staticmethod
            def sync_status() -> typing.Type['ExternalFileSyncStatuses']:
                return ExternalFileSyncStatuses
            
            
            class sync_error_message(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'sync_error_message':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class last_sync(
                schemas.DateTimeBase,
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                class MetaOapg:
                    format = 'date-time'
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, datetime, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'last_sync':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
        
            @staticmethod
            def file_statistics() -> typing.Type['FileStatisticsNullable']:
                return FileStatisticsNullable
            
            
            class file_metadata(
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
                ) -> 'file_metadata':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                        **kwargs,
                    )
        
            @staticmethod
            def embedding_properties() -> typing.Type['UserFileEmbeddingProperties']:
                return UserFileEmbeddingProperties
            
            
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
        
            @staticmethod
            def chunk_properties() -> typing.Type['ChunkPropertiesNullable']:
                return ChunkPropertiesNullable
            ocr_properties = schemas.DictSchema
            
            
            class ocr_job_started_at(
                schemas.DateTimeBase,
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                class MetaOapg:
                    format = 'date-time'
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, datetime, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'ocr_job_started_at':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class name(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'name':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class parent_id(
                schemas.IntBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneDecimalMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, decimal.Decimal, int, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'parent_id':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class enable_auto_sync(
                schemas.BoolBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneBoolMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, bool, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'enable_auto_sync':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class presigned_url(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'presigned_url':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class parsed_text_url(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'parsed_text_url':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class additional_presigned_urls(
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
                ) -> 'additional_presigned_urls':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                        **kwargs,
                    )
            skip_embedding_generation = schemas.BoolSchema
            
            
            class source_created_at(
                schemas.DateTimeBase,
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                class MetaOapg:
                    format = 'date-time'
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, datetime, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'source_created_at':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class generate_sparse_vectors(
                schemas.BoolBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneBoolMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, bool, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'generate_sparse_vectors':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class request_id(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'request_id':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            sync_properties = schemas.DictSchema
            messages_metadata = schemas.DictSchema
            file_contents_deleted = schemas.BoolSchema
            created_at = schemas.DateTimeSchema
            updated_at = schemas.DateTimeSchema
            __annotations__ = {
                "tags": tags,
                "id": id,
                "source": source,
                "organization_id": organization_id,
                "organization_supplied_user_id": organization_supplied_user_id,
                "organization_user_data_source_id": organization_user_data_source_id,
                "external_file_id": external_file_id,
                "external_url": external_url,
                "sync_status": sync_status,
                "sync_error_message": sync_error_message,
                "last_sync": last_sync,
                "file_statistics": file_statistics,
                "file_metadata": file_metadata,
                "embedding_properties": embedding_properties,
                "chunk_size": chunk_size,
                "chunk_overlap": chunk_overlap,
                "chunk_properties": chunk_properties,
                "ocr_properties": ocr_properties,
                "ocr_job_started_at": ocr_job_started_at,
                "name": name,
                "parent_id": parent_id,
                "enable_auto_sync": enable_auto_sync,
                "presigned_url": presigned_url,
                "parsed_text_url": parsed_text_url,
                "additional_presigned_urls": additional_presigned_urls,
                "skip_embedding_generation": skip_embedding_generation,
                "source_created_at": source_created_at,
                "generate_sparse_vectors": generate_sparse_vectors,
                "request_id": request_id,
                "sync_properties": sync_properties,
                "messages_metadata": messages_metadata,
                "file_contents_deleted": file_contents_deleted,
                "created_at": created_at,
                "updated_at": updated_at,
            }
    
    additional_presigned_urls: MetaOapg.properties.additional_presigned_urls
    chunk_size: MetaOapg.properties.chunk_size
    embedding_properties: 'UserFileEmbeddingProperties'
    organization_user_data_source_id: MetaOapg.properties.organization_user_data_source_id
    sync_properties: MetaOapg.properties.sync_properties
    created_at: MetaOapg.properties.created_at
    external_file_id: MetaOapg.properties.external_file_id
    file_contents_deleted: MetaOapg.properties.file_contents_deleted
    file_metadata: MetaOapg.properties.file_metadata
    source: 'DataSourceType'
    organization_supplied_user_id: MetaOapg.properties.organization_supplied_user_id
    external_url: MetaOapg.properties.external_url
    source_created_at: MetaOapg.properties.source_created_at
    ocr_job_started_at: MetaOapg.properties.ocr_job_started_at
    updated_at: MetaOapg.properties.updated_at
    generate_sparse_vectors: MetaOapg.properties.generate_sparse_vectors
    enable_auto_sync: MetaOapg.properties.enable_auto_sync
    id: MetaOapg.properties.id
    chunk_properties: 'ChunkPropertiesNullable'
    last_sync: MetaOapg.properties.last_sync
    ocr_properties: MetaOapg.properties.ocr_properties
    presigned_url: MetaOapg.properties.presigned_url
    tags: MetaOapg.properties.tags
    sync_error_message: MetaOapg.properties.sync_error_message
    parent_id: MetaOapg.properties.parent_id
    organization_id: MetaOapg.properties.organization_id
    name: MetaOapg.properties.name
    file_statistics: 'FileStatisticsNullable'
    sync_status: 'ExternalFileSyncStatuses'
    parsed_text_url: MetaOapg.properties.parsed_text_url
    chunk_overlap: MetaOapg.properties.chunk_overlap
    request_id: MetaOapg.properties.request_id
    skip_embedding_generation: MetaOapg.properties.skip_embedding_generation
    messages_metadata: MetaOapg.properties.messages_metadata
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["tags"]) -> MetaOapg.properties.tags: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["id"]) -> MetaOapg.properties.id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["source"]) -> 'DataSourceType': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["organization_id"]) -> MetaOapg.properties.organization_id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["organization_supplied_user_id"]) -> MetaOapg.properties.organization_supplied_user_id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["organization_user_data_source_id"]) -> MetaOapg.properties.organization_user_data_source_id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["external_file_id"]) -> MetaOapg.properties.external_file_id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["external_url"]) -> MetaOapg.properties.external_url: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["sync_status"]) -> 'ExternalFileSyncStatuses': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["sync_error_message"]) -> MetaOapg.properties.sync_error_message: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["last_sync"]) -> MetaOapg.properties.last_sync: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["file_statistics"]) -> 'FileStatisticsNullable': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["file_metadata"]) -> MetaOapg.properties.file_metadata: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["embedding_properties"]) -> 'UserFileEmbeddingProperties': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["chunk_size"]) -> MetaOapg.properties.chunk_size: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["chunk_overlap"]) -> MetaOapg.properties.chunk_overlap: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["chunk_properties"]) -> 'ChunkPropertiesNullable': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["ocr_properties"]) -> MetaOapg.properties.ocr_properties: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["ocr_job_started_at"]) -> MetaOapg.properties.ocr_job_started_at: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["name"]) -> MetaOapg.properties.name: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["parent_id"]) -> MetaOapg.properties.parent_id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["enable_auto_sync"]) -> MetaOapg.properties.enable_auto_sync: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["presigned_url"]) -> MetaOapg.properties.presigned_url: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["parsed_text_url"]) -> MetaOapg.properties.parsed_text_url: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["additional_presigned_urls"]) -> MetaOapg.properties.additional_presigned_urls: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["skip_embedding_generation"]) -> MetaOapg.properties.skip_embedding_generation: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["source_created_at"]) -> MetaOapg.properties.source_created_at: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["generate_sparse_vectors"]) -> MetaOapg.properties.generate_sparse_vectors: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["request_id"]) -> MetaOapg.properties.request_id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["sync_properties"]) -> MetaOapg.properties.sync_properties: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["messages_metadata"]) -> MetaOapg.properties.messages_metadata: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["file_contents_deleted"]) -> MetaOapg.properties.file_contents_deleted: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["created_at"]) -> MetaOapg.properties.created_at: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["updated_at"]) -> MetaOapg.properties.updated_at: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["tags", "id", "source", "organization_id", "organization_supplied_user_id", "organization_user_data_source_id", "external_file_id", "external_url", "sync_status", "sync_error_message", "last_sync", "file_statistics", "file_metadata", "embedding_properties", "chunk_size", "chunk_overlap", "chunk_properties", "ocr_properties", "ocr_job_started_at", "name", "parent_id", "enable_auto_sync", "presigned_url", "parsed_text_url", "additional_presigned_urls", "skip_embedding_generation", "source_created_at", "generate_sparse_vectors", "request_id", "sync_properties", "messages_metadata", "file_contents_deleted", "created_at", "updated_at", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["tags"]) -> MetaOapg.properties.tags: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["id"]) -> MetaOapg.properties.id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["source"]) -> 'DataSourceType': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["organization_id"]) -> MetaOapg.properties.organization_id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["organization_supplied_user_id"]) -> MetaOapg.properties.organization_supplied_user_id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["organization_user_data_source_id"]) -> MetaOapg.properties.organization_user_data_source_id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["external_file_id"]) -> MetaOapg.properties.external_file_id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["external_url"]) -> MetaOapg.properties.external_url: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["sync_status"]) -> 'ExternalFileSyncStatuses': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["sync_error_message"]) -> MetaOapg.properties.sync_error_message: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["last_sync"]) -> MetaOapg.properties.last_sync: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["file_statistics"]) -> 'FileStatisticsNullable': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["file_metadata"]) -> MetaOapg.properties.file_metadata: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["embedding_properties"]) -> 'UserFileEmbeddingProperties': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["chunk_size"]) -> MetaOapg.properties.chunk_size: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["chunk_overlap"]) -> MetaOapg.properties.chunk_overlap: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["chunk_properties"]) -> 'ChunkPropertiesNullable': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["ocr_properties"]) -> MetaOapg.properties.ocr_properties: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["ocr_job_started_at"]) -> MetaOapg.properties.ocr_job_started_at: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["name"]) -> MetaOapg.properties.name: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["parent_id"]) -> MetaOapg.properties.parent_id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["enable_auto_sync"]) -> MetaOapg.properties.enable_auto_sync: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["presigned_url"]) -> MetaOapg.properties.presigned_url: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["parsed_text_url"]) -> MetaOapg.properties.parsed_text_url: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["additional_presigned_urls"]) -> MetaOapg.properties.additional_presigned_urls: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["skip_embedding_generation"]) -> MetaOapg.properties.skip_embedding_generation: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["source_created_at"]) -> MetaOapg.properties.source_created_at: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["generate_sparse_vectors"]) -> MetaOapg.properties.generate_sparse_vectors: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["request_id"]) -> MetaOapg.properties.request_id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["sync_properties"]) -> MetaOapg.properties.sync_properties: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["messages_metadata"]) -> MetaOapg.properties.messages_metadata: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["file_contents_deleted"]) -> MetaOapg.properties.file_contents_deleted: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["created_at"]) -> MetaOapg.properties.created_at: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["updated_at"]) -> MetaOapg.properties.updated_at: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["tags", "id", "source", "organization_id", "organization_supplied_user_id", "organization_user_data_source_id", "external_file_id", "external_url", "sync_status", "sync_error_message", "last_sync", "file_statistics", "file_metadata", "embedding_properties", "chunk_size", "chunk_overlap", "chunk_properties", "ocr_properties", "ocr_job_started_at", "name", "parent_id", "enable_auto_sync", "presigned_url", "parsed_text_url", "additional_presigned_urls", "skip_embedding_generation", "source_created_at", "generate_sparse_vectors", "request_id", "sync_properties", "messages_metadata", "file_contents_deleted", "created_at", "updated_at", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        additional_presigned_urls: typing.Union[MetaOapg.properties.additional_presigned_urls, dict, frozendict.frozendict, None, ],
        chunk_size: typing.Union[MetaOapg.properties.chunk_size, None, decimal.Decimal, int, ],
        embedding_properties: 'UserFileEmbeddingProperties',
        organization_user_data_source_id: typing.Union[MetaOapg.properties.organization_user_data_source_id, None, decimal.Decimal, int, ],
        sync_properties: typing.Union[MetaOapg.properties.sync_properties, dict, frozendict.frozendict, ],
        created_at: typing.Union[MetaOapg.properties.created_at, str, datetime, ],
        external_file_id: typing.Union[MetaOapg.properties.external_file_id, str, ],
        file_contents_deleted: typing.Union[MetaOapg.properties.file_contents_deleted, bool, ],
        file_metadata: typing.Union[MetaOapg.properties.file_metadata, dict, frozendict.frozendict, None, ],
        source: 'DataSourceType',
        organization_supplied_user_id: typing.Union[MetaOapg.properties.organization_supplied_user_id, str, ],
        external_url: typing.Union[MetaOapg.properties.external_url, None, str, ],
        source_created_at: typing.Union[MetaOapg.properties.source_created_at, None, str, datetime, ],
        ocr_job_started_at: typing.Union[MetaOapg.properties.ocr_job_started_at, None, str, datetime, ],
        updated_at: typing.Union[MetaOapg.properties.updated_at, str, datetime, ],
        generate_sparse_vectors: typing.Union[MetaOapg.properties.generate_sparse_vectors, None, bool, ],
        enable_auto_sync: typing.Union[MetaOapg.properties.enable_auto_sync, None, bool, ],
        id: typing.Union[MetaOapg.properties.id, decimal.Decimal, int, ],
        chunk_properties: 'ChunkPropertiesNullable',
        last_sync: typing.Union[MetaOapg.properties.last_sync, None, str, datetime, ],
        ocr_properties: typing.Union[MetaOapg.properties.ocr_properties, dict, frozendict.frozendict, ],
        presigned_url: typing.Union[MetaOapg.properties.presigned_url, None, str, ],
        tags: typing.Union[MetaOapg.properties.tags, dict, frozendict.frozendict, None, ],
        sync_error_message: typing.Union[MetaOapg.properties.sync_error_message, None, str, ],
        parent_id: typing.Union[MetaOapg.properties.parent_id, None, decimal.Decimal, int, ],
        organization_id: typing.Union[MetaOapg.properties.organization_id, decimal.Decimal, int, ],
        name: typing.Union[MetaOapg.properties.name, None, str, ],
        file_statistics: 'FileStatisticsNullable',
        sync_status: 'ExternalFileSyncStatuses',
        parsed_text_url: typing.Union[MetaOapg.properties.parsed_text_url, None, str, ],
        chunk_overlap: typing.Union[MetaOapg.properties.chunk_overlap, None, decimal.Decimal, int, ],
        request_id: typing.Union[MetaOapg.properties.request_id, None, str, ],
        skip_embedding_generation: typing.Union[MetaOapg.properties.skip_embedding_generation, bool, ],
        messages_metadata: typing.Union[MetaOapg.properties.messages_metadata, dict, frozendict.frozendict, ],
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'UserFile':
        return super().__new__(
            cls,
            *args,
            additional_presigned_urls=additional_presigned_urls,
            chunk_size=chunk_size,
            embedding_properties=embedding_properties,
            organization_user_data_source_id=organization_user_data_source_id,
            sync_properties=sync_properties,
            created_at=created_at,
            external_file_id=external_file_id,
            file_contents_deleted=file_contents_deleted,
            file_metadata=file_metadata,
            source=source,
            organization_supplied_user_id=organization_supplied_user_id,
            external_url=external_url,
            source_created_at=source_created_at,
            ocr_job_started_at=ocr_job_started_at,
            updated_at=updated_at,
            generate_sparse_vectors=generate_sparse_vectors,
            enable_auto_sync=enable_auto_sync,
            id=id,
            chunk_properties=chunk_properties,
            last_sync=last_sync,
            ocr_properties=ocr_properties,
            presigned_url=presigned_url,
            tags=tags,
            sync_error_message=sync_error_message,
            parent_id=parent_id,
            organization_id=organization_id,
            name=name,
            file_statistics=file_statistics,
            sync_status=sync_status,
            parsed_text_url=parsed_text_url,
            chunk_overlap=chunk_overlap,
            request_id=request_id,
            skip_embedding_generation=skip_embedding_generation,
            messages_metadata=messages_metadata,
            _configuration=_configuration,
            **kwargs,
        )

from carbon.model.chunk_properties_nullable import ChunkPropertiesNullable
from carbon.model.data_source_type import DataSourceType
from carbon.model.external_file_sync_statuses import ExternalFileSyncStatuses
from carbon.model.file_statistics_nullable import FileStatisticsNullable
from carbon.model.user_file_embedding_properties import UserFileEmbeddingProperties
