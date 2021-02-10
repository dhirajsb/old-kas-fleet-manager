// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package services

import (
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/api"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/errors"
	"sync"
)

// Ensure, that KasFleetshardOperatorAddonMock does implement KasFleetshardOperatorAddon.
// If this is not the case, regenerate this file with moq.
var _ KasFleetshardOperatorAddon = &KasFleetshardOperatorAddonMock{}

// KasFleetshardOperatorAddonMock is a mock implementation of KasFleetshardOperatorAddon.
//
//     func TestSomethingThatUsesKasFleetshardOperatorAddon(t *testing.T) {
//
//         // make and configure a mocked KasFleetshardOperatorAddon
//         mockedKasFleetshardOperatorAddon := &KasFleetshardOperatorAddonMock{
//             ProvisionFunc: func(cluster api.Cluster) (bool, *errors.ServiceError) {
// 	               panic("mock out the Provision method")
//             },
//         }
//
//         // use mockedKasFleetshardOperatorAddon in code that requires KasFleetshardOperatorAddon
//         // and then make assertions.
//
//     }
type KasFleetshardOperatorAddonMock struct {
	// ProvisionFunc mocks the Provision method.
	ProvisionFunc func(cluster api.Cluster) (bool, *errors.ServiceError)

	// calls tracks calls to the methods.
	calls struct {
		// Provision holds details about calls to the Provision method.
		Provision []struct {
			// Cluster is the cluster argument value.
			Cluster api.Cluster
		}
	}
	lockProvision sync.RWMutex
}

// Provision calls ProvisionFunc.
func (mock *KasFleetshardOperatorAddonMock) Provision(cluster api.Cluster) (bool, *errors.ServiceError) {
	if mock.ProvisionFunc == nil {
		panic("KasFleetshardOperatorAddonMock.ProvisionFunc: method is nil but KasFleetshardOperatorAddon.Provision was just called")
	}
	callInfo := struct {
		Cluster api.Cluster
	}{
		Cluster: cluster,
	}
	mock.lockProvision.Lock()
	mock.calls.Provision = append(mock.calls.Provision, callInfo)
	mock.lockProvision.Unlock()
	return mock.ProvisionFunc(cluster)
}

// ProvisionCalls gets all the calls that were made to Provision.
// Check the length with:
//     len(mockedKasFleetshardOperatorAddon.ProvisionCalls())
func (mock *KasFleetshardOperatorAddonMock) ProvisionCalls() []struct {
	Cluster api.Cluster
} {
	var calls []struct {
		Cluster api.Cluster
	}
	mock.lockProvision.RLock()
	calls = mock.calls.Provision
	mock.lockProvision.RUnlock()
	return calls
}
