package main

import (
	"fmt"
	"time"

	paseto "aidanwoods.dev/go-paseto"
)

func step1() {

	fmt.Println("CREATING TOKEN...")

	token := paseto.NewToken()

	token.SetIssuedAt(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))

	token.SetString("user-id", "user-123")

	// ENCRYPTE TOKEN DISINI
	key := paseto.NewV4SymmetricKey() // don't share this!!
	encrypted := token.V4Encrypt(key, nil)

	fmt.Println("TOKEN CREATED (not encrypted yet)")
	fmt.Println(token)

	fmt.Println("ENCRYPTED TOKEN")
	fmt.Println(encrypted)

}
