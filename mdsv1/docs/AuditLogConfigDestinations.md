# AuditLogConfigDestinations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapServers** | **[]string** | A list of Kafka broker endpoints, used when configuring the Kafka producer(s) that will emit audit log events as they occur.  See Kafka Producer Configuration &#x60;&#x60;bootstrap.servers&#x60;&#x60;: https://kafka.apache.org/documentation/#producerconfigs  | [optional] 
**Topics** | [**map[string]AuditLogConfigDestinationConfig**](AuditLogConfigDestinationConfig.md) | The destination Kafka topics that can receive audit log events. The keys are Kafka topic names used in routes within the audit log configuration specification. The values are the configuration details for each destination topic. Note that topic names must match the pattern &#x60;&#x60;^confluent-audit-log-events[-_a-zA-Z0-9]*$&#x60;&#x60; and be 249 characters or less.  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


