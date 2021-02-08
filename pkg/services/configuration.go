package services

import (
	"errors"
	"fmt"

	"gitlab.cee.redhat.com/service/managed-services-api/pkg/config"
)

// ConfigService is a service used for managing and accessing the various configurations used by the overall service
type ConfigService interface {
	// GetSupportedProviders returns the current supported providers in this service
	GetSupportedProviders() config.ProviderList
	// GetDefaultProvider returns the default provider in the supported providers configuration, if multiple providers
	// are specified as default then the first will be returned
	GetDefaultProvider() (config.Provider, error)
	// GetDefaultRegionForProvider returns the default region specified in the provider, if multiple regions are
	// specified as default then the first will be returned
	GetDefaultRegionForProvider(provider config.Provider) (config.Region, error)
	// IsProviderSupported returns true if the given provider is supported
	IsProviderSupported(providerName string) bool
	// IsRegionSupportedForProvider returns true if the provided region is supported in the given provider
	IsRegionSupportedForProvider(providerName, regionName string) bool
	// IsAllowListEnabled returns true if the allow access list is feature is enable for access control
	IsAllowListEnabled() bool
	// GetOrganisationById returns the organisaion by the given id
	GetOrganisationById(orgId string) (config.Organisation, bool)
	// GetAllowedAccountByUsernameAndOrgId returns the allowed user in a given organisation (if found organisation is found), else return user by from the global list
	GetAllowedAccountByUsernameAndOrgId(username string, orgId string) (config.AllowedAccount, bool)
	// GetServiceAccountByUsername returns allowed account by from the list of service accounts
	GetServiceAccountByUsername(username string) (config.AllowedAccount, bool)
	// Validate ensures all configuration managed by the service contains correct and valid values
	Validate() error
	// IsAutoCreateOSDEnabled returns true if the automatic creation of OSD cluster is enabled, false otherwise.
	IsAutoCreateOSDEnabled() bool
	// GetObservabilityConfiguration returns ObservabilityConfiguration.
	GetObservabilityConfiguration() config.ObservabilityConfiguration
	// IsKasFleetshardOperatorEnabled returns if the agent operator should be used when creating a new OSD cluster
	IsKasFleetshardOperatorEnabled() bool
}

var _ ConfigService = &configService{}

// configService is an internal implementation of ConfigService
type configService struct {
	// providersConfig is the supported providers managed by the service
	providersConfig config.ProviderConfiguration

	// allowListConfig is the list of users allowed to access the service
	allowListConfig config.AllowListConfig

	// clusterCreationConfig is the cluster creation configuration
	clusterCreationConfig config.ClusterCreationConfig

	// Observability coniguration
	observabilityConfig config.ObservabilityConfiguration
}

// NewConfigService returns a new default implementation of ConfigService
// TODO pass ApplicationConfiguration instead
func NewConfigService(providersConfig config.ProviderConfiguration, allowListConfig config.AllowListConfig, clusterCreationConfig config.ClusterCreationConfig, observabilityConfig config.ObservabilityConfiguration) ConfigService {
	return &configService{
		providersConfig:       providersConfig,
		allowListConfig:       allowListConfig,
		clusterCreationConfig: clusterCreationConfig,
		observabilityConfig:   observabilityConfig,
	}
}

func (c configService) GetSupportedProviders() config.ProviderList {
	return c.providersConfig.SupportedProviders
}

func (c configService) GetDefaultProvider() (config.Provider, error) {
	for _, p := range c.providersConfig.SupportedProviders {
		if p.Default {
			return p, nil
		}
	}
	return config.Provider{}, errors.New("no default provider found in list of supported providers")
}

func (c configService) GetDefaultRegionForProvider(provider config.Provider) (config.Region, error) {
	for _, r := range provider.Regions {
		if r.Default {
			return r, nil
		}
	}
	return config.Region{}, fmt.Errorf("no default region found for provider %s", provider.Name)
}

func (c configService) IsProviderSupported(providerName string) bool {
	_, ok := c.providersConfig.SupportedProviders.GetByName(providerName)
	return ok
}

func (c configService) IsRegionSupportedForProvider(providerName, regionName string) bool {
	provider, ok := c.providersConfig.SupportedProviders.GetByName(providerName)
	if !ok {
		return false
	}
	_, ok = provider.Regions.GetByName(regionName)
	return ok
}

func (c configService) IsAllowListEnabled() bool {
	return c.allowListConfig.EnableAllowList
}

func (c configService) GetOrganisationById(orgId string) (config.Organisation, bool) {
	return c.allowListConfig.AllowList.Organisations.GetById(orgId)
}

// GetServiceAccountByUsername returns allowed account by from the list of service accounts
func (c configService) GetServiceAccountByUsername(username string) (config.AllowedAccount, bool) {
	return c.allowListConfig.AllowList.ServiceAccounts.GetByUsername(username)
}

// GetAllowedAccountByUsernameAndOrgId returns the allowed user in a given organisation (if found organisation is found),
// else return user by from the global list
func (c configService) GetAllowedAccountByUsernameAndOrgId(username string, orgId string) (config.AllowedAccount, bool) {
	var user config.AllowedAccount
	var found bool
	org, _ := c.GetOrganisationById(orgId)
	user, found = org.AllowedAccounts.GetByUsername(username)
	if found {
		return user, found
	}

	return c.GetServiceAccountByUsername(username)
}

func (c configService) Validate() error {
	providerDefaultCount := 0
	for _, p := range c.providersConfig.SupportedProviders {
		if err := c.validateProvider(p); err != nil {
			return err
		}
		if p.Default {
			providerDefaultCount++
		}
	}
	if providerDefaultCount != 1 {
		return fmt.Errorf("expected 1 default provider in provider list, got %d", providerDefaultCount)
	}
	return nil
}

// validateProvider ensures a given provider contains correct and valid values, including it's regions
func (c configService) validateProvider(provider config.Provider) error {
	defaultCount := 0
	for _, p := range provider.Regions {
		if p.Default {
			defaultCount++
		}
	}
	if defaultCount != 1 {
		return fmt.Errorf("expected 1 default region in provider %s, got %d", provider.Name, defaultCount)
	}
	return nil
}

func (c configService) IsAutoCreateOSDEnabled() bool {
	return c.clusterCreationConfig.AutoOSDCreation
}

// GetObservabilityConfiguration returns ObservabilityConfiguration.
func (c configService) GetObservabilityConfiguration() config.ObservabilityConfiguration {
	return c.observabilityConfig
}

func (c configService) IsKasFleetshardOperatorEnabled() bool {
	return c.clusterCreationConfig.EnableKasFleetshardOperator
}
