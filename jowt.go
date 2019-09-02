package jowt

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	// HS512alg is used to implement HS512 Algorithm
	HS512alg string = "HS512"
	// HS256alg is used to implement HS512 Algorithm
	HS256alg string = "HS256"
	// typeJWT is for fill the typ in header data
	typeJWT = "JWT"
)

type (
	head struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}
	// Payload is a JWT data
	Payload map[string]interface{}
	// JWT struct is used to store header and payload data to be created as a JWT Token Based on the first function name.
	JWT struct {
		head
		Payload
		secretKey string
	}
	// VerifyToken struct is used to store header and payload data that has been given from Verify function.
	VerifyToken struct {
		head
		Payload
		encHeader    string
		encPayload   string
		encSignature string
		secretKey    string
	}

	// Security is used to store JWT to be encrypted.
	Security struct {
		SecretKey         string
		WhiteListURI      []string
		BlackListURI      []string
		Message           interface{}
		MiddlewarePayload map[string]interface{}
	}
)

// middlewarePayload :
var middlewarePayload map[string]interface{}

// SetPayloadFromMiddleware :
func (s *Security) SetPayloadFromMiddleware(DecrypedPayload Payload) {
	var stepOne map[string]interface{}
	stepOne = DecrypedPayload
	s.MiddlewarePayload = stepOne
}

// HS512 is used to start make an easy HS512 JWT Token
func HS512(secret string) *JWT {
	newJWT := new(JWT)
	newJWT.secretKey = secret
	newJWT.setHeader(head{
		Alg: HS512alg,
		Typ: typeJWT,
	})
	return newJWT
}

// HS256 is used to start make an easy HS512 JWT Token
func HS256(secret string) *JWT {
	newJWT := new(JWT)
	newJWT.secretKey = secret
	newJWT.setHeader(head{
		Alg: HS256alg,
		Typ: typeJWT,
	})
	return newJWT
}

func (jwt *JWT) setHeader(h head) *JWT {
	jwt.head.Alg = h.Alg
	jwt.head.Typ = h.Typ
	return jwt
}

// SetPayload is used to fill data into the payload to be encrypted,
// this function only can be passed using payload type from this library.
func (jwt *JWT) SetPayload(data Payload) *JWT {
	jwt.Payload = data
	return jwt
}

// Get is used to generate a JWT token.
func (jwt *JWT) Get() (string, bool) {
	head, err := json.Marshal(&jwt.head)
	if err != nil {
		return "", false
	}
	payload, err := json.Marshal(&jwt.Payload)
	if err != nil {
		return "", false
	}

	token := string(encode(string(head)) + "." + encode(string(payload)))
	fixToken := token + "." + encode(string(encrypt(jwt.head.Alg, jwt.secretKey, token)))

	return fixToken, true
}

// Verify is used to verify your last JWT token based on secret key.
func Verify(secret string) *VerifyToken {
	verifyToken := new(VerifyToken)
	verifyToken.secretKey = secret
	return verifyToken
}

// SetToken is used to fill JWT token to be filled to VerifyToken Struct.
func (verif *VerifyToken) SetToken(token string) *VerifyToken {
	split := strings.Split(token, ".")
	if len(split) != 3 {
		return verif
	}
	header := decode(split[0])
	payload := decode(split[1])
	json.Unmarshal([]byte(header), &verif.head)
	json.Unmarshal([]byte(payload), &verif.Payload)
	verif.encHeader = split[0]
	verif.encPayload = split[1]
	verif.encSignature = split[2]
	return verif
}

// Status function is used to get a status from JWT Token is original or fake.
func (verif *VerifyToken) Status() bool {
	if (verif.encHeader == "") || (verif.encPayload == "") || (verif.encSignature == "") || (verif.secretKey == "") {
		return false
	}
	token := string(verif.encHeader + "." + verif.encPayload)
	generatedToken := encrypt(verif.head.Alg, verif.secretKey, token)
	decodedSignature := decode(verif.encSignature)
	isSame := hmac.Equal(generatedToken, []byte(decodedSignature))
	return isSame
}

func encrypt(alg, secret, token string) []byte {
	switch alg {
	case HS512alg:
		{
			h := hmac.New(sha512.New, []byte(secret))
			h.Write([]byte(token))
			return h.Sum(nil)
		}
	case HS256alg:
		{
			h := hmac.New(sha256.New, []byte(secret))
			h.Write([]byte(token))
			return h.Sum(nil)
		}
	}
	return []byte("")
}

func encode(d string) string {
	ceking := base64.URLEncoding.EncodeToString([]byte(d))
	r := strings.NewReplacer("+", "", "/", "")
	result := r.Replace(ceking)
	return strings.TrimRight(result, "=")
}

func decode(d string) string {
	r := strings.NewReplacer("-_", "+/")
	result := r.Replace(d)
	length := len(result) % 4
	if length > 0 {
		result += strings.Repeat("=", 3)
	}
	decoded, _ := base64.URLEncoding.DecodeString(result)
	return string(decoded)
}

// JWTMiddleware is implemented from middleware. this function are implemented with http.Handle param and it will return http.Handler
func (s *Security) JWTMiddleware(next http.Handler) http.Handler {
	// Importing from Middleware Package
	return Jwt(s, next)
}
