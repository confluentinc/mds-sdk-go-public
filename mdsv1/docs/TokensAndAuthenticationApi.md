# \TokensAndAuthenticationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetToken**](TokensAndAuthenticationApi.md#GetToken) | **Get** /security/1.0/authenticate | Get a bearer token.



## GetToken

> AuthenticationResponse GetToken(ctx, )

Get a bearer token.

Callable by LDAP users.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**AuthenticationResponse**](AuthenticationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

