package exception

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"qbit_case/helper"
	"qbit_case/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	if internalServerError(writer, request, err) {
		return
	}

}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	exceptionService, ok2 := err.(BadRequestError)

	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else if ok2 {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exceptionService.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true

	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	logger := logrus.New()
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0666)
	defer file.Close()
	logger.SetOutput(file)
	logger.Error(err)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err.(error).Error(),
	}

	helper.WriteToResponseBody(writer, webResponse)
	return true
}
