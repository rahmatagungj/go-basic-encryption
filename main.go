package main

import (
	"fmt"

	encrypt "github.com/rahmatagungj/encrypt/formula"
)

func main() {
	test := encrypt.Encrypt("12", "rahmat")
	fmt.Println(test)
}
