package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ShoppingCartController interface {
	GetAllItemsCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	AddCartItem(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteCartItem(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateCartItem(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
