package workers

import (
	"errors"
	"reflect"
	"testing"
	"time"

	ingressoperatorv1 "github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api/ingressoperator/v1"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/services/syncsetresources"
	storagev1 "k8s.io/api/storage/v1"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/config"

	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/operator-framework/api/pkg/operators/v1alpha1"
	"github.com/operator-framework/api/pkg/operators/v1alpha2"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/ocm"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api"
	ocmErrors "github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/services"
	. "github.com/onsi/gomega"
	clustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"

	projectv1 "github.com/openshift/api/project/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	k8sCoreV1 "k8s.io/api/core/v1"

	apiErrors "github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
)

var (
	testRegion   = "us-west-1"
	testProvider = "aws"
)

func TestClusterManager_reconcileClusterStatus(t *testing.T) {
	type fields struct {
		ocmClient      ocm.Client
		clusterService services.ClusterService
		timer          *time.Timer
	}
	type args struct {
		cluster *api.Cluster
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.Cluster
		wantErr bool
	}{
		{
			name: "error when getting cluster status from ocm fails",
			fields: fields{
				ocmClient: &ocm.ClientMock{
					GetClusterStatusFunc: func(id string) (status *clustersmgmtv1.ClusterStatus, e error) {
						return nil, errors.New("test")
					},
				},
			},
			args: args{
				cluster: &api.Cluster{
					ClusterID: "test",
				},
			},
			wantErr: true,
		},
		{
			name: "error when updating status in database fails",
			fields: fields{
				ocmClient: &ocm.ClientMock{
					GetClusterStatusFunc: func(id string) (status *clustersmgmtv1.ClusterStatus, e error) {
						clusterStatus, err := clustersmgmtv1.NewClusterStatus().State(clustersmgmtv1.ClusterStateReady).Build()
						if err != nil {
							panic(err)
						}
						return clusterStatus, nil
					},
				},
				clusterService: &services.ClusterServiceMock{
					UpdateStatusFunc: func(cluster api.Cluster, status api.ClusterStatus) error {
						return errors.New("test")
					},
				},
			},
			args: args{
				cluster: &api.Cluster{
					ClusterID: "test",
					Status:    api.ClusterProvisioning,
				},
			},
			wantErr: true,
		},
		{
			name: "database update not invoked when update not needed",
			fields: fields{
				ocmClient: &ocm.ClientMock{
					GetClusterStatusFunc: func(id string) (status *clustersmgmtv1.ClusterStatus, e error) {
						clusterStatus, err := clustersmgmtv1.NewClusterStatus().State(clustersmgmtv1.ClusterStateInstalling).Build()
						if err != nil {
							panic(err)
						}
						return clusterStatus, nil
					},
				},
				clusterService: &services.ClusterServiceMock{
					UpdateStatusFunc: func(cluster api.Cluster, status api.ClusterStatus) error {
						// this should never be invoked as the cluster state is already accurate
						return errors.New("test")
					},
				},
			},
			args: args{
				cluster: &api.Cluster{
					ClusterID: "test",
					Status:    api.ClusterProvisioning,
				},
			},
			want: &api.Cluster{
				ClusterID: "test",
				Status:    api.ClusterProvisioning,
			},
		},
		{
			name: "pending state is set when internal cluster status is empty",
			fields: fields{
				ocmClient: &ocm.ClientMock{
					GetClusterStatusFunc: func(id string) (status *clustersmgmtv1.ClusterStatus, e error) {
						clusterStatus, err := clustersmgmtv1.NewClusterStatus().State(clustersmgmtv1.ClusterStatePending).Build()
						if err != nil {
							panic(err)
						}
						return clusterStatus, nil
					},
				},
				clusterService: &services.ClusterServiceMock{
					UpdateStatusFunc: func(cluster api.Cluster, status api.ClusterStatus) error {
						return nil
					},
				},
			},
			args: args{
				cluster: &api.Cluster{
					ClusterID: "test",
					Status:    "",
				},
			},
			want: &api.Cluster{
				ClusterID: "test",
				Status:    api.ClusterProvisioning,
			},
		},
		{
			name: "state is failed when underlying ocm cluster failed",
			fields: fields{
				ocmClient: &ocm.ClientMock{
					GetClusterStatusFunc: func(id string) (status *clustersmgmtv1.ClusterStatus, e error) {
						clusterStatus, err := clustersmgmtv1.NewClusterStatus().State(clustersmgmtv1.ClusterStateError).Build()
						if err != nil {
							panic(err)
						}
						return clusterStatus, nil
					},
				},
				clusterService: &services.ClusterServiceMock{
					UpdateStatusFunc: func(cluster api.Cluster, status api.ClusterStatus) error {
						return nil
					},
				},
			},
			args: args{
				cluster: &api.Cluster{
					ClusterID: "test",
					Status:    api.ClusterProvisioning,
				},
			},
			want: &api.Cluster{
				ClusterID: "test",
				Status:    api.ClusterFailed,
			},
		},
		{
			name: "successful reconcile",
			fields: fields{
				ocmClient: &ocm.ClientMock{
					GetClusterStatusFunc: func(id string) (status *clustersmgmtv1.ClusterStatus, e error) {
						clusterStatus, err := clustersmgmtv1.NewClusterStatus().State(clustersmgmtv1.ClusterStateReady).Build()
						if err != nil {
							panic(err)
						}
						return clusterStatus, nil
					},
				},
				clusterService: &services.ClusterServiceMock{
					UpdateStatusFunc: func(cluster api.Cluster, status api.ClusterStatus) error {
						return nil
					},
				},
			},
			args: args{
				cluster: &api.Cluster{
					ClusterID: "test",
					Status:    api.ClusterProvisioning,
				},
			},
			want: &api.Cluster{
				ClusterID: "test",
				Status:    api.ClusterProvisioned,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ClusterManager{
				ocmClient:      tt.fields.ocmClient,
				clusterService: tt.fields.clusterService,
				timer:          tt.fields.timer,
			}
			got, err := c.reconcileClusterStatus(tt.args.cluster)
			if (err != nil) != tt.wantErr {
				t.Errorf("reconcileClusterStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconcileClusterStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClusterManager_reconcileStrimziOperator(t *testing.T) {
	type fields struct {
		ocmClient      ocm.Client
		timer          *time.Timer
		clusterService services.ClusterService
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "error when getting managed kafka addon from ocm fails",
			fields: fields{
				ocmClient: &ocm.ClientMock{
					GetAddonFunc: func(clusterId string, addonId string) (status *clustersmgmtv1.AddOnInstallation, e error) {
						return nil, errors.New("error when getting managed kafka addon from ocm")
					},
				},
			},
			wantErr: true,
		},
		{
			name: "empty state returned when managed kafka addon not found",
			fields: fields{
				ocmClient: &ocm.ClientMock{
					GetAddonFunc: func(clusterId string, addonId string) (status *clustersmgmtv1.AddOnInstallation, e error) {
						managedKafkaAddon := &clustersmgmtv1.AddOnInstallation{}
						return managedKafkaAddon, nil
					},
					CreateAddonFunc: func(clusterId string, addonId string) (status *clustersmgmtv1.AddOnInstallation, e error) {
						managedKafkaAddon := &clustersmgmtv1.AddOnInstallation{}
						return managedKafkaAddon, nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "empty state returned when managed kafka addon is found but with no state",
			fields: fields{
				ocmClient: &ocm.ClientMock{
					GetAddonFunc: func(clusterId string, addonId string) (status *clustersmgmtv1.AddOnInstallation, e error) {
						managedKafkaAddon, err := clustersmgmtv1.NewAddOnInstallation().ID(api.ManagedKafkaAddonID).Build()
						if err != nil {
							panic(err)
						}
						return managedKafkaAddon, nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "failed state returned when managed kafka addon is found but with a AddOnInstallationStateFailed state",
			fields: fields{
				ocmClient: &ocm.ClientMock{
					GetAddonFunc: func(clusterId string, addonId string) (status *clustersmgmtv1.AddOnInstallation, e error) {
						managedKafkaAddon, err := clustersmgmtv1.NewAddOnInstallation().ID(api.ManagedKafkaAddonID).State(clustersmgmtv1.AddOnInstallationStateFailed).Build()
						if err != nil {
							panic(err)
						}
						return managedKafkaAddon, nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ready state returned when managed kafka addon is found but with a AddOnInstallationStateReady state",
			fields: fields{
				ocmClient: &ocm.ClientMock{
					GetAddonFunc: func(clusterId string, addonId string) (status *clustersmgmtv1.AddOnInstallation, e error) {
						managedKafkaAddon, err := clustersmgmtv1.NewAddOnInstallation().ID(api.ManagedKafkaAddonID).State(clustersmgmtv1.AddOnInstallationStateReady).Build()
						if err != nil {
							panic(err)
						}
						return managedKafkaAddon, nil
					},
					CreateSyncSetFunc: func(clusterID string, syncset *clustersmgmtv1.Syncset) (*clustersmgmtv1.Syncset, error) {
						return &clustersmgmtv1.Syncset{}, nil
					},
				},
				clusterService: &services.ClusterServiceMock{
					GetClusterDNSFunc: func(clusterID string) (string, *apiErrors.ServiceError) {
						return "apps.example.com", nil
					},
					UpdateStatusFunc: func(cluster api.Cluster, status api.ClusterStatus) error {
						return nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error when creating managed kafka addon from ocm fails",
			fields: fields{
				ocmClient: &ocm.ClientMock{
					GetAddonFunc: func(clusterId string, addonId string) (status *clustersmgmtv1.AddOnInstallation, e error) {
						managedKafkaAddon := &clustersmgmtv1.AddOnInstallation{}
						return managedKafkaAddon, nil
					},
					CreateAddonFunc: func(clusterId string, addonId string) (status *clustersmgmtv1.AddOnInstallation, e error) {
						return nil, errors.New("error when creating managed kafka addon from ocm")
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ClusterManager{
				ocmClient:      tt.fields.ocmClient,
				clusterService: tt.fields.clusterService,
				timer:          tt.fields.timer,
				configService: services.NewConfigService(
					config.ApplicationConfig{
						SupportedProviders:         &config.ProviderConfig{},
						AllowList:                  &config.AllowListConfig{},
						ObservabilityConfiguration: &config.ObservabilityConfiguration{},
						ClusterCreationConfig:      &config.ClusterCreationConfig{},
					},
				),
			}
			err := c.reconcileStrimziOperator(api.Cluster{
				ClusterID: "clusterId",
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("reconcileStrimziOperator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClusterManager_reconcileAcceptedCluster(t *testing.T) {
	type fields struct {
		providerLst           []string
		clusterService        services.ClusterService
		providersConfig       config.ProviderConfig
		clusterCreationConfig config.ClusterCreationConfig
	}

	tests := []struct {
		name    string
		wantErr bool
		fields  fields
	}{
		{
			name: "reconcile cluster with cluster creation requests",
			fields: fields{
				providerLst: []string{"us-east-1"},
				clusterService: &services.ClusterServiceMock{
					ListGroupByProviderAndRegionFunc: func(providers []string, regions []string, status []string) (m []*services.ResGroupCPRegion, e *ocmErrors.ServiceError) {
						var res []*services.ResGroupCPRegion
						return res, nil
					},
					CreateFunc: func(Cluster *api.Cluster) (cls *v1.Cluster, e *ocmErrors.ServiceError) {
						sample, _ := v1.NewCluster().Build()
						return sample, nil
					},
				},
				clusterCreationConfig: config.ClusterCreationConfig{
					AutoOSDCreation: true,
				},
				providersConfig: config.ProviderConfig{
					ProvidersConfig: config.ProviderConfiguration{
						SupportedProviders: config.ProviderList{
							config.Provider{
								Name: "aws",
								Regions: config.RegionList{
									config.Region{
										Name: "us-east-1",
									},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ClusterManager{
				clusterService: tt.fields.clusterService,
				configService: services.NewConfigService(
					config.ApplicationConfig{
						SupportedProviders:         &tt.fields.providersConfig,
						AllowList:                  &config.AllowListConfig{},
						ObservabilityConfiguration: &config.ObservabilityConfiguration{},
						ClusterCreationConfig:      &tt.fields.clusterCreationConfig,
					}),
			}

			clusterRequest := &api.Cluster{
				Region:        testRegion,
				CloudProvider: testProvider,
				Status:        "cluster_accepted",
			}

			err := c.reconcileAcceptedCluster(clusterRequest)
			if err != nil && !tt.wantErr {
				t.Errorf("reconcileAcceptedCluster() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClusterManager_reconcileClustersForRegions(t *testing.T) {
	type fields struct {
		providerLst           []string
		clusterService        services.ClusterService
		providersConfig       config.ProviderConfig
		clusterCreationConfig config.ClusterCreationConfig
	}

	tests := []struct {
		name    string
		wantErr bool
		fields  fields
	}{
		{
			name: "creates a missing OSD cluster request automatically",
			fields: fields{
				providerLst: []string{"us-east-1"},
				clusterService: &services.ClusterServiceMock{
					ListGroupByProviderAndRegionFunc: func(providers []string, regions []string, status []string) (m []*services.ResGroupCPRegion, e *ocmErrors.ServiceError) {
						var res []*services.ResGroupCPRegion
						return res, nil
					},
					RegisterClusterJobFunc: func(clusterReq *api.Cluster) *apiErrors.ServiceError {
						return nil
					},
				},
				clusterCreationConfig: config.ClusterCreationConfig{
					AutoOSDCreation: true,
				},
				providersConfig: config.ProviderConfig{
					ProvidersConfig: config.ProviderConfiguration{
						SupportedProviders: config.ProviderList{
							config.Provider{
								Name: "aws",
								Regions: config.RegionList{
									config.Region{
										Name: "us-east-1",
									},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "failed to create OSD request",
			fields: fields{
				providerLst: []string{"us-east-1"},
				clusterService: &services.ClusterServiceMock{
					ListGroupByProviderAndRegionFunc: func(providers []string, regions []string, status []string) (m []*services.ResGroupCPRegion, e *ocmErrors.ServiceError) {
						var res []*services.ResGroupCPRegion
						return res, nil
					},
					RegisterClusterJobFunc: func(clusterReq *api.Cluster) *apiErrors.ServiceError {
						return apiErrors.GeneralError("failed to create cluster request")
					},
				},
				clusterCreationConfig: config.ClusterCreationConfig{
					AutoOSDCreation: true,
				},
				providersConfig: config.ProviderConfig{
					ProvidersConfig: config.ProviderConfiguration{
						SupportedProviders: config.ProviderList{
							config.Provider{
								Name: "aws",
								Regions: config.RegionList{
									config.Region{
										Name: "us-east-1",
									},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "failed to retrieve OSD cluster info from database",
			fields: fields{
				providerLst: []string{"us-east-1"},
				clusterService: &services.ClusterServiceMock{
					ListGroupByProviderAndRegionFunc: func(providers []string, regions []string, status []string) (m []*services.ResGroupCPRegion, e *ocmErrors.ServiceError) {
						return nil, ocmErrors.New(ocmErrors.ErrorGeneral, "Database retrieval failed")
					},
				},
				clusterCreationConfig: config.ClusterCreationConfig{
					AutoOSDCreation: true,
				},
				providersConfig: config.ProviderConfig{
					ProvidersConfig: config.ProviderConfiguration{
						SupportedProviders: config.ProviderList{
							config.Provider{
								Name: "aws",
								Regions: config.RegionList{
									config.Region{
										Name: "us-east-1",
									},
								},
							},
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ClusterManager{
				clusterService: tt.fields.clusterService,
				configService: services.NewConfigService(config.ApplicationConfig{
					SupportedProviders:         &tt.fields.providersConfig,
					AllowList:                  &config.AllowListConfig{},
					ClusterCreationConfig:      &config.ClusterCreationConfig{},
					ObservabilityConfiguration: &config.ObservabilityConfiguration{},
				}),
			}
			err := c.reconcileClustersForRegions()
			if err != nil && !tt.wantErr {
				t.Errorf("reconcileClustersForRegions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClusterManager_createSyncSet(t *testing.T) {
	const ingressDNS = "foo.bar.example.com"
	observabilityConfig := buildObservabilityConfig()
	syncset, err := buildSyncSet(observabilityConfig, ingressDNS)

	if err != nil {
		t.Fatal("Unable to create test syncset")
		return
	}

	type fields struct {
		ocmClient ocm.Client
		timer     *time.Timer
	}

	type result struct {
		err     error
		syncset *clustersmgmtv1.Syncset
	}
	tests := []struct {
		name   string
		fields fields
		want   result
	}{
		{
			name: "throw an error when syncset creation fails",
			fields: fields{
				ocmClient: &ocm.ClientMock{
					CreateSyncSetFunc: func(clusterId string, syncset *clustersmgmtv1.Syncset) (*clustersmgmtv1.Syncset, error) {
						return nil, errors.New("error when creating syncset")
					},
				},
			},
			want: result{
				err:     errors.New("error when creating syncset"),
				syncset: nil,
			},
		},
		{
			name: "returns created syncset",
			fields: fields{
				ocmClient: &ocm.ClientMock{
					CreateSyncSetFunc: func(clusterId string, syncset *clustersmgmtv1.Syncset) (*clustersmgmtv1.Syncset, error) {
						return syncset, nil
					},
				},
			},
			want: result{
				err:     nil,
				syncset: syncset,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterTestingT(t)

			c := &ClusterManager{
				ocmClient: tt.fields.ocmClient,
				timer:     tt.fields.timer,
				configService: services.NewConfigService(config.ApplicationConfig{
					SupportedProviders:         &config.ProviderConfig{},
					AllowList:                  &config.AllowListConfig{},
					ObservabilityConfiguration: &observabilityConfig,
					ClusterCreationConfig:      &config.ClusterCreationConfig{},
				}),
			}
			got, err := c.createSyncSet("clusterId", ingressDNS)
			Expect(got).To(Equal(tt.want.syncset))
			if err != nil {
				Expect(err).To(MatchError(tt.want.err))
			}
		})
	}
}

func TestClusterManager_reconcileAddonOperator(t *testing.T) {
	type fields struct {
		agentOperator  services.KasFleetshardOperatorAddon
		clusterService services.ClusterService
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "agent operator is ready",
			fields: fields{
				agentOperator: &services.KasFleetshardOperatorAddonMock{
					ProvisionFunc: func(cluster api.Cluster) (bool, *apiErrors.ServiceError) {
						return true, nil
					},
				},
				clusterService: &services.ClusterServiceMock{
					UpdateStatusFunc: func(cluster api.Cluster, status api.ClusterStatus) error {
						return nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "agent operator is not ready",
			fields: fields{
				agentOperator: &services.KasFleetshardOperatorAddonMock{
					ProvisionFunc: func(cluster api.Cluster) (bool, *apiErrors.ServiceError) {
						return false, nil
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ClusterManager{
				clusterService: tt.fields.clusterService,
				configService: services.NewConfigService(config.ApplicationConfig{
					SupportedProviders:         &config.ProviderConfig{},
					AllowList:                  &config.AllowListConfig{},
					ClusterCreationConfig:      &config.ClusterCreationConfig{EnableKasFleetshardOperator: true},
					ObservabilityConfiguration: &config.ObservabilityConfiguration{},
				}),
				kasFleetshardOperatorAddon: tt.fields.agentOperator,
			}

			err := c.reconcileAddonOperator(api.Cluster{
				ClusterID: "test-cluster-id",
			})
			if err != nil && !tt.wantErr {
				t.Errorf("reconcileAddonOperator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// buildObservabilityConfig builds a observability config used for testing
func buildObservabilityConfig() config.ObservabilityConfiguration {
	observabilityConfig := config.ObservabilityConfiguration{
		DexUrl:                         "dex-url",
		DexPassword:                    "dex-password",
		DexUsername:                    "dex-username",
		DexSecret:                      "dex-secret",
		ObservatoriumTenant:            "tenant",
		ObservatoriumGateway:           "gateway",
		ObservabilityConfigRepo:        "obs-config-repo",
		ObservabilityConfigChannel:     "obs-config-channel",
		ObservabilityConfigAccessToken: "obs-config-token",
	}
	return observabilityConfig
}

// buildSyncSet builds a syncset used for testing
func buildSyncSet(observabilityConfig config.ObservabilityConfiguration, ingressDNS string) (*clustersmgmtv1.Syncset, error) {
	reclaimDelete := k8sCoreV1.PersistentVolumeReclaimDelete
	expansion := true
	consumer := storagev1.VolumeBindingWaitForFirstConsumer
	r := ingressReplicas

	syncset, err := clustersmgmtv1.NewSyncset().
		ID(syncsetName).
		Resources([]interface{}{
			&storagev1.StorageClass{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "storage.k8s.io/v1",
					Kind:       "StorageClass",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: syncsetresources.KafkaStorageClass,
				},
				Parameters: map[string]string{
					"encrypted": "false",
					"type":      "gp2",
				},
				Provisioner:          "kubernetes.io/aws-ebs",
				ReclaimPolicy:        &reclaimDelete,
				AllowVolumeExpansion: &expansion,
				VolumeBindingMode:    &consumer,
			},
			&ingressoperatorv1.IngressController{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "operator.openshift.io/v1",
					Kind:       "IngressController",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "sharded",
					Namespace: openshiftIngressNamespace,
				},
				Spec: ingressoperatorv1.IngressControllerSpec{
					Domain: ingressDNS,
					RouteSelector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							syncsetresources.IngressLabelName: syncsetresources.IngressLabelValue,
						},
					},
					Replicas: &r,
					NodePlacement: &ingressoperatorv1.NodePlacement{
						NodeSelector: &metav1.LabelSelector{
							MatchLabels: map[string]string{
								"node-role.kubernetes.io/worker": "",
							},
						},
					},
				},
			},
			&projectv1.Project{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "project.openshift.io/v1",
					Kind:       "Project",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: observabilityNamespace,
				},
			},
			&k8sCoreV1.Secret{
				TypeMeta: metav1.TypeMeta{
					APIVersion: metav1.SchemeGroupVersion.Version,
					Kind:       "Secret",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      observabilityDexCredentials,
					Namespace: observabilityNamespace,
				},
				Type: k8sCoreV1.SecretTypeOpaque,
				StringData: map[string]string{
					"password": observabilityConfig.DexPassword,
					"secret":   observabilityConfig.DexSecret,
					"username": observabilityConfig.DexUsername,
				},
			},
			&v1alpha1.CatalogSource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "operators.coreos.com/v1alpha1",
					Kind:       "CatalogSource",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      observabilityCatalogSourceName,
					Namespace: observabilityNamespace,
				},
				Spec: v1alpha1.CatalogSourceSpec{
					SourceType: v1alpha1.SourceTypeGrpc,
					Image:      observabilityCatalogSourceImage,
				},
			},
			&v1alpha2.OperatorGroup{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "operators.coreos.com/v1alpha2",
					Kind:       "OperatorGroup",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      observabilityOperatorGroupName,
					Namespace: observabilityNamespace,
				},
				Spec: v1alpha2.OperatorGroupSpec{
					TargetNamespaces: []string{observabilityNamespace},
				},
			},
			&v1alpha1.Subscription{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "operators.coreos.com/v1alpha1",
					Kind:       "Subscription",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      observabilitySubscriptionName,
					Namespace: observabilityNamespace,
				},
				Spec: &v1alpha1.SubscriptionSpec{
					CatalogSource:          observabilityCatalogSourceName,
					Channel:                "alpha",
					CatalogSourceNamespace: observabilityNamespace,
					StartingCSV:            "observability-operator.v2.0.0",
					InstallPlanApproval:    v1alpha1.ApprovalAutomatic,
					Package:                observabilitySubscriptionName,
				},
			},
			&k8sCoreV1.ConfigMap{
				TypeMeta: metav1.TypeMeta{
					APIVersion: k8sCoreV1.SchemeGroupVersion.String(),
					Kind:       "ConfigMap",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      observabilityKafkaConfiguration,
					Namespace: observabilityNamespace,
					Labels: map[string]string{
						"configures": "observability-operator",
					},
				},
				Data: map[string]string{
					"access_token": observabilityConfig.ObservabilityConfigAccessToken,
					"channel":      observabilityConfig.ObservabilityConfigChannel,
					"repository":   observabilityConfig.ObservabilityConfigRepo,
				},
			},
		}...).
		Build()
	return syncset, err
}
