package status

import "k8s.io/apimachinery/pkg/types"

type ClustersPerPolicy struct {
	PolicyId        types.UID `json:"policyId"`
	Clusters        []string  `json:"clusters"`
	ResourceVersion string    `json:"resourceVersion"`
}

type BaseClustersPerPolicyBundle struct {
	Objects     []*ClustersPerPolicy `json:"objects"`
	LeafHubName string               `json:"leafHubName"`
	Generation  uint64               `json:"generation"`
}
