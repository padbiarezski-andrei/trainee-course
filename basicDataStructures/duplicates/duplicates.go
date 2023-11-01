package duplicates

func removeDuplicates(ss []string) []string {
	res := make([]string, 0, len(ss))
	set := make(map[string]struct{}, len(ss))
	for _, s := range ss {
		if _, ok := set[s]; !ok { //not found
			set[s] = struct{}{}
			res = append(res, s)
		}
	}

	return res
}
