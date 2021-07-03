package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

const key = "asdbasduiaus"

func keyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(key), nil
}

func TestSignMethod(t *testing.T) {
	cfg := Config{
		Audience:            "test",
		Issuer:              "test",
		Expiration:          2 * time.Second,
		PrivateKey:          key,
		Method:              "HS256",
	}

	claims := cfg.Standardize()
	fmt.Println(claims)
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(key))
	if err != nil {
		t.Fatal(err)
	}

	_, err = jwt.ParseWithClaims(token, jwt.MapClaims{}, keyFunc)
	if err != nil {
		t.Fatal(err)
	}
}
