package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MerchantController interface {
	RegisterMerchant(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	AddProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
