# \PrivateRBACUIMyRoleBindingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MyRoleBindingsAll**](PrivateRBACUIMyRoleBindingsApi.md#MyRoleBindingsAll) | **Get** /security/1.0/lookup/rolebindings/principal/{principal} | List all rolebindings for the specifed principal for all scopes and clusters that have any rolebindings.
[**MyRoleBindingsSingleScope**](PrivateRBACUIMyRoleBindingsApi.md#MyRoleBindingsSingleScope) | **Post** /security/1.0/lookup/rolebindings/principal/{principal} | List all rolebindings for the specifed principal and scope.



## MyRoleBindingsAll

> []ScopeRoleBindingMapping MyRoleBindingsAll(ctx, principal, optional)

List all rolebindings for the specifed principal for all scopes and clusters that have any rolebindings.

Be aware that this simply looks at the rolebinding data, and does not mean that the clusters actually exist.  Callable by Admins+User. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
 **optional** | ***MyRoleBindingsAllOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MyRoleBindingsAllOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterType** | **optional.String**| Filter down by a cluster type. | 

### Return type

[**[]ScopeRoleBindingMapping**](ScopeRoleBindingMapping.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MyRoleBindingsSingleScope

> ScopeRoleBindingMapping MyRoleBindingsSingleScope(ctx, principal, scope)

List all rolebindings for the specifed principal and scope.

Callable by Admins+User.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
**scope** | [**Scope**](Scope.md)|  | 

### Return type

[**ScopeRoleBindingMapping**](ScopeRoleBindingMapping.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

