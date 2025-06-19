package main

import "strings"

func containsAny(s string, words []string) bool {
	for _, word := range words {
		if strings.Contains(strings.ToLower(s), word) {
			return true
		}
	}
	return false
}

func replaceWords(s string, words []string) string {
	for _, word := range strings.Split(s, " ") {
		masked := strings.Repeat("*", 4)

		if containsAny(strings.ToLower(word), words) {
			s = strings.Replace(s, word, masked, -1)
		}

	}
	return s
}
