# \MetadataServiceOperationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Activenodes**](MetadataServiceOperationsApi.md#Activenodes) | **Get** /security/1.0/activenodes/{protocol} | Returns all the nodes running the Metadata Service REST API.
[**Features**](MetadataServiceOperationsApi.md#Features) | **Get** /security/1.0/features | What features the running version of MDS supports.
[**MetadataClusterId**](MetadataServiceOperationsApi.md#MetadataClusterId) | **Get** /security/1.0/metadataClusterId | Returns the ID of the Kafka cluster that MDS is running on.



## Activenodes

> []string Activenodes(ctx, protocol)

Returns all the nodes running the Metadata Service REST API.

Clients are expected to round robin call to these endpoints if they don't set up a load balancer in front of the Metadata Service nodes.  Callable by LDAP users. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**protocol** | **string**| Should be \&quot;http\&quot; or \&quot;https\&quot;. | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Features

> FeaturesInfo Features(ctx, )

What features the running version of MDS supports.

This endpoint does not require authentication.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**FeaturesInfo**](FeaturesInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataClusterId

> string MetadataClusterId(ctx, )

Returns the ID of the Kafka cluster that MDS is running on.

Callable by LDAP users.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

