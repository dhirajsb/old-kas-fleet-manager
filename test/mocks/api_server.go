package mocks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"

	"k8s.io/apimachinery/pkg/util/wait"

	"time"

	"github.com/gorilla/mux"
	clustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	ocmErrors "gitlab.cee.redhat.com/service/managed-services-api/pkg/errors"
)

const (
	// EndpointPathClusters ocm clusters management service clusters endpoint
	EndpointPathClusters = "/api/clusters_mgmt/v1/clusters"
	// EndpointPathClusters ocm clusters management service clusters endpoint
	EndpointPathCluster = "/api/clusters_mgmt/v1/clusters/{id}"
	// EndpointPathSyncsets ocm clusters management service syncset endpoint
	EndpointPathSyncsets = "/api/clusters_mgmt/v1/clusters/{id}/external_configuration/syncsets"
	// EndpointPathSyncsetsDelete ocm clusters management service syncset endpoint delete
	EndpointPathSyncsetsDelete = "/api/clusters_mgmt/v1/clusters/{id}/external_configuration/syncsets/{syncsetID}"
	// EndpointPathIngresses ocm cluster management ingress endpoint
	EndpointPathIngresses = "/api/clusters_mgmt/v1/clusters/{id}/ingresses"
	// EndpointPathCloudProviders ocm cluster management cloud providers endpoint
	EndpointPathCloudProviders = "/api/clusters_mgmt/v1/cloud_providers"
	// EndpointPathCloudProvider ocm cluster management cloud provider endpoint
	EndpointPathCloudProvider = "/api/clusters_mgmt/v1/cloud_providers/{id}"
	// EndpointPathCloudProviderRegions ocm cluster management cloud provider regions endpoint
	EndpointPathCloudProviderRegions = "/api/clusters_mgmt/v1/cloud_providers/{id}/regions"
	// EndpointPathCloudProviderRegion ocm cluster management cloud provider region endpoint
	EndpointPathCloudProviderRegion = "/api/clusters_mgmt/v1/cloud_providers/{providerID}/regions/{regionID}"
	// EndpointPathClusterStatus ocm cluster management cluster status endpoint
	EndpointPathClusterStatus = "/api/clusters_mgmt/v1/clusters/{id}/status"
	// EndpointPathClusterAddons ocm cluster management cluster addons endpoint
	EndpointPathClusterAddons = "/api/clusters_mgmt/v1/clusters/{id}/addons"
	// EndpointPathMachinePools ocm cluster management machine pools endpoint
	EndpointPathMachinePools = "/api/clusters_mgmt/v1/clusters/{id}/machine_pools"
	// EndpointPathMachinePool ocm cluster management machine pool endpoint
	EndpointPathMachinePool = "/api/clusters_mgmt/v1/clusters/{id}/machine_pools/{machinePoolId}"
)

// variables for endpoints
var (
	EndpointClusterGet              = Endpoint{EndpointPathCluster, http.MethodGet}
	EndpointKafkaDelete             = Endpoint{EndpointPathSyncsetsDelete, http.MethodDelete}
	EndpointClustersGet             = Endpoint{EndpointPathClusters, http.MethodGet}
	EndpointClustersPost            = Endpoint{EndpointPathClusters, http.MethodPost}
	EndpointClusterSyncsetPost      = Endpoint{EndpointPathSyncsets, http.MethodPost}
	EndpointClusterIngressGet       = Endpoint{EndpointPathIngresses, http.MethodGet}
	EndpointCloudProvidersGet       = Endpoint{EndpointPathCloudProviders, http.MethodGet}
	EndpointCloudProviderGet        = Endpoint{EndpointPathCloudProvider, http.MethodGet}
	EndpointCloudProviderRegionsGet = Endpoint{EndpointPathCloudProviderRegions, http.MethodGet}
	EndpointCloudProviderRegionGet  = Endpoint{EndpointPathCloudProviderRegion, http.MethodGet}
	EndpointClusterStatusGet        = Endpoint{EndpointPathClusterStatus, http.MethodGet}
	EndpointClusterAddonsGet        = Endpoint{EndpointPathClusterAddons, http.MethodGet}
	EndpointClusterAddonPost        = Endpoint{EndpointPathClusterAddons, http.MethodPost}
	EndpointMachinePoolsGet         = Endpoint{EndpointPathMachinePools, http.MethodGet}
	EndpointMachinePoolPost         = Endpoint{EndpointPathMachinePools, http.MethodPost}
	EndpointMachinePoolPatch        = Endpoint{EndpointPathMachinePool, http.MethodPatch}
	EndpointMachinePoolGet          = Endpoint{EndpointPathMachinePool, http.MethodGet}
)

