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

from carbon.type.embedding_generators import EmbeddingGenerators
from carbon.type.webscrape_request_css_classes_to_skip import WebscrapeRequestCssClassesToSkip
from carbon.type.webscrape_request_css_selectors_to_skip import WebscrapeRequestCssSelectorsToSkip
from carbon.type.webscrape_request_html_tags_to_skip import WebscrapeRequestHtmlTagsToSkip
from carbon.type.webscrape_request_tags import WebscrapeRequestTags
from carbon.type.webscrape_request_url_paths_to_include import WebscrapeRequestUrlPathsToInclude

class RequiredWebscrapeRequest(TypedDict):
    url: str


class OptionalWebscrapeRequest(TypedDict, total=False):
    tags: typing.Optional[WebscrapeRequestTags]

    recursion_depth: typing.Optional[int]

    max_pages_to_scrape: typing.Optional[int]

    chunk_size: typing.Optional[int]

    chunk_overlap: typing.Optional[int]

    skip_embedding_generation: typing.Optional[bool]

    enable_auto_sync: typing.Optional[bool]

    generate_sparse_vectors: typing.Optional[bool]

    prepend_filename_to_chunks: typing.Optional[bool]

    html_tags_to_skip: typing.Optional[WebscrapeRequestHtmlTagsToSkip]

    css_classes_to_skip: typing.Optional[WebscrapeRequestCssClassesToSkip]

    css_selectors_to_skip: typing.Optional[WebscrapeRequestCssSelectorsToSkip]

    embedding_model: EmbeddingGenerators

    url_paths_to_include: typing.Optional[WebscrapeRequestUrlPathsToInclude]

    # Whether the scraper should download css and media from the page (images, fonts, etc). Scrapes          might take longer to finish with this flag enabled, but the success rate is improved.
    download_css_and_media: typing.Optional[bool]

    # If this flag is enabled, the file will be chunked and stored with Carbon,           but no embeddings will be generated. This overrides the skip_embedding_generation flag.
    generate_chunks_only: bool

class WebscrapeRequest(RequiredWebscrapeRequest, OptionalWebscrapeRequest):
    pass
