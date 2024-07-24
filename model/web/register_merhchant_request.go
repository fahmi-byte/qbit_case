package web

type RegisterMerchantRequest struct {
	UserId  int    `json:"user_id"`
	Name    string `validate:"required" json:"name"`
	Address string `validate:"required" json:"address"`
	City    string `validate:"required" json:"city"`
}