// variables for mocked ocm types
//
// these are the default types that will be returned by the emulated ocm api
// to override these values, do not set them directly e.g. mocks.MockSyncset = ...
// instead use the Set*Response functions provided by MockConfigurableServerBuilder e.g. SetClusterGetResponse(...)
var (
	MockSyncset                      *clustersmgmtv1.Syncset
	MockIngressList                  *clustersmgmtv1.IngressList
	MockCloudProvider                *clustersmgmtv1.CloudProvider
	MockCloudProviderList            *clustersmgmtv1.CloudProviderList
	MockCloudProviderRegion          *clustersmgmtv1.CloudRegion
	MockCloudProviderRegionList      *clustersmgmtv1.CloudRegionList
	MockClusterStatus                *clustersmgmtv1.ClusterStatus
	MockClusterAddonInstallation     *clustersmgmtv1.AddOnInstallation
	MockClusterAddonInstallationList *clustersmgmtv1.AddOnInstallationList
	MockMachinePoolList              *clustersmgmtv1.MachinePoolList
	MockMachinePool                  *clustersmgmtv1.MachinePool
	MockCluster                      *clustersmgmtv1.Cluster
)

// routerSwapper is an http.Handler that allows you to swap mux routers.
type routerSwapper struct {
	mu     sync.Mutex
	router *mux.Router
}

// Swap changes the old router with the new one.
func (rs *routerSwapper) Swap(newRouter *mux.Router) {
	rs.mu.Lock()
	rs.router = newRouter
	rs.mu.Unlock()
}

var router *mux.Router

// rSwapper is required if any change to the Router for mocked OCM server is needed
var rSwapper *routerSwapper

// Endpoint is a wrapper around an endpoint and the method used to interact with that endpoint e.g. GET /clusters
type Endpoint struct {
	Path   string
	Method string
}

// HandlerRegister is a cache that maps Endpoints to their handlers
type HandlerRegister map[Endpoint]func(w http.ResponseWriter, r *http.Request)

// MockConfigurableServerBuilder allows mock ocm api servers to be built
type MockConfigurableServerBuilder struct {
	// handlerRegister cache of endpoints and handlers to be used when the mock ocm api server is built
	handlerRegister HandlerRegister
}

// NewMockConfigurableServerBuilder returns a new builder that can be used to define a mock ocm api server
func NewMockConfigurableServerBuilder() *MockConfigurableServerBuilder {
	// get the default endpoint handlers that'll be used if they're not overridden
	handlerRegister, err := getDefaultHandlerRegister()
	if err != nil {
		panic(err)
	}
	return &MockConfigurableServerBuilder{
		handlerRegister: handlerRegister,
	}
}

// SetClusterGetResponse set a mock response cluster or error for the POST /api/clusters_mgmt/v1/clusters endpoint
func (b *MockConfigurableServerBuilder) SetClusterGetResponse(cluster *clustersmgmtv1.Cluster, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointClusterGet] = buildMockRequestHandler(cluster, err)
}

// SetKafkaDeleteResponse set a mock response cluster or error for the DELETE /api/clusters_mgmt/v1/clusters/{id}/external_configuration/syncsets/{syncsetID} endpoint
func (b *MockConfigurableServerBuilder) SetKafkaDeleteResponse(syncset *clustersmgmtv1.Syncset, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointKafkaDelete] = buildMockRequestHandler(syncset, err)
}

// SetClustersPostResponse set a mock response cluster or error for the POST /api/clusters_mgmt/v1/clusters endpoint
func (b *MockConfigurableServerBuilder) SetClustersPostResponse(cluster *clustersmgmtv1.Cluster, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointClustersPost] = buildMockRequestHandler(cluster, err)
}

// SetClustersGetResponse set a mock response cluster or error for the GET /api/clusters_mgmt/v1/clusters endpoint
func (b *MockConfigurableServerBuilder) SetClustersGetResponse(cluster *clustersmgmtv1.Cluster, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointClustersGet] = buildMockRequestHandler(cluster, err)
}

// SetClusterSyncsetPostResponse set a mock response syncset or error for the POST /api/clusters_mgmt/v1/clusters/{id}/syncsets endpoint
func (b *MockConfigurableServerBuilder) SetClusterSyncsetPostResponse(syncset *clustersmgmtv1.Syncset, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointClusterSyncsetPost] = buildMockRequestHandler(syncset, err)
}

