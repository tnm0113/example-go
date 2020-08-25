package main

import (
	"github.com/toannm/example-go/json"
	"github.com/toannm/example-go/jwt"
)

func main() {
	jwt.GenJwkSignAndParse()
	json.FileReadWrite()
}