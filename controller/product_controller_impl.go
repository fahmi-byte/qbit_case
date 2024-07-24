package controller

import (
	"github.com/gorilla/schema"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"qbit_case/helper"
	"qbit_case/model/web"
	"qbit_case/service"
)

type ProductControllerImpl struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{productService: productService}
}

func (controller *ProductControllerImpl) GetAllProducts(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	queryParams := web.ProductParams{}
	decoder := schema.NewDecoder()
	err := decoder.Decode(&queryParams, request.URL.Query())
	helper.PanicIfError(err)

	products := controller.productService.GetAllProductsData(request.Context(), queryParams)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   products,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
