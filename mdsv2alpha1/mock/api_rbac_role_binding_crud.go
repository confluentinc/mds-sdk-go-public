// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_rbac_role_binding_crud.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_mds_sdk_go_mdsv2alpha1 "github.com/confluentinc/mds-sdk-go/mdsv2alpha1"
)

// RBACRoleBindingCRUDApi is a mock of RBACRoleBindingCRUDApi interface
type RBACRoleBindingCRUDApi struct {
	lockAddRoleForPrincipal sync.Mutex
	AddRoleForPrincipalFunc func(ctx context.Context, principal, roleName string, scope github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope) (*net_http.Response, error)

	lockAddRoleResourcesForPrincipal sync.Mutex
	AddRoleResourcesForPrincipalFunc func(ctx context.Context, principal, roleName string, resourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest) (*net_http.Response, error)

	lockDeleteAllRolesForPrincipal sync.Mutex
	DeleteAllRolesForPrincipalFunc func(ctx context.Context, principal string, scope github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope) (*net_http.Response, error)

	lockDeleteRoleForPrincipal sync.Mutex
	DeleteRoleForPrincipalFunc func(ctx context.Context, principal, roleName string, scope github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope) (*net_http.Response, error)

	lockGetRoleResourcesForPrincipal sync.Mutex
	GetRoleResourcesForPrincipalFunc func(ctx context.Context, principal, roleName string, scope github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope) ([]github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcePattern, *net_http.Response, error)

	lockRemoveRoleResourcesForPrincipal sync.Mutex
	RemoveRoleResourcesForPrincipalFunc func(ctx context.Context, principal, roleName string, resourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest) (*net_http.Response, error)

	lockSetRoleResourcesForPrincipal sync.Mutex
	SetRoleResourcesForPrincipalFunc func(ctx context.Context, principal, roleName string, resourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest) (*net_http.Response, error)

	calls struct {
		AddRoleForPrincipal []struct {
			Ctx       context.Context
			Principal string
			RoleName  string
			Scope     github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope
		}
		AddRoleResourcesForPrincipal []struct {
			Ctx              context.Context
			Principal        string
			RoleName         string
			ResourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest
		}
		DeleteAllRolesForPrincipal []struct {
			Ctx       context.Context
			Principal string
			Scope     github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope
		}
		DeleteRoleForPrincipal []struct {
			Ctx       context.Context
			Principal string
			RoleName  string
			Scope     github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope
		}
		GetRoleResourcesForPrincipal []struct {
			Ctx       context.Context
			Principal string
			RoleName  string
			Scope     github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope
		}
		RemoveRoleResourcesForPrincipal []struct {
			Ctx              context.Context
			Principal        string
			RoleName         string
			ResourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest
		}
		SetRoleResourcesForPrincipal []struct {
			Ctx              context.Context
			Principal        string
			RoleName         string
			ResourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest
		}
	}
}

// AddRoleForPrincipal mocks base method by wrapping the associated func.
func (m *RBACRoleBindingCRUDApi) AddRoleForPrincipal(ctx context.Context, principal, roleName string, scope github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope) (*net_http.Response, error) {
	m.lockAddRoleForPrincipal.Lock()
	defer m.lockAddRoleForPrincipal.Unlock()

	if m.AddRoleForPrincipalFunc == nil {
		panic("mocker: RBACRoleBindingCRUDApi.AddRoleForPrincipalFunc is nil but RBACRoleBindingCRUDApi.AddRoleForPrincipal was called.")
	}

	call := struct {
		Ctx       context.Context
		Principal string
		RoleName  string
		Scope     github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope
	}{
		Ctx:       ctx,
		Principal: principal,
		RoleName:  roleName,
		Scope:     scope,
	}

	m.calls.AddRoleForPrincipal = append(m.calls.AddRoleForPrincipal, call)

	return m.AddRoleForPrincipalFunc(ctx, principal, roleName, scope)
}

// AddRoleForPrincipalCalled returns true if AddRoleForPrincipal was called at least once.
func (m *RBACRoleBindingCRUDApi) AddRoleForPrincipalCalled() bool {
	m.lockAddRoleForPrincipal.Lock()
	defer m.lockAddRoleForPrincipal.Unlock()

	return len(m.calls.AddRoleForPrincipal) > 0
}

// AddRoleForPrincipalCalls returns the calls made to AddRoleForPrincipal.
func (m *RBACRoleBindingCRUDApi) AddRoleForPrincipalCalls() []struct {
	Ctx       context.Context
	Principal string
	RoleName  string
	Scope     github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope
} {
	m.lockAddRoleForPrincipal.Lock()
	defer m.lockAddRoleForPrincipal.Unlock()

	return m.calls.AddRoleForPrincipal
}

