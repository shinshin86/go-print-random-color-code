package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	bytes := make([]byte, 3)

	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}

	fmt.Println("#" + hex.EncodeToString(bytes))
}