// SetClusterIngressGetResponse set a mock response ingress or error for the GET /api/clusters_mgmt/v1/clusters/{id}/ingresses endpoint
func (b *MockConfigurableServerBuilder) SetClusterIngressGetResponse(ingress *clustersmgmtv1.Ingress, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointClusterIngressGet] = buildMockRequestHandler(ingress, err)
}

// SetCloudProvidersGetResponse set a mock response provider list or error for GET /api/clusters_mgmt/v1/cloud_providers
func (b *MockConfigurableServerBuilder) SetCloudProvidersGetResponse(providers *clustersmgmtv1.CloudProviderList, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointCloudProvidersGet] = buildMockRequestHandler(providers, err)
}

// SetCloudRegionsGetResponse set a mock response region list or error for GET /api/clusters_mgmt/v1/cloud_providers/{id}/regions
func (b *MockConfigurableServerBuilder) SetCloudRegionsGetResponse(regions *clustersmgmtv1.CloudRegionList, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointCloudProviderRegionsGet] = buildMockRequestHandler(regions, err)
}

// SetCloudRegionGetResponse set a mock response region or error for GET /api/clusters_mgmt/v1/cloud_providers/{id}/regions/{regionId}
func (b *MockConfigurableServerBuilder) SetCloudRegionGetResponse(region *clustersmgmtv1.CloudRegion, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointCloudProviderRegionGet] = buildMockRequestHandler(region, err)
}

// SetClusterStatusGetResponse set a mock response cluster status or error for GET /api/clusters_mgmt/v1/clusters/{id}/status
func (b *MockConfigurableServerBuilder) SetClusterStatusGetResponse(status *clustersmgmtv1.ClusterStatus, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointClusterStatusGet] = buildMockRequestHandler(status, err)
}

// SetClusterAddonsGetResponse set a mock response addon list or error for GET /api/clusters_mgmt/v1/clusters/{id}/addons
func (b *MockConfigurableServerBuilder) SetClusterAddonsGetResponse(addons *clustersmgmtv1.AddOnInstallationList, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointClusterAddonsGet] = buildMockRequestHandler(addons, err)
}

// SetClusterAddonPostResponse set a mock response addon or error for POST /api/clusters_mgmt/v1/clusters/{id}/addons
func (b *MockConfigurableServerBuilder) SetClusterAddonPostResponse(addon *clustersmgmtv1.AddOnInstallation, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointClusterAddonPost] = buildMockRequestHandler(addon, err)
}

// SetMachinePoolsGetResponse set a mock response machine pool or error for Get /api/clusters_mgmt/v1/clusters/{id}/machine_pools
func (b *MockConfigurableServerBuilder) SetMachinePoolsGetResponse(mp *clustersmgmtv1.MachinePoolList, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointMachinePoolsGet] = buildMockRequestHandler(mp, err)
}

// SetMachinePoolGetResponse set a mock response machine pool list or error for Get /api/clusters_mgmt/v1/clusters/{id}/machine_pools/{machinePoolId}
func (b *MockConfigurableServerBuilder) SetMachinePoolGetResponse(mp *clustersmgmtv1.MachinePoolList, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointMachinePoolGet] = buildMockRequestHandler(mp, err)
}

// SetMachinePoolPostResponse set a mock response for Post /api/clusters_mgmt/v1/clusters/{id}/machine_pools
func (b *MockConfigurableServerBuilder) SetMachinePoolPostResponse(mp *clustersmgmtv1.MachinePool, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointMachinePoolPost] = buildMockRequestHandler(mp, err)
}

// SetMachinePoolPatchResponse set a mock response for Patch /api/clusters_mgmt/v1/clusters/{id}/machine_pools/{machinePoolId}
func (b *MockConfigurableServerBuilder) SetMachinePoolPatchResponse(mp *clustersmgmtv1.MachinePool, err *ocmErrors.ServiceError) {
	b.handlerRegister[EndpointMachinePoolPatch] = buildMockRequestHandler(mp, err)
}

// Build builds the mock ocm api server using the endpoint handlers that have been set in the builder
func (b *MockConfigurableServerBuilder) Build() *httptest.Server {
	router = mux.NewRouter()
	rSwapper = &routerSwapper{sync.Mutex{}, router}

	// set up handlers from the builder
	for endpoint, handleFn := range b.handlerRegister {
		router.HandleFunc(endpoint.Path, handleFn).Methods(endpoint.Method)
	}
	server := httptest.NewUnstartedServer(rSwapper)
	l, err := net.Listen("tcp", "127.0.0.1:9876")
	if err != nil {
		log.Fatal(err)
	}
	server.Listener = l
	server.Start()
	err = wait.PollImmediate(time.Second, 10*time.Second, func() (done bool, err error) {
		_, err = http.Get("http://127.0.0.1:9876/api/clusters_mgmt/v1/cloud_providers/aws/regions")
		return err == nil, nil
	})
	if err != nil {
		log.Fatal("Timed out waiting for mock server to start.")
		panic(err)
	}
	return server
}

