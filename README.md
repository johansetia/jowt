# JOWT (JWT Simple Package for Golang)
[![GoDoc](https://godoc.org/github.com/johansetia/jowt?status.svg)](https://godoc.org/github.com/johansetia/jowt)
[![Go Report Card](https://goreportcard.com/badge/github.com/johansetia/jowt)](https://goreportcard.com/report/github.com/johansetia/jowt)
[![Build Status](https://travis-ci.org/johansetia/jowt.svg?branch=master)](https://travis-ci.org/johansetia/jowt)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3126/badge)](https://bestpractices.coreinfrastructure.org/projects/3126)

JOWT is a Simple Package for Golang to generate and verify JWT Token using HMAC + RSA256 OR RSA512.

## 
## How to Install ?
```bash
$ go get github.com/johansetia/jowt
```
## How to Use ?
### 1. Importing Package
```go
package main

import (
    "github.com/johansetia/jowt"
)
```
### 2. Make a Payload and Key
you can use a structure from this package to create payload data.
```go
key := "qwertyuiopasdfghjklzxcvbnm123456"
data := jowt.Payload{
		"iss": "johan.com",
        "iat": 1566959060,
        "exp": 1598495060,
        "aud": "www.johan.com",
        "sub": "admin@johan.com",
        "UserName": "Johan",
        "Email": "johan@example.com",
        "userData": jowt.Payload{
                "Level":     "Administrator",
                "Role":      "UserAdministrator",
            },
	}
```
### 3. Getting value
Using **HS512 Function** to get a JWT Token
```go
    token, ok := jowt.HS512(key).SetPayload(dataPayload).Get()
    if !ok {
        panic("ERROR GET TOKEN")
    }
    fmt.Println(token)
    
```
Your HS512 token is 
```
eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImpvaGFuQGV4YW1wbGUuY29tIiwiVXNlck5hbWUiOiJKb2hhbiIsImF1ZCI6Ind3dy5qb2hhbi5jb20iLCJleHAiOjE1OTg0OTUwNjAsImlhdCI6MTU2Njk1OTA2MCwiaXNzIjoiam9oYW4uY29tIiwic3ViIjoiYWRtaW5Aam9oYW4uY29tIiwidXNlckRhdGEiOnsiTGV2ZWwiOiJBZG1pbmlzdHJhdG9yIiwiUm9sZSI6IlVzZXJBZG1pbmlzdHJhdG9yIn19.M1xhZDFeXwWj9dxVyMYGBgZ45NFqVobe8ZoPm6JVDBjSc6TSQCA-Ja9_DnIqzhNP1JMCMCdam5SIY6xn5ijLIw
```
Using **HS256 Function** to get a JWT Token
```go
    token, ok := jowt.HS256(key).SetPayload(dataPayload).Get()
    if !ok {
        panic("ERROR GET TOKEN")
    }
    fmt.Println(token)
    
```
Your HS256 token is 
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImpvaGFuQGV4YW1wbGUuY29tIiwiVXNlck5hbWUiOiJKb2hhbiIsImF1ZCI6Ind3dy5qb2hhbi5jb20iLCJleHAiOjE1OTg0OTUwNjAsImlhdCI6MTU2Njk1OTA2MCwiaXNzIjoiam9oYW4uY29tIiwic3ViIjoiYWRtaW5Aam9oYW4uY29tIiwidXNlckRhdGEiOnsiTGV2ZWwiOiJBZG1pbmlzdHJhdG9yIiwiUm9sZSI6IlVzZXJBZG1pbmlzdHJhdG9yIn19.yoFC0whrAh80A0bc7bGMlDR3XRW_dL-YwwEXQ3qRlJQ
```

### 4. Verify your JWT using key
```go
    verify := jowt.Verify(key).SetToken(jwt).Status()
    if verify {
        fmt.Println("USER ACCEPTED")
    }else{
        fmt.Println("USER DECIDED")
    }
```
## Structure
```bash
.
├── example
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── go.mod
├── jowt.go
├── LICENSE
└── README.md

1 directory, 7 files
```

