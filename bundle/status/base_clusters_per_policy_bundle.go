package status

import v1 "github.com/open-cluster-management/governance-policy-propagator/pkg/apis/policy/v1"

type ClustersPerPolicy struct {
	PolicyId          string               `json:"policyId"`
	Clusters          []string             `json:"clusters"`
	RemediationAction v1.RemediationAction `json:"remediationAction"`
	ResourceVersion   string               `json:"resourceVersion"`
}

type BaseClustersPerPolicyBundle struct {
	Objects     []*ClustersPerPolicy `json:"objects"`
	LeafHubName string               `json:"leafHubName"`
	Generation  uint64               `json:"generation"`
}