// SwapRouterResponse and update the router to handle this response
func (b *MockConfigurableServerBuilder) SwapRouterResponse(path string, method string, successType interface{}, serviceErr *ocmErrors.ServiceError) {
	b.handlerRegister[Endpoint{
		Path:   path,
		Method: method,
	}] = buildMockRequestHandler(successType, serviceErr)

	router = mux.NewRouter()
	for endpoint, handleFn := range b.handlerRegister {
		router.HandleFunc(endpoint.Path, handleFn).Methods(endpoint.Method)
	}

	rSwapper.Swap(router)
}

// ServeHTTP makes the routerSwapper to implement the http.Handler interface
// so that routerSwapper can be used by httptest.NewServer()
func (rs *routerSwapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rs.mu.Lock()
	router := rs.router
	rs.mu.Unlock()
	router.ServeHTTP(w, r)
}

// getDefaultHandlerRegister returns a set of default endpoints and handlers used in the mock ocm api server
func getDefaultHandlerRegister() (HandlerRegister, error) {
	// define a list of default endpoints and handlers in the mock ocm api server, when new endpoints are used in the
	// managed-services-api service, a default ocm response should also be added here
	return HandlerRegister{
		EndpointClusterGet:              buildMockRequestHandler(MockCluster, nil),
		EndpointKafkaDelete:             buildMockRequestHandler(MockSyncset, nil),
		EndpointClustersGet:             buildMockRequestHandler(MockCluster, nil),
		EndpointClustersPost:            buildMockRequestHandler(MockCluster, nil),
		EndpointClusterSyncsetPost:      buildMockRequestHandler(MockSyncset, nil),
		EndpointClusterIngressGet:       buildMockRequestHandler(MockIngressList, nil),
		EndpointCloudProvidersGet:       buildMockRequestHandler(MockCloudProviderList, nil),
		EndpointCloudProviderGet:        buildMockRequestHandler(MockCloudProvider, nil),
		EndpointCloudProviderRegionsGet: buildMockRequestHandler(MockCloudProviderRegionList, nil),
		EndpointCloudProviderRegionGet:  buildMockRequestHandler(MockCloudProviderRegion, nil),
		EndpointClusterStatusGet:        buildMockRequestHandler(MockClusterStatus, nil),
		EndpointClusterAddonsGet:        buildMockRequestHandler(MockClusterAddonInstallationList, nil),
		EndpointClusterAddonPost:        buildMockRequestHandler(MockClusterAddonInstallation, nil),
		EndpointMachinePoolsGet:         buildMockRequestHandler(MockMachinePoolList, nil),
		EndpointMachinePoolGet:          buildMockRequestHandler(MockMachinePool, nil),
		EndpointMachinePoolPatch:        buildMockRequestHandler(MockMachinePool, nil),
		EndpointMachinePoolPost:         buildMockRequestHandler(MockMachinePool, nil),
	}, nil
}

// buildMockRequestHandler builds a generic handler for all ocm api server responses
// one of successType of serviceErr should be defined
// if serviceErr is defined, it will be provided as an ocm error response
// if successType is defined, it will be provided as an ocm success response
func buildMockRequestHandler(successType interface{}, serviceErr *ocmErrors.ServiceError) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if serviceErr != nil {
			w.WriteHeader(serviceErr.HttpCode)
			if err := marshalOCMType(serviceErr, w); err != nil {
				panic(err)
			}
		} else if successType != nil {
			if err := marshalOCMType(successType, w); err != nil {
				panic(err)
			}
		} else {
			panic("no response was defined")
		}
	}
}

