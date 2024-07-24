package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"qbit_case/constant"
	"qbit_case/helper"
	"qbit_case/model/domain"
	"qbit_case/model/web"
	"qbit_case/service"
	"strconv"
)

type MerchantControllerImpl struct {
	Validate        *validator.Validate
	MerchantService service.MerchantService
}

func NewMerchantController(validate *validator.Validate, merchantService service.MerchantService) *MerchantControllerImpl {
	return &MerchantControllerImpl{Validate: validate, MerchantService: merchantService}
}

func (controller *MerchantControllerImpl) RegisterMerchant(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	registerMerchantRequest := web.RegisterMerchantRequest{}
	helper.ReadFromRequest(request, &registerMerchantRequest)
	err := controller.Validate.Struct(registerMerchantRequest)
	helper.PanicIfError(err)

	claims := request.Context().Value("claims").(domain.JWTAuthClaims)
	role := claims.Role
	if role != constant.ROLE[constant.USER] {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	userId, err := strconv.Atoi(claims.Id)
	helper.PanicIfError(err)
	registerMerchantRequest.UserId = userId

	err = controller.MerchantService.RegisterMerchant(request.Context(), registerMerchantRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Register Merchant Successfully!",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MerchantControllerImpl) AddProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	addProductRequest := web.AddProductRequest{}
	helper.ReadFromRequest(request, &addProductRequest)
	err := controller.Validate.Struct(addProductRequest)
	helper.PanicIfError(err)

	claims := request.Context().Value("claims").(domain.JWTAuthClaims)
	role := claims.Role
	if role != constant.ROLE[constant.MERCHANT] {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	merchantId, err := strconv.Atoi(claims.MerchantId)
	helper.PanicIfError(err)

	fmt.Println(merchantId, "apa isi nya")

	addProductRequest.MerchantId = merchantId

	err = controller.MerchantService.AddProductData(request.Context(), addProductRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Create Product Successfully!",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
