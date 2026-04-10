	words := strings.Fields(text)

	for i := len(words) - 1; i >= 0; i-- {
		if strings.HasPrefix(words[i], "(up,") && i+1 < len(words) {
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, _ := strconv.Atoi(numStr)

			for j := 1; j <= n; j++ {
				if i-j >= 0 {
					words[i-j] = strings.ToUpper(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)
			continue
		}
	}