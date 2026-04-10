package main

import (
	"fmt"
	"strings"
)

func FixArticles(text string) string {
	word := strings.Fields(text)

	for i := 0; i < len(word); i++ {

		if word[i] == "a" && strings.ContainsRune("aeiouhAEIOUH", rune(word[i+1][0])) {
			word[i] = "an"
		} else if word[i] == "A" && strings.ContainsRune("aeiouhAEIOUH", rune(word[i+1][0])) {
			word[i] = "An"
		}
	}

	return strings.Join(word, " ")
}
