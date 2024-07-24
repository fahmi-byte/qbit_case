package web

type RegisterRequest struct {
	Username    string `validate:"required" json:"username"`
	Email       string `validate:"required" json:"email"`
	Password    string `validate:"required" json:"password"`
	FullName    string `validate:"required" json:"full_name"`
	PhoneNumber string `validate:"required" json:"phone_number"`
	Address     string `validate:"required" json:"address"`
	City        string `validate:"required" json:"city"`
	RoleId      int    `validate:"required" json:"role_id"`
}