// marshalOCMType marshals known ocm types to a provided io.Writer using the ocm sdk marshallers
func marshalOCMType(t interface{}, w io.Writer) error {
	switch t.(type) {
	// handle cluster types
	case *clustersmgmtv1.Cluster:
		return clustersmgmtv1.MarshalCluster(t.(*clustersmgmtv1.Cluster), w)
	// handle cluster status types
	case *clustersmgmtv1.ClusterStatus:
		return clustersmgmtv1.MarshalClusterStatus(t.(*clustersmgmtv1.ClusterStatus), w)
	// handle syncset types
	case *clustersmgmtv1.Syncset:
		return clustersmgmtv1.MarshalSyncset(t.(*clustersmgmtv1.Syncset), w)
	// handle ingress types
	case *clustersmgmtv1.Ingress:
		return clustersmgmtv1.MarshalIngress(t.(*clustersmgmtv1.Ingress), w)
	case []*clustersmgmtv1.Ingress:
		return clustersmgmtv1.MarshalIngressList(t.([]*clustersmgmtv1.Ingress), w)
	// for any <type>List ocm type we'll need to follow this pattern to ensure the array of objects
	// is wrapped with an OCMList object
	case *clustersmgmtv1.IngressList:
		ocmList, err := NewOCMList().WithItems(t.(*clustersmgmtv1.IngressList).Slice())
		if err != nil {
			return err
		}
		return json.NewEncoder(w).Encode(ocmList)
	// handle cloud provider types
	case *clustersmgmtv1.CloudProvider:
		return clustersmgmtv1.MarshalCloudProvider(t.(*clustersmgmtv1.CloudProvider), w)
	case []*clustersmgmtv1.CloudProvider:
		return clustersmgmtv1.MarshalCloudProviderList(t.([]*clustersmgmtv1.CloudProvider), w)
	case *clustersmgmtv1.CloudProviderList:
		ocmList, err := NewOCMList().WithItems(t.(*clustersmgmtv1.CloudProviderList).Slice())
		if err != nil {
			return err
		}
		return json.NewEncoder(w).Encode(ocmList)
	// handle cloud region types
	case *clustersmgmtv1.CloudRegion:
		return clustersmgmtv1.MarshalCloudRegion(t.(*clustersmgmtv1.CloudRegion), w)
	case []*clustersmgmtv1.CloudRegion:
		return clustersmgmtv1.MarshalCloudRegionList(t.([]*clustersmgmtv1.CloudRegion), w)
	case *clustersmgmtv1.CloudRegionList:
		ocmList, err := NewOCMList().WithItems(t.(*clustersmgmtv1.CloudRegionList).Slice())
		if err != nil {
			return err
		}
		return json.NewEncoder(w).Encode(ocmList)
	// handle cluster addon installations
	case *clustersmgmtv1.AddOnInstallation:
		return clustersmgmtv1.MarshalAddOnInstallation(t.(*clustersmgmtv1.AddOnInstallation), w)
	case []*clustersmgmtv1.AddOnInstallation:
		return clustersmgmtv1.MarshalAddOnInstallationList(t.([]*clustersmgmtv1.AddOnInstallation), w)
	case *clustersmgmtv1.AddOnInstallationList:
		ocmList, err := NewOCMList().WithItems(t.(*clustersmgmtv1.AddOnInstallationList).Slice())
		if err != nil {
			return err
		}
		return json.NewEncoder(w).Encode(ocmList)
	case *clustersmgmtv1.MachinePool:
		return clustersmgmtv1.MarshalMachinePool(t.(*clustersmgmtv1.MachinePool), w)
	case []*clustersmgmtv1.MachinePool:
		return clustersmgmtv1.MarshalMachinePoolList(t.([]*clustersmgmtv1.MachinePool), w)
	case *clustersmgmtv1.MachinePoolList:
		ocmList, err := NewOCMList().WithItems(t.(*clustersmgmtv1.MachinePoolList).Slice())
		if err != nil {
			return err
		}
		return json.NewEncoder(w).Encode(ocmList)
	// handle the generic ocm list type
	case *ocmList:
		return json.NewEncoder(w).Encode(t)
	// handle ocm error type
	case *ocmErrors.ServiceError:
		return json.NewEncoder(w).Encode(t.(*ocmErrors.ServiceError).AsOpenapiError(""))
	}
	return errors.New(fmt.Sprintf("could not recognise type %s in ocm type marshaller", reflect.TypeOf(t).String()))
}

// basic wrapper to emulate the the ocm list types as they're private
type ocmList struct {
	HREF  *string         `json:"href"`
	Link  bool            `json:"link"`
	Items json.RawMessage `json:"items"`
}

func NewOCMList() *ocmList {
	return &ocmList{
		HREF:  nil,
		Link:  false,
		Items: nil,
	}
}

func (l *ocmList) WithHREF(href string) *ocmList {
	l.HREF = &href
	return l
}

func (l *ocmList) WithLink(link bool) *ocmList {
	l.Link = link
	return l
}

