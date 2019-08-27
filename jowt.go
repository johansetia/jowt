package jowt

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	// HS512 is used to implement HS512 Algorithm
	HS512 = "HS512"
	// HS314 is used to implement HS512 Algorithm
	HS314 = "HS314"
	// HS256 is used to implement HS512 Algorithm
	HS256 = "HS256"
	typ   = "JWT"
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
		Head      *Head
		Payload   Payload
		Signature Signature
	}

	// EncyptedJWT is JWT data that has been encrypted
	EncyptedJWT struct {
		Head      string
		Payload   string
		Signature string
	}
)

var algorithm string
var secretKey string

// Make is used to start make an easy JWT Token
func Make(alg string) *JWT {
	var h Head
	h.Typ = typ
	switch alg {
	case HS512:
		{
			h.Alg = HS512
			algorithm = HS512
		}
	case HS314:
		{
			h.Alg = HS314
			algorithm = HS314
		}
	case HS256:
		{
			h.Alg = HS256
			algorithm = HS256
		}
	default:
		{
			panic("ERROR ALG NOT ALLOWED")
			os.Exit(71)
		}
	}
	return &JWT{}
}

// SetSecret is to fill the secret key string.
func (jwt *JWT) SetSecret(secret string) *JWT {

	secretKey = secret
	return jwt
}

// SetPayload is used to fill data into the payload to be encrypted, this function only can be passed using payload type from this library.
func (jwt *JWT) SetPayload(data Payload) {
	fmt.Println(data)
}

// Get is used to make a JWT token.
func (jwt *JWT) Get() string {
	data, err := json.Marshal(jwt.Head)
	if err != nil {
		panic("ERRRR")
	}
	return string(data)
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
