package environments

import (
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/db"
)

var productionConfigDefaults map[string]string = map[string]string{
	"v":                                 "1",
	"ocm-debug":                         "false",
	"enable-ocm-mock":                   "false",
	"enable-sentry":                     "true",
	"enable-allow-list":                 "true",
	"max-allowed-instances":             "1",
	"auto-osd-creation":                 "true",
	"mas-sso-realm":                     "mas-sso",
	"mas-sso-base-url":                  "https://keycloak-edge-redhat-rhoam-user-sso.apps.mas-sso-stage.1gzl.s1.devshift.org",
	"enable-kafka-external-certificate": "true",
	"cluster-compute-machine-type":      "m5.4xlarge",
}

func loadProduction(env *Env) error {
	env.DBFactory = db.NewConnectionFactory(env.Config.Database)

	err := env.LoadClients()
	if err != nil {
		return err
	}
	err = env.LoadServices()
	if err != nil {
		return err
	}

	return env.InitializeSentry()
}
