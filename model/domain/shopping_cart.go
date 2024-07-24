package domain

import "time"

type ShoppingCart struct {
	Id        int        `json:"id"`
	UserId    int        `json:"user_id"`
	Items     []CartItem `json:"items"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
