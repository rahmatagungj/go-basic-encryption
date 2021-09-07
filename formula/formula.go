package formula

func charCodeAt(s string) rune {
	for _, r := range s {
		return r
	}
	return 0
}

func Encrypt(str string) []int {
	runes := []rune(str)

	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}

	return result
}

func Decrypt(str []int) string {
	var result []rune

	for i := 0; i < len(str); i++ {
		result = append(result, rune(str[i]))
	}

	return string(result)
}
