package status

import v1 "github.com/open-cluster-management/governance-policy-propagator/pkg/apis/policy/v1"

type MinimalPolicyComplianceStatus struct {
	PolicyId             string               `json:"policyId"`
	RemediationAction    v1.RemediationAction `json:"remediationAction"`
	NonCompliantClusters uint                 `json:"nonCompliantClusters"`
	AppliedClusters      uint                 `json:"appliedClusters"`
}

type BaseMinimalComplianceStatusBundle struct {
	Objects     []*MinimalPolicyComplianceStatus `json:"objects"`
	LeafHubName string                           `json:"leafHubName"`
	Generation  uint64                           `json:"generation"`
}
