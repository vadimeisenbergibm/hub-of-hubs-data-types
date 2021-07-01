package status

import (
	"github.com/open-cluster-management/governance-policy-propagator/pkg/apis/policy/v1"
)

type PolicyComplianceStatus struct {
	PolicyId                  string               `json:"policyId"`
	NonCompliantClusters      []string             `json:"nonCompliantClusters"`
	UnknownComplianceClusters []string             `json:"unknownComplianceClusters"`
	RemediationAction         v1.RemediationAction `json:"remediationAction"`
	ResourceVersion           string               `json:"resourceVersion"`
}

type BaseComplianceStatusBundle struct {
	Objects              []*PolicyComplianceStatus `json:"objects"`
	LeafHubName          string                    `json:"leafHubName"`
	BaseBundleGeneration uint64                    `json:"baseBundleGeneration"`
	Generation           uint64                    `json:"generation"`
}
