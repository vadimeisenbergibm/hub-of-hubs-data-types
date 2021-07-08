package status

type PolicyComplianceStatus struct {
	PolicyId                  string   `json:"policyId"`
	NonCompliantClusters      []string `json:"nonCompliantClusters"`
	UnknownComplianceClusters []string `json:"unknownComplianceClusters"`
	ResourceVersion           string   `json:"resourceVersion"`
}

type BaseComplianceStatusBundle struct {
	Objects              []*PolicyComplianceStatus `json:"objects"`
	LeafHubName          string                    `json:"leafHubName"`
	BaseBundleGeneration uint64                    `json:"baseBundleGeneration"`
	Generation           uint64                    `json:"generation"`
}
