# \PrivateRBACUIClusterVisibilityApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListManagedClusters**](PrivateRBACUIClusterVisibilityApi.md#ListManagedClusters) | **Get** /security/1.0/lookup/managed/clusters/principal/{principal} | Identifies the scopes for the rolebindings that a user can see.
[**Visibility**](PrivateRBACUIClusterVisibilityApi.md#Visibility) | **Post** /security/1.0/lookup/principals/{principal}/visibility | Endpoint for Confluent Control Center to determine visibilty of Kafka and its sub-clusters for the specifed principal.



## ListManagedClusters

> []Scope ListManagedClusters(ctx, principal, optional)

Identifies the scopes for the rolebindings that a user can see.

May include rolebindings from scopes and clusters that never existed or previously existed (in other words, rolebindings that have been decommissioned, but are still defined in the system).  Callable by Admins+User. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
 **optional** | ***ListManagedClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListManagedClustersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterType** | **optional.String**| Filter down by cluster type. | 

### Return type

[**[]Scope**](Scope.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Visibility

> VisibilityResponse Visibility(ctx, principal, visibilityRequest)

Endpoint for Confluent Control Center to determine visibilty of Kafka and its sub-clusters for the specifed principal.

The intent is that this endpoint is called with cluster IDs that actually exist.  Callable by Admins+User. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
**visibilityRequest** | [**[]VisibilityRequest**](VisibilityRequest.md)|  | 

### Return type

[**VisibilityResponse**](VisibilityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

