package constant

const ApiUrlPaymentGateway = "https://sandbox.duitku.com/webapi/api/merchant/v2/inquiry"
const CallbackUrl = "https://6251-103-154-109-53.ngrok-free.app/api/orders/payment-callback"

const (
	ADMIN    = 1
	USER     = 2
	MERCHANT = 3
)

var ROLE = map[int]string{
	ADMIN:    "admin",
	USER:     "user",
	MERCHANT: "merchant",
}

const (
	SUCCESS = "00"
	FAILED  = "01"
)

var PAYMENT_STATUS = map[string]string{
	SUCCESS: "Success",
	FAILED:  "Failed",
}

const (
	PENDING    = "Pending"
	SHIPPED    = "Shipped"
	DELIVERED  = "Delivered"
	CANCELED   = "Canceled"
	SUCCESSFUL = "Success"
)
