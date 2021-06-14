package bundle

import (
	clusterv1 "github.com/open-cluster-management/api/cluster/v1"
)

func NewClustersStatusBundle() *ClustersStatusBundle {
	return &ClustersStatusBundle{
		Objects:        make([]clusterv1.ManagedCluster, 0),
		DeletedObjects: make([]clusterv1.ManagedCluster, 0),
	}
}

type ClustersStatusBundle struct {
	Objects        []clusterv1.ManagedCluster `json:"objects"`
	DeletedObjects []clusterv1.ManagedCluster `json:"deletedObjects"`
}

func (bundle *ClustersStatusBundle) AddObject(cluster clusterv1.ManagedCluster) {
	bundle.Objects = append(bundle.Objects, cluster)
}

func (bundle *ClustersStatusBundle) AddDeletedObject(cluster clusterv1.ManagedCluster) {
	bundle.DeletedObjects = append(bundle.DeletedObjects, cluster)
}
