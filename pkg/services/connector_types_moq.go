// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package services

import (
	"context"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api"
	apiErrors "github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
	"sync"
)

// Ensure, that ConnectorTypesServiceMock does implement ConnectorTypesService.
// If this is not the case, regenerate this file with moq.
var _ ConnectorTypesService = &ConnectorTypesServiceMock{}

// ConnectorTypesServiceMock is a mock implementation of ConnectorTypesService.
//
// 	func TestSomethingThatUsesConnectorTypesService(t *testing.T) {
//
// 		// make and configure a mocked ConnectorTypesService
// 		mockedConnectorTypesService := &ConnectorTypesServiceMock{
// 			DiscoverExtensionsFunc: func() error {
// 				panic("mock out the DiscoverExtensions method")
// 			},
// 			GetFunc: func(id string) (*api.ConnectorType, *apiErrors.ServiceError) {
// 				panic("mock out the Get method")
// 			},
// 			GetServiceAddressFunc: func(id string) (string, *apiErrors.ServiceError) {
// 				panic("mock out the GetServiceAddress method")
// 			},
// 			ListFunc: func(ctx context.Context, listArgs *ListArguments) (api.ConnectorTypeList, *api.PagingMeta, *apiErrors.ServiceError) {
// 				panic("mock out the List method")
// 			},
// 		}
//
// 		// use mockedConnectorTypesService in code that requires ConnectorTypesService
// 		// and then make assertions.
//
// 	}
type ConnectorTypesServiceMock struct {
	// DiscoverExtensionsFunc mocks the DiscoverExtensions method.
	DiscoverExtensionsFunc func() error

	// GetFunc mocks the Get method.
	GetFunc func(id string) (*api.ConnectorType, *apiErrors.ServiceError)

	// GetServiceAddressFunc mocks the GetServiceAddress method.
	GetServiceAddressFunc func(id string) (string, *apiErrors.ServiceError)

	// ListFunc mocks the List method.
	ListFunc func(ctx context.Context, listArgs *ListArguments) (api.ConnectorTypeList, *api.PagingMeta, *apiErrors.ServiceError)

	// calls tracks calls to the methods.
	calls struct {
		// DiscoverExtensions holds details about calls to the DiscoverExtensions method.
		DiscoverExtensions []struct {
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// ID is the id argument value.
			ID string
		}
		// GetServiceAddress holds details about calls to the GetServiceAddress method.
		GetServiceAddress []struct {
			// ID is the id argument value.
			ID string
		}
		// List holds details about calls to the List method.
		List []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ListArgs is the listArgs argument value.
			ListArgs *ListArguments
		}
	}
	lockDiscoverExtensions sync.RWMutex
	lockGet                sync.RWMutex
	lockGetServiceAddress  sync.RWMutex
	lockList               sync.RWMutex
}

// DiscoverExtensions calls DiscoverExtensionsFunc.
func (mock *ConnectorTypesServiceMock) DiscoverExtensions() error {
	if mock.DiscoverExtensionsFunc == nil {
		panic("ConnectorTypesServiceMock.DiscoverExtensionsFunc: method is nil but ConnectorTypesService.DiscoverExtensions was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDiscoverExtensions.Lock()
	mock.calls.DiscoverExtensions = append(mock.calls.DiscoverExtensions, callInfo)
	mock.lockDiscoverExtensions.Unlock()
	return mock.DiscoverExtensionsFunc()
}

// DiscoverExtensionsCalls gets all the calls that were made to DiscoverExtensions.
// Check the length with:
//     len(mockedConnectorTypesService.DiscoverExtensionsCalls())
func (mock *ConnectorTypesServiceMock) DiscoverExtensionsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDiscoverExtensions.RLock()
	calls = mock.calls.DiscoverExtensions
	mock.lockDiscoverExtensions.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *ConnectorTypesServiceMock) Get(id string) (*api.ConnectorType, *apiErrors.ServiceError) {
	if mock.GetFunc == nil {
		panic("ConnectorTypesServiceMock.GetFunc: method is nil but ConnectorTypesService.Get was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(id)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedConnectorTypesService.GetCalls())
func (mock *ConnectorTypesServiceMock) GetCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// GetServiceAddress calls GetServiceAddressFunc.
func (mock *ConnectorTypesServiceMock) GetServiceAddress(id string) (string, *apiErrors.ServiceError) {
	if mock.GetServiceAddressFunc == nil {
		panic("ConnectorTypesServiceMock.GetServiceAddressFunc: method is nil but ConnectorTypesService.GetServiceAddress was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGetServiceAddress.Lock()
	mock.calls.GetServiceAddress = append(mock.calls.GetServiceAddress, callInfo)
	mock.lockGetServiceAddress.Unlock()
	return mock.GetServiceAddressFunc(id)
}

// GetServiceAddressCalls gets all the calls that were made to GetServiceAddress.
// Check the length with:
//     len(mockedConnectorTypesService.GetServiceAddressCalls())
func (mock *ConnectorTypesServiceMock) GetServiceAddressCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetServiceAddress.RLock()
	calls = mock.calls.GetServiceAddress
	mock.lockGetServiceAddress.RUnlock()
	return calls
}

// List calls ListFunc.
func (mock *ConnectorTypesServiceMock) List(ctx context.Context, listArgs *ListArguments) (api.ConnectorTypeList, *api.PagingMeta, *apiErrors.ServiceError) {
	if mock.ListFunc == nil {
		panic("ConnectorTypesServiceMock.ListFunc: method is nil but ConnectorTypesService.List was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		ListArgs *ListArguments
	}{
		Ctx:      ctx,
		ListArgs: listArgs,
	}
	mock.lockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	mock.lockList.Unlock()
	return mock.ListFunc(ctx, listArgs)
}

// ListCalls gets all the calls that were made to List.
// Check the length with:
//     len(mockedConnectorTypesService.ListCalls())
func (mock *ConnectorTypesServiceMock) ListCalls() []struct {
	Ctx      context.Context
	ListArgs *ListArguments
} {
	var calls []struct {
		Ctx      context.Context
		ListArgs *ListArguments
	}
	mock.lockList.RLock()
	calls = mock.calls.List
	mock.lockList.RUnlock()
	return calls
}
