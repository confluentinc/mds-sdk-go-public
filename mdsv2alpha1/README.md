# Go API client for mdsv2alpha1

## Confluent Metadata API - Swagger UI
---
This tool (SwaggerUI) and the Open API spec file are provided _for development / test
purposes only_:

- **Do _not_ enable in Production.**
- **This tool only works with HTTP.**

### Authenticating
Authentication is supported in this version only for testing purposes. Real users
get their tokens from cc-auth-service.

Authentication is performed by HTTP Basic Auth or by presenting a bearer token.
In this UI, click **Authorize** to enter credentials.

To get a bearer token, first call the authenticate endpoint with basic auth, and then extract the
auth_token part of the request, and pass that as the bearer token.

### Access Restrictions - Who can call what?

Some endpoints can be called by any authenticated user, while others can only be called by \"admins\".
Additionally, many of the endpoints in the API involve two users: the user who is calling the
endpoint (the \"calling\" principal) and the user that the API call is about (the \"target\" principal).

Example: User u-abc12, who has the UserAdmin role, and is identifed by her basic auth credentials or a
bearer token, calls the CRUD endpoint to modify role bindings about user u-45def.

To document what access restrictions each endpoint has, use the following legend, which
lists access in order from least restrictive to most restrictive:

*  **Users**: Any authenticated user
*  **Admins+Users**: Admins or the user requesting information about themself
*  **Admins**: Admins only, which can be UserAdmin, SystemAdmin, broker super.user, or SecurityAdmin as \"Read\"

### Overview of Responses

**Valid**

* 200 - Successful call with a return body.
* 204 - Sucessuful call with **no** return body.

**Errors**

* 400 - Invalid request.  JSON parsing error, or otherwise incorrect request.
* 401 - Not Authenticated.  You need to pass valid basic auth credentials or a user bearer token.
* 403 - Not Authorized.  Valid request, but you aren't authorized to perform the requested action.
* 404 - Invalid URL.  If you get this error from the authenticate endpoint, it means bearer token authentication needs to be enabled in the configuration.
    * ``confluent.metadata.server.authentication.method=BEARER``
* 405 - Method Not Allowed.  Using the wrong HTTP method on a valid endpoint (for example, GET instead of POST).
* 415 - Invalid Content Type.  Usually, not sending \"application/json\" as request body header.
* 500 - Server Error.

### Special Resource Types

Cluster and KsqlCluster are special ResourceTypes because they grant resource-scoped roles
like ResourceOwner and DeveloperManage limited access to cluster-level operations (for example,
Describe Configs on Kafka clusters).
These special resource types only accept LITERAL patterns with the values \"kafka-cluster\" and
\"kql-cluster\", respectively.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v2alpha1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoDeprecatedClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./mdsv2alpha1"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/security/v2alpha1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthorizationApi* | [**Authorize**](docs/AuthorizationApi.md#authorize) | **Put** /authorize | Authorize operations against resourceType for a given user.  Callable by Admins+User.
*CloudLifecycleApi* | [**DuplicateRolesForOrg**](docs/CloudLifecycleApi.md#duplicaterolesfororg) | **Post** /cloudlifecycle/rolebindings/{sourceOrgId}/copy | Duplicate all role bindings from the source organization to the destination organization
*CloudLifecycleApi* | [**RemoveAllRoleBindingsForScope**](docs/CloudLifecycleApi.md#removeallrolebindingsforscope) | **Delete** /cloudlifecycle/rolebindings | Delete all role bindings at the given scope
*CloudLifecycleApi* | [**ScopeUndelete**](docs/CloudLifecycleApi.md#scopeundelete) | **Post** /cloudlifecycle/scope/undelete | Undelete all role bindings for a given scope and reason.
*CloudLifecycleApi* | [**UserUndelete**](docs/CloudLifecycleApi.md#userundelete) | **Post** /cloudlifecycle/user/undelete | Undelete all role bindings for a given user, org, and reason.
*MetadataServiceOperationsApi* | [**Activenodes**](docs/MetadataServiceOperationsApi.md#activenodes) | **Get** /activenodes/{protocol} | Returns all the nodes running the Metadata Service REST API. Clients are expected to round robin call to these endpoints if they don&#39;t set up a load balancer in front of the Metadata Service nodes. Callable by authenticated users.
*MetadataServiceOperationsApi* | [**MetadataClusterId**](docs/MetadataServiceOperationsApi.md#metadataclusterid) | **Get** /metadataClusterId | Returns the ID of the Kafka cluster that MDS is running on.  Callable by LDAP users.
*RBACRoleBindingCRUDApi* | [**AddRoleForPrincipal**](docs/RBACRoleBindingCRUDApi.md#addroleforprincipal) | **Post** /principals/{principal}/roles/{roleName} | Binds the principal to a role for a specific cluster or in the given scope. Callable by Admins.
*RBACRoleBindingCRUDApi* | [**AddRoleResourcesForPrincipal**](docs/RBACRoleBindingCRUDApi.md#addroleresourcesforprincipal) | **Post** /principals/{principal}/roles/{roleName}/bindings | Incrementally grant the resources to the principal at the given scope/cluster using the given role.
*RBACRoleBindingCRUDApi* | [**DeleteAllRolesForPrincipal**](docs/RBACRoleBindingCRUDApi.md#deleteallrolesforprincipal) | **Delete** /principals/{principal}/roles | Remove all roles for the principal at the given scope and all contained scopes. Callable by Admins. 
*RBACRoleBindingCRUDApi* | [**DeleteRoleForPrincipal**](docs/RBACRoleBindingCRUDApi.md#deleteroleforprincipal) | **Delete** /principals/{principal}/roles/{roleName} | Remove the role from the principal at the given scope/cluster. No-op if the user doesn&#39;t have the role.  Callable by Admins. 
*RBACRoleBindingCRUDApi* | [**GetRoleResourcesForPrincipal**](docs/RBACRoleBindingCRUDApi.md#getroleresourcesforprincipal) | **Post** /principals/{principal}/roles/{roleName}/resources | Look up the rolebindings for the principal at the given scope/cluster using the given role.
*RBACRoleBindingCRUDApi* | [**RemoveRoleResourcesForPrincipal**](docs/RBACRoleBindingCRUDApi.md#removeroleresourcesforprincipal) | **Delete** /principals/{principal}/roles/{roleName}/bindings | Incrementally remove the resources from the principal at the given scope/cluster using the given role.
*RBACRoleBindingCRUDApi* | [**SetRoleResourcesForPrincipal**](docs/RBACRoleBindingCRUDApi.md#setroleresourcesforprincipal) | **Put** /principals/{principal}/roles/{roleName}/bindings | Overwrite existing resource grants.
*RBACRoleBindingSummariesApi* | [**LookupPrincipalsWithRole**](docs/RBACRoleBindingSummariesApi.md#lookupprincipalswithrole) | **Post** /lookup/role/{roleName} | Look up the KafkaPrincipals who have the given role for the given scope.  Callable by Admins.
*RBACRoleBindingSummariesApi* | [**LookupPrincipalsWithRoleOnResource**](docs/RBACRoleBindingSummariesApi.md#lookupprincipalswithroleonresource) | **Post** /lookup/role/{roleName}/resource/{resourceType}/name/{resourceName} | Look up the KafkaPrincipals who have the given role on the specified resource for the given scope.
*RBACRoleBindingSummariesApi* | [**ManagedNonResourceRoleBindingsAtScope**](docs/RBACRoleBindingSummariesApi.md#managednonresourcerolebindingsatscope) | **Post** /lookup/managed/rolebindings | Returns all non-resource rolebindings in the given scope for all users (not just the calling user) that the calling user has permission to see. A user can see, but not alter rolebindings for scopes that they have Describe access on, and alter rolebindings for scopes that they have Alter access on. Callable by Admins+Users. 
*RBACRoleBindingSummariesApi* | [**MyAllowedResources**](docs/RBACRoleBindingSummariesApi.md#myallowedresources) | **Post** /lookup/resources/{resourceType}/operation/{operation} | List all resource patterns of the specified resourceType that the caller is allowed to perform the specified operation on. If the caller is not allowed to perform the operation on any resources, the list will be empty. Overlapping resource patterns will be \&quot;squashed\&quot; to eliminate redundancy, for example if you have access on both the prefix \&quot;topic-*\&quot; and the literal \&quot;topic-1\&quot;, only \&quot;topic-*\&quot; will be returned.  Callable by Admins+Users. 
*RBACRoleBindingSummariesApi* | [**MyRoleBindings**](docs/RBACRoleBindingSummariesApi.md#myrolebindings) | **Post** /lookup/rolebindings/principal/{principal} | List all rolebindings for the specifed principal in the scope and all contained scopes. Be aware that this simply looks at the rolebinding data, and does not mean that the scopes actually exist. Callable by Admins+Users. 
*RBACRoleDefinitionsApi* | [**RoleDetail**](docs/RBACRoleDefinitionsApi.md#roledetail) | **Get** /roles/{roleName} | List the resourceType and operations allowed for a given role. Callable by Users. 
*RBACRoleDefinitionsApi* | [**Rolenames**](docs/RBACRoleDefinitionsApi.md#rolenames) | **Get** /roleNames | Returns the names of all the roles defined in the system. For information and developer purposes. Callable by Users. 
*RBACRoleDefinitionsApi* | [**Roles**](docs/RBACRoleDefinitionsApi.md#roles) | **Get** /roles | Returns all the public roles defined in the system.  For information and developer purposes. Callable by Users. 
*TokensAndAuthenticationApi* | [**GetToken**](docs/TokensAndAuthenticationApi.md#gettoken) | **Get** /authenticate | Get a token. This is here for testing purposes. Real users get tokens from cc-auth-service. 


## Documentation For Models

 - [AccessPolicy](docs/AccessPolicy.md)
 - [Action](docs/Action.md)
 - [AuthenticationResponse](docs/AuthenticationResponse.md)
 - [AuthorizeRequest](docs/AuthorizeRequest.md)
 - [DuplicateRequest](docs/DuplicateRequest.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ManagedRoleBinding](docs/ManagedRoleBinding.md)
 - [Operation](docs/Operation.md)
 - [ResourcePattern](docs/ResourcePattern.md)
 - [ResourcesRequest](docs/ResourcesRequest.md)
 - [Role](docs/Role.md)
 - [RoleBinding](docs/RoleBinding.md)
 - [Scope](docs/Scope.md)
 - [ScopeClusters](docs/ScopeClusters.md)
 - [ScopeRoleBindingMapping](docs/ScopeRoleBindingMapping.md)
 - [ScopeUndeleteRequest](docs/ScopeUndeleteRequest.md)
 - [UserUndeleteRequest](docs/UserUndeleteRequest.md)


## Documentation For Authorization



## basicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), mdsv2alpha1.ContextBasicAuth, mdsv2alpha1.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), mdsv2alpha1.ContextAccessToken, "BEARERTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```



## Author



