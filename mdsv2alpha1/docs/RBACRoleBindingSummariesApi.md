# \RBACRoleBindingSummariesApi

All URIs are relative to *http://localhost/security/v2alpha1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LookupPrincipalsWithRole**](RBACRoleBindingSummariesApi.md#LookupPrincipalsWithRole) | **Post** /lookup/role/{roleName} | Look up the KafkaPrincipals who have the given role for the given scope.  Callable by Admins.
[**LookupPrincipalsWithRoleOnResource**](RBACRoleBindingSummariesApi.md#LookupPrincipalsWithRoleOnResource) | **Post** /lookup/role/{roleName}/resource/{resourceType}/name/{resourceName} | Look up the KafkaPrincipals who have the given role on the specified resource for the given scope.
[**ManagedNonResourceRoleBindingsAtScope**](RBACRoleBindingSummariesApi.md#ManagedNonResourceRoleBindingsAtScope) | **Post** /lookup/managed/rolebindings | Returns all non-resource rolebindings in the given scope for all users (not just the calling user) that the calling user has permission to see. A user can see, but not alter rolebindings for scopes that they have Describe access on, and alter rolebindings for scopes that they have Alter access on. Callable by Admins+Users. 
[**MyAllowedResources**](RBACRoleBindingSummariesApi.md#MyAllowedResources) | **Post** /lookup/resources/{resourceType}/operation/{operation} | List all resource patterns of the specified resourceType that the caller is allowed to perform the specified operation on. If the caller is not allowed to perform the operation on any resources, the list will be empty. Overlapping resource patterns will be \&quot;squashed\&quot; to eliminate redundancy, for example if you have access on both the prefix \&quot;topic-*\&quot; and the literal \&quot;topic-1\&quot;, only \&quot;topic-*\&quot; will be returned.  Callable by Admins+Users. 
[**MyRoleBindings**](RBACRoleBindingSummariesApi.md#MyRoleBindings) | **Post** /lookup/rolebindings/principal/{principal} | List all rolebindings for the specifed principal in the scope and all contained scopes. Be aware that this simply looks at the rolebinding data, and does not mean that the scopes actually exist. Callable by Admins+Users. 



## LookupPrincipalsWithRole

> []string LookupPrincipalsWithRole(ctx, roleName, scope)

Look up the KafkaPrincipals who have the given role for the given scope.  Callable by Admins.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string**| Role name to look up. | 
**scope** | [**Scope**](Scope.md)|  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupPrincipalsWithRoleOnResource

> []string LookupPrincipalsWithRoleOnResource(ctx, roleName, resourceType, resourceName, scope)

Look up the KafkaPrincipals who have the given role on the specified resource for the given scope.

Callable by Admins.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string**| Role name to look up. | 
**resourceType** | **string**| Type of resource to look up. | 
**resourceName** | **string**| Name of resource to look up. | 
**scope** | [**Scope**](Scope.md)|  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedNonResourceRoleBindingsAtScope

> []ManagedRoleBinding ManagedNonResourceRoleBindingsAtScope(ctx, scope)

Returns all non-resource rolebindings in the given scope for all users (not just the calling user) that the calling user has permission to see. A user can see, but not alter rolebindings for scopes that they have Describe access on, and alter rolebindings for scopes that they have Alter access on. Callable by Admins+Users. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | [**Scope**](Scope.md)|  | 

### Return type

[**[]ManagedRoleBinding**](ManagedRoleBinding.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MyAllowedResources

> []ResourcePattern MyAllowedResources(ctx, resourceType, operation, scope)

List all resource patterns of the specified resourceType that the caller is allowed to perform the specified operation on. If the caller is not allowed to perform the operation on any resources, the list will be empty. Overlapping resource patterns will be \"squashed\" to eliminate redundancy, for example if you have access on both the prefix \"topic-*\" and the literal \"topic-1\", only \"topic-*\" will be returned.  Callable by Admins+Users. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of resource we want patterns for. | 
**operation** | **string**| The operation we want to check. | 
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


## MyRoleBindings

> []ScopeRoleBindingMapping MyRoleBindings(ctx, principal, scope)

List all rolebindings for the specifed principal in the scope and all contained scopes. Be aware that this simply looks at the rolebinding data, and does not mean that the scopes actually exist. Callable by Admins+Users. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user. | 
**scope** | [**Scope**](Scope.md)|  | 

### Return type

[**[]ScopeRoleBindingMapping**](ScopeRoleBindingMapping.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

