# \PrivateRBACUICachedUserStoreInformationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGroupsList**](PrivateRBACUICachedUserStoreInformationApi.md#UsersGroupsList) | **Post** /security/1.0/rbac/principals | List of MDS cached users and groups.



## UsersGroupsList

> []string UsersGroupsList(ctx, scope, optional)

List of MDS cached users and groups.

For use by a rolebinding admin on the provided scope.  Callable by Admins+ResourceOwners, but not broker super.users. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | [**Scope**](Scope.md)|  | 
 **optional** | ***UsersGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsersGroupsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| The type of principals requested. | 

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