// AddRoleResourcesForPrincipal mocks base method by wrapping the associated func.
func (m *RBACRoleBindingCRUDApi) AddRoleResourcesForPrincipal(ctx context.Context, principal, roleName string, resourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest) (*net_http.Response, error) {
	m.lockAddRoleResourcesForPrincipal.Lock()
	defer m.lockAddRoleResourcesForPrincipal.Unlock()

	if m.AddRoleResourcesForPrincipalFunc == nil {
		panic("mocker: RBACRoleBindingCRUDApi.AddRoleResourcesForPrincipalFunc is nil but RBACRoleBindingCRUDApi.AddRoleResourcesForPrincipal was called.")
	}

	call := struct {
		Ctx              context.Context
		Principal        string
		RoleName         string
		ResourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest
	}{
		Ctx:              ctx,
		Principal:        principal,
		RoleName:         roleName,
		ResourcesRequest: resourcesRequest,
	}

	m.calls.AddRoleResourcesForPrincipal = append(m.calls.AddRoleResourcesForPrincipal, call)

	return m.AddRoleResourcesForPrincipalFunc(ctx, principal, roleName, resourcesRequest)
}

// AddRoleResourcesForPrincipalCalled returns true if AddRoleResourcesForPrincipal was called at least once.
func (m *RBACRoleBindingCRUDApi) AddRoleResourcesForPrincipalCalled() bool {
	m.lockAddRoleResourcesForPrincipal.Lock()
	defer m.lockAddRoleResourcesForPrincipal.Unlock()

	return len(m.calls.AddRoleResourcesForPrincipal) > 0
}

// AddRoleResourcesForPrincipalCalls returns the calls made to AddRoleResourcesForPrincipal.
func (m *RBACRoleBindingCRUDApi) AddRoleResourcesForPrincipalCalls() []struct {
	Ctx              context.Context
	Principal        string
	RoleName         string
	ResourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest
} {
	m.lockAddRoleResourcesForPrincipal.Lock()
	defer m.lockAddRoleResourcesForPrincipal.Unlock()

	return m.calls.AddRoleResourcesForPrincipal
}

// DeleteAllRolesForPrincipal mocks base method by wrapping the associated func.
func (m *RBACRoleBindingCRUDApi) DeleteAllRolesForPrincipal(ctx context.Context, principal string, scope github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope) (*net_http.Response, error) {
	m.lockDeleteAllRolesForPrincipal.Lock()
	defer m.lockDeleteAllRolesForPrincipal.Unlock()

	if m.DeleteAllRolesForPrincipalFunc == nil {
		panic("mocker: RBACRoleBindingCRUDApi.DeleteAllRolesForPrincipalFunc is nil but RBACRoleBindingCRUDApi.DeleteAllRolesForPrincipal was called.")
	}

	call := struct {
		Ctx       context.Context
		Principal string
		Scope     github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope
	}{
		Ctx:       ctx,
		Principal: principal,
		Scope:     scope,
	}

	m.calls.DeleteAllRolesForPrincipal = append(m.calls.DeleteAllRolesForPrincipal, call)

	return m.DeleteAllRolesForPrincipalFunc(ctx, principal, scope)
}

// DeleteAllRolesForPrincipalCalled returns true if DeleteAllRolesForPrincipal was called at least once.
func (m *RBACRoleBindingCRUDApi) DeleteAllRolesForPrincipalCalled() bool {
	m.lockDeleteAllRolesForPrincipal.Lock()
	defer m.lockDeleteAllRolesForPrincipal.Unlock()

	return len(m.calls.DeleteAllRolesForPrincipal) > 0
}

// DeleteAllRolesForPrincipalCalls returns the calls made to DeleteAllRolesForPrincipal.
func (m *RBACRoleBindingCRUDApi) DeleteAllRolesForPrincipalCalls() []struct {
	Ctx       context.Context
	Principal string
	Scope     github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope
} {
	m.lockDeleteAllRolesForPrincipal.Lock()
	defer m.lockDeleteAllRolesForPrincipal.Unlock()

	return m.calls.DeleteAllRolesForPrincipal
}

// DeleteRoleForPrincipal mocks base method by wrapping the associated func.
func (m *RBACRoleBindingCRUDApi) DeleteRoleForPrincipal(ctx context.Context, principal, roleName string, scope github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope) (*net_http.Response, error) {
	m.lockDeleteRoleForPrincipal.Lock()
	defer m.lockDeleteRoleForPrincipal.Unlock()

	if m.DeleteRoleForPrincipalFunc == nil {
		panic("mocker: RBACRoleBindingCRUDApi.DeleteRoleForPrincipalFunc is nil but RBACRoleBindingCRUDApi.DeleteRoleForPrincipal was called.")
	}

	call := struct {
		Ctx       context.Context
		Principal string
		RoleName  string
		Scope     github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope
	}{
		Ctx:       ctx,
		Principal: principal,
		RoleName:  roleName,
		Scope:     scope,
	}

	m.calls.DeleteRoleForPrincipal = append(m.calls.DeleteRoleForPrincipal, call)

	return m.DeleteRoleForPrincipalFunc(ctx, principal, roleName, scope)
}

