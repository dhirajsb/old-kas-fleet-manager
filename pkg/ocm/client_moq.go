// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package ocm

import (
	"github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"sync"
)

// Ensure, that ClientMock does implement Client.
// If this is not the case, regenerate this file with moq.
var _ Client = &ClientMock{}

// ClientMock is a mock implementation of Client.
//
//     func TestSomethingThatUsesClient(t *testing.T) {
//
//         // make and configure a mocked Client
//         mockedClient := &ClientMock{
//             CreateAddonFunc: func(clusterId string, addonId string) (*v1.AddOnInstallation, error) {
// 	               panic("mock out the CreateAddon method")
//             },
//             CreateAddonWithParamsFunc: func(clusterId string, addonId string, parameters []AddonParameter) (*v1.AddOnInstallation, error) {
// 	               panic("mock out the CreateAddonWithParams method")
//             },
//             CreateClusterFunc: func(cluster *v1.Cluster) (*v1.Cluster, error) {
// 	               panic("mock out the CreateCluster method")
//             },
//             CreateSyncSetFunc: func(clusterID string, syncset *v1.Syncset) (*v1.Syncset, error) {
// 	               panic("mock out the CreateSyncSet method")
//             },
//             DeleteSyncSetFunc: func(clusterID string, syncsetID string) (int, error) {
// 	               panic("mock out the DeleteSyncSet method")
//             },
//             GetAddonFunc: func(clusterId string, addonId string) (*v1.AddOnInstallation, error) {
// 	               panic("mock out the GetAddon method")
//             },
//             GetCloudProvidersFunc: func() (*v1.CloudProviderList, error) {
// 	               panic("mock out the GetCloudProviders method")
//             },
//             GetClusterDNSFunc: func(clusterID string) (string, error) {
// 	               panic("mock out the GetClusterDNS method")
//             },
//             GetClusterIngressesFunc: func(clusterID string) (*v1.IngressesListResponse, error) {
// 	               panic("mock out the GetClusterIngresses method")
//             },
//             GetClusterStatusFunc: func(id string) (*v1.ClusterStatus, error) {
// 	               panic("mock out the GetClusterStatus method")
//             },
//             GetRegionsFunc: func(provider *v1.CloudProvider) (*v1.CloudRegionList, error) {
// 	               panic("mock out the GetRegions method")
//             },
//             GetSyncSetFunc: func(clusterID string, syncSetID string) (*v1.Syncset, error) {
// 	               panic("mock out the GetSyncSet method")
//             },
//             ScaleDownComputeNodesFunc: func(clusterID string, decrement int) (*v1.Cluster, error) {
// 	               panic("mock out the ScaleDownComputeNodes method")
//             },
//             ScaleUpComputeNodesFunc: func(clusterID string, increment int) (*v1.Cluster, error) {
// 	               panic("mock out the ScaleUpComputeNodes method")
//             },
//         }
//
//         // use mockedClient in code that requires Client
//         // and then make assertions.
//
//     }
type ClientMock struct {
	// CreateAddonFunc mocks the CreateAddon method.
	CreateAddonFunc func(clusterId string, addonId string) (*v1.AddOnInstallation, error)

	// CreateAddonWithParamsFunc mocks the CreateAddonWithParams method.
	CreateAddonWithParamsFunc func(clusterId string, addonId string, parameters []AddonParameter) (*v1.AddOnInstallation, error)

	// CreateClusterFunc mocks the CreateCluster method.
	CreateClusterFunc func(cluster *v1.Cluster) (*v1.Cluster, error)

	// CreateSyncSetFunc mocks the CreateSyncSet method.
	CreateSyncSetFunc func(clusterID string, syncset *v1.Syncset) (*v1.Syncset, error)

	// DeleteSyncSetFunc mocks the DeleteSyncSet method.
	DeleteSyncSetFunc func(clusterID string, syncsetID string) (int, error)

	// GetAddonFunc mocks the GetAddon method.
	GetAddonFunc func(clusterId string, addonId string) (*v1.AddOnInstallation, error)

	// GetCloudProvidersFunc mocks the GetCloudProviders method.
	GetCloudProvidersFunc func() (*v1.CloudProviderList, error)

	// GetClusterDNSFunc mocks the GetClusterDNS method.
	GetClusterDNSFunc func(clusterID string) (string, error)

	// GetClusterIngressesFunc mocks the GetClusterIngresses method.
	GetClusterIngressesFunc func(clusterID string) (*v1.IngressesListResponse, error)

	// GetClusterStatusFunc mocks the GetClusterStatus method.
	GetClusterStatusFunc func(id string) (*v1.ClusterStatus, error)

	// GetRegionsFunc mocks the GetRegions method.
	GetRegionsFunc func(provider *v1.CloudProvider) (*v1.CloudRegionList, error)

	// GetSyncSetFunc mocks the GetSyncSet method.
	GetSyncSetFunc func(clusterID string, syncSetID string) (*v1.Syncset, error)

	// ScaleDownComputeNodesFunc mocks the ScaleDownComputeNodes method.
	ScaleDownComputeNodesFunc func(clusterID string, decrement int) (*v1.Cluster, error)

	// ScaleUpComputeNodesFunc mocks the ScaleUpComputeNodes method.
	ScaleUpComputeNodesFunc func(clusterID string, increment int) (*v1.Cluster, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateAddon holds details about calls to the CreateAddon method.
		CreateAddon []struct {
			// ClusterId is the clusterId argument value.
			ClusterId string
			// AddonId is the addonId argument value.
			AddonId string
		}
		// CreateAddonWithParams holds details about calls to the CreateAddonWithParams method.
		CreateAddonWithParams []struct {
			// ClusterId is the clusterId argument value.
			ClusterId string
			// AddonId is the addonId argument value.
			AddonId string
			// Parameters is the parameters argument value.
			Parameters []AddonParameter
		}
		// CreateCluster holds details about calls to the CreateCluster method.
		CreateCluster []struct {
			// Cluster is the cluster argument value.
			Cluster *v1.Cluster
		}
		// CreateSyncSet holds details about calls to the CreateSyncSet method.
		CreateSyncSet []struct {
			// ClusterID is the clusterID argument value.
			ClusterID string
			// Syncset is the syncset argument value.
			Syncset *v1.Syncset
		}
		// DeleteSyncSet holds details about calls to the DeleteSyncSet method.
		DeleteSyncSet []struct {
			// ClusterID is the clusterID argument value.
			ClusterID string
			// SyncsetID is the syncsetID argument value.
			SyncsetID string
		}
		// GetAddon holds details about calls to the GetAddon method.
		GetAddon []struct {
			// ClusterId is the clusterId argument value.
			ClusterId string
			// AddonId is the addonId argument value.
			AddonId string
		}
		// GetCloudProviders holds details about calls to the GetCloudProviders method.
		GetCloudProviders []struct {
		}
		// GetClusterDNS holds details about calls to the GetClusterDNS method.
		GetClusterDNS []struct {
			// ClusterID is the clusterID argument value.
			ClusterID string
		}
		// GetClusterIngresses holds details about calls to the GetClusterIngresses method.
		GetClusterIngresses []struct {
			// ClusterID is the clusterID argument value.
			ClusterID string
		}
		// GetClusterStatus holds details about calls to the GetClusterStatus method.
		GetClusterStatus []struct {
			// ID is the id argument value.
			ID string
		}
		// GetRegions holds details about calls to the GetRegions method.
		GetRegions []struct {
			// Provider is the provider argument value.
			Provider *v1.CloudProvider
		}
		// GetSyncSet holds details about calls to the GetSyncSet method.
		GetSyncSet []struct {
			// ClusterID is the clusterID argument value.
			ClusterID string
			// SyncSetID is the syncSetID argument value.
			SyncSetID string
		}
		// ScaleDownComputeNodes holds details about calls to the ScaleDownComputeNodes method.
		ScaleDownComputeNodes []struct {
			// ClusterID is the clusterID argument value.
			ClusterID string
			// Decrement is the decrement argument value.
			Decrement int
		}
		// ScaleUpComputeNodes holds details about calls to the ScaleUpComputeNodes method.
		ScaleUpComputeNodes []struct {
			// ClusterID is the clusterID argument value.
			ClusterID string
			// Increment is the increment argument value.
			Increment int
		}
	}
	lockCreateAddon           sync.RWMutex
	lockCreateAddonWithParams sync.RWMutex
	lockCreateCluster         sync.RWMutex
	lockCreateSyncSet         sync.RWMutex
	lockDeleteSyncSet         sync.RWMutex
	lockGetAddon              sync.RWMutex
	lockGetCloudProviders     sync.RWMutex
	lockGetClusterDNS         sync.RWMutex
	lockGetClusterIngresses   sync.RWMutex
	lockGetClusterStatus      sync.RWMutex
	lockGetRegions            sync.RWMutex
	lockGetSyncSet            sync.RWMutex
	lockScaleDownComputeNodes sync.RWMutex
	lockScaleUpComputeNodes   sync.RWMutex
}

