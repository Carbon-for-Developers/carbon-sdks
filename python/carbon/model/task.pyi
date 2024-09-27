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


class Task(
    schemas.DictSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        required = {
            "owner",
            "is_deleted",
            "updated_at",
            "subject",
            "contact",
            "created_at",
            "description",
            "id",
            "priority",
            "account",
            "remote_data",
            "status",
        }
        
        class properties:
            
            
            class description(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'description':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            id = schemas.StrSchema
        
            @staticmethod
            def owner() -> typing.Type['PartialOwner']:
                return PartialOwner
            
            
            class subject(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'subject':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class status(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'status':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class priority(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'priority':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
        
            @staticmethod
            def account() -> typing.Type['PartialAccountNullable']:
                return PartialAccountNullable
        
            @staticmethod
            def contact() -> typing.Type['PartialContactNullable']:
                return PartialContactNullable
            created_at = schemas.StrSchema
            updated_at = schemas.StrSchema
            is_deleted = schemas.BoolSchema
            
            
            class remote_data(
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
                ) -> 'remote_data':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                        **kwargs,
                    )
            __annotations__ = {
                "description": description,
                "id": id,
                "owner": owner,
                "subject": subject,
                "status": status,
                "priority": priority,
                "account": account,
                "contact": contact,
                "created_at": created_at,
                "updated_at": updated_at,
                "is_deleted": is_deleted,
                "remote_data": remote_data,
            }
    
    owner: 'PartialOwner'
    is_deleted: MetaOapg.properties.is_deleted
    updated_at: MetaOapg.properties.updated_at
    subject: MetaOapg.properties.subject
    contact: 'PartialContactNullable'
    created_at: MetaOapg.properties.created_at
    description: MetaOapg.properties.description
    id: MetaOapg.properties.id
    priority: MetaOapg.properties.priority
    account: 'PartialAccountNullable'
    remote_data: MetaOapg.properties.remote_data
    status: MetaOapg.properties.status
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["description"]) -> MetaOapg.properties.description: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["id"]) -> MetaOapg.properties.id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["owner"]) -> 'PartialOwner': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["subject"]) -> MetaOapg.properties.subject: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["status"]) -> MetaOapg.properties.status: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["priority"]) -> MetaOapg.properties.priority: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["account"]) -> 'PartialAccountNullable': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["contact"]) -> 'PartialContactNullable': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["created_at"]) -> MetaOapg.properties.created_at: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["updated_at"]) -> MetaOapg.properties.updated_at: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["is_deleted"]) -> MetaOapg.properties.is_deleted: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["remote_data"]) -> MetaOapg.properties.remote_data: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["description", "id", "owner", "subject", "status", "priority", "account", "contact", "created_at", "updated_at", "is_deleted", "remote_data", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["description"]) -> MetaOapg.properties.description: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["id"]) -> MetaOapg.properties.id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["owner"]) -> 'PartialOwner': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["subject"]) -> MetaOapg.properties.subject: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["status"]) -> MetaOapg.properties.status: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["priority"]) -> MetaOapg.properties.priority: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["account"]) -> 'PartialAccountNullable': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["contact"]) -> 'PartialContactNullable': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["created_at"]) -> MetaOapg.properties.created_at: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["updated_at"]) -> MetaOapg.properties.updated_at: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["is_deleted"]) -> MetaOapg.properties.is_deleted: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["remote_data"]) -> MetaOapg.properties.remote_data: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["description", "id", "owner", "subject", "status", "priority", "account", "contact", "created_at", "updated_at", "is_deleted", "remote_data", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        owner: 'PartialOwner',
        is_deleted: typing.Union[MetaOapg.properties.is_deleted, bool, ],
        updated_at: typing.Union[MetaOapg.properties.updated_at, str, ],
        subject: typing.Union[MetaOapg.properties.subject, None, str, ],
        contact: 'PartialContactNullable',
        created_at: typing.Union[MetaOapg.properties.created_at, str, ],
        description: typing.Union[MetaOapg.properties.description, None, str, ],
        id: typing.Union[MetaOapg.properties.id, str, ],
        priority: typing.Union[MetaOapg.properties.priority, None, str, ],
        account: 'PartialAccountNullable',
        remote_data: typing.Union[MetaOapg.properties.remote_data, dict, frozendict.frozendict, None, ],
        status: typing.Union[MetaOapg.properties.status, None, str, ],
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'Task':
        return super().__new__(
            cls,
            *args,
            owner=owner,
            is_deleted=is_deleted,
            updated_at=updated_at,
            subject=subject,
            contact=contact,
            created_at=created_at,
            description=description,
            id=id,
            priority=priority,
            account=account,
            remote_data=remote_data,
            status=status,
            _configuration=_configuration,
            **kwargs,
        )

from carbon.model.partial_account_nullable import PartialAccountNullable
from carbon.model.partial_contact_nullable import PartialContactNullable
from carbon.model.partial_owner import PartialOwner
