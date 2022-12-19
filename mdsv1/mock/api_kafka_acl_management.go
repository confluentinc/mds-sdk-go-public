// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_kafka_acl_management.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_mds_sdk_go_public_mdsv1 "github.com/confluentinc/mds-sdk-go-public/mdsv1"
)

// KafkaACLManagementApi is a mock of KafkaACLManagementApi interface
type KafkaACLManagementApi struct {
	lockAddAclBinding sync.Mutex
	AddAclBindingFunc func(ctx context.Context, createAclRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.CreateAclRequest) (*net_http.Response, error)

	lockRemoveAclBindings sync.Mutex
	RemoveAclBindingsFunc func(ctx context.Context, aclFilterRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.AclFilterRequest) ([]github_com_confluentinc_mds_sdk_go_public_mdsv1.AclBinding, *net_http.Response, error)

	lockSearchAclBinding sync.Mutex
	SearchAclBindingFunc func(ctx context.Context, aclFilterRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.AclFilterRequest) ([]github_com_confluentinc_mds_sdk_go_public_mdsv1.AclBinding, *net_http.Response, error)

	calls struct {
		AddAclBinding []struct {
			Ctx              context.Context
			CreateAclRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.CreateAclRequest
		}
		RemoveAclBindings []struct {
			Ctx              context.Context
			AclFilterRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.AclFilterRequest
		}
		SearchAclBinding []struct {
			Ctx              context.Context
			AclFilterRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.AclFilterRequest
		}
	}
}

// AddAclBinding mocks base method by wrapping the associated func.
func (m *KafkaACLManagementApi) AddAclBinding(ctx context.Context, createAclRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.CreateAclRequest) (*net_http.Response, error) {
	m.lockAddAclBinding.Lock()
	defer m.lockAddAclBinding.Unlock()

	if m.AddAclBindingFunc == nil {
		panic("mocker: KafkaACLManagementApi.AddAclBindingFunc is nil but KafkaACLManagementApi.AddAclBinding was called.")
	}

	call := struct {
		Ctx              context.Context
		CreateAclRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.CreateAclRequest
	}{
		Ctx:              ctx,
		CreateAclRequest: createAclRequest,
	}

	m.calls.AddAclBinding = append(m.calls.AddAclBinding, call)

	return m.AddAclBindingFunc(ctx, createAclRequest)
}

// AddAclBindingCalled returns true if AddAclBinding was called at least once.
func (m *KafkaACLManagementApi) AddAclBindingCalled() bool {
	m.lockAddAclBinding.Lock()
	defer m.lockAddAclBinding.Unlock()

	return len(m.calls.AddAclBinding) > 0
}

// AddAclBindingCalls returns the calls made to AddAclBinding.
func (m *KafkaACLManagementApi) AddAclBindingCalls() []struct {
	Ctx              context.Context
	CreateAclRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.CreateAclRequest
} {
	m.lockAddAclBinding.Lock()
	defer m.lockAddAclBinding.Unlock()

	return m.calls.AddAclBinding
}

// RemoveAclBindings mocks base method by wrapping the associated func.
func (m *KafkaACLManagementApi) RemoveAclBindings(ctx context.Context, aclFilterRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.AclFilterRequest) ([]github_com_confluentinc_mds_sdk_go_public_mdsv1.AclBinding, *net_http.Response, error) {
	m.lockRemoveAclBindings.Lock()
	defer m.lockRemoveAclBindings.Unlock()

	if m.RemoveAclBindingsFunc == nil {
		panic("mocker: KafkaACLManagementApi.RemoveAclBindingsFunc is nil but KafkaACLManagementApi.RemoveAclBindings was called.")
	}

	call := struct {
		Ctx              context.Context
		AclFilterRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.AclFilterRequest
	}{
		Ctx:              ctx,
		AclFilterRequest: aclFilterRequest,
	}

	m.calls.RemoveAclBindings = append(m.calls.RemoveAclBindings, call)

	return m.RemoveAclBindingsFunc(ctx, aclFilterRequest)
}

// RemoveAclBindingsCalled returns true if RemoveAclBindings was called at least once.
func (m *KafkaACLManagementApi) RemoveAclBindingsCalled() bool {
	m.lockRemoveAclBindings.Lock()
	defer m.lockRemoveAclBindings.Unlock()

	return len(m.calls.RemoveAclBindings) > 0
}

// RemoveAclBindingsCalls returns the calls made to RemoveAclBindings.
func (m *KafkaACLManagementApi) RemoveAclBindingsCalls() []struct {
	Ctx              context.Context
	AclFilterRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.AclFilterRequest
} {
	m.lockRemoveAclBindings.Lock()
	defer m.lockRemoveAclBindings.Unlock()

	return m.calls.RemoveAclBindings
}

// SearchAclBinding mocks base method by wrapping the associated func.
func (m *KafkaACLManagementApi) SearchAclBinding(ctx context.Context, aclFilterRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.AclFilterRequest) ([]github_com_confluentinc_mds_sdk_go_public_mdsv1.AclBinding, *net_http.Response, error) {
	m.lockSearchAclBinding.Lock()
	defer m.lockSearchAclBinding.Unlock()

	if m.SearchAclBindingFunc == nil {
		panic("mocker: KafkaACLManagementApi.SearchAclBindingFunc is nil but KafkaACLManagementApi.SearchAclBinding was called.")
	}

	call := struct {
		Ctx              context.Context
		AclFilterRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.AclFilterRequest
	}{
		Ctx:              ctx,
		AclFilterRequest: aclFilterRequest,
	}

	m.calls.SearchAclBinding = append(m.calls.SearchAclBinding, call)

	return m.SearchAclBindingFunc(ctx, aclFilterRequest)
}

// SearchAclBindingCalled returns true if SearchAclBinding was called at least once.
func (m *KafkaACLManagementApi) SearchAclBindingCalled() bool {
	m.lockSearchAclBinding.Lock()
	defer m.lockSearchAclBinding.Unlock()

	return len(m.calls.SearchAclBinding) > 0
}

// SearchAclBindingCalls returns the calls made to SearchAclBinding.
func (m *KafkaACLManagementApi) SearchAclBindingCalls() []struct {
	Ctx              context.Context
	AclFilterRequest github_com_confluentinc_mds_sdk_go_public_mdsv1.AclFilterRequest
} {
	m.lockSearchAclBinding.Lock()
	defer m.lockSearchAclBinding.Unlock()

	return m.calls.SearchAclBinding
}

// Reset resets the calls made to the mocked methods.
func (m *KafkaACLManagementApi) Reset() {
	m.lockAddAclBinding.Lock()
	m.calls.AddAclBinding = nil
	m.lockAddAclBinding.Unlock()
	m.lockRemoveAclBindings.Lock()
	m.calls.RemoveAclBindings = nil
	m.lockRemoveAclBindings.Unlock()
	m.lockSearchAclBinding.Lock()
	m.calls.SearchAclBinding = nil
	m.lockSearchAclBinding.Unlock()
}