// CreateAddon calls CreateAddonFunc.
func (mock *ClientMock) CreateAddon(clusterId string, addonId string) (*v1.AddOnInstallation, error) {
	if mock.CreateAddonFunc == nil {
		panic("ClientMock.CreateAddonFunc: method is nil but Client.CreateAddon was just called")
	}
	callInfo := struct {
		ClusterId string
		AddonId   string
	}{
		ClusterId: clusterId,
		AddonId:   addonId,
	}
	mock.lockCreateAddon.Lock()
	mock.calls.CreateAddon = append(mock.calls.CreateAddon, callInfo)
	mock.lockCreateAddon.Unlock()
	return mock.CreateAddonFunc(clusterId, addonId)
}

// CreateAddonCalls gets all the calls that were made to CreateAddon.
// Check the length with:
//     len(mockedClient.CreateAddonCalls())
func (mock *ClientMock) CreateAddonCalls() []struct {
	ClusterId string
	AddonId   string
} {
	var calls []struct {
		ClusterId string
		AddonId   string
	}
	mock.lockCreateAddon.RLock()
	calls = mock.calls.CreateAddon
	mock.lockCreateAddon.RUnlock()
	return calls
}

// CreateAddonWithParams calls CreateAddonWithParamsFunc.
func (mock *ClientMock) CreateAddonWithParams(clusterId string, addonId string, parameters []AddonParameter) (*v1.AddOnInstallation, error) {
	if mock.CreateAddonWithParamsFunc == nil {
		panic("ClientMock.CreateAddonWithParamsFunc: method is nil but Client.CreateAddonWithParams was just called")
	}
	callInfo := struct {
		ClusterId  string
		AddonId    string
		Parameters []AddonParameter
	}{
		ClusterId:  clusterId,
		AddonId:    addonId,
		Parameters: parameters,
	}
	mock.lockCreateAddonWithParams.Lock()
	mock.calls.CreateAddonWithParams = append(mock.calls.CreateAddonWithParams, callInfo)
	mock.lockCreateAddonWithParams.Unlock()
	return mock.CreateAddonWithParamsFunc(clusterId, addonId, parameters)
}

