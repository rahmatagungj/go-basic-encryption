package main

import (
	"fmt"

	"github.com/rahmatagungj/encrypt/formula"
)

func main() {
	secured := formula.Encrypt("rahmat")
	decrypted := formula.Decrypt(secured)

	fmt.Println("Ecnrypted: ", secured)
	fmt.Println("Decrypted: ", decrypted)
}
