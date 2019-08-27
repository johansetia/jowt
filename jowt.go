package jowt

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
)

const secretKey = "CONFIDENTAL_SECRET_KEY"

func main() {

}

// Fill is used to fill data into the payload to be encrypted.
func (jwt *JWT) Fill(data Payload) {

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
