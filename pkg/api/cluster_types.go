package api

import (
	"github.com/jinzhu/gorm"
)

type ClusterStatus string

func (k ClusterStatus) String() string {
	return string(k)
}

const (
	// The create cluster request has been recorder
	ClusterAccepted ClusterStatus = "cluster_accepted"
	// ClusterProvisioning the underlying ocm cluster is provisioning
	ClusterProvisioning ClusterStatus = "cluster_provisioning"
	// ClusterProvisioned the underlying ocm cluster is provisioned
	ClusterProvisioned ClusterStatus = "cluster_provisioned"
	// ClusterFailed the cluster failed to become ready
	ClusterFailed ClusterStatus = "failed"
	// ManagedKafkaAddonID the ID of the managed Kafka addon
	ManagedKafkaAddonID = "managed-kafka"
	// ClusterReady the cluster is terraformed and ready for kafka instances
	ClusterReady ClusterStatus = "ready"
	// ClusterAddonInstalled addon is installed on the cluster
	AddonInstalled ClusterStatus = "addon_installed"
	// KasFleetshardOperatorAddonId the ID of the kas-fleetshard-operator addon
	KasFleetshardOperatorAddonId = "kas-fleetshard-operator"
)

// This represents the valid statuses of a OSD cluster
var StatusForValidCluster = []string{string(ClusterProvisioning), string(ClusterProvisioned), string(ClusterReady), string(ClusterAccepted)}

type Cluster struct {
	Meta
	CloudProvider string        `json:"cloud_provider"`
	ClusterID     string        `json:"cluster_id"`
	ExternalID    string        `json:"external_id"`
	MultiAZ       bool          `json:"multi_az"`
	Region        string        `json:"region"`
	BYOC          bool          `json:"byoc"`
	Managed       bool          `json:"managed"`
	Status        ClusterStatus `json:"status"`
}

type ClusterList []*Cluster
type ClusterIndex map[string]*Cluster

func (c ClusterList) Index() ClusterIndex {
	index := ClusterIndex{}
	for _, o := range c {
		index[o.ID] = o
	}
	return index
}

func (org *Cluster) BeforeCreate(scope *gorm.Scope) error {
	if org.Status == "" {
		if err := scope.SetColumn("status", ClusterAccepted); err != nil {
			return err
		}
	}

	id := org.ID
	if id == "" {
		id = NewID()
	}

	return scope.SetColumn("ID", id)
}
