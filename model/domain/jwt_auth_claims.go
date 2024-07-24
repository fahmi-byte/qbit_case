package domain

import "github.com/dgrijalva/jwt-go"

type JWTAuthClaims struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phone_number"`
	MerchantId  string `json:"merchant_id"`
	jwt.StandardClaims
}
