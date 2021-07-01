package status

type ClustersPerPolicy struct {
	PolicyId        string   `json:"policyId"`
	Clusters        []string `json:"clusters"`
	ResourceVersion string   `json:"resourceVersion"`
}

type BaseClustersPerPolicyBundle struct {
	Objects     []*ClustersPerPolicy `json:"objects"`
	LeafHubName string               `json:"leafHubName"`
	Generation  uint64               `json:"generation"`
}