// DeleteRoleForPrincipalCalled returns true if DeleteRoleForPrincipal was called at least once.
func (m *RBACRoleBindingCRUDApi) DeleteRoleForPrincipalCalled() bool {
	m.lockDeleteRoleForPrincipal.Lock()
	defer m.lockDeleteRoleForPrincipal.Unlock()

	return len(m.calls.DeleteRoleForPrincipal) > 0
}

// DeleteRoleForPrincipalCalls returns the calls made to DeleteRoleForPrincipal.
func (m *RBACRoleBindingCRUDApi) DeleteRoleForPrincipalCalls() []struct {
	Ctx       context.Context
	Principal string
	RoleName  string
	Scope     github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope
} {
	m.lockDeleteRoleForPrincipal.Lock()
	defer m.lockDeleteRoleForPrincipal.Unlock()

	return m.calls.DeleteRoleForPrincipal
}

// GetRoleResourcesForPrincipal mocks base method by wrapping the associated func.
func (m *RBACRoleBindingCRUDApi) GetRoleResourcesForPrincipal(ctx context.Context, principal, roleName string, scope github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope) ([]github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcePattern, *net_http.Response, error) {
	m.lockGetRoleResourcesForPrincipal.Lock()
	defer m.lockGetRoleResourcesForPrincipal.Unlock()

	if m.GetRoleResourcesForPrincipalFunc == nil {
		panic("mocker: RBACRoleBindingCRUDApi.GetRoleResourcesForPrincipalFunc is nil but RBACRoleBindingCRUDApi.GetRoleResourcesForPrincipal was called.")
	}

	call := struct {
		Ctx       context.Context
		Principal string
		RoleName  string
		Scope     github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope
	}{
		Ctx:       ctx,
		Principal: principal,
		RoleName:  roleName,
		Scope:     scope,
	}

	m.calls.GetRoleResourcesForPrincipal = append(m.calls.GetRoleResourcesForPrincipal, call)

	return m.GetRoleResourcesForPrincipalFunc(ctx, principal, roleName, scope)
}

// GetRoleResourcesForPrincipalCalled returns true if GetRoleResourcesForPrincipal was called at least once.
func (m *RBACRoleBindingCRUDApi) GetRoleResourcesForPrincipalCalled() bool {
	m.lockGetRoleResourcesForPrincipal.Lock()
	defer m.lockGetRoleResourcesForPrincipal.Unlock()

	return len(m.calls.GetRoleResourcesForPrincipal) > 0
}

// GetRoleResourcesForPrincipalCalls returns the calls made to GetRoleResourcesForPrincipal.
func (m *RBACRoleBindingCRUDApi) GetRoleResourcesForPrincipalCalls() []struct {
	Ctx       context.Context
	Principal string
	RoleName  string
	Scope     github_com_confluentinc_mds_sdk_go_mdsv2alpha1.Scope
} {
	m.lockGetRoleResourcesForPrincipal.Lock()
	defer m.lockGetRoleResourcesForPrincipal.Unlock()

	return m.calls.GetRoleResourcesForPrincipal
}

// RemoveRoleResourcesForPrincipal mocks base method by wrapping the associated func.
func (m *RBACRoleBindingCRUDApi) RemoveRoleResourcesForPrincipal(ctx context.Context, principal, roleName string, resourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest) (*net_http.Response, error) {
	m.lockRemoveRoleResourcesForPrincipal.Lock()
	defer m.lockRemoveRoleResourcesForPrincipal.Unlock()

	if m.RemoveRoleResourcesForPrincipalFunc == nil {
		panic("mocker: RBACRoleBindingCRUDApi.RemoveRoleResourcesForPrincipalFunc is nil but RBACRoleBindingCRUDApi.RemoveRoleResourcesForPrincipal was called.")
	}

	call := struct {
		Ctx              context.Context
		Principal        string
		RoleName         string
		ResourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest
	}{
		Ctx:              ctx,
		Principal:        principal,
		RoleName:         roleName,
		ResourcesRequest: resourcesRequest,
	}

	m.calls.RemoveRoleResourcesForPrincipal = append(m.calls.RemoveRoleResourcesForPrincipal, call)

	return m.RemoveRoleResourcesForPrincipalFunc(ctx, principal, roleName, resourcesRequest)
}

