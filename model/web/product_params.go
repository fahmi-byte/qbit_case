package web

type ProductParams struct {
	Name       string  `schema:"name" json:"name"`
	City       string  `schema:"city" json:"city"`
	StartPrice float32 `schema:"start_price" json:"start_price"`
	EndPrice   float32 `schema:"end_price" json:"end_price"`
}
