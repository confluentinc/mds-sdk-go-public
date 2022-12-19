# \AuthorizationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authorize**](AuthorizationApi.md#Authorize) | **Put** /security/1.0/authorize | Authorize operations against resourceType for a given user.



## Authorize

> []string Authorize(ctx, authorizeRequest)

Authorize operations against resourceType for a given user.

Callable by Admins+User.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorizeRequest** | [**AuthorizeRequest**](AuthorizeRequest.md)|  | 

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

