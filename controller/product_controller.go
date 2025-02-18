package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ProductController interface {
	GetAllProducts(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
