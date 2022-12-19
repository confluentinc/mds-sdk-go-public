# \AuditLogConfigurationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfig**](AuditLogConfigurationApi.md#GetConfig) | **Get** /security/1.0/audit/config | Get the entire audit log configuration, including the live retention time policy values (&#x60;&#x60;retention_ms&#x60;&#x60;) for the destination topics.
[**ListRoutes**](AuditLogConfigurationApi.md#ListRoutes) | **Get** /security/1.0/audit/routes | Lists all currently defined routes that match the queried resource or its sub-resources.
[**PutConfig**](AuditLogConfigurationApi.md#PutConfig) | **Put** /security/1.0/audit/config | Update the entire audit log configuration on the MDS cluster and all Kafka clusters known to the cluster registry.
[**ResolveResourceRoute**](AuditLogConfigurationApi.md#ResolveResourceRoute) | **Get** /security/1.0/audit/lookup | Returns the route describing how messages regarding this CRN would be routed.



## GetConfig

> AuditLogConfigSpec GetConfig(ctx, )

Get the entire audit log configuration, including the live retention time policy values (``retention_ms``) for the destination topics.

Requires the \"AuditAdmin\" role on the metadata service (MDS) cluster and every Kafka cluster in the cluster registry.  Callable by Admins. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**AuditLogConfigSpec**](AuditLogConfigSpec.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoutes

> AuditLogConfigListRoutesResponse ListRoutes(ctx, optional)

Lists all currently defined routes that match the queried resource or its sub-resources.

Multiple routes may match a resource, but only the most specific route will be selected for events related to the resource. This endpoint returns all matching routes regardless of whether or not they would actually be selected, or ignored in favor of a more specific route.  Requires the \"AuditAdmin\" role on the metadata service (MDS) cluster and every Kafka cluster in the cluster registry.  Callable by Admins.  The CRN patterns in the audit log config routes can contain wildcards. So a route with a CRN pattern like ``crn://mds.example.com/kafka=*_/topic=finance-*`` would match events associated with the topic at address ``crn://mds.example.com/kafka=abc123/topic=finance-deposits``, or events associated with the topic at ``crn://mds.example.com/kafka=xyz789/topic=finance-chargebacks``, but would not match events associated with the topic ``crn://mds.example.com/kafka=abc123/topic=server-deployments``. So a route's CRN pattern can match events from more than one resource, based on where the pattern's wildcards are.  It is possible to write multiple routes with different CRN patterns that match a given resource's CRN. For example: the resource at ``crn://mds.example.com/kafka=abc123/topic=finance-chargebacks`` is matched by any of the following route CRN patterns:  * ``crn://mds.example.com/kafka=*_/topic=*`` * ``crn://mds.example.com/kafka=abc123/topic=*`` * ``crn://mds.example.com/kafka=*_/topic=finance-*``  When there are multiple matching routes for an event, we select the matching route with the most specific CRN pattern. The most specific CRN pattern is the one with the greatest length before its first wildcard. So in the above example, ``crn://mds.example.com/kafka=abc123/topic=*`` wins.  To break a tie, ignore the prefix that the patterns have in common. So, for example ``crn://mds.example.com/kafka=*_/topic=finance-*`` is more specific than ``crn://mds.example.com/kafka=*_/topic=*``.  This endpoint lists all currently defined routes that match the queried resource or its sub-resources, regardless of whether or not they would actually be selected, or ignored in favor of a more specific route.  A query pattern like ...  ``crn://mds1.example.com/kafka=abcde_FGHIJKL-01234567/connect=qa-test``  ... would match all of the following routes ...  ``crn://mds1.example.com/kafka=abcde_FGHIJKL-01234567/connect=qa-test/connector=from-db4``  ``crn://mds1.example.com/kafka=abcde_FGHIJKL-01234567/connect=qa-test/connector=*``  ``crn://mds1.example.com/kafka=abcde_FGHIJKL-01234567/connect=*_/connector=*``  ``crn://mds1.example.com/kafka=abcde_FGHIJKL-01234567/connect=qa-*``  ``crn://mds1.example.com/kafka=abcde_FGHIJKL-01234567/connect=*``  ``crn://mds1.example.com/kafka=*_/connect=qa-*``  ``crn://mds1.example.com/kafka=*_/connect=qa-*_/connector=*``  ... but would not match any of these routes ...  ``crn://mds1.example.com/kafka=*_/ksql=*``  ``crn://mds1.example.com/kafka=abcde_FGHIJKL-01234567``  ``crn://mds1.example.com/kafka=abcde_FGHIJKL-01234567/connect=stg-*``  ``crn://mds1.example.com/kafka=zyxwv-UTSRQPO_98765432/connect=qa-*``  ``crn://mds1.example.com/kafka=abcde_FGHIJKL-01234567/topic=qa-*`` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListRoutesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRoutesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| A Confluent resource name (CRN) .  | 

### Return type

[**AuditLogConfigListRoutesResponse**](AuditLogConfigListRoutesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConfig

> AuditLogConfigSpec PutConfig(ctx, auditLogConfigSpec)

Update the entire audit log configuration on the MDS cluster and all Kafka clusters known to the cluster registry.

Also creates missing destination topics on the destination cluster and updates the retention time policy of destination topics, if necessary.  Requires the \"AuditAdmin\" role on the MDS cluster and every Kafka cluster in the cluster registry.  May result in a 409 conflict error status if the ``resource_version`` in the JSON body of the request does not match the current version.  Callable by Admins. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditLogConfigSpec** | [**AuditLogConfigSpec**](AuditLogConfigSpec.md)|  | 

### Return type

[**AuditLogConfigSpec**](AuditLogConfigSpec.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveResourceRoute

> AuditLogConfigResolveResourceRouteResponse ResolveResourceRoute(ctx, optional)

Returns the route describing how messages regarding this CRN would be routed.

Requires the \"AuditAdmin\" role on the metadata service (MDS) cluster and every Kafka cluster in the cluster registry.  The CRN patterns in the audit log config routes can contain wildcards. So a route with a CRN pattern like ``crn://mds.example.com/kafka=*_/topic=finance-*`` would match events associated with the topic at address ``crn://mds.example.com/kafka=abc123/topic=finance-deposits``, or events associated with the topic at ``crn://mds.example.com/kafka=xyz789/topic=finance-chargebacks``, but would not match events associated with the topic ``crn://mds.example.com/kafka=abc123/topic=server-deployments``. So a route's CRN pattern can match events from more than one resource, based on where the pattern's wildcards are.  It is possible to write multiple routes with different CRN patterns that match a given resource's CRN. For example: the resource at ``crn://mds.example.com/kafka=abc123/topic=finance-chargebacks`` is matched by any of the following route CRN patterns:  * ``crn://mds.example.com/kafka=*_/topic=*`` * ``crn://mds.example.com/kafka=abc123/topic=*`` * ``crn://mds.example.com/kafka=*_/topic=finance-*``  When there are multiple matching routes for an event, we select the matching route with the most specific CRN pattern. The most specific CRN pattern is the one with the greatest length before its first wildcard. So in the above example, ``crn://mds.example.com/kafka=abc123/topic=*`` wins.  To break a tie, ignore the prefix that the patterns have in common. So, for example ``crn://mds.example.com/kafka=*_/topic=finance-*`` is more specific than ``crn://mds.example.com/kafka=*_/topic=*``.  This endpoint resolves the matching and precedence rules of all configured audit log routes for you, and returns the one selected (most specific) matching route describing how messages regarding the given CRN would be routed. Callable by Admins. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ResolveResourceRouteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ResolveResourceRouteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crn** | **optional.String**| A Confluent resource name (CRN). | 

### Return type

[**AuditLogConfigResolveResourceRouteResponse**](AuditLogConfigResolveResourceRouteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

