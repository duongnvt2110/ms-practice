package dto

type RegisterRequestForm struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	Username     string `json:"username" validate:"required,max=100"`
	Avatar       string `json:"avatar" validate:"omitempty,url"`
	Birthday     string `json:"birthday" validate:"required,datetime=2006-01-02"`
	MobileNumber string `json:"mobile_number" validate:"required"`
}
