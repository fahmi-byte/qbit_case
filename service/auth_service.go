package service

import (
	"context"
	"qbit_case/model/web"
)

type AuthService interface {
	AuthRegister(ctx context.Context, request web.RegisterRequest) error
	AuthLogin(ctx context.Context, request web.LoginRequest) string
}
