

# ConnectorType

Represents a connector type supported by the API

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **String** |  |  [optional]
**kind** | **String** |  |  [optional]
**href** | **String** |  |  [optional]
**name** | **String** | Name of the connector type. | 
**version** | **String** | Version of the connector type. | 
**channels** | **List&lt;Channel&gt;** | Channels of the connector type. |  [optional]
**description** | **String** | A description of the connector. |  [optional]
**iconHref** | **String** | URL to an icon of the connector. |  [optional]
**labels** | **List&lt;String&gt;** | Labels used to categorize the connector |  [optional]
**annotations** | **Map&lt;String, String&gt;** | Name-value string annotations for resource |  [optional]
**featuredRank** | **Integer** | Ranking for featured connectors |  [optional]
**capabilities** | **List&lt;String&gt;** | The capabilities supported by the connector |  [optional]
**schema** | **Object** | A json schema that can be used to validate a ConnectorRequest connector field. | 



