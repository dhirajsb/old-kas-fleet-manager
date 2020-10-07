package kafka

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"gitlab.cee.redhat.com/service/managed-services-api/cmd/managed-services-api/environments"
	"gitlab.cee.redhat.com/service/managed-services-api/cmd/managed-services-api/flags"
	customOcm "gitlab.cee.redhat.com/service/managed-services-api/pkg/ocm"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/services"
)

// NewDeleteCommand command for deleting kafkas.
func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a managed-services-api kafka request",
		Long:  "Delete a managed-services-api kafka request.",
		Run:   runDelete,
	}
	err := environments.Environment().AddFlags(cmd.PersistentFlags())
	if err != nil {
		glog.Fatalf("Unable to add environment flags to serve command: %s", err.Error())
	}

	cmd.Flags().String(FlagID, "", "Kafka id")

	return cmd
}

func runDelete(cmd *cobra.Command, _ []string) {
	id := flags.MustGetDefinedString(FlagID, cmd.Flags())

	if err := environments.Environment().Initialize(); err != nil {
		glog.Fatalf("Unable to initialize environment: %s", err.Error())
	}

	env := environments.Environment()

	// setup required services
	ocmClient := customOcm.NewClient(env.Clients.OCM.Connection)
	clusterService := services.NewClusterService(env.DBFactory, ocmClient, env.Config.AWS)
	syncsetService := services.NewSyncsetService(env.Clients.OCM.Connection)
	kafkaService := services.NewKafkaService(env.DBFactory, syncsetService, clusterService)

	err := kafkaService.Delete(id)
	if err != nil {
		glog.Fatalf("Unable to delete kafka request: %s", err.Error())
	}

	glog.V(10).Infof("Deleted kafka request with id %s", id)
}