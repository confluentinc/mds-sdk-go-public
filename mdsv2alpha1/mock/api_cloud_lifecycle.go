// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_cloud_lifecycle.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1 "github.com/confluentinc/mds-sdk-go-public/mdsv2alpha1"
)

// CloudLifecycleApi is a mock of CloudLifecycleApi interface
type CloudLifecycleApi struct {
	lockDuplicateRolesForOrg sync.Mutex
	DuplicateRolesForOrgFunc func(ctx context.Context, sourceOrgId string, duplicateRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.DuplicateRequest) (*net_http.Response, error)

	lockRemoveAllRoleBindingsForScope sync.Mutex
	RemoveAllRoleBindingsForScopeFunc func(ctx context.Context, transactionId string, scope github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.Scope) (*net_http.Response, error)

	lockScopeUndelete sync.Mutex
	ScopeUndeleteFunc func(ctx context.Context, scopeUndeleteRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.ScopeUndeleteRequest) ([]string, *net_http.Response, error)

	lockUserUndelete sync.Mutex
	UserUndeleteFunc func(ctx context.Context, userUndeleteRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.UserUndeleteRequest) ([]string, *net_http.Response, error)

	calls struct {
		DuplicateRolesForOrg []struct {
			Ctx              context.Context
			SourceOrgId      string
			DuplicateRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.DuplicateRequest
		}
		RemoveAllRoleBindingsForScope []struct {
			Ctx           context.Context
			TransactionId string
			Scope         github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.Scope
		}
		ScopeUndelete []struct {
			Ctx                  context.Context
			ScopeUndeleteRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.ScopeUndeleteRequest
		}
		UserUndelete []struct {
			Ctx                 context.Context
			UserUndeleteRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.UserUndeleteRequest
		}
	}
}

// DuplicateRolesForOrg mocks base method by wrapping the associated func.
func (m *CloudLifecycleApi) DuplicateRolesForOrg(ctx context.Context, sourceOrgId string, duplicateRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.DuplicateRequest) (*net_http.Response, error) {
	m.lockDuplicateRolesForOrg.Lock()
	defer m.lockDuplicateRolesForOrg.Unlock()

	if m.DuplicateRolesForOrgFunc == nil {
		panic("mocker: CloudLifecycleApi.DuplicateRolesForOrgFunc is nil but CloudLifecycleApi.DuplicateRolesForOrg was called.")
	}

	call := struct {
		Ctx              context.Context
		SourceOrgId      string
		DuplicateRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.DuplicateRequest
	}{
		Ctx:              ctx,
		SourceOrgId:      sourceOrgId,
		DuplicateRequest: duplicateRequest,
	}

	m.calls.DuplicateRolesForOrg = append(m.calls.DuplicateRolesForOrg, call)

	return m.DuplicateRolesForOrgFunc(ctx, sourceOrgId, duplicateRequest)
}

// DuplicateRolesForOrgCalled returns true if DuplicateRolesForOrg was called at least once.
func (m *CloudLifecycleApi) DuplicateRolesForOrgCalled() bool {
	m.lockDuplicateRolesForOrg.Lock()
	defer m.lockDuplicateRolesForOrg.Unlock()

	return len(m.calls.DuplicateRolesForOrg) > 0
}

// DuplicateRolesForOrgCalls returns the calls made to DuplicateRolesForOrg.
func (m *CloudLifecycleApi) DuplicateRolesForOrgCalls() []struct {
	Ctx              context.Context
	SourceOrgId      string
	DuplicateRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.DuplicateRequest
} {
	m.lockDuplicateRolesForOrg.Lock()
	defer m.lockDuplicateRolesForOrg.Unlock()

	return m.calls.DuplicateRolesForOrg
}

// RemoveAllRoleBindingsForScope mocks base method by wrapping the associated func.
func (m *CloudLifecycleApi) RemoveAllRoleBindingsForScope(ctx context.Context, transactionId string, scope github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.Scope) (*net_http.Response, error) {
	m.lockRemoveAllRoleBindingsForScope.Lock()
	defer m.lockRemoveAllRoleBindingsForScope.Unlock()

	if m.RemoveAllRoleBindingsForScopeFunc == nil {
		panic("mocker: CloudLifecycleApi.RemoveAllRoleBindingsForScopeFunc is nil but CloudLifecycleApi.RemoveAllRoleBindingsForScope was called.")
	}

	call := struct {
		Ctx           context.Context
		TransactionId string
		Scope         github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.Scope
	}{
		Ctx:           ctx,
		TransactionId: transactionId,
		Scope:         scope,
	}

	m.calls.RemoveAllRoleBindingsForScope = append(m.calls.RemoveAllRoleBindingsForScope, call)

	return m.RemoveAllRoleBindingsForScopeFunc(ctx, transactionId, scope)
}

