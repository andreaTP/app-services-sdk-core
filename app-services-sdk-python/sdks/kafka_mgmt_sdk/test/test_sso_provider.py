"""
    Kafka Management API

    Kafka Management API is a REST API to manage Kafka instances  # noqa: E501

    The version of the OpenAPI document: 1.14.0
    Contact: rhosak-support@redhat.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import rhoas_kafka_mgmt_sdk
from rhoas_kafka_mgmt_sdk.model.object_reference import ObjectReference
from rhoas_kafka_mgmt_sdk.model.sso_provider_all_of import SsoProviderAllOf
globals()['ObjectReference'] = ObjectReference
globals()['SsoProviderAllOf'] = SsoProviderAllOf
from rhoas_kafka_mgmt_sdk.model.sso_provider import SsoProvider


class TestSsoProvider(unittest.TestCase):
    """SsoProvider unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testSsoProvider(self):
        """Test SsoProvider"""
        # FIXME: construct object with mandatory attributes with example values
        # model = SsoProvider()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
