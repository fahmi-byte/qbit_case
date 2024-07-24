package controller

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"qbit_case/constant"
	"qbit_case/helper"
	"qbit_case/model/domain"
	"qbit_case/model/web"
	"qbit_case/service"
	"strconv"
	"strings"
)

type PaymentConfig interface {
	AccessApiKey() string
	AccessMerchantCode() string
}

type OrderControllerImpl struct {
	Validate      *validator.Validate
	OrderService  service.OrderService
	PaymentConfig PaymentConfig
}

func NewOrderController(validate *validator.Validate, orderService service.OrderService, paymentConfig PaymentConfig) *OrderControllerImpl {
	return &OrderControllerImpl{Validate: validate, OrderService: orderService, PaymentConfig: paymentConfig}
}

func (controller *OrderControllerImpl) NewOrder(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderRequest := web.OrderRequest{}
	helper.ReadFromRequest(request, &orderRequest)
	err := controller.Validate.Struct(orderRequest)
	helper.PanicIfError(err)

	claims := request.Context().Value("claims").(domain.JWTAuthClaims)
	userId, err := strconv.Atoi(claims.Id)
	orderRequest.UserId = userId

	orderNumber, err := controller.OrderService.CreateNewOrder(request.Context(), orderRequest)
	helper.PanicIfError(err)

	merchantCode := controller.PaymentConfig.AccessMerchantCode()
	apiKey := controller.PaymentConfig.AccessApiKey()
	signature := helper.GenerateMD5Hash(merchantCode, orderNumber, strconv.Itoa(int(orderRequest.TotalAmount)), apiKey)

	inquiryRequest := web.PaymentGatewayInquiryRequest{
		MerchantCode:    merchantCode,
		PaymentAmount:   int(orderRequest.TotalAmount),
		PaymentMethod:   "M2",
		MerchantOrderId: orderNumber,
		ProductDetails:  "Pembayaran cicilan Kredit Plus",
		Email:           claims.Email,
		CustomerVaName:  claims.Username,
		CallbackUrl:     constant.CallbackUrl,
		ReturnUrl:       constant.CallbackUrl,
		Signature:       signature,
		ExpiryPeriod:    1,
	}

	requestBody, err := json.Marshal(inquiryRequest)
	if err != nil {
		http.Error(writer, "Failed to create request body", http.StatusInternalServerError)
		return
	}

	httpReq, err := http.NewRequest(http.MethodPost, constant.ApiUrlPaymentGateway, bytes.NewBuffer(requestBody))
	if err != nil {
		http.Error(writer, "Failed to create request", http.StatusInternalServerError)
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		http.Error(writer, "Request failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Baca respons body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(writer, "Failed to read response", http.StatusInternalServerError)
		return
	}

	// Parse respons JSON
	var inquiryResp web.InquiryResponse
	if err := json.Unmarshal(respBody, &inquiryResp); err != nil {
		http.Error(writer, "Failed to parse response", http.StatusInternalServerError)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   inquiryResp,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) GetUserOrders(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	claims := request.Context().Value("claims").(domain.JWTAuthClaims)
	userId, err := strconv.Atoi(claims.Id)
	helper.PanicIfError(err)

	orders := controller.OrderService.GetUserOrdersData(request.Context(), userId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orders,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) PaymentGatewayCallback(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	status := request.FormValue("resultCode")
	amountStr := request.FormValue("amount")
	amountFloat64, err := strconv.ParseFloat(amountStr, 32)
	helper.PanicIfError(err)

	amountFloat32 := float32(amountFloat64)
	orderNumber := request.FormValue("merchantOrderId")

	orderNumber = strings.TrimPrefix(orderNumber, "ORD-")

	orderId, err := strconv.Atoi(orderNumber)
	helper.PanicIfError(err)

	controller.OrderService.PaymentOrderCallback(request.Context(), orderId, status, amountFloat32)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
