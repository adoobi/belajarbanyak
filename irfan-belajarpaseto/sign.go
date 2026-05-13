package main

import (
	"encoding/hex"
	"fmt"
	"time"

	paseto "aidanwoods.dev/go-paseto"
)

func main() {

	fmt.Println("CREATING TOKEN...")

	token := paseto.NewToken()

	token.SetIssuedAt(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))

	token.SetString("user-id", "user-123")

	// ENCRYPT TOKEN DISINI
	// key := paseto.NewV4SymmetricKey() // don't share this!!
	// encrypted := token.V4Encrypt(key, nil)

	// ENCRYPT PUBLIC
	secretKey := paseto.NewV4AsymmetricSecretKey() // don't share this!!!
	publicKey := secretKey.Public()                // DO share this one

	signed := token.V4Sign(secretKey, nil)

	fmt.Println("\nTOKEN CREATED (not encrypted yet)")
	fmt.Println(token)

	// fmt.Println("\ENCRYPT KEY")
	// fmt.Println(encrypted)

	fmt.Println("\nPUBLIC KEY")
	fmt.Println(hex.EncodeToString(publicKey.ExportBytes()))

	fmt.Println("\nSIGNED TOKEN:")
	fmt.Println(signed)

}
