package domain

type Merchant struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	UserId  int    `json:"user_id"`
}
