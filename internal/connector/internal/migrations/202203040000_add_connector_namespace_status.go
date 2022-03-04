package migrations

// Migrations should NEVER use types from other packages. Types can change
// and then migrations run on a _new_ database will fail or behave unexpectedly.
// Instead of importing types, always re-create the type in the migration, as
// is done here, even though the same type is defined in pkg/api

import (
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/db"
	"github.com/go-gormigrate/gormigrate/v2"
)

func addConnectorNamespaceStatus(migrationId string) *gormigrate.Migration {

	type ConnectorNamespacePhaseEnum = string

	type ConnectorNamespaceStatus struct {
		Phase              ConnectorNamespacePhaseEnum `gorm:"not null;index"`
		ConnectorsDeployed int                         `gorm:"not null"`
	}

	type ConnectorNamespace struct {
		Status ConnectorNamespaceStatus `gorm:"embedded;embeddedPrefix:status_"`
	}

	return db.CreateMigrationFromActions(migrationId,
		// add namespace status
		db.AddTableColumnsAction(&ConnectorNamespace{}),
	)
}
