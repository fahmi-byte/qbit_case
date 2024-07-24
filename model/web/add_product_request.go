package web

type AddProductRequest struct {
	MerchantId  int     `json:"merchant_id"`
	Name        string  `validate:"required" json:"name"`
	Description string  `validate:"required" json:"description"`
	Price       float32 `validate:"required" json:"price"`
	Discount    float32 `validate:"required" json:"discount"`
	Stock       int     `validate:"required" json:"stock"`
	CategoryId  int     `validate:"required" json:"category_id"`
}
