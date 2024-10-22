# coding: utf-8

"""
    Carbon

    Connect external data to LLMs, no matter the source.

    The version of the OpenAPI document: 1.0.0
    Generated by: https://konfigthis.com
"""

from carbon.paths.data_sources_tags_add.post import AddTagsRaw
from carbon.paths.user_data_sources.post import QueryUserDataSourcesRaw
from carbon.paths.data_sources_tags_remove.post import RemoveTagsRaw
from carbon.paths.revoke_access_token.post import RevokeAccessTokenRaw


class DataSourcesApiRaw(
    AddTagsRaw,
    QueryUserDataSourcesRaw,
    RemoveTagsRaw,
    RevokeAccessTokenRaw,
):
    """NOTE:
    This class is auto generated by Konfig (https://konfigthis.com)
    """
    pass
