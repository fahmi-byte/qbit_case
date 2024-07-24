package middleware

import (
	"context"
	"database/sql"
	"net/http"
	"qbit_case/helper"
	"qbit_case/model/repository"
	"qbit_case/model/web"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
	DB      *sql.DB
	crypto  repository.CryptoRepository
}

func NewAuthMiddleware(handler http.Handler, DB *sql.DB, crypto repository.CryptoRepository) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler, DB: DB, crypto: crypto}
}

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	url := request.URL.String()

	if strings.HasPrefix(url, "/api/auth") || url == "/api/orders/payment-callback" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		ctx := middleware.TokenValidationMiddleware(writer, request)
		if ctx != nil {
			middleware.Handler.ServeHTTP(writer, request.WithContext(ctx))
		} else {
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			helper.WriteToResponseBody(writer, webResponse)
		}
	}

}

func (middleware AuthMiddleware) TokenValidationMiddleware(w http.ResponseWriter, r *http.Request) context.Context {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return nil
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims, err := middleware.crypto.ParseAndValidateJWT(tokenString)
	if err != nil {
		return nil
	}

	// Set claims in context
	ctx := context.WithValue(r.Context(), "claims", claims)
	return ctx
}