func (l *ocmList) WithItems(items interface{}) (*ocmList, error) {
	var b bytes.Buffer
	if err := marshalOCMType(items, &b); err != nil {
		return l, err
	}
	l.Items = b.Bytes()
	return l, nil
}

// init the shared mock types, panic if we fail, this should never fail
func init() {
	// mockClusterID default mock cluster id used in the mock ocm server
	mockClusterID := "2aad9fc1-c40e-471f-8d57-fdaecc7d3686"
	// mockCloudProviderID default mock provider ID
	mockCloudProviderID := "aws"
	// mockClusterExternalID default mock cluster external ID
	mockClusterExternalID := "2aad9fc1-c40e-471f-8d57-fdaecc7d3686"
	// mockClusterState default mock cluster state
	mockClusterState := clustersmgmtv1.ClusterStateReady
	// mockCloudProviderDisplayName default mock provider display name
	mockCloudProviderDisplayName := "AWS"
	// mockCloudRegionID default mock cluster region
	mockCloudRegionID := "eu-west-1"
	// mockCloudRegionDisplayName default mock cloud region display name
	mockCloudRegionDisplayName := "EU, Ireland"
	// mockSyncsetID default mock syncset id used in the mock ocm server
	mockSyncsetID := "ext-8a41f783-b5e4-4692-a7cd-c0b9c8eeede9"
	// mockIngressID default mock ingress id used in the mock ocm server
	mockIngressID := "s1h5"
	// mockIngressDNS default mock ingress dns used in the mock ocm server
	mockIngressDNS := "apps.ms-btq2d1h8d3b1.b3k3.s1.devshift.org"
	// mockIngressHref default mock ingress HREF used in the mock ocm server
	mockIngressHref := "/api/clusters_mgmt/v1/clusters/000/ingresses/i8y1"
	// mockIngressListening default mock ingress listening used in the mock ocm server
	mockIngressListening := clustersmgmtv1.ListeningMethodExternal
	// mockClusterAddonID default mock cluster addon ID
	mockClusterAddonID := "managed-kafka"
	// mockClusterAddonState default mock cluster addon state
	mockClusterAddonState := clustersmgmtv1.AddOnInstallationStateReady
	// mockClusterAddonDescription default mock cluster addon description
	mockClusterAddonDescription := "InstallWaiting"
	// mockMachinePoolID default machine pool ID
	mockMachinePoolID := "managed"
	// mockMachinePoolReplicas default number of machine pool replicas
	mockMachinePoolReplicas := 2

	var err error

	// mock syncsets
	mockMockSyncsetBuilder := GetMockSyncsetBuilder(mockSyncsetID, mockClusterID)

	MockSyncset, err = GetMockSyncset(mockMockSyncsetBuilder)
	if err != nil {
		panic(err)
	}

	// mock ingresses
	MockIngressList, err = GetMockIngressList(mockIngressID, mockIngressDNS, true, mockIngressListening, mockIngressHref)
	if err != nil {
		panic(err)
	}

	// mock cloud providers
	mockCloudProviderBuilder := GetMockCloudProviderBuilder(mockCloudProviderID, mockCloudProviderDisplayName)

	MockCloudProvider, err = GetMockCloudProvider(mockCloudProviderBuilder)
	if err != nil {
		panic(err)
	}

	MockCloudProviderList, err = GetMockCloudProviderList(mockCloudProviderBuilder)
	if err != nil {
		panic(err)
	}

	// mock cloud provider regions/cloud regions
	mockCloudProviderRegionBuilder := GetMockCloudProviderRegionBuilder(mockCloudRegionID, mockCloudProviderID, mockCloudRegionDisplayName, mockCloudProviderBuilder, true, true)

	MockCloudProviderRegion, err = GetMockCloudProviderRegion(mockCloudProviderRegionBuilder)
	if err != nil {
		panic(err)
	}

	MockCloudProviderRegionList, err = GetMockCloudProviderRegionList(mockCloudProviderRegionBuilder)
	if err != nil {
		panic(err)
	}

	// mock cluster status
	MockClusterStatus, err = GetMockClusterStatus(mockClusterID, mockClusterState)

	// mock cluster addons
	mockClusterAddonBuilder := GetMockClusterAddonBuilder(mockClusterAddonID)

	// mock cluster addon installations
	mockClusterAddonInstallationBuilder := GetMockClusterAddonInstallationBuilder(mockClusterAddonID, mockClusterID, mockClusterAddonBuilder, mockClusterAddonState, mockClusterAddonDescription)

	MockClusterAddonInstallation, err = GetMockClusterAddonInstallation(mockClusterAddonInstallationBuilder)
	if err != nil {
		panic(err)
	}

	MockClusterAddonInstallationList, err = GetMockClusterAddonInstallationList(mockClusterAddonInstallationBuilder)
	if err != nil {
		panic(err)
	}

	// mock cluster
	mockClusterBuilder := GetMockClusterBuilder(mockClusterID, mockClusterExternalID, mockClusterState, mockCloudProviderBuilder, mockCloudProviderRegionBuilder)

	MockCluster, err = GetMockCluster(mockClusterBuilder)

	if err != nil {
		panic(err)
	}

	// Mock machine pools
	mockMachinePoolBuilder := GetMockMachineBuilder(mockMachinePoolID, mockClusterID, mockMachinePoolReplicas, mockClusterBuilder)
	MockMachinePoolList, err = GetMachinePoolList(mockMachinePoolBuilder)
	if err != nil {
		panic(err)
	}

	MockMachinePool, err = GetMockMachinePool(mockMachinePoolBuilder)
	if err != nil {
		panic(err)
	}
}

