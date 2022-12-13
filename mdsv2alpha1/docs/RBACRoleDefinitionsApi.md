# \RBACRoleDefinitionsApi

All URIs are relative to *http://localhost/security/v2alpha1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoleDetail**](RBACRoleDefinitionsApi.md#RoleDetail) | **Get** /roles/{roleName} | List the resourceType and operations allowed for a given role. Callable by Users. 
[**Rolenames**](RBACRoleDefinitionsApi.md#Rolenames) | **Get** /roleNames | Returns the names of all the roles defined in the system. For information and developer purposes. Callable by Users. 
[**Roles**](RBACRoleDefinitionsApi.md#Roles) | **Get** /roles | Returns all the public roles defined in the system.  For information and developer purposes. Callable by Users. 



## RoleDetail

> Role RoleDetail(ctx, roleName, optional)

List the resourceType and operations allowed for a given role. Callable by Users. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string**| Role name to look up. | 
 **optional** | ***RoleDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RoleDetailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **optional.String**| Return the role definitions available in the specified namespace. If no namespace is specified, return the public roles. May be a comma-separated list.  | 

### Return type

[**Role**](Role.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Rolenames

> []string Rolenames(ctx, optional)

Returns the names of all the roles defined in the system. For information and developer purposes. Callable by Users. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RolenamesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RolenamesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **optional.String**| Return the role names available in the specified namespace. If no namespace is specified, return the public roles. May be a comma-separated list.  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Roles

> []Role Roles(ctx, optional)

Returns all the public roles defined in the system.  For information and developer purposes. Callable by Users. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RolesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RolesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **optional.String**| Return the role definitions available in the specified namespace. If no namespace is specified, return the public roles. May be a comma-separated list. | 

### Return type

[**[]Role**](Role.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

