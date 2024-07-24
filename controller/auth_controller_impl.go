package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"qbit_case/helper"
	"qbit_case/model/web"
	"qbit_case/service"
)

type AuthControllerImpl struct {
	Validate    *validator.Validate
	AuthService service.AuthService
}

func NewAuthController(validate *validator.Validate, authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{Validate: validate, AuthService: authService}
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	registerRequest := web.RegisterRequest{}
	helper.ReadFromRequest(request, &registerRequest)
	err := controller.Validate.Struct(registerRequest)
	helper.PanicIfError(err)

	err = controller.AuthService.AuthRegister(request.Context(), registerRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Create User Successfully!",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.LoginRequest{}
	helper.ReadFromRequest(request, &loginRequest)
	err := controller.Validate.Struct(loginRequest)
	helper.PanicIfError(err)

	token := controller.AuthService.AuthLogin(request.Context(), loginRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   token,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
