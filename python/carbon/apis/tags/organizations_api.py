# coding: utf-8

"""
    Carbon

    Connect external data to LLMs, no matter the source.

    The version of the OpenAPI document: 1.0.0
    Generated by: https://konfigthis.com
"""

from carbon.paths.organization.get import Get
from carbon.apis.tags.organizations_api_raw import OrganizationsApiRaw


class OrganizationsApi(
    Get,
):
    """NOTE:
    This class is auto generated by Konfig (https://konfigthis.com)
    """
    raw: OrganizationsApiRaw

    def __init__(self, api_client=None):
        super().__init__(api_client)
        self.raw = OrganizationsApiRaw(api_client)
