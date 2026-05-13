package main

import (
	"fmt"

	paseto "aidanwoods.dev/go-paseto"
)

func main() {
	secretKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := secretKey.Public()

	fmt.Println(secretKey.ExportHex())
	fmt.Println(publicKey.ExportHex())
}
