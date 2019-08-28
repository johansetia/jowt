package main

import (
	"fmt"

	"github.com/johansetia/jowt"
)

const key = "qwertyuiopasdfghjklzxcvbnm123456"

func main() {
	dataPayload := jowt.Payload{
		"iss":      "johan.com",
		"iat":      1566959060,
		"exp":      1598495060,
		"aud":      "www.johan.com",
		"sub":      "admin@johan.com",
		"UserName": "Johan",
		"Email":    "johan@johan.com",
		"userData": jowt.Payload{
			"Level": "Administrator",
			"Role":  "UserAdministrator",
		},
	}
	jwt, ok := jowt.HS256(key).SetPayload(dataPayload).Get()
	if !ok {
		panic("CANT CREATE")
	}
	fmt.Println("JWT TOKEN=======================================================")
	fmt.Println(jwt)
	fmt.Println("VERIFY JWT TOKEN=======================================================")
	verif := jowt.Verify(key).SetToken(jwt).Status()
	if !verif {
		panic("NOT VERIFIED")
	} else {
		panic("VERIFIED")
	}

}
