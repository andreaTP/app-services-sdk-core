"""
    Service Registry API

    Service Registry Instance API  NOTE: This API cannot be called directly from the portal.  # noqa: E501

    The version of the OpenAPI document: 2.2.5.Final
    Contact: apicurio@lists.jboss.org
    Generated by: https://openapi-generator.tech
"""


import unittest

import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api.system_api import SystemApi  # noqa: E501


class TestSystemApi(unittest.TestCase):
    """SystemApi unit test stubs"""

    def setUp(self):
        self.api = SystemApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_get_resource_limits(self):
        """Test case for get_resource_limits

        Get resource limits information  # noqa: E501
        """
        pass

    def test_get_system_info(self):
        """Test case for get_system_info

        Get system information  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
