package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decimal(num string, base int) (int64, error) {
	return strconv.ParseInt(num, base, 64)
}

func capitaliseFirstWORD(cap1 string) string {
	return strings.ToUpper(string(cap1[0])) + strings.ToLower(cap1[1:])
}
func capitalise(cap3 string) string {
	words := strings.Fields(cap3)
	for i := range words {
		words[i] = strings.ToUpper(words[i][:1]) + strings.ToLower(words[i][1:])
	}
	return strings.Join(words, " ")
}
func capitaliseLastprevWord(cap2 []string, n int) []string {
	for i := len(cap2) - n; i < len(cap2); i++ {
		cap2[i] = strings.ToUpper(cap2[i])
	}
	return cap2
}
func isPunc(pun1 string) bool {
	if pun1 == "," || pun1 == "!" {
		return true
	} else {
		return false
	}
}
func isPunct(pun2 string) bool {
	return strings.ContainsAny(pun2, ".,!?;:")
}

func removespacefrompunc(s []string) string {
	res := strings.Join(s, " ")
	for _, p := range []string{" .", " ,", " !", " ?", " :", " ;"} {
		res = strings.ReplaceAll(res, p, p[1:])
	}
	return res
}

func article(word string) string {
	if strings.Contains("aeiouh", strings.ToLower(string(word[0]))) {
		return "an"
	}
	return "a"
}
func fixArticle(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" && strings.Contains("aeiou", strings.ToLower(string(words[i+1][0]))) {
			words[i] = "an"
		}
	}

	return strings.Join(words, " ")
}
func CapitalizeALL(cap1 string) string {
	return strings.Title(strings.ToLower(cap1))
}
func Is_Detect(word string) string {
	if strings.ContainsAny("!,.;'", word) {
		return "puntuation"
	} else {
		if strings.ContainsAny("0123456789", word) {
			return "number"
		}
		if strings.ContainsAny("ABCDEFGH", word) {
			return "letter"
		}
	}
	return word
}
func main() {
	fmt.Println(decimal("1E", 16))
	fmt.Println(decimal("1010", 2))
	fmt.Println(decimal("72", 8))
	fmt.Println(capitaliseFirstWORD("heLLo"))
	fmt.Println(capitaliseFirstWORD("woRLD"))
	fmt.Println(capitalise("david is a goat"))
	fmt.Println(capitaliseLastprevWord([]string{"it", "is", "very", "excitng"}, 2))
	fmt.Println(isPunc("."))
	fmt.Println(isPunc("!"))
	fmt.Println(isPunc(","))
	fmt.Println(isPunc("x"))
	fmt.Println(article("Apple"))
	fmt.Println(article("bpple"))
	fmt.Println(fixArticle("this is a apple which is a orange"))
	fmt.Println(capitalise("david is a goat"))
	fmt.Println(isPunct(","))
	fmt.Println(removespacefrompunc([]string{"Hello , World !"}))
	fmt.Println(CapitalizeALL("hello world"))
	fmt.Println(Is_Detect("!"))
	fmt.Println(Is_Detect("A"))
	fmt.Println(Is_Detect("7"))

}
