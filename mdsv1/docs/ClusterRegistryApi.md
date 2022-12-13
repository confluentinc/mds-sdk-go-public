# \ClusterRegistryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterRegistryList**](ClusterRegistryApi.md#ClusterRegistryList) | **Get** /security/1.0/registry/clusters | Returns a list of all clusters in the registry, optionally filtered by cluster type.
[**DeleteNamedCluster**](ClusterRegistryApi.md#DeleteNamedCluster) | **Delete** /security/1.0/registry/clusters/{clusterName} | Remove a named cluster from the registry.
[**GetNamedCluster**](ClusterRegistryApi.md#GetNamedCluster) | **Get** /security/1.0/registry/clusters/{clusterName} | Returns the information for a single named cluster, assuming the cluster exists and is visible to the calling principal.
[**UpdateClusters**](ClusterRegistryApi.md#UpdateClusters) | **Post** /security/1.0/registry/clusters | Define/overwrite named clusters.



## ClusterRegistryList

> []ClusterInfo ClusterRegistryList(ctx, optional)

Returns a list of all clusters in the registry, optionally filtered by cluster type.

If the calling principal doesn't have permissions to see the full cluster info, some information (\"hosts\", \"protocol\", etc) is redacted.  Callable by Admins+User. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClusterRegistryListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClusterRegistryListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterType** | **optional.String**| Optionally filter down by cluster type. | 

### Return type

[**[]ClusterInfo**](ClusterInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamedCluster

> DeleteNamedCluster(ctx, clusterName)

Remove a named cluster from the registry.

Callable by Admins.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string**| The name of cluster (ASCII printable characters without spaces). | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamedCluster

> ClusterInfo GetNamedCluster(ctx, clusterName)

Returns the information for a single named cluster, assuming the cluster exists and is visible to the calling principal.

Callable by Admins+User.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string**| The name of cluster (ASCII printable characters without spaces). | 

### Return type

[**ClusterInfo**](ClusterInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClusters

> UpdateClusters(ctx, clusterInfo)

Define/overwrite named clusters.

May result in a 409 Conflict if the name and scope combination of any cluster conflicts with existing clusters in the registry.  Callable by Admins. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterInfo** | [**[]ClusterInfo**](ClusterInfo.md)|  | 

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

