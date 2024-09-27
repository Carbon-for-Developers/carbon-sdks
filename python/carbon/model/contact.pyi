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


class Contact(
    schemas.DictSchema
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        required = {
            "owner",
            "addresses",
            "created_at",
            "description",
            "last_name",
            "title",
            "remote_data",
            "emails",
            "phone_numbers",
            "is_deleted",
            "updated_at",
            "last_activity_at",
            "name",
            "id",
            "department",
            "first_name",
            "account",
            "events",
            "tasks",
        }
        
        class properties:
            
            
            class title(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'title':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
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
            def owner() -> typing.Type['PartialOwnerNullable']:
                return PartialOwnerNullable
            
            
            class first_name(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'first_name':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class last_name(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'last_name':
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
            
            
            class department(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'department':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class addresses(
                schemas.ListSchema
            ):
            
            
                class MetaOapg:
                    
                    @staticmethod
                    def items() -> typing.Type['Address']:
                        return Address
            
                def __new__(
                    cls,
                    arg: typing.Union[typing.Tuple['Address'], typing.List['Address']],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'addresses':
                    return super().__new__(
                        cls,
                        arg,
                        _configuration=_configuration,
                    )
            
                def __getitem__(self, i: int) -> 'Address':
                    return super().__getitem__(i)
            
            
            class phone_numbers(
                schemas.ListSchema
            ):
            
            
                class MetaOapg:
                    
                    @staticmethod
                    def items() -> typing.Type['PhoneNumber']:
                        return PhoneNumber
            
                def __new__(
                    cls,
                    arg: typing.Union[typing.Tuple['PhoneNumber'], typing.List['PhoneNumber']],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'phone_numbers':
                    return super().__new__(
                        cls,
                        arg,
                        _configuration=_configuration,
                    )
            
                def __getitem__(self, i: int) -> 'PhoneNumber':
                    return super().__getitem__(i)
            
            
            class emails(
                schemas.ListSchema
            ):
            
            
                class MetaOapg:
                    
                    @staticmethod
                    def items() -> typing.Type['Email']:
                        return Email
            
                def __new__(
                    cls,
                    arg: typing.Union[typing.Tuple['Email'], typing.List['Email']],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'emails':
                    return super().__new__(
                        cls,
                        arg,
                        _configuration=_configuration,
                    )
            
                def __getitem__(self, i: int) -> 'Email':
                    return super().__getitem__(i)
        
            @staticmethod
            def account() -> typing.Type['PartialAccountNullable']:
                return PartialAccountNullable
            
            
            class last_activity_at(
                schemas.StrBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneStrMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, str, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'last_activity_at':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            created_at = schemas.StrSchema
            updated_at = schemas.StrSchema
            is_deleted = schemas.BoolSchema
            
            
            class tasks(
                schemas.ListBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneTupleMixin
            ):
            
            
                class MetaOapg:
                    
                    @staticmethod
                    def items() -> typing.Type['Task']:
                        return Task
            
            
                def __new__(
                    cls,
                    *args: typing.Union[list, tuple, None, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'tasks':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class events(
                schemas.ListBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneTupleMixin
            ):
            
            
                class MetaOapg:
                    
                    @staticmethod
                    def items() -> typing.Type['Event']:
                        return Event
            
            
                def __new__(
                    cls,
                    *args: typing.Union[list, tuple, None, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'events':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
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
                "title": title,
                "description": description,
                "id": id,
                "owner": owner,
                "first_name": first_name,
                "last_name": last_name,
                "name": name,
                "department": department,
                "addresses": addresses,
                "phone_numbers": phone_numbers,
                "emails": emails,
                "account": account,
                "last_activity_at": last_activity_at,
                "created_at": created_at,
                "updated_at": updated_at,
                "is_deleted": is_deleted,
                "tasks": tasks,
                "events": events,
                "remote_data": remote_data,
            }
    
    owner: 'PartialOwnerNullable'
    addresses: MetaOapg.properties.addresses
    created_at: MetaOapg.properties.created_at
    description: MetaOapg.properties.description
    last_name: MetaOapg.properties.last_name
    title: MetaOapg.properties.title
    remote_data: MetaOapg.properties.remote_data
    emails: MetaOapg.properties.emails
    phone_numbers: MetaOapg.properties.phone_numbers
    is_deleted: MetaOapg.properties.is_deleted
    updated_at: MetaOapg.properties.updated_at
    last_activity_at: MetaOapg.properties.last_activity_at
    name: MetaOapg.properties.name
    id: MetaOapg.properties.id
    department: MetaOapg.properties.department
    first_name: MetaOapg.properties.first_name
    account: 'PartialAccountNullable'
    events: MetaOapg.properties.events
    tasks: MetaOapg.properties.tasks
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["title"]) -> MetaOapg.properties.title: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["description"]) -> MetaOapg.properties.description: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["id"]) -> MetaOapg.properties.id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["owner"]) -> 'PartialOwnerNullable': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["first_name"]) -> MetaOapg.properties.first_name: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["last_name"]) -> MetaOapg.properties.last_name: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["name"]) -> MetaOapg.properties.name: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["department"]) -> MetaOapg.properties.department: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["addresses"]) -> MetaOapg.properties.addresses: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["phone_numbers"]) -> MetaOapg.properties.phone_numbers: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["emails"]) -> MetaOapg.properties.emails: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["account"]) -> 'PartialAccountNullable': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["last_activity_at"]) -> MetaOapg.properties.last_activity_at: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["created_at"]) -> MetaOapg.properties.created_at: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["updated_at"]) -> MetaOapg.properties.updated_at: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["is_deleted"]) -> MetaOapg.properties.is_deleted: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["tasks"]) -> MetaOapg.properties.tasks: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["events"]) -> MetaOapg.properties.events: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["remote_data"]) -> MetaOapg.properties.remote_data: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["title", "description", "id", "owner", "first_name", "last_name", "name", "department", "addresses", "phone_numbers", "emails", "account", "last_activity_at", "created_at", "updated_at", "is_deleted", "tasks", "events", "remote_data", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["title"]) -> MetaOapg.properties.title: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["description"]) -> MetaOapg.properties.description: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["id"]) -> MetaOapg.properties.id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["owner"]) -> 'PartialOwnerNullable': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["first_name"]) -> MetaOapg.properties.first_name: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["last_name"]) -> MetaOapg.properties.last_name: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["name"]) -> MetaOapg.properties.name: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["department"]) -> MetaOapg.properties.department: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["addresses"]) -> MetaOapg.properties.addresses: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["phone_numbers"]) -> MetaOapg.properties.phone_numbers: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["emails"]) -> MetaOapg.properties.emails: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["account"]) -> 'PartialAccountNullable': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["last_activity_at"]) -> MetaOapg.properties.last_activity_at: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["created_at"]) -> MetaOapg.properties.created_at: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["updated_at"]) -> MetaOapg.properties.updated_at: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["is_deleted"]) -> MetaOapg.properties.is_deleted: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["tasks"]) -> MetaOapg.properties.tasks: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["events"]) -> MetaOapg.properties.events: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["remote_data"]) -> MetaOapg.properties.remote_data: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["title", "description", "id", "owner", "first_name", "last_name", "name", "department", "addresses", "phone_numbers", "emails", "account", "last_activity_at", "created_at", "updated_at", "is_deleted", "tasks", "events", "remote_data", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        owner: 'PartialOwnerNullable',
        addresses: typing.Union[MetaOapg.properties.addresses, list, tuple, ],
        created_at: typing.Union[MetaOapg.properties.created_at, str, ],
        description: typing.Union[MetaOapg.properties.description, None, str, ],
        last_name: typing.Union[MetaOapg.properties.last_name, None, str, ],
        title: typing.Union[MetaOapg.properties.title, None, str, ],
        remote_data: typing.Union[MetaOapg.properties.remote_data, dict, frozendict.frozendict, None, ],
        emails: typing.Union[MetaOapg.properties.emails, list, tuple, ],
        phone_numbers: typing.Union[MetaOapg.properties.phone_numbers, list, tuple, ],
        is_deleted: typing.Union[MetaOapg.properties.is_deleted, bool, ],
        updated_at: typing.Union[MetaOapg.properties.updated_at, str, ],
        last_activity_at: typing.Union[MetaOapg.properties.last_activity_at, None, str, ],
        name: typing.Union[MetaOapg.properties.name, None, str, ],
        id: typing.Union[MetaOapg.properties.id, str, ],
        department: typing.Union[MetaOapg.properties.department, None, str, ],
        first_name: typing.Union[MetaOapg.properties.first_name, None, str, ],
        account: 'PartialAccountNullable',
        events: typing.Union[MetaOapg.properties.events, list, tuple, None, ],
        tasks: typing.Union[MetaOapg.properties.tasks, list, tuple, None, ],
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'Contact':
        return super().__new__(
            cls,
            *args,
            owner=owner,
            addresses=addresses,
            created_at=created_at,
            description=description,
            last_name=last_name,
            title=title,
            remote_data=remote_data,
            emails=emails,
            phone_numbers=phone_numbers,
            is_deleted=is_deleted,
            updated_at=updated_at,
            last_activity_at=last_activity_at,
            name=name,
            id=id,
            department=department,
            first_name=first_name,
            account=account,
            events=events,
            tasks=tasks,
            _configuration=_configuration,
            **kwargs,
        )

from carbon.model.address import Address
from carbon.model.email import Email
from carbon.model.event import Event
from carbon.model.partial_account_nullable import PartialAccountNullable
from carbon.model.partial_owner_nullable import PartialOwnerNullable
from carbon.model.phone_number import PhoneNumber
from carbon.model.task import Task
