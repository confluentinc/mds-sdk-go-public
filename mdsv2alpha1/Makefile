ALL_SRC = $(shell find . -name \*.go -print)

.PHONY: mock
mock:
	mocker --prefix "" --dst mock/api_authorization.go --pkg mock api_authorization.go AuthorizationApi
	mocker --prefix "" --dst mock/api_cloud_lifecycle.go --pkg mock api_cloud_lifecycle.go CloudLifecycleApi
	mocker --prefix "" --dst mock/api_rbac_role_binding_crud.go --pkg mock api_rbac_role_binding_crud.go RBACRoleBindingCRUDApi
	mocker --prefix "" --dst mock/api_rbac_role_binding_summaries.go --pkg mock api_rbac_role_binding_summaries.go RBACRoleBindingSummariesApi
	mocker --prefix "" --dst mock/api_rbac_role_definitions.go --pkg mock api_rbac_role_definitions.go RBACRoleDefinitionsApi
	mocker --prefix "" --dst mock/api_tokens_and_authentication.go --pkg mock api_tokens_and_authentication.go TokensAndAuthenticationApi

.PHONY: fmt
# Format entire codebase
fmt:
	@gofmt -e -s -l -w $(ALL_SRC)
