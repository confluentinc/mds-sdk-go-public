# \PrivateRBACUIManageRoleBindingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterAccessInfo**](PrivateRBACUIManageRoleBindingsApi.md#ClusterAccessInfo) | **Post** /security/1.0/lookup/managed/clusters/principal/{principal} | Identify the rolebinding abilities (view vs manage) the user has on the specified scope.
[**ManagedRoleBindings**](PrivateRBACUIManageRoleBindingsApi.md#ManagedRoleBindings) | **Post** /security/1.0/lookup/managed/rolebindings/principal/{principal} | Identify the rolebindings this user can see and manage.



## ClusterAccessInfo

> ClusterAccessInfo ClusterAccessInfo(ctx, principal, scope)

Identify the rolebinding abilities (view vs manage) the user has on the specified scope.

Used by the Confluent Control Center UI to control access to rolebinding add/remove buttons.  Callable by Admins+ResourceOwners. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
**scope** | [**Scope**](Scope.md)|  | 

### Return type

[**ClusterAccessInfo**](ClusterAccessInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedRoleBindings

> ManagedRoleBindings ManagedRoleBindings(ctx, principal, scope, optional)

Identify the rolebindings this user can see and manage.

Callable by Admins+ResourceOwners.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
**scope** | [**Scope**](Scope.md)|  | 
 **optional** | ***ManagedRoleBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ManagedRoleBindingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceType** | **optional.String**| Filter down by resource type. | 

### Return type

[**ManagedRoleBindings**](ManagedRoleBindings.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

