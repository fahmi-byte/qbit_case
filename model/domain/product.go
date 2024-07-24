package domain

type Product struct {
	Id           int     `json:"id"`
	MerchantId   int     `json:"merchant_id"`
	MerchantName string  `json:"merchant_name"`
	MerchantCity string  `json:"merchant_city"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float32 `json:"price"`
	FinalPrice   float32 `json:"final_price"`
	Discount     float32 `json:"discount"`
	Stock        int     `json:"stock"`
	CategoryId   int     `json:"category_id"`
	CategoryName string  `json:"category_name"`
}
