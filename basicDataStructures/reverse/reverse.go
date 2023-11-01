package reverse

// func reverse(s string) string {
// 	rune_arr := []rune(s)
// 	if len(rune_arr) < 2 {
// 		return s
// 	}
// 	rev := make([]rune, len(rune_arr), len(rune_arr))
// 	j := 0
// 	for i := len(rune_arr) - 1; i >= 0; i-- {
// 		rev[j] = rune_arr[i]
// 		j++
// 	}
// 	return string(rev)
// }

func reverse(s []rune) { //inplace
	if len(s) < 2 {
		return
	}

	l := 0
	r := len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

	return
}

func reverseWords1(s string) string {
	rune_arr := []rune(s)
	l := 0
	r := 0
	for i := 0; i < len(rune_arr); i++ {
		for i < len(rune_arr) && rune_arr[i] != ' ' { // char isn't space
			r++
			i++
		}

		reverse(rune_arr[l:r])

		for i < len(rune_arr) && rune_arr[i] == ' ' { // char is space
			r++
			i++
		}
		l = r
		r++
	}
	return string(rune_arr)
}

// func reverseWords2(s string) string {
// 	panic("not implemented")
// 	return ""
// }
