package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"testing"
)

func TestSignMethod(t *testing.T) {
	method := jwt.GetSigningMethod("ES256")
	method = jwt.GetSigningMethod("ES384")
	method = jwt.GetSigningMethod("ES512")
	method = jwt.GetSigningMethod("HS256")
	method = jwt.GetSigningMethod("HS384")
	method = jwt.GetSigningMethod("HS512")
	method = jwt.GetSigningMethod("RS256")
	method = jwt.GetSigningMethod("RS384")
	method = jwt.GetSigningMethod("RS512")
	method = jwt.GetSigningMethod("PS256")
	method = jwt.GetSigningMethod("PS384")
	method = jwt.GetSigningMethod("PS512")
	fmt.Println(method.Alg())
}
