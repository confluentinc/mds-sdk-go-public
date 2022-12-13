# \AuthorizationApi

All URIs are relative to *http://localhost/security/v2alpha1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authorize**](AuthorizationApi.md#Authorize) | **Put** /authorize | Authorize operations against resourceType for a given user.  Callable by Admins+User.



## Authorize

> []string Authorize(ctx, authorizeRequest, optional)

Authorize operations against resourceType for a given user.  Callable by Admins+User.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorizeRequest** | [**AuthorizeRequest**](AuthorizeRequest.md)|  | 
 **optional** | ***AuthorizeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorizeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**|  | 

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

