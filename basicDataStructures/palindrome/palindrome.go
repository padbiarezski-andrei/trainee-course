package palindrome

import (
	"regexp"
	"strings"
	"unicode"
)

func Palindrome1(s string) bool {
	//regexp
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	clearedStr := nonAlphanumericRegex.ReplaceAllString(s, "")

	runeArr := []rune(strings.ToLower(clearedStr))
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

func Palindrome2Helper(s string) bool {
	//recursion
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return Palindrome2Helper(s[1 : len(s)-1])
}

func Palindrome2(s string) bool {
	// unicode
	clearedStr := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)

	return Palindrome2Helper(string(clearedStr))
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