// RemoveRoleResourcesForPrincipalCalled returns true if RemoveRoleResourcesForPrincipal was called at least once.
func (m *RBACRoleBindingCRUDApi) RemoveRoleResourcesForPrincipalCalled() bool {
	m.lockRemoveRoleResourcesForPrincipal.Lock()
	defer m.lockRemoveRoleResourcesForPrincipal.Unlock()

	return len(m.calls.RemoveRoleResourcesForPrincipal) > 0
}

// RemoveRoleResourcesForPrincipalCalls returns the calls made to RemoveRoleResourcesForPrincipal.
func (m *RBACRoleBindingCRUDApi) RemoveRoleResourcesForPrincipalCalls() []struct {
	Ctx              context.Context
	Principal        string
	RoleName         string
	ResourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest
} {
	m.lockRemoveRoleResourcesForPrincipal.Lock()
	defer m.lockRemoveRoleResourcesForPrincipal.Unlock()

	return m.calls.RemoveRoleResourcesForPrincipal
}

// SetRoleResourcesForPrincipal mocks base method by wrapping the associated func.
func (m *RBACRoleBindingCRUDApi) SetRoleResourcesForPrincipal(ctx context.Context, principal, roleName string, resourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest) (*net_http.Response, error) {
	m.lockSetRoleResourcesForPrincipal.Lock()
	defer m.lockSetRoleResourcesForPrincipal.Unlock()

	if m.SetRoleResourcesForPrincipalFunc == nil {
		panic("mocker: RBACRoleBindingCRUDApi.SetRoleResourcesForPrincipalFunc is nil but RBACRoleBindingCRUDApi.SetRoleResourcesForPrincipal was called.")
	}

	call := struct {
		Ctx              context.Context
		Principal        string
		RoleName         string
		ResourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest
	}{
		Ctx:              ctx,
		Principal:        principal,
		RoleName:         roleName,
		ResourcesRequest: resourcesRequest,
	}

	m.calls.SetRoleResourcesForPrincipal = append(m.calls.SetRoleResourcesForPrincipal, call)

	return m.SetRoleResourcesForPrincipalFunc(ctx, principal, roleName, resourcesRequest)
}

// SetRoleResourcesForPrincipalCalled returns true if SetRoleResourcesForPrincipal was called at least once.
func (m *RBACRoleBindingCRUDApi) SetRoleResourcesForPrincipalCalled() bool {
	m.lockSetRoleResourcesForPrincipal.Lock()
	defer m.lockSetRoleResourcesForPrincipal.Unlock()

	return len(m.calls.SetRoleResourcesForPrincipal) > 0
}

// SetRoleResourcesForPrincipalCalls returns the calls made to SetRoleResourcesForPrincipal.
func (m *RBACRoleBindingCRUDApi) SetRoleResourcesForPrincipalCalls() []struct {
	Ctx              context.Context
	Principal        string
	RoleName         string
	ResourcesRequest github_com_confluentinc_mds_sdk_go_mdsv2alpha1.ResourcesRequest
} {
	m.lockSetRoleResourcesForPrincipal.Lock()
	defer m.lockSetRoleResourcesForPrincipal.Unlock()

	return m.calls.SetRoleResourcesForPrincipal
}

// Reset resets the calls made to the mocked methods.
func (m *RBACRoleBindingCRUDApi) Reset() {
	m.lockAddRoleForPrincipal.Lock()
	m.calls.AddRoleForPrincipal = nil
	m.lockAddRoleForPrincipal.Unlock()
	m.lockAddRoleResourcesForPrincipal.Lock()
	m.calls.AddRoleResourcesForPrincipal = nil
	m.lockAddRoleResourcesForPrincipal.Unlock()
	m.lockDeleteAllRolesForPrincipal.Lock()
	m.calls.DeleteAllRolesForPrincipal = nil
	m.lockDeleteAllRolesForPrincipal.Unlock()
	m.lockDeleteRoleForPrincipal.Lock()
	m.calls.DeleteRoleForPrincipal = nil
	m.lockDeleteRoleForPrincipal.Unlock()
	m.lockGetRoleResourcesForPrincipal.Lock()
	m.calls.GetRoleResourcesForPrincipal = nil
	m.lockGetRoleResourcesForPrincipal.Unlock()
	m.lockRemoveRoleResourcesForPrincipal.Lock()
	m.calls.RemoveRoleResourcesForPrincipal = nil
	m.lockRemoveRoleResourcesForPrincipal.Unlock()
	m.lockSetRoleResourcesForPrincipal.Lock()
	m.calls.SetRoleResourcesForPrincipal = nil
	m.lockSetRoleResourcesForPrincipal.Unlock()
}