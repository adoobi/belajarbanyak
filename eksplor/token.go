package main

import (
	"fmt"
	"time"

	paseto "aidanwoods.dev/go-paseto"
)

func main() {
	// 1. Private Key (Secret) - Jangan disebar!
	// Ini adalah kunci untuk Menandatangani (Sign) token.
	privateKeyHex := "c22ae99b184baf6fc650b374ccab6c1246354a23c1911916318829e44fb5c84b32691e3baa8cd3a9a55ad1e37bb8fcc9e85b8853d987611f2eda891b0e50edbe"

	// Ubah hex ke object key
	privateKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKeyHex)
	if err != nil {
		panic(fmt.Sprintf("Private key tidak valid: %v", err))
	}

	// 2. Buat Token
	token := paseto.NewToken()
	token.SetString("username", "admin")
	token.SetString("role", "owner")
	token.SetExpiration(time.Now().Add(1 * time.Hour))

	// 3. Sign Token (Proses Tanda Tangan)
	signed := token.V4Sign(privateKey, nil)

	fmt.Println("--- TOKEN HASIL GENERATE ---")
	fmt.Println(signed)
	fmt.Println("----------------------------")
	fmt.Println("Gunakan token di atas untuk dimasukkan ke verify.go")
}
