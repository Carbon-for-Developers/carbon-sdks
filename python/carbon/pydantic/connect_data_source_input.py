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

from carbon.pydantic.confluence_authentication import ConfluenceAuthentication
from carbon.pydantic.freskdesk_authentication import FreskdeskAuthentication
from carbon.pydantic.gitbook_authetication import GitbookAuthetication
from carbon.pydantic.github_authentication import GithubAuthentication
from carbon.pydantic.notion_authentication import NotionAuthentication
from carbon.pydantic.o_auth_authentication import OAuthAuthentication
from carbon.pydantic.s3_authentication import S3Authentication
from carbon.pydantic.salesforce_authentication import SalesforceAuthentication
from carbon.pydantic.sharepoint_authentication import SharepointAuthentication
from carbon.pydantic.sync_options import SyncOptions
from carbon.pydantic.zendesk_authentication import ZendeskAuthentication
from carbon.pydantic.zotero_authentication import ZoteroAuthentication

class ConnectDataSourceInput(BaseModel):
    authentication: typing.Union[OAuthAuthentication, NotionAuthentication, SharepointAuthentication, ConfluenceAuthentication, ZendeskAuthentication, ZoteroAuthentication, GitbookAuthetication, SalesforceAuthentication, FreskdeskAuthentication, S3Authentication, GithubAuthentication] = Field(alias='authentication')

    sync_options: typing.Optional[SyncOptions] = Field(None, alias='sync_options')

    model_config = ConfigDict(
        protected_namespaces=(),
        arbitrary_types_allowed=True
    )
