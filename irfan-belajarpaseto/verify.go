package main

import (
	"fmt"

	paseto "aidanwoods.dev/go-paseto"
)

func main() {

	publicKey, _ := paseto.NewV4AsymmetricPublicKeyFromHex("645766a0c56e018fb885b7fe26370f44e42972442e7e7523537669ba888c5026")

	signed := "v4.public.eyJleHAiOiIyMDI2LTA1LTEzVDE0OjE4OjIzKzA3OjAwIiwiaWF0IjoiMjAyNi0wNS0xM1QxMjoxODoyMyswNzowMCIsInVzZXItaWQiOiJ1c2VyLTEyMyJ9ysrTjjmrSgEJw4HtdZ3q1cfEL3DZBHqR2m_doN5_7keuc6da2lSGd63OrW-EwPCx8vXtcLKNBHSLJ16Ct5B8DA"

	parser := paseto.NewParser()

	token, err := parser.ParseV4Public(publicKey, signed, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("USER-ID:")
	fmt.Println(token.GetString("user-id"))
}
