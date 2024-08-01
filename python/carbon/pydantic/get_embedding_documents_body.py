# coding: utf-8

"""
    Carbon

    Connect external data to LLMs, no matter the source.

    The version of the OpenAPI document: 1.0.0
    Generated by: https://konfigthis.com
"""

from datetime import datetime, date
import typing
from enum import Enum
from typing_extensions import TypedDict, Literal, TYPE_CHECKING
from pydantic import BaseModel, Field, RootModel, ConfigDict

from carbon.pydantic.embedding_generators_nullable import EmbeddingGeneratorsNullable
from carbon.pydantic.file_content_types_nullable import FileContentTypesNullable
from carbon.pydantic.get_embedding_documents_body_file_ids import GetEmbeddingDocumentsBodyFileIds
from carbon.pydantic.get_embedding_documents_body_parent_file_ids import GetEmbeddingDocumentsBodyParentFileIds
from carbon.pydantic.get_embedding_documents_body_query_vector import GetEmbeddingDocumentsBodyQueryVector
from carbon.pydantic.get_embedding_documents_body_tags import GetEmbeddingDocumentsBodyTags
from carbon.pydantic.helpdesk_file_types import HelpdeskFileTypes
from carbon.pydantic.hybrid_search_tuning_params_nullable import HybridSearchTuningParamsNullable
from carbon.pydantic.rerank_params_nullable import RerankParamsNullable

class GetEmbeddingDocumentsBody(BaseModel):
    # Query for which to get related chunks and embeddings.
    query: str = Field(alias='query')

    # Number of related chunks to return.
    k: int = Field(alias='k')

    tags: typing.Optional[GetEmbeddingDocumentsBodyTags] = Field(None, alias='tags')

    query_vector: typing.Optional[GetEmbeddingDocumentsBodyQueryVector] = Field(None, alias='query_vector')

    file_ids: typing.Optional[GetEmbeddingDocumentsBodyFileIds] = Field(None, alias='file_ids')

    # WARNING: This property is deprecated
    parent_file_ids: typing.Optional[GetEmbeddingDocumentsBodyParentFileIds] = Field(None, alias='parent_file_ids')

    # Flag to control whether or not to include all children of filtered files in the embedding search.
    include_all_children: typing.Optional[bool] = Field(None, alias='include_all_children')

    # A set of tags to limit the search to. Use this instead of `tags`, which is deprecated.
    tags_v2: typing.Optional[typing.Optional[typing.Dict[str, typing.Union[bool, date, datetime, dict, float, int, list, str, None]]]] = Field(None, alias='tags_v2')

    # Flag to control whether or not to include tags for each chunk in the response.
    include_tags: typing.Optional[typing.Optional[bool]] = Field(None, alias='include_tags')

    # Flag to control whether or not to include embedding vectors in the response.
    include_vectors: typing.Optional[typing.Optional[bool]] = Field(None, alias='include_vectors')

    # Flag to control whether or not to include a signed URL to the raw file containing each chunk         in the response.
    include_raw_file: typing.Optional[typing.Optional[bool]] = Field(None, alias='include_raw_file')

    # Flag to control whether or not to perform hybrid search.
    hybrid_search: typing.Optional[typing.Optional[bool]] = Field(None, alias='hybrid_search')

    hybrid_search_tuning_parameters: typing.Optional[HybridSearchTuningParamsNullable] = Field(None, alias='hybrid_search_tuning_parameters')

    media_type: typing.Optional[FileContentTypesNullable] = Field(None, alias='media_type')

    embedding_model: typing.Optional[EmbeddingGeneratorsNullable] = Field(None, alias='embedding_model')

    # Flag to control whether or not to include file-level metadata in the response. This metadata         will be included in the `content_metadata` field of each document along with chunk/embedding level metadata.
    include_file_level_metadata: typing.Optional[typing.Optional[bool]] = Field(None, alias='include_file_level_metadata')

    # Flag to control whether or not to perform a high accuracy embedding search. By default, this is set to false.         If true, the search may return more accurate results, but may take longer to complete.
    high_accuracy: typing.Optional[typing.Optional[bool]] = Field(None, alias='high_accuracy')

    rerank: typing.Optional[RerankParamsNullable] = Field(None, alias='rerank')

    # Filter files based on their type at the source (for example help center tickets and articles)
    file_types_at_source: typing.Optional[typing.Optional[typing.List[HelpdeskFileTypes]]] = Field(None, alias='file_types_at_source')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
