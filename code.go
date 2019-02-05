package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	hash := md5.New()
	hash.Write([]byte("Corned Beef"))
	fmt.Println(hex.EncodeToString(hash.Sum(nil)))
}
