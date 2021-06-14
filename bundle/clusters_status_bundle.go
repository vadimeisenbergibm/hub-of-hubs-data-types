package bundle

import (
	clusterv1 "github.com/open-cluster-management/api/cluster/v1"
	"k8s.io/apimachinery/pkg/types"
)

func NewClustersStatusBundle() *ClustersStatusBundle {
	return &ClustersStatusBundle{
		Objects:            make([]clusterv1.ManagedCluster, 0),
		DeletedObjectsUIDs: make([]types.UID, 0),
	}
}

type ClustersStatusBundle struct {
	Objects            []clusterv1.ManagedCluster `json:"objects"`
	DeletedObjectsUIDs []types.UID                `json:"deletedObjectsUIDs"`
}

func (bundle *ClustersStatusBundle) AddObject(cluster clusterv1.ManagedCluster) {
	bundle.Objects = append(bundle.Objects, cluster)
}

func (bundle *ClustersStatusBundle) AddDeletedObjectUID(clusterUID types.UID) {
	bundle.DeletedObjectsUIDs = append(bundle.DeletedObjectsUIDs, clusterUID)
}
