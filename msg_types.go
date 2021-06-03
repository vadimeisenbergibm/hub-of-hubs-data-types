package hub_of_hubs_data_types

const (
	PolicyType = "PoliciesBundle"
	PlacementRuleType = "PlacementRulesBundle"
	PlacementBindingType = "PlacementBindingsBundle"
)

func GetBundleTypes() []string {
	var bundleTypes []string
	bundleTypes = append(bundleTypes, PolicyType)
	bundleTypes = append(bundleTypes, PlacementRuleType)
	bundleTypes = append(bundleTypes, PlacementBindingType)

}