// CreateAddonWithParamsCalls gets all the calls that were made to CreateAddonWithParams.
// Check the length with:
//     len(mockedClient.CreateAddonWithParamsCalls())
func (mock *ClientMock) CreateAddonWithParamsCalls() []struct {
	ClusterId  string
	AddonId    string
	Parameters []AddonParameter
} {
	var calls []struct {
		ClusterId  string
		AddonId    string
		Parameters []AddonParameter
	}
	mock.lockCreateAddonWithParams.RLock()
	calls = mock.calls.CreateAddonWithParams
	mock.lockCreateAddonWithParams.RUnlock()
	return calls
}

// CreateCluster calls CreateClusterFunc.
func (mock *ClientMock) CreateCluster(cluster *v1.Cluster) (*v1.Cluster, error) {
	if mock.CreateClusterFunc == nil {
		panic("ClientMock.CreateClusterFunc: method is nil but Client.CreateCluster was just called")
	}
	callInfo := struct {
		Cluster *v1.Cluster
	}{
		Cluster: cluster,
	}
	mock.lockCreateCluster.Lock()
	mock.calls.CreateCluster = append(mock.calls.CreateCluster, callInfo)
	mock.lockCreateCluster.Unlock()
	return mock.CreateClusterFunc(cluster)
}

// CreateClusterCalls gets all the calls that were made to CreateCluster.
// Check the length with:
//     len(mockedClient.CreateClusterCalls())
func (mock *ClientMock) CreateClusterCalls() []struct {
	Cluster *v1.Cluster
} {
	var calls []struct {
		Cluster *v1.Cluster
	}
	mock.lockCreateCluster.RLock()
	calls = mock.calls.CreateCluster
	mock.lockCreateCluster.RUnlock()
	return calls
}

