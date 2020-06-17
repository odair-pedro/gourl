package slice

func Filter(slice []interface{}, condition func(interface{}) bool) []interface{} {
	var result []interface{}
	for _, s := range slice {
		if condition(s) {
			result = append(result, s)
		}
	}
	return result
}

func Any(slice []interface{}, condition func(interface{}) bool) bool {
	for _, s := range slice {
		if condition(s) {
			return true
		}
	}
	return false
}

func Sum(slice []interface{}, predicate func(interface{}) int) int {
	var result int
	for _, s := range slice {
		result += predicate(s)
	}
	return result
}
