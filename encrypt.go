package main

import (
	"flag"
	"fmt"
)

func main() {

	textPtr := flag.String("t", "Hello, World!", "a text")
	keyPtr := flag.Int("k", 3, "an encryption key (shift)")

	flag.Parse()

	encrypted := encryptDecryptText(textPtr, *keyPtr, lowSlice, capitalSlice, 'e')

	fmt.Println(encrypted)
}
