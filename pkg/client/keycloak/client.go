package keycloak

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/Nerzal/gocloak/v8"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/config"
)

//go:generate moq -out client_moq.go . KcClient
type KcClient interface {
	CreateClient(client gocloak.Client, accessToken string) (string, error)
	GetToken() (string, error)
	GetClientSecret(internalClientId string, accessToken string) (string, error)
	DeleteClient(internalClientID string, accessToken string) error
	GetClient(clientId string, accessToken string) (*gocloak.Client, error)
	IsClientExist(clientId string, accessToken string) (string, error)
	GetConfig() *config.KeycloakConfig
	GetClientById(id string, accessToken string) (*gocloak.Client, error)
	ClientConfig(client ClientRepresentation) gocloak.Client
	CreateProtocolMapperConfig(string) []gocloak.ProtocolMapperRepresentation
	GetClientServiceAccount(accessToken string, internalClient string) (*gocloak.User, error)
	UpdateServiceAccountUser(accessToken string, serviceAccountUser gocloak.User) error
	GetClients(accessToken string, first int, max int) ([]*gocloak.Client, error)
	IsSameOrg(client *gocloak.Client, orgId string) bool
	RegenerateClientSecret(accessToken string, id string) (*gocloak.CredentialRepresentation, error)
	GetRealmRole(accessToken string, roleName string) (*gocloak.Role, error)
	CreateRealmRole(accessToken string, roleName string) (*gocloak.Role, error)
	UserHasRealmRole(accessToken string, userId string, roleName string) (*gocloak.Role, error)
	AddRealmRoleToUser(accessToken string, userId string, role gocloak.Role) error
}

type ClientRepresentation struct {
	Name                         string
	ClientID                     string
	ServiceAccountsEnabled       bool
	Secret                       string
	StandardFlowEnabled          bool
	Attributes                   map[string]string
	AuthorizationServicesEnabled bool
	ProtocolMappers              []gocloak.ProtocolMapperRepresentation
	Description                  string
}

type kcClient struct {
	kcClient gocloak.GoCloak
	ctx      context.Context
	config   *config.KeycloakConfig
}

var _ KcClient = &kcClient{}

func NewClient(config *config.KeycloakConfig) *kcClient {
	setTokenEndpoints(config)
	client := gocloak.NewClient(config.BaseURL)
	client.RestyClient().SetDebug(config.Debug)
	client.RestyClient().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: config.InsecureSkipVerify})
	return &kcClient{
		kcClient: client,
		ctx:      context.Background(),
		config:   config,
	}
}

func (kc *kcClient) ClientConfig(client ClientRepresentation) gocloak.Client {
	return gocloak.Client{
		Name:                         &client.Name,
		ClientID:                     &client.ClientID,
		ServiceAccountsEnabled:       &client.ServiceAccountsEnabled,
		StandardFlowEnabled:          &client.StandardFlowEnabled,
		Attributes:                   &client.Attributes,
		AuthorizationServicesEnabled: &client.AuthorizationServicesEnabled,
		ProtocolMappers:              &client.ProtocolMappers,
		Description:                  &client.Description,
	}
}

func (kc *kcClient) CreateProtocolMapperConfig(name string) []gocloak.ProtocolMapperRepresentation {
	proto := "openid-connect"
	mapper := "oidc-usermodel-attribute-mapper"
	protocolMapper := []gocloak.ProtocolMapperRepresentation{
		{
			Name:           &name,
			Protocol:       &proto,
			ProtocolMapper: &mapper,
			Config: &map[string]string{
				"access.token.claim":   "true",
				"claim.name":           name,
				"id.token.claim":       "true",
				"jsonType.label":       "String",
				"user.attribute":       name,
				"userinfo.token.claim": "true",
			},
		},
	}
	return protocolMapper
}

func setTokenEndpoints(config *config.KeycloakConfig) {
	config.JwksEndpointURI = config.BaseURL + "/auth/realms/" + config.Realm + "/protocol/openid-connect/certs"
	config.TokenEndpointURI = config.BaseURL + "/auth/realms/" + config.Realm + "/protocol/openid-connect/token"
	config.ValidIssuerURI = config.BaseURL + "/auth/realms/" + config.Realm
}

func (kc *kcClient) CreateClient(client gocloak.Client, accessToken string) (string, error) {
	internalClientID, err := kc.kcClient.CreateClient(kc.ctx, accessToken, kc.config.Realm, client)
	if err != nil {
		return "", fmt.Errorf("%+v", err.Error())
	}
	return internalClientID, err
}

func (kc *kcClient) GetClient(clientId string, accessToken string) (*gocloak.Client, error) {
	params := gocloak.GetClientsParams{
		ClientID: &clientId,
	}
	client, err := kc.kcClient.GetClients(kc.ctx, accessToken, kc.config.Realm, params)
	if err != nil {
		return nil, fmt.Errorf("%+v", err.Error())
	}
	if len(client) > 0 {
		return client[0], nil
	} else {
		return nil, nil
	}
}

func (kc *kcClient) GetToken() (string, error) {
	options := gocloak.TokenOptions{
		ClientID:     &kc.config.ClientID,
		GrantType:    &kc.config.GrantType,
		ClientSecret: &kc.config.ClientSecret,
	}
	tokenResp, err := kc.kcClient.GetToken(kc.ctx, kc.config.Realm, options)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve the token: %+v", err.Error())
	}
	return tokenResp.AccessToken, nil
}

func (kc *kcClient) GetClientSecret(internalClientId string, accessToken string) (string, error) {
	resp, err := kc.kcClient.GetClientSecret(kc.ctx, accessToken, kc.config.Realm, internalClientId)
	if err != nil {
		return "", fmt.Errorf("%+v", err.Error())
	}
	value := *resp.Value
	return value, err
}

