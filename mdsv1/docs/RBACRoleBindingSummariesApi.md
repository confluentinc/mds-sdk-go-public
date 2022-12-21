# \RBACRoleBindingSummariesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LookupPrincipalsWithRole**](RBACRoleBindingSummariesApi.md#LookupPrincipalsWithRole) | **Post** /security/1.0/lookup/role/{roleName} | Look up the KafkaPrincipals who have the given role for the given scope.
[**LookupPrincipalsWithRoleOnResource**](RBACRoleBindingSummariesApi.md#LookupPrincipalsWithRoleOnResource) | **Post** /security/1.0/lookup/role/{roleName}/resource/{resourceType}/name/{resourceName} | Look up the KafkaPrincipals who have the given role on the specified resource for the given scope.
[**LookupResourcesForPrincipal**](RBACRoleBindingSummariesApi.md#LookupResourcesForPrincipal) | **Post** /security/1.0/lookup/principal/{principal}/resources | Look up the resource bindings for the principal at the given scope/cluster.
[**ScopedPrincipalRolenames**](RBACRoleBindingSummariesApi.md#ScopedPrincipalRolenames) | **Post** /security/1.0/lookup/principals/{principal}/roleNames | Returns the effective list of role names for a principal.



## LookupPrincipalsWithRole

> []string LookupPrincipalsWithRole(ctx, roleName, mdsScope)

Look up the KafkaPrincipals who have the given role for the given scope.

Callable by Admins.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string**| Role name to look up. | 
**mdsScope** | [**MdsScope**](MdsScope.md)|  | 

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

> []string LookupPrincipalsWithRoleOnResource(ctx, roleName, resourceType, resourceName, mdsScope)

Look up the KafkaPrincipals who have the given role on the specified resource for the given scope.

Callable by Admins.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string**| Role name to look up. | 
**resourceType** | **string**| Type of resource to look up. | 
**resourceName** | **string**| Name of resource to look up. | 
**mdsScope** | [**MdsScope**](MdsScope.md)|  | 

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


## LookupResourcesForPrincipal

> map[string]map[string][]ResourcePattern LookupResourcesForPrincipal(ctx, principal, mdsScope)

Look up the resource bindings for the principal at the given scope/cluster.

Includes bindings from groups that the user belongs to.  Callable by Admins+User. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
**mdsScope** | [**MdsScope**](MdsScope.md)|  | 

### Return type

[**map[string]map[string][]ResourcePattern**](map.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScopedPrincipalRolenames

> []string ScopedPrincipalRolenames(ctx, principal, mdsScope)

Returns the effective list of role names for a principal.

For groups, these are the roles that are bound.  For users, this is the combination of roles granted to the specific user and roles granted to the user's groups.  Callable by Admins+User. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
**mdsScope** | [**MdsScope**](MdsScope.md)|  | 

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

