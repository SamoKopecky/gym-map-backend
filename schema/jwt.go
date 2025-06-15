package schema

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	RealmAccess map[string][]string `json:"realm_access"`
	jwt.RegisteredClaims
}
