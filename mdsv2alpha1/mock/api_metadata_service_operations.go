// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_metadata_service_operations.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"
)

// MetadataServiceOperationsApi is a mock of MetadataServiceOperationsApi interface
type MetadataServiceOperationsApi struct {
	lockActivenodes sync.Mutex
	ActivenodesFunc func(ctx context.Context, protocol string) ([]string, *net_http.Response, error)

	lockMetadataClusterId sync.Mutex
	MetadataClusterIdFunc func(ctx context.Context) (string, *net_http.Response, error)

	calls struct {
		Activenodes []struct {
			Ctx      context.Context
			Protocol string
		}
		MetadataClusterId []struct {
			Ctx context.Context
		}
	}
}

// Activenodes mocks base method by wrapping the associated func.
func (m *MetadataServiceOperationsApi) Activenodes(ctx context.Context, protocol string) ([]string, *net_http.Response, error) {
	m.lockActivenodes.Lock()
	defer m.lockActivenodes.Unlock()

	if m.ActivenodesFunc == nil {
		panic("mocker: MetadataServiceOperationsApi.ActivenodesFunc is nil but MetadataServiceOperationsApi.Activenodes was called.")
	}

	call := struct {
		Ctx      context.Context
		Protocol string
	}{
		Ctx:      ctx,
		Protocol: protocol,
	}

	m.calls.Activenodes = append(m.calls.Activenodes, call)

	return m.ActivenodesFunc(ctx, protocol)
}

// ActivenodesCalled returns true if Activenodes was called at least once.
func (m *MetadataServiceOperationsApi) ActivenodesCalled() bool {
	m.lockActivenodes.Lock()
	defer m.lockActivenodes.Unlock()

	return len(m.calls.Activenodes) > 0
}

// ActivenodesCalls returns the calls made to Activenodes.
func (m *MetadataServiceOperationsApi) ActivenodesCalls() []struct {
	Ctx      context.Context
	Protocol string
} {
	m.lockActivenodes.Lock()
	defer m.lockActivenodes.Unlock()

	return m.calls.Activenodes
}

// MetadataClusterId mocks base method by wrapping the associated func.
func (m *MetadataServiceOperationsApi) MetadataClusterId(ctx context.Context) (string, *net_http.Response, error) {
	m.lockMetadataClusterId.Lock()
	defer m.lockMetadataClusterId.Unlock()

	if m.MetadataClusterIdFunc == nil {
		panic("mocker: MetadataServiceOperationsApi.MetadataClusterIdFunc is nil but MetadataServiceOperationsApi.MetadataClusterId was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.MetadataClusterId = append(m.calls.MetadataClusterId, call)

	return m.MetadataClusterIdFunc(ctx)
}

// MetadataClusterIdCalled returns true if MetadataClusterId was called at least once.
func (m *MetadataServiceOperationsApi) MetadataClusterIdCalled() bool {
	m.lockMetadataClusterId.Lock()
	defer m.lockMetadataClusterId.Unlock()

	return len(m.calls.MetadataClusterId) > 0
}

// MetadataClusterIdCalls returns the calls made to MetadataClusterId.
func (m *MetadataServiceOperationsApi) MetadataClusterIdCalls() []struct {
	Ctx context.Context
} {
	m.lockMetadataClusterId.Lock()
	defer m.lockMetadataClusterId.Unlock()

	return m.calls.MetadataClusterId
}

// Reset resets the calls made to the mocked methods.
func (m *MetadataServiceOperationsApi) Reset() {
	m.lockActivenodes.Lock()
	m.calls.Activenodes = nil
	m.lockActivenodes.Unlock()
	m.lockMetadataClusterId.Lock()
	m.calls.MetadataClusterId = nil
	m.lockMetadataClusterId.Unlock()
}
