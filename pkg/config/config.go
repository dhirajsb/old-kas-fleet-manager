package config

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
)

var projectRootDirectory = GetProjectRootDir()

type ApplicationConfig struct {
	Server                     *ServerConfig               `json:"server"`
	Metrics                    *MetricsConfig              `json:"metrics"`
	HealthCheck                *HealthCheckConfig          `json:"health_check"`
	Database                   *DatabaseConfig             `json:"database"`
	OCM                        *OCMConfig                  `json:"ocm"`
	Sentry                     *SentryConfig               `json:"sentry"`
	AWS                        *AWSConfig                  `json:"aws"`
	SupportedProviders         *ProviderConfig             `json:"providers"`
	AllowList                  *AllowListConfig            `json:"allow_list"`
	ObservabilityConfiguration *ObservabilityConfiguration `json:"observability"`
	Keycloak                   *KeycloakConfig             `json:"keycloak"`
	Kafka                      *KafkaConfig                `json:"kafka_tls"`
	ClusterCreationConfig      *ClusterCreationConfig      `json:"cluster_creation"`
	ConnectorsConfig           *ConnectorsConfig           `json:"connectors"`
}

func NewApplicationConfig() *ApplicationConfig {
	return &ApplicationConfig{
		Server:                     NewServerConfig(),
		Metrics:                    NewMetricsConfig(),
		HealthCheck:                NewHealthCheckConfig(),
		Database:                   NewDatabaseConfig(),
		OCM:                        NewOCMConfig(),
		Sentry:                     NewSentryConfig(),
		AWS:                        NewAWSConfig(),
		SupportedProviders:         NewSupportedProvidersConfig(),
		AllowList:                  NewAllowListConfig(),
		ObservabilityConfiguration: NewObservabilityConfigurationConfig(),
		Keycloak:                   NewKeycloakConfig(),
		Kafka:                      NewKafkaConfig(),
		ClusterCreationConfig:      NewClusterCreationConfig(),
		ConnectorsConfig:           NewConnectorsConfig(),
	}
}

func (c *ApplicationConfig) AddFlags(flagset *pflag.FlagSet) {
	flagset.AddGoFlagSet(flag.CommandLine)
	c.Server.AddFlags(flagset)
	c.Metrics.AddFlags(flagset)
	c.HealthCheck.AddFlags(flagset)
	c.Database.AddFlags(flagset)
	c.OCM.AddFlags(flagset)
	c.Sentry.AddFlags(flagset)
	c.AWS.AddFlags(flagset)
	c.SupportedProviders.AddFlags(flagset)
	c.AllowList.AddFlags(flagset)
	c.ObservabilityConfiguration.AddFlags(flagset)
	c.Keycloak.AddFlags(flagset)
	c.Kafka.AddFlags(flagset)
	c.ClusterCreationConfig.AddFlags(flagset)
	c.ConnectorsConfig.AddFlags(flagset)
}

func (c *ApplicationConfig) ReadFiles() error {
	err := c.Server.ReadFiles()
	if err != nil {
		return err
	}
	err = c.Metrics.ReadFiles()
	if err != nil {
		return err
	}
	err = c.HealthCheck.ReadFiles()
	if err != nil {
		return err
	}
	err = c.Database.ReadFiles()
	if err != nil {
		return err
	}
	err = c.OCM.ReadFiles()
	if err != nil {
		return err
	}
	err = c.Sentry.ReadFiles()
	if err != nil {
		return err
	}
	err = c.AWS.ReadFiles()
	if err != nil {
		return err
	}
	err = c.SupportedProviders.ReadFiles()
	if err != nil {
		return err
	}
	err = c.ObservabilityConfiguration.ReadFiles()
	if err != nil {
		return err
	}
	err = c.Keycloak.ReadFiles()
	if err != nil {
		return err
	}
	if c.AllowList.EnableAllowList {
		err = c.AllowList.ReadFiles()
		if err != nil {
			return err
		}
	}
	err = c.Kafka.ReadFiles()
	if err != nil {
		return err
	}
	if c.ConnectorsConfig.Enabled {
		err = c.ConnectorsConfig.ReadFiles()
		if err != nil {
			return err
		}
	}
	return nil
}

// Read the contents of file into integer value
func readFileValueInt(file string, val *int) error {
	fileContents, err := readFile(file)
	if err != nil {
		return err
	}

	*val, err = strconv.Atoi(fileContents)
	return err
}

// Read the contents of file into string value
func readFileValueString(file string, val *string) error {
	fileContents, err := readFile(file)
	if err != nil {
		return err
	}

	*val = strings.TrimSuffix(fileContents, "\n")
	return err
}

// Read the contents of file into boolean value
func readFileValueBool(file string, val *bool) error {
	fileContents, err := readFile(file)
	if err != nil {
		return err
	}

	*val, err = strconv.ParseBool(fileContents)
	return err
}

func readFile(file string) (string, error) {
	// If the value is in quotes, unquote it
	unquotedFile, err := strconv.Unquote(file)
	if err != nil {
		// values without quotes will raise an error, ignore it.
		unquotedFile = file
	}

	// If no file is provided, leave val unchanged.
	if unquotedFile == "" {
		return "", nil
	}

	// Ensure the absolute file path is used
	absFilePath := unquotedFile
	if !filepath.IsAbs(unquotedFile) {
		absFilePath = filepath.Join(projectRootDirectory, unquotedFile)
	}

	// Read the file
	buf, err := ioutil.ReadFile(absFilePath)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// GetProjectRootDir returns the root directory of the project.
// The root directory of the project is the directory that contains the go.mod file which contains
// the "github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager" module name.
func GetProjectRootDir() string {
	workingDir, err := os.Getwd()
	if err != nil {
		glog.Fatal(err)
	}
	dirs := strings.Split(workingDir, "/")
	var goModPath string
	var rootPath string
	for _, d := range dirs {
		rootPath = rootPath + "/" + d
		goModPath = rootPath + "/go.mod"
		goModFile, err := ioutil.ReadFile(goModPath)
		if err != nil { // if the file doesn't exist, continue searching
			continue
		}
		// The project root directory is obtained based on the assumption that module name,
		// "github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager", is contained in the 'go.mod' file.
		// Should the module name change in the code repo then it needs to be changed here too.
		if strings.Contains(string(goModFile), "github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager") {
			break
		}
	}
	return rootPath
}