func (kc *kcClient) DeleteClient(internalClientID string, accessToken string) error {
	err := kc.kcClient.DeleteClient(kc.ctx, accessToken, kc.config.Realm, internalClientID)
	if err != nil {
		return fmt.Errorf("%+v", err.Error())
	}
	return err
}

func (kc *kcClient) getClient(clientId string, accessToken string) ([]*gocloak.Client, error) {
	params := gocloak.GetClientsParams{
		ClientID: &clientId,
	}
	client, err := kc.kcClient.GetClients(kc.ctx, accessToken, kc.config.Realm, params)
	if err != nil {
		return nil, fmt.Errorf("%+v", err.Error())
	}
	return client, err
}

func (kc *kcClient) GetClientById(id string, accessToken string) (*gocloak.Client, error) {
	client, err := kc.kcClient.GetClient(kc.ctx, accessToken, kc.config.Realm, id)
	if err != nil {
		return nil, fmt.Errorf("client id: %s:%+v", id, err.Error())
	}
	return client, err
}

func (kc *kcClient) GetConfig() *config.KeycloakConfig {
	return kc.config
}

func (kc *kcClient) IsClientExist(clientId string, accessToken string) (string, error) {
	client, err := kc.getClient(clientId, accessToken)
	var internalClientID string
	if err != nil {
		return internalClientID, fmt.Errorf("client id: %s:%+v", clientId, err.Error())
	}
	if len(client) > 0 {
		internalClientID = *client[0].ID
		return internalClientID, nil
	}
	return internalClientID, err
}

func (kc *kcClient) GetClientServiceAccount(accessToken string, internalClient string) (*gocloak.User, error) {
	serviceAccountUser, err := kc.kcClient.GetClientServiceAccount(kc.ctx, accessToken, kc.config.Realm, internalClient)
	if err != nil {
		return nil, fmt.Errorf("%+v", err.Error())
	}
	return serviceAccountUser, err
}

func (kc *kcClient) UpdateServiceAccountUser(accessToken string, serviceAccountUser gocloak.User) error {
	err := kc.kcClient.UpdateUser(kc.ctx, accessToken, kc.config.Realm, serviceAccountUser)
	if err != nil {
		return fmt.Errorf("%+v", err.Error())
	}
	return err
}

func (kc *kcClient) GetClients(accessToken string, first int, max int) ([]*gocloak.Client, error) {
	params := gocloak.GetClientsParams{}
	if first > 0 && max > 0 {
		params = gocloak.GetClientsParams{
			First: &first,
			Max:   &max,
		}
	}
	clients, err := kc.kcClient.GetClients(kc.ctx, accessToken, kc.config.Realm, params)
	if err != nil {
		return nil, fmt.Errorf("%+v", err.Error())
	}
	return clients, err
}

func (kc *kcClient) IsSameOrg(client *gocloak.Client, orgId string) bool {
	if orgId == "" {
		return false
	}
	attributes := *client.Attributes
	return attributes["rh-org-id"] == orgId
}

func (kc *kcClient) RegenerateClientSecret(accessToken string, id string) (*gocloak.CredentialRepresentation, error) {
	credRep, err := kc.kcClient.RegenerateClientSecret(kc.ctx, accessToken, kc.config.Realm, id)
	if err != nil {
		return nil, fmt.Errorf("client id: %s:%+v", id, err.Error())
	}
	return credRep, err
}

func (kc *kcClient) GetRealmRole(accessToken string, roleName string) (*gocloak.Role, error) {
	r, err := kc.kcClient.GetRealmRole(kc.ctx, accessToken, kc.config.Realm, roleName)
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("realm role: %s:%+v", roleName, err.Error())
	}
	return r, err
}

func (kc *kcClient) CreateRealmRole(accessToken string, roleName string) (*gocloak.Role, error) {
	r := &gocloak.Role{
		Name: &roleName,
	}
	_, err := kc.kcClient.CreateRealmRole(kc.ctx, accessToken, kc.config.Realm, *r)
	if err != nil {
		return nil, fmt.Errorf("realm role: %s:%+v", roleName, err.Error())
	}
	// for some reason, the internal id of the role is not returned by kcClient.CreateRealmRole, so we have to get the role again to get the full details
	r, err = kc.kcClient.GetRealmRole(kc.ctx, accessToken, kc.config.Realm, roleName)
	if err != nil {
		return nil, fmt.Errorf("failed to create realm role: %s:%+v", roleName, err.Error())
	}
	return r, nil
}

func (kc *kcClient) UserHasRealmRole(accessToken string, userId string, roleName string) (*gocloak.Role, error) {
	roles, err := kc.kcClient.GetRealmRolesByUserID(kc.ctx, accessToken, kc.config.Realm, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to list realm roles for user: %s:%+v", userId, err.Error())
	}
	for _, r := range roles {
		if *r.Name == roleName {
			return r, nil
		}
	}
	return nil, nil
}

func (kc *kcClient) AddRealmRoleToUser(accessToken string, userId string, role gocloak.Role) error {
	roles := []gocloak.Role{role}
	err := kc.kcClient.AddRealmRoleToUser(kc.ctx, accessToken, kc.config.Realm, userId, roles)
	if err != nil {
		return fmt.Errorf("failed to add realm role to user: %s:%s:%+v", userId, *role.Name, err.Error())
	}
	return nil
}

func isNotFoundError(err error) bool {
	if e, ok := err.(*gocloak.APIError); ok {
		if e.Code == 404 {
			return true
		}
	}
	return false
}
