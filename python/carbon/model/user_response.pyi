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


class UserResponse(
    schemas.DictSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        required = {
            "num_characters_synced",
            "aggregate_file_size",
            "aggregate_num_characters",
            "aggregate_num_files_by_file_format",
            "custom_limits",
            "created_at",
            "deleted_at",
            "organization_supplied_user_id",
            "enabled_features",
            "updated_at",
            "aggregate_num_embeddings",
            "aggregate_num_files_by_source",
            "num_tokens_synced",
            "organization_id",
            "unique_file_tags",
            "id",
            "aggregate_num_tokens",
            "auto_sync_enabled_sources",
            "num_files_synced",
        }
        
        class properties:
            id = schemas.IntSchema
            organization_id = schemas.IntSchema
            organization_supplied_user_id = schemas.StrSchema
            created_at = schemas.DateTimeSchema
            updated_at = schemas.DateTimeSchema
            
            
            class deleted_at(
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
                ) -> 'deleted_at':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            num_files_synced = schemas.IntSchema
            num_characters_synced = schemas.IntSchema
            num_tokens_synced = schemas.IntSchema
            aggregate_file_size = schemas.DictSchema
            aggregate_num_characters = schemas.DictSchema
            aggregate_num_tokens = schemas.DictSchema
            aggregate_num_embeddings = schemas.DictSchema
            aggregate_num_files_by_source = schemas.DictSchema
            aggregate_num_files_by_file_format = schemas.DictSchema
        
            @staticmethod
            def unique_file_tags() -> typing.Type['UserResponseUniqueFileTags']:
                return UserResponseUniqueFileTags
            
            
            class enabled_features(
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
                ) -> 'enabled_features':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                        **kwargs,
                    )
            custom_limits = schemas.DictSchema
        
            @staticmethod
            def auto_sync_enabled_sources() -> typing.Type['UserResponseAutoSyncEnabledSources']:
                return UserResponseAutoSyncEnabledSources
            __annotations__ = {
                "id": id,
                "organization_id": organization_id,
                "organization_supplied_user_id": organization_supplied_user_id,
                "created_at": created_at,
                "updated_at": updated_at,
                "deleted_at": deleted_at,
                "num_files_synced": num_files_synced,
                "num_characters_synced": num_characters_synced,
                "num_tokens_synced": num_tokens_synced,
                "aggregate_file_size": aggregate_file_size,
                "aggregate_num_characters": aggregate_num_characters,
                "aggregate_num_tokens": aggregate_num_tokens,
                "aggregate_num_embeddings": aggregate_num_embeddings,
                "aggregate_num_files_by_source": aggregate_num_files_by_source,
                "aggregate_num_files_by_file_format": aggregate_num_files_by_file_format,
                "unique_file_tags": unique_file_tags,
                "enabled_features": enabled_features,
                "custom_limits": custom_limits,
                "auto_sync_enabled_sources": auto_sync_enabled_sources,
            }
    
    num_characters_synced: MetaOapg.properties.num_characters_synced
    aggregate_file_size: MetaOapg.properties.aggregate_file_size
    aggregate_num_characters: MetaOapg.properties.aggregate_num_characters
    aggregate_num_files_by_file_format: MetaOapg.properties.aggregate_num_files_by_file_format
    custom_limits: MetaOapg.properties.custom_limits
    created_at: MetaOapg.properties.created_at
    deleted_at: MetaOapg.properties.deleted_at
    organization_supplied_user_id: MetaOapg.properties.organization_supplied_user_id
    enabled_features: MetaOapg.properties.enabled_features
    updated_at: MetaOapg.properties.updated_at
    aggregate_num_embeddings: MetaOapg.properties.aggregate_num_embeddings
    aggregate_num_files_by_source: MetaOapg.properties.aggregate_num_files_by_source
    num_tokens_synced: MetaOapg.properties.num_tokens_synced
    organization_id: MetaOapg.properties.organization_id
    unique_file_tags: 'UserResponseUniqueFileTags'
    id: MetaOapg.properties.id
    aggregate_num_tokens: MetaOapg.properties.aggregate_num_tokens
    auto_sync_enabled_sources: 'UserResponseAutoSyncEnabledSources'
    num_files_synced: MetaOapg.properties.num_files_synced
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["id"]) -> MetaOapg.properties.id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["organization_id"]) -> MetaOapg.properties.organization_id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["organization_supplied_user_id"]) -> MetaOapg.properties.organization_supplied_user_id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["created_at"]) -> MetaOapg.properties.created_at: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["updated_at"]) -> MetaOapg.properties.updated_at: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["deleted_at"]) -> MetaOapg.properties.deleted_at: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["num_files_synced"]) -> MetaOapg.properties.num_files_synced: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["num_characters_synced"]) -> MetaOapg.properties.num_characters_synced: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["num_tokens_synced"]) -> MetaOapg.properties.num_tokens_synced: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["aggregate_file_size"]) -> MetaOapg.properties.aggregate_file_size: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["aggregate_num_characters"]) -> MetaOapg.properties.aggregate_num_characters: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["aggregate_num_tokens"]) -> MetaOapg.properties.aggregate_num_tokens: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["aggregate_num_embeddings"]) -> MetaOapg.properties.aggregate_num_embeddings: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["aggregate_num_files_by_source"]) -> MetaOapg.properties.aggregate_num_files_by_source: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["aggregate_num_files_by_file_format"]) -> MetaOapg.properties.aggregate_num_files_by_file_format: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["unique_file_tags"]) -> 'UserResponseUniqueFileTags': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["enabled_features"]) -> MetaOapg.properties.enabled_features: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["custom_limits"]) -> MetaOapg.properties.custom_limits: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["auto_sync_enabled_sources"]) -> 'UserResponseAutoSyncEnabledSources': ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["id", "organization_id", "organization_supplied_user_id", "created_at", "updated_at", "deleted_at", "num_files_synced", "num_characters_synced", "num_tokens_synced", "aggregate_file_size", "aggregate_num_characters", "aggregate_num_tokens", "aggregate_num_embeddings", "aggregate_num_files_by_source", "aggregate_num_files_by_file_format", "unique_file_tags", "enabled_features", "custom_limits", "auto_sync_enabled_sources", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["id"]) -> MetaOapg.properties.id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["organization_id"]) -> MetaOapg.properties.organization_id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["organization_supplied_user_id"]) -> MetaOapg.properties.organization_supplied_user_id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["created_at"]) -> MetaOapg.properties.created_at: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["updated_at"]) -> MetaOapg.properties.updated_at: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["deleted_at"]) -> MetaOapg.properties.deleted_at: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["num_files_synced"]) -> MetaOapg.properties.num_files_synced: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["num_characters_synced"]) -> MetaOapg.properties.num_characters_synced: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["num_tokens_synced"]) -> MetaOapg.properties.num_tokens_synced: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["aggregate_file_size"]) -> MetaOapg.properties.aggregate_file_size: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["aggregate_num_characters"]) -> MetaOapg.properties.aggregate_num_characters: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["aggregate_num_tokens"]) -> MetaOapg.properties.aggregate_num_tokens: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["aggregate_num_embeddings"]) -> MetaOapg.properties.aggregate_num_embeddings: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["aggregate_num_files_by_source"]) -> MetaOapg.properties.aggregate_num_files_by_source: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["aggregate_num_files_by_file_format"]) -> MetaOapg.properties.aggregate_num_files_by_file_format: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["unique_file_tags"]) -> 'UserResponseUniqueFileTags': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["enabled_features"]) -> MetaOapg.properties.enabled_features: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["custom_limits"]) -> MetaOapg.properties.custom_limits: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["auto_sync_enabled_sources"]) -> 'UserResponseAutoSyncEnabledSources': ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["id", "organization_id", "organization_supplied_user_id", "created_at", "updated_at", "deleted_at", "num_files_synced", "num_characters_synced", "num_tokens_synced", "aggregate_file_size", "aggregate_num_characters", "aggregate_num_tokens", "aggregate_num_embeddings", "aggregate_num_files_by_source", "aggregate_num_files_by_file_format", "unique_file_tags", "enabled_features", "custom_limits", "auto_sync_enabled_sources", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        num_characters_synced: typing.Union[MetaOapg.properties.num_characters_synced, decimal.Decimal, int, ],
        aggregate_file_size: typing.Union[MetaOapg.properties.aggregate_file_size, dict, frozendict.frozendict, ],
        aggregate_num_characters: typing.Union[MetaOapg.properties.aggregate_num_characters, dict, frozendict.frozendict, ],
        aggregate_num_files_by_file_format: typing.Union[MetaOapg.properties.aggregate_num_files_by_file_format, dict, frozendict.frozendict, ],
        custom_limits: typing.Union[MetaOapg.properties.custom_limits, dict, frozendict.frozendict, ],
        created_at: typing.Union[MetaOapg.properties.created_at, str, datetime, ],
        deleted_at: typing.Union[MetaOapg.properties.deleted_at, None, str, datetime, ],
        organization_supplied_user_id: typing.Union[MetaOapg.properties.organization_supplied_user_id, str, ],
        enabled_features: typing.Union[MetaOapg.properties.enabled_features, dict, frozendict.frozendict, None, ],
        updated_at: typing.Union[MetaOapg.properties.updated_at, str, datetime, ],
        aggregate_num_embeddings: typing.Union[MetaOapg.properties.aggregate_num_embeddings, dict, frozendict.frozendict, ],
        aggregate_num_files_by_source: typing.Union[MetaOapg.properties.aggregate_num_files_by_source, dict, frozendict.frozendict, ],
        num_tokens_synced: typing.Union[MetaOapg.properties.num_tokens_synced, decimal.Decimal, int, ],
        organization_id: typing.Union[MetaOapg.properties.organization_id, decimal.Decimal, int, ],
        unique_file_tags: 'UserResponseUniqueFileTags',
        id: typing.Union[MetaOapg.properties.id, decimal.Decimal, int, ],
        aggregate_num_tokens: typing.Union[MetaOapg.properties.aggregate_num_tokens, dict, frozendict.frozendict, ],
        auto_sync_enabled_sources: 'UserResponseAutoSyncEnabledSources',
        num_files_synced: typing.Union[MetaOapg.properties.num_files_synced, decimal.Decimal, int, ],
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'UserResponse':
        return super().__new__(
            cls,
            *args,
            num_characters_synced=num_characters_synced,
            aggregate_file_size=aggregate_file_size,
            aggregate_num_characters=aggregate_num_characters,
            aggregate_num_files_by_file_format=aggregate_num_files_by_file_format,
            custom_limits=custom_limits,
            created_at=created_at,
            deleted_at=deleted_at,
            organization_supplied_user_id=organization_supplied_user_id,
            enabled_features=enabled_features,
            updated_at=updated_at,
            aggregate_num_embeddings=aggregate_num_embeddings,
            aggregate_num_files_by_source=aggregate_num_files_by_source,
            num_tokens_synced=num_tokens_synced,
            organization_id=organization_id,
            unique_file_tags=unique_file_tags,
            id=id,
            aggregate_num_tokens=aggregate_num_tokens,
            auto_sync_enabled_sources=auto_sync_enabled_sources,
            num_files_synced=num_files_synced,
            _configuration=_configuration,
            **kwargs,
        )

from carbon.model.user_response_auto_sync_enabled_sources import UserResponseAutoSyncEnabledSources
from carbon.model.user_response_unique_file_tags import UserResponseUniqueFileTags
