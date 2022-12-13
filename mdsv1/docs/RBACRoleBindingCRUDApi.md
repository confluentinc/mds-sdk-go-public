# \RBACRoleBindingCRUDApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoleForPrincipal**](RBACRoleBindingCRUDApi.md#AddRoleForPrincipal) | **Post** /security/1.0/principals/{principal}/roles/{roleName} | Binds the principal to a cluster-scoped role for a specific cluster or in the given scope.
[**AddRoleResourcesForPrincipal**](RBACRoleBindingCRUDApi.md#AddRoleResourcesForPrincipal) | **Post** /security/1.0/principals/{principal}/roles/{roleName}/bindings | Incrementally grant the resources to the principal at the given scope/cluster using the given role.
[**DeleteRoleForPrincipal**](RBACRoleBindingCRUDApi.md#DeleteRoleForPrincipal) | **Delete** /security/1.0/principals/{principal}/roles/{roleName} | Remove the role (cluster or resource scoped) from the principal at the given scope/cluster.
[**GetRoleResourcesForPrincipal**](RBACRoleBindingCRUDApi.md#GetRoleResourcesForPrincipal) | **Post** /security/1.0/principals/{principal}/roles/{roleName}/resources | Look up the rolebindings for the principal at the given scope/cluster using the given role.
[**RemoveRoleResourcesForPrincipal**](RBACRoleBindingCRUDApi.md#RemoveRoleResourcesForPrincipal) | **Delete** /security/1.0/principals/{principal}/roles/{roleName}/bindings | Incrementally remove the resources from the principal at the given scope/cluster using the given role.
[**SetRoleResourcesForPrincipal**](RBACRoleBindingCRUDApi.md#SetRoleResourcesForPrincipal) | **Put** /security/1.0/principals/{principal}/roles/{roleName}/bindings | Overwrite existing resource grants.



## AddRoleForPrincipal

> AddRoleForPrincipal(ctx, principal, roleName, mdsScope)

Binds the principal to a cluster-scoped role for a specific cluster or in the given scope.

Callable by Admins.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
**roleName** | **string**| The name of the cluster-scoped role to bind the user to. | 
**mdsScope** | [**MdsScope**](MdsScope.md)|  | 

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


## AddRoleResourcesForPrincipal

> AddRoleResourcesForPrincipal(ctx, principal, roleName, resourcesRequest)

Incrementally grant the resources to the principal at the given scope/cluster using the given role.

Callable by Admins+ResourceOwners.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
**roleName** | **string**| The name of the role. | 
**resourcesRequest** | [**ResourcesRequest**](ResourcesRequest.md)|  | 

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


## DeleteRoleForPrincipal

> DeleteRoleForPrincipal(ctx, principal, roleName, mdsScope)

Remove the role (cluster or resource scoped) from the principal at the given scope/cluster.

No-op if the user doesn't have the role.  Callable by Admins. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
**roleName** | **string**| The name of the role. | 
**mdsScope** | [**MdsScope**](MdsScope.md)|  | 

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


## GetRoleResourcesForPrincipal

> []ResourcePattern GetRoleResourcesForPrincipal(ctx, principal, roleName, mdsScope)

Look up the rolebindings for the principal at the given scope/cluster using the given role.

Callable by Admins.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
**roleName** | **string**| The name of the role. | 
**mdsScope** | [**MdsScope**](MdsScope.md)|  | 

### Return type

[**[]ResourcePattern**](ResourcePattern.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRoleResourcesForPrincipal

> RemoveRoleResourcesForPrincipal(ctx, principal, roleName, resourcesRequest)

Incrementally remove the resources from the principal at the given scope/cluster using the given role.

Callable by Admins+ResourceOwners.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
**roleName** | **string**| The name of the role. | 
**resourcesRequest** | [**ResourcesRequest**](ResourcesRequest.md)|  | 

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


## SetRoleResourcesForPrincipal

> SetRoleResourcesForPrincipal(ctx, principal, roleName, resourcesRequest)

Overwrite existing resource grants.

Callable by Admins+ResourceOwners.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
**roleName** | **string**| The name of the role. | 
**resourcesRequest** | [**ResourcesRequest**](ResourcesRequest.md)|  | 

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

