package main

import (
	"fmt"
	"time"

	paseto "aidanwoods.dev/go-paseto/v2"
)

func main() {

	// Generate key pair
	privateKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := privateKey.Public()

	// Buat token
	token := paseto.NewToken()

	token.SetIssuedAt(time.Now())
	token.SetExpiration(time.Now().Add(1 * time.Hour))

	token.SetString("username", "azim")
	token.SetString("role", "admin")

	// Sign token
	signed := token.V4Sign(privateKey, nil)

	fmt.Println("TOKEN:")
	fmt.Println(signed)

	fmt.Println()

	// Parse token
	parser := paseto.NewParser()

	parsed, err := parser.ParseV4Public(
		publicKey,
		signed,
		nil,
	)

	if err != nil {
		panic(err)
	}

	// Ambil data
	username, _ := parsed.GetString("username")
	role, _ := parsed.GetString("role")

	fmt.Println("HASIL PARSE:")
	fmt.Println("Username:", username)
	fmt.Println("Role:", role)
}