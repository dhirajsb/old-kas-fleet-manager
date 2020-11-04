package environments

import (
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/db"
)

var stageConfigDefaults map[string]string = map[string]string{
	"ocm-base-url":    "https://api.stage.openshift.com",
	"enable-ocm-mock": "false",
}

func loadStage(env *Env) error {
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
