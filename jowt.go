package jowt

import (
	"fmt"
)

const (
	// HS512 is used to implement HS512 Algorithm
	HS512 = "HS512"
	// HS314 is used to implement HS512 Algorithm
	HS314 = "HS314"
	// HS256 is used to implement HS512 Algorithm
	HS256 = "HS256"
)

type (
	// Head is a JWT Header
	Head struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}
	// Payload is a JWT data
	Payload map[string]interface{}

	// Signature is a JWT Signature Data
	Signature string

	// JWT is the real of unenctypted data JWT
	JWT struct {
		Head      Head
		Payload   Payload
		Signature Signature
	}

	// Secret is a string data to encrypt your jwt signature
	Secret string
)

// Fill is used to fill data into the payload to be encrypted, this function only can be passed using payload type from this library.
func (jwt *JWT) Fill(data Payload) {
	fmt.Println(data)
}

// Generate is used to make a JWT token.
func (jwt *JWT) Generate() {

}

// Verify is used to verify your last JWT token from secret key.
func Verify() {

}

func encrypt(d string) string {
	return d
}

func decrypt(d string) string {
	return d
}
