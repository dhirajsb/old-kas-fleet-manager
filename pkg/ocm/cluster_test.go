package ocm

import (
	clustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/api"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/clusterservicetest"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/config"
	"reflect"
	"testing"
)

func Test_clusterBuilder_NewOCMClusterFromCluster(t *testing.T) {
	awsConfig := &config.AWSConfig{}
	clusterAWS := clustersmgmtv1.NewAWS().AccountID(awsConfig.AccountID).AccessKeyID(awsConfig.AccessKey).SecretAccessKey(awsConfig.SecretAccessKey)
	type fields struct {
		idGenerator IDGenerator
		awsConfig   *config.AWSConfig
	}
	type args struct {
		cluster *api.Cluster
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantFn  func() *clustersmgmtv1.Cluster
		wantErr bool
	}{
		{
			name: "nil aws config results in error",
			fields: fields{
				idGenerator: NewIDGenerator(""),
				awsConfig:   nil,
			},
			args: args{
				cluster: &api.Cluster{},
			},
			wantErr: true,
		},
		{
			name: "nil cluster results in error",
			fields: fields{
				idGenerator: NewIDGenerator(""),
				awsConfig:   awsConfig,
			},
			args: args{
				cluster: nil,
			},
			wantErr: true,
		},
		{
			name: "nil id generator results in error",
			fields: fields{
				idGenerator: nil,
				awsConfig:   awsConfig,
			},
			args: args{
				cluster: &api.Cluster{},
			},
			wantErr: true,
		},
		{
			name: "successful conversion of all supported provided values",
			fields: fields{
				idGenerator: &IDGeneratorMock{
					GenerateFunc: func() string {
						return ""
					},
				},
				awsConfig: awsConfig,
			},
			args: args{
				cluster: &api.Cluster{
					CloudProvider: clusterservicetest.MockClusterCloudProvider,
					ClusterID:     clusterservicetest.MockClusterID,
					ExternalID:    clusterservicetest.MockClusterExternalID,
					Region:        clusterservicetest.MockClusterRegion,
					State:         clusterservicetest.MockClusterState,
					BYOC:          clusterservicetest.MockClusterBYOC,
					Managed:       clusterservicetest.MockClusterManaged,
					MultiAZ:       clusterservicetest.MockClusterMultiAZ,
				},
			},
			wantFn: func() *clustersmgmtv1.Cluster {
				cluster, err := clusterservicetest.NewMockCluster(func(builder *clustersmgmtv1.ClusterBuilder) {
					// these values will be ignored by the conversion as they're unsupported. so expect different
					// values than we provide.
					builder.BYOC(true)
					builder.Managed(true)
					builder.Name("")
					builder.AWS(clusterAWS)
				})
				if err != nil {
					panic(err)
				}
				return cluster
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if tt.wantFn == nil {
			tt.wantFn = func() *clustersmgmtv1.Cluster {
				return nil
			}
		}
		t.Run(tt.name, func(t *testing.T) {
			r := clusterBuilder{
				idGenerator: tt.fields.idGenerator,
				awsConfig:   tt.fields.awsConfig,
			}
			got, err := r.NewOCMClusterFromCluster(tt.args.cluster)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOCMClusterFromCluster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantFn()) {
				t.Errorf("NewOCMClusterFromCluster() got = %+v, want %+v", got, tt.wantFn())
			}
		})
	}
}
