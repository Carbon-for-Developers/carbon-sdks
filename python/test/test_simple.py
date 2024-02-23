# coding: utf-8

"""
    Carbon

    Connect external data to LLMs, no matter the source.

    The version of the OpenAPI document: 1.0.0
    Generated by: https://konfigthis.com
"""

import unittest

import os
from pprint import pprint
from carbon import Carbon

class TestSimple(unittest.TestCase):
    def setUp(self):
        pass

    def test_auth_schemes(self):
        carbon = Carbon(
            api_key = 'YOUR_API_KEY',
            customer_id = 'YOUR_API_KEY',
            host = "http://127.0.0.1:4010"
        )
        token = carbon.auth.get_access_token()
        self.assertIsNotNone(token)
        pprint(token)
        carbon = Carbon(
            access_token = token.access_token,
            host = "http://127.0.0.1:4010"
        )
        self.assertIsNotNone(carbon)
        white_labeling = carbon.auth.get_white_labeling()
        self.assertIsNotNone(white_labeling)
        pprint(white_labeling)

    def test_various_endpoints(self):
        carbon = Carbon(
            api_key = 'YOUR_API_KEY',
            customer_id = 'YOUR_API_KEY',
            host = "http://127.0.0.1:4010"
        )
        data_sources = carbon.data_sources.query_user_data_sources(
            pagination={
                "limit": 10,
                "offset": 0,
            },
            order_by="created_at",
            order_dir="desc",
            filters={
                "source": "GOOGLE_DRIVE",
            },
        )
        pprint(data_sources)
        self.assertIsNotNone(data_sources)

        files = carbon.integrations.sync_s3_files(ids=[{"id": "ID", "bucket": "BUCKET"}])
        pprint(files)
        self.assertIsNotNone(files)

        status = carbon.health.check()
        pprint(status)
        self.assertIsNotNone(status)

        file = carbon.files.upload(file=open("README.md", "rb"))
        pprint(file)
        self.assertIsNotNone(file)

    def tearDown(self):
        pass


if __name__ == "__main__":
    unittest.main()
