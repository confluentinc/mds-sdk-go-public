# \KafkaACLManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAclBinding**](KafkaACLManagementApi.md#AddAclBinding) | **Post** /security/1.0/acls | Creates Kafka ACLs for given AclBinding.
[**RemoveAclBindings**](KafkaACLManagementApi.md#RemoveAclBindings) | **Delete** /security/1.0/acls | Deletes Kafka ACLs according to the supplied filter.
[**SearchAclBinding**](KafkaACLManagementApi.md#SearchAclBinding) | **Post** /security/1.0/acls:search | Lists Kafka ACLs according to the supplied filter.



## AddAclBinding

> AddAclBinding(ctx, createAclRequest)

Creates Kafka ACLs for given AclBinding.

Callable by Admins+AclUsers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createAclRequest** | [**CreateAclRequest**](CreateAclRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAclBindings

> []AclBinding RemoveAclBindings(ctx, aclFilterRequest)

Deletes Kafka ACLs according to the supplied filter.

Callable by Admins+AclUsers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aclFilterRequest** | [**AclFilterRequest**](AclFilterRequest.md)|  | 

### Return type

[**[]AclBinding**](AclBinding.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAclBinding

> []AclBinding SearchAclBinding(ctx, aclFilterRequest)

Lists Kafka ACLs according to the supplied filter.

Callable by Admins+AclUsers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aclFilterRequest** | [**AclFilterRequest**](AclFilterRequest.md)|  | 

### Return type

[**[]AclBinding**](AclBinding.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

