package helpers

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	*jwt.StandardClaims
	UserID      uint64 `json:"user_id"`
	UserProfile string `json:"user_profile"`
}

func CreateToken() (string, error) {
	return "", nil
}

func ParseToken(token string) (*TokenClaims, error) {
	return nil, nil
}

func IsTokenValid(t TokenClaims) bool {
	return false
}