// CreateSyncSet calls CreateSyncSetFunc.
func (mock *ClientMock) CreateSyncSet(clusterID string, syncset *v1.Syncset) (*v1.Syncset, error) {
	if mock.CreateSyncSetFunc == nil {
		panic("ClientMock.CreateSyncSetFunc: method is nil but Client.CreateSyncSet was just called")
	}
	callInfo := struct {
		ClusterID string
		Syncset   *v1.Syncset
	}{
		ClusterID: clusterID,
		Syncset:   syncset,
	}
	mock.lockCreateSyncSet.Lock()
	mock.calls.CreateSyncSet = append(mock.calls.CreateSyncSet, callInfo)
	mock.lockCreateSyncSet.Unlock()
	return mock.CreateSyncSetFunc(clusterID, syncset)
}

// CreateSyncSetCalls gets all the calls that were made to CreateSyncSet.
// Check the length with:
//     len(mockedClient.CreateSyncSetCalls())
func (mock *ClientMock) CreateSyncSetCalls() []struct {
	ClusterID string
	Syncset   *v1.Syncset
} {
	var calls []struct {
		ClusterID string
		Syncset   *v1.Syncset
	}
	mock.lockCreateSyncSet.RLock()
	calls = mock.calls.CreateSyncSet
	mock.lockCreateSyncSet.RUnlock()
	return calls
}

// DeleteSyncSet calls DeleteSyncSetFunc.
func (mock *ClientMock) DeleteSyncSet(clusterID string, syncsetID string) (int, error) {
	if mock.DeleteSyncSetFunc == nil {
		panic("ClientMock.DeleteSyncSetFunc: method is nil but Client.DeleteSyncSet was just called")
	}
	callInfo := struct {
		ClusterID string
		SyncsetID string
	}{
		ClusterID: clusterID,
		SyncsetID: syncsetID,
	}
	mock.lockDeleteSyncSet.Lock()
	mock.calls.DeleteSyncSet = append(mock.calls.DeleteSyncSet, callInfo)
	mock.lockDeleteSyncSet.Unlock()
	return mock.DeleteSyncSetFunc(clusterID, syncsetID)
}

// DeleteSyncSetCalls gets all the calls that were made to DeleteSyncSet.
// Check the length with:
//     len(mockedClient.DeleteSyncSetCalls())
func (mock *ClientMock) DeleteSyncSetCalls() []struct {
	ClusterID string
	SyncsetID string
} {
	var calls []struct {
		ClusterID string
		SyncsetID string
	}
	mock.lockDeleteSyncSet.RLock()
	calls = mock.calls.DeleteSyncSet
	mock.lockDeleteSyncSet.RUnlock()
	return calls
}

// GetAddon calls GetAddonFunc.
func (mock *ClientMock) GetAddon(clusterId string, addonId string) (*v1.AddOnInstallation, error) {
	if mock.GetAddonFunc == nil {
		panic("ClientMock.GetAddonFunc: method is nil but Client.GetAddon was just called")
	}
	callInfo := struct {
		ClusterId string
		AddonId   string
	}{
		ClusterId: clusterId,
		AddonId:   addonId,
	}
	mock.lockGetAddon.Lock()
	mock.calls.GetAddon = append(mock.calls.GetAddon, callInfo)
	mock.lockGetAddon.Unlock()
	return mock.GetAddonFunc(clusterId, addonId)
}

// GetAddonCalls gets all the calls that were made to GetAddon.
// Check the length with:
//     len(mockedClient.GetAddonCalls())
func (mock *ClientMock) GetAddonCalls() []struct {
	ClusterId string
	AddonId   string
} {
	var calls []struct {
		ClusterId string
		AddonId   string
	}
	mock.lockGetAddon.RLock()
	calls = mock.calls.GetAddon
	mock.lockGetAddon.RUnlock()
	return calls
}

