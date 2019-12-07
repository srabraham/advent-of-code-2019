package seanmath

// sliceDiff returns elements of a that are not in b
func SliceDiff(a, b []string) []string {
	bMap := make(map[string]bool)
	for _, e := range b {
		bMap[e] = true
	}
	result := make([]string, 0)
	for _, e := range a {
		if !bMap[e] {
			result = append(result, e)
		}
	}
	return result
}
