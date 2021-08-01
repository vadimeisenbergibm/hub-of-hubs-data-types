package status

// PolicyComplianceStatus struct holds information for policy compliance status.
type PolicyComplianceStatus struct {
	PolicyID                  string   `json:"policyId"`
	NonCompliantClusters      []string `json:"nonCompliantClusters"`
	UnknownComplianceClusters []string `json:"unknownComplianceClusters"`
	ResourceVersion           string   `json:"resourceVersion"`
}

// BaseComplianceStatusBundle is the base struct for compliance status bundle.
type BaseComplianceStatusBundle struct {
	Objects              []*PolicyComplianceStatus `json:"objects"`
	LeafHubName          string                    `json:"leafHubName"`
	BaseBundleGeneration uint64                    `json:"baseBundleGeneration"`
	Generation           uint64                    `json:"generation"`
}