// GetCloudProviders calls GetCloudProvidersFunc.
func (mock *ClientMock) GetCloudProviders() (*v1.CloudProviderList, error) {
	if mock.GetCloudProvidersFunc == nil {
		panic("ClientMock.GetCloudProvidersFunc: method is nil but Client.GetCloudProviders was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetCloudProviders.Lock()
	mock.calls.GetCloudProviders = append(mock.calls.GetCloudProviders, callInfo)
	mock.lockGetCloudProviders.Unlock()
	return mock.GetCloudProvidersFunc()
}

// GetCloudProvidersCalls gets all the calls that were made to GetCloudProviders.
// Check the length with:
//     len(mockedClient.GetCloudProvidersCalls())
func (mock *ClientMock) GetCloudProvidersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetCloudProviders.RLock()
	calls = mock.calls.GetCloudProviders
	mock.lockGetCloudProviders.RUnlock()
	return calls
}

// GetClusterDNS calls GetClusterDNSFunc.
func (mock *ClientMock) GetClusterDNS(clusterID string) (string, error) {
	if mock.GetClusterDNSFunc == nil {
		panic("ClientMock.GetClusterDNSFunc: method is nil but Client.GetClusterDNS was just called")
	}
	callInfo := struct {
		ClusterID string
	}{
		ClusterID: clusterID,
	}
	mock.lockGetClusterDNS.Lock()
	mock.calls.GetClusterDNS = append(mock.calls.GetClusterDNS, callInfo)
	mock.lockGetClusterDNS.Unlock()
	return mock.GetClusterDNSFunc(clusterID)
}

// GetClusterDNSCalls gets all the calls that were made to GetClusterDNS.
// Check the length with:
//     len(mockedClient.GetClusterDNSCalls())
func (mock *ClientMock) GetClusterDNSCalls() []struct {
	ClusterID string
} {
	var calls []struct {
		ClusterID string
	}
	mock.lockGetClusterDNS.RLock()
	calls = mock.calls.GetClusterDNS
	mock.lockGetClusterDNS.RUnlock()
	return calls
}

// GetClusterIngresses calls GetClusterIngressesFunc.
func (mock *ClientMock) GetClusterIngresses(clusterID string) (*v1.IngressesListResponse, error) {
	if mock.GetClusterIngressesFunc == nil {
		panic("ClientMock.GetClusterIngressesFunc: method is nil but Client.GetClusterIngresses was just called")
	}
	callInfo := struct {
		ClusterID string
	}{
		ClusterID: clusterID,
	}
	mock.lockGetClusterIngresses.Lock()
	mock.calls.GetClusterIngresses = append(mock.calls.GetClusterIngresses, callInfo)
	mock.lockGetClusterIngresses.Unlock()
	return mock.GetClusterIngressesFunc(clusterID)
}

// GetClusterIngressesCalls gets all the calls that were made to GetClusterIngresses.
// Check the length with:
//     len(mockedClient.GetClusterIngressesCalls())
func (mock *ClientMock) GetClusterIngressesCalls() []struct {
	ClusterID string
} {
	var calls []struct {
		ClusterID string
	}
	mock.lockGetClusterIngresses.RLock()
	calls = mock.calls.GetClusterIngresses
	mock.lockGetClusterIngresses.RUnlock()
	return calls
}

// GetClusterStatus calls GetClusterStatusFunc.
func (mock *ClientMock) GetClusterStatus(id string) (*v1.ClusterStatus, error) {
	if mock.GetClusterStatusFunc == nil {
		panic("ClientMock.GetClusterStatusFunc: method is nil but Client.GetClusterStatus was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGetClusterStatus.Lock()
	mock.calls.GetClusterStatus = append(mock.calls.GetClusterStatus, callInfo)
	mock.lockGetClusterStatus.Unlock()
	return mock.GetClusterStatusFunc(id)
}

// GetClusterStatusCalls gets all the calls that were made to GetClusterStatus.
// Check the length with:
//     len(mockedClient.GetClusterStatusCalls())
func (mock *ClientMock) GetClusterStatusCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetClusterStatus.RLock()
	calls = mock.calls.GetClusterStatus
	mock.lockGetClusterStatus.RUnlock()
	return calls
}

// GetRegions calls GetRegionsFunc.
func (mock *ClientMock) GetRegions(provider *v1.CloudProvider) (*v1.CloudRegionList, error) {
	if mock.GetRegionsFunc == nil {
		panic("ClientMock.GetRegionsFunc: method is nil but Client.GetRegions was just called")
	}
	callInfo := struct {
		Provider *v1.CloudProvider
	}{
		Provider: provider,
	}
	mock.lockGetRegions.Lock()
	mock.calls.GetRegions = append(mock.calls.GetRegions, callInfo)
	mock.lockGetRegions.Unlock()
	return mock.GetRegionsFunc(provider)
}

// GetRegionsCalls gets all the calls that were made to GetRegions.
// Check the length with:
//     len(mockedClient.GetRegionsCalls())
func (mock *ClientMock) GetRegionsCalls() []struct {
	Provider *v1.CloudProvider
} {
	var calls []struct {
		Provider *v1.CloudProvider
	}
	mock.lockGetRegions.RLock()
	calls = mock.calls.GetRegions
	mock.lockGetRegions.RUnlock()
	return calls
}

// GetSyncSet calls GetSyncSetFunc.
func (mock *ClientMock) GetSyncSet(clusterID string, syncSetID string) (*v1.Syncset, error) {
	if mock.GetSyncSetFunc == nil {
		panic("ClientMock.GetSyncSetFunc: method is nil but Client.GetSyncSet was just called")
	}
	callInfo := struct {
		ClusterID string
		SyncSetID string
	}{
		ClusterID: clusterID,
		SyncSetID: syncSetID,
	}
	mock.lockGetSyncSet.Lock()
	mock.calls.GetSyncSet = append(mock.calls.GetSyncSet, callInfo)
	mock.lockGetSyncSet.Unlock()
	return mock.GetSyncSetFunc(clusterID, syncSetID)
}

// GetSyncSetCalls gets all the calls that were made to GetSyncSet.
// Check the length with:
//     len(mockedClient.GetSyncSetCalls())
func (mock *ClientMock) GetSyncSetCalls() []struct {
	ClusterID string
	SyncSetID string
} {
	var calls []struct {
		ClusterID string
		SyncSetID string
	}
	mock.lockGetSyncSet.RLock()
	calls = mock.calls.GetSyncSet
	mock.lockGetSyncSet.RUnlock()
	return calls
}

// ScaleDownComputeNodes calls ScaleDownComputeNodesFunc.
func (mock *ClientMock) ScaleDownComputeNodes(clusterID string, decrement int) (*v1.Cluster, error) {
	if mock.ScaleDownComputeNodesFunc == nil {
		panic("ClientMock.ScaleDownComputeNodesFunc: method is nil but Client.ScaleDownComputeNodes was just called")
	}
	callInfo := struct {
		ClusterID string
		Decrement int
	}{
		ClusterID: clusterID,
		Decrement: decrement,
	}
	mock.lockScaleDownComputeNodes.Lock()
	mock.calls.ScaleDownComputeNodes = append(mock.calls.ScaleDownComputeNodes, callInfo)
	mock.lockScaleDownComputeNodes.Unlock()
	return mock.ScaleDownComputeNodesFunc(clusterID, decrement)
}

// ScaleDownComputeNodesCalls gets all the calls that were made to ScaleDownComputeNodes.
// Check the length with:
//     len(mockedClient.ScaleDownComputeNodesCalls())
func (mock *ClientMock) ScaleDownComputeNodesCalls() []struct {
	ClusterID string
	Decrement int
} {
	var calls []struct {
		ClusterID string
		Decrement int
	}
	mock.lockScaleDownComputeNodes.RLock()
	calls = mock.calls.ScaleDownComputeNodes
	mock.lockScaleDownComputeNodes.RUnlock()
	return calls
}

// ScaleUpComputeNodes calls ScaleUpComputeNodesFunc.
func (mock *ClientMock) ScaleUpComputeNodes(clusterID string, increment int) (*v1.Cluster, error) {
	if mock.ScaleUpComputeNodesFunc == nil {
		panic("ClientMock.ScaleUpComputeNodesFunc: method is nil but Client.ScaleUpComputeNodes was just called")
	}
	callInfo := struct {
		ClusterID string
		Increment int
	}{
		ClusterID: clusterID,
		Increment: increment,
	}
	mock.lockScaleUpComputeNodes.Lock()
	mock.calls.ScaleUpComputeNodes = append(mock.calls.ScaleUpComputeNodes, callInfo)
	mock.lockScaleUpComputeNodes.Unlock()
	return mock.ScaleUpComputeNodesFunc(clusterID, increment)
}

// ScaleUpComputeNodesCalls gets all the calls that were made to ScaleUpComputeNodes.
// Check the length with:
//     len(mockedClient.ScaleUpComputeNodesCalls())
func (mock *ClientMock) ScaleUpComputeNodesCalls() []struct {
	ClusterID string
	Increment int
} {
	var calls []struct {
		ClusterID string
		Increment int
	}
	mock.lockScaleUpComputeNodes.RLock()
	calls = mock.calls.ScaleUpComputeNodes
	mock.lockScaleUpComputeNodes.RUnlock()
	return calls
}
