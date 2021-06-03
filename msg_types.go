package datatypes

const (
	PolicyType = "PoliciesBundle"
	PlacementRuleType = "PlacementRulesBundle"
	PlacementBindingType = "PlacementBindingsBundle"
)

func GetBundleTypes() []string {
	return []string {PolicyType, PlacementRuleType, PlacementBindingType}
}