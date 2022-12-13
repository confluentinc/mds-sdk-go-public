// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_authorization.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_mds_sdk_go_mdsv1 "github.com/confluentinc/mds-sdk-go/mdsv1"
)

// AuthorizationApi is a mock of AuthorizationApi interface
type AuthorizationApi struct {
	lockAuthorize sync.Mutex
	AuthorizeFunc func(ctx context.Context, authorizeRequest github_com_confluentinc_mds_sdk_go_mdsv1.AuthorizeRequest) ([]string, *net_http.Response, error)

	calls struct {
		Authorize []struct {
			Ctx              context.Context
			AuthorizeRequest github_com_confluentinc_mds_sdk_go_mdsv1.AuthorizeRequest
		}
	}
}

// Authorize mocks base method by wrapping the associated func.
func (m *AuthorizationApi) Authorize(ctx context.Context, authorizeRequest github_com_confluentinc_mds_sdk_go_mdsv1.AuthorizeRequest) ([]string, *net_http.Response, error) {
	m.lockAuthorize.Lock()
	defer m.lockAuthorize.Unlock()

	if m.AuthorizeFunc == nil {
		panic("mocker: AuthorizationApi.AuthorizeFunc is nil but AuthorizationApi.Authorize was called.")
	}

	call := struct {
		Ctx              context.Context
		AuthorizeRequest github_com_confluentinc_mds_sdk_go_mdsv1.AuthorizeRequest
	}{
		Ctx:              ctx,
		AuthorizeRequest: authorizeRequest,
	}

	m.calls.Authorize = append(m.calls.Authorize, call)

	return m.AuthorizeFunc(ctx, authorizeRequest)
}

// AuthorizeCalled returns true if Authorize was called at least once.
func (m *AuthorizationApi) AuthorizeCalled() bool {
	m.lockAuthorize.Lock()
	defer m.lockAuthorize.Unlock()

	return len(m.calls.Authorize) > 0
}

// AuthorizeCalls returns the calls made to Authorize.
func (m *AuthorizationApi) AuthorizeCalls() []struct {
	Ctx              context.Context
	AuthorizeRequest github_com_confluentinc_mds_sdk_go_mdsv1.AuthorizeRequest
} {
	m.lockAuthorize.Lock()
	defer m.lockAuthorize.Unlock()

	return m.calls.Authorize
}

// Reset resets the calls made to the mocked methods.
func (m *AuthorizationApi) Reset() {
	m.lockAuthorize.Lock()
	m.calls.Authorize = nil
	m.lockAuthorize.Unlock()
}
