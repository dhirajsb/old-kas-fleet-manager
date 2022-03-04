package phase

import (
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/connector/internal/api/dbapi"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
	"github.com/looplab/fsm"
)

type ConnectorOperation string

const (
	AssignConnector      ConnectorOperation = "assign"
	UpdateConnector      ConnectorOperation = "update"
	StopConnector        ConnectorOperation = "stop"
	ProvisionConnector   ConnectorOperation = "provision"
	DeprovisionConnector ConnectorOperation = "deprovision"
	DeleteConnector      ConnectorOperation = "delete"
)

// ConnectorFSM handles connector phase changes within it's cluster's current phase
type ConnectorFSM struct {
	namespacePhase dbapi.ConnectorNamespacePhaseEnum
	Connector      *dbapi.Connector
	fsm            *fsm.FSM
}

var connectorEvents = map[dbapi.ConnectorNamespacePhaseEnum][]fsm.EventDesc{
	dbapi.ConnectorNamespacePhaseDisconnected: {
		{
			Name: string(AssignConnector),
			Src: []string{
				dbapi.ConnectorStatusPhaseAssigning,
			},
			Dst: dbapi.ConnectorStatusPhaseAssigned,
		},
		{
			Name: string(UpdateConnector),
			Src: []string{
				dbapi.ConnectorStatusPhaseAssigned,
			},
			Dst: dbapi.ConnectorStatusPhaseUpdating,
		},
	},
	dbapi.ConnectorNamespacePhaseReady: {
		{
			Name: string(AssignConnector),
			Src: []string{
				dbapi.ConnectorStatusPhaseAssigning,
			},
			Dst: dbapi.ConnectorStatusPhaseAssigned,
		},
		{
			Name: string(UpdateConnector),
			Src: []string{
				dbapi.ConnectorStatusPhaseAssigned,
			},
			Dst: dbapi.ConnectorStatusPhaseUpdating,
		},
	},
	dbapi.ConnectorNamespacePhaseDeleting: {
		{
			Name: string(DeleteConnector),
			Src: []string{
				dbapi.ConnectorStatusPhaseDeleting,
			},
			Dst: dbapi.ConnectorStatusPhaseDeleting,
		},
	},
}

func NewConnectorFSM(namespace *dbapi.ConnectorNamespace, connector *dbapi.Connector) *ConnectorFSM {
	return &ConnectorFSM{
		Connector:      connector,
		namespacePhase: namespace.Status.Phase,
		fsm:            fsm.NewFSM(string(connector.Status.Phase), connectorEvents[namespace.Status.Phase], nil),
	}
}

// Perform tries to perform the given operation and updates the connector phase,
// first return value is true if the phase was changed and
// second value is an error if operation is not permitted in connector's present phase
func (c *ConnectorFSM) Perform(operation ConnectorOperation) (bool, *errors.ServiceError) {
	// make sure FSM phase is current
	c.fsm.SetState(string(c.Connector.Status.Phase))
	if err := c.fsm.Event(string(operation)); err != nil {
		return false, errors.BadRequest("Cannot perform operation %q on Connector in Namespace in phase %q because %s",
			operation, c.namespacePhase, err)
	}

	newState := c.fsm.Current()
	updated := c.Connector.Status.Phase != newState
	c.Connector.Status.Phase = newState

	return updated, nil
}

// PerformConnectorOperation is a utility method to change a connector's phase
// first return value is true if the phase was changed and
// second value is an error if operation is not permitted in connector's present phase
func PerformConnectorOperation(namespace *dbapi.ConnectorNamespace, connector *dbapi.Connector, operation ConnectorOperation) (bool, *errors.ServiceError) {
	return NewConnectorFSM(namespace, connector).Perform(operation)
}
