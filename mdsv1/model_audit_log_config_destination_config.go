/*
 * MDS API
 *
 * ## Confluent Metadata API - Swagger UI --- This tool (SwaggerUI) and the Open API spec file are provided _for development / test purposes only_:  - **Do _not_ enable in Production.** - **This tool only works with HTTP.**  ### Authenticating Authentication is performed by HTTP Basic Auth or by presenting a bearer token. In this UI, click **Authorize** to enter credentials.  To get a bearer token, first call the authenticate endpoint with basic auth, and then extract the auth_token part of the request, and pass that as the bearer token.  ### Access Restrictions - Who can call what?  Some endpoints can be called by any authenticated user, while others can only be called by \"admins\". Additionally, many of the endpoints in the API involve two users: the user who is calling the endpoint (the \"calling\" principal) and the user that the API call is about (the \"target\" principal).  Example: User \"alice\", who has the UserAdmin role, and is identifed by her basic auth credentials or a bearer token, calls the CRUD endpoint to modify role bindings about user \"bob\".  To document what access restrictions each endpoint has, use the following legend, which lists access in order from least restrictive to most restrictive:  *  **LDAP**: Any authenticated LDAP user *  **Admins+User**: Admins or the user requesting information about themself *  **Admins+ResourceOwners**: Admins or users with ResourceOwner role *  **Admins+AclUsers**: Admins or the user having the required ACL permissions *  **Admins**: Admins only, which can be UserAdmin, SystemAdmin, broker super.user, or SecurityAdmin as \"Read\"  ### Overview of Responses  **Valid**  * 200 - Successful call with a return body. * 204 - Sucessuful call with **no** return body.  **Errors**  * 400 - Invalid request.  JSON parsing error, or otherwise incorrect request. * 401 - Not Authenticated.  You need to pass valid basic auth credentials or a user bearer token. * 403 - Not Authorized.  Valid request, but you aren't authorized to perform the requested action. * 404 - Invalid URL.  If you get this error from the authenticate endpoint, it means bearer token authentication needs to be enabled in the configuration.     * ``confluent.metadata.server.authentication.method=BEARER`` * 405 - Method Not Allowed.  Using the wrong HTTP method on a valid endpoint (for example, GET instead of POST). * 409 - Conflict.  Adding a new resource or updating an existing resource which would result in a conflict with existing state.     * can be thrown by Audit Logs and Cluster Registry APIs * 415 - Invalid Content Type.  Usually, not sending \"application/json\" as request body header. * 500 - Server Error.  ### Special Resource Types  Cluster and KsqlCluster are special ResourceTypes because they grant resource-scoped roles like ResourceOwner and DeveloperManage limited access to cluster-level operations (for example, Describe Configs on Kafka clusters). These special resource types only accept LITERAL patterns with the values \"kafka-cluster\" and \"kql-cluster\", respectively.  ### Private RBAC UI Endpoints  These endpoints were developed specifically to power the Confluent Control Center UI. As such, they only focus on those use cases and have only been tested in the context of Confluent Control Center. These endpoints have not been tested, nor has their usability been evaluated with respect to manual API calls.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mdsv1

// AuditLogConfigDestinationConfig struct for AuditLogConfigDestinationConfig
type AuditLogConfigDestinationConfig struct {
	// The number of milliseconds to keep events sent to this topic
	RetentionMs int64 `json:"retention_ms"`
}
