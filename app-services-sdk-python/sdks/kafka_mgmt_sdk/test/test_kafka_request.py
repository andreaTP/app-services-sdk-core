"""
    Kafka Management API

    Kafka Management API is a REST API to manage Kafka instances  # noqa: E501

    The version of the OpenAPI document: 1.16.0
    Contact: rhosak-support@redhat.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import rhoas_kafka_mgmt_sdk
from rhoas_kafka_mgmt_sdk.model.kafka_request_all_of import KafkaRequestAllOf
from rhoas_kafka_mgmt_sdk.model.object_reference import ObjectReference
from rhoas_kafka_mgmt_sdk.model.supported_kafka_size_bytes_value_item import SupportedKafkaSizeBytesValueItem
globals()['KafkaRequestAllOf'] = KafkaRequestAllOf
globals()['ObjectReference'] = ObjectReference
globals()['SupportedKafkaSizeBytesValueItem'] = SupportedKafkaSizeBytesValueItem
from rhoas_kafka_mgmt_sdk.model.kafka_request import KafkaRequest


class TestKafkaRequest(unittest.TestCase):
    """KafkaRequest unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testKafkaRequest(self):
        """Test KafkaRequest"""
        # FIXME: construct object with mandatory attributes with example values
        # model = KafkaRequest()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
