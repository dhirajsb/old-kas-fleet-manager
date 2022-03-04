package phase

import (
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/connector/internal/api/dbapi"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
	"github.com/looplab/fsm"
)

type NamespaceOperation string

const (
	CreateNamespace     NamespaceOperation = "create"
	ConnectNamespace    NamespaceOperation = "connect"
	DisconnectNamespace NamespaceOperation = "disconnect"
	DeleteNamespace     NamespaceOperation = "delete"
)

// NamespaceFSM handles namespace phase changes within it's cluster's current phase
type NamespaceFSM struct {
	clusterPhase dbapi.ConnectorClusterPhaseEnum
	Namespace    *dbapi.ConnectorNamespace
	fsm          *fsm.FSM
}

var namespaceEvents = map[dbapi.ConnectorClusterPhaseEnum][]fsm.EventDesc{
	dbapi.ConnectorClusterPhaseDisconnected: {
		{
			Name: string(CreateNamespace),
			Src: []string{
				"", // phase default value from presenter
			},
			Dst: dbapi.ConnectorNamespacePhaseDisconnected,
		},
		{
			Name: string(DisconnectNamespace),
			Src: []string{
				dbapi.ConnectorNamespacePhaseDisconnected, dbapi.ConnectorNamespacePhaseReady,
			},
			Dst: dbapi.ConnectorNamespacePhaseDisconnected,
		},
		{
			Name: string(DeleteNamespace),
			Src: []string{
				dbapi.ConnectorNamespacePhaseDisconnected,
			},
			Dst: dbapi.ConnectorNamespacePhaseDeleting,
		},
	},
	dbapi.ConnectorClusterPhaseReady: {
		{
			Name: string(CreateNamespace),
			Src: []string{
				"",
			},
			Dst: dbapi.ConnectorNamespacePhaseDisconnected,
		},
		{
			Name: string(ConnectNamespace),
			Src: []string{
				dbapi.ConnectorNamespacePhaseDisconnected, dbapi.ConnectorNamespacePhaseReady,
			},
			Dst: dbapi.ConnectorNamespacePhaseReady,
		},
		{
			Name: string(ConnectNamespace),
			Src: []string{
				dbapi.ConnectorNamespacePhaseDeleting,
			},
			Dst: dbapi.ConnectorNamespacePhaseDeleting,
		},
		{
			Name: string(DisconnectNamespace),
			Src: []string{
				dbapi.ConnectorNamespacePhaseDisconnected, dbapi.ConnectorNamespacePhaseReady,
			},
			Dst: dbapi.ConnectorNamespacePhaseDisconnected,
		},
		{
			Name: string(DeleteNamespace),
			Src: []string{
				dbapi.ConnectorNamespacePhaseDisconnected, dbapi.ConnectorNamespacePhaseReady, dbapi.ConnectorNamespacePhaseDeleting,
			},
			Dst: dbapi.ConnectorNamespacePhaseDeleting,
		},
	},
	dbapi.ConnectorClusterPhaseDeleting: {
		{
			Name: string(ConnectNamespace),
			Src: []string{
				dbapi.ConnectorClusterPhaseDeleting,
			},
			Dst: dbapi.ConnectorNamespacePhaseDeleting,
		},
		{
			Name: string(DeleteNamespace),
			Src: []string{
				dbapi.ConnectorNamespacePhaseDeleting,
			},
			Dst: dbapi.ConnectorNamespacePhaseDeleting,
		},
	},
}

func NewNamespaceFSM(cluster *dbapi.ConnectorCluster, namespace *dbapi.ConnectorNamespace) *NamespaceFSM {
	return &NamespaceFSM{
		Namespace:    namespace,
		clusterPhase: cluster.Status.Phase,
		fsm:          fsm.NewFSM(string(namespace.Status.Phase), namespaceEvents[cluster.Status.Phase], nil),
	}
}

// Perform tries to perform the given operation and updates the namespace phase,
// first return value is true if the phase was changed and
// second value is an error if operation is not permitted in namespace's present phase
func (c *NamespaceFSM) Perform(operation NamespaceOperation) (bool, *errors.ServiceError) {
	// make sure FSM phase is current
	c.fsm.SetState(c.Namespace.Status.Phase)
	if err := c.fsm.Event(string(operation)); err != nil {
		switch err.(type) {
		case fsm.NoTransitionError:
			return false, nil
		default:
			return false, errors.BadRequest("Cannot perform operation %q on Namespace in Cluster in phase %q because %s",
				operation, c.clusterPhase, err)
		}
	}

	newState := c.fsm.Current()
	updated := c.Namespace.Status.Phase != newState
	c.Namespace.Status.Phase = newState

	return updated, nil
}

// PerformNamespaceOperation is a utility method to change a namespace's phase
// first return value is true if the phase was changed and
// second value is an error if operation is not permitted in namespace's present phase
func PerformNamespaceOperation(cluster *dbapi.ConnectorCluster, namespace *dbapi.ConnectorNamespace, operation NamespaceOperation) (bool, *errors.ServiceError) {
	return NewNamespaceFSM(cluster, namespace).Perform(operation)
}
