package status

import v1 "github.com/open-cluster-management/governance-policy-propagator/pkg/apis/policy/v1"

// ClustersPerPolicy struct holds the base information about clusters per policy.
type ClustersPerPolicy struct {
	PolicyID          string               `json:"policyId"`
	Clusters          []string             `json:"clusters"`
	RemediationAction v1.RemediationAction `json:"remediationAction"`
	ResourceVersion   string               `json:"resourceVersion"`
}

// BaseClustersPerPolicyBundle is the base struct for clusters per policy bundle.
type BaseClustersPerPolicyBundle struct {
	Objects     []*ClustersPerPolicy `json:"objects"`
	LeafHubName string               `json:"leafHubName"`
	Generation  uint64               `json:"generation"`
}
