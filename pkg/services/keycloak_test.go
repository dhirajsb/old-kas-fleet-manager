package services

import (
	"context"
	"reflect"
	"testing"

	gocloak "github.com/Nerzal/gocloak/v8"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/auth"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/client/keycloak"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/config"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
)

const (
	token        = "token"
	testClientID = "12221"
	secret       = "secret"
)

func TestKeycloakService_RegisterKafkaClientInSSO(t *testing.T) {
	type fields struct {
		kcClient keycloak.KcClient
	}

	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "fetch kafka client secret from sso when client already exists",
			fields: fields{
				kcClient: &keycloak.KcClientMock{
					GetTokenFunc: func() (string, error) {
						return token, nil
					},
					GetConfigFunc: func() *config.KeycloakConfig {
						return config.NewKeycloakConfig()
					},
					IsClientExistFunc: func(clientId string, accessToken string) (string, error) {
						return testClientID, nil
					},
					GetClientSecretFunc: func(internalClientId string, accessToken string) (string, error) {
						return secret, nil
					},
				},
			},
			want:    secret,
			wantErr: false,
		},
		{
			name: "successfully register a new sso client for the kafka cluster",
			fields: fields{
				kcClient: &keycloak.KcClientMock{
					GetTokenFunc: func() (string, error) {
						return token, nil
					},
					GetConfigFunc: func() *config.KeycloakConfig {
						return config.NewKeycloakConfig()
					},
					IsClientExistFunc: func(clientId string, accessToken string) (string, error) {
						return "", nil
					},
					GetClientSecretFunc: func(internalClientId string, accessToken string) (string, error) {
						return secret, nil
					},
					CreateClientFunc: func(client gocloak.Client, accessToken string) (string, error) {
						return testClientID, nil
					},
					ClientConfigFunc: func(client keycloak.ClientRepresentation) gocloak.Client {
						testID := "12221"
						return gocloak.Client{
							ClientID: &testID,
						}
					},
				},
			},
			want:    secret,
			wantErr: false,
		},
		{
			name: "failed to register sso client for the kafka cluster",
			fields: fields{
				kcClient: &keycloak.KcClientMock{
					GetTokenFunc: func() (string, error) {
						return token, nil
					},
					GetConfigFunc: func() *config.KeycloakConfig {
						return config.NewKeycloakConfig()
					},
					IsClientExistFunc: func(clientId string, accessToken string) (string, error) {
						return "", nil
					},
					GetClientSecretFunc: func(internalClientId string, accessToken string) (string, error) {
						return secret, nil
					},
					CreateClientFunc: func(client gocloak.Client, accessToken string) (string, error) {
						return "", errors.GeneralError("failed to create the sso client")
					},
					ClientConfigFunc: func(client keycloak.ClientRepresentation) gocloak.Client {
						testID := "12221"
						return gocloak.Client{
							ClientID: &testID,
						}
					},
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keycloakService := keycloakService{
				tt.fields.kcClient,
			}
			got, err := keycloakService.RegisterKafkaClientInSSO("kafka-12212", "121212")
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterKafkaClientInSSO() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterKafkaClientInSSO() got = %+v, want %+v", got, tt.want)
			}
		})
	}

}

