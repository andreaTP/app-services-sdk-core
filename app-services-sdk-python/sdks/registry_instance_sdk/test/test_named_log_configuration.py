"""
    Service Registry API

    Service Registry Instance API  NOTE: This API cannot be called directly from the portal.  # noqa: E501

    The version of the OpenAPI document: 2.2.5.Final
    Contact: apicurio@lists.jboss.org
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.model.log_configuration import LogConfiguration
from rhoas_registry_instance_sdk.model.log_level import LogLevel
from rhoas_registry_instance_sdk.model.named_log_configuration_all_of import NamedLogConfigurationAllOf
globals()['LogConfiguration'] = LogConfiguration
globals()['LogLevel'] = LogLevel
globals()['NamedLogConfigurationAllOf'] = NamedLogConfigurationAllOf
from rhoas_registry_instance_sdk.model.named_log_configuration import NamedLogConfiguration


class TestNamedLogConfiguration(unittest.TestCase):
    """NamedLogConfiguration unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testNamedLogConfiguration(self):
        """Test NamedLogConfiguration"""
        # FIXME: construct object with mandatory attributes with example values
        # model = NamedLogConfiguration()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
