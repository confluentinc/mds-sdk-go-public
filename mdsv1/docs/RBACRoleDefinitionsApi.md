# \RBACRoleDefinitionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoleDetail**](RBACRoleDefinitionsApi.md#RoleDetail) | **Get** /security/1.0/roles/{roleName} | List the resourceType and operations allowed for a given role.
[**Rolenames**](RBACRoleDefinitionsApi.md#Rolenames) | **Get** /security/1.0/roleNames | Returns the names of all the roles defined in the system.
[**Roles**](RBACRoleDefinitionsApi.md#Roles) | **Get** /security/1.0/roles | Returns all the roles defined in the system.



## RoleDetail

> Role RoleDetail(ctx, roleName)

List the resourceType and operations allowed for a given role.

Callable by LDAP users.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string**| Role name to look up. | 

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

> []string Rolenames(ctx, )

Returns the names of all the roles defined in the system.

For information and developer purposes.  Callable by LDAP users. 

### Required Parameters

This endpoint does not need any parameter.

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

> []Role Roles(ctx, )

Returns all the roles defined in the system.

For information and developer purposes.  Callable by LDAP users. 

### Required Parameters

This endpoint does not need any parameter.

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

