package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type OrderController interface {
	NewOrder(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetUserOrders(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	PaymentGatewayCallback(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
