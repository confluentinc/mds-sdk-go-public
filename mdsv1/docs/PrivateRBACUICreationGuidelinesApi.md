# \PrivateRBACUICreationGuidelinesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LookupCreateGuidelines**](PrivateRBACUICreationGuidelinesApi.md#LookupCreateGuidelines) | **Post** /security/1.0/lookup/principal/{principal}/resource/{resourceType}/operation/{operation} | Summarizes what resources and rolebindings this principal is allowed to create.



## LookupCreateGuidelines

> map[string]interface{} LookupCreateGuidelines(ctx, principal, resourceType, operation, scope)

Summarizes what resources and rolebindings this principal is allowed to create.

Callable by Admins+User.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principal** | **string**| Fully-qualified KafkaPrincipal string for a user or group. | 
**resourceType** | **string**| The type of resource to create or the type of resource to specify when creating a new rolebinding. | 
**operation** | **string**| \&quot;Create\&quot; for creating an actual resource, \&quot;AlterAccess\&quot; for creating a rolebinding for a user.  | 
**scope** | [**Scope**](Scope.md)|  | 

### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

