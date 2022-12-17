# \CloudLifecycleApi

All URIs are relative to *http://localhost/security/v2alpha1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DuplicateRolesForOrg**](CloudLifecycleApi.md#DuplicateRolesForOrg) | **Post** /cloudlifecycle/rolebindings/{sourceOrgId}/copy | Duplicate all role bindings from the source organization to the destination organization
[**RemoveAllRoleBindingsForScope**](CloudLifecycleApi.md#RemoveAllRoleBindingsForScope) | **Delete** /cloudlifecycle/rolebindings | Delete all role bindings at the given scope
[**ScopeUndelete**](CloudLifecycleApi.md#ScopeUndelete) | **Post** /cloudlifecycle/scope/undelete | Undelete all role bindings for a given scope and reason.
[**UserUndelete**](CloudLifecycleApi.md#UserUndelete) | **Post** /cloudlifecycle/user/undelete | Undelete all role bindings for a given user, org, and reason.



## DuplicateRolesForOrg

> DuplicateRolesForOrg(ctx, sourceOrgId, duplicateRequest)

Duplicate all role bindings from the source organization to the destination organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceOrgId** | **string**| The UUID resource Identifier for the source organization | 
**duplicateRequest** | [**DuplicateRequest**](DuplicateRequest.md)|  | 

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


## RemoveAllRoleBindingsForScope

> RemoveAllRoleBindingsForScope(ctx, transactionId, scope)

Delete all role bindings at the given scope

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string**| Transaction Identifier supplied by calling service, stored in reason column | 
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


## ScopeUndelete

> []string ScopeUndelete(ctx, scopeUndeleteRequest)

Undelete all role bindings for a given scope and reason.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopeUndeleteRequest** | [**ScopeUndeleteRequest**](ScopeUndeleteRequest.md)|  | 

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


## UserUndelete

> []string UserUndelete(ctx, userUndeleteRequest)

Undelete all role bindings for a given user, org, and reason.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUndeleteRequest** | [**UserUndeleteRequest**](UserUndeleteRequest.md)|  | 

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

