package domain

type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Role        string `json:"role"`
}
