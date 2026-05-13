package main

import (
	"fmt"

	paseto "aidanwoods.dev/go-paseto"
)

func main() {
	// 1. Public Key (Pasangan sah dari Private Key di atas)
	// Public key ini aman untuk disebar ke server lain untuk verifikasi.
	publicKeyHex := "32691e3baa8cd3a9a55ad1e37bb8fcc9e85b8853d987611f2eda891b0e50edbe"

	// 2. Token yang didapat dari token.go (Copy hasil output token.go ke sini)
	// Saya isi contoh token sementara (ganti dengan hasil running kamu nanti)
	signedToken := "v4.public.eyJleHAiOiIyMDI2LTA1LTEzVDEyOjI5OjIyKzA3OjAwIiwicm9sZSI6Im93bmVyIiwidXNlcm5hbWUiOiJhZG1pbiJ9Hxi4hhmODP1cEEPZqDOohjz-MlJzpZeP6YeFtlxfeZRJ8ygbC4k6imorWUtbxN0sw0q9Rb-fG_R66MnVWnf3BQ"

	// Ubah hex ke object public key
	publicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(publicKeyHex)
	if err != nil {
		fmt.Printf("Error: Public Key tidak valid: %v\n", err)
		return
	}

	// 3. Verifikasi Token
	parser := paseto.NewParser()
	token, err := parser.ParseV4Public(publicKey, signedToken, nil)

	if err != nil {
		fmt.Printf("❌ Verifikasi Gagal: %v\n", err)
	} else {
		fmt.Println("✅ Token Valid!")

		username, _ := token.GetString("username")
		role, _ := token.GetString("role")

		fmt.Printf("User: %s | Role: %s\n", username, role)
	}
}
