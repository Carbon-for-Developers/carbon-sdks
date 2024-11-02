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


class ListItemsFiltersNullable(
    schemas.DictBase,
    schemas.NoneBase,
    schemas.Schema,
    schemas.NoneFrozenDictMixin
):
    """
    This class is auto generated by Konfig (https://konfigthis.com)
    """


    class MetaOapg:
        
        class properties:
        
            @staticmethod
            def external_ids() -> typing.Type['ListItemsFiltersNullableExternalIds']:
                return ListItemsFiltersNullableExternalIds
        
            @staticmethod
            def ids() -> typing.Type['ListItemsFiltersNullableIds']:
                return ListItemsFiltersNullableIds
            
            
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
            
            
            class root_files_only(
                schemas.BoolBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneBoolMixin
            ):
            
            
                def __new__(
                    cls,
                    *args: typing.Union[None, bool, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'root_files_only':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
            
            
            class file_formats(
                schemas.ListBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneTupleMixin
            ):
            
            
                class MetaOapg:
                    
                    @staticmethod
                    def items() -> typing.Type['StorageFileFormats']:
                        return StorageFileFormats
            
            
                def __new__(
                    cls,
                    *args: typing.Union[list, tuple, None, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'file_formats':
                    return super().__new__(
                        cls,
                        *args,
                        _configuration=_configuration,
                    )
        
            @staticmethod
            def item_types() -> typing.Type['ListItemsFiltersNullableItemTypes']:
                return ListItemsFiltersNullableItemTypes
            __annotations__ = {
                "external_ids": external_ids,
                "ids": ids,
                "name": name,
                "root_files_only": root_files_only,
                "file_formats": file_formats,
                "item_types": item_types,
            }

    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["external_ids"]) -> 'ListItemsFiltersNullableExternalIds': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["ids"]) -> 'ListItemsFiltersNullableIds': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["name"]) -> MetaOapg.properties.name: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["root_files_only"]) -> MetaOapg.properties.root_files_only: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["file_formats"]) -> MetaOapg.properties.file_formats: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["item_types"]) -> 'ListItemsFiltersNullableItemTypes': ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["external_ids", "ids", "name", "root_files_only", "file_formats", "item_types", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["external_ids"]) -> typing.Union['ListItemsFiltersNullableExternalIds', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["ids"]) -> typing.Union['ListItemsFiltersNullableIds', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["name"]) -> typing.Union[MetaOapg.properties.name, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["root_files_only"]) -> typing.Union[MetaOapg.properties.root_files_only, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["file_formats"]) -> typing.Union[MetaOapg.properties.file_formats, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["item_types"]) -> typing.Union['ListItemsFiltersNullableItemTypes', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["external_ids", "ids", "name", "root_files_only", "file_formats", "item_types", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, None, ],
        external_ids: typing.Union['ListItemsFiltersNullableExternalIds', schemas.Unset] = schemas.unset,
        ids: typing.Union['ListItemsFiltersNullableIds', schemas.Unset] = schemas.unset,
        name: typing.Union[MetaOapg.properties.name, None, str, schemas.Unset] = schemas.unset,
        root_files_only: typing.Union[MetaOapg.properties.root_files_only, None, bool, schemas.Unset] = schemas.unset,
        file_formats: typing.Union[MetaOapg.properties.file_formats, list, tuple, None, schemas.Unset] = schemas.unset,
        item_types: typing.Union['ListItemsFiltersNullableItemTypes', schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'ListItemsFiltersNullable':
        return super().__new__(
            cls,
            *args,
            external_ids=external_ids,
            ids=ids,
            name=name,
            root_files_only=root_files_only,
            file_formats=file_formats,
            item_types=item_types,
            _configuration=_configuration,
            **kwargs,
        )

from carbon.model.list_items_filters_nullable_external_ids import ListItemsFiltersNullableExternalIds
from carbon.model.list_items_filters_nullable_ids import ListItemsFiltersNullableIds
from carbon.model.list_items_filters_nullable_item_types import ListItemsFiltersNullableItemTypes
from carbon.model.storage_file_formats import StorageFileFormats
