/*
Connector Management API

Testing ConnectorTypesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package connectormgmtclient

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_connectormgmtclient_ConnectorTypesApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConnectorTypesApiService GetConnectorTypeByID", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var connectorTypeId string

        resp, httpRes, err := apiClient.ConnectorTypesApi.GetConnectorTypeByID(context.Background(), connectorTypeId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConnectorTypesApiService GetConnectorTypeLabels", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConnectorTypesApi.GetConnectorTypeLabels(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConnectorTypesApiService GetConnectorTypes", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConnectorTypesApi.GetConnectorTypes(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