// GetMockSyncsetBuilder for emulated OCM server
func GetMockSyncsetBuilder(syncsetID string, clusterID string) *clustersmgmtv1.SyncsetBuilder {
	return clustersmgmtv1.NewSyncset().
		ID(syncsetID).
		HREF(fmt.Sprintf("/api/clusters_mgmt/v1/clusters/%s/external_configuration/syncsets/%s", clusterID, syncsetID))
}

// GetMockSyncset for emulated OCM server
func GetMockSyncset(syncsetBuilder *clustersmgmtv1.SyncsetBuilder) (*clustersmgmtv1.Syncset, error) {
	return syncsetBuilder.Build()
}

// GetMockIngressList for emulated OCM server
func GetMockIngressList(ingressID string, ingressDNS string, isDefault bool, listening v1.ListeningMethod, href string) (*clustersmgmtv1.IngressList, error) {
	return clustersmgmtv1.NewIngressList().Items(
		clustersmgmtv1.NewIngress().ID(ingressID).DNSName(ingressDNS).Default(isDefault).Listening(listening).HREF(href)).Build()
}

// GetMockCloudProviderBuilder for emulated OCM server
func GetMockCloudProviderBuilder(cloudProviderID string, cloudProviderDisplayName string) *clustersmgmtv1.CloudProviderBuilder {
	return clustersmgmtv1.NewCloudProvider().
		ID(cloudProviderID).
		Name(cloudProviderID).
		DisplayName(cloudProviderDisplayName).
		HREF(fmt.Sprintf("/api/clusters_mgmt/v1/cloud_providers/%s", cloudProviderID))
}

// GetMockCloudProvider for emulated OCM server
func GetMockCloudProvider(cloudProviderBuilder *clustersmgmtv1.CloudProviderBuilder) (*clustersmgmtv1.CloudProvider, error) {
	return cloudProviderBuilder.Build()
}

// GetMockCloudProviderList for emulated OCM server
func GetMockCloudProviderList(cloudProviderBuilder *clustersmgmtv1.CloudProviderBuilder) (*clustersmgmtv1.CloudProviderList, error) {
	return clustersmgmtv1.NewCloudProviderList().
		Items(cloudProviderBuilder).
		Build()
}

// GetMockCloudProviderRegionBuilder for emulated OCM server
func GetMockCloudProviderRegionBuilder(cloudRegionID string, cloudProviderID string, cloudRegionDisplayName string, cloudProviderBuilder *clustersmgmtv1.CloudProviderBuilder, enabled bool, multiAZ bool) *clustersmgmtv1.CloudRegionBuilder {
	return clustersmgmtv1.NewCloudRegion().
		ID(cloudRegionID).
		HREF(fmt.Sprintf("/api/clusters_mgmt/v1/cloud_providers/%s/regions/%s", cloudProviderID, cloudRegionID)).
		DisplayName(cloudRegionDisplayName).
		CloudProvider(cloudProviderBuilder).
		Enabled(enabled).
		SupportsMultiAZ(multiAZ)
}

// GetMockCloudProviderRegion for emulated OCM server
func GetMockCloudProviderRegion(cloudProviderRegionBuilder *clustersmgmtv1.CloudRegionBuilder) (*clustersmgmtv1.CloudRegion, error) {
	return cloudProviderRegionBuilder.Build()
}

