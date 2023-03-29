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
from rhoas_kafka_mgmt_sdk.model.supported_kafka_size import SupportedKafkaSize
from rhoas_kafka_mgmt_sdk.model.supported_kafka_size_bytes_value_item import SupportedKafkaSizeBytesValueItem
globals()['SupportedKafkaSize'] = SupportedKafkaSize
globals()['SupportedKafkaSizeBytesValueItem'] = SupportedKafkaSizeBytesValueItem
from rhoas_kafka_mgmt_sdk.model.supported_kafka_instance_type_sizes_inner import SupportedKafkaInstanceTypeSizesInner


class TestSupportedKafkaInstanceTypeSizesInner(unittest.TestCase):
    """SupportedKafkaInstanceTypeSizesInner unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testSupportedKafkaInstanceTypeSizesInner(self):
        """Test SupportedKafkaInstanceTypeSizesInner"""
        # FIXME: construct object with mandatory attributes with example values
        # model = SupportedKafkaInstanceTypeSizesInner()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
