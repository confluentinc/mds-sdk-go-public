# \SSODeviceAuthorizationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckDeviceAuth**](SSODeviceAuthorizationApi.md#CheckDeviceAuth) | **Post** /security/1.0/oidc/device/check-auth | Provides information about current status of user authentication
[**ExtendDeviceAuth**](SSODeviceAuthorizationApi.md#ExtendDeviceAuth) | **Get** /security/1.0/oidc/device/extend-auth | Extend auth by generating a new token
[**Security10OidcDeviceAuthenticatePost**](SSODeviceAuthorizationApi.md#Security10OidcDeviceAuthenticatePost) | **Post** /security/1.0/oidc/device/authenticate | Provides user authentication details and device polling for authentication status



## CheckDeviceAuth

> CheckDeviceAuthResponse CheckDeviceAuth(ctx, checkDeviceAuthRequest)

Provides information about current status of user authentication

Checks if the user has authorized and logged in. If the user has authorized, then the response contains the auth token, and the complete flag is set to true.   If the user has not authorized in right time or some other error occurred, then the response contains the error details, and again complete flag is set to true.   Otherwise, in case of pending auth, the response contains the status, description and the complete flag is set to false. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkDeviceAuthRequest** | [**CheckDeviceAuthRequest**](CheckDeviceAuthRequest.md)|  | 

### Return type

[**CheckDeviceAuthResponse**](CheckDeviceAuthResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtendDeviceAuth

> ExtendAuthResponse ExtendDeviceAuth(ctx, )

Extend auth by generating a new token

Attempts to refresh the Confluent token if applicable, based on the provided JwtPrincipal.   If refresh token is configured to be used, the session is extended until expiry time of new ID token requested using the refresh token.  Else session is extended until min(`mex`, `currentTime`+ `sessionTokenExpiryConfig`) where `mex` is the claim already present in auth token.  The token cannot be extended beyond value of config `confluent.oidc.session.max.timeout.ms`.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ExtendAuthResponse**](ExtendAuthResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Security10OidcDeviceAuthenticatePost

> InitDeviceAuthResponse Security10OidcDeviceAuthenticatePost(ctx, )

Provides user authentication details and device polling for authentication status

This initiates user authentication in CLI.  Response contains  1. `user_code` and `verification_uri` which are used by user to authenticate from the identity provider (IdP).  2. `key` is used to poll the IdP to check if the user has authorized and logged in. 3. `interval` to hint client about polling frequency. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InitDeviceAuthResponse**](InitDeviceAuthResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

