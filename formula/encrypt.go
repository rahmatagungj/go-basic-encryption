package encrypt

import (
	"fmt"
)

func Encrypt(key, text string) string {
	var result string
	for i := 0; i < len(text); i++ {
		result += fmt.Sprintf("%x", text[i]^key[i%len(key)])
	}
	return result
}
