package repository

import (
	"github.com/dgrijalva/jwt-go"
	"qbit_case/model/domain"
)

type CryptoRepository interface {
	GenerateHashPassword(password string) []byte
	CompareHashAndPassword(hash string, password string) bool
	ValidateWT(tokenString string) (*jwt.Token, error)
	GenerateJWT(user domain.User, merchantId string) (string, error)
	GenerateRefreshToken(user domain.User) (string, error)
	ParseAndValidateJWT(token string) (domain.JWTAuthClaims, error)
}
