# coding: utf-8

"""
    Carbon

    Connect external data to LLMs, no matter the source.

    The version of the OpenAPI document: 1.0.0
    Generated by: https://konfigthis.com
"""

import unittest
from unittest.mock import patch

import urllib3

import carbon
from carbon.paths.auth_v1_access_token import get
from carbon import configuration, schemas, api_client

from .. import ApiTestMixin


class TestAuthV1AccessToken(ApiTestMixin, unittest.TestCase):
    """
    AuthV1AccessToken unit test stubs
        Get Access Token
    """

    def setUp(self):
        pass

    def tearDown(self):
        pass

    response_status = 200




if __name__ == '__main__':
    unittest.main()
