/*
 * MDS API
 *
 * ## Confluent Metadata API - Swagger UI --- This tool (SwaggerUI) and the Open API spec file are provided _for development / test purposes only_:  - **Do _not_ enable in Production.** - **This tool only works with HTTP.**  ### Authenticating Authentication is performed by HTTP Basic Auth or by presenting a bearer token. In this UI, click **Authorize** to enter credentials.  To get a bearer token, first call the authenticate endpoint with basic auth, and then extract the auth_token part of the request, and pass that as the bearer token.  ### Access Restrictions - Who can call what?  Some endpoints can be called by any authenticated user, while others can only be called by \"admins\". Additionally, many of the endpoints in the API involve two users: the user who is calling the endpoint (the \"calling\" principal) and the user that the API call is about (the \"target\" principal).  Example: User \"alice\", who has the UserAdmin role, and is identifed by her basic auth credentials or a bearer token, calls the CRUD endpoint to modify role bindings about user \"bob\".  To document what access restrictions each endpoint has, use the following legend, which lists access in order from least restrictive to most restrictive:  *  **LDAP**: Any authenticated LDAP user *  **Admins+User**: Admins or the user requesting information about themself *  **Admins+ResourceOwners**: Admins or users with ResourceOwner role *  **Admins+AclUsers**: Admins or the user having the required ACL permissions *  **Admins**: Admins only, which can be UserAdmin, SystemAdmin, broker super.user, or SecurityAdmin as \"Read\"  ### Overview of Responses  **Valid**  * 200 - Successful call with a return body. * 204 - Sucessuful call with **no** return body.  **Errors**  * 400 - Invalid request.  JSON parsing error, or otherwise incorrect request. * 401 - Not Authenticated.  You need to pass valid basic auth credentials or a user bearer token. * 403 - Not Authorized.  Valid request, but you aren't authorized to perform the requested action. * 404 - Invalid URL.  If you get this error from the authenticate endpoint, it means bearer token authentication needs to be enabled in the configuration.     * ``confluent.metadata.server.authentication.method=BEARER`` * 405 - Method Not Allowed.  Using the wrong HTTP method on a valid endpoint (for example, GET instead of POST). * 409 - Conflict.  Adding a new resource or updating an existing resource which would result in a conflict with existing state.     * can be thrown by Audit Logs and Cluster Registry APIs * 415 - Invalid Content Type.  Usually, not sending \"application/json\" as request body header. * 500 - Server Error.  ### Special Resource Types  Cluster and KsqlCluster are special ResourceTypes because they grant resource-scoped roles like ResourceOwner and DeveloperManage limited access to cluster-level operations (for example, Describe Configs on Kafka clusters). These special resource types only accept LITERAL patterns with the values \"kafka-cluster\" and \"kql-cluster\", respectively.  ### Private RBAC UI Endpoints  These endpoints were developed specifically to power the Confluent Control Center UI. As such, they only focus on those use cases and have only been tested in the context of Confluent Control Center. These endpoints have not been tested, nor has their usability been evaluated with respect to manual API calls.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mdsv1

import (
	_bytes "bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type KafkaACLManagementApi interface {

	/*
	 * AddAclBinding Creates Kafka ACLs for given AclBinding.
	 *
	 * Callable by Admins+AclUsers.
	 *
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param createAclRequest
	 */
	AddAclBinding(ctx _context.Context, createAclRequest CreateAclRequest) (*_nethttp.Response, error)

	/*
	 * RemoveAclBindings Deletes Kafka ACLs according to the supplied filter.
	 *
	 * Callable by Admins+AclUsers.
	 *
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param aclFilterRequest
	 * @return []AclBinding
	 */
	RemoveAclBindings(ctx _context.Context, aclFilterRequest AclFilterRequest) ([]AclBinding, *_nethttp.Response, error)

	/*
	 * SearchAclBinding Lists Kafka ACLs according to the supplied filter.
	 *
	 * Callable by Admins+AclUsers.
	 *
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param aclFilterRequest
	 * @return []AclBinding
	 */
	SearchAclBinding(ctx _context.Context, aclFilterRequest AclFilterRequest) ([]AclBinding, *_nethttp.Response, error)
}

// KafkaACLManagementApiService KafkaACLManagementApi service
type KafkaACLManagementApiService service

/*
 * AddAclBinding Creates Kafka ACLs for given AclBinding.
 *
 * Callable by Admins+AclUsers.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param createAclRequest
 */
func (a *KafkaACLManagementApiService) AddAclBinding(ctx _context.Context, createAclRequest CreateAclRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/security/1.0/acls"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &createAclRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
 * RemoveAclBindings Deletes Kafka ACLs according to the supplied filter.
 *
 * Callable by Admins+AclUsers.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param aclFilterRequest
 * @return []AclBinding
 */
func (a *KafkaACLManagementApiService) RemoveAclBindings(ctx _context.Context, aclFilterRequest AclFilterRequest) ([]AclBinding, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []AclBinding
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/security/1.0/acls"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &aclFilterRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
 * SearchAclBinding Lists Kafka ACLs according to the supplied filter.
 *
 * Callable by Admins+AclUsers.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param aclFilterRequest
 * @return []AclBinding
 */
func (a *KafkaACLManagementApiService) SearchAclBinding(ctx _context.Context, aclFilterRequest AclFilterRequest) ([]AclBinding, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []AclBinding
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/security/1.0/acls:search"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &aclFilterRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
