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
            "updated_at",
            "num_tokens_synced",
            "organization_id",
            "unique_file_tags",
            "created_at",
            "id",
            "deleted_at",
            "organization_supplied_user_id",
            "enabled_features",
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
                "unique_file_tags": unique_file_tags,
                "enabled_features": enabled_features,
            }
    
    num_characters_synced: MetaOapg.properties.num_characters_synced
    updated_at: MetaOapg.properties.updated_at
    num_tokens_synced: MetaOapg.properties.num_tokens_synced
    organization_id: MetaOapg.properties.organization_id
    unique_file_tags: 'UserResponseUniqueFileTags'
    created_at: MetaOapg.properties.created_at
    id: MetaOapg.properties.id
    deleted_at: MetaOapg.properties.deleted_at
    organization_supplied_user_id: MetaOapg.properties.organization_supplied_user_id
    enabled_features: MetaOapg.properties.enabled_features
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
    def __getitem__(self, name: typing_extensions.Literal["unique_file_tags"]) -> 'UserResponseUniqueFileTags': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["enabled_features"]) -> MetaOapg.properties.enabled_features: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["id", "organization_id", "organization_supplied_user_id", "created_at", "updated_at", "deleted_at", "num_files_synced", "num_characters_synced", "num_tokens_synced", "unique_file_tags", "enabled_features", ], str]):
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
    def get_item_oapg(self, name: typing_extensions.Literal["unique_file_tags"]) -> 'UserResponseUniqueFileTags': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["enabled_features"]) -> MetaOapg.properties.enabled_features: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["id", "organization_id", "organization_supplied_user_id", "created_at", "updated_at", "deleted_at", "num_files_synced", "num_characters_synced", "num_tokens_synced", "unique_file_tags", "enabled_features", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        num_characters_synced: typing.Union[MetaOapg.properties.num_characters_synced, decimal.Decimal, int, ],
        updated_at: typing.Union[MetaOapg.properties.updated_at, str, datetime, ],
        num_tokens_synced: typing.Union[MetaOapg.properties.num_tokens_synced, decimal.Decimal, int, ],
        organization_id: typing.Union[MetaOapg.properties.organization_id, decimal.Decimal, int, ],
        unique_file_tags: 'UserResponseUniqueFileTags',
        created_at: typing.Union[MetaOapg.properties.created_at, str, datetime, ],
        id: typing.Union[MetaOapg.properties.id, decimal.Decimal, int, ],
        deleted_at: typing.Union[MetaOapg.properties.deleted_at, None, str, datetime, ],
        organization_supplied_user_id: typing.Union[MetaOapg.properties.organization_supplied_user_id, str, ],
        enabled_features: typing.Union[MetaOapg.properties.enabled_features, dict, frozendict.frozendict, None, ],
        num_files_synced: typing.Union[MetaOapg.properties.num_files_synced, decimal.Decimal, int, ],
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'UserResponse':
        return super().__new__(
            cls,
            *args,
            num_characters_synced=num_characters_synced,
            updated_at=updated_at,
            num_tokens_synced=num_tokens_synced,
            organization_id=organization_id,
            unique_file_tags=unique_file_tags,
            created_at=created_at,
            id=id,
            deleted_at=deleted_at,
            organization_supplied_user_id=organization_supplied_user_id,
            enabled_features=enabled_features,
            num_files_synced=num_files_synced,
            _configuration=_configuration,
            **kwargs,
        )

from carbon.model.user_response_unique_file_tags import UserResponseUniqueFileTags
