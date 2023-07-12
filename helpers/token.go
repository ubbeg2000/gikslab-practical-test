package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var AppSecret []byte
var TokenBlacklist map[string]bool

type TokenClaims struct {
	jwt.StandardClaims
	UserID      uint64 `json:"user_id"`
	UserProfile string `json:"user_profile"`
}

func init() {
	AppSecret = []byte(os.Getenv("APP_SECRET"))
	TokenBlacklist = make(map[string]bool)
}

func GetClaims(c *fiber.Ctx) (*TokenClaims, error) {
	token := c.Query("token", "")
	if token == "" {
		return nil, errors.New("token not found")
	}

	t, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return AppSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(*TokenClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func CreateToken(userID uint64, userProfile string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    os.Getenv("APP_NAME"),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 3600*int64(time.Second),
			NotBefore: time.Now().Unix(),
			Id:        uuid.NewString(),
		},
		UserID:      userID,
		UserProfile: userProfile,
	}).SignedString(AppSecret)
}

func ParseToken(token string) (*TokenClaims, error) {
	t, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return AppSecret, nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	if !t.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := t.Claims.(*TokenClaims)
	if !ok {
		return nil, err
	}

	if _, ok := TokenBlacklist[claims.Id]; ok {
		return nil, errors.New("token blacklisted")
	}

	return claims, nil
}

func BlacklistToken(claims *TokenClaims) {
	TokenBlacklist[claims.Id] = true
}
