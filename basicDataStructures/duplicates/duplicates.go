package duplicates

func removeDuplicates(ss []string) []string {
	result := make([]string, 0, len(ss))
	string_set := make(map[string]struct{}, len(ss))
	for _, s := range ss {
		if _, ok := string_set[s]; !ok {
			string_set[s] = struct{}{}
			result = append(result, s)
		}
	}

	return result
}