func TestNewKeycloakService_DeRegisterKafkaClientInSSO(t *testing.T) {
	type fields struct {
		kcClient keycloak.KcClient
	}

	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "successful deleted the kafka client in sso",
			fields: fields{
				kcClient: &keycloak.KcClientMock{
					GetTokenFunc: func() (string, error) {
						return token, nil
					},
					GetConfigFunc: func() *config.KeycloakConfig {
						return config.NewKeycloakConfig()
					},
					IsClientExistFunc: func(clientId string, accessToken string) (string, error) {
						return testClientID, nil
					},
					DeleteClientFunc: func(internalClientID string, accessToken string) error {
						return nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "failed to delete the kafka client from sso",
			fields: fields{
				kcClient: &keycloak.KcClientMock{
					GetTokenFunc: func() (string, error) {
						return token, nil
					},
					GetConfigFunc: func() *config.KeycloakConfig {
						return config.NewKeycloakConfig()
					},
					IsClientExistFunc: func(clientId string, accessToken string) (string, error) {
						return testClientID, nil
					},
					DeleteClientFunc: func(internalClientID string, accessToken string) error {
						return errors.GeneralError("failed to delete")
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keycloakService := keycloakService{
				tt.fields.kcClient,
			}
			err := keycloakService.DeRegisterKafkaClientInSSO(testClientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterKafkaClientInSSO() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

func TestKeycloakService_GetSecretForRegisteredKafkaClient(t *testing.T) {
	type fields struct {
		kcClient keycloak.KcClient
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "fetch kafka client secret for the existing client",
			fields: fields{
				kcClient: &keycloak.KcClientMock{
					GetTokenFunc: func() (string, error) {
						return token, nil
					},
					GetConfigFunc: func() *config.KeycloakConfig {
						return config.NewKeycloakConfig()
					},
					IsClientExistFunc: func(clientId string, accessToken string) (string, error) {
						return testClientID, nil
					},
					GetClientSecretFunc: func(internalClientId string, accessToken string) (string, error) {
						return secret, nil
					},
				},
			},
			want:    secret,
			wantErr: false,
		},
		{
			name: "failed to get the kafka client secret for an existing client",
			fields: fields{
				kcClient: &keycloak.KcClientMock{
					GetTokenFunc: func() (string, error) {
						return token, nil
					},
					GetConfigFunc: func() *config.KeycloakConfig {
						return config.NewKeycloakConfig()
					},
					IsClientExistFunc: func(clientId string, accessToken string) (string, error) {
						return "", errors.GeneralError("failed to get client")
					},
					GetClientSecretFunc: func(internalClientId string, accessToken string) (string, error) {
						return secret, nil
					},
				},
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keycloakService := keycloakService{
				tt.fields.kcClient,
			}
			got, err := keycloakService.GetSecretForRegisteredKafkaClient(testClientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterKafkaClientInSSO() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterKafkaClientInSSO() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestKeycloakService_CreateServiceAccount(t *testing.T) {
	type fields struct {
		kcClient keycloak.KcClient
	}

	type args struct {
		serviceAccountRequest *api.ServiceAccountRequest
		ctx                   context.Context
	}

	testServiceAccount := api.ServiceAccount{
		ID:           testClientID,
		ClientSecret: secret,
		Name:         "test-svc",
		Description:  "desc",
		ClientID:     "srvc-acct-cca1a262-9465-4878-9f76-c3bb59d4b4b5",
	}

	tests := []struct {
		name    string
		fields  fields
		want    *api.ServiceAccount
		wantErr bool
		args    args
	}{
		{
			name: "successfully created a service account in sso",
			fields: fields{
				kcClient: &keycloak.KcClientMock{
					GetTokenFunc: func() (string, error) {
						return token, nil
					},
					GetConfigFunc: func() *config.KeycloakConfig {
						return config.NewKeycloakConfig()
					},
					IsClientExistFunc: func(clientId string, accessToken string) (string, error) {
						return "", nil
					},
					GetClientSecretFunc: func(internalClientId string, accessToken string) (string, error) {
						return secret, nil
					},
					CreateClientFunc: func(client gocloak.Client, accessToken string) (string, error) {
						return testClientID, nil
					},
					ClientConfigFunc: func(client keycloak.ClientRepresentation) gocloak.Client {
						testID := "12221"
						return gocloak.Client{
							ClientID: &testID,
						}
					},
					GetClientServiceAccountFunc: func(accessToken string, internalClient string) (*gocloak.User, error) {
						id := "1"
						return &gocloak.User{
							ID: &id,
						}, nil
					},
					UpdateServiceAccountUserFunc: func(accessToken string, serviceAccountUser gocloak.User) error {
						return nil
					},
					CreateProtocolMapperConfigFunc: func(name string) []gocloak.ProtocolMapperRepresentation {
						return []gocloak.ProtocolMapperRepresentation{}
					},
				},
			},
			args: args{
				serviceAccountRequest: &api.ServiceAccountRequest{
					Name:        "test-svc",
					Description: "desc",
				},
				ctx: auth.SetOrgIdContext(context.TODO(), testClientID),
			},
			want:    &testServiceAccount,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keycloakService := keycloakService{
				tt.fields.kcClient,
			}
			got, err := keycloakService.CreateServiceAccount(tt.args.serviceAccountRequest, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateServiceAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
			//over-riding the random generate id
			got.ClientID = "srvc-acct-cca1a262-9465-4878-9f76-c3bb59d4b4b5"
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateServiceAccount() got = %+v, want %+v", got, tt.want)
			}
		})
	}

}

func TestKeycloakService_DeleteServiceAccount(t *testing.T) {
	type fields struct {
		kcClient keycloak.KcClient
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
		args    args
	}{
		{name: "successfully deleted service account",
			fields: fields{
				kcClient: &keycloak.KcClientMock{
					GetTokenFunc: func() (string, error) {
						return token, nil
					},
					GetConfigFunc: func() *config.KeycloakConfig {
						return config.NewKeycloakConfig()
					},
					IsClientExistFunc: func(clientId string, accessToken string) (string, error) {
						return "", nil
					},
					ClientConfigFunc: func(client keycloak.ClientRepresentation) gocloak.Client {
						testID := "12221"
						return gocloak.Client{
							ClientID: &testID,
						}
					},
					DeleteClientFunc: func(internalClientID string, accessToken string) error {
						return nil
					},
					GetClientByIdFunc: func(id string, accessToken string) (*gocloak.Client, error) {
						testID := "12221"
						return &gocloak.Client{
							ClientID: &testID,
						}, nil
					},
					IsSameOrgFunc: func(client *gocloak.Client, orgId string) bool {
						return true
					},
				},
			},
			args: args{
				ctx: auth.SetOrgIdContext(context.TODO(), testClientID),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keycloakService := keycloakService{
				tt.fields.kcClient,
			}
			err := keycloakService.DeleteServiceAccount(tt.args.ctx, testClientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterKafkaClientInSSO() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func TestKeycloakService_ListServiceAcc(t *testing.T) {
	type fields struct {
		kcClient keycloak.KcClient
	}

	type args struct {
		ctx context.Context
	}

	var testServiceAcc []api.ServiceAccount

	tests := []struct {
		name    string
		fields  fields
		want    []api.ServiceAccount
		wantErr bool
		args    args
	}{
		{
			name: "list service account",
			fields: fields{
				kcClient: &keycloak.KcClientMock{
					GetTokenFunc: func() (string, error) {
						return token, nil
					},
					GetConfigFunc: func() *config.KeycloakConfig {
						return config.NewKeycloakConfig()
					},
					IsClientExistFunc: func(clientId string, accessToken string) (string, error) {
						return "", nil
					},
					ClientConfigFunc: func(client keycloak.ClientRepresentation) gocloak.Client {
						testID := "12221"
						return gocloak.Client{
							ClientID: &testID,
						}
					},
					IsSameOrgFunc: func(client *gocloak.Client, orgId string) bool {
						return true
					},
					GetClientsFunc: func(accessToken string, first int, max int) ([]*gocloak.Client, error) {
						testClient := []*gocloak.Client{}
						return testClient, nil
					},
				},
			},
			args: args{
				ctx: auth.SetOrgIdContext(context.TODO(), "12221"),
			},
			want:    testServiceAcc,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keycloakService := keycloakService{
				tt.fields.kcClient,
			}
			got, err := keycloakService.ListServiceAcc(tt.args.ctx, 0, 10)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterKafkaClientInSSO() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterKafkaClientInSSO() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestKeycloakService_ResetServiceAccountCredentials(t *testing.T) {
	type fields struct {
		kcClient keycloak.KcClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		want    *api.ServiceAccount
		args    args
		wantErr bool
	}{
		{
			name: "Reset service account credentials",
			fields: fields{
				kcClient: &keycloak.KcClientMock{
					GetTokenFunc: func() (string, error) {
						return token, nil
					},
					GetConfigFunc: func() *config.KeycloakConfig {
						return config.NewKeycloakConfig()
					},
					IsClientExistFunc: func(clientId string, accessToken string) (string, error) {
						return "", nil
					},
					ClientConfigFunc: func(client keycloak.ClientRepresentation) gocloak.Client {
						testID := "12221"
						return gocloak.Client{
							ClientID: &testID,
						}
					},
					IsSameOrgFunc: func(client *gocloak.Client, orgId string) bool {
						return true
					},
					GetClientByIdFunc: func(id string, accessToken string) (*gocloak.Client, error) {
						testID := "12221"
						return &gocloak.Client{
							ID:       &testID,
							ClientID: &testID,
						}, nil
					},
					RegenerateClientSecretFunc: func(accessToken string, id string) (*gocloak.CredentialRepresentation, error) {
						sec := "secret"
						testID := "12221"
						return &gocloak.CredentialRepresentation{
							Value: &sec,
							ID:    &testID,
						}, nil
					},
				},
			},
			args: args{
				ctx: auth.SetOrgIdContext(context.TODO(), "12221"),
			},
			want: &api.ServiceAccount{
				ID:           "12221",
				ClientID:     "12221",
				ClientSecret: secret,
				Name:         "",
				Description:  "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keycloakService := keycloakService{
				tt.fields.kcClient,
			}
			got, err := keycloakService.ResetServiceAccountCredentials(tt.args.ctx, testClientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterKafkaClientInSSO() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterKafkaClientInSSO() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestKeycloakService_RegisterKasFleetshardOperatorServiceAccount(t *testing.T) {
	type fields struct {
		kcClient keycloak.KcClient
	}
	type args struct {
		clusterId string
		roleName  string
	}
	fakeRoleId := "1234"
	fakeClientId := "test-client-id"
	fakeClientSecret := "test-client-secret"
	fakeUserId := "test-user-id"
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.ServiceAccount
		wantErr bool
	}{
		{
			name: "test registering serviceaccount for agent operator first time",
			fields: fields{
				kcClient: &keycloak.KcClientMock{
					GetTokenFunc: func() (string, error) {
						return token, nil
					},
					AddRealmRoleToUserFunc: func(accessToken string, userId string, role gocloak.Role) error {
						return nil
					},
					CreateRealmRoleFunc: func(accessToken string, roleName string) (*gocloak.Role, error) {
						return &gocloak.Role{
							ID:   &fakeRoleId,
							Name: &roleName,
						}, nil
					},
					CreateClientFunc: func(client gocloak.Client, accessToken string) (string, error) {
						return fakeClientId, nil
					},
					GetClientFunc: func(clientId string, accessToken string) (*gocloak.Client, error) {
						return nil, nil
					},
					GetClientSecretFunc: func(internalClientId string, accessToken string) (string, error) {
						return fakeClientSecret, nil
					},
					GetClientServiceAccountFunc: func(accessToken string, internalClient string) (*gocloak.User, error) {
						return &gocloak.User{
							ID: &fakeUserId,
						}, nil
					},
					GetRealmRoleFunc: func(accessToken string, roleName string) (*gocloak.Role, error) {
						return nil, nil
					},
					UpdateServiceAccountUserFunc: func(accessToken string, serviceAccountUser gocloak.User) error {
						return nil
					},
					UserHasRealmRoleFunc: func(accessToken string, userId string, roleName string) (*gocloak.Role, error) {
						return nil, nil
					},
					CreateProtocolMapperConfigFunc: func(in1 string) []gocloak.ProtocolMapperRepresentation {
						return []gocloak.ProtocolMapperRepresentation{{}}
					},
					ClientConfigFunc: func(client keycloak.ClientRepresentation) gocloak.Client {
						return gocloak.Client{}
					},
				},
			},
			args: args{
				clusterId: "test-cluster-id",
				roleName:  "test-role-name",
			},
			want: &api.ServiceAccount{
				ID:           fakeClientId,
				ClientID:     "kas-fleetshard-agent-test-cluster-id",
				ClientSecret: fakeClientSecret,
				Name:         "kas-fleetshard-agent-test-cluster-id",
				Description:  "service account for agent on cluster test-cluster-id",
			},
			wantErr: false,
		},
		{
			name: "test registering serviceaccount for agent operator second time",
			fields: fields{
				kcClient: &keycloak.KcClientMock{
					GetTokenFunc: func() (string, error) {
						return token, nil
					},
					GetClientFunc: func(clientId string, accessToken string) (*gocloak.Client, error) {
						return &gocloak.Client{
							ID: &fakeClientId,
						}, nil
					},
					GetClientSecretFunc: func(internalClientId string, accessToken string) (string, error) {
						return fakeClientSecret, nil
					},
					GetClientServiceAccountFunc: func(accessToken string, internalClient string) (*gocloak.User, error) {
						return &gocloak.User{
							ID: &fakeUserId,
							Attributes: &map[string][]string{
								clusterId: {"test-cluster-id"},
							},
						}, nil
					},
					GetRealmRoleFunc: func(accessToken string, roleName string) (*gocloak.Role, error) {
						return &gocloak.Role{
							ID: &fakeRoleId,
						}, nil
					},
					UserHasRealmRoleFunc: func(accessToken string, userId string, roleName string) (*gocloak.Role, error) {
						return &gocloak.Role{
							ID: &fakeRoleId,
						}, nil
					},
					CreateProtocolMapperConfigFunc: func(in1 string) []gocloak.ProtocolMapperRepresentation {
						return []gocloak.ProtocolMapperRepresentation{{}}
					},
					ClientConfigFunc: func(client keycloak.ClientRepresentation) gocloak.Client {
						return gocloak.Client{}
					},
				},
			},
			args: args{
				clusterId: "test-cluster-id",
				roleName:  "test-role-name",
			},
			want: &api.ServiceAccount{
				ID:           fakeClientId,
				ClientID:     "kas-fleetshard-agent-test-cluster-id",
				ClientSecret: fakeClientSecret,
				Name:         "kas-fleetshard-agent-test-cluster-id",
				Description:  "service account for agent on cluster test-cluster-id",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keycloakService := keycloakService{
				tt.fields.kcClient,
			}
			got, err := keycloakService.RegisterKasFleetshardOperatorServiceAccount(tt.args.clusterId, tt.args.roleName)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterKasFleetshardOperatorServiceAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterKasFleetshardOperatorServiceAccount() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestKeycloakService_GetServiceAccountById(t *testing.T) {
	type fields struct {
		kcClient keycloak.KcClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		want    *api.ServiceAccount
		args    args
		wantErr bool
	}{
		{
			name: "Get service account by id",
			fields: fields{
				kcClient: &keycloak.KcClientMock{
					GetTokenFunc: func() (string, error) {
						return token, nil
					},
					GetConfigFunc: func() *config.KeycloakConfig {
						return config.NewKeycloakConfig()
					},
					IsClientExistFunc: func(clientId string, accessToken string) (string, error) {
						return "", nil
					},
					ClientConfigFunc: func(client keycloak.ClientRepresentation) gocloak.Client {
						testID := "12221"
						return gocloak.Client{
							ClientID: &testID,
						}
					},
					IsSameOrgFunc: func(client *gocloak.Client, orgId string) bool {
						return true
					},
					GetClientByIdFunc: func(id string, accessToken string) (*gocloak.Client, error) {
						testID := "12221"
						return &gocloak.Client{
							ID:       &testID,
							ClientID: &testID,
						}, nil
					},
				},
			},
			args: args{
				ctx: auth.SetOrgIdContext(context.TODO(), "12221"),
			},
			want: &api.ServiceAccount{
				ID:          "12221",
				ClientID:    "12221",
				Name:        "",
				Description: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keycloakService := keycloakService{
				tt.fields.kcClient,
			}
			got, err := keycloakService.GetServiceAccountById(tt.args.ctx, testClientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetServiceAccountById() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetServiceAccountById() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}
