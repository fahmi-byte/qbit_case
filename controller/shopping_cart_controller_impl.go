package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"qbit_case/helper"
	"qbit_case/model/domain"
	"qbit_case/model/web"
	"qbit_case/service"
	"strconv"
)

type ShoppingCartControllerImpl struct {
	Validate            *validator.Validate
	ShoppingCartService service.ShoppingCartService
}

func NewShoppingCartController(validate *validator.Validate, shoppingCartService service.ShoppingCartService) *ShoppingCartControllerImpl {
	return &ShoppingCartControllerImpl{Validate: validate, ShoppingCartService: shoppingCartService}
}

func (controller *ShoppingCartControllerImpl) GetAllItemsCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claims := request.Context().Value("claims").(domain.JWTAuthClaims)
	userId, err := strconv.Atoi(claims.Id)
	helper.PanicIfError(err)

	cartItems := controller.ShoppingCartService.GetAllItemsCartData(request.Context(), userId)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartItems,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ShoppingCartControllerImpl) AddCartItem(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartItemRequest := web.CartItemRequest{}
	helper.ReadFromRequest(request, &cartItemRequest)
	err := controller.Validate.Struct(cartItemRequest)
	helper.PanicIfError(err)

	claims := request.Context().Value("claims").(domain.JWTAuthClaims)
	userId, err := strconv.Atoi(claims.Id)
	cartItemRequest.UserId = userId

	err = controller.ShoppingCartService.AddCartItemData(request.Context(), cartItemRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Added item to cart successfully!",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ShoppingCartControllerImpl) DeleteCartItem(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	deleteCartItemRequest := web.DeleteCartItemRequest{}
	helper.ReadFromRequest(request, &deleteCartItemRequest)
	err := controller.Validate.Struct(deleteCartItemRequest)
	helper.PanicIfError(err)

	claims := request.Context().Value("claims").(domain.JWTAuthClaims)
	userId, err := strconv.Atoi(claims.Id)
	deleteCartItemRequest.UserId = userId

	err = controller.ShoppingCartService.DeleteCartItemData(request.Context(), deleteCartItemRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Delete item from cart successfully!",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ShoppingCartControllerImpl) UpdateCartItem(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateCartItemRequest := web.UpdateCartItemRequest{}
	helper.ReadFromRequest(request, &updateCartItemRequest)
	err := controller.Validate.Struct(updateCartItemRequest)
	helper.PanicIfError(err)

	claims := request.Context().Value("claims").(domain.JWTAuthClaims)
	userId, err := strconv.Atoi(claims.Id)
	updateCartItemRequest.UserId = userId

	err = controller.ShoppingCartService.UpdateCartItemData(request.Context(), updateCartItemRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Update item from cart successfully!",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
