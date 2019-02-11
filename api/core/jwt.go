package core

import (
	"os"

	"time"

	"github.com/dgrijalva/jwt-go"
	h "github.com/jmilagroso/api/helpers"
	m "github.com/jmilagroso/api/models"
)

// Reference: https://medium.com/@raul_11817/securing-golang-api-using-json-web-token-jwt-2dc363792a48

// JWTAuthenticationBackend struct
type JWTAuthenticationBackend struct {
}

const (
	tokenDuration = 72
	expireOffset  = 3600
)

var authBackendInstance *JWTAuthenticationBackend

// InitJWTAuthenticationBackend instance
func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{}
	}

	return authBackendInstance
}

// GenerateToken generates token
func (backend *JWTAuthenticationBackend) GenerateToken(userID string) m.JWT {

	exp := time.Now().Add(time.Hour * time.Duration(tokenDuration)).Unix()

	claims := jwt.MapClaims{
		"exp": exp,
		"iat": time.Now().Unix(),
		"sub": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("PRIVATE_KEY")))
	h.Error(err)

	return m.JWT{Token: ss, Expiration: exp}
}

// ValidateToken validates token
func (backend *JWTAuthenticationBackend) ValidateToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("PRIVATE_KEY")), nil
	})

	return token, err
}
