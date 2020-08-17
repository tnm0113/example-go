package jwt

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

// GenJwkSignAndParse : gen rsa key, create jwk with public key, sign with private key, parse
func GenJwkSignAndParse() {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("failed to generate private key: %s\n", err)
		return
	}

	key, err := jwk.New(&privKey.PublicKey)
	key.Set(jwk.KeyUsageKey, "sig")
	jwk.AssignKeyID(key)
	if err != nil {
		log.Printf("failed to create JWK: %s", err)
		return
	}

	jsonbuf, err := json.MarshalIndent(key, "", "  ")
	if err != nil {
		log.Printf("failed to generate JSON: %s", err)
		return
	}


	var payload []byte
	{ // Create signed payload
		token := jwt.New()
		token.Set(`foo`, `bar`)
		payload, err = jwt.Sign(token, jwa.RS256, privKey)
		if err != nil {
			fmt.Printf("failed to generate signed payload: %s\n", err)
			return
		}
	}

	os.Stdout.Write(jsonbuf)

	{ // Parse signed payload
		// Use jwt.ParseVerify if you want to make absolutely sure that you
		// are going to verify the signatures every time
		token, err := jwt.Parse(bytes.NewReader(payload), jwt.WithVerify(jwa.RS256, &privKey.PublicKey))
		if err != nil {
			fmt.Printf("failed to parse JWT token: %s\n", err)
			return
		}
		buf, err := json.MarshalIndent(token, "", "  ")
		if err != nil {
			fmt.Printf("failed to generate JSON: %s\n", err)
			return
		}
		fmt.Printf("%s\n", buf)
	}
}

// VerifyUsingJwks : verify token from a jwks.json url
func VerifyUsingJwks() {
	payload := "eyJhbGciOiJSUzI1NiIsImtpZCI6InB1YmxpYzo5NzhhZjE5My0zMzRhLTQ2ZWYtYmQxYi03ZDUwNmYzM2M3NzQiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOltdLCJjbGllbnRfaWQiOiJteS1jbGllbnQiLCJleHAiOjE1OTc2MzA5NjgsImV4dCI6e30sImlhdCI6MTU5NzYyNzM2OCwiaXNzIjoiaHR0cDovLzEyNy4wLjAuMTo0NDQ0LyIsImp0aSI6IjgyYzdjZTNiLTI3NjQtNDEwYy1iN2ZhLTg0M2IwN2RmMDllMCIsIm5iZiI6MTU5NzYyNzM2OCwic2NwIjpbXSwic3ViIjoibXktY2xpZW50In0.wK4pid7ouyyaYX9y5CUyNZjpSF6OLK6XwYSxFC1E-wjHYxPz_-3hl7x2dCndlrajLcLRWAK5XUZgKQIxVphNViIF-znWaKdhiYGMauv25FtxZPWMFrnOZDYLOV-injK0m2HFAKDztZMhOjVzCwetPE6wqeou6fQ8yM1J-6_GRuinmnCf2jGQZTk2fe8vUXVKWMeDz590ZqcTk4M7ZT2jEKmMm8MOa77Eo9_SqoJyPt1k-E1xqgBRfexDf35D8fl0ny-vLgmtqmGpkVSFavFoQ1HJfgxM0CqaTy4k-9UQOJRCYl6nT52h313NVfAhiiaRO48zlePnItlE23LHWt2LlCiPrdGMQZPL60fVWKP2HoA5ahI8iZWLC1mYGJXdUMkiaXxKzyIp9QeiR2S3jprqAlpondNER9Ib4BDOwaVAOk8zQe56QrEBavHrGyugQgJqIhz13U-aGIEfiYGkJrS3EhVElSHSAa2F21CYCgciaobTUw-DtFtHs-UeLeLStrYzums87vEOPmXjFb13sAJFYc6WpKBMjfJX3nWG7Pyybpsa1JGhGOlDCqwtBT8lqLsDo1VzFfV7HB3hamhVClasAMSSLlFuM6wo5vYa6Ua51pdlV54Xfmpf7UINXtMTljA9i_B_LNjyw5EylM1fBrzj7hS9q8suFbsMtNGhSd_FN1c"

	set, err := jwk.Fetch("http://localhost:4444/.well-known/jwks.json")
	if err != nil {
		fmt.Printf("failed to parsed jwk : %s", err)
		return
	}

	token, err := jwt.Parse(bytes.NewReader([]byte(payload)), jwt.WithKeySet(set))
	if err != nil {
		fmt.Printf("failed to parse signed payload: %s\n", err)
		return
	}
	buf, err := json.MarshalIndent(token, "", "  ")
	fmt.Printf("%s\n", buf)
}
 
// CreateTokenWithClaims : create token with claims, get claims
func CreateTokenWithClaims() {
	const aLongLongTimeAgo = 233431200
	t := jwt.New()
	t.Set(jwt.SubjectKey, `https://github.com/lestrrat-go/jwx/jwt`)
	t.Set(jwt.AudienceKey, `Golang Users`)
	t.Set(jwt.IssuedAtKey, time.Unix(aLongLongTimeAgo, 0))
	t.Set(`privateClaimKey`, `Hello, World!`)

	buf, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		fmt.Printf("failed to generate JSON: %s\n", err)
		return
	}

	fmt.Printf("%s\n", buf)
	fmt.Printf("aud -> '%s'\n", t.Audience())
	fmt.Printf("iat -> '%s'\n", t.IssuedAt().Format(time.RFC3339))
	if v, ok := t.Get(`privateClaimKey`); ok {
		fmt.Printf("privateClaimKey -> '%s'\n", v)
	}
	fmt.Printf("sub -> '%s'\n", t.Subject())
}