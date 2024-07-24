package repository

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"qbit_case/helper"
	"qbit_case/model/domain"
	"strconv"
	"time"
)

type CryptoConfig interface {
	AccessTokenSecret() string
	AccessTokenExpiresDate() time.Time
	RefreshTokenExpiresDate() time.Time
}

type CryptoRepositoryImpl struct {
	Config CryptoConfig
}

func NewCryptoRepository(config CryptoConfig) *CryptoRepositoryImpl {
	return &CryptoRepositoryImpl{Config: config}
}

func (repository *CryptoRepositoryImpl) GenerateHashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	helper.PanicIfError(err)
	return hashedPassword
}

func (repository *CryptoRepositoryImpl) CompareHashAndPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		panic(err)
	}
	return true
}

func (repository *CryptoRepositoryImpl) ValidateWT(tokenString string) (*jwt.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *CryptoRepositoryImpl) GenerateJWT(user domain.User, merchantId string) (string, error) {
	claims := &domain.JWTAuthClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   strconv.Itoa(user.Id),
			ExpiresAt: repository.Config.AccessTokenExpiresDate().Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
		// Add the user ID, username, email, user_clients and role to the JWT claims
		Id:          strconv.Itoa(user.Id),
		Username:    user.Username,
		Email:       user.Email,
		Role:        user.Role,
		PhoneNumber: user.PhoneNumber,
		MerchantId:  merchantId,
	}

	token, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(repository.Config.AccessTokenSecret()))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (repository *CryptoRepositoryImpl) GenerateRefreshToken(user domain.User) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *CryptoRepositoryImpl) ParseAndValidateJWT(token string) (domain.JWTAuthClaims, error) {
	parsedJwt := domain.JWTAuthClaims{}
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(repository.Config.AccessTokenSecret()), nil
	})
	if err != nil {
		return parsedJwt, err
	}
	if !parsedToken.Valid {
		return parsedJwt, errors.New("Token is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return parsedJwt, errors.New("Token is not valid")
	}

	payload := make(map[string]interface{})

	for key, value := range claims {
		payload[key] = value
	}

	parsedJwt.Id = payload["id"].(string)
	parsedJwt.Role = payload["role"].(string)
	parsedJwt.Email = payload["email"].(string)
	parsedJwt.Username = payload["username"].(string)
	parsedJwt.PhoneNumber = payload["phone_number"].(string)
	parsedJwt.MerchantId = payload["merchant_id"].(string)

	return parsedJwt, nil
}