// RemoveAllRoleBindingsForScopeCalled returns true if RemoveAllRoleBindingsForScope was called at least once.
func (m *CloudLifecycleApi) RemoveAllRoleBindingsForScopeCalled() bool {
	m.lockRemoveAllRoleBindingsForScope.Lock()
	defer m.lockRemoveAllRoleBindingsForScope.Unlock()

	return len(m.calls.RemoveAllRoleBindingsForScope) > 0
}

// RemoveAllRoleBindingsForScopeCalls returns the calls made to RemoveAllRoleBindingsForScope.
func (m *CloudLifecycleApi) RemoveAllRoleBindingsForScopeCalls() []struct {
	Ctx           context.Context
	TransactionId string
	Scope         github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.Scope
} {
	m.lockRemoveAllRoleBindingsForScope.Lock()
	defer m.lockRemoveAllRoleBindingsForScope.Unlock()

	return m.calls.RemoveAllRoleBindingsForScope
}

// ScopeUndelete mocks base method by wrapping the associated func.
func (m *CloudLifecycleApi) ScopeUndelete(ctx context.Context, scopeUndeleteRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.ScopeUndeleteRequest) ([]string, *net_http.Response, error) {
	m.lockScopeUndelete.Lock()
	defer m.lockScopeUndelete.Unlock()

	if m.ScopeUndeleteFunc == nil {
		panic("mocker: CloudLifecycleApi.ScopeUndeleteFunc is nil but CloudLifecycleApi.ScopeUndelete was called.")
	}

	call := struct {
		Ctx                  context.Context
		ScopeUndeleteRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.ScopeUndeleteRequest
	}{
		Ctx:                  ctx,
		ScopeUndeleteRequest: scopeUndeleteRequest,
	}

	m.calls.ScopeUndelete = append(m.calls.ScopeUndelete, call)

	return m.ScopeUndeleteFunc(ctx, scopeUndeleteRequest)
}

// ScopeUndeleteCalled returns true if ScopeUndelete was called at least once.
func (m *CloudLifecycleApi) ScopeUndeleteCalled() bool {
	m.lockScopeUndelete.Lock()
	defer m.lockScopeUndelete.Unlock()

	return len(m.calls.ScopeUndelete) > 0
}

// ScopeUndeleteCalls returns the calls made to ScopeUndelete.
func (m *CloudLifecycleApi) ScopeUndeleteCalls() []struct {
	Ctx                  context.Context
	ScopeUndeleteRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.ScopeUndeleteRequest
} {
	m.lockScopeUndelete.Lock()
	defer m.lockScopeUndelete.Unlock()

	return m.calls.ScopeUndelete
}

// UserUndelete mocks base method by wrapping the associated func.
func (m *CloudLifecycleApi) UserUndelete(ctx context.Context, userUndeleteRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.UserUndeleteRequest) ([]string, *net_http.Response, error) {
	m.lockUserUndelete.Lock()
	defer m.lockUserUndelete.Unlock()

	if m.UserUndeleteFunc == nil {
		panic("mocker: CloudLifecycleApi.UserUndeleteFunc is nil but CloudLifecycleApi.UserUndelete was called.")
	}

	call := struct {
		Ctx                 context.Context
		UserUndeleteRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.UserUndeleteRequest
	}{
		Ctx:                 ctx,
		UserUndeleteRequest: userUndeleteRequest,
	}

	m.calls.UserUndelete = append(m.calls.UserUndelete, call)

	return m.UserUndeleteFunc(ctx, userUndeleteRequest)
}

// UserUndeleteCalled returns true if UserUndelete was called at least once.
func (m *CloudLifecycleApi) UserUndeleteCalled() bool {
	m.lockUserUndelete.Lock()
	defer m.lockUserUndelete.Unlock()

	return len(m.calls.UserUndelete) > 0
}

// UserUndeleteCalls returns the calls made to UserUndelete.
func (m *CloudLifecycleApi) UserUndeleteCalls() []struct {
	Ctx                 context.Context
	UserUndeleteRequest github_com_confluentinc_mds_sdk_go_public_mdsv2alpha1.UserUndeleteRequest
} {
	m.lockUserUndelete.Lock()
	defer m.lockUserUndelete.Unlock()

	return m.calls.UserUndelete
}

// Reset resets the calls made to the mocked methods.
func (m *CloudLifecycleApi) Reset() {
	m.lockDuplicateRolesForOrg.Lock()
	m.calls.DuplicateRolesForOrg = nil
	m.lockDuplicateRolesForOrg.Unlock()
	m.lockRemoveAllRoleBindingsForScope.Lock()
	m.calls.RemoveAllRoleBindingsForScope = nil
	m.lockRemoveAllRoleBindingsForScope.Unlock()
	m.lockScopeUndelete.Lock()
	m.calls.ScopeUndelete = nil
	m.lockScopeUndelete.Unlock()
	m.lockUserUndelete.Lock()
	m.calls.UserUndelete = nil
	m.lockUserUndelete.Unlock()
}
