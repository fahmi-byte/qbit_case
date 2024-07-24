package app

import (
	"github.com/julienschmidt/httprouter"
	"qbit_case/controller"
	"qbit_case/exception"
)

func NewRouter(authController controller.AuthController, merchantController controller.MerchantController, productController controller.ProductController, shoppingCartController controller.ShoppingCartController, orderController controller.OrderController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/auth/register", authController.Register)
	router.POST("/api/auth/login", authController.Login)
	router.POST("/api/merchant/register", merchantController.RegisterMerchant)
	router.POST("/api/merchant/add-product", merchantController.AddProduct)

	router.GET("/api/orders", orderController.GetUserOrders)
	router.POST("/api/orders", orderController.NewOrder)
	router.POST("/api/orders/payment-callback", orderController.PaymentGatewayCallback)

	router.GET("/api/shopping-cart", shoppingCartController.GetAllItemsCart)
	router.POST("/api/shopping-cart", shoppingCartController.AddCartItem)
	router.PUT("/api/shopping-cart", shoppingCartController.UpdateCartItem)
	router.DELETE("/api/shopping-cart", shoppingCartController.DeleteCartItem)

	router.GET("/api/products", productController.GetAllProducts)
	router.PanicHandler = exception.ErrorHandler

	return router
}
