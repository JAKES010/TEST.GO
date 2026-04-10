func FixPunct(text string) string {
	re1 := regexp.MustCompile(`\text+([.,;?!']+)`)
	text = re1.ReplaceAllString(text, "$1")

	re2 := regexp.MustCompile(`([.,;?!']+)([A-Za-z])`)
	text = re2.ReplaceAllString(text, "$1 $2")

	text = strings.ReplaceAll(text, " ,", ",")
	text = strings.ReplaceAll(text, " .", ".")
	text = strings.ReplaceAll(text, " ?", "?")
	text = strings.ReplaceAll(text, " !", "!")
	text = strings.ReplaceAll(text, " :", ":")
	text = strings.ReplaceAll(text, " ;", ";")
	text = strings.ReplaceAll(text, " ...", "...")

	return text
}