// GetMockCloudProviderRegionList for emulated OCM server
func GetMockCloudProviderRegionList(cloudProviderRegionBuilder *clustersmgmtv1.CloudRegionBuilder) (*clustersmgmtv1.CloudRegionList, error) {
	return clustersmgmtv1.NewCloudRegionList().Items(cloudProviderRegionBuilder).Build()
}

// GetMockClusterStatus for emulated OCM server
func GetMockClusterStatus(clusterID string, clusterState clustersmgmtv1.ClusterState) (*clustersmgmtv1.ClusterStatus, error) {
	return clustersmgmtv1.NewClusterStatus().
		ID(clusterID).
		HREF(fmt.Sprintf("/api/clusters_mgmt/v1/clusters/%s/status", clusterID)).
		State(clusterState).
		Description("").
		Build()
}

// GetMockClusterAddonBuilder for emulated OCM server
func GetMockClusterAddonBuilder(clusterAddonID string) *clustersmgmtv1.AddOnBuilder {
	return clustersmgmtv1.NewAddOn().
		ID(clusterAddonID).
		HREF(fmt.Sprintf("/api/clusters_mgmt/v1/addons/%s", clusterAddonID))
}

// GetMockClusterAddonInstallationBuilder for emulated OCM server
func GetMockClusterAddonInstallationBuilder(clusterAddonID string, clusterID string, clusterAddonBuilder *clustersmgmtv1.AddOnBuilder, addonInstallatonState clustersmgmtv1.AddOnInstallationState, clusterAddonDescription string) *clustersmgmtv1.AddOnInstallationBuilder {
	return clustersmgmtv1.NewAddOnInstallation().
		ID(clusterAddonID).
		HREF(fmt.Sprintf("/api/clusters_mgmt/v1/clusters/%s/addons/managed-kafka", clusterID)).
		Addon(clusterAddonBuilder).
		State(addonInstallatonState).
		StateDescription(clusterAddonDescription)
}

// GetMockClusterAddonInstallation for emulated OCM server
func GetMockClusterAddonInstallation(clusterAddonInstallationBuilder *clustersmgmtv1.AddOnInstallationBuilder) (*clustersmgmtv1.AddOnInstallation, error) {
	return clusterAddonInstallationBuilder.Build()
}

// GetMockClusterAddonInstallationList for emulated OCM server
func GetMockClusterAddonInstallationList(clusterAddonInstallationBuilder *clustersmgmtv1.AddOnInstallationBuilder) (*clustersmgmtv1.AddOnInstallationList, error) {
	return clustersmgmtv1.NewAddOnInstallationList().Items(clusterAddonInstallationBuilder).Build()
}

// GetMockClusterBuilder for emulated OCM server
func GetMockClusterBuilder(clusterID string, clusterExternalID string, clusterState clustersmgmtv1.ClusterState, clusterProviderBuilder *clustersmgmtv1.CloudProviderBuilder, cloudProviderRegionBuilder *clustersmgmtv1.CloudRegionBuilder) *clustersmgmtv1.ClusterBuilder {
	return clustersmgmtv1.NewCluster().
		ID(clusterID).
		ExternalID(clusterExternalID).
		State(clusterState).
		CloudProvider(clusterProviderBuilder).
		Region(cloudProviderRegionBuilder)
}

// GetMockCluster for emulated OCM server
func GetMockCluster(clusterBuilder *clustersmgmtv1.ClusterBuilder) (*clustersmgmtv1.Cluster, error) {
	return clusterBuilder.Build()
}

// GetMockMachineBuilder for emulated OCM server
func GetMockMachineBuilder(machinePoolID string, clusterID string, replicas int, clusterBuilder *clustersmgmtv1.ClusterBuilder) *clustersmgmtv1.MachinePoolBuilder {
	return clustersmgmtv1.NewMachinePool().
		ID(machinePoolID).
		HREF(fmt.Sprintf("/api/clusters_mgmt/v1/clusters/%s/machine_pools/%s", clusterID, machinePoolID)).
		Replicas(replicas).
		Cluster(clusterBuilder)
}

// GetMachinePoolList for emulated OCM server
func GetMachinePoolList(machinePoolBuilder *clustersmgmtv1.MachinePoolBuilder) (*clustersmgmtv1.MachinePoolList, error) {
	return clustersmgmtv1.NewMachinePoolList().Items(machinePoolBuilder).Build()
}

// GetMockMachinePool for emulated OCM server
func GetMockMachinePool(machinePoolBuilder *clustersmgmtv1.MachinePoolBuilder) (*clustersmgmtv1.MachinePool, error) {
	return machinePoolBuilder.Build()
}
