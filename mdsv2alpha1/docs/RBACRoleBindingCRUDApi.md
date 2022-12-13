# \RBACRoleBindingCRUDApi

All URIs are relative to *http://localhost/security/v2alpha1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoleForPrincipal**](RBACRoleBindingCRUDApi.md#AddRoleForPrincipal) | **Post** /principals/{principal}/roles/{roleName} | Binds the principal to a role for a specific cluster or in the given scope. Callable by Admins.
[**AddRoleResourcesForPrincipal**](RBACRoleBindingCRUDApi.md#AddRoleResourcesForPrincipal) | **Post** /principals/{principal}/roles/{roleName}/bindings | Incrementally grant the resources to the principal at the given scope/cluster using the given role.
[**DeleteAllRolesForPrincipal**](RBACRoleBindingCRUDApi.md#DeleteAllRolesForPrincipal) | **Delete** /principals/{principal}/roles | Remove all roles for the principal at the given scope and all contained scopes. Callable by Admins. 
[**DeleteRoleForPrincipal**](RBACRoleBindingCRUDApi.md#DeleteRoleForPrincipal) | **Delete** /principals/{principal}/roles/{roleName} | Remove the role from the principal at the given scope/cluster. No-op if the user doesn&#39;t have the role.  Callable by Admins. 
[**GetRoleResourcesForPrincipal**](RBACRoleBindingCRUDApi.md#GetRoleResourcesForPrincipal) | **Post** /principals/{principal}/roles/{roleName}/resources | Look up the rolebindings for the principal at the given scope/cluster using the given role.
[**RemoveRoleResourcesForPrincipal**](RBACRoleBindingCRUDApi.md#RemoveRoleResourcesForPrincipal) | **Delete** /principals/{principal}/roles/{roleName}/bindings | Incrementally remove the resources from the principal at the given scope/cluster using the given role.
[**SetRoleResourcesForPrincipal**](RBACRoleBindingCRUDApi.md#SetRoleResourcesForPrincipal) | **Put** /principals/{principal}/roles/{roleName}/bindings | Overwrite existing resource grants.



## AddRoleForPrincipal

> AddRoleForPrincipal(ctx, principal, roleName, scope)

Binds the principal to a role for a specific cluster or in the given scope. Callable by Admins.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user. | 
**roleName** | **string**| The name of the role to bind the user to. | 
**scope** | [**Scope**](Scope.md)|  | 

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
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user. | 
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


## DeleteAllRolesForPrincipal

> DeleteAllRolesForPrincipal(ctx, principal, scope)

Remove all roles for the principal at the given scope and all contained scopes. Callable by Admins. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user. | 
**scope** | [**Scope**](Scope.md)|  | 

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

> DeleteRoleForPrincipal(ctx, principal, roleName, scope)

Remove the role from the principal at the given scope/cluster. No-op if the user doesn't have the role.  Callable by Admins. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user. | 
**roleName** | **string**| The name of the role. | 
**scope** | [**Scope**](Scope.md)|  | 

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

> []ResourcePattern GetRoleResourcesForPrincipal(ctx, principal, roleName, scope)

Look up the rolebindings for the principal at the given scope/cluster using the given role.

Callable by Admins and ResourceOwners.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user. | 
**roleName** | **string**| The name of the role. | 
**scope** | [**Scope**](Scope.md)|  | 

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
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user. | 
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
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user. | 
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

