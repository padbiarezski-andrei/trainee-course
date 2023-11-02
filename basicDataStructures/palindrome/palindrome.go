package palindrome

import (
	"regexp"
	"strings"
	"unicode"
)

func Palindrome1(s string) bool {
	//regexp
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	nonAlphanumericRegex.ReplaceAllString(s, "")
	clearedStr := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)

	runeArr := []rune(clearedStr)
	if len(runeArr) < 2 {
		return true
	}
	l := 0
	r := len(runeArr) - 1
	for l < r {
		if runeArr[l] != runeArr[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func Palindrome2Helper(str string) bool {
	if len(str) <= 1 {
		return true
	}
	if str[0] != str[len(str)-1] {
		return false
	}
	return Palindrome2Helper(str[1 : len(str)-1])
}

func Palindrome2(s string) bool {
	result := make([]rune, 0)
	for _, ch := range s {
		if unicode.IsNumber(ch) || unicode.IsLetter(ch) {
			result = append(result, unicode.ToLower(ch))
		}
	}

	return Palindrome2Helper(string(result))
}

func Palindrome3(s string) bool {
	// unicode
	clearedStr := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)

	runeArr := []rune(clearedStr)
	if len(runeArr) < 2 {
		return true
	}
	l := 0
	r := len(runeArr) - 1
	for l < r {
		if runeArr[l] != runeArr[r] {
			return false
		}
		l++
		r--
	}

	return true
}
