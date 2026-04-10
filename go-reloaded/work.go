package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Caseing(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "(up)" {
			if i > 0 {
				s[i-1] = strings.ToUpper(s[i-1])
			}
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
func Capitalize(s []string) []string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == "(low," && i+1 < len(s) {
			numstr := strings.TrimSuffix(s[i+1], ")")
			data, _ := strconv.Atoi(numstr)

			// 1. Calculate the range BEFORE deleting the tags
			start := i - data
			if start < 0 {
				start = 0
			}

			// 2. Lowercase the words leading up to index i
			for j := start; j < i; j++ {
				s[j] = strings.ToLower(s[j])
			}

			// 3. NOW delete the tags
			s = append(s[:i], s[i+2:]...)
		}
	}
	return s
}
func HexToDecimal(str []string) []string {
	for i := 0; i < len(str); i++ {
		if str[i] == "(hex)" && i > 0 {
			val, _ := strconv.ParseInt(str[i-1], 16, 64)
			str[i-1] = strconv.FormatInt(val, 10)
			str = append(str[:i], str[i+1:]...)
		}
	}
	return str
}
func main() {
	work := []string{"my", "name", "isn't", "(up)", "REALLY", "WHAT", "YOU", "(low,", "3)", "think", "it", "is", "your", "at", "least", "1F", "(hex)", "percent", "wrong"}
	fmt.Println(Caseing(work))
	fmt.Println(Capitalize(work))
	fmt.Println(HexToDecimal(work))
}
