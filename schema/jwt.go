package schema

import (
	"github.com/golang-jwt/jwt/v5"
	"slices"
)

const (
	ROLES_KEY    = "roles"
	ADMIN_ROLE   = "admin"
	TRAINER_ROLE = "trainer"
)

type JwtClaims struct {
	RealmAccess map[string][]string `json:"realm_access"`
	jwt.RegisteredClaims
}

func (jc JwtClaims) isRole(roleKey string) bool {
	if roles, ok := jc.RealmAccess[ROLES_KEY]; ok {
		if slices.Contains(roles, roleKey) {
			return true
		}
	}
	return false
}

func (jc JwtClaims) IsTrainer() bool {
	return jc.isRole(TRAINER_ROLE)
}

func (jc JwtClaims) IsAdmin() bool {
	return jc.isRole(ADMIN_ROLE)
}
