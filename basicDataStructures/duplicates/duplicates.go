package duplicates

func removeDuplicates(ss []string) []string {
	result := make([]string, 0, len(ss))
	stringSet := make(map[string]struct{}, len(ss))
	for _, s := range ss {
		if _, ok := stringSet[s]; !ok {
			stringSet[s] = struct{}{}
			result = append(result, s)
		}
	}

	return result
